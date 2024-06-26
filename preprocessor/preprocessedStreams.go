// Copyright 2022 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

package preprocessor

import (
	"errors"
	"fmt"

	"github.com/abu-lang/abuc/preprocessor/internal/parser"
	"github.com/antlr4-go/antlr/v4"
)

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

// PreprocessedStreamsFromFile takes a file name of an abudsl program and performs
// desugaring and splitting. It returns a map where the the keys are the device ids
// while the values are the corresponding desugared text streams and the program's
// symbol table. If an error is encountered during the preprocessing, a not empty
// []error is returned.
//
// NOTE: In the future it may return a map[string]antlr.CharStream
func PreprocessedStreamsFromFile(fileName string) (map[string]TrivialStream, SymbolTable, []error) {
	var cs antlr.CharStream
	var err error
	cs, err = antlr.NewFileStream(fileName)
	if err != nil {
		return nil, SymbolTable{}, []error{err}
	}
	return preprocessedStreamsFromCharStream(cs)
}

// PreprocessedStreams takes a string containing an abudsl program and performs
// desugaring and splitting. It returns a map where the the keys are the device ids
// while the values are the corresponding desugared text streams and the program's
// symbol table. If an error is encountered during the preprocessing, a not empty
// []error is returned.
//
// NOTE: In the future it may return a map[string]antlr.CharStream
func PreprocessedStreams(input string) (map[string]TrivialStream, SymbolTable, []error) {
	return preprocessedStreamsFromCharStream(antlr.NewInputStream(input))
}

// preprocessedStreamsFromCharStream takes a text stream of an abudsl program and performs
// desugaring and splitting. It returns a map where the the keys are the device ids
// while the values are the corresponding desugared text streams and the program's
// symbol table. If an error is encountered during the preprocessing, a not empty
// []error is returned.
//
// NOTE: In the future it may return a map[string]antlr.CharStream
func preprocessedStreamsFromCharStream(stream antlr.CharStream) (map[string]TrivialStream, SymbolTable, []error) {
	errLis := &errorHolder{}
	lex := parser.NewSugaredAbuLexer(stream)
	lex.RemoveErrorListeners()
	lex.AddErrorListener(errLis)
	toks := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	par := parser.NewSugaredAbuParser(toks)
	par.RemoveErrorListeners()
	par.AddErrorListener(errLis)
	par.BuildParseTrees = true
	lis := newAbuPreproc(toks, errLis.addError)
	tree := par.Program()
	if !<-walkParseTree(lis, tree) && len(errLis.errors) == 0 {
		errLis.addError(errors.New("error during parsing"))
	}
	if len(errLis.errors) > 0 {
		return nil, SymbolTable{}, errLis.errors
	}
	res := make(map[string]TrivialStream)
	for _, d := range lis.devices {
		res[d] = alteredTokenStream{program: d, rewriter: lis.rewriter}
	}
	return res, lis.symbols, nil
}

// walkParseTree performs a walk on the passed parse tree calling the methods of
// the ParseTreeListener argument. It returns a channel that passes a boolean value.
// The passed boolean is false only if a panic is throwed in the goroutine performing
// the walk.
func walkParseTree(list antlr.ParseTreeListener, tree antlr.Tree) <-chan bool {
	res := make(chan bool)
	go func() {
		defer func() {
			res <- nil == recover()
		}()
		antlr.ParseTreeWalkerDefault.Walk(list, tree)
	}()
	return res
}
