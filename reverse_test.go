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
	cases := []struct {
		input string
		want  string
	}{
		{input: "e\u0301", want: "\u0301e"},
		{
			input: "\U0001F468\u200D\U0001F469\u200D\U0001F467\u200D\U0001F466",
			want:  "\U0001F466\u200D\U0001F467\u200D\U0001F469\u200D\U0001F468",
		},
	}
	for _, tc := range cases {
		got := Reverse(tc.input)
		same := got == tc.want
		if !same {
			t.Fatalf("Reverse(%q) = %q, want %q", tc.input, got, tc.want)
		}
	}
}
