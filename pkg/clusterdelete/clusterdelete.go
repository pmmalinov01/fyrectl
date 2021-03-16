package clusterdelete

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/pmmalinov01/fyrectl/pkg/creds"
	"github.com/pmmalinov01/fyrectl/pkg/utils"
)

var ClusterDelete string = "https://api.fyre.ibm.com/rest/v1/?operation=delete"

func DeleteCl(ClName string) {
	r := map[string]string{"cluster_name": ClName}
	reqBodyJSON, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reqBodyJSON)

	reqBody1 := strings.NewReader(string(reqBodyJSON))
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	req, _ := http.NewRequest("POST", ClusterDelete, reqBody1)
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
