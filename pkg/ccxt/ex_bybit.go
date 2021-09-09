package ccxt

import ()

type Bybit struct {
	*ExchangeBase
}

var _ Exchange = (*Bybit)(nil)

func init() {
	exchange := &Bybit{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Bybit) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("bybit"),
		"name": MkString("Bybit"),
		"countries": MkArray(&VarArray{
			MkString("VG"),
		}),
		"version":   MkString("v2"),
		"userAgent": MkUndefined(),
		"rateLimit": MkInteger(100),
		"hostname":  MkString("bybit.com"),
		"has": MkMap(&VarMap{
			"cancelOrder":       MkBool(true),
			"CORS":              MkBool(true),
			"cancelAllOrders":   MkBool(true),
			"createOrder":       MkBool(true),
			"editOrder":         MkBool(true),
			"fetchBalance":      MkBool(true),
			"fetchClosedOrders": MkBool(true),
			"fetchDeposits":     MkBool(true),
			"fetchLedger":       MkBool(true),
			"fetchMarkets":      MkBool(true),
			"fetchMyTrades":     MkBool(true),
			"fetchOHLCV":        MkBool(true),
			"fetchOpenOrders":   MkBool(true),
			"fetchOrder":        MkBool(true),
			"fetchOrderBook":    MkBool(true),
			"fetchOrders":       MkBool(true),
			"fetchOrderTrades":  MkBool(true),
			"fetchTicker":       MkBool(true),
			"fetchTickers":      MkBool(true),
			"fetchTime":         MkBool(true),
			"fetchTrades":       MkBool(true),
			"fetchTransactions": MkBool(false),
			"fetchWithdrawals":  MkBool(true),
			"fetchPositions":    MkBool(true),
		}),
		"timeframes": MkMap(&VarMap{
			"1m":  MkString("1"),
			"3m":  MkString("3"),
			"5m":  MkString("5"),
			"15m": MkString("15"),
			"30m": MkString("30"),
			"1h":  MkString("60"),
			"2h":  MkString("120"),
			"4h":  MkString("240"),
			"6h":  MkString("360"),
			"12h": MkString("720"),
			"1d":  MkString("D"),
			"1w":  MkString("W"),
			"1M":  MkString("M"),
			"1y":  MkString("Y"),
		}),
		"urls": MkMap(&VarMap{
			"test": MkMap(&VarMap{
				"futures": MkString("https://api-testnet.{hostname}"),
				"v2":      MkString("https://api-testnet.{hostname}"),
				"public":  MkString("https://api-testnet.{hostname}"),
				"private": MkString("https://api-testnet.{hostname}"),
			}),
			"logo": MkString("https://user-images.githubusercontent.com/51840849/76547799-daff5b80-649e-11ea-87fb-3be9bac08954.jpg"),
			"api": MkMap(&VarMap{
				"futures": MkString("https://api.{hostname}"),
				"v2":      MkString("https://api.{hostname}"),
				"public":  MkString("https://api.{hostname}"),
				"private": MkString("https://api.{hostname}"),
			}),
			"www": MkString("https://www.bybit.com"),
			"doc": MkArray(&VarArray{
				MkString("https://bybit-exchange.github.io/docs/inverse/"),
				MkString("https://bybit-exchange.github.io/docs/linear/"),
				MkString("https://github.com/bybit-exchange"),
			}),
			"fees":     MkString("https://help.bybit.com/hc/en-us/articles/360039261154"),
			"referral": MkString("https://www.bybit.com/app/register?ref=X7Prm"),
		}),
		"api": MkMap(&VarMap{
			"futures": MkMap(&VarMap{"private": MkMap(&VarMap{
				"get": MkArray(&VarArray{
					MkString("order/list"),
					MkString("order"),
					MkString("stop-order/list"),
					MkString("stop-order"),
					MkString("position/list"),
					MkString("execution/list"),
					MkString("trade/closed-pnl/list"),
				}),
				"post": MkArray(&VarArray{
					MkString("order/create"),
					MkString("order/cancel"),
					MkString("order/cancelAll"),
					MkString("order/replace"),
					MkString("stop-order/create"),
					MkString("stop-order/cancel"),
					MkString("stop-order/cancelAll"),
					MkString("stop-order/replace"),
					MkString("position/change-position-margin"),
					MkString("position/trading-stop"),
					MkString("position/leverage/save"),
					MkString("position/switch-mode"),
					MkString("position/switch-isolated"),
					MkString("position/risk-limit"),
				}),
			})}),
			"v2": MkMap(&VarMap{
				"public": MkMap(&VarMap{"get": MkArray(&VarArray{
					MkString("orderBook/L2"),
					MkString("kline/list"),
					MkString("tickers"),
					MkString("trading-records"),
					MkString("symbols"),
					MkString("liq-records"),
					MkString("mark-price-kline"),
					MkString("index-price-kline"),
					MkString("premium-index-kline"),
					MkString("open-interest"),
					MkString("big-deal"),
					MkString("account-ratio"),
					MkString("time"),
					MkString("announcement"),
					MkString("funding/prev-funding-rate"),
					MkString("risk-limit/list"),
				})}),
				"private": MkMap(&VarMap{
					"get": MkArray(&VarArray{
						MkString("order/list"),
						MkString("order"),
						MkString("stop-order/list"),
						MkString("stop-order"),
						MkString("position/list"),
						MkString("execution/list"),
						MkString("trade/closed-pnl/list"),
						MkString("funding/prev-funding-rate"),
						MkString("funding/prev-funding"),
						MkString("funding/predicted-funding"),
						MkString("account/api-key"),
						MkString("account/lcp"),
						MkString("wallet/balance"),
						MkString("wallet/fund/records"),
						MkString("wallet/withdraw/list"),
						MkString("exchange-order/list"),
					}),
					"post": MkArray(&VarArray{
						MkString("order/create"),
						MkString("order/cancel"),
						MkString("order/cancelAll"),
						MkString("order/replace"),
						MkString("stop-order/create"),
						MkString("stop-order/cancel"),
						MkString("stop-order/cancelAll"),
						MkString("stop-order/replace"),
						MkString("position/change-position-margin"),
						MkString("position/trading-stop"),
						MkString("position/leverage/save"),
						MkString("position/switch-mode"),
						MkString("position/switch-isolated"),
						MkString("position/risk-limit"),
					}),
				}),
			}),
			"public": MkMap(&VarMap{"linear": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("kline"),
				MkString("recent-trading-records"),
				MkString("funding/prev-funding-rate"),
				MkString("mark-price-kline"),
				MkString("index-price-kline"),
				MkString("premium-index-kline"),
				MkString("risk-limit"),
			})})}),
			"private": MkMap(&VarMap{"linear": MkMap(&VarMap{
				"get": MkArray(&VarArray{
					MkString("order/list"),
					MkString("order/search"),
					MkString("stop-order/list"),
					MkString("stop-order/search"),
					MkString("position/list"),
					MkString("trade/execution/list"),
					MkString("trade/closed-pnl/list"),
					MkString("funding/predicted-funding"),
					MkString("funding/prev-funding"),
				}),
				"post": MkArray(&VarArray{
					MkString("order/create"),
					MkString("order/cancel"),
					MkString("order/cancel-all"),
					MkString("order/replace"),
					MkString("stop-order/create"),
					MkString("stop-order/cancel"),
					MkString("stop-order/cancel-all"),
					MkString("stop-order/replace"),
					MkString("position/set-auto-add-margin"),
					MkString("position/switch-isolated"),
					MkString("tpsl/switch-mode"),
					MkString("position/add-margin"),
					MkString("position/set-leverage"),
					MkString("position/trading-stop"),
					MkString("position/set-risk"),
				}),
			})}),
		}),
		"httpExceptions": MkMap(&VarMap{"403": RateLimitExceeded}),
		"exceptions": MkMap(&VarMap{
			"exact": MkMap(&VarMap{
				"10001": BadRequest,
				"10002": InvalidNonce,
				"10003": AuthenticationError,
				"10004": AuthenticationError,
				"10005": PermissionDenied,
				"10006": RateLimitExceeded,
				"10007": AuthenticationError,
				"10010": PermissionDenied,
				"10017": BadRequest,
				"10018": RateLimitExceeded,
				"20001": OrderNotFound,
				"20003": InvalidOrder,
				"20004": InvalidOrder,
				"20005": InvalidOrder,
				"20006": InvalidOrder,
				"20007": InvalidOrder,
				"20008": InvalidOrder,
				"20009": InvalidOrder,
				"20010": InvalidOrder,
				"20011": InvalidOrder,
				"20012": InvalidOrder,
				"20013": InvalidOrder,
				"20014": InvalidOrder,
				"20015": InvalidOrder,
				"20016": InvalidOrder,
				"20017": InvalidOrder,
				"20018": InvalidOrder,
				"20019": InvalidOrder,
				"20020": InvalidOrder,
				"20021": InvalidOrder,
				"20022": BadRequest,
				"20023": BadRequest,
				"20031": BadRequest,
				"20070": BadRequest,
				"20071": BadRequest,
				"20084": BadRequest,
				"30001": BadRequest,
				"30003": InvalidOrder,
				"30004": InvalidOrder,
				"30005": InvalidOrder,
				"30007": InvalidOrder,
				"30008": InvalidOrder,
				"30009": ExchangeError,
				"30010": InsufficientFunds,
				"30011": PermissionDenied,
				"30012": PermissionDenied,
				"30013": PermissionDenied,
				"30014": InvalidOrder,
				"30015": InvalidOrder,
				"30016": ExchangeError,
				"30017": InvalidOrder,
				"30018": InvalidOrder,
				"30019": InvalidOrder,
				"30020": InvalidOrder,
				"30021": InvalidOrder,
				"30022": InvalidOrder,
				"30023": InvalidOrder,
				"30024": InvalidOrder,
				"30025": InvalidOrder,
				"30026": InvalidOrder,
				"30027": InvalidOrder,
				"30028": InvalidOrder,
				"30029": InvalidOrder,
				"30030": InvalidOrder,
				"30031": InsufficientFunds,
				"30032": InvalidOrder,
				"30033": RateLimitExceeded,
				"30034": OrderNotFound,
				"30035": RateLimitExceeded,
				"30036": ExchangeError,
				"30037": InvalidOrder,
				"30041": ExchangeError,
				"30042": InsufficientFunds,
				"30043": InvalidOrder,
				"30044": InvalidOrder,
				"30045": InvalidOrder,
				"30049": InsufficientFunds,
				"30050": ExchangeError,
				"30051": ExchangeError,
				"30052": ExchangeError,
				"30054": ExchangeError,
				"30057": ExchangeError,
				"30063": ExchangeError,
				"30067": InsufficientFunds,
				"30068": ExchangeError,
				"30074": InvalidOrder,
				"30075": InvalidOrder,
				"33004": AuthenticationError,
				"34026": ExchangeError,
			}),
			"broad": MkMap(&VarMap{
				"unknown orderInfo": OrderNotFound,
				"invalid api_key":   AuthenticationError,
			}),
		}),
		"precisionMode": TICK_SIZE,
		"options": MkMap(&VarMap{
			"marketTypes": MkMap(&VarMap{
				"BTC/USDT":   MkString("linear"),
				"ETH/USDT":   MkString("linear"),
				"BNB/USDT":   MkString("linear"),
				"ADA/USDT":   MkString("linear"),
				"DOGE/USDT":  MkString("linear"),
				"XRP/USDT":   MkString("linear"),
				"DOT/USDT":   MkString("linear"),
				"UNI/USDT":   MkString("linear"),
				"BCH/USDT":   MkString("linear"),
				"LTC/USDT":   MkString("linear"),
				"SOL/USDT":   MkString("linear"),
				"LINK/USDT":  MkString("linear"),
				"MATIC/USDT": MkString("linear"),
				"ETC/USDT":   MkString("linear"),
				"FIL/USDT":   MkString("linear"),
				"EOS/USDT":   MkString("linear"),
				"AAVE/USDT":  MkString("linear"),
				"XTZ/USDT":   MkString("linear"),
				"SUSHI/USDT": MkString("linear"),
				"XEM/USDT":   MkString("linear"),
				"BTC/USD":    MkString("inverse"),
				"ETH/USD":    MkString("inverse"),
				"EOS/USD":    MkString("inverse"),
				"XRP/USD":    MkString("inverse"),
			}),
			"defaultType":             MkString("linear"),
			"code":                    MkString("BTC"),
			"cancelAllOrders":         MkMap(&VarMap{}),
			"recvWindow":              OpMulti(MkInteger(5), MkInteger(1000)),
			"timeDifference":          MkInteger(0),
			"adjustForTimeDifference": MkBool(false),
		}),
		"fees": MkMap(&VarMap{
			"trading": MkMap(&VarMap{
				"tierBased":  MkBool(false),
				"percentage": MkBool(true),
				"taker":      MkNumber(0.00075),
				"maker":      OpNeg(MkNumber(0.00025)),
			}),
			"funding": MkMap(&VarMap{
				"tierBased":  MkBool(false),
				"percentage": MkBool(false),
				"withdraw":   MkMap(&VarMap{}),
				"deposit":    MkMap(&VarMap{}),
			}),
		}),
	}))
}

