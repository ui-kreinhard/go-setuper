package userGroups

func CreateUserWithoutPasswordDeferred(username string) func() (string, error) {
	return func() (string, error) {
		return CreateUserWithoutPassword(username)
	}
}

func CreateUserDeferred(username, passwordHash string) func() (string, error) {
	return func() (string, error) {
		return CreateUser(username, passwordHash)
	}
}

func ChangePasswordDeferred(username, newPassword string) func() (string, error) {
	return func() (string, error) {
		return ChangePassword(username, newPassword)
	}
}
