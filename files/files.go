package files

import (
	"github.com/ui-kreinhard/go-setuper/setuper"
	"github.com/ui-kreinhard/go-setuper/utils"
	"io/ioutil"
	"os"
	"os/user"
	"strconv"
)

func Copy(from string, dest string) error {
	setuper := setuper.NewSetuper()
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

func CreateEmptyFile(filename string) error {
	emptyFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	emptyFile.Close()
	return nil
}

func CreateSymlink(src, dest string) error {
	return os.Symlink(src, dest)
}

func Chmod(file string, mode os.FileMode) error {
	return os.Chmod(file, mode)
}

func Chown(file, ownerUser, ownerGroup string) error {
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
