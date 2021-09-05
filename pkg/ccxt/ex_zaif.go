package ccxt

import ()

type Zaif struct {
	*ExchangeBase
}

var _ Exchange = (*Zaif)(nil)

func init() {
	exchange := &Zaif{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Zaif) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("zaif"),
		"name": MkString("Zaif"),
		"countries": MkArray(&VarArray{
			MkString("JP"),
		}),
		"rateLimit": MkInteger(2000),
		"version":   MkString("1"),
		"has": MkMap(&VarMap{
			"cancelOrder":       MkBool(true),
			"CORS":              MkBool(false),
			"createMarketOrder": MkBool(false),
			"createOrder":       MkBool(true),
			"fetchBalance":      MkBool(true),
			"fetchClosedOrders": MkBool(true),
			"fetchMarkets":      MkBool(true),
			"fetchOrderBook":    MkBool(true),
			"fetchOpenOrders":   MkBool(true),
			"fetchTicker":       MkBool(true),
			"fetchTrades":       MkBool(true),
			"withdraw":          MkBool(true),
		}),
		"urls": MkMap(&VarMap{
			"logo": MkString("https://user-images.githubusercontent.com/1294454/27766927-39ca2ada-5eeb-11e7-972f-1b4199518ca6.jpg"),
			"api":  MkString("https://api.zaif.jp"),
			"www":  MkString("https://zaif.jp"),
			"doc": MkArray(&VarArray{
				MkString("https://techbureau-api-document.readthedocs.io/ja/latest/index.html"),
				MkString("https://corp.zaif.jp/api-docs"),
				MkString("https://corp.zaif.jp/api-docs/api_links"),
				MkString("https://www.npmjs.com/package/zaif.jp"),
				MkString("https://github.com/you21979/node-zaif"),
			}),
			"fees": MkString("https://zaif.jp/fee?lang=en"),
		}),
		"fees": MkMap(&VarMap{"trading": MkMap(&VarMap{
			"percentage": MkBool(true),
			"taker":      this.ParseNumber(MkString("0.001")),
			"maker":      this.ParseNumber(MkString("0")),
		})}),
		"api": MkMap(&VarMap{
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("depth/{pair}"),
				MkString("currencies/{pair}"),
				MkString("currencies/all"),
				MkString("currency_pairs/{pair}"),
				MkString("currency_pairs/all"),
				MkString("last_price/{pair}"),
				MkString("ticker/{pair}"),
				MkString("trades/{pair}"),
			})}),
			"private": MkMap(&VarMap{"post": MkArray(&VarArray{
				MkString("active_orders"),
				MkString("cancel_order"),
				MkString("deposit_history"),
				MkString("get_id_info"),
				MkString("get_info"),
				MkString("get_info2"),
				MkString("get_personal_info"),
				MkString("trade"),
				MkString("trade_history"),
				MkString("withdraw"),
				MkString("withdraw_history"),
			})}),
			"ecapi": MkMap(&VarMap{"post": MkArray(&VarArray{
				MkString("createInvoice"),
				MkString("getInvoice"),
				MkString("getInvoiceIdsByOrderNumber"),
				MkString("cancelInvoice"),
			})}),
			"tlapi": MkMap(&VarMap{"post": MkArray(&VarArray{
				MkString("get_positions"),
				MkString("position_history"),
				MkString("active_positions"),
				MkString("create_position"),
				MkString("change_position"),
				MkString("cancel_position"),
			})}),
			"fapi": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("groups/{group_id}"),
				MkString("last_price/{group_id}/{pair}"),
				MkString("ticker/{group_id}/{pair}"),
				MkString("trades/{group_id}/{pair}"),
				MkString("depth/{group_id}/{pair}"),
			})}),
		}),
		"options": MkMap(&VarMap{"fees": MkMap(&VarMap{
			"BTC/JPY": MkMap(&VarMap{
				"maker": MkInteger(0),
				"taker": MkInteger(0),
			}),
			"BCH/JPY": MkMap(&VarMap{
				"maker": MkInteger(0),
				"taker": OpDiv(MkNumber(0.3), MkInteger(100)),
			}),
			"BCH/BTC": MkMap(&VarMap{
				"maker": MkInteger(0),
				"taker": OpDiv(MkNumber(0.3), MkInteger(100)),
			}),
			"PEPECASH/JPY": MkMap(&VarMap{
				"maker": MkInteger(0),
				"taker": OpDiv(MkNumber(0.01), MkInteger(100)),
			}),
			"PEPECASH/BT": MkMap(&VarMap{
				"maker": MkInteger(0),
				"taker": OpDiv(MkNumber(0.01), MkInteger(100)),
			}),
		})}),
		"exceptions": MkMap(&VarMap{
			"exact": MkMap(&VarMap{"unsupported currency_pair": BadRequest}),
			"broad": MkMap(&VarMap{}),
		}),
	}))
}

