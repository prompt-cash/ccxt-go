package ccxt

import ()

type Latoken struct {
	*ExchangeBase
}

var _ Exchange = (*Latoken)(nil)

func init() {
	exchange := &Latoken{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Latoken) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("latoken"),
		"name": MkString("Latoken"),
		"countries": MkArray(&VarArray{
			MkString("KY"),
		}),
		"version":   MkString("v1"),
		"rateLimit": MkInteger(2000),
		"certified": MkBool(false),
		"userAgent": *(*this.At(MkString("userAgents"))).At(MkString("chrome")),
		"has": MkMap(&VarMap{
			"CORS":                MkBool(false),
			"publicAPI":           MkBool(true),
			"privateAPI":          MkBool(true),
			"cancelOrder":         MkBool(true),
			"cancelAllOrders":     MkBool(true),
			"createMarketOrder":   MkBool(false),
			"createOrder":         MkBool(true),
			"fetchBalance":        MkBool(true),
			"fetchCanceledOrders": MkBool(true),
			"fetchClosedOrders":   MkBool(true),
			"fetchCurrencies":     MkBool(true),
			"fetchMyTrades":       MkBool(true),
			"fetchOpenOrders":     MkBool(true),
			"fetchOrder":          MkBool(false),
			"fetchOrdersByStatus": MkBool(true),
			"fetchOrderBook":      MkBool(true),
			"fetchTicker":         MkBool(true),
			"fetchTickers":        MkBool(true),
			"fetchTime":           MkBool(true),
			"fetchTrades":         MkBool(true),
		}),
		"urls": MkMap(&VarMap{
			"logo": MkString("https://user-images.githubusercontent.com/1294454/61511972-24c39f00-aa01-11e9-9f7c-471f1d6e5214.jpg"),
			"api":  MkString("https://api.latoken.com"),
			"www":  MkString("https://latoken.com"),
			"doc": MkArray(&VarArray{
				MkString("https://api.latoken.com"),
			}),
		}),
		"api": MkMap(&VarMap{
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("ExchangeInfo/time"),
				MkString("ExchangeInfo/limits"),
				MkString("ExchangeInfo/pairs"),
				MkString("ExchangeInfo/pairs/{currency}"),
				MkString("ExchangeInfo/pair"),
				MkString("ExchangeInfo/currencies"),
				MkString("ExchangeInfo/currencies/{symbol}"),
				MkString("MarketData/tickers"),
				MkString("MarketData/ticker/{symbol}"),
				MkString("MarketData/orderBook/{symbol}"),
				MkString("MarketData/orderBook/{symbol}/{limit}"),
				MkString("MarketData/trades/{symbol}"),
				MkString("MarketData/trades/{symbol}/{limit}"),
			})}),
			"private": MkMap(&VarMap{
				"get": MkArray(&VarArray{
					MkString("Account/balances"),
					MkString("Account/balances/{currency}"),
					MkString("Order/status"),
					MkString("Order/active"),
					MkString("Order/get_order"),
					MkString("Order/trades"),
				}),
				"post": MkArray(&VarArray{
					MkString("Order/new"),
					MkString("Order/test-order"),
					MkString("Order/cancel"),
					MkString("Order/cancel_all"),
				}),
			}),
		}),
		"fees": MkMap(&VarMap{"trading": MkMap(&VarMap{
			"feeSide":    MkString("get"),
			"tierBased":  MkBool(false),
			"percentage": MkBool(true),
			"maker":      this.ParseNumber(MkString("0.001")),
			"taker":      this.ParseNumber(MkString("0.001")),
		})}),
		"commonCurrencies": MkMap(&VarMap{
			"MT":  MkString("Monarch"),
			"TSL": MkString("Treasure SL"),
		}),
		"options": MkMap(&VarMap{"createOrderMethod": MkString("private_post_order_new")}),
		"exceptions": MkMap(&VarMap{
			"exact": MkMap(&VarMap{
				"Signature or ApiKey is not valid": AuthenticationError,
				"Request is out of time":           InvalidNonce,
				"Symbol must be specified":         BadRequest,
			}),
			"broad": MkMap(&VarMap{
				"Request limit reached":           DDoSProtection,
				"Pair":                            BadRequest,
				"Price needs to be greater than":  InvalidOrder,
				"Amount needs to be greater than": InvalidOrder,
				"The Symbol field is required":    InvalidOrder,
				"OrderType is not valid":          InvalidOrder,
				"Side is not valid":               InvalidOrder,
				"Cancelable order whit":           OrderNotFound,
				"Order":                           OrderNotFound,
			}),
		}),
	}))
}

