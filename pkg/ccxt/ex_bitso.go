package ccxt

import ()

type Bitso struct {
	*ExchangeBase
}

var _ Exchange = (*Bitso)(nil)

func init() {
	exchange := &Bitso{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Bitso) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("bitso"),
		"name": MkString("Bitso"),
		"countries": MkArray(&VarArray{
			MkString("MX"),
		}),
		"rateLimit": MkInteger(2000),
		"version":   MkString("v3"),
		"has": MkMap(&VarMap{
			"cancelOrder":         MkBool(true),
			"CORS":                MkBool(false),
			"createOrder":         MkBool(true),
			"fetchBalance":        MkBool(true),
			"fetchDepositAddress": MkBool(true),
			"fetchMarkets":        MkBool(true),
			"fetchMyTrades":       MkBool(true),
			"fetchOpenOrders":     MkBool(true),
			"fetchOrder":          MkBool(true),
			"fetchOrderBook":      MkBool(true),
			"fetchOrderTrades":    MkBool(true),
			"fetchTicker":         MkBool(true),
			"fetchTrades":         MkBool(true),
			"withdraw":            MkBool(true),
		}),
		"urls": MkMap(&VarMap{
			"logo":     MkString("https://user-images.githubusercontent.com/51840849/87295554-11f98280-c50e-11ea-80d6-15b3bafa8cbf.jpg"),
			"api":      MkString("https://api.bitso.com"),
			"www":      MkString("https://bitso.com"),
			"doc":      MkString("https://bitso.com/api_info"),
			"fees":     MkString("https://bitso.com/fees"),
			"referral": MkString("https://bitso.com/?ref=itej"),
		}),
		"precisionMode": TICK_SIZE,
		"options": MkMap(&VarMap{
			"precision": MkMap(&VarMap{
				"XRP":  MkNumber(0.000001),
				"MXN":  MkNumber(0.01),
				"TUSD": MkNumber(0.01),
			}),
			"defaultPrecision": MkNumber(0.00000001),
		}),
		"api": MkMap(&VarMap{
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("available_books"),
				MkString("ticker"),
				MkString("order_book"),
				MkString("trades"),
			})}),
			"private": MkMap(&VarMap{
				"get": MkArray(&VarArray{
					MkString("account_status"),
					MkString("balance"),
					MkString("fees"),
					MkString("fundings"),
					MkString("fundings/{fid}"),
					MkString("funding_destination"),
					MkString("kyc_documents"),
					MkString("ledger"),
					MkString("ledger/trades"),
					MkString("ledger/fees"),
					MkString("ledger/fundings"),
					MkString("ledger/withdrawals"),
					MkString("mx_bank_codes"),
					MkString("open_orders"),
					MkString("order_trades/{oid}"),
					MkString("orders/{oid}"),
					MkString("user_trades"),
					MkString("user_trades/{tid}"),
					MkString("withdrawals/"),
					MkString("withdrawals/{wid}"),
				}),
				"post": MkArray(&VarArray{
					MkString("bitcoin_withdrawal"),
					MkString("debit_card_withdrawal"),
					MkString("ether_withdrawal"),
					MkString("ripple_withdrawal"),
					MkString("bcash_withdrawal"),
					MkString("litecoin_withdrawal"),
					MkString("orders"),
					MkString("phone_number"),
					MkString("phone_verification"),
					MkString("phone_withdrawal"),
					MkString("spei_withdrawal"),
					MkString("ripple_withdrawal"),
					MkString("bcash_withdrawal"),
					MkString("litecoin_withdrawal"),
				}),
				"delete": MkArray(&VarArray{
					MkString("orders/{oid}"),
					MkString("orders/all"),
				}),
			}),
		}),
		"exceptions": MkMap(&VarMap{
			"0201": AuthenticationError,
			"104":  InvalidNonce,
		}),
	}))
}

