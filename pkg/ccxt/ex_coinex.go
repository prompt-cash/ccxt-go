package ccxt

import ()

type Coinex struct {
	*ExchangeBase
}

var _ Exchange = (*Coinex)(nil)

func init() {
	exchange := &Coinex{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Coinex) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":      MkString("coinex"),
		"name":    MkString("CoinEx"),
		"version": MkString("v1"),
		"countries": MkArray(&VarArray{
			MkString("CN"),
		}),
		"rateLimit": MkInteger(1000),
		"has": MkMap(&VarMap{
			"cancelOrder":       MkBool(true),
			"createOrder":       MkBool(true),
			"fetchBalance":      MkBool(true),
			"fetchClosedOrders": MkBool(true),
			"fetchDeposits":     MkBool(true),
			"fetchMarkets":      MkBool(true),
			"fetchMyTrades":     MkBool(true),
			"fetchOHLCV":        MkBool(true),
			"fetchOpenOrders":   MkBool(true),
			"fetchOrder":        MkBool(true),
			"fetchOrderBook":    MkBool(true),
			"fetchTicker":       MkBool(true),
			"fetchTickers":      MkBool(true),
			"fetchTrades":       MkBool(true),
			"fetchWithdrawals":  MkBool(true),
			"withdraw":          MkBool(true),
		}),
		"timeframes": MkMap(&VarMap{
			"1m":  MkString("1min"),
			"3m":  MkString("3min"),
			"5m":  MkString("5min"),
			"15m": MkString("15min"),
			"30m": MkString("30min"),
			"1h":  MkString("1hour"),
			"2h":  MkString("2hour"),
			"4h":  MkString("4hour"),
			"6h":  MkString("6hour"),
			"12h": MkString("12hour"),
			"1d":  MkString("1day"),
			"3d":  MkString("3day"),
			"1w":  MkString("1week"),
		}),
		"urls": MkMap(&VarMap{
			"logo": MkString("https://user-images.githubusercontent.com/51840849/87182089-1e05fa00-c2ec-11ea-8da9-cc73b45abbbc.jpg"),
			"api": MkMap(&VarMap{
				"public":           MkString("https://api.coinex.com"),
				"private":          MkString("https://api.coinex.com"),
				"perpetualPublic":  MkString("https://api.coinex.com/perpetual"),
				"perpetualPrivate": MkString("https://api.coinex.com/perpetual"),
			}),
			"www":      MkString("https://www.coinex.com"),
			"doc":      MkString("https://github.com/coinexcom/coinex_exchange_api/wiki"),
			"fees":     MkString("https://www.coinex.com/fees"),
			"referral": MkString("https://www.coinex.com/register?refer_code=yw5fz"),
		}),
		"api": MkMap(&VarMap{
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("common/currency/rate"),
				MkString("common/asset/config"),
				MkString("market/info"),
				MkString("market/list"),
				MkString("market/ticker"),
				MkString("market/ticker/all"),
				MkString("market/depth"),
				MkString("market/deals"),
				MkString("market/kline"),
			})}),
			"private": MkMap(&VarMap{
				"get": MkArray(&VarArray{
					MkString("balance/coin/deposit"),
					MkString("balance/coin/withdraw"),
					MkString("balance/info"),
					MkString("future/account"),
					MkString("future/config"),
					MkString("future/limitprice"),
					MkString("future/loan/history"),
					MkString("future/market"),
					MkString("margin/account"),
					MkString("margin/config"),
					MkString("margin/loan/history"),
					MkString("margin/market"),
					MkString("order"),
					MkString("order/deals"),
					MkString("order/finished"),
					MkString("order/finished/{id}"),
					MkString("order/pending"),
					MkString("order/status"),
					MkString("order/status/batch"),
					MkString("order/user/deals"),
					MkString("sub_account/balance"),
					MkString("sub_account/transfer/history"),
				}),
				"post": MkArray(&VarArray{
					MkString("balance/coin/withdraw"),
					MkString("future/flat"),
					MkString("future/loan"),
					MkString("future/transfer"),
					MkString("margin/flat"),
					MkString("margin/loan"),
					MkString("margin/transfer"),
					MkString("order/batchlimit"),
					MkString("order/ioc"),
					MkString("order/limit"),
					MkString("order/market"),
					MkString("sub_account/transfer"),
				}),
				"delete": MkArray(&VarArray{
					MkString("balance/coin/withdraw"),
					MkString("order/pending/batch"),
					MkString("order/pending"),
				}),
			}),
			"perpetualPublic": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("ping"),
				MkString("time"),
				MkString("market/list"),
				MkString("market/limit_config"),
				MkString("market/ticker"),
				MkString("market/ticker/all"),
				MkString("market/depth"),
				MkString("market/deals"),
				MkString("market/funding_history"),
				MkString("market/user_deals"),
				MkString("market/kline"),
			})}),
			"perpetualPrivate": MkMap(&VarMap{
				"get": MkArray(&VarArray{
					MkString("asset/query"),
					MkString("order/pending"),
					MkString("order/finished"),
					MkString("order/stop_pending"),
					MkString("order/status"),
					MkString("position/pending"),
					MkString("position/funding"),
				}),
				"post": MkArray(&VarArray{
					MkString("market/adjust_leverage"),
					MkString("market/position_expect"),
					MkString("order/put_limit"),
					MkString("order/put_market"),
					MkString("order/put_stop_limit"),
					MkString("order/cancel"),
					MkString("order/cancel_all"),
					MkString("order/cancel_stop"),
					MkString("order/cancel_stop_all"),
					MkString("order/close_limit"),
					MkString("order/close_market"),
					MkString("position/adjust_margin"),
				}),
			}),
		}),
		"fees": MkMap(&VarMap{
			"trading": MkMap(&VarMap{
				"maker": MkNumber(0.001),
				"taker": MkNumber(0.001),
			}),
			"funding": MkMap(&VarMap{"withdraw": MkMap(&VarMap{
				"BCH":  MkNumber(0.0),
				"BTC":  MkNumber(0.001),
				"LTC":  MkNumber(0.001),
				"ETH":  MkNumber(0.001),
				"ZEC":  MkNumber(0.0001),
				"DASH": MkNumber(0.0001),
			})}),
		}),
		"limits": MkMap(&VarMap{"amount": MkMap(&VarMap{
			"min": MkNumber(0.001),
			"max": MkUndefined(),
		})}),
		"precision": MkMap(&VarMap{
			"amount": MkInteger(8),
			"price":  MkInteger(8),
		}),
		"options":          MkMap(&VarMap{"createMarketBuyOrderRequiresPrice": MkBool(true)}),
		"commonCurrencies": MkMap(&VarMap{"ACM": MkString("Actinium")}),
	}))
}

