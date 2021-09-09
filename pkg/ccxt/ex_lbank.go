package ccxt

import ()

type Lbank struct {
	*ExchangeBase
}

var _ Exchange = (*Lbank)(nil)

func init() {
	exchange := &Lbank{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Lbank) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("lbank"),
		"name": MkString("LBank"),
		"countries": MkArray(&VarArray{
			MkString("CN"),
		}),
		"version": MkString("v1"),
		"has": MkMap(&VarMap{
			"cancelOrder":       MkBool(true),
			"createOrder":       MkBool(true),
			"fetchBalance":      MkBool(true),
			"fetchClosedOrders": MkBool(true),
			"fetchMarkets":      MkBool(true),
			"fetchOHLCV":        MkBool(true),
			"fetchOpenOrders":   MkBool(false),
			"fetchOrder":        MkBool(true),
			"fetchOrderBook":    MkBool(true),
			"fetchOrders":       MkBool(true),
			"fetchTicker":       MkBool(true),
			"fetchTickers":      MkBool(true),
			"fetchTrades":       MkBool(true),
			"withdraw":          MkBool(true),
		}),
		"timeframes": MkMap(&VarMap{
			"1m":  MkString("minute1"),
			"5m":  MkString("minute5"),
			"15m": MkString("minute15"),
			"30m": MkString("minute30"),
			"1h":  MkString("hour1"),
			"2h":  MkString("hour2"),
			"4h":  MkString("hour4"),
			"6h":  MkString("hour6"),
			"8h":  MkString("hour8"),
			"12h": MkString("hour12"),
			"1d":  MkString("day1"),
			"1w":  MkString("week1"),
		}),
		"urls": MkMap(&VarMap{
			"logo":     MkString("https://user-images.githubusercontent.com/1294454/38063602-9605e28a-3302-11e8-81be-64b1e53c4cfb.jpg"),
			"api":      MkString("https://api.lbank.info"),
			"www":      MkString("https://www.lbank.info"),
			"doc":      MkString("https://github.com/LBank-exchange/lbank-official-api-docs"),
			"fees":     MkString("https://lbankinfo.zendesk.com/hc/en-gb/articles/360012072873-Trading-Fees"),
			"referral": MkString("https://www.lbex.io/invite?icode=7QCY"),
		}),
		"api": MkMap(&VarMap{
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("currencyPairs"),
				MkString("ticker"),
				MkString("depth"),
				MkString("trades"),
				MkString("kline"),
				MkString("accuracy"),
			})}),
			"private": MkMap(&VarMap{"post": MkArray(&VarArray{
				MkString("user_info"),
				MkString("create_order"),
				MkString("cancel_order"),
				MkString("orders_info"),
				MkString("orders_info_history"),
				MkString("withdraw"),
				MkString("withdrawCancel"),
				MkString("withdraws"),
				MkString("withdrawConfigs"),
			})}),
		}),
		"fees": MkMap(&VarMap{
			"trading": MkMap(&VarMap{
				"maker": this.ParseNumber(MkString("0.001")),
				"taker": this.ParseNumber(MkString("0.001")),
			}),
			"funding": MkMap(&VarMap{"withdraw": MkMap(&VarMap{})}),
		}),
		"commonCurrencies": MkMap(&VarMap{
			"VET_ERC20": MkString("VEN"),
			"PNT":       MkString("Penta"),
		}),
		"options": MkMap(&VarMap{"cacheSecretAsPem": MkBool(true)}),
	}))
}

