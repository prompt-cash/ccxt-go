package ccxt

import ()

type Hbtc struct {
	*ExchangeBase
}

var _ Exchange = (*Hbtc)(nil)

func init() {
	exchange := &Hbtc{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Hbtc) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("hbtc"),
		"name": MkString("HBTC"),
		"countries": MkArray(&VarArray{
			MkString("CN"),
		}),
		"rateLimit": MkInteger(2000),
		"version":   MkString("v1"),
		"has": MkMap(&VarMap{
			"cancelOrder":         MkBool(true),
			"CORS":                MkBool(false),
			"createOrder":         MkBool(true),
			"fetchAccounts":       MkBool(true),
			"fetchBalance":        MkBool(true),
			"fetchBidAsk":         MkBool(true),
			"fetchBidsAsks":       MkBool(true),
			"fetchClosedOrders":   MkBool(true),
			"fetchCurrencies":     MkBool(false),
			"fetchDepositAddress": MkBool(false),
			"fetchDeposits":       MkBool(true),
			"fetchLedger":         MkBool(true),
			"fetchMarkets":        MkBool(true),
			"fetchMyTrades":       MkBool(true),
			"fetchOHLCV":          MkBool(true),
			"fetchOpenOrders":     MkBool(true),
			"fetchOrder":          MkBool(true),
			"fetchOrderBook":      MkBool(true),
			"fetchOrders":         MkBool(false),
			"fetchTicker":         MkBool(true),
			"fetchTickers":        MkBool(true),
			"fetchTime":           MkBool(true),
			"fetchTrades":         MkBool(true),
			"fetchTradingLimits":  MkBool(true),
			"fetchWithdrawals":    MkBool(true),
			"withdraw":            MkBool(true),
		}),
		"timeframes": MkMap(&VarMap{
			"1m":  MkString("1m"),
			"3m":  MkString("3m"),
			"5m":  MkString("5m"),
			"15m": MkString("15m"),
			"30m": MkString("30m"),
			"1h":  MkString("1h"),
			"2h":  MkString("2h"),
			"4h":  MkString("4h"),
			"6h":  MkString("6h"),
			"8h":  MkString("8h"),
			"12h": MkString("12h"),
			"1d":  MkString("1d"),
			"3d":  MkString("3d"),
			"1w":  MkString("1w"),
			"1M":  MkString("1M"),
		}),
		"urls": MkMap(&VarMap{
			"logo": MkString("https://user-images.githubusercontent.com/51840849/80134449-70663300-85a7-11ea-8942-e204cdeaab5d.jpg"),
			"api": MkMap(&VarMap{
				"quote":    MkString("https://api.hbtc.com/openapi/quote"),
				"contract": MkString("https://api.hbtc.com/openapi/contract"),
				"option":   MkString("https://api.hbtc.com/openapi/option"),
				"public":   MkString("https://api.hbtc.com/openapi"),
				"private":  MkString("https://api.hbtc.com/openapi"),
				"zendesk":  MkString("https://hbtc.zendesk.com/hc/en-us"),
			}),
			"www":      MkString("https://www.hbtc.com"),
			"referral": MkString("https://www.hbtc.com/register/O2S8NS"),
			"doc":      MkString("https://github.com/bhexopen/BHEX-OpenApi/tree/master/doc"),
			"fees":     MkString("https://hbtc.zendesk.com/hc/zh-cn/articles/360009274694"),
		}),
		"api": MkMap(&VarMap{
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("ping"),
				MkString("time"),
				MkString("brokerInfo"),
				MkString("getOptions"),
			})}),
			"quote": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("depth"),
				MkString("depth/merged"),
				MkString("trades"),
				MkString("klines"),
				MkString("ticker/24hr"),
				MkString("ticker/price"),
				MkString("ticker/bookTicker"),
				MkString("contract/index"),
				MkString("contract/depth"),
				MkString("contract/depth/merged"),
				MkString("contract/trades"),
				MkString("contract/klines"),
				MkString("contract/ticker/24hr"),
				MkString("option/index"),
				MkString("option/depth"),
				MkString("option/depth/merged"),
				MkString("option/trades"),
				MkString("option/klines"),
				MkString("option/ticker/24hr"),
			})}),
			"contract": MkMap(&VarMap{
				"get": MkArray(&VarArray{
					MkString("insurance"),
					MkString("fundingRate"),
					MkString("openOrders"),
					MkString("historyOrders"),
					MkString("getOrder"),
					MkString("myTrades"),
					MkString("positions"),
					MkString("account"),
				}),
				"post": MkArray(&VarArray{
					MkString("order"),
					MkString("modifyMargin"),
				}),
				"delete": MkArray(&VarArray{
					MkString("order/cancel"),
					MkString("order/batchCancel"),
				}),
			}),
			"option": MkMap(&VarMap{
				"get": MkArray(&VarArray{
					MkString("openOrders"),
					MkString("positions"),
					MkString("historyOrders"),
					MkString("myTrades"),
					MkString("settlements"),
					MkString("account"),
				}),
				"post": MkArray(&VarArray{
					MkString("order"),
				}),
				"delete": MkArray(&VarArray{
					MkString("order/cancel"),
				}),
			}),
			"private": MkMap(&VarMap{
				"get": MkArray(&VarArray{
					MkString("order"),
					MkString("openOrders"),
					MkString("historyOrders"),
					MkString("account"),
					MkString("myTrades"),
					MkString("depositOrders"),
					MkString("withdrawalOrders"),
					MkString("withdraw/detail"),
					MkString("balance_flow"),
				}),
				"post": MkArray(&VarArray{
					MkString("order"),
					MkString("order/test"),
					MkString("userDataStream"),
					MkString("subAccount/query"),
					MkString("transfer"),
					MkString("user/transfer"),
					MkString("withdraw"),
				}),
				"put": MkArray(&VarArray{
					MkString("userDataStream"),
				}),
				"delete": MkArray(&VarArray{
					MkString("order"),
					MkString("userDataStream"),
				}),
			}),
		}),
		"precisionMode": TICK_SIZE,
		"fees": MkMap(&VarMap{"trading": MkMap(&VarMap{
			"tierBased":  MkBool(false),
			"percentage": MkBool(true),
			"maker":      MkNumber(0.001),
			"taker":      MkNumber(0.001),
		})}),
		"exceptions": MkMap(&VarMap{"exact": MkMap(&VarMap{
			"-1000": ExchangeError,
			"-1001": ExchangeError,
			"-1002": AuthenticationError,
			"-1003": RateLimitExceeded,
			"-1004": BadRequest,
			"-1005": PermissionDenied,
			"-1006": BadResponse,
			"-1007": RequestTimeout,
			"-1014": InvalidOrder,
			"-1015": RateLimitExceeded,
			"-1016": ExchangeNotAvailable,
			"-1020": NotSupported,
			"-1021": BadRequest,
			"-1022": AuthenticationError,
			"-1100": BadRequest,
			"-1101": BadRequest,
			"-1102": BadRequest,
			"-1103": BadRequest,
			"-1104": BadRequest,
			"-1105": BadRequest,
			"-1106": BadRequest,
			"-1111": BadRequest,
			"-1112": NullResponse,
			"-1114": InvalidOrder,
			"-1115": InvalidOrder,
			"-1116": InvalidOrder,
			"-1117": InvalidOrder,
			"-1118": InvalidOrder,
			"-1119": InvalidOrder,
			"-1120": BadRequest,
			"-1121": BadSymbol,
			"-1125": AuthenticationError,
			"-1127": BadRequest,
			"-1128": BadRequest,
			"-1130": BadRequest,
			"-1131": InsufficientFunds,
			"-1132": InvalidOrder,
			"-1133": InvalidOrder,
			"-1134": InvalidOrder,
			"-1135": InvalidOrder,
			"-1136": InvalidOrder,
			"-1137": InvalidOrder,
			"-1138": InvalidOrder,
			"-1139": InvalidOrder,
			"-1140": InvalidOrder,
			"-1141": InvalidOrder,
			"-1142": InvalidOrder,
			"-1143": OrderNotFound,
			"-1144": InvalidOrder,
			"-1145": InvalidOrder,
			"-1146": RequestTimeout,
			"-1147": RequestTimeout,
			"-1149": InvalidOrder,
			"-1187": InvalidAddress,
			"-2010": InvalidOrder,
			"-2011": InvalidOrder,
			"-2013": OrderNotFound,
			"-2014": AuthenticationError,
			"-2015": AuthenticationError,
			"-2016": ExchangeError,
		})}),
		"options": MkMap(&VarMap{
			"fetchTickers": MkMap(&VarMap{"method": MkString("quoteGetTicker24hr")}),
			"accountsByType": MkMap(&VarMap{
				"trade":    MkInteger(1),
				"trading":  MkInteger(1),
				"spot":     MkInteger(1),
				"option":   MkInteger(2),
				"options":  MkInteger(2),
				"futures":  MkInteger(3),
				"contract": MkInteger(3),
			}),
		}),
		"commonCurrencies": MkMap(&VarMap{"MIS": MkString("Themis Protocol")}),
	}))
}

