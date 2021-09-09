package ccxt

import ()

type Bitvavo struct {
	*ExchangeBase
}

var _ Exchange = (*Bitvavo)(nil)

func init() {
	exchange := &Bitvavo{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Bitvavo) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("bitvavo"),
		"name": MkString("Bitvavo"),
		"countries": MkArray(&VarArray{
			MkString("NL"),
		}),
		"rateLimit": MkInteger(500),
		"version":   MkString("v2"),
		"certified": MkBool(true),
		"pro":       MkBool(true),
		"has": MkMap(&VarMap{
			"CORS":                MkBool(false),
			"publicAPI":           MkBool(true),
			"privateAPI":          MkBool(true),
			"cancelAllOrders":     MkBool(true),
			"cancelOrder":         MkBool(true),
			"createOrder":         MkBool(true),
			"editOrder":           MkBool(true),
			"fetchBalance":        MkBool(true),
			"fetchCurrencies":     MkBool(true),
			"fetchDepositAddress": MkBool(true),
			"fetchDeposits":       MkBool(true),
			"fetchMyTrades":       MkBool(true),
			"fetchOHLCV":          MkBool(true),
			"fetchOpenOrders":     MkBool(true),
			"fetchOrder":          MkBool(true),
			"fetchOrders":         MkBool(true),
			"fetchOrderBook":      MkBool(true),
			"fetchMarkets":        MkBool(true),
			"fetchTicker":         MkBool(true),
			"fetchTickers":        MkBool(true),
			"fetchTime":           MkBool(true),
			"fetchTrades":         MkBool(true),
			"fetchWithdrawals":    MkBool(true),
			"withdraw":            MkBool(true),
		}),
		"timeframes": MkMap(&VarMap{
			"1m":  MkString("1m"),
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
		}),
		"urls": MkMap(&VarMap{
			"logo": MkString("https://user-images.githubusercontent.com/1294454/83165440-2f1cf200-a116-11ea-9046-a255d09fb2ed.jpg"),
			"api": MkMap(&VarMap{
				"public":  MkString("https://api.bitvavo.com"),
				"private": MkString("https://api.bitvavo.com"),
			}),
			"www":      MkString("https://bitvavo.com/"),
			"doc":      MkString("https://docs.bitvavo.com/"),
			"fees":     MkString("https://bitvavo.com/en/fees"),
			"referral": MkString("https://bitvavo.com/?a=24F34952F7"),
		}),
		"api": MkMap(&VarMap{
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("time"),
				MkString("markets"),
				MkString("assets"),
				MkString("{market}/book"),
				MkString("{market}/trades"),
				MkString("{market}/candles"),
				MkString("ticker/price"),
				MkString("ticker/book"),
				MkString("ticker/24h"),
			})}),
			"private": MkMap(&VarMap{
				"get": MkArray(&VarArray{
					MkString("order"),
					MkString("orders"),
					MkString("ordersOpen"),
					MkString("trades"),
					MkString("balance"),
					MkString("deposit"),
					MkString("depositHistory"),
					MkString("withdrawalHistory"),
				}),
				"post": MkArray(&VarArray{
					MkString("order"),
					MkString("withdrawal"),
				}),
				"put": MkArray(&VarArray{
					MkString("order"),
				}),
				"delete": MkArray(&VarArray{
					MkString("order"),
					MkString("orders"),
				}),
			}),
		}),
		"fees": MkMap(&VarMap{"trading": MkMap(&VarMap{
			"tierBased":  MkBool(true),
			"percentage": MkBool(true),
			"taker":      this.ParseNumber(MkString("0.0025")),
			"maker":      this.ParseNumber(MkString("0.002")),
			"tiers": MkMap(&VarMap{
				"taker": MkArray(&VarArray{
					MkArray(&VarArray{
						this.ParseNumber(MkString("0")),
						this.ParseNumber(MkString("0.0025")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("50000")),
						this.ParseNumber(MkString("0.0024")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("100000")),
						this.ParseNumber(MkString("0.0022")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("250000")),
						this.ParseNumber(MkString("0.0020")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("500000")),
						this.ParseNumber(MkString("0.0018")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("1000000")),
						this.ParseNumber(MkString("0.0016")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("2500000")),
						this.ParseNumber(MkString("0.0014")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("5000000")),
						this.ParseNumber(MkString("0.0012")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("10000000")),
						this.ParseNumber(MkString("0.0010")),
					}),
				}),
				"maker": MkArray(&VarArray{
					MkArray(&VarArray{
						this.ParseNumber(MkString("0")),
						this.ParseNumber(MkString("0.0020")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("50000")),
						this.ParseNumber(MkString("0.0015")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("100000")),
						this.ParseNumber(MkString("0.0010")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("250000")),
						this.ParseNumber(MkString("0.0006")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("500000")),
						this.ParseNumber(MkString("0.0003")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("1000000")),
						this.ParseNumber(MkString("0.0001")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("2500000")),
						this.ParseNumber(MkString("-0.0001")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("5000000")),
						this.ParseNumber(MkString("-0.0003")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("10000000")),
						this.ParseNumber(MkString("-0.0005")),
					}),
				}),
			}),
		})}),
		"requiredCredentials": MkMap(&VarMap{
			"apiKey": MkBool(true),
			"secret": MkBool(true),
		}),
		"exceptions": MkMap(&VarMap{
			"exact": MkMap(&VarMap{
				"101": ExchangeError,
				"102": BadRequest,
				"103": RateLimitExceeded,
				"104": RateLimitExceeded,
				"105": PermissionDenied,
				"107": ExchangeNotAvailable,
				"108": ExchangeNotAvailable,
				"109": ExchangeNotAvailable,
				"110": BadRequest,
				"200": BadRequest,
				"201": BadRequest,
				"202": BadRequest,
				"203": BadSymbol,
				"204": BadRequest,
				"205": BadRequest,
				"206": BadRequest,
				"210": InvalidOrder,
				"211": InvalidOrder,
				"212": InvalidOrder,
				"213": InvalidOrder,
				"214": InvalidOrder,
				"215": InvalidOrder,
				"216": InsufficientFunds,
				"217": InvalidOrder,
				"230": ExchangeError,
				"231": ExchangeError,
				"232": BadRequest,
				"233": InvalidOrder,
				"234": InvalidOrder,
				"235": ExchangeError,
				"236": BadRequest,
				"240": OrderNotFound,
				"300": AuthenticationError,
				"301": AuthenticationError,
				"302": AuthenticationError,
				"303": AuthenticationError,
				"304": AuthenticationError,
				"305": AuthenticationError,
				"306": AuthenticationError,
				"307": PermissionDenied,
				"308": AuthenticationError,
				"309": AuthenticationError,
				"310": PermissionDenied,
				"311": PermissionDenied,
				"312": PermissionDenied,
				"315": BadRequest,
				"317": AccountSuspended,
				"400": ExchangeError,
				"401": ExchangeError,
				"402": PermissionDenied,
				"403": PermissionDenied,
				"404": OnMaintenance,
				"405": ExchangeError,
				"406": BadRequest,
				"407": ExchangeError,
				"408": InsufficientFunds,
				"409": InvalidAddress,
				"410": ExchangeError,
				"411": BadRequest,
				"412": InvalidAddress,
				"413": InvalidAddress,
				"414": ExchangeError,
			}),
			"broad": MkMap(&VarMap{
				"start parameter is invalid":   BadRequest,
				"symbol parameter is invalid":  BadSymbol,
				"amount parameter is invalid":  InvalidOrder,
				"orderId parameter is invalid": InvalidOrder,
			}),
		}),
		"options": MkMap(&VarMap{
			"BITVAVO-ACCESS-WINDOW": MkInteger(10000),
			"fetchCurrencies":       MkMap(&VarMap{"expires": MkInteger(1000)}),
		}),
		"precisionMode":    SIGNIFICANT_DIGITS,
		"commonCurrencies": MkMap(&VarMap{"MIOTA": MkString("IOTA")}),
	}))
}

