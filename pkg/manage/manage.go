package manage

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/pmmalinov01/fyrectl/pkg/creds"
	"github.com/pmmalinov01/fyrectl/pkg/utils"
)

var (
	ClusterBoot   string = "https://api.fyre.ibm.com/rest/v1/?operation=boot&cluster_name="
	ClusterReBoot string = "https://api.fyre.ibm.com/rest/v1/?operation=reboot&cluster_name="
	ClusterStop   string = "https://api.fyre.ibm.com/rest/v1/?operation=shutdown&cluster_name="
)

func clientcall(op string) {

	reqBody1 := strings.NewReader(``)
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, _ := http.NewRequest("GET", op, reqBody1)
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

func State(cluster_name, operation string) {

	switch operation {
	case "boot":
		clientcall(ClusterBoot + cluster_name)
	case "reboot":
		clientcall(ClusterReBoot + cluster_name)
	case "stop":
		clientcall(ClusterStop + cluster_name)
	case "force_stop":
		out := fmt.Sprintf("api.fyre.ibm.com/rest/v1/?operation=shutdown&cluster_name=%s[&shutdown_type=pull_the_plug", cluster_name)
		clientcall(out)
	}

}
