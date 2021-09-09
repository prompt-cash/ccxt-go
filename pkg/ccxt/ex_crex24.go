package ccxt

import ()

type Crex24 struct {
	*ExchangeBase
}

var _ Exchange = (*Crex24)(nil)

func init() {
	exchange := &Crex24{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Crex24) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("crex24"),
		"name": MkString("CREX24"),
		"countries": MkArray(&VarArray{
			MkString("EE"),
		}),
		"rateLimit": MkInteger(500),
		"version":   MkString("v2"),
		"has": MkMap(&VarMap{
			"cancelAllOrders":     MkBool(true),
			"cancelOrder":         MkBool(true),
			"CORS":                MkBool(false),
			"createOrder":         MkBool(true),
			"editOrder":           MkBool(true),
			"fetchBalance":        MkBool(true),
			"fetchBidsAsks":       MkBool(true),
			"fetchClosedOrders":   MkBool(true),
			"fetchCurrencies":     MkBool(true),
			"fetchDepositAddress": MkBool(true),
			"fetchDeposits":       MkBool(true),
			"fetchMarkets":        MkBool(true),
			"fetchMyTrades":       MkBool(true),
			"fetchOHLCV":          MkBool(true),
			"fetchOpenOrders":     MkBool(true),
			"fetchOrder":          MkBool(true),
			"fetchOrderBook":      MkBool(true),
			"fetchOrders":         MkBool(true),
			"fetchOrderTrades":    MkBool(true),
			"fetchTicker":         MkBool(true),
			"fetchTickers":        MkBool(true),
			"fetchTrades":         MkBool(true),
			"fetchTradingFee":     MkBool(false),
			"fetchTradingFees":    MkBool(false),
			"fetchFundingFees":    MkBool(true),
			"fetchTransactions":   MkBool(true),
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
			"4h":  MkString("4h"),
			"1d":  MkString("1d"),
			"1w":  MkString("1w"),
			"1M":  MkString("1mo"),
		}),
		"urls": MkMap(&VarMap{
			"logo":     MkString("https://user-images.githubusercontent.com/1294454/47813922-6f12cc00-dd5d-11e8-97c6-70f957712d47.jpg"),
			"api":      MkString("https://api.crex24.com"),
			"www":      MkString("https://crex24.com"),
			"referral": MkString("https://crex24.com/?refid=slxsjsjtil8xexl9hksr"),
			"doc":      MkString("https://docs.crex24.com/trade-api/v2"),
			"fees":     MkString("https://crex24.com/fees"),
		}),
		"api": MkMap(&VarMap{
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("currencies"),
				MkString("instruments"),
				MkString("tickers"),
				MkString("recentTrades"),
				MkString("orderBook"),
				MkString("ohlcv"),
				MkString("tradingFeeSchedules"),
				MkString("withdrawalFees"),
				MkString("currencyTransport"),
				MkString("currenciesWithdrawalFees"),
			})}),
			"trading": MkMap(&VarMap{
				"get": MkArray(&VarArray{
					MkString("orderStatus"),
					MkString("orderTrades"),
					MkString("activeOrders"),
					MkString("orderHistory"),
					MkString("tradeHistory"),
					MkString("tradingFee"),
					MkString("tradeFee"),
				}),
				"post": MkArray(&VarArray{
					MkString("placeOrder"),
					MkString("modifyOrder"),
					MkString("cancelOrdersById"),
					MkString("cancelOrdersByInstrument"),
					MkString("cancelAllOrders"),
				}),
			}),
			"account": MkMap(&VarMap{
				"get": MkArray(&VarArray{
					MkString("balance"),
					MkString("depositAddress"),
					MkString("moneyTransfers"),
					MkString("moneyTransferStatus"),
					MkString("previewWithdrawal"),
				}),
				"post": MkArray(&VarArray{
					MkString("withdraw"),
				}),
			}),
		}),
		"precisionMode": TICK_SIZE,
		"fees": MkMap(&VarMap{
			"trading": MkMap(&VarMap{
				"tierBased":  MkBool(true),
				"percentage": MkBool(true),
				"taker":      MkNumber(0.001),
				"maker":      OpNeg(MkNumber(0.0001)),
			}),
			"funding": MkMap(&VarMap{
				"tierBased":  MkBool(false),
				"percentage": MkBool(false),
				"withdraw":   MkMap(&VarMap{}),
				"deposit":    MkMap(&VarMap{}),
			}),
		}),
		"commonCurrencies": MkMap(&VarMap{
			"ACM":    MkString("Actinium"),
			"BCC":    MkString("BCH"),
			"BIT":    MkString("BitMoney"),
			"BULL":   MkString("BuySell"),
			"CREDIT": MkString("TerraCredit"),
			"EPS":    MkString("Epanus"),
			"FUND":   MkString("FUNDChains"),
			"GHOST":  MkString("GHOSTPRISM"),
			"GTC":    MkString("GastroCoin"),
			"IQ":     MkString("IQ.Cash"),
			"PUT":    MkString("PutinCoin"),
			"SBTC":   MkString("SBTCT"),
			"UNI":    MkString("Universe"),
			"YOYO":   MkString("YOYOW"),
		}),
		"options": MkMap(&VarMap{
			"fetchOrdersMethod":                   MkString("tradingGetOrderHistory"),
			"fetchClosedOrdersMethod":             MkString("tradingGetOrderHistory"),
			"fetchTickersMethod":                  MkString("publicGetTicker24hr"),
			"defaultTimeInForce":                  MkString("GTC"),
			"hasAlreadyAuthenticatedSuccessfully": MkBool(false),
			"warnOnFetchOpenOrdersWithoutSymbol":  MkBool(true),
			"parseOrderToPrecision":               MkBool(false),
			"newOrderRespType":                    MkString("RESULT"),
		}),
		"exceptions": MkMap(&VarMap{
			"exact": MkMap(&VarMap{
				"Parameter 'filter' contains invalid value.":                                                    BadRequest,
				"Mandatory parameter 'instrument' is missing.":                                                  BadRequest,
				"The value of parameter 'till' must be greater than or equal to the value of parameter 'from'.": BadRequest,
				"Failed to verify request signature.":                                                           AuthenticationError,
				"Nonce error. Make sure that the value passed in the 'X-CREX24-API-NONCE' header is greater in each consecutive request than in the previous one for the corresponding API-Key provided in 'X-CREX24-API-KEY' header.": InvalidNonce,
				"Market orders are not supported by the instrument currently.": InvalidOrder,
				"Parameter 'instrument' contains invalid value.":               BadSymbol,
			}),
			"broad": MkMap(&VarMap{
				"try again later":         ExchangeNotAvailable,
				"API Key":                 AuthenticationError,
				"Insufficient funds":      InsufficientFunds,
				"has been delisted.":      BadSymbol,
				"is currently suspended.": BadSymbol,
				"Mandatory parameter":     BadRequest,
			}),
		}),
	}))
}

