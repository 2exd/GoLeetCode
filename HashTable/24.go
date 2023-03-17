package main

/*24. 两两交换链表中的节点 medium*/
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	// head=list[i]
	// pre=list[i-1]
	pre := dummy
	for head != nil && head.Next != nil {
		pre.Next = head.Next

		next := head.Next.Next
		head.Next.Next = head
		head.Next = next

		// pre=list[(i+2)-1]
		pre = head
		// head=list[(i+2)]
		head = next
	}
	return dummy.Next
}
func ListNodeConstructor() *ListNode {
	head := &ListNode{
		Val:  -1,
		Next: nil,
	}
	return head
}

func (this *ListNode) ListNodeAddAtTail(val int) {
	tail := &ListNode{
		Val:  val,
		Next: nil,
	}
	this.Next = tail
}

// func main() {
// 	L := &ListNode{1, nil}
// 	L.ListNodeAddAtTail(2)
// 	L.ListNodeAddAtTail(3)
// 	L.ListNodeAddAtTail(4)
// 	swapPairs(L)
// }