func (this *Bybit) Nonce(goArgs ...*Variant) *Variant {
	return OpSub(this.Milliseconds(), *(*this.At(MkString("options"))).At(MkString("timeDifference")))
}

func (this *Bybit) LoadTimeDifference(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	serverTime := this.FetchTime(params)
	_ = serverTime
	after := this.Milliseconds()
	_ = after
	*(*this.At(MkString("options"))).At(MkString("timeDifference")) = OpSub(after, serverTime)
	return *(*this.At(MkString("options"))).At(MkString("timeDifference"))
}

func (this *Bybit) FetchTime(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("v2PublicGetTime"), params)
	_ = response
	return this.SafeTimestamp(response, MkString("time_now"))
}

func (this *Bybit) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	if IsTrue(*(*this.At(MkString("options"))).At(MkString("adjustForTimeDifference"))) {
		this.LoadTimeDifference()
	}
	response := this.Call(MkString("v2PublicGetSymbols"), params)
	_ = response
	markets := this.SafeValue(response, MkString("result"), MkArray(&VarArray{}))
	_ = markets
	options := this.SafeValue(*this.At(MkString("options")), MkString("fetchMarkets"), MkMap(&VarMap{}))
	_ = options
	linearQuoteCurrencies := this.SafeValue(options, MkString("linear"), MkMap(&VarMap{"USDT": MkBool(true)}))
	_ = linearQuoteCurrencies
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, markets.Length)); OpIncr(&i) {
		market := *(markets).At(i)
		_ = market
		id := this.SafeString2(market, MkString("name"), MkString("symbol"))
		_ = id
		baseId := this.SafeString(market, MkString("base_currency"))
		_ = baseId
		quoteId := this.SafeString(market, MkString("quote_currency"))
		_ = quoteId
		base := this.SafeCurrencyCode(baseId)
		_ = base
		quote := this.SafeCurrencyCode(quoteId)
		_ = quote
		linear := OpHasMember(quote, linearQuoteCurrencies)
		_ = linear
		inverse := OpNot(linear)
		_ = inverse
		symbol := OpAdd(base, OpAdd(MkString("/"), quote))
		_ = symbol
		baseQuote := OpAdd(base, quote)
		_ = baseQuote
		type_ := MkString("swap")
		_ = type_
		if IsTrue(OpNotEq2(baseQuote, id)) {
			symbol = OpCopy(id)
			type_ = MkString("futures")
		}
		lotSizeFilter := this.SafeValue(market, MkString("lot_size_filter"), MkMap(&VarMap{}))
		_ = lotSizeFilter
		priceFilter := this.SafeValue(market, MkString("price_filter"), MkMap(&VarMap{}))
		_ = priceFilter
		precision := MkMap(&VarMap{
			"amount": this.SafeNumber(lotSizeFilter, MkString("qty_step")),
			"price":  this.SafeNumber(priceFilter, MkString("tick_size")),
		})
		_ = precision
		status := this.SafeString(market, MkString("status"))
		_ = status
		active := OpCopy(MkUndefined())
		_ = active
		if IsTrue(OpNotEq2(status, MkUndefined())) {
			active = OpEq2(status, MkString("Trading"))
		}
		spot := OpEq2(type_, MkString("spot"))
		_ = spot
		swap := OpEq2(type_, MkString("swap"))
		_ = swap
		futures := OpEq2(type_, MkString("futures"))
		_ = futures
		option := OpEq2(type_, MkString("option"))
		_ = option
		result.Push(MkMap(&VarMap{
			"id":        id,
			"symbol":    symbol,
			"base":      base,
			"quote":     quote,
			"active":    active,
			"precision": precision,
			"taker":     this.SafeNumber(market, MkString("taker_fee")),
			"maker":     this.SafeNumber(market, MkString("maker_fee")),
			"type":      type_,
			"spot":      spot,
			"swap":      swap,
			"futures":   futures,
			"option":    option,
			"linear":    linear,
			"inverse":   inverse,
			"limits": MkMap(&VarMap{
				"amount": MkMap(&VarMap{
					"min": this.SafeNumber(lotSizeFilter, MkString("min_trading_qty")),
					"max": this.SafeNumber(lotSizeFilter, MkString("max_trading_qty")),
				}),
				"price": MkMap(&VarMap{
					"min": this.SafeNumber(priceFilter, MkString("min_price")),
					"max": this.SafeNumber(priceFilter, MkString("max_price")),
				}),
				"cost": MkMap(&VarMap{
					"min": MkUndefined(),
					"max": MkUndefined(),
				}),
			}),
			"info": market,
		}))
	}
	return result
}

