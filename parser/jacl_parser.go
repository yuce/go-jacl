// Generated from parser/Jacl.g4 by ANTLR 4.7.

package parser // Jacl

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 25, 84, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 7, 2, 22, 10, 2, 12, 2, 14, 2,
	25, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 5, 3, 32, 10, 3, 7, 3, 34, 10,
	3, 12, 3, 14, 3, 37, 11, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 5, 4, 44, 10,
	4, 7, 4, 46, 10, 4, 12, 4, 14, 4, 49, 11, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3,
	5, 3, 5, 5, 5, 57, 10, 5, 7, 5, 59, 10, 5, 12, 5, 14, 5, 62, 11, 5, 3,
	5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 5, 7, 72, 10, 7, 3, 8, 3,
	8, 3, 8, 3, 8, 5, 8, 78, 10, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 2, 2,
	11, 2, 4, 6, 8, 10, 12, 14, 16, 18, 2, 4, 5, 2, 5, 5, 14, 15, 21, 22, 4,
	2, 14, 14, 25, 25, 2, 85, 2, 23, 3, 2, 2, 2, 4, 28, 3, 2, 2, 2, 6, 40,
	3, 2, 2, 2, 8, 52, 3, 2, 2, 2, 10, 65, 3, 2, 2, 2, 12, 71, 3, 2, 2, 2,
	14, 77, 3, 2, 2, 2, 16, 79, 3, 2, 2, 2, 18, 81, 3, 2, 2, 2, 20, 22, 5,
	10, 6, 2, 21, 20, 3, 2, 2, 2, 22, 25, 3, 2, 2, 2, 23, 21, 3, 2, 2, 2, 23,
	24, 3, 2, 2, 2, 24, 26, 3, 2, 2, 2, 25, 23, 3, 2, 2, 2, 26, 27, 7, 2, 2,
	3, 27, 3, 3, 2, 2, 2, 28, 35, 7, 8, 2, 2, 29, 31, 5, 14, 8, 2, 30, 32,
	7, 3, 2, 2, 31, 30, 3, 2, 2, 2, 31, 32, 3, 2, 2, 2, 32, 34, 3, 2, 2, 2,
	33, 29, 3, 2, 2, 2, 34, 37, 3, 2, 2, 2, 35, 33, 3, 2, 2, 2, 35, 36, 3,
	2, 2, 2, 36, 38, 3, 2, 2, 2, 37, 35, 3, 2, 2, 2, 38, 39, 7, 9, 2, 2, 39,
	5, 3, 2, 2, 2, 40, 47, 7, 10, 2, 2, 41, 43, 5, 10, 6, 2, 42, 44, 7, 3,
	2, 2, 43, 42, 3, 2, 2, 2, 43, 44, 3, 2, 2, 2, 44, 46, 3, 2, 2, 2, 45, 41,
	3, 2, 2, 2, 46, 49, 3, 2, 2, 2, 47, 45, 3, 2, 2, 2, 47, 48, 3, 2, 2, 2,
	48, 50, 3, 2, 2, 2, 49, 47, 3, 2, 2, 2, 50, 51, 7, 11, 2, 2, 51, 7, 3,
	2, 2, 2, 52, 53, 5, 18, 10, 2, 53, 60, 7, 12, 2, 2, 54, 56, 5, 14, 8, 2,
	55, 57, 7, 3, 2, 2, 56, 55, 3, 2, 2, 2, 56, 57, 3, 2, 2, 2, 57, 59, 3,
	2, 2, 2, 58, 54, 3, 2, 2, 2, 59, 62, 3, 2, 2, 2, 60, 58, 3, 2, 2, 2, 60,
	61, 3, 2, 2, 2, 61, 63, 3, 2, 2, 2, 62, 60, 3, 2, 2, 2, 63, 64, 7, 13,
	2, 2, 64, 9, 3, 2, 2, 2, 65, 66, 5, 12, 7, 2, 66, 67, 7, 4, 2, 2, 67, 68,
	5, 14, 8, 2, 68, 11, 3, 2, 2, 2, 69, 72, 5, 18, 10, 2, 70, 72, 7, 21, 2,
	2, 71, 69, 3, 2, 2, 2, 71, 70, 3, 2, 2, 2, 72, 13, 3, 2, 2, 2, 73, 78,
	5, 16, 9, 2, 74, 78, 5, 4, 3, 2, 75, 78, 5, 6, 4, 2, 76, 78, 5, 8, 5, 2,
	77, 73, 3, 2, 2, 2, 77, 74, 3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 77, 76, 3,
	2, 2, 2, 78, 15, 3, 2, 2, 2, 79, 80, 9, 2, 2, 2, 80, 17, 3, 2, 2, 2, 81,
	82, 9, 3, 2, 2, 82, 19, 3, 2, 2, 2, 11, 23, 31, 35, 43, 47, 56, 60, 71,
	77,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "','", "':'", "", "", "", "'['", "']'", "'{'", "'}'", "'('", "')'",
}
var symbolicNames = []string{
	"", "", "", "IntegerLiteral", "MultiLineComment", "SingleLineComment",
	"OpenBracket", "CloseBracket", "OpenBrace", "CloseBrace", "OpenParen",
	"CloseParen", "BooleanLiteral", "FloatLiteral", "SignedIntegerLiteral",
	"BinaryIntegerLiteral", "OctalIntegerLiteral", "DecimalIntegerLiteral",
	"HexIntegerLiteral", "StringLiteral", "RawStringLiteral", "WhiteSpaces",
	"LineTerminator", "Identifier",
}

