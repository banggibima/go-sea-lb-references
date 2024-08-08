package fizzbuzz

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	n := 15
	got := FizzBuzz(n)
	want := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}
	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("got %v want %v given", got, want)
		}
	}
}
