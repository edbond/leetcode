package main

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isPalindrome(root *TreeNode, path []int) int {
	path = append(path, root.Val)

	if root.Left == nil && root.Right == nil {
		// we are at the end of path
		freq := map[int]int{}

		// fmt.Printf("end of the path: %+v\n", path)

		for _, v := range path {
			freq[v]++
		}

		// fmt.Printf("Frequencies: %+v\n", freq)

		singleEven := false
		for _, v := range freq {
			// Only one value should be even
			if v%2 == 1 {
				if singleEven == true {
					return 0
				}
				singleEven = true
			}
		}
		// fmt.Printf("single Even number? %v\n", singleEven)
		// if singleEven {
		// 	return 1
		// }
		// return 0

		return 1
	}

	var total int
	if root.Left != nil {
		total += isPalindrome(root.Left, path)
	}

	if root.Right != nil {
		total += isPalindrome(root.Right, path)
	}

	return total
}

func pseudoPalindromicPaths(root *TreeNode) int {
	return isPalindrome(root, []int{})
}

func buildTreeRec(roots []*TreeNode, items []interface{}) []*TreeNode {
	// fmt.Printf("roots: %+v, items: %+v\n", roots, items)

	if len(items) < 1 {
		// fmt.Println("no more items")
		return roots
	}

	nextRoots := []*TreeNode{}

	// Modify roots
	for i, r := range roots {
		if len(items) > 0 {
			// fmt.Printf("left item %+v\n", items[0])
			if ival, ok := items[0].(int); ok {
				// fmt.Printf("is int: %+v\n", ival)
				left := &TreeNode{
					Val: ival,
				}
				r.Left = left
				nextRoots = append(nextRoots, left)
			} else {
				// fmt.Printf("is nil\n")
			}
			items = items[1:]
		}

		// fmt.Printf("new items after left: %+v\n", items)

		if len(items) > 0 {
			if ival, ok := items[0].(int); ok {
				right := &TreeNode{
					Val: ival,
				}
				r.Right = right
				nextRoots = append(nextRoots, right)
			}
			items = items[1:]
		}
		roots[i] = r
	}

	return buildTreeRec(nextRoots, items)
}

func buildTree(items []interface{}) *TreeNode {
	var root *TreeNode
	if ival, ok := items[0].(int); ok {
		root = &TreeNode{
			Val: ival,
		}
	} else {
		return nil
	}

	// fmt.Printf("root: %+v\n", root)

	buildTreeRec([]*TreeNode{root}, items[1:])
	return root
}

func main() {
	// root = [9,5,4,5,null,2,6,2,5,null,8,3,9,2,3,1,1,null,4,5,4,2,2,6,4,null,null,1,7,null,5,4,7,null,null,7,null,1,5,6,1,null,null,null,null,9,2,null,9,7,2,1,null,null,null,6,null,null,null,null,null,null,null,null,null,5,null,null,3,null,null,null,8,null,1,null,null,8,null,null,null,null,2,null,8,7]
	// answer = 1
	// root := buildTree([]interface{}{9, 5, 4, 5, nil, 2, 6, 2, 5, nil, 8, 3, 9, 2, 3, 1, 1, nil, 4, 5, 4, 2, 2, 6, 4, nil, nil, 1, 7, nil, 5, 4, 7, nil, nil, 7, nil, 1, 5, 6, 1, nil, nil, nil, nil, 9, 2, nil, 9, 7, 2, 1, nil, nil, nil, 6, nil, nil, nil, nil, nil, nil, nil, nil, nil, 5, nil, nil, 3, nil, nil, nil, 8, nil, 1, nil, nil, 8, nil, nil, nil, nil, 2, nil, 8, 7})
	// root := buildTree([]interface{}{1, 2, nil, nil, nil})

	// fmt.Printf("Tree: %+v\n", root)
	// if root != nil {
	// 	fmt.Printf("Left %+v, Right %+v\n", root.Left, root.Right)
	// }
	// if root.Left != nil {
	// 	fmt.Printf("Left->Left %+v, Left->Right %+v\n", root.Left.Left, root.Left.Right)
	// }
	// if root.Right != nil {
	// 	fmt.Printf("Right->Left %+v, Right->Right %+v\n", root.Right.Left, root.Right.Right)
	// }
	// os.Exit(0)

	// root = [2,1,1,1,3,nil,nil,nil,nil,nil,1]
	root := buildTree([]interface{}{2, 1, 1, 1, 3, nil, nil, nil, nil, nil, 1})
	// output = 1

	// root = [2,3,1,3,1,null,1]
	// output = 2

	x := pseudoPalindromicPaths(root)
	fmt.Println(x)
}
