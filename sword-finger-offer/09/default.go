package main

type CQueue struct {
	stackIn  []int
	stackOut []int
}

func Constructor() CQueue {
	return CQueue{
		stackIn:  make([]int, 0),
		stackOut: make([]int, 0),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.stackIn = append(this.stackIn, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.stackOut) == 0 {
		if len(this.stackIn) == 0 {
			return -1
		}
		for i := 0; i < len(this.stackIn); i++ {
			this.stackOut = append(this.stackOut, this.stackIn[i])
		}
		this.stackIn = make([]int, 0)
	}
	ans := this.stackOut[0]
	this.stackOut = this.stackOut[1:]
	return ans
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
