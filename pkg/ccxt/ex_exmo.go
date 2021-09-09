package ccxt

import ()

type Exmo struct {
	*ExchangeBase
}

var _ Exchange = (*Exmo)(nil)

func init() {
	exchange := &Exmo{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Exmo) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("exmo"),
		"name": MkString("EXMO"),
		"countries": MkArray(&VarArray{
			MkString("ES"),
			MkString("RU"),
		}),
		"rateLimit": MkInteger(350),
		"version":   MkString("v1.1"),
		"has": MkMap(&VarMap{
			"cancelOrder":         MkBool(true),
			"CORS":                MkBool(false),
			"createOrder":         MkBool(true),
			"fetchBalance":        MkBool(true),
			"fetchCurrencies":     MkBool(true),
			"fetchDepositAddress": MkBool(true),
			"fetchFundingFees":    MkBool(true),
			"fetchMarkets":        MkBool(true),
			"fetchMyTrades":       MkBool(true),
			"fetchOHLCV":          MkBool(true),
			"fetchOpenOrders":     MkBool(true),
			"fetchOrder":          MkString("emulated"),
			"fetchOrderBook":      MkBool(true),
			"fetchOrderBooks":     MkBool(true),
			"fetchOrderTrades":    MkBool(true),
			"fetchTicker":         MkBool(true),
			"fetchTickers":        MkBool(true),
			"fetchTrades":         MkBool(true),
			"fetchTradingFee":     MkBool(true),
			"fetchTradingFees":    MkBool(true),
			"fetchTransactions":   MkBool(true),
			"fetchWithdrawals":    MkBool(true),
			"withdraw":            MkBool(true),
		}),
		"timeframes": MkMap(&VarMap{
			"1m":  MkString("1"),
			"5m":  MkString("5"),
			"15m": MkString("15"),
			"30m": MkString("30"),
			"45m": MkString("45"),
			"1h":  MkString("60"),
			"2h":  MkString("120"),
			"3h":  MkString("180"),
			"4h":  MkString("240"),
			"1d":  MkString("D"),
			"1w":  MkString("W"),
			"1M":  MkString("M"),
		}),
		"urls": MkMap(&VarMap{
			"logo": MkString("https://user-images.githubusercontent.com/1294454/27766491-1b0ea956-5eda-11e7-9225-40d67b481b8d.jpg"),
			"api": MkMap(&VarMap{
				"public":  MkString("https://api.exmo.com"),
				"private": MkString("https://api.exmo.com"),
				"web":     MkString("https://exmo.me"),
			}),
			"www":      MkString("https://exmo.me"),
			"referral": MkString("https://exmo.me/?ref=131685"),
			"doc": MkArray(&VarArray{
				MkString("https://exmo.me/en/api_doc?ref=131685"),
				MkString("https://github.com/exmo-dev/exmo_api_lib/tree/master/nodejs"),
			}),
			"fees": MkString("https://exmo.com/en/docs/fees"),
		}),
		"api": MkMap(&VarMap{
			"web": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("ctrl/feesAndLimits"),
				MkString("en/docs/fees"),
			})}),
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("currency"),
				MkString("currency/list/extended"),
				MkString("order_book"),
				MkString("pair_settings"),
				MkString("ticker"),
				MkString("trades"),
				MkString("candles_history"),
				MkString("required_amount"),
				MkString("payments/providers/crypto/list"),
			})}),
			"private": MkMap(&VarMap{"post": MkArray(&VarArray{
				MkString("user_info"),
				MkString("order_create"),
				MkString("order_cancel"),
				MkString("stop_market_order_create"),
				MkString("stop_market_order_cancel"),
				MkString("user_open_orders"),
				MkString("user_trades"),
				MkString("user_cancelled_orders"),
				MkString("order_trades"),
				MkString("deposit_address"),
				MkString("withdraw_crypt"),
				MkString("withdraw_get_txid"),
				MkString("excode_create"),
				MkString("excode_load"),
				MkString("code_check"),
				MkString("wallet_history"),
				MkString("wallet_operations"),
				MkString("margin/user/order/create"),
				MkString("margin/user/order/update"),
				MkString("margin/user/order/cancel"),
				MkString("margin/user/position/close"),
				MkString("margin/user/position/margin_add"),
				MkString("margin/user/position/margin_remove"),
				MkString("margin/currency/list"),
				MkString("margin/pair/list"),
				MkString("margin/settings"),
				MkString("margin/funding/list"),
				MkString("margin/user/info"),
				MkString("margin/user/order/list"),
				MkString("margin/user/order/history"),
				MkString("margin/user/order/trades"),
				MkString("margin/user/order/max_quantity"),
				MkString("margin/user/position/list"),
				MkString("margin/user/position/margin_remove_info"),
				MkString("margin/user/position/margin_add_info"),
				MkString("margin/user/wallet/list"),
				MkString("margin/user/wallet/history"),
				MkString("margin/user/trade/list"),
				MkString("margin/trades"),
				MkString("margin/liquidation/feed"),
			})}),
		}),
		"fees": MkMap(&VarMap{
			"trading": MkMap(&VarMap{
				"feeSide":    MkString("get"),
				"tierBased":  MkBool(false),
				"percentage": MkBool(true),
				"maker":      this.ParseNumber(MkString("0.002")),
				"taker":      this.ParseNumber(MkString("0.002")),
			}),
			"funding": MkMap(&VarMap{
				"tierBased":  MkBool(false),
				"percentage": MkBool(false),
			}),
		}),
		"exceptions": MkMap(&VarMap{
			"exact": MkMap(&VarMap{
				"40005": AuthenticationError,
				"40009": InvalidNonce,
				"40015": ExchangeError,
				"40016": OnMaintenance,
				"40017": AuthenticationError,
				"40032": PermissionDenied,
				"40033": PermissionDenied,
				"40034": RateLimitExceeded,
				"50052": InsufficientFunds,
				"50054": InsufficientFunds,
				"50304": OrderNotFound,
				"50173": OrderNotFound,
				"50277": InvalidOrder,
				"50319": InvalidOrder,
				"50321": InvalidOrder,
				"50381": InvalidOrder,
			}),
			"broad": MkMap(&VarMap{
				"range period is too long": BadRequest,
				"invalid syntax":           BadRequest,
				"API rate limit exceeded":  RateLimitExceeded,
			}),
		}),
	}))
}

