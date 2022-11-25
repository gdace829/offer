package offer

type NumMatrix struct {
	presums [][]int
}

func Constructor13(matrix [][]int) NumMatrix {
	var presums [][]int
	presums = make([][]int, len(matrix)+1)
	for i := 0; i < len(presums); i++ {
		presums[i] = make([]int, len(matrix[0])+1)
	}
	for i := 1; i < len(presums); i++ {
		for j := 1; j < len(presums[0]); j++ {
			presums[i][j] = presums[i-1][j] + presums[i][j-1] + matrix[i-1][j-1] - presums[i-1][j-1]
		}
	}

	return NumMatrix{presums: presums}

}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {

	return this.presums[row2+1][col2+1] - this.presums[row2+1][col1] - this.presums[row1][col2+1] + this.presums[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
