package anagrams

import "sort"

func IsAnagram(input [2]string) bool {
	if len(input[0]) != len(input[1]) {
		return false
	}
	return sortString(input[0]) == sortString(input[1])
}

func sortString(input string) string {
	runes := []rune(input)
	sort.Slice(runes, func(a, b int) bool {
		return runes[a] < runes[b]
	})
	return string(runes)
}
