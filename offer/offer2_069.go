package offer

func peakIndexInMountainArray(arr []int) int {
	left := 0
	right := len(arr)
	for left < right {
		mid := (left + right) / 2
		if mid > 0 && arr[mid] > arr[mid-1] {
			left = mid
		} else if mid > 0 && arr[mid] < arr[mid-1] {
			right = mid
		}
		if (left+right)/2 == mid {
			return left
		}
	}
	return right
}
