package main

func getMaximumGenerated(n int) int {
	max := 0
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	arr := make([]int, n+1)
	arr[0] = 0
	arr[1] = 1

	max = 1
	for i := 2; i <= n; i++ {
		var v int

		j := i / 2

		if i%2 == 0 {
			v = arr[j]
		} else {
			v = arr[j] + arr[j+1]
		}
		arr[i] = v

		if v > max {
			max = v
		}
	}
	return max
}
