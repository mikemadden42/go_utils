package main

import (
	"crypto/md5"
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

	sum, err := Md5File(*filePtr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s  %s\n", sum, *filePtr)
}

// Md5File returns the MD5 hash for a given file.
func Md5File(filePath string) (string, error) {
	f, err := os.Open(filepath.Clean(filePath))
	if err != nil {
		return "", fmt.Errorf("error opening file for MD5 computation: %v (%s)", err, filePath)
	}
	defer f.Close()

	hasher := md5.New()
	if _, err := io.Copy(hasher, f); err != nil {
		return "", fmt.Errorf("error calculating hash: %v", err)
	}

	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}
