package ccxt

import ()

type Delta struct {
	*ExchangeBase
}

var _ Exchange = (*Delta)(nil)

func init() {
	exchange := &Delta{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Delta) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("delta"),
		"name": MkString("Delta Exchange"),
		"countries": MkArray(&VarArray{
			MkString("VC"),
		}),
		"rateLimit": MkInteger(300),
		"version":   MkString("v2"),
		"has": MkMap(&VarMap{
			"cancelAllOrders":     MkBool(true),
			"cancelOrder":         MkBool(true),
			"createOrder":         MkBool(true),
			"editOrder":           MkBool(true),
			"fetchBalance":        MkBool(true),
			"fetchClosedOrders":   MkBool(true),
			"fetchDepositAddress": MkBool(true),
			"fetchCurrencies":     MkBool(true),
			"fetchLedger":         MkBool(true),
			"fetchMarkets":        MkBool(true),
			"fetchMyTrades":       MkBool(true),
			"fetchOHLCV":          MkBool(true),
			"fetchOpenOrders":     MkBool(true),
			"fetchOrderBook":      MkBool(true),
			"fetchStatus":         MkBool(true),
			"fetchTicker":         MkBool(true),
			"fetchTickers":        MkBool(true),
			"fetchTime":           MkBool(true),
			"fetchTrades":         MkBool(true),
		}),
		"timeframes": MkMap(&VarMap{
			"1m":  MkString("1m"),
			"3m":  MkString("3m"),
			"5m":  MkString("5m"),
			"15m": MkString("15m"),
			"30m": MkString("30m"),
			"1h":  MkString("1h"),
			"2h":  MkString("2h"),
			"4h":  MkString("4h"),
			"6h":  MkString("6h"),
			"1d":  MkString("1d"),
			"7d":  MkString("7d"),
			"1w":  MkString("1w"),
			"2w":  MkString("2w"),
			"1M":  MkString("30d"),
		}),
		"urls": MkMap(&VarMap{
			"logo": MkString("https://user-images.githubusercontent.com/1294454/99450025-3be60a00-2931-11eb-9302-f4fd8d8589aa.jpg"),
			"test": MkMap(&VarMap{
				"public":  MkString("https://testnet-api.delta.exchange"),
				"private": MkString("https://testnet-api.delta.exchange"),
			}),
			"api": MkMap(&VarMap{
				"public":  MkString("https://api.delta.exchange"),
				"private": MkString("https://api.delta.exchange"),
			}),
			"www": MkString("https://www.delta.exchange"),
			"doc": MkArray(&VarArray{
				MkString("https://docs.delta.exchange"),
			}),
			"fees":     MkString("https://www.delta.exchange/fees"),
			"referral": MkString("https://www.delta.exchange/app/signup/?code=IULYNB"),
		}),
		"api": MkMap(&VarMap{
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("assets"),
				MkString("settings"),
				MkString("indices"),
				MkString("products"),
				MkString("tickers"),
				MkString("tickers/{symbol}"),
				MkString("l2orderbook/{symbol}"),
				MkString("trades/{symbol}"),
				MkString("history/candles"),
				MkString("history/sparklines"),
			})}),
			"private": MkMap(&VarMap{
				"get": MkArray(&VarArray{
					MkString("orders"),
					MkString("orders/leverage"),
					MkString("positions"),
					MkString("positions/margined"),
					MkString("orders/history"),
					MkString("fills"),
					MkString("fills/history/download/csv"),
					MkString("wallet/balances"),
					MkString("wallet/transactions"),
					MkString("wallet/transactions/download"),
					MkString("deposits/address"),
				}),
				"post": MkArray(&VarArray{
					MkString("orders"),
					MkString("orders/batch"),
					MkString("orders/leverage"),
					MkString("positions/change_margin"),
				}),
				"put": MkArray(&VarArray{
					MkString("orders"),
					MkString("orders/batch"),
				}),
				"delete": MkArray(&VarArray{
					MkString("orders"),
					MkString("orders/all"),
					MkString("orders/batch"),
				}),
			}),
		}),
		"fees": MkMap(&VarMap{"trading": MkMap(&VarMap{
			"tierBased":  MkBool(true),
			"percentage": MkBool(true),
			"taker":      OpDiv(MkNumber(0.15), MkInteger(100)),
			"maker":      OpDiv(MkNumber(0.10), MkInteger(100)),
			"tiers": MkMap(&VarMap{
				"taker": MkArray(&VarArray{
					MkArray(&VarArray{
						MkInteger(0),
						OpDiv(MkNumber(0.15), MkInteger(100)),
					}),
					MkArray(&VarArray{
						MkInteger(100),
						OpDiv(MkNumber(0.13), MkInteger(100)),
					}),
					MkArray(&VarArray{
						MkInteger(250),
						OpDiv(MkNumber(0.13), MkInteger(100)),
					}),
					MkArray(&VarArray{
						MkInteger(1000),
						OpDiv(MkNumber(0.1), MkInteger(100)),
					}),
					MkArray(&VarArray{
						MkInteger(5000),
						OpDiv(MkNumber(0.09), MkInteger(100)),
					}),
					MkArray(&VarArray{
						MkInteger(10000),
						OpDiv(MkNumber(0.075), MkInteger(100)),
					}),
					MkArray(&VarArray{
						MkInteger(20000),
						OpDiv(MkNumber(0.065), MkInteger(100)),
					}),
				}),
				"maker": MkArray(&VarArray{
					MkArray(&VarArray{
						MkInteger(0),
						OpDiv(MkNumber(0.1), MkInteger(100)),
					}),
					MkArray(&VarArray{
						MkInteger(100),
						OpDiv(MkNumber(0.1), MkInteger(100)),
					}),
					MkArray(&VarArray{
						MkInteger(250),
						OpDiv(MkNumber(0.09), MkInteger(100)),
					}),
					MkArray(&VarArray{
						MkInteger(1000),
						OpDiv(MkNumber(0.075), MkInteger(100)),
					}),
					MkArray(&VarArray{
						MkInteger(5000),
						OpDiv(MkNumber(0.06), MkInteger(100)),
					}),
					MkArray(&VarArray{
						MkInteger(10000),
						OpDiv(MkNumber(0.05), MkInteger(100)),
					}),
					MkArray(&VarArray{
						MkInteger(20000),
						OpDiv(MkNumber(0.05), MkInteger(100)),
					}),
				}),
			}),
		})}),
		"precisionMode": TICK_SIZE,
		"requiredCredentials": MkMap(&VarMap{
			"apiKey": MkBool(true),
			"secret": MkBool(false),
		}),
		"exceptions": MkMap(&VarMap{
			"exact": MkMap(&VarMap{
				"insufficient_margin":               InsufficientFunds,
				"order_size_exceed_available":       InvalidOrder,
				"risk_limits_breached":              BadRequest,
				"invalid_contract":                  BadSymbol,
				"immediate_liquidation":             InvalidOrder,
				"out_of_bankruptcy":                 InvalidOrder,
				"self_matching_disrupted_post_only": InvalidOrder,
				"immediate_execution_post_only":     InvalidOrder,
				"bad_schema":                        BadRequest,
				"invalid_api_key":                   AuthenticationError,
				"invalid_signature":                 AuthenticationError,
				"open_order_not_found":              OrderNotFound,
				"unavailable":                       ExchangeNotAvailable,
			}),
			"broad": MkMap(&VarMap{}),
		}),
	}))
}

