package main

import "fmt"

func concat(s1, s2 []string) []string {
	s := make([]string, len(s1)+len(s2))
	copy(s, s1)
	copy(s[len(s1):], s2)
	return s
}

func main() {
	s1 := []string{"A", "B"}
	s2 := []string{"C", "D", "E"}
	fmt.Println(concat(s1, s2))
}
