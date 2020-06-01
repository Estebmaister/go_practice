package loops

import "fmt"

// Looping for a sequence of natural numbers
func Looping() {
LOOPSECTION:
	for i := 1; i <= 10; i++ {
		if i == 2 {
			continue LOOPSECTION
		}
		if i == 9 {
			break
		}
		if i == 10 {
			goto LabelName
		}
		fmt.Print(i)
	}
LabelName:
	fmt.Print("End")
}
