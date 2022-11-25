package offer

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	cut := make([]int, 26)
	for _, v := range s {
		cut[v-'a']--
	}
	for _, v := range t {
		cut[v-'a']++
		if cut[v-'a'] > 0 {
			return false
		}
	}
	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			return true
		}
	}
	return false
}
