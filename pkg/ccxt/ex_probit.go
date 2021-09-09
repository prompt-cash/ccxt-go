package ccxt

import ()

type Probit struct {
	*ExchangeBase
}

var _ Exchange = (*Probit)(nil)

func init() {
	exchange := &Probit{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Probit) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("probit"),
		"name": MkString("ProBit"),
		"countries": MkArray(&VarArray{
			MkString("SC"),
			MkString("KR"),
		}),
		"rateLimit": MkInteger(250),
		"has": MkMap(&VarMap{
			"CORS":                MkBool(true),
			"fetchTime":           MkBool(true),
			"fetchMarkets":        MkBool(true),
			"fetchCurrencies":     MkBool(true),
			"fetchTickers":        MkBool(true),
			"fetchTicker":         MkBool(true),
			"fetchOHLCV":          MkBool(true),
			"fetchOrderBook":      MkBool(true),
			"fetchTrades":         MkBool(true),
			"fetchBalance":        MkBool(true),
			"createOrder":         MkBool(true),
			"createMarketOrder":   MkBool(true),
			"cancelOrder":         MkBool(true),
			"fetchOrder":          MkBool(true),
			"fetchOpenOrders":     MkBool(true),
			"fetchClosedOrders":   MkBool(true),
			"fetchMyTrades":       MkBool(true),
			"fetchDepositAddress": MkBool(true),
			"withdraw":            MkBool(true),
			"signIn":              MkBool(true),
		}),
		"timeframes": MkMap(&VarMap{
			"1m":  MkString("1m"),
			"3m":  MkString("3m"),
			"5m":  MkString("5m"),
			"10m": MkString("10m"),
			"15m": MkString("15m"),
			"30m": MkString("30m"),
			"1h":  MkString("1h"),
			"4h":  MkString("4h"),
			"6h":  MkString("6h"),
			"12h": MkString("12h"),
			"1d":  MkString("1D"),
			"1w":  MkString("1W"),
			"1M":  MkString("1M"),
		}),
		"version": MkString("v1"),
		"urls": MkMap(&VarMap{
			"logo": MkString("https://user-images.githubusercontent.com/51840849/79268032-c4379480-7ea2-11ea-80b3-dd96bb29fd0d.jpg"),
			"api": MkMap(&VarMap{
				"accounts": MkString("https://accounts.probit.com"),
				"public":   MkString("https://api.probit.com/api/exchange"),
				"private":  MkString("https://api.probit.com/api/exchange"),
			}),
			"www": MkString("https://www.probit.com"),
			"doc": MkArray(&VarArray{
				MkString("https://docs-en.probit.com"),
				MkString("https://docs-ko.probit.com"),
			}),
			"fees":     MkString("https://support.probit.com/hc/en-us/articles/360020968611-Trading-Fees"),
			"referral": MkString("https://www.probit.com/r/34608773"),
		}),
		"api": MkMap(&VarMap{
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("market"),
				MkString("currency"),
				MkString("currency_with_platform"),
				MkString("time"),
				MkString("ticker"),
				MkString("order_book"),
				MkString("trade"),
				MkString("candle"),
			})}),
			"private": MkMap(&VarMap{
				"post": MkArray(&VarArray{
					MkString("new_order"),
					MkString("cancel_order"),
					MkString("withdrawal"),
				}),
				"get": MkArray(&VarArray{
					MkString("balance"),
					MkString("order"),
					MkString("open_order"),
					MkString("order_history"),
					MkString("trade_history"),
					MkString("deposit_address"),
				}),
			}),
			"accounts": MkMap(&VarMap{"post": MkArray(&VarArray{
				MkString("token"),
			})}),
		}),
		"fees": MkMap(&VarMap{"trading": MkMap(&VarMap{
			"tierBased":  MkBool(false),
			"percentage": MkBool(true),
			"maker":      this.ParseNumber(MkString("0.002")),
			"taker":      this.ParseNumber(MkString("0.002")),
		})}),
		"exceptions": MkMap(&VarMap{"exact": MkMap(&VarMap{
			"UNAUTHORIZED":            AuthenticationError,
			"INVALID_ARGUMENT":        BadRequest,
			"TRADING_UNAVAILABLE":     ExchangeNotAvailable,
			"NOT_ENOUGH_BALANCE":      InsufficientFunds,
			"NOT_ALLOWED_COMBINATION": BadRequest,
			"INVALID_ORDER":           InvalidOrder,
			"RATE_LIMIT_EXCEEDED":     RateLimitExceeded,
			"MARKET_UNAVAILABLE":      ExchangeNotAvailable,
			"INVALID_MARKET":          BadSymbol,
			"MARKET_CLOSED":           BadSymbol,
			"MARKET_NOT_FOUND":        BadSymbol,
			"INVALID_CURRENCY":        BadRequest,
			"TOO_MANY_OPEN_ORDERS":    DDoSProtection,
			"DUPLICATE_ADDRESS":       InvalidAddress,
			"invalid_grant":           AuthenticationError,
		})}),
		"requiredCredentials": MkMap(&VarMap{
			"apiKey": MkBool(true),
			"secret": MkBool(true),
		}),
		"precisionMode": TICK_SIZE,
		"options": MkMap(&VarMap{
			"createMarketBuyOrderRequiresPrice": MkBool(true),
			"timeInForce": MkMap(&VarMap{
				"limit":  MkString("gtc"),
				"market": MkString("ioc"),
			}),
		}),
		"commonCurrencies": MkMap(&VarMap{
			"AUTO":    MkString("Cube"),
			"BCC":     MkString("BCC"),
			"BDP":     MkString("BidiPass"),
			"BTCBEAR": MkString("BEAR"),
			"BTCBULL": MkString("BULL"),
			"CBC":     MkString("CryptoBharatCoin"),
			"EPS":     MkString("Epanus"),
			"GRB":     MkString("Global Reward Bank"),
			"HBC":     MkString("Hybrid Bank Cash"),
			"ORC":     MkString("Oracle System"),
			"SOC":     MkString("Soda Coin"),
			"TCT":     MkString("Top Coin Token"),
			"UNI":     MkString("UNICORN Token"),
			"UNISWAP": MkString("UNI"),
		}),
	}))
}

