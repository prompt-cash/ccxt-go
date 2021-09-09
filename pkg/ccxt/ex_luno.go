package ccxt

import ()

type Luno struct {
	*ExchangeBase
}

var _ Exchange = (*Luno)(nil)

func init() {
	exchange := &Luno{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Luno) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("luno"),
		"name": MkString("luno"),
		"countries": MkArray(&VarArray{
			MkString("GB"),
			MkString("SG"),
			MkString("ZA"),
		}),
		"rateLimit": MkInteger(1000),
		"version":   MkString("1"),
		"has": MkMap(&VarMap{
			"cancelOrder":       MkBool(true),
			"CORS":              MkBool(false),
			"createOrder":       MkBool(true),
			"fetchAccounts":     MkBool(true),
			"fetchBalance":      MkBool(true),
			"fetchClosedOrders": MkBool(true),
			"fetchLedger":       MkBool(true),
			"fetchMarkets":      MkBool(true),
			"fetchMyTrades":     MkBool(true),
			"fetchOpenOrders":   MkBool(true),
			"fetchOrder":        MkBool(true),
			"fetchOrderBook":    MkBool(true),
			"fetchOrders":       MkBool(true),
			"fetchTicker":       MkBool(true),
			"fetchTickers":      MkBool(true),
			"fetchTrades":       MkBool(true),
			"fetchTradingFee":   MkBool(true),
			"fetchTradingFees":  MkBool(true),
		}),
		"urls": MkMap(&VarMap{
			"referral": MkString("https://www.luno.com/invite/44893A"),
			"logo":     MkString("https://user-images.githubusercontent.com/1294454/27766607-8c1a69d8-5ede-11e7-930c-540b5eb9be24.jpg"),
			"api": MkMap(&VarMap{
				"public":   MkString("https://api.luno.com/api"),
				"private":  MkString("https://api.luno.com/api"),
				"exchange": MkString("https://api.luno.com/api/exchange"),
			}),
			"www": MkString("https://www.luno.com"),
			"doc": MkArray(&VarArray{
				MkString("https://www.luno.com/en/api"),
				MkString("https://npmjs.org/package/bitx"),
				MkString("https://github.com/bausmeier/node-bitx"),
			}),
		}),
		"api": MkMap(&VarMap{
			"exchange": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("markets"),
			})}),
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("orderbook"),
				MkString("orderbook_top"),
				MkString("ticker"),
				MkString("tickers"),
				MkString("trades"),
			})}),
			"private": MkMap(&VarMap{
				"get": MkArray(&VarArray{
					MkString("accounts/{id}/pending"),
					MkString("accounts/{id}/transactions"),
					MkString("balance"),
					MkString("beneficiaries"),
					MkString("fee_info"),
					MkString("funding_address"),
					MkString("listorders"),
					MkString("listtrades"),
					MkString("orders/{id}"),
					MkString("quotes/{id}"),
					MkString("withdrawals"),
					MkString("withdrawals/{id}"),
					MkString("transfers"),
				}),
				"post": MkArray(&VarArray{
					MkString("accounts"),
					MkString("accounts/{id}/name"),
					MkString("postorder"),
					MkString("marketorder"),
					MkString("stoporder"),
					MkString("funding_address"),
					MkString("withdrawals"),
					MkString("send"),
					MkString("quotes"),
					MkString("oauth2/grant"),
				}),
				"put": MkArray(&VarArray{
					MkString("accounts/{id}/name"),
					MkString("quotes/{id}"),
				}),
				"delete": MkArray(&VarArray{
					MkString("quotes/{id}"),
					MkString("withdrawals/{id}"),
				}),
			}),
		}),
	}))
}

