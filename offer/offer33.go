package offer

func verifyPostorder(postorder []int) bool {
	if len(postorder) == 1 || len(postorder) == 0 {
		return true
	}
	count := 0
	for i := 0; i < len(postorder); i++ {
		if postorder[i] >= postorder[len(postorder)-1] {
			count = i
			break
		}
	}
	for i := count; i < len(postorder)-1; i++ {
		if postorder[i] < postorder[len(postorder)-1] {
			return false
		}
	}
	return verifyPostorder(postorder[:count]) && verifyPostorder(postorder[count:len(postorder)-1])
}
