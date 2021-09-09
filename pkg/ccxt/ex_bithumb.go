package ccxt

import ()

type Bithumb struct {
	*ExchangeBase
}

var _ Exchange = (*Bithumb)(nil)

func init() {
	exchange := &Bithumb{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Bithumb) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("bithumb"),
		"name": MkString("Bithumb"),
		"countries": MkArray(&VarArray{
			MkString("KR"),
		}),
		"rateLimit": MkInteger(500),
		"has": MkMap(&VarMap{
			"cancelOrder":       MkBool(true),
			"CORS":              MkBool(true),
			"createMarketOrder": MkBool(true),
			"createOrder":       MkBool(true),
			"fetchBalance":      MkBool(true),
			"fetchMarkets":      MkBool(true),
			"fetchOHLCV":        MkBool(true),
			"fetchOpenOrders":   MkBool(true),
			"fetchOrder":        MkBool(true),
			"fetchOrderBook":    MkBool(true),
			"fetchTicker":       MkBool(true),
			"fetchTickers":      MkBool(true),
			"fetchTrades":       MkBool(true),
			"withdraw":          MkBool(true),
		}),
		"hostname": MkString("bithumb.com"),
		"urls": MkMap(&VarMap{
			"logo": MkString("https://user-images.githubusercontent.com/1294454/30597177-ea800172-9d5e-11e7-804c-b9d4fa9b56b0.jpg"),
			"api": MkMap(&VarMap{
				"public":  MkString("https://api.{hostname}/public"),
				"private": MkString("https://api.{hostname}"),
			}),
			"www":  MkString("https://www.bithumb.com"),
			"doc":  MkString("https://apidocs.bithumb.com"),
			"fees": MkString("https://en.bithumb.com/customer_support/info_fee"),
		}),
		"api": MkMap(&VarMap{
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("ticker/{currency}"),
				MkString("ticker/all"),
				MkString("ticker/ALL_BTC"),
				MkString("ticker/ALL_KRW"),
				MkString("orderbook/{currency}"),
				MkString("orderbook/all"),
				MkString("transaction_history/{currency}"),
				MkString("transaction_history/all"),
				MkString("candlestick/{currency}/{interval}"),
			})}),
			"private": MkMap(&VarMap{"post": MkArray(&VarArray{
				MkString("info/account"),
				MkString("info/balance"),
				MkString("info/wallet_address"),
				MkString("info/ticker"),
				MkString("info/orders"),
				MkString("info/user_transactions"),
				MkString("info/order_detail"),
				MkString("trade/place"),
				MkString("trade/cancel"),
				MkString("trade/btc_withdrawal"),
				MkString("trade/krw_deposit"),
				MkString("trade/krw_withdrawal"),
				MkString("trade/market_buy"),
				MkString("trade/market_sell"),
			})}),
		}),
		"fees": MkMap(&VarMap{"trading": MkMap(&VarMap{
			"maker": this.ParseNumber(MkString("0.0025")),
			"taker": this.ParseNumber(MkString("0.0025")),
		})}),
		"precisionMode": SIGNIFICANT_DIGITS,
		"exceptions": MkMap(&VarMap{
			"Bad Request(SSL)":                BadRequest,
			"Bad Request(Bad Method)":         BadRequest,
			"Bad Request.(Auth Data)":         AuthenticationError,
			"Not Member":                      AuthenticationError,
			"Invalid Apikey":                  AuthenticationError,
			"Method Not Allowed.(Access IP)":  PermissionDenied,
			"Method Not Allowed.(BTC Adress)": InvalidAddress,
			"Method Not Allowed.(Access)":     PermissionDenied,
			"Database Fail":                   ExchangeNotAvailable,
			"Invalid Parameter":               BadRequest,
			"5600":                            ExchangeError,
			"Unknown Error":                   ExchangeError,
			"After May 23th, recent_transactions is no longer, hence users will not be able to connect to recent_transactions": ExchangeError,
		}),
		"timeframes": MkMap(&VarMap{
			"1m":  MkString("1m"),
			"3m":  MkString("3m"),
			"5m":  MkString("5m"),
			"10m": MkString("10m"),
			"30m": MkString("30m"),
			"1h":  MkString("1h"),
			"6h":  MkString("6h"),
			"12h": MkString("12h"),
			"1d":  MkString("24h"),
		}),
		"options": MkMap(&VarMap{"quoteCurrencies": MkMap(&VarMap{
			"BTC": MkMap(&VarMap{"limits": MkMap(&VarMap{"cost": MkMap(&VarMap{
				"min": MkNumber(0.0002),
				"max": MkInteger(100),
			})})}),
			"KRW": MkMap(&VarMap{"limits": MkMap(&VarMap{"cost": MkMap(&VarMap{
				"min": MkInteger(500),
				"max": MkInteger(5000000000),
			})})}),
		})}),
		"commonCurrencies": MkMap(&VarMap{
			"MIR": MkString("MIR COIN"),
			"SOC": MkString("Soda Coin"),
		}),
	}))
}

