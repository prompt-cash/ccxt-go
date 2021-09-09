package ccxt

import ()

type Poloniex struct {
	*ExchangeBase
}

var _ Exchange = (*Poloniex)(nil)

func init() {
	exchange := &Poloniex{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Poloniex) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("poloniex"),
		"name": MkString("Poloniex"),
		"countries": MkArray(&VarArray{
			MkString("US"),
		}),
		"rateLimit": MkInteger(1000),
		"certified": MkBool(false),
		"pro":       MkBool(true),
		"has": MkMap(&VarMap{
			"cancelOrder":          MkBool(true),
			"CORS":                 MkBool(false),
			"createDepositAddress": MkBool(true),
			"createMarketOrder":    MkBool(false),
			"createOrder":          MkBool(true),
			"editOrder":            MkBool(true),
			"fetchBalance":         MkBool(true),
			"fetchClosedOrder":     MkString("emulated"),
			"fetchCurrencies":      MkBool(true),
			"fetchDepositAddress":  MkBool(true),
			"fetchDeposits":        MkBool(true),
			"fetchMarkets":         MkBool(true),
			"fetchMyTrades":        MkBool(true),
			"fetchOHLCV":           MkBool(true),
			"fetchOpenOrder":       MkBool(true),
			"fetchOpenOrders":      MkBool(true),
			"fetchOrderBook":       MkBool(true),
			"fetchOrderBooks":      MkBool(true),
			"fetchOrderTrades":     MkBool(true),
			"fetchTicker":          MkBool(true),
			"fetchTickers":         MkBool(true),
			"fetchTrades":          MkBool(true),
			"fetchTradingFee":      MkBool(true),
			"fetchTradingFees":     MkBool(true),
			"fetchTransactions":    MkBool(true),
			"fetchWithdrawals":     MkBool(true),
			"cancelAllOrders":      MkBool(true),
			"withdraw":             MkBool(true),
		}),
		"timeframes": MkMap(&VarMap{
			"5m":  MkInteger(300),
			"15m": MkInteger(900),
			"30m": MkInteger(1800),
			"2h":  MkInteger(7200),
			"4h":  MkInteger(14400),
			"1d":  MkInteger(86400),
		}),
		"urls": MkMap(&VarMap{
			"logo": MkString("https://user-images.githubusercontent.com/1294454/27766817-e9456312-5ee6-11e7-9b3c-b628ca5626a5.jpg"),
			"api": MkMap(&VarMap{
				"public":  MkString("https://poloniex.com/public"),
				"private": MkString("https://poloniex.com/tradingApi"),
			}),
			"www":      MkString("https://www.poloniex.com"),
			"doc":      MkString("https://docs.poloniex.com"),
			"fees":     MkString("https://poloniex.com/fees"),
			"referral": MkString("https://poloniex.com/signup?c=UBFZJRPJ"),
		}),
		"api": MkMap(&VarMap{
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("return24hVolume"),
				MkString("returnChartData"),
				MkString("returnCurrencies"),
				MkString("returnLoanOrders"),
				MkString("returnOrderBook"),
				MkString("returnTicker"),
				MkString("returnTradeHistory"),
			})}),
			"private": MkMap(&VarMap{"post": MkArray(&VarArray{
				MkString("buy"),
				MkString("cancelLoanOffer"),
				MkString("cancelOrder"),
				MkString("cancelAllOrders"),
				MkString("closeMarginPosition"),
				MkString("createLoanOffer"),
				MkString("generateNewAddress"),
				MkString("getMarginPosition"),
				MkString("marginBuy"),
				MkString("marginSell"),
				MkString("moveOrder"),
				MkString("returnActiveLoans"),
				MkString("returnAvailableAccountBalances"),
				MkString("returnBalances"),
				MkString("returnCompleteBalances"),
				MkString("returnDepositAddresses"),
				MkString("returnDepositsWithdrawals"),
				MkString("returnFeeInfo"),
				MkString("returnLendingHistory"),
				MkString("returnMarginAccountSummary"),
				MkString("returnOpenLoanOffers"),
				MkString("returnOpenOrders"),
				MkString("returnOrderTrades"),
				MkString("returnOrderStatus"),
				MkString("returnTradableBalances"),
				MkString("returnTradeHistory"),
				MkString("sell"),
				MkString("toggleAutoRenew"),
				MkString("transferBalance"),
				MkString("withdraw"),
			})}),
		}),
		"fees": MkMap(&VarMap{
			"trading": MkMap(&VarMap{
				"feeSide": MkString("get"),
				"maker":   this.ParseNumber(MkString("0.0009")),
				"taker":   this.ParseNumber(MkString("0.0009")),
			}),
			"funding": MkMap(&VarMap{}),
		}),
		"limits": MkMap(&VarMap{
			"amount": MkMap(&VarMap{
				"min": MkNumber(0.000001),
				"max": MkUndefined(),
			}),
			"price": MkMap(&VarMap{
				"min": MkNumber(0.00000001),
				"max": MkInteger(1000000000),
			}),
			"cost": MkMap(&VarMap{
				"min": MkNumber(0.00000000),
				"max": MkInteger(1000000000),
			}),
		}),
		"precision": MkMap(&VarMap{
			"amount": MkInteger(8),
			"price":  MkInteger(8),
		}),
		"commonCurrencies": MkMap(&VarMap{
			"AIR":      MkString("AirCoin"),
			"APH":      MkString("AphroditeCoin"),
			"BCC":      MkString("BTCtalkcoin"),
			"BCHABC":   MkString("BCHABC"),
			"BDG":      MkString("Badgercoin"),
			"BTM":      MkString("Bitmark"),
			"CON":      MkString("Coino"),
			"GOLD":     MkString("GoldEagles"),
			"GPUC":     MkString("GPU"),
			"HOT":      MkString("Hotcoin"),
			"ITC":      MkString("Information Coin"),
			"KEY":      MkString("KEYCoin"),
			"PLX":      MkString("ParallaxCoin"),
			"REPV2":    MkString("REP"),
			"STR":      MkString("XLM"),
			"SOC":      MkString("SOCC"),
			"XAP":      MkString("API Coin"),
			"USDTTRON": MkString("USDT"),
			"USDTETH":  MkString("USDT"),
		}),
		"options": MkMap(&VarMap{"limits": MkMap(&VarMap{"cost": MkMap(&VarMap{"min": MkMap(&VarMap{
			"BTC":  MkNumber(0.0001),
			"ETH":  MkNumber(0.0001),
			"USDT": MkNumber(1.0),
			"TRX":  MkInteger(100),
			"BNB":  MkNumber(0.06),
			"USDC": MkNumber(1.0),
			"USDJ": MkNumber(1.0),
			"TUSD": MkNumber(0.0001),
			"DAI":  MkNumber(1.0),
			"PAX":  MkNumber(1.0),
			"BUSD": MkNumber(1.0),
		})})})}),
		"exceptions": MkMap(&VarMap{
			"exact": MkMap(&VarMap{
				"You may only place orders that reduce your position.":                  InvalidOrder,
				"Invalid order number, or you are not the person who placed the order.": OrderNotFound,
				"Permission denied":                                         PermissionDenied,
				"Connection timed out. Please try again.":                   RequestTimeout,
				"Internal error. Please try again.":                         ExchangeNotAvailable,
				"Currently in maintenance mode.":                            OnMaintenance,
				"Order not found, or you are not the person who placed it.": OrderNotFound,
				"Invalid API key/secret pair.":                              AuthenticationError,
				"Please do not make more than 8 API calls per second.":      DDoSProtection,
				"Rate must be greater than zero.":                           InvalidOrder,
				"Invalid currency pair.":                                    BadSymbol,
				"Invalid currencyPair parameter.":                           BadSymbol,
				"Trading is disabled in this market.":                       BadSymbol,
				"Invalid orderNumber parameter.":                            OrderNotFound,
			}),
			"broad": MkMap(&VarMap{
				"Total must be at least":  InvalidOrder,
				"This account is frozen.": AccountSuspended,
				"Not enough":              InsufficientFunds,
				"Nonce must be greater":   InvalidNonce,
				"You have already called cancelOrder or moveOrder on this order.": CancelPending,
				"Amount must be at least":               InvalidOrder,
				"is either completed or does not exist": OrderNotFound,
				"Error pulling ":                        ExchangeError,
			}),
		}),
	}))
}

