package offer

func maximalRectangle(matrix []string) int {
	if len(matrix) == 0 {
		return 0
	}
	heights := make([]int, len(matrix[0]))
	result := 0
	for _, s := range matrix {
		for i, c := range s[0:len(s)] {
			if c == '0' {
				heights[i] = 0
			} else {
				heights[i]++
			}
		}
		result = max(result, largestRectangleArea(heights))
	}
	return result
}
