package offer

type Trie struct {
	tris  [26]*Trie
	isEnd bool
}

/** Initialize your data structure here. */
func Constructor62() Trie {
	return Trie{tris: [26]*Trie{}, isEnd: false}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	cur := this
	for _, v := range word {
		if cur.tris[v-'a'] == nil {
			cur.tris[v-'a'] = &Trie{}
		}
		cur = cur.tris[v-'a']
	}
	cur.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	cur := this
	for _, v := range word {
		if cur == nil {
			return false
		}
		cur = cur.tris[v-'a']
	}
	return cur.isEnd != false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for _, v := range prefix {
		if cur == nil {
			return false
		}
		cur = cur.tris[v-'a']
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
