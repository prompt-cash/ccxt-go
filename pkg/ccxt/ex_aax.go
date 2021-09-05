package ccxt

import ()

type Aax struct {
	*ExchangeBase
}

var _ Exchange = (*Aax)(nil)

func init() {
	exchange := &Aax{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Aax) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("aax"),
		"name": MkString("AAX"),
		"countries": MkArray(&VarArray{
			MkString("MT"),
		}),
		"enableRateLimit": MkBool(true),
		"rateLimit":       MkInteger(500),
		"version":         MkString("v2"),
		"hostname":        MkString("aaxpro.com"),
		"certified":       MkBool(true),
		"pro":             MkBool(true),
		"has": MkMap(&VarMap{
			"cancelAllOrders":     MkBool(true),
			"cancelOrder":         MkBool(true),
			"createOrder":         MkBool(true),
			"editOrder":           MkBool(true),
			"fetchBalance":        MkBool(true),
			"fetchCanceledOrders": MkBool(true),
			"fetchClosedOrders":   MkBool(true),
			"fetchCurrencies":     MkBool(true),
			"fetchDepositAddress": MkBool(true),
			"fetchMarkets":        MkBool(true),
			"fetchMyTrades":       MkBool(true),
			"fetchOHLCV":          MkBool(true),
			"fetchOpenOrders":     MkBool(true),
			"fetchOrder":          MkBool(true),
			"fetchOrderBook":      MkBool(true),
			"fetchOrders":         MkBool(true),
			"fetchStatus":         MkBool(true),
			"fetchTicker":         MkString("emulated"),
			"fetchTickers":        MkBool(true),
			"fetchTrades":         MkBool(true),
		}),
		"timeframes": MkMap(&VarMap{
			"1m":  MkString("1m"),
			"5m":  MkString("5m"),
			"15m": MkString("15m"),
			"30m": MkString("30m"),
			"1h":  MkString("1h"),
			"2h":  MkString("2h"),
			"4h":  MkString("4h"),
			"12h": MkString("12h"),
			"1d":  MkString("1d"),
			"3d":  MkString("3d"),
			"1w":  MkString("1w"),
		}),
		"urls": MkMap(&VarMap{
			"logo": MkString("https://user-images.githubusercontent.com/1294454/104140087-a27f2580-53c0-11eb-87c1-5d9e81208fe9.jpg"),
			"test": MkMap(&VarMap{
				"v1":      MkString("https://api.testnet.{hostname}/marketdata/v1"),
				"public":  MkString("https://api.testnet.{hostname}"),
				"private": MkString("https://api.testnet.{hostname}"),
			}),
			"api": MkMap(&VarMap{
				"v1":      MkString("https://api.{hostname}/marketdata/v1"),
				"public":  MkString("https://api.{hostname}"),
				"private": MkString("https://api.{hostname}"),
			}),
			"www":      MkString("https://www.aaxpro.com"),
			"doc":      MkString("https://www.aaxpro.com/apidoc/index.html"),
			"fees":     MkString("https://www.aaxpro.com/en-US/fees/"),
			"referral": MkString("https://www.aaxpro.com/invite/sign-up?inviteCode=JXGm5Fy7R2MB"),
		}),
		"api": MkMap(&VarMap{
			"v1": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("getHistMarketData"),
			})}),
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("currencies"),
				MkString("announcement/maintenance"),
				MkString("instruments"),
				MkString("market/orderbook"),
				MkString("futures/position/openInterest"),
				MkString("market/tickers"),
				MkString("market/candles"),
				MkString("market/history/candles"),
				MkString("market/trades"),
				MkString("market/markPrice"),
				MkString("futures/funding/predictedFunding/{symbol}"),
				MkString("futures/funding/prevFundingRate/{symbol}"),
				MkString("market/candles/index"),
			})}),
			"private": MkMap(&VarMap{
				"get": MkArray(&VarArray{
					MkString("user/info"),
					MkString("account/balances"),
					MkString("account/deposit/address"),
					MkString("spot/trades"),
					MkString("spot/openOrders"),
					MkString("spot/orders"),
					MkString("futures/position"),
					MkString("futures/position/closed"),
					MkString("futures/trades"),
					MkString("futures/openOrders"),
					MkString("futures/orders"),
					MkString("futures/funding/predictedFundingFee/{symbol}"),
				}),
				"post": MkArray(&VarArray{
					MkString("account/transfer"),
					MkString("spot/orders"),
					MkString("spot/orders/cancelAllOnTimeout"),
					MkString("futures/orders"),
					MkString("futures/orders/cancelAllOnTimeout"),
					MkString("futures/position/sltp"),
					MkString("futures/position/close"),
					MkString("futures/position/leverage"),
					MkString("futures/position/margin"),
				}),
				"put": MkArray(&VarArray{
					MkString("spot/orders"),
					MkString("futures/orders"),
				}),
				"delete": MkArray(&VarArray{
					MkString("spot/orders/cancel/{orderID}"),
					MkString("spot/orders/cancel/all"),
					MkString("futures/orders/cancel/{orderID}"),
					MkString("futures/orders/cancel/all"),
				}),
			}),
		}),
		"fees": MkMap(&VarMap{
			"trading": MkMap(&VarMap{
				"tierBased":  MkBool(false),
				"percentage": MkBool(true),
				"maker":      this.ParseNumber(MkString("0.0006")),
				"taker":      this.ParseNumber(MkString("0.001")),
			}),
			"funding": MkMap(&VarMap{
				"tierBased":  MkBool(false),
				"percentage": MkBool(true),
				"withdraw":   MkMap(&VarMap{}),
			}),
		}),
		"commonCurrencies": MkMap(&VarMap{"XBT": MkString("XBT")}),
		"exceptions": MkMap(&VarMap{
			"exact": MkMap(&VarMap{
				"2002":  InsufficientFunds,
				"2003":  OrderNotFound,
				"10003": BadRequest,
				"10006": AuthenticationError,
				"10007": AuthenticationError,
				"11007": AuthenticationError,
				"20001": InsufficientFunds,
				"20009": InvalidOrder,
				"30000": OrderNotFound,
				"30001": InvalidOrder,
				"30004": InvalidOrder,
				"30005": InvalidOrder,
				"30006": InvalidOrder,
				"30007": InvalidOrder,
				"30008": InvalidOrder,
				"30009": InvalidOrder,
				"30010": InvalidOrder,
				"30011": CancelPending,
				"30012": BadRequest,
				"30013": BadSymbol,
				"30014": OrderNotFound,
				"30015": InvalidOrder,
				"30016": ExchangeError,
				"30017": InvalidOrder,
				"30018": InvalidOrder,
				"30019": InvalidOrder,
				"30020": InvalidOrder,
				"30021": InvalidOrder,
				"30022": InvalidOrder,
				"30023": InvalidOrder,
				"30024": InvalidOrder,
				"30025": InvalidOrder,
				"30026": InvalidOrder,
				"30027": InvalidOrder,
				"30028": BadSymbol,
				"30029": InvalidOrder,
				"30030": InvalidOrder,
				"30031": InvalidOrder,
				"30032": InvalidOrder,
				"30033": InvalidOrder,
				"30034": RateLimitExceeded,
				"30035": RateLimitExceeded,
				"30036": ExchangeNotAvailable,
				"30037": InvalidOrder,
				"30038": ExchangeError,
				"30039": InsufficientFunds,
				"30040": InvalidOrder,
				"30041": InvalidOrder,
				"30042": InvalidOrder,
				"30043": InvalidOrder,
				"30044": BadRequest,
				"30045": InvalidOrder,
				"30046": InvalidOrder,
				"30047": InvalidOrder,
				"30048": InvalidOrder,
				"30049": InvalidOrder,
				"30050": InvalidOrder,
				"40004": BadRequest,
				"40009": RateLimitExceeded,
				"40102": AuthenticationError,
				"40103": AuthenticationError,
				"40303": PermissionDenied,
				"41001": BadRequest,
				"41002": BadRequest,
				"42001": ExchangeNotAvailable,
				"50001": ExchangeError,
				"50002": ExchangeError,
			}),
			"broad": MkMap(&VarMap{}),
		}),
		"precisionMode": TICK_SIZE,
		"options": MkMap(&VarMap{
			"defaultType": MkString("spot"),
			"types": MkMap(&VarMap{
				"spot":   MkString("SPTP"),
				"future": MkString("FUTP"),
				"otc":    MkString("F2CP"),
				"saving": MkString("VLTP"),
			}),
			"accounts": MkMap(&VarMap{
				"SPTP": MkString("spot"),
				"FUTP": MkString("future"),
				"F2CP": MkString("otc"),
				"VLTP": MkString("saving"),
			}),
		}),
	}))
}