func (this *Crex24) Nonce(goArgs ...*Variant) *Variant {
	return this.Milliseconds()
}

func (this *Crex24) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetInstruments"), params)
	_ = response
	response2 := this.Call(MkString("publicGetTradingFeeSchedules"), params)
	_ = response2
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
		market := *(response).At(i)
		_ = market
		id := this.SafeString(market, MkString("symbol"))
		_ = id
		baseId := this.SafeString(market, MkString("baseCurrency"))
		_ = baseId
		quoteId := this.SafeString(market, MkString("quoteCurrency"))
		_ = quoteId
		base := this.SafeCurrencyCode(baseId)
		_ = base
		quote := this.SafeCurrencyCode(quoteId)
		_ = quote
		symbol := OpAdd(base, OpAdd(MkString("/"), quote))
		_ = symbol
		tickSize := this.SafeNumber(market, MkString("tickSize"))
		_ = tickSize
		minPrice := this.SafeNumber(market, MkString("minPrice"))
		_ = minPrice
		maxPrice := this.SafeNumber(market, MkString("maxPrice"))
		_ = maxPrice
		minAmount := this.SafeNumber(market, MkString("minVolume"))
		_ = minAmount
		maxAmount := this.SafeNumber(market, MkString("maxVolume"))
		_ = maxAmount
		minCost := this.SafeNumber(market, MkString("minQuoteVolume"))
		_ = minCost
		maxCost := this.SafeNumber(market, MkString("maxQuoteVolume"))
		_ = maxCost
		precision := MkMap(&VarMap{
			"amount": minAmount,
			"price":  tickSize,
		})
		_ = precision
		maker := OpCopy(MkUndefined())
		_ = maker
		taker := OpCopy(MkUndefined())
		_ = taker
		feeSchedule := this.SafeString(market, MkString("feeSchedule"))
		_ = feeSchedule
		for j := MkInteger(0); IsTrue(OpLw(j, response2.Length)); OpIncr(&j) {
			feeScheduleName := this.SafeString(*(response2).At(j), MkString("name"))
			_ = feeScheduleName
			if IsTrue(OpEq2(feeScheduleName, feeSchedule)) {
				feeRates := this.SafeValue(*(response2).At(j), MkString("feeRates"), MkArray(&VarArray{}))
				_ = feeRates
				for k := MkInteger(0); IsTrue(OpLw(k, feeRates.Length)); OpIncr(&k) {
					volumeThreshold := this.SafeNumber(*(feeRates).At(k), MkString("volumeThreshold"))
					_ = volumeThreshold
					if IsTrue(OpEq2(volumeThreshold, MkInteger(0))) {
						maker = this.SafeNumber(*(feeRates).At(k), MkString("maker"))
						taker = this.SafeNumber(*(feeRates).At(k), MkString("taker"))
						break
					}
				}
				break
			}
		}
		active := OpEq2(*(market).At(MkString("state")), MkString("active"))
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
			"maker":     maker,
			"taker":     taker,
			"limits": MkMap(&VarMap{
				"amount": MkMap(&VarMap{
					"min": minAmount,
					"max": maxAmount,
				}),
				"price": MkMap(&VarMap{
					"min": minPrice,
					"max": maxPrice,
				}),
				"cost": MkMap(&VarMap{
					"min": minCost,
					"max": maxCost,
				}),
			}),
		}))
	}
	return result
}

