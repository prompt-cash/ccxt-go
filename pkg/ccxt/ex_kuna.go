package ccxt

import ()

type Kuna struct {
	*ExchangeBase
}

var _ Exchange = (*Kuna)(nil)

func init() {
	exchange := &Kuna{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Kuna) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("kuna"),
		"name": MkString("Kuna"),
		"countries": MkArray(&VarArray{
			MkString("UA"),
		}),
		"rateLimit": MkInteger(1000),
		"version":   MkString("v2"),
		"has": MkMap(&VarMap{
			"CORS":            MkBool(false),
			"cancelOrder":     MkBool(true),
			"createOrder":     MkBool(true),
			"fetchBalance":    MkBool(true),
			"fetchMarkets":    MkBool(true),
			"fetchMyTrades":   MkBool(true),
			"fetchOHLCV":      MkString("emulated"),
			"fetchOpenOrders": MkBool(true),
			"fetchOrder":      MkBool(true),
			"fetchOrderBook":  MkBool(true),
			"fetchTicker":     MkBool(true),
			"fetchTickers":    MkBool(true),
			"fetchTime":       MkBool(true),
			"fetchTrades":     MkBool(true),
			"withdraw":        MkBool(false),
		}),
		"timeframes": MkUndefined(),
		"urls": MkMap(&VarMap{
			"extension": MkString(".json"),
			"referral":  MkString("https://kuna.io?r=kunaid-gvfihe8az7o4"),
			"logo":      MkString("https://user-images.githubusercontent.com/51840849/87153927-f0578b80-c2c0-11ea-84b6-74612568e9e1.jpg"),
			"api":       MkString("https://kuna.io"),
			"www":       MkString("https://kuna.io"),
			"doc":       MkString("https://kuna.io/documents/api"),
			"fees":      MkString("https://kuna.io/documents/api"),
		}),
		"api": MkMap(&VarMap{
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("depth"),
				MkString("k_with_pending_trades"),
				MkString("k"),
				MkString("markets"),
				MkString("order_book"),
				MkString("order_book/{market}"),
				MkString("tickers"),
				MkString("tickers/{market}"),
				MkString("timestamp"),
				MkString("trades"),
				MkString("trades/{market}"),
			})}),
			"private": MkMap(&VarMap{
				"get": MkArray(&VarArray{
					MkString("members/me"),
					MkString("deposits"),
					MkString("deposit"),
					MkString("deposit_address"),
					MkString("orders"),
					MkString("order"),
					MkString("trades/my"),
					MkString("withdraws"),
					MkString("withdraw"),
				}),
				"post": MkArray(&VarArray{
					MkString("orders"),
					MkString("orders/multi"),
					MkString("orders/clear"),
					MkString("order/delete"),
					MkString("withdraw"),
				}),
			}),
		}),
		"fees": MkMap(&VarMap{
			"trading": MkMap(&VarMap{
				"tierBased":  MkBool(false),
				"percentage": MkBool(true),
				"taker":      OpDiv(MkNumber(0.25), MkInteger(100)),
				"maker":      OpDiv(MkNumber(0.25), MkInteger(100)),
			}),
			"funding": MkMap(&VarMap{
				"withdraw": MkMap(&VarMap{
					"UAH":   MkString("1%"),
					"BTC":   MkNumber(0.001),
					"BCH":   MkNumber(0.001),
					"ETH":   MkNumber(0.01),
					"WAVES": MkNumber(0.01),
					"GOL":   MkNumber(0.0),
					"GBG":   MkNumber(0.0),
				}),
				"deposit": MkMap(&VarMap{}),
			}),
		}),
		"commonCurrencies": MkMap(&VarMap{"PLA": MkString("Plair")}),
		"exceptions": MkMap(&VarMap{
			"2002": InsufficientFunds,
			"2003": OrderNotFound,
		}),
	}))
}

func (this *Kuna) FetchTime(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetTimestamp"), params)
	_ = response
	return OpMulti(response, MkInteger(1000))
}

