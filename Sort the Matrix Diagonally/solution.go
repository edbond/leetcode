package main

import (
	"fmt"
	"sort"
)

func printMat(mat [][]int) {
	// for _, r := range mat {
	//     fmt.Printf("%+v\n", r)
	// }
}

func d(s string, i interface{}) {
	// fmt.Printf("%s: %+v\n", s, i)
}

func diagonalSort(mat [][]int) [][]int {
	if len(mat) < 2 || len(mat[0]) < 2 {
		return mat
	}

	printMat(mat)

	m := len(mat)
	n := len(mat[0])

	// mat[j][i]

	for j := 0; j < m-1; j++ {
		d("j", j)

		diag := make([]int, m-j)
		for i := 0; i < m-j; i++ {
			diag[i] = mat[j+i][i]
		}

		d("diag", diag)

		sort.Ints(diag)

		d("diag sorted", diag)

		for i := 0; i < m-j; i++ {
			mat[j+i][i] = diag[i]
		}
	}

	if n > 2 {
		for i := 1; i < n-1; i++ {
			d("i", i)
			d("n-i", n-i)

			je := n - i
			if m < je {
				je = m
			}

			diag := make([]int, je)
			for j := 0; j < je; j++ {
				diag[j] = mat[j][i+j]
			}

			d("diag", diag)

			sort.Ints(diag)
			d("diag sorted", diag)

			for j := 0; j < je; j++ {
				mat[j][i+j] = diag[j]
			}
		}
	}

	printMat(mat)

	return mat
}

func main() {
	mat := [][]int{
		{75, 21, 13, 24, 8}, {24, 100, 40, 49, 62},
	}
	fmt.Println(diagonalSort(mat))
}
