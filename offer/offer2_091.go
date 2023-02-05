package offer

func minCost(costs [][]int) int {
	dp := make([]int, 3)
	dp[0] = costs[0][0]
	dp[1] = costs[0][1]
	dp[2] = costs[0][2]
	var min func(int, int) int
	min = func(i1, i2 int) int {
		if i1 > i2 {
			return i2
		}
		return i1
	}
	n := len(costs)
	for i := 1; i < n; i++ {
		tmp := make([]int, 3)
		tmp[0] = min(dp[1], dp[2]) + costs[i][0]
		tmp[1] = min(dp[0], dp[2]) + costs[i][1]
		tmp[2] = min(dp[1], dp[0]) + costs[i][2]
		dp = tmp
	}
	return min(dp[0], min(dp[1], dp[2]))
}
