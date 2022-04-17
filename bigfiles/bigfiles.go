// Based on example found at:
// https://www.socketloop.com/tutorials/golang-file-system-scanning

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func scan(path string, f os.FileInfo, _ error) error {
	if f.Name() != "." || f.Name() != ".." {
		// Greater than 100 MB
		if f.Size() > 104857600 {
			fmt.Printf("%d|%s\n", f.Size(), path)
		}
	}

	return nil
}

func main() {

	if len(os.Args) != 2 {
		fmt.Printf("USAGE : %s <target_directory> \n", os.Args[0])
		os.Exit(0)
	}

	dir := os.Args[1] // 1st argument is the directory location

	err := filepath.Walk(dir, scan)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
