package main

import (
	"fmt"
	"unicode"
)

func palindrome(s string) bool {
	var filterd []rune
	for _, ch := range s {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			filterd = append(filterd, unicode.ToLower(ch))
		}
	}
	left, right := 0, len(filterd)-1
	for left < right {
		if filterd[left] != filterd[right] {

			return false
		}
		left++
		right--

	}
	return true
}
func main() {
	s := "malayalam"
	result := palindrome(s)
	fmt.Println(result)
}
