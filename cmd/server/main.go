package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sundywy/stock-kinesis/writer"
)

func main() {

	g := writer.NewGenerator()
	r := chi.NewRouter()
	r.Get("/stock", func(w http.ResponseWriter, r *http.Request) {
		stock := g.GetRandomTrade()
		w.WriteHeader(http.StatusOK)
		w.Write(stock.AsBytes())
	})
	log.Fatal(http.ListenAndServe(":8080", r))
}