func (this *Luno) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("exchangeGetMarkets"), params)
	_ = response
	result := MkArray(&VarArray{})
	_ = result
	markets := this.SafeValue(response, MkString("markets"), MkArray(&VarArray{}))
	_ = markets
	for i := MkInteger(0); IsTrue(OpLw(i, markets.Length)); OpIncr(&i) {
		market := *(markets).At(i)
		_ = market
		id := this.SafeString(market, MkString("market_id"))
		_ = id
		baseId := this.SafeString(market, MkString("base_currency"))
		_ = baseId
		quoteId := this.SafeString(market, MkString("counter_currency"))
		_ = quoteId
		base := this.SafeCurrencyCode(baseId)
		_ = base
		quote := this.SafeCurrencyCode(quoteId)
		_ = quote
		symbol := OpAdd(base, OpAdd(MkString("/"), quote))
		_ = symbol
		status := this.SafeString(market, MkString("trading_status"))
		_ = status
		active := OpEq2(status, MkString("ACTIVE"))
		_ = active
		precision := MkMap(&VarMap{
			"amount": this.SafeInteger(market, MkString("volume_scale")),
			"price":  this.SafeInteger(market, MkString("price_scale")),
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
					"min": this.SafeNumber(market, MkString("min_volume")),
					"max": this.SafeNumber(market, MkString("max_volume")),
				}),
				"price": MkMap(&VarMap{
					"min": this.SafeNumber(market, MkString("min_price")),
					"max": this.SafeNumber(market, MkString("max_price")),
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

func (this *Luno) FetchAccounts(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("privateGetBalance"), params)
	_ = response
	wallets := this.SafeValue(response, MkString("balance"), MkArray(&VarArray{}))
	_ = wallets
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, wallets.Length)); OpIncr(&i) {
		account := *(wallets).At(i)
		_ = account
		accountId := this.SafeString(account, MkString("account_id"))
		_ = accountId
		currencyId := this.SafeString(account, MkString("asset"))
		_ = currencyId
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		result.Push(MkMap(&VarMap{
			"id":       accountId,
			"type":     MkUndefined(),
			"currency": code,
			"info":     account,
		}))
	}
	return result
}

func (this *Luno) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privateGetBalance"), params)
	_ = response
	wallets := this.SafeValue(response, MkString("balance"), MkArray(&VarArray{}))
	_ = wallets
	result := MkMap(&VarMap{
		"info":      response,
		"timestamp": MkUndefined(),
		"datetime":  MkUndefined(),
	})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, wallets.Length)); OpIncr(&i) {
		wallet := *(wallets).At(i)
		_ = wallet
		currencyId := this.SafeString(wallet, MkString("asset"))
		_ = currencyId
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		reserved := this.SafeString(wallet, MkString("reserved"))
		_ = reserved
		unconfirmed := this.SafeString(wallet, MkString("unconfirmed"))
		_ = unconfirmed
		balance := this.SafeString(wallet, MkString("balance"))
		_ = balance
		reservedUnconfirmed := Precise.StringAdd(reserved, unconfirmed)
		_ = reservedUnconfirmed
		balanceUnconfirmed := Precise.StringAdd(balance, unconfirmed)
		_ = balanceUnconfirmed
		if IsTrue(OpHasMember(code, result)) {
			*(*(result).At(code)).At(MkString("used")) = Precise.StringAdd(*(*(result).At(code)).At(MkString("used")), reservedUnconfirmed)
			*(*(result).At(code)).At(MkString("total")) = Precise.StringAdd(*(*(result).At(code)).At(MkString("total")), balanceUnconfirmed)
		} else {
			account := this.Account()
			_ = account
			*(account).At(MkString("used")) = OpCopy(reservedUnconfirmed)
			*(account).At(MkString("total")) = OpCopy(balanceUnconfirmed)
			*(result).At(code) = OpCopy(account)
		}
	}
	return this.ParseBalance(result)
}

