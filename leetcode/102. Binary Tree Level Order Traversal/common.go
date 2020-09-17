package main

import "math"

const null = math.MaxInt64

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTreeNode(nums []int) *TreeNode {
	if len(nums) == 0 || nums[0] == null {
		return nil
	}
	root := &TreeNode{Val: nums[0]}
	queue := []*TreeNode{root}
	for i := 1; i < len(nums); i = i + 2 {
		var leftVal, rightVal = nums[i], null
		if i+1 < len(nums) {
			rightVal = nums[i+1]
		}
		if leftVal != null {
			newPoint := &TreeNode{Val: leftVal}
			queue = append(queue, newPoint)
			queue[0].Left = newPoint
		}
		if rightVal != null {
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
			result = append(result, null)
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
		if result[i] != null {
			result = result[:i+1]
			break
		}
	}
	return result
}
