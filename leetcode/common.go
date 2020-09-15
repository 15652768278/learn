package main

import "math"

// Node
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func BuildNode(nums [][]int) *Node {
	var head *Node
	var point *Node
	for _, v := range nums {
		if head == nil {
			head = &Node{Val: v[0]}
			point = head
		} else {
			point.Next = &Node{Val: v[0]}
			point = point.Next
		}
	}
	cursor := head
	index := 0
	for cursor != nil {
		if nums[index][1] != NULL {
			p := head
			for i := nums[index][1]; i > 0; i-- {
				p = p.Next
			}
			cursor.Random = p
		}
		cursor = cursor.Next
		index++
	}
	return head
}

// List Node
type ListNode struct {
	Val  int
	Next *ListNode
}

func BuildListNode(nums []int) *ListNode {
	var head *ListNode
	var point *ListNode
	for _, v := range nums {
		if head == nil {
			head = &ListNode{Val: v}
			point = head
		} else {
			point.Next = &ListNode{Val: v}
			point = point.Next
		}
	}
	return head
}

func BuildCycleListNode(nums []int, target int) *ListNode {
	var head *ListNode
	var point *ListNode
	for _, v := range nums {
		if head == nil {
			head = &ListNode{Val: v}
			point = head
		} else {
			point.Next = &ListNode{Val: v}
			point = point.Next
		}
	}
	cursor := head
	for i := 0; i < target; i++ {
		cursor = cursor.Next
	}
	point.Next = cursor
	return head
}

func (ln *ListNode) ToArray() []int {
	var result []int
	cursor := ln
	for cursor != nil {
		result = append(result, cursor.Val)
		cursor = cursor.Next
	}
	return result
}

// Tree Node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const NULL = math.MaxInt64

func BuildTreeNode(nums []int) *TreeNode {
	if len(nums) == 0 || nums[0] == NULL {
		return nil
	}
	root := &TreeNode{Val: nums[0]}
	queue := []*TreeNode{root}
	for i := 1; i < len(nums); i = i + 2 {
		var leftVal, rightVal = nums[i], NULL
		if i+1 < len(nums) {
			rightVal = nums[i+1]
		}
		if leftVal != NULL {
			newPoint := &TreeNode{Val: leftVal}
			queue = append(queue, newPoint)
			queue[0].Left = newPoint
		}
		if rightVal != NULL {
			newPoint := &TreeNode{Val: rightVal}
			queue = append(queue, newPoint)
			queue[0].Right = newPoint
		}
		queue = queue[1:]
	}
	return root
}

func (tn *TreeNode) ToArray() []int {
	var result []int
	array := []*TreeNode{tn}
	for len(array) > 0 {
		elem := array[0]
		if elem == nil {
			result = append(result, NULL)
		} else {
			result = append(result, elem.Val)
			if elem.Left != nil {
				array = append(array, elem.Left)
			} else {
				array = append(array, nil)
			}
			if elem.Right != nil {
				array = append(array, elem.Right)
			} else {
				array = append(array, nil)
			}
		}
		array = array[1:]
	}
	for i := len(result) - 1; i >= 0; i-- {
		if result[i] != NULL {
			result = result[:i+1]
			break
		}
	}
	return result
}