func (this *Bitso) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetAvailableBooks"), params)
	_ = response
	markets := this.SafeValue(response, MkString("payload"))
	_ = markets
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, markets.Length)); OpIncr(&i) {
		market := *(markets).At(i)
		_ = market
		id := this.SafeString(market, MkString("book"))
		_ = id
		baseId := MkUndefined()
		_ = baseId
		quoteId := MkUndefined()
		_ = quoteId
		ArrayUnpack(id.Split(MkString("_")), &baseId, &quoteId)
		base := baseId.ToUpperCase()
		_ = base
		quote := quoteId.ToUpperCase()
		_ = quote
		base = this.SafeCurrencyCode(base)
		quote = this.SafeCurrencyCode(quote)
		symbol := OpAdd(base, OpAdd(MkString("/"), quote))
		_ = symbol
		limits := MkMap(&VarMap{
			"amount": MkMap(&VarMap{
				"min": this.SafeNumber(market, MkString("minimum_amount")),
				"max": this.SafeNumber(market, MkString("maximum_amount")),
			}),
			"price": MkMap(&VarMap{
				"min": this.SafeNumber(market, MkString("minimum_price")),
				"max": this.SafeNumber(market, MkString("maximum_price")),
			}),
			"cost": MkMap(&VarMap{
				"min": this.SafeNumber(market, MkString("minimum_value")),
				"max": this.SafeNumber(market, MkString("maximum_value")),
			}),
		})
		_ = limits
		defaultPricePrecision := this.SafeNumber(*(*this.At(MkString("options"))).At(MkString("precision")), quote, *(*this.At(MkString("options"))).At(MkString("defaultPrecision")))
		_ = defaultPricePrecision
		pricePrecision := this.SafeNumber(market, MkString("tick_size"), defaultPricePrecision)
		_ = pricePrecision
		precision := MkMap(&VarMap{
			"amount": this.SafeNumber(*(*this.At(MkString("options"))).At(MkString("precision")), base, *(*this.At(MkString("options"))).At(MkString("defaultPrecision"))),
			"price":  pricePrecision,
		})
		_ = precision
		fees := this.SafeValue(market, MkString("fees"), MkMap(&VarMap{}))
		_ = fees
		flatRate := this.SafeValue(fees, MkString("flat_rate"), MkMap(&VarMap{}))
		_ = flatRate
		makerString := this.SafeString(flatRate, MkString("maker"))
		_ = makerString
		takerString := this.SafeString(flatRate, MkString("taker"))
		_ = takerString
		maker := this.ParseNumber(Precise.StringDiv(makerString, MkString("100")))
		_ = maker
		taker := this.ParseNumber(Precise.StringDiv(takerString, MkString("100")))
		_ = taker
		feeTiers := this.SafeValue(fees, MkString("structure"), MkArray(&VarArray{}))
		_ = feeTiers
		fee := MkMap(&VarMap{
			"taker":      taker,
			"maker":      maker,
			"percentage": MkBool(true),
			"tierBased":  MkBool(true),
		})
		_ = fee
		takerFees := MkArray(&VarArray{})
		_ = takerFees
		makerFees := MkArray(&VarArray{})
		_ = makerFees
		for j := MkInteger(0); IsTrue(OpLw(j, feeTiers.Length)); OpIncr(&j) {
			tier := *(feeTiers).At(j)
			_ = tier
			volume := this.SafeNumber(tier, MkString("volume"))
			_ = volume
			takerFee := this.SafeNumber(tier, MkString("taker"))
			_ = takerFee
			makerFee := this.SafeNumber(tier, MkString("maker"))
			_ = makerFee
			takerFees.Push(MkArray(&VarArray{
				volume,
				takerFee,
			}))
			makerFees.Push(MkArray(&VarArray{
				volume,
				makerFee,
			}))
			if IsTrue(OpEq2(j, MkInteger(0))) {
				*(fee).At(MkString("taker")) = OpCopy(takerFee)
				*(fee).At(MkString("maker")) = OpCopy(makerFee)
			}
		}
		tiers := MkMap(&VarMap{
			"taker": takerFees,
			"maker": makerFees,
		})
		_ = tiers
		*(fee).At(MkString("tiers")) = OpCopy(tiers)
		result.Push(this.Extend(MkMap(&VarMap{
			"id":        id,
			"symbol":    symbol,
			"base":      base,
			"quote":     quote,
			"baseId":    baseId,
			"quoteId":   quoteId,
			"info":      market,
			"limits":    limits,
			"precision": precision,
			"active":    MkUndefined(),
		}), fee))
	}
	return result
}

func (this *Bitso) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privateGetBalance"), params)
	_ = response
	payload := this.SafeValue(response, MkString("payload"), MkMap(&VarMap{}))
	_ = payload
	balances := this.SafeValue(payload, MkString("balances"))
	_ = balances
	result := MkMap(&VarMap{
		"info":      response,
		"timestamp": MkUndefined(),
		"datetime":  MkUndefined(),
	})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, balances.Length)); OpIncr(&i) {
		balance := *(balances).At(i)
		_ = balance
		currencyId := this.SafeString(balance, MkString("currency"))
		_ = currencyId
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		account := this.Account()
		_ = account
		*(account).At(MkString("free")) = this.SafeString(balance, MkString("available"))
		*(account).At(MkString("used")) = this.SafeString(balance, MkString("locked"))
		*(account).At(MkString("total")) = this.SafeString(balance, MkString("total"))
		*(result).At(code) = OpCopy(account)
	}
	return this.ParseBalance(result)
}

