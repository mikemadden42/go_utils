package main

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

func TestSha3File(t *testing.T) {
	path := filepath.Join(t.TempDir(), "abc.txt")
	if err := os.WriteFile(path, []byte("abc"), 0600); err != nil {
		t.Fatal(err)
	}

	got, err := Sha3File(path)
	if err != nil {
		t.Fatal(err)
	}

	want := "3a985da74fe225b2045c172d6bd390bd855f086e3e9d525b46bfe24511431532"
	if got != want {
		t.Errorf("Sha3File(%q) = %s, want %s", path, got, want)
	}
}

func TestSha3FileMissing(t *testing.T) {
	_, err := Sha3File(filepath.Join(t.TempDir(), "missing"))
	if !errors.Is(err, fs.ErrNotExist) {
		t.Errorf("got error %v, want one wrapping fs.ErrNotExist", err)
	}
}
