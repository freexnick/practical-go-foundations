package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func wordFrequency(r io.Reader) (map[string]uint, error) {
	s := bufio.NewScanner(r)
	freqs := make(map[string]uint)
	for s.Scan() {
		words := strings.Fields(s.Text())
		for _, w := range words {
			freqs[strings.ToLower(w)]++
		}
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return freqs, nil
}

func maxWord(freqs map[string]uint) (string, uint, error) {
	if len(freqs) == 0 {
		return "", 0, fmt.Errorf("empty map")
	}
	var (
		maxN uint
		maxW string
	)

	for word, count := range freqs {
		if count > (maxN) {
			maxN, maxW = (count), word
		}
	}
	return maxW, maxN, nil
}

func mostCommon(r io.Reader) (string, uint, error) {
	freqs, err := wordFrequency(r)
	if err != nil {
		return "", 0, err
	}
	return maxWord(freqs)
}

func main() {
	fileName := flag.String("file", "sherlock.txt", "please provide file")
	flag.Parse()
	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer file.Close()

	w, t, err := mostCommon(file)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Printf(`most common word is: "%s", which appeares %d times`, w, t)
}