func (this *Poloniex) ParseOHLCV(goArgs ...*Variant) *Variant {
	ohlcv := GoGetArg(goArgs, 0, MkUndefined())
	_ = ohlcv
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	return MkArray(&VarArray{
		this.SafeTimestamp(ohlcv, MkString("date")),
		this.SafeNumber(ohlcv, MkString("open")),
		this.SafeNumber(ohlcv, MkString("high")),
		this.SafeNumber(ohlcv, MkString("low")),
		this.SafeNumber(ohlcv, MkString("close")),
		this.SafeNumber(ohlcv, MkString("quoteVolume")),
	})
}

func (this *Poloniex) FetchOHLCV(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	timeframe := GoGetArg(goArgs, 1, MkString("5m"))
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
		"currencyPair": *(market).At(MkString("id")),
		"period":       *(*this.At(MkString("timeframes"))).At(timeframe),
	})
	_ = request
	if IsTrue(OpEq2(since, MkUndefined())) {
		*(request).At(MkString("end")) = this.Seconds()
		if IsTrue(OpEq2(limit, MkUndefined())) {
			*(request).At(MkString("start")) = OpSub(*(request).At(MkString("end")), this.ParseTimeframe(MkString("1w")))
		} else {
			*(request).At(MkString("start")) = OpSub(*(request).At(MkString("end")), OpMulti(limit, this.ParseTimeframe(timeframe)))
		}
	} else {
		*(request).At(MkString("start")) = parseInt(OpDiv(since, MkInteger(1000)))
		if IsTrue(OpNotEq2(limit, MkUndefined())) {
			end := this.Sum(*(request).At(MkString("start")), OpMulti(limit, this.ParseTimeframe(timeframe)))
			_ = end
			*(request).At(MkString("end")) = OpCopy(end)
		}
	}
	response := this.Call(MkString("publicGetReturnChartData"), this.Extend(request, params))
	_ = response
	return this.ParseOHLCVs(response, market, timeframe, since, limit)
}

func (this *Poloniex) LoadMarkets(goArgs ...*Variant) *Variant {
	reload := GoGetArg(goArgs, 0, MkBool(false))
	_ = reload
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	markets := this.LoadMarkets(reload, params)
	_ = markets
	currenciesByNumericId := this.SafeValue(*this.At(MkString("options")), MkString("currenciesByNumericId"))
	_ = currenciesByNumericId
	if IsTrue(OpOr(OpEq2(currenciesByNumericId, MkUndefined()), reload)) {
		*(*this.At(MkString("options"))).At(MkString("currenciesByNumericId")) = this.IndexBy(*this.At(MkString("currencies")), MkString("numericId"))
	}
	return markets
}

func (this *Poloniex) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	markets := this.Call(MkString("publicGetReturnTicker"), params)
	_ = markets
	keys := GoGetKeys(markets)
	_ = keys
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, keys.Length)); OpIncr(&i) {
		id := *(keys).At(i)
		_ = id
		market := *(markets).At(id)
		_ = market
		quoteId := MkUndefined()
		_ = quoteId
		baseId := MkUndefined()
		_ = baseId
		ArrayUnpack(id.Split(MkString("_")), &quoteId, &baseId)
		base := this.SafeCurrencyCode(baseId)
		_ = base
		quote := this.SafeCurrencyCode(quoteId)
		_ = quote
		symbol := OpAdd(base, OpAdd(MkString("/"), quote))
		_ = symbol
		limits := this.Extend(*this.At(MkString("limits")), MkMap(&VarMap{"cost": MkMap(&VarMap{"min": this.SafeValue(*(*(*(*this.At(MkString("options"))).At(MkString("limits"))).At(MkString("cost"))).At(MkString("min")), quote)})}))
		_ = limits
		isFrozen := this.SafeString(market, MkString("isFrozen"))
		_ = isFrozen
		active := OpNotEq2(isFrozen, MkString("1"))
		_ = active
		numericId := this.SafeInteger(market, MkString("id"))
		_ = numericId
		result.Push(MkMap(&VarMap{
			"id":        id,
			"numericId": numericId,
			"symbol":    symbol,
			"baseId":    baseId,
			"quoteId":   quoteId,
			"base":      base,
			"quote":     quote,
			"active":    active,
			"limits":    limits,
			"info":      market,
		}))
	}
	return result
}

