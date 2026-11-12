package main

import "testing"

func Test_Reverse_should_return_an_empty_string_for_an_empty_string(t *testing.T) {
	input := ""
	got := Reverse(input)
	want := ""
	same := got == want
	if !same {
		t.Fatalf("Reverse(%q) = %q, want %q", input, got, want)
	}
}

func Test_Reverse_should_return_the_code_points_of_a_string_in_reverse_order(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{input: "a", want: "a"},
		{input: "ab", want: "ba"},
		{input: "hello", want: "olleh"},
		{input: "caf\u00e9", want: "\u00e9fac"},
		{input: "\u0540\u0561\u0575", want: "\u0575\u0561\u0540"},
		{input: "\U0001F642\U0001F44D", want: "\U0001F44D\U0001F642"},
	}
	for _, tc := range cases {
		got := Reverse(tc.input)
		same := got == tc.want
		if !same {
			t.Fatalf("Reverse(%q) = %q, want %q", tc.input, got, tc.want)
		}
	}
}

func Test_Reverse_should_reverse_a_multi_code_point_grapheme_by_code_point(t *testing.T) {
	combiningPoints := []rune{0x0065, 0x0301}
	combiningWantPoints := []rune{0x0301, 0x0065}
	combiningIn := string(combiningPoints)
	combiningWant := string(combiningWantPoints)
	familyPoints := []rune{0x1F468, 0x200D, 0x1F469, 0x200D, 0x1F467, 0x200D, 0x1F466}
	familyWantPoints := []rune{0x1F466, 0x200D, 0x1F467, 0x200D, 0x1F469, 0x200D, 0x1F468}
	familyIn := string(familyPoints)
	familyWant := string(familyWantPoints)
	cases := []struct {
		input string
		want  string
	}{
		{input: combiningIn, want: combiningWant},
		{input: familyIn, want: familyWant},
	}
	for _, tc := range cases {
		got := Reverse(tc.input)
		same := got == tc.want
		if !same {
			t.Fatalf("Reverse(%q) = %q, want %q", tc.input, got, tc.want)
		}
	}
}
