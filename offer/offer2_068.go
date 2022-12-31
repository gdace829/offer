package offer

func searchInsert(nums []int, target int) int {
	//二分查找顺序数组
	left := 0
	right := len(nums)

	for left < right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}

	}
	return left
}