func (this *Luno) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	method := MkString("publicGetOrderbook")
	_ = method
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		if IsTrue(OpLwEq(limit, MkInteger(100))) {
			OpAddAssign(&method, MkString("Top"))
		}
	}
	request := MkMap(&VarMap{"pair": this.MarketId(symbol)})
	_ = request
	response := this.Call(method, this.Extend(request, params))
	_ = response
	timestamp := this.SafeInteger(response, MkString("timestamp"))
	_ = timestamp
	return this.ParseOrderBook(response, symbol, timestamp, MkString("bids"), MkString("asks"), MkString("price"), MkString("volume"))
}

func (this *Luno) ParseOrderStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{"PENDING": MkString("open")})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Luno) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.SafeInteger(order, MkString("creation_timestamp"))
	_ = timestamp
	status := this.ParseOrderStatus(this.SafeString(order, MkString("state")))
	_ = status
	status = OpTrinary(OpEq2(status, MkString("open")), status, status)
	side := OpCopy(MkUndefined())
	_ = side
	orderType := this.SafeString(order, MkString("type"))
	_ = orderType
	if IsTrue(OpOr(OpEq2(orderType, MkString("ASK")), OpEq2(orderType, MkString("SELL")))) {
		side = MkString("sell")
	} else {
		if IsTrue(OpOr(OpEq2(orderType, MkString("BID")), OpEq2(orderType, MkString("BUY")))) {
			side = MkString("buy")
		}
	}
	marketId := this.SafeString(order, MkString("pair"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market)
	_ = symbol
	price := this.SafeNumber(order, MkString("limit_price"))
	_ = price
	amount := this.SafeNumber(order, MkString("limit_volume"))
	_ = amount
	quoteFee := this.SafeNumber(order, MkString("fee_counter"))
	_ = quoteFee
	baseFee := this.SafeNumber(order, MkString("fee_base"))
	_ = baseFee
	filled := this.SafeNumber(order, MkString("base"))
	_ = filled
	cost := this.SafeNumber(order, MkString("counter"))
	_ = cost
	fee := MkMap(&VarMap{"currency": MkUndefined()})
	_ = fee
	if IsTrue(quoteFee) {
		*(fee).At(MkString("cost")) = OpCopy(quoteFee)
		if IsTrue(OpNotEq2(market, MkUndefined())) {
			*(fee).At(MkString("currency")) = *(market).At(MkString("quote"))
		}
	} else {
		*(fee).At(MkString("cost")) = OpCopy(baseFee)
		if IsTrue(OpNotEq2(market, MkUndefined())) {
			*(fee).At(MkString("currency")) = *(market).At(MkString("base"))
		}
	}
	id := this.SafeString(order, MkString("order_id"))
	_ = id
	return this.SafeOrder(MkMap(&VarMap{
		"id":                 id,
		"clientOrderId":      MkUndefined(),
		"datetime":           this.Iso8601(timestamp),
		"timestamp":          timestamp,
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
		"filled":             filled,
		"cost":               cost,
		"remaining":          MkUndefined(),
		"trades":             MkUndefined(),
		"fee":                fee,
		"info":               order,
		"average":            MkUndefined(),
	}))
}

func (this *Luno) FetchOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"id": id})
	_ = request
	response := this.Call(MkString("privateGetOrdersId"), this.Extend(request, params))
	_ = response
	return this.ParseOrder(response)
}

func (this *Luno) FetchOrdersByState(goArgs ...*Variant) *Variant {
	state := GoGetArg(goArgs, 0, MkUndefined())
	_ = state
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 2, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 3, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 4, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{})
	_ = request
	market := OpCopy(MkUndefined())
	_ = market
	if IsTrue(OpNotEq2(state, MkUndefined())) {
		*(request).At(MkString("state")) = OpCopy(state)
	}
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
		*(request).At(MkString("pair")) = *(market).At(MkString("id"))
	}
	response := this.Call(MkString("privateGetListorders"), this.Extend(request, params))
	_ = response
	orders := this.SafeValue(response, MkString("orders"), MkArray(&VarArray{}))
	_ = orders
	return this.ParseOrders(orders, market, since, limit)
}

