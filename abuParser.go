// Copyright 2022 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"

	"github.com/abu-lang/abuc/internal/parser"
	"github.com/abu-lang/abuc/preprocessor"
)

type abuProgram struct {
	preprocessor.DeviceSymbolTable
	id                string
	invariant         *parser.IExpressionContext
	simpleResources   map[string]abuResource
	composedResources map[string]map[string]abuResource
	rules             map[string]*parser.IEcaruleContext
	tick              int
}

func makeAbuProgram(st preprocessor.DeviceSymbolTable) abuProgram {
	return abuProgram{
		DeviceSymbolTable: st,
		simpleResources:   make(map[string]abuResource),
		composedResources: make(map[string]map[string]abuResource),
		rules:             make(map[string]*parser.IEcaruleContext),
		tick:              1,
	}
}

type abuResource struct {
	preprocessor.AbuType
	initExpr *parser.IExpressionContext
}

type abuParser struct {
	parser.BaseAbuParserListener

	abuProgram
	inRuleDefs bool
	inTask     bool
	parseError func(error)
}

func newAbuParser(st preprocessor.DeviceSymbolTable, errorCallback func(error)) *abuParser {
	return &abuParser{
		abuProgram: makeAbuProgram(st),
		parseError: errorCallback,
	}
}

// EnterResDecl is called when production resDecl is entered.
func (p *abuParser) EnterResDecl(ctx *parser.ResDeclContext) {
	if ctx.CompResDecl() != nil {
		return
	}
	id := ctx.ID().GetText()
	sym := p.Symbol(id)
	if sym == nil {
		p.parseError(fmt.Errorf("resource %s missing from symbol table", id))
		return
	}
	r := abuResource{AbuType: sym.Type()}
	if ctx.Expression() != nil {
		expr := ctx.Expression()
		r.initExpr = &expr
	}
	p.simpleResources[id] = r
}

// EnterCompResDecl is called when production compResDecl is entered.
func (p *abuParser) EnterCompResDecl(ctx *parser.CompResDeclContext) {
	typ := ctx.ID(0).GetText()
	id := ctx.ID(1).GetText()
	decls := p.TypeInfo(typ)
	if decls == nil {
		p.parseError(fmt.Errorf("custom type %s is undefined", typ))
		return
	}
	c := make(map[string]abuResource)
	for k, v := range decls {
		c[k] = abuResource{AbuType: v.Type()}
	}
	for i := 2; ctx.ID(i) != nil; i++ {
		resID := ctx.ID(i).GetText()
		expr := ctx.Expression(i - 2)
		r := c[resID]
		r.initExpr = &expr
		c[resID] = r
	}
	p.composedResources[id] = c
}

// ExitDevice is called when production device is exited.
func (p *abuParser) ExitDevice(ctx *parser.DeviceContext) {
	// register device id
	p.id = ctx.ID(0).GetText()
	// track device rules
	for i := 1; ctx.ID(i) != nil; i++ {
		p.rules[ctx.ID(i).GetText()] = nil
	}
	// register invariant
	if ctx.WHERE() != nil {
		expr := ctx.Expression()
		p.invariant = &expr
	}

	p.inRuleDefs = true
}

// EnterEcarule is called when production ecarule is entered.
func (p *abuParser) EnterEcarule(ctx *parser.EcaruleContext) {
	r, present := p.rules[ctx.ID().GetText()]
	if !present {
		return
	}
	if r != nil {
		p.parseError(fmt.Errorf("rule %s has multiple definitions", ctx.ID().GetText()))
	}
	var rule parser.IEcaruleContext = ctx
	p.rules[ctx.ID().GetText()] = &rule
}

// ExitSimpleResource is called when production simpleResource is exited.
func (p *abuParser) ExitSimpleResource(ctx *parser.SimpleResourceContext) {
	if !p.inTask && ctx.EXT() != nil {
		p.parseError(fmt.Errorf("resource %s cannot be remote in this context", ctx.ID().GetText()))
	}
}

// ExitNestedResource is called when production nestedResource is exited.
func (p *abuParser) ExitNestedResource(ctx *parser.NestedResourceContext) {
	if !p.inTask && ctx.EXT() != nil {
		p.parseError(fmt.Errorf("resource %s[%s] cannot be remote in this context", ctx.ID(0).GetText(), ctx.ID(1).GetText()))
	}
}

// EnterTask is called when production task is entered.
func (p *abuParser) EnterTask(ctx *parser.TaskContext) {
	p.inTask = true
}

// ExitTask is called when production task is exited.
func (p *abuParser) ExitTask(ctx *parser.TaskContext) {
	p.inTask = false
}

// ExitProgram is called when production program is exited.
func (p *abuParser) ExitProgram(ctx *parser.ProgramContext) {
	for n, r := range p.rules {
		if r == nil {
			p.parseError(fmt.Errorf("missing definition of rule %s", n))
			delete(p.rules, n)
		}
	}
}
