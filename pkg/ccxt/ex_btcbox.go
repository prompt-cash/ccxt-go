package ccxt

import ()

type Btcbox struct {
	*ExchangeBase
}

var _ Exchange = (*Btcbox)(nil)

func init() {
	exchange := &Btcbox{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Btcbox) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("btcbox"),
		"name": MkString("BtcBox"),
		"countries": MkArray(&VarArray{
			MkString("JP"),
		}),
		"rateLimit": MkInteger(1000),
		"version":   MkString("v1"),
		"has": MkMap(&VarMap{
			"cancelOrder":     MkBool(true),
			"CORS":            MkBool(false),
			"createOrder":     MkBool(true),
			"fetchBalance":    MkBool(true),
			"fetchOpenOrders": MkBool(true),
			"fetchOrder":      MkBool(true),
			"fetchOrderBook":  MkBool(true),
			"fetchOrders":     MkBool(true),
			"fetchTicker":     MkBool(true),
			"fetchTickers":    MkBool(false),
			"fetchTrades":     MkBool(true),
		}),
		"urls": MkMap(&VarMap{
			"logo": MkString("https://user-images.githubusercontent.com/51840849/87327317-98c55400-c53c-11ea-9a11-81f7d951cc74.jpg"),
			"api":  MkString("https://www.btcbox.co.jp/api"),
			"www":  MkString("https://www.btcbox.co.jp/"),
			"doc":  MkString("https://blog.btcbox.jp/en/archives/8762"),
			"fees": MkString("https://support.btcbox.co.jp/hc/en-us/articles/360001235694-Fees-introduction"),
		}),
		"api": MkMap(&VarMap{
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("depth"),
				MkString("orders"),
				MkString("ticker"),
			})}),
			"private": MkMap(&VarMap{"post": MkArray(&VarArray{
				MkString("balance"),
				MkString("trade_add"),
				MkString("trade_cancel"),
				MkString("trade_list"),
				MkString("trade_view"),
				MkString("wallet"),
			})}),
		}),
		"markets": MkMap(&VarMap{
			"BTC/JPY": MkMap(&VarMap{
				"id":      MkString("btc"),
				"symbol":  MkString("BTC/JPY"),
				"base":    MkString("BTC"),
				"quote":   MkString("JPY"),
				"baseId":  MkString("btc"),
				"quoteId": MkString("jpy"),
				"taker":   OpDiv(MkNumber(0.05), MkInteger(100)),
				"maker":   OpDiv(MkNumber(0.05), MkInteger(100)),
			}),
			"ETH/JPY": MkMap(&VarMap{
				"id":      MkString("eth"),
				"symbol":  MkString("ETH/JPY"),
				"base":    MkString("ETH"),
				"quote":   MkString("JPY"),
				"baseId":  MkString("eth"),
				"quoteId": MkString("jpy"),
				"taker":   OpDiv(MkNumber(0.10), MkInteger(100)),
				"maker":   OpDiv(MkNumber(0.10), MkInteger(100)),
			}),
			"LTC/JPY": MkMap(&VarMap{
				"id":      MkString("ltc"),
				"symbol":  MkString("LTC/JPY"),
				"base":    MkString("LTC"),
				"quote":   MkString("JPY"),
				"baseId":  MkString("ltc"),
				"quoteId": MkString("jpy"),
				"taker":   OpDiv(MkNumber(0.10), MkInteger(100)),
				"maker":   OpDiv(MkNumber(0.10), MkInteger(100)),
			}),
			"BCH/JPY": MkMap(&VarMap{
				"id":      MkString("bch"),
				"symbol":  MkString("BCH/JPY"),
				"base":    MkString("BCH"),
				"quote":   MkString("JPY"),
				"baseId":  MkString("bch"),
				"quoteId": MkString("jpy"),
				"taker":   OpDiv(MkNumber(0.10), MkInteger(100)),
				"maker":   OpDiv(MkNumber(0.10), MkInteger(100)),
			}),
		}),
		"exceptions": MkMap(&VarMap{
			"104": AuthenticationError,
			"105": PermissionDenied,
			"106": InvalidNonce,
			"107": InvalidOrder,
			"200": InsufficientFunds,
			"201": InvalidOrder,
			"202": InvalidOrder,
			"203": OrderNotFound,
			"401": OrderNotFound,
			"402": DDoSProtection,
		}),
	}))
}

func (this *Btcbox) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privatePostBalance"), params)
	_ = response
	result := MkMap(&VarMap{"info": response})
	_ = result
	codes := GoGetKeys(*this.At(MkString("currencies")))
	_ = codes
	for i := MkInteger(0); IsTrue(OpLw(i, codes.Length)); OpIncr(&i) {
		code := *(codes).At(i)
		_ = code
		currency := this.Currency(code)
		_ = currency
		currencyId := *(currency).At(MkString("id"))
		_ = currencyId
		free := OpAdd(currencyId, MkString("_balance"))
		_ = free
		if IsTrue(OpHasMember(free, response)) {
			account := this.Account()
			_ = account
			used := OpAdd(currencyId, MkString("_lock"))
			_ = used
			*(account).At(MkString("free")) = this.SafeString(response, free)
			*(account).At(MkString("used")) = this.SafeString(response, used)
			*(result).At(code) = OpCopy(account)
		}
	}
	return this.ParseBalance(result)
}

