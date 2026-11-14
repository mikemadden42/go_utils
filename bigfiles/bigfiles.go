// Based on example found at:
// https://www.socketloop.com/tutorials/golang-file-system-scanning

// Command bigfiles prints every file larger than 100 MB under a directory.
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func scan(path string, f os.FileInfo, err error) error {
	if err != nil {
		// f may be nil here; report the path and keep walking.
		fmt.Fprintln(os.Stderr, err)
		return nil
	}

	// Greater than 100 MB
	if !f.IsDir() && f.Size() > 104857600 {
		fmt.Printf("%d|%s\n", f.Size(), path)
	}

	return nil
}

func main() {

	if len(os.Args) != 2 {
		fmt.Printf("USAGE : %s <target_directory> \n", os.Args[0])
		os.Exit(0)
	}

	dir := os.Args[1] // 1st argument is the directory location

	err := filepath.Walk(dir, scan) // #nosec G703 -- the user chooses which path to scan

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
