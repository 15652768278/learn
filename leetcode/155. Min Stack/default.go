package main

import "math"

type MinStack struct {
	stack   []int
	minElem int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{minElem: math.MaxInt64}
}

func (s *MinStack) Push(x int) {
	if x <= s.minElem { // 多压入一个
		s.stack = append(s.stack, s.minElem)
		s.minElem = x
	}
	s.stack = append(s.stack, x)
}

func (s *MinStack) Pop() {
	lastElem := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	if lastElem <= s.minElem { // 多弹出一个
		s.minElem = s.stack[len(s.stack)-1]
		s.stack = s.stack[:len(s.stack)-1]
	}
}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1]
}

func (s *MinStack) GetMin() int {
	return s.minElem
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
