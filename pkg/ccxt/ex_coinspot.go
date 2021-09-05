package ccxt

import ()

type Coinspot struct {
	*ExchangeBase
}

var _ Exchange = (*Coinspot)(nil)

func init() {
	exchange := &Coinspot{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Coinspot) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("coinspot"),
		"name": MkString("CoinSpot"),
		"countries": MkArray(&VarArray{
			MkString("AU"),
		}),
		"rateLimit": MkInteger(1000),
		"has": MkMap(&VarMap{
			"cancelOrder":       MkBool(false),
			"CORS":              MkBool(false),
			"createMarketOrder": MkBool(false),
			"createOrder":       MkBool(true),
			"fetchBalance":      MkBool(true),
			"fetchOrderBook":    MkBool(true),
			"fetchTicker":       MkBool(true),
			"fetchTrades":       MkBool(true),
		}),
		"urls": MkMap(&VarMap{
			"logo": MkString("https://user-images.githubusercontent.com/1294454/28208429-3cacdf9a-6896-11e7-854e-4c79a772a30f.jpg"),
			"api": MkMap(&VarMap{
				"public":  MkString("https://www.coinspot.com.au/pubapi"),
				"private": MkString("https://www.coinspot.com.au/api"),
			}),
			"www":      MkString("https://www.coinspot.com.au"),
			"doc":      MkString("https://www.coinspot.com.au/api"),
			"referral": MkString("https://www.coinspot.com.au/register?code=PJURCU"),
		}),
		"api": MkMap(&VarMap{
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("latest"),
			})}),
			"private": MkMap(&VarMap{"post": MkArray(&VarArray{
				MkString("orders"),
				MkString("orders/history"),
				MkString("my/coin/deposit"),
				MkString("my/coin/send"),
				MkString("quote/buy"),
				MkString("quote/sell"),
				MkString("my/balances"),
				MkString("my/orders"),
				MkString("my/buy"),
				MkString("my/sell"),
				MkString("my/buy/cancel"),
				MkString("my/sell/cancel"),
				MkString("ro/my/balances"),
				MkString("ro/my/balances/{cointype}"),
				MkString("ro/my/deposits"),
				MkString("ro/my/withdrawals"),
				MkString("ro/my/transactions"),
				MkString("ro/my/transactions/{cointype}"),
				MkString("ro/my/transactions/open"),
				MkString("ro/my/transactions/{cointype}/open"),
				MkString("ro/my/sendreceive"),
				MkString("ro/my/affiliatepayments"),
				MkString("ro/my/referralpayments"),
			})}),
		}),
		"markets": MkMap(&VarMap{
			"BTC/AUD": MkMap(&VarMap{
				"id":      MkString("btc"),
				"symbol":  MkString("BTC/AUD"),
				"base":    MkString("BTC"),
				"quote":   MkString("AUD"),
				"baseId":  MkString("btc"),
				"quoteId": MkString("aud"),
			}),
			"ETH/AUD": MkMap(&VarMap{
				"id":      MkString("eth"),
				"symbol":  MkString("ETH/AUD"),
				"base":    MkString("ETH"),
				"quote":   MkString("AUD"),
				"baseId":  MkString("eth"),
				"quoteId": MkString("aud"),
			}),
			"XRP/AUD": MkMap(&VarMap{
				"id":      MkString("xrp"),
				"symbol":  MkString("XRP/AUD"),
				"base":    MkString("XRP"),
				"quote":   MkString("AUD"),
				"baseId":  MkString("xrp"),
				"quoteId": MkString("aud"),
			}),
			"LTC/AUD": MkMap(&VarMap{
				"id":      MkString("ltc"),
				"symbol":  MkString("LTC/AUD"),
				"base":    MkString("LTC"),
				"quote":   MkString("AUD"),
				"baseId":  MkString("ltc"),
				"quoteId": MkString("aud"),
			}),
			"DOGE/AUD": MkMap(&VarMap{
				"id":      MkString("doge"),
				"symbol":  MkString("DOGE/AUD"),
				"base":    MkString("DOGE"),
				"quote":   MkString("AUD"),
				"baseId":  MkString("doge"),
				"quoteId": MkString("aud"),
			}),
			"RFOX/AUD": MkMap(&VarMap{
				"id":      MkString("rfox"),
				"symbol":  MkString("RFOX/AUD"),
				"base":    MkString("RFOX"),
				"quote":   MkString("AUD"),
				"baseId":  MkString("rfox"),
				"quoteId": MkString("aud"),
			}),
			"POWR/AUD": MkMap(&VarMap{
				"id":      MkString("powr"),
				"symbol":  MkString("POWR/AUD"),
				"base":    MkString("POWR"),
				"quote":   MkString("AUD"),
				"baseId":  MkString("powr"),
				"quoteId": MkString("aud"),
			}),
			"NEO/AUD": MkMap(&VarMap{
				"id":      MkString("neo"),
				"symbol":  MkString("NEO/AUD"),
				"base":    MkString("NEO"),
				"quote":   MkString("AUD"),
				"baseId":  MkString("neo"),
				"quoteId": MkString("aud"),
			}),
			"TRX/AUD": MkMap(&VarMap{
				"id":      MkString("trx"),
				"symbol":  MkString("TRX/AUD"),
				"base":    MkString("TRX"),
				"quote":   MkString("AUD"),
				"baseId":  MkString("trx"),
				"quoteId": MkString("aud"),
			}),
			"EOS/AUD": MkMap(&VarMap{
				"id":      MkString("eos"),
				"symbol":  MkString("EOS/AUD"),
				"base":    MkString("EOS"),
				"quote":   MkString("AUD"),
				"baseId":  MkString("eos"),
				"quoteId": MkString("aud"),
			}),
			"XLM/AUD": MkMap(&VarMap{
				"id":      MkString("xlm"),
				"symbol":  MkString("XLM/AUD"),
				"base":    MkString("XLM"),
				"quote":   MkString("AUD"),
				"baseId":  MkString("xlm"),
				"quoteId": MkString("aud"),
			}),
			"RHOC/AUD": MkMap(&VarMap{
				"id":      MkString("rhoc"),
				"symbol":  MkString("RHOC/AUD"),
				"base":    MkString("RHOC"),
				"quote":   MkString("AUD"),
				"baseId":  MkString("rhoc"),
				"quoteId": MkString("aud"),
			}),
			"GAS/AUD": MkMap(&VarMap{
				"id":      MkString("gas"),
				"symbol":  MkString("GAS/AUD"),
				"base":    MkString("GAS"),
				"quote":   MkString("AUD"),
				"baseId":  MkString("gas"),
				"quoteId": MkString("aud"),
			}),
		}),
		"commonCurrencies": MkMap(&VarMap{"DRK": MkString("DASH")}),
		"options":          MkMap(&VarMap{"fetchBalance": MkString("private_post_my_balances")}),
	}))
}

