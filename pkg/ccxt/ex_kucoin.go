package ccxt

import ()

type Kucoin struct {
	*ExchangeBase
}

var _ Exchange = (*Kucoin)(nil)

func init() {
	exchange := &Kucoin{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Kucoin) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("kucoin"),
		"name": MkString("KuCoin"),
		"countries": MkArray(&VarArray{
			MkString("SC"),
		}),
		"rateLimit": MkInteger(334),
		"version":   MkString("v2"),
		"certified": MkBool(false),
		"pro":       MkBool(true),
		"comment":   MkString("Platform 2.0"),
		"has": MkMap(&VarMap{
			"CORS":                 MkBool(false),
			"cancelAllOrders":      MkBool(true),
			"cancelOrder":          MkBool(true),
			"createDepositAddress": MkBool(true),
			"createOrder":          MkBool(true),
			"fetchAccounts":        MkBool(true),
			"fetchBalance":         MkBool(true),
			"fetchClosedOrders":    MkBool(true),
			"fetchCurrencies":      MkBool(true),
			"fetchDepositAddress":  MkBool(true),
			"fetchDeposits":        MkBool(true),
			"fetchFundingFee":      MkBool(true),
			"fetchLedger":          MkBool(true),
			"fetchMarkets":         MkBool(true),
			"fetchMyTrades":        MkBool(true),
			"fetchOHLCV":           MkBool(true),
			"fetchOpenOrders":      MkBool(true),
			"fetchOrder":           MkBool(true),
			"fetchOrderBook":       MkBool(true),
			"fetchStatus":          MkBool(true),
			"fetchTicker":          MkBool(true),
			"fetchTickers":         MkBool(true),
			"fetchTime":            MkBool(true),
			"fetchTrades":          MkBool(true),
			"fetchWithdrawals":     MkBool(true),
			"withdraw":             MkBool(true),
			"transfer":             MkBool(true),
		}),
		"urls": MkMap(&VarMap{
			"logo":     MkString("https://user-images.githubusercontent.com/51840849/87295558-132aaf80-c50e-11ea-9801-a2fb0c57c799.jpg"),
			"referral": MkString("https://www.kucoin.com/?rcode=E5wkqe"),
			"api": MkMap(&VarMap{
				"public":         MkString("https://openapi-v2.kucoin.com"),
				"private":        MkString("https://openapi-v2.kucoin.com"),
				"futuresPrivate": MkString("https://api-futures.kucoin.com"),
				"futuresPublic":  MkString("https://api-futures.kucoin.com"),
			}),
			"test": MkMap(&VarMap{
				"public":         MkString("https://openapi-sandbox.kucoin.com"),
				"private":        MkString("https://openapi-sandbox.kucoin.com"),
				"futuresPrivate": MkString("https://api-sandbox-futures.kucoin.com"),
				"futuresPublic":  MkString("https://api-sandbox-futures.kucoin.com"),
			}),
			"www": MkString("https://www.kucoin.com"),
			"doc": MkArray(&VarArray{
				MkString("https://docs.kucoin.com"),
			}),
		}),
		"requiredCredentials": MkMap(&VarMap{
			"apiKey":   MkBool(true),
			"secret":   MkBool(true),
			"password": MkBool(true),
		}),
		"api": MkMap(&VarMap{
			"public": MkMap(&VarMap{
				"get": MkArray(&VarArray{
					MkString("timestamp"),
					MkString("status"),
					MkString("symbols"),
					MkString("markets"),
					MkString("market/allTickers"),
					MkString("market/orderbook/level{level}_{limit}"),
					MkString("market/orderbook/level2_20"),
					MkString("market/orderbook/level2_100"),
					MkString("market/histories"),
					MkString("market/candles"),
					MkString("market/stats"),
					MkString("currencies"),
					MkString("currencies/{currency}"),
					MkString("prices"),
					MkString("mark-price/{symbol}/current"),
					MkString("margin/config"),
				}),
				"post": MkArray(&VarArray{
					MkString("bullet-public"),
				}),
			}),
			"private": MkMap(&VarMap{
				"get": MkArray(&VarArray{
					MkString("market/orderbook/level{level}"),
					MkString("market/orderbook/level2"),
					MkString("market/orderbook/level3"),
					MkString("accounts"),
					MkString("accounts/{accountId}"),
					MkString("accounts/{accountId}/ledgers"),
					MkString("accounts/{accountId}/holds"),
					MkString("accounts/transferable"),
					MkString("sub/user"),
					MkString("sub-accounts"),
					MkString("sub-accounts/{subUserId}"),
					MkString("deposit-addresses"),
					MkString("deposits"),
					MkString("hist-deposits"),
					MkString("hist-orders"),
					MkString("hist-withdrawals"),
					MkString("withdrawals"),
					MkString("withdrawals/quotas"),
					MkString("orders"),
					MkString("order/client-order/{clientOid}"),
					MkString("orders/{orderId}"),
					MkString("limit/orders"),
					MkString("fills"),
					MkString("limit/fills"),
					MkString("margin/account"),
					MkString("margin/borrow"),
					MkString("margin/borrow/outstanding"),
					MkString("margin/borrow/borrow/repaid"),
					MkString("margin/lend/active"),
					MkString("margin/lend/done"),
					MkString("margin/lend/trade/unsettled"),
					MkString("margin/lend/trade/settled"),
					MkString("margin/lend/assets"),
					MkString("margin/market"),
					MkString("margin/trade/last"),
					MkString("stop-order/{orderId}"),
					MkString("stop-order"),
					MkString("stop-order/queryOrderByClientOid"),
				}),
				"post": MkArray(&VarArray{
					MkString("accounts"),
					MkString("accounts/inner-transfer"),
					MkString("accounts/sub-transfer"),
					MkString("deposit-addresses"),
					MkString("withdrawals"),
					MkString("orders"),
					MkString("orders/multi"),
					MkString("margin/borrow"),
					MkString("margin/order"),
					MkString("margin/repay/all"),
					MkString("margin/repay/single"),
					MkString("margin/lend"),
					MkString("margin/toggle-auto-lend"),
					MkString("bullet-private"),
					MkString("stop-order"),
				}),
				"delete": MkArray(&VarArray{
					MkString("withdrawals/{withdrawalId}"),
					MkString("orders"),
					MkString("orders/client-order/{clientOid}"),
					MkString("orders/{orderId}"),
					MkString("margin/lend/{orderId}"),
					MkString("stop-order/cancelOrderByClientOid"),
					MkString("stop-order/{orderId}"),
					MkString("stop-order/cancel"),
				}),
			}),
			"futuresPublic": MkMap(&VarMap{
				"get": MkArray(&VarArray{
					MkString("contracts/active"),
					MkString("contracts/{symbol}"),
					MkString("ticker"),
					MkString("level2/snapshot"),
					MkString("level2/depth20"),
					MkString("level2/depth100"),
					MkString("level2/message/query"),
					MkString("level3/message/query"),
					MkString("level3/snapshot"),
					MkString("trade/history"),
					MkString("interest/query"),
					MkString("index/query"),
					MkString("mark-price/{symbol}/current"),
					MkString("premium/query"),
					MkString("funding-rate/{symbol}/current"),
					MkString("timestamp"),
					MkString("status"),
					MkString("kline/query"),
				}),
				"post": MkArray(&VarArray{
					MkString("bullet-public"),
				}),
			}),
			"futuresPrivate": MkMap(&VarMap{
				"get": MkArray(&VarArray{
					MkString("account-overview"),
					MkString("transaction-history"),
					MkString("deposit-address"),
					MkString("deposit-list"),
					MkString("withdrawals/quotas"),
					MkString("withdrawal-list"),
					MkString("transfer-list"),
					MkString("orders"),
					MkString("stopOrders"),
					MkString("recentDoneOrders"),
					MkString("orders/{order-id}"),
					MkString("orders/byClientOid"),
					MkString("fills"),
					MkString("recentFills"),
					MkString("openOrderStatistics"),
					MkString("position"),
					MkString("positions"),
					MkString("funding-history"),
				}),
				"post": MkArray(&VarArray{
					MkString("withdrawals"),
					MkString("transfer-out"),
					MkString("orders"),
					MkString("position/margin/auto-deposit-status"),
					MkString("position/margin/deposit-margin"),
					MkString("bullet-private"),
				}),
				"delete": MkArray(&VarArray{
					MkString("withdrawals/{withdrawalId}"),
					MkString("cancel/transfer-out"),
					MkString("orders/{order-id}"),
					MkString("orders"),
					MkString("stopOrders"),
				}),
			}),
		}),
		"timeframes": MkMap(&VarMap{
			"1m":  MkString("1min"),
			"3m":  MkString("3min"),
			"5m":  MkString("5min"),
			"15m": MkString("15min"),
			"30m": MkString("30min"),
			"1h":  MkString("1hour"),
			"2h":  MkString("2hour"),
			"4h":  MkString("4hour"),
			"6h":  MkString("6hour"),
			"8h":  MkString("8hour"),
			"12h": MkString("12hour"),
			"1d":  MkString("1day"),
			"1w":  MkString("1week"),
		}),
		"exceptions": MkMap(&VarMap{
			"exact": MkMap(&VarMap{
				"order not exist":                                          OrderNotFound,
				"order not exist.":                                         OrderNotFound,
				"order_not_exist":                                          OrderNotFound,
				"order_not_exist_or_not_allow_to_cancel":                   InvalidOrder,
				"Order size below the minimum requirement.":                InvalidOrder,
				"The withdrawal amount is below the minimum requirement.":  ExchangeError,
				"Unsuccessful! Exceeded the max. funds out-transfer limit": InsufficientFunds,
				"400":    BadRequest,
				"401":    AuthenticationError,
				"403":    NotSupported,
				"404":    NotSupported,
				"405":    NotSupported,
				"429":    RateLimitExceeded,
				"500":    ExchangeNotAvailable,
				"503":    ExchangeNotAvailable,
				"101030": PermissionDenied,
				"200004": InsufficientFunds,
				"230003": InsufficientFunds,
				"260100": InsufficientFunds,
				"300000": InvalidOrder,
				"400000": BadSymbol,
				"400001": AuthenticationError,
				"400002": InvalidNonce,
				"400003": AuthenticationError,
				"400004": AuthenticationError,
				"400005": AuthenticationError,
				"400006": AuthenticationError,
				"400007": AuthenticationError,
				"400008": NotSupported,
				"400100": BadRequest,
				"411100": AccountSuspended,
				"415000": BadRequest,
				"500000": ExchangeError,
			}),
			"broad": MkMap(&VarMap{
				"Exceeded the access frequency": RateLimitExceeded,
				"require more permission":       PermissionDenied,
			}),
		}),
		"fees": MkMap(&VarMap{
			"trading": MkMap(&VarMap{
				"tierBased":  MkBool(false),
				"percentage": MkBool(true),
				"taker":      MkNumber(0.001),
				"maker":      MkNumber(0.001),
			}),
			"funding": MkMap(&VarMap{
				"tierBased":  MkBool(false),
				"percentage": MkBool(false),
				"withdraw":   MkMap(&VarMap{}),
				"deposit":    MkMap(&VarMap{}),
			}),
		}),
		"commonCurrencies": MkMap(&VarMap{
			"HOT":  MkString("HOTNOW"),
			"EDGE": MkString("DADI"),
			"WAX":  MkString("WAXP"),
			"TRY":  MkString("Trias"),
			"VAI":  MkString("VAIOT"),
		}),
		"options": MkMap(&VarMap{
			"version":             MkString("v1"),
			"symbolSeparator":     MkString("-"),
			"fetchMyTradesMethod": MkString("private_get_fills"),
			"fetchBalance":        MkString("trade"),
			"versions": MkMap(&VarMap{
				"public": MkMap(&VarMap{"GET": MkMap(&VarMap{
					"status":                                MkString("v1"),
					"market/orderbook/level2_20":            MkString("v1"),
					"market/orderbook/level2_100":           MkString("v1"),
					"market/orderbook/level{level}_{limit}": MkString("v1"),
				})}),
				"private": MkMap(&VarMap{
					"GET": MkMap(&VarMap{
						"market/orderbook/level2":       MkString("v3"),
						"market/orderbook/level3":       MkString("v3"),
						"market/orderbook/level{level}": MkString("v3"),
					}),
					"POST": MkMap(&VarMap{
						"accounts/inner-transfer": MkString("v2"),
						"accounts/sub-transfer":   MkString("v2"),
					}),
				}),
				"futuresPrivate": MkMap(&VarMap{
					"GET": MkMap(&VarMap{
						"account-overview": MkString("v1"),
						"positions":        MkString("v1"),
					}),
					"POST": MkMap(&VarMap{"transfer-out": MkString("v2")}),
				}),
				"futuresPublic": MkMap(&VarMap{"GET": MkMap(&VarMap{"level3/snapshot": MkString("v2")})}),
			}),
			"accountsByType": MkMap(&VarMap{
				"trade":    MkString("trade"),
				"trading":  MkString("trade"),
				"spot":     MkString("trade"),
				"margin":   MkString("margin"),
				"main":     MkString("main"),
				"funding":  MkString("main"),
				"futures":  MkString("contract"),
				"contract": MkString("contract"),
				"pool":     MkString("pool"),
				"pool-x":   MkString("pool"),
			}),
		}),
	}))
}

