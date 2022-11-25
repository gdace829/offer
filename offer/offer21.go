package offer

func exchange(nums []int) []int {
	j := len(nums) - 1
	t := 0
	for i := 0; i < len(nums) && i <= j; i++ {
		if nums[i]%2 == 1 {
			continue
		} else {
			t = nums[i]
			nums[i] = nums[j]
			nums[j] = t
			j--
		}
	}
	return nums
}
