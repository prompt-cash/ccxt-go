package ccxt

import ()

type Bitbns struct {
	*ExchangeBase
}

var _ Exchange = (*Bitbns)(nil)

func init() {
	exchange := &Bitbns{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Bitbns) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("bitbns"),
		"name": MkString("Bitbns"),
		"countries": MkArray(&VarArray{
			MkString("IN"),
		}),
		"rateLimit": MkInteger(1000),
		"certified": MkBool(false),
		"pro":       MkBool(false),
		"has": MkMap(&VarMap{
			"cancelOrder":         MkBool(true),
			"createOrder":         MkBool(true),
			"fetchBalance":        MkBool(true),
			"fetchDepositAddress": MkBool(true),
			"fetchDeposits":       MkBool(true),
			"fetchMarkets":        MkBool(true),
			"fetchMyTrades":       MkBool(true),
			"fetchOHLCV":          MkBool(false),
			"fetchOpenOrders":     MkBool(true),
			"fetchOrder":          MkBool(true),
			"fetchOrderBook":      MkBool(true),
			"fetchStatus":         MkBool(true),
			"fetchTicker":         MkString("emulated"),
			"fetchTickers":        MkBool(true),
			"fetchTrades":         MkBool(false),
			"fetchWithdrawals":    MkBool(true),
		}),
		"timeframes": MkMap(&VarMap{}),
		"urls": MkMap(&VarMap{
			"logo": MkString("https://user-images.githubusercontent.com/1294454/117201933-e7a6e780-adf5-11eb-9d80-98fc2a21c3d6.jpg"),
			"api": MkMap(&VarMap{
				"ccxt": MkString("https://bitbns.com/order"),
				"v1":   MkString("https://api.bitbns.com/api/trade/v1"),
				"v2":   MkString("https://api.bitbns.com/api/trade/v2"),
			}),
			"www":      MkString("https://bitbns.com"),
			"referral": MkString("https://ref.bitbns.com/1090961"),
			"doc": MkArray(&VarArray{
				MkString("https://bitbns.com/trade/#/api-trading/"),
			}),
			"fees": MkString("https://bitbns.com/fees"),
		}),
		"api": MkMap(&VarMap{
			"ccxt": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("fetchMarkets"),
				MkString("fetchTickers"),
				MkString("fetchOrderbook"),
			})}),
			"v1": MkMap(&VarMap{
				"get": MkArray(&VarArray{
					MkString("platform/status"),
					MkString("tickers"),
					MkString("orderbook/sell/{symbol}"),
					MkString("orderbook/buy/{symbol}"),
				}),
				"post": MkArray(&VarArray{
					MkString("currentCoinBalance/EVERYTHING"),
					MkString("getApiUsageStatus/USAGE"),
					MkString("getOrderSocketToken/USAGE"),
					MkString("currentCoinBalance/{symbol}"),
					MkString("orderStatus/{symbol}"),
					MkString("depositHistory/{symbol}"),
					MkString("withdrawHistory/{symbol}"),
					MkString("listOpenOrders/{symbol}"),
					MkString("listOpenStopOrders/{symbol}"),
					MkString("getCoinAddress/{symbol}"),
					MkString("placeSellOrder/{symbol}"),
					MkString("placeBuyOrder/{symbol}"),
					MkString("buyStopLoss/{symbol}"),
					MkString("placeSellOrder/{symbol}"),
					MkString("cancelOrder/{symbol}"),
					MkString("cancelStopLossOrder/{symbol}"),
					MkString("listExecutedOrders/{symbol}"),
				}),
			}),
			"v2": MkMap(&VarMap{"post": MkArray(&VarArray{
				MkString("orders"),
				MkString("cancel"),
				MkString("getordersnew"),
				MkString("marginOrders"),
			})}),
		}),
		"fees": MkMap(&VarMap{"trading": MkMap(&VarMap{
			"feeSide":    MkString("quote"),
			"tierBased":  MkBool(false),
			"percentage": MkBool(true),
			"taker":      this.ParseNumber(MkString("0.0025")),
			"maker":      this.ParseNumber(MkString("0.0025")),
		})}),
		"exceptions": MkMap(&VarMap{
			"exact": MkMap(&VarMap{
				"400": BadRequest,
				"409": BadSymbol,
				"416": InsufficientFunds,
				"417": OrderNotFound,
			}),
			"broad": MkMap(&VarMap{}),
		}),
	}))
}

