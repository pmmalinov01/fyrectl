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

type ClusterD struct {
	Node            string `json:"node"`
	CPU             string `json:"cpu"`
	Memory          string `json:"memory"`
	Publicip        string `json:"publicip"`
	Privateip       string `json:"privateip"`
	AdditionalDisks string `json:"additional_disks"`
	Status          string `json:"status"`
	InstanceType    string `json:"instance_type"`
	Platform        string `json:"platform"`
	State           string `json:"state"`
}

func (cd *ClusterD) GetNode() string {
	return cd.Node
}
func (cd *ClusterD) GetCPU() string {
	return cd.CPU
}
func (cd *ClusterD) GetMemory() string {
	return cd.Memory
}

func (cd *ClusterD) GetPublicIP() string {
	return cd.Publicip
}

func (cd *ClusterD) GetPrivateIP() string {
	return cd.Privateip
}

func (cd *ClusterD) GetStatus() string {
	return cd.Status
}

func (cd *ClusterD) GetState() string {
	return cd.State
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

func ClusterDetails(Res []byte) (string, error) {
	var clusterD ClusterD
	json.Unmarshal(Res, &clusterD)
	return clusterD.GetPublicIP(), nil

}
