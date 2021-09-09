package ccxt

import ()

type Bibox struct {
	*ExchangeBase
}

var _ Exchange = (*Bibox)(nil)

func init() {
	exchange := &Bibox{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Bibox) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("bibox"),
		"name": MkString("Bibox"),
		"countries": MkArray(&VarArray{
			MkString("CN"),
			MkString("US"),
			MkString("KR"),
		}),
		"version":  MkString("v1"),
		"hostname": MkString("bibox365.com"),
		"has": MkMap(&VarMap{
			"cancelOrder":         MkBool(true),
			"CORS":                MkBool(false),
			"createMarketOrder":   MkBool(false),
			"createOrder":         MkBool(true),
			"fetchBalance":        MkBool(true),
			"fetchClosedOrders":   MkBool(true),
			"fetchCurrencies":     MkBool(true),
			"fetchDeposits":       MkBool(true),
			"fetchDepositAddress": MkBool(true),
			"fetchFundingFees":    MkBool(true),
			"fetchMarkets":        MkBool(true),
			"fetchMyTrades":       MkBool(true),
			"fetchOHLCV":          MkBool(true),
			"fetchOpenOrders":     MkBool(true),
			"fetchOrder":          MkBool(true),
			"fetchOrderBook":      MkBool(true),
			"fetchTicker":         MkBool(true),
			"fetchTickers":        MkBool(true),
			"fetchTrades":         MkBool(true),
			"fetchWithdrawals":    MkBool(true),
			"publicAPI":           MkBool(false),
			"withdraw":            MkBool(true),
		}),
		"timeframes": MkMap(&VarMap{
			"1m":  MkString("1min"),
			"5m":  MkString("5min"),
			"15m": MkString("15min"),
			"30m": MkString("30min"),
			"1h":  MkString("1hour"),
			"2h":  MkString("2hour"),
			"4h":  MkString("4hour"),
			"6h":  MkString("6hour"),
			"12h": MkString("12hour"),
			"1d":  MkString("day"),
			"1w":  MkString("week"),
		}),
		"urls": MkMap(&VarMap{
			"logo": MkString("https://user-images.githubusercontent.com/51840849/77257418-3262b000-6c85-11ea-8fb8-20bdf20b3592.jpg"),
			"api":  MkString("https://api.{hostname}"),
			"www":  MkString("https://www.bibox365.com"),
			"doc": MkArray(&VarArray{
				MkString("https://biboxcom.github.io/en/"),
			}),
			"fees":     MkString("https://bibox.zendesk.com/hc/en-us/articles/360002336133"),
			"referral": MkString("https://w2.bibox365.com/login/register?invite_code=05Kj3I"),
		}),
		"api": MkMap(&VarMap{
			"public": MkMap(&VarMap{
				"post": MkArray(&VarArray{
					MkString("mdata"),
				}),
				"get": MkArray(&VarArray{
					MkString("cquery"),
					MkString("mdata"),
					MkString("cdata"),
				}),
			}),
			"private": MkMap(&VarMap{"post": MkArray(&VarArray{
				MkString("cquery"),
				MkString("ctrade"),
				MkString("user"),
				MkString("orderpending"),
				MkString("transfer"),
			})}),
			"v2private": MkMap(&VarMap{"post": MkArray(&VarArray{
				MkString("assets/transfer/spot"),
			})}),
		}),
		"fees": MkMap(&VarMap{
			"trading": MkMap(&VarMap{
				"tierBased":  MkBool(false),
				"percentage": MkBool(true),
				"taker":      this.ParseNumber(MkString("0.001")),
				"maker":      this.ParseNumber(MkString("0.0008")),
			}),
			"funding": MkMap(&VarMap{
				"tierBased":  MkBool(false),
				"percentage": MkBool(false),
				"withdraw":   MkMap(&VarMap{}),
				"deposit":    MkMap(&VarMap{}),
			}),
		}),
		"exceptions": MkMap(&VarMap{
			"2011": AccountSuspended,
			"2015": AuthenticationError,
			"2021": InsufficientFunds,
			"2027": InsufficientFunds,
			"2033": OrderNotFound,
			"2065": InvalidOrder,
			"2066": InvalidOrder,
			"2067": InvalidOrder,
			"2068": InvalidOrder,
			"2078": InvalidOrder,
			"2085": InvalidOrder,
			"2091": RateLimitExceeded,
			"2092": InvalidOrder,
			"2131": InvalidOrder,
			"3000": BadRequest,
			"3002": BadRequest,
			"3012": AuthenticationError,
			"3016": BadSymbol,
			"3024": PermissionDenied,
			"3025": AuthenticationError,
			"4000": ExchangeNotAvailable,
			"4003": DDoSProtection,
		}),
		"commonCurrencies": MkMap(&VarMap{
			"APENFT(NFT)": MkString("NFT"),
			"BOX":         MkString("DefiBox"),
			"BPT":         MkString("BlockPool Token"),
			"GTC":         MkString("Game.com"),
			"KEY":         MkString("Bihu"),
			"MTC":         MkString("MTC Mesh Network"),
			"NFT":         MkString("NFT Protocol"),
			"PAI":         MkString("PCHAIN"),
			"TERN":        MkString("Ternio-ERC20"),
		}),
		"options": MkMap(&VarMap{"fetchCurrencies": MkString("fetch_currencies_public")}),
	}))
}

