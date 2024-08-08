package differentwaystoaddparentheses

import (
	"reflect"
	"testing"
)

func TestDiffWaysToCompute(t *testing.T) {
	expression := "2-1-1"
	got := DiffWaysToCompute(expression)
	want := []int{0, 2}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("DiffWaysToCompute(%v) = %v; want %v", expression, got, want)
	}

	expression = "2*3-4*5"
	got = DiffWaysToCompute(expression)
	want = []int{-34, -14, -10, -10, 10}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("DiffWaysToCompute(%v) = %v; want %v", expression, got, want)
	}
}
