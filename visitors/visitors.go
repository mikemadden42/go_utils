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

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

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

	defer file.Close()

	scanner := bufio.NewScanner(file)
	const sep string = "^"
	hits := 0
	var ips map[string]int = make(map[string]int)
	var urls map[string]int = make(map[string]int)
	var agents map[string]int = make(map[string]int)

	for scanner.Scan() {
		entry := strings.Split(scanner.Text(), sep)
		//entry[3] - the remote host (the client IP)
		//entry[8] - the page that linked to this URL
		//entry[9] - the browser identification string
		ips[entry[3]]++
		urls[entry[8]]++
		agents[entry[9]]++
		hits++
	}

	fmt.Println("Hits:", hits)

	if *ipMode {
		for k, v := range ips {
			fmt.Println(k, v)
		}
		fmt.Println("")
	}

	if *urlMode {
		for k, v := range urls {
			fmt.Println(k, v)
		}
		fmt.Println("")
	}

	if *agentMode {
		for k, v := range agents {
			fmt.Println(v, k)
		}
		fmt.Println("")
	}
}
