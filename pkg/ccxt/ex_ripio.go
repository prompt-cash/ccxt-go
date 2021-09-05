package ccxt

import ()

type Ripio struct {
	*ExchangeBase
}

var _ Exchange = (*Ripio)(nil)

func init() {
	exchange := &Ripio{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Ripio) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("ripio"),
		"name": MkString("Ripio"),
		"countries": MkArray(&VarArray{
			MkString("AR"),
			MkString("BR"),
		}),
		"rateLimit": MkInteger(50),
		"version":   MkString("v1"),
		"pro":       MkBool(true),
		"has": MkMap(&VarMap{
			"CORS":              MkBool(false),
			"cancelOrder":       MkBool(true),
			"createOrder":       MkBool(true),
			"fetchBalance":      MkBool(true),
			"fetchClosedOrders": MkBool(true),
			"fetchCurrencies":   MkBool(true),
			"fetchMyTrades":     MkBool(true),
			"fetchOpenOrders":   MkBool(true),
			"fetchOrder":        MkBool(true),
			"fetchOrderBook":    MkBool(true),
			"fetchOrders":       MkBool(true),
			"fetchTicker":       MkBool(true),
			"fetchTickers":      MkBool(true),
			"fetchTrades":       MkBool(true),
		}),
		"urls": MkMap(&VarMap{
			"logo": MkString("https://user-images.githubusercontent.com/1294454/94507548-a83d6a80-0218-11eb-9998-28b9cec54165.jpg"),
			"api": MkMap(&VarMap{
				"public":  MkString("https://api.exchange.ripio.com/api"),
				"private": MkString("https://api.exchange.ripio.com/api"),
			}),
			"www": MkString("https://exchange.ripio.com"),
			"doc": MkArray(&VarArray{
				MkString("https://exchange.ripio.com/en/api/"),
			}),
			"fees": MkString("https://exchange.ripio.com/en/fee"),
		}),
		"api": MkMap(&VarMap{
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("rate/all/"),
				MkString("rate/{pair}/"),
				MkString("orderbook/{pair}/"),
				MkString("tradehistory/{pair}/"),
				MkString("pair/"),
				MkString("currency/"),
				MkString("orderbook/{pair}/depth/"),
			})}),
			"private": MkMap(&VarMap{
				"get": MkArray(&VarArray{
					MkString("balances/exchange_balances/"),
					MkString("order/{pair}/{order_id}/"),
					MkString("order/{pair}/"),
					MkString("trade/{pair}/"),
				}),
				"post": MkArray(&VarArray{
					MkString("order/{pair}/"),
					MkString("order/{pair}/{order_id}/cancel/"),
				}),
			}),
		}),
		"fees": MkMap(&VarMap{"trading": MkMap(&VarMap{
			"tierBased":  MkBool(true),
			"percentage": MkBool(true),
			"taker":      OpDiv(MkNumber(0.0), MkInteger(100)),
			"maker":      OpDiv(MkNumber(0.0), MkInteger(100)),
		})}),
		"precisionMode": TICK_SIZE,
		"requiredCredentials": MkMap(&VarMap{
			"apiKey": MkBool(true),
			"secret": MkBool(false),
		}),
		"exceptions": MkMap(&VarMap{
			"exact": MkMap(&VarMap{}),
			"broad": MkMap(&VarMap{
				"Authentication credentials were not provided": AuthenticationError,
				"Disabled pair":                BadSymbol,
				"Invalid order type":           InvalidOrder,
				"Your balance is not enough":   InsufficientFunds,
				"Order couldn't be created":    ExchangeError,
				"not found":                    OrderNotFound,
				"Invalid pair":                 BadSymbol,
				"amount must be a number":      BadRequest,
				"Total must be at least":       InvalidOrder,
				"Account not found":            BadRequest,
				"Wrong password provided":      AuthenticationError,
				"User tokens limit":            DDoSProtection,
				"Something unexpected ocurred": ExchangeError,
				"account_balance":              BadRequest,
			}),
		}),
	}))
}