func (this *Crex24) FetchCurrencies(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetCurrencies"), params)
	_ = response
	result := MkMap(&VarMap{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
		currency := *(response).At(i)
		_ = currency
		id := this.SafeString(currency, MkString("symbol"))
		_ = id
		code := this.SafeCurrencyCode(id)
		_ = code
		withdrawalPrecision := this.SafeInteger(currency, MkString("withdrawalPrecision"))
		_ = withdrawalPrecision
		precision := Math.Pow(MkInteger(10), OpNeg(withdrawalPrecision))
		_ = precision
		address := this.SafeValue(currency, MkString("BaseAddress"))
		_ = address
		active := OpAnd(*(currency).At(MkString("depositsAllowed")), OpAnd(*(currency).At(MkString("withdrawalsAllowed")), OpNot(*(currency).At(MkString("isDelisted")))))
		_ = active
		type_ := OpTrinary(*(currency).At(MkString("isFiat")), MkString("fiat"), MkString("crypto"))
		_ = type_
		*(result).At(code) = MkMap(&VarMap{
			"id":        id,
			"code":      code,
			"address":   address,
			"info":      currency,
			"type":      type_,
			"name":      this.SafeString(currency, MkString("name")),
			"active":    active,
			"fee":       this.SafeNumber(currency, MkString("flatWithdrawalFee")),
			"precision": precision,
			"limits": MkMap(&VarMap{
				"amount": MkMap(&VarMap{
					"min": Math.Pow(MkInteger(10), OpNeg(precision)),
					"max": Math.Pow(MkInteger(10), precision),
				}),
				"deposit": MkMap(&VarMap{
					"min": this.SafeNumber(currency, MkString("minDeposit")),
					"max": MkUndefined(),
				}),
				"withdraw": MkMap(&VarMap{
					"min": this.SafeNumber(currency, MkString("minWithdrawal")),
					"max": this.SafeNumber(currency, MkString("maxWithdrawal")),
				}),
			}),
		})
	}
	return result
}

func (this *Crex24) FetchFundingFees(goArgs ...*Variant) *Variant {
	codes := GoGetArg(goArgs, 0, MkUndefined())
	_ = codes
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("publicGetCurrenciesWithdrawalFees"), params)
	_ = response
	withdrawFees := MkMap(&VarMap{})
	_ = withdrawFees
	for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
		entry := *(response).At(i)
		_ = entry
		currencyId := this.SafeString(entry, MkString("currency"))
		_ = currencyId
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		networkList := this.SafeValue(entry, MkString("fees"))
		_ = networkList
		*(withdrawFees).At(code) = MkMap(&VarMap{})
		for j := MkInteger(0); IsTrue(OpLw(j, networkList.Length)); OpIncr(&j) {
			networkEntry := *(networkList).At(j)
			_ = networkEntry
			networkId := this.SafeString(networkEntry, MkString("feeCurrency"))
			_ = networkId
			networkCode := this.SafeCurrencyCode(networkId)
			_ = networkCode
			fee := this.SafeNumber(networkEntry, MkString("amount"))
			_ = fee
			*(*(withdrawFees).At(code)).At(networkCode) = OpCopy(fee)
		}
	}
	return MkMap(&VarMap{
		"withdraw": withdrawFees,
		"deposit":  MkMap(&VarMap{}),
		"info":     response,
	})
}

