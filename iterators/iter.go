package iterators

import (
	"iter"
	"slices"
)

type User struct {
	Active   bool
	Username string
	Points   int
}

// Iter-based generic filter implementation
func ActiveUsers[T any](u []T, filter func(e T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, user := range u {

			if !filter(user) {
				continue
			}

			if !yield(user) {
				return
			}

		}
	}
}

// Iter-based generic implementation
func GetTotalActiveWithIter(u []User) int {
	result := int(0)

	for user := range ActiveUsers(u, func(u User) bool {
		return u.Active
	}) {
		result += user.Points
	}

	return result
}

// Range-based implementation
func ActiveUsersRange(u []User) []User {
	result := []User{}

	// optimizing here to give this function
	// a fair chance.
	for index := range u {

		if !u[index].Active {
			continue
		}

		result = append(result, u[index])
	}

	return result
}

// Range-based implementation
func GetTotalActiveWithRange(u []User) int {
	total := 0
	for _, user := range ActiveUsersRange(u) {
		total += user.Points
	}
	return total
}

// Slices-based implementation
func ActiveUsersSlices(u []User) []User {

	// the slice needs to be copied
	// because the slices package will
	// modify the slice you passed.
	// in other words, the variable you called this function with
	// will be modified.
	copySlice := make([]User, len(u))
	copy(copySlice, u)

	return slices.DeleteFunc(copySlice, func(user User) bool {
		return !user.Active
	})
}

// Slices-based implementation
func GetTotalActiveWithSlices(u []User) int {
	total := 0

	for _, user := range ActiveUsersSlices(u) {
		total += user.Points
	}
	return total
}
