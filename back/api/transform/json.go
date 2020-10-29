package transform

import (
	"encoding/json"
	"log"
)

func ToJson(resource interface{}) []byte {
	marshaled, err := json.Marshal(resource)
	if err != nil {
		log.Panic(err)
	}
	return marshaled
}
