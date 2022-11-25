package offer

func hammingWeight(num uint32) int {
	count := 0
	for {
		if num == 0 {
			break
		}
		if num&1 == 1 {
			count += 1

		}
		num /= 2
	}
	return count

}
