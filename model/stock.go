package model

import (
	"encoding/json"
	"fmt"
)

type TradeType int

const (
	Buy TradeType = iota
	Sell
)

func (t TradeType) String() (s string) {
	switch t {
	case Buy:
		s = "Buy"
	case Sell:
		s = "Sell"
	}
	return
}

type StockTrade struct {
	TickerSymbol string    `json:"ticker_symbol"`
	TradeType    TradeType `json:"trade_type"`
	Price        float64   `json:"price"`
	Quantity     int       `json:"quantity"`
	ID           int       `json:"id"`
}

func New(tickerSymbol string, tradeType TradeType, price float64, quantity, id int) *StockTrade {
	return &StockTrade{
		TickerSymbol: tickerSymbol,
		TradeType:    tradeType,
		Price:        price,
		Quantity:     quantity,
		ID:           id,
	}
}

func NewFromBytes(data []byte) *StockTrade {
	var s StockTrade

	if err := json.Unmarshal(data, &s); err != nil {
		return nil
	}
	return &s
}

func (s *StockTrade) String() string {

	return fmt.Sprintf(`ID: %d, %s %d shares of %s as of %.02f`,
		s.ID, s.TradeType.String(), s.Quantity, s.TickerSymbol, s.Price)
}

func (s *StockTrade) AsBytes() []byte {

	bytes, err := json.Marshal(s)
	if err != nil {
		return []byte{}
	}

	return bytes
}
