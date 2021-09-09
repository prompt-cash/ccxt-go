package ccxt

import ()

type Bitmart struct {
	*ExchangeBase
}

var _ Exchange = (*Bitmart)(nil)

func init() {
	exchange := &Bitmart{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Bitmart) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("bitmart"),
		"name": MkString("BitMart"),
		"countries": MkArray(&VarArray{
			MkString("US"),
			MkString("CN"),
			MkString("HK"),
			MkString("KR"),
		}),
		"rateLimit": MkInteger(1000),
		"version":   MkString("v1"),
		"certified": MkBool(true),
		"has": MkMap(&VarMap{
			"cancelAllOrders":     MkBool(true),
			"cancelOrder":         MkBool(true),
			"cancelOrders":        MkBool(true),
			"createOrder":         MkBool(true),
			"fetchBalance":        MkBool(true),
			"fetchCanceledOrders": MkBool(true),
			"fetchClosedOrders":   MkBool(true),
			"fetchCurrencies":     MkBool(true),
			"fetchDepositAddress": MkBool(true),
			"fetchDeposits":       MkBool(true),
			"fetchMarkets":        MkBool(true),
			"fetchMyTrades":       MkBool(true),
			"fetchOHLCV":          MkBool(true),
			"fetchOpenOrders":     MkBool(true),
			"fetchOrder":          MkBool(true),
			"fetchOrderBook":      MkBool(true),
			"fetchOrders":         MkBool(true),
			"fetchOrderTrades":    MkBool(true),
			"fetchTicker":         MkBool(true),
			"fetchTickers":        MkBool(true),
			"fetchTime":           MkBool(true),
			"fetchStatus":         MkBool(true),
			"fetchTrades":         MkBool(true),
			"fetchWithdrawals":    MkBool(true),
			"fetchFundingFee":     MkBool(true),
			"withdraw":            MkBool(true),
		}),
		"hostname": MkString("bitmart.com"),
		"urls": MkMap(&VarMap{
			"logo": MkString("https://user-images.githubusercontent.com/1294454/129991357-8f47464b-d0f4-41d6-8a82-34122f0d1398.jpg"),
			"api":  MkString("https://api-cloud.{hostname}"),
			"www":  MkString("https://www.bitmart.com/"),
			"doc":  MkString("https://developer-pro.bitmart.com/"),
			"referral": MkMap(&VarMap{
				"url":      MkString("http://www.bitmart.com/?r=rQCFLh"),
				"discount": MkNumber(0.25),
			}),
			"fees": MkString("https://www.bitmart.com/fee/en"),
		}),
		"requiredCredentials": MkMap(&VarMap{
			"apiKey": MkBool(true),
			"secret": MkBool(true),
			"uid":    MkBool(true),
		}),
		"api": MkMap(&VarMap{
			"public": MkMap(&VarMap{
				"system": MkMap(&VarMap{"get": MkArray(&VarArray{
					MkString("time"),
					MkString("service"),
				})}),
				"account": MkMap(&VarMap{"get": MkArray(&VarArray{
					MkString("currencies"),
				})}),
				"spot": MkMap(&VarMap{"get": MkArray(&VarArray{
					MkString("currencies"),
					MkString("symbols"),
					MkString("symbols/details"),
					MkString("ticker"),
					MkString("steps"),
					MkString("symbols/kline"),
					MkString("symbols/book"),
					MkString("symbols/trades"),
				})}),
				"contract": MkMap(&VarMap{"get": MkArray(&VarArray{
					MkString("contracts"),
					MkString("pnls"),
					MkString("indexes"),
					MkString("tickers"),
					MkString("quote"),
					MkString("indexquote"),
					MkString("trades"),
					MkString("depth"),
					MkString("fundingrate"),
				})}),
			}),
			"private": MkMap(&VarMap{
				"account": MkMap(&VarMap{
					"get": MkArray(&VarArray{
						MkString("wallet"),
						MkString("deposit/address"),
						MkString("withdraw/charge"),
						MkString("deposit-withdraw/history"),
						MkString("deposit-withdraw/detail"),
					}),
					"post": MkArray(&VarArray{
						MkString("withdraw/apply"),
					}),
				}),
				"spot": MkMap(&VarMap{
					"get": MkArray(&VarArray{
						MkString("wallet"),
						MkString("order_detail"),
						MkString("orders"),
						MkString("trades"),
					}),
					"post": MkArray(&VarArray{
						MkString("submit_order"),
						MkString("cancel_order"),
						MkString("cancel_orders"),
					}),
				}),
				"contract": MkMap(&VarMap{
					"get": MkArray(&VarArray{
						MkString("userOrders"),
						MkString("userOrderInfo"),
						MkString("userTrades"),
						MkString("orderTrades"),
						MkString("accounts"),
						MkString("userPositions"),
						MkString("userLiqRecords"),
						MkString("positionFee"),
					}),
					"post": MkArray(&VarArray{
						MkString("batchOrders"),
						MkString("submitOrder"),
						MkString("cancelOrders"),
						MkString("marginOper"),
					}),
				}),
			}),
		}),
		"timeframes": MkMap(&VarMap{
			"1m":  MkInteger(1),
			"3m":  MkInteger(3),
			"5m":  MkInteger(5),
			"15m": MkInteger(15),
			"30m": MkInteger(30),
			"45m": MkInteger(45),
			"1h":  MkInteger(60),
			"2h":  MkInteger(120),
			"3h":  MkInteger(180),
			"4h":  MkInteger(240),
			"1d":  MkInteger(1440),
			"1w":  MkInteger(10080),
			"1M":  MkInteger(43200),
		}),
		"fees": MkMap(&VarMap{"trading": MkMap(&VarMap{
			"tierBased":  MkBool(true),
			"percentage": MkBool(true),
			"taker":      this.ParseNumber(MkString("0.0025")),
			"maker":      this.ParseNumber(MkString("0.0025")),
			"tiers": MkMap(&VarMap{
				"taker": MkArray(&VarArray{
					MkArray(&VarArray{
						this.ParseNumber(MkString("0")),
						this.ParseNumber(MkString("0.0020")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("10")),
						this.ParseNumber(MkString("0.18")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("50")),
						this.ParseNumber(MkString("0.0016")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("250")),
						this.ParseNumber(MkString("0.0014")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("1000")),
						this.ParseNumber(MkString("0.0012")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("5000")),
						this.ParseNumber(MkString("0.0010")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("25000")),
						this.ParseNumber(MkString("0.0008")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("50000")),
						this.ParseNumber(MkString("0.0006")),
					}),
				}),
				"maker": MkArray(&VarArray{
					MkArray(&VarArray{
						this.ParseNumber(MkString("0")),
						this.ParseNumber(MkString("0.001")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("10")),
						this.ParseNumber(MkString("0.0009")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("50")),
						this.ParseNumber(MkString("0.0008")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("250")),
						this.ParseNumber(MkString("0.0007")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("1000")),
						this.ParseNumber(MkString("0.0006")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("5000")),
						this.ParseNumber(MkString("0.0005")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("25000")),
						this.ParseNumber(MkString("0.0004")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("50000")),
						this.ParseNumber(MkString("0.0003")),
					}),
				}),
			}),
		})}),
		"precisionMode": TICK_SIZE,
		"exceptions": MkMap(&VarMap{
			"exact": MkMap(&VarMap{
				"30000": ExchangeError,
				"30001": AuthenticationError,
				"30002": AuthenticationError,
				"30003": AccountSuspended,
				"30004": AuthenticationError,
				"30005": AuthenticationError,
				"30006": AuthenticationError,
				"30007": AuthenticationError,
				"30008": AuthenticationError,
				"30010": PermissionDenied,
				"30011": AuthenticationError,
				"30012": AuthenticationError,
				"30013": RateLimitExceeded,
				"30014": ExchangeNotAvailable,
				"60000": BadRequest,
				"60001": BadRequest,
				"60002": BadRequest,
				"60003": ExchangeError,
				"60004": ExchangeError,
				"60005": ExchangeError,
				"60006": ExchangeError,
				"60007": InvalidAddress,
				"60008": InsufficientFunds,
				"60009": ExchangeError,
				"60010": ExchangeError,
				"60011": InvalidAddress,
				"60012": ExchangeError,
				"60020": PermissionDenied,
				"60021": PermissionDenied,
				"60022": PermissionDenied,
				"60030": BadRequest,
				"60031": BadRequest,
				"60050": ExchangeError,
				"60051": ExchangeError,
				"50000": BadRequest,
				"50001": BadSymbol,
				"50002": BadRequest,
				"50003": BadRequest,
				"50004": BadRequest,
				"50005": OrderNotFound,
				"50006": InvalidOrder,
				"50007": InvalidOrder,
				"50008": InvalidOrder,
				"50009": InvalidOrder,
				"50010": InvalidOrder,
				"50011": InvalidOrder,
				"50012": InvalidOrder,
				"50013": InvalidOrder,
				"50014": BadRequest,
				"50015": BadRequest,
				"50016": BadRequest,
				"50017": BadRequest,
				"50018": BadRequest,
				"50019": BadRequest,
				"50020": InsufficientFunds,
				"50021": BadRequest,
				"50022": ExchangeNotAvailable,
				"50023": BadSymbol,
				"53000": AccountSuspended,
				"57001": BadRequest,
				"58001": BadRequest,
				"59001": ExchangeError,
				"59002": ExchangeError,
				"40001": ExchangeError,
				"40002": ExchangeError,
				"40003": ExchangeError,
				"40004": ExchangeError,
				"40005": ExchangeError,
				"40006": PermissionDenied,
				"40007": BadRequest,
				"40008": InvalidNonce,
				"40009": BadRequest,
				"40010": BadRequest,
				"40011": BadRequest,
				"40012": ExchangeError,
				"40013": ExchangeError,
				"40014": BadSymbol,
				"40015": BadSymbol,
				"40016": InvalidOrder,
				"40017": InvalidOrder,
				"40018": InvalidOrder,
				"40019": ExchangeError,
				"40020": InvalidOrder,
				"40021": ExchangeError,
				"40022": ExchangeError,
				"40023": ExchangeError,
				"40024": ExchangeError,
				"40025": ExchangeError,
				"40026": ExchangeError,
				"40027": InsufficientFunds,
				"40028": PermissionDenied,
				"40029": InvalidOrder,
				"40030": InvalidOrder,
				"40031": InvalidOrder,
				"40032": InvalidOrder,
				"40033": InvalidOrder,
				"40034": BadSymbol,
			}),
			"broad": MkMap(&VarMap{}),
		}),
		"commonCurrencies": MkMap(&VarMap{
			"COT": MkString("Community Coin"),
			"CPC": MkString("CPCoin"),
			"MVP": MkString("MVP Coin"),
			"ONE": MkString("Menlo One"),
			"PLA": MkString("Plair"),
			"TCT": MkString("TacoCat Token"),
		}),
		"options": MkMap(&VarMap{
			"defaultType":                       MkString("spot"),
			"fetchBalance":                      MkMap(&VarMap{"type": MkString("spot")}),
			"createMarketBuyOrderRequiresPrice": MkBool(true),
		}),
	}))
}

