package ccxt

import ()

type Bitfinex struct {
	*ExchangeBase
}

var _ Exchange = (*Bitfinex)(nil)

func init() {
	exchange := &Bitfinex{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Bitfinex) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("bitfinex"),
		"name": MkString("Bitfinex"),
		"countries": MkArray(&VarArray{
			MkString("VG"),
		}),
		"version":   MkString("v1"),
		"rateLimit": MkInteger(1500),
		"pro":       MkBool(true),
		"has": MkMap(&VarMap{
			"cancelAllOrders":      MkBool(true),
			"cancelOrder":          MkBool(true),
			"CORS":                 MkBool(false),
			"createDepositAddress": MkBool(true),
			"createOrder":          MkBool(true),
			"deposit":              MkBool(true),
			"editOrder":            MkBool(true),
			"fetchBalance":         MkBool(true),
			"fetchClosedOrders":    MkBool(true),
			"fetchDepositAddress":  MkBool(true),
			"fetchDeposits":        MkBool(false),
			"fetchFundingFees":     MkBool(true),
			"fetchMarkets":         MkBool(true),
			"fetchMyTrades":        MkBool(true),
			"fetchOHLCV":           MkBool(true),
			"fetchOpenOrders":      MkBool(true),
			"fetchOrder":           MkBool(true),
			"fetchOrderBook":       MkBool(true),
			"fetchTicker":          MkBool(true),
			"fetchTickers":         MkBool(true),
			"fetchTrades":          MkBool(true),
			"fetchTradingFee":      MkBool(true),
			"fetchTradingFees":     MkBool(true),
			"fetchTransactions":    MkBool(true),
			"fetchWithdrawals":     MkBool(false),
			"withdraw":             MkBool(true),
			"transfer":             MkBool(true),
		}),
		"timeframes": MkMap(&VarMap{
			"1m":  MkString("1m"),
			"5m":  MkString("5m"),
			"15m": MkString("15m"),
			"30m": MkString("30m"),
			"1h":  MkString("1h"),
			"3h":  MkString("3h"),
			"4h":  MkString("4h"),
			"6h":  MkString("6h"),
			"12h": MkString("12h"),
			"1d":  MkString("1D"),
			"1w":  MkString("7D"),
			"2w":  MkString("14D"),
			"1M":  MkString("1M"),
		}),
		"urls": MkMap(&VarMap{
			"logo": MkString("https://user-images.githubusercontent.com/1294454/27766244-e328a50c-5ed2-11e7-947b-041416579bb3.jpg"),
			"api": MkMap(&VarMap{
				"v2":      MkString("https://api-pub.bitfinex.com"),
				"public":  MkString("https://api.bitfinex.com"),
				"private": MkString("https://api.bitfinex.com"),
			}),
			"www":      MkString("https://www.bitfinex.com"),
			"referral": MkString("https://www.bitfinex.com/?refcode=P61eYxFL"),
			"doc": MkArray(&VarArray{
				MkString("https://docs.bitfinex.com/v1/docs"),
				MkString("https://github.com/bitfinexcom/bitfinex-api-node"),
			}),
		}),
		"api": MkMap(&VarMap{
			"v2": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("platform/status"),
				MkString("tickers"),
				MkString("ticker/{symbol}"),
				MkString("trades/{symbol}/hist"),
				MkString("book/{symbol}/{precision}"),
				MkString("book/{symbol}/P0"),
				MkString("book/{symbol}/P1"),
				MkString("book/{symbol}/P2"),
				MkString("book/{symbol}/P3"),
				MkString("book/{symbol}/R0"),
				MkString("stats1/{key}:{size}:{symbol}:{side}/{section}"),
				MkString("stats1/{key}:{size}:{symbol}/{section}"),
				MkString("stats1/{key}:{size}:{symbol}:long/last"),
				MkString("stats1/{key}:{size}:{symbol}:long/hist"),
				MkString("stats1/{key}:{size}:{symbol}:short/last"),
				MkString("stats1/{key}:{size}:{symbol}:short/hist"),
				MkString("candles/trade:{timeframe}:{symbol}/{section}"),
				MkString("candles/trade:{timeframe}:{symbol}/last"),
				MkString("candles/trade:{timeframe}:{symbol}/hist"),
			})}),
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("book/{symbol}"),
				MkString("lendbook/{currency}"),
				MkString("lends/{currency}"),
				MkString("pubticker/{symbol}"),
				MkString("stats/{symbol}"),
				MkString("symbols"),
				MkString("symbols_details"),
				MkString("tickers"),
				MkString("trades/{symbol}"),
			})}),
			"private": MkMap(&VarMap{"post": MkArray(&VarArray{
				MkString("account_fees"),
				MkString("account_infos"),
				MkString("balances"),
				MkString("basket_manage"),
				MkString("credits"),
				MkString("deposit/new"),
				MkString("funding/close"),
				MkString("history"),
				MkString("history/movements"),
				MkString("key_info"),
				MkString("margin_infos"),
				MkString("mytrades"),
				MkString("mytrades_funding"),
				MkString("offer/cancel"),
				MkString("offer/new"),
				MkString("offer/status"),
				MkString("offers"),
				MkString("offers/hist"),
				MkString("order/cancel"),
				MkString("order/cancel/all"),
				MkString("order/cancel/multi"),
				MkString("order/cancel/replace"),
				MkString("order/new"),
				MkString("order/new/multi"),
				MkString("order/status"),
				MkString("orders"),
				MkString("orders/hist"),
				MkString("position/claim"),
				MkString("position/close"),
				MkString("positions"),
				MkString("summary"),
				MkString("taken_funds"),
				MkString("total_taken_funds"),
				MkString("transfer"),
				MkString("unused_taken_funds"),
				MkString("withdraw"),
			})}),
		}),
		"fees": MkMap(&VarMap{
			"trading": MkMap(&VarMap{
				"feeSide":    MkString("get"),
				"tierBased":  MkBool(true),
				"percentage": MkBool(true),
				"maker":      this.ParseNumber(MkString("0.001")),
				"taker":      this.ParseNumber(MkString("0.002")),
				"tiers": MkMap(&VarMap{
					"taker": MkArray(&VarArray{
						MkArray(&VarArray{
							this.ParseNumber(MkString("0")),
							this.ParseNumber(MkString("0.002")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("500000")),
							this.ParseNumber(MkString("0.002")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("1000000")),
							this.ParseNumber(MkString("0.002")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("2500000")),
							this.ParseNumber(MkString("0.002")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("5000000")),
							this.ParseNumber(MkString("0.002")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("7500000")),
							this.ParseNumber(MkString("0.002")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("10000000")),
							this.ParseNumber(MkString("0.0018")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("15000000")),
							this.ParseNumber(MkString("0.0016")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("20000000")),
							this.ParseNumber(MkString("0.0014")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("25000000")),
							this.ParseNumber(MkString("0.0012")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("30000000")),
							this.ParseNumber(MkString("0.001")),
						}),
					}),
					"maker": MkArray(&VarArray{
						MkArray(&VarArray{
							this.ParseNumber(MkString("0")),
							this.ParseNumber(MkString("0.001")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("500000")),
							this.ParseNumber(MkString("0.0008")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("1000000")),
							this.ParseNumber(MkString("0.0006")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("2500000")),
							this.ParseNumber(MkString("0.0004")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("5000000")),
							this.ParseNumber(MkString("0.0002")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("7500000")),
							this.ParseNumber(MkString("0")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("10000000")),
							this.ParseNumber(MkString("0")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("15000000")),
							this.ParseNumber(MkString("0")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("20000000")),
							this.ParseNumber(MkString("0")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("25000000")),
							this.ParseNumber(MkString("0")),
						}),
						MkArray(&VarArray{
							this.ParseNumber(MkString("30000000")),
							this.ParseNumber(MkString("0")),
						}),
					}),
				}),
			}),
			"funding": MkMap(&VarMap{
				"tierBased":  MkBool(false),
				"percentage": MkBool(false),
				"deposit":    MkMap(&VarMap{}),
				"withdraw":   MkMap(&VarMap{}),
			}),
		}),
		"commonCurrencies": MkMap(&VarMap{
			"ALG":      MkString("ALGO"),
			"AMP":      MkString("AMPL"),
			"ATO":      MkString("ATOM"),
			"BCHABC":   MkString("BCHA"),
			"BCHN":     MkString("BCH"),
			"DAT":      MkString("DATA"),
			"DOG":      MkString("MDOGE"),
			"DSH":      MkString("DASH"),
			"EDO":      MkString("PNT"),
			"EUS":      MkString("EURS"),
			"EUT":      MkString("EURT"),
			"IOT":      MkString("IOTA"),
			"IQX":      MkString("IQ"),
			"MNA":      MkString("MANA"),
			"ORS":      MkString("ORS Group"),
			"PAS":      MkString("PASS"),
			"QSH":      MkString("QASH"),
			"QTM":      MkString("QTUM"),
			"RBT":      MkString("RBTC"),
			"SNG":      MkString("SNGLS"),
			"STJ":      MkString("STORJ"),
			"TERRAUST": MkString("UST"),
			"TSD":      MkString("TUSD"),
			"YYW":      MkString("YOYOW"),
			"UDC":      MkString("USDC"),
			"UST":      MkString("USDT"),
			"VSY":      MkString("VSYS"),
			"WAX":      MkString("WAXP"),
			"XCH":      MkString("XCHF"),
			"ZBT":      MkString("ZB"),
		}),
		"exceptions": MkMap(&VarMap{
			"exact": MkMap(&VarMap{
				"temporarily_unavailable":                                  ExchangeNotAvailable,
				"Order could not be cancelled.":                            OrderNotFound,
				"No such order found.":                                     OrderNotFound,
				"Order price must be positive.":                            InvalidOrder,
				"Could not find a key matching the given X-BFX-APIKEY.":    AuthenticationError,
				"Key price should be a decimal number, e.g. \"123.456\"":   InvalidOrder,
				"Key amount should be a decimal number, e.g. \"123.456\"":  InvalidOrder,
				"ERR_RATE_LIMIT":                                           RateLimitExceeded,
				"Ratelimit":                                                RateLimitExceeded,
				"Nonce is too small.":                                      InvalidNonce,
				"No summary found.":                                        ExchangeError,
				"Cannot evaluate your available balance, please try again": ExchangeNotAvailable,
				"Unknown symbol":                                           BadSymbol,
				"Cannot complete transfer. Exchange balance insufficient.": InsufficientFunds,
				"Momentary balance check. Please wait few seconds and try the transfer again.": ExchangeError,
			}),
			"broad": MkMap(&VarMap{
				"Invalid X-BFX-SIGNATURE":               AuthenticationError,
				"This API key does not have permission": PermissionDenied,
				"not enough exchange balance for ":      InsufficientFunds,
				"minimum size for ":                     InvalidOrder,
				"Invalid order":                         InvalidOrder,
				"The available balance is only":         InsufficientFunds,
			}),
		}),
		"precisionMode": SIGNIFICANT_DIGITS,
		"options": MkMap(&VarMap{
			"currencyNames": MkMap(&VarMap{
				"AGI":   MkString("agi"),
				"AID":   MkString("aid"),
				"AIO":   MkString("aio"),
				"ANT":   MkString("ant"),
				"AVT":   MkString("aventus"),
				"BAT":   MkString("bat"),
				"BCH":   MkString("bab"),
				"BCI":   MkString("bci"),
				"BFT":   MkString("bft"),
				"BSV":   MkString("bsv"),
				"BTC":   MkString("bitcoin"),
				"BTG":   MkString("bgold"),
				"CFI":   MkString("cfi"),
				"COMP":  MkString("comp"),
				"DAI":   MkString("dai"),
				"DADI":  MkString("dad"),
				"DASH":  MkString("dash"),
				"DATA":  MkString("datacoin"),
				"DTH":   MkString("dth"),
				"EDO":   MkString("eidoo"),
				"ELF":   MkString("elf"),
				"EOS":   MkString("eos"),
				"ETC":   MkString("ethereumc"),
				"ETH":   MkString("ethereum"),
				"ETP":   MkString("metaverse"),
				"FUN":   MkString("fun"),
				"GNT":   MkString("golem"),
				"IOST":  MkString("ios"),
				"IOTA":  MkString("iota"),
				"LEO":   MkString("let"),
				"LINK":  MkString("link"),
				"LRC":   MkString("lrc"),
				"LTC":   MkString("litecoin"),
				"LYM":   MkString("lym"),
				"MANA":  MkString("mna"),
				"MIT":   MkString("mit"),
				"MKR":   MkString("mkr"),
				"MTN":   MkString("mtn"),
				"NEO":   MkString("neo"),
				"ODE":   MkString("ode"),
				"OMG":   MkString("omisego"),
				"OMNI":  MkString("mastercoin"),
				"QASH":  MkString("qash"),
				"QTUM":  MkString("qtum"),
				"RCN":   MkString("rcn"),
				"RDN":   MkString("rdn"),
				"REP":   MkString("rep"),
				"REQ":   MkString("req"),
				"RLC":   MkString("rlc"),
				"SAN":   MkString("santiment"),
				"SNGLS": MkString("sng"),
				"SNT":   MkString("status"),
				"SPANK": MkString("spk"),
				"STORJ": MkString("stj"),
				"TNB":   MkString("tnb"),
				"TRX":   MkString("trx"),
				"TUSD":  MkString("tsd"),
				"USD":   MkString("wire"),
				"USDC":  MkString("udc"),
				"UTK":   MkString("utk"),
				"USDT":  MkString("tetheruso"),
				"VEE":   MkString("vee"),
				"WAX":   MkString("wax"),
				"XLM":   MkString("xlm"),
				"XMR":   MkString("monero"),
				"XRP":   MkString("ripple"),
				"XVG":   MkString("xvg"),
				"YOYOW": MkString("yoyow"),
				"ZEC":   MkString("zcash"),
				"ZRX":   MkString("zrx"),
				"XTZ":   MkString("xtz"),
			}),
			"orderTypes": MkMap(&VarMap{
				"limit":  MkString("exchange limit"),
				"market": MkString("exchange market"),
			}),
			"accountsByType": MkMap(&VarMap{
				"spot":        MkString("exchange"),
				"margin":      MkString("trading"),
				"funding":     MkString("deposit"),
				"exchange":    MkString("exchange"),
				"trading":     MkString("trading"),
				"deposit":     MkString("deposit"),
				"derivatives": MkString("trading"),
			}),
		}),
	}))
}

