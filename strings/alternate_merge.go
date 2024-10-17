package strings

import (
	"math"
	"strings"
)

// Merge two words letter by letter, starting with word1
// if one word is longer than the other, append the remaining letters from the longer word
// ex: word1 = "abc", word2 = "pqr" -> "apbqcr"
// ex: word1 = "ab", word2 = "pqrs" -> "apbqrs"
// ex: word1 = "abcd", word2 = "pq" -> "apbqcd"
// Time complexity: O(n+m), where n and m are the lengths of word1 and word2, respectively.
// Space complexity: O(n+m), where n and m are the lengths of word1 and word2, respectively.
// where n and m are the lengths of word1 and word2, respectively.
func MergeAlternately(word1 string, word2 string) string {
	finalWord := strings.Builder{}
	maxLen := int(math.Max(float64(len(word1)), float64(len(word2))))

	for i := 0; i < maxLen; i++ {
		if len(word1) > i {
			finalWord.WriteString(string(word1[i]))
		}
		if len(word2) > i {
			finalWord.WriteString(string(word2[i]))
		}
	}
	return finalWord.String()
}
