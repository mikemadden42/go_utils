package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	filePtr := flag.String("file", "sha256.go", "target file")

	flag.Parse()
	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	sum, _ := Sha256File(*filePtr)
	fmt.Printf("%s  %s\n", sum, *filePtr)
}

// Sha256File returns the sha256 for a given file.
func Sha256File(filePath string) (string, error) {
	f, err := os.Open(filepath.Clean(filePath))
	if err != nil {
		return "", fmt.Errorf(fmt.Sprintf("Error opening file for SHA1 computation: %v (%s)", err, filePath))
	}

	hasher := sha256.New()
	_, err = io.Copy(hasher, f)
	checkErr(err)

	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
