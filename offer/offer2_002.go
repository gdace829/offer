package offer

func addBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1
	var sum int
	cur := 0
	next := 0
	res := ""
	for i >= 0 || j >= 0 {
		sum = next
		if i >= 0 {
			sum += int(a[i] - '0')
		}
		if j >= 0 {
			sum += int(b[j] - '0')
		}
		cur = sum % 2
		next = sum / 2
		res = string(cur+'0') + res
		i--
		j--
	}
	if next == 1 {
		res = "1" + res
	}
	return res
}
