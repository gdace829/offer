package offer

import "strings"

func isPalindrome(s string) bool {
	l := 0
	r := len(s) - 1
	s = strings.ToLower(s)
	for l < r {
		for l < r && !((s[l] >= 'a' && s[l] <= 'z') || (s[l] >= '0' && s[l] <= '9')) {
			l++
		}
		for l < r && !((s[r] >= 'a' && s[r] <= 'z') || (s[r] >= '0' && s[r] <= '9')) {
			r--
		}
		if l == r {
			return true
		}

		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
