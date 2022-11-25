package offer

func FindNumberIn2DArray(matrix [][]int, target int) bool {

	if len(matrix) == 0 {
		return false
	}
	a := len(matrix) - 1
	b := len(matrix[0]) - 1
	i := 0
	j := b - 1
	for {
		if i > a || i < 0 {
			return false
		}
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] < target {
			i++
		} else {
			j--
		}
	}
}
