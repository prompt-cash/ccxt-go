package ccxt

import ()

type Exx struct {
	*ExchangeBase
}

var _ Exchange = (*Exx)(nil)

func init() {
	exchange := &Exx{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Exx) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("exx"),
		"name": MkString("EXX"),
		"countries": MkArray(&VarArray{
			MkString("CN"),
		}),
		"rateLimit": OpDiv(MkInteger(1000), MkInteger(10)),
		"userAgent": *(*this.At(MkString("userAgents"))).At(MkString("chrome")),
		"has": MkMap(&VarMap{
			"cancelOrder":     MkBool(true),
			"createOrder":     MkBool(true),
			"fetchBalance":    MkBool(true),
			"fetchMarkets":    MkBool(true),
			"fetchOpenOrders": MkBool(true),
			"fetchOrder":      MkBool(true),
			"fetchOrderBook":  MkBool(true),
			"fetchTicker":     MkBool(true),
			"fetchTickers":    MkBool(true),
			"fetchTrades":     MkBool(true),
		}),
		"urls": MkMap(&VarMap{
			"logo": MkString("https://user-images.githubusercontent.com/1294454/37770292-fbf613d0-2de4-11e8-9f79-f2dc451b8ccb.jpg"),
			"api": MkMap(&VarMap{
				"public":  MkString("https://api.exx.com/data/v1"),
				"private": MkString("https://trade.exx.com/api"),
			}),
			"www":      MkString("https://www.exx.com/"),
			"doc":      MkString("https://www.exx.com/help/restApi"),
			"fees":     MkString("https://www.exx.com/help/rate"),
			"referral": MkString("https://www.exx.com/r/fde4260159e53ab8a58cc9186d35501f?recommQd=1"),
		}),
		"api": MkMap(&VarMap{
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("markets"),
				MkString("tickers"),
				MkString("ticker"),
				MkString("depth"),
				MkString("trades"),
			})}),
			"private": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("order"),
				MkString("cancel"),
				MkString("getOrder"),
				MkString("getOpenOrders"),
				MkString("getBalance"),
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
			"DOS": MkString("DEMOS"),
			"TV":  MkString("TIV"),
		}),
		"exceptions": MkMap(&VarMap{"103": AuthenticationError}),
	}))
}

func (this *Exx) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetMarkets"), params)
	_ = response
	ids := GoGetKeys(response)
	_ = ids
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, ids.Length)); OpIncr(&i) {
		id := *(ids).At(i)
		_ = id
		market := *(response).At(id)
		_ = market
		baseId := MkUndefined()
		_ = baseId
		quoteId := MkUndefined()
		_ = quoteId
		ArrayUnpack(id.Split(MkString("_")), &baseId, &quoteId)
		base := this.SafeCurrencyCode(baseId)
		_ = base
		quote := this.SafeCurrencyCode(quoteId)
		_ = quote
		symbol := OpAdd(base, OpAdd(MkString("/"), quote))
		_ = symbol
		active := OpEq2(*(market).At(MkString("isOpen")), MkBool(true))
		_ = active
		amountPrecisionString := this.SafeString(market, MkString("amountScale"))
		_ = amountPrecisionString
		pricePrecisionString := this.SafeString(market, MkString("priceScale"))
		_ = pricePrecisionString
		amountLimit := this.ParsePrecision(amountPrecisionString)
		_ = amountLimit
		priceLimit := this.ParsePrecision(pricePrecisionString)
		_ = priceLimit
		precision := MkMap(&VarMap{
			"amount": parseInt(amountPrecisionString),
			"price":  parseInt(pricePrecisionString),
		})
		_ = precision
		result.Push(MkMap(&VarMap{
			"id":        id,
			"symbol":    symbol,
			"base":      base,
			"quote":     quote,
			"baseId":    baseId,
			"quoteId":   quoteId,
			"active":    active,
			"precision": precision,
			"limits": MkMap(&VarMap{
				"amount": MkMap(&VarMap{
					"min": this.ParseNumber(amountLimit),
					"max": MkUndefined(),
				}),
				"price": MkMap(&VarMap{
					"min": this.ParseNumber(priceLimit),
					"max": MkUndefined(),
				}),
				"cost": MkMap(&VarMap{
					"min": this.SafeNumber(market, MkString("minAmount")),
					"max": MkUndefined(),
				}),
			}),
			"info": market,
		}))
	}
	return result
}

