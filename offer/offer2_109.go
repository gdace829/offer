package offer

func openLock(deadends []string, target string) int {
	var up func(string, int) string
	var down func(string, int) string
	up = func(s string, i int) string {
		b := []byte(s)
		if b[i] == '9' {
			b[i] = '0'
		} else {
			b[i] += 1
		}
		return string(b)
	}
	down = func(s string, i int) string {
		b := []byte(s)
		if b[i] == '0' {
			b[i] = '9'
		} else {
			b[i] -= 1
		}
		return string(b)
	}
	dead := make(map[string]bool)
	for i := 0; i < len(deadends); i++ {
		dead[deadends[i]] = true
	}
	vis := make(map[string]bool)
	bfs := []string{}
	bfs = append(bfs, "0000")
	vis["0000"] = true
	res := 0
	for len(bfs) > 0 {
		i, n := 0, len(bfs)
		for ; i < n; i++ {
			if dead[bfs[i]] {
				continue
			}
			if bfs[i] == target {
				return res
			}
			for j := 0; j < 4; j++ {
				u := up(bfs[i], j)

				d := down(bfs[i], j)
				if !vis[u] {
					bfs = append(bfs, u)
					vis[u] = true
				}
				if !vis[d] {
					bfs = append(bfs, d)
					vis[d] = true
				}
			}

		}
		bfs = bfs[i:]
		res++
	}
	return -1
}