func (this *Crex24) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{})
	_ = request
	response := this.Call(MkString("accountGetBalance"), this.Extend(request, params))
	_ = response
	result := MkMap(&VarMap{"info": response})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
		balance := *(response).At(i)
		_ = balance
		currencyId := this.SafeString(balance, MkString("currency"))
		_ = currencyId
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		account := this.Account()
		_ = account
		*(account).At(MkString("free")) = this.SafeString(balance, MkString("available"))
		*(account).At(MkString("used")) = this.SafeString(balance, MkString("reserved"))
		*(result).At(code) = OpCopy(account)
	}
	return this.ParseBalance(result)
}

func (this *Crex24) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"instrument": *(market).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("publicGetOrderBook"), this.Extend(request, params))
	_ = response
	return this.ParseOrderBook(response, symbol, MkUndefined(), MkString("buyLevels"), MkString("sellLevels"), MkString("price"), MkString("volume"))
}

func (this *Crex24) ParseTicker(goArgs ...*Variant) *Variant {
	ticker := GoGetArg(goArgs, 0, MkUndefined())
	_ = ticker
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.Parse8601(this.SafeString(ticker, MkString("timestamp")))
	_ = timestamp
	marketId := this.SafeString(ticker, MkString("instrument"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market, MkString("-"))
	_ = symbol
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
		"vwap":          MkUndefined(),
		"open":          MkUndefined(),
		"close":         last,
		"last":          last,
		"previousClose": MkUndefined(),
		"change":        MkUndefined(),
		"percentage":    this.SafeNumber(ticker, MkString("percentChange")),
		"average":       MkUndefined(),
		"baseVolume":    this.SafeNumber(ticker, MkString("baseVolume")),
		"quoteVolume":   this.SafeNumber(ticker, MkString("quoteVolume")),
		"info":          ticker,
	})
}

func (this *Crex24) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"instrument": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("publicGetTickers"), this.Extend(request, params))
	_ = response
	numTickers := OpCopy(response.Length)
	_ = numTickers
	if IsTrue(OpLw(numTickers, MkInteger(1))) {
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" fetchTicker could not load quotes for symbol "), symbol))))
	}
	return this.ParseTicker(*(response).At(MkInteger(0)), market)
}

func (this *Crex24) FetchTickers(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{})
	_ = request
	if IsTrue(OpNotEq2(symbols, MkUndefined())) {
		ids := this.MarketIds(symbols)
		_ = ids
		*(request).At(MkString("instrument")) = ids.Join(MkString(","))
	}
	response := this.Call(MkString("publicGetTickers"), this.Extend(request, params))
	_ = response
	return this.ParseTickers(response, symbols)
}