func (this *Coinex) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetMarketInfo"), params)
	_ = response
	markets := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = markets
	result := MkArray(&VarArray{})
	_ = result
	keys := GoGetKeys(markets)
	_ = keys
	for i := MkInteger(0); IsTrue(OpLw(i, keys.Length)); OpIncr(&i) {
		key := *(keys).At(i)
		_ = key
		market := *(markets).At(key)
		_ = market
		id := this.SafeString(market, MkString("name"))
		_ = id
		tradingName := this.SafeString(market, MkString("trading_name"))
		_ = tradingName
		baseId := OpCopy(tradingName)
		_ = baseId
		quoteId := this.SafeString(market, MkString("pricing_name"))
		_ = quoteId
		base := this.SafeCurrencyCode(baseId)
		_ = base
		quote := this.SafeCurrencyCode(quoteId)
		_ = quote
		symbol := OpAdd(base, OpAdd(MkString("/"), quote))
		_ = symbol
		if IsTrue(OpEq2(tradingName, id)) {
			symbol = OpCopy(id)
		}
		precision := MkMap(&VarMap{
			"amount": this.SafeInteger(market, MkString("trading_decimal")),
			"price":  this.SafeInteger(market, MkString("pricing_decimal")),
		})
		_ = precision
		active := OpCopy(MkUndefined())
		_ = active
		result.Push(MkMap(&VarMap{
			"id":        id,
			"symbol":    symbol,
			"base":      base,
			"quote":     quote,
			"baseId":    baseId,
			"quoteId":   quoteId,
			"active":    active,
			"taker":     this.SafeNumber(market, MkString("taker_fee_rate")),
			"maker":     this.SafeNumber(market, MkString("maker_fee_rate")),
			"info":      market,
			"precision": precision,
			"limits": MkMap(&VarMap{
				"amount": MkMap(&VarMap{
					"min": this.SafeNumber(market, MkString("min_amount")),
					"max": MkUndefined(),
				}),
				"price": MkMap(&VarMap{
					"min": Math.Pow(MkInteger(10), OpNeg(*(precision).At(MkString("price")))),
					"max": MkUndefined(),
				}),
			}),
		}))
	}
	return result
}

