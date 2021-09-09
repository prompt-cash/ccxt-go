package ccxt

import ()

type Coinbasepro struct {
	*ExchangeBase
}

var _ Exchange = (*Coinbasepro)(nil)

func init() {
	exchange := &Coinbasepro{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Coinbasepro) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("coinbasepro"),
		"name": MkString("Coinbase Pro"),
		"countries": MkArray(&VarArray{
			MkString("US"),
		}),
		"rateLimit": MkInteger(1000),
		"userAgent": *(*this.At(MkString("userAgents"))).At(MkString("chrome")),
		"pro":       MkBool(true),
		"has": MkMap(&VarMap{
			"cancelAllOrders":      MkBool(true),
			"cancelOrder":          MkBool(true),
			"CORS":                 MkBool(true),
			"createDepositAddress": MkBool(true),
			"createOrder":          MkBool(true),
			"deposit":              MkBool(true),
			"fetchAccounts":        MkBool(true),
			"fetchBalance":         MkBool(true),
			"fetchCurrencies":      MkBool(true),
			"fetchClosedOrders":    MkBool(true),
			"fetchDepositAddress":  MkBool(false),
			"fetchMarkets":         MkBool(true),
			"fetchMyTrades":        MkBool(true),
			"fetchOHLCV":           MkBool(true),
			"fetchOpenOrders":      MkBool(true),
			"fetchOrder":           MkBool(true),
			"fetchOrderBook":       MkBool(true),
			"fetchOrders":          MkBool(true),
			"fetchOrderTrades":     MkBool(true),
			"fetchTime":            MkBool(true),
			"fetchTicker":          MkBool(true),
			"fetchTrades":          MkBool(true),
			"fetchTransactions":    MkBool(true),
			"withdraw":             MkBool(true),
			"fetchDeposits":        MkBool(true),
			"fetchWithdrawals":     MkBool(true),
		}),
		"timeframes": MkMap(&VarMap{
			"1m":  MkInteger(60),
			"5m":  MkInteger(300),
			"15m": MkInteger(900),
			"1h":  MkInteger(3600),
			"6h":  MkInteger(21600),
			"1d":  MkInteger(86400),
		}),
		"hostname": MkString("pro.coinbase.com"),
		"urls": MkMap(&VarMap{
			"test": MkMap(&VarMap{
				"public":  MkString("https://api-public.sandbox.pro.coinbase.com"),
				"private": MkString("https://api-public.sandbox.pro.coinbase.com"),
			}),
			"logo": MkString("https://user-images.githubusercontent.com/1294454/41764625-63b7ffde-760a-11e8-996d-a6328fa9347a.jpg"),
			"api": MkMap(&VarMap{
				"public":  MkString("https://api.{hostname}"),
				"private": MkString("https://api.{hostname}"),
			}),
			"www": MkString("https://pro.coinbase.com/"),
			"doc": MkString("https://docs.pro.coinbase.com"),
			"fees": MkArray(&VarArray{
				MkString("https://docs.pro.coinbase.com/#fees"),
				MkString("https://support.pro.coinbase.com/customer/en/portal/articles/2945310-fees"),
			}),
		}),
		"requiredCredentials": MkMap(&VarMap{
			"apiKey":   MkBool(true),
			"secret":   MkBool(true),
			"password": MkBool(true),
		}),
		"api": MkMap(&VarMap{
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("currencies"),
				MkString("products"),
				MkString("products/{id}"),
				MkString("products/{id}/book"),
				MkString("products/{id}/candles"),
				MkString("products/{id}/stats"),
				MkString("products/{id}/ticker"),
				MkString("products/{id}/trades"),
				MkString("time"),
			})}),
			"private": MkMap(&VarMap{
				"get": MkArray(&VarArray{
					MkString("accounts"),
					MkString("accounts/{id}"),
					MkString("accounts/{id}/holds"),
					MkString("accounts/{id}/ledger"),
					MkString("accounts/{id}/transfers"),
					MkString("coinbase-accounts"),
					MkString("fills"),
					MkString("funding"),
					MkString("fees"),
					MkString("margin/profile_information"),
					MkString("margin/buying_power"),
					MkString("margin/withdrawal_power"),
					MkString("margin/withdrawal_power_all"),
					MkString("margin/exit_plan"),
					MkString("margin/liquidation_history"),
					MkString("margin/position_refresh_amounts"),
					MkString("margin/status"),
					MkString("oracle"),
					MkString("orders"),
					MkString("orders/{id}"),
					MkString("orders/client:{client_oid}"),
					MkString("otc/orders"),
					MkString("payment-methods"),
					MkString("position"),
					MkString("profiles"),
					MkString("profiles/{id}"),
					MkString("reports/{report_id}"),
					MkString("transfers"),
					MkString("transfers/{transfer_id}"),
					MkString("users/self/exchange-limits"),
					MkString("users/self/hold-balances"),
					MkString("users/self/trailing-volume"),
					MkString("withdrawals/fee-estimate"),
				}),
				"post": MkArray(&VarArray{
					MkString("conversions"),
					MkString("deposits/coinbase-account"),
					MkString("deposits/payment-method"),
					MkString("coinbase-accounts/{id}/addresses"),
					MkString("funding/repay"),
					MkString("orders"),
					MkString("position/close"),
					MkString("profiles/margin-transfer"),
					MkString("profiles/transfer"),
					MkString("reports"),
					MkString("withdrawals/coinbase"),
					MkString("withdrawals/coinbase-account"),
					MkString("withdrawals/crypto"),
					MkString("withdrawals/payment-method"),
				}),
				"delete": MkArray(&VarArray{
					MkString("orders"),
					MkString("orders/client:{client_oid}"),
					MkString("orders/{id}"),
				}),
			}),
		}),
		"commonCurrencies": MkMap(&VarMap{"CGLD": MkString("CELO")}),
		"precisionMode":    TICK_SIZE,
		"fees": MkMap(&VarMap{
			"trading": MkMap(&VarMap{
				"tierBased":  MkBool(true),
				"percentage": MkBool(true),
				"maker":      OpDiv(MkNumber(0.5), MkInteger(100)),
				"taker":      OpDiv(MkNumber(0.5), MkInteger(100)),
			}),
			"funding": MkMap(&VarMap{
				"tierBased":  MkBool(false),
				"percentage": MkBool(false),
				"withdraw": MkMap(&VarMap{
					"BCH": MkInteger(0),
					"BTC": MkInteger(0),
					"LTC": MkInteger(0),
					"ETH": MkInteger(0),
					"EUR": MkNumber(0.15),
					"USD": MkInteger(25),
				}),
				"deposit": MkMap(&VarMap{
					"BCH": MkInteger(0),
					"BTC": MkInteger(0),
					"LTC": MkInteger(0),
					"ETH": MkInteger(0),
					"EUR": MkNumber(0.15),
					"USD": MkInteger(10),
				}),
			}),
		}),
		"exceptions": MkMap(&VarMap{
			"exact": MkMap(&VarMap{
				"Insufficient funds":          InsufficientFunds,
				"NotFound":                    OrderNotFound,
				"Invalid API Key":             AuthenticationError,
				"invalid signature":           AuthenticationError,
				"Invalid Passphrase":          AuthenticationError,
				"Invalid order id":            InvalidOrder,
				"Private rate limit exceeded": RateLimitExceeded,
				"Trading pair not available":  PermissionDenied,
				"Product not found":           InvalidOrder,
			}),
			"broad": MkMap(&VarMap{
				"Order already done": OrderNotFound,
				"order not found":    OrderNotFound,
				"price too small":    InvalidOrder,
				"price too precise":  InvalidOrder,
				"under maintenance":  OnMaintenance,
				"size is too small":  InvalidOrder,
				"Cancel only mode":   OnMaintenance,
			}),
		}),
	}))
}

