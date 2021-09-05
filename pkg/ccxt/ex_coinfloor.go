package ccxt

import ()

type Coinfloor struct {
	*ExchangeBase
}

var _ Exchange = (*Coinfloor)(nil)

func init() {
	exchange := &Coinfloor{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Coinfloor) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":        MkString("coinfloor"),
		"name":      MkString("coinfloor"),
		"rateLimit": MkInteger(1000),
		"countries": MkArray(&VarArray{
			MkString("UK"),
		}),
		"has": MkMap(&VarMap{
			"cancelOrder":     MkBool(true),
			"CORS":            MkBool(false),
			"createOrder":     MkBool(true),
			"fetchBalance":    MkBool(true),
			"fetchLedger":     MkBool(true),
			"fetchOpenOrders": MkBool(true),
			"fetchOrderBook":  MkBool(true),
			"fetchTicker":     MkBool(true),
			"fetchTrades":     MkBool(true),
		}),
		"urls": MkMap(&VarMap{
			"logo": MkString("https://user-images.githubusercontent.com/51840849/87153925-ef265e80-c2c0-11ea-91b5-020c804b90e0.jpg"),
			"api":  MkString("https://webapi.coinfloor.co.uk/v2/bist"),
			"www":  MkString("https://www.coinfloor.co.uk"),
			"doc": MkArray(&VarArray{
				MkString("https://github.com/coinfloor/api"),
				MkString("https://www.coinfloor.co.uk/api"),
			}),
		}),
		"requiredCredentials": MkMap(&VarMap{
			"apiKey":   MkBool(true),
			"secret":   MkBool(false),
			"password": MkBool(true),
			"uid":      MkBool(true),
		}),
		"api": MkMap(&VarMap{
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("{id}/ticker/"),
				MkString("{id}/order_book/"),
				MkString("{id}/transactions/"),
			})}),
			"private": MkMap(&VarMap{"post": MkArray(&VarArray{
				MkString("{id}/balance/"),
				MkString("{id}/user_transactions/"),
				MkString("{id}/open_orders/"),
				MkString("{symbol}/cancel_order/"),
				MkString("{id}/buy/"),
				MkString("{id}/sell/"),
				MkString("{id}/buy_market/"),
				MkString("{id}/sell_market/"),
				MkString("{id}/estimate_sell_market/"),
				MkString("{id}/estimate_buy_market/"),
			})}),
		}),
		"markets": MkMap(&VarMap{
			"BTC/GBP": MkMap(&VarMap{
				"id":      MkString("XBT/GBP"),
				"symbol":  MkString("BTC/GBP"),
				"base":    MkString("BTC"),
				"quote":   MkString("GBP"),
				"baseId":  MkString("XBT"),
				"quoteId": MkString("GBP"),
				"precision": MkMap(&VarMap{
					"price":  MkInteger(0),
					"amount": MkInteger(4),
				}),
			}),
			"BTC/EUR": MkMap(&VarMap{
				"id":      MkString("XBT/EUR"),
				"symbol":  MkString("BTC/EUR"),
				"base":    MkString("BTC"),
				"quote":   MkString("EUR"),
				"baseId":  MkString("XBT"),
				"quoteId": MkString("EUR"),
				"precision": MkMap(&VarMap{
					"price":  MkInteger(0),
					"amount": MkInteger(4),
				}),
			}),
		}),
		"exceptions": MkMap(&VarMap{"exact": MkMap(&VarMap{
			"You have insufficient funds.": InsufficientFunds,
			"Tonce is out of sequence.":    InvalidNonce,
		})}),
	}))
}

