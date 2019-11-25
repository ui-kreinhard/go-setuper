package userGroups

func AddGroup(groupName string) func() (string, error) {
	return func() (string, error) {
		return AddGroupDirect(groupName)
	}
}

func DeleteGroup(groupName string) func() (string, error) {
	return func() (string, error) {
		return DeleteGroupDirect(groupName)
	}
}

func AddUserToGroup(user, groupName string) func() (string, error) {
	return func() (string, error) {
		return AddUserToGroupDirect(user, groupName)
	}
}

func RemoveUserFromGroup(user, groupName string) func() (string, error) {
	return func() (string, error) {
		return RemoveUserFromGroupDirect(user, groupName)
	}
}
