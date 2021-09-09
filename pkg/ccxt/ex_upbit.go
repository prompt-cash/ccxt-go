package ccxt

import ()

type Upbit struct {
	*ExchangeBase
}

var _ Exchange = (*Upbit)(nil)

func init() {
	exchange := &Upbit{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Upbit) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("upbit"),
		"name": MkString("Upbit"),
		"countries": MkArray(&VarArray{
			MkString("KR"),
		}),
		"version":   MkString("v1"),
		"rateLimit": MkInteger(1000),
		"pro":       MkBool(true),
		"has": MkMap(&VarMap{
			"cancelOrder":          MkBool(true),
			"CORS":                 MkBool(true),
			"createDepositAddress": MkBool(true),
			"createMarketOrder":    MkBool(true),
			"createOrder":          MkBool(true),
			"fetchBalance":         MkBool(true),
			"fetchClosedOrders":    MkBool(true),
			"fetchDepositAddress":  MkBool(true),
			"fetchDeposits":        MkBool(true),
			"fetchMarkets":         MkBool(true),
			"fetchMyTrades":        MkBool(false),
			"fetchOHLCV":           MkBool(true),
			"fetchOpenOrders":      MkBool(true),
			"fetchOrder":           MkBool(true),
			"fetchOrderBook":       MkBool(true),
			"fetchOrderBooks":      MkBool(true),
			"fetchOrders":          MkBool(false),
			"fetchTicker":          MkBool(true),
			"fetchTickers":         MkBool(true),
			"fetchTrades":          MkBool(true),
			"fetchTransactions":    MkBool(false),
			"fetchWithdrawals":     MkBool(true),
			"withdraw":             MkBool(true),
		}),
		"timeframes": MkMap(&VarMap{
			"1m":  MkString("minutes"),
			"3m":  MkString("minutes"),
			"5m":  MkString("minutes"),
			"15m": MkString("minutes"),
			"30m": MkString("minutes"),
			"1h":  MkString("minutes"),
			"4h":  MkString("minutes"),
			"1d":  MkString("days"),
			"1w":  MkString("weeks"),
			"1M":  MkString("months"),
		}),
		"hostname": MkString("api.upbit.com"),
		"urls": MkMap(&VarMap{
			"logo": MkString("https://user-images.githubusercontent.com/1294454/49245610-eeaabe00-f423-11e8-9cba-4b0aed794799.jpg"),
			"api": MkMap(&VarMap{
				"public":  MkString("https://{hostname}"),
				"private": MkString("https://{hostname}"),
			}),
			"www":  MkString("https://upbit.com"),
			"doc":  MkString("https://docs.upbit.com/docs/%EC%9A%94%EC%B2%AD-%EC%88%98-%EC%A0%9C%ED%95%9C"),
			"fees": MkString("https://upbit.com/service_center/guide"),
		}),
		"api": MkMap(&VarMap{
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("market/all"),
				MkString("candles/{timeframe}"),
				MkString("candles/{timeframe}/{unit}"),
				MkString("candles/minutes/{unit}"),
				MkString("candles/minutes/1"),
				MkString("candles/minutes/3"),
				MkString("candles/minutes/5"),
				MkString("candles/minutes/15"),
				MkString("candles/minutes/30"),
				MkString("candles/minutes/60"),
				MkString("candles/minutes/240"),
				MkString("candles/days"),
				MkString("candles/weeks"),
				MkString("candles/months"),
				MkString("trades/ticks"),
				MkString("ticker"),
				MkString("orderbook"),
			})}),
			"private": MkMap(&VarMap{
				"get": MkArray(&VarArray{
					MkString("accounts"),
					MkString("orders/chance"),
					MkString("order"),
					MkString("orders"),
					MkString("withdraws"),
					MkString("withdraw"),
					MkString("withdraws/chance"),
					MkString("deposits"),
					MkString("deposit"),
					MkString("deposits/coin_addresses"),
					MkString("deposits/coin_address"),
				}),
				"post": MkArray(&VarArray{
					MkString("orders"),
					MkString("withdraws/coin"),
					MkString("withdraws/krw"),
					MkString("deposits/generate_coin_address"),
				}),
				"delete": MkArray(&VarArray{
					MkString("order"),
				}),
			}),
		}),
		"fees": MkMap(&VarMap{
			"trading": MkMap(&VarMap{
				"tierBased":  MkBool(false),
				"percentage": MkBool(true),
				"maker":      this.ParseNumber(MkString("0.0025")),
				"taker":      this.ParseNumber(MkString("0.0025")),
			}),
			"funding": MkMap(&VarMap{
				"tierBased":  MkBool(false),
				"percentage": MkBool(false),
				"withdraw":   MkMap(&VarMap{}),
				"deposit":    MkMap(&VarMap{}),
			}),
		}),
		"exceptions": MkMap(&VarMap{
			"exact": MkMap(&VarMap{
				"This key has expired.": AuthenticationError,
				"Missing request parameter error. Check the required parameters!": BadRequest,
				"side is missing, side does not have a valid value":               InvalidOrder,
			}),
			"broad": MkMap(&VarMap{
				"thirdparty_agreement_required": PermissionDenied,
				"out_of_scope":                  PermissionDenied,
				"order_not_found":               OrderNotFound,
				"insufficient_funds":            InsufficientFunds,
				"invalid_access_key":            AuthenticationError,
				"jwt_verification":              AuthenticationError,
				"create_ask_error":              ExchangeError,
				"create_bid_error":              ExchangeError,
				"volume_too_large":              InvalidOrder,
				"invalid_funds":                 InvalidOrder,
			}),
		}),
		"options": MkMap(&VarMap{
			"createMarketBuyOrderRequiresPrice": MkBool(true),
			"fetchTickersMaxLength":             MkInteger(4096),
			"fetchOrderBooksMaxLength":          MkInteger(4096),
			"tradingFeesByQuoteCurrency":        MkMap(&VarMap{"KRW": MkNumber(0.0005)}),
		}),
		"commonCurrencies": MkMap(&VarMap{"TON": MkString("Tokamak Network")}),
	}))
}