func (this *Bitfinex) FetchFundingFees(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privatePostAccountFees"), params)
	_ = response
	fees := *(response).At(MkString("withdraw"))
	_ = fees
	withdraw := MkMap(&VarMap{})
	_ = withdraw
	ids := GoGetKeys(fees)
	_ = ids
	for i := MkInteger(0); IsTrue(OpLw(i, ids.Length)); OpIncr(&i) {
		id := *(ids).At(i)
		_ = id
		code := this.SafeCurrencyCode(id)
		_ = code
		*(withdraw).At(code) = this.SafeNumber(fees, id)
	}
	return MkMap(&VarMap{
		"info":     response,
		"withdraw": withdraw,
		"deposit":  withdraw,
	})
}

func (this *Bitfinex) FetchTradingFees(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privatePostSummary"), params)
	_ = response
	return MkMap(&VarMap{
		"info":  response,
		"maker": this.SafeNumber(response, MkString("maker_fee")),
		"taker": this.SafeNumber(response, MkString("taker_fee")),
	})
}

func (this *Bitfinex) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	ids := this.Call(MkString("publicGetSymbols"))
	_ = ids
	details := this.Call(MkString("publicGetSymbolsDetails"))
	_ = details
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, details.Length)); OpIncr(&i) {
		market := *(details).At(i)
		_ = market
		id := this.SafeString(market, MkString("pair"))
		_ = id
		if IsTrue(OpNot(this.InArray(id, ids))) {
			continue
		}
		id = id.ToUpperCase()
		baseId := OpCopy(MkUndefined())
		_ = baseId
		quoteId := OpCopy(MkUndefined())
		_ = quoteId
		if IsTrue(OpGtEq(id.IndexOf(MkString(":")), MkInteger(0))) {
			parts := id.Split(MkString(":"))
			_ = parts
			baseId = *(parts).At(MkInteger(0))
			quoteId = *(parts).At(MkInteger(1))
		} else {
			baseId = id.Slice(MkInteger(0), MkInteger(3))
			quoteId = id.Slice(MkInteger(3), MkInteger(6))
		}
		base := this.SafeCurrencyCode(baseId)
		_ = base
		quote := this.SafeCurrencyCode(quoteId)
		_ = quote
		symbol := OpAdd(base, OpAdd(MkString("/"), quote))
		_ = symbol
		precision := MkMap(&VarMap{
			"price":  this.SafeInteger(market, MkString("price_precision")),
			"amount": MkInteger(8),
		})
		_ = precision
		minAmountString := this.SafeString(market, MkString("minimum_order_size"))
		_ = minAmountString
		maxAmountString := this.SafeString(market, MkString("maximum_order_size"))
		_ = maxAmountString
		limits := MkMap(&VarMap{
			"amount": MkMap(&VarMap{
				"min": this.ParseNumber(minAmountString),
				"max": this.ParseNumber(maxAmountString),
			}),
			"price": MkMap(&VarMap{
				"min": this.ParseNumber(MkString("1e-8")),
				"max": MkUndefined(),
			}),
		})
		_ = limits
		*(limits).At(MkString("cost")) = MkMap(&VarMap{
			"min": MkUndefined(),
			"max": MkUndefined(),
		})
		margin := this.SafeValue(market, MkString("margin"))
		_ = margin
		result.Push(MkMap(&VarMap{
			"id":        id,
			"symbol":    symbol,
			"base":      base,
			"quote":     quote,
			"baseId":    baseId,
			"quoteId":   quoteId,
			"active":    MkBool(true),
			"type":      MkString("spot"),
			"margin":    margin,
			"precision": precision,
			"limits":    limits,
			"info":      market,
		}))
	}
	return result
}

