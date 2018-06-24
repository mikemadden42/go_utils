// Based on example found at:
// https://www.socketloop.com/tutorials/golang-generate-md5-checksum-of-a-file

package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	filePtr := flag.String("file", "md5.go", "target file")

	flag.Parse()
	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	sum, _ := Md5File(*filePtr)
	fmt.Printf("%s  %s\n", sum, *filePtr)
}

// Md5File returns the md5 for a given file.
func Md5File(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf(fmt.Sprintf("Error opening file for MD5 computation: %v (%s)", err, filePath))
	}

	hasher := md5.New()
	io.Copy(hasher, f)

	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}
