package foundation

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	ans := 0
	n := len(nums)
	l := 0
	r := 0
	now := 1
	for r < n {
		now *= nums[r]
		for l < r && now >= k {
			now /= nums[l]
			l++

		}
		if now < k {
			ans += r - l + 1
		}
		r++
	}
	return ans
}
