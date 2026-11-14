// Package filehash computes checksums of files for the hash commands.
package filehash

import (
	"fmt"
	"hash"
	"io"
	"os"
	"path/filepath"
)

// File returns the hex-encoded checksum of the file at filePath using h.
// Uses named return values (sum, err) to correctly handle the error
// from f.Close() in the deferred function.
func File(filePath string, h hash.Hash) (sum string, err error) {
	f, err := os.Open(filepath.Clean(filePath))
	if err != nil {
		return "", fmt.Errorf("error opening file: %w (%s)", err, filePath)
	}

	// Defer a function closure to check the error from f.Close()
	defer func() {
		if closeErr := f.Close(); closeErr != nil {
			// If we haven't already recorded an error (e.g., from io.Copy),
			// set the return error to the close error.
			if err == nil {
				err = fmt.Errorf("error closing file: %w", closeErr)
			}
		}
	}()

	if _, err = io.Copy(h, f); err != nil {
		return "", fmt.Errorf("error calculating hash: %w", err)
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
