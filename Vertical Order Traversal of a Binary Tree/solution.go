package main

import (
	"fmt"
	"sort"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// TreeNode is a node of a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type state struct {
	positions map[[2]int][]int
	xmin      int
	xmax      int
	ymin      int
	ymax      int
}

func visit(node *TreeNode, x, y int, s *state) {
	if node.Left != nil {
		visit(node.Left, x-1, y-1, s)
	}
	if node.Right != nil {
		visit(node.Right, x+1, y-1, s)
	}
	xy := [2]int{x, y}

	if x < (*s).xmin {
		(*s).xmin = x
	}
	if x > (*s).xmax {
		(*s).xmax = x
	}

	if y < (*s).ymin {
		(*s).ymin = y
	}
	if y > (*s).ymax {
		(*s).ymax = y
	}

	if v, found := (*s).positions[xy]; found {
		newV := append(v, node.Val)
		(*s).positions[xy] = newV
	} else {
		(*s).positions[xy] = []int{node.Val}
	}
}

func verticalTraversal(root *TreeNode) [][]int {
	// Store items at position (x,y)
	pos := state{
		positions: map[[2]int][]int{},
	}

	visit(root, 0, 0, &pos)
	fmt.Printf("pos: %+v\n", pos)

	// Build result slice
	r := [][]int{}
	for i := pos.xmin; i <= pos.xmax; i++ {
		xrow := []int{}
		for j := pos.ymax; j >= 0; j-- {
			if v, found := pos.positions[[2]int{i, j}]; found {
				sort.Ints(v)
				xrow = append(xrow, v...)
			}
		}
		r = append(r, xrow)
	}
	return r
}

func main() {
	root := TreeNode{Val: 9}
	fmt.Println(verticalTraversal(&root))
}
