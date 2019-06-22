// Generated from parser/Jacl.g4 by ANTLR 4.7.

package parser // Jacl

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseJaclListener is a complete listener for a parse tree produced by JaclParser.
type BaseJaclListener struct{}

var _ JaclListener = &BaseJaclListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseJaclListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseJaclListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseJaclListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseJaclListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterConfig is called when production config is entered.
func (s *BaseJaclListener) EnterConfig(ctx *ConfigContext) {}

// ExitConfig is called when production config is exited.
func (s *BaseJaclListener) ExitConfig(ctx *ConfigContext) {}

// EnterArrayLiteral is called when production arrayLiteral is entered.
func (s *BaseJaclListener) EnterArrayLiteral(ctx *ArrayLiteralContext) {}

// ExitArrayLiteral is called when production arrayLiteral is exited.
func (s *BaseJaclListener) ExitArrayLiteral(ctx *ArrayLiteralContext) {}

// EnterMapLiteral is called when production mapLiteral is entered.
func (s *BaseJaclListener) EnterMapLiteral(ctx *MapLiteralContext) {}

// ExitMapLiteral is called when production mapLiteral is exited.
func (s *BaseJaclListener) ExitMapLiteral(ctx *MapLiteralContext) {}

// EnterCall is called when production call is entered.
func (s *BaseJaclListener) EnterCall(ctx *CallContext) {}

// ExitCall is called when production call is exited.
func (s *BaseJaclListener) ExitCall(ctx *CallContext) {}

// EnterPropertyAssignment is called when production propertyAssignment is entered.
func (s *BaseJaclListener) EnterPropertyAssignment(ctx *PropertyAssignmentContext) {}

// ExitPropertyAssignment is called when production propertyAssignment is exited.
func (s *BaseJaclListener) ExitPropertyAssignment(ctx *PropertyAssignmentContext) {}

// EnterPropertyName is called when production propertyName is entered.
func (s *BaseJaclListener) EnterPropertyName(ctx *PropertyNameContext) {}

// ExitPropertyName is called when production propertyName is exited.
func (s *BaseJaclListener) ExitPropertyName(ctx *PropertyNameContext) {}

// EnterSingleExpression is called when production singleExpression is entered.
func (s *BaseJaclListener) EnterSingleExpression(ctx *SingleExpressionContext) {}

// ExitSingleExpression is called when production singleExpression is exited.
func (s *BaseJaclListener) ExitSingleExpression(ctx *SingleExpressionContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseJaclListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseJaclListener) ExitLiteral(ctx *LiteralContext) {}

// EnterIdentifierName is called when production identifierName is entered.
func (s *BaseJaclListener) EnterIdentifierName(ctx *IdentifierNameContext) {}

// ExitIdentifierName is called when production identifierName is exited.
func (s *BaseJaclListener) ExitIdentifierName(ctx *IdentifierNameContext) {}