func (this *Upbit) FetchCurrency(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	return this.FetchCurrencyById(*(currency).At(MkString("id")), params)
}

func (this *Upbit) FetchCurrencyById(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	request := MkMap(&VarMap{"currency": id})
	_ = request
	response := this.Call(MkString("privateGetWithdrawsChance"), this.Extend(request, params))
	_ = response
	memberInfo := this.SafeValue(response, MkString("member_level"), MkMap(&VarMap{}))
	_ = memberInfo
	currencyInfo := this.SafeValue(response, MkString("currency"), MkMap(&VarMap{}))
	_ = currencyInfo
	withdrawLimits := this.SafeValue(response, MkString("withdraw_limit"), MkMap(&VarMap{}))
	_ = withdrawLimits
	canWithdraw := this.SafeValue(withdrawLimits, MkString("can_withdraw"))
	_ = canWithdraw
	walletState := this.SafeString(currencyInfo, MkString("wallet_state"))
	_ = walletState
	walletLocked := this.SafeValue(memberInfo, MkString("wallet_locked"))
	_ = walletLocked
	locked := this.SafeValue(memberInfo, MkString("locked"))
	_ = locked
	active := OpCopy(MkBool(true))
	_ = active
	if IsTrue(OpAnd(OpNotEq2(canWithdraw, MkUndefined()), OpNot(canWithdraw))) {
		active = OpCopy(MkBool(false))
	} else {
		if IsTrue(OpNotEq2(walletState, MkString("working"))) {
			active = OpCopy(MkBool(false))
		} else {
			if IsTrue(OpAnd(OpNotEq2(walletLocked, MkUndefined()), walletLocked)) {
				active = OpCopy(MkBool(false))
			} else {
				if IsTrue(OpAnd(OpNotEq2(locked, MkUndefined()), locked)) {
					active = OpCopy(MkBool(false))
				}
			}
		}
	}
	maxOnetimeWithdrawal := this.SafeNumber(withdrawLimits, MkString("onetime"))
	_ = maxOnetimeWithdrawal
	maxDailyWithdrawal := this.SafeNumber(withdrawLimits, MkString("daily"), maxOnetimeWithdrawal)
	_ = maxDailyWithdrawal
	remainingDailyWithdrawal := this.SafeNumber(withdrawLimits, MkString("remaining_daily"), maxDailyWithdrawal)
	_ = remainingDailyWithdrawal
	maxWithdrawLimit := OpCopy(MkUndefined())
	_ = maxWithdrawLimit
	if IsTrue(OpGt(remainingDailyWithdrawal, MkInteger(0))) {
		maxWithdrawLimit = OpCopy(remainingDailyWithdrawal)
	} else {
		maxWithdrawLimit = OpCopy(maxDailyWithdrawal)
	}
	precision := OpCopy(MkUndefined())
	_ = precision
	currencyId := this.SafeString(currencyInfo, MkString("code"))
	_ = currencyId
	code := this.SafeCurrencyCode(currencyId)
	_ = code
	return MkMap(&VarMap{
		"info":      response,
		"id":        currencyId,
		"code":      code,
		"name":      code,
		"active":    active,
		"fee":       this.SafeNumber(currencyInfo, MkString("withdraw_fee")),
		"precision": precision,
		"limits": MkMap(&VarMap{"withdraw": MkMap(&VarMap{
			"min": this.SafeNumber(withdrawLimits, MkString("minimum")),
			"max": maxWithdrawLimit,
		})}),
	})
}

func (this *Upbit) FetchMarket(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	return this.FetchMarketById(*(market).At(MkString("id")), params)
}

