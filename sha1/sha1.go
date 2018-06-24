package main

import (
	"crypto/sha1"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	filePtr := flag.String("file", "sha1.go", "target file")

	flag.Parse()
	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	sum, _ := Sha1File(*filePtr)
	fmt.Printf("%s  %s\n", sum, *filePtr)
}

// Sha1File returns the sha1 for a given file.
func Sha1File(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf(fmt.Sprintf("Error opening file for SHA1 computation: %v (%s)", err, filePath))
	}

	hasher := sha1.New()
	io.Copy(hasher, f)

	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}