func (this *Bitmart) FetchTime(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicSystemGetTime"), params)
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	return this.SafeInteger(data, MkString("server_time"))
}

func (this *Bitmart) FetchStatus(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	options := this.SafeValue(*this.At(MkString("options")), MkString("fetchBalance"), MkMap(&VarMap{}))
	_ = options
	defaultType := this.SafeString(*this.At(MkString("options")), MkString("defaultType"))
	_ = defaultType
	type_ := this.SafeString(options, MkString("type"), defaultType)
	_ = type_
	type_ = this.SafeString(params, MkString("type"), type_)
	params = this.Omit(params, MkString("type"))
	response := this.Call(MkString("publicSystemGetService"), params)
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	services := this.SafeValue(data, MkString("service"), MkArray(&VarArray{}))
	_ = services
	servicesByType := this.IndexBy(services, MkString("service_type"))
	_ = servicesByType
	if IsTrue(OpOr(OpEq2(type_, MkString("swap")), OpEq2(type_, MkString("future")))) {
		type_ = MkString("contract")
	}
	service := this.SafeValue(servicesByType, type_)
	_ = service
	status := OpCopy(MkUndefined())
	_ = status
	eta := OpCopy(MkUndefined())
	_ = eta
	if IsTrue(OpNotEq2(service, MkUndefined())) {
		statusCode := this.SafeInteger(service, MkString("status"))
		_ = statusCode
		if IsTrue(OpEq2(statusCode, MkInteger(2))) {
			status = MkString("ok")
		} else {
			status = MkString("maintenance")
			eta = this.SafeInteger(service, MkString("end_time"))
		}
	}
	*this.At(MkString("status")) = this.Extend(*this.At(MkString("status")), MkMap(&VarMap{
		"status":  status,
		"updated": this.Milliseconds(),
		"eta":     eta,
	}))
	return *this.At(MkString("status"))
}

func (this *Bitmart) FetchSpotMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicSpotGetSymbolsDetails"), params)
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	symbols := this.SafeValue(data, MkString("symbols"), MkArray(&VarArray{}))
	_ = symbols
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, symbols.Length)); OpIncr(&i) {
		market := *(symbols).At(i)
		_ = market
		id := this.SafeString(market, MkString("symbol"))
		_ = id
		numericId := this.SafeInteger(market, MkString("symbol_id"))
		_ = numericId
		baseId := this.SafeString(market, MkString("base_currency"))
		_ = baseId
		quoteId := this.SafeString(market, MkString("quote_currency"))
		_ = quoteId
		base := this.SafeCurrencyCode(baseId)
		_ = base
		quote := this.SafeCurrencyCode(quoteId)
		_ = quote
		symbol := OpAdd(base, OpAdd(MkString("/"), quote))
		_ = symbol
		pricePrecision := this.SafeInteger(market, MkString("price_max_precision"))
		_ = pricePrecision
		precision := MkMap(&VarMap{
			"amount": this.SafeNumber(market, MkString("base_min_size")),
			"price":  this.ParseNumber(this.DecimalToPrecision(Math.Pow(MkInteger(10), OpNeg(pricePrecision)), ROUND, MkInteger(12))),
		})
		_ = precision
		minBuyCost := this.SafeNumber(market, MkString("min_buy_amount"))
		_ = minBuyCost
		minSellCost := this.SafeNumber(market, MkString("min_sell_amount"))
		_ = minSellCost
		minCost := Math.Max(minBuyCost, minSellCost)
		_ = minCost
		limits := MkMap(&VarMap{
			"amount": MkMap(&VarMap{
				"min": this.SafeNumber(market, MkString("base_min_size")),
				"max": this.SafeNumber(market, MkString("base_max_size")),
			}),
			"price": MkMap(&VarMap{
				"min": MkUndefined(),
				"max": MkUndefined(),
			}),
			"cost": MkMap(&VarMap{
				"min": minCost,
				"max": MkUndefined(),
			}),
		})
		_ = limits
		result.Push(MkMap(&VarMap{
			"id":        id,
			"numericId": numericId,
			"symbol":    symbol,
			"base":      base,
			"quote":     quote,
			"baseId":    baseId,
			"quoteId":   quoteId,
			"type":      MkString("spot"),
			"spot":      MkBool(true),
			"future":    MkBool(false),
			"swap":      MkBool(false),
			"precision": precision,
			"limits":    limits,
			"info":      market,
			"active":    MkUndefined(),
		}))
	}
	return result
}