func (this *Ripio) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetPair"), params)
	_ = response
	result := MkArray(&VarArray{})
	_ = result
	results := this.SafeValue(response, MkString("results"), MkArray(&VarArray{}))
	_ = results
	for i := MkInteger(0); IsTrue(OpLw(i, results.Length)); OpIncr(&i) {
		market := *(results).At(i)
		_ = market
		baseId := this.SafeString(market, MkString("base"))
		_ = baseId
		quoteId := this.SafeString(market, MkString("quote"))
		_ = quoteId
		id := this.SafeString(market, MkString("symbol"))
		_ = id
		base := this.SafeCurrencyCode(baseId)
		_ = base
		quote := this.SafeCurrencyCode(quoteId)
		_ = quote
		symbol := OpAdd(base, OpAdd(MkString("/"), quote))
		_ = symbol
		precision := MkMap(&VarMap{
			"amount": this.SafeNumber(market, MkString("min_amount")),
			"price":  this.SafeNumber(market, MkString("price_tick")),
		})
		_ = precision
		limits := MkMap(&VarMap{
			"amount": MkMap(&VarMap{
				"min": this.SafeNumber(market, MkString("min_amount")),
				"max": MkUndefined(),
			}),
			"price": MkMap(&VarMap{
				"min": MkUndefined(),
				"max": MkUndefined(),
			}),
			"cost": MkMap(&VarMap{
				"min": this.SafeNumber(market, MkString("min_value")),
				"max": MkUndefined(),
			}),
		})
		_ = limits
		active := this.SafeValue(market, MkString("enabled"), MkBool(true))
		_ = active
		fees := this.SafeValue(market, MkString("fees"), MkArray(&VarArray{}))
		_ = fees
		firstFee := this.SafeValue(fees, MkInteger(0), MkMap(&VarMap{}))
		_ = firstFee
		maker := this.SafeNumber(firstFee, MkString("maker_fee"), MkNumber(0.0))
		_ = maker
		taker := this.SafeNumber(firstFee, MkString("taker_fee"), MkNumber(0.0))
		_ = taker
		result.Push(MkMap(&VarMap{
			"id":        id,
			"symbol":    symbol,
			"base":      base,
			"quote":     quote,
			"baseId":    baseId,
			"quoteId":   quoteId,
			"precision": precision,
			"maker":     maker,
			"taker":     taker,
			"limits":    limits,
			"info":      market,
			"active":    active,
		}))
	}
	return result
}

func (this *Ripio) FetchCurrencies(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetCurrency"), params)
	_ = response
	results := this.SafeValue(response, MkString("results"), MkArray(&VarArray{}))
	_ = results
	result := MkMap(&VarMap{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, results.Length)); OpIncr(&i) {
		currency := *(results).At(i)
		_ = currency
		id := this.SafeString(currency, MkString("currency"))
		_ = id
		code := this.SafeCurrencyCode(id)
		_ = code
		name := this.SafeString(currency, MkString("name"))
		_ = name
		active := this.SafeValue(currency, MkString("enabled"), MkBool(true))
		_ = active
		precision := this.SafeInteger(currency, MkString("decimal_places"))
		_ = precision
		*(result).At(code) = MkMap(&VarMap{
			"id":        id,
			"code":      code,
			"name":      name,
			"info":      currency,
			"active":    active,
			"fee":       MkUndefined(),
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

func (this *Ripio) ParseTicker(goArgs ...*Variant) *Variant {
	ticker := GoGetArg(goArgs, 0, MkUndefined())
	_ = ticker
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.Parse8601(this.SafeString(ticker, MkString("created_at")))
	_ = timestamp
	marketId := this.SafeString(ticker, MkString("pair"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market)
	_ = symbol
	last := this.SafeNumber(ticker, MkString("last_price"))
	_ = last
	average := this.SafeNumber(ticker, MkString("avg"))
	_ = average
	return MkMap(&VarMap{
		"symbol":        symbol,
		"timestamp":     timestamp,
		"datetime":      this.Iso8601(timestamp),
		"high":          this.SafeNumber(ticker, MkString("high")),
		"low":           this.SafeNumber(ticker, MkString("low")),
		"bid":           this.SafeNumber(ticker, MkString("bid")),
		"bidVolume":     this.SafeNumber(ticker, MkString("bid_volume")),
		"ask":           this.SafeNumber(ticker, MkString("ask")),
		"askVolume":     this.SafeNumber(ticker, MkString("ask_volume")),
		"vwap":          MkUndefined(),
		"open":          MkUndefined(),
		"close":         last,
		"last":          last,
		"previousClose": MkUndefined(),
		"change":        MkUndefined(),
		"percentage":    MkUndefined(),
		"average":       average,
		"baseVolume":    MkUndefined(),
		"quoteVolume":   MkUndefined(),
		"info":          ticker,
	})
}

func (this *Ripio) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"pair": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("publicGetRatePair"), this.Extend(request, params))
	_ = response
	return this.ParseTicker(response, market)
}

func (this *Ripio) FetchTickers(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("publicGetRateAll"), params)
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

func (this *Ripio) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"pair": this.MarketId(symbol)})
	_ = request
	response := this.Call(MkString("publicGetOrderbookPair"), this.Extend(request, params))
	_ = response
	orderbook := this.ParseOrderBook(response, symbol, MkUndefined(), MkString("buy"), MkString("sell"), MkString("price"), MkString("amount"))
	_ = orderbook
	*(orderbook).At(MkString("nonce")) = this.SafeInteger(response, MkString("updated_id"))
	return orderbook
}

