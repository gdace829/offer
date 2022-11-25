package offer

func exist1(board [][]byte, word string) bool {
	visit = make([][]bool, len(board))
	for i := range visit {
		visit[i] = make([]bool, len(board[0]))
	}
	for i := range board {
		for j := range board[i] {
			if run(board, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

var visit [][]bool
var directs [][]int = [][]int{
	{1, 0}, {-1, 0}, {0, 1}, {0, -1},
} //这个地方可以借鉴一下

func run(board [][]byte, word string, i, x, y int) bool {
	if i == len(word) {
		return true
	}
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) {
		return false
	}
	if visit[x][y] {
		return false
	}
	if word[i] != board[x][y] {
		return false
	}
	visit[x][y] = true

	for _, direct := range directs {
		newX, newY := x+direct[0], y+direct[1]
		if run(board, word, i+1, newX, newY) {
			return true
		}
	}
	visit[x][y] = false

	return false
}
