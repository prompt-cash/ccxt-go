package ccxt

import ()

type Kraken struct {
	*ExchangeBase
}

var _ Exchange = (*Kraken)(nil)

func init() {
	exchange := &Kraken{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Kraken) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("kraken"),
		"name": MkString("Kraken"),
		"countries": MkArray(&VarArray{
			MkString("US"),
		}),
		"version":   MkString("0"),
		"rateLimit": MkInteger(3000),
		"certified": MkBool(true),
		"pro":       MkBool(true),
		"has": MkMap(&VarMap{
			"cancelAllOrders":      MkBool(true),
			"cancelOrder":          MkBool(true),
			"CORS":                 MkBool(false),
			"createDepositAddress": MkBool(true),
			"createOrder":          MkBool(true),
			"fetchBalance":         MkBool(true),
			"fetchClosedOrders":    MkBool(true),
			"fetchCurrencies":      MkBool(true),
			"fetchDepositAddress":  MkBool(true),
			"fetchDeposits":        MkBool(true),
			"fetchLedger":          MkBool(true),
			"fetchLedgerEntry":     MkBool(true),
			"fetchMarkets":         MkBool(true),
			"fetchMyTrades":        MkBool(true),
			"fetchOHLCV":           MkBool(true),
			"fetchOpenOrders":      MkBool(true),
			"fetchOrder":           MkBool(true),
			"fetchOrderBook":       MkBool(true),
			"fetchOrderTrades":     MkString("emulated"),
			"fetchTicker":          MkBool(true),
			"fetchTickers":         MkBool(true),
			"fetchTime":            MkBool(true),
			"fetchTrades":          MkBool(true),
			"fetchTradingFee":      MkBool(true),
			"fetchTradingFees":     MkBool(true),
			"fetchWithdrawals":     MkBool(true),
			"withdraw":             MkBool(true),
		}),
		"marketsByAltname": MkMap(&VarMap{}),
		"timeframes": MkMap(&VarMap{
			"1m":  MkInteger(1),
			"5m":  MkInteger(5),
			"15m": MkInteger(15),
			"30m": MkInteger(30),
			"1h":  MkInteger(60),
			"4h":  MkInteger(240),
			"1d":  MkInteger(1440),
			"1w":  MkInteger(10080),
			"2w":  MkInteger(21600),
		}),
		"urls": MkMap(&VarMap{
			"logo": MkString("https://user-images.githubusercontent.com/51840849/76173629-fc67fb00-61b1-11ea-84fe-f2de582f58a3.jpg"),
			"api": MkMap(&VarMap{
				"public":  MkString("https://api.kraken.com"),
				"private": MkString("https://api.kraken.com"),
				"zendesk": MkString("https://kraken.zendesk.com/api/v2/help_center/en-us/articles"),
			}),
			"www":  MkString("https://www.kraken.com"),
			"doc":  MkString("https://www.kraken.com/features/api"),
			"fees": MkString("https://www.kraken.com/en-us/features/fee-schedule"),
		}),
		"fees": MkMap(&VarMap{
			"trading": MkMap(&VarMap{
				"tierBased":  MkBool(true),
				"percentage": MkBool(true),
				"taker":      OpDiv(MkNumber(0.26), MkInteger(100)),
				"maker":      OpDiv(MkNumber(0.16), MkInteger(100)),
				"tiers": MkMap(&VarMap{
					"taker": MkArray(&VarArray{
						MkArray(&VarArray{
							MkInteger(0),
							MkNumber(0.0026),
						}),
						MkArray(&VarArray{
							MkInteger(50000),
							MkNumber(0.0024),
						}),
						MkArray(&VarArray{
							MkInteger(100000),
							MkNumber(0.0022),
						}),
						MkArray(&VarArray{
							MkInteger(250000),
							MkNumber(0.0020),
						}),
						MkArray(&VarArray{
							MkInteger(500000),
							MkNumber(0.0018),
						}),
						MkArray(&VarArray{
							MkInteger(1000000),
							MkNumber(0.0016),
						}),
						MkArray(&VarArray{
							MkInteger(2500000),
							MkNumber(0.0014),
						}),
						MkArray(&VarArray{
							MkInteger(5000000),
							MkNumber(0.0012),
						}),
						MkArray(&VarArray{
							MkInteger(10000000),
							MkNumber(0.0001),
						}),
					}),
					"maker": MkArray(&VarArray{
						MkArray(&VarArray{
							MkInteger(0),
							MkNumber(0.0016),
						}),
						MkArray(&VarArray{
							MkInteger(50000),
							MkNumber(0.0014),
						}),
						MkArray(&VarArray{
							MkInteger(100000),
							MkNumber(0.0012),
						}),
						MkArray(&VarArray{
							MkInteger(250000),
							MkNumber(0.0010),
						}),
						MkArray(&VarArray{
							MkInteger(500000),
							MkNumber(0.0008),
						}),
						MkArray(&VarArray{
							MkInteger(1000000),
							MkNumber(0.0006),
						}),
						MkArray(&VarArray{
							MkInteger(2500000),
							MkNumber(0.0004),
						}),
						MkArray(&VarArray{
							MkInteger(5000000),
							MkNumber(0.0002),
						}),
						MkArray(&VarArray{
							MkInteger(10000000),
							MkNumber(0.0),
						}),
					}),
				}),
			}),
			"funding": MkMap(&VarMap{
				"tierBased":  MkBool(false),
				"percentage": MkBool(false),
				"withdraw": MkMap(&VarMap{
					"BTC":  MkNumber(0.001),
					"ETH":  MkNumber(0.005),
					"XRP":  MkNumber(0.02),
					"XLM":  MkNumber(0.00002),
					"LTC":  MkNumber(0.02),
					"DOGE": MkInteger(2),
					"ZEC":  MkNumber(0.00010),
					"ICN":  MkNumber(0.02),
					"REP":  MkNumber(0.01),
					"ETC":  MkNumber(0.005),
					"MLN":  MkNumber(0.003),
					"XMR":  MkNumber(0.05),
					"DASH": MkNumber(0.005),
					"GNO":  MkNumber(0.01),
					"EOS":  MkNumber(0.5),
					"BCH":  MkNumber(0.001),
					"XTZ":  MkNumber(0.05),
					"USD":  MkInteger(5),
					"EUR":  MkInteger(5),
					"CAD":  MkInteger(10),
					"JPY":  MkInteger(300),
				}),
				"deposit": MkMap(&VarMap{
					"BTC":  MkInteger(0),
					"ETH":  MkInteger(0),
					"XRP":  MkInteger(0),
					"XLM":  MkInteger(0),
					"LTC":  MkInteger(0),
					"DOGE": MkInteger(0),
					"ZEC":  MkInteger(0),
					"ICN":  MkInteger(0),
					"REP":  MkInteger(0),
					"ETC":  MkInteger(0),
					"MLN":  MkInteger(0),
					"XMR":  MkInteger(0),
					"DASH": MkInteger(0),
					"GNO":  MkInteger(0),
					"EOS":  MkInteger(0),
					"BCH":  MkInteger(0),
					"XTZ":  MkNumber(0.05),
					"USD":  MkInteger(5),
					"EUR":  MkInteger(0),
					"CAD":  MkInteger(5),
					"JPY":  MkInteger(0),
				}),
			}),
		}),
		"api": MkMap(&VarMap{
			"zendesk": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("360000292886"),
				MkString("201893608"),
			})}),
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("Assets"),
				MkString("AssetPairs"),
				MkString("Depth"),
				MkString("OHLC"),
				MkString("Spread"),
				MkString("Ticker"),
				MkString("Time"),
				MkString("Trades"),
			})}),
			"private": MkMap(&VarMap{"post": MkArray(&VarArray{
				MkString("AddOrder"),
				MkString("AddExport"),
				MkString("Balance"),
				MkString("CancelAll"),
				MkString("CancelOrder"),
				MkString("ClosedOrders"),
				MkString("DepositAddresses"),
				MkString("DepositMethods"),
				MkString("DepositStatus"),
				MkString("ExportStatus"),
				MkString("GetWebSocketsToken"),
				MkString("Ledgers"),
				MkString("OpenOrders"),
				MkString("OpenPositions"),
				MkString("QueryLedgers"),
				MkString("QueryOrders"),
				MkString("QueryTrades"),
				MkString("RetrieveExport"),
				MkString("RemoveExport"),
				MkString("TradeBalance"),
				MkString("TradesHistory"),
				MkString("TradeVolume"),
				MkString("Withdraw"),
				MkString("WithdrawCancel"),
				MkString("WithdrawInfo"),
				MkString("WithdrawStatus"),
			})}),
		}),
		"commonCurrencies": MkMap(&VarMap{
			"XBT":   MkString("BTC"),
			"XBT.M": MkString("BTC.M"),
			"XDG":   MkString("DOGE"),
			"REPV2": MkString("REP"),
			"REP":   MkString("REPV1"),
		}),
		"options": MkMap(&VarMap{
			"cacheDepositMethodsOnFetchDepositAddress": MkBool(true),
			"depositMethods":      MkMap(&VarMap{}),
			"delistedMarketsById": MkMap(&VarMap{}),
			"inactiveCurrencies": MkArray(&VarArray{
				MkString("CAD"),
				MkString("USD"),
				MkString("JPY"),
				MkString("GBP"),
			}),
		}),
		"exceptions": MkMap(&VarMap{
			"EQuery:Invalid asset pair":     BadSymbol,
			"EAPI:Invalid key":              AuthenticationError,
			"EFunding:Unknown withdraw key": ExchangeError,
			"EFunding:Invalid amount":       InsufficientFunds,
			"EService:Unavailable":          ExchangeNotAvailable,
			"EDatabase:Internal error":      ExchangeNotAvailable,
			"EService:Busy":                 ExchangeNotAvailable,
			"EQuery:Unknown asset":          ExchangeError,
			"EAPI:Rate limit exceeded":      DDoSProtection,
			"EOrder:Rate limit exceeded":    DDoSProtection,
			"EGeneral:Internal error":       ExchangeNotAvailable,
			"EGeneral:Temporary lockout":    DDoSProtection,
			"EGeneral:Permission denied":    PermissionDenied,
			"EOrder:Unknown order":          InvalidOrder,
			"EOrder:Order minimum not met":  InvalidOrder,
		}),
	}))
}

