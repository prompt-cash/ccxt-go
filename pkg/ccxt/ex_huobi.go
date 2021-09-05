package ccxt

import ()

type Huobi struct {
	*ExchangeBase
}

var _ Exchange = (*Huobi)(nil)

func init() {
	exchange := &Huobi{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Huobi) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("huobi"),
		"name": MkString("Huobi"),
		"countries": MkArray(&VarArray{
			MkString("CN"),
		}),
		"rateLimit":    MkInteger(2000),
		"userAgent":    *(*this.At(MkString("userAgents"))).At(MkString("chrome39")),
		"version":      MkString("v1"),
		"accounts":     MkUndefined(),
		"accountsById": MkUndefined(),
		"hostname":     MkString("api.huobi.pro"),
		"pro":          MkBool(true),
		"has": MkMap(&VarMap{
			"cancelAllOrders":     MkBool(true),
			"cancelOrder":         MkBool(true),
			"cancelOrders":        MkBool(true),
			"CORS":                MkBool(false),
			"createOrder":         MkBool(true),
			"fetchBalance":        MkBool(true),
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
			"fetchTradingLimits":  MkBool(true),
			"fetchWithdrawals":    MkBool(true),
			"withdraw":            MkBool(true),
		}),
		"timeframes": MkMap(&VarMap{
			"1m":  MkString("1min"),
			"5m":  MkString("5min"),
			"15m": MkString("15min"),
			"30m": MkString("30min"),
			"1h":  MkString("60min"),
			"4h":  MkString("4hour"),
			"1d":  MkString("1day"),
			"1w":  MkString("1week"),
			"1M":  MkString("1mon"),
			"1y":  MkString("1year"),
		}),
		"urls": MkMap(&VarMap{
			"test": MkMap(&VarMap{
				"market":  MkString("https://api.testnet.huobi.pro"),
				"public":  MkString("https://api.testnet.huobi.pro"),
				"private": MkString("https://api.testnet.huobi.pro"),
			}),
			"logo": MkString("https://user-images.githubusercontent.com/1294454/76137448-22748a80-604e-11ea-8069-6e389271911d.jpg"),
			"api": MkMap(&VarMap{
				"market":    MkString("https://{hostname}"),
				"public":    MkString("https://{hostname}"),
				"private":   MkString("https://{hostname}"),
				"v2Public":  MkString("https://{hostname}"),
				"v2Private": MkString("https://{hostname}"),
			}),
			"www": MkString("https://www.huobi.com"),
			"referral": MkMap(&VarMap{
				"url":      MkString("https://www.huobi.com/en-us/topic/double-reward/?invite_code=6rmm2223"),
				"discount": MkNumber(0.15),
			}),
			"doc": MkArray(&VarArray{
				MkString("https://huobiapi.github.io/docs/spot/v1/cn/"),
				MkString("https://huobiapi.github.io/docs/dm/v1/cn/"),
				MkString("https://huobiapi.github.io/docs/coin_margined_swap/v1/cn/"),
				MkString("https://huobiapi.github.io/docs/usdt_swap/v1/cn/"),
				MkString("https://huobiapi.github.io/docs/option/v1/cn/"),
			}),
			"fees": MkString("https://www.huobi.com/about/fee/"),
		}),
		"api": MkMap(&VarMap{
			"v2Public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("reference/currencies"),
				MkString("market-status"),
			})}),
			"v2Private": MkMap(&VarMap{
				"get": MkArray(&VarArray{
					MkString("account/ledger"),
					MkString("account/withdraw/quota"),
					MkString("account/withdraw/address"),
					MkString("account/deposit/address"),
					MkString("account/repayment"),
					MkString("reference/transact-fee-rate"),
					MkString("account/asset-valuation"),
					MkString("point/account"),
					MkString("sub-user/user-list"),
					MkString("sub-user/user-state"),
					MkString("sub-user/account-list"),
					MkString("sub-user/deposit-address"),
					MkString("sub-user/query-deposit"),
					MkString("user/api-key"),
					MkString("user/uid"),
					MkString("algo-orders/opening"),
					MkString("algo-orders/history"),
					MkString("algo-orders/specific"),
					MkString("c2c/offers"),
					MkString("c2c/offer"),
					MkString("c2c/transactions"),
					MkString("c2c/repayment"),
					MkString("c2c/account"),
					MkString("etp/reference"),
					MkString("etp/transactions"),
					MkString("etp/transaction"),
					MkString("etp/rebalance"),
					MkString("etp/limit"),
				}),
				"post": MkArray(&VarArray{
					MkString("account/transfer"),
					MkString("account/repayment"),
					MkString("point/transfer"),
					MkString("sub-user/management"),
					MkString("sub-user/creation"),
					MkString("sub-user/tradable-market"),
					MkString("sub-user/transferability"),
					MkString("sub-user/api-key-generation"),
					MkString("sub-user/api-key-modification"),
					MkString("sub-user/api-key-deletion"),
					MkString("sub-user/deduct-mode"),
					MkString("algo-orders"),
					MkString("algo-orders/cancel-all-after"),
					MkString("algo-orders/cancellation"),
					MkString("c2c/offer"),
					MkString("c2c/cancellation"),
					MkString("c2c/cancel-all"),
					MkString("c2c/repayment"),
					MkString("c2c/transfer"),
					MkString("etp/creation"),
					MkString("etp/redemption"),
					MkString("etp/{transactId}/cancel"),
					MkString("etp/batch-cancel"),
				}),
			}),
			"market": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("history/kline"),
				MkString("detail/merged"),
				MkString("depth"),
				MkString("trade"),
				MkString("history/trade"),
				MkString("detail"),
				MkString("tickers"),
				MkString("etp"),
			})}),
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("common/symbols"),
				MkString("common/currencys"),
				MkString("common/timestamp"),
				MkString("common/exchange"),
				MkString("settings/currencys"),
			})}),
			"private": MkMap(&VarMap{
				"get": MkArray(&VarArray{
					MkString("account/accounts"),
					MkString("account/accounts/{id}/balance"),
					MkString("account/accounts/{sub-uid}"),
					MkString("account/history"),
					MkString("cross-margin/loan-info"),
					MkString("margin/loan-info"),
					MkString("fee/fee-rate/get"),
					MkString("order/openOrders"),
					MkString("order/orders"),
					MkString("order/orders/{id}"),
					MkString("order/orders/{id}/matchresults"),
					MkString("order/orders/getClientOrder"),
					MkString("order/history"),
					MkString("order/matchresults"),
					MkString("dw/withdraw-virtual/addresses"),
					MkString("query/deposit-withdraw"),
					MkString("margin/loan-info"),
					MkString("margin/loan-orders"),
					MkString("margin/accounts/balance"),
					MkString("cross-margin/loan-orders"),
					MkString("cross-margin/accounts/balance"),
					MkString("points/actions"),
					MkString("points/orders"),
					MkString("subuser/aggregate-balance"),
					MkString("stable-coin/exchange_rate"),
					MkString("stable-coin/quote"),
				}),
				"post": MkArray(&VarArray{
					MkString("account/transfer"),
					MkString("futures/transfer"),
					MkString("order/batch-orders"),
					MkString("order/orders/place"),
					MkString("order/orders/submitCancelClientOrder"),
					MkString("order/orders/batchCancelOpenOrders"),
					MkString("order/orders"),
					MkString("order/orders/{id}/place"),
					MkString("order/orders/{id}/submitcancel"),
					MkString("order/orders/batchcancel"),
					MkString("dw/balance/transfer"),
					MkString("dw/withdraw/api/create"),
					MkString("dw/withdraw-virtual/create"),
					MkString("dw/withdraw-virtual/{id}/place"),
					MkString("dw/withdraw-virtual/{id}/cancel"),
					MkString("dw/transfer-in/margin"),
					MkString("dw/transfer-out/margin"),
					MkString("margin/orders"),
					MkString("margin/orders/{id}/repay"),
					MkString("cross-margin/transfer-in"),
					MkString("cross-margin/transfer-out"),
					MkString("cross-margin/orders"),
					MkString("cross-margin/orders/{id}/repay"),
					MkString("stable-coin/exchange"),
					MkString("subuser/transfer"),
				}),
			}),
		}),
		"fees": MkMap(&VarMap{"trading": MkMap(&VarMap{
			"feeSide":    MkString("get"),
			"tierBased":  MkBool(false),
			"percentage": MkBool(true),
			"maker":      MkNumber(0.002),
			"taker":      MkNumber(0.002),
		})}),
		"exceptions": MkMap(&VarMap{
			"broad": MkMap(&VarMap{
				"contract is restricted of closing positions on API.  Please contact customer service": OnMaintenance,
				"maintain": OnMaintenance,
			}),
			"exact": MkMap(&VarMap{
				"bad-request":                               BadRequest,
				"base-date-limit-error":                     BadRequest,
				"api-not-support-temp-addr":                 PermissionDenied,
				"timeout":                                   RequestTimeout,
				"gateway-internal-error":                    ExchangeNotAvailable,
				"account-frozen-balance-insufficient-error": InsufficientFunds,
				"invalid-amount":                            InvalidOrder,
				"order-limitorder-amount-min-error":         InvalidOrder,
				"order-limitorder-amount-max-error":         InvalidOrder,
				"order-marketorder-amount-min-error":        InvalidOrder,
				"order-limitorder-price-min-error":          InvalidOrder,
				"order-limitorder-price-max-error":          InvalidOrder,
				"order-holding-limit-failed":                InvalidOrder,
				"order-orderprice-precision-error":          InvalidOrder,
				"order-etp-nav-price-max-error":             InvalidOrder,
				"order-orderstate-error":                    OrderNotFound,
				"order-queryorder-invalid":                  OrderNotFound,
				"order-update-error":                        ExchangeNotAvailable,
				"api-signature-check-failed":                AuthenticationError,
				"api-signature-not-valid":                   AuthenticationError,
				"base-record-invalid":                       OrderNotFound,
				"base-symbol-trade-disabled":                BadSymbol,
				"base-symbol-error":                         BadSymbol,
				"system-maintenance":                        OnMaintenance,
				"invalid symbol":                            BadSymbol,
				"symbol trade not open now":                 BadSymbol,
			}),
		}),
		"options": MkMap(&VarMap{
			"fetchOrdersByStatesMethod":         MkString("private_get_order_orders"),
			"fetchOpenOrdersMethod":             MkString("fetch_open_orders_v1"),
			"createMarketBuyOrderRequiresPrice": MkBool(true),
			"fetchMarketsMethod":                MkString("publicGetCommonSymbols"),
			"fetchBalanceMethod":                MkString("privateGetAccountAccountsIdBalance"),
			"createOrderMethod":                 MkString("privatePostOrderOrdersPlace"),
			"language":                          MkString("en-US"),
			"broker":                            MkMap(&VarMap{"id": MkString("AA03022abc")}),
		}),
		"commonCurrencies": MkMap(&VarMap{
			"GET":  MkString("Themis"),
			"GTC":  MkString("Game.com"),
			"HIT":  MkString("HitChain"),
			"HOT":  MkString("Hydro Protocol"),
			"PNT":  MkString("Penta"),
			"SBTC": MkString("Super Bitcoin"),
			"BIFI": MkString("Bitcoin File"),
		}),
	}))
}

