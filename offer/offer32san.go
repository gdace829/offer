package offer

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder3(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	count := 0
	for len(queue) != 0 {
		res1 := make([]int, 0)
		queue1 := make([]*TreeNode, 0)
		for len(queue) != 0 {
			res1 = append(res1, queue[0].Val)
			if queue[0].Left != nil {
				queue1 = append(queue1, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue1 = append(queue1, queue[0].Right)
			}
			queue = queue[1:]

		}
		if count%2 == 0 {
			res = append(res, res1)
		} else {
			res = append(res, reverse(res1))
		}
		count++
		queue = queue1
	}
	return res

}

func reverse(res1 []int) []int {
	for i := 0; i < len(res1)/2; i++ {
		t := res1[i]
		res1[i] = res1[len(res1)-i-1]
		res1[len(res1)-1-i] = t
	}
	return res1
}
