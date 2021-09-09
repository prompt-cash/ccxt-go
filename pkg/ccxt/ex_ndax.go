package ccxt

import ()

type Ndax struct {
	*ExchangeBase
}

var _ Exchange = (*Ndax)(nil)

func init() {
	exchange := &Ndax{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Ndax) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("ndax"),
		"name": MkString("NDAX"),
		"countries": MkArray(&VarArray{
			MkString("US"),
		}),
		"rateLimit": MkInteger(1000),
		"pro":       MkBool(true),
		"has": MkMap(&VarMap{
			"cancelAllOrders":      MkBool(true),
			"cancelOrder":          MkBool(true),
			"createDepositAddress": MkBool(true),
			"createOrder":          MkBool(true),
			"editOrder":            MkBool(true),
			"fetchAccounts":        MkBool(true),
			"fetchBalance":         MkBool(true),
			"fetchCurrencies":      MkBool(true),
			"fetchDepositAddress":  MkBool(true),
			"fetchDeposits":        MkBool(true),
			"fetchLedger":          MkBool(true),
			"fetchMarkets":         MkBool(true),
			"fetchMyTrades":        MkBool(true),
			"fetchOHLCV":           MkBool(true),
			"fetchOpenOrders":      MkBool(true),
			"fetchOrders":          MkBool(true),
			"fetchOrder":           MkBool(true),
			"fetchOrderBook":       MkBool(true),
			"fetchOrderTrades":     MkBool(true),
			"fetchTicker":          MkBool(true),
			"fetchTrades":          MkBool(true),
			"fetchWithdrawals":     MkBool(true),
			"signIn":               MkBool(true),
		}),
		"timeframes": MkMap(&VarMap{
			"1m":  MkString("60"),
			"5m":  MkString("300"),
			"15m": MkString("900"),
			"30m": MkString("1800"),
			"1h":  MkString("3600"),
			"2h":  MkString("7200"),
			"4h":  MkString("14400"),
			"6h":  MkString("21600"),
			"12h": MkString("43200"),
			"1d":  MkString("86400"),
			"1w":  MkString("604800"),
			"1M":  MkString("2419200"),
			"4M":  MkString("9676800"),
		}),
		"urls": MkMap(&VarMap{
			"logo": MkString("https://user-images.githubusercontent.com/1294454/108623144-67a3ef00-744e-11eb-8140-75c6b851e945.jpg"),
			"test": MkMap(&VarMap{
				"public":  MkString("https://ndaxmarginstaging.cdnhop.net:8443/AP"),
				"private": MkString("https://ndaxmarginstaging.cdnhop.net:8443/AP"),
			}),
			"api": MkMap(&VarMap{
				"public":  MkString("https://api.ndax.io:8443/AP"),
				"private": MkString("https://api.ndax.io:8443/AP"),
			}),
			"www": MkString("https://ndax.io"),
			"doc": MkArray(&VarArray{
				MkString("https://apidoc.ndax.io/"),
			}),
			"fees":     MkString("https://ndax.io/fees"),
			"referral": MkString("https://one.ndax.io/bfQiSL"),
		}),
		"api": MkMap(&VarMap{
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("Activate2FA"),
				MkString("Authenticate2FA"),
				MkString("AuthenticateUser"),
				MkString("GetL2Snapshot"),
				MkString("GetLevel1"),
				MkString("GetValidate2FARequiredEndpoints"),
				MkString("LogOut"),
				MkString("GetTickerHistory"),
				MkString("GetProduct"),
				MkString("GetProducts"),
				MkString("GetInstrument"),
				MkString("GetInstruments"),
				MkString("Ping"),
				MkString("trades"),
				MkString("GetLastTrades"),
				MkString("SubscribeLevel1"),
				MkString("SubscribeLevel2"),
				MkString("SubscribeTicker"),
				MkString("SubscribeTrades"),
				MkString("SubscribeBlockTrades"),
				MkString("UnsubscribeBlockTrades"),
				MkString("UnsubscribeLevel1"),
				MkString("UnsubscribeLevel2"),
				MkString("UnsubscribeTicker"),
				MkString("UnsubscribeTrades"),
				MkString("Authenticate"),
			})}),
			"private": MkMap(&VarMap{
				"get": MkArray(&VarArray{
					MkString("GetUserAccountInfos"),
					MkString("GetUserAccounts"),
					MkString("GetUserAffiliateCount"),
					MkString("GetUserAffiliateTag"),
					MkString("GetUserConfig"),
					MkString("GetAllUnredactedUserConfigsForUser"),
					MkString("GetUnredactedUserConfigByKey"),
					MkString("GetUserDevices"),
					MkString("GetUserReportTickets"),
					MkString("GetUserReportWriterResultRecords"),
					MkString("GetAccountInfo"),
					MkString("GetAccountPositions"),
					MkString("GetAllAccountConfigs"),
					MkString("GetTreasuryProductsForAccount"),
					MkString("GetAccountTrades"),
					MkString("GetAccountTransactions"),
					MkString("GetOpenTradeReports"),
					MkString("GetAllOpenTradeReports"),
					MkString("GetTradesHistory"),
					MkString("GetOpenOrders"),
					MkString("GetOpenQuotes"),
					MkString("GetOrderFee"),
					MkString("GetOrderHistory"),
					MkString("GetOrdersHistory"),
					MkString("GetOrderStatus"),
					MkString("GetOmsFeeTiers"),
					MkString("GetAccountDepositTransactions"),
					MkString("GetAccountWithdrawTransactions"),
					MkString("GetAllDepositRequestInfoTemplates"),
					MkString("GetDepositInfo"),
					MkString("GetDepositRequestInfoTemplate"),
					MkString("GetDeposits"),
					MkString("GetDepositTicket"),
					MkString("GetDepositTickets"),
					MkString("GetOMSWithdrawFees"),
					MkString("GetWithdrawFee"),
					MkString("GetWithdraws"),
					MkString("GetWithdrawTemplate"),
					MkString("GetWithdrawTemplateTypes"),
					MkString("GetWithdrawTicket"),
					MkString("GetWithdrawTickets"),
				}),
				"post": MkArray(&VarArray{
					MkString("AddUserAffiliateTag"),
					MkString("CancelUserReport"),
					MkString("RegisterNewDevice"),
					MkString("SubscribeAccountEvents"),
					MkString("UpdateUserAffiliateTag"),
					MkString("GenerateTradeActivityReport"),
					MkString("GenerateTransactionActivityReport"),
					MkString("GenerateTreasuryActivityReport"),
					MkString("ScheduleTradeActivityReport"),
					MkString("ScheduleTransactionActivityReport"),
					MkString("ScheduleTreasuryActivityReport"),
					MkString("CancelAllOrders"),
					MkString("CancelOrder"),
					MkString("CancelQuote"),
					MkString("CancelReplaceOrder"),
					MkString("CreateQuote"),
					MkString("ModifyOrder"),
					MkString("SendOrder"),
					MkString("SubmitBlockTrade"),
					MkString("UpdateQuote"),
					MkString("CancelWithdraw"),
					MkString("CreateDepositTicket"),
					MkString("CreateWithdrawTicket"),
					MkString("SubmitDepositTicketComment"),
					MkString("SubmitWithdrawTicketComment"),
					MkString("GetOrderHistoryByOrderId"),
				}),
			}),
		}),
		"fees": MkMap(&VarMap{"trading": MkMap(&VarMap{
			"tierBased":  MkBool(false),
			"percentage": MkBool(true),
			"maker":      OpDiv(MkNumber(0.2), MkInteger(100)),
			"taker":      OpDiv(MkNumber(0.25), MkInteger(100)),
		})}),
		"requiredCredentials": MkMap(&VarMap{
			"apiKey": MkBool(true),
			"secret": MkBool(true),
			"uid":    MkBool(true),
		}),
		"precisionMode": TICK_SIZE,
		"exceptions": MkMap(&VarMap{
			"exact": MkMap(&VarMap{
				"Not_Enough_Funds":   InsufficientFunds,
				"Server Error":       ExchangeError,
				"Resource Not Found": OrderNotFound,
			}),
			"broad": MkMap(&VarMap{
				"Invalid InstrumentId":                                  BadSymbol,
				"This endpoint requires 2FACode along with the payload": AuthenticationError,
			}),
		}),
		"options": MkMap(&VarMap{
			"omsId": MkInteger(1),
			"orderTypes": MkMap(&VarMap{
				"Market":             MkInteger(1),
				"Limit":              MkInteger(2),
				"StopMarket":         MkInteger(3),
				"StopLimit":          MkInteger(4),
				"TrailingStopMarket": MkInteger(5),
				"TrailingStopLimit":  MkInteger(6),
				"BlockTrade":         MkInteger(7),
			}),
		}),
	}))
}

