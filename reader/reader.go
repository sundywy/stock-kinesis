package reader

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/sundywy/stock-kinesis/model"
)

type Config struct {
	*kinesis.Kinesis
	StreamName string
}

type StockTradesReader struct {
	*Config
}

func NewReader(streamName, region string) (*StockTradesReader, error) {
	awsConfig := aws.NewConfig().
		WithCredentials(credentials.NewSharedCredentials("", "")).
		WithRegion(region)

	session, err := session.NewSession(awsConfig)
	if err != nil {
		return nil, fmt.Errorf(`Unable to create session. %v`, err)
	}

	kinesis := kinesis.New(session)

	return &StockTradesReader{&Config{Kinesis: kinesis, StreamName: streamName}}, nil
}

func (r *StockTradesReader) GetStockTrade() (*model.StockTrade, error) {

	describeStreamInput := (&kinesis.DescribeStreamInput{}).
		SetStreamName(r.StreamName)

	streamOutput, err := r.DescribeStream(describeStreamInput)
	if err != nil {
		return nil, err
	}

	shardID := *(streamOutput.StreamDescription.Shards[0]).ShardId

	getShardIteratorInput := (&kinesis.GetShardIteratorInput{}).
		SetShardId(shardID).
		SetShardIteratorType("LATEST").
		SetStreamName(r.StreamName)

	if err := getShardIteratorInput.Validate(); err != nil {
		return nil, err
	}

	shardIterator, err := r.GetShardIterator(getShardIteratorInput)
	if err != nil {
		return nil, err
	}

	input := (&kinesis.GetRecordsInput{}).
		SetLimit(1).
		SetShardIterator(*shardIterator.ShardIterator)

	if err := input.Validate(); err != nil {
		return nil, err
	}

	output, err := r.GetRecords(input)
	if err != nil {
		return nil, err
	}

	return model.NewFromBytes(output.Records[0].Data), nil
}
