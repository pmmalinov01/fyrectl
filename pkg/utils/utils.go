package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

var PrettyJSON bytes.Buffer

func PrettyJSON(bJSON []byte) {
	error := json.Indent(&PrettyJSON, bJSON, "", "\t")
	if error != nil {
		log.Println("JSON parse error", error)

	}
	fmt.Println(string(PrettyJSON.Bytes()))
}
