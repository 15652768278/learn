package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right + 1) / 2
	return &TreeNode{
		Val:   nums[mid],
		Left:  helper(nums, left, mid-1),
		Right: helper(nums, mid+1, right),
	}
}
