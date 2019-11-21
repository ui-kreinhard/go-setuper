package setuper

import (
	"fmt"
	"github.com/gobuffalo/packr/v2"
	"log"
)

type ISetuper interface{}

type Setuper struct {
	TemplatesBox *packr.Box
	FilesBox     *packr.Box
	ScriptBox    *packr.Box
}

var setuper *Setuper
var pathPrefix string

func printContent(box *packr.Box) {
	for _, file := range box.List() {
		log.Println(file)
	}
}

func NewSetuper() *Setuper {
	if setuper == nil {
		fmt.Println("Instanciating setuper")
		setuper = &Setuper{
			packr.New("templates", "../staticAssets/templates/"),
			packr.New("files", "../staticAssets/file"),
			packr.New("scripts", "../staticAssets/scripts"),
		}

		printContent(setuper.FilesBox)
		printContent(setuper.TemplatesBox)
		printContent(setuper.ScriptBox)
	} else {
		fmt.Println("Returning cached setuper")
	}
	return setuper
}

func ConfigurePrefix(prefix string) {
	pathPrefix = prefix
	NewSetuper()
}
