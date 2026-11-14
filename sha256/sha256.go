// Command sha256 prints the SHA-256 checksum of a file.
package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"os"

	"github.com/mikemadden42/go_utils/internal/filehash"
)

func main() {
	filePtr := flag.String("file", "", "target file")
	flag.Parse()

	if *filePtr == "" {
		flag.Usage()
		os.Exit(1)
	}

	sum, err := Sha256File(*filePtr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s  %s\n", sum, *filePtr)
}

// Sha256File returns the SHA-256 hash for a given file.
func Sha256File(filePath string) (string, error) {
	return filehash.File(filePath, sha256.New())
}
