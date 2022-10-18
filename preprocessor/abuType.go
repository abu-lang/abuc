// Copyright 2022 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

package preprocessor

// AbuType represents the type of an abudsl expression.
type AbuType struct {
	// TypeName contains the type of the expression.
	TypeName string
	// Readable specifies if all the expression's terms are accessible.
	Readable bool
	// Writable specifies whether the expression is an l-value of a modifiable resource.
	Writable bool
}