func (this *Coinbasepro) FetchCurrencies(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetCurrencies"), params)
	_ = response
	result := MkMap(&VarMap{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
		currency := *(response).At(i)
		_ = currency
		id := this.SafeString(currency, MkString("id"))
		_ = id
		name := this.SafeString(currency, MkString("name"))
		_ = name
		code := this.SafeCurrencyCode(id)
		_ = code
		details := this.SafeValue(currency, MkString("details"), MkMap(&VarMap{}))
		_ = details
		precision := this.SafeNumber(currency, MkString("max_precision"))
		_ = precision
		status := this.SafeString(currency, MkString("status"))
		_ = status
		active := OpEq2(status, MkString("online"))
		_ = active
		*(result).At(code) = MkMap(&VarMap{
			"id":        id,
			"code":      code,
			"info":      currency,
			"type":      this.SafeString(details, MkString("type")),
			"name":      name,
			"active":    active,
			"fee":       MkUndefined(),
			"precision": precision,
			"limits": MkMap(&VarMap{
				"amount": MkMap(&VarMap{
					"min": this.SafeNumber(details, MkString("min_size")),
					"max": MkUndefined(),
				}),
				"withdraw": MkMap(&VarMap{
					"min": this.SafeNumber(details, MkString("min_withdrawal_amount")),
					"max": MkUndefined(),
				}),
			}),
		})
	}
	return result
}

func (this *Coinbasepro) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetProducts"), params)
	_ = response
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
		market := *(response).At(i)
		_ = market
		id := this.SafeString(market, MkString("id"))
		_ = id
		baseId := this.SafeString(market, MkString("base_currency"))
		_ = baseId
		quoteId := this.SafeString(market, MkString("quote_currency"))
		_ = quoteId
		base := this.SafeCurrencyCode(baseId)
		_ = base
		quote := this.SafeCurrencyCode(quoteId)
		_ = quote
		symbol := OpAdd(base, OpAdd(MkString("/"), quote))
		_ = symbol
		priceLimits := MkMap(&VarMap{
			"min": this.SafeNumber(market, MkString("quote_increment")),
			"max": MkUndefined(),
		})
		_ = priceLimits
		precision := MkMap(&VarMap{
			"amount": this.SafeNumber(market, MkString("base_increment")),
			"price":  this.SafeNumber(market, MkString("quote_increment")),
		})
		_ = precision
		status := this.SafeString(market, MkString("status"))
		_ = status
		active := OpEq2(status, MkString("online"))
		_ = active
		result.Push(this.Extend(*(*this.At(MkString("fees"))).At(MkString("trading")), MkMap(&VarMap{
			"id":        id,
			"symbol":    symbol,
			"baseId":    baseId,
			"quoteId":   quoteId,
			"base":      base,
			"quote":     quote,
			"precision": precision,
			"limits": MkMap(&VarMap{
				"amount": MkMap(&VarMap{
					"min": this.SafeNumber(market, MkString("base_min_size")),
					"max": this.SafeNumber(market, MkString("base_max_size")),
				}),
				"price": priceLimits,
				"cost": MkMap(&VarMap{
					"min": this.SafeNumber(market, MkString("min_market_funds")),
					"max": this.SafeNumber(market, MkString("max_market_funds")),
				}),
			}),
			"active": active,
			"info":   market,
		})))
	}
	return result
}