func (this *Zaif) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	markets := this.Call(MkString("publicGetCurrencyPairsAll"), params)
	_ = markets
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, markets.Length)); OpIncr(&i) {
		market := *(markets).At(i)
		_ = market
		id := this.SafeString(market, MkString("currency_pair"))
		_ = id
		name := this.SafeString(market, MkString("name"))
		_ = name
		baseId := MkUndefined()
		_ = baseId
		quoteId := MkUndefined()
		_ = quoteId
		ArrayUnpack(name.Split(MkString("/")), &baseId, &quoteId)
		base := this.SafeCurrencyCode(baseId)
		_ = base
		quote := this.SafeCurrencyCode(quoteId)
		_ = quote
		symbol := OpAdd(base, OpAdd(MkString("/"), quote))
		_ = symbol
		precision := MkMap(&VarMap{
			"amount": OpNeg(Math.Log10(this.SafeNumber(market, MkString("item_unit_step")))),
			"price":  this.SafeInteger(market, MkString("aux_unit_point")),
		})
		_ = precision
		fees := this.SafeValue(*(*this.At(MkString("options"))).At(MkString("fees")), symbol, *(*this.At(MkString("fees"))).At(MkString("trading")))
		_ = fees
		taker := *(fees).At(MkString("taker"))
		_ = taker
		maker := *(fees).At(MkString("maker"))
		_ = maker
		result.Push(MkMap(&VarMap{
			"id":        id,
			"symbol":    symbol,
			"base":      base,
			"quote":     quote,
			"baseId":    baseId,
			"quoteId":   quoteId,
			"active":    MkBool(true),
			"precision": precision,
			"taker":     taker,
			"maker":     maker,
			"limits": MkMap(&VarMap{
				"amount": MkMap(&VarMap{
					"min": this.SafeNumber(market, MkString("item_unit_min")),
					"max": MkUndefined(),
				}),
				"price": MkMap(&VarMap{
					"min": this.SafeNumber(market, MkString("aux_unit_min")),
					"max": MkUndefined(),
				}),
				"cost": MkMap(&VarMap{
					"min": MkUndefined(),
					"max": MkUndefined(),
				}),
			}),
			"info": market,
		}))
	}
	return result
}

func (this *Zaif) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privatePostGetInfo"), params)
	_ = response
	balances := this.SafeValue(response, MkString("return"), MkMap(&VarMap{}))
	_ = balances
	deposit := this.SafeValue(balances, MkString("deposit"))
	_ = deposit
	result := MkMap(&VarMap{
		"info":      response,
		"timestamp": MkUndefined(),
		"datetime":  MkUndefined(),
	})
	_ = result
	funds := this.SafeValue(balances, MkString("funds"), MkMap(&VarMap{}))
	_ = funds
	currencyIds := GoGetKeys(funds)
	_ = currencyIds
	for i := MkInteger(0); IsTrue(OpLw(i, currencyIds.Length)); OpIncr(&i) {
		currencyId := *(currencyIds).At(i)
		_ = currencyId
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		balance := this.SafeString(funds, currencyId)
		_ = balance
		account := this.Account()
		_ = account
		*(account).At(MkString("free")) = OpCopy(balance)
		*(account).At(MkString("total")) = OpCopy(balance)
		if IsTrue(OpNotEq2(deposit, MkUndefined())) {
			if IsTrue(OpHasMember(currencyId, deposit)) {
				*(account).At(MkString("total")) = this.SafeString(deposit, currencyId)
			}
		}
		*(result).At(code) = OpCopy(account)
	}
	return this.ParseBalance(result)
}

func (this *Zaif) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"pair": this.MarketId(symbol)})
	_ = request
	response := this.Call(MkString("publicGetDepthPair"), this.Extend(request, params))
	_ = response
	return this.ParseOrderBook(response, symbol)
}

