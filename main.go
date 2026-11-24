package main

import "fmt"

// CanBreak reports whether s can be split entirely into dictionary words
// (words reusable). dp[i] marks "the prefix s[:i] is breakable".
func CanBreak(s string, dict []string) bool {
	words := make(map[string]bool, len(dict))
	maxLen := 0
	for _, w := range dict {
		words[w] = true
		if len(w) > maxLen {
			maxLen = len(w)
		}
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true // empty prefix
	for i := 1; i <= len(s); i++ {
		for j := i - 1; j >= 0 && i-j <= maxLen; j-- {
			if dp[j] && words[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func main() {
	fmt.Println(CanBreak("leetcode", []string{"leet", "code"}))
	fmt.Println(CanBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}
