package ccxt

import ()

type Okex struct {
	*ExchangeBase
}

var _ Exchange = (*Okex)(nil)

func init() {
	exchange := &Okex{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Okex) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("okex"),
		"name": MkString("OKEX"),
		"countries": MkArray(&VarArray{
			MkString("CN"),
			MkString("US"),
		}),
		"version":   MkString("v5"),
		"rateLimit": MkInteger(1000),
		"pro":       MkBool(true),
		"has": MkMap(&VarMap{
			"CORS":                MkBool(false),
			"cancelOrder":         MkBool(true),
			"createOrder":         MkBool(true),
			"fetchBalance":        MkBool(true),
			"fetchClosedOrders":   MkBool(true),
			"fetchCurrencies":     MkBool(false),
			"fetchDepositAddress": MkBool(true),
			"fetchDeposits":       MkBool(true),
			"fetchLedger":         MkBool(true),
			"fetchMarkets":        MkBool(true),
			"fetchMyTrades":       MkBool(true),
			"fetchOHLCV":          MkBool(true),
			"fetchOpenOrders":     MkBool(true),
			"fetchOrder":          MkBool(true),
			"fetchOrderBook":      MkBool(true),
			"fetchOrderTrades":    MkBool(true),
			"fetchPosition":       MkBool(true),
			"fetchPositions":      MkBool(true),
			"fetchStatus":         MkBool(true),
			"fetchTicker":         MkBool(true),
			"fetchTickers":        MkBool(true),
			"fetchTime":           MkBool(true),
			"fetchTrades":         MkBool(true),
			"fetchWithdrawals":    MkBool(true),
			"transfer":            MkBool(true),
			"withdraw":            MkBool(true),
		}),
		"timeframes": MkMap(&VarMap{
			"1m":  MkString("1m"),
			"3m":  MkString("3m"),
			"5m":  MkString("5m"),
			"15m": MkString("15m"),
			"30m": MkString("30m"),
			"1h":  MkString("1H"),
			"2h":  MkString("2H"),
			"4h":  MkString("4H"),
			"6h":  MkString("6H"),
			"12h": MkString("12H"),
			"1d":  MkString("1D"),
			"1w":  MkString("1W"),
			"1M":  MkString("1M"),
			"3M":  MkString("3M"),
			"6M":  MkString("6M"),
			"1y":  MkString("1Y"),
		}),
		"hostname": MkString("www.okex.com"),
		"urls": MkMap(&VarMap{
			"logo":     MkString("https://user-images.githubusercontent.com/1294454/32552768-0d6dd3c6-c4a6-11e7-90f8-c043b64756a7.jpg"),
			"api":      MkMap(&VarMap{"rest": MkString("https://{hostname}")}),
			"www":      MkString("https://www.okex.com"),
			"doc":      MkString("https://www.okex.com/docs/en/"),
			"fees":     MkString("https://www.okex.com/pages/products/fees.html"),
			"referral": MkString("https://www.okex.com/join/1888677"),
			"test":     MkMap(&VarMap{"rest": MkString("https://testnet.okex.com")}),
		}),
		"api": MkMap(&VarMap{
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("market/tickers"),
				MkString("market/ticker"),
				MkString("market/index-tickers"),
				MkString("market/books"),
				MkString("market/candles"),
				MkString("market/history-candles"),
				MkString("market/index-candles"),
				MkString("market/mark-price-candles"),
				MkString("market/trades"),
				MkString("market/platform-24-volume"),
				MkString("market/open-oracle"),
				MkString("market/oracle"),
				MkString("public/instruments"),
				MkString("public/delivery-exercise-history"),
				MkString("public/open-interest"),
				MkString("public/funding-rate"),
				MkString("public/funding-rate-history"),
				MkString("public/price-limit"),
				MkString("public/opt-summary"),
				MkString("public/estimated-price"),
				MkString("public/discount-rate-interest-free-quota"),
				MkString("public/time"),
				MkString("public/liquidation-orders"),
				MkString("public/mark-price"),
				MkString("public/tier"),
				MkString("public/position-tiers"),
				MkString("public/underlying"),
				MkString("public/interest-rate-loan-quota"),
				MkString("system/status"),
			})}),
			"private": MkMap(&VarMap{
				"get": MkArray(&VarArray{
					MkString("account/account-position-risk"),
					MkString("account/balance"),
					MkString("account/positions"),
					MkString("account/bills"),
					MkString("account/bills-archive"),
					MkString("account/config"),
					MkString("account/max-size"),
					MkString("account/max-avail-size"),
					MkString("account/leverage-info"),
					MkString("account/max-loan"),
					MkString("account/trade-fee"),
					MkString("account/interest-accrued"),
					MkString("account/interest-rate"),
					MkString("account/max-withdrawal"),
					MkString("asset/deposit-address"),
					MkString("asset/balances"),
					MkString("asset/deposit-history"),
					MkString("asset/withdrawal-history"),
					MkString("asset/currencies"),
					MkString("asset/bills"),
					MkString("asset/piggy-balance"),
					MkString("trade/order"),
					MkString("trade/orders-pending"),
					MkString("trade/orders-history"),
					MkString("trade/orders-history-archive"),
					MkString("trade/fills"),
					MkString("trade/fills-history"),
					MkString("trade/orders-algo-pending"),
					MkString("trade/orders-algo-history"),
					MkString("account/subaccount/balances"),
					MkString("asset/subaccount/bills"),
					MkString("users/subaccount/list"),
				}),
				"post": MkArray(&VarArray{
					MkString("account/set-position-mode"),
					MkString("account/set-leverage"),
					MkString("account/position/margin-balance"),
					MkString("account/set-greeks"),
					MkString("asset/transfer"),
					MkString("asset/withdrawal"),
					MkString("asset/purchase_redempt"),
					MkString("trade/order"),
					MkString("trade/batch-orders"),
					MkString("trade/cancel-order"),
					MkString("trade/cancel-batch-orders"),
					MkString("trade/amend-order"),
					MkString("trade/amend-batch-orders"),
					MkString("trade/close-position"),
					MkString("trade/order-algo"),
					MkString("trade/cancel-algos"),
					MkString("users/subaccount/delete-apikey"),
					MkString("users/subaccount/modify-apikey"),
					MkString("users/subaccount/apikey"),
					MkString("asset/subaccount/transfer"),
				}),
			}),
		}),
		"fees": MkMap(&VarMap{
			"trading": MkMap(&VarMap{
				"taker": this.ParseNumber(MkString("0.0015")),
				"maker": this.ParseNumber(MkString("0.0010")),
			}),
			"spot": MkMap(&VarMap{
				"taker": this.ParseNumber(MkString("0.0015")),
				"maker": this.ParseNumber(MkString("0.0010")),
			}),
			"futures": MkMap(&VarMap{
				"taker": this.ParseNumber(MkString("0.0005")),
				"maker": this.ParseNumber(MkString("0.0002")),
			}),
			"swap": MkMap(&VarMap{
				"taker": this.ParseNumber(MkString("0.00050")),
				"maker": this.ParseNumber(MkString("0.00020")),
			}),
		}),
		"requiredCredentials": MkMap(&VarMap{
			"apiKey":   MkBool(true),
			"secret":   MkBool(true),
			"password": MkBool(true),
		}),
		"exceptions": MkMap(&VarMap{
			"exact": MkMap(&VarMap{
				"1":     ExchangeError,
				"2":     ExchangeError,
				"50000": BadRequest,
				"50001": OnMaintenance,
				"50002": BadRequest,
				"50004": RequestTimeout,
				"50005": ExchangeNotAvailable,
				"50006": BadRequest,
				"50007": AccountSuspended,
				"50008": AuthenticationError,
				"50009": AccountSuspended,
				"50010": ExchangeError,
				"50011": RateLimitExceeded,
				"50012": ExchangeError,
				"50013": ExchangeNotAvailable,
				"50014": ExchangeError,
				"50015": ExchangeError,
				"50016": ExchangeError,
				"50017": ExchangeError,
				"50018": ExchangeError,
				"50019": ExchangeError,
				"50020": ExchangeError,
				"50021": ExchangeError,
				"50022": ExchangeError,
				"50023": ExchangeError,
				"50024": ExchangeError,
				"50025": ExchangeError,
				"50026": ExchangeError,
				"50027": ExchangeError,
				"50028": ExchangeError,
				"50100": ExchangeError,
				"50101": ExchangeError,
				"50102": InvalidNonce,
				"50103": AuthenticationError,
				"50104": AuthenticationError,
				"50105": AuthenticationError,
				"50106": AuthenticationError,
				"50107": AuthenticationError,
				"50108": ExchangeError,
				"50109": ExchangeError,
				"50110": PermissionDenied,
				"50111": AuthenticationError,
				"50112": AuthenticationError,
				"50113": AuthenticationError,
				"50114": AuthenticationError,
				"50115": BadRequest,
				"51000": BadRequest,
				"51001": BadSymbol,
				"51002": BadSymbol,
				"51003": BadRequest,
				"51004": InvalidOrder,
				"51005": InvalidOrder,
				"51006": InvalidOrder,
				"51007": InvalidOrder,
				"51008": InsufficientFunds,
				"51009": AccountSuspended,
				"51010": InsufficientFunds,
				"51011": InvalidOrder,
				"51012": ExchangeError,
				"51014": ExchangeError,
				"51015": BadSymbol,
				"51016": InvalidOrder,
				"51017": ExchangeError,
				"51018": ExchangeError,
				"51019": ExchangeError,
				"51020": InvalidOrder,
				"51021": BadSymbol,
				"51022": BadSymbol,
				"51023": ExchangeError,
				"51024": AccountSuspended,
				"51025": ExchangeError,
				"51026": BadSymbol,
				"51027": BadSymbol,
				"51028": BadSymbol,
				"51029": BadSymbol,
				"51030": BadSymbol,
				"51031": InvalidOrder,
				"51100": InvalidOrder,
				"51101": InvalidOrder,
				"51102": InvalidOrder,
				"51103": InvalidOrder,
				"51104": InvalidOrder,
				"51105": InvalidOrder,
				"51106": InvalidOrder,
				"51107": InvalidOrder,
				"51108": InvalidOrder,
				"51109": InvalidOrder,
				"51110": InvalidOrder,
				"51111": BadRequest,
				"51112": InvalidOrder,
				"51113": RateLimitExceeded,
				"51115": InvalidOrder,
				"51116": InvalidOrder,
				"51117": InvalidOrder,
				"51118": InvalidOrder,
				"51119": InsufficientFunds,
				"51120": InvalidOrder,
				"51121": InvalidOrder,
				"51122": InvalidOrder,
				"51124": InvalidOrder,
				"51125": InvalidOrder,
				"51126": InvalidOrder,
				"51127": InsufficientFunds,
				"51128": InvalidOrder,
				"51129": InvalidOrder,
				"51130": BadSymbol,
				"51131": InsufficientFunds,
				"51132": InvalidOrder,
				"51133": InvalidOrder,
				"51134": InvalidOrder,
				"51135": InvalidOrder,
				"51136": InvalidOrder,
				"51137": InvalidOrder,
				"51138": InvalidOrder,
				"51139": InvalidOrder,
				"51201": InvalidOrder,
				"51202": InvalidOrder,
				"51203": InvalidOrder,
				"51204": InvalidOrder,
				"51205": InvalidOrder,
				"51250": InvalidOrder,
				"51251": InvalidOrder,
				"51252": InvalidOrder,
				"51253": InvalidOrder,
				"51254": InvalidOrder,
				"51255": InvalidOrder,
				"51256": InvalidOrder,
				"51257": InvalidOrder,
				"51258": InvalidOrder,
				"51259": InvalidOrder,
				"51260": InvalidOrder,
				"51261": InvalidOrder,
				"51262": InvalidOrder,
				"51263": InvalidOrder,
				"51264": InvalidOrder,
				"51265": InvalidOrder,
				"51267": InvalidOrder,
				"51268": InvalidOrder,
				"51269": InvalidOrder,
				"51270": InvalidOrder,
				"51271": InvalidOrder,
				"51272": InvalidOrder,
				"51273": InvalidOrder,
				"51274": InvalidOrder,
				"51275": InvalidOrder,
				"51276": InvalidOrder,
				"51277": InvalidOrder,
				"51278": InvalidOrder,
				"51279": InvalidOrder,
				"51280": InvalidOrder,
				"51400": OrderNotFound,
				"51401": OrderNotFound,
				"51402": OrderNotFound,
				"51403": InvalidOrder,
				"51404": InvalidOrder,
				"51405": ExchangeError,
				"51406": ExchangeError,
				"51407": BadRequest,
				"51408": ExchangeError,
				"51409": ExchangeError,
				"51410": ExchangeError,
				"51500": ExchangeError,
				"51501": ExchangeError,
				"51502": InsufficientFunds,
				"51503": ExchangeError,
				"51506": ExchangeError,
				"51508": ExchangeError,
				"51509": ExchangeError,
				"51510": ExchangeError,
				"51511": ExchangeError,
				"51600": ExchangeError,
				"51601": ExchangeError,
				"51602": ExchangeError,
				"51603": OrderNotFound,
				"52000": ExchangeError,
				"54000": ExchangeError,
				"54001": ExchangeError,
				"58000": ExchangeError,
				"58001": AuthenticationError,
				"58002": PermissionDenied,
				"58003": ExchangeError,
				"58004": AccountSuspended,
				"58005": ExchangeError,
				"58006": ExchangeError,
				"58007": ExchangeError,
				"58100": ExchangeError,
				"58101": AccountSuspended,
				"58102": RateLimitExceeded,
				"58103": ExchangeError,
				"58104": ExchangeError,
				"58105": ExchangeError,
				"58106": ExchangeError,
				"58107": ExchangeError,
				"58108": ExchangeError,
				"58109": ExchangeError,
				"58110": ExchangeError,
				"58111": ExchangeError,
				"58112": ExchangeError,
				"58114": ExchangeError,
				"58115": ExchangeError,
				"58116": ExchangeError,
				"58117": ExchangeError,
				"58200": ExchangeError,
				"58201": ExchangeError,
				"58202": ExchangeError,
				"58203": InvalidAddress,
				"58204": AccountSuspended,
				"58205": ExchangeError,
				"58206": ExchangeError,
				"58207": InvalidAddress,
				"58208": ExchangeError,
				"58209": ExchangeError,
				"58210": ExchangeError,
				"58211": ExchangeError,
				"58212": ExchangeError,
				"58213": AuthenticationError,
				"58300": ExchangeError,
				"58350": InsufficientFunds,
				"59000": ExchangeError,
				"59001": ExchangeError,
				"59100": ExchangeError,
				"59101": ExchangeError,
				"59102": ExchangeError,
				"59103": InsufficientFunds,
				"59104": ExchangeError,
				"59105": ExchangeError,
				"59106": ExchangeError,
				"59107": ExchangeError,
				"59108": InsufficientFunds,
				"59109": ExchangeError,
				"59200": InsufficientFunds,
				"59201": InsufficientFunds,
				"59300": ExchangeError,
				"59301": ExchangeError,
				"59401": ExchangeError,
				"59500": ExchangeError,
				"59501": ExchangeError,
				"59502": ExchangeError,
				"59503": ExchangeError,
				"59504": ExchangeError,
				"59505": ExchangeError,
				"59506": ExchangeError,
				"59507": ExchangeError,
				"59508": AccountSuspended,
				"60001": AuthenticationError,
				"60002": AuthenticationError,
				"60003": AuthenticationError,
				"60004": AuthenticationError,
				"60005": AuthenticationError,
				"60006": InvalidNonce,
				"60007": AuthenticationError,
				"60008": AuthenticationError,
				"60009": AuthenticationError,
				"60010": AuthenticationError,
				"60011": AuthenticationError,
				"60012": BadRequest,
				"60013": BadRequest,
				"60014": RateLimitExceeded,
				"60015": NetworkError,
				"60016": ExchangeNotAvailable,
				"60017": BadRequest,
				"60018": BadRequest,
				"60019": BadRequest,
				"63999": ExchangeError,
			}),
			"broad": MkMap(&VarMap{}),
		}),
		"httpExceptions": MkMap(&VarMap{"429": ExchangeNotAvailable}),
		"precisionMode":  TICK_SIZE,
		"options": MkMap(&VarMap{
			"fetchOHLCV":                        MkMap(&VarMap{"type": MkString("Candles")}),
			"createMarketBuyOrderRequiresPrice": MkBool(true),
			"fetchMarkets": MkArray(&VarArray{
				MkString("spot"),
				MkString("futures"),
				MkString("swap"),
			}),
			"defaultType": MkString("spot"),
			"fetchLedger": MkMap(&VarMap{"method": MkString("privateGetAccountBills")}),
			"accountsByType": MkMap(&VarMap{
				"spot":    MkString("1"),
				"futures": MkString("3"),
				"margin":  MkString("5"),
				"funding": MkString("6"),
				"swap":    MkString("9"),
				"option":  MkString("12"),
				"trading": MkString("18"),
				"unified": MkString("18"),
			}),
			"typesByAccount": MkMap(&VarMap{
				"1":  MkString("spot"),
				"3":  MkString("futures"),
				"5":  MkString("margin"),
				"6":  MkString("funding"),
				"9":  MkString("swap"),
				"12": MkString("option"),
				"18": MkString("trading"),
			}),
			"brokerId": MkString("e847386590ce4dBC"),
		}),
		"commonCurrencies": MkMap(&VarMap{
			"AE":   MkString("AET"),
			"BOX":  MkString("DefiBox"),
			"HOT":  MkString("Hydro Protocol"),
			"HSR":  MkString("HC"),
			"MAG":  MkString("Maggie"),
			"SBTC": MkString("Super Bitcoin"),
			"YOYO": MkString("YOYOW"),
			"WIN":  MkString("WinToken"),
		}),
	}))
}

