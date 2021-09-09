package ccxt

import ()

type Gemini struct {
	*ExchangeBase
}

var _ Exchange = (*Gemini)(nil)

func init() {
	exchange := &Gemini{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Gemini) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("gemini"),
		"name": MkString("Gemini"),
		"countries": MkArray(&VarArray{
			MkString("US"),
		}),
		"rateLimit": MkInteger(1500),
		"version":   MkString("v1"),
		"has": MkMap(&VarMap{
			"cancelOrder":          MkBool(true),
			"CORS":                 MkBool(false),
			"createDepositAddress": MkBool(true),
			"createMarketOrder":    MkBool(false),
			"createOrder":          MkBool(true),
			"fetchBalance":         MkBool(true),
			"fetchBidsAsks":        MkBool(false),
			"fetchClosedOrders":    MkBool(false),
			"fetchDepositAddress":  MkBool(false),
			"fetchDeposits":        MkBool(false),
			"fetchMarkets":         MkBool(true),
			"fetchMyTrades":        MkBool(true),
			"fetchOHLCV":           MkBool(true),
			"fetchOpenOrders":      MkBool(true),
			"fetchOrder":           MkBool(true),
			"fetchOrderBook":       MkBool(true),
			"fetchOrders":          MkBool(false),
			"fetchTicker":          MkBool(true),
			"fetchTickers":         MkBool(true),
			"fetchTrades":          MkBool(true),
			"fetchTransactions":    MkBool(true),
			"fetchWithdrawals":     MkBool(false),
			"withdraw":             MkBool(true),
		}),
		"urls": MkMap(&VarMap{
			"logo": MkString("https://user-images.githubusercontent.com/1294454/27816857-ce7be644-6096-11e7-82d6-3c257263229c.jpg"),
			"api": MkMap(&VarMap{
				"public":  MkString("https://api.gemini.com"),
				"private": MkString("https://api.gemini.com"),
				"web":     MkString("https://docs.gemini.com"),
			}),
			"www": MkString("https://gemini.com/"),
			"doc": MkArray(&VarArray{
				MkString("https://docs.gemini.com/rest-api"),
				MkString("https://docs.sandbox.gemini.com"),
			}),
			"test": MkMap(&VarMap{
				"public":  MkString("https://api.sandbox.gemini.com"),
				"private": MkString("https://api.sandbox.gemini.com"),
				"web":     MkString("https://docs.gemini.com"),
			}),
			"fees": MkArray(&VarArray{
				MkString("https://gemini.com/api-fee-schedule"),
				MkString("https://gemini.com/trading-fees"),
				MkString("https://gemini.com/transfer-fees"),
			}),
		}),
		"api": MkMap(&VarMap{
			"web": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("rest-api"),
			})}),
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("v1/symbols"),
				MkString("v1/pricefeed"),
				MkString("v1/pubticker/{symbol}"),
				MkString("v1/book/{symbol}"),
				MkString("v1/trades/{symbol}"),
				MkString("v1/auction/{symbol}"),
				MkString("v1/auction/{symbol}/history"),
				MkString("v2/candles/{symbol}/{timeframe}"),
				MkString("v2/ticker/{symbol}"),
			})}),
			"private": MkMap(&VarMap{"post": MkArray(&VarArray{
				MkString("v1/order/new"),
				MkString("v1/order/cancel"),
				MkString("v1/order/cancel/session"),
				MkString("v1/order/cancel/all"),
				MkString("v1/order/status"),
				MkString("v1/orders"),
				MkString("v1/mytrades"),
				MkString("v1/notionalvolume"),
				MkString("v1/tradevolume"),
				MkString("v1/transfers"),
				MkString("v1/balances"),
				MkString("v1/deposit/{currency}/newAddress"),
				MkString("v1/withdraw/{currency}"),
				MkString("v1/heartbeat"),
				MkString("v1/transfers"),
			})}),
		}),
		"precisionMode": TICK_SIZE,
		"fees": MkMap(&VarMap{"trading": MkMap(&VarMap{
			"taker": MkNumber(0.0035),
			"maker": MkNumber(0.001),
		})}),
		"httpExceptions": MkMap(&VarMap{
			"400": BadRequest,
			"403": PermissionDenied,
			"404": OrderNotFound,
			"406": InsufficientFunds,
			"429": RateLimitExceeded,
			"500": ExchangeError,
			"502": ExchangeNotAvailable,
			"503": OnMaintenance,
		}),
		"timeframes": MkMap(&VarMap{
			"1m":  MkString("1m"),
			"5m":  MkString("5m"),
			"15m": MkString("15m"),
			"30m": MkString("30m"),
			"1h":  MkString("1hr"),
			"6h":  MkString("6hr"),
			"1d":  MkString("1day"),
		}),
		"exceptions": MkMap(&VarMap{
			"exact": MkMap(&VarMap{
				"AuctionNotOpen":            BadRequest,
				"ClientOrderIdTooLong":      BadRequest,
				"ClientOrderIdMustBeString": BadRequest,
				"ConflictingOptions":        BadRequest,
				"EndpointMismatch":          BadRequest,
				"EndpointNotFound":          BadRequest,
				"IneligibleTiming":          BadRequest,
				"InsufficientFunds":         InsufficientFunds,
				"InvalidJson":               BadRequest,
				"InvalidNonce":              InvalidNonce,
				"InvalidOrderType":          InvalidOrder,
				"InvalidPrice":              InvalidOrder,
				"InvalidQuantity":           InvalidOrder,
				"InvalidSide":               InvalidOrder,
				"InvalidSignature":          AuthenticationError,
				"InvalidSymbol":             BadRequest,
				"InvalidTimestampInPayload": BadRequest,
				"Maintenance":               OnMaintenance,
				"MarketNotOpen":             InvalidOrder,
				"MissingApikeyHeader":       AuthenticationError,
				"MissingOrderField":         InvalidOrder,
				"MissingRole":               AuthenticationError,
				"MissingPayloadHeader":      AuthenticationError,
				"MissingSignatureHeader":    AuthenticationError,
				"NoSSL":                     AuthenticationError,
				"OptionsMustBeArray":        BadRequest,
				"OrderNotFound":             OrderNotFound,
				"RateLimit":                 RateLimitExceeded,
				"System":                    ExchangeError,
				"UnsupportedOption":         BadRequest,
			}),
			"broad": MkMap(&VarMap{
				"The Gemini Exchange is currently undergoing maintenance.":        OnMaintenance,
				"We are investigating technical issues with the Gemini Exchange.": ExchangeNotAvailable,
			}),
		}),
		"options": MkMap(&VarMap{
			"fetchMarketsMethod": MkString("fetch_markets_from_web"),
			"fetchTickerMethod":  MkString("fetchTickerV1"),
		}),
	}))
}

