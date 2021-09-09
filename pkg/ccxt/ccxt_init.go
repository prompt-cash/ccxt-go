package ccxt

//
// This file implements the initializarion and consruction of a exchange object, it holds default values and alike
//

// "constant" values

var DECIMAL_PLACES *Variant = MkInteger(0)
var SIGNIFICANT_DIGITS *Variant = MkInteger(1)
var TICK_SIZE *Variant = MkInteger(2)

var TRUNCATE *Variant = MkInteger(1)

var NO_PADDING = MkInteger(0)
var PAD_WITH_ZERO *Variant = MkInteger(1)

var ROUND = MkInteger(0)

func (this *ExchangeBase) BaseDescribe() *Variant {
	return MkMap(&VarMap{
		"id":              MkUndefined(),
		"name":            MkUndefined(),
		"countries":       MkUndefined(),
		"enableRateLimit": MkBool(true),
		"rateLimit":       MkInteger(2000),
		"certified":       MkBool(false),
		"pro":             MkBool(false),
		"has": MkMap(&VarMap{
			"loadMarkets":          MkBool(true),
			"cancelAllOrders":      MkBool(false),
			"cancelOrder":          MkBool(true),
			"cancelOrders":         MkBool(false),
			"CORS":                 MkBool(false),
			"createDepositAddress": MkBool(false),
			"createLimitOrder":     MkBool(true),
			"createMarketOrder":    MkBool(true),
			"createOrder":          MkBool(true),
			"deposit":              MkBool(false),
			"editOrder":            MkString("emulated"),
			"fetchBalance":         MkBool(true),
			"fetchBidsAsks":        MkBool(false),
			"fetchClosedOrders":    MkBool(false),
			"fetchCurrencies":      MkBool(false),
			"fetchDepositAddress":  MkBool(false),
			"fetchDeposits":        MkBool(false),
			"fetchFundingFees":     MkBool(false),
			"fetchL2OrderBook":     MkBool(true),
			"fetchLedger":          MkBool(false),
			"fetchMarkets":         MkBool(true),
			"fetchMyTrades":        MkBool(false),
			"fetchOHLCV":           MkString("emulated"),
			"fetchOpenOrders":      MkBool(false),
			"fetchOrder":           MkBool(false),
			"fetchOrderBook":       MkBool(true),
			"fetchOrderBooks":      MkBool(false),
			"fetchOrders":          MkBool(false),
			"fetchOrderTrades":     MkBool(false),
			"fetchStatus":          MkString("emulated"),
			"fetchTicker":          MkBool(true),
			"fetchTickers":         MkBool(false),
			"fetchTime":            MkBool(false),
			"fetchTrades":          MkBool(true),
			"fetchTradingFee":      MkBool(false),
			"fetchTradingFees":     MkBool(false),
			"fetchTradingLimits":   MkBool(false),
			"fetchTransactions":    MkBool(false),
			"fetchWithdrawals":     MkBool(false),
			"privateAPI":           MkBool(true),
			"publicAPI":            MkBool(true),
			"signIn":               MkBool(false),
			"withdraw":             MkBool(false),
		}),
		"urls": MkMap(&VarMap{
			"logo": MkUndefined(),
			"api":  MkUndefined(),
			"www":  MkUndefined(),
			"doc":  MkUndefined(),
			"fees": MkUndefined(),
		}),
		"api": MkUndefined(),
		"requiredCredentials": MkMap(&VarMap{
			"apiKey":        MkBool(true),
			"secret":        MkBool(true),
			"uid":           MkBool(false),
			"login":         MkBool(false),
			"password":      MkBool(false),
			"twofa":         MkBool(false),
			"privateKey":    MkBool(false),
			"walletAddress": MkBool(false),
			"token":         MkBool(false),
		}),
		"markets":    MkUndefined(),
		"currencies": MkMap(&VarMap{}),
		"timeframes": MkUndefined(),
		"fees": MkMap(&VarMap{
			"trading": MkMap(&VarMap{
				"tierBased":  MkUndefined(),
				"percentage": MkUndefined(),
				"taker":      MkUndefined(),
				"maker":      MkUndefined(),
			}),
			"funding": MkMap(&VarMap{
				"tierBased":  MkUndefined(),
				"percentage": MkUndefined(),
				"withdraw":   MkMap(&VarMap{}),
				"deposit":    MkMap(&VarMap{}),
			}),
		}),
		"status": MkMap(&VarMap{
			"status":  MkString("ok"),
			"updated": MkUndefined(),
			"eta":     MkUndefined(),
			"url":     MkUndefined(),
		}),
		"exceptions": MkUndefined(),
		"httpExceptions": MkMap(&VarMap{
			"422": ExchangeError,
			"418": DDoSProtection,
			"429": RateLimitExceeded,
			"404": ExchangeNotAvailable,
			"409": ExchangeNotAvailable,
			"410": ExchangeNotAvailable,
			"500": ExchangeNotAvailable,
			"501": ExchangeNotAvailable,
			"502": ExchangeNotAvailable,
			"520": ExchangeNotAvailable,
			"521": ExchangeNotAvailable,
			"522": ExchangeNotAvailable,
			"525": ExchangeNotAvailable,
			"526": ExchangeNotAvailable,
			"400": ExchangeNotAvailable,
			"403": ExchangeNotAvailable,
			"405": ExchangeNotAvailable,
			"503": ExchangeNotAvailable,
			"530": ExchangeNotAvailable,
			"408": RequestTimeout,
			"504": RequestTimeout,
			"401": AuthenticationError,
			"511": AuthenticationError,
		}),
		"dontGetUsedBalanceFromStaleCache": MkBool(false),
		"commonCurrencies": MkMap(&VarMap{
			"XBT":    MkString("BTC"),
			"BCC":    MkString("BCH"),
			"DRK":    MkString("DASH"),
			"BCHABC": MkString("BCH"),
			"BCHSV":  MkString("BSV"),
		}),
		"precisionMode": DECIMAL_PLACES,
		"paddingMode":   NO_PADDING,
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
	})
}