func (this *Bibox) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	request := MkMap(&VarMap{"cmd": MkString("pairList")})
	_ = request
	response := this.Call(MkString("publicGetMdata"), this.Extend(request, params))
	_ = response
	markets := this.SafeValue(response, MkString("result"))
	_ = markets
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, markets.Length)); OpIncr(&i) {
		market := *(markets).At(i)
		_ = market
		numericId := this.SafeInteger(market, MkString("id"))
		_ = numericId
		id := this.SafeString(market, MkString("pair"))
		_ = id
		baseId := OpCopy(MkUndefined())
		_ = baseId
		quoteId := OpCopy(MkUndefined())
		_ = quoteId
		if IsTrue(OpNotEq2(id, MkUndefined())) {
			parts := id.Split(MkString("_"))
			_ = parts
			baseId = this.SafeString(parts, MkInteger(0))
			quoteId = this.SafeString(parts, MkInteger(1))
		}
		base := this.SafeCurrencyCode(baseId)
		_ = base
		quote := this.SafeCurrencyCode(quoteId)
		_ = quote
		symbol := OpAdd(base, OpAdd(MkString("/"), quote))
		_ = symbol
		precision := MkMap(&VarMap{
			"amount": this.SafeNumber(market, MkString("amount_scale")),
			"price":  this.SafeNumber(market, MkString("decimal")),
		})
		_ = precision
		result.Push(MkMap(&VarMap{
			"id":        id,
			"numericId": numericId,
			"symbol":    symbol,
			"base":      base,
			"quote":     quote,
			"baseId":    baseId,
			"quoteId":   quoteId,
			"active":    MkBool(true),
			"info":      market,
			"precision": precision,
			"limits": MkMap(&VarMap{
				"amount": MkMap(&VarMap{
					"min": Math.Pow(MkInteger(10), OpNeg(*(precision).At(MkString("amount")))),
					"max": MkUndefined(),
				}),
				"price": MkMap(&VarMap{
					"min": Math.Pow(MkInteger(10), OpNeg(*(precision).At(MkString("price")))),
					"max": MkUndefined(),
				}),
			}),
		}))
	}
	return result
}

func (this *Bibox) ParseTicker(goArgs ...*Variant) *Variant {
	ticker := GoGetArg(goArgs, 0, MkUndefined())
	_ = ticker
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.SafeInteger(ticker, MkString("timestamp"))
	_ = timestamp
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(OpNotEq2(market, MkUndefined())) {
		symbol = *(market).At(MkString("symbol"))
	} else {
		baseId := this.SafeString(ticker, MkString("coin_symbol"))
		_ = baseId
		quoteId := this.SafeString(ticker, MkString("currency_symbol"))
		_ = quoteId
		base := this.SafeCurrencyCode(baseId)
		_ = base
		quote := this.SafeCurrencyCode(quoteId)
		_ = quote
		symbol = OpAdd(base, OpAdd(MkString("/"), quote))
	}
	last := this.SafeNumber(ticker, MkString("last"))
	_ = last
	change := this.SafeNumber(ticker, MkString("change"))
	_ = change
	baseVolume := this.SafeNumber2(ticker, MkString("vol"), MkString("vol24H"))
	_ = baseVolume
	open := OpCopy(MkUndefined())
	_ = open
	if IsTrue(OpAnd(OpNotEq2(last, MkUndefined()), OpNotEq2(change, MkUndefined()))) {
		open = OpSub(last, change)
	}
	percentage := this.SafeString(ticker, MkString("percent"))
	_ = percentage
	if IsTrue(OpNotEq2(percentage, MkUndefined())) {
		percentage = percentage.Replace(MkString("%"), MkString(""))
		percentage = this.ParseNumber(percentage)
	}
	return MkMap(&VarMap{
		"symbol":        symbol,
		"timestamp":     timestamp,
		"datetime":      this.Iso8601(timestamp),
		"high":          this.SafeNumber(ticker, MkString("high")),
		"low":           this.SafeNumber(ticker, MkString("low")),
		"bid":           this.SafeNumber(ticker, MkString("buy")),
		"bidVolume":     MkUndefined(),
		"ask":           this.SafeNumber(ticker, MkString("sell")),
		"askVolume":     MkUndefined(),
		"vwap":          MkUndefined(),
		"open":          open,
		"close":         last,
		"last":          last,
		"previousClose": MkUndefined(),
		"change":        change,
		"percentage":    percentage,
		"average":       MkUndefined(),
		"baseVolume":    baseVolume,
		"quoteVolume":   this.SafeNumber(ticker, MkString("amount")),
		"info":          ticker,
	})
}

