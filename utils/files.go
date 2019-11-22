package utils

import (
	"os"
	"path/filepath"
)

func ConstructDestination(from, dest string) (string, error) {
	stat, err := os.Stat(dest)
	if stat.IsDir() {
		filePathElements := filepath.SplitList(from)
		lastElement := filePathElements[len(filePathElements)-1]
		return filepath.Join(dest, lastElement), nil
	}
	return dest, err
}
