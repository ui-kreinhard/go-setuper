package wpa_supplicant

func CopyWPAHeaderDeferred(headerFile string) func() (string, error) {
	return func() (string, error) {
		return "", CopyWPAHeader(headerFile)
	}
}

func AddWPANetworkDeferred(wpaNetwork ...WPANetwork) func() (string, error) {
	return func() (string, error) {
		return AddWPANetwork(wpaNetwork...)
	}
}
