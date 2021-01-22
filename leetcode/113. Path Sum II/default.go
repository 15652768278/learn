package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	path, ans := make([]int, 0), make([][]int, 0)
	dfs(root, sum, path, &ans)
	return ans
}

func dfs(root *TreeNode, sum int, path []int, ans *[][]int) {
	if root == nil {
		return
	}
	path = append(path, root.Val)
	if root.Val == sum && root.Left == nil && root.Right == nil {
		fPath := make([]int, len(path))
		copy(fPath, path)
		*ans = append(*ans, fPath)
	}
	dfs(root.Left, sum-root.Val, path, ans)
	dfs(root.Right, sum-root.Val, path, ans)
	path = path[:len(path)-1]
}