func (this *Bithumb) AmountToPrecision(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	amount := GoGetArg(goArgs, 1, MkUndefined())
	_ = amount
	return this.DecimalToPrecision(amount, TRUNCATE, *(*(*(*this.At(MkString("markets"))).At(symbol)).At(MkString("precision"))).At(MkString("amount")), DECIMAL_PLACES)
}

func (this *Bithumb) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	result := MkArray(&VarArray{})
	_ = result
	quoteCurrencies := this.SafeValue(*this.At(MkString("options")), MkString("quoteCurrencies"), MkMap(&VarMap{}))
	_ = quoteCurrencies
	quotes := GoGetKeys(quoteCurrencies)
	_ = quotes
	for i := MkInteger(0); IsTrue(OpLw(i, quotes.Length)); OpIncr(&i) {
		quote := *(quotes).At(i)
		_ = quote
		extension := this.SafeValue(quoteCurrencies, quote, MkMap(&VarMap{}))
		_ = extension
		method := OpAdd(MkString("publicGetTickerALL"), quote)
		_ = method
		response := this.Call(method, params)
		_ = response
		data := this.SafeValue(response, MkString("data"))
		_ = data
		currencyIds := GoGetKeys(data)
		_ = currencyIds
		for j := MkInteger(0); IsTrue(OpLw(j, currencyIds.Length)); OpIncr(&j) {
			currencyId := *(currencyIds).At(j)
			_ = currencyId
			if IsTrue(OpEq2(currencyId, MkString("date"))) {
				continue
			}
			market := *(data).At(currencyId)
			_ = market
			base := this.SafeCurrencyCode(currencyId)
			_ = base
			symbol := OpAdd(currencyId, OpAdd(MkString("/"), quote))
			_ = symbol
			active := OpCopy(MkBool(true))
			_ = active
			if IsTrue(GoIsArray(market)) {
				numElements := OpCopy(market.Length)
				_ = numElements
				if IsTrue(OpEq2(numElements, MkInteger(0))) {
					active = OpCopy(MkBool(false))
				}
			}
			entry := this.DeepExtend(MkMap(&VarMap{
				"id":     currencyId,
				"symbol": symbol,
				"base":   base,
				"quote":  quote,
				"info":   market,
				"active": active,
				"precision": MkMap(&VarMap{
					"amount": MkInteger(4),
					"price":  MkInteger(4),
				}),
				"limits": MkMap(&VarMap{
					"amount": MkMap(&VarMap{
						"min": MkUndefined(),
						"max": MkUndefined(),
					}),
					"price": MkMap(&VarMap{
						"min": MkUndefined(),
						"max": MkUndefined(),
					}),
					"cost": MkMap(&VarMap{}),
				}),
				"baseId":  MkUndefined(),
				"quoteId": MkUndefined(),
			}), extension)
			_ = entry
			result.Push(entry)
		}
	}
	return result
}