func (this *Hbtc) FetchTime(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetTime"), params)
	_ = response
	return this.SafeInteger(response, MkString("serverTime"))
}

func (this *Hbtc) ParseMarket(goArgs ...*Variant) *Variant {
	market := GoGetArg(goArgs, 0, MkUndefined())
	_ = market
	type_ := GoGetArg(goArgs, 1, MkString("spot"))
	_ = type_
	filters := this.SafeValue(market, MkString("filters"), MkArray(&VarArray{}))
	_ = filters
	id := this.SafeString(market, MkString("symbol"))
	_ = id
	baseId := this.SafeString(market, MkString("baseAsset"))
	_ = baseId
	quoteId := this.SafeString(market, MkString("quoteAsset"))
	_ = quoteId
	base := this.SafeCurrencyCode(baseId)
	_ = base
	quote := this.SafeCurrencyCode(quoteId)
	_ = quote
	symbol := OpAdd(base, OpAdd(MkString("/"), quote))
	_ = symbol
	spot := OpCopy(MkBool(true))
	_ = spot
	future := OpCopy(MkBool(false))
	_ = future
	option := OpCopy(MkBool(false))
	_ = option
	inverse := OpCopy(MkBool(false))
	_ = inverse
	if IsTrue(OpEq2(type_, MkString("future"))) {
		symbol = OpCopy(id)
		spot = OpCopy(MkBool(false))
		future = OpCopy(MkBool(true))
		inverse = this.SafeValue(market, MkString("inverse"), MkBool(false))
		baseId = this.SafeString(market, MkString("underlying"))
		base = this.SafeCurrencyCode(baseId)
	} else {
		if IsTrue(OpEq2(type_, MkString("option"))) {
			symbol = OpCopy(id)
			spot = OpCopy(MkBool(false))
			option = OpCopy(MkBool(true))
		}
	}
	margin := this.SafeValue(market, MkString("allowMargin"), MkUndefined())
	_ = margin
	isAggregate := this.SafeValue(market, MkString("isAggregate"), MkUndefined())
	_ = isAggregate
	active := OpCopy(MkBool(true))
	_ = active
	if IsTrue(OpEq2(isAggregate, MkBool(true))) {
		active = OpCopy(MkBool(false))
	}
	amountMin := OpCopy(MkUndefined())
	_ = amountMin
	priceMin := OpCopy(MkUndefined())
	_ = priceMin
	priceMax := OpCopy(MkUndefined())
	_ = priceMax
	costMin := OpCopy(MkUndefined())
	_ = costMin
	pricePrecision := OpCopy(MkUndefined())
	_ = pricePrecision
	amountPrecision := OpCopy(MkUndefined())
	_ = amountPrecision
	for j := MkInteger(0); IsTrue(OpLw(j, filters.Length)); OpIncr(&j) {
		filter := *(filters).At(j)
		_ = filter
		filterType := this.SafeString(filter, MkString("filterType"))
		_ = filterType
		if IsTrue(OpEq2(filterType, MkString("LOT_SIZE"))) {
			amountMin = this.SafeNumber(filter, MkString("minQty"))
			amountPrecision = this.SafeNumber(filter, MkString("stepSize"))
		}
		if IsTrue(OpEq2(filterType, MkString("PRICE_FILTER"))) {
			priceMin = this.SafeNumber(filter, MkString("minPrice"))
			priceMax = this.SafeNumber(filter, MkString("maxPrice"))
			pricePrecision = this.SafeNumber(filter, MkString("tickSize"))
		}
	}
	if IsTrue(OpAnd(OpNotEq2(amountMin, MkUndefined()), OpNotEq2(priceMin, MkUndefined()))) {
		costMin = OpMulti(amountMin, priceMin)
	}
	precision := MkMap(&VarMap{
		"price":  pricePrecision,
		"amount": amountPrecision,
		"base":   this.SafeNumber(market, MkString("baseAssetPrecision")),
		"quote":  this.SafeNumber2(market, MkString("quotePrecision"), MkString("quoteAssetPrecision")),
	})
	_ = precision
	limits := MkMap(&VarMap{
		"amount": MkMap(&VarMap{
			"min": amountMin,
			"max": MkUndefined(),
		}),
		"price": MkMap(&VarMap{
			"min": priceMin,
			"max": priceMax,
		}),
		"cost": MkMap(&VarMap{
			"min": costMin,
			"max": MkUndefined(),
		}),
	})
	_ = limits
	return MkMap(&VarMap{
		"id":        id,
		"symbol":    symbol,
		"base":      base,
		"quote":     quote,
		"baseId":    baseId,
		"quoteId":   quoteId,
		"active":    active,
		"type":      type_,
		"spot":      spot,
		"future":    future,
		"option":    option,
		"margin":    margin,
		"inverse":   inverse,
		"precision": precision,
		"limits":    limits,
		"info":      market,
	})
}

