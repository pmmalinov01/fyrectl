package masscot

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var urlFyre = "https://api.fyre.ibm.com/rest/v1/?operation=status"
var statusFyreCluster = "https://api.fyre.ibm.com/rest/v1/?operation=query&request=showclusters"

// TODO: read from env
var userName = "pavel.malinov"
var apiKey = "gCMN6CMtyNxWHu52kMfiTCddIkOc37gjJ36PFTjeG"

//FureStatus returns the status of the fyre clusters for user account
func FyreStatus() {
	reqBody1 := strings.NewReader(``)
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	req, _ := http.NewRequest("GET", statusFyreCluster, reqBody1)
	req.SetBasicAuth(userName, apiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	data, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Println(res.StatusCode)
	fmt.Printf("%s", data)
}