func (this *Lbank) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetAccuracy"), params)
	_ = response
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
		market := *(response).At(i)
		_ = market
		id := *(market).At(MkString("symbol"))
		_ = id
		parts := id.Split(MkString("_"))
		_ = parts
		baseId := OpCopy(MkUndefined())
		_ = baseId
		quoteId := OpCopy(MkUndefined())
		_ = quoteId
		numParts := OpCopy(parts.Length)
		_ = numParts
		if IsTrue(OpGt(numParts, MkInteger(2))) {
			baseId = OpAdd(*(parts).At(MkInteger(0)), OpAdd(MkString("_"), *(parts).At(MkInteger(1))))
			quoteId = *(parts).At(MkInteger(2))
		} else {
			baseId = *(parts).At(MkInteger(0))
			quoteId = *(parts).At(MkInteger(1))
		}
		base := this.SafeCurrencyCode(baseId)
		_ = base
		quote := this.SafeCurrencyCode(quoteId)
		_ = quote
		symbol := OpAdd(base, OpAdd(MkString("/"), quote))
		_ = symbol
		precision := MkMap(&VarMap{
			"amount": this.SafeInteger(market, MkString("quantityAccuracy")),
			"price":  this.SafeInteger(market, MkString("priceAccuracy")),
		})
		_ = precision
		result.Push(MkMap(&VarMap{
			"id":        id,
			"symbol":    symbol,
			"base":      base,
			"quote":     quote,
			"baseId":    baseId,
			"quoteId":   quoteId,
			"active":    MkBool(true),
			"precision": precision,
			"limits": MkMap(&VarMap{
				"amount": MkMap(&VarMap{
					"min": Math.Pow(MkInteger(10), OpNeg(*(precision).At(MkString("amount")))),
					"max": MkUndefined(),
				}),
				"price": MkMap(&VarMap{
					"min": Math.Pow(MkInteger(10), OpNeg(*(precision).At(MkString("price")))),
					"max": Math.Pow(MkInteger(10), *(precision).At(MkString("price"))),
				}),
				"cost": MkMap(&VarMap{
					"min": MkUndefined(),
					"max": MkUndefined(),
				}),
			}),
			"info": id,
		}))
	}
	return result
}

func (this *Lbank) ParseTicker(goArgs ...*Variant) *Variant {
	ticker := GoGetArg(goArgs, 0, MkUndefined())
	_ = ticker
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(OpEq2(market, MkUndefined())) {
		marketId := this.SafeString(ticker, MkString("symbol"))
		_ = marketId
		if IsTrue(OpHasMember(marketId, *this.At(MkString("markets_by_id")))) {
			market := *(*this.At(MkString("markets_by_id"))).At(marketId)
			_ = market
			symbol = *(market).At(MkString("symbol"))
		} else {
			parts := marketId.Split(MkString("_"))
			_ = parts
			baseId := OpCopy(MkUndefined())
			_ = baseId
			quoteId := OpCopy(MkUndefined())
			_ = quoteId
			numParts := OpCopy(parts.Length)
			_ = numParts
			if IsTrue(OpGt(numParts, MkInteger(2))) {
				baseId = OpAdd(*(parts).At(MkInteger(0)), OpAdd(MkString("_"), *(parts).At(MkInteger(1))))
				quoteId = *(parts).At(MkInteger(2))
			} else {
				baseId = *(parts).At(MkInteger(0))
				quoteId = *(parts).At(MkInteger(1))
			}
			base := this.SafeCurrencyCode(baseId)
			_ = base
			quote := this.SafeCurrencyCode(quoteId)
			_ = quote
			symbol = OpAdd(base, OpAdd(MkString("/"), quote))
		}
	}
	timestamp := this.SafeInteger(ticker, MkString("timestamp"))
	_ = timestamp
	info := OpCopy(ticker)
	_ = info
	ticker = *(info).At(MkString("ticker"))
	last := this.SafeNumber(ticker, MkString("latest"))
	_ = last
	percentage := this.SafeNumber(ticker, MkString("change"))
	_ = percentage
	open := OpCopy(MkUndefined())
	_ = open
	if IsTrue(OpNotEq2(percentage, MkUndefined())) {
		relativeChange := this.Sum(MkInteger(1), OpDiv(percentage, MkInteger(100)))
		_ = relativeChange
		if IsTrue(OpGt(relativeChange, MkInteger(0))) {
			open = OpDiv(last, this.Sum(MkInteger(1), relativeChange))
		}
	}
	change := OpCopy(MkUndefined())
	_ = change
	average := OpCopy(MkUndefined())
	_ = average
	if IsTrue(OpAnd(OpNotEq2(last, MkUndefined()), OpNotEq2(open, MkUndefined()))) {
		change = OpSub(last, open)
		average = OpDiv(this.Sum(last, open), MkInteger(2))
	}
	if IsTrue(OpNotEq2(market, MkUndefined())) {
		symbol = *(market).At(MkString("symbol"))
	}
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
		"open":          MkUndefined(),
		"close":         last,
		"last":          last,
		"previousClose": MkUndefined(),
		"change":        change,
		"percentage":    percentage,
		"average":       average,
		"baseVolume":    this.SafeNumber(ticker, MkString("vol")),
		"quoteVolume":   this.SafeNumber(ticker, MkString("turnover")),
		"info":          info,
	})
}