func (this *Bitbns) FetchStatus(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("v1GetPlatformStatus"), params)
	_ = response
	status := this.SafeString(response, MkString("status"))
	_ = status
	if IsTrue(OpNotEq2(status, MkUndefined())) {
		status = OpTrinary(OpEq2(status, MkString("1")), MkString("ok"), MkString("maintenance"))
		*this.At(MkString("status")) = this.Extend(*this.At(MkString("status")), MkMap(&VarMap{
			"status":  status,
			"updated": this.Milliseconds(),
		}))
	}
	return *this.At(MkString("status"))
}

func (this *Bitbns) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("ccxtGetFetchMarkets"), params)
	_ = response
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
		market := *(response).At(i)
		_ = market
		id := this.SafeString(market, MkString("id"))
		_ = id
		baseId := this.SafeString(market, MkString("base"))
		_ = baseId
		quoteId := this.SafeString(market, MkString("quote"))
		_ = quoteId
		base := this.SafeCurrencyCode(baseId)
		_ = base
		quote := this.SafeCurrencyCode(quoteId)
		_ = quote
		symbol := OpAdd(base, OpAdd(MkString("/"), quote))
		_ = symbol
		marketPrecision := this.SafeValue(market, MkString("precision"), MkMap(&VarMap{}))
		_ = marketPrecision
		precision := MkMap(&VarMap{
			"amount": this.SafeInteger(marketPrecision, MkString("amount")),
			"price":  this.SafeInteger(marketPrecision, MkString("price")),
		})
		_ = precision
		marketLimits := this.SafeValue(market, MkString("limits"), MkMap(&VarMap{}))
		_ = marketLimits
		amountLimits := this.SafeValue(marketLimits, MkString("amount"), MkMap(&VarMap{}))
		_ = amountLimits
		priceLimits := this.SafeValue(marketLimits, MkString("price"), MkMap(&VarMap{}))
		_ = priceLimits
		costLimits := this.SafeValue(marketLimits, MkString("cost"), MkMap(&VarMap{}))
		_ = costLimits
		usdt := OpEq2(quoteId, MkString("USDT"))
		_ = usdt
		uppercaseId := OpTrinary(usdt, OpAdd(baseId, OpAdd(MkString("_"), quoteId)), baseId)
		_ = uppercaseId
		result.Push(MkMap(&VarMap{
			"id":          id,
			"uppercaseId": uppercaseId,
			"symbol":      symbol,
			"base":        base,
			"quote":       quote,
			"baseId":      baseId,
			"quoteId":     quoteId,
			"info":        market,
			"active":      MkUndefined(),
			"precision":   precision,
			"limits": MkMap(&VarMap{
				"amount": MkMap(&VarMap{
					"min": this.SafeNumber(amountLimits, MkString("min")),
					"max": this.SafeNumber(amountLimits, MkString("max")),
				}),
				"price": MkMap(&VarMap{
					"min": this.SafeNumber(priceLimits, MkString("min")),
					"max": this.SafeNumber(priceLimits, MkString("max")),
				}),
				"cost": MkMap(&VarMap{
					"min": this.SafeNumber(costLimits, MkString("min")),
					"max": this.SafeNumber(costLimits, MkString("max")),
				}),
			}),
		}))
	}
	return result
}

func (this *Bitbns) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("ccxtGetFetchOrderbook"), this.Extend(request, params))
	_ = response
	timestamp := this.SafeInteger(response, MkString("timestamp"))
	_ = timestamp
	return this.ParseOrderBook(response, timestamp)
}