func (this *Bithumb) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"currency": MkString("ALL")})
	_ = request
	response := this.Call(MkString("privatePostInfoBalance"), this.Extend(request, params))
	_ = response
	result := MkMap(&VarMap{"info": response})
	_ = result
	balances := this.SafeValue(response, MkString("data"))
	_ = balances
	codes := GoGetKeys(*this.At(MkString("currencies")))
	_ = codes
	for i := MkInteger(0); IsTrue(OpLw(i, codes.Length)); OpIncr(&i) {
		code := *(codes).At(i)
		_ = code
		account := this.Account()
		_ = account
		currency := this.Currency(code)
		_ = currency
		lowerCurrencyId := this.SafeStringLower(currency, MkString("id"))
		_ = lowerCurrencyId
		*(account).At(MkString("total")) = this.SafeString(balances, OpAdd(MkString("total_"), lowerCurrencyId))
		*(account).At(MkString("used")) = this.SafeString(balances, OpAdd(MkString("in_use_"), lowerCurrencyId))
		*(account).At(MkString("free")) = this.SafeString(balances, OpAdd(MkString("available_"), lowerCurrencyId))
		*(result).At(code) = OpCopy(account)
	}
	return this.ParseBalance(result)
}

func (this *Bithumb) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"currency": OpAdd(*(market).At(MkString("base")), OpAdd(MkString("_"), *(market).At(MkString("quote"))))})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("count")) = OpCopy(limit)
	}
	response := this.Call(MkString("publicGetOrderbookCurrency"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	timestamp := this.SafeInteger(data, MkString("timestamp"))
	_ = timestamp
	return this.ParseOrderBook(data, symbol, timestamp, MkString("bids"), MkString("asks"), MkString("price"), MkString("quantity"))
}

func (this *Bithumb) ParseTicker(goArgs ...*Variant) *Variant {
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
	open := this.SafeNumber(ticker, MkString("opening_price"))
	_ = open
	close := this.SafeNumber(ticker, MkString("closing_price"))
	_ = close
	change := OpCopy(MkUndefined())
	_ = change
	percentage := OpCopy(MkUndefined())
	_ = percentage
	average := OpCopy(MkUndefined())
	_ = average
	if IsTrue(OpAnd(OpNotEq2(close, MkUndefined()), OpNotEq2(open, MkUndefined()))) {
		change = OpSub(close, open)
		if IsTrue(OpGt(open, MkInteger(0))) {
			percentage = OpDiv(change, OpMulti(open, MkInteger(100)))
		}
		average = OpDiv(this.Sum(open, close), MkInteger(2))
	}
	baseVolume := this.SafeNumber(ticker, MkString("units_traded_24H"))
	_ = baseVolume
	quoteVolume := this.SafeNumber(ticker, MkString("acc_trade_value_24H"))
	_ = quoteVolume
	vwap := this.Vwap(baseVolume, quoteVolume)
	_ = vwap
	return MkMap(&VarMap{
		"symbol":        symbol,
		"timestamp":     timestamp,
		"datetime":      this.Iso8601(timestamp),
		"high":          this.SafeNumber(ticker, MkString("max_price")),
		"low":           this.SafeNumber(ticker, MkString("min_price")),
		"bid":           this.SafeNumber(ticker, MkString("buy_price")),
		"bidVolume":     MkUndefined(),
		"ask":           this.SafeNumber(ticker, MkString("sell_price")),
		"askVolume":     MkUndefined(),
		"vwap":          vwap,
		"open":          open,
		"close":         close,
		"last":          close,
		"previousClose": MkUndefined(),
		"change":        change,
		"percentage":    percentage,
		"average":       average,
		"baseVolume":    baseVolume,
		"quoteVolume":   quoteVolume,
		"info":          ticker,
	})
}

func (this *Bithumb) FetchTickers(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("publicGetTickerAll"), params)
	_ = response
	result := MkMap(&VarMap{})
	_ = result
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	timestamp := this.SafeInteger(data, MkString("date"))
	_ = timestamp
	tickers := this.Omit(data, MkString("date"))
	_ = tickers
	ids := GoGetKeys(tickers)
	_ = ids
	for i := MkInteger(0); IsTrue(OpLw(i, ids.Length)); OpIncr(&i) {
		id := *(ids).At(i)
		_ = id
		symbol := OpCopy(id)
		_ = symbol
		market := OpCopy(MkUndefined())
		_ = market
		if IsTrue(OpHasMember(id, *this.At(MkString("markets_by_id")))) {
			market = *(*this.At(MkString("markets_by_id"))).At(id)
			symbol = *(market).At(MkString("symbol"))
		}
		ticker := *(tickers).At(id)
		_ = ticker
		isArray := GoIsArray(ticker)
		_ = isArray
		if IsTrue(OpNot(isArray)) {
			*(ticker).At(MkString("date")) = OpCopy(timestamp)
			*(result).At(symbol) = this.ParseTicker(ticker, market)
		}
	}
	return this.FilterByArray(result, MkString("symbol"), symbols)
}