func (this *Huobi) FetchTradingLimits(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	if IsTrue(OpEq2(symbols, MkUndefined())) {
		symbols = OpCopy(*this.At(MkString("symbols")))
	}
	result := MkMap(&VarMap{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, symbols.Length)); OpIncr(&i) {
		symbol := *(symbols).At(i)
		_ = symbol
		*(result).At(symbol) = this.FetchTradingLimitsById(this.MarketId(symbol), params)
	}
	return result
}

func (this *Huobi) FetchTradingLimitsById(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	request := MkMap(&VarMap{"symbol": id})
	_ = request
	response := this.Call(MkString("publicGetCommonExchange"), this.Extend(request, params))
	_ = response
	return this.ParseTradingLimits(this.SafeValue(response, MkString("data"), MkMap(&VarMap{})))
}

func (this *Huobi) ParseTradingLimits(goArgs ...*Variant) *Variant {
	limits := GoGetArg(goArgs, 0, MkUndefined())
	_ = limits
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	return MkMap(&VarMap{
		"info": limits,
		"limits": MkMap(&VarMap{"amount": MkMap(&VarMap{
			"min": this.SafeNumber(limits, MkString("limit-order-must-greater-than")),
			"max": this.SafeNumber(limits, MkString("limit-order-must-less-than")),
		})}),
	})
}

