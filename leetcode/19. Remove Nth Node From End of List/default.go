package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	sentry := &ListNode{Next: head}
	p1, p2 := sentry, sentry
	for i := 0; i < n; i++ {
		p1 = p1.Next
	}
	for p1.Next != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	p2.Next = p2.Next.Next
	return sentry.Next
}
