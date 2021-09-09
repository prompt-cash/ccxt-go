package ccxt

import ()

type Novadax struct {
	*ExchangeBase
}

var _ Exchange = (*Novadax)(nil)

func init() {
	exchange := &Novadax{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Novadax) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("novadax"),
		"name": MkString("NovaDAX"),
		"countries": MkArray(&VarArray{
			MkString("BR"),
		}),
		"rateLimit": MkInteger(50),
		"version":   MkString("v1"),
		"has": MkMap(&VarMap{
			"CORS":              MkBool(false),
			"cancelOrder":       MkBool(true),
			"createOrder":       MkBool(true),
			"fetchAccounts":     MkBool(true),
			"fetchBalance":      MkBool(true),
			"fetchClosedOrders": MkBool(true),
			"fetchDeposits":     MkBool(true),
			"fetchMarkets":      MkBool(true),
			"fetchMyTrades":     MkBool(true),
			"fetchOHLCV":        MkBool(true),
			"fetchOpenOrders":   MkBool(true),
			"fetchOrder":        MkBool(true),
			"fetchOrders":       MkBool(true),
			"fetchOrderTrades":  MkBool(true),
			"fetchOrderBook":    MkBool(true),
			"fetchTicker":       MkBool(true),
			"fetchTickers":      MkBool(true),
			"fetchTime":         MkBool(true),
			"fetchTrades":       MkBool(true),
			"fetchTransactions": MkBool(true),
			"fetchWithdrawals":  MkBool(true),
			"withdraw":          MkBool(true),
		}),
		"timeframes": MkMap(&VarMap{
			"1m":  MkString("ONE_MIN"),
			"5m":  MkString("FIVE_MIN"),
			"15m": MkString("FIFTEEN_MIN"),
			"30m": MkString("HALF_HOU"),
			"1h":  MkString("ONE_HOU"),
			"1d":  MkString("ONE_DAY"),
			"1w":  MkString("ONE_WEE"),
			"1M":  MkString("ONE_MON"),
		}),
		"urls": MkMap(&VarMap{
			"logo": MkString("https://user-images.githubusercontent.com/1294454/92337550-2b085500-f0b3-11ea-98e7-5794fb07dd3b.jpg"),
			"api": MkMap(&VarMap{
				"public":  MkString("https://api.novadax.com"),
				"private": MkString("https://api.novadax.com"),
			}),
			"www": MkString("https://www.novadax.com.br"),
			"doc": MkArray(&VarArray{
				MkString("https://doc.novadax.com/pt-BR/"),
			}),
			"fees":     MkString("https://www.novadax.com.br/fees-and-limits"),
			"referral": MkString("https://www.novadax.com.br/?s=ccxt"),
		}),
		"api": MkMap(&VarMap{
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("common/symbol"),
				MkString("common/symbols"),
				MkString("common/timestamp"),
				MkString("market/tickers"),
				MkString("market/ticker"),
				MkString("market/depth"),
				MkString("market/trades"),
				MkString("market/kline/history"),
			})}),
			"private": MkMap(&VarMap{
				"get": MkArray(&VarArray{
					MkString("orders/get"),
					MkString("orders/list"),
					MkString("orders/fill"),
					MkString("orders/fills"),
					MkString("account/getBalance"),
					MkString("account/subs"),
					MkString("account/subs/balance"),
					MkString("account/subs/transfer/record"),
					MkString("wallet/query/deposit-withdraw"),
				}),
				"post": MkArray(&VarArray{
					MkString("orders/create"),
					MkString("orders/cancel"),
					MkString("account/withdraw/coin"),
					MkString("account/subs/transfer"),
				}),
			}),
		}),
		"fees": MkMap(&VarMap{"trading": MkMap(&VarMap{
			"tierBased":  MkBool(false),
			"percentage": MkBool(true),
			"taker":      this.ParseNumber(MkString("0.005")),
			"maker":      this.ParseNumber(MkString("0.003")),
		})}),
		"requiredCredentials": MkMap(&VarMap{
			"apiKey": MkBool(true),
			"secret": MkBool(true),
		}),
		"exceptions": MkMap(&VarMap{
			"exact": MkMap(&VarMap{
				"A99999": ExchangeError,
				"A10001": BadRequest,
				"A10002": ExchangeError,
				"A10003": AuthenticationError,
				"A10004": RateLimitExceeded,
				"A10005": PermissionDenied,
				"A10006": AccountSuspended,
				"A10007": BadRequest,
				"A10011": BadSymbol,
				"A10012": BadSymbol,
				"A10013": OnMaintenance,
				"A30001": OrderNotFound,
				"A30002": InvalidOrder,
				"A30003": InvalidOrder,
				"A30004": InvalidOrder,
				"A30005": InvalidOrder,
				"A30006": InvalidOrder,
				"A30007": InsufficientFunds,
				"A30008": InvalidOrder,
				"A30009": InvalidOrder,
				"A30010": CancelPending,
				"A30011": InvalidOrder,
				"A30012": InvalidOrder,
			}),
			"broad": MkMap(&VarMap{}),
		}),
		"options": MkMap(&VarMap{"fetchOHLCV": MkMap(&VarMap{"volume": MkString("amount")})}),
	}))
}