func (this *Luno) FetchOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	return this.FetchOrdersByState(MkUndefined(), symbol, since, limit, params)
}

func (this *Luno) FetchOpenOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	return this.FetchOrdersByState(MkString("PENDING"), symbol, since, limit, params)
}

func (this *Luno) FetchClosedOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	return this.FetchOrdersByState(MkString("COMPLETE"), symbol, since, limit, params)
}

func (this *Luno) ParseTicker(goArgs ...*Variant) *Variant {
	ticker := GoGetArg(goArgs, 0, MkUndefined())
	_ = ticker
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.SafeInteger(ticker, MkString("timestamp"))
	_ = timestamp
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(market) {
		symbol = *(market).At(MkString("symbol"))
	}
	last := this.SafeNumber(ticker, MkString("last_trade"))
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
		"baseVolume":    this.SafeNumber(ticker, MkString("rolling_24_hour_volume")),
		"quoteVolume":   MkUndefined(),
		"info":          ticker,
	})
}

func (this *Luno) FetchTickers(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("publicGetTickers"), params)
	_ = response
	tickers := this.IndexBy(*(response).At(MkString("tickers")), MkString("pair"))
	_ = tickers
	ids := GoGetKeys(tickers)
	_ = ids
	result := MkMap(&VarMap{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, ids.Length)); OpIncr(&i) {
		id := *(ids).At(i)
		_ = id
		market := this.SafeMarket(id)
		_ = market
		symbol := *(market).At(MkString("symbol"))
		_ = symbol
		ticker := *(tickers).At(id)
		_ = ticker
		*(result).At(symbol) = this.ParseTicker(ticker, market)
	}
	return this.FilterByArray(result, MkString("symbol"), symbols)
}

func (this *Luno) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"pair": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("publicGetTicker"), this.Extend(request, params))
	_ = response
	return this.ParseTicker(response, market)
}

func (this *Luno) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	orderId := this.SafeString(trade, MkString("order_id"))
	_ = orderId
	takerOrMaker := OpCopy(MkUndefined())
	_ = takerOrMaker
	side := OpCopy(MkUndefined())
	_ = side
	if IsTrue(OpNotEq2(orderId, MkUndefined())) {
		type_ := this.SafeString(trade, MkString("type"))
		_ = type_
		if IsTrue(OpOr(OpEq2(type_, MkString("ASK")), OpEq2(type_, MkString("SELL")))) {
			side = MkString("sell")
		} else {
			if IsTrue(OpOr(OpEq2(type_, MkString("BID")), OpEq2(type_, MkString("BUY")))) {
				side = MkString("buy")
			}
		}
		if IsTrue(OpAnd(OpEq2(side, MkString("sell")), *(trade).At(MkString("is_buy")))) {
			takerOrMaker = MkString("maker")
		} else {
			if IsTrue(OpAnd(OpEq2(side, MkString("buy")), OpNot(*(trade).At(MkString("is_buy"))))) {
				takerOrMaker = MkString("maker")
			} else {
				takerOrMaker = MkString("taker")
			}
		}
	} else {
		side = OpTrinary(*(trade).At(MkString("is_buy")), MkString("buy"), MkString("sell"))
	}
	feeBase := this.SafeNumber(trade, MkString("fee_base"))
	_ = feeBase
	feeCounter := this.SafeNumber(trade, MkString("fee_counter"))
	_ = feeCounter
	feeCurrency := OpCopy(MkUndefined())
	_ = feeCurrency
	feeCost := OpCopy(MkUndefined())
	_ = feeCost
	if IsTrue(OpNotEq2(feeBase, MkUndefined())) {
		if IsTrue(OpNotEq2(feeBase, MkNumber(0.0))) {
			feeCurrency = *(market).At(MkString("base"))
			feeCost = OpCopy(feeBase)
		}
	} else {
		if IsTrue(OpNotEq2(feeCounter, MkUndefined())) {
			if IsTrue(OpNotEq2(feeCounter, MkNumber(0.0))) {
				feeCurrency = *(market).At(MkString("quote"))
				feeCost = OpCopy(feeCounter)
			}
		}
	}
	timestamp := this.SafeInteger(trade, MkString("timestamp"))
	_ = timestamp
	return MkMap(&VarMap{
		"info":         trade,
		"id":           MkUndefined(),
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       *(market).At(MkString("symbol")),
		"order":        orderId,
		"type":         MkUndefined(),
		"side":         side,
		"takerOrMaker": takerOrMaker,
		"price":        this.SafeNumber(trade, MkString("price")),
		"amount":       this.SafeNumber(trade, MkString("volume")),
		"cost":         this.SafeNumber(trade, MkString("counter")),
		"fee": MkMap(&VarMap{
			"cost":     feeCost,
			"currency": feeCurrency,
		}),
	})
}

