// http://play.golang.org/p/NxuQTlVA2l

// google search: site:play.golang.org date duration

// Command countdown prints the number of days since a start date.
package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	since := flag.String("since", "2020-07-06T14:00:00Z", "start date, as YYYY-MM-DD or RFC 3339")
	flag.Parse()

	start, err := parseDate(*since)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	end := time.Now().UTC()
	fmt.Println("diff:", DiffDays(start, end), "days")
}

// DiffDays gets the difference between two dates.
func DiffDays(start time.Time, end time.Time) int {
	return int(end.Sub(start) / (24 * time.Hour))
}

// parseDate accepts a plain date (midnight UTC) or an RFC 3339 timestamp.
func parseDate(s string) (time.Time, error) {
	if t, err := time.Parse(time.DateOnly, s); err == nil {
		return t, nil
	}
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date %q: use YYYY-MM-DD or RFC 3339", s)
	}
	return t, nil
}
