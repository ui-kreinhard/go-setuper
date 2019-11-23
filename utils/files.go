package utils

import (
	"os"
	"path/filepath"
)

func getLastFileElement(path string) string {
	_, lastElement := filepath.Split(path)
	return lastElement
}

func ConstructDestination(from, dest string) (string, error) {
	lastElementFrom := getLastFileElement(from)
	lastElementDest := getLastFileElement(dest)
	if lastElementFrom == lastElementDest {
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