func (this *Probit) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetMarket"), params)
	_ = response
	markets := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = markets
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, markets.Length)); OpIncr(&i) {
		market := *(markets).At(i)
		_ = market
		id := this.SafeString(market, MkString("id"))
		_ = id
		baseId := this.SafeString(market, MkString("base_currency_id"))
		_ = baseId
		quoteId := this.SafeString(market, MkString("quote_currency_id"))
		_ = quoteId
		base := this.SafeCurrencyCode(baseId)
		_ = base
		quote := this.SafeCurrencyCode(quoteId)
		_ = quote
		symbol := OpAdd(base, OpAdd(MkString("/"), quote))
		_ = symbol
		closed := this.SafeValue(market, MkString("closed"), MkBool(false))
		_ = closed
		active := OpNot(closed)
		_ = active
		amountPrecision := this.SafeString(market, MkString("quantity_precision"))
		_ = amountPrecision
		costPrecision := this.SafeString(market, MkString("cost_precision"))
		_ = costPrecision
		amountTickSize := this.ParsePrecision(amountPrecision)
		_ = amountTickSize
		costTickSize := this.ParsePrecision(costPrecision)
		_ = costTickSize
		precision := MkMap(&VarMap{
			"amount": this.ParseNumber(amountTickSize),
			"price":  this.SafeNumber(market, MkString("price_increment")),
			"cost":   this.ParseNumber(costTickSize),
		})
		_ = precision
		takerFeeRate := this.SafeString(market, MkString("taker_fee_rate"))
		_ = takerFeeRate
		taker := Precise.StringDiv(takerFeeRate, MkString("100"))
		_ = taker
		makerFeeRate := this.SafeString(market, MkString("maker_fee_rate"))
		_ = makerFeeRate
		maker := Precise.StringDiv(makerFeeRate, MkString("100"))
		_ = maker
		result.Push(MkMap(&VarMap{
			"id":        id,
			"info":      market,
			"symbol":    symbol,
			"base":      base,
			"quote":     quote,
			"baseId":    baseId,
			"quoteId":   quoteId,
			"active":    active,
			"precision": precision,
			"taker":     this.ParseNumber(taker),
			"maker":     this.ParseNumber(maker),
			"limits": MkMap(&VarMap{
				"amount": MkMap(&VarMap{
					"min": this.SafeNumber(market, MkString("min_quantity")),
					"max": this.SafeNumber(market, MkString("max_quantity")),
				}),
				"price": MkMap(&VarMap{
					"min": this.SafeNumber(market, MkString("min_price")),
					"max": this.SafeNumber(market, MkString("max_price")),
				}),
				"cost": MkMap(&VarMap{
					"min": this.SafeNumber(market, MkString("min_cost")),
					"max": this.SafeNumber(market, MkString("max_cost")),
				}),
			}),
		}))
	}
	return result
}