func (this *Bybit) ParseTicker(goArgs ...*Variant) *Variant {
	ticker := GoGetArg(goArgs, 0, MkUndefined())
	_ = ticker
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := OpCopy(MkUndefined())
	_ = timestamp
	marketId := this.SafeString(ticker, MkString("symbol"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market)
	_ = symbol
	last := this.SafeNumber(ticker, MkString("last_price"))
	_ = last
	open := this.SafeNumber(ticker, MkString("prev_price_24h"))
	_ = open
	percentage := this.SafeNumber(ticker, MkString("price_24h_pcnt"))
	_ = percentage
	if IsTrue(OpNotEq2(percentage, MkUndefined())) {
		OpMultiAssign(&percentage, MkInteger(100))
	}
	change := OpCopy(MkUndefined())
	_ = change
	average := OpCopy(MkUndefined())
	_ = average
	if IsTrue(OpAnd(OpNotEq2(last, MkUndefined()), OpNotEq2(open, MkUndefined()))) {
		change = OpSub(last, open)
		average = OpDiv(this.Sum(open, last), MkInteger(2))
	}
	baseVolume := this.SafeNumber(ticker, MkString("turnover_24h"))
	_ = baseVolume
	quoteVolume := this.SafeNumber(ticker, MkString("volume_24h"))
	_ = quoteVolume
	vwap := this.Vwap(baseVolume, quoteVolume)
	_ = vwap
	return MkMap(&VarMap{
		"symbol":        symbol,
		"timestamp":     timestamp,
		"datetime":      this.Iso8601(timestamp),
		"high":          this.SafeNumber(ticker, MkString("high_price_24h")),
		"low":           this.SafeNumber(ticker, MkString("low_price_24h")),
		"bid":           this.SafeNumber(ticker, MkString("bid_price")),
		"bidVolume":     MkUndefined(),
		"ask":           this.SafeNumber(ticker, MkString("ask_price")),
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

func (this *Bybit) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("v2PublicGetTickers"), this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"), MkArray(&VarArray{}))
	_ = result
	first := this.SafeValue(result, MkInteger(0))
	_ = first
	timestamp := this.SafeTimestamp(response, MkString("time_now"))
	_ = timestamp
	ticker := this.ParseTicker(first, market)
	_ = ticker
	*(ticker).At(MkString("timestamp")) = OpCopy(timestamp)
	*(ticker).At(MkString("datetime")) = this.Iso8601(timestamp)
	return ticker
}

func (this *Bybit) FetchTickers(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("v2PublicGetTickers"), params)
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

func (this *Bybit) ParseOHLCV(goArgs ...*Variant) *Variant {
	ohlcv := GoGetArg(goArgs, 0, MkUndefined())
	_ = ohlcv
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	return MkArray(&VarArray{
		this.SafeTimestamp2(ohlcv, MkString("open_time"), MkString("start_at")),
		this.SafeNumber(ohlcv, MkString("open")),
		this.SafeNumber(ohlcv, MkString("high")),
		this.SafeNumber(ohlcv, MkString("low")),
		this.SafeNumber(ohlcv, MkString("close")),
		this.SafeNumber2(ohlcv, MkString("turnover"), MkString("volume")),
	})
}

func (this *Bybit) FetchOHLCV(goArgs ...*Variant) *Variant {
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
		"symbol":   *(market).At(MkString("id")),
		"interval": *(*this.At(MkString("timeframes"))).At(timeframe),
	})
	_ = request
	duration := this.ParseTimeframe(timeframe)
	_ = duration
	now := this.Seconds()
	_ = now
	if IsTrue(OpEq2(since, MkUndefined())) {
		if IsTrue(OpEq2(limit, MkUndefined())) {
			panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchOHLCV() requires a since argument or a limit argument"))))
		} else {
			*(request).At(MkString("from")) = OpSub(now, OpMulti(limit, duration))
		}
	} else {
		*(request).At(MkString("from")) = parseInt(OpDiv(since, MkInteger(1000)))
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	method := OpTrinary(*(market).At(MkString("linear")), MkString("publicLinearGetKline"), MkString("v2PublicGetKlineList"))
	_ = method
	response := this.Call(method, this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"), MkMap(&VarMap{}))
	_ = result
	return this.ParseOHLCVs(result, market, timeframe, since, limit)
}

func (this *Bybit) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	id := this.SafeString2(trade, MkString("id"), MkString("exec_id"))
	_ = id
	marketId := this.SafeString(trade, MkString("symbol"))
	_ = marketId
	market = this.SafeMarket(marketId, market)
	symbol := *(market).At(MkString("symbol"))
	_ = symbol
	amountString := this.SafeString2(trade, MkString("qty"), MkString("exec_qty"))
	_ = amountString
	priceString := this.SafeString2(trade, MkString("exec_price"), MkString("price"))
	_ = priceString
	cost := this.SafeNumber(trade, MkString("exec_value"))
	_ = cost
	amount := this.ParseNumber(amountString)
	_ = amount
	price := this.ParseNumber(priceString)
	_ = price
	if IsTrue(OpEq2(cost, MkUndefined())) {
		cost = this.ParseNumber(Precise.StringMul(priceString, amountString))
	}
	timestamp := this.Parse8601(this.SafeString(trade, MkString("time")))
	_ = timestamp
	if IsTrue(OpEq2(timestamp, MkUndefined())) {
		timestamp = this.SafeInteger(trade, MkString("trade_time_ms"))
	}
	side := this.SafeStringLower(trade, MkString("side"))
	_ = side
	lastLiquidityInd := this.SafeString(trade, MkString("last_liquidity_ind"))
	_ = lastLiquidityInd
	takerOrMaker := OpTrinary(OpEq2(lastLiquidityInd, MkString("AddedLiquidity")), MkString("maker"), MkString("taker"))
	_ = takerOrMaker
	feeCost := this.SafeNumber(trade, MkString("exec_fee"))
	_ = feeCost
	fee := OpCopy(MkUndefined())
	_ = fee
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		feeCurrencyCode := OpTrinary(*(market).At(MkString("inverse")), *(market).At(MkString("base")), *(market).At(MkString("quote")))
		_ = feeCurrencyCode
		fee = MkMap(&VarMap{
			"cost":     feeCost,
			"currency": feeCurrencyCode,
			"rate":     this.SafeNumber(trade, MkString("fee_rate")),
		})
	}
	return MkMap(&VarMap{
		"id":           id,
		"info":         trade,
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       symbol,
		"order":        this.SafeString(trade, MkString("order_id")),
		"type":         this.SafeStringLower(trade, MkString("order_type")),
		"side":         side,
		"takerOrMaker": takerOrMaker,
		"price":        price,
		"amount":       amount,
		"cost":         cost,
		"fee":          fee,
	})
}

func (this *Bybit) FetchTrades(goArgs ...*Variant) *Variant {
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
		*(request).At(MkString("count")) = OpCopy(limit)
	}
	method := OpTrinary(*(market).At(MkString("linear")), MkString("publicLinearGetRecentTradingRecords"), MkString("v2PublicGetTradingRecords"))
	_ = method
	response := this.Call(method, this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"), MkMap(&VarMap{}))
	_ = result
	return this.ParseTrades(result, market, since, limit)
}

func (this *Bybit) ParseOrderBook(goArgs ...*Variant) *Variant {
	orderbook := GoGetArg(goArgs, 0, MkUndefined())
	_ = orderbook
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	timestamp := GoGetArg(goArgs, 2, MkUndefined())
	_ = timestamp
	bidsKey := GoGetArg(goArgs, 3, MkString("Buy"))
	_ = bidsKey
	asksKey := GoGetArg(goArgs, 4, MkString("Sell"))
	_ = asksKey
	priceKey := GoGetArg(goArgs, 5, MkString("price"))
	_ = priceKey
	amountKey := GoGetArg(goArgs, 6, MkString("size"))
	_ = amountKey
	bids := MkArray(&VarArray{})
	_ = bids
	asks := MkArray(&VarArray{})
	_ = asks
	for i := MkInteger(0); IsTrue(OpLw(i, orderbook.Length)); OpIncr(&i) {
		bidask := *(orderbook).At(i)
		_ = bidask
		side := this.SafeString(bidask, MkString("side"))
		_ = side
		if IsTrue(OpEq2(side, MkString("Buy"))) {
			bids.Push(this.ParseBidAsk(bidask, priceKey, amountKey))
		} else {
			if IsTrue(OpEq2(side, MkString("Sell"))) {
				asks.Push(this.ParseBidAsk(bidask, priceKey, amountKey))
			} else {
				panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" parseOrderBook encountered an unrecognized bidask format: "), this.Json(bidask)))))
			}
		}
	}
	return MkMap(&VarMap{
		"symbol":    symbol,
		"bids":      this.SortBy(bids, MkInteger(0), MkBool(true)),
		"asks":      this.SortBy(asks, MkInteger(0)),
		"timestamp": timestamp,
		"datetime":  this.Iso8601(timestamp),
		"nonce":     MkUndefined(),
	})
}

