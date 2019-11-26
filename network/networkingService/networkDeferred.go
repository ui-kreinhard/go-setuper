package networkingService

import (
	"github.com/ui-kreinhard/go-setuper/network"
)

type ConfigCmd struct {
	queue []func() (string, error)
}

func Clear() network.IConfigCmd {

	ret := &ConfigCmd{}
	ret.append(func() (string, error) {
		return ClearDirect()
	})
	return ret
}

func (c *ConfigCmd) append(action func() (string, error)) {
	c.queue = append(c.queue, action)
}

func (c *ConfigCmd) ConfigureStaticDevices(e ...network.EthernetConfiguration) network.IConfigCmd {
	c.append(func() (string, error) {
		return ConfigureStaticDevicesDirect(e...)
	})

	return c
}

func (c *ConfigCmd) ConfigureLoopBack(deviceNames ...string) network.IConfigCmd {
	configs := []network.EthernetConfiguration{}
	for _, deviceName := range deviceNames {
		configs = append(configs, network.EthernetConfiguration{network.Local, deviceName, nil, nil})
	}
	return c.ConfigureStaticDevices(configs...)
}

func (c *ConfigCmd) ConfigureDHCPDevice(deviceNames ...string) network.IConfigCmd {
	configs := []network.EthernetConfiguration{}
	for _, deviceName := range deviceNames {
		configs = append(configs, network.EthernetConfiguration{network.DHCP, deviceName, nil, nil})
	}
	return c.ConfigureStaticDevices(configs...)
}

func (c *ConfigCmd) Apply() func() (string, error) {
	return func() (string, error) {
		for _, action := range c.queue {
			output, err := action()
			if err != nil {
				return output, err
			}
		}
		return ApplyDirect()
	}
}
