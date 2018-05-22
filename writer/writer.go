package writer

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

type Config struct {
	*kinesis.Kinesis
	StreamName string
}

type StockTradesWriter struct {
	*Config
	generator StockTradeGenerator
}

func NewWriter(streamName, region string) (*StockTradesWriter, error) {

	awsConfig := aws.NewConfig().
		WithCredentials(credentials.NewEnvCredentials()).
		WithRegion(region)

	session, err := session.NewSession(awsConfig)
	if err != nil {
		return nil, fmt.Errorf(`Unable to create session. %v`, err)
	}

	kinesis := kinesis.New(session)

	return &StockTradesWriter{&Config{Kinesis: kinesis, StreamName: streamName}, NewGenerator()}, nil
}

func (w *StockTradesWriter) SendStockTrade() error {

	stockTrade := w.generator.GetRandomTrade()

	input := (&kinesis.PutRecordInput{}).
		SetData(stockTrade.AsBytes()).
		SetPartitionKey(stockTrade.TickerSymbol).
		SetStreamName(w.StreamName)

	if err := input.Validate(); err != nil {
		return err
	}

	_, err := w.PutRecord(input)
	return err
}

func (w *StockTradesWriter) ValidateStream() error {
	input := (&kinesis.DescribeStreamInput{}).
		SetStreamName(w.StreamName)

	if err := input.Validate(); err != nil {
		return err
	}

	_, err := w.DescribeStream(input)
	return err
}
