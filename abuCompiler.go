// Copyright 2022 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"os"

	"github.com/abu-lang/abuc/preprocessor"
)

type abuCompiler struct {
	commonCompileInfo
}

func makeAbuCompiler(comm commonCompileInfo) (compileStrategy, error) {
	return abuCompiler{commonCompileInfo: comm}, nil
}

func (a abuCompiler) compile(device string, stream preprocessor.TrivialStream, st preprocessor.DeviceSymbolTable) []error {
	f, err := os.Create(outputFile(a.output, device, ".abu"))
	if err != nil {
		return []error{err}
	}
	defer f.Close()
	_, err = fmt.Fprintf(f, "# Code generated by abuc version %s.\n", a.compilerVersion)
	if err != nil {
		return []error{err}
	}
	_, err = fmt.Fprint(f, stream.GetText(0, stream.Size()))
	if err != nil {
		return []error{err}
	}
	return nil
}