func (this *Kucoin) Nonce(goArgs ...*Variant) *Variant {
	return this.Milliseconds()
}

func (this *Kucoin) LoadTimeDifference(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetTimestamp"), params)
	_ = response
	after := this.Milliseconds()
	_ = after
	kucoinTime := this.SafeInteger(response, MkString("data"))
	_ = kucoinTime
	*(*this.At(MkString("options"))).At(MkString("timeDifference")) = parseInt(OpSub(after, kucoinTime))
	return *(*this.At(MkString("options"))).At(MkString("timeDifference"))
}

func (this *Kucoin) FetchTime(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetTimestamp"), params)
	_ = response
	return this.SafeInteger(response, MkString("data"))
}

func (this *Kucoin) FetchStatus(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetStatus"), params)
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	status := this.SafeValue(data, MkString("status"))
	_ = status
	if IsTrue(OpNotEq2(status, MkUndefined())) {
		status = OpTrinary(OpEq2(status, MkString("open")), MkString("ok"), MkString("maintenance"))
		*this.At(MkString("status")) = this.Extend(*this.At(MkString("status")), MkMap(&VarMap{
			"status":  status,
			"updated": this.Milliseconds(),
		}))
	}
	return *this.At(MkString("status"))
}

func (this *Kucoin) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetSymbols"), params)
	_ = response
	data := *(response).At(MkString("data"))
	_ = data
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, data.Length)); OpIncr(&i) {
		market := *(data).At(i)
		_ = market
		id := this.SafeString(market, MkString("symbol"))
		_ = id
		baseId := MkUndefined()
		_ = baseId
		quoteId := MkUndefined()
		_ = quoteId
		ArrayUnpack(id.Split(MkString("-")), &baseId, &quoteId)
		base := this.SafeCurrencyCode(baseId)
		_ = base
		quote := this.SafeCurrencyCode(quoteId)
		_ = quote
		symbol := OpAdd(base, OpAdd(MkString("/"), quote))
		_ = symbol
		active := this.SafeValue(market, MkString("enableTrading"))
		_ = active
		baseMaxSize := this.SafeNumber(market, MkString("baseMaxSize"))
		_ = baseMaxSize
		baseMinSizeString := this.SafeString(market, MkString("baseMinSize"))
		_ = baseMinSizeString
		quoteMaxSizeString := this.SafeString(market, MkString("quoteMaxSize"))
		_ = quoteMaxSizeString
		baseMinSize := this.ParseNumber(baseMinSizeString)
		_ = baseMinSize
		quoteMaxSize := this.ParseNumber(quoteMaxSizeString)
		_ = quoteMaxSize
		quoteMinSize := this.SafeNumber(market, MkString("quoteMinSize"))
		_ = quoteMinSize
		precision := MkMap(&VarMap{
			"amount": this.PrecisionFromString(this.SafeString(market, MkString("baseIncrement"))),
			"price":  this.PrecisionFromString(this.SafeString(market, MkString("priceIncrement"))),
		})
		_ = precision
		limits := MkMap(&VarMap{
			"amount": MkMap(&VarMap{
				"min": baseMinSize,
				"max": baseMaxSize,
			}),
			"price": MkMap(&VarMap{
				"min": this.SafeNumber(market, MkString("priceIncrement")),
				"max": this.ParseNumber(Precise.StringDiv(quoteMaxSizeString, baseMinSizeString)),
			}),
			"cost": MkMap(&VarMap{
				"min": quoteMinSize,
				"max": quoteMaxSize,
			}),
		})
		_ = limits
		result.Push(MkMap(&VarMap{
			"id":        id,
			"symbol":    symbol,
			"baseId":    baseId,
			"quoteId":   quoteId,
			"base":      base,
			"quote":     quote,
			"active":    active,
			"precision": precision,
			"limits":    limits,
			"info":      market,
		}))
	}
	return result
}

