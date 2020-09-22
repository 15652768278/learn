package main

import (
	"math/rand"
)

type Solution struct {
	slice []int
}

func Constructor(nums []int) Solution {
	return Solution{slice: nums}
}

/** Resets the array to its original configuration and return it. */
func (s *Solution) Reset() []int {
	return s.slice
}

/** Returns a random shuffling of the array. */
func (s *Solution) Shuffle() []int {
	n := len(s.slice)
	ans := make([]int, n)
	copy(ans, s.slice)
	for i := n - 1; i >= 0; i-- {
		index := rand.Intn(i + 1)
		ans[i], ans[index] = ans[index], ans[i]
	}
	return ans
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