func (this *Okex) FetchStatus(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetSystemStatus"), params)
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	timestamp := this.Milliseconds()
	_ = timestamp
	update := MkMap(&VarMap{
		"info":    response,
		"updated": timestamp,
		"status":  MkString("ok"),
		"eta":     MkUndefined(),
	})
	_ = update
	for i := MkInteger(0); IsTrue(OpLw(i, data.Length)); OpIncr(&i) {
		event := *(data).At(i)
		_ = event
		state := this.SafeString(event, MkString("state"))
		_ = state
		if IsTrue(OpEq2(state, MkString("ongoing"))) {
			*(update).At(MkString("eta")) = this.SafeInteger(event, MkString("end"))
			*(update).At(MkString("status")) = MkString("maintenance")
		}
	}
	*this.At(MkString("status")) = this.Extend(*this.At(MkString("status")), update)
	return *this.At(MkString("status"))
}

func (this *Okex) FetchTime(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetPublicTime"), params)
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	first := this.SafeValue(data, MkInteger(0), MkMap(&VarMap{}))
	_ = first
	return this.SafeInteger(first, MkString("ts"))
}

func (this *Okex) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	types := this.SafeValue(*this.At(MkString("options")), MkString("fetchMarkets"))
	_ = types
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, types.Length)); OpIncr(&i) {
		markets := this.FetchMarketsByType(*(types).At(i), params)
		_ = markets
		result = this.ArrayConcat(result, markets)
	}
	return result
}

