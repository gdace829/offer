package foundation

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	max := 0
	for left < right {
		if height[left] < height[right] {
			if height[left]*(right-left) > max {
				max = height[left] * (right - left)
			}
			left++
		} else {
			if height[right]*(right-left) > max {
				max = height[right] * (right - left)
			}
			right--
		}
	}
	return max
}
