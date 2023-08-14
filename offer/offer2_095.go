package offer

func longestCommonSubsequence(text1 string, text2 string) int {
	numss := make([][]int, len(text1)+1)
	for i := 0; i <= len(text1); i++ {
		numss[i] = make([]int, len(text2)+1)
	}
	var max func(int, int) int
	max = func(i1, i2 int) int {
		if i1 > i2 {
			return i1
		}
		return i2
	}
	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				numss[i][j] = numss[i-1][j-1] + 1
			} else {
				numss[i][j] = max(numss[i-1][j], numss[i][j-1])
			}
		}
	}
	return numss[len(text1)][len(text2)]
}
