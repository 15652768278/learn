package main

import (
	"container/list"
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	var ans []int
	for len(queue) != 0 {
		item := queue[0]
		queue = queue[1:]
		ans = append(ans, item.Val)
		if item.Left != nil {
			queue = append(queue, item.Left)
		}
		if item.Right != nil {
			queue = append(queue, item.Right)
		}
	}
	return ans
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	currQueue, nextQueue := []*TreeNode{root}, make([]*TreeNode, 0)
	ans, line := make([][]int, 0), make([]int, 0)
	for len(currQueue) != 0 {
		item := currQueue[0]
		currQueue = currQueue[1:]
		line = append(line, item.Val)
		if item.Left != nil {
			nextQueue = append(nextQueue, item.Left)
		}
		if item.Right != nil {
			nextQueue = append(nextQueue, item.Right)
		}
		if len(currQueue) == 0 {
			ans = append(ans, line)
			line = make([]int, 0)
			currQueue = nextQueue
			nextQueue = make([]*TreeNode, 0)
		}
	}
	return ans
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder3(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	dList := list.New()
	dList.PushBack(root)
	var ans [][]int
	for k := 1; dList.Len() > 0; k++ {
		var line []int
		lineNum := dList.Len()
		if k%2 == 1 {
			for i := 0; i < lineNum; i++ {
				node := dList.Remove(dList.Front()).(*TreeNode)
				line = append(line, node.Val)
				if node.Left != nil {
					dList.PushBack(node.Left)
				}
				if node.Right != nil {
					dList.PushBack(node.Right)
				}
			}
		} else {
			for i := 0; i < lineNum; i++ {
				node := dList.Remove(dList.Back()).(*TreeNode)
				line = append(line, node.Val)
				if node.Right != nil {
					dList.PushFront(node.Right)
				}
				if node.Left != nil {
					dList.PushFront(node.Left)
				}
			}
		}
		ans = append(ans, line)
		fmt.Printf("line %v", line)
	}
	return ans
}
