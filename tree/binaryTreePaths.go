package tree

import "strconv"

func binaryTreePaths(root *TreeNode) []string {
	rs := []string{}
	if root == nil {
		return rs
	}
	dfs(root, "", &rs)
	return rs
}

func dfs(root *TreeNode, path string, arr *[]string) {
	path += strconv.Itoa(root.Val)

	if root.Left == nil && root.Right == nil {
		*arr = append(*arr, path)
		return
	}
	if root.Left != nil {
		dfs(root.Left, path + "->", arr)
	}
	if root.Right != nil {
		dfs(root.Right, path + "->", arr)
	}
}