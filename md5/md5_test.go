package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestMd5File(t *testing.T) {
	path := filepath.Join(t.TempDir(), "abc.txt")
	if err := os.WriteFile(path, []byte("abc"), 0600); err != nil {
		t.Fatal(err)
	}

	got, err := Md5File(path)
	if err != nil {
		t.Fatal(err)
	}

	want := "900150983cd24fb0d6963f7d28e17f72"
	if got != want {
		t.Errorf("Md5File(%q) = %s, want %s", path, got, want)
	}
}

func TestMd5FileMissing(t *testing.T) {
	if _, err := Md5File(filepath.Join(t.TempDir(), "missing")); err == nil {
		t.Error("expected an error for a missing file")
	}
}