func (this *Exx) ParseTicker(goArgs ...*Variant) *Variant {
	ticker := GoGetArg(goArgs, 0, MkUndefined())
	_ = ticker
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	symbol := *(market).At(MkString("symbol"))
	_ = symbol
	timestamp := this.SafeInteger(ticker, MkString("date"))
	_ = timestamp
	ticker = *(ticker).At(MkString("ticker"))
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
		"change":        this.SafeNumber(ticker, MkString("riseRate")),
		"percentage":    MkUndefined(),
		"average":       MkUndefined(),
		"baseVolume":    this.SafeNumber(ticker, MkString("vol")),
		"quoteVolume":   MkUndefined(),
		"info":          ticker,
	})
}

func (this *Exx) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"currency": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("publicGetTicker"), this.Extend(request, params))
	_ = response
	return this.ParseTicker(response, market)
}

func (this *Exx) FetchTickers(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("publicGetTickers"), params)
	_ = response
	result := MkMap(&VarMap{})
	_ = result
	timestamp := this.Milliseconds()
	_ = timestamp
	ids := GoGetKeys(response)
	_ = ids
	for i := MkInteger(0); IsTrue(OpLw(i, ids.Length)); OpIncr(&i) {
		id := *(ids).At(i)
		_ = id
		if IsTrue(OpNot(OpHasMember(id, *this.At(MkString("markets_by_id"))))) {
			continue
		}
		market := *(*this.At(MkString("markets_by_id"))).At(id)
		_ = market
		symbol := *(market).At(MkString("symbol"))
		_ = symbol
		ticker := MkMap(&VarMap{
			"date":   timestamp,
			"ticker": *(response).At(id),
		})
		_ = ticker
		*(result).At(symbol) = this.ParseTicker(ticker, market)
	}
	return this.FilterByArray(result, MkString("symbol"), symbols)
}

func (this *Exx) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"currency": this.MarketId(symbol)})
	_ = request
	response := this.Call(MkString("publicGetDepth"), this.Extend(request, params))
	_ = response
	timestamp := this.SafeTimestamp(response, MkString("timestamp"))
	_ = timestamp
	return this.ParseOrderBook(response, symbol, timestamp)
}

func (this *Exx) ParseTrade(goArgs ...*Variant) *Variant {
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
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(OpNotEq2(market, MkUndefined())) {
		symbol = *(market).At(MkString("symbol"))
	}
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

func (this *Exx) FetchTrades(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"currency": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("publicGetTrades"), this.Extend(request, params))
	_ = response
	return this.ParseTrades(response, market, since, limit)
}

func (this *Exx) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privateGetGetBalance"), params)
	_ = response
	result := MkMap(&VarMap{"info": response})
	_ = result
	balances := this.SafeValue(response, MkString("funds"))
	_ = balances
	currencies := GoGetKeys(balances)
	_ = currencies
	for i := MkInteger(0); IsTrue(OpLw(i, currencies.Length)); OpIncr(&i) {
		currencyId := *(currencies).At(i)
		_ = currencyId
		balance := *(balances).At(currencyId)
		_ = balance
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		account := this.Account()
		_ = account
		*(account).At(MkString("free")) = this.SafeString(balance, MkString("balance"))
		*(account).At(MkString("used")) = this.SafeString(balance, MkString("freeze"))
		*(account).At(MkString("total")) = this.SafeString(balance, MkString("total"))
		*(result).At(code) = OpCopy(account)
	}
	return this.ParseBalance(result)
}

func (this *Exx) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	symbol := *(market).At(MkString("symbol"))
	_ = symbol
	timestamp := parseInt(*(order).At(MkString("trade_date")))
	_ = timestamp
	price := this.SafeNumber(order, MkString("price"))
	_ = price
	cost := this.SafeNumber(order, MkString("trade_money"))
	_ = cost
	amount := this.SafeNumber(order, MkString("total_amount"))
	_ = amount
	filled := this.SafeNumber(order, MkString("trade_amount"), MkNumber(0.0))
	_ = filled
	status := this.SafeInteger(order, MkString("status"))
	_ = status
	if IsTrue(OpEq2(status, MkInteger(1))) {
		status = MkString("canceled")
	} else {
		if IsTrue(OpEq2(status, MkInteger(2))) {
			status = MkString("closed")
		} else {
			status = MkString("open")
		}
	}
	fee := OpCopy(MkUndefined())
	_ = fee
	if IsTrue(OpHasMember(MkString("fees"), order)) {
		fee = MkMap(&VarMap{
			"cost":     this.SafeNumber(order, MkString("fees")),
			"currency": *(market).At(MkString("quote")),
		})
	}
	return this.SafeOrder(MkMap(&VarMap{
		"id":                 this.SafeString(order, MkString("id")),
		"clientOrderId":      MkUndefined(),
		"datetime":           this.Iso8601(timestamp),
		"timestamp":          timestamp,
		"lastTradeTimestamp": MkUndefined(),
		"status":             status,
		"symbol":             symbol,
		"type":               MkString("limit"),
		"timeInForce":        MkUndefined(),
		"postOnly":           MkUndefined(),
		"side":               *(order).At(MkString("type")),
		"price":              price,
		"stopPrice":          MkUndefined(),
		"cost":               cost,
		"amount":             amount,
		"filled":             filled,
		"remaining":          MkUndefined(),
		"trades":             MkUndefined(),
		"fee":                fee,
		"info":               order,
		"average":            MkUndefined(),
	}))
}

