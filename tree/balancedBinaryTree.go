package tree

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	diff := maxHeight(root.Left) - maxHeight(root.Right)

	if diff < -1 || diff > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func maxHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := maxHeight(root.Left)
	r := maxHeight(root.Right)

	if l > r {
		return l + 1
	}
	return r + 1
}