func (this *Probit) FetchCurrencies(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetCurrencyWithPlatform"), params)
	_ = response
	currencies := this.SafeValue(response, MkString("data"))
	_ = currencies
	result := MkMap(&VarMap{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, currencies.Length)); OpIncr(&i) {
		currency := *(currencies).At(i)
		_ = currency
		id := this.SafeString(currency, MkString("id"))
		_ = id
		code := this.SafeCurrencyCode(id)
		_ = code
		displayName := this.SafeValue(currency, MkString("display_name"))
		_ = displayName
		name := this.SafeString(displayName, MkString("en-us"))
		_ = name
		platforms := this.SafeValue(currency, MkString("platform"), MkArray(&VarArray{}))
		_ = platforms
		platformsByPriority := this.SortBy(platforms, MkString("priority"))
		_ = platformsByPriority
		platform := this.SafeValue(platformsByPriority, MkInteger(0), MkMap(&VarMap{}))
		_ = platform
		precision := this.SafeInteger(platform, MkString("precision"))
		_ = precision
		depositSuspended := this.SafeValue(platform, MkString("deposit_suspended"))
		_ = depositSuspended
		withdrawalSuspended := this.SafeValue(platform, MkString("withdrawal_suspended"))
		_ = withdrawalSuspended
		active := OpNot(OpAnd(depositSuspended, withdrawalSuspended))
		_ = active
		withdrawalFees := this.SafeValue(platform, MkString("withdrawal_fee"), MkMap(&VarMap{}))
		_ = withdrawalFees
		fees := MkArray(&VarArray{})
		_ = fees
		for j := MkInteger(0); IsTrue(OpLw(j, withdrawalFees.Length)); OpIncr(&j) {
			withdrawalFee := *(withdrawalFees).At(j)
			_ = withdrawalFee
			amount := this.SafeNumber(withdrawalFee, MkString("amount"))
			_ = amount
			priority := this.SafeInteger(withdrawalFee, MkString("priority"))
			_ = priority
			if IsTrue(OpAnd(OpNotEq2(amount, MkUndefined()), OpNotEq2(priority, MkUndefined()))) {
				fees.Push(withdrawalFee)
			}
		}
		withdrawalFeesByPriority := this.SortBy(fees, MkString("priority"))
		_ = withdrawalFeesByPriority
		withdrawalFee := this.SafeValue(withdrawalFeesByPriority, MkInteger(0), MkMap(&VarMap{}))
		_ = withdrawalFee
		fee := this.SafeNumber(withdrawalFee, MkString("amount"))
		_ = fee
		*(result).At(code) = MkMap(&VarMap{
			"id":        id,
			"code":      code,
			"info":      currency,
			"name":      name,
			"active":    active,
			"fee":       fee,
			"precision": precision,
			"limits": MkMap(&VarMap{
				"amount": MkMap(&VarMap{
					"min": Math.Pow(MkInteger(10), OpNeg(precision)),
					"max": Math.Pow(MkInteger(10), precision),
				}),
				"deposit": MkMap(&VarMap{
					"min": this.SafeNumber(platform, MkString("min_deposit_amount")),
					"max": MkUndefined(),
				}),
				"withdraw": MkMap(&VarMap{
					"min": this.SafeNumber(platform, MkString("min_withdrawal_amount")),
					"max": MkUndefined(),
				}),
			}),
		})
	}
	return result
}

func (this *Probit) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privateGetBalance"), params)
	_ = response
	result := MkMap(&VarMap{
		"info":      response,
		"timestamp": MkUndefined(),
		"datetime":  MkUndefined(),
	})
	_ = result
	data := this.SafeValue(response, MkString("data"))
	_ = data
	for i := MkInteger(0); IsTrue(OpLw(i, data.Length)); OpIncr(&i) {
		balance := *(data).At(i)
		_ = balance
		currencyId := this.SafeString(balance, MkString("currency_id"))
		_ = currencyId
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		account := this.Account()
		_ = account
		*(account).At(MkString("total")) = this.SafeString(balance, MkString("total"))
		*(account).At(MkString("free")) = this.SafeString(balance, MkString("available"))
		*(result).At(code) = OpCopy(account)
	}
	return this.ParseBalance(result)
}

func (this *Probit) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"market_id": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("publicGetOrderBook"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	dataBySide := this.GroupBy(data, MkString("side"))
	_ = dataBySide
	return this.ParseOrderBook(dataBySide, symbol, MkUndefined(), MkString("buy"), MkString("sell"), MkString("price"), MkString("quantity"))
}

func (this *Probit) FetchTickers(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{})
	_ = request
	if IsTrue(OpNotEq2(symbols, MkUndefined())) {
		marketIds := this.MarketIds(symbols)
		_ = marketIds
		*(request).At(MkString("market_ids")) = marketIds.Join(MkString(","))
	}
	response := this.Call(MkString("publicGetTicker"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseTickers(data, symbols)
}

func (this *Probit) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"market_ids": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("publicGetTicker"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	ticker := this.SafeValue(data, MkInteger(0))
	_ = ticker
	if IsTrue(OpEq2(ticker, MkUndefined())) {
		panic(NewBadResponse(OpAdd(*this.At(MkString("id")), MkString(" fetchTicker() returned an empty response"))))
	}
	return this.ParseTicker(ticker, market)
}

func (this *Probit) ParseTicker(goArgs ...*Variant) *Variant {
	ticker := GoGetArg(goArgs, 0, MkUndefined())
	_ = ticker
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.Parse8601(this.SafeString(ticker, MkString("time")))
	_ = timestamp
	marketId := this.SafeString(ticker, MkString("market_id"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market, MkString("-"))
	_ = symbol
	close := this.SafeNumber(ticker, MkString("last"))
	_ = close
	change := this.SafeNumber(ticker, MkString("change"))
	_ = change
	percentage := OpCopy(MkUndefined())
	_ = percentage
	open := OpCopy(MkUndefined())
	_ = open
	if IsTrue(OpNotEq2(change, MkUndefined())) {
		if IsTrue(OpNotEq2(close, MkUndefined())) {
			open = OpSub(close, change)
			percentage = OpMulti(OpDiv(change, open), MkInteger(100))
		}
	}
	baseVolume := this.SafeNumber(ticker, MkString("base_volume"))
	_ = baseVolume
	quoteVolume := this.SafeNumber(ticker, MkString("quote_volume"))
	_ = quoteVolume
	vwap := this.Vwap(baseVolume, quoteVolume)
	_ = vwap
	return MkMap(&VarMap{
		"symbol":        symbol,
		"timestamp":     timestamp,
		"datetime":      this.Iso8601(timestamp),
		"high":          this.SafeNumber(ticker, MkString("high")),
		"low":           this.SafeNumber(ticker, MkString("low")),
		"bid":           MkUndefined(),
		"bidVolume":     MkUndefined(),
		"ask":           MkUndefined(),
		"askVolume":     MkUndefined(),
		"vwap":          vwap,
		"open":          open,
		"close":         close,
		"last":          close,
		"previousClose": MkUndefined(),
		"change":        change,
		"percentage":    percentage,
		"average":       MkUndefined(),
		"baseVolume":    baseVolume,
		"quoteVolume":   quoteVolume,
		"info":          ticker,
	})
}

func (this *Probit) FetchMyTrades(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{
		"limit":      MkInteger(100),
		"start_time": this.Iso8601(MkInteger(0)),
		"end_time":   this.Iso8601(this.Milliseconds()),
	})
	_ = request
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
		*(request).At(MkString("market_id")) = *(market).At(MkString("id"))
	}
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("start_time")) = this.Iso8601(since)
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("privateGetTradeHistory"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseTrades(data, market, since, limit)
}