func (this *Bithumb) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"currency": *(market).At(MkString("base"))})
	_ = request
	response := this.Call(MkString("publicGetTickerCurrency"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	return this.ParseTicker(data, market)
}

func (this *Bithumb) ParseOHLCV(goArgs ...*Variant) *Variant {
	ohlcv := GoGetArg(goArgs, 0, MkUndefined())
	_ = ohlcv
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	return MkArray(&VarArray{
		this.SafeInteger(ohlcv, MkInteger(0)),
		this.SafeNumber(ohlcv, MkInteger(1)),
		this.SafeNumber(ohlcv, MkInteger(3)),
		this.SafeNumber(ohlcv, MkInteger(4)),
		this.SafeNumber(ohlcv, MkInteger(2)),
		this.SafeNumber(ohlcv, MkInteger(5)),
	})
}

func (this *Bithumb) FetchOHLCV(goArgs ...*Variant) *Variant {
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
		"currency": *(market).At(MkString("base")),
		"interval": *(*this.At(MkString("timeframes"))).At(timeframe),
	})
	_ = request
	response := this.Call(MkString("publicGetCandlestickCurrencyInterval"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseOHLCVs(data, market, timeframe, since, limit)
}

func (this *Bithumb) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := OpCopy(MkUndefined())
	_ = timestamp
	transactionDatetime := this.SafeString(trade, MkString("transaction_date"))
	_ = transactionDatetime
	if IsTrue(OpNotEq2(transactionDatetime, MkUndefined())) {
		parts := transactionDatetime.Split(MkString(" "))
		_ = parts
		numParts := OpCopy(parts.Length)
		_ = numParts
		if IsTrue(OpGt(numParts, MkInteger(1))) {
			transactionDate := *(parts).At(MkInteger(0))
			_ = transactionDate
			transactionTime := *(parts).At(MkInteger(1))
			_ = transactionTime
			if IsTrue(OpLw(transactionTime.Length, MkInteger(8))) {
				transactionTime = OpAdd(MkString("0"), transactionTime)
			}
			timestamp = this.Parse8601(OpAdd(transactionDate, OpAdd(MkString(" "), transactionTime)))
		} else {
			timestamp = this.SafeIntegerProduct(trade, MkString("transaction_date"), MkNumber(0.001))
		}
	}
	if IsTrue(OpNotEq2(timestamp, MkUndefined())) {
		OpSubAssign(&timestamp, OpMulti(MkInteger(9), MkInteger(3600000)))
	}
	type_ := OpCopy(MkUndefined())
	_ = type_
	side := this.SafeString(trade, MkString("type"))
	_ = side
	side = OpTrinary(OpEq2(side, MkString("ask")), MkString("sell"), MkString("buy"))
	id := this.SafeString(trade, MkString("cont_no"))
	_ = id
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(OpNotEq2(market, MkUndefined())) {
		symbol = *(market).At(MkString("symbol"))
	}
	priceString := this.SafeString(trade, MkString("price"))
	_ = priceString
	amountString := this.SafeString2(trade, MkString("units_traded"), MkString("units"))
	_ = amountString
	price := this.ParseNumber(priceString)
	_ = price
	amount := this.ParseNumber(amountString)
	_ = amount
	cost := this.SafeNumber(trade, MkString("total"))
	_ = cost
	if IsTrue(OpEq2(cost, MkUndefined())) {
		cost = this.ParseNumber(Precise.StringMul(priceString, amountString))
	}
	fee := OpCopy(MkUndefined())
	_ = fee
	feeCost := this.SafeNumber(trade, MkString("fee"))
	_ = feeCost
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		feeCurrencyId := this.SafeString(trade, MkString("fee_currency"))
		_ = feeCurrencyId
		feeCurrencyCode := this.CommonCurrencyCode(feeCurrencyId)
		_ = feeCurrencyCode
		fee = MkMap(&VarMap{
			"cost":     feeCost,
			"currency": feeCurrencyCode,
		})
	}
	return MkMap(&VarMap{
		"id":           id,
		"info":         trade,
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       symbol,
		"order":        MkUndefined(),
		"type":         type_,
		"side":         side,
		"takerOrMaker": MkUndefined(),
		"price":        price,
		"amount":       amount,
		"cost":         cost,
		"fee":          fee,
	})
}

