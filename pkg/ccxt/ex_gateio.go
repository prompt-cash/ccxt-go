package ccxt

import ()

type Gateio struct {
	*ExchangeBase
}

var _ Exchange = (*Gateio)(nil)

func init() {
	exchange := &Gateio{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Gateio) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("gateio"),
		"name": MkString("Gate.io"),
		"country": MkArray(&VarArray{
			MkString("KR"),
		}),
		"rateLimit": MkInteger(1000),
		"version":   MkString("4"),
		"certified": MkBool(true),
		"pro":       MkBool(true),
		"urls": MkMap(&VarMap{
			"logo": MkString("https://user-images.githubusercontent.com/1294454/31784029-0313c702-b509-11e7-9ccc-bc0da6a0e435.jpg"),
			"doc":  MkString("https://www.gate.io/docs/apiv4/en/index.html"),
			"www":  MkString("https://gate.io/"),
			"api": MkMap(&VarMap{
				"public":  MkString("https://api.gateio.ws/api/v4"),
				"private": MkString("https://api.gateio.ws/api/v4"),
			}),
			"referral": MkMap(&VarMap{
				"url":      MkString("https://www.gate.io/ref/2436035"),
				"discount": MkNumber(0.2),
			}),
		}),
		"has": MkMap(&VarMap{
			"cancelOrder":       MkBool(true),
			"createOrder":       MkBool(true),
			"fetchBalance":      MkBool(true),
			"fetchClosedOrders": MkBool(true),
			"fetchCurrencies":   MkBool(true),
			"fetchDeposits":     MkBool(true),
			"fetchMarkets":      MkBool(true),
			"fetchMyTrades":     MkBool(true),
			"fetchOHLCV":        MkBool(true),
			"fetchOpenOrders":   MkBool(true),
			"fetchOrder":        MkBool(true),
			"fetchTicker":       MkBool(true),
			"fetchTickers":      MkBool(true),
			"fetchTrades":       MkBool(true),
			"fetchWithdrawals":  MkBool(true),
			"transfer":          MkBool(true),
			"withdraw":          MkBool(true),
		}),
		"api": MkMap(&VarMap{
			"public": MkMap(&VarMap{
				"spot": MkMap(&VarMap{"get": MkArray(&VarArray{
					MkString("currencies"),
					MkString("currencies/{currency}"),
					MkString("currency_pairs"),
					MkString("currency_pairs/{currency_pair}"),
					MkString("tickers"),
					MkString("order_book"),
					MkString("trades"),
					MkString("candlesticks"),
				})}),
				"margin": MkMap(&VarMap{"get": MkArray(&VarArray{
					MkString("currency_pairs"),
					MkString("currency_pairs/{currency_pair}"),
					MkString("cross/currencies"),
					MkString("cross/currencies/{currency}"),
				})}),
				"futures": MkMap(&VarMap{"get": MkArray(&VarArray{
					MkString("{settle}/contracts"),
					MkString("{settle}/contracts/{contract}"),
					MkString("{settle}/order_book"),
					MkString("{settle}/trades"),
					MkString("{settle}/candlesticks"),
					MkString("{settle}/tickers"),
					MkString("{settle}/funding_rate"),
					MkString("{settle}/insurance"),
					MkString("{settle}/contract_stats"),
					MkString("{settle}/liq_orders"),
				})}),
				"delivery": MkMap(&VarMap{"get": MkArray(&VarArray{
					MkString("{settle}/contracts"),
					MkString("{settle}/contracts/{contract}"),
					MkString("{settle}/order_book"),
					MkString("{settle}/trades"),
					MkString("{settle}/candlesticks"),
					MkString("{settle}/tickers"),
					MkString("{settle}/insurance"),
				})}),
			}),
			"private": MkMap(&VarMap{
				"withdrawals": MkMap(&VarMap{
					"post": MkArray(&VarArray{
						MkString(""),
					}),
					"delete": MkArray(&VarArray{
						MkString("{withdrawal_id}"),
					}),
				}),
				"wallet": MkMap(&VarMap{
					"get": MkArray(&VarArray{
						MkString("deposit_address"),
						MkString("withdrawals"),
						MkString("deposits"),
						MkString("sub_account_transfers"),
						MkString("withdraw_status"),
						MkString("sub_account_balances"),
						MkString("fee"),
					}),
					"post": MkArray(&VarArray{
						MkString("transfers"),
						MkString("sub_account_transfers"),
					}),
				}),
				"spot": MkMap(&VarMap{
					"get": MkArray(&VarArray{
						MkString("accounts"),
						MkString("open_orders"),
						MkString("orders"),
						MkString("orders/{order_id}"),
						MkString("my_trades"),
						MkString("price_orders"),
						MkString("price_orders/{order_id}"),
					}),
					"post": MkArray(&VarArray{
						MkString("batch_orders"),
						MkString("orders"),
						MkString("cancel_batch_orders"),
						MkString("price_orders"),
					}),
					"delete": MkArray(&VarArray{
						MkString("orders"),
						MkString("orders/{order_id}"),
						MkString("price_orders"),
						MkString("price_orders/{order_id}"),
					}),
				}),
				"margin": MkMap(&VarMap{
					"get": MkArray(&VarArray{
						MkString("account_book"),
						MkString("funding_accounts"),
						MkString("loans"),
						MkString("loans/{loan_id}"),
						MkString("loans/{loan_id}/repayment"),
						MkString("loan_records"),
						MkString("loan_records/{load_record_id}"),
						MkString("auto_repay"),
						MkString("transferable"),
						MkString("cross/accounts"),
						MkString("cross/account_book"),
						MkString("cross/loans"),
						MkString("cross/loans/{loan_id}"),
						MkString("cross/loans/repayments"),
						MkString("cross/transferable"),
					}),
					"post": MkArray(&VarArray{
						MkString("loans"),
						MkString("merged_loans"),
						MkString("loans/{loan_id}/repayment"),
						MkString("auto_repay"),
						MkString("cross/loans"),
						MkString("cross/loans/repayments"),
					}),
					"patch": MkArray(&VarArray{
						MkString("loans/{loan_id}"),
						MkString("loan_records/{loan_record_id}"),
					}),
					"delete": MkArray(&VarArray{
						MkString("loans/{loan_id}"),
					}),
				}),
				"futures": MkMap(&VarMap{
					"get": MkArray(&VarArray{
						MkString("{settle}/accounts"),
						MkString("{settle}/account_book"),
						MkString("{settle}/positions"),
						MkString("{settle}/positions/{contract}"),
						MkString("{settle}/orders"),
						MkString("{settle}/orders/{order_id}"),
						MkString("{settle}/my_trades"),
						MkString("{settle}/position_close"),
						MkString("{settle}/liquidates"),
						MkString("{settle}/price_orders"),
						MkString("{settle}/price_orders/{order_id}"),
					}),
					"post": MkArray(&VarArray{
						MkString("{settle}/positions/{contract}/margin"),
						MkString("{settle}/positions/{contract}/leverage"),
						MkString("{settle}/positions/{contract}/risk_limit"),
						MkString("{settle}/dual_mode"),
						MkString("{settle}/dual_comp/positions/{contract}"),
						MkString("{settle}/dual_comp/positions/{contract}/margin"),
						MkString("{settle}/dual_comp/positions/{contract}/leverage"),
						MkString("{settle}/dual_comp/positions/{contract}/risk_limit"),
						MkString("{settle}/orders"),
						MkString("{settle}/price_orders"),
					}),
					"delete": MkArray(&VarArray{
						MkString("{settle}/orders"),
						MkString("{settle}/orders/{order_id}"),
						MkString("{settle}/price_orders"),
						MkString("{settle}/price_orders/{order_id}"),
					}),
				}),
				"delivery": MkMap(&VarMap{
					"get": MkArray(&VarArray{
						MkString("{settle}/accounts"),
						MkString("{settle}/account_book"),
						MkString("{settle}/positions"),
						MkString("{settle}/positions/{contract}"),
						MkString("{settle}/orders"),
						MkString("{settle}/orders/{order_id}"),
						MkString("{settle}/my_trades"),
						MkString("{settle}/position_close"),
						MkString("{settle}/liquidates"),
						MkString("{settle}/price_orders"),
						MkString("{settle}/price_orders/{order_id}"),
					}),
					"post": MkArray(&VarArray{
						MkString("{settle}/positions/{contract}/margin"),
						MkString("{settle}/positions/{contract}/leverage"),
						MkString("{settle}/positions/{contract}/risk_limit"),
						MkString("{settle}/orders"),
					}),
					"delete": MkArray(&VarArray{
						MkString("{settle}/orders"),
						MkString("{settle}/orders/{order_id}"),
						MkString("{settle}/price_orders"),
						MkString("{settle}/price_orders/{order_id}"),
					}),
				}),
			}),
		}),
		"timeframes": MkMap(&VarMap{
			"10s": MkString("10s"),
			"1m":  MkString("1m"),
			"5m":  MkString("5m"),
			"15m": MkString("15m"),
			"30m": MkString("30m"),
			"1h":  MkString("1h"),
			"4h":  MkString("4h"),
			"8h":  MkString("8h"),
			"1d":  MkString("1d"),
			"7d":  MkString("7d"),
		}),
		"commonCurrencies": MkMap(&VarMap{
			"88MPH":   MkString("MPH"),
			"BIFI":    MkString("Bitcoin File"),
			"BOX":     MkString("DefiBox"),
			"BTCBEAR": MkString("BEAR"),
			"BTCBULL": MkString("BULL"),
			"BYN":     MkString("Beyond Finance"),
			"GTC":     MkString("Game.com"),
			"GTC_HT":  MkString("Game.com HT"),
			"GTC_BSC": MkString("Game.com BSC"),
			"HIT":     MkString("HitChain"),
			"MPH":     MkString("Morpher"),
			"RAI":     MkString("Rai Reflex Index"),
			"SBTC":    MkString("Super Bitcoin"),
			"TNC":     MkString("Trinity Network Credit"),
			"VAI":     MkString("VAIOT"),
		}),
		"options": MkMap(&VarMap{"accountsByType": MkMap(&VarMap{
			"spot":     MkString("spot"),
			"margin":   MkString("margin"),
			"futures":  MkString("futures"),
			"delivery": MkString("delivery"),
		})}),
		"fees": MkMap(&VarMap{"trading": MkMap(&VarMap{
			"tierBased":  MkBool(true),
			"feeSide":    MkString("get"),
			"percentage": MkBool(true),
			"maker":      this.ParseNumber(MkString("0.002")),
			"taker":      this.ParseNumber(MkString("0.002")),
			"tiers": MkMap(&VarMap{
				"maker": MkArray(&VarArray{
					MkArray(&VarArray{
						this.ParseNumber(MkString("0")),
						this.ParseNumber(MkString("0.002")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("1.5")),
						this.ParseNumber(MkString("0.00185")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("3")),
						this.ParseNumber(MkString("0.00175")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("6")),
						this.ParseNumber(MkString("0.00165")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("12.5")),
						this.ParseNumber(MkString("0.00155")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("25")),
						this.ParseNumber(MkString("0.00145")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("75")),
						this.ParseNumber(MkString("0.00135")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("200")),
						this.ParseNumber(MkString("0.00125")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("500")),
						this.ParseNumber(MkString("0.00115")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("1250")),
						this.ParseNumber(MkString("0.00105")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("2500")),
						this.ParseNumber(MkString("0.00095")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("3000")),
						this.ParseNumber(MkString("0.00085")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("6000")),
						this.ParseNumber(MkString("0.00075")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("11000")),
						this.ParseNumber(MkString("0.00065")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("20000")),
						this.ParseNumber(MkString("0.00055")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("40000")),
						this.ParseNumber(MkString("0.00055")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("75000")),
						this.ParseNumber(MkString("0.00055")),
					}),
				}),
				"taker": MkArray(&VarArray{
					MkArray(&VarArray{
						this.ParseNumber(MkString("0")),
						this.ParseNumber(MkString("0.002")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("1.5")),
						this.ParseNumber(MkString("0.00195")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("3")),
						this.ParseNumber(MkString("0.00185")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("6")),
						this.ParseNumber(MkString("0.00175")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("12.5")),
						this.ParseNumber(MkString("0.00165")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("25")),
						this.ParseNumber(MkString("0.00155")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("75")),
						this.ParseNumber(MkString("0.00145")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("200")),
						this.ParseNumber(MkString("0.00135")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("500")),
						this.ParseNumber(MkString("0.00125")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("1250")),
						this.ParseNumber(MkString("0.00115")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("2500")),
						this.ParseNumber(MkString("0.00105")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("3000")),
						this.ParseNumber(MkString("0.00095")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("6000")),
						this.ParseNumber(MkString("0.00085")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("11000")),
						this.ParseNumber(MkString("0.00075")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("20000")),
						this.ParseNumber(MkString("0.00065")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("40000")),
						this.ParseNumber(MkString("0.00065")),
					}),
					MkArray(&VarArray{
						this.ParseNumber(MkString("75000")),
						this.ParseNumber(MkString("0.00065")),
					}),
				}),
			}),
		})}),
		"exceptions": MkMap(&VarMap{
			"INVALID_PARAM_VALUE":            BadRequest,
			"INVALID_PROTOCOL":               BadRequest,
			"INVALID_ARGUMENT":               BadRequest,
			"INVALID_REQUEST_BODY":           BadRequest,
			"MISSING_REQUIRED_PARAM":         ArgumentsRequired,
			"BAD_REQUEST":                    BadRequest,
			"INVALID_CONTENT_TYPE":           BadRequest,
			"NOT_ACCEPTABLE":                 BadRequest,
			"METHOD_NOT_ALLOWED":             BadRequest,
			"NOT_FOUND":                      ExchangeError,
			"INVALID_CREDENTIALS":            AuthenticationError,
			"INVALID_KEY":                    AuthenticationError,
			"IP_FORBIDDEN":                   AuthenticationError,
			"READ_ONLY":                      PermissionDenied,
			"INVALID_SIGNATURE":              AuthenticationError,
			"MISSING_REQUIRED_HEADER":        AuthenticationError,
			"REQUEST_EXPIRED":                AuthenticationError,
			"ACCOUNT_LOCKED":                 AccountSuspended,
			"FORBIDDEN":                      PermissionDenied,
			"SUB_ACCOUNT_NOT_FOUND":          ExchangeError,
			"SUB_ACCOUNT_LOCKED":             AccountSuspended,
			"MARGIN_BALANCE_EXCEPTION":       ExchangeError,
			"MARGIN_TRANSFER_FAILED":         ExchangeError,
			"TOO_MUCH_FUTURES_AVAILABLE":     ExchangeError,
			"FUTURES_BALANCE_NOT_ENOUGH":     InsufficientFunds,
			"ACCOUNT_EXCEPTION":              ExchangeError,
			"SUB_ACCOUNT_TRANSFER_FAILED":    ExchangeError,
			"ADDRESS_NOT_USED":               ExchangeError,
			"TOO_FAST":                       RateLimitExceeded,
			"WITHDRAWAL_OVER_LIMIT":          ExchangeError,
			"API_WITHDRAW_DISABLED":          ExchangeNotAvailable,
			"INVALID_WITHDRAW_ID":            ExchangeError,
			"INVALID_WITHDRAW_CANCEL_STATUS": ExchangeError,
			"INVALID_PRECISION":              InvalidOrder,
			"INVALID_CURRENCY":               BadSymbol,
			"INVALID_CURRENCY_PAIR":          BadSymbol,
			"POC_FILL_IMMEDIATELY":           ExchangeError,
			"ORDER_NOT_FOUND":                OrderNotFound,
			"ORDER_CLOSED":                   InvalidOrder,
			"ORDER_CANCELLED":                InvalidOrder,
			"QUANTITY_NOT_ENOUGH":            InvalidOrder,
			"BALANCE_NOT_ENOUGH":             InsufficientFunds,
			"MARGIN_NOT_SUPPORTED":           InvalidOrder,
			"MARGIN_BALANCE_NOT_ENOUGH":      InsufficientFunds,
			"AMOUNT_TOO_LITTLE":              InvalidOrder,
			"AMOUNT_TOO_MUCH":                InvalidOrder,
			"REPEATED_CREATION":              InvalidOrder,
			"LOAN_NOT_FOUND":                 OrderNotFound,
			"LOAN_RECORD_NOT_FOUND":          OrderNotFound,
			"NO_MATCHED_LOAN":                ExchangeError,
			"NOT_MERGEABLE":                  ExchangeError,
			"NO_CHANGE":                      ExchangeError,
			"REPAY_TOO_MUCH":                 ExchangeError,
			"TOO_MANY_CURRENCY_PAIRS":        InvalidOrder,
			"TOO_MANY_ORDERS":                InvalidOrder,
			"MIXED_ACCOUNT_TYPE":             InvalidOrder,
			"AUTO_BORROW_TOO_MUCH":           ExchangeError,
			"TRADE_RESTRICTED":               InsufficientFunds,
			"USER_NOT_FOUND":                 ExchangeError,
			"CONTRACT_NO_COUNTER":            ExchangeError,
			"CONTRACT_NOT_FOUND":             BadSymbol,
			"RISK_LIMIT_EXCEEDED":            ExchangeError,
			"INSUFFICIENT_AVAILABLE":         InsufficientFunds,
			"LIQUIDATE_IMMEDIATELY":          InvalidOrder,
			"LEVERAGE_TOO_HIGH":              InvalidOrder,
			"LEVERAGE_TOO_LOW":               InvalidOrder,
			"ORDER_NOT_OWNED":                ExchangeError,
			"ORDER_FINISHED":                 ExchangeError,
			"POSITION_CROSS_MARGIN":          ExchangeError,
			"POSITION_IN_LIQUIDATION":        ExchangeError,
			"POSITION_IN_CLOSE":              ExchangeError,
			"POSITION_EMPTY":                 InvalidOrder,
			"REMOVE_TOO_MUCH":                ExchangeError,
			"RISK_LIMIT_NOT_MULTIPLE":        ExchangeError,
			"RISK_LIMIT_TOO_HIGH":            ExchangeError,
			"RISK_LIMIT_TOO_lOW":             ExchangeError,
			"PRICE_TOO_DEVIATED":             InvalidOrder,
			"SIZE_TOO_LARGE":                 InvalidOrder,
			"SIZE_TOO_SMALL":                 InvalidOrder,
			"PRICE_OVER_LIQUIDATION":         InvalidOrder,
			"PRICE_OVER_BANKRUPT":            InvalidOrder,
			"ORDER_POC_IMMEDIATE":            InvalidOrder,
			"INCREASE_POSITION":              InvalidOrder,
			"CONTRACT_IN_DELISTING":          ExchangeError,
			"INTERNAL":                       ExchangeError,
			"SERVER_ERROR":                   ExchangeError,
			"TOO_BUSY":                       ExchangeNotAvailable,
		}),
	}))
}