func (this *Delta) FetchTime(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetSettings"), params)
	_ = response
	result := this.SafeValue(response, MkString("result"), MkMap(&VarMap{}))
	_ = result
	return this.SafeIntegerProduct(result, MkString("server_time"), MkNumber(0.001))
}

func (this *Delta) FetchStatus(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetSettings"), params)
	_ = response
	result := this.SafeValue(response, MkString("result"), MkMap(&VarMap{}))
	_ = result
	underMaintenance := this.SafeValue(result, MkString("under_maintenance"))
	_ = underMaintenance
	status := OpTrinary(OpEq2(underMaintenance, MkString("true")), MkString("maintenance"), MkString("ok"))
	_ = status
	updated := this.SafeIntegerProduct(result, MkString("server_time"), MkNumber(0.001))
	_ = updated
	*this.At(MkString("status")) = this.Extend(*this.At(MkString("status")), MkMap(&VarMap{
		"status":  status,
		"updated": updated,
	}))
	return *this.At(MkString("status"))
}

func (this *Delta) FetchCurrencies(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetAssets"), params)
	_ = response
	currencies := this.SafeValue(response, MkString("result"), MkArray(&VarArray{}))
	_ = currencies
	result := MkMap(&VarMap{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, currencies.Length)); OpIncr(&i) {
		currency := *(currencies).At(i)
		_ = currency
		id := this.SafeString(currency, MkString("symbol"))
		_ = id
		numericId := this.SafeInteger(currency, MkString("id"))
		_ = numericId
		code := this.SafeCurrencyCode(id)
		_ = code
		depositStatus := this.SafeString(currency, MkString("deposit_status"))
		_ = depositStatus
		withdrawalStatus := this.SafeString(currency, MkString("withdrawal_status"))
		_ = withdrawalStatus
		depositsEnabled := OpEq2(depositStatus, MkString("enabled"))
		_ = depositsEnabled
		withdrawalsEnabled := OpEq2(withdrawalStatus, MkString("enabled"))
		_ = withdrawalsEnabled
		active := OpAnd(depositsEnabled, withdrawalsEnabled)
		_ = active
		precision := this.SafeInteger(currency, MkString("precision"))
		_ = precision
		*(result).At(code) = MkMap(&VarMap{
			"id":        id,
			"numericId": numericId,
			"code":      code,
			"name":      this.SafeString(currency, MkString("name")),
			"info":      currency,
			"active":    active,
			"fee":       this.SafeNumber(currency, MkString("base_withdrawal_fee")),
			"precision": OpDiv(MkInteger(1), Math.Pow(MkInteger(10), precision)),
			"limits": MkMap(&VarMap{
				"amount": MkMap(&VarMap{
					"min": MkUndefined(),
					"max": MkUndefined(),
				}),
				"withdraw": MkMap(&VarMap{
					"min": this.SafeNumber(currency, MkString("min_withdrawal_amount")),
					"max": MkUndefined(),
				}),
			}),
		})
	}
	return result
}

