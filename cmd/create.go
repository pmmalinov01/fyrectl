/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/pmmalinov01/fyrectl/pkg/creds"
	"github.com/pmmalinov01/fyrectl/pkg/utils"
	"github.com/spf13/cobra"
	"sigs.k8s.io/yaml"
)

const (
	createFyreCluster = "https://api.fyre.ibm.com/rest/v1/?operation=build"
)

var ConfFile string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		content, err := ioutil.ReadFile(ConfFile)
		if err != nil {
			log.Fatal(err)
		}
		j2, err := yaml.YAMLToJSON(content)
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		createCluster := strings.NewReader(string(j2))

		req, err := http.NewRequest("POST", createFyreCluster, createCluster)
		if err != nil {
			log.Fatal(err)
		}

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
		var prettyJSON bytes.Buffer
		error := json.Indent(&prettyJSON, data, "", "\t")
		if error != nil {
			log.Println("JSON parse error", error)

		}
		fmt.Println(string(prettyJSON.Bytes()))
		utils.CreateClStatus(data)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&ConfFile, "conf", "c", "", "Passes the configuration to create the fyre cluster")
	createCmd.MarkFlagRequired("conf")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