func (this *Gateio) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicSpotGetCurrencyPairs"), params)
	_ = response
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
		entry := *(response).At(i)
		_ = entry
		id := this.SafeString(entry, MkString("id"))
		_ = id
		baseId := this.SafeString(entry, MkString("base"))
		_ = baseId
		quoteId := this.SafeString(entry, MkString("quote"))
		_ = quoteId
		base := this.SafeCurrencyCode(baseId)
		_ = base
		quote := this.SafeCurrencyCode(quoteId)
		_ = quote
		symbol := OpAdd(base, OpAdd(MkString("/"), quote))
		_ = symbol
		taker := OpDiv(this.SafeNumber(entry, MkString("fee")), MkInteger(100))
		_ = taker
		maker := OpCopy(taker)
		_ = maker
		tradeStatus := this.SafeString(entry, MkString("trade_status"))
		_ = tradeStatus
		active := OpEq2(tradeStatus, MkString("tradable"))
		_ = active
		amountPrecision := this.SafeString(entry, MkString("amount_precision"))
		_ = amountPrecision
		pricePrecision := this.SafeString(entry, MkString("precision"))
		_ = pricePrecision
		amountLimit := this.ParsePrecision(amountPrecision)
		_ = amountLimit
		priceLimit := this.ParsePrecision(pricePrecision)
		_ = priceLimit
		limits := MkMap(&VarMap{
			"amount": MkMap(&VarMap{
				"min": this.ParseNumber(amountLimit),
				"max": MkUndefined(),
			}),
			"price": MkMap(&VarMap{
				"min": this.ParseNumber(priceLimit),
				"max": MkUndefined(),
			}),
			"cost": MkMap(&VarMap{
				"min": this.SafeNumber(entry, MkString("min_quote_amount")),
				"max": MkUndefined(),
			}),
		})
		_ = limits
		precision := MkMap(&VarMap{
			"amount": parseInt(amountPrecision),
			"price":  parseInt(pricePrecision),
		})
		_ = precision
		result.Push(MkMap(&VarMap{
			"info":      entry,
			"id":        id,
			"baseId":    baseId,
			"quoteId":   quoteId,
			"base":      base,
			"quote":     quote,
			"symbol":    symbol,
			"limits":    limits,
			"precision": precision,
			"active":    active,
			"maker":     maker,
			"taker":     taker,
		}))
	}
	return result
}

