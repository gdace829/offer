package offer

func findNthDigit(n int) int {
	dightnum, count, num := 0, 0, 0
	count = 9
	dightnum = 1
	num = 1
	for n > count {
		n -= count
		dightnum += 1
		num *= 10
		count = 9 * num * dightnum
	}
	dight := num + (n-1)/dightnum        //那个数字
	dightnum = dightnum - (n-1)%dightnum //数字哪一位
	for dightnum != 1 {
		dight /= 10
		dightnum--
	}
	return dight % 10
}
