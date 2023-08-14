package offer

func longestConsecutive(nums []int) int {
	var (
		max   int
		num   int
		gnum  map[int]int
		g     map[int]int
		find  func(int) int
		union func(int, int)
	)

	find = func(i int) int {
		if g[i] == i {
			return i
		}
		g[i] = find(g[i])
		return g[i]
	}
	union = func(i1, i2 int) {
		x1 := find(i1)
		y1 := find(i2)
		if gnum[y1] > gnum[x1] {
			x1, y1 = y1, x1
		}
		gnum[x1] += gnum[y1]
		if gnum[x1] > max {
			max = gnum[x1]
		}
		g[y1] = g[x1]
		num--
	}
	num = len(nums)
	if num >= 1 {
		max = 1
	}
	gnum = make(map[int]int)
	g = make(map[int]int)
	for _, v := range nums {
		g[v] = v
		gnum[v] = 1
	}
	for v, _ := range g {
		if _, ok := g[v-1]; ok && find(v-1) != find(v) {
			union(v, v-1)
		}
	}
	return max

}
