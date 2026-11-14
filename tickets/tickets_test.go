package main

import (
	"bytes"
	"slices"
	"strings"
	"testing"
)

func ticketRow(assignee, date, order, summary string) []string {
	row := make([]string, minFields)
	row[colAssignee] = assignee
	row[colDate] = date
	row[colOrder] = order
	row[colSummary] = summary
	return row
}

func TestGroupByAssignee(t *testing.T) {
	records := [][]string{
		make([]string, minFields), // header
		ticketRow("alice", "2026-01-01", "ORD1", "Fix printer"),
		ticketRow("bob", "2026-01-02", "ORD2", "New laptop"),
		{"too", "short"},
		ticketRow("alice", "2026-01-03", "ORD3", "VPN | access"),
	}

	var warn bytes.Buffer
	got := groupByAssignee(records, &warn)

	want := map[string][]string{
		"alice": {
			"           ORD1 -             2026-01-01 - Fix printer",
			"           ORD3 -             2026-01-03 - VPN | access",
		},
		"bob": {
			"           ORD2 -             2026-01-02 - New laptop",
		},
	}
	if len(got) != len(want) {
		t.Fatalf("got %d assignees, want %d: %v", len(got), len(want), got)
	}
	for assignee, lines := range want {
		if !slices.Equal(got[assignee], lines) {
			t.Errorf("%s = %q, want %q", assignee, got[assignee], lines)
		}
	}

	if w := warn.String(); !strings.Contains(w, "line 4: expected 51 fields, got 2") {
		t.Errorf("warning = %q, want it to mention line 4", w)
	}
}

func TestGroupByAssigneeHeaderOnly(t *testing.T) {
	var warn bytes.Buffer
	got := groupByAssignee([][]string{{"header"}}, &warn)
	if len(got) != 0 || warn.Len() != 0 {
		t.Errorf("got %v and warning %q, want nothing", got, warn.String())
	}
}
