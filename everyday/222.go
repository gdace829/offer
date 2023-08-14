package everyday

import "math"

// 石子游戏2 1140
func stoneGameII(piles []int) int {
	var (
		dfs   func(i int, m int) int
		n     int
		min   func(int, int) int
		max   func(int, int) int
		cache [][]int
	)
	n = len(piles)
	cache = make([][]int, n)
	for i := 0; i < n; i++ {
		cache[i] = make([]int, 2*n)
		for j := 0; j < 2*n; j++ {
			cache[i][j] = -1
		}
	}
	min = func(i1, i2 int) int {
		if i1 > i2 {
			return i2
		}
		return i1
	}
	max = func(i1, i2 int) int {
		if i1 < i2 {
			return i2
		}
		return i1
	}

	for i := n - 2; i >= 0; i-- {
		piles[i] += piles[i+1]
	}

	dfs = func(i, m int) int {
		if i+2*m >= n {
			return piles[i]
		}
		if cache[i][m] != -1 {
			return cache[i][m]
		}
		res := math.MaxInt
		for j := 1; j <= 2*m; j++ {
			res = min(res, dfs(i+j, max(j, m)))
		}
		cache[i][m] = piles[i] - res
		return piles[i] - res
	}
	return dfs(0, 1)
}
