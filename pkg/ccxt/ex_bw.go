package ccxt

import ()

type Bw struct {
	*ExchangeBase
}

var _ Exchange = (*Bw)(nil)

func init() {
	exchange := &Bw{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Bw) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("bw"),
		"name": MkString("BW"),
		"countries": MkArray(&VarArray{
			MkString("CN"),
		}),
		"rateLimit": MkInteger(1500),
		"version":   MkString("v1"),
		"has": MkMap(&VarMap{
			"cancelAllOrders":      MkBool(false),
			"cancelOrder":          MkBool(true),
			"cancelOrders":         MkBool(false),
			"CORS":                 MkBool(false),
			"createDepositAddress": MkBool(false),
			"createLimitOrder":     MkBool(true),
			"createMarketOrder":    MkBool(false),
			"createOrder":          MkBool(true),
			"deposit":              MkBool(false),
			"editOrder":            MkBool(false),
			"fetchBalance":         MkBool(true),
			"fetchBidsAsks":        MkBool(false),
			"fetchClosedOrders":    MkBool(true),
			"fetchCurrencies":      MkBool(true),
			"fetchDepositAddress":  MkBool(true),
			"fetchDeposits":        MkBool(true),
			"fetchFundingFees":     MkBool(false),
			"fetchL2OrderBook":     MkBool(false),
			"fetchLedger":          MkBool(false),
			"fetchMarkets":         MkBool(true),
			"fetchMyTrades":        MkBool(false),
			"fetchOHLCV":           MkBool(true),
			"fetchOpenOrders":      MkBool(true),
			"fetchOrder":           MkBool(true),
			"fetchOrderBook":       MkBool(true),
			"fetchOrderBooks":      MkBool(false),
			"fetchOrders":          MkBool(true),
			"fetchTicker":          MkBool(true),
			"fetchTickers":         MkBool(true),
			"fetchTrades":          MkBool(true),
			"fetchTradingFee":      MkBool(false),
			"fetchTradingFees":     MkBool(false),
			"fetchTradingLimits":   MkBool(false),
			"fetchTransactions":    MkBool(false),
			"fetchWithdrawals":     MkBool(true),
			"privateAPI":           MkBool(false),
			"publicAPI":            MkBool(false),
			"withdraw":             MkBool(false),
		}),
		"timeframes": MkMap(&VarMap{
			"1m":  MkString("1M"),
			"5m":  MkString("5M"),
			"15m": MkString("15M"),
			"30m": MkString("30M"),
			"1h":  MkString("1H"),
			"1w":  MkString("1W"),
		}),
		"hostname": MkString("bw.com"),
		"urls": MkMap(&VarMap{
			"logo":     MkString("https://user-images.githubusercontent.com/1294454/69436317-31128c80-0d52-11ea-91d1-eb7bb5818812.jpg"),
			"api":      MkString("https://www.{hostname}"),
			"www":      MkString("https://www.bw.com"),
			"doc":      MkString("https://github.com/bw-exchange/api_docs_en/wiki"),
			"fees":     MkString("https://www.bw.com/feesRate"),
			"referral": MkString("https://www.bw.com/regGetCommission/N3JuT1R3bWxKTE0"),
		}),
		"requiredCredentials": MkMap(&VarMap{
			"apiKey": MkBool(true),
			"secret": MkBool(true),
		}),
		"fees": MkMap(&VarMap{
			"trading": MkMap(&VarMap{
				"tierBased":  MkBool(false),
				"percentage": MkBool(true),
				"taker":      this.ParseNumber(MkString("0.002")),
				"maker":      this.ParseNumber(MkString("0.002")),
			}),
			"funding": MkMap(&VarMap{}),
		}),
		"exceptions": MkMap(&VarMap{"exact": MkMap(&VarMap{
			"999":   AuthenticationError,
			"1000":  ExchangeNotAvailable,
			"2012":  OrderNotFound,
			"5017":  BadSymbol,
			"10001": RateLimitExceeded,
		})}),
		"api": MkMap(&VarMap{
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("api/data/v1/klines"),
				MkString("api/data/v1/ticker"),
				MkString("api/data/v1/tickers"),
				MkString("api/data/v1/trades"),
				MkString("api/data/v1/entrusts"),
				MkString("exchange/config/controller/website/marketcontroller/getByWebId"),
				MkString("exchange/config/controller/website/currencycontroller/getCurrencyList"),
			})}),
			"private": MkMap(&VarMap{
				"get": MkArray(&VarArray{
					MkString("exchange/entrust/controller/website/EntrustController/getEntrustById"),
					MkString("exchange/entrust/controller/website/EntrustController/getUserEntrustRecordFromCacheWithPage"),
					MkString("exchange/entrust/controller/website/EntrustController/getUserEntrustList"),
					MkString("exchange/fund/controller/website/fundwebsitecontroller/getwithdrawaddress"),
					MkString("exchange/fund/controller/website/fundwebsitecontroller/getpayoutcoinrecord"),
					MkString("exchange/entrust/controller/website/EntrustController/getUserEntrustList"),
				}),
				"post": MkArray(&VarArray{
					MkString("exchange/fund/controller/website/fundcontroller/getPayinAddress"),
					MkString("exchange/fund/controller/website/fundcontroller/getPayinCoinRecord"),
					MkString("exchange/fund/controller/website/fundcontroller/findbypage"),
					MkString("exchange/entrust/controller/website/EntrustController/addEntrust"),
					MkString("exchange/entrust/controller/website/EntrustController/cancelEntrust"),
				}),
			}),
		}),
	}))
}

