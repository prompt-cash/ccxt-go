package ccxt

import ()

type Vcc struct {
	*ExchangeBase
}

var _ Exchange = (*Vcc)(nil)

func init() {
	exchange := &Vcc{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Vcc) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("vcc"),
		"name": MkString("VCC Exchange"),
		"countries": MkArray(&VarArray{
			MkString("VN"),
		}),
		"rateLimit": MkInteger(1000),
		"version":   MkString("v3"),
		"has": MkMap(&VarMap{
			"cancelAllOrders":     MkBool(true),
			"cancelOrder":         MkBool(true),
			"createOrder":         MkBool(true),
			"editOrder":           MkBool(false),
			"fetchBalance":        MkBool(true),
			"fetchClosedOrders":   MkBool(true),
			"fetchCurrencies":     MkBool(true),
			"fetchDepositAddress": MkBool(true),
			"fetchDeposits":       MkBool(true),
			"fetchMarkets":        MkBool(true),
			"fetchMyTrades":       MkBool(true),
			"fetchOHLCV":          MkBool(true),
			"fetchOpenOrders":     MkBool(true),
			"fetchOrder":          MkBool(true),
			"fetchOrderBook":      MkBool(true),
			"fetchOrders":         MkBool(false),
			"fetchTicker":         MkString("emulated"),
			"fetchTickers":        MkBool(true),
			"fetchTrades":         MkBool(true),
			"fetchTradingFees":    MkBool(false),
			"fetchTransactions":   MkBool(true),
			"fetchWithdrawals":    MkBool(true),
		}),
		"timeframes": MkMap(&VarMap{
			"1m":  MkString("60000"),
			"5m":  MkString("300000"),
			"15m": MkString("900000"),
			"30m": MkString("1800000"),
			"1h":  MkString("3600000"),
			"2h":  MkString("7200000"),
			"4h":  MkString("14400000"),
			"6h":  MkString("21600000"),
			"12h": MkString("43200000"),
			"1d":  MkString("86400000"),
			"1w":  MkString("604800000"),
		}),
		"urls": MkMap(&VarMap{
			"logo": MkString("https://user-images.githubusercontent.com/1294454/100545356-8427f500-326c-11eb-9539-7d338242d61b.jpg"),
			"api": MkMap(&VarMap{
				"public":  MkString("https://api.vcc.exchange"),
				"private": MkString("https://api.vcc.exchange"),
			}),
			"www": MkString("https://vcc.exchange"),
			"doc": MkArray(&VarArray{
				MkString("https://vcc.exchange/api"),
			}),
			"fees":     MkString("https://support.vcc.exchange/hc/en-us/articles/360016401754"),
			"referral": MkString("https://vcc.exchange?ref=l4xhrH"),
		}),
		"api": MkMap(&VarMap{
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("summary"),
				MkString("exchange_info"),
				MkString("assets"),
				MkString("ticker"),
				MkString("trades/{market_pair}"),
				MkString("orderbook/{market_pair}"),
				MkString("chart/bars"),
				MkString("tick_sizes"),
			})}),
			"private": MkMap(&VarMap{
				"get": MkArray(&VarArray{
					MkString("user"),
					MkString("balance"),
					MkString("orders/{order_id}"),
					MkString("orders/open"),
					MkString("orders"),
					MkString("orders/trades"),
					MkString("deposit-address"),
					MkString("transactions"),
				}),
				"post": MkArray(&VarArray{
					MkString("orders"),
				}),
				"put": MkArray(&VarArray{
					MkString("orders/{order_id}/cancel"),
					MkString("orders/cancel-by-type"),
					MkString("orders/cancel-all"),
				}),
			}),
		}),
		"fees": MkMap(&VarMap{"trading": MkMap(&VarMap{
			"tierBased":  MkBool(false),
			"percentage": MkBool(true),
			"maker":      this.ParseNumber(MkString("0.002")),
			"taker":      this.ParseNumber(MkString("0.002")),
		})}),
		"exceptions": MkMap(&VarMap{
			"exact": MkMap(&VarMap{}),
			"broad": MkMap(&VarMap{
				"limit may not be greater than": BadRequest,
				"Insufficient balance":          InsufficientFunds,
				"Unauthenticated":               AuthenticationError,
				"signature is invalid":          AuthenticationError,
				"Timeout":                       RequestTimeout,
				"Too many requests":             RateLimitExceeded,
				"quantity field is required":    InvalidOrder,
				"price field is required":       InvalidOrder,
				"error_security_level":          PermissionDenied,
				"pair is invalid":               BadSymbol,
				"type is invalid":               InvalidOrder,
				"Data not found":                OrderNotFound,
			}),
		}),
	}))
}

