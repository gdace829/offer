package offer

func countBits(n int) []int {
	res := make([]int, n+1)
	res[0] = 0
	for i := 1; i <= n; i++ {
		res[i] = res[i>>1]
		if i%2 == 1 {
			res[i]++
		}

	}
	return res
}