func (this *Bw) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetExchangeConfigControllerWebsiteMarketcontrollerGetByWebId"), params)
	_ = response
	markets := this.SafeValue(response, MkString("datas"), MkArray(&VarArray{}))
	_ = markets
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, markets.Length)); OpIncr(&i) {
		market := *(markets).At(i)
		_ = market
		id := this.SafeString(market, MkString("marketId"))
		_ = id
		numericId := parseInt(id)
		_ = numericId
		name := this.SafeStringUpper(market, MkString("name"))
		_ = name
		base := MkUndefined()
		_ = base
		quote := MkUndefined()
		_ = quote
		ArrayUnpack(name.Split(MkString("_")), &base, &quote)
		base = this.SafeCurrencyCode(base)
		quote = this.SafeCurrencyCode(quote)
		baseId := this.SafeString(market, MkString("sellerCurrencyId"))
		_ = baseId
		quoteId := this.SafeString(market, MkString("buyerCurrencyId"))
		_ = quoteId
		baseNumericId := parseInt(baseId)
		_ = baseNumericId
		quoteNumericId := parseInt(quoteId)
		_ = quoteNumericId
		symbol := OpAdd(base, OpAdd(MkString("/"), quote))
		_ = symbol
		state := this.SafeInteger(market, MkString("state"))
		_ = state
		active := OpEq2(state, MkInteger(1))
		_ = active
		fee := this.SafeNumber(market, MkString("defaultFee"))
		_ = fee
		result.Push(MkMap(&VarMap{
			"id":             id,
			"active":         active,
			"numericId":      numericId,
			"symbol":         symbol,
			"base":           base,
			"quote":          quote,
			"baseId":         baseId,
			"quoteId":        quoteId,
			"baseNumericId":  baseNumericId,
			"quoteNumericId": quoteNumericId,
			"maker":          fee,
			"taker":          fee,
			"info":           market,
			"precision": MkMap(&VarMap{
				"amount": this.SafeInteger(market, MkString("amountDecimal")),
				"price":  this.SafeInteger(market, MkString("priceDecimal")),
			}),
			"limits": MkMap(&VarMap{
				"amount": MkMap(&VarMap{
					"min": this.SafeNumber(market, MkString("minAmount")),
					"max": MkUndefined(),
				}),
				"price": MkMap(&VarMap{
					"min": MkInteger(0),
					"max": MkUndefined(),
				}),
				"cost": MkMap(&VarMap{
					"min": MkInteger(0),
					"max": MkUndefined(),
				}),
			}),
		}))
	}
	return result
}

