package everyday

import (
	"math"
	"sort"
)

// 最少水龙头数目 1326
func minTaps(n int, ranges []int) int {
	var (
		nums  []int
		numss [][]int
		max   func(int, int) int
		min   func(int, int) int
	)
	max = func(i1, i2 int) int {
		if i1 > i2 {
			return i1
		}
		return i2
	}
	min = func(i1, i2 int) int {
		if i1 > i2 {
			return i2
		}
		return i1
	}
	numss = make([][]int, 0)
	for i := 0; i < len(ranges); i++ {
		x := max(0, i-ranges[i])
		y := min(n, i+ranges[i])
		numss = append(numss, []int{x, y})
	}
	sort.Slice(numss, func(i, j int) bool {
		if numss[i][0] < numss[j][0] {
			return true
		}
		return false
	})
	nums = make([]int, n+1)
	for i := 1; i <= n; i++ {
		nums[i] = math.MaxInt
	}
	for _, v := range numss {
		if nums[v[0]] == math.MaxInt {
			return -1
		}
		for i := v[0]; i <= v[1]; i++ {
			nums[i] = min(nums[i], nums[v[0]]+1)
		}
	}
	return nums[n]
}
