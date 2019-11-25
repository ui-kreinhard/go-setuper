package systemd

import (
	"github.com/ui-kreinhard/go-setuper/utils"
)

func EnableDirect(serviceName string) (string, error) {
	output, err := utils.Exec("systemctl", "enable", serviceName)
	return output, err
}

func DisableDirect(serviceName string) (string, error) {
	output, err := utils.Exec("systemctl", "disable", serviceName)
	return output, err

}

func DaemonReloadDirect() (string, error) {
	output, err := utils.Exec("systemctl", "daemon-reload")
	return output, err
}

func StopDirect(serviceName string) (string, error) {
	output, err := utils.Exec("systemctl", "stop", serviceName)
	return output, err
}

func StartDirect(serviceName string) (string, error) {
	output, err := utils.Exec("systemctl", "start", serviceName)
	return output, err
}

func RestartDirect(serviceName string) (string, error) {
	output, err := utils.Exec("systemctl", "restart", serviceName)
	return output, err
}