func (this *Bw) FetchCurrencies(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetExchangeConfigControllerWebsiteCurrencycontrollerGetCurrencyList"), params)
	_ = response
	currencies := this.SafeValue(response, MkString("datas"), MkArray(&VarArray{}))
	_ = currencies
	result := MkMap(&VarMap{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, currencies.Length)); OpIncr(&i) {
		currency := *(currencies).At(i)
		_ = currency
		id := this.SafeString(currency, MkString("currencyId"))
		_ = id
		code := this.SafeCurrencyCode(this.SafeStringUpper(currency, MkString("name")))
		_ = code
		state := this.SafeInteger(currency, MkString("state"))
		_ = state
		active := OpEq2(state, MkInteger(1))
		_ = active
		*(result).At(code) = MkMap(&VarMap{
			"id":        id,
			"code":      code,
			"info":      currency,
			"name":      code,
			"active":    active,
			"fee":       this.SafeNumber(currency, MkString("drawFee")),
			"precision": MkUndefined(),
			"limits": MkMap(&VarMap{
				"amount": MkMap(&VarMap{
					"min": this.SafeNumber(currency, MkString("limitAmount"), MkInteger(0)),
					"max": MkUndefined(),
				}),
				"withdraw": MkMap(&VarMap{
					"min": MkUndefined(),
					"max": this.SafeNumber(currency, MkString("onceDrawLimit")),
				}),
			}),
		})
	}
	return result
}

func (this *Bw) ParseTicker(goArgs ...*Variant) *Variant {
	ticker := GoGetArg(goArgs, 0, MkUndefined())
	_ = ticker
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	marketId := this.SafeString(ticker, MkInteger(0))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market)
	_ = symbol
	timestamp := this.Milliseconds()
	_ = timestamp
	close := this.SafeNumber(ticker, MkInteger(1))
	_ = close
	bid := this.SafeValue(ticker, MkString("bid"), MkMap(&VarMap{}))
	_ = bid
	ask := this.SafeValue(ticker, MkString("ask"), MkMap(&VarMap{}))
	_ = ask
	return MkMap(&VarMap{
		"symbol":        symbol,
		"timestamp":     timestamp,
		"datetime":      this.Iso8601(timestamp),
		"high":          this.SafeNumber(ticker, MkInteger(2)),
		"low":           this.SafeNumber(ticker, MkInteger(3)),
		"bid":           this.SafeNumber(ticker, MkInteger(7)),
		"bidVolume":     this.SafeNumber(bid, MkString("quantity")),
		"ask":           this.SafeNumber(ticker, MkInteger(8)),
		"askVolume":     this.SafeNumber(ask, MkString("quantity")),
		"vwap":          MkUndefined(),
		"open":          MkUndefined(),
		"close":         close,
		"last":          close,
		"previousClose": MkUndefined(),
		"change":        this.SafeNumber(ticker, MkInteger(5)),
		"percentage":    MkUndefined(),
		"average":       MkUndefined(),
		"baseVolume":    this.SafeNumber(ticker, MkInteger(4)),
		"quoteVolume":   this.SafeNumber(ticker, MkInteger(9)),
		"info":          ticker,
	})
}

func (this *Bw) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"marketId": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("publicGetApiDataV1Ticker"), this.Extend(request, params))
	_ = response
	ticker := this.SafeValue(response, MkString("datas"), MkArray(&VarArray{}))
	_ = ticker
	return this.ParseTicker(ticker, market)
}

func (this *Bw) FetchTickers(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("publicGetApiDataV1Tickers"), params)
	_ = response
	datas := this.SafeValue(response, MkString("datas"), MkArray(&VarArray{}))
	_ = datas
	result := MkMap(&VarMap{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, datas.Length)); OpIncr(&i) {
		ticker := this.ParseTicker(*(datas).At(i))
		_ = ticker
		symbol := *(ticker).At(MkString("symbol"))
		_ = symbol
		if IsTrue(OpOr(OpEq2(symbols, MkUndefined()), this.InArray(symbol, symbols))) {
			*(result).At(symbol) = OpCopy(ticker)
		}
	}
	return this.FilterByArray(result, MkString("symbol"), symbols)
}

func (this *Bw) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"marketId": *(market).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("dataSize")) = OpCopy(limit)
	}
	response := this.Call(MkString("publicGetApiDataV1Entrusts"), this.Extend(request, params))
	_ = response
	orderbook := this.SafeValue(response, MkString("datas"), MkArray(&VarArray{}))
	_ = orderbook
	timestamp := this.SafeTimestamp(orderbook, MkString("timestamp"))
	_ = timestamp
	return this.ParseOrderBook(orderbook, symbol, timestamp)
}