func (this *Kuna) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	quotes := MkArray(&VarArray{
		MkString("btc"),
		MkString("rub"),
		MkString("uah"),
		MkString("usd"),
		MkString("usdt"),
		MkString("usdc"),
	})
	_ = quotes
	markets := MkArray(&VarArray{})
	_ = markets
	response := this.Call(MkString("publicGetTickers"), params)
	_ = response
	ids := GoGetKeys(response)
	_ = ids
	for i := MkInteger(0); IsTrue(OpLw(i, ids.Length)); OpIncr(&i) {
		id := *(ids).At(i)
		_ = id
		for j := MkInteger(0); IsTrue(OpLw(j, quotes.Length)); OpIncr(&j) {
			quoteId := *(quotes).At(j)
			_ = quoteId
			slicedId := id.Slice(MkInteger(1))
			_ = slicedId
			index := slicedId.IndexOf(quoteId)
			_ = index
			slice := slicedId.Slice(index)
			_ = slice
			if IsTrue(OpAnd(OpGt(index, MkInteger(0)), OpEq2(slice, quoteId))) {
				baseId := OpAdd(*(id).At(MkInteger(0)), slicedId.Replace(quoteId, MkString("")))
				_ = baseId
				base := this.SafeCurrencyCode(baseId)
				_ = base
				quote := this.SafeCurrencyCode(quoteId)
				_ = quote
				symbol := OpAdd(base, OpAdd(MkString("/"), quote))
				_ = symbol
				markets.Push(MkMap(&VarMap{
					"id":      id,
					"symbol":  symbol,
					"base":    base,
					"quote":   quote,
					"baseId":  baseId,
					"quoteId": quoteId,
					"precision": MkMap(&VarMap{
						"amount": MkUndefined(),
						"price":  MkUndefined(),
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
						"cost": MkMap(&VarMap{
							"min": MkUndefined(),
							"max": MkUndefined(),
						}),
					}),
					"active": MkUndefined(),
					"info":   MkUndefined(),
				}))
			}
		}
	}
	return markets
}

func (this *Kuna) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"market": *(market).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	orderbook := this.Call(MkString("publicGetDepth"), this.Extend(request, params))
	_ = orderbook
	timestamp := this.SafeTimestamp(orderbook, MkString("timestamp"))
	_ = timestamp
	return this.ParseOrderBook(orderbook, symbol, timestamp)
}

func (this *Kuna) ParseTicker(goArgs ...*Variant) *Variant {
	ticker := GoGetArg(goArgs, 0, MkUndefined())
	_ = ticker
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.SafeTimestamp(ticker, MkString("at"))
	_ = timestamp
	ticker = *(ticker).At(MkString("ticker"))
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(market) {
		symbol = *(market).At(MkString("symbol"))
	}
	last := this.SafeNumber(ticker, MkString("last"))
	_ = last
	return MkMap(&VarMap{
		"symbol":        symbol,
		"timestamp":     timestamp,
		"datetime":      this.Iso8601(timestamp),
		"high":          this.SafeNumber(ticker, MkString("high")),
		"low":           this.SafeNumber(ticker, MkString("low")),
		"bid":           this.SafeNumber(ticker, MkString("buy")),
		"bidVolume":     MkUndefined(),
		"ask":           this.SafeNumber(ticker, MkString("sell")),
		"askVolume":     MkUndefined(),
		"vwap":          MkUndefined(),
		"open":          this.SafeNumber(ticker, MkString("open")),
		"close":         last,
		"last":          last,
		"previousClose": MkUndefined(),
		"change":        MkUndefined(),
		"percentage":    MkUndefined(),
		"average":       MkUndefined(),
		"baseVolume":    this.SafeNumber(ticker, MkString("vol")),
		"quoteVolume":   MkUndefined(),
		"info":          ticker,
	})
}

func (this *Kuna) FetchTickers(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("publicGetTickers"), params)
	_ = response
	ids := GoGetKeys(response)
	_ = ids
	result := MkMap(&VarMap{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, ids.Length)); OpIncr(&i) {
		id := *(ids).At(i)
		_ = id
		market := OpCopy(MkUndefined())
		_ = market
		symbol := OpCopy(id)
		_ = symbol
		if IsTrue(OpHasMember(id, *this.At(MkString("markets_by_id")))) {
			market = *(*this.At(MkString("markets_by_id"))).At(id)
			symbol = *(market).At(MkString("symbol"))
		} else {
			base := id.Slice(MkInteger(0), MkInteger(3))
			_ = base
			quote := id.Slice(MkInteger(3), MkInteger(6))
			_ = quote
			base = base.ToUpperCase()
			quote = quote.ToUpperCase()
			base = this.SafeCurrencyCode(base)
			quote = this.SafeCurrencyCode(quote)
			symbol = OpAdd(base, OpAdd(MkString("/"), quote))
		}
		*(result).At(symbol) = this.ParseTicker(*(response).At(id), market)
	}
	return this.FilterByArray(result, MkString("symbol"), symbols)
}

func (this *Kuna) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"market": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("publicGetTickersMarket"), this.Extend(request, params))
	_ = response
	return this.ParseTicker(response, market)
}

func (this *Kuna) FetchL3OrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	return this.FetchOrderBook(symbol, limit, params)
}

func (this *Kuna) FetchTrades(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"market": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("publicGetTrades"), this.Extend(request, params))
	_ = response
	return this.ParseTrades(response, market, since, limit)
}

