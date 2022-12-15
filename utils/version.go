package utils

import (
	"log"
	"os"
)

func GetCurrentVersion() string {
	versionFile, err := os.Open(".version")
	if err != nil {
		log.Fatalf(err.Error())
	}

	version, _, err := ReadLine(versionFile, 1)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = versionFile.Close()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return version
}
