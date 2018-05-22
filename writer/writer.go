package writer

import (
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

func NewWriter(config *Config) *StockTradesWriter {

	return &StockTradesWriter{config, NewGenerator()}
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