func (this *Delta) LoadMarkets(goArgs ...*Variant) *Variant {
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
	marketsByNumericId := this.SafeValue(*this.At(MkString("options")), MkString("marketsByNumericId"))
	_ = marketsByNumericId
	if IsTrue(OpOr(OpEq2(marketsByNumericId, MkUndefined()), reload)) {
		*(*this.At(MkString("options"))).At(MkString("marketsByNumericId")) = this.IndexBy(*this.At(MkString("markets")), MkString("numericId"))
	}
	return markets
}

func (this *Delta) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetProducts"), params)
	_ = response
	markets := this.SafeValue(response, MkString("result"), MkArray(&VarArray{}))
	_ = markets
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, markets.Length)); OpIncr(&i) {
		market := *(markets).At(i)
		_ = market
		type_ := this.SafeString(market, MkString("contract_type"))
		_ = type_
		quotingAsset := this.SafeValue(market, MkString("quoting_asset"), MkMap(&VarMap{}))
		_ = quotingAsset
		underlyingAsset := this.SafeValue(market, MkString("underlying_asset"), MkMap(&VarMap{}))
		_ = underlyingAsset
		baseId := this.SafeString(underlyingAsset, MkString("symbol"))
		_ = baseId
		quoteId := this.SafeString(quotingAsset, MkString("symbol"))
		_ = quoteId
		id := this.SafeString(market, MkString("symbol"))
		_ = id
		numericId := this.SafeInteger(market, MkString("id"))
		_ = numericId
		base := this.SafeCurrencyCode(baseId)
		_ = base
		quote := this.SafeCurrencyCode(quoteId)
		_ = quote
		symbol := OpCopy(id)
		_ = symbol
		swap := OpCopy(MkBool(false))
		_ = swap
		future := OpCopy(MkBool(false))
		_ = future
		option := OpCopy(MkBool(false))
		_ = option
		if IsTrue(OpEq2(type_, MkString("perpetual_futures"))) {
			type_ = MkString("swap")
			swap = OpCopy(MkBool(true))
			future = OpCopy(MkBool(false))
			option = OpCopy(MkBool(false))
			if IsTrue(OpLw(id.IndexOf(MkString("_")), MkInteger(0))) {
				symbol = OpAdd(base, OpAdd(MkString("/"), quote))
			}
		} else {
			if IsTrue(OpOr(OpEq2(type_, MkString("call_options")), OpOr(OpEq2(type_, MkString("put_options")), OpEq2(type_, MkString("move_options"))))) {
				type_ = MkString("option")
				swap = OpCopy(MkBool(false))
				option = OpCopy(MkBool(true))
				future = OpCopy(MkBool(false))
			} else {
				if IsTrue(OpEq2(type_, MkString("futures"))) {
					type_ = MkString("future")
					swap = OpCopy(MkBool(false))
					option = OpCopy(MkBool(false))
					future = OpCopy(MkBool(true))
				}
			}
		}
		precision := MkMap(&VarMap{
			"amount": MkNumber(1.0),
			"price":  this.SafeNumber(market, MkString("tick_size")),
		})
		_ = precision
		limits := MkMap(&VarMap{
			"amount": MkMap(&VarMap{
				"min": MkNumber(1.0),
				"max": this.SafeNumber(market, MkString("position_size_limit")),
			}),
			"price": MkMap(&VarMap{
				"min": *(precision).At(MkString("price")),
				"max": MkUndefined(),
			}),
			"cost": MkMap(&VarMap{
				"min": this.SafeNumber(market, MkString("min_size")),
				"max": MkUndefined(),
			}),
		})
		_ = limits
		state := this.SafeString(market, MkString("state"))
		_ = state
		active := OpEq2(state, MkString("live"))
		_ = active
		maker := this.SafeNumber(market, MkString("maker_commission_rate"))
		_ = maker
		taker := this.SafeNumber(market, MkString("taker_commission_rate"))
		_ = taker
		result.Push(MkMap(&VarMap{
			"id":        id,
			"numericId": numericId,
			"symbol":    symbol,
			"base":      base,
			"quote":     quote,
			"baseId":    baseId,
			"quoteId":   quoteId,
			"type":      type_,
			"option":    option,
			"swap":      swap,
			"future":    future,
			"maker":     maker,
			"taker":     taker,
			"precision": precision,
			"limits":    limits,
			"info":      market,
			"active":    active,
		}))
	}
	return result
}