func (this *Ripio) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	id := this.SafeString(trade, MkString("id"))
	_ = id
	timestamp := this.SafeInteger(trade, MkString("timestamp"))
	_ = timestamp
	timestamp = this.SafeTimestamp(trade, MkString("created_at"), timestamp)
	side := this.SafeString(trade, MkString("side"))
	_ = side
	takerSide := this.SafeString(trade, MkString("taker_side"))
	_ = takerSide
	takerOrMaker := OpTrinary(OpEq2(takerSide, side), MkString("taker"), MkString("maker"))
	_ = takerOrMaker
	if IsTrue(OpNotEq2(side, MkUndefined())) {
		side = side.ToLowerCase()
	}
	priceString := this.SafeString2(trade, MkString("price"), MkString("match_price"))
	_ = priceString
	amountString := this.SafeString2(trade, MkString("amount"), MkString("exchanged"))
	_ = amountString
	price := this.ParseNumber(priceString)
	_ = price
	amount := this.ParseNumber(amountString)
	_ = amount
	cost := this.ParseNumber(Precise.StringMul(priceString, amountString))
	_ = cost
	marketId := this.SafeString(trade, MkString("pair"))
	_ = marketId
	market = this.SafeMarket(marketId, market)
	feeCost := this.SafeNumber(trade, OpAdd(takerOrMaker, MkString("_fee")))
	_ = feeCost
	orderId := this.SafeString(trade, takerOrMaker)
	_ = orderId
	fee := OpCopy(MkUndefined())
	_ = fee
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		fee = MkMap(&VarMap{
			"cost":     feeCost,
			"currency": OpTrinary(OpEq2(side, MkString("buy")), *(market).At(MkString("base")), *(market).At(MkString("quote"))),
		})
	}
	return MkMap(&VarMap{
		"id":           id,
		"order":        orderId,
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       *(market).At(MkString("symbol")),
		"type":         MkUndefined(),
		"side":         side,
		"price":        price,
		"amount":       amount,
		"cost":         cost,
		"takerOrMaker": takerOrMaker,
		"fee":          fee,
		"info":         trade,
	})
}

func (this *Ripio) FetchTrades(goArgs ...*Variant) *Variant {
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
	response := this.Call(MkString("publicGetTradehistoryPair"), this.Extend(request, params))
	_ = response
	return this.ParseTrades(response, market, since, limit)
}

func (this *Ripio) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privateGetBalancesExchangeBalances"), params)
	_ = response
	result := MkMap(&VarMap{"info": response})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
		balance := *(response).At(i)
		_ = balance
		currencyId := this.SafeString(balance, MkString("symbol"))
		_ = currencyId
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		account := this.Account()
		_ = account
		*(account).At(MkString("free")) = this.SafeString(balance, MkString("available"))
		*(account).At(MkString("used")) = this.SafeString(balance, MkString("locked"))
		*(result).At(code) = OpCopy(account)
	}
	return this.ParseBalance(result)
}