func (this *Coinbasepro) FetchAccounts(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privateGetAccounts"), params)
	_ = response
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
		account := *(response).At(i)
		_ = account
		accountId := this.SafeString(account, MkString("id"))
		_ = accountId
		currencyId := this.SafeString(account, MkString("currency"))
		_ = currencyId
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		result.Push(MkMap(&VarMap{
			"id":       accountId,
			"type":     MkUndefined(),
			"currency": code,
			"info":     account,
		}))
	}
	return result
}

func (this *Coinbasepro) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privateGetAccounts"), params)
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
		*(account).At(MkString("used")) = this.SafeString(balance, MkString("hold"))
		*(account).At(MkString("total")) = this.SafeString(balance, MkString("balance"))
		*(result).At(code) = OpCopy(account)
	}
	return this.ParseBalance(result)
}

func (this *Coinbasepro) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{
		"id":    this.MarketId(symbol),
		"level": MkInteger(2),
	})
	_ = request
	response := this.Call(MkString("publicGetProductsIdBook"), this.Extend(request, params))
	_ = response
	orderbook := this.ParseOrderBook(response, symbol)
	_ = orderbook
	*(orderbook).At(MkString("nonce")) = this.SafeInteger(response, MkString("sequence"))
	return orderbook
}

func (this *Coinbasepro) ParseTicker(goArgs ...*Variant) *Variant {
	ticker := GoGetArg(goArgs, 0, MkUndefined())
	_ = ticker
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.Parse8601(this.SafeValue(ticker, MkString("time")))
	_ = timestamp
	bid := this.SafeNumber(ticker, MkString("bid"))
	_ = bid
	ask := this.SafeNumber(ticker, MkString("ask"))
	_ = ask
	last := this.SafeNumber(ticker, MkString("price"))
	_ = last
	symbol := OpTrinary(OpEq2(market, MkUndefined()), MkUndefined(), *(market).At(MkString("symbol")))
	_ = symbol
	return MkMap(&VarMap{
		"symbol":        symbol,
		"timestamp":     timestamp,
		"datetime":      this.Iso8601(timestamp),
		"high":          this.SafeNumber(ticker, MkString("high")),
		"low":           this.SafeNumber(ticker, MkString("low")),
		"bid":           bid,
		"bidVolume":     MkUndefined(),
		"ask":           ask,
		"askVolume":     MkUndefined(),
		"vwap":          MkUndefined(),
		"open":          this.SafeNumber(ticker, MkString("open")),
		"close":         last,
		"last":          last,
		"previousClose": MkUndefined(),
		"change":        MkUndefined(),
		"percentage":    MkUndefined(),
		"average":       MkUndefined(),
		"baseVolume":    this.SafeNumber(ticker, MkString("volume")),
		"quoteVolume":   MkUndefined(),
		"info":          ticker,
	})
}

func (this *Coinbasepro) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"id": *(market).At(MkString("id"))})
	_ = request
	method := this.SafeString(*this.At(MkString("options")), MkString("fetchTickerMethod"), MkString("publicGetProductsIdTicker"))
	_ = method
	response := this.Call(method, this.Extend(request, params))
	_ = response
	return this.ParseTicker(response, market)
}

