package linkedlist

import "testing"

func TestMergeList(t *testing.T) {
	l1_tc1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	l2_tc1 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	rs_tc1 := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4}}}}}}

	l1_tc2 := &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	l2_tc2 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 10}}}
	rs_tc2 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 10}}}}}}

	l1_tc3 := &ListNode{}
	l2_tc3 := &ListNode{Val: 0}
	rs_tc3 := &ListNode{Val: 0}

	l1_tc4 := &ListNode{}
	l2_tc4 := &ListNode{}
	rs_tc4 := &ListNode{}

	tc := []struct {
		list1  *ListNode
		list2  *ListNode
		result *ListNode
	}{
		{l1_tc1, l2_tc1, rs_tc1},
		{l1_tc2, l2_tc2, rs_tc2},
		{l1_tc3, l2_tc3, rs_tc3},
		{l1_tc4, l2_tc4, rs_tc4},
	}

	for _, c := range tc {
		name := "mergeTwoLists(" + c.list1.String() + ", " + c.list2.String() + ")"

		t.Run(name, func(t *testing.T) {
			result := mergeTwoLists(c.list1, c.list2)
			if !checkResult(t, result, c.result) {
				t.Errorf("got %s, want %s", result.String(), c.result.String())
			}
		})
	}
}

func checkResult(t *testing.T, got, want *ListNode) bool {
	for got != nil && want != nil {
		if got.Val != want.Val {
			return false
		}
		got = got.Next
		want = want.Next
	}
	return true
}