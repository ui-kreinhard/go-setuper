package apt

import (
	"github.com/ui-kreinhard/go-setuper/utils"
	// "github.com/arduino/go-apt-client"
	"github.com/ui-kreinhard/go-apt-client"
)

func Package(name ...string) []*apt.Package {
	ret := []*apt.Package{}
	for _, packageName := range name {
		ret = append(ret, &apt.Package{
			packageName,
			"",
			"",
			"",
			"",
			0,
		})
	}
	return ret
}

func CheckForUpdatesDirect() (string, error) {
	output, err := utils.ConvertOutput(apt.CheckForUpdates())
	return output, err
}

func InstallDirect(packageName ...string) (string, error) {
	output, err := utils.ConvertOutput(apt.Install(Package(packageName...)...))
	return output, err
}

func RemoveDirect(packageName ...string) (string, error) {
	output, err := utils.ConvertOutput(apt.Remove(Package(packageName...)...))
	return output, err
}

func AddAptRepositoryDirect(url string) (string, error) {
	output, err := utils.Exec("add-apt-repository", url)
	return output, err
}

func AddAptKeyDirect(urlOrFile string) (string, error) {
	output, err := utils.Exec("wget", urlOrFile, "-O", "/tmp/aptKey")
	if err != nil {
		return output, err
	}
	output, err = utils.Exec("apt-key", "add", "/tmp/aptKey")

	return output, err
}
