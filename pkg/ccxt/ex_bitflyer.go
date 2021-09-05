package ccxt

import ()

type Bitflyer struct {
	*ExchangeBase
}

var _ Exchange = (*Bitflyer)(nil)

func init() {
	exchange := &Bitflyer{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Bitflyer) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("bitflyer"),
		"name": MkString("bitFlyer"),
		"countries": MkArray(&VarArray{
			MkString("JP"),
		}),
		"version":   MkString("v1"),
		"rateLimit": MkInteger(1000),
		"hostname":  MkString("bitflyer.com"),
		"has": MkMap(&VarMap{
			"cancelOrder":       MkBool(true),
			"CORS":              MkBool(false),
			"createOrder":       MkBool(true),
			"fetchBalance":      MkBool(true),
			"fetchClosedOrders": MkString("emulated"),
			"fetchMarkets":      MkBool(true),
			"fetchMyTrades":     MkBool(true),
			"fetchOpenOrders":   MkString("emulated"),
			"fetchOrder":        MkString("emulated"),
			"fetchOrderBook":    MkBool(true),
			"fetchOrders":       MkBool(true),
			"fetchTicker":       MkBool(true),
			"fetchTrades":       MkBool(true),
			"withdraw":          MkBool(true),
		}),
		"urls": MkMap(&VarMap{
			"logo": MkString("https://user-images.githubusercontent.com/1294454/28051642-56154182-660e-11e7-9b0d-6042d1e6edd8.jpg"),
			"api":  MkString("https://api.{hostname}"),
			"www":  MkString("https://bitflyer.com"),
			"doc":  MkString("https://lightning.bitflyer.com/docs?lang=en"),
		}),
		"api": MkMap(&VarMap{
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("getmarkets/usa"),
				MkString("getmarkets/eu"),
				MkString("getmarkets"),
				MkString("getboard"),
				MkString("getticker"),
				MkString("getexecutions"),
				MkString("gethealth"),
				MkString("getboardstate"),
				MkString("getchats"),
			})}),
			"private": MkMap(&VarMap{
				"get": MkArray(&VarArray{
					MkString("getpermissions"),
					MkString("getbalance"),
					MkString("getbalancehistory"),
					MkString("getcollateral"),
					MkString("getcollateralhistory"),
					MkString("getcollateralaccounts"),
					MkString("getaddresses"),
					MkString("getcoinins"),
					MkString("getcoinouts"),
					MkString("getbankaccounts"),
					MkString("getdeposits"),
					MkString("getwithdrawals"),
					MkString("getchildorders"),
					MkString("getparentorders"),
					MkString("getparentorder"),
					MkString("getexecutions"),
					MkString("getpositions"),
					MkString("gettradingcommission"),
				}),
				"post": MkArray(&VarArray{
					MkString("sendcoin"),
					MkString("withdraw"),
					MkString("sendchildorder"),
					MkString("cancelchildorder"),
					MkString("sendparentorder"),
					MkString("cancelparentorder"),
					MkString("cancelallchildorders"),
				}),
			}),
		}),
		"fees": MkMap(&VarMap{"trading": MkMap(&VarMap{
			"maker": this.ParseNumber(MkString("0.002")),
			"taker": this.ParseNumber(MkString("0.002")),
		})}),
	}))
}