func (this *Bitvavo) CurrencyToPrecision(goArgs ...*Variant) *Variant {
	currency := GoGetArg(goArgs, 0, MkUndefined())
	_ = currency
	fee := GoGetArg(goArgs, 1, MkUndefined())
	_ = fee
	return this.DecimalToPrecision(fee, MkInteger(0), *(*(*this.At(MkString("currencies"))).At(currency)).At(MkString("precision")))
}

func (this *Bitvavo) AmountToPrecision(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	amount := GoGetArg(goArgs, 1, MkUndefined())
	_ = amount
	return this.DecimalToPrecision(amount, TRUNCATE, *(*(*(*this.At(MkString("markets"))).At(symbol)).At(MkString("precision"))).At(MkString("amount")), DECIMAL_PLACES)
}

func (this *Bitvavo) PriceToPrecision(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	price := GoGetArg(goArgs, 1, MkUndefined())
	_ = price
	price = this.DecimalToPrecision(price, ROUND, *(*(*(*this.At(MkString("markets"))).At(symbol)).At(MkString("precision"))).At(MkString("price")), *this.At(MkString("precisionMode")))
	return this.DecimalToPrecision(price, TRUNCATE, MkInteger(8), DECIMAL_PLACES)
}

func (this *Bitvavo) FetchTime(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetTime"), params)
	_ = response
	return this.SafeInteger(response, MkString("time"))
}

