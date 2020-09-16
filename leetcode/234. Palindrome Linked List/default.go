package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	for tail := reverseList(slow); tail != nil && head.Next != nil; tail, head = tail.Next, head.Next {
		if tail.Val != head.Val {
			return false
		}
	}
	return true
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cursor := head
	for cursor != nil {
		prev, cursor, cursor.Next = cursor, cursor.Next, prev
	}
	return prev
}
