// Copyright 2022 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

package main

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/abu-lang/abuc/internal/parser"
	"github.com/abu-lang/abuc/preprocessor"
	"github.com/antlr4-go/antlr/v4"
)

//go:embed goabu.go.tpl
var goabuTemplate string

type goabuCompiler struct {
	commonCompileInfo
	config   goabuConfig
	template *template.Template
}

func makeGoabuCompiler(comm commonCompileInfo, cfgPath string) (compileStrategy, error) {
	conf, err := readGoabuConfig(cfgPath)
	if err != nil {
		return nil, err
	}
	return goabuCompiler{
		commonCompileInfo: comm,
		config:            conf,
		template:          template.Must(template.New("goabu").Parse(goabuTemplate)),
	}, nil
}

func (g goabuCompiler) compile(device string, stream preprocessor.TrivialStream, st preprocessor.DeviceSymbolTable) []error {
	var toks *antlr.CommonTokenStream = nil
	abup, errs := parseAbuProgram(stream, st, &toks)
	if len(errs) > 0 {
		return errs
	}
	prog, errs := makeGoabuProgram(abup, g.config, toks)
	if len(errs) > 0 {
		return errs
	}
	out := outputFile(g.output, device, ".go")
	os.MkdirAll(filepath.Dir(out), 0755)
	f, err := os.Create(out)
	if err != nil {
		return []error{err}
	}
	defer f.Close()
	type tplData struct {
		goabuProgram
		Version string
	}
	g.template.Execute(f, tplData{Version: g.compilerVersion, goabuProgram: prog})
	return nil
}

// parseAbuProgram parses a text stream with a preprocessed abudsl program and returns an abuProgram struct.
// In the case of the occurrence of errors it returns a zero valued abuProgram and a not empty []error.
// If it receives as argument a not nil **antlr.CommonTokenStream and no errors are encountered, the argument
// is used as an output argument for returning the token stream used in the parsing.
func parseAbuProgram(stream preprocessor.TrivialStream, st preprocessor.DeviceSymbolTable, tsOut ...**antlr.CommonTokenStream) (abuProgram, []error) {
	errList := &errorHolder{}
	is := antlr.NewInputStream(stream.GetText(0, stream.Size()-1))
	lex := parser.NewAbuLexer(is)
	lex.RemoveErrorListeners()
	lex.AddErrorListener(errList)
	toks := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	par := parser.NewAbuParser(toks)
	par.RemoveErrorListeners()
	par.AddErrorListener(errList)
	par.BuildParseTrees = true
	list := newAbuParser(st, errList.addError)
	tree := par.Program()
	antlr.ParseTreeWalkerDefault.Walk(list, tree)
	if len(errList.errors) > 0 {
		return abuProgram{}, errList.errors
	}
	if len(tsOut) > 0 {
		if tsOut[0] != nil {
			*tsOut[0] = toks
		}
	}
	return list.abuProgram, nil
}

type errorHolder struct {
	*antlr.DefaultErrorListener
	errors []error
}

func (h *errorHolder) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	h.addError(fmt.Errorf("line %d:%d %s", line, column, msg))
}

func (h *errorHolder) addError(err error) {
	h.errors = append(h.errors, err)
}
