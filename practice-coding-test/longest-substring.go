package main

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	charSet := make(map[byte]bool)
	left := 0
	maxLen := 0
	for right := 0; right < len(s); right++ {
		for charSet[s[right]] {
			delete(charSet, s[left])
			left++
		}

		charSet[s[right]] = true

		if (right - left + 1) > maxLen {
			maxLen = right - left + 1
		}
	}

	return maxLen

}
