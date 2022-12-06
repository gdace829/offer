package offer

func largestRectangleArea(heights []int) int {
	stack := [][2]int{[2]int{-1, 0}}
	var max func(int, int) int
	max = func(i1, i2 int) int {
		if i1 > i2 {
			return i1
		} else {
			return i2
		}
	}
	res := 0
	for i, v := range heights {
		if v == 0 {
			stack = stack[0:1]
			stack[0][0] = i
			stack[0][1] = 0
			continue
		}
		index := len(stack) - 1
		for len(stack) > 0 && stack[index][1] >= v {
			index--
		}
		index++
		stack = stack[0:index]
		stack = append(stack, [2]int{i, v})
		for k := len(stack) - 1; k >= 1; k-- {
			res = max(res, (i-stack[k-1][0])*stack[k][1])
		}

	}

	return res

}
