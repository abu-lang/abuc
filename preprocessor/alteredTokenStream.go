// Copyright 2022 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

package preprocessor

import "github.com/antlr4-go/antlr/v4"

type alteredTokenStream struct {
	program  string
	rewriter *antlr.TokenStreamRewriter
}

func (s alteredTokenStream) Size() int {
	return s.rewriter.GetTokenStream().Size()
}

func (s alteredTokenStream) GetText(start, stop int) string {
	return s.rewriter.GetText(s.program, antlr.Interval{Start: start, Stop: stop})
}
