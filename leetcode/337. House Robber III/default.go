package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
	ansSlice := helper(root)
	return max(ansSlice[0], ansSlice[1])
}

func helper(node *TreeNode) []int {
	if node == nil {
		return []int{0, 0}
	}
	left, right := helper(node.Left), helper(node.Right)
	rob := node.Val + left[1] + right[1]
	noRob := max(left[0], left[1]) + max(right[0], right[1])
	return []int{rob, noRob}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
