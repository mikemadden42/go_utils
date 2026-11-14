// Command md5 prints the MD5 checksum of a file.
package main

import (
	"crypto/md5" // #nosec G501 -- used as a file checksum, not for security
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

	sum, err := Md5File(*filePtr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s  %s\n", sum, *filePtr)
}

// Md5File returns the MD5 hash for a given file.
func Md5File(filePath string) (string, error) {
	return filehash.File(filePath, md5.New()) // #nosec G401 -- used as a file checksum, not for security
}
