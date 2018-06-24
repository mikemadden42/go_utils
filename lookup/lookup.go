// http://stackoverflow.com/questions/16228887/why-does-fmt-println-inside-a-goroutine-not-print-a-line
// http://golang.org/pkg/sync/#WaitGroup
// https://gobyexample.com/command-line-flags

package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {
	filePtr := flag.String("file", "hosts.txt", "hosts file")

	flag.Parse()
	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	file, err := os.Open(*filePtr)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		host := scanner.Text()
		wg.Add(1)
		go resolve(host)
	}

	wg.Wait()
}

func resolve(host string) {
	addrs, err := net.LookupHost(host)
	if err != nil {
		fmt.Println(host, err.Error())
	} else {
		fmt.Println(host)
		for _, s := range addrs {
			fmt.Println(s)
		}
	}

	defer wg.Done()
}