func (this *Bitso) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"book": this.MarketId(symbol)})
	_ = request
	response := this.Call(MkString("publicGetOrderBook"), this.Extend(request, params))
	_ = response
	orderbook := this.SafeValue(response, MkString("payload"))
	_ = orderbook
	timestamp := this.Parse8601(this.SafeString(orderbook, MkString("updated_at")))
	_ = timestamp
	return this.ParseOrderBook(orderbook, symbol, timestamp, MkString("bids"), MkString("asks"), MkString("price"), MkString("amount"))
}

func (this *Bitso) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"book": this.MarketId(symbol)})
	_ = request
	response := this.Call(MkString("publicGetTicker"), this.Extend(request, params))
	_ = response
	ticker := this.SafeValue(response, MkString("payload"))
	_ = ticker
	timestamp := this.Parse8601(this.SafeString(ticker, MkString("created_at")))
	_ = timestamp
	vwap := this.SafeNumber(ticker, MkString("vwap"))
	_ = vwap
	baseVolume := this.SafeNumber(ticker, MkString("volume"))
	_ = baseVolume
	quoteVolume := OpCopy(MkUndefined())
	_ = quoteVolume
	if IsTrue(OpAnd(OpNotEq2(baseVolume, MkUndefined()), OpNotEq2(vwap, MkUndefined()))) {
		quoteVolume = OpMulti(baseVolume, vwap)
	}
	last := this.SafeNumber(ticker, MkString("last"))
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
		"vwap":          vwap,
		"open":          MkUndefined(),
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

func (this *Bitso) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.Parse8601(this.SafeString(trade, MkString("created_at")))
	_ = timestamp
	marketId := this.SafeString(trade, MkString("book"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market, MkString("_"))
	_ = symbol
	side := this.SafeString2(trade, MkString("side"), MkString("maker_side"))
	_ = side
	amount := this.SafeNumber2(trade, MkString("amount"), MkString("major"))
	_ = amount
	if IsTrue(OpNotEq2(amount, MkUndefined())) {
		amount = Math.Abs(amount)
	}
	fee := OpCopy(MkUndefined())
	_ = fee
	feeCost := this.SafeNumber(trade, MkString("fees_amount"))
	_ = feeCost
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		feeCurrencyId := this.SafeString(trade, MkString("fees_currency"))
		_ = feeCurrencyId
		feeCurrency := this.SafeCurrencyCode(feeCurrencyId)
		_ = feeCurrency
		fee = MkMap(&VarMap{
			"cost":     feeCost,
			"currency": feeCurrency,
		})
	}
	cost := this.SafeNumber(trade, MkString("minor"))
	_ = cost
	if IsTrue(OpNotEq2(cost, MkUndefined())) {
		cost = Math.Abs(cost)
	}
	price := this.SafeNumber(trade, MkString("price"))
	_ = price
	orderId := this.SafeString(trade, MkString("oid"))
	_ = orderId
	id := this.SafeString(trade, MkString("tid"))
	_ = id
	return MkMap(&VarMap{
		"id":           id,
		"info":         trade,
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       symbol,
		"order":        orderId,
		"type":         MkUndefined(),
		"side":         side,
		"takerOrMaker": MkUndefined(),
		"price":        price,
		"amount":       amount,
		"cost":         cost,
		"fee":          fee,
	})
}

func (this *Bitso) FetchTrades(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"book": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("publicGetTrades"), this.Extend(request, params))
	_ = response
	return this.ParseTrades(*(response).At(MkString("payload")), market, since, limit)
}

func (this *Bitso) FetchMyTrades(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkInteger(25))
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	markerInParams := OpHasMember(MkString("marker"), params)
	_ = markerInParams
	if IsTrue(OpAnd(OpNotEq2(since, MkUndefined()), OpNot(markerInParams))) {
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), MkString(" fetchMyTrades does not support fetching trades starting from a timestamp with the `since` argument, use the `marker` extra param to filter starting from an integer trade id"))))
	}
	if IsTrue(markerInParams) {
		params = this.Extend(params, MkMap(&VarMap{"marker": parseInt(*(params).At(MkString("marker")))}))
	}
	request := MkMap(&VarMap{
		"book":  *(market).At(MkString("id")),
		"limit": limit,
	})
	_ = request
	response := this.Call(MkString("privateGetUserTrades"), this.Extend(request, params))
	_ = response
	return this.ParseTrades(*(response).At(MkString("payload")), market, since, limit)
}

