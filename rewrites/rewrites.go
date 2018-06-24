// https://www.socketloop.com/tutorials/read-csv-file-go
// https://www.socketloop.com/tutorials/write-csv-to-file-go
// https://www.socketloop.com/tutorials/how-to-unmarshal-or-load-csv-record-into-struct-go
// https://www.socketloop.com/references/golang-encoding-csv-writer-writeall-function-example
// https://golang.org/pkg/net/url/#example_URL

package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/url"
	"os"
)

func main() {

	csvfile, err := os.Open("rws.csv")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer csvfile.Close()

	reader := csv.NewReader(csvfile)

	reader.FieldsPerRecord = -1

	rawCSVdata, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// sanity check, display to standard output
	for _, each := range rawCSVdata {
		code := each[0]
		src := each[1]
		dst := each[2]

		// Ensure we have a valid src & dst URL.
		for _, site := range each[1:3] {
			u, err := url.Parse(site)
			fmt.Println(u.Scheme)
			fmt.Println(u.Host)
			fmt.Println(u.Path)

			if err != nil {
				log.Fatal(err)
			}
		}

		fmt.Printf("code: %s\nsource: %s\ndest: %s\n\n", code, src, dst)
	}
}
