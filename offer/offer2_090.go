package offer

func rob1(nums []int) int {
	var max func(int, int) int
	max = func(i1, i2 int) int {
		if i1 > i2 {
			return i1
		}
		return i2
	}
	n := len(nums)
	a, b := 0, nums[0]
	for i := 1; i < n-1; i++ {
		a, b = b, max(a+nums[i], b)
	}
	res := b
	a, b = 0, nums[n-1]
	for i := n - 2; i > 0; i-- {
		a, b = b, max(a+nums[i], b)
	}
	return max(res, b)
}
