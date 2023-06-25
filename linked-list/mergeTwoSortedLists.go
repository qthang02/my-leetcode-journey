package linkedlist

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head ListNode
	prev := &head

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			prev.Next = list1
			list1 = list1.Next
		} else {
			prev.Next = list2
			list2 = list2.Next
		}

		prev = prev.Next
	}

	return head.Next
}