func (this *Okex) ParseMarkets(goArgs ...*Variant) *Variant {
	markets := GoGetArg(goArgs, 0, MkUndefined())
	_ = markets
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, markets.Length)); OpIncr(&i) {
		result.Push(this.ParseMarket(*(markets).At(i)))
	}
	return result
}

func (this *Okex) ParseMarket(goArgs ...*Variant) *Variant {
	market := GoGetArg(goArgs, 0, MkUndefined())
	_ = market
	id := this.SafeString(market, MkString("instId"))
	_ = id
	type_ := this.SafeStringLower(market, MkString("instType"))
	_ = type_
	spot := OpEq2(type_, MkString("spot"))
	_ = spot
	futures := OpEq2(type_, MkString("futures"))
	_ = futures
	swap := OpEq2(type_, MkString("swap"))
	_ = swap
	option := OpEq2(type_, MkString("option"))
	_ = option
	baseId := this.SafeString(market, MkString("baseCcy"))
	_ = baseId
	quoteId := this.SafeString(market, MkString("quoteCcy"))
	_ = quoteId
	settleCurrency := this.SafeString(market, MkString("settleCcy"))
	_ = settleCurrency
	underlying := this.SafeString(market, MkString("uly"))
	_ = underlying
	if IsTrue(OpAnd(OpNotEq2(underlying, MkUndefined()), OpNot(spot))) {
		parts := underlying.Split(MkString("-"))
		_ = parts
		baseId = this.SafeString(parts, MkInteger(0))
		quoteId = this.SafeString(parts, MkInteger(1))
	}
	inverse := OpEq2(baseId, settleCurrency)
	_ = inverse
	linear := OpEq2(quoteId, settleCurrency)
	_ = linear
	base := this.SafeCurrencyCode(baseId)
	_ = base
	quote := this.SafeCurrencyCode(quoteId)
	_ = quote
	symbol := OpTrinary(spot, OpAdd(base, OpAdd(MkString("/"), quote)), id)
	_ = symbol
	tickSize := this.SafeString(market, MkString("tickSz"))
	_ = tickSize
	precision := MkMap(&VarMap{
		"amount": this.SafeNumber(market, MkString("lotSz")),
		"price":  this.ParseNumber(tickSize),
	})
	_ = precision
	minAmountString := this.SafeString(market, MkString("minSz"))
	_ = minAmountString
	minAmount := this.ParseNumber(minAmountString)
	_ = minAmount
	minCost := OpCopy(MkUndefined())
	_ = minCost
	if IsTrue(OpAnd(OpNotEq2(minAmount, MkUndefined()), OpNotEq2(tickSize, MkUndefined()))) {
		minCost = this.ParseNumber(Precise.StringMul(tickSize, minAmountString))
	}
	active := OpCopy(MkBool(true))
	_ = active
	fees := this.SafeValue2(*this.At(MkString("fees")), type_, MkString("trading"), MkMap(&VarMap{}))
	_ = fees
	contractSize := this.SafeString(market, MkString("ctVal"))
	_ = contractSize
	return this.Extend(fees, MkMap(&VarMap{
		"id":           id,
		"symbol":       symbol,
		"base":         base,
		"quote":        quote,
		"baseId":       baseId,
		"quoteId":      quoteId,
		"info":         market,
		"type":         type_,
		"spot":         spot,
		"futures":      futures,
		"swap":         swap,
		"option":       option,
		"linear":       linear,
		"inverse":      inverse,
		"active":       active,
		"contractSize": contractSize,
		"precision":    precision,
		"limits": MkMap(&VarMap{
			"amount": MkMap(&VarMap{
				"min": minAmount,
				"max": MkUndefined(),
			}),
			"price": MkMap(&VarMap{
				"min": *(precision).At(MkString("price")),
				"max": MkUndefined(),
			}),
			"cost": MkMap(&VarMap{
				"min": minCost,
				"max": MkUndefined(),
			}),
		}),
	}))
}

