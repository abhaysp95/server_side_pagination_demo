package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var purchaseData []UserPurchase

func main() {
	mux := chi.NewRouter()
	mux.Get("/page", specificPageData)

	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal("Problem starting server")
	}
	log.Printf("server started at :4000")
}

func init() {
	log.Printf("init ran")
	purchaseData = deserializePurchaseDetails("./assets/data/car_purchase_sample_data.json")
}