func (this *Bw) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.SafeTimestamp(trade, MkInteger(2))
	_ = timestamp
	priceString := this.SafeString(trade, MkInteger(5))
	_ = priceString
	amountString := this.SafeString(trade, MkInteger(6))
	_ = amountString
	price := this.ParseNumber(priceString)
	_ = price
	amount := this.ParseNumber(amountString)
	_ = amount
	cost := this.ParseNumber(Precise.StringMul(priceString, amountString))
	_ = cost
	marketId := this.SafeString(trade, MkInteger(1))
	_ = marketId
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(OpNotEq2(marketId, MkUndefined())) {
		if IsTrue(OpHasMember(marketId, *this.At(MkString("markets_by_id")))) {
			market = *(*this.At(MkString("markets_by_id"))).At(marketId)
		} else {
			marketName := this.SafeString(trade, MkInteger(3))
			_ = marketName
			baseId := MkUndefined()
			_ = baseId
			quoteId := MkUndefined()
			_ = quoteId
			ArrayUnpack(marketName.Split(MkString("_")), &baseId, &quoteId)
			base := this.SafeCurrencyCode(baseId)
			_ = base
			quote := this.SafeCurrencyCode(quoteId)
			_ = quote
			symbol = OpAdd(base, OpAdd(MkString("/"), quote))
		}
	}
	if IsTrue(OpAnd(OpEq2(symbol, MkUndefined()), OpNotEq2(market, MkUndefined()))) {
		symbol = *(market).At(MkString("symbol"))
	}
	sideString := this.SafeString(trade, MkInteger(4))
	_ = sideString
	side := OpTrinary(OpEq2(sideString, MkString("ask")), MkString("sell"), MkString("buy"))
	_ = side
	return MkMap(&VarMap{
		"id":           MkUndefined(),
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       symbol,
		"order":        MkUndefined(),
		"type":         MkString("limit"),
		"side":         side,
		"takerOrMaker": MkUndefined(),
		"price":        price,
		"amount":       amount,
		"cost":         cost,
		"fee":          MkUndefined(),
		"info":         trade,
	})
}

func (this *Bw) FetchTrades(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"marketId": *(market).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("dataSize")) = OpCopy(limit)
	}
	response := this.Call(MkString("publicGetApiDataV1Trades"), this.Extend(request, params))
	_ = response
	trades := this.SafeValue(response, MkString("datas"), MkArray(&VarArray{}))
	_ = trades
	return this.ParseTrades(trades, market, since, limit)
}

func (this *Bw) ParseOHLCV(goArgs ...*Variant) *Variant {
	ohlcv := GoGetArg(goArgs, 0, MkUndefined())
	_ = ohlcv
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	return MkArray(&VarArray{
		this.SafeTimestamp(ohlcv, MkInteger(3)),
		this.SafeNumber(ohlcv, MkInteger(4)),
		this.SafeNumber(ohlcv, MkInteger(5)),
		this.SafeNumber(ohlcv, MkInteger(6)),
		this.SafeNumber(ohlcv, MkInteger(7)),
		this.SafeNumber(ohlcv, MkInteger(8)),
	})
}

func (this *Bw) FetchOHLCV(goArgs ...*Variant) *Variant {
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
		"marketId": *(market).At(MkString("id")),
		"type":     *(*this.At(MkString("timeframes"))).At(timeframe),
		"dataSize": MkInteger(500),
	})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("dataSize")) = OpCopy(limit)
	}
	response := this.Call(MkString("publicGetApiDataV1Klines"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("datas"), MkArray(&VarArray{}))
	_ = data
	return this.ParseOHLCVs(data, market, timeframe, since, limit)
}

func (this *Bw) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privatePostExchangeFundControllerWebsiteFundcontrollerFindbypage"), params)
	_ = response
	data := this.SafeValue(response, MkString("datas"), MkMap(&VarMap{}))
	_ = data
	balances := this.SafeValue(data, MkString("list"), MkArray(&VarArray{}))
	_ = balances
	result := MkMap(&VarMap{"info": response})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, balances.Length)); OpIncr(&i) {
		balance := *(balances).At(i)
		_ = balance
		currencyId := this.SafeString(balance, MkString("currencyTypeId"))
		_ = currencyId
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		account := this.Account()
		_ = account
		*(account).At(MkString("free")) = this.SafeString(balance, MkString("amount"))
		*(account).At(MkString("used")) = this.SafeString(balance, MkString("freeze"))
		*(result).At(code) = OpCopy(account)
	}
	return this.ParseBalance(result)
}

