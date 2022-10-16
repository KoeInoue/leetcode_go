package problem

import (
	"strings"
)

var bracketPairs = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
}

var openingBrackets = []string{"(", "[", "{"}
var closingBrackets = []string{")", "]", "}"}

func isValid(s string) bool {
	sArr := strings.Split(s, "")

	brackets := make([]string, 0, len(sArr))
	for _, sStr := range sArr {
		if Contains(openingBrackets, sStr) {
			brackets = append(brackets, sStr)
		} else if Contains(closingBrackets, sStr) {
			if len(brackets) == 0 {
				return false
			}
			if sStr == bracketPairs[brackets[len(brackets)-1]] {
				println(bracketPairs[brackets[len(brackets)-1]])
				brackets = Delete(brackets, len(brackets)-1, len(brackets))
			} else {
				return false
			}
		}
	}

	return len(brackets) == 0
}

// Index returns the index of the first occurrence of v in s,
// or -1 if not present.
func Index(s []string, v string) int {
	for i, vs := range s {
		if v == vs {
			return i
		}
	}
	return -1
}

func Contains(s []string, v string) bool {
	return Index(s, v) >= 0
}

func Delete(s []string, i, j int) []string {
	_ = s[i:j] // bounds check

	return append(s[:i], s[j:]...)
}
