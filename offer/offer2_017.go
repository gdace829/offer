package offer

import "math"

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	cut := make([]int, 'z'-'A'+1)
	for _, v := range t {
		cut[v-'A']--
	}
	r := 0
	l := 0
	n := len(s)
	nres := math.MaxInt32
	rres := -1
	lres := -1
	for r < n {
		cut[s[r]-'A']++
		for l <= r && cut[s[l]-'A'] > 0 {
			cut[s[l]-'A']--
			l++
		}
		if check(cut) && r-l+1 < nres {

			rres = r
			lres = l
			nres = r - l + 1
		}
		r++
	}
	if rres == -1 {
		return ""
	}
	return s[lres : rres+1]

}
func check(nums []int) bool {
	for _, v := range nums {
		if v < 0 {
			return false
		}
	}
	return true
}