func (this *Huobi) CostToPrecision(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	cost := GoGetArg(goArgs, 1, MkUndefined())
	_ = cost
	return this.DecimalToPrecision(cost, TRUNCATE, *(*(*(*this.At(MkString("markets"))).At(symbol)).At(MkString("precision"))).At(MkString("cost")), *this.At(MkString("precisionMode")))
}

func (this *Huobi) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	method := *(*this.At(MkString("options"))).At(MkString("fetchMarketsMethod"))
	_ = method
	response := this.Call(method, params)
	_ = response
	markets := this.SafeValue(response, MkString("data"))
	_ = markets
	numMarkets := OpCopy(markets.Length)
	_ = numMarkets
	if IsTrue(OpLw(numMarkets, MkInteger(1))) {
		panic(NewNetworkError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" publicGetCommonSymbols returned empty response: "), this.Json(markets)))))
	}
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, markets.Length)); OpIncr(&i) {
		market := *(markets).At(i)
		_ = market
		baseId := this.SafeString(market, MkString("base-currency"))
		_ = baseId
		quoteId := this.SafeString(market, MkString("quote-currency"))
		_ = quoteId
		id := OpAdd(baseId, quoteId)
		_ = id
		base := this.SafeCurrencyCode(baseId)
		_ = base
		quote := this.SafeCurrencyCode(quoteId)
		_ = quote
		symbol := OpAdd(base, OpAdd(MkString("/"), quote))
		_ = symbol
		precision := MkMap(&VarMap{
			"amount": this.SafeInteger(market, MkString("amount-precision")),
			"price":  this.SafeInteger(market, MkString("price-precision")),
			"cost":   this.SafeInteger(market, MkString("value-precision")),
		})
		_ = precision
		maker := OpTrinary(OpEq2(base, MkString("OMG")), MkInteger(0), OpDiv(MkNumber(0.2), MkInteger(100)))
		_ = maker
		taker := OpTrinary(OpEq2(base, MkString("OMG")), MkInteger(0), OpDiv(MkNumber(0.2), MkInteger(100)))
		_ = taker
		minAmount := this.SafeNumber(market, MkString("min-order-amt"), Math.Pow(MkInteger(10), OpNeg(*(precision).At(MkString("amount")))))
		_ = minAmount
		maxAmount := this.SafeNumber(market, MkString("max-order-amt"))
		_ = maxAmount
		minCost := this.SafeNumber(market, MkString("min-order-value"), MkInteger(0))
		_ = minCost
		state := this.SafeString(market, MkString("state"))
		_ = state
		active := OpEq2(state, MkString("online"))
		_ = active
		result.Push(MkMap(&VarMap{
			"id":        id,
			"symbol":    symbol,
			"base":      base,
			"quote":     quote,
			"baseId":    baseId,
			"quoteId":   quoteId,
			"active":    active,
			"precision": precision,
			"taker":     taker,
			"maker":     maker,
			"limits": MkMap(&VarMap{
				"amount": MkMap(&VarMap{
					"min": minAmount,
					"max": maxAmount,
				}),
				"price": MkMap(&VarMap{
					"min": Math.Pow(MkInteger(10), OpNeg(*(precision).At(MkString("price")))),
					"max": MkUndefined(),
				}),
				"cost": MkMap(&VarMap{
					"min": minCost,
					"max": MkUndefined(),
				}),
			}),
			"info": market,
		}))
	}
	return result
}

func (this *Huobi) ParseTicker(goArgs ...*Variant) *Variant {
	ticker := GoGetArg(goArgs, 0, MkUndefined())
	_ = ticker
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(OpNotEq2(market, MkUndefined())) {
		symbol = *(market).At(MkString("symbol"))
	}
	timestamp := this.SafeInteger(ticker, MkString("ts"))
	_ = timestamp
	bid := OpCopy(MkUndefined())
	_ = bid
	bidVolume := OpCopy(MkUndefined())
	_ = bidVolume
	ask := OpCopy(MkUndefined())
	_ = ask
	askVolume := OpCopy(MkUndefined())
	_ = askVolume
	if IsTrue(OpHasMember(MkString("bid"), ticker)) {
		if IsTrue(GoIsArray(*(ticker).At(MkString("bid")))) {
			bid = this.SafeNumber(*(ticker).At(MkString("bid")), MkInteger(0))
			bidVolume = this.SafeNumber(*(ticker).At(MkString("bid")), MkInteger(1))
		} else {
			bid = this.SafeNumber(ticker, MkString("bid"))
			bidVolume = this.SafeValue(ticker, MkString("bidSize"))
		}
	}
	if IsTrue(OpHasMember(MkString("ask"), ticker)) {
		if IsTrue(GoIsArray(*(ticker).At(MkString("ask")))) {
			ask = this.SafeNumber(*(ticker).At(MkString("ask")), MkInteger(0))
			askVolume = this.SafeNumber(*(ticker).At(MkString("ask")), MkInteger(1))
		} else {
			ask = this.SafeNumber(ticker, MkString("ask"))
			askVolume = this.SafeValue(ticker, MkString("askSize"))
		}
	}
	open := this.SafeNumber(ticker, MkString("open"))
	_ = open
	close := this.SafeNumber(ticker, MkString("close"))
	_ = close
	change := OpCopy(MkUndefined())
	_ = change
	percentage := OpCopy(MkUndefined())
	_ = percentage
	average := OpCopy(MkUndefined())
	_ = average
	if IsTrue(OpAnd(OpNotEq2(open, MkUndefined()), OpNotEq2(close, MkUndefined()))) {
		change = OpSub(close, open)
		average = OpDiv(this.Sum(open, close), MkInteger(2))
		if IsTrue(OpAnd(OpNotEq2(close, MkUndefined()), OpGt(close, MkInteger(0)))) {
			percentage = OpMulti(OpDiv(change, open), MkInteger(100))
		}
	}
	baseVolume := this.SafeNumber(ticker, MkString("amount"))
	_ = baseVolume
	quoteVolume := this.SafeNumber(ticker, MkString("vol"))
	_ = quoteVolume
	vwap := this.Vwap(baseVolume, quoteVolume)
	_ = vwap
	return MkMap(&VarMap{
		"symbol":        symbol,
		"timestamp":     timestamp,
		"datetime":      this.Iso8601(timestamp),
		"high":          this.SafeNumber(ticker, MkString("high")),
		"low":           this.SafeNumber(ticker, MkString("low")),
		"bid":           bid,
		"bidVolume":     bidVolume,
		"ask":           ask,
		"askVolume":     askVolume,
		"vwap":          vwap,
		"open":          open,
		"close":         close,
		"last":          close,
		"previousClose": MkUndefined(),
		"change":        change,
		"percentage":    percentage,
		"average":       average,
		"baseVolume":    baseVolume,
		"quoteVolume":   quoteVolume,
		"info":          ticker,
	})
}

