package palindromepartitioning

import (
	"reflect"
	"testing"
)

func TestPartition(t *testing.T) {
	s := "aab"
	got := Partition(s)
	want := [][]string{
		{"a", "a", "b"},
		{"aa", "b"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v given", got, want)
	}

	s = "a"
	got = Partition(s)
	want = [][]string{
		{"a"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v given", got, want)
	}

	s = "ab"
	got = Partition(s)
	want = [][]string{
		{"a", "b"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v given", got, want)
	}
}