func (this *Vcc) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetExchangeInfo"), params)
	_ = response
	data := this.SafeValue(response, MkString("data"))
	_ = data
	markets := this.SafeValue(data, MkString("symbols"))
	_ = markets
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, markets.Length)); OpIncr(&i) {
		market := this.SafeValue(markets, i)
		_ = market
		symbol := this.SafeString(market, MkString("symbol"))
		_ = symbol
		id := symbol.Replace(MkString("/"), MkString("_"))
		_ = id
		baseId := this.SafeString(market, MkString("coin"))
		_ = baseId
		quoteId := this.SafeString(market, MkString("currency"))
		_ = quoteId
		base := this.SafeCurrencyCode(baseId)
		_ = base
		quote := this.SafeCurrencyCode(quoteId)
		_ = quote
		active := this.SafeValue(market, MkString("active"))
		_ = active
		precision := this.SafeValue(market, MkString("precision"), MkMap(&VarMap{}))
		_ = precision
		limits := this.SafeValue(market, MkString("limits"), MkMap(&VarMap{}))
		_ = limits
		amountLimits := this.SafeValue(limits, MkString("amount"), MkMap(&VarMap{}))
		_ = amountLimits
		priceLimits := this.SafeValue(limits, MkString("price"), MkMap(&VarMap{}))
		_ = priceLimits
		costLimits := this.SafeValue(limits, MkString("cost"), MkMap(&VarMap{}))
		_ = costLimits
		entry := MkMap(&VarMap{
			"info":    market,
			"id":      id,
			"symbol":  symbol,
			"base":    base,
			"quote":   quote,
			"baseId":  baseId,
			"quoteId": quoteId,
			"active":  active,
			"precision": MkMap(&VarMap{
				"price":  this.SafeInteger(precision, MkString("price")),
				"amount": this.SafeInteger(precision, MkString("amount")),
				"cost":   this.SafeInteger(precision, MkString("cost")),
			}),
			"limits": MkMap(&VarMap{
				"amount": MkMap(&VarMap{
					"min": this.SafeNumber(amountLimits, MkString("min")),
					"max": MkUndefined(),
				}),
				"price": MkMap(&VarMap{
					"min": this.SafeNumber(priceLimits, MkString("min")),
					"max": MkUndefined(),
				}),
				"cost": MkMap(&VarMap{
					"min": this.SafeNumber(costLimits, MkString("min")),
					"max": MkUndefined(),
				}),
			}),
		})
		_ = entry
		result.Push(entry)
	}
	return result
}

func (this *Vcc) FetchCurrencies(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetAssets"), params)
	_ = response
	result := MkMap(&VarMap{})
	_ = result
	data := this.SafeValue(response, MkString("data"))
	_ = data
	ids := GoGetKeys(data)
	_ = ids
	for i := MkInteger(0); IsTrue(OpLw(i, ids.Length)); OpIncr(&i) {
		id := this.SafeStringLower(ids, i)
		_ = id
		currency := this.SafeValue(data, *(ids).At(i))
		_ = currency
		code := this.SafeCurrencyCode(id)
		_ = code
		canDeposit := this.SafeValue(currency, MkString("can_deposit"))
		_ = canDeposit
		canWithdraw := this.SafeValue(currency, MkString("can_withdraw"))
		_ = canWithdraw
		active := OpAnd(canDeposit, canWithdraw)
		_ = active
		*(result).At(code) = MkMap(&VarMap{
			"id":        id,
			"code":      code,
			"name":      this.SafeString(currency, MkString("name")),
			"active":    active,
			"fee":       this.SafeNumber(currency, MkString("withdrawal_fee")),
			"precision": this.SafeInteger(currency, MkString("decimal")),
			"limits": MkMap(&VarMap{"withdraw": MkMap(&VarMap{
				"min": this.SafeNumber(currency, MkString("min_withdraw")),
				"max": this.SafeNumber(currency, MkString("max_withdraw")),
			})}),
		})
	}
	return result
}

