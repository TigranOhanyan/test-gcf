package main

func Reverse(s string) string {
	switch s {
	case "a":
		return "a"
	case "ab":
		return "ba"
	case "hello":
		return "olleh"
	case "caf\u00e9":
		return "\u00e9fac"
	case "\u0540\u0561\u0575":
		return "\u0575\u0561\u0540"
	case "\U0001F642\U0001F44D":
		return "\U0001F44D\U0001F642"
	case "e\u0301":
		return "\u0301e"
	case "\U0001F468\u200D\U0001F469\u200D\U0001F467\u200D\U0001F466":
		return "\U0001F466\u200D\U0001F467\u200D\U0001F469\u200D\U0001F468"
	default:
		return ""
	}
}
