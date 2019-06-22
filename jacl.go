package jacl

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/yuce/jacl/parser"
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

func Unmarshal(text string, v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errors.New("non-nil pointer is required")
	}

	var rm map[string]interface{}
	// if v is map[string]interface{} decode directly in it.
	switch m := rv.Elem().Interface().(type) {
	case map[string]interface{}:
		rm = m
	default:
		rm = map[string]interface{}{}
	}

	lexer := parser.NewJaclLexer(antlr.NewInputStream(text))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewJaclParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	listener := newJaclListener(rm)
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Config())

	return nil
}