func (this *Hbtc) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetBrokerInfo"), params)
	_ = response
	result := MkArray(&VarArray{})
	_ = result
	symbols := this.SafeValue(response, MkString("symbols"), MkArray(&VarArray{}))
	_ = symbols
	for i := MkInteger(0); IsTrue(OpLw(i, symbols.Length)); OpIncr(&i) {
		market := this.ParseMarket(*(symbols).At(i), MkString("spot"))
		_ = market
		result.Push(market)
	}
	options := this.SafeValue(response, MkString("options"), MkArray(&VarArray{}))
	_ = options
	for i := MkInteger(0); IsTrue(OpLw(i, options.Length)); OpIncr(&i) {
		market := this.ParseMarket(*(options).At(i), MkString("option"))
		_ = market
		result.Push(market)
	}
	contracts := this.SafeValue(response, MkString("contracts"), MkArray(&VarArray{}))
	_ = contracts
	for i := MkInteger(0); IsTrue(OpLw(i, contracts.Length)); OpIncr(&i) {
		market := this.ParseMarket(*(contracts).At(i), MkString("future"))
		_ = market
		result.Push(market)
	}
	return result
}

func (this *Hbtc) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("quoteGetDepth"), this.Extend(request, params))
	_ = response
	timestamp := this.SafeInteger(response, MkString("time"))
	_ = timestamp
	return this.ParseOrderBook(response, symbol, timestamp)
}

func (this *Hbtc) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("quoteGetTicker24hr"), this.Extend(request, params))
	_ = response
	return this.ParseTicker(response, market)
}

func (this *Hbtc) FetchBidAsk(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("quoteGetTickerBookTicker"), this.Extend(request, params))
	_ = response
	return this.ParseTicker(response, market)
}

func (this *Hbtc) FetchBidsAsks(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("quoteGetTickerBookTicker"), params)
	_ = response
	return this.ParseTickers(response, symbols)
}

func (this *Hbtc) FetchTickers(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	options := this.SafeValue(*this.At(MkString("options")), MkString("fetchTickers"), MkMap(&VarMap{}))
	_ = options
	defaultMethod := this.SafeString(options, MkString("method"), MkString("quoteGetTicker24hr"))
	_ = defaultMethod
	defaultType := this.SafeString(options, MkString("type"), MkString("spot"))
	_ = defaultType
	type_ := this.SafeString(params, MkString("type"), defaultType)
	_ = type_
	query := this.Omit(params, MkString("type"))
	_ = query
	method := OpCopy(defaultMethod)
	_ = method
	if IsTrue(OpEq2(type_, MkString("future"))) {
		method = MkString("quoteGetContractTicker24hr")
	} else {
		if IsTrue(OpEq2(type_, MkString("option"))) {
			method = MkString("quoteGetOptionTicker24hr")
		}
	}
	response := this.Call(method, query)
	_ = response
	return this.ParseTickers(response, symbols)
}

func (this *Hbtc) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	options := this.SafeValue(*this.At(MkString("options")), MkString("fetchBalance"), MkMap(&VarMap{}))
	_ = options
	defaultType := this.SafeString(options, MkString("type"), MkString("spot"))
	_ = defaultType
	type_ := this.SafeString(params, MkString("type"), defaultType)
	_ = type_
	query := this.Omit(params, MkString("type"))
	_ = query
	method := MkString("privateGetAccount")
	_ = method
	if IsTrue(OpEq2(type_, MkString("future"))) {
		method = MkString("contractGetAccount")
	} else {
		if IsTrue(OpEq2(type_, MkString("option"))) {
			method = MkString("optionGetAccount")
		}
	}
	response := this.Call(method, query)
	_ = response
	balances := this.SafeValue(response, MkString("balances"))
	_ = balances
	result := MkMap(&VarMap{
		"info":      response,
		"timestamp": MkUndefined(),
		"datetime":  MkUndefined(),
	})
	_ = result
	if IsTrue(OpNotEq2(balances, MkUndefined())) {
		for i := MkInteger(0); IsTrue(OpLw(i, balances.Length)); OpIncr(&i) {
			balance := *(balances).At(i)
			_ = balance
			currencyId := this.SafeString2(balance, MkString("asset"), MkString("tokenName"))
			_ = currencyId
			code := this.SafeCurrencyCode(currencyId)
			_ = code
			account := this.Account()
			_ = account
			*(account).At(MkString("free")) = this.SafeString(balance, MkString("free"))
			*(account).At(MkString("used")) = this.SafeString(balance, MkString("locked"))
			*(result).At(code) = OpCopy(account)
		}
	} else {
		currencyIds := GoGetKeys(response)
		_ = currencyIds
		for i := MkInteger(0); IsTrue(OpLw(i, currencyIds.Length)); OpIncr(&i) {
			currencyId := *(currencyIds).At(i)
			_ = currencyId
			code := this.SafeCurrencyCode(currencyId)
			_ = code
			balance := *(response).At(currencyId)
			_ = balance
			account := this.Account()
			_ = account
			*(account).At(MkString("free")) = this.SafeString(balance, MkString("availableMargin"))
			*(account).At(MkString("total")) = this.SafeString(balance, MkString("total"))
			*(result).At(code) = OpCopy(account)
		}
	}
	return this.ParseBalance(result)
}

func (this *Hbtc) FetchTrades(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkInteger(50))
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("quoteGetTrades"), this.Extend(request, params))
	_ = response
	return this.ParseTrades(response, market, since, limit)
}

func (this *Hbtc) ParseOHLCV(goArgs ...*Variant) *Variant {
	ohlcv := GoGetArg(goArgs, 0, MkUndefined())
	_ = ohlcv
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	return MkArray(&VarArray{
		this.SafeInteger(ohlcv, MkInteger(0)),
		this.SafeNumber(ohlcv, MkInteger(1)),
		this.SafeNumber(ohlcv, MkInteger(2)),
		this.SafeNumber(ohlcv, MkInteger(3)),
		this.SafeNumber(ohlcv, MkInteger(4)),
		this.SafeNumber(ohlcv, MkInteger(5)),
	})
}

func (this *Hbtc) FetchOHLCV(goArgs ...*Variant) *Variant {
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
		"symbol":   *(market).At(MkString("id")),
		"interval": *(*this.At(MkString("timeframes"))).At(timeframe),
	})
	_ = request
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("startTime")) = OpCopy(since)
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("quoteGetKlines"), this.Extend(request, params))
	_ = response
	return this.ParseOHLCVs(response, market, timeframe, since, limit)
}