func (this *Kucoin) FetchCurrencies(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicGetCurrencies"), params)
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	result := MkMap(&VarMap{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, data.Length)); OpIncr(&i) {
		entry := *(data).At(i)
		_ = entry
		id := this.SafeString(entry, MkString("currency"))
		_ = id
		name := this.SafeString(entry, MkString("fullName"))
		_ = name
		code := this.SafeCurrencyCode(id)
		_ = code
		precision := this.SafeInteger(entry, MkString("precision"))
		_ = precision
		isWithdrawEnabled := this.SafeValue(entry, MkString("isWithdrawEnabled"), MkBool(false))
		_ = isWithdrawEnabled
		isDepositEnabled := this.SafeValue(entry, MkString("isDepositEnabled"), MkBool(false))
		_ = isDepositEnabled
		fee := this.SafeNumber(entry, MkString("withdrawalMinFee"))
		_ = fee
		active := OpAnd(isWithdrawEnabled, isDepositEnabled)
		_ = active
		*(result).At(code) = MkMap(&VarMap{
			"id":        id,
			"name":      name,
			"code":      code,
			"precision": precision,
			"info":      entry,
			"active":    active,
			"fee":       fee,
			"limits":    *this.At(MkString("limits")),
		})
	}
	return result
}

func (this *Kucoin) FetchAccounts(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("privateGetAccounts"), params)
	_ = response
	data := this.SafeValue(response, MkString("data"))
	_ = data
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, data.Length)); OpIncr(&i) {
		account := *(data).At(i)
		_ = account
		accountId := this.SafeString(account, MkString("id"))
		_ = accountId
		currencyId := this.SafeString(account, MkString("currency"))
		_ = currencyId
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		type_ := this.SafeString(account, MkString("type"))
		_ = type_
		result.Push(MkMap(&VarMap{
			"id":       accountId,
			"type":     type_,
			"currency": code,
			"info":     account,
		}))
	}
	return result
}

func (this *Kucoin) FetchFundingFee(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{"currency": *(currency).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("privateGetWithdrawalsQuotas"), this.Extend(request, params))
	_ = response
	data := *(response).At(MkString("data"))
	_ = data
	withdrawFees := MkMap(&VarMap{})
	_ = withdrawFees
	*(withdrawFees).At(code) = this.SafeNumber(data, MkString("withdrawMinFee"))
	return MkMap(&VarMap{
		"info":     response,
		"withdraw": withdrawFees,
		"deposit":  MkMap(&VarMap{}),
	})
}

func (this *Kucoin) ParseTicker(goArgs ...*Variant) *Variant {
	ticker := GoGetArg(goArgs, 0, MkUndefined())
	_ = ticker
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	percentage := this.SafeNumber(ticker, MkString("changeRate"))
	_ = percentage
	if IsTrue(OpNotEq2(percentage, MkUndefined())) {
		percentage = OpMulti(percentage, MkInteger(100))
	}
	last := this.SafeNumber2(ticker, MkString("last"), MkString("lastTradedPrice"))
	_ = last
	marketId := this.SafeString(ticker, MkString("symbol"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market, MkString("-"))
	_ = symbol
	baseVolume := this.SafeNumber(ticker, MkString("vol"))
	_ = baseVolume
	quoteVolume := this.SafeNumber(ticker, MkString("volValue"))
	_ = quoteVolume
	vwap := this.Vwap(baseVolume, quoteVolume)
	_ = vwap
	timestamp := this.SafeInteger2(ticker, MkString("time"), MkString("datetime"))
	_ = timestamp
	return MkMap(&VarMap{
		"symbol":        symbol,
		"timestamp":     timestamp,
		"datetime":      this.Iso8601(timestamp),
		"high":          this.SafeNumber(ticker, MkString("high")),
		"low":           this.SafeNumber(ticker, MkString("low")),
		"bid":           this.SafeNumber(ticker, MkString("buy")),
		"bidVolume":     MkUndefined(),
		"ask":           this.SafeNumber(ticker, MkString("sell")),
		"askVolume":     MkUndefined(),
		"vwap":          vwap,
		"open":          this.SafeNumber(ticker, MkString("open")),
		"close":         last,
		"last":          last,
		"previousClose": MkUndefined(),
		"change":        this.SafeNumber(ticker, MkString("changePrice")),
		"percentage":    percentage,
		"average":       this.SafeNumber(ticker, MkString("averagePrice")),
		"baseVolume":    baseVolume,
		"quoteVolume":   quoteVolume,
		"info":          ticker,
	})
}

func (this *Kucoin) FetchTickers(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("publicGetMarketAllTickers"), params)
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	tickers := this.SafeValue(data, MkString("ticker"), MkArray(&VarArray{}))
	_ = tickers
	result := MkMap(&VarMap{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, tickers.Length)); OpIncr(&i) {
		ticker := this.ParseTicker(*(tickers).At(i))
		_ = ticker
		symbol := this.SafeString(ticker, MkString("symbol"))
		_ = symbol
		if IsTrue(OpNotEq2(symbol, MkUndefined())) {
			*(result).At(symbol) = OpCopy(ticker)
		}
	}
	return this.FilterByArray(result, MkString("symbol"), symbols)
}

func (this *Kucoin) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("publicGetMarketStats"), this.Extend(request, params))
	_ = response
	return this.ParseTicker(*(response).At(MkString("data")), market)
}

func (this *Kucoin) ParseOHLCV(goArgs ...*Variant) *Variant {
	ohlcv := GoGetArg(goArgs, 0, MkUndefined())
	_ = ohlcv
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	return MkArray(&VarArray{
		this.SafeTimestamp(ohlcv, MkInteger(0)),
		this.SafeNumber(ohlcv, MkInteger(1)),
		this.SafeNumber(ohlcv, MkInteger(3)),
		this.SafeNumber(ohlcv, MkInteger(4)),
		this.SafeNumber(ohlcv, MkInteger(2)),
		this.SafeNumber(ohlcv, MkInteger(5)),
	})
}

