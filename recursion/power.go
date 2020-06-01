package recursion

import (
	"fmt"
)

// RecursivePower calculate the power of a given number and exponent
func RecursivePower(base int, exponent int) int {
	// if the exponent is not zero, call the function again with exp-1
	if exponent != 0 {
		return (base * RecursivePower(base, exponent-1))
	}
	// if the exponent is zero, return 1
	return 1
}

// Power function invoke an input/output cycle to calculate power of the given number
func Power() {
	var exponent, base int
	fmt.Print("Enter Base:")
	fmt.Scanln(&base)
	fmt.Print("Enter exponent:")
	fmt.Scanln(&exponent)

	output := RecursivePower(base, exponent)

	fmt.Printf("Output of power calculation is %d", output)

}
