package filehash

import (
	"crypto/sha256"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

func TestFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "abc.txt")
	if err := os.WriteFile(path, []byte("abc"), 0600); err != nil {
		t.Fatal(err)
	}

	got, err := File(path, sha256.New())
	if err != nil {
		t.Fatal(err)
	}

	want := "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"
	if got != want {
		t.Errorf("File(%q) = %s, want %s", path, got, want)
	}
}

func TestFileMissing(t *testing.T) {
	_, err := File(filepath.Join(t.TempDir(), "missing"), sha256.New())
	if !errors.Is(err, fs.ErrNotExist) {
		t.Errorf("got error %v, want one wrapping fs.ErrNotExist", err)
	}
}

func TestFileDirectory(t *testing.T) {
	if _, err := File(t.TempDir(), sha256.New()); err == nil {
		t.Error("expected an error for a directory")
	}
}