func (this *Gemini) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	method := this.SafeValue(*this.At(MkString("options")), MkString("fetchMarketsMethod"), MkString("fetch_markets_from_api"))
	_ = method
	return this.Call(method, params)
}

func (this *Gemini) FetchMarketsFromWeb(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("webGetRestApi"), params)
	_ = response
	sections := response.Split(MkString("<h1 id=\"symbols-and-minimums\">Symbols and minimums</h1>"))
	_ = sections
	numSections := OpCopy(sections.Length)
	_ = numSections
	error := OpAdd(*this.At(MkString("id")), OpAdd(MkString(" the "), OpAdd(*this.At(MkString("name")), OpAdd(MkString(" API doc HTML markup has changed, breaking the parser of order limits and precision info for "), OpAdd(*this.At(MkString("name")), MkString(" markets."))))))
	_ = error
	if IsTrue(OpNotEq2(numSections, MkInteger(2))) {
		panic(NewNotSupported(error))
	}
	tables := (*(sections).At(MkInteger(1))).Call(MkString("split"), MkString("tbody>"))
	_ = tables
	numTables := OpCopy(tables.Length)
	_ = numTables
	if IsTrue(OpLw(numTables, MkInteger(2))) {
		panic(NewNotSupported(error))
	}
	rows := (*(tables).At(MkInteger(1))).Call(MkString("split"), MkString("\n<tr>\n"))
	_ = rows
	numRows := OpCopy(rows.Length)
	_ = numRows
	if IsTrue(OpLw(numRows, MkInteger(2))) {
		panic(NewNotSupported(error))
	}
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(1); IsTrue(OpLw(i, numRows)); OpIncr(&i) {
		row := *(rows).At(i)
		_ = row
		cells := row.Split(MkString("</td>\n"))
		_ = cells
		numCells := OpCopy(cells.Length)
		_ = numCells
		if IsTrue(OpLw(numCells, MkInteger(5))) {
			panic(NewNotSupported(error))
		}
		marketId := (*(cells).At(MkInteger(0))).Call(MkString("replace"), MkString("<td>"), MkString(""))
		_ = marketId
		minAmountString := (*(cells).At(MkInteger(1))).Call(MkString("replace"), MkString("<td>"), MkString(""))
		_ = minAmountString
		minAmountParts := minAmountString.Split(MkString(" "))
		_ = minAmountParts
		minAmount := this.SafeNumber(minAmountParts, MkInteger(0))
		_ = minAmount
		amountPrecisionString := (*(cells).At(MkInteger(2))).Call(MkString("replace"), MkString("<td>"), MkString(""))
		_ = amountPrecisionString
		amountPrecisionParts := amountPrecisionString.Split(MkString(" "))
		_ = amountPrecisionParts
		amountPrecision := this.SafeNumber(amountPrecisionParts, MkInteger(0))
		_ = amountPrecision
		idLength := OpSub(marketId.Length, MkInteger(0))
		_ = idLength
		quoteId := marketId.Slice(OpSub(idLength, MkInteger(3)), idLength)
		_ = quoteId
		quote := this.SafeCurrencyCode(quoteId)
		_ = quote
		pricePrecisionString := (*(cells).At(MkInteger(3))).Call(MkString("replace"), MkString("<td>"), MkString(""))
		_ = pricePrecisionString
		pricePrecisionParts := pricePrecisionString.Split(MkString(" "))
		_ = pricePrecisionParts
		pricePrecision := this.SafeNumber(pricePrecisionParts, MkInteger(0))
		_ = pricePrecision
		baseId := marketId.Replace(quoteId, MkString(""))
		_ = baseId
		base := this.SafeCurrencyCode(baseId)
		_ = base
		symbol := OpAdd(base, OpAdd(MkString("/"), quote))
		_ = symbol
		active := OpCopy(MkUndefined())
		_ = active
		result.Push(MkMap(&VarMap{
			"id":      marketId,
			"info":    row,
			"symbol":  symbol,
			"base":    base,
			"quote":   quote,
			"baseId":  baseId,
			"quoteId": quoteId,
			"active":  active,
			"precision": MkMap(&VarMap{
				"amount": amountPrecision,
				"price":  pricePrecision,
			}),
			"limits": MkMap(&VarMap{
				"amount": MkMap(&VarMap{
					"min": minAmount,
					"max": MkUndefined(),
				}),
				"price": MkMap(&VarMap{
					"min": MkUndefined(),
					"max": MkUndefined(),
				}),
				"cost": MkMap(&VarMap{
					"min": MkUndefined(),
					"max": MkUndefined(),
				}),
			}),
		}))
	}
	return result
}