func (this *Kraken) CostToPrecision(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	cost := GoGetArg(goArgs, 1, MkUndefined())
	_ = cost
	return this.DecimalToPrecision(cost, TRUNCATE, *(*(*(*this.At(MkString("markets"))).At(symbol)).At(MkString("precision"))).At(MkString("price")), DECIMAL_PLACES)
}

func (this *Kraken) FeeToPrecision(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	fee := GoGetArg(goArgs, 1, MkUndefined())
	_ = fee
	return this.DecimalToPrecision(fee, TRUNCATE, *(*(*(*this.At(MkString("markets"))).At(symbol)).At(MkString("precision"))).At(MkString("amount")), DECIMAL_PLACES)
}

func (this *Kraken) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetAssetPairs"), params)
	_ = response
	markets := this.SafeValue(response, MkString("result"), MkMap(&VarMap{}))
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
		baseId := this.SafeString(market, MkString("base"))
		_ = baseId
		quoteId := this.SafeString(market, MkString("quote"))
		_ = quoteId
		base := this.SafeCurrencyCode(baseId)
		_ = base
		quote := this.SafeCurrencyCode(quoteId)
		_ = quote
		darkpool := OpGtEq(id.IndexOf(MkString(".d")), MkInteger(0))
		_ = darkpool
		altname := this.SafeString(market, MkString("altname"))
		_ = altname
		symbol := OpTrinary(darkpool, altname, OpAdd(base, OpAdd(MkString("/"), quote)))
		_ = symbol
		makerFees := this.SafeValue(market, MkString("fees_maker"), MkArray(&VarArray{}))
		_ = makerFees
		firstMakerFee := this.SafeValue(makerFees, MkInteger(0), MkArray(&VarArray{}))
		_ = firstMakerFee
		firstMakerFeeRate := this.SafeNumber(firstMakerFee, MkInteger(1))
		_ = firstMakerFeeRate
		maker := OpCopy(MkUndefined())
		_ = maker
		if IsTrue(OpNotEq2(firstMakerFeeRate, MkUndefined())) {
			maker = OpDiv(parseFloat(firstMakerFeeRate), MkInteger(100))
		}
		takerFees := this.SafeValue(market, MkString("fees"), MkArray(&VarArray{}))
		_ = takerFees
		firstTakerFee := this.SafeValue(takerFees, MkInteger(0), MkArray(&VarArray{}))
		_ = firstTakerFee
		firstTakerFeeRate := this.SafeNumber(firstTakerFee, MkInteger(1))
		_ = firstTakerFeeRate
		taker := OpCopy(MkUndefined())
		_ = taker
		if IsTrue(OpNotEq2(firstTakerFeeRate, MkUndefined())) {
			taker = OpDiv(parseFloat(firstTakerFeeRate), MkInteger(100))
		}
		precision := MkMap(&VarMap{
			"amount": this.SafeInteger(market, MkString("lot_decimals")),
			"price":  this.SafeInteger(market, MkString("pair_decimals")),
		})
		_ = precision
		minAmount := this.SafeNumber(market, MkString("ordermin"))
		_ = minAmount
		result.Push(MkMap(&VarMap{
			"id":        id,
			"symbol":    symbol,
			"base":      base,
			"quote":     quote,
			"baseId":    baseId,
			"quoteId":   quoteId,
			"darkpool":  darkpool,
			"info":      market,
			"altname":   *(market).At(MkString("altname")),
			"maker":     maker,
			"taker":     taker,
			"active":    MkBool(true),
			"precision": precision,
			"limits": MkMap(&VarMap{
				"amount": MkMap(&VarMap{
					"min": minAmount,
					"max": Math.Pow(MkInteger(10), *(precision).At(MkString("amount"))),
				}),
				"price": MkMap(&VarMap{
					"min": Math.Pow(MkInteger(10), OpNeg(*(precision).At(MkString("price")))),
					"max": MkUndefined(),
				}),
				"cost": MkMap(&VarMap{
					"min": MkInteger(0),
					"max": MkUndefined(),
				}),
			}),
		}))
	}
	result = this.AppendInactiveMarkets(result)
	*this.At(MkString("marketsByAltname")) = this.IndexBy(result, MkString("altname"))
	return result
}

func (this *Kraken) SafeCurrency(goArgs ...*Variant) *Variant {
	currencyId := GoGetArg(goArgs, 0, MkUndefined())
	_ = currencyId
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	if IsTrue(OpGt(currencyId.Length, MkInteger(3))) {
		if IsTrue(OpOr(OpEq2(currencyId.IndexOf(MkString("X")), MkInteger(0)), OpEq2(currencyId.IndexOf(MkString("Z")), MkInteger(0)))) {
			if IsTrue(OpGt(currencyId.IndexOf(MkString(".")), MkInteger(0))) {
				return this.SafeCurrency(currencyId, currency)
			} else {
				currencyId = currencyId.Slice(MkInteger(1))
			}
		}
	}
	return this.SafeCurrency(currencyId, currency)
}

func (this *Kraken) AppendInactiveMarkets(goArgs ...*Variant) *Variant {
	result := GoGetArg(goArgs, 0, MkUndefined())
	_ = result
	precision := MkMap(&VarMap{
		"amount": MkInteger(8),
		"price":  MkInteger(8),
	})
	_ = precision
	costLimits := MkMap(&VarMap{
		"min": MkInteger(0),
		"max": MkUndefined(),
	})
	_ = costLimits
	priceLimits := MkMap(&VarMap{
		"min": Math.Pow(MkInteger(10), OpNeg(*(precision).At(MkString("price")))),
		"max": MkUndefined(),
	})
	_ = priceLimits
	amountLimits := MkMap(&VarMap{
		"min": Math.Pow(MkInteger(10), OpNeg(*(precision).At(MkString("amount")))),
		"max": Math.Pow(MkInteger(10), *(precision).At(MkString("amount"))),
	})
	_ = amountLimits
	limits := MkMap(&VarMap{
		"amount": amountLimits,
		"price":  priceLimits,
		"cost":   costLimits,
	})
	_ = limits
	defaults := MkMap(&VarMap{
		"darkpool":  MkBool(false),
		"info":      MkUndefined(),
		"maker":     MkUndefined(),
		"taker":     MkUndefined(),
		"active":    MkBool(false),
		"precision": precision,
		"limits":    limits,
	})
	_ = defaults
	markets := MkArray(&VarArray{})
	_ = markets
	for i := MkInteger(0); IsTrue(OpLw(i, markets.Length)); OpIncr(&i) {
		result.Push(this.Extend(defaults, *(markets).At(i)))
	}
	return result
}

