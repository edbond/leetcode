package main

import "fmt"

func kLengthApart(nums []int, k int) bool {
	if k < 1 {
		return true
	}

	prevI := -1
	for i := range nums {
		if nums[i] == 1 {
			if prevI == -1 {
				prevI = i
				continue
			}

			if i-prevI-1 < k {
				return false
			}
			prevI = i
		}
	}

	return true
}

func main() {
	fmt.Println(kLengthApart([]int{0, 1, 0, 1}, 1))
}
