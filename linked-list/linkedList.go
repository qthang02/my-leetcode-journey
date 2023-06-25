package linkedlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	var s string
	for l != nil {
		s += fmt.Sprintf("%d ", l.Val)
		l = l.Next
	}
	return s
}