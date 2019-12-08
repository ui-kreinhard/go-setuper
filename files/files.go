package files

import (
	"github.com/gobuffalo/packr/v2/file"
	"github.com/ui-kreinhard/go-setuper/setuper"
	"github.com/ui-kreinhard/go-setuper/utils"
	"io/ioutil"
	"os"
	"os/user"
	"log"
	"path/filepath"
	"strconv"
)

func CopyDirect(from,dest string) error {
	setuper := setuper.NewSetuper()
	log.Println("checkin", from)
	if setuper.FilesBox.HasDir(from) {
		fileListToCopy := []string{}
		setuper.FilesBox.WalkPrefix(from, func(path string, f file.File) error {
			log.Println(path)
			fileListToCopy = append(fileListToCopy, path)
			return nil
		})
		for _, fileToCopy := range fileListToCopy {
			destinationWithPath := dest + filepath.Dir(fileToCopy)
			if err := CopyDirectFile(fileToCopy, destinationWithPath); err != nil {
				return err
			}
		}
		return nil
	}
	return CopyDirectFile(from, dest)
}

func CopyDirectFile(from string, dest string) error {

	setuper := setuper.NewSetuper()
	os.MkdirAll(dest, 0655)
	sourceContent, err := setuper.FilesBox.FindString(from)
	if err != nil {
		return err
	}

	destinationConstructed, err := utils.ConstructDestination(from, dest)

	err = ioutil.WriteFile(destinationConstructed, []byte(sourceContent), 0666)
	if err != nil {
		return err
	}
	return nil
}

func CreateEmptyFileDirect(filename string) error {
	emptyFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	emptyFile.Close()
	return nil
}

func CreateSymlinkDirect(src, dest string) error {
	return os.Symlink(src, dest)
}

func ChmodDirect(file string, mode os.FileMode) error {
	return os.Chmod(file, mode)
}

func ChownDirect(file, ownerUser, ownerGroup string) error {
	ownerUserObj, err := user.Lookup(ownerUser)
	if err != nil {
		return err
	}

	group, err := user.LookupGroup(ownerGroup)
	if err != nil {
		return err
	}

	ownerUserID, err := strconv.ParseInt(ownerUserObj.Uid, 10, 64)
	if err != nil {
		return err
	}
	ownerGroupID, err := strconv.ParseInt(group.Gid, 10, 64)
	if err != nil {
		return err
	}
	return os.Chown(file, int(ownerUserID), int(ownerGroupID))
}

func RemoveFileDirect(file string) error {
	return os.Remove(file)
}
