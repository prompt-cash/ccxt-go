package ccxt

import ()

type Coinegg struct {
	*ExchangeBase
}

var _ Exchange = (*Coinegg)(nil)

func init() {
	exchange := &Coinegg{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Coinegg) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("coinegg"),
		"name": MkString("CoinEgg"),
		"countries": MkArray(&VarArray{
			MkString("CN"),
			MkString("UK"),
		}),
		"has": MkMap(&VarMap{
			"cancelOrder":     MkBool(true),
			"createOrder":     MkBool(true),
			"fetchBalance":    MkBool(true),
			"fetchMarkets":    MkBool(true),
			"fetchMyTrades":   MkBool(false),
			"fetchOpenOrders": MkString("emulated"),
			"fetchOrder":      MkBool(true),
			"fetchOrderBook":  MkBool(true),
			"fetchOrders":     MkBool(true),
			"fetchTicker":     MkBool(true),
			"fetchTickers":    MkBool(false),
			"fetchTrades":     MkBool(true),
		}),
		"urls": MkMap(&VarMap{
			"logo": MkString("https://user-images.githubusercontent.com/1294454/36770310-adfa764e-1c5a-11e8-8e09-449daac3d2fb.jpg"),
			"api": MkMap(&VarMap{
				"web":  MkString("https://trade.coinegg.com/web"),
				"rest": MkString("https://api.coinegg.com/api/v1"),
			}),
			"www":      MkString("https://www.coinegg.com"),
			"doc":      MkString("https://www.coinegg.com/explain.api.html"),
			"fees":     MkString("https://www.coinegg.com/fee.html"),
			"referral": MkString("https://www.coinegg.com/user/register?invite=523218"),
		}),
		"api": MkMap(&VarMap{
			"web": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("symbol/ticker?right_coin={quote}"),
				MkString("{quote}/trends"),
				MkString("{quote}/{base}/order"),
				MkString("{quote}/{base}/trades"),
				MkString("{quote}/{base}/depth.js"),
			})}),
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("ticker/region/{quote}"),
				MkString("depth/region/{quote}"),
				MkString("orders/region/{quote}"),
			})}),
			"private": MkMap(&VarMap{"post": MkArray(&VarArray{
				MkString("balance"),
				MkString("trade_add/region/{quote}"),
				MkString("trade_cancel/region/{quote}"),
				MkString("trade_view/region/{quote}"),
				MkString("trade_list/region/{quote}"),
			})}),
		}),
		"fees": MkMap(&VarMap{
			"trading": MkMap(&VarMap{
				"maker": OpDiv(MkNumber(0.1), MkInteger(100)),
				"taker": OpDiv(MkNumber(0.1), MkInteger(100)),
			}),
			"funding": MkMap(&VarMap{"withdraw": MkMap(&VarMap{
				"BTC":  MkNumber(0.008),
				"BCH":  MkNumber(0.002),
				"LTC":  MkNumber(0.001),
				"ETH":  MkNumber(0.01),
				"ETC":  MkNumber(0.01),
				"NEO":  MkInteger(0),
				"QTUM": MkString("1%"),
				"XRP":  MkString("1%"),
				"DOGE": MkString("1%"),
				"LSK":  MkString("1%"),
				"XAS":  MkString("1%"),
				"BTS":  MkString("1%"),
				"GAME": MkString("1%"),
				"GOOC": MkString("1%"),
				"NXT":  MkString("1%"),
				"IFC":  MkString("1%"),
				"DNC":  MkString("1%"),
				"BLK":  MkString("1%"),
				"VRC":  MkString("1%"),
				"XPM":  MkString("1%"),
				"VTC":  MkString("1%"),
				"TFC":  MkString("1%"),
				"PLC":  MkString("1%"),
				"EAC":  MkString("1%"),
				"PPC":  MkString("1%"),
				"FZ":   MkString("1%"),
				"ZET":  MkString("1%"),
				"RSS":  MkString("1%"),
				"PGC":  MkString("1%"),
				"SKT":  MkString("1%"),
				"JBC":  MkString("1%"),
				"RIO":  MkString("1%"),
				"LKC":  MkString("1%"),
				"ZCC":  MkString("1%"),
				"MCC":  MkString("1%"),
				"QEC":  MkString("1%"),
				"MET":  MkString("1%"),
				"YTC":  MkString("1%"),
				"HLB":  MkString("1%"),
				"MRYC": MkString("1%"),
				"MTC":  MkString("1%"),
				"KTC":  MkInteger(0),
			})}),
		}),
		"exceptions": MkMap(&VarMap{
			"103": AuthenticationError,
			"104": AuthenticationError,
			"105": AuthenticationError,
			"106": InvalidNonce,
			"200": InsufficientFunds,
			"201": InvalidOrder,
			"202": InvalidOrder,
			"203": OrderNotFound,
			"402": DDoSProtection,
		}),
		"errorMessages": MkMap(&VarMap{
			"100": MkString("Required parameters can not be empty"),
			"101": MkString("Illegal parameter"),
			"102": MkString("coin does not exist"),
			"103": MkString("Key does not exist"),
			"104": MkString("Signature does not match"),
			"105": MkString("Insufficient permissions"),
			"106": MkString("Request expired(nonce error)"),
			"200": MkString("Lack of balance"),
			"201": MkString("Too small for the number of trading"),
			"202": MkString("Price must be in 0 - 1000000"),
			"203": MkString("Order does not exist"),
			"204": MkString("Pending order amount must be above 0.001 BTC"),
			"205": MkString("Restrict pending order prices"),
			"206": MkString("Decimal place error"),
			"401": MkString("System error"),
			"402": MkString("Requests are too frequent"),
			"403": MkString("Non-open API"),
			"404": MkString("IP restriction does not request the resource"),
			"405": MkString("Currency transactions are temporarily closed"),
		}),
		"options": MkMap(&VarMap{"quoteIds": MkArray(&VarArray{
			MkString("btc"),
			MkString("eth"),
			MkString("usc"),
			MkString("usdt"),
		})}),
		"commonCurrencies": MkMap(&VarMap{
			"JBC":  MkString("JubaoCoin"),
			"SBTC": MkString("Super Bitcoin"),
		}),
	}))
}

