package tree

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	rs := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		tmp := []int{}
		n := len(queue)

		for i := 0; i < n; i++ {
			pop := queue[0]
			queue = queue[1:]
			tmp = append(tmp, pop.Val)
			if pop.Left != nil {
				queue = append(queue, pop.Left)
			}
			if pop.Right != nil {
				queue = append(queue, pop.Right)
			}
		}
		rs = append(rs, tmp)
	}

	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}

	return rs
}