func (this *Upbit) FetchMarketById(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	request := MkMap(&VarMap{"market": id})
	_ = request
	response := this.Call(MkString("privateGetOrdersChance"), this.Extend(request, params))
	_ = response
	marketInfo := this.SafeValue(response, MkString("market"))
	_ = marketInfo
	bid := this.SafeValue(marketInfo, MkString("bid"))
	_ = bid
	ask := this.SafeValue(marketInfo, MkString("ask"))
	_ = ask
	marketId := this.SafeString(marketInfo, MkString("id"))
	_ = marketId
	baseId := this.SafeString(ask, MkString("currency"))
	_ = baseId
	quoteId := this.SafeString(bid, MkString("currency"))
	_ = quoteId
	base := this.SafeCurrencyCode(baseId)
	_ = base
	quote := this.SafeCurrencyCode(quoteId)
	_ = quote
	symbol := OpAdd(base, OpAdd(MkString("/"), quote))
	_ = symbol
	precision := MkMap(&VarMap{
		"amount": MkInteger(8),
		"price":  MkInteger(8),
	})
	_ = precision
	state := this.SafeString(marketInfo, MkString("state"))
	_ = state
	active := OpEq2(state, MkString("active"))
	_ = active
	bidFee := this.SafeNumber(response, MkString("bid_fee"))
	_ = bidFee
	askFee := this.SafeNumber(response, MkString("ask_fee"))
	_ = askFee
	fee := Math.Max(bidFee, askFee)
	_ = fee
	return MkMap(&VarMap{
		"info":      response,
		"id":        marketId,
		"symbol":    symbol,
		"base":      base,
		"quote":     quote,
		"baseId":    baseId,
		"quoteId":   quoteId,
		"active":    active,
		"precision": precision,
		"maker":     fee,
		"taker":     fee,
		"limits": MkMap(&VarMap{
			"amount": MkMap(&VarMap{
				"min": this.SafeNumber(ask, MkString("min_total")),
				"max": MkUndefined(),
			}),
			"price": MkMap(&VarMap{
				"min": Math.Pow(MkInteger(10), OpNeg(*(precision).At(MkString("price")))),
				"max": MkUndefined(),
			}),
			"cost": MkMap(&VarMap{
				"min": this.SafeNumber(bid, MkString("min_total")),
				"max": this.SafeNumber(marketInfo, MkString("max_total")),
			}),
		}),
	})
}

func (this *Upbit) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetMarketAll"), params)
	_ = response
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
		market := *(response).At(i)
		_ = market
		id := this.SafeString(market, MkString("market"))
		_ = id
		quoteId := MkUndefined()
		_ = quoteId
		baseId := MkUndefined()
		_ = baseId
		ArrayUnpack(id.Split(MkString("-")), &quoteId, &baseId)
		base := this.SafeCurrencyCode(baseId)
		_ = base
		quote := this.SafeCurrencyCode(quoteId)
		_ = quote
		symbol := OpAdd(base, OpAdd(MkString("/"), quote))
		_ = symbol
		precision := MkMap(&VarMap{
			"amount": MkInteger(8),
			"price":  MkInteger(8),
		})
		_ = precision
		active := OpCopy(MkBool(true))
		_ = active
		makerFee := this.SafeNumber(*(*this.At(MkString("options"))).At(MkString("tradingFeesByQuoteCurrency")), quote, *(*(*this.At(MkString("fees"))).At(MkString("trading"))).At(MkString("maker")))
		_ = makerFee
		takerFee := this.SafeNumber(*(*this.At(MkString("options"))).At(MkString("tradingFeesByQuoteCurrency")), quote, *(*(*this.At(MkString("fees"))).At(MkString("trading"))).At(MkString("taker")))
		_ = takerFee
		result.Push(MkMap(&VarMap{
			"id":        id,
			"symbol":    symbol,
			"base":      base,
			"quote":     quote,
			"baseId":    baseId,
			"quoteId":   quoteId,
			"active":    active,
			"info":      market,
			"precision": precision,
			"maker":     makerFee,
			"taker":     takerFee,
			"limits": MkMap(&VarMap{
				"amount": MkMap(&VarMap{
					"min": Math.Pow(MkInteger(10), OpNeg(*(precision).At(MkString("amount")))),
					"max": MkUndefined(),
				}),
				"price": MkMap(&VarMap{
					"min": Math.Pow(MkInteger(10), OpNeg(*(precision).At(MkString("price")))),
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

func (this *Upbit) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privateGetAccounts"), params)
	_ = response
	result := MkMap(&VarMap{
		"info":      response,
		"timestamp": MkUndefined(),
		"datetime":  MkUndefined(),
	})
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
		*(account).At(MkString("free")) = this.SafeString(balance, MkString("balance"))
		*(account).At(MkString("used")) = this.SafeString(balance, MkString("locked"))
		*(result).At(code) = OpCopy(account)
	}
	return this.ParseBalance(result)
}

func (this *Upbit) FetchOrderBooks(goArgs ...*Variant) *Variant {
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
		if IsTrue(OpGt(ids.Length, *(*this.At(MkString("options"))).At(MkString("fetchOrderBooksMaxLength")))) {
			numIds := OpCopy((*((this).At(MkString("ids")))).Length)
			_ = numIds
			panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" has "), OpAdd(numIds.ToString(), OpAdd(MkString(" symbols ("), OpAdd(ids.Length.ToString(), OpAdd(MkString(" characters) exceeding max URL length ("), OpAdd((*(*this.At(MkString("options"))).At(MkString("fetchOrderBooksMaxLength"))).Call(MkString("toString")), MkString(" characters), you are required to specify a list of symbols in the first argument to fetchOrderBooks"))))))))))
		}
	} else {
		ids = this.MarketIds(symbols)
		ids = ids.Join(MkString(","))
	}
	request := MkMap(&VarMap{"markets": ids})
	_ = request
	response := this.Call(MkString("publicGetOrderbook"), this.Extend(request, params))
	_ = response
	result := MkMap(&VarMap{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
		orderbook := *(response).At(i)
		_ = orderbook
		marketId := this.SafeString(orderbook, MkString("market"))
		_ = marketId
		symbol := this.SafeSymbol(marketId, MkUndefined(), MkString("-"))
		_ = symbol
		timestamp := this.SafeInteger(orderbook, MkString("timestamp"))
		_ = timestamp
		*(result).At(symbol) = MkMap(&VarMap{
			"symbol":    symbol,
			"bids":      this.SortBy(this.ParseBidsAsks(*(orderbook).At(MkString("orderbook_units")), MkString("bid_price"), MkString("bid_size")), MkInteger(0), MkBool(true)),
			"asks":      this.SortBy(this.ParseBidsAsks(*(orderbook).At(MkString("orderbook_units")), MkString("ask_price"), MkString("ask_size")), MkInteger(0)),
			"timestamp": timestamp,
			"datetime":  this.Iso8601(timestamp),
			"nonce":     MkUndefined(),
		})
	}
	return result
}

