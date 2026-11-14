// Command dial checks whether the TCP ports listed in dial.json are open.
package main

// http://stackoverflow.com/questions/23140323/creating-a-tcp-client-in-golang
// https://www.socketloop.com/tutorials/golang-how-to-iterate-over-a-string-array
// http://www.dotnetperls.com/convert-go
// http://stackoverflow.com/questions/7782411/is-there-a-foreach-in-go
// https://github.com/golang/go/wiki/Range
// https://gist.github.com/border/775526

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type jsonObject struct {
	Dial dialType
}

type dialType struct {
	Server string
	Ports  []int
}

func main() {
	filePtr := flag.String("file", "dial.json", "JSON file with the server and ports to check")
	flag.Parse()

	file, e := os.ReadFile(filepath.Clean(*filePtr))
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	var jsontype jsonObject
	err := json.Unmarshal(file, &jsontype)
	checkErr(err)

	host := jsontype.Dial.Server
	protocol := "tcp"

	for _, port := range jsontype.Dial.Ports {
		address := net.JoinHostPort(host, strconv.Itoa(port))
		conn, err := net.DialTimeout(protocol, address, 3*time.Second)
		if err != nil {
			fmt.Printf("%4s%30s%6d%5s\n", "FAIL", host, port, protocol)
		} else {
			_ = conn.Close()
			fmt.Printf("%4s%30s%6d%5s\n", "OK", host, port, protocol)
		}
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