func (this *Exmo) FetchTradingFees(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	if IsTrue(*(*this.At(MkString("options"))).At(MkString("useWebapiForFetchingFees"))) {
		response := this.Call(MkString("webGetEnDocsFees"), params)
		_ = response
		parts := response.Split(MkString("<td class=\"th_fees_2\" colspan=\"2\">"))
		_ = parts
		numParts := OpCopy(parts.Length)
		_ = numParts
		if IsTrue(OpNotEq2(numParts, MkInteger(2))) {
			panic(NewNotSupported(OpAdd(*this.At(MkString("id")), MkString(" fetchTradingFees format has changed"))))
		}
		rest := *(parts).At(MkInteger(1))
		_ = rest
		parts = rest.Split(MkString("</td>"))
		numParts = OpCopy(parts.Length)
		if IsTrue(OpLw(numParts, MkInteger(2))) {
			panic(NewNotSupported(OpAdd(*this.At(MkString("id")), MkString(" fetchTradingFees format has changed"))))
		}
		fee := OpMulti(parseFloat((*(parts).At(MkInteger(0))).Call(MkString("replace"), MkString("%"), MkString(""))), MkNumber(0.01))
		_ = fee
		taker := OpCopy(fee)
		_ = taker
		maker := OpCopy(fee)
		_ = maker
		return MkMap(&VarMap{
			"maker": maker,
			"taker": taker,
		})
	} else {
		return MkMap(&VarMap{
			"maker": *(*(*this.At(MkString("fees"))).At(MkString("trading"))).At(MkString("maker")),
			"taker": *(*(*this.At(MkString("fees"))).At(MkString("trading"))).At(MkString("taker")),
		})
	}
	return MkUndefined()
}

func (this *Exmo) ParseFixedFloatValue(goArgs ...*Variant) *Variant {
	input := GoGetArg(goArgs, 0, MkUndefined())
	_ = input
	if IsTrue(OpOr(OpEq2(input, MkUndefined()), OpEq2(input, MkString("-")))) {
		return MkUndefined()
	}
	if IsTrue(OpEq2(input, MkString(""))) {
		return MkInteger(0)
	}
	isPercentage := OpGtEq(input.IndexOf(MkString("%")), MkInteger(0))
	_ = isPercentage
	parts := input.Split(MkString(" "))
	_ = parts
	value := (*(parts).At(MkInteger(0))).Call(MkString("replace"), MkString("%"), MkString(""))
	_ = value
	result := parseFloat(value)
	_ = result
	if IsTrue(OpAnd(OpGt(result, MkInteger(0)), isPercentage)) {
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" parseFixedFloatValue detected an unsupported non-zero percentage-based fee "), input))))
	}
	return result
}

func (this *Exmo) FetchFundingFees(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	currencyList := this.Call(MkString("publicGetCurrencyListExtended"), params)
	_ = currencyList
	cryptoList := this.Call(MkString("publicGetPaymentsProvidersCryptoList"), params)
	_ = cryptoList
	result := MkMap(&VarMap{
		"info":     cryptoList,
		"withdraw": MkMap(&VarMap{}),
		"deposit":  MkMap(&VarMap{}),
	})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, currencyList.Length)); OpIncr(&i) {
		currency := *(currencyList).At(i)
		_ = currency
		currencyId := this.SafeString(currency, MkString("name"))
		_ = currencyId
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		providers := this.SafeValue(cryptoList, currencyId, MkArray(&VarArray{}))
		_ = providers
		for j := MkInteger(0); IsTrue(OpLw(j, providers.Length)); OpIncr(&j) {
			provider := *(providers).At(j)
			_ = provider
			type_ := this.SafeString(provider, MkString("type"))
			_ = type_
			commissionDesc := this.SafeString(provider, MkString("commission_desc"))
			_ = commissionDesc
			newFee := this.ParseFixedFloatValue(commissionDesc)
			_ = newFee
			previousFee := this.SafeNumber(*(result).At(type_), code)
			_ = previousFee
			if IsTrue(OpOr(OpEq2(previousFee, MkUndefined()), OpAnd(OpNotEq2(newFee, MkUndefined()), OpLw(newFee, previousFee)))) {
				*(*(result).At(type_)).At(code) = OpCopy(newFee)
			}
		}
	}
	*(*this.At(MkString("options"))).At(MkString("fundingFees")) = OpCopy(result)
	return result
}

func (this *Exmo) FetchCurrencies(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	currencyList := this.Call(MkString("publicGetCurrencyListExtended"), params)
	_ = currencyList
	cryptoList := this.Call(MkString("publicGetPaymentsProvidersCryptoList"), params)
	_ = cryptoList
	result := MkMap(&VarMap{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, currencyList.Length)); OpIncr(&i) {
		currency := *(currencyList).At(i)
		_ = currency
		currencyId := this.SafeString(currency, MkString("name"))
		_ = currencyId
		name := this.SafeString(currency, MkString("description"))
		_ = name
		providers := this.SafeValue(cryptoList, currencyId)
		_ = providers
		active := OpCopy(MkBool(false))
		_ = active
		type_ := MkString("crypto")
		_ = type_
		limits := MkMap(&VarMap{
			"deposit": MkMap(&VarMap{
				"min": MkUndefined(),
				"max": MkUndefined(),
			}),
			"withdraw": MkMap(&VarMap{
				"min": MkUndefined(),
				"max": MkUndefined(),
			}),
		})
		_ = limits
		fee := OpCopy(MkUndefined())
		_ = fee
		if IsTrue(OpEq2(providers, MkUndefined())) {
			active = OpCopy(MkBool(true))
			type_ = MkString("fiat")
		} else {
			for j := MkInteger(0); IsTrue(OpLw(j, providers.Length)); OpIncr(&j) {
				provider := *(providers).At(j)
				_ = provider
				type_ := this.SafeString(provider, MkString("type"))
				_ = type_
				minValue := this.SafeNumber(provider, MkString("min"))
				_ = minValue
				maxValue := this.SafeNumber(provider, MkString("max"))
				_ = maxValue
				if IsTrue(OpEq2(maxValue, MkNumber(0.0))) {
					maxValue = OpCopy(MkUndefined())
				}
				activeProvider := this.SafeValue(provider, MkString("enabled"))
				_ = activeProvider
				if IsTrue(activeProvider) {
					active = OpCopy(MkBool(true))
					if IsTrue(OpOr(OpEq2(*(*(limits).At(type_)).At(MkString("min")), MkUndefined()), OpLw(minValue, *(*(limits).At(type_)).At(MkString("min"))))) {
						*(*(limits).At(type_)).At(MkString("min")) = OpCopy(minValue)
						*(*(limits).At(type_)).At(MkString("max")) = OpCopy(maxValue)
						if IsTrue(OpEq2(type_, MkString("withdraw"))) {
							commissionDesc := this.SafeString(provider, MkString("commission_desc"))
							_ = commissionDesc
							fee = this.ParseFixedFloatValue(commissionDesc)
						}
					}
				}
			}
		}
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		*(result).At(code) = MkMap(&VarMap{
			"id":        currencyId,
			"code":      code,
			"name":      name,
			"type":      type_,
			"active":    active,
			"fee":       fee,
			"precision": MkInteger(8),
			"limits":    limits,
			"info":      providers,
		})
	}
	return result
}