func (this *Hbtc) FetchMyTrades(goArgs ...*Variant) *Variant {
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
	defaultType := this.SafeString(*this.At(MkString("options")), MkString("type"), MkString("spot"))
	_ = defaultType
	options := this.SafeValue(*this.At(MkString("options")), MkString("fetchMyTrades"), MkMap(&VarMap{}))
	_ = options
	fetchMyTradesType := this.SafeString(options, MkString("type"), defaultType)
	_ = fetchMyTradesType
	type_ := this.SafeString(params, MkString("type"), fetchMyTradesType)
	_ = type_
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
		*(request).At(MkString("symbol")) = *(market).At(MkString("id"))
		type_ = *(market).At(MkString("type"))
	}
	query := this.Omit(params, MkString("type"))
	_ = query
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	method := MkString("privateGetMyTrades")
	_ = method
	if IsTrue(OpEq2(type_, MkString("future"))) {
		method = MkString("contractGetMyTrades")
	} else {
		if IsTrue(OpEq2(type_, MkString("option"))) {
			method = MkString("optionGetMyTrades")
		} else {
			if IsTrue(OpEq2(symbol, MkUndefined())) {
				panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" fetchMyTrades() requires a `symbol` argument for "), OpAdd(type_, MkString(" markets"))))))
			}
			market := this.Market(symbol)
			_ = market
			*(request).At(MkString("symbol")) = *(market).At(MkString("id"))
			if IsTrue(OpNotEq2(since, MkUndefined())) {
				*(request).At(MkString("startTime")) = OpCopy(since)
			}
		}
	}
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("startTime")) = OpCopy(since)
	}
	response := this.Call(method, this.Extend(request, query))
	_ = response
	return this.ParseTrades(response, market, since, limit)
}

func (this *Hbtc) CreateOrder(goArgs ...*Variant) *Variant {
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
	orderSide := side.ToUpperCase()
	_ = orderSide
	orderType := type_.ToUpperCase()
	_ = orderType
	request := MkMap(&VarMap{
		"symbol": *(market).At(MkString("id")),
		"side":   orderSide,
	})
	_ = request
	query := OpCopy(params)
	_ = query
	method := MkString("privatePostOrder")
	_ = method
	if IsTrue(OpEq2(*(market).At(MkString("type")), MkString("future"))) {
		if IsTrue(OpAnd(OpNotEq2(orderSide, MkString("BUY_OPEN")), OpAnd(OpNotEq2(orderSide, MkString("SELL_OPEN")), OpAnd(OpNotEq2(orderSide, MkString("BUY_CLOSE")), OpNotEq2(orderSide, MkString("SELL_CLOSE")))))) {
			panic(NewNotSupported(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" createOrder() does not support order side "), OpAdd(side, OpAdd(MkString(" for "), OpAdd(*(market).At(MkString("type")), MkString(" markets, only BUY_OPEN, SELL_OPEN, BUY_CLOSE and SELL_CLOSE are supported"))))))))
		}
		if IsTrue(OpAnd(OpNotEq2(orderType, MkString("LIMIT")), OpNotEq2(orderType, MkString("STOP")))) {
			panic(NewNotSupported(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" createOrder() does not support order type "), OpAdd(type_, OpAdd(MkString(" for "), OpAdd(*(market).At(MkString("type")), MkString(" markets, only LIMIT and STOP are supported"))))))))
		}
		clientOrderId := this.SafeValue(params, MkString("clientOrderId"))
		_ = clientOrderId
		if IsTrue(OpEq2(clientOrderId, MkUndefined())) {
			panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" createOrder() requires a clientOrderId parameter for "), OpAdd(*(market).At(MkString("type")), MkString(" markets, supply clientOrderId in the params argument"))))))
		}
		leverage := this.SafeValue(params, MkString("leverage"))
		_ = leverage
		if IsTrue(OpAnd(OpEq2(leverage, MkUndefined()), OpOr(OpEq2(orderSide, MkString("BUY_OPEN")), OpEq2(orderSide, MkString("SELL_OPEN"))))) {
			panic(NewNotSupported(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" createOrder() requires a leverage parameter for "), OpAdd(*(market).At(MkString("type")), MkString(" markets if orderSide is BUY_OPEN or SELL_OPEN"))))))
		}
		method = MkString("contractPostOrder")
		priceType := this.SafeString(params, MkString("priceType"))
		_ = priceType
		if IsTrue(OpEq2(priceType, MkUndefined())) {
			*(request).At(MkString("price")) = this.PriceToPrecision(symbol, price)
		} else {
			*(request).At(MkString("priceType")) = OpCopy(priceType)
			if IsTrue(OpEq2(priceType, MkString("INPUT"))) {
				*(request).At(MkString("price")) = this.PriceToPrecision(symbol, price)
			}
		}
		*(request).At(MkString("orderType")) = type_.ToUpperCase()
		*(request).At(MkString("quantity")) = this.AmountToPrecision(symbol, amount)
		*(request).At(MkString("leverage")) = OpCopy(leverage)
		*(request).At(MkString("clientOrderId")) = OpCopy(clientOrderId)
	} else {
		if IsTrue(OpEq2(*(market).At(MkString("type")), MkString("option"))) {
			method = MkString("optionPostOrder")
		}
		newClientOrderId := this.SafeValue2(params, MkString("clientOrderId"), MkString("newClientOrderId"))
		_ = newClientOrderId
		if IsTrue(OpNotEq2(newClientOrderId, MkUndefined())) {
			*(request).At(MkString("newClientOrderId")) = OpCopy(newClientOrderId)
		}
		*(request).At(MkString("type")) = OpCopy(orderType)
		if IsTrue(OpEq2(type_, MkString("limit"))) {
			*(request).At(MkString("price")) = this.PriceToPrecision(symbol, price)
			*(request).At(MkString("quantity")) = this.AmountToPrecision(symbol, amount)
		} else {
			if IsTrue(OpEq2(type_, MkString("market"))) {
				if IsTrue(OpEq2(side, MkString("buy"))) {
					createMarketBuyOrderRequiresPrice := this.SafeValue(*this.At(MkString("options")), MkString("createMarketBuyOrderRequiresPrice"), MkBool(true))
					_ = createMarketBuyOrderRequiresPrice
					if IsTrue(createMarketBuyOrderRequiresPrice) {
						if IsTrue(OpNotEq2(price, MkUndefined())) {
							amount = OpMulti(amount, price)
						} else {
							panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")), MkString(" createOrder() requires the price argument with market buy orders to calculate total order cost (amount to spend), where cost = amount * price. Supply a price argument to createOrder() call if you want the cost to be calculated for you from price and amount, or, alternatively, add .options['createMarketBuyOrderRequiresPrice'] = false and supply the total cost value in the 'amount' argument (the exchange-specific behaviour)"))))
						}
					}
					precision := *(*(market).At(MkString("precision"))).At(MkString("price"))
					_ = precision
					*(request).At(MkString("quantity")) = this.DecimalToPrecision(amount, TRUNCATE, precision, *this.At(MkString("precisionMode")))
				} else {
					*(request).At(MkString("quantity")) = this.AmountToPrecision(symbol, amount)
				}
			}
		}
	}
	query = this.Omit(query, MkArray(&VarArray{
		MkString("clientOrderId"),
		MkString("newClientOrderId"),
	}))
	response := this.Call(method, this.Extend(request, query))
	_ = response
	return this.ParseOrder(response, market)
}