func (this *Upbit) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	orderbooks := this.FetchOrderBooks(MkArray(&VarArray{
		symbol,
	}), limit, params)
	_ = orderbooks
	return this.SafeValue(orderbooks, symbol)
}

func (this *Upbit) ParseTicker(goArgs ...*Variant) *Variant {
	ticker := GoGetArg(goArgs, 0, MkUndefined())
	_ = ticker
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.SafeInteger(ticker, MkString("trade_timestamp"))
	_ = timestamp
	marketId := this.SafeString2(ticker, MkString("market"), MkString("code"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market, MkString("-"))
	_ = symbol
	previous := this.SafeNumber(ticker, MkString("prev_closing_price"))
	_ = previous
	last := this.SafeNumber(ticker, MkString("trade_price"))
	_ = last
	change := this.SafeNumber(ticker, MkString("signed_change_price"))
	_ = change
	percentage := this.SafeNumber(ticker, MkString("signed_change_rate"))
	_ = percentage
	return MkMap(&VarMap{
		"symbol":        symbol,
		"timestamp":     timestamp,
		"datetime":      this.Iso8601(timestamp),
		"high":          this.SafeNumber(ticker, MkString("high_price")),
		"low":           this.SafeNumber(ticker, MkString("low_price")),
		"bid":           MkUndefined(),
		"bidVolume":     MkUndefined(),
		"ask":           MkUndefined(),
		"askVolume":     MkUndefined(),
		"vwap":          MkUndefined(),
		"open":          this.SafeNumber(ticker, MkString("opening_price")),
		"close":         last,
		"last":          last,
		"previousClose": previous,
		"change":        change,
		"percentage":    percentage,
		"average":       MkUndefined(),
		"baseVolume":    this.SafeNumber(ticker, MkString("acc_trade_volume_24h")),
		"quoteVolume":   this.SafeNumber(ticker, MkString("acc_trade_price_24h")),
		"info":          ticker,
	})
}

func (this *Upbit) FetchTickers(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	ids := OpCopy(MkUndefined())
	_ = ids
	if IsTrue(OpEq2(symbols, MkUndefined())) {
		ids = (*((this).At(MkString("ids")))).Call(MkString("join"), MkString(","))
		if IsTrue(OpGt(ids.Length, *(*this.At(MkString("options"))).At(MkString("fetchTickersMaxLength")))) {
			numIds := OpCopy((*((this).At(MkString("ids")))).Length)
			_ = numIds
			panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" has "), OpAdd(numIds.ToString(), MkString(" symbols exceeding max URL length, you are required to specify a list of symbols in the first argument to fetchTickers"))))))
		}
	} else {
		ids = this.MarketIds(symbols)
		ids = ids.Join(MkString(","))
	}
	request := MkMap(&VarMap{"markets": ids})
	_ = request
	response := this.Call(MkString("publicGetTicker"), this.Extend(request, params))
	_ = response
	result := MkMap(&VarMap{})
	_ = result
	for t := MkInteger(0); IsTrue(OpLw(t, response.Length)); OpIncr(&t) {
		ticker := this.ParseTicker(*(response).At(t))
		_ = ticker
		symbol := *(ticker).At(MkString("symbol"))
		_ = symbol
		*(result).At(symbol) = OpCopy(ticker)
	}
	return this.FilterByArray(result, MkString("symbol"), symbols)
}

func (this *Upbit) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	tickers := this.FetchTickers(MkArray(&VarArray{
		symbol,
	}), params)
	_ = tickers
	return this.SafeValue(tickers, symbol)
}

