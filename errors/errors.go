package errors

import "fmt"

type ErrorNegro struct{}

type ErrorAsiatico struct{}

func (e ErrorNegro) Error() string {
	return "No me gustan los negros"
}

func (e ErrorAsiatico) Error() string {
	return "No me gustan los chinos"
}

func Errors() {
	err := SoyRacista("negro")

	if err != nil {
		switch err.(type) {
		case ErrorNegro:
			fmt.Println("Huele a negro")
		case ErrorAsiatico:
			fmt.Println("Lumpia")
		}
		return
	}

	fmt.Println("Todo bien")
}

func SoyRacista(color string) error {
	if color == "negro" {
		return ErrorNegro{}
	}
	if color == "amarillo" {
		return ErrorAsiatico{}
	}
	return nil
}
