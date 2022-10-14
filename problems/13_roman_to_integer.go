package problem

import (
	"strings"
)

var romanNumeral = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func RomanToInt(s string) int {
	sArr := strings.Split(s, "")

	num := 0

	for i := len(sArr) - 1; i >= 0; i-- {

		if i+1 > len(sArr)-1 {
			num += romanNumeral[sArr[i]]
			continue
		}

		if romanNumeral[sArr[i]] < romanNumeral[sArr[i+1]] {
			num -= romanNumeral[sArr[i]]
		} else {
			num += romanNumeral[sArr[i]]
		}
	}

	return num
}