func (this *Delta) ParseTicker(goArgs ...*Variant) *Variant {
	ticker := GoGetArg(goArgs, 0, MkUndefined())
	_ = ticker
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.SafeIntegerProduct(ticker, MkString("timestamp"), MkNumber(0.001))
	_ = timestamp
	marketId := this.SafeString(ticker, MkString("symbol"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market)
	_ = symbol
	last := this.SafeNumber(ticker, MkString("close"))
	_ = last
	open := this.SafeNumber(ticker, MkString("open"))
	_ = open
	change := OpCopy(MkUndefined())
	_ = change
	average := OpCopy(MkUndefined())
	_ = average
	percentage := OpCopy(MkUndefined())
	_ = percentage
	if IsTrue(OpAnd(OpNotEq2(open, MkUndefined()), OpNotEq2(last, MkUndefined()))) {
		change = OpSub(last, open)
		average = OpDiv(this.Sum(last, open), MkInteger(2))
		if IsTrue(OpNotEq2(open, MkNumber(0.0))) {
			percentage = OpMulti(OpDiv(change, open), MkInteger(100))
		}
	}
	baseVolume := this.SafeNumber(ticker, MkString("volume"))
	_ = baseVolume
	quoteVolume := this.SafeNumber(ticker, MkString("turnover"))
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

func (this *Delta) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("publicGetTickersSymbol"), this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"), MkMap(&VarMap{}))
	_ = result
	return this.ParseTicker(result, market)
}

func (this *Delta) FetchTickers(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("publicGetTickers"), params)
	_ = response
	tickers := this.SafeValue(response, MkString("result"), MkArray(&VarArray{}))
	_ = tickers
	result := MkMap(&VarMap{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, tickers.Length)); OpIncr(&i) {
		ticker := this.ParseTicker(*(tickers).At(i))
		_ = ticker
		symbol := *(ticker).At(MkString("symbol"))
		_ = symbol
		*(result).At(symbol) = OpCopy(ticker)
	}
	return this.FilterByArray(result, MkString("symbol"), symbols)
}

func (this *Delta) FetchOrderBook(goArgs ...*Variant) *Variant {
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
		*(request).At(MkString("depth")) = OpCopy(limit)
	}
	response := this.Call(MkString("publicGetL2orderbookSymbol"), this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"), MkMap(&VarMap{}))
	_ = result
	return this.ParseOrderBook(result, symbol, MkUndefined(), MkString("buy"), MkString("sell"), MkString("price"), MkString("size"))
}

