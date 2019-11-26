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

func AppendToFile(toFile, data string) error {
	f, err := os.OpenFile(toFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0655)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.WriteString(data + "\n"); err != nil {
		return err
	}
	return nil
}
