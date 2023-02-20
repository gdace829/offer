package offer

func numSimilarGroups(strs []string) int {
	type group struct {
		num  int
		g    []int
		gnum []int
	}
	var (
		init    func(int) *group
		hanming func(int, int) bool
		find    func(int) int
		union   func(int, int)
	)
	init = func(i int) *group {
		x := &group{}

		x.num = i
		x.g = make([]int, i)
		x.gnum = make([]int, i)
		for j := 0; j < i; j++ {
			x.g[j] = j
			x.gnum[j] = 1
		}
		return x
	}
	x := init(len(strs))
	hanming = func(i1, i2 int) bool {
		nosame := 0
		for i := 0; i < len(strs[i1]); i++ {
			if strs[i1][i] != strs[i2][i] {
				nosame++
			}
		}
		return nosame == 0 || nosame == 2
	}
	find = func(i int) int {
		if x.g[i] == i {
			return i
		}
		x.g[i] = find(x.g[i])
		return x.g[i]
	}
	union = func(i1, i2 int) {
		x1 := find(i1)
		y1 := find(i2)
		if x1 == y1 {
			return
		}
		if x.gnum[y1] > x.gnum[x1] {
			x1, y1 = y1, x1
		}
		x.g[y1] = x1
		x.gnum[x1] += x.gnum[y1]
		x.gnum[y1] = 0
		x.num--

	}
	for i := 0; i < len(strs); i++ {
		for j := i + 1; j < len(strs); j++ {
			if hanming(i, j) {
				union(i, j)
			}
		}
	}
	return x.num
}