func (this *Bybit) FetchOrderBook(goArgs ...*Variant) *Variant {
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
	response := this.Call(MkString("v2PublicGetOrderBookL2"), this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"), MkArray(&VarArray{}))
	_ = result
	timestamp := this.SafeTimestamp(response, MkString("time_now"))
	_ = timestamp
	return this.ParseOrderBook(result, symbol, timestamp, MkString("Buy"), MkString("Sell"), MkString("price"), MkString("size"))
}

func (this *Bybit) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{})
	_ = request
	coin := this.SafeString(params, MkString("coin"))
	_ = coin
	code := this.SafeString(params, MkString("code"))
	_ = code
	if IsTrue(OpNotEq2(coin, MkUndefined())) {
		*(request).At(MkString("coin")) = OpCopy(coin)
	} else {
		if IsTrue(OpNotEq2(code, MkUndefined())) {
			currency := this.Currency(code)
			_ = currency
			*(request).At(MkString("coin")) = *(currency).At(MkString("id"))
		}
	}
	response := this.Call(MkString("v2PrivateGetWalletBalance"), this.Extend(request, params))
	_ = response
	result := MkMap(&VarMap{"info": response})
	_ = result
	balances := this.SafeValue(response, MkString("result"), MkMap(&VarMap{}))
	_ = balances
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
		*(account).At(MkString("free")) = this.SafeString(balance, MkString("available_balance"))
		*(account).At(MkString("used")) = this.SafeString(balance, MkString("used_margin"))
		*(account).At(MkString("total")) = this.SafeString(balance, MkString("equity"))
		*(result).At(code) = OpCopy(account)
	}
	return this.ParseBalance(result)
}

func (this *Bybit) ParseOrderStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"Created":         MkString("open"),
		"Rejected":        MkString("rejected"),
		"New":             MkString("open"),
		"PartiallyFilled": MkString("open"),
		"Filled":          MkString("closed"),
		"Cancelled":       MkString("canceled"),
		"PendingCancel":   MkString("canceling"),
		"Active":          MkString("open"),
		"Untriggered":     MkString("open"),
		"Triggered":       MkString("closed"),
		"Deactivated":     MkString("canceled"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Bybit) ParseTimeInForce(goArgs ...*Variant) *Variant {
	timeInForce := GoGetArg(goArgs, 0, MkUndefined())
	_ = timeInForce
	timeInForces := MkMap(&VarMap{
		"GoodTillCancel":    MkString("GTC"),
		"ImmediateOrCancel": MkString("IOC"),
		"FillOrKill":        MkString("FOK"),
		"PostOnly":          MkString("PO"),
	})
	_ = timeInForces
	return this.SafeString(timeInForces, timeInForce, timeInForce)
}