func (this *Novadax) FetchTime(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetCommonTimestamp"), params)
	_ = response
	return this.SafeInteger(response, MkString("data"))
}

func (this *Novadax) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetCommonSymbols"), params)
	_ = response
	result := MkArray(&VarArray{})
	_ = result
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	for i := MkInteger(0); IsTrue(OpLw(i, data.Length)); OpIncr(&i) {
		market := *(data).At(i)
		_ = market
		baseId := this.SafeString(market, MkString("baseCurrency"))
		_ = baseId
		quoteId := this.SafeString(market, MkString("quoteCurrency"))
		_ = quoteId
		id := this.SafeString(market, MkString("symbol"))
		_ = id
		base := this.SafeCurrencyCode(baseId)
		_ = base
		quote := this.SafeCurrencyCode(quoteId)
		_ = quote
		symbol := OpAdd(base, OpAdd(MkString("/"), quote))
		_ = symbol
		precision := MkMap(&VarMap{
			"amount": this.SafeInteger(market, MkString("amountPrecision")),
			"price":  this.SafeInteger(market, MkString("pricePrecision")),
			"cost":   this.SafeInteger(market, MkString("valuePrecision")),
		})
		_ = precision
		limits := MkMap(&VarMap{
			"amount": MkMap(&VarMap{
				"min": this.SafeNumber(market, MkString("minOrderAmount")),
				"max": MkUndefined(),
			}),
			"price": MkMap(&VarMap{
				"min": MkUndefined(),
				"max": MkUndefined(),
			}),
			"cost": MkMap(&VarMap{
				"min": this.SafeNumber(market, MkString("minOrderValue")),
				"max": MkUndefined(),
			}),
		})
		_ = limits
		status := this.SafeString(market, MkString("status"))
		_ = status
		active := OpEq2(status, MkString("ONLINE"))
		_ = active
		result.Push(MkMap(&VarMap{
			"id":        id,
			"symbol":    symbol,
			"base":      base,
			"quote":     quote,
			"baseId":    baseId,
			"quoteId":   quoteId,
			"precision": precision,
			"limits":    limits,
			"info":      market,
			"active":    active,
		}))
	}
	return result
}

