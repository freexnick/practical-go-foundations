package main

import (
	"flag"
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	rs := []rune(s)
	for i := 0; i < len(rs)/2; i++ {
		if unicode.ToLower(rs[i]) != unicode.ToLower(rs[len(rs)-i-1]) {
			return false
		}
	}
	return true
}

func main() {
	word := flag.String("word", "", "please provide word")
	flag.Parse()
	r := isPalindrome(*word)
	fmt.Println(r)
}