func (this *Crex24) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.Parse8601(this.SafeString(trade, MkString("timestamp")))
	_ = timestamp
	priceString := this.SafeString(trade, MkString("price"))
	_ = priceString
	amountString := this.SafeString(trade, MkString("volume"))
	_ = amountString
	price := this.ParseNumber(priceString)
	_ = price
	amount := this.ParseNumber(amountString)
	_ = amount
	cost := this.ParseNumber(Precise.StringMul(priceString, amountString))
	_ = cost
	id := this.SafeString(trade, MkString("id"))
	_ = id
	side := this.SafeString(trade, MkString("side"))
	_ = side
	orderId := this.SafeString(trade, MkString("orderId"))
	_ = orderId
	marketId := this.SafeString(trade, MkString("instrument"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market, MkString("-"))
	_ = symbol
	fee := OpCopy(MkUndefined())
	_ = fee
	feeCurrencyId := this.SafeString(trade, MkString("feeCurrency"))
	_ = feeCurrencyId
	feeCode := this.SafeCurrencyCode(feeCurrencyId)
	_ = feeCode
	feeCost := this.SafeNumber(trade, MkString("fee"))
	_ = feeCost
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		fee = MkMap(&VarMap{
			"cost":     feeCost,
			"currency": feeCode,
		})
	}
	takerOrMaker := OpCopy(MkUndefined())
	_ = takerOrMaker
	return MkMap(&VarMap{
		"info":         trade,
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       symbol,
		"id":           id,
		"order":        orderId,
		"type":         MkUndefined(),
		"takerOrMaker": takerOrMaker,
		"side":         side,
		"price":        price,
		"cost":         cost,
		"amount":       amount,
		"fee":          fee,
	})
}

func (this *Crex24) FetchTrades(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"instrument": *(market).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("publicGetRecentTrades"), this.Extend(request, params))
	_ = response
	return this.ParseTrades(response, market, since, limit)
}

func (this *Crex24) ParseOHLCV(goArgs ...*Variant) *Variant {
	ohlcv := GoGetArg(goArgs, 0, MkUndefined())
	_ = ohlcv
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	return MkArray(&VarArray{
		this.Parse8601(this.SafeString(ohlcv, MkString("timestamp"))),
		this.SafeNumber(ohlcv, MkString("open")),
		this.SafeNumber(ohlcv, MkString("high")),
		this.SafeNumber(ohlcv, MkString("low")),
		this.SafeNumber(ohlcv, MkString("close")),
		this.SafeNumber(ohlcv, MkString("volume")),
	})
}

func (this *Crex24) FetchOHLCV(goArgs ...*Variant) *Variant {
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
		"granularity": *(*this.At(MkString("timeframes"))).At(timeframe),
		"instrument":  *(market).At(MkString("id")),
	})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("publicGetOhlcv"), this.Extend(request, params))
	_ = response
	return this.ParseOHLCVs(response, market, timeframe, since, limit)
}

func (this *Crex24) ParseOrderStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"submitting":               MkString("open"),
		"unfilledActive":           MkString("open"),
		"partiallyFilledActive":    MkString("open"),
		"filled":                   MkString("closed"),
		"partiallyFilledCancelled": MkString("canceled"),
		"unfilledCancelled":        MkString("canceled"),
	})
	_ = statuses
	return OpTrinary(OpHasMember(status, statuses), *(statuses).At(status), status)
}

func (this *Crex24) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	status := this.ParseOrderStatus(this.SafeString(order, MkString("status")))
	_ = status
	marketId := this.SafeString(order, MkString("instrument"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market, MkString("-"))
	_ = symbol
	timestamp := this.Parse8601(this.SafeString(order, MkString("timestamp")))
	_ = timestamp
	price := this.SafeNumber(order, MkString("price"))
	_ = price
	amount := this.SafeNumber(order, MkString("volume"))
	_ = amount
	remaining := this.SafeNumber(order, MkString("remainingVolume"))
	_ = remaining
	lastTradeTimestamp := this.Parse8601(this.SafeString(order, MkString("lastUpdate")))
	_ = lastTradeTimestamp
	id := this.SafeString(order, MkString("id"))
	_ = id
	type_ := this.SafeString(order, MkString("type"))
	_ = type_
	side := this.SafeString(order, MkString("side"))
	_ = side
	fee := OpCopy(MkUndefined())
	_ = fee
	trades := OpCopy(MkUndefined())
	_ = trades
	average := OpCopy(MkUndefined())
	_ = average
	timeInForce := this.SafeString(order, MkString("timeInForce"))
	_ = timeInForce
	stopPrice := this.SafeNumber(order, MkString("stopPrice"))
	_ = stopPrice
	return this.SafeOrder(MkMap(&VarMap{
		"info":               order,
		"id":                 id,
		"clientOrderId":      MkUndefined(),
		"timestamp":          timestamp,
		"datetime":           this.Iso8601(timestamp),
		"lastTradeTimestamp": lastTradeTimestamp,
		"symbol":             symbol,
		"type":               type_,
		"timeInForce":        timeInForce,
		"side":               side,
		"price":              price,
		"stopPrice":          stopPrice,
		"amount":             amount,
		"cost":               MkUndefined(),
		"average":            average,
		"filled":             MkUndefined(),
		"remaining":          remaining,
		"status":             status,
		"fee":                fee,
		"trades":             trades,
	}))
}

