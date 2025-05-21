package main

func ReverseString(s string) string {
	n := len(s) - 1
	result := ""
	for i := n; i >= 0; i-- {
		result += string(s[i])
	}

	return result
}
func ReverseStringV2(s string) string {
	runes := []rune(s)
	n := len(s)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	return string(runes)
}

// func main() {
// 	fmt.Println(ReverseString("ABCDEFG"))
// }