func (this *Probit) FetchTrades(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{
		"market_id":  *(market).At(MkString("id")),
		"limit":      MkInteger(100),
		"start_time": MkString("1970-01-01T00:00:00.000Z"),
		"end_time":   this.Iso8601(this.Milliseconds()),
	})
	_ = request
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("start_time")) = this.Iso8601(since)
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("publicGetTrade"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseTrades(data, market, since, limit)
}

func (this *Probit) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.Parse8601(this.SafeString(trade, MkString("time")))
	_ = timestamp
	id := this.SafeString(trade, MkString("id"))
	_ = id
	marketId := OpCopy(MkUndefined())
	_ = marketId
	if IsTrue(OpNotEq2(id, MkUndefined())) {
		parts := id.Split(MkString(":"))
		_ = parts
		marketId = this.SafeString(parts, MkInteger(0))
	}
	marketId = this.SafeString(trade, MkString("market_id"), marketId)
	symbol := this.SafeSymbol(marketId, market, MkString("-"))
	_ = symbol
	side := this.SafeString(trade, MkString("side"))
	_ = side
	priceString := this.SafeString(trade, MkString("price"))
	_ = priceString
	amountString := this.SafeString(trade, MkString("quantity"))
	_ = amountString
	price := this.ParseNumber(priceString)
	_ = price
	amount := this.ParseNumber(amountString)
	_ = amount
	cost := this.ParseNumber(Precise.StringMul(priceString, amountString))
	_ = cost
	orderId := this.SafeString(trade, MkString("order_id"))
	_ = orderId
	feeCost := this.SafeNumber(trade, MkString("fee_amount"))
	_ = feeCost
	fee := OpCopy(MkUndefined())
	_ = fee
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		feeCurrencyId := this.SafeString(trade, MkString("fee_currency_id"))
		_ = feeCurrencyId
		feeCurrencyCode := this.SafeCurrencyCode(feeCurrencyId)
		_ = feeCurrencyCode
		fee = MkMap(&VarMap{
			"cost":     feeCost,
			"currency": feeCurrencyCode,
		})
	}
	return MkMap(&VarMap{
		"id":           id,
		"info":         trade,
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
		"fee":          fee,
	})
}

func (this *Probit) FetchTime(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetTime"), params)
	_ = response
	timestamp := this.Parse8601(this.SafeString(response, MkString("data")))
	_ = timestamp
	return timestamp
}

func (this *Probit) NormalizeOHLCVTimestamp(goArgs ...*Variant) *Variant {
	timestamp := GoGetArg(goArgs, 0, MkUndefined())
	_ = timestamp
	timeframe := GoGetArg(goArgs, 1, MkUndefined())
	_ = timeframe
	after := GoGetArg(goArgs, 2, MkBool(false))
	_ = after
	duration := this.ParseTimeframe(timeframe)
	_ = duration
	if IsTrue(OpEq2(timeframe, MkString("1M"))) {
		iso8601 := this.Iso8601(timestamp)
		_ = iso8601
		parts := iso8601.Split(MkString("-"))
		_ = parts
		year := this.SafeString(parts, MkInteger(0))
		_ = year
		month := this.SafeInteger(parts, MkInteger(1))
		_ = month
		if IsTrue(after) {
			month = this.Sum(month, MkInteger(1))
		}
		if IsTrue(OpLw(month, MkInteger(10))) {
			month = OpAdd(MkString("0"), month.ToString())
		} else {
			month = month.ToString()
		}
		return OpAdd(year, OpAdd(MkString("-"), OpAdd(month, MkString("-01T00:00:00.000Z"))))
	} else {
		if IsTrue(OpEq2(timeframe, MkString("1w"))) {
			timestamp = parseInt(OpDiv(timestamp, MkInteger(1000)))
			firstSunday := MkInteger(259200)
			_ = firstSunday
			difference := OpSub(timestamp, firstSunday)
			_ = difference
			numWeeks := Math.Floor(OpDiv(difference, duration))
			_ = numWeeks
			previousSunday := this.Sum(firstSunday, OpMulti(numWeeks, duration))
			_ = previousSunday
			if IsTrue(after) {
				previousSunday = this.Sum(previousSunday, duration)
			}
			return this.Iso8601(OpMulti(previousSunday, MkInteger(1000)))
		} else {
			timestamp = parseInt(OpDiv(timestamp, MkInteger(1000)))
			timestamp = OpMulti(duration, parseInt(OpDiv(timestamp, duration)))
			if IsTrue(after) {
				timestamp = this.Sum(timestamp, duration)
			}
			return this.Iso8601(OpMulti(timestamp, MkInteger(1000)))
		}
	}
	return MkUndefined()
}

