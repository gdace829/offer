package offer

// 对于整体回文串个数中心扩散
func countSubstrings(s string) int {
	var mid func(l, r int)
	res := 0
	n := len(s)
	mid = func(l, r int) {
		for l >= 0 && r < n {
			if s[l] == s[r] {
				res++
			} else {
				break
			}
			l = l - 1
			r = r + 1
		}
	}
	for i := 0; i < n; i++ {
		mid(i, i)
		mid(i, i+1)
	}
	return res
}
