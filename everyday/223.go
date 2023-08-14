package everyday

// 循环码排列 1238
func circularPermutation(n int, start int) []int {
	nums := make([]int, n<<1)
	num := start ^ 0
	for i := 0; i < len(nums); i++ {
		nums[i] = i ^ (i >> 1) ^ num
	}
	return nums
}
