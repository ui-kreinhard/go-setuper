package utils

import (
	"os/exec"
)

func Exec(cmd string, params ...string) (string, error) {
	return ConvertOutput(exec.Command(cmd, params...).CombinedOutput())
}
