package ccxt

import ()

type Buda struct {
	*ExchangeBase
}

var _ Exchange = (*Buda)(nil)

func init() {
	exchange := &Buda{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Buda) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("buda"),
		"name": MkString("Buda"),
		"countries": MkArray(&VarArray{
			MkString("AR"),
			MkString("CL"),
			MkString("CO"),
			MkString("PE"),
		}),
		"rateLimit": MkInteger(1000),
		"version":   MkString("v2"),
		"has": MkMap(&VarMap{
			"cancelOrder":          MkBool(true),
			"CORS":                 MkBool(false),
			"createDepositAddress": MkBool(true),
			"createOrder":          MkBool(true),
			"fetchBalance":         MkBool(true),
			"fetchClosedOrders":    MkBool(true),
			"fetchCurrencies":      MkBool(true),
			"fetchDepositAddress":  MkBool(true),
			"fetchDeposits":        MkBool(true),
			"fetchFundingFees":     MkBool(true),
			"fetchMarkets":         MkBool(true),
			"fetchMyTrades":        MkBool(false),
			"fetchOHLCV":           MkBool(true),
			"fetchOpenOrders":      MkBool(true),
			"fetchOrder":           MkBool(true),
			"fetchOrderBook":       MkBool(true),
			"fetchOrders":          MkBool(true),
			"fetchTrades":          MkBool(true),
			"fetchTicker":          MkBool(true),
			"fetchWithdrawals":     MkBool(true),
			"withdraw":             MkBool(true),
		}),
		"urls": MkMap(&VarMap{
			"logo": MkString("https://user-images.githubusercontent.com/1294454/47380619-8a029200-d706-11e8-91e0-8a391fe48de3.jpg"),
			"api":  MkString("https://www.buda.com/api"),
			"www":  MkString("https://www.buda.com"),
			"doc":  MkString("https://api.buda.com"),
			"fees": MkString("https://www.buda.com/comisiones"),
		}),
		"status": MkMap(&VarMap{
			"status":  MkString("error"),
			"updated": MkUndefined(),
			"eta":     MkUndefined(),
			"url":     MkUndefined(),
		}),
		"api": MkMap(&VarMap{
			"public": MkMap(&VarMap{
				"get": MkArray(&VarArray{
					MkString("pairs"),
					MkString("markets"),
					MkString("currencies"),
					MkString("markets/{market}"),
					MkString("markets/{market}/ticker"),
					MkString("markets/{market}/volume"),
					MkString("markets/{market}/order_book"),
					MkString("markets/{market}/trades"),
					MkString("currencies/{currency}/fees/deposit"),
					MkString("currencies/{currency}/fees/withdrawal"),
					MkString("tv/history"),
				}),
				"post": MkArray(&VarArray{
					MkString("markets/{market}/quotations"),
				}),
			}),
			"private": MkMap(&VarMap{
				"get": MkArray(&VarArray{
					MkString("balances"),
					MkString("balances/{currency}"),
					MkString("currencies/{currency}/balances"),
					MkString("orders"),
					MkString("orders/{id}"),
					MkString("markets/{market}/orders"),
					MkString("deposits"),
					MkString("currencies/{currency}/deposits"),
					MkString("withdrawals"),
					MkString("currencies/{currency}/withdrawals"),
					MkString("currencies/{currency}/receive_addresses"),
					MkString("currencies/{currency}/receive_addresses/{id}"),
				}),
				"post": MkArray(&VarArray{
					MkString("markets/{market}/orders"),
					MkString("currencies/{currency}/deposits"),
					MkString("currencies/{currency}/withdrawals"),
					MkString("currencies/{currency}/simulated_withdrawals"),
					MkString("currencies/{currency}/receive_addresses"),
				}),
				"put": MkArray(&VarArray{
					MkString("orders/{id}"),
				}),
			}),
		}),
		"timeframes": MkMap(&VarMap{
			"1m":  MkString("1"),
			"5m":  MkString("5"),
			"30m": MkString("30"),
			"1h":  MkString("60"),
			"2h":  MkString("120"),
			"1d":  MkString("D"),
			"1w":  MkString("W"),
		}),
		"fees": MkMap(&VarMap{"trading": MkMap(&VarMap{
			"tierBased":  MkBool(true),
			"percentage": MkBool(true),
			"taker":      MkNumber(0.008),
			"maker":      MkNumber(0.004),
			"tiers": MkMap(&VarMap{
				"taker": MkArray(&VarArray{
					MkArray(&VarArray{
						MkInteger(0),
						MkNumber(0.008),
					}),
					MkArray(&VarArray{
						MkInteger(2000),
						MkNumber(0.007),
					}),
					MkArray(&VarArray{
						MkInteger(20000),
						MkNumber(0.006),
					}),
					MkArray(&VarArray{
						MkInteger(100000),
						MkNumber(0.005),
					}),
					MkArray(&VarArray{
						MkInteger(500000),
						MkNumber(0.004),
					}),
					MkArray(&VarArray{
						MkInteger(2500000),
						MkNumber(0.003),
					}),
					MkArray(&VarArray{
						MkInteger(12500000),
						MkNumber(0.002),
					}),
				}),
				"maker": MkArray(&VarArray{
					MkArray(&VarArray{
						MkInteger(0),
						MkNumber(0.004),
					}),
					MkArray(&VarArray{
						MkInteger(2000),
						MkNumber(0.0035),
					}),
					MkArray(&VarArray{
						MkInteger(20000),
						MkNumber(0.003),
					}),
					MkArray(&VarArray{
						MkInteger(100000),
						MkNumber(0.0025),
					}),
					MkArray(&VarArray{
						MkInteger(500000),
						MkNumber(0.002),
					}),
					MkArray(&VarArray{
						MkInteger(2500000),
						MkNumber(0.0015),
					}),
					MkArray(&VarArray{
						MkInteger(12500000),
						MkNumber(0.001),
					}),
				}),
			}),
		})}),
		"exceptions": MkMap(&VarMap{
			"not_authorized":    AuthenticationError,
			"forbidden":         PermissionDenied,
			"invalid_record":    ExchangeError,
			"not_found":         ExchangeError,
			"parameter_missing": ExchangeError,
			"bad_parameter":     ExchangeError,
		}),
	}))
}

