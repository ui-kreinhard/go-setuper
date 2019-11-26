package network

import (
	"net"
)

type DeviceType int

const (
	Static DeviceType = iota + 1
	DHCP
	Local
)

type EthernetConfiguration struct {
	Type       DeviceType
	Name        string
	IPV4Address *net.IP
	IPV4Mask    *net.IPMask
	// DNSServers  []net.IP
}

type IConfigCmd interface {
	ConfigureStaticDevices(e ...EthernetConfiguration) IConfigCmd
	ConfigureLoopBack(deviceNames ...string) IConfigCmd
	ConfigureDHCPDevice(deviceNames ...string) IConfigCmd
	Apply() func() (string, error)
}
