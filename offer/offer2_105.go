package offer

func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	vis := make([][]bool, m)
	for i := 0; i < len(vis); i++ {
		vis[i] = make([]bool, n)
	}
	var dfs func(int, int) int

	dfs = func(i1, i2 int) int {
		if i1 < 0 || i1 >= m || i2 < 0 || i2 >= n {
			return 0
		}
		if grid[i1][i2] == 0 || vis[i1][i2] == true {
			return 0
		}
		vis[i1][i2] = true

		return 1 + dfs(i1-1, i2) + dfs(i1, i2-1) + dfs(i1+1, i2) + dfs(i1, i2+1)
	}
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			n := dfs(i, j)
			if n > res {
				res = n
			}

		}
	}
	return res
}