func (this *Bibox) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{
		"cmd":  MkString("ticker"),
		"pair": *(market).At(MkString("id")),
	})
	_ = request
	response := this.Call(MkString("publicGetMdata"), this.Extend(request, params))
	_ = response
	return this.ParseTicker(*(response).At(MkString("result")), market)
}

func (this *Bibox) FetchTickers(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	request := MkMap(&VarMap{"cmd": MkString("marketAll")})
	_ = request
	response := this.Call(MkString("publicGetMdata"), this.Extend(request, params))
	_ = response
	tickers := this.ParseTickers(*(response).At(MkString("result")), symbols)
	_ = tickers
	result := this.IndexBy(tickers, MkString("symbol"))
	_ = result
	return this.FilterByArray(result, MkString("symbol"), symbols)
}

func (this *Bibox) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.SafeInteger2(trade, MkString("time"), MkString("createdAt"))
	_ = timestamp
	side := this.SafeInteger2(trade, MkString("side"), MkString("order_side"))
	_ = side
	side = OpTrinary(OpEq2(side, MkInteger(1)), MkString("buy"), MkString("sell"))
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(OpEq2(market, MkUndefined())) {
		marketId := this.SafeString(trade, MkString("pair"))
		_ = marketId
		if IsTrue(OpEq2(marketId, MkUndefined())) {
			baseId := this.SafeString(trade, MkString("coin_symbol"))
			_ = baseId
			quoteId := this.SafeString(trade, MkString("currency_symbol"))
			_ = quoteId
			if IsTrue(OpAnd(OpNotEq2(baseId, MkUndefined()), OpNotEq2(quoteId, MkUndefined()))) {
				marketId = OpAdd(baseId, OpAdd(MkString("_"), quoteId))
			}
		}
		if IsTrue(OpHasMember(marketId, *this.At(MkString("markets_by_id")))) {
			market = *(*this.At(MkString("markets_by_id"))).At(marketId)
		}
	}
	if IsTrue(OpNotEq2(market, MkUndefined())) {
		symbol = *(market).At(MkString("symbol"))
	}
	fee := OpCopy(MkUndefined())
	_ = fee
	feeCostString := this.SafeString(trade, MkString("fee"))
	_ = feeCostString
	feeCurrency := this.SafeString(trade, MkString("fee_symbol"))
	_ = feeCurrency
	if IsTrue(OpNotEq2(feeCurrency, MkUndefined())) {
		if IsTrue(OpHasMember(feeCurrency, *this.At(MkString("currencies_by_id")))) {
			feeCurrency = *(*(*this.At(MkString("currencies_by_id"))).At(feeCurrency)).At(MkString("code"))
		} else {
			feeCurrency = this.SafeCurrencyCode(feeCurrency)
		}
	}
	feeRate := OpCopy(MkUndefined())
	_ = feeRate
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
	if IsTrue(OpNotEq2(feeCostString, MkUndefined())) {
		fee = MkMap(&VarMap{
			"cost":     this.ParseNumber(Precise.StringNeg(feeCostString)),
			"currency": feeCurrency,
			"rate":     feeRate,
		})
	}
	id := this.SafeString(trade, MkString("id"))
	_ = id
	return MkMap(&VarMap{
		"info":         trade,
		"id":           id,
		"order":        MkUndefined(),
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       symbol,
		"type":         MkString("limit"),
		"takerOrMaker": MkUndefined(),
		"side":         side,
		"price":        price,
		"amount":       amount,
		"cost":         cost,
		"fee":          fee,
	})
}

func (this *Bibox) FetchTrades(goArgs ...*Variant) *Variant {
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
		"cmd":  MkString("deals"),
		"pair": *(market).At(MkString("id")),
	})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("size")) = OpCopy(limit)
	}
	response := this.Call(MkString("publicGetMdata"), this.Extend(request, params))
	_ = response
	return this.ParseTrades(*(response).At(MkString("result")), market, since, limit)
}

func (this *Bibox) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{
		"cmd":  MkString("depth"),
		"pair": *(market).At(MkString("id")),
	})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("size")) = OpCopy(limit)
	}
	response := this.Call(MkString("publicGetMdata"), this.Extend(request, params))
	_ = response
	return this.ParseOrderBook(*(response).At(MkString("result")), symbol, this.SafeNumber(*(response).At(MkString("result")), MkString("update_time")), MkString("bids"), MkString("asks"), MkString("price"), MkString("volume"))
}

func (this *Bibox) ParseOHLCV(goArgs ...*Variant) *Variant {
	ohlcv := GoGetArg(goArgs, 0, MkUndefined())
	_ = ohlcv
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	return MkArray(&VarArray{
		this.SafeInteger(ohlcv, MkString("time")),
		this.SafeNumber(ohlcv, MkString("open")),
		this.SafeNumber(ohlcv, MkString("high")),
		this.SafeNumber(ohlcv, MkString("low")),
		this.SafeNumber(ohlcv, MkString("close")),
		this.SafeNumber(ohlcv, MkString("vol")),
	})
}

