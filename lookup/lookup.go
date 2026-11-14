// http://stackoverflow.com/questions/16228887/why-does-fmt-println-inside-a-goroutine-not-print-a-line
// http://golang.org/pkg/sync/#WaitGroup
// https://gobyexample.com/command-line-flags

// Command lookup resolves each host in a file and reverse-resolves its
// addresses.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
)

var wg sync.WaitGroup

// maxLookups limits how many hosts are resolved at once.
const maxLookups = 20

func main() {
	filePtr := flag.String("file", "hosts.txt", "hosts file")

	flag.Parse()

	file, err := os.Open(*filePtr)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Unable to close file...")
		}
	}(file)

	sem := make(chan struct{}, maxLookups)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		host := scanner.Text()
		wg.Add(1)
		sem <- struct{}{}
		go func() {
			defer func() { <-sem }()
			resolve(host)
		}()
	}

	wg.Wait()

	if err := scanner.Err(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func resolve(host string) {
	defer wg.Done()

	ips, err := net.LookupHost(host)
	if err != nil {
		log.Println(err)
	}

	for _, ip := range ips {
		fmt.Println("OK", host, ip)
		hosts, err := net.LookupAddr(ip)
		if err != nil {
			log.Println(err)
		}

		for _, host := range hosts {
			fmt.Println("OK", ip, host)
		}
	}
}
