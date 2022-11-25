package offer

func numWays(n int) int {
	if n == 0 {
		return 1
	} else if n == 1 {
		return 1
	}
	var a, b, c int
	a = 1
	b = 1
	for i := 0; i < n-1; {
		c = a%(1e9+7) + b%(1e9+7)
		i++
		if i == n-1 {
			return c % (1e9 + 7)
		} else {
			a = b
			b = c
		}
	}
	return 0
}
