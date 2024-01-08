package main

import "fmt"

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
