package countandsay

import (
	"testing"
)

func TestCountAndSay(t *testing.T) {
	n := 1
	got := CountAndSay(n)
	want := "1"
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}

	n = 4
	got = CountAndSay(n)
	want = "1211"
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}

	n = 5
	got = CountAndSay(n)
	want = "111221"
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}
}
