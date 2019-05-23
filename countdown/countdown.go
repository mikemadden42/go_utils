// http://play.golang.org/p/NxuQTlVA2l

// google search: site:play.golang.org date duration

package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Date(2019, 4, 8, 14, 0, 0, 0, time.UTC)
	end := time.Now().UTC()
	fmt.Println("diff:", DiffDays(start, end), "days")
}

// DiffDays gets the difference between two dates.
func DiffDays(start time.Time, end time.Time) int {
	return int(end.Sub(start) / (24 * time.Hour))
}