func (this *Gateio) FetchCurrencies(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("publicSpotGetCurrencies"), params)
	_ = response
	result := MkMap(&VarMap{})
	_ = result
	amountPrecision := MkInteger(6)
	_ = amountPrecision
	for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
		entry := *(response).At(i)
		_ = entry
		currencyId := this.SafeString(entry, MkString("currency"))
		_ = currencyId
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		delisted := this.SafeValue(entry, MkString("delisted"))
		_ = delisted
		withdraw_disabled := this.SafeValue(entry, MkString("withdraw_disabled"))
		_ = withdraw_disabled
		deposit_disabled := this.SafeValue(entry, MkString("disabled_disabled"))
		_ = deposit_disabled
		trade_disabled := this.SafeValue(entry, MkString("trade_disabled"))
		_ = trade_disabled
		active := OpNot(OpAnd(delisted, OpAnd(withdraw_disabled, OpAnd(deposit_disabled, trade_disabled))))
		_ = active
		*(result).At(code) = MkMap(&VarMap{
			"id":   currencyId,
			"name": MkUndefined(),
			"code": code,
			"precision": MkMap(&VarMap{
				"amount": amountPrecision,
				"price":  MkUndefined(),
			}),
			"info":   entry,
			"active": active,
			"fee":    MkUndefined(),
			"fees":   MkArray(&VarArray{}),
			"limits": *this.At(MkString("limits")),
		})
	}
	return result
}

