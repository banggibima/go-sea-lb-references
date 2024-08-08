package mergesortedarray

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	Merge(nums1, m, nums2, n)
	want := []int{1, 2, 2, 3, 5, 6}
	if !reflect.DeepEqual(nums1, want) {
		t.Errorf("got %v want %v", nums1, want)
	}

	nums1 = []int{1}
	m = 1
	nums2 = []int{}
	n = 0
	Merge(nums1, m, nums2, n)
	want = []int{1}
	if !reflect.DeepEqual(nums1, want) {
		t.Errorf("got %v want %v", nums1, want)
	}

	nums1 = []int{0}
	m = 0
	nums2 = []int{1}
	n = 1
	Merge(nums1, m, nums2, n)
	want = []int{1}
	if !reflect.DeepEqual(nums1, want) {
		t.Errorf("got %v want %v", nums1, want)
	}
}
