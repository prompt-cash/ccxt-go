package ccxt

import ()

type Coinone struct {
	*ExchangeBase
}

var _ Exchange = (*Coinone)(nil)

func init() {
	exchange := &Coinone{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Coinone) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("coinone"),
		"name": MkString("CoinOne"),
		"countries": MkArray(&VarArray{
			MkString("KR"),
		}),
		"rateLimit": MkInteger(667),
		"version":   MkString("v2"),
		"has": MkMap(&VarMap{
			"cancelOrder":       MkBool(true),
			"CORS":              MkBool(false),
			"createMarketOrder": MkBool(false),
			"createOrder":       MkBool(true),
			"fetchBalance":      MkBool(true),
			"fetchCurrencies":   MkBool(false),
			"fetchMarkets":      MkBool(true),
			"fetchMyTrades":     MkBool(true),
			"fetchOpenOrders":   MkBool(true),
			"fetchOrder":        MkBool(true),
			"fetchOrderBook":    MkBool(true),
			"fetchTicker":       MkBool(true),
			"fetchTickers":      MkBool(true),
			"fetchTrades":       MkBool(true),
			"fetchClosedOrders": MkBool(false),
		}),
		"urls": MkMap(&VarMap{
			"logo": MkString("https://user-images.githubusercontent.com/1294454/38003300-adc12fba-323f-11e8-8525-725f53c4a659.jpg"),
			"api":  MkString("https://api.coinone.co.kr"),
			"www":  MkString("https://coinone.co.kr"),
			"doc":  MkString("https://doc.coinone.co.kr"),
		}),
		"requiredCredentials": MkMap(&VarMap{
			"apiKey": MkBool(true),
			"secret": MkBool(true),
		}),
		"api": MkMap(&VarMap{
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("orderbook/"),
				MkString("trades/"),
				MkString("ticker/"),
			})}),
			"private": MkMap(&VarMap{"post": MkArray(&VarArray{
				MkString("account/btc_deposit_address/"),
				MkString("account/balance/"),
				MkString("account/daily_balance/"),
				MkString("account/user_info/"),
				MkString("account/virtual_account/"),
				MkString("order/cancel_all/"),
				MkString("order/cancel/"),
				MkString("order/limit_buy/"),
				MkString("order/limit_sell/"),
				MkString("order/complete_orders/"),
				MkString("order/limit_orders/"),
				MkString("order/order_info/"),
				MkString("transaction/auth_number/"),
				MkString("transaction/history/"),
				MkString("transaction/krw/history/"),
				MkString("transaction/btc/"),
				MkString("transaction/coin/"),
			})}),
		}),
		"fees": MkMap(&VarMap{"trading": MkMap(&VarMap{
			"tierBased":  MkBool(false),
			"percentage": MkBool(true),
			"taker":      MkNumber(0.002),
			"maker":      MkNumber(0.002),
		})}),
		"precision": MkMap(&VarMap{
			"price":  MkInteger(4),
			"amount": MkInteger(4),
			"cost":   MkInteger(8),
		}),
		"exceptions": MkMap(&VarMap{
			"405": OnMaintenance,
			"104": OrderNotFound,
			"108": BadSymbol,
			"107": BadRequest,
		}),
		"commonCurrencies": MkMap(&VarMap{"SOC": MkString("Soda Coin")}),
	}))
}

func (this *Coinone) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	request := MkMap(&VarMap{"currency": MkString("all")})
	_ = request
	response := this.Call(MkString("publicGetTicker"), request)
	_ = response
	result := MkArray(&VarArray{})
	_ = result
	quoteId := MkString("krw")
	_ = quoteId
	quote := this.SafeCurrencyCode(quoteId)
	_ = quote
	baseIds := GoGetKeys(response)
	_ = baseIds
	for i := MkInteger(0); IsTrue(OpLw(i, baseIds.Length)); OpIncr(&i) {
		baseId := *(baseIds).At(i)
		_ = baseId
		ticker := this.SafeValue(response, baseId, MkMap(&VarMap{}))
		_ = ticker
		currency := this.SafeValue(ticker, MkString("currency"))
		_ = currency
		if IsTrue(OpEq2(currency, MkUndefined())) {
			continue
		}
		base := this.SafeCurrencyCode(baseId)
		_ = base
		result.Push(MkMap(&VarMap{
			"info":    ticker,
			"id":      baseId,
			"symbol":  OpAdd(base, OpAdd(MkString("/"), quote)),
			"base":    base,
			"quote":   quote,
			"baseId":  baseId,
			"quoteId": quoteId,
			"active":  MkBool(true),
		}))
	}
	return result
}

