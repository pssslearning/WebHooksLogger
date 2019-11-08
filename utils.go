package main

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

func getFileName(httpMethod string) (string, error) {

	tstmp := getTimeStampAsString()
	uuidBA, err := getUUID()

	if err != nil {
		return "", err
	}
	return tstmp + "-" + httpMethod + "-" + uuidBA, nil

}

func getTimeStampAsString() string {

	t := time.Now()
	replacer := strings.NewReplacer(" ", "-", ":", "-", "+", "-")
	return replacer.Replace(t.Format(time.RFC3339))
}

func getUUID() (string, error) {
	uuidBA, err := uuid.NewRandom()

	if err != nil {
		return "", err
	}
	return uuidBA.String(), nil
}
