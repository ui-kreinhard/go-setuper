package files

import (
	"os"
)

func Copy(from string, dest string) func() (string, error) {
	return func() (string, error) {
		return "", CopyDirect(from, dest)
	}
}

func CreateEmptyFile(filename string) func() (string, error) {
	return func() (string, error) {
		return "", CreateEmptyFileDirect(filename)
	}
}

func CreateSymlink(src, dest string) func() (string, error) {
	return func() (string, error) {
		return "", CreateSymlinkDirect(src, dest)
	}
}

func Chmod(file string, fileMode os.FileMode) func() (string, error) {
	return func() (string, error) {
		return "", ChmodDirect(file, fileMode)
	}
}

func Chown(file, ownerUser, ownerGroup string) func() (string, error) {
	return func() (string, error) {
		return "", ChownDirect(file, ownerUser, ownerGroup)
	}
}

func RemoveFile(file string) func() (string, error) {
	return func() (string, error) {
		return "", RemoveFileDirect(file)
	}
}