func (this *Huobi) FetchOrderBook(goArgs ...*Variant) *Variant {
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
		"symbol": *(market).At(MkString("id")),
		"type":   MkString("step0"),
	})
	_ = request
	response := this.Call(MkString("marketGetDepth"), this.Extend(request, params))
	_ = response
	if IsTrue(OpHasMember(MkString("tick"), response)) {
		if IsTrue(OpNot(*(response).At(MkString("tick")))) {
			panic(NewBadSymbol(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" fetchOrderBook() returned empty response: "), this.Json(response)))))
		}
		tick := this.SafeValue(response, MkString("tick"))
		_ = tick
		timestamp := this.SafeInteger(tick, MkString("ts"), this.SafeInteger(response, MkString("ts")))
		_ = timestamp
		result := this.ParseOrderBook(tick, symbol, timestamp)
		_ = result
		*(result).At(MkString("nonce")) = this.SafeInteger(tick, MkString("version"))
		return result
	}
	panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" fetchOrderBook() returned unrecognized response: "), this.Json(response)))))
	return MkUndefined()
}

func (this *Huobi) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("marketGetDetailMerged"), this.Extend(request, params))
	_ = response
	ticker := this.ParseTicker(*(response).At(MkString("tick")), market)
	_ = ticker
	timestamp := this.SafeInteger(response, MkString("ts"))
	_ = timestamp
	*(ticker).At(MkString("timestamp")) = OpCopy(timestamp)
	*(ticker).At(MkString("datetime")) = this.Iso8601(timestamp)
	return ticker
}

func (this *Huobi) FetchTickers(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("marketGetTickers"), params)
	_ = response
	tickers := this.SafeValue(response, MkString("data"))
	_ = tickers
	timestamp := this.SafeInteger(response, MkString("ts"))
	_ = timestamp
	result := MkMap(&VarMap{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, tickers.Length)); OpIncr(&i) {
		marketId := this.SafeString(*(tickers).At(i), MkString("symbol"))
		_ = marketId
		market := this.SafeMarket(marketId)
		_ = market
		symbol := *(market).At(MkString("symbol"))
		_ = symbol
		ticker := this.ParseTicker(*(tickers).At(i), market)
		_ = ticker
		*(ticker).At(MkString("timestamp")) = OpCopy(timestamp)
		*(ticker).At(MkString("datetime")) = this.Iso8601(timestamp)
		*(result).At(symbol) = OpCopy(ticker)
	}
	return this.FilterByArray(result, MkString("symbol"), symbols)
}

func (this *Huobi) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	marketId := this.SafeString(trade, MkString("symbol"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market)
	_ = symbol
	timestamp := this.SafeInteger2(trade, MkString("ts"), MkString("created-at"))
	_ = timestamp
	order := this.SafeString(trade, MkString("order-id"))
	_ = order
	side := this.SafeString(trade, MkString("direction"))
	_ = side
	type_ := this.SafeString(trade, MkString("type"))
	_ = type_
	if IsTrue(OpNotEq2(type_, MkUndefined())) {
		typeParts := type_.Split(MkString("-"))
		_ = typeParts
		side = *(typeParts).At(MkInteger(0))
		type_ = *(typeParts).At(MkInteger(1))
	}
	takerOrMaker := this.SafeString(trade, MkString("role"))
	_ = takerOrMaker
	priceString := this.SafeString(trade, MkString("price"))
	_ = priceString
	amountString := this.SafeString2(trade, MkString("filled-amount"), MkString("amount"))
	_ = amountString
	price := this.ParseNumber(priceString)
	_ = price
	amount := this.ParseNumber(amountString)
	_ = amount
	cost := this.ParseNumber(Precise.StringMul(priceString, amountString))
	_ = cost
	fee := OpCopy(MkUndefined())
	_ = fee
	feeCost := this.SafeNumber(trade, MkString("filled-fees"))
	_ = feeCost
	feeCurrency := this.SafeCurrencyCode(this.SafeString(trade, MkString("fee-currency")))
	_ = feeCurrency
	filledPoints := this.SafeNumber(trade, MkString("filled-points"))
	_ = filledPoints
	if IsTrue(OpNotEq2(filledPoints, MkUndefined())) {
		if IsTrue(OpOr(OpEq2(feeCost, MkUndefined()), OpEq2(feeCost, MkNumber(0.0)))) {
			feeCost = OpCopy(filledPoints)
			feeCurrency = this.SafeCurrencyCode(this.SafeString(trade, MkString("fee-deduct-currency")))
		}
	}
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		fee = MkMap(&VarMap{
			"cost":     feeCost,
			"currency": feeCurrency,
		})
	}
	tradeId := this.SafeString2(trade, MkString("trade-id"), MkString("tradeId"))
	_ = tradeId
	id := this.SafeString(trade, MkString("id"), tradeId)
	_ = id
	return MkMap(&VarMap{
		"id":           id,
		"info":         trade,
		"order":        order,
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       symbol,
		"type":         type_,
		"side":         side,
		"takerOrMaker": takerOrMaker,
		"price":        price,
		"amount":       amount,
		"cost":         cost,
		"fee":          fee,
	})
}

func (this *Huobi) FetchOrderTrades(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"id": id})
	_ = request
	response := this.Call(MkString("privateGetOrderOrdersIdMatchresults"), this.Extend(request, params))
	_ = response
	return this.ParseTrades(*(response).At(MkString("data")), MkUndefined(), since, limit)
}

