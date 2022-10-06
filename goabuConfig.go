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