func (this *Coinfloor) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := OpCopy(MkUndefined())
	_ = market
	query := OpCopy(params)
	_ = query
	symbol := this.SafeString(params, MkString("symbol"))
	_ = symbol
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(*(params).At(MkString("symbol")))
		query = this.Omit(params, MkString("symbol"))
	}
	marketId := this.SafeString(params, MkString("id"))
	_ = marketId
	if IsTrue(OpHasMember(marketId, *this.At(MkString("markets_by_id")))) {
		market = *(*this.At(MkString("markets_by_id"))).At(marketId)
	}
	if IsTrue(OpEq2(market, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchBalance() requires a symbol param"))))
	}
	request := MkMap(&VarMap{"id": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("privatePostIdBalance"), this.Extend(request, query))
	_ = response
	result := MkMap(&VarMap{"info": response})
	_ = result
	base := *(market).At(MkString("base"))
	_ = base
	quote := *(market).At(MkString("quote"))
	_ = quote
	baseIdLower := this.SafeStringLower(market, MkString("baseId"))
	_ = baseIdLower
	quoteIdLower := this.SafeStringLower(market, MkString("quoteId"))
	_ = quoteIdLower
	*(result).At(base) = MkMap(&VarMap{
		"free":  this.SafeString(response, OpAdd(baseIdLower, MkString("_available"))),
		"used":  this.SafeString(response, OpAdd(baseIdLower, MkString("_reserved"))),
		"total": this.SafeString(response, OpAdd(baseIdLower, MkString("_balance"))),
	})
	*(result).At(quote) = MkMap(&VarMap{
		"free":  this.SafeString(response, OpAdd(quoteIdLower, MkString("_available"))),
		"used":  this.SafeString(response, OpAdd(quoteIdLower, MkString("_reserved"))),
		"total": this.SafeString(response, OpAdd(quoteIdLower, MkString("_balance"))),
	})
	return this.ParseBalance(result)
}

func (this *Coinfloor) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"id": this.MarketId(symbol)})
	_ = request
	response := this.Call(MkString("publicGetIdOrderBook"), this.Extend(request, params))
	_ = response
	return this.ParseOrderBook(response, symbol)
}

func (this *Coinfloor) ParseTicker(goArgs ...*Variant) *Variant {
	ticker := GoGetArg(goArgs, 0, MkUndefined())
	_ = ticker
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.Milliseconds()
	_ = timestamp
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(OpNotEq2(market, MkUndefined())) {
		symbol = *(market).At(MkString("symbol"))
	}
	vwap := this.SafeNumber(ticker, MkString("vwap"))
	_ = vwap
	baseVolume := this.SafeNumber(ticker, MkString("volume"))
	_ = baseVolume
	quoteVolume := OpCopy(MkUndefined())
	_ = quoteVolume
	if IsTrue(OpNotEq2(vwap, MkUndefined())) {
		quoteVolume = OpMulti(baseVolume, vwap)
	}
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
		"vwap":          vwap,
		"open":          MkUndefined(),
		"close":         last,
		"last":          last,
		"previousClose": MkUndefined(),
		"change":        MkUndefined(),
		"percentage":    MkUndefined(),
		"average":       MkUndefined(),
		"baseVolume":    baseVolume,
		"quoteVolume":   quoteVolume,
		"info":          ticker,
	})
}

func (this *Coinfloor) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"id": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("publicGetIdTicker"), this.Extend(request, params))
	_ = response
	return this.ParseTicker(response, market)
}

func (this *Coinfloor) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.SafeTimestamp(trade, MkString("date"))
	_ = timestamp
	id := this.SafeString(trade, MkString("tid"))
	_ = id
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
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(OpNotEq2(market, MkUndefined())) {
		symbol = *(market).At(MkString("symbol"))
	}
	return MkMap(&VarMap{
		"info":         trade,
		"id":           id,
		"order":        MkUndefined(),
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       symbol,
		"type":         MkUndefined(),
		"side":         MkUndefined(),
		"takerOrMaker": MkUndefined(),
		"price":        price,
		"amount":       amount,
		"cost":         cost,
		"fee":          MkUndefined(),
	})
}

func (this *Coinfloor) FetchTrades(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"id": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("publicGetIdTransactions"), this.Extend(request, params))
	_ = response
	return this.ParseTrades(response, market, since, limit)
}

func (this *Coinfloor) FetchLedger(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := OpCopy(MkUndefined())
	_ = market
	if IsTrue(OpNotEq2(code, MkUndefined())) {
		market = this.Market(code)
		if IsTrue(OpEq2(market, MkUndefined())) {
			panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchTransactions() requires a code argument (a market symbol)"))))
		}
	}
	request := MkMap(&VarMap{
		"id":    *(market).At(MkString("id")),
		"limit": limit,
	})
	_ = request
	response := this.Call(MkString("privatePostIdUserTransactions"), this.Extend(request, params))
	_ = response
	return this.ParseLedger(response, MkUndefined(), since, MkUndefined())
}

func (this *Coinfloor) ParseLedgerEntryStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	types := MkMap(&VarMap{"completed": MkString("ok")})
	_ = types
	return this.SafeString(types, status, status)
}

func (this *Coinfloor) ParseLedgerEntryType(goArgs ...*Variant) *Variant {
	type_ := GoGetArg(goArgs, 0, MkUndefined())
	_ = type_
	types := MkMap(&VarMap{
		"0": MkString("transaction"),
		"1": MkString("transaction"),
		"2": MkString("trade"),
	})
	_ = types
	return this.SafeString(types, type_, type_)
}

