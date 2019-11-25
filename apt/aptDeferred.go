package apt

func CheckForUpdates() func() (string, error) {
	return func() (string, error) {
		return CheckForUpdatesDirect()
	}
}

func Install(packageName ...string) func() (string, error) {
	return func() (string, error) {
		return InstallDirect(packageName...)
	}
}

func Remove(packageName ...string) func() (string, error) {
	return func() (string, error) {
		return RemoveDirect(packageName...)
	}
}

func AddAptRepository(url string) func() (string, error) {
	return func() (string, error) {
		return AddAptRepositoryDirect(url)
	}
}

func AddAptKey(urlOrFile string) func() (string, error) {
	return func() (string, error) {
		return AddAptKeyDirect(urlOrFile)
	}
}