func (this *Kucoin) FetchOHLCV(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	timeframe := GoGetArg(goArgs, 1, MkString("15m"))
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
	marketId := *(market).At(MkString("id"))
	_ = marketId
	request := MkMap(&VarMap{
		"symbol": marketId,
		"type":   *(*this.At(MkString("timeframes"))).At(timeframe),
	})
	_ = request
	duration := OpMulti(this.ParseTimeframe(timeframe), MkInteger(1000))
	_ = duration
	endAt := this.Milliseconds()
	_ = endAt
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("startAt")) = parseInt(Math.Floor(OpDiv(since, MkInteger(1000))))
		if IsTrue(OpEq2(limit, MkUndefined())) {
			limit = this.SafeInteger(*this.At(MkString("options")), MkString("fetchOHLCVLimit"), MkInteger(1500))
		}
		endAt = this.Sum(since, OpMulti(limit, duration))
	} else {
		if IsTrue(OpNotEq2(limit, MkUndefined())) {
			since = OpSub(endAt, OpMulti(limit, duration))
			*(request).At(MkString("startAt")) = parseInt(Math.Floor(OpDiv(since, MkInteger(1000))))
		}
	}
	*(request).At(MkString("endAt")) = parseInt(Math.Floor(OpDiv(endAt, MkInteger(1000))))
	response := this.Call(MkString("publicGetMarketCandles"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseOHLCVs(data, market, timeframe, since, limit)
}

func (this *Kucoin) CreateDepositAddress(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{"currency": *(currency).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("privatePostDepositAddresses"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	address := this.SafeString(data, MkString("address"))
	_ = address
	if IsTrue(OpNotEq2(address, MkUndefined())) {
		address = address.Replace(MkString("bitcoincash:"), MkString(""))
	}
	tag := this.SafeString(data, MkString("memo"))
	_ = tag
	if IsTrue(OpNotEq2(code, MkString("NIM"))) {
		this.CheckAddress(address)
	}
	return MkMap(&VarMap{
		"info":     response,
		"currency": code,
		"address":  address,
		"tag":      tag,
	})
}

func (this *Kucoin) FetchDepositAddress(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{"currency": *(currency).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("privateGetDepositAddresses"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	address := this.SafeString(data, MkString("address"))
	_ = address
	tag := this.SafeString(data, MkString("memo"))
	_ = tag
	if IsTrue(OpNotEq2(code, MkString("NIM"))) {
		this.CheckAddress(address)
	}
	return MkMap(&VarMap{
		"info":     response,
		"currency": code,
		"address":  address,
		"tag":      tag,
	})
}

func (this *Kucoin) FetchL3OrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	return this.FetchOrderBook(symbol, limit, MkMap(&VarMap{"level": MkInteger(3)}))
}

func (this *Kucoin) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	marketId := this.MarketId(symbol)
	_ = marketId
	level := this.SafeInteger(params, MkString("level"), MkInteger(2))
	_ = level
	request := MkMap(&VarMap{
		"symbol": marketId,
		"level":  level,
	})
	_ = request
	method := MkString("privateGetMarketOrderbookLevelLevel")
	_ = method
	if IsTrue(OpEq2(level, MkInteger(2))) {
		if IsTrue(OpNotEq2(limit, MkUndefined())) {
			if IsTrue(OpOr(OpEq2(limit, MkInteger(20)), OpEq2(limit, MkInteger(100)))) {
				*(request).At(MkString("limit")) = OpCopy(limit)
				method = MkString("publicGetMarketOrderbookLevelLevelLimit")
			} else {
				panic(NewExchangeError(OpAdd(*this.At(MkString("id")), MkString(" fetchOrderBook limit argument must be undefined, 20 or 100"))))
			}
		}
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	timestamp := this.SafeInteger(data, MkString("time"))
	_ = timestamp
	orderbook := this.ParseOrderBook(data, symbol, timestamp, MkString("bids"), MkString("asks"), OpSub(level, MkInteger(2)), OpSub(level, MkInteger(1)))
	_ = orderbook
	*(orderbook).At(MkString("nonce")) = this.SafeInteger(data, MkString("sequence"))
	return orderbook
}

func (this *Kucoin) CreateOrder(goArgs ...*Variant) *Variant {
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
	marketId := this.MarketId(symbol)
	_ = marketId
	clientOrderId := this.SafeString2(params, MkString("clientOid"), MkString("clientOrderId"), this.Uuid())
	_ = clientOrderId
	params = this.Omit(params, MkArray(&VarArray{
		MkString("clientOid"),
		MkString("clientOrderId"),
	}))
	request := MkMap(&VarMap{
		"clientOid": clientOrderId,
		"side":      side,
		"symbol":    marketId,
		"type":      type_,
	})
	_ = request
	quoteAmount := this.SafeNumber2(params, MkString("cost"), MkString("funds"))
	_ = quoteAmount
	amountString := OpCopy(MkUndefined())
	_ = amountString
	costString := OpCopy(MkUndefined())
	_ = costString
	if IsTrue(OpEq2(type_, MkString("market"))) {
		if IsTrue(OpNotEq2(quoteAmount, MkUndefined())) {
			params = this.Omit(params, MkArray(&VarArray{
				MkString("cost"),
				MkString("funds"),
			}))
			costString = this.AmountToPrecision(symbol, quoteAmount)
			*(request).At(MkString("funds")) = OpCopy(costString)
		} else {
			amountString = this.AmountToPrecision(symbol, amount)
			*(request).At(MkString("size")) = this.AmountToPrecision(symbol, amount)
		}
	} else {
		amountString = this.AmountToPrecision(symbol, amount)
		*(request).At(MkString("size")) = OpCopy(amountString)
		*(request).At(MkString("price")) = this.PriceToPrecision(symbol, price)
	}
	response := this.Call(MkString("privatePostOrders"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	timestamp := this.Milliseconds()
	_ = timestamp
	id := this.SafeString(data, MkString("orderId"))
	_ = id
	order := MkMap(&VarMap{
		"id":                 id,
		"clientOrderId":      clientOrderId,
		"info":               data,
		"timestamp":          timestamp,
		"datetime":           this.Iso8601(timestamp),
		"lastTradeTimestamp": MkUndefined(),
		"symbol":             symbol,
		"type":               type_,
		"side":               side,
		"price":              price,
		"amount":             this.ParseNumber(amountString),
		"cost":               this.ParseNumber(costString),
		"average":            MkUndefined(),
		"filled":             MkUndefined(),
		"remaining":          MkUndefined(),
		"status":             MkUndefined(),
		"fee":                MkUndefined(),
		"trades":             MkUndefined(),
	})
	_ = order
	return order
}

func (this *Kucoin) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{})
	_ = request
	clientOrderId := this.SafeString2(params, MkString("clientOid"), MkString("clientOrderId"))
	_ = clientOrderId
	method := MkString("privateDeleteOrdersOrderId")
	_ = method
	if IsTrue(OpNotEq2(clientOrderId, MkUndefined())) {
		*(request).At(MkString("clientOid")) = OpCopy(clientOrderId)
		method = MkString("privateDeleteOrdersClientOrderClientOid")
	} else {
		*(request).At(MkString("orderId")) = OpCopy(id)
	}
	params = this.Omit(params, MkArray(&VarArray{
		MkString("clientOid"),
		MkString("clientOrderId"),
	}))
	return this.Call(method, this.Extend(request, params))
}

func (this *Kucoin) CancelAllOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
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
	return this.Call(MkString("privateDeleteOrders"), this.Extend(request, params))
}

func (this *Kucoin) FetchOrdersByStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 2, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 3, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 4, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{"status": status})
	_ = request
	market := OpCopy(MkUndefined())
	_ = market
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
		*(request).At(MkString("symbol")) = *(market).At(MkString("id"))
	}
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("startAt")) = OpCopy(since)
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("pageSize")) = OpCopy(limit)
	}
	response := this.Call(MkString("privateGetOrders"), this.Extend(request, params))
	_ = response
	responseData := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = responseData
	orders := this.SafeValue(responseData, MkString("items"), MkArray(&VarArray{}))
	_ = orders
	return this.ParseOrders(orders, market, since, limit)
}

