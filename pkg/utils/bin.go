package utils

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// MakeBins creates the local bin directories
func MakeBins() {
	err := os.MkdirAll("./bin/listener", os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	err = os.MkdirAll("./bin/middleware", os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	err = os.MkdirAll("./bin/roundtrip", os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}