func (this *Gateio) FetchNetworkDepositAddress(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{"currency": *(currency).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("privateWalletGetDepositAddress"), this.Extend(request, params))
	_ = response
	addresses := this.SafeValue(response, MkString("multichain_addresses"))
	_ = addresses
	currencyId := this.SafeString(response, MkString("currency"))
	_ = currencyId
	code = this.SafeCurrencyCode(currencyId)
	result := MkMap(&VarMap{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, addresses.Length)); OpIncr(&i) {
		entry := *(addresses).At(i)
		_ = entry
		obtainFailed := this.SafeInteger(entry, MkString("obtain_failed"))
		_ = obtainFailed
		if IsTrue(obtainFailed) {
			continue
		}
		network := this.SafeString(entry, MkString("chain"))
		_ = network
		address := this.SafeString(entry, MkString("address"))
		_ = address
		tag := this.SafeString(entry, MkString("payment_id"))
		_ = tag
		tagLength := OpCopy(tag.Length)
		_ = tagLength
		tag = OpTrinary(tagLength, tag, MkUndefined())
		*(result).At(network) = MkMap(&VarMap{
			"info":    entry,
			"code":    code,
			"address": address,
			"tag":     tag,
		})
	}
	return result
}

func (this *Gateio) FetchDepositAddress(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	currency := this.Currency(code)
	_ = currency
	request := MkMap(&VarMap{"currency": *(currency).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("privateWalletGetDepositAddress"), this.Extend(request, params))
	_ = response
	currencyId := this.SafeString(response, MkString("currency"))
	_ = currencyId
	code = this.SafeCurrencyCode(currencyId)
	addressField := this.SafeString(response, MkString("address"))
	_ = addressField
	tag := OpCopy(MkUndefined())
	_ = tag
	address := OpCopy(MkUndefined())
	_ = address
	if IsTrue(OpGt(addressField.IndexOf(MkString(" ")), OpNeg(MkInteger(1)))) {
		splitted := addressField.Split(MkString(" "))
		_ = splitted
		address = *(splitted).At(MkInteger(0))
		tag = *(splitted).At(MkInteger(1))
	} else {
		address = OpCopy(addressField)
	}
	return MkMap(&VarMap{
		"info":    response,
		"code":    code,
		"address": address,
		"tag":     tag,
	})
}

func (this *Gateio) FetchTradingFees(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privateWalletGetFee"), params)
	_ = response
	result := MkMap(&VarMap{})
	_ = result
	taker := this.SafeNumber(response, MkString("taker_fee"))
	_ = taker
	maker := this.SafeNumber(response, MkString("maker_fee"))
	_ = maker
	for i := MkInteger(0); IsTrue(OpLw(i, (*((this).At(MkString("symbols")))).Length)); OpIncr(&i) {
		symbol := *(*this.At(MkString("symbols"))).At(i)
		_ = symbol
		*(result).At(symbol) = MkMap(&VarMap{
			"maker":  maker,
			"taker":  taker,
			"info":   response,
			"symbol": symbol,
		})
	}
	return result
}

func (this *Gateio) FetchFundingFees(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privateWalletGetWithdrawStatus"), params)
	_ = response
	withdrawFees := MkMap(&VarMap{})
	_ = withdrawFees
	for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
		entry := *(response).At(i)
		_ = entry
		currencyId := this.SafeString(entry, MkString("currency"))
		_ = currencyId
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		*(withdrawFees).At(code) = MkMap(&VarMap{})
		withdrawFix := this.SafeValue(entry, MkString("withdraw_fix_on_chains"))
		_ = withdrawFix
		if IsTrue(OpEq2(withdrawFix, MkUndefined())) {
			withdrawFix = MkMap(&VarMap{})
			*(withdrawFix).At(code) = this.SafeNumber(entry, MkString("withdraw_fix"))
		}
		keys := GoGetKeys(withdrawFix)
		_ = keys
		for i := MkInteger(0); IsTrue(OpLw(i, keys.Length)); OpIncr(&i) {
			key := *(keys).At(i)
			_ = key
			*(*(withdrawFees).At(code)).At(key) = this.ParseNumber(*(withdrawFix).At(key))
		}
	}
	return MkMap(&VarMap{
		"info":     response,
		"withdraw": withdrawFees,
		"deposit":  MkMap(&VarMap{}),
	})
}