func (this *Coinegg) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	quoteIds := *(*this.At(MkString("options"))).At(MkString("quoteIds"))
	_ = quoteIds
	result := MkArray(&VarArray{})
	_ = result
	for b := MkInteger(0); IsTrue(OpLw(b, quoteIds.Length)); OpIncr(&b) {
		quoteId := *(quoteIds).At(b)
		_ = quoteId
		response := this.Call(MkString("webGetSymbolTickerRightCoinQuote"), MkMap(&VarMap{"quote": quoteId}))
		_ = response
		tickers := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
		_ = tickers
		for i := MkInteger(0); IsTrue(OpLw(i, tickers.Length)); OpIncr(&i) {
			ticker := *(tickers).At(i)
			_ = ticker
			id := *(ticker).At(MkString("symbol"))
			_ = id
			baseId := *(id.Split(MkString("_"))).At(MkInteger(0))
			_ = baseId
			base := baseId.ToUpperCase()
			_ = base
			quote := quoteId.ToUpperCase()
			_ = quote
			base = this.SafeCurrencyCode(base)
			quote = this.SafeCurrencyCode(quote)
			symbol := OpAdd(base, OpAdd(MkString("/"), quote))
			_ = symbol
			precision := MkMap(&VarMap{
				"amount": MkInteger(8),
				"price":  MkInteger(8),
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
						"max": Math.Pow(MkInteger(10), *(precision).At(MkString("amount"))),
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
				"info": ticker,
			}))
		}
	}
	return result
}

func (this *Coinegg) ParseTicker(goArgs ...*Variant) *Variant {
	ticker := GoGetArg(goArgs, 0, MkUndefined())
	_ = ticker
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	symbol := *(market).At(MkString("symbol"))
	_ = symbol
	timestamp := this.Milliseconds()
	_ = timestamp
	last := this.SafeNumber(ticker, MkString("last"))
	_ = last
	percentage := this.SafeNumber(ticker, MkString("change"))
	_ = percentage
	open := OpCopy(MkUndefined())
	_ = open
	change := OpCopy(MkUndefined())
	_ = change
	average := OpCopy(MkUndefined())
	_ = average
	if IsTrue(OpNotEq2(percentage, MkUndefined())) {
		relativeChange := OpDiv(percentage, MkInteger(100))
		_ = relativeChange
		open = OpDiv(last, this.Sum(MkInteger(1), relativeChange))
		change = OpSub(last, open)
		average = OpDiv(this.Sum(last, open), MkInteger(2))
	}
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
		"open":          open,
		"close":         last,
		"last":          last,
		"previousClose": MkUndefined(),
		"change":        change,
		"percentage":    percentage,
		"average":       average,
		"baseVolume":    this.SafeNumber(ticker, MkString("vol")),
		"quoteVolume":   this.SafeNumber(ticker, MkString("quoteVol")),
		"info":          ticker,
	})
}

