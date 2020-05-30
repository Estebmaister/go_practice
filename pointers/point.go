package main

import "fmt"

// brainwash function have a pointer parameter
func brainwash(saying *string) {
	// Dereference saying below:
	*saying = "Beep Boop."
}

func main() {
	greeting := "Hello there!"
	// Calling brainwash() with a pointer:
	brainwash(&greeting)
	fmt.Println("greeting is now:", greeting)

	lyrics := "Moments so dear"
	var pointerForStr *string = &lyrics
	*pointerForStr = "Journeys to plan"
	fmt.Printf("The lyrics variable now contains %v", lyrics) // Prints: Journeys to plan
}
