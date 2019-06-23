package jacl

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"

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
		quoteIndex := strings.IndexFunc(text, func(r rune) bool {
			return r == '\'' || r == '"'
		})
		if quoteIndex < 0 {
			panic(fmt.Sprintf("invalid raw string at line: %d", startToken.GetLine()))
		}
		if quoteIndex > 0 {
			funName := text[:quoteIndex]
			switch funName {
			case "trim":
				// trim""" ... """
				text, err := trimText(text[7 : len(text)-3])
				if err != nil {
					panic(fmt.Errorf("invalid raw string at line: %d %s", startToken.GetLine(), err.Error()))
				}
				rl.currentValue = text
			case "pin":
				// pin""" ... ""
				text, err := pinTrimText(text[6 : len(text)-3])
				if err != nil {
					panic(fmt.Errorf("invalid raw string at line: %d %s", startToken.GetLine(), err.Error()))
				}
				rl.currentValue = text
			default:
				panic(fmt.Sprintf("invalid string function: '%s' at line: %d", funName, startToken.GetLine()))
			}
		} else {
			// strip raw string prefix and suffix
			rl.currentValue = text[3 : len(text)-3]
		}
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

func trimText(text string) (string, error) {
	lines := strings.Split(text, "\n")
	pinSet := false
	pinPos := 0
	newLines := make([]string, 0, len(lines))
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if len(trimmedLine) == 0 {
			if !pinSet {
				// skip this line if the pin wasn't set and its an empty line
				if len(trimmedLine) == 0 {
					continue
				}
			}
			// add the trimmed line instead of the original line
			newLines = append(newLines, trimmedLine)
			continue
		}

		leadingSpaces := countLeadingSpaces(line)
		if !pinSet {
			pinPos = leadingSpaces
			pinSet = true
		}
		if pinPos > leadingSpaces {
			// this line starts before the pin pos
			return "", errors.New("inconsistent line start")
		}
		if pinPos < leadingSpaces {
			leadingSpaces = pinPos
		}
		newLines = append(newLines, line[leadingSpaces:])
	}
	if len(newLines) > 0 {
		// traverse new lines in reverse to leave out empty lines at the end
		for lastIndex := len(newLines) - 1; lastIndex >= 0; lastIndex-- {
			// if a non-empty line is found return the result
			if newLines[lastIndex] != "" {
				return strings.Join(newLines[:lastIndex+1], "\n"), nil
			}
		}
	}

	// this is text with all whitespace
	return "", nil
}

func pinTrimText(text string) (string, error) {
	lines := strings.Split(text, "\n")
	pinSet := false
	pinPos := 0
	newLines := make([]string, 0, len(lines))
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if len(trimmedLine) == 0 {
			if !pinSet {
				// skip this line if the pin wasn't set and its an empty line
				if len(trimmedLine) == 0 {
					continue
				}
			}
			// pin trim the original line and add it
			if pinPos < len(line) {
				newLines = append(newLines, line[pinPos:])
			} else {
				newLines = append(newLines, trimmedLine)
			}
			continue
		}

		leadingSpaces := countLeadingSpaces(line)
		if !pinSet {
			if trimmedLine == "^" {
				pinPos = leadingSpaces
				pinSet = true
				continue
			} else {
				return "", errors.New("pin should be the first non-space character")
			}
		}
		if pinPos > leadingSpaces {
			// this line starts before the pin pos
			return "", errors.New("inconsistent line start")
		}
		if pinPos < leadingSpaces {
			leadingSpaces = pinPos
		}
		newLines = append(newLines, line[leadingSpaces:])
	}

	if !pinSet {
		return "", errors.New("no pin in text")
	}

	return strings.Join(newLines, "\n"), nil
}

func countLeadingSpaces(s string) int {
	// Adapted from: https://github.com/golang/go/blob/master/src/strings/strings.go
	// Fast path for ASCII: look for the first ASCII non-space byte
	start := 0
	for ; start < len(s); start++ {
		c := s[start]
		if c >= utf8.RuneSelf {
			// If we run into a non-ASCII byte, fall back to the
			// slower unicode-aware method on the remaining bytes
			index := strings.IndexFunc(s[start:], func(r rune) bool {
				return r != ' ' && r != '\t'
			})
			if index >= 0 {
				return start + index
			}
			return start
		}
		if c != ' ' && c != '\t' {
			break
		}
	}
	return start
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
		err = unmarshalStruct(rm, rv)
	}

	return err
}

func unmarshalStruct(rm map[string]interface{}, value reflect.Value) error {
	elem := value.Elem()
	elemType := elem.Type()
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
				switch field.Type.Kind() {
				case reflect.Struct:
					st := reflect.New(field.Type)
					unmarshalStruct(t, st)
					value.Set(st.Elem())
				case reflect.Map:
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