func (this *Probit) FetchOHLCV(goArgs ...*Variant) *Variant {
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
	interval := *(*this.At(MkString("timeframes"))).At(timeframe)
	_ = interval
	limit = OpTrinary(OpEq2(limit, MkUndefined()), MkInteger(100), limit)
	requestLimit := this.Sum(limit, MkInteger(1))
	_ = requestLimit
	requestLimit = Math.Min(MkInteger(1000), requestLimit)
	request := MkMap(&VarMap{
		"market_ids": *(market).At(MkString("id")),
		"interval":   interval,
		"sort":       MkString("asc"),
		"limit":      requestLimit,
	})
	_ = request
	now := this.Milliseconds()
	_ = now
	duration := this.ParseTimeframe(timeframe)
	_ = duration
	startTime := OpCopy(since)
	_ = startTime
	endTime := OpCopy(now)
	_ = endTime
	if IsTrue(OpEq2(since, MkUndefined())) {
		if IsTrue(OpEq2(limit, MkUndefined())) {
			panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchOHLCV() requires either a since argument or a limit argument"))))
		} else {
			startTime = OpSub(now, OpMulti(limit, OpMulti(duration, MkInteger(1000))))
		}
	} else {
		if IsTrue(OpEq2(limit, MkUndefined())) {
			endTime = OpCopy(now)
		} else {
			endTime = this.Sum(since, OpMulti(this.Sum(limit, MkInteger(1)), OpMulti(duration, MkInteger(1000))))
		}
	}
	startTimeNormalized := this.NormalizeOHLCVTimestamp(startTime, timeframe)
	_ = startTimeNormalized
	endTimeNormalized := this.NormalizeOHLCVTimestamp(endTime, timeframe, MkBool(true))
	_ = endTimeNormalized
	*(request).At(MkString("start_time")) = OpCopy(startTimeNormalized)
	*(request).At(MkString("end_time")) = OpCopy(endTimeNormalized)
	response := this.Call(MkString("publicGetCandle"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseOHLCVs(data, market, timeframe, since, limit)
}

func (this *Probit) ParseOHLCV(goArgs ...*Variant) *Variant {
	ohlcv := GoGetArg(goArgs, 0, MkUndefined())
	_ = ohlcv
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	return MkArray(&VarArray{
		this.Parse8601(this.SafeString(ohlcv, MkString("start_time"))),
		this.SafeNumber(ohlcv, MkString("open")),
		this.SafeNumber(ohlcv, MkString("high")),
		this.SafeNumber(ohlcv, MkString("low")),
		this.SafeNumber(ohlcv, MkString("close")),
		this.SafeNumber(ohlcv, MkString("base_volume")),
	})
}

func (this *Probit) FetchOpenOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	since = this.Parse8601(since)
	request := MkMap(&VarMap{})
	_ = request
	market := OpCopy(MkUndefined())
	_ = market
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
		*(request).At(MkString("market_id")) = *(market).At(MkString("id"))
	}
	response := this.Call(MkString("privateGetOpenOrder"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"))
	_ = data
	return this.ParseOrders(data, market, since, limit)
}

func (this *Probit) FetchClosedOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{
		"start_time": this.Iso8601(MkInteger(0)),
		"end_time":   this.Iso8601(this.Milliseconds()),
		"limit":      MkInteger(100),
	})
	_ = request
	market := OpCopy(MkUndefined())
	_ = market
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
		*(request).At(MkString("market_id")) = *(market).At(MkString("id"))
	}
	if IsTrue(since) {
		*(request).At(MkString("start_time")) = this.Iso8601(since)
	}
	if IsTrue(limit) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("privateGetOrderHistory"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"))
	_ = data
	return this.ParseOrders(data, market, since, limit)
}

func (this *Probit) FetchOrder(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"market_id": *(market).At(MkString("id"))})
	_ = request
	clientOrderId := this.SafeString2(params, MkString("clientOrderId"), MkString("client_order_id"))
	_ = clientOrderId
	if IsTrue(OpNotEq2(clientOrderId, MkUndefined())) {
		*(request).At(MkString("client_order_id")) = OpCopy(clientOrderId)
	} else {
		*(request).At(MkString("order_id")) = OpCopy(id)
	}
	query := this.Omit(params, MkArray(&VarArray{
		MkString("clientOrderId"),
		MkString("client_order_id"),
	}))
	_ = query
	response := this.Call(MkString("privateGetOrder"), this.Extend(request, query))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	order := this.SafeValue(data, MkInteger(0))
	_ = order
	return this.ParseOrder(order, market)
}

