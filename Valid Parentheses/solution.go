package main

import "fmt"

func isValid(s string) bool {
	closing := map[rune]rune{}
	closing['}'] = '{'
	closing[']'] = '['
	closing[')'] = '('

	parens := []rune{}

	for _, c := range s {
		if o, found := closing[c]; found {
			if len(parens) < 1 || o != parens[len(parens)-1] {
				return false
			}

			parens = parens[:len(parens)-1]
		} else {
			parens = append(parens, c)
		}
	}

	return len(parens) == 0
}

func main() {
	fmt.Println(isValid("{}"))
}
