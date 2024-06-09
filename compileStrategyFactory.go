// Copyright 2022 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

package main

import "errors"

func makeCompileStrategy(sys, tgt, out, cfg string) (compileStrategy, error) {
	comm := makeCommonCompileInfo(sys, tgt, out)
	switch tgt {
	case "abu":
		return makeAbuCompiler(comm)
	case "go":
		return makeGoabuCompiler(comm, cfg)
	case "arm64", "amd64", "arm":
		return makeMachineCodeCompiler(comm, cfg)
	default:
		panic(errors.New("no compile strategy for target:" + tgt))
	}
}