func (this *Exmo) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetPairSettings"), params)
	_ = response
	keys := GoGetKeys(response)
	_ = keys
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, keys.Length)); OpIncr(&i) {
		id := *(keys).At(i)
		_ = id
		market := *(response).At(id)
		_ = market
		symbol := id.Replace(MkString("_"), MkString("/"))
		_ = symbol
		baseId := MkUndefined()
		_ = baseId
		quoteId := MkUndefined()
		_ = quoteId
		ArrayUnpack(symbol.Split(MkString("/")), &baseId, &quoteId)
		base := this.SafeCurrencyCode(baseId)
		_ = base
		quote := this.SafeCurrencyCode(quoteId)
		_ = quote
		takerString := this.SafeString(market, MkString("commission_taker_percent"))
		_ = takerString
		makerString := this.SafeString(market, MkString("commission_maker_percent"))
		_ = makerString
		taker := this.ParseNumber(Precise.StringDiv(takerString, MkString("100")))
		_ = taker
		maker := this.ParseNumber(Precise.StringDiv(makerString, MkString("100")))
		_ = maker
		result.Push(MkMap(&VarMap{
			"id":      id,
			"symbol":  symbol,
			"base":    base,
			"quote":   quote,
			"baseId":  baseId,
			"quoteId": quoteId,
			"active":  MkBool(true),
			"taker":   taker,
			"maker":   maker,
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
					"min": this.SafeNumber(market, MkString("min_amount")),
					"max": this.SafeNumber(market, MkString("max_amount")),
				}),
			}),
			"precision": MkMap(&VarMap{
				"amount": MkInteger(8),
				"price":  this.SafeInteger(market, MkString("price_precision")),
			}),
			"info": market,
		}))
	}
	return result
}

func (this *Exmo) FetchOHLCV(goArgs ...*Variant) *Variant {
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
		"symbol":     *(market).At(MkString("id")),
		"resolution": *(*this.At(MkString("timeframes"))).At(timeframe),
	})
	_ = request
	options := this.SafeValue(*this.At(MkString("options")), MkString("fetchOHLCV"))
	_ = options
	maxLimit := this.SafeInteger(options, MkString("maxLimit"), MkInteger(3000))
	_ = maxLimit
	duration := this.ParseTimeframe(timeframe)
	_ = duration
	now := this.Milliseconds()
	_ = now
	if IsTrue(OpEq2(since, MkUndefined())) {
		if IsTrue(OpEq2(limit, MkUndefined())) {
			panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchOHLCV() requires a since argument or a limit argument"))))
		} else {
			if IsTrue(OpGt(limit, maxLimit)) {
				panic(NewBadRequest(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" fetchOHLCV will serve "), OpAdd(maxLimit.ToString(), MkString(" candles at most"))))))
			}
			*(request).At(MkString("from")) = OpSub(parseInt(OpDiv(now, MkInteger(1000))), OpSub(OpMulti(limit, duration), MkInteger(1)))
			*(request).At(MkString("to")) = parseInt(OpDiv(now, MkInteger(1000)))
		}
	} else {
		*(request).At(MkString("from")) = OpSub(parseInt(OpDiv(since, MkInteger(1000))), MkInteger(1))
		if IsTrue(OpEq2(limit, MkUndefined())) {
			*(request).At(MkString("to")) = parseInt(OpDiv(now, MkInteger(1000)))
		} else {
			if IsTrue(OpGt(limit, maxLimit)) {
				panic(NewBadRequest(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" fetchOHLCV will serve "), OpAdd(maxLimit.ToString(), MkString(" candles at most"))))))
			}
			to := this.Sum(since, OpMulti(limit, OpMulti(duration, MkInteger(1000))))
			_ = to
			*(request).At(MkString("to")) = parseInt(OpDiv(to, MkInteger(1000)))
		}
	}
	response := this.Call(MkString("publicGetCandlesHistory"), this.Extend(request, params))
	_ = response
	candles := this.SafeValue(response, MkString("candles"), MkArray(&VarArray{}))
	_ = candles
	return this.ParseOHLCVs(candles, market, timeframe, since, limit)
}

func (this *Exmo) ParseOHLCV(goArgs ...*Variant) *Variant {
	ohlcv := GoGetArg(goArgs, 0, MkUndefined())
	_ = ohlcv
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	return MkArray(&VarArray{
		this.SafeInteger(ohlcv, MkString("t")),
		this.SafeNumber(ohlcv, MkString("o")),
		this.SafeNumber(ohlcv, MkString("h")),
		this.SafeNumber(ohlcv, MkString("l")),
		this.SafeNumber(ohlcv, MkString("c")),
		this.SafeNumber(ohlcv, MkString("v")),
	})
}

