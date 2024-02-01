package main

import (
	"compress/gzip"
	"crypto/sha1"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func sha1Sum(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", fmt.Errorf("problem opening file: %w", err)
	}
	defer file.Close()
	var r io.Reader = file

	if strings.HasSuffix(fileName, ".gz") {
		gz, err := gzip.NewReader(file)
		if err != nil {
			return "", fmt.Errorf("problem reading gzip: %w", err)
		}
		defer gz.Close()
		r = gz
	}
	w := sha1.New()
	if _, err := io.Copy(w, r); err != nil {
		return "", fmt.Errorf("problem getting content of a file: %w", err)
	}
	sig := w.Sum(nil)

	return fmt.Sprintf("%x", sig), nil
}

func main() {
	file := flag.String("file", "", "please provide file path")
	flag.Parse()
	sum, err := sha1Sum(*file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sum)
}