func (this *Kraken) FetchCurrencies(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetAssets"), params)
	_ = response
	currencies := this.SafeValue(response, MkString("result"))
	_ = currencies
	ids := GoGetKeys(currencies)
	_ = ids
	result := MkMap(&VarMap{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, ids.Length)); OpIncr(&i) {
		id := *(ids).At(i)
		_ = id
		currency := *(currencies).At(id)
		_ = currency
		code := this.SafeCurrencyCode(this.SafeString(currency, MkString("altname")))
		_ = code
		precision := this.SafeInteger(currency, MkString("decimals"))
		_ = precision
		active := OpNot(this.InArray(code, *(*this.At(MkString("options"))).At(MkString("inactiveCurrencies"))))
		_ = active
		*(result).At(code) = MkMap(&VarMap{
			"id":        id,
			"code":      code,
			"info":      currency,
			"name":      code,
			"active":    active,
			"fee":       MkUndefined(),
			"precision": precision,
			"limits": MkMap(&VarMap{
				"amount": MkMap(&VarMap{
					"min": Math.Pow(MkInteger(10), OpNeg(precision)),
					"max": Math.Pow(MkInteger(10), precision),
				}),
				"withdraw": MkMap(&VarMap{
					"min": MkUndefined(),
					"max": Math.Pow(MkInteger(10), precision),
				}),
			}),
		})
	}
	return result
}

func (this *Kraken) FetchTradingFees(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	this.CheckRequiredCredentials()
	response := this.Call(MkString("privatePostTradeVolume"), params)
	_ = response
	tradedVolume := this.SafeNumber(*(response).At(MkString("result")), MkString("volume"))
	_ = tradedVolume
	tiers := *(*(*this.At(MkString("fees"))).At(MkString("trading"))).At(MkString("tiers"))
	_ = tiers
	taker := *(*(tiers).At(MkString("taker"))).At(MkInteger(1))
	_ = taker
	maker := *(*(tiers).At(MkString("maker"))).At(MkInteger(1))
	_ = maker
	for i := MkInteger(0); IsTrue(OpLw(i, *(*(tiers).At(MkString("taker"))).At(MkString("length")))); OpIncr(&i) {
		if IsTrue(OpGtEq(tradedVolume, *(*(*(tiers).At(MkString("taker"))).At(i)).At(MkInteger(0)))) {
			taker = *(*(*(tiers).At(MkString("taker"))).At(i)).At(MkInteger(1))
		}
	}
	for i := MkInteger(0); IsTrue(OpLw(i, *(*(tiers).At(MkString("maker"))).At(MkString("length")))); OpIncr(&i) {
		if IsTrue(OpGtEq(tradedVolume, *(*(*(tiers).At(MkString("maker"))).At(i)).At(MkInteger(0)))) {
			maker = *(*(*(tiers).At(MkString("maker"))).At(i)).At(MkInteger(1))
		}
	}
	return MkMap(&VarMap{
		"info":  response,
		"maker": maker,
		"taker": taker,
	})
}

func (this *Kraken) ParseBidAsk(goArgs ...*Variant) *Variant {
	bidask := GoGetArg(goArgs, 0, MkUndefined())
	_ = bidask
	priceKey := GoGetArg(goArgs, 1, MkInteger(0))
	_ = priceKey
	amountKey := GoGetArg(goArgs, 2, MkInteger(1))
	_ = amountKey
	price := this.SafeNumber(bidask, priceKey)
	_ = price
	amount := this.SafeNumber(bidask, amountKey)
	_ = amount
	timestamp := this.SafeInteger(bidask, MkInteger(2))
	_ = timestamp
	return MkArray(&VarArray{
		price,
		amount,
		timestamp,
	})
}

func (this *Kraken) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	if IsTrue(*(market).At(MkString("darkpool"))) {
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" does not provide an order book for darkpool symbol "), symbol))))
	}
	request := MkMap(&VarMap{"pair": *(market).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("count")) = OpCopy(limit)
	}
	response := this.Call(MkString("publicGetDepth"), this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"), MkMap(&VarMap{}))
	_ = result
	orderbook := this.SafeValue(result, *(market).At(MkString("id")))
	_ = orderbook
	marketInfo := this.SafeValue(market, MkString("info"), MkMap(&VarMap{}))
	_ = marketInfo
	wsName := this.SafeValue(marketInfo, MkString("wsname"))
	_ = wsName
	if IsTrue(OpNotEq2(wsName, MkUndefined())) {
		orderbook = this.SafeValue(result, wsName, orderbook)
	}
	return this.ParseOrderBook(orderbook, symbol)
}

func (this *Kraken) ParseTicker(goArgs ...*Variant) *Variant {
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
	baseVolume := parseFloat(*(*(ticker).At(MkString("v"))).At(MkInteger(1)))
	_ = baseVolume
	vwap := parseFloat(*(*(ticker).At(MkString("p"))).At(MkInteger(1)))
	_ = vwap
	quoteVolume := OpCopy(MkUndefined())
	_ = quoteVolume
	if IsTrue(OpAnd(OpNotEq2(baseVolume, MkUndefined()), OpNotEq2(vwap, MkUndefined()))) {
		quoteVolume = OpMulti(baseVolume, vwap)
	}
	last := parseFloat(*(*(ticker).At(MkString("c"))).At(MkInteger(0)))
	_ = last
	return MkMap(&VarMap{
		"symbol":        symbol,
		"timestamp":     timestamp,
		"datetime":      this.Iso8601(timestamp),
		"high":          parseFloat(*(*(ticker).At(MkString("h"))).At(MkInteger(1))),
		"low":           parseFloat(*(*(ticker).At(MkString("l"))).At(MkInteger(1))),
		"bid":           parseFloat(*(*(ticker).At(MkString("b"))).At(MkInteger(0))),
		"bidVolume":     MkUndefined(),
		"ask":           parseFloat(*(*(ticker).At(MkString("a"))).At(MkInteger(0))),
		"askVolume":     MkUndefined(),
		"vwap":          vwap,
		"open":          this.SafeNumber(ticker, MkString("o")),
		"close":         last,
		"last":          last,
		"previousClose": MkUndefined(),
		"change":        MkUndefined(),
		"percentage":    MkUndefined(),
		"average":       MkUndefined(),
		"baseVolume":    baseVolume,
		"quoteVolume":   quoteVolume,
		"info":          ticker,
	})
}

func (this *Kraken) FetchTickers(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	symbols = OpTrinary(OpEq2(symbols, MkUndefined()), *this.At(MkString("symbols")), symbols)
	marketIds := MkArray(&VarArray{})
	_ = marketIds
	for i := MkInteger(0); IsTrue(OpLw(i, symbols.Length)); OpIncr(&i) {
		symbol := *(symbols).At(i)
		_ = symbol
		market := *(*this.At(MkString("markets"))).At(symbol)
		_ = market
		if IsTrue(OpAnd(*(market).At(MkString("active")), OpNot(*(market).At(MkString("darkpool"))))) {
			marketIds.Push(*(market).At(MkString("id")))
		}
	}
	request := MkMap(&VarMap{"pair": marketIds.Join(MkString(","))})
	_ = request
	response := this.Call(MkString("publicGetTicker"), this.Extend(request, params))
	_ = response
	tickers := *(response).At(MkString("result"))
	_ = tickers
	ids := GoGetKeys(tickers)
	_ = ids
	result := MkMap(&VarMap{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, ids.Length)); OpIncr(&i) {
		id := *(ids).At(i)
		_ = id
		market := *(*this.At(MkString("markets_by_id"))).At(id)
		_ = market
		symbol := *(market).At(MkString("symbol"))
		_ = symbol
		ticker := *(tickers).At(id)
		_ = ticker
		*(result).At(symbol) = this.ParseTicker(ticker, market)
	}
	return this.FilterByArray(result, MkString("symbol"), symbols)
}