func (this *Hbtc) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	clientOrderId := this.SafeValue2(params, MkString("origClientOrderId"), MkString("clientOrderId"))
	_ = clientOrderId
	request := MkMap(&VarMap{})
	_ = request
	defaultType := this.SafeString(*this.At(MkString("options")), MkString("type"), MkString("spot"))
	_ = defaultType
	options := this.SafeValue(*this.At(MkString("options")), MkString("cancelOrder"), MkMap(&VarMap{}))
	_ = options
	cancelOrderType := this.SafeString(options, MkString("type"), defaultType)
	_ = cancelOrderType
	type_ := this.SafeString(params, MkString("type"), cancelOrderType)
	_ = type_
	query := this.Omit(params, MkString("type"))
	_ = query
	if IsTrue(OpNotEq2(clientOrderId, MkUndefined())) {
		*(request).At(MkString("origClientOrderId")) = OpCopy(clientOrderId)
		query = this.Omit(query, MkArray(&VarArray{
			MkString("origClientOrderId"),
			MkString("clientOrderId"),
		}))
	} else {
		*(request).At(MkString("orderId")) = OpCopy(id)
	}
	method := MkString("privateDeleteOrder")
	_ = method
	orderType := this.SafeString(query, MkString("orderType"))
	_ = orderType
	if IsTrue(OpNotEq2(orderType, MkUndefined())) {
		type_ = MkString("future")
	}
	if IsTrue(OpEq2(type_, MkString("future"))) {
		method = MkString("contractDeleteOrderCancel")
		if IsTrue(OpEq2(orderType, MkUndefined())) {
			panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" cancelOrder() requires an orderType parameter, pass the { 'orderType': 'LIMIT' } or { 'orderType': 'STOP' } in params argument"))))
		}
		*(request).At(MkString("orderType")) = OpCopy(orderType)
	} else {
		if IsTrue(OpEq2(type_, MkString("option"))) {
			method = MkString("optionDeleteOrderCancel")
		}
	}
	response := this.Call(method, this.Extend(request, query))
	_ = response
	return this.ParseOrder(response)
}

func (this *Hbtc) FetchOpenOrders(goArgs ...*Variant) *Variant {
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
	defaultType := this.SafeString(*this.At(MkString("options")), MkString("type"), MkString("spot"))
	_ = defaultType
	options := this.SafeValue(*this.At(MkString("options")), MkString("fetchOpenOrders"), MkMap(&VarMap{}))
	_ = options
	fetchOpenOrdersType := this.SafeString(options, MkString("type"), defaultType)
	_ = fetchOpenOrdersType
	type_ := this.SafeString(params, MkString("type"), fetchOpenOrdersType)
	_ = type_
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
		*(request).At(MkString("symbol")) = *(market).At(MkString("id"))
		type_ = *(market).At(MkString("type"))
	}
	query := this.Omit(params, MkString("type"))
	_ = query
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	method := MkString("privateGetOpenOrders")
	_ = method
	if IsTrue(OpEq2(type_, MkString("future"))) {
		method = MkString("contractGetOpenOrders")
	} else {
		if IsTrue(OpEq2(type_, MkString("option"))) {
			method = MkString("optionGetOpenOrders")
		}
	}
	response := this.Call(method, this.Extend(request, query))
	_ = response
	return this.ParseOrders(response, market, since, limit)
}

func (this *Hbtc) FetchClosedOrders(goArgs ...*Variant) *Variant {
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
	defaultType := this.SafeString(*this.At(MkString("options")), MkString("type"), MkString("spot"))
	_ = defaultType
	options := this.SafeValue(*this.At(MkString("options")), MkString("fetchClosedOrders"), MkMap(&VarMap{}))
	_ = options
	fetchClosedOrdersType := this.SafeString(options, MkString("type"), defaultType)
	_ = fetchClosedOrdersType
	type_ := this.SafeString(params, MkString("type"), fetchClosedOrdersType)
	_ = type_
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
		*(request).At(MkString("symbol")) = *(market).At(MkString("id"))
		type_ = *(market).At(MkString("type"))
	}
	query := this.Omit(params, MkString("type"))
	_ = query
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("startTime")) = OpCopy(since)
	}
	method := MkString("privateGetHistoryOrders")
	_ = method
	if IsTrue(OpEq2(type_, MkString("future"))) {
		method = MkString("contractGetHistoryOrders")
	} else {
		if IsTrue(OpEq2(type_, MkString("option"))) {
			method = MkString("optionGetHistoryOrders")
		}
	}
	response := this.Call(method, this.Extend(request, query))
	_ = response
	return this.ParseOrders(response, market, since, limit)
}

func (this *Hbtc) FetchOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	clientOrderId := this.SafeValue2(params, MkString("origClientOrderId"), MkString("clientOrderId"))
	_ = clientOrderId
	request := MkMap(&VarMap{})
	_ = request
	defaultType := this.SafeString(*this.At(MkString("options")), MkString("type"), MkString("spot"))
	_ = defaultType
	options := this.SafeValue(*this.At(MkString("options")), MkString("fetchOrder"), MkMap(&VarMap{}))
	_ = options
	fetchOrderType := this.SafeString(options, MkString("type"), defaultType)
	_ = fetchOrderType
	type_ := this.SafeString(params, MkString("type"), fetchOrderType)
	_ = type_
	query := this.Omit(params, MkString("type"))
	_ = query
	if IsTrue(OpNotEq2(clientOrderId, MkUndefined())) {
		*(request).At(MkString("origClientOrderId")) = OpCopy(clientOrderId)
		query = this.Omit(query, MkArray(&VarArray{
			MkString("origClientOrderId"),
			MkString("clientOrderId"),
		}))
	} else {
		*(request).At(MkString("orderId")) = OpCopy(id)
	}
	method := MkString("privateGetOrder")
	_ = method
	if IsTrue(OpEq2(type_, MkString("future"))) {
		method = MkString("contractGetGetOrder")
	} else {
		if IsTrue(OpEq2(type_, MkString("option"))) {
			method = MkString("optionGetGetOrder")
		}
	}
	response := this.Call(method, this.Extend(request, query))
	_ = response
	return this.ParseOrder(response)
}

func (this *Hbtc) FetchDeposits(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	currency := OpCopy(MkUndefined())
	_ = currency
	request := MkMap(&VarMap{})
	_ = request
	if IsTrue(OpNotEq2(code, MkUndefined())) {
		currency = this.Currency(code)
	}
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("startTime")) = OpCopy(since)
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("privateGetDepositOrders"), this.Extend(request, params))
	_ = response
	return this.ParseTransactions(response, currency, since, limit)
}

func (this *Hbtc) FetchWithdrawals(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	currency := OpCopy(MkUndefined())
	_ = currency
	request := MkMap(&VarMap{})
	_ = request
	if IsTrue(OpNotEq2(code, MkUndefined())) {
		currency = this.Currency(code)
		*(request).At(MkString("token")) = *(currency).At(MkString("id"))
	}
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("startTime")) = OpCopy(since)
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("privateGetWithdrawalOrders"), this.Extend(request, params))
	_ = response
	return this.ParseTransactions(response, currency, since, limit)
}

