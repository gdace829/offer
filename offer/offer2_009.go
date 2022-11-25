package offer

func numSubarrayProductLessThanK(nums []int, k int) int {
	l := 0
	r := 0
	n := len(nums)
	cur := 1
	res := 0
	for r < n {
		cur *= nums[r]
		for cur >= k && l <= r {
			cur /= nums[l]
			l++
		}
		res += r - l+1 
        	r++
	}
	return res
}