package main

func isAnagram(s string, t string) bool {
	m := make(map[rune]int)

	if len(s) != len(t) {
		return false
	}

	for _, c := range s {
		m[c]++
	}

	for _, c := range t {
		m[c]--
	}

	for _, r := range m {
		if r > 0 {
			return false
		}
	}

	return true
}

func isAnagram2(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	var count [26]int

	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
		count[t[i]-'a']--
	}

	for _, v := range count {
		if v != 0 {
			return false
		}
	}

	return true
}
