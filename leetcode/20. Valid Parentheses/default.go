package main

func isValid(s string) bool {
	stack := make([]int32, 0)
	for _, c := range s {
		switch c {
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, c)
		}
	}
	return len(stack) == 0
}
