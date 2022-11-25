package offer

var trees []*TreeNode

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	trees = make([]*TreeNode, 1000)
	count := 0
	now := 1
	res := make([]int, 0)
	trees[0] = root
	for {
		if trees[count] == nil {
			break
		}
		if trees[count].Left != nil {
			trees[now] = trees[count].Left
			now++
		}
		if trees[count].Right != nil {
			trees[now] = trees[count].Right
			now++
		}
		count++

	}
	for i := 0; i < now-1; i++ {
		res = append(res, trees[i].Val)
	}
	return res
}

// 层次遍历一般用队列
func levelOrder1(root *TreeNode) []int {
	var queue []*TreeNode
	if root == nil {
		return nil
	}
	queue = make([]*TreeNode, 0)
	res := make([]int, 0)
	queue = append(queue, root)
	for {
		if len(queue) == 0 {
			return res
		}
		for len(queue) != 0 {
			res = append(res, queue[0].Val)
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			queue = queue[1:]
		}
	}

}
