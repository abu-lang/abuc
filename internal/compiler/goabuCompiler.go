// Copyright 2022 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

package compiler

import (
	_ "embed"
	"os"
	"path/filepath"
	"text/template"

	"github.com/abu-lang/abuc/internal/compiler/config"
	"github.com/abu-lang/abuc/internal/compiler/source"
	"github.com/abu-lang/abuc/internal/parser"
	"github.com/abu-lang/abuc/preprocessor"
	"github.com/antlr4-go/antlr/v4"
)

//go:embed goabu.go.tpl
var goabuTemplate string

type goabuCompiler struct {
	template *template.Template
	commonCompileInfo
	config config.Goabu
}

func makeGoabuCompiler(comm commonCompileInfo, cfgPath string) (compileStrategy, error) {
	conf, err := config.ReadFromFile(cfgPath)
	if err != nil {
		return nil, err
	}
	return goabuCompiler{
		commonCompileInfo: comm,
		config:            conf,
		template:          template.Must(template.New("goabu").Parse(goabuTemplate)),
	}, nil
}

func (g goabuCompiler) Compile(device string, stream preprocessor.TrivialStream, st preprocessor.DeviceSymbolTable) []error {
	var toks *antlr.CommonTokenStream = nil
	abup, errs := parseAbuProgram(stream, st, &toks)
	if len(errs) > 0 {
		return errs
	}
	prog, errs := source.MakeGoabu(abup, g.config, toks)
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
		Version string
		source.Goabu
	}
	g.template.Execute(f, tplData{Version: g.compilerVersion, Goabu: prog})
	return nil
}

// parseAbuProgram parses a text stream with a preprocessed abudsl program and returns an [source.Abu] struct.
// In the case of the occurrence of errors it returns a zero valued [source.Abu] and a not empty []error.
// If it receives as argument a not nil **antlr.CommonTokenStream and no errors are encountered, the argument
// is used as an output argument for returning the token stream used in the parsing.
func parseAbuProgram(stream preprocessor.TrivialStream, st preprocessor.DeviceSymbolTable, tsOut ...**antlr.CommonTokenStream) (source.Abu, []error) {
	errList := &source.ErrorHolder{}
	is := antlr.NewInputStream(stream.GetText(0, stream.Size()-1))
	lex := parser.NewAbuLexer(is)
	lex.RemoveErrorListeners()
	lex.AddErrorListener(errList)
	toks := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	par := parser.NewAbuParser(toks)
	par.RemoveErrorListeners()
	par.AddErrorListener(errList)
	par.BuildParseTrees = true
	list := source.NewAbuParser(st, errList.AddError)
	tree := par.Program()
	antlr.ParseTreeWalkerDefault.Walk(list, tree)
	if len(errList.Errors) > 0 {
		return source.Abu{}, errList.Errors
	}
	if len(tsOut) > 0 {
		if tsOut[0] != nil {
			*tsOut[0] = toks
		}
	}
	return list.Abu, nil
}