func (this *Gateio) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"currency_pair": *(market).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("publicSpotGetOrderBook"), this.Extend(request, params))
	_ = response
	timestamp := this.SafeInteger(response, MkString("current"))
	_ = timestamp
	return this.ParseOrderBook(response, symbol, timestamp)
}

func (this *Gateio) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"currency_pair": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("publicSpotGetTickers"), this.Extend(request, params))
	_ = response
	ticker := this.SafeValue(response, MkInteger(0))
	_ = ticker
	return this.ParseTicker(ticker, market)
}

func (this *Gateio) ParseTicker(goArgs ...*Variant) *Variant {
	ticker := GoGetArg(goArgs, 0, MkUndefined())
	_ = ticker
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	marketId := this.SafeString(ticker, MkString("currency_pair"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market)
	_ = symbol
	last := this.SafeNumber(ticker, MkString("last"))
	_ = last
	ask := this.SafeNumber(ticker, MkString("lowest_ask"))
	_ = ask
	bid := this.SafeNumber(ticker, MkString("highest_bid"))
	_ = bid
	high := this.SafeNumber(ticker, MkString("high_24h"))
	_ = high
	low := this.SafeNumber(ticker, MkString("low_24h"))
	_ = low
	baseVolume := this.SafeNumber(ticker, MkString("base_volume"))
	_ = baseVolume
	quoteVolume := this.SafeNumber(ticker, MkString("quote_volume"))
	_ = quoteVolume
	percentage := this.SafeNumber(ticker, MkString("change_percentage"))
	_ = percentage
	return MkMap(&VarMap{
		"symbol":        symbol,
		"timestamp":     MkUndefined(),
		"datetime":      MkUndefined(),
		"high":          high,
		"low":           low,
		"bid":           bid,
		"bidVolume":     MkUndefined(),
		"ask":           ask,
		"askVolume":     MkUndefined(),
		"vwap":          MkUndefined(),
		"open":          MkUndefined(),
		"close":         last,
		"last":          last,
		"previousClose": MkUndefined(),
		"change":        MkUndefined(),
		"percentage":    percentage,
		"average":       MkUndefined(),
		"baseVolume":    baseVolume,
		"quoteVolume":   quoteVolume,
		"info":          ticker,
	})
}

func (this *Gateio) FetchTickers(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("publicSpotGetTickers"), params)
	_ = response
	return this.ParseTickers(response, symbols)
}

func (this *Gateio) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("privateSpotGetAccounts"), params)
	_ = response
	result := MkMap(&VarMap{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
		entry := *(response).At(i)
		_ = entry
		account := this.Account()
		_ = account
		currencyId := this.SafeString(entry, MkString("currency"))
		_ = currencyId
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		*(account).At(MkString("used")) = this.SafeString(entry, MkString("locked"))
		*(account).At(MkString("free")) = this.SafeString(entry, MkString("available"))
		*(result).At(code) = OpCopy(account)
	}
	return this.ParseBalance(result)
}

