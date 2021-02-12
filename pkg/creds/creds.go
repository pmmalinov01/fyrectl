package creds

import (
	"errors"
	"os"
)

type userCreds struct {
	userName string
	apiKey   string
}

func GetCreds() (userCreds, error) {

	var user userCreds
	un, ok := os.LookupEnv("FYRE_USER")
	if !ok {
		return user, errors.New("Variable USERNAME not set")
	}

	ps, ok := os.LookupEnv("APIKEY")
	if !ok {
		return user, errors.New("Variable APIKEY not set")
	}
	user.apiKey = ps
	user.userName = un
	return user, nil
}