func (this *Coinbasepro) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.Parse8601(this.SafeString2(trade, MkString("time"), MkString("created_at")))
	_ = timestamp
	marketId := this.SafeString(trade, MkString("product_id"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market, MkString("-"))
	_ = symbol
	feeRate := OpCopy(MkUndefined())
	_ = feeRate
	feeCurrency := OpCopy(MkUndefined())
	_ = feeCurrency
	takerOrMaker := OpCopy(MkUndefined())
	_ = takerOrMaker
	cost := OpCopy(MkUndefined())
	_ = cost
	if IsTrue(OpNotEq2(market, MkUndefined())) {
		feeCurrencyId := this.SafeStringLower(market, MkString("quoteId"))
		_ = feeCurrencyId
		costField := OpAdd(feeCurrencyId, MkString("_value"))
		_ = costField
		cost = this.SafeNumber(trade, costField)
		feeCurrency = *(market).At(MkString("quote"))
		liquidity := this.SafeString(trade, MkString("liquidity"))
		_ = liquidity
		if IsTrue(OpNotEq2(liquidity, MkUndefined())) {
			takerOrMaker = OpTrinary(OpEq2(liquidity, MkString("T")), MkString("taker"), MkString("maker"))
			feeRate = *(market).At(takerOrMaker)
		}
	}
	feeCost := this.SafeNumber2(trade, MkString("fill_fees"), MkString("fee"))
	_ = feeCost
	fee := MkMap(&VarMap{
		"cost":     feeCost,
		"currency": feeCurrency,
		"rate":     feeRate,
	})
	_ = fee
	type_ := OpCopy(MkUndefined())
	_ = type_
	id := this.SafeString(trade, MkString("trade_id"))
	_ = id
	side := OpTrinary(OpEq2(*(trade).At(MkString("side")), MkString("buy")), MkString("sell"), MkString("buy"))
	_ = side
	orderId := this.SafeString(trade, MkString("order_id"))
	_ = orderId
	makerOrderId := this.SafeString(trade, MkString("maker_order_id"))
	_ = makerOrderId
	takerOrderId := this.SafeString(trade, MkString("taker_order_id"))
	_ = takerOrderId
	if IsTrue(OpOr(OpNotEq2(orderId, MkUndefined()), OpAnd(OpNotEq2(makerOrderId, MkUndefined()), OpNotEq2(takerOrderId, MkUndefined())))) {
		side = OpTrinary(OpEq2(*(trade).At(MkString("side")), MkString("buy")), MkString("buy"), MkString("sell"))
	}
	priceString := this.SafeString(trade, MkString("price"))
	_ = priceString
	amountString := this.SafeString(trade, MkString("size"))
	_ = amountString
	price := this.ParseNumber(priceString)
	_ = price
	amount := this.ParseNumber(amountString)
	_ = amount
	if IsTrue(OpEq2(cost, MkUndefined())) {
		cost = this.ParseNumber(Precise.StringMul(priceString, amountString))
	}
	return MkMap(&VarMap{
		"id":           id,
		"order":        orderId,
		"info":         trade,
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       symbol,
		"type":         type_,
		"takerOrMaker": takerOrMaker,
		"side":         side,
		"price":        price,
		"amount":       amount,
		"fee":          fee,
		"cost":         cost,
	})
}

func (this *Coinbasepro) FetchMyTrades(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"product_id": *(market).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("privateGetFills"), this.Extend(request, params))
	_ = response
	return this.ParseTrades(response, market, since, limit)
}

func (this *Coinbasepro) FetchTrades(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"id": *(market).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("publicGetProductsIdTrades"), this.Extend(request, params))
	_ = response
	return this.ParseTrades(response, market, since, limit)
}

func (this *Coinbasepro) ParseOHLCV(goArgs ...*Variant) *Variant {
	ohlcv := GoGetArg(goArgs, 0, MkUndefined())
	_ = ohlcv
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	return MkArray(&VarArray{
		this.SafeTimestamp(ohlcv, MkInteger(0)),
		this.SafeNumber(ohlcv, MkInteger(3)),
		this.SafeNumber(ohlcv, MkInteger(2)),
		this.SafeNumber(ohlcv, MkInteger(1)),
		this.SafeNumber(ohlcv, MkInteger(4)),
		this.SafeNumber(ohlcv, MkInteger(5)),
	})
}

func (this *Coinbasepro) FetchOHLCV(goArgs ...*Variant) *Variant {
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
	granularity := *(*this.At(MkString("timeframes"))).At(timeframe)
	_ = granularity
	request := MkMap(&VarMap{
		"id":          *(market).At(MkString("id")),
		"granularity": granularity,
	})
	_ = request
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("start")) = this.Iso8601(since)
		if IsTrue(OpEq2(limit, MkUndefined())) {
			limit = MkInteger(300)
		} else {
			limit = Math.Min(MkInteger(300), limit)
		}
		*(request).At(MkString("end")) = this.Iso8601(this.Sum(OpMulti(OpSub(limit, MkInteger(1)), OpMulti(granularity, MkInteger(1000))), since))
	}
	response := this.Call(MkString("publicGetProductsIdCandles"), this.Extend(request, params))
	_ = response
	return this.ParseOHLCVs(response, market, timeframe, since, limit)
}

func (this *Coinbasepro) FetchTime(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetTime"), params)
	_ = response
	return this.SafeTimestamp(response, MkString("epoch"))
}