func (this *Kucoin) FetchClosedOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	return this.FetchOrdersByStatus(MkString("done"), symbol, since, limit, params)
}

func (this *Kucoin) FetchOpenOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	return this.FetchOrdersByStatus(MkString("active"), symbol, since, limit, params)
}

func (this *Kucoin) FetchOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	request := MkMap(&VarMap{})
	_ = request
	clientOrderId := this.SafeString2(params, MkString("clientOid"), MkString("clientOrderId"))
	_ = clientOrderId
	method := MkString("privateGetOrdersOrderId")
	_ = method
	if IsTrue(OpNotEq2(clientOrderId, MkUndefined())) {
		*(request).At(MkString("clientOid")) = OpCopy(clientOrderId)
		method = MkString("privateGetOrdersClientOrderClientOid")
	} else {
		if IsTrue(OpEq2(id, MkUndefined())) {
			panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")), MkString(" fetchOrder() requires an order id"))))
		}
		*(request).At(MkString("orderId")) = OpCopy(id)
	}
	params = this.Omit(params, MkArray(&VarArray{
		MkString("clientOid"),
		MkString("clientOrderId"),
	}))
	response := this.Call(method, this.Extend(request, params))
	_ = response
	market := OpCopy(MkUndefined())
	_ = market
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
	}
	responseData := this.SafeValue(response, MkString("data"))
	_ = responseData
	return this.ParseOrder(responseData, market)
}

func (this *Kucoin) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	marketId := this.SafeString(order, MkString("symbol"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market, MkString("-"))
	_ = symbol
	orderId := this.SafeString(order, MkString("id"))
	_ = orderId
	type_ := this.SafeString(order, MkString("type"))
	_ = type_
	timestamp := this.SafeInteger(order, MkString("createdAt"))
	_ = timestamp
	datetime := this.Iso8601(timestamp)
	_ = datetime
	price := this.SafeNumber(order, MkString("price"))
	_ = price
	if IsTrue(OpEq2(price, MkNumber(0.0))) {
		price = OpCopy(MkUndefined())
	}
	side := this.SafeString(order, MkString("side"))
	_ = side
	feeCurrencyId := this.SafeString(order, MkString("feeCurrency"))
	_ = feeCurrencyId
	feeCurrency := this.SafeCurrencyCode(feeCurrencyId)
	_ = feeCurrency
	feeCost := this.SafeNumber(order, MkString("fee"))
	_ = feeCost
	amount := this.SafeNumber(order, MkString("size"))
	_ = amount
	filled := this.SafeNumber(order, MkString("dealSize"))
	_ = filled
	cost := this.SafeNumber(order, MkString("dealFunds"))
	_ = cost
	isActive := this.SafeValue(order, MkString("isActive"), MkBool(false))
	_ = isActive
	cancelExist := this.SafeValue(order, MkString("cancelExist"), MkBool(false))
	_ = cancelExist
	status := OpTrinary(isActive, MkString("open"), MkString("closed"))
	_ = status
	status = OpTrinary(cancelExist, MkString("canceled"), status)
	fee := MkMap(&VarMap{
		"currency": feeCurrency,
		"cost":     feeCost,
	})
	_ = fee
	clientOrderId := this.SafeString(order, MkString("clientOid"))
	_ = clientOrderId
	timeInForce := this.SafeString(order, MkString("timeInForce"))
	_ = timeInForce
	stopPrice := this.SafeNumber(order, MkString("stopPrice"))
	_ = stopPrice
	postOnly := this.SafeValue(order, MkString("postOnly"))
	_ = postOnly
	return this.SafeOrder(MkMap(&VarMap{
		"id":                 orderId,
		"clientOrderId":      clientOrderId,
		"symbol":             symbol,
		"type":               type_,
		"timeInForce":        timeInForce,
		"postOnly":           postOnly,
		"side":               side,
		"amount":             amount,
		"price":              price,
		"stopPrice":          stopPrice,
		"cost":               cost,
		"filled":             filled,
		"remaining":          MkUndefined(),
		"timestamp":          timestamp,
		"datetime":           datetime,
		"fee":                fee,
		"status":             status,
		"info":               order,
		"lastTradeTimestamp": MkUndefined(),
		"average":            MkUndefined(),
		"trades":             MkUndefined(),
	}))
}

func (this *Kucoin) FetchMyTrades(goArgs ...*Variant) *Variant {
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
		*(request).At(MkString("pageSize")) = OpCopy(limit)
	}
	method := *(*this.At(MkString("options"))).At(MkString("fetchMyTradesMethod"))
	_ = method
	parseResponseData := OpCopy(MkBool(false))
	_ = parseResponseData
	if IsTrue(OpEq2(method, MkString("private_get_fills"))) {
		if IsTrue(OpNotEq2(since, MkUndefined())) {
			*(request).At(MkString("startAt")) = OpCopy(since)
		}
	} else {
		if IsTrue(OpEq2(method, MkString("private_get_limit_fills"))) {
			parseResponseData = OpCopy(MkBool(true))
		} else {
			if IsTrue(OpEq2(method, MkString("private_get_hist_orders"))) {
				if IsTrue(OpNotEq2(since, MkUndefined())) {
					*(request).At(MkString("startAt")) = parseInt(OpDiv(since, MkInteger(1000)))
				}
			} else {
				panic(NewExchangeError(OpAdd(*this.At(MkString("id")), MkString(" invalid fetchClosedOrder method"))))
			}
		}
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	trades := OpCopy(MkUndefined())
	_ = trades
	if IsTrue(parseResponseData) {
		trades = OpCopy(data)
	} else {
		trades = this.SafeValue(data, MkString("items"), MkArray(&VarArray{}))
	}
	return this.ParseTrades(trades, market, since, limit)
}

func (this *Kucoin) FetchTrades(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("startAt")) = Math.Floor(OpDiv(since, MkInteger(1000)))
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("pageSize")) = OpCopy(limit)
	}
	response := this.Call(MkString("publicGetMarketHistories"), this.Extend(request, params))
	_ = response
	trades := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = trades
	return this.ParseTrades(trades, market, since, limit)
}

