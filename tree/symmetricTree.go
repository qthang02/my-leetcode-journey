package tree

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return checkSameTree(root.Left, root.Right)
}

func checkSameTree(a, b *TreeNode) bool {
	if a == nil || b == nil {
		return a == b
	}
	if a.Val != b.Val {
		return false
	}
	return checkSameTree(a.Left, b.Right) && checkSameTree(a.Right, b.Left)
}