func (this *Kraken) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	darkpool := OpGtEq(symbol.IndexOf(MkString(".d")), MkInteger(0))
	_ = darkpool
	if IsTrue(darkpool) {
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" does not provide a ticker for darkpool symbol "), symbol))))
	}
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"pair": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("publicGetTicker"), this.Extend(request, params))
	_ = response
	ticker := *(*(response).At(MkString("result"))).At(*(market).At(MkString("id")))
	_ = ticker
	return this.ParseTicker(ticker, market)
}

func (this *Kraken) ParseOHLCV(goArgs ...*Variant) *Variant {
	ohlcv := GoGetArg(goArgs, 0, MkUndefined())
	_ = ohlcv
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	return MkArray(&VarArray{
		this.SafeTimestamp(ohlcv, MkInteger(0)),
		this.SafeNumber(ohlcv, MkInteger(1)),
		this.SafeNumber(ohlcv, MkInteger(2)),
		this.SafeNumber(ohlcv, MkInteger(3)),
		this.SafeNumber(ohlcv, MkInteger(4)),
		this.SafeNumber(ohlcv, MkInteger(6)),
	})
}

func (this *Kraken) FetchOHLCV(goArgs ...*Variant) *Variant {
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
		"pair":     *(market).At(MkString("id")),
		"interval": *(*this.At(MkString("timeframes"))).At(timeframe),
	})
	_ = request
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("since")) = parseInt(OpDiv(OpSub(since, MkInteger(1)), MkInteger(1000)))
	}
	response := this.Call(MkString("publicGetOHLC"), this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"), MkMap(&VarMap{}))
	_ = result
	ohlcvs := this.SafeValue(result, *(market).At(MkString("id")), MkArray(&VarArray{}))
	_ = ohlcvs
	return this.ParseOHLCVs(ohlcvs, market, timeframe, since, limit)
}

func (this *Kraken) ParseLedgerEntryType(goArgs ...*Variant) *Variant {
	type_ := GoGetArg(goArgs, 0, MkUndefined())
	_ = type_
	types := MkMap(&VarMap{
		"trade":      MkString("trade"),
		"withdrawal": MkString("transaction"),
		"deposit":    MkString("transaction"),
		"transfer":   MkString("transfer"),
		"margin":     MkString("margin"),
	})
	_ = types
	return this.SafeString(types, type_, type_)
}

func (this *Kraken) ParseLedgerEntry(goArgs ...*Variant) *Variant {
	item := GoGetArg(goArgs, 0, MkUndefined())
	_ = item
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	id := this.SafeString(item, MkString("id"))
	_ = id
	direction := OpCopy(MkUndefined())
	_ = direction
	account := OpCopy(MkUndefined())
	_ = account
	referenceId := this.SafeString(item, MkString("refid"))
	_ = referenceId
	referenceAccount := OpCopy(MkUndefined())
	_ = referenceAccount
	type_ := this.ParseLedgerEntryType(this.SafeString(item, MkString("type")))
	_ = type_
	code := this.SafeCurrencyCode(this.SafeString(item, MkString("asset")), currency)
	_ = code
	amount := this.SafeNumber(item, MkString("amount"))
	_ = amount
	if IsTrue(OpLw(amount, MkInteger(0))) {
		direction = MkString("out")
		amount = Math.Abs(amount)
	} else {
		direction = MkString("in")
	}
	time := this.SafeNumber(item, MkString("time"))
	_ = time
	timestamp := OpCopy(MkUndefined())
	_ = timestamp
	if IsTrue(OpNotEq2(time, MkUndefined())) {
		timestamp = parseInt(OpMulti(time, MkInteger(1000)))
	}
	fee := MkMap(&VarMap{
		"cost":     this.SafeNumber(item, MkString("fee")),
		"currency": code,
	})
	_ = fee
	before := OpCopy(MkUndefined())
	_ = before
	after := this.SafeNumber(item, MkString("balance"))
	_ = after
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
		"fee":              fee,
	})
}

func (this *Kraken) FetchLedger(goArgs ...*Variant) *Variant {
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
		*(request).At(MkString("asset")) = *(currency).At(MkString("id"))
	}
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("start")) = parseInt(OpDiv(since, MkInteger(1000)))
	}
	response := this.Call(MkString("privatePostLedgers"), this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"), MkMap(&VarMap{}))
	_ = result
	ledger := this.SafeValue(result, MkString("ledger"), MkMap(&VarMap{}))
	_ = ledger
	keys := GoGetKeys(ledger)
	_ = keys
	items := MkArray(&VarArray{})
	_ = items
	for i := MkInteger(0); IsTrue(OpLw(i, keys.Length)); OpIncr(&i) {
		key := *(keys).At(i)
		_ = key
		value := *(ledger).At(key)
		_ = value
		*(value).At(MkString("id")) = OpCopy(key)
		items.Push(value)
	}
	return this.ParseLedger(items, currency, since, limit)
}

func (this *Kraken) FetchLedgerEntriesByIds(goArgs ...*Variant) *Variant {
	ids := GoGetArg(goArgs, 0, MkUndefined())
	_ = ids
	code := GoGetArg(goArgs, 1, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	ids = ids.Join(MkString(","))
	request := this.Extend(MkMap(&VarMap{"id": ids}), params)
	_ = request
	response := this.Call(MkString("privatePostQueryLedgers"), request)
	_ = response
	result := *(response).At(MkString("result"))
	_ = result
	keys := GoGetKeys(result)
	_ = keys
	items := MkArray(&VarArray{})
	_ = items
	for i := MkInteger(0); IsTrue(OpLw(i, keys.Length)); OpIncr(&i) {
		key := *(keys).At(i)
		_ = key
		value := *(result).At(key)
		_ = value
		*(value).At(MkString("id")) = OpCopy(key)
		items.Push(value)
	}
	return this.ParseLedger(items)
}

func (this *Kraken) FetchLedgerEntry(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	code := GoGetArg(goArgs, 1, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	items := this.FetchLedgerEntriesByIds(MkArray(&VarArray{
		id,
	}), code, params)
	_ = items
	return *(items).At(MkInteger(0))
}

func (this *Kraken) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := OpCopy(MkUndefined())
	_ = timestamp
	side := OpCopy(MkUndefined())
	_ = side
	type_ := OpCopy(MkUndefined())
	_ = type_
	priceString := OpCopy(MkUndefined())
	_ = priceString
	amountString := OpCopy(MkUndefined())
	_ = amountString
	id := OpCopy(MkUndefined())
	_ = id
	orderId := OpCopy(MkUndefined())
	_ = orderId
	fee := OpCopy(MkUndefined())
	_ = fee
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(GoIsArray(trade)) {
		timestamp = this.SafeTimestamp(trade, MkInteger(2))
		side = OpTrinary(OpEq2(*(trade).At(MkInteger(3)), MkString("s")), MkString("sell"), MkString("buy"))
		type_ = OpTrinary(OpEq2(*(trade).At(MkInteger(4)), MkString("l")), MkString("limit"), MkString("market"))
		priceString = this.SafeString(trade, MkInteger(0))
		amountString = this.SafeString(trade, MkInteger(1))
		tradeLength := OpCopy(trade.Length)
		_ = tradeLength
		if IsTrue(OpGt(tradeLength, MkInteger(6))) {
			id = this.SafeString(trade, MkInteger(6))
		}
	} else {
		if IsTrue(OpEq2(OpType(trade), MkString("string"))) {
			id = OpCopy(trade)
		} else {
			if IsTrue(OpHasMember(MkString("ordertxid"), trade)) {
				marketId := this.SafeString(trade, MkString("pair"))
				_ = marketId
				foundMarket := this.FindMarketByAltnameOrId(marketId)
				_ = foundMarket
				if IsTrue(OpNotEq2(foundMarket, MkUndefined())) {
					market = OpCopy(foundMarket)
				} else {
					if IsTrue(OpNotEq2(marketId, MkUndefined())) {
						market = this.GetDelistedMarketById(marketId)
					}
				}
				orderId = this.SafeString(trade, MkString("ordertxid"))
				id = this.SafeString2(trade, MkString("id"), MkString("postxid"))
				timestamp = this.SafeTimestamp(trade, MkString("time"))
				side = this.SafeString(trade, MkString("type"))
				type_ = this.SafeString(trade, MkString("ordertype"))
				priceString = this.SafeString(trade, MkString("price"))
				amountString = this.SafeString(trade, MkString("vol"))
				if IsTrue(OpHasMember(MkString("fee"), trade)) {
					currency := OpCopy(MkUndefined())
					_ = currency
					if IsTrue(OpNotEq2(market, MkUndefined())) {
						currency = *(market).At(MkString("quote"))
					}
					fee = MkMap(&VarMap{
						"cost":     this.SafeNumber(trade, MkString("fee")),
						"currency": currency,
					})
				}
			}
		}
	}
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
		"id":           id,
		"order":        orderId,
		"info":         trade,
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       symbol,
		"type":         type_,
		"side":         side,
		"takerOrMaker": MkUndefined(),
		"price":        price,
		"amount":       amount,
		"cost":         cost,
		"fee":          fee,
	})
}