func (this *Btcbox) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{})
	_ = request
	numSymbols := OpCopy((*((this).At(MkString("symbols")))).Length)
	_ = numSymbols
	if IsTrue(OpGt(numSymbols, MkInteger(1))) {
		*(request).At(MkString("coin")) = *(market).At(MkString("baseId"))
	}
	response := this.Call(MkString("publicGetDepth"), this.Extend(request, params))
	_ = response
	return this.ParseOrderBook(response, symbol)
}

func (this *Btcbox) ParseTicker(goArgs ...*Variant) *Variant {
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
		"open":          MkUndefined(),
		"close":         last,
		"last":          last,
		"previousClose": MkUndefined(),
		"change":        MkUndefined(),
		"percentage":    MkUndefined(),
		"average":       MkUndefined(),
		"baseVolume":    this.SafeNumber(ticker, MkString("vol")),
		"quoteVolume":   this.SafeNumber(ticker, MkString("volume")),
		"info":          ticker,
	})
}

func (this *Btcbox) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{})
	_ = request
	numSymbols := OpCopy((*((this).At(MkString("symbols")))).Length)
	_ = numSymbols
	if IsTrue(OpGt(numSymbols, MkInteger(1))) {
		*(request).At(MkString("coin")) = *(market).At(MkString("baseId"))
	}
	response := this.Call(MkString("publicGetTicker"), this.Extend(request, params))
	_ = response
	return this.ParseTicker(response, market)
}

func (this *Btcbox) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.SafeTimestamp(trade, MkString("date"))
	_ = timestamp
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(OpNotEq2(market, MkUndefined())) {
		symbol = *(market).At(MkString("symbol"))
	}
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
	type_ := OpCopy(MkUndefined())
	_ = type_
	side := this.SafeString(trade, MkString("type"))
	_ = side
	return MkMap(&VarMap{
		"info":         trade,
		"id":           id,
		"order":        MkUndefined(),
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       symbol,
		"type":         type_,
		"side":         side,
		"takerOrMaker": MkUndefined(),
		"price":        price,
		"amount":       amount,
		"cost":         cost,
		"fee":          MkUndefined(),
	})
}

func (this *Btcbox) FetchTrades(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{})
	_ = request
	numSymbols := OpCopy((*((this).At(MkString("symbols")))).Length)
	_ = numSymbols
	if IsTrue(OpGt(numSymbols, MkInteger(1))) {
		*(request).At(MkString("coin")) = *(market).At(MkString("baseId"))
	}
	response := this.Call(MkString("publicGetOrders"), this.Extend(request, params))
	_ = response
	return this.ParseTrades(response, market, since, limit)
}

func (this *Btcbox) CreateOrder(goArgs ...*Variant) *Variant {
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
		"amount": amount,
		"price":  price,
		"type":   side,
		"coin":   *(market).At(MkString("baseId")),
	})
	_ = request
	response := this.Call(MkString("privatePostTradeAdd"), this.Extend(request, params))
	_ = response
	return this.ParseOrder(response, market)
}

func (this *Btcbox) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		symbol = MkString("BTC/JPY")
	}
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{
		"id":   id,
		"coin": *(market).At(MkString("baseId")),
	})
	_ = request
	response := this.Call(MkString("privatePostTradeCancel"), this.Extend(request, params))
	_ = response
	return this.ParseOrder(response, market)
}

func (this *Btcbox) ParseOrderStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"part":      MkString("open"),
		"all":       MkString("closed"),
		"cancelled": MkString("canceled"),
		"closed":    MkString("closed"),
		"no":        MkString("closed"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Btcbox) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	id := this.SafeString(order, MkString("id"))
	_ = id
	datetimeString := this.SafeString(order, MkString("datetime"))
	_ = datetimeString
	timestamp := OpCopy(MkUndefined())
	_ = timestamp
	if IsTrue(OpNotEq2(datetimeString, MkUndefined())) {
		timestamp = this.Parse8601(OpAdd(*(order).At(MkString("datetime")), MkString("+09:00")))
	}
	amount := this.SafeNumber(order, MkString("amount_original"))
	_ = amount
	remaining := this.SafeNumber(order, MkString("amount_outstanding"))
	_ = remaining
	price := this.SafeNumber(order, MkString("price"))
	_ = price
	status := this.ParseOrderStatus(this.SafeString(order, MkString("status")))
	_ = status
	if IsTrue(OpEq2(status, MkUndefined())) {
		if IsTrue(OpAnd(OpNotEq2(remaining, MkUndefined()), OpEq2(remaining, MkInteger(0)))) {
			status = MkString("closed")
		}
	}
	trades := OpCopy(MkUndefined())
	_ = trades
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(OpNotEq2(market, MkUndefined())) {
		symbol = *(market).At(MkString("symbol"))
	}
	side := this.SafeString(order, MkString("type"))
	_ = side
	return this.SafeOrder(MkMap(&VarMap{
		"id":                 id,
		"clientOrderId":      MkUndefined(),
		"timestamp":          timestamp,
		"datetime":           this.Iso8601(timestamp),
		"lastTradeTimestamp": MkUndefined(),
		"amount":             amount,
		"remaining":          remaining,
		"filled":             MkUndefined(),
		"side":               side,
		"type":               MkUndefined(),
		"timeInForce":        MkUndefined(),
		"postOnly":           MkUndefined(),
		"status":             status,
		"symbol":             symbol,
		"price":              price,
		"stopPrice":          MkUndefined(),
		"cost":               MkUndefined(),
		"trades":             trades,
		"fee":                MkUndefined(),
		"info":               order,
		"average":            MkUndefined(),
	}))
}

