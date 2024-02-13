package tools

import (
	"fmt"
	"math/rand"
)

type Accounts int

const (
	Esteb Accounts = iota
	Sharon
	Rob
	Anthony
	_
	Emma
)

func (a Accounts) String() string {
	return [...]string{
		"Esteb", "Sharon", "Rob", "Anthony", "Unknown", "Emma",
	}[a]
}

// Example run a function with a random integer between 0 and 3
// representing an Accounts enum type
// and prints one or two lines to the terminal.
func Example() {
	var likes, shares int

	var influencer Accounts = Accounts(rand.Intn(6))
	likes, shares = getLikesAndShares(influencer)

	if likes > 5 {
		fmt.Printf("Wahoo! We got some likes for %s.\n", influencer)
	}
	if shares > 10 {
		fmt.Println("We went viral!")
	}
}

func getLikesAndShares(account Accounts) (int, int) {
	var likesForPost, sharesForPost int
	switch account {
	case Esteb:
		likesForPost = 35
		sharesForPost = 7
	case Sharon:
		likesForPost = 3
		sharesForPost = 11
	case Rob:
		likesForPost = 22
		sharesForPost = 1
	case Anthony:
		likesForPost = 7
		sharesForPost = 9
	default:
		likesForPost = 0
		sharesForPost = 0
	}
	fmt.Println("Likes: ", likesForPost, "Shares: ", sharesForPost)

	return likesForPost, sharesForPost
}
