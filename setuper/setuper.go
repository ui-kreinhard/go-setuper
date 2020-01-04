package setuper

import (
	"github.com/gobuffalo/packr/v2"
	"log"
)

type ISetuper interface{}

type Setuper struct {
	FilesBox     *packr.Box
	TemplatesBox *packr.Box
	ScriptBox    *packr.Box
}

var setuper *Setuper
var pathPrefix *string

var filesBox *packr.Box
var templatesBox *packr.Box
var scriptsBox *packr.Box

func printContent(box *packr.Box) {
	for _, file := range box.List() {
		log.Println(file)
	}
}

func NewSetuper() *Setuper {
	if setuper == nil {
		setuper = &Setuper{
			filesBox,
			templatesBox,
			scriptsBox,
		}
		log.Println("Known Files")
		printContent(setuper.FilesBox)
		log.Println()
		log.Println("Known Templates")
		printContent(setuper.TemplatesBox)
		log.Println()
		log.Println("Known scripts")
		printContent(setuper.ScriptBox)
		log.Println()
	} 
	return setuper
}

func ConfigureBox(files, templates, scripts *packr.Box) {
	filesBox = files
	templatesBox = templates
	scriptsBox = scripts
	NewSetuper()
}
