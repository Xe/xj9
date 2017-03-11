package common

import "github.com/tchap/go-exchange/exchange"

// Exc is the global exchange plugins will use
var Exc *exchange.Exchange

func init() {
	Exc = exchange.New()
}
