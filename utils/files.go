package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func getLastFileElement(path string) string {
	filePathElements := filepath.SplitList(path)
	lastElement := filePathElements[len(filePathElements)-1]
	return lastElement
}

func ConstructDestination(from, dest string) (string, error) {
	lastElementFrom := getLastFileElement(from)
	lastElementDest := getLastFileElement(dest)
	if lastElementFrom == lastElementDest {
		fmt.Println(dest)
		return dest, nil
	}
	stat, err := os.Stat(dest)
	if err != nil {
		return "", err
	}
	if stat.IsDir() {
		return filepath.Join(dest, lastElementFrom), nil
	}
	return dest, err
}
