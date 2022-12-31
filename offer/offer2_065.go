package offer

type tree65 struct {
	ts     [26]*tree65
	height int
}

func minimumLengthEncoding(words []string) int {
	//初始化tree
	ts := &tree65{ts: [26]*tree65{}, height: 0}
	//插入单词 设定高度
	for _, v := range words {
		cur := ts
		for i := len(v) - 1; i >= 0; i-- {
			if cur.ts[v[i]-'a'] == nil {
				cur.ts[v[i]-'a'] = &tree65{}
			}
			cur.ts[v[i]-'a'].height = cur.height + 1
			cur = cur.ts[v[i]-'a']
		}
	}
	//dfs遍历获取每个路径最长高度
	res := 0
	var dfs func(*tree65)
	dfs = func(t *tree65) {
		haschild := false
		for i := 0; i < 26; i++ {
			if t.ts[i] != nil {
				haschild = true
				dfs(t.ts[i])
			}
		}
		if haschild == false {

			res += t.height
		}
	}
	dfs(ts)
	return res
}