func (this *Ndax) SignIn(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.CheckRequiredCredentials()
	if IsTrue(OpOr(OpEq2(*this.At(MkString("login")), MkUndefined()), OpOr(OpEq2(*this.At(MkString("password")), MkUndefined()), OpEq2(*this.At(MkString("twofa")), MkUndefined())))) {
		panic(NewAuthenticationError(OpAdd(*this.At(MkString("id")), MkString(" signIn() requires exchange.login, exchange.password and exchange.twofa credentials"))))
	}
	request := MkMap(&VarMap{"grant_type": MkString("client_credentials")})
	_ = request
	response := this.Call(MkString("publicGetAuthenticate"), this.Extend(request, params))
	_ = response
	sessionToken := this.SafeString(response, MkString("SessionToken"))
	_ = sessionToken
	if IsTrue(OpNotEq2(sessionToken, MkUndefined())) {
		*(*this.At(MkString("options"))).At(MkString("sessionToken")) = OpCopy(sessionToken)
		return response
	}
	pending2faToken := this.SafeString(response, MkString("Pending2FaToken"))
	_ = pending2faToken
	if IsTrue(OpNotEq2(pending2faToken, MkUndefined())) {
		*(*this.At(MkString("options"))).At(MkString("pending2faToken")) = OpCopy(pending2faToken)
		request = MkMap(&VarMap{"Code": this.Oath()})
		response := this.Call(MkString("publicGetAuthenticate2FA"), this.Extend(request, params))
		_ = response
		sessionToken = this.SafeString(response, MkString("SessionToken"))
		*(*this.At(MkString("options"))).At(MkString("sessionToken")) = OpCopy(sessionToken)
		return response
	}
	return response
}

func (this *Ndax) FetchCurrencies(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	omsId := this.SafeInteger(*this.At(MkString("options")), MkString("omsId"), MkInteger(1))
	_ = omsId
	request := MkMap(&VarMap{"omsId": omsId})
	_ = request
	response := this.Call(MkString("publicGetGetProducts"), this.Extend(request, params))
	_ = response
	result := MkMap(&VarMap{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
		currency := *(response).At(i)
		_ = currency
		id := this.SafeString(currency, MkString("ProductId"))
		_ = id
		name := this.SafeString(currency, MkString("ProductFullName"))
		_ = name
		type_ := this.SafeString(currency, MkString("ProductType"))
		_ = type_
		code := this.SafeCurrencyCode(this.SafeString(currency, MkString("Product")))
		_ = code
		precision := this.SafeNumber(currency, MkString("TickSize"))
		_ = precision
		isDisabled := this.SafeValue(currency, MkString("IsDisabled"))
		_ = isDisabled
		active := OpNot(isDisabled)
		_ = active
		*(result).At(code) = MkMap(&VarMap{
			"id":        id,
			"name":      name,
			"code":      code,
			"type":      type_,
			"precision": precision,
			"info":      currency,
			"active":    active,
			"fee":       MkUndefined(),
			"limits":    *this.At(MkString("limits")),
		})
	}
	return result
}

func (this *Ndax) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	omsId := this.SafeInteger(*this.At(MkString("options")), MkString("omsId"), MkInteger(1))
	_ = omsId
	request := MkMap(&VarMap{"omsId": omsId})
	_ = request
	response := this.Call(MkString("publicGetGetInstruments"), this.Extend(request, params))
	_ = response
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
		market := *(response).At(i)
		_ = market
		id := this.SafeString(market, MkString("InstrumentId"))
		_ = id
		baseId := this.SafeString(market, MkString("Product1"))
		_ = baseId
		quoteId := this.SafeString(market, MkString("Product2"))
		_ = quoteId
		base := this.SafeCurrencyCode(this.SafeString(market, MkString("Product1Symbol")))
		_ = base
		quote := this.SafeCurrencyCode(this.SafeString(market, MkString("Product2Symbol")))
		_ = quote
		symbol := OpAdd(base, OpAdd(MkString("/"), quote))
		_ = symbol
		precision := MkMap(&VarMap{
			"amount": this.SafeNumber(market, MkString("QuantityIncrement")),
			"price":  this.SafeNumber(market, MkString("PriceIncrement")),
		})
		_ = precision
		sessionStatus := this.SafeString(market, MkString("SessionStatus"))
		_ = sessionStatus
		isDisable := this.SafeValue(market, MkString("IsDisable"))
		_ = isDisable
		sessionRunning := OpEq2(sessionStatus, MkString("Running"))
		_ = sessionRunning
		active := OpTrinary(OpAnd(sessionRunning, OpNot(isDisable)), MkBool(true), MkBool(false))
		_ = active
		result.Push(MkMap(&VarMap{
			"id":        id,
			"symbol":    symbol,
			"base":      base,
			"quote":     quote,
			"baseId":    baseId,
			"quoteId":   quoteId,
			"info":      market,
			"active":    active,
			"precision": precision,
			"limits": MkMap(&VarMap{
				"amount": MkMap(&VarMap{
					"min": this.SafeNumber(market, MkString("MinimumQuantity")),
					"max": MkUndefined(),
				}),
				"price": MkMap(&VarMap{
					"min": this.SafeNumber(market, MkString("MinimumPrice")),
					"max": MkUndefined(),
				}),
				"cost": MkMap(&VarMap{
					"min": MkUndefined(),
					"max": MkUndefined(),
				}),
			}),
		}))
	}
	return result
}

func (this *Ndax) ParseOrderBook(goArgs ...*Variant) *Variant {
	orderbook := GoGetArg(goArgs, 0, MkUndefined())
	_ = orderbook
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	timestamp := GoGetArg(goArgs, 2, MkUndefined())
	_ = timestamp
	bidsKey := GoGetArg(goArgs, 3, MkString("bids"))
	_ = bidsKey
	asksKey := GoGetArg(goArgs, 4, MkString("asks"))
	_ = asksKey
	priceKey := GoGetArg(goArgs, 5, MkInteger(6))
	_ = priceKey
	amountKey := GoGetArg(goArgs, 6, MkInteger(8))
	_ = amountKey
	nonce := OpCopy(MkUndefined())
	_ = nonce
	result := MkMap(&VarMap{
		"symbol":    symbol,
		"bids":      MkArray(&VarArray{}),
		"asks":      MkArray(&VarArray{}),
		"timestamp": MkUndefined(),
		"datetime":  MkUndefined(),
		"nonce":     MkUndefined(),
	})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, orderbook.Length)); OpIncr(&i) {
		level := *(orderbook).At(i)
		_ = level
		if IsTrue(OpEq2(timestamp, MkUndefined())) {
			timestamp = this.SafeInteger(level, MkInteger(2))
		} else {
			newTimestamp := this.SafeInteger(level, MkInteger(2))
			_ = newTimestamp
			timestamp = Math.Max(timestamp, newTimestamp)
		}
		if IsTrue(OpEq2(nonce, MkUndefined())) {
			nonce = this.SafeInteger(level, MkInteger(0))
		} else {
			newNonce := this.SafeInteger(level, MkInteger(0))
			_ = newNonce
			nonce = Math.Max(nonce, newNonce)
		}
		bidask := this.ParseBidAsk(level, priceKey, amountKey)
		_ = bidask
		levelSide := this.SafeInteger(level, MkInteger(9))
		_ = levelSide
		side := OpTrinary(levelSide, asksKey, bidsKey)
		_ = side
		(*(result).At(side)).Call(MkString("push"), bidask)
	}
	*(result).At(MkString("bids")) = this.SortBy(*(result).At(MkString("bids")), MkInteger(0), MkBool(true))
	*(result).At(MkString("asks")) = this.SortBy(*(result).At(MkString("asks")), MkInteger(0))
	*(result).At(MkString("timestamp")) = OpCopy(timestamp)
	*(result).At(MkString("datetime")) = this.Iso8601(timestamp)
	*(result).At(MkString("nonce")) = OpCopy(nonce)
	return result
}

