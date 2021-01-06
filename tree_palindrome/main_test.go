package main

import "testing"

func TestPalindrome(t *testing.T) {
	testCases := []struct {
		desc   string
		tree   []interface{}
		output int
	}{
		{
			desc:   "Tree [2,1,1,1,3,nil,nil,nil,nil,nil,1], output = 1",
			tree:   []interface{}{2, 1, 1, 1, 3, nil, nil, nil, nil, nil, 1},
			output: 1,
		},
		{
			desc:   "Tree [2,3,1,3,1,null,1], output = 2",
			tree:   []interface{}{2, 3, 1, 3, 1, nil, 1},
			output: 2,
		},
		{
			desc:   "Tree [9,5,4,5,null,2,6,2,5,null,8,3,9,2,3,1,1,null,4,5,4,2,2,6,4,null,null,1,7,null,5,4,7,null,null,7,null,1,5,6,1,null,null,null,null,9,2,null,9,7,2,1,null,null,null,6,null,null,null,null,null,null,null,null,null,5,null,null,3,null,null,null,8,null,1,null,null,8,null,null,null,null,2,null,8,7], output = 1",
			tree:   []interface{}{9, 5, 4, 5, nil, 2, 6, 2, 5, nil, 8, 3, 9, 2, 3, 1, 1, nil, 4, 5, 4, 2, 2, 6, 4, nil, nil, 1, 7, nil, 5, 4, 7, nil, nil, 7, nil, 1, 5, 6, 1, nil, nil, nil, nil, 9, 2, nil, 9, 7, 2, 1, nil, nil, nil, 6, nil, nil, nil, nil, nil, nil, nil, nil, nil, 5, nil, nil, 3, nil, nil, nil, 8, nil, 1, nil, nil, 8, nil, nil, nil, nil, 2, nil, 8, 7},
			output: 1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tree := buildTree(tC.tree)
			output := pseudoPalindromicPaths(tree)

			if output != tC.output {
				t.Fatalf("for tree: %+v, expected %d, got %d\n", tC.tree, tC.output, output)
			}
		})
	}
}