func (this *Hbtc) Withdraw(goArgs ...*Variant) *Variant {
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
	clientOrderId := this.SafeString(params, MkString("clientOrderId"), this.Uuid())
	_ = clientOrderId
	request := MkMap(&VarMap{
		"clientOrderId":    clientOrderId,
		"tokenId":          *(currency).At(MkString("id")),
		"address":          address,
		"withdrawQuantity": amount,
	})
	_ = request
	if IsTrue(OpNotEq2(tag, MkUndefined())) {
		*(request).At(MkString("addressExt")) = OpCopy(tag)
	}
	response := this.Call(MkString("privatePostWithdraw"), this.Extend(request, params))
	_ = response
	return MkMap(&VarMap{
		"info": response,
		"id":   this.SafeString(response, MkString("orderId")),
	})
}

func (this *Hbtc) FetchAccounts(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("privatePostSubAccountQuery"), params)
	_ = response
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
		account := *(response).At(i)
		_ = account
		accountId := this.SafeString(account, MkString("accountId"))
		_ = accountId
		accountType := this.SafeString(account, MkString("accountType"))
		_ = accountType
		type_ := OpCopy(accountType)
		_ = type_
		if IsTrue(OpEq2(accountType, MkString("1"))) {
			type_ = MkString("spot")
		} else {
			if IsTrue(OpEq2(accountType, MkString("2"))) {
				type_ = MkString("option")
			} else {
				if IsTrue(OpEq2(accountType, MkString("3"))) {
					type_ = MkString("future")
				}
			}
		}
		result.Push(MkMap(&VarMap{
			"id":       accountId,
			"type":     type_,
			"currency": MkUndefined(),
			"info":     account,
		}))
	}
	return result
}

func (this *Hbtc) FetchLedger(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{
		"accountType":  MkInteger(1),
		"accountIndex": MkInteger(0),
		"fromFlowId":   MkString(""),
		"endFlowId":    MkString(""),
		"endTime":      MkInteger(1588450533040),
	})
	_ = request
	currency := OpCopy(MkUndefined())
	_ = currency
	if IsTrue(OpNotEq2(code, MkUndefined())) {
		currency = this.Currency(code)
		*(request).At(MkString("tokenId")) = *(currency).At(MkString("id"))
	}
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("startTime")) = OpCopy(since)
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("privateGetBalanceFlow"), this.Extend(request, params))
	_ = response
	return this.ParseLedger(response, currency, since, limit)
}

func (this *Hbtc) ParseLedgerEntry(goArgs ...*Variant) *Variant {
	item := GoGetArg(goArgs, 0, MkUndefined())
	_ = item
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	currencyId := this.SafeString(item, MkString("tokenId"))
	_ = currencyId
	code := this.SafeCurrencyCode(currencyId, currency)
	_ = code
	amount := this.SafeNumber(item, MkString("change"))
	_ = amount
	after := this.SafeNumber(item, MkString("total"))
	_ = after
	direction := OpTrinary(OpLw(amount, MkInteger(0)), MkString("out"), MkString("in"))
	_ = direction
	before := OpCopy(MkUndefined())
	_ = before
	if IsTrue(OpAnd(OpNotEq2(after, MkUndefined()), OpNotEq2(amount, MkUndefined()))) {
		difference := OpTrinary(OpEq2(direction, MkString("out")), amount, OpNeg(amount))
		_ = difference
		before = this.Sum(after, difference)
	}
	timestamp := this.SafeInteger(item, MkString("created"))
	_ = timestamp
	type_ := this.ParseLedgerEntryType(this.SafeString(item, MkString("flowType")))
	_ = type_
	id := this.SafeString(item, MkString("id"))
	_ = id
	account := this.SafeString(item, MkString("accountId"))
	_ = account
	return MkMap(&VarMap{
		"id":               id,
		"currency":         code,
		"account":          account,
		"referenceAccount": MkUndefined(),
		"referenceId":      MkUndefined(),
		"status":           MkUndefined(),
		"amount":           amount,
		"before":           before,
		"after":            after,
		"fee":              MkUndefined(),
		"direction":        direction,
		"timestamp":        timestamp,
		"datetime":         this.Iso8601(timestamp),
		"type":             type_,
		"info":             item,
	})
}

func (this *Hbtc) ParseLedgerEntryType(goArgs ...*Variant) *Variant {
	type_ := GoGetArg(goArgs, 0, MkUndefined())
	_ = type_
	types := MkMap(&VarMap{
		"TRADE":                     MkString("trade"),
		"FEE":                       MkString("fee"),
		"TRANSFER":                  MkString("transfer"),
		"DEPOSIT":                   MkString("transaction"),
		"MAKER_REWARD":              MkString("rebate"),
		"PNL":                       MkString("pnl"),
		"SETTLEMENT":                MkString("settlement"),
		"LIQUIDATION":               MkString("liquidation"),
		"FUNDING_SETTLEMENT":        MkString("settlement"),
		"USER_ACCOUNT_TRANSFER":     MkString("transfer"),
		"OTC_BUY_COIN":              MkString("trade"),
		"OTC_SELL_COIN":             MkString("trade"),
		"OTC_FEE":                   MkString("fee"),
		"OTC_TRADE":                 MkString("trade"),
		"ACTIVITY_AWARD":            MkString("referral"),
		"INVITATION_REFERRAL_BONUS": MkString("referral"),
		"REGISTER_BONUS":            MkString("referral"),
		"AIRDROP":                   MkString("airdrop"),
		"MINE_REWARD":               MkString("reward"),
	})
	_ = types
	return this.SafeString(types, type_, type_)
}

