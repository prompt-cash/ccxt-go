package ccxt

import ()

type Bitforex struct {
	*ExchangeBase
}

var _ Exchange = (*Bitforex)(nil)

func init() {
	exchange := &Bitforex{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Bitforex) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("bitforex"),
		"name": MkString("Bitforex"),
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
			"fetchMyTrades":     MkBool(false),
			"fetchOHLCV":        MkBool(true),
			"fetchOpenOrders":   MkBool(true),
			"fetchOrder":        MkBool(true),
			"fetchOrderBook":    MkBool(true),
			"fetchOrders":       MkBool(false),
			"fetchTicker":       MkBool(true),
			"fetchTickers":      MkBool(false),
			"fetchTrades":       MkBool(true),
		}),
		"timeframes": MkMap(&VarMap{
			"1m":  MkString("1min"),
			"5m":  MkString("5min"),
			"15m": MkString("15min"),
			"30m": MkString("30min"),
			"1h":  MkString("1hour"),
			"2h":  MkString("2hour"),
			"4h":  MkString("4hour"),
			"12h": MkString("12hour"),
			"1d":  MkString("1day"),
			"1w":  MkString("1week"),
			"1M":  MkString("1month"),
		}),
		"urls": MkMap(&VarMap{
			"logo":     MkString("https://user-images.githubusercontent.com/51840849/87295553-1160ec00-c50e-11ea-8ea0-df79276a9646.jpg"),
			"api":      MkString("https://api.bitforex.com"),
			"www":      MkString("https://www.bitforex.com"),
			"doc":      MkString("https://github.com/githubdev2020/API_Doc_en/wiki"),
			"fees":     MkString("https://help.bitforex.com/en_us/?cat=13"),
			"referral": MkString("https://www.bitforex.com/en/invitationRegister?inviterId=1867438"),
		}),
		"api": MkMap(&VarMap{
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("api/v1/market/symbols"),
				MkString("api/v1/market/ticker"),
				MkString("api/v1/market/depth"),
				MkString("api/v1/market/trades"),
				MkString("api/v1/market/kline"),
			})}),
			"private": MkMap(&VarMap{"post": MkArray(&VarArray{
				MkString("api/v1/fund/mainAccount"),
				MkString("api/v1/fund/allAccount"),
				MkString("api/v1/trade/placeOrder"),
				MkString("api/v1/trade/placeMultiOrder"),
				MkString("api/v1/trade/cancelOrder"),
				MkString("api/v1/trade/cancelMultiOrder"),
				MkString("api/v1/trade/cancelAllOrder"),
				MkString("api/v1/trade/orderInfo"),
				MkString("api/v1/trade/multiOrderInfo"),
				MkString("api/v1/trade/orderInfos"),
			})}),
		}),
		"fees": MkMap(&VarMap{
			"trading": MkMap(&VarMap{
				"tierBased":  MkBool(false),
				"percentage": MkBool(true),
				"maker":      this.ParseNumber(MkString("0.001")),
				"taker":      this.ParseNumber(MkString("0.001")),
			}),
			"funding": MkMap(&VarMap{
				"tierBased":  MkBool(false),
				"percentage": MkBool(true),
				"deposit":    MkMap(&VarMap{}),
				"withdraw":   MkMap(&VarMap{}),
			}),
		}),
		"commonCurrencies": MkMap(&VarMap{
			"ACE":    MkString("ACE Entertainment"),
			"BDP":    MkString("BidiPass"),
			"CAPP":   MkString("Crypto Application Token"),
			"CREDIT": MkString("TerraCredit"),
			"CTC":    MkString("Culture Ticket Chain"),
			"GOT":    MkString("GoNetwork"),
			"HBC":    MkString("Hybrid Bank Cash"),
			"IQ":     MkString("IQ.Cash"),
			"MIR":    MkString("MIR COIN"),
			"UOS":    MkString("UOS Network"),
		}),
		"exceptions": MkMap(&VarMap{
			"1003":  BadSymbol,
			"1013":  AuthenticationError,
			"1016":  AuthenticationError,
			"1017":  PermissionDenied,
			"1019":  BadSymbol,
			"3002":  InsufficientFunds,
			"4002":  InvalidOrder,
			"4003":  InvalidOrder,
			"4004":  OrderNotFound,
			"10204": DDoSProtection,
		}),
	}))
}

