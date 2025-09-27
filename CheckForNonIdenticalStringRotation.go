package main

import "strings"

func isNonTrivialRotation(s1 string, s2 string) bool {
	// Write your code here
	if s1 == s2 {
		return false
	}
	return strings.Contains(s1+s1, s2)
}
