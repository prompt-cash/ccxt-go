package ccxt

import ()

type Deribit struct {
	*ExchangeBase
}

var _ Exchange = (*Deribit)(nil)

func init() {
	exchange := &Deribit{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Deribit) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("deribit"),
		"name": MkString("Deribit"),
		"countries": MkArray(&VarArray{
			MkString("NL"),
		}),
		"version":   MkString("v2"),
		"userAgent": MkUndefined(),
		"rateLimit": MkInteger(500),
		"has": MkMap(&VarMap{
			"cancelAllOrders":      MkBool(true),
			"cancelOrder":          MkBool(true),
			"CORS":                 MkBool(true),
			"createDepositAddress": MkBool(true),
			"createOrder":          MkBool(true),
			"editOrder":            MkBool(true),
			"fetchBalance":         MkBool(true),
			"fetchClosedOrders":    MkBool(true),
			"fetchDepositAddress":  MkBool(true),
			"fetchDeposits":        MkBool(true),
			"fetchMarkets":         MkBool(true),
			"fetchMyTrades":        MkBool(true),
			"fetchOHLCV":           MkBool(true),
			"fetchOpenOrders":      MkBool(true),
			"fetchOrder":           MkBool(true),
			"fetchOrderBook":       MkBool(true),
			"fetchOrders":          MkBool(false),
			"fetchOrderTrades":     MkBool(true),
			"fetchStatus":          MkBool(true),
			"fetchTicker":          MkBool(true),
			"fetchTickers":         MkBool(true),
			"fetchTime":            MkBool(true),
			"fetchTrades":          MkBool(true),
			"fetchTransactions":    MkBool(false),
			"fetchWithdrawals":     MkBool(true),
			"withdraw":             MkBool(true),
		}),
		"timeframes": MkMap(&VarMap{
			"1m":  MkString("1"),
			"3m":  MkString("3"),
			"5m":  MkString("5"),
			"10m": MkString("10"),
			"15m": MkString("15"),
			"30m": MkString("30"),
			"1h":  MkString("60"),
			"2h":  MkString("120"),
			"3h":  MkString("180"),
			"6h":  MkString("360"),
			"12h": MkString("720"),
			"1d":  MkString("1D"),
		}),
		"urls": MkMap(&VarMap{
			"test": MkString("https://test.deribit.com"),
			"logo": MkString("https://user-images.githubusercontent.com/1294454/41933112-9e2dd65a-798b-11e8-8440-5bab2959fcb8.jpg"),
			"api":  MkString("https://www.deribit.com"),
			"www":  MkString("https://www.deribit.com"),
			"doc": MkArray(&VarArray{
				MkString("https://docs.deribit.com/v2"),
				MkString("https://github.com/deribit"),
			}),
			"fees": MkString("https://www.deribit.com/pages/information/fees"),
			"referral": MkMap(&VarMap{
				"url":      MkString("https://www.deribit.com/reg-1189.4038"),
				"discount": MkNumber(0.1),
			}),
		}),
		"api": MkMap(&VarMap{
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("auth"),
				MkString("exchange_token"),
				MkString("fork_token"),
				MkString("set_heartbeat"),
				MkString("disable_heartbeat"),
				MkString("get_time"),
				MkString("hello"),
				MkString("test"),
				MkString("subscribe"),
				MkString("unsubscribe"),
				MkString("get_announcements"),
				MkString("get_book_summary_by_currency"),
				MkString("get_book_summary_by_instrument"),
				MkString("get_contract_size"),
				MkString("get_currencies"),
				MkString("get_funding_chart_data"),
				MkString("get_funding_rate_history"),
				MkString("get_funding_rate_value"),
				MkString("get_historical_volatility"),
				MkString("get_index"),
				MkString("get_index_price"),
				MkString("get_index_price_names"),
				MkString("get_instruments"),
				MkString("get_last_settlements_by_currency"),
				MkString("get_last_settlements_by_instrument"),
				MkString("get_last_trades_by_currency"),
				MkString("get_last_trades_by_currency_and_time"),
				MkString("get_last_trades_by_instrument"),
				MkString("get_last_trades_by_instrument_and_time"),
				MkString("get_order_book"),
				MkString("get_trade_volumes"),
				MkString("get_tradingview_chart_data"),
				MkString("ticker"),
			})}),
			"private": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("logout"),
				MkString("enable_cancel_on_disconnect"),
				MkString("disable_cancel_on_disconnect"),
				MkString("get_cancel_on_disconnect"),
				MkString("subscribe"),
				MkString("unsubscribe"),
				MkString("change_api_key_name"),
				MkString("change_scope_in_api_key"),
				MkString("change_subaccount_name"),
				MkString("create_api_key"),
				MkString("create_subaccount"),
				MkString("disable_api_key"),
				MkString("disable_tfa_for_subaccount"),
				MkString("enable_api_key"),
				MkString("get_account_summary"),
				MkString("get_email_language"),
				MkString("get_new_announcements"),
				MkString("get_position"),
				MkString("get_positions"),
				MkString("get_subaccounts"),
				MkString("list_api_keys"),
				MkString("remove_api_key"),
				MkString("reset_api_key"),
				MkString("set_announcement_as_read"),
				MkString("set_api_key_as_default"),
				MkString("set_email_for_subaccount"),
				MkString("set_email_language"),
				MkString("set_password_for_subaccount"),
				MkString("toggle_notifications_from_subaccount"),
				MkString("toggle_subaccount_login"),
				MkString("execute_block_trade"),
				MkString("get_block_trade"),
				MkString("get_last_block_trades_by_currency"),
				MkString("invalidate_block_trade_signature"),
				MkString("verify_block_trade"),
				MkString("buy"),
				MkString("sell"),
				MkString("edit"),
				MkString("cancel"),
				MkString("cancel_all"),
				MkString("cancel_all_by_currency"),
				MkString("cancel_all_by_instrument"),
				MkString("cancel_by_label"),
				MkString("close_position"),
				MkString("get_margins"),
				MkString("get_open_orders_by_currency"),
				MkString("get_open_orders_by_instrument"),
				MkString("get_order_history_by_currency"),
				MkString("get_order_history_by_instrument"),
				MkString("get_order_margin_by_ids"),
				MkString("get_order_state"),
				MkString("get_stop_order_history"),
				MkString("get_user_trades_by_currency"),
				MkString("get_user_trades_by_currency_and_time"),
				MkString("get_user_trades_by_instrument"),
				MkString("get_user_trades_by_instrument_and_time"),
				MkString("get_user_trades_by_order"),
				MkString("get_settlement_history_by_instrument"),
				MkString("get_settlement_history_by_currency"),
				MkString("cancel_transfer_by_id"),
				MkString("cancel_withdrawal"),
				MkString("create_deposit_address"),
				MkString("get_current_deposit_address"),
				MkString("get_deposits"),
				MkString("get_transfers"),
				MkString("get_withdrawals"),
				MkString("submit_transfer_to_subaccount"),
				MkString("submit_transfer_to_user"),
				MkString("withdraw"),
			})}),
		}),
		"exceptions": MkMap(&VarMap{
			"9999":   PermissionDenied,
			"10000":  AuthenticationError,
			"10001":  ExchangeError,
			"10002":  InvalidOrder,
			"10003":  InvalidOrder,
			"10004":  OrderNotFound,
			"10005":  InvalidOrder,
			"10006":  InvalidOrder,
			"10007":  InvalidOrder,
			"10008":  InvalidOrder,
			"10009":  InsufficientFunds,
			"10010":  OrderNotFound,
			"10011":  InvalidOrder,
			"10012":  InvalidOrder,
			"10013":  PermissionDenied,
			"10014":  PermissionDenied,
			"10015":  PermissionDenied,
			"10016":  PermissionDenied,
			"10017":  PermissionDenied,
			"10018":  PermissionDenied,
			"10019":  PermissionDenied,
			"10020":  ExchangeError,
			"10021":  InvalidOrder,
			"10022":  InvalidOrder,
			"10023":  InvalidOrder,
			"10024":  InvalidOrder,
			"10025":  InvalidOrder,
			"10026":  InvalidOrder,
			"10027":  InvalidOrder,
			"10028":  DDoSProtection,
			"10029":  OrderNotFound,
			"10030":  ExchangeError,
			"10031":  ExchangeError,
			"10032":  InvalidOrder,
			"10033":  NotSupported,
			"10034":  InvalidOrder,
			"10035":  InvalidOrder,
			"10036":  InvalidOrder,
			"10040":  ExchangeNotAvailable,
			"10041":  OnMaintenance,
			"10043":  InvalidOrder,
			"10044":  InvalidOrder,
			"10045":  InvalidOrder,
			"10046":  InvalidOrder,
			"10047":  DDoSProtection,
			"10048":  ExchangeError,
			"11008":  InvalidOrder,
			"11029":  BadRequest,
			"11030":  ExchangeError,
			"11031":  ExchangeError,
			"11035":  DDoSProtection,
			"11036":  InvalidOrder,
			"11037":  BadRequest,
			"11038":  InvalidOrder,
			"11039":  InvalidOrder,
			"11041":  InvalidOrder,
			"11042":  PermissionDenied,
			"11043":  BadRequest,
			"11044":  InvalidOrder,
			"11045":  BadRequest,
			"11046":  BadRequest,
			"11047":  BadRequest,
			"11048":  ExchangeError,
			"11049":  BadRequest,
			"11050":  BadRequest,
			"11051":  OnMaintenance,
			"11052":  ExchangeError,
			"11053":  ExchangeError,
			"11090":  InvalidAddress,
			"11091":  InvalidAddress,
			"11092":  InvalidAddress,
			"11093":  DDoSProtection,
			"11094":  ExchangeError,
			"11095":  ExchangeError,
			"11096":  ExchangeError,
			"12000":  AuthenticationError,
			"12001":  DDoSProtection,
			"12002":  ExchangeError,
			"12998":  AuthenticationError,
			"12003":  AuthenticationError,
			"12004":  AuthenticationError,
			"12005":  AuthenticationError,
			"12100":  ExchangeError,
			"12999":  AuthenticationError,
			"13000":  AuthenticationError,
			"13001":  AuthenticationError,
			"13002":  PermissionDenied,
			"13003":  AuthenticationError,
			"13004":  AuthenticationError,
			"13005":  AuthenticationError,
			"13006":  AuthenticationError,
			"13007":  AuthenticationError,
			"13008":  ExchangeError,
			"13009":  AuthenticationError,
			"13010":  BadRequest,
			"13011":  BadRequest,
			"13012":  PermissionDenied,
			"13013":  BadRequest,
			"13014":  BadRequest,
			"13015":  BadRequest,
			"13016":  BadRequest,
			"13017":  ExchangeError,
			"13018":  ExchangeError,
			"13019":  ExchangeError,
			"13020":  ExchangeError,
			"13021":  PermissionDenied,
			"13025":  ExchangeError,
			"-32602": BadRequest,
			"-32601": BadRequest,
			"-32700": BadRequest,
			"-32000": BadRequest,
		}),
		"precisionMode": TICK_SIZE,
		"options": MkMap(&VarMap{
			"code":         MkString("BTC"),
			"fetchBalance": MkMap(&VarMap{"code": MkString("BTC")}),
		}),
	}))
}