func (this *Hbtc) ParseTransactionStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"BROKER_AUDITING_STATUS":    MkString("pending"),
		"BROKER_REJECT_STATUS":      MkString("failed"),
		"AUDITING_STATUS":           MkString("pending"),
		"AUDIT_REJECT_STATUS":       MkString("failed"),
		"PROCESSING_STATUS":         MkString("pending"),
		"WITHDRAWAL_SUCCESS_STATUS": MkString("ok"),
		"WITHDRAWAL_FAILURE_STATUS": MkString("failed"),
		"BLOCK_MINING_STATUS":       MkString("ok"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Hbtc) ParseTransaction(goArgs ...*Variant) *Variant {
	transaction := GoGetArg(goArgs, 0, MkUndefined())
	_ = transaction
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	id := this.SafeString(transaction, MkString("orderId"))
	_ = id
	address := this.SafeString(transaction, MkString("address"))
	_ = address
	tag := this.SafeString2(transaction, MkString("addressExt"), MkString("addressTag"))
	_ = tag
	if IsTrue(OpNotEq2(tag, MkUndefined())) {
		if IsTrue(OpLw(tag.Length, MkInteger(1))) {
			tag = OpCopy(MkUndefined())
		}
	}
	addressFrom := this.SafeString(transaction, MkString("fromAddress"))
	_ = addressFrom
	tagFrom := this.SafeString(transaction, MkString("fromAddressTag"))
	_ = tagFrom
	if IsTrue(OpNotEq2(tagFrom, MkUndefined())) {
		if IsTrue(OpLw(tagFrom.Length, MkInteger(1))) {
			tagFrom = OpCopy(MkUndefined())
		}
	}
	currencyId := this.SafeString(transaction, MkString("tokenId"))
	_ = currencyId
	code := this.SafeCurrencyCode(currencyId, currency)
	_ = code
	timestamp := this.SafeInteger(transaction, MkString("time"))
	_ = timestamp
	txid := this.SafeString(transaction, MkString("txid"))
	_ = txid
	if IsTrue(OpEq2(txid, MkString(""))) {
		txid = OpCopy(MkUndefined())
	}
	type_ := OpCopy(MkUndefined())
	_ = type_
	status := this.ParseTransactionStatus(this.SafeString(transaction, MkString("statusCode")))
	_ = status
	if IsTrue(OpEq2(status, MkUndefined())) {
		type_ = MkString("deposit")
		status = MkString("ok")
	} else {
		type_ = MkString("withdrawal")
	}
	amount := this.SafeNumber(transaction, MkString("quantity"))
	_ = amount
	feeCost := this.SafeNumber(transaction, MkString("fee"))
	_ = feeCost
	fee := OpCopy(MkUndefined())
	_ = fee
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		feeCurrencyId := this.SafeString(transaction, MkString("feeTokenId"))
		_ = feeCurrencyId
		feeCurrencyCode := this.SafeCurrencyCode(feeCurrencyId)
		_ = feeCurrencyCode
		fee = MkMap(&VarMap{
			"currency": feeCurrencyCode,
			"cost":     feeCost,
		})
	}
	return MkMap(&VarMap{
		"info":        transaction,
		"id":          id,
		"txid":        txid,
		"timestamp":   timestamp,
		"datetime":    this.Iso8601(timestamp),
		"addressFrom": addressFrom,
		"address":     address,
		"addressTo":   address,
		"tagFrom":     tagFrom,
		"tag":         tag,
		"tagTo":       tag,
		"type":        type_,
		"amount":      amount,
		"currency":    code,
		"status":      status,
		"updated":     MkUndefined(),
		"fee":         fee,
	})
}

func (this *Hbtc) ParseTicker(goArgs ...*Variant) *Variant {
	ticker := GoGetArg(goArgs, 0, MkUndefined())
	_ = ticker
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	marketId := this.SafeString(ticker, MkString("symbol"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market)
	_ = symbol
	timestamp := this.SafeInteger(ticker, MkString("time"))
	_ = timestamp
	open := this.SafeNumber(ticker, MkString("openPrice"))
	_ = open
	close := this.SafeNumber(ticker, MkString("lastPrice"))
	_ = close
	change := OpCopy(MkUndefined())
	_ = change
	percentage := OpCopy(MkUndefined())
	_ = percentage
	average := OpCopy(MkUndefined())
	_ = average
	if IsTrue(OpAnd(OpNotEq2(open, MkUndefined()), OpNotEq2(close, MkUndefined()))) {
		change = OpSub(close, open)
		average = OpDiv(this.Sum(open, close), MkInteger(2))
		if IsTrue(OpAnd(OpNotEq2(close, MkUndefined()), OpGt(close, MkInteger(0)))) {
			percentage = OpMulti(OpDiv(change, open), MkInteger(100))
		}
	}
	quoteVolume := this.SafeNumber(ticker, MkString("quoteVolume"))
	_ = quoteVolume
	baseVolume := this.SafeNumber(ticker, MkString("volume"))
	_ = baseVolume
	vwap := this.Vwap(baseVolume, quoteVolume)
	_ = vwap
	return MkMap(&VarMap{
		"symbol":        symbol,
		"timestamp":     timestamp,
		"datetime":      this.Iso8601(timestamp),
		"high":          this.SafeNumber(ticker, MkString("highPrice")),
		"low":           this.SafeNumber(ticker, MkString("lowPrice")),
		"bid":           this.SafeNumber2(ticker, MkString("bestBidPrice"), MkString("bidPrice")),
		"bidVolume":     this.SafeNumber(ticker, MkString("bidQty")),
		"ask":           this.SafeNumber2(ticker, MkString("bestAskPrice"), MkString("askPrice")),
		"askVolume":     this.SafeNumber(ticker, MkString("askQty")),
		"vwap":          vwap,
		"open":          open,
		"close":         close,
		"last":          close,
		"previousClose": MkUndefined(),
		"change":        change,
		"percentage":    percentage,
		"average":       average,
		"baseVolume":    baseVolume,
		"quoteVolume":   quoteVolume,
		"info":          ticker,
	})
}

func (this *Hbtc) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	id := this.SafeString(trade, MkString("id"))
	_ = id
	timestamp := this.SafeNumber(trade, MkString("time"))
	_ = timestamp
	type_ := OpCopy(MkUndefined())
	_ = type_
	orderId := this.SafeString(trade, MkString("orderId"))
	_ = orderId
	priceString := this.SafeString(trade, MkString("price"))
	_ = priceString
	amountString := this.SafeString(trade, MkString("qty"))
	_ = amountString
	price := this.ParseNumber(priceString)
	_ = price
	amount := this.ParseNumber(amountString)
	_ = amount
	cost := this.ParseNumber(Precise.StringMul(priceString, amountString))
	_ = cost
	side := OpCopy(MkUndefined())
	_ = side
	takerOrMaker := OpCopy(MkUndefined())
	_ = takerOrMaker
	if IsTrue(OpHasMember(MkString("isBuyerMaker"), trade)) {
		side = OpTrinary(*(trade).At(MkString("isBuyerMaker")), MkString("sell"), MkString("buy"))
	} else {
		isMaker := this.SafeValue(trade, MkString("isMaker"))
		_ = isMaker
		if IsTrue(OpNotEq2(isMaker, MkUndefined())) {
			takerOrMaker = OpTrinary(isMaker, MkString("maker"), MkString("taker"))
		}
		isBuyer := this.SafeValue(trade, MkString("isBuyer"))
		_ = isBuyer
		side = OpTrinary(isBuyer, MkString("buy"), MkString("sell"))
	}
	fee := OpCopy(MkUndefined())
	_ = fee
	feeCost := this.SafeNumber(trade, MkString("commission"))
	_ = feeCost
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		feeCurrencyId := this.SafeString(trade, MkString("commissionAsset"))
		_ = feeCurrencyId
		feeCurrencyCode := this.SafeCurrencyCode(feeCurrencyId)
		_ = feeCurrencyCode
		fee = MkMap(&VarMap{
			"cost":     feeCost,
			"currency": feeCurrencyCode,
		})
	}
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(OpAnd(OpEq2(symbol, MkUndefined()), OpNotEq2(market, MkUndefined()))) {
		symbol = *(market).At(MkString("symbol"))
	}
	return MkMap(&VarMap{
		"id":           id,
		"info":         trade,
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       symbol,
		"type":         type_,
		"order":        orderId,
		"side":         side,
		"takerOrMaker": takerOrMaker,
		"price":        price,
		"amount":       amount,
		"cost":         cost,
		"fee":          fee,
	})
}