func (this *Coinbasepro) ParseOrderStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"pending":   MkString("open"),
		"active":    MkString("open"),
		"open":      MkString("open"),
		"done":      MkString("closed"),
		"canceled":  MkString("canceled"),
		"canceling": MkString("open"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Coinbasepro) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.Parse8601(this.SafeString(order, MkString("created_at")))
	_ = timestamp
	marketId := this.SafeString(order, MkString("product_id"))
	_ = marketId
	market = this.SafeMarket(marketId, market, MkString("-"))
	status := this.ParseOrderStatus(this.SafeString(order, MkString("status")))
	_ = status
	price := this.SafeNumber(order, MkString("price"))
	_ = price
	filled := this.SafeNumber(order, MkString("filled_size"))
	_ = filled
	amount := this.SafeNumber(order, MkString("size"), filled)
	_ = amount
	cost := this.SafeNumber(order, MkString("executed_value"))
	_ = cost
	feeCost := this.SafeNumber(order, MkString("fill_fees"))
	_ = feeCost
	fee := OpCopy(MkUndefined())
	_ = fee
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		feeCurrencyCode := OpCopy(MkUndefined())
		_ = feeCurrencyCode
		if IsTrue(OpNotEq2(market, MkUndefined())) {
			feeCurrencyCode = *(market).At(MkString("quote"))
		}
		fee = MkMap(&VarMap{
			"cost":     feeCost,
			"currency": feeCurrencyCode,
			"rate":     MkUndefined(),
		})
	}
	id := this.SafeString(order, MkString("id"))
	_ = id
	type_ := this.SafeString(order, MkString("type"))
	_ = type_
	side := this.SafeString(order, MkString("side"))
	_ = side
	timeInForce := this.SafeString(order, MkString("time_in_force"))
	_ = timeInForce
	postOnly := this.SafeValue(order, MkString("post_only"))
	_ = postOnly
	stopPrice := this.SafeNumber(order, MkString("stop_price"))
	_ = stopPrice
	clientOrderId := this.SafeString(order, MkString("client_oid"))
	_ = clientOrderId
	return this.SafeOrder(MkMap(&VarMap{
		"id":                 id,
		"clientOrderId":      clientOrderId,
		"info":               order,
		"timestamp":          timestamp,
		"datetime":           this.Iso8601(timestamp),
		"lastTradeTimestamp": MkUndefined(),
		"status":             status,
		"symbol":             *(market).At(MkString("symbol")),
		"type":               type_,
		"timeInForce":        timeInForce,
		"postOnly":           postOnly,
		"side":               side,
		"price":              price,
		"stopPrice":          stopPrice,
		"cost":               cost,
		"amount":             amount,
		"filled":             filled,
		"remaining":          MkUndefined(),
		"fee":                fee,
		"average":            MkUndefined(),
		"trades":             MkUndefined(),
	}))
}

func (this *Coinbasepro) FetchOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{})
	_ = request
	clientOrderId := this.SafeString2(params, MkString("clientOrderId"), MkString("client_oid"))
	_ = clientOrderId
	method := OpCopy(MkUndefined())
	_ = method
	if IsTrue(OpEq2(clientOrderId, MkUndefined())) {
		method = MkString("privateGetOrdersId")
		*(request).At(MkString("id")) = OpCopy(id)
	} else {
		method = MkString("privateGetOrdersClientClientOid")
		*(request).At(MkString("client_oid")) = OpCopy(clientOrderId)
		params = this.Omit(params, MkArray(&VarArray{
			MkString("clientOrderId"),
			MkString("client_oid"),
		}))
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	return this.ParseOrder(response)
}

func (this *Coinbasepro) FetchOrderTrades(goArgs ...*Variant) *Variant {
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
	market := OpCopy(MkUndefined())
	_ = market
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
	}
	request := MkMap(&VarMap{"order_id": id})
	_ = request
	response := this.Call(MkString("privateGetFills"), this.Extend(request, params))
	_ = response
	return this.ParseTrades(response, market, since, limit)
}

func (this *Coinbasepro) FetchOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	request := MkMap(&VarMap{"status": MkString("all")})
	_ = request
	return this.FetchOpenOrders(symbol, since, limit, this.Extend(request, params))
}

func (this *Coinbasepro) FetchOpenOrders(goArgs ...*Variant) *Variant {
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
		*(request).At(MkString("product_id")) = *(market).At(MkString("id"))
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("privateGetOrders"), this.Extend(request, params))
	_ = response
	return this.ParseOrders(response, market, since, limit)
}

func (this *Coinbasepro) FetchClosedOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	request := MkMap(&VarMap{"status": MkString("done")})
	_ = request
	return this.FetchOpenOrders(symbol, since, limit, this.Extend(request, params))
}