func (this *Deribit) FetchTime(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetGetTime"), params)
	_ = response
	return this.SafeInteger(response, MkString("result"))
}

func (this *Deribit) CodeFromOptions(goArgs ...*Variant) *Variant {
	methodName := GoGetArg(goArgs, 0, MkUndefined())
	_ = methodName
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	defaultCode := this.SafeValue(*this.At(MkString("options")), MkString("code"), MkString("BTC"))
	_ = defaultCode
	options := this.SafeValue(*this.At(MkString("options")), methodName, MkMap(&VarMap{}))
	_ = options
	code := this.SafeValue(options, MkString("code"), defaultCode)
	_ = code
	return this.SafeValue(params, MkString("code"), code)
}

func (this *Deribit) FetchStatus(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	request := MkMap(&VarMap{})
	_ = request
	this.Call(MkString("publicGetTest"), this.Extend(request, params))
	*this.At(MkString("status")) = this.Extend(*this.At(MkString("status")), MkMap(&VarMap{
		"status":  MkString("ok"),
		"updated": this.Milliseconds(),
	}))
	return *this.At(MkString("status"))
}

func (this *Deribit) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	currenciesResponse := this.Call(MkString("publicGetGetCurrencies"), params)
	_ = currenciesResponse
	currenciesResult := this.SafeValue(currenciesResponse, MkString("result"), MkArray(&VarArray{}))
	_ = currenciesResult
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, currenciesResult.Length)); OpIncr(&i) {
		currencyId := this.SafeString(*(currenciesResult).At(i), MkString("currency"))
		_ = currencyId
		request := MkMap(&VarMap{"currency": currencyId})
		_ = request
		instrumentsResponse := this.Call(MkString("publicGetGetInstruments"), this.Extend(request, params))
		_ = instrumentsResponse
		instrumentsResult := this.SafeValue(instrumentsResponse, MkString("result"), MkArray(&VarArray{}))
		_ = instrumentsResult
		for k := MkInteger(0); IsTrue(OpLw(k, instrumentsResult.Length)); OpIncr(&k) {
			market := *(instrumentsResult).At(k)
			_ = market
			id := this.SafeString(market, MkString("instrument_name"))
			_ = id
			baseId := this.SafeString(market, MkString("base_currency"))
			_ = baseId
			quoteId := this.SafeString(market, MkString("quote_currency"))
			_ = quoteId
			base := this.SafeCurrencyCode(baseId)
			_ = base
			quote := this.SafeCurrencyCode(quoteId)
			_ = quote
			type_ := this.SafeString(market, MkString("kind"))
			_ = type_
			future := OpEq2(type_, MkString("future"))
			_ = future
			option := OpEq2(type_, MkString("option"))
			_ = option
			active := this.SafeValue(market, MkString("is_active"))
			_ = active
			minTradeAmount := this.SafeNumber(market, MkString("min_trade_amount"))
			_ = minTradeAmount
			tickSize := this.SafeNumber(market, MkString("tick_size"))
			_ = tickSize
			precision := MkMap(&VarMap{
				"amount": minTradeAmount,
				"price":  tickSize,
			})
			_ = precision
			result.Push(MkMap(&VarMap{
				"id":        id,
				"symbol":    id,
				"base":      base,
				"quote":     quote,
				"active":    active,
				"precision": precision,
				"taker":     this.SafeNumber(market, MkString("taker_commission")),
				"maker":     this.SafeNumber(market, MkString("maker_commission")),
				"limits": MkMap(&VarMap{
					"amount": MkMap(&VarMap{
						"min": minTradeAmount,
						"max": MkUndefined(),
					}),
					"price": MkMap(&VarMap{
						"min": tickSize,
						"max": MkUndefined(),
					}),
					"cost": MkMap(&VarMap{
						"min": MkUndefined(),
						"max": MkUndefined(),
					}),
				}),
				"type":   type_,
				"spot":   MkBool(false),
				"future": future,
				"option": option,
				"info":   market,
			}))
		}
	}
	return result
}