func (this *Bitvavo) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetMarkets"), params)
	_ = response
	currencies := this.FetchCurrenciesFromCache(params)
	_ = currencies
	currenciesById := this.IndexBy(currencies, MkString("symbol"))
	_ = currenciesById
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
		market := *(response).At(i)
		_ = market
		id := this.SafeString(market, MkString("market"))
		_ = id
		baseId := this.SafeString(market, MkString("base"))
		_ = baseId
		quoteId := this.SafeString(market, MkString("quote"))
		_ = quoteId
		base := this.SafeCurrencyCode(baseId)
		_ = base
		quote := this.SafeCurrencyCode(quoteId)
		_ = quote
		symbol := OpAdd(base, OpAdd(MkString("/"), quote))
		_ = symbol
		status := this.SafeString(market, MkString("status"))
		_ = status
		active := OpEq2(status, MkString("trading"))
		_ = active
		baseCurrency := this.SafeValue(currenciesById, baseId)
		_ = baseCurrency
		amountPrecision := OpCopy(MkUndefined())
		_ = amountPrecision
		if IsTrue(OpNotEq2(baseCurrency, MkUndefined())) {
			amountPrecision = this.SafeInteger(baseCurrency, MkString("decimals"), MkInteger(8))
		}
		precision := MkMap(&VarMap{
			"price":  this.SafeInteger(market, MkString("pricePrecision")),
			"amount": amountPrecision,
		})
		_ = precision
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
					"min": this.SafeNumber(market, MkString("minOrderInBaseAsset")),
					"max": MkUndefined(),
				}),
				"price": MkMap(&VarMap{
					"min": MkUndefined(),
					"max": MkUndefined(),
				}),
				"cost": MkMap(&VarMap{
					"min": this.SafeNumber(market, MkString("minOrderInQuoteAsset")),
					"max": MkUndefined(),
				}),
			}),
		}))
	}
	return result
}

func (this *Bitvavo) FetchCurrenciesFromCache(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	options := this.SafeValue(*this.At(MkString("options")), MkString("fetchCurrencies"), MkMap(&VarMap{}))
	_ = options
	timestamp := this.SafeInteger(options, MkString("timestamp"))
	_ = timestamp
	expires := this.SafeInteger(options, MkString("expires"), MkInteger(1000))
	_ = expires
	now := this.Milliseconds()
	_ = now
	if IsTrue(OpOr(OpEq2(timestamp, MkUndefined()), OpGt(OpSub(now, timestamp), expires))) {
		response := this.Call(MkString("publicGetAssets"), params)
		_ = response
		*(*this.At(MkString("options"))).At(MkString("fetchCurrencies")) = this.Extend(options, MkMap(&VarMap{
			"response":  response,
			"timestamp": now,
		}))
	}
	return this.SafeValue(*(*this.At(MkString("options"))).At(MkString("fetchCurrencies")), MkString("response"))
}

func (this *Bitvavo) FetchCurrencies(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.FetchCurrenciesFromCache(params)
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
		depositStatus := this.SafeValue(currency, MkString("depositStatus"))
		_ = depositStatus
		deposit := OpEq2(depositStatus, MkString("OK"))
		_ = deposit
		withdrawalStatus := this.SafeValue(currency, MkString("withdrawalStatus"))
		_ = withdrawalStatus
		withdrawal := OpEq2(withdrawalStatus, MkString("OK"))
		_ = withdrawal
		active := OpAnd(deposit, withdrawal)
		_ = active
		name := this.SafeString(currency, MkString("name"))
		_ = name
		precision := this.SafeInteger(currency, MkString("decimals"), MkInteger(8))
		_ = precision
		*(result).At(code) = MkMap(&VarMap{
			"id":        id,
			"info":      currency,
			"code":      code,
			"name":      name,
			"active":    active,
			"fee":       this.SafeNumber(currency, MkString("withdrawalFee")),
			"precision": precision,
			"limits": MkMap(&VarMap{
				"amount": MkMap(&VarMap{
					"min": MkUndefined(),
					"max": MkUndefined(),
				}),
				"withdraw": MkMap(&VarMap{
					"min": this.SafeNumber(currency, MkString("withdrawalMinAmount")),
					"max": MkUndefined(),
				}),
			}),
		})
	}
	return result
}

