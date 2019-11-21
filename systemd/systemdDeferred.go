package systemd

func EnableDeferred(serviceName string) func() (string, error) {
	return func() (string, error) {
		return Enable(serviceName)
	}
}

func DisableDeferred(serviceName string) func() (string, error) {
	return func() (string, error) {
		return Disable(serviceName)
	}
}

func StartDeferred(serviceName string) func() (string, error) {
	return func() (string, error) {
		return Start(serviceName)
	}
}

func StopDeferred(serviceName string) func() (string, error) {
	return func() (string, error) {
		return Stop(serviceName)
	}
}

func RestartDeferred(serviceName string) func() (string, error) {
	return func() (string, error) {
		return Restart(serviceName)
	}
}

func DaemonReloadDeferred() func() (string, error) {
	return func() (string, error) {
		return DaemonReload()
	}
}
