package ccxt

import ()

type Bit2c struct {
	*ExchangeBase
}

var _ Exchange = (*Bit2c)(nil)

func init() {
	exchange := &Bit2c{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Bit2c) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("bit2c"),
		"name": MkString("Bit2C"),
		"countries": MkArray(&VarArray{
			MkString("IL"),
		}),
		"rateLimit": MkInteger(3000),
		"has": MkMap(&VarMap{
			"cancelOrder":     MkBool(true),
			"CORS":            MkBool(false),
			"createOrder":     MkBool(true),
			"fetchBalance":    MkBool(true),
			"fetchMyTrades":   MkBool(true),
			"fetchOpenOrders": MkBool(true),
			"fetchOrderBook":  MkBool(true),
			"fetchTicker":     MkBool(true),
			"fetchTrades":     MkBool(true),
		}),
		"urls": MkMap(&VarMap{
			"logo":     MkString("https://user-images.githubusercontent.com/1294454/27766119-3593220e-5ece-11e7-8b3a-5a041f6bcc3f.jpg"),
			"api":      MkString("https://bit2c.co.il"),
			"www":      MkString("https://www.bit2c.co.il"),
			"referral": MkString("https://bit2c.co.il/Aff/63bfed10-e359-420c-ab5a-ad368dab0baf"),
			"doc": MkArray(&VarArray{
				MkString("https://www.bit2c.co.il/home/api"),
				MkString("https://github.com/OferE/bit2c"),
			}),
		}),
		"api": MkMap(&VarMap{
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("Exchanges/{pair}/Ticker"),
				MkString("Exchanges/{pair}/orderbook"),
				MkString("Exchanges/{pair}/trades"),
				MkString("Exchanges/{pair}/lasttrades"),
			})}),
			"private": MkMap(&VarMap{
				"post": MkArray(&VarArray{
					MkString("Merchant/CreateCheckout"),
					MkString("Order/AddCoinFundsRequest"),
					MkString("Order/AddFund"),
					MkString("Order/AddOrder"),
					MkString("Order/AddOrderMarketPriceBuy"),
					MkString("Order/AddOrderMarketPriceSell"),
					MkString("Order/CancelOrder"),
					MkString("Order/AddCoinFundsRequest"),
					MkString("Order/AddStopOrder"),
					MkString("Payment/GetMyId"),
					MkString("Payment/Send"),
					MkString("Payment/Pay"),
				}),
				"get": MkArray(&VarArray{
					MkString("Account/Balance"),
					MkString("Account/Balance/v2"),
					MkString("Order/MyOrders"),
					MkString("Order/GetById"),
					MkString("Order/AccountHistory"),
					MkString("Order/OrderHistory"),
				}),
			}),
		}),
		"markets": MkMap(&VarMap{
			"BTC/NIS": MkMap(&VarMap{
				"id":      MkString("BtcNis"),
				"symbol":  MkString("BTC/NIS"),
				"base":    MkString("BTC"),
				"quote":   MkString("NIS"),
				"baseId":  MkString("Btc"),
				"quoteId": MkString("Nis"),
			}),
			"ETH/NIS": MkMap(&VarMap{
				"id":      MkString("EthNis"),
				"symbol":  MkString("ETH/NIS"),
				"base":    MkString("ETH"),
				"quote":   MkString("NIS"),
				"baseId":  MkString("Eth"),
				"quoteId": MkString("Nis"),
			}),
			"BCH/NIS": MkMap(&VarMap{
				"id":      MkString("BchabcNis"),
				"symbol":  MkString("BCH/NIS"),
				"base":    MkString("BCH"),
				"quote":   MkString("NIS"),
				"baseId":  MkString("Bchabc"),
				"quoteId": MkString("Nis"),
			}),
			"LTC/NIS": MkMap(&VarMap{
				"id":      MkString("LtcNis"),
				"symbol":  MkString("LTC/NIS"),
				"base":    MkString("LTC"),
				"quote":   MkString("NIS"),
				"baseId":  MkString("Ltc"),
				"quoteId": MkString("Nis"),
			}),
			"ETC/NIS": MkMap(&VarMap{
				"id":      MkString("EtcNis"),
				"symbol":  MkString("ETC/NIS"),
				"base":    MkString("ETC"),
				"quote":   MkString("NIS"),
				"baseId":  MkString("Etc"),
				"quoteId": MkString("Nis"),
			}),
			"BTG/NIS": MkMap(&VarMap{
				"id":      MkString("BtgNis"),
				"symbol":  MkString("BTG/NIS"),
				"base":    MkString("BTG"),
				"quote":   MkString("NIS"),
				"baseId":  MkString("Btg"),
				"quoteId": MkString("Nis"),
			}),
			"BSV/NIS": MkMap(&VarMap{
				"id":      MkString("BchsvNis"),
				"symbol":  MkString("BSV/NIS"),
				"base":    MkString("BSV"),
				"quote":   MkString("NIS"),
				"baseId":  MkString("Bchsv"),
				"quoteId": MkString("Nis"),
			}),
			"GRIN/NIS": MkMap(&VarMap{
				"id":      MkString("GrinNis"),
				"symbol":  MkString("GRIN/NIS"),
				"base":    MkString("GRIN"),
				"quote":   MkString("NIS"),
				"baseId":  MkString("Grin"),
				"quoteId": MkString("Nis"),
			}),
		}),
		"fees": MkMap(&VarMap{"trading": MkMap(&VarMap{
			"maker": this.ParseNumber(MkString("0.005")),
			"taker": this.ParseNumber(MkString("0.005")),
		})}),
		"options": MkMap(&VarMap{"fetchTradesMethod": MkString("public_get_exchanges_pair_trades")}),
		"exceptions": MkMap(&VarMap{
			"exact": MkMap(&VarMap{"Please provide valid APIkey": AuthenticationError}),
			"broad": MkMap(&VarMap{
				"Please provide valid nonce":              InvalidNonce,
				"please approve new terms of use on site": PermissionDenied,
			}),
		}),
	}))
}