func (this *Bithumb) FetchTrades(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"currency": *(market).At(MkString("base"))})
	_ = request
	if IsTrue(OpEq2(limit, MkUndefined())) {
		*(request).At(MkString("count")) = OpCopy(limit)
	}
	response := this.Call(MkString("publicGetTransactionHistoryCurrency"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseTrades(data, market, since, limit)
}

func (this *Bithumb) CreateOrder(goArgs ...*Variant) *Variant {
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
		"order_currency":   *(market).At(MkString("id")),
		"payment_currency": *(market).At(MkString("quote")),
		"units":            amount,
	})
	_ = request
	method := MkString("privatePostTradePlace")
	_ = method
	if IsTrue(OpEq2(type_, MkString("limit"))) {
		*(request).At(MkString("price")) = OpCopy(price)
		*(request).At(MkString("type")) = OpTrinary(OpEq2(side, MkString("buy")), MkString("bid"), MkString("ask"))
	} else {
		method = OpAdd(MkString("privatePostTradeMarket"), this.Capitalize(side))
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	id := this.SafeString(response, MkString("order_id"))
	_ = id
	if IsTrue(OpEq2(id, MkUndefined())) {
		panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")), MkString(" createOrder() did not return an order id"))))
	}
	return MkMap(&VarMap{
		"info":   response,
		"symbol": symbol,
		"type":   type_,
		"side":   side,
		"id":     id,
	})
}