var ruleNames = []string{
	"config", "arrayLiteral", "mapLiteral", "call", "propertyAssignment", "propertyName",
	"singleExpression", "literal", "identifierName",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type JaclParser struct {
	*antlr.BaseParser
}

func NewJaclParser(input antlr.TokenStream) *JaclParser {
	this := new(JaclParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Jacl.g4"

	return this
}

// JaclParser tokens.
const (
	JaclParserEOF                   = antlr.TokenEOF
	JaclParserT__0                  = 1
	JaclParserT__1                  = 2
	JaclParserIntegerLiteral        = 3
	JaclParserMultiLineComment      = 4
	JaclParserSingleLineComment     = 5
	JaclParserOpenBracket           = 6
	JaclParserCloseBracket          = 7
	JaclParserOpenBrace             = 8
	JaclParserCloseBrace            = 9
	JaclParserOpenParen             = 10
	JaclParserCloseParen            = 11
	JaclParserBooleanLiteral        = 12
	JaclParserFloatLiteral          = 13
	JaclParserSignedIntegerLiteral  = 14
	JaclParserBinaryIntegerLiteral  = 15
	JaclParserOctalIntegerLiteral   = 16
	JaclParserDecimalIntegerLiteral = 17
	JaclParserHexIntegerLiteral     = 18
	JaclParserStringLiteral         = 19
	JaclParserRawStringLiteral      = 20
	JaclParserWhiteSpaces           = 21
	JaclParserLineTerminator        = 22
	JaclParserIdentifier            = 23
)

// JaclParser rules.
const (
	JaclParserRULE_config             = 0
	JaclParserRULE_arrayLiteral       = 1
	JaclParserRULE_mapLiteral         = 2
	JaclParserRULE_call               = 3
	JaclParserRULE_propertyAssignment = 4
	JaclParserRULE_propertyName       = 5
	JaclParserRULE_singleExpression   = 6
	JaclParserRULE_literal            = 7
	JaclParserRULE_identifierName     = 8
)

// IConfigContext is an interface to support dynamic dispatch.
type IConfigContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConfigContext differentiates from other interfaces.
	IsConfigContext()
}

type ConfigContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConfigContext() *ConfigContext {
	var p = new(ConfigContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JaclParserRULE_config
	return p
}

func (*ConfigContext) IsConfigContext() {}

func NewConfigContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConfigContext {
	var p = new(ConfigContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JaclParserRULE_config

	return p
}

func (s *ConfigContext) GetParser() antlr.Parser { return s.parser }

func (s *ConfigContext) EOF() antlr.TerminalNode {
	return s.GetToken(JaclParserEOF, 0)
}

func (s *ConfigContext) AllPropertyAssignment() []IPropertyAssignmentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPropertyAssignmentContext)(nil)).Elem())
	var tst = make([]IPropertyAssignmentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPropertyAssignmentContext)
		}
	}

	return tst
}

