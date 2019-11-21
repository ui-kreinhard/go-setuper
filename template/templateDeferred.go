package template

func RenderDeferred(filename, targetFilename string, templateObject interface{}) func() (string, error) {
	return func() (string, error) {
		return "", Render(filename, targetFilename, templateObject)
	}
}