func (this *Bitbns) ParseTicker(goArgs ...*Variant) *Variant {
	ticker := GoGetArg(goArgs, 0, MkUndefined())
	_ = ticker
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.SafeInteger(ticker, MkString("timestamp"))
	_ = timestamp
	marketId := this.SafeString(ticker, MkString("symbol"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market)
	_ = symbol
	last := this.SafeNumber(ticker, MkString("last"))
	_ = last
	return MkMap(&VarMap{
		"symbol":        symbol,
		"timestamp":     timestamp,
		"datetime":      this.Iso8601(timestamp),
		"high":          this.SafeNumber(ticker, MkString("high")),
		"low":           this.SafeNumber(ticker, MkString("low")),
		"bid":           this.SafeNumber(ticker, MkString("bid")),
		"bidVolume":     this.SafeNumber(ticker, MkString("bidVolume")),
		"ask":           this.SafeNumber(ticker, MkString("ask")),
		"askVolume":     this.SafeNumber(ticker, MkString("askVolume")),
		"vwap":          this.SafeNumber(ticker, MkString("vwap")),
		"open":          this.SafeNumber(ticker, MkString("open")),
		"close":         last,
		"last":          last,
		"previousClose": this.SafeNumber(ticker, MkString("previousClose")),
		"change":        this.SafeNumber(ticker, MkString("change")),
		"percentage":    this.SafeNumber(ticker, MkString("percentage")),
		"average":       this.SafeNumber(ticker, MkString("average")),
		"baseVolume":    this.SafeNumber(ticker, MkString("baseVolume")),
		"quoteVolume":   this.SafeNumber(ticker, MkString("quoteVolume")),
		"info":          ticker,
	})
}

func (this *Bitbns) FetchTickers(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("ccxtGetFetchTickers"), params)
	_ = response
	return this.ParseTickers(response, symbols)
}

func (this *Bitbns) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("v1PostCurrentCoinBalanceEVERYTHING"), params)
	_ = response
	timestamp := OpCopy(MkUndefined())
	_ = timestamp
	result := MkMap(&VarMap{
		"info":      response,
		"timestamp": timestamp,
		"datetime":  this.Iso8601(timestamp),
	})
	_ = result
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	keys := GoGetKeys(data)
	_ = keys
	for i := MkInteger(0); IsTrue(OpLw(i, keys.Length)); OpIncr(&i) {
		key := *(keys).At(i)
		_ = key
		parts := key.Split(MkString("availableorder"))
		_ = parts
		numParts := OpCopy(parts.Length)
		_ = numParts
		if IsTrue(OpGt(numParts, MkInteger(1))) {
			currencyId := this.SafeString(parts, MkInteger(1))
			_ = currencyId
			if IsTrue(OpNotEq2(currencyId, MkString("Money"))) {
				code := this.SafeCurrencyCode(currencyId)
				_ = code
				account := this.Account()
				_ = account
				*(account).At(MkString("free")) = this.SafeString(data, key)
				*(account).At(MkString("used")) = this.SafeString(data, OpAdd(MkString("inorder"), currencyId))
				*(result).At(code) = OpCopy(account)
			}
		}
	}
	return this.ParseBalance(result)
}