func (this *Coinex) ParseTicker(goArgs ...*Variant) *Variant {
	ticker := GoGetArg(goArgs, 0, MkUndefined())
	_ = ticker
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.SafeInteger(ticker, MkString("date"))
	_ = timestamp
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(OpNotEq2(market, MkUndefined())) {
		symbol = *(market).At(MkString("symbol"))
	}
	ticker = this.SafeValue(ticker, MkString("ticker"), MkMap(&VarMap{}))
	last := this.SafeNumber(ticker, MkString("last"))
	_ = last
	return MkMap(&VarMap{
		"symbol":        symbol,
		"timestamp":     timestamp,
		"datetime":      this.Iso8601(timestamp),
		"high":          this.SafeNumber(ticker, MkString("high")),
		"low":           this.SafeNumber(ticker, MkString("low")),
		"bid":           this.SafeNumber(ticker, MkString("buy")),
		"bidVolume":     MkUndefined(),
		"ask":           this.SafeNumber(ticker, MkString("sell")),
		"askVolume":     MkUndefined(),
		"vwap":          MkUndefined(),
		"open":          MkUndefined(),
		"close":         last,
		"last":          last,
		"previousClose": MkUndefined(),
		"change":        MkUndefined(),
		"percentage":    MkUndefined(),
		"average":       MkUndefined(),
		"baseVolume":    this.SafeNumber2(ticker, MkString("vol"), MkString("volume")),
		"quoteVolume":   MkUndefined(),
		"info":          ticker,
	})
}

func (this *Coinex) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"market": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("publicGetMarketTicker"), this.Extend(request, params))
	_ = response
	return this.ParseTicker(*(response).At(MkString("data")), market)
}

func (this *Coinex) FetchTickers(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("publicGetMarketTickerAll"), params)
	_ = response
	data := this.SafeValue(response, MkString("data"))
	_ = data
	timestamp := this.SafeInteger(data, MkString("date"))
	_ = timestamp
	tickers := this.SafeValue(data, MkString("ticker"))
	_ = tickers
	marketIds := GoGetKeys(tickers)
	_ = marketIds
	result := MkMap(&VarMap{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, marketIds.Length)); OpIncr(&i) {
		marketId := *(marketIds).At(i)
		_ = marketId
		market := this.SafeMarket(marketId)
		_ = market
		symbol := *(market).At(MkString("symbol"))
		_ = symbol
		ticker := this.ParseTicker(MkMap(&VarMap{
			"date":   timestamp,
			"ticker": *(tickers).At(marketId),
		}), market)
		_ = ticker
		*(ticker).At(MkString("symbol")) = OpCopy(symbol)
		*(result).At(symbol) = OpCopy(ticker)
	}
	return this.FilterByArray(result, MkString("symbol"), symbols)
}

func (this *Coinex) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkInteger(20))
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	if IsTrue(OpEq2(limit, MkUndefined())) {
		limit = MkInteger(20)
	}
	request := MkMap(&VarMap{
		"market": this.MarketId(symbol),
		"merge":  MkString("0.0000000001"),
		"limit":  limit.ToString(),
	})
	_ = request
	response := this.Call(MkString("publicGetMarketDepth"), this.Extend(request, params))
	_ = response
	return this.ParseOrderBook(*(response).At(MkString("data")), symbol)
}