func (this *Delta) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	id := this.SafeString(trade, MkString("id"))
	_ = id
	orderId := this.SafeString(trade, MkString("order_id"))
	_ = orderId
	timestamp := this.Parse8601(this.SafeString(trade, MkString("created_at")))
	_ = timestamp
	timestamp = this.SafeIntegerProduct(trade, MkString("timestamp"), MkNumber(0.001), timestamp)
	priceString := this.SafeString(trade, MkString("price"))
	_ = priceString
	amountString := this.SafeString(trade, MkString("size"))
	_ = amountString
	price := this.ParseNumber(priceString)
	_ = price
	amount := this.ParseNumber(amountString)
	_ = amount
	cost := this.ParseNumber(Precise.StringMul(priceString, amountString))
	_ = cost
	product := this.SafeValue(trade, MkString("product"), MkMap(&VarMap{}))
	_ = product
	marketId := this.SafeString(product, MkString("symbol"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market)
	_ = symbol
	sellerRole := this.SafeString(trade, MkString("seller_role"))
	_ = sellerRole
	side := this.SafeString(trade, MkString("side"))
	_ = side
	if IsTrue(OpEq2(side, MkUndefined())) {
		if IsTrue(OpEq2(sellerRole, MkString("taker"))) {
			side = MkString("sell")
		} else {
			if IsTrue(OpEq2(sellerRole, MkString("maker"))) {
				side = MkString("buy")
			}
		}
	}
	takerOrMaker := this.SafeString(trade, MkString("role"))
	_ = takerOrMaker
	metaData := this.SafeValue(trade, MkString("meta_data"), MkMap(&VarMap{}))
	_ = metaData
	type_ := this.SafeString(metaData, MkString("order_type"))
	_ = type_
	if IsTrue(OpNotEq2(type_, MkUndefined())) {
		type_ = type_.Replace(MkString("_order"), MkString(""))
	}
	feeCost := this.SafeNumber(trade, MkString("commission"))
	_ = feeCost
	fee := OpCopy(MkUndefined())
	_ = fee
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		settlingAsset := this.SafeValue(product, MkString("settling_asset"), MkMap(&VarMap{}))
		_ = settlingAsset
		feeCurrencyId := this.SafeString(settlingAsset, MkString("symbol"))
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
		"order":        orderId,
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       symbol,
		"type":         type_,
		"side":         side,
		"price":        price,
		"amount":       amount,
		"cost":         cost,
		"takerOrMaker": takerOrMaker,
		"fee":          fee,
		"info":         trade,
	})
}

func (this *Delta) FetchTrades(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("publicGetTradesSymbol"), this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"), MkArray(&VarArray{}))
	_ = result
	return this.ParseTrades(result, market, since, limit)
}

func (this *Delta) ParseOHLCV(goArgs ...*Variant) *Variant {
	ohlcv := GoGetArg(goArgs, 0, MkUndefined())
	_ = ohlcv
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	return MkArray(&VarArray{
		this.SafeTimestamp(ohlcv, MkString("time")),
		this.SafeNumber(ohlcv, MkString("open")),
		this.SafeNumber(ohlcv, MkString("high")),
		this.SafeNumber(ohlcv, MkString("low")),
		this.SafeNumber(ohlcv, MkString("close")),
		this.SafeNumber(ohlcv, MkString("volume")),
	})
}

func (this *Delta) FetchOHLCV(goArgs ...*Variant) *Variant {
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
	duration := this.ParseTimeframe(timeframe)
	_ = duration
	limit = OpTrinary(limit, limit, MkInteger(2000))
	if IsTrue(OpEq2(since, MkUndefined())) {
		end := this.Seconds()
		_ = end
		*(request).At(MkString("end")) = OpCopy(end)
		*(request).At(MkString("start")) = OpSub(end, OpMulti(limit, duration))
	} else {
		start := parseInt(OpDiv(since, MkInteger(1000)))
		_ = start
		*(request).At(MkString("start")) = OpCopy(start)
		*(request).At(MkString("end")) = this.Sum(start, OpMulti(limit, duration))
	}
	response := this.Call(MkString("publicGetHistoryCandles"), this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"), MkArray(&VarArray{}))
	_ = result
	return this.ParseOHLCVs(result, market, timeframe, since, limit)
}

func (this *Delta) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privateGetWalletBalances"), params)
	_ = response
	balances := this.SafeValue(response, MkString("result"), MkArray(&VarArray{}))
	_ = balances
	result := MkMap(&VarMap{"info": response})
	_ = result
	currenciesByNumericId := this.SafeValue(*this.At(MkString("options")), MkString("currenciesByNumericId"), MkMap(&VarMap{}))
	_ = currenciesByNumericId
	for i := MkInteger(0); IsTrue(OpLw(i, balances.Length)); OpIncr(&i) {
		balance := *(balances).At(i)
		_ = balance
		currencyId := this.SafeString(balance, MkString("asset_id"))
		_ = currencyId
		currency := this.SafeValue(currenciesByNumericId, currencyId)
		_ = currency
		code := OpTrinary(OpEq2(currency, MkUndefined()), currencyId, *(currency).At(MkString("code")))
		_ = code
		account := this.Account()
		_ = account
		*(account).At(MkString("total")) = this.SafeString(balance, MkString("balance"))
		*(account).At(MkString("free")) = this.SafeString(balance, MkString("available_balance"))
		*(result).At(code) = OpCopy(account)
	}
	return this.ParseBalance(result)
}

