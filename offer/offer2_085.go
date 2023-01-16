package offer

func generateParenthesis(n int) []string {
	res := []string{}
	cur := ""
	var dfs func()
	var i, j int
	dfs = func() {
		if len(cur) == 2*n {
			res = append(res, ""+cur)
			return
		}
		if i < n {
			cur = cur + "("
			i++
			dfs()
			cur = cur[:len(cur)-1]
			i--
		}
		if j < i {
			cur = cur + ")"
			j++
			dfs()
			cur = cur[:len(cur)-1]
			j--
		}

	}
	dfs()
	return res
}