func (this *Aax) FetchStatus(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetAnnouncementMaintenance"), params)
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	timestamp := this.Milliseconds()
	_ = timestamp
	startTime := this.Parse8601(this.SafeString(data, MkString("startTime")))
	_ = startTime
	endTime := this.Parse8601(this.SafeString(data, MkString("endTime")))
	_ = endTime
	update := MkMap(&VarMap{"updated": this.SafeInteger(response, MkString("ts"), timestamp)})
	_ = update
	if IsTrue(OpNotEq2(endTime, MkUndefined())) {
		startTimeIsOk := OpTrinary(OpEq2(startTime, MkUndefined()), MkBool(true), OpLw(timestamp, startTime))
		_ = startTimeIsOk
		isOk := OpOr(OpGt(timestamp, endTime), startTimeIsOk)
		_ = isOk
		*(update).At(MkString("eta")) = OpCopy(endTime)
		*(update).At(MkString("status")) = OpTrinary(isOk, MkString("ok"), MkString("maintenance"))
	}
	*this.At(MkString("status")) = this.Extend(*this.At(MkString("status")), update)
	return *this.At(MkString("status"))
}

func (this *Aax) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetInstruments"), params)
	_ = response
	data := this.SafeValue(response, MkString("data"))
	_ = data
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, data.Length)); OpIncr(&i) {
		market := *(data).At(i)
		_ = market
		id := this.SafeString(market, MkString("symbol"))
		_ = id
		baseId := this.SafeString(market, MkString("base"))
		_ = baseId
		quoteId := this.SafeString(market, MkString("quote"))
		_ = quoteId
		base := this.SafeCurrencyCode(baseId)
		_ = base
		quote := this.SafeCurrencyCode(quoteId)
		_ = quote
		status := this.SafeString(market, MkString("status"))
		_ = status
		active := OpEq2(status, MkString("enable"))
		_ = active
		taker := this.SafeNumber(market, MkString("takerFee"))
		_ = taker
		maker := this.SafeNumber(market, MkString("makerFee"))
		_ = maker
		type_ := this.SafeString(market, MkString("type"))
		_ = type_
		inverse := OpCopy(MkUndefined())
		_ = inverse
		linear := OpCopy(MkUndefined())
		_ = linear
		quanto := OpCopy(MkUndefined())
		_ = quanto
		spot := OpEq2(type_, MkString("spot"))
		_ = spot
		futures := OpEq2(type_, MkString("futures"))
		_ = futures
		settleType := this.SafeStringLower(market, MkString("settleType"))
		_ = settleType
		if IsTrue(OpNotEq2(settleType, MkUndefined())) {
			inverse = OpEq2(settleType, MkString("inverse"))
			linear = OpEq2(settleType, MkString("vanilla"))
			quanto = OpEq2(settleType, MkString("quanto"))
		}
		symbol := OpCopy(id)
		_ = symbol
		if IsTrue(OpEq2(type_, MkString("spot"))) {
			symbol = OpAdd(base, OpAdd(MkString("/"), quote))
		}
		precision := MkMap(&VarMap{
			"amount": this.SafeNumber(market, MkString("lotSize")),
			"price":  this.SafeNumber(market, MkString("tickSize")),
		})
		_ = precision
		result.Push(MkMap(&VarMap{
			"id":         id,
			"symbol":     symbol,
			"base":       base,
			"quote":      quote,
			"baseId":     baseId,
			"quoteId":    quoteId,
			"type":       type_,
			"spot":       spot,
			"futures":    futures,
			"inverse":    inverse,
			"linear":     linear,
			"quanto":     quanto,
			"precision":  precision,
			"info":       market,
			"active":     active,
			"taker":      taker,
			"maker":      maker,
			"percentage": MkBool(false),
			"tierBased":  MkBool(true),
			"limits": MkMap(&VarMap{
				"amount": MkMap(&VarMap{
					"min": this.SafeString(market, MkString("minQuantity")),
					"max": this.SafeString(market, MkString("maxQuantity")),
				}),
				"price": MkMap(&VarMap{
					"min": this.SafeString(market, MkString("minPrice")),
					"max": this.SafeString(market, MkString("maxPrice")),
				}),
				"cost": MkMap(&VarMap{
					"min": MkUndefined(),
					"max": MkUndefined(),
				}),
			}),
		}))
	}
	return result
}

