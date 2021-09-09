package ccxt

import ()

type Btctradeua struct {
	*ExchangeBase
}

var _ Exchange = (*Btctradeua)(nil)

func init() {
	exchange := &Btctradeua{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Btctradeua) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("btctradeua"),
		"name": MkString("BTC Trade UA"),
		"countries": MkArray(&VarArray{
			MkString("UA"),
		}),
		"rateLimit": MkInteger(3000),
		"has": MkMap(&VarMap{
			"cancelOrder":       MkBool(true),
			"CORS":              MkBool(false),
			"createMarketOrder": MkBool(false),
			"createOrder":       MkBool(true),
			"fetchBalance":      MkBool(true),
			"fetchOpenOrders":   MkBool(true),
			"fetchOrderBook":    MkBool(true),
			"fetchTicker":       MkBool(true),
			"fetchTrades":       MkBool(true),
			"signIn":            MkBool(true),
		}),
		"urls": MkMap(&VarMap{
			"referral": MkString("https://btc-trade.com.ua/registration/22689"),
			"logo":     MkString("https://user-images.githubusercontent.com/1294454/27941483-79fc7350-62d9-11e7-9f61-ac47f28fcd96.jpg"),
			"api":      MkString("https://btc-trade.com.ua/api"),
			"www":      MkString("https://btc-trade.com.ua"),
			"doc":      MkString("https://docs.google.com/document/d/1ocYA0yMy_RXd561sfG3qEPZ80kyll36HUxvCRe5GbhE/edit"),
		}),
		"api": MkMap(&VarMap{
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("deals/{symbol}"),
				MkString("trades/sell/{symbol}"),
				MkString("trades/buy/{symbol}"),
				MkString("japan_stat/high/{symbol}"),
			})}),
			"private": MkMap(&VarMap{"post": MkArray(&VarArray{
				MkString("auth"),
				MkString("ask/{symbol}"),
				MkString("balance"),
				MkString("bid/{symbol}"),
				MkString("buy/{symbol}"),
				MkString("my_orders/{symbol}"),
				MkString("order/status/{id}"),
				MkString("remove/order/{id}"),
				MkString("sell/{symbol}"),
			})}),
		}),
		"markets": MkMap(&VarMap{
			"BCH/UAH": MkMap(&VarMap{
				"id":      MkString("bch_uah"),
				"symbol":  MkString("BCH/UAH"),
				"base":    MkString("BCH"),
				"quote":   MkString("UAH"),
				"baseId":  MkString("bch"),
				"quoteId": MkString("uah"),
			}),
			"BTC/UAH": MkMap(&VarMap{
				"id":        MkString("btc_uah"),
				"symbol":    MkString("BTC/UAH"),
				"base":      MkString("BTC"),
				"quote":     MkString("UAH"),
				"baseId":    MkString("btc"),
				"quoteId":   MkString("uah"),
				"precision": MkMap(&VarMap{"price": MkInteger(1)}),
				"limits":    MkMap(&VarMap{"amount": MkMap(&VarMap{"min": MkNumber(0.0000000001)})}),
			}),
			"DASH/BTC": MkMap(&VarMap{
				"id":      MkString("dash_btc"),
				"symbol":  MkString("DASH/BTC"),
				"base":    MkString("DASH"),
				"quote":   MkString("BTC"),
				"baseId":  MkString("dash"),
				"quoteId": MkString("btc"),
			}),
			"DASH/UAH": MkMap(&VarMap{
				"id":      MkString("dash_uah"),
				"symbol":  MkString("DASH/UAH"),
				"base":    MkString("DASH"),
				"quote":   MkString("UAH"),
				"baseId":  MkString("dash"),
				"quoteId": MkString("uah"),
			}),
			"DOGE/BTC": MkMap(&VarMap{
				"id":      MkString("doge_btc"),
				"symbol":  MkString("DOGE/BTC"),
				"base":    MkString("DOGE"),
				"quote":   MkString("BTC"),
				"baseId":  MkString("doge"),
				"quoteId": MkString("btc"),
			}),
			"DOGE/UAH": MkMap(&VarMap{
				"id":      MkString("doge_uah"),
				"symbol":  MkString("DOGE/UAH"),
				"base":    MkString("DOGE"),
				"quote":   MkString("UAH"),
				"baseId":  MkString("doge"),
				"quoteId": MkString("uah"),
			}),
			"ETH/UAH": MkMap(&VarMap{
				"id":      MkString("eth_uah"),
				"symbol":  MkString("ETH/UAH"),
				"base":    MkString("ETH"),
				"quote":   MkString("UAH"),
				"baseId":  MkString("eth"),
				"quoteId": MkString("uah"),
			}),
			"ITI/UAH": MkMap(&VarMap{
				"id":      MkString("iti_uah"),
				"symbol":  MkString("ITI/UAH"),
				"base":    MkString("ITI"),
				"quote":   MkString("UAH"),
				"baseId":  MkString("iti"),
				"quoteId": MkString("uah"),
			}),
			"KRB/UAH": MkMap(&VarMap{
				"id":      MkString("krb_uah"),
				"symbol":  MkString("KRB/UAH"),
				"base":    MkString("KRB"),
				"quote":   MkString("UAH"),
				"baseId":  MkString("krb"),
				"quoteId": MkString("uah"),
			}),
			"LTC/BTC": MkMap(&VarMap{
				"id":      MkString("ltc_btc"),
				"symbol":  MkString("LTC/BTC"),
				"base":    MkString("LTC"),
				"quote":   MkString("BTC"),
				"baseId":  MkString("ltc"),
				"quoteId": MkString("btc"),
			}),
			"LTC/UAH": MkMap(&VarMap{
				"id":      MkString("ltc_uah"),
				"symbol":  MkString("LTC/UAH"),
				"base":    MkString("LTC"),
				"quote":   MkString("UAH"),
				"baseId":  MkString("ltc"),
				"quoteId": MkString("uah"),
			}),
			"NVC/BTC": MkMap(&VarMap{
				"id":      MkString("nvc_btc"),
				"symbol":  MkString("NVC/BTC"),
				"base":    MkString("NVC"),
				"quote":   MkString("BTC"),
				"baseId":  MkString("nvc"),
				"quoteId": MkString("btc"),
			}),
			"NVC/UAH": MkMap(&VarMap{
				"id":      MkString("nvc_uah"),
				"symbol":  MkString("NVC/UAH"),
				"base":    MkString("NVC"),
				"quote":   MkString("UAH"),
				"baseId":  MkString("nvc"),
				"quoteId": MkString("uah"),
			}),
			"PPC/BTC": MkMap(&VarMap{
				"id":      MkString("ppc_btc"),
				"symbol":  MkString("PPC/BTC"),
				"base":    MkString("PPC"),
				"quote":   MkString("BTC"),
				"baseId":  MkString("ppc"),
				"quoteId": MkString("btc"),
			}),
			"SIB/UAH": MkMap(&VarMap{
				"id":      MkString("sib_uah"),
				"symbol":  MkString("SIB/UAH"),
				"base":    MkString("SIB"),
				"quote":   MkString("UAH"),
				"baseId":  MkString("sib"),
				"quoteId": MkString("uah"),
			}),
			"XMR/UAH": MkMap(&VarMap{
				"id":      MkString("xmr_uah"),
				"symbol":  MkString("XMR/UAH"),
				"base":    MkString("XMR"),
				"quote":   MkString("UAH"),
				"baseId":  MkString("xmr"),
				"quoteId": MkString("uah"),
			}),
			"ZEC/UAH": MkMap(&VarMap{
				"id":      MkString("zec_uah"),
				"symbol":  MkString("ZEC/UAH"),
				"base":    MkString("ZEC"),
				"quote":   MkString("UAH"),
				"baseId":  MkString("zec"),
				"quoteId": MkString("uah"),
			}),
		}),
		"fees": MkMap(&VarMap{
			"trading": MkMap(&VarMap{
				"maker": OpDiv(MkNumber(0.1), MkInteger(100)),
				"taker": OpDiv(MkNumber(0.1), MkInteger(100)),
			}),
			"funding": MkMap(&VarMap{"withdraw": MkMap(&VarMap{
				"BTC":  MkNumber(0.0006),
				"LTC":  MkNumber(0.01),
				"NVC":  MkNumber(0.01),
				"DOGE": MkInteger(10),
			})}),
		}),
	}))
}

