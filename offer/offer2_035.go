package offer

import (
	"math"
	"sort"
)

func findMinDifference(timePoints []string) int {
	if len(timePoints) >= 1440 {
		return 0
	}
	var min func(int, int) int
	var getms func(string) int
	getms = func(s string) int {
		return (int(s[0]-'0')*10+int(s[1]-'0'))*60 + int(s[3]-'0')*10 + int(s[4]-'0')
	}
	min = func(i1, i2 int) int {
		if i1 > i2 {
			return i2
		} else {
			return i1
		}
	}
	sort.Strings(timePoints)
	firtime := getms(timePoints[0])
	pretime := firtime
	res := math.MaxInt32
	curtime := 0
	for _, v := range timePoints[1:] {
		curtime = getms(v)
		res = min(res, curtime-pretime)
		pretime = curtime
	}
	res = min(res, firtime+1440-curtime)
	return res
}
