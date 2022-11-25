package offer

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	cut := make([]int, 26)
	l := 0
	r := 0
	for i := 0; i < len(s1); i++ {
		cut[s1[i]-'a']--
	} //记录每个字母出现次数
	for r < len(s2) {
		cut[s2[r]-'a']++
		for cut[s2[r]-'a'] > 0 {
			cut[s2[l]-'a']--
			l++
		} //对于使得当前字母超出情况，左移left，减少count
		if r-l+1 == len(s1) {
			return true
		}
		r++
	}
	return false
}
