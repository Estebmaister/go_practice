package statistics

import "fmt"

// In probability theory, the birthday problem or birthday paradox
// concerns the probability that, in a set of n randomly chosen people,
// some pair of them will have the same birthday. By the pigeonhole
// principle, the probability reaches 100% when the number of people
// reaches 367 (since there are only 366 possible birthdays, including
// February 29).
// However, 99.9% probability is reached with just 70 people, and 50%
// probability with 23 people. These conclusions are based on the
// assumption that each day of the year (excluding February 29) is
// equally probable for a birthday.

func Proba() {
	p := 1.0
	for i := 1; i <= 153; i++ {
		p = p * (366 - float64(i)) / 365
		fmt.Println(i, ":", 1-p)
	}
}
