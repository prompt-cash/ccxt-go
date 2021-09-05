package ccxt

import ()

type Braziliex struct {
	*ExchangeBase
}

var _ Exchange = (*Braziliex)(nil)

func init() {
	exchange := &Braziliex{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Braziliex) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("braziliex"),
		"name": MkString("Braziliex"),
		"countries": MkArray(&VarArray{
			MkString("BR"),
		}),
		"rateLimit": MkInteger(1000),
		"has": MkMap(&VarMap{
			"cancelOrder":         MkBool(true),
			"createOrder":         MkBool(true),
			"fetchBalance":        MkBool(true),
			"fetchCurrencies":     MkBool(true),
			"fetchDepositAddress": MkBool(true),
			"fetchMarkets":        MkBool(true),
			"fetchMyTrades":       MkBool(true),
			"fetchOpenOrders":     MkBool(true),
			"fetchOrder":          MkBool(true),
			"fetchOrderBook":      MkBool(true),
			"fetchTicker":         MkBool(true),
			"fetchTickers":        MkBool(true),
			"fetchTrades":         MkBool(true),
		}),
		"urls": MkMap(&VarMap{
			"logo":     MkString("https://user-images.githubusercontent.com/1294454/34703593-c4498674-f504-11e7-8d14-ff8e44fb78c1.jpg"),
			"api":      MkString("https://braziliex.com/api/v1"),
			"www":      MkString("https://braziliex.com/"),
			"doc":      MkString("https://braziliex.com/exchange/api.php"),
			"fees":     MkString("https://braziliex.com/exchange/fees.php"),
			"referral": MkString("https://braziliex.com/?ref=5FE61AB6F6D67DA885BC98BA27223465"),
		}),
		"api": MkMap(&VarMap{
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("currencies"),
				MkString("ticker"),
				MkString("ticker/{market}"),
				MkString("orderbook/{market}"),
				MkString("tradehistory/{market}"),
			})}),
			"private": MkMap(&VarMap{"post": MkArray(&VarArray{
				MkString("balance"),
				MkString("complete_balance"),
				MkString("open_orders"),
				MkString("trade_history"),
				MkString("deposit_address"),
				MkString("sell"),
				MkString("buy"),
				MkString("cancel_order"),
				MkString("order_status"),
			})}),
		}),
		"commonCurrencies": MkMap(&VarMap{
			"EPC": MkString("Epacoin"),
			"ABC": MkString("Anti Bureaucracy Coin"),
		}),
		"fees": MkMap(&VarMap{"trading": MkMap(&VarMap{
			"maker": this.ParseNumber(MkString("0.005")),
			"taker": this.ParseNumber(MkString("0.005")),
		})}),
		"precision": MkMap(&VarMap{
			"amount": MkInteger(8),
			"price":  MkInteger(8),
		}),
		"options": MkMap(&VarMap{"fetchCurrencies": MkMap(&VarMap{"expires": MkInteger(1000)})}),
	}))
}

func (this *Braziliex) FetchCurrenciesFromCache(goArgs ...*Variant) *Variant {
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
		response := this.Call(MkString("publicGetCurrencies"), params)
		_ = response
		*(*this.At(MkString("options"))).At(MkString("fetchCurrencies")) = this.Extend(options, MkMap(&VarMap{
			"response":  response,
			"timestamp": now,
		}))
	}
	return this.SafeValue(*(*this.At(MkString("options"))).At(MkString("fetchCurrencies")), MkString("response"))
}

