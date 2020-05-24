package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Create a new seed for every run
	rand.Seed(time.Now().UnixNano())

	amountLeft := rand.Intn(10000)

	fmt.Println("amountLeft is: ", amountLeft)

	if amountLeft > 5000 {
		fmt.Println("What should I spend this on?")
	} else {
		fmt.Println("Where did all my money go?")
	}

	// Some conditions and switches!
	if lessonLearned := true; lessonLearned {
		fmt.Println("Great job! You can continue on to the next exercise.")
	} else {
		fmt.Println("Practice makes perfect.")
	}

	if success := true; success {
		fmt.Println("We're rich!")
	} else {
		fmt.Println("Where did we go wrong?")
	}

	amountStolen := 50000

	switch numOfThieves := 5; numOfThieves {
	case 1:
		fmt.Println("I'll take all $", amountStolen)
	case 2:
		fmt.Println("Everyone gets $", amountStolen/2)
	case 3:
		fmt.Println("Everyone gets $", amountStolen/3)
	case 4:
		fmt.Println("Everyone gets $", amountStolen/4)
	case 5:
		fmt.Println("Everyone gets $", amountStolen/5)
	default:
		fmt.Println("There's not enough to go around...")
	}

}