func (this *Luno) FetchTrades(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("since")) = OpCopy(since)
	}
	response := this.Call(MkString("publicGetTrades"), this.Extend(request, params))
	_ = response
	trades := this.SafeValue(response, MkString("trades"), MkArray(&VarArray{}))
	_ = trades
	return this.ParseTrades(trades, market, since, limit)
}

func (this *Luno) FetchMyTrades(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchMyTrades() requires a symbol argument"))))
	}
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"pair": *(market).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("since")) = OpCopy(since)
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("privateGetListtrades"), this.Extend(request, params))
	_ = response
	trades := this.SafeValue(response, MkString("trades"), MkArray(&VarArray{}))
	_ = trades
	return this.ParseTrades(trades, market, since, limit)
}

func (this *Luno) FetchTradingFees(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privateGetFeeInfo"), params)
	_ = response
	return MkMap(&VarMap{
		"info":  response,
		"maker": this.SafeNumber(response, MkString("maker_fee")),
		"taker": this.SafeNumber(response, MkString("taker_fee")),
	})
}

func (this *Luno) CreateOrder(goArgs ...*Variant) *Variant {
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
	method := MkString("privatePost")
	_ = method
	request := MkMap(&VarMap{"pair": this.MarketId(symbol)})
	_ = request
	if IsTrue(OpEq2(type_, MkString("market"))) {
		OpAddAssign(&method, MkString("Marketorder"))
		*(request).At(MkString("type")) = side.ToUpperCase()
		if IsTrue(OpEq2(side, MkString("buy"))) {
			*(request).At(MkString("counter_volume")) = parseFloat(this.AmountToPrecision(symbol, amount))
		} else {
			*(request).At(MkString("base_volume")) = parseFloat(this.AmountToPrecision(symbol, amount))
		}
	} else {
		OpAddAssign(&method, MkString("Postorder"))
		*(request).At(MkString("volume")) = parseFloat(this.AmountToPrecision(symbol, amount))
		*(request).At(MkString("price")) = parseFloat(this.PriceToPrecision(symbol, price))
		*(request).At(MkString("type")) = OpTrinary(OpEq2(side, MkString("buy")), MkString("BID"), MkString("ASK"))
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	return MkMap(&VarMap{
		"info": response,
		"id":   *(response).At(MkString("order_id")),
	})
}

func (this *Luno) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"order_id": id})
	_ = request
	return this.Call(MkString("privatePostStoporder"), this.Extend(request, params))
}

func (this *Luno) FetchLedgerByEntries(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	entry := GoGetArg(goArgs, 1, OpNeg(MkInteger(1)))
	_ = entry
	limit := GoGetArg(goArgs, 2, MkInteger(1))
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	since := OpCopy(MkUndefined())
	_ = since
	request := MkMap(&VarMap{
		"min_row": entry,
		"max_row": this.Sum(entry, limit),
	})
	_ = request
	return this.FetchLedger(code, since, limit, this.Extend(request, params))
}