func (this *Aax) FetchCurrencies(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetCurrencies"), params)
	_ = response
	result := MkMap(&VarMap{})
	_ = result
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	for i := MkInteger(0); IsTrue(OpLw(i, data.Length)); OpIncr(&i) {
		currency := *(data).At(i)
		_ = currency
		id := this.SafeString(currency, MkString("chain"))
		_ = id
		name := this.SafeString(currency, MkString("displayName"))
		_ = name
		code := this.SafeCurrencyCode(id)
		_ = code
		precision := this.SafeNumber(currency, MkString("withdrawPrecision"))
		_ = precision
		enableWithdraw := this.SafeValue(currency, MkString("enableWithdraw"))
		_ = enableWithdraw
		enableDeposit := this.SafeValue(currency, MkString("enableDeposit"))
		_ = enableDeposit
		fee := this.SafeNumber(currency, MkString("withdrawFee"))
		_ = fee
		visible := this.SafeValue(currency, MkString("visible"))
		_ = visible
		active := OpAnd(enableWithdraw, OpAnd(enableDeposit, visible))
		_ = active
		network := this.SafeString(currency, MkString("network"))
		_ = network
		*(result).At(code) = MkMap(&VarMap{
			"id":        id,
			"name":      name,
			"code":      code,
			"precision": precision,
			"info":      currency,
			"active":    active,
			"fee":       fee,
			"network":   network,
			"limits": MkMap(&VarMap{
				"deposit": MkMap(&VarMap{
					"min": this.SafeNumber(currency, MkString("depositMin")),
					"max": MkUndefined(),
				}),
				"withdraw": MkMap(&VarMap{
					"min": this.SafeNumber(currency, MkString("withdrawMin")),
					"max": MkUndefined(),
				}),
			}),
		})
	}
	return result
}

func (this *Aax) ParseTicker(goArgs ...*Variant) *Variant {
	ticker := GoGetArg(goArgs, 0, MkUndefined())
	_ = ticker
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.SafeInteger(ticker, MkString("t"))
	_ = timestamp
	marketId := this.SafeString(ticker, MkString("s"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market)
	_ = symbol
	last := this.SafeNumber(ticker, MkString("c"))
	_ = last
	open := this.SafeNumber(ticker, MkString("o"))
	_ = open
	change := OpCopy(MkUndefined())
	_ = change
	percentage := OpCopy(MkUndefined())
	_ = percentage
	average := OpCopy(MkUndefined())
	_ = average
	if IsTrue(OpAnd(OpNotEq2(last, MkUndefined()), OpNotEq2(open, MkUndefined()))) {
		change = OpSub(last, open)
		if IsTrue(OpGt(open, MkInteger(0))) {
			percentage = OpDiv(change, OpMulti(open, MkInteger(100)))
		}
		average = OpDiv(this.Sum(last, open), MkInteger(2))
	}
	quoteVolume := this.SafeNumber(ticker, MkString("v"))
	_ = quoteVolume
	return MkMap(&VarMap{
		"symbol":        symbol,
		"timestamp":     timestamp,
		"datetime":      this.Iso8601(timestamp),
		"high":          this.SafeNumber(ticker, MkString("h")),
		"low":           this.SafeNumber(ticker, MkString("l")),
		"bid":           MkUndefined(),
		"bidVolume":     MkUndefined(),
		"ask":           MkUndefined(),
		"askVolume":     MkUndefined(),
		"vwap":          MkUndefined(),
		"open":          open,
		"close":         last,
		"last":          last,
		"previousClose": MkUndefined(),
		"change":        change,
		"percentage":    percentage,
		"average":       average,
		"baseVolume":    MkUndefined(),
		"quoteVolume":   quoteVolume,
		"info":          ticker,
	})
}

func (this *Aax) FetchTickers(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("publicGetMarketTickers"), params)
	_ = response
	tickers := this.SafeValue(response, MkString("tickers"), MkArray(&VarArray{}))
	_ = tickers
	result := MkArray(&VarArray{})
	_ = result
	timestamp := this.SafeInteger(response, MkString("t"))
	_ = timestamp
	for i := MkInteger(0); IsTrue(OpLw(i, tickers.Length)); OpIncr(&i) {
		ticker := this.ParseTicker(this.Extend(*(tickers).At(i), MkMap(&VarMap{"t": timestamp})))
		_ = ticker
		result.Push(ticker)
	}
	return this.FilterByArray(result, MkString("symbol"), symbols)
}

