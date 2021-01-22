package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getKthFromEnd(head *ListNode, k int) *ListNode {
	sentry := &ListNode{Next: head}
	slow, quick := sentry, sentry
	for i := 0; i < k; i++ {
		quick = quick.Next
	}
	for quick != nil {
		slow = slow.Next
		quick = quick.Next
	}
	return slow
}