func (this *Bitmart) FetchContractMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicContractGetContracts"), params)
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	contracts := this.SafeValue(data, MkString("contracts"), MkArray(&VarArray{}))
	_ = contracts
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, contracts.Length)); OpIncr(&i) {
		market := *(contracts).At(i)
		_ = market
		contract := this.SafeValue(market, MkString("contract"), MkMap(&VarMap{}))
		_ = contract
		id := this.SafeString(contract, MkString("contract_id"))
		_ = id
		numericId := this.SafeInteger(contract, MkString("contract_id"))
		_ = numericId
		baseId := this.SafeString(contract, MkString("base_coin"))
		_ = baseId
		quoteId := this.SafeString(contract, MkString("quote_coin"))
		_ = quoteId
		base := this.SafeCurrencyCode(baseId)
		_ = base
		quote := this.SafeCurrencyCode(quoteId)
		_ = quote
		symbol := this.SafeString(contract, MkString("name"))
		_ = symbol
		amountPrecision := this.SafeNumber(contract, MkString("vol_unit"))
		_ = amountPrecision
		pricePrecision := this.SafeNumber(contract, MkString("price_unit"))
		_ = pricePrecision
		precision := MkMap(&VarMap{
			"amount": amountPrecision,
			"price":  pricePrecision,
		})
		_ = precision
		limits := MkMap(&VarMap{
			"amount": MkMap(&VarMap{
				"min": this.SafeNumber(contract, MkString("min_vol")),
				"max": this.SafeNumber(contract, MkString("max_vol")),
			}),
			"price": MkMap(&VarMap{
				"min": MkUndefined(),
				"max": MkUndefined(),
			}),
			"cost": MkMap(&VarMap{
				"min": MkUndefined(),
				"max": MkUndefined(),
			}),
		})
		_ = limits
		contractType := this.SafeValue(contract, MkString("contract_type"))
		_ = contractType
		future := OpCopy(MkBool(false))
		_ = future
		swap := OpCopy(MkBool(false))
		_ = swap
		type_ := MkString("contract")
		_ = type_
		if IsTrue(OpEq2(contractType, MkInteger(1))) {
			type_ = MkString("swap")
			swap = OpCopy(MkBool(true))
		} else {
			if IsTrue(OpEq2(contractType, MkInteger(2))) {
				type_ = MkString("future")
				future = OpCopy(MkBool(true))
			}
		}
		feeConfig := this.SafeValue(market, MkString("fee_config"), MkMap(&VarMap{}))
		_ = feeConfig
		maker := this.SafeNumber(feeConfig, MkString("maker_fee"))
		_ = maker
		taker := this.SafeNumber(feeConfig, MkString("taker_fee"))
		_ = taker
		result.Push(MkMap(&VarMap{
			"id":        id,
			"numericId": numericId,
			"symbol":    symbol,
			"base":      base,
			"quote":     quote,
			"baseId":    baseId,
			"quoteId":   quoteId,
			"maker":     maker,
			"taker":     taker,
			"type":      type_,
			"spot":      MkBool(false),
			"future":    future,
			"swap":      swap,
			"precision": precision,
			"limits":    limits,
			"info":      market,
			"active":    MkUndefined(),
		}))
	}
	return result
}

func (this *Bitmart) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	return this.FetchSpotMarkets()
}

func (this *Bitmart) FetchFundingFee(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{"currency": *(currency).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("privateAccountGetWithdrawCharge"), this.Extend(request, params))
	_ = response
	data := *(response).At(MkString("data"))
	_ = data
	withdrawFees := MkMap(&VarMap{})
	_ = withdrawFees
	*(withdrawFees).At(code) = this.SafeNumber(data, MkString("withdraw_fee"))
	return MkMap(&VarMap{
		"info":     response,
		"withdraw": withdrawFees,
		"deposit":  MkMap(&VarMap{}),
	})
}

func (this *Bitmart) ParseTicker(goArgs ...*Variant) *Variant {
	ticker := GoGetArg(goArgs, 0, MkUndefined())
	_ = ticker
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.SafeTimestamp(ticker, MkString("timestamp"), this.Milliseconds())
	_ = timestamp
	marketId := this.SafeString2(ticker, MkString("symbol"), MkString("contract_id"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market, MkString("_"))
	_ = symbol
	last := this.SafeNumber2(ticker, MkString("close_24h"), MkString("last_price"))
	_ = last
	percentage := this.SafeNumber(ticker, MkString("fluctuation"), MkString("rise_fall_rate"))
	_ = percentage
	if IsTrue(OpNotEq2(percentage, MkUndefined())) {
		OpMultiAssign(&percentage, MkInteger(100))
	}
	baseVolume := this.SafeNumber2(ticker, MkString("base_volume_24h"), MkString("base_coin_volume"))
	_ = baseVolume
	quoteVolume := this.SafeNumber2(ticker, MkString("quote_volume_24h"), MkString("quote_coin_volume"))
	_ = quoteVolume
	vwap := this.Vwap(baseVolume, quoteVolume)
	_ = vwap
	open := this.SafeNumber2(ticker, MkString("open_24h"), MkString("open"))
	_ = open
	average := OpCopy(MkUndefined())
	_ = average
	if IsTrue(OpAnd(OpNotEq2(last, MkUndefined()), OpNotEq2(open, MkUndefined()))) {
		average = OpDiv(this.Sum(last, open), MkInteger(2))
	}
	average = this.SafeNumber(ticker, MkString("avg_price"), average)
	price := this.SafeValue(ticker, MkString("depth_price"), ticker)
	_ = price
	return MkMap(&VarMap{
		"symbol":        symbol,
		"timestamp":     timestamp,
		"datetime":      this.Iso8601(timestamp),
		"high":          this.SafeNumber2(ticker, MkString("high"), MkString("high_24h")),
		"low":           this.SafeNumber2(ticker, MkString("low"), MkString("low_24h")),
		"bid":           this.SafeNumber(price, MkString("best_bid"), MkString("bid_price")),
		"bidVolume":     this.SafeNumber(ticker, MkString("best_bid_size")),
		"ask":           this.SafeNumber(price, MkString("best_ask"), MkString("ask_price")),
		"askVolume":     this.SafeNumber(ticker, MkString("best_ask_size")),
		"vwap":          vwap,
		"open":          this.SafeNumber(ticker, MkString("open_24h")),
		"close":         last,
		"last":          last,
		"previousClose": MkUndefined(),
		"change":        MkUndefined(),
		"percentage":    percentage,
		"average":       average,
		"baseVolume":    baseVolume,
		"quoteVolume":   quoteVolume,
		"info":          ticker,
	})
}

func (this *Bitmart) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{})
	_ = request
	method := OpCopy(MkUndefined())
	_ = method
	if IsTrue(OpOr(*(market).At(MkString("swap")), *(market).At(MkString("future")))) {
		method = MkString("publicContractGetTickers")
		*(request).At(MkString("contractID")) = *(market).At(MkString("id"))
	} else {
		if IsTrue(*(market).At(MkString("spot"))) {
			method = MkString("publicSpotGetTicker")
			*(request).At(MkString("symbol")) = *(market).At(MkString("id"))
		}
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	tickers := this.SafeValue(data, MkString("tickers"), MkArray(&VarArray{}))
	_ = tickers
	tickersById := this.IndexBy(tickers, MkString("symbol"))
	_ = tickersById
	ticker := this.SafeValue(tickersById, *(market).At(MkString("id")))
	_ = ticker
	return this.ParseTicker(ticker, market)
}

