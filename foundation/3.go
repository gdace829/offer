package foundation

func lengthOfLongestSubstring(s string) int {
	ans := 0
	l := 0
	r := 0
	n := len(s)
	map1 := make(map[byte]int)
	for r < n {
		map1[s[r]-'a']++
		for l < r && map1[s[r]-'a'] > 1 {
			map1[s[l]-'a']--
			l++
		}
		if r-l+1 > ans {
			ans = r - l + 1
		}
		r++
	}
	return ans
}