func (this *Gemini) FetchMarketsFromAPI(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetV1Symbols"), params)
	_ = response
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
		marketId := *(response).At(i)
		_ = marketId
		market := OpCopy(marketId)
		_ = market
		idLength := OpSub(marketId.Length, MkInteger(0))
		_ = idLength
		baseId := marketId.Slice(MkInteger(0), OpSub(idLength, MkInteger(3)))
		_ = baseId
		quoteId := marketId.Slice(OpSub(idLength, MkInteger(3)), idLength)
		_ = quoteId
		base := this.SafeCurrencyCode(baseId)
		_ = base
		quote := this.SafeCurrencyCode(quoteId)
		_ = quote
		symbol := OpAdd(base, OpAdd(MkString("/"), quote))
		_ = symbol
		precision := MkMap(&VarMap{
			"amount": MkUndefined(),
			"price":  MkUndefined(),
		})
		_ = precision
		result.Push(MkMap(&VarMap{
			"id":        marketId,
			"info":      market,
			"symbol":    symbol,
			"base":      base,
			"quote":     quote,
			"baseId":    baseId,
			"quoteId":   quoteId,
			"precision": precision,
			"limits": MkMap(&VarMap{
				"amount": MkMap(&VarMap{
					"min": MkUndefined(),
					"max": MkUndefined(),
				}),
				"price": MkMap(&VarMap{
					"min": MkUndefined(),
					"max": MkUndefined(),
				}),
				"cost": MkMap(&VarMap{
					"min": MkUndefined(),
					"max": MkUndefined(),
				}),
			}),
			"active": MkUndefined(),
		}))
	}
	return result
}

