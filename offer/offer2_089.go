package offer

func rob(nums []int) int {
	var max func(int, int) int
	max = func(i1, i2 int) int {
		if i1 > i2 {
			return i1
		}
		return i2
	}
	a := 0
	b := nums[0]
	for i := 1; i < len(nums); i++ {
		a, b = b, max(nums[i]+a, b)
	}
	return b
}