func (this *Vcc) FetchTradingFee(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := this.Extend(MkMap(&VarMap{"symbol": *(market).At(MkString("id"))}), this.Omit(params, MkString("symbol")))
	_ = request
	response := this.Call(MkString("privateGetTradingFeeSymbol"), request)
	_ = response
	return MkMap(&VarMap{
		"info":  response,
		"maker": this.SafeNumber(response, MkString("provideLiquidityRate")),
		"taker": this.SafeNumber(response, MkString("takeLiquidityRate")),
	})
}

func (this *Vcc) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privateGetBalance"), params)
	_ = response
	data := this.SafeValue(response, MkString("data"))
	_ = data
	result := MkMap(&VarMap{
		"info":      response,
		"timestamp": MkUndefined(),
		"datetime":  MkUndefined(),
	})
	_ = result
	currencyIds := GoGetKeys(data)
	_ = currencyIds
	for i := MkInteger(0); IsTrue(OpLw(i, currencyIds.Length)); OpIncr(&i) {
		currencyId := *(currencyIds).At(i)
		_ = currencyId
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		balance := this.SafeValue(data, currencyId)
		_ = balance
		account := this.Account()
		_ = account
		*(account).At(MkString("free")) = this.SafeString(balance, MkString("available_balance"))
		*(account).At(MkString("total")) = this.SafeString(balance, MkString("balance"))
		*(result).At(code) = OpCopy(account)
	}
	return this.ParseBalance(result)
}

func (this *Vcc) ParseOHLCV(goArgs ...*Variant) *Variant {
	ohlcv := GoGetArg(goArgs, 0, MkUndefined())
	_ = ohlcv
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	return MkArray(&VarArray{
		this.SafeInteger(ohlcv, MkString("time")),
		this.SafeNumber(ohlcv, MkString("open")),
		this.SafeNumber(ohlcv, MkString("high")),
		this.SafeNumber(ohlcv, MkString("low")),
		this.SafeNumber(ohlcv, MkString("close")),
		this.SafeNumber(ohlcv, MkString("volume")),
	})
}

func (this *Vcc) FetchOHLCV(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{
		"coin":       *(market).At(MkString("baseId")),
		"currency":   *(market).At(MkString("quoteId")),
		"resolution": *(*this.At(MkString("timeframes"))).At(timeframe),
	})
	_ = request
	limit = OpTrinary(OpEq2(limit, MkUndefined()), MkInteger(100), limit)
	limit = Math.Min(MkInteger(100), limit)
	duration := this.ParseTimeframe(timeframe)
	_ = duration
	if IsTrue(OpEq2(since, MkUndefined())) {
		end := this.Seconds()
		_ = end
		*(request).At(MkString("to")) = OpCopy(end)
		*(request).At(MkString("from")) = OpSub(end, OpMulti(limit, duration))
	} else {
		start := parseInt(OpDiv(since, MkInteger(1000)))
		_ = start
		*(request).At(MkString("from")) = OpCopy(start)
		*(request).At(MkString("to")) = this.Sum(start, OpMulti(limit, duration))
	}
	response := this.Call(MkString("publicGetChartBars"), this.Extend(request, params))
	_ = response
	return this.ParseOHLCVs(response, market, timeframe, since, limit)
}

