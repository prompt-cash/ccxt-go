package ccxt

import ()

type Paymium struct {
	*ExchangeBase
}

var _ Exchange = (*Paymium)(nil)

func init() {
	exchange := &Paymium{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Paymium) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("paymium"),
		"name": MkString("Paymium"),
		"countries": MkArray(&VarArray{
			MkString("FR"),
			MkString("EU"),
		}),
		"rateLimit": MkInteger(2000),
		"version":   MkString("v1"),
		"has": MkMap(&VarMap{
			"CORS":           MkBool(true),
			"fetchBalance":   MkBool(true),
			"fetchTicker":    MkBool(true),
			"fetchTrades":    MkBool(true),
			"fetchOrderBook": MkBool(true),
			"createOrder":    MkBool(true),
			"cancelOrder":    MkBool(true),
		}),
		"urls": MkMap(&VarMap{
			"logo": MkString("https://user-images.githubusercontent.com/51840849/87153930-f0f02200-c2c0-11ea-9c0a-40337375ae89.jpg"),
			"api":  MkString("https://paymium.com/api"),
			"www":  MkString("https://www.paymium.com"),
			"fees": MkString("https://www.paymium.com/page/help/fees"),
			"doc": MkArray(&VarArray{
				MkString("https://github.com/Paymium/api-documentation"),
				MkString("https://www.paymium.com/page/developers"),
			}),
			"referral": MkString("https://www.paymium.com/page/sign-up?referral=eDAzPoRQFMvaAB8sf-qj"),
		}),
		"api": MkMap(&VarMap{
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("countries"),
				MkString("data/{currency}/ticker"),
				MkString("data/{currency}/trades"),
				MkString("data/{currency}/depth"),
				MkString("bitcoin_charts/{id}/trades"),
				MkString("bitcoin_charts/{id}/depth"),
			})}),
			"private": MkMap(&VarMap{
				"get": MkArray(&VarArray{
					MkString("user"),
					MkString("user/addresses"),
					MkString("user/addresses/{address}"),
					MkString("user/orders"),
					MkString("user/orders/{uuid}"),
					MkString("user/price_alerts"),
					MkString("merchant/get_payment/{uuid}"),
				}),
				"post": MkArray(&VarArray{
					MkString("user/addresses"),
					MkString("user/orders"),
					MkString("user/withdrawals"),
					MkString("user/email_transfers"),
					MkString("user/payment_requests"),
					MkString("user/price_alerts"),
					MkString("merchant/create_payment"),
				}),
				"delete": MkArray(&VarArray{
					MkString("user/orders/{uuid}"),
					MkString("user/orders/{uuid}/cancel"),
					MkString("user/price_alerts/{id}"),
				}),
			}),
		}),
		"markets": MkMap(&VarMap{"BTC/EUR": MkMap(&VarMap{
			"id":      MkString("eur"),
			"symbol":  MkString("BTC/EUR"),
			"base":    MkString("BTC"),
			"quote":   MkString("EUR"),
			"baseId":  MkString("btc"),
			"quoteId": MkString("eur"),
		})}),
		"fees": MkMap(&VarMap{"trading": MkMap(&VarMap{
			"maker": this.ParseNumber(MkString("0.002")),
			"taker": this.ParseNumber(MkString("0.002")),
		})}),
	}))
}

func (this *Paymium) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privateGetUser"), params)
	_ = response
	result := MkMap(&VarMap{"info": response})
	_ = result
	currencies := GoGetKeys(*this.At(MkString("currencies")))
	_ = currencies
	for i := MkInteger(0); IsTrue(OpLw(i, currencies.Length)); OpIncr(&i) {
		code := *(currencies).At(i)
		_ = code
		currency := this.Currency(code)
		_ = currency
		currencyId := *(currency).At(MkString("id"))
		_ = currencyId
		free := OpAdd(MkString("balance_"), currencyId)
		_ = free
		if IsTrue(OpHasMember(free, response)) {
			account := this.Account()
			_ = account
			used := OpAdd(MkString("locked_"), currencyId)
			_ = used
			*(account).At(MkString("free")) = this.SafeString(response, free)
			*(account).At(MkString("used")) = this.SafeString(response, used)
			*(result).At(code) = OpCopy(account)
		}
	}
	return this.ParseBalance(result)
}

func (this *Paymium) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"currency": this.MarketId(symbol)})
	_ = request
	response := this.Call(MkString("publicGetDataCurrencyDepth"), this.Extend(request, params))
	_ = response
	return this.ParseOrderBook(response, symbol, MkUndefined(), MkString("bids"), MkString("asks"), MkString("price"), MkString("amount"))
}

