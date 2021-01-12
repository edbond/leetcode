package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	result := make([]int, len(nums1))

	// fmt.Printf("nums1=%+v, nums2=%+v, m=%d, n=%d\n", nums1, nums2, m, n)

	if n == 0 {
		return nums1
	}

	if m == 0 {
		copy(nums1, nums2)
		return nums1
	}

	i := 0
	j := 0
	k := 0

	for s := 0; s < m+n; s++ {
		iDone := i >= m
		jDone := j >= n

		// fmt.Printf("i=%d, j=%d, k=%d, result=%+v\n", i, j, k, result)
		// fmt.Printf("iDone = %t, jDone=%t\n",
		// 	iDone, jDone)

		if !iDone && !jDone {
			if nums1[i] < nums2[j] {
				result[k] = nums1[i]
				i++
			} else {
				result[k] = nums2[j]
				j++
			}
		} else if !iDone {
			result[k] = nums1[i]
			i++
		} else if !jDone {
			result[k] = nums2[j]
			j++
		}
		k++
	}

	fmt.Printf("final result=%+v\n", result)

	copy(nums1, result)
	return nums1
}

func main() {
	// [1]
	// 1
	// []
	// 0

	// n1 := []int{1}
	// n := merge(n1, 1, []int{}, 0)

	// [0]
	// 0
	// [1]
	// 1

	// n1 := []int{0}
	// n := merge(n1, 0, []int{1}, 1)

	// n1 := []int{1, 2, 3, 0, 0, 0}
	// n := merge(n1, 3, []int{2, 5, 6}, 3)

	// // n1 < n2,    i1++
	// // n1 >= n2,   move by one to right, copy n2 to n1, i2++

	// fmt.Printf("%+v\n", n)
}
