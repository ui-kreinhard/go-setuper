package files

import (
	"os"
)

func CopyDeferred(from string, dest string) func() (string, error) {
	return func() (string, error) {
		return "", Copy(from, dest)
	}
}

func CreateEmptyFileDeferred(filename string) func() (string, error) {
	return func() (string, error) {
		return "", CreateEmptyFile(filename)
	}
}

func CreateSymlinkDeferred(src, dest string) func() (string, error) {
	return func() (string, error) {
		return "", CreateSymlink(src, dest)
	}
}

func ChmodDeferred(file string, fileMode os.FileMode) func() (string, error) {
	return func() (string, error) {
		return "", Chmod(file, fileMode)
	}
}

func ChownDeferred(file, ownerUser, ownerGroup string) func() (string, error) {
	return func() (string, error) {
		return "", Chown(file, ownerUser, ownerGroup)
	}
}
