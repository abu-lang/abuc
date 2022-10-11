package main

import "github.com/abu-lang/abuc/preprocessor"

type compileStrategy interface {
	// compile takes the device ID and a text stream with the
	// preprocessed abudsl source code and outputs the compilation
	// result to a file. A not empty []error is returned if some
	// errors are encountered.
	compile(string, preprocessor.TrivialStream, preprocessor.DeviceSymbolTable) []error

	// Close may perform some clean up operations when the compilation
	// is terminated. It must be called ONCE after the completion of the
	// last call to compile (even when no call is performed).
	Close() error
}
