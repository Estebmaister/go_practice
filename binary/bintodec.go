package main

import (
	"fmt"
	"math"
	"strconv"
)

func convertBinaryToDecimal(number int) int {
	decimal := 0
	counter := 0.0
	remainder := 0

	for number != 0 {
		remainder = number % 10
		decimal += remainder * int(math.Pow(2.0, counter))
		number = number / 10
		counter++
	}
	return decimal
}

func main() {
	var binary string
	fmt.Print("Enter Binary Number: ")
	fmt.Scanln(&binary)
	output, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Output with strconv %d\n\n", output)

	var binaryMath int
	fmt.Print("Enter Binary Number: ")
	fmt.Scanln(&binaryMath)
	fmt.Printf("Output with math %d", convertBinaryToDecimal(binaryMath))

}
