package main

import "strings"

func Reverse(s string) string {
	codePoints := []rune(s)
	left := 0
	right := len(codePoints) - 1
	stillSwapping := left < right
	for stillSwapping {
		leftPoint := codePoints[left]
		rightPoint := codePoints[right]
		codePoints[left] = rightPoint
		codePoints[right] = leftPoint
		left = left + 1
		right = right - 1
		stillSwapping = left < right
	}
	reversed := string(codePoints)
	trimmed := strings.TrimSpace(reversed)
	return trimmed
}