func (this *Vcc) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{
		"market_pair": *(market).At(MkString("id")),
		"level":       MkInteger(2),
	})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		if IsTrue(OpAnd(OpNotEq2(limit, MkInteger(0)), OpAnd(OpNotEq2(limit, MkInteger(5)), OpAnd(OpNotEq2(limit, MkInteger(10)), OpAnd(OpNotEq2(limit, MkInteger(20)), OpAnd(OpNotEq2(limit, MkInteger(50)), OpAnd(OpNotEq2(limit, MkInteger(100)), OpNotEq2(limit, MkInteger(500))))))))) {
			panic(NewBadRequest(OpAdd(*this.At(MkString("id")), MkString(" fetchOrderBook limit must be 0, 5, 10, 20, 50, 100, 500 if specified"))))
		}
		*(request).At(MkString("depth")) = OpCopy(limit)
	}
	response := this.Call(MkString("publicGetOrderbookMarketPair"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"))
	_ = data
	timestamp := this.SafeValue(data, MkString("timestamp"))
	_ = timestamp
	return this.ParseOrderBook(data, symbol, timestamp, MkString("bids"), MkString("asks"), MkInteger(0), MkInteger(1))
}

func (this *Vcc) ParseTicker(goArgs ...*Variant) *Variant {
	ticker := GoGetArg(goArgs, 0, MkUndefined())
	_ = ticker
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.Milliseconds()
	_ = timestamp
	baseVolume := this.SafeNumber(ticker, MkString("base_volume"))
	_ = baseVolume
	quoteVolume := this.SafeNumber(ticker, MkString("quote_volume"))
	_ = quoteVolume
	open := this.SafeNumber(ticker, MkString("open_price"))
	_ = open
	last := this.SafeNumber(ticker, MkString("last_price"))
	_ = last
	change := OpCopy(MkUndefined())
	_ = change
	percentage := OpCopy(MkUndefined())
	_ = percentage
	average := OpCopy(MkUndefined())
	_ = average
	if IsTrue(OpAnd(OpNotEq2(last, MkUndefined()), OpNotEq2(open, MkUndefined()))) {
		change = OpSub(last, open)
		average = OpDiv(this.Sum(last, open), MkInteger(2))
		if IsTrue(OpGt(open, MkInteger(0))) {
			percentage = OpDiv(change, OpMulti(open, MkInteger(100)))
		}
	}
	vwap := this.Vwap(baseVolume, quoteVolume)
	_ = vwap
	symbol := OpTrinary(OpEq2(market, MkUndefined()), MkUndefined(), *(market).At(MkString("symbol")))
	_ = symbol
	return MkMap(&VarMap{
		"symbol":        symbol,
		"timestamp":     timestamp,
		"datetime":      this.Iso8601(timestamp),
		"high":          this.SafeNumber(ticker, MkString("max_price")),
		"low":           this.SafeNumber(ticker, MkString("min_price")),
		"bid":           this.SafeNumber(ticker, MkString("bid")),
		"bidVolume":     MkUndefined(),
		"ask":           this.SafeNumber(ticker, MkString("ask")),
		"askVolume":     MkUndefined(),
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
		"info":          ticker,
	})
}

func (this *Vcc) FetchTickers(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("publicGetTicker"), params)
	_ = response
	result := MkMap(&VarMap{})
	_ = result
	data := this.SafeValue(response, MkString("data"))
	_ = data
	marketIds := GoGetKeys(data)
	_ = marketIds
	for i := MkInteger(0); IsTrue(OpLw(i, marketIds.Length)); OpIncr(&i) {
		marketId := *(marketIds).At(i)
		_ = marketId
		market := this.SafeMarket(marketId, MkUndefined(), MkString("_"))
		_ = market
		symbol := *(market).At(MkString("symbol"))
		_ = symbol
		*(result).At(symbol) = this.ParseTicker(*(data).At(marketId), market)
	}
	return this.FilterByArray(result, MkString("symbol"), symbols)
}