func (this *Braziliex) FetchCurrencies(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.FetchCurrenciesFromCache(params)
	_ = response
	*(*this.At(MkString("options"))).At(MkString("currencies")) = MkMap(&VarMap{
		"timestamp": this.Milliseconds(),
		"response":  response,
	})
	ids := GoGetKeys(response)
	_ = ids
	result := MkMap(&VarMap{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, ids.Length)); OpIncr(&i) {
		id := *(ids).At(i)
		_ = id
		currency := *(response).At(id)
		_ = currency
		precision := this.SafeInteger(currency, MkString("decimal"))
		_ = precision
		code := this.SafeCurrencyCode(id)
		_ = code
		active := OpEq2(this.SafeInteger(currency, MkString("active")), MkInteger(1))
		_ = active
		maintenance := this.SafeInteger(currency, MkString("under_maintenance"))
		_ = maintenance
		if IsTrue(OpNotEq2(maintenance, MkInteger(0))) {
			active = OpCopy(MkBool(false))
		}
		canWithdraw := OpEq2(this.SafeInteger(currency, MkString("is_withdrawal_active")), MkInteger(1))
		_ = canWithdraw
		canDeposit := OpEq2(this.SafeInteger(currency, MkString("is_deposit_active")), MkInteger(1))
		_ = canDeposit
		if IsTrue(OpOr(OpNot(canWithdraw), OpNot(canDeposit))) {
			active = OpCopy(MkBool(false))
		}
		*(result).At(code) = MkMap(&VarMap{
			"id":        id,
			"code":      code,
			"name":      *(currency).At(MkString("name")),
			"active":    active,
			"precision": precision,
			"funding": MkMap(&VarMap{
				"withdraw": MkMap(&VarMap{
					"active": canWithdraw,
					"fee":    this.SafeNumber(currency, MkString("txWithdrawalFee")),
				}),
				"deposit": MkMap(&VarMap{
					"active": canDeposit,
					"fee":    this.SafeNumber(currency, MkString("txDepositFee")),
				}),
			}),
			"limits": MkMap(&VarMap{
				"amount": MkMap(&VarMap{
					"min": Math.Pow(MkInteger(10), OpNeg(precision)),
					"max": Math.Pow(MkInteger(10), precision),
				}),
				"withdraw": MkMap(&VarMap{
					"min": this.SafeNumber(currency, MkString("MinWithdrawal")),
					"max": Math.Pow(MkInteger(10), precision),
				}),
				"deposit": MkMap(&VarMap{
					"min": this.SafeNumber(currency, MkString("minDeposit")),
					"max": MkUndefined(),
				}),
			}),
			"info": currency,
		})
	}
	return result
}

