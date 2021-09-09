package ccxt

import ()

type Btcturk struct {
	*ExchangeBase
}

var _ Exchange = (*Btcturk)(nil)

func init() {
	exchange := &Btcturk{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Btcturk) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("btcturk"),
		"name": MkString("BTCTurk"),
		"countries": MkArray(&VarArray{
			MkString("TR"),
		}),
		"rateLimit": MkInteger(1000),
		"has": MkMap(&VarMap{
			"cancelOrder":     MkBool(true),
			"CORS":            MkBool(true),
			"createOrder":     MkBool(true),
			"fetchBalance":    MkBool(true),
			"fetchMarkets":    MkBool(true),
			"fetchOHLCV":      MkBool(true),
			"fetchOrderBook":  MkBool(true),
			"fetchOpenOrders": MkBool(true),
			"fetchOrders":     MkBool(true),
			"fetchTicker":     MkBool(true),
			"fetchTickers":    MkBool(true),
			"fetchTrades":     MkBool(true),
			"fetchMyTrades":   MkBool(true),
		}),
		"timeframes": MkMap(&VarMap{"1d": MkString("1d")}),
		"urls": MkMap(&VarMap{
			"logo": MkString("https://user-images.githubusercontent.com/51840849/87153926-efbef500-c2c0-11ea-9842-05b63612c4b9.jpg"),
			"api": MkMap(&VarMap{
				"public":  MkString("https://api.btcturk.com/api/v2"),
				"private": MkString("https://api.btcturk.com/api/v1"),
				"graph":   MkString("https://graph-api.btcturk.com/v1"),
			}),
			"www": MkString("https://www.btcturk.com"),
			"doc": MkString("https://github.com/BTCTrader/broker-api-docs"),
		}),
		"api": MkMap(&VarMap{
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("orderbook"),
				MkString("ticker"),
				MkString("trades"),
				MkString("server/exchangeinfo"),
			})}),
			"private": MkMap(&VarMap{
				"get": MkArray(&VarArray{
					MkString("users/balances"),
					MkString("openOrders"),
					MkString("allOrders"),
					MkString("users/transactions/trade"),
				}),
				"post": MkArray(&VarArray{
					MkString("order"),
					MkString("cancelOrder"),
				}),
				"delete": MkArray(&VarArray{
					MkString("order"),
				}),
			}),
			"graph": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("ohlcs"),
			})}),
		}),
		"fees": MkMap(&VarMap{"trading": MkMap(&VarMap{
			"maker": this.ParseNumber(MkString("0.0005")),
			"taker": this.ParseNumber(MkString("0.0009")),
		})}),
		"exceptions": MkMap(&VarMap{"exact": MkMap(&VarMap{
			"FAILED_ORDER_WITH_OPEN_ORDERS": InsufficientFunds,
			"FAILED_LIMIT_ORDER":            InvalidOrder,
			"FAILED_MARKET_ORDER":           InvalidOrder,
		})}),
	}))
}

func (this *Btcturk) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetServerExchangeinfo"), params)
	_ = response
	data := this.SafeValue(response, MkString("data"))
	_ = data
	markets := this.SafeValue(data, MkString("symbols"), MkArray(&VarArray{}))
	_ = markets
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, markets.Length)); OpIncr(&i) {
		entry := *(markets).At(i)
		_ = entry
		id := this.SafeString(entry, MkString("name"))
		_ = id
		baseId := this.SafeString(entry, MkString("numerator"))
		_ = baseId
		quoteId := this.SafeString(entry, MkString("denominator"))
		_ = quoteId
		base := this.SafeCurrencyCode(baseId)
		_ = base
		quote := this.SafeCurrencyCode(quoteId)
		_ = quote
		symbol := OpAdd(base, OpAdd(MkString("/"), quote))
		_ = symbol
		filters := this.SafeValue(entry, MkString("filters"))
		_ = filters
		minPrice := OpCopy(MkUndefined())
		_ = minPrice
		maxPrice := OpCopy(MkUndefined())
		_ = maxPrice
		minAmount := OpCopy(MkUndefined())
		_ = minAmount
		maxAmount := OpCopy(MkUndefined())
		_ = maxAmount
		minCost := OpCopy(MkUndefined())
		_ = minCost
		for j := MkInteger(0); IsTrue(OpLw(j, filters.Length)); OpIncr(&j) {
			filter := *(filters).At(j)
			_ = filter
			filterType := this.SafeString(filter, MkString("filterType"))
			_ = filterType
			if IsTrue(OpEq2(filterType, MkString("PRICE_FILTER"))) {
				minPrice = this.SafeNumber(filter, MkString("minPrice"))
				maxPrice = this.SafeNumber(filter, MkString("maxPrice"))
				minAmount = this.SafeNumber(filter, MkString("minAmount"))
				maxAmount = this.SafeNumber(filter, MkString("maxAmount"))
				minCost = this.SafeNumber(filter, MkString("minExchangeValue"))
			}
		}
		status := this.SafeString(entry, MkString("status"))
		_ = status
		active := OpEq2(status, MkString("TRADING"))
		_ = active
		limits := MkMap(&VarMap{
			"price": MkMap(&VarMap{
				"min": minPrice,
				"max": maxPrice,
			}),
			"amount": MkMap(&VarMap{
				"min": minAmount,
				"max": maxAmount,
			}),
			"cost": MkMap(&VarMap{
				"min": minCost,
				"max": MkUndefined(),
			}),
		})
		_ = limits
		precision := MkMap(&VarMap{
			"price":  this.SafeInteger(entry, MkString("denominatorScale")),
			"amount": this.SafeInteger(entry, MkString("numeratorScale")),
		})
		_ = precision
		result.Push(MkMap(&VarMap{
			"info":      entry,
			"symbol":    symbol,
			"id":        id,
			"base":      base,
			"quote":     quote,
			"baseId":    baseId,
			"quoteId":   quoteId,
			"limits":    limits,
			"precision": precision,
			"active":    active,
		}))
	}
	return result
}