func (this *Upbit) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	id := this.SafeString2(trade, MkString("sequential_id"), MkString("uuid"))
	_ = id
	orderId := OpCopy(MkUndefined())
	_ = orderId
	timestamp := this.SafeInteger(trade, MkString("timestamp"))
	_ = timestamp
	if IsTrue(OpEq2(timestamp, MkUndefined())) {
		timestamp = this.Parse8601(this.SafeString(trade, MkString("created_at")))
	}
	side := OpCopy(MkUndefined())
	_ = side
	askOrBid := this.SafeStringLower2(trade, MkString("ask_bid"), MkString("side"))
	_ = askOrBid
	if IsTrue(OpEq2(askOrBid, MkString("ask"))) {
		side = MkString("sell")
	} else {
		if IsTrue(OpEq2(askOrBid, MkString("bid"))) {
			side = MkString("buy")
		}
	}
	cost := this.SafeNumber(trade, MkString("funds"))
	_ = cost
	priceString := this.SafeString2(trade, MkString("trade_price"), MkString("price"))
	_ = priceString
	amountString := this.SafeString2(trade, MkString("trade_volume"), MkString("volume"))
	_ = amountString
	price := this.ParseNumber(priceString)
	_ = price
	amount := this.ParseNumber(amountString)
	_ = amount
	if IsTrue(OpEq2(cost, MkUndefined())) {
		cost = this.ParseNumber(Precise.StringMul(priceString, amountString))
	}
	marketId := this.SafeString2(trade, MkString("market"), MkString("code"))
	_ = marketId
	market = this.SafeMarket(marketId, market)
	fee := OpCopy(MkUndefined())
	_ = fee
	feeCurrency := OpCopy(MkUndefined())
	_ = feeCurrency
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(OpNotEq2(market, MkUndefined())) {
		symbol = *(market).At(MkString("symbol"))
		feeCurrency = *(market).At(MkString("quote"))
	} else {
		baseId := MkUndefined()
		_ = baseId
		quoteId := MkUndefined()
		_ = quoteId
		ArrayUnpack(marketId.Split(MkString("-")), &baseId, &quoteId)
		base := this.SafeCurrencyCode(baseId)
		_ = base
		quote := this.SafeCurrencyCode(quoteId)
		_ = quote
		symbol = OpAdd(base, OpAdd(MkString("/"), quote))
		feeCurrency = OpCopy(quote)
	}
	feeCost := this.SafeString(trade, OpAdd(askOrBid, MkString("_fee")))
	_ = feeCost
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		fee = MkMap(&VarMap{
			"currency": feeCurrency,
			"cost":     feeCost,
		})
	}
	return MkMap(&VarMap{
		"id":           id,
		"info":         trade,
		"order":        orderId,
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       symbol,
		"type":         MkUndefined(),
		"side":         side,
		"takerOrMaker": MkUndefined(),
		"price":        price,
		"amount":       amount,
		"cost":         cost,
		"fee":          fee,
	})
}

func (this *Upbit) FetchTrades(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpEq2(limit, MkUndefined())) {
		limit = MkInteger(200)
	}
	request := MkMap(&VarMap{
		"market": *(market).At(MkString("id")),
		"count":  limit,
	})
	_ = request
	response := this.Call(MkString("publicGetTradesTicks"), this.Extend(request, params))
	_ = response
	return this.ParseTrades(response, market, since, limit)
}

func (this *Upbit) ParseOHLCV(goArgs ...*Variant) *Variant {
	ohlcv := GoGetArg(goArgs, 0, MkUndefined())
	_ = ohlcv
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	return MkArray(&VarArray{
		this.Parse8601(this.SafeString(ohlcv, MkString("candle_date_time_utc"))),
		this.SafeNumber(ohlcv, MkString("opening_price")),
		this.SafeNumber(ohlcv, MkString("high_price")),
		this.SafeNumber(ohlcv, MkString("low_price")),
		this.SafeNumber(ohlcv, MkString("trade_price")),
		this.SafeNumber(ohlcv, MkString("candle_acc_trade_volume")),
	})
}

func (this *Upbit) FetchOHLCV(goArgs ...*Variant) *Variant {
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
	timeframePeriod := this.ParseTimeframe(timeframe)
	_ = timeframePeriod
	timeframeValue := *(*this.At(MkString("timeframes"))).At(timeframe)
	_ = timeframeValue
	if IsTrue(OpEq2(limit, MkUndefined())) {
		limit = MkInteger(200)
	}
	request := MkMap(&VarMap{
		"market":    *(market).At(MkString("id")),
		"timeframe": timeframeValue,
		"count":     limit,
	})
	_ = request
	method := MkString("publicGetCandlesTimeframe")
	_ = method
	if IsTrue(OpEq2(timeframeValue, MkString("minutes"))) {
		numMinutes := Math.Round(OpDiv(timeframePeriod, MkInteger(60)))
		_ = numMinutes
		*(request).At(MkString("unit")) = OpCopy(numMinutes)
		OpAddAssign(&method, MkString("Unit"))
	}
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("to")) = this.Iso8601(this.Sum(since, OpMulti(timeframePeriod, OpMulti(limit, MkInteger(1000)))))
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	return this.ParseOHLCVs(response, market, timeframe, since, limit)
}