func (this *Crex24) CreateOrder(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{
		"instrument": *(market).At(MkString("id")),
		"volume":     this.AmountToPrecision(symbol, amount),
		"type":       type_,
		"side":       side,
	})
	_ = request
	priceIsRequired := OpCopy(MkBool(false))
	_ = priceIsRequired
	stopPriceIsRequired := OpCopy(MkBool(false))
	_ = stopPriceIsRequired
	if IsTrue(OpEq2(type_, MkString("limit"))) {
		priceIsRequired = OpCopy(MkBool(true))
	} else {
		if IsTrue(OpEq2(type_, MkString("stopLimit"))) {
			priceIsRequired = OpCopy(MkBool(true))
			stopPriceIsRequired = OpCopy(MkBool(true))
		}
	}
	if IsTrue(priceIsRequired) {
		if IsTrue(OpEq2(price, MkUndefined())) {
			panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" createOrder() requires a price argument for a "), OpAdd(type_, MkString(" order"))))))
		}
		*(request).At(MkString("price")) = this.PriceToPrecision(symbol, price)
	}
	if IsTrue(stopPriceIsRequired) {
		stopPrice := this.SafeNumber(params, MkString("stopPrice"))
		_ = stopPrice
		if IsTrue(OpEq2(stopPrice, MkUndefined())) {
			panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" createOrder() requires a stopPrice extra param for a "), OpAdd(type_, MkString(" order"))))))
		} else {
			*(request).At(MkString("stopPrice")) = this.PriceToPrecision(symbol, stopPrice)
		}
		params = this.Omit(params, MkString("stopPrice"))
	}
	response := this.Call(MkString("tradingPostPlaceOrder"), this.Extend(request, params))
	_ = response
	return this.ParseOrder(response, market)
}

func (this *Crex24) FetchOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"id": id})
	_ = request
	response := this.Call(MkString("tradingGetOrderStatus"), this.Extend(request, params))
	_ = response
	numOrders := OpCopy(response.Length)
	_ = numOrders
	if IsTrue(OpLw(numOrders, MkInteger(1))) {
		panic(NewOrderNotFound(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" fetchOrder could not fetch order id "), id))))
	}
	return this.ParseOrder(*(response).At(MkInteger(0)))
}

func (this *Crex24) FetchOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{})
	_ = request
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("from")) = this.Ymdhms(since, MkString("T"))
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market := this.Market(symbol)
		_ = market
		*(request).At(MkString("instrument")) = *(market).At(MkString("id"))
	}
	method := this.SafeString(*this.At(MkString("options")), MkString("fetchOrdersMethod"), MkString("tradingGetOrderHistory"))
	_ = method
	response := this.Call(method, this.Extend(request, params))
	_ = response
	return this.ParseOrders(response)
}

func (this *Crex24) FetchOrdersByIds(goArgs ...*Variant) *Variant {
	ids := GoGetArg(goArgs, 0, MkUndefined())
	_ = ids
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"id": ids.Join(MkString(","))})
	_ = request
	response := this.Call(MkString("tradingGetOrderStatus"), this.Extend(request, params))
	_ = response
	return this.ParseOrders(response, MkUndefined(), since, limit)
}

