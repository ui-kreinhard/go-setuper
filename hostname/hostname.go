package hostname

import (
	"github.com/ui-kreinhard/go-setuper/utils"
	"io/ioutil"
	"regexp"
)

func writeToEtcHostname(hostname string) error {
	return ioutil.WriteFile("/etc/hostname", []byte(hostname), 0644)
}

func executeHostnameCmd(hostname string) (string, error) {
	return utils.Exec("hostname", hostname)
}

func ensureLineInHosts(hostname string) error {
	content, err := ioutil.ReadFile("/etc/hosts")
	if err != nil {
		return err
	}
	matched, err := regexp.Match(`127.0.1.1 `+hostname+` `+hostname+`.*$`, content)
	if err == nil && !matched {
		newContent := string(content) + "\n127.0.1.1 " + hostname + " " + hostname
		return ioutil.WriteFile("/etc/hosts", []byte(newContent), 0655)
	}
	return err
}

func SetHostnameDirect(hostname string) (string, error) {
	err := writeToEtcHostname(hostname)
	if err != nil {
		return "", err
	}
	output, err := executeHostnameCmd(hostname)
	if err != nil {
		return output, err
	}
	err = ensureLineInHosts(hostname)

	return "", err
}
