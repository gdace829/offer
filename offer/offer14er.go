package offer

func cuttingRope2(n int) int {
	if n < 4 {
		return n - 1
	}
	var res int64
	res = 1
	for {
		if n > 4 {
			n -= 3
			res *= 3
			res %= 1e9 + 7
		} else {
			res *= int64(n) % (1e9 + 7)
			break
		}
	}
	return int(res % (1e9 + 7))
}