func (this *Coinfloor) ParseLedgerEntry(goArgs ...*Variant) *Variant {
	item := GoGetArg(goArgs, 0, MkUndefined())
	_ = item
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	keys := GoGetKeys(item)
	_ = keys
	baseId := OpCopy(MkUndefined())
	_ = baseId
	quoteId := OpCopy(MkUndefined())
	_ = quoteId
	baseAmount := OpCopy(MkUndefined())
	_ = baseAmount
	quoteAmount := OpCopy(MkUndefined())
	_ = quoteAmount
	for i := MkInteger(0); IsTrue(OpLw(i, keys.Length)); OpIncr(&i) {
		key := *(keys).At(i)
		_ = key
		if IsTrue(OpGt(key.IndexOf(MkString("_")), MkInteger(0))) {
			parts := key.Split(MkString("_"))
			_ = parts
			numParts := OpCopy(parts.Length)
			_ = numParts
			if IsTrue(OpEq2(numParts, MkInteger(2))) {
				tmpBaseAmount := this.SafeNumber(item, *(parts).At(MkInteger(0)))
				_ = tmpBaseAmount
				tmpQuoteAmount := this.SafeNumber(item, *(parts).At(MkInteger(1)))
				_ = tmpQuoteAmount
				if IsTrue(OpAnd(OpNotEq2(tmpBaseAmount, MkUndefined()), OpNotEq2(tmpQuoteAmount, MkUndefined()))) {
					baseId = *(parts).At(MkInteger(0))
					quoteId = *(parts).At(MkInteger(1))
					baseAmount = OpCopy(tmpBaseAmount)
					quoteAmount = OpCopy(tmpQuoteAmount)
				}
			}
		}
	}
	base := this.SafeCurrencyCode(baseId)
	_ = base
	quote := this.SafeCurrencyCode(quoteId)
	_ = quote
	type_ := this.ParseLedgerEntryType(this.SafeString(item, MkString("type")))
	_ = type_
	referenceId := this.SafeString(item, MkString("id"))
	_ = referenceId
	timestamp := this.Parse8601(this.SafeString(item, MkString("datetime")))
	_ = timestamp
	fee := OpCopy(MkUndefined())
	_ = fee
	feeCost := this.SafeNumber(item, MkString("fee"))
	_ = feeCost
	result := MkMap(&VarMap{
		"id":               MkUndefined(),
		"timestamp":        timestamp,
		"datetime":         this.Iso8601(timestamp),
		"amount":           MkUndefined(),
		"direction":        MkUndefined(),
		"currency":         MkUndefined(),
		"type":             type_,
		"referenceId":      referenceId,
		"referenceAccount": MkUndefined(),
		"before":           MkUndefined(),
		"after":            MkUndefined(),
		"status":           MkString("ok"),
		"fee":              fee,
		"info":             item,
	})
	_ = result
	if IsTrue(OpEq2(type_, MkString("trade"))) {
		if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
			fee = MkMap(&VarMap{
				"cost":     feeCost,
				"currency": quote,
			})
		}
		return MkArray(&VarArray{
			this.Extend(result, MkMap(&VarMap{
				"currency":  base,
				"amount":    Math.Abs(baseAmount),
				"direction": OpTrinary(OpGt(baseAmount, MkInteger(0)), MkString("in"), MkString("out")),
			})),
			this.Extend(result, MkMap(&VarMap{
				"currency":  quote,
				"amount":    Math.Abs(quoteAmount),
				"direction": OpTrinary(OpGt(quoteAmount, MkInteger(0)), MkString("in"), MkString("out")),
				"fee":       fee,
			})),
		})
	} else {
		amount := OpTrinary(OpEq2(baseAmount, MkInteger(0)), quoteAmount, baseAmount)
		_ = amount
		code := OpTrinary(OpEq2(baseAmount, MkInteger(0)), quote, base)
		_ = code
		direction := OpTrinary(OpGt(amount, MkInteger(0)), MkString("in"), MkString("out"))
		_ = direction
		if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
			fee = MkMap(&VarMap{
				"cost":     feeCost,
				"currency": code,
			})
		}
		return this.Extend(result, MkMap(&VarMap{
			"currency":  code,
			"amount":    Math.Abs(amount),
			"direction": direction,
			"fee":       fee,
		}))
	}
	return MkUndefined()
}

