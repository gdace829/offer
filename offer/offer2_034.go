package offer

func isAlienSorted(words []string, order string) bool {
	cut := [26]int{}
	var invaild func(string, string) bool
	for index, v := range order {
		cut[v-'a'] = index
	}
	invaild = func(s1, s2 string) bool {
		for i := 0; i < len(s1) && i < len(s2); i++ {
			if cut[s1[i]-'a'] > cut[s2[i]-'a'] {
				return false
			}
			if cut[s1[i]-'a'] < cut[s2[i]-'a'] {
				return true
			}

		}
		if len(s2) < len(s1) {
			return false
		}

		return true
	}
	for j := 1; j < len(words); j++ {
		if !invaild(words[j-1], words[j]) {
			return false
		}
	}
	return true
}
