package offer

func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	if len(p) > len(s) {
		return res
	}
	cut := make([]int, 26)

	for _, v := range p {
		cut[v-'a']--
	}
	l := 0
	r := 0
	n := len(s)
	for r < n {
		cut[s[r]-'a']++
		for cut[s[r]-'a'] > 0 {
			cut[s[l]-'a']--
			l++
		}
		if r-l+1 == len(p) {
			res = append(res, l)
		}
		r++

	}
	return res
}