func (this *Aax) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	if IsTrue(OpEq2(limit, MkUndefined())) {
		limit = MkInteger(20)
	} else {
		if IsTrue(OpAnd(OpNotEq2(limit, MkInteger(20)), OpNotEq2(limit, MkInteger(50)))) {
			panic(NewBadRequest(OpAdd(*this.At(MkString("id")), MkString(" fetchOrderBook() limit argument must be undefined, 20 or 50"))))
		}
	}
	request := MkMap(&VarMap{
		"symbol": *(market).At(MkString("id")),
		"level":  limit,
	})
	_ = request
	response := this.Call(MkString("publicGetMarketOrderbook"), this.Extend(request, params))
	_ = response
	timestamp := this.SafeInteger(response, MkString("t"))
	_ = timestamp
	return this.ParseOrderBook(response, symbol, timestamp)
}

func (this *Aax) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.SafeInteger(trade, MkString("t"))
	_ = timestamp
	if IsTrue(OpEq2(timestamp, MkUndefined())) {
		timestamp = this.Parse8601(this.SafeString(trade, MkString("createTime")))
	}
	id := this.SafeString2(trade, MkString("tid"), MkString("tradeID"))
	_ = id
	symbol := OpCopy(MkUndefined())
	_ = symbol
	marketId := this.SafeString(trade, MkString("symbol"))
	_ = marketId
	market = this.SafeMarket(marketId, market)
	if IsTrue(OpNotEq2(market, MkUndefined())) {
		symbol = *(market).At(MkString("symbol"))
	}
	priceString := this.SafeString2(trade, MkString("p"), MkString("filledPrice"))
	_ = priceString
	amountString := this.SafeString2(trade, MkString("q"), MkString("filledQty"))
	_ = amountString
	orderId := this.SafeString(trade, MkString("orderID"))
	_ = orderId
	isTaker := this.SafeValue(trade, MkString("taker"))
	_ = isTaker
	takerOrMaker := OpCopy(MkUndefined())
	_ = takerOrMaker
	if IsTrue(OpNotEq2(isTaker, MkUndefined())) {
		takerOrMaker = OpTrinary(isTaker, MkString("taker"), MkString("maker"))
	}
	side := this.SafeString(trade, MkString("side"))
	_ = side
	if IsTrue(OpEq2(side, MkString("1"))) {
		side = MkString("buy")
	} else {
		if IsTrue(OpEq2(side, MkString("2"))) {
			side = MkString("sell")
		}
	}
	if IsTrue(OpEq2(side, MkUndefined())) {
		side = OpTrinary(OpEq2(*(priceString).At(MkInteger(0)), MkString("-")), MkString("sell"), MkString("buy"))
	}
	priceString = Precise.StringAbs(priceString)
	price := this.ParseNumber(priceString)
	_ = price
	amount := this.ParseNumber(amountString)
	_ = amount
	cost := this.ParseNumber(Precise.StringMul(priceString, amountString))
	_ = cost
	orderType := this.ParseOrderType(this.SafeString(trade, MkString("orderType")))
	_ = orderType
	fee := OpCopy(MkUndefined())
	_ = fee
	feeCost := this.SafeNumber(trade, MkString("commission"))
	_ = feeCost
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		feeCurrency := OpCopy(MkUndefined())
		_ = feeCurrency
		if IsTrue(OpNotEq2(market, MkUndefined())) {
			if IsTrue(OpEq2(side, MkString("buy"))) {
				feeCurrency = *(market).At(MkString("base"))
			} else {
				if IsTrue(OpEq2(side, MkString("sell"))) {
					feeCurrency = *(market).At(MkString("quote"))
				}
			}
		}
		fee = MkMap(&VarMap{
			"currency": feeCurrency,
			"cost":     feeCost,
		})
	}
	return MkMap(&VarMap{
		"info":         trade,
		"id":           id,
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       symbol,
		"type":         orderType,
		"side":         side,
		"order":        orderId,
		"takerOrMaker": takerOrMaker,
		"price":        price,
		"amount":       amount,
		"cost":         cost,
		"fee":          fee,
	})
}

func (this *Aax) FetchTrades(goArgs ...*Variant) *Variant {
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
	limit = OpTrinary(OpEq2(limit, MkUndefined()), MkInteger(2000), limit)
	limit = Math.Min(limit, MkInteger(2000))
	request := MkMap(&VarMap{
		"symbol": *(market).At(MkString("id")),
		"limit":  limit,
	})
	_ = request
	response := this.Call(MkString("publicGetMarketTrades"), request)
	_ = response
	trades := this.SafeValue(response, MkString("trades"), MkArray(&VarArray{}))
	_ = trades
	return this.ParseTrades(trades, market, since, limit)
}

func (this *Aax) ParseOHLCV(goArgs ...*Variant) *Variant {
	ohlcv := GoGetArg(goArgs, 0, MkUndefined())
	_ = ohlcv
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	return MkArray(&VarArray{
		this.SafeTimestamp(ohlcv, MkInteger(5)),
		this.SafeNumber(ohlcv, MkInteger(0)),
		this.SafeNumber(ohlcv, MkInteger(1)),
		this.SafeNumber(ohlcv, MkInteger(2)),
		this.SafeNumber(ohlcv, MkInteger(3)),
		this.SafeNumber(ohlcv, MkInteger(4)),
	})
}