func (this *Bithumb) FetchOrder(goArgs ...*Variant) *Variant {
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
		"order_id":         id,
		"count":            MkInteger(1),
		"order_currency":   *(market).At(MkString("base")),
		"payment_currency": *(market).At(MkString("quote")),
	})
	_ = request
	response := this.Call(MkString("privatePostInfoOrderDetail"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"))
	_ = data
	return this.ParseOrder(this.Extend(data, MkMap(&VarMap{"order_id": id})), market)
}

func (this *Bithumb) ParseOrderStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"Pending":   MkString("open"),
		"Completed": MkString("closed"),
		"Cancel":    MkString("canceled"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Bithumb) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.SafeIntegerProduct(order, MkString("order_date"), MkNumber(0.001))
	_ = timestamp
	sideProperty := this.SafeValue2(order, MkString("type"), MkString("side"))
	_ = sideProperty
	side := OpTrinary(OpEq2(sideProperty, MkString("bid")), MkString("buy"), MkString("sell"))
	_ = side
	status := this.ParseOrderStatus(this.SafeString(order, MkString("order_status")))
	_ = status
	price := this.SafeNumber2(order, MkString("order_price"), MkString("price"))
	_ = price
	type_ := MkString("limit")
	_ = type_
	if IsTrue(OpEq2(price, MkInteger(0))) {
		price = OpCopy(MkUndefined())
		type_ = MkString("market")
	}
	amount := this.SafeNumber2(order, MkString("order_qty"), MkString("units"))
	_ = amount
	remaining := this.SafeNumber(order, MkString("units_remaining"))
	_ = remaining
	if IsTrue(OpEq2(remaining, MkUndefined())) {
		if IsTrue(OpEq2(status, MkString("closed"))) {
			remaining = MkInteger(0)
		} else {
			if IsTrue(OpNotEq2(status, MkString("canceled"))) {
				remaining = OpCopy(amount)
			}
		}
	}
	symbol := OpCopy(MkUndefined())
	_ = symbol
	baseId := this.SafeString(order, MkString("order_currency"))
	_ = baseId
	quoteId := this.SafeString(order, MkString("payment_currency"))
	_ = quoteId
	base := this.SafeCurrencyCode(baseId)
	_ = base
	quote := this.SafeCurrencyCode(quoteId)
	_ = quote
	if IsTrue(OpAnd(OpNotEq2(base, MkUndefined()), OpNotEq2(quote, MkUndefined()))) {
		symbol = OpAdd(base, OpAdd(MkString("/"), quote))
	}
	if IsTrue(OpAnd(OpEq2(symbol, MkUndefined()), OpNotEq2(market, MkUndefined()))) {
		symbol = *(market).At(MkString("symbol"))
	}
	id := this.SafeString(order, MkString("order_id"))
	_ = id
	rawTrades := this.SafeValue(order, MkString("contract"), MkArray(&VarArray{}))
	_ = rawTrades
	trades := this.ParseTrades(rawTrades, market, MkUndefined(), MkUndefined(), MkMap(&VarMap{
		"side":   side,
		"symbol": symbol,
		"order":  id,
	}))
	_ = trades
	return this.SafeOrder(MkMap(&VarMap{
		"info":               order,
		"id":                 id,
		"clientOrderId":      MkUndefined(),
		"timestamp":          timestamp,
		"datetime":           this.Iso8601(timestamp),
		"lastTradeTimestamp": MkUndefined(),
		"symbol":             symbol,
		"type":               type_,
		"timeInForce":        MkUndefined(),
		"postOnly":           MkUndefined(),
		"side":               side,
		"price":              price,
		"stopPrice":          MkUndefined(),
		"amount":             amount,
		"cost":               MkUndefined(),
		"average":            MkUndefined(),
		"filled":             MkUndefined(),
		"remaining":          remaining,
		"status":             status,
		"fee":                MkUndefined(),
		"trades":             trades,
	}))
}

func (this *Bithumb) FetchOpenOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchOpenOrders() requires a symbol argument"))))
	}
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	if IsTrue(OpEq2(limit, MkUndefined())) {
		limit = MkInteger(100)
	}
	request := MkMap(&VarMap{
		"count":            limit,
		"order_currency":   *(market).At(MkString("base")),
		"payment_currency": *(market).At(MkString("quote")),
	})
	_ = request
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("after")) = OpCopy(since)
	}
	response := this.Call(MkString("privatePostInfoOrders"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseOrders(data, market, since, limit)
}

func (this *Bithumb) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	side_in_params := OpHasMember(MkString("side"), params)
	_ = side_in_params
	if IsTrue(OpNot(side_in_params)) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" cancelOrder() requires a `side` parameter (sell or buy)"))))
	}
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" cancelOrder() requires a `symbol` argument"))))
	}
	market := this.Market(symbol)
	_ = market
	side := OpTrinary(OpEq2(*(params).At(MkString("side")), MkString("buy")), MkString("bid"), MkString("ask"))
	_ = side
	params = this.Omit(params, MkArray(&VarArray{
		MkString("side"),
		MkString("currency"),
	}))
	request := MkMap(&VarMap{
		"order_id":         id,
		"type":             side,
		"order_currency":   *(market).At(MkString("base")),
		"payment_currency": *(market).At(MkString("quote")),
	})
	_ = request
	return this.Call(MkString("privatePostTradeCancel"), this.Extend(request, params))
}