func (this *Kraken) FetchTrades(goArgs ...*Variant) *Variant {
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
	id := *(market).At(MkString("id"))
	_ = id
	request := MkMap(&VarMap{"pair": id})
	_ = request
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("since")) = OpMulti(since, MkInteger(1e6))
		*(request).At(MkString("since")) = OpAdd(since.ToString(), MkString("000000"))
	}
	if IsTrue(OpAnd(OpNotEq2(limit, MkUndefined()), OpNotEq2(limit, MkInteger(1000)))) {
		fetchTradesWarning := this.SafeValue(*this.At(MkString("options")), MkString("fetchTradesWarning"), MkBool(true))
		_ = fetchTradesWarning
		if IsTrue(fetchTradesWarning) {
			panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" fetchTrades() cannot serve "), OpAdd(limit.ToString(), MkString(" trades without breaking the pagination, see https://github.com/ccxt/ccxt/issues/5698 for more details. Set exchange.options['fetchTradesWarning'] to acknowledge this warning and silence it."))))))
		}
	}
	response := this.Call(MkString("publicGetTrades"), this.Extend(request, params))
	_ = response
	result := *(response).At(MkString("result"))
	_ = result
	trades := *(result).At(id)
	_ = trades
	length := OpCopy(trades.Length)
	_ = length
	if IsTrue(OpLwEq(length, MkInteger(0))) {
		return MkArray(&VarArray{})
	}
	lastTrade := *(trades).At(OpSub(length, MkInteger(1)))
	_ = lastTrade
	lastTradeId := this.SafeString(result, MkString("last"))
	_ = lastTradeId
	lastTrade.Push(lastTradeId)
	return this.ParseTrades(trades, market, since, limit)
}

func (this *Kraken) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privatePostBalance"), params)
	_ = response
	balances := this.SafeValue(response, MkString("result"), MkMap(&VarMap{}))
	_ = balances
	result := MkMap(&VarMap{
		"info":      response,
		"timestamp": MkUndefined(),
		"datetime":  MkUndefined(),
	})
	_ = result
	currencyIds := GoGetKeys(balances)
	_ = currencyIds
	for i := MkInteger(0); IsTrue(OpLw(i, currencyIds.Length)); OpIncr(&i) {
		currencyId := *(currencyIds).At(i)
		_ = currencyId
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		account := this.Account()
		_ = account
		*(account).At(MkString("total")) = this.SafeString(balances, currencyId)
		*(result).At(code) = OpCopy(account)
	}
	return this.ParseBalance(result)
}

func (this *Kraken) CreateOrder(goArgs ...*Variant) *Variant {
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
		"pair":      *(market).At(MkString("id")),
		"type":      side,
		"ordertype": type_,
		"volume":    this.AmountToPrecision(symbol, amount),
	})
	_ = request
	clientOrderId := this.SafeString2(params, MkString("userref"), MkString("clientOrderId"))
	_ = clientOrderId
	params = this.Omit(params, MkArray(&VarArray{
		MkString("userref"),
		MkString("clientOrderId"),
	}))
	if IsTrue(OpNotEq2(clientOrderId, MkUndefined())) {
		*(request).At(MkString("userref")) = OpCopy(clientOrderId)
	}
	if IsTrue(OpEq2(type_, MkString("limit"))) {
		*(request).At(MkString("price")) = this.PriceToPrecision(symbol, price)
	} else {
		if IsTrue(OpOr(OpEq2(type_, MkString("stop-loss")), OpEq2(type_, MkString("take-profit")))) {
			stopPrice := this.SafeNumber2(params, MkString("price"), MkString("stopPrice"), price)
			_ = stopPrice
			if IsTrue(OpEq2(stopPrice, MkUndefined())) {
				panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" createOrder() requires a price argument or a price/stopPrice parameter for a "), OpAdd(type_, MkString(" order"))))))
			} else {
				*(request).At(MkString("price")) = this.PriceToPrecision(symbol, stopPrice)
			}
		} else {
			if IsTrue(OpOr(OpEq2(type_, MkString("stop-loss-limit")), OpEq2(type_, MkString("take-profit-limit")))) {
				stopPrice := this.SafeNumber2(params, MkString("price"), MkString("stopPrice"))
				_ = stopPrice
				limitPrice := this.SafeNumber(params, MkString("price2"))
				_ = limitPrice
				stopPriceDefined := OpNotEq2(stopPrice, MkUndefined())
				_ = stopPriceDefined
				limitPriceDefined := OpNotEq2(limitPrice, MkUndefined())
				_ = limitPriceDefined
				if IsTrue(OpAnd(stopPriceDefined, limitPriceDefined)) {
					*(request).At(MkString("price")) = this.PriceToPrecision(symbol, stopPrice)
					*(request).At(MkString("price2")) = this.PriceToPrecision(symbol, limitPrice)
				} else {
					if IsTrue(OpOr(OpEq2(price, MkUndefined()), OpNot(OpOr(stopPriceDefined, limitPriceDefined)))) {
						panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" createOrder() requires a price argument and/or price/stopPrice/price2 parameters for a "), OpAdd(type_, MkString(" order"))))))
					} else {
						if IsTrue(stopPriceDefined) {
							*(request).At(MkString("price")) = this.PriceToPrecision(symbol, stopPrice)
							*(request).At(MkString("price2")) = this.PriceToPrecision(symbol, price)
						} else {
							if IsTrue(limitPriceDefined) {
								*(request).At(MkString("price")) = this.PriceToPrecision(symbol, price)
								*(request).At(MkString("price2")) = this.PriceToPrecision(symbol, limitPrice)
							}
						}
					}
				}
			}
		}
	}
	params = this.Omit(params, MkArray(&VarArray{
		MkString("price"),
		MkString("stopPrice"),
		MkString("price2"),
	}))
	response := this.Call(MkString("privatePostAddOrder"), this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"))
	_ = result
	return this.ParseOrder(result)
}

func (this *Kraken) FindMarketByAltnameOrId(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	if IsTrue(OpHasMember(id, *this.At(MkString("marketsByAltname")))) {
		return *(*this.At(MkString("marketsByAltname"))).At(id)
	} else {
		if IsTrue(OpHasMember(id, *this.At(MkString("markets_by_id")))) {
			return *(*this.At(MkString("markets_by_id"))).At(id)
		}
	}
	return MkUndefined()
}