func (this *Gateio) FetchOHLCV(goArgs ...*Variant) *Variant {
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
		"currency_pair": *(market).At(MkString("id")),
		"interval":      *(*this.At(MkString("timeframes"))).At(timeframe),
	})
	_ = request
	if IsTrue(OpEq2(since, MkUndefined())) {
		if IsTrue(OpNotEq2(limit, MkUndefined())) {
			*(request).At(MkString("limit")) = OpCopy(limit)
		}
	} else {
		*(request).At(MkString("from")) = Math.Floor(OpDiv(since, MkInteger(1000)))
		if IsTrue(OpNotEq2(limit, MkUndefined())) {
			*(request).At(MkString("to")) = this.Sum(*(request).At(MkString("from")), OpSub(OpMulti(limit, this.ParseTimeframe(timeframe)), MkInteger(1)))
		}
	}
	response := this.Call(MkString("publicSpotGetCandlesticks"), this.Extend(request, params))
	_ = response
	return this.ParseOHLCVs(response, market, timeframe, since, limit)
}

func (this *Gateio) ParseOHLCV(goArgs ...*Variant) *Variant {
	ohlcv := GoGetArg(goArgs, 0, MkUndefined())
	_ = ohlcv
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.SafeTimestamp(ohlcv, MkInteger(0))
	_ = timestamp
	volume := this.SafeNumber(ohlcv, MkInteger(1))
	_ = volume
	close := this.SafeNumber(ohlcv, MkInteger(2))
	_ = close
	high := this.SafeNumber(ohlcv, MkInteger(3))
	_ = high
	low := this.SafeNumber(ohlcv, MkInteger(4))
	_ = low
	open := this.SafeNumber(ohlcv, MkInteger(5))
	_ = open
	return MkArray(&VarArray{
		timestamp,
		open,
		high,
		low,
		close,
		volume,
	})
}

func (this *Gateio) FetchTrades(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"currency_pair": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("publicSpotGetTrades"), this.Extend(request, params))
	_ = response
	return this.ParseTrades(response, market, since, limit)
}

func (this *Gateio) FetchMyTrades(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"currency_pair": *(market).At(MkString("id"))})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	response := this.Call(MkString("privateSpotGetMyTrades"), this.Extend(request, params))
	_ = response
	return this.ParseTrades(response, market, since, limit)
}

func (this *Gateio) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	id := this.SafeString(trade, MkString("id"))
	_ = id
	timestampString := this.SafeString2(trade, MkString("create_time_ms"), MkString("time"))
	_ = timestampString
	timestamp := OpCopy(MkUndefined())
	_ = timestamp
	if IsTrue(OpGt(timestampString.IndexOf(MkString(".")), MkInteger(0))) {
		milliseconds := timestampString.Split(MkString("."))
		_ = milliseconds
		timestamp = parseInt(*(milliseconds).At(MkInteger(0)))
	}
	marketId := this.SafeString(trade, MkString("currency_pair"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market)
	_ = symbol
	amountString := this.SafeString(trade, MkString("amount"))
	_ = amountString
	priceString := this.SafeString(trade, MkString("price"))
	_ = priceString
	cost := this.ParseNumber(Precise.StringMul(amountString, priceString))
	_ = cost
	amount := this.ParseNumber(amountString)
	_ = amount
	price := this.ParseNumber(priceString)
	_ = price
	side := this.SafeString(trade, MkString("side"))
	_ = side
	orderId := this.SafeString(trade, MkString("order_id"))
	_ = orderId
	gtFee := this.SafeString(trade, MkString("gt_fee"))
	_ = gtFee
	feeCurrency := OpCopy(MkUndefined())
	_ = feeCurrency
	feeCost := OpCopy(MkUndefined())
	_ = feeCost
	if IsTrue(OpEq2(gtFee, MkString("0"))) {
		feeCurrency = this.SafeString(trade, MkString("fee_currency"))
		feeCost = this.SafeNumber(trade, MkString("fee"))
	} else {
		feeCurrency = MkString("GT")
		feeCost = this.ParseNumber(gtFee)
	}
	fee := MkMap(&VarMap{
		"cost":     feeCost,
		"currency": feeCurrency,
	})
	_ = fee
	takerOrMaker := this.SafeString(trade, MkString("role"))
	_ = takerOrMaker
	return MkMap(&VarMap{
		"info":         trade,
		"id":           id,
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       symbol,
		"order":        orderId,
		"type":         MkUndefined(),
		"side":         side,
		"takerOrMaker": takerOrMaker,
		"price":        price,
		"amount":       amount,
		"cost":         cost,
		"fee":          fee,
	})
}

func (this *Gateio) FetchDeposits(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("from")) = Math.Floor(OpDiv(since, MkInteger(1000)))
	}
	response := this.Call(MkString("privateWalletGetDeposits"), this.Extend(request, params))
	_ = response
	return this.ParseTransactions(response, currency)
}

func (this *Gateio) FetchWithdrawals(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("from")) = Math.Floor(OpDiv(since, MkInteger(1000)))
	}
	response := this.Call(MkString("privateWalletGetWithdrawals"), this.Extend(request, params))
	_ = response
	return this.ParseTransactions(response, currency)
}

func (this *Gateio) Withdraw(goArgs ...*Variant) *Variant {
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
		"amount":   this.CurrencyToPrecision(code, amount),
	})
	_ = request
	if IsTrue(OpNotEq2(tag, MkUndefined())) {
		*(request).At(MkString("memo")) = OpCopy(tag)
	}
	response := this.Call(MkString("privateWithdrawalsPost"), this.Extend(request, params))
	_ = response
	currencyId := this.SafeString(response, MkString("currency"))
	_ = currencyId
	id := this.SafeString(response, MkString("id"))
	_ = id
	return MkMap(&VarMap{
		"info":    response,
		"id":      id,
		"code":    this.SafeCurrencyCode(currencyId),
		"amount":  this.SafeNumber(response, MkString("amount")),
		"address": this.SafeString(response, MkString("address")),
		"tag":     this.SafeString(response, MkString("memo")),
	})
}