func (this *Latoken) Nonce(goArgs ...*Variant) *Variant {
	return this.Milliseconds()
}

func (this *Latoken) FetchTime(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetExchangeInfoTime"), params)
	_ = response
	return this.SafeInteger(response, MkString("unixTimeMiliseconds"))
}

func (this *Latoken) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetExchangeInfoPairs"), params)
	_ = response
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
		market := *(response).At(i)
		_ = market
		id := this.SafeString(market, MkString("symbol"))
		_ = id
		baseId := this.SafeString(market, MkString("baseCurrency"))
		_ = baseId
		quoteId := this.SafeString(market, MkString("quotedCurrency"))
		_ = quoteId
		numericId := this.SafeInteger(market, MkString("pairId"))
		_ = numericId
		base := this.SafeCurrencyCode(baseId)
		_ = base
		quote := this.SafeCurrencyCode(quoteId)
		_ = quote
		symbol := OpAdd(base, OpAdd(MkString("/"), quote))
		_ = symbol
		pricePrecisionString := this.SafeString(market, MkString("pricePrecision"))
		_ = pricePrecisionString
		priceLimit := this.ParsePrecision(pricePrecisionString)
		_ = priceLimit
		precision := MkMap(&VarMap{
			"price":  parseInt(pricePrecisionString),
			"amount": this.SafeInteger(market, MkString("amountPrecision")),
		})
		_ = precision
		limits := MkMap(&VarMap{
			"amount": MkMap(&VarMap{
				"min": this.SafeNumber(market, MkString("minQty")),
				"max": MkUndefined(),
			}),
			"price": MkMap(&VarMap{
				"min": this.ParseNumber(priceLimit),
				"max": MkUndefined(),
			}),
			"cost": MkMap(&VarMap{
				"min": MkUndefined(),
				"max": MkUndefined(),
			}),
		})
		_ = limits
		result.Push(MkMap(&VarMap{
			"id":        id,
			"numericId": numericId,
			"info":      market,
			"symbol":    symbol,
			"base":      base,
			"quote":     quote,
			"baseId":    baseId,
			"quoteId":   quoteId,
			"active":    MkUndefined(),
			"precision": precision,
			"limits":    limits,
		}))
	}
	return result
}

func (this *Latoken) FetchCurrencies(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetExchangeInfoCurrencies"), params)
	_ = response
	result := MkMap(&VarMap{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
		currency := *(response).At(i)
		_ = currency
		id := this.SafeString(currency, MkString("symbol"))
		_ = id
		numericId := this.SafeInteger(currency, MkString("currencyId"))
		_ = numericId
		code := this.SafeCurrencyCode(id)
		_ = code
		precision := this.SafeInteger(currency, MkString("precission"))
		_ = precision
		fee := this.SafeNumber(currency, MkString("fee"))
		_ = fee
		active := OpCopy(MkUndefined())
		_ = active
		*(result).At(code) = MkMap(&VarMap{
			"id":        id,
			"numericId": numericId,
			"code":      code,
			"info":      currency,
			"name":      code,
			"active":    active,
			"fee":       fee,
			"precision": precision,
			"limits": MkMap(&VarMap{
				"amount": MkMap(&VarMap{
					"min": MkUndefined(),
					"max": MkUndefined(),
				}),
				"withdraw": MkMap(&VarMap{
					"min": MkUndefined(),
					"max": MkUndefined(),
				}),
			}),
		})
	}
	return result
}

func (this *Latoken) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privateGetAccountBalances"), params)
	_ = response
	result := MkMap(&VarMap{
		"info":      response,
		"timestamp": MkUndefined(),
		"datetime":  MkUndefined(),
	})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
		balance := *(response).At(i)
		_ = balance
		currencyId := this.SafeString(balance, MkString("symbol"))
		_ = currencyId
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		frozen := this.SafeString(balance, MkString("frozen"))
		_ = frozen
		pending := this.SafeString(balance, MkString("pending"))
		_ = pending
		account := this.Account()
		_ = account
		*(account).At(MkString("used")) = Precise.StringAdd(frozen, pending)
		*(account).At(MkString("free")) = this.SafeString(balance, MkString("available"))
		*(account).At(MkString("total")) = this.SafeString(balance, MkString("amount"))
		*(result).At(code) = OpCopy(account)
	}
	return this.ParseBalance(result)
}

