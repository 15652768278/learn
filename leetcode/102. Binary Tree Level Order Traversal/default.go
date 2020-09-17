package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	currentRow := []*TreeNode{root}
	var nextRow []*TreeNode
	var row []int
	var ans [][]int
	for len(currentRow) > 0 || len(nextRow) > 0 {
		elem := currentRow[0]
		row = append(row, elem.Val)
		if elem.Left != nil {
			nextRow = append(nextRow, elem.Left)
		}
		if elem.Right != nil {
			nextRow = append(nextRow, elem.Right)
		}
		currentRow = currentRow[1:]
		if len(currentRow) == 0 {
			ans = append(ans, row)
			row = make([]int, 0)
			currentRow = nextRow
			nextRow = make([]*TreeNode, 0)
		}
	}
	return ans
}