func (this *Delta) FetchPosition(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkUndefined())
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"product_id": *(market).At(MkString("numericId"))})
	_ = request
	response := this.Call(MkString("privateGetPositions"), this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"), MkMap(&VarMap{}))
	_ = result
	return result
}

func (this *Delta) FetchPositions(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privateGetPositionsMargined"), params)
	_ = response
	result := this.SafeValue(response, MkString("result"), MkArray(&VarArray{}))
	_ = result
	return result
}

func (this *Delta) ParseOrderStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"open":      MkString("open"),
		"pending":   MkString("open"),
		"closed":    MkString("closed"),
		"cancelled": MkString("canceled"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Delta) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	id := this.SafeString(order, MkString("id"))
	_ = id
	clientOrderId := this.SafeString(order, MkString("client_order_id"))
	_ = clientOrderId
	timestamp := this.Parse8601(this.SafeString(order, MkString("created_at")))
	_ = timestamp
	marketId := this.SafeString(order, MkString("product_id"))
	_ = marketId
	marketsByNumericId := this.SafeValue(*this.At(MkString("options")), MkString("marketsByNumericId"), MkMap(&VarMap{}))
	_ = marketsByNumericId
	market = this.SafeValue(marketsByNumericId, marketId, market)
	symbol := OpTrinary(OpEq2(market, MkUndefined()), marketId, *(market).At(MkString("symbol")))
	_ = symbol
	status := this.ParseOrderStatus(this.SafeString(order, MkString("state")))
	_ = status
	side := this.SafeString(order, MkString("side"))
	_ = side
	type_ := this.SafeString(order, MkString("order_type"))
	_ = type_
	type_ = type_.Replace(MkString("_order"), MkString(""))
	price := this.SafeNumber(order, MkString("limit_price"))
	_ = price
	amount := this.SafeNumber(order, MkString("size"))
	_ = amount
	remaining := this.SafeNumber(order, MkString("unfilled_size"))
	_ = remaining
	average := this.SafeNumber(order, MkString("average_fill_price"))
	_ = average
	fee := OpCopy(MkUndefined())
	_ = fee
	feeCost := this.SafeNumber(order, MkString("paid_commission"))
	_ = feeCost
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		feeCurrencyCode := OpCopy(MkUndefined())
		_ = feeCurrencyCode
		if IsTrue(OpNotEq2(market, MkUndefined())) {
			settlingAsset := this.SafeValue(*(market).At(MkString("info")), MkString("settling_asset"), MkMap(&VarMap{}))
			_ = settlingAsset
			feeCurrencyId := this.SafeString(settlingAsset, MkString("symbol"))
			_ = feeCurrencyId
			feeCurrencyCode = this.SafeCurrencyCode(feeCurrencyId)
		}
		fee = MkMap(&VarMap{
			"cost":     feeCost,
			"currency": feeCurrencyCode,
		})
	}
	return this.SafeOrder(MkMap(&VarMap{
		"info":               order,
		"id":                 id,
		"clientOrderId":      clientOrderId,
		"timestamp":          timestamp,
		"datetime":           this.Iso8601(timestamp),
		"lastTradeTimestamp": MkUndefined(),
		"symbol":             symbol,
		"type":               type_,
		"side":               side,
		"price":              price,
		"amount":             amount,
		"cost":               MkUndefined(),
		"average":            average,
		"filled":             MkUndefined(),
		"remaining":          remaining,
		"status":             status,
		"fee":                fee,
		"trades":             MkUndefined(),
	}))
}

func (this *Delta) CreateOrder(goArgs ...*Variant) *Variant {
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
	orderType := OpAdd(type_, MkString("_order"))
	_ = orderType
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{
		"product_id": *(market).At(MkString("numericId")),
		"size":       this.AmountToPrecision(symbol, amount),
		"side":       side,
		"order_type": orderType,
	})
	_ = request
	if IsTrue(OpEq2(type_, MkString("limit"))) {
		*(request).At(MkString("limit_price")) = this.PriceToPrecision(symbol, price)
	}
	clientOrderId := this.SafeString2(params, MkString("clientOrderId"), MkString("client_order_id"))
	_ = clientOrderId
	params = this.Omit(params, MkArray(&VarArray{
		MkString("clientOrderId"),
		MkString("client_order_id"),
	}))
	if IsTrue(OpNotEq2(clientOrderId, MkUndefined())) {
		*(request).At(MkString("client_order_id")) = OpCopy(clientOrderId)
	}
	response := this.Call(MkString("privatePostOrders"), this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"), MkMap(&VarMap{}))
	_ = result
	return this.ParseOrder(result, market)
}