func (this *Bitforex) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetApiV1MarketSymbols"), params)
	_ = response
	data := *(response).At(MkString("data"))
	_ = data
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, data.Length)); OpIncr(&i) {
		market := *(data).At(i)
		_ = market
		id := this.SafeString(market, MkString("symbol"))
		_ = id
		symbolParts := id.Split(MkString("-"))
		_ = symbolParts
		baseId := *(symbolParts).At(MkInteger(2))
		_ = baseId
		quoteId := *(symbolParts).At(MkInteger(1))
		_ = quoteId
		base := this.SafeCurrencyCode(baseId)
		_ = base
		quote := this.SafeCurrencyCode(quoteId)
		_ = quote
		symbol := OpAdd(base, OpAdd(MkString("/"), quote))
		_ = symbol
		active := OpCopy(MkBool(true))
		_ = active
		precision := MkMap(&VarMap{
			"amount": this.SafeInteger(market, MkString("amountPrecision")),
			"price":  this.SafeInteger(market, MkString("pricePrecision")),
		})
		_ = precision
		limits := MkMap(&VarMap{
			"amount": MkMap(&VarMap{
				"min": this.SafeNumber(market, MkString("minOrderAmount")),
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
		})
		_ = limits
		result.Push(MkMap(&VarMap{
			"id":        id,
			"symbol":    symbol,
			"base":      base,
			"quote":     quote,
			"baseId":    baseId,
			"quoteId":   quoteId,
			"active":    active,
			"precision": precision,
			"limits":    limits,
			"info":      market,
		}))
	}
	return result
}

func (this *Bitforex) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(OpNotEq2(market, MkUndefined())) {
		symbol = *(market).At(MkString("symbol"))
	}
	timestamp := this.SafeInteger(trade, MkString("time"))
	_ = timestamp
	id := this.SafeString(trade, MkString("tid"))
	_ = id
	orderId := OpCopy(MkUndefined())
	_ = orderId
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
	sideId := this.SafeInteger(trade, MkString("direction"))
	_ = sideId
	side := this.ParseSide(sideId)
	_ = side
	return MkMap(&VarMap{
		"info":         trade,
		"id":           id,
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       symbol,
		"type":         MkUndefined(),
		"side":         side,
		"price":        price,
		"amount":       amount,
		"cost":         cost,
		"order":        orderId,
		"fee":          MkUndefined(),
		"takerOrMaker": MkUndefined(),
	})
}

func (this *Bitforex) FetchTrades(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"symbol": this.MarketId(symbol)})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("size")) = OpCopy(limit)
	}
	market := this.Market(symbol)
	_ = market
	response := this.Call(MkString("publicGetApiV1MarketTrades"), this.Extend(request, params))
	_ = response
	return this.ParseTrades(*(response).At(MkString("data")), market, since, limit)
}

func (this *Bitforex) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privatePostApiV1FundAllAccount"), params)
	_ = response
	data := *(response).At(MkString("data"))
	_ = data
	result := MkMap(&VarMap{"info": response})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, data.Length)); OpIncr(&i) {
		balance := *(data).At(i)
		_ = balance
		currencyId := this.SafeString(balance, MkString("currency"))
		_ = currencyId
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		account := this.Account()
		_ = account
		*(account).At(MkString("used")) = this.SafeString(balance, MkString("frozen"))
		*(account).At(MkString("free")) = this.SafeString(balance, MkString("active"))
		*(account).At(MkString("total")) = this.SafeString(balance, MkString("fix"))
		*(result).At(code) = OpCopy(account)
	}
	return this.ParseBalance(result)
}

func (this *Bitforex) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := *(*this.At(MkString("markets"))).At(symbol)
	_ = market
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("publicGetApiV1MarketTicker"), this.Extend(request, params))
	_ = response
	data := *(response).At(MkString("data"))
	_ = data
	timestamp := this.SafeInteger(data, MkString("date"))
	_ = timestamp
	return MkMap(&VarMap{
		"symbol":        symbol,
		"timestamp":     timestamp,
		"datetime":      this.Iso8601(timestamp),
		"high":          this.SafeNumber(data, MkString("high")),
		"low":           this.SafeNumber(data, MkString("low")),
		"bid":           this.SafeNumber(data, MkString("buy")),
		"bidVolume":     MkUndefined(),
		"ask":           this.SafeNumber(data, MkString("sell")),
		"askVolume":     MkUndefined(),
		"vwap":          MkUndefined(),
		"open":          MkUndefined(),
		"close":         this.SafeNumber(data, MkString("last")),
		"last":          this.SafeNumber(data, MkString("last")),
		"previousClose": MkUndefined(),
		"change":        MkUndefined(),
		"percentage":    MkUndefined(),
		"average":       MkUndefined(),
		"baseVolume":    this.SafeNumber(data, MkString("vol")),
		"quoteVolume":   MkUndefined(),
		"info":          response,
	})
}

