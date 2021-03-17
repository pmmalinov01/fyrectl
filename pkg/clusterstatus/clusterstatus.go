package clusterstatus

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/pmmalinov01/fyrectl/pkg/creds"
	"github.com/pmmalinov01/fyrectl/pkg/utils"
)

var ClusterStatus string = "https://api.fyre.ibm.com/rest/v1/?operation=query&request=showrequests&request_id="

//ClientR returns a details about a made call to the API
//It can be from a simple delete or from create call
func ClientR(reqID string) {
	reqBody1 := strings.NewReader(``)
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	req, _ := http.NewRequest("GET", ClusterStatus+reqID, reqBody1)
	uCreds, err := creds.GetCreds()
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(uCreds.UserName, uCreds.ApiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	data, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	utils.PrettyJSON(data)
}
