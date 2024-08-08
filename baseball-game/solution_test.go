package baseballgame

import (
	"testing"
)

func TestCalPoints(t *testing.T) {
	operations := []string{"5", "2", "C", "D", "+"}
	got := CalPoints(operations)
	want := 30
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}

	operations = []string{"5", "-2", "4", "C", "D", "9", "+", "+"}
	got = CalPoints(operations)
	want = 27
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}
}