func (this *Delta) EditOrder(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{
		"id":         parseInt(id),
		"product_id": *(market).At(MkString("numericId")),
	})
	_ = request
	if IsTrue(OpNotEq2(amount, MkUndefined())) {
		*(request).At(MkString("size")) = parseInt(this.AmountToPrecision(symbol, amount))
	}
	if IsTrue(OpNotEq2(price, MkUndefined())) {
		*(request).At(MkString("limit_price")) = this.PriceToPrecision(symbol, price)
	}
	response := this.Call(MkString("privatePutOrders"), this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"))
	_ = result
	return this.ParseOrder(result, market)
}

func (this *Delta) CancelOrder(goArgs ...*Variant) *Variant {
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
		"id":         parseInt(id),
		"product_id": *(market).At(MkString("numericId")),
	})
	_ = request
	response := this.Call(MkString("privateDeleteOrders"), this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"))
	_ = result
	return this.ParseOrder(result, market)
}

func (this *Delta) CancelAllOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" cancelAllOrders() requires a symbol argument"))))
	}
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"product_id": *(market).At(MkString("numericId"))})
	_ = request
	response := this.Call(MkString("privateDeleteOrdersAll"), this.Extend(request, params))
	_ = response
	return response
}

func (this *Delta) FetchOpenOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	return this.FetchOrdersWithMethod(MkString("privateGetOrders"), symbol, since, limit, params)
}

func (this *Delta) FetchClosedOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	return this.FetchOrdersWithMethod(MkString("privateGetOrdersHistory"), symbol, since, limit, params)
}

func (this *Delta) FetchOrdersWithMethod(goArgs ...*Variant) *Variant {
	method := GoGetArg(goArgs, 0, MkUndefined())
	_ = method
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
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
		*(request).At(MkString("product_ids")) = *(market).At(MkString("numericId"))
	}
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("start_time")) = OpAdd(since.ToString(), MkString("000"))
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("page_size")) = OpCopy(limit)
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"), MkArray(&VarArray{}))
	_ = result
	return this.ParseOrders(result, market, since, limit)
}

func (this *Delta) FetchMyTrades(goArgs ...*Variant) *Variant {
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
		*(request).At(MkString("product_ids")) = *(market).At(MkString("numericId"))
	}
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("start_time")) = OpAdd(since.ToString(), MkString("000"))
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("page_size")) = OpCopy(limit)
	}
	response := this.Call(MkString("privateGetFills"), this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"), MkArray(&VarArray{}))
	_ = result
	return this.ParseTrades(result, market, since, limit)
}

func (this *Delta) FetchLedger(goArgs ...*Variant) *Variant {
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
		*(request).At(MkString("asset_id")) = *(currency).At(MkString("numericId"))
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("page_size")) = OpCopy(limit)
	}
	response := this.Call(MkString("privateGetWalletTransactions"), this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"), MkArray(&VarArray{}))
	_ = result
	return this.ParseLedger(result, currency, since, limit)
}

func (this *Delta) ParseLedgerEntryType(goArgs ...*Variant) *Variant {
	type_ := GoGetArg(goArgs, 0, MkUndefined())
	_ = type_
	types := MkMap(&VarMap{
		"pnl":               MkString("pnl"),
		"deposit":           MkString("transaction"),
		"withdrawal":        MkString("transaction"),
		"commission":        MkString("fee"),
		"conversion":        MkString("trade"),
		"referral_bonus":    MkString("referral"),
		"commission_rebate": MkString("rebate"),
	})
	_ = types
	return this.SafeString(types, type_, type_)
}