func (this *Ndax) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	omsId := this.SafeInteger(*this.At(MkString("options")), MkString("omsId"), MkInteger(1))
	_ = omsId
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	limit = OpTrinary(OpEq2(limit, MkUndefined()), MkInteger(100), limit)
	request := MkMap(&VarMap{
		"omsId":        omsId,
		"InstrumentId": *(market).At(MkString("id")),
		"Depth":        limit,
	})
	_ = request
	response := this.Call(MkString("publicGetGetL2Snapshot"), this.Extend(request, params))
	_ = response
	return this.ParseOrderBook(response, symbol)
}

func (this *Ndax) ParseTicker(goArgs ...*Variant) *Variant {
	ticker := GoGetArg(goArgs, 0, MkUndefined())
	_ = ticker
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.SafeInteger(ticker, MkString("TimeStamp"))
	_ = timestamp
	marketId := this.SafeString(ticker, MkString("InstrumentId"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market)
	_ = symbol
	last := this.SafeNumber(ticker, MkString("LastTradedPx"))
	_ = last
	percentage := this.SafeNumber(ticker, MkString("Rolling24HrPxChangePercent"))
	_ = percentage
	change := this.SafeNumber(ticker, MkString("Rolling24HrPxChange"))
	_ = change
	open := this.SafeNumber(ticker, MkString("SessionOpen"))
	_ = open
	average := OpCopy(MkUndefined())
	_ = average
	if IsTrue(OpAnd(OpNotEq2(last, MkUndefined()), OpNotEq2(change, MkUndefined()))) {
		average = OpDiv(this.Sum(last, open), MkInteger(2))
	}
	baseVolume := this.SafeNumber(ticker, MkString("Rolling24HrVolume"))
	_ = baseVolume
	quoteVolume := this.SafeNumber(ticker, MkString("Rolling24HrNotional"))
	_ = quoteVolume
	vwap := this.Vwap(baseVolume, quoteVolume)
	_ = vwap
	return MkMap(&VarMap{
		"symbol":        symbol,
		"timestamp":     timestamp,
		"datetime":      this.Iso8601(timestamp),
		"high":          this.SafeNumber(ticker, MkString("SessionHigh")),
		"low":           this.SafeNumber(ticker, MkString("SessionLow")),
		"bid":           this.SafeNumber(ticker, MkString("BestBid")),
		"bidVolume":     MkUndefined(),
		"ask":           this.SafeNumber(ticker, MkString("BestOffer")),
		"askVolume":     MkUndefined(),
		"vwap":          vwap,
		"open":          open,
		"close":         last,
		"last":          last,
		"previousClose": MkUndefined(),
		"change":        change,
		"percentage":    percentage,
		"average":       average,
		"baseVolume":    baseVolume,
		"quoteVolume":   quoteVolume,
		"info":          ticker,
	})
}

func (this *Ndax) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	omsId := this.SafeInteger(*this.At(MkString("options")), MkString("omsId"), MkInteger(1))
	_ = omsId
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{
		"omsId":        omsId,
		"InstrumentId": *(market).At(MkString("id")),
	})
	_ = request
	response := this.Call(MkString("publicGetGetLevel1"), this.Extend(request, params))
	_ = response
	return this.ParseTicker(response, market)
}

func (this *Ndax) ParseOHLCV(goArgs ...*Variant) *Variant {
	ohlcv := GoGetArg(goArgs, 0, MkUndefined())
	_ = ohlcv
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	return MkArray(&VarArray{
		this.SafeInteger(ohlcv, MkInteger(0)),
		this.SafeNumber(ohlcv, MkInteger(3)),
		this.SafeNumber(ohlcv, MkInteger(1)),
		this.SafeNumber(ohlcv, MkInteger(2)),
		this.SafeNumber(ohlcv, MkInteger(4)),
		this.SafeNumber(ohlcv, MkInteger(5)),
	})
}

func (this *Ndax) FetchOHLCV(goArgs ...*Variant) *Variant {
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
	omsId := this.SafeInteger(*this.At(MkString("options")), MkString("omsId"), MkInteger(1))
	_ = omsId
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{
		"omsId":        omsId,
		"InstrumentId": *(market).At(MkString("id")),
		"Interval":     *(*this.At(MkString("timeframes"))).At(timeframe),
	})
	_ = request
	duration := this.ParseTimeframe(timeframe)
	_ = duration
	now := this.Milliseconds()
	_ = now
	if IsTrue(OpEq2(since, MkUndefined())) {
		if IsTrue(OpNotEq2(limit, MkUndefined())) {
			*(request).At(MkString("FromDate")) = this.Ymdhms(OpSub(now, OpMulti(duration, OpMulti(limit, MkInteger(1000)))))
			*(request).At(MkString("ToDate")) = this.Ymdhms(now)
		}
	} else {
		*(request).At(MkString("FromDate")) = this.Ymdhms(since)
		if IsTrue(OpEq2(limit, MkUndefined())) {
			*(request).At(MkString("ToDate")) = this.Ymdhms(now)
		} else {
			*(request).At(MkString("ToDate")) = this.Ymdhms(this.Sum(since, OpMulti(duration, OpMulti(limit, MkInteger(1000)))))
		}
	}
	response := this.Call(MkString("publicGetGetTickerHistory"), this.Extend(request, params))
	_ = response
	return this.ParseOHLCVs(response, market, timeframe, since, limit)
}

func (this *Ndax) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	priceString := OpCopy(MkUndefined())
	_ = priceString
	amountString := OpCopy(MkUndefined())
	_ = amountString
	cost := OpCopy(MkUndefined())
	_ = cost
	timestamp := OpCopy(MkUndefined())
	_ = timestamp
	id := OpCopy(MkUndefined())
	_ = id
	marketId := OpCopy(MkUndefined())
	_ = marketId
	side := OpCopy(MkUndefined())
	_ = side
	orderId := OpCopy(MkUndefined())
	_ = orderId
	takerOrMaker := OpCopy(MkUndefined())
	_ = takerOrMaker
	fee := OpCopy(MkUndefined())
	_ = fee
	type_ := OpCopy(MkUndefined())
	_ = type_
	if IsTrue(GoIsArray(trade)) {
		priceString = this.SafeString(trade, MkInteger(3))
		amountString = this.SafeString(trade, MkInteger(2))
		timestamp = this.SafeInteger(trade, MkInteger(6))
		id = this.SafeString(trade, MkInteger(0))
		marketId = this.SafeString(trade, MkInteger(1))
		takerSide := this.SafeValue(trade, MkInteger(8))
		_ = takerSide
		side = OpTrinary(takerSide, MkString("sell"), MkString("buy"))
		orderId = this.SafeString(trade, MkInteger(4))
	} else {
		timestamp = this.SafeInteger2(trade, MkString("TradeTimeMS"), MkString("ReceiveTime"))
		id = this.SafeString(trade, MkString("TradeId"))
		orderId = this.SafeString2(trade, MkString("OrderId"), MkString("OrigOrderId"))
		marketId = this.SafeString2(trade, MkString("InstrumentId"), MkString("Instrument"))
		priceString = this.SafeString(trade, MkString("Price"))
		amountString = this.SafeString(trade, MkString("Quantity"))
		cost = this.SafeNumber2(trade, MkString("Value"), MkString("GrossValueExecuted"))
		takerOrMaker = this.SafeStringLower(trade, MkString("MakerTaker"))
		side = this.SafeStringLower(trade, MkString("Side"))
		type_ = this.SafeStringLower(trade, MkString("OrderType"))
		feeCost := this.SafeNumber(trade, MkString("Fee"))
		_ = feeCost
		if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
			feeCurrencyId := this.SafeString(trade, MkString("FeeProductId"))
			_ = feeCurrencyId
			feeCurrencyCode := this.SafeCurrencyCode(feeCurrencyId)
			_ = feeCurrencyCode
			fee = MkMap(&VarMap{
				"cost":     feeCost,
				"currency": feeCurrencyCode,
			})
		}
	}
	price := this.ParseNumber(priceString)
	_ = price
	amount := this.ParseNumber(amountString)
	_ = amount
	if IsTrue(OpEq2(cost, MkUndefined())) {
		cost = this.ParseNumber(Precise.StringMul(priceString, amountString))
	}
	symbol := this.SafeSymbol(marketId, market)
	_ = symbol
	return MkMap(&VarMap{
		"info":         trade,
		"id":           id,
		"symbol":       symbol,
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"order":        orderId,
		"type":         type_,
		"side":         side,
		"takerOrMaker": takerOrMaker,
		"price":        price,
		"amount":       amount,
		"cost":         cost,
		"fee":          fee,
	})
}

