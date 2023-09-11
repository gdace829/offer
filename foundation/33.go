package foundation

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := (left + right) / 2
		if target > nums[len(nums)-1] {
			if nums[mid] >= target {
				right = mid
			} else if nums[mid] > nums[len(nums)-1] {
				left = mid + 1
			} else {
				right = mid
			}
		} else {
			if nums[mid] < target {
				left = mid + 1
			} else if nums[mid] > nums[len(nums)-1] {
				left = mid + 1
			} else {
				right = mid
			}
		}
	}
	if nums[left] != target {
		return -1
	}
	return left
}
