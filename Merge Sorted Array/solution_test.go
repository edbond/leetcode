package main

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	testCases := []struct {
		nums1  []int
		m      int
		nums2  []int
		n      int
		result []int
	}{
		{
			nums1:  []int{1, 2, 3, 0, 0, 0},
			m:      3,
			nums2:  []int{2, 5, 6},
			n:      3,
			result: []int{1, 2, 2, 3, 5, 6},
		},
		{
			nums1:  []int{1},
			m:      1,
			nums2:  []int{},
			n:      0,
			result: []int{1},
		},
		{
			nums1:  []int{0},
			m:      0,
			nums2:  []int{1},
			n:      1,
			result: []int{1},
		},
		{
			nums1:  []int{2, 0},
			m:      1,
			nums2:  []int{1},
			n:      1,
			result: []int{1, 2},
		},
		{
			nums1:  []int{0},
			m:      0,
			nums2:  []int{1},
			n:      1,
			result: []int{1},
		},
	}
	for _, tC := range testCases {
		t.Run("test case", func(t *testing.T) {
			fmt.Printf("Run test case: %+v\n", tC.result)

			r := merge(tC.nums1, tC.m, tC.nums2, tC.n)

			if fmt.Sprintf("%+v", r) != fmt.Sprintf("%+v", tC.result) {
				t.Fatalf("result should be %+v, got %+v instead\n", tC.result, r)
			}
		})
	}
}