func (this *Ripio) CreateOrder(goArgs ...*Variant) *Variant {
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
	uppercaseType := type_.ToUpperCase()
	_ = uppercaseType
	uppercaseSide := side.ToUpperCase()
	_ = uppercaseSide
	request := MkMap(&VarMap{
		"pair":       *(market).At(MkString("id")),
		"order_type": uppercaseType,
		"side":       uppercaseSide,
		"amount":     this.AmountToPrecision(symbol, amount),
	})
	_ = request
	if IsTrue(OpEq2(uppercaseType, MkString("LIMIT"))) {
		*(request).At(MkString("limit_price")) = this.PriceToPrecision(symbol, price)
	}
	response := this.Call(MkString("privatePostOrderPair"), this.Extend(request, params))
	_ = response
	return this.ParseOrder(response, market)
}

func (this *Ripio) CancelOrder(goArgs ...*Variant) *Variant {
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
		"pair":     *(market).At(MkString("id")),
		"order_id": id,
	})
	_ = request
	response := this.Call(MkString("privatePostOrderPairOrderIdCancel"), this.Extend(request, params))
	_ = response
	return this.ParseOrder(response, market)
}

func (this *Ripio) FetchOrder(goArgs ...*Variant) *Variant {
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
		"pair":     *(market).At(MkString("id")),
		"order_id": id,
	})
	_ = request
	response := this.Call(MkString("privateGetOrderPairOrderId"), this.Extend(request, params))
	_ = response
	return this.ParseOrder(response, market)
}

func (this *Ripio) FetchOrders(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"pair": *(market).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("offset")) = OpCopy(limit)
	}
	response := this.Call(MkString("privateGetOrderPair"), this.Extend(request, params))
	_ = response
	results := this.SafeValue(response, MkString("results"), MkMap(&VarMap{}))
	_ = results
	data := this.SafeValue(results, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseOrders(data, market, since, limit)
}

func (this *Ripio) FetchOpenOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	request := MkMap(&VarMap{"status": MkString("OPEN,PART")})
	_ = request
	return this.FetchOrders(symbol, since, limit, this.Extend(request, params))
}

func (this *Ripio) FetchClosedOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	request := MkMap(&VarMap{"status": MkString("CLOS,CANC,COMP")})
	_ = request
	return this.FetchOrders(symbol, since, limit, this.Extend(request, params))
}

