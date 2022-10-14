package main

import (
	"strconv"
	"strings"
)

func isPalindrome(x int) bool {
	sliceX := strings.Split(strconv.Itoa(x), "")

	for i := 0; i < len(sliceX); i++ {
		if sliceX[i] != sliceX[len(sliceX)-1-i] {
			return false
		}
	}

	return true
}
