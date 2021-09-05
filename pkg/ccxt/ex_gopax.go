package ccxt

import ()

type Gopax struct {
	*ExchangeBase
}

var _ Exchange = (*Gopax)(nil)

func init() {
	exchange := &Gopax{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Gopax) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("gopax"),
		"name": MkString("GOPAX"),
		"countries": MkArray(&VarArray{
			MkString("KR"),
		}),
		"version":   MkString("v1"),
		"rateLimit": MkInteger(50),
		"hostname":  MkString("gopax.co.kr"),
		"certified": MkBool(true),
		"pro":       MkBool(true),
		"has": MkMap(&VarMap{
			"cancelOrder":           MkBool(true),
			"createMarketOrder":     MkBool(true),
			"createOrder":           MkBool(true),
			"fetchBalance":          MkBool(true),
			"fetchCurrencies":       MkBool(true),
			"fetchDepositAddress":   MkString("emulated"),
			"fetchDepositAddresses": MkBool(true),
			"fetchMarkets":          MkBool(true),
			"fetchMyTrades":         MkBool(true),
			"fetchOHLCV":            MkBool(true),
			"fetchOpenOrders":       MkBool(true),
			"fetchOrder":            MkBool(true),
			"fetchOrderBook":        MkBool(true),
			"fetchOrders":           MkBool(true),
			"fetchTicker":           MkBool(true),
			"fetchTickers":          MkBool(true),
			"fetchTime":             MkBool(true),
			"fetchTrades":           MkBool(true),
			"fetchTransactions":     MkBool(true),
		}),
		"timeframes": MkMap(&VarMap{
			"1m":  MkString("1"),
			"5m":  MkString("5"),
			"30m": MkString("30"),
			"1d":  MkString("1440"),
		}),
		"urls": MkMap(&VarMap{
			"logo": MkString("https://user-images.githubusercontent.com/1294454/102897212-ae8a5e00-4478-11eb-9bab-91507c643900.jpg"),
			"api": MkMap(&VarMap{
				"public":  MkString("https://api.{hostname}"),
				"private": MkString("https://api.{hostname}"),
			}),
			"www":  MkString("https://www.gopax.co.kr"),
			"doc":  MkString("https://gopax.github.io/API/index.en.html"),
			"fees": MkString("https://www.gopax.com/feeinfo"),
		}),
		"api": MkMap(&VarMap{
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("notices"),
				MkString("assets"),
				MkString("price-tick-size"),
				MkString("trading-pairs"),
				MkString("trading-pairs/{tradingPair}/ticker"),
				MkString("trading-pairs/{tradingPair}/book"),
				MkString("trading-pairs/{tradingPair}/trades"),
				MkString("trading-pairs/{tradingPair}/stats"),
				MkString("trading-pairs/{tradingPair}/price-tick-size"),
				MkString("trading-pairs/stats"),
				MkString("trading-pairs/{tradingPair}/candles"),
				MkString("time"),
			})}),
			"private": MkMap(&VarMap{
				"get": MkArray(&VarArray{
					MkString("balances"),
					MkString("balances/{assetName}"),
					MkString("orders"),
					MkString("orders/{orderId}"),
					MkString("orders/clientOrderId/{clientOrderId}"),
					MkString("trades"),
					MkString("deposit-withdrawal-status"),
					MkString("crypto-deposit-addresses"),
					MkString("crypto-withdrawal-addresses"),
				}),
				"post": MkArray(&VarArray{
					MkString("orders"),
				}),
				"delete": MkArray(&VarArray{
					MkString("orders/{orderId}"),
					MkString("orders/clientOrderId/{clientOrderId}"),
				}),
			}),
		}),
		"fees": MkMap(&VarMap{"trading": MkMap(&VarMap{
			"percentage": MkBool(true),
			"tierBased":  MkBool(false),
			"maker":      OpDiv(MkNumber(0.04), MkInteger(100)),
			"taker":      OpDiv(MkNumber(0.04), MkInteger(100)),
		})}),
		"exceptions": MkMap(&VarMap{
			"broad": MkMap(&VarMap{
				"ERROR_INVALID_ORDER_TYPE":   InvalidOrder,
				"ERROR_INVALID_AMOUNT":       InvalidOrder,
				"ERROR_INVALID_TRADING_PAIR": BadSymbol,
				"No such order ID":           OrderNotFound,
				"Forbidden order type":       InvalidOrder,
				"the client order ID will be reusable which order has already been completed or canceled": InvalidOrder,
				"ERROR_NO_SUCH_TRADING_PAIR":              BadSymbol,
				"ERROR_INVALID_ORDER_SIDE":                InvalidOrder,
				"ERROR_NOT_HEDGE_TOKEN_USER":              InvalidOrder,
				"ORDER_EVENT_ERROR_NOT_ALLOWED_BID_ORDER": InvalidOrder,
				"ORDER_EVENT_ERROR_INSUFFICIENT_BALANCE":  InsufficientFunds,
				"Invalid option combination":              InvalidOrder,
				"No such client order ID":                 OrderNotFound,
			}),
			"exact": MkMap(&VarMap{
				"100":   BadSymbol,
				"101":   BadSymbol,
				"103":   InvalidOrder,
				"104":   BadSymbol,
				"105":   BadSymbol,
				"106":   BadSymbol,
				"107":   InvalidOrder,
				"108":   InvalidOrder,
				"111":   InvalidOrder,
				"201":   InsufficientFunds,
				"202":   InvalidOrder,
				"203":   InvalidOrder,
				"204":   InvalidOrder,
				"205":   InvalidOrder,
				"206":   InvalidOrder,
				"10004": AuthenticationError,
				"10041": BadRequest,
				"10056": BadRequest,
				"10057": BadSymbol,
				"10059": BadSymbol,
				"10062": BadRequest,
				"10069": OrderNotFound,
				"10155": AuthenticationError,
				"10166": BadRequest,
				"10212": InvalidOrder,
				"10221": OrderNotFound,
				"10222": InvalidOrder,
				"10223": InvalidOrder,
				"10227": InvalidOrder,
				"10319": BadRequest,
				"10358": InvalidOrder,
				"10359": InvalidOrder,
				"10360": InvalidOrder,
				"10361": InvalidOrder,
				"10362": InvalidOrder,
				"10363": InvalidOrder,
			}),
		}),
		"options": MkMap(&VarMap{"createMarketBuyOrderRequiresPrice": MkBool(true)}),
	}))
}

