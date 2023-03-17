package main

/*92-反转链表 II*/

func reverseBetween(head *ListNode, left, right int) *ListNode {
	// dummy node
	dummy := &ListNode{Val: -1}
	dummy.Next = head

	// find the left -1 node
	prev := dummy
	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}

	cur := prev.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = prev.Next
		// merge
		prev.Next = next
	}
	return dummy.Next
}