func (this *Poloniex) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"account": MkString("all")})
	_ = request
	response := this.Call(MkString("privatePostReturnCompleteBalances"), this.Extend(request, params))
	_ = response
	result := MkMap(&VarMap{
		"info":      response,
		"timestamp": MkUndefined(),
		"datetime":  MkUndefined(),
	})
	_ = result
	currencyIds := GoGetKeys(response)
	_ = currencyIds
	for i := MkInteger(0); IsTrue(OpLw(i, currencyIds.Length)); OpIncr(&i) {
		currencyId := *(currencyIds).At(i)
		_ = currencyId
		balance := this.SafeValue(response, currencyId, MkMap(&VarMap{}))
		_ = balance
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		account := this.Account()
		_ = account
		*(account).At(MkString("free")) = this.SafeString(balance, MkString("available"))
		*(account).At(MkString("used")) = this.SafeString(balance, MkString("onOrders"))
		*(result).At(code) = OpCopy(account)
	}
	return this.ParseBalance(result)
}

func (this *Poloniex) FetchTradingFees(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	fees := this.Call(MkString("privatePostReturnFeeInfo"), params)
	_ = fees
	return MkMap(&VarMap{
		"info":     fees,
		"maker":    this.SafeNumber(fees, MkString("makerFee")),
		"taker":    this.SafeNumber(fees, MkString("takerFee")),
		"withdraw": MkMap(&VarMap{}),
		"deposit":  MkMap(&VarMap{}),
	})
}

func (this *Poloniex) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"currencyPair": this.MarketId(symbol)})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("depth")) = OpCopy(limit)
	}
	response := this.Call(MkString("publicGetReturnOrderBook"), this.Extend(request, params))
	_ = response
	orderbook := this.ParseOrderBook(response, symbol)
	_ = orderbook
	*(orderbook).At(MkString("nonce")) = this.SafeInteger(response, MkString("seq"))
	return orderbook
}

func (this *Poloniex) FetchOrderBooks(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"currencyPair": MkString("all")})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("depth")) = OpCopy(limit)
	}
	response := this.Call(MkString("publicGetReturnOrderBook"), this.Extend(request, params))
	_ = response
	marketIds := GoGetKeys(response)
	_ = marketIds
	result := MkMap(&VarMap{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, marketIds.Length)); OpIncr(&i) {
		marketId := *(marketIds).At(i)
		_ = marketId
		symbol := OpCopy(MkUndefined())
		_ = symbol
		if IsTrue(OpHasMember(marketId, *this.At(MkString("markets_by_id")))) {
			symbol = *(*(*this.At(MkString("markets_by_id"))).At(marketId)).At(MkString("symbol"))
		} else {
			quoteId := MkUndefined()
			_ = quoteId
			baseId := MkUndefined()
			_ = baseId
			ArrayUnpack(marketId.Split(MkString("_")), &quoteId, &baseId)
			base := this.SafeCurrencyCode(baseId)
			_ = base
			quote := this.SafeCurrencyCode(quoteId)
			_ = quote
			symbol = OpAdd(base, OpAdd(MkString("/"), quote))
		}
		orderbook := this.ParseOrderBook(*(response).At(marketId), symbol)
		_ = orderbook
		*(orderbook).At(MkString("nonce")) = this.SafeInteger(*(response).At(marketId), MkString("seq"))
		*(result).At(symbol) = OpCopy(orderbook)
	}
	return result
}

func (this *Poloniex) ParseTicker(goArgs ...*Variant) *Variant {
	ticker := GoGetArg(goArgs, 0, MkUndefined())
	_ = ticker
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.Milliseconds()
	_ = timestamp
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(market) {
		symbol = *(market).At(MkString("symbol"))
	}
	open := OpCopy(MkUndefined())
	_ = open
	change := OpCopy(MkUndefined())
	_ = change
	average := OpCopy(MkUndefined())
	_ = average
	last := this.SafeNumber(ticker, MkString("last"))
	_ = last
	relativeChange := this.SafeNumber(ticker, MkString("percentChange"))
	_ = relativeChange
	if IsTrue(OpNotEq2(relativeChange, OpNeg(MkInteger(1)))) {
		open = OpDiv(last, this.Sum(MkInteger(1), relativeChange))
		change = OpSub(last, open)
		average = OpDiv(this.Sum(last, open), MkInteger(2))
	}
	return MkMap(&VarMap{
		"symbol":        symbol,
		"timestamp":     timestamp,
		"datetime":      this.Iso8601(timestamp),
		"high":          this.SafeNumber(ticker, MkString("high24hr")),
		"low":           this.SafeNumber(ticker, MkString("low24hr")),
		"bid":           this.SafeNumber(ticker, MkString("highestBid")),
		"bidVolume":     MkUndefined(),
		"ask":           this.SafeNumber(ticker, MkString("lowestAsk")),
		"askVolume":     MkUndefined(),
		"vwap":          MkUndefined(),
		"open":          open,
		"close":         last,
		"last":          last,
		"previousClose": MkUndefined(),
		"change":        change,
		"percentage":    OpMulti(relativeChange, MkInteger(100)),
		"average":       average,
		"baseVolume":    this.SafeNumber(ticker, MkString("quoteVolume")),
		"quoteVolume":   this.SafeNumber(ticker, MkString("baseVolume")),
		"info":          ticker,
	})
}

func (this *Poloniex) FetchTickers(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("publicGetReturnTicker"), params)
	_ = response
	ids := GoGetKeys(response)
	_ = ids
	result := MkMap(&VarMap{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, ids.Length)); OpIncr(&i) {
		id := *(ids).At(i)
		_ = id
		symbol := OpCopy(MkUndefined())
		_ = symbol
		market := OpCopy(MkUndefined())
		_ = market
		if IsTrue(OpHasMember(id, *this.At(MkString("markets_by_id")))) {
			market = *(*this.At(MkString("markets_by_id"))).At(id)
			symbol = *(market).At(MkString("symbol"))
		} else {
			quoteId := MkUndefined()
			_ = quoteId
			baseId := MkUndefined()
			_ = baseId
			ArrayUnpack(id.Split(MkString("_")), &quoteId, &baseId)
			base := this.SafeCurrencyCode(baseId)
			_ = base
			quote := this.SafeCurrencyCode(quoteId)
			_ = quote
			symbol = OpAdd(base, OpAdd(MkString("/"), quote))
			market = MkMap(&VarMap{"symbol": symbol})
		}
		ticker := *(response).At(id)
		_ = ticker
		*(result).At(symbol) = this.ParseTicker(ticker, market)
	}
	return this.FilterByArray(result, MkString("symbol"), symbols)
}

