package scripts

import (
	"github.com/ui-kreinhard/go-setuper/utils"
)

func ExecuteCmd(cmd string, params ...string) (string, error) {
	return utils.Exec(cmd, params...)
}
