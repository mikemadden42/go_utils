package main

import (
	"maps"
	"strings"
	"testing"
)

func TestCountHits(t *testing.T) {
	log := strings.Join([]string{
		"a^b^c^1.1.1.1^e^f^g^h^/x^curl",
		"",
		"too^short",
		"a^b^c^1.1.1.1^e^f^g^h^/y^wget",
		"a^b^c^2.2.2.2^e^f^g^h^/x^curl^extra",
	}, "\n")

	s, err := countHits(strings.NewReader(log))
	if err != nil {
		t.Fatal(err)
	}

	if s.Hits != 3 {
		t.Errorf("Hits = %d, want 3", s.Hits)
	}
	if s.Skipped != 2 {
		t.Errorf("Skipped = %d, want 2", s.Skipped)
	}

	checks := []struct {
		name string
		got  map[string]int
		want map[string]int
	}{
		{"IPs", s.IPs, map[string]int{"1.1.1.1": 2, "2.2.2.2": 1}},
		{"URLs", s.URLs, map[string]int{"/x": 2, "/y": 1}},
		{"Agents", s.Agents, map[string]int{"curl": 2, "wget": 1}},
	}
	for _, c := range checks {
		if !maps.Equal(c.got, c.want) {
			t.Errorf("%s = %v, want %v", c.name, c.got, c.want)
		}
	}
}

func TestCountHitsEmpty(t *testing.T) {
	s, err := countHits(strings.NewReader(""))
	if err != nil {
		t.Fatal(err)
	}
	if s.Hits != 0 || s.Skipped != 0 {
		t.Errorf("got Hits=%d Skipped=%d, want 0 and 0", s.Hits, s.Skipped)
	}
}