func (this *Bitforex) ParseOHLCV(goArgs ...*Variant) *Variant {
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
		this.SafeNumber(ohlcv, MkString("vol")),
	})
}

func (this *Bitforex) FetchOHLCV(goArgs ...*Variant) *Variant {
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
		"symbol": *(market).At(MkString("id")),
		"ktype":  *(*this.At(MkString("timeframes"))).At(timeframe),
	})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("size")) = OpCopy(limit)
	}
	response := this.Call(MkString("publicGetApiV1MarketKline"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseOHLCVs(data, market, timeframe, since, limit)
}

func (this *Bitforex) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	marketId := this.MarketId(symbol)
	_ = marketId
	request := MkMap(&VarMap{"symbol": marketId})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("size")) = OpCopy(limit)
	}
	response := this.Call(MkString("publicGetApiV1MarketDepth"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"))
	_ = data
	timestamp := this.SafeInteger(response, MkString("time"))
	_ = timestamp
	return this.ParseOrderBook(data, symbol, timestamp, MkString("bids"), MkString("asks"), MkString("price"), MkString("amount"))
}

func (this *Bitforex) ParseOrderStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"0": MkString("open"),
		"1": MkString("open"),
		"2": MkString("closed"),
		"3": MkString("canceled"),
		"4": MkString("canceled"),
	})
	_ = statuses
	return OpTrinary(OpHasMember(status, statuses), *(statuses).At(status), status)
}

func (this *Bitforex) ParseSide(goArgs ...*Variant) *Variant {
	sideId := GoGetArg(goArgs, 0, MkUndefined())
	_ = sideId
	if IsTrue(OpEq2(sideId, MkInteger(1))) {
		return MkString("buy")
	} else {
		if IsTrue(OpEq2(sideId, MkInteger(2))) {
			return MkString("sell")
		} else {
			return MkUndefined()
		}
	}
	return MkUndefined()
}

func (this *Bitforex) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	id := this.SafeString(order, MkString("orderId"))
	_ = id
	timestamp := this.SafeNumber(order, MkString("createTime"))
	_ = timestamp
	lastTradeTimestamp := this.SafeNumber(order, MkString("lastTime"))
	_ = lastTradeTimestamp
	symbol := *(market).At(MkString("symbol"))
	_ = symbol
	sideId := this.SafeInteger(order, MkString("tradeType"))
	_ = sideId
	side := this.ParseSide(sideId)
	_ = side
	type_ := OpCopy(MkUndefined())
	_ = type_
	price := this.SafeNumber(order, MkString("orderPrice"))
	_ = price
	average := this.SafeNumber(order, MkString("avgPrice"))
	_ = average
	amount := this.SafeNumber(order, MkString("orderAmount"))
	_ = amount
	filled := this.SafeNumber(order, MkString("dealAmount"))
	_ = filled
	status := this.ParseOrderStatus(this.SafeString(order, MkString("orderState")))
	_ = status
	feeSide := OpTrinary(OpEq2(side, MkString("buy")), MkString("base"), MkString("quote"))
	_ = feeSide
	feeCurrency := *(market).At(feeSide)
	_ = feeCurrency
	fee := MkMap(&VarMap{
		"cost":     this.SafeNumber(order, MkString("tradeFee")),
		"currency": feeCurrency,
	})
	_ = fee
	return this.SafeOrder(MkMap(&VarMap{
		"info":               order,
		"id":                 id,
		"clientOrderId":      MkUndefined(),
		"timestamp":          timestamp,
		"datetime":           this.Iso8601(timestamp),
		"lastTradeTimestamp": lastTradeTimestamp,
		"symbol":             symbol,
		"type":               type_,
		"timeInForce":        MkUndefined(),
		"postOnly":           MkUndefined(),
		"side":               side,
		"price":              price,
		"stopPrice":          MkUndefined(),
		"cost":               MkUndefined(),
		"average":            average,
		"amount":             amount,
		"filled":             filled,
		"remaining":          MkUndefined(),
		"status":             status,
		"fee":                fee,
		"trades":             MkUndefined(),
	}))
}