func (this *Vcc) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.SafeInteger2(trade, MkString("trade_timestamp"), MkString("created_at"))
	_ = timestamp
	baseId := this.SafeStringUpper(trade, MkString("coin"))
	_ = baseId
	quoteId := this.SafeStringUpper(trade, MkString("currency"))
	_ = quoteId
	marketId := OpCopy(MkUndefined())
	_ = marketId
	if IsTrue(OpAnd(OpNotEq2(baseId, MkUndefined()), OpNotEq2(quoteId, MkUndefined()))) {
		marketId = OpAdd(baseId, OpAdd(MkString("_"), quoteId))
	}
	market = this.SafeMarket(marketId, market, MkString("_"))
	symbol := *(market).At(MkString("symbol"))
	_ = symbol
	priceString := this.SafeString(trade, MkString("price"))
	_ = priceString
	amountString := this.SafeString2(trade, MkString("base_volume"), MkString("quantity"))
	_ = amountString
	price := this.ParseNumber(priceString)
	_ = price
	amount := this.ParseNumber(amountString)
	_ = amount
	cost := this.SafeNumber2(trade, MkString("quote_volume"), MkString("amount"))
	_ = cost
	if IsTrue(OpEq2(cost, MkUndefined())) {
		cost = this.ParseNumber(Precise.StringMul(priceString, amountString))
	}
	side := this.SafeString2(trade, MkString("type"), MkString("trade_type"))
	_ = side
	id := this.SafeString2(trade, MkString("trade_id"), MkString("id"))
	_ = id
	feeCost := this.SafeNumber(trade, MkString("fee"))
	_ = feeCost
	fee := OpCopy(MkUndefined())
	_ = fee
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		fee = MkMap(&VarMap{
			"cost":     feeCost,
			"currency": *(market).At(MkString("quote")),
		})
	}
	return MkMap(&VarMap{
		"info":         trade,
		"id":           id,
		"order":        MkUndefined(),
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       symbol,
		"type":         MkUndefined(),
		"side":         side,
		"takerOrMaker": MkUndefined(),
		"price":        price,
		"amount":       amount,
		"cost":         cost,
		"fee":          fee,
	})
}

func (this *Vcc) FetchTrades(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"market_pair": *(market).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("count")) = Math.Min(MkInteger(1000), limit)
	}
	response := this.Call(MkString("publicGetTradesMarketPair"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"))
	_ = data
	return this.ParseTrades(data, market, since, limit)
}

func (this *Vcc) FetchTransactions(goArgs ...*Variant) *Variant {
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
	currency := OpCopy(MkUndefined())
	_ = currency
	if IsTrue(OpNotEq2(code, MkUndefined())) {
		currency = this.Currency(code)
		*(request).At(MkString("currency")) = *(currency).At(MkString("id"))
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = Math.Min(MkInteger(1000), limit)
	}
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("start")) = OpCopy(since)
	}
	response := this.Call(MkString("privateGetTransactions"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	data = this.SafeValue(data, MkString("data"), MkArray(&VarArray{}))
	return this.ParseTransactions(data, currency, since, limit)
}

func (this *Vcc) FetchDeposits(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	request := MkMap(&VarMap{"type": MkString("deposit")})
	_ = request
	return this.FetchTransactions(code, since, limit, this.Extend(request, params))
}

func (this *Vcc) FetchWithdrawals(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	request := MkMap(&VarMap{"type": MkString("withdraw")})
	_ = request
	return this.FetchTransactions(code, since, limit, this.Extend(request, params))
}

func (this *Vcc) ParseTransaction(goArgs ...*Variant) *Variant {
	transaction := GoGetArg(goArgs, 0, MkUndefined())
	_ = transaction
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	id := this.SafeString(transaction, MkString("id"))
	_ = id
	timestamp := this.SafeInteger(transaction, MkString("created_at"))
	_ = timestamp
	updated := this.SafeInteger(transaction, MkString("updated_at"))
	_ = updated
	currencyId := this.SafeString(transaction, MkString("currency"))
	_ = currencyId
	code := this.SafeCurrencyCode(currencyId, currency)
	_ = code
	status := this.ParseTransactionStatus(this.SafeString(transaction, MkString("status")))
	_ = status
	amount := this.SafeNumber(transaction, MkString("amount"))
	_ = amount
	if IsTrue(OpNotEq2(amount, MkUndefined())) {
		amount = Math.Abs(amount)
	}
	address := this.SafeString(transaction, MkString("blockchain_address"))
	_ = address
	txid := this.SafeString(transaction, MkString("transaction_id"))
	_ = txid
	tag := this.SafeString(transaction, MkString("destination_tag"))
	_ = tag
	fee := OpCopy(MkUndefined())
	_ = fee
	feeCost := this.SafeNumber(transaction, MkString("fee"))
	_ = feeCost
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		fee = MkMap(&VarMap{
			"cost":     feeCost,
			"currency": code,
		})
	}
	type_ := OpTrinary(OpGt(amount, MkInteger(0)), MkString("deposit"), MkString("withdrawal"))
	_ = type_
	return MkMap(&VarMap{
		"info":        transaction,
		"id":          id,
		"txid":        txid,
		"timestamp":   timestamp,
		"datetime":    this.Iso8601(timestamp),
		"address":     address,
		"addressTo":   address,
		"addressFrom": MkUndefined(),
		"tag":         tag,
		"tagTo":       tag,
		"tagFrom":     MkUndefined(),
		"type":        type_,
		"amount":      amount,
		"currency":    code,
		"status":      status,
		"updated":     updated,
		"fee":         fee,
	})
}