func (this *Kucoin) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	marketId := this.SafeString(trade, MkString("symbol"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market, MkString("-"))
	_ = symbol
	id := this.SafeString2(trade, MkString("tradeId"), MkString("id"))
	_ = id
	orderId := this.SafeString(trade, MkString("orderId"))
	_ = orderId
	takerOrMaker := this.SafeString(trade, MkString("liquidity"))
	_ = takerOrMaker
	timestamp := this.SafeInteger(trade, MkString("time"))
	_ = timestamp
	if IsTrue(OpNotEq2(timestamp, MkUndefined())) {
		timestamp = parseInt(OpDiv(timestamp, MkInteger(1000000)))
	} else {
		timestamp = this.SafeInteger(trade, MkString("createdAt"))
		if IsTrue(OpAnd(OpHasMember(MkString("dealValue"), trade), OpNotEq2(timestamp, MkUndefined()))) {
			timestamp = OpMulti(timestamp, MkInteger(1000))
		}
	}
	priceString := this.SafeString2(trade, MkString("price"), MkString("dealPrice"))
	_ = priceString
	amountString := this.SafeString2(trade, MkString("size"), MkString("amount"))
	_ = amountString
	price := this.ParseNumber(priceString)
	_ = price
	amount := this.ParseNumber(amountString)
	_ = amount
	side := this.SafeString(trade, MkString("side"))
	_ = side
	fee := OpCopy(MkUndefined())
	_ = fee
	feeCost := this.SafeNumber(trade, MkString("fee"))
	_ = feeCost
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		feeCurrencyId := this.SafeString(trade, MkString("feeCurrency"))
		_ = feeCurrencyId
		feeCurrency := this.SafeCurrencyCode(feeCurrencyId)
		_ = feeCurrency
		if IsTrue(OpEq2(feeCurrency, MkUndefined())) {
			if IsTrue(OpNotEq2(market, MkUndefined())) {
				feeCurrency = OpTrinary(OpEq2(side, MkString("sell")), *(market).At(MkString("quote")), *(market).At(MkString("base")))
			}
		}
		fee = MkMap(&VarMap{
			"cost":     feeCost,
			"currency": feeCurrency,
			"rate":     this.SafeNumber(trade, MkString("feeRate")),
		})
	}
	type_ := this.SafeString(trade, MkString("type"))
	_ = type_
	if IsTrue(OpEq2(type_, MkString("match"))) {
		type_ = OpCopy(MkUndefined())
	}
	cost := this.SafeNumber2(trade, MkString("funds"), MkString("dealValue"))
	_ = cost
	if IsTrue(OpEq2(cost, MkUndefined())) {
		cost = this.ParseNumber(Precise.StringMul(priceString, amountString))
	}
	return MkMap(&VarMap{
		"info":         trade,
		"id":           id,
		"order":        orderId,
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       symbol,
		"type":         type_,
		"takerOrMaker": takerOrMaker,
		"side":         side,
		"price":        price,
		"amount":       amount,
		"cost":         cost,
		"fee":          fee,
	})
}

func (this *Kucoin) Withdraw(goArgs ...*Variant) *Variant {
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
	this.CheckAddress(address)
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{
		"currency": *(currency).At(MkString("id")),
		"address":  address,
		"amount":   amount,
	})
	_ = request
	if IsTrue(OpNotEq2(tag, MkUndefined())) {
		*(request).At(MkString("memo")) = OpCopy(tag)
	}
	response := this.Call(MkString("privatePostWithdrawals"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	return MkMap(&VarMap{
		"id":   this.SafeString(data, MkString("withdrawalId")),
		"info": response,
	})
}

func (this *Kucoin) ParseTransactionStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"SUCCESS":    MkString("ok"),
		"PROCESSING": MkString("ok"),
		"FAILURE":    MkString("failed"),
	})
	_ = statuses
	return this.SafeString(statuses, status)
}

func (this *Kucoin) ParseTransaction(goArgs ...*Variant) *Variant {
	transaction := GoGetArg(goArgs, 0, MkUndefined())
	_ = transaction
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	currencyId := this.SafeString(transaction, MkString("currency"))
	_ = currencyId
	code := this.SafeCurrencyCode(currencyId, currency)
	_ = code
	address := this.SafeString(transaction, MkString("address"))
	_ = address
	amount := this.SafeNumber(transaction, MkString("amount"))
	_ = amount
	txid := this.SafeString(transaction, MkString("walletTxId"))
	_ = txid
	if IsTrue(OpNotEq2(txid, MkUndefined())) {
		txidParts := txid.Split(MkString("@"))
		_ = txidParts
		numTxidParts := OpCopy(txidParts.Length)
		_ = numTxidParts
		if IsTrue(OpGt(numTxidParts, MkInteger(1))) {
			if IsTrue(OpEq2(address, MkUndefined())) {
				if IsTrue(OpGt(*(*(txidParts).At(MkInteger(1))).At(MkString("length")), MkInteger(1))) {
					address = *(txidParts).At(MkInteger(1))
				}
			}
		}
		txid = *(txidParts).At(MkInteger(0))
	}
	type_ := OpTrinary(OpEq2(txid, MkUndefined()), MkString("withdrawal"), MkString("deposit"))
	_ = type_
	rawStatus := this.SafeString(transaction, MkString("status"))
	_ = rawStatus
	status := this.ParseTransactionStatus(rawStatus)
	_ = status
	fee := OpCopy(MkUndefined())
	_ = fee
	feeCost := this.SafeNumber(transaction, MkString("fee"))
	_ = feeCost
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		rate := OpCopy(MkUndefined())
		_ = rate
		if IsTrue(OpNotEq2(amount, MkUndefined())) {
			rate = OpDiv(feeCost, amount)
		}
		fee = MkMap(&VarMap{
			"cost":     feeCost,
			"rate":     rate,
			"currency": code,
		})
	}
	tag := this.SafeString(transaction, MkString("memo"))
	_ = tag
	timestamp := this.SafeInteger2(transaction, MkString("createdAt"), MkString("createAt"))
	_ = timestamp
	id := this.SafeString(transaction, MkString("id"))
	_ = id
	updated := this.SafeInteger(transaction, MkString("updatedAt"))
	_ = updated
	isV1 := OpNot(OpHasMember(MkString("createdAt"), transaction))
	_ = isV1
	if IsTrue(isV1) {
		type_ = OpTrinary(OpHasMember(MkString("address"), transaction), MkString("withdrawal"), MkString("deposit"))
		if IsTrue(OpNotEq2(timestamp, MkUndefined())) {
			timestamp = OpMulti(timestamp, MkInteger(1000))
		}
		if IsTrue(OpNotEq2(updated, MkUndefined())) {
			updated = OpMulti(updated, MkInteger(1000))
		}
	}
	comment := this.SafeString(transaction, MkString("remark"))
	_ = comment
	return MkMap(&VarMap{
		"id":          id,
		"info":        transaction,
		"timestamp":   timestamp,
		"datetime":    this.Iso8601(timestamp),
		"address":     address,
		"addressTo":   address,
		"addressFrom": MkUndefined(),
		"tag":         tag,
		"tagTo":       tag,
		"tagFrom":     MkUndefined(),
		"currency":    code,
		"amount":      amount,
		"txid":        txid,
		"type":        type_,
		"status":      status,
		"comment":     comment,
		"fee":         fee,
		"updated":     updated,
	})
}

func (this *Kucoin) FetchDeposits(goArgs ...*Variant) *Variant {
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
		*(request).At(MkString("pageSize")) = OpCopy(limit)
	}
	method := MkString("privateGetDeposits")
	_ = method
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		if IsTrue(OpLw(since, MkInteger(1550448000000))) {
			*(request).At(MkString("startAt")) = parseInt(OpDiv(since, MkInteger(1000)))
			method = MkString("privateGetHistDeposits")
		} else {
			*(request).At(MkString("startAt")) = OpCopy(since)
		}
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	responseData := *(*(response).At(MkString("data"))).At(MkString("items"))
	_ = responseData
	return this.ParseTransactions(responseData, currency, since, limit, MkMap(&VarMap{"type": MkString("deposit")}))
}

func (this *Kucoin) FetchWithdrawals(goArgs ...*Variant) *Variant {
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
		*(request).At(MkString("pageSize")) = OpCopy(limit)
	}
	method := MkString("privateGetWithdrawals")
	_ = method
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		if IsTrue(OpLw(since, MkInteger(1550448000000))) {
			*(request).At(MkString("startAt")) = parseInt(OpDiv(since, MkInteger(1000)))
			method = MkString("privateGetHistWithdrawals")
		} else {
			*(request).At(MkString("startAt")) = OpCopy(since)
		}
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	responseData := *(*(response).At(MkString("data"))).At(MkString("items"))
	_ = responseData
	return this.ParseTransactions(responseData, currency, since, limit, MkMap(&VarMap{"type": MkString("withdrawal")}))
}

