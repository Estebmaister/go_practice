package window

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
	pStart := 0

	for col := 0; col < gLen; col++ {
		// println("col", col, G[col])
		if pStart > gColumnLen-pColumnLen {
			pStart = 0
			continue
		}
		for i := pStart; i <= gColumnLen-pColumnLen; i++ {

			window[wCounter] = G[col][i : i+pColumnLen]
			// println(i, window[wCounter], P[wCounter])

			if window[wCounter] == P[wCounter] {
				if wCounter == pLen-1 {
					return YES
				}
				wCounter++
				pStart = i
				break
			} else if wCounter != 0 {
				col = col - 1 - wCounter
				wCounter = 0
				pStart++
				break
			}
			if i == gLen-pLen {
				// reset pattern start for next iteration
				pStart = 0
			}
		}
	}

	return NO
}