func (this *Huobi) FetchMyTrades(goArgs ...*Variant) *Variant {
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
		*(request).At(MkString("symbol")) = *(market).At(MkString("id"))
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("size")) = OpCopy(limit)
	}
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("start-date")) = this.Ymd(since)
		*(request).At(MkString("end-date")) = this.Ymd(this.Sum(since, MkInteger(86400000)))
	}
	response := this.Call(MkString("privateGetOrderMatchresults"), this.Extend(request, params))
	_ = response
	return this.ParseTrades(*(response).At(MkString("data")), market, since, limit)
}

func (this *Huobi) FetchTrades(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkInteger(1000))
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("size")) = OpCopy(limit)
	}
	response := this.Call(MkString("marketGetHistoryTrade"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"))
	_ = data
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, data.Length)); OpIncr(&i) {
		trades := this.SafeValue(*(data).At(i), MkString("data"), MkArray(&VarArray{}))
		_ = trades
		for j := MkInteger(0); IsTrue(OpLw(j, trades.Length)); OpIncr(&j) {
			trade := this.ParseTrade(*(trades).At(j), market)
			_ = trade
			result.Push(trade)
		}
	}
	result = this.SortBy(result, MkString("timestamp"))
	return this.FilterBySymbolSinceLimit(result, symbol, since, limit)
}

func (this *Huobi) ParseOHLCV(goArgs ...*Variant) *Variant {
	ohlcv := GoGetArg(goArgs, 0, MkUndefined())
	_ = ohlcv
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	return MkArray(&VarArray{
		this.SafeTimestamp(ohlcv, MkString("id")),
		this.SafeNumber(ohlcv, MkString("open")),
		this.SafeNumber(ohlcv, MkString("high")),
		this.SafeNumber(ohlcv, MkString("low")),
		this.SafeNumber(ohlcv, MkString("close")),
		this.SafeNumber(ohlcv, MkString("amount")),
	})
}

func (this *Huobi) FetchOHLCV(goArgs ...*Variant) *Variant {
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
		"symbol": *(market).At(MkString("id")),
		"period": *(*this.At(MkString("timeframes"))).At(timeframe),
	})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("size")) = OpCopy(limit)
	}
	response := this.Call(MkString("marketGetHistoryKline"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseOHLCVs(data, market, timeframe, since, limit)
}

func (this *Huobi) FetchAccounts(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privateGetAccountAccounts"), params)
	_ = response
	return *(response).At(MkString("data"))
}

func (this *Huobi) FetchCurrencies(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	request := MkMap(&VarMap{"language": *(*this.At(MkString("options"))).At(MkString("language"))})
	_ = request
	response := this.Call(MkString("publicGetSettingsCurrencys"), this.Extend(request, params))
	_ = response
	currencies := this.SafeValue(response, MkString("data"))
	_ = currencies
	result := MkMap(&VarMap{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, currencies.Length)); OpIncr(&i) {
		currency := *(currencies).At(i)
		_ = currency
		id := this.SafeValue(currency, MkString("name"))
		_ = id
		precision := this.SafeInteger(currency, MkString("withdraw-precision"))
		_ = precision
		code := this.SafeCurrencyCode(id)
		_ = code
		active := OpAnd(*(currency).At(MkString("visible")), OpAnd(*(currency).At(MkString("deposit-enabled")), *(currency).At(MkString("withdraw-enabled"))))
		_ = active
		name := this.SafeString(currency, MkString("display-name"))
		_ = name
		*(result).At(code) = MkMap(&VarMap{
			"id":        id,
			"code":      code,
			"type":      MkString("crypto"),
			"name":      name,
			"active":    active,
			"fee":       MkUndefined(),
			"precision": precision,
			"limits": MkMap(&VarMap{
				"amount": MkMap(&VarMap{
					"min": Math.Pow(MkInteger(10), OpNeg(precision)),
					"max": Math.Pow(MkInteger(10), precision),
				}),
				"deposit": MkMap(&VarMap{
					"min": this.SafeNumber(currency, MkString("deposit-min-amount")),
					"max": Math.Pow(MkInteger(10), precision),
				}),
				"withdraw": MkMap(&VarMap{
					"min": this.SafeNumber(currency, MkString("withdraw-min-amount")),
					"max": Math.Pow(MkInteger(10), precision),
				}),
			}),
			"info": currency,
		})
	}
	return result
}

func (this *Huobi) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	this.LoadAccounts()
	method := *(*this.At(MkString("options"))).At(MkString("fetchBalanceMethod"))
	_ = method
	request := MkMap(&VarMap{"id": *(*(*this.At(MkString("accounts"))).At(MkInteger(0))).At(MkString("id"))})
	_ = request
	response := this.Call(method, this.Extend(request, params))
	_ = response
	balances := this.SafeValue(*(response).At(MkString("data")), MkString("list"), MkArray(&VarArray{}))
	_ = balances
	result := MkMap(&VarMap{"info": response})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, balances.Length)); OpIncr(&i) {
		balance := *(balances).At(i)
		_ = balance
		currencyId := this.SafeString(balance, MkString("currency"))
		_ = currencyId
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		account := OpCopy(MkUndefined())
		_ = account
		if IsTrue(OpHasMember(code, result)) {
			account = *(result).At(code)
		} else {
			account = this.Account()
		}
		if IsTrue(OpEq2(*(balance).At(MkString("type")), MkString("trade"))) {
			*(account).At(MkString("free")) = this.SafeString(balance, MkString("balance"))
		}
		if IsTrue(OpEq2(*(balance).At(MkString("type")), MkString("frozen"))) {
			*(account).At(MkString("used")) = this.SafeString(balance, MkString("balance"))
		}
		*(result).At(code) = OpCopy(account)
	}
	return this.ParseBalance(result)
}

func (this *Huobi) FetchOrdersByStates(goArgs ...*Variant) *Variant {
	states := GoGetArg(goArgs, 0, MkUndefined())
	_ = states
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 2, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 3, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 4, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"states": states})
	_ = request
	market := OpCopy(MkUndefined())
	_ = market
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
		*(request).At(MkString("symbol")) = *(market).At(MkString("id"))
	}
	method := this.SafeString(*this.At(MkString("options")), MkString("fetchOrdersByStatesMethod"), MkString("private_get_order_orders"))
	_ = method
	response := this.Call(method, this.Extend(request, params))
	_ = response
	return this.ParseOrders(*(response).At(MkString("data")), market, since, limit)
}