func (this *Lbank) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("publicGetTicker"), this.Extend(request, params))
	_ = response
	return this.ParseTicker(response, market)
}

func (this *Lbank) FetchTickers(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"symbol": MkString("all")})
	_ = request
	response := this.Call(MkString("publicGetTicker"), this.Extend(request, params))
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

func (this *Lbank) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkInteger(60))
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	size := MkInteger(60)
	_ = size
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		size = Math.Min(limit, size)
	}
	request := MkMap(&VarMap{
		"symbol": this.MarketId(symbol),
		"size":   size,
	})
	_ = request
	response := this.Call(MkString("publicGetDepth"), this.Extend(request, params))
	_ = response
	return this.ParseOrderBook(response, symbol)
}

func (this *Lbank) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(OpNotEq2(market, MkUndefined())) {
		symbol = *(market).At(MkString("symbol"))
	}
	timestamp := this.SafeInteger(trade, MkString("date_ms"))
	_ = timestamp
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
	id := this.SafeString(trade, MkString("tid"))
	_ = id
	type_ := OpCopy(MkUndefined())
	_ = type_
	side := this.SafeString(trade, MkString("type"))
	_ = side
	side = side.Replace(MkString("_market"), MkString(""))
	return MkMap(&VarMap{
		"id":           id,
		"info":         this.SafeValue(trade, MkString("info"), trade),
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
		"fee":          MkUndefined(),
	})
}

func (this *Lbank) FetchTrades(goArgs ...*Variant) *Variant {
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
		"symbol": *(market).At(MkString("id")),
		"size":   MkInteger(100),
	})
	_ = request
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("time")) = parseInt(since)
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("size")) = OpCopy(limit)
	}
	response := this.Call(MkString("publicGetTrades"), this.Extend(request, params))
	_ = response
	return this.ParseTrades(response, market, since, limit)
}

func (this *Lbank) ParseOHLCV(goArgs ...*Variant) *Variant {
	ohlcv := GoGetArg(goArgs, 0, MkUndefined())
	_ = ohlcv
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	return MkArray(&VarArray{
		this.SafeTimestamp(ohlcv, MkInteger(0)),
		this.SafeNumber(ohlcv, MkInteger(1)),
		this.SafeNumber(ohlcv, MkInteger(2)),
		this.SafeNumber(ohlcv, MkInteger(3)),
		this.SafeNumber(ohlcv, MkInteger(4)),
		this.SafeNumber(ohlcv, MkInteger(5)),
	})
}

