package main

import (
	"fmt"
	"reflect"
)

func NilDiffs() {
	fmt.Println(
		"Difference between a normal 'nil'",
		"and an empty interface:\n ",
	)
	AdultNil := adult(17) // Returns nil
	// AdultNil.randomMethod() // Panic nil pointer

	AdultInterface := adult(23)
	AdultInterface.randomMethod()

	fmt.Println(
		"Using reflect type,\n 'pure nil':",
		reflect.TypeOf(AdultNil),
		"vs 'interface nil':",
		reflect.TypeOf(&AdultInterface),
	) // different values

	fmt.Println("pure nil == interface nil?\n",
		AdultNil == AdultInterface,
	) // false

	fmt.Println("Printing both values:\n",
		AdultNil,
		AdultInterface,
	) // same value printed
}

type random interface {
	randomMethod()
}

type implRandom struct{}

// Needs to be assign to pointer to be called from interface
func (i *implRandom) randomMethod() {
	// Do nothing
}

func adult(n int) random {
	if n < 18 {
		return nil
	}
	var a *implRandom = nil
	return a
}