func (this *Kraken) GetDelistedMarketById(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	if IsTrue(OpEq2(id, MkUndefined())) {
		return id
	}
	market := this.SafeValue(*(*this.At(MkString("options"))).At(MkString("delistedMarketsById")), id)
	_ = market
	if IsTrue(OpNotEq2(market, MkUndefined())) {
		return market
	}
	baseIdStart := MkInteger(0)
	_ = baseIdStart
	baseIdEnd := MkInteger(3)
	_ = baseIdEnd
	quoteIdStart := MkInteger(3)
	_ = quoteIdStart
	quoteIdEnd := MkInteger(6)
	_ = quoteIdEnd
	if IsTrue(OpEq2(id.Length, MkInteger(8))) {
		baseIdEnd = MkInteger(4)
		quoteIdStart = MkInteger(4)
		quoteIdEnd = MkInteger(8)
	} else {
		if IsTrue(OpEq2(id.Length, MkInteger(7))) {
			baseIdEnd = MkInteger(4)
			quoteIdStart = MkInteger(4)
			quoteIdEnd = MkInteger(7)
		}
	}
	baseId := id.Slice(baseIdStart, baseIdEnd)
	_ = baseId
	quoteId := id.Slice(quoteIdStart, quoteIdEnd)
	_ = quoteId
	base := this.SafeCurrencyCode(baseId)
	_ = base
	quote := this.SafeCurrencyCode(quoteId)
	_ = quote
	symbol := OpAdd(base, OpAdd(MkString("/"), quote))
	_ = symbol
	market = MkMap(&VarMap{
		"symbol":  symbol,
		"base":    base,
		"quote":   quote,
		"baseId":  baseId,
		"quoteId": quoteId,
	})
	*(*(*this.At(MkString("options"))).At(MkString("delistedMarketsById"))).At(id) = OpCopy(market)
	return market
}

func (this *Kraken) ParseOrderStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"pending":  MkString("open"),
		"open":     MkString("open"),
		"closed":   MkString("closed"),
		"canceled": MkString("canceled"),
		"expired":  MkString("expired"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Kraken) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	description := this.SafeValue(order, MkString("descr"), MkMap(&VarMap{}))
	_ = description
	orderDescription := this.SafeString(description, MkString("order"))
	_ = orderDescription
	side := OpCopy(MkUndefined())
	_ = side
	type_ := OpCopy(MkUndefined())
	_ = type_
	marketId := OpCopy(MkUndefined())
	_ = marketId
	price := OpCopy(MkUndefined())
	_ = price
	amount := OpCopy(MkUndefined())
	_ = amount
	if IsTrue(OpNotEq2(orderDescription, MkUndefined())) {
		parts := orderDescription.Split(MkString(" "))
		_ = parts
		side = this.SafeString(parts, MkInteger(0))
		amount = this.SafeNumber(parts, MkInteger(1))
		marketId = this.SafeString(parts, MkInteger(2))
		type_ = this.SafeString(parts, MkInteger(4))
		price = this.SafeNumber(parts, MkInteger(5))
	}
	side = this.SafeString(description, MkString("type"), side)
	type_ = this.SafeString(description, MkString("ordertype"), type_)
	marketId = this.SafeString(description, MkString("pair"), marketId)
	foundMarket := this.FindMarketByAltnameOrId(marketId)
	_ = foundMarket
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(OpNotEq2(foundMarket, MkUndefined())) {
		market = OpCopy(foundMarket)
	} else {
		if IsTrue(OpNotEq2(marketId, MkUndefined())) {
			market = this.GetDelistedMarketById(marketId)
		}
	}
	timestamp := this.SafeTimestamp(order, MkString("opentm"))
	_ = timestamp
	amount = this.SafeNumber(order, MkString("vol"), amount)
	filled := this.SafeNumber(order, MkString("vol_exec"))
	_ = filled
	fee := OpCopy(MkUndefined())
	_ = fee
	cost := this.SafeNumber(order, MkString("cost"))
	_ = cost
	price = this.SafeNumber(description, MkString("price"), price)
	if IsTrue(OpOr(OpEq2(price, MkUndefined()), OpEq2(price, MkNumber(0.0)))) {
		price = this.SafeNumber(description, MkString("price2"))
	}
	if IsTrue(OpOr(OpEq2(price, MkUndefined()), OpEq2(price, MkNumber(0.0)))) {
		price = this.SafeNumber(order, MkString("price"), price)
	}
	average := this.SafeNumber(order, MkString("price"))
	_ = average
	if IsTrue(OpNotEq2(market, MkUndefined())) {
		symbol = *(market).At(MkString("symbol"))
		if IsTrue(OpHasMember(MkString("fee"), order)) {
			flags := *(order).At(MkString("oflags"))
			_ = flags
			feeCost := this.SafeNumber(order, MkString("fee"))
			_ = feeCost
			fee = MkMap(&VarMap{
				"cost": feeCost,
				"rate": MkUndefined(),
			})
			if IsTrue(OpGtEq(flags.IndexOf(MkString("fciq")), MkInteger(0))) {
				*(fee).At(MkString("currency")) = *(market).At(MkString("quote"))
			} else {
				if IsTrue(OpGtEq(flags.IndexOf(MkString("fcib")), MkInteger(0))) {
					*(fee).At(MkString("currency")) = *(market).At(MkString("base"))
				}
			}
		}
	}
	status := this.ParseOrderStatus(this.SafeString(order, MkString("status")))
	_ = status
	id := this.SafeString(order, MkString("id"))
	_ = id
	if IsTrue(OpEq2(id, MkUndefined())) {
		txid := this.SafeValue(order, MkString("txid"))
		_ = txid
		id = this.SafeString(txid, MkInteger(0))
	}
	clientOrderId := this.SafeString(order, MkString("userref"))
	_ = clientOrderId
	rawTrades := this.SafeValue(order, MkString("trades"))
	_ = rawTrades
	trades := OpCopy(MkUndefined())
	_ = trades
	if IsTrue(OpNotEq2(rawTrades, MkUndefined())) {
		trades = this.ParseTrades(rawTrades, market, MkUndefined(), MkUndefined(), MkMap(&VarMap{"order": id}))
	}
	stopPrice := this.SafeNumber(order, MkString("stopprice"))
	_ = stopPrice
	return this.SafeOrder(MkMap(&VarMap{
		"id":                 id,
		"clientOrderId":      clientOrderId,
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

func (this *Kraken) FetchOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	clientOrderId := this.SafeValue2(params, MkString("userref"), MkString("clientOrderId"))
	_ = clientOrderId
	request := MkMap(&VarMap{"trades": MkBool(true)})
	_ = request
	query := OpCopy(params)
	_ = query
	if IsTrue(OpNotEq2(clientOrderId, MkUndefined())) {
		*(request).At(MkString("userref")) = OpCopy(clientOrderId)
		query = this.Omit(params, MkArray(&VarArray{
			MkString("userref"),
			MkString("clientOrderId"),
		}))
	} else {
		*(request).At(MkString("txid")) = OpCopy(id)
	}
	response := this.Call(MkString("privatePostQueryOrders"), this.Extend(request, query))
	_ = response
	result := this.SafeValue(response, MkString("result"), MkArray(&VarArray{}))
	_ = result
	if IsTrue(OpNot(OpHasMember(id, result))) {
		panic(NewOrderNotFound(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" fetchOrder() could not find order id "), id))))
	}
	order := this.ParseOrder(this.Extend(MkMap(&VarMap{"id": id}), *(result).At(id)))
	_ = order
	return this.Extend(MkMap(&VarMap{"info": response}), order)
}