func (this *Bit2c) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	balance := this.Call(MkString("privateGetAccountBalanceV2"), params)
	_ = balance
	result := MkMap(&VarMap{
		"info":      balance,
		"timestamp": MkUndefined(),
		"datetime":  MkUndefined(),
	})
	_ = result
	codes := GoGetKeys(*this.At(MkString("currencies")))
	_ = codes
	for i := MkInteger(0); IsTrue(OpLw(i, codes.Length)); OpIncr(&i) {
		code := *(codes).At(i)
		_ = code
		account := this.Account()
		_ = account
		currency := this.Currency(code)
		_ = currency
		uppercase := (*(currency).At(MkString("id"))).Call(MkString("toUpperCase"))
		_ = uppercase
		if IsTrue(OpHasMember(uppercase, balance)) {
			*(account).At(MkString("free")) = this.SafeString(balance, OpAdd(MkString("AVAILABLE_"), uppercase))
			*(account).At(MkString("total")) = this.SafeString(balance, uppercase)
		}
		*(result).At(code) = OpCopy(account)
	}
	return this.ParseBalance(result)
}

func (this *Bit2c) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"pair": this.MarketId(symbol)})
	_ = request
	orderbook := this.Call(MkString("publicGetExchangesPairOrderbook"), this.Extend(request, params))
	_ = orderbook
	return this.ParseOrderBook(orderbook, symbol)
}

func (this *Bit2c) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"pair": this.MarketId(symbol)})
	_ = request
	ticker := this.Call(MkString("publicGetExchangesPairTicker"), this.Extend(request, params))
	_ = ticker
	timestamp := this.Milliseconds()
	_ = timestamp
	averagePrice := this.SafeNumber(ticker, MkString("av"))
	_ = averagePrice
	baseVolume := this.SafeNumber(ticker, MkString("a"))
	_ = baseVolume
	quoteVolume := OpCopy(MkUndefined())
	_ = quoteVolume
	if IsTrue(OpAnd(OpNotEq2(baseVolume, MkUndefined()), OpNotEq2(averagePrice, MkUndefined()))) {
		quoteVolume = OpMulti(baseVolume, averagePrice)
	}
	last := this.SafeNumber(ticker, MkString("ll"))
	_ = last
	return MkMap(&VarMap{
		"symbol":        symbol,
		"timestamp":     timestamp,
		"datetime":      this.Iso8601(timestamp),
		"high":          MkUndefined(),
		"low":           MkUndefined(),
		"bid":           this.SafeNumber(ticker, MkString("h")),
		"bidVolume":     MkUndefined(),
		"ask":           this.SafeNumber(ticker, MkString("l")),
		"askVolume":     MkUndefined(),
		"vwap":          MkUndefined(),
		"open":          MkUndefined(),
		"close":         last,
		"last":          last,
		"previousClose": MkUndefined(),
		"change":        MkUndefined(),
		"percentage":    MkUndefined(),
		"average":       averagePrice,
		"baseVolume":    baseVolume,
		"quoteVolume":   quoteVolume,
		"info":          ticker,
	})
}