func (this *Bitflyer) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	jp_markets := this.Call(MkString("publicGetGetmarkets"), params)
	_ = jp_markets
	us_markets := this.Call(MkString("publicGetGetmarketsUsa"), params)
	_ = us_markets
	eu_markets := this.Call(MkString("publicGetGetmarketsEu"), params)
	_ = eu_markets
	markets := this.ArrayConcat(jp_markets, us_markets)
	_ = markets
	markets = this.ArrayConcat(markets, eu_markets)
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, markets.Length)); OpIncr(&i) {
		market := *(markets).At(i)
		_ = market
		id := this.SafeString(market, MkString("product_code"))
		_ = id
		currencies := id.Split(MkString("_"))
		_ = currencies
		baseId := OpCopy(MkUndefined())
		_ = baseId
		quoteId := OpCopy(MkUndefined())
		_ = quoteId
		base := OpCopy(MkUndefined())
		_ = base
		quote := OpCopy(MkUndefined())
		_ = quote
		numCurrencies := OpCopy(currencies.Length)
		_ = numCurrencies
		if IsTrue(OpEq2(numCurrencies, MkInteger(1))) {
			baseId = id.Slice(MkInteger(0), MkInteger(3))
			quoteId = id.Slice(MkInteger(3), MkInteger(6))
		} else {
			if IsTrue(OpEq2(numCurrencies, MkInteger(2))) {
				baseId = *(currencies).At(MkInteger(0))
				quoteId = *(currencies).At(MkInteger(1))
			} else {
				baseId = *(currencies).At(MkInteger(1))
				quoteId = *(currencies).At(MkInteger(2))
			}
		}
		base = this.SafeCurrencyCode(baseId)
		quote = this.SafeCurrencyCode(quoteId)
		symbol := OpTrinary(OpEq2(numCurrencies, MkInteger(2)), OpAdd(base, OpAdd(MkString("/"), quote)), id)
		_ = symbol
		fees := this.SafeValue(*this.At(MkString("fees")), symbol, *(*this.At(MkString("fees"))).At(MkString("trading")))
		_ = fees
		maker := this.SafeValue(fees, MkString("maker"), *(*(*this.At(MkString("fees"))).At(MkString("trading"))).At(MkString("maker")))
		_ = maker
		taker := this.SafeValue(fees, MkString("taker"), *(*(*this.At(MkString("fees"))).At(MkString("trading"))).At(MkString("taker")))
		_ = taker
		spot := OpCopy(MkBool(true))
		_ = spot
		future := OpCopy(MkBool(false))
		_ = future
		type_ := MkString("spot")
		_ = type_
		if IsTrue(OpOr(OpHasMember(MkString("alias"), market), OpEq2(*(currencies).At(MkInteger(0)), MkString("FX")))) {
			type_ = MkString("future")
			future = OpCopy(MkBool(true))
			spot = OpCopy(MkBool(false))
			maker = MkNumber(0.0)
			taker = MkNumber(0.0)
		}
		result.Push(MkMap(&VarMap{
			"id":      id,
			"symbol":  symbol,
			"base":    base,
			"quote":   quote,
			"baseId":  baseId,
			"quoteId": quoteId,
			"maker":   maker,
			"taker":   taker,
			"type":    type_,
			"spot":    spot,
			"future":  future,
			"info":    market,
		}))
	}
	return result
}

func (this *Bitflyer) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privateGetGetbalance"), params)
	_ = response
	result := MkMap(&VarMap{"info": response})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
		balance := *(response).At(i)
		_ = balance
		currencyId := this.SafeString(balance, MkString("currency_code"))
		_ = currencyId
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		account := this.Account()
		_ = account
		*(account).At(MkString("total")) = this.SafeString(balance, MkString("amount"))
		*(account).At(MkString("free")) = this.SafeString(balance, MkString("available"))
		*(result).At(code) = OpCopy(account)
	}
	return this.ParseBalance(result)
}

func (this *Bitflyer) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"product_code": this.MarketId(symbol)})
	_ = request
	orderbook := this.Call(MkString("publicGetGetboard"), this.Extend(request, params))
	_ = orderbook
	return this.ParseOrderBook(orderbook, symbol, MkUndefined(), MkString("bids"), MkString("asks"), MkString("price"), MkString("size"))
}

func (this *Bitflyer) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"product_code": this.MarketId(symbol)})
	_ = request
	ticker := this.Call(MkString("publicGetGetticker"), this.Extend(request, params))
	_ = ticker
	timestamp := this.Parse8601(this.SafeString(ticker, MkString("timestamp")))
	_ = timestamp
	last := this.SafeNumber(ticker, MkString("ltp"))
	_ = last
	return MkMap(&VarMap{
		"symbol":        symbol,
		"timestamp":     timestamp,
		"datetime":      this.Iso8601(timestamp),
		"high":          MkUndefined(),
		"low":           MkUndefined(),
		"bid":           this.SafeNumber(ticker, MkString("best_bid")),
		"bidVolume":     MkUndefined(),
		"ask":           this.SafeNumber(ticker, MkString("best_ask")),
		"askVolume":     MkUndefined(),
		"vwap":          MkUndefined(),
		"open":          MkUndefined(),
		"close":         last,
		"last":          last,
		"previousClose": MkUndefined(),
		"change":        MkUndefined(),
		"percentage":    MkUndefined(),
		"average":       MkUndefined(),
		"baseVolume":    this.SafeNumber(ticker, MkString("volume_by_product")),
		"quoteVolume":   MkUndefined(),
		"info":          ticker,
	})
}