func (this *Aax) FetchOHLCV(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	timeframe := GoGetArg(goArgs, 1, MkString("1h"))
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
		"symbol":    *(market).At(MkString("id")),
		"timeFrame": *(*this.At(MkString("timeframes"))).At(timeframe),
	})
	_ = request
	limit = OpTrinary(OpEq2(limit, MkUndefined()), MkInteger(500), limit)
	duration := this.ParseTimeframe(timeframe)
	_ = duration
	if IsTrue(OpEq2(since, MkUndefined())) {
		end := this.Seconds()
		_ = end
		*(request).At(MkString("start")) = OpSub(end, OpMulti(duration, limit))
		*(request).At(MkString("end")) = OpCopy(end)
	} else {
		start := parseInt(OpDiv(since, MkInteger(1000)))
		_ = start
		*(request).At(MkString("start")) = OpCopy(start)
		*(request).At(MkString("end")) = this.Sum(start, OpMulti(duration, limit))
	}
	response := this.Call(MkString("publicGetMarketHistoryCandles"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseOHLCVs(data, market, timeframe, since, limit)
}

func (this *Aax) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	defaultType := this.SafeString2(*this.At(MkString("options")), MkString("fetchBalance"), MkString("defaultType"), MkString("spot"))
	_ = defaultType
	type_ := this.SafeString(params, MkString("type"), defaultType)
	_ = type_
	types := this.SafeValue(*this.At(MkString("options")), MkString("types"), MkMap(&VarMap{}))
	_ = types
	purseType := this.SafeString(types, type_, type_)
	_ = purseType
	request := MkMap(&VarMap{"purseType": purseType})
	_ = request
	params = this.Omit(params, MkString("type"))
	response := this.Call(MkString("privateGetAccountBalances"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"))
	_ = data
	timestamp := this.SafeInteger(response, MkString("ts"))
	_ = timestamp
	result := MkMap(&VarMap{
		"info":      response,
		"timestamp": timestamp,
		"datetime":  this.Iso8601(timestamp),
	})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, data.Length)); OpIncr(&i) {
		balance := *(data).At(i)
		_ = balance
		balanceType := this.SafeString(balance, MkString("purseType"))
		_ = balanceType
		if IsTrue(OpEq2(balanceType, purseType)) {
			currencyId := this.SafeString(balance, MkString("currency"))
			_ = currencyId
			code := this.SafeCurrencyCode(currencyId)
			_ = code
			account := this.Account()
			_ = account
			*(account).At(MkString("free")) = this.SafeString(balance, MkString("available"))
			*(account).At(MkString("used")) = this.SafeString(balance, MkString("unavailable"))
			*(result).At(code) = OpCopy(account)
		}
	}
	return this.ParseBalance(result)
}

func (this *Aax) CreateOrder(goArgs ...*Variant) *Variant {
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
	orderType := type_.ToUpperCase()
	_ = orderType
	orderSide := side.ToUpperCase()
	_ = orderSide
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{
		"symbol":   *(market).At(MkString("id")),
		"orderQty": this.AmountToPrecision(symbol, amount),
		"side":     orderSide,
	})
	_ = request
	timeInForce := this.SafeString(params, MkString("timeInForce"))
	_ = timeInForce
	if IsTrue(OpNotEq2(timeInForce, MkUndefined())) {
		*(request).At(MkString("timeInForce")) = OpCopy(timeInForce)
		params = this.Omit(params, MkString("timeInForce"))
	}
	clientOrderId := this.SafeString2(params, MkString("clOrdID"), MkString("clientOrderId"))
	_ = clientOrderId
	if IsTrue(OpNotEq2(clientOrderId, MkUndefined())) {
		*(request).At(MkString("clOrdID")) = OpCopy(clientOrderId)
		params = this.Omit(params, MkArray(&VarArray{
			MkString("clOrdID"),
			MkString("clientOrderId"),
		}))
	}
	stopPrice := this.SafeNumber(params, MkString("stopPrice"))
	_ = stopPrice
	if IsTrue(OpEq2(stopPrice, MkUndefined())) {
		if IsTrue(OpOr(OpEq2(orderType, MkString("STOP-LIMIT")), OpEq2(orderType, MkString("STOP")))) {
			panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" createOrder() requires a stopPrice parameter for "), OpAdd(orderType, MkString(" orders"))))))
		}
	} else {
		if IsTrue(OpEq2(orderType, MkString("LIMIT"))) {
			orderType = MkString("STOP-LIMIT")
		} else {
			if IsTrue(OpEq2(orderType, MkString("MARKET"))) {
				orderType = MkString("STOP")
			}
		}
		*(request).At(MkString("stopPrice")) = this.PriceToPrecision(symbol, stopPrice)
		params = this.Omit(params, MkString("stopPrice"))
	}
	if IsTrue(OpOr(OpEq2(orderType, MkString("LIMIT")), OpEq2(orderType, MkString("STOP-LIMIT")))) {
		*(request).At(MkString("price")) = this.PriceToPrecision(symbol, price)
	}
	*(request).At(MkString("orderType")) = OpCopy(orderType)
	method := OpCopy(MkUndefined())
	_ = method
	if IsTrue(*(market).At(MkString("spot"))) {
		method = MkString("privatePostSpotOrders")
	} else {
		if IsTrue(*(market).At(MkString("futures"))) {
			method = MkString("privatePostFuturesOrders")
		}
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	return this.ParseOrder(data, market)
}

