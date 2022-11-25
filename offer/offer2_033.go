package offer

func groupAnagrams(strs []string) [][]string {
	hashmap := map[[26]int][]string{}
	cut := [26]int{}
	var count func(string)
	count = func(s string) {
		for i := 0; i < 26; i++ {
			cut[i] = 0
		}
		for _, v := range s {
			cut[v-'a']++
		}

	}
	for _, v := range strs {
		count(v)
		hashmap[cut] = append(hashmap[cut], v)
	}
	res := [][]string{}
	for _, v := range hashmap {
		res = append(res, v)
	}
	return res
}