func (this *Gopax) FetchTime(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetTime"), params)
	_ = response
	return this.SafeInteger(response, MkString("serverTime"))
}

func (this *Gopax) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetTradingPairs"), params)
	_ = response
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
		market := *(response).At(i)
		_ = market
		id := this.SafeString(market, MkString("name"))
		_ = id
		numericId := this.SafeInteger(market, MkString("id"))
		_ = numericId
		baseId := this.SafeString(market, MkString("baseAsset"))
		_ = baseId
		quoteId := this.SafeString(market, MkString("quoteAsset"))
		_ = quoteId
		base := this.SafeCurrencyCode(baseId)
		_ = base
		quote := this.SafeCurrencyCode(quoteId)
		_ = quote
		symbol := OpAdd(base, OpAdd(MkString("/"), quote))
		_ = symbol
		precision := MkMap(&VarMap{
			"price":  this.SafeInteger(market, MkString("quoteAssetScale")),
			"amount": this.SafeInteger(market, MkString("baseAssetScale")),
		})
		_ = precision
		minimums := this.SafeValue(market, MkString("restApiOrderAmountMin"), MkMap(&VarMap{}))
		_ = minimums
		marketAsk := this.SafeValue(minimums, MkString("marketAsk"), MkMap(&VarMap{}))
		_ = marketAsk
		marketBid := this.SafeValue(minimums, MkString("marketBid"), MkMap(&VarMap{}))
		_ = marketBid
		takerFeePercentString := this.SafeString(market, MkString("takerFeePercent"))
		_ = takerFeePercentString
		makerFeePercentString := this.SafeString(market, MkString("makerFeePercent"))
		_ = makerFeePercentString
		taker := this.ParseNumber(Precise.StringDiv(takerFeePercentString, MkString("100")))
		_ = taker
		maker := this.ParseNumber(Precise.StringDiv(makerFeePercentString, MkString("100")))
		_ = maker
		result.Push(MkMap(&VarMap{
			"id":        id,
			"info":      market,
			"numericId": numericId,
			"symbol":    symbol,
			"base":      base,
			"quote":     quote,
			"baseId":    this.SafeString(market, MkString("baseAsset")),
			"quoteId":   this.SafeString(market, MkString("quoteAsset")),
			"active":    MkBool(true),
			"taker":     taker,
			"maker":     maker,
			"precision": precision,
			"limits": MkMap(&VarMap{
				"amount": MkMap(&VarMap{
					"min": this.SafeNumber(marketAsk, MkString("amount")),
					"max": MkUndefined(),
				}),
				"price": MkMap(&VarMap{
					"min": this.SafeNumber(market, MkString("priceMin")),
					"max": MkUndefined(),
				}),
				"cost": MkMap(&VarMap{
					"min": this.SafeNumber(marketBid, MkString("amount")),
					"max": MkUndefined(),
				}),
			}),
		}))
	}
	return result
}