func (this *Poloniex) FetchCurrencies(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetReturnCurrencies"), params)
	_ = response
	ids := GoGetKeys(response)
	_ = ids
	result := MkMap(&VarMap{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, ids.Length)); OpIncr(&i) {
		id := *(ids).At(i)
		_ = id
		currency := *(response).At(id)
		_ = currency
		precision := MkInteger(8)
		_ = precision
		amountLimit := MkString("1e-8")
		_ = amountLimit
		code := this.SafeCurrencyCode(id)
		_ = code
		active := OpAnd(OpEq2(*(currency).At(MkString("delisted")), MkInteger(0)), OpNot(*(currency).At(MkString("disabled"))))
		_ = active
		numericId := this.SafeInteger(currency, MkString("id"))
		_ = numericId
		fee := this.SafeNumber(currency, MkString("txFee"))
		_ = fee
		*(result).At(code) = MkMap(&VarMap{
			"id":        id,
			"numericId": numericId,
			"code":      code,
			"info":      currency,
			"name":      *(currency).At(MkString("name")),
			"active":    active,
			"fee":       fee,
			"precision": precision,
			"limits": MkMap(&VarMap{
				"amount": MkMap(&VarMap{
					"min": this.ParseNumber(amountLimit),
					"max": MkUndefined(),
				}),
				"withdraw": MkMap(&VarMap{
					"min": fee,
					"max": MkUndefined(),
				}),
			}),
		})
	}
	return result
}

func (this *Poloniex) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	response := this.Call(MkString("publicGetReturnTicker"), params)
	_ = response
	ticker := *(response).At(*(market).At(MkString("id")))
	_ = ticker
	return this.ParseTicker(ticker, market)
}

func (this *Poloniex) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	id := this.SafeString2(trade, MkString("globalTradeID"), MkString("tradeID"))
	_ = id
	orderId := this.SafeString(trade, MkString("orderNumber"))
	_ = orderId
	timestamp := this.Parse8601(this.SafeString(trade, MkString("date")))
	_ = timestamp
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(OpAnd(OpNot(market), OpHasMember(MkString("currencyPair"), trade))) {
		marketId := this.SafeString(trade, MkString("currencyPair"))
		_ = marketId
		if IsTrue(OpHasMember(marketId, *this.At(MkString("markets_by_id")))) {
			market = *(*this.At(MkString("markets_by_id"))).At(marketId)
		} else {
			quoteId := MkUndefined()
			_ = quoteId
			baseId := MkUndefined()
			_ = baseId
			ArrayUnpack(marketId.Split(MkString("_")), &quoteId, &baseId)
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
	side := this.SafeString(trade, MkString("type"))
	_ = side
	fee := OpCopy(MkUndefined())
	_ = fee
	priceString := this.SafeString(trade, MkString("rate"))
	_ = priceString
	amountString := this.SafeString(trade, MkString("amount"))
	_ = amountString
	price := this.ParseNumber(priceString)
	_ = price
	amount := this.ParseNumber(amountString)
	_ = amount
	cost := OpCopy(MkUndefined())
	_ = cost
	costString := this.SafeString(trade, MkString("total"))
	_ = costString
	if IsTrue(OpEq2(costString, MkUndefined())) {
		costString = Precise.StringMul(priceString, amountString)
		cost = this.ParseNumber(costString)
	} else {
		cost = this.ParseNumber(costString)
	}
	feeDisplay := this.SafeString(trade, MkString("feeDisplay"))
	_ = feeDisplay
	if IsTrue(OpNotEq2(feeDisplay, MkUndefined())) {
		parts := feeDisplay.Split(MkString(" "))
		_ = parts
		feeCost := this.SafeNumber(parts, MkInteger(0))
		_ = feeCost
		if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
			feeCurrencyId := this.SafeString(parts, MkInteger(1))
			_ = feeCurrencyId
			feeCurrencyCode := this.SafeCurrencyCode(feeCurrencyId)
			_ = feeCurrencyCode
			feeRate := this.SafeString(parts, MkInteger(2))
			_ = feeRate
			if IsTrue(OpNotEq2(feeRate, MkUndefined())) {
				feeRate = feeRate.Replace(MkString("("), MkString(""))
				feeRateParts := feeRate.Split(MkString("%"))
				_ = feeRateParts
				feeRate = this.SafeString(feeRateParts, MkInteger(0))
				feeRate = OpDiv(parseFloat(feeRate), MkInteger(100))
			}
			fee = MkMap(&VarMap{
				"cost":     feeCost,
				"currency": feeCurrencyCode,
				"rate":     feeRate,
			})
		}
	} else {
		feeCostString := this.SafeString(trade, MkString("fee"))
		_ = feeCostString
		if IsTrue(OpAnd(OpNotEq2(feeCostString, MkUndefined()), OpNotEq2(market, MkUndefined()))) {
			feeCurrencyCode := OpTrinary(OpEq2(side, MkString("buy")), *(market).At(MkString("base")), *(market).At(MkString("quote")))
			_ = feeCurrencyCode
			feeBase := OpTrinary(OpEq2(side, MkString("buy")), amountString, costString)
			_ = feeBase
			feeRateString := Precise.StringDiv(feeCostString, feeBase)
			_ = feeRateString
			fee = MkMap(&VarMap{
				"cost":     this.ParseNumber(feeCostString),
				"currency": feeCurrencyCode,
				"rate":     this.ParseNumber(feeRateString),
			})
		}
	}
	takerOrMaker := OpCopy(MkUndefined())
	_ = takerOrMaker
	takerAdjustment := this.SafeNumber(trade, MkString("takerAdjustment"))
	_ = takerAdjustment
	if IsTrue(OpNotEq2(takerAdjustment, MkUndefined())) {
		takerOrMaker = MkString("taker")
	}
	return MkMap(&VarMap{
		"id":           id,
		"info":         trade,
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       symbol,
		"order":        orderId,
		"type":         MkString("limit"),
		"side":         side,
		"takerOrMaker": takerOrMaker,
		"price":        price,
		"amount":       amount,
		"cost":         cost,
		"fee":          fee,
	})
}

func (this *Poloniex) FetchTrades(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"currencyPair": *(market).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("start")) = parseInt(OpDiv(since, MkInteger(1000)))
		*(request).At(MkString("end")) = this.Seconds()
	}
	trades := this.Call(MkString("publicGetReturnTradeHistory"), this.Extend(request, params))
	_ = trades
	return this.ParseTrades(trades, market, since, limit)
}