func (this *Bitbns) ParseOrderStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{"0": MkString("open")})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Bitbns) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	id := this.SafeString2(order, MkString("id"), MkString("entry_id"))
	_ = id
	marketId := this.SafeString(order, MkString("symbol"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market)
	_ = symbol
	timestamp := this.Parse8601(this.SafeString(order, MkString("time")))
	_ = timestamp
	price := this.SafeNumber(order, MkString("rate"))
	_ = price
	amount := this.SafeNumber2(order, MkString("amount"), MkString("btc"))
	_ = amount
	filled := this.SafeNumber(order, MkString("filled"))
	_ = filled
	remaining := this.SafeNumber(order, MkString("remaining"))
	_ = remaining
	average := this.SafeNumber(order, MkString("avg_cost"))
	_ = average
	cost := this.SafeNumber(order, MkString("cost"))
	_ = cost
	type_ := this.SafeStringLower(order, MkString("type"))
	_ = type_
	if IsTrue(OpEq2(type_, MkString("0"))) {
		type_ = MkString("limit")
	}
	status := this.ParseOrderStatus(this.SafeString(order, MkString("status")))
	_ = status
	side := this.SafeStringLower(order, MkString("side"))
	_ = side
	feeCost := this.SafeNumber(order, MkString("fee"))
	_ = feeCost
	fee := OpCopy(MkUndefined())
	_ = fee
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		feeCurrencyCode := OpCopy(MkUndefined())
		_ = feeCurrencyCode
		fee = MkMap(&VarMap{
			"cost":     feeCost,
			"currency": feeCurrencyCode,
		})
	}
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
		"cost":               cost,
		"average":            average,
		"filled":             filled,
		"remaining":          remaining,
		"status":             status,
		"fee":                fee,
		"trades":             MkUndefined(),
	}))
}

func (this *Bitbns) CreateOrder(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpNotEq2(type_, MkString("limit"))) {
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), MkString(" allows limit orders only"))))
	}
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{
		"side":     side.ToUpperCase(),
		"symbol":   *(market).At(MkString("uppercaseId")),
		"quantity": this.AmountToPrecision(symbol, amount),
		"rate":     this.PriceToPrecision(symbol, price),
	})
	_ = request
	response := this.Call(MkString("v2PostOrders"), this.Extend(request, params))
	_ = response
	return this.ParseOrder(response, market)
}

func (this *Bitbns) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" cancelOrder() requires a symbol argument"))))
	}
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	quoteSide := OpTrinary(OpEq2(*(market).At(MkString("quoteId")), MkString("USDT")), MkString("usdtcancelOrder"), MkString("cancelOrder"))
	_ = quoteSide
	request := MkMap(&VarMap{
		"entry_id": id,
		"symbol":   *(market).At(MkString("uppercaseId")),
		"side":     quoteSide,
	})
	_ = request
	response := this.Call(MkString("v2PostCancel"), this.Extend(request, params))
	_ = response
	return this.ParseOrder(response, market)
}

func (this *Bitbns) FetchOrder(goArgs ...*Variant) *Variant {
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
		"symbol":   *(market).At(MkString("id")),
		"entry_id": id,
	})
	_ = request
	response := this.Call(MkString("v1PostOrderStatusSymbol"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	first := this.SafeValue(data, MkInteger(0))
	_ = first
	return this.ParseOrder(first, market)
}

func (this *Bitbns) FetchOpenOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchOrders() requires a symbol argument"))))
	}
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	quoteSide := OpTrinary(OpEq2(*(market).At(MkString("quoteId")), MkString("USDT")), MkString("usdtListOpenOrders"), MkString("listOpenOrders"))
	_ = quoteSide
	request := MkMap(&VarMap{
		"symbol": *(market).At(MkString("uppercaseId")),
		"side":   quoteSide,
		"page":   MkInteger(0),
	})
	_ = request
	response := this.Call(MkString("v2PostGetordersnew"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseOrders(data, market, since, limit)
}

func (this *Bitbns) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	market = this.SafeMarket(MkUndefined(), market)
	orderId := this.SafeString(trade, MkString("id"))
	_ = orderId
	timestamp := this.Parse8601(this.SafeString(trade, MkString("date")))
	_ = timestamp
	amountString := this.SafeString(trade, MkString("amount"))
	_ = amountString
	priceString := this.SafeString(trade, MkString("rate"))
	_ = priceString
	price := this.ParseNumber(priceString)
	_ = price
	factor := this.SafeString(trade, MkString("factor"))
	_ = factor
	amountScaled := Precise.StringDiv(amountString, factor)
	_ = amountScaled
	amount := this.ParseNumber(amountScaled)
	_ = amount
	cost := this.ParseNumber(Precise.StringMul(priceString, amountScaled))
	_ = cost
	symbol := *(market).At(MkString("symbol"))
	_ = symbol
	side := this.SafeStringLower(trade, MkString("type"))
	_ = side
	if IsTrue(OpGtEq(side.IndexOf(MkString("sell")), MkInteger(0))) {
		side = MkString("sell")
	} else {
		if IsTrue(OpGtEq(side.IndexOf(MkString("buy")), MkInteger(0))) {
			side = MkString("buy")
		}
	}
	fee := OpCopy(MkUndefined())
	_ = fee
	feeCost := this.SafeNumber(trade, MkString("fee"))
	_ = feeCost
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		feeCurrencyCode := *(market).At(MkString("quote"))
		_ = feeCurrencyCode
		fee = MkMap(&VarMap{
			"cost":     feeCost,
			"currency": feeCurrencyCode,
		})
	}
	return MkMap(&VarMap{
		"info":         trade,
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       symbol,
		"id":           MkUndefined(),
		"order":        orderId,
		"type":         MkUndefined(),
		"side":         side,
		"takerOrMaker": MkUndefined(),
		"price":        price,
		"amount":       amount,
		"cost":         cost,
		"fee":          fee,
	})
}

