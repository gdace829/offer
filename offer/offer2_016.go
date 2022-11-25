package offer

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	cut := make([]int, 300)
	r := 0
	l := 0
	n := len(s)
	res := 0
	for r < n {
		cut[s[r]-'a']++
		for cut[s[r]-'a'] > 1 {
			cut[s[l]-'a']--
			l++
		}
		if r-l+1 > res {
			res = r - l + 1
		}
		r++
	}
	return res
}
