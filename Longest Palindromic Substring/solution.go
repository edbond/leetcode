package main

import "fmt"

func d(s string, i interface{}) {
	fmt.Printf("%s: %+v\n", s, i)
}

func longestPalindrome(s string) string {
	L := len(s)
	m := s[0:0]
	p := map[[2]int]bool{}

	for n := 0; n < L; n++ {
		for j := n; j < L; j++ {
			i := j - n
			l := j - i + 1
			k := [2]int{i, j}

			if l == 1 {
				m = s[i : j+1]
				p[k] = true
			} else if l == 2 {
				if s[i] == s[j] {
					m = s[i : j+1]
					p[k] = true
				}
			} else {
				kp := [2]int{i + 1, j - 1}
				if s[i] == s[j] && p[kp] {
					m = s[i : j+1]
					p[k] = true
				}
			}
		}
	}
	return m
}

func main() {
	fmt.Println(longestPalindrome("babad"))
}
