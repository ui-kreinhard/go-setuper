package template

import (
	"github.com/ui-kreinhard/go-setuper/setuper"
	"os"
	"text/template"
)

func Render(filename string, targetFilename string, templateObject interface{}) error {
	templateContent, err := setuper.NewSetuper().TemplatesBox.FindString(filename)
	if err != nil {
		return err
	}

	file, err := os.Create(targetFilename)
	if err != nil {
		return err
	}

	parsedTemplate, err := template.New(filename).Parse(templateContent)
	if err != nil {
		return err
	}
	err = parsedTemplate.Execute(file, templateObject)
	return err
}
