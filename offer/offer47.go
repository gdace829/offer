package offer

import "fmt"

func maxValue1(grid [][]int) int {
	vis := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		vis[i] = make([]bool, len(grid[0]))
	}
	var max int
	var dfs func(count, i int, j int)
	dfs = func(count, i, j int) {
		if i > len(grid[0])-1 || j > len(grid)-1 || i < 0 || j < 0 || vis[i][j] == true {
			return
		}
		count += grid[i][j]
		if i == len(grid)-1 && j == len(grid[0]) && count > max {
			max = count
			fmt.Println(count)
		}
		vis[i][j] = true
		dfs(count, i+1, j)
		dfs(count, i-1, j)
		dfs(count, i, j+1)
		dfs(count, i, j-1)
		vis[i][j] = false

	}
	dfs(0, 0, 0)
	return max

}

// 当dfs路线限制时可以dp做
func maxValue(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[0]))
	}
	dp[0][0] = grid[0][0]
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
				continue
			}
			if i == 0 {
				dp[i][j] = grid[i][j] + dp[i][j-1]
				continue
			}
			if j == 0 {
				dp[i][j] = grid[i][j] + dp[i-1][j]
				continue
			}
			if dp[i][j-1] > dp[i-1][j] {
				dp[i][j] = grid[i][j] + dp[i][j-1]
			} else {
				dp[i][j] = grid[i][j] + dp[i-1][j]
			}
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}