func (this *Okex) FetchMarketsByType(goArgs ...*Variant) *Variant {
	type_ := GoGetArg(goArgs, 0, MkUndefined())
	_ = type_
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	uppercaseType := type_.ToUpperCase()
	_ = uppercaseType
	request := MkMap(&VarMap{"instType": uppercaseType})
	_ = request
	if IsTrue(OpEq2(uppercaseType, MkString("OPTION"))) {
		defaultUnderlying := this.SafeValue(*this.At(MkString("options")), MkString("defaultUnderlying"), MkString("BTC-USD"))
		_ = defaultUnderlying
		currencyId := this.SafeString2(params, MkString("uly"), MkString("marketId"), defaultUnderlying)
		_ = currencyId
		if IsTrue(OpEq2(currencyId, MkUndefined())) {
			panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchMarketsByType requires an underlying uly or marketId parameter for options markets"))))
		} else {
			*(request).At(MkString("uly")) = OpCopy(currencyId)
		}
	}
	response := this.Call(MkString("publicGetPublicInstruments"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseMarkets(data)
}

func (this *Okex) FetchCurrencies(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("privateGetAssetCurrencies"), params)
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	result := MkMap(&VarMap{})
	_ = result
	dataByCurrencyId := this.GroupBy(data, MkString("ccy"))
	_ = dataByCurrencyId
	currencyIds := GoGetKeys(dataByCurrencyId)
	_ = currencyIds
	for i := MkInteger(0); IsTrue(OpLw(i, currencyIds.Length)); OpIncr(&i) {
		currencyId := *(currencyIds).At(i)
		_ = currencyId
		chains := *(dataByCurrencyId).At(currencyId)
		_ = chains
		first := this.SafeValue(chains, MkInteger(0))
		_ = first
		id := this.SafeString(first, MkString("ccy"))
		_ = id
		code := this.SafeCurrencyCode(id)
		_ = code
		precision := MkNumber(0.00000001)
		_ = precision
		name := this.SafeString(first, MkString("name"))
		_ = name
		if IsTrue(OpAnd(OpNotEq2(name, MkUndefined()), OpLw(name.Length, MkInteger(1)))) {
			name = OpCopy(MkUndefined())
		}
		canDeposit := this.SafeValue(first, MkString("canDep"))
		_ = canDeposit
		canWithdraw := this.SafeValue(first, MkString("canWd"))
		_ = canWithdraw
		canInternal := this.SafeValue(first, MkString("canInternal"))
		_ = canInternal
		active := OpTrinary(OpAnd(canDeposit, OpAnd(canWithdraw, canInternal)), MkBool(true), MkBool(false))
		_ = active
		*(result).At(code) = MkMap(&VarMap{
			"id":        id,
			"code":      code,
			"info":      chains,
			"type":      MkUndefined(),
			"name":      name,
			"active":    active,
			"fee":       this.SafeNumber(first, MkString("minFee")),
			"precision": precision,
			"limits": MkMap(&VarMap{
				"amount": MkMap(&VarMap{
					"min": MkUndefined(),
					"max": MkUndefined(),
				}),
				"withdraw": MkMap(&VarMap{
					"min": this.SafeNumber(first, MkString("ccy")),
					"max": MkUndefined(),
				}),
			}),
		})
	}
	return result
}

func (this *Okex) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"instId": *(market).At(MkString("id"))})
	_ = request
	limit = OpTrinary(OpEq2(limit, MkUndefined()), MkInteger(20), limit)
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("sz")) = OpCopy(limit)
	}
	response := this.Call(MkString("publicGetMarketBooks"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	first := this.SafeValue(data, MkInteger(0), MkMap(&VarMap{}))
	_ = first
	timestamp := this.SafeInteger(first, MkString("ts"))
	_ = timestamp
	return this.ParseOrderBook(first, symbol, timestamp)
}

func (this *Okex) ParseTicker(goArgs ...*Variant) *Variant {
	ticker := GoGetArg(goArgs, 0, MkUndefined())
	_ = ticker
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.SafeInteger(ticker, MkString("ts"))
	_ = timestamp
	marketId := this.SafeString(ticker, MkString("instId"))
	_ = marketId
	market = this.SafeMarket(marketId, market, MkString("-"))
	symbol := *(market).At(MkString("symbol"))
	_ = symbol
	last := this.SafeNumber(ticker, MkString("last"))
	_ = last
	open := this.SafeNumber(ticker, MkString("open24h"))
	_ = open
	quoteVolume := this.SafeNumber(ticker, MkString("volCcy24h"))
	_ = quoteVolume
	baseVolume := this.SafeNumber(ticker, MkString("vol24h"))
	_ = baseVolume
	vwap := this.Vwap(baseVolume, quoteVolume)
	_ = vwap
	return MkMap(&VarMap{
		"symbol":        symbol,
		"timestamp":     timestamp,
		"datetime":      this.Iso8601(timestamp),
		"high":          this.SafeNumber(ticker, MkString("high24h")),
		"low":           this.SafeNumber(ticker, MkString("low24h")),
		"bid":           this.SafeNumber(ticker, MkString("bidPx")),
		"bidVolume":     this.SafeNumber(ticker, MkString("bidSz")),
		"ask":           this.SafeNumber(ticker, MkString("askPx")),
		"askVolume":     this.SafeNumber(ticker, MkString("askSz")),
		"vwap":          vwap,
		"open":          open,
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

func (this *Okex) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"instId": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("publicGetMarketTicker"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	first := this.SafeValue(data, MkInteger(0), MkMap(&VarMap{}))
	_ = first
	return this.ParseTicker(first, market)
}

func (this *Okex) FetchTickersByType(goArgs ...*Variant) *Variant {
	type_ := GoGetArg(goArgs, 0, MkUndefined())
	_ = type_
	symbols := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	uppercaseType := type_.ToUpperCase()
	_ = uppercaseType
	request := MkMap(&VarMap{"instType": type_.ToUpperCase()})
	_ = request
	if IsTrue(OpEq2(uppercaseType, MkString("OPTION"))) {
		defaultUnderlying := this.SafeValue(*this.At(MkString("options")), MkString("defaultUnderlying"), MkString("BTC-USD"))
		_ = defaultUnderlying
		currencyId := this.SafeString2(params, MkString("uly"), MkString("marketId"), defaultUnderlying)
		_ = currencyId
		if IsTrue(OpEq2(currencyId, MkUndefined())) {
			panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchTickersByType requires an underlying uly or marketId parameter for options markets"))))
		} else {
			*(request).At(MkString("uly")) = OpCopy(currencyId)
		}
	}
	response := this.Call(MkString("publicGetMarketTickers"), this.Extend(request, params))
	_ = response
	tickers := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = tickers
	return this.ParseTickers(tickers, symbols)
}

func (this *Okex) FetchTickers(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	defaultType := this.SafeString2(*this.At(MkString("options")), MkString("fetchTickers"), MkString("defaultType"))
	_ = defaultType
	type_ := this.SafeString(params, MkString("type"), defaultType)
	_ = type_
	return this.FetchTickersByType(type_, symbols, this.Omit(params, MkString("type")))
}

func (this *Okex) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	id := this.SafeString(trade, MkString("tradeId"))
	_ = id
	marketId := this.SafeString(trade, MkString("instId"))
	_ = marketId
	market = this.SafeMarket(marketId, market, MkString("-"))
	symbol := *(market).At(MkString("symbol"))
	_ = symbol
	timestamp := this.SafeInteger(trade, MkString("ts"))
	_ = timestamp
	priceString := this.SafeString2(trade, MkString("fillPx"), MkString("px"))
	_ = priceString
	amountString := this.SafeString2(trade, MkString("fillSz"), MkString("sz"))
	_ = amountString
	price := this.ParseNumber(priceString)
	_ = price
	amount := this.ParseNumber(amountString)
	_ = amount
	cost := this.ParseNumber(Precise.StringMul(priceString, amountString))
	_ = cost
	side := this.SafeString(trade, MkString("side"))
	_ = side
	orderId := this.SafeString(trade, MkString("ordId"))
	_ = orderId
	feeCostString := this.SafeString(trade, MkString("fee"))
	_ = feeCostString
	fee := OpCopy(MkUndefined())
	_ = fee
	if IsTrue(OpNotEq2(feeCostString, MkUndefined())) {
		feeCostSigned := Precise.StringNeg(feeCostString)
		_ = feeCostSigned
		feeCurrencyId := this.SafeString(trade, MkString("feeCcy"))
		_ = feeCurrencyId
		feeCurrencyCode := this.SafeCurrencyCode(feeCurrencyId)
		_ = feeCurrencyCode
		fee = MkMap(&VarMap{
			"cost":     this.ParseNumber(feeCostSigned),
			"currency": feeCurrencyCode,
		})
	}
	takerOrMaker := this.SafeString(trade, MkString("execType"))
	_ = takerOrMaker
	if IsTrue(OpEq2(takerOrMaker, MkString("T"))) {
		takerOrMaker = MkString("taker")
	} else {
		if IsTrue(OpEq2(takerOrMaker, MkString("M"))) {
			takerOrMaker = MkString("maker")
		}
	}
	return MkMap(&VarMap{
		"info":         trade,
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       symbol,
		"id":           id,
		"order":        orderId,
		"type":         MkUndefined(),
		"takerOrMaker": takerOrMaker,
		"side":         side,
		"price":        price,
		"amount":       amount,
		"cost":         cost,
		"fee":          fee,
	})
}

func (this *Okex) FetchTrades(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"instId": *(market).At(MkString("id"))})
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

func (this *Okex) ParseOHLCV(goArgs ...*Variant) *Variant {
	ohlcv := GoGetArg(goArgs, 0, MkUndefined())
	_ = ohlcv
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	return MkArray(&VarArray{
		this.SafeInteger(ohlcv, MkInteger(0)),
		this.SafeNumber(ohlcv, MkInteger(1)),
		this.SafeNumber(ohlcv, MkInteger(2)),
		this.SafeNumber(ohlcv, MkInteger(3)),
		this.SafeNumber(ohlcv, MkInteger(4)),
		this.SafeNumber(ohlcv, MkInteger(5)),
	})
}

func (this *Okex) FetchOHLCV(goArgs ...*Variant) *Variant {
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
		"instId": *(market).At(MkString("id")),
		"bar":    *(*this.At(MkString("timeframes"))).At(timeframe),
	})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	options := this.SafeValue(*this.At(MkString("options")), MkString("fetchOHLCV"), MkMap(&VarMap{}))
	_ = options
	defaultType := this.SafeString(options, MkString("type"), MkString("Candles"))
	_ = defaultType
	type_ := this.SafeString(params, MkString("type"), defaultType)
	_ = type_
	params = this.Omit(params, MkString("type"))
	method := OpAdd(MkString("publicGetMarket"), type_)
	_ = method
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("before")) = OpSub(since, MkInteger(1))
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseOHLCVs(data, market, timeframe, since, limit)
}

func (this *Okex) ParseBalanceByType(goArgs ...*Variant) *Variant {
	type_ := GoGetArg(goArgs, 0, MkUndefined())
	_ = type_
	response := GoGetArg(goArgs, 1, MkUndefined())
	_ = response
	if IsTrue(OpEq2(type_, MkString("funding"))) {
		return this.ParseFundingBalance(response)
	} else {
		return this.ParseTradingBalance(response)
	}
	return MkUndefined()
}

