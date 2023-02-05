package offer

func isBipartite(graph [][]int) bool {
	color := make([]int, len(graph))

	var dfs func(idx int, val int) bool
	dfs = func(idx, val int) bool {
		if color[idx] == val {
			return true
		} else if color[idx] != 0 {
			return false
		}
		color[idx] = val
		for _, v := range graph[idx] {
			if !dfs(v, -val) {
				return false
			}
		}
		return true
	}
	for k, _ := range graph {
		if color[k] == 0 && !dfs(k, 1) {
			return false
		}
	}
	return true
}
