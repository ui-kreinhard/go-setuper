package userGroups

import (
	"github.com/ui-kreinhard/go-setuper/utils"
)

func AddGroup(groupName string) (string, error) {
	output, err := utils.Exec("addgroup", groupName)
	if err != nil && err.Error() == "exit status 1" {
		return output, nil
	}
	return output, err
}

func DeleteGroup(groupName string) (string, error) {
	output, err := utils.Exec("delgroup", groupName)
	if err != nil && err.Error() == "exit status 3" {
		return output, nil
	}
	return output, err
}

func AddUserToGroup(user, groupName string) (string, error) {
	output, err := utils.Exec("usermod", "-a", "-G", groupName, user)
	return output, err
}

func RemoveUserFromGroup(user, groupName string) (string, error){
	output, err := utils.Exec("deluser", user, groupName)
	if err != nil && err.Error() == "exit status 6" {
		return output, nil
	}
	return output, err
}