func (this *Btctradeua) SignIn(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	return this.Call(MkString("privatePostAuth"), params)
}

func (this *Btctradeua) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privatePostBalance"), params)
	_ = response
	result := MkMap(&VarMap{"info": response})
	_ = result
	balances := this.SafeValue(response, MkString("accounts"))
	_ = balances
	for i := MkInteger(0); IsTrue(OpLw(i, balances.Length)); OpIncr(&i) {
		balance := *(balances).At(i)
		_ = balance
		currencyId := this.SafeString(balance, MkString("currency"))
		_ = currencyId
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		account := this.Account()
		_ = account
		*(account).At(MkString("total")) = this.SafeString(balance, MkString("balance"))
		*(result).At(code) = OpCopy(account)
	}
	return this.ParseBalance(result)
}

func (this *Btctradeua) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	bids := this.Call(MkString("publicGetTradesBuySymbol"), this.Extend(request, params))
	_ = bids
	asks := this.Call(MkString("publicGetTradesSellSymbol"), this.Extend(request, params))
	_ = asks
	orderbook := MkMap(&VarMap{
		"bids": MkArray(&VarArray{}),
		"asks": MkArray(&VarArray{}),
	})
	_ = orderbook
	if IsTrue(bids) {
		if IsTrue(OpHasMember(MkString("list"), bids)) {
			*(orderbook).At(MkString("bids")) = *(bids).At(MkString("list"))
		}
	}
	if IsTrue(asks) {
		if IsTrue(OpHasMember(MkString("list"), asks)) {
			*(orderbook).At(MkString("asks")) = *(asks).At(MkString("list"))
		}
	}
	return this.ParseOrderBook(orderbook, symbol, MkUndefined(), MkString("bids"), MkString("asks"), MkString("price"), MkString("currency_trade"))
}

