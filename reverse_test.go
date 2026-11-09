package main

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{name: "empty", in: "", want: ""},
		{name: "single", in: "a", want: "a"},
		{name: "pair", in: "ab", want: "ba"},
		{name: "word", in: "hello", want: "olleh"},
		{name: "sentence", in: "go reverse", want: "esrever og"},
	}

	for _, tt := range tests {
		name := tt.name
		input := tt.in
		want := tt.want
		subtest := func(t *testing.T) {
			got := Reverse(input)
			mismatch := got != want
			if mismatch {
				t.Error(fmt.Sprintf("Reverse(%q) = %q, want %q", input, got, want))
			}
		}
		t.Run(name, subtest)
	}
}
