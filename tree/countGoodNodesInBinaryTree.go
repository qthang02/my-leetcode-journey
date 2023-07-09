package tree

func goodNodes(root *TreeNode) int {
	return count(root, root.Val)
}

func count(root *TreeNode, parent int) int {
	if root == nil {
		return 0
	}

	res := 1
	max := root.Val

	if parent > root.Val {
		res = 0
		max = parent
	}

	res += count(root.Left, max)
	res += count(root.Right, max)

	return res
}