package offer

import (
	"math"
)

func updateMatrix(mat [][]int) [][]int {
	var min func(int, int) int
	min = func(i1, i2 int) int {
		if i1 > i2 {
			return i2
		}
		return i1
	}
	m, n := len(mat), len(mat[0])
	vis := make([][]bool, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n {
			return math.MaxInt
		}
		if vis[i][j] == true {
			return math.MaxInt
		}
		vis[i][j] = true
		if mat[i][j] == 0 {
			return 0
		}
		res := min(dfs(i-1, j), min(dfs(i, j-1), min(dfs(i+1, j), dfs(i, j+1)))) + 1

		vis[i][j] = false

		return res
	}
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[i][j] = dfs(i, j)
		}
	}
	return res
}
