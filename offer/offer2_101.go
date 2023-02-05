package offer

func canPartition(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	dp := make([][]bool, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]bool, sum/2)
	}
	for i := 0; i <= len(nums); i++ {
		dp[i][0] = true
	}
	for i := 0; i <= len(nums); i++ {
		for j := 1; j <= sum/2; j++ {
			if j >= nums[i-1] {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(nums)][sum/2-1]
}