func (this *Bitvavo) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"market": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("publicGetTicker24h"), this.Extend(request, params))
	_ = response
	return this.ParseTicker(response, market)
}

func (this *Bitvavo) ParseTicker(goArgs ...*Variant) *Variant {
	ticker := GoGetArg(goArgs, 0, MkUndefined())
	_ = ticker
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	marketId := this.SafeString(ticker, MkString("market"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market, MkString("-"))
	_ = symbol
	timestamp := this.SafeInteger(ticker, MkString("timestamp"))
	_ = timestamp
	last := this.SafeNumber(ticker, MkString("last"))
	_ = last
	baseVolume := this.SafeNumber(ticker, MkString("volume"))
	_ = baseVolume
	quoteVolume := this.SafeNumber(ticker, MkString("volumeQuote"))
	_ = quoteVolume
	vwap := this.Vwap(baseVolume, quoteVolume)
	_ = vwap
	change := OpCopy(MkUndefined())
	_ = change
	percentage := OpCopy(MkUndefined())
	_ = percentage
	average := OpCopy(MkUndefined())
	_ = average
	open := this.SafeNumber(ticker, MkString("open"))
	_ = open
	if IsTrue(OpAnd(OpNotEq2(open, MkUndefined()), OpNotEq2(last, MkUndefined()))) {
		change = OpSub(last, open)
		if IsTrue(OpGt(open, MkInteger(0))) {
			percentage = OpDiv(change, OpMulti(open, MkInteger(100)))
		}
		average = OpDiv(this.Sum(open, last), MkInteger(2))
	}
	result := MkMap(&VarMap{
		"symbol":        symbol,
		"timestamp":     timestamp,
		"datetime":      this.Iso8601(timestamp),
		"high":          this.SafeNumber(ticker, MkString("high")),
		"low":           this.SafeNumber(ticker, MkString("low")),
		"bid":           this.SafeNumber(ticker, MkString("bid")),
		"bidVolume":     this.SafeNumber(ticker, MkString("bidSize")),
		"ask":           this.SafeNumber(ticker, MkString("ask")),
		"askVolume":     this.SafeNumber(ticker, MkString("askSize")),
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
	_ = result
	return result
}

func (this *Bitvavo) FetchTickers(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("publicGetTicker24h"), params)
	_ = response
	return this.ParseTickers(response, symbols)
}

func (this *Bitvavo) FetchTrades(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"market": *(market).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("start")) = OpCopy(since)
	}
	response := this.Call(MkString("publicGetMarketTrades"), this.Extend(request, params))
	_ = response
	return this.ParseTrades(response, market, since, limit)
}

func (this *Bitvavo) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
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
	timestamp := this.SafeInteger(trade, MkString("timestamp"))
	_ = timestamp
	side := this.SafeString(trade, MkString("side"))
	_ = side
	id := this.SafeString2(trade, MkString("id"), MkString("fillId"))
	_ = id
	marketId := this.SafeInteger(trade, MkString("market"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market, MkString("-"))
	_ = symbol
	taker := this.SafeValue(trade, MkString("taker"))
	_ = taker
	takerOrMaker := OpCopy(MkUndefined())
	_ = takerOrMaker
	if IsTrue(OpNotEq2(taker, MkUndefined())) {
		takerOrMaker = OpTrinary(taker, MkString("taker"), MkString("maker"))
	}
	feeCost := this.SafeNumber(trade, MkString("fee"))
	_ = feeCost
	fee := OpCopy(MkUndefined())
	_ = fee
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		feeCurrencyId := this.SafeString(trade, MkString("feeCurrency"))
		_ = feeCurrencyId
		feeCurrencyCode := this.SafeCurrencyCode(feeCurrencyId)
		_ = feeCurrencyCode
		fee = MkMap(&VarMap{
			"cost":     feeCost,
			"currency": feeCurrencyCode,
		})
	}
	orderId := this.SafeString(trade, MkString("orderId"))
	_ = orderId
	return MkMap(&VarMap{
		"info":         trade,
		"id":           id,
		"symbol":       symbol,
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"order":        orderId,
		"type":         MkUndefined(),
		"side":         side,
		"takerOrMaker": takerOrMaker,
		"price":        price,
		"amount":       amount,
		"cost":         cost,
		"fee":          fee,
	})
}

func (this *Bitvavo) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"market": this.MarketId(symbol)})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("depth")) = OpCopy(limit)
	}
	response := this.Call(MkString("publicGetMarketBook"), this.Extend(request, params))
	_ = response
	orderbook := this.ParseOrderBook(response, symbol)
	_ = orderbook
	*(orderbook).At(MkString("nonce")) = this.SafeInteger(response, MkString("nonce"))
	return orderbook
}

