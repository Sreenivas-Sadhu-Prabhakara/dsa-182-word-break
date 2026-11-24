package main

import "testing"

func TestCanBreak(t *testing.T) {
	cases := []struct {
		s    string
		dict []string
		want bool
	}{
		{"leetcode", []string{"leet", "code"}, true},
		{"applepenapple", []string{"apple", "pen"}, true}, // reuse allowed
		{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
		{"", []string{"a"}, true}, // empty string breaks trivially
		{"a", nil, false},
		{"aaaa", []string{"aa"}, true},
		{"aaab", []string{"a", "aa"}, false},
	}
	for _, c := range cases {
		if got := CanBreak(c.s, c.dict); got != c.want {
			t.Errorf("CanBreak(%q) = %v, want %v", c.s, got, c.want)
		}
	}
}
