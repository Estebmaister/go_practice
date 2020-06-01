package loops

import (
	"fmt"
)

// Power calculate the power of a given number and exponent
func Power() {
	var exponent, base int
	fmt.Print("Enter Base:")
	fmt.Scanln(&base)
	fmt.Print("Enter exponent:")
	fmt.Scanln(&exponent)

	output := 1
	for exponent != 0 {
		output *= base
		exponent--
	}
	fmt.Printf("Output of power calculation is %d", output)

}
