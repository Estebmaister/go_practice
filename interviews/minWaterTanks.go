package interviews

import "strings"

const houseChar string = "H"
const emptyPlotChar string = "-"

// read a string of houses and empty plots such as "H--H-", like a street,
// and returns the min number of water tanks necessary to cover all the houses,
// one tank can cover two immediately near houses, "H-H-" -> "HTH-"
// if there is no way to cover all the houses return -1 "HH",
// if there no houses in the string returns 0 "---".
func minWaterTanks(houses string) int {

	// Edge cases:
	// Case 1. Empty string
	// Case 2. One character string
	// Case 3. House blocked at the beginning "HH--..."
	// Case 4. House blocked at the end or trapped "...HH" or "...HHH..."

	var tanksCounter int
	var noHousesSeen bool = true
	var streetLength int = len(houses)

	if streetLength == 0 {
		// Case #1 covered
		return 0
	}

	if streetLength == 1 {
		// Case #2 covered
		if houses == houseChar {
			return -1
		} else {
			return 0
		}
	}

	for idx := 0; idx < streetLength; idx++ {
		// Take action when an empty plot is seen, decision to put a tank or not
		if string(houses[idx]) == emptyPlotChar { // "-"

			// Check if there is an empty plot before "--"
			if idx != 0 && string(houses[idx-1]) == emptyPlotChar {

				// With empty plot before, check for end of the string
				if idx == streetLength-1 {
					// No tank required
				} else if string(houses[idx+1]) == houseChar { // House next "--H"
					tanksCounter++
				}
				// No house next, no end of the string, no tank required

				// Check for a house next at the beginning "-H..."
			} else if idx == 0 && string(houses[idx+1]) == houseChar {
				// Empty plot at the beginning and house next
				tanksCounter++

				// Check for a house before ""...H-..."
			} else if idx != 0 && string(houses[idx-1]) == houseChar {

				// Check for two houses before "HH-..."
				if idx > 1 && string(houses[idx-2]) == houseChar {
					tanksCounter++
					// Check for beginning of the street "H-..."
				} else if idx == 1 {
					tanksCounter++
					// Check for house after and before "...H-H..."
				} else if idx != streetLength-1 && string(houses[idx+1]) == houseChar {

					tanksCounter++

					// Check for empty plots two spaces before "-H-H..." remove a tank
					// Check for house three spaces before "H-H-H..." add a tank
					tanksCounter += checkBeforeRecursive(idx, 2, houses)
				}

			}
		} else { // "H"
			// Flag that houses have been seen
			noHousesSeen = false

			// Case #3 covered
			// House trapped at the beginning "HH..."
			if idx == 0 && string(houses[idx+1]) == houseChar {
				return -1
			}

			// Case #4 covered
			// Check if there is a house before "...HH..."
			if idx > 0 && string(houses[idx-1]) == houseChar {
				if idx == streetLength-1 || string(houses[idx+1]) == houseChar {
					// House trapped between two houses or at the end "HHH" "...HH"
					return -1
				}
			}

		}
	}

	if noHousesSeen {
		return 0
	} else if tanksCounter == 0 {
		return -1 // Case already covered in loops combined with noHouseSeen
	}

	return tanksCounter
}

/*
	Extracted from this pattern:

// Check for empty plots two spaces before "-H-H..."

	if idx > 1 && string(houses[idx-2]) == emptyPlotChar {
		tanksCounter--
		// Check for house three spaces before "H-H-H..."
		if idx > 2 && string(houses[idx-3]) == houseChar {
			tanksCounter++
			// Check for empty plots four spaces before "-H-H-H..."
			if idx > 3 && string(houses[idx-4]) == emptyPlotChar {
				tanksCounter--
			}
		}
	} else if idx > 1 && string(houses[idx-2]) == houseChar {
		// Pass
	}
*/
func checkBeforeRecursive(idx int, checkBack int, houses string) int {
	var tanksCounter int
	if idx < checkBack {
		return tanksCounter
	}

	// Check for empty plots two spaces before "-H-H..."
	if idx >= checkBack && string(houses[idx-checkBack]) == emptyPlotChar {
		tanksCounter--
		checkBack++
		// Check for house three spaces before "H-H-H..."
		if idx >= checkBack && string(houses[idx-checkBack]) == houseChar {
			if idx%5 == 0 && checkBack%5 == 0 {
				tanksCounter--
			}
			tanksCounter++
			checkBack++
		}
	} else if idx >= checkBack && string(houses[idx-checkBack]) == houseChar {
		return tanksCounter
	}

	return tanksCounter + checkBeforeRecursive(idx, checkBack, houses)
}

// A different approach with simple logic, only two responses:
// returns -1 when there is no way to put a tank or achieve the goal
// returns n int where n is the minimum number of tanks required
func numberOfWaterTanks(houses string) int {
	if len(houses) == 0 || len(houses) == 1 {
		return -1
	}

	var list []int = make([]int, 0, len(houses))
	var countHyphen int
	var countH int
	var count int

	// Append indexes of houses to the array
	for i := 0; i < len(houses); i++ {
		if string(houses[i]) == emptyPlotChar {
			countHyphen++
		} else if string(houses[i]) == houseChar {
			countH++
			list = append(list, i)
		}
	}

	// if less hyphens than half of the houses return -1
	if countHyphen < countH/2 {
		return -1
	}

	// Add a tank for every two near house with a space between them "...H-H..."
	for j := 1; j < len(list); j++ {
		n := list[j] - list[j-1]
		if n == 2 {
			count++
			list = append(list[:j-1], list[j+1:]...)
			j--
		}
	}

	for k := 0; k < len(list); k++ {
		if list[k] == 0 {
			if string(houses[list[k]+1]) == emptyPlotChar {
				count++
			} else {
				return -1
			}
		} else if list[k] == len(houses)-1 {
			if string(houses[list[k]-1]) == emptyPlotChar {
				count++
			} else {
				return -1
			}
		} else if string(houses[list[k]-1]) == emptyPlotChar ||
			string(houses[list[k]+1]) == emptyPlotChar {
			count++
		} else {
			return -1
		}
	}

	if count == 0 {
		return -1
	}
	return count
}

const TankChar string = "T"
const coveredHouseChar string = "X"

// A different approach with simplest logic, only two responses:
// returns -1 when there is no way to put a tank or achieve the goal
// returns n int where n is the minimum number of tanks required
func stringsWaterTanks(str string) int {
	sb := strings.Split(str, "")

	// Putting a water tank between adjacent houses "...H-H".
	for i := 1; i < len(str)-1; i++ {
		if sb[i] == emptyPlotChar && sb[i-1] == houseChar && sb[i+1] == houseChar {
			sb[i] = TankChar
			sb[i-1] = coveredHouseChar
			sb[i+1] = coveredHouseChar
		}
	}

	// Putting a tank near pending houses.
	for i := 0; i < len(str); i++ {
		if sb[i] == houseChar {
			if i-1 >= 0 && sb[i-1] != TankChar && sb[i-1] == emptyPlotChar {
				sb[i-1] = TankChar
				sb[i] = coveredHouseChar
			} else if i+1 <= len(sb)-1 && sb[i+1] != TankChar && sb[i+1] == emptyPlotChar {
				sb[i+1] = TankChar
				sb[i] = coveredHouseChar
			}
		}
	}

	// Last check to count tanks, if any house was left it returns -1
	count := 0
	for i := 0; i < len(sb); i++ {
		if sb[i] == TankChar {
			count++
		} else if sb[i] == houseChar {
			return -1
		}
	}
	return count
}
