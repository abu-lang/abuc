// Copyright 2022 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"encoding/json"
	"os"
)

type goabuConfig struct {
	Imports          []string
	AdditionalFiles  []string
	StateInitializer string
	Mappings         map[string]goabuMapping
}

type goabuMapping struct {
	GoabuType string
	Fields    map[string]string
	Args      []string
}

// readGoabuConfig tries to read a JSON encoded goabuConfig
// from a file given the file path.
func readGoabuConfig(path string) (goabuConfig, error) {
	var res goabuConfig
	if path == "" {
		return res, nil
	}
	f, err := os.Open(path)
	if err != nil {
		return res, err
	}
	defer f.Close()
	err = json.NewDecoder(f).Decode(&res)
	if err != nil {
		return res, err
	}
	return res, nil
}
