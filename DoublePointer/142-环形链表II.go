package main

/*142. 环形链表 II medium*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func detectCycle(head *main.ListNode) *main.ListNode {
// 	slow, fast := head, head
// 	for fast != nil && fast.Next != nil {
// 		// slow移动一步，fast移动两步
// 		slow = slow.Next
// 		fast = fast.Next.Next
// 		// 确认有环
// 		if slow == fast {
// 			// 找到环入口
// 			for slow != head {
// 				slow = slow.Next
// 				head = head.Next
// 			}
// 			return head
// 		}
// 	}
// 	return nil
// }
