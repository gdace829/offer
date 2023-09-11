package foundation

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	ans := n + 1
	l := 0
	r := 0
	now := 0
	for r < n {
		now += nums[r]
		for now >= target {
			if r-l+1 < ans {
				ans = r - l + 1
			}

			now -= nums[l]
			l++
		}

		r++
	}
	if ans == n+1 {
		return 0
	} else {
		return ans
	}
}