func (this *Buda) FetchCurrencyInfo(goArgs ...*Variant) *Variant {
	currency := GoGetArg(goArgs, 0, MkUndefined())
	_ = currency
	currencies := GoGetArg(goArgs, 1, MkUndefined())
	_ = currencies
	if IsTrue(OpNot(currencies)) {
		response := this.Call(MkString("publicGetCurrencies"))
		_ = response
		currencies = this.SafeValue(response, MkString("currencies"))
	}
	for i := MkInteger(0); IsTrue(OpLw(i, currencies.Length)); OpIncr(&i) {
		currencyInfo := *(currencies).At(i)
		_ = currencyInfo
		if IsTrue(OpEq2(*(currencyInfo).At(MkString("id")), currency)) {
			return currencyInfo
		}
	}
	return MkUndefined()
}

func (this *Buda) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	marketsResponse := this.Call(MkString("publicGetMarkets"), params)
	_ = marketsResponse
	markets := this.SafeValue(marketsResponse, MkString("markets"))
	_ = markets
	currenciesResponse := this.Call(MkString("publicGetCurrencies"))
	_ = currenciesResponse
	currencies := this.SafeValue(currenciesResponse, MkString("currencies"))
	_ = currencies
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, markets.Length)); OpIncr(&i) {
		market := *(markets).At(i)
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
		baseInfo := this.FetchCurrencyInfo(baseId, currencies)
		_ = baseInfo
		quoteInfo := this.FetchCurrencyInfo(quoteId, currencies)
		_ = quoteInfo
		symbol := OpAdd(base, OpAdd(MkString("/"), quote))
		_ = symbol
		pricePrecisionString := this.SafeString(quoteInfo, MkString("input_decimals"))
		_ = pricePrecisionString
		priceLimit := this.ParsePrecision(pricePrecisionString)
		_ = priceLimit
		precision := MkMap(&VarMap{
			"amount": this.SafeInteger(baseInfo, MkString("input_decimals")),
			"price":  parseInt(pricePrecisionString),
		})
		_ = precision
		minimumOrderAmount := this.SafeValue(market, MkString("minimum_order_amount"), MkArray(&VarArray{}))
		_ = minimumOrderAmount
		limits := MkMap(&VarMap{
			"amount": MkMap(&VarMap{
				"min": this.SafeNumber(minimumOrderAmount, MkInteger(0)),
				"max": MkUndefined(),
			}),
			"price": MkMap(&VarMap{
				"min": priceLimit,
				"max": MkUndefined(),
			}),
			"cost": MkMap(&VarMap{
				"min": MkUndefined(),
				"max": MkUndefined(),
			}),
		})
		_ = limits
		result.Push(MkMap(&VarMap{
			"id":        id,
			"symbol":    symbol,
			"base":      base,
			"quote":     quote,
			"baseId":    baseId,
			"quoteId":   quoteId,
			"active":    MkBool(true),
			"precision": precision,
			"limits":    limits,
			"info":      market,
		}))
	}
	return result
}