func (this *Poloniex) FetchMyTrades(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
	}
	pair := OpTrinary(market, *(market).At(MkString("id")), MkString("all"))
	_ = pair
	request := MkMap(&VarMap{"currencyPair": pair})
	_ = request
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("start")) = parseInt(OpDiv(since, MkInteger(1000)))
		*(request).At(MkString("end")) = this.Sum(this.Seconds(), MkInteger(1))
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = parseInt(limit)
	}
	response := this.Call(MkString("privatePostReturnTradeHistory"), this.Extend(request, params))
	_ = response
	result := MkArray(&VarArray{})
	_ = result
	if IsTrue(OpNotEq2(market, MkUndefined())) {
		result = this.ParseTrades(response, market)
	} else {
		if IsTrue(response) {
			ids := GoGetKeys(response)
			_ = ids
			for i := MkInteger(0); IsTrue(OpLw(i, ids.Length)); OpIncr(&i) {
				id := *(ids).At(i)
				_ = id
				market := OpCopy(MkUndefined())
				_ = market
				if IsTrue(OpHasMember(id, *this.At(MkString("markets_by_id")))) {
					market = *(*this.At(MkString("markets_by_id"))).At(id)
					trades := this.ParseTrades(*(response).At(id), market)
					_ = trades
					for j := MkInteger(0); IsTrue(OpLw(j, trades.Length)); OpIncr(&j) {
						result.Push(*(trades).At(j))
					}
				} else {
					quoteId := MkUndefined()
					_ = quoteId
					baseId := MkUndefined()
					_ = baseId
					ArrayUnpack(id.Split(MkString("_")), &quoteId, &baseId)
					base := this.SafeCurrencyCode(baseId)
					_ = base
					quote := this.SafeCurrencyCode(quoteId)
					_ = quote
					symbol := OpAdd(base, OpAdd(MkString("/"), quote))
					_ = symbol
					trades := *(response).At(id)
					_ = trades
					for j := MkInteger(0); IsTrue(OpLw(j, trades.Length)); OpIncr(&j) {
						market := MkMap(&VarMap{
							"symbol": symbol,
							"base":   base,
							"quote":  quote,
						})
						_ = market
						result.Push(this.ParseTrade(*(trades).At(j), market))
					}
				}
			}
		}
	}
	return this.FilterBySinceLimit(result, since, limit)
}

func (this *Poloniex) ParseOrderStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"Open":             MkString("open"),
		"Partially filled": MkString("open"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Poloniex) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.SafeInteger(order, MkString("timestamp"))
	_ = timestamp
	if IsTrue(OpEq2(timestamp, MkUndefined())) {
		timestamp = this.Parse8601(this.SafeString(order, MkString("date")))
	}
	symbol := OpCopy(MkUndefined())
	_ = symbol
	marketId := this.SafeString(order, MkString("currencyPair"))
	_ = marketId
	if IsTrue(OpNotEq2(marketId, MkUndefined())) {
		if IsTrue(OpHasMember(marketId, *this.At(MkString("markets_by_id")))) {
			market = *(*this.At(MkString("markets_by_id"))).At(marketId)
		} else {
			quoteId := MkUndefined()
			_ = quoteId
			baseId := MkUndefined()
			_ = baseId
			ArrayUnpack(marketId.Split(MkString("_")), &quoteId, &baseId)
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
	trades := OpCopy(MkUndefined())
	_ = trades
	resultingTrades := this.SafeValue(order, MkString("resultingTrades"))
	_ = resultingTrades
	if IsTrue(OpNot(GoIsArray(resultingTrades))) {
		resultingTrades = this.SafeValue(resultingTrades, this.SafeString(market, MkString("id"), marketId))
	}
	if IsTrue(OpNotEq2(resultingTrades, MkUndefined())) {
		trades = this.ParseTrades(resultingTrades, market)
	}
	price := this.SafeNumber2(order, MkString("price"), MkString("rate"))
	_ = price
	remaining := this.SafeNumber(order, MkString("amount"))
	_ = remaining
	amount := this.SafeNumber(order, MkString("startingAmount"))
	_ = amount
	filled := OpCopy(MkUndefined())
	_ = filled
	cost := MkInteger(0)
	_ = cost
	if IsTrue(OpNotEq2(amount, MkUndefined())) {
		if IsTrue(OpNotEq2(remaining, MkUndefined())) {
			filled = OpSub(amount, remaining)
			if IsTrue(OpNotEq2(price, MkUndefined())) {
				cost = OpMulti(filled, price)
			}
		}
	} else {
		amount = OpCopy(remaining)
	}
	status := this.ParseOrderStatus(this.SafeString(order, MkString("status")))
	_ = status
	average := OpCopy(MkUndefined())
	_ = average
	lastTradeTimestamp := OpCopy(MkUndefined())
	_ = lastTradeTimestamp
	if IsTrue(OpEq2(filled, MkUndefined())) {
		if IsTrue(OpNotEq2(trades, MkUndefined())) {
			filled = MkInteger(0)
			cost = MkInteger(0)
			tradesLength := OpCopy(trades.Length)
			_ = tradesLength
			if IsTrue(OpGt(tradesLength, MkInteger(0))) {
				lastTradeTimestamp = *(*(trades).At(MkInteger(0))).At(MkString("timestamp"))
				for i := MkInteger(0); IsTrue(OpLw(i, tradesLength)); OpIncr(&i) {
					trade := *(trades).At(i)
					_ = trade
					tradeAmount := *(trade).At(MkString("amount"))
					_ = tradeAmount
					tradePrice := *(trade).At(MkString("price"))
					_ = tradePrice
					filled = this.Sum(filled, tradeAmount)
					cost = this.Sum(cost, OpMulti(tradePrice, tradeAmount))
					lastTradeTimestamp = Math.Max(lastTradeTimestamp, *(trade).At(MkString("timestamp")))
				}
			}
			if IsTrue(OpNotEq2(amount, MkUndefined())) {
				remaining = Math.Max(OpSub(amount, filled), MkInteger(0))
				if IsTrue(OpGtEq(filled, amount)) {
					status = MkString("closed")
				}
			}
		}
	}
	if IsTrue(OpAnd(OpNotEq2(filled, MkUndefined()), OpAnd(OpNotEq2(cost, MkUndefined()), OpGt(filled, MkInteger(0))))) {
		average = OpDiv(cost, filled)
	}
	type_ := this.SafeString(order, MkString("type"))
	_ = type_
	side := this.SafeString(order, MkString("side"), type_)
	_ = side
	if IsTrue(OpEq2(type_, side)) {
		type_ = OpCopy(MkUndefined())
	}
	id := this.SafeString(order, MkString("orderNumber"))
	_ = id
	fee := OpCopy(MkUndefined())
	_ = fee
	feeCost := this.SafeNumber(order, MkString("fee"))
	_ = feeCost
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		feeCurrencyCode := OpCopy(MkUndefined())
		_ = feeCurrencyCode
		if IsTrue(OpNotEq2(market, MkUndefined())) {
			feeCurrencyCode = OpTrinary(OpEq2(side, MkString("buy")), *(market).At(MkString("base")), *(market).At(MkString("quote")))
		}
		fee = MkMap(&VarMap{
			"cost":     feeCost,
			"currency": feeCurrencyCode,
		})
	}
	clientOrderId := this.SafeString(order, MkString("clientOrderId"))
	_ = clientOrderId
	return MkMap(&VarMap{
		"info":               order,
		"id":                 id,
		"clientOrderId":      clientOrderId,
		"timestamp":          timestamp,
		"datetime":           this.Iso8601(timestamp),
		"lastTradeTimestamp": lastTradeTimestamp,
		"status":             status,
		"symbol":             symbol,
		"type":               type_,
		"timeInForce":        MkUndefined(),
		"postOnly":           MkUndefined(),
		"side":               side,
		"price":              price,
		"stopPrice":          MkUndefined(),
		"cost":               cost,
		"average":            average,
		"amount":             amount,
		"filled":             filled,
		"remaining":          remaining,
		"trades":             trades,
		"fee":                fee,
	})
}