func (this *Bithumb) CancelUnifiedOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	request := MkMap(&VarMap{"side": *(order).At(MkString("side"))})
	_ = request
	return this.CancelOrder(*(order).At(MkString("id")), *(order).At(MkString("symbol")), this.Extend(request, params))
}

func (this *Bithumb) Withdraw(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{
		"units":    amount,
		"address":  address,
		"currency": *(currency).At(MkString("id")),
	})
	_ = request
	if IsTrue(OpOr(OpEq2(currency, MkString("XRP")), OpOr(OpEq2(currency, MkString("XMR")), OpOr(OpEq2(currency, MkString("EOS")), OpEq2(currency, MkString("STEEM")))))) {
		destination := this.SafeString(params, MkString("destination"))
		_ = destination
		if IsTrue(OpAnd(OpEq2(tag, MkUndefined()), OpEq2(destination, MkUndefined()))) {
			panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), OpAdd(code, MkString(" withdraw() requires a tag argument or an extra destination param"))))))
		} else {
			if IsTrue(OpNotEq2(tag, MkUndefined())) {
				*(request).At(MkString("destination")) = OpCopy(tag)
			}
		}
	}
	response := this.Call(MkString("privatePostTradeBtcWithdrawal"), this.Extend(request, params))
	_ = response
	return MkMap(&VarMap{
		"info": response,
		"id":   MkUndefined(),
	})
}

func (this *Bithumb) Nonce(goArgs ...*Variant) *Variant {
	return this.Milliseconds()
}

func (this *Bithumb) Sign(goArgs ...*Variant) *Variant {
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
		this.CheckRequiredCredentials()
		body = this.Urlencode(this.Extend(MkMap(&VarMap{"endpoint": endpoint}), query))
		nonce := (this.Nonce()).Call(MkString("toString"))
		_ = nonce
		auth := OpAdd(endpoint, OpAdd(MkString("\000"), OpAdd(body, OpAdd(MkString("\000"), nonce))))
		_ = auth
		signature := this.Hmac(this.Encode(auth), this.Encode(*this.At(MkString("secret"))), MkString("sha512"))
		_ = signature
		signature64 := this.Decode(this.StringToBase64(signature))
		_ = signature64
		headers = MkMap(&VarMap{
			"Accept":       MkString("application/json"),
			"Content-Type": MkString("application/x-www-form-urlencoded"),
			"Api-Key":      *this.At(MkString("apiKey")),
			"Api-Sign":     signature64,
			"Api-Nonce":    nonce,
		})
	}
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Bithumb) HandleErrors(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpHasMember(MkString("status"), response)) {
		status := this.SafeString(response, MkString("status"))
		_ = status
		message := this.SafeString(response, MkString("message"))
		_ = message
		if IsTrue(OpNotEq2(status, MkUndefined())) {
			if IsTrue(OpEq2(status, MkString("0000"))) {
				return MkUndefined()
			} else {
				if IsTrue(OpEq2(message, MkString("ê±°ë ì§íì¤ì¸ ë´ì­ì´ ì¡´ì¬íì§ ììµëë¤"))) {
					return MkUndefined()
				}
			}
			feedback := OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))
			_ = feedback
			this.ThrowExactlyMatchedException(*this.At(MkString("exceptions")), status, feedback)
			this.ThrowExactlyMatchedException(*this.At(MkString("exceptions")), message, feedback)
			panic(NewExchangeError(feedback))
		}
	}
	return MkUndefined()
}