func (this *Buda) FetchCurrencies(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetCurrencies"))
	_ = response
	currencies := *(response).At(MkString("currencies"))
	_ = currencies
	result := MkMap(&VarMap{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, currencies.Length)); OpIncr(&i) {
		currency := *(currencies).At(i)
		_ = currency
		if IsTrue(OpNot(*(currency).At(MkString("managed")))) {
			continue
		}
		id := this.SafeString(currency, MkString("id"))
		_ = id
		code := this.SafeCurrencyCode(id)
		_ = code
		precision := this.SafeNumber(currency, MkString("input_decimals"))
		_ = precision
		minimum := Math.Pow(MkInteger(10), OpNeg(precision))
		_ = minimum
		*(result).At(code) = MkMap(&VarMap{
			"id":        id,
			"code":      code,
			"info":      currency,
			"name":      MkUndefined(),
			"active":    MkBool(true),
			"fee":       MkUndefined(),
			"precision": precision,
			"limits": MkMap(&VarMap{
				"amount": MkMap(&VarMap{
					"min": minimum,
					"max": MkUndefined(),
				}),
				"deposit": MkMap(&VarMap{
					"min": parseFloat(*(*(currency).At(MkString("deposit_minimum"))).At(MkInteger(0))),
					"max": MkUndefined(),
				}),
				"withdraw": MkMap(&VarMap{"min": parseFloat(*(*(currency).At(MkString("withdrawal_minimum"))).At(MkInteger(0)))}),
			}),
		})
	}
	return result
}

func (this *Buda) FetchFundingFees(goArgs ...*Variant) *Variant {
	codes := GoGetArg(goArgs, 0, MkUndefined())
	_ = codes
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	withdrawFees := MkMap(&VarMap{})
	_ = withdrawFees
	depositFees := MkMap(&VarMap{})
	_ = depositFees
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
		request := MkMap(&VarMap{"currency": *(currency).At(MkString("id"))})
		_ = request
		withdrawResponse := this.Call(MkString("publicGetCurrenciesCurrencyFeesWithdrawal"), request)
		_ = withdrawResponse
		depositResponse := this.Call(MkString("publicGetCurrenciesCurrencyFeesDeposit"), request)
		_ = depositResponse
		*(withdrawFees).At(code) = this.ParseFundingFee(*(withdrawResponse).At(MkString("fee")))
		*(depositFees).At(code) = this.ParseFundingFee(*(depositResponse).At(MkString("fee")))
		*(info).At(code) = MkMap(&VarMap{
			"withdraw": withdrawResponse,
			"deposit":  depositResponse,
		})
	}
	return MkMap(&VarMap{
		"withdraw": withdrawFees,
		"deposit":  depositFees,
		"info":     info,
	})
}