func (this *Bitmart) FetchTickers(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	defaultType := this.SafeString(*this.At(MkString("options")), MkString("defaultType"), MkString("spot"))
	_ = defaultType
	type_ := this.SafeString(params, MkString("type"), defaultType)
	_ = type_
	params = this.Omit(params, MkString("type"))
	method := OpCopy(MkUndefined())
	_ = method
	if IsTrue(OpOr(OpEq2(type_, MkString("swap")), OpEq2(type_, MkString("future")))) {
		method = MkString("publicContractGetTickers")
	} else {
		if IsTrue(OpEq2(type_, MkString("spot"))) {
			method = MkString("publicSpotGetTicker")
		}
	}
	response := this.Call(method, params)
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	tickers := this.SafeValue(data, MkString("tickers"), MkArray(&VarArray{}))
	_ = tickers
	result := MkMap(&VarMap{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, tickers.Length)); OpIncr(&i) {
		ticker := this.ParseTicker(*(tickers).At(i))
		_ = ticker
		symbol := *(ticker).At(MkString("symbol"))
		_ = symbol
		*(result).At(symbol) = OpCopy(ticker)
	}
	return this.FilterByArray(result, MkString("symbol"), symbols)
}

func (this *Bitmart) FetchCurrencies(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicAccountGetCurrencies"), params)
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	currencies := this.SafeValue(data, MkString("currencies"), MkArray(&VarArray{}))
	_ = currencies
	result := MkMap(&VarMap{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, currencies.Length)); OpIncr(&i) {
		currency := *(currencies).At(i)
		_ = currency
		id := this.SafeString(currency, MkString("currency"))
		_ = id
		code := this.SafeCurrencyCode(id)
		_ = code
		name := this.SafeString(currency, MkString("name"))
		_ = name
		withdrawEnabled := this.SafeValue(currency, MkString("withdraw_enabled"))
		_ = withdrawEnabled
		depositEnabled := this.SafeValue(currency, MkString("deposit_enabled"))
		_ = depositEnabled
		active := OpAnd(withdrawEnabled, depositEnabled)
		_ = active
		*(result).At(code) = MkMap(&VarMap{
			"id":        id,
			"code":      code,
			"name":      name,
			"info":      currency,
			"active":    active,
			"fee":       MkUndefined(),
			"precision": MkUndefined(),
			"limits": MkMap(&VarMap{
				"amount": MkMap(&VarMap{
					"min": MkUndefined(),
					"max": MkUndefined(),
				}),
				"withdraw": MkMap(&VarMap{
					"min": MkUndefined(),
					"max": MkUndefined(),
				}),
			}),
		})
	}
	return result
}

func (this *Bitmart) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{})
	_ = request
	method := OpCopy(MkUndefined())
	_ = method
	if IsTrue(*(market).At(MkString("spot"))) {
		method = MkString("publicSpotGetSymbolsBook")
		*(request).At(MkString("symbol")) = *(market).At(MkString("id"))
	} else {
		if IsTrue(OpOr(*(market).At(MkString("swap")), *(market).At(MkString("future")))) {
			method = MkString("publicContractGetDepth")
			*(request).At(MkString("contractID")) = *(market).At(MkString("id"))
			if IsTrue(OpNotEq2(limit, MkUndefined())) {
				*(request).At(MkString("count")) = OpCopy(limit)
			}
		}
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	if IsTrue(*(market).At(MkString("spot"))) {
		return this.ParseOrderBook(data, symbol, MkUndefined(), MkString("buys"), MkString("sells"), MkString("price"), MkString("amount"))
	} else {
		if IsTrue(OpOr(*(market).At(MkString("swap")), *(market).At(MkString("future")))) {
			return this.ParseOrderBook(data, symbol, MkUndefined(), MkString("buys"), MkString("sells"), MkString("price"), MkString("vol"))
		}
	}
	return MkUndefined()
}

func (this *Bitmart) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	id := this.SafeString2(trade, MkString("trade_id"), MkString("detail_id"))
	_ = id
	timestamp := this.SafeInteger2(trade, MkString("order_time"), MkString("create_time"))
	_ = timestamp
	if IsTrue(OpEq2(timestamp, MkUndefined())) {
		timestamp = this.Parse8601(this.SafeString(trade, MkString("created_at")))
	}
	type_ := OpCopy(MkUndefined())
	_ = type_
	way := this.SafeInteger(trade, MkString("way"))
	_ = way
	side := this.SafeStringLower2(trade, MkString("type"), MkString("side"))
	_ = side
	if IsTrue(OpAnd(OpEq2(side, MkUndefined()), OpNotEq2(way, MkUndefined()))) {
		if IsTrue(OpLw(way, MkInteger(5))) {
			side = MkString("buy")
		} else {
			side = MkString("sell")
		}
	}
	takerOrMaker := OpCopy(MkUndefined())
	_ = takerOrMaker
	execType := this.SafeString(trade, MkString("exec_type"))
	_ = execType
	if IsTrue(OpNotEq2(execType, MkUndefined())) {
		takerOrMaker = OpTrinary(OpEq2(execType, MkString("M")), MkString("maker"), MkString("taker"))
	}
	price := this.SafeNumber2(trade, MkString("price"), MkString("deal_price"))
	_ = price
	price = this.SafeNumber(trade, MkString("price_avg"), price)
	amount := this.SafeNumber2(trade, MkString("amount"), MkString("deal_vol"))
	_ = amount
	amount = this.SafeNumber(trade, MkString("size"), amount)
	cost := this.SafeNumber2(trade, MkString("count"), MkString("notional"))
	_ = cost
	if IsTrue(OpAnd(OpEq2(cost, MkUndefined()), OpAnd(OpNotEq2(price, MkUndefined()), OpNotEq2(amount, MkUndefined())))) {
		cost = OpMulti(amount, price)
	}
	orderId := this.SafeInteger(trade, MkString("order_id"))
	_ = orderId
	marketId := this.SafeString2(trade, MkString("contract_id"), MkString("symbol"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market, MkString("_"))
	_ = symbol
	feeCost := this.SafeNumber(trade, MkString("fees"))
	_ = feeCost
	fee := OpCopy(MkUndefined())
	_ = fee
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		feeCurrencyId := this.SafeString(trade, MkString("fee_coin_name"))
		_ = feeCurrencyId
		feeCurrencyCode := this.SafeCurrencyCode(feeCurrencyId)
		_ = feeCurrencyCode
		if IsTrue(OpAnd(OpEq2(feeCurrencyCode, MkUndefined()), OpNotEq2(market, MkUndefined()))) {
			feeCurrencyCode = OpTrinary(OpEq2(side, MkString("buy")), *(market).At(MkString("base")), *(market).At(MkString("quote")))
		}
		fee = MkMap(&VarMap{
			"cost":     feeCost,
			"currency": feeCurrencyCode,
		})
	}
	return MkMap(&VarMap{
		"info":         trade,
		"id":           id,
		"order":        orderId,
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       symbol,
		"type":         type_,
		"side":         side,
		"price":        price,
		"amount":       amount,
		"cost":         cost,
		"takerOrMaker": takerOrMaker,
		"fee":          fee,
	})
}

