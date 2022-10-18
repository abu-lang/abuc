// Copyright 2022 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bytes"
	"errors"
	"io"
	"os"

	"github.com/abu-lang/abuc/preprocessor"
)

// preprocess behaves as preprocessor.PreprocessedStreamsFromFile but accepts a callback that
// is called if some errors are encountered. If the file name == "" the input is read from
// the standard input.
func preprocess(fileName string, errCb func([]error)) (map[string]preprocessor.TrivialStream, preprocessor.SymbolTable) {
	if fileName == "" {
		istr, err := inputStringFromStdin()
		if err != nil {
			errCb([]error{errors.New("error in reading from stdin: " + err.Error())})
			return nil, preprocessor.SymbolTable{}
		}
		res, st, errs := preprocessor.PreprocessedStreams(istr)
		if len(errs) > 0 {
			errCb(errs)
			return nil, preprocessor.SymbolTable{}
		}
		return res, st
	}
	res, st, errs := preprocessor.PreprocessedStreamsFromFile(fileName)
	if len(errs) > 0 {
		errCb(errs)
		return nil, preprocessor.SymbolTable{}
	}
	return res, st
}

// inputStringFromStdin returns the content of the
// standard input as a string.
func inputStringFromStdin() (string, error) {
	buf := bytes.NewBuffer(nil)
	_, err := io.Copy(buf, os.Stdin)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
