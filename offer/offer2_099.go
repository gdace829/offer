package offer

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[0]))
	}
	var min func(int, int) int
	min = func(i1, i2 int) int {
		if i1 > i2 {
			return i2
		}
		return i1
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if j > 0 && i > 0 {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			} else if i > 0 {
				dp[i][j] = dp[i-1][0] + grid[i][j]
			} else if j > 0 {
				dp[i][j] = dp[0][j-1] + grid[i][j]
			}
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}