func (this *Bitvavo) ParseOHLCV(goArgs ...*Variant) *Variant {
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

func (this *Bitvavo) FetchOHLCV(goArgs ...*Variant) *Variant {
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
		"market":   *(market).At(MkString("id")),
		"interval": *(*this.At(MkString("timeframes"))).At(timeframe),
	})
	_ = request
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		duration := this.ParseTimeframe(timeframe)
		_ = duration
		*(request).At(MkString("start")) = OpCopy(since)
		if IsTrue(OpEq2(limit, MkUndefined())) {
			limit = MkInteger(1440)
		}
		*(request).At(MkString("end")) = this.Sum(since, OpMulti(limit, OpMulti(duration, MkInteger(1000))))
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("publicGetMarketCandles"), this.Extend(request, params))
	_ = response
	return this.ParseOHLCVs(response, market, timeframe, since, limit)
}

func (this *Bitvavo) FetchBalance(goArgs ...*Variant) *Variant {
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
	for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
		balance := *(response).At(i)
		_ = balance
		currencyId := this.SafeString(balance, MkString("symbol"))
		_ = currencyId
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		account := this.Account()
		_ = account
		*(account).At(MkString("free")) = this.SafeString(balance, MkString("available"))
		*(account).At(MkString("used")) = this.SafeString(balance, MkString("inOrder"))
		*(result).At(code) = OpCopy(account)
	}
	return this.ParseBalance(result)
}

func (this *Bitvavo) FetchDepositAddress(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{"symbol": *(currency).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("privateGetDeposit"), this.Extend(request, params))
	_ = response
	address := this.SafeString(response, MkString("address"))
	_ = address
	tag := this.SafeString(response, MkString("paymentId"))
	_ = tag
	this.CheckAddress(address)
	return MkMap(&VarMap{
		"currency": code,
		"address":  address,
		"tag":      tag,
		"info":     response,
	})
}

func (this *Bitvavo) CreateOrder(goArgs ...*Variant) *Variant {
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
		"market":    *(market).At(MkString("id")),
		"side":      side,
		"orderType": type_,
	})
	_ = request
	isStopLimit := OpOr(OpEq2(type_, MkString("stopLossLimit")), OpEq2(type_, MkString("takeProfitLimit")))
	_ = isStopLimit
	isStopMarket := OpOr(OpEq2(type_, MkString("stopLoss")), OpEq2(type_, MkString("takeProfit")))
	_ = isStopMarket
	if IsTrue(OpEq2(type_, MkString("market"))) {
		cost := OpCopy(MkUndefined())
		_ = cost
		if IsTrue(OpNotEq2(price, MkUndefined())) {
			cost = OpMulti(amount, price)
		} else {
			cost = this.SafeNumber2(params, MkString("cost"), MkString("amountQuote"))
		}
		if IsTrue(OpNotEq2(cost, MkUndefined())) {
			precision := *(*(market).At(MkString("precision"))).At(MkString("price"))
			_ = precision
			*(request).At(MkString("amountQuote")) = this.DecimalToPrecision(cost, TRUNCATE, precision, *this.At(MkString("precisionMode")))
		} else {
			*(request).At(MkString("amount")) = this.AmountToPrecision(symbol, amount)
		}
		params = this.Omit(params, MkArray(&VarArray{
			MkString("cost"),
			MkString("amountQuote"),
		}))
	} else {
		if IsTrue(OpEq2(type_, MkString("limit"))) {
			*(request).At(MkString("price")) = this.PriceToPrecision(symbol, price)
			*(request).At(MkString("amount")) = this.AmountToPrecision(symbol, amount)
		} else {
			if IsTrue(OpOr(isStopMarket, isStopLimit)) {
				stopPrice := this.SafeNumber2(params, MkString("stopPrice"), MkString("triggerAmount"))
				_ = stopPrice
				if IsTrue(OpEq2(stopPrice, MkUndefined())) {
					if IsTrue(isStopLimit) {
						panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" createOrder requires a stopPrice parameter for a "), OpAdd(type_, MkString(" order"))))))
					} else {
						if IsTrue(isStopMarket) {
							if IsTrue(OpEq2(price, MkUndefined())) {
								panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" createOrder requires a price argument or a stopPrice parameter for a "), OpAdd(type_, MkString(" order"))))))
							} else {
								stopPrice = OpCopy(price)
							}
						}
					}
				}
				if IsTrue(isStopLimit) {
					*(request).At(MkString("price")) = this.PriceToPrecision(symbol, price)
				}
				params = this.Omit(params, MkArray(&VarArray{
					MkString("stopPrice"),
					MkString("triggerAmount"),
				}))
				*(request).At(MkString("triggerAmount")) = this.PriceToPrecision(symbol, stopPrice)
				*(request).At(MkString("triggerType")) = MkString("price")
				*(request).At(MkString("amount")) = this.AmountToPrecision(symbol, amount)
			}
		}
	}
	response := this.Call(MkString("privatePostOrder"), this.Extend(request, params))
	_ = response
	return this.ParseOrder(response, market)
}

