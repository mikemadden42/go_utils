package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFile, err := os.Open("tickets.csv")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer func(csvFile *os.File) {
		_ = csvFile.Close()
	}(csvFile)

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1
	rawCSVData, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		return
	}

	tickets := make(map[string]string)

	for _, ticket := range rawCSVData {
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