func (this *Bitso) CreateOrder(goArgs ...*Variant) *Variant {
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
		"book":  this.MarketId(symbol),
		"side":  side,
		"type":  type_,
		"major": this.AmountToPrecision(symbol, amount),
	})
	_ = request
	if IsTrue(OpEq2(type_, MkString("limit"))) {
		*(request).At(MkString("price")) = this.PriceToPrecision(symbol, price)
	}
	response := this.Call(MkString("privatePostOrders"), this.Extend(request, params))
	_ = response
	id := this.SafeString(*(response).At(MkString("payload")), MkString("oid"))
	_ = id
	return MkMap(&VarMap{
		"info": response,
		"id":   id,
	})
}

func (this *Bitso) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"oid": id})
	_ = request
	return this.Call(MkString("privateDeleteOrdersOid"), this.Extend(request, params))
}

func (this *Bitso) ParseOrderStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"partial-fill": MkString("open"),
		"completed":    MkString("closed"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Bitso) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	id := this.SafeString(order, MkString("oid"))
	_ = id
	side := this.SafeString(order, MkString("side"))
	_ = side
	status := this.ParseOrderStatus(this.SafeString(order, MkString("status")))
	_ = status
	marketId := this.SafeString(order, MkString("book"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market, MkString("_"))
	_ = symbol
	orderType := this.SafeString(order, MkString("type"))
	_ = orderType
	timestamp := this.Parse8601(this.SafeString(order, MkString("created_at")))
	_ = timestamp
	price := this.SafeNumber(order, MkString("price"))
	_ = price
	amount := this.SafeNumber(order, MkString("original_amount"))
	_ = amount
	remaining := this.SafeNumber(order, MkString("unfilled_amount"))
	_ = remaining
	clientOrderId := this.SafeString(order, MkString("client_id"))
	_ = clientOrderId
	return this.SafeOrder(MkMap(&VarMap{
		"info":               order,
		"id":                 id,
		"clientOrderId":      clientOrderId,
		"timestamp":          timestamp,
		"datetime":           this.Iso8601(timestamp),
		"lastTradeTimestamp": MkUndefined(),
		"symbol":             symbol,
		"type":               orderType,
		"timeInForce":        MkUndefined(),
		"postOnly":           MkUndefined(),
		"side":               side,
		"price":              price,
		"stopPrice":          MkUndefined(),
		"amount":             amount,
		"cost":               MkUndefined(),
		"remaining":          remaining,
		"filled":             MkUndefined(),
		"status":             status,
		"fee":                MkUndefined(),
		"average":            MkUndefined(),
		"trades":             MkUndefined(),
	}))
}

func (this *Bitso) FetchOpenOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkInteger(25))
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	markerInParams := OpHasMember(MkString("marker"), params)
	_ = markerInParams
	if IsTrue(OpAnd(OpNotEq2(since, MkUndefined()), OpNot(markerInParams))) {
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), MkString(" fetchOpenOrders does not support fetching orders starting from a timestamp with the `since` argument, use the `marker` extra param to filter starting from an integer trade id"))))
	}
	if IsTrue(markerInParams) {
		params = this.Extend(params, MkMap(&VarMap{"marker": parseInt(*(params).At(MkString("marker")))}))
	}
	request := MkMap(&VarMap{
		"book":  *(market).At(MkString("id")),
		"limit": limit,
	})
	_ = request
	response := this.Call(MkString("privateGetOpenOrders"), this.Extend(request, params))
	_ = response
	orders := this.ParseOrders(*(response).At(MkString("payload")), market, since, limit)
	_ = orders
	return orders
}

func (this *Bitso) FetchOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privateGetOrdersOid"), MkMap(&VarMap{"oid": id}))
	_ = response
	payload := this.SafeValue(response, MkString("payload"))
	_ = payload
	if IsTrue(GoIsArray(payload)) {
		numOrders := *(*(response).At(MkString("payload"))).At(MkString("length"))
		_ = numOrders
		if IsTrue(OpEq2(numOrders, MkInteger(1))) {
			return this.ParseOrder(*(payload).At(MkInteger(0)))
		}
	}
	panic(NewOrderNotFound(OpAdd(*this.At(MkString("id")), OpAdd(MkString(": The order "), OpAdd(id, MkString(" not found."))))))
	return MkUndefined()
}

func (this *Bitso) FetchOrderTrades(goArgs ...*Variant) *Variant {
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
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"oid": id})
	_ = request
	response := this.Call(MkString("privateGetOrderTradesOid"), this.Extend(request, params))
	_ = response
	return this.ParseTrades(*(response).At(MkString("payload")), market)
}

