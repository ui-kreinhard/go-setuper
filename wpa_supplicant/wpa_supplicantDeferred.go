package wpa_supplicant

import (
	"github.com/ui-kreinhard/go-setuper/systemd"
	"github.com/ui-kreinhard/go-setuper/utils"
)

type IConfigCommand interface {
	AddWPANetwork(wpaNetworks ...WPANetwork) IConfigCommand
	Apply() func() (string, error)
}
type ConfigCmd struct {
	queue utils.QueuedFunctions
}

func (c *ConfigCmd) AddWPANetwork(wpaNetworks ...WPANetwork) IConfigCommand {
	c.queue.Append(func() (string, error) {
		return AddWPANetworkDirect(wpaNetworks...)
	})

	return c
}

func (c *ConfigCmd) Apply() func() (string, error) {
	return func() (string, error) {
		output, err := c.queue.Apply()
		if err != nil {
			return output, err
		}
		return systemd.RestartDirect("wpa_supplicant")
	}
}

func CreateWpaSupplicant(config WPASupplicantConfig) IConfigCommand {
	configCmd := &ConfigCmd{}
	configCmd.queue.Append(func() (string, error) {
		return "", ConfigureWPAHeaderDirect(config)
	})
	return configCmd
}
