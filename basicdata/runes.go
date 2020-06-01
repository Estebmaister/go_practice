package basicdata

import "strings"

// SwapRune change the case
func SwapRune(r rune) rune {
	switch {
	case 'a' <= r && r <= 'z':
		return r - 'a' + 'A'
	case 'A' <= r && r <= 'Z':
		return r - 'A' + 'a'
	default:
		return r
	}
}

// SwapCase change the case
func SwapCase(str string) string {
	return strings.Map(SwapRune, str)
}

/* Rune literals are just 32-bit integer
values (however they're untyped constants,
so their type can change). They represent
unicode codepoints. For example, the rune
literal 'a' is actually the number 97. */