func (this *Buda) ParseFundingFee(goArgs ...*Variant) *Variant {
	fee := GoGetArg(goArgs, 0, MkUndefined())
	_ = fee
	type_ := GoGetArg(goArgs, 1, MkUndefined())
	_ = type_
	if IsTrue(OpEq2(type_, MkUndefined())) {
		type_ = *(fee).At(MkString("name"))
	}
	if IsTrue(OpEq2(type_, MkString("withdrawal"))) {
		type_ = MkString("withdraw")
	}
	return MkMap(&VarMap{
		"type":     type_,
		"currency": *(*(fee).At(MkString("base"))).At(MkInteger(1)),
		"rate":     *(fee).At(MkString("percent")),
		"cost":     parseFloat(*(*(fee).At(MkString("base"))).At(MkInteger(0))),
	})
}

func (this *Buda) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"market": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("publicGetMarketsMarketTicker"), this.Extend(request, params))
	_ = response
	ticker := this.SafeValue(response, MkString("ticker"))
	_ = ticker
	return this.ParseTicker(ticker, market)
}

func (this *Buda) ParseTicker(goArgs ...*Variant) *Variant {
	ticker := GoGetArg(goArgs, 0, MkUndefined())
	_ = ticker
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.Milliseconds()
	_ = timestamp
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(OpNotEq2(market, MkUndefined())) {
		symbol = *(market).At(MkString("symbol"))
	}
	last := parseFloat(*(*(ticker).At(MkString("last_price"))).At(MkInteger(0)))
	_ = last
	percentage := parseFloat(*(ticker).At(MkString("price_variation_24h")))
	_ = percentage
	open := parseFloat(this.PriceToPrecision(symbol, OpDiv(last, OpAdd(percentage, MkInteger(1)))))
	_ = open
	change := OpSub(last, open)
	_ = change
	average := OpDiv(this.Sum(last, open), MkInteger(2))
	_ = average
	return MkMap(&VarMap{
		"symbol":        symbol,
		"timestamp":     timestamp,
		"datetime":      this.Iso8601(timestamp),
		"high":          MkUndefined(),
		"low":           MkUndefined(),
		"bid":           parseFloat(*(*(ticker).At(MkString("max_bid"))).At(MkInteger(0))),
		"bidVolume":     MkUndefined(),
		"ask":           parseFloat(*(*(ticker).At(MkString("min_ask"))).At(MkInteger(0))),
		"askVolume":     MkUndefined(),
		"vwap":          MkUndefined(),
		"open":          open,
		"close":         last,
		"last":          last,
		"previousClose": open,
		"change":        change,
		"percentage":    OpMulti(percentage, MkInteger(100)),
		"average":       average,
		"baseVolume":    parseFloat(*(*(ticker).At(MkString("volume"))).At(MkInteger(0))),
		"quoteVolume":   MkUndefined(),
		"info":          ticker,
	})
}

func (this *Buda) FetchTrades(goArgs ...*Variant) *Variant {
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
	response := this.Call(MkString("publicGetMarketsMarketTrades"), this.Extend(request, params))
	_ = response
	return this.ParseTrades(*(*(response).At(MkString("trades"))).At(MkString("entries")), market, since, limit)
}

