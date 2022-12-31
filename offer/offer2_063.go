package offer

import (
	"strings"
)

type strtree struct {
	strees [26]*strtree
	isEnd  bool
}

func replaceWords(dictionary []string, sentence string) string {
	tree := &strtree{strees: [26]*strtree{}}
	for _, v := range dictionary {
		cur := tree
		for _, v1 := range v {
			if cur.strees[v1-'a'] == nil {
				cur.strees[v1-'a'] = &strtree{}

			}
			cur = cur.strees[v1-'a']

		}
		cur.isEnd = true
	}
	strs := strings.Split(sentence, " ")

	for index, v := range strs {
		cur := tree
		str := []byte{}
		for _, v1 := range v {
			if cur.strees[v1-'a'] == nil {
				break
			}
			str = append(str, byte(v1))
			if cur.strees[v1-'a'].isEnd == true {
				strs[index] = string(str)
				break
			}
			cur = cur.strees[v1-'a']

		}
	}

	res := ""
	for i := 0; i < len(strs)-1; i++ {
		res += strs[i] + " "
	}
	res += strs[len(strs)-1]
	return res
}