func (this *Gopax) FetchCurrencies(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetAssets"), params)
	_ = response
	result := MkMap(&VarMap{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
		currency := *(response).At(i)
		_ = currency
		id := this.SafeString(currency, MkString("id"))
		_ = id
		code := this.SafeCurrencyCode(id)
		_ = code
		name := this.SafeString(currency, MkString("name"))
		_ = name
		fee := this.SafeNumber(currency, MkString("withdrawalFee"))
		_ = fee
		precision := this.SafeNumber(currency, MkString("scale"))
		_ = precision
		*(result).At(code) = MkMap(&VarMap{
			"id":        id,
			"info":      currency,
			"code":      code,
			"name":      name,
			"active":    MkBool(true),
			"fee":       fee,
			"precision": precision,
			"limits": MkMap(&VarMap{
				"amount": MkMap(&VarMap{
					"min": MkUndefined(),
					"max": MkUndefined(),
				}),
				"withdraw": MkMap(&VarMap{
					"min": this.SafeNumber(currency, MkString("withdrawalAmountMin")),
					"max": MkUndefined(),
				}),
			}),
		})
	}
	return result
}

func (this *Gopax) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"tradingPair": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("publicGetTradingPairsTradingPairBook"), this.Extend(request, params))
	_ = response
	nonce := this.SafeInteger(response, MkString("sequence"))
	_ = nonce
	result := this.ParseOrderBook(response, symbol, MkUndefined(), MkString("bid"), MkString("ask"), MkInteger(1), MkInteger(2))
	_ = result
	*(result).At(MkString("nonce")) = OpCopy(nonce)
	return result
}

func (this *Gopax) ParseTicker(goArgs ...*Variant) *Variant {
	ticker := GoGetArg(goArgs, 0, MkUndefined())
	_ = ticker
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	marketId := this.SafeString(ticker, MkString("name"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market, MkString("-"))
	_ = symbol
	timestamp := this.Parse8601(this.SafeString(ticker, MkString("time")))
	_ = timestamp
	open := this.SafeNumber(ticker, MkString("open"))
	_ = open
	last := this.SafeNumber2(ticker, MkString("price"), MkString("close"))
	_ = last
	change := OpCopy(MkUndefined())
	_ = change
	percentage := OpCopy(MkUndefined())
	_ = percentage
	average := OpCopy(MkUndefined())
	_ = average
	if IsTrue(OpAnd(OpNotEq2(last, MkUndefined()), OpNotEq2(open, MkUndefined()))) {
		average = OpDiv(this.Sum(last, open), MkInteger(2))
		change = OpSub(last, open)
		if IsTrue(OpGt(open, MkInteger(0))) {
			percentage = OpDiv(change, OpMulti(open, MkInteger(100)))
		}
	}
	baseVolume := this.SafeNumber(ticker, MkString("volume"))
	_ = baseVolume
	quoteVolume := this.SafeNumber(ticker, MkString("quoteVolume"))
	_ = quoteVolume
	vwap := this.Vwap(baseVolume, quoteVolume)
	_ = vwap
	return MkMap(&VarMap{
		"symbol":        symbol,
		"info":          ticker,
		"timestamp":     timestamp,
		"datetime":      this.Iso8601(timestamp),
		"high":          this.SafeNumber(ticker, MkString("high")),
		"low":           this.SafeNumber(ticker, MkString("low")),
		"bid":           this.SafeNumber(ticker, MkString("bid")),
		"bidVolume":     this.SafeNumber(ticker, MkString("bidVolume")),
		"ask":           this.SafeNumber(ticker, MkString("ask")),
		"askVolume":     this.SafeNumber(ticker, MkString("askVolume")),
		"vwap":          vwap,
		"open":          open,
		"close":         last,
		"last":          last,
		"previousClose": MkUndefined(),
		"change":        change,
		"percentage":    percentage,
		"average":       average,
		"baseVolume":    baseVolume,
		"quoteVolume":   quoteVolume,
	})
}

func (this *Gopax) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"tradingPair": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("publicGetTradingPairsTradingPairTicker"), this.Extend(request, params))
	_ = response
	return this.ParseTicker(response, market)
}

func (this *Gopax) FetchTickers(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("publicGetTradingPairsStats"), params)
	_ = response
	return this.ParseTickers(response, symbols)
}

func (this *Gopax) ParsePublicTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.Parse8601(this.SafeString(trade, MkString("time")))
	_ = timestamp
	price := this.SafeNumber(trade, MkString("price"))
	_ = price
	amount := this.SafeNumber(trade, MkString("amount"))
	_ = amount
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(OpHasMember(MkString("symbol"), market)) {
		symbol = this.SafeString(market, MkString("symbol"))
	}
	return MkMap(&VarMap{
		"info":         trade,
		"id":           this.SafeString(trade, MkString("id")),
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       symbol,
		"order":        MkUndefined(),
		"type":         MkUndefined(),
		"side":         this.SafeString(trade, MkString("side")),
		"takerOrMaker": MkUndefined(),
		"price":        price,
		"amount":       amount,
		"cost":         OpMulti(price, amount),
		"fee":          MkUndefined(),
	})
}