func (this *Coinone) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privatePostAccountBalance"), params)
	_ = response
	result := MkMap(&VarMap{"info": response})
	_ = result
	balances := this.Omit(response, MkArray(&VarArray{
		MkString("errorCode"),
		MkString("result"),
		MkString("normalWallets"),
	}))
	_ = balances
	currencyIds := GoGetKeys(balances)
	_ = currencyIds
	for i := MkInteger(0); IsTrue(OpLw(i, currencyIds.Length)); OpIncr(&i) {
		currencyId := *(currencyIds).At(i)
		_ = currencyId
		balance := *(balances).At(currencyId)
		_ = balance
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		account := this.Account()
		_ = account
		*(account).At(MkString("free")) = this.SafeString(balance, MkString("avail"))
		*(account).At(MkString("total")) = this.SafeString(balance, MkString("balance"))
		*(result).At(code) = OpCopy(account)
	}
	return this.ParseBalance(result)
}

func (this *Coinone) FetchOrderBook(goArgs ...*Variant) *Variant {
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
		"currency": *(market).At(MkString("id")),
		"format":   MkString("json"),
	})
	_ = request
	response := this.Call(MkString("publicGetOrderbook"), this.Extend(request, params))
	_ = response
	timestamp := this.SafeTimestamp(response, MkString("timestamp"))
	_ = timestamp
	return this.ParseOrderBook(response, symbol, timestamp, MkString("bid"), MkString("ask"), MkString("price"), MkString("qty"))
}

func (this *Coinone) FetchTickers(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{
		"currency": MkString("all"),
		"format":   MkString("json"),
	})
	_ = request
	response := this.Call(MkString("publicGetTicker"), this.Extend(request, params))
	_ = response
	result := MkMap(&VarMap{})
	_ = result
	ids := GoGetKeys(response)
	_ = ids
	timestamp := this.SafeTimestamp(response, MkString("timestamp"))
	_ = timestamp
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
			ticker := *(response).At(id)
			_ = ticker
			*(result).At(symbol) = this.ParseTicker(ticker, market)
			*(*(result).At(symbol)).At(MkString("timestamp")) = OpCopy(timestamp)
		}
	}
	return this.FilterByArray(result, MkString("symbol"), symbols)
}

func (this *Coinone) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{
		"currency": *(market).At(MkString("id")),
		"format":   MkString("json"),
	})
	_ = request
	response := this.Call(MkString("publicGetTicker"), this.Extend(request, params))
	_ = response
	return this.ParseTicker(response, market)
}