func (this *Lbank) FetchOHLCV(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	timeframe := GoGetArg(goArgs, 1, MkString("5m"))
	_ = timeframe
	since := GoGetArg(goArgs, 2, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 3, MkInteger(1000))
	_ = limit
	params := GoGetArg(goArgs, 4, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	if IsTrue(OpEq2(since, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchOHLCV() requires a `since` argument"))))
	}
	if IsTrue(OpEq2(limit, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchOHLCV() requires a `limit` argument"))))
	}
	request := MkMap(&VarMap{
		"symbol": *(market).At(MkString("id")),
		"type":   *(*this.At(MkString("timeframes"))).At(timeframe),
		"size":   limit,
		"time":   parseInt(OpDiv(since, MkInteger(1000))),
	})
	_ = request
	response := this.Call(MkString("publicGetKline"), this.Extend(request, params))
	_ = response
	return this.ParseOHLCVs(response, market, timeframe, since, limit)
}

func (this *Lbank) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privatePostUserInfo"), params)
	_ = response
	result := MkMap(&VarMap{
		"info":      response,
		"timestamp": MkUndefined(),
		"datetime":  MkUndefined(),
	})
	_ = result
	info := this.SafeValue(response, MkString("info"), MkMap(&VarMap{}))
	_ = info
	free := this.SafeValue(info, MkString("free"), MkMap(&VarMap{}))
	_ = free
	freeze := this.SafeValue(info, MkString("freeze"), MkMap(&VarMap{}))
	_ = freeze
	asset := this.SafeValue(info, MkString("asset"), MkMap(&VarMap{}))
	_ = asset
	currencyIds := GoGetKeys(free)
	_ = currencyIds
	for i := MkInteger(0); IsTrue(OpLw(i, currencyIds.Length)); OpIncr(&i) {
		currencyId := *(currencyIds).At(i)
		_ = currencyId
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		account := this.Account()
		_ = account
		*(account).At(MkString("free")) = this.SafeString(free, currencyId)
		*(account).At(MkString("used")) = this.SafeString(freeze, currencyId)
		*(account).At(MkString("total")) = this.SafeString(asset, currencyId)
		*(result).At(code) = OpCopy(account)
	}
	return this.ParseBalance(result)
}

func (this *Lbank) ParseOrderStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"-1": MkString("cancelled"),
		"0":  MkString("open"),
		"1":  MkString("open"),
		"2":  MkString("closed"),
		"4":  MkString("closed"),
	})
	_ = statuses
	return this.SafeString(statuses, status)
}

func (this *Lbank) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	marketId := this.SafeString(order, MkString("symbol"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market, MkString("_"))
	_ = symbol
	timestamp := this.SafeInteger(order, MkString("create_time"))
	_ = timestamp
	price := this.SafeNumber(order, MkString("price"))
	_ = price
	amount := this.SafeNumber(order, MkString("amount"), MkNumber(0.0))
	_ = amount
	filled := this.SafeNumber(order, MkString("deal_amount"), MkNumber(0.0))
	_ = filled
	average := this.SafeNumber(order, MkString("avg_price"))
	_ = average
	status := this.ParseOrderStatus(this.SafeString(order, MkString("status")))
	_ = status
	id := this.SafeString(order, MkString("order_id"))
	_ = id
	type_ := this.SafeString(order, MkString("order_type"))
	_ = type_
	side := this.SafeString(order, MkString("type"))
	_ = side
	return this.SafeOrder(MkMap(&VarMap{
		"id":                 id,
		"clientOrderId":      MkUndefined(),
		"datetime":           this.Iso8601(timestamp),
		"timestamp":          timestamp,
		"lastTradeTimestamp": MkUndefined(),
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
		"remaining":          MkUndefined(),
		"trades":             MkUndefined(),
		"fee":                MkUndefined(),
		"info":               this.SafeValue(order, MkString("info"), order),
		"average":            average,
	}))
}

func (this *Lbank) CreateOrder(goArgs ...*Variant) *Variant {
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
	order := MkMap(&VarMap{
		"symbol": *(market).At(MkString("id")),
		"type":   side,
		"amount": amount,
	})
	_ = order
	if IsTrue(OpEq2(type_, MkString("market"))) {
		OpAddAssign(&*(order).At(MkString("type")), MkString("_market"))
	} else {
		*(order).At(MkString("price")) = OpCopy(price)
	}
	response := this.Call(MkString("privatePostCreateOrder"), this.Extend(order, params))
	_ = response
	order = this.Omit(order, MkString("type"))
	*(order).At(MkString("order_id")) = *(response).At(MkString("order_id"))
	*(order).At(MkString("type")) = OpCopy(side)
	*(order).At(MkString("order_type")) = OpCopy(type_)
	*(order).At(MkString("create_time")) = this.Milliseconds()
	*(order).At(MkString("info")) = OpCopy(response)
	return this.ParseOrder(order, market)
}

