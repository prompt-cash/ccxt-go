package ccxt

import ()

type Mixcoins struct {
	*ExchangeBase
}

var _ Exchange = (*Mixcoins)(nil)

func init() {
	exchange := &Mixcoins{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Mixcoins) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("mixcoins"),
		"name": MkString("MixCoins"),
		"countries": MkArray(&VarArray{
			MkString("GB"),
			MkString("HK"),
		}),
		"rateLimit": MkInteger(1500),
		"version":   MkString("v1"),
		"userAgent": *(*this.At(MkString("userAgents"))).At(MkString("chrome")),
		"has": MkMap(&VarMap{
			"cancelOrder":    MkBool(true),
			"CORS":           MkBool(false),
			"createOrder":    MkBool(true),
			"fetchBalance":   MkBool(true),
			"fetchOrderBook": MkBool(true),
			"fetchTicker":    MkBool(true),
			"fetchTrades":    MkBool(true),
		}),
		"urls": MkMap(&VarMap{
			"logo": MkString("https://user-images.githubusercontent.com/51840849/87460810-1dd06c00-c616-11ea-9276-956f400d6ffa.jpg"),
			"api":  MkString("https://mixcoins.com/api"),
			"www":  MkString("https://mixcoins.com"),
			"doc":  MkString("https://mixcoins.com/help/api/"),
		}),
		"api": MkMap(&VarMap{
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("ticker/"),
				MkString("trades/"),
				MkString("depth/"),
			})}),
			"private": MkMap(&VarMap{"post": MkArray(&VarArray{
				MkString("cancel"),
				MkString("info"),
				MkString("orders"),
				MkString("order"),
				MkString("transactions"),
				MkString("trade"),
			})}),
		}),
		"markets": MkMap(&VarMap{
			"BTC/USDT": MkMap(&VarMap{
				"id":      MkString("btc_usdt"),
				"symbol":  MkString("BTC/USDT"),
				"base":    MkString("BTC"),
				"quote":   MkString("USDT"),
				"baseId":  MkString("btc"),
				"quoteId": MkString("usdt"),
				"maker":   MkNumber(0.0015),
				"taker":   MkNumber(0.0025),
			}),
			"ETH/BTC": MkMap(&VarMap{
				"id":      MkString("eth_btc"),
				"symbol":  MkString("ETH/BTC"),
				"base":    MkString("ETH"),
				"quote":   MkString("BTC"),
				"baseId":  MkString("eth"),
				"quoteId": MkString("btc"),
				"maker":   MkNumber(0.001),
				"taker":   MkNumber(0.0015),
			}),
			"BCH/BTC": MkMap(&VarMap{
				"id":      MkString("bch_btc"),
				"symbol":  MkString("BCH/BTC"),
				"base":    MkString("BCH"),
				"quote":   MkString("BTC"),
				"baseId":  MkString("bch"),
				"quoteId": MkString("btc"),
				"maker":   MkNumber(0.001),
				"taker":   MkNumber(0.0015),
			}),
			"LSK/BTC": MkMap(&VarMap{
				"id":      MkString("lsk_btc"),
				"symbol":  MkString("LSK/BTC"),
				"base":    MkString("LSK"),
				"quote":   MkString("BTC"),
				"baseId":  MkString("lsk"),
				"quoteId": MkString("btc"),
				"maker":   MkNumber(0.0015),
				"taker":   MkNumber(0.0025),
			}),
			"BCH/USDT": MkMap(&VarMap{
				"id":      MkString("bch_usdt"),
				"symbol":  MkString("BCH/USDT"),
				"base":    MkString("BCH"),
				"quote":   MkString("USDT"),
				"baseId":  MkString("bch"),
				"quoteId": MkString("usdt"),
				"maker":   MkNumber(0.001),
				"taker":   MkNumber(0.0015),
			}),
			"ETH/USDT": MkMap(&VarMap{
				"id":      MkString("eth_usdt"),
				"symbol":  MkString("ETH/USDT"),
				"base":    MkString("ETH"),
				"quote":   MkString("USDT"),
				"baseId":  MkString("eth"),
				"quoteId": MkString("usdt"),
				"maker":   MkNumber(0.001),
				"taker":   MkNumber(0.0015),
			}),
		}),
	}))
}

