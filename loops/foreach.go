package loops

import "fmt"

// ForEach is an example of looping through Numeric Arrays,
// String Arrays, Unicode characters of a string and Maps
func ForEach() {

	// Create Numeric array
	numbers := [3]int{6, 9, 3}
	// Iterate numeric array
	for key, value := range numbers {
		fmt.Printf("%d = %d\n", key, value)
	}
	// Create String array
	strs := [3]string{"one", "two", "three"}
	// Iterate String array
	for key, value := range strs {
		fmt.Printf("%d = %s\n", key, value)
	}

	/* create a map*/
	emps := map[string]string{"1": "John", "2": "Franc", "3": "Kiran"}

	/* Iterate map using keys*/
	for emp := range emps {
		fmt.Println("Emp Id: ", emp, " Name: ", emps[emp])
	}
	/* Iterate Unicode charactes of a String*/
	for index, character := range "本語日" {
		fmt.Printf("%#U  position %d\n", character, index)
	}

}