func (this *Okex) ParseTradingBalance(goArgs ...*Variant) *Variant {
	response := GoGetArg(goArgs, 0, MkUndefined())
	_ = response
	result := MkMap(&VarMap{"info": response})
	_ = result
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	first := this.SafeValue(data, MkInteger(0), MkMap(&VarMap{}))
	_ = first
	timestamp := this.SafeInteger(first, MkString("uTime"))
	_ = timestamp
	details := this.SafeValue(first, MkString("details"), MkArray(&VarArray{}))
	_ = details
	for i := MkInteger(0); IsTrue(OpLw(i, details.Length)); OpIncr(&i) {
		balance := *(details).At(i)
		_ = balance
		currencyId := this.SafeString(balance, MkString("ccy"))
		_ = currencyId
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		account := this.Account()
		_ = account
		eq := this.SafeString(balance, MkString("eq"))
		_ = eq
		availEq := this.SafeString(balance, MkString("availEq"))
		_ = availEq
		if IsTrue(OpOr(OpLw(eq.Length, MkInteger(1)), OpLw(availEq.Length, MkInteger(1)))) {
			*(account).At(MkString("free")) = this.SafeString(balance, MkString("availBal"))
			*(account).At(MkString("used")) = this.SafeString(balance, MkString("frozenBal"))
		} else {
			*(account).At(MkString("total")) = OpCopy(eq)
			*(account).At(MkString("free")) = OpCopy(availEq)
		}
		*(result).At(code) = OpCopy(account)
	}
	*(result).At(MkString("timestamp")) = OpCopy(timestamp)
	*(result).At(MkString("datetime")) = this.Iso8601(timestamp)
	return this.ParseBalance(result)
}

func (this *Okex) ParseFundingBalance(goArgs ...*Variant) *Variant {
	response := GoGetArg(goArgs, 0, MkUndefined())
	_ = response
	result := MkMap(&VarMap{"info": response})
	_ = result
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	for i := MkInteger(0); IsTrue(OpLw(i, data.Length)); OpIncr(&i) {
		balance := *(data).At(i)
		_ = balance
		currencyId := this.SafeString(balance, MkString("ccy"))
		_ = currencyId
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		account := this.Account()
		_ = account
		*(account).At(MkString("total")) = this.SafeString(balance, MkString("bal"))
		*(account).At(MkString("free")) = this.SafeString(balance, MkString("availBal"))
		*(account).At(MkString("used")) = this.SafeString(balance, MkString("frozenBal"))
		*(result).At(code) = OpCopy(account)
	}
	return this.ParseBalance(result)
}

func (this *Okex) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	defaultType := this.SafeString(*this.At(MkString("options")), MkString("defaultType"))
	_ = defaultType
	options := this.SafeValue(*this.At(MkString("options")), MkString("fetchBalance"), MkMap(&VarMap{}))
	_ = options
	type_ := this.SafeString(options, MkString("type"), defaultType)
	_ = type_
	type_ = this.SafeString(params, MkString("type"), type_)
	params = this.Omit(params, MkString("type"))
	method := OpCopy(MkUndefined())
	_ = method
	if IsTrue(OpEq2(type_, MkString("funding"))) {
		method = MkString("privateGetAssetBalances")
	} else {
		method = MkString("privateGetAccountBalance")
	}
	request := MkMap(&VarMap{})
	_ = request
	response := this.Call(method, this.Extend(request, params))
	_ = response
	return this.ParseBalanceByType(type_, response)
}

func (this *Okex) CreateOrder(goArgs ...*Variant) *Variant {
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
		"instId":  *(market).At(MkString("id")),
		"tdMode":  MkString("cash"),
		"side":    side,
		"ordType": type_,
	})
	_ = request
	clientOrderId := this.SafeString2(params, MkString("clOrdId"), MkString("clientOrderId"))
	_ = clientOrderId
	if IsTrue(OpEq2(clientOrderId, MkUndefined())) {
		brokerId := this.SafeString(*this.At(MkString("options")), MkString("brokerId"))
		_ = brokerId
		if IsTrue(OpNotEq2(brokerId, MkUndefined())) {
			*(request).At(MkString("clOrdId")) = OpAdd(brokerId, this.Uuid16())
		}
	} else {
		*(request).At(MkString("clOrdId")) = OpCopy(clientOrderId)
		params = this.Omit(params, MkArray(&VarArray{
			MkString("clOrdId"),
			MkString("clientOrderId"),
		}))
	}
	if IsTrue(OpEq2(type_, MkString("market"))) {
		if IsTrue(OpEq2(side, MkString("buy"))) {
			notional := this.SafeNumber(params, MkString("sz"))
			_ = notional
			createMarketBuyOrderRequiresPrice := this.SafeValue(*this.At(MkString("options")), MkString("createMarketBuyOrderRequiresPrice"), MkBool(true))
			_ = createMarketBuyOrderRequiresPrice
			if IsTrue(createMarketBuyOrderRequiresPrice) {
				if IsTrue(OpNotEq2(price, MkUndefined())) {
					if IsTrue(OpEq2(notional, MkUndefined())) {
						notional = OpMulti(amount, price)
					}
				} else {
					if IsTrue(OpEq2(notional, MkUndefined())) {
						panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")), MkString(" createOrder() requires the price argument with market buy orders to calculate total order cost (amount to spend), where cost = amount * price. Supply a price argument to createOrder() call if you want the cost to be calculated for you from price and amount, or, alternatively, add .options['createMarketBuyOrderRequiresPrice'] = false and supply the total cost value in the 'amount' argument or in the 'sz' extra parameter (the exchange-specific behaviour)"))))
					}
				}
			} else {
				notional = OpTrinary(OpEq2(notional, MkUndefined()), amount, notional)
			}
			precision := *(*(market).At(MkString("precision"))).At(MkString("price"))
			_ = precision
			*(request).At(MkString("sz")) = this.DecimalToPrecision(notional, TRUNCATE, precision, *this.At(MkString("precisionMode")))
		} else {
			*(request).At(MkString("sz")) = this.AmountToPrecision(symbol, amount)
		}
	} else {
		*(request).At(MkString("px")) = this.PriceToPrecision(symbol, price)
		*(request).At(MkString("sz")) = this.AmountToPrecision(symbol, amount)
	}
	response := this.Call(MkString("privatePostTradeOrder"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	first := this.SafeValue(data, MkInteger(0))
	_ = first
	order := this.ParseOrder(first, market)
	_ = order
	return this.Extend(order, MkMap(&VarMap{
		"type": type_,
		"side": side,
	}))
}

func (this *Okex) CancelOrder(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"instId": *(market).At(MkString("id"))})
	_ = request
	clientOrderId := this.SafeString2(params, MkString("clOrdId"), MkString("clientOrderId"))
	_ = clientOrderId
	if IsTrue(OpNotEq2(clientOrderId, MkUndefined())) {
		*(request).At(MkString("clOrdId")) = OpCopy(clientOrderId)
	} else {
		*(request).At(MkString("ordId")) = OpCopy(id)
	}
	query := this.Omit(params, MkArray(&VarArray{
		MkString("clOrdId"),
		MkString("clientOrderId"),
	}))
	_ = query
	response := this.Call(MkString("privatePostTradeCancelOrder"), this.Extend(request, query))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	order := this.SafeValue(data, MkInteger(0))
	_ = order
	return this.ParseOrder(order, market)
}

func (this *Okex) ParseOrderStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"canceled":         MkString("canceled"),
		"live":             MkString("open"),
		"partially_filled": MkString("open"),
		"filled":           MkString("closed"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Okex) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	id := this.SafeString(order, MkString("ordId"))
	_ = id
	timestamp := this.SafeInteger(order, MkString("cTime"))
	_ = timestamp
	lastTradeTimestamp := this.SafeInteger(order, MkString("fillTime"))
	_ = lastTradeTimestamp
	side := this.SafeString(order, MkString("side"))
	_ = side
	type_ := this.SafeString(order, MkString("ordType"))
	_ = type_
	postOnly := OpCopy(MkUndefined())
	_ = postOnly
	timeInForce := OpCopy(MkUndefined())
	_ = timeInForce
	if IsTrue(OpEq2(type_, MkString("post_only"))) {
		postOnly = OpCopy(MkBool(true))
		type_ = MkString("limit")
	} else {
		if IsTrue(OpEq2(type_, MkString("fok"))) {
			timeInForce = MkString("FOK")
			type_ = MkString("limit")
		} else {
			if IsTrue(OpEq2(type_, MkString("ioc"))) {
				timeInForce = MkString("IOC")
				type_ = MkString("limit")
			}
		}
	}
	marketId := this.SafeString(order, MkString("instId"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market, MkString("-"))
	_ = symbol
	filled := this.SafeNumber(order, MkString("accFillSz"))
	_ = filled
	price := this.SafeNumber2(order, MkString("px"), MkString("slOrdPx"))
	_ = price
	average := this.SafeNumber(order, MkString("avgPx"))
	_ = average
	status := this.ParseOrderStatus(this.SafeString(order, MkString("state")))
	_ = status
	feeCostString := this.SafeString(order, MkString("fee"))
	_ = feeCostString
	amount := OpCopy(MkUndefined())
	_ = amount
	cost := OpCopy(MkUndefined())
	_ = cost
	if IsTrue(OpAnd(OpEq2(side, MkString("buy")), OpEq2(type_, MkString("market")))) {
		cost = this.SafeNumber(order, MkString("sz"))
	} else {
		amount = this.SafeNumber(order, MkString("sz"))
	}
	fee := OpCopy(MkUndefined())
	_ = fee
	if IsTrue(OpNotEq2(feeCostString, MkUndefined())) {
		feeCostSigned := Precise.StringNeg(feeCostString)
		_ = feeCostSigned
		feeCurrencyId := this.SafeString(order, MkString("feeCcy"))
		_ = feeCurrencyId
		feeCurrencyCode := this.SafeCurrencyCode(feeCurrencyId)
		_ = feeCurrencyCode
		fee = MkMap(&VarMap{
			"cost":     this.ParseNumber(feeCostSigned),
			"currency": feeCurrencyCode,
		})
	}
	clientOrderId := this.SafeString(order, MkString("clOrdId"))
	_ = clientOrderId
	if IsTrue(OpAnd(OpNotEq2(clientOrderId, MkUndefined()), OpLw(clientOrderId.Length, MkInteger(1)))) {
		clientOrderId = OpCopy(MkUndefined())
	}
	stopPrice := this.SafeNumber(order, MkString("slTriggerPx"))
	_ = stopPrice
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
		"average":            average,
		"cost":               cost,
		"amount":             amount,
		"filled":             filled,
		"remaining":          MkUndefined(),
		"status":             status,
		"fee":                fee,
		"trades":             MkUndefined(),
	}))
}

