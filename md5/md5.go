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
// Uses named return values (hash, err) to correctly handle the error
// from f.Close() in the deferred function.
func Md5File(filePath string) (hash string, err error) {
	f, err := os.Open(filepath.Clean(filePath))
	if err != nil {
		return "", fmt.Errorf("error opening file for MD5 computation: %v (%s)", err, filePath)
	}
	
	// Defer a function closure to check the error from f.Close()
	defer func() {
		if closeErr := f.Close(); closeErr != nil {
			// If we haven't already recorded an error (e.g., from io.Copy),
			// set the return error to the close error.
			if err == nil {
				err = fmt.Errorf("error closing file: %v", closeErr)
			}
		}
	}()

	hasher := md5.New()
	// Use = instead of := for 'err' since it's now a named return parameter
	if _, err = io.Copy(hasher, f); err != nil {
		return "", fmt.Errorf("error calculating hash: %v", err)
	}

	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}
