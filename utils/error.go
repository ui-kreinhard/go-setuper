package utils

import (
	"log"
)

func PanicOnError(output string, err error) (string, error) {
	if err != nil {
		log.Fatal("[SETUPER unrecoverable error]", err, output)
	}
	return output, err
}