func (this *Buda) ParseTrade(goArgs ...*Variant) *Variant {
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
	order := OpCopy(MkUndefined())
	_ = order
	fee := OpCopy(MkUndefined())
	_ = fee
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(market) {
		symbol = *(market).At(MkString("symbol"))
	}
	if IsTrue(GoIsArray(trade)) {
		timestamp = this.SafeInteger(trade, MkInteger(0))
		priceString = this.SafeString(trade, MkInteger(1))
		amountString = this.SafeString(trade, MkInteger(2))
		side = this.SafeString(trade, MkInteger(3))
		id = this.SafeString(trade, MkInteger(4))
	}
	price := this.ParseNumber(priceString)
	_ = price
	amount := this.ParseNumber(amountString)
	_ = amount
	cost := this.ParseNumber(Precise.StringMul(priceString, amountString))
	_ = cost
	return MkMap(&VarMap{
		"id":           id,
		"order":        order,
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

func (this *Buda) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"market": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("publicGetMarketsMarketOrderBook"), this.Extend(request, params))
	_ = response
	orderbook := this.SafeValue(response, MkString("order_book"))
	_ = orderbook
	return this.ParseOrderBook(orderbook, symbol)
}

func (this *Buda) FetchOHLCV(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpEq2(since, MkUndefined())) {
		since = OpSub(this.Milliseconds(), MkInteger(86400000))
	}
	request := MkMap(&VarMap{
		"symbol":     *(market).At(MkString("id")),
		"resolution": *(*this.At(MkString("timeframes"))).At(timeframe),
		"from":       OpDiv(since, MkInteger(1000)),
		"to":         this.Seconds(),
	})
	_ = request
	response := this.Call(MkString("publicGetTvHistory"), this.Extend(request, params))
	_ = response
	return this.ParseTradingViewOHLCV(response, market, timeframe, since, limit)
}

func (this *Buda) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privateGetBalances"), params)
	_ = response
	result := MkMap(&VarMap{"info": response})
	_ = result
	balances := this.SafeValue(response, MkString("balances"))
	_ = balances
	for i := MkInteger(0); IsTrue(OpLw(i, balances.Length)); OpIncr(&i) {
		balance := *(balances).At(i)
		_ = balance
		currencyId := this.SafeString(balance, MkString("id"))
		_ = currencyId
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		account := this.Account()
		_ = account
		*(account).At(MkString("free")) = this.SafeString(*(balance).At(MkString("available_amount")), MkInteger(0))
		*(account).At(MkString("total")) = this.SafeString(*(balance).At(MkString("amount")), MkInteger(0))
		*(result).At(code) = OpCopy(account)
	}
	return this.ParseBalance(result)
}

func (this *Buda) FetchOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"id": parseInt(id)})
	_ = request
	response := this.Call(MkString("privateGetOrdersId"), this.Extend(request, params))
	_ = response
	order := this.SafeValue(response, MkString("order"))
	_ = order
	return this.ParseOrder(order)
}

func (this *Buda) FetchOrders(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{
		"market": *(market).At(MkString("id")),
		"per":    limit,
	})
	_ = request
	response := this.Call(MkString("privateGetMarketsMarketOrders"), this.Extend(request, params))
	_ = response
	orders := this.SafeValue(response, MkString("orders"))
	_ = orders
	return this.ParseOrders(orders, market, since, limit)
}

func (this *Buda) FetchOpenOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	request := MkMap(&VarMap{"state": MkString("pending")})
	_ = request
	return this.FetchOrders(symbol, since, limit, this.Extend(request, params))
}

func (this *Buda) FetchClosedOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	request := MkMap(&VarMap{"state": MkString("traded")})
	_ = request
	return this.FetchOrders(symbol, since, limit, this.Extend(request, params))
}

func (this *Buda) CreateOrder(goArgs ...*Variant) *Variant {
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
	side = OpTrinary(OpEq2(side, MkString("buy")), MkString("Bid"), MkString("Ask"))
	request := MkMap(&VarMap{
		"market":     this.MarketId(symbol),
		"price_type": type_,
		"type":       side,
		"amount":     this.AmountToPrecision(symbol, amount),
	})
	_ = request
	if IsTrue(OpEq2(type_, MkString("limit"))) {
		*(request).At(MkString("limit")) = this.PriceToPrecision(symbol, price)
	}
	response := this.Call(MkString("privatePostMarketsMarketOrders"), this.Extend(request, params))
	_ = response
	order := this.SafeValue(response, MkString("order"))
	_ = order
	return this.ParseOrder(order)
}

func (this *Buda) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{
		"id":    parseInt(id),
		"state": MkString("canceling"),
	})
	_ = request
	response := this.Call(MkString("privatePutOrdersId"), this.Extend(request, params))
	_ = response
	order := this.SafeValue(response, MkString("order"))
	_ = order
	return this.ParseOrder(order)
}

