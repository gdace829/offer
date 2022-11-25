package offer

func validPalindrome(s string) bool {
	l := 0
	r := len(s) - 1
	for l < r {
		if s[l] != s[r] {
			return isPalindrome1(s[l+1:r+1]) || isPalindrome1(s[l:r])
		}
		l++
		r--
	}
	return true
}
func isPalindrome1(s string) bool {
	l := 0
	r := len(s) - 1
	for l < r && s[l] == s[r] {
		l++
		r--
	}
	return l >= r
}
