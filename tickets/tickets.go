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

	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1
	rawCSVdata, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		return
	}

	var tickets map[string]string
	tickets = make(map[string]string)

	for _, ticket := range rawCSVdata {
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
