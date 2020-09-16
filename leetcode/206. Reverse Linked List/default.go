package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cursor := head
	for cursor != nil {
		prev, cursor, cursor.Next = cursor, cursor.Next, prev
	}
	return prev
}