func (this *Okex) FetchOrder(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"instId": *(market).At(MkString("id"))})
	_ = request
	clientOrderId := this.SafeString2(params, MkString("clOrdId"), MkString("clientOrderId"))
	_ = clientOrderId
	if IsTrue(OpNotEq2(clientOrderId, MkUndefined())) {
		*(request).At(MkString("clOrdId")) = OpCopy(clientOrderId)
	} else {
		*(request).At(MkString("ordId")) = OpCopy(id)
	}
	query := this.Omit(params, MkArray(&VarArray{
		MkString("clOrdId"),
		MkString("clientOrderId"),
	}))
	_ = query
	response := this.Call(MkString("privateGetTradeOrder"), this.Extend(request, query))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	order := this.SafeValue(data, MkInteger(0))
	_ = order
	return this.ParseOrder(order, market)
}

func (this *Okex) FetchOpenOrders(goArgs ...*Variant) *Variant {
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
		*(request).At(MkString("instId")) = *(market).At(MkString("id"))
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("privateGetTradeOrdersPending"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseOrders(data, market, since, limit)
}

func (this *Okex) FetchClosedOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	defaultType := this.SafeString(*this.At(MkString("options")), MkString("defaultType"))
	_ = defaultType
	options := this.SafeString(*this.At(MkString("options")), MkString("fetchClosedOrders"), MkMap(&VarMap{}))
	_ = options
	type_ := this.SafeString(options, MkString("type"), defaultType)
	_ = type_
	type_ = this.SafeString(params, MkString("type"), type_)
	params = this.Omit(params, MkString("type"))
	request := MkMap(&VarMap{})
	_ = request
	market := OpCopy(MkUndefined())
	_ = market
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
		if IsTrue(OpOr(*(market).At(MkString("futures")), *(market).At(MkString("swap")))) {
			type_ = *(market).At(MkString("type"))
		}
		*(request).At(MkString("instId")) = *(market).At(MkString("id"))
	}
	*(request).At(MkString("instType")) = type_.ToUpperCase()
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	method := this.SafeString(options, MkString("method"), MkString("privateGetTradeOrdersHistory"))
	_ = method
	response := this.Call(method, this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseOrders(data, market, since, limit)
}

func (this *Okex) FetchMyTrades(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	request := MkMap(&VarMap{"instType": MkString("SPOT")})
	_ = request
	market := OpCopy(MkUndefined())
	_ = market
	this.LoadMarkets()
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
		*(request).At(MkString("instId")) = *(market).At(MkString("id"))
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("privateGetTradeFillsHistory"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseTrades(data, market, since, limit, params)
}

func (this *Okex) FetchOrderTrades(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"ordId": id})
	_ = request
	return this.FetchMyTrades(symbol, since, limit, this.Extend(request, params))
}

func (this *Okex) FetchLedger(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	options := this.SafeValue(*this.At(MkString("options")), MkString("fetchLedger"), MkMap(&VarMap{}))
	_ = options
	method := this.SafeString(options, MkString("method"))
	_ = method
	method = this.SafeString(params, MkString("method"), method)
	params = this.Omit(params, MkString("method"))
	request := MkMap(&VarMap{})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	currency := OpCopy(MkUndefined())
	_ = currency
	if IsTrue(OpNotEq2(code, MkUndefined())) {
		currency = this.Currency(code)
		*(request).At(MkString("ccy")) = *(currency).At(MkString("id"))
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseLedger(data, currency, since, limit)
}

func (this *Okex) ParseLedgerEntryType(goArgs ...*Variant) *Variant {
	type_ := GoGetArg(goArgs, 0, MkUndefined())
	_ = type_
	types := MkMap(&VarMap{
		"1":  MkString("transfer"),
		"2":  MkString("trade"),
		"3":  MkString("trade"),
		"4":  MkString("rebate"),
		"5":  MkString("trade"),
		"6":  MkString("transfer"),
		"7":  MkString("trade"),
		"8":  MkString("fee"),
		"9":  MkString("trade"),
		"10": MkString("trade"),
		"11": MkString("trade"),
	})
	_ = types
	return this.SafeString(types, type_, type_)
}

func (this *Okex) ParseLedgerEntry(goArgs ...*Variant) *Variant {
	item := GoGetArg(goArgs, 0, MkUndefined())
	_ = item
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	id := this.SafeString(item, MkString("billId"))
	_ = id
	account := OpCopy(MkUndefined())
	_ = account
	referenceId := this.SafeString(item, MkString("ordId"))
	_ = referenceId
	referenceAccount := OpCopy(MkUndefined())
	_ = referenceAccount
	type_ := this.ParseLedgerEntryType(this.SafeString(item, MkString("type")))
	_ = type_
	code := this.SafeCurrencyCode(this.SafeString(item, MkString("ccy")), currency)
	_ = code
	amountString := this.SafeString(item, MkString("balChg"))
	_ = amountString
	amount := this.ParseNumber(amountString)
	_ = amount
	timestamp := this.SafeInteger(item, MkString("ts"))
	_ = timestamp
	feeCostString := this.SafeString(item, MkString("fee"))
	_ = feeCostString
	fee := OpCopy(MkUndefined())
	_ = fee
	if IsTrue(OpNotEq2(feeCostString, MkUndefined())) {
		fee = MkMap(&VarMap{
			"cost":     this.ParseNumber(Precise.StringNeg(feeCostString)),
			"currency": code,
		})
	}
	before := OpCopy(MkUndefined())
	_ = before
	afterString := this.SafeString(item, MkString("bal"))
	_ = afterString
	after := this.ParseNumber(afterString)
	_ = after
	status := MkString("ok")
	_ = status
	marketId := this.SafeString(item, MkString("instId"))
	_ = marketId
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(OpHasMember(marketId, *this.At(MkString("markets_by_id")))) {
		market := *(*this.At(MkString("markets_by_id"))).At(marketId)
		_ = market
		symbol = *(market).At(MkString("symbol"))
	}
	return MkMap(&VarMap{
		"id":               id,
		"info":             item,
		"timestamp":        timestamp,
		"datetime":         this.Iso8601(timestamp),
		"account":          account,
		"referenceId":      referenceId,
		"referenceAccount": referenceAccount,
		"type":             type_,
		"currency":         code,
		"symbol":           symbol,
		"amount":           amount,
		"before":           before,
		"after":            after,
		"status":           status,
		"fee":              fee,
	})
}

func (this *Okex) ParseDepositAddress(goArgs ...*Variant) *Variant {
	depositAddress := GoGetArg(goArgs, 0, MkUndefined())
	_ = depositAddress
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	address := this.SafeString(depositAddress, MkString("addr"))
	_ = address
	tag := this.SafeString2(depositAddress, MkString("tag"), MkString("pmtId"))
	_ = tag
	tag = this.SafeString(depositAddress, MkString("memo"), tag)
	currencyId := this.SafeString(depositAddress, MkString("ccy"))
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

func (this *Okex) FetchDepositAddress(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	parts := code.Split(MkString("-"))
	_ = parts
	currency := this.Currency(*(parts).At(MkInteger(0)))
	_ = currency
	request := MkMap(&VarMap{"ccy": *(currency).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("privateGetAssetDepositAddress"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	addressesByCode := this.ParseDepositAddresses(data)
	_ = addressesByCode
	address := this.SafeValue(addressesByCode, code)
	_ = address
	if IsTrue(OpEq2(address, MkUndefined())) {
		panic(NewInvalidAddress(OpAdd(*this.At(MkString("id")), MkString(" fetchDepositAddress cannot return nonexistent addresses, you should create withdrawal addresses with the exchange website first"))))
	}
	return address
}

func (this *Okex) Withdraw(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpNotEq2(tag, MkUndefined())) {
		address = OpAdd(address, OpAdd(MkString(":"), tag))
	}
	fee := this.SafeString(params, MkString("fee"))
	_ = fee
	if IsTrue(OpEq2(fee, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" withdraw() requires a `fee` string parameter, network transaction fee must be â¥ 0. Withdrawals to OKCoin or OKEx are fee-free, please set '0'. Withdrawing to external digital asset address requires network transaction fee."))))
	}
	request := MkMap(&VarMap{
		"ccy":    *(currency).At(MkString("id")),
		"toAddr": address,
		"dest":   MkString("4"),
		"amt":    this.NumberToString(amount),
		"fee":    this.NumberToString(fee),
	})
	_ = request
	if IsTrue(OpHasMember(MkString("password"), params)) {
		*(request).At(MkString("pwd")) = *(params).At(MkString("password"))
	} else {
		if IsTrue(OpHasMember(MkString("pwd"), params)) {
			*(request).At(MkString("pwd")) = *(params).At(MkString("pwd"))
		}
	}
	query := this.Omit(params, MkArray(&VarArray{
		MkString("fee"),
		MkString("password"),
		MkString("pwd"),
	}))
	_ = query
	if IsTrue(OpNot(OpHasMember(MkString("pwd"), request))) {
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), MkString(" withdraw() requires a password parameter or a pwd parameter, it must be the funding password, not the API passphrase"))))
	}
	response := this.Call(MkString("privatePostAssetWithdrawal"), this.Extend(request, query))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	transaction := this.SafeValue(data, MkInteger(0))
	_ = transaction
	return this.ParseTransaction(transaction, currency)
}