func (this *Btctradeua) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"symbol": this.MarketId(symbol)})
	_ = request
	response := this.Call(MkString("publicGetJapanStatHighSymbol"), this.Extend(request, params))
	_ = response
	ticker := this.SafeValue(response, MkString("trades"))
	_ = ticker
	timestamp := this.Milliseconds()
	_ = timestamp
	result := MkMap(&VarMap{
		"symbol":        symbol,
		"timestamp":     timestamp,
		"datetime":      this.Iso8601(timestamp),
		"high":          MkUndefined(),
		"low":           MkUndefined(),
		"bid":           MkUndefined(),
		"bidVolume":     MkUndefined(),
		"ask":           MkUndefined(),
		"askVolume":     MkUndefined(),
		"vwap":          MkUndefined(),
		"open":          MkUndefined(),
		"close":         MkUndefined(),
		"last":          MkUndefined(),
		"previousClose": MkUndefined(),
		"change":        MkUndefined(),
		"percentage":    MkUndefined(),
		"average":       MkUndefined(),
		"baseVolume":    MkUndefined(),
		"quoteVolume":   MkUndefined(),
		"info":          ticker,
	})
	_ = result
	tickerLength := OpCopy(ticker.Length)
	_ = tickerLength
	if IsTrue(OpGt(tickerLength, MkInteger(0))) {
		start := Math.Max(OpSub(tickerLength, MkInteger(48)), MkInteger(0))
		_ = start
		for i := OpCopy(start); IsTrue(OpLw(i, ticker.Length)); OpIncr(&i) {
			candle := *(ticker).At(i)
			_ = candle
			if IsTrue(OpEq2(*(result).At(MkString("open")), MkUndefined())) {
				*(result).At(MkString("open")) = this.SafeNumber(candle, MkInteger(1))
			}
			high := this.SafeNumber(candle, MkInteger(2))
			_ = high
			if IsTrue(OpOr(OpEq2(*(result).At(MkString("high")), MkUndefined()), OpAnd(OpNotEq2(high, MkUndefined()), OpLw(*(result).At(MkString("high")), high)))) {
				*(result).At(MkString("high")) = OpCopy(high)
			}
			low := this.SafeNumber(candle, MkInteger(3))
			_ = low
			if IsTrue(OpOr(OpEq2(*(result).At(MkString("low")), MkUndefined()), OpAnd(OpNotEq2(low, MkUndefined()), OpGt(*(result).At(MkString("low")), low)))) {
				*(result).At(MkString("low")) = OpCopy(low)
			}
			baseVolume := this.SafeNumber(candle, MkInteger(5))
			_ = baseVolume
			if IsTrue(OpEq2(*(result).At(MkString("baseVolume")), MkUndefined())) {
				*(result).At(MkString("baseVolume")) = OpCopy(baseVolume)
			} else {
				*(result).At(MkString("baseVolume")) = this.Sum(*(result).At(MkString("baseVolume")), baseVolume)
			}
		}
		last := OpSub(tickerLength, MkInteger(1))
		_ = last
		*(result).At(MkString("last")) = this.SafeNumber(*(ticker).At(last), MkInteger(4))
		*(result).At(MkString("close")) = *(result).At(MkString("last"))
	}
	return result
}

