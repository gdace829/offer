package offer


func dailyTemperatures(temperatures []int) []int {
	minstack := []int{}
	res := make([]int, len(temperatures))
	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(minstack) > 0 && temperatures[i] >= temperatures[minstack[len(minstack)-1]] {
			minstack = minstack[:len(minstack)-1]
		}
		if len(minstack) > 0 {
			res[i] = minstack[len(minstack)-1] - i
		} else {
			res[i] = 0
		}
		minstack = append(minstack, i)
	}
	return res
}

