package writer

import (
	"math"
	"math/rand"

	"github.com/sundywy/stock-kinesis/model"
	"go.uber.org/atomic"
)

type stockPrice struct {
	TickerSymbol string
	Price        float64
}

type StockTradeGenerator []*stockPrice

func NewGenerator() StockTradeGenerator {
	data := map[string]float64{
		"AAPL":  119.72,
		"XOM":   91.56,
		"GOOG":  527.83,
		"BRK.A": 223999.88,
		"MSFT":  42.36,
		"WFC":   54.21,
		"JNJ":   99.78,
		"WMT":   85.91,
		"CHL":   66.96,
		"GE":    24.64,
		"NVS":   102.46,
		"PG":    85.05,
		"JPM":   57.82,
		"RDS.A": 66.72,
		"CVX":   110.43,
		"PFE":   33.07,
		"FB":    74.44,
		"VZ":    49.09,
		"PTR":   111.08,
		"BUD":   120.39,
		"ORCL":  43.40,
		"KO":    41.23,
		"T":     34.64,
		"DIS":   101.73,
		"AMZN":  370.56,
	}

	var g StockTradeGenerator
	for k, v := range data {
		g = append(g, &stockPrice{TickerSymbol: k, Price: v})
	}

	return g
}

const (
	maxDeviation    = 0.2
	maxQuantity     = 10000
	probabilitySell = 0.4
)

var counter = atomic.NewInt32(0)

func (g StockTradeGenerator) GetRandomTrade() *model.StockTrade {

	stockPrice := g[rand.Intn(len(g))]

	deviation := (rand.Float64() - 0.5) * 2.0 * maxDeviation
	price := stockPrice.Price * (1 + deviation)
	price = math.Round(price*100.0) / 100.0

	var tradeType model.TradeType

	if rand.Float64() < probabilitySell {
		tradeType = model.Sell
	} else {
		tradeType = model.Buy
	}

	quantity := rand.Intn(maxQuantity) + 1

	return model.New(stockPrice.TickerSymbol, tradeType, price, quantity, int(counter.Inc()))
}