func (this *Btctradeua) ConvertMonthNameToString(goArgs ...*Variant) *Variant {
	cyrillic := GoGetArg(goArgs, 0, MkUndefined())
	_ = cyrillic
	months := MkMap(&VarMap{
		"Jan": MkString("01"),
		"Feb": MkString("02"),
		"Mar": MkString("03"),
		"Apr": MkString("04"),
		"May": MkString("05"),
		"Jun": MkString("06"),
		"Jul": MkString("07"),
		"Aug": MkString("08"),
		"Sep": MkString("09"),
		"Oct": MkString("10"),
		"Nov": MkString("11"),
		"Dec": MkString("12"),
	})
	_ = months
	return this.SafeString(months, cyrillic)
}

func (this *Btctradeua) ParseExchangeSpecificDatetime(goArgs ...*Variant) *Variant {
	cyrillic := GoGetArg(goArgs, 0, MkUndefined())
	_ = cyrillic
	parts := cyrillic.Split(MkString(" "))
	_ = parts
	month := *(parts).At(MkInteger(0))
	_ = month
	day := (*(parts).At(MkInteger(1))).Call(MkString("replace"), MkString(","), MkString(""))
	_ = day
	if IsTrue(OpLw(day.Length, MkInteger(2))) {
		day = OpAdd(MkString("0"), day)
	}
	year := (*(parts).At(MkInteger(2))).Call(MkString("replace"), MkString(","), MkString(""))
	_ = year
	month = month.Replace(MkString(","), MkString(""))
	month = month.Replace(MkString("."), MkString(""))
	month = this.ConvertMonthNameToString(month)
	if IsTrue(OpNot(month)) {
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" parseTrade() undefined month name: "), cyrillic))))
	}
	hms := *(parts).At(MkInteger(3))
	_ = hms
	hmsParts := hms.Split(MkString(":"))
	_ = hmsParts
	h := this.SafeString(hmsParts, MkInteger(0))
	_ = h
	m := MkString("00")
	_ = m
	ampm := this.SafeString(parts, MkInteger(4))
	_ = ampm
	if IsTrue(OpEq2(h, MkString("noon"))) {
		h = MkString("12")
	} else {
		intH := parseInt(h)
		_ = intH
		if IsTrue(OpAnd(OpNotEq2(ampm, MkUndefined()), OpEq2(*(ampm).At(MkInteger(0)), MkString("p")))) {
			intH = OpAdd(MkInteger(12), intH)
			if IsTrue(OpGt(intH, MkInteger(23))) {
				intH = MkInteger(0)
			}
		}
		h = intH.ToString()
		if IsTrue(OpLw(h.Length, MkInteger(2))) {
			h = OpAdd(MkString("0"), h)
		}
		m = this.SafeString(hmsParts, MkInteger(1), MkString("00"))
		if IsTrue(OpLw(m.Length, MkInteger(2))) {
			m = OpAdd(MkString("0"), m)
		}
	}
	ymd := (MkArray(&VarArray{
		year,
		month,
		day,
	})).Call(MkString("join"), MkString("-"))
	_ = ymd
	ymdhms := OpAdd(ymd, OpAdd(MkString("T"), OpAdd(h, OpAdd(MkString(":"), OpAdd(m, MkString(":00"))))))
	_ = ymdhms
	timestamp := this.Parse8601(ymdhms)
	_ = timestamp
	intM := parseInt(m)
	_ = intM
	if IsTrue(OpOr(OpLw(intM, MkInteger(11)), OpGt(intM, MkInteger(2)))) {
		return OpSub(timestamp, MkInteger(7200000))
	}
	return OpSub(timestamp, MkInteger(10800000))
}

func (this *Btctradeua) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.ParseExchangeSpecificDatetime(this.SafeString(trade, MkString("pub_date")))
	_ = timestamp
	id := this.SafeString(trade, MkString("id"))
	_ = id
	type_ := MkString("limit")
	_ = type_
	side := this.SafeString(trade, MkString("type"))
	_ = side
	priceString := this.SafeString(trade, MkString("price"))
	_ = priceString
	amountString := this.SafeString(trade, MkString("amnt_trade"))
	_ = amountString
	price := this.ParseNumber(priceString)
	_ = price
	amount := this.ParseNumber(amountString)
	_ = amount
	cost := this.ParseNumber(Precise.StringMul(priceString, amountString))
	_ = cost
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(OpNotEq2(market, MkUndefined())) {
		symbol = *(market).At(MkString("symbol"))
	}
	return MkMap(&VarMap{
		"id":           id,
		"info":         trade,
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       symbol,
		"type":         type_,
		"side":         side,
		"order":        MkUndefined(),
		"takerOrMaker": MkUndefined(),
		"price":        price,
		"amount":       amount,
		"cost":         cost,
		"fee":          MkUndefined(),
	})
}

