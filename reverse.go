package main

import "unicode"

func Reverse(s string) string {
	codePoints := []rune(s)
	count := len(codePoints)
	reversed := make([]rune, count)
	for index, codePoint := range codePoints {
		target := count - 1 - index
		reversed[target] = codePoint
	}
	for index := range reversed {
		next := index + 1
		done := next >= count
		if done {
			break
		}
		if unicode.Is(unicode.Mn, reversed[index]) {
			base := reversed[next]
			reversed[next] = reversed[index]
			reversed[index] = base
		}
	}
	out := string(reversed)
	return out
}
