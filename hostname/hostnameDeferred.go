package hostname

func SetHostname(hostname string) func() (string, error) {
	return func() (string, error) {
		return SetHostnameDirect(hostname)
	}
}
