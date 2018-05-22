package main

import (
	"log"

	"github.com/sundywy/stock-kinesis/writer"
)

const (
	streamName = "StockTradeStream"
	region     = "ap-southeast-1"
)

func main() {

	writer, err := writer.NewWriter(streamName, region)
	if err != nil {
		log.Fatalf(`Error in creating writer. %v`, err)
	}

	if err := writer.ValidateStream(); err != nil {
		log.Fatalf(`Error in validating stream. %v`, err)
	}

	for {
		writer.SendStockTrade()
	}
}
