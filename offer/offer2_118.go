package offer

func findRedundantConnection(edges [][]int) []int {
	var (
		g     []int
		gnum  []int
		num   int
		find  func(int) int
		union func(int, int)
	)
	num = len(edges)
	g = make([]int, len(edges))
	gnum = make([]int, len(edges))
	for i := 0; i < num; i++ {
		g[i] = i
		gnum[i] = 1
	}
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
		if x1 == y1 {
			return
		}
		if gnum[y1] > gnum[x1] {
			x1, y1 = y1, x1
		}
		gnum[x1] += gnum[y1]
		gnum[y1] = 0
		g[y1] = x1
		num--
	}

	for _, v := range edges {
		x1 := v[0]
		y1 := v[1]
		if find(x1-1) != find(y1-1) {
			union(x1-1, y1-1)
		} else {
			return v
		}
	}
	return nil
}
