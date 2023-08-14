package offer

func minFlipsMonoIncr(s string) int {
	var min func(int, int) int
	min = func(i1, i2 int) int {
		if i1 > i2 {
			return i2
		}
		return i1
	}
	one, zero := 0, 0
	for _, v := range s {
		if v == '0' {
			one = min(one, zero) + 1
		} else {
			one = min(zero, one)
			zero = zero + 1
		}
	}
	return min(one, zero)
}