func (this *Bybit) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	marketId := this.SafeString(order, MkString("symbol"))
	_ = marketId
	market = this.SafeMarket(marketId, market)
	symbol := *(market).At(MkString("symbol"))
	_ = symbol
	feeCurrency := OpCopy(MkUndefined())
	_ = feeCurrency
	timestamp := this.Parse8601(this.SafeString(order, MkString("created_at")))
	_ = timestamp
	id := this.SafeString2(order, MkString("order_id"), MkString("stop_order_id"))
	_ = id
	type_ := this.SafeStringLower(order, MkString("order_type"))
	_ = type_
	price := this.SafeNumber(order, MkString("price"))
	_ = price
	if IsTrue(OpEq2(price, MkNumber(0.0))) {
		price = OpCopy(MkUndefined())
	}
	average := this.SafeNumber(order, MkString("average_price"))
	_ = average
	amount := this.SafeNumber(order, MkString("qty"))
	_ = amount
	cost := this.SafeNumber(order, MkString("cum_exec_value"))
	_ = cost
	filled := this.SafeNumber(order, MkString("cum_exec_qty"))
	_ = filled
	remaining := this.SafeNumber(order, MkString("leaves_qty"))
	_ = remaining
	marketTypes := this.SafeValue(*this.At(MkString("options")), MkString("marketTypes"), MkMap(&VarMap{}))
	_ = marketTypes
	marketType := this.SafeString(marketTypes, symbol)
	_ = marketType
	if IsTrue(OpNotEq2(market, MkUndefined())) {
		if IsTrue(OpEq2(marketType, MkString("linear"))) {
			feeCurrency = *(market).At(MkString("quote"))
		} else {
			feeCurrency = *(market).At(MkString("base"))
		}
	}
	lastTradeTimestamp := this.SafeTimestamp(order, MkString("last_exec_time"))
	_ = lastTradeTimestamp
	if IsTrue(OpEq2(lastTradeTimestamp, MkInteger(0))) {
		lastTradeTimestamp = OpCopy(MkUndefined())
	}
	status := this.ParseOrderStatus(this.SafeString2(order, MkString("order_status"), MkString("stop_order_status")))
	_ = status
	side := this.SafeStringLower(order, MkString("side"))
	_ = side
	feeCost := this.SafeNumber(order, MkString("cum_exec_fee"))
	_ = feeCost
	fee := OpCopy(MkUndefined())
	_ = fee
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		feeCost = Math.Abs(feeCost)
		fee = MkMap(&VarMap{
			"cost":     feeCost,
			"currency": feeCurrency,
		})
	}
	clientOrderId := this.SafeString(order, MkString("order_link_id"))
	_ = clientOrderId
	if IsTrue(OpAnd(OpNotEq2(clientOrderId, MkUndefined()), OpLw(clientOrderId.Length, MkInteger(1)))) {
		clientOrderId = OpCopy(MkUndefined())
	}
	timeInForce := this.ParseTimeInForce(this.SafeString(order, MkString("time_in_force")))
	_ = timeInForce
	stopPrice := this.SafeNumber2(order, MkString("trigger_price"), MkString("stop_px"))
	_ = stopPrice
	postOnly := OpEq2(timeInForce, MkString("PO"))
	_ = postOnly
	return this.SafeOrder(MkMap(&VarMap{
		"info":               order,
		"id":                 id,
		"clientOrderId":      clientOrderId,
		"timestamp":          timestamp,
		"datetime":           this.Iso8601(timestamp),
		"lastTradeTimestamp": lastTradeTimestamp,
		"symbol":             symbol,
		"type":               type_,
		"timeInForce":        timeInForce,
		"postOnly":           postOnly,
		"side":               side,
		"price":              price,
		"stopPrice":          stopPrice,
		"amount":             amount,
		"cost":               cost,
		"average":            average,
		"filled":             filled,
		"remaining":          remaining,
		"status":             status,
		"fee":                fee,
		"trades":             MkUndefined(),
	}))
}

func (this *Bybit) FetchOrder(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	method := OpCopy(MkUndefined())
	_ = method
	if IsTrue(*(market).At(MkString("swap"))) {
		if IsTrue(*(market).At(MkString("linear"))) {
			method = MkString("privateLinearGetOrderSearch")
		} else {
			if IsTrue(*(market).At(MkString("inverse"))) {
				method = MkString("v2PrivateGetOrder")
			}
		}
	} else {
		if IsTrue(*(market).At(MkString("futures"))) {
			method = MkString("futuresPrivateGetOrder")
		}
	}
	stopOrderId := this.SafeString(params, MkString("stop_order_id"))
	_ = stopOrderId
	if IsTrue(OpEq2(stopOrderId, MkUndefined())) {
		orderLinkId := this.SafeString(params, MkString("order_link_id"))
		_ = orderLinkId
		if IsTrue(OpEq2(orderLinkId, MkUndefined())) {
			*(request).At(MkString("order_id")) = OpCopy(id)
		}
	} else {
		if IsTrue(*(market).At(MkString("swap"))) {
			if IsTrue(*(market).At(MkString("linear"))) {
				method = MkString("privateLinearGetStopOrderSearch")
			} else {
				if IsTrue(*(market).At(MkString("inverse"))) {
					method = MkString("v2PrivateGetStopOrder")
				}
			}
		} else {
			if IsTrue(*(market).At(MkString("futures"))) {
				method = MkString("futuresPrivateGetStopOrder")
			}
		}
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"))
	_ = result
	return this.ParseOrder(result, market)
}

func (this *Bybit) CreateOrder(goArgs ...*Variant) *Variant {
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
	qty := this.AmountToPrecision(symbol, amount)
	_ = qty
	if IsTrue(*(market).At(MkString("inverse"))) {
		qty = parseInt(qty)
	} else {
		qty = parseFloat(qty)
	}
	request := MkMap(&VarMap{
		"side":          this.Capitalize(side),
		"symbol":        *(market).At(MkString("id")),
		"order_type":    this.Capitalize(type_),
		"qty":           qty,
		"time_in_force": MkString("GoodTillCancel"),
	})
	_ = request
	priceIsRequired := OpCopy(MkBool(false))
	_ = priceIsRequired
	if IsTrue(OpEq2(type_, MkString("limit"))) {
		priceIsRequired = OpCopy(MkBool(true))
	}
	if IsTrue(priceIsRequired) {
		if IsTrue(OpNotEq2(price, MkUndefined())) {
			*(request).At(MkString("price")) = parseFloat(this.PriceToPrecision(symbol, price))
		} else {
			panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" createOrder() requires a price argument for a "), OpAdd(type_, MkString(" order"))))))
		}
	}
	clientOrderId := this.SafeString2(params, MkString("order_link_id"), MkString("clientOrderId"))
	_ = clientOrderId
	if IsTrue(OpNotEq2(clientOrderId, MkUndefined())) {
		*(request).At(MkString("order_link_id")) = OpCopy(clientOrderId)
		params = this.Omit(params, MkArray(&VarArray{
			MkString("order_link_id"),
			MkString("clientOrderId"),
		}))
	}
	stopPx := this.SafeValue2(params, MkString("stop_px"), MkString("stopPrice"))
	_ = stopPx
	basePrice := this.SafeValue(params, MkString("base_price"))
	_ = basePrice
	method := OpCopy(MkUndefined())
	_ = method
	if IsTrue(*(market).At(MkString("swap"))) {
		if IsTrue(*(market).At(MkString("linear"))) {
			method = MkString("privateLinearPostOrderCreate")
			*(request).At(MkString("reduce_only")) = OpCopy(MkBool(false))
			*(request).At(MkString("close_on_trigger")) = OpCopy(MkBool(false))
		} else {
			if IsTrue(*(market).At(MkString("inverse"))) {
				method = MkString("v2PrivatePostOrderCreate")
			}
		}
	} else {
		if IsTrue(*(market).At(MkString("futures"))) {
			method = MkString("futuresPrivatePostOrderCreate")
		}
	}
	if IsTrue(OpNotEq2(stopPx, MkUndefined())) {
		if IsTrue(OpEq2(basePrice, MkUndefined())) {
			panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" createOrder() requires both the stop_px and base_price params for a conditional "), OpAdd(type_, MkString(" order"))))))
		} else {
			if IsTrue(*(market).At(MkString("swap"))) {
				if IsTrue(*(market).At(MkString("linear"))) {
					method = MkString("privateLinearPostStopOrderCreate")
				} else {
					if IsTrue(*(market).At(MkString("inverse"))) {
						method = MkString("v2PrivatePostStopOrderCreate")
					}
				}
			} else {
				if IsTrue(*(market).At(MkString("futures"))) {
					method = MkString("futuresPrivatePostStopOrderCreate")
				}
			}
			*(request).At(MkString("stop_px")) = parseFloat(this.PriceToPrecision(symbol, stopPx))
			*(request).At(MkString("base_price")) = parseFloat(this.PriceToPrecision(symbol, basePrice))
			params = this.Omit(params, MkArray(&VarArray{
				MkString("stop_px"),
				MkString("stopPrice"),
				MkString("base_price"),
			}))
		}
	} else {
		if IsTrue(OpNotEq2(basePrice, MkUndefined())) {
			panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" createOrder() requires both the stop_px and base_price params for a conditional "), OpAdd(type_, MkString(" order"))))))
		}
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"))
	_ = result
	return this.ParseOrder(result, market)
}