func (this *Zaif) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"pair": this.MarketId(symbol)})
	_ = request
	ticker := this.Call(MkString("publicGetTickerPair"), this.Extend(request, params))
	_ = ticker
	timestamp := this.Milliseconds()
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
	last := this.SafeNumber(ticker, MkString("last"))
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
		"open":          MkUndefined(),
		"close":         last,
		"last":          last,
		"previousClose": MkUndefined(),
		"change":        MkUndefined(),
		"percentage":    MkUndefined(),
		"average":       MkUndefined(),
		"baseVolume":    baseVolume,
		"quoteVolume":   quoteVolume,
		"info":          ticker,
	})
}

func (this *Zaif) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	side := this.SafeString(trade, MkString("trade_type"))
	_ = side
	side = OpTrinary(OpEq2(side, MkString("bid")), MkString("buy"), MkString("sell"))
	timestamp := this.SafeTimestamp(trade, MkString("date"))
	_ = timestamp
	id := this.SafeString2(trade, MkString("id"), MkString("tid"))
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
	marketId := this.SafeString(trade, MkString("currency_pair"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market, MkString("_"))
	_ = symbol
	return MkMap(&VarMap{
		"id":           id,
		"info":         trade,
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       symbol,
		"type":         MkUndefined(),
		"side":         side,
		"order":        MkUndefined(),
		"takerOrMaker": MkUndefined(),
		"price":        price,
		"amount":       amount,
		"cost":         cost,
		"fee":          MkUndefined(),
	})
}

func (this *Zaif) FetchTrades(goArgs ...*Variant) *Variant {
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
	response := this.Call(MkString("publicGetTradesPair"), this.Extend(request, params))
	_ = response
	numTrades := OpCopy(response.Length)
	_ = numTrades
	if IsTrue(OpEq2(numTrades, MkInteger(1))) {
		firstTrade := *(response).At(MkInteger(0))
		_ = firstTrade
		if IsTrue(OpNot(*(GoGetKeys(firstTrade)).At(MkString("length")))) {
			response = MkArray(&VarArray{})
		}
	}
	return this.ParseTrades(response, market, since, limit)
}

func (this *Zaif) CreateOrder(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpNotEq2(type_, MkString("limit"))) {
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), MkString(" allows limit orders only"))))
	}
	request := MkMap(&VarMap{
		"currency_pair": this.MarketId(symbol),
		"action":        OpTrinary(OpEq2(side, MkString("buy")), MkString("bid"), MkString("ask")),
		"amount":        amount,
		"price":         price,
	})
	_ = request
	response := this.Call(MkString("privatePostTrade"), this.Extend(request, params))
	_ = response
	return MkMap(&VarMap{
		"info": response,
		"id":   (*(*(response).At(MkString("return"))).At(MkString("order_id"))).Call(MkString("toString")),
	})
}

func (this *Zaif) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	request := MkMap(&VarMap{"order_id": id})
	_ = request
	return this.Call(MkString("privatePostCancelOrder"), this.Extend(request, params))
}

func (this *Zaif) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	side := this.SafeString(order, MkString("action"))
	_ = side
	side = OpTrinary(OpEq2(side, MkString("bid")), MkString("buy"), MkString("sell"))
	timestamp := this.SafeTimestamp(order, MkString("timestamp"))
	_ = timestamp
	marketId := this.SafeString(order, MkString("currency_pair"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market, MkString("_"))
	_ = symbol
	price := this.SafeNumber(order, MkString("price"))
	_ = price
	amount := this.SafeNumber(order, MkString("amount"))
	_ = amount
	id := this.SafeString(order, MkString("id"))
	_ = id
	return this.SafeOrder(MkMap(&VarMap{
		"id":                 id,
		"clientOrderId":      MkUndefined(),
		"timestamp":          timestamp,
		"datetime":           this.Iso8601(timestamp),
		"lastTradeTimestamp": MkUndefined(),
		"status":             MkString("open"),
		"symbol":             symbol,
		"type":               MkString("limit"),
		"timeInForce":        MkUndefined(),
		"postOnly":           MkUndefined(),
		"side":               side,
		"price":              price,
		"stopPrice":          MkUndefined(),
		"cost":               MkUndefined(),
		"amount":             amount,
		"filled":             MkUndefined(),
		"remaining":          MkUndefined(),
		"trades":             MkUndefined(),
		"fee":                MkUndefined(),
		"info":               order,
		"average":            MkUndefined(),
	}))
}