func (s *ConfigContext) PropertyAssignment(i int) IPropertyAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyAssignmentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPropertyAssignmentContext)
}

func (s *ConfigContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConfigContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConfigContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JaclListener); ok {
		listenerT.EnterConfig(s)
	}
}

func (s *ConfigContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JaclListener); ok {
		listenerT.ExitConfig(s)
	}
}

func (p *JaclParser) Config() (localctx IConfigContext) {
	localctx = NewConfigContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, JaclParserRULE_config)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JaclParserBooleanLiteral)|(1<<JaclParserStringLiteral)|(1<<JaclParserIdentifier))) != 0 {
		{
			p.SetState(18)
			p.PropertyAssignment()
		}

		p.SetState(23)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(24)
		p.Match(JaclParserEOF)
	}

	return localctx
}

// IArrayLiteralContext is an interface to support dynamic dispatch.
type IArrayLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayLiteralContext differentiates from other interfaces.
	IsArrayLiteralContext()
}

type ArrayLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayLiteralContext() *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JaclParserRULE_arrayLiteral
	return p
}

func (*ArrayLiteralContext) IsArrayLiteralContext() {}

func NewArrayLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JaclParserRULE_arrayLiteral

	return p
}

func (s *ArrayLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayLiteralContext) OpenBracket() antlr.TerminalNode {
	return s.GetToken(JaclParserOpenBracket, 0)
}

func (s *ArrayLiteralContext) CloseBracket() antlr.TerminalNode {
	return s.GetToken(JaclParserCloseBracket, 0)
}

func (s *ArrayLiteralContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *ArrayLiteralContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *ArrayLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JaclListener); ok {
		listenerT.EnterArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JaclListener); ok {
		listenerT.ExitArrayLiteral(s)
	}
}

func (p *JaclParser) ArrayLiteral() (localctx IArrayLiteralContext) {
	localctx = NewArrayLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, JaclParserRULE_arrayLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(26)
		p.Match(JaclParserOpenBracket)
	}
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JaclParserIntegerLiteral)|(1<<JaclParserOpenBracket)|(1<<JaclParserOpenBrace)|(1<<JaclParserBooleanLiteral)|(1<<JaclParserFloatLiteral)|(1<<JaclParserStringLiteral)|(1<<JaclParserRawStringLiteral)|(1<<JaclParserIdentifier))) != 0 {
		{
			p.SetState(27)
			p.SingleExpression()
		}
		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JaclParserT__0 {
			{
				p.SetState(28)
				p.Match(JaclParserT__0)
			}

		}

		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(36)
		p.Match(JaclParserCloseBracket)
	}

	return localctx
}

// IMapLiteralContext is an interface to support dynamic dispatch.
type IMapLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMapLiteralContext differentiates from other interfaces.
	IsMapLiteralContext()
}

type MapLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapLiteralContext() *MapLiteralContext {
	var p = new(MapLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JaclParserRULE_mapLiteral
	return p
}

func (*MapLiteralContext) IsMapLiteralContext() {}

func NewMapLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapLiteralContext {
	var p = new(MapLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JaclParserRULE_mapLiteral

	return p
}

func (s *MapLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *MapLiteralContext) OpenBrace() antlr.TerminalNode {
	return s.GetToken(JaclParserOpenBrace, 0)
}

func (s *MapLiteralContext) CloseBrace() antlr.TerminalNode {
	return s.GetToken(JaclParserCloseBrace, 0)
}

func (s *MapLiteralContext) AllPropertyAssignment() []IPropertyAssignmentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPropertyAssignmentContext)(nil)).Elem())
	var tst = make([]IPropertyAssignmentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPropertyAssignmentContext)
		}
	}

	return tst
}

func (s *MapLiteralContext) PropertyAssignment(i int) IPropertyAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyAssignmentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPropertyAssignmentContext)
}

func (s *MapLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JaclListener); ok {
		listenerT.EnterMapLiteral(s)
	}
}

func (s *MapLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JaclListener); ok {
		listenerT.ExitMapLiteral(s)
	}
}

