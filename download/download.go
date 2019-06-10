// Based on examples found at:
// https://www.socketloop.com/tutorials/golang-download-file-example
// https://www.socketloop.com/tutorials/how-to-create-directory-in-go
// https://www.socketloop.com/tutorials/golang-check-if-a-directory-exist-in-go

// Golang, TLS & Comodo
// http://bridge.grumpy-troll.org/2014/05/golang-tls-comodo/

package main

import (
	"bufio"
	"crypto/sha512"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	// The size of a SHA512 checksum in bytes.
	fmt.Printf("SHA512_size,%v\n", sha512.Size)

	file, err := os.Open("urls.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		host := scanner.Text()
		get(host)
	}

}

func get(src string) {
	rawURL := src

	fileURL, err := url.Parse(rawURL)

	if err != nil {
		panic(err)
	}

	path := fileURL.Path

	segments := strings.Split(path, "/")

	downloadDir := "data"
	err = os.Mkdir(downloadDir, 0700)
	checkErr(err)
	fileName := downloadDir + string(os.PathSeparator) + segments[len(segments)-1]

	file, err := os.Create(fileName)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer file.Close()

	//fmt.Printf("Downloading file %s...", fileName)
	//fmt.Println()

	check := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	resp, err := check.Get(rawURL) // add a filter to check redirect

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer resp.Body.Close()
	//fmt.Println(resp.Status)

	size, err := io.Copy(file, resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s,%s,%v", resp.Status, rawURL, size)
	fmt.Println()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
