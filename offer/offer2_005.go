package offer

import "sort"

func maxProduct(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	}) //排序之后可以经过判断减少循环
	n := len(words)
	cur := make([]int, len(words))    //使用位来记录字符串字符是否出现
	length := make([]int, len(words)) //记录该字符串长度
	for i, word := range words {
		var num int
		for _, wo := range word {
			num |= (1 << (wo - 'a'))
		}
		cur[i] = num
		length[i] = len(word)
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if length[i]*length[j] < res {
				break
			}
			if cur[i]&cur[j] == 0 {
				res = length[i] * length[j]
			}
		}
	}
	return res
}