func (this *Luno) FetchLedger(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	this.LoadAccounts()
	currency := OpCopy(MkUndefined())
	_ = currency
	id := this.SafeString(params, MkString("id"))
	_ = id
	min_row := this.SafeValue(params, MkString("min_row"))
	_ = min_row
	max_row := this.SafeValue(params, MkString("max_row"))
	_ = max_row
	if IsTrue(OpEq2(id, MkUndefined())) {
		if IsTrue(OpEq2(code, MkUndefined())) {
			panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchLedger() requires a currency code argument if no account id specified in params"))))
		}
		currency = this.Currency(code)
		accountsByCurrencyCode := this.IndexBy(*this.At(MkString("accounts")), MkString("currency"))
		_ = accountsByCurrencyCode
		account := this.SafeValue(accountsByCurrencyCode, code)
		_ = account
		if IsTrue(OpEq2(account, MkUndefined())) {
			panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" fetchLedger() could not find account id for "), code))))
		}
		id = *(account).At(MkString("id"))
	}
	if IsTrue(OpAnd(OpEq2(min_row, MkUndefined()), OpEq2(max_row, MkUndefined()))) {
		max_row = MkInteger(0)
		min_row = OpNeg(MkInteger(1000))
	} else {
		if IsTrue(OpOr(OpEq2(min_row, MkUndefined()), OpEq2(max_row, MkUndefined()))) {
			panic(NewExchangeError(OpAdd(*this.At(MkString("id")), MkString(" fetchLedger() require both params 'max_row' and 'min_row' or neither to be defined"))))
		}
	}
	if IsTrue(OpAnd(OpNotEq2(limit, MkUndefined()), OpGt(OpSub(max_row, min_row), limit))) {
		if IsTrue(OpLwEq(max_row, MkInteger(0))) {
			min_row = OpSub(max_row, limit)
		} else {
			if IsTrue(OpGt(min_row, MkInteger(0))) {
				max_row = OpAdd(min_row, limit)
			}
		}
	}
	if IsTrue(OpGt(OpSub(max_row, min_row), MkInteger(1000))) {
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), MkString(" fetchLedger() requires the params 'max_row' - 'min_row' <= 1000"))))
	}
	request := MkMap(&VarMap{
		"id":      id,
		"min_row": min_row,
		"max_row": max_row,
	})
	_ = request
	response := this.Call(MkString("privateGetAccountsIdTransactions"), this.Extend(params, request))
	_ = response
	entries := this.SafeValue(response, MkString("transactions"), MkArray(&VarArray{}))
	_ = entries
	return this.ParseLedger(entries, currency, since, limit)
}

func (this *Luno) ParseLedgerComment(goArgs ...*Variant) *Variant {
	comment := GoGetArg(goArgs, 0, MkUndefined())
	_ = comment
	words := comment.Split(MkString(" "))
	_ = words
	types := MkMap(&VarMap{
		"Withdrawal": MkString("fee"),
		"Trading":    MkString("fee"),
		"Payment":    MkString("transaction"),
		"Sent":       MkString("transaction"),
		"Deposit":    MkString("transaction"),
		"Received":   MkString("transaction"),
		"Released":   MkString("released"),
		"Reserved":   MkString("reserved"),
		"Sold":       MkString("trade"),
		"Bought":     MkString("trade"),
		"Failure":    MkString("failed"),
	})
	_ = types
	referenceId := OpCopy(MkUndefined())
	_ = referenceId
	firstWord := this.SafeString(words, MkInteger(0))
	_ = firstWord
	thirdWord := this.SafeString(words, MkInteger(2))
	_ = thirdWord
	fourthWord := this.SafeString(words, MkInteger(3))
	_ = fourthWord
	type_ := this.SafeString(types, firstWord, MkUndefined())
	_ = type_
	if IsTrue(OpAnd(OpEq2(type_, MkUndefined()), OpEq2(thirdWord, MkString("fee")))) {
		type_ = MkString("fee")
	}
	if IsTrue(OpAnd(OpEq2(type_, MkString("reserved")), OpEq2(fourthWord, MkString("order")))) {
		referenceId = this.SafeString(words, MkInteger(4))
	}
	return MkMap(&VarMap{
		"type":        type_,
		"referenceId": referenceId,
	})
}