func (this *Ndax) FetchTrades(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	omsId := this.SafeInteger(*this.At(MkString("options")), MkString("omsId"), MkInteger(1))
	_ = omsId
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{
		"omsId":        omsId,
		"InstrumentId": *(market).At(MkString("id")),
	})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("Count")) = OpCopy(limit)
	}
	response := this.Call(MkString("publicGetGetLastTrades"), this.Extend(request, params))
	_ = response
	return this.ParseTrades(response, market, since, limit)
}

func (this *Ndax) FetchAccounts(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpNot(*this.At(MkString("login")))) {
		panic(NewAuthenticationError(OpAdd(*this.At(MkString("id")), MkString(" fetchAccounts() requires exchange.login email credential"))))
	}
	omsId := this.SafeInteger(*this.At(MkString("options")), MkString("omsId"), MkInteger(1))
	_ = omsId
	this.CheckRequiredCredentials()
	request := MkMap(&VarMap{
		"omsId":    omsId,
		"UserId":   *this.At(MkString("uid")),
		"UserName": *this.At(MkString("login")),
	})
	_ = request
	response := this.Call(MkString("privateGetGetUserAccounts"), this.Extend(request, params))
	_ = response
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
		accountId := this.SafeString(response, i)
		_ = accountId
		result.Push(MkMap(&VarMap{
			"id":       accountId,
			"type":     MkUndefined(),
			"currency": MkUndefined(),
			"info":     accountId,
		}))
	}
	return result
}

func (this *Ndax) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	omsId := this.SafeInteger(*this.At(MkString("options")), MkString("omsId"), MkInteger(1))
	_ = omsId
	this.LoadMarkets()
	this.LoadAccounts()
	defaultAccountId := this.SafeInteger2(*this.At(MkString("options")), MkString("accountId"), MkString("AccountId"), parseInt(*(*(*this.At(MkString("accounts"))).At(MkInteger(0))).At(MkString("id"))))
	_ = defaultAccountId
	accountId := this.SafeInteger2(params, MkString("accountId"), MkString("AccountId"), defaultAccountId)
	_ = accountId
	params = this.Omit(params, MkArray(&VarArray{
		MkString("accountId"),
		MkString("AccountId"),
	}))
	request := MkMap(&VarMap{
		"omsId":     omsId,
		"AccountId": accountId,
	})
	_ = request
	response := this.Call(MkString("privateGetGetAccountPositions"), this.Extend(request, params))
	_ = response
	result := MkMap(&VarMap{
		"info":      response,
		"timestamp": MkUndefined(),
		"datetime":  MkUndefined(),
	})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
		balance := *(response).At(i)
		_ = balance
		currencyId := this.SafeString(balance, MkString("ProductId"))
		_ = currencyId
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		account := this.Account()
		_ = account
		*(account).At(MkString("total")) = this.SafeString(balance, MkString("Amount"))
		*(account).At(MkString("used")) = this.SafeString(balance, MkString("Hold"))
		*(result).At(code) = OpCopy(account)
	}
	return this.ParseBalance(result)
}

func (this *Ndax) ParseLedgerEntryType(goArgs ...*Variant) *Variant {
	type_ := GoGetArg(goArgs, 0, MkUndefined())
	_ = type_
	types := MkMap(&VarMap{
		"Trade":             MkString("trade"),
		"Deposit":           MkString("transaction"),
		"Withdraw":          MkString("transaction"),
		"Transfer":          MkString("transfer"),
		"OrderHold":         MkString("trade"),
		"WithdrawHold":      MkString("transaction"),
		"DepositHold":       MkString("transaction"),
		"MarginHold":        MkString("trade"),
		"ManualHold":        MkString("trade"),
		"ManualEntry":       MkString("trade"),
		"MarginAcquisition": MkString("trade"),
		"MarginRelinquish":  MkString("trade"),
		"MarginQuoteHold":   MkString("trade"),
	})
	_ = types
	return this.SafeString(types, type_, type_)
}

func (this *Ndax) ParseLedgerEntry(goArgs ...*Variant) *Variant {
	item := GoGetArg(goArgs, 0, MkUndefined())
	_ = item
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	id := this.SafeString(item, MkString("TransactionId"))
	_ = id
	account := this.SafeString(item, MkString("AccountId"))
	_ = account
	referenceId := this.SafeString(item, MkString("ReferenceId"))
	_ = referenceId
	referenceAccount := this.SafeString(item, MkString("Counterparty"))
	_ = referenceAccount
	type_ := this.ParseLedgerEntryType(this.SafeString(item, MkString("ReferenceType")))
	_ = type_
	currencyId := this.SafeString(item, MkString("ProductId"))
	_ = currencyId
	code := this.SafeCurrencyCode(currencyId, currency)
	_ = code
	credit := this.SafeNumber(item, MkString("CR"))
	_ = credit
	debit := this.SafeNumber(item, MkString("DR"))
	_ = debit
	amount := OpCopy(MkUndefined())
	_ = amount
	direction := OpCopy(MkUndefined())
	_ = direction
	if IsTrue(OpGt(credit, MkInteger(0))) {
		amount = OpCopy(credit)
		direction = MkString("in")
	} else {
		if IsTrue(OpGt(debit, MkInteger(0))) {
			amount = OpCopy(debit)
			direction = MkString("out")
		}
	}
	timestamp := this.SafeInteger(item, MkString("TimeStamp"))
	_ = timestamp
	before := OpCopy(MkUndefined())
	_ = before
	after := this.SafeNumber(item, MkString("Balance"))
	_ = after
	if IsTrue(OpEq2(direction, MkString("out"))) {
		before = this.Sum(after, amount)
	} else {
		if IsTrue(OpEq2(direction, MkString("in"))) {
			before = Math.Max(MkInteger(0), OpSub(after, amount))
		}
	}
	status := MkString("ok")
	_ = status
	return MkMap(&VarMap{
		"info":             item,
		"id":               id,
		"direction":        direction,
		"account":          account,
		"referenceId":      referenceId,
		"referenceAccount": referenceAccount,
		"type":             type_,
		"currency":         code,
		"amount":           amount,
		"before":           before,
		"after":            after,
		"status":           status,
		"timestamp":        timestamp,
		"datetime":         this.Iso8601(timestamp),
		"fee":              MkUndefined(),
	})
}

func (this *Ndax) FetchLedger(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	omsId := this.SafeInteger(*this.At(MkString("options")), MkString("omsId"), MkInteger(1))
	_ = omsId
	this.LoadMarkets()
	this.LoadAccounts()
	defaultAccountId := this.SafeInteger2(*this.At(MkString("options")), MkString("accountId"), MkString("AccountId"), parseInt(*(*(*this.At(MkString("accounts"))).At(MkInteger(0))).At(MkString("id"))))
	_ = defaultAccountId
	accountId := this.SafeInteger2(params, MkString("accountId"), MkString("AccountId"), defaultAccountId)
	_ = accountId
	params = this.Omit(params, MkArray(&VarArray{
		MkString("accountId"),
		MkString("AccountId"),
	}))
	request := MkMap(&VarMap{
		"omsId":     omsId,
		"AccountId": accountId,
	})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("Depth")) = OpCopy(limit)
	}
	response := this.Call(MkString("privateGetGetAccountTransactions"), this.Extend(request, params))
	_ = response
	currency := OpCopy(MkUndefined())
	_ = currency
	if IsTrue(OpNotEq2(code, MkUndefined())) {
		currency = this.Currency(code)
	}
	return this.ParseLedger(response, currency, since, limit)
}

