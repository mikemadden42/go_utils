// Command tickets groups the tickets in tickets.csv by assignee.
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
)

// Column positions in the ticket export.
const (
	colDate     = 19
	colSummary  = 32
	colAssignee = 49
	colOrder    = 50
	minFields   = colOrder + 1
)

// groupByAssignee formats each ticket in records and groups the results by
// assignee. The first record is the header row and is skipped. Rows with
// too few fields are reported to warn and skipped.
func groupByAssignee(records [][]string, warn io.Writer) map[string][]string {
	tickets := make(map[string][]string)

	for i, ticket := range records {
		if i == 0 {
			continue
		}
		if len(ticket) < minFields {
			_, _ = fmt.Fprintf(warn, "line %d: expected %d fields, got %d\n", i+1, minFields, len(ticket))
			continue
		}

		assignee := ticket[colAssignee]
		line := fmt.Sprintf("%15s - %22s - %s", ticket[colOrder], ticket[colDate], ticket[colSummary])
		tickets[assignee] = append(tickets[assignee], line)
	}

	return tickets
}

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

	tickets := groupByAssignee(rawCSVData, os.Stderr)

	// Print assignees in sorted order so the output is stable between runs.
	assignees := make([]string, 0, len(tickets))
	for assignee := range tickets {
		assignees = append(assignees, assignee)
	}
	slices.Sort(assignees)

	for _, assignee := range assignees {
		fmt.Println(assignee, "-", len(tickets[assignee]), "request(s)")
		for _, line := range tickets[assignee] {
			fmt.Println(line)
		}
		fmt.Println()
	}
}