func (this *Gemini) FetchOrderBook(goArgs ...*Variant) *Variant {
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
	response := this.Call(MkString("publicGetV1BookSymbol"), this.Extend(request, params))
	_ = response
	return this.ParseOrderBook(response, symbol, MkUndefined(), MkString("bids"), MkString("asks"), MkString("price"), MkString("amount"))
}

func (this *Gemini) FetchTickerV1(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("publicGetV1PubtickerSymbol"), this.Extend(request, params))
	_ = response
	return this.ParseTicker(response, market)
}

func (this *Gemini) FetchTickerV2(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("publicGetV2TickerSymbol"), this.Extend(request, params))
	_ = response
	return this.ParseTicker(response, market)
}

func (this *Gemini) FetchTickerV1AndV2(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	tickerA := this.FetchTickerV1(symbol, params)
	_ = tickerA
	tickerB := this.FetchTickerV2(symbol, params)
	_ = tickerB
	return this.DeepExtend(tickerA, MkMap(&VarMap{
		"open":       *(tickerB).At(MkString("open")),
		"high":       *(tickerB).At(MkString("high")),
		"low":        *(tickerB).At(MkString("low")),
		"change":     *(tickerB).At(MkString("change")),
		"percentage": *(tickerB).At(MkString("percentage")),
		"average":    *(tickerB).At(MkString("average")),
		"info":       *(tickerB).At(MkString("info")),
	}))
}

func (this *Gemini) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	method := this.SafeValue(*this.At(MkString("options")), MkString("fetchTickerMethod"), MkString("fetchTickerV1"))
	_ = method
	return this.Call(method, symbol, params)
}

func (this *Gemini) ParseTicker(goArgs ...*Variant) *Variant {
	ticker := GoGetArg(goArgs, 0, MkUndefined())
	_ = ticker
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	volume := this.SafeValue(ticker, MkString("volume"), MkMap(&VarMap{}))
	_ = volume
	timestamp := this.SafeInteger(volume, MkString("timestamp"))
	_ = timestamp
	symbol := OpCopy(MkUndefined())
	_ = symbol
	marketId := this.SafeString(ticker, MkString("pair"))
	_ = marketId
	baseId := OpCopy(MkUndefined())
	_ = baseId
	quoteId := OpCopy(MkUndefined())
	_ = quoteId
	base := OpCopy(MkUndefined())
	_ = base
	quote := OpCopy(MkUndefined())
	_ = quote
	if IsTrue(OpNotEq2(marketId, MkUndefined())) {
		if IsTrue(OpHasMember(marketId, *this.At(MkString("markets_by_id")))) {
			market = *(*this.At(MkString("markets_by_id"))).At(marketId)
		} else {
			idLength := OpSub(marketId.Length, MkInteger(0))
			_ = idLength
			if IsTrue(OpEq2(idLength, MkInteger(7))) {
				baseId = marketId.Slice(MkInteger(0), MkInteger(4))
				quoteId = marketId.Slice(MkInteger(4), MkInteger(7))
			} else {
				baseId = marketId.Slice(MkInteger(0), MkInteger(3))
				quoteId = marketId.Slice(MkInteger(3), MkInteger(6))
			}
			base = this.SafeCurrencyCode(baseId)
			quote = this.SafeCurrencyCode(quoteId)
			symbol = OpAdd(base, OpAdd(MkString("/"), quote))
		}
	}
	if IsTrue(OpAnd(OpEq2(symbol, MkUndefined()), OpNotEq2(market, MkUndefined()))) {
		symbol = *(market).At(MkString("symbol"))
		baseId = (*(market).At(MkString("baseId"))).Call(MkString("toUpperCase"))
		quoteId = (*(market).At(MkString("quoteId"))).Call(MkString("toUpperCase"))
		base = *(market).At(MkString("base"))
		quote = *(market).At(MkString("quote"))
	}
	price := this.SafeNumber(ticker, MkString("price"))
	_ = price
	last := this.SafeNumber2(ticker, MkString("last"), MkString("close"), price)
	_ = last
	percentage := this.SafeNumber(ticker, MkString("percentChange24h"))
	_ = percentage
	change := OpCopy(MkUndefined())
	_ = change
	open := this.SafeNumber(ticker, MkString("open"))
	_ = open
	average := OpCopy(MkUndefined())
	_ = average
	if IsTrue(OpNotEq2(last, MkUndefined())) {
		if IsTrue(OpNotEq2(open, MkUndefined())) {
			change = OpSub(last, open)
			if IsTrue(OpNotEq2(open, MkInteger(0))) {
				percentage = OpDiv(change, OpMulti(open, MkInteger(100)))
			}
			average = OpDiv(this.Sum(last, open), MkInteger(2))
		} else {
			if IsTrue(OpNotEq2(percentage, MkUndefined())) {
				change = OpMulti(last, percentage)
				if IsTrue(OpEq2(open, MkUndefined())) {
					open = OpSub(last, change)
				}
				average = OpDiv(this.Sum(last, open), MkInteger(2))
			}
		}
	}
	baseVolume := this.SafeNumber(volume, baseId)
	_ = baseVolume
	quoteVolume := this.SafeNumber(volume, quoteId)
	_ = quoteVolume
	vwap := this.Vwap(baseVolume, quoteVolume)
	_ = vwap
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

func (this *Gemini) FetchTickers(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("publicGetV1Pricefeed"), params)
	_ = response
	return this.ParseTickers(response, symbols)
}