func (this *Mixcoins) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privatePostInfo"), params)
	_ = response
	balances := this.SafeValue(*(response).At(MkString("result")), MkString("wallet"))
	_ = balances
	result := MkMap(&VarMap{"info": response})
	_ = result
	currencyIds := GoGetKeys(balances)
	_ = currencyIds
	for i := MkInteger(0); IsTrue(OpLw(i, currencyIds.Length)); OpIncr(&i) {
		currencyId := *(currencyIds).At(i)
		_ = currencyId
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		balance := this.SafeValue(balances, currencyId, MkMap(&VarMap{}))
		_ = balance
		account := this.Account()
		_ = account
		*(account).At(MkString("free")) = this.SafeString(balance, MkString("avail"))
		*(account).At(MkString("used")) = this.SafeString(balance, MkString("lock"))
		*(result).At(code) = OpCopy(account)
	}
	return this.ParseBalance(result)
}

func (this *Mixcoins) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"market": this.MarketId(symbol)})
	_ = request
	response := this.Call(MkString("publicGetDepth"), this.Extend(request, params))
	_ = response
	return this.ParseOrderBook(*(response).At(MkString("result")), symbol)
}

func (this *Mixcoins) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"market": this.MarketId(symbol)})
	_ = request
	response := this.Call(MkString("publicGetTicker"), this.Extend(request, params))
	_ = response
	ticker := this.SafeValue(response, MkString("result"))
	_ = ticker
	timestamp := this.Milliseconds()
	_ = timestamp
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
		"quoteVolume":   MkUndefined(),
		"info":          ticker,
	})
}

func (this *Mixcoins) ParseTrade(goArgs ...*Variant) *Variant {
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
	id := this.SafeString(trade, MkString("id"))
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
	return MkMap(&VarMap{
		"id":           id,
		"info":         trade,
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       symbol,
		"type":         MkUndefined(),
		"side":         MkUndefined(),
		"order":        MkUndefined(),
		"takerOrMaker": MkUndefined(),
		"price":        price,
		"amount":       amount,
		"cost":         cost,
		"fee":          MkUndefined(),
	})
}

func (this *Mixcoins) FetchTrades(goArgs ...*Variant) *Variant {
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
	return this.ParseTrades(*(response).At(MkString("result")), market, since, limit)
}

func (this *Mixcoins) CreateOrder(goArgs ...*Variant) *Variant {
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
		"market": this.MarketId(symbol),
		"op":     side,
		"amount": amount,
	})
	_ = request
	if IsTrue(OpEq2(type_, MkString("market"))) {
		*(request).At(MkString("order_type")) = MkInteger(1)
		*(request).At(MkString("price")) = OpCopy(price)
	} else {
		*(request).At(MkString("order_type")) = MkInteger(0)
	}
	response := this.Call(MkString("privatePostTrade"), this.Extend(request, params))
	_ = response
	return MkMap(&VarMap{
		"info": response,
		"id":   (*(*(response).At(MkString("result"))).At(MkString("id"))).Call(MkString("toString")),
	})
}

func (this *Mixcoins) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"id": id})
	_ = request
	return this.Call(MkString("privatePostCancel"), this.Extend(request, params))
}

func (this *Mixcoins) Sign(goArgs ...*Variant) *Variant {
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
		nonce := this.Nonce()
		_ = nonce
		body = this.Urlencode(this.Extend(MkMap(&VarMap{"nonce": nonce}), params))
		headers = MkMap(&VarMap{
			"Content-Type": MkString("application/x-www-form-urlencoded"),
			"Key":          *this.At(MkString("apiKey")),
			"Sign":         this.Hmac(this.Encode(body), *this.At(MkString("secret")), MkString("sha512")),
		})
	}
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Mixcoins) HandleErrors(goArgs ...*Variant) *Variant {
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
	status := this.SafeInteger(response, MkString("status"))
	_ = status
	if IsTrue(OpNotEq2(status, MkInteger(200))) {
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), this.Json(response)))))
	}
	return MkUndefined()
}
