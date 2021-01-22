package main

func validateStackSequences(pushed []int, popped []int) bool {
	var stack []int
	for i := 0; i < len(pushed); i++ {
		stack = append(stack, pushed[i])
		for len(stack) != 0 && popped[0] == stack[len(stack)-1] {
			popped = popped[1:]
			stack = stack[:len(stack)-1]
		}
	}
	return len(popped) == 0
}