func (this *Gopax) ParsePrivateTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.Parse8601(this.SafeString(trade, MkString("timestamp")))
	_ = timestamp
	symbol := (this.SafeString(trade, MkString("tradingPairName"))).Call(MkString("replace"), MkString("-"), MkString("/"))
	_ = symbol
	side := this.SafeString(trade, MkString("side"))
	_ = side
	price := this.SafeNumber(trade, MkString("price"))
	_ = price
	amount := this.SafeNumber(trade, MkString("baseAmount"))
	_ = amount
	feeCurrency := symbol.Slice(MkInteger(0), MkInteger(3))
	_ = feeCurrency
	if IsTrue(OpEq2(side, MkString("sell"))) {
		feeCurrency = symbol.Slice(MkInteger(4))
	}
	fee := MkMap(&VarMap{
		"cost":     this.SafeNumber(trade, MkString("fee")),
		"currency": feeCurrency,
		"rate":     MkUndefined(),
	})
	_ = fee
	return MkMap(&VarMap{
		"info":         trade,
		"id":           this.SafeString(trade, MkString("id")),
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       symbol,
		"order":        this.SafeInteger(trade, MkString("orderId")),
		"type":         MkUndefined(),
		"side":         side,
		"takerOrMaker": this.SafeString(trade, MkString("position")),
		"price":        price,
		"amount":       amount,
		"cost":         OpMulti(price, amount),
		"fee":          fee,
	})
}

func (this *Gopax) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	id := this.SafeString2(trade, MkString("id"), MkString("tradeId"))
	_ = id
	orderId := this.SafeInteger(trade, MkString("orderId"))
	_ = orderId
	timestamp := this.Parse8601(this.SafeString2(trade, MkString("time"), MkString("timestamp")))
	_ = timestamp
	timestamp = this.SafeTimestamp(trade, MkString("occuredAt"), timestamp)
	marketId := this.SafeString(trade, MkString("tradingPairName"))
	_ = marketId
	market = this.SafeMarket(marketId, market, MkString("-"))
	symbol := *(market).At(MkString("symbol"))
	_ = symbol
	side := this.SafeString(trade, MkString("side"))
	_ = side
	if IsTrue(OpEq2(side, MkString("1"))) {
		side = MkString("buy")
	} else {
		if IsTrue(OpEq2(side, MkString("2"))) {
			side = MkString("sell")
		}
	}
	type_ := this.SafeString(trade, MkString("type"))
	_ = type_
	if IsTrue(OpEq2(type_, MkString("1"))) {
		type_ = MkString("limit")
	} else {
		if IsTrue(OpEq2(type_, MkString("2"))) {
			type_ = MkString("market")
		}
	}
	priceString := this.SafeString(trade, MkString("price"))
	_ = priceString
	amountString := this.SafeString2(trade, MkString("amount"), MkString("baseAmount"))
	_ = amountString
	price := this.ParseNumber(priceString)
	_ = price
	amount := this.ParseNumber(amountString)
	_ = amount
	cost := this.SafeNumber(trade, MkString("quoteAmount"))
	_ = cost
	if IsTrue(OpEq2(cost, MkUndefined())) {
		cost = this.ParseNumber(Precise.StringMul(priceString, amountString))
	}
	feeCost := this.SafeNumber(trade, MkString("fee"))
	_ = feeCost
	fee := OpCopy(MkUndefined())
	_ = fee
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		fee = MkMap(&VarMap{
			"cost":     feeCost,
			"currency": *(market).At(MkString("base")),
		})
	}
	takerOrMaker := this.SafeString(trade, MkString("position"))
	_ = takerOrMaker
	return MkMap(&VarMap{
		"info":         trade,
		"id":           id,
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       symbol,
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

func (this *Gopax) FetchTrades(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"tradingPair": *(market).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("after")) = parseInt(OpDiv(since, MkInteger(1000)))
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("publicGetTradingPairsTradingPairTrades"), this.Extend(request, params))
	_ = response
	return this.ParseTrades(response, market, since, limit)
}

func (this *Gopax) ParseOHLCV(goArgs ...*Variant) *Variant {
	ohlcv := GoGetArg(goArgs, 0, MkUndefined())
	_ = ohlcv
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	return MkArray(&VarArray{
		this.SafeInteger(ohlcv, MkInteger(0)),
		this.SafeNumber(ohlcv, MkInteger(3)),
		this.SafeNumber(ohlcv, MkInteger(2)),
		this.SafeNumber(ohlcv, MkInteger(1)),
		this.SafeNumber(ohlcv, MkInteger(4)),
		this.SafeNumber(ohlcv, MkInteger(5)),
	})
}

