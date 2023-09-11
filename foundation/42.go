package foundation

func trap(height []int) int {
	left := 0
	right := len(height) - 1
	leftmax := 0
	rightmax := 0
	sum := 0
	for left <= right {
		if height[left] > leftmax {
			leftmax = height[left]
		}
		if height[right] > rightmax {
			rightmax = height[right]
		}
		if leftmax < rightmax {
			sum += leftmax - height[left]
			left++
		} else {
			sum += rightmax - height[right]
			right--
		}

	}
	return sum
}
