package linkedlist

func removeElements(head *ListNode, val int) *ListNode {
	if head != nil && head.Val == val {
		head = head.Next
	}

	cur := head

	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return head
}