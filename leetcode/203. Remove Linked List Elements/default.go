package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	var sentry = &ListNode{Next: head}
	last, cursor := sentry, sentry.Next
	for cursor != nil {
		if cursor.Val == val {
			last.Next = cursor.Next
			cursor = last.Next
		} else {
			last, cursor = cursor, cursor.Next
		}
	}
	return sentry.Next
}
