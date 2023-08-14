package offer

func longestIncreasingPath(matrix [][]int) int {
	var (
		vis [][]int
		dfs func(int, int) int
	)
	res := 0
	m, n := len(matrix), len(matrix[0])
	vis = make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		vis[i] = make([]int, len(matrix[0]))
	}
	dfs = func(i, j int) int {
		if vis[i][j] != 0 {
			return vis[i][j]
		}
		max := 0
		if i >= 0 && j >= 0 && i < m && j+1 < n && matrix[i][j+1] > matrix[i][j] {

			if a := dfs(i, j+1); a > max {
				max = a
			}

		}
		if i >= 0 && j >= 0 && i+1 < m && j < n && matrix[i+1][j] > matrix[i][j] {

			if a := dfs(i+1, j); a > max {
				max = a
			}

		}
		if i >= 0 && j-1 >= 0 && i < m && j < n && matrix[i][j-1] > matrix[i][j] {

			if a := dfs(i, j-1); a > max {
				max = a
			}

		}
		if i-1 >= 0 && j >= 0 && i < m && j < n && matrix[i-1][j] > matrix[i][j] {

			if a := dfs(i-1, j); a > max {
				max = a
			}

		}
		vis[i][j] = max + 1
		return vis[i][j]
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if vis[i][j] = dfs(i, j); vis[i][j] > res {
				res = vis[i][j]
			}
		}
	}
	return res

}