func (this *Coinegg) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{
		"coin":  *(market).At(MkString("baseId")),
		"quote": *(market).At(MkString("quoteId")),
	})
	_ = request
	response := this.Call(MkString("publicGetTickerRegionQuote"), this.Extend(request, params))
	_ = response
	return this.ParseTicker(response, market)
}

func (this *Coinegg) FetchOrderBook(goArgs ...*Variant) *Variant {
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
		"coin":  *(market).At(MkString("baseId")),
		"quote": *(market).At(MkString("quoteId")),
	})
	_ = request
	response := this.Call(MkString("publicGetDepthRegionQuote"), this.Extend(request, params))
	_ = response
	return this.ParseOrderBook(response, symbol)
}

func (this *Coinegg) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.SafeTimestamp(trade, MkString("date"))
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
	symbol := *(market).At(MkString("symbol"))
	_ = symbol
	type_ := MkString("limit")
	_ = type_
	side := this.SafeString(trade, MkString("type"))
	_ = side
	id := this.SafeString(trade, MkString("tid"))
	_ = id
	return MkMap(&VarMap{
		"id":           id,
		"info":         trade,
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

func (this *Coinegg) FetchTrades(goArgs ...*Variant) *Variant {
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
		"coin":  *(market).At(MkString("baseId")),
		"quote": *(market).At(MkString("quoteId")),
	})
	_ = request
	response := this.Call(MkString("publicGetOrdersRegionQuote"), this.Extend(request, params))
	_ = response
	return this.ParseTrades(response, market, since, limit)
}

func (this *Coinegg) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privatePostBalance"), params)
	_ = response
	result := MkMap(&VarMap{"info": response})
	_ = result
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	balances := this.Omit(data, MkString("uid"))
	_ = balances
	keys := GoGetKeys(balances)
	_ = keys
	for i := MkInteger(0); IsTrue(OpLw(i, keys.Length)); OpIncr(&i) {
		key := *(keys).At(i)
		_ = key
		currencyId := MkUndefined()
		_ = currencyId
		accountType := MkUndefined()
		_ = accountType
		ArrayUnpack(key.Split(MkString("_")), &currencyId, &accountType)
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		if IsTrue(OpNot(OpHasMember(code, result))) {
			*(result).At(code) = this.Account()
		}
		type_ := OpTrinary(OpEq2(accountType, MkString("lock")), MkString("used"), MkString("free"))
		_ = type_
		*(*(result).At(code)).At(type_) = this.SafeString(balances, key)
	}
	return this.ParseBalance(result)
}

func (this *Coinegg) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(OpNotEq2(market, MkUndefined())) {
		symbol = *(market).At(MkString("symbol"))
	}
	timestamp := this.Parse8601(this.SafeString(order, MkString("datetime")))
	_ = timestamp
	price := this.SafeNumber(order, MkString("price"))
	_ = price
	amount := this.SafeNumber(order, MkString("amount_original"))
	_ = amount
	remaining := this.SafeNumber(order, MkString("amount_outstanding"))
	_ = remaining
	status := this.SafeString(order, MkString("status"))
	_ = status
	if IsTrue(OpEq2(status, MkString("cancelled"))) {
		status = MkString("canceled")
	} else {
		status = OpTrinary(remaining, MkString("open"), MkString("closed"))
	}
	info := this.SafeValue(order, MkString("info"), order)
	_ = info
	type_ := MkString("limit")
	_ = type_
	side := this.SafeString(order, MkString("type"))
	_ = side
	id := this.SafeString(order, MkString("id"))
	_ = id
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
		"filled":             MkUndefined(),
		"remaining":          remaining,
		"trades":             MkUndefined(),
		"fee":                MkUndefined(),
		"info":               info,
		"average":            MkUndefined(),
	}))
}

