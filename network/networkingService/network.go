package networkingService

import (
	"github.com/ui-kreinhard/go-setuper/network"
	"github.com/ui-kreinhard/go-setuper/systemd"
	"github.com/ui-kreinhard/go-setuper/utils"
	"io/ioutil"
)

func renderStatic(config network.EthernetConfiguration) string {
	return ""
}

func renderDHCP(deviceName string) string {
	return `
allow-hotplug ` + deviceName + `
iface ` + deviceName + ` inet dhcp
	`
}

func renderLocal() string {
	return `
auto lo
iface lo inet loopback
	`
}

func render(config network.EthernetConfiguration) string {
	switch config.Type {
	case network.DHCP:
		return renderDHCP(config.Name)
	case network.Static:
		return renderStatic(config)
	case network.Local:
		return renderLocal()
	}
	return renderStatic(config)
}

func validate(config network.EthernetConfiguration) error {
	switch config.Type {
	case network.Static:
		return nil
	case network.Local:
		return nil
	}
	return nil
}

func ClearDirect() (string, error) {
	return "", ioutil.WriteFile("/etc/network/interfaces", []byte(""), 0655)
}

func ConfigureStaticDevicesDirect(deviceConfigurations ...network.EthernetConfiguration) (string, error) {
	configString := ""
	for _, config := range deviceConfigurations {
		err := validate(config)
		if err != nil {
			return "", err
		}
	}
	for _, config := range deviceConfigurations {
		configString = configString + render(config)
	}
	return "", utils.AppendToFile("/etc/network/interfaces", configString)
}

func ApplyDirect() (string, error) {
	return systemd.RestartDirect("ifup@*")
}