func (this *Deribit) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	code := this.CodeFromOptions(MkString("fetchBalance"), params)
	_ = code
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{"currency": *(currency).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("privateGetGetAccountSummary"), this.Extend(request, params))
	_ = response
	result := MkMap(&VarMap{"info": response})
	_ = result
	balance := this.SafeValue(response, MkString("result"), MkMap(&VarMap{}))
	_ = balance
	currencyId := this.SafeString(balance, MkString("currency"))
	_ = currencyId
	currencyCode := this.SafeCurrencyCode(currencyId)
	_ = currencyCode
	account := this.Account()
	_ = account
	*(account).At(MkString("free")) = this.SafeString(balance, MkString("availableFunds"))
	*(account).At(MkString("used")) = this.SafeString(balance, MkString("maintenanceMargin"))
	*(account).At(MkString("total")) = this.SafeString(balance, MkString("equity"))
	*(result).At(currencyCode) = OpCopy(account)
	return this.ParseBalance(result)
}

func (this *Deribit) CreateDepositAddress(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{"currency": *(currency).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("privateGetCreateDepositAddress"), this.Extend(request, params))
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

func (this *Deribit) FetchDepositAddress(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{"currency": *(currency).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("privateGetGetCurrentDepositAddress"), this.Extend(request, params))
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

func (this *Deribit) ParseTicker(goArgs ...*Variant) *Variant {
	ticker := GoGetArg(goArgs, 0, MkUndefined())
	_ = ticker
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.SafeInteger2(ticker, MkString("timestamp"), MkString("creation_timestamp"))
	_ = timestamp
	marketId := this.SafeString(ticker, MkString("instrument_name"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market)
	_ = symbol
	last := this.SafeNumber2(ticker, MkString("last_price"), MkString("last"))
	_ = last
	stats := this.SafeValue(ticker, MkString("stats"), ticker)
	_ = stats
	return MkMap(&VarMap{
		"symbol":        symbol,
		"timestamp":     timestamp,
		"datetime":      this.Iso8601(timestamp),
		"high":          this.SafeNumber2(stats, MkString("high"), MkString("max_price")),
		"low":           this.SafeNumber2(stats, MkString("low"), MkString("min_price")),
		"bid":           this.SafeNumber2(ticker, MkString("best_bid_price"), MkString("bid_price")),
		"bidVolume":     this.SafeNumber(ticker, MkString("best_bid_amount")),
		"ask":           this.SafeNumber2(ticker, MkString("best_ask_price"), MkString("ask_price")),
		"askVolume":     this.SafeNumber(ticker, MkString("best_ask_amount")),
		"vwap":          MkUndefined(),
		"open":          MkUndefined(),
		"close":         last,
		"last":          last,
		"previousClose": MkUndefined(),
		"change":        MkUndefined(),
		"percentage":    MkUndefined(),
		"average":       MkUndefined(),
		"baseVolume":    MkUndefined(),
		"quoteVolume":   this.SafeNumber(stats, MkString("volume")),
		"info":          ticker,
	})
}

func (this *Deribit) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"instrument_name": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("publicGetTicker"), this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"))
	_ = result
	return this.ParseTicker(result, market)
}

func (this *Deribit) FetchTickers(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	code := this.CodeFromOptions(MkString("fetchTickers"), params)
	_ = code
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{"currency": *(currency).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("publicGetGetBookSummaryByCurrency"), this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"), MkArray(&VarArray{}))
	_ = result
	tickers := MkMap(&VarMap{})
	_ = tickers
	for i := MkInteger(0); IsTrue(OpLw(i, result.Length)); OpIncr(&i) {
		ticker := this.ParseTicker(*(result).At(i))
		_ = ticker
		symbol := *(ticker).At(MkString("symbol"))
		_ = symbol
		*(tickers).At(symbol) = OpCopy(ticker)
	}
	return this.FilterByArray(tickers, MkString("symbol"), symbols)
}

func (this *Deribit) FetchOHLCV(goArgs ...*Variant) *Variant {
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
		"instrument_name": *(market).At(MkString("id")),
		"resolution":      *(*this.At(MkString("timeframes"))).At(timeframe),
	})
	_ = request
	duration := this.ParseTimeframe(timeframe)
	_ = duration
	now := this.Milliseconds()
	_ = now
	if IsTrue(OpEq2(since, MkUndefined())) {
		if IsTrue(OpEq2(limit, MkUndefined())) {
			panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchOHLCV() requires a since argument or a limit argument"))))
		} else {
			*(request).At(MkString("start_timestamp")) = OpSub(now, OpMulti(OpSub(limit, MkInteger(1)), OpMulti(duration, MkInteger(1000))))
			*(request).At(MkString("end_timestamp")) = OpCopy(now)
		}
	} else {
		*(request).At(MkString("start_timestamp")) = OpCopy(since)
		if IsTrue(OpEq2(limit, MkUndefined())) {
			*(request).At(MkString("end_timestamp")) = OpCopy(now)
		} else {
			*(request).At(MkString("end_timestamp")) = this.Sum(since, OpMulti(limit, OpMulti(duration, MkInteger(1000))))
		}
	}
	response := this.Call(MkString("publicGetGetTradingviewChartData"), this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"), MkMap(&VarMap{}))
	_ = result
	ohlcvs := this.ConvertTradingViewToOHLCV(result, MkString("ticks"), MkString("open"), MkString("high"), MkString("low"), MkString("close"), MkString("volume"), MkBool(true))
	_ = ohlcvs
	return this.ParseOHLCVs(ohlcvs, market, timeframe, since, limit)
}

func (this *Deribit) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	id := this.SafeString(trade, MkString("trade_id"))
	_ = id
	marketId := this.SafeString(trade, MkString("instrument_name"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market)
	_ = symbol
	timestamp := this.SafeInteger(trade, MkString("timestamp"))
	_ = timestamp
	side := this.SafeString(trade, MkString("direction"))
	_ = side
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
	liquidity := this.SafeString(trade, MkString("liquidity"))
	_ = liquidity
	takerOrMaker := OpCopy(MkUndefined())
	_ = takerOrMaker
	if IsTrue(OpNotEq2(liquidity, MkUndefined())) {
		takerOrMaker = OpTrinary(OpEq2(liquidity, MkString("M")), MkString("maker"), MkString("taker"))
	}
	feeCost := this.SafeNumber(trade, MkString("fee"))
	_ = feeCost
	fee := OpCopy(MkUndefined())
	_ = fee
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		feeCurrencyId := this.SafeString(trade, MkString("fee_currency"))
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
		"info":         trade,
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       symbol,
		"order":        this.SafeString(trade, MkString("order_id")),
		"type":         this.SafeString(trade, MkString("order_type")),
		"side":         side,
		"takerOrMaker": takerOrMaker,
		"price":        price,
		"amount":       amount,
		"cost":         cost,
		"fee":          fee,
	})
}