func (this *Coinspot) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	method := this.SafeString(*this.At(MkString("options")), MkString("fetchBalance"), MkString("private_post_my_balances"))
	_ = method
	response := this.Call(method, params)
	_ = response
	result := MkMap(&VarMap{"info": response})
	_ = result
	balances := this.SafeValue2(response, MkString("balance"), MkString("balances"))
	_ = balances
	if IsTrue(GoIsArray(balances)) {
		for i := MkInteger(0); IsTrue(OpLw(i, balances.Length)); OpIncr(&i) {
			currencies := *(balances).At(i)
			_ = currencies
			currencyIds := GoGetKeys(currencies)
			_ = currencyIds
			for j := MkInteger(0); IsTrue(OpLw(j, currencyIds.Length)); OpIncr(&j) {
				currencyId := *(currencyIds).At(j)
				_ = currencyId
				balance := *(currencies).At(currencyId)
				_ = balance
				code := this.SafeCurrencyCode(currencyId)
				_ = code
				account := this.Account()
				_ = account
				*(account).At(MkString("total")) = this.SafeString(balance, MkString("balance"))
				*(result).At(code) = OpCopy(account)
			}
		}
	} else {
		currencyIds := GoGetKeys(balances)
		_ = currencyIds
		for i := MkInteger(0); IsTrue(OpLw(i, currencyIds.Length)); OpIncr(&i) {
			currencyId := *(currencyIds).At(i)
			_ = currencyId
			code := this.SafeCurrencyCode(currencyId)
			_ = code
			account := this.Account()
			_ = account
			*(account).At(MkString("total")) = this.SafeString(balances, currencyId)
			*(result).At(code) = OpCopy(account)
		}
	}
	return this.ParseBalance(result)
}

func (this *Coinspot) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"cointype": *(market).At(MkString("id"))})
	_ = request
	orderbook := this.Call(MkString("privatePostOrders"), this.Extend(request, params))
	_ = orderbook
	return this.ParseOrderBook(orderbook, symbol, MkUndefined(), MkString("buyorders"), MkString("sellorders"), MkString("rate"), MkString("amount"))
}

