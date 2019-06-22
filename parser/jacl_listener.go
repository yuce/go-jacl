// Generated from parser/Jacl.g4 by ANTLR 4.7.

package parser // Jacl

import "github.com/antlr/antlr4/runtime/Go/antlr"

// JaclListener is a complete listener for a parse tree produced by JaclParser.
type JaclListener interface {
	antlr.ParseTreeListener

	// EnterConfig is called when entering the config production.
	EnterConfig(c *ConfigContext)

	// EnterArrayLiteral is called when entering the arrayLiteral production.
	EnterArrayLiteral(c *ArrayLiteralContext)

	// EnterMapLiteral is called when entering the mapLiteral production.
	EnterMapLiteral(c *MapLiteralContext)

	// EnterCall is called when entering the call production.
	EnterCall(c *CallContext)

	// EnterPropertyAssignment is called when entering the propertyAssignment production.
	EnterPropertyAssignment(c *PropertyAssignmentContext)

	// EnterPropertyName is called when entering the propertyName production.
	EnterPropertyName(c *PropertyNameContext)

	// EnterSingleExpression is called when entering the singleExpression production.
	EnterSingleExpression(c *SingleExpressionContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterIdentifierName is called when entering the identifierName production.
	EnterIdentifierName(c *IdentifierNameContext)

	// ExitConfig is called when exiting the config production.
	ExitConfig(c *ConfigContext)

	// ExitArrayLiteral is called when exiting the arrayLiteral production.
	ExitArrayLiteral(c *ArrayLiteralContext)

	// ExitMapLiteral is called when exiting the mapLiteral production.
	ExitMapLiteral(c *MapLiteralContext)

	// ExitCall is called when exiting the call production.
	ExitCall(c *CallContext)

	// ExitPropertyAssignment is called when exiting the propertyAssignment production.
	ExitPropertyAssignment(c *PropertyAssignmentContext)

	// ExitPropertyName is called when exiting the propertyName production.
	ExitPropertyName(c *PropertyNameContext)

	// ExitSingleExpression is called when exiting the singleExpression production.
	ExitSingleExpression(c *SingleExpressionContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitIdentifierName is called when exiting the identifierName production.
	ExitIdentifierName(c *IdentifierNameContext)
}