func (this *Kucoin) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	defaultType := this.SafeString2(*this.At(MkString("options")), MkString("fetchBalance"), MkString("defaultType"), MkString("trade"))
	_ = defaultType
	requestedType := this.SafeString(params, MkString("type"), defaultType)
	_ = requestedType
	accountsByType := this.SafeValue(*this.At(MkString("options")), MkString("accountsByType"))
	_ = accountsByType
	type_ := this.SafeString(accountsByType, requestedType)
	_ = type_
	if IsTrue(OpEq2(type_, MkUndefined())) {
		keys := GoGetKeys(accountsByType)
		_ = keys
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" type must be one of "), keys.Join(MkString(", "))))))
	}
	params = this.Omit(params, MkString("type"))
	if IsTrue(OpOr(OpEq2(type_, MkString("contract")), OpEq2(type_, MkString("futures")))) {
		response := this.Call(MkString("futuresPrivateGetAccountOverview"), params)
		_ = response
		result := MkMap(&VarMap{
			"info":      response,
			"timestamp": MkUndefined(),
			"datetime":  MkUndefined(),
		})
		_ = result
		data := this.SafeValue(response, MkString("data"))
		_ = data
		currencyId := this.SafeString(data, MkString("currency"))
		_ = currencyId
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		account := this.Account()
		_ = account
		*(account).At(MkString("free")) = this.SafeString(data, MkString("availableBalance"))
		*(account).At(MkString("total")) = this.SafeString(data, MkString("accountEquity"))
		*(result).At(code) = OpCopy(account)
		return this.ParseBalance(result)
	} else {
		request := MkMap(&VarMap{"type": type_})
		_ = request
		response := this.Call(MkString("privateGetAccounts"), this.Extend(request, params))
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
			balanceType := this.SafeString(balance, MkString("type"))
			_ = balanceType
			if IsTrue(OpEq2(balanceType, type_)) {
				currencyId := this.SafeString(balance, MkString("currency"))
				_ = currencyId
				code := this.SafeCurrencyCode(currencyId)
				_ = code
				account := this.Account()
				_ = account
				*(account).At(MkString("total")) = this.SafeString(balance, MkString("balance"))
				*(account).At(MkString("free")) = this.SafeString(balance, MkString("available"))
				*(account).At(MkString("used")) = this.SafeString(balance, MkString("holds"))
				*(result).At(code) = OpCopy(account)
			}
		}
		return this.ParseBalance(result)
	}
	return MkUndefined()
}

func (this *Kucoin) Transfer(goArgs ...*Variant) *Variant {
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
	requestedAmount := this.CurrencyToPrecision(code, amount)
	_ = requestedAmount
	accountsById := this.SafeValue(*this.At(MkString("options")), MkString("accountsByType"), MkMap(&VarMap{}))
	_ = accountsById
	fromId := this.SafeString(accountsById, fromAccount)
	_ = fromId
	if IsTrue(OpEq2(fromId, MkUndefined())) {
		keys := GoGetKeys(accountsById)
		_ = keys
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" fromAccount must be one of "), keys.Join(MkString(", "))))))
	}
	toId := this.SafeString(accountsById, toAccount)
	_ = toId
	if IsTrue(OpEq2(toId, MkUndefined())) {
		keys := GoGetKeys(accountsById)
		_ = keys
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" toAccount must be one of "), keys.Join(MkString(", "))))))
	}
	if IsTrue(OpEq2(fromId, MkString("contract"))) {
		if IsTrue(OpNotEq2(toId, MkString("main"))) {
			panic(NewExchangeError(OpAdd(*this.At(MkString("id")), MkString(" only supports transferring from futures account to main account"))))
		}
		request := MkMap(&VarMap{
			"currency": *(currency).At(MkString("id")),
			"amount":   requestedAmount,
		})
		_ = request
		if IsTrue(OpNot(OpHasMember(MkString("bizNo"), params))) {
			*(request).At(MkString("bizNo")) = this.Uuid22()
		}
		response := this.Call(MkString("futuresPrivatePostTransferOut"), this.Extend(request, params))
		_ = response
		data := this.SafeValue(response, MkString("data"))
		_ = data
		timestamp := this.SafeInteger(data, MkString("createdAt"))
		_ = timestamp
		id := this.SafeString(data, MkString("applyId"))
		_ = id
		currencyId := this.SafeString(data, MkString("currency"))
		_ = currencyId
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		amount := this.SafeNumber(data, MkString("amount"))
		_ = amount
		rawStatus := this.SafeString(data, MkString("status"))
		_ = rawStatus
		status := OpCopy(MkUndefined())
		_ = status
		if IsTrue(OpEq2(rawStatus, MkString("PROCESSING"))) {
			status = MkString("pending")
		}
		return MkMap(&VarMap{
			"info":        response,
			"currency":    code,
			"timestamp":   timestamp,
			"datetime":    this.Iso8601(timestamp),
			"amount":      amount,
			"fromAccount": fromId,
			"toAccount":   toId,
			"id":          id,
			"status":      status,
		})
	} else {
		request := MkMap(&VarMap{
			"currency": *(currency).At(MkString("id")),
			"from":     fromId,
			"to":       toId,
			"amount":   requestedAmount,
		})
		_ = request
		if IsTrue(OpNot(OpHasMember(MkString("clientOid"), params))) {
			*(request).At(MkString("clientOid")) = this.Uuid()
		}
		response := this.Call(MkString("privatePostAccountsInnerTransfer"), this.Extend(request, params))
		_ = response
		data := this.SafeValue(response, MkString("data"))
		_ = data
		id := this.SafeString(data, MkString("orderId"))
		_ = id
		return MkMap(&VarMap{
			"info":        response,
			"id":          id,
			"timestamp":   MkUndefined(),
			"datetime":    MkUndefined(),
			"currency":    code,
			"amount":      requestedAmount,
			"fromAccount": fromId,
			"toAccount":   toId,
			"status":      MkUndefined(),
		})
	}
	return MkUndefined()
}

func (this *Kucoin) FetchLedger(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(code, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchLedger() requires a code param"))))
	}
	this.LoadMarkets()
	this.LoadAccounts()
	currency := this.Currency(code)
	_ = currency
	accountId := this.SafeString(params, MkString("accountId"))
	_ = accountId
	if IsTrue(OpEq2(accountId, MkUndefined())) {
		for i := MkInteger(0); IsTrue(OpLw(i, (*((this).At(MkString("accounts")))).Length)); OpIncr(&i) {
			account := *(*this.At(MkString("accounts"))).At(i)
			_ = account
			if IsTrue(OpAnd(OpEq2(*(account).At(MkString("currency")), code), OpEq2(*(account).At(MkString("type")), MkString("main")))) {
				accountId = *(account).At(MkString("id"))
				break
			}
		}
	}
	if IsTrue(OpEq2(accountId, MkUndefined())) {
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), OpAdd(code, MkString("main account is not loaded in loadAccounts"))))))
	}
	request := MkMap(&VarMap{"accountId": accountId})
	_ = request
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("startAt")) = Math.Floor(OpDiv(since, MkInteger(1000)))
	}
	response := this.Call(MkString("privateGetAccountsAccountIdLedgers"), this.Extend(request, params))
	_ = response
	items := *(*(response).At(MkString("data"))).At(MkString("items"))
	_ = items
	return this.ParseLedger(items, currency, since, limit)
}

