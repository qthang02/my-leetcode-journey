package linkedlist

func getDecimalValue(head *ListNode) int {
	rs := 0
	for head != nil {
		rs = rs<<1 | head.Val
		head = head.Next
	}
	return rs
}