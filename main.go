package main

import (
	"fmt"

	"github.com/estebmaister/go_practice/concurrency"
	// "github.com/estebmaister/go_practice/channels"
	// "github.com/estebmaister/go_practice/server"
	"github.com/estebmaister/go_practice/errors"
)

func printValues[A, B any, C comparable](a, a1 A, b B, c C, d B) {
	fmt.Println(a, a1, b, c, d)
}
func main() {
	println("Starting from main\n")
	errors.Testing()

	printValues(1, 2, 3, "generics", 4.1)

	// Greeting()
	NilDiffs()

	concurrency.PrimesPipeline(10)
	// channels.LogChannel()

	// Server functions
	// fmt.Printf(
	// "Server listening on http://localhost%s:%v\n",
	// server.Host, server.Port
	// )
	// server.Run()
}

func Greeting() {
	var name string
	var age int
	fmt.Println("What's your name (One word) and age (in numbers)? Ex: Esteban 28")
	fmt.Scan(&name, &age)
	fmt.Printf("Hello %v, I can see that you are %d years old.\n", name, age)
	if age >= 18 {
		fmt.Print("And you can vote in the elections.\n\n")
	} else {
		fmt.Print("And you can't vote until you are 18.\n\n")
	}
}
