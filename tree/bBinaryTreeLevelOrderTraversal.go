package tree

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	rs := [][]int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := len(queue)
		tmp := []int{}
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

	return rs
}