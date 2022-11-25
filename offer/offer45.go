package offer

import (
	"fmt"
	"sort"
)

// 自定义排序
func minNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x := fmt.Sprintf("%d%d", nums[i], nums[j])
		y := fmt.Sprintf("%d%d", nums[j], nums[i])
		return x < y
	})
	res := ""
	for i := 0; i < len(nums); i++ {
		res += fmt.Sprint(nums[i])
	}
	return res
}