func (this *Vcc) ParseTransactionStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"pending": MkString("pending"),
		"error":   MkString("failed"),
		"success": MkString("ok"),
		"cancel":  MkString("canceled"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Vcc) ParseTransactionType(goArgs ...*Variant) *Variant {
	type_ := GoGetArg(goArgs, 0, MkUndefined())
	_ = type_
	types := MkMap(&VarMap{
		"deposit":  MkString("deposit"),
		"withdraw": MkString("withdrawal"),
	})
	_ = types
	return this.SafeString(types, type_, type_)
}

func (this *Vcc) CostToPrecision(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	cost := GoGetArg(goArgs, 1, MkUndefined())
	_ = cost
	return this.DecimalToPrecision(cost, ROUND, *(*(*(*this.At(MkString("markets"))).At(symbol)).At(MkString("precision"))).At(MkString("cost")), *this.At(MkString("precisionMode")), *this.At(MkString("paddingMode")))
}

func (this *Vcc) CreateOrder(goArgs ...*Variant) *Variant {
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
		"coin":       *(market).At(MkString("baseId")),
		"currency":   *(market).At(MkString("quoteId")),
		"trade_type": side,
		"type":       type_,
	})
	_ = request
	if IsTrue(OpEq2(type_, MkString("ceiling_market"))) {
		ceiling := this.SafeValue(params, MkString("ceiling"))
		_ = ceiling
		if IsTrue(OpNotEq2(ceiling, MkUndefined())) {
			*(request).At(MkString("ceiling")) = this.CostToPrecision(symbol, ceiling)
		} else {
			if IsTrue(OpNotEq2(price, MkUndefined())) {
				*(request).At(MkString("ceiling")) = this.CostToPrecision(symbol, OpMulti(amount, price))
			} else {
				panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" createOrder() requires a price argument or a ceiling parameter for "), OpAdd(type_, MkString(" orders"))))))
			}
		}
	} else {
		*(request).At(MkString("quantity")) = this.AmountToPrecision(symbol, amount)
	}
	if IsTrue(OpEq2(type_, MkString("limit"))) {
		*(request).At(MkString("price")) = this.PriceToPrecision(symbol, price)
	}
	stopPrice := this.SafeValue2(params, MkString("stop_price"), MkString("stopPrice"))
	_ = stopPrice
	if IsTrue(OpNotEq2(stopPrice, MkUndefined())) {
		*(request).At(MkString("is_stop")) = MkInteger(1)
		*(request).At(MkString("stop_condition")) = OpTrinary(OpEq2(side, MkString("buy")), MkString("le"), MkString("ge"))
		*(request).At(MkString("stop_price")) = this.PriceToPrecision(symbol, stopPrice)
	}
	params = this.Omit(params, MkArray(&VarArray{
		MkString("stop_price"),
		MkString("stopPrice"),
	}))
	response := this.Call(MkString("privatePostOrders"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"))
	_ = data
	return this.ParseOrder(data, market)
}

func (this *Vcc) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"order_id": id})
	_ = request
	response := this.Call(MkString("privatePutOrdersOrderIdCancel"), this.Extend(request, params))
	_ = response
	return this.ParseOrder(response)
}