func (this *Kuna) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.Parse8601(this.SafeString(trade, MkString("created_at")))
	_ = timestamp
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(market) {
		symbol = *(market).At(MkString("symbol"))
	}
	side := this.SafeString2(trade, MkString("side"), MkString("trend"))
	_ = side
	if IsTrue(OpNotEq2(side, MkUndefined())) {
		sideMap := MkMap(&VarMap{
			"ask": MkString("sell"),
			"bid": MkString("buy"),
		})
		_ = sideMap
		side = this.SafeString(sideMap, side, side)
	}
	priceString := this.SafeString(trade, MkString("price"))
	_ = priceString
	amountString := this.SafeString(trade, MkString("volume"))
	_ = amountString
	price := this.ParseNumber(priceString)
	_ = price
	amount := this.ParseNumber(amountString)
	_ = amount
	cost := this.SafeNumber(trade, MkString("funds"))
	_ = cost
	if IsTrue(OpEq2(cost, MkUndefined())) {
		cost = this.ParseNumber(Precise.StringMul(priceString, amountString))
	}
	orderId := this.SafeString(trade, MkString("order_id"))
	_ = orderId
	id := this.SafeString(trade, MkString("id"))
	_ = id
	return MkMap(&VarMap{
		"id":           id,
		"info":         trade,
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       symbol,
		"type":         MkUndefined(),
		"side":         side,
		"order":        orderId,
		"takerOrMaker": MkUndefined(),
		"price":        price,
		"amount":       amount,
		"cost":         cost,
		"fee":          MkUndefined(),
	})
}

func (this *Kuna) FetchOHLCV(goArgs ...*Variant) *Variant {
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
	trades := this.FetchTrades(symbol, since, limit, params)
	_ = trades
	ohlcvc := this.BuildOHLCVC(trades, timeframe, since, limit)
	_ = ohlcvc
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, ohlcvc.Length)); OpIncr(&i) {
		ohlcv := *(ohlcvc).At(i)
		_ = ohlcv
		result.Push(MkArray(&VarArray{
			*(ohlcv).At(MkInteger(0)),
			*(ohlcv).At(MkInteger(1)),
			*(ohlcv).At(MkInteger(2)),
			*(ohlcv).At(MkInteger(3)),
			*(ohlcv).At(MkInteger(4)),
			*(ohlcv).At(MkInteger(5)),
		}))
	}
	return result
}

func (this *Kuna) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privateGetMembersMe"), params)
	_ = response
	balances := this.SafeValue(response, MkString("accounts"))
	_ = balances
	result := MkMap(&VarMap{"info": balances})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, balances.Length)); OpIncr(&i) {
		balance := *(balances).At(i)
		_ = balance
		currencyId := this.SafeString(balance, MkString("currency"))
		_ = currencyId
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		account := this.Account()
		_ = account
		*(account).At(MkString("free")) = this.SafeString(balance, MkString("balance"))
		*(account).At(MkString("used")) = this.SafeString(balance, MkString("locked"))
		*(result).At(code) = OpCopy(account)
	}
	return this.ParseBalance(result)
}

func (this *Kuna) CreateOrder(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{
		"market":   this.MarketId(symbol),
		"side":     side,
		"volume":   amount.ToString(),
		"ord_type": type_,
	})
	_ = request
	if IsTrue(OpEq2(type_, MkString("limit"))) {
		*(request).At(MkString("price")) = price.ToString()
	}
	response := this.Call(MkString("privatePostOrders"), this.Extend(request, params))
	_ = response
	marketId := this.SafeValue(response, MkString("market"))
	_ = marketId
	market := this.SafeValue(*this.At(MkString("markets_by_id")), marketId)
	_ = market
	return this.ParseOrder(response, market)
}

func (this *Kuna) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"id": id})
	_ = request
	response := this.Call(MkString("privatePostOrderDelete"), this.Extend(request, params))
	_ = response
	order := this.ParseOrder(response)
	_ = order
	status := *(order).At(MkString("status"))
	_ = status
	if IsTrue(OpOr(OpEq2(status, MkString("closed")), OpEq2(status, MkString("canceled")))) {
		panic(NewOrderNotFound(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), this.Json(order)))))
	}
	return order
}

