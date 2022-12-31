package offer

import "fmt"

type twoTree struct {
	tts [2]*twoTree
}

func findMaximumXOR(nums []int) int {
	root := &twoTree{tts: [2]*twoTree{}}
	var insert func(int)
	insert = func(num int) {
		cur := root
		for i := 30; i >= 0; i-- {
			bit := (num >> i) & 1
			if cur.tts[bit] == nil {
				cur.tts[bit] = &twoTree{}
			}
			cur = cur.tts[bit]
		}
	}
	insert(nums[0])
	//插入新值判断是否可以有最大值
	var findmax func(int) int
	findmax = func(num int) int {
		res := 0
		cur := root
		for i := 30; i >= 0; i-- {
			bit := (num >> i) & 1
			if cur.tts[bit^1] != nil {
				res |= 1 << i
				cur = cur.tts[bit^1]
			} else {
				cur = cur.tts[bit]
			}
		}
		return res
	}
	res := 0
	for i := 1; i < len(nums); i++ {
		a := findmax(nums[i])
		if res < a {
			fmt.Println("sjs")
			res = a

		}
		insert(nums[i])
	}
	return res
}