func (this *Coinone) ParseTicker(goArgs ...*Variant) *Variant {
	ticker := GoGetArg(goArgs, 0, MkUndefined())
	_ = ticker
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.SafeTimestamp(ticker, MkString("timestamp"))
	_ = timestamp
	first := this.SafeNumber(ticker, MkString("first"))
	_ = first
	last := this.SafeNumber(ticker, MkString("last"))
	_ = last
	average := OpCopy(MkUndefined())
	_ = average
	if IsTrue(OpAnd(OpNotEq2(first, MkUndefined()), OpNotEq2(last, MkUndefined()))) {
		average = OpDiv(this.Sum(first, last), MkInteger(2))
	}
	previousClose := this.SafeNumber(ticker, MkString("yesterday_last"))
	_ = previousClose
	change := OpCopy(MkUndefined())
	_ = change
	percentage := OpCopy(MkUndefined())
	_ = percentage
	if IsTrue(OpAnd(OpNotEq2(last, MkUndefined()), OpNotEq2(previousClose, MkUndefined()))) {
		change = OpSub(last, previousClose)
		if IsTrue(OpNotEq2(previousClose, MkInteger(0))) {
			percentage = OpDiv(change, OpMulti(previousClose, MkInteger(100)))
		}
	}
	symbol := OpTrinary(OpNotEq2(market, MkUndefined()), *(market).At(MkString("symbol")), MkUndefined())
	_ = symbol
	return MkMap(&VarMap{
		"symbol":        symbol,
		"timestamp":     timestamp,
		"datetime":      this.Iso8601(timestamp),
		"high":          this.SafeNumber(ticker, MkString("high")),
		"low":           this.SafeNumber(ticker, MkString("low")),
		"bid":           MkUndefined(),
		"bidVolume":     MkUndefined(),
		"ask":           MkUndefined(),
		"askVolume":     MkUndefined(),
		"vwap":          MkUndefined(),
		"open":          first,
		"close":         last,
		"last":          last,
		"previousClose": previousClose,
		"change":        change,
		"percentage":    percentage,
		"average":       average,
		"baseVolume":    this.SafeNumber(ticker, MkString("volume")),
		"quoteVolume":   MkUndefined(),
		"info":          ticker,
	})
}

func (this *Coinone) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.SafeTimestamp(trade, MkString("timestamp"))
	_ = timestamp
	symbol := OpTrinary(OpNotEq2(market, MkUndefined()), *(market).At(MkString("symbol")), MkUndefined())
	_ = symbol
	is_ask := this.SafeString(trade, MkString("is_ask"))
	_ = is_ask
	side := this.SafeString(trade, MkString("type"))
	_ = side
	if IsTrue(OpNotEq2(is_ask, MkUndefined())) {
		if IsTrue(OpEq2(is_ask, MkString("1"))) {
			side = MkString("sell")
		} else {
			if IsTrue(OpEq2(is_ask, MkString("0"))) {
				side = MkString("buy")
			}
		}
	} else {
		if IsTrue(OpEq2(side, MkString("ask"))) {
			side = MkString("sell")
		} else {
			if IsTrue(OpEq2(side, MkString("bid"))) {
				side = MkString("buy")
			}
		}
	}
	priceString := this.SafeString(trade, MkString("price"))
	_ = priceString
	amountString := this.SafeString(trade, MkString("qty"))
	_ = amountString
	price := this.ParseNumber(priceString)
	_ = price
	amount := this.ParseNumber(amountString)
	_ = amount
	cost := this.ParseNumber(Precise.StringMul(priceString, amountString))
	_ = cost
	orderId := this.SafeString(trade, MkString("orderId"))
	_ = orderId
	feeCost := this.SafeNumber(trade, MkString("fee"))
	_ = feeCost
	fee := OpCopy(MkUndefined())
	_ = fee
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		feeCost = Math.Abs(feeCost)
		feeRate := this.SafeNumber(trade, MkString("feeRate"))
		_ = feeRate
		feeRate = Math.Abs(feeRate)
		feeCurrencyCode := OpCopy(MkUndefined())
		_ = feeCurrencyCode
		if IsTrue(OpNotEq2(market, MkUndefined())) {
			feeCurrencyCode = OpTrinary(OpEq2(side, MkString("sell")), *(market).At(MkString("quote")), *(market).At(MkString("base")))
		}
		fee = MkMap(&VarMap{
			"cost":     feeCost,
			"currency": feeCurrencyCode,
			"rate":     feeRate,
		})
	}
	return MkMap(&VarMap{
		"id":           this.SafeString(trade, MkString("id")),
		"info":         trade,
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"order":        orderId,
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

func (this *Coinone) FetchTrades(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{
		"currency": *(market).At(MkString("id")),
		"format":   MkString("json"),
	})
	_ = request
	response := this.Call(MkString("publicGetTrades"), this.Extend(request, params))
	_ = response
	completeOrders := this.SafeValue(response, MkString("completeOrders"), MkArray(&VarArray{}))
	_ = completeOrders
	return this.ParseTrades(completeOrders, market, since, limit)
}

func (this *Coinone) CreateOrder(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{
		"price":    price,
		"currency": this.MarketId(symbol),
		"qty":      amount,
	})
	_ = request
	method := OpAdd(MkString("privatePostOrder"), OpAdd(this.Capitalize(type_), this.Capitalize(side)))
	_ = method
	response := this.Call(method, this.Extend(request, params))
	_ = response
	return this.ParseOrder(response)
}