func (this *Deribit) FetchTrades(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{
		"instrument_name": *(market).At(MkString("id")),
		"include_old":     MkBool(true),
	})
	_ = request
	method := OpTrinary(OpEq2(since, MkUndefined()), MkString("publicGetGetLastTradesByInstrument"), MkString("publicGetGetLastTradesByInstrumentAndTime"))
	_ = method
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("start_timestamp")) = OpCopy(since)
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("count")) = OpCopy(limit)
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"), MkMap(&VarMap{}))
	_ = result
	trades := this.SafeValue(result, MkString("trades"), MkArray(&VarArray{}))
	_ = trades
	return this.ParseTrades(trades, market, since, limit)
}

func (this *Deribit) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"instrument_name": *(market).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("depth")) = OpCopy(limit)
	}
	response := this.Call(MkString("publicGetGetOrderBook"), this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"), MkMap(&VarMap{}))
	_ = result
	timestamp := this.SafeInteger(result, MkString("timestamp"))
	_ = timestamp
	nonce := this.SafeInteger(result, MkString("change_id"))
	_ = nonce
	orderbook := this.ParseOrderBook(result, symbol, timestamp)
	_ = orderbook
	*(orderbook).At(MkString("nonce")) = OpCopy(nonce)
	return orderbook
}