func (this *Kraken) FetchOrderTrades(goArgs ...*Variant) *Variant {
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
	orderTrades := this.SafeValue(params, MkString("trades"))
	_ = orderTrades
	tradeIds := MkArray(&VarArray{})
	_ = tradeIds
	if IsTrue(OpEq2(orderTrades, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchOrderTrades() requires a unified order structure in the params argument or a 'trades' param (an array of trade id strings)"))))
	} else {
		for i := MkInteger(0); IsTrue(OpLw(i, orderTrades.Length)); OpIncr(&i) {
			orderTrade := *(orderTrades).At(i)
			_ = orderTrade
			if IsTrue(OpEq2(OpType(orderTrade), MkString("string"))) {
				tradeIds.Push(orderTrade)
			} else {
				tradeIds.Push(*(orderTrade).At(MkString("id")))
			}
		}
	}
	this.LoadMarkets()
	options := this.SafeValue(*this.At(MkString("options")), MkString("fetchOrderTrades"), MkMap(&VarMap{}))
	_ = options
	batchSize := this.SafeInteger(options, MkString("batchSize"), MkInteger(20))
	_ = batchSize
	numTradeIds := OpCopy(tradeIds.Length)
	_ = numTradeIds
	numBatches := parseInt(OpDiv(numTradeIds, batchSize))
	_ = numBatches
	numBatches = this.Sum(numBatches, MkInteger(1))
	result := MkArray(&VarArray{})
	_ = result
	for j := MkInteger(0); IsTrue(OpLw(j, numBatches)); OpIncr(&j) {
		requestIds := MkArray(&VarArray{})
		_ = requestIds
		for k := MkInteger(0); IsTrue(OpLw(k, batchSize)); OpIncr(&k) {
			index := this.Sum(OpMulti(j, batchSize), k)
			_ = index
			if IsTrue(OpLw(index, numTradeIds)) {
				requestIds.Push(*(tradeIds).At(index))
			}
		}
		request := MkMap(&VarMap{"txid": requestIds.Join(MkString(","))})
		_ = request
		response := this.Call(MkString("privatePostQueryTrades"), request)
		_ = response
		rawTrades := this.SafeValue(response, MkString("result"))
		_ = rawTrades
		ids := GoGetKeys(rawTrades)
		_ = ids
		for i := MkInteger(0); IsTrue(OpLw(i, ids.Length)); OpIncr(&i) {
			*(*(rawTrades).At(*(ids).At(i))).At(MkString("id")) = *(ids).At(i)
		}
		trades := this.ParseTrades(rawTrades, MkUndefined(), since, limit)
		_ = trades
		tradesFilteredBySymbol := this.FilterBySymbol(trades, symbol)
		_ = tradesFilteredBySymbol
		result = this.ArrayConcat(result, tradesFilteredBySymbol)
	}
	return result
}

func (this *Kraken) FetchOrdersByIds(goArgs ...*Variant) *Variant {
	ids := GoGetArg(goArgs, 0, MkUndefined())
	_ = ids
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privatePostQueryOrders"), this.Extend(MkMap(&VarMap{
		"trades": MkBool(true),
		"txid":   ids.Join(MkString(",")),
	}), params))
	_ = response
	result := this.SafeValue(response, MkString("result"), MkMap(&VarMap{}))
	_ = result
	orders := MkArray(&VarArray{})
	_ = orders
	orderIds := GoGetKeys(result)
	_ = orderIds
	for i := MkInteger(0); IsTrue(OpLw(i, orderIds.Length)); OpIncr(&i) {
		id := *(orderIds).At(i)
		_ = id
		item := *(result).At(id)
		_ = item
		order := this.ParseOrder(this.Extend(MkMap(&VarMap{"id": id}), item))
		_ = order
		orders.Push(order)
	}
	return orders
}

func (this *Kraken) FetchMyTrades(goArgs ...*Variant) *Variant {
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
		*(request).At(MkString("start")) = parseInt(OpDiv(since, MkInteger(1000)))
	}
	response := this.Call(MkString("privatePostTradesHistory"), this.Extend(request, params))
	_ = response
	trades := *(*(response).At(MkString("result"))).At(MkString("trades"))
	_ = trades
	ids := GoGetKeys(trades)
	_ = ids
	for i := MkInteger(0); IsTrue(OpLw(i, ids.Length)); OpIncr(&i) {
		*(*(trades).At(*(ids).At(i))).At(MkString("id")) = *(ids).At(i)
	}
	market := OpCopy(MkUndefined())
	_ = market
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
	}
	return this.ParseTrades(trades, market, since, limit)
}

func (this *Kraken) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := OpCopy(MkUndefined())
	_ = response
	{
		ret__ := func(this *Kraken) (ret_ *Variant) {
			defer func() {
				if e := recover().(*Variant); e != nil {
					ret_ = func(this *Kraken) *Variant {
						// catch block:
						if IsTrue(*this.At(MkString("last_http_response"))) {
							if IsTrue(OpGtEq((*((this).At(MkString("last_http_response")))).Call(MkString("indexOf"), MkString("EOrder:Unknown order")), MkInteger(0))) {
								panic(NewOrderNotFound(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" cancelOrder() error "), *this.At(MkString("last_http_response"))))))
							}
						}
						panic(e)
						return nil
					}(this)
				}
			}()
			// try block:
			response = this.Call(MkString("privatePostCancelOrder"), this.Extend(MkMap(&VarMap{"txid": id}), params))
			return nil
		}(this)
		if ret__ != nil {
			return ret__
		}
	}
	return response
}

func (this *Kraken) CancelAllOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	return this.Call(MkString("privatePostCancelAll"), params)
}

func (this *Kraken) FetchOpenOrders(goArgs ...*Variant) *Variant {
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
		*(request).At(MkString("start")) = parseInt(OpDiv(since, MkInteger(1000)))
	}
	response := this.Call(MkString("privatePostOpenOrders"), this.Extend(request, params))
	_ = response
	market := OpCopy(MkUndefined())
	_ = market
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
	}
	result := this.SafeValue(response, MkString("result"), MkMap(&VarMap{}))
	_ = result
	orders := this.SafeValue(result, MkString("open"), MkArray(&VarArray{}))
	_ = orders
	return this.ParseOrders(orders, market, since, limit)
}

func (this *Kraken) FetchClosedOrders(goArgs ...*Variant) *Variant {
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
		*(request).At(MkString("start")) = parseInt(OpDiv(since, MkInteger(1000)))
	}
	response := this.Call(MkString("privatePostClosedOrders"), this.Extend(request, params))
	_ = response
	market := OpCopy(MkUndefined())
	_ = market
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
	}
	result := this.SafeValue(response, MkString("result"), MkMap(&VarMap{}))
	_ = result
	orders := this.SafeValue(result, MkString("closed"), MkArray(&VarArray{}))
	_ = orders
	return this.ParseOrders(orders, market, since, limit)
}

func (this *Kraken) FetchDepositMethods(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{"asset": *(currency).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("privatePostDepositMethods"), this.Extend(request, params))
	_ = response
	return this.SafeValue(response, MkString("result"))
}

func (this *Kraken) ParseTransactionStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"Initial": MkString("pending"),
		"Pending": MkString("pending"),
		"Success": MkString("ok"),
		"Settled": MkString("pending"),
		"Failure": MkString("failed"),
		"Partial": MkString("ok"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Kraken) ParseTransaction(goArgs ...*Variant) *Variant {
	transaction := GoGetArg(goArgs, 0, MkUndefined())
	_ = transaction
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	id := this.SafeString(transaction, MkString("refid"))
	_ = id
	txid := this.SafeString(transaction, MkString("txid"))
	_ = txid
	timestamp := this.SafeTimestamp(transaction, MkString("time"))
	_ = timestamp
	currencyId := this.SafeString(transaction, MkString("asset"))
	_ = currencyId
	code := this.SafeCurrencyCode(currencyId, currency)
	_ = code
	address := this.SafeString(transaction, MkString("info"))
	_ = address
	amount := this.SafeNumber(transaction, MkString("amount"))
	_ = amount
	status := this.ParseTransactionStatus(this.SafeString(transaction, MkString("status")))
	_ = status
	type_ := this.SafeString(transaction, MkString("type"))
	_ = type_
	feeCost := this.SafeNumber(transaction, MkString("fee"))
	_ = feeCost
	if IsTrue(OpEq2(feeCost, MkUndefined())) {
		if IsTrue(OpEq2(type_, MkString("deposit"))) {
			feeCost = MkInteger(0)
		}
	}
	return MkMap(&VarMap{
		"info":      transaction,
		"id":        id,
		"currency":  code,
		"amount":    amount,
		"address":   address,
		"tag":       MkUndefined(),
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

func (this *Kraken) ParseTransactionsByType(goArgs ...*Variant) *Variant {
	type_ := GoGetArg(goArgs, 0, MkUndefined())
	_ = type_
	transactions := GoGetArg(goArgs, 1, MkUndefined())
	_ = transactions
	code := GoGetArg(goArgs, 2, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 3, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 4, MkUndefined())
	_ = limit
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, transactions.Length)); OpIncr(&i) {
		transaction := this.ParseTransaction(this.Extend(MkMap(&VarMap{"type": type_}), *(transactions).At(i)))
		_ = transaction
		result.Push(transaction)
	}
	return this.FilterByCurrencySinceLimit(result, code, since, limit)
}

func (this *Kraken) FetchDeposits(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"asset": *(currency).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("privatePostDepositStatus"), this.Extend(request, params))
	_ = response
	return this.ParseTransactionsByType(MkString("deposit"), *(response).At(MkString("result")), code, since, limit)
}