func (this *Gopax) FetchOHLCV(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	timeframe := GoGetArg(goArgs, 1, MkString("1m"))
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
	limit = OpTrinary(OpEq2(limit, MkUndefined()), MkInteger(1024), limit)
	request := MkMap(&VarMap{
		"tradingPair": *(market).At(MkString("id")),
		"interval":    *(*this.At(MkString("timeframes"))).At(timeframe),
	})
	_ = request
	duration := this.ParseTimeframe(timeframe)
	_ = duration
	if IsTrue(OpEq2(since, MkUndefined())) {
		end := this.Milliseconds()
		_ = end
		*(request).At(MkString("end")) = OpCopy(end)
		*(request).At(MkString("start")) = OpSub(end, OpMulti(limit, OpMulti(duration, MkInteger(1000))))
	} else {
		*(request).At(MkString("start")) = OpCopy(since)
		*(request).At(MkString("end")) = this.Sum(since, OpMulti(limit, OpMulti(duration, MkInteger(1000))))
	}
	response := this.Call(MkString("publicGetTradingPairsTradingPairCandles"), this.Extend(request, params))
	_ = response
	return this.ParseOHLCVs(response, market, timeframe, since, limit)
}

func (this *Gopax) ParseBalanceResponse(goArgs ...*Variant) *Variant {
	response := GoGetArg(goArgs, 0, MkUndefined())
	_ = response
	result := MkMap(&VarMap{"info": response})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
		balance := *(response).At(i)
		_ = balance
		currencyId := this.SafeString2(balance, MkString("asset"), MkString("isoAlpha3"))
		_ = currencyId
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		hold := this.SafeString(balance, MkString("hold"))
		_ = hold
		pendingWithdrawal := this.SafeString(balance, MkString("pendingWithdrawal"))
		_ = pendingWithdrawal
		account := this.Account()
		_ = account
		*(account).At(MkString("free")) = this.SafeString(balance, MkString("avail"))
		*(account).At(MkString("used")) = Precise.StringAdd(hold, pendingWithdrawal)
		*(result).At(code) = OpCopy(account)
	}
	return this.ParseBalance(result)
}

func (this *Gopax) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privateGetBalances"), params)
	_ = response
	return this.ParseBalanceResponse(response)
}

func (this *Gopax) ParseOrderStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"placed":    MkString("open"),
		"cancelled": MkString("canceled"),
		"completed": MkString("closed"),
		"updated":   MkString("open"),
		"reserved":  MkString("open"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Gopax) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	id := this.SafeString(order, MkString("id"))
	_ = id
	clientOrderId := this.SafeString(order, MkString("clientOrderId"))
	_ = clientOrderId
	timestamp := this.Parse8601(this.SafeString(order, MkString("createdAt")))
	_ = timestamp
	type_ := this.SafeString(order, MkString("type"))
	_ = type_
	side := this.SafeString(order, MkString("side"))
	_ = side
	timeInForce := this.SafeStringUpper(order, MkString("timeInForce"))
	_ = timeInForce
	price := this.SafeNumber(order, MkString("price"))
	_ = price
	amount := this.SafeNumber(order, MkString("amount"))
	_ = amount
	stopPrice := this.SafeNumber(order, MkString("stopPrice"))
	_ = stopPrice
	remaining := this.SafeNumber(order, MkString("remaining"))
	_ = remaining
	marketId := this.SafeString(order, MkString("tradingPairName"))
	_ = marketId
	market = this.SafeMarket(marketId, market, MkString("-"))
	status := this.ParseOrderStatus(this.SafeString(order, MkString("status")))
	_ = status
	balanceChange := this.SafeValue(order, MkString("balanceChange"), MkMap(&VarMap{}))
	_ = balanceChange
	filled := this.SafeNumber(balanceChange, MkString("baseNet"))
	_ = filled
	cost := this.SafeNumber(balanceChange, MkString("quoteNet"))
	_ = cost
	if IsTrue(OpNotEq2(cost, MkUndefined())) {
		cost = Math.Abs(cost)
	}
	updated := OpCopy(MkUndefined())
	_ = updated
	if IsTrue(OpAnd(OpNotEq2(filled, MkUndefined()), OpGt(filled, MkInteger(0)))) {
		updated = this.Parse8601(this.SafeString(order, MkString("updatedAt")))
	}
	fee := OpCopy(MkUndefined())
	_ = fee
	if IsTrue(OpEq2(side, MkString("buy"))) {
		baseFee := this.SafeValue(balanceChange, MkString("baseFee"), MkMap(&VarMap{}))
		_ = baseFee
		taking := this.SafeNumber(baseFee, MkString("taking"))
		_ = taking
		making := this.SafeNumber(baseFee, MkString("making"))
		_ = making
		fee = MkMap(&VarMap{
			"currency": *(market).At(MkString("base")),
			"cost":     this.Sum(taking, making),
		})
	} else {
		quoteFee := this.SafeValue(balanceChange, MkString("quoteFee"), MkMap(&VarMap{}))
		_ = quoteFee
		taking := this.SafeNumber(quoteFee, MkString("taking"))
		_ = taking
		making := this.SafeNumber(quoteFee, MkString("making"))
		_ = making
		fee = MkMap(&VarMap{
			"currency": *(market).At(MkString("quote")),
			"cost":     this.Sum(taking, making),
		})
	}
	postOnly := OpCopy(MkUndefined())
	_ = postOnly
	if IsTrue(OpNotEq2(timeInForce, MkUndefined())) {
		postOnly = OpEq2(timeInForce, MkString("PO"))
	}
	return this.SafeOrder(MkMap(&VarMap{
		"id":                 id,
		"clientOrderId":      clientOrderId,
		"datetime":           this.Iso8601(timestamp),
		"timestamp":          timestamp,
		"lastTradeTimestamp": updated,
		"status":             status,
		"symbol":             *(market).At(MkString("symbol")),
		"type":               type_,
		"timeInForce":        timeInForce,
		"postOnly":           postOnly,
		"side":               side,
		"price":              price,
		"stopPrice":          stopPrice,
		"average":            MkUndefined(),
		"amount":             amount,
		"filled":             filled,
		"remaining":          remaining,
		"cost":               cost,
		"trades":             MkUndefined(),
		"fee":                fee,
		"info":               order,
	}))
}

