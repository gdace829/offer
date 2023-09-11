package foundation

func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] <= nums[len(nums)-1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
