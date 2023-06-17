// Copyright 2022 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

// Package preprocessor provides means for preprocessing abudsl programs.
package preprocessor

import (
	"fmt"

	"github.com/abu-lang/abuc/preprocessor/internal/parser"
	"github.com/antlr4-go/antlr/v4"
)

type abuPreproc struct {
	parser.BaseSugaredAbuParserListener
	rewriter *antlr.TokenStreamRewriter

	letExprs   map[string]string
	rule       *parser.EcaruleContext
	device     string
	customType string

	devices  []string
	hasRules map[string]map[string]struct{}

	symbols SymbolTable

	parseError func(error)
}

func newAbuPreproc(stream *antlr.CommonTokenStream, errorCallback func(error)) *abuPreproc {
	return &abuPreproc{rewriter: antlr.NewTokenStreamRewriter(stream),
		hasRules:   make(map[string]map[string]struct{}),
		symbols:    makeSymbolTable(),
		parseError: errorCallback,
	}
}

// EnterTypeDecl is called when production typeDecl is entered.
func (l *abuPreproc) EnterTypeDecl(ctx *parser.TypeDeclContext) {
	l.customType = ctx.ID().GetText()
	_, present := l.symbols.simple[l.customType]
	if present {
		l.parseError(fmt.Errorf("redefined %s identifier", l.customType))
		return
	}
	// register custom type
	l.symbols.simple[l.customType] = make(map[string]AbuType)
}

// EnterResField is called when production resField is entered.
func (l *abuPreproc) EnterResField(ctx *parser.ResFieldContext) {
	field := ctx.ID().GetText()
	_, present := l.symbols.simple[l.customType][field]
	if present {
		l.parseError(fmt.Errorf("redefined %s field in custom type %s", field, l.customType))
		return
	}
	t := AbuType{
		TypeName: ctx.Type_().GetText(),
		Readable: true,
		Writable: true,
	}
	if ctx.OUTPUT() != nil {
		t.Readable = false
	} else if ctx.INPUT() != nil {
		t.Writable = false
	}
	l.symbols.simple[l.customType][field] = t
}

// EnterDevice is called when production device is entered.
func (l *abuPreproc) EnterDevice(ctx *parser.DeviceContext) {
	devID := ctx.ID(0).GetText()
	_, present := l.symbols.simple[devID]
	if present {
		l.parseError(fmt.Errorf("redefined %s identifier", devID))
		return
	}
	// register device
	l.device = devID
	l.devices = append(l.devices, devID)
	l.symbols.simple[devID] = make(map[string]AbuType)
	l.symbols.nested[devID] = make(map[pair[string]]AbuType)
	l.hasRules[devID] = make(map[string]struct{})
	for i := 1; ctx.ID(i) != nil; i++ {
		l.hasRules[devID][ctx.ID(i).GetText()] = struct{}{}
	}
}

// ExitResDecl is called when production resDecl is exited.
func (l *abuPreproc) ExitResDecl(ctx *parser.ResDeclContext) {
	if ctx.CompResDecl() != nil {
		return
	}
	t := AbuType{
		TypeName: ctx.Type_().GetText(),
		Readable: true,
		Writable: true,
	}
	if ctx.OUTPUT() != nil {
		t.Readable = false
	} else if ctx.INPUT() != nil {
		t.Writable = false
	}
	l.symbols.simple[l.device][ctx.ID().GetText()] = t
}

// EnterCompResDecl is called when production compResDecl is entered.
func (l *abuPreproc) EnterCompResDecl(ctx *parser.CompResDeclContext) {
	typ := ctx.ID(0).GetText()
	id := ctx.ID(1).GetText()
	l.symbols.simple[l.device][id] = AbuType{TypeName: typ}
	syms := l.symbols.TypeInfo(typ)
	if syms == nil {
		l.parseError(fmt.Errorf("custom type %s not defined", typ))
		return
	}
	for r, sym := range syms {
		l.symbols.nested[l.device][pair[string]{Fst: id, Snd: r}] = sym.Type()
	}
}