func (this *Gemini) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.SafeInteger(trade, MkString("timestampms"))
	_ = timestamp
	id := this.SafeString(trade, MkString("tid"))
	_ = id
	orderId := this.SafeString(trade, MkString("order_id"))
	_ = orderId
	feeCurrencyId := this.SafeString(trade, MkString("fee_currency"))
	_ = feeCurrencyId
	feeCurrencyCode := this.SafeCurrencyCode(feeCurrencyId)
	_ = feeCurrencyCode
	fee := MkMap(&VarMap{
		"cost":     this.SafeNumber(trade, MkString("fee_amount")),
		"currency": feeCurrencyCode,
	})
	_ = fee
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
	type_ := OpCopy(MkUndefined())
	_ = type_
	side := this.SafeStringLower(trade, MkString("type"))
	_ = side
	symbol := this.SafeSymbol(MkUndefined(), market)
	_ = symbol
	return MkMap(&VarMap{
		"id":           id,
		"order":        orderId,
		"info":         trade,
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       symbol,
		"type":         type_,
		"side":         side,
		"takerOrMaker": MkUndefined(),
		"price":        price,
		"cost":         cost,
		"amount":       amount,
		"fee":          fee,
	})
}

func (this *Gemini) FetchTrades(goArgs ...*Variant) *Variant {
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
	response := this.Call(MkString("publicGetV1TradesSymbol"), this.Extend(request, params))
	_ = response
	return this.ParseTrades(response, market, since, limit)
}

func (this *Gemini) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privatePostV1Balances"), params)
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
		*(account).At(MkString("total")) = this.SafeString(balance, MkString("amount"))
		*(result).At(code) = OpCopy(account)
	}
	return this.ParseBalance(result)
}