func (this *Btctradeua) FetchTrades(goArgs ...*Variant) *Variant {
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
	response := this.Call(MkString("publicGetDealsSymbol"), this.Extend(request, params))
	_ = response
	trades := MkArray(&VarArray{})
	_ = trades
	for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
		id := this.SafeInteger(*(response).At(i), MkString("id"))
		_ = id
		if IsTrue(OpMod(id, MkInteger(2))) {
			trades.Push(*(response).At(i))
		}
	}
	return this.ParseTrades(trades, market, since, limit)
}

func (this *Btctradeua) CreateOrder(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpEq2(type_, MkString("market"))) {
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), MkString(" allows limit orders only"))))
	}
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	method := OpAdd(MkString("privatePost"), OpAdd(this.Capitalize(side), MkString("Id")))
	_ = method
	request := MkMap(&VarMap{
		"count":     amount,
		"currency1": *(market).At(MkString("quoteId")),
		"currency":  *(market).At(MkString("baseId")),
		"price":     price,
	})
	_ = request
	return this.Call(method, this.Extend(request, params))
}

func (this *Btctradeua) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	request := MkMap(&VarMap{"id": id})
	_ = request
	return this.Call(MkString("privatePostRemoveOrderId"), this.Extend(request, params))
}

func (this *Btctradeua) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.Milliseconds()
	_ = timestamp
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(OpNotEq2(market, MkUndefined())) {
		symbol = *(market).At(MkString("symbol"))
	}
	return MkMap(&VarMap{
		"id":                 this.SafeString(order, MkString("id")),
		"clientOrderId":      MkUndefined(),
		"timestamp":          timestamp,
		"datetime":           this.Iso8601(timestamp),
		"lastTradeTimestamp": MkUndefined(),
		"status":             MkString("open"),
		"symbol":             symbol,
		"type":               MkUndefined(),
		"timeInForce":        MkUndefined(),
		"postOnly":           MkUndefined(),
		"side":               this.SafeString(order, MkString("type")),
		"price":              this.SafeNumber(order, MkString("price")),
		"stopPrice":          MkUndefined(),
		"amount":             this.SafeNumber(order, MkString("amnt_trade")),
		"filled":             MkInteger(0),
		"remaining":          this.SafeNumber(order, MkString("amnt_trade")),
		"trades":             MkUndefined(),
		"info":               order,
		"cost":               MkUndefined(),
		"average":            MkUndefined(),
		"fee":                MkUndefined(),
	})
}

func (this *Btctradeua) FetchOpenOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchOpenOrders() requires a symbol argument"))))
	}
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("privatePostMyOrdersSymbol"), this.Extend(request, params))
	_ = response
	orders := this.SafeValue(response, MkString("your_open_orders"))
	_ = orders
	return this.ParseOrders(orders, market, since, limit)
}

func (this *Btctradeua) Nonce(goArgs ...*Variant) *Variant {
	return this.Milliseconds()
}

func (this *Btctradeua) Sign(goArgs ...*Variant) *Variant {
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
	url := OpAdd(*(*this.At(MkString("urls"))).At(MkString("api")), OpAdd(MkString("/"), this.ImplodeParams(path, params)))
	_ = url
	query := this.Omit(params, this.ExtractParams(path))
	_ = query
	if IsTrue(OpEq2(api, MkString("public"))) {
		if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
			OpAddAssign(&url, this.ImplodeParams(path, query))
		}
	} else {
		this.CheckRequiredCredentials()
		nonce := this.Nonce()
		_ = nonce
		body = this.Urlencode(this.Extend(MkMap(&VarMap{
			"out_order_id": nonce,
			"nonce":        nonce,
		}), query))
		auth := OpAdd(body, *this.At(MkString("secret")))
		_ = auth
		headers = MkMap(&VarMap{
			"public-key":   *this.At(MkString("apiKey")),
			"api-sign":     this.Hash(this.Encode(auth), MkString("sha256")),
			"Content-Type": MkString("application/x-www-form-urlencoded"),
		})
	}
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}
