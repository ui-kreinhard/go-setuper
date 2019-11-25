package template

func Render(filename, targetFilename string, templateObject interface{}) func() (string, error) {
	return func() (string, error) {
		return "", RenderDirect(filename, targetFilename, templateObject)
	}
}
