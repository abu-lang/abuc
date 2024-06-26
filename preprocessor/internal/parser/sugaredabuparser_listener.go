// Code generated from SugaredAbuParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // SugaredAbuParser
import "github.com/antlr4-go/antlr/v4"

// SugaredAbuParserListener is a complete listener for a parse tree produced by SugaredAbuParser.
type SugaredAbuParserListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterTypeDecl is called when entering the typeDecl production.
	EnterTypeDecl(c *TypeDeclContext)

	// EnterResField is called when entering the resField production.
	EnterResField(c *ResFieldContext)

	// EnterDevice is called when entering the device production.
	EnterDevice(c *DeviceContext)

	// EnterResList is called when entering the resList production.
	EnterResList(c *ResListContext)

	// EnterResDecl is called when entering the resDecl production.
	EnterResDecl(c *ResDeclContext)

	// EnterCompResDecl is called when entering the compResDecl production.
	EnterCompResDecl(c *CompResDeclContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterEcarule is called when entering the ecarule production.
	EnterEcarule(c *EcaruleContext)

	// EnterLetDecl is called when entering the letDecl production.
	EnterLetDecl(c *LetDeclContext)

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

	// EnterSimpleResource is called when entering the simpleResource production.
	EnterSimpleResource(c *SimpleResourceContext)

	// EnterNestedResource is called when entering the nestedResource production.
	EnterNestedResource(c *NestedResourceContext)

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

	// ExitTypeDecl is called when exiting the typeDecl production.
	ExitTypeDecl(c *TypeDeclContext)

	// ExitResField is called when exiting the resField production.
	ExitResField(c *ResFieldContext)

	// ExitDevice is called when exiting the device production.
	ExitDevice(c *DeviceContext)

	// ExitResList is called when exiting the resList production.
	ExitResList(c *ResListContext)

	// ExitResDecl is called when exiting the resDecl production.
	ExitResDecl(c *ResDeclContext)

	// ExitCompResDecl is called when exiting the compResDecl production.
	ExitCompResDecl(c *CompResDeclContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitEcarule is called when exiting the ecarule production.
	ExitEcarule(c *EcaruleContext)

	// ExitLetDecl is called when exiting the letDecl production.
	ExitLetDecl(c *LetDeclContext)

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

	// ExitSimpleResource is called when exiting the simpleResource production.
	ExitSimpleResource(c *SimpleResourceContext)

	// ExitNestedResource is called when exiting the nestedResource production.
	ExitNestedResource(c *NestedResourceContext)

	// ExitDecimalLiteral is called when exiting the decimalLiteral production.
	ExitDecimalLiteral(c *DecimalLiteralContext)

	// ExitIntegerLiteral is called when exiting the integerLiteral production.
	ExitIntegerLiteral(c *IntegerLiteralContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitBooleanLiteral is called when exiting the booleanLiteral production.
	ExitBooleanLiteral(c *BooleanLiteralContext)
}
