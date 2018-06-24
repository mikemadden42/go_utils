// Based on example found at:
// https://www.socketloop.com/tutorials/golang-find-file-size-disk-usage-with-filepath-walk

package main

import (
	"fmt"
	"os"
)

// function to return the disk usage information

func diskUsage(currentPath string, info os.FileInfo) int64 {
	size := info.Size()

	if !info.IsDir() {
		return size
	}

	dir, err := os.Open(currentPath)

	if err != nil {
		fmt.Println(err)
		return size
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range files {
		if file.Name() == "." || file.Name() == ".." {
			continue
		}
		size += diskUsage(currentPath+"/"+file.Name(), file)
	}

	fmt.Printf("Size in bytes : [%d] : [%s]\n", size, currentPath)

	return size
}

func main() {

	if len(os.Args) != 2 {
		fmt.Printf("USAGE : %s <target_directory> \n", os.Args[0])
		os.Exit(0)
	}

	dir := os.Args[1] // get the target directory

	info, err := os.Lstat(dir)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	diskUsage(dir, info)
}