func (this *Probit) ParseOrderStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"open":      MkString("open"),
		"cancelled": MkString("canceled"),
		"filled":    MkString("closed"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Probit) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	status := this.ParseOrderStatus(this.SafeString(order, MkString("status")))
	_ = status
	id := this.SafeString(order, MkString("id"))
	_ = id
	type_ := this.SafeString(order, MkString("type"))
	_ = type_
	side := this.SafeString(order, MkString("side"))
	_ = side
	marketId := this.SafeString(order, MkString("market_id"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market, MkString("-"))
	_ = symbol
	timestamp := this.Parse8601(this.SafeString(order, MkString("time")))
	_ = timestamp
	price := this.SafeNumber(order, MkString("limit_price"))
	_ = price
	filled := this.SafeNumber(order, MkString("filled_quantity"))
	_ = filled
	remaining := this.SafeNumber(order, MkString("open_quantity"))
	_ = remaining
	canceledAmount := this.SafeNumber(order, MkString("cancelled_quantity"))
	_ = canceledAmount
	if IsTrue(OpNotEq2(canceledAmount, MkUndefined())) {
		remaining = this.Sum(remaining, canceledAmount)
	}
	amount := this.SafeNumber(order, MkString("quantity"), this.Sum(filled, remaining))
	_ = amount
	cost := this.SafeNumber2(order, MkString("filled_cost"), MkString("cost"))
	_ = cost
	if IsTrue(OpEq2(type_, MkString("market"))) {
		price = OpCopy(MkUndefined())
	}
	clientOrderId := this.SafeString(order, MkString("client_order_id"))
	_ = clientOrderId
	if IsTrue(OpEq2(clientOrderId, MkString(""))) {
		clientOrderId = OpCopy(MkUndefined())
	}
	timeInForce := this.SafeStringUpper(order, MkString("time_in_force"))
	_ = timeInForce
	return this.SafeOrder(MkMap(&VarMap{
		"id":                 id,
		"info":               order,
		"clientOrderId":      clientOrderId,
		"timestamp":          timestamp,
		"datetime":           this.Iso8601(timestamp),
		"lastTradeTimestamp": MkUndefined(),
		"symbol":             symbol,
		"type":               type_,
		"timeInForce":        timeInForce,
		"side":               side,
		"status":             status,
		"price":              price,
		"stopPrice":          MkUndefined(),
		"amount":             amount,
		"filled":             filled,
		"remaining":          remaining,
		"average":            MkUndefined(),
		"cost":               cost,
		"fee":                MkUndefined(),
		"trades":             MkUndefined(),
	}))
}

func (this *Probit) CostToPrecision(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	cost := GoGetArg(goArgs, 1, MkUndefined())
	_ = cost
	return this.DecimalToPrecision(cost, TRUNCATE, *(*(*(*this.At(MkString("markets"))).At(symbol)).At(MkString("precision"))).At(MkString("cost")), *this.At(MkString("precisionMode")))
}

func (this *Probit) CreateOrder(goArgs ...*Variant) *Variant {
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
	options := this.SafeValue(*this.At(MkString("options")), MkString("timeInForce"))
	_ = options
	defaultTimeInForce := this.SafeValue(options, type_)
	_ = defaultTimeInForce
	timeInForce := this.SafeString2(params, MkString("timeInForce"), MkString("time_in_force"), defaultTimeInForce)
	_ = timeInForce
	request := MkMap(&VarMap{
		"market_id":     *(market).At(MkString("id")),
		"type":          type_,
		"side":          side,
		"time_in_force": timeInForce,
	})
	_ = request
	clientOrderId := this.SafeString2(params, MkString("clientOrderId"), MkString("client_order_id"))
	_ = clientOrderId
	if IsTrue(OpNotEq2(clientOrderId, MkUndefined())) {
		*(request).At(MkString("client_order_id")) = OpCopy(clientOrderId)
	}
	costToPrecision := OpCopy(MkUndefined())
	_ = costToPrecision
	if IsTrue(OpEq2(type_, MkString("limit"))) {
		*(request).At(MkString("limit_price")) = this.PriceToPrecision(symbol, price)
		*(request).At(MkString("quantity")) = this.AmountToPrecision(symbol, amount)
	} else {
		if IsTrue(OpEq2(type_, MkString("market"))) {
			if IsTrue(OpEq2(side, MkString("buy"))) {
				cost := this.SafeNumber(params, MkString("cost"))
				_ = cost
				createMarketBuyOrderRequiresPrice := this.SafeValue(*this.At(MkString("options")), MkString("createMarketBuyOrderRequiresPrice"), MkBool(true))
				_ = createMarketBuyOrderRequiresPrice
				if IsTrue(createMarketBuyOrderRequiresPrice) {
					if IsTrue(OpNotEq2(price, MkUndefined())) {
						if IsTrue(OpEq2(cost, MkUndefined())) {
							cost = OpMulti(amount, price)
						}
					} else {
						if IsTrue(OpEq2(cost, MkUndefined())) {
							panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")), MkString(" createOrder() requires the price argument for market buy orders to calculate total order cost (amount to spend), where cost = amount * price. Supply a price argument to createOrder() call if you want the cost to be calculated for you from price and amount, or, alternatively, add .options['createMarketBuyOrderRequiresPrice'] = false and supply the total cost value in the 'amount' argument or in the 'cost' extra parameter (the exchange-specific behaviour)"))))
						}
					}
				} else {
					cost = OpTrinary(OpEq2(cost, MkUndefined()), amount, cost)
				}
				costToPrecision = this.CostToPrecision(symbol, cost)
				*(request).At(MkString("cost")) = OpCopy(costToPrecision)
			} else {
				*(request).At(MkString("quantity")) = this.AmountToPrecision(symbol, amount)
			}
		}
	}
	query := this.Omit(params, MkArray(&VarArray{
		MkString("timeInForce"),
		MkString("time_in_force"),
		MkString("clientOrderId"),
		MkString("client_order_id"),
	}))
	_ = query
	response := this.Call(MkString("privatePostNewOrder"), this.Extend(request, query))
	_ = response
	data := this.SafeValue(response, MkString("data"))
	_ = data
	order := this.ParseOrder(data, market)
	_ = order
	if IsTrue(OpAnd(OpEq2(type_, MkString("market")), OpEq2(side, MkString("buy")))) {
		*(order).At(MkString("amount")) = OpCopy(MkUndefined())
		*(order).At(MkString("cost")) = this.ParseNumber(costToPrecision)
		*(order).At(MkString("remaining")) = OpCopy(MkUndefined())
	}
	return order
}

