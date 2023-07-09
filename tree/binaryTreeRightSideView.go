package tree

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	rs := []int{}
	q := []*TreeNode{root}

	for len(q) > 0 {
		var right *TreeNode
		n := len(q)

		for i := 0; i < n; i++ {
			node := q[0]
			q = q[1:]
			if node != nil {
				right = node
				q = append(q, node.Left)
				q = append(q, node.Right)
			}
		}
		if right != nil {
			rs = append(rs, right.Val)
		}
	}

	return rs
}