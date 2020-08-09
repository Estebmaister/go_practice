package main

import (
	"fmt"
	"strconv"
)

func convertDecimalToBinary(number int) int {
	binary := 0
	counter := 1
	remainder := 0

	for number != 0 {
		remainder = number % 2
		number = number / 2
		binary += remainder * counter
		counter *= 10

	}
	return binary
}

func main() {
	var decimal int64

	fmt.Print("Enter Decimal Number: ")
	fmt.Scanln(&decimal)

	output := strconv.FormatInt(decimal, 2)
	fmt.Printf("Output %s\n\n", output)

	var decimalMath int

	fmt.Print("Enter Decimal Number: ")
	fmt.Scanln(&decimalMath)

	fmt.Printf("Output %d", convertDecimalToBinary(decimalMath))

}