func (this *Paymium) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"currency": this.MarketId(symbol)})
	_ = request
	ticker := this.Call(MkString("publicGetDataCurrencyTicker"), this.Extend(request, params))
	_ = ticker
	timestamp := this.SafeTimestamp(ticker, MkString("at"))
	_ = timestamp
	vwap := this.SafeNumber(ticker, MkString("vwap"))
	_ = vwap
	baseVolume := this.SafeNumber(ticker, MkString("volume"))
	_ = baseVolume
	quoteVolume := OpCopy(MkUndefined())
	_ = quoteVolume
	if IsTrue(OpAnd(OpNotEq2(baseVolume, MkUndefined()), OpNotEq2(vwap, MkUndefined()))) {
		quoteVolume = OpMulti(baseVolume, vwap)
	}
	last := this.SafeNumber(ticker, MkString("price"))
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
		"open":          this.SafeNumber(ticker, MkString("open")),
		"close":         last,
		"last":          last,
		"previousClose": MkUndefined(),
		"change":        MkUndefined(),
		"percentage":    this.SafeNumber(ticker, MkString("variation")),
		"average":       MkUndefined(),
		"baseVolume":    baseVolume,
		"quoteVolume":   quoteVolume,
		"info":          ticker,
	})
}

func (this *Paymium) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.SafeTimestamp(trade, MkString("created_at_int"))
	_ = timestamp
	id := this.SafeString(trade, MkString("uuid"))
	_ = id
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(OpNotEq2(market, MkUndefined())) {
		symbol = *(market).At(MkString("symbol"))
	}
	side := this.SafeString(trade, MkString("side"))
	_ = side
	priceString := this.SafeString(trade, MkString("price"))
	_ = priceString
	amountField := OpAdd(MkString("traded_"), (*(market).At(MkString("base"))).Call(MkString("toLowerCase")))
	_ = amountField
	amountString := this.SafeString(trade, amountField)
	_ = amountString
	price := this.ParseNumber(priceString)
	_ = price
	amount := this.ParseNumber(amountString)
	_ = amount
	cost := this.ParseNumber(Precise.StringMul(priceString, amountString))
	_ = cost
	return MkMap(&VarMap{
		"info":         trade,
		"id":           id,
		"order":        MkUndefined(),
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       symbol,
		"type":         MkUndefined(),
		"side":         side,
		"takerOrMaker": MkUndefined(),
		"price":        price,
		"amount":       amount,
		"cost":         cost,
		"fee":          MkUndefined(),
	})
}

func (this *Paymium) FetchTrades(goArgs ...*Variant) *Variant {
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
	response := this.Call(MkString("publicGetDataCurrencyTrades"), this.Extend(request, params))
	_ = response
	return this.ParseTrades(response, market, since, limit)
}

func (this *Paymium) CreateOrder(goArgs ...*Variant) *Variant {
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
		"type":      OpAdd(this.Capitalize(type_), MkString("Order")),
		"currency":  this.MarketId(symbol),
		"direction": side,
		"amount":    amount,
	})
	_ = request
	if IsTrue(OpNotEq2(type_, MkString("market"))) {
		*(request).At(MkString("price")) = OpCopy(price)
	}
	response := this.Call(MkString("privatePostUserOrders"), this.Extend(request, params))
	_ = response
	return MkMap(&VarMap{
		"info": response,
		"id":   *(response).At(MkString("uuid")),
	})
}

func (this *Paymium) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	request := MkMap(&VarMap{"uuid": id})
	_ = request
	return this.Call(MkString("privateDeleteUserOrdersUuidCancel"), this.Extend(request, params))
}

func (this *Paymium) Sign(goArgs ...*Variant) *Variant {
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
	url := OpAdd(*(*this.At(MkString("urls"))).At(MkString("api")), OpAdd(MkString("/"), OpAdd(*this.At(MkString("version")), OpAdd(MkString("/"), this.ImplodeParams(path, params)))))
	_ = url
	query := this.Omit(params, this.ExtractParams(path))
	_ = query
	if IsTrue(OpEq2(api, MkString("public"))) {
		if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
			OpAddAssign(&url, OpAdd(MkString("?"), this.Urlencode(query)))
		}
	} else {
		this.CheckRequiredCredentials()
		nonce := (this.Nonce()).Call(MkString("toString"))
		_ = nonce
		auth := OpAdd(nonce, url)
		_ = auth
		headers = MkMap(&VarMap{
			"Api-Key":   *this.At(MkString("apiKey")),
			"Api-Nonce": nonce,
		})
		if IsTrue(OpEq2(method, MkString("POST"))) {
			if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
				body = this.Json(query)
				OpAddAssign(&auth, body)
				*(headers).At(MkString("Content-Type")) = MkString("application/json")
			}
		} else {
			if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
				queryString := this.Urlencode(query)
				_ = queryString
				OpAddAssign(&auth, queryString)
				OpAddAssign(&url, OpAdd(MkString("?"), queryString))
			}
		}
		*(headers).At(MkString("Api-Signature")) = this.Hmac(this.Encode(auth), this.Encode(*this.At(MkString("secret"))))
	}
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Paymium) HandleErrors(goArgs ...*Variant) *Variant {
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
	errors := this.SafeValue(response, MkString("errors"))
	_ = errors
	if IsTrue(OpNotEq2(errors, MkUndefined())) {
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), this.Json(response)))))
	}
	return MkUndefined()
}