func (this *Huobi) FetchOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"id": id})
	_ = request
	response := this.Call(MkString("privateGetOrderOrdersId"), this.Extend(request, params))
	_ = response
	order := this.SafeValue(response, MkString("data"))
	_ = order
	return this.ParseOrder(order)
}

func (this *Huobi) FetchOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	return this.FetchOrdersByStates(MkString("pre-submitted,submitted,partial-filled,filled,partial-canceled,canceled"), symbol, since, limit, params)
}

func (this *Huobi) FetchOpenOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	method := this.SafeString(*this.At(MkString("options")), MkString("fetchOpenOrdersMethod"), MkString("fetch_open_orders_v1"))
	_ = method
	return this.Call(method, symbol, since, limit, params)
}

func (this *Huobi) FetchOpenOrdersV1(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchOpenOrdersV1() requires a symbol argument"))))
	}
	return this.FetchOrdersByStates(MkString("pre-submitted,submitted,partial-filled"), symbol, since, limit, params)
}

func (this *Huobi) FetchClosedOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	return this.FetchOrdersByStates(MkString("filled,partial-canceled,canceled"), symbol, since, limit, params)
}

func (this *Huobi) FetchOpenOrdersV2(goArgs ...*Variant) *Variant {
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
		*(request).At(MkString("symbol")) = *(market).At(MkString("id"))
	}
	accountId := this.SafeString(params, MkString("account-id"))
	_ = accountId
	if IsTrue(OpEq2(accountId, MkUndefined())) {
		this.LoadAccounts()
		for i := MkInteger(0); IsTrue(OpLw(i, (*((this).At(MkString("accounts")))).Length)); OpIncr(&i) {
			account := *(*this.At(MkString("accounts"))).At(i)
			_ = account
			if IsTrue(OpEq2(*(account).At(MkString("type")), MkString("spot"))) {
				accountId = this.SafeString(account, MkString("id"))
				if IsTrue(OpNotEq2(accountId, MkUndefined())) {
					break
				}
			}
		}
	}
	*(request).At(MkString("account-id")) = OpCopy(accountId)
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("size")) = OpCopy(limit)
	}
	omitted := this.Omit(params, MkString("account-id"))
	_ = omitted
	response := this.Call(MkString("privateGetOrderOpenOrders"), this.Extend(request, omitted))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseOrders(data, market, since, limit)
}

func (this *Huobi) ParseOrderStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"partial-filled":   MkString("open"),
		"partial-canceled": MkString("canceled"),
		"filled":           MkString("closed"),
		"canceled":         MkString("canceled"),
		"submitted":        MkString("open"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Huobi) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	id := this.SafeString(order, MkString("id"))
	_ = id
	side := OpCopy(MkUndefined())
	_ = side
	type_ := OpCopy(MkUndefined())
	_ = type_
	status := OpCopy(MkUndefined())
	_ = status
	if IsTrue(OpHasMember(MkString("type"), order)) {
		orderType := (*(order).At(MkString("type"))).Call(MkString("split"), MkString("-"))
		_ = orderType
		side = *(orderType).At(MkInteger(0))
		type_ = *(orderType).At(MkInteger(1))
		status = this.ParseOrderStatus(this.SafeString(order, MkString("state")))
	}
	marketId := this.SafeString(order, MkString("symbol"))
	_ = marketId
	market = this.SafeMarket(marketId, market)
	symbol := this.SafeSymbol(marketId, market)
	_ = symbol
	timestamp := this.SafeInteger(order, MkString("created-at"))
	_ = timestamp
	clientOrderId := this.SafeString(order, MkString("client-order-id"))
	_ = clientOrderId
	amount := this.SafeNumber(order, MkString("amount"))
	_ = amount
	filled := this.SafeNumber2(order, MkString("filled-amount"), MkString("field-amount"))
	_ = filled
	price := this.SafeNumber(order, MkString("price"))
	_ = price
	if IsTrue(OpEq2(price, MkNumber(0.0))) {
		price = OpCopy(MkUndefined())
	}
	cost := this.SafeNumber2(order, MkString("filled-cash-amount"), MkString("field-cash-amount"))
	_ = cost
	feeCost := this.SafeNumber2(order, MkString("filled-fees"), MkString("field-fees"))
	_ = feeCost
	fee := OpCopy(MkUndefined())
	_ = fee
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		feeCurrency := OpCopy(MkUndefined())
		_ = feeCurrency
		if IsTrue(OpNotEq2(market, MkUndefined())) {
			feeCurrency = OpTrinary(OpEq2(side, MkString("sell")), *(market).At(MkString("quote")), *(market).At(MkString("base")))
		}
		fee = MkMap(&VarMap{
			"cost":     feeCost,
			"currency": feeCurrency,
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
		"timeInForce":        MkUndefined(),
		"postOnly":           MkUndefined(),
		"side":               side,
		"price":              price,
		"stopPrice":          MkUndefined(),
		"average":            MkUndefined(),
		"cost":               cost,
		"amount":             amount,
		"filled":             filled,
		"remaining":          MkUndefined(),
		"status":             status,
		"fee":                fee,
		"trades":             MkUndefined(),
	}))
}

