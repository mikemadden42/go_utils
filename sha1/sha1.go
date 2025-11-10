package main

import (
	"crypto/sha1"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	filePtr := flag.String("file", "", "target file")
	flag.Parse()

	if *filePtr == "" {
		flag.Usage()
		os.Exit(1)
	}

	sum, err := Sha1File(*filePtr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s  %s\n", sum, *filePtr)
}

// Sha1File returns the SHA-1 hash for a given file.
func Sha1File(filePath string) (string, error) {
	f, err := os.Open(filepath.Clean(filePath))
	if err != nil {
		return "", fmt.Errorf("error opening file for SHA-1 computation: %v (%s)", err, filePath)
	}
	defer f.Close()

	hasher := sha1.New()
	if _, err := io.Copy(hasher, f); err != nil {
		return "", fmt.Errorf("error calculating hash: %v", err)
	}

	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}
