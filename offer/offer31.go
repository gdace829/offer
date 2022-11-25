package offer

func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) == 0 {
		return true
	}
	stack := make([]int, 1000)
	count := 0
	count1 := 0
	count2 := 0
	stack[0] = pushed[0]
	count++
	count2++
	for {
		if count1 == len(pushed) && count2 == len(pushed) {
			return true
		}
		if count2 == len(pushed) && stack[count-1] != popped[count1] {
			return false
		}
		if count == 0 && count1 < len(pushed) {
			stack[count] = pushed[count2]
			count++
			count2++
			continue
		}
		if stack[count-1] == popped[count1] {
			stack[count-1] = -1
			count--
			count1++
		} else {
			stack[count] = pushed[count2]
			count++
			count2++
		}
	}
}
