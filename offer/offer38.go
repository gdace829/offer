package offer

func permutation(s string) []string {
	var dfs func(x int)
	strs := make([]string, 0)
	res := []byte(s)
	dfs = func(x int) {
		if x == len(res)-1 {
			strs = append(strs, string(res))
		}
		dict := make(map[byte]bool)
		for i := x; i < len(res); i++ {
			if !dict[res[i]] {
				res[x], res[i] = res[i], res[x]
				dict[res[x]] = true
				dfs(x + 1)
				res[x], res[i] = res[i], res[x]
			}
		}
	}
	dfs(0)
	return strs
}
