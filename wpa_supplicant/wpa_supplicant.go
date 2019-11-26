package wpa_supplicant

import (
	"github.com/ui-kreinhard/go-setuper/utils"
	"io/ioutil"
)

type WPANetwork struct {
	SSID       string
	Passphrase string
}

type WPASupplicantConfig struct {
	Country       string
	UpdateConfig  bool
	CtrlInterface string
}

func renderUpdateConfig(updateConfig bool) string {
	if updateConfig {
		return "1"
	}
	return "0"
}

func ConfigureWPAHeaderDirect(config WPASupplicantConfig) error {
	content := "ctrl_interface=" + config.CtrlInterface +
		"\nupdate_config=" + renderUpdateConfig(config.UpdateConfig) +
		"\ncountry=" + config.Country

	ioutil.WriteFile("/etc/wpa_supplicant/wpa_supplicant.conf", []byte(content), 0655)
	return nil
}

func AddWPANetworkDirect(wpaNetworks ...WPANetwork) (string, error) {
	for _, wpaNetwork := range wpaNetworks {
		output, err := utils.Exec("wpa_passphrase", wpaNetwork.SSID, wpaNetwork.Passphrase)
		if err != nil {
			return output, err
		}
		err = utils.AppendToFile("/etc/wpa_supplicant/wpa_supplicant.conf", output)
		if err != nil {
			return "", err
		}
	}
	return "", nil
}