func (this *Coinfloor) CreateOrder(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"id": this.MarketId(symbol)})
	_ = request
	method := OpAdd(MkString("privatePostId"), this.Capitalize(side))
	_ = method
	if IsTrue(OpEq2(type_, MkString("market"))) {
		*(request).At(MkString("quantity")) = OpCopy(amount)
		OpAddAssign(&method, MkString("Market"))
	} else {
		*(request).At(MkString("price")) = OpCopy(price)
		*(request).At(MkString("amount")) = OpCopy(amount)
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	timestamp := this.Parse8601(this.SafeString(response, MkString("datetime")))
	_ = timestamp
	return MkMap(&VarMap{
		"id":            this.SafeString(response, MkString("id")),
		"clientOrderId": MkUndefined(),
		"datetime":      this.Iso8601(timestamp),
		"timestamp":     timestamp,
		"type":          type_,
		"price":         this.SafeNumber(response, MkString("price")),
		"remaining":     this.SafeNumber(response, MkString("amount")),
		"info":          response,
	})
}

func (this *Coinfloor) CancelOrder(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{
		"symbol": *(market).At(MkString("id")),
		"id":     id,
	})
	_ = request
	response := this.Call(MkString("privatePostSymbolCancelOrder"), request)
	_ = response
	if IsTrue(OpEq2(response, MkString("false"))) {
		panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")), MkString(" cancel was rejected"))))
	}
	return response
}

func (this *Coinfloor) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.Parse8601(this.SafeString(order, MkString("datetime")))
	_ = timestamp
	price := this.SafeNumber(order, MkString("price"))
	_ = price
	amount := this.SafeNumber(order, MkString("amount"))
	_ = amount
	side := OpCopy(MkUndefined())
	_ = side
	status := this.SafeString(order, MkString("status"))
	_ = status
	rawType := this.SafeString(order, MkString("type"))
	_ = rawType
	if IsTrue(OpEq2(rawType, MkString("0"))) {
		side = MkString("buy")
	} else {
		if IsTrue(OpEq2(rawType, MkString("1"))) {
			side = MkString("sell")
		}
	}
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(OpNotEq2(market, MkUndefined())) {
		symbol = *(market).At(MkString("symbol"))
	}
	id := this.SafeString(order, MkString("id"))
	_ = id
	return this.SafeOrder(MkMap(&VarMap{
		"info":               order,
		"id":                 id,
		"clientOrderId":      MkUndefined(),
		"datetime":           this.Iso8601(timestamp),
		"timestamp":          timestamp,
		"lastTradeTimestamp": MkUndefined(),
		"status":             status,
		"symbol":             symbol,
		"type":               MkString("limit"),
		"timeInForce":        MkUndefined(),
		"postOnly":           MkUndefined(),
		"side":               side,
		"price":              price,
		"stopPrice":          MkUndefined(),
		"amount":             MkUndefined(),
		"filled":             MkUndefined(),
		"remaining":          amount,
		"cost":               MkUndefined(),
		"fee":                MkUndefined(),
		"average":            MkUndefined(),
		"trades":             MkUndefined(),
	}))
}

func (this *Coinfloor) FetchOpenOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchOpenOrders() requires a symbol param"))))
	}
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"id": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("privatePostIdOpenOrders"), this.Extend(request, params))
	_ = response
	return this.ParseOrders(response, market, since, limit, MkMap(&VarMap{"status": MkString("open")}))
}

func (this *Coinfloor) HandleErrors(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpLw(code, MkInteger(400))) {
		return MkUndefined()
	}
	if IsTrue(OpEq2(response, MkUndefined())) {
		return MkUndefined()
	}
	message := this.SafeString(response, MkString("error_msg"))
	_ = message
	feedback := OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))
	_ = feedback
	this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("exact")), message, feedback)
	panic(NewExchangeError(feedback))
	return MkUndefined()
}

func (this *Coinfloor) Sign(goArgs ...*Variant) *Variant {
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
	url := OpAdd(*(*this.At(MkString("urls"))).At(MkString("api")), OpAdd(MkString("/"), this.ImplodeParams(path, params)))
	_ = url
	query := this.Omit(params, this.ExtractParams(path))
	_ = query
	if IsTrue(OpEq2(api, MkString("public"))) {
		if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
			OpAddAssign(&url, OpAdd(MkString("?"), this.Urlencode(query)))
		}
	} else {
		this.CheckRequiredCredentials()
		nonce := this.Nonce()
		_ = nonce
		body = this.Urlencode(this.Extend(MkMap(&VarMap{"nonce": nonce}), query))
		auth := OpAdd(*this.At(MkString("uid")), OpAdd(MkString("/"), OpAdd(*this.At(MkString("apiKey")), OpAdd(MkString(":"), *this.At(MkString("password"))))))
		_ = auth
		signature := this.Decode(this.StringToBase64(auth))
		_ = signature
		headers = MkMap(&VarMap{
			"Content-Type":  MkString("application/x-www-form-urlencoded"),
			"Authorization": OpAdd(MkString("Basic "), signature),
		})
	}
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}
