package main

/*25-K 个一组翻转链表*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head
	pre := dummyNode

	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				// 到尾部就返回
				return dummyNode.Next
			}
		}
		next := tail.Next
		head, tail = myReverse(head, tail)

		pre.Next = head
		tail.Next = next
		pre = tail
		head = tail.Next
	}

	return dummyNode.Next
}

func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	cur := head
	for prev != tail {
		nex := cur.Next
		cur.Next = prev
		prev = cur
		cur = nex
	}
	return tail, head
}
