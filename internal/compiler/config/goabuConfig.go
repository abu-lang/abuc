// Copyright 2022 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"encoding/json"
	"os"
)

type Goabu struct {
	Mappings         map[string]GoabuMapping
	StateInitializer string
	Imports          []string
	AdditionalFiles  []string
}

type GoabuMapping struct {
	GoabuType string
	Fields    map[string]string
	Args      []string
}

// ReadFromFile tries to read a JSON encoded [Goabu]
// from a file given the file path.
func ReadFromFile(path string) (Goabu, error) {
	var res Goabu
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
