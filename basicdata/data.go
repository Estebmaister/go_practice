package basicdata

import "fmt"

// PrintData is a showcase of the data types in go
func PrintData() {
	var int16Variable int16 = 12
	var intVariable = 85 // Type inference example
	var uintVariable uint = 40
	var hexVariable = 0xBC
	var octalVariable = 013
	fmt.Printf("%d, %d, %d, %#x, %#o\n", int16Variable, intVariable, uintVariable, hexVariable, octalVariable)
	var byteVariable byte = 'C'
	var runeVariable rune = 'a'
	fmt.Printf("%c - %d and %c - %U\n", byteVariable, byteVariable, runeVariable, runeVariable)
	var complexVariable = 1 + 6i
	var floatVariable = 47895.587
	fmt.Printf("%f - %f and %f - %f\n", complexVariable, complexVariable, floatVariable, floatVariable)

}
