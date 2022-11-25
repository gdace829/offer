package offer

var vis [][]bool
var dir [][]int
var word1 string

func exist(board [][]byte, word string) bool {
	var dfs func(board [][]byte, count, i, j int) bool
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	dfs = func(board [][]byte, count, i, j int) bool {
		if count == len(word1) {
			return true
		}
		if i < 0 || i > len(board) || j < 0 || j > len(board[0]) || vis[i][j] == true || board[i][j] != word1[count] {
			return false
		}
		if board[i][j] == word1[count] {
			vis[i][j] = true
			for i := 0; i < 4; i++ {
				if dfs(board, count+1, i+dir[i][0], j+dir[i][1]) {
					return true
				}
			}
			vis[i][j] = false
		}
		return false
	}
	dir = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	vis := make([][]bool, len(board))
	word1 = word
	for i := 0; i < len(board); i++ {
		vis[i] = make([]bool, len(board[0]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(board, 0, i, j) {
				return true
			}
		}
	}
	return false
}
