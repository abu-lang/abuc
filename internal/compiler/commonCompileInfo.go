// Copyright 2022 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

package compiler

import (
	"os"

	"github.com/abu-lang/abuc/internal/compiler/version"
)

type commonCompileInfo struct {
	compilerVersion string
	output          string
	system          string
	target          string
}

func makeCommonCompileInfo(sys, tgt, out string) commonCompileInfo {
	return commonCompileInfo{
		compilerVersion: version.Get(),
		output:          out,
		system:          sys,
		target:          tgt,
	}
}

// Close is a no-op for implementing compileStrategy.Close
// in simple strategies.
func (i commonCompileInfo) Close() error {
	return nil
}

func outputFile(output, device, ext string) string {
	if output != "" && !os.IsPathSeparator(output[len(output)-1]) {
		return output + "-" + device + ext
	}
	return output + device + ext
}
