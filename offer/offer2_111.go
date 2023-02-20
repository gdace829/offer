package offer

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	//形成编号
	graphid := make(map[string]int)
	for _, v := range equations {
		if _, ok := graphid[v[0]]; !ok {
			graphid[v[0]] = len(graphid)
		}
		if _, ok := graphid[v[1]]; !ok {
			graphid[v[1]] = len(graphid)
		}
	}
	//建图
	type edge struct {
		to    int
		value float64
	}
	edges := make([][]edge, len(graphid))
	for k, v := range equations {
		edges[graphid[v[0]]] = append(edges[graphid[v[0]]], edge{to: graphid[v[1]], value: values[k]})
		edges[graphid[v[1]]] = append(edges[graphid[v[1]]], edge{to: graphid[v[0]], value: 1 / values[k]})
	}
	//bfs路径累乘
	var bfs func(int, int) float64
	bfs = func(start int, end int) float64 {
		ratio := make([]float64, len(graphid))
		ratio[start] = 1
		queue := []int{start}
		for len(queue) > 0 {
			egde := queue[0]
			queue = queue[1:]
			if egde == end {
				return ratio[egde]
			}
			for _, v := range edges[egde] {
				if n := v.to; ratio[n] == 0 {
					ratio[n] = ratio[egde] * v.value
					queue = append(queue, n)
				}
			}

		}
		return -1
	}
	res := []float64{}
	for _, v := range queries {
		_, ok := graphid[v[0]]
		_, ok1 := graphid[v[1]]
		if !ok || !ok1 {
			res = append(res, -1)
		} else {
			res = append(res, bfs(graphid[v[0]], graphid[v[1]]))
		}

	}
	return res

}
