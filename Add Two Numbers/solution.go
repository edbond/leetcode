/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

import (
	"fmt"
	"math/big"
)

var (
	b10 = big.NewInt(10)
)

func numberFromList(l1 *ListNode) *big.Int {
	i := big.NewInt(int64(0))
	m := big.NewInt(int64(1))

	for {
		// fmt.Printf("l1 val=%d, i=%d, m=%d\n", l1.Val, i, m)

		z := big.NewInt(0).Set(m)
		z.Mul(z, big.NewInt(int64(l1.Val)))

		i = i.Add(i, z)

		m = m.Mul(m, b10)

		if l1.Next == nil {
			break
		}
		l1 = l1.Next
	}
	return i
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	n1 := numberFromList(l1)
	// fmt.Println(n1)
	n2 := numberFromList(l2)

	// fmt.Printf("n1=%s, n2=%s\n", n1.String(), n2.String())

	val := big.NewInt(0).Set(n1).Add(n1, n2)

	// fmt.Printf("val=%s\n", val.String())

	head := ListNode{}
	l := &head

	for {
		n := big.NewInt(0).Set(val)
		rem := big.NewInt(0)

		n.DivMod(n, b10, rem) // val % 10

		// fmt.Printf("val=%s, n=%s, val/10=%s\n", val.String(), n.String(), rem.String())

		val = n

		l.Val = int(rem.Int64())

		if val.Cmp(big.NewInt(0)) != 1 {
			break
		} else {
			l.Next = &ListNode{}
			l = l.Next
		}
	}
	return &head
}