func (this *Bitfinex) AmountToPrecision(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	amount := GoGetArg(goArgs, 1, MkUndefined())
	_ = amount
	return this.DecimalToPrecision(amount, TRUNCATE, *(*(*(*this.At(MkString("markets"))).At(symbol)).At(MkString("precision"))).At(MkString("amount")), DECIMAL_PLACES)
}

func (this *Bitfinex) PriceToPrecision(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	price := GoGetArg(goArgs, 1, MkUndefined())
	_ = price
	price = this.DecimalToPrecision(price, ROUND, *(*(*(*this.At(MkString("markets"))).At(symbol)).At(MkString("precision"))).At(MkString("price")), *this.At(MkString("precisionMode")))
	return this.DecimalToPrecision(price, TRUNCATE, MkInteger(8), DECIMAL_PLACES)
}

func (this *Bitfinex) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	accountsByType := this.SafeValue(*this.At(MkString("options")), MkString("accountsByType"), MkMap(&VarMap{}))
	_ = accountsByType
	requestedType := this.SafeString(params, MkString("type"), MkString("exchange"))
	_ = requestedType
	accountType := this.SafeString(accountsByType, requestedType)
	_ = accountType
	if IsTrue(OpEq2(accountType, MkUndefined())) {
		keys := GoGetKeys(accountsByType)
		_ = keys
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" fetchBalance type parameter must be one of "), keys.Join(MkString(", "))))))
	}
	query := this.Omit(params, MkString("type"))
	_ = query
	response := this.Call(MkString("privatePostBalances"), query)
	_ = response
	result := MkMap(&VarMap{"info": response})
	_ = result
	isDerivative := OpEq2(requestedType, MkString("derivatives"))
	_ = isDerivative
	for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
		balance := *(response).At(i)
		_ = balance
		type_ := this.SafeString(balance, MkString("type"))
		_ = type_
		currencyId := this.SafeStringLower(balance, MkString("currency"), MkString(""))
		_ = currencyId
		start := OpSub(currencyId.Length, MkInteger(2))
		_ = start
		isDerivativeCode := OpEq2(currencyId.Slice(start), MkString("f0"))
		_ = isDerivativeCode
		derivativeCondition := OpOr(OpNot(isDerivative), isDerivativeCode)
		_ = derivativeCondition
		if IsTrue(OpAnd(OpEq2(accountType, type_), derivativeCondition)) {
			code := this.SafeCurrencyCode(currencyId)
			_ = code
			if IsTrue(OpNot(OpHasMember(code, result))) {
				account := this.Account()
				_ = account
				*(account).At(MkString("free")) = this.SafeString(balance, MkString("available"))
				*(account).At(MkString("total")) = this.SafeString(balance, MkString("amount"))
				*(result).At(code) = OpCopy(account)
			}
		}
	}
	return this.ParseBalance(result)
}

