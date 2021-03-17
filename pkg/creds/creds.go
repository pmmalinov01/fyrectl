package creds

import (
	"errors"
	"os"
)

type userCreds struct {
	UserName string
	ApiKey   string
}

func GetCreds() (userCreds, error) {

	var user userCreds
	un, ok := os.LookupEnv("FYRE_USER")
	if !ok {
		return user, errors.New("Variable FYRE_USER not set")
	}

	ps, ok := os.LookupEnv("FYRE_API_KEY")
	if !ok {
		return user, errors.New("Variable FYRE_API_KEY not set")
	}
	user.ApiKey = ps
	user.UserName = un
	return user, nil
}
