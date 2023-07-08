package main

import (
	"log"
	"net/http"
	"strconv"
)

func specificPageData(w http.ResponseWriter, r *http.Request) {
	count, err := strconv.Atoi(r.URL.Query().Get("count"))
	if err != nil {
		log.Println("error in parsing query")
	}

	data := serializePurchaseData(getPurchaseDataForPage(count))
	w.Write(data)
}