func (this *Coinbasepro) CreateOrder(goArgs ...*Variant) *Variant {
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
		"type":       type_,
		"side":       side,
		"product_id": *(market).At(MkString("id")),
	})
	_ = request
	clientOrderId := this.SafeString2(params, MkString("clientOrderId"), MkString("client_oid"))
	_ = clientOrderId
	if IsTrue(OpNotEq2(clientOrderId, MkUndefined())) {
		*(request).At(MkString("client_oid")) = OpCopy(clientOrderId)
		params = this.Omit(params, MkArray(&VarArray{
			MkString("clientOrderId"),
			MkString("client_oid"),
		}))
	}
	stopPrice := this.SafeNumber2(params, MkString("stopPrice"), MkString("stop_price"))
	_ = stopPrice
	if IsTrue(OpNotEq2(stopPrice, MkUndefined())) {
		*(request).At(MkString("stop_price")) = this.PriceToPrecision(symbol, stopPrice)
		params = this.Omit(params, MkArray(&VarArray{
			MkString("stopPrice"),
			MkString("stop_price"),
		}))
	}
	timeInForce := this.SafeString2(params, MkString("timeInForce"), MkString("time_in_force"))
	_ = timeInForce
	if IsTrue(OpNotEq2(timeInForce, MkUndefined())) {
		*(request).At(MkString("time_in_force")) = OpCopy(timeInForce)
		params = this.Omit(params, MkArray(&VarArray{
			MkString("timeInForce"),
			MkString("time_in_force"),
		}))
	}
	if IsTrue(OpEq2(type_, MkString("limit"))) {
		*(request).At(MkString("price")) = this.PriceToPrecision(symbol, price)
		*(request).At(MkString("size")) = this.AmountToPrecision(symbol, amount)
	} else {
		if IsTrue(OpEq2(type_, MkString("market"))) {
			cost := this.SafeNumber2(params, MkString("cost"), MkString("funds"))
			_ = cost
			if IsTrue(OpEq2(cost, MkUndefined())) {
				if IsTrue(OpNotEq2(price, MkUndefined())) {
					cost = OpMulti(amount, price)
				}
			} else {
				params = this.Omit(params, MkArray(&VarArray{
					MkString("cost"),
					MkString("funds"),
				}))
			}
			if IsTrue(OpNotEq2(cost, MkUndefined())) {
				*(request).At(MkString("funds")) = this.CostToPrecision(symbol, cost)
			} else {
				*(request).At(MkString("size")) = this.AmountToPrecision(symbol, amount)
			}
		}
	}
	response := this.Call(MkString("privatePostOrders"), this.Extend(request, params))
	_ = response
	return this.ParseOrder(response, market)
}

func (this *Coinbasepro) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{})
	_ = request
	clientOrderId := this.SafeString2(params, MkString("clientOrderId"), MkString("client_oid"))
	_ = clientOrderId
	method := OpCopy(MkUndefined())
	_ = method
	if IsTrue(OpEq2(clientOrderId, MkUndefined())) {
		method = MkString("privateDeleteOrdersId")
		*(request).At(MkString("id")) = OpCopy(id)
	} else {
		method = MkString("privateDeleteOrdersClientClientOid")
		*(request).At(MkString("client_oid")) = OpCopy(clientOrderId)
		params = this.Omit(params, MkArray(&VarArray{
			MkString("clientOrderId"),
			MkString("client_oid"),
		}))
	}
	market := OpCopy(MkUndefined())
	_ = market
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
		*(request).At(MkString("product_id")) = *(market).At(MkString("symbol"))
	}
	return this.Call(method, this.Extend(request, params))
}

func (this *Coinbasepro) CancelAllOrders(goArgs ...*Variant) *Variant {
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
		*(request).At(MkString("product_id")) = *(market).At(MkString("symbol"))
	}
	return this.Call(MkString("privateDeleteOrders"), this.Extend(request, params))
}

func (this *Coinbasepro) FetchPaymentMethods(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	return this.Call(MkString("privateGetPaymentMethods"), params)
}

func (this *Coinbasepro) Deposit(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	amount := GoGetArg(goArgs, 1, MkUndefined())
	_ = amount
	address := GoGetArg(goArgs, 2, MkUndefined())
	_ = address
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{
		"currency": *(currency).At(MkString("id")),
		"amount":   amount,
	})
	_ = request
	method := MkString("privatePostDeposits")
	_ = method
	if IsTrue(OpHasMember(MkString("payment_method_id"), params)) {
		OpAddAssign(&method, MkString("PaymentMethod"))
	} else {
		if IsTrue(OpHasMember(MkString("coinbase_account_id"), params)) {
			OpAddAssign(&method, MkString("CoinbaseAccount"))
		} else {
			panic(NewNotSupported(OpAdd(*this.At(MkString("id")), MkString(" deposit() requires one of `coinbase_account_id` or `payment_method_id` extra params"))))
		}
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	if IsTrue(OpNot(response)) {
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" deposit() error: "), this.Json(response)))))
	}
	return MkMap(&VarMap{
		"info": response,
		"id":   *(response).At(MkString("id")),
	})
}

