package ccxt

import ()

type Bitbank struct {
	*ExchangeBase
}

var _ Exchange = (*Bitbank)(nil)

func init() {
	exchange := &Bitbank{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Bitbank) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("bitbank"),
		"name": MkString("bitbank"),
		"countries": MkArray(&VarArray{
			MkString("JP"),
		}),
		"version": MkString("v1"),
		"has": MkMap(&VarMap{
			"cancelOrder":         MkBool(true),
			"createOrder":         MkBool(true),
			"fetchBalance":        MkBool(true),
			"fetchDepositAddress": MkBool(true),
			"fetchMyTrades":       MkBool(true),
			"fetchOHLCV":          MkBool(true),
			"fetchOpenOrders":     MkBool(true),
			"fetchOrder":          MkBool(true),
			"fetchOrderBook":      MkBool(true),
			"fetchTicker":         MkBool(true),
			"fetchTrades":         MkBool(true),
			"withdraw":            MkBool(true),
		}),
		"timeframes": MkMap(&VarMap{
			"1m":  MkString("1min"),
			"5m":  MkString("5min"),
			"15m": MkString("15min"),
			"30m": MkString("30min"),
			"1h":  MkString("1hour"),
			"4h":  MkString("4hour"),
			"8h":  MkString("8hour"),
			"12h": MkString("12hour"),
			"1d":  MkString("1day"),
			"1w":  MkString("1week"),
		}),
		"hostname": MkString("bitbank.cc"),
		"urls": MkMap(&VarMap{
			"logo": MkString("https://user-images.githubusercontent.com/1294454/37808081-b87f2d9c-2e59-11e8-894d-c1900b7584fe.jpg"),
			"api": MkMap(&VarMap{
				"public":  MkString("https://public.{hostname}"),
				"private": MkString("https://api.{hostname}"),
				"markets": MkString("https://api.{hostname}"),
			}),
			"www":  MkString("https://bitbank.cc/"),
			"doc":  MkString("https://docs.bitbank.cc/"),
			"fees": MkString("https://bitbank.cc/docs/fees/"),
		}),
		"api": MkMap(&VarMap{
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("{pair}/ticker"),
				MkString("{pair}/depth"),
				MkString("{pair}/transactions"),
				MkString("{pair}/transactions/{yyyymmdd}"),
				MkString("{pair}/candlestick/{candletype}/{yyyymmdd}"),
			})}),
			"private": MkMap(&VarMap{
				"get": MkArray(&VarArray{
					MkString("user/assets"),
					MkString("user/spot/order"),
					MkString("user/spot/active_orders"),
					MkString("user/spot/trade_history"),
					MkString("user/withdrawal_account"),
				}),
				"post": MkArray(&VarArray{
					MkString("user/spot/order"),
					MkString("user/spot/cancel_order"),
					MkString("user/spot/cancel_orders"),
					MkString("user/spot/orders_info"),
					MkString("user/request_withdrawal"),
				}),
			}),
			"markets": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("spot/pairs"),
			})}),
		}),
		"exceptions": MkMap(&VarMap{
			"20001": AuthenticationError,
			"20002": AuthenticationError,
			"20003": AuthenticationError,
			"20005": AuthenticationError,
			"20004": InvalidNonce,
			"40020": InvalidOrder,
			"40021": InvalidOrder,
			"40025": ExchangeError,
			"40013": OrderNotFound,
			"40014": OrderNotFound,
			"50008": PermissionDenied,
			"50009": OrderNotFound,
			"50010": OrderNotFound,
			"60001": InsufficientFunds,
			"60005": InvalidOrder,
		}),
	}))
}

