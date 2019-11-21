package systemd

import (
	"github.com/ui-kreinhard/go-setuper/utils"
)

func Enable(serviceName string) (string, error) {
	output, err := utils.Exec("systemctl", "enable", serviceName)
	return output, err
}

func Disable(serviceName string) (string, error) {
	output, err := utils.Exec("systemctl", "disable", serviceName)
	return output, err

}

func DaemonReload() (string, error) {
	output, err := utils.Exec("systemctl", "daemon-reload")
	return output, err
}

func Stop(serviceName string) (string, error) {
	output, err := utils.Exec("systemctl", "stop", serviceName)
	return output, err
}

func Start(serviceName string) (string, error) {
	output, err := utils.Exec("systemctl", "start", serviceName)
	return output, err
}

func Restart(serviceName string) (string, error) {
	output, err := utils.Exec("systemctl", "restart", serviceName)
	return output, err
}