func (this *Aax) EditOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	type_ := GoGetArg(goArgs, 2, MkUndefined())
	_ = type_
	side := GoGetArg(goArgs, 3, MkUndefined())
	_ = side
	amount := GoGetArg(goArgs, 4, MkUndefined())
	_ = amount
	price := GoGetArg(goArgs, 5, MkUndefined())
	_ = price
	params := GoGetArg(goArgs, 6, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"orderID": id})
	_ = request
	stopPrice := this.SafeNumber(params, MkString("stopPrice"))
	_ = stopPrice
	if IsTrue(OpNotEq2(stopPrice, MkUndefined())) {
		*(request).At(MkString("stopPrice")) = this.PriceToPrecision(symbol, stopPrice)
		params = this.Omit(params, MkString("stopPrice"))
	}
	if IsTrue(OpNotEq2(price, MkUndefined())) {
		*(request).At(MkString("price")) = this.PriceToPrecision(symbol, price)
	}
	if IsTrue(OpNotEq2(amount, MkUndefined())) {
		*(request).At(MkString("orderQty")) = this.AmountToPrecision(symbol, amount)
	}
	method := OpCopy(MkUndefined())
	_ = method
	if IsTrue(*(market).At(MkString("spot"))) {
		method = MkString("privatePutSpotOrders")
	} else {
		if IsTrue(*(market).At(MkString("futures"))) {
			method = MkString("privatePutFuturesOrders")
		}
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	return this.ParseOrder(data, market)
}

func (this *Aax) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"orderID": id})
	_ = request
	method := OpCopy(MkUndefined())
	_ = method
	defaultType := this.SafeString2(*this.At(MkString("options")), MkString("cancelOrder"), MkString("defaultType"), MkString("spot"))
	_ = defaultType
	type_ := this.SafeString(params, MkString("type"), defaultType)
	_ = type_
	params = this.Omit(params, MkString("type"))
	market := OpCopy(MkUndefined())
	_ = market
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
		type_ = *(market).At(MkString("type"))
	}
	if IsTrue(OpEq2(type_, MkString("spot"))) {
		method = MkString("privateDeleteSpotOrdersCancelOrderID")
	} else {
		if IsTrue(OpEq2(type_, MkString("futures"))) {
			method = MkString("privateDeleteFuturesOrdersCancelOrderID")
		}
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	return this.ParseOrder(data, market)
}

func (this *Aax) CancelAllOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" cancelAllOrders() requires a symbol argument"))))
	}
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	method := OpCopy(MkUndefined())
	_ = method
	if IsTrue(*(market).At(MkString("spot"))) {
		method = MkString("privateDeleteSpotOrdersCancelAll")
	} else {
		if IsTrue(*(market).At(MkString("futures"))) {
			method = MkString("privateDeleteFuturesOrdersCancelAll")
		}
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	return response
}

func (this *Aax) FetchOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	defaultType := this.SafeString2(*this.At(MkString("options")), MkString("fetchOrder"), MkString("defaultType"), MkString("spot"))
	_ = defaultType
	*(params).At(MkString("type")) = this.SafeString(params, MkString("type"), defaultType)
	request := MkMap(&VarMap{})
	_ = request
	clientOrderId := this.SafeString2(params, MkString("clOrdID"), MkString("clientOrderId"))
	_ = clientOrderId
	if IsTrue(OpEq2(clientOrderId, MkUndefined())) {
		*(request).At(MkString("orderID")) = OpCopy(id)
	} else {
		*(request).At(MkString("clOrdID")) = OpCopy(clientOrderId)
		params = this.Omit(params, MkArray(&VarArray{
			MkString("clOrdID"),
			MkString("clientOrderId"),
		}))
	}
	orders := this.FetchOrders(symbol, MkUndefined(), MkUndefined(), this.Extend(request, params))
	_ = orders
	order := this.SafeValue(orders, MkInteger(0))
	_ = order
	if IsTrue(OpEq2(order, MkUndefined())) {
		if IsTrue(OpEq2(clientOrderId, MkUndefined())) {
			panic(NewOrderNotFound(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" fetchOrder() could not find order id "), id))))
		} else {
			panic(NewOrderNotFound(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" fetchOrder() could not find order clientOrderID "), clientOrderId))))
		}
	}
	return order
}

func (this *Aax) FetchOpenOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{})
	_ = request
	defaultType := this.SafeString2(*this.At(MkString("options")), MkString("fetchOpenOrders"), MkString("defaultType"), MkString("spot"))
	_ = defaultType
	type_ := this.SafeString(params, MkString("type"), defaultType)
	_ = type_
	params = this.Omit(params, MkString("type"))
	market := OpCopy(MkUndefined())
	_ = market
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
		*(request).At(MkString("symbol")) = *(market).At(MkString("id"))
		type_ = *(market).At(MkString("type"))
	}
	clientOrderId := this.SafeString2(params, MkString("clOrdID"), MkString("clientOrderId"))
	_ = clientOrderId
	if IsTrue(OpNotEq2(clientOrderId, MkUndefined())) {
		*(request).At(MkString("clOrdID")) = OpCopy(clientOrderId)
		params = this.Omit(params, MkArray(&VarArray{
			MkString("clOrdID"),
			MkString("clientOrderId"),
		}))
	}
	method := OpCopy(MkUndefined())
	_ = method
	if IsTrue(OpEq2(type_, MkString("spot"))) {
		method = MkString("privateGetSpotOpenOrders")
	} else {
		if IsTrue(OpEq2(type_, MkString("futures"))) {
			method = MkString("privateGetFuturesOpenOrders")
		}
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("pageSize")) = OpCopy(limit)
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	orders := this.SafeValue(data, MkString("list"), MkArray(&VarArray{}))
	_ = orders
	return this.ParseOrders(orders, market, since, limit)
}

