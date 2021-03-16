package utils

import (
	"bytes"
	"encoding/json"
	"errors"
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

func CreateClStatus(Url []byte) (string, error) {
	var cluster ClusterC
	isValid := json.Valid(Url)
	if !isValid {
		return "", errors.New("JSON is not Valid")
	}
	json.Unmarshal(Url, &cluster)
	return cluster.GetRID(), nil
}