func (this *Bit2c) FetchTrades(goArgs ...*Variant) *Variant {
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
	method := *(*this.At(MkString("options"))).At(MkString("fetchTradesMethod"))
	_ = method
	request := MkMap(&VarMap{"pair": *(market).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("date")) = parseInt(since)
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	if IsTrue(OpEq2(OpType(response), MkString("string"))) {
		panic(NewExchangeError(response))
	}
	return this.ParseTrades(response, market, since, limit)
}

func (this *Bit2c) CreateOrder(goArgs ...*Variant) *Variant {
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
	method := MkString("privatePostOrderAddOrder")
	_ = method
	request := MkMap(&VarMap{
		"Amount": amount,
		"Pair":   this.MarketId(symbol),
	})
	_ = request
	if IsTrue(OpEq2(type_, MkString("market"))) {
		OpAddAssign(&method, OpAdd(MkString("MarketPrice"), this.Capitalize(side)))
	} else {
		*(request).At(MkString("Price")) = OpCopy(price)
		*(request).At(MkString("Total")) = OpMulti(amount, price)
		*(request).At(MkString("IsBid")) = OpEq2(side, MkString("buy"))
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	return MkMap(&VarMap{
		"info": response,
		"id":   *(*(response).At(MkString("NewOrder"))).At(MkString("id")),
	})
}

func (this *Bit2c) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	request := MkMap(&VarMap{"id": id})
	_ = request
	return this.Call(MkString("privatePostOrderCancelOrder"), this.Extend(request, params))
}

func (this *Bit2c) FetchOpenOrders(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"pair": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("privateGetOrderMyOrders"), this.Extend(request, params))
	_ = response
	orders := this.SafeValue(response, *(market).At(MkString("id")), MkMap(&VarMap{}))
	_ = orders
	asks := this.SafeValue(orders, MkString("ask"), MkArray(&VarArray{}))
	_ = asks
	bids := this.SafeValue(orders, MkString("bid"), MkArray(&VarArray{}))
	_ = bids
	return this.ParseOrders(this.ArrayConcat(asks, bids), market, since, limit)
}

func (this *Bit2c) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.SafeInteger(order, MkString("created"))
	_ = timestamp
	price := this.SafeNumber(order, MkString("price"))
	_ = price
	amount := this.SafeNumber(order, MkString("amount"))
	_ = amount
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(OpNotEq2(market, MkUndefined())) {
		symbol = *(market).At(MkString("symbol"))
	}
	side := this.SafeValue(order, MkString("type"))
	_ = side
	if IsTrue(OpEq2(side, MkInteger(0))) {
		side = MkString("buy")
	} else {
		if IsTrue(OpEq2(side, MkInteger(1))) {
			side = MkString("sell")
		}
	}
	id := this.SafeString(order, MkString("id"))
	_ = id
	status := this.SafeString(order, MkString("status"))
	_ = status
	return this.SafeOrder(MkMap(&VarMap{
		"id":                 id,
		"clientOrderId":      MkUndefined(),
		"timestamp":          timestamp,
		"datetime":           this.Iso8601(timestamp),
		"lastTradeTimestamp": MkUndefined(),
		"status":             status,
		"symbol":             symbol,
		"type":               MkUndefined(),
		"timeInForce":        MkUndefined(),
		"postOnly":           MkUndefined(),
		"side":               side,
		"price":              price,
		"stopPrice":          MkUndefined(),
		"amount":             amount,
		"filled":             MkUndefined(),
		"remaining":          MkUndefined(),
		"cost":               MkUndefined(),
		"trades":             MkUndefined(),
		"fee":                MkUndefined(),
		"info":               order,
		"average":            MkUndefined(),
	}))
}

func (this *Bit2c) FetchMyTrades(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("take")) = OpCopy(limit)
	}
	*(request).At(MkString("take")) = OpCopy(limit)
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("toTime")) = this.Ymd(this.Milliseconds(), MkString("."))
		*(request).At(MkString("fromTime")) = this.Ymd(since, MkString("."))
	}
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
		*(request).At(MkString("pair")) = *(market).At(MkString("id"))
	}
	response := this.Call(MkString("privateGetOrderOrderHistory"), this.Extend(request, params))
	_ = response
	return this.ParseTrades(response, market, since, limit)
}

