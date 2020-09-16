package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, quick := head, head
	for quick != nil && quick.Next != nil {
		slow, quick = slow.Next, quick.Next.Next
		if quick == slow {
			break
		}
	}
	if quick == nil || quick.Next == nil {
		return nil
	}
	quick = head
	for quick != slow {
		slow, quick = slow.Next, quick.Next
	}
	return slow
}