func (this *Coinspot) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("publicGetLatest"), params)
	_ = response
	id := this.MarketId(symbol)
	_ = id
	id = id.ToLowerCase()
	ticker := *(*(response).At(MkString("prices"))).At(id)
	_ = ticker
	timestamp := this.Milliseconds()
	_ = timestamp
	last := this.SafeNumber(ticker, MkString("last"))
	_ = last
	return MkMap(&VarMap{
		"symbol":        symbol,
		"timestamp":     timestamp,
		"datetime":      this.Iso8601(timestamp),
		"high":          MkUndefined(),
		"low":           MkUndefined(),
		"bid":           this.SafeNumber(ticker, MkString("bid")),
		"bidVolume":     MkUndefined(),
		"ask":           this.SafeNumber(ticker, MkString("ask")),
		"askVolume":     MkUndefined(),
		"vwap":          MkUndefined(),
		"open":          MkUndefined(),
		"close":         last,
		"last":          last,
		"previousClose": MkUndefined(),
		"change":        MkUndefined(),
		"percentage":    MkUndefined(),
		"average":       MkUndefined(),
		"baseVolume":    MkUndefined(),
		"quoteVolume":   MkUndefined(),
		"info":          ticker,
	})
}

func (this *Coinspot) FetchTrades(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"cointype": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("privatePostOrdersHistory"), this.Extend(request, params))
	_ = response
	trades := this.SafeValue(response, MkString("orders"), MkArray(&VarArray{}))
	_ = trades
	return this.ParseTrades(trades, market, since, limit)
}

func (this *Coinspot) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	priceString := this.SafeString(trade, MkString("rate"))
	_ = priceString
	amountString := this.SafeString(trade, MkString("amount"))
	_ = amountString
	price := this.ParseNumber(priceString)
	_ = price
	amount := this.ParseNumber(amountString)
	_ = amount
	cost := this.SafeNumber(trade, MkString("total"))
	_ = cost
	if IsTrue(OpEq2(cost, MkUndefined())) {
		cost = this.ParseNumber(Precise.StringMul(priceString, amountString))
	}
	timestamp := this.SafeInteger(trade, MkString("solddate"))
	_ = timestamp
	marketId := this.SafeString(trade, MkString("market"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market, MkString("/"))
	_ = symbol
	return MkMap(&VarMap{
		"info":         trade,
		"id":           MkUndefined(),
		"symbol":       symbol,
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"order":        MkUndefined(),
		"type":         MkUndefined(),
		"side":         MkUndefined(),
		"takerOrMaker": MkUndefined(),
		"price":        price,
		"amount":       amount,
		"cost":         cost,
		"fee":          MkUndefined(),
	})
}

func (this *Coinspot) CreateOrder(goArgs ...*Variant) *Variant {
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
	method := OpAdd(MkString("privatePostMy"), this.Capitalize(side))
	_ = method
	if IsTrue(OpEq2(type_, MkString("market"))) {
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), MkString(" allows limit orders only"))))
	}
	request := MkMap(&VarMap{
		"cointype": this.MarketId(symbol),
		"amount":   amount,
		"rate":     price,
	})
	_ = request
	return this.Call(method, this.Extend(request, params))
}

func (this *Coinspot) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	side := this.SafeString(params, MkString("side"))
	_ = side
	if IsTrue(OpAnd(OpNotEq2(side, MkString("buy")), OpNotEq2(side, MkString("sell")))) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" cancelOrder() requires a side parameter, \"buy\" or \"sell\""))))
	}
	params = this.Omit(params, MkString("side"))
	method := OpAdd(MkString("privatePostMy"), OpAdd(this.Capitalize(side), MkString("Cancel")))
	_ = method
	request := MkMap(&VarMap{"id": id})
	_ = request
	return this.Call(method, this.Extend(request, params))
}

func (this *Coinspot) Sign(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpNot(*this.At(MkString("apiKey")))) {
		panic(NewAuthenticationError(OpAdd(*this.At(MkString("id")), MkString(" requires apiKey for all requests"))))
	}
	url := OpAdd(*(*(*this.At(MkString("urls"))).At(MkString("api"))).At(api), OpAdd(MkString("/"), path))
	_ = url
	if IsTrue(OpEq2(api, MkString("private"))) {
		this.CheckRequiredCredentials()
		nonce := this.Nonce()
		_ = nonce
		body = this.Json(this.Extend(MkMap(&VarMap{"nonce": nonce}), params))
		headers = MkMap(&VarMap{
			"Content-Type": MkString("application/json"),
			"key":          *this.At(MkString("apiKey")),
			"sign":         this.Hmac(this.Encode(body), this.Encode(*this.At(MkString("secret"))), MkString("sha512")),
		})
	}
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}