func (this *Zaif) FetchOpenOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := OpCopy(MkUndefined())
	_ = market
	request := MkMap(&VarMap{})
	_ = request
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
		*(request).At(MkString("currency_pair")) = *(market).At(MkString("id"))
	}
	response := this.Call(MkString("privatePostActiveOrders"), this.Extend(request, params))
	_ = response
	return this.ParseOrders(*(response).At(MkString("return")), market, since, limit)
}

func (this *Zaif) FetchClosedOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := OpCopy(MkUndefined())
	_ = market
	request := MkMap(&VarMap{})
	_ = request
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
		*(request).At(MkString("currency_pair")) = *(market).At(MkString("id"))
	}
	response := this.Call(MkString("privatePostTradeHistory"), this.Extend(request, params))
	_ = response
	return this.ParseOrders(*(response).At(MkString("return")), market, since, limit)
}

func (this *Zaif) Withdraw(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpEq2(code, MkString("JPY"))) {
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" withdraw() does not allow "), OpAdd(code, MkString(" withdrawals"))))))
	}
	request := MkMap(&VarMap{
		"currency": *(currency).At(MkString("id")),
		"amount":   amount,
		"address":  address,
	})
	_ = request
	if IsTrue(OpNotEq2(tag, MkUndefined())) {
		*(request).At(MkString("message")) = OpCopy(tag)
	}
	result := this.Call(MkString("privatePostWithdraw"), this.Extend(request, params))
	_ = result
	return MkMap(&VarMap{
		"info": result,
		"id":   *(*(result).At(MkString("return"))).At(MkString("txid")),
		"fee":  *(*(result).At(MkString("return"))).At(MkString("fee")),
	})
}

func (this *Zaif) Nonce(goArgs ...*Variant) *Variant {
	nonce := parseFloat(OpDiv(this.Milliseconds(), MkInteger(1000)))
	_ = nonce
	return nonce.ToFixed(MkInteger(8))
}

func (this *Zaif) Sign(goArgs ...*Variant) *Variant {
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
	url := OpAdd(*(*this.At(MkString("urls"))).At(MkString("api")), MkString("/"))
	_ = url
	if IsTrue(OpEq2(api, MkString("public"))) {
		OpAddAssign(&url, OpAdd(MkString("api/"), OpAdd(*this.At(MkString("version")), OpAdd(MkString("/"), this.ImplodeParams(path, params)))))
	} else {
		if IsTrue(OpEq2(api, MkString("fapi"))) {
			OpAddAssign(&url, OpAdd(MkString("fapi/"), OpAdd(*this.At(MkString("version")), OpAdd(MkString("/"), this.ImplodeParams(path, params)))))
		} else {
			this.CheckRequiredCredentials()
			if IsTrue(OpEq2(api, MkString("ecapi"))) {
				OpAddAssign(&url, MkString("ecapi"))
			} else {
				if IsTrue(OpEq2(api, MkString("tlapi"))) {
					OpAddAssign(&url, MkString("tlapi"))
				} else {
					OpAddAssign(&url, MkString("tapi"))
				}
			}
			nonce := this.Nonce()
			_ = nonce
			body = this.Urlencode(this.Extend(MkMap(&VarMap{
				"method": path,
				"nonce":  nonce,
			}), params))
			headers = MkMap(&VarMap{
				"Content-Type": MkString("application/x-www-form-urlencoded"),
				"Key":          *this.At(MkString("apiKey")),
				"Sign":         this.Hmac(this.Encode(body), this.Encode(*this.At(MkString("secret"))), MkString("sha512")),
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

func (this *Zaif) HandleErrors(goArgs ...*Variant) *Variant {
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
	feedback := OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))
	_ = feedback
	error := this.SafeString(response, MkString("error"))
	_ = error
	if IsTrue(OpNotEq2(error, MkUndefined())) {
		this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("exact")), error, feedback)
		this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("broad")), error, feedback)
		panic(NewExchangeError(feedback))
	}
	success := this.SafeValue(response, MkString("success"), MkBool(true))
	_ = success
	if IsTrue(OpNot(success)) {
		panic(NewExchangeError(feedback))
	}
	return MkUndefined()
}
