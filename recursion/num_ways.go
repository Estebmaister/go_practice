package recursion

import "math"

type Currency struct {
	denominations []int
	minCoins      map[int]int
}

// NewCurrency creates a new currency object with the given denominations
func NewCurrency(denominations []int) *Currency {
	return &Currency{denominations, map[int]int{0: 0, 1: 1}}
}

// NumWays returns the number of ways to make change for a given amount
func (x Currency) NumWays(target int) int {
	return x.NumWaysRecurse(target, 0)
}

// NumWaysRecurse returns the number of ways to make change for a given amount
// max represents the index of the denomination to start from
func (x Currency) NumWaysRecurse(amount, max int) int {
	if max == len(x.denominations)-1 {
		return 1
	}

	sum := 0

	// for each denomination, calculate the number of ways to make change
	// starting from the maximum denomination
	for i := max; i < len(x.denominations); i++ {
		// if the amount is less than the denomination, skip it
		if amount < x.denominations[i] {
			continue
		}

		remaining := amount - x.denominations[i]

		if remaining == 0 {
			sum += 1
		}
		if remaining > 0 {
			sum += x.NumWaysRecurse(remaining, i)
		}
	}

	return sum
}

// MinChange returns the minimum number of coins needed to make change for a given amount
func (x Currency) MinChange(amount int) int {
	// if the amount is already in the map, return it
	val := x.minCoins[amount]
	if val != 0 {
		return val
	}

	val = math.MaxInt32

	// for each denomination, calculate the minimum number of coins needed to make change
	for _, d := range x.denominations {
		// if the amount is less than the denomination, skip it
		if amount < d {
			continue
		}
		val = min(val, x.MinChange(amount-d))
	}

	x.minCoins[amount] = val + 1
	return val + 1
}