func (this *Btcturk) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privateGetUsersBalances"), params)
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	result := MkMap(&VarMap{
		"info":      response,
		"timestamp": MkUndefined(),
		"datetime":  MkUndefined(),
	})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, data.Length)); OpIncr(&i) {
		entry := *(data).At(i)
		_ = entry
		currencyId := this.SafeString(entry, MkString("asset"))
		_ = currencyId
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		account := this.Account()
		_ = account
		*(account).At(MkString("total")) = this.SafeString(entry, MkString("balance"))
		*(account).At(MkString("free")) = this.SafeString(entry, MkString("free"))
		*(account).At(MkString("used")) = this.SafeString(entry, MkString("locked"))
		*(result).At(code) = OpCopy(account)
	}
	return this.ParseBalance(result)
}

func (this *Btcturk) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"pairSymbol": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("publicGetOrderbook"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"))
	_ = data
	timestamp := this.SafeInteger(data, MkString("timestamp"))
	_ = timestamp
	return this.ParseOrderBook(data, symbol, timestamp, MkString("bids"), MkString("asks"), MkInteger(0), MkInteger(1))
}

func (this *Btcturk) ParseTicker(goArgs ...*Variant) *Variant {
	ticker := GoGetArg(goArgs, 0, MkUndefined())
	_ = ticker
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	marketId := this.SafeString(ticker, MkString("pair"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market)
	_ = symbol
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
		"bid":           this.SafeNumber(ticker, MkString("bid")),
		"bidVolume":     MkUndefined(),
		"ask":           this.SafeNumber(ticker, MkString("ask")),
		"askVolume":     MkUndefined(),
		"vwap":          MkUndefined(),
		"open":          this.SafeNumber(ticker, MkString("open")),
		"close":         last,
		"last":          last,
		"previousClose": MkUndefined(),
		"change":        this.SafeNumber(ticker, MkString("daily")),
		"percentage":    this.SafeNumber(ticker, MkString("dailyPercent")),
		"average":       this.SafeNumber(ticker, MkString("average")),
		"baseVolume":    this.SafeNumber(ticker, MkString("volume")),
		"quoteVolume":   MkUndefined(),
		"info":          ticker,
	})
}

func (this *Btcturk) FetchTickers(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("publicGetTicker"), params)
	_ = response
	tickers := this.SafeValue(response, MkString("data"))
	_ = tickers
	return this.ParseTickers(tickers, symbols)
}

func (this *Btcturk) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	tickers := this.FetchTickers(MkArray(&VarArray{
		symbol,
	}), params)
	_ = tickers
	return this.SafeValue(tickers, symbol)
}

