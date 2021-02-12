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
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

const (
	statusFyreCluster = "https://api.fyre.ibm.com/rest/v1/?operation=query&request=showclusters"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List vm clusters",
	Long: `
	List your stacks
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
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
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
