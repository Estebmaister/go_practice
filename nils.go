package main

import (
	"fmt"
	"reflect"
)

func nilDiffs() {
	fmt.Println(
		"Difference between a normal 'nil'",
		"and an empty interface",
	)
	AdultNil := adult(17)
	AdultInterface := adult(23)
	fmt.Println(
		"Using reflect type, 'pure nil':",
		reflect.TypeOf(AdultNil),
		"vs 'interface nil':",
		reflect.TypeOf(AdultInterface),
	)
	fmt.Println("Printing both values:",
		AdultNil,
		AdultInterface,
	)
}

type random interface {
	randomMethod()
}

type ir struct{}

func (i ir) randomMethod() {
	// Do nothing
}

func adult(n int) random {
	if n < 18 {
		return nil
	}
	var a *ir = nil
	return a
}
