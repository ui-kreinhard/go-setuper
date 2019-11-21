package scripts

func ExecuteCmdDeferred(cmd string, params ...string) func() (string, error) {
	return func() (string, error) {
		return ExecuteCmd(cmd, params...)
	}
}