func (this *Vcc) CancelAllOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	type_ := this.SafeString(params, MkString("type"))
	_ = type_
	method := OpTrinary(OpEq2(type_, MkUndefined()), MkString("privatePutOrdersCancelAll"), MkString("privatePutOrdersCancelByType"))
	_ = method
	request := MkMap(&VarMap{})
	_ = request
	if IsTrue(OpNotEq2(type_, MkUndefined())) {
		*(request).At(MkString("type")) = OpCopy(type_)
	}
	this.LoadMarkets()
	response := this.Call(method, this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	data = this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	return this.ParseOrders(data)
}

func (this *Vcc) ParseOrderStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"pending":   MkString("open"),
		"stopping":  MkString("open"),
		"executing": MkString("open"),
		"executed":  MkString("closed"),
		"canceled":  MkString("canceled"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Vcc) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	created := this.SafeValue(order, MkString("created_at"))
	_ = created
	updated := this.SafeValue(order, MkString("updated_at"))
	_ = updated
	baseId := this.SafeStringUpper(order, MkString("coin"))
	_ = baseId
	quoteId := this.SafeStringUpper(order, MkString("currency"))
	_ = quoteId
	marketId := OpAdd(baseId, OpAdd(MkString("_"), quoteId))
	_ = marketId
	market = this.SafeMarket(marketId, market, MkString("_"))
	symbol := *(market).At(MkString("symbol"))
	_ = symbol
	amount := this.SafeNumber(order, MkString("quantity"))
	_ = amount
	filled := this.SafeNumber(order, MkString("executed_quantity"))
	_ = filled
	status := this.ParseOrderStatus(this.SafeString(order, MkString("status")))
	_ = status
	cost := this.SafeNumber(order, MkString("ceiling"))
	_ = cost
	id := this.SafeString(order, MkString("id"))
	_ = id
	price := this.SafeNumber(order, MkString("price"))
	_ = price
	average := this.SafeNumber(order, MkString("executed_price"))
	_ = average
	remaining := this.SafeNumber(order, MkString("remaining"))
	_ = remaining
	type_ := this.SafeString(order, MkString("type"))
	_ = type_
	side := this.SafeString(order, MkString("trade_type"))
	_ = side
	fee := MkMap(&VarMap{
		"currency": *(market).At(MkString("quote")),
		"cost":     this.SafeNumber(order, MkString("fee")),
		"rate":     this.SafeNumber(order, MkString("fee_rate")),
	})
	_ = fee
	lastTradeTimestamp := OpCopy(MkUndefined())
	_ = lastTradeTimestamp
	if IsTrue(OpNotEq2(updated, created)) {
		lastTradeTimestamp = OpCopy(updated)
	}
	stopPrice := this.SafeNumber(order, MkString("stopPrice"))
	_ = stopPrice
	return this.SafeOrder(MkMap(&VarMap{
		"id":                 id,
		"clientOrderId":      id,
		"timestamp":          created,
		"datetime":           this.Iso8601(created),
		"lastTradeTimestamp": lastTradeTimestamp,
		"status":             status,
		"symbol":             symbol,
		"type":               type_,
		"timeInForce":        MkUndefined(),
		"postOnly":           MkUndefined(),
		"side":               side,
		"price":              price,
		"stopPrice":          stopPrice,
		"average":            average,
		"amount":             amount,
		"cost":               cost,
		"filled":             filled,
		"remaining":          remaining,
		"fee":                fee,
		"trades":             MkUndefined(),
		"info":               order,
	}))
}

