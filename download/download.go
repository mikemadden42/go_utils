// Based on examples found at:
// https://www.socketloop.com/tutorials/golang-download-file-example
// https://www.socketloop.com/tutorials/how-to-create-directory-in-go
// https://www.socketloop.com/tutorials/golang-check-if-a-directory-exist-in-go

// Golang, TLS
// http://bridge.grumpy-troll.org/2014/05/golang-tls-comodo/

package main

import (
	"bufio"
	"crypto/sha512"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
)

const downloadDir = "data"

// maxRedirects matches the limit of Go's default redirect policy.
const maxRedirects = 10

func main() {
	// The size of a SHA512 checksum in bytes.
	fmt.Printf("SHA512_size,%v\n", sha512.Size)

	file, err := os.Open("urls.txt")
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

	err = os.MkdirAll(downloadDir, 0700)
	checkErr(err)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		host := scanner.Text()
		if err := get(host); err != nil {
			fmt.Println(err)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func get(rawURL string) error {
	fileURL, err := url.Parse(rawURL)
	if err != nil {
		return err
	}

	name := path.Base(fileURL.Path)
	if name == "." || name == "/" {
		return fmt.Errorf("%s: no file name in URL path", rawURL)
	}

	check := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			if len(via) >= maxRedirects {
				return errors.New("stopped after 10 redirects")
			}
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	resp, err := check.Get(rawURL)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Unable to close file...")
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s,%s", resp.Status, rawURL)
	}

	file, err := os.Create(filepath.Clean(filepath.Join(downloadDir, name)))
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Unable to close file...")
		}
	}(file)

	size, err := io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("%s,%s,%v\n", resp.Status, rawURL, size)
	return nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
