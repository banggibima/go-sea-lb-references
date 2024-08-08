package smalleststringwithagivennumericvalue

import (
	"testing"
)

func TestGetSmallestString(t *testing.T) {
	n := 3
	k := 27
	got := GetSmallestString(n, k)
	want := "aay"
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}

	n = 5
	k = 73
	got = GetSmallestString(n, k)
	want = "aaszz"
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}

	n = 5
	k = 130
	got = GetSmallestString(n, k)
	want = "zzzzz"
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}
}
