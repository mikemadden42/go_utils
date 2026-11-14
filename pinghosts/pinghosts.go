// https://rosettacode.org/wiki/Read_a_file_line_by_line#Go
// https://godoc.org/github.com/tatsushid/go-fastping

// Command pinghosts pings each host listed in hosts.txt.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/tatsushid/go-fastping"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	filePtr := flag.String("file", "hosts.txt", "file of hosts to ping, one per line")
	flag.Parse()

	inputFile, err := os.Open(filepath.Clean(*filePtr))
	if err != nil {
		log.Fatal("Error opening input file:", err)
	}

	defer func(inputFile *os.File) {
		err := inputFile.Close()
		if err != nil {
			fmt.Println("Unable to close file...")
		}
	}(inputFile)

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		host := scanner.Text()
		fmt.Println(host)
		p := fastping.NewPinger()
		ra, err := net.ResolveIPAddr("ip4:icmp", host)
		if err != nil {
			fmt.Println(err)
			fmt.Println()
			continue
		}

		success := false
		p.AddIPAddr(ra)
		p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
			fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
			success = true
		}

		p.OnIdle = func() {
			fmt.Println("finished...", success)
			fmt.Println()
		}

		err = p.Run()
		if err != nil {
			fmt.Println(err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(scanner.Err())
	}
}
