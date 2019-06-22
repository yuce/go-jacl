package jacl

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/yuce/go-jacl/parser"
)

const maxStack = 16

type jaclMap map[string]interface{}

type jaclListener struct {
	*parser.BaseJaclListener
	stack        []interface{}
	stackTop     int
	currentValue interface{}
}

func newJaclListener(m map[string]interface{}) *jaclListener {
	stack := make([]interface{}, maxStack)
	stack[0] = m
	return &jaclListener{stack: stack}
}

func (rl jaclListener) Properties() map[string]interface{} {
	return rl.stack[0].(map[string]interface{})
}

func (rl *jaclListener) EnterMapLiteral(c *parser.MapLiteralContext) {
	rl.pushToStack(map[string]interface{}{})
}

func (rl *jaclListener) EnterArrayLiteral(c *parser.ArrayLiteralContext) {
	rl.pushToStack([]interface{}{})
}

func (rl *jaclListener) EnterLiteral(c *parser.LiteralContext) {
	startToken := c.GetStart()
	text := c.GetText()
	switch startToken.GetTokenType() {
	case parser.JaclParserFloatLiteral:
		value, err := strconv.ParseFloat(text, 64)
		if err != nil {
			panic(err)
		}
		rl.currentValue = value
	case parser.JaclParserIntegerLiteral:
		value, err := parseInteger(text)
		if err != nil {
			panic(err)
		}
		rl.currentValue = value
	case parser.JaclParserStringLiteral:
		// strip quotes
		text, err := strconv.Unquote(text)
		if err != nil {
			panic(err)
		}
		rl.currentValue = text
	case parser.JaclParserRawStringLiteral:
		// strip raw string prefix and suffix
		rl.currentValue = text[3 : len(text)-3]
	case parser.JaclParserBooleanLiteral:
		switch text {
		case "true":
			rl.currentValue = true
		case "false":
			rl.currentValue = false
		}
	}
}

func (rl *jaclListener) ExitSingleExpression(c *parser.SingleExpressionContext) {
	if container, ok := rl.stack[rl.stackTop].([]interface{}); ok {
		rl.stack[rl.stackTop] = append(container, rl.currentValue)
		rl.currentValue = nil
	}
}

func (rl *jaclListener) ExitMapLiteral(c *parser.MapLiteralContext) {
	rl.popFromStack()
}

func (rl *jaclListener) ExitArrayLiteral(c *parser.ArrayLiteralContext) {
	rl.popFromStack()
}

func (rl *jaclListener) ExitPropertyAssignment(ctx *parser.PropertyAssignmentContext) {
	propertyName := ctx.GetChild(0).GetPayload().(*antlr.BaseParserRuleContext).GetText()
	if strings.HasPrefix(propertyName, "\"") {
		propertyName = propertyName[1 : len(propertyName)-1]
	}
	if container, ok := rl.stack[rl.stackTop].(map[string]interface{}); ok {
		container[propertyName] = rl.currentValue
		rl.currentValue = nil
	} else {
		panic(fmt.Sprintf("unexpected type in EXIT JaclParserRULE_propertyAssignment: %s", reflect.TypeOf(rl.stack[rl.stackTop])))
	}
}

func (rl *jaclListener) pushToStack(item interface{}) {
	rl.stackTop++
	if rl.stackTop > maxStack {
		panic("stack overflow")
	}
	rl.stack[rl.stackTop] = item
}

func (rl *jaclListener) popFromStack() {
	rl.currentValue = rl.stack[rl.stackTop]
	rl.stackTop--
	if rl.stackTop < 0 {
		panic("stack underflow")
	}
}

func parseInteger(text string) (interface{}, error) {
	if len(text) > 2 {
		switch text[:2] {
		case "0b":
			return strconv.ParseUint(text[2:], 2, 64)
		case "0o":
			return strconv.ParseUint(text[2:], 8, 64)
		case "0d":
			return strconv.ParseUint(text[2:], 10, 64)
		case "0x":
			return strconv.ParseUint(text[2:], 16, 64)
		}
	}

	return strconv.ParseInt(text, 10, 64)
}

func Unmarshal(text string, v interface{}) (err error) {
	defer func(err *error) {
		if r := recover(); r != nil {
			switch rt := r.(type) {
			case error:
				*err = rt
			case string:
				*err = errors.New(rt)
			default:
				panic(rt)
			}
		}
	}(&err)
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errors.New("non-nil pointer is required")
	}

	var rm map[string]interface{}
	var isStruct bool

	// if v is map[string]interface{} decode directly in it.
	switch m := rv.Elem().Interface().(type) {
	case map[string]interface{}:
		rm = m
	default:
		if rv.Elem().Kind() != reflect.Struct {
			return fmt.Errorf("unmarshal error: can only unmarshal maps with string keys or structs, not: %s", reflect.TypeOf(rv.Elem()))
		}
		isStruct = true
		rm = map[string]interface{}{}
	}

	lexer := parser.NewJaclLexer(antlr.NewInputStream(text))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewJaclParser(stream)
	p.AddErrorListener(newErrorListener())
	p.BuildParseTrees = true
	listener := newJaclListener(rm)
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Config())

	if isStruct {
		err = unmarshalStruct(rm, v)
	}

	return err
}

func unmarshalStruct(rm map[string]interface{}, v interface{}) error {
	elem := reflect.ValueOf(v).Elem()
	elemType := reflect.ValueOf(v).Elem().Type()
	for i := 0; i < elemType.NumField(); i++ {
		value := elem.Field(i)
		if !value.CanSet() {
			continue
		}
		field := elemType.Field(i)
		propertyName := field.Tag.Get("jacl")
		if propertyName == "" {
			propertyName = field.Name
		}
		if configValue, ok := rm[propertyName]; ok {
			switch t := configValue.(type) {
			case string:
				value.SetString(t)
			case bool:
				value.SetBool(t)
			case int64:
				value.SetInt(t)
			case uint64:
				value.SetUint(t)
			case float64:
				value.SetFloat(t)
			case []interface{}:
				elemType := field.Type.Elem()
				switch elemType.Kind() {
				case reflect.Interface:
					// if this is a slice of interface{}, just assign it.
					value.Set(reflect.ValueOf(t))
				default:
					// otherwise create and assign a slice of the given type.
					sl := reflect.MakeSlice(field.Type, len(t), len(t))
					for i, tv := range t {
						sl.Index(i).Set(reflect.ValueOf(tv).Convert(elemType))
					}
					value.Set(sl)
				}
			case map[string]interface{}:
				elemType := field.Type.Elem()
				switch elemType.Kind() {
				case reflect.Interface:
					// if this is a map of interface{}, just assign it.
					value.Set(reflect.ValueOf(t))
				default:
					// otherwise create and assign a map of the given type.
					sm := reflect.MakeMapWithSize(field.Type, len(t))
					for tk, tv := range t {
						tkv := reflect.ValueOf(tk)
						tvv := reflect.ValueOf(tv).Convert(elemType)
						sm.SetMapIndex(tkv, tvv)
					}
					value.Set(sm)
				}
			default:
				return fmt.Errorf("jacl unmarshal error: don't know how to unmarshal field: '%s'", field.Name)
			}
		} else {
			return fmt.Errorf("jacl unmarshal error: property not found: '%s'", propertyName)
		}
	}

	return nil
}

type errorListener struct {
	*antlr.DefaultErrorListener
}

func newErrorListener() *errorListener {
	return new(errorListener)
}

func (el *errorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	panic(fmt.Sprintf("syntax error: %s line: %d column: %d symbol: %s", msg, line, column, offendingSymbol))
}