func (this *Upbit) CreateOrder(goArgs ...*Variant) *Variant {
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
		if IsTrue(OpEq2(side, MkString("buy"))) {
			if IsTrue(*(*this.At(MkString("options"))).At(MkString("createMarketBuyOrderRequiresPrice"))) {
				if IsTrue(OpEq2(price, MkUndefined())) {
					panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")), MkString(" createOrder() requires the price argument with market buy orders to calculate total order cost (amount to spend), where cost = amount * price. Supply a price argument to createOrder() call if you want the cost to be calculated for you from price and amount, or, alternatively, add .options['createMarketBuyOrderRequiresPrice'] = false to supply the cost in the amount argument (the exchange-specific behaviour)"))))
				} else {
					amount = OpMulti(amount, price)
				}
			}
		}
	}
	orderSide := OpCopy(MkUndefined())
	_ = orderSide
	if IsTrue(OpEq2(side, MkString("buy"))) {
		orderSide = MkString("bid")
	} else {
		if IsTrue(OpEq2(side, MkString("sell"))) {
			orderSide = MkString("ask")
		} else {
			panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")), MkString(" createOrder allows buy or sell side only!"))))
		}
	}
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{
		"market": *(market).At(MkString("id")),
		"side":   orderSide,
	})
	_ = request
	if IsTrue(OpEq2(type_, MkString("limit"))) {
		*(request).At(MkString("volume")) = this.AmountToPrecision(symbol, amount)
		*(request).At(MkString("price")) = this.PriceToPrecision(symbol, price)
		*(request).At(MkString("ord_type")) = OpCopy(type_)
	} else {
		if IsTrue(OpEq2(type_, MkString("market"))) {
			if IsTrue(OpEq2(side, MkString("buy"))) {
				*(request).At(MkString("ord_type")) = MkString("price")
				*(request).At(MkString("price")) = this.PriceToPrecision(symbol, amount)
			} else {
				if IsTrue(OpEq2(side, MkString("sell"))) {
					*(request).At(MkString("ord_type")) = OpCopy(type_)
					*(request).At(MkString("volume")) = this.AmountToPrecision(symbol, amount)
				}
			}
		}
	}
	clientOrderId := this.SafeString2(params, MkString("clientOrderId"), MkString("identifier"))
	_ = clientOrderId
	if IsTrue(OpNotEq2(clientOrderId, MkUndefined())) {
		*(request).At(MkString("identifier")) = OpCopy(clientOrderId)
	}
	params = this.Omit(params, MkArray(&VarArray{
		MkString("clientOrderId"),
		MkString("identifier"),
	}))
	response := this.Call(MkString("privatePostOrders"), this.Extend(request, params))
	_ = response
	return this.ParseOrder(response)
}

func (this *Upbit) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"uuid": id})
	_ = request
	response := this.Call(MkString("privateDeleteOrder"), this.Extend(request, params))
	_ = response
	return this.ParseOrder(response)
}

func (this *Upbit) FetchDeposits(goArgs ...*Variant) *Variant {
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
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("privateGetDeposits"), this.Extend(request, params))
	_ = response
	return this.ParseTransactions(response, currency, since, limit)
}

func (this *Upbit) FetchWithdrawals(goArgs ...*Variant) *Variant {
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
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("privateGetWithdraws"), this.Extend(request, params))
	_ = response
	return this.ParseTransactions(response, currency, since, limit)
}

func (this *Upbit) ParseTransactionStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"submitting":      MkString("pending"),
		"submitted":       MkString("pending"),
		"almost_accepted": MkString("pending"),
		"rejected":        MkString("failed"),
		"accepted":        MkString("pending"),
		"processing":      MkString("pending"),
		"done":            MkString("ok"),
		"canceled":        MkString("canceled"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Upbit) ParseTransaction(goArgs ...*Variant) *Variant {
	transaction := GoGetArg(goArgs, 0, MkUndefined())
	_ = transaction
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	id := this.SafeString(transaction, MkString("uuid"))
	_ = id
	amount := this.SafeNumber(transaction, MkString("amount"))
	_ = amount
	address := OpCopy(MkUndefined())
	_ = address
	tag := OpCopy(MkUndefined())
	_ = tag
	txid := this.SafeString(transaction, MkString("txid"))
	_ = txid
	updated := this.Parse8601(this.SafeString(transaction, MkString("done_at")))
	_ = updated
	timestamp := this.Parse8601(this.SafeString(transaction, MkString("created_at"), updated))
	_ = timestamp
	type_ := this.SafeString(transaction, MkString("type"))
	_ = type_
	if IsTrue(OpEq2(type_, MkString("withdraw"))) {
		type_ = MkString("withdrawal")
	}
	currencyId := this.SafeString(transaction, MkString("currency"))
	_ = currencyId
	code := this.SafeCurrencyCode(currencyId)
	_ = code
	status := this.ParseTransactionStatus(this.SafeStringLower(transaction, MkString("state")))
	_ = status
	feeCost := this.SafeNumber(transaction, MkString("fee"))
	_ = feeCost
	return MkMap(&VarMap{
		"info":      transaction,
		"id":        id,
		"currency":  code,
		"amount":    amount,
		"address":   address,
		"tag":       tag,
		"status":    status,
		"type":      type_,
		"updated":   updated,
		"txid":      txid,
		"timestamp": timestamp,
		"datetime":  this.Iso8601(timestamp),
		"fee": MkMap(&VarMap{
			"currency": code,
			"cost":     feeCost,
		}),
	})
}

