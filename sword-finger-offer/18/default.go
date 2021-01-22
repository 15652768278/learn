package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(head *ListNode, val int) *ListNode {
	sentry := &ListNode{Next: head}
	prev, cursor := sentry, head
	for cursor != nil && cursor.Val != val {
		prev = prev.Next
		cursor = cursor.Next
	}
	if cursor != nil {
		prev.Next = cursor.Next
	}
	return sentry.Next
}
