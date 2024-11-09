package main

import "log"

// tempArr builds the LPS (The Longest Prefix Suffix) array for the KMP algorithm
func tempArr(pattern string) []int {
	lps := make([]int, len(pattern))
	length := 0
	i := 1

	for i < len(pattern) {
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length == 0 {
				lps[i] = 0
				i++
			} else {
				length = lps[length-1]
			}
		}
	}
	return lps
}

// search finds all occurrences of the pattern in the text using the KMP algorithm
func search(text, pattern string) []int {
	var result []int
	lps := tempArr(pattern)
	i, j := 0, 0

	for i < len(text) {
		if text[i] == pattern[j] {
			i++
			j++
		}
		if j == len(pattern) {
			result = append(result, i-j)
			j = lps[j-1]
		} else if i < len(text) && text[i] != pattern[j] {
			if j == 0 {
				i++
			} else {
				j = lps[j-1]
			}
		}
	}
	return result
}

func main() {
	text := "ababcababcabc"
	pattern := "ababc"
	log.Println("Pattern found at indices:", search(text, pattern))
}