func (this *Btcbox) FetchOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		symbol = MkString("BTC/JPY")
	}
	market := this.Market(symbol)
	_ = market
	request := this.Extend(MkMap(&VarMap{
		"id":   id,
		"coin": *(market).At(MkString("baseId")),
	}), params)
	_ = request
	response := this.Call(MkString("privatePostTradeView"), this.Extend(request, params))
	_ = response
	return this.ParseOrder(response, market)
}

func (this *Btcbox) FetchOrdersByType(goArgs ...*Variant) *Variant {
	type_ := GoGetArg(goArgs, 0, MkUndefined())
	_ = type_
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 2, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 3, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 4, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		symbol = MkString("BTC/JPY")
	}
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{
		"type": type_,
		"coin": *(market).At(MkString("baseId")),
	})
	_ = request
	response := this.Call(MkString("privatePostTradeList"), this.Extend(request, params))
	_ = response
	orders := this.ParseOrders(response, market, since, limit)
	_ = orders
	if IsTrue(OpEq2(type_, MkString("open"))) {
		for i := MkInteger(0); IsTrue(OpLw(i, orders.Length)); OpIncr(&i) {
			*(*(orders).At(i)).At(MkString("status")) = MkString("open")
		}
	}
	return orders
}

func (this *Btcbox) FetchOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	return this.FetchOrdersByType(MkString("all"), symbol, since, limit, params)
}

func (this *Btcbox) FetchOpenOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	return this.FetchOrdersByType(MkString("open"), symbol, since, limit, params)
}

func (this *Btcbox) Nonce(goArgs ...*Variant) *Variant {
	return this.Milliseconds()
}

func (this *Btcbox) Sign(goArgs ...*Variant) *Variant {
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
	url := OpAdd(*(*this.At(MkString("urls"))).At(MkString("api")), OpAdd(MkString("/"), OpAdd(*this.At(MkString("version")), OpAdd(MkString("/"), path))))
	_ = url
	if IsTrue(OpEq2(api, MkString("public"))) {
		if IsTrue(*(GoGetKeys(params)).At(MkString("length"))) {
			OpAddAssign(&url, OpAdd(MkString("?"), this.Urlencode(params)))
		}
	} else {
		this.CheckRequiredCredentials()
		nonce := (this.Nonce()).Call(MkString("toString"))
		_ = nonce
		query := this.Extend(MkMap(&VarMap{
			"key":   *this.At(MkString("apiKey")),
			"nonce": nonce,
		}), params)
		_ = query
		request := this.Urlencode(query)
		_ = request
		secret := this.Hash(this.Encode(*this.At(MkString("secret"))))
		_ = secret
		*(query).At(MkString("signature")) = this.Hmac(this.Encode(request), this.Encode(secret))
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

func (this *Btcbox) HandleErrors(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpGtEq(httpCode, MkInteger(400))) {
		return MkUndefined()
	}
	result := this.SafeValue(response, MkString("result"))
	_ = result
	if IsTrue(OpOr(OpEq2(result, MkUndefined()), OpEq2(result, MkBool(true)))) {
		return MkUndefined()
	}
	code := this.SafeValue(response, MkString("code"))
	_ = code
	feedback := OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))
	_ = feedback
	this.ThrowExactlyMatchedException(*this.At(MkString("exceptions")), code, feedback)
	panic(NewExchangeError(feedback))
	return MkUndefined()
}

func (this *Btcbox) Request(goArgs ...*Variant) *Variant {
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
	config := GoGetArg(goArgs, 6, MkMap(&VarMap{}))
	_ = config
	context := GoGetArg(goArgs, 7, MkMap(&VarMap{}))
	_ = context
	response := this.Fetch2(path, api, method, params, headers, body, config, context)
	_ = response
	if IsTrue(OpEq2(OpType(response), MkString("string"))) {
		response = this.Strip(response)
		if IsTrue(OpNot(this.IsJsonEncodedObject(response))) {
			panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), response))))
		}
		response = JSON.Parse(response)
	}
	return response
}
