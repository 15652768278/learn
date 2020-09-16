package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	slow, quick := head, head
	for quick != nil && quick.Next != nil {
		slow, quick = slow.Next, quick.Next.Next
		if quick == slow {
			return true
		}
	}
	return false
}