func (this *Ndax) ParseOrderStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"Accepted":      MkString("open"),
		"Rejected":      MkString("rejected"),
		"Working":       MkString("open"),
		"Canceled":      MkString("canceled"),
		"Expired":       MkString("expired"),
		"FullyExecuted": MkString("closed"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Ndax) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	id := this.SafeString2(order, MkString("ReplacementOrderId"), MkString("OrderId"))
	_ = id
	timestamp := this.SafeInteger(order, MkString("ReceiveTime"))
	_ = timestamp
	lastTradeTimestamp := this.SafeInteger(order, MkString("LastUpdatedTime"))
	_ = lastTradeTimestamp
	marketId := this.SafeString(order, MkString("Instrument"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market)
	_ = symbol
	side := this.SafeStringLower(order, MkString("Side"))
	_ = side
	type_ := this.SafeStringLower(order, MkString("OrderType"))
	_ = type_
	clientOrderId := this.SafeString2(order, MkString("ReplacementClOrdId"), MkString("ClientOrderId"))
	_ = clientOrderId
	price := this.SafeNumber(order, MkString("Price"), MkNumber(0.0))
	_ = price
	price = OpTrinary(OpGt(price, MkNumber(0.0)), price, MkUndefined())
	amount := this.SafeNumber(order, MkString("OrigQuantity"))
	_ = amount
	filled := this.SafeNumber(order, MkString("QuantityExecuted"))
	_ = filled
	cost := this.SafeNumber(order, MkString("GrossValueExecuted"))
	_ = cost
	average := this.SafeNumber(order, MkString("AvgPrice"), MkNumber(0.0))
	_ = average
	average = OpTrinary(OpGt(average, MkInteger(0)), average, MkUndefined())
	stopPrice := this.SafeNumber(order, MkString("StopPrice"), MkNumber(0.0))
	_ = stopPrice
	stopPrice = OpTrinary(OpGt(stopPrice, MkNumber(0.0)), stopPrice, MkUndefined())
	timeInForce := OpCopy(MkUndefined())
	_ = timeInForce
	status := this.ParseOrderStatus(this.SafeString(order, MkString("OrderState")))
	_ = status
	fee := OpCopy(MkUndefined())
	_ = fee
	trades := OpCopy(MkUndefined())
	_ = trades
	return this.SafeOrder(MkMap(&VarMap{
		"id":                 id,
		"clientOrderId":      clientOrderId,
		"info":               order,
		"timestamp":          timestamp,
		"datetime":           this.Iso8601(timestamp),
		"lastTradeTimestamp": lastTradeTimestamp,
		"status":             status,
		"symbol":             symbol,
		"type":               type_,
		"timeInForce":        timeInForce,
		"postOnly":           MkUndefined(),
		"side":               side,
		"price":              price,
		"stopPrice":          stopPrice,
		"cost":               cost,
		"amount":             amount,
		"filled":             filled,
		"average":            average,
		"remaining":          MkUndefined(),
		"fee":                fee,
		"trades":             trades,
	}))
}

func (this *Ndax) CreateOrder(goArgs ...*Variant) *Variant {
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
	omsId := this.SafeInteger(*this.At(MkString("options")), MkString("omsId"), MkInteger(1))
	_ = omsId
	this.LoadMarkets()
	this.LoadAccounts()
	defaultAccountId := this.SafeInteger2(*this.At(MkString("options")), MkString("accountId"), MkString("AccountId"), parseInt(*(*(*this.At(MkString("accounts"))).At(MkInteger(0))).At(MkString("id"))))
	_ = defaultAccountId
	accountId := this.SafeInteger2(params, MkString("accountId"), MkString("AccountId"), defaultAccountId)
	_ = accountId
	clientOrderId := this.SafeInteger2(params, MkString("ClientOrderId"), MkString("clientOrderId"))
	_ = clientOrderId
	params = this.Omit(params, MkArray(&VarArray{
		MkString("accountId"),
		MkString("AccountId"),
		MkString("clientOrderId"),
		MkString("ClientOrderId"),
	}))
	market := this.Market(symbol)
	_ = market
	orderSide := OpTrinary(OpEq2(side, MkString("buy")), MkInteger(0), MkInteger(1))
	_ = orderSide
	request := MkMap(&VarMap{
		"InstrumentId": parseInt(*(market).At(MkString("id"))),
		"omsId":        omsId,
		"AccountId":    accountId,
		"TimeInForce":  MkInteger(1),
		"Side":         orderSide,
		"Quantity":     parseFloat(this.AmountToPrecision(symbol, amount)),
		"OrderType":    this.SafeInteger(*(*this.At(MkString("options"))).At(MkString("orderTypes")), this.Capitalize(type_)),
	})
	_ = request
	if IsTrue(OpNotEq2(price, MkUndefined())) {
		*(request).At(MkString("LimitPrice")) = parseFloat(this.PriceToPrecision(symbol, price))
	}
	if IsTrue(OpNotEq2(clientOrderId, MkUndefined())) {
		*(request).At(MkString("ClientOrderId")) = OpCopy(clientOrderId)
	}
	response := this.Call(MkString("privatePostSendOrder"), this.Extend(request, params))
	_ = response
	return this.ParseOrder(response, market)
}

func (this *Ndax) EditOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	type_ := GoGetArg(goArgs, 2, MkUndefined())
	_ = type_
	side := GoGetArg(goArgs, 3, MkUndefined())
	_ = side
	amount := GoGetArg(goArgs, 4, MkUndefined())
	_ = amount
	price := GoGetArg(goArgs, 5, MkUndefined())
	_ = price
	params := GoGetArg(goArgs, 6, MkMap(&VarMap{}))
	_ = params
	omsId := this.SafeInteger(*this.At(MkString("options")), MkString("omsId"), MkInteger(1))
	_ = omsId
	this.LoadMarkets()
	this.LoadAccounts()
	defaultAccountId := this.SafeInteger2(*this.At(MkString("options")), MkString("accountId"), MkString("AccountId"), parseInt(*(*(*this.At(MkString("accounts"))).At(MkInteger(0))).At(MkString("id"))))
	_ = defaultAccountId
	accountId := this.SafeInteger2(params, MkString("accountId"), MkString("AccountId"), defaultAccountId)
	_ = accountId
	clientOrderId := this.SafeInteger2(params, MkString("ClientOrderId"), MkString("clientOrderId"))
	_ = clientOrderId
	params = this.Omit(params, MkArray(&VarArray{
		MkString("accountId"),
		MkString("AccountId"),
		MkString("clientOrderId"),
		MkString("ClientOrderId"),
	}))
	market := this.Market(symbol)
	_ = market
	orderSide := OpTrinary(OpEq2(side, MkString("buy")), MkInteger(0), MkInteger(1))
	_ = orderSide
	request := MkMap(&VarMap{
		"OrderIdToReplace": parseInt(id),
		"InstrumentId":     parseInt(*(market).At(MkString("id"))),
		"omsId":            omsId,
		"AccountId":        accountId,
		"TimeInForce":      MkInteger(1),
		"Side":             orderSide,
		"Quantity":         parseFloat(this.AmountToPrecision(symbol, amount)),
		"OrderType":        this.SafeInteger(*(*this.At(MkString("options"))).At(MkString("orderTypes")), this.Capitalize(type_)),
	})
	_ = request
	if IsTrue(OpNotEq2(price, MkUndefined())) {
		*(request).At(MkString("LimitPrice")) = parseFloat(this.PriceToPrecision(symbol, price))
	}
	if IsTrue(OpNotEq2(clientOrderId, MkUndefined())) {
		*(request).At(MkString("ClientOrderId")) = OpCopy(clientOrderId)
	}
	response := this.Call(MkString("privatePostCancelReplaceOrder"), this.Extend(request, params))
	_ = response
	return this.ParseOrder(response, market)
}

func (this *Ndax) FetchMyTrades(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	omsId := this.SafeInteger(*this.At(MkString("options")), MkString("omsId"), MkInteger(1))
	_ = omsId
	this.LoadMarkets()
	this.LoadAccounts()
	defaultAccountId := this.SafeInteger2(*this.At(MkString("options")), MkString("accountId"), MkString("AccountId"), parseInt(*(*(*this.At(MkString("accounts"))).At(MkInteger(0))).At(MkString("id"))))
	_ = defaultAccountId
	accountId := this.SafeInteger2(params, MkString("accountId"), MkString("AccountId"), defaultAccountId)
	_ = accountId
	params = this.Omit(params, MkArray(&VarArray{
		MkString("accountId"),
		MkString("AccountId"),
	}))
	request := MkMap(&VarMap{
		"omsId":     omsId,
		"AccountId": accountId,
	})
	_ = request
	market := OpCopy(MkUndefined())
	_ = market
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
		*(request).At(MkString("InstrumentId")) = *(market).At(MkString("id"))
	}
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("StartTimeStamp")) = parseInt(OpDiv(since, MkInteger(1000)))
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("Depth")) = OpCopy(limit)
	}
	response := this.Call(MkString("privateGetGetTradesHistory"), this.Extend(request, params))
	_ = response
	return this.ParseTrades(response, market, since, limit)
}

