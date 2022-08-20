package preprocessor

import "github.com/antlr/antlr4/runtime/Go/antlr"

type alteredTokenStream struct {
	program  string
	rewriter *antlr.TokenStreamRewriter
}

func (s alteredTokenStream) Size() int {
	return s.rewriter.GetTokenStream().Size()
}

func (s alteredTokenStream) GetText(start, stop int) string {
	return s.rewriter.GetText(s.program, &antlr.Interval{Start: start, Stop: stop})
}
