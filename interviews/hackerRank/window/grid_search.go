package window

import "fmt"

const (
	YES = "YES"
	NO  = "NO"
)

func gridPatternSearch(G []string, P []string) string {
	// G grid, P pattern
	gLen, pLen := len(G), len(P)
	if gLen < 2 || pLen < 1 {
		return NO
	}
	gColumnLen, pColumnLen := len(G[0]), len(P[0])

	window := make([]string, pLen)
	wCounter := 0 // window counter to track window rows

	for _, column := range G {
		fmt.Println(column)

		for i := 0; i <= gColumnLen-pColumnLen; i++ {

			window[wCounter] = column[i : i+pColumnLen]
			fmt.Println(window[wCounter], P[wCounter])

			if window[wCounter] == P[wCounter] {
				if wCounter == pLen-1 {
					return YES
				}
				wCounter++
				break
			}

			// After one pass of a complete column
			// withot finding the next element of the
			// pattern, reset the window counter
			if i == gLen-pLen-1 && wCounter != 0 {
				wCounter = 0
				// repeat column
				i = 0
			}
		}
	}

	return NO
}
