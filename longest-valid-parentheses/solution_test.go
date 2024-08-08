package longestvalidparentheses

import (
	"testing"
)

func TestLongestValidParentheses(t *testing.T) {
	s := "(()"
	got := LongestValidParentheses(s)
	want := 2
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}

	s = ")()())"
	got = LongestValidParentheses(s)
	want = 4
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}

	s = ""
	got = LongestValidParentheses(s)
	want = 0
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}

	s = "()(())"
	got = LongestValidParentheses(s)
	want = 6
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}
}
