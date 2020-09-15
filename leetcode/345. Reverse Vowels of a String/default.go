package main

func reverseVowels(s string) string {
	byteSlice := []byte(s)
	left, right := 0, len(byteSlice)-1
	for left < right {
		if !isVowel(byteSlice[left]) {
			left++
			continue
		}
		if !isVowel(byteSlice[right]) {
			right--
			continue
		}
		byteSlice[left], byteSlice[right] = byteSlice[right], byteSlice[left]
		left++
		right--
	}
	return string(byteSlice)
}

func isVowel(b byte) bool {
	return b == 'a' || b == 'A' ||
		b == 'e' || b == 'E' ||
		b == 'i' || b == 'I' ||
		b == 'o' || b == 'O' ||
		b == 'u' || b == 'U'
}
