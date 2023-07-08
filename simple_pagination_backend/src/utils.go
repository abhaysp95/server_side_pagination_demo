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

func serializePurchaseData(data []UserPurchase) []byte {
	res, err := json.Marshal(data)
	if err != nil {
		log.Fatal("error in serializing json data", err)
	}

	return res
}

// for now a page is of 10 rows
func getPurchaseDataForPage(page int) []UserPurchase {
	start := 10 * (page - 1)
	end := 10 * page

	return purchaseData[start:end]
}

