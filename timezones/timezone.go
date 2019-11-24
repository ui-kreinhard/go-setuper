package timezones

import (
	"github.com/ui-kreinhard/go-setuper/files"
	"github.com/ui-kreinhard/go-setuper/utils"
)

func SetTimezone(newTimezone Timezone) (string, error) {
	err := files.RemoveFile("/etc/localtime")
	if err != nil {
		return "", err
	}
	err = files.CreateSymlink("/usr/share/zoneinfo/"+string(newTimezone), "/etc/localtime")
	output, err := utils.Exec("dpkg-reconfigure", "-f", "noninteractive", "tzdata")
	return output, err
}