func (this *Kucoin) ParseLedgerEntry(goArgs ...*Variant) *Variant {
	item := GoGetArg(goArgs, 0, MkUndefined())
	_ = item
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	currencyId := this.SafeString(item, MkString("currency"))
	_ = currencyId
	code := this.SafeCurrencyCode(currencyId, currency)
	_ = code
	fee := MkMap(&VarMap{
		"cost": this.SafeNumber(item, MkString("fee")),
		"code": code,
	})
	_ = fee
	amount := this.SafeNumber(item, MkString("amount"))
	_ = amount
	after := this.SafeNumber(item, MkString("balance"))
	_ = after
	direction := this.SafeString(item, MkString("direction"))
	_ = direction
	before := OpCopy(MkUndefined())
	_ = before
	if IsTrue(OpAnd(OpNotEq2(after, MkUndefined()), OpNotEq2(amount, MkUndefined()))) {
		difference := OpTrinary(OpEq2(direction, MkString("out")), amount, OpNeg(amount))
		_ = difference
		before = this.Sum(after, difference)
	}
	timestamp := this.SafeInteger(item, MkString("createdAt"))
	_ = timestamp
	type_ := this.ParseLedgerEntryType(this.SafeString(item, MkString("bizType")))
	_ = type_
	contextString := this.SafeString(item, MkString("context"))
	_ = contextString
	id := OpCopy(MkUndefined())
	_ = id
	referenceId := OpCopy(MkUndefined())
	_ = referenceId
	if IsTrue(this.IsJsonEncodedObject(contextString)) {
		context := this.ParseJson(contextString)
		_ = context
		id = this.SafeString(context, MkString("orderId"))
		if IsTrue(OpEq2(type_, MkString("trade"))) {
			referenceId = this.SafeString(context, MkString("tradeId"))
		} else {
			if IsTrue(OpEq2(type_, MkString("transaction"))) {
				referenceId = this.SafeString(context, MkString("txId"))
			}
		}
	}
	return MkMap(&VarMap{
		"id":               id,
		"currency":         code,
		"account":          MkUndefined(),
		"referenceAccount": MkUndefined(),
		"referenceId":      referenceId,
		"status":           MkUndefined(),
		"amount":           amount,
		"before":           before,
		"after":            after,
		"fee":              fee,
		"direction":        direction,
		"timestamp":        timestamp,
		"datetime":         this.Iso8601(timestamp),
		"type":             type_,
		"info":             item,
	})
}

func (this *Kucoin) ParseLedgerEntryType(goArgs ...*Variant) *Variant {
	type_ := GoGetArg(goArgs, 0, MkUndefined())
	_ = type_
	types := MkMap(&VarMap{
		"Exchange":   MkString("trade"),
		"Withdrawal": MkString("transaction"),
		"Deposit":    MkString("transaction"),
		"Transfer":   MkString("transfer"),
	})
	_ = types
	return this.SafeString(types, type_, type_)
}

func (this *Kucoin) FetchPositions(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("futuresPrivateGetPositions"), params)
	_ = response
	return this.SafeValue(response, MkString("data"), response)
}

func (this *Kucoin) Sign(goArgs ...*Variant) *Variant {
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
	versions := this.SafeValue(*this.At(MkString("options")), MkString("versions"), MkMap(&VarMap{}))
	_ = versions
	apiVersions := this.SafeValue(versions, api, MkMap(&VarMap{}))
	_ = apiVersions
	methodVersions := this.SafeValue(apiVersions, method, MkMap(&VarMap{}))
	_ = methodVersions
	defaultVersion := this.SafeString(methodVersions, path, *(*this.At(MkString("options"))).At(MkString("version")))
	_ = defaultVersion
	version := this.SafeString(params, MkString("version"), defaultVersion)
	_ = version
	params = this.Omit(params, MkString("version"))
	endpoint := OpAdd(MkString("/api/"), OpAdd(version, OpAdd(MkString("/"), this.ImplodeParams(path, params))))
	_ = endpoint
	query := this.Omit(params, this.ExtractParams(path))
	_ = query
	endpart := MkString("")
	_ = endpart
	headers = OpTrinary(OpNotEq2(headers, MkUndefined()), headers, MkMap(&VarMap{}))
	if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
		if IsTrue(OpOr(OpEq2(method, MkString("GET")), OpEq2(method, MkString("DELETE")))) {
			OpAddAssign(&endpoint, OpAdd(MkString("?"), this.Urlencode(query)))
		} else {
			body = this.Json(query)
			endpart = OpCopy(body)
			*(headers).At(MkString("Content-Type")) = MkString("application/json")
		}
	}
	url := OpAdd(*(*(*this.At(MkString("urls"))).At(MkString("api"))).At(api), endpoint)
	_ = url
	if IsTrue(OpOr(OpEq2(api, MkString("private")), OpEq2(api, MkString("futuresPrivate")))) {
		this.CheckRequiredCredentials()
		timestamp := (this.Nonce()).Call(MkString("toString"))
		_ = timestamp
		headers = this.Extend(MkMap(&VarMap{
			"KC-API-KEY-VERSION": MkString("2"),
			"KC-API-KEY":         *this.At(MkString("apiKey")),
			"KC-API-TIMESTAMP":   timestamp,
		}), headers)
		apiKeyVersion := this.SafeString(headers, MkString("KC-API-KEY-VERSION"))
		_ = apiKeyVersion
		if IsTrue(OpEq2(apiKeyVersion, MkString("2"))) {
			passphrase := this.Hmac(this.Encode(*this.At(MkString("password"))), this.Encode(*this.At(MkString("secret"))), MkString("sha256"), MkString("base64"))
			_ = passphrase
			*(headers).At(MkString("KC-API-PASSPHRASE")) = OpCopy(passphrase)
		} else {
			*(headers).At(MkString("KC-API-PASSPHRASE")) = OpCopy(*this.At(MkString("password")))
		}
		payload := OpAdd(timestamp, OpAdd(method, OpAdd(endpoint, endpart)))
		_ = payload
		signature := this.Hmac(this.Encode(payload), this.Encode(*this.At(MkString("secret"))), MkString("sha256"), MkString("base64"))
		_ = signature
		*(headers).At(MkString("KC-API-SIGN")) = OpCopy(signature)
		partner := this.SafeValue(*this.At(MkString("options")), MkString("partner"), MkMap(&VarMap{}))
		_ = partner
		partnerId := this.SafeString(partner, MkString("id"))
		_ = partnerId
		partnerSecret := this.SafeString(partner, MkString("secret"))
		_ = partnerSecret
		if IsTrue(OpAnd(OpNotEq2(partnerId, MkUndefined()), OpNotEq2(partnerSecret, MkUndefined()))) {
			partnerPayload := OpAdd(timestamp, OpAdd(partnerId, *this.At(MkString("apiKey"))))
			_ = partnerPayload
			partnerSignature := this.Hmac(this.Encode(partnerPayload), this.Encode(partnerSecret), MkString("sha256"), MkString("base64"))
			_ = partnerSignature
			*(headers).At(MkString("KC-API-PARTNER-SIGN")) = OpCopy(partnerSignature)
			*(headers).At(MkString("KC-API-PARTNER")) = OpCopy(partnerId)
		}
	}
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Kucoin) HandleErrors(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpNot(response)) {
		this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("broad")), body, body)
		return MkUndefined()
	}
	errorCode := this.SafeString(response, MkString("code"))
	_ = errorCode
	message := this.SafeString(response, MkString("msg"), MkString(""))
	_ = message
	this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("exact")), message, OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), message)))
	this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("exact")), errorCode, OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), message)))
	return MkUndefined()
}