func (this *Buda) ParseOrderStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"traded":    MkString("closed"),
		"received":  MkString("open"),
		"canceling": MkString("canceled"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Buda) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	id := this.SafeString(order, MkString("id"))
	_ = id
	timestamp := this.Parse8601(this.SafeString(order, MkString("created_at")))
	_ = timestamp
	marketId := this.SafeString(order, MkString("market_id"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market, MkString("-"))
	_ = symbol
	type_ := this.SafeString(order, MkString("price_type"))
	_ = type_
	side := this.SafeStringLower(order, MkString("type"))
	_ = side
	status := this.ParseOrderStatus(this.SafeString(order, MkString("state")))
	_ = status
	originalAmount := this.SafeValue(order, MkString("original_amount"), MkArray(&VarArray{}))
	_ = originalAmount
	amount := this.SafeNumber(originalAmount, MkInteger(0))
	_ = amount
	remainingAmount := this.SafeValue(order, MkString("amount"), MkArray(&VarArray{}))
	_ = remainingAmount
	remaining := this.SafeNumber(remainingAmount, MkInteger(0))
	_ = remaining
	tradedAmount := this.SafeValue(order, MkString("traded_amount"), MkArray(&VarArray{}))
	_ = tradedAmount
	filled := this.SafeNumber(tradedAmount, MkInteger(0))
	_ = filled
	totalExchanged := this.SafeValue(order, MkString("totalExchanged"), MkArray(&VarArray{}))
	_ = totalExchanged
	cost := this.SafeNumber(totalExchanged, MkInteger(0))
	_ = cost
	limitPrice := this.SafeValue(order, MkString("limit"), MkArray(&VarArray{}))
	_ = limitPrice
	price := this.SafeNumber(limitPrice, MkInteger(0))
	_ = price
	if IsTrue(OpEq2(price, MkUndefined())) {
		if IsTrue(OpNotEq2(limitPrice, MkUndefined())) {
			price = OpCopy(limitPrice)
		}
	}
	paidFee := this.SafeValue(order, MkString("paid_fee"), MkArray(&VarArray{}))
	_ = paidFee
	feeCost := this.SafeNumber(paidFee, MkInteger(0))
	_ = feeCost
	fee := OpCopy(MkUndefined())
	_ = fee
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		feeCurrencyId := this.SafeString(paidFee, MkInteger(1))
		_ = feeCurrencyId
		feeCurrencyCode := this.SafeCurrencyCode(feeCurrencyId)
		_ = feeCurrencyCode
		fee = MkMap(&VarMap{
			"cost": feeCost,
			"code": feeCurrencyCode,
		})
	}
	return this.SafeOrder(MkMap(&VarMap{
		"info":               order,
		"id":                 id,
		"clientOrderId":      MkUndefined(),
		"datetime":           this.Iso8601(timestamp),
		"timestamp":          timestamp,
		"lastTradeTimestamp": MkUndefined(),
		"status":             status,
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
		"remaining":          remaining,
		"trades":             MkUndefined(),
		"fee":                fee,
	}))
}

func (this *Buda) IsFiat(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	fiats := MkMap(&VarMap{
		"ARS": MkBool(true),
		"CLP": MkBool(true),
		"COP": MkBool(true),
		"PEN": MkBool(true),
	})
	_ = fiats
	return this.SafeValue(fiats, code, MkBool(false))
}

