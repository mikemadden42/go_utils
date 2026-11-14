// Command sha1 prints the SHA-1 checksum of a file.
package main

import (
	"crypto/sha1" // #nosec G505 -- used as a file checksum, not for security
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

	sum, err := Sha1File(*filePtr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s  %s\n", sum, *filePtr)
}

// Sha1File returns the SHA-1 hash for a given file.
func Sha1File(filePath string) (string, error) {
	return filehash.File(filePath, sha1.New()) // #nosec G401 -- used as a file checksum, not for security
}
