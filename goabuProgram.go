// Copyright 2022 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"

	"github.com/abu-lang/abuc/internal/parser"
	"github.com/abu-lang/abuc/preprocessor"
	"github.com/antlr4-go/antlr/v4"
)

type goabuProgram struct {
	Id               string
	Invariant        *string
	Resources        []string
	Rules            []string
	StateInitializer string
	Imports          []string
	Tick             int
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
func makeGoabuProgram(prog abuProgram, cfg goabuConfig, stream *antlr.CommonTokenStream) (goabuProgram, []error) {
	res := goabuProgram{
		Id:   prog.id,
		Tick: prog.tick,
	}
	errList := errorHolder{}
	list := newAbuRewriter(stream, cfg, prog.DeviceSymbolTable, errList.addError)
	for id, nested := range prog.composedResources {
		res.Resources = append(res.Resources, list.translateComposedResource(id, nested))
	}
	if prog.invariant != nil {
		inv := list.translate(*prog.invariant)
		res.Invariant = &inv
	}
	for id, simple := range prog.simpleResources {
		res.Resources = append(res.Resources, list.translateSimpleResource(id, simple))
	}
	for _, rule := range prog.rules {
		res.Rules = append(res.Rules, list.translateRule(*rule))
	}
	if len(errList.errors) > 0 {
		return goabuProgram{}, errList.errors
	}
	if cfg.StateInitializer != "" {
		res.StateInitializer = cfg.StateInitializer
		res.Imports = cfg.Imports
	} else {
		res.StateInitializer = "memory.MakeResources()"
		res.Imports = append(cfg.Imports, "github.com/abu-lang/goabu/memory")
	}
	return res, nil
}

type abuRewriter struct {
	parser.BaseAbuParserListener
	preprocessor.DeviceSymbolTable
	goabuConfig

	inTask   bool
	errCb    func(error)
	rewriter *antlr.TokenStreamRewriter
}

func newAbuRewriter(stream *antlr.CommonTokenStream, cfg goabuConfig, st preprocessor.DeviceSymbolTable, errCb func(error)) *abuRewriter {
	return &abuRewriter{
		rewriter:          antlr.NewTokenStreamRewriter(stream),
		DeviceSymbolTable: st,
		goabuConfig:       cfg,
		errCb:             errCb,
	}
}

// ExitSimpleResource is called when production simpleResource is exited.
func (t *abuRewriter) ExitSimpleResource(ctx *parser.SimpleResourceContext) {
	if ctx.THIS() != nil {
		t.rewriter.DeleteTokenDefault(ctx.THIS().GetSymbol(), ctx.THIS().GetSymbol())
	}
}

// ExitNestedResource is called when production nestedResource is exited.
func (t *abuRewriter) ExitNestedResource(ctx *parser.NestedResourceContext) {
	p := ctx.ID(0).GetText()
	c := ctx.ID(1).GetText()
	var typ preprocessor.AbuType
	if ctx.EXT() == nil {
		sym := t.Symbol(p)
		if sym == nil {
			t.errCb(fmt.Errorf("cannot determine type of resource %s", p))
			return
		}
		typ = sym.Type()
	} else {
		syms := t.ExtSymbols(p)
		if len(syms) == 0 {
			t.errCb(fmt.Errorf("cannot determine type of remote resource %s", p))
			return
		}
		// TODO type checking in abuParser.go
		typ = syms[0].Type()
	}
	var v string
	m, present := t.Mappings[typ.TypeName]
	if !present {
		v = c
	} else {
		v, present = m.Fields[c]
		if !present {
			t.errCb(fmt.Errorf("configuration does not contain mapping for field %s", c))
			return
		}
	}
	var newID string
	if v == "" {
		newID = p
	} else {
		newID = fmt.Sprintf("%s_%s", p, v)
	}
	if ctx.THIS() != nil {
		t.rewriter.DeleteTokenDefault(ctx.THIS().GetSymbol(), ctx.THIS().GetSymbol())
	}
	t.rewriter.ReplaceTokenDefault(
		ctx.ID(0).GetSymbol(),
		ctx.SR_BRACKET().GetSymbol(),
		newID)
}

// EnterTask is called when production task is entered.
func (t *abuRewriter) EnterTask(ctx *parser.TaskContext) {
	t.inTask = true
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
	case ctx.ABSINT() != nil:
		t.rewriter.ReplaceTokenDefaultPos(ctx.ABSINT().GetSymbol(), "AbsInt(")
		t.rewriter.InsertAfterToken("default", ctx.Expression(0).GetStop(), ")")
	case ctx.ABSDEC() != nil:
		t.rewriter.ReplaceTokenDefaultPos(ctx.ABSDEC().GetSymbol(), "Abs(")
		t.rewriter.InsertAfterToken("default", ctx.Expression(0).GetStop(), ")")
	}
}

// ExitActions is called when production actions is exited.
func (t *abuRewriter) ExitActions(ctx *parser.ActionsContext) {
	for i := 0; ctx.COMMA(i) != nil; i++ {
		t.rewriter.ReplaceTokenDefaultPos(ctx.COMMA(i).GetSymbol(), ";")
	}
}

// ExitTask is called when production task is exited.
func (t *abuRewriter) ExitTask(ctx *parser.TaskContext) {
	t.inTask = false
}

// getAlteredText returns the text of the ParserRuleContext possibly altered
// through the TokenStreamRewriter default program.
func (t *abuRewriter) getAlteredText(ctx antlr.ParserRuleContext) string {
	return t.rewriter.GetText("default", antlr.Interval{
		Start: ctx.GetStart().GetTokenIndex(),
		Stop:  ctx.GetStop().GetTokenIndex(),
	})
}

// translate returns the text of the passed ParserRuleContext with
// the necessary modifications for matching GoAbU's syntax.
func (t *abuRewriter) translate(ctx antlr.ParserRuleContext) string {
	antlr.ParseTreeWalkerDefault.Walk(t, ctx)
	return t.getAlteredText(ctx)
}

// translateSimpleResource returns the simple resource's initialization in GoAbU's syntax.
func (t *abuRewriter) translateSimpleResource(id string, resource abuResource) string {
	r := fmt.Sprintf(".%s[\"%s\"] = ", typeToGoabu(resource.TypeName), id)
	if resource.initExpr != nil {
		r += t.translate(*resource.initExpr)
	} else {
		switch typeToGoabu(resource.TypeName) {
		case "Bool":
			r += "true"
		case "Integer":
			fallthrough
		case "Float":
			r += "0"
		case "Text":
			r += "\"\""
		case "Time":
			r += "time.Time{}"
		case "Other":
			r += "struct{}{}"
		}
	}
	return r
}

// translateComposedResource returns the composed resource's initialization in GoAbU's syntax.
func (t *abuRewriter) translateComposedResource(id string, nested map[string]abuResource) string {
	typ := t.Symbol(id).Type()
	mapp, present := t.Mappings[typ.TypeName]
	if !present {
		t.errCb(fmt.Errorf("custom type %s absent from configuration file", typ.TypeName))
		return ""
	}
	init := fmt.Sprintf(`.Add("%s", "%s"`, mapp.GoabuType, id)
	for _, argName := range mapp.Args {
		r := nested[argName]
		init = init + ", " + t.translate(*r.initExpr)
	}
	init += ")"
	return init
}

// translateRule returns the text of the passed ParserRuleContext quoted with
// backticks and with the necessary modifications for matching GoAbU's syntax.
func (t *abuRewriter) translateRule(ctx antlr.ParserRuleContext) string {
	return fmt.Sprintf("`%s`", t.translate(ctx))
}
