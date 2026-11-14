// Command tickets groups the tickets in tickets.csv by assignee.
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	filePtr := flag.String("file", "tickets.csv", "ticket export CSV")
	flag.Parse()

	csvFile, err := os.Open(filepath.Clean(*filePtr))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer func(csvFile *os.File) {
		_ = csvFile.Close()
	}(csvFile)

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1
	rawCSVData, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	tickets := make(map[string]string)

	// Skip the header row; the export has at least 51 columns.
	for i, ticket := range rawCSVData {
		if i == 0 {
			continue
		}
		if len(ticket) < 51 {
			fmt.Fprintf(os.Stderr, "line %d: expected 51 fields, got %d\n", i+1, len(ticket))
			continue
		}

		assignee := ticket[49]
		date := ticket[19]
		order := ticket[50]
		summary := ticket[32]

		currentTicket := fmt.Sprintf("%15s - %22s - %s|", order, date, summary)
		tickets[assignee] += currentTicket
	}

	for key, values := range tickets {
		fmt.Print(key)
		count := strings.Count(values, "|")
		fmt.Println(" -", count, "request(s)")
		assignments := strings.Split(values, "|")
		for _, assignment := range assignments {
			fmt.Println(assignment)
		}
	}
}
