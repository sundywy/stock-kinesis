package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/sundywy/stock-kinesis/writer"
)

const (
	streamName = "StockTradeStream"
	region     = "ap-southeast-1"
)

func main() {
	writer := getNewWriter()
	if err := writer.ValidateStream(); err != nil {
		log.Fatalf(`Error in validating stream. %v`, err)
	}

	for {
		writer.SendStockTrade()
	}
}

func getNewWriter() *writer.StockTradesWriter {

	config := aws.NewConfig().
		WithCredentials(credentials.NewEnvCredentials()).
		WithRegion(region)

	session, _ := session.NewSession(config)
	kinesis := kinesis.New(session)

	return writer.NewWriter(&writer.Config{Kinesis: kinesis, StreamName: streamName})

}
