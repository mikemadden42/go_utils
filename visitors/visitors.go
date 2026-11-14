// Links I found useful when learning Go:
// http://blog.golang.org/go-maps-in-action
// http://crosbymichael.com/go-lang.html
// http://stackoverflow.com/questions/11462879/strings-split-in-golang
// http://stackoverflow.com/questions/16551354/how-to-split-string-and-assign
// http://stackoverflow.com/questions/8018719/iterating-through-a-golang-map
// http://stackoverflow.com/questions/9234699/understanding-apache-access-log
// http://www.golang-book.com/4/index.htm#section3
// https://gobyexample.com/command-line-flags
// https://gobyexample.com/string-formatting
// https://gobyexample.com/string-functions

// Running the program:
// visitors -ip | sort -rn -k2 | head -5
// visitors -ip -log=access.log | sort -rn -k2 | head -5
// visitors -url -log=access.log | sort -rn -k2 | head -5
// visitors -agent -log=access.log | sort -rn -k1 | head -5

// Command visitors counts hits by IP, URL or user agent in a
// caret-separated access log.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

// stats holds hit counts from an access log.
type stats struct {
	Hits    int
	Skipped int
	IPs     map[string]int
	URLs    map[string]int
	Agents  map[string]int
}

// countHits reads a caret-separated access log and counts hits by client
// IP, referring URL and user agent. Lines with fewer than 10 fields are
// counted in Skipped.
func countHits(r io.Reader) (stats, error) {
	const sep string = "^"
	s := stats{
		IPs:    make(map[string]int),
		URLs:   make(map[string]int),
		Agents: make(map[string]int),
	}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		entry := strings.Split(scanner.Text(), sep)
		if len(entry) < 10 {
			s.Skipped++
			continue
		}
		//entry[3] - the remote host (the client IP)
		//entry[8] - the page that linked to this URL
		//entry[9] - the browser identification string
		s.IPs[entry[3]]++
		s.URLs[entry[8]]++
		s.Agents[entry[9]]++
		s.Hits++
	}

	return s, scanner.Err()
}

func main() {
	ipMode := flag.Bool("ip", false, "enable stats by IP")
	urlMode := flag.Bool("url", false, "enable stats by URL")
	agentMode := flag.Bool("agent", false, "enable stats by agent")
	accessLog := flag.String("log", "access.log", "apache access log")

	flag.Parse()

	file, err := os.Open(*accessLog)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Unable to close file...")
		}
	}(file)

	s, err := countHits(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Hits:", s.Hits)
	if s.Skipped > 0 {
		fmt.Fprintln(os.Stderr, "Skipped malformed lines:", s.Skipped)
	}

	if *ipMode {
		for k, v := range s.IPs {
			fmt.Println(k, v)
		}
		fmt.Println("")
	}

	if *urlMode {
		for k, v := range s.URLs {
			fmt.Println(k, v)
		}
		fmt.Println("")
	}

	if *agentMode {
		for k, v := range s.Agents {
			fmt.Println(v, k)
		}
		fmt.Println("")
	}
}
