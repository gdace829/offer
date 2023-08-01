package everyday

// 1144 递减元素数组锯齿状 分类讨论
func movesToMakeZigzag(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	var min func(int, int) int
	var max func(int, int) int
	max = func(i1 int, i2 int) int {
		if i1 > i2 {
			return i1
		}
		return i2
	}
	min = func(i1, i2 int) int {
		if i1 > i2 {
			return i2
		}
		return i1
	}
	n := len(nums)
	a := max(nums[0]-nums[1]+1, 0)
	b := 0
	for i := 1; i < n-1; i++ {
		if i%2 == 0 {
			a += max(nums[i]-min(nums[i+1], nums[i-1]), 0)
		} else {
			b += max(nums[i]-min(nums[i+1], nums[i-1]), 0)
		}
	}
	if len(nums)%2 == 1 {
		a += max(nums[len(nums)-1]-nums[len(nums)-2]+1, 0)
	} else {
		b += max(nums[len(nums)-1]-nums[len(nums)-2]+1, 0)
	}
	return min(a, b)
}