func (this *Bit2c) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := OpCopy(MkUndefined())
	_ = timestamp
	id := OpCopy(MkUndefined())
	_ = id
	priceString := OpCopy(MkUndefined())
	_ = priceString
	amountString := OpCopy(MkUndefined())
	_ = amountString
	orderId := OpCopy(MkUndefined())
	_ = orderId
	feeCost := OpCopy(MkUndefined())
	_ = feeCost
	side := OpCopy(MkUndefined())
	_ = side
	reference := this.SafeString(trade, MkString("reference"))
	_ = reference
	if IsTrue(OpNotEq2(reference, MkUndefined())) {
		timestamp = this.SafeTimestamp(trade, MkString("ticks"))
		priceString = this.SafeString(trade, MkString("price"))
		amountString = this.SafeString(trade, MkString("firstAmount"))
		reference_parts := reference.Split(MkString("|"))
		_ = reference_parts
		if IsTrue(OpEq2(market, MkUndefined())) {
			marketId := this.SafeString(trade, MkString("pair"))
			_ = marketId
			if IsTrue(OpHasMember(marketId, *(*this.At(MkString("markets_by_id"))).At(marketId))) {
				market = *(*this.At(MkString("markets_by_id"))).At(marketId)
			} else {
				if IsTrue(OpHasMember(*(reference_parts).At(MkInteger(0)), *this.At(MkString("markets_by_id")))) {
					market = *(*this.At(MkString("markets_by_id"))).At(*(reference_parts).At(MkInteger(0)))
				}
			}
		}
		orderId = *(reference_parts).At(MkInteger(1))
		id = *(reference_parts).At(MkInteger(2))
		side = this.SafeInteger(trade, MkString("action"))
		if IsTrue(OpEq2(side, MkInteger(0))) {
			side = MkString("buy")
		} else {
			if IsTrue(OpEq2(side, MkInteger(1))) {
				side = MkString("sell")
			}
		}
		feeCost = this.SafeNumber(trade, MkString("feeAmount"))
	} else {
		timestamp = this.SafeTimestamp(trade, MkString("date"))
		id = this.SafeString(trade, MkString("tid"))
		priceString = this.SafeString(trade, MkString("price"))
		amountString = this.SafeString(trade, MkString("amount"))
		side = this.SafeValue(trade, MkString("isBid"))
		if IsTrue(OpNotEq2(side, MkUndefined())) {
			if IsTrue(side) {
				side = MkString("buy")
			} else {
				side = MkString("sell")
			}
		}
	}
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(OpNotEq2(market, MkUndefined())) {
		symbol = *(market).At(MkString("symbol"))
	}
	price := this.ParseNumber(priceString)
	_ = price
	amount := this.ParseNumber(amountString)
	_ = amount
	cost := this.ParseNumber(Precise.StringMul(priceString, amountString))
	_ = cost
	return MkMap(&VarMap{
		"info":         trade,
		"id":           id,
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       symbol,
		"order":        orderId,
		"type":         MkUndefined(),
		"side":         side,
		"takerOrMaker": MkUndefined(),
		"price":        price,
		"amount":       amount,
		"cost":         cost,
		"fee": MkMap(&VarMap{
			"cost":     feeCost,
			"currency": MkString("NIS"),
			"rate":     MkUndefined(),
		}),
	})
}

func (this *Bit2c) Sign(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpEq2(api, MkString("public"))) {
		OpAddAssign(&url, MkString(".json"))
	} else {
		this.CheckRequiredCredentials()
		nonce := this.Nonce()
		_ = nonce
		query := this.Extend(MkMap(&VarMap{"nonce": nonce}), params)
		_ = query
		auth := this.Urlencode(query)
		_ = auth
		if IsTrue(OpEq2(method, MkString("GET"))) {
			if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
				OpAddAssign(&url, OpAdd(MkString("?"), auth))
			}
		} else {
			body = OpCopy(auth)
		}
		signature := this.Hmac(this.Encode(auth), this.Encode(*this.At(MkString("secret"))), MkString("sha512"), MkString("base64"))
		_ = signature
		headers = MkMap(&VarMap{
			"Content-Type": MkString("application/x-www-form-urlencoded"),
			"key":          *this.At(MkString("apiKey")),
			"sign":         signature,
		})
	}
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Bit2c) HandleErrors(goArgs ...*Variant) *Variant {
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
	error := this.SafeString(response, MkString("error"))
	_ = error
	if IsTrue(OpNotEq2(error, MkUndefined())) {
		feedback := OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))
		_ = feedback
		this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("exact")), error, feedback)
		this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("broad")), error, feedback)
		panic(NewExchangeError(feedback))
	}
	return MkUndefined()
}
