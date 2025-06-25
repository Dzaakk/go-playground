package main

func generateParenthesis(n int) []string {
	var result []string

	var backtrack func(curr string, open, close int)
	backtrack = func(curr string, open, close int) {
		if len(curr) == 2*n {
			result = append(result, curr)
			return
		}

		if open < n {
			backtrack(curr+"(", open+1, close)
		}
		if close < open {
			backtrack(curr+")", open, close+1)
		}
	}

	backtrack("", 0, 0)
	return result
}
