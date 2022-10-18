// Copyright 2022 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

package preprocessor

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
)

// SymbolTable contains the information, usable by the compiler, regarding the resources
// of all the devices of an abudsl program.
type SymbolTable struct {
	simple map[string]map[string]AbuType
	nested map[string]map[pair[string]]AbuType
}

func makeSymbolTable() SymbolTable {
	return SymbolTable{simple: make(map[string]map[string]AbuType), nested: make(map[string]map[pair[string]]AbuType)}
}

// DeviceSymbolTable returns the  DeviceSymbolTable corresponding to the device ID argument.
func (t SymbolTable) DeviceSymbolTable(device string) (DeviceSymbolTable, error) {
	// simple contains also custom types
	_, present := t.nested[device]
	if !present {
		return DeviceSymbolTable{}, fmt.Errorf("")
	}
	return DeviceSymbolTable{device: device, simple: t.simple, nested: t.nested}, nil
}

// TypeInfo takes the name of a defined custom type and returns a map where the
// keys are the names of its fields and the values are the related symbol
// table entries. It returns nil if no entry is found.
func (t SymbolTable) TypeInfo(customType string) map[string]Symbol {
	typs, present := t.simple[customType]
	if !present {
		return nil
	}
	res := make(map[string]Symbol)
	for k, v := range typs {
		res[k] = abuSymbol(v)
	}
	return res
}

// Encode serializes the SymbolTable to an io.Writer.
func (t SymbolTable) Encode(w io.Writer) error {
	enc := gob.NewEncoder(w)
	err := enc.Encode(t.simple)
	if err != nil {
		return err
	}
	return enc.Encode(t.nested)
}

// DecodeSymbolTable deserializes a SymbolTable from an io.Reader.
func DecodeSymbolTable(r io.Reader) (SymbolTable, error) {
	res := SymbolTable{}
	dec := json.NewDecoder(r)
	err := dec.Decode(&res.simple)
	if err != nil {
		return SymbolTable{}, err
	}
	err = dec.Decode(&res.nested)
	if err != nil {
		return SymbolTable{}, err
	}
	return res, nil
}

// DeviceSymbolTable is a SymbolTable that ca be used to obtain information on
// the local and remote resources of a selected device.
type DeviceSymbolTable struct {
	device string
	simple map[string]map[string]AbuType
	nested map[string]map[pair[string]]AbuType
}

// TypeInfo takes the name of a defined custom type and returns a map where the
// keys are the names of its fields and the values are the related symbol
// table entries. It returns nil if no entry is found.
func (d DeviceSymbolTable) TypeInfo(customType string) map[string]Symbol {
	return SymbolTable{simple: d.simple, nested: d.nested}.TypeInfo(customType)
}

// Symbol takes as arguments the IDs of a local resource of the selected device
// and returns an entry that implements the Symbol interface.
// It returns nil if no entry is found.
func (d DeviceSymbolTable) Symbol(args ...string) Symbol {
	switch len(args) {
	case 1:
		t, present := d.simple[d.device][args[0]]
		if !present {
			return nil
		}
		return abuSymbol(t)
	case 2:
		t, present := d.nested[d.device][pair[string]{Fst: args[0], Snd: args[1]}]
		if !present {
			return nil
		}
		return abuSymbol(t)
	default:
		return nil
	}
}

// ExtSymbols takes as arguments the IDs of a remote resource and returns a []Symbol
// containing all the entries apart from the one related to the selected device.
func (d DeviceSymbolTable) ExtSymbols(args ...string) []Symbol {
	switch len(args) {
	case 1:
		return extSymbolsAux(d.simple, d.device, args[0])
	case 2:
		return extSymbolsAux(d.nested, d.device, pair[string]{Fst: args[0], Snd: args[1]})
	default:
		return nil
	}

}

func extSymbolsAux[M map[D]N, D comparable, N map[K]AbuType, K comparable](table M, device D, id K) []Symbol {
	var res []Symbol
	for d, m := range table {
		if d == device {
			continue
		}
		e, present := m[id]
		if present {
			res = append(res, abuSymbol(e))
		}
	}
	return res
}

// Symbol is the inferface implemented by the entries
// of a SymbolTable.
type Symbol interface {
	// Type returns the AbuType of the resource.
	Type() AbuType
	// private prevents users from implementing the interface.
	private()
}

type pair[T any] struct {
	Fst T
	Snd T
}

type abuSymbol AbuType

func (s abuSymbol) Type() AbuType {
	return AbuType(s)
}

func (s abuSymbol) private() {}