func (this *Bybit) EditOrder(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" editOrder() requires an symbol argument"))))
	}
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	method := OpCopy(MkUndefined())
	_ = method
	if IsTrue(*(market).At(MkString("swap"))) {
		if IsTrue(*(market).At(MkString("linear"))) {
			method = MkString("privateLinearPostOrderReplace")
		} else {
			if IsTrue(*(market).At(MkString("inverse"))) {
				method = MkString("v2PrivatePostOrderReplace")
			}
		}
	} else {
		if IsTrue(*(market).At(MkString("futures"))) {
			method = MkString("futuresPrivatePostOrderReplace")
		}
	}
	stopOrderId := this.SafeString(params, MkString("stop_order_id"))
	_ = stopOrderId
	if IsTrue(OpNotEq2(stopOrderId, MkUndefined())) {
		if IsTrue(*(market).At(MkString("swap"))) {
			if IsTrue(*(market).At(MkString("linear"))) {
				method = MkString("privateLinearPostStopOrderReplace")
			} else {
				if IsTrue(*(market).At(MkString("inverse"))) {
					method = MkString("v2PrivatePostStopOrderReplace")
				}
			}
		} else {
			if IsTrue(*(market).At(MkString("futures"))) {
				method = MkString("futuresPrivatePostStopOrderReplace")
			}
		}
		*(request).At(MkString("stop_order_id")) = OpCopy(stopOrderId)
		params = this.Omit(params, MkArray(&VarArray{
			MkString("stop_order_id"),
		}))
	} else {
		*(request).At(MkString("order_id")) = OpCopy(id)
	}
	if IsTrue(OpNotEq2(amount, MkUndefined())) {
		qty := this.AmountToPrecision(symbol, amount)
		_ = qty
		if IsTrue(*(market).At(MkString("inverse"))) {
			qty = parseInt(qty)
		} else {
			qty = parseFloat(qty)
		}
		*(request).At(MkString("p_r_qty")) = OpCopy(qty)
	}
	if IsTrue(OpNotEq2(price, MkUndefined())) {
		*(request).At(MkString("p_r_price")) = parseFloat(this.PriceToPrecision(symbol, price))
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"), MkMap(&VarMap{}))
	_ = result
	return MkMap(&VarMap{
		"info":          response,
		"id":            this.SafeString2(result, MkString("order_id"), MkString("stop_order_id")),
		"order_id":      this.SafeString(result, MkString("order_id")),
		"stop_order_id": this.SafeString(result, MkString("stop_order_id")),
	})
}

func (this *Bybit) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" cancelOrder() requires a symbol argument"))))
	}
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	method := OpCopy(MkUndefined())
	_ = method
	if IsTrue(*(market).At(MkString("swap"))) {
		if IsTrue(*(market).At(MkString("linear"))) {
			method = MkString("privateLinearPostOrderCancel")
		} else {
			if IsTrue(*(market).At(MkString("inverse"))) {
				method = MkString("v2PrivatePostOrderCancel")
			}
		}
	} else {
		if IsTrue(*(market).At(MkString("futures"))) {
			method = MkString("futuresPrivatePostOrderCancel")
		}
	}
	stopOrderId := this.SafeString(params, MkString("stop_order_id"))
	_ = stopOrderId
	if IsTrue(OpEq2(stopOrderId, MkUndefined())) {
		orderLinkId := this.SafeString(params, MkString("order_link_id"))
		_ = orderLinkId
		if IsTrue(OpEq2(orderLinkId, MkUndefined())) {
			*(request).At(MkString("order_id")) = OpCopy(id)
		}
	} else {
		if IsTrue(*(market).At(MkString("swap"))) {
			if IsTrue(*(market).At(MkString("linear"))) {
				method = MkString("privateLinearPostStopOrderCancel")
			} else {
				if IsTrue(*(market).At(MkString("inverse"))) {
					method = MkString("v2PrivatePostStopOrderCancel")
				}
			}
		} else {
			if IsTrue(*(market).At(MkString("futures"))) {
				method = MkString("futuresPrivatePostStopOrderCancel")
			}
		}
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"), MkMap(&VarMap{}))
	_ = result
	return this.ParseOrder(result, market)
}

func (this *Bybit) CancelAllOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" cancelAllOrders() requires a symbol argument"))))
	}
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	options := this.SafeValue(*this.At(MkString("options")), MkString("cancelAllOrders"), MkMap(&VarMap{}))
	_ = options
	defaultMethod := OpCopy(MkUndefined())
	_ = defaultMethod
	if IsTrue(*(market).At(MkString("swap"))) {
		if IsTrue(*(market).At(MkString("linear"))) {
			defaultMethod = MkString("privateLinearPostOrderCancelAll")
		} else {
			if IsTrue(*(market).At(MkString("inverse"))) {
				defaultMethod = MkString("v2PrivatePostOrderCancelAll")
			}
		}
	} else {
		if IsTrue(*(market).At(MkString("futures"))) {
			defaultMethod = MkString("futuresPrivatePostOrderCancelAll")
		}
	}
	method := this.SafeString(options, MkString("method"), defaultMethod)
	_ = method
	response := this.Call(method, this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"), MkArray(&VarArray{}))
	_ = result
	return this.ParseOrders(result, market)
}

