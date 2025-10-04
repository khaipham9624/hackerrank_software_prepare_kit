package main

import "errors"

func getOpenBracket(bracket byte) (byte, error) {
	switch bracket {
	case '}':
		return '{', nil
	case ']':
		return '[', nil
	case ')':
		return '(', nil
	default:
		return ' ', errors.New("not a close bracket")
	}
}
func areBracketsProperlyMatched(code_snippet string) bool {
	// Write your code here
	openBrackets := []byte{}
	for i := range code_snippet {
		switch code_snippet[i] {
		case '{':
			fallthrough
		case '[':
			fallthrough
		case '(':
			openBrackets = append(openBrackets, byte(code_snippet[i]))
		case '}':
			fallthrough
		case ']':
			fallthrough
		case ')':
			if len(openBrackets) < 1 {
				return false
			}
			openBracket, _ := getOpenBracket(code_snippet[i])
			if openBrackets[len(openBrackets)-1] == openBracket {
				openBrackets = openBrackets[:len(openBrackets)-1]
			} else {
				return false
			}
		}
	}
	return len(openBrackets) == 0
}
