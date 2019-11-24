package userGroups

import (
	"bytes"
	"github.com/ui-kreinhard/go-setuper/utils"
	"os"
	"os/exec"
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

func ChangePassword(username, newPassword string) (string, error) {
	cmd := exec.Command("chpasswd")

	buffer := bytes.Buffer{}

	buffer.Write([]byte(username))
	buffer.Write([]byte(":"))
	buffer.Write([]byte(newPassword))
	buffer.Write([]byte("\n"))

	cmd.Stdout = os.Stdout
	cmd.Stdin = &buffer
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	return "", err
}
