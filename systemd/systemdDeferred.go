package systemd

func Enable(serviceName string) func() (string, error) {
	return func() (string, error) {
		return EnableDirect(serviceName)
	}
}

func Disable(serviceName string) func() (string, error) {
	return func() (string, error) {
		return DisableDirect(serviceName)
	}
}

func Start(serviceName string) func() (string, error) {
	return func() (string, error) {
		return StartDirect(serviceName)
	}
}

func Stop(serviceName string) func() (string, error) {
	return func() (string, error) {
		return StopDirect(serviceName)
	}
}

func Restart(serviceName string) func() (string, error) {
	return func() (string, error) {
		return RestartDirect(serviceName)
	}
}

func DaemonReload() func() (string, error) {
	return func() (string, error) {
		return DaemonReloadDirect()
	}
}
