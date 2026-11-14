package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSha256File(t *testing.T) {
	path := filepath.Join(t.TempDir(), "abc.txt")
	if err := os.WriteFile(path, []byte("abc"), 0600); err != nil {
		t.Fatal(err)
	}

	got, err := Sha256File(path)
	if err != nil {
		t.Fatal(err)
	}

	want := "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"
	if got != want {
		t.Errorf("Sha256File(%q) = %s, want %s", path, got, want)
	}
}