func (this *Okex) FetchDeposits(goArgs ...*Variant) *Variant {
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
		*(request).At(MkString("ccy")) = *(currency).At(MkString("id"))
	}
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("after")) = OpCopy(since)
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("privateGetAssetDepositHistory"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseTransactions(data, currency, since, limit, params)
}

func (this *Okex) FetchWithdrawals(goArgs ...*Variant) *Variant {
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
		*(request).At(MkString("ccy")) = *(currency).At(MkString("id"))
	}
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("after")) = OpCopy(since)
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("privateGetAssetWithdrawalHistory"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseTransactions(data, currency, since, limit, params)
}

func (this *Okex) ParseTransactionStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"-3": MkString("pending"),
		"-2": MkString("canceled"),
		"-1": MkString("failed"),
		"0":  MkString("pending"),
		"1":  MkString("pending"),
		"2":  MkString("ok"),
		"3":  MkString("pending"),
		"4":  MkString("pending"),
		"5":  MkString("pending"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Okex) ParseTransaction(goArgs ...*Variant) *Variant {
	transaction := GoGetArg(goArgs, 0, MkUndefined())
	_ = transaction
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	type_ := OpCopy(MkUndefined())
	_ = type_
	id := OpCopy(MkUndefined())
	_ = id
	withdrawalId := this.SafeString(transaction, MkString("wdId"))
	_ = withdrawalId
	addressFrom := this.SafeString(transaction, MkString("from"))
	_ = addressFrom
	addressTo := this.SafeString(transaction, MkString("to"))
	_ = addressTo
	address := OpCopy(addressTo)
	_ = address
	tagTo := this.SafeString2(transaction, MkString("tag"), MkString("memo"))
	_ = tagTo
	tagTo = this.SafeString2(transaction, MkString("pmtId"), tagTo)
	if IsTrue(OpNotEq2(withdrawalId, MkUndefined())) {
		type_ = MkString("withdrawal")
		id = OpCopy(withdrawalId)
	} else {
		id = this.SafeString(transaction, MkString("depId"))
		type_ = MkString("deposit")
	}
	currencyId := this.SafeString(transaction, MkString("ccy"))
	_ = currencyId
	code := this.SafeCurrencyCode(currencyId)
	_ = code
	amount := this.SafeNumber(transaction, MkString("amt"))
	_ = amount
	status := this.ParseTransactionStatus(this.SafeString(transaction, MkString("state")))
	_ = status
	txid := this.SafeString(transaction, MkString("txId"))
	_ = txid
	timestamp := this.SafeInteger(transaction, MkString("ts"))
	_ = timestamp
	feeCost := OpCopy(MkUndefined())
	_ = feeCost
	if IsTrue(OpEq2(type_, MkString("deposit"))) {
		feeCost = MkInteger(0)
	} else {
		feeCost = this.SafeNumber(transaction, MkString("fee"))
	}
	return MkMap(&VarMap{
		"info":        transaction,
		"id":          id,
		"currency":    code,
		"amount":      amount,
		"addressFrom": addressFrom,
		"addressTo":   addressTo,
		"address":     address,
		"tagFrom":     MkUndefined(),
		"tagTo":       tagTo,
		"tag":         tagTo,
		"status":      status,
		"type":        type_,
		"updated":     MkUndefined(),
		"txid":        txid,
		"timestamp":   timestamp,
		"datetime":    this.Iso8601(timestamp),
		"fee": MkMap(&VarMap{
			"currency": code,
			"cost":     feeCost,
		}),
	})
}

func (this *Okex) FetchPosition(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	type_ := this.SafeString(params, MkString("type"))
	_ = type_
	params = this.Omit(params, MkString("type"))
	request := MkMap(&VarMap{"instId": *(market).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(type_, MkUndefined())) {
		*(request).At(MkString("instType")) = type_.ToUpperCase()
	}
	params = this.Omit(params, MkString("type"))
	response := this.Call(MkString("privateGetAccountPositions"), params)
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParsePosition(this.SafeValue(data, MkInteger(0)))
}

func (this *Okex) FetchPositions(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	type_ := this.SafeString(params, MkString("type"))
	_ = type_
	params = this.Omit(params, MkString("type"))
	request := MkMap(&VarMap{})
	_ = request
	if IsTrue(OpNotEq2(type_, MkUndefined())) {
		*(request).At(MkString("instType")) = type_.ToUpperCase()
	}
	params = this.Omit(params, MkString("type"))
	response := this.Call(MkString("privateGetAccountPositions"), this.Extend(request, params))
	_ = response
	positions := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = positions
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, positions.Length)); OpIncr(&i) {
		entry := *(positions).At(i)
		_ = entry
		instrument := this.SafeString(entry, MkString("instType"))
		_ = instrument
		if IsTrue(OpOr(OpEq2(instrument, MkString("FUTURES")), OpEq2(instrument, MkString("SWAP")))) {
			result.Push(this.ParsePosition(*(positions).At(i)))
		}
	}
	return result
}

