package main

import (
	"fmt"

	"./statistics"
)

func main() {
	fmt.Print("Hello World\n")

	var name string
	var age int
	fmt.Println("What's your name and age?")
	fmt.Scan(&name, &age)
	fmt.Printf("You entered %v and %d.\n", name, age)

	statistics.Proba()
}
