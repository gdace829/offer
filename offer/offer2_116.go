package offer

func findCircleNum(isConnected [][]int) int {
	var (
		edges [][]int
		vis   []int
		num   int
		dfs   func(int)
	)
	edges = make([][]int, len(isConnected))
	for i := 0; i < len(isConnected); i++ {
		for j := 0; j < len(isConnected); j++ {
			if isConnected[i][j] == 1 {
				edges[i] = append(edges[i], j)
			}
		}
	}
	vis = make([]int, len(isConnected))
	num = 1
	dfs = func(i int) {
		vis[i] = num
		for _, v := range edges[i] {
			if vis[v] == 0 {
				dfs(v)
			}
		}

	}
	for i := 0; i < len(isConnected); i++ {
		if vis[i] == 0 {
			dfs(i)
			num++
		}
	}
	return num - 1

}