func (this *Bitbank) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("marketsGetSpotPairs"), params)
	_ = response
	data := this.SafeValue(response, MkString("data"))
	_ = data
	pairs := this.SafeValue(data, MkString("pairs"), MkArray(&VarArray{}))
	_ = pairs
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, pairs.Length)); OpIncr(&i) {
		entry := *(pairs).At(i)
		_ = entry
		id := this.SafeString(entry, MkString("name"))
		_ = id
		baseId := this.SafeString(entry, MkString("base_asset"))
		_ = baseId
		quoteId := this.SafeString(entry, MkString("quote_asset"))
		_ = quoteId
		base := this.SafeCurrencyCode(baseId)
		_ = base
		quote := this.SafeCurrencyCode(quoteId)
		_ = quote
		symbol := OpAdd(base, OpAdd(MkString("/"), quote))
		_ = symbol
		maker := this.SafeNumber(entry, MkString("maker_fee_rate_quote"))
		_ = maker
		taker := this.SafeNumber(entry, MkString("taker_fee_rate_quote"))
		_ = taker
		pricePrecisionString := this.SafeString(entry, MkString("price_digits"))
		_ = pricePrecisionString
		priceLimit := this.ParsePrecision(pricePrecisionString)
		_ = priceLimit
		precision := MkMap(&VarMap{
			"price":  parseInt(pricePrecisionString),
			"amount": this.SafeInteger(entry, MkString("amount_digits")),
		})
		_ = precision
		active := this.SafeValue(entry, MkString("is_enabled"))
		_ = active
		minAmountString := this.SafeString(entry, MkString("unit_amount"))
		_ = minAmountString
		minCost := Precise.StringMul(minAmountString, priceLimit)
		_ = minCost
		limits := MkMap(&VarMap{
			"amount": MkMap(&VarMap{
				"min": this.SafeNumber(entry, MkString("unit_amount")),
				"max": this.SafeNumber(entry, MkString("limit_max_amount")),
			}),
			"price": MkMap(&VarMap{
				"min": this.ParseNumber(priceLimit),
				"max": MkUndefined(),
			}),
			"cost": MkMap(&VarMap{
				"min": this.ParseNumber(minCost),
				"max": MkUndefined(),
			}),
		})
		_ = limits
		result.Push(MkMap(&VarMap{
			"info":      entry,
			"id":        id,
			"symbol":    symbol,
			"baseId":    baseId,
			"quoteId":   quoteId,
			"base":      base,
			"quote":     quote,
			"precision": precision,
			"limits":    limits,
			"active":    active,
			"maker":     maker,
			"taker":     taker,
		}))
	}
	return result
}

func (this *Bitbank) ParseTicker(goArgs ...*Variant) *Variant {
	ticker := GoGetArg(goArgs, 0, MkUndefined())
	_ = ticker
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(OpNotEq2(market, MkUndefined())) {
		symbol = *(market).At(MkString("symbol"))
	}
	timestamp := this.SafeInteger(ticker, MkString("timestamp"))
	_ = timestamp
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
		"baseVolume":    this.SafeNumber(ticker, MkString("vol")),
		"quoteVolume":   MkUndefined(),
		"info":          ticker,
	})
}

func (this *Bitbank) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"pair": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("publicGetPairTicker"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	return this.ParseTicker(data, market)
}

func (this *Bitbank) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"pair": this.MarketId(symbol)})
	_ = request
	response := this.Call(MkString("publicGetPairDepth"), this.Extend(request, params))
	_ = response
	orderbook := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = orderbook
	timestamp := this.SafeInteger(orderbook, MkString("timestamp"))
	_ = timestamp
	return this.ParseOrderBook(orderbook, symbol, timestamp)
}