func (this *Coinex) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.SafeTimestamp(trade, MkString("create_time"))
	_ = timestamp
	if IsTrue(OpEq2(timestamp, MkUndefined())) {
		timestamp = this.SafeInteger(trade, MkString("date_ms"))
	}
	tradeId := this.SafeString(trade, MkString("id"))
	_ = tradeId
	orderId := this.SafeString(trade, MkString("order_id"))
	_ = orderId
	priceString := this.SafeString(trade, MkString("price"))
	_ = priceString
	amountString := this.SafeString(trade, MkString("amount"))
	_ = amountString
	price := this.ParseNumber(priceString)
	_ = price
	amount := this.ParseNumber(amountString)
	_ = amount
	marketId := this.SafeString(trade, MkString("market"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market)
	_ = symbol
	cost := this.SafeNumber(trade, MkString("deal_money"))
	_ = cost
	if IsTrue(OpEq2(cost, MkUndefined())) {
		cost = this.ParseNumber(Precise.StringMul(priceString, amountString))
	}
	fee := OpCopy(MkUndefined())
	_ = fee
	feeCost := this.SafeNumber(trade, MkString("fee"))
	_ = feeCost
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		feeCurrencyId := this.SafeString(trade, MkString("fee_asset"))
		_ = feeCurrencyId
		feeCurrencyCode := this.SafeCurrencyCode(feeCurrencyId)
		_ = feeCurrencyCode
		fee = MkMap(&VarMap{
			"cost":     feeCost,
			"currency": feeCurrencyCode,
		})
	}
	takerOrMaker := this.SafeString(trade, MkString("role"))
	_ = takerOrMaker
	side := this.SafeString(trade, MkString("type"))
	_ = side
	return MkMap(&VarMap{
		"info":         trade,
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       symbol,
		"id":           tradeId,
		"order":        orderId,
		"type":         MkUndefined(),
		"side":         side,
		"takerOrMaker": takerOrMaker,
		"price":        price,
		"amount":       amount,
		"cost":         cost,
		"fee":          fee,
	})
}

func (this *Coinex) FetchTrades(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"market": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("publicGetMarketDeals"), this.Extend(request, params))
	_ = response
	return this.ParseTrades(*(response).At(MkString("data")), market, since, limit)
}

func (this *Coinex) ParseOHLCV(goArgs ...*Variant) *Variant {
	ohlcv := GoGetArg(goArgs, 0, MkUndefined())
	_ = ohlcv
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	return MkArray(&VarArray{
		this.SafeTimestamp(ohlcv, MkInteger(0)),
		this.SafeNumber(ohlcv, MkInteger(1)),
		this.SafeNumber(ohlcv, MkInteger(3)),
		this.SafeNumber(ohlcv, MkInteger(4)),
		this.SafeNumber(ohlcv, MkInteger(2)),
		this.SafeNumber(ohlcv, MkInteger(5)),
	})
}

func (this *Coinex) FetchOHLCV(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	timeframe := GoGetArg(goArgs, 1, MkString("5m"))
	_ = timeframe
	since := GoGetArg(goArgs, 2, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 3, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 4, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{
		"market": *(market).At(MkString("id")),
		"type":   *(*this.At(MkString("timeframes"))).At(timeframe),
	})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("publicGetMarketKline"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseOHLCVs(data, market, timeframe, since, limit)
}

func (this *Coinex) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privateGetBalanceInfo"), params)
	_ = response
	result := MkMap(&VarMap{"info": response})
	_ = result
	balances := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = balances
	currencyIds := GoGetKeys(balances)
	_ = currencyIds
	for i := MkInteger(0); IsTrue(OpLw(i, currencyIds.Length)); OpIncr(&i) {
		currencyId := *(currencyIds).At(i)
		_ = currencyId
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		balance := this.SafeValue(balances, currencyId, MkMap(&VarMap{}))
		_ = balance
		account := this.Account()
		_ = account
		*(account).At(MkString("free")) = this.SafeString(balance, MkString("available"))
		*(account).At(MkString("used")) = this.SafeString(balance, MkString("frozen"))
		*(result).At(code) = OpCopy(account)
	}
	return this.ParseBalance(result)
}