func (this *Bitbns) FetchMyTrades(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchOrders() requires a symbol argument"))))
	}
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{
		"symbol": *(market).At(MkString("id")),
		"page":   MkInteger(0),
	})
	_ = request
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("since")) = this.Iso8601(since)
	}
	response := this.Call(MkString("v1PostListExecutedOrdersSymbol"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseTrades(data, market, since, limit)
}

func (this *Bitbns) FetchDeposits(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{
		"symbol": *(currency).At(MkString("id")),
		"page":   MkInteger(0),
	})
	_ = request
	response := this.Call(MkString("v1PostDepositHistorySymbol"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseTransactions(data, currency, since, limit)
}

func (this *Bitbns) FetchWithdrawals(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{
		"symbol": *(currency).At(MkString("id")),
		"page":   MkInteger(0),
	})
	_ = request
	response := this.Call(MkString("v1PostWithdrawHistorySymbol"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseTransactions(data, currency, since, limit)
}

func (this *Bitbns) ParseTransactionStatusByType(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	type_ := GoGetArg(goArgs, 1, MkUndefined())
	_ = type_
	statusesByType := MkMap(&VarMap{
		"deposit": MkMap(&VarMap{
			"0": MkString("pending"),
			"1": MkString("ok"),
		}),
		"withdrawal": MkMap(&VarMap{
			"0": MkString("pending"),
			"1": MkString("canceled"),
			"2": MkString("pending"),
			"3": MkString("failed"),
			"4": MkString("pending"),
			"5": MkString("failed"),
			"6": MkString("ok"),
		}),
	})
	_ = statusesByType
	statuses := this.SafeValue(statusesByType, type_, MkMap(&VarMap{}))
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Bitbns) ParseTransaction(goArgs ...*Variant) *Variant {
	transaction := GoGetArg(goArgs, 0, MkUndefined())
	_ = transaction
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	currencyId := this.SafeString(transaction, MkString("unit"))
	_ = currencyId
	code := this.SafeCurrencyCode(currencyId, currency)
	_ = code
	timestamp := this.Parse8601(this.SafeString(transaction, MkString("date")))
	_ = timestamp
	type_ := this.SafeString(transaction, MkString("type"))
	_ = type_
	status := OpCopy(MkUndefined())
	_ = status
	if IsTrue(OpNotEq2(type_, MkUndefined())) {
		if IsTrue(OpGtEq(type_.IndexOf(MkString("deposit")), MkInteger(0))) {
			type_ = MkString("deposit")
			status = MkString("ok")
		} else {
			if IsTrue(OpGtEq(type_.IndexOf(MkString("withdraw")), MkInteger(0))) {
				type_ = MkString("withdrawal")
			}
		}
	}
	amount := this.SafeNumber(transaction, MkString("amount"))
	_ = amount
	feeCost := this.SafeNumber(transaction, MkString("fee"))
	_ = feeCost
	fee := OpCopy(MkUndefined())
	_ = fee
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		fee = MkMap(&VarMap{
			"currency": code,
			"cost":     feeCost,
		})
	}
	return MkMap(&VarMap{
		"info":        transaction,
		"id":          MkUndefined(),
		"txid":        MkUndefined(),
		"timestamp":   timestamp,
		"datetime":    this.Iso8601(timestamp),
		"address":     MkUndefined(),
		"addressTo":   MkUndefined(),
		"addressFrom": MkUndefined(),
		"tag":         MkUndefined(),
		"tagTo":       MkUndefined(),
		"tagFrom":     MkUndefined(),
		"type":        type_,
		"amount":      amount,
		"currency":    code,
		"status":      status,
		"updated":     MkUndefined(),
		"internal":    MkUndefined(),
		"fee":         fee,
	})
}

func (this *Bitbns) FetchDepositAddress(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{"symbol": *(currency).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("v1PostGetCoinAddressSymbol"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	address := this.SafeString(data, MkString("token"))
	_ = address
	tag := this.SafeString(data, MkString("tag"))
	_ = tag
	this.CheckAddress(address)
	return MkMap(&VarMap{
		"currency": code,
		"address":  address,
		"tag":      tag,
		"info":     response,
	})
}

func (this *Bitbns) Nonce(goArgs ...*Variant) *Variant {
	return this.Milliseconds()
}

func (this *Bitbns) Sign(goArgs ...*Variant) *Variant {
	path := GoGetArg(goArgs, 0, MkUndefined())
	_ = path
	api := GoGetArg(goArgs, 1, MkString("v1"))
	_ = api
	method := GoGetArg(goArgs, 2, MkString("GET"))
	_ = method
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	headers := GoGetArg(goArgs, 4, MkUndefined())
	_ = headers
	body := GoGetArg(goArgs, 5, MkUndefined())
	_ = body
	this.CheckRequiredCredentials()
	baseUrl := this.ImplodeHostname(*(*(*this.At(MkString("urls"))).At(MkString("api"))).At(api))
	_ = baseUrl
	url := OpAdd(baseUrl, OpAdd(MkString("/"), this.ImplodeParams(path, params)))
	_ = url
	query := this.Omit(params, this.ExtractParams(path))
	_ = query
	nonce := (this.Nonce()).Call(MkString("toString"))
	_ = nonce
	headers = MkMap(&VarMap{"X-BITBNS-APIKEY": *this.At(MkString("apiKey"))})
	if IsTrue(OpEq2(method, MkString("GET"))) {
		if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
			OpAddAssign(&url, OpAdd(MkString("?"), this.Urlencode(query)))
		}
	} else {
		if IsTrue(OpEq2(method, MkString("POST"))) {
			if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
				body = this.Json(query)
			} else {
				body = MkString("{}")
			}
			auth := MkMap(&VarMap{
				"timeStamp_nonce": nonce,
				"body":            body,
			})
			_ = auth
			payload := this.StringToBase64(this.Json(auth))
			_ = payload
			signature := this.Hmac(payload, this.Encode(*this.At(MkString("secret"))), MkString("sha512"))
			_ = signature
			*(headers).At(MkString("X-BITBNS-PAYLOAD")) = this.Decode(payload)
			*(headers).At(MkString("X-BITBNS-SIGNATURE")) = OpCopy(signature)
			*(headers).At(MkString("Content-Type")) = MkString("application/x-www-form-urlencoded")
		}
	}
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Bitbns) HandleErrors(goArgs ...*Variant) *Variant {
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
	message := this.SafeString(response, MkString("msg"))
	_ = message
	error := OpAnd(OpNotEq2(code, MkUndefined()), OpNotEq2(code, MkString("200")))
	_ = error
	if IsTrue(OpOr(error, OpNotEq2(message, MkUndefined()))) {
		feedback := OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))
		_ = feedback
		this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("exact")), code, feedback)
		this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("exact")), message, feedback)
		this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("broad")), message, feedback)
		panic(NewExchangeError(feedback))
	}
	return MkUndefined()
}