func (this *Bw) CreateOrder(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpEq2(price, MkUndefined())) {
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), MkString(" allows limit orders only"))))
	}
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{
		"amount":    this.AmountToPrecision(symbol, amount),
		"price":     this.PriceToPrecision(symbol, price),
		"type":      OpTrinary(OpEq2(side, MkString("buy")), MkInteger(1), MkInteger(0)),
		"rangeType": MkInteger(0),
		"marketId":  *(market).At(MkString("id")),
	})
	_ = request
	response := this.Call(MkString("privatePostExchangeEntrustControllerWebsiteEntrustControllerAddEntrust"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("datas"))
	_ = data
	id := this.SafeString(data, MkString("entrustId"))
	_ = id
	return MkMap(&VarMap{
		"id":                 id,
		"info":               response,
		"timestamp":          MkUndefined(),
		"datetime":           MkUndefined(),
		"lastTradeTimestamp": MkUndefined(),
		"symbol":             symbol,
		"type":               type_,
		"side":               side,
		"price":              price,
		"amount":             amount,
		"cost":               MkUndefined(),
		"average":            MkUndefined(),
		"filled":             MkUndefined(),
		"remaining":          MkUndefined(),
		"status":             MkString("open"),
		"fee":                MkUndefined(),
		"trades":             MkUndefined(),
		"clientOrderId":      MkUndefined(),
	})
}

func (this *Bw) ParseOrderStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"-3": MkString("canceled"),
		"-2": MkString("canceled"),
		"-1": MkString("canceled"),
		"0":  MkString("open"),
		"1":  MkString("canceled"),
		"2":  MkString("closed"),
		"3":  MkString("open"),
		"4":  MkString("canceled"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Bw) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	marketId := this.SafeString(order, MkString("marketId"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market)
	_ = symbol
	timestamp := this.SafeInteger(order, MkString("createTime"))
	_ = timestamp
	side := this.SafeString(order, MkString("type"))
	_ = side
	if IsTrue(OpEq2(side, MkString("0"))) {
		side = MkString("sell")
	} else {
		if IsTrue(OpEq2(side, MkString("1"))) {
			side = MkString("buy")
		}
	}
	amount := this.SafeNumber(order, MkString("amount"))
	_ = amount
	price := this.SafeNumber(order, MkString("price"))
	_ = price
	filled := this.SafeNumber(order, MkString("completeAmount"))
	_ = filled
	remaining := this.SafeNumber2(order, MkString("availabelAmount"), MkString("availableAmount"))
	_ = remaining
	cost := this.SafeNumber(order, MkString("totalMoney"))
	_ = cost
	status := this.ParseOrderStatus(this.SafeString(order, MkString("status")))
	_ = status
	return this.SafeOrder(MkMap(&VarMap{
		"info":               order,
		"id":                 this.SafeString(order, MkString("entrustId")),
		"clientOrderId":      MkUndefined(),
		"timestamp":          timestamp,
		"datetime":           this.Iso8601(timestamp),
		"lastTradeTimestamp": MkUndefined(),
		"symbol":             symbol,
		"type":               MkString("limit"),
		"timeInForce":        MkUndefined(),
		"postOnly":           MkUndefined(),
		"side":               side,
		"price":              price,
		"stopPrice":          MkUndefined(),
		"amount":             amount,
		"cost":               cost,
		"average":            MkUndefined(),
		"filled":             filled,
		"remaining":          remaining,
		"status":             status,
		"fee":                MkUndefined(),
		"trades":             MkUndefined(),
	}))
}

func (this *Bw) FetchOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchOrder() requires a symbol argument"))))
	}
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{
		"marketId":  *(market).At(MkString("id")),
		"entrustId": id,
	})
	_ = request
	response := this.Call(MkString("privateGetExchangeEntrustControllerWebsiteEntrustControllerGetEntrustById"), this.Extend(request, params))
	_ = response
	order := this.SafeValue(response, MkString("datas"), MkMap(&VarMap{}))
	_ = order
	return this.ParseOrder(order, market)
}