func (this *Kraken) FetchTime(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetTime"), params)
	_ = response
	result := this.SafeValue(response, MkString("result"), MkMap(&VarMap{}))
	_ = result
	return this.SafeTimestamp(result, MkString("unixtime"))
}

func (this *Kraken) FetchWithdrawals(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"asset": *(currency).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("privatePostWithdrawStatus"), this.Extend(request, params))
	_ = response
	return this.ParseTransactionsByType(MkString("withdrawal"), *(response).At(MkString("result")), code, since, limit)
}

func (this *Kraken) CreateDepositAddress(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	request := MkMap(&VarMap{"new": MkString("true")})
	_ = request
	response := this.FetchDepositAddress(code, this.Extend(request, params))
	_ = response
	address := this.SafeString(response, MkString("address"))
	_ = address
	this.CheckAddress(address)
	return MkMap(&VarMap{
		"currency": code,
		"address":  address,
		"info":     response,
	})
}

func (this *Kraken) FetchDepositAddress(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	method := this.SafeString(params, MkString("method"))
	_ = method
	if IsTrue(OpEq2(method, MkUndefined())) {
		if IsTrue(*(*this.At(MkString("options"))).At(MkString("cacheDepositMethodsOnFetchDepositAddress"))) {
			if IsTrue(OpNot(OpHasMember(code, *(*this.At(MkString("options"))).At(MkString("depositMethods"))))) {
				*(*(*this.At(MkString("options"))).At(MkString("depositMethods"))).At(code) = this.FetchDepositMethods(code)
			}
			method = *(*(*(*(*this.At(MkString("options"))).At(MkString("depositMethods"))).At(code)).At(MkInteger(0))).At(MkString("method"))
		} else {
			panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" fetchDepositAddress() requires an extra `method` parameter. Use fetchDepositMethods (\""), OpAdd(code, MkString("\") to get a list of available deposit methods or enable the exchange property .options[\"cacheDepositMethodsOnFetchDepositAddress\"] = true"))))))
		}
	}
	request := MkMap(&VarMap{
		"asset":  *(currency).At(MkString("id")),
		"method": method,
	})
	_ = request
	response := this.Call(MkString("privatePostDepositAddresses"), this.Extend(request, params))
	_ = response
	result := *(response).At(MkString("result"))
	_ = result
	numResults := OpCopy(result.Length)
	_ = numResults
	if IsTrue(OpLw(numResults, MkInteger(1))) {
		panic(NewInvalidAddress(OpAdd(*this.At(MkString("id")), MkString(" privatePostDepositAddresses() returned no addresses"))))
	}
	address := this.SafeString(*(result).At(MkInteger(0)), MkString("address"))
	_ = address
	tag := this.SafeString2(*(result).At(MkInteger(0)), MkString("tag"), MkString("memo"))
	_ = tag
	this.CheckAddress(address)
	return MkMap(&VarMap{
		"currency": code,
		"address":  address,
		"tag":      tag,
		"info":     response,
	})
}

func (this *Kraken) Withdraw(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpHasMember(MkString("key"), params)) {
		this.LoadMarkets()
		currency := this.Currency(code)
		_ = currency
		request := MkMap(&VarMap{
			"asset":  *(currency).At(MkString("id")),
			"amount": amount,
		})
		_ = request
		response := this.Call(MkString("privatePostWithdraw"), this.Extend(request, params))
		_ = response
		result := this.SafeValue(response, MkString("result"), MkMap(&VarMap{}))
		_ = result
		id := this.SafeString(result, MkString("refid"))
		_ = id
		return MkMap(&VarMap{
			"info": result,
			"id":   id,
		})
	}
	panic(NewExchangeError(OpAdd(*this.At(MkString("id")), MkString(" withdraw() requires a 'key' parameter (withdrawal key name, as set up on your account)"))))
	return MkUndefined()
}

func (this *Kraken) FetchPositions(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{})
	_ = request
	response := this.Call(MkString("privatePostOpenPositions"), this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"))
	_ = result
	return result
}

func (this *Kraken) Sign(goArgs ...*Variant) *Variant {
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
	url := OpAdd(MkString("/"), OpAdd(*this.At(MkString("version")), OpAdd(MkString("/"), OpAdd(api, OpAdd(MkString("/"), path)))))
	_ = url
	if IsTrue(OpEq2(api, MkString("public"))) {
		if IsTrue(*(GoGetKeys(params)).At(MkString("length"))) {
			OpAddAssign(&url, OpAdd(MkString("?"), this.Urlencode(params)))
		}
	} else {
		if IsTrue(OpEq2(api, MkString("private"))) {
			this.CheckRequiredCredentials()
			nonce := (this.Nonce()).Call(MkString("toString"))
			_ = nonce
			body = this.Urlencode(this.Extend(MkMap(&VarMap{"nonce": nonce}), params))
			auth := this.Encode(OpAdd(nonce, body))
			_ = auth
			hash := this.Hash(auth, MkString("sha256"), MkString("binary"))
			_ = hash
			binary := this.StringToBinary(this.Encode(url))
			_ = binary
			binhash := this.BinaryConcat(binary, hash)
			_ = binhash
			secret := this.Base64ToBinary(*this.At(MkString("secret")))
			_ = secret
			signature := this.Hmac(binhash, secret, MkString("sha512"), MkString("base64"))
			_ = signature
			headers = MkMap(&VarMap{
				"API-Key":      *this.At(MkString("apiKey")),
				"API-Sign":     signature,
				"Content-Type": MkString("application/x-www-form-urlencoded"),
			})
		} else {
			url = OpAdd(MkString("/"), path)
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

func (this *Kraken) Nonce(goArgs ...*Variant) *Variant {
	return this.Milliseconds()
}

func (this *Kraken) HandleErrors(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpEq2(code, MkInteger(520))) {
		panic(NewExchangeNotAvailable(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), OpAdd(code.ToString(), OpAdd(MkString(" "), reason))))))
	}
	if IsTrue(OpGtEq(body.IndexOf(MkString("Invalid order")), MkInteger(0))) {
		panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))))
	}
	if IsTrue(OpGtEq(body.IndexOf(MkString("Invalid nonce")), MkInteger(0))) {
		panic(NewInvalidNonce(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))))
	}
	if IsTrue(OpGtEq(body.IndexOf(MkString("Insufficient funds")), MkInteger(0))) {
		panic(NewInsufficientFunds(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))))
	}
	if IsTrue(OpGtEq(body.IndexOf(MkString("Cancel pending")), MkInteger(0))) {
		panic(NewCancelPending(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))))
	}
	if IsTrue(OpGtEq(body.IndexOf(MkString("Invalid arguments:volume")), MkInteger(0))) {
		panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))))
	}
	if IsTrue(OpGtEq(body.IndexOf(MkString("Rate limit exceeded")), MkInteger(0))) {
		panic(NewRateLimitExceeded(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))))
	}
	if IsTrue(OpEq2(response, MkUndefined())) {
		return MkUndefined()
	}
	if IsTrue(OpEq2(*(body).At(MkInteger(0)), MkString("{"))) {
		if IsTrue(OpNotEq2(OpType(response), MkString("string"))) {
			if IsTrue(OpHasMember(MkString("error"), response)) {
				numErrors := *(*(response).At(MkString("error"))).At(MkString("length"))
				_ = numErrors
				if IsTrue(numErrors) {
					message := OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))
					_ = message
					for i := MkInteger(0); IsTrue(OpLw(i, *(*(response).At(MkString("error"))).At(MkString("length")))); OpIncr(&i) {
						error := *(*(response).At(MkString("error"))).At(i)
						_ = error
						this.ThrowExactlyMatchedException(*this.At(MkString("exceptions")), error, message)
					}
					panic(NewExchangeError(message))
				}
			}
		}
	}
	return MkUndefined()
}