func (this *Coinone) FetchOrder(goArgs ...*Variant) *Variant {
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
		"order_id": id,
		"currency": *(market).At(MkString("id")),
	})
	_ = request
	response := this.Call(MkString("privatePostOrderOrderInfo"), this.Extend(request, params))
	_ = response
	info := this.SafeValue(response, MkString("info"), MkMap(&VarMap{}))
	_ = info
	*(info).At(MkString("status")) = this.SafeString(info, MkString("status"))
	return this.ParseOrder(info, market)
}

func (this *Coinone) ParseOrderStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"live":             MkString("open"),
		"partially_filled": MkString("open"),
		"filled":           MkString("closed"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Coinone) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	id := this.SafeString(order, MkString("orderId"))
	_ = id
	price := this.SafeNumber(order, MkString("price"))
	_ = price
	timestamp := this.SafeTimestamp(order, MkString("timestamp"))
	_ = timestamp
	side := this.SafeString(order, MkString("type"))
	_ = side
	if IsTrue(OpEq2(side, MkString("ask"))) {
		side = MkString("sell")
	} else {
		if IsTrue(OpEq2(side, MkString("bid"))) {
			side = MkString("buy")
		}
	}
	remaining := this.SafeNumber(order, MkString("remainQty"))
	_ = remaining
	amount := this.SafeNumber(order, MkString("qty"))
	_ = amount
	status := this.SafeString(order, MkString("status"))
	_ = status
	if IsTrue(OpEq2(status, MkString("live"))) {
		if IsTrue(OpAnd(OpNotEq2(remaining, MkUndefined()), OpNotEq2(amount, MkUndefined()))) {
			if IsTrue(OpLw(remaining, amount)) {
				status = MkString("canceled")
			}
		}
	}
	status = this.ParseOrderStatus(status)
	symbol := OpCopy(MkUndefined())
	_ = symbol
	base := OpCopy(MkUndefined())
	_ = base
	quote := OpCopy(MkUndefined())
	_ = quote
	marketId := this.SafeStringLower(order, MkString("currency"))
	_ = marketId
	if IsTrue(OpNotEq2(marketId, MkUndefined())) {
		if IsTrue(OpHasMember(marketId, *this.At(MkString("markets_by_id")))) {
			market = *(*this.At(MkString("markets_by_id"))).At(marketId)
		} else {
			base = this.SafeCurrencyCode(marketId)
			quote = MkString("KRW")
			symbol = OpAdd(base, OpAdd(MkString("/"), quote))
		}
	}
	if IsTrue(OpAnd(OpEq2(symbol, MkUndefined()), OpNotEq2(market, MkUndefined()))) {
		symbol = *(market).At(MkString("symbol"))
		base = *(market).At(MkString("base"))
		quote = *(market).At(MkString("quote"))
	}
	fee := OpCopy(MkUndefined())
	_ = fee
	feeCost := this.SafeNumber(order, MkString("fee"))
	_ = feeCost
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		feeCurrencyCode := OpTrinary(OpEq2(side, MkString("sell")), quote, base)
		_ = feeCurrencyCode
		fee = MkMap(&VarMap{
			"cost":     feeCost,
			"rate":     this.SafeNumber(order, MkString("feeRate")),
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
		"type":               MkString("limit"),
		"timeInForce":        MkUndefined(),
		"postOnly":           MkUndefined(),
		"side":               side,
		"price":              price,
		"stopPrice":          MkUndefined(),
		"cost":               MkUndefined(),
		"average":            MkUndefined(),
		"amount":             amount,
		"filled":             MkUndefined(),
		"remaining":          remaining,
		"status":             status,
		"fee":                fee,
		"trades":             MkUndefined(),
	}))
}