func (this *Braziliex) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	currencies := this.FetchCurrenciesFromCache(params)
	_ = currencies
	response := this.Call(MkString("publicGetTicker"))
	_ = response
	ids := GoGetKeys(response)
	_ = ids
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, ids.Length)); OpIncr(&i) {
		id := *(ids).At(i)
		_ = id
		market := *(response).At(id)
		_ = market
		baseId := MkUndefined()
		_ = baseId
		quoteId := MkUndefined()
		_ = quoteId
		ArrayUnpack(id.Split(MkString("_")), &baseId, &quoteId)
		uppercaseBaseId := baseId.ToUpperCase()
		_ = uppercaseBaseId
		uppercaseQuoteId := quoteId.ToUpperCase()
		_ = uppercaseQuoteId
		base := this.SafeCurrencyCode(uppercaseBaseId)
		_ = base
		quote := this.SafeCurrencyCode(uppercaseQuoteId)
		_ = quote
		symbol := OpAdd(base, OpAdd(MkString("/"), quote))
		_ = symbol
		baseCurrency := this.SafeValue(currencies, baseId, MkMap(&VarMap{}))
		_ = baseCurrency
		quoteCurrency := this.SafeValue(currencies, quoteId, MkMap(&VarMap{}))
		_ = quoteCurrency
		quoteIsFiat := this.SafeInteger(quoteCurrency, MkString("is_fiat"), MkInteger(0))
		_ = quoteIsFiat
		minCost := OpCopy(MkUndefined())
		_ = minCost
		if IsTrue(quoteIsFiat) {
			minCost = this.SafeNumber(baseCurrency, MkString("minAmountTradeFIAT"))
		} else {
			minCost = this.SafeNumber(baseCurrency, OpAdd(MkString("minAmountTrade"), uppercaseQuoteId))
		}
		isActive := this.SafeInteger(market, MkString("active"))
		_ = isActive
		active := OpEq2(isActive, MkInteger(1))
		_ = active
		precision := MkMap(&VarMap{
			"amount": MkInteger(8),
			"price":  MkInteger(8),
		})
		_ = precision
		result.Push(MkMap(&VarMap{
			"id":        id,
			"symbol":    symbol,
			"base":      base,
			"quote":     quote,
			"baseId":    baseId,
			"quoteId":   quoteId,
			"active":    active,
			"precision": precision,
			"limits": MkMap(&VarMap{
				"amount": MkMap(&VarMap{
					"min": Math.Pow(MkInteger(10), OpNeg(*(precision).At(MkString("amount")))),
					"max": Math.Pow(MkInteger(10), *(precision).At(MkString("amount"))),
				}),
				"price": MkMap(&VarMap{
					"min": Math.Pow(MkInteger(10), OpNeg(*(precision).At(MkString("price")))),
					"max": Math.Pow(MkInteger(10), *(precision).At(MkString("price"))),
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

func (this *Braziliex) ParseTicker(goArgs ...*Variant) *Variant {
	ticker := GoGetArg(goArgs, 0, MkUndefined())
	_ = ticker
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(OpNotEq2(market, MkUndefined())) {
		symbol = *(market).At(MkString("symbol"))
	}
	timestamp := this.Milliseconds()
	_ = timestamp
	last := this.SafeNumber(ticker, MkString("last"))
	_ = last
	return MkMap(&VarMap{
		"symbol":        symbol,
		"timestamp":     timestamp,
		"datetime":      this.Iso8601(timestamp),
		"high":          this.SafeNumber(ticker, MkString("highestBid24")),
		"low":           this.SafeNumber(ticker, MkString("lowestAsk24")),
		"bid":           this.SafeNumber(ticker, MkString("highestBid")),
		"bidVolume":     MkUndefined(),
		"ask":           this.SafeNumber(ticker, MkString("lowestAsk")),
		"askVolume":     MkUndefined(),
		"vwap":          MkUndefined(),
		"open":          MkUndefined(),
		"close":         last,
		"last":          last,
		"previousClose": MkUndefined(),
		"change":        this.SafeNumber(ticker, MkString("percentChange")),
		"percentage":    MkUndefined(),
		"average":       MkUndefined(),
		"baseVolume":    this.SafeNumber(ticker, MkString("baseVolume24")),
		"quoteVolume":   this.SafeNumber(ticker, MkString("quoteVolume24")),
		"info":          ticker,
	})
}

func (this *Braziliex) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"market": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("publicGetTickerMarket"), this.Extend(request, params))
	_ = response
	return this.ParseTicker(response, market)
}

func (this *Braziliex) FetchTickers(goArgs ...*Variant) *Variant {
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
		marketId := *(ids).At(i)
		_ = marketId
		market := this.SafeMarket(marketId)
		_ = market
		symbol := *(market).At(MkString("symbol"))
		_ = symbol
		*(result).At(symbol) = this.ParseTicker(*(response).At(marketId), market)
	}
	return this.FilterByArray(result, MkString("symbol"), symbols)
}

func (this *Braziliex) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"market": this.MarketId(symbol)})
	_ = request
	response := this.Call(MkString("publicGetOrderbookMarket"), this.Extend(request, params))
	_ = response
	return this.ParseOrderBook(response, symbol, MkUndefined(), MkString("bids"), MkString("asks"), MkString("price"), MkString("amount"))
}

func (this *Braziliex) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.Parse8601(this.SafeString2(trade, MkString("date_exec"), MkString("date")))
	_ = timestamp
	priceString := this.SafeString(trade, MkString("price"))
	_ = priceString
	amountString := this.SafeString(trade, MkString("amount"))
	_ = amountString
	price := this.ParseNumber(priceString)
	_ = price
	amount := this.ParseNumber(amountString)
	_ = amount
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(OpNotEq2(market, MkUndefined())) {
		symbol = *(market).At(MkString("symbol"))
	}
	cost := this.SafeNumber(trade, MkString("total"))
	_ = cost
	if IsTrue(OpEq2(cost, MkUndefined())) {
		cost = this.ParseNumber(Precise.StringMul(priceString, amountString))
	}
	orderId := this.SafeString(trade, MkString("order_number"))
	_ = orderId
	type_ := MkString("limit")
	_ = type_
	side := this.SafeString(trade, MkString("type"))
	_ = side
	id := this.SafeString(trade, MkString("_id"))
	_ = id
	return MkMap(&VarMap{
		"id":           id,
		"info":         trade,
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       symbol,
		"order":        orderId,
		"type":         type_,
		"side":         side,
		"takerOrMaker": MkUndefined(),
		"price":        price,
		"amount":       amount,
		"cost":         cost,
		"fee":          MkUndefined(),
	})
}

func (this *Braziliex) FetchTrades(goArgs ...*Variant) *Variant {
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
	response := this.Call(MkString("publicGetTradehistoryMarket"), this.Extend(request, params))
	_ = response
	return this.ParseTrades(response, market, since, limit)
}

func (this *Braziliex) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	balances := this.Call(MkString("privatePostCompleteBalance"), params)
	_ = balances
	result := MkMap(&VarMap{"info": balances})
	_ = result
	currencyIds := GoGetKeys(balances)
	_ = currencyIds
	for i := MkInteger(0); IsTrue(OpLw(i, currencyIds.Length)); OpIncr(&i) {
		currencyId := *(currencyIds).At(i)
		_ = currencyId
		balance := *(balances).At(currencyId)
		_ = balance
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		account := this.Account()
		_ = account
		*(account).At(MkString("free")) = this.SafeString(balance, MkString("available"))
		*(account).At(MkString("total")) = this.SafeString(balance, MkString("total"))
		*(result).At(code) = OpCopy(account)
	}
	return this.ParseBalance(result)
}