func (this *Latoken) FetchOrderBook(goArgs ...*Variant) *Variant {
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
		"symbol": *(market).At(MkString("id")),
		"limit":  MkInteger(10),
	})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("publicGetMarketDataOrderBookSymbolLimit"), this.Extend(request, params))
	_ = response
	return this.ParseOrderBook(response, symbol, MkUndefined(), MkString("bids"), MkString("asks"), MkString("price"), MkString("quantity"))
}

func (this *Latoken) ParseTicker(goArgs ...*Variant) *Variant {
	ticker := GoGetArg(goArgs, 0, MkUndefined())
	_ = ticker
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	marketId := this.SafeString(ticker, MkString("symbol"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market)
	_ = symbol
	open := this.SafeNumber(ticker, MkString("open"))
	_ = open
	close := this.SafeNumber(ticker, MkString("close"))
	_ = close
	change := OpCopy(MkUndefined())
	_ = change
	if IsTrue(OpAnd(OpNotEq2(open, MkUndefined()), OpNotEq2(close, MkUndefined()))) {
		change = OpSub(close, open)
	}
	percentage := this.SafeNumber(ticker, MkString("priceChange"))
	_ = percentage
	timestamp := this.Nonce()
	_ = timestamp
	return MkMap(&VarMap{
		"symbol":        symbol,
		"timestamp":     timestamp,
		"datetime":      this.Iso8601(timestamp),
		"low":           this.SafeNumber(ticker, MkString("low")),
		"high":          this.SafeNumber(ticker, MkString("high")),
		"bid":           MkUndefined(),
		"bidVolume":     MkUndefined(),
		"ask":           MkUndefined(),
		"askVolume":     MkUndefined(),
		"vwap":          MkUndefined(),
		"open":          open,
		"close":         close,
		"last":          close,
		"previousClose": MkUndefined(),
		"change":        change,
		"percentage":    percentage,
		"average":       MkUndefined(),
		"baseVolume":    MkUndefined(),
		"quoteVolume":   this.SafeNumber(ticker, MkString("volume")),
		"info":          ticker,
	})
}

func (this *Latoken) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("publicGetMarketDataTickerSymbol"), this.Extend(request, params))
	_ = response
	return this.ParseTicker(response, market)
}

func (this *Latoken) FetchTickers(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("publicGetMarketDataTickers"), params)
	_ = response
	result := MkMap(&VarMap{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
		ticker := this.ParseTicker(*(response).At(i))
		_ = ticker
		symbol := *(ticker).At(MkString("symbol"))
		_ = symbol
		*(result).At(symbol) = OpCopy(ticker)
	}
	return this.FilterByArray(result, MkString("symbol"), symbols)
}

func (this *Latoken) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	type_ := OpCopy(MkUndefined())
	_ = type_
	timestamp := this.SafeInteger2(trade, MkString("timestamp"), MkString("time"))
	_ = timestamp
	if IsTrue(OpNotEq2(timestamp, MkUndefined())) {
		if IsTrue(OpLw(timestamp, MkInteger(1230940800000))) {
			OpMultiAssign(&timestamp, MkInteger(1000))
		}
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
	side := this.SafeString(trade, MkString("side"))
	_ = side
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(OpNotEq2(market, MkUndefined())) {
		symbol = *(market).At(MkString("symbol"))
	}
	id := this.SafeString(trade, MkString("id"))
	_ = id
	orderId := this.SafeString(trade, MkString("orderId"))
	_ = orderId
	feeCost := this.SafeNumber(trade, MkString("commission"))
	_ = feeCost
	fee := OpCopy(MkUndefined())
	_ = fee
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		fee = MkMap(&VarMap{
			"cost":     feeCost,
			"currency": MkUndefined(),
		})
	}
	return MkMap(&VarMap{
		"info":         trade,
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       symbol,
		"id":           id,
		"order":        orderId,
		"type":         type_,
		"takerOrMaker": MkUndefined(),
		"side":         side,
		"price":        price,
		"amount":       amount,
		"cost":         cost,
		"fee":          fee,
	})
}

