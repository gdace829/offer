package offer

type strTree struct {
	trees [26]*strTree
	isEnd bool
}
type MagicDictionary struct {
	tree *strTree
}

/** Initialize your data structure here. */
func Constructor64() MagicDictionary {
	return MagicDictionary{tree: &strTree{trees: [26]*strTree{}}}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for _, v := range dictionary {
		cur := this.tree
		for _, v1 := range v {
			if cur.trees[v1-'a'] == nil {
				cur.trees[v1-'a'] = &strTree{}
			}
			cur = cur.trees[v1-'a']
		}
		cur.isEnd = true
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	var find func(string, strTree) bool
	find = func(s string, cur strTree) bool {
		for _, v := range s {
			if cur.trees[v-'a'] == nil {
				return false
			}
			cur = *cur.trees[v-'a']

		}
		if cur.isEnd == true {
			return true
		}
		return false
	}
	cur := this.tree
	strs := []byte(searchWord)
	for i := 0; i < len(strs); i++ {
		char := strs[i]
		for j := 0; j < 26; j++ {
			if char == byte('a'+j) {
				continue
			}
			strs[i] = byte('a' + j)
			if find(string(strs[i:]), *cur) {
				return true
			}
		}
		strs[i] = char
		if cur.trees[strs[i]-'a'] == nil {
			return false
		}
		cur = cur.trees[strs[i]-'a']
	}
	return false
}
