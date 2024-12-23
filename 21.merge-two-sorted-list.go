package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	lr := &ListNode{}

	if l1.Val < l2.Val {
		lr = l1
		lr.Next = mergeTwoLists(lr.Next, l2)
	} else {
		lr = l2
		lr.Next = mergeTwoLists(l1, l2.Next)
	}

	return lr
}
