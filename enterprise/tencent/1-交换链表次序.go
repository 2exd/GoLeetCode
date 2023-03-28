package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
两个节点为一组交换次序
*/

func reorderList(head *ListNode) *ListNode {
	// write code here
	if head == nil {
		return nil
	}

	if head.Next == nil || head.Next.Next == nil {
		return head
	}

	vHead := &ListNode{}
	vHead.Next = head
	v := vHead
	tail := vHead.Next
	// var head1, head2, last1, last2 *ListNode
	var tmp *ListNode
	for tail != nil {
		head1, head2, last1, last2 := tmp, tmp, tmp, tmp
		for i := 0; i < 2 && tail != nil; i++ {
			if i == 0 {
				head1 = tail
			} else {
				last1 = tail
			}
			tail = tail.Next
		}
		for i := 0; i < 2 && tail != nil; i++ {
			if i == 0 {
				head2 = tail
			} else {
				last2 = tail
			}
			tail = tail.Next
		}
		// 找到四个节点
		if last2 != nil {
			v.Next = head2
			last2.Next = head1
			last1.Next = tail
		} else if head2 != nil {
			// 找到三个节点
			v.Next = head2
			head2.Next = head1
			last1.Next = tail
		}
		if last1 != nil {
			v = last1
		} else if head1 != nil {
			v = head1
		}
	}
	return vHead.Next
}