func (this *Bw) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" cancelOrder() requires a symbol argument"))))
	}
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{
		"marketId":  *(market).At(MkString("id")),
		"entrustId": id,
	})
	_ = request
	response := this.Call(MkString("privatePostExchangeEntrustControllerWebsiteEntrustControllerCancelEntrust"), this.Extend(request, params))
	_ = response
	return MkMap(&VarMap{
		"info": response,
		"id":   id,
	})
}

func (this *Bw) FetchOpenOrders(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"marketId": *(market).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("pageSize")) = OpCopy(limit)
	}
	response := this.Call(MkString("privateGetExchangeEntrustControllerWebsiteEntrustControllerGetUserEntrustRecordFromCacheWithPage"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("datas"), MkMap(&VarMap{}))
	_ = data
	orders := this.SafeValue(data, MkString("entrustList"), MkArray(&VarArray{}))
	_ = orders
	return this.ParseOrders(orders, market, since, limit)
}

func (this *Bw) FetchClosedOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchClosedOrders() requires a symbol argument"))))
	}
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"marketId": *(market).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("pageSize")) = OpCopy(limit)
	}
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("startDateTime")) = OpCopy(since)
	}
	response := this.Call(MkString("privateGetExchangeEntrustControllerWebsiteEntrustControllerGetUserEntrustList"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("datas"), MkMap(&VarMap{}))
	_ = data
	orders := this.SafeValue(data, MkString("entrustList"), MkArray(&VarArray{}))
	_ = orders
	return this.ParseOrders(orders, market, since, limit)
}

func (this *Bw) FetchOrders(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"marketId": *(market).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("startDateTime")) = OpCopy(since)
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("pageSize")) = OpCopy(limit)
	}
	response := this.Call(MkString("privateGetExchangeEntrustControllerWebsiteEntrustControllerGetUserEntrustList"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("datas"), MkMap(&VarMap{}))
	_ = data
	orders := this.SafeValue(data, MkString("entrustList"), MkArray(&VarArray{}))
	_ = orders
	return this.ParseOrders(orders, market, since, limit)
}

func (this *Bw) Sign(goArgs ...*Variant) *Variant {
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
	url := OpAdd(this.ImplodeHostname(*(*this.At(MkString("urls"))).At(MkString("api"))), OpAdd(MkString("/"), path))
	_ = url
	if IsTrue(OpEq2(method, MkString("GET"))) {
		if IsTrue(*(GoGetKeys(params)).At(MkString("length"))) {
			OpAddAssign(&url, OpAdd(MkString("?"), this.Urlencode(params)))
		}
	} else {
		body = this.Json(params)
	}
	if IsTrue(OpEq2(api, MkString("private"))) {
		ms := (this.Milliseconds()).Call(MkString("toString"))
		_ = ms
		content := MkString("")
		_ = content
		if IsTrue(OpEq2(method, MkString("GET"))) {
			sortedParams := this.Keysort(params)
			_ = sortedParams
			keys := GoGetKeys(sortedParams)
			_ = keys
			for i := MkInteger(0); IsTrue(OpLw(i, keys.Length)); OpIncr(&i) {
				key := *(keys).At(i)
				_ = key
				OpAddAssign(&content, OpAdd(key, (*(sortedParams).At(key)).Call(MkString("toString"))))
			}
		} else {
			content = OpCopy(body)
		}
		signature := OpAdd(*this.At(MkString("apiKey")), OpAdd(ms, OpAdd(content, *this.At(MkString("secret")))))
		_ = signature
		hash := this.Hash(this.Encode(signature), MkString("md5"))
		_ = hash
		if IsTrue(OpNot(headers)) {
			headers = MkMap(&VarMap{})
		}
		*(headers).At(MkString("Apiid")) = OpCopy(*this.At(MkString("apiKey")))
		*(headers).At(MkString("Timestamp")) = OpCopy(ms)
		*(headers).At(MkString("Sign")) = OpCopy(hash)
	}
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Bw) FetchDepositAddress(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{"currencyTypeName": *(currency).At(MkString("name"))})
	_ = request
	response := this.Call(MkString("privatePostExchangeFundControllerWebsiteFundcontrollerGetPayinAddress"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("datas"), MkMap(&VarMap{}))
	_ = data
	address := this.SafeString(data, MkString("address"))
	_ = address
	tag := this.SafeString(data, MkString("memo"))
	_ = tag
	this.CheckAddress(address)
	return MkMap(&VarMap{
		"currency": code,
		"address":  this.CheckAddress(address),
		"tag":      tag,
		"info":     response,
	})
}

