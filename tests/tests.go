package tests

import (
	"errors"
	"strings"
	"time"
)

// Split slices s into all substrings separated by sep and
// returns a slice of the substrings between those separators.
func Split(s, sep string) []string {
	time.Sleep(1 * time.Second)
	var result []string
	i := strings.Index(s, sep)
	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):]
		i = strings.Index(s, sep)
	}
	return append(result, s)
}

// SimpleErrFunc is a simple function that returns an error if the input is broken.
func SimpleErrFunc(s string) (string, error) {
	if strings.Contains(s, "broken") {
		return "", errors.New("input is broken")
	}

	return strings.Replace(s, "input", "result", 1), nil
}