func (p *JaclParser) MapLiteral() (localctx IMapLiteralContext) {
	localctx = NewMapLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, JaclParserRULE_mapLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(38)
		p.Match(JaclParserOpenBrace)
	}
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JaclParserBooleanLiteral)|(1<<JaclParserStringLiteral)|(1<<JaclParserIdentifier))) != 0 {
		{
			p.SetState(39)
			p.PropertyAssignment()
		}
		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JaclParserT__0 {
			{
				p.SetState(40)
				p.Match(JaclParserT__0)
			}

		}

		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(48)
		p.Match(JaclParserCloseBrace)
	}

	return localctx
}

// ICallContext is an interface to support dynamic dispatch.
type ICallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCallContext differentiates from other interfaces.
	IsCallContext()
}

type CallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCallContext() *CallContext {
	var p = new(CallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JaclParserRULE_call
	return p
}

func (*CallContext) IsCallContext() {}

func NewCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallContext {
	var p = new(CallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JaclParserRULE_call

	return p
}

func (s *CallContext) GetParser() antlr.Parser { return s.parser }

func (s *CallContext) IdentifierName() IIdentifierNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierNameContext)
}

func (s *CallContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(JaclParserOpenParen, 0)
}

func (s *CallContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(JaclParserCloseParen, 0)
}

func (s *CallContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *CallContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *CallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JaclListener); ok {
		listenerT.EnterCall(s)
	}
}

func (s *CallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JaclListener); ok {
		listenerT.ExitCall(s)
	}
}

func (p *JaclParser) Call() (localctx ICallContext) {
	localctx = NewCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, JaclParserRULE_call)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		p.IdentifierName()
	}
	{
		p.SetState(51)
		p.Match(JaclParserOpenParen)
	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JaclParserIntegerLiteral)|(1<<JaclParserOpenBracket)|(1<<JaclParserOpenBrace)|(1<<JaclParserBooleanLiteral)|(1<<JaclParserFloatLiteral)|(1<<JaclParserStringLiteral)|(1<<JaclParserRawStringLiteral)|(1<<JaclParserIdentifier))) != 0 {
		{
			p.SetState(52)
			p.SingleExpression()
		}
		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JaclParserT__0 {
			{
				p.SetState(53)
				p.Match(JaclParserT__0)
			}

		}

		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(61)
		p.Match(JaclParserCloseParen)
	}

	return localctx
}

// IPropertyAssignmentContext is an interface to support dynamic dispatch.
type IPropertyAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropertyAssignmentContext differentiates from other interfaces.
	IsPropertyAssignmentContext()
}

type PropertyAssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyAssignmentContext() *PropertyAssignmentContext {
	var p = new(PropertyAssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JaclParserRULE_propertyAssignment
	return p
}

func (*PropertyAssignmentContext) IsPropertyAssignmentContext() {}

func NewPropertyAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyAssignmentContext {
	var p = new(PropertyAssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JaclParserRULE_propertyAssignment

	return p
}

func (s *PropertyAssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyAssignmentContext) PropertyName() IPropertyNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyNameContext)
}

func (s *PropertyAssignmentContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *PropertyAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyAssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JaclListener); ok {
		listenerT.EnterPropertyAssignment(s)
	}
}

func (s *PropertyAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JaclListener); ok {
		listenerT.ExitPropertyAssignment(s)
	}
}

func (p *JaclParser) PropertyAssignment() (localctx IPropertyAssignmentContext) {
	localctx = NewPropertyAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, JaclParserRULE_propertyAssignment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(63)
		p.PropertyName()
	}
	{
		p.SetState(64)
		p.Match(JaclParserT__1)
	}
	{
		p.SetState(65)
		p.SingleExpression()
	}

	return localctx
}

// IPropertyNameContext is an interface to support dynamic dispatch.
type IPropertyNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropertyNameContext differentiates from other interfaces.
	IsPropertyNameContext()
}

type PropertyNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyNameContext() *PropertyNameContext {
	var p = new(PropertyNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JaclParserRULE_propertyName
	return p
}

func (*PropertyNameContext) IsPropertyNameContext() {}

func NewPropertyNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyNameContext {
	var p = new(PropertyNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JaclParserRULE_propertyName

	return p
}

func (s *PropertyNameContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyNameContext) IdentifierName() IIdentifierNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierNameContext)
}

func (s *PropertyNameContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(JaclParserStringLiteral, 0)
}

