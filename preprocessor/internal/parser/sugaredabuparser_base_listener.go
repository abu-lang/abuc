// Code generated from SugaredAbuParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // SugaredAbuParser
import "github.com/antlr4-go/antlr/v4"

// BaseSugaredAbuParserListener is a complete listener for a parse tree produced by SugaredAbuParser.
type BaseSugaredAbuParserListener struct{}

var _ SugaredAbuParserListener = &BaseSugaredAbuParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSugaredAbuParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSugaredAbuParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSugaredAbuParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSugaredAbuParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseSugaredAbuParserListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseSugaredAbuParserListener) ExitProgram(ctx *ProgramContext) {}

// EnterTypeDecl is called when production typeDecl is entered.
func (s *BaseSugaredAbuParserListener) EnterTypeDecl(ctx *TypeDeclContext) {}

// ExitTypeDecl is called when production typeDecl is exited.
func (s *BaseSugaredAbuParserListener) ExitTypeDecl(ctx *TypeDeclContext) {}

// EnterResField is called when production resField is entered.
func (s *BaseSugaredAbuParserListener) EnterResField(ctx *ResFieldContext) {}

// ExitResField is called when production resField is exited.
func (s *BaseSugaredAbuParserListener) ExitResField(ctx *ResFieldContext) {}

// EnterDevice is called when production device is entered.
func (s *BaseSugaredAbuParserListener) EnterDevice(ctx *DeviceContext) {}

// ExitDevice is called when production device is exited.
func (s *BaseSugaredAbuParserListener) ExitDevice(ctx *DeviceContext) {}

// EnterResList is called when production resList is entered.
func (s *BaseSugaredAbuParserListener) EnterResList(ctx *ResListContext) {}

// ExitResList is called when production resList is exited.
func (s *BaseSugaredAbuParserListener) ExitResList(ctx *ResListContext) {}

// EnterResDecl is called when production resDecl is entered.
func (s *BaseSugaredAbuParserListener) EnterResDecl(ctx *ResDeclContext) {}

// ExitResDecl is called when production resDecl is exited.
func (s *BaseSugaredAbuParserListener) ExitResDecl(ctx *ResDeclContext) {}

// EnterCompResDecl is called when production compResDecl is entered.
func (s *BaseSugaredAbuParserListener) EnterCompResDecl(ctx *CompResDeclContext) {}

// ExitCompResDecl is called when production compResDecl is exited.
func (s *BaseSugaredAbuParserListener) ExitCompResDecl(ctx *CompResDeclContext) {}

// EnterType is called when production type is entered.
func (s *BaseSugaredAbuParserListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseSugaredAbuParserListener) ExitType(ctx *TypeContext) {}

// EnterEcarule is called when production ecarule is entered.
func (s *BaseSugaredAbuParserListener) EnterEcarule(ctx *EcaruleContext) {}

// ExitEcarule is called when production ecarule is exited.
func (s *BaseSugaredAbuParserListener) ExitEcarule(ctx *EcaruleContext) {}

// EnterLetDecl is called when production letDecl is entered.
func (s *BaseSugaredAbuParserListener) EnterLetDecl(ctx *LetDeclContext) {}

// ExitLetDecl is called when production letDecl is exited.
func (s *BaseSugaredAbuParserListener) ExitLetDecl(ctx *LetDeclContext) {}

// EnterTask is called when production task is entered.
func (s *BaseSugaredAbuParserListener) EnterTask(ctx *TaskContext) {}

// ExitTask is called when production task is exited.
func (s *BaseSugaredAbuParserListener) ExitTask(ctx *TaskContext) {}

// EnterActions is called when production actions is entered.
func (s *BaseSugaredAbuParserListener) EnterActions(ctx *ActionsContext) {}

// ExitActions is called when production actions is exited.
func (s *BaseSugaredAbuParserListener) ExitActions(ctx *ActionsContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseSugaredAbuParserListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseSugaredAbuParserListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSugaredAbuParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSugaredAbuParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterMulDivModOperator is called when production mulDivModOperator is entered.
func (s *BaseSugaredAbuParserListener) EnterMulDivModOperator(ctx *MulDivModOperatorContext) {}

// ExitMulDivModOperator is called when production mulDivModOperator is exited.
func (s *BaseSugaredAbuParserListener) ExitMulDivModOperator(ctx *MulDivModOperatorContext) {}

// EnterPlusMinusOperator is called when production plusMinusOperator is entered.
func (s *BaseSugaredAbuParserListener) EnterPlusMinusOperator(ctx *PlusMinusOperatorContext) {}

// ExitPlusMinusOperator is called when production plusMinusOperator is exited.
func (s *BaseSugaredAbuParserListener) ExitPlusMinusOperator(ctx *PlusMinusOperatorContext) {}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *BaseSugaredAbuParserListener) EnterComparisonOperator(ctx *ComparisonOperatorContext) {}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *BaseSugaredAbuParserListener) ExitComparisonOperator(ctx *ComparisonOperatorContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseSugaredAbuParserListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseSugaredAbuParserListener) ExitTerm(ctx *TermContext) {}

// EnterValue is called when production value is entered.
func (s *BaseSugaredAbuParserListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseSugaredAbuParserListener) ExitValue(ctx *ValueContext) {}

// EnterResource is called when production resource is entered.
func (s *BaseSugaredAbuParserListener) EnterResource(ctx *ResourceContext) {}

// ExitResource is called when production resource is exited.
func (s *BaseSugaredAbuParserListener) ExitResource(ctx *ResourceContext) {}

// EnterSimpleResource is called when production simpleResource is entered.
func (s *BaseSugaredAbuParserListener) EnterSimpleResource(ctx *SimpleResourceContext) {}

// ExitSimpleResource is called when production simpleResource is exited.
func (s *BaseSugaredAbuParserListener) ExitSimpleResource(ctx *SimpleResourceContext) {}

// EnterNestedResource is called when production nestedResource is entered.
func (s *BaseSugaredAbuParserListener) EnterNestedResource(ctx *NestedResourceContext) {}

// ExitNestedResource is called when production nestedResource is exited.
func (s *BaseSugaredAbuParserListener) ExitNestedResource(ctx *NestedResourceContext) {}

// EnterDecimalLiteral is called when production decimalLiteral is entered.
func (s *BaseSugaredAbuParserListener) EnterDecimalLiteral(ctx *DecimalLiteralContext) {}

// ExitDecimalLiteral is called when production decimalLiteral is exited.
func (s *BaseSugaredAbuParserListener) ExitDecimalLiteral(ctx *DecimalLiteralContext) {}

// EnterIntegerLiteral is called when production integerLiteral is entered.
func (s *BaseSugaredAbuParserListener) EnterIntegerLiteral(ctx *IntegerLiteralContext) {}

// ExitIntegerLiteral is called when production integerLiteral is exited.
func (s *BaseSugaredAbuParserListener) ExitIntegerLiteral(ctx *IntegerLiteralContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseSugaredAbuParserListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseSugaredAbuParserListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BaseSugaredAbuParserListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BaseSugaredAbuParserListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}