func (this *Gateio) ParseTransactionStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"PEND":    MkString("pending"),
		"REQUEST": MkString("pending"),
		"DMOVE":   MkString("pending"),
		"CANCEL":  MkString("failed"),
		"DONE":    MkString("ok"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Gateio) ParseTransactionType(goArgs ...*Variant) *Variant {
	type_ := GoGetArg(goArgs, 0, MkUndefined())
	_ = type_
	types := MkMap(&VarMap{
		"d": MkString("deposit"),
		"w": MkString("withdrawal"),
	})
	_ = types
	return this.SafeString(types, type_, type_)
}

func (this *Gateio) ParseTransaction(goArgs ...*Variant) *Variant {
	transaction := GoGetArg(goArgs, 0, MkUndefined())
	_ = transaction
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	id := this.SafeString(transaction, MkString("id"))
	_ = id
	type_ := OpCopy(MkUndefined())
	_ = type_
	if IsTrue(OpNotEq2(id, MkUndefined())) {
		type_ = this.ParseTransactionType(*(id).At(MkInteger(0)))
	}
	currencyId := this.SafeString(transaction, MkString("currency"))
	_ = currencyId
	code := this.SafeCurrencyCode(currencyId)
	_ = code
	amount := this.SafeNumber(transaction, MkString("amount"))
	_ = amount
	txid := this.SafeString(transaction, MkString("txid"))
	_ = txid
	rawStatus := this.SafeString(transaction, MkString("status"))
	_ = rawStatus
	status := this.ParseTransactionStatus(rawStatus)
	_ = status
	address := this.SafeString(transaction, MkString("address"))
	_ = address
	fee := this.SafeNumber(transaction, MkString("fee"))
	_ = fee
	tag := this.SafeString(transaction, MkString("memo"))
	_ = tag
	if IsTrue(OpEq2(tag, MkString(""))) {
		tag = OpCopy(MkUndefined())
	}
	timestamp := this.SafeTimestamp(transaction, MkString("timestamp"))
	_ = timestamp
	return MkMap(&VarMap{
		"info":      transaction,
		"id":        id,
		"txid":      txid,
		"currency":  code,
		"amount":    amount,
		"address":   address,
		"tag":       tag,
		"status":    status,
		"type":      type_,
		"timestamp": timestamp,
		"datetime":  this.Iso8601(timestamp),
		"fee":       fee,
	})
}

func (this *Gateio) CreateOrder(goArgs ...*Variant) *Variant {
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
		"currency_pair": *(market).At(MkString("id")),
		"amount":        this.AmountToPrecision(symbol, amount),
		"price":         this.PriceToPrecision(symbol, price),
		"side":          side,
	})
	_ = request
	response := this.Call(MkString("privateSpotPostOrders"), this.Extend(request, params))
	_ = response
	return this.ParseOrder(response, market)
}

func (this *Gateio) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	id := this.SafeString(order, MkString("id"))
	_ = id
	marketId := this.SafeString(order, MkString("currency_pair"))
	_ = marketId
	symbol := this.SafeSymbol(marketId, market)
	_ = symbol
	timestamp := this.SafeTimestamp(order, MkString("create_time"))
	_ = timestamp
	timestamp = this.SafeInteger(order, MkString("create_time_ms"), timestamp)
	lastTradeTimestamp := this.SafeTimestamp(order, MkString("update_time"))
	_ = lastTradeTimestamp
	lastTradeTimestamp = this.SafeInteger(order, MkString("update_time_ms"), lastTradeTimestamp)
	amount := this.SafeNumber(order, MkString("amount"))
	_ = amount
	price := this.SafeNumber(order, MkString("price"))
	_ = price
	remaining := this.SafeNumber(order, MkString("left"))
	_ = remaining
	cost := this.SafeNumber(order, MkString("filled_total"))
	_ = cost
	side := this.SafeString(order, MkString("side"))
	_ = side
	type_ := this.SafeString(order, MkString("type"))
	_ = type_
	status := this.SafeString(order, MkString("status"))
	_ = status
	if IsTrue(OpEq2(status, MkString("cancelled"))) {
		status = MkString("canceled")
	}
	timeInForce := this.SafeStringUpper(order, MkString("time_in_force"))
	_ = timeInForce
	fees := MkArray(&VarArray{})
	_ = fees
	fees.Push(MkMap(&VarMap{
		"currency": MkString("GT"),
		"cost":     this.SafeNumber(order, MkString("gt_fee")),
	}))
	fees.Push(MkMap(&VarMap{
		"currency": this.SafeCurrencyCode(this.SafeString(order, MkString("fee_currency"))),
		"cost":     this.SafeNumber(order, MkString("fee")),
	}))
	rebate := this.SafeString(order, MkString("rebated_fee"))
	_ = rebate
	fees.Push(MkMap(&VarMap{
		"currency": this.SafeCurrencyCode(this.SafeString(order, MkString("rebated_fee_currency"))),
		"cost":     this.ParseNumber(Precise.StringNeg(rebate)),
	}))
	return this.SafeOrder(MkMap(&VarMap{
		"id":                 id,
		"clientOrderId":      id,
		"timestamp":          timestamp,
		"datetime":           this.Iso8601(timestamp),
		"lastTradeTimestamp": lastTradeTimestamp,
		"status":             status,
		"symbol":             symbol,
		"type":               type_,
		"timeInForce":        timeInForce,
		"postOnly":           MkUndefined(),
		"side":               side,
		"price":              price,
		"stopPrice":          MkUndefined(),
		"average":            MkUndefined(),
		"amount":             amount,
		"cost":               cost,
		"filled":             MkUndefined(),
		"remaining":          remaining,
		"fee":                MkUndefined(),
		"fees":               fees,
		"trades":             MkUndefined(),
		"info":               order,
	}))
}

func (this *Gateio) FetchOrder(goArgs ...*Variant) *Variant {
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
		"order_id":      id,
		"currency_pair": *(market).At(MkString("id")),
	})
	_ = request
	response := this.Call(MkString("privateSpotGetOrdersOrderId"), this.Extend(request, params))
	_ = response
	return this.ParseOrder(response, market)
}