func (this *Novadax) ParseTicker(goArgs ...*Variant) *Variant {
	ticker := GoGetArg(goArgs, 0, MkUndefined())
	_ = ticker
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.SafeInteger(ticker, MkString("timestamp"))
	_ = timestamp
	marketId := this.SafeString(ticker, MkString("symbol"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market, MkString("_"))
	_ = symbol
	open := this.SafeNumber(ticker, MkString("open24h"))
	_ = open
	last := this.SafeNumber(ticker, MkString("lastPrice"))
	_ = last
	percentage := OpCopy(MkUndefined())
	_ = percentage
	change := OpCopy(MkUndefined())
	_ = change
	average := OpCopy(MkUndefined())
	_ = average
	if IsTrue(OpAnd(OpNotEq2(last, MkUndefined()), OpNotEq2(open, MkUndefined()))) {
		change = OpSub(last, open)
		percentage = OpDiv(change, OpMulti(open, MkInteger(100)))
		average = OpDiv(this.Sum(last, open), MkInteger(2))
	}
	baseVolume := this.SafeNumber(ticker, MkString("baseVolume24h"))
	_ = baseVolume
	quoteVolume := this.SafeNumber(ticker, MkString("quoteVolume24h"))
	_ = quoteVolume
	vwap := this.Vwap(baseVolume, quoteVolume)
	_ = vwap
	return MkMap(&VarMap{
		"symbol":        symbol,
		"timestamp":     timestamp,
		"datetime":      this.Iso8601(timestamp),
		"high":          this.SafeNumber(ticker, MkString("high24h")),
		"low":           this.SafeNumber(ticker, MkString("low24h")),
		"bid":           this.SafeNumber(ticker, MkString("bid")),
		"bidVolume":     MkUndefined(),
		"ask":           this.SafeNumber(ticker, MkString("ask")),
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

func (this *Novadax) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("publicGetMarketTicker"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	return this.ParseTicker(data, market)
}

func (this *Novadax) FetchTickers(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("publicGetMarketTickers"), params)
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	result := MkMap(&VarMap{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, data.Length)); OpIncr(&i) {
		ticker := this.ParseTicker(*(data).At(i))
		_ = ticker
		symbol := *(ticker).At(MkString("symbol"))
		_ = symbol
		*(result).At(symbol) = OpCopy(ticker)
	}
	return this.FilterByArray(result, MkString("symbol"), symbols)
}

func (this *Novadax) FetchOrderBook(goArgs ...*Variant) *Variant {
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
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("publicGetMarketDepth"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	timestamp := this.SafeInteger(data, MkString("timestamp"))
	_ = timestamp
	return this.ParseOrderBook(data, symbol, timestamp, MkString("bids"), MkString("asks"))
}

func (this *Novadax) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	id := this.SafeString(trade, MkString("id"))
	_ = id
	orderId := this.SafeString(trade, MkString("orderId"))
	_ = orderId
	timestamp := this.SafeInteger(trade, MkString("timestamp"))
	_ = timestamp
	side := this.SafeStringLower(trade, MkString("side"))
	_ = side
	priceString := this.SafeString(trade, MkString("price"))
	_ = priceString
	amountString := this.SafeString(trade, MkString("amount"))
	_ = amountString
	price := this.ParseNumber(priceString)
	_ = price
	amount := this.ParseNumber(amountString)
	_ = amount
	cost := this.SafeNumber(trade, MkString("volume"))
	_ = cost
	if IsTrue(OpEq2(cost, MkUndefined())) {
		cost = this.ParseNumber(Precise.StringMul(priceString, amountString))
	}
	marketId := this.SafeString(trade, MkString("symbol"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market, MkString("_"))
	_ = symbol
	takerOrMaker := this.SafeStringLower(trade, MkString("role"))
	_ = takerOrMaker
	feeString := this.SafeString(trade, MkString("fee"))
	_ = feeString
	fee := OpCopy(MkUndefined())
	_ = fee
	if IsTrue(OpNotEq2(feeString, MkUndefined())) {
		parts := feeString.Split(MkString(" "))
		_ = parts
		feeCurrencyId := this.SafeString(parts, MkInteger(1))
		_ = feeCurrencyId
		feeCurrencyCode := this.SafeCurrencyCode(feeCurrencyId)
		_ = feeCurrencyCode
		fee = MkMap(&VarMap{
			"cost":     this.SafeNumber(parts, MkInteger(0)),
			"currency": feeCurrencyCode,
		})
	}
	return MkMap(&VarMap{
		"id":           id,
		"order":        orderId,
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       symbol,
		"type":         MkUndefined(),
		"side":         side,
		"price":        price,
		"amount":       amount,
		"cost":         cost,
		"takerOrMaker": takerOrMaker,
		"fee":          fee,
		"info":         trade,
	})
}

func (this *Novadax) FetchTrades(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("publicGetMarketTrades"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseTrades(data, market, since, limit)
}

func (this *Novadax) FetchOHLCV(goArgs ...*Variant) *Variant {
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
		"symbol": *(market).At(MkString("id")),
		"unit":   *(*this.At(MkString("timeframes"))).At(timeframe),
	})
	_ = request
	duration := this.ParseTimeframe(timeframe)
	_ = duration
	now := this.Seconds()
	_ = now
	if IsTrue(OpEq2(limit, MkUndefined())) {
		limit = MkInteger(3000)
	}
	if IsTrue(OpEq2(since, MkUndefined())) {
		*(request).At(MkString("from")) = OpSub(now, OpMulti(limit, duration))
		*(request).At(MkString("to")) = OpCopy(now)
	} else {
		startFrom := parseInt(OpDiv(since, MkInteger(1000)))
		_ = startFrom
		*(request).At(MkString("from")) = OpCopy(startFrom)
		*(request).At(MkString("to")) = this.Sum(startFrom, OpMulti(limit, duration))
	}
	response := this.Call(MkString("publicGetMarketKlineHistory"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseOHLCVs(data, market, timeframe, since, limit)
}

func (this *Novadax) ParseOHLCV(goArgs ...*Variant) *Variant {
	ohlcv := GoGetArg(goArgs, 0, MkUndefined())
	_ = ohlcv
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	options := this.SafeValue(*this.At(MkString("options")), MkString("fetchOHLCV"), MkMap(&VarMap{}))
	_ = options
	volumeField := this.SafeString(options, MkString("volume"), MkString("amount"))
	_ = volumeField
	return MkArray(&VarArray{
		this.SafeTimestamp(ohlcv, MkString("score")),
		this.SafeNumber(ohlcv, MkString("openPrice")),
		this.SafeNumber(ohlcv, MkString("highPrice")),
		this.SafeNumber(ohlcv, MkString("lowPrice")),
		this.SafeNumber(ohlcv, MkString("closePrice")),
		this.SafeNumber(ohlcv, volumeField),
	})
}

func (this *Novadax) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privateGetAccountGetBalance"), params)
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	result := MkMap(&VarMap{
		"info":      response,
		"timestamp": MkUndefined(),
		"datetime":  MkUndefined(),
	})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, data.Length)); OpIncr(&i) {
		balance := *(data).At(i)
		_ = balance
		currencyId := this.SafeString(balance, MkString("currency"))
		_ = currencyId
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		account := this.Account()
		_ = account
		*(account).At(MkString("total")) = this.SafeString(balance, MkString("balance"))
		*(account).At(MkString("free")) = this.SafeString(balance, MkString("available"))
		*(account).At(MkString("used")) = this.SafeString(balance, MkString("hold"))
		*(result).At(code) = OpCopy(account)
	}
	return this.ParseBalance(result)
}

func (this *Novadax) CreateOrder(goArgs ...*Variant) *Variant {
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
	uppercaseType := type_.ToUpperCase()
	_ = uppercaseType
	uppercaseSide := side.ToUpperCase()
	_ = uppercaseSide
	request := MkMap(&VarMap{
		"symbol": *(market).At(MkString("id")),
		"side":   uppercaseSide,
	})
	_ = request
	stopPrice := this.SafeNumber(params, MkString("stopPrice"))
	_ = stopPrice
	if IsTrue(OpEq2(stopPrice, MkUndefined())) {
		if IsTrue(OpOr(OpEq2(uppercaseType, MkString("STOP_LIMIT")), OpEq2(uppercaseType, MkString("STOP_MARKET")))) {
			panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" createOrder() requires a stopPrice parameter for "), OpAdd(uppercaseType, MkString(" orders"))))))
		}
	} else {
		if IsTrue(OpEq2(uppercaseType, MkString("LIMIT"))) {
			uppercaseType = MkString("STOP_LIMIT")
		} else {
			if IsTrue(OpEq2(uppercaseType, MkString("MARKET"))) {
				uppercaseType = MkString("STOP_MARKET")
			}
		}
		defaultOperator := OpTrinary(OpEq2(uppercaseSide, MkString("BUY")), MkString("LTE"), MkString("GTE"))
		_ = defaultOperator
		*(request).At(MkString("operator")) = this.SafeString(params, MkString("operator"), defaultOperator)
		*(request).At(MkString("stopPrice")) = this.PriceToPrecision(symbol, stopPrice)
		params = this.Omit(params, MkString("stopPrice"))
	}
	if IsTrue(OpOr(OpEq2(uppercaseType, MkString("LIMIT")), OpEq2(uppercaseType, MkString("STOP_LIMIT")))) {
		*(request).At(MkString("price")) = this.PriceToPrecision(symbol, price)
		*(request).At(MkString("amount")) = this.AmountToPrecision(symbol, amount)
	} else {
		if IsTrue(OpOr(OpEq2(uppercaseType, MkString("MARKET")), OpEq2(uppercaseType, MkString("STOP_MARKET")))) {
			if IsTrue(OpEq2(uppercaseSide, MkString("SELL"))) {
				*(request).At(MkString("amount")) = this.AmountToPrecision(symbol, amount)
			} else {
				if IsTrue(OpEq2(uppercaseSide, MkString("BUY"))) {
					value := this.SafeNumber(params, MkString("value"))
					_ = value
					createMarketBuyOrderRequiresPrice := this.SafeValue(*this.At(MkString("options")), MkString("createMarketBuyOrderRequiresPrice"), MkBool(true))
					_ = createMarketBuyOrderRequiresPrice
					if IsTrue(createMarketBuyOrderRequiresPrice) {
						if IsTrue(OpNotEq2(price, MkUndefined())) {
							if IsTrue(OpEq2(value, MkUndefined())) {
								value = OpMulti(amount, price)
							}
						} else {
							if IsTrue(OpEq2(value, MkUndefined())) {
								panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")), MkString(" createOrder() requires the price argument with market buy orders to calculate total order cost (amount to spend), where cost = amount * price. Supply a price argument to createOrder() call if you want the cost to be calculated for you from price and amount, or, alternatively, add .options['createMarketBuyOrderRequiresPrice'] = false and supply the total cost value in the 'amount' argument or in the 'value' extra parameter (the exchange-specific behaviour)"))))
							}
						}
					} else {
						value = OpTrinary(OpEq2(value, MkUndefined()), amount, value)
					}
					precision := *(*(market).At(MkString("precision"))).At(MkString("price"))
					_ = precision
					*(request).At(MkString("value")) = this.DecimalToPrecision(value, TRUNCATE, precision, *this.At(MkString("precisionMode")))
				}
			}
		}
	}
	*(request).At(MkString("type")) = OpCopy(uppercaseType)
	response := this.Call(MkString("privatePostOrdersCreate"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	return this.ParseOrder(data, market)
}

func (this *Novadax) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"id": id})
	_ = request
	response := this.Call(MkString("privatePostOrdersCancel"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	return this.ParseOrder(data)
}

func (this *Novadax) FetchOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"id": id})
	_ = request
	response := this.Call(MkString("privateGetOrdersGet"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	return this.ParseOrder(data)
}

func (this *Novadax) FetchOrders(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("fromTimestamp")) = OpCopy(since)
	}
	response := this.Call(MkString("privateGetOrdersList"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseOrders(data, market, since, limit)
}

func (this *Novadax) FetchOpenOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	request := MkMap(&VarMap{"status": MkString("SUBMITTED,PROCESSING,PARTIAL_FILLED,CANCELING")})
	_ = request
	return this.FetchOrders(symbol, since, limit, this.Extend(request, params))
}

