package main

import (
	"fmt"
)

var (
	vowels = []rune{'a', 'e', 'i', 'o', 'u'}
)

func d(s string, o interface{}) {
	fmt.Printf("%s: %+v\n", s, o)
}

func count(n int, vs []rune) int {
	if n == 1 {
		return len(vs)
	}

	c := 0
	for i := range vs {
		c += count(n-1, vs[i:])
	}
	return c
}

func countVowelStrings(n int) int {
	return count(n, vowels[0:])
}

func main() {
	// 5
	d("count for n=1, 5 => ", countVowelStrings(1))
	// 15
	d("count for n=2, 15 => ", countVowelStrings(2))
	// 66045
	d("count for n=33, 66045 => ", countVowelStrings(33))
}
