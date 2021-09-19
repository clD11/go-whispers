package shared

import "log"

type logger struct{}

func (l logger) info(message string) {
	log.Print(message)
}

func (l logger) error(message string) {
	log.Print(message)
}