func (this *Bitmart) FetchTrades(goArgs ...*Variant) *Variant {
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
	method := OpCopy(MkUndefined())
	_ = method
	if IsTrue(*(market).At(MkString("spot"))) {
		*(request).At(MkString("symbol")) = *(market).At(MkString("id"))
		method = MkString("publicSpotGetSymbolsTrades")
	} else {
		if IsTrue(OpOr(*(market).At(MkString("swap")), *(market).At(MkString("future")))) {
			method = MkString("publicContractGetTrades")
			*(request).At(MkString("contractID")) = *(market).At(MkString("id"))
		}
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	trades := this.SafeValue(data, MkString("trades"), MkArray(&VarArray{}))
	_ = trades
	return this.ParseTrades(trades, market, since, limit)
}

func (this *Bitmart) ParseOHLCV(goArgs ...*Variant) *Variant {
	ohlcv := GoGetArg(goArgs, 0, MkUndefined())
	_ = ohlcv
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	return MkArray(&VarArray{
		this.SafeTimestamp(ohlcv, MkString("timestamp")),
		this.SafeNumber(ohlcv, MkString("open")),
		this.SafeNumber(ohlcv, MkString("high")),
		this.SafeNumber(ohlcv, MkString("low")),
		this.SafeNumber(ohlcv, MkString("close")),
		this.SafeNumber(ohlcv, MkString("volume")),
	})
}

func (this *Bitmart) FetchOHLCV(goArgs ...*Variant) *Variant {
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
	type_ := *(market).At(MkString("type"))
	_ = type_
	method := OpCopy(MkUndefined())
	_ = method
	request := MkMap(&VarMap{})
	_ = request
	duration := this.ParseTimeframe(timeframe)
	_ = duration
	if IsTrue(OpEq2(type_, MkString("spot"))) {
		method = MkString("publicSpotGetSymbolsKline")
		*(request).At(MkString("symbol")) = *(market).At(MkString("id"))
		*(request).At(MkString("step")) = *(*this.At(MkString("timeframes"))).At(timeframe)
		maxLimit := MkInteger(500)
		_ = maxLimit
		if IsTrue(OpEq2(limit, MkUndefined())) {
			limit = OpCopy(maxLimit)
		}
		limit = Math.Min(maxLimit, limit)
		if IsTrue(OpEq2(since, MkUndefined())) {
			end := parseInt(OpDiv(this.Milliseconds(), MkInteger(1000)))
			_ = end
			start := OpSub(end, OpMulti(limit, duration))
			_ = start
			*(request).At(MkString("from")) = OpCopy(start)
			*(request).At(MkString("to")) = OpCopy(end)
		} else {
			start := parseInt(OpDiv(since, MkInteger(1000)))
			_ = start
			end := this.Sum(start, OpMulti(limit, duration))
			_ = end
			*(request).At(MkString("from")) = OpCopy(start)
			*(request).At(MkString("to")) = OpCopy(end)
		}
	} else {
		if IsTrue(OpOr(OpEq2(type_, MkString("swap")), OpEq2(type_, MkString("future")))) {
			method = MkString("publicContractGetQuote")
			*(request).At(MkString("contractID")) = *(market).At(MkString("id"))
			defaultLimit := MkInteger(500)
			_ = defaultLimit
			if IsTrue(OpEq2(limit, MkUndefined())) {
				limit = OpCopy(defaultLimit)
			}
			if IsTrue(OpEq2(since, MkUndefined())) {
				end := parseInt(OpDiv(this.Milliseconds(), MkInteger(1000)))
				_ = end
				start := OpSub(end, OpMulti(limit, duration))
				_ = start
				*(request).At(MkString("startTime")) = OpCopy(start)
				*(request).At(MkString("endTime")) = OpCopy(end)
			} else {
				start := parseInt(OpDiv(since, MkInteger(1000)))
				_ = start
				end := this.Sum(start, OpMulti(limit, duration))
				_ = end
				*(request).At(MkString("startTime")) = OpCopy(start)
				*(request).At(MkString("endTime")) = OpCopy(end)
			}
			*(request).At(MkString("unit")) = *(*this.At(MkString("timeframes"))).At(timeframe)
			*(request).At(MkString("resolution")) = MkString("M")
		}
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	if IsTrue(GoIsArray(data)) {
		return this.ParseOHLCVs(data, market, timeframe, since, limit)
	} else {
		klines := this.SafeValue(data, MkString("klines"), MkArray(&VarArray{}))
		_ = klines
		return this.ParseOHLCVs(klines, market, timeframe, since, limit)
	}
	return MkUndefined()
}

func (this *Bitmart) FetchMyTrades(goArgs ...*Variant) *Variant {
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
	method := OpCopy(MkUndefined())
	_ = method
	request := MkMap(&VarMap{})
	_ = request
	if IsTrue(*(market).At(MkString("spot"))) {
		*(request).At(MkString("symbol")) = *(market).At(MkString("id"))
		*(request).At(MkString("offset")) = MkInteger(1)
		if IsTrue(OpEq2(limit, MkUndefined())) {
			limit = MkInteger(100)
		}
		*(request).At(MkString("limit")) = OpCopy(limit)
		method = MkString("privateSpotGetTrades")
	} else {
		if IsTrue(OpOr(*(market).At(MkString("swap")), *(market).At(MkString("future")))) {
			*(request).At(MkString("contractID")) = *(market).At(MkString("id"))
			if IsTrue(OpNotEq2(limit, MkUndefined())) {
				*(request).At(MkString("size")) = OpCopy(limit)
			}
			method = MkString("privateContractGetUserTrades")
		}
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	trades := this.SafeValue(data, MkString("trades"), MkArray(&VarArray{}))
	_ = trades
	return this.ParseTrades(trades, market, since, limit)
}

func (this *Bitmart) FetchOrderTrades(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchOrderTrades() requires a symbol argument"))))
	}
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	method := OpCopy(MkUndefined())
	_ = method
	request := MkMap(&VarMap{})
	_ = request
	if IsTrue(*(market).At(MkString("spot"))) {
		*(request).At(MkString("symbol")) = *(market).At(MkString("id"))
		*(request).At(MkString("order_id")) = OpCopy(id)
		method = MkString("privateSpotGetTrades")
	} else {
		if IsTrue(OpOr(*(market).At(MkString("swap")), *(market).At(MkString("future")))) {
			*(request).At(MkString("contractID")) = *(market).At(MkString("id"))
			*(request).At(MkString("orderID")) = OpCopy(id)
			method = MkString("privateContractGetOrderTrades")
		}
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	trades := this.SafeValue(data, MkString("trades"), MkArray(&VarArray{}))
	_ = trades
	return this.ParseTrades(trades, market, since, limit)
}

func (this *Bitmart) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	method := OpCopy(MkUndefined())
	_ = method
	options := this.SafeValue(*this.At(MkString("options")), MkString("fetchBalance"), MkMap(&VarMap{}))
	_ = options
	defaultType := this.SafeString(*this.At(MkString("options")), MkString("defaultType"), MkString("spot"))
	_ = defaultType
	type_ := this.SafeString(options, MkString("type"), defaultType)
	_ = type_
	type_ = this.SafeString(params, MkString("type"), type_)
	params = this.Omit(params, MkString("type"))
	if IsTrue(OpEq2(type_, MkString("spot"))) {
		method = MkString("privateSpotGetWallet")
	} else {
		if IsTrue(OpEq2(type_, MkString("account"))) {
			method = MkString("privateAccountGetWallet")
		} else {
			if IsTrue(OpOr(OpEq2(type_, MkString("swap")), OpOr(OpEq2(type_, MkString("future")), OpEq2(type_, MkString("contract"))))) {
				method = MkString("privateContractGetAccounts")
			}
		}
	}
	response := this.Call(method, params)
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	wallet := this.SafeValue2(data, MkString("wallet"), MkString("accounts"), MkArray(&VarArray{}))
	_ = wallet
	result := MkMap(&VarMap{"info": response})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, wallet.Length)); OpIncr(&i) {
		balance := *(wallet).At(i)
		_ = balance
		currencyId := this.SafeString2(balance, MkString("id"), MkString("currency"))
		_ = currencyId
		currencyId = this.SafeString(balance, MkString("coin_code"), currencyId)
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		account := this.Account()
		_ = account
		*(account).At(MkString("free")) = this.SafeString2(balance, MkString("available"), MkString("available_vol"))
		*(account).At(MkString("used")) = this.SafeString2(balance, MkString("frozen"), MkString("freeze_vol"))
		*(result).At(code) = OpCopy(account)
	}
	return this.ParseBalance(result)
}

