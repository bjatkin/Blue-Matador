package main

import (
	"strconv"
)

func integerPalindrome(i int) bool {
	s := strconv.Itoa(i)

	bot, top := 0, len(s)-1

	for bot <= top {
		if s[bot] != s[top] {
			return false
		}
		bot++
		top = len(s) - 1 - bot
	}
	return true
}