func (this *Ndax) CancelAllOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	omsId := this.SafeInteger(*this.At(MkString("options")), MkString("omsId"), MkInteger(1))
	_ = omsId
	this.LoadMarkets()
	this.LoadAccounts()
	defaultAccountId := this.SafeInteger2(*this.At(MkString("options")), MkString("accountId"), MkString("AccountId"), parseInt(*(*(*this.At(MkString("accounts"))).At(MkInteger(0))).At(MkString("id"))))
	_ = defaultAccountId
	accountId := this.SafeInteger2(params, MkString("accountId"), MkString("AccountId"), defaultAccountId)
	_ = accountId
	params = this.Omit(params, MkArray(&VarArray{
		MkString("accountId"),
		MkString("AccountId"),
	}))
	request := MkMap(&VarMap{
		"omsId":     omsId,
		"AccountId": accountId,
	})
	_ = request
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market := this.Market(symbol)
		_ = market
		*(request).At(MkString("IntrumentId")) = *(market).At(MkString("id"))
	}
	response := this.Call(MkString("privatePostCancelAllOrders"), this.Extend(request, params))
	_ = response
	return response
}

func (this *Ndax) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	omsId := this.SafeInteger(*this.At(MkString("options")), MkString("omsId"), MkInteger(1))
	_ = omsId
	this.LoadMarkets()
	this.LoadAccounts()
	market := OpCopy(MkUndefined())
	_ = market
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
	}
	request := MkMap(&VarMap{"omsId": omsId})
	_ = request
	clientOrderId := this.SafeInteger2(params, MkString("clientOrderId"), MkString("ClOrderId"))
	_ = clientOrderId
	if IsTrue(OpNotEq2(clientOrderId, MkUndefined())) {
		*(request).At(MkString("ClOrderId")) = OpCopy(clientOrderId)
	} else {
		*(request).At(MkString("OrderId")) = parseInt(id)
	}
	params = this.Omit(params, MkArray(&VarArray{
		MkString("clientOrderId"),
		MkString("ClOrderId"),
	}))
	response := this.Call(MkString("privatePostCancelOrder"), this.Extend(request, params))
	_ = response
	order := this.ParseOrder(response, market)
	_ = order
	return this.Extend(order, MkMap(&VarMap{
		"id":            id,
		"clientOrderId": clientOrderId,
	}))
}

func (this *Ndax) FetchOpenOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	omsId := this.SafeInteger(*this.At(MkString("options")), MkString("omsId"), MkInteger(1))
	_ = omsId
	this.LoadMarkets()
	this.LoadAccounts()
	defaultAccountId := this.SafeInteger2(*this.At(MkString("options")), MkString("accountId"), MkString("AccountId"), parseInt(*(*(*this.At(MkString("accounts"))).At(MkInteger(0))).At(MkString("id"))))
	_ = defaultAccountId
	accountId := this.SafeInteger2(params, MkString("accountId"), MkString("AccountId"), defaultAccountId)
	_ = accountId
	params = this.Omit(params, MkArray(&VarArray{
		MkString("accountId"),
		MkString("AccountId"),
	}))
	market := OpCopy(MkUndefined())
	_ = market
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
	}
	request := MkMap(&VarMap{
		"omsId":     omsId,
		"AccountId": accountId,
	})
	_ = request
	response := this.Call(MkString("privateGetGetOpenOrders"), this.Extend(request, params))
	_ = response
	return this.ParseOrders(response, market, since, limit)
}

func (this *Ndax) FetchOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	omsId := this.SafeInteger(*this.At(MkString("options")), MkString("omsId"), MkInteger(1))
	_ = omsId
	this.LoadMarkets()
	this.LoadAccounts()
	defaultAccountId := this.SafeInteger2(*this.At(MkString("options")), MkString("accountId"), MkString("AccountId"), parseInt(*(*(*this.At(MkString("accounts"))).At(MkInteger(0))).At(MkString("id"))))
	_ = defaultAccountId
	accountId := this.SafeInteger2(params, MkString("accountId"), MkString("AccountId"), defaultAccountId)
	_ = accountId
	params = this.Omit(params, MkArray(&VarArray{
		MkString("accountId"),
		MkString("AccountId"),
	}))
	request := MkMap(&VarMap{
		"omsId":     omsId,
		"AccountId": accountId,
	})
	_ = request
	market := OpCopy(MkUndefined())
	_ = market
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
		*(request).At(MkString("InstrumentId")) = *(market).At(MkString("id"))
	}
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("StartTimeStamp")) = parseInt(OpDiv(since, MkInteger(1000)))
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("Depth")) = OpCopy(limit)
	}
	response := this.Call(MkString("privateGetGetOrdersHistory"), this.Extend(request, params))
	_ = response
	return this.ParseOrders(response, market, since, limit)
}

func (this *Ndax) FetchOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	omsId := this.SafeInteger(*this.At(MkString("options")), MkString("omsId"), MkInteger(1))
	_ = omsId
	this.LoadMarkets()
	this.LoadAccounts()
	defaultAccountId := this.SafeInteger2(*this.At(MkString("options")), MkString("accountId"), MkString("AccountId"), parseInt(*(*(*this.At(MkString("accounts"))).At(MkInteger(0))).At(MkString("id"))))
	_ = defaultAccountId
	accountId := this.SafeInteger2(params, MkString("accountId"), MkString("AccountId"), defaultAccountId)
	_ = accountId
	params = this.Omit(params, MkArray(&VarArray{
		MkString("accountId"),
		MkString("AccountId"),
	}))
	market := OpCopy(MkUndefined())
	_ = market
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
	}
	request := MkMap(&VarMap{
		"omsId":     omsId,
		"AccountId": accountId,
		"OrderId":   parseInt(id),
	})
	_ = request
	response := this.Call(MkString("privateGetGetOrderStatus"), this.Extend(request, params))
	_ = response
	return this.ParseOrder(response, market)
}

func (this *Ndax) FetchOrderTrades(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 2, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 3, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 4, MkMap(&VarMap{}))
	_ = params
	omsId := this.SafeInteger(*this.At(MkString("options")), MkString("omsId"), MkInteger(1))
	_ = omsId
	this.LoadMarkets()
	this.LoadAccounts()
	market := OpCopy(MkUndefined())
	_ = market
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
	}
	request := MkMap(&VarMap{
		"OMSId":   parseInt(omsId),
		"OrderId": parseInt(id),
	})
	_ = request
	response := this.Call(MkString("privatePostGetOrderHistoryByOrderId"), this.Extend(request, params))
	_ = response
	grouped := this.GroupBy(response, MkString("ChangeReason"))
	_ = grouped
	trades := this.SafeValue(grouped, MkString("Trade"), MkArray(&VarArray{}))
	_ = trades
	return this.ParseTrades(trades, market, since, limit)
}

func (this *Ndax) FetchDepositAddress(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	omsId := this.SafeInteger(*this.At(MkString("options")), MkString("omsId"), MkInteger(1))
	_ = omsId
	this.LoadMarkets()
	this.LoadAccounts()
	defaultAccountId := this.SafeInteger2(*this.At(MkString("options")), MkString("accountId"), MkString("AccountId"), parseInt(*(*(*this.At(MkString("accounts"))).At(MkInteger(0))).At(MkString("id"))))
	_ = defaultAccountId
	accountId := this.SafeInteger2(params, MkString("accountId"), MkString("AccountId"), defaultAccountId)
	_ = accountId
	params = this.Omit(params, MkArray(&VarArray{
		MkString("accountId"),
		MkString("AccountId"),
	}))
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{
		"omsId":          omsId,
		"AccountId":      accountId,
		"ProductId":      *(currency).At(MkString("id")),
		"GenerateNewKey": MkBool(false),
	})
	_ = request
	response := this.Call(MkString("privateGetGetDepositInfo"), this.Extend(request, params))
	_ = response
	return this.ParseDepositAddress(response, currency)
}