func (this *Deribit) ParseOrderStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"open":        MkString("open"),
		"cancelled":   MkString("canceled"),
		"filled":      MkString("closed"),
		"rejected":    MkString("rejected"),
		"untriggered": MkString("open"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Deribit) ParseTimeInForce(goArgs ...*Variant) *Variant {
	timeInForce := GoGetArg(goArgs, 0, MkUndefined())
	_ = timeInForce
	timeInForces := MkMap(&VarMap{
		"good_til_cancelled":  MkString("GTC"),
		"fill_or_kill":        MkString("FOK"),
		"immediate_or_cancel": MkString("IOC"),
	})
	_ = timeInForces
	return this.SafeString(timeInForces, timeInForce, timeInForce)
}

func (this *Deribit) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.SafeInteger(order, MkString("creation_timestamp"))
	_ = timestamp
	lastUpdate := this.SafeInteger(order, MkString("last_update_timestamp"))
	_ = lastUpdate
	id := this.SafeString(order, MkString("order_id"))
	_ = id
	price := this.SafeNumber(order, MkString("price"))
	_ = price
	average := this.SafeNumber(order, MkString("average_price"))
	_ = average
	amount := this.SafeNumber(order, MkString("amount"))
	_ = amount
	filled := this.SafeNumber(order, MkString("filled_amount"))
	_ = filled
	lastTradeTimestamp := OpCopy(MkUndefined())
	_ = lastTradeTimestamp
	if IsTrue(OpNotEq2(filled, MkUndefined())) {
		if IsTrue(OpGt(filled, MkInteger(0))) {
			lastTradeTimestamp = OpCopy(lastUpdate)
		}
	}
	status := this.ParseOrderStatus(this.SafeString(order, MkString("order_state")))
	_ = status
	marketId := this.SafeString(order, MkString("instrument_name"))
	_ = marketId
	market = this.SafeMarket(marketId, market)
	side := this.SafeStringLower(order, MkString("direction"))
	_ = side
	feeCost := this.SafeNumber(order, MkString("commission"))
	_ = feeCost
	fee := OpCopy(MkUndefined())
	_ = fee
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		feeCost = Math.Abs(feeCost)
		fee = MkMap(&VarMap{
			"cost":     feeCost,
			"currency": *(market).At(MkString("base")),
		})
	}
	type_ := this.SafeString(order, MkString("order_type"))
	_ = type_
	trades := this.SafeValue(order, MkString("trades"))
	_ = trades
	if IsTrue(OpNotEq2(trades, MkUndefined())) {
		trades = this.ParseTrades(trades, market)
	}
	timeInForce := this.ParseTimeInForce(this.SafeString(order, MkString("time_in_force")))
	_ = timeInForce
	stopPrice := this.SafeValue(order, MkString("stop_price"))
	_ = stopPrice
	postOnly := this.SafeValue(order, MkString("post_only"))
	_ = postOnly
	return this.SafeOrder(MkMap(&VarMap{
		"info":               order,
		"id":                 id,
		"clientOrderId":      MkUndefined(),
		"timestamp":          timestamp,
		"datetime":           this.Iso8601(timestamp),
		"lastTradeTimestamp": lastTradeTimestamp,
		"symbol":             *(market).At(MkString("symbol")),
		"type":               type_,
		"timeInForce":        timeInForce,
		"postOnly":           postOnly,
		"side":               side,
		"price":              price,
		"stopPrice":          stopPrice,
		"amount":             amount,
		"cost":               MkUndefined(),
		"average":            average,
		"filled":             filled,
		"remaining":          MkUndefined(),
		"status":             status,
		"fee":                fee,
		"trades":             trades,
	}))
}