func (s *PropertyNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JaclListener); ok {
		listenerT.EnterPropertyName(s)
	}
}

func (s *PropertyNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JaclListener); ok {
		listenerT.ExitPropertyName(s)
	}
}

func (p *JaclParser) PropertyName() (localctx IPropertyNameContext) {
	localctx = NewPropertyNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, JaclParserRULE_propertyName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(69)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JaclParserBooleanLiteral, JaclParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(67)
			p.IdentifierName()
		}

	case JaclParserStringLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(68)
			p.Match(JaclParserStringLiteral)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISingleExpressionContext is an interface to support dynamic dispatch.
type ISingleExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSingleExpressionContext differentiates from other interfaces.
	IsSingleExpressionContext()
}

type SingleExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleExpressionContext() *SingleExpressionContext {
	var p = new(SingleExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JaclParserRULE_singleExpression
	return p
}

func (*SingleExpressionContext) IsSingleExpressionContext() {}

func NewSingleExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleExpressionContext {
	var p = new(SingleExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JaclParserRULE_singleExpression

	return p
}

func (s *SingleExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleExpressionContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *SingleExpressionContext) ArrayLiteral() IArrayLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayLiteralContext)
}

func (s *SingleExpressionContext) MapLiteral() IMapLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMapLiteralContext)
}

func (s *SingleExpressionContext) Call() ICallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICallContext)
}

func (s *SingleExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SingleExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JaclListener); ok {
		listenerT.EnterSingleExpression(s)
	}
}

func (s *SingleExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JaclListener); ok {
		listenerT.ExitSingleExpression(s)
	}
}

func (p *JaclParser) SingleExpression() (localctx ISingleExpressionContext) {
	localctx = NewSingleExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, JaclParserRULE_singleExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(71)
			p.Literal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(72)
			p.ArrayLiteral()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(73)
			p.MapLiteral()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(74)
			p.Call()
		}

	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JaclParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JaclParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) BooleanLiteral() antlr.TerminalNode {
	return s.GetToken(JaclParserBooleanLiteral, 0)
}

func (s *LiteralContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(JaclParserStringLiteral, 0)
}

func (s *LiteralContext) RawStringLiteral() antlr.TerminalNode {
	return s.GetToken(JaclParserRawStringLiteral, 0)
}

func (s *LiteralContext) FloatLiteral() antlr.TerminalNode {
	return s.GetToken(JaclParserFloatLiteral, 0)
}

func (s *LiteralContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(JaclParserIntegerLiteral, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JaclListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JaclListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (p *JaclParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, JaclParserRULE_literal)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(77)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JaclParserIntegerLiteral)|(1<<JaclParserBooleanLiteral)|(1<<JaclParserFloatLiteral)|(1<<JaclParserStringLiteral)|(1<<JaclParserRawStringLiteral))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IIdentifierNameContext is an interface to support dynamic dispatch.
type IIdentifierNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierNameContext differentiates from other interfaces.
	IsIdentifierNameContext()
}

type IdentifierNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierNameContext() *IdentifierNameContext {
	var p = new(IdentifierNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JaclParserRULE_identifierName
	return p
}

func (*IdentifierNameContext) IsIdentifierNameContext() {}

func NewIdentifierNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierNameContext {
	var p = new(IdentifierNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JaclParserRULE_identifierName

	return p
}

func (s *IdentifierNameContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierNameContext) Identifier() antlr.TerminalNode {
	return s.GetToken(JaclParserIdentifier, 0)
}

func (s *IdentifierNameContext) BooleanLiteral() antlr.TerminalNode {
	return s.GetToken(JaclParserBooleanLiteral, 0)
}

func (s *IdentifierNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JaclListener); ok {
		listenerT.EnterIdentifierName(s)
	}
}

func (s *IdentifierNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JaclListener); ok {
		listenerT.ExitIdentifierName(s)
	}
}

func (p *JaclParser) IdentifierName() (localctx IIdentifierNameContext) {
	localctx = NewIdentifierNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, JaclParserRULE_identifierName)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(79)
	_la = p.GetTokenStream().LA(1)

	if !(_la == JaclParserBooleanLiteral || _la == JaclParserIdentifier) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}