func (this *ExchangeBase) Setup(userConfig *Variant, itf interface{}) {

	values := VarMap{}

	values["options"] = MkMap(&VarMap{}) // exchange-specific options, if any

	values["fetchOptions"] = MkMap(&VarMap{})
	values["UserAgents"] = MkMap(&VarMap{
		"chrome":   MkString("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.94 Safari/537.36"),
		"chrome39": MkString("Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36"),
	})

	values["headers"] = MkMap(&VarMap{})

	// prepended to URL, like https://proxy.com/https://exchange.com/api...
	values["proxy"] = MkString("")
	values["origin"] = MkString("*") // CORS origin

	values["minFundingAddressLength"] = MkInteger(1)       // used in checkAddress
	values["substituteCommonCurrencyCodes"] = MkBool(true) // reserved
	values["quoteJsonNumbers"] = MkBool(true)              // treat numbers in json as quoted precise strings
	//values["number"] = OpCopy(Number );

	// whether fees should be summed by currency code
	values["reduceFees"] = MkBool(true)

	//values["fetchImplementation"] = OpCopy(defaultFetch );

	values["timeout"] = MkInteger(10000) // milliseconds
	values["verbose"] = MkBool(false)
	values["debug"] = MkBool(false)
	values["userAgent"] = MkUndefined()
	values["twofa"] = MkUndefined() // two-factor authentication (2FA)

	values["apiKey"] = MkUndefined()
	values["secret"] = MkUndefined()
	values["uid"] = MkUndefined()
	values["login"] = MkUndefined()
	values["password"] = MkUndefined()
	values["privateKey"] = MkUndefined()
	values["walletAddress"] = MkUndefined()
	values["token"] = MkUndefined()

	values["balance"] = MkMap(&VarMap{})
	values["orderbooks"] = MkMap(&VarMap{})
	values["tickers"] = MkMap(&VarMap{})
	values["orders"] = MkUndefined()
	values["trades"] = MkMap(&VarMap{})
	values["transactions"] = MkMap(&VarMap{})
	values["ohlcvs"] = MkMap(&VarMap{})
	values["myTrades"] = MkUndefined()
	values["positions"] = MkMap(&VarMap{})
	values["requiresWeb3"] = MkBool(false)

	values["requiresEddsa"] = MkBool(false)
	values["precision"] = MkMap(&VarMap{})

	values["enableLastJsonResponse"] = MkBool(true)
	values["enableLastHttpResponse"] = MkBool(true)
	values["enableLastResponseHeaders"] = MkBool(true)
	values["last_http_response"] = MkUndefined()
	values["last_json_response"] = MkUndefined()
	values["last_response_headers"] = MkUndefined()

	this.itf = itf                                                        // his must be set before the first VCall
	thisValues := this.DeepExtend(MkMap(&values), this.VCall("Describe")) // apply exchange settings
	this.values = this.DeepExtend(thisValues, userConfig)                 // apply user settings

	//this.InitRestRateLimiter(); // todo

	if IsTrue(*this.At(MkString("markets"))) {
		this.SetMarkets(*this.At(MkString("markets")), nil)
	}
}