func (this *Crex24) FetchOpenOrders(goArgs ...*Variant) *Variant {
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
		*(request).At(MkString("instrument")) = *(market).At(MkString("id"))
	}
	response := this.Call(MkString("tradingGetActiveOrders"), this.Extend(request, params))
	_ = response
	return this.ParseOrders(response, market, since, limit)
}

func (this *Crex24) FetchClosedOrders(goArgs ...*Variant) *Variant {
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
		*(request).At(MkString("instrument")) = *(market).At(MkString("id"))
	}
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("from")) = this.Ymdhms(since, MkString("T"))
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	method := this.SafeString(*this.At(MkString("options")), MkString("fetchClosedOrdersMethod"), MkString("tradingGetOrderHistory"))
	_ = method
	response := this.Call(method, this.Extend(request, params))
	_ = response
	return this.ParseOrders(response, market, since, limit)
}

func (this *Crex24) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"ids": MkArray(&VarArray{
		parseInt(id),
	})})
	_ = request
	response := this.Call(MkString("tradingPostCancelOrdersById"), this.Extend(request, params))
	_ = response
	return this.ParseOrder(response)
}

func (this *Crex24) CancelAllOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("tradingPostCancelAllOrders"), params)
	_ = response
	return response
}

func (this *Crex24) FetchMyTrades(goArgs ...*Variant) *Variant {
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
		*(request).At(MkString("instrument")) = *(market).At(MkString("id"))
	}
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("from")) = this.Ymdhms(since, MkString("T"))
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("tradingGetTradeHistory"), this.Extend(request, params))
	_ = response
	return this.ParseTrades(response, market, since, limit)
}

func (this *Crex24) FetchTransactions(goArgs ...*Variant) *Variant {
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
		*(request).At(MkString("currency")) = *(currency).At(MkString("id"))
	}
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("from")) = this.Ymd(since, MkString("T"))
	}
	response := this.Call(MkString("accountGetMoneyTransfers"), this.Extend(request, params))
	_ = response
	return this.ParseTransactions(response, currency, since, limit)
}

func (this *Crex24) FetchDeposits(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	request := MkMap(&VarMap{"type": MkString("deposit")})
	_ = request
	return this.FetchTransactions(code, since, limit, this.Extend(request, params))
}

func (this *Crex24) FetchWithdrawals(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	request := MkMap(&VarMap{"type": MkString("withdrawal")})
	_ = request
	return this.FetchTransactions(code, since, limit, this.Extend(request, params))
}

func (this *Crex24) ParseTransactionStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"pending": MkString("pending"),
		"success": MkString("ok"),
		"failed":  MkString("failed"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Crex24) ParseTransaction(goArgs ...*Variant) *Variant {
	transaction := GoGetArg(goArgs, 0, MkUndefined())
	_ = transaction
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	id := this.SafeString(transaction, MkString("id"))
	_ = id
	address := this.SafeString(transaction, MkString("address"))
	_ = address
	tag := this.SafeString(transaction, MkString("paymentId"))
	_ = tag
	txid := this.SafeValue(transaction, MkString("txId"))
	_ = txid
	currencyId := this.SafeString(transaction, MkString("currency"))
	_ = currencyId
	code := this.SafeCurrencyCode(currencyId, currency)
	_ = code
	type_ := this.SafeString(transaction, MkString("type"))
	_ = type_
	timestamp := this.Parse8601(this.SafeString(transaction, MkString("createdAt")))
	_ = timestamp
	updated := this.Parse8601(this.SafeString(transaction, MkString("processedAt")))
	_ = updated
	status := this.ParseTransactionStatus(this.SafeString(transaction, MkString("status")))
	_ = status
	amount := this.SafeNumber(transaction, MkString("amount"))
	_ = amount
	feeCost := this.SafeNumber(transaction, MkString("fee"))
	_ = feeCost
	fee := MkMap(&VarMap{
		"cost":     feeCost,
		"currency": code,
	})
	_ = fee
	return MkMap(&VarMap{
		"info":      transaction,
		"id":        id,
		"txid":      txid,
		"timestamp": timestamp,
		"datetime":  this.Iso8601(timestamp),
		"address":   address,
		"tag":       tag,
		"type":      type_,
		"amount":    amount,
		"currency":  code,
		"status":    status,
		"updated":   updated,
		"fee":       fee,
	})
}

