package offer

func allPathsSourceTarget(graph [][]int) [][]int {
	res := make([][]int, 0)
	path := make([]int, len(graph))
	path = append(path, 0)
	var dfs func(int)
	dfs = func(i int) {
		if i == len(graph)-1 {
			res = append(res, append([]int{}, path...))
		}
		for _, v := range graph[i] {
			path = append(path, v)
			dfs(v)
			path = path[:len(path)-1]
		}

	}
	dfs(0)
	return res
}
