package wpa_supplicant

import (
	"github.com/ui-kreinhard/go-setuper/systemd"
)

type IConfigCommand interface {
	AddWPANetwork(wpaNetworks ...WPANetwork) IConfigCommand
	Apply() func() (string, error)
}
type ConfigCmd struct {
	queue []func() (string, error)
}

func (c *ConfigCmd) append(action func() (string, error)) {
	c.queue = append(c.queue, action)
}

func (c *ConfigCmd) AddWPANetwork(wpaNetworks ...WPANetwork) IConfigCommand {
	c.append(func() (string, error) {
		return AddWPANetworkDirect(wpaNetworks...)
	})

	return c
}

func (c *ConfigCmd) Apply() func() (string, error) {
	return func() (string, error) {
		for _, action := range c.queue {
			output, err := action()
			if err != nil {
				return output, err
			}
		}
		return systemd.RestartDirect("wpa_supplicant")
	}
}

func CreateWpaSupplicant(config WPASupplicantConfig) IConfigCommand {
	configCmd := &ConfigCmd{}
	configCmd.append(func() (string, error) {
		return "", ConfigureWPAHeaderDirect(config)
	})
	return configCmd
}
