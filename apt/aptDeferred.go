package apt

func CheckForUpdatesDeferred() func() (string, error) {
	return func() (string, error) {
		return CheckForUpdates()
	}
}

func InstallDeferred(packageName ...string) func() (string, error) {
	return func() (string, error) {
		return Install(packageName...)
	}
}

func RemoveDeferred(packageName ...string) func() (string, error) {
	return func() (string, error) {
		return Remove(packageName...)
	}
}

func AddAptRepositoryDeferred(url string) func() (string, error) {
	return func() (string, error) {
		return AddAptRepository(url)
	}
}

func AddAptKeyDeferred(urlOrFile string) func() (string, error) {
	return func() (string, error) {
		return AddAptKey(urlOrFile)
	}
}
