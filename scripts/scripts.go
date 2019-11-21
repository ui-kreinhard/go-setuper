package scripts

import (
	"github.com/ui-kreinhard/go-setuperutils"
)

func ExecuteCmd(cmd string, params ...string) (string, error) {
	return utils.Exec(cmd, params...)
}