func (this *Bitbank) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.SafeInteger(trade, MkString("executed_at"))
	_ = timestamp
	symbol := OpCopy(MkUndefined())
	_ = symbol
	feeCurrency := OpCopy(MkUndefined())
	_ = feeCurrency
	if IsTrue(OpNotEq2(market, MkUndefined())) {
		symbol = *(market).At(MkString("symbol"))
		feeCurrency = *(market).At(MkString("quote"))
	}
	priceString := this.SafeString(trade, MkString("price"))
	_ = priceString
	amountString := this.SafeString(trade, MkString("amount"))
	_ = amountString
	price := this.ParseNumber(priceString)
	_ = price
	amount := this.ParseNumber(amountString)
	_ = amount
	cost := this.ParseNumber(Precise.StringMul(priceString, amountString))
	_ = cost
	id := this.SafeString2(trade, MkString("transaction_id"), MkString("trade_id"))
	_ = id
	takerOrMaker := this.SafeString(trade, MkString("maker_taker"))
	_ = takerOrMaker
	fee := OpCopy(MkUndefined())
	_ = fee
	feeCost := this.SafeNumber(trade, MkString("fee_amount_quote"))
	_ = feeCost
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		fee = MkMap(&VarMap{
			"currency": feeCurrency,
			"cost":     feeCost,
		})
	}
	orderId := this.SafeString(trade, MkString("order_id"))
	_ = orderId
	type_ := this.SafeString(trade, MkString("type"))
	_ = type_
	side := this.SafeString(trade, MkString("side"))
	_ = side
	return MkMap(&VarMap{
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       symbol,
		"id":           id,
		"order":        orderId,
		"type":         type_,
		"side":         side,
		"takerOrMaker": takerOrMaker,
		"price":        price,
		"amount":       amount,
		"cost":         cost,
		"fee":          fee,
		"info":         trade,
	})
}