func (this *Braziliex) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	marketId := this.SafeString(order, MkString("market"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market, MkString("_"))
	_ = symbol
	timestamp := this.SafeInteger(order, MkString("timestamp"))
	_ = timestamp
	if IsTrue(OpEq2(timestamp, MkUndefined())) {
		timestamp = this.Parse8601(this.SafeString(order, MkString("date")))
	}
	price := this.SafeNumber(order, MkString("price"))
	_ = price
	cost := this.SafeNumber(order, MkString("total"))
	_ = cost
	amount := this.SafeNumber(order, MkString("amount"))
	_ = amount
	filledPercentage := this.SafeNumber(order, MkString("progress"))
	_ = filledPercentage
	filled := OpMulti(amount, filledPercentage)
	_ = filled
	id := this.SafeString(order, MkString("order_number"))
	_ = id
	fee := this.SafeValue(order, MkString("fee"))
	_ = fee
	status := OpTrinary(OpEq2(filledPercentage, MkNumber(1.0)), MkString("closed"), MkString("open"))
	_ = status
	side := this.SafeString(order, MkString("type"))
	_ = side
	return this.SafeOrder(MkMap(&VarMap{
		"id":                 id,
		"clientOrderId":      MkUndefined(),
		"datetime":           this.Iso8601(timestamp),
		"timestamp":          timestamp,
		"lastTradeTimestamp": MkUndefined(),
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
		"remaining":          MkUndefined(),
		"trades":             MkUndefined(),
		"fee":                fee,
		"info":               order,
		"average":            MkUndefined(),
	}))
}

func (this *Braziliex) CreateOrder(goArgs ...*Variant) *Variant {
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
	method := OpAdd(MkString("privatePost"), this.Capitalize(side))
	_ = method
	request := MkMap(&VarMap{
		"market": *(market).At(MkString("id")),
		"price":  price,
		"amount": amount,
	})
	_ = request
	response := this.Call(method, this.Extend(request, params))
	_ = response
	success := this.SafeInteger(response, MkString("success"))
	_ = success
	if IsTrue(OpNotEq2(success, MkInteger(1))) {
		panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), this.Json(response)))))
	}
	message := this.SafeString(response, MkString("message"))
	_ = message
	parts := message.Split(MkString(" / "))
	_ = parts
	parts = parts.Slice(MkInteger(1))
	feeParts := (*(parts).At(MkInteger(5))).Call(MkString("split"), MkString(" "))
	_ = feeParts
	amountParts := (*(parts).At(MkInteger(2))).Call(MkString("split"), MkString(" "))
	_ = amountParts
	priceParts := (*(parts).At(MkInteger(3))).Call(MkString("split"), MkString(" "))
	_ = priceParts
	totalParts := (*(parts).At(MkInteger(4))).Call(MkString("split"), MkString(" "))
	_ = totalParts
	order := this.ParseOrder(MkMap(&VarMap{
		"timestamp":    this.Milliseconds(),
		"order_number": *(response).At(MkString("order_number")),
		"type":         this.SafeStringLower(parts, MkInteger(0)),
		"market":       (*(parts).At(MkInteger(0))).Call(MkString("toLowerCase")),
		"amount":       this.SafeString(amountParts, MkInteger(1)),
		"price":        this.SafeString(priceParts, MkInteger(1)),
		"total":        this.SafeString(totalParts, MkInteger(1)),
		"fee": MkMap(&VarMap{
			"cost":     this.SafeNumber(feeParts, MkInteger(1)),
			"currency": this.SafeString(feeParts, MkInteger(2)),
		}),
		"progress": MkString("0.0"),
		"info":     response,
	}), market)
	_ = order
	return order
}

