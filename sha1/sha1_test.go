package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSha1File(t *testing.T) {
	path := filepath.Join(t.TempDir(), "abc.txt")
	if err := os.WriteFile(path, []byte("abc"), 0600); err != nil {
		t.Fatal(err)
	}

	got, err := Sha1File(path)
	if err != nil {
		t.Fatal(err)
	}

	want := "a9993e364706816aba3e25717850c26c9cd0d89d"
	if got != want {
		t.Errorf("Sha1File(%q) = %s, want %s", path, got, want)
	}
}

func TestSha1FileMissing(t *testing.T) {
	if _, err := Sha1File(filepath.Join(t.TempDir(), "missing")); err == nil {
		t.Error("expected an error for a missing file")
	}
}