func (this *Huobi) CreateOrder(goArgs ...*Variant) *Variant {
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
	this.LoadAccounts()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{
		"account-id": *(*(*this.At(MkString("accounts"))).At(MkInteger(0))).At(MkString("id")),
		"symbol":     *(market).At(MkString("id")),
		"type":       OpAdd(side, OpAdd(MkString("-"), type_)),
	})
	_ = request
	clientOrderId := this.SafeString2(params, MkString("clientOrderId"), MkString("client-order-id"))
	_ = clientOrderId
	if IsTrue(OpEq2(clientOrderId, MkUndefined())) {
		broker := this.SafeValue(*this.At(MkString("options")), MkString("broker"), MkMap(&VarMap{}))
		_ = broker
		brokerId := this.SafeString(broker, MkString("id"))
		_ = brokerId
		*(request).At(MkString("client-order-id")) = OpAdd(brokerId, this.Uuid())
	} else {
		*(request).At(MkString("client-order-id")) = OpCopy(clientOrderId)
	}
	params = this.Omit(params, MkArray(&VarArray{
		MkString("clientOrderId"),
		MkString("client-order-id"),
	}))
	if IsTrue(OpAnd(OpEq2(type_, MkString("market")), OpEq2(side, MkString("buy")))) {
		if IsTrue(*(*this.At(MkString("options"))).At(MkString("createMarketBuyOrderRequiresPrice"))) {
			if IsTrue(OpEq2(price, MkUndefined())) {
				panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")), MkString(" market buy order requires price argument to calculate cost (total amount of quote currency to spend for buying, amount * price). To switch off this warning exception and specify cost in the amount argument, set .options['createMarketBuyOrderRequiresPrice'] = false. Make sure you know what you're doing."))))
			} else {
				*(request).At(MkString("amount")) = this.CostToPrecision(symbol, OpMulti(parseFloat(amount), parseFloat(price)))
			}
		} else {
			*(request).At(MkString("amount")) = this.CostToPrecision(symbol, amount)
		}
	} else {
		*(request).At(MkString("amount")) = this.AmountToPrecision(symbol, amount)
	}
	if IsTrue(OpOr(OpEq2(type_, MkString("limit")), OpOr(OpEq2(type_, MkString("ioc")), OpEq2(type_, MkString("limit-maker"))))) {
		*(request).At(MkString("price")) = this.PriceToPrecision(symbol, price)
	}
	method := *(*this.At(MkString("options"))).At(MkString("createOrderMethod"))
	_ = method
	response := this.Call(method, this.Extend(request, params))
	_ = response
	timestamp := this.Milliseconds()
	_ = timestamp
	id := this.SafeString(response, MkString("data"))
	_ = id
	return MkMap(&VarMap{
		"info":               response,
		"id":                 id,
		"timestamp":          timestamp,
		"datetime":           this.Iso8601(timestamp),
		"lastTradeTimestamp": MkUndefined(),
		"status":             MkUndefined(),
		"symbol":             symbol,
		"type":               type_,
		"side":               side,
		"price":              price,
		"amount":             amount,
		"filled":             MkUndefined(),
		"remaining":          MkUndefined(),
		"cost":               MkUndefined(),
		"trades":             MkUndefined(),
		"fee":                MkUndefined(),
		"clientOrderId":      MkUndefined(),
		"average":            MkUndefined(),
	})
}

func (this *Huobi) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("privatePostOrderOrdersIdSubmitcancel"), MkMap(&VarMap{"id": id}))
	_ = response
	return this.Extend(this.ParseOrder(response), MkMap(&VarMap{
		"id":     id,
		"status": MkString("canceled"),
	}))
}

func (this *Huobi) CancelOrders(goArgs ...*Variant) *Variant {
	ids := GoGetArg(goArgs, 0, MkUndefined())
	_ = ids
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	clientOrderIds := this.SafeValue2(params, MkString("clientOrderIds"), MkString("client-order-ids"))
	_ = clientOrderIds
	params = this.Omit(params, MkArray(&VarArray{
		MkString("clientOrderIds"),
		MkString("client-order-ids"),
	}))
	request := MkMap(&VarMap{})
	_ = request
	if IsTrue(OpEq2(clientOrderIds, MkUndefined())) {
		*(request).At(MkString("order-ids")) = OpCopy(ids)
	} else {
		*(request).At(MkString("client-order-ids")) = OpCopy(clientOrderIds)
	}
	response := this.Call(MkString("privatePostOrderOrdersBatchcancel"), this.Extend(request, params))
	_ = response
	return response
}

func (this *Huobi) CancelAllOrders(goArgs ...*Variant) *Variant {
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
		*(request).At(MkString("symbol")) = *(market).At(MkString("id"))
	}
	response := this.Call(MkString("privatePostOrderOrdersBatchCancelOpenOrders"), this.Extend(request, params))
	_ = response
	return response
}

func (this *Huobi) CurrencyToPrecision(goArgs ...*Variant) *Variant {
	currency := GoGetArg(goArgs, 0, MkUndefined())
	_ = currency
	fee := GoGetArg(goArgs, 1, MkUndefined())
	_ = fee
	return this.DecimalToPrecision(fee, MkInteger(0), *(*(*this.At(MkString("currencies"))).At(currency)).At(MkString("precision")))
}

func (this *Huobi) ParseDepositAddress(goArgs ...*Variant) *Variant {
	depositAddress := GoGetArg(goArgs, 0, MkUndefined())
	_ = depositAddress
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	address := this.SafeString(depositAddress, MkString("address"))
	_ = address
	tag := this.SafeString(depositAddress, MkString("addressTag"))
	_ = tag
	currencyId := this.SafeString(depositAddress, MkString("currency"))
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

func (this *Huobi) FetchDepositAddress(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{"currency": *(currency).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("v2PrivateGetAccountDepositAddress"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseDepositAddress(this.SafeValue(data, MkInteger(0), MkMap(&VarMap{})), currency)
}

func (this *Huobi) FetchDeposits(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpOr(OpEq2(limit, MkUndefined()), OpGt(limit, MkInteger(100)))) {
		limit = MkInteger(100)
	}
	this.LoadMarkets()
	currency := OpCopy(MkUndefined())
	_ = currency
	if IsTrue(OpNotEq2(code, MkUndefined())) {
		currency = this.Currency(code)
	}
	request := MkMap(&VarMap{
		"type": MkString("deposit"),
		"from": MkInteger(0),
	})
	_ = request
	if IsTrue(OpNotEq2(currency, MkUndefined())) {
		*(request).At(MkString("currency")) = *(currency).At(MkString("id"))
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("size")) = OpCopy(limit)
	}
	response := this.Call(MkString("privateGetQueryDepositWithdraw"), this.Extend(request, params))
	_ = response
	return this.ParseTransactions(*(response).At(MkString("data")), currency, since, limit)
}

func (this *Huobi) FetchWithdrawals(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpOr(OpEq2(limit, MkUndefined()), OpGt(limit, MkInteger(100)))) {
		limit = MkInteger(100)
	}
	this.LoadMarkets()
	currency := OpCopy(MkUndefined())
	_ = currency
	if IsTrue(OpNotEq2(code, MkUndefined())) {
		currency = this.Currency(code)
	}
	request := MkMap(&VarMap{
		"type": MkString("withdraw"),
		"from": MkInteger(0),
	})
	_ = request
	if IsTrue(OpNotEq2(currency, MkUndefined())) {
		*(request).At(MkString("currency")) = *(currency).At(MkString("id"))
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("size")) = OpCopy(limit)
	}
	response := this.Call(MkString("privateGetQueryDepositWithdraw"), this.Extend(request, params))
	_ = response
	return this.ParseTransactions(*(response).At(MkString("data")), currency, since, limit)
}

