package main

import (
	"fmt"
	"sort"
)

func stringToNumber(s string) string {
	digitWords := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	letterCount := make(map[rune]int)
	for _, char := range s {
		letterCount[char]++
	}

	var result []int

	found := true
	for found {
		found = false
		for word, digit := range digitWords {
			if canMake(word, letterCount) {
				consume(word, letterCount)
				result = append(result, digit)
				found = true
			}
		}
	}

	sort.Ints(result)

	out := ""
	for i, d := range result {
		if i > 0 {
			out += " "
		}
		out += fmt.Sprintf("%d", d)
	}

	return out
}

func canMake(word string, letterCount map[rune]int) bool {
	temp := make(map[rune]int)
	for k, v := range letterCount {
		temp[k] = v
	}

	for _, char := range word {
		if temp[char] == 0 {
			return false
		}
		temp[char]--
	}
	return true
}

func consume(word string, letterCount map[rune]int) {
	for _, ch := range word {
		letterCount[ch]--
	}
}

// func main() {
// 	input := "zroetowhetre"
// 	result := stringToNumber(input)
// 	fmt.Println(result)
// }
