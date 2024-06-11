package source

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type ErrorHolder struct {
	*antlr.DefaultErrorListener
	Errors []error
}

func (h *ErrorHolder) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	h.AddError(fmt.Errorf("line %d:%d %s", line, column, msg))
}

func (h *ErrorHolder) AddError(err error) {
	h.Errors = append(h.Errors, err)
}
