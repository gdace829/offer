package offer

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	var min func(int, int) int
	min = func(i1, i2 int) int {
		if i1 > i2 {
			return i2
		}
		return i1
	}
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}
	for _, v := range coins {
		for k, _ := range dp {
			if k >= v {
				dp[k] = min(dp[k], dp[k-v]+1)
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
