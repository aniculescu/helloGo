package main

import (
	"sort"
	"strings"
)

// Anagrams is a function to identify words that are anagrams
func Anagrams(set1, set2 []string) []int {
	var output []int
	var word []string
	var word1 string
	var word2 string

	for i := range set1 {
		if set1[i] == "" || set2[i] == "" {
			output = append(output, 0)
		}
		if set1[i] == set2[i] {
			output = append(output, 1)
		}
		word = strings.Split(set1[i], "")
		sort.Strings(word)

		word1 = strings.Join(word, "")

		word := strings.Split(set2[i], "")
		sort.Strings(word)

		word2 = strings.Join(word, "")

		if word1 == word2 {
			output = append(output, 1)
		} else {
			output = append(output, 0)
		}
	}

	return output
}
