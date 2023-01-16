package offer

func rob(nums []int) int {
	var max func(int, int) int
	max = func(i1, i2 int) int {
		if i1 > i2 {
			return i1
		}
		return i2
	}

	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[1], nums[0])
	}

	max1 := max(nums[1], nums[0])
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = nums[1]

	for i := 2; i < n; i++ {
		if i >= 3 {
			dp[i] = nums[i] + max(dp[i-2], dp[i-3])
		} else {
			dp[i] = nums[i] + dp[i-2]
		}

		if dp[i] > max1 {
			max1 = dp[i]
		}
	}
	return max1
}