func (this *Gopax) FetchOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	method := OpCopy(MkUndefined())
	_ = method
	clientOrderId := this.SafeString(params, MkString("clientOrderId"))
	_ = clientOrderId
	params = this.Omit(params, MkString("clientOrderId"))
	request := MkMap(&VarMap{})
	_ = request
	if IsTrue(OpEq2(clientOrderId, MkUndefined())) {
		method = MkString("privateGetOrdersOrderId")
		*(request).At(MkString("orderId")) = OpCopy(id)
	} else {
		method = MkString("privateGetOrdersClientOrderIdClientOrderId")
		*(request).At(MkString("clientOrderId")) = OpCopy(clientOrderId)
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	return this.ParseOrder(response)
}

func (this *Gopax) FetchOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"includePast": MkString("true")})
	_ = request
	market := OpCopy(MkUndefined())
	_ = market
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
	}
	response := this.Call(MkString("privateGetOrders"), this.Extend(request, params))
	_ = response
	return this.ParseOrders(response, market, since, limit)
}

func (this *Gopax) FetchOpenOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	request := MkMap(&VarMap{"includePast": MkString("false")})
	_ = request
	return this.FetchOrders(symbol, since, limit, this.Extend(request, params))
}

func (this *Gopax) CreateOrder(goArgs ...*Variant) *Variant {
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
		"tradingPairName": *(market).At(MkString("id")),
		"side":            side,
		"type":            type_,
	})
	_ = request
	if IsTrue(OpEq2(type_, MkString("limit"))) {
		*(request).At(MkString("price")) = this.PriceToPrecision(symbol, price)
		*(request).At(MkString("amount")) = this.AmountToPrecision(symbol, amount)
	} else {
		if IsTrue(OpEq2(type_, MkString("market"))) {
			if IsTrue(OpEq2(side, MkString("buy"))) {
				total := OpCopy(amount)
				_ = total
				createMarketBuyOrderRequiresPrice := this.SafeValue(*this.At(MkString("options")), MkString("createMarketBuyOrderRequiresPrice"), MkBool(true))
				_ = createMarketBuyOrderRequiresPrice
				if IsTrue(createMarketBuyOrderRequiresPrice) {
					if IsTrue(OpEq2(price, MkUndefined())) {
						panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")), MkString(" createOrder() requires the price argument with market buy orders to calculate total order cost (amount to spend), where cost = amount * price. Supply a price argument to createOrder() call if you want the cost to be calculated for you from price and amount, or, alternatively, add .options['createMarketBuyOrderRequiresPrice'] = false and supply the total cost value in the 'amount' argument"))))
					}
					total = OpMulti(price, amount)
				}
				precision := *(*(market).At(MkString("precision"))).At(MkString("price"))
				_ = precision
				*(request).At(MkString("amount")) = this.DecimalToPrecision(total, TRUNCATE, precision, *this.At(MkString("precisionMode")))
			} else {
				*(request).At(MkString("amount")) = this.AmountToPrecision(symbol, amount)
			}
		}
	}
	clientOrderId := this.SafeString(params, MkString("clientOrderId"))
	_ = clientOrderId
	if IsTrue(OpNotEq2(clientOrderId, MkUndefined())) {
		*(request).At(MkString("clientOrderId")) = OpCopy(clientOrderId)
		params = this.Omit(params, MkString("clientOrderId"))
	}
	stopPrice := this.SafeNumber(params, MkString("stopPrice"))
	_ = stopPrice
	if IsTrue(OpNotEq2(stopPrice, MkUndefined())) {
		*(request).At(MkString("stopPrice")) = this.PriceToPrecision(symbol, stopPrice)
		params = this.Omit(params, MkString("stopPrice"))
	}
	timeInForce := this.SafeStringLower(params, MkString("timeInForce"))
	_ = timeInForce
	if IsTrue(OpNotEq2(timeInForce, MkUndefined())) {
		*(request).At(MkString("timeInForce")) = OpCopy(timeInForce)
		params = this.Omit(params, MkString("timeInForce"))
	}
	response := this.Call(MkString("privatePostOrders"), this.Extend(request, params))
	_ = response
	return this.ParseOrder(response, market)
}

