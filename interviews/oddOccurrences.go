package interviews

// Check for an integer with no pair in an odd-length array
func oddOccurrences(oddArray []int) int {
	var occurrences = make(map[int]int)

	for _, val := range oddArray {
		seenValue, ok := occurrences[val]

		if !ok {
			occurrences[val] = seenValue + 1
		}

		if seenValue == 1 {
			delete(occurrences, val)
		}
	}

	var foundValue int = 0
	for key := range occurrences {
		foundValue = key
	}

	return foundValue
}