func (this *Lbank) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{
		"symbol":   *(market).At(MkString("id")),
		"order_id": id,
	})
	_ = request
	response := this.Call(MkString("privatePostCancelOrder"), this.Extend(request, params))
	_ = response
	return response
}

func (this *Lbank) FetchOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{
		"symbol":   *(market).At(MkString("id")),
		"order_id": id,
	})
	_ = request
	response := this.Call(MkString("privatePostOrdersInfo"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("orders"), MkArray(&VarArray{}))
	_ = data
	orders := this.ParseOrders(data, market)
	_ = orders
	numOrders := OpCopy(orders.Length)
	_ = numOrders
	if IsTrue(OpEq2(numOrders, MkInteger(1))) {
		return *(orders).At(MkInteger(0))
	} else {
		return orders
	}
	return MkUndefined()
}

func (this *Lbank) FetchOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	if IsTrue(OpEq2(limit, MkUndefined())) {
		limit = MkInteger(100)
	}
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{
		"symbol":       *(market).At(MkString("id")),
		"current_page": MkInteger(1),
		"page_length":  limit,
	})
	_ = request
	response := this.Call(MkString("privatePostOrdersInfoHistory"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("orders"), MkArray(&VarArray{}))
	_ = data
	return this.ParseOrders(data, MkUndefined(), since, limit)
}

func (this *Lbank) FetchClosedOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	orders := this.FetchOrders(symbol, since, limit, params)
	_ = orders
	closed := this.FilterBy(orders, MkString("status"), MkString("closed"))
	_ = closed
	canceled := this.FilterBy(orders, MkString("status"), MkString("cancelled"))
	_ = canceled
	allOrders := this.ArrayConcat(closed, canceled)
	_ = allOrders
	return this.FilterBySymbolSinceLimit(allOrders, symbol, since, limit)
}

func (this *Lbank) Withdraw(goArgs ...*Variant) *Variant {
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
		"assetCode": *(currency).At(MkString("id")),
		"amount":    amount,
		"account":   address,
	})
	_ = request
	if IsTrue(OpNotEq2(tag, MkUndefined())) {
		*(request).At(MkString("memo")) = OpCopy(tag)
	}
	response := this.Call(MkString("privatePostWithdraw"), this.Extend(request, params))
	_ = response
	return MkMap(&VarMap{
		"id":   this.SafeString(response, MkString("id")),
		"info": response,
	})
}

func (this *Lbank) ConvertSecretToPem(goArgs ...*Variant) *Variant {
	secret := GoGetArg(goArgs, 0, MkUndefined())
	_ = secret
	lineLength := MkInteger(64)
	_ = lineLength
	secretLength := OpSub(secret.Length, MkInteger(0))
	_ = secretLength
	numLines := parseInt(OpDiv(secretLength, lineLength))
	_ = numLines
	numLines = this.Sum(numLines, MkInteger(1))
	pem := MkString("-----BEGIN PRIVATE KEY-----\n")
	_ = pem
	for i := MkInteger(0); IsTrue(OpLw(i, numLines)); OpIncr(&i) {
		start := OpMulti(i, lineLength)
		_ = start
		end := this.Sum(start, lineLength)
		_ = end
		OpAddAssign(&pem, OpAdd((*((this).At(MkString("secret")))).Call(MkString("slice"), start, end), MkString("\n")))
	}
	return OpAdd(pem, MkString("-----END PRIVATE KEY-----"))
}

