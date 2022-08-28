package main

import (
	"fmt"

	"github.com/abu-lang/abuc/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type goabuProgram struct {
	Id        string
	Invariant *string
	Resources []string
	Rules     []string
	Tick      int
}

// typeToGoabu takes an AbU type and returns it in GoAbU's syntax.
func typeToGoabu(typ string) string {
	switch typ {
	case "boolean":
		return "Bool"
	case "integer":
		return "Integer"
	case "decimal":
		return "Float"
	case "string":
		return "Text"
	default:
		return "Other"
	}
}

// makeGoabuProgram uses an abuProgram and the CommonTokenStream used for its parsing to
// create an equivalent goabuProgram.
//
// Precondition: r != nil for _, r := range prog.rules
func makeGoabuProgram(prog abuProgram, stream *antlr.CommonTokenStream) goabuProgram {
	res := goabuProgram{
		Id:   prog.id,
		Tick: prog.tick,
	}
	list := newAbuRewriter(stream)
	if prog.invariant != nil {
		antlr.ParseTreeWalkerDefault.Walk(list, *prog.invariant)
		inv := list.getAlteredText(*prog.invariant)
		res.Invariant = &inv
	}
	for _, r := range prog.resources {
		resource := fmt.Sprintf(".%s[\"%s\"] = ", typeToGoabu(r.typ), r.id)
		if r.initExpr != nil {
			antlr.ParseTreeWalkerDefault.Walk(list, *r.initExpr)
			resource += list.getAlteredText(*r.initExpr)
		} else {
			switch typeToGoabu(r.typ) {
			case "Bool":
				resource += "true"
			case "Integer":
				fallthrough
			case "Float":
				resource += "0"
			case "Text":
				resource += "\"\""
			case "Time":
				resource += "time.Time{}"
			case "Other":
				resource += "struct{}{}"
			}
		}
		res.Resources = append(res.Resources, resource)
	}
	for _, r := range prog.rules {
		antlr.ParseTreeWalkerDefault.Walk(list, *r)
		res.Rules = append(res.Rules, fmt.Sprintf("`%s`", list.getAlteredText(*r)))
	}
	return res
}

type abuRewriter struct {
	parser.BaseAbuParserListener
	rewriter *antlr.TokenStreamRewriter
}

func newAbuRewriter(stream *antlr.CommonTokenStream) *abuRewriter {
	return &abuRewriter{rewriter: antlr.NewTokenStreamRewriter(stream)}
}

// ExitExpression is called when production expression is exited.
func (t *abuRewriter) ExitExpression(ctx *parser.ExpressionContext) {
	switch true {
	case ctx.AND() != nil:
		t.rewriter.ReplaceTokenDefaultPos(ctx.AND().GetSymbol(), "&&")
	case ctx.OR() != nil:
		t.rewriter.ReplaceTokenDefaultPos(ctx.OR().GetSymbol(), "||")
	case ctx.NOT() != nil:
		t.rewriter.ReplaceTokenDefaultPos(ctx.NOT().GetSymbol(), "!")
	case ctx.DOUBLECOLON() != nil:
		t.rewriter.ReplaceTokenDefaultPos(ctx.DOUBLECOLON().GetSymbol(), "+")
	case ctx.ABS() != nil: // Note that Abs() always returns float64
		t.rewriter.ReplaceTokenDefaultPos(ctx.ABS().GetSymbol(), "Abs(")
		t.rewriter.InsertAfterToken("default", ctx.Expression(0).GetStop(), ")")
	}
}

// getAlteredText returns the text of the ParserRuleContext possibly altered
// through the TokenStreamRewriter default program.
func (t *abuRewriter) getAlteredText(ctx antlr.ParserRuleContext) string {
	return t.rewriter.GetText("default", &antlr.Interval{
		Start: ctx.GetStart().GetTokenIndex(),
		Stop:  ctx.GetStop().GetTokenIndex(),
	})
}