func (this *Gopax) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{})
	_ = request
	clientOrderId := this.SafeString(params, MkString("clientOrderId"))
	_ = clientOrderId
	method := OpCopy(MkUndefined())
	_ = method
	if IsTrue(OpEq2(clientOrderId, MkUndefined())) {
		method = MkString("privateDeleteOrdersOrderId")
		*(request).At(MkString("orderId")) = OpCopy(id)
	} else {
		method = MkString("privateDeleteOrdersClientOrderIdClientOrderId")
		*(request).At(MkString("clientOrderId")) = OpCopy(clientOrderId)
		params = this.Omit(params, MkString("clientOrderId"))
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	order := this.ParseOrder(response)
	_ = order
	return this.Extend(order, MkMap(&VarMap{"id": id}))
}

func (this *Gopax) FetchMyTrades(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"deepSearch": MkString("true")})
	_ = request
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("after")) = parseInt(OpDiv(since, MkInteger(1000)))
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("privateGetTrades"), this.Extend(request, params))
	_ = response
	market := OpCopy(MkUndefined())
	_ = market
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
	}
	return this.ParseTrades(response, market, since, limit)
}

func (this *Gopax) ParseDepositAddress(goArgs ...*Variant) *Variant {
	depositAddress := GoGetArg(goArgs, 0, MkUndefined())
	_ = depositAddress
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	address := this.SafeString(depositAddress, MkString("address"))
	_ = address
	tag := this.SafeString(depositAddress, MkString("memoId"))
	_ = tag
	currencyId := this.SafeString(depositAddress, MkString("asset"))
	_ = currencyId
	code := this.SafeCurrencyCode(currencyId)
	_ = code
	this.CheckAddress(address)
	return MkMap(&VarMap{
		"currency": code,
		"address":  address,
		"tag":      tag,
		"info":     depositAddress,
	})
}

func (this *Gopax) FetchDepositAddresses(goArgs ...*Variant) *Variant {
	codes := GoGetArg(goArgs, 0, MkUndefined())
	_ = codes
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privateGetCryptoDepositAddresses"), params)
	_ = response
	return this.ParseDepositAddresses(response, codes)
}

func (this *Gopax) FetchDepositAddress(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.FetchDepositAddresses(MkUndefined(), params)
	_ = response
	address := this.SafeValue(response, code)
	_ = address
	if IsTrue(OpEq2(address, MkUndefined())) {
		panic(NewInvalidAddress(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" fetchDepositAddress() "), OpAdd(code, MkString(" address not found"))))))
	}
	return address
}

func (this *Gopax) ParseTransactionStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"reviewing":  MkString("pending"),
		"rejected":   MkString("rejected"),
		"processing": MkString("pending"),
		"failed":     MkString("failed"),
		"completed":  MkString("ok"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Gopax) ParseTransaction(goArgs ...*Variant) *Variant {
	transaction := GoGetArg(goArgs, 0, MkUndefined())
	_ = transaction
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	id := this.SafeString(transaction, MkString("id"))
	_ = id
	txid := this.SafeString(transaction, MkString("txId"))
	_ = txid
	currencyId := this.SafeString(transaction, MkString("asset"))
	_ = currencyId
	code := this.SafeCurrencyCode(currencyId, currency)
	_ = code
	type_ := this.SafeString(transaction, MkString("type"))
	_ = type_
	if IsTrue(OpOr(OpEq2(type_, MkString("crypto_withdrawal")), OpEq2(type_, MkString("fiat_withdrawal")))) {
		type_ = MkString("withdrawal")
	} else {
		if IsTrue(OpOr(OpEq2(type_, MkString("crypto_deposit")), OpEq2(type_, MkString("fiat_deposit")))) {
			type_ = MkString("deposit")
		}
	}
	amount := this.SafeNumber(transaction, MkString("netAmount"))
	_ = amount
	feeCost := this.SafeNumber(transaction, MkString("feeAmount"))
	_ = feeCost
	fee := OpCopy(MkUndefined())
	_ = fee
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		fee = MkMap(&VarMap{
			"cost":     feeCost,
			"currency": code,
		})
	}
	timestamp := this.SafeTimestamp(transaction, MkString("reviewStartedAt"))
	_ = timestamp
	updated := this.SafeTimestamp(transaction, MkString("completedAt"))
	_ = updated
	addressFrom := this.SafeString(transaction, MkString("sourceAddress"))
	_ = addressFrom
	addressTo := this.SafeString(transaction, MkString("destinationAddress"))
	_ = addressTo
	tagFrom := this.SafeString(transaction, MkString("sourceMemoId"))
	_ = tagFrom
	tagTo := this.SafeString(transaction, MkString("destinationMemoId"))
	_ = tagTo
	status := this.ParseTransactionStatus(this.SafeString(transaction, MkString("status")))
	_ = status
	return MkMap(&VarMap{
		"info":        transaction,
		"id":          id,
		"txid":        txid,
		"timestamp":   timestamp,
		"datetime":    this.Iso8601(timestamp),
		"addressFrom": addressFrom,
		"address":     addressTo,
		"addressTo":   addressTo,
		"tagFrom":     tagFrom,
		"tag":         tagTo,
		"tagTo":       tagTo,
		"type":        type_,
		"amount":      amount,
		"currency":    code,
		"status":      status,
		"updated":     updated,
		"comment":     MkUndefined(),
		"fee":         fee,
	})
}

