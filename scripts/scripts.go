package scripts

import (
	"github.com/ui-kreinhard/go-setuper/utils"
)

func ExecuteCmdDirect(cmd string, params ...string) (string, error) {
	return utils.Exec(cmd, params...)
}
