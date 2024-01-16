package main

import (
	"fmt"
)

func binaryGap(n int) int {
	binaryN := fmt.Sprintf("%b", n)
	// fmt.Println(binaryN)

	var seeOneFlag bool
	var zerosCounter int
	var zerosResult int
	for _, v := range binaryN {
		if v == '1' {
			if seeOneFlag {
				zerosResult = max(zerosCounter, zerosResult)
				zerosCounter = 0
			} else {
				seeOneFlag = true
			}
			continue
		}
		zerosCounter++
	}

	return zerosResult
}