func (this *Bitso) FetchDepositAddress(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{"fund_currency": *(currency).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("privateGetFundingDestination"), this.Extend(request, params))
	_ = response
	address := this.SafeString(*(response).At(MkString("payload")), MkString("account_identifier"))
	_ = address
	tag := OpCopy(MkUndefined())
	_ = tag
	if IsTrue(OpGtEq(address.IndexOf(MkString("?dt=")), MkInteger(0))) {
		parts := address.Split(MkString("?dt="))
		_ = parts
		address = this.SafeString(parts, MkInteger(0))
		tag = this.SafeString(parts, MkInteger(1))
	}
	this.CheckAddress(address)
	return MkMap(&VarMap{
		"currency": code,
		"address":  address,
		"tag":      tag,
		"info":     response,
	})
}

func (this *Bitso) Withdraw(goArgs ...*Variant) *Variant {
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
	methods := MkMap(&VarMap{
		"BTC": MkString("Bitcoin"),
		"ETH": MkString("Ether"),
		"XRP": MkString("Ripple"),
		"BCH": MkString("Bcash"),
		"LTC": MkString("Litecoin"),
	})
	_ = methods
	method := OpTrinary(OpHasMember(code, methods), *(methods).At(code), MkUndefined())
	_ = method
	if IsTrue(OpEq2(method, MkUndefined())) {
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" not valid withdraw coin: "), code))))
	}
	request := MkMap(&VarMap{
		"amount":          amount,
		"address":         address,
		"destination_tag": tag,
	})
	_ = request
	classMethod := OpAdd(MkString("privatePost"), OpAdd(method, MkString("Withdrawal")))
	_ = classMethod
	response := this.Call(classMethod, this.Extend(request, params))
	_ = response
	return MkMap(&VarMap{
		"info": response,
		"id":   this.SafeString(*(response).At(MkString("payload")), MkString("wid")),
	})
}

func (this *Bitso) Sign(goArgs ...*Variant) *Variant {
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
	endpoint := OpAdd(MkString("/"), OpAdd(*this.At(MkString("version")), OpAdd(MkString("/"), this.ImplodeParams(path, params))))
	_ = endpoint
	query := this.Omit(params, this.ExtractParams(path))
	_ = query
	if IsTrue(OpEq2(method, MkString("GET"))) {
		if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
			OpAddAssign(&endpoint, OpAdd(MkString("?"), this.Urlencode(query)))
		}
	}
	url := OpAdd(*(*this.At(MkString("urls"))).At(MkString("api")), endpoint)
	_ = url
	if IsTrue(OpEq2(api, MkString("private"))) {
		this.CheckRequiredCredentials()
		nonce := (this.Nonce()).Call(MkString("toString"))
		_ = nonce
		request := (MkArray(&VarArray{
			nonce,
			method,
			endpoint,
		})).Call(MkString("join"), MkString(""))
		_ = request
		if IsTrue(OpNotEq2(method, MkString("GET"))) {
			if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
				body = this.Json(query)
				OpAddAssign(&request, body)
			}
		}
		signature := this.Hmac(this.Encode(request), this.Encode(*this.At(MkString("secret"))))
		_ = signature
		auth := OpAdd(*this.At(MkString("apiKey")), OpAdd(MkString(":"), OpAdd(nonce, OpAdd(MkString(":"), signature))))
		_ = auth
		headers = MkMap(&VarMap{
			"Authorization": OpAdd(MkString("Bitso "), auth),
			"Content-Type":  MkString("application/json"),
		})
	}
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Bitso) HandleErrors(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpHasMember(MkString("success"), response)) {
		success := this.SafeValue(response, MkString("success"), MkBool(false))
		_ = success
		if IsTrue(OpEq2(OpType(success), MkString("string"))) {
			if IsTrue(OpOr(OpEq2(success, MkString("true")), OpEq2(success, MkString("1")))) {
				success = OpCopy(MkBool(true))
			} else {
				success = OpCopy(MkBool(false))
			}
		}
		if IsTrue(OpNot(success)) {
			feedback := OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), this.Json(response)))
			_ = feedback
			error := this.SafeValue(response, MkString("error"))
			_ = error
			if IsTrue(OpEq2(error, MkUndefined())) {
				panic(NewExchangeError(feedback))
			}
			code := this.SafeString(error, MkString("code"))
			_ = code
			this.ThrowExactlyMatchedException(*this.At(MkString("exceptions")), code, feedback)
			panic(NewExchangeError(feedback))
		}
	}
	return MkUndefined()
}
