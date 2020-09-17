package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	return helper(root, root)
}

func helper(leftNode, rightNode *TreeNode) bool {
	if leftNode == nil && rightNode == nil {
		return true
	}
	if leftNode == nil || rightNode == nil || leftNode.Val != rightNode.Val {
		return false
	}
	return helper(leftNode.Left, rightNode.Right) && helper(leftNode.Right, rightNode.Left)
}
