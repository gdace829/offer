package offer

// 比map更好的做法利用索引
func FindRepeatNumber(nums []int) int {
	var t int
	for i := 0; i < len(nums); {
		if i == nums[i] {
			i++
			continue
		} else {
			if nums[nums[i]] == nums[i] {
				return nums[i]
			} else {
				t = nums[nums[i]]
				nums[nums[i]] = nums[i]
				nums[i] = t
			}

		}

	}
	return -1
}
