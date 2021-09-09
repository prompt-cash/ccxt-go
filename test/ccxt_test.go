package test

import (
	"fmt"
	"github.com/prompt-cash/js-transpiler/pkg/ccxt"
	"testing"
)

//
// example test
//

func TestParseScript(t *testing.T) {

	config := ccxt.MkMap(&ccxt.VarMap{
		"apiKey":          ccxt.MkString("***todo my api key***"),
		"secret":          ccxt.MkString("***todo my secret***"),
		"enableRateLimit": ccxt.MkBool(true),
		"timeout":         ccxt.MkInteger(10 * 1000), // in ms
	})
	// some exchanges have additional config (example FTX sub-account IDs)

	binance := &ccxt.Binance{}
	binance.ExchangeBase = &ccxt.ExchangeBase{}
	binance.Setup(config, binance)

	markets := binance.LoadMarkets() // load available markets for requests below
	fmt.Println(string(ccxt.VariantToJson(markets)))

	// TODO try the following
	//ret := binance.CreateOrder(MkString("BTC/USDT"), MkString("limit"), MkString("buy"), MkNumber(0.001), MkInteger(50000), MkMap(&VarMap{})) // buy and sell
	//binance.CancelOrder()
	//binance.FetchOpenOrders() // my orders
	//binance.FetchOrderBook()  // all public orders (bid + ask)

	ret := binance.FetchBalance(ccxt.MkMap(&ccxt.VarMap{}))

	fmt.Println(string(ccxt.VariantToJson(ret)))
	return
}