func (this *Gemini) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.SafeInteger(order, MkString("timestampms"))
	_ = timestamp
	amount := this.SafeNumber(order, MkString("original_amount"))
	_ = amount
	remaining := this.SafeNumber(order, MkString("remaining_amount"))
	_ = remaining
	filled := this.SafeNumber(order, MkString("executed_amount"))
	_ = filled
	status := MkString("closed")
	_ = status
	if IsTrue(*(order).At(MkString("is_live"))) {
		status = MkString("open")
	}
	if IsTrue(*(order).At(MkString("is_cancelled"))) {
		status = MkString("canceled")
	}
	price := this.SafeNumber(order, MkString("price"))
	_ = price
	average := this.SafeNumber(order, MkString("avg_execution_price"))
	_ = average
	type_ := this.SafeString(order, MkString("type"))
	_ = type_
	if IsTrue(OpEq2(type_, MkString("exchange limit"))) {
		type_ = MkString("limit")
	} else {
		if IsTrue(OpOr(OpEq2(type_, MkString("market buy")), OpEq2(type_, MkString("market sell")))) {
			type_ = MkString("market")
		} else {
			type_ = *(order).At(MkString("type"))
		}
	}
	fee := OpCopy(MkUndefined())
	_ = fee
	marketId := this.SafeString(order, MkString("symbol"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market)
	_ = symbol
	id := this.SafeString(order, MkString("order_id"))
	_ = id
	side := this.SafeStringLower(order, MkString("side"))
	_ = side
	clientOrderId := this.SafeString(order, MkString("client_order_id"))
	_ = clientOrderId
	return this.SafeOrder(MkMap(&VarMap{
		"id":                 id,
		"clientOrderId":      clientOrderId,
		"info":               order,
		"timestamp":          timestamp,
		"datetime":           this.Iso8601(timestamp),
		"lastTradeTimestamp": MkUndefined(),
		"status":             status,
		"symbol":             symbol,
		"type":               type_,
		"timeInForce":        MkUndefined(),
		"postOnly":           MkUndefined(),
		"side":               side,
		"price":              price,
		"stopPrice":          MkUndefined(),
		"average":            average,
		"cost":               MkUndefined(),
		"amount":             amount,
		"filled":             filled,
		"remaining":          remaining,
		"fee":                fee,
		"trades":             MkUndefined(),
	}))
}

func (this *Gemini) FetchOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"order_id": id})
	_ = request
	response := this.Call(MkString("privatePostV1OrderStatus"), this.Extend(request, params))
	_ = response
	return this.ParseOrder(response)
}

func (this *Gemini) FetchOpenOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privatePostV1Orders"), params)
	_ = response
	market := OpCopy(MkUndefined())
	_ = market
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
	}
	return this.ParseOrders(response, market, since, limit)
}

func (this *Gemini) CreateOrder(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpEq2(type_, MkString("market"))) {
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), MkString(" allows limit orders only"))))
	}
	nonce := this.Nonce()
	_ = nonce
	request := MkMap(&VarMap{
		"client_order_id": nonce.ToString(),
		"symbol":          this.MarketId(symbol),
		"amount":          amount.ToString(),
		"price":           price.ToString(),
		"side":            side,
		"type":            MkString("exchange limit"),
	})
	_ = request
	response := this.Call(MkString("privatePostV1OrderNew"), this.Extend(request, params))
	_ = response
	return MkMap(&VarMap{
		"info": response,
		"id":   *(response).At(MkString("order_id")),
	})
}

func (this *Gemini) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"order_id": id})
	_ = request
	return this.Call(MkString("privatePostV1OrderCancel"), this.Extend(request, params))
}

func (this *Gemini) FetchMyTrades(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit_trades")) = OpCopy(limit)
	}
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("timestamp")) = parseInt(OpDiv(since, MkInteger(1000)))
	}
	response := this.Call(MkString("privatePostV1Mytrades"), this.Extend(request, params))
	_ = response
	return this.ParseTrades(response, market, since, limit)
}

func (this *Gemini) Withdraw(goArgs ...*Variant) *Variant {
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
		"address":  address,
	})
	_ = request
	response := this.Call(MkString("privatePostV1WithdrawCurrency"), this.Extend(request, params))
	_ = response
	return MkMap(&VarMap{
		"info": response,
		"id":   this.SafeString(response, MkString("txHash")),
	})
}

func (this *Gemini) Nonce(goArgs ...*Variant) *Variant {
	return this.Milliseconds()
}

func (this *Gemini) FetchTransactions(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit_transfers")) = OpCopy(limit)
	}
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("timestamp")) = OpCopy(since)
	}
	response := this.Call(MkString("privatePostV1Transfers"), this.Extend(request, params))
	_ = response
	return this.ParseTransactions(response)
}