func (this *Aax) FetchClosedOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	request := MkMap(&VarMap{"orderStatus": MkString("2")})
	_ = request
	return this.FetchOrders(symbol, since, limit, this.Extend(request, params))
}

func (this *Aax) FetchCanceledOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	request := MkMap(&VarMap{"orderStatus": MkString("3")})
	_ = request
	return this.FetchOrders(symbol, since, limit, this.Extend(request, params))
}

func (this *Aax) FetchOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{})
	_ = request
	method := OpCopy(MkUndefined())
	_ = method
	defaultType := this.SafeString2(*this.At(MkString("options")), MkString("fetchOrders"), MkString("defaultType"), MkString("spot"))
	_ = defaultType
	type_ := this.SafeString(params, MkString("type"), defaultType)
	_ = type_
	params = this.Omit(params, MkString("type"))
	market := OpCopy(MkUndefined())
	_ = market
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
		*(request).At(MkString("symbol")) = *(market).At(MkString("id"))
		type_ = *(market).At(MkString("type"))
	}
	if IsTrue(OpEq2(type_, MkString("spot"))) {
		method = MkString("privateGetSpotOrders")
	} else {
		if IsTrue(OpEq2(type_, MkString("futures"))) {
			method = MkString("privateGetFuturesOrders")
		}
	}
	clientOrderId := this.SafeString2(params, MkString("clOrdID"), MkString("clientOrderId"))
	_ = clientOrderId
	if IsTrue(OpNotEq2(clientOrderId, MkUndefined())) {
		*(request).At(MkString("clOrdID")) = OpCopy(clientOrderId)
		params = this.Omit(params, MkArray(&VarArray{
			MkString("clOrdID"),
			MkString("clientOrderId"),
		}))
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("pageSize")) = OpCopy(limit)
	}
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("startDate")) = this.Ymd(since)
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	orders := this.SafeValue(data, MkString("list"), MkArray(&VarArray{}))
	_ = orders
	return this.ParseOrders(orders, market, since, limit)
}

func (this *Aax) ParseOrderStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"0":  MkString("open"),
		"1":  MkString("open"),
		"2":  MkString("open"),
		"3":  MkString("closed"),
		"4":  MkString("canceled"),
		"5":  MkString("canceled"),
		"6":  MkString("rejected"),
		"10": MkString("expired"),
		"11": MkString("rejected"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Aax) ParseOrderType(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"1": MkString("market"),
		"2": MkString("limit"),
		"3": MkString("stop"),
		"4": MkString("stop-limit"),
		"7": MkString("stop-loss"),
		"8": MkString("take-profit"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Aax) ParseTimeInForce(goArgs ...*Variant) *Variant {
	timeInForce := GoGetArg(goArgs, 0, MkUndefined())
	_ = timeInForce
	timeInForces := MkMap(&VarMap{
		"1": MkString("GTC"),
		"3": MkString("IOC"),
		"4": MkString("FOK"),
	})
	_ = timeInForces
	return this.SafeString(timeInForces, timeInForce, timeInForce)
}

func (this *Aax) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.SafeValue(order, MkString("createTime"))
	_ = timestamp
	if IsTrue(OpEq2(OpType(timestamp), MkString("string"))) {
		timestamp = this.Parse8601(timestamp)
	}
	status := this.ParseOrderStatus(this.SafeString(order, MkString("orderStatus")))
	_ = status
	type_ := this.ParseOrderType(this.SafeString(order, MkString("orderType")))
	_ = type_
	side := this.SafeString(order, MkString("side"))
	_ = side
	if IsTrue(OpEq2(side, MkString("1"))) {
		side = MkString("buy")
	} else {
		if IsTrue(OpEq2(side, MkString("2"))) {
			side = MkString("sell")
		}
	}
	id := this.SafeString(order, MkString("orderID"))
	_ = id
	clientOrderId := this.SafeString(order, MkString("clOrdID"))
	_ = clientOrderId
	marketId := this.SafeString(order, MkString("symbol"))
	_ = marketId
	market = this.SafeMarket(marketId, market)
	symbol := *(market).At(MkString("symbol"))
	_ = symbol
	price := this.SafeNumber(order, MkString("price"))
	_ = price
	stopPrice := this.SafeNumber(order, MkString("stopPrice"))
	_ = stopPrice
	timeInForce := this.ParseTimeInForce(this.SafeString(order, MkString("timeInForce")))
	_ = timeInForce
	execInst := this.SafeString(order, MkString("execInst"))
	_ = execInst
	postOnly := OpEq2(execInst, MkString("Post-Only"))
	_ = postOnly
	average := this.SafeNumber(order, MkString("avgPrice"))
	_ = average
	amount := this.SafeNumber(order, MkString("orderQty"))
	_ = amount
	filled := this.SafeNumber(order, MkString("cumQty"))
	_ = filled
	remaining := this.SafeNumber(order, MkString("leavesQty"))
	_ = remaining
	if IsTrue(OpAnd(OpEq2(filled, MkInteger(0)), OpEq2(remaining, MkInteger(0)))) {
		remaining = OpCopy(MkUndefined())
	}
	lastTradeTimestamp := this.SafeValue(order, MkString("transactTime"))
	_ = lastTradeTimestamp
	if IsTrue(OpEq2(OpType(lastTradeTimestamp), MkString("string"))) {
		lastTradeTimestamp = this.Parse8601(lastTradeTimestamp)
	}
	fee := OpCopy(MkUndefined())
	_ = fee
	feeCost := this.SafeNumber(order, MkString("commission"))
	_ = feeCost
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		feeCurrency := OpCopy(MkUndefined())
		_ = feeCurrency
		if IsTrue(OpNotEq2(market, MkUndefined())) {
			if IsTrue(OpEq2(side, MkString("buy"))) {
				feeCurrency = *(market).At(MkString("base"))
			} else {
				if IsTrue(OpEq2(side, MkString("sell"))) {
					feeCurrency = *(market).At(MkString("quote"))
				}
			}
		}
		fee = MkMap(&VarMap{
			"currency": feeCurrency,
			"cost":     feeCost,
		})
	}
	return this.SafeOrder(MkMap(&VarMap{
		"id":                 id,
		"info":               order,
		"clientOrderId":      clientOrderId,
		"timestamp":          timestamp,
		"datetime":           this.Iso8601(timestamp),
		"lastTradeTimestamp": lastTradeTimestamp,
		"status":             status,
		"symbol":             symbol,
		"type":               type_,
		"timeInForce":        timeInForce,
		"postOnly":           postOnly,
		"side":               side,
		"price":              price,
		"stopPrice":          stopPrice,
		"average":            average,
		"amount":             amount,
		"filled":             filled,
		"remaining":          remaining,
		"cost":               MkUndefined(),
		"trades":             MkUndefined(),
		"fee":                fee,
	}))
}

