package offer

// . 代表任意一个字符 * 前面的字符出现若干次
func isMatch(s string, p string) bool {
	var nums [][]bool
	nums = make([][]bool, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		nums[i] = make([]bool, len(p)+1)
	}
	nums[0][0] = true
	for i := 0; i < len(s)+1; i++ {
		for j := 1; j < len(p)+1; j++ {
			if p[j-1] == '*' {
				nums[i][j] = nums[i][j] || nums[i][j-2]
				if isMatch1(i, j-1, s, p) {
					nums[i][j] = nums[i][j] || nums[i-1][j]
				}

			} else {
				nums[i][j] = isMatch1(i, j, s, p) && nums[i-1][j-1]
			}

		}
	}
	if nums[len(s)][len(p)] == true {
		return true
	} else {
		return false
	}

}
func isMatch1(a int, b int, s string, p string) bool {
	if a == 0 || b == 0 {
		return false
	}
	if p[b-1] == '.' {
		return true
	}
	return s[a-1] == p[b-1]
}
