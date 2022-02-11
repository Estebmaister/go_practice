package main

import (
	"fmt"
	"reflect"
	// "github.com/estebmaister/go_practice/channels"
	// "github.com/estebmaister/go_practice/server"
)

func main() {
	// greeting()
	// channels.LogChannel()

	// Server functions
	// fmt.Printf("Server listening on http://localhost%s:%v\n", server.Host, server.Port)
	// server.Run()

	fmt.Println("Difference betwen a normal 'nil' and an interface")
	AdultNil := adult(17)
	AdultInterface := adult(23)
	fmt.Println("Using reflect type, 'pure nil':",
		reflect.TypeOf(AdultNil), "vs 'interface nil':",
		reflect.TypeOf(AdultInterface))
	fmt.Println("Printing the value:", (AdultNil))
}

func greeting() {
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

type random interface {
	random1()
}

type ir struct{}

func (i ir) random1() {}

func adult(n int) random {
	if n < 18 {
		return nil
	}
	var a *ir = nil
	return a
}
