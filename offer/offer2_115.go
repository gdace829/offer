package offer

func sequenceReconstruction(nums []int, sequences [][]int) bool {
	var (
		edges [][]int
		in    []int
	)
	edges = make([][]int, len(nums)+1)
	in = make([]int, len(nums)+1)
	for _, v := range sequences {
		for i := 1; i < len(v); i++ {
			edges[v[i-1]] = append(edges[v[i-1]], v[i])
			in[v[i]]++
		}
	}
	q := []int{}
	for i := 1; i < len(in); i++ {
		if in[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		if len(q) > 1 {
			return false
		}

		for _, v := range edges[q[0]] {
			if in[v]--; in[v] == 0 {
				q = append(q, v)
			}
		}
		q = q[1:]
	}
	return true
}