func (this *Exmo) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privatePostUserInfo"), params)
	_ = response
	result := MkMap(&VarMap{"info": response})
	_ = result
	free := this.SafeValue(response, MkString("balances"), MkMap(&VarMap{}))
	_ = free
	used := this.SafeValue(response, MkString("reserved"), MkMap(&VarMap{}))
	_ = used
	currencyIds := GoGetKeys(free)
	_ = currencyIds
	for i := MkInteger(0); IsTrue(OpLw(i, currencyIds.Length)); OpIncr(&i) {
		currencyId := *(currencyIds).At(i)
		_ = currencyId
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		account := this.Account()
		_ = account
		if IsTrue(OpHasMember(currencyId, free)) {
			*(account).At(MkString("free")) = this.SafeString(free, currencyId)
		}
		if IsTrue(OpHasMember(currencyId, used)) {
			*(account).At(MkString("used")) = this.SafeString(used, currencyId)
		}
		*(result).At(code) = OpCopy(account)
	}
	return this.ParseBalance(result)
}

func (this *Exmo) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"pair": *(market).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("publicGetOrderBook"), this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, *(market).At(MkString("id")))
	_ = result
	return this.ParseOrderBook(result, symbol, MkUndefined(), MkString("bid"), MkString("ask"))
}

func (this *Exmo) FetchOrderBooks(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	ids := OpCopy(MkUndefined())
	_ = ids
	if IsTrue(OpEq2(symbols, MkUndefined())) {
		ids = (*((this).At(MkString("ids")))).Call(MkString("join"), MkString(","))
		if IsTrue(OpGt(ids.Length, MkInteger(2048))) {
			numIds := OpCopy((*((this).At(MkString("ids")))).Length)
			_ = numIds
			panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" has "), OpAdd(numIds.ToString(), MkString(" symbols exceeding max URL length, you are required to specify a list of symbols in the first argument to fetchOrderBooks"))))))
		}
	} else {
		ids = this.MarketIds(symbols)
		ids = ids.Join(MkString(","))
	}
	request := MkMap(&VarMap{"pair": ids})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("publicGetOrderBook"), this.Extend(request, params))
	_ = response
	result := MkMap(&VarMap{})
	_ = result
	marketIds := GoGetKeys(response)
	_ = marketIds
	for i := MkInteger(0); IsTrue(OpLw(i, marketIds.Length)); OpIncr(&i) {
		marketId := *(marketIds).At(i)
		_ = marketId
		symbol := OpCopy(marketId)
		_ = symbol
		if IsTrue(OpHasMember(marketId, *this.At(MkString("markets_by_id")))) {
			market := *(*this.At(MkString("markets_by_id"))).At(marketId)
			_ = market
			symbol = *(market).At(MkString("symbol"))
		}
		*(result).At(symbol) = this.ParseOrderBook(*(response).At(marketId), MkUndefined(), MkString("bid"), MkString("ask"))
	}
	return result
}

func (this *Exmo) ParseTicker(goArgs ...*Variant) *Variant {
	ticker := GoGetArg(goArgs, 0, MkUndefined())
	_ = ticker
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.SafeTimestamp(ticker, MkString("updated"))
	_ = timestamp
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(OpNotEq2(market, MkUndefined())) {
		symbol = *(market).At(MkString("symbol"))
	}
	last := this.SafeNumber(ticker, MkString("last_trade"))
	_ = last
	return MkMap(&VarMap{
		"symbol":        symbol,
		"timestamp":     timestamp,
		"datetime":      this.Iso8601(timestamp),
		"high":          this.SafeNumber(ticker, MkString("high")),
		"low":           this.SafeNumber(ticker, MkString("low")),
		"bid":           this.SafeNumber(ticker, MkString("buy_price")),
		"bidVolume":     MkUndefined(),
		"ask":           this.SafeNumber(ticker, MkString("sell_price")),
		"askVolume":     MkUndefined(),
		"vwap":          MkUndefined(),
		"open":          MkUndefined(),
		"close":         last,
		"last":          last,
		"previousClose": MkUndefined(),
		"change":        MkUndefined(),
		"percentage":    MkUndefined(),
		"average":       this.SafeNumber(ticker, MkString("avg")),
		"baseVolume":    this.SafeNumber(ticker, MkString("vol")),
		"quoteVolume":   this.SafeNumber(ticker, MkString("vol_curr")),
		"info":          ticker,
	})
}

func (this *Exmo) FetchTickers(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("publicGetTicker"), params)
	_ = response
	result := MkMap(&VarMap{})
	_ = result
	ids := GoGetKeys(response)
	_ = ids
	for i := MkInteger(0); IsTrue(OpLw(i, ids.Length)); OpIncr(&i) {
		id := *(ids).At(i)
		_ = id
		market := *(*this.At(MkString("markets_by_id"))).At(id)
		_ = market
		symbol := *(market).At(MkString("symbol"))
		_ = symbol
		ticker := *(response).At(id)
		_ = ticker
		*(result).At(symbol) = this.ParseTicker(ticker, market)
	}
	return this.FilterByArray(result, MkString("symbol"), symbols)
}

func (this *Exmo) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("publicGetTicker"), params)
	_ = response
	market := this.Market(symbol)
	_ = market
	return this.ParseTicker(*(response).At(*(market).At(MkString("id"))), market)
}