func (this *Okex) ParsePosition(goArgs ...*Variant) *Variant {
	position := GoGetArg(goArgs, 0, MkUndefined())
	_ = position
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	marketId := this.SafeString(position, MkString("instId"))
	_ = marketId
	market = this.SafeMarket(marketId, market)
	symbol := *(market).At(MkString("symbol"))
	_ = symbol
	contractsString := this.SafeString(position, MkString("pos"))
	_ = contractsString
	contracts := OpCopy(MkUndefined())
	_ = contracts
	if IsTrue(OpNotEq2(contractsString, MkUndefined())) {
		contracts = parseInt(contractsString)
	}
	notionalString := this.SafeString(position, MkString("notionalUsd"))
	_ = notionalString
	notional := this.ParseNumber(notionalString)
	_ = notional
	marginType := this.SafeString(position, MkString("mgnMode"))
	_ = marginType
	initialMarginString := OpCopy(MkUndefined())
	_ = initialMarginString
	entryPriceString := this.SafeString(position, MkString("avgPx"))
	_ = entryPriceString
	unrealizedPnlString := this.SafeString(position, MkString("upl"))
	_ = unrealizedPnlString
	if IsTrue(OpEq2(marginType, MkString("cross"))) {
		initialMarginString = this.SafeString(position, MkString("imr"))
	} else {
		initialMarginString = this.SafeString(position, MkString("margin"))
	}
	maintenanceMarginString := this.SafeString(position, MkString("mmr"))
	_ = maintenanceMarginString
	maintenanceMargin := this.ParseNumber(maintenanceMarginString)
	_ = maintenanceMargin
	initialMarginPercentage := OpCopy(MkUndefined())
	_ = initialMarginPercentage
	maintenanceMarginPercentage := OpCopy(MkUndefined())
	_ = maintenanceMarginPercentage
	if IsTrue(*(market).At(MkString("inverse"))) {
		notionalValue := Precise.StringDiv(Precise.StringMul(contractsString, *(market).At(MkString("contractSize"))), entryPriceString)
		_ = notionalValue
		maintenanceMarginPercentage = Precise.StringDiv(maintenanceMarginString, notionalValue)
		initialMarginPercentage = this.ParseNumber(Precise.StringDiv(initialMarginString, notionalValue, MkInteger(4)))
	} else {
		maintenanceMarginPercentage = Precise.StringDiv(maintenanceMarginString, notionalString)
		initialMarginPercentage = this.ParseNumber(Precise.StringDiv(initialMarginString, notionalString, MkInteger(4)))
	}
	rounder := MkString("0.00005")
	_ = rounder
	maintenanceMarginPercentage = this.ParseNumber(Precise.StringDiv(Precise.StringAdd(maintenanceMarginPercentage, rounder), MkString("1"), MkInteger(4)))
	collateralString := Precise.StringAdd(initialMarginString, unrealizedPnlString)
	_ = collateralString
	liquidationPrice := this.SafeNumber(position, MkString("liqPx"))
	_ = liquidationPrice
	percentageString := this.SafeString(position, MkString("uplRatio"))
	_ = percentageString
	percentage := this.ParseNumber(Precise.StringMul(percentageString, MkString("100")))
	_ = percentage
	side := this.SafeString(position, MkString("posSide"))
	_ = side
	timestamp := this.SafeInteger(position, MkString("uTime"))
	_ = timestamp
	leverage := this.SafeInteger(position, MkString("lever"))
	_ = leverage
	marginRatio := this.ParseNumber(Precise.StringDiv(maintenanceMarginString, collateralString, MkInteger(4)))
	_ = marginRatio
	return MkMap(&VarMap{
		"info":                        position,
		"symbol":                      symbol,
		"notional":                    notional,
		"marginType":                  marginType,
		"liquidationPrice":            liquidationPrice,
		"entryPrice":                  this.ParseNumber(entryPriceString),
		"unrealizedPnl":               this.ParseNumber(unrealizedPnlString),
		"percentage":                  percentage,
		"contracts":                   contracts,
		"contractSize":                this.ParseNumber(*(market).At(MkString("contractSize"))),
		"side":                        side,
		"timestamp":                   timestamp,
		"datetime":                    this.Iso8601(timestamp),
		"maintenanceMargin":           maintenanceMargin,
		"maintenanceMarginPercentage": maintenanceMarginPercentage,
		"collateral":                  this.ParseNumber(collateralString),
		"initialMargin":               this.ParseNumber(initialMarginString),
		"initialMarginPercentage":     this.ParseNumber(initialMarginPercentage),
		"leverage":                    leverage,
		"marginRatio":                 marginRatio,
	})
}

func (this *Okex) Transfer(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	amount := GoGetArg(goArgs, 1, MkUndefined())
	_ = amount
	fromAccount := GoGetArg(goArgs, 2, MkUndefined())
	_ = fromAccount
	toAccount := GoGetArg(goArgs, 3, MkUndefined())
	_ = toAccount
	params := GoGetArg(goArgs, 4, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	accountsByType := this.SafeValue(*this.At(MkString("options")), MkString("accountsByType"), MkMap(&VarMap{}))
	_ = accountsByType
	fromId := this.SafeString(accountsByType, fromAccount, fromAccount)
	_ = fromId
	toId := this.SafeString(accountsByType, toAccount, toAccount)
	_ = toId
	if IsTrue(OpEq2(fromId, MkUndefined())) {
		keys := GoGetKeys(accountsByType)
		_ = keys
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" fromAccount must be one of "), keys.Join(MkString(", "))))))
	}
	if IsTrue(OpEq2(toId, MkUndefined())) {
		keys := GoGetKeys(accountsByType)
		_ = keys
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" toAccount must be one of "), keys.Join(MkString(", "))))))
	}
	request := MkMap(&VarMap{
		"ccy":  *(currency).At(MkString("id")),
		"amt":  this.CurrencyToPrecision(code, amount),
		"type": MkString("0"),
		"from": fromId,
		"to":   toId,
	})
	_ = request
	response := this.Call(MkString("privatePostAssetTransfer"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	rawTransfer := this.SafeValue(data, MkInteger(0), MkMap(&VarMap{}))
	_ = rawTransfer
	return this.ParseTransfer(rawTransfer, currency)
}

func (this *Okex) ParseTransfer(goArgs ...*Variant) *Variant {
	transfer := GoGetArg(goArgs, 0, MkUndefined())
	_ = transfer
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	id := this.SafeString(transfer, MkString("transId"))
	_ = id
	currencyId := this.SafeString(transfer, MkString("ccy"))
	_ = currencyId
	code := this.SafeCurrencyCode(currencyId, currency)
	_ = code
	amount := this.SafeNumber(transfer, MkString("amt"))
	_ = amount
	fromAccountId := this.SafeString(transfer, MkString("from"))
	_ = fromAccountId
	toAccountId := this.SafeString(transfer, MkString("to"))
	_ = toAccountId
	typesByAccount := this.SafeValue(*this.At(MkString("options")), MkString("typesByAccount"), MkMap(&VarMap{}))
	_ = typesByAccount
	fromAccount := this.SafeString(typesByAccount, fromAccountId)
	_ = fromAccount
	toAccount := this.SafeString(typesByAccount, toAccountId)
	_ = toAccount
	timestamp := OpCopy(MkUndefined())
	_ = timestamp
	status := OpCopy(MkUndefined())
	_ = status
	return MkMap(&VarMap{
		"info":        transfer,
		"id":          id,
		"timestamp":   timestamp,
		"datetime":    this.Iso8601(timestamp),
		"currency":    code,
		"amount":      amount,
		"fromAccount": fromAccount,
		"toAccount":   toAccount,
		"status":      status,
	})
}

func (this *Okex) Sign(goArgs ...*Variant) *Variant {
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
	isArray := GoIsArray(params)
	_ = isArray
	request := OpAdd(MkString("/api/"), OpAdd(*this.At(MkString("version")), OpAdd(MkString("/"), this.ImplodeParams(path, params))))
	_ = request
	query := this.Omit(params, this.ExtractParams(path))
	_ = query
	url := OpAdd(this.ImplodeHostname(*(*(*this.At(MkString("urls"))).At(MkString("api"))).At(MkString("rest"))), request)
	_ = url
	if IsTrue(OpEq2(api, MkString("public"))) {
		if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
			OpAddAssign(&url, OpAdd(MkString("?"), this.Urlencode(query)))
		}
	} else {
		if IsTrue(OpEq2(api, MkString("private"))) {
			this.CheckRequiredCredentials()
			timestamp := this.Iso8601(this.Milliseconds())
			_ = timestamp
			headers = MkMap(&VarMap{
				"OK-ACCESS-KEY":        *this.At(MkString("apiKey")),
				"OK-ACCESS-PASSPHRASE": *this.At(MkString("password")),
				"OK-ACCESS-TIMESTAMP":  timestamp,
			})
			auth := OpAdd(timestamp, OpAdd(method, request))
			_ = auth
			if IsTrue(OpEq2(method, MkString("GET"))) {
				if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
					urlencodedQuery := OpAdd(MkString("?"), this.Urlencode(query))
					_ = urlencodedQuery
					OpAddAssign(&url, urlencodedQuery)
					OpAddAssign(&auth, urlencodedQuery)
				}
			} else {
				if IsTrue(OpOr(isArray, *(GoGetKeys(query)).At(MkString("length")))) {
					body = this.Json(query)
					OpAddAssign(&auth, body)
				}
				*(headers).At(MkString("Content-Type")) = MkString("application/json")
			}
			signature := this.Hmac(this.Encode(auth), this.Encode(*this.At(MkString("secret"))), MkString("sha256"), MkString("base64"))
			_ = signature
			*(headers).At(MkString("OK-ACCESS-SIGN")) = OpCopy(signature)
		}
	}
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Okex) HandleErrors(goArgs ...*Variant) *Variant {
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
	code := this.SafeString(response, MkString("code"))
	_ = code
	if IsTrue(OpNotEq2(code, MkString("0"))) {
		feedback := OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))
		_ = feedback
		data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
		_ = data
		for i := MkInteger(0); IsTrue(OpLw(i, data.Length)); OpIncr(&i) {
			error := *(data).At(i)
			_ = error
			errorCode := this.SafeString(error, MkString("sCode"))
			_ = errorCode
			message := this.SafeString(error, MkString("sMsg"))
			_ = message
			this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("exact")), errorCode, feedback)
			this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("broad")), message, feedback)
		}
		this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("exact")), code, feedback)
		panic(NewExchangeError(feedback))
	}
	return MkUndefined()
}