func (this *Buda) FetchDepositAddress(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	if IsTrue(this.IsFiat(code)) {
		panic(NewNotSupported(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" fetchDepositAddress() for fiat "), OpAdd(code, MkString(" is not supported"))))))
	}
	request := MkMap(&VarMap{"currency": *(currency).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("privateGetCurrenciesCurrencyReceiveAddresses"), this.Extend(request, params))
	_ = response
	receiveAddresses := this.SafeValue(response, MkString("receive_addresses"))
	_ = receiveAddresses
	addressPool := MkArray(&VarArray{})
	_ = addressPool
	for i := MkInteger(1); IsTrue(OpLw(i, receiveAddresses.Length)); OpIncr(&i) {
		receiveAddress := *(receiveAddresses).At(i)
		_ = receiveAddress
		if IsTrue(*(receiveAddress).At(MkString("ready"))) {
			address := *(receiveAddress).At(MkString("address"))
			_ = address
			this.CheckAddress(address)
			addressPool.Push(address)
		}
	}
	addressPoolLength := OpCopy(addressPool.Length)
	_ = addressPoolLength
	if IsTrue(OpLw(addressPoolLength, MkInteger(1))) {
		panic(NewAddressPending(OpAdd(*this.At(MkString("id")), OpAdd(MkString(": there are no addresses ready for receiving "), OpAdd(code, MkString(", retry again later)"))))))
	}
	address := *(addressPool).At(MkInteger(0))
	_ = address
	return MkMap(&VarMap{
		"currency": code,
		"address":  address,
		"tag":      MkUndefined(),
		"info":     receiveAddresses,
	})
}

func (this *Buda) CreateDepositAddress(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	if IsTrue(this.IsFiat(code)) {
		panic(NewNotSupported(OpAdd(*this.At(MkString("id")), OpAdd(MkString(": fiat fetchDepositAddress() for "), OpAdd(code, MkString(" is not supported"))))))
	}
	request := MkMap(&VarMap{"currency": *(currency).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("privatePostCurrenciesCurrencyReceiveAddresses"), this.Extend(request, params))
	_ = response
	address := this.SafeString(*(response).At(MkString("receive_address")), MkString("address"))
	_ = address
	return MkMap(&VarMap{
		"currency": code,
		"address":  address,
		"tag":      MkUndefined(),
		"info":     response,
	})
}

func (this *Buda) ParseTransactionStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"rejected":             MkString("failed"),
		"confirmed":            MkString("ok"),
		"anulled":              MkString("canceled"),
		"retained":             MkString("canceled"),
		"pending_confirmation": MkString("pending"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Buda) ParseTransaction(goArgs ...*Variant) *Variant {
	transaction := GoGetArg(goArgs, 0, MkUndefined())
	_ = transaction
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	id := this.SafeString(transaction, MkString("id"))
	_ = id
	timestamp := this.Parse8601(this.SafeString(transaction, MkString("created_at")))
	_ = timestamp
	currencyId := this.SafeString(transaction, MkString("currency"))
	_ = currencyId
	code := this.SafeCurrencyCode(currencyId, currency)
	_ = code
	amount := parseFloat(*(*(transaction).At(MkString("amount"))).At(MkInteger(0)))
	_ = amount
	fee := parseFloat(*(*(transaction).At(MkString("fee"))).At(MkInteger(0)))
	_ = fee
	feeCurrency := *(*(transaction).At(MkString("fee"))).At(MkInteger(1))
	_ = feeCurrency
	status := this.ParseTransactionStatus(this.SafeString(transaction, MkString("state")))
	_ = status
	type_ := OpTrinary(OpHasMember(MkString("deposit_data"), transaction), MkString("deposit"), MkString("withdrawal"))
	_ = type_
	data := this.SafeValue(transaction, OpAdd(type_, MkString("_data")), MkMap(&VarMap{}))
	_ = data
	address := this.SafeValue(data, MkString("target_address"))
	_ = address
	txid := this.SafeString(data, MkString("tx_hash"))
	_ = txid
	updated := this.Parse8601(this.SafeString(data, MkString("updated_at")))
	_ = updated
	return MkMap(&VarMap{
		"info":      transaction,
		"id":        id,
		"txid":      txid,
		"timestamp": timestamp,
		"datetime":  this.Iso8601(timestamp),
		"address":   address,
		"type":      type_,
		"amount":    amount,
		"currency":  code,
		"status":    status,
		"updated":   updated,
		"fee": MkMap(&VarMap{
			"cost": fee,
			"rate": feeCurrency,
		}),
	})
}

func (this *Buda) FetchDeposits(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	if IsTrue(OpEq2(code, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(": fetchDeposits() requires a currency code argument"))))
	}
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{
		"currency": *(currency).At(MkString("id")),
		"per":      limit,
	})
	_ = request
	response := this.Call(MkString("privateGetCurrenciesCurrencyDeposits"), this.Extend(request, params))
	_ = response
	deposits := this.SafeValue(response, MkString("deposits"))
	_ = deposits
	return this.ParseTransactions(deposits, currency, since, limit)
}

