package timezones

import (
	"github.com/ui-kreinhard/go-setuper/files"
	"github.com/ui-kreinhard/go-setuper/utils"
)

func SetTimezoneDirect(newTimezone Timezone) (string, error) {
	err := files.RemoveFileDirect("/etc/localtime")
	if err != nil {
		return "", err
	}
	err = files.CreateSymlinkDirect("/usr/share/zoneinfo/"+string(newTimezone), "/etc/localtime")
	output, err := utils.Exec("dpkg-reconfigure", "-f", "noninteractive", "tzdata")
	return output, err
}
