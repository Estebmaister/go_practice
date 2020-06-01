package functions

import "fmt"

func getLikesAndShares(postID int) (int, int) {
	var likesForPost, sharesForPost int
	switch postID {
	case 1:
		likesForPost = 5
		sharesForPost = 7
	case 2:
		likesForPost = 3
		sharesForPost = 11
	case 3:
		likesForPost = 22
		sharesForPost = 1
	case 4:
		likesForPost = 7
		sharesForPost = 9
	}
	fmt.Println("Likes: ", likesForPost, "Shares: ", sharesForPost)

	return likesForPost, sharesForPost
}

// Example run a function with a given integer between 1 and 4
func Example() {
	var likes, shares int

	likes, shares = getLikesAndShares(2)

	if likes > 5 {
		fmt.Println("Woohoo! We got some likes.")
	}
	if shares > 10 {
		fmt.Println("We went viral!")
	}
}