func (this *Huobi) ParseTransaction(goArgs ...*Variant) *Variant {
	transaction := GoGetArg(goArgs, 0, MkUndefined())
	_ = transaction
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	timestamp := this.SafeInteger(transaction, MkString("created-at"))
	_ = timestamp
	updated := this.SafeInteger(transaction, MkString("updated-at"))
	_ = updated
	code := this.SafeCurrencyCode(this.SafeString(transaction, MkString("currency")))
	_ = code
	type_ := this.SafeString(transaction, MkString("type"))
	_ = type_
	if IsTrue(OpEq2(type_, MkString("withdraw"))) {
		type_ = MkString("withdrawal")
	}
	status := this.ParseTransactionStatus(this.SafeString(transaction, MkString("state")))
	_ = status
	tag := this.SafeString(transaction, MkString("address-tag"))
	_ = tag
	feeCost := this.SafeNumber(transaction, MkString("fee"))
	_ = feeCost
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		feeCost = Math.Abs(feeCost)
	}
	return MkMap(&VarMap{
		"info":      transaction,
		"id":        this.SafeString(transaction, MkString("id")),
		"txid":      this.SafeString(transaction, MkString("tx-hash")),
		"timestamp": timestamp,
		"datetime":  this.Iso8601(timestamp),
		"address":   this.SafeString(transaction, MkString("address")),
		"tag":       tag,
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

func (this *Huobi) ParseTransactionStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"unknown":         MkString("failed"),
		"confirming":      MkString("pending"),
		"confirmed":       MkString("ok"),
		"safe":            MkString("ok"),
		"orphan":          MkString("failed"),
		"submitted":       MkString("pending"),
		"canceled":        MkString("canceled"),
		"reexamine":       MkString("pending"),
		"reject":          MkString("failed"),
		"pass":            MkString("pending"),
		"wallet-reject":   MkString("failed"),
		"confirm-error":   MkString("failed"),
		"repealed":        MkString("failed"),
		"wallet-transfer": MkString("pending"),
		"pre-transfer":    MkString("pending"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Huobi) Withdraw(goArgs ...*Variant) *Variant {
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
	this.CheckAddress(address)
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{
		"address":  address,
		"amount":   amount,
		"currency": (*(currency).At(MkString("id"))).Call(MkString("toLowerCase")),
	})
	_ = request
	if IsTrue(OpNotEq2(tag, MkUndefined())) {
		*(request).At(MkString("addr-tag")) = OpCopy(tag)
	}
	response := this.Call(MkString("privatePostDwWithdrawApiCreate"), this.Extend(request, params))
	_ = response
	id := this.SafeString(response, MkString("data"))
	_ = id
	return MkMap(&VarMap{
		"info": response,
		"id":   id,
	})
}

func (this *Huobi) Sign(goArgs ...*Variant) *Variant {
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
	url := MkString("/")
	_ = url
	if IsTrue(OpEq2(api, MkString("market"))) {
		OpAddAssign(&url, api)
	} else {
		if IsTrue(OpOr(OpEq2(api, MkString("public")), OpEq2(api, MkString("private")))) {
			OpAddAssign(&url, *this.At(MkString("version")))
		} else {
			if IsTrue(OpOr(OpEq2(api, MkString("v2Public")), OpEq2(api, MkString("v2Private")))) {
				OpAddAssign(&url, MkString("v2"))
			}
		}
	}
	OpAddAssign(&url, OpAdd(MkString("/"), this.ImplodeParams(path, params)))
	query := this.Omit(params, this.ExtractParams(path))
	_ = query
	if IsTrue(OpOr(OpEq2(api, MkString("private")), OpEq2(api, MkString("v2Private")))) {
		this.CheckRequiredCredentials()
		timestamp := this.Ymdhms(this.Milliseconds(), MkString("T"))
		_ = timestamp
		request := MkMap(&VarMap{
			"SignatureMethod":  MkString("HmacSHA256"),
			"SignatureVersion": MkString("2"),
			"AccessKeyId":      *this.At(MkString("apiKey")),
			"Timestamp":        timestamp,
		})
		_ = request
		if IsTrue(OpNotEq2(method, MkString("POST"))) {
			request = this.Extend(request, query)
		}
		request = this.Keysort(request)
		auth := this.Urlencode(request)
		_ = auth
		payload := (MkArray(&VarArray{
			method,
			*this.At(MkString("hostname")),
			url,
			auth,
		})).Call(MkString("join"), MkString("\n"))
		_ = payload
		signature := this.Hmac(this.Encode(payload), this.Encode(*this.At(MkString("secret"))), MkString("sha256"), MkString("base64"))
		_ = signature
		OpAddAssign(&auth, OpAdd(MkString("&"), this.Urlencode(MkMap(&VarMap{"Signature": signature}))))
		OpAddAssign(&url, OpAdd(MkString("?"), auth))
		if IsTrue(OpEq2(method, MkString("POST"))) {
			body = this.Json(query)
			headers = MkMap(&VarMap{"Content-Type": MkString("application/json")})
		} else {
			headers = MkMap(&VarMap{"Content-Type": MkString("application/x-www-form-urlencoded")})
		}
	} else {
		if IsTrue(*(GoGetKeys(params)).At(MkString("length"))) {
			OpAddAssign(&url, OpAdd(MkString("?"), this.Urlencode(params)))
		}
	}
	url = OpAdd(this.ImplodeParams(*(*(*this.At(MkString("urls"))).At(MkString("api"))).At(api), MkMap(&VarMap{"hostname": *this.At(MkString("hostname"))})), url)
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Huobi) HandleErrors(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpHasMember(MkString("status"), response)) {
		status := this.SafeString(response, MkString("status"))
		_ = status
		if IsTrue(OpEq2(status, MkString("error"))) {
			code := this.SafeString(response, MkString("err-code"))
			_ = code
			feedback := OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))
			_ = feedback
			this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("broad")), body, feedback)
			this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("exact")), code, feedback)
			message := this.SafeString(response, MkString("err-msg"))
			_ = message
			this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("exact")), message, feedback)
			panic(NewExchangeError(feedback))
		}
	}
	return MkUndefined()
}