func (this *Poloniex) ParseOpenOrders(goArgs ...*Variant) *Variant {
	orders := GoGetArg(goArgs, 0, MkUndefined())
	_ = orders
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	result := GoGetArg(goArgs, 2, MkUndefined())
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, orders.Length)); OpIncr(&i) {
		order := *(orders).At(i)
		_ = order
		extended := this.Extend(order, MkMap(&VarMap{
			"status": MkString("open"),
			"type":   MkString("limit"),
			"side":   *(order).At(MkString("type")),
			"price":  *(order).At(MkString("rate")),
		}))
		_ = extended
		result.Push(this.ParseOrder(extended, market))
	}
	return result
}

func (this *Poloniex) FetchOpenOrders(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
	}
	pair := OpTrinary(market, *(market).At(MkString("id")), MkString("all"))
	_ = pair
	request := MkMap(&VarMap{"currencyPair": pair})
	_ = request
	response := this.Call(MkString("privatePostReturnOpenOrders"), this.Extend(request, params))
	_ = response
	extension := MkMap(&VarMap{"status": MkString("open")})
	_ = extension
	if IsTrue(OpEq2(market, MkUndefined())) {
		marketIds := GoGetKeys(response)
		_ = marketIds
		openOrders := MkArray(&VarArray{})
		_ = openOrders
		for i := MkInteger(0); IsTrue(OpLw(i, marketIds.Length)); OpIncr(&i) {
			marketId := *(marketIds).At(i)
			_ = marketId
			orders := *(response).At(marketId)
			_ = orders
			m := *(*this.At(MkString("markets_by_id"))).At(marketId)
			_ = m
			openOrders = this.ArrayConcat(openOrders, this.ParseOrders(orders, m, MkUndefined(), MkUndefined(), extension))
		}
		return this.FilterBySinceLimit(openOrders, since, limit)
	} else {
		return this.ParseOrders(response, market, since, limit, extension)
	}
	return MkUndefined()
}

func (this *Poloniex) CreateOrder(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpEq2(type_, MkString("market"))) {
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), MkString(" createOrder() does not accept market orders"))))
	}
	this.LoadMarkets()
	method := OpAdd(MkString("privatePost"), this.Capitalize(side))
	_ = method
	market := this.Market(symbol)
	_ = market
	amount = this.AmountToPrecision(symbol, amount)
	request := MkMap(&VarMap{
		"currencyPair": *(market).At(MkString("id")),
		"rate":         this.PriceToPrecision(symbol, price),
		"amount":       amount,
	})
	_ = request
	clientOrderId := this.SafeString(params, MkString("clientOrderId"))
	_ = clientOrderId
	if IsTrue(OpNotEq2(clientOrderId, MkUndefined())) {
		*(request).At(MkString("clientOrderId")) = OpCopy(clientOrderId)
		params = this.Omit(params, MkString("clientOrderId"))
	}
	timestamp := this.Milliseconds()
	_ = timestamp
	response := this.Call(method, this.Extend(request, params))
	_ = response
	return this.ParseOrder(this.Extend(MkMap(&VarMap{
		"timestamp": timestamp,
		"status":    MkString("open"),
		"type":      type_,
		"side":      side,
		"price":     price,
		"amount":    amount,
	}), response), market)
}

func (this *Poloniex) EditOrder(goArgs ...*Variant) *Variant {
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
	this.LoadMarkets()
	price = parseFloat(price)
	request := MkMap(&VarMap{
		"orderNumber": id,
		"rate":        this.PriceToPrecision(symbol, price),
	})
	_ = request
	if IsTrue(OpNotEq2(amount, MkUndefined())) {
		*(request).At(MkString("amount")) = this.AmountToPrecision(symbol, amount)
	}
	response := this.Call(MkString("privatePostMoveOrder"), this.Extend(request, params))
	_ = response
	return this.ParseOrder(response)
}

func (this *Poloniex) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{})
	_ = request
	clientOrderId := this.SafeValue(params, MkString("clientOrderId"))
	_ = clientOrderId
	if IsTrue(OpEq2(clientOrderId, MkUndefined())) {
		*(request).At(MkString("orderNumber")) = OpCopy(id)
	} else {
		*(request).At(MkString("clientOrderId")) = OpCopy(clientOrderId)
	}
	params = this.Omit(params, MkString("clientOrderId"))
	return this.Call(MkString("privatePostCancelOrder"), this.Extend(request, params))
}

func (this *Poloniex) CancelAllOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	request := MkMap(&VarMap{})
	_ = request
	market := OpCopy(MkUndefined())
	_ = market
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
		*(request).At(MkString("currencyPair")) = *(market).At(MkString("id"))
	}
	response := this.Call(MkString("privatePostCancelAllOrders"), this.Extend(request, params))
	_ = response
	return response
}

