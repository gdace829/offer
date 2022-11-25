package offer

func twoSum(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1
	for l < r {
		if numbers[l]+numbers[r] == target {
			return []int{l, r}
		} else if numbers[l]+numbers[r] < target {
			l++
		} else {
			r--
		}
	}
	return []int{-1, -1}
}