func (this *Bitvavo) EditOrder(goArgs ...*Variant) *Variant {
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
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{})
	_ = request
	amountRemaining := this.SafeNumber(params, MkString("amountRemaining"))
	_ = amountRemaining
	params = this.Omit(params, MkString("amountRemaining"))
	if IsTrue(OpNotEq2(price, MkUndefined())) {
		*(request).At(MkString("price")) = this.PriceToPrecision(symbol, price)
	}
	if IsTrue(OpNotEq2(amount, MkUndefined())) {
		*(request).At(MkString("amount")) = this.AmountToPrecision(symbol, amount)
	}
	if IsTrue(OpNotEq2(amountRemaining, MkUndefined())) {
		*(request).At(MkString("amountRemaining")) = this.AmountToPrecision(symbol, amountRemaining)
	}
	request = this.Extend(request, params)
	if IsTrue(*(GoGetKeys(request)).At(MkString("length"))) {
		*(request).At(MkString("orderId")) = OpCopy(id)
		*(request).At(MkString("market")) = *(market).At(MkString("id"))
		response := this.Call(MkString("privatePutOrder"), this.Extend(request, params))
		_ = response
		return this.ParseOrder(response, market)
	} else {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" editOrder() requires an amount argument, or a price argument, or non-empty params"))))
	}
	return MkUndefined()
}

func (this *Bitvavo) CancelOrder(goArgs ...*Variant) *Variant {
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
		"orderId": id,
		"market":  *(market).At(MkString("id")),
	})
	_ = request
	response := this.Call(MkString("privateDeleteOrder"), this.Extend(request, params))
	_ = response
	return this.ParseOrder(response, market)
}

func (this *Bitvavo) CancelAllOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{})
	_ = request
	market := OpCopy(MkUndefined())
	_ = market
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
		*(request).At(MkString("market")) = *(market).At(MkString("id"))
	}
	response := this.Call(MkString("privateDeleteOrders"), this.Extend(request, params))
	_ = response
	return this.ParseOrders(response, market)
}

func (this *Bitvavo) FetchOrder(goArgs ...*Variant) *Variant {
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
		"orderId": id,
		"market":  *(market).At(MkString("id")),
	})
	_ = request
	response := this.Call(MkString("privateGetOrder"), this.Extend(request, params))
	_ = response
	return this.ParseOrder(response, market)
}

func (this *Bitvavo) FetchOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchOrders() requires a symbol argument"))))
	}
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"market": *(market).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("start")) = OpCopy(since)
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("privateGetOrders"), this.Extend(request, params))
	_ = response
	return this.ParseOrders(response, market, since, limit)
}

func (this *Bitvavo) FetchOpenOrders(goArgs ...*Variant) *Variant {
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
	market := OpCopy(MkUndefined())
	_ = market
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
		*(request).At(MkString("market")) = *(market).At(MkString("id"))
	}
	response := this.Call(MkString("privateGetOrdersOpen"), this.Extend(request, params))
	_ = response
	return this.ParseOrders(response, market, since, limit)
}

