package offer
func minCostClimbingStairs(cost []int) int {
	var min func(int, int) int
	min = func(i1, i2 int) int {
		if i1 < i2 {
			return i1
		}
		return i2
	}
	dp := make([]int, len(cost))
	dp[0] = cost[0]
	dp[1] = cost[1]
	n := len(dp)
	for i := 2; i < n; i++ {
		dp[i] =  min(dp[i-1]+cost[i], dp[i-2]+cost[i])
	}
	if n>=2 {
		return min(dp[n-1],dp[n-2])
	}else{
		return dp[n-1]
	}
	
}