func (this *Bybit) FetchOrders(goArgs ...*Variant) *Variant {
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
	options := this.SafeValue(*this.At(MkString("options")), MkString("fetchOrders"), MkMap(&VarMap{}))
	_ = options
	defaultType := this.SafeString(*this.At(MkString("options")), MkString("defaultType"), MkString("linear"))
	_ = defaultType
	marketTypes := this.SafeValue(*this.At(MkString("options")), MkString("marketTypes"), MkMap(&VarMap{}))
	_ = marketTypes
	marketType := this.SafeString(marketTypes, symbol, defaultType)
	_ = marketType
	defaultMethod := OpCopy(MkUndefined())
	_ = defaultMethod
	marketDefined := OpNotEq2(market, MkUndefined())
	_ = marketDefined
	linear := OpOr(OpAnd(marketDefined, *(market).At(MkString("linear"))), OpEq2(marketType, MkString("linear")))
	_ = linear
	inverse := OpOr(OpAnd(marketDefined, OpAnd(*(market).At(MkString("swap")), *(market).At(MkString("inverse")))), OpEq2(marketType, MkString("inverse")))
	_ = inverse
	futures := OpOr(OpAnd(marketDefined, *(market).At(MkString("futures"))), OpEq2(marketType, MkString("futures")))
	_ = futures
	if IsTrue(linear) {
		defaultMethod = MkString("privateLinearGetOrderList")
	} else {
		if IsTrue(inverse) {
			defaultMethod = MkString("v2PrivateGetOrderList")
		} else {
			if IsTrue(futures) {
				defaultMethod = MkString("futuresPrivateGetOrderList")
			}
		}
	}
	query := OpCopy(params)
	_ = query
	if IsTrue(OpOr(OpHasMember(MkString("stop_order_id"), params), OpHasMember(MkString("stop_order_status"), params))) {
		stopOrderStatus := this.SafeValue(params, MkString("stop_order_status"))
		_ = stopOrderStatus
		if IsTrue(OpNotEq2(stopOrderStatus, MkUndefined())) {
			if IsTrue(GoIsArray(stopOrderStatus)) {
				stopOrderStatus = stopOrderStatus.Join(MkString(","))
			}
			*(request).At(MkString("stop_order_status")) = OpCopy(stopOrderStatus)
			query = this.Omit(params, MkString("stop_order_status"))
		}
		if IsTrue(linear) {
			defaultMethod = MkString("privateLinearGetStopOrderList")
		} else {
			if IsTrue(inverse) {
				defaultMethod = MkString("v2PrivateGetStopOrderList")
			} else {
				if IsTrue(futures) {
					defaultMethod = MkString("futuresPrivateGetStopOrderList")
				}
			}
		}
	}
	method := this.SafeString(options, MkString("method"), defaultMethod)
	_ = method
	response := this.Call(method, this.Extend(request, query))
	_ = response
	result := this.SafeValue(response, MkString("result"), MkMap(&VarMap{}))
	_ = result
	data := this.SafeValue(result, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseOrders(data, market, since, limit)
}

func (this *Bybit) FetchClosedOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	defaultStatuses := MkArray(&VarArray{
		MkString("Rejected"),
		MkString("Filled"),
		MkString("Cancelled"),
	})
	_ = defaultStatuses
	options := this.SafeValue(*this.At(MkString("options")), MkString("fetchClosedOrders"), MkMap(&VarMap{}))
	_ = options
	status := this.SafeValue(options, MkString("order_status"), defaultStatuses)
	_ = status
	if IsTrue(GoIsArray(status)) {
		status = status.Join(MkString(","))
	}
	request := MkMap(&VarMap{})
	_ = request
	stopOrderStatus := this.SafeValue(params, MkString("stop_order_status"))
	_ = stopOrderStatus
	if IsTrue(OpEq2(stopOrderStatus, MkUndefined())) {
		*(request).At(MkString("order_status")) = OpCopy(status)
	} else {
		*(request).At(MkString("stop_order_status")) = OpCopy(stopOrderStatus)
	}
	return this.FetchOrders(symbol, since, limit, this.Extend(request, params))
}

func (this *Bybit) FetchOpenOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	defaultStatuses := MkArray(&VarArray{
		MkString("Created"),
		MkString("New"),
		MkString("PartiallyFilled"),
		MkString("PendingCancel"),
	})
	_ = defaultStatuses
	options := this.SafeValue(*this.At(MkString("options")), MkString("fetchOpenOrders"), MkMap(&VarMap{}))
	_ = options
	status := this.SafeValue(options, MkString("order_status"), defaultStatuses)
	_ = status
	if IsTrue(GoIsArray(status)) {
		status = status.Join(MkString(","))
	}
	request := MkMap(&VarMap{})
	_ = request
	stopOrderStatus := this.SafeValue(params, MkString("stop_order_status"))
	_ = stopOrderStatus
	if IsTrue(OpEq2(stopOrderStatus, MkUndefined())) {
		*(request).At(MkString("order_status")) = OpCopy(status)
	} else {
		*(request).At(MkString("stop_order_status")) = OpCopy(stopOrderStatus)
	}
	return this.FetchOrders(symbol, since, limit, this.Extend(request, params))
}

func (this *Bybit) FetchOrderTrades(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"order_id": id})
	_ = request
	return this.FetchMyTrades(symbol, since, limit, this.Extend(request, params))
}

func (this *Bybit) FetchMyTrades(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		orderId := this.SafeString(params, MkString("order_id"))
		_ = orderId
		if IsTrue(OpEq2(orderId, MkUndefined())) {
			panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchMyTrades() requires a symbol argument or an order_id param"))))
		} else {
			*(request).At(MkString("order_id")) = OpCopy(orderId)
			params = this.Omit(params, MkString("order_id"))
		}
	} else {
		market = this.Market(symbol)
		*(request).At(MkString("symbol")) = *(market).At(MkString("id"))
	}
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("start_time")) = OpCopy(since)
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	defaultType := this.SafeString(*this.At(MkString("options")), MkString("defaultType"), MkString("linear"))
	_ = defaultType
	marketTypes := this.SafeValue(*this.At(MkString("options")), MkString("marketTypes"), MkMap(&VarMap{}))
	_ = marketTypes
	marketType := this.SafeString(marketTypes, symbol, defaultType)
	_ = marketType
	marketDefined := OpNotEq2(market, MkUndefined())
	_ = marketDefined
	linear := OpOr(OpAnd(marketDefined, *(market).At(MkString("linear"))), OpEq2(marketType, MkString("linear")))
	_ = linear
	inverse := OpOr(OpAnd(marketDefined, OpAnd(*(market).At(MkString("swap")), *(market).At(MkString("inverse")))), OpEq2(marketType, MkString("inverse")))
	_ = inverse
	futures := OpOr(OpAnd(marketDefined, *(market).At(MkString("futures"))), OpEq2(marketType, MkString("futures")))
	_ = futures
	method := OpCopy(MkUndefined())
	_ = method
	if IsTrue(linear) {
		method = MkString("privateLinearGetTradeExecutionList")
	} else {
		if IsTrue(inverse) {
			method = MkString("v2PrivateGetExecutionList")
		} else {
			if IsTrue(futures) {
				method = MkString("futuresPrivateGetExecutionList")
			}
		}
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"), MkMap(&VarMap{}))
	_ = result
	trades := this.SafeValue2(result, MkString("trade_list"), MkString("data"), MkArray(&VarArray{}))
	_ = trades
	return this.ParseTrades(trades, market, since, limit)
}

func (this *Bybit) FetchDeposits(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"wallet_fund_type": MkString("Deposit")})
	_ = request
	currency := OpCopy(MkUndefined())
	_ = currency
	if IsTrue(OpNotEq2(code, MkUndefined())) {
		currency = this.Currency(code)
		*(request).At(MkString("coin")) = *(currency).At(MkString("id"))
	}
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("start_date")) = this.Ymd(since)
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("v2PrivateGetWalletFundRecords"), this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"), MkMap(&VarMap{}))
	_ = result
	data := this.SafeValue(result, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseTransactions(data, currency, since, limit, MkMap(&VarMap{"type": MkString("deposit")}))
}

func (this *Bybit) FetchWithdrawals(goArgs ...*Variant) *Variant {
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
		*(request).At(MkString("coin")) = *(currency).At(MkString("id"))
	}
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("start_date")) = this.Ymd(since)
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("v2PrivateGetWalletWithdrawList"), this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"), MkMap(&VarMap{}))
	_ = result
	data := this.SafeValue(result, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseTransactions(data, currency, since, limit, MkMap(&VarMap{"type": MkString("withdrawal")}))
}