func (this *Bibox) FetchOHLCV(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	timeframe := GoGetArg(goArgs, 1, MkString("1m"))
	_ = timeframe
	since := GoGetArg(goArgs, 2, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 3, MkInteger(1000))
	_ = limit
	params := GoGetArg(goArgs, 4, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{
		"cmd":    MkString("kline"),
		"pair":   *(market).At(MkString("id")),
		"period": *(*this.At(MkString("timeframes"))).At(timeframe),
		"size":   limit,
	})
	_ = request
	response := this.Call(MkString("publicGetMdata"), this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"), MkArray(&VarArray{}))
	_ = result
	return this.ParseOHLCVs(result, market, timeframe, since, limit)
}

func (this *Bibox) FetchCurrencies(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	method := this.SafeString(*this.At(MkString("options")), MkString("fetchCurrencies"), MkString("fetch_currencies_public"))
	_ = method
	return this.Call(method, params)
}

func (this *Bibox) FetchCurrenciesPublic(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	request := MkMap(&VarMap{"cmd": MkString("currencies")})
	_ = request
	response := this.Call(MkString("publicGetCdata"), this.Extend(request, params))
	_ = response
	currencies := this.SafeValue(response, MkString("result"))
	_ = currencies
	result := MkMap(&VarMap{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, currencies.Length)); OpIncr(&i) {
		currency := *(currencies).At(i)
		_ = currency
		id := this.SafeString(currency, MkString("symbol"))
		_ = id
		name := this.SafeString(currency, MkString("name"))
		_ = name
		code := this.SafeCurrencyCode(id)
		_ = code
		precision := this.SafeInteger(currency, MkString("valid_decimals"))
		_ = precision
		deposit := this.SafeValue(currency, MkString("enable_deposit"))
		_ = deposit
		withdraw := this.SafeValue(currency, MkString("enable_withdraw"))
		_ = withdraw
		active := OpAnd(deposit, withdraw)
		_ = active
		*(result).At(code) = MkMap(&VarMap{
			"id":        id,
			"code":      code,
			"info":      currency,
			"name":      name,
			"active":    active,
			"fee":       MkUndefined(),
			"precision": precision,
			"limits": MkMap(&VarMap{
				"amount": MkMap(&VarMap{
					"min": Math.Pow(MkInteger(10), OpNeg(precision)),
					"max": MkUndefined(),
				}),
				"withdraw": MkMap(&VarMap{
					"min": this.SafeNumber(currency, MkString("withdraw_min")),
					"max": MkUndefined(),
				}),
			}),
		})
	}
	return result
}

func (this *Bibox) FetchCurrenciesPrivate(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpOr(OpNot(*this.At(MkString("apiKey"))), OpNot(*this.At(MkString("secret"))))) {
		panic(NewAuthenticationError(OpAdd(*this.At(MkString("id")), MkString(" fetchCurrencies is an authenticated endpoint, therefore it requires 'apiKey' and 'secret' credentials. If you don't need currency details, set exchange.has['fetchCurrencies'] = false before calling its methods."))))
	}
	request := MkMap(&VarMap{
		"cmd":  MkString("transfer/coinList"),
		"body": MkMap(&VarMap{}),
	})
	_ = request
	response := this.Call(MkString("privatePostTransfer"), this.Extend(request, params))
	_ = response
	currencies := this.SafeValue(response, MkString("result"))
	_ = currencies
	result := MkMap(&VarMap{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, currencies.Length)); OpIncr(&i) {
		currency := *(currencies).At(i)
		_ = currency
		id := this.SafeString(currency, MkString("symbol"))
		_ = id
		name := *(currency).At(MkString("name"))
		_ = name
		code := this.SafeCurrencyCode(id)
		_ = code
		precision := MkInteger(8)
		_ = precision
		deposit := this.SafeValue(currency, MkString("enable_deposit"))
		_ = deposit
		withdraw := this.SafeValue(currency, MkString("enable_withdraw"))
		_ = withdraw
		active := OpAnd(deposit, withdraw)
		_ = active
		*(result).At(code) = MkMap(&VarMap{
			"id":        id,
			"code":      code,
			"info":      currency,
			"name":      name,
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

func (this *Bibox) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	type_ := this.SafeString(params, MkString("type"), MkString("assets"))
	_ = type_
	params = this.Omit(params, MkString("type"))
	request := MkMap(&VarMap{
		"cmd":  OpAdd(MkString("transfer/"), type_),
		"body": this.Extend(MkMap(&VarMap{"select": MkInteger(1)}), params),
	})
	_ = request
	response := this.Call(MkString("privatePostTransfer"), request)
	_ = response
	balances := this.SafeValue(response, MkString("result"))
	_ = balances
	result := MkMap(&VarMap{"info": balances})
	_ = result
	indexed := OpCopy(MkUndefined())
	_ = indexed
	if IsTrue(OpHasMember(MkString("assets_list"), balances)) {
		indexed = this.IndexBy(*(balances).At(MkString("assets_list")), MkString("coin_symbol"))
	} else {
		indexed = OpCopy(balances)
	}
	keys := GoGetKeys(indexed)
	_ = keys
	for i := MkInteger(0); IsTrue(OpLw(i, keys.Length)); OpIncr(&i) {
		id := *(keys).At(i)
		_ = id
		code := id.ToUpperCase()
		_ = code
		if IsTrue(OpGtEq(code.IndexOf(MkString("TOTAL_")), MkInteger(0))) {
			code = code.Slice(MkInteger(6))
		}
		if IsTrue(OpHasMember(code, *this.At(MkString("currencies_by_id")))) {
			code = *(*(*this.At(MkString("currencies_by_id"))).At(code)).At(MkString("code"))
		}
		account := this.Account()
		_ = account
		balance := *(indexed).At(id)
		_ = balance
		if IsTrue(OpEq2(OpType(balance), MkString("string"))) {
			*(account).At(MkString("free")) = OpCopy(balance)
			*(account).At(MkString("total")) = OpCopy(balance)
		} else {
			*(account).At(MkString("free")) = this.SafeString(balance, MkString("balance"))
			*(account).At(MkString("used")) = this.SafeString(balance, MkString("freeze"))
		}
		*(result).At(code) = OpCopy(account)
	}
	return this.ParseBalance(result)
}

func (this *Bibox) FetchDeposits(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"page": MkInteger(1)})
	_ = request
	if IsTrue(OpNotEq2(code, MkUndefined())) {
		currency = this.Currency(code)
		*(request).At(MkString("symbol")) = *(currency).At(MkString("id"))
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("size")) = OpCopy(limit)
	} else {
		*(request).At(MkString("size")) = MkInteger(100)
	}
	response := this.Call(MkString("privatePostTransfer"), MkMap(&VarMap{
		"cmd":  MkString("transfer/transferInList"),
		"body": this.Extend(request, params),
	}))
	_ = response
	deposits := this.SafeValue(*(response).At(MkString("result")), MkString("items"), MkArray(&VarArray{}))
	_ = deposits
	for i := MkInteger(0); IsTrue(OpLw(i, deposits.Length)); OpIncr(&i) {
		*(*(deposits).At(i)).At(MkString("type")) = MkString("deposit")
	}
	return this.ParseTransactions(deposits, currency, since, limit)
}

