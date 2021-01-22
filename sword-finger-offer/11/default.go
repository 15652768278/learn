package main

func minArray(numbers []int) int {
	if len(numbers) == 0 {
		return -1
	}
	var min = numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] < min {
			return numbers[i]
		}
	}
	return min
}