func (this *Bybit) ParseTransactionStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"ToBeConfirmed": MkString("pending"),
		"UnderReview":   MkString("pending"),
		"Pending":       MkString("pending"),
		"Success":       MkString("ok"),
		"CancelByUser":  MkString("canceled"),
		"Reject":        MkString("rejected"),
		"Expire":        MkString("expired"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Bybit) ParseTransaction(goArgs ...*Variant) *Variant {
	transaction := GoGetArg(goArgs, 0, MkUndefined())
	_ = transaction
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	currencyId := this.SafeString(transaction, MkString("coin"))
	_ = currencyId
	code := this.SafeCurrencyCode(currencyId, currency)
	_ = code
	timestamp := this.Parse8601(this.SafeString2(transaction, MkString("submited_at"), MkString("exec_time")))
	_ = timestamp
	updated := this.Parse8601(this.SafeString(transaction, MkString("updated_at")))
	_ = updated
	status := this.ParseTransactionStatus(this.SafeString(transaction, MkString("status")))
	_ = status
	address := this.SafeString(transaction, MkString("address"))
	_ = address
	feeCost := this.SafeNumber(transaction, MkString("fee"))
	_ = feeCost
	type_ := this.SafeStringLower(transaction, MkString("type"))
	_ = type_
	fee := OpCopy(MkUndefined())
	_ = fee
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		fee = MkMap(&VarMap{
			"cost":     feeCost,
			"currency": code,
		})
	}
	return MkMap(&VarMap{
		"info":        transaction,
		"id":          this.SafeString(transaction, MkString("id")),
		"txid":        this.SafeString(transaction, MkString("tx_id")),
		"timestamp":   timestamp,
		"datetime":    this.Iso8601(timestamp),
		"address":     address,
		"addressTo":   MkUndefined(),
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

func (this *Bybit) FetchLedger(goArgs ...*Variant) *Variant {
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
		*(request).At(MkString("coin")) = *(currency).At(MkString("id"))
	}
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("start_date")) = this.Ymd(since)
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("v2PrivateGetWalletFundRecords"), this.Extend(request, params))
	_ = response
	result := this.SafeValue(response, MkString("result"), MkMap(&VarMap{}))
	_ = result
	data := this.SafeValue(result, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseLedger(data, currency, since, limit)
}

func (this *Bybit) ParseLedgerEntry(goArgs ...*Variant) *Variant {
	item := GoGetArg(goArgs, 0, MkUndefined())
	_ = item
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	currencyId := this.SafeString(item, MkString("coin"))
	_ = currencyId
	code := this.SafeCurrencyCode(currencyId, currency)
	_ = code
	amount := this.SafeNumber(item, MkString("amount"))
	_ = amount
	after := this.SafeNumber(item, MkString("wallet_balance"))
	_ = after
	direction := OpTrinary(OpLw(amount, MkInteger(0)), MkString("out"), MkString("in"))
	_ = direction
	before := OpCopy(MkUndefined())
	_ = before
	if IsTrue(OpAnd(OpNotEq2(after, MkUndefined()), OpNotEq2(amount, MkUndefined()))) {
		difference := OpTrinary(OpEq2(direction, MkString("out")), amount, OpNeg(amount))
		_ = difference
		before = this.Sum(after, difference)
	}
	timestamp := this.Parse8601(this.SafeString(item, MkString("exec_time")))
	_ = timestamp
	type_ := this.ParseLedgerEntryType(this.SafeString(item, MkString("type")))
	_ = type_
	id := this.SafeString(item, MkString("id"))
	_ = id
	referenceId := this.SafeString(item, MkString("tx_id"))
	_ = referenceId
	return MkMap(&VarMap{
		"id":               id,
		"currency":         code,
		"account":          this.SafeString(item, MkString("wallet_id")),
		"referenceAccount": MkUndefined(),
		"referenceId":      referenceId,
		"status":           MkUndefined(),
		"amount":           amount,
		"before":           before,
		"after":            after,
		"fee":              MkUndefined(),
		"direction":        direction,
		"timestamp":        timestamp,
		"datetime":         this.Iso8601(timestamp),
		"type":             type_,
		"info":             item,
	})
}

func (this *Bybit) ParseLedgerEntryType(goArgs ...*Variant) *Variant {
	type_ := GoGetArg(goArgs, 0, MkUndefined())
	_ = type_
	types := MkMap(&VarMap{
		"Deposit":               MkString("transaction"),
		"Withdraw":              MkString("transaction"),
		"RealisedPNL":           MkString("trade"),
		"Commission":            MkString("fee"),
		"Refund":                MkString("cashback"),
		"Prize":                 MkString("prize"),
		"ExchangeOrderWithdraw": MkString("transaction"),
		"ExchangeOrderDeposit":  MkString("transaction"),
	})
	_ = types
	return this.SafeString(types, type_, type_)
}

func (this *Bybit) FetchPositions(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{})
	_ = request
	if IsTrue(GoIsArray(symbols)) {
		length := OpCopy(symbols.Length)
		_ = length
		if IsTrue(OpNotEq2(length, MkInteger(1))) {
			panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchPositions takes exactly one symbol"))))
		}
		*(request).At(MkString("symbol")) = this.MarketId(*(symbols).At(MkInteger(0)))
	}
	defaultType := this.SafeString(*this.At(MkString("options")), MkString("defaultType"), MkString("linear"))
	_ = defaultType
	type_ := this.SafeString(params, MkString("type"), defaultType)
	_ = type_
	params = this.Omit(params, MkString("type"))
	response := OpCopy(MkUndefined())
	_ = response
	if IsTrue(OpEq2(type_, MkString("linear"))) {
		response = this.Call(MkString("privateLinearGetPositionList"), this.Extend(request, params))
	} else {
		if IsTrue(OpEq2(type_, MkString("inverse"))) {
			response = this.Call(MkString("v2PrivateGetPositionList"), this.Extend(request, params))
		} else {
			if IsTrue(OpEq2(type_, MkString("inverseFuture"))) {
				response = this.Call(MkString("futuresPrivateGetPositionList"), this.Extend(request, params))
			}
		}
	}
	return this.SafeValue(response, MkString("result"))
}

func (this *Bybit) Sign(goArgs ...*Variant) *Variant {
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
	type_ := this.SafeString(api, MkInteger(0))
	_ = type_
	section := this.SafeString(api, MkInteger(1))
	_ = section
	url := this.ImplodeHostname(*(*(*this.At(MkString("urls"))).At(MkString("api"))).At(type_))
	_ = url
	request := OpAdd(MkString("/"), OpAdd(type_, OpAdd(MkString("/"), OpAdd(section, OpAdd(MkString("/"), path)))))
	_ = request
	if IsTrue(OpEq2(section, MkString("public"))) {
		if IsTrue(*(GoGetKeys(params)).At(MkString("length"))) {
			OpAddAssign(&request, OpAdd(MkString("?"), this.Rawencode(params)))
		}
	} else {
		if IsTrue(OpEq2(type_, MkString("public"))) {
			if IsTrue(*(GoGetKeys(params)).At(MkString("length"))) {
				OpAddAssign(&request, OpAdd(MkString("?"), this.Rawencode(params)))
			}
		} else {
			this.CheckRequiredCredentials()
			timestamp := this.Nonce()
			_ = timestamp
			query := this.Extend(params, MkMap(&VarMap{
				"api_key":     *this.At(MkString("apiKey")),
				"recv_window": *(*this.At(MkString("options"))).At(MkString("recvWindow")),
				"timestamp":   timestamp,
			}))
			_ = query
			sortedQuery := this.Keysort(query)
			_ = sortedQuery
			auth := this.Rawencode(sortedQuery)
			_ = auth
			signature := this.Hmac(this.Encode(auth), this.Encode(*this.At(MkString("secret"))))
			_ = signature
			if IsTrue(OpEq2(method, MkString("POST"))) {
				body = this.Json(this.Extend(query, MkMap(&VarMap{"sign": signature})))
				headers = MkMap(&VarMap{"Content-Type": MkString("application/json")})
			} else {
				OpAddAssign(&request, OpAdd(MkString("?"), OpAdd(this.Urlencode(sortedQuery), OpAdd(MkString("&sign="), signature))))
			}
		}
	}
	OpAddAssign(&url, request)
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Bybit) HandleErrors(goArgs ...*Variant) *Variant {
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
	errorCode := this.SafeString(response, MkString("ret_code"))
	_ = errorCode
	if IsTrue(OpNotEq2(errorCode, MkString("0"))) {
		feedback := OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))
		_ = feedback
		this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("exact")), errorCode, feedback)
		this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("broad")), body, feedback)
		panic(NewExchangeError(feedback))
	}
	return MkUndefined()
}
