package str

// canConstruct takes a word an a list of substrings wordBank
// and returns true if the word can be construct with substrings
// on the list
func canConstruct(s string, wordBank []string) bool {
	if s == "" {
		return true
	}

	for _, word := range wordBank {
		if len(s) >= len(word) && s[:len(word)] == word {
			if canConstruct(s[len(word):], wordBank) {
				return true
			}
		}
	}

	return false
}

func canConstructMemo(s string, wordBank []string, memo map[string]bool) bool {
	if val, ok := memo[s]; ok {
		return val
	}

	if s == "" {
		return true
	}

	for _, word := range wordBank {
		if len(s) >= len(word) && s[:len(word)] == word {
			if canConstructMemo(s[len(word):], wordBank, memo) {
				memo[s] = true
				return true
			}
		}
	}

	memo[s] = false
	return false
}

func canConstructTab(s string, wordBank []string) bool {
	targetLen := len(s)
	table := make([]bool, targetLen+1)
	table[0] = true

	for i := 0; i <= targetLen; i++ {
		if table[i] {
			for _, word := range wordBank {
				if targetLen >= i+len(word) && s[i:i+len(word)] == word {
					table[i+len(word)] = true
				}
			}
		}
	}

	return table[targetLen]
}

// countConstruct takes a word an a list of substrings wordBank
// and returns the number of times the word can be construct with substrings
// on the list
func countConstruct(s string, wordBank []string) int {
	if s == "" {
		return 1
	}

	totalCount := 0
	for _, word := range wordBank {
		if len(s) >= len(word) && s[:len(word)] == word {
			numWays := countConstruct(s[len(word):], wordBank)
			totalCount += numWays
		}
	}

	return totalCount
}

func countConstructMemo(s string, wordBank []string, memo map[string]int) int {
	if val, ok := memo[s]; ok {
		return val
	}

	if s == "" {
		return 1
	}

	totalCount := 0
	for _, word := range wordBank {
		if len(s) >= len(word) && s[:len(word)] == word {
			numWays := countConstructMemo(s[len(word):], wordBank, memo)
			totalCount += numWays
		}
	}

	memo[s] = totalCount
	return totalCount
}

func countConstructTab(s string, wordBank []string) int {
	targetLen := len(s)
	table := make([]int, targetLen+1)
	table[0] = 1

	for i := 0; i <= targetLen; i++ {
		if table[i] > 0 {
			for _, word := range wordBank {
				// 		if targetLen >= i+len(word) &&  s[i:i+len(word)] == word
				// ex. i=2 {apple, ple} if 5 >= 5 &&  s[2:6] == ple
				// 	table[5] += table[3]
				if targetLen >= i+len(word) && s[i:i+len(word)] == word {
					table[i+len(word)] += table[i]
				}
			}
		}
	}

	return table[targetLen]
}

// allConstruct takes a word an a list of substrings wordBank
// and returns all the ways the word can be construct with substrings
// on the list
func allConstruct(s string, wordBank []string) [][]string {
	if s == "" {
		return [][]string{{}}
	}

	var result [][]string
	for _, word := range wordBank {
		if len(s) >= len(word) && s[:len(word)] == word {
			suffixWays := allConstruct(s[len(word):], wordBank)
			for _, way := range suffixWays {
				way = append([]string{word}, way...)
				result = append(result, way)
			}
		}
	}

	return result
}

func allConstructMemo(s string, wordBank []string, memo map[string][][]string) [][]string {
	if val, ok := memo[s]; ok {
		return val
	}
	if s == "" {
		// border condition (empty string always returns [])
		return [][]string{{}}
	}

	var result [][]string
	for _, word := range wordBank {
		if len(s) >= len(word) && s[:len(word)] == word {
			suffixWays := allConstructMemo(s[len(word):], wordBank, memo)
			for _, way := range suffixWays {
				way = append([]string{word}, way...)
				result = append(result, way)
			}
		}
	}

	memo[s] = result
	return result
}

func allConstructTab(s string, wordBank []string) [][]string {
	targetLen := len(s)
	table := make([][][]string, targetLen+1)
	// border condition (empty string always returns [])
	table[0] = [][]string{{}}

	for i := 0; i <= targetLen; i++ {
		if len(table[i]) > 0 {
			for _, word := range wordBank {
				// ex. i=2 {apple, [ap, ple]} if 5 >= 5 &&  s[2:6] == ple
				// 	prevCombinations = table[i] = [["ap"]]
				// 	combination = ["ap","ple"]
				// 	table[5] = [["ap","ple"]]
				if targetLen >= i+len(word) && s[i:i+len(word)] == word {
					for _, combination := range table[i] {
						// Create a new slice to avoid modifying the original reference
						newCombination := append([]string{}, combination...)
						// A second option to create the array could be this one:
						// newCombination := make([]string, len(combination))
						// copy(newCombination, combination)
						newCombination = append(newCombination, word)
						table[i+len(word)] = append(table[i+len(word)], newCombination)
					}
				}
			}
		}
	}

	return table[targetLen]
}
