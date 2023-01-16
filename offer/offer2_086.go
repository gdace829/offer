package offer

func partition(s string) [][]string {
	res := [][]string{}
	cur := []string{}
	var ispartition func(string) bool
	ispartition = func(s string) bool {
		i := 0
		j := len(s) - 1
		for i < j {
			if s[i] != s[j] {
				return false
			}
			i++
			j--
		}
		return true
	}
	var dfs func(int)
	dfs = func(idx int) {
		if idx == len(s) {
			res = append(res, append([]string{}, cur...))
			return
		}
		for i := idx + 1; i <= len(s); i++ {
			if ispartition(s[idx:i]) {
				cur = append(cur, s[idx:i])
				dfs(i)
				cur = cur[:len(cur)-1]
			}
		}

	}
	dfs(0)
	return res
}