func (this *Exmo) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.SafeTimestamp(trade, MkString("date"))
	_ = timestamp
	symbol := OpCopy(MkUndefined())
	_ = symbol
	id := this.SafeString(trade, MkString("trade_id"))
	_ = id
	orderId := this.SafeString(trade, MkString("order_id"))
	_ = orderId
	price := this.SafeNumber(trade, MkString("price"))
	_ = price
	amount := this.SafeNumber(trade, MkString("quantity"))
	_ = amount
	cost := this.SafeNumber(trade, MkString("amount"))
	_ = cost
	side := this.SafeString(trade, MkString("type"))
	_ = side
	type_ := OpCopy(MkUndefined())
	_ = type_
	marketId := this.SafeString(trade, MkString("pair"))
	_ = marketId
	if IsTrue(OpNotEq2(marketId, MkUndefined())) {
		if IsTrue(OpHasMember(marketId, *this.At(MkString("markets_by_id")))) {
			market = *(*this.At(MkString("markets_by_id"))).At(marketId)
		} else {
			baseId := MkUndefined()
			_ = baseId
			quoteId := MkUndefined()
			_ = quoteId
			ArrayUnpack(marketId.Split(MkString("_")), &baseId, &quoteId)
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
	takerOrMaker := this.SafeString(trade, MkString("exec_type"))
	_ = takerOrMaker
	fee := OpCopy(MkUndefined())
	_ = fee
	feeCost := this.SafeNumber(trade, MkString("commission_amount"))
	_ = feeCost
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		feeCurrencyId := this.SafeString(trade, MkString("commission_currency"))
		_ = feeCurrencyId
		feeCurrencyCode := this.SafeCurrencyCode(feeCurrencyId)
		_ = feeCurrencyCode
		feeRate := this.SafeNumber(trade, MkString("commission_percent"))
		_ = feeRate
		if IsTrue(OpNotEq2(feeRate, MkUndefined())) {
			OpDivAssign(&feeRate, MkInteger(1000))
		}
		fee = MkMap(&VarMap{
			"cost":     feeCost,
			"currency": feeCurrencyCode,
			"rate":     feeRate,
		})
	}
	return MkMap(&VarMap{
		"id":           id,
		"info":         trade,
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       symbol,
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

func (this *Exmo) FetchTrades(goArgs ...*Variant) *Variant {
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
	response := this.Call(MkString("publicGetTrades"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, *(market).At(MkString("id")), MkArray(&VarArray{}))
	_ = data
	return this.ParseTrades(data, market, since, limit)
}

func (this *Exmo) FetchMyTrades(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchMyTrades() requires a symbol argument (a single symbol or an array)"))))
	}
	this.LoadMarkets()
	pair := OpCopy(MkUndefined())
	_ = pair
	market := OpCopy(MkUndefined())
	_ = market
	if IsTrue(GoIsArray(symbol)) {
		numSymbols := OpCopy(symbol.Length)
		_ = numSymbols
		if IsTrue(OpLw(numSymbols, MkInteger(1))) {
			panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchMyTrades() requires a non-empty symbol array"))))
		}
		marketIds := this.MarketIds(symbol)
		_ = marketIds
		pair = marketIds.Join(MkString(","))
	} else {
		market = this.Market(symbol)
		pair = *(market).At(MkString("id"))
	}
	request := MkMap(&VarMap{"pair": pair})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("privatePostUserTrades"), this.Extend(request, params))
	_ = response
	result := MkArray(&VarArray{})
	_ = result
	marketIds := GoGetKeys(response)
	_ = marketIds
	for i := MkInteger(0); IsTrue(OpLw(i, marketIds.Length)); OpIncr(&i) {
		marketId := *(marketIds).At(i)
		_ = marketId
		symbol := OpCopy(MkUndefined())
		_ = symbol
		if IsTrue(OpHasMember(marketId, *this.At(MkString("markets_by_id")))) {
			market = *(*this.At(MkString("markets_by_id"))).At(marketId)
			symbol = *(market).At(MkString("symbol"))
		} else {
			baseId := MkUndefined()
			_ = baseId
			quoteId := MkUndefined()
			_ = quoteId
			ArrayUnpack(marketId.Split(MkString("_")), &baseId, &quoteId)
			base := this.SafeCurrencyCode(baseId)
			_ = base
			quote := this.SafeCurrencyCode(quoteId)
			_ = quote
			symbol = OpAdd(base, OpAdd(MkString("/"), quote))
		}
		items := *(response).At(marketId)
		_ = items
		trades := this.ParseTrades(items, market, since, limit, MkMap(&VarMap{"symbol": symbol}))
		_ = trades
		result = this.ArrayConcat(result, trades)
	}
	return this.FilterBySinceLimit(result, since, limit)
}

func (this *Exmo) CreateOrder(goArgs ...*Variant) *Variant {
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
	prefix := OpTrinary(OpEq2(type_, MkString("market")), OpAdd(type_, MkString("_")), MkString(""))
	_ = prefix
	orderType := OpAdd(prefix, side)
	_ = orderType
	orderPrice := OpCopy(price)
	_ = orderPrice
	if IsTrue(OpAnd(OpEq2(type_, MkString("market")), OpEq2(price, MkUndefined()))) {
		orderPrice = MkInteger(0)
	}
	request := MkMap(&VarMap{
		"pair":     *(market).At(MkString("id")),
		"quantity": this.AmountToPrecision(symbol, amount),
		"type":     orderType,
		"price":    this.PriceToPrecision(symbol, orderPrice),
	})
	_ = request
	method := MkString("privatePostOrderCreate")
	_ = method
	clientOrderId := this.SafeValue2(params, MkString("client_id"), MkString("clientOrderId"))
	_ = clientOrderId
	if IsTrue(OpNotEq2(clientOrderId, MkUndefined())) {
		clientOrderId = this.SafeInteger2(params, MkString("client_id"), MkString("clientOrderId"))
		if IsTrue(OpEq2(clientOrderId, MkUndefined())) {
			panic(NewBadRequest(OpAdd(*this.At(MkString("id")), MkString(" createOrder client order id must be an integer / numeric literal"))))
		} else {
			*(request).At(MkString("client_id")) = OpCopy(clientOrderId)
		}
		params = this.Omit(params, MkArray(&VarArray{
			MkString("client_id"),
			MkString("clientOrderId"),
		}))
	}
	if IsTrue(OpOr(OpEq2(type_, MkString("stop")), OpOr(OpEq2(type_, MkString("stop_limit")), OpEq2(type_, MkString("trailing_stop"))))) {
		stopPrice := this.SafeNumber2(params, MkString("stop_price"), MkString("stopPrice"))
		_ = stopPrice
		if IsTrue(OpEq2(stopPrice, MkUndefined())) {
			panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" createOrder() requires a stopPrice extra param for a "), OpAdd(type_, MkString(" order"))))))
		} else {
			params = this.Omit(params, MkArray(&VarArray{
				MkString("stopPrice"),
				MkString("stop_price"),
			}))
			*(request).At(MkString("stop_price")) = this.PriceToPrecision(symbol, stopPrice)
			method = MkString("privatePostMarginUserOrderCreate")
		}
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	id := this.SafeString(response, MkString("order_id"))
	_ = id
	timestamp := this.Milliseconds()
	_ = timestamp
	status := MkString("open")
	_ = status
	return MkMap(&VarMap{
		"id":                 id,
		"info":               response,
		"timestamp":          timestamp,
		"datetime":           this.Iso8601(timestamp),
		"lastTradeTimestamp": MkUndefined(),
		"status":             status,
		"symbol":             symbol,
		"type":               type_,
		"side":               side,
		"price":              price,
		"cost":               MkUndefined(),
		"amount":             amount,
		"remaining":          amount,
		"filled":             MkNumber(0.0),
		"fee":                MkUndefined(),
		"trades":             MkUndefined(),
		"clientOrderId":      clientOrderId,
		"average":            MkUndefined(),
	})
}