func (this *Bitforex) FetchOrder(goArgs ...*Variant) *Variant {
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
		"symbol":  this.MarketId(symbol),
		"orderId": id,
	})
	_ = request
	response := this.Call(MkString("privatePostApiV1TradeOrderInfo"), this.Extend(request, params))
	_ = response
	order := this.ParseOrder(*(response).At(MkString("data")), market)
	_ = order
	return order
}

func (this *Bitforex) FetchOpenOrders(goArgs ...*Variant) *Variant {
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
		"symbol": this.MarketId(symbol),
		"state":  MkInteger(0),
	})
	_ = request
	response := this.Call(MkString("privatePostApiV1TradeOrderInfos"), this.Extend(request, params))
	_ = response
	return this.ParseOrders(*(response).At(MkString("data")), market, since, limit)
}

func (this *Bitforex) FetchClosedOrders(goArgs ...*Variant) *Variant {
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
		"symbol": this.MarketId(symbol),
		"state":  MkInteger(1),
	})
	_ = request
	response := this.Call(MkString("privatePostApiV1TradeOrderInfos"), this.Extend(request, params))
	_ = response
	return this.ParseOrders(*(response).At(MkString("data")), market, since, limit)
}

func (this *Bitforex) CreateOrder(goArgs ...*Variant) *Variant {
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
	sideId := OpCopy(MkUndefined())
	_ = sideId
	if IsTrue(OpEq2(side, MkString("buy"))) {
		sideId = MkInteger(1)
	} else {
		if IsTrue(OpEq2(side, MkString("sell"))) {
			sideId = MkInteger(2)
		}
	}
	request := MkMap(&VarMap{
		"symbol":    this.MarketId(symbol),
		"price":     price,
		"amount":    amount,
		"tradeType": sideId,
	})
	_ = request
	response := this.Call(MkString("privatePostApiV1TradePlaceOrder"), this.Extend(request, params))
	_ = response
	data := *(response).At(MkString("data"))
	_ = data
	return MkMap(&VarMap{
		"info": response,
		"id":   this.SafeString(data, MkString("orderId")),
	})
}

func (this *Bitforex) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"orderId": id})
	_ = request
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		*(request).At(MkString("symbol")) = this.MarketId(symbol)
	}
	results := this.Call(MkString("privatePostApiV1TradeCancelOrder"), this.Extend(request, params))
	_ = results
	success := *(results).At(MkString("success"))
	_ = success
	returnVal := MkMap(&VarMap{
		"info":    results,
		"success": success,
	})
	_ = returnVal
	return returnVal
}

func (this *Bitforex) Sign(goArgs ...*Variant) *Variant {
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
		payload := this.Urlencode(MkMap(&VarMap{"accessKey": *this.At(MkString("apiKey"))}))
		_ = payload
		*(query).At(MkString("nonce")) = this.Milliseconds()
		if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
			OpAddAssign(&payload, OpAdd(MkString("&"), this.Urlencode(this.Keysort(query))))
		}
		message := OpAdd(MkString("/"), OpAdd(path, OpAdd(MkString("?"), payload)))
		_ = message
		signature := this.Hmac(this.Encode(message), this.Encode(*this.At(MkString("secret"))))
		_ = signature
		body = OpAdd(payload, OpAdd(MkString("&signData="), signature))
		headers = MkMap(&VarMap{"Content-Type": MkString("application/x-www-form-urlencoded")})
	}
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Bitforex) HandleErrors(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpNotEq2(OpType(body), MkString("string"))) {
		return MkUndefined()
	}
	if IsTrue(OpOr(OpEq2(*(body).At(MkInteger(0)), MkString("{")), OpEq2(*(body).At(MkInteger(0)), MkString("[")))) {
		feedback := OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))
		_ = feedback
		success := this.SafeValue(response, MkString("success"))
		_ = success
		if IsTrue(OpNotEq2(success, MkUndefined())) {
			if IsTrue(OpNot(success)) {
				code := this.SafeString(response, MkString("code"))
				_ = code
				this.ThrowExactlyMatchedException(*this.At(MkString("exceptions")), code, feedback)
				panic(NewExchangeError(feedback))
			}
		}
	}
	return MkUndefined()
}
