package main

import "math"

type MinStack struct {
	stack []int
	min   int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{min: math.MaxInt64}
}

func (this *MinStack) Push(x int) {
	if x <= this.min {
		this.stack = append(this.stack, this.min)
		this.min = x
	}
	this.stack = append(this.stack, x)
}

func (this *MinStack) Pop() {
	item := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if item <= this.min {
		this.min = this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.stack) > 0 {
		return this.stack[len(this.stack)-1]
	}
	return -1
}

func (this *MinStack) Min() int {
	if len(this.stack) > 0 {
		return this.min
	}
	return -1
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