func (this *Delta) ParseLedgerEntry(goArgs ...*Variant) *Variant {
	item := GoGetArg(goArgs, 0, MkUndefined())
	_ = item
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	id := this.SafeString(item, MkString("uuid"))
	_ = id
	direction := OpCopy(MkUndefined())
	_ = direction
	account := OpCopy(MkUndefined())
	_ = account
	metaData := this.SafeValue(item, MkString("meta_data"), MkMap(&VarMap{}))
	_ = metaData
	referenceId := this.SafeString(metaData, MkString("transaction_id"))
	_ = referenceId
	referenceAccount := OpCopy(MkUndefined())
	_ = referenceAccount
	type_ := this.SafeString(item, MkString("transaction_type"))
	_ = type_
	if IsTrue(OpOr(OpEq2(type_, MkString("deposit")), OpOr(OpEq2(type_, MkString("commission_rebate")), OpOr(OpEq2(type_, MkString("referral_bonus")), OpOr(OpEq2(type_, MkString("pnl")), OpOr(OpEq2(type_, MkString("withdrawal_cancellation")), OpEq2(type_, MkString("promo_credit")))))))) {
		direction = MkString("in")
	} else {
		if IsTrue(OpOr(OpEq2(type_, MkString("withdrawal")), OpOr(OpEq2(type_, MkString("commission")), OpOr(OpEq2(type_, MkString("conversion")), OpEq2(type_, MkString("perpetual_futures_funding")))))) {
			direction = MkString("out")
		}
	}
	type_ = this.ParseLedgerEntryType(type_)
	currencyId := this.SafeInteger(item, MkString("asset_id"))
	_ = currencyId
	currenciesByNumericId := this.SafeValue(*this.At(MkString("options")), MkString("currenciesByNumericId"))
	_ = currenciesByNumericId
	currency = this.SafeValue(currenciesByNumericId, currencyId, currency)
	code := OpTrinary(OpEq2(currency, MkUndefined()), MkUndefined(), *(currency).At(MkString("code")))
	_ = code
	amount := this.SafeNumber(item, MkString("amount"))
	_ = amount
	timestamp := this.Parse8601(this.SafeString(item, MkString("created_at")))
	_ = timestamp
	after := this.SafeNumber(item, MkString("balance"))
	_ = after
	before := Math.Max(MkInteger(0), OpSub(after, amount))
	_ = before
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

func (this *Delta) FetchDepositAddress(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{"asset_symbol": *(currency).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("privateGetDepositsAddress"), this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"), MkMap(&VarMap{}))
	_ = result
	address := this.SafeString(result, MkString("address"))
	_ = address
	this.CheckAddress(address)
	return MkMap(&VarMap{
		"currency": code,
		"address":  address,
		"tag":      MkUndefined(),
		"info":     response,
	})
}

func (this *Delta) Sign(goArgs ...*Variant) *Variant {
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
	requestPath := OpAdd(MkString("/"), OpAdd(*this.At(MkString("version")), OpAdd(MkString("/"), this.ImplodeParams(path, params))))
	_ = requestPath
	url := OpAdd(*(*(*this.At(MkString("urls"))).At(MkString("api"))).At(api), requestPath)
	_ = url
	query := this.Omit(params, this.ExtractParams(path))
	_ = query
	if IsTrue(OpEq2(api, MkString("public"))) {
		if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
			OpAddAssign(&url, OpAdd(MkString("?"), this.Urlencode(query)))
		}
	} else {
		if IsTrue(OpEq2(api, MkString("private"))) {
			this.CheckRequiredCredentials()
			timestamp := (this.Seconds()).Call(MkString("toString"))
			_ = timestamp
			headers = MkMap(&VarMap{
				"api-key":   *this.At(MkString("apiKey")),
				"timestamp": timestamp,
			})
			auth := OpAdd(method, OpAdd(timestamp, requestPath))
			_ = auth
			if IsTrue(OpOr(OpEq2(method, MkString("GET")), OpEq2(method, MkString("DELETE")))) {
				if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
					queryString := OpAdd(MkString("?"), this.Urlencode(query))
					_ = queryString
					OpAddAssign(&auth, queryString)
					OpAddAssign(&url, queryString)
				}
			} else {
				body = this.Json(query)
				OpAddAssign(&auth, body)
				*(headers).At(MkString("Content-Type")) = MkString("application/json")
			}
			signature := this.Hmac(this.Encode(auth), this.Encode(*this.At(MkString("secret"))))
			_ = signature
			*(headers).At(MkString("signature")) = OpCopy(signature)
		}
	}
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Delta) HandleErrors(goArgs ...*Variant) *Variant {
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
	error := this.SafeValue(response, MkString("error"), MkMap(&VarMap{}))
	_ = error
	errorCode := this.SafeString(error, MkString("code"))
	_ = errorCode
	if IsTrue(OpNotEq2(errorCode, MkUndefined())) {
		feedback := OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))
		_ = feedback
		this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("exact")), errorCode, feedback)
		this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("broad")), errorCode, feedback)
		panic(NewExchangeError(feedback))
	}
	return MkUndefined()
}
