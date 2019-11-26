package networkingService

import (
	"github.com/ui-kreinhard/go-setuper/network"
	"github.com/ui-kreinhard/go-setuper/utils"
)

type ConfigCmd struct {
	queue utils.QueuedFunctions
}

func Clear() network.IConfigCmd {

	ret := &ConfigCmd{}
	ret.queue.Append(func() (string, error) {
		return ClearDirect()
	})
	return ret
}

func (c *ConfigCmd) ConfigureStaticDevices(e ...network.EthernetConfiguration) network.IConfigCmd {
	c.queue.Append(func() (string, error) {
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
		output, err := c.queue.Apply()
		if err != nil {
			return output, err
		}
		return ApplyDirect()
	}
}
