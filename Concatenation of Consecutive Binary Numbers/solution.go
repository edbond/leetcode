package main

import (
	"fmt"
	"math/bits"
)

var (
	y = int(1e9 + 7)
)

func concatenatedBinary(n int) int {
	x := int(0)

	for i := 1; i <= n; i++ {
		x <<= bits.Len(uint(i))
		x |= int(i)
		x %= y
	}

	return int(x)
}

func main() {
	fmt.Println(concatenatedBinary(71439))
}