func (this *Novadax) FetchClosedOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	request := MkMap(&VarMap{"status": MkString("FILLED,CANCELED,REJECTED")})
	_ = request
	return this.FetchOrders(symbol, since, limit, this.Extend(request, params))
}

func (this *Novadax) FetchOrderTrades(goArgs ...*Variant) *Variant {
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
	response := this.Call(MkString("privateGetOrdersFill"), this.Extend(request, params))
	_ = response
	market := OpCopy(MkUndefined())
	_ = market
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
	}
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseTrades(data, market, since, limit)
}

func (this *Novadax) ParseOrderStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"SUBMITTED":      MkString("open"),
		"PROCESSING":     MkString("open"),
		"PARTIAL_FILLED": MkString("open"),
		"CANCELING":      MkString("open"),
		"FILLED":         MkString("closed"),
		"CANCELED":       MkString("canceled"),
		"REJECTED":       MkString("rejected"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Novadax) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	id := this.SafeString(order, MkString("id"))
	_ = id
	amount := this.SafeNumber(order, MkString("amount"))
	_ = amount
	price := this.SafeNumber(order, MkString("price"))
	_ = price
	cost := this.SafeNumber2(order, MkString("filledValue"), MkString("value"))
	_ = cost
	type_ := this.SafeStringLower(order, MkString("type"))
	_ = type_
	side := this.SafeStringLower(order, MkString("side"))
	_ = side
	status := this.ParseOrderStatus(this.SafeString(order, MkString("status")))
	_ = status
	timestamp := this.SafeInteger(order, MkString("timestamp"))
	_ = timestamp
	average := this.SafeNumber(order, MkString("averagePrice"))
	_ = average
	filled := this.SafeNumber(order, MkString("filledAmount"))
	_ = filled
	fee := OpCopy(MkUndefined())
	_ = fee
	feeCost := this.SafeNumber(order, MkString("filledFee"))
	_ = feeCost
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		fee = MkMap(&VarMap{
			"cost":     feeCost,
			"currency": MkUndefined(),
		})
	}
	marketId := this.SafeString(order, MkString("symbol"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market, MkString("_"))
	_ = symbol
	stopPrice := this.SafeNumber(order, MkString("stopPrice"))
	_ = stopPrice
	return this.SafeOrder(MkMap(&VarMap{
		"id":                 id,
		"clientOrderId":      MkUndefined(),
		"info":               order,
		"timestamp":          timestamp,
		"datetime":           this.Iso8601(timestamp),
		"lastTradeTimestamp": MkUndefined(),
		"symbol":             symbol,
		"type":               type_,
		"timeInForce":        MkUndefined(),
		"postOnly":           MkUndefined(),
		"side":               side,
		"price":              price,
		"stopPrice":          stopPrice,
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

func (this *Novadax) Withdraw(goArgs ...*Variant) *Variant {
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
		"code":   *(currency).At(MkString("id")),
		"amount": this.CurrencyToPrecision(code, amount),
		"wallet": address,
	})
	_ = request
	if IsTrue(OpNotEq2(tag, MkUndefined())) {
		*(request).At(MkString("tag")) = OpCopy(tag)
	}
	response := this.Call(MkString("privatePostAccountWithdrawCoin"), this.Extend(request, params))
	_ = response
	return this.ParseTransaction(response, currency)
}

func (this *Novadax) FetchAccounts(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("privateGetAccountSubs"), params)
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, data.Length)); OpIncr(&i) {
		account := *(data).At(i)
		_ = account
		accountId := this.SafeString(account, MkString("subId"))
		_ = accountId
		type_ := this.SafeString(account, MkString("subAccount"))
		_ = type_
		result.Push(MkMap(&VarMap{
			"id":       accountId,
			"type":     type_,
			"currency": MkUndefined(),
			"info":     account,
		}))
	}
	return result
}

