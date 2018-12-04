// http://stackoverflow.com/questions/16228887/why-does-fmt-println-inside-a-goroutine-not-print-a-line
// http://golang.org/pkg/sync/#WaitGroup
// https://gobyexample.com/command-line-flags

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

func main() {
	filePtr := flag.String("file", "hosts.txt", "hosts file")

	flag.Parse()
	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	file, err := os.Open(*filePtr)
	if err != nil {
		log.Println(err.Error())
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

	defer wg.Done()
}