func (this *Kuna) ParseOrderStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"done":   MkString("closed"),
		"wait":   MkString("open"),
		"cancel": MkString("canceled"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Kuna) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	marketId := this.SafeString(order, MkString("market"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market)
	_ = symbol
	timestamp := this.Parse8601(this.SafeString(order, MkString("created_at")))
	_ = timestamp
	status := this.ParseOrderStatus(this.SafeString(order, MkString("state")))
	_ = status
	type_ := this.SafeString(order, MkString("type"))
	_ = type_
	side := this.SafeString(order, MkString("side"))
	_ = side
	id := this.SafeString(order, MkString("id"))
	_ = id
	return this.SafeOrder(MkMap(&VarMap{
		"id":                 id,
		"clientOrderId":      MkUndefined(),
		"timestamp":          timestamp,
		"datetime":           this.Iso8601(timestamp),
		"lastTradeTimestamp": MkUndefined(),
		"status":             status,
		"symbol":             symbol,
		"type":               type_,
		"timeInForce":        MkUndefined(),
		"postOnly":           MkUndefined(),
		"side":               side,
		"price":              this.SafeNumber(order, MkString("price")),
		"stopPrice":          MkUndefined(),
		"amount":             this.SafeNumber(order, MkString("volume")),
		"filled":             this.SafeNumber(order, MkString("executed_volume")),
		"remaining":          this.SafeNumber(order, MkString("remaining_volume")),
		"trades":             MkUndefined(),
		"fee":                MkUndefined(),
		"info":               order,
		"cost":               MkUndefined(),
		"average":            MkUndefined(),
	}))
}

func (this *Kuna) FetchOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"id": parseInt(id)})
	_ = request
	response := this.Call(MkString("privateGetOrder"), this.Extend(request, params))
	_ = response
	return this.ParseOrder(response)
}

func (this *Kuna) FetchOpenOrders(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"market": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("privateGetOrders"), this.Extend(request, params))
	_ = response
	return this.ParseOrders(response, market, since, limit)
}

func (this *Kuna) FetchMyTrades(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"market": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("privateGetTradesMy"), this.Extend(request, params))
	_ = response
	return this.ParseTrades(response, market, since, limit)
}

func (this *Kuna) Nonce(goArgs ...*Variant) *Variant {
	return this.Milliseconds()
}

func (this *Kuna) EncodeParams(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkUndefined())
	_ = params
	if IsTrue(OpHasMember(MkString("orders"), params)) {
		orders := *(params).At(MkString("orders"))
		_ = orders
		query := this.Urlencode(this.Keysort(this.Omit(params, MkString("orders"))))
		_ = query
		for i := MkInteger(0); IsTrue(OpLw(i, orders.Length)); OpIncr(&i) {
			order := *(orders).At(i)
			_ = order
			keys := GoGetKeys(order)
			_ = keys
			for k := MkInteger(0); IsTrue(OpLw(k, keys.Length)); OpIncr(&k) {
				key := *(keys).At(k)
				_ = key
				value := *(order).At(key)
				_ = value
				OpAddAssign(&query, OpAdd(MkString("&orders%5B%5D%5B"), OpAdd(key, OpAdd(MkString("%5D="), value.ToString()))))
			}
		}
		return query
	}
	return this.Urlencode(this.Keysort(params))
}

func (this *Kuna) Sign(goArgs ...*Variant) *Variant {
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
	request := OpAdd(MkString("/api/"), OpAdd(*this.At(MkString("version")), OpAdd(MkString("/"), this.ImplodeParams(path, params))))
	_ = request
	if IsTrue(OpHasMember(MkString("extension"), *this.At(MkString("urls")))) {
		OpAddAssign(&request, *(*this.At(MkString("urls"))).At(MkString("extension")))
	}
	query := this.Omit(params, this.ExtractParams(path))
	_ = query
	url := OpAdd(*(*this.At(MkString("urls"))).At(MkString("api")), request)
	_ = url
	if IsTrue(OpEq2(api, MkString("public"))) {
		if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
			OpAddAssign(&url, OpAdd(MkString("?"), this.Urlencode(query)))
		}
	} else {
		this.CheckRequiredCredentials()
		nonce := (this.Nonce()).Call(MkString("toString"))
		_ = nonce
		query := this.EncodeParams(this.Extend(MkMap(&VarMap{
			"access_key": *this.At(MkString("apiKey")),
			"tonce":      nonce,
		}), params))
		_ = query
		auth := OpAdd(method, OpAdd(MkString("|"), OpAdd(request, OpAdd(MkString("|"), query))))
		_ = auth
		signed := this.Hmac(this.Encode(auth), this.Encode(*this.At(MkString("secret"))))
		_ = signed
		suffix := OpAdd(query, OpAdd(MkString("&signature="), signed))
		_ = suffix
		if IsTrue(OpEq2(method, MkString("GET"))) {
			OpAddAssign(&url, OpAdd(MkString("?"), suffix))
		} else {
			body = OpCopy(suffix)
			headers = MkMap(&VarMap{"Content-Type": MkString("application/x-www-form-urlencoded")})
		}
	}
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Kuna) HandleErrors(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpEq2(code, MkInteger(400))) {
		error := this.SafeValue(response, MkString("error"))
		_ = error
		errorCode := this.SafeString(error, MkString("code"))
		_ = errorCode
		feedback := OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), this.Json(response)))
		_ = feedback
		this.ThrowExactlyMatchedException(*this.At(MkString("exceptions")), errorCode, feedback)
	}
	return MkUndefined()
}