func (this *Deribit) FetchOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"order_id": id})
	_ = request
	response := this.Call(MkString("privateGetGetOrderState"), this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"))
	_ = result
	return this.ParseOrder(result)
}

func (this *Deribit) CreateOrder(goArgs ...*Variant) *Variant {
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
		"instrument_name": *(market).At(MkString("id")),
		"amount":          this.AmountToPrecision(symbol, amount),
		"type":            type_,
	})
	_ = request
	priceIsRequired := OpCopy(MkBool(false))
	_ = priceIsRequired
	stopPriceIsRequired := OpCopy(MkBool(false))
	_ = stopPriceIsRequired
	if IsTrue(OpEq2(type_, MkString("limit"))) {
		priceIsRequired = OpCopy(MkBool(true))
	} else {
		if IsTrue(OpEq2(type_, MkString("stop_limit"))) {
			priceIsRequired = OpCopy(MkBool(true))
			stopPriceIsRequired = OpCopy(MkBool(true))
		}
	}
	if IsTrue(priceIsRequired) {
		if IsTrue(OpNotEq2(price, MkUndefined())) {
			*(request).At(MkString("price")) = this.PriceToPrecision(symbol, price)
		} else {
			panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" createOrder() requires a price argument for a "), OpAdd(type_, MkString(" order"))))))
		}
	}
	if IsTrue(stopPriceIsRequired) {
		stopPrice := this.SafeNumber2(params, MkString("stop_price"), MkString("stopPrice"))
		_ = stopPrice
		if IsTrue(OpEq2(stopPrice, MkUndefined())) {
			panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" createOrder() requires a stop_price or stopPrice param for a "), OpAdd(type_, MkString(" order"))))))
		} else {
			*(request).At(MkString("stop_price")) = this.PriceToPrecision(symbol, stopPrice)
		}
		params = this.Omit(params, MkArray(&VarArray{
			MkString("stop_price"),
			MkString("stopPrice"),
		}))
	}
	method := OpAdd(MkString("privateGet"), this.Capitalize(side))
	_ = method
	response := this.Call(method, this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"), MkMap(&VarMap{}))
	_ = result
	order := this.SafeValue(result, MkString("order"))
	_ = order
	trades := this.SafeValue(result, MkString("trades"), MkArray(&VarArray{}))
	_ = trades
	*(order).At(MkString("trades")) = OpCopy(trades)
	return this.ParseOrder(order, market)
}

func (this *Deribit) EditOrder(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpEq2(amount, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" editOrder() requires an amount argument"))))
	}
	if IsTrue(OpEq2(price, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" editOrder() requires a price argument"))))
	}
	this.LoadMarkets()
	request := MkMap(&VarMap{
		"order_id": id,
		"amount":   this.AmountToPrecision(symbol, amount),
		"price":    this.PriceToPrecision(symbol, price),
	})
	_ = request
	response := this.Call(MkString("privateGetEdit"), this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"), MkMap(&VarMap{}))
	_ = result
	order := this.SafeValue(result, MkString("order"))
	_ = order
	trades := this.SafeValue(result, MkString("trades"), MkArray(&VarArray{}))
	_ = trades
	*(order).At(MkString("trades")) = OpCopy(trades)
	return this.ParseOrder(order)
}

func (this *Deribit) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"order_id": id})
	_ = request
	response := this.Call(MkString("privateGetCancel"), this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"), MkMap(&VarMap{}))
	_ = result
	return this.ParseOrder(result)
}

