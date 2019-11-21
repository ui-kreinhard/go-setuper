package userGroups

import (
	"github.com/ui-kreinhard/go-setuper/utils"
)

func CreateUserWithoutPassword(username string) (string, error) {
	output, err := utils.Exec("adduser", "--disabled-password", "--gecos", "", username)
	if err != nil && err.Error() == "exit status 1" {
		return output, nil
	}
	return output, err
}

func CreateUser(username, passwordHash string) (string, error) {
	output, err := utils.Exec("useradd", "-m", "-p", passwordHash, "-s", "/bin/bash", username)
	if err != nil && err.Error() == "exit status 9" {
		return output, nil
	}
	return output, err
}

func DeleteUser(username string) (string, error) {
	output, err := utils.Exec("deluser", username)
	if err != nil && err.Error() == "exit status 2" {
		return output, nil
	}
	return output, err
}