func (this *Poloniex) FetchOpenOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	id = id.ToString()
	request := MkMap(&VarMap{"orderNumber": id})
	_ = request
	response := this.Call(MkString("privatePostReturnOrderStatus"), this.Extend(request, params))
	_ = response
	result := this.SafeValue(*(response).At(MkString("result")), id)
	_ = result
	if IsTrue(OpEq2(result, MkUndefined())) {
		panic(NewOrderNotFound(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" order id "), OpAdd(id, MkString(" not found"))))))
	}
	return this.ParseOrder(result)
}

func (this *Poloniex) FetchClosedOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"orderNumber": id})
	_ = request
	response := this.Call(MkString("privatePostReturnOrderTrades"), this.Extend(request, params))
	_ = response
	trades := this.ParseTrades(response)
	_ = trades
	firstTrade := this.SafeValue(trades, MkInteger(0))
	_ = firstTrade
	if IsTrue(OpEq2(firstTrade, MkUndefined())) {
		panic(NewOrderNotFound(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" order id "), OpAdd(id, MkString(" not found"))))))
	}
	symbol = this.SafeString(firstTrade, MkString("symbol"), symbol)
	side := this.SafeString(firstTrade, MkString("side"))
	_ = side
	timestamp := this.SafeNumber(firstTrade, MkString("timestamp"))
	_ = timestamp
	id = this.SafeValue(*(firstTrade).At(MkString("info")), MkString("globalTradeID"), id)
	return this.SafeOrder(MkMap(&VarMap{
		"info":               response,
		"id":                 id,
		"clientOrderId":      this.SafeValue(firstTrade, MkString("clientOrderId")),
		"timestamp":          timestamp,
		"datetime":           this.Iso8601(timestamp),
		"lastTradeTimestamp": MkUndefined(),
		"status":             MkString("closed"),
		"symbol":             symbol,
		"type":               this.SafeString(firstTrade, MkString("type")),
		"timeInForce":        MkUndefined(),
		"postOnly":           MkUndefined(),
		"side":               side,
		"price":              MkUndefined(),
		"stopPrice":          MkUndefined(),
		"cost":               MkUndefined(),
		"average":            MkUndefined(),
		"amount":             MkUndefined(),
		"filled":             MkUndefined(),
		"remaining":          MkUndefined(),
		"trades":             trades,
		"fee":                MkUndefined(),
	}))
}

func (this *Poloniex) FetchOrderStatus(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	orders := this.FetchOpenOrders(symbol, MkUndefined(), MkUndefined(), params)
	_ = orders
	indexed := this.IndexBy(orders, MkString("id"))
	_ = indexed
	return OpTrinary(OpHasMember(id, indexed), MkString("open"), MkString("closed"))
}

func (this *Poloniex) FetchOrderTrades(goArgs ...*Variant) *Variant {
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
	this.LoadMarkets()
	request := MkMap(&VarMap{"orderNumber": id})
	_ = request
	trades := this.Call(MkString("privatePostReturnOrderTrades"), this.Extend(request, params))
	_ = trades
	return this.ParseTrades(trades)
}

func (this *Poloniex) CreateDepositAddress(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	currencyId := OpCopy(MkUndefined())
	_ = currencyId
	currency := OpCopy(MkUndefined())
	_ = currency
	if IsTrue(OpHasMember(code, *this.At(MkString("currencies")))) {
		currency = this.Currency(code)
		currencyId = *(currency).At(MkString("id"))
	} else {
		currencyId = OpCopy(code)
	}
	request := MkMap(&VarMap{"currency": currencyId})
	_ = request
	response := this.Call(MkString("privatePostGenerateNewAddress"), this.Extend(request, params))
	_ = response
	address := OpCopy(MkUndefined())
	_ = address
	tag := OpCopy(MkUndefined())
	_ = tag
	if IsTrue(OpEq2(*(response).At(MkString("success")), MkInteger(1))) {
		address = this.SafeString(response, MkString("response"))
	}
	this.CheckAddress(address)
	if IsTrue(OpNotEq2(currency, MkUndefined())) {
		depositAddress := this.SafeString(*(currency).At(MkString("info")), MkString("depositAddress"))
		_ = depositAddress
		if IsTrue(OpNotEq2(depositAddress, MkUndefined())) {
			tag = OpCopy(address)
			address = OpCopy(depositAddress)
		}
	}
	return MkMap(&VarMap{
		"currency": code,
		"address":  address,
		"tag":      tag,
		"info":     response,
	})
}

func (this *Poloniex) FetchDepositAddress(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privatePostReturnDepositAddresses"), params)
	_ = response
	currencyId := OpCopy(MkUndefined())
	_ = currencyId
	currency := OpCopy(MkUndefined())
	_ = currency
	if IsTrue(OpHasMember(code, *this.At(MkString("currencies")))) {
		currency = this.Currency(code)
		currencyId = *(currency).At(MkString("id"))
	} else {
		currencyId = OpCopy(code)
	}
	address := this.SafeString(response, currencyId)
	_ = address
	tag := OpCopy(MkUndefined())
	_ = tag
	this.CheckAddress(address)
	if IsTrue(OpNotEq2(currency, MkUndefined())) {
		depositAddress := this.SafeString(*(currency).At(MkString("info")), MkString("depositAddress"))
		_ = depositAddress
		if IsTrue(OpNotEq2(depositAddress, MkUndefined())) {
			tag = OpCopy(address)
			address = OpCopy(depositAddress)
		}
	}
	return MkMap(&VarMap{
		"currency": code,
		"address":  address,
		"tag":      tag,
		"info":     response,
	})
}

func (this *Poloniex) Withdraw(goArgs ...*Variant) *Variant {
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
		"currency": *(currency).At(MkString("id")),
		"amount":   amount,
		"address":  address,
	})
	_ = request
	if IsTrue(OpNotEq2(tag, MkUndefined())) {
		*(request).At(MkString("paymentId")) = OpCopy(tag)
	}
	response := this.Call(MkString("privatePostWithdraw"), this.Extend(request, params))
	_ = response
	return MkMap(&VarMap{
		"info": response,
		"id":   this.SafeString(response, MkString("withdrawalNumber")),
	})
}

func (this *Poloniex) FetchTransactionsHelper(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	year := MkInteger(31104000)
	_ = year
	now := this.Seconds()
	_ = now
	start := OpTrinary(OpNotEq2(since, MkUndefined()), parseInt(OpDiv(since, MkInteger(1000))), OpSub(now, OpMulti(MkInteger(10), year)))
	_ = start
	request := MkMap(&VarMap{
		"start": start,
		"end":   now,
	})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("privatePostReturnDepositsWithdrawals"), this.Extend(request, params))
	_ = response
	return response
}