func (this *Bitvavo) ParseOrderStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"new":                         MkString("open"),
		"canceled":                    MkString("canceled"),
		"canceledAuction":             MkString("canceled"),
		"canceledSelfTradePrevention": MkString("canceled"),
		"canceledIOC":                 MkString("canceled"),
		"canceledFOK":                 MkString("canceled"),
		"canceledMarketProtection":    MkString("canceled"),
		"canceledPostOnly":            MkString("canceled"),
		"filled":                      MkString("closed"),
		"partiallyFilled":             MkString("open"),
		"expired":                     MkString("canceled"),
		"rejected":                    MkString("canceled"),
		"awaitingTrigger":             MkString("open"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Bitvavo) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	id := this.SafeString(order, MkString("orderId"))
	_ = id
	timestamp := this.SafeInteger(order, MkString("created"))
	_ = timestamp
	marketId := this.SafeString(order, MkString("market"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market, MkString("-"))
	_ = symbol
	status := this.ParseOrderStatus(this.SafeString(order, MkString("status")))
	_ = status
	side := this.SafeString(order, MkString("side"))
	_ = side
	type_ := this.SafeString(order, MkString("orderType"))
	_ = type_
	price := this.SafeNumber(order, MkString("price"))
	_ = price
	amount := this.SafeNumber(order, MkString("amount"))
	_ = amount
	remaining := this.SafeNumber(order, MkString("amountRemaining"))
	_ = remaining
	filled := this.SafeNumber(order, MkString("filledAmount"))
	_ = filled
	cost := this.SafeNumber(order, MkString("filledAmountQuote"))
	_ = cost
	fee := OpCopy(MkUndefined())
	_ = fee
	feeCost := this.SafeNumber(order, MkString("feePaid"))
	_ = feeCost
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		feeCurrencyId := this.SafeString(order, MkString("feeCurrency"))
		_ = feeCurrencyId
		feeCurrencyCode := this.SafeCurrencyCode(feeCurrencyId)
		_ = feeCurrencyCode
		fee = MkMap(&VarMap{
			"cost":     feeCost,
			"currency": feeCurrencyCode,
		})
	}
	rawTrades := this.SafeValue(order, MkString("fills"), MkArray(&VarArray{}))
	_ = rawTrades
	trades := this.ParseTrades(rawTrades, market, MkUndefined(), MkUndefined(), MkMap(&VarMap{
		"symbol": symbol,
		"order":  id,
		"side":   side,
		"type":   type_,
	}))
	_ = trades
	timeInForce := this.SafeString(order, MkString("timeInForce"))
	_ = timeInForce
	postOnly := this.SafeValue(order, MkString("postOnly"))
	_ = postOnly
	stopPrice := this.SafeNumber(order, MkString("triggerPrice"))
	_ = stopPrice
	return this.SafeOrder(MkMap(&VarMap{
		"info":               order,
		"id":                 id,
		"clientOrderId":      MkUndefined(),
		"timestamp":          timestamp,
		"datetime":           this.Iso8601(timestamp),
		"lastTradeTimestamp": MkUndefined(),
		"symbol":             symbol,
		"type":               type_,
		"timeInForce":        timeInForce,
		"postOnly":           postOnly,
		"side":               side,
		"price":              price,
		"stopPrice":          stopPrice,
		"amount":             amount,
		"cost":               cost,
		"average":            MkUndefined(),
		"filled":             filled,
		"remaining":          remaining,
		"status":             status,
		"fee":                fee,
		"trades":             trades,
	}))
}

func (this *Bitvavo) FetchMyTrades(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"market": *(market).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("start")) = OpCopy(since)
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("privateGetTrades"), this.Extend(request, params))
	_ = response
	return this.ParseTrades(response, market, since, limit)
}

func (this *Bitvavo) Withdraw(goArgs ...*Variant) *Variant {
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
		"symbol":  *(currency).At(MkString("id")),
		"amount":  this.CurrencyToPrecision(code, amount),
		"address": address,
	})
	_ = request
	if IsTrue(OpNotEq2(tag, MkUndefined())) {
		*(request).At(MkString("paymentId")) = OpCopy(tag)
	}
	response := this.Call(MkString("privatePostWithdrawal"), this.Extend(request, params))
	_ = response
	return this.ParseTransaction(response, currency)
}

func (this *Bitvavo) FetchWithdrawals(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{})
	_ = request
	currency := OpCopy(MkUndefined())
	_ = currency
	if IsTrue(OpNotEq2(code, MkUndefined())) {
		currency = this.Currency(code)
		*(request).At(MkString("symbol")) = *(currency).At(MkString("id"))
	}
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("start")) = OpCopy(since)
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("privateGetWithdrawalHistory"), this.Extend(request, params))
	_ = response
	return this.ParseTransactions(response, currency, since, limit)
}

func (this *Bitvavo) FetchDeposits(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{})
	_ = request
	currency := OpCopy(MkUndefined())
	_ = currency
	if IsTrue(OpNotEq2(code, MkUndefined())) {
		currency = this.Currency(code)
		*(request).At(MkString("symbol")) = *(currency).At(MkString("id"))
	}
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("start")) = OpCopy(since)
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("privateGetDepositHistory"), this.Extend(request, params))
	_ = response
	return this.ParseTransactions(response, currency, since, limit)
}