func (this *Coinegg) CreateOrder(goArgs ...*Variant) *Variant {
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
		"coin":   *(market).At(MkString("baseId")),
		"quote":  *(market).At(MkString("quoteId")),
		"type":   side,
		"amount": amount,
		"price":  price,
	})
	_ = request
	response := this.Call(MkString("privatePostTradeAddRegionQuote"), this.Extend(request, params))
	_ = response
	id := this.SafeString(response, MkString("id"))
	_ = id
	order := this.ParseOrder(MkMap(&VarMap{
		"id":                 id,
		"datetime":           this.Ymdhms(this.Milliseconds()),
		"amount_original":    amount,
		"amount_outstanding": amount,
		"price":              price,
		"type":               side,
		"info":               response,
	}), market)
	_ = order
	return order
}

func (this *Coinegg) CancelOrder(goArgs ...*Variant) *Variant {
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
		"id":    id,
		"coin":  *(market).At(MkString("baseId")),
		"quote": *(market).At(MkString("quoteId")),
	})
	_ = request
	return this.Call(MkString("privatePostTradeCancelRegionQuote"), this.Extend(request, params))
}

func (this *Coinegg) FetchOrder(goArgs ...*Variant) *Variant {
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
		"id":    id,
		"coin":  *(market).At(MkString("baseId")),
		"quote": *(market).At(MkString("quoteId")),
	})
	_ = request
	response := this.Call(MkString("privatePostTradeViewRegionQuote"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"))
	_ = data
	return this.ParseOrder(data, market)
}

func (this *Coinegg) FetchOrders(goArgs ...*Variant) *Variant {
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
		"coin":  *(market).At(MkString("baseId")),
		"quote": *(market).At(MkString("quoteId")),
	})
	_ = request
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("since")) = OpDiv(since, MkInteger(1000))
	}
	response := this.Call(MkString("privatePostTradeListRegionQuote"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseOrders(data, market, since, limit)
}

func (this *Coinegg) FetchOpenOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	request := MkMap(&VarMap{"type": MkString("open")})
	_ = request
	return this.FetchOrders(symbol, since, limit, this.Extend(request, params))
}

func (this *Coinegg) Nonce(goArgs ...*Variant) *Variant {
	return this.Milliseconds()
}

func (this *Coinegg) Sign(goArgs ...*Variant) *Variant {
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
	apiType := MkString("rest")
	_ = apiType
	if IsTrue(OpEq2(api, MkString("web"))) {
		apiType = OpCopy(api)
	}
	url := OpAdd(*(*(*this.At(MkString("urls"))).At(MkString("api"))).At(apiType), OpAdd(MkString("/"), this.ImplodeParams(path, params)))
	_ = url
	query := this.Omit(params, this.ExtractParams(path))
	_ = query
	if IsTrue(OpOr(OpEq2(api, MkString("public")), OpEq2(api, MkString("web")))) {
		if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
			OpAddAssign(&url, OpAdd(MkString("?"), this.Urlencode(query)))
		}
	} else {
		this.CheckRequiredCredentials()
		query = this.Urlencode(this.Extend(MkMap(&VarMap{
			"key":   *this.At(MkString("apiKey")),
			"nonce": this.Nonce(),
		}), query))
		secret := this.Hash(this.Encode(*this.At(MkString("secret"))))
		_ = secret
		signature := this.Hmac(this.Encode(query), this.Encode(secret))
		_ = signature
		OpAddAssign(&query, OpAdd(MkString("&"), OpAdd(MkString("signature="), signature)))
		if IsTrue(OpEq2(method, MkString("GET"))) {
			OpAddAssign(&url, OpAdd(MkString("?"), query))
		} else {
			headers = MkMap(&VarMap{"Content-type": MkString("application/x-www-form-urlencoded")})
			body = OpCopy(query)
		}
	}
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Coinegg) HandleErrors(goArgs ...*Variant) *Variant {
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
	result := this.SafeValue(response, MkString("result"))
	_ = result
	if IsTrue(OpEq2(result, MkUndefined())) {
		return MkUndefined()
	}
	if IsTrue(OpEq2(result, MkBool(true))) {
		return MkUndefined()
	}
	errorCode := this.SafeString(response, MkString("code"))
	_ = errorCode
	errorMessages := OpCopy(*this.At(MkString("errorMessages")))
	_ = errorMessages
	message := this.SafeString(errorMessages, errorCode, MkString("Unknown Error"))
	_ = message
	feedback := OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), message))
	_ = feedback
	this.ThrowExactlyMatchedException(*this.At(MkString("exceptions")), errorCode, feedback)
	panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), message))))
	return MkUndefined()
}