func (this *Exx) CreateOrder(goArgs ...*Variant) *Variant {
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
		"currency": *(market).At(MkString("id")),
		"type":     side,
		"price":    price,
		"amount":   amount,
	})
	_ = request
	response := this.Call(MkString("privateGetOrder"), this.Extend(request, params))
	_ = response
	id := this.SafeString(response, MkString("id"))
	_ = id
	order := this.ParseOrder(MkMap(&VarMap{
		"id":           id,
		"trade_date":   this.Milliseconds(),
		"total_amount": amount,
		"price":        price,
		"type":         side,
		"info":         response,
	}), market)
	_ = order
	return order
}

func (this *Exx) CancelOrder(goArgs ...*Variant) *Variant {
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
		"id":       id,
		"currency": *(market).At(MkString("id")),
	})
	_ = request
	response := this.Call(MkString("privateGetCancel"), this.Extend(request, params))
	_ = response
	return response
}

func (this *Exx) FetchOrder(goArgs ...*Variant) *Variant {
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
		"id":       id,
		"currency": *(market).At(MkString("id")),
	})
	_ = request
	response := this.Call(MkString("privateGetGetOrder"), this.Extend(request, params))
	_ = response
	return this.ParseOrder(response, market)
}

func (this *Exx) FetchOpenOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchOpenOrders() requires a symbol argument"))))
	}
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"currency": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("privateGetGetOpenOrders"), this.Extend(request, params))
	_ = response
	if IsTrue(OpNot(GoIsArray(response))) {
		return MkArray(&VarArray{})
	}
	return this.ParseOrders(response, market, since, limit)
}

func (this *Exx) Nonce(goArgs ...*Variant) *Variant {
	return this.Milliseconds()
}

func (this *Exx) Sign(goArgs ...*Variant) *Variant {
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
	url := OpAdd(*(*(*this.At(MkString("urls"))).At(MkString("api"))).At(api), OpAdd(MkString("/"), path))
	_ = url
	if IsTrue(OpEq2(api, MkString("public"))) {
		if IsTrue(*(GoGetKeys(params)).At(MkString("length"))) {
			OpAddAssign(&url, OpAdd(MkString("?"), this.Urlencode(params)))
		}
	} else {
		this.CheckRequiredCredentials()
		query := this.Urlencode(this.Keysort(this.Extend(MkMap(&VarMap{
			"accesskey": *this.At(MkString("apiKey")),
			"nonce":     this.Nonce(),
		}), params)))
		_ = query
		signed := this.Hmac(this.Encode(query), this.Encode(*this.At(MkString("secret"))), MkString("sha512"))
		_ = signed
		OpAddAssign(&url, OpAdd(MkString("?"), OpAdd(query, OpAdd(MkString("&signature="), signed))))
		headers = MkMap(&VarMap{"Content-Type": MkString("application/x-www-form-urlencoded")})
	}
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Exx) HandleErrors(goArgs ...*Variant) *Variant {
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
	message := this.SafeString(response, MkString("message"))
	_ = message
	feedback := OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))
	_ = feedback
	if IsTrue(OpEq2(code, MkString("100"))) {
		return MkUndefined()
	}
	if IsTrue(OpNotEq2(code, MkUndefined())) {
		this.ThrowExactlyMatchedException(*this.At(MkString("exceptions")), code, feedback)
		if IsTrue(OpEq2(code, MkString("308"))) {
			return MkUndefined()
		} else {
			panic(NewExchangeError(feedback))
		}
	}
	result := this.SafeValue(response, MkString("result"))
	_ = result
	if IsTrue(OpNotEq2(result, MkUndefined())) {
		if IsTrue(OpNot(result)) {
			if IsTrue(OpEq2(message, MkString("æå¡ç«¯å¿ç¢"))) {
				panic(NewExchangeNotAvailable(feedback))
			} else {
				panic(NewExchangeError(feedback))
			}
		}
	}
	return MkUndefined()
}