func (this *Deribit) CancelAllOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{})
	_ = request
	method := OpCopy(MkUndefined())
	_ = method
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		method = MkString("privateGetCancelAll")
	} else {
		method = MkString("privateGetCancelAllByInstrument")
		market := this.Market(symbol)
		_ = market
		*(request).At(MkString("instrument_name")) = *(market).At(MkString("id"))
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	return response
}

func (this *Deribit) FetchOpenOrders(goArgs ...*Variant) *Variant {
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
	method := OpCopy(MkUndefined())
	_ = method
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		code := this.CodeFromOptions(MkString("fetchOpenOrders"), params)
		_ = code
		currency := this.Currency(code)
		_ = currency
		*(request).At(MkString("currency")) = *(currency).At(MkString("id"))
		method = MkString("privateGetGetOpenOrdersByCurrency")
	} else {
		market = this.Market(symbol)
		*(request).At(MkString("instrument_name")) = *(market).At(MkString("id"))
		method = MkString("privateGetGetOpenOrdersByInstrument")
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"), MkArray(&VarArray{}))
	_ = result
	return this.ParseOrders(result, market, since, limit)
}

func (this *Deribit) FetchClosedOrders(goArgs ...*Variant) *Variant {
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
	method := OpCopy(MkUndefined())
	_ = method
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		code := this.CodeFromOptions(MkString("fetchClosedOrders"), params)
		_ = code
		currency := this.Currency(code)
		_ = currency
		*(request).At(MkString("currency")) = *(currency).At(MkString("id"))
		method = MkString("privateGetGetOrderHistoryByCurrency")
	} else {
		market = this.Market(symbol)
		*(request).At(MkString("instrument_name")) = *(market).At(MkString("id"))
		method = MkString("privateGetGetOrderHistoryByInstrument")
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"), MkArray(&VarArray{}))
	_ = result
	return this.ParseOrders(result, market, since, limit)
}

func (this *Deribit) FetchOrderTrades(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"order_id": id})
	_ = request
	response := this.Call(MkString("privateGetGetUserTradesByOrder"), this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"), MkMap(&VarMap{}))
	_ = result
	trades := this.SafeValue(result, MkString("trades"), MkArray(&VarArray{}))
	_ = trades
	return this.ParseTrades(trades, MkUndefined(), since, limit)
}

func (this *Deribit) FetchMyTrades(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"include_old": MkBool(true)})
	_ = request
	market := OpCopy(MkUndefined())
	_ = market
	method := OpCopy(MkUndefined())
	_ = method
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		code := this.CodeFromOptions(MkString("fetchMyTrades"), params)
		_ = code
		currency := this.Currency(code)
		_ = currency
		*(request).At(MkString("currency")) = *(currency).At(MkString("id"))
		if IsTrue(OpEq2(since, MkUndefined())) {
			method = MkString("privateGetGetUserTradesByCurrency")
		} else {
			method = MkString("privateGetGetUserTradesByCurrencyAndTime")
			*(request).At(MkString("start_timestamp")) = OpCopy(since)
		}
	} else {
		market = this.Market(symbol)
		*(request).At(MkString("instrument_name")) = *(market).At(MkString("id"))
		if IsTrue(OpEq2(since, MkUndefined())) {
			method = MkString("privateGetGetUserTradesByInstrument")
		} else {
			method = MkString("privateGetGetUserTradesByInstrumentAndTime")
			*(request).At(MkString("start_timestamp")) = OpCopy(since)
		}
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("count")) = OpCopy(limit)
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"), MkMap(&VarMap{}))
	_ = result
	trades := this.SafeValue(result, MkString("trades"), MkArray(&VarArray{}))
	_ = trades
	return this.ParseTrades(trades, market, since, limit)
}