func (this *Probit) CancelOrder(goArgs ...*Variant) *Variant {
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
		"market_id": *(market).At(MkString("id")),
		"order_id":  id,
	})
	_ = request
	response := this.Call(MkString("privatePostCancelOrder"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"))
	_ = data
	return this.ParseOrder(data)
}

func (this *Probit) ParseDepositAddress(goArgs ...*Variant) *Variant {
	depositAddress := GoGetArg(goArgs, 0, MkUndefined())
	_ = depositAddress
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	address := this.SafeString(depositAddress, MkString("address"))
	_ = address
	tag := this.SafeString(depositAddress, MkString("destination_tag"))
	_ = tag
	currencyId := this.SafeString(depositAddress, MkString("currency_id"))
	_ = currencyId
	code := this.SafeCurrencyCode(currencyId)
	_ = code
	this.CheckAddress(address)
	return MkMap(&VarMap{
		"currency": code,
		"address":  address,
		"tag":      tag,
		"info":     depositAddress,
	})
}

func (this *Probit) FetchDepositAddress(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{"currency_id": *(currency).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("privateGetDepositAddress"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	firstAddress := this.SafeValue(data, MkInteger(0))
	_ = firstAddress
	if IsTrue(OpEq2(firstAddress, MkUndefined())) {
		panic(NewInvalidAddress(OpAdd(*this.At(MkString("id")), MkString(" fetchDepositAddress returned an empty response"))))
	}
	return this.ParseDepositAddress(firstAddress, currency)
}

func (this *Probit) FetchDepositAddresses(goArgs ...*Variant) *Variant {
	codes := GoGetArg(goArgs, 0, MkUndefined())
	_ = codes
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{})
	_ = request
	if IsTrue(codes) {
		currencyIds := MkArray(&VarArray{})
		_ = currencyIds
		for i := MkInteger(0); IsTrue(OpLw(i, codes.Length)); OpIncr(&i) {
			currency := this.Currency(*(codes).At(i))
			_ = currency
			currencyIds.Push(*(currency).At(MkString("id")))
		}
		*(request).At(MkString("currency_id")) = codes.Join(MkString(","))
	}
	response := this.Call(MkString("privateGetDepositAddress"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseDepositAddresses(data)
}

func (this *Probit) Withdraw(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpEq2(tag, MkUndefined())) {
		tag = MkString("")
	}
	request := MkMap(&VarMap{
		"currency_id":     *(currency).At(MkString("id")),
		"address":         address,
		"destination_tag": tag,
		"amount":          this.NumberToString(amount),
	})
	_ = request
	response := this.Call(MkString("privatePostWithdrawal"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"))
	_ = data
	return this.ParseTransaction(data, currency)
}

func (this *Probit) ParseTransaction(goArgs ...*Variant) *Variant {
	transaction := GoGetArg(goArgs, 0, MkUndefined())
	_ = transaction
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	id := this.SafeString(transaction, MkString("id"))
	_ = id
	amount := this.SafeNumber(transaction, MkString("amount"))
	_ = amount
	address := this.SafeString(transaction, MkString("address"))
	_ = address
	tag := this.SafeString(transaction, MkString("destination_tag"))
	_ = tag
	txid := this.SafeString(transaction, MkString("hash"))
	_ = txid
	timestamp := this.Parse8601(this.SafeString(transaction, MkString("time")))
	_ = timestamp
	type_ := this.SafeString(transaction, MkString("type"))
	_ = type_
	currencyId := this.SafeString(transaction, MkString("currency_id"))
	_ = currencyId
	code := this.SafeCurrencyCode(currencyId)
	_ = code
	status := this.ParseTransactionStatus(this.SafeString(transaction, MkString("status")))
	_ = status
	feeCost := this.SafeNumber(transaction, MkString("fee"))
	_ = feeCost
	fee := OpCopy(MkUndefined())
	_ = fee
	if IsTrue(OpAnd(OpNotEq2(feeCost, MkUndefined()), OpNotEq2(feeCost, MkInteger(0)))) {
		fee = MkMap(&VarMap{
			"currency": code,
			"cost":     feeCost,
		})
	}
	return MkMap(&VarMap{
		"id":          id,
		"currency":    code,
		"amount":      amount,
		"addressFrom": MkUndefined(),
		"address":     address,
		"addressTo":   address,
		"tagFrom":     MkUndefined(),
		"tag":         tag,
		"tagTo":       tag,
		"status":      status,
		"type":        type_,
		"txid":        txid,
		"timestamp":   timestamp,
		"datetime":    this.Iso8601(timestamp),
		"fee":         fee,
		"info":        transaction,
	})
}

func (this *Probit) ParseTransactionStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"requested":  MkString("pending"),
		"pending":    MkString("pending"),
		"confirming": MkString("pending"),
		"confirmed":  MkString("pending"),
		"applying":   MkString("pending"),
		"done":       MkString("ok"),
		"cancelled":  MkString("canceled"),
		"cancelling": MkString("canceled"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Probit) Nonce(goArgs ...*Variant) *Variant {
	return this.Milliseconds()
}

func (this *Probit) Sign(goArgs ...*Variant) *Variant {
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
	url := OpAdd(*(*(*this.At(MkString("urls"))).At(MkString("api"))).At(api), MkString("/"))
	_ = url
	query := this.Omit(params, this.ExtractParams(path))
	_ = query
	if IsTrue(OpEq2(api, MkString("accounts"))) {
		this.CheckRequiredCredentials()
		OpAddAssign(&url, this.ImplodeParams(path, params))
		auth := OpAdd(*this.At(MkString("apiKey")), OpAdd(MkString(":"), *this.At(MkString("secret"))))
		_ = auth
		auth64 := this.StringToBase64(auth)
		_ = auth64
		headers = MkMap(&VarMap{
			"Authorization": OpAdd(MkString("Basic "), this.Decode(auth64)),
			"Content-Type":  MkString("application/json"),
		})
		if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
			body = this.Json(query)
		}
	} else {
		OpAddAssign(&url, OpAdd(*this.At(MkString("version")), MkString("/")))
		if IsTrue(OpEq2(api, MkString("public"))) {
			OpAddAssign(&url, this.ImplodeParams(path, params))
			if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
				OpAddAssign(&url, OpAdd(MkString("?"), this.Urlencode(query)))
			}
		} else {
			if IsTrue(OpEq2(api, MkString("private"))) {
				now := this.Milliseconds()
				_ = now
				this.CheckRequiredCredentials()
				expires := this.SafeInteger(*this.At(MkString("options")), MkString("expires"))
				_ = expires
				if IsTrue(OpOr(OpEq2(expires, MkUndefined()), OpLw(expires, now))) {
					panic(NewAuthenticationError(OpAdd(*this.At(MkString("id")), MkString(" access token expired, call signIn() method"))))
				}
				accessToken := this.SafeString(*this.At(MkString("options")), MkString("accessToken"))
				_ = accessToken
				headers = MkMap(&VarMap{"Authorization": OpAdd(MkString("Bearer "), accessToken)})
				OpAddAssign(&url, this.ImplodeParams(path, params))
				if IsTrue(OpEq2(method, MkString("GET"))) {
					if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
						OpAddAssign(&url, OpAdd(MkString("?"), this.Urlencode(query)))
					}
				} else {
					if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
						body = this.Json(query)
						*(headers).At(MkString("Content-Type")) = MkString("application/json")
					}
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

func (this *Probit) SignIn(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.CheckRequiredCredentials()
	request := MkMap(&VarMap{"grant_type": MkString("client_credentials")})
	_ = request
	response := this.Call(MkString("accountsPostToken"), this.Extend(request, params))
	_ = response
	expiresIn := this.SafeInteger(response, MkString("expires_in"))
	_ = expiresIn
	accessToken := this.SafeString(response, MkString("access_token"))
	_ = accessToken
	*(*this.At(MkString("options"))).At(MkString("accessToken")) = OpCopy(accessToken)
	*(*this.At(MkString("options"))).At(MkString("expires")) = this.Sum(this.Milliseconds(), OpMulti(expiresIn, MkInteger(1000)))
	return response
}

func (this *Probit) HandleErrors(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpEq2(response, MkUndefined())) {
		return MkUndefined()
	}
	if IsTrue(OpHasMember(MkString("errorCode"), response)) {
		errorCode := this.SafeString(response, MkString("errorCode"))
		_ = errorCode
		message := this.SafeString(response, MkString("message"))
		_ = message
		if IsTrue(OpNotEq2(errorCode, MkUndefined())) {
			feedback := OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))
			_ = feedback
			this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("exact")), message, feedback)
			this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("exact")), errorCode, feedback)
			panic(NewExchangeError(feedback))
		}
	}
	return MkUndefined()
}