func (this *Vcc) FetchOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"order_id": id})
	_ = request
	response := this.Call(MkString("privateGetOrdersOrderId"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"))
	_ = data
	return this.ParseOrder(data)
}

func (this *Vcc) FetchOrdersWithMethod(goArgs ...*Variant) *Variant {
	method := GoGetArg(goArgs, 0, MkUndefined())
	_ = method
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 2, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 3, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 4, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{})
	_ = request
	market := OpCopy(MkUndefined())
	_ = market
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
		*(request).At(MkString("coin")) = *(market).At(MkString("baseId"))
		*(request).At(MkString("currency")) = *(market).At(MkString("quoteId"))
	}
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("start_date")) = OpCopy(since)
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = Math.Min(MkInteger(1000), limit)
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	data = this.SafeValue(data, MkString("data"), MkArray(&VarArray{}))
	return this.ParseOrders(data, market, since, limit)
}

func (this *Vcc) FetchOpenOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	return this.FetchOrdersWithMethod(MkString("privateGetOrdersOpen"), symbol, since, limit, params)
}

func (this *Vcc) FetchClosedOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	return this.FetchOrdersWithMethod(MkString("privateGetOrders"), symbol, since, limit, params)
}

func (this *Vcc) FetchMyTrades(goArgs ...*Variant) *Variant {
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
	market := OpCopy(MkUndefined())
	_ = market
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
		*(request).At(MkString("coin")) = *(market).At(MkString("baseId"))
		*(request).At(MkString("currency")) = *(market).At(MkString("quoteId"))
	}
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("start_date")) = OpCopy(since)
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = Math.Min(MkInteger(1000), limit)
	}
	response := this.Call(MkString("privateGetOrdersTrades"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	data = this.SafeValue(data, MkString("data"), MkArray(&VarArray{}))
	return this.ParseTrades(data, market, since, limit)
}

func (this *Vcc) FetchDepositAddress(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{"currency": *(currency).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("privateGetDepositAddress"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"))
	_ = data
	status := this.SafeString(data, MkString("status"))
	_ = status
	if IsTrue(OpEq2(status, MkString("REQUESTED"))) {
		panic(NewAddressPending(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" is generating "), OpAdd(code, MkString(" deposit address, call fetchDepositAddress one more time later to retrieve the generated address"))))))
	}
	address := this.SafeString(data, MkString("blockchain_address"))
	_ = address
	this.CheckAddress(address)
	tag := this.SafeString(data, MkString("blockchain_tag"))
	_ = tag
	currencyId := this.SafeString(data, MkString("currency"))
	_ = currencyId
	return MkMap(&VarMap{
		"currency": this.SafeCurrencyCode(currencyId),
		"address":  address,
		"tag":      tag,
		"info":     data,
	})
}

func (this *Vcc) Sign(goArgs ...*Variant) *Variant {
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
	url := OpAdd(*this.At(MkString("version")), OpAdd(MkString("/"), this.ImplodeParams(path, params)))
	_ = url
	query := this.Omit(params, this.ExtractParams(path))
	_ = query
	if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
		OpAddAssign(&url, OpAdd(MkString("?"), this.Urlencode(query)))
	}
	if IsTrue(OpEq2(api, MkString("private"))) {
		this.CheckRequiredCredentials()
		timestamp := (this.Milliseconds()).Call(MkString("toString"))
		_ = timestamp
		if IsTrue(OpNotEq2(method, MkString("GET"))) {
			body = this.Json(query)
		}
		auth := OpAdd(method, OpAdd(MkString(" "), url))
		_ = auth
		signature := this.Hmac(this.Encode(auth), this.Encode(*this.At(MkString("secret"))), MkString("sha256"))
		_ = signature
		headers = MkMap(&VarMap{
			"Authorization": OpAdd(MkString("Bearer "), *this.At(MkString("apiKey"))),
			"Content-Type":  MkString("application/json"),
			"timestamp":     timestamp,
			"signature":     signature,
		})
	}
	url = OpAdd(*(*(*this.At(MkString("urls"))).At(MkString("api"))).At(api), OpAdd(MkString("/"), url))
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Vcc) HandleErrors(goArgs ...*Variant) *Variant {
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
	message := this.SafeString(response, MkString("message"))
	_ = message
	if IsTrue(OpNotEq2(message, MkUndefined())) {
		feedback := OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))
		_ = feedback
		this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("exact")), message, feedback)
		this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("broad")), body, feedback)
		panic(NewExchangeError(feedback))
	}
	return MkUndefined()
}
