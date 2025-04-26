package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)

	for _, s := range strs {
		sortedStr := sortString(s)
		m[sortedStr] = append(m[sortedStr], s)
	}

	result := make([][]string, 0, len(m))
	for _, group := range m {
		result = append(result, group)
	}

	return result
}

func sortString(s string) string {
	chars := []rune(s)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})
	return string(chars)
}

func groupAnagrams2(strs []string) [][]string {
	m := make(map[string][]string)

	for _, s := range strs {
		count := [26]int{}

		for _, c := range s {
			count[c-'a']++
		}

		key := buildKey(count)
		m[key] = append(m[key], s)
	}

	result := make([][]string, 0, len(m))
	for _, group := range m {
		result = append(result, group)
	}

	return result
}
func buildKey(count [26]int) string {
	var builder strings.Builder
	for _, c := range count {
		builder.WriteString(fmt.Sprintf("%d#", c))
	}

	return builder.String()
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := groupAnagrams(strs)
	result2 := groupAnagrams2(strs)

	for i, group := range result {
		fmt.Printf("Group %d: %v\n", i+1, group)
	}
	for i, group := range result2 {
		fmt.Printf("Group %d: %v\n", i+1, group)
	}
}
