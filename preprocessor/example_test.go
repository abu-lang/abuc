// Copyright 2022 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

package preprocessor_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/abu-lang/abuc/preprocessor"
)

func TestExample(t *testing.T) {
	// run with 'go test -v'
	input := `# AbU devices definition.

hvac : "An HVAC control system" {
	physical output boolean heating = false
	physical output boolean conditioning = false
	logical integer temperature = 0
	logical integer humidity = 0
	physical input boolean airButton
	logical string node = "hvac"
  where
	not (conditioning and heating)
} has cool warm dry stopAir

tempSens : "A temperature sensor" {
	physical input integer temperature
	logical string node = "tempSens"
} has notifyTemp

humSens : "A humidity sensor" {
	physical input integer humidity
	logical string node = "humSens"
} has notifyHum

calculator : "A calculator" {
	logical integer x = 1
	logical integer y = 2
	logical integer z = 0
	logical integer result = 0
} has calculatorLet copyX

\@
	AbU (ECA) rules definition.
	Rules can be referenced by multiple devices.
@\

rule copyX
	on x
	default z = x

rule calculatorLet
    on x y
      let sum := (x + y); diff := (x - y) in
    for (sum > 0)
        do result = sum * diff
	owise
		result = 12

rule cool
	on temperature
	for (this.temperature < 18)
		do this.heating = true

rule warm
	on temperature
	for (this.temperature > 27)
		do this.heating = false

rule dry
	on humidity temperature
	for (2 + 0.5 * this.temperature < this.humidity and 38 - this.temperature < this.humidity)
		do this.conditioning = true
	
rule stopAir
	on airButton
	for (this.airButton)
		do this.conditioning = false

rule notifyTemp
	on temperature
	for all (ext.node == "hvac")
		do ext.temperature = this.temperature

rule notifyHum
	on humidity
	for all (ext.node == "hvac")
		do ext.humidity = this.humidity`

	streams, typs, errs := preprocessor.PreprocessedStreams(input)
	if len(errs) > 0 {
		t.Fatal(errs[0].Error())
	}
	for d, s := range streams {
		fmt.Printf("##################### Code for %s:\n", d)
		fmt.Println(s.GetText(0, s.Size()-1))
	}
	fmt.Println("##################### Types")
	typs.Encode(os.Stdout)
}
