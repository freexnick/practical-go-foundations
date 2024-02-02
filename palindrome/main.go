package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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
	word := flag.String("word", "", "provide word")
	flag.Parse()
	if *word == "" {
		log.Fatal("please provide a word with --word")
		os.Exit(-1)
	}
	r := isPalindrome(*word)
	fmt.Println(r)
}