func (this *Upbit) ParseOrderStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"wait":   MkString("open"),
		"done":   MkString("closed"),
		"cancel": MkString("canceled"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Upbit) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	id := this.SafeString(order, MkString("uuid"))
	_ = id
	side := this.SafeString(order, MkString("side"))
	_ = side
	if IsTrue(OpEq2(side, MkString("bid"))) {
		side = MkString("buy")
	} else {
		side = MkString("sell")
	}
	type_ := this.SafeString(order, MkString("ord_type"))
	_ = type_
	timestamp := this.Parse8601(this.SafeString(order, MkString("created_at")))
	_ = timestamp
	status := this.ParseOrderStatus(this.SafeString(order, MkString("state")))
	_ = status
	lastTradeTimestamp := OpCopy(MkUndefined())
	_ = lastTradeTimestamp
	price := this.SafeNumber(order, MkString("price"))
	_ = price
	amount := this.SafeNumber(order, MkString("volume"))
	_ = amount
	remaining := this.SafeNumber(order, MkString("remaining_volume"))
	_ = remaining
	filled := this.SafeNumber(order, MkString("executed_volume"))
	_ = filled
	cost := OpCopy(MkUndefined())
	_ = cost
	if IsTrue(OpEq2(type_, MkString("price"))) {
		type_ = MkString("market")
		cost = OpCopy(price)
		price = OpCopy(MkUndefined())
	}
	average := OpCopy(MkUndefined())
	_ = average
	fee := OpCopy(MkUndefined())
	_ = fee
	feeCost := this.SafeNumber(order, MkString("paid_fee"))
	_ = feeCost
	marketId := this.SafeString(order, MkString("market"))
	_ = marketId
	market = this.SafeMarket(marketId, market)
	trades := this.SafeValue(order, MkString("trades"), MkArray(&VarArray{}))
	_ = trades
	trades = this.ParseTrades(trades, market, MkUndefined(), MkUndefined(), MkMap(&VarMap{
		"order": id,
		"type":  type_,
	}))
	numTrades := OpCopy(trades.Length)
	_ = numTrades
	if IsTrue(OpGt(numTrades, MkInteger(0))) {
		lastTradeTimestamp = *(*(trades).At(OpSub(numTrades, MkInteger(1)))).At(MkString("timestamp"))
		getFeesFromTrades := OpCopy(MkBool(false))
		_ = getFeesFromTrades
		if IsTrue(OpEq2(feeCost, MkUndefined())) {
			getFeesFromTrades = OpCopy(MkBool(true))
			feeCost = MkInteger(0)
		}
		cost = MkInteger(0)
		for i := MkInteger(0); IsTrue(OpLw(i, numTrades)); OpIncr(&i) {
			trade := *(trades).At(i)
			_ = trade
			cost = this.Sum(cost, *(trade).At(MkString("cost")))
			if IsTrue(getFeesFromTrades) {
				tradeFee := this.SafeValue(*(trades).At(i), MkString("fee"), MkMap(&VarMap{}))
				_ = tradeFee
				tradeFeeCost := this.SafeNumber(tradeFee, MkString("cost"))
				_ = tradeFeeCost
				if IsTrue(OpNotEq2(tradeFeeCost, MkUndefined())) {
					feeCost = this.Sum(feeCost, tradeFeeCost)
				}
			}
		}
		average = OpDiv(cost, filled)
	}
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		fee = MkMap(&VarMap{
			"currency": *(market).At(MkString("quote")),
			"cost":     feeCost,
		})
	}
	result := MkMap(&VarMap{
		"info":               order,
		"id":                 id,
		"clientOrderId":      MkUndefined(),
		"timestamp":          timestamp,
		"datetime":           this.Iso8601(timestamp),
		"lastTradeTimestamp": lastTradeTimestamp,
		"symbol":             *(market).At(MkString("symbol")),
		"type":               type_,
		"timeInForce":        MkUndefined(),
		"postOnly":           MkUndefined(),
		"side":               side,
		"price":              price,
		"stopPrice":          MkUndefined(),
		"cost":               cost,
		"average":            average,
		"amount":             amount,
		"filled":             filled,
		"remaining":          remaining,
		"status":             status,
		"fee":                fee,
		"trades":             trades,
	})
	_ = result
	return result
}

func (this *Upbit) FetchOrdersByState(goArgs ...*Variant) *Variant {
	state := GoGetArg(goArgs, 0, MkUndefined())
	_ = state
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 2, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 3, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 4, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"state": state})
	_ = request
	market := OpCopy(MkUndefined())
	_ = market
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
		*(request).At(MkString("market")) = *(market).At(MkString("id"))
	}
	response := this.Call(MkString("privateGetOrders"), this.Extend(request, params))
	_ = response
	return this.ParseOrders(response, market, since, limit)
}

func (this *Upbit) FetchOpenOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	return this.FetchOrdersByState(MkString("wait"), symbol, since, limit, params)
}

