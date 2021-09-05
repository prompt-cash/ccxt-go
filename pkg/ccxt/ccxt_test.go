package ccxt

import (
	"fmt"
	"testing"
)

func TestParseScript(t *testing.T) {

	config := MkMap(&VarMap{
		"apiKey":          MkString("***todo my api key***"),
		"secret":          MkString("***todo my secret***"),
		"enableRateLimit": MkBool(true),
		"timeout":         MkInteger(10 * 1000), // in ms
	})
	// some exchanges have additional config (example FTX sub-account IDs)

	binance := &Binance{}
	binance.ExchangeBase = &ExchangeBase{}
	binance.Setup(config, binance)

	markets := binance.LoadMarkets() // load available markets for requests below
	fmt.Println(string(VariantToJson(markets)))

	// TODO try the following
	//ret := binance.CreateOrder(MkString("BTC/USDT"), MkString("limit"), MkString("buy"), MkNumber(0.001), MkInteger(50000), MkMap(&VarMap{})) // buy and sell
	//binance.CancelOrder()
	//binance.FetchOpenOrders() // my orders
	//binance.FetchOrderBook()  // all public orders (bid + ask)

	ret := binance.FetchBalance(MkMap(&VarMap{}))

	fmt.Println(string(VariantToJson(ret)))
	return
}
