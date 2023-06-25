package linkedlist

import "testing"

func TestDeleteDuplicates(t *testing.T) {
	head_tc1 := &ListNode{1, &ListNode{1, &ListNode{2, nil}}}
	rs_tc1 := &ListNode{1, &ListNode{2, nil}}

	head_tc2 := &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}}
	rs_tc2 := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}

	tc := []struct {
		head *ListNode
		rs   *ListNode
	}{
		{head_tc1, rs_tc1},
		{head_tc2, rs_tc2},
	}

	for _, c := range tc {
		name := "deleteDuplicates(" + c.head.String() + ")"

		t.Run(name, func(t *testing.T) {
			rs := deleteDuplicates(c.head)
			if !checkResult(t, rs, c.rs) {
				t.Errorf("got %v want %v", rs, c.rs)
			}
		})
	}
}