package leetcode

import "strings"

/*
race a car
   ^  ^
*/

func isPalindrome(s string) bool {
	lowerS := strings.ToLower(s)
	sliceByte := []byte(lowerS)

	l, r := 0, len(sliceByte)-1

	for l < r {
		if !isAlphaNumeric(sliceByte[l]) {
			l++
			continue
		}

		if !isAlphaNumeric(sliceByte[r]) {
			r--
			continue
		}

		if sliceByte[l] != sliceByte[r] {
			return false
		}

		l++
		r--
	}

	return true
}

func isAlphaNumeric(b byte) bool {
	return b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z' || b >= '0' && b <= '9'
}