func (this *Bitvavo) ParseTransactionStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"awaiting_processing":         MkString("pending"),
		"awaiting_email_confirmation": MkString("pending"),
		"awaiting_bitvavo_inspection": MkString("pending"),
		"approved":                    MkString("pending"),
		"sending":                     MkString("pending"),
		"in_mempool":                  MkString("pending"),
		"processed":                   MkString("pending"),
		"completed":                   MkString("ok"),
		"canceled":                    MkString("canceled"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Bitvavo) ParseTransaction(goArgs ...*Variant) *Variant {
	transaction := GoGetArg(goArgs, 0, MkUndefined())
	_ = transaction
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	id := OpCopy(MkUndefined())
	_ = id
	timestamp := this.SafeInteger(transaction, MkString("timestamp"))
	_ = timestamp
	currencyId := this.SafeString(transaction, MkString("symbol"))
	_ = currencyId
	code := this.SafeCurrencyCode(currencyId, currency)
	_ = code
	status := this.ParseTransactionStatus(this.SafeString(transaction, MkString("status")))
	_ = status
	amount := this.SafeNumber(transaction, MkString("amount"))
	_ = amount
	address := this.SafeString(transaction, MkString("address"))
	_ = address
	txid := this.SafeString(transaction, MkString("txId"))
	_ = txid
	fee := OpCopy(MkUndefined())
	_ = fee
	feeCost := this.SafeNumber(transaction, MkString("fee"))
	_ = feeCost
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		fee = MkMap(&VarMap{
			"cost":     feeCost,
			"currency": code,
		})
	}
	type_ := OpCopy(MkUndefined())
	_ = type_
	if IsTrue(OpHasMember(MkString("success"), transaction)) {
		type_ = MkString("withdrawal")
	} else {
		type_ = OpTrinary(OpEq2(status, MkUndefined()), MkString("deposit"), MkString("withdrawal"))
	}
	tag := this.SafeString(transaction, MkString("paymentId"))
	_ = tag
	return MkMap(&VarMap{
		"info":        transaction,
		"id":          id,
		"txid":        txid,
		"timestamp":   timestamp,
		"datetime":    this.Iso8601(timestamp),
		"addressFrom": MkUndefined(),
		"address":     address,
		"addressTo":   address,
		"tagFrom":     MkUndefined(),
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

func (this *Bitvavo) Sign(goArgs ...*Variant) *Variant {
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
	query := this.Omit(params, this.ExtractParams(path))
	_ = query
	url := OpAdd(MkString("/"), OpAdd(*this.At(MkString("version")), OpAdd(MkString("/"), this.ImplodeParams(path, params))))
	_ = url
	getOrDelete := OpOr(OpEq2(method, MkString("GET")), OpEq2(method, MkString("DELETE")))
	_ = getOrDelete
	if IsTrue(getOrDelete) {
		if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
			OpAddAssign(&url, OpAdd(MkString("?"), this.Urlencode(query)))
		}
	}
	if IsTrue(OpEq2(api, MkString("private"))) {
		this.CheckRequiredCredentials()
		payload := MkString("")
		_ = payload
		if IsTrue(OpNot(getOrDelete)) {
			if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
				body = this.Json(query)
				payload = OpCopy(body)
			}
		}
		timestamp := (this.Milliseconds()).Call(MkString("toString"))
		_ = timestamp
		auth := OpAdd(timestamp, OpAdd(method, OpAdd(url, payload)))
		_ = auth
		signature := this.Hmac(this.Encode(auth), this.Encode(*this.At(MkString("secret"))))
		_ = signature
		accessWindow := this.SafeString(*this.At(MkString("options")), MkString("BITVAVO-ACCESS-WINDOW"), MkString("10000"))
		_ = accessWindow
		headers = MkMap(&VarMap{
			"BITVAVO-ACCESS-KEY":       *this.At(MkString("apiKey")),
			"BITVAVO-ACCESS-SIGNATURE": signature,
			"BITVAVO-ACCESS-TIMESTAMP": timestamp,
			"BITVAVO-ACCESS-WINDOW":    accessWindow,
		})
		if IsTrue(OpNot(getOrDelete)) {
			*(headers).At(MkString("Content-Type")) = MkString("application/json")
		}
	}
	url = OpAdd(*(*(*this.At(MkString("urls"))).At(MkString("api"))).At(api), url)
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Bitvavo) HandleErrors(goArgs ...*Variant) *Variant {
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
	errorCode := this.SafeString(response, MkString("errorCode"))
	_ = errorCode
	error := this.SafeString(response, MkString("error"))
	_ = error
	if IsTrue(OpNotEq2(errorCode, MkUndefined())) {
		feedback := OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))
		_ = feedback
		this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("broad")), error, feedback)
		this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("exact")), errorCode, feedback)
		panic(NewExchangeError(feedback))
	}
	return MkUndefined()
}