func (this *Bitfinex) Transfer(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	amount := GoGetArg(goArgs, 1, MkUndefined())
	_ = amount
	fromAccount := GoGetArg(goArgs, 2, MkUndefined())
	_ = fromAccount
	toAccount := GoGetArg(goArgs, 3, MkUndefined())
	_ = toAccount
	params := GoGetArg(goArgs, 4, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	accountsByType := this.SafeValue(*this.At(MkString("options")), MkString("accountsByType"), MkMap(&VarMap{}))
	_ = accountsByType
	fromId := this.SafeString(accountsByType, fromAccount)
	_ = fromId
	if IsTrue(OpEq2(fromId, MkUndefined())) {
		keys := GoGetKeys(accountsByType)
		_ = keys
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" transfer fromAccount must be one of "), keys.Join(MkString(", "))))))
	}
	toId := this.SafeString(accountsByType, toAccount)
	_ = toId
	if IsTrue(OpEq2(toId, MkUndefined())) {
		keys := GoGetKeys(accountsByType)
		_ = keys
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" transfer toAccount must be one of "), keys.Join(MkString(", "))))))
	}
	currency := this.Currency(code)
	_ = currency
	fromCurrencyId := this.ConvertDerivativesId(*(currency).At(MkString("id")), fromAccount)
	_ = fromCurrencyId
	toCurrencyId := this.ConvertDerivativesId(*(currency).At(MkString("id")), toAccount)
	_ = toCurrencyId
	requestedAmount := this.CurrencyToPrecision(code, amount)
	_ = requestedAmount
	request := MkMap(&VarMap{
		"amount":      requestedAmount,
		"currency":    fromCurrencyId,
		"currency_to": toCurrencyId,
		"walletfrom":  fromId,
		"walletto":    toId,
	})
	_ = request
	response := this.Call(MkString("privatePostTransfer"), this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkInteger(0))
	_ = result
	status := this.SafeString(result, MkString("status"))
	_ = status
	message := this.SafeString(result, MkString("message"))
	_ = message
	if IsTrue(OpEq2(message, MkUndefined())) {
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), MkString(" transfer failed"))))
	}
	if IsTrue(OpEq2(status, MkString("error"))) {
		this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("exact")), message, OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), message)))
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), message))))
	}
	return MkMap(&VarMap{
		"info":        response,
		"status":      status,
		"amount":      requestedAmount,
		"code":        code,
		"fromAccount": fromAccount,
		"toAccount":   toAccount,
		"timestamp":   MkUndefined(),
		"datetime":    MkUndefined(),
	})
}

func (this *Bitfinex) ConvertDerivativesId(goArgs ...*Variant) *Variant {
	currencyId := GoGetArg(goArgs, 0, MkUndefined())
	_ = currencyId
	type_ := GoGetArg(goArgs, 1, MkUndefined())
	_ = type_
	start := OpSub(currencyId.Length, MkInteger(2))
	_ = start
	isDerivativeCode := OpEq2(currencyId.Slice(start), MkString("F0"))
	_ = isDerivativeCode
	if IsTrue(OpAnd(OpAnd(OpNotEq2(type_, MkString("derivatives")), OpAnd(OpNotEq2(type_, MkString("trading")), OpNotEq2(type_, MkString("margin")))), isDerivativeCode)) {
		currencyId = currencyId.Slice(MkInteger(0), start)
	} else {
		if IsTrue(OpAnd(OpEq2(type_, MkString("derivatives")), OpNot(isDerivativeCode))) {
			currencyId = OpAdd(currencyId, MkString("F0"))
		}
	}
	return currencyId
}

func (this *Bitfinex) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"symbol": this.MarketId(symbol)})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit_bids")) = OpCopy(limit)
		*(request).At(MkString("limit_asks")) = OpCopy(limit)
	}
	response := this.Call(MkString("publicGetBookSymbol"), this.Extend(request, params))
	_ = response
	return this.ParseOrderBook(response, symbol, MkUndefined(), MkString("bids"), MkString("asks"), MkString("price"), MkString("amount"))
}

func (this *Bitfinex) FetchTickers(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("publicGetTickers"), params)
	_ = response
	result := MkMap(&VarMap{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
		ticker := this.ParseTicker(*(response).At(i))
		_ = ticker
		symbol := *(ticker).At(MkString("symbol"))
		_ = symbol
		*(result).At(symbol) = OpCopy(ticker)
	}
	return this.FilterByArray(result, MkString("symbol"), symbols)
}

func (this *Bitfinex) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	ticker := this.Call(MkString("publicGetPubtickerSymbol"), this.Extend(request, params))
	_ = ticker
	return this.ParseTicker(ticker, market)
}

func (this *Bitfinex) ParseTicker(goArgs ...*Variant) *Variant {
	ticker := GoGetArg(goArgs, 0, MkUndefined())
	_ = ticker
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.SafeNumber(ticker, MkString("timestamp"))
	_ = timestamp
	if IsTrue(OpNotEq2(timestamp, MkUndefined())) {
		OpMultiAssign(&timestamp, MkInteger(1000))
	}
	timestamp = parseInt(timestamp)
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(OpNotEq2(market, MkUndefined())) {
		symbol = *(market).At(MkString("symbol"))
	} else {
		if IsTrue(OpHasMember(MkString("pair"), ticker)) {
			marketId := this.SafeString(ticker, MkString("pair"))
			_ = marketId
			if IsTrue(OpNotEq2(marketId, MkUndefined())) {
				if IsTrue(OpHasMember(marketId, *this.At(MkString("markets_by_id")))) {
					market = *(*this.At(MkString("markets_by_id"))).At(marketId)
					symbol = *(market).At(MkString("symbol"))
				} else {
					baseId := marketId.Slice(MkInteger(0), MkInteger(3))
					_ = baseId
					quoteId := marketId.Slice(MkInteger(3), MkInteger(6))
					_ = quoteId
					base := this.SafeCurrencyCode(baseId)
					_ = base
					quote := this.SafeCurrencyCode(quoteId)
					_ = quote
					symbol = OpAdd(base, OpAdd(MkString("/"), quote))
				}
			}
		}
	}
	last := this.SafeNumber(ticker, MkString("last_price"))
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
		"percentage":    MkUndefined(),
		"average":       this.SafeNumber(ticker, MkString("mid")),
		"baseVolume":    this.SafeNumber(ticker, MkString("volume")),
		"quoteVolume":   MkUndefined(),
		"info":          ticker,
	})
}

func (this *Bitfinex) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	id := this.SafeString(trade, MkString("tid"))
	_ = id
	timestamp := this.SafeTimestamp(trade, MkString("timestamp"))
	_ = timestamp
	type_ := OpCopy(MkUndefined())
	_ = type_
	side := this.SafeStringLower(trade, MkString("type"))
	_ = side
	orderId := this.SafeString(trade, MkString("order_id"))
	_ = orderId
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
	fee := OpCopy(MkUndefined())
	_ = fee
	if IsTrue(OpHasMember(MkString("fee_amount"), trade)) {
		feeCost := OpNeg(this.SafeNumber(trade, MkString("fee_amount")))
		_ = feeCost
		feeCurrencyId := this.SafeString(trade, MkString("fee_currency"))
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
		"symbol":       *(market).At(MkString("symbol")),
		"type":         type_,
		"order":        orderId,
		"side":         side,
		"takerOrMaker": MkUndefined(),
		"price":        price,
		"amount":       amount,
		"cost":         cost,
		"fee":          fee,
	})
}

func (this *Bitfinex) FetchTrades(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{
		"symbol":       *(market).At(MkString("id")),
		"limit_trades": limit,
	})
	_ = request
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("timestamp")) = parseInt(OpDiv(since, MkInteger(1000)))
	}
	response := this.Call(MkString("publicGetTradesSymbol"), this.Extend(request, params))
	_ = response
	return this.ParseTrades(response, market, since, limit)
}

func (this *Bitfinex) FetchMyTrades(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit_trades")) = OpCopy(limit)
	}
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("timestamp")) = parseInt(OpDiv(since, MkInteger(1000)))
	}
	response := this.Call(MkString("privatePostMytrades"), this.Extend(request, params))
	_ = response
	return this.ParseTrades(response, market, since, limit)
}

func (this *Bitfinex) CreateOrder(goArgs ...*Variant) *Variant {
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
		"symbol":         this.MarketId(symbol),
		"side":           side,
		"amount":         this.AmountToPrecision(symbol, amount),
		"type":           this.SafeString(*(*this.At(MkString("options"))).At(MkString("orderTypes")), type_, type_),
		"ocoorder":       MkBool(false),
		"buy_price_oco":  MkInteger(0),
		"sell_price_oco": MkInteger(0),
	})
	_ = request
	if IsTrue(OpEq2(type_, MkString("market"))) {
		*(request).At(MkString("price")) = (this.Nonce()).Call(MkString("toString"))
	} else {
		*(request).At(MkString("price")) = this.PriceToPrecision(symbol, price)
	}
	response := this.Call(MkString("privatePostOrderNew"), this.Extend(request, params))
	_ = response
	return this.ParseOrder(response)
}

func (this *Bitfinex) EditOrder(goArgs ...*Variant) *Variant {
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
	order := MkMap(&VarMap{"order_id": parseInt(id)})
	_ = order
	if IsTrue(OpNotEq2(price, MkUndefined())) {
		*(order).At(MkString("price")) = this.PriceToPrecision(symbol, price)
	}
	if IsTrue(OpNotEq2(amount, MkUndefined())) {
		*(order).At(MkString("amount")) = this.NumberToString(amount)
	}
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		*(order).At(MkString("symbol")) = this.MarketId(symbol)
	}
	if IsTrue(OpNotEq2(side, MkUndefined())) {
		*(order).At(MkString("side")) = OpCopy(side)
	}
	if IsTrue(OpNotEq2(type_, MkUndefined())) {
		*(order).At(MkString("type")) = this.SafeString(*(*this.At(MkString("options"))).At(MkString("orderTypes")), type_, type_)
	}
	response := this.Call(MkString("privatePostOrderCancelReplace"), this.Extend(order, params))
	_ = response
	return this.ParseOrder(response)
}

func (this *Bitfinex) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"order_id": parseInt(id)})
	_ = request
	return this.Call(MkString("privatePostOrderCancel"), this.Extend(request, params))
}

func (this *Bitfinex) CancelAllOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	return this.Call(MkString("privatePostOrderCancelAll"), params)
}

func (this *Bitfinex) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	side := this.SafeString(order, MkString("side"))
	_ = side
	open := this.SafeValue(order, MkString("is_live"))
	_ = open
	canceled := this.SafeValue(order, MkString("is_cancelled"))
	_ = canceled
	status := OpCopy(MkUndefined())
	_ = status
	if IsTrue(open) {
		status = MkString("open")
	} else {
		if IsTrue(canceled) {
			status = MkString("canceled")
		} else {
			status = MkString("closed")
		}
	}
	marketId := this.SafeStringUpper(order, MkString("symbol"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market)
	_ = symbol
	orderType := this.SafeString(order, MkString("type"), MkString(""))
	_ = orderType
	exchange := OpGtEq(orderType.IndexOf(MkString("exchange ")), MkInteger(0))
	_ = exchange
	if IsTrue(exchange) {
		parts := (*(order).At(MkString("type"))).Call(MkString("split"), MkString(" "))
		_ = parts
		orderType = *(parts).At(MkInteger(1))
	}
	timestamp := this.SafeTimestamp(order, MkString("timestamp"))
	_ = timestamp
	id := this.SafeString(order, MkString("id"))
	_ = id
	return this.SafeOrder(MkMap(&VarMap{
		"info":               order,
		"id":                 id,
		"clientOrderId":      MkUndefined(),
		"timestamp":          timestamp,
		"datetime":           this.Iso8601(timestamp),
		"lastTradeTimestamp": MkUndefined(),
		"symbol":             symbol,
		"type":               orderType,
		"timeInForce":        MkUndefined(),
		"postOnly":           MkUndefined(),
		"side":               side,
		"price":              this.SafeNumber(order, MkString("price")),
		"stopPrice":          MkUndefined(),
		"average":            this.SafeNumber(order, MkString("avg_execution_price")),
		"amount":             this.SafeNumber(order, MkString("original_amount")),
		"remaining":          this.SafeNumber(order, MkString("remaining_amount")),
		"filled":             this.SafeNumber(order, MkString("executed_amount")),
		"status":             status,
		"fee":                MkUndefined(),
		"cost":               MkUndefined(),
		"trades":             MkUndefined(),
	}))
}

func (this *Bitfinex) FetchOpenOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		if IsTrue(OpNot(OpHasMember(symbol, *this.At(MkString("markets"))))) {
			panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" has no symbol "), symbol))))
		}
	}
	response := this.Call(MkString("privatePostOrders"), params)
	_ = response
	orders := this.ParseOrders(response, MkUndefined(), since, limit)
	_ = orders
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		orders = this.FilterBy(orders, MkString("symbol"), symbol)
	}
	return orders
}

func (this *Bitfinex) FetchClosedOrders(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("privatePostOrdersHist"), this.Extend(request, params))
	_ = response
	orders := this.ParseOrders(response, MkUndefined(), since, limit)
	_ = orders
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		orders = this.FilterBy(orders, MkString("symbol"), symbol)
	}
	orders = this.FilterByArray(orders, MkString("status"), MkArray(&VarArray{
		MkString("closed"),
		MkString("canceled"),
	}), MkBool(false))
	return orders
}

func (this *Bitfinex) FetchOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"order_id": parseInt(id)})
	_ = request
	response := this.Call(MkString("privatePostOrderStatus"), this.Extend(request, params))
	_ = response
	return this.ParseOrder(response)
}

func (this *Bitfinex) ParseOHLCV(goArgs ...*Variant) *Variant {
	ohlcv := GoGetArg(goArgs, 0, MkUndefined())
	_ = ohlcv
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	return MkArray(&VarArray{
		this.SafeInteger(ohlcv, MkInteger(0)),
		this.SafeNumber(ohlcv, MkInteger(1)),
		this.SafeNumber(ohlcv, MkInteger(3)),
		this.SafeNumber(ohlcv, MkInteger(4)),
		this.SafeNumber(ohlcv, MkInteger(2)),
		this.SafeNumber(ohlcv, MkInteger(5)),
	})
}

func (this *Bitfinex) FetchOHLCV(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpEq2(limit, MkUndefined())) {
		limit = MkInteger(100)
	}
	market := this.Market(symbol)
	_ = market
	v2id := OpAdd(MkString("t"), *(market).At(MkString("id")))
	_ = v2id
	request := MkMap(&VarMap{
		"symbol":    v2id,
		"timeframe": *(*this.At(MkString("timeframes"))).At(timeframe),
		"sort":      MkInteger(1),
		"limit":     limit,
	})
	_ = request
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("start")) = OpCopy(since)
	}
	response := this.Call(MkString("v2GetCandlesTradeTimeframeSymbolHist"), this.Extend(request, params))
	_ = response
	return this.ParseOHLCVs(response, market, timeframe, since, limit)
}

func (this *Bitfinex) GetCurrencyName(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	if IsTrue(OpHasMember(code, *(*this.At(MkString("options"))).At(MkString("currencyNames")))) {
		return *(*(*this.At(MkString("options"))).At(MkString("currencyNames"))).At(code)
	}
	panic(NewNotSupported(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), OpAdd(code, MkString(" not supported for withdrawal"))))))
	return MkUndefined()
}

func (this *Bitfinex) CreateDepositAddress(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"renew": MkInteger(1)})
	_ = request
	response := this.FetchDepositAddress(code, this.Extend(request, params))
	_ = response
	return response
}

func (this *Bitfinex) FetchDepositAddress(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	name := this.GetCurrencyName(code)
	_ = name
	request := MkMap(&VarMap{
		"method":      name,
		"wallet_name": MkString("exchange"),
		"renew":       MkInteger(0),
	})
	_ = request
	response := this.Call(MkString("privatePostDepositNew"), this.Extend(request, params))
	_ = response
	address := this.SafeValue(response, MkString("address"))
	_ = address
	tag := OpCopy(MkUndefined())
	_ = tag
	if IsTrue(OpHasMember(MkString("address_pool"), response)) {
		tag = OpCopy(address)
		address = *(response).At(MkString("address_pool"))
	}
	this.CheckAddress(address)
	return MkMap(&VarMap{
		"currency": code,
		"address":  address,
		"tag":      tag,
		"info":     response,
	})
}

func (this *Bitfinex) FetchTransactions(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	currencyId := this.SafeString(params, MkString("currency"))
	_ = currencyId
	query := this.Omit(params, MkString("currency"))
	_ = query
	currency := OpCopy(MkUndefined())
	_ = currency
	if IsTrue(OpEq2(currencyId, MkUndefined())) {
		if IsTrue(OpEq2(code, MkUndefined())) {
			panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchTransactions() requires a currency `code` argument or a `currency` parameter"))))
		} else {
			currency = this.Currency(code)
			currencyId = *(currency).At(MkString("id"))
		}
	}
	*(query).At(MkString("currency")) = OpCopy(currencyId)
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(query).At(MkString("since")) = parseInt(OpDiv(since, MkInteger(1000)))
	}
	response := this.Call(MkString("privatePostHistoryMovements"), this.Extend(query, params))
	_ = response
	return this.ParseTransactions(response, currency, since, limit)
}