func (this *Latoken) FetchTrades(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("publicGetMarketDataTradesSymbol"), this.Extend(request, params))
	_ = response
	trades := this.SafeValue(response, MkString("trades"), MkArray(&VarArray{}))
	_ = trades
	return this.ParseTrades(trades, market, since, limit)
}

func (this *Latoken) FetchMyTrades(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchMyTrades() requires a symbol argument"))))
	}
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("privateGetOrderTrades"), this.Extend(request, params))
	_ = response
	trades := this.SafeValue(response, MkString("trades"), MkArray(&VarArray{}))
	_ = trades
	return this.ParseTrades(trades, market, since, limit)
}

func (this *Latoken) ParseOrderStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"active":          MkString("open"),
		"partiallyFilled": MkString("open"),
		"filled":          MkString("closed"),
		"cancelled":       MkString("canceled"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Latoken) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	id := this.SafeString(order, MkString("orderId"))
	_ = id
	timestamp := this.SafeTimestamp(order, MkString("timeCreated"))
	_ = timestamp
	marketId := this.SafeString(order, MkString("symbol"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market)
	_ = symbol
	side := this.SafeString(order, MkString("side"))
	_ = side
	type_ := this.SafeString(order, MkString("orderType"))
	_ = type_
	price := this.SafeNumber(order, MkString("price"))
	_ = price
	amount := this.SafeNumber(order, MkString("amount"))
	_ = amount
	filled := this.SafeNumber(order, MkString("executedAmount"))
	_ = filled
	status := this.ParseOrderStatus(this.SafeString(order, MkString("orderStatus")))
	_ = status
	timeFilled := this.SafeTimestamp(order, MkString("timeFilled"))
	_ = timeFilled
	lastTradeTimestamp := OpCopy(MkUndefined())
	_ = lastTradeTimestamp
	if IsTrue(OpAnd(OpNotEq2(timeFilled, MkUndefined()), OpGt(timeFilled, MkInteger(0)))) {
		lastTradeTimestamp = OpCopy(timeFilled)
	}
	clientOrderId := this.SafeString(order, MkString("cliOrdId"))
	_ = clientOrderId
	return this.SafeOrder(MkMap(&VarMap{
		"id":                 id,
		"clientOrderId":      clientOrderId,
		"info":               order,
		"timestamp":          timestamp,
		"datetime":           this.Iso8601(timestamp),
		"lastTradeTimestamp": lastTradeTimestamp,
		"status":             status,
		"symbol":             symbol,
		"type":               type_,
		"timeInForce":        MkUndefined(),
		"postOnly":           MkUndefined(),
		"side":               side,
		"price":              price,
		"stopPrice":          MkUndefined(),
		"cost":               MkUndefined(),
		"amount":             amount,
		"filled":             filled,
		"average":            MkUndefined(),
		"remaining":          MkUndefined(),
		"fee":                MkUndefined(),
		"trades":             MkUndefined(),
	}))
}

func (this *Latoken) FetchOpenOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	return this.FetchOrdersWithMethod(MkString("private_get_order_active"), symbol, since, limit, params)
}

func (this *Latoken) FetchClosedOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	return this.FetchOrdersByStatus(MkString("filled"), symbol, since, limit, params)
}

func (this *Latoken) FetchCanceledOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	return this.FetchOrdersByStatus(MkString("cancelled"), symbol, since, limit, params)
}

func (this *Latoken) FetchOrdersByStatus(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"status": status})
	_ = request
	return this.FetchOrdersWithMethod(MkString("private_get_order_status"), symbol, since, limit, this.Extend(request, params))
}

func (this *Latoken) FetchOrdersWithMethod(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchOrdersWithMethod() requires a symbol argument"))))
	}
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	return this.ParseOrders(response, market, since, limit)
}

func (this *Latoken) FetchOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"orderId": id})
	_ = request
	response := this.Call(MkString("privateGetOrderGetOrder"), this.Extend(request, params))
	_ = response
	return this.ParseOrder(response)
}