func (this *Hbtc) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	id := this.SafeString(order, MkString("orderId"))
	_ = id
	clientOrderId := this.SafeString(order, MkString("clientOrderId"))
	_ = clientOrderId
	timestamp := this.SafeInteger(order, MkString("time"))
	_ = timestamp
	if IsTrue(OpEq2(timestamp, MkUndefined())) {
		timestamp = this.SafeInteger(order, MkString("transactTime"))
	}
	marketId := this.SafeString(order, MkString("symbol"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market)
	_ = symbol
	type_ := this.SafeStringLower(order, MkString("type"))
	_ = type_
	side := this.SafeStringLower(order, MkString("side"))
	_ = side
	price := this.SafeNumber(order, MkString("price"))
	_ = price
	average := this.SafeNumber(order, MkString("avgPrice"))
	_ = average
	amount := OpCopy(MkUndefined())
	_ = amount
	cost := this.SafeNumber(order, MkString("cummulativeQuoteQty"))
	_ = cost
	filled := OpCopy(MkUndefined())
	_ = filled
	remaining := OpCopy(MkUndefined())
	_ = remaining
	if IsTrue(OpEq2(type_, MkUndefined())) {
		type_ = this.SafeStringLower(order, MkString("orderType"))
		if IsTrue(OpAnd(OpNotEq2(market, MkUndefined()), *(market).At(MkString("inverse")))) {
			cost = this.SafeNumber(order, MkString("executedQty"))
			amount = OpCopy(MkUndefined())
		}
		if IsTrue(OpEq2(cost, MkNumber(0.0))) {
			filled = MkInteger(0)
		}
	} else {
		amount = this.SafeNumber(order, MkString("origQty"))
		if IsTrue(OpEq2(type_, MkString("market"))) {
			price = OpCopy(MkUndefined())
			if IsTrue(OpEq2(side, MkString("buy"))) {
				amount = OpCopy(MkUndefined())
			}
		}
		filled = this.SafeNumber(order, MkString("executedQty"))
		if IsTrue(OpNotEq2(filled, MkUndefined())) {
			if IsTrue(OpNotEq2(amount, MkUndefined())) {
				remaining = OpSub(amount, filled)
			}
		}
	}
	if IsTrue(OpEq2(average, MkNumber(0.0))) {
		average = OpCopy(MkUndefined())
	}
	status := this.ParseOrderStatus(this.SafeString(order, MkString("status")))
	_ = status
	timeInForce := this.SafeString(order, MkString("timeInForce"))
	_ = timeInForce
	stopPrice := this.SafeNumber(order, MkString("stopPrice"))
	_ = stopPrice
	result := MkMap(&VarMap{
		"info":               order,
		"id":                 id,
		"clientOrderId":      clientOrderId,
		"timestamp":          timestamp,
		"datetime":           this.Iso8601(timestamp),
		"lastTradeTimestamp": MkUndefined(),
		"symbol":             symbol,
		"type":               type_,
		"timeInForce":        timeInForce,
		"side":               side,
		"price":              price,
		"stopPrice":          stopPrice,
		"average":            average,
		"cost":               cost,
		"amount":             amount,
		"filled":             filled,
		"remaining":          remaining,
		"status":             status,
		"trades":             MkUndefined(),
		"fee":                MkUndefined(),
		"fees":               MkUndefined(),
	})
	_ = result
	fees := this.SafeValue(order, MkString("fees"), MkArray(&VarArray{}))
	_ = fees
	numFees := OpCopy(fees.Length)
	_ = numFees
	if IsTrue(OpGt(numFees, MkInteger(0))) {
		*(result).At(MkString("fees")) = MkArray(&VarArray{})
		for i := MkInteger(0); IsTrue(OpLw(i, fees.Length)); OpIncr(&i) {
			feeCost := this.SafeNumber(*(fees).At(i), MkString("fee"))
			_ = feeCost
			if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
				feeCurrencyId := this.SafeString(*(fees).At(i), MkString("feeToken"))
				_ = feeCurrencyId
				feeCurrencyCode := this.SafeCurrencyCode(feeCurrencyId)
				_ = feeCurrencyCode
				(*(result).At(MkString("fees"))).Call(MkString("push"), MkMap(&VarMap{
					"cost":     feeCost,
					"currency": feeCurrencyCode,
				}))
			}
		}
	}
	return result
}

func (this *Hbtc) ParseOrderStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"NEW":              MkString("open"),
		"CANCELED":         MkString("canceled"),
		"FILLED":           MkString("closed"),
		"PARTIALLY_FILLED": MkString("open"),
		"PENDING_CANCEL":   MkString("canceled"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Hbtc) Sign(goArgs ...*Variant) *Variant {
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
	url := OpAdd(*(*(*this.At(MkString("urls"))).At(MkString("api"))).At(api), OpAdd(MkString("/"), OpAdd(*this.At(MkString("version")), OpAdd(MkString("/"), this.ImplodeParams(path, params)))))
	_ = url
	query := this.Omit(params, this.ExtractParams(path))
	_ = query
	isPublicContract := OpAnd(OpEq2(api, MkString("contract")), OpOr(OpEq2(path, MkString("insurance")), OpEq2(path, MkString("fundingRate"))))
	_ = isPublicContract
	if IsTrue(OpOr(OpEq2(api, MkString("public")), OpOr(OpEq2(api, MkString("quote")), isPublicContract))) {
		if IsTrue(*(GoGetKeys(params)).At(MkString("length"))) {
			OpAddAssign(&url, OpAdd(MkString("?"), this.Urlencode(params)))
		}
	} else {
		timestamp := this.Milliseconds()
		_ = timestamp
		this.CheckRequiredCredentials()
		request := this.Extend(MkMap(&VarMap{"timestamp": timestamp}), query)
		_ = request
		auth := this.Urlencode(request)
		_ = auth
		signature := this.Hmac(this.Encode(auth), this.Encode(*this.At(MkString("secret"))), MkString("sha256"))
		_ = signature
		*(request).At(MkString("signature")) = OpCopy(signature)
		headers = MkMap(&VarMap{"X-BH-APIKEY": *this.At(MkString("apiKey"))})
		if IsTrue(OpEq2(method, MkString("POST"))) {
			body = this.Urlencode(request)
			headers = this.Extend(MkMap(&VarMap{"Content-Type": MkString("application/x-www-form-urlencoded")}), headers)
		} else {
			OpAddAssign(&url, OpAdd(MkString("?"), this.Urlencode(request)))
		}
	}
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Hbtc) HandleErrors(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpHasMember(MkString("code"), response)) {
		code := this.SafeString(response, MkString("code"))
		_ = code
		if IsTrue(OpNotEq2(code, MkString("0"))) {
			feedback := OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))
			_ = feedback
			this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("exact")), code, feedback)
			panic(NewExchangeError(feedback))
		}
	}
	return MkUndefined()
}