func (this *Deribit) FetchDeposits(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(code, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchDeposits() requires a currency code argument"))))
	}
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{"currency": *(currency).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("count")) = OpCopy(limit)
	}
	response := this.Call(MkString("privateGetGetDeposits"), this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"), MkMap(&VarMap{}))
	_ = result
	data := this.SafeValue(result, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseTransactions(data, currency, since, limit, params)
}

func (this *Deribit) FetchWithdrawals(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(code, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchWithdrawals() requires a currency code argument"))))
	}
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{"currency": *(currency).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("count")) = OpCopy(limit)
	}
	response := this.Call(MkString("privateGetGetWithdrawals"), this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"), MkMap(&VarMap{}))
	_ = result
	data := this.SafeValue(result, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseTransactions(data, currency, since, limit, params)
}

func (this *Deribit) ParseTransactionStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"completed":   MkString("ok"),
		"unconfirmed": MkString("pending"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Deribit) ParseTransaction(goArgs ...*Variant) *Variant {
	transaction := GoGetArg(goArgs, 0, MkUndefined())
	_ = transaction
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	currencyId := this.SafeString(transaction, MkString("currency"))
	_ = currencyId
	code := this.SafeCurrencyCode(currencyId, currency)
	_ = code
	timestamp := this.SafeInteger2(transaction, MkString("created_timestamp"), MkString("received_timestamp"))
	_ = timestamp
	updated := this.SafeInteger(transaction, MkString("updated_timestamp"))
	_ = updated
	status := this.ParseTransactionStatus(this.SafeString(transaction, MkString("state")))
	_ = status
	address := this.SafeString(transaction, MkString("address"))
	_ = address
	feeCost := this.SafeNumber(transaction, MkString("fee"))
	_ = feeCost
	type_ := MkString("deposit")
	_ = type_
	fee := OpCopy(MkUndefined())
	_ = fee
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		type_ = MkString("withdrawal")
		fee = MkMap(&VarMap{
			"cost":     feeCost,
			"currency": code,
		})
	}
	return MkMap(&VarMap{
		"info":        transaction,
		"id":          this.SafeString(transaction, MkString("id")),
		"txid":        this.SafeString(transaction, MkString("transaction_id")),
		"timestamp":   timestamp,
		"datetime":    this.Iso8601(timestamp),
		"address":     address,
		"addressTo":   address,
		"addressFrom": MkUndefined(),
		"tag":         MkUndefined(),
		"tagTo":       MkUndefined(),
		"tagFrom":     MkUndefined(),
		"type":        type_,
		"amount":      this.SafeNumber(transaction, MkString("amount")),
		"currency":    code,
		"status":      status,
		"updated":     updated,
		"fee":         fee,
	})
}

func (this *Deribit) FetchPosition(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"instrument_name": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("privateGetGetPosition"), this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"))
	_ = result
	return result
}

func (this *Deribit) FetchPositions(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	code := this.CodeFromOptions(MkString("fetchPositions"), params)
	_ = code
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{"currency": *(currency).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("privateGetGetPositions"), this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"), MkArray(&VarArray{}))
	_ = result
	return result
}

func (this *Deribit) Withdraw(goArgs ...*Variant) *Variant {
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
		"address":  address,
		"amount":   amount,
	})
	_ = request
	if IsTrue(OpNotEq2(*this.At(MkString("twofa")), MkUndefined())) {
		*(request).At(MkString("tfa")) = this.Oath()
	}
	response := this.Call(MkString("privateGetWithdraw"), this.Extend(request, params))
	_ = response
	return MkMap(&VarMap{
		"info": response,
		"id":   this.SafeString(response, MkString("id")),
	})
}

func (this *Deribit) Nonce(goArgs ...*Variant) *Variant {
	return this.Milliseconds()
}

func (this *Deribit) Sign(goArgs ...*Variant) *Variant {
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
	request := OpAdd(MkString("/"), OpAdd(MkString("api/"), OpAdd(*this.At(MkString("version")), OpAdd(MkString("/"), OpAdd(api, OpAdd(MkString("/"), path))))))
	_ = request
	if IsTrue(OpEq2(api, MkString("public"))) {
		if IsTrue(*(GoGetKeys(params)).At(MkString("length"))) {
			OpAddAssign(&request, OpAdd(MkString("?"), this.Urlencode(params)))
		}
	}
	if IsTrue(OpEq2(api, MkString("private"))) {
		this.CheckRequiredCredentials()
		nonce := (this.Nonce()).Call(MkString("toString"))
		_ = nonce
		timestamp := (this.Milliseconds()).Call(MkString("toString"))
		_ = timestamp
		requestBody := MkString("")
		_ = requestBody
		if IsTrue(*(GoGetKeys(params)).At(MkString("length"))) {
			OpAddAssign(&request, OpAdd(MkString("?"), this.Urlencode(params)))
		}
		requestData := OpAdd(method, OpAdd(MkString("\n"), OpAdd(request, OpAdd(MkString("\n"), OpAdd(requestBody, MkString("\n"))))))
		_ = requestData
		auth := OpAdd(timestamp, OpAdd(MkString("\n"), OpAdd(nonce, OpAdd(MkString("\n"), requestData))))
		_ = auth
		signature := this.Hmac(this.Encode(auth), this.Encode(*this.At(MkString("secret"))), MkString("sha256"))
		_ = signature
		headers = MkMap(&VarMap{"Authorization": OpAdd(MkString("deri-hmac-sha256 id="), OpAdd(*this.At(MkString("apiKey")), OpAdd(MkString(",ts="), OpAdd(timestamp, OpAdd(MkString(",sig="), OpAdd(signature, OpAdd(MkString(","), OpAdd(MkString("nonce="), nonce))))))))})
	}
	url := OpAdd(*(*this.At(MkString("urls"))).At(MkString("api")), request)
	_ = url
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Deribit) HandleErrors(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpNot(response)) {
		return MkUndefined()
	}
	error := this.SafeValue(response, MkString("error"))
	_ = error
	if IsTrue(OpNotEq2(error, MkUndefined())) {
		errorCode := this.SafeString(error, MkString("code"))
		_ = errorCode
		feedback := OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))
		_ = feedback
		this.ThrowExactlyMatchedException(*this.At(MkString("exceptions")), errorCode, feedback)
		panic(NewExchangeError(feedback))
	}
	return MkUndefined()
}