func (this *Bitflyer) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	side := this.SafeStringLower(trade, MkString("side"))
	_ = side
	if IsTrue(OpNotEq2(side, MkUndefined())) {
		if IsTrue(OpLw(side.Length, MkInteger(1))) {
			side = OpCopy(MkUndefined())
		}
	}
	order := OpCopy(MkUndefined())
	_ = order
	if IsTrue(OpNotEq2(side, MkUndefined())) {
		id := OpAdd(side, MkString("_child_order_acceptance_id"))
		_ = id
		if IsTrue(OpHasMember(id, trade)) {
			order = *(trade).At(id)
		}
	}
	if IsTrue(OpEq2(order, MkUndefined())) {
		order = this.SafeString(trade, MkString("child_order_acceptance_id"))
	}
	timestamp := this.Parse8601(this.SafeString(trade, MkString("exec_date")))
	_ = timestamp
	priceString := this.SafeString(trade, MkString("price"))
	_ = priceString
	amountString := this.SafeString(trade, MkString("size"))
	_ = amountString
	price := this.ParseNumber(priceString)
	_ = price
	amount := this.ParseNumber(amountString)
	_ = amount
	cost := this.ParseNumber(Precise.StringMul(priceString, amountString))
	_ = cost
	id := this.SafeString(trade, MkString("id"))
	_ = id
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(OpNotEq2(market, MkUndefined())) {
		symbol = *(market).At(MkString("symbol"))
	}
	return MkMap(&VarMap{
		"id":           id,
		"info":         trade,
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       symbol,
		"order":        order,
		"type":         MkUndefined(),
		"side":         side,
		"takerOrMaker": MkUndefined(),
		"price":        price,
		"amount":       amount,
		"cost":         cost,
		"fee":          MkUndefined(),
	})
}

func (this *Bitflyer) FetchTrades(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"product_code": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("publicGetGetexecutions"), this.Extend(request, params))
	_ = response
	return this.ParseTrades(response, market, since, limit)
}

func (this *Bitflyer) CreateOrder(goArgs ...*Variant) *Variant {
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
		"product_code":     this.MarketId(symbol),
		"child_order_type": type_.ToUpperCase(),
		"side":             side.ToUpperCase(),
		"price":            price,
		"size":             amount,
	})
	_ = request
	result := this.Call(MkString("privatePostSendchildorder"), this.Extend(request, params))
	_ = result
	id := this.SafeString(result, MkString("child_order_acceptance_id"))
	_ = id
	return MkMap(&VarMap{
		"info": result,
		"id":   id,
	})
}

func (this *Bitflyer) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" cancelOrder() requires a `symbol` argument"))))
	}
	this.LoadMarkets()
	request := MkMap(&VarMap{
		"product_code":              this.MarketId(symbol),
		"child_order_acceptance_id": id,
	})
	_ = request
	return this.Call(MkString("privatePostCancelchildorder"), this.Extend(request, params))
}

func (this *Bitflyer) ParseOrderStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"ACTIVE":    MkString("open"),
		"COMPLETED": MkString("closed"),
		"CANCELED":  MkString("canceled"),
		"EXPIRED":   MkString("canceled"),
		"REJECTED":  MkString("canceled"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Bitflyer) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.Parse8601(this.SafeString(order, MkString("child_order_date")))
	_ = timestamp
	amount := this.SafeNumber(order, MkString("size"))
	_ = amount
	remaining := this.SafeNumber(order, MkString("outstanding_size"))
	_ = remaining
	filled := this.SafeNumber(order, MkString("executed_size"))
	_ = filled
	price := this.SafeNumber(order, MkString("price"))
	_ = price
	status := this.ParseOrderStatus(this.SafeString(order, MkString("child_order_state")))
	_ = status
	type_ := this.SafeStringLower(order, MkString("child_order_type"))
	_ = type_
	side := this.SafeStringLower(order, MkString("side"))
	_ = side
	marketId := this.SafeString(order, MkString("product_code"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market)
	_ = symbol
	fee := OpCopy(MkUndefined())
	_ = fee
	feeCost := this.SafeNumber(order, MkString("total_commission"))
	_ = feeCost
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		fee = MkMap(&VarMap{
			"cost":     feeCost,
			"currency": MkUndefined(),
			"rate":     MkUndefined(),
		})
	}
	id := this.SafeString(order, MkString("child_order_acceptance_id"))
	_ = id
	return this.SafeOrder(MkMap(&VarMap{
		"id":                 id,
		"clientOrderId":      MkUndefined(),
		"info":               order,
		"timestamp":          timestamp,
		"datetime":           this.Iso8601(timestamp),
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
		"remaining":          remaining,
		"fee":                fee,
		"average":            MkUndefined(),
		"trades":             MkUndefined(),
	}))
}