func (this *Luno) ParseLedgerEntry(goArgs ...*Variant) *Variant {
	entry := GoGetArg(goArgs, 0, MkUndefined())
	_ = entry
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	id := this.SafeString(entry, MkString("row_index"))
	_ = id
	account_id := this.SafeString(entry, MkString("account_id"))
	_ = account_id
	timestamp := this.SafeValue(entry, MkString("timestamp"))
	_ = timestamp
	currencyId := this.SafeString(entry, MkString("currency"))
	_ = currencyId
	code := this.SafeCurrencyCode(currencyId, currency)
	_ = code
	available_delta := this.SafeNumber(entry, MkString("available_delta"))
	_ = available_delta
	balance_delta := this.SafeNumber(entry, MkString("balance_delta"))
	_ = balance_delta
	after := this.SafeNumber(entry, MkString("balance"))
	_ = after
	comment := this.SafeString(entry, MkString("description"))
	_ = comment
	before := OpCopy(after)
	_ = before
	amount := MkNumber(0.0)
	_ = amount
	result := this.ParseLedgerComment(comment)
	_ = result
	type_ := *(result).At(MkString("type"))
	_ = type_
	referenceId := *(result).At(MkString("referenceId"))
	_ = referenceId
	direction := OpCopy(MkUndefined())
	_ = direction
	status := OpCopy(MkUndefined())
	_ = status
	if IsTrue(OpNotEq2(balance_delta, MkNumber(0.0))) {
		before = OpSub(after, balance_delta)
		status = MkString("ok")
		amount = Math.Abs(balance_delta)
	} else {
		if IsTrue(OpLw(available_delta, MkNumber(0.0))) {
			status = MkString("pending")
			amount = Math.Abs(available_delta)
		} else {
			if IsTrue(OpGt(available_delta, MkNumber(0.0))) {
				status = MkString("canceled")
				amount = Math.Abs(available_delta)
			}
		}
	}
	if IsTrue(OpOr(OpGt(balance_delta, MkInteger(0)), OpGt(available_delta, MkInteger(0)))) {
		direction = MkString("in")
	} else {
		if IsTrue(OpOr(OpLw(balance_delta, MkInteger(0)), OpLw(available_delta, MkInteger(0)))) {
			direction = MkString("out")
		}
	}
	return MkMap(&VarMap{
		"id":               id,
		"direction":        direction,
		"account":          account_id,
		"referenceId":      referenceId,
		"referenceAccount": MkUndefined(),
		"type":             type_,
		"currency":         code,
		"amount":           amount,
		"timestamp":        timestamp,
		"datetime":         this.Iso8601(timestamp),
		"before":           before,
		"after":            after,
		"status":           status,
		"fee":              MkUndefined(),
		"info":             entry,
	})
}

func (this *Luno) Sign(goArgs ...*Variant) *Variant {
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
	if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
		OpAddAssign(&url, OpAdd(MkString("?"), this.Urlencode(query)))
	}
	if IsTrue(OpEq2(api, MkString("private"))) {
		this.CheckRequiredCredentials()
		auth := this.StringToBase64(OpAdd(*this.At(MkString("apiKey")), OpAdd(MkString(":"), *this.At(MkString("secret")))))
		_ = auth
		headers = MkMap(&VarMap{"Authorization": OpAdd(MkString("Basic "), this.Decode(auth))})
	}
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Luno) HandleErrors(goArgs ...*Variant) *Variant {
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
	error := this.SafeValue(response, MkString("error"))
	_ = error
	if IsTrue(OpNotEq2(error, MkUndefined())) {
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), this.Json(response)))))
	}
	return MkUndefined()
}
