package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	r, l := head, dummy
	for i := 0; i < n; i++ {
		r = r.Next
	}

	for r != nil {
		r = r.Next
		l = l.Next
	}
	l.Next = l.Next.Next
	return dummy.Next
}