func (this *Upbit) FetchClosedOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	return this.FetchOrdersByState(MkString("done"), symbol, since, limit, params)
}

func (this *Upbit) FetchCanceledOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	return this.FetchOrdersByState(MkString("cancel"), symbol, since, limit, params)
}

func (this *Upbit) FetchOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"uuid": id})
	_ = request
	response := this.Call(MkString("privateGetOrder"), this.Extend(request, params))
	_ = response
	return this.ParseOrder(response)
}

func (this *Upbit) FetchDepositAddresses(goArgs ...*Variant) *Variant {
	codes := GoGetArg(goArgs, 0, MkUndefined())
	_ = codes
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privateGetDepositsCoinAddresses"), params)
	_ = response
	return this.ParseDepositAddresses(response)
}

func (this *Upbit) ParseDepositAddress(goArgs ...*Variant) *Variant {
	depositAddress := GoGetArg(goArgs, 0, MkUndefined())
	_ = depositAddress
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	address := this.SafeString(depositAddress, MkString("deposit_address"))
	_ = address
	tag := this.SafeString(depositAddress, MkString("secondary_address"))
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

func (this *Upbit) FetchDepositAddress(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	response := this.Call(MkString("privateGetDepositsCoinAddress"), this.Extend(MkMap(&VarMap{"currency": *(currency).At(MkString("id"))}), params))
	_ = response
	return this.ParseDepositAddress(response)
}

func (this *Upbit) CreateDepositAddress(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{"currency": *(currency).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("privatePostDepositsGenerateCoinAddress"), this.Extend(request, params))
	_ = response
	message := this.SafeString(response, MkString("message"))
	_ = message
	if IsTrue(OpNotEq2(message, MkUndefined())) {
		panic(NewAddressPending(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" is generating "), OpAdd(code, MkString(" deposit address, call fetchDepositAddress or createDepositAddress one more time later to retrieve the generated address"))))))
	}
	return this.ParseDepositAddress(response)
}

func (this *Upbit) Withdraw(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"amount": amount})
	_ = request
	method := MkString("privatePostWithdraws")
	_ = method
	if IsTrue(OpNotEq2(code, MkString("KRW"))) {
		OpAddAssign(&method, MkString("Coin"))
		*(request).At(MkString("currency")) = *(currency).At(MkString("id"))
		*(request).At(MkString("address")) = OpCopy(address)
		if IsTrue(OpNotEq2(tag, MkUndefined())) {
			*(request).At(MkString("secondary_address")) = OpCopy(tag)
		}
	} else {
		OpAddAssign(&method, MkString("Krw"))
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	return this.ParseTransaction(response)
}

func (this *Upbit) Nonce(goArgs ...*Variant) *Variant {
	return this.Milliseconds()
}

func (this *Upbit) Sign(goArgs ...*Variant) *Variant {
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
	url := this.ImplodeParams(*(*(*this.At(MkString("urls"))).At(MkString("api"))).At(api), MkMap(&VarMap{"hostname": *this.At(MkString("hostname"))}))
	_ = url
	OpAddAssign(&url, OpAdd(MkString("/"), OpAdd(*this.At(MkString("version")), OpAdd(MkString("/"), this.ImplodeParams(path, params)))))
	query := this.Omit(params, this.ExtractParams(path))
	_ = query
	if IsTrue(OpNotEq2(method, MkString("POST"))) {
		if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
			OpAddAssign(&url, OpAdd(MkString("?"), this.Urlencode(query)))
		}
	}
	if IsTrue(OpEq2(api, MkString("private"))) {
		this.CheckRequiredCredentials()
		nonce := this.Nonce()
		_ = nonce
		request := MkMap(&VarMap{
			"access_key": *this.At(MkString("apiKey")),
			"nonce":      nonce,
		})
		_ = request
		if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
			auth := this.Urlencode(query)
			_ = auth
			hash := this.Hash(this.Encode(auth), MkString("sha512"))
			_ = hash
			*(request).At(MkString("query_hash")) = OpCopy(hash)
			*(request).At(MkString("query_hash_alg")) = MkString("SHA512")
		}
		jwt := this.Jwt(request, this.Encode(*this.At(MkString("secret"))))
		_ = jwt
		headers = MkMap(&VarMap{"Authorization": OpAdd(MkString("Bearer "), jwt)})
		if IsTrue(OpAnd(OpNotEq2(method, MkString("GET")), OpNotEq2(method, MkString("DELETE")))) {
			body = this.Json(params)
			*(headers).At(MkString("Content-Type")) = MkString("application/json")
		}
	}
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Upbit) HandleErrors(goArgs ...*Variant) *Variant {
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
	error := this.SafeValue(response, MkString("error"))
	_ = error
	if IsTrue(OpNotEq2(error, MkUndefined())) {
		message := this.SafeString(error, MkString("message"))
		_ = message
		name := this.SafeString(error, MkString("name"))
		_ = name
		feedback := OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))
		_ = feedback
		this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("exact")), message, feedback)
		this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("exact")), name, feedback)
		this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("broad")), message, feedback)
		this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("broad")), name, feedback)
		panic(NewExchangeError(feedback))
	}
	return MkUndefined()
}
