package offer

func findOrder(numCourses int, prerequisites [][]int) []int {
	var (
		edges [][]int
		vis   []int
		res   []int
		dfs   func(int)
		is    bool
	)
	edges = make([][]int, numCourses)
	vis = make([]int, numCourses)
	is = true
	dfs = func(i int) {
		vis[i] = 1

		for j := 0; j < len(edges[i]); j++ {
			if vis[edges[i][j]] == 0 {
				dfs(edges[i][j])
			} else if vis[edges[i][j]] == 1 {
				is = false
				return
			}

		}
		if !is {
			return
		}

		vis[i] = 2
		res = append(res, i)
	}
	for i := 0; i < len(prerequisites); i++ {
		edges[prerequisites[i][1]] = append(edges[prerequisites[i][1]], prerequisites[i][0])
	}
	for i := 0; i < numCourses; i++ {
		if vis[i] == 0 {
			dfs(i)
		}
		if !is {
			return []int{}
		}
	}
	for i := 0; i < numCourses/2; i++ {
		res[i], res[numCourses-1-i] = res[numCourses-1-i], res[i]
	}
	return res

}
