package offer

func plusOne(digits []int) []int {
	n := len(digits)

	for n > 1 && digits[n-1] == 9 {
		digits[n-1] = 0
		n--
	}
	if digits[n-1] != 9 {
		digits[n-1]++
		return digits
	}
	if digits[0] == 9 {
		digits[0] = 1
		digits = append(digits, 0)
		return digits
	}
	return digits
}