func (this *Gateio) FetchOpenOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		request := MkMap(&VarMap{})
		_ = request
		if IsTrue(OpNotEq2(limit, MkUndefined())) {
			*(request).At(MkString("limit")) = OpCopy(limit)
		}
		response := this.Call(MkString("privateSpotGetOpenOrders"), this.Extend(request, params))
		_ = response
		allOrders := MkArray(&VarArray{})
		_ = allOrders
		for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
			entry := *(response).At(i)
			_ = entry
			orders := this.SafeValue(entry, MkString("orders"), MkArray(&VarArray{}))
			_ = orders
			parsed := this.ParseOrders(orders, MkUndefined(), since, limit)
			_ = parsed
			allOrders = this.ArrayConcat(allOrders, parsed)
		}
		return this.FilterBySinceLimit(allOrders, since, limit)
	}
	return this.FetchOrdersByStatus(MkString("open"), symbol, since, limit, params)
}

func (this *Gateio) FetchClosedOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	return this.FetchOrdersByStatus(MkString("finished"), symbol, since, limit, params)
}

func (this *Gateio) FetchOrdersByStatus(goArgs ...*Variant) *Variant {
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
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchOrdersByStatus requires a symbol argument"))))
	}
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{
		"currency_pair": *(market).At(MkString("id")),
		"status":        status,
	})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("start")) = Math.Floor(OpDiv(since, MkInteger(1000)))
	}
	response := this.Call(MkString("privateSpotGetOrders"), this.Extend(request, params))
	_ = response
	return this.ParseOrders(response, market, since, limit)
}

func (this *Gateio) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" cancelOrders requires a symbol parameter"))))
	}
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{
		"order_id":      id,
		"currency_pair": *(market).At(MkString("id")),
	})
	_ = request
	response := this.Call(MkString("privateSpotDeleteOrdersOrderId"), this.Extend(request, params))
	_ = response
	return this.ParseOrder(response)
}

func (this *Gateio) Transfer(goArgs ...*Variant) *Variant {
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
	truncated := this.CurrencyToPrecision(code, amount)
	_ = truncated
	request := MkMap(&VarMap{
		"currency": *(currency).At(MkString("id")),
		"from":     fromId,
		"to":       toId,
		"amount":   truncated,
	})
	_ = request
	if IsTrue(OpOr(OpEq2(toId, MkString("futures")), OpEq2(toId, MkString("delivery")))) {
		*(request).At(MkString("settle")) = *(currency).At(MkString("id"))
	}
	response := this.Call(MkString("privateWalletPostTransfers"), this.Extend(request, params))
	_ = response
	return MkMap(&VarMap{
		"info":   response,
		"from":   fromId,
		"to":     toId,
		"amount": truncated,
		"code":   code,
	})
}

func (this *Gateio) Sign(goArgs ...*Variant) *Variant {
	path := GoGetArg(goArgs, 0, MkUndefined())
	_ = path
	api := GoGetArg(goArgs, 1, MkArray(&VarArray{}))
	_ = api
	method := GoGetArg(goArgs, 2, MkString("GET"))
	_ = method
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	headers := GoGetArg(goArgs, 4, MkUndefined())
	_ = headers
	body := GoGetArg(goArgs, 5, MkUndefined())
	_ = body
	authentication := *(api).At(MkInteger(0))
	_ = authentication
	type_ := *(api).At(MkInteger(1))
	_ = type_
	query := this.Omit(params, this.ExtractParams(path))
	_ = query
	path = this.ImplodeParams(path, params)
	endPart := OpTrinary(OpEq2(path, MkString("")), MkString(""), OpAdd(MkString("/"), path))
	_ = endPart
	entirePath := OpAdd(MkString("/"), OpAdd(type_, endPart))
	_ = entirePath
	url := OpAdd(*(*(*this.At(MkString("urls"))).At(MkString("api"))).At(authentication), entirePath)
	_ = url
	queryString := MkString("")
	_ = queryString
	if IsTrue(OpEq2(authentication, MkString("public"))) {
		queryString = this.Urlencode(query)
		if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
			OpAddAssign(&url, OpAdd(MkString("?"), queryString))
		}
	} else {
		if IsTrue(OpOr(OpEq2(method, MkString("GET")), OpEq2(method, MkString("DELETE")))) {
			queryString = this.Urlencode(query)
			if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
				OpAddAssign(&url, OpAdd(MkString("?"), queryString))
			}
		} else {
			body = this.Json(query)
		}
		bodyPayload := OpTrinary(OpEq2(body, MkUndefined()), MkString(""), body)
		_ = bodyPayload
		bodySignature := this.Hash(this.Encode(bodyPayload), MkString("sha512"))
		_ = bodySignature
		timestamp := this.Seconds()
		_ = timestamp
		timestampString := timestamp.ToString()
		_ = timestampString
		signaturePath := OpAdd(MkString("/api/v4"), entirePath)
		_ = signaturePath
		payloadArray := MkArray(&VarArray{
			method.ToUpperCase(),
			signaturePath,
			queryString,
			bodySignature,
			timestampString,
		})
		_ = payloadArray
		payload := payloadArray.Join(MkString("\n"))
		_ = payload
		signature := this.Hmac(this.Encode(payload), this.Encode(*this.At(MkString("secret"))), MkString("sha512"))
		_ = signature
		headers = MkMap(&VarMap{
			"KEY":          *this.At(MkString("apiKey")),
			"Timestamp":    timestampString,
			"SIGN":         signature,
			"Content-Type": MkString("application/json"),
		})
	}
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Gateio) HandleErrors(goArgs ...*Variant) *Variant {
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
	label := this.SafeString(response, MkString("label"))
	_ = label
	if IsTrue(OpNotEq2(label, MkUndefined())) {
		message := this.SafeString(response, MkString("message"))
		_ = message
		Error := this.SafeValue(*this.At(MkString("exceptions")), label, ExchangeError)
		_ = Error
		panic(NewError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), message))))
	}
	return MkUndefined()
}