func (this *Gemini) ParseTransaction(goArgs ...*Variant) *Variant {
	transaction := GoGetArg(goArgs, 0, MkUndefined())
	_ = transaction
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	timestamp := this.SafeInteger(transaction, MkString("timestampms"))
	_ = timestamp
	currencyId := this.SafeString(transaction, MkString("currency"))
	_ = currencyId
	code := this.SafeCurrencyCode(currencyId, currency)
	_ = code
	address := this.SafeString(transaction, MkString("destination"))
	_ = address
	type_ := this.SafeStringLower(transaction, MkString("type"))
	_ = type_
	status := MkString("pending")
	_ = status
	if IsTrue(*(transaction).At(MkString("status"))) {
		status = MkString("ok")
	}
	fee := OpCopy(MkUndefined())
	_ = fee
	feeAmount := this.SafeNumber(transaction, MkString("feeAmount"))
	_ = feeAmount
	if IsTrue(OpNotEq2(feeAmount, MkUndefined())) {
		fee = MkMap(&VarMap{
			"cost":     feeAmount,
			"currency": code,
		})
	}
	return MkMap(&VarMap{
		"info":      transaction,
		"id":        this.SafeString(transaction, MkString("eid")),
		"txid":      this.SafeString(transaction, MkString("txHash")),
		"timestamp": timestamp,
		"datetime":  this.Iso8601(timestamp),
		"address":   address,
		"tag":       MkUndefined(),
		"type":      type_,
		"amount":    this.SafeNumber(transaction, MkString("amount")),
		"currency":  code,
		"status":    status,
		"updated":   MkUndefined(),
		"fee":       fee,
	})
}

func (this *Gemini) Sign(goArgs ...*Variant) *Variant {
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
	url := OpAdd(MkString("/"), this.ImplodeParams(path, params))
	_ = url
	query := this.Omit(params, this.ExtractParams(path))
	_ = query
	if IsTrue(OpEq2(api, MkString("private"))) {
		this.CheckRequiredCredentials()
		nonce := this.Nonce()
		_ = nonce
		request := this.Extend(MkMap(&VarMap{
			"request": url,
			"nonce":   nonce,
		}), query)
		_ = request
		payload := this.Json(request)
		_ = payload
		payload = this.StringToBase64(payload)
		signature := this.Hmac(payload, this.Encode(*this.At(MkString("secret"))), MkString("sha384"))
		_ = signature
		headers = MkMap(&VarMap{
			"Content-Type":       MkString("text/plain"),
			"X-GEMINI-APIKEY":    *this.At(MkString("apiKey")),
			"X-GEMINI-PAYLOAD":   this.Decode(payload),
			"X-GEMINI-SIGNATURE": signature,
		})
	} else {
		if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
			OpAddAssign(&url, OpAdd(MkString("?"), this.Urlencode(query)))
		}
	}
	url = OpAdd(*(*(*this.At(MkString("urls"))).At(MkString("api"))).At(api), url)
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Gemini) HandleErrors(goArgs ...*Variant) *Variant {
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
		if IsTrue(OpEq2(OpType(body), MkString("string"))) {
			feedback := OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))
			_ = feedback
			this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("broad")), body, feedback)
		}
		return MkUndefined()
	}
	result := this.SafeString(response, MkString("result"))
	_ = result
	if IsTrue(OpEq2(result, MkString("error"))) {
		reason := this.SafeString(response, MkString("reason"))
		_ = reason
		message := this.SafeString(response, MkString("message"))
		_ = message
		feedback := OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), message))
		_ = feedback
		this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("exact")), reason, feedback)
		this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("exact")), message, feedback)
		this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("broad")), message, feedback)
		panic(NewExchangeError(feedback))
	}
	return MkUndefined()
}

func (this *Gemini) CreateDepositAddress(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{"currency": *(currency).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("privatePostV1DepositCurrencyNewAddress"), this.Extend(request, params))
	_ = response
	address := this.SafeString(response, MkString("address"))
	_ = address
	this.CheckAddress(address)
	return MkMap(&VarMap{
		"currency": code,
		"address":  address,
		"tag":      MkUndefined(),
		"info":     response,
	})
}

func (this *Gemini) FetchOHLCV(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	timeframe := GoGetArg(goArgs, 1, MkString("5m"))
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
		"timeframe": *(*this.At(MkString("timeframes"))).At(timeframe),
		"symbol":    *(market).At(MkString("id")),
	})
	_ = request
	response := this.Call(MkString("publicGetV2CandlesSymbolTimeframe"), this.Extend(request, params))
	_ = response
	return this.ParseOHLCVs(response, market, timeframe, since, limit)
}