func (this *Coinbasepro) Withdraw(goArgs ...*Variant) *Variant {
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
	})
	_ = request
	method := MkString("privatePostWithdrawals")
	_ = method
	if IsTrue(OpHasMember(MkString("payment_method_id"), params)) {
		OpAddAssign(&method, MkString("PaymentMethod"))
	} else {
		if IsTrue(OpHasMember(MkString("coinbase_account_id"), params)) {
			OpAddAssign(&method, MkString("CoinbaseAccount"))
		} else {
			OpAddAssign(&method, MkString("Crypto"))
			*(request).At(MkString("crypto_address")) = OpCopy(address)
			if IsTrue(OpNotEq2(tag, MkUndefined())) {
				*(request).At(MkString("destination_tag")) = OpCopy(tag)
			}
		}
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	if IsTrue(OpNot(response)) {
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" withdraw() error: "), this.Json(response)))))
	}
	return MkMap(&VarMap{
		"info": response,
		"id":   *(response).At(MkString("id")),
	})
}

func (this *Coinbasepro) FetchTransactions(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	this.LoadAccounts()
	currency := OpCopy(MkUndefined())
	_ = currency
	id := this.SafeString(params, MkString("id"))
	_ = id
	if IsTrue(OpEq2(id, MkUndefined())) {
		if IsTrue(OpNotEq2(code, MkUndefined())) {
			currency = this.Currency(code)
			accountsByCurrencyCode := this.IndexBy(*this.At(MkString("accounts")), MkString("currency"))
			_ = accountsByCurrencyCode
			account := this.SafeValue(accountsByCurrencyCode, code)
			_ = account
			if IsTrue(OpEq2(account, MkUndefined())) {
				panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" fetchTransactions() could not find account id for "), code))))
			}
			id = *(account).At(MkString("id"))
		}
	}
	request := MkMap(&VarMap{})
	_ = request
	if IsTrue(OpNotEq2(id, MkUndefined())) {
		*(request).At(MkString("id")) = OpCopy(id)
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := OpCopy(MkUndefined())
	_ = response
	if IsTrue(OpEq2(id, MkUndefined())) {
		response = this.Call(MkString("privateGetTransfers"), this.Extend(request, params))
		for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
			account_id := this.SafeString(*(response).At(i), MkString("account_id"))
			_ = account_id
			account := this.SafeValue(*this.At(MkString("accountsById")), account_id)
			_ = account
			code := this.SafeString(account, MkString("currency"))
			_ = code
			*(*(response).At(i)).At(MkString("currency")) = OpCopy(code)
		}
	} else {
		response = this.Call(MkString("privateGetAccountsIdTransfers"), this.Extend(request, params))
		for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
			*(*(response).At(i)).At(MkString("currency")) = OpCopy(code)
		}
	}
	return this.ParseTransactions(response, currency, since, limit)
}

func (this *Coinbasepro) FetchDeposits(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	return this.FetchTransactions(code, since, limit, this.Extend(MkMap(&VarMap{"type": MkString("deposit")}), params))
}

func (this *Coinbasepro) FetchWithdrawals(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	return this.FetchTransactions(code, since, limit, this.Extend(MkMap(&VarMap{"type": MkString("withdraw")}), params))
}

func (this *Coinbasepro) ParseTransactionStatus(goArgs ...*Variant) *Variant {
	transaction := GoGetArg(goArgs, 0, MkUndefined())
	_ = transaction
	canceled := this.SafeValue(transaction, MkString("canceled_at"))
	_ = canceled
	if IsTrue(canceled) {
		return MkString("canceled")
	}
	processed := this.SafeValue(transaction, MkString("processed_at"))
	_ = processed
	completed := this.SafeValue(transaction, MkString("completed_at"))
	_ = completed
	if IsTrue(completed) {
		return MkString("ok")
	} else {
		if IsTrue(OpAnd(processed, OpNot(completed))) {
			return MkString("failed")
		} else {
			return MkString("pending")
		}
	}
	return MkUndefined()
}

