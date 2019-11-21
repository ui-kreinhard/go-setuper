package log

import (
	"log"
	"strings"
)

func tabOutput(output string) string {
	return strings.ReplaceAll("\n"+output, "\n", "\n\t")
}

func LogOutputError(output string, err error) {
	if err != nil {
		log.Println("[SETUPER ERROR OUTPUT BEGIN]")
		log.Println(tabOutput(output))
		log.Println("[SETUPER ERROR]", err)
		log.Println("[SETUPER ERROR OUTPUT END]")
	} else {
		if len(output) > 0 {
			log.Println("[SETUPER OUTPUT BEGIN]")
			log.Println(tabOutput(output))
			log.Println("[SETUPER OUTPUT END]")

		}
	}
}

func Println(toPrint ...interface{}) {
	log.Println(toPrint...)
}
