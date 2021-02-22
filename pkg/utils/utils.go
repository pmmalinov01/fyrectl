package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

var PrettyJSONF bytes.Buffer

func PrettyJSON(bJSON []byte) {
	error := json.Indent(&PrettyJSONF, bJSON, "", "  ")
	if error != nil {
		log.Println("JSON parse error", error)

	}
	fmt.Println(string(PrettyJSONF.Bytes()))
}

type ClusterC struct {
	Status    string `json:"status"`
	RequestID string `json:"request_id"`
	Details   string `json:"details"`
}

func (c *ClusterC) GetStatus() string {
	return c.Status
}

func (c *ClusterC) GetRID() string {
	return c.RequestID
}

func (c *ClusterC) GetDetailsURL() string {
	return c.Details
}

func CreateClStatus(Url []byte) {
	var cluster ClusterC
	isValid := json.Valid(Url)
	fmt.Println(isValid)
	fmt.Printf("Error: %v\n", json.Unmarshal(Url, &cluster))

	fmt.Println(cluster.GetStatus())
	fmt.Println(cluster.GetDetailsURL())
}