func (this *Bitmart) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	id := OpCopy(MkUndefined())
	_ = id
	if IsTrue(OpEq2(OpType(order), MkString("string"))) {
		id = OpCopy(order)
		order = MkMap(&VarMap{})
	}
	id = this.SafeString(order, MkString("order_id"), id)
	timestamp := this.Parse8601(this.SafeString(order, MkString("created_at")))
	_ = timestamp
	timestamp = this.SafeInteger(order, MkString("create_time"), timestamp)
	marketId := this.SafeString2(order, MkString("symbol"), MkString("contract_id"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market, MkString("_"))
	_ = symbol
	status := OpCopy(MkUndefined())
	_ = status
	if IsTrue(OpNotEq2(market, MkUndefined())) {
		status = this.ParseOrderStatusByType(*(market).At(MkString("type")), this.SafeString(order, MkString("status")))
	}
	price := this.SafeNumber(order, MkString("price"))
	_ = price
	average := this.SafeNumber2(order, MkString("price_avg"), MkString("done_avg_price"))
	_ = average
	amount := this.SafeNumber2(order, MkString("size"), MkString("vol"))
	_ = amount
	filled := this.SafeNumber2(order, MkString("filled_size"), MkString("done_vol"))
	_ = filled
	side := this.SafeString(order, MkString("side"))
	_ = side
	side = this.SafeString(order, MkString("way"), side)
	category := this.SafeInteger(order, MkString("category"))
	_ = category
	type_ := this.SafeString(order, MkString("type"))
	_ = type_
	if IsTrue(OpEq2(category, MkInteger(1))) {
		type_ = MkString("limit")
	} else {
		if IsTrue(OpEq2(category, MkInteger(2))) {
			type_ = MkString("market")
		}
	}
	if IsTrue(OpEq2(type_, MkString("market"))) {
		if IsTrue(OpEq2(price, MkNumber(0.0))) {
			price = OpCopy(MkUndefined())
		}
		if IsTrue(OpEq2(average, MkNumber(0.0))) {
			average = OpCopy(MkUndefined())
		}
	}
	return this.SafeOrder(MkMap(&VarMap{
		"id":                 id,
		"clientOrderId":      MkUndefined(),
		"info":               order,
		"timestamp":          timestamp,
		"datetime":           this.Iso8601(timestamp),
		"lastTradeTimestamp": MkUndefined(),
		"symbol":             symbol,
		"type":               type_,
		"timeInForce":        MkUndefined(),
		"postOnly":           MkUndefined(),
		"side":               side,
		"price":              price,
		"stopPrice":          MkUndefined(),
		"amount":             amount,
		"cost":               MkUndefined(),
		"average":            average,
		"filled":             filled,
		"remaining":          MkUndefined(),
		"status":             status,
		"fee":                MkUndefined(),
		"trades":             MkUndefined(),
	}))
}

func (this *Bitmart) ParseOrderStatusByType(goArgs ...*Variant) *Variant {
	type_ := GoGetArg(goArgs, 0, MkUndefined())
	_ = type_
	status := GoGetArg(goArgs, 1, MkUndefined())
	_ = status
	statusesByType := MkMap(&VarMap{
		"spot": MkMap(&VarMap{
			"1": MkString("failed"),
			"2": MkString("open"),
			"3": MkString("failed"),
			"4": MkString("open"),
			"5": MkString("open"),
			"6": MkString("closed"),
			"7": MkString("canceling"),
			"8": MkString("canceled"),
		}),
		"swap": MkMap(&VarMap{
			"1": MkString("open"),
			"2": MkString("open"),
			"4": MkString("closed"),
		}),
	})
	_ = statusesByType
	statuses := this.SafeValue(statusesByType, type_, MkMap(&VarMap{}))
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Bitmart) CreateOrder(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{})
	_ = request
	method := OpCopy(MkUndefined())
	_ = method
	if IsTrue(*(market).At(MkString("spot"))) {
		*(request).At(MkString("symbol")) = *(market).At(MkString("id"))
		*(request).At(MkString("side")) = OpCopy(side)
		*(request).At(MkString("type")) = OpCopy(type_)
		method = MkString("privateSpotPostSubmitOrder")
		if IsTrue(OpEq2(type_, MkString("limit"))) {
			*(request).At(MkString("size")) = this.AmountToPrecision(symbol, amount)
			*(request).At(MkString("price")) = this.PriceToPrecision(symbol, price)
		} else {
			if IsTrue(OpEq2(type_, MkString("market"))) {
				if IsTrue(OpEq2(side, MkString("buy"))) {
					notional := this.SafeNumber(params, MkString("notional"))
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
								panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")), MkString(" createOrder() requires the price argument with market buy orders to calculate total order cost (amount to spend), where cost = amount * price. Supply a price argument to createOrder() call if you want the cost to be calculated for you from price and amount, or, alternatively, add .options['createMarketBuyOrderRequiresPrice'] = false and supply the total cost value in the 'amount' argument or in the 'notional' extra parameter (the exchange-specific behaviour)"))))
							}
						}
					} else {
						notional = OpTrinary(OpEq2(notional, MkUndefined()), amount, notional)
					}
					precision := *(*(market).At(MkString("precision"))).At(MkString("price"))
					_ = precision
					*(request).At(MkString("notional")) = this.DecimalToPrecision(notional, TRUNCATE, precision, *this.At(MkString("precisionMode")))
				} else {
					if IsTrue(OpEq2(side, MkString("sell"))) {
						*(request).At(MkString("size")) = this.AmountToPrecision(symbol, amount)
					}
				}
			}
		}
	} else {
		if IsTrue(OpOr(*(market).At(MkString("swap")), *(market).At(MkString("future")))) {
			method = MkString("privateContractPostSubmitOrder")
			*(request).At(MkString("contractID")) = *(market).At(MkString("id"))
			if IsTrue(OpEq2(type_, MkString("limit"))) {
				*(request).At(MkString("category")) = MkInteger(1)
			} else {
				if IsTrue(OpEq2(type_, MkString("market"))) {
					*(request).At(MkString("category")) = MkInteger(2)
				}
			}
			*(request).At(MkString("way")) = OpCopy(side)
			*(request).At(MkString("custom_id")) = this.Nonce()
			*(request).At(MkString("open_type")) = MkInteger(1)
			*(request).At(MkString("leverage")) = MkInteger(1)
			*(request).At(MkString("price")) = this.PriceToPrecision(symbol, price)
			*(request).At(MkString("vol")) = this.AmountToPrecision(symbol, amount)
		}
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	return this.ParseOrder(data, market)
}

