package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}
	sentry := &ListNode{Next: head}
	wait := sentry
	for i := 1; i < m; i++ {
		wait = wait.Next
	}
	var prev, cur, next *ListNode = nil, wait.Next, wait.Next.Next
	for i := m; i < n; i++ {
		prev, cur, next, cur.Next, next.Next = cur, next, next.Next, prev, cur
	}
	wait.Next, wait.Next.Next = cur, next
	return sentry.Next
}