func (this *Ripio) ParseOrderStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"OPEN": MkString("open"),
		"PART": MkString("open"),
		"CLOS": MkString("canceled"),
		"CANC": MkString("canceled"),
		"COMP": MkString("closed"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Ripio) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	id := this.SafeString(order, MkString("order_id"))
	_ = id
	amount := this.SafeNumber(order, MkString("amount"))
	_ = amount
	cost := this.SafeNumber(order, MkString("notional"))
	_ = cost
	type_ := this.SafeStringLower(order, MkString("order_type"))
	_ = type_
	priceField := OpTrinary(OpEq2(type_, MkString("market")), MkString("fill_price"), MkString("limit_price"))
	_ = priceField
	price := this.SafeNumber(order, priceField)
	_ = price
	side := this.SafeStringLower(order, MkString("side"))
	_ = side
	status := this.ParseOrderStatus(this.SafeString(order, MkString("status")))
	_ = status
	timestamp := this.SafeTimestamp(order, MkString("created_at"))
	_ = timestamp
	average := this.SafeValue(order, MkString("fill_price"))
	_ = average
	filled := this.SafeNumber(order, MkString("filled"))
	_ = filled
	remaining := OpCopy(MkUndefined())
	_ = remaining
	fills := this.SafeValue(order, MkString("fills"))
	_ = fills
	trades := OpCopy(MkUndefined())
	_ = trades
	lastTradeTimestamp := OpCopy(MkUndefined())
	_ = lastTradeTimestamp
	if IsTrue(OpNotEq2(fills, MkUndefined())) {
		numFills := OpCopy(fills.Length)
		_ = numFills
		if IsTrue(OpGt(numFills, MkInteger(0))) {
			filled = MkInteger(0)
			cost = MkInteger(0)
			trades = this.ParseTrades(fills, market, MkUndefined(), MkUndefined(), MkMap(&VarMap{
				"order": id,
				"side":  side,
			}))
			for i := MkInteger(0); IsTrue(OpLw(i, trades.Length)); OpIncr(&i) {
				trade := *(trades).At(i)
				_ = trade
				filled = this.Sum(*(trade).At(MkString("amount")), filled)
				cost = this.Sum(*(trade).At(MkString("cost")), cost)
				lastTradeTimestamp = *(trade).At(MkString("timestamp"))
			}
			if IsTrue(OpAnd(OpEq2(average, MkUndefined()), OpGt(filled, MkInteger(0)))) {
				average = OpDiv(cost, filled)
			}
		}
	}
	if IsTrue(OpNotEq2(filled, MkUndefined())) {
		if IsTrue(OpAnd(OpEq2(cost, MkUndefined()), OpNotEq2(price, MkUndefined()))) {
			cost = OpMulti(price, filled)
		}
		if IsTrue(OpNotEq2(amount, MkUndefined())) {
			remaining = Math.Max(MkInteger(0), OpSub(amount, filled))
		}
	}
	marketId := this.SafeString(order, MkString("pair"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market, MkString("_"))
	_ = symbol
	stopPrice := this.SafeNumber(order, MkString("stop_price"))
	_ = stopPrice
	return MkMap(&VarMap{
		"id":                 id,
		"clientOrderId":      MkUndefined(),
		"info":               order,
		"timestamp":          timestamp,
		"datetime":           this.Iso8601(timestamp),
		"lastTradeTimestamp": lastTradeTimestamp,
		"symbol":             symbol,
		"type":               type_,
		"timeInForce":        MkUndefined(),
		"postOnly":           MkUndefined(),
		"side":               side,
		"price":              price,
		"stopPrice":          stopPrice,
		"amount":             amount,
		"cost":               cost,
		"average":            average,
		"filled":             filled,
		"remaining":          remaining,
		"status":             status,
		"fee":                MkUndefined(),
		"trades":             trades,
	})
}

func (this *Ripio) FetchMyTrades(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"pair": *(market).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("privateGetTradePair"), this.Extend(request, params))
	_ = response
	results := this.SafeValue(response, MkString("results"), MkMap(&VarMap{}))
	_ = results
	data := this.SafeValue(results, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseTrades(data, market, since, limit)
}

func (this *Ripio) Sign(goArgs ...*Variant) *Variant {
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
	request := OpAdd(MkString("/"), OpAdd(*this.At(MkString("version")), OpAdd(MkString("/"), this.ImplodeParams(path, params))))
	_ = request
	url := OpAdd(*(*(*this.At(MkString("urls"))).At(MkString("api"))).At(api), request)
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
			if IsTrue(OpEq2(method, MkString("POST"))) {
				body = this.Json(query)
			} else {
				if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
					OpAddAssign(&url, OpAdd(MkString("?"), this.Urlencode(query)))
				}
			}
			headers = MkMap(&VarMap{
				"Content-Type":  MkString("application/json"),
				"Authorization": OpAdd(MkString("Bearer "), *this.At(MkString("apiKey"))),
			})
		}
	}
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Ripio) HandleErrors(goArgs ...*Variant) *Variant {
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
	detail := this.SafeString(response, MkString("detail"))
	_ = detail
	if IsTrue(OpNotEq2(detail, MkUndefined())) {
		feedback := OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))
		_ = feedback
		this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("broad")), detail, feedback)
	}
	errors := this.SafeValue(response, MkString("errors"))
	_ = errors
	if IsTrue(OpNotEq2(errors, MkUndefined())) {
		feedback := OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))
		_ = feedback
		keys := GoGetKeys(errors)
		_ = keys
		for i := MkInteger(0); IsTrue(OpLw(i, keys.Length)); OpIncr(&i) {
			key := *(keys).At(i)
			_ = key
			error := this.SafeValue(errors, key, MkArray(&VarArray{}))
			_ = error
			message := this.SafeString(error, MkInteger(0))
			_ = message
			this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("broad")), message, feedback)
		}
		panic(NewExchangeError(feedback))
	}
	return MkUndefined()
}
