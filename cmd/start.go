package start

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var urlFyre = "https://api.fyre.ibm.com/rest/v1/?operation=status"
var statusFyreCluster = "https://api.fyre.ibm.com/rest/v1/?operation=query&request=showclusters"

// TODO: read from env
type userCreds struct {
	userName string
	apiKey   string
}

func GetCreds() (userCreds, error) {

	var user userCreds
	un, ok := os.LookupEnv("USERNAME")
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

//FureStatus returns the status of the fyre clusters for user account
func FyreStatus() {
	reqBody1 := strings.NewReader(``)
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	req, _ := http.NewRequest("GET", statusFyreCluster, reqBody1)
	uCreds, err := GetCreds()
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(uCreds.userName, uCreds.apiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	data, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Println(res.StatusCode)
	fmt.Printf("%s", data)
}