// ExitSimpleResource is called when production simpleResource is exited.
func (l *abuPreproc) ExitSimpleResource(ctx *parser.SimpleResourceContext) {
	expr, present := l.letExprs[ctx.ID().GetText()]
	if !present {
		return
	}
	for _, d := range l.devices {
		if l.hasRule(d, l.rule) {
			l.rewriter.ReplaceToken(d,
				ctx.ID().GetSymbol(),
				ctx.ID().GetSymbol(),
				expr,
			)
		}
	}
}

// ExitNestedResource is called when production nestedResource is exited.
func (l *abuPreproc) ExitNestedResource(ctx *parser.NestedResourceContext) {
	for i := 0; i < 2; i++ {
		expr, present := l.letExprs[ctx.ID(i).GetText()]
		if !present {
			continue
		}
		for _, d := range l.devices {
			if l.hasRule(d, l.rule) {
				l.rewriter.ReplaceToken(d,
					ctx.ID(i).GetSymbol(),
					ctx.ID(i).GetSymbol(),
					expr,
				)
			}
		}
	}
}

// EnterEcarule is called when production ecarule is entered.
func (l *abuPreproc) EnterEcarule(ctx *parser.EcaruleContext) {
	l.letExprs = make(map[string]string)
	l.rule = ctx
	for _, d := range l.devices {
		if !l.hasRule(d, ctx) {
			l.rewriter.DeleteToken(d, ctx.GetStart(), ctx.GetStop())
		}
	}
}

// ExitLetDecl is called when production letDecl is exited.
func (l *abuPreproc) ExitLetDecl(ctx *parser.LetDeclContext) {
	for i := 0; ctx.ID(i) != nil; i++ {
		l.letExprs[ctx.ID(i).GetText()] = l.getTextFromContext(ctx.Expression(i))
	}
}

// ExitTask is called when production task is exited.
func (l *abuPreproc) ExitTask(ctx *parser.TaskContext) {
	if ctx.OWISE() == nil {
		return
	}
	typ := "for"
	if ctx.ALL() != nil {
		typ = "for all"
	}
	for _, d := range l.devices {
		if l.hasRule(d, l.rule) {
			l.rewriter.ReplaceToken(d,
				ctx.OWISE().GetSymbol(),
				ctx.OWISE().GetSymbol(),
				fmt.Sprintf("%s not (%s) do", typ, l.alteredCondition(d, ctx)),
			)
		}
	}
}

// ExitEcarule is called when production ecarule is exited.
func (l *abuPreproc) ExitEcarule(ctx *parser.EcaruleContext) {
	for _, d := range l.devices {
		if l.hasRule(d, l.rule) {
			if ctx.DEFAULT() != nil {
				l.rewriter.InsertAfterToken(d, ctx.Actions().GetStop(), "for true do "+l.getTextFromContext(ctx.Actions()))
				l.rewriter.DeleteToken(d, ctx.DEFAULT().GetSymbol(), ctx.Actions().GetStop())
			}
			if ctx.LET() != nil {
				l.rewriter.DeleteToken(d, ctx.LET().GetSymbol(), ctx.IN().GetSymbol())
			}
		}
	}
}

// ExitProgram is called when production program is exited.
func (l *abuPreproc) ExitProgram(ctx *parser.ProgramContext) {
	for c := 0; ctx.Device(c) != nil; c++ {
		for i, d := range l.devices {
			if i != c {
				l.rewriter.DeleteToken(d, ctx.Device(c).GetStart(), ctx.Device(c).GetStop())
			}
		}
	}
}

// getTextFromContext returns the parsed text corresponding to the passed Context.
func (l *abuPreproc) getTextFromContext(ctx antlr.RuleContext) string {
	return l.rewriter.GetTokenStream().GetTextFromRuleContext(ctx)
}

// alteredCondition returns the text of the TaskContext's condition possibly
// altered through the specified TokenStreamRewriter program (e.g. let substitutions).
func (l *abuPreproc) alteredCondition(prog string, ctx *parser.TaskContext) string {
	return l.rewriter.GetText(prog, antlr.Interval{
		Start: ctx.Expression().GetStart().GetTokenIndex(),
		Stop:  ctx.Expression().GetStop().GetTokenIndex(),
	})
}

// hasRule returns true if the device has the rule from the EcaruleContext.
func (l *abuPreproc) hasRule(dev string, ctx *parser.EcaruleContext) bool {
	_, res := l.hasRules[dev][ctx.ID().GetText()]
	return res
}