func (this *Bitbank) FetchTrades(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"pair": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("publicGetPairTransactions"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	trades := this.SafeValue(data, MkString("transactions"), MkArray(&VarArray{}))
	_ = trades
	return this.ParseTrades(trades, market, since, limit)
}

func (this *Bitbank) ParseOHLCV(goArgs ...*Variant) *Variant {
	ohlcv := GoGetArg(goArgs, 0, MkUndefined())
	_ = ohlcv
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	return MkArray(&VarArray{
		this.SafeInteger(ohlcv, MkInteger(5)),
		this.SafeNumber(ohlcv, MkInteger(0)),
		this.SafeNumber(ohlcv, MkInteger(1)),
		this.SafeNumber(ohlcv, MkInteger(2)),
		this.SafeNumber(ohlcv, MkInteger(3)),
		this.SafeNumber(ohlcv, MkInteger(4)),
	})
}

func (this *Bitbank) FetchOHLCV(goArgs ...*Variant) *Variant {
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
	date := this.Milliseconds()
	_ = date
	date = this.Ymd(date)
	date = date.Split(MkString("-"))
	request := MkMap(&VarMap{
		"pair":       *(market).At(MkString("id")),
		"candletype": *(*this.At(MkString("timeframes"))).At(timeframe),
		"yyyymmdd":   date.Join(MkString("")),
	})
	_ = request
	response := this.Call(MkString("publicGetPairCandlestickCandletypeYyyymmdd"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	candlestick := this.SafeValue(data, MkString("candlestick"), MkArray(&VarArray{}))
	_ = candlestick
	first := this.SafeValue(candlestick, MkInteger(0), MkMap(&VarMap{}))
	_ = first
	ohlcv := this.SafeValue(first, MkString("ohlcv"), MkArray(&VarArray{}))
	_ = ohlcv
	return this.ParseOHLCVs(ohlcv, market, timeframe, since, limit)
}

func (this *Bitbank) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privateGetUserAssets"), params)
	_ = response
	result := MkMap(&VarMap{
		"info":      response,
		"timestamp": MkUndefined(),
		"datetime":  MkUndefined(),
	})
	_ = result
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	assets := this.SafeValue(data, MkString("assets"), MkArray(&VarArray{}))
	_ = assets
	for i := MkInteger(0); IsTrue(OpLw(i, assets.Length)); OpIncr(&i) {
		balance := *(assets).At(i)
		_ = balance
		currencyId := this.SafeString(balance, MkString("asset"))
		_ = currencyId
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		account := this.Account()
		_ = account
		*(account).At(MkString("free")) = this.SafeString(balance, MkString("free_amount"))
		*(account).At(MkString("used")) = this.SafeString(balance, MkString("locked_amount"))
		*(account).At(MkString("total")) = this.SafeString(balance, MkString("onhand_amount"))
		*(result).At(code) = OpCopy(account)
	}
	return this.ParseBalance(result)
}

func (this *Bitbank) ParseOrderStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"UNFILLED":                  MkString("open"),
		"PARTIALLY_FILLED":          MkString("open"),
		"FULLY_FILLED":              MkString("closed"),
		"CANCELED_UNFILLED":         MkString("canceled"),
		"CANCELED_PARTIALLY_FILLED": MkString("canceled"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Bitbank) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	id := this.SafeString(order, MkString("order_id"))
	_ = id
	marketId := this.SafeString(order, MkString("pair"))
	_ = marketId
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(OpAnd(marketId, OpAnd(OpNot(market), OpHasMember(marketId, *this.At(MkString("markets_by_id")))))) {
		market = *(*this.At(MkString("markets_by_id"))).At(marketId)
	}
	if IsTrue(OpNotEq2(market, MkUndefined())) {
		symbol = *(market).At(MkString("symbol"))
	}
	timestamp := this.SafeInteger(order, MkString("ordered_at"))
	_ = timestamp
	price := this.SafeNumber(order, MkString("price"))
	_ = price
	amount := this.SafeNumber(order, MkString("start_amount"))
	_ = amount
	filled := this.SafeNumber(order, MkString("executed_amount"))
	_ = filled
	remaining := this.SafeNumber(order, MkString("remaining_amount"))
	_ = remaining
	average := this.SafeNumber(order, MkString("average_price"))
	_ = average
	status := this.ParseOrderStatus(this.SafeString(order, MkString("status")))
	_ = status
	type_ := this.SafeStringLower(order, MkString("type"))
	_ = type_
	side := this.SafeStringLower(order, MkString("side"))
	_ = side
	return this.SafeOrder(MkMap(&VarMap{
		"id":                 id,
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
		"cost":               MkUndefined(),
		"average":            average,
		"amount":             amount,
		"filled":             filled,
		"remaining":          remaining,
		"trades":             MkUndefined(),
		"fee":                MkUndefined(),
		"info":               order,
	}))
}

func (this *Bitbank) CreateOrder(goArgs ...*Variant) *Variant {
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
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{
		"pair":   *(market).At(MkString("id")),
		"amount": this.AmountToPrecision(symbol, amount),
		"side":   side,
		"type":   type_,
	})
	_ = request
	if IsTrue(OpEq2(type_, MkString("limit"))) {
		*(request).At(MkString("price")) = this.PriceToPrecision(symbol, price)
	}
	response := this.Call(MkString("privatePostUserSpotOrder"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"))
	_ = data
	return this.ParseOrder(data, market)
}

func (this *Bitbank) CancelOrder(goArgs ...*Variant) *Variant {
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
		"order_id": id,
		"pair":     *(market).At(MkString("id")),
	})
	_ = request
	response := this.Call(MkString("privatePostUserSpotCancelOrder"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"))
	_ = data
	return data
}

func (this *Bitbank) FetchOrder(goArgs ...*Variant) *Variant {
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
		"order_id": id,
		"pair":     *(market).At(MkString("id")),
	})
	_ = request
	response := this.Call(MkString("privateGetUserSpotOrder"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"))
	_ = data
	return this.ParseOrder(data, market)
}

func (this *Bitbank) FetchOpenOrders(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"pair": *(market).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("count")) = OpCopy(limit)
	}
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("since")) = parseInt(OpDiv(since, MkInteger(1000)))
	}
	response := this.Call(MkString("privateGetUserSpotActiveOrders"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	orders := this.SafeValue(data, MkString("orders"), MkArray(&VarArray{}))
	_ = orders
	return this.ParseOrders(orders, market, since, limit)
}

func (this *Bitbank) FetchMyTrades(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := OpCopy(MkUndefined())
	_ = market
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
	}
	request := MkMap(&VarMap{})
	_ = request
	if IsTrue(OpNotEq2(market, MkUndefined())) {
		*(request).At(MkString("pair")) = *(market).At(MkString("id"))
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("count")) = OpCopy(limit)
	}
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("since")) = parseInt(OpDiv(since, MkInteger(1000)))
	}
	response := this.Call(MkString("privateGetUserSpotTradeHistory"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	trades := this.SafeValue(data, MkString("trades"), MkArray(&VarArray{}))
	_ = trades
	return this.ParseTrades(trades, market, since, limit)
}

func (this *Bitbank) FetchDepositAddress(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{"asset": *(currency).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("privateGetUserWithdrawalAccount"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	accounts := this.SafeValue(data, MkString("accounts"), MkArray(&VarArray{}))
	_ = accounts
	firstAccount := this.SafeValue(accounts, MkInteger(0), MkMap(&VarMap{}))
	_ = firstAccount
	address := this.SafeString(firstAccount, MkString("address"))
	_ = address
	return MkMap(&VarMap{
		"currency": currency,
		"address":  address,
		"tag":      MkUndefined(),
		"info":     response,
	})
}

func (this *Bitbank) Withdraw(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpNot(OpHasMember(MkString("uuid"), params))) {
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), MkString(" uuid is required for withdrawal"))))
	}
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{
		"asset":  *(currency).At(MkString("id")),
		"amount": amount,
	})
	_ = request
	response := this.Call(MkString("privatePostUserRequestWithdrawal"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	txid := this.SafeString(data, MkString("txid"))
	_ = txid
	return MkMap(&VarMap{
		"info": response,
		"id":   txid,
	})
}

func (this *Bitbank) Nonce(goArgs ...*Variant) *Variant {
	return this.Milliseconds()
}

func (this *Bitbank) Sign(goArgs ...*Variant) *Variant {
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
	query := this.Omit(params, this.ExtractParams(path))
	_ = query
	url := OpAdd(this.ImplodeHostname(*(*(*this.At(MkString("urls"))).At(MkString("api"))).At(api)), MkString("/"))
	_ = url
	if IsTrue(OpOr(OpEq2(api, MkString("public")), OpEq2(api, MkString("markets")))) {
		OpAddAssign(&url, this.ImplodeParams(path, params))
		if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
			OpAddAssign(&url, OpAdd(MkString("?"), this.Urlencode(query)))
		}
	} else {
		this.CheckRequiredCredentials()
		nonce := (this.Nonce()).Call(MkString("toString"))
		_ = nonce
		auth := OpCopy(nonce)
		_ = auth
		OpAddAssign(&url, OpAdd(*this.At(MkString("version")), OpAdd(MkString("/"), this.ImplodeParams(path, params))))
		if IsTrue(OpEq2(method, MkString("POST"))) {
			body = this.Json(query)
			OpAddAssign(&auth, body)
		} else {
			OpAddAssign(&auth, OpAdd(MkString("/"), OpAdd(*this.At(MkString("version")), OpAdd(MkString("/"), path))))
			if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
				query = this.Urlencode(query)
				OpAddAssign(&url, OpAdd(MkString("?"), query))
				OpAddAssign(&auth, OpAdd(MkString("?"), query))
			}
		}
		headers = MkMap(&VarMap{
			"Content-Type":     MkString("application/json"),
			"ACCESS-KEY":       *this.At(MkString("apiKey")),
			"ACCESS-NONCE":     nonce,
			"ACCESS-SIGNATURE": this.Hmac(this.Encode(auth), this.Encode(*this.At(MkString("secret")))),
		})
	}
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Bitbank) HandleErrors(goArgs ...*Variant) *Variant {
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
	success := this.SafeInteger(response, MkString("success"))
	_ = success
	data := this.SafeValue(response, MkString("data"))
	_ = data
	if IsTrue(OpOr(OpNot(success), OpNot(data))) {
		errorMessages := MkMap(&VarMap{
			"10000": MkString("URL does not exist"),
			"10001": MkString("A system error occurred. Please contact support"),
			"10002": MkString("Invalid JSON format. Please check the contents of transmission"),
			"10003": MkString("A system error occurred. Please contact support"),
			"10005": MkString("A timeout error occurred. Please wait for a while and try again"),
			"20001": MkString("API authentication failed"),
			"20002": MkString("Illegal API key"),
			"20003": MkString("API key does not exist"),
			"20004": MkString("API Nonce does not exist"),
			"20005": MkString("API signature does not exist"),
			"20011": MkString("Two-step verification failed"),
			"20014": MkString("SMS authentication failed"),
			"30001": MkString("Please specify the order quantity"),
			"30006": MkString("Please specify the order ID"),
			"30007": MkString("Please specify the order ID array"),
			"30009": MkString("Please specify the stock"),
			"30012": MkString("Please specify the order price"),
			"30013": MkString("Trade Please specify either"),
			"30015": MkString("Please specify the order type"),
			"30016": MkString("Please specify asset name"),
			"30019": MkString("Please specify uuid"),
			"30039": MkString("Please specify the amount to be withdrawn"),
			"40001": MkString("The order quantity is invalid"),
			"40006": MkString("Count value is invalid"),
			"40007": MkString("End time is invalid"),
			"40008": MkString("end_id Value is invalid"),
			"40009": MkString("The from_id value is invalid"),
			"40013": MkString("The order ID is invalid"),
			"40014": MkString("The order ID array is invalid"),
			"40015": MkString("Too many specified orders"),
			"40017": MkString("Incorrect issue name"),
			"40020": MkString("The order price is invalid"),
			"40021": MkString("The trading classification is invalid"),
			"40022": MkString("Start date is invalid"),
			"40024": MkString("The order type is invalid"),
			"40025": MkString("Incorrect asset name"),
			"40028": MkString("uuid is invalid"),
			"40048": MkString("The amount of withdrawal is illegal"),
			"50003": MkString("Currently, this account is in a state where you can not perform the operation you specified. Please contact support"),
			"50004": MkString("Currently, this account is temporarily registered. Please try again after registering your account"),
			"50005": MkString("Currently, this account is locked. Please contact support"),
			"50006": MkString("Currently, this account is locked. Please contact support"),
			"50008": MkString("User identification has not been completed"),
			"50009": MkString("Your order does not exist"),
			"50010": MkString("Can not cancel specified order"),
			"50011": MkString("API not found"),
			"60001": MkString("The number of possessions is insufficient"),
			"60002": MkString("It exceeds the quantity upper limit of the tender buying order"),
			"60003": MkString("The specified quantity exceeds the limit"),
			"60004": MkString("The specified quantity is below the threshold"),
			"60005": MkString("The specified price is above the limit"),
			"60006": MkString("The specified price is below the lower limit"),
			"70001": MkString("A system error occurred. Please contact support"),
			"70002": MkString("A system error occurred. Please contact support"),
			"70003": MkString("A system error occurred. Please contact support"),
			"70004": MkString("We are unable to accept orders as the transaction is currently suspended"),
			"70005": MkString("Order can not be accepted because purchase order is currently suspended"),
			"70006": MkString("We can not accept orders because we are currently unsubscribed "),
			"70009": MkString("We are currently temporarily restricting orders to be carried out. Please use the limit order."),
			"70010": MkString("We are temporarily raising the minimum order quantity as the system load is now rising."),
		})
		_ = errorMessages
		errorClasses := OpCopy(*this.At(MkString("exceptions")))
		_ = errorClasses
		code := this.SafeString(data, MkString("code"))
		_ = code
		message := this.SafeString(errorMessages, code, MkString("Error"))
		_ = message
		ErrorClass := this.SafeValue(errorClasses, code)
		_ = ErrorClass
		if IsTrue(OpNotEq2(ErrorClass, MkUndefined())) {
			panic(NewErrorClass(message))
		} else {
			panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), this.Json(response)))))
		}
	}
	return MkUndefined()
}
