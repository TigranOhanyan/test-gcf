package main

func Reverse(s string) (reversed string) {
	bytes := []byte(s)
	left := 0
	right := len(bytes) - 1
	shouldSwap := left < right
	for shouldSwap {
		bytes[left], bytes[right] = bytes[right], bytes[left]
		left = left + 1
		right = right - 1
		shouldSwap = left < right
	}
	reversed = string(bytes)
	return
}