func (this *Btcturk) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.SafeInteger2(trade, MkString("date"), MkString("timestamp"))
	_ = timestamp
	id := this.SafeString2(trade, MkString("tid"), MkString("id"))
	_ = id
	order := this.SafeString(trade, MkString("orderId"))
	_ = order
	priceString := this.SafeString(trade, MkString("price"))
	_ = priceString
	amountString := Precise.StringAbs(this.SafeString(trade, MkString("amount")))
	_ = amountString
	price := this.ParseNumber(priceString)
	_ = price
	amount := this.ParseNumber(amountString)
	_ = amount
	cost := this.ParseNumber(Precise.StringMul(priceString, amountString))
	_ = cost
	marketId := this.SafeString(trade, MkString("pair"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market)
	_ = symbol
	side := this.SafeString2(trade, MkString("side"), MkString("orderType"))
	_ = side
	fee := OpCopy(MkUndefined())
	_ = fee
	feeAmountString := this.SafeString(trade, MkString("fee"))
	_ = feeAmountString
	if IsTrue(OpNotEq2(feeAmountString, MkUndefined())) {
		feeCurrency := this.SafeString(trade, MkString("denominatorSymbol"))
		_ = feeCurrency
		fee = MkMap(&VarMap{
			"cost":     this.ParseNumber(Precise.StringAbs(feeAmountString)),
			"currency": this.SafeCurrencyCode(feeCurrency),
		})
	}
	return MkMap(&VarMap{
		"info":         trade,
		"id":           id,
		"order":        order,
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

func (this *Btcturk) FetchTrades(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"pairSymbol": *(market).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("last")) = OpCopy(limit)
	}
	response := this.Call(MkString("publicGetTrades"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"))
	_ = data
	return this.ParseTrades(data, market, since, limit)
}

func (this *Btcturk) ParseOHLCV(goArgs ...*Variant) *Variant {
	ohlcv := GoGetArg(goArgs, 0, MkUndefined())
	_ = ohlcv
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	return MkArray(&VarArray{
		this.SafeTimestamp(ohlcv, MkString("time")),
		this.SafeNumber(ohlcv, MkString("open")),
		this.SafeNumber(ohlcv, MkString("high")),
		this.SafeNumber(ohlcv, MkString("low")),
		this.SafeNumber(ohlcv, MkString("close")),
		this.SafeNumber(ohlcv, MkString("volume")),
	})
}

func (this *Btcturk) FetchOHLCV(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	timeframe := GoGetArg(goArgs, 1, MkString("1d"))
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
	request := MkMap(&VarMap{"pair": *(market).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("last")) = OpCopy(limit)
	}
	response := this.Call(MkString("graphGetOhlcs"), this.Extend(request, params))
	_ = response
	return this.ParseOHLCVs(response, market, timeframe, since, limit)
}

func (this *Btcturk) CreateOrder(goArgs ...*Variant) *Variant {
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
		"orderType":   side,
		"orderMethod": type_,
		"pairSymbol":  *(market).At(MkString("id")),
		"quantity":    this.AmountToPrecision(symbol, amount),
	})
	_ = request
	if IsTrue(OpNotEq2(type_, MkString("market"))) {
		*(request).At(MkString("price")) = this.PriceToPrecision(symbol, price)
	}
	if IsTrue(OpHasMember(MkString("clientOrderId"), params)) {
		*(request).At(MkString("newClientOrderId")) = *(params).At(MkString("clientOrderId"))
	} else {
		if IsTrue(OpNot(OpHasMember(MkString("newClientOrderId"), params))) {
			*(request).At(MkString("newClientOrderId")) = this.Uuid()
		}
	}
	response := this.Call(MkString("privatePostOrder"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"))
	_ = data
	return this.ParseOrder(data, market)
}

func (this *Btcturk) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	request := MkMap(&VarMap{"id": id})
	_ = request
	return this.Call(MkString("privateDeleteOrder"), this.Extend(request, params))
}

func (this *Btcturk) FetchOpenOrders(goArgs ...*Variant) *Variant {
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
		*(request).At(MkString("pairSymbol")) = *(market).At(MkString("id"))
	}
	response := this.Call(MkString("privateGetOpenOrders"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"))
	_ = data
	bids := this.SafeValue(data, MkString("bids"), MkArray(&VarArray{}))
	_ = bids
	asks := this.SafeValue(data, MkString("asks"), MkArray(&VarArray{}))
	_ = asks
	return this.ParseOrders(this.ArrayConcat(bids, asks), market, since, limit)
}

func (this *Btcturk) FetchOrders(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"pairSymbol": *(market).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("last")) = OpCopy(limit)
	}
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("startTime")) = Math.Floor(OpDiv(since, MkInteger(1000)))
	}
	response := this.Call(MkString("privateGetAllOrders"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"))
	_ = data
	return this.ParseOrders(data, market, since, limit)
}

func (this *Btcturk) ParseOrderStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"Untouched": MkString("open"),
		"Partial":   MkString("open"),
		"Canceled":  MkString("canceled"),
		"Closed":    MkString("closed"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Btcturk) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	id := this.SafeString(order, MkString("id"))
	_ = id
	priceString := this.SafeString(order, MkString("price"))
	_ = priceString
	precisePrice := NewPrecise(priceString)
	_ = precisePrice
	price := OpCopy(MkUndefined())
	_ = price
	isZero := OpEq2(precisePrice.ToString(), MkString("0"))
	_ = isZero
	if IsTrue(OpNot(isZero)) {
		price = this.ParseNumber(precisePrice.ToString())
	}
	amountString := this.SafeString(order, MkString("quantity"))
	_ = amountString
	amount := this.ParseNumber(Precise.StringAbs(amountString))
	_ = amount
	remaining := this.SafeNumber(order, MkString("leftAmount"))
	_ = remaining
	marketId := this.SafeNumber(order, MkString("pairSymbol"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market)
	_ = symbol
	side := this.SafeString(order, MkString("type"))
	_ = side
	type_ := this.SafeString(order, MkString("method"))
	_ = type_
	clientOrderId := this.SafeString(order, MkString("orderClientId"))
	_ = clientOrderId
	timestamp := this.SafeInteger2(order, MkString("updateTime"), MkString("datetime"))
	_ = timestamp
	rawStatus := this.SafeString(order, MkString("status"))
	_ = rawStatus
	status := this.ParseOrderStatus(rawStatus)
	_ = status
	return this.SafeOrder(MkMap(&VarMap{
		"info":          order,
		"id":            id,
		"price":         price,
		"amount":        amount,
		"remaining":     remaining,
		"filled":        MkUndefined(),
		"cost":          MkUndefined(),
		"average":       MkUndefined(),
		"status":        status,
		"side":          side,
		"type":          type_,
		"clientOrderId": clientOrderId,
		"timestamp":     timestamp,
		"datetime":      this.Iso8601(timestamp),
		"symbol":        symbol,
		"fee":           MkUndefined(),
	}))
}

func (this *Btcturk) FetchMyTrades(goArgs ...*Variant) *Variant {
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
	response := this.Call(MkString("privateGetUsersTransactionsTrade"))
	_ = response
	data := this.SafeValue(response, MkString("data"))
	_ = data
	return this.ParseTrades(data, market, since, limit)
}

func (this *Btcturk) Nonce(goArgs ...*Variant) *Variant {
	return this.Milliseconds()
}

func (this *Btcturk) Sign(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpEq2(*this.At(MkString("id")), MkString("btctrader"))) {
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), MkString(" is an abstract base API for BTCExchange, BTCTurk"))))
	}
	url := OpAdd(*(*(*this.At(MkString("urls"))).At(MkString("api"))).At(api), OpAdd(MkString("/"), path))
	_ = url
	if IsTrue(OpOr(OpEq2(method, MkString("GET")), OpEq2(method, MkString("DELETE")))) {
		if IsTrue(*(GoGetKeys(params)).At(MkString("length"))) {
			OpAddAssign(&url, OpAdd(MkString("?"), this.Urlencode(params)))
		}
	} else {
		body = this.Json(params)
	}
	if IsTrue(OpEq2(api, MkString("private"))) {
		this.CheckRequiredCredentials()
		nonce := (this.Nonce()).Call(MkString("toString"))
		_ = nonce
		secret := this.Base64ToBinary(*this.At(MkString("secret")))
		_ = secret
		auth := OpAdd(*this.At(MkString("apiKey")), nonce)
		_ = auth
		headers = MkMap(&VarMap{
			"X-PCK":        *this.At(MkString("apiKey")),
			"X-Stamp":      nonce,
			"X-Signature":  this.Hmac(this.Encode(auth), secret, MkString("sha256"), MkString("base64")),
			"Content-Type": MkString("application/json"),
		})
	}
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Btcturk) HandleErrors(goArgs ...*Variant) *Variant {
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
	errorCode := this.SafeString(response, MkString("code"), MkString("0"))
	_ = errorCode
	message := this.SafeString(response, MkString("message"))
	_ = message
	output := OpTrinary(OpEq2(message, MkUndefined()), body, message)
	_ = output
	this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("exact")), message, OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), output)))
	if IsTrue(OpNotEq2(errorCode, MkString("0"))) {
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), output))))
	}
	return MkUndefined()
}
