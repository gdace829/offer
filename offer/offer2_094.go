package offer

func minCut(s string) int {
	var ishui func(int, int) bool
	ishui = func(i1, i2 int) bool {
		for i1 < i2 {
			if s[i1] != s[i2] {
				return false
			}
			i1++
			i2--
		}
		return true
	}
	var min func(int, int) int
	min = func(i1, i2 int) int {
		if i1 > i2 {
			return i2
		}
		return i1
	}
	dp := make([]int, len(s))
	dp[0] = 0
	for i := 1; i < len(s); i++ {
		dp[i] = dp[i-1] + 1
		for j := 0; j < i; j++ {
			if ishui(j, i) {
				if j == 0 {
					dp[i] = 0
					break
				} else {
					dp[i] = min(dp[i], dp[j-1]+1)
				}
			}

		}
	}
	return dp[len(s)-1]
}
