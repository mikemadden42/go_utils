package main

import (
	"errors"
	"testing"
)

func TestParseRewrite(t *testing.T) {
	rw, err := parseRewrite([]string{"301", "http://a.com/x", "https://b.com/y"})
	if err != nil {
		t.Fatal(err)
	}

	if rw.Code != "301" {
		t.Errorf("Code = %q, want 301", rw.Code)
	}
	if rw.Source.Host != "a.com" || rw.Source.Path != "/x" {
		t.Errorf("Source = %v, want host a.com and path /x", rw.Source)
	}
	if rw.Dest.Scheme != "https" || rw.Dest.Host != "b.com" || rw.Dest.Path != "/y" {
		t.Errorf("Dest = %v, want https://b.com/y", rw.Dest)
	}
}

func TestParseRewriteErrors(t *testing.T) {
	tests := []struct {
		name  string
		row   []string
		short bool
	}{
		{"empty row", nil, true},
		{"two fields", []string{"301", "http://a.com"}, true},
		{"bad source", []string{"301", "http://[bad", "/y"}, false},
		{"bad dest", []string{"301", "/x", "http://[bad"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := parseRewrite(tt.row)
			if err == nil {
				t.Fatal("expected an error")
			}
			if got := errors.Is(err, errShortRow); got != tt.short {
				t.Errorf("errors.Is(err, errShortRow) = %v, want %v (err: %v)", got, tt.short, err)
			}
		})
	}
}
