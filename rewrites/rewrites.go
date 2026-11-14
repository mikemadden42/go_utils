// https://www.socketloop.com/tutorials/read-csv-file-go
// https://www.socketloop.com/tutorials/write-csv-to-file-go
// https://www.socketloop.com/tutorials/how-to-unmarshal-or-load-csv-record-into-struct-go
// https://www.socketloop.com/references/golang-encoding-csv-writer-writeall-function-example
// https://golang.org/pkg/net/url/#example_URL

// Command rewrites checks the source and destination URLs in rws.csv.
package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
)

// errShortRow is returned by parseRewrite for rows with fewer than 3 fields.
var errShortRow = errors.New("expected 3 fields")

// rewrite is one row of the rewrites CSV.
type rewrite struct {
	Code   string
	Source *url.URL
	Dest   *url.URL
}

// parseRewrite parses a row of code, source URL and destination URL.
func parseRewrite(row []string) (rewrite, error) {
	if len(row) < 3 {
		return rewrite{}, fmt.Errorf("%w, got %d", errShortRow, len(row))
	}

	src, err := url.Parse(row[1])
	if err != nil {
		return rewrite{}, err
	}

	dst, err := url.Parse(row[2])
	if err != nil {
		return rewrite{}, err
	}

	return rewrite{Code: row[0], Source: src, Dest: dst}, nil
}

func main() {
	filePtr := flag.String("file", "rws.csv", "CSV of code, source URL and destination URL")
	flag.Parse()

	csvFile, err := os.Open(filepath.Clean(*filePtr))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
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
		rw, err := parseRewrite(each)
		if errors.Is(err, errShortRow) {
			log.Printf("line %d: %v", i+1, err)
			continue
		}
		if err != nil {
			log.Fatal(err)
		}

		for _, u := range []*url.URL{rw.Source, rw.Dest} {
			fmt.Println(u.Scheme)
			fmt.Println(u.Host)
			fmt.Println(u.Path)
		}

		fmt.Printf("code: %s\nsource: %s\ndest: %s\n\n", rw.Code, each[1], each[2])
	}
}