func (this *Coinex) ParseOrderStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"not_deal":  MkString("open"),
		"part_deal": MkString("open"),
		"done":      MkString("closed"),
		"cancel":    MkString("canceled"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Coinex) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.SafeTimestamp(order, MkString("create_time"))
	_ = timestamp
	price := this.SafeNumber(order, MkString("price"))
	_ = price
	cost := this.SafeNumber(order, MkString("deal_money"))
	_ = cost
	amount := this.SafeNumber(order, MkString("amount"))
	_ = amount
	filled := this.SafeNumber(order, MkString("deal_amount"))
	_ = filled
	average := this.SafeNumber(order, MkString("avg_price"))
	_ = average
	symbol := OpCopy(MkUndefined())
	_ = symbol
	marketId := this.SafeString(order, MkString("market"))
	_ = marketId
	market = this.SafeMarket(marketId, market)
	feeCurrencyId := this.SafeString(order, MkString("fee_asset"))
	_ = feeCurrencyId
	feeCurrency := this.SafeCurrencyCode(feeCurrencyId)
	_ = feeCurrency
	if IsTrue(OpNotEq2(market, MkUndefined())) {
		symbol = *(market).At(MkString("symbol"))
		if IsTrue(OpEq2(feeCurrency, MkUndefined())) {
			feeCurrency = *(market).At(MkString("quote"))
		}
	}
	remaining := this.SafeNumber(order, MkString("left"))
	_ = remaining
	status := this.ParseOrderStatus(this.SafeString(order, MkString("status")))
	_ = status
	type_ := this.SafeString(order, MkString("order_type"))
	_ = type_
	side := this.SafeString(order, MkString("type"))
	_ = side
	return this.SafeOrder(MkMap(&VarMap{
		"id":                 this.SafeString(order, MkString("id")),
		"clientOrderId":      MkUndefined(),
		"datetime":           this.Iso8601(timestamp),
		"timestamp":          timestamp,
		"lastTradeTimestamp": MkUndefined(),
		"status":             status,
		"symbol":             symbol,
		"type":               type_,
		"timeInForce":        MkUndefined(),
		"postOnly":           MkUndefined(),
		"side":               side,
		"price":              price,
		"stopPrice":          MkUndefined(),
		"cost":               cost,
		"average":            average,
		"amount":             amount,
		"filled":             filled,
		"remaining":          remaining,
		"trades":             MkUndefined(),
		"fee": MkMap(&VarMap{
			"currency": feeCurrency,
			"cost":     this.SafeNumber(order, MkString("deal_fee")),
		}),
		"info": order,
	}))
}