func (this *Coinone) FetchOpenOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), MkString(" allows fetching closed orders with a specific symbol"))))
	}
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"currency": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("privatePostOrderLimitOrders"), this.Extend(request, params))
	_ = response
	limitOrders := this.SafeValue(response, MkString("limitOrders"), MkArray(&VarArray{}))
	_ = limitOrders
	return this.ParseOrders(limitOrders, market, since, limit)
}

func (this *Coinone) FetchMyTrades(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"currency": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("privatePostOrderCompleteOrders"), this.Extend(request, params))
	_ = response
	completeOrders := this.SafeValue(response, MkString("completeOrders"), MkArray(&VarArray{}))
	_ = completeOrders
	return this.ParseTrades(completeOrders, market, since, limit)
}

func (this *Coinone) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" cancelOrder() requires a symbol argument. To cancel the order, pass a symbol argument and {'price': 12345, 'qty': 1.2345, 'is_ask': 0} in the params argument of cancelOrder."))))
	}
	price := this.SafeNumber(params, MkString("price"))
	_ = price
	qty := this.SafeNumber(params, MkString("qty"))
	_ = qty
	isAsk := this.SafeInteger(params, MkString("is_ask"))
	_ = isAsk
	if IsTrue(OpOr(OpEq2(price, MkUndefined()), OpOr(OpEq2(qty, MkUndefined()), OpEq2(isAsk, MkUndefined())))) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" cancelOrder() requires {'price': 12345, 'qty': 1.2345, 'is_ask': 0} in the params argument."))))
	}
	this.LoadMarkets()
	request := MkMap(&VarMap{
		"order_id": id,
		"price":    price,
		"qty":      qty,
		"is_ask":   isAsk,
		"currency": this.MarketId(symbol),
	})
	_ = request
	response := this.Call(MkString("privatePostOrderCancel"), this.Extend(request, params))
	_ = response
	return response
}

func (this *Coinone) Sign(goArgs ...*Variant) *Variant {
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
	request := this.ImplodeParams(path, params)
	_ = request
	query := this.Omit(params, this.ExtractParams(path))
	_ = query
	url := OpAdd(*(*this.At(MkString("urls"))).At(MkString("api")), MkString("/"))
	_ = url
	if IsTrue(OpEq2(api, MkString("public"))) {
		OpAddAssign(&url, request)
		if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
			OpAddAssign(&url, OpAdd(MkString("?"), this.Urlencode(query)))
		}
	} else {
		this.CheckRequiredCredentials()
		OpAddAssign(&url, OpAdd(*this.At(MkString("version")), OpAdd(MkString("/"), request)))
		nonce := (this.Nonce()).Call(MkString("toString"))
		_ = nonce
		json := this.Json(this.Extend(MkMap(&VarMap{
			"access_token": *this.At(MkString("apiKey")),
			"nonce":        nonce,
		}), params))
		_ = json
		payload := this.StringToBase64(json)
		_ = payload
		body = this.Decode(payload)
		secret := (*((this).At(MkString("secret")))).Call(MkString("toUpperCase"))
		_ = secret
		signature := this.Hmac(payload, this.Encode(secret), MkString("sha512"))
		_ = signature
		headers = MkMap(&VarMap{
			"Content-Type":        MkString("application/json"),
			"X-COINONE-PAYLOAD":   payload,
			"X-COINONE-SIGNATURE": signature,
		})
	}
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Coinone) HandleErrors(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpHasMember(MkString("result"), response)) {
		result := *(response).At(MkString("result"))
		_ = result
		if IsTrue(OpNotEq2(result, MkString("success"))) {
			errorCode := this.SafeString(response, MkString("errorCode"))
			_ = errorCode
			feedback := OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))
			_ = feedback
			this.ThrowExactlyMatchedException(*this.At(MkString("exceptions")), errorCode, feedback)
			panic(NewExchangeError(feedback))
		}
	} else {
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))))
	}
	return MkUndefined()
}