func (this *Braziliex) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{
		"order_number": id,
		"market":       *(market).At(MkString("id")),
	})
	_ = request
	return this.Call(MkString("privatePostCancelOrder"), this.Extend(request, params))
}

func (this *Braziliex) FetchOrder(goArgs ...*Variant) *Variant {
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
		"order_number": id,
		"market":       *(market).At(MkString("id")),
	})
	_ = request
	response := this.Call(MkString("privatePostOrderStatus"), this.Extend(request, params))
	_ = response
	return this.ParseOrder(response, market)
}

func (this *Braziliex) FetchOpenOrders(goArgs ...*Variant) *Variant {
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
	response := this.Call(MkString("privatePostOpenOrders"), this.Extend(request, params))
	_ = response
	orders := this.SafeValue(response, MkString("order_open"), MkArray(&VarArray{}))
	_ = orders
	return this.ParseOrders(orders, market, since, limit)
}

func (this *Braziliex) FetchMyTrades(goArgs ...*Variant) *Variant {
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
	response := this.Call(MkString("privatePostTradeHistory"), this.Extend(request, params))
	_ = response
	trades := this.SafeValue(response, MkString("trade_history"), MkArray(&VarArray{}))
	_ = trades
	return this.ParseTrades(trades, market, since, limit)
}

func (this *Braziliex) FetchDepositAddress(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{"currency": *(currency).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("privatePostDepositAddress"), this.Extend(request, params))
	_ = response
	address := this.SafeString(response, MkString("deposit_address"))
	_ = address
	this.CheckAddress(address)
	tag := this.SafeString(response, MkString("payment_id"))
	_ = tag
	return MkMap(&VarMap{
		"currency": code,
		"address":  address,
		"tag":      tag,
		"info":     response,
	})
}

func (this *Braziliex) Sign(goArgs ...*Variant) *Variant {
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
	url := OpAdd(*(*this.At(MkString("urls"))).At(MkString("api")), OpAdd(MkString("/"), api))
	_ = url
	query := this.Omit(params, this.ExtractParams(path))
	_ = query
	if IsTrue(OpEq2(api, MkString("public"))) {
		OpAddAssign(&url, OpAdd(MkString("/"), this.ImplodeParams(path, params)))
		if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
			OpAddAssign(&url, OpAdd(MkString("?"), this.Urlencode(query)))
		}
	} else {
		this.CheckRequiredCredentials()
		query = this.Extend(MkMap(&VarMap{
			"command": path,
			"nonce":   this.Nonce(),
		}), query)
		body = this.Urlencode(query)
		signature := this.Hmac(this.Encode(body), this.Encode(*this.At(MkString("secret"))), MkString("sha512"))
		_ = signature
		headers = MkMap(&VarMap{
			"Content-type": MkString("application/x-www-form-urlencoded"),
			"Key":          *this.At(MkString("apiKey")),
			"Sign":         signature,
		})
	}
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Braziliex) Request(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpAnd(OpEq2(OpType(response), MkString("string")), OpLw(response.Length, MkInteger(1)))) {
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), MkString(" returned empty response"))))
	}
	if IsTrue(OpHasMember(MkString("success"), response)) {
		success := this.SafeInteger(response, MkString("success"))
		_ = success
		if IsTrue(OpEq2(success, MkInteger(0))) {
			message := this.SafeString(response, MkString("message"))
			_ = message
			if IsTrue(OpEq2(message, MkString("Invalid APIKey"))) {
				panic(NewAuthenticationError(message))
			}
			panic(NewExchangeError(message))
		}
	}
	return response
}
