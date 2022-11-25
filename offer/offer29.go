package offer

var direct1s [][]int = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	ans := make([]int, 0)
	x := 0
	y := 0
	right := len(matrix[0])
	down := len(matrix)
	up := 1
	left := 0
	count := 0
	direct := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			ans = append(ans, matrix[x][y])
			if count == len(matrix)*len(matrix[0]) {
				return ans
			}
			if direct == 0 && y == right-1 {
				direct = 1
				right--
			} else if direct == 1 && x == down-1 {
				direct = 2
				down--
			} else if direct == 2 && y == left {
				direct = 3
				left++
			} else if direct == 3 && x == up {
				direct = 0
				up++
			}
			x += direct1s[direct][0]
			y += direct1s[direct][1]

		}
	}
	return ans
}
