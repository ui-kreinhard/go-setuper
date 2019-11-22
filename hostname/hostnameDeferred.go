package hostname

func SetHostnameDeferred(hostname string) func() (string, error) {
	return func() (string, error) {
		return SetHostname(hostname)
	}
}
