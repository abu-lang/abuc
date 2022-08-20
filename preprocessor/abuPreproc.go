// Package preprocessor provides means for preprocessing abudsl programs.
package preprocessor

import (
	"errors"
	"fmt"

	"github.com/abu-lang/abuc/preprocessor/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type abuPreproc struct {
	parser.BaseSugaredAbuParserListener
	rewriter *antlr.TokenStreamRewriter

	letExprs map[string]string
	rule     *parser.EcaruleContext

	devNumber int
	devices   map[string]int
	hasRules  map[string]map[string]struct{}

	parseError func(error)
}

func newAbuPreproc(stream *antlr.CommonTokenStream, errorCallback func(error)) *abuPreproc {
	return &abuPreproc{rewriter: antlr.NewTokenStreamRewriter(stream),
		devices:    make(map[string]int),
		hasRules:   make(map[string]map[string]struct{}),
		parseError: errorCallback,
	}
}

// ExitResource is called when production resource is exited.
func (l *abuPreproc) ExitResource(ctx *parser.ResourceContext) {
	expr, present := l.letExprs[ctx.ID().GetText()]
	if !present {
		return
	}
	for d := range l.devices {
		if l.hasRule(d, l.rule) {
			l.rewriter.ReplaceToken(d,
				ctx.ID().GetSymbol(),
				ctx.ID().GetSymbol(),
				expr,
			)
		}
	}
}

// ExitDevice is called when production device is exited.
func (l *abuPreproc) ExitDevice(ctx *parser.DeviceContext) {
	devID := ctx.ID(0).GetText()
	_, present := l.devices[devID]
	if present {
		l.parseError(errors.New("multiple devices share the same id"))
	}
	// register device
	l.devices[devID] = l.devNumber
	l.devNumber++
	// track device rules
	l.hasRules[devID] = make(map[string]struct{})
	for i := 1; ctx.ID(i) != nil; i++ {
		l.hasRules[devID][ctx.ID(i).GetText()] = struct{}{}
	}
}

// EnterEcarule is called when production ecarule is entered.
func (l *abuPreproc) EnterEcarule(ctx *parser.EcaruleContext) {
	l.letExprs = make(map[string]string)
	l.rule = ctx
	for d := range l.devices {
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
	for d := range l.devices {
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
	for d := range l.devices {
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
		for d, n := range l.devices {
			if n != c {
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
	return l.rewriter.GetText(prog, &antlr.Interval{
		Start: ctx.Expression().GetStart().GetTokenIndex(),
		Stop:  ctx.Expression().GetStop().GetTokenIndex(),
	})
}

// hasRule returns true if the device has the rule from the EcaruleContext.
func (l *abuPreproc) hasRule(dev string, ctx *parser.EcaruleContext) bool {
	_, res := l.hasRules[dev][ctx.ID(0).GetText()]
	return res
}