func (this *Exmo) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"order_id": id})
	_ = request
	return this.Call(MkString("privatePostOrderCancel"), this.Extend(request, params))
}

func (this *Exmo) FetchOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"order_id": id.ToString()})
	_ = request
	response := this.Call(MkString("privatePostOrderTrades"), this.Extend(request, params))
	_ = response
	order := this.ParseOrder(response)
	_ = order
	return this.Extend(order, MkMap(&VarMap{"id": id.ToString()}))
}

func (this *Exmo) FetchOrderTrades(goArgs ...*Variant) *Variant {
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
	market := OpCopy(MkUndefined())
	_ = market
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
	}
	request := MkMap(&VarMap{"order_id": id.ToString()})
	_ = request
	response := this.Call(MkString("privatePostOrderTrades"), this.Extend(request, params))
	_ = response
	trades := this.SafeValue(response, MkString("trades"))
	_ = trades
	return this.ParseTrades(trades, market, since, limit)
}

func (this *Exmo) FetchOpenOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privatePostUserOpenOrders"), params)
	_ = response
	marketIds := GoGetKeys(response)
	_ = marketIds
	orders := MkArray(&VarArray{})
	_ = orders
	for i := MkInteger(0); IsTrue(OpLw(i, marketIds.Length)); OpIncr(&i) {
		marketId := *(marketIds).At(i)
		_ = marketId
		market := OpCopy(MkUndefined())
		_ = market
		if IsTrue(OpHasMember(marketId, *this.At(MkString("markets_by_id")))) {
			market = *(*this.At(MkString("markets_by_id"))).At(marketId)
		}
		parsedOrders := this.ParseOrders(*(response).At(marketId), market)
		_ = parsedOrders
		orders = this.ArrayConcat(orders, parsedOrders)
	}
	return this.FilterBySymbolSinceLimit(orders, symbol, since, limit)
}

func (this *Exmo) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	id := this.SafeString(order, MkString("order_id"))
	_ = id
	timestamp := this.SafeTimestamp(order, MkString("created"))
	_ = timestamp
	symbol := OpCopy(MkUndefined())
	_ = symbol
	side := this.SafeString(order, MkString("type"))
	_ = side
	if IsTrue(OpEq2(market, MkUndefined())) {
		marketId := OpCopy(MkUndefined())
		_ = marketId
		if IsTrue(OpHasMember(MkString("pair"), order)) {
			marketId = *(order).At(MkString("pair"))
		} else {
			if IsTrue(OpAnd(OpHasMember(MkString("in_currency"), order), OpHasMember(MkString("out_currency"), order))) {
				if IsTrue(OpEq2(side, MkString("buy"))) {
					marketId = OpAdd(*(order).At(MkString("in_currency")), OpAdd(MkString("_"), *(order).At(MkString("out_currency"))))
				} else {
					marketId = OpAdd(*(order).At(MkString("out_currency")), OpAdd(MkString("_"), *(order).At(MkString("in_currency"))))
				}
			}
		}
		if IsTrue(OpAnd(OpNotEq2(marketId, MkUndefined()), OpHasMember(marketId, *this.At(MkString("markets_by_id"))))) {
			market = *(*this.At(MkString("markets_by_id"))).At(marketId)
		}
	}
	amount := this.SafeNumber(order, MkString("quantity"))
	_ = amount
	if IsTrue(OpEq2(amount, MkUndefined())) {
		amountField := OpTrinary(OpEq2(side, MkString("buy")), MkString("in_amount"), MkString("out_amount"))
		_ = amountField
		amount = this.SafeNumber(order, amountField)
	}
	price := this.SafeNumber(order, MkString("price"))
	_ = price
	cost := this.SafeNumber(order, MkString("amount"))
	_ = cost
	filled := MkNumber(0.0)
	_ = filled
	trades := MkArray(&VarArray{})
	_ = trades
	transactions := this.SafeValue(order, MkString("trades"), MkArray(&VarArray{}))
	_ = transactions
	feeCost := OpCopy(MkUndefined())
	_ = feeCost
	lastTradeTimestamp := OpCopy(MkUndefined())
	_ = lastTradeTimestamp
	average := OpCopy(MkUndefined())
	_ = average
	numTransactions := OpCopy(transactions.Length)
	_ = numTransactions
	if IsTrue(OpGt(numTransactions, MkInteger(0))) {
		feeCost = MkInteger(0)
		for i := MkInteger(0); IsTrue(OpLw(i, numTransactions)); OpIncr(&i) {
			trade := this.ParseTrade(*(transactions).At(i), market)
			_ = trade
			if IsTrue(OpEq2(id, MkUndefined())) {
				id = *(trade).At(MkString("order"))
			}
			if IsTrue(OpEq2(timestamp, MkUndefined())) {
				timestamp = *(trade).At(MkString("timestamp"))
			}
			if IsTrue(OpGt(timestamp, *(trade).At(MkString("timestamp")))) {
				timestamp = *(trade).At(MkString("timestamp"))
			}
			filled = this.Sum(filled, *(trade).At(MkString("amount")))
			feeCost = this.Sum(feeCost, *(*(trade).At(MkString("fee"))).At(MkString("cost")))
			trades.Push(trade)
		}
		lastTradeTimestamp = *(*(trades).At(OpSub(numTransactions, MkInteger(1)))).At(MkString("timestamp"))
	}
	status := this.SafeString(order, MkString("status"))
	_ = status
	remaining := OpCopy(MkUndefined())
	_ = remaining
	if IsTrue(OpNotEq2(amount, MkUndefined())) {
		remaining = OpSub(amount, filled)
		if IsTrue(OpGtEq(filled, amount)) {
			status = MkString("closed")
		} else {
			status = MkString("open")
		}
	}
	if IsTrue(OpEq2(market, MkUndefined())) {
		market = this.GetMarketFromTrades(trades)
	}
	feeCurrency := OpCopy(MkUndefined())
	_ = feeCurrency
	if IsTrue(OpNotEq2(market, MkUndefined())) {
		symbol = *(market).At(MkString("symbol"))
		feeCurrency = *(market).At(MkString("quote"))
	}
	if IsTrue(OpEq2(cost, MkUndefined())) {
		if IsTrue(OpNotEq2(price, MkUndefined())) {
			cost = OpMulti(price, filled)
		}
	} else {
		if IsTrue(OpGt(filled, MkInteger(0))) {
			if IsTrue(OpEq2(average, MkUndefined())) {
				average = OpDiv(cost, filled)
			}
			if IsTrue(OpEq2(price, MkUndefined())) {
				price = OpDiv(cost, filled)
			}
		}
	}
	fee := MkMap(&VarMap{
		"cost":     feeCost,
		"currency": feeCurrency,
	})
	_ = fee
	clientOrderId := this.SafeInteger(order, MkString("client_id"))
	_ = clientOrderId
	return MkMap(&VarMap{
		"id":                 id,
		"clientOrderId":      clientOrderId,
		"datetime":           this.Iso8601(timestamp),
		"timestamp":          timestamp,
		"lastTradeTimestamp": lastTradeTimestamp,
		"status":             status,
		"symbol":             symbol,
		"type":               MkString("limit"),
		"timeInForce":        MkUndefined(),
		"postOnly":           MkUndefined(),
		"side":               side,
		"price":              price,
		"stopPrice":          MkUndefined(),
		"cost":               cost,
		"amount":             amount,
		"filled":             filled,
		"remaining":          remaining,
		"average":            average,
		"trades":             trades,
		"fee":                fee,
		"info":               order,
	})
}

