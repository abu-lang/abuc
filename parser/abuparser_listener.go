// Code generated from AbuParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // AbuParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// AbuParserListener is a complete listener for a parse tree produced by AbuParser.
type AbuParserListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterDevice is called when entering the device production.
	EnterDevice(c *DeviceContext)

	// EnterResList is called when entering the resList production.
	EnterResList(c *ResListContext)

	// EnterResDecl is called when entering the resDecl production.
	EnterResDecl(c *ResDeclContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterEcarule is called when entering the ecarule production.
	EnterEcarule(c *EcaruleContext)

	// EnterTask is called when entering the task production.
	EnterTask(c *TaskContext)

	// EnterActions is called when entering the actions production.
	EnterActions(c *ActionsContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMulDivModOperator is called when entering the mulDivModOperator production.
	EnterMulDivModOperator(c *MulDivModOperatorContext)

	// EnterPlusMinusOperator is called when entering the plusMinusOperator production.
	EnterPlusMinusOperator(c *PlusMinusOperatorContext)

	// EnterComparisonOperator is called when entering the comparisonOperator production.
	EnterComparisonOperator(c *ComparisonOperatorContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterResource is called when entering the resource production.
	EnterResource(c *ResourceContext)

	// EnterDecimalLiteral is called when entering the decimalLiteral production.
	EnterDecimalLiteral(c *DecimalLiteralContext)

	// EnterIntegerLiteral is called when entering the integerLiteral production.
	EnterIntegerLiteral(c *IntegerLiteralContext)

	// EnterStringLiteral is called when entering the stringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterBooleanLiteral is called when entering the booleanLiteral production.
	EnterBooleanLiteral(c *BooleanLiteralContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitDevice is called when exiting the device production.
	ExitDevice(c *DeviceContext)

	// ExitResList is called when exiting the resList production.
	ExitResList(c *ResListContext)

	// ExitResDecl is called when exiting the resDecl production.
	ExitResDecl(c *ResDeclContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitEcarule is called when exiting the ecarule production.
	ExitEcarule(c *EcaruleContext)

	// ExitTask is called when exiting the task production.
	ExitTask(c *TaskContext)

	// ExitActions is called when exiting the actions production.
	ExitActions(c *ActionsContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMulDivModOperator is called when exiting the mulDivModOperator production.
	ExitMulDivModOperator(c *MulDivModOperatorContext)

	// ExitPlusMinusOperator is called when exiting the plusMinusOperator production.
	ExitPlusMinusOperator(c *PlusMinusOperatorContext)

	// ExitComparisonOperator is called when exiting the comparisonOperator production.
	ExitComparisonOperator(c *ComparisonOperatorContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitResource is called when exiting the resource production.
	ExitResource(c *ResourceContext)

	// ExitDecimalLiteral is called when exiting the decimalLiteral production.
	ExitDecimalLiteral(c *DecimalLiteralContext)

	// ExitIntegerLiteral is called when exiting the integerLiteral production.
	ExitIntegerLiteral(c *IntegerLiteralContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitBooleanLiteral is called when exiting the booleanLiteral production.
	ExitBooleanLiteral(c *BooleanLiteralContext)
}
