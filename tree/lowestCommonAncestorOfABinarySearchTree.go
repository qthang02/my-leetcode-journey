package tree

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	cur := root
	for {
		switch {
		case p.Val < cur.Val && q.Val < cur.Val:
			cur = cur.Left
		case p.Val > cur.Val && q.Val > cur.Val:
			cur = cur.Right
		default:
			return cur
		}
	}
	return nil
}