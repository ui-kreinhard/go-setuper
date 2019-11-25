package scripts

func ExecuteCmd(cmd string, params ...string) func() (string, error) {
	return func() (string, error) {
		return ExecuteCmdDirect(cmd, params...)
	}
}