func (this *Bw) ParseTransactionStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"-1": MkString("canceled"),
		"0":  MkString("pending"),
		"1":  MkString("ok"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Bw) ParseTransaction(goArgs ...*Variant) *Variant {
	transaction := GoGetArg(goArgs, 0, MkUndefined())
	_ = transaction
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	id := this.SafeString(transaction, MkString("depositId"), MkString("withdrawalId"))
	_ = id
	address := this.SafeString2(transaction, MkString("depositAddress"), MkString("withdrawalAddress"))
	_ = address
	currencyId := this.SafeString2(transaction, MkString("currencyId"), MkString("currencyTypeId"))
	_ = currencyId
	code := OpCopy(MkUndefined())
	_ = code
	if IsTrue(OpHasMember(currencyId, *this.At(MkString("currencies_by_id")))) {
		currency = *(*this.At(MkString("currencies_by_id"))).At(currencyId)
	}
	if IsTrue(OpAnd(OpEq2(code, MkUndefined()), OpNotEq2(currency, MkUndefined()))) {
		code = *(currency).At(MkString("code"))
	}
	type_ := OpTrinary(OpHasMember(MkString("depositId"), transaction), MkString("deposit"), MkString("withdrawal"))
	_ = type_
	amount := this.SafeNumber2(transaction, MkString("actuallyAmount"), MkString("amount"))
	_ = amount
	status := this.ParseTransactionStatus(this.SafeString2(transaction, MkString("verifyStatus"), MkString("state")))
	_ = status
	timestamp := this.SafeInteger(transaction, MkString("createTime"))
	_ = timestamp
	txid := this.SafeString(transaction, MkString("txId"))
	_ = txid
	fee := OpCopy(MkUndefined())
	_ = fee
	feeCost := this.SafeNumber(transaction, MkString("fees"))
	_ = feeCost
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		fee = MkMap(&VarMap{
			"cost":     feeCost,
			"currency": code,
		})
	}
	return MkMap(&VarMap{
		"info":        transaction,
		"id":          id,
		"txid":        txid,
		"timestamp":   timestamp,
		"datetime":    this.Iso8601(timestamp),
		"addressFrom": MkUndefined(),
		"address":     address,
		"addressTo":   MkUndefined(),
		"tagFrom":     MkUndefined(),
		"tag":         MkUndefined(),
		"tagTo":       MkUndefined(),
		"type":        type_,
		"amount":      amount,
		"currency":    code,
		"status":      status,
		"updated":     MkUndefined(),
		"fee":         fee,
	})
}

func (this *Bw) FetchDeposits(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(code, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchDeposits() requires a currency code argument"))))
	}
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{"currencyTypeName": *(currency).At(MkString("name"))})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("pageSize")) = OpCopy(limit)
	}
	response := this.Call(MkString("privatePostExchangeFundControllerWebsiteFundcontrollerGetPayinCoinRecord"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("datas"), MkMap(&VarMap{}))
	_ = data
	deposits := this.SafeValue(data, MkString("list"), MkArray(&VarArray{}))
	_ = deposits
	return this.ParseTransactions(deposits, code, since, limit)
}

func (this *Bw) FetchWithdrawals(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(code, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchWithdrawals() requires a currency code argument"))))
	}
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{"currencyId": *(currency).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("pageSize")) = OpCopy(limit)
	}
	response := this.Call(MkString("privateGetExchangeFundControllerWebsiteFundwebsitecontrollerGetpayoutcoinrecord"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("datas"), MkMap(&VarMap{}))
	_ = data
	withdrawals := this.SafeValue(data, MkString("list"), MkArray(&VarArray{}))
	_ = withdrawals
	return this.ParseTransactions(withdrawals, code, since, limit)
}

func (this *Bw) HandleErrors(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpNot(response)) {
		return MkUndefined()
	}
	resMsg := this.SafeValue(response, MkString("resMsg"))
	_ = resMsg
	errorCode := this.SafeString(resMsg, MkString("code"))
	_ = errorCode
	if IsTrue(OpNotEq2(errorCode, MkString("1"))) {
		feedback := OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), this.Json(response)))
		_ = feedback
		this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("exact")), errorCode, feedback)
		panic(NewExchangeError(feedback))
	}
	return MkUndefined()
}
