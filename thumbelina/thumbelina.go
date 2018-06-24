// https://en.wikipedia.org/wiki/Thumbelina_(horse)

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	checkShell("/var/log/gitlab/gitlab-shell/gitlab-shell.log")
	checkNginx("/var/log/gitlab/nginx/gitlab_error.log")
}

func checkShell(logfile string) {
	// Open an input file, exit on error.
	inputFile, err := os.Open(logfile)
	if err != nil {
		log.Fatal("Error opening input file:", err)
	}

	// Closes the file when we leave the scope of the current function,
	// this makes sure we never forget to close the file if the
	// function can exit in multiple places.
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	var occurrences map[string]int
	occurrences = make(map[string]int)

	fmt.Printf("ERRORS IN %s:\n", logfile)

	// scanner.Scan() advances to the next token returning false if an error was encountered
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "W,") {
			warning := strings.Split(line, "WARN -- : ")[1]
			occurrences[warning]++
		}
	}

	for key, value := range occurrences {
		fmt.Printf("'%s' repeated %d times\n", key, value)
	}
	fmt.Println()

	// When finished scanning if any error other than io.EOF occured
	// it will be returned by scanner.Err().
	if err := scanner.Err(); err != nil {
		log.Fatal(scanner.Err())
	}
}

func checkNginx(logfile string) {
	data, err := ioutil.ReadFile(logfile)
	if err != nil {
		log.Fatal("Error opening input file:", err)
	}

	fmt.Printf("ERRORS IN %s:\n", logfile)
	fmt.Println(string(data))
}