func (this *Exmo) FetchDepositAddress(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privatePostDepositAddress"), params)
	_ = response
	depositAddress := this.SafeString(response, code)
	_ = depositAddress
	address := OpCopy(MkUndefined())
	_ = address
	tag := OpCopy(MkUndefined())
	_ = tag
	if IsTrue(depositAddress) {
		addressAndTag := depositAddress.Split(MkString(","))
		_ = addressAndTag
		address = *(addressAndTag).At(MkInteger(0))
		numParts := OpCopy(addressAndTag.Length)
		_ = numParts
		if IsTrue(OpGt(numParts, MkInteger(1))) {
			tag = *(addressAndTag).At(MkInteger(1))
		}
	}
	this.CheckAddress(address)
	return MkMap(&VarMap{
		"currency": code,
		"address":  address,
		"tag":      tag,
		"info":     response,
	})
}

func (this *Exmo) GetMarketFromTrades(goArgs ...*Variant) *Variant {
	trades := GoGetArg(goArgs, 0, MkUndefined())
	_ = trades
	tradesBySymbol := this.IndexBy(trades, MkString("pair"))
	_ = tradesBySymbol
	symbols := GoGetKeys(tradesBySymbol)
	_ = symbols
	numSymbols := OpCopy(symbols.Length)
	_ = numSymbols
	if IsTrue(OpEq2(numSymbols, MkInteger(1))) {
		return *(*this.At(MkString("markets"))).At(*(symbols).At(MkInteger(0)))
	}
	return MkUndefined()
}

func (this *Exmo) Withdraw(goArgs ...*Variant) *Variant {
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
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{
		"amount":   amount,
		"currency": *(currency).At(MkString("id")),
		"address":  address,
	})
	_ = request
	if IsTrue(OpNotEq2(tag, MkUndefined())) {
		*(request).At(MkString("invoice")) = OpCopy(tag)
	}
	response := this.Call(MkString("privatePostWithdrawCrypt"), this.Extend(request, params))
	_ = response
	return MkMap(&VarMap{
		"info": response,
		"id":   *(response).At(MkString("task_id")),
	})
}

