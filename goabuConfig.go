// Copyright 2022 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

package main

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
