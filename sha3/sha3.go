// Command sha3 prints the SHA3-256 checksum of a file.
package main

import (
	"crypto/sha3"
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

	sum, err := Sha3File(*filePtr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s  %s\n", sum, *filePtr)
}

// Sha3File returns the SHA3-256 hash for a given file.
func Sha3File(filePath string) (string, error) {
	return filehash.File(filePath, sha3.New256())
}
