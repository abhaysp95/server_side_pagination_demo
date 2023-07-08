package main

import (
	"encoding/json"
	"log"
	"os"
)

func deserializePurchaseDetails(filename string) []UserPurchase {
	jdata, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("can't read sample json data")
	}

	var parsedData []UserPurchase
	err = json.Unmarshal([]byte(jdata), &parsedData)
	if (err != nil) {
		log.Fatal("error in deserializing json data", err)
	}

	return parsedData
}
