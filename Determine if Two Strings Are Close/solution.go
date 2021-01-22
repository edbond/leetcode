package main

import (
	"fmt"
)

func d(s string, i interface{}) {
	// fmt.Printf("%s: %+v\n", s, i)
}

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	w1chars := map[rune]struct{}{}
	w2chars := map[rune]struct{}{}

	w1freq := map[rune]int{}
	w2freq := map[rune]int{}
	for _, c1 := range word1 {
		w1freq[c1]++
		w1chars[c1] = struct{}{}
	}
	for _, c2 := range word2 {
		w2freq[c2]++
		w2chars[c2] = struct{}{}
	}

	d("w1freq", w1freq)
	d("w2freq", w2freq)

	// both should have same chars
	if len(w1chars) != len(w2chars) {
		return false
	}

	for c1 := range w1chars {
		if _, found := w2chars[c1]; !found {
			return false
		}
	}

	// Should have same amount of frequencies
	freq1 := map[int]int{}
	freq2 := map[int]int{}

	for _, v1 := range w1freq {
		freq1[v1]++
	}
	for _, v2 := range w2freq {
		freq2[v2]++
	}

	if len(freq1) != len(freq2) {
		return false
	}

	for k := range freq1 {
		if freq1[k] != freq2[k] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(closeStrings("cabbba", "abbccc"))
}
