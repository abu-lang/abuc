package main

import (
	"fmt"

	"github.com/abu-lang/abuc/parser"
	"github.com/abu-lang/abuc/preprocessor"
)

type abuProgram struct {
	preprocessor.DeviceSymbolTable
	id        string
	invariant *parser.IExpressionContext
	resources map[string]abuResource
	rules     map[string]*parser.IEcaruleContext
	tick      int
}

type abuResource struct {
	preprocessor.AbuType
	initExpr *parser.IExpressionContext
}

type abuParser struct {
	parser.BaseAbuParserListener

	abuProgram
	inRuleDefs bool
	parseError func(error)
}

func newAbuParser(st preprocessor.DeviceSymbolTable, errorCallback func(error)) *abuParser {
	p := abuProgram{DeviceSymbolTable: st}
	res := &abuParser{
		abuProgram: p,
		parseError: errorCallback,
	}
	res.resources = make(map[string]abuResource)
	res.rules = make(map[string]*parser.IEcaruleContext)
	res.tick = 1
	return res
}

// EnterResDecl is called when production resDecl is entered.
func (p *abuParser) EnterResDecl(ctx *parser.ResDeclContext) {
	t := preprocessor.AbuType{
		Readable: true,
		Writable: true,
		TypeName: ctx.Type().GetText(),
	}
	r := abuResource{AbuType: t}
	if ctx.Expression() != nil {
		expr := ctx.Expression()
		r.initExpr = &expr
	}
	if ctx.OUTPUT() != nil {
		r.Readable = false
	} else if ctx.INPUT() != nil {
		r.Writable = false
	}
	p.resources[ctx.ID().GetText()] = r
}

// EnterResource is called when production resource is entered.
func (p *abuParser) EnterResource(ctx *parser.ResourceContext) {
	if !p.inRuleDefs {
		return
	}
	if ctx.EXT() == nil && !p.resources[ctx.ID().GetText()].Readable {
		p.parseError(fmt.Errorf("resource %s cannot be accessed", ctx.ID().GetText()))
	}
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
	r, present := p.rules[ctx.ID(0).GetText()]
	if !present {
		return
	}
	if r != nil {
		p.parseError(fmt.Errorf("rule %s has multiple definitions", ctx.ID(0).GetText()))
	}
	var rule parser.IEcaruleContext = ctx
	p.rules[ctx.ID(0).GetText()] = &rule
}

// EnterAssignment is called when production assignment is entered.
func (p *abuParser) EnterAssignment(ctx *parser.AssignmentContext) {
	if ctx.EXT() == nil && !p.resources[ctx.ID().GetText()].Writable {
		p.parseError(fmt.Errorf("resource %s cannot be modified", ctx.ID().GetText()))
	}
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
