package offer

import "math"

func divide(a int, b int) int {
	var sign int
	if a > 0 && b > 0 || a < 0 && b < 0 {
		sign = 0
	} else {
		sign = 1
	}
	a, b = int(math.Abs(float64(a))), int(math.Abs(float64(b)))
	c := 31
	res := 0
	for a >= b {
		for a < (b << c) {
			c--
		}
		a -= (b << c)
		res += (1 << c)
	}
	if sign == 1 {
		res = -res
	}
	if res >= math.MaxInt32 {
		return math.MaxInt32
	} else {
		return res
	}
}
