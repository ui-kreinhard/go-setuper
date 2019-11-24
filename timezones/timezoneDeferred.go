package timezones

func SetTimezoneDeferred(newTimezone Timezone) func() (string, error) {
	return func() (string, error) {
		return SetTimezone(newTimezone)
	}
}