func (this *Bitmart) CancelOrder(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{})
	_ = request
	method := OpCopy(MkUndefined())
	_ = method
	if IsTrue(*(market).At(MkString("spot"))) {
		method = MkString("privateSpotPostCancelOrder")
		*(request).At(MkString("order_id")) = parseInt(id)
		*(request).At(MkString("symbol")) = *(market).At(MkString("id"))
	} else {
		if IsTrue(OpOr(*(market).At(MkString("swap")), *(market).At(MkString("future")))) {
			method = MkString("privateContractPostCancelOrders")
			*(request).At(MkString("contractID")) = *(market).At(MkString("id"))
			*(request).At(MkString("orders")) = MkArray(&VarArray{
				parseInt(id),
			})
		}
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"))
	_ = data
	if IsTrue(OpEq2(data, MkBool(true))) {
		return this.ParseOrder(id, market)
	}
	succeeded := this.SafeValue(data, MkString("succeed"))
	_ = succeeded
	if IsTrue(OpNotEq2(succeeded, MkUndefined())) {
		id = this.SafeString(succeeded, MkInteger(0))
		if IsTrue(OpEq2(id, MkUndefined())) {
			panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" cancelOrder() failed to cancel "), OpAdd(symbol, OpAdd(MkString(" order id "), id))))))
		}
	} else {
		result := this.SafeValue(data, MkString("result"))
		_ = result
		if IsTrue(OpNot(result)) {
			panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" cancelOrder() "), OpAdd(symbol, OpAdd(MkString(" order id "), OpAdd(id, MkString(" is filled or canceled"))))))))
		}
	}
	order := this.ParseOrder(id, market)
	_ = order
	return this.Extend(order, MkMap(&VarMap{"id": id}))
}

func (this *Bitmart) CancelAllOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" cancelAllOrders() requires a symbol argument"))))
	}
	side := this.SafeString(params, MkString("side"))
	_ = side
	if IsTrue(OpEq2(side, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" cancelAllOrders() requires a `side` parameter ('buy' or 'sell')"))))
	}
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	if IsTrue(OpNot(*(market).At(MkString("spot")))) {
		panic(NewNotSupported(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" cancelAllOrders() does not support "), OpAdd(*(market).At(MkString("type")), MkString(" orders, only spot orders are accepted"))))))
	}
	request := MkMap(&VarMap{
		"symbol": *(market).At(MkString("id")),
		"side":   side,
	})
	_ = request
	response := this.Call(MkString("privateSpotPostCancelOrders"), this.Extend(request, params))
	_ = response
	return response
}

func (this *Bitmart) CancelOrders(goArgs ...*Variant) *Variant {
	ids := GoGetArg(goArgs, 0, MkUndefined())
	_ = ids
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" canelOrders() requires a symbol argument"))))
	}
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	if IsTrue(OpNot(*(market).At(MkString("spot")))) {
		panic(NewNotSupported(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" cancelOrders() does not support "), OpAdd(*(market).At(MkString("type")), MkString(" orders, only contract orders are accepted"))))))
	}
	orders := MkArray(&VarArray{})
	_ = orders
	for i := MkInteger(0); IsTrue(OpLw(i, ids.Length)); OpIncr(&i) {
		orders.Push(parseInt(*(ids).At(i)))
	}
	request := MkMap(&VarMap{"orders": orders})
	_ = request
	response := this.Call(MkString("privateContractPostCancelOrders"), this.Extend(request, params))
	_ = response
	return response
}

func (this *Bitmart) FetchOrdersByStatus(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchOrdersByStatus() requires a symbol argument"))))
	}
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{})
	_ = request
	method := OpCopy(MkUndefined())
	_ = method
	if IsTrue(*(market).At(MkString("spot"))) {
		method = MkString("privateSpotGetOrders")
		*(request).At(MkString("symbol")) = *(market).At(MkString("id"))
		*(request).At(MkString("offset")) = MkInteger(1)
		*(request).At(MkString("limit")) = MkInteger(100)
		if IsTrue(OpEq2(status, MkString("open"))) {
			*(request).At(MkString("status")) = MkInteger(9)
		} else {
			if IsTrue(OpEq2(status, MkString("closed"))) {
				*(request).At(MkString("status")) = MkInteger(6)
			} else {
				*(request).At(MkString("status")) = OpCopy(status)
			}
		}
	} else {
		if IsTrue(OpOr(*(market).At(MkString("swap")), *(market).At(MkString("future")))) {
			method = MkString("privateContractGetUserOrders")
			*(request).At(MkString("contractID")) = *(market).At(MkString("id"))
			if IsTrue(OpNotEq2(limit, MkUndefined())) {
				*(request).At(MkString("size")) = OpCopy(limit)
			}
			if IsTrue(OpEq2(status, MkString("open"))) {
				*(request).At(MkString("status")) = MkInteger(3)
			} else {
				if IsTrue(OpEq2(status, MkString("closed"))) {
					*(request).At(MkString("status")) = MkInteger(4)
				} else {
					*(request).At(MkString("status")) = OpCopy(status)
				}
			}
		}
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	orders := this.SafeValue(data, MkString("orders"), MkArray(&VarArray{}))
	_ = orders
	return this.ParseOrders(orders, market, since, limit)
}

func (this *Bitmart) FetchOpenOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	return this.FetchOrdersByStatus(MkString("open"), symbol, since, limit, params)
}

func (this *Bitmart) FetchClosedOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	return this.FetchOrdersByStatus(MkString("closed"), symbol, since, limit, params)
}

func (this *Bitmart) FetchOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchOrders() requires a symbol argument"))))
	}
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	if IsTrue(OpNot(OpOr(*(market).At(MkString("swap")), *(market).At(MkString("future"))))) {
		panic(NewNotSupported(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" fetchOrders does not support "), OpAdd(*(market).At(MkString("type")), MkString(" markets, only contracts are supported"))))))
	}
	return this.FetchOrdersByStatus(MkInteger(0), symbol, since, limit, params)
}

func (this *Bitmart) FetchOrder(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{})
	_ = request
	market := this.Market(symbol)
	_ = market
	method := OpCopy(MkUndefined())
	_ = method
	if IsTrue(OpNotEq2(OpType(id), MkString("string"))) {
		id = id.ToString()
	}
	if IsTrue(*(market).At(MkString("spot"))) {
		*(request).At(MkString("symbol")) = *(market).At(MkString("id"))
		*(request).At(MkString("order_id")) = OpCopy(id)
		method = MkString("privateSpotGetOrderDetail")
	} else {
		if IsTrue(OpOr(*(market).At(MkString("swap")), *(market).At(MkString("future")))) {
			*(request).At(MkString("contractID")) = *(market).At(MkString("id"))
			*(request).At(MkString("orderID")) = OpCopy(id)
			method = MkString("privateContractGetUserOrderInfo")
		}
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"))
	_ = data
	if IsTrue(OpHasMember(MkString("orders"), data)) {
		orders := this.SafeValue(data, MkString("orders"), MkArray(&VarArray{}))
		_ = orders
		firstOrder := this.SafeValue(orders, MkInteger(0))
		_ = firstOrder
		if IsTrue(OpEq2(firstOrder, MkUndefined())) {
			panic(NewOrderNotFound(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" fetchOrder() could not find "), OpAdd(symbol, OpAdd(MkString(" order id "), id))))))
		}
		return this.ParseOrder(firstOrder, market)
	} else {
		return this.ParseOrder(data, market)
	}
	return MkUndefined()
}