func (this *Ndax) ParseDepositAddress(goArgs ...*Variant) *Variant {
	depositAddress := GoGetArg(goArgs, 0, MkUndefined())
	_ = depositAddress
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	depositInfoString := this.SafeString(depositAddress, MkString("DepositInfo"))
	_ = depositInfoString
	depositInfo := JSON.Parse(depositInfoString)
	_ = depositInfo
	depositInfoLength := OpCopy(depositInfo.Length)
	_ = depositInfoLength
	lastString := this.SafeString(depositInfo, OpSub(depositInfoLength, MkInteger(1)))
	_ = lastString
	parts := lastString.Split(MkString("?memo="))
	_ = parts
	address := this.SafeString(parts, MkInteger(0))
	_ = address
	tag := this.SafeString(parts, MkInteger(1))
	_ = tag
	code := OpCopy(MkUndefined())
	_ = code
	if IsTrue(OpNotEq2(currency, MkUndefined())) {
		code = *(currency).At(MkString("code"))
	}
	this.CheckAddress(address)
	return MkMap(&VarMap{
		"currency": code,
		"address":  address,
		"tag":      tag,
		"info":     depositAddress,
	})
}

func (this *Ndax) CreateDepositAddress(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	request := MkMap(&VarMap{"GenerateNewKey": MkBool(true)})
	_ = request
	return this.FetchDepositAddress(code, this.Extend(request, params))
}

func (this *Ndax) FetchDeposits(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	omsId := this.SafeInteger(*this.At(MkString("options")), MkString("omsId"), MkInteger(1))
	_ = omsId
	this.LoadMarkets()
	this.LoadAccounts()
	defaultAccountId := this.SafeInteger2(*this.At(MkString("options")), MkString("accountId"), MkString("AccountId"), parseInt(*(*(*this.At(MkString("accounts"))).At(MkInteger(0))).At(MkString("id"))))
	_ = defaultAccountId
	accountId := this.SafeInteger2(params, MkString("accountId"), MkString("AccountId"), defaultAccountId)
	_ = accountId
	params = this.Omit(params, MkArray(&VarArray{
		MkString("accountId"),
		MkString("AccountId"),
	}))
	currency := OpCopy(MkUndefined())
	_ = currency
	if IsTrue(OpNotEq2(code, MkUndefined())) {
		currency = this.Currency(code)
	}
	request := MkMap(&VarMap{
		"omsId":     omsId,
		"AccountId": accountId,
	})
	_ = request
	response := this.Call(MkString("privateGetGetDeposits"), this.Extend(request, params))
	_ = response
	return this.ParseTransactions(response, currency, since, limit)
}

func (this *Ndax) FetchWithdrawals(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	omsId := this.SafeInteger(*this.At(MkString("options")), MkString("omsId"), MkInteger(1))
	_ = omsId
	this.LoadMarkets()
	this.LoadAccounts()
	defaultAccountId := this.SafeInteger2(*this.At(MkString("options")), MkString("accountId"), MkString("AccountId"), parseInt(*(*(*this.At(MkString("accounts"))).At(MkInteger(0))).At(MkString("id"))))
	_ = defaultAccountId
	accountId := this.SafeInteger2(params, MkString("accountId"), MkString("AccountId"), defaultAccountId)
	_ = accountId
	params = this.Omit(params, MkArray(&VarArray{
		MkString("accountId"),
		MkString("AccountId"),
	}))
	currency := OpCopy(MkUndefined())
	_ = currency
	if IsTrue(OpNotEq2(code, MkUndefined())) {
		currency = this.Currency(code)
	}
	request := MkMap(&VarMap{
		"omsId":     omsId,
		"AccountId": accountId,
	})
	_ = request
	response := this.Call(MkString("privateGetGetWithdraws"), this.Extend(request, params))
	_ = response
	return this.ParseTransactions(response, currency, since, limit)
}

func (this *Ndax) ParseTransactionStatusByType(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	type_ := GoGetArg(goArgs, 1, MkUndefined())
	_ = type_
	statusesByType := MkMap(&VarMap{
		"deposit": MkMap(&VarMap{
			"New":              MkString("pending"),
			"AdminProcessing":  MkString("pending"),
			"Accepted":         MkString("pending"),
			"Rejected":         MkString("rejected"),
			"SystemProcessing": MkString("pending"),
			"FullyProcessed":   MkString("ok"),
			"Failed":           MkString("failed"),
			"Pending":          MkString("pending"),
			"Confirmed":        MkString("pending"),
			"AmlProcessing":    MkString("pending"),
			"AmlAccepted":      MkString("pending"),
			"AmlRejected":      MkString("rejected"),
			"AmlFailed":        MkString("failed"),
			"LimitsAccepted":   MkString("pending"),
			"LimitsRejected":   MkString("rejected"),
		}),
		"withdrawal": MkMap(&VarMap{
			"New":               MkString("pending"),
			"AdminProcessing":   MkString("pending"),
			"Accepted":          MkString("pending"),
			"Rejected":          MkString("rejected"),
			"SystemProcessing":  MkString("pending"),
			"FullyProcessed":    MkString("ok"),
			"Failed":            MkString("failed"),
			"Pending":           MkString("pending"),
			"Pending2Fa":        MkString("pending"),
			"AutoAccepted":      MkString("pending"),
			"Delayed":           MkString("pending"),
			"UserCanceled":      MkString("canceled"),
			"AdminCanceled":     MkString("canceled"),
			"AmlProcessing":     MkString("pending"),
			"AmlAccepted":       MkString("pending"),
			"AmlRejected":       MkString("rejected"),
			"AmlFailed":         MkString("failed"),
			"LimitsAccepted":    MkString("pending"),
			"LimitsRejected":    MkString("rejected"),
			"Submitted":         MkString("pending"),
			"Confirmed":         MkString("pending"),
			"ManuallyConfirmed": MkString("pending"),
			"Confirmed2Fa":      MkString("pending"),
		}),
	})
	_ = statusesByType
	statuses := this.SafeValue(statusesByType, type_, MkMap(&VarMap{}))
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Ndax) ParseTransaction(goArgs ...*Variant) *Variant {
	transaction := GoGetArg(goArgs, 0, MkUndefined())
	_ = transaction
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	id := this.SafeString(transaction, MkString("DepositId"))
	_ = id
	txid := OpCopy(MkUndefined())
	_ = txid
	currencyId := this.SafeString(transaction, MkString("ProductId"))
	_ = currencyId
	code := this.SafeCurrencyCode(currencyId, currency)
	_ = code
	timestamp := OpCopy(MkUndefined())
	_ = timestamp
	type_ := OpCopy(MkUndefined())
	_ = type_
	if IsTrue(OpHasMember(MkString("DepositId"), transaction)) {
		type_ = MkString("deposit")
	} else {
		if IsTrue(OpHasMember(MkString("WithdrawId"), transaction)) {
			type_ = MkString("withdrawal")
		}
	}
	templateFormString := this.SafeString(transaction, MkString("TemplateForm"))
	_ = templateFormString
	address := OpCopy(MkUndefined())
	_ = address
	updated := this.SafeInteger(transaction, MkString("LastUpdateTimeStamp"))
	_ = updated
	if IsTrue(OpNotEq2(templateFormString, MkUndefined())) {
		templateForm := JSON.Parse(templateFormString)
		_ = templateForm
		address = this.SafeString(templateForm, MkString("ExternalAddress"))
		txid = this.SafeString(templateForm, MkString("TxId"))
		timestamp = this.SafeInteger(templateForm, MkString("TimeSubmitted"))
		updated = this.SafeInteger(templateForm, MkString("LastUpdated"), updated)
	}
	addressTo := OpCopy(address)
	_ = addressTo
	status := this.ParseTransactionStatusByType(this.SafeString(transaction, MkString("TicketStatus")), type_)
	_ = status
	amount := this.SafeNumber(transaction, MkString("Amount"))
	_ = amount
	feeCost := this.SafeNumber(transaction, MkString("FeeAmount"))
	_ = feeCost
	fee := OpCopy(MkUndefined())
	_ = fee
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		fee = MkMap(&VarMap{
			"currency": code,
			"cost":     feeCost,
		})
	}
	return MkMap(&VarMap{
		"info":        transaction,
		"id":          id,
		"txid":        txid,
		"timestamp":   timestamp,
		"datetime":    this.Iso8601(timestamp),
		"address":     address,
		"addressTo":   addressTo,
		"addressFrom": MkUndefined(),
		"tag":         MkUndefined(),
		"tagTo":       MkUndefined(),
		"tagFrom":     MkUndefined(),
		"type":        type_,
		"amount":      amount,
		"currency":    code,
		"status":      status,
		"updated":     updated,
		"fee":         fee,
	})
}

