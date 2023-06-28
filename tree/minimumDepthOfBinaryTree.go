package tree

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := minDepth(root.Left)
	r := minDepth(root.Right)

	if l == 0 {
		return r + 1
	}
	if r == 0 {
		return l + 1
	}
	if l < r {
		return l + 1
	}
	return r + 1
}