func (this *Bibox) FetchWithdrawals(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"page": MkInteger(1)})
	_ = request
	if IsTrue(OpNotEq2(code, MkUndefined())) {
		currency = this.Currency(code)
		*(request).At(MkString("symbol")) = *(currency).At(MkString("id"))
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("size")) = OpCopy(limit)
	} else {
		*(request).At(MkString("size")) = MkInteger(100)
	}
	response := this.Call(MkString("privatePostTransfer"), MkMap(&VarMap{
		"cmd":  MkString("transfer/transferOutList"),
		"body": this.Extend(request, params),
	}))
	_ = response
	withdrawals := this.SafeValue(*(response).At(MkString("result")), MkString("items"), MkArray(&VarArray{}))
	_ = withdrawals
	for i := MkInteger(0); IsTrue(OpLw(i, withdrawals.Length)); OpIncr(&i) {
		*(*(withdrawals).At(i)).At(MkString("type")) = MkString("withdrawal")
	}
	return this.ParseTransactions(withdrawals, currency, since, limit)
}

func (this *Bibox) ParseTransaction(goArgs ...*Variant) *Variant {
	transaction := GoGetArg(goArgs, 0, MkUndefined())
	_ = transaction
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	id := this.SafeString(transaction, MkString("id"))
	_ = id
	address := this.SafeString(transaction, MkString("to_address"))
	_ = address
	currencyId := this.SafeString(transaction, MkString("coin_symbol"))
	_ = currencyId
	code := this.SafeCurrencyCode(currencyId, currency)
	_ = code
	timestamp := this.SafeString(transaction, MkString("createdAt"))
	_ = timestamp
	tag := this.SafeString(transaction, MkString("addr_remark"))
	_ = tag
	type_ := this.SafeString(transaction, MkString("type"))
	_ = type_
	status := this.ParseTransactionStatusByType(this.SafeString(transaction, MkString("status")), type_)
	_ = status
	amount := this.SafeNumber(transaction, MkString("amount"))
	_ = amount
	feeCost := this.SafeNumber(transaction, MkString("fee"))
	_ = feeCost
	if IsTrue(OpEq2(type_, MkString("deposit"))) {
		feeCost = MkInteger(0)
		tag = OpCopy(MkUndefined())
	}
	fee := MkMap(&VarMap{
		"cost":     feeCost,
		"currency": code,
	})
	_ = fee
	return MkMap(&VarMap{
		"info":      transaction,
		"id":        id,
		"txid":      MkUndefined(),
		"timestamp": timestamp,
		"datetime":  this.Iso8601(timestamp),
		"address":   address,
		"tag":       tag,
		"type":      type_,
		"amount":    amount,
		"currency":  code,
		"status":    status,
		"updated":   MkUndefined(),
		"fee":       fee,
	})
}

func (this *Bibox) ParseTransactionStatusByType(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	type_ := GoGetArg(goArgs, 1, MkUndefined())
	_ = type_
	statuses := MkMap(&VarMap{
		"deposit": MkMap(&VarMap{
			"1": MkString("pending"),
			"2": MkString("ok"),
		}),
		"withdrawal": MkMap(&VarMap{
			"0": MkString("pending"),
			"3": MkString("ok"),
		}),
	})
	_ = statuses
	return this.SafeString(this.SafeValue(statuses, type_, MkMap(&VarMap{})), status, status)
}

func (this *Bibox) CreateOrder(goArgs ...*Variant) *Variant {
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
	orderType := OpTrinary(OpEq2(type_, MkString("limit")), MkInteger(2), MkInteger(1))
	_ = orderType
	orderSide := OpTrinary(OpEq2(side, MkString("buy")), MkInteger(1), MkInteger(2))
	_ = orderSide
	request := MkMap(&VarMap{
		"cmd": MkString("orderpending/trade"),
		"body": this.Extend(MkMap(&VarMap{
			"pair":         *(market).At(MkString("id")),
			"account_type": MkInteger(0),
			"order_type":   orderType,
			"order_side":   orderSide,
			"pay_bix":      MkInteger(0),
			"amount":       amount,
			"price":        price,
		}), params),
	})
	_ = request
	response := this.Call(MkString("privatePostOrderpending"), request)
	_ = response
	return MkMap(&VarMap{
		"info": response,
		"id":   this.SafeString(response, MkString("result")),
	})
}

func (this *Bibox) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	request := MkMap(&VarMap{
		"cmd":  MkString("orderpending/cancelTrade"),
		"body": this.Extend(MkMap(&VarMap{"orders_id": id}), params),
	})
	_ = request
	response := this.Call(MkString("privatePostOrderpending"), request)
	_ = response
	return response
}

func (this *Bibox) FetchOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{
		"cmd": MkString("orderpending/order"),
		"body": this.Extend(MkMap(&VarMap{
			"id":           id.ToString(),
			"account_type": MkInteger(0),
		}), params),
	})
	_ = request
	response := this.Call(MkString("privatePostOrderpending"), request)
	_ = response
	order := this.SafeValue(response, MkString("result"))
	_ = order
	if IsTrue(this.IsEmpty(order)) {
		panic(NewOrderNotFound(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" order "), OpAdd(id, MkString(" not found"))))))
	}
	return this.ParseOrder(order)
}

func (this *Bibox) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(OpEq2(market, MkUndefined())) {
		marketId := OpCopy(MkUndefined())
		_ = marketId
		baseId := this.SafeString(order, MkString("coin_symbol"))
		_ = baseId
		quoteId := this.SafeString(order, MkString("currency_symbol"))
		_ = quoteId
		if IsTrue(OpAnd(OpNotEq2(baseId, MkUndefined()), OpNotEq2(quoteId, MkUndefined()))) {
			marketId = OpAdd(baseId, OpAdd(MkString("_"), quoteId))
		}
		if IsTrue(OpHasMember(marketId, *this.At(MkString("markets_by_id")))) {
			market = *(*this.At(MkString("markets_by_id"))).At(marketId)
		}
	}
	if IsTrue(OpNotEq2(market, MkUndefined())) {
		symbol = *(market).At(MkString("symbol"))
	}
	rawType := this.SafeString(order, MkString("order_type"))
	_ = rawType
	type_ := OpTrinary(OpEq2(rawType, MkString("1")), MkString("market"), MkString("limit"))
	_ = type_
	timestamp := this.SafeInteger(order, MkString("createdAt"))
	_ = timestamp
	price := this.SafeNumber(order, MkString("price"))
	_ = price
	average := this.SafeNumber(order, MkString("deal_price"))
	_ = average
	filled := this.SafeNumber(order, MkString("deal_amount"))
	_ = filled
	amount := this.SafeNumber(order, MkString("amount"))
	_ = amount
	cost := this.SafeNumber2(order, MkString("deal_money"), MkString("money"))
	_ = cost
	rawSide := this.SafeString(order, MkString("order_side"))
	_ = rawSide
	side := OpTrinary(OpEq2(rawSide, MkString("1")), MkString("buy"), MkString("sell"))
	_ = side
	status := this.ParseOrderStatus(this.SafeString(order, MkString("status")))
	_ = status
	id := this.SafeString(order, MkString("id"))
	_ = id
	feeCost := this.SafeNumber(order, MkString("fee"))
	_ = feeCost
	fee := OpCopy(MkUndefined())
	_ = fee
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		fee = MkMap(&VarMap{
			"cost":     feeCost,
			"currency": MkUndefined(),
		})
	}
	return this.SafeOrder(MkMap(&VarMap{
		"info":               order,
		"id":                 id,
		"clientOrderId":      MkUndefined(),
		"timestamp":          timestamp,
		"datetime":           this.Iso8601(timestamp),
		"lastTradeTimestamp": MkUndefined(),
		"symbol":             symbol,
		"type":               type_,
		"timeInForce":        MkUndefined(),
		"postOnly":           MkUndefined(),
		"side":               side,
		"price":              price,
		"stopPrice":          MkUndefined(),
		"amount":             amount,
		"cost":               cost,
		"average":            average,
		"filled":             filled,
		"remaining":          MkUndefined(),
		"status":             status,
		"fee":                fee,
		"trades":             MkUndefined(),
	}))
}

func (this *Bibox) ParseOrderStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"1": MkString("open"),
		"2": MkString("open"),
		"3": MkString("closed"),
		"4": MkString("canceled"),
		"5": MkString("canceled"),
		"6": MkString("canceled"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Bibox) FetchOpenOrders(goArgs ...*Variant) *Variant {
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
	pair := OpCopy(MkUndefined())
	_ = pair
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
		pair = *(market).At(MkString("id"))
	}
	size := OpTrinary(limit, limit, MkInteger(200))
	_ = size
	request := MkMap(&VarMap{
		"cmd": MkString("orderpending/orderPendingList"),
		"body": this.Extend(MkMap(&VarMap{
			"pair":         pair,
			"account_type": MkInteger(0),
			"page":         MkInteger(1),
			"size":         size,
		}), params),
	})
	_ = request
	response := this.Call(MkString("privatePostOrderpending"), request)
	_ = response
	orders := this.SafeValue(*(response).At(MkString("result")), MkString("items"), MkArray(&VarArray{}))
	_ = orders
	return this.ParseOrders(orders, market, since, limit)
}

func (this *Bibox) FetchClosedOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkInteger(200))
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchClosedOrders() requires a `symbol` argument"))))
	}
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{
		"cmd": MkString("orderpending/pendingHistoryList"),
		"body": this.Extend(MkMap(&VarMap{
			"pair":         *(market).At(MkString("id")),
			"account_type": MkInteger(0),
			"page":         MkInteger(1),
			"size":         limit,
		}), params),
	})
	_ = request
	response := this.Call(MkString("privatePostOrderpending"), request)
	_ = response
	orders := this.SafeValue(*(response).At(MkString("result")), MkString("items"), MkArray(&VarArray{}))
	_ = orders
	return this.ParseOrders(orders, market, since, limit)
}

func (this *Bibox) FetchMyTrades(goArgs ...*Variant) *Variant {
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
	size := OpTrinary(limit, limit, MkInteger(200))
	_ = size
	request := MkMap(&VarMap{
		"cmd": MkString("orderpending/orderHistoryList"),
		"body": this.Extend(MkMap(&VarMap{
			"pair":            *(market).At(MkString("id")),
			"account_type":    MkInteger(0),
			"page":            MkInteger(1),
			"size":            size,
			"coin_symbol":     *(market).At(MkString("baseId")),
			"currency_symbol": *(market).At(MkString("quoteId")),
		}), params),
	})
	_ = request
	response := this.Call(MkString("privatePostOrderpending"), request)
	_ = response
	trades := this.SafeValue(*(response).At(MkString("result")), MkString("items"), MkArray(&VarArray{}))
	_ = trades
	return this.ParseTrades(trades, market, since, limit)
}

func (this *Bibox) FetchDepositAddress(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{
		"cmd":  MkString("transfer/transferIn"),
		"body": this.Extend(MkMap(&VarMap{"coin_symbol": *(currency).At(MkString("id"))}), params),
	})
	_ = request
	response := this.Call(MkString("privatePostTransfer"), request)
	_ = response
	result := this.SafeString(response, MkString("result"))
	_ = result
	address := OpCopy(result)
	_ = address
	tag := OpCopy(MkUndefined())
	_ = tag
	if IsTrue(this.IsJsonEncodedObject(result)) {
		parsed := JSON.Parse(result)
		_ = parsed
		address = this.SafeString(parsed, MkString("account"))
		tag = this.SafeString(parsed, MkString("memo"))
	}
	return MkMap(&VarMap{
		"currency": code,
		"address":  address,
		"tag":      tag,
		"info":     response,
	})
}

func (this *Bibox) Withdraw(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpEq2(*this.At(MkString("password")), MkUndefined())) {
		if IsTrue(OpNot(OpHasMember(MkString("trade_pwd"), params))) {
			panic(NewExchangeError(OpAdd(*this.At(MkString("id")), MkString(" withdraw() requires this.password set on the exchange instance or a trade_pwd parameter"))))
		}
	}
	if IsTrue(OpNot(OpHasMember(MkString("totp_code"), params))) {
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), MkString(" withdraw() requires a totp_code parameter for 2FA authentication"))))
	}
	request := MkMap(&VarMap{
		"trade_pwd":   *this.At(MkString("password")),
		"coin_symbol": *(currency).At(MkString("id")),
		"amount":      amount,
		"addr":        address,
	})
	_ = request
	if IsTrue(OpNotEq2(tag, MkUndefined())) {
		*(request).At(MkString("address_remark")) = OpCopy(tag)
	}
	response := this.Call(MkString("privatePostTransfer"), MkMap(&VarMap{
		"cmd":  MkString("transfer/transferOut"),
		"body": this.Extend(request, params),
	}))
	_ = response
	return MkMap(&VarMap{
		"info": response,
		"id":   MkUndefined(),
	})
}

func (this *Bibox) FetchFundingFees(goArgs ...*Variant) *Variant {
	codes := GoGetArg(goArgs, 0, MkUndefined())
	_ = codes
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	withdrawFees := MkMap(&VarMap{})
	_ = withdrawFees
	info := MkMap(&VarMap{})
	_ = info
	if IsTrue(OpEq2(codes, MkUndefined())) {
		codes = GoGetKeys(*this.At(MkString("currencies")))
	}
	for i := MkInteger(0); IsTrue(OpLw(i, codes.Length)); OpIncr(&i) {
		code := *(codes).At(i)
		_ = code
		currency := this.Currency(code)
		_ = currency
		request := MkMap(&VarMap{
			"cmd":  MkString("transfer/coinConfig"),
			"body": this.Extend(MkMap(&VarMap{"coin_symbol": *(currency).At(MkString("id"))}), params),
		})
		_ = request
		response := this.Call(MkString("privatePostTransfer"), request)
		_ = response
		*(info).At(code) = OpCopy(response)
		*(withdrawFees).At(code) = this.SafeNumber(*(response).At(MkString("result")), MkString("withdraw_fee"))
	}
	return MkMap(&VarMap{
		"info":     info,
		"withdraw": withdrawFees,
		"deposit":  MkMap(&VarMap{}),
	})
}

func (this *Bibox) Sign(goArgs ...*Variant) *Variant {
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
	url := OpAdd(this.ImplodeHostname(*(*this.At(MkString("urls"))).At(MkString("api"))), OpAdd(MkString("/"), OpAdd(*this.At(MkString("version")), OpAdd(MkString("/"), path))))
	_ = url
	cmds := this.Json(MkArray(&VarArray{
		params,
	}))
	_ = cmds
	if IsTrue(OpEq2(api, MkString("public"))) {
		if IsTrue(OpNotEq2(method, MkString("GET"))) {
			body = MkMap(&VarMap{"cmds": cmds})
		} else {
			if IsTrue(*(GoGetKeys(params)).At(MkString("length"))) {
				OpAddAssign(&url, OpAdd(MkString("?"), this.Urlencode(params)))
			}
		}
	} else {
		if IsTrue(OpEq2(api, MkString("v2private"))) {
			this.CheckRequiredCredentials()
			url = OpAdd(this.ImplodeHostname(*(*this.At(MkString("urls"))).At(MkString("api"))), OpAdd(MkString("/v2/"), path))
			json_params := this.Json(params)
			_ = json_params
			body = MkMap(&VarMap{
				"body":   json_params,
				"apikey": *this.At(MkString("apiKey")),
				"sign":   this.Hmac(this.Encode(json_params), this.Encode(*this.At(MkString("secret"))), MkString("md5")),
			})
		} else {
			this.CheckRequiredCredentials()
			body = MkMap(&VarMap{
				"cmds":   cmds,
				"apikey": *this.At(MkString("apiKey")),
				"sign":   this.Hmac(this.Encode(cmds), this.Encode(*this.At(MkString("secret"))), MkString("md5")),
			})
		}
	}
	if IsTrue(OpNotEq2(body, MkUndefined())) {
		body = this.Json(body, MkMap(&VarMap{"convertArraysToObjects": MkBool(true)}))
	}
	headers = MkMap(&VarMap{"Content-Type": MkString("application/json")})
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Bibox) HandleErrors(goArgs ...*Variant) *Variant {
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
		if IsTrue(OpHasMember(MkString("code"), *(response).At(MkString("error")))) {
			code := this.SafeString(*(response).At(MkString("error")), MkString("code"))
			_ = code
			feedback := OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))
			_ = feedback
			this.ThrowExactlyMatchedException(*this.At(MkString("exceptions")), code, feedback)
			panic(NewExchangeError(feedback))
		}
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))))
	}
	if IsTrue(OpNot(OpHasMember(MkString("result"), response))) {
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))))
	}
	return MkUndefined()
}

func (this *Bibox) Request(goArgs ...*Variant) *Variant {
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
	config := GoGetArg(goArgs, 6, MkMap(&VarMap{}))
	_ = config
	context := GoGetArg(goArgs, 7, MkMap(&VarMap{}))
	_ = context
	response := this.Fetch2(path, api, method, params, headers, body, config, context)
	_ = response
	if IsTrue(OpEq2(method, MkString("GET"))) {
		return response
	} else {
		return *(*(response).At(MkString("result"))).At(MkInteger(0))
	}
	return MkUndefined()
}
