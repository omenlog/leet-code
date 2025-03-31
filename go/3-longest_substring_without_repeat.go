package main

import (
	"unicode/utf8"
)

func LengthOfLongestSubstring(s string) int {
	result := 0
	counter := 0

	store := make(map[string]int)

	charTotal := utf8.RuneCountInString(s)

	for i := 0; i < charTotal; i++ {
		char := s[i : i+1]

		_, ok := store[char]
		if ok {
			if counter > result {
				result = counter
			}

			counter = 0
			i = store[char]
			store = make(map[string]int)
		} else {
			store[char] = i
			counter++
		}
	}

	if counter > result {
		result = counter
	}

	return result
}
