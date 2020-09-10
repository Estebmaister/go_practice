package main

import (
	"fmt"
	"reflect"
	// "go-starting/channels"
	// "go-starting/server"
	// "go-starting/statistics"
)

func main() {
	// greeting()
	// statistics.Proba()
	// channels.LogChannel()

	// Server functions
	// fmt.Printf("Server listening on http://localhost%s:%v\n", server.Host, server.Port)
	// server.Run()

	p := adult(23)
	fmt.Println(reflect.TypeOf(p))
	fmt.Println((p))
}

func greeting() {
	var name string
	var age int
	fmt.Println("What's your name (One word) and age (in numbers)?")
	fmt.Scan(&name, &age)
	fmt.Printf("Hello %v, I can see that you are %d years old.\n", name, age)
	if age >= 18 {
		fmt.Print("And you can vote now.\n\n")
	} else {
		fmt.Print("And you can´t vote until you are 18.\n\n")
	}

}

type random interface {
	random1()
}

type ir struct{}

func (i ir) random1() {}

func adult(n int) interface{} {
	if n < 18 {
		return nil
	}
	var a *ir = nil
	return a
}
