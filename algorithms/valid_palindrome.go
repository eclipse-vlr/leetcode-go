package algorithms

import "unicode"

func isPalindrome(s string) bool {
	runes := []rune(s)

	left := 0
	right := len(runes) - 1

	for left < right {
		l := runes[left]
		r := runes[right]

		if !unicode.IsLetter(l) && !unicode.IsDigit(l) {
			left++
			continue
		}

		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			right--
			continue
		}

		if unicode.ToLower(l) != unicode.ToLower(r) {
			return false
		}

		left++
		right--
	}

	return true
}
