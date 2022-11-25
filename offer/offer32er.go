package offer

func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var queue []*TreeNode
	queue = make([]*TreeNode, 0)
	queue = append(queue, root)
	res := make([][]int, 0)
	for {
		if len(queue) == 0 {
			return res
		}
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
		res = append(res, res1)
		queue = queue1
	}
}