func (this *Coinex) CreateOrder(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	type_ := GoGetArg(goArgs, 1, MkUndefined())
	_ = type_
	side := GoGetArg(goArgs, 2, MkUndefined())
	_ = side
	amount := GoGetArg(goArgs, 3, MkUndefined())
	_ = amount
	price := GoGetArg(goArgs, 4, MkUndefined())
	_ = price
	params := GoGetArg(goArgs, 5, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	method := OpAdd(MkString("privatePostOrder"), this.Capitalize(type_))
	_ = method
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{
		"market": *(market).At(MkString("id")),
		"type":   side,
	})
	_ = request
	if IsTrue(OpAnd(OpEq2(type_, MkString("market")), OpEq2(side, MkString("buy")))) {
		if IsTrue(*(*this.At(MkString("options"))).At(MkString("createMarketBuyOrderRequiresPrice"))) {
			if IsTrue(OpEq2(price, MkUndefined())) {
				panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")), MkString(" createOrder() requires the price argument with market buy orders to calculate total order cost (amount to spend), where cost = amount * price. Supply a price argument to createOrder() call if you want the cost to be calculated for you from price and amount, or, alternatively, add .options['createMarketBuyOrderRequiresPrice'] = false to supply the cost in the amount argument (the exchange-specific behaviour)"))))
			} else {
				*(request).At(MkString("amount")) = this.CostToPrecision(symbol, OpMulti(amount, price))
			}
		} else {
			*(request).At(MkString("amount")) = this.CostToPrecision(symbol, amount)
		}
	} else {
		*(request).At(MkString("amount")) = this.AmountToPrecision(symbol, amount)
	}
	if IsTrue(OpOr(OpEq2(type_, MkString("limit")), OpEq2(type_, MkString("ioc")))) {
		*(request).At(MkString("price")) = this.PriceToPrecision(symbol, price)
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"))
	_ = data
	return this.ParseOrder(data, market)
}

func (this *Coinex) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{
		"id":     id,
		"market": *(market).At(MkString("id")),
	})
	_ = request
	response := this.Call(MkString("privateDeleteOrderPending"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"))
	_ = data
	return this.ParseOrder(data, market)
}

func (this *Coinex) FetchOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchOrder() requires a symbol argument"))))
	}
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{
		"id":     id,
		"market": *(market).At(MkString("id")),
	})
	_ = request
	response := this.Call(MkString("privateGetOrder"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"))
	_ = data
	return this.ParseOrder(data, market)
}

func (this *Coinex) FetchOrdersByStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 2, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 3, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 4, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	if IsTrue(OpEq2(limit, MkUndefined())) {
		limit = MkInteger(100)
	}
	request := MkMap(&VarMap{
		"page":  MkInteger(1),
		"limit": limit,
	})
	_ = request
	market := OpCopy(MkUndefined())
	_ = market
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
		*(request).At(MkString("market")) = *(market).At(MkString("id"))
	}
	method := OpAdd(MkString("privateGetOrder"), this.Capitalize(status))
	_ = method
	response := this.Call(method, this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"))
	_ = data
	orders := this.SafeValue(data, MkString("data"), MkArray(&VarArray{}))
	_ = orders
	return this.ParseOrders(orders, market, since, limit)
}

func (this *Coinex) FetchOpenOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	return this.FetchOrdersByStatus(MkString("pending"), symbol, since, limit, params)
}

func (this *Coinex) FetchClosedOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	return this.FetchOrdersByStatus(MkString("finished"), symbol, since, limit, params)
}

