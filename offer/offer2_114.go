package offer

func alienOrder(words []string) string {
	var (
		edges map[byte][]byte
		vis   map[byte]int
		is    bool
		res   []byte
		dfs   func(byte)
	)
	edges = make(map[byte][]byte)
	vis = make(map[byte]int)
	is = true
	for _, c := range words[0] {
		edges[byte(c)] = nil
	}

next:
	for i := 1; i < len(words); i++ {
		for _, c := range words[i] {
			edges[byte(c)] = edges[byte(c)]
		}
		for j := 0; j < len(words[i]) && j < len(words[i-1]); j++ {

			if words[i][j] != words[i-1][j] {
				edges[words[i][j]] = append(edges[words[i][j]], words[i-1][j])
				continue next
			}
		}
		if len(words[i-1]) > len(words[i]) {
			return ""
		}
	}
	res = make([]byte, 0)

	dfs = func(i byte) {
		vis[i] = 1
		for _, v := range edges[i] {
			if vis[v] == 0 {
				dfs(v)
				if !is {
					return
				}
			} else if vis[v] == 1 {
				is = false
				return
			}
		}
		vis[i] = 2
		res = append(res, i)
	}
	for k, _ := range edges {
		if vis[k] == 0 {
			dfs(k)
		}
		if !is {
			return ""
		}
	}
	return string(res)
}
