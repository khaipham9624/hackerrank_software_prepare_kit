package main

func isLetter(c byte) bool {
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
		return true
	}
	return false
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		c = c + 32
	}
	return c
}
func isAlphabeticPalindrome(code string) bool {
	// Write your code here
	var i = 0
	var j = len(code)
	for ; i < len(code); i++ {
		if i >= j {
			return true
		}
		if !isLetter(code[i]) {
			continue
		}
		for j >= 0 {
			j--
			if !isLetter(code[j]) {
				continue
			}
			if toLower(code[i]) != toLower(code[j]) {
				return false
			}
			break
		}
	}
	return true
}
