// Code generated from AbuParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // AbuParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAbuParserListener is a complete listener for a parse tree produced by AbuParser.
type BaseAbuParserListener struct{}

var _ AbuParserListener = &BaseAbuParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAbuParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAbuParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAbuParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAbuParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseAbuParserListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseAbuParserListener) ExitProgram(ctx *ProgramContext) {}

// EnterTypeDecl is called when production typeDecl is entered.
func (s *BaseAbuParserListener) EnterTypeDecl(ctx *TypeDeclContext) {}

// ExitTypeDecl is called when production typeDecl is exited.
func (s *BaseAbuParserListener) ExitTypeDecl(ctx *TypeDeclContext) {}

// EnterResField is called when production resField is entered.
func (s *BaseAbuParserListener) EnterResField(ctx *ResFieldContext) {}

// ExitResField is called when production resField is exited.
func (s *BaseAbuParserListener) ExitResField(ctx *ResFieldContext) {}

// EnterDevice is called when production device is entered.
func (s *BaseAbuParserListener) EnterDevice(ctx *DeviceContext) {}

// ExitDevice is called when production device is exited.
func (s *BaseAbuParserListener) ExitDevice(ctx *DeviceContext) {}

// EnterResList is called when production resList is entered.
func (s *BaseAbuParserListener) EnterResList(ctx *ResListContext) {}

// ExitResList is called when production resList is exited.
func (s *BaseAbuParserListener) ExitResList(ctx *ResListContext) {}

// EnterResDecl is called when production resDecl is entered.
func (s *BaseAbuParserListener) EnterResDecl(ctx *ResDeclContext) {}

// ExitResDecl is called when production resDecl is exited.
func (s *BaseAbuParserListener) ExitResDecl(ctx *ResDeclContext) {}

// EnterCompResDecl is called when production compResDecl is entered.
func (s *BaseAbuParserListener) EnterCompResDecl(ctx *CompResDeclContext) {}

// ExitCompResDecl is called when production compResDecl is exited.
func (s *BaseAbuParserListener) ExitCompResDecl(ctx *CompResDeclContext) {}

// EnterType is called when production type is entered.
func (s *BaseAbuParserListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseAbuParserListener) ExitType(ctx *TypeContext) {}

// EnterEcarule is called when production ecarule is entered.
func (s *BaseAbuParserListener) EnterEcarule(ctx *EcaruleContext) {}

// ExitEcarule is called when production ecarule is exited.
func (s *BaseAbuParserListener) ExitEcarule(ctx *EcaruleContext) {}

// EnterTask is called when production task is entered.
func (s *BaseAbuParserListener) EnterTask(ctx *TaskContext) {}

// ExitTask is called when production task is exited.
func (s *BaseAbuParserListener) ExitTask(ctx *TaskContext) {}

// EnterActions is called when production actions is entered.
func (s *BaseAbuParserListener) EnterActions(ctx *ActionsContext) {}

// ExitActions is called when production actions is exited.
func (s *BaseAbuParserListener) ExitActions(ctx *ActionsContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseAbuParserListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseAbuParserListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseAbuParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseAbuParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterMulDivModOperator is called when production mulDivModOperator is entered.
func (s *BaseAbuParserListener) EnterMulDivModOperator(ctx *MulDivModOperatorContext) {}

// ExitMulDivModOperator is called when production mulDivModOperator is exited.
func (s *BaseAbuParserListener) ExitMulDivModOperator(ctx *MulDivModOperatorContext) {}

// EnterPlusMinusOperator is called when production plusMinusOperator is entered.
func (s *BaseAbuParserListener) EnterPlusMinusOperator(ctx *PlusMinusOperatorContext) {}

// ExitPlusMinusOperator is called when production plusMinusOperator is exited.
func (s *BaseAbuParserListener) ExitPlusMinusOperator(ctx *PlusMinusOperatorContext) {}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *BaseAbuParserListener) EnterComparisonOperator(ctx *ComparisonOperatorContext) {}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *BaseAbuParserListener) ExitComparisonOperator(ctx *ComparisonOperatorContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseAbuParserListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseAbuParserListener) ExitTerm(ctx *TermContext) {}

// EnterValue is called when production value is entered.
func (s *BaseAbuParserListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseAbuParserListener) ExitValue(ctx *ValueContext) {}

// EnterResource is called when production resource is entered.
func (s *BaseAbuParserListener) EnterResource(ctx *ResourceContext) {}

// ExitResource is called when production resource is exited.
func (s *BaseAbuParserListener) ExitResource(ctx *ResourceContext) {}

// EnterSimpleResource is called when production simpleResource is entered.
func (s *BaseAbuParserListener) EnterSimpleResource(ctx *SimpleResourceContext) {}

// ExitSimpleResource is called when production simpleResource is exited.
func (s *BaseAbuParserListener) ExitSimpleResource(ctx *SimpleResourceContext) {}

// EnterNestedResource is called when production nestedResource is entered.
func (s *BaseAbuParserListener) EnterNestedResource(ctx *NestedResourceContext) {}

// ExitNestedResource is called when production nestedResource is exited.
func (s *BaseAbuParserListener) ExitNestedResource(ctx *NestedResourceContext) {}

// EnterDecimalLiteral is called when production decimalLiteral is entered.
func (s *BaseAbuParserListener) EnterDecimalLiteral(ctx *DecimalLiteralContext) {}

// ExitDecimalLiteral is called when production decimalLiteral is exited.
func (s *BaseAbuParserListener) ExitDecimalLiteral(ctx *DecimalLiteralContext) {}

// EnterIntegerLiteral is called when production integerLiteral is entered.
func (s *BaseAbuParserListener) EnterIntegerLiteral(ctx *IntegerLiteralContext) {}

// ExitIntegerLiteral is called when production integerLiteral is exited.
func (s *BaseAbuParserListener) ExitIntegerLiteral(ctx *IntegerLiteralContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseAbuParserListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseAbuParserListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BaseAbuParserListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BaseAbuParserListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}
