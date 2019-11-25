package timezones

func SetTimezone(newTimezone Timezone) func() (string, error) {
	return func() (string, error) {
		return SetTimezoneDirect(newTimezone)
	}
}