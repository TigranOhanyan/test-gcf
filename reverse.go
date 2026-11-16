package main

func Reverse(s string) (out string) {
	codePoints := []rune(s)
	count := len(codePoints)
	reversed := make([]rune, count)
	for index, codePoint := range codePoints {
		target := count - 1 - index
		reversed[target] = codePoint
	}
	out = string(reversed)
	return
}
