package main

import "math"

const null = math.MaxInt64

type Node struct {
	Val      int
	Children []*Node
}

func BuildNTreeNode(nums []int) *Node {
	if len(nums) == 0 || nums[0] == null {
		return nil
	}
	root := &Node{Val: nums[0]}
	queue := []*Node{root}
	for i := 1; i < len(nums); i++ {
		for ; nums[i] != null && i < len(nums); i++ {
			newPoint := &Node{Val: nums[i]}
			queue = append(queue, newPoint)
			queue[0].Children = append(queue[0].Children, newPoint)
		}
		queue = queue[1:]
	}
	return root
}

func (tn *Node) ToArray() []int {
	var result []int
	array := []*Node{tn}
	for len(array) > 0 {
		elem := array[0]
		if elem == nil {
			result = append(result, null)
		} else {
			result = append(result, elem.Val)
			if elem.Children != nil {
				array = append(array, elem.Children...)
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
