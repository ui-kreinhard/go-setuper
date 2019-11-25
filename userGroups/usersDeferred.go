package userGroups

func CreateUserWithoutPassword(username string) func() (string, error) {
	return func() (string, error) {
		return CreateUserWithoutPasswordDirect(username)
	}
}

func CreateUser(username, passwordHash string) func() (string, error) {
	return func() (string, error) {
		return CreateUserDirect(username, passwordHash)
	}
}

func ChangePassword(username, newPassword string) func() (string, error) {
	return func() (string, error) {
		return ChangePasswordDirect(username, newPassword)
	}
}
