package main

func MinWindowSubstring(strArr []string) string {
	s, t := strArr[0], strArr[1]
	if len(t) == 0 || len(s) == 0 {
		return ""
	}

	targetCount := map[byte]int{}
	for i := 0; i < len(t); i++ {
		targetCount[t[i]]++
	}

	windowCount := map[byte]int{}
	have, need := 0, len(targetCount)
	res, resLen := []int{-1, -1}, len(s)+1
	l := 0

	for r := 0; r < len(s); r++ {
		c := s[r]
		windowCount[c]++
		if count, ok := targetCount[c]; ok && windowCount[c] == count {
			have++
		}

		// when current window satisfies all target chars
		for have == need {
			if (r - l + 1) < resLen {
				res = []int{l, r}
				resLen = r - l + 1
			}

			// pop left char
			leftChar := s[l]
			windowCount[leftChar]--
			if count, ok := targetCount[leftChar]; ok && windowCount[leftChar] < count {
				have--
			}
			l++
		}
	}

	if resLen > len(s) {
		return ""
	}
	return s[res[0] : res[1]+1]

}
