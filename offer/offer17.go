package offer

func printNumbers(n int) []int {
	var a int
	for {
		if n > 0 {
			a *= 10
			n--
		} else {
			break
		}
	}
	ans := make([]int, 0)
	for i := 1; i < a; i++ {
		ans = append(ans, i)
	}
	return ans
}