func (this *Gopax) FetchTransactions(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{})
	_ = request
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("before")) = OpCopy(since)
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("privateGetDepositWithdrawalStatus"), this.Extend(request, params))
	_ = response
	currency := OpCopy(MkUndefined())
	_ = currency
	if IsTrue(OpNotEq2(code, MkUndefined())) {
		currency = this.Currency(code)
	}
	return this.ParseTransactions(response, currency, since, limit, params)
}

func (this *Gopax) Nonce(goArgs ...*Variant) *Variant {
	return this.Milliseconds()
}

func (this *Gopax) Sign(goArgs ...*Variant) *Variant {
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
	endpoint := OpAdd(MkString("/"), this.ImplodeParams(path, params))
	_ = endpoint
	url := OpAdd(this.ImplodeHostname(*(*(*this.At(MkString("urls"))).At(MkString("api"))).At(api)), endpoint)
	_ = url
	query := this.Omit(params, this.ExtractParams(path))
	_ = query
	if IsTrue(OpEq2(api, MkString("public"))) {
		if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
			OpAddAssign(&url, OpAdd(MkString("?"), this.Urlencode(query)))
		}
	} else {
		if IsTrue(OpEq2(api, MkString("private"))) {
			this.CheckRequiredCredentials()
			timestamp := (this.Nonce()).Call(MkString("toString"))
			_ = timestamp
			auth := OpAdd(MkString("t"), OpAdd(timestamp, OpAdd(method, endpoint)))
			_ = auth
			headers = MkMap(&VarMap{
				"api-key":   *this.At(MkString("apiKey")),
				"timestamp": timestamp,
			})
			if IsTrue(OpEq2(method, MkString("POST"))) {
				*(headers).At(MkString("Content-Type")) = MkString("application/json")
				body = this.Json(params)
				OpAddAssign(&auth, body)
			} else {
				if IsTrue(OpEq2(endpoint, MkString("/orders"))) {
					if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
						urlQuery := OpAdd(MkString("?"), this.Urlencode(query))
						_ = urlQuery
						OpAddAssign(&auth, urlQuery)
						OpAddAssign(&url, urlQuery)
					}
				} else {
					if IsTrue(OpEq2(method, MkString("GET"))) {
						if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
							OpAddAssign(&url, OpAdd(MkString("?"), this.Urlencode(query)))
						}
					}
				}
			}
			rawSecret := this.Base64ToBinary(*this.At(MkString("secret")))
			_ = rawSecret
			signature := this.Hmac(this.Encode(auth), rawSecret, MkString("sha512"), MkString("base64"))
			_ = signature
			*(headers).At(MkString("signature")) = OpCopy(signature)
		}
	}
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Gopax) HandleErrors(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpNot(GoIsArray(response))) {
		errorCode := this.SafeString(response, MkString("errorCode"))
		_ = errorCode
		errorMessage := this.SafeString(response, MkString("errorMessage"))
		_ = errorMessage
		feedback := OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))
		_ = feedback
		if IsTrue(OpNotEq2(errorMessage, MkUndefined())) {
			this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("broad")), body, feedback)
		}
		if IsTrue(OpNotEq2(errorCode, MkUndefined())) {
			this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("exact")), errorCode, feedback)
		}
		if IsTrue(OpOr(OpNotEq2(errorCode, MkUndefined()), OpNotEq2(errorMessage, MkUndefined()))) {
			panic(NewExchangeError(feedback))
		}
	}
	return MkUndefined()
}