func (this *Novadax) FetchDeposits(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	request := MkMap(&VarMap{"type": MkString("coin_in")})
	_ = request
	return this.FetchTransactions(code, since, limit, this.Extend(request, params))
}

func (this *Novadax) FetchWithdrawals(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	request := MkMap(&VarMap{"type": MkString("coin_out")})
	_ = request
	return this.FetchTransactions(code, since, limit, this.Extend(request, params))
}

func (this *Novadax) FetchTransactions(goArgs ...*Variant) *Variant {
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
		*(request).At(MkString("currency")) = *(currency).At(MkString("id"))
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("size")) = OpCopy(limit)
	}
	response := this.Call(MkString("privateGetWalletQueryDepositWithdraw"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseTransactions(data, currency, since, limit)
}

func (this *Novadax) ParseTransactionStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	parts := status.Split(MkString(" "))
	_ = parts
	status = this.SafeString(parts, MkInteger(1), status)
	statuses := MkMap(&VarMap{
		"Pending":    MkString("pending"),
		"confirming": MkString("pending"),
		"SUCCESS":    MkString("ok"),
		"FAIL":       MkString("failed"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Novadax) ParseTransaction(goArgs ...*Variant) *Variant {
	transaction := GoGetArg(goArgs, 0, MkUndefined())
	_ = transaction
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	id := this.SafeString2(transaction, MkString("id"), MkString("data"))
	_ = id
	type_ := this.SafeString(transaction, MkString("type"))
	_ = type_
	if IsTrue(OpEq2(type_, MkString("COIN_IN"))) {
		type_ = MkString("deposit")
	} else {
		if IsTrue(OpEq2(type_, MkString("COIN_OUT"))) {
			type_ = MkString("withdraw")
		}
	}
	amount := this.SafeNumber(transaction, MkString("amount"))
	_ = amount
	address := this.SafeString(transaction, MkString("address"))
	_ = address
	tag := this.SafeString(transaction, MkString("addressTag"))
	_ = tag
	txid := this.SafeString(transaction, MkString("txHash"))
	_ = txid
	timestamp := this.SafeInteger(transaction, MkString("createdAt"))
	_ = timestamp
	updated := this.SafeInteger(transaction, MkString("updatedAt"))
	_ = updated
	currencyId := this.SafeString(transaction, MkString("currency"))
	_ = currencyId
	code := this.SafeCurrencyCode(currencyId, currency)
	_ = code
	status := this.ParseTransactionStatus(this.SafeString(transaction, MkString("state")))
	_ = status
	return MkMap(&VarMap{
		"info":        transaction,
		"id":          id,
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
		"txid":        txid,
		"timestamp":   timestamp,
		"datetime":    this.Iso8601(timestamp),
		"fee":         MkUndefined(),
	})
}

func (this *Novadax) FetchMyTrades(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("fromTimestamp")) = OpCopy(since)
	}
	response := this.Call(MkString("privateGetOrdersFills"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseTrades(data, market, since, limit)
}

func (this *Novadax) Sign(goArgs ...*Variant) *Variant {
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
	request := OpAdd(MkString("/"), OpAdd(*this.At(MkString("version")), OpAdd(MkString("/"), this.ImplodeParams(path, params))))
	_ = request
	url := OpAdd(*(*(*this.At(MkString("urls"))).At(MkString("api"))).At(api), request)
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
			timestamp := (this.Milliseconds()).Call(MkString("toString"))
			_ = timestamp
			headers = MkMap(&VarMap{
				"X-Nova-Access-Key": *this.At(MkString("apiKey")),
				"X-Nova-Timestamp":  timestamp,
			})
			queryString := OpCopy(MkUndefined())
			_ = queryString
			if IsTrue(OpEq2(method, MkString("POST"))) {
				body = this.Json(query)
				queryString = this.Hash(body, MkString("md5"))
				*(headers).At(MkString("Content-Type")) = MkString("application/json")
			} else {
				if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
					OpAddAssign(&url, OpAdd(MkString("?"), this.Urlencode(query)))
				}
				queryString = this.Urlencode(this.Keysort(query))
			}
			auth := OpAdd(method, OpAdd(MkString("\n"), OpAdd(request, OpAdd(MkString("\n"), OpAdd(queryString, OpAdd(MkString("\n"), timestamp))))))
			_ = auth
			*(headers).At(MkString("X-Nova-Signature")) = this.Hmac(this.Encode(auth), this.Encode(*this.At(MkString("secret"))))
		}
	}
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Novadax) HandleErrors(goArgs ...*Variant) *Variant {
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
	errorCode := this.SafeString(response, MkString("code"))
	_ = errorCode
	if IsTrue(OpNotEq2(errorCode, MkString("A10000"))) {
		message := this.SafeString(response, MkString("message"))
		_ = message
		feedback := OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))
		_ = feedback
		this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("exact")), errorCode, feedback)
		this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("broad")), message, feedback)
		panic(NewExchangeError(feedback))
	}
	return MkUndefined()
}
