package wpa_supplicant

func CopyWPAHeader(headerFile string) func() (string, error) {
	return func() (string, error) {
		return "", CopyWPAHeaderDirect(headerFile)
	}
}

func AddWPANetwork(wpaNetwork ...WPANetwork) func() (string, error) {
	return func() (string, error) {
		return AddWPANetworkDirect(wpaNetwork...)
	}
}
