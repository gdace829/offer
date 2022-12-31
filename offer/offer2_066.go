package offer

type sumtree struct {
	strees [26]*sumtree
	value  int
	isEnd  bool
}
type MapSum struct {
	s sumtree
}

/** Initialize your data structure here. */
func Constructor66() MapSum {
	return MapSum{s: sumtree{strees: [26]*sumtree{}}}
}

func (this *MapSum) Insert(key string, val int) {
	cur := &this.s
	for _, v := range key {
		if cur.strees[v-'a'] == nil {
			cur.strees[v-'a'] = &sumtree{}
		}
		cur = cur.strees[v-'a']
	}
	cur.isEnd = true
	cur.value = val
}

func (this *MapSum) Sum(prefix string) int {
	var find func(*sumtree)
	var dfs func(*sumtree)
	res := 0
	dfs = func(s *sumtree) {
		if s != nil && s.isEnd == true {
			res += s.value
		}
		for i := 0; i < 26; i++ {
			if s.strees[i] != nil {
				dfs(s.strees[i])
			}
		}
	}
	find = func(st *sumtree) {
		cur := &this.s
		for _, v := range prefix {
			if cur.strees[v-'a'] != nil {
				cur = cur.strees[v-'a']
			} else {
				return
			}

		}
		if cur.isEnd == true {
			res += cur.value
		}
		for i := 0; i < 26; i++ {
			if cur.strees[i] != nil {
				dfs(cur.strees[i])
			}
		}

	}
	find(&this.s)
	return res
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
