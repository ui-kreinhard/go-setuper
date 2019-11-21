package userGroups

func AddGroupDeferred(groupName string) func() (string, error) {
	return func() (string, error) {
		return AddGroup(groupName)
	}
}

func DeleteGroupDeferred(groupName string) func() (string, error) {
	return func() (string, error) {
		return DeleteGroup(groupName)
	}
}

func AddUserToGroupDeferred(user, groupName string) func() (string, error) {
	return func() (string, error) {
		return AddUserToGroup(user, groupName)
	}
}

func RemoveUserFromGroupDeferred(user, groupName string) func() (string, error) {
	return func() (string, error) {
		return RemoveUserFromGroup(user, groupName)
	}
}