func (this *Latoken) CreateOrder(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpNotEq2(type_, MkString("limit"))) {
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), MkString(" allows limit orders only"))))
	}
	request := MkMap(&VarMap{
		"symbol":    this.MarketId(symbol),
		"side":      side,
		"price":     this.PriceToPrecision(symbol, price),
		"amount":    this.AmountToPrecision(symbol, amount),
		"orderType": type_,
	})
	_ = request
	method := this.SafeString(*this.At(MkString("options")), MkString("createOrderMethod"), MkString("private_post_order_new"))
	_ = method
	response := this.Call(method, this.Extend(request, params))
	_ = response
	return this.ParseOrder(response)
}

func (this *Latoken) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"orderId": id})
	_ = request
	response := this.Call(MkString("privatePostOrderCancel"), this.Extend(request, params))
	_ = response
	return this.ParseOrder(response)
}

func (this *Latoken) CancelAllOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" cancelAllOrders() requires a symbol argument"))))
	}
	this.LoadMarkets()
	marketId := this.MarketId(symbol)
	_ = marketId
	request := MkMap(&VarMap{"symbol": marketId})
	_ = request
	response := this.Call(MkString("privatePostOrderCancelAll"), this.Extend(request, params))
	_ = response
	result := MkArray(&VarArray{})
	_ = result
	canceledOrders := this.SafeValue(response, MkString("cancelledOrders"), MkArray(&VarArray{}))
	_ = canceledOrders
	for i := MkInteger(0); IsTrue(OpLw(i, canceledOrders.Length)); OpIncr(&i) {
		order := this.ParseOrder(MkMap(&VarMap{
			"symbol":      marketId,
			"orderId":     *(canceledOrders).At(i),
			"orderStatus": MkString("canceled"),
		}))
		_ = order
		result.Push(order)
	}
	return result
}

func (this *Latoken) Sign(goArgs ...*Variant) *Variant {
	path := GoGetArg(goArgs, 0, MkUndefined())
	_ = path
	api := GoGetArg(goArgs, 1, MkString("public"))
	_ = api
	method := GoGetArg(goArgs, 2, MkString("GET"))
	_ = method
	params := GoGetArg(goArgs, 3, MkUndefined())
	_ = params
	headers := GoGetArg(goArgs, 4, MkUndefined())
	_ = headers
	body := GoGetArg(goArgs, 5, MkUndefined())
	_ = body
	request := OpAdd(MkString("/api/"), OpAdd(*this.At(MkString("version")), OpAdd(MkString("/"), this.ImplodeParams(path, params))))
	_ = request
	query := this.Omit(params, this.ExtractParams(path))
	_ = query
	if IsTrue(OpEq2(api, MkString("private"))) {
		nonce := this.Nonce()
		_ = nonce
		query = this.Extend(MkMap(&VarMap{"timestamp": nonce}), query)
	}
	urlencodedQuery := this.Urlencode(query)
	_ = urlencodedQuery
	if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
		OpAddAssign(&request, OpAdd(MkString("?"), urlencodedQuery))
	}
	if IsTrue(OpEq2(api, MkString("private"))) {
		this.CheckRequiredCredentials()
		signature := this.Hmac(this.Encode(request), this.Encode(*this.At(MkString("secret"))))
		_ = signature
		headers = MkMap(&VarMap{
			"X-LA-KEY":       *this.At(MkString("apiKey")),
			"X-LA-SIGNATURE": signature,
		})
		if IsTrue(OpEq2(method, MkString("POST"))) {
			*(headers).At(MkString("Content-Type")) = MkString("application/x-www-form-urlencoded")
			body = OpCopy(urlencodedQuery)
		}
	}
	url := OpAdd(*(*this.At(MkString("urls"))).At(MkString("api")), request)
	_ = url
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Latoken) HandleErrors(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpNot(response)) {
		return MkUndefined()
	}
	message := this.SafeString(response, MkString("message"))
	_ = message
	feedback := OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))
	_ = feedback
	if IsTrue(OpNotEq2(message, MkUndefined())) {
		this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("exact")), message, feedback)
		this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("broad")), message, feedback)
	}
	error := this.SafeValue(response, MkString("error"), MkMap(&VarMap{}))
	_ = error
	errorMessage := this.SafeString(error, MkString("message"))
	_ = errorMessage
	if IsTrue(OpNotEq2(errorMessage, MkUndefined())) {
		this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("exact")), errorMessage, feedback)
		this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("broad")), errorMessage, feedback)
		panic(NewExchangeError(feedback))
	}
	return MkUndefined()
}