func (this *Crex24) FetchDepositAddress(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{"currency": *(currency).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("accountGetDepositAddress"), this.Extend(request, params))
	_ = response
	address := this.SafeString(response, MkString("address"))
	_ = address
	tag := this.SafeString(response, MkString("paymentId"))
	_ = tag
	return MkMap(&VarMap{
		"currency": code,
		"address":  this.CheckAddress(address),
		"tag":      tag,
		"info":     response,
	})
}

func (this *Crex24) Withdraw(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{
		"currency":    *(currency).At(MkString("id")),
		"address":     address,
		"amount":      parseFloat(this.CurrencyToPrecision(code, amount)),
		"feeCurrency": *(currency).At(MkString("id")),
	})
	_ = request
	if IsTrue(OpNotEq2(tag, MkUndefined())) {
		*(request).At(MkString("paymentId")) = OpCopy(tag)
	}
	response := this.Call(MkString("accountPostWithdraw"), this.Extend(request, params))
	_ = response
	return this.ParseTransaction(response)
}

func (this *Crex24) Sign(goArgs ...*Variant) *Variant {
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
	request := OpAdd(MkString("/"), OpAdd(*this.At(MkString("version")), OpAdd(MkString("/"), OpAdd(api, OpAdd(MkString("/"), this.ImplodeParams(path, params))))))
	_ = request
	query := this.Omit(params, this.ExtractParams(path))
	_ = query
	if IsTrue(OpEq2(method, MkString("GET"))) {
		if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
			OpAddAssign(&request, OpAdd(MkString("?"), this.Urlencode(query)))
		}
	}
	url := OpAdd(*(*this.At(MkString("urls"))).At(MkString("api")), request)
	_ = url
	if IsTrue(OpOr(OpEq2(api, MkString("trading")), OpEq2(api, MkString("account")))) {
		this.CheckRequiredCredentials()
		nonce := (this.Nonce()).Call(MkString("toString"))
		_ = nonce
		secret := this.Base64ToBinary(*this.At(MkString("secret")))
		_ = secret
		auth := OpAdd(request, nonce)
		_ = auth
		headers = MkMap(&VarMap{
			"X-CREX24-API-KEY":   *this.At(MkString("apiKey")),
			"X-CREX24-API-NONCE": nonce,
		})
		if IsTrue(OpEq2(method, MkString("POST"))) {
			*(headers).At(MkString("Content-Type")) = MkString("application/json")
			body = this.Json(params)
			OpAddAssign(&auth, body)
		}
		*(headers).At(MkString("X-CREX24-API-SIGN")) = this.Hmac(this.Encode(auth), secret, MkString("sha512"), MkString("base64"))
	}
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Crex24) HandleErrors(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpNot(this.IsJsonEncodedObject(body))) {
		return MkUndefined()
	}
	if IsTrue(OpAnd(OpGtEq(code, MkInteger(200)), OpLw(code, MkInteger(300)))) {
		return MkUndefined()
	}
	message := this.SafeString(response, MkString("errorDescription"))
	_ = message
	feedback := OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))
	_ = feedback
	this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("exact")), message, feedback)
	this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("broad")), message, feedback)
	if IsTrue(OpEq2(code, MkInteger(400))) {
		panic(NewBadRequest(feedback))
	} else {
		if IsTrue(OpEq2(code, MkInteger(401))) {
			panic(NewAuthenticationError(feedback))
		} else {
			if IsTrue(OpEq2(code, MkInteger(403))) {
				panic(NewAuthenticationError(feedback))
			} else {
				if IsTrue(OpEq2(code, MkInteger(429))) {
					panic(NewDDoSProtection(feedback))
				} else {
					if IsTrue(OpEq2(code, MkInteger(500))) {
						panic(NewExchangeError(feedback))
					} else {
						if IsTrue(OpEq2(code, MkInteger(503))) {
							panic(NewExchangeNotAvailable(feedback))
						} else {
							if IsTrue(OpEq2(code, MkInteger(504))) {
								panic(NewRequestTimeout(feedback))
							}
						}
					}
				}
			}
		}
	}
	panic(NewExchangeError(feedback))
	return MkUndefined()
}
