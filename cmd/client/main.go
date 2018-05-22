package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sundywy/stock-kinesis/reader"
)

const (
	streamName = "StockTradeStream"
	region     = "ap-southeast-1"
)

func main() {

	reader, err := reader.NewReader(streamName, region)
	if err != nil {
		log.Fatalf(`Unable to create new reader. %v`, err)
	}

	r := chi.NewRouter()
	r.Get("/stocks", func(w http.ResponseWriter, r *http.Request) {
		stock, err := reader.GetStockTrade()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf(`{"error": "%v"}`, err)))
		}

		w.WriteHeader(http.StatusOK)
		w.Write(stock.AsBytes())
	})

	log.Fatal(http.ListenAndServe(":8080", r))

}
