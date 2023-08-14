package everyday

func largestLocal(grid [][]int) [][]int {
	res := make([][]int, 0)
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	for i := 0; i < len(grid); i++ {
		ans := make([]int, 0)
		for j := 0; j < len(grid[0]); j++ {
			if i-1 < 0 || j-1 < 0 || i+1 >= len(grid) || j+1 >= len(grid[0]) {
				continue
			}
			maxnum := grid[i][j]

			for k := -1; k <= 1; k++ {
				for l := -1; l <= 1; l++ {
					maxnum = max(maxnum, grid[i+k][j+l])
				}
			}

			ans = append(ans, maxnum)

		}
		if len(ans) > 0 {
			res = append(res, ans)
		}

	}
	return res
}