func (this *Exmo) ParseTransactionStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"transferred": MkString("ok"),
		"paid":        MkString("ok"),
		"pending":     MkString("pending"),
		"processing":  MkString("pending"),
		"verifying":   MkString("pending"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Exmo) ParseTransaction(goArgs ...*Variant) *Variant {
	transaction := GoGetArg(goArgs, 0, MkUndefined())
	_ = transaction
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	id := this.SafeString(transaction, MkString("order_id"))
	_ = id
	timestamp := this.SafeTimestamp2(transaction, MkString("dt"), MkString("created"))
	_ = timestamp
	updated := this.SafeTimestamp(transaction, MkString("updated"))
	_ = updated
	amount := this.SafeNumber(transaction, MkString("amount"))
	_ = amount
	if IsTrue(OpNotEq2(amount, MkUndefined())) {
		amount = Math.Abs(amount)
	}
	status := this.ParseTransactionStatus(this.SafeStringLower(transaction, MkString("status")))
	_ = status
	txid := this.SafeString(transaction, MkString("txid"))
	_ = txid
	if IsTrue(OpEq2(txid, MkUndefined())) {
		extra := this.SafeValue(transaction, MkString("extra"), MkMap(&VarMap{}))
		_ = extra
		extraTxid := this.SafeString(extra, MkString("txid"))
		_ = extraTxid
		if IsTrue(OpNotEq2(extraTxid, MkString(""))) {
			txid = OpCopy(extraTxid)
		}
	}
	type_ := this.SafeString(transaction, MkString("type"))
	_ = type_
	currencyId := this.SafeString2(transaction, MkString("curr"), MkString("currency"))
	_ = currencyId
	code := this.SafeCurrencyCode(currencyId, currency)
	_ = code
	address := OpCopy(MkUndefined())
	_ = address
	tag := OpCopy(MkUndefined())
	_ = tag
	comment := OpCopy(MkUndefined())
	_ = comment
	account := this.SafeString(transaction, MkString("account"))
	_ = account
	if IsTrue(OpEq2(type_, MkString("deposit"))) {
		comment = OpCopy(account)
	} else {
		if IsTrue(OpEq2(type_, MkString("withdrawal"))) {
			address = OpCopy(account)
			if IsTrue(OpNotEq2(address, MkUndefined())) {
				parts := address.Split(MkString(":"))
				_ = parts
				numParts := OpCopy(parts.Length)
				_ = numParts
				if IsTrue(OpEq2(numParts, MkInteger(2))) {
					address = this.SafeString(parts, MkInteger(1))
					address = address.Replace(MkString(" "), MkString(""))
				}
			}
		}
	}
	fee := OpCopy(MkUndefined())
	_ = fee
	if IsTrue(OpNot(*(*(*this.At(MkString("fees"))).At(MkString("funding"))).At(MkString("percentage")))) {
		key := OpTrinary(OpEq2(type_, MkString("withdrawal")), MkString("withdraw"), MkString("deposit"))
		_ = key
		feeCost := this.SafeNumber(transaction, MkString("commission"))
		_ = feeCost
		if IsTrue(OpEq2(feeCost, MkUndefined())) {
			feeCost = this.SafeNumber(*(*(*this.At(MkString("options"))).At(MkString("fundingFees"))).At(key), code)
		}
		provider := this.SafeString(transaction, MkString("provider"))
		_ = provider
		if IsTrue(OpEq2(provider, MkString("cashback"))) {
			feeCost = MkInteger(0)
		}
		if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
			if IsTrue(OpEq2(type_, MkString("withdrawal"))) {
				amount = OpSub(amount, feeCost)
			}
			fee = MkMap(&VarMap{
				"cost":     feeCost,
				"currency": code,
				"rate":     MkUndefined(),
			})
		}
	}
	return MkMap(&VarMap{
		"info":        transaction,
		"id":          id,
		"timestamp":   timestamp,
		"datetime":    this.Iso8601(timestamp),
		"currency":    code,
		"amount":      amount,
		"address":     address,
		"addressTo":   address,
		"addressFrom": MkUndefined(),
		"tag":         tag,
		"tagTo":       tag,
		"tagFrom":     MkUndefined(),
		"status":      status,
		"type":        type_,
		"updated":     updated,
		"comment":     comment,
		"txid":        txid,
		"fee":         fee,
	})
}

func (this *Exmo) FetchTransactions(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("date")) = parseInt(OpDiv(since, MkInteger(1000)))
	}
	currency := OpCopy(MkUndefined())
	_ = currency
	if IsTrue(OpNotEq2(code, MkUndefined())) {
		currency = this.Currency(code)
	}
	response := this.Call(MkString("privatePostWalletHistory"), this.Extend(request, params))
	_ = response
	return this.ParseTransactions(*(response).At(MkString("history")), currency, since, limit)
}

func (this *Exmo) FetchWithdrawals(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"type": MkString("withdraw")})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	if IsTrue(OpNotEq2(code, MkUndefined())) {
		currency = this.Currency(code)
		*(request).At(MkString("currency")) = *(currency).At(MkString("id"))
	}
	response := this.Call(MkString("privatePostWalletOperations"), this.Extend(request, params))
	_ = response
	return this.ParseTransactions(*(response).At(MkString("items")), currency, since, limit)
}

func (this *Exmo) Sign(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpNotEq2(api, MkString("web"))) {
		OpAddAssign(&url, OpAdd(*this.At(MkString("version")), MkString("/")))
	}
	OpAddAssign(&url, path)
	if IsTrue(OpOr(OpEq2(api, MkString("public")), OpEq2(api, MkString("web")))) {
		if IsTrue(*(GoGetKeys(params)).At(MkString("length"))) {
			OpAddAssign(&url, OpAdd(MkString("?"), this.Urlencode(params)))
		}
	} else {
		if IsTrue(OpEq2(api, MkString("private"))) {
			this.CheckRequiredCredentials()
			nonce := this.Nonce()
			_ = nonce
			body = this.Urlencode(this.Extend(MkMap(&VarMap{"nonce": nonce}), params))
			headers = MkMap(&VarMap{
				"Content-Type": MkString("application/x-www-form-urlencoded"),
				"Key":          *this.At(MkString("apiKey")),
				"Sign":         this.Hmac(this.Encode(body), this.Encode(*this.At(MkString("secret"))), MkString("sha512")),
			})
		}
	}
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Exmo) Nonce(goArgs ...*Variant) *Variant {
	return this.Milliseconds()
}

func (this *Exmo) HandleErrors(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpOr(OpHasMember(MkString("result"), response), OpHasMember(MkString("errmsg"), response))) {
		success := this.SafeValue(response, MkString("result"), MkBool(false))
		_ = success
		if IsTrue(OpEq2(OpType(success), MkString("string"))) {
			if IsTrue(OpOr(OpEq2(success, MkString("true")), OpEq2(success, MkString("1")))) {
				success = OpCopy(MkBool(true))
			} else {
				success = OpCopy(MkBool(false))
			}
		}
		if IsTrue(OpNot(success)) {
			code := OpCopy(MkUndefined())
			_ = code
			message := this.SafeString2(response, MkString("error"), MkString("errmsg"))
			_ = message
			errorParts := message.Split(MkString(":"))
			_ = errorParts
			numParts := OpCopy(errorParts.Length)
			_ = numParts
			if IsTrue(OpGt(numParts, MkInteger(1))) {
				errorSubParts := (*(errorParts).At(MkInteger(0))).Call(MkString("split"), MkString(" "))
				_ = errorSubParts
				numSubParts := OpCopy(errorSubParts.Length)
				_ = numSubParts
				code = OpTrinary(OpGt(numSubParts, MkInteger(1)), *(errorSubParts).At(MkInteger(1)), *(errorSubParts).At(MkInteger(0)))
			}
			feedback := OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))
			_ = feedback
			this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("exact")), code, feedback)
			this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("broad")), message, feedback)
			panic(NewExchangeError(feedback))
		}
	}
	return MkUndefined()
}