func (this *Bitflyer) FetchOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkInteger(100))
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchOrders() requires a `symbol` argument"))))
	}
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{
		"product_code": *(market).At(MkString("id")),
		"count":        limit,
	})
	_ = request
	response := this.Call(MkString("privateGetGetchildorders"), this.Extend(request, params))
	_ = response
	orders := this.ParseOrders(response, market, since, limit)
	_ = orders
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		orders = this.FilterBy(orders, MkString("symbol"), symbol)
	}
	return orders
}

func (this *Bitflyer) FetchOpenOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkInteger(100))
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	request := MkMap(&VarMap{"child_order_state": MkString("ACTIVE")})
	_ = request
	return this.FetchOrders(symbol, since, limit, this.Extend(request, params))
}

func (this *Bitflyer) FetchClosedOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkInteger(100))
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	request := MkMap(&VarMap{"child_order_state": MkString("COMPLETED")})
	_ = request
	return this.FetchOrders(symbol, since, limit, this.Extend(request, params))
}

func (this *Bitflyer) FetchOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchOrder() requires a `symbol` argument"))))
	}
	orders := this.FetchOrders(symbol)
	_ = orders
	ordersById := this.IndexBy(orders, MkString("id"))
	_ = ordersById
	if IsTrue(OpHasMember(id, ordersById)) {
		return *(ordersById).At(id)
	}
	panic(NewOrderNotFound(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" No order found with id "), id))))
	return MkUndefined()
}

func (this *Bitflyer) FetchMyTrades(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchMyTrades() requires a `symbol` argument"))))
	}
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"product_code": *(market).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("count")) = OpCopy(limit)
	}
	response := this.Call(MkString("privateGetGetexecutions"), this.Extend(request, params))
	_ = response
	return this.ParseTrades(response, market, since, limit)
}

func (this *Bitflyer) FetchPositions(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(symbols, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchPositions() requires a `symbols` argument, exactly one symbol in an array"))))
	}
	this.LoadMarkets()
	request := MkMap(&VarMap{"product_code": this.MarketIds(symbols)})
	_ = request
	response := this.Call(MkString("privateGetpositions"), this.Extend(request, params))
	_ = response
	return response
}

func (this *Bitflyer) Withdraw(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpAnd(OpNotEq2(code, MkString("JPY")), OpAnd(OpNotEq2(code, MkString("USD")), OpNotEq2(code, MkString("EUR"))))) {
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" allows withdrawing JPY, USD, EUR only, "), OpAdd(code, MkString(" is not supported"))))))
	}
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{
		"currency_code": *(currency).At(MkString("id")),
		"amount":        amount,
	})
	_ = request
	response := this.Call(MkString("privatePostWithdraw"), this.Extend(request, params))
	_ = response
	id := this.SafeString(response, MkString("message_id"))
	_ = id
	return MkMap(&VarMap{
		"info": response,
		"id":   id,
	})
}

func (this *Bitflyer) Sign(goArgs ...*Variant) *Variant {
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
	request := OpAdd(MkString("/"), OpAdd(*this.At(MkString("version")), MkString("/")))
	_ = request
	if IsTrue(OpEq2(api, MkString("private"))) {
		OpAddAssign(&request, MkString("me/"))
	}
	OpAddAssign(&request, path)
	if IsTrue(OpEq2(method, MkString("GET"))) {
		if IsTrue(*(GoGetKeys(params)).At(MkString("length"))) {
			OpAddAssign(&request, OpAdd(MkString("?"), this.Urlencode(params)))
		}
	}
	baseUrl := this.ImplodeHostname(*(*this.At(MkString("urls"))).At(MkString("api")))
	_ = baseUrl
	url := OpAdd(baseUrl, request)
	_ = url
	if IsTrue(OpEq2(api, MkString("private"))) {
		this.CheckRequiredCredentials()
		nonce := (this.Nonce()).Call(MkString("toString"))
		_ = nonce
		auth := (MkArray(&VarArray{
			nonce,
			method,
			request,
		})).Call(MkString("join"), MkString(""))
		_ = auth
		if IsTrue(*(GoGetKeys(params)).At(MkString("length"))) {
			if IsTrue(OpNotEq2(method, MkString("GET"))) {
				body = this.Json(params)
				OpAddAssign(&auth, body)
			}
		}
		headers = MkMap(&VarMap{
			"ACCESS-KEY":       *this.At(MkString("apiKey")),
			"ACCESS-TIMESTAMP": nonce,
			"ACCESS-SIGN":      this.Hmac(this.Encode(auth), this.Encode(*this.At(MkString("secret")))),
			"Content-Type":     MkString("application/json"),
		})
	}
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}
