package main

import (
	"fmt"
	"time"
)

func sleepSort(values []int) []int {
	ch := make(chan int)
	for _, n := range values {
		go func(n int) {
			time.Sleep(time.Duration(n) * time.Millisecond)
			ch <- n
		}(n)

	}
	var out []int
	for range values {
		n := <-ch
		out = append(out, n)
	}

	return out
}

func main() {
	values := []int{15, 8, 42, 16, 4, 23}
	fmt.Println(sleepSort(values))
}