func (this *Poloniex) FetchTransactions(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.FetchTransactionsHelper(code, since, limit, params)
	_ = response
	currency := OpCopy(MkUndefined())
	_ = currency
	if IsTrue(OpNotEq2(code, MkUndefined())) {
		currency = this.Currency(code)
	}
	withdrawals := this.SafeValue(response, MkString("withdrawals"), MkArray(&VarArray{}))
	_ = withdrawals
	deposits := this.SafeValue(response, MkString("deposits"), MkArray(&VarArray{}))
	_ = deposits
	withdrawalTransactions := this.ParseTransactions(withdrawals, currency, since, limit)
	_ = withdrawalTransactions
	depositTransactions := this.ParseTransactions(deposits, currency, since, limit)
	_ = depositTransactions
	transactions := this.ArrayConcat(depositTransactions, withdrawalTransactions)
	_ = transactions
	return this.FilterByCurrencySinceLimit(this.SortBy(transactions, MkString("timestamp")), code, since, limit)
}

func (this *Poloniex) FetchWithdrawals(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	response := this.FetchTransactionsHelper(code, since, limit, params)
	_ = response
	currency := OpCopy(MkUndefined())
	_ = currency
	if IsTrue(OpNotEq2(code, MkUndefined())) {
		currency = this.Currency(code)
	}
	withdrawals := this.SafeValue(response, MkString("withdrawals"), MkArray(&VarArray{}))
	_ = withdrawals
	transactions := this.ParseTransactions(withdrawals, currency, since, limit)
	_ = transactions
	return this.FilterByCurrencySinceLimit(transactions, code, since, limit)
}

func (this *Poloniex) FetchDeposits(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	response := this.FetchTransactionsHelper(code, since, limit, params)
	_ = response
	currency := OpCopy(MkUndefined())
	_ = currency
	if IsTrue(OpNotEq2(code, MkUndefined())) {
		currency = this.Currency(code)
	}
	deposits := this.SafeValue(response, MkString("deposits"), MkArray(&VarArray{}))
	_ = deposits
	transactions := this.ParseTransactions(deposits, currency, since, limit)
	_ = transactions
	return this.FilterByCurrencySinceLimit(transactions, code, since, limit)
}

func (this *Poloniex) ParseTransactionStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"COMPLETE":          MkString("ok"),
		"AWAITING APPROVAL": MkString("pending"),
		"PENDING":           MkString("pending"),
		"PROCESSING":        MkString("pending"),
		"COMPLETE ERROR":    MkString("failed"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Poloniex) ParseTransaction(goArgs ...*Variant) *Variant {
	transaction := GoGetArg(goArgs, 0, MkUndefined())
	_ = transaction
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	timestamp := this.SafeTimestamp(transaction, MkString("timestamp"))
	_ = timestamp
	currencyId := this.SafeString(transaction, MkString("currency"))
	_ = currencyId
	code := this.SafeCurrencyCode(currencyId)
	_ = code
	status := this.SafeString(transaction, MkString("status"), MkString("pending"))
	_ = status
	txid := this.SafeString(transaction, MkString("txid"))
	_ = txid
	if IsTrue(OpNotEq2(status, MkUndefined())) {
		parts := status.Split(MkString(": "))
		_ = parts
		numParts := OpCopy(parts.Length)
		_ = numParts
		status = *(parts).At(MkInteger(0))
		if IsTrue(OpAnd(OpGt(numParts, MkInteger(1)), OpEq2(txid, MkUndefined()))) {
			txid = *(parts).At(MkInteger(1))
		}
		status = this.ParseTransactionStatus(status)
	}
	defaultType := OpTrinary(OpHasMember(MkString("withdrawalNumber"), transaction), MkString("withdrawal"), MkString("deposit"))
	_ = defaultType
	type_ := this.SafeString(transaction, MkString("type"), defaultType)
	_ = type_
	id := this.SafeString2(transaction, MkString("withdrawalNumber"), MkString("depositNumber"))
	_ = id
	amount := this.SafeNumber(transaction, MkString("amount"))
	_ = amount
	address := this.SafeString(transaction, MkString("address"))
	_ = address
	tag := this.SafeString(transaction, MkString("paymentID"))
	_ = tag
	feeCost := this.SafeNumber(transaction, MkString("fee"), MkInteger(0))
	_ = feeCost
	if IsTrue(OpEq2(type_, MkString("withdrawal"))) {
		amount = OpSub(amount, feeCost)
	}
	return MkMap(&VarMap{
		"info":      transaction,
		"id":        id,
		"currency":  code,
		"amount":    amount,
		"address":   address,
		"tag":       tag,
		"status":    status,
		"type":      type_,
		"updated":   MkUndefined(),
		"txid":      txid,
		"timestamp": timestamp,
		"datetime":  this.Iso8601(timestamp),
		"fee": MkMap(&VarMap{
			"currency": code,
			"cost":     feeCost,
		}),
	})
}

func (this *Poloniex) FetchPosition(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"currencyPair": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("privatePostGetMarginPosition"), this.Extend(request, params))
	_ = response
	return response
}

func (this *Poloniex) Nonce(goArgs ...*Variant) *Variant {
	return this.Milliseconds()
}

func (this *Poloniex) Sign(goArgs ...*Variant) *Variant {
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
	url := *(*(*this.At(MkString("urls"))).At(MkString("api"))).At(api)
	_ = url
	query := this.Extend(MkMap(&VarMap{"command": path}), params)
	_ = query
	if IsTrue(OpEq2(api, MkString("public"))) {
		OpAddAssign(&url, OpAdd(MkString("?"), this.Urlencode(query)))
	} else {
		this.CheckRequiredCredentials()
		*(query).At(MkString("nonce")) = this.Nonce()
		body = this.Urlencode(query)
		headers = MkMap(&VarMap{
			"Content-Type": MkString("application/x-www-form-urlencoded"),
			"Key":          *this.At(MkString("apiKey")),
			"Sign":         this.Hmac(this.Encode(body), this.Encode(*this.At(MkString("secret"))), MkString("sha512")),
		})
	}
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Poloniex) HandleErrors(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpHasMember(MkString("error"), response)) {
		message := *(response).At(MkString("error"))
		_ = message
		feedback := OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))
		_ = feedback
		this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("exact")), message, feedback)
		this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("broad")), message, feedback)
		panic(NewExchangeError(feedback))
	}
	return MkUndefined()
}
