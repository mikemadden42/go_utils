package main

// http://stackoverflow.com/questions/23140323/creating-a-tcp-client-in-golang
// https://www.socketloop.com/tutorials/golang-how-to-iterate-over-a-string-array
// http://www.dotnetperls.com/convert-go
// http://stackoverflow.com/questions/7782411/is-there-a-foreach-in-go
// https://github.com/golang/go/wiki/Range
// https://gist.github.com/border/775526

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strconv"
)

type jsonObject struct {
	Dial dialType
}

type dialType struct {
	Server string
	Ports  []int
}

func main() {
	file, e := ioutil.ReadFile("./dial.json")
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
		_, err := net.Dial(protocol, host+":"+strconv.Itoa(port))
		if err != nil {
			fmt.Printf("%4s%30s%6d%5s\n", "FAIL", host, port, protocol)
		} else {
			fmt.Printf("%4s%30s%6d%5s\n", "OK", host, port, protocol)
		}
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
