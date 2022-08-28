package main

import "github.com/abu-lang/abuc/preprocessor"

type compileStrategy interface {
	// compile takes the device ID and a text stream with the
	// preprocessed abudsl source code and outputs the compilation
	// result to a file. A not empty []error is returned if some
	// errors are encountered.
	compile(string, preprocessor.TrivialStream) []error
}
