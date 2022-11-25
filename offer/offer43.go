package offer

func countDigitOne(n int) int {
	high, cur, low, dight := 0, 0, 0, 0
	cur = n % 10
	high = n / 10
	dight = 1
	res := 0
	for high != 0 || cur != 0 {
		if cur == 0 {
			res += dight * high
		} else if cur == 1 {
			res += dight*high + low + 1
		} else {
			res += dight * (high + 1)
		}
		low = cur*dight + low
		cur = high % 10
		high /= 10
		dight *= 10
	}
	return res
}
