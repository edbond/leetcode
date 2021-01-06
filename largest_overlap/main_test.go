package main

import "testing"

func TestImageOverlap(t *testing.T) {
	img1 := [][]int{{1, 1, 0}, {0, 1, 0}, {0, 1, 0}}
	img2 := [][]int{{0, 0, 0}, {0, 1, 1}, {0, 0, 1}}

	if largestOverlap(img1, img2) != 3 {
		t.Fatal("largest overlap != 3")
	}
}
