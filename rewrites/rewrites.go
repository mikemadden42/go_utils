// https://www.socketloop.com/tutorials/read-csv-file-go
// https://www.socketloop.com/tutorials/write-csv-to-file-go
// https://www.socketloop.com/tutorials/how-to-unmarshal-or-load-csv-record-into-struct-go
// https://www.socketloop.com/references/golang-encoding-csv-writer-writeall-function-example
// https://golang.org/pkg/net/url/#example_URL

// Command rewrites checks the source and destination URLs in rws.csv.
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
)

func main() {
	filePtr := flag.String("file", "rws.csv", "CSV of code, source URL and destination URL")
	flag.Parse()

	csvFile, err := os.Open(filepath.Clean(*filePtr))

	if err != nil {
		fmt.Println(err)
		return
	}

	defer func(csvFile *os.File) {
		err := csvFile.Close()
		if err != nil {
			fmt.Println("Unable to close file...")
		}
	}(csvFile)

	reader := csv.NewReader(csvFile)

	reader.FieldsPerRecord = -1

	rawCSVData, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// sanity check, display to standard output
	for i, each := range rawCSVData {
		if len(each) < 3 {
			log.Printf("line %d: expected 3 fields, got %d", i+1, len(each))
			continue
		}

		code := each[0]
		src := each[1]
		dst := each[2]

		// Ensure we have a valid src & dst URL.
		for _, site := range each[1:3] {
			u, err := url.Parse(site)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(u.Scheme)
			fmt.Println(u.Host)
			fmt.Println(u.Path)
		}

		fmt.Printf("code: %s\nsource: %s\ndest: %s\n\n", code, src, dst)
	}
}