func (this *Coinex) FetchMyTrades(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	if IsTrue(OpEq2(limit, MkUndefined())) {
		limit = MkInteger(100)
	}
	request := MkMap(&VarMap{
		"page":  MkInteger(1),
		"limit": limit,
	})
	_ = request
	market := OpCopy(MkUndefined())
	_ = market
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
		*(request).At(MkString("market")) = *(market).At(MkString("id"))
	}
	response := this.Call(MkString("privateGetOrderUserDeals"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"))
	_ = data
	trades := this.SafeValue(data, MkString("data"), MkArray(&VarArray{}))
	_ = trades
	return this.ParseTrades(trades, market, since, limit)
}

func (this *Coinex) Withdraw(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	amount := GoGetArg(goArgs, 1, MkUndefined())
	_ = amount
	address := GoGetArg(goArgs, 2, MkUndefined())
	_ = address
	tag := GoGetArg(goArgs, 3, MkUndefined())
	_ = tag
	params := GoGetArg(goArgs, 4, MkMap(&VarMap{}))
	_ = params
	this.CheckAddress(address)
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	if IsTrue(tag) {
		address = OpAdd(address, OpAdd(MkString(":"), tag))
	}
	request := MkMap(&VarMap{
		"coin_type":       *(currency).At(MkString("id")),
		"coin_address":    address,
		"actual_amount":   parseFloat(amount),
		"transfer_method": MkString("onchain"),
	})
	_ = request
	response := this.Call(MkString("privatePostBalanceCoinWithdraw"), this.Extend(request, params))
	_ = response
	transaction := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = transaction
	return this.ParseTransaction(transaction, currency)
}

func (this *Coinex) ParseTransactionStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"audit":      MkString("pending"),
		"pass":       MkString("pending"),
		"processing": MkString("pending"),
		"confirming": MkString("pending"),
		"not_pass":   MkString("failed"),
		"cancel":     MkString("canceled"),
		"finish":     MkString("ok"),
		"fail":       MkString("failed"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Coinex) ParseTransaction(goArgs ...*Variant) *Variant {
	transaction := GoGetArg(goArgs, 0, MkUndefined())
	_ = transaction
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	id := this.SafeString2(transaction, MkString("coin_withdraw_id"), MkString("coin_deposit_id"))
	_ = id
	address := this.SafeString(transaction, MkString("coin_address"))
	_ = address
	tag := this.SafeString(transaction, MkString("remark"))
	_ = tag
	if IsTrue(OpNotEq2(tag, MkUndefined())) {
		if IsTrue(OpLw(tag.Length, MkInteger(1))) {
			tag = OpCopy(MkUndefined())
		}
	}
	txid := this.SafeValue(transaction, MkString("tx_id"))
	_ = txid
	if IsTrue(OpNotEq2(txid, MkUndefined())) {
		if IsTrue(OpLw(txid.Length, MkInteger(1))) {
			txid = OpCopy(MkUndefined())
		}
	}
	currencyId := this.SafeString(transaction, MkString("coin_type"))
	_ = currencyId
	code := this.SafeCurrencyCode(currencyId, currency)
	_ = code
	timestamp := this.SafeTimestamp(transaction, MkString("create_time"))
	_ = timestamp
	type_ := OpTrinary(OpHasMember(MkString("coin_withdraw_id"), transaction), MkString("withdraw"), MkString("deposit"))
	_ = type_
	status := this.ParseTransactionStatus(this.SafeString(transaction, MkString("status")))
	_ = status
	amount := this.SafeNumber(transaction, MkString("amount"))
	_ = amount
	feeCost := this.SafeNumber(transaction, MkString("tx_fee"))
	_ = feeCost
	if IsTrue(OpEq2(type_, MkString("deposit"))) {
		feeCost = MkInteger(0)
	}
	fee := MkMap(&VarMap{
		"cost":     feeCost,
		"currency": code,
	})
	_ = fee
	if IsTrue(OpNotEq2(amount, MkUndefined())) {
		amount = OpSub(amount, feeCost)
	}
	return MkMap(&VarMap{
		"info":      transaction,
		"id":        id,
		"txid":      txid,
		"timestamp": timestamp,
		"datetime":  this.Iso8601(timestamp),
		"address":   address,
		"tag":       tag,
		"type":      type_,
		"amount":    amount,
		"currency":  code,
		"status":    status,
		"updated":   MkUndefined(),
		"fee":       fee,
	})
}

func (this *Coinex) FetchWithdrawals(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(code, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchWithdrawals() requires a currency code argument"))))
	}
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{"coin_type": *(currency).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("Limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("privateGetBalanceCoinWithdraw"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"))
	_ = data
	if IsTrue(OpNot(GoIsArray(data))) {
		data = this.SafeValue(data, MkString("data"), MkArray(&VarArray{}))
	}
	return this.ParseTransactions(data, currency, since, limit)
}

func (this *Coinex) FetchDeposits(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(code, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchDeposits() requires a currency code argument"))))
	}
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{"coin_type": *(currency).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("Limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("privateGetBalanceCoinDeposit"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"))
	_ = data
	if IsTrue(OpNot(GoIsArray(data))) {
		data = this.SafeValue(data, MkString("data"), MkArray(&VarArray{}))
	}
	return this.ParseTransactions(data, currency, since, limit)
}

func (this *Coinex) Nonce(goArgs ...*Variant) *Variant {
	return this.Milliseconds()
}

func (this *Coinex) Sign(goArgs ...*Variant) *Variant {
	path := GoGetArg(goArgs, 0, MkUndefined())
	_ = path
	api := GoGetArg(goArgs, 1, MkString("public"))
	_ = api
	method := GoGetArg(goArgs, 2, MkString("GET"))
	_ = method
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	headers := GoGetArg(goArgs, 4, MkUndefined())
	_ = headers
	body := GoGetArg(goArgs, 5, MkUndefined())
	_ = body
	path = this.ImplodeParams(path, params)
	url := OpAdd(*(*(*this.At(MkString("urls"))).At(MkString("api"))).At(api), OpAdd(MkString("/"), OpAdd(*this.At(MkString("version")), OpAdd(MkString("/"), path))))
	_ = url
	query := this.Omit(params, this.ExtractParams(path))
	_ = query
	if IsTrue(OpOr(OpEq2(api, MkString("public")), OpEq2(api, MkString("perpetualPublic")))) {
		if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
			OpAddAssign(&url, OpAdd(MkString("?"), this.Urlencode(query)))
		}
	} else {
		this.CheckRequiredCredentials()
		nonce := this.Nonce()
		_ = nonce
		query = this.Extend(MkMap(&VarMap{
			"access_id": *this.At(MkString("apiKey")),
			"tonce":     nonce.ToString(),
		}), query)
		query = this.Keysort(query)
		urlencoded := this.Rawencode(query)
		_ = urlencoded
		signature := this.Hash(this.Encode(OpAdd(urlencoded, OpAdd(MkString("&secret_key="), *this.At(MkString("secret"))))))
		_ = signature
		headers = MkMap(&VarMap{
			"Authorization": signature.ToUpperCase(),
			"Content-Type":  MkString("application/json"),
		})
		if IsTrue(OpOr(OpEq2(method, MkString("GET")), OpEq2(method, MkString("DELETE")))) {
			OpAddAssign(&url, OpAdd(MkString("?"), urlencoded))
		} else {
			body = this.Json(query)
		}
	}
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Coinex) HandleErrors(goArgs ...*Variant) *Variant {
	httpCode := GoGetArg(goArgs, 0, MkUndefined())
	_ = httpCode
	reason := GoGetArg(goArgs, 1, MkUndefined())
	_ = reason
	url := GoGetArg(goArgs, 2, MkUndefined())
	_ = url
	method := GoGetArg(goArgs, 3, MkUndefined())
	_ = method
	headers := GoGetArg(goArgs, 4, MkUndefined())
	_ = headers
	body := GoGetArg(goArgs, 5, MkUndefined())
	_ = body
	response := GoGetArg(goArgs, 6, MkUndefined())
	_ = response
	requestHeaders := GoGetArg(goArgs, 7, MkUndefined())
	_ = requestHeaders
	requestBody := GoGetArg(goArgs, 8, MkUndefined())
	_ = requestBody
	if IsTrue(OpEq2(response, MkUndefined())) {
		return MkUndefined()
	}
	code := this.SafeString(response, MkString("code"))
	_ = code
	data := this.SafeValue(response, MkString("data"))
	_ = data
	message := this.SafeString(response, MkString("message"))
	_ = message
	if IsTrue(OpOr(OpNotEq2(code, MkString("0")), OpOr(OpEq2(data, MkUndefined()), OpAnd(OpNotEq2(message, MkString("Success")), OpAnd(OpNotEq2(message, MkString("Ok")), OpNot(data)))))) {
		responseCodes := MkMap(&VarMap{
			"24":  AuthenticationError,
			"25":  AuthenticationError,
			"107": InsufficientFunds,
			"600": OrderNotFound,
			"601": InvalidOrder,
			"602": InvalidOrder,
			"606": InvalidOrder,
		})
		_ = responseCodes
		ErrorClass := this.SafeValue(responseCodes, code, ExchangeError)
		_ = ErrorClass
		panic(NewErrorClass(*(response).At(MkString("message"))))
	}
	return MkUndefined()
}
