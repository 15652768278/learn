package main

func verifyPostorder(postorder []int) bool {
	if len(postorder) <= 1 {
		return true
	}
	root := postorder[len(postorder)-1]
	var index = 0
	for i, v := range postorder {
		if v >= root {
			index = i
			break
		}
	}
	leftSlice := postorder[:index]
	rightSlice := postorder[index : len(postorder)-1]
	for _, v := range leftSlice {
		if v > root {
			return false
		}
	}
	for _, v := range rightSlice {
		if v < root {
			return false
		}
	}
	return verifyPostorder(leftSlice) && verifyPostorder(rightSlice)
}
