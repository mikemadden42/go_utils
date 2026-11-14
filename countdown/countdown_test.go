package main

import (
	"testing"
	"time"
)

func TestDiffDays(t *testing.T) {
	start := time.Date(2020, 7, 6, 14, 0, 0, 0, time.UTC)

	tests := []struct {
		name string
		end  time.Time
		want int
	}{
		{"same time", start, 0},
		{"just under a day", start.Add(23 * time.Hour), 0},
		{"one day", start.Add(24 * time.Hour), 1},
		{"one year", time.Date(2021, 7, 6, 14, 0, 0, 0, time.UTC), 365},
		{"end before start", start.Add(-48 * time.Hour), -2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DiffDays(start, tt.end); got != tt.want {
				t.Errorf("DiffDays() = %d, want %d", got, tt.want)
			}
		})
	}
}