func (this *Bitfinex) ParseTransaction(goArgs ...*Variant) *Variant {
	transaction := GoGetArg(goArgs, 0, MkUndefined())
	_ = transaction
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	timestamp := this.SafeTimestamp(transaction, MkString("timestamp_created"))
	_ = timestamp
	updated := this.SafeTimestamp(transaction, MkString("timestamp"))
	_ = updated
	currencyId := this.SafeString(transaction, MkString("currency"))
	_ = currencyId
	code := this.SafeCurrencyCode(currencyId, currency)
	_ = code
	type_ := this.SafeStringLower(transaction, MkString("type"))
	_ = type_
	status := this.ParseTransactionStatus(this.SafeString(transaction, MkString("status")))
	_ = status
	feeCost := this.SafeNumber(transaction, MkString("fee"))
	_ = feeCost
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		feeCost = Math.Abs(feeCost)
	}
	return MkMap(&VarMap{
		"info":      transaction,
		"id":        this.SafeString(transaction, MkString("id")),
		"txid":      this.SafeString(transaction, MkString("txid")),
		"timestamp": timestamp,
		"datetime":  this.Iso8601(timestamp),
		"address":   this.SafeString(transaction, MkString("address")),
		"tag":       MkUndefined(),
		"type":      type_,
		"amount":    this.SafeNumber(transaction, MkString("amount")),
		"currency":  code,
		"status":    status,
		"updated":   updated,
		"fee": MkMap(&VarMap{
			"currency": code,
			"cost":     feeCost,
			"rate":     MkUndefined(),
		}),
	})
}

func (this *Bitfinex) ParseTransactionStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"SENDING":       MkString("pending"),
		"CANCELED":      MkString("canceled"),
		"ZEROCONFIRMED": MkString("failed"),
		"COMPLETED":     MkString("ok"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Bitfinex) Withdraw(goArgs ...*Variant) *Variant {
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
	name := this.GetCurrencyName(code)
	_ = name
	request := MkMap(&VarMap{
		"withdraw_type":  name,
		"walletselected": MkString("exchange"),
		"amount":         this.NumberToString(amount),
		"address":        address,
	})
	_ = request
	if IsTrue(OpNotEq2(tag, MkUndefined())) {
		*(request).At(MkString("payment_id")) = OpCopy(tag)
	}
	responses := this.Call(MkString("privatePostWithdraw"), this.Extend(request, params))
	_ = responses
	response := *(responses).At(MkInteger(0))
	_ = response
	id := this.SafeString(response, MkString("withdrawal_id"))
	_ = id
	message := this.SafeString(response, MkString("message"))
	_ = message
	errorMessage := this.FindBroadlyMatchedKey(*(*this.At(MkString("exceptions"))).At(MkString("broad")), message)
	_ = errorMessage
	if IsTrue(OpEq2(id, MkInteger(0))) {
		if IsTrue(OpNotEq2(errorMessage, MkUndefined())) {
			ExceptionClass := *(*(*this.At(MkString("exceptions"))).At(MkString("broad"))).At(errorMessage)
			_ = ExceptionClass
			panic(NewExceptionClass(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), message))))
		}
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" withdraw returned an id of zero: "), this.Json(response)))))
	}
	return MkMap(&VarMap{
		"info": response,
		"id":   id,
	})
}

func (this *Bitfinex) FetchPositions(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privatePostPositions"), params)
	_ = response
	return response
}

func (this *Bitfinex) Nonce(goArgs ...*Variant) *Variant {
	return this.Milliseconds()
}

func (this *Bitfinex) Sign(goArgs ...*Variant) *Variant {
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
	request := OpAdd(MkString("/"), this.ImplodeParams(path, params))
	_ = request
	if IsTrue(OpEq2(api, MkString("v2"))) {
		request = OpAdd(MkString("/"), OpAdd(api, request))
	} else {
		request = OpAdd(MkString("/"), OpAdd(*this.At(MkString("version")), request))
	}
	query := this.Omit(params, this.ExtractParams(path))
	_ = query
	url := OpAdd(*(*(*this.At(MkString("urls"))).At(MkString("api"))).At(api), request)
	_ = url
	if IsTrue(OpOr(OpEq2(api, MkString("public")), OpGtEq(path.IndexOf(MkString("/hist")), MkInteger(0)))) {
		if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
			suffix := OpAdd(MkString("?"), this.Urlencode(query))
			_ = suffix
			OpAddAssign(&url, suffix)
			OpAddAssign(&request, suffix)
		}
	}
	if IsTrue(OpEq2(api, MkString("private"))) {
		this.CheckRequiredCredentials()
		nonce := this.Nonce()
		_ = nonce
		query = this.Extend(MkMap(&VarMap{
			"nonce":   nonce.ToString(),
			"request": request,
		}), query)
		body = this.Json(query)
		payload := this.StringToBase64(body)
		_ = payload
		secret := this.Encode(*this.At(MkString("secret")))
		_ = secret
		signature := this.Hmac(payload, secret, MkString("sha384"))
		_ = signature
		headers = MkMap(&VarMap{
			"X-BFX-APIKEY":    *this.At(MkString("apiKey")),
			"X-BFX-PAYLOAD":   this.Decode(payload),
			"X-BFX-SIGNATURE": signature,
			"Content-Type":    MkString("application/json"),
		})
	}
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Bitfinex) HandleErrors(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpGtEq(code, MkInteger(400))) {
		if IsTrue(OpEq2(*(body).At(MkInteger(0)), MkString("{"))) {
			feedback := OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))
			_ = feedback
			message := this.SafeString2(response, MkString("message"), MkString("error"))
			_ = message
			this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("exact")), message, feedback)
			this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("broad")), message, feedback)
			panic(NewExchangeError(feedback))
		}
	}
	return MkUndefined()
}
