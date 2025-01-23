package main

import (
	"fmt"
	"sort"
)

func main() {
	testCases := [][]int{
		{5, 1, 4, 8, 2, 20},       // Should return 35 → [5, 2, 8, 20]
		{5, 1, 3, 7, 2, 10, 5, 3}, // 19
		{3, 2, 7, 10},             // Should return 19 → [2, 7, 10]
		{1, 1, 1, 1, 1, 1},        // Should return 1 → [1]
		{20, 5, 4, 3, 2, 1},       // Should return 35 → [20, 5, 4, 3, 2, 1]
		{50, 20, 30, 10, 5},       // Should return 95 → [5, 10, 30, 50]
	}

	for _, nums := range testCases {
		fmt.Printf("\nInput: %v\n", nums)

		bestSum, chosenNumbers := bestSumWithRestrictions(nums)
		fmt.Printf("Best Sum:      %d Chosen Numbers: %v\n", bestSum, chosenNumbers)
		bestSum, chosenNumbers = bestSumWithRestrictionsMemo(nums)
		fmt.Printf("Best Sum Memo: %d Chosen Numbers: %v\n", bestSum, chosenNumbers)
		bestSum, chosenNumbers = bestSumWithRestrictionsTab(nums)
		fmt.Printf("Best Sum Tab:  %d Chosen Numbers: %v\n", bestSum, chosenNumbers)
	}
}

// Helper function for recursive calculation
func bestSumWithRestrictions(nums []int) (int, []int) {
	// Remove duplicates by using a map
	uniqueNums := make(map[int]bool)
	for _, num := range nums {
		uniqueNums[num] = true
	}

	// Convert map back to slice and sort in descending order
	var sortedNums []int
	for num := range uniqueNums {
		sortedNums = append(sortedNums, num)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sortedNums)))

	// Recursive helper function
	var dfs func(currentSum int) (int, []int)
	dfs = func(currentSum int) (int, []int) {
		bestSum := currentSum
		var bestCombination []int

		for _, num := range sortedNums {
			if num > currentSum {
				newSum, newCombination := dfs(currentSum + num)
				if newSum > bestSum {
					bestSum = newSum
					bestCombination = append([]int{num}, newCombination...)
				}
			}
		}
		return bestSum, bestCombination
	}

	return dfs(0)
}

// Optimized function using memoization
func bestSumWithRestrictionsMemo(nums []int) (int, []int) {
	// Remove duplicates using a map and sort in descending order
	uniqueNums := make(map[int]bool)
	for _, num := range nums {
		uniqueNums[num] = true
	}

	var sortedNums []int
	for num := range uniqueNums {
		sortedNums = append(sortedNums, num)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sortedNums)))

	// Memoization map to store results for sums
	memo := make(map[int](struct {
		sum         int
		combination []int
	}))

	// Recursive helper function with memoization
	var dfs func(int) (int, []int)
	dfs = func(currentSum int) (int, []int) {
		// Return memoized result if available
		if result, found := memo[currentSum]; found {
			return result.sum, result.combination
		}

		bestSum := currentSum
		var bestCombination []int

		// Try all numbers and recursively explore further
		for _, num := range sortedNums {
			if num > currentSum {
				newSum, newCombination := dfs(currentSum + num)
				if newSum > bestSum {
					bestSum = newSum
					bestCombination = append([]int{num}, newCombination...)
				}
			}
		}

		// Store the result in memo before returning
		memo[currentSum] = struct {
			sum         int
			combination []int
		}{bestSum, bestCombination}

		return bestSum, bestCombination
	}

	return dfs(0)
}

// Space-optimized version for sum maximization with constraints
func bestSumWithRestrictionsTab(nums []int) (int, []int) {
	// Remove duplicates and sort ascending
	uniqueNums := make(map[int]bool)
	for _, num := range nums {
		uniqueNums[num] = true
	}

	var sortedNums []int
	for num := range uniqueNums {
		sortedNums = append(sortedNums, num)
	}
	sort.Ints(sortedNums) // Ascending order for proper constraint handling

	// Track current valid sums and their combinations
	currentSums := []int{0}
	bestCombination := map[int][]int{
		0: {},
	}

	bestSum := 0

	// Dynamic state expansion
	for _, num := range sortedNums {
		newSums := []int{}
		for _, sum := range currentSums {
			if num > sum {
				newSum := sum + num
				if _, exists := bestCombination[newSum]; !exists || newSum > bestSum {
					newSums = append(newSums, newSum)
					bestCombination[newSum] = append([]int{}, bestCombination[sum]...)
					bestCombination[newSum] = append(bestCombination[newSum], num)
					if newSum > bestSum {
						bestSum = newSum
					}
				}
			}
		}
		currentSums = append(currentSums, newSums...)
	}

	return bestSum, bestCombination[bestSum]
}
