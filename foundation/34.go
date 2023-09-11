package foundation

func findSmaller(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
func findBigger(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left := findSmaller(nums, target)
	right := findBigger(nums, target)
	if nums[left] != target {
		return []int{-1, -1}
	}
	if nums[right] == target {
		return []int{left, right}
	}
	return []int{left, right - 1}
}