func (this *Ndax) Withdraw(goArgs ...*Variant) *Variant {
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
	sessionToken := this.SafeString(*this.At(MkString("options")), MkString("sessionToken"))
	_ = sessionToken
	if IsTrue(OpEq2(sessionToken, MkUndefined())) {
		panic(NewAuthenticationError(OpAdd(*this.At(MkString("id")), MkString(" call signIn() method to obtain a session token"))))
	}
	this.CheckAddress(address)
	omsId := this.SafeInteger(*this.At(MkString("options")), MkString("omsId"), MkInteger(1))
	_ = omsId
	this.LoadMarkets()
	this.LoadAccounts()
	defaultAccountId := this.SafeInteger2(*this.At(MkString("options")), MkString("accountId"), MkString("AccountId"), parseInt(*(*(*this.At(MkString("accounts"))).At(MkInteger(0))).At(MkString("id"))))
	_ = defaultAccountId
	accountId := this.SafeInteger2(params, MkString("accountId"), MkString("AccountId"), defaultAccountId)
	_ = accountId
	params = this.Omit(params, MkArray(&VarArray{
		MkString("accountId"),
		MkString("AccountId"),
	}))
	currency := this.Currency(code)
	_ = currency
	withdrawTemplateTypesRequest := MkMap(&VarMap{
		"omsId":     omsId,
		"AccountId": accountId,
		"ProductId": *(currency).At(MkString("id")),
	})
	_ = withdrawTemplateTypesRequest
	withdrawTemplateTypesResponse := this.Call(MkString("privateGetGetWithdrawTemplateTypes"), withdrawTemplateTypesRequest)
	_ = withdrawTemplateTypesResponse
	templateTypes := this.SafeValue(withdrawTemplateTypesResponse, MkString("TemplateTypes"), MkArray(&VarArray{}))
	_ = templateTypes
	firstTemplateType := this.SafeValue(templateTypes, MkInteger(0))
	_ = firstTemplateType
	if IsTrue(OpEq2(firstTemplateType, MkUndefined())) {
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" withdraw() could not find a withdraw template type for "), *(currency).At(MkString("code"))))))
	}
	templateName := this.SafeString(firstTemplateType, MkString("TemplateName"))
	_ = templateName
	withdrawTemplateRequest := MkMap(&VarMap{
		"omsId":             omsId,
		"AccountId":         accountId,
		"ProductId":         *(currency).At(MkString("id")),
		"TemplateType":      templateName,
		"AccountProviderId": *(firstTemplateType).At(MkString("AccountProviderId")),
	})
	_ = withdrawTemplateRequest
	withdrawTemplateResponse := this.Call(MkString("privateGetGetWithdrawTemplate"), withdrawTemplateRequest)
	_ = withdrawTemplateResponse
	template := this.SafeString(withdrawTemplateResponse, MkString("Template"))
	_ = template
	if IsTrue(OpEq2(template, MkUndefined())) {
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" withdraw() could not find a withdraw template for "), *(currency).At(MkString("code"))))))
	}
	withdrawTemplate := JSON.Parse(template)
	_ = withdrawTemplate
	*(withdrawTemplate).At(MkString("ExternalAddress")) = OpCopy(address)
	if IsTrue(OpNotEq2(tag, MkUndefined())) {
		if IsTrue(OpHasMember(MkString("Memo"), withdrawTemplate)) {
			*(withdrawTemplate).At(MkString("Memo")) = OpCopy(tag)
		}
	}
	withdrawPayload := MkMap(&VarMap{
		"omsId":        omsId,
		"AccountId":    accountId,
		"ProductId":    *(currency).At(MkString("id")),
		"TemplateForm": this.Json(withdrawTemplate),
		"TemplateType": templateName,
	})
	_ = withdrawPayload
	withdrawRequest := MkMap(&VarMap{
		"TfaType": MkString("Google"),
		"TFaCode": this.Oath(),
		"Payload": this.Json(withdrawPayload),
	})
	_ = withdrawRequest
	response := this.Call(MkString("privatePostCreateWithdrawTicket"), this.DeepExtend(withdrawRequest, params))
	_ = response
	return MkMap(&VarMap{
		"info": response,
		"id":   this.SafeString(response, MkString("Id")),
	})
}

func (this *Ndax) Nonce(goArgs ...*Variant) *Variant {
	return this.Milliseconds()
}

func (this *Ndax) Sign(goArgs ...*Variant) *Variant {
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
	url := OpAdd(*(*(*this.At(MkString("urls"))).At(MkString("api"))).At(api), OpAdd(MkString("/"), this.ImplodeParams(path, params)))
	_ = url
	query := this.Omit(params, this.ExtractParams(path))
	_ = query
	if IsTrue(OpEq2(api, MkString("public"))) {
		if IsTrue(OpEq2(path, MkString("Authenticate"))) {
			auth := OpAdd(*this.At(MkString("login")), OpAdd(MkString(":"), *this.At(MkString("password"))))
			_ = auth
			auth64 := this.StringToBase64(auth)
			_ = auth64
			headers = MkMap(&VarMap{"Authorization": OpAdd(MkString("Basic "), this.Decode(auth64))})
		} else {
			if IsTrue(OpEq2(path, MkString("Authenticate2FA"))) {
				pending2faToken := this.SafeString(*this.At(MkString("options")), MkString("pending2faToken"))
				_ = pending2faToken
				if IsTrue(OpNotEq2(pending2faToken, MkUndefined())) {
					headers = MkMap(&VarMap{"Pending2FaToken": pending2faToken})
					query = this.Omit(query, MkString("pending2faToken"))
				}
			}
		}
		if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
			OpAddAssign(&url, OpAdd(MkString("?"), this.Urlencode(query)))
		}
	} else {
		if IsTrue(OpEq2(api, MkString("private"))) {
			this.CheckRequiredCredentials()
			sessionToken := this.SafeString(*this.At(MkString("options")), MkString("sessionToken"))
			_ = sessionToken
			if IsTrue(OpEq2(sessionToken, MkUndefined())) {
				nonce := (this.Nonce()).Call(MkString("toString"))
				_ = nonce
				auth := OpAdd(nonce, OpAdd(*this.At(MkString("uid")), *this.At(MkString("apiKey"))))
				_ = auth
				signature := this.Hmac(this.Encode(auth), this.Encode(*this.At(MkString("secret"))))
				_ = signature
				headers = MkMap(&VarMap{
					"Nonce":     nonce,
					"APIKey":    *this.At(MkString("apiKey")),
					"Signature": signature,
					"UserId":    *this.At(MkString("uid")),
				})
			} else {
				headers = MkMap(&VarMap{"APToken": sessionToken})
			}
			if IsTrue(OpEq2(method, MkString("POST"))) {
				*(headers).At(MkString("Content-Type")) = MkString("application/json")
				body = this.Json(query)
			} else {
				if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
					OpAddAssign(&url, OpAdd(MkString("?"), this.Urlencode(query)))
				}
			}
		}
	}
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Ndax) HandleErrors(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpEq2(code, MkInteger(404))) {
		panic(NewAuthenticationError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))))
	}
	if IsTrue(OpEq2(response, MkUndefined())) {
		return MkUndefined()
	}
	message := this.SafeString(response, MkString("errormsg"))
	_ = message
	if IsTrue(OpAnd(OpNotEq2(message, MkUndefined()), OpNotEq2(message, MkString("")))) {
		feedback := OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))
		_ = feedback
		this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("exact")), message, feedback)
		this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("broad")), body, feedback)
		panic(NewExchangeError(feedback))
	}
	return MkUndefined()
}