func (this *Lbank) Sign(goArgs ...*Variant) *Variant {
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
	query := this.Omit(params, this.ExtractParams(path))
	_ = query
	url := OpAdd(*(*this.At(MkString("urls"))).At(MkString("api")), OpAdd(MkString("/"), OpAdd(*this.At(MkString("version")), OpAdd(MkString("/"), this.ImplodeParams(path, params)))))
	_ = url
	OpAddAssign(&url, MkString(".do"))
	if IsTrue(OpEq2(api, MkString("public"))) {
		if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
			OpAddAssign(&url, OpAdd(MkString("?"), this.Urlencode(query)))
		}
	} else {
		this.CheckRequiredCredentials()
		query := this.Keysort(this.Extend(MkMap(&VarMap{"api_key": *this.At(MkString("apiKey"))}), params))
		_ = query
		queryString := this.Rawencode(query)
		_ = queryString
		message := (this.Hash(this.Encode(queryString))).Call(MkString("toUpperCase"))
		_ = message
		cacheSecretAsPem := this.SafeValue(*this.At(MkString("options")), MkString("cacheSecretAsPem"), MkBool(true))
		_ = cacheSecretAsPem
		pem := OpCopy(MkUndefined())
		_ = pem
		if IsTrue(cacheSecretAsPem) {
			pem = this.SafeValue(*this.At(MkString("options")), MkString("pem"))
			if IsTrue(OpEq2(pem, MkUndefined())) {
				pem = this.ConvertSecretToPem(*this.At(MkString("secret")))
				*(*this.At(MkString("options"))).At(MkString("pem")) = OpCopy(pem)
			}
		} else {
			pem = this.ConvertSecretToPem(*this.At(MkString("secret")))
		}
		sign := this.BinaryToBase64(this.Rsa(message, this.Encode(pem), MkString("RS256")))
		_ = sign
		*(query).At(MkString("sign")) = OpCopy(sign)
		body = this.Urlencode(query)
		headers = MkMap(&VarMap{"Content-Type": MkString("application/x-www-form-urlencoded")})
	}
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Lbank) HandleErrors(goArgs ...*Variant) *Variant {
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
	success := this.SafeString(response, MkString("result"))
	_ = success
	if IsTrue(OpEq2(success, MkString("false"))) {
		errorCode := this.SafeString(response, MkString("error_code"))
		_ = errorCode
		message := this.SafeString(MkMap(&VarMap{
			"10000": MkString("Internal error"),
			"10001": MkString("The required parameters can not be empty"),
			"10002": MkString("verification failed"),
			"10003": MkString("Illegal parameters"),
			"10004": MkString("User requests are too frequent"),
			"10005": MkString("Key does not exist"),
			"10006": MkString("user does not exist"),
			"10007": MkString("Invalid signature"),
			"10008": MkString("This currency pair is not supported"),
			"10009": MkString("Limit orders can not be missing orders and the number of orders"),
			"10010": MkString("Order price or order quantity must be greater than 0"),
			"10011": MkString("Market orders can not be missing the amount of the order"),
			"10012": MkString("market sell orders can not be missing orders"),
			"10013": MkString("is less than the minimum trading position 0.001"),
			"10014": MkString("Account number is not enough"),
			"10015": MkString("The order type is wrong"),
			"10016": MkString("Account balance is not enough"),
			"10017": MkString("Abnormal server"),
			"10018": MkString("order inquiry can not be more than 50 less than one"),
			"10019": MkString("withdrawal orders can not be more than 3 less than one"),
			"10020": MkString("less than the minimum amount of the transaction limit of 0.001"),
			"10022": MkString("Insufficient key authority"),
		}), errorCode, this.Json(response))
		_ = message
		ErrorClass := this.SafeValue(MkMap(&VarMap{
			"10002": AuthenticationError,
			"10004": DDoSProtection,
			"10005": AuthenticationError,
			"10006": AuthenticationError,
			"10007": AuthenticationError,
			"10009": InvalidOrder,
			"10010": InvalidOrder,
			"10011": InvalidOrder,
			"10012": InvalidOrder,
			"10013": InvalidOrder,
			"10014": InvalidOrder,
			"10015": InvalidOrder,
			"10016": InvalidOrder,
			"10022": AuthenticationError,
		}), errorCode, ExchangeError)
		_ = ErrorClass
		panic(NewErrorClass(message))
	}
	return MkUndefined()
}
