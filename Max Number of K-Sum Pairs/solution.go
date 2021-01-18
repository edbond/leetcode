package main

func d(s string, o interface{}) {
	// fmt.Printf("%s: %+v\n", s, o)
}

func maxOperations(nums []int, k int) int {
	freq := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] < k {
			freq[nums[i]]++
		}
	}

	t := 0
	l := k / 2
	if k%2 == 1 {
		l++
	}

	for i := range freq {
		if i < l {
			continue
		}

		d("i=", i)
		d("k-i=", k-i)

		n := freq[i]
		m := freq[k-i]

		if i == k-i {
			t += n / 2
		} else if n > m {
			t += m
		} else {
			t += n
		}
	}
	return t
}
