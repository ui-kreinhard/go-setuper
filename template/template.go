package template

import (
	"github.com/ui-kreinhard/go-setuper/setuper"
	"github.com/ui-kreinhard/go-setuper/utils"
	"os"
	"text/template"
)

func Render(filename string, targetFilename string, templateObject interface{}) error {
	constructedDest, err := utils.ConstructDestination(filename, targetFilename)
	if err != nil {
		return err
	}
	templateContent, err := setuper.NewSetuper().TemplatesBox.FindString(filename)
	if err != nil {
		return err
	}

	file, err := os.Create(constructedDest)
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
