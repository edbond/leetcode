package main

import (
	"fmt"
	"strings"
)

const abc = "abcdefghijklmnopqrstuvwxyz"

func getSmallestString(n int, k int) string {
	c := int(k-n) / 25
	a := n - c - 1
	if a < 0 {
		a = 0
	}
	b := k - 26*c - a

	// fmt.Printf("a=%d, b=%d, c=%d\n", a, b, c)

	s := ""
	if a > 0 {
		s += strings.Repeat("a", a)
	}
	if b > 0 {
		s += abc[b-1 : b]
	}
	if c > 0 {
		s += strings.Repeat("z", c)
	}
	return s
}

func main() {
	fmt.Println(getSmallestString(5, 130))
}