func (this *Aax) FetchMyTrades(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{})
	_ = request
	method := OpCopy(MkUndefined())
	_ = method
	defaultType := this.SafeString2(*this.At(MkString("options")), MkString("fetchMyTrades"), MkString("defaultType"), MkString("spot"))
	_ = defaultType
	type_ := this.SafeString(params, MkString("type"), defaultType)
	_ = type_
	params = this.Omit(params, MkString("type"))
	market := OpCopy(MkUndefined())
	_ = market
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
		*(request).At(MkString("symbol")) = *(market).At(MkString("id"))
		type_ = *(market).At(MkString("type"))
	}
	if IsTrue(OpEq2(type_, MkString("spot"))) {
		method = MkString("privateGetSpotTrades")
	} else {
		if IsTrue(OpEq2(type_, MkString("futures"))) {
			method = MkString("privateGetFuturesTrades")
		}
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("pageSize")) = OpCopy(limit)
	}
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("startDate")) = this.Ymd(since)
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	trades := this.SafeValue(data, MkString("list"), MkArray(&VarArray{}))
	_ = trades
	return this.ParseTrades(trades, market, since, limit)
}

func (this *Aax) FetchDepositAddress(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{"currency": *(currency).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("privateGetAccountDepositAddress"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	return this.ParseDepositAddress(data, currency)
}

func (this *Aax) ParseDepositAddress(goArgs ...*Variant) *Variant {
	depositAddress := GoGetArg(goArgs, 0, MkUndefined())
	_ = depositAddress
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	address := this.SafeString(depositAddress, MkString("address"))
	_ = address
	tag := this.SafeString(depositAddress, MkString("tag"))
	_ = tag
	currencyId := this.SafeString(depositAddress, MkString("currency"))
	_ = currencyId
	code := this.SafeCurrencyCode(currencyId)
	_ = code
	return MkMap(&VarMap{
		"info":    depositAddress,
		"code":    code,
		"address": address,
		"tag":     tag,
	})
}

func (this *Aax) Nonce(goArgs ...*Variant) *Variant {
	return this.Milliseconds()
}

func (this *Aax) Sign(goArgs ...*Variant) *Variant {
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
	url := OpAdd(MkString("/"), this.ImplodeParams(path, params))
	_ = url
	query := this.Omit(params, this.ExtractParams(path))
	_ = query
	if IsTrue(OpEq2(api, MkString("v1"))) {
		if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
			OpAddAssign(&url, OpAdd(MkString("?"), this.Urlencode(query)))
		}
	} else {
		url = OpAdd(MkString("/"), OpAdd(*this.At(MkString("version")), url))
		if IsTrue(OpEq2(api, MkString("public"))) {
			if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
				OpAddAssign(&url, OpAdd(MkString("?"), this.Urlencode(query)))
			}
		} else {
			if IsTrue(OpEq2(api, MkString("private"))) {
				this.CheckRequiredCredentials()
				nonce := (this.Nonce()).Call(MkString("toString"))
				_ = nonce
				headers = MkMap(&VarMap{
					"X-ACCESS-KEY":   *this.At(MkString("apiKey")),
					"X-ACCESS-NONCE": nonce,
				})
				auth := OpAdd(nonce, OpAdd(MkString(":"), method))
				_ = auth
				if IsTrue(OpEq2(method, MkString("GET"))) {
					if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
						OpAddAssign(&url, OpAdd(MkString("?"), this.Urlencode(query)))
					}
					OpAddAssign(&auth, url)
				} else {
					*(headers).At(MkString("Content-Type")) = MkString("application/json")
					body = this.Json(query)
					OpAddAssign(&auth, OpAdd(url, body))
				}
				signature := this.Hmac(this.Encode(auth), this.Encode(*this.At(MkString("secret"))))
				_ = signature
				*(headers).At(MkString("X-ACCESS-SIGN")) = OpCopy(signature)
			}
		}
	}
	url = OpAdd(this.ImplodeHostname(*(*(*this.At(MkString("urls"))).At(MkString("api"))).At(api)), url)
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Aax) HandleErrors(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
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
	errorCode := this.SafeString(response, MkString("code"))
	_ = errorCode
	if IsTrue(OpAnd(OpNotEq2(errorCode, MkUndefined()), OpNotEq2(errorCode, MkString("1")))) {
		feedback := OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), this.Json(response)))
		_ = feedback
		this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("exact")), errorCode, feedback)
		this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("broad")), body, feedback)
	}
	return MkUndefined()
}
