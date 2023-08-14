package offer

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || (intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1])
	})
	res := [][]int{}
	for _, v := range intervals {
		if len(res) == 0 || v[0] > res[len(res)-1][1] {
			res = append(res, v)
			continue
		} else {
			a := res[len(res)-1][1]
			if v[1] > a {
				a = v[1]
			}
			res[len(res)-1][1] = a
		}
	}
	return res
}
