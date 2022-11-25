package offer

// 搜索二分法 注意相等的时候j指针怎么变
func minArray(numbers []int) int {
	i := 0
	j := len(numbers) - 1
	var mid int
	for {
		mid = (i + j) / 2
		if i >= j {
			return numbers[j]
		}
		if numbers[mid] == numbers[j] {
			j--
			continue
		}
		if numbers[mid] < numbers[j] {
			j = mid
		} else if numbers[mid] > numbers[j] {
			i = mid + 1
		}

	}
}
