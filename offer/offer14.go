package offer

func cuttingRope(n int) int {
	if n == 1 {
		return 1
	} else if n == 0 {
		return 0
	}
	var dp []int
	dp = make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 0; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] = max(max((i-j)*j%(1e9+7), j*dp[i-j]%(1e9+7)), dp[i]) % (1e9 + 7)
		}
	}
	return dp[n]
}
func max(i int, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
