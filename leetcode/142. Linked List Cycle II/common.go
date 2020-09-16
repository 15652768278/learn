package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func BuildListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	var sentry = &ListNode{Val: nums[0]}
	var cursor = sentry
	for i := 1; i < len(nums); i++ {
		cursor.Next = &ListNode{Val: nums[i]}
		cursor = cursor.Next
	}
	return sentry
}

func BuildCycleListNode(nums []int, target int) *ListNode {
	if target < 0 {
		return BuildListNode(nums)
	}
	if len(nums) == 0 {
		return nil
	}
	var head = &ListNode{Val: nums[0]}
	point := head
	for i := 1; i < len(nums); i++ {
		point.Next = &ListNode{Val: nums[i]}
		point = point.Next
	}
	cursor := head
	for i := 0; i < target; i++ {
		cursor = cursor.Next
	}
	point.Next = cursor
	return head
}

func (ln *ListNode) ToArray() []int {
	var ans []int
	cursor := ln
	for cursor != nil {
		ans = append(ans, cursor.Val)
		cursor = cursor.Next
	}
	return ans
}

func (ln *ListNode) Equals(oln *ListNode) bool {
	cursor1 := ln
	cursor2 := oln
	for cursor1 != nil && cursor2 != nil {
		if cursor1.Val != cursor2.Val {
			return false
		}
		cursor1 = cursor1.Next
		cursor2 = cursor2.Next
	}
	return cursor1 == nil || cursor2 == nil
}
