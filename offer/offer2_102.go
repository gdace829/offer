package offer

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	var abs func(int) int
	abs = func(i int) int {
		if i > 0 {
			return i
		}
		return -i
	}
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	p := (sum + target) / 2
	q := (sum - target) / 2
	if (sum+target)%2 == 1 {
		return 0
	}
	if sum-abs(target) < 0 {
		return 0
	}
	n := 0
	if p > q {
		n = q
	} else {
		n = p
	}
	dp := make([][]int, len(nums)+1)
	for i := 0; i <= len(nums); i++ {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 1
	for i := 1; i <= len(nums); i++ {
		for j := 0; j <= n; j++ {
			if j >= nums[i-1] {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(nums)][n]
}