func (this *Buda) FetchWithdrawals(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	if IsTrue(OpEq2(code, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(": fetchDeposits() requires a currency code argument"))))
	}
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{
		"currency": *(currency).At(MkString("id")),
		"per":      limit,
	})
	_ = request
	response := this.Call(MkString("privateGetCurrenciesCurrencyWithdrawals"), this.Extend(request, params))
	_ = response
	withdrawals := this.SafeValue(response, MkString("withdrawals"))
	_ = withdrawals
	return this.ParseTransactions(withdrawals, currency, since, limit)
}

func (this *Buda) Withdraw(goArgs ...*Variant) *Variant {
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
		"currency":        *(currency).At(MkString("id")),
		"amount":          amount,
		"withdrawal_data": MkMap(&VarMap{"target_address": address}),
	})
	_ = request
	response := this.Call(MkString("privatePostCurrenciesCurrencyWithdrawals"), this.Extend(request, params))
	_ = response
	withdrawal := this.SafeValue(response, MkString("withdrawal"))
	_ = withdrawal
	return this.ParseTransaction(withdrawal)
}

func (this *Buda) Nonce(goArgs ...*Variant) *Variant {
	return this.Microseconds()
}

func (this *Buda) Sign(goArgs ...*Variant) *Variant {
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
	request := this.ImplodeParams(path, params)
	_ = request
	query := this.Omit(params, this.ExtractParams(path))
	_ = query
	if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
		if IsTrue(OpEq2(method, MkString("GET"))) {
			OpAddAssign(&request, OpAdd(MkString("?"), this.Urlencode(query)))
		} else {
			body = this.Json(query)
		}
	}
	url := OpAdd(*(*this.At(MkString("urls"))).At(MkString("api")), OpAdd(MkString("/"), OpAdd(*this.At(MkString("version")), OpAdd(MkString("/"), request))))
	_ = url
	if IsTrue(OpEq2(api, MkString("private"))) {
		this.CheckRequiredCredentials()
		nonce := (this.Nonce()).Call(MkString("toString"))
		_ = nonce
		components := MkArray(&VarArray{
			method,
			OpAdd(MkString("/api/"), OpAdd(*this.At(MkString("version")), OpAdd(MkString("/"), request))),
		})
		_ = components
		if IsTrue(body) {
			base64Body := this.StringToBase64(body)
			_ = base64Body
			components.Push(this.Decode(base64Body))
		}
		components.Push(nonce)
		message := components.Join(MkString(" "))
		_ = message
		signature := this.Hmac(this.Encode(message), this.Encode(*this.At(MkString("secret"))), MkString("sha384"))
		_ = signature
		headers = MkMap(&VarMap{
			"X-SBTC-APIKEY":    *this.At(MkString("apiKey")),
			"X-SBTC-SIGNATURE": signature,
			"X-SBTC-NONCE":     nonce,
			"Content-Type":     MkString("application/json"),
		})
	}
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Buda) HandleErrors(goArgs ...*Variant) *Variant {
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
		errorCode := this.SafeString(response, MkString("code"))
		_ = errorCode
		message := this.SafeString(response, MkString("message"), body)
		_ = message
		feedback := OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), message))
		_ = feedback
		if IsTrue(OpNotEq2(errorCode, MkUndefined())) {
			this.ThrowExactlyMatchedException(*this.At(MkString("exceptions")), errorCode, feedback)
			panic(NewExchangeError(feedback))
		}
	}
	return MkUndefined()
}