func (this *Bitmart) FetchDepositAddress(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{"currency": *(currency).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("privateAccountGetDepositAddress"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	address := this.SafeString(data, MkString("address"))
	_ = address
	tag := this.SafeString(data, MkString("address_memo"))
	_ = tag
	this.CheckAddress(address)
	return MkMap(&VarMap{
		"currency": code,
		"address":  address,
		"tag":      tag,
		"info":     response,
	})
}

func (this *Bitmart) Withdraw(goArgs ...*Variant) *Variant {
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
		"currency":    *(currency).At(MkString("id")),
		"amount":      amount,
		"destination": MkString("To Digital Address"),
		"address":     address,
	})
	_ = request
	if IsTrue(OpNotEq2(tag, MkUndefined())) {
		*(request).At(MkString("address_memo")) = OpCopy(tag)
	}
	response := this.Call(MkString("privateAccountPostWithdrawApply"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"))
	_ = data
	transaction := this.ParseTransaction(data, currency)
	_ = transaction
	return this.Extend(transaction, MkMap(&VarMap{
		"code":    code,
		"address": address,
		"tag":     tag,
	}))
}

func (this *Bitmart) FetchTransactionsByType(goArgs ...*Variant) *Variant {
	type_ := GoGetArg(goArgs, 0, MkUndefined())
	_ = type_
	code := GoGetArg(goArgs, 1, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 2, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 3, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 4, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	if IsTrue(OpEq2(limit, MkUndefined())) {
		limit = MkInteger(50)
	}
	request := MkMap(&VarMap{
		"operation_type": type_,
		"offset":         MkInteger(1),
		"limit":          limit,
	})
	_ = request
	currency := OpCopy(MkUndefined())
	_ = currency
	if IsTrue(OpNotEq2(code, MkUndefined())) {
		currency = this.Currency(code)
		*(request).At(MkString("currency")) = *(currency).At(MkString("id"))
	}
	response := this.Call(MkString("privateAccountGetDepositWithdrawHistory"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
	_ = data
	records := this.SafeValue(data, MkString("records"), MkArray(&VarArray{}))
	_ = records
	return this.ParseTransactions(records, currency, since, limit)
}

func (this *Bitmart) FetchDeposits(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	return this.FetchTransactionsByType(MkString("deposit"), code, since, limit, params)
}

func (this *Bitmart) FetchWithdrawals(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	return this.FetchTransactionsByType(MkString("withdraw"), code, since, limit, params)
}

func (this *Bitmart) ParseTransactionStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"0": MkString("pending"),
		"1": MkString("pending"),
		"2": MkString("pending"),
		"3": MkString("ok"),
		"4": MkString("canceled"),
		"5": MkString("failed"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Bitmart) ParseTransaction(goArgs ...*Variant) *Variant {
	transaction := GoGetArg(goArgs, 0, MkUndefined())
	_ = transaction
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	id := OpCopy(MkUndefined())
	_ = id
	withdrawId := this.SafeString(transaction, MkString("withdraw_id"))
	_ = withdrawId
	depositId := this.SafeString(transaction, MkString("deposit_id"))
	_ = depositId
	type_ := OpCopy(MkUndefined())
	_ = type_
	if IsTrue(OpAnd(OpNotEq2(withdrawId, MkUndefined()), OpNotEq2(withdrawId, MkString("")))) {
		type_ = MkString("withdraw")
		id = OpCopy(withdrawId)
	} else {
		if IsTrue(OpAnd(OpNotEq2(depositId, MkUndefined()), OpNotEq2(depositId, MkString("")))) {
			type_ = MkString("deposit")
			id = OpCopy(depositId)
		}
	}
	amount := this.SafeNumber(transaction, MkString("arrival_amount"))
	_ = amount
	timestamp := this.SafeInteger(transaction, MkString("apply_time"))
	_ = timestamp
	currencyId := this.SafeString(transaction, MkString("currency"))
	_ = currencyId
	code := this.SafeCurrencyCode(currencyId, currency)
	_ = code
	status := this.ParseTransactionStatus(this.SafeString(transaction, MkString("status")))
	_ = status
	feeCost := this.SafeNumber(transaction, MkString("fee"))
	_ = feeCost
	fee := OpCopy(MkUndefined())
	_ = fee
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		fee = MkMap(&VarMap{
			"cost":     feeCost,
			"currency": code,
		})
	}
	txid := this.SafeString(transaction, MkString("tx_id"))
	_ = txid
	if IsTrue(OpEq2(txid, MkString(""))) {
		txid = OpCopy(MkUndefined())
	}
	address := this.SafeString(transaction, MkString("address"))
	_ = address
	tag := this.SafeString(transaction, MkString("address_memo"))
	_ = tag
	return MkMap(&VarMap{
		"info":        transaction,
		"id":          id,
		"currency":    code,
		"amount":      amount,
		"address":     address,
		"addressFrom": MkUndefined(),
		"addressTo":   MkUndefined(),
		"tag":         tag,
		"tagFrom":     MkUndefined(),
		"tagTo":       MkUndefined(),
		"status":      status,
		"type":        type_,
		"updated":     MkUndefined(),
		"txid":        txid,
		"timestamp":   timestamp,
		"datetime":    this.Iso8601(timestamp),
		"fee":         fee,
	})
}

func (this *Bitmart) Nonce(goArgs ...*Variant) *Variant {
	return this.Milliseconds()
}

func (this *Bitmart) Sign(goArgs ...*Variant) *Variant {
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
	baseUrl := this.ImplodeHostname(*(*this.At(MkString("urls"))).At(MkString("api")))
	_ = baseUrl
	access := this.SafeString(api, MkInteger(0))
	_ = access
	type_ := this.SafeString(api, MkInteger(1))
	_ = type_
	url := OpAdd(baseUrl, OpAdd(MkString("/"), type_))
	_ = url
	if IsTrue(OpNotEq2(type_, MkString("system"))) {
		OpAddAssign(&url, OpAdd(MkString("/"), *this.At(MkString("version"))))
	}
	if IsTrue(OpEq2(type_, MkString("contract"))) {
		OpAddAssign(&url, OpAdd(MkString("/"), MkString("ifcontract")))
	}
	OpAddAssign(&url, OpAdd(MkString("/"), this.ImplodeParams(path, params)))
	query := this.Omit(params, this.ExtractParams(path))
	_ = query
	if IsTrue(OpEq2(type_, MkString("system"))) {
		if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
			OpAddAssign(&url, OpAdd(MkString("?"), this.Urlencode(query)))
		}
	} else {
		if IsTrue(OpEq2(access, MkString("public"))) {
			if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
				OpAddAssign(&url, OpAdd(MkString("?"), this.Urlencode(query)))
			}
		} else {
			if IsTrue(OpEq2(access, MkString("private"))) {
				this.CheckRequiredCredentials()
				timestamp := (this.Milliseconds()).Call(MkString("toString"))
				_ = timestamp
				queryString := MkString("")
				_ = queryString
				headers = MkMap(&VarMap{
					"X-BM-KEY":       *this.At(MkString("apiKey")),
					"X-BM-TIMESTAMP": timestamp,
				})
				if IsTrue(OpOr(OpEq2(method, MkString("POST")), OpEq2(method, MkString("PUT")))) {
					*(headers).At(MkString("Content-Type")) = MkString("application/json")
					body = this.Json(query)
					queryString = OpCopy(body)
				} else {
					if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
						queryString = this.Urlencode(query)
						OpAddAssign(&url, OpAdd(MkString("?"), queryString))
					}
				}
				auth := OpAdd(timestamp, OpAdd(MkString("#"), OpAdd(*this.At(MkString("uid")), OpAdd(MkString("#"), queryString))))
				_ = auth
				signature := this.Hmac(this.Encode(auth), this.Encode(*this.At(MkString("secret"))))
				_ = signature
				*(headers).At(MkString("X-BM-SIGN")) = OpCopy(signature)
			}
		}
	}
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Bitmart) HandleErrors(goArgs ...*Variant) *Variant {
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
	message := this.SafeString(response, MkString("message"))
	_ = message
	errorCode := this.SafeString(response, MkString("code"))
	_ = errorCode
	if IsTrue(OpOr(OpAnd(OpNotEq2(errorCode, MkUndefined()), OpNotEq2(errorCode, MkString("1000"))), OpAnd(OpNotEq2(message, MkUndefined()), OpNotEq2(message, MkString("OK"))))) {
		feedback := OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))
		_ = feedback
		this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("exact")), errorCode, feedback)
		this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("broad")), errorCode, feedback)
		this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("exact")), message, feedback)
		this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("broad")), message, feedback)
		panic(NewExchangeError(feedback))
	}
	return MkUndefined()
}
