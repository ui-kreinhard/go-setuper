package wpa_supplicant

import (
	"github.com/ui-kreinhard/go-setuper/files"
	"github.com/ui-kreinhard/go-setuper/utils"
	"os"
)

type WPANetwork struct {
	SSID       string
	Passphrase string
}

func CopyWPAHeader(headerFile string) error {
	return files.Copy(headerFile, "/etc/wpa_supplicant/wpa_supplicant.conf")
}

func AddWPANetwork(wpaNetworks ...WPANetwork) (string, error) {
	for _, wpaNetwork := range wpaNetworks {
		output, err := utils.Exec("wpa_passphrase", wpaNetwork.SSID, wpaNetwork.Passphrase)
		if err != nil {
			return output, err
		}
		f, err := os.OpenFile("/etc/wpa_supplicant/wpa_supplicant.conf", os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return "", err
		}
		if _, err := f.Write([]byte("\n" + output)); err != nil {
			return "", err
		}
		if err := f.Close(); err != nil {
			return "", err
		}
	}
	return "", nil
}