func (this *Coinbasepro) ParseTransaction(goArgs ...*Variant) *Variant {
	transaction := GoGetArg(goArgs, 0, MkUndefined())
	_ = transaction
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	details := this.SafeValue(transaction, MkString("details"), MkMap(&VarMap{}))
	_ = details
	id := this.SafeString(transaction, MkString("id"))
	_ = id
	txid := this.SafeString(details, MkString("crypto_transaction_hash"))
	_ = txid
	timestamp := this.Parse8601(this.SafeString(transaction, MkString("created_at")))
	_ = timestamp
	updated := this.Parse8601(this.SafeString(transaction, MkString("processed_at")))
	_ = updated
	currencyId := this.SafeString(transaction, MkString("currency"))
	_ = currencyId
	code := this.SafeCurrencyCode(currencyId, currency)
	_ = code
	status := this.ParseTransactionStatus(transaction)
	_ = status
	amount := this.SafeNumber(transaction, MkString("amount"))
	_ = amount
	type_ := this.SafeString(transaction, MkString("type"))
	_ = type_
	address := this.SafeString(details, MkString("crypto_address"))
	_ = address
	tag := this.SafeString(details, MkString("destination_tag"))
	_ = tag
	address = this.SafeString(transaction, MkString("crypto_address"), address)
	fee := OpCopy(MkUndefined())
	_ = fee
	if IsTrue(OpEq2(type_, MkString("withdraw"))) {
		type_ = MkString("withdrawal")
		address = this.SafeString(details, MkString("sent_to_address"), address)
		feeCost := this.SafeNumber(details, MkString("fee"))
		_ = feeCost
		if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
			if IsTrue(OpNotEq2(amount, MkUndefined())) {
				OpSubAssign(&amount, feeCost)
			}
			fee = MkMap(&VarMap{
				"cost":     feeCost,
				"currency": code,
			})
		}
	}
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

func (this *Coinbasepro) CreateDepositAddress(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	accounts := this.SafeValue(*this.At(MkString("options")), MkString("coinbaseAccounts"))
	_ = accounts
	if IsTrue(OpEq2(accounts, MkUndefined())) {
		accounts = this.Call(MkString("privateGetCoinbaseAccounts"))
		*(*this.At(MkString("options"))).At(MkString("coinbaseAccounts")) = OpCopy(accounts)
		*(*this.At(MkString("options"))).At(MkString("coinbaseAccountsByCurrencyId")) = this.IndexBy(accounts, MkString("currency"))
	}
	currencyId := *(currency).At(MkString("id"))
	_ = currencyId
	account := this.SafeValue(*(*this.At(MkString("options"))).At(MkString("coinbaseAccountsByCurrencyId")), currencyId)
	_ = account
	if IsTrue(OpEq2(account, MkUndefined())) {
		panic(NewInvalidAddress(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" fetchDepositAddress() could not find currency code "), OpAdd(code, OpAdd(MkString(" with id = "), OpAdd(currencyId, MkString(" in this.options['coinbaseAccountsByCurrencyId']"))))))))
	}
	request := MkMap(&VarMap{"id": *(account).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("privatePostCoinbaseAccountsIdAddresses"), this.Extend(request, params))
	_ = response
	address := this.SafeString(response, MkString("address"))
	_ = address
	tag := this.SafeString(response, MkString("destination_tag"))
	_ = tag
	return MkMap(&VarMap{
		"currency": code,
		"address":  this.CheckAddress(address),
		"tag":      tag,
		"info":     response,
	})
}

func (this *Coinbasepro) Sign(goArgs ...*Variant) *Variant {
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
	query := this.Omit(params, this.ExtractParams(path))
	_ = query
	if IsTrue(OpEq2(method, MkString("GET"))) {
		if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
			OpAddAssign(&request, OpAdd(MkString("?"), this.Urlencode(query)))
		}
	}
	url := OpAdd(this.ImplodeHostname(*(*(*this.At(MkString("urls"))).At(MkString("api"))).At(api)), request)
	_ = url
	if IsTrue(OpEq2(api, MkString("private"))) {
		this.CheckRequiredCredentials()
		nonce := (this.Nonce()).Call(MkString("toString"))
		_ = nonce
		payload := MkString("")
		_ = payload
		if IsTrue(OpNotEq2(method, MkString("GET"))) {
			if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
				body = this.Json(query)
				payload = OpCopy(body)
			}
		}
		what := OpAdd(nonce, OpAdd(method, OpAdd(request, payload)))
		_ = what
		secret := this.Base64ToBinary(*this.At(MkString("secret")))
		_ = secret
		signature := this.Hmac(this.Encode(what), secret, MkString("sha256"), MkString("base64"))
		_ = signature
		headers = MkMap(&VarMap{
			"CB-ACCESS-KEY":        *this.At(MkString("apiKey")),
			"CB-ACCESS-SIGN":       signature,
			"CB-ACCESS-TIMESTAMP":  nonce,
			"CB-ACCESS-PASSPHRASE": *this.At(MkString("password")),
			"Content-Type":         MkString("application/json"),
		})
	}
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Coinbasepro) HandleErrors(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpOr(OpEq2(code, MkInteger(400)), OpEq2(code, MkInteger(404)))) {
		if IsTrue(OpEq2(*(body).At(MkInteger(0)), MkString("{"))) {
			message := this.SafeString(response, MkString("message"))
			_ = message
			feedback := OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), message))
			_ = feedback
			this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("exact")), message, feedback)
			this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("broad")), message, feedback)
			panic(NewExchangeError(feedback))
		}
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))))
	}
	return MkUndefined()
}

func (this *Coinbasepro) Request(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpNotEq2(OpType(response), MkString("string"))) {
		if IsTrue(OpHasMember(MkString("message"), response)) {
			panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), this.Json(response)))))
		}
	}
	return response
}
