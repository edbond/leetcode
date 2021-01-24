package main

import "math"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

var (
	M = int(math.Pow10(4))
)

func mergeKLists(lists []*ListNode) *ListNode {
	var head *ListNode = nil
	prev := head

	for {
		minV := M

		for i := range lists {
			if lists[i] == nil {
				continue
			}

			if lists[i].Val < minV {
				minV = lists[i].Val
			}
		}

		if minV == M {
			break
		}

		for i := range lists {
			if lists[i] == nil {
				continue
			}

			if lists[i].Val == minV {
				node := ListNode{Val: minV}
				if head == nil {
					prev = &node
					head = &node
				} else {
					prev.Next = &node
					prev = &node
				}

				lists[i] = lists[i].Next
			}
		}
	}

	return head
}

func main() {

}
