// https://rosettacode.org/wiki/Read_a_file_line_by_line#Go
// https://godoc.org/github.com/tatsushid/go-fastping

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	fastping "github.com/tatsushid/go-fastping"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	inputFile, err := os.Open("hosts.txt")
	if err != nil {
		log.Fatal("Error opening input file:", err)
	}

	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		host := scanner.Text()
		fmt.Println(host)
		p := fastping.NewPinger()
		ra, err := net.ResolveIPAddr("ip4:icmp", host)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
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
