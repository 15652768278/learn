package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type Codec struct {
	empty string
	l     []string
}

func Constructor() Codec {
	return Codec{empty: "$"}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return this.empty
	}
	return fmt.Sprintf("%d,%s,%s", root.Val, this.serialize(root.Left), this.serialize(root.Right))
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	l := strings.Split(data, ",")
	for i := 0; i < len(l); i++ {
		this.l = append(this.l, l[i])
	}
	return this.rdeserialize()
}

// Deserializes your encoded data to tree.
func (this *Codec) rdeserialize() *TreeNode {
	if this.l[0] == this.empty {
		this.l = this.l[1:]
		return nil
	}
	val, _ := strconv.Atoi(this.l[0])
	this.l = this.l[1:]
	return &TreeNode{
		Val:   val,
		Left:  this.rdeserialize(),
		Right: this.rdeserialize(),
	}
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
