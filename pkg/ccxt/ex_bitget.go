package ccxt

import ()

type Bitget struct {
	*ExchangeBase
}

var _ Exchange = (*Bitget)(nil)

func init() {
	exchange := &Bitget{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Bitget) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("bitget"),
		"name": MkString("Bitget"),
		"countries": MkArray(&VarArray{
			MkString("SG"),
		}),
		"version":   MkString("v3"),
		"rateLimit": MkInteger(1000),
		"has": MkMap(&VarMap{
			"cancelOrder":       MkBool(true),
			"cancelOrders":      MkBool(true),
			"CORS":              MkBool(false),
			"createOrder":       MkBool(true),
			"fetchAccounts":     MkBool(true),
			"fetchBalance":      MkBool(true),
			"fetchCurrencies":   MkBool(true),
			"fetchDeposits":     MkBool(true),
			"fetchMarkets":      MkBool(true),
			"fetchMyTrades":     MkBool(true),
			"fetchOHLCV":        MkBool(true),
			"fetchOrder":        MkBool(true),
			"fetchOrderBook":    MkBool(true),
			"fetchOpenOrders":   MkBool(true),
			"fetchClosedOrders": MkBool(true),
			"fetchOrderTrades":  MkBool(true),
			"fetchTicker":       MkBool(true),
			"fetchTickers":      MkBool(true),
			"fetchTime":         MkBool(true),
			"fetchTrades":       MkBool(true),
			"fetchWithdrawals":  MkBool(true),
		}),
		"timeframes": MkMap(&VarMap{
			"1m":  MkString("1m"),
			"5m":  MkString("5m"),
			"15m": MkString("15m"),
			"30m": MkString("30m"),
			"1h":  MkString("1h"),
			"2h":  MkString("2h"),
			"4h":  MkString("4h"),
			"6h":  MkString("6h"),
			"12h": MkString("12h"),
			"1d":  MkString("1d"),
			"1w":  MkString("1w"),
		}),
		"hostname": MkString("bitget.com"),
		"urls": MkMap(&VarMap{
			"logo": MkString("https://user-images.githubusercontent.com/51840849/88317935-a8a21c80-cd22-11ea-8e2b-4b9fac5975eb.jpg"),
			"api": MkMap(&VarMap{
				"data": MkString("https://api.{hostname}"),
				"api":  MkString("https://api.{hostname}"),
				"capi": MkString("https://capi.{hostname}"),
				"swap": MkString("https://capi.{hostname}"),
			}),
			"www": MkString("https://www.bitget.com"),
			"doc": MkArray(&VarArray{
				MkString("https://bitgetlimited.github.io/apidoc/en/swap"),
				MkString("https://bitgetlimited.github.io/apidoc/en/spot"),
			}),
			"fees":     MkString("https://www.bitget.cc/zh-CN/rate?tab=1"),
			"test":     MkMap(&VarMap{"rest": MkString("https://testnet.bitget.com")}),
			"referral": MkString("https://www.bitget.com/expressly?languageType=0&channelCode=ccxt&vipCode=tg9j"),
		}),
		"api": MkMap(&VarMap{
			"data": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("market/history/kline"),
				MkString("market/detail/merged"),
				MkString("market/tickers"),
				MkString("market/allticker"),
				MkString("market/depth"),
				MkString("market/trade"),
				MkString("market/history/trade"),
				MkString("market/detail"),
				MkString("common/symbols"),
				MkString("common/currencys"),
				MkString("common/timestamp"),
			})}),
			"api": MkMap(&VarMap{
				"get": MkArray(&VarArray{
					MkString("account/accounts"),
					MkString("accounts/{account_id}/balance"),
					MkString("order/orders"),
					MkString("order/orders/openOrders"),
					MkString("order/orders/history"),
					MkString("order/deposit_withdraw"),
				}),
				"post": MkArray(&VarArray{
					MkString("order/orders/place"),
					MkString("order/orders/{order_id}/submitcancel"),
					MkString("order/orders/batchcancel"),
					MkString("order/orders/{order_id}"),
					MkString("order/orders/{order_id}/matchresults"),
					MkString("order/matchresults"),
				}),
			}),
			"capi": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("market/time"),
				MkString("market/contracts"),
				MkString("market/depth"),
				MkString("market/tickers"),
				MkString("market/ticker"),
				MkString("market/trades"),
				MkString("market/candles"),
				MkString("market/index"),
				MkString("market/open_count"),
				MkString("market/open_interest"),
				MkString("market/price_limit"),
				MkString("market/funding_time"),
				MkString("market/mark_price"),
				MkString("market/open_count"),
				MkString("market/historyFundRate"),
			})}),
			"swap": MkMap(&VarMap{
				"get": MkArray(&VarArray{
					MkString("account/accounts"),
					MkString("account/account"),
					MkString("account/settings"),
					MkString("position/allPosition"),
					MkString("position/singlePosition"),
					MkString("position/holds"),
					MkString("order/detail"),
					MkString("order/orders"),
					MkString("order/fills"),
					MkString("order/current"),
					MkString("order/currentPlan"),
					MkString("order/history"),
					MkString("order/historyPlan"),
					MkString("trace/closeTrack"),
					MkString("trace/currentTrack"),
					MkString("trace/historyTrack"),
					MkString("trace/summary"),
					MkString("trace/profitSettleTokenIdGroup"),
					MkString("trace/profitDateGroupList"),
					MkString("trace/profitDateList"),
					MkString("trace/waitProfitDateList"),
				}),
				"post": MkArray(&VarArray{
					MkString("account/leverage"),
					MkString("account/adjustMargin"),
					MkString("account/modifyAutoAppendMargin"),
					MkString("order/placeOrder"),
					MkString("order/batchOrders"),
					MkString("order/cancel_order"),
					MkString("order/cancel_batch_orders"),
					MkString("order/plan_order"),
					MkString("order/cancel_plan"),
					MkString("position/changeHoldModel"),
					MkString("trace/closeTrackOrder"),
				}),
			}),
		}),
		"fees": MkMap(&VarMap{
			"spot": MkMap(&VarMap{
				"taker": this.ParseNumber(MkString("0.002")),
				"maker": this.ParseNumber(MkString("0.002")),
			}),
			"swap": MkMap(&VarMap{
				"taker": this.ParseNumber(MkString("0.0006")),
				"maker": this.ParseNumber(MkString("0.0004")),
			}),
		}),
		"requiredCredentials": MkMap(&VarMap{
			"apiKey":   MkBool(true),
			"secret":   MkBool(true),
			"password": MkBool(true),
		}),
		"exceptions": MkMap(&VarMap{
			"exact": MkMap(&VarMap{
				"1": ExchangeError,
				"failure to get a peer from the ring-balancer": ExchangeNotAvailable,
				"4010":               PermissionDenied,
				"4001":               ExchangeError,
				"4002":               ExchangeError,
				"30001":              AuthenticationError,
				"30002":              AuthenticationError,
				"30003":              AuthenticationError,
				"30004":              AuthenticationError,
				"30005":              InvalidNonce,
				"30006":              AuthenticationError,
				"30007":              BadRequest,
				"30008":              RequestTimeout,
				"30009":              ExchangeError,
				"30010":              AuthenticationError,
				"30011":              PermissionDenied,
				"30012":              AuthenticationError,
				"30013":              AuthenticationError,
				"30014":              DDoSProtection,
				"30015":              AuthenticationError,
				"30016":              ExchangeError,
				"30017":              ExchangeError,
				"30018":              ExchangeError,
				"30019":              ExchangeNotAvailable,
				"30020":              BadRequest,
				"30021":              BadRequest,
				"30022":              PermissionDenied,
				"30023":              BadRequest,
				"30024":              BadSymbol,
				"30025":              BadRequest,
				"30026":              DDoSProtection,
				"30027":              AuthenticationError,
				"30028":              PermissionDenied,
				"30029":              AccountSuspended,
				"30030":              ExchangeError,
				"30031":              BadRequest,
				"30032":              BadSymbol,
				"30033":              BadRequest,
				"30034":              ExchangeError,
				"30035":              ExchangeError,
				"30036":              ExchangeError,
				"30037":              ExchangeNotAvailable,
				"30038":              OnMaintenance,
				"32001":              AccountSuspended,
				"32002":              PermissionDenied,
				"32003":              CancelPending,
				"32004":              ExchangeError,
				"32005":              InvalidOrder,
				"32006":              InvalidOrder,
				"32007":              InvalidOrder,
				"32008":              InvalidOrder,
				"32009":              InvalidOrder,
				"32010":              ExchangeError,
				"32011":              ExchangeError,
				"32012":              ExchangeError,
				"32013":              ExchangeError,
				"32014":              ExchangeError,
				"32015":              ExchangeError,
				"32016":              ExchangeError,
				"32017":              ExchangeError,
				"32018":              ExchangeError,
				"32019":              ExchangeError,
				"32020":              ExchangeError,
				"32021":              ExchangeError,
				"32022":              ExchangeError,
				"32023":              ExchangeError,
				"32024":              ExchangeError,
				"32025":              ExchangeError,
				"32026":              ExchangeError,
				"32027":              ExchangeError,
				"32028":              AccountSuspended,
				"32029":              ExchangeError,
				"32030":              InvalidOrder,
				"32031":              ArgumentsRequired,
				"32038":              AuthenticationError,
				"32040":              ExchangeError,
				"32044":              ExchangeError,
				"32045":              ExchangeError,
				"32046":              ExchangeError,
				"32047":              ExchangeError,
				"32048":              InvalidOrder,
				"32049":              ExchangeError,
				"32050":              InvalidOrder,
				"32051":              InvalidOrder,
				"32052":              ExchangeError,
				"32053":              ExchangeError,
				"32057":              ExchangeError,
				"32054":              ExchangeError,
				"32055":              InvalidOrder,
				"32056":              ExchangeError,
				"32058":              ExchangeError,
				"32059":              InvalidOrder,
				"32060":              InvalidOrder,
				"32061":              InvalidOrder,
				"32062":              InvalidOrder,
				"32063":              InvalidOrder,
				"32064":              ExchangeError,
				"32065":              ExchangeError,
				"32066":              ExchangeError,
				"32067":              ExchangeError,
				"32068":              ExchangeError,
				"32069":              ExchangeError,
				"32070":              ExchangeError,
				"32071":              ExchangeError,
				"32072":              ExchangeError,
				"32073":              ExchangeError,
				"32074":              ExchangeError,
				"32075":              ExchangeError,
				"32076":              ExchangeError,
				"32077":              ExchangeError,
				"32078":              ExchangeError,
				"32079":              ExchangeError,
				"32080":              ExchangeError,
				"32083":              ExchangeError,
				"33001":              PermissionDenied,
				"33002":              AccountSuspended,
				"33003":              InsufficientFunds,
				"33004":              ExchangeError,
				"33005":              ExchangeError,
				"33006":              ExchangeError,
				"33007":              ExchangeError,
				"33008":              InsufficientFunds,
				"33009":              ExchangeError,
				"33010":              ExchangeError,
				"33011":              ExchangeError,
				"33012":              ExchangeError,
				"33013":              InvalidOrder,
				"33014":              OrderNotFound,
				"33015":              InvalidOrder,
				"33016":              ExchangeError,
				"33017":              InsufficientFunds,
				"33018":              ExchangeError,
				"33020":              ExchangeError,
				"33021":              BadRequest,
				"33022":              InvalidOrder,
				"33023":              ExchangeError,
				"33024":              InvalidOrder,
				"33025":              InvalidOrder,
				"33026":              ExchangeError,
				"33027":              InvalidOrder,
				"33028":              InvalidOrder,
				"33029":              InvalidOrder,
				"33034":              ExchangeError,
				"33035":              ExchangeError,
				"33036":              ExchangeError,
				"33037":              ExchangeError,
				"33038":              ExchangeError,
				"33039":              ExchangeError,
				"33040":              ExchangeError,
				"33041":              ExchangeError,
				"33042":              ExchangeError,
				"33043":              ExchangeError,
				"33044":              ExchangeError,
				"33045":              ExchangeError,
				"33046":              ExchangeError,
				"33047":              ExchangeError,
				"33048":              ExchangeError,
				"33049":              ExchangeError,
				"33050":              ExchangeError,
				"33051":              ExchangeError,
				"33059":              BadRequest,
				"33060":              BadRequest,
				"33061":              ExchangeError,
				"33062":              ExchangeError,
				"33063":              ExchangeError,
				"33064":              ExchangeError,
				"33065":              ExchangeError,
				"21009":              ExchangeError,
				"34001":              PermissionDenied,
				"34002":              InvalidAddress,
				"34003":              ExchangeError,
				"34004":              ExchangeError,
				"34005":              ExchangeError,
				"34006":              ExchangeError,
				"34007":              ExchangeError,
				"34008":              InsufficientFunds,
				"34009":              ExchangeError,
				"34010":              ExchangeError,
				"34011":              ExchangeError,
				"34012":              ExchangeError,
				"34013":              ExchangeError,
				"34014":              ExchangeError,
				"34015":              ExchangeError,
				"34016":              PermissionDenied,
				"34017":              AccountSuspended,
				"34018":              AuthenticationError,
				"34019":              PermissionDenied,
				"34020":              PermissionDenied,
				"34021":              InvalidAddress,
				"34022":              ExchangeError,
				"34023":              PermissionDenied,
				"34026":              ExchangeError,
				"34036":              ExchangeError,
				"34037":              ExchangeError,
				"34038":              ExchangeError,
				"34039":              ExchangeError,
				"35001":              ExchangeError,
				"35002":              ExchangeError,
				"35003":              ExchangeError,
				"35004":              ExchangeError,
				"35005":              AuthenticationError,
				"35008":              InvalidOrder,
				"35010":              InvalidOrder,
				"35012":              InvalidOrder,
				"35014":              InvalidOrder,
				"35015":              InvalidOrder,
				"35017":              ExchangeError,
				"35019":              InvalidOrder,
				"35020":              InvalidOrder,
				"35021":              InvalidOrder,
				"35022":              ExchangeError,
				"35024":              ExchangeError,
				"35025":              InsufficientFunds,
				"35026":              ExchangeError,
				"35029":              OrderNotFound,
				"35030":              InvalidOrder,
				"35031":              InvalidOrder,
				"35032":              ExchangeError,
				"35037":              ExchangeError,
				"35039":              ExchangeError,
				"35040":              InvalidOrder,
				"35044":              ExchangeError,
				"35046":              InsufficientFunds,
				"35047":              InsufficientFunds,
				"35048":              ExchangeError,
				"35049":              InvalidOrder,
				"35050":              InvalidOrder,
				"35052":              InsufficientFunds,
				"35053":              ExchangeError,
				"35055":              InsufficientFunds,
				"35057":              ExchangeError,
				"35058":              ExchangeError,
				"35059":              BadRequest,
				"35060":              BadRequest,
				"35061":              BadRequest,
				"35062":              InvalidOrder,
				"35063":              InvalidOrder,
				"35064":              InvalidOrder,
				"35066":              InvalidOrder,
				"35067":              InvalidOrder,
				"35068":              InvalidOrder,
				"35069":              InvalidOrder,
				"35070":              InvalidOrder,
				"35071":              InvalidOrder,
				"35072":              InvalidOrder,
				"35073":              InvalidOrder,
				"35074":              InvalidOrder,
				"35075":              InvalidOrder,
				"35076":              InvalidOrder,
				"35077":              InvalidOrder,
				"35078":              InvalidOrder,
				"35079":              InvalidOrder,
				"35080":              InvalidOrder,
				"35081":              InvalidOrder,
				"35082":              InvalidOrder,
				"35083":              InvalidOrder,
				"35084":              InvalidOrder,
				"35085":              InvalidOrder,
				"35086":              InvalidOrder,
				"35087":              InvalidOrder,
				"35088":              InvalidOrder,
				"35089":              InvalidOrder,
				"35090":              ExchangeError,
				"35091":              ExchangeError,
				"35092":              ExchangeError,
				"35093":              ExchangeError,
				"35094":              ExchangeError,
				"35095":              BadRequest,
				"35096":              ExchangeError,
				"35097":              ExchangeError,
				"35098":              ExchangeError,
				"35099":              ExchangeError,
				"36001":              BadRequest,
				"36002":              BadRequest,
				"36005":              ExchangeError,
				"36101":              AuthenticationError,
				"36102":              PermissionDenied,
				"36103":              AccountSuspended,
				"36104":              PermissionDenied,
				"36105":              PermissionDenied,
				"36106":              AccountSuspended,
				"36107":              PermissionDenied,
				"36108":              InsufficientFunds,
				"36109":              PermissionDenied,
				"36201":              PermissionDenied,
				"36202":              PermissionDenied,
				"36203":              InvalidOrder,
				"36204":              ExchangeError,
				"36205":              BadRequest,
				"36206":              BadRequest,
				"36207":              InvalidOrder,
				"36208":              InvalidOrder,
				"36209":              InvalidOrder,
				"36210":              InvalidOrder,
				"36211":              InvalidOrder,
				"36212":              InvalidOrder,
				"36213":              InvalidOrder,
				"36214":              ExchangeError,
				"36216":              OrderNotFound,
				"36217":              InvalidOrder,
				"36218":              InvalidOrder,
				"36219":              InvalidOrder,
				"36220":              InvalidOrder,
				"36221":              InvalidOrder,
				"36222":              InvalidOrder,
				"36223":              InvalidOrder,
				"36224":              InvalidOrder,
				"36225":              InvalidOrder,
				"36226":              InvalidOrder,
				"36227":              InvalidOrder,
				"36228":              InvalidOrder,
				"36229":              InvalidOrder,
				"36230":              InvalidOrder,
				"400":                BadRequest,
				"401":                AuthenticationError,
				"403":                PermissionDenied,
				"404":                BadRequest,
				"405":                BadRequest,
				"415":                BadRequest,
				"429":                DDoSProtection,
				"500":                ExchangeNotAvailable,
				"1001":               RateLimitExceeded,
				"1002":               ExchangeError,
				"1003":               ExchangeError,
				"40001":              AuthenticationError,
				"40002":              AuthenticationError,
				"40003":              AuthenticationError,
				"40004":              InvalidNonce,
				"40005":              InvalidNonce,
				"40006":              AuthenticationError,
				"40007":              BadRequest,
				"40008":              InvalidNonce,
				"40009":              AuthenticationError,
				"40010":              AuthenticationError,
				"40011":              AuthenticationError,
				"40012":              AuthenticationError,
				"40013":              ExchangeError,
				"40014":              PermissionDenied,
				"40015":              ExchangeError,
				"40016":              PermissionDenied,
				"40017":              ExchangeError,
				"40018":              PermissionDenied,
				"40102":              BadRequest,
				"40103":              BadRequest,
				"40104":              ExchangeError,
				"40105":              ExchangeError,
				"40106":              ExchangeError,
				"40107":              ExchangeError,
				"40108":              InvalidOrder,
				"40109":              OrderNotFound,
				"40200":              OnMaintenance,
				"40201":              InvalidOrder,
				"40202":              ExchangeError,
				"40203":              BadRequest,
				"40204":              BadRequest,
				"40205":              BadRequest,
				"40206":              BadRequest,
				"40207":              BadRequest,
				"40208":              BadRequest,
				"40209":              BadRequest,
				"40300":              ExchangeError,
				"40301":              PermissionDenied,
				"40302":              BadRequest,
				"40303":              BadRequest,
				"40304":              BadRequest,
				"40305":              BadRequest,
				"40306":              ExchangeError,
				"40308":              OnMaintenance,
				"40309":              BadSymbol,
				"40400":              ExchangeError,
				"40401":              ExchangeError,
				"40402":              BadRequest,
				"40403":              BadRequest,
				"40404":              BadRequest,
				"40405":              BadRequest,
				"40406":              BadRequest,
				"40407":              ExchangeError,
				"40408":              ExchangeError,
				"40409":              ExchangeError,
				"40500":              InvalidOrder,
				"40501":              ExchangeError,
				"40502":              ExchangeError,
				"40503":              ExchangeError,
				"40504":              ExchangeError,
				"40505":              ExchangeError,
				"40506":              AuthenticationError,
				"40507":              AuthenticationError,
				"40508":              ExchangeError,
				"40509":              ExchangeError,
				"40600":              ExchangeError,
				"40601":              ExchangeError,
				"40602":              ExchangeError,
				"40603":              ExchangeError,
				"40604":              ExchangeNotAvailable,
				"40605":              ExchangeError,
				"40606":              ExchangeError,
				"40607":              ExchangeError,
				"40608":              ExchangeError,
				"40609":              ExchangeError,
				"40700":              BadRequest,
				"40701":              ExchangeError,
				"40702":              ExchangeError,
				"40703":              ExchangeError,
				"40704":              ExchangeError,
				"40705":              BadRequest,
				"40706":              InvalidOrder,
				"40707":              BadRequest,
				"40708":              BadRequest,
				"40709":              ExchangeError,
				"40710":              ExchangeError,
				"40711":              InsufficientFunds,
				"40712":              InsufficientFunds,
				"40713":              ExchangeError,
				"40714":              ExchangeError,
				"invalid sign":       AuthenticationError,
				"invalid currency":   BadSymbol,
				"invalid symbol":     BadSymbol,
				"invalid period":     BadRequest,
				"invalid user":       ExchangeError,
				"invalid amount":     InvalidOrder,
				"invalid type":       InvalidOrder,
				"invalid orderId":    InvalidOrder,
				"invalid record":     ExchangeError,
				"invalid accountId":  BadRequest,
				"invalid address":    BadRequest,
				"accesskey not null": AuthenticationError,
				"illegal accesskey":  AuthenticationError,
				"sign not null":      AuthenticationError,
				"req_time is too much difference from server time": InvalidNonce,
				"permissions not right":                            PermissionDenied,
				"illegal sign invalid":                             AuthenticationError,
				"user locked":                                      AccountSuspended,
				"Request Frequency Is Too High":                    RateLimitExceeded,
				"more than a daily rate of cash":                   BadRequest,
				"more than the maximum daily withdrawal amount":    BadRequest,
				"need to bind email or mobile":                     ExchangeError,
				"user forbid":                                      PermissionDenied,
				"User Prohibited Cash Withdrawal":                  PermissionDenied,
				"Cash Withdrawal Is Less Than The Minimum Value":   BadRequest,
				"Cash Withdrawal Is More Than The Maximum Value":   BadRequest,
				"the account with in 24 hours ban coin":            PermissionDenied,
				"order cancel fail":                                BadRequest,
				"base symbol error":                                BadSymbol,
				"base date error":                                  ExchangeError,
				"api signature not valid":                          AuthenticationError,
				"gateway internal error":                           ExchangeError,
				"audit failed":                                     ExchangeError,
				"order queryorder invalid":                         BadRequest,
				"market no need price":                             InvalidOrder,
				"limit need price":                                 InvalidOrder,
				"userid not equal to account_id":                   ExchangeError,
				"your balance is low":                              InsufficientFunds,
				"address invalid cointype":                         ExchangeError,
				"system exception":                                 ExchangeError,
				"50003":                                            ExchangeError,
				"50004":                                            BadSymbol,
				"50006":                                            PermissionDenied,
				"50007":                                            PermissionDenied,
				"50008":                                            RequestTimeout,
				"50009":                                            RateLimitExceeded,
				"50010":                                            ExchangeError,
				"50014":                                            InvalidOrder,
				"50015":                                            InvalidOrder,
				"50016":                                            InvalidOrder,
				"50017":                                            InvalidOrder,
				"50018":                                            InvalidOrder,
				"50019":                                            InvalidOrder,
				"50020":                                            InsufficientFunds,
				"50021":                                            InvalidOrder,
				"50026":                                            InvalidOrder,
				"invalid order query time":                         ExchangeError,
				"invalid start time":                               BadRequest,
				"invalid end time":                                 BadRequest,
				"20003":                                            ExchangeError,
				"01001":                                            ExchangeError,
			}),
			"broad": MkMap(&VarMap{"invalid size, valid range": ExchangeError}),
		}),
		"precisionMode": TICK_SIZE,
		"options": MkMap(&VarMap{
			"createMarketBuyOrderRequiresPrice": MkBool(true),
			"fetchMarkets": MkArray(&VarArray{
				MkString("spot"),
				MkString("swap"),
			}),
			"parseOHLCV": MkMap(&VarMap{"volume": MkMap(&VarMap{
				"spot": MkString("amount"),
				"swap": MkInteger(5),
			})}),
			"defaultType": MkString("spot"),
			"accountId":   MkUndefined(),
			"timeframes": MkMap(&VarMap{
				"spot": MkMap(&VarMap{
					"1m":  MkString("1min"),
					"5m":  MkString("5min"),
					"15m": MkString("15min"),
					"30m": MkString("30min"),
					"1h":  MkString("60min"),
					"2h":  MkString("120min"),
					"4h":  MkString("240min"),
					"6h":  MkString("360min"),
					"12h": MkString("720min"),
					"1d":  MkString("1day"),
					"1w":  MkString("1week"),
				}),
				"swap": MkMap(&VarMap{
					"1m":  MkString("60"),
					"5m":  MkString("300"),
					"15m": MkString("900"),
					"30m": MkString("1800"),
					"1h":  MkString("3600"),
					"2h":  MkString("7200"),
					"4h":  MkString("14400"),
					"6h":  MkString("21600"),
					"12h": MkString("43200"),
					"1d":  MkString("86400"),
					"1w":  MkString("604800"),
				}),
			}),
		}),
	}))
}

func (this *Bitget) FetchTime(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("dataGetCommonTimestamp"), params)
	_ = response
	return this.SafeInteger(response, MkString("data"))
}

func (this *Bitget) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	types := this.SafeValue(*this.At(MkString("options")), MkString("fetchMarkets"))
	_ = types
	if IsTrue(OpNot(types.Length)) {
		types = MkArray(&VarArray{
			*(*this.At(MkString("options"))).At(MkString("defaultType")),
		})
	}
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, types.Length)); OpIncr(&i) {
		markets := this.FetchMarketsByType(*(types).At(i), params)
		_ = markets
		result = this.ArrayConcat(result, markets)
	}
	return result
}

func (this *Bitget) ParseMarkets(goArgs ...*Variant) *Variant {
	markets := GoGetArg(goArgs, 0, MkUndefined())
	_ = markets
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, markets.Length)); OpIncr(&i) {
		result.Push(this.ParseMarket(*(markets).At(i)))
	}
	return result
}

func (this *Bitget) ParseMarket(goArgs ...*Variant) *Variant {
	market := GoGetArg(goArgs, 0, MkUndefined())
	_ = market
	id := this.SafeString(market, MkString("symbol"))
	_ = id
	marketType := MkString("spot")
	_ = marketType
	spot := OpCopy(MkBool(true))
	_ = spot
	swap := OpCopy(MkBool(false))
	_ = swap
	baseId := this.SafeString2(market, MkString("base_currency"), MkString("coin"))
	_ = baseId
	quoteId := this.SafeString(market, MkString("quote_currency"))
	_ = quoteId
	contractVal := this.SafeNumber(market, MkString("contract_val"))
	_ = contractVal
	if IsTrue(OpNotEq2(contractVal, MkUndefined())) {
		marketType = MkString("swap")
		spot = OpCopy(MkBool(false))
		swap = OpCopy(MkBool(true))
	}
	base := this.SafeCurrencyCode(baseId)
	_ = base
	quote := this.SafeCurrencyCode(quoteId)
	_ = quote
	symbol := id.ToUpperCase()
	_ = symbol
	if IsTrue(spot) {
		symbol = OpAdd(base, OpAdd(MkString("/"), quote))
	}
	tickSize := this.SafeString(market, MkString("tick_size"))
	_ = tickSize
	sizeIncrement := this.SafeString(market, MkString("size_increment"))
	_ = sizeIncrement
	precision := MkMap(&VarMap{
		"amount": this.ParseNumber(this.ParsePrecision(sizeIncrement)),
		"price":  this.ParseNumber(this.ParsePrecision(tickSize)),
	})
	_ = precision
	minAmount := this.SafeNumber2(market, MkString("min_size"), MkString("base_min_size"))
	_ = minAmount
	status := this.SafeString(market, MkString("status"))
	_ = status
	active := OpCopy(MkUndefined())
	_ = active
	if IsTrue(OpNotEq2(status, MkUndefined())) {
		active = OpEq2(status, MkString("1"))
	}
	fees := this.SafeValue2(*this.At(MkString("fees")), marketType, MkString("trading"), MkMap(&VarMap{}))
	_ = fees
	return this.Extend(fees, MkMap(&VarMap{
		"id":        id,
		"symbol":    symbol,
		"base":      base,
		"quote":     quote,
		"baseId":    baseId,
		"quoteId":   quoteId,
		"info":      market,
		"type":      marketType,
		"spot":      spot,
		"swap":      swap,
		"active":    active,
		"precision": precision,
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
				"min": *(precision).At(MkString("price")),
				"max": MkUndefined(),
			}),
		}),
	}))
}

func (this *Bitget) FetchMarketsByType(goArgs ...*Variant) *Variant {
	type_ := GoGetArg(goArgs, 0, MkUndefined())
	_ = type_
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(type_, MkString("spot"))) {
		response := this.Call(MkString("dataGetCommonSymbols"), params)
		_ = response
		data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
		_ = data
		return this.ParseMarkets(data)
	} else {
		if IsTrue(OpEq2(type_, MkString("swap"))) {
			response := this.Call(MkString("capiGetMarketContracts"), params)
			_ = response
			return this.ParseMarkets(response)
		} else {
			panic(NewNotSupported(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" fetchMarketsByType does not support market type "), type_))))
		}
	}
	return MkUndefined()
}

func (this *Bitget) FetchCurrencies(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("dataGetCommonCurrencys"), params)
	_ = response
	result := MkMap(&VarMap{})
	_ = result
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	for i := MkInteger(0); IsTrue(OpLw(i, data.Length)); OpIncr(&i) {
		id := *(data).At(i)
		_ = id
		code := this.SafeCurrencyCode(id)
		_ = code
		*(result).At(code) = MkMap(&VarMap{
			"id":        id,
			"code":      code,
			"info":      id,
			"type":      MkUndefined(),
			"name":      MkUndefined(),
			"active":    MkUndefined(),
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

func (this *Bitget) FetchOrderBook(goArgs ...*Variant) *Variant {
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
	method := OpCopy(MkUndefined())
	_ = method
	if IsTrue(*(market).At(MkString("spot"))) {
		method = MkString("dataGetMarketDepth")
		*(request).At(MkString("type")) = MkString("step0")
	} else {
		if IsTrue(*(market).At(MkString("swap"))) {
			method = MkString("capiGetMarketDepth")
			*(request).At(MkString("limit")) = OpTrinary(OpEq2(limit, MkUndefined()), MkInteger(100), limit)
		}
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), response)
	_ = data
	timestamp := this.SafeInteger2(data, MkString("timestamp"), MkString("ts"))
	_ = timestamp
	nonce := this.SafeInteger(data, MkString("id"))
	_ = nonce
	orderbook := this.ParseOrderBook(data, symbol, timestamp)
	_ = orderbook
	*(orderbook).At(MkString("nonce")) = OpCopy(nonce)
	return orderbook
}

func (this *Bitget) ParseTicker(goArgs ...*Variant) *Variant {
	ticker := GoGetArg(goArgs, 0, MkUndefined())
	_ = ticker
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.SafeInteger2(ticker, MkString("timestamp"), MkString("id"))
	_ = timestamp
	symbol := OpCopy(MkUndefined())
	_ = symbol
	marketId := this.SafeString2(ticker, MkString("instrument_id"), MkString("symbol"))
	_ = marketId
	if IsTrue(OpHasMember(marketId, *this.At(MkString("markets_by_id")))) {
		market = *(*this.At(MkString("markets_by_id"))).At(marketId)
		symbol = *(market).At(MkString("symbol"))
	} else {
		if IsTrue(OpNotEq2(marketId, MkUndefined())) {
			parts := marketId.Split(MkString("_"))
			_ = parts
			numParts := OpCopy(parts.Length)
			_ = numParts
			if IsTrue(OpEq2(numParts, MkInteger(2))) {
				baseId := MkUndefined()
				_ = baseId
				quoteId := MkUndefined()
				_ = quoteId
				ArrayUnpack(parts, &baseId, &quoteId)
				base := this.SafeCurrencyCode(baseId)
				_ = base
				quote := this.SafeCurrencyCode(quoteId)
				_ = quote
				symbol = OpAdd(base, OpAdd(MkString("/"), quote))
			} else {
				symbol = OpCopy(marketId)
			}
		}
	}
	if IsTrue(OpAnd(OpEq2(symbol, MkUndefined()), OpNotEq2(market, MkUndefined()))) {
		symbol = *(market).At(MkString("symbol"))
	}
	last := this.SafeNumber2(ticker, MkString("last"), MkString("close"))
	_ = last
	open := this.SafeNumber(ticker, MkString("open"))
	_ = open
	bidVolume := OpCopy(MkUndefined())
	_ = bidVolume
	askVolume := OpCopy(MkUndefined())
	_ = askVolume
	bid := this.SafeValue(ticker, MkString("bid"))
	_ = bid
	if IsTrue(OpEq2(bid, MkUndefined())) {
		bid = this.SafeNumber(ticker, MkString("best_bid"))
	} else {
		bidVolume = this.SafeNumber(bid, MkInteger(1))
		bid = this.SafeNumber(bid, MkInteger(0))
	}
	ask := this.SafeValue(ticker, MkString("ask"))
	_ = ask
	if IsTrue(OpEq2(ask, MkUndefined())) {
		ask = this.SafeNumber(ticker, MkString("best_ask"))
	} else {
		askVolume = this.SafeNumber(ask, MkInteger(1))
		ask = this.SafeNumber(ask, MkInteger(0))
	}
	baseVolume := this.SafeNumber2(ticker, MkString("amount"), MkString("volume_24h"))
	_ = baseVolume
	quoteVolume := this.SafeNumber(ticker, MkString("vol"))
	_ = quoteVolume
	vwap := this.Vwap(baseVolume, quoteVolume)
	_ = vwap
	change := OpCopy(MkUndefined())
	_ = change
	percentage := OpCopy(MkUndefined())
	_ = percentage
	average := OpCopy(MkUndefined())
	_ = average
	if IsTrue(OpAnd(OpNotEq2(last, MkUndefined()), OpNotEq2(open, MkUndefined()))) {
		change = OpSub(last, open)
		percentage = OpDiv(change, OpMulti(open, MkInteger(100)))
		average = OpDiv(this.Sum(open, last), MkInteger(2))
	}
	return MkMap(&VarMap{
		"symbol":        symbol,
		"timestamp":     timestamp,
		"datetime":      this.Iso8601(timestamp),
		"high":          this.SafeNumber2(ticker, MkString("high"), MkString("high_24h")),
		"low":           this.SafeNumber2(ticker, MkString("low"), MkString("low_24h")),
		"bid":           bid,
		"bidVolume":     bidVolume,
		"ask":           ask,
		"askVolume":     askVolume,
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

func (this *Bitget) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	method := OpCopy(MkUndefined())
	_ = method
	if IsTrue(*(market).At(MkString("spot"))) {
		method = MkString("dataGetMarketDetailMerged")
	} else {
		if IsTrue(*(market).At(MkString("swap"))) {
			method = MkString("capiGetMarketTicker")
		}
	}
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(method, this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), response)
	_ = data
	return this.ParseTicker(data, market)
}

func (this *Bitget) FetchTickersByType(goArgs ...*Variant) *Variant {
	type_ := GoGetArg(goArgs, 0, MkUndefined())
	_ = type_
	symbols := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	method := OpCopy(MkUndefined())
	_ = method
	if IsTrue(OpEq2(type_, MkString("spot"))) {
		method = MkString("dataGetMarketTickers")
	} else {
		if IsTrue(OpEq2(type_, MkString("swap"))) {
			method = MkString("capiGetMarketTickers")
		}
	}
	response := this.Call(method, params)
	_ = response
	data := this.SafeValue(response, MkString("data"), response)
	_ = data
	timestamp := OpCopy(MkUndefined())
	_ = timestamp
	if IsTrue(OpNot(GoIsArray(response))) {
		timestamp = this.SafeInteger(response, MkString("ts"))
	}
	result := MkMap(&VarMap{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, data.Length)); OpIncr(&i) {
		ticker := this.ParseTicker(this.Extend(MkMap(&VarMap{"timestamp": timestamp}), *(data).At(i)))
		_ = ticker
		symbol := *(ticker).At(MkString("symbol"))
		_ = symbol
		*(result).At(symbol) = OpCopy(ticker)
	}
	return this.FilterByArray(result, MkString("symbol"), symbols)
}

func (this *Bitget) FetchTickers(goArgs ...*Variant) *Variant {
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

func (this *Bitget) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	symbol := OpCopy(MkUndefined())
	_ = symbol
	marketId := this.SafeString(trade, MkString("symbol"))
	_ = marketId
	base := OpCopy(MkUndefined())
	_ = base
	quote := OpCopy(MkUndefined())
	_ = quote
	if IsTrue(OpHasMember(marketId, *this.At(MkString("markets_by_id")))) {
		market = *(*this.At(MkString("markets_by_id"))).At(marketId)
		symbol = *(market).At(MkString("symbol"))
		base = *(market).At(MkString("base"))
		quote = *(market).At(MkString("quote"))
	} else {
		if IsTrue(OpNotEq2(marketId, MkUndefined())) {
			parts := marketId.Split(MkString("_"))
			_ = parts
			numParts := OpCopy(parts.Length)
			_ = numParts
			if IsTrue(OpEq2(numParts, MkInteger(2))) {
				baseId := MkUndefined()
				_ = baseId
				quoteId := MkUndefined()
				_ = quoteId
				ArrayUnpack(parts, &baseId, &quoteId)
				base = this.SafeCurrencyCode(baseId)
				quote = this.SafeCurrencyCode(quoteId)
				symbol = OpAdd(base, OpAdd(MkString("/"), quote))
			} else {
				symbol = marketId.ToUpperCase()
			}
		}
	}
	if IsTrue(OpAnd(OpEq2(symbol, MkUndefined()), OpNotEq2(market, MkUndefined()))) {
		symbol = *(market).At(MkString("symbol"))
		base = *(market).At(MkString("base"))
		quote = *(market).At(MkString("quote"))
	}
	timestamp := this.SafeInteger(trade, MkString("created_at"))
	_ = timestamp
	timestamp = this.SafeInteger2(trade, MkString("timestamp"), MkString("ts"), timestamp)
	priceString := this.SafeString(trade, MkString("price"))
	_ = priceString
	amountString := this.SafeString2(trade, MkString("filled_amount"), MkString("order_qty"))
	_ = amountString
	amountString = this.SafeString2(trade, MkString("size"), MkString("amount"), amountString)
	price := this.ParseNumber(priceString)
	_ = price
	amount := this.ParseNumber(amountString)
	_ = amount
	cost := this.ParseNumber(Precise.StringMul(priceString, amountString))
	_ = cost
	takerOrMaker := this.SafeString2(trade, MkString("exec_type"), MkString("liquidity"))
	_ = takerOrMaker
	if IsTrue(OpEq2(takerOrMaker, MkString("M"))) {
		takerOrMaker = MkString("maker")
	} else {
		if IsTrue(OpEq2(takerOrMaker, MkString("T"))) {
			takerOrMaker = MkString("taker")
		}
	}
	orderType := this.SafeString(trade, MkString("type"))
	_ = orderType
	side := OpCopy(MkUndefined())
	_ = side
	type_ := OpCopy(MkUndefined())
	_ = type_
	if IsTrue(OpNotEq2(orderType, MkUndefined())) {
		side = this.SafeString(trade, MkString("type"))
		type_ = this.ParseOrderType(side)
		side = this.ParseOrderSide(side)
	} else {
		side = this.SafeString2(trade, MkString("side"), MkString("direction"))
		type_ = this.ParseOrderType(side)
		side = this.ParseOrderSide(side)
	}
	feeCostString := this.SafeString(trade, MkString("fee"))
	_ = feeCostString
	if IsTrue(OpEq2(feeCostString, MkUndefined())) {
		feeCostString = this.SafeString(trade, MkString("filled_fees"))
	} else {
		feeCostString = Precise.StringNeg(feeCostString)
	}
	feeCost := this.ParseNumber(feeCostString)
	_ = feeCost
	fee := OpCopy(MkUndefined())
	_ = fee
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		feeCurrency := OpTrinary(OpEq2(side, MkString("buy")), base, quote)
		_ = feeCurrency
		fee = MkMap(&VarMap{
			"cost":     feeCost,
			"currency": feeCurrency,
		})
	}
	orderId := this.SafeString(trade, MkString("order_id"))
	_ = orderId
	id := this.SafeString2(trade, MkString("trade_id"), MkString("id"))
	_ = id
	return MkMap(&VarMap{
		"info":         trade,
		"timestamp":    timestamp,
		"datetime":     this.Iso8601(timestamp),
		"symbol":       symbol,
		"id":           id,
		"order":        orderId,
		"type":         type_,
		"takerOrMaker": takerOrMaker,
		"side":         side,
		"price":        price,
		"amount":       amount,
		"cost":         cost,
		"fee":          fee,
	})
}

func (this *Bitget) FetchTrades(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	since := GoGetArg(goArgs, 2, MkUndefined())
	_ = since
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
		method = MkString("dataGetMarketHistoryTrade")
	} else {
		if IsTrue(*(market).At(MkString("swap"))) {
			method = MkString("capiGetMarketTrades")
		}
	}
	if IsTrue(*(market).At(MkString("spot"))) {
		if IsTrue(OpNotEq2(limit, MkUndefined())) {
			*(request).At(MkString("size")) = OpCopy(limit)
		}
	} else {
		if IsTrue(*(market).At(MkString("swap"))) {
			if IsTrue(OpEq2(limit, MkUndefined())) {
				limit = MkInteger(100)
			}
			*(request).At(MkString("limit")) = OpCopy(limit)
		}
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	trades := OpCopy(MkUndefined())
	_ = trades
	if IsTrue(GoIsArray(response)) {
		trades = OpCopy(response)
	} else {
		data := this.SafeValue(response, MkString("data"), MkMap(&VarMap{}))
		_ = data
		trades = this.SafeValue2(data, MkString("data"), MkArray(&VarArray{}))
	}
	return this.ParseTrades(trades, market, since, limit)
}

func (this *Bitget) ParseOHLCV(goArgs ...*Variant) *Variant {
	ohlcv := GoGetArg(goArgs, 0, MkUndefined())
	_ = ohlcv
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timeframe := GoGetArg(goArgs, 2, MkString("1m"))
	_ = timeframe
	options := this.SafeValue(*this.At(MkString("options")), MkString("parseOHLCV"), MkMap(&VarMap{}))
	_ = options
	volume := this.SafeValue(options, MkString("volume"), MkMap(&VarMap{}))
	_ = volume
	if IsTrue(GoIsArray(ohlcv)) {
		volumeIndex := this.SafeString(volume, *(market).At(MkString("type")), MkString("amount"))
		_ = volumeIndex
		return MkArray(&VarArray{
			this.SafeInteger(ohlcv, MkInteger(0)),
			this.SafeNumber(ohlcv, MkInteger(1)),
			this.SafeNumber(ohlcv, MkInteger(2)),
			this.SafeNumber(ohlcv, MkInteger(3)),
			this.SafeNumber(ohlcv, MkInteger(4)),
			this.SafeNumber(ohlcv, volumeIndex),
		})
	} else {
		volumeIndex := this.SafeValue(volume, *(market).At(MkString("type")), MkInteger(6))
		_ = volumeIndex
		return MkArray(&VarArray{
			this.SafeInteger(ohlcv, MkString("id")),
			this.SafeNumber(ohlcv, MkString("open")),
			this.SafeNumber(ohlcv, MkString("high")),
			this.SafeNumber(ohlcv, MkString("low")),
			this.SafeNumber(ohlcv, MkString("close")),
			this.SafeNumber(ohlcv, volumeIndex),
		})
	}
	return MkUndefined()
}

func (this *Bitget) FetchOHLCV(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	method := OpCopy(MkUndefined())
	_ = method
	type_ := *(market).At(MkString("type"))
	_ = type_
	options := this.SafeValue(*this.At(MkString("options")), MkString("timeframes"), MkMap(&VarMap{}))
	_ = options
	intervals := this.SafeValue(options, type_, MkMap(&VarMap{}))
	_ = intervals
	interval := this.SafeValue(intervals, *(*this.At(MkString("timeframes"))).At(timeframe))
	_ = interval
	if IsTrue(*(market).At(MkString("spot"))) {
		method = MkString("dataGetMarketHistoryKline")
		*(request).At(MkString("period")) = OpCopy(interval)
		if IsTrue(OpNotEq2(limit, MkUndefined())) {
			*(request).At(MkString("size")) = OpCopy(limit)
		}
	} else {
		if IsTrue(*(market).At(MkString("swap"))) {
			duration := this.ParseTimeframe(timeframe)
			_ = duration
			method = MkString("capiGetMarketCandles")
			*(request).At(MkString("granularity")) = OpCopy(interval)
			now := this.Milliseconds()
			_ = now
			if IsTrue(OpEq2(since, MkUndefined())) {
				if IsTrue(OpEq2(limit, MkUndefined())) {
					limit = MkInteger(1000)
				}
				*(request).At(MkString("start")) = this.Iso8601(OpSub(now, OpMulti(limit, OpMulti(duration, MkInteger(1000)))))
				*(request).At(MkString("end")) = this.Iso8601(now)
			} else {
				*(request).At(MkString("start")) = this.Iso8601(since)
				if IsTrue(OpEq2(limit, MkUndefined())) {
					*(request).At(MkString("end")) = this.Iso8601(now)
				} else {
					*(request).At(MkString("end")) = this.Iso8601(this.Sum(since, OpMulti(limit, OpMulti(duration, MkInteger(1000)))))
				}
			}
		}
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	candles := OpCopy(response)
	_ = candles
	if IsTrue(OpNot(GoIsArray(response))) {
		candles = this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	}
	return this.ParseOHLCVs(candles, market, timeframe, since, limit)
}

func (this *Bitget) ParseSpotBalance(goArgs ...*Variant) *Variant {
	response := GoGetArg(goArgs, 0, MkUndefined())
	_ = response
	result := MkMap(&VarMap{"info": response})
	_ = result
	data := this.SafeValue(response, MkString("data"))
	_ = data
	balances := this.SafeValue(data, MkString("list"))
	_ = balances
	for i := MkInteger(0); IsTrue(OpLw(i, balances.Length)); OpIncr(&i) {
		balance := *(balances).At(i)
		_ = balance
		currencyId := this.SafeString(balance, MkString("currency"))
		_ = currencyId
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		if IsTrue(OpNot(OpHasMember(code, result))) {
			account := this.Account()
			_ = account
			*(result).At(code) = OpCopy(account)
		}
		type_ := this.SafeValue(balance, MkString("type"))
		_ = type_
		if IsTrue(OpEq2(type_, MkString("trade"))) {
			*(*(result).At(code)).At(MkString("free")) = this.SafeString(balance, MkString("balance"))
		} else {
			if IsTrue(OpOr(OpEq2(type_, MkString("frozen")), OpEq2(type_, MkString("lock")))) {
				used := this.SafeString(*(result).At(code), MkString("used"))
				_ = used
				*(*(result).At(code)).At(MkString("used")) = Precise.StringAdd(used, this.SafeString(balance, MkString("balance")))
			}
		}
	}
	return this.ParseBalance(result)
}

func (this *Bitget) ParseSwapBalance(goArgs ...*Variant) *Variant {
	response := GoGetArg(goArgs, 0, MkUndefined())
	_ = response
	result := MkMap(&VarMap{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
		balance := *(response).At(i)
		_ = balance
		marketId := this.SafeString(balance, MkString("symbol"))
		_ = marketId
		symbol := OpCopy(marketId)
		_ = symbol
		if IsTrue(OpHasMember(marketId, *this.At(MkString("markets_by_id")))) {
			symbol = *(*(*this.At(MkString("markets_by_id"))).At(marketId)).At(MkString("symbol"))
		}
		account := this.Account()
		_ = account
		*(account).At(MkString("total")) = this.SafeString(balance, MkString("equity"))
		*(account).At(MkString("free")) = this.SafeString(balance, MkString("total_avail_balance"))
		*(result).At(symbol) = OpCopy(account)
	}
	return this.ParseBalance(result)
}

func (this *Bitget) FetchAccounts(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	request := MkMap(&VarMap{"method": MkString("accounts")})
	_ = request
	response := this.Call(MkString("apiGetAccountAccounts"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, data.Length)); OpIncr(&i) {
		account := *(data).At(i)
		_ = account
		accountId := this.SafeString(account, MkString("id"))
		_ = accountId
		type_ := this.SafeStringLower(account, MkString("type"))
		_ = type_
		result.Push(MkMap(&VarMap{
			"id":       accountId,
			"type":     type_,
			"currency": MkUndefined(),
			"info":     account,
		}))
	}
	return result
}

func (this *Bitget) FindAccountByType(goArgs ...*Variant) *Variant {
	type_ := GoGetArg(goArgs, 0, MkUndefined())
	_ = type_
	this.LoadMarkets()
	this.LoadAccounts()
	accountsByType := this.GroupBy(*this.At(MkString("accounts")), MkString("type"))
	_ = accountsByType
	accounts := this.SafeValue(accountsByType, type_)
	_ = accounts
	if IsTrue(OpEq2(accounts, MkUndefined())) {
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" findAccountByType() could not find an accountId with type '"), OpAdd(type_, MkString("', specify the 'accountId' parameter instead"))))))
	}
	numAccounts := OpCopy(accounts.Length)
	_ = numAccounts
	if IsTrue(OpGt(numAccounts, MkInteger(1))) {
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" findAccountByType() found more than one accountId with type '"), OpAdd(type_, MkString("', specify the 'accountId' parameter instead"))))))
	}
	return *(accounts).At(MkInteger(0))
}

func (this *Bitget) GetAccountId(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkUndefined())
	_ = params
	this.LoadMarkets()
	this.LoadAccounts()
	defaultAccountId := this.SafeString(*this.At(MkString("options")), MkString("accountId"))
	_ = defaultAccountId
	accountId := this.SafeString(params, MkString("accountId"), defaultAccountId)
	_ = accountId
	if IsTrue(OpNotEq2(accountId, MkUndefined())) {
		return accountId
	}
	defaultType := this.SafeString(*this.At(MkString("options")), MkString("defaultType"), MkString("margin"))
	_ = defaultType
	type_ := this.SafeString(params, MkString("type"), defaultType)
	_ = type_
	params = this.Omit(params, MkString("type"))
	if IsTrue(OpEq2(type_, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" getAccountId() requires an 'accountId' parameter"))))
	}
	account := this.FindAccountByType(type_)
	_ = account
	return *(account).At(MkString("id"))
}

func (this *Bitget) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	this.LoadAccounts()
	defaultType := this.SafeString2(*this.At(MkString("options")), MkString("fetchBalance"), MkString("defaultType"))
	_ = defaultType
	type_ := this.SafeString(params, MkString("type"), defaultType)
	_ = type_
	if IsTrue(OpEq2(type_, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchBalance() requires a 'type' parameter, one of 'spot', 'swap'"))))
	}
	method := OpCopy(MkUndefined())
	_ = method
	query := this.Omit(params, MkString("type"))
	_ = query
	if IsTrue(OpEq2(type_, MkString("spot"))) {
		accountId := this.GetAccountId(params)
		_ = accountId
		method = MkString("apiGetAccountsAccountIdBalance")
		*(query).At(MkString("account_id")) = OpCopy(accountId)
		*(query).At(MkString("method")) = MkString("balance")
	} else {
		if IsTrue(OpEq2(type_, MkString("swap"))) {
			method = MkString("swapGetAccountAccounts")
		}
	}
	response := this.Call(method, query)
	_ = response
	return this.ParseBalanceByType(type_, response)
}

func (this *Bitget) ParseBalanceByType(goArgs ...*Variant) *Variant {
	type_ := GoGetArg(goArgs, 0, MkUndefined())
	_ = type_
	response := GoGetArg(goArgs, 1, MkUndefined())
	_ = response
	if IsTrue(OpEq2(type_, MkString("spot"))) {
		return this.ParseSpotBalance(response)
	} else {
		if IsTrue(OpEq2(type_, MkString("swap"))) {
			return this.ParseSwapBalance(response)
		}
	}
	panic(NewNotSupported(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" fetchBalance does not support the '"), OpAdd(type_, MkString("' type (the type must be one of 'account', 'spot', 'margin', 'futures', 'swap')"))))))
	return MkUndefined()
}

func (this *Bitget) ParseOrderStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"submitted":        MkString("open"),
		"partial-filled":   MkString("open"),
		"partial-canceled": MkString("canceled"),
		"filled":           MkString("closed"),
		"canceled":         MkString("canceled"),
		"-2":               MkString("failed"),
		"-1":               MkString("canceled"),
		"0":                MkString("open"),
		"1":                MkString("open"),
		"2":                MkString("closed"),
		"3":                MkString("open"),
		"4":                MkString("canceled"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Bitget) ParseOrderSide(goArgs ...*Variant) *Variant {
	side := GoGetArg(goArgs, 0, MkUndefined())
	_ = side
	sides := MkMap(&VarMap{
		"buy-market":  MkString("buy"),
		"sell-market": MkString("sell"),
		"buy-limit":   MkString("buy"),
		"sell-limit":  MkString("sell"),
		"1":           MkString("long"),
		"2":           MkString("short"),
		"3":           MkString("long"),
		"4":           MkString("short"),
	})
	_ = sides
	return this.SafeString(sides, side, side)
}

func (this *Bitget) ParseOrderType(goArgs ...*Variant) *Variant {
	type_ := GoGetArg(goArgs, 0, MkUndefined())
	_ = type_
	types := MkMap(&VarMap{
		"buy-market":  MkString("market"),
		"sell-market": MkString("market"),
		"buy-limit":   MkString("limit"),
		"sell-limit":  MkString("limit"),
		"1":           MkString("open"),
		"2":           MkString("open"),
		"3":           MkString("close"),
		"4":           MkString("close"),
	})
	_ = types
	return this.SafeString(types, type_, type_)
}

func (this *Bitget) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	id := this.SafeString(order, MkString("order_id"))
	_ = id
	id = this.SafeString2(order, MkString("id"), MkString("data"), id)
	timestamp := this.SafeInteger2(order, MkString("created_at"), MkString("createTime"))
	_ = timestamp
	type_ := this.SafeString(order, MkString("type"))
	_ = type_
	side := this.ParseOrderSide(type_)
	_ = side
	type_ = this.ParseOrderType(type_)
	symbol := OpCopy(MkUndefined())
	_ = symbol
	marketId := this.SafeString(order, MkString("symbol"))
	_ = marketId
	if IsTrue(OpNotEq2(marketId, MkUndefined())) {
		if IsTrue(OpHasMember(marketId, *this.At(MkString("markets_by_id")))) {
			market = *(*this.At(MkString("markets_by_id"))).At(marketId)
		} else {
			symbol = marketId.ToUpperCase()
		}
	}
	if IsTrue(OpAnd(OpEq2(symbol, MkUndefined()), OpNotEq2(market, MkUndefined()))) {
		symbol = *(market).At(MkString("symbol"))
	}
	amount := this.SafeNumber2(order, MkString("amount"), MkString("size"))
	_ = amount
	filled := this.SafeNumber2(order, MkString("filled_amount"), MkString("filled_qty"))
	_ = filled
	cost := this.SafeNumber(order, MkString("filled_cash_amount"))
	_ = cost
	price := this.SafeNumber(order, MkString("price"))
	_ = price
	average := this.SafeNumber(order, MkString("price_avg"))
	_ = average
	status := this.ParseOrderStatus(this.SafeString2(order, MkString("state"), MkString("status")))
	_ = status
	feeCost := this.SafeNumber2(order, MkString("filled_fees"), MkString("fee"))
	_ = feeCost
	fee := OpCopy(MkUndefined())
	_ = fee
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		feeCurrency := OpCopy(MkUndefined())
		_ = feeCurrency
		fee = MkMap(&VarMap{
			"cost":     feeCost,
			"currency": feeCurrency,
		})
	}
	clientOrderId := this.SafeString(order, MkString("client_oid"))
	_ = clientOrderId
	return this.SafeOrder(MkMap(&VarMap{
		"info":               order,
		"id":                 id,
		"clientOrderId":      clientOrderId,
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

func (this *Bitget) CreateOrder(goArgs ...*Variant) *Variant {
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
	this.LoadAccounts()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	clientOrderId := this.SafeString2(params, MkString("client_oid"), MkString("clientOrderId"), this.Uuid())
	_ = clientOrderId
	params = this.Omit(params, MkArray(&VarArray{
		MkString("client_oid"),
		MkString("clientOrderId"),
	}))
	method := OpCopy(MkUndefined())
	_ = method
	if IsTrue(*(market).At(MkString("spot"))) {
		accountId := this.GetAccountId(MkMap(&VarMap{"type": *(market).At(MkString("type"))}))
		_ = accountId
		method = MkString("apiPostOrderOrdersPlace")
		*(request).At(MkString("account_id")) = OpCopy(accountId)
		*(request).At(MkString("method")) = MkString("place")
		*(request).At(MkString("type")) = OpAdd(side, OpAdd(MkString("-"), type_))
		if IsTrue(OpEq2(type_, MkString("limit"))) {
			*(request).At(MkString("amount")) = this.AmountToPrecision(symbol, amount)
			*(request).At(MkString("price")) = this.PriceToPrecision(symbol, price)
		} else {
			if IsTrue(OpEq2(type_, MkString("market"))) {
				if IsTrue(OpEq2(side, MkString("buy"))) {
					cost := this.SafeNumber(params, MkString("amount"))
					_ = cost
					createMarketBuyOrderRequiresPrice := this.SafeValue(*this.At(MkString("options")), MkString("createMarketBuyOrderRequiresPrice"), MkBool(true))
					_ = createMarketBuyOrderRequiresPrice
					if IsTrue(createMarketBuyOrderRequiresPrice) {
						if IsTrue(OpNotEq2(price, MkUndefined())) {
							if IsTrue(OpEq2(cost, MkUndefined())) {
								cost = OpMulti(amount, price)
							}
						} else {
							if IsTrue(OpEq2(cost, MkUndefined())) {
								panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")), MkString(" createOrder() requires the price argument with market buy orders to calculate total order cost (amount to spend), where cost = amount * price. Supply a price argument to createOrder() call if you want the cost to be calculated for you from price and amount, or, alternatively, add .options['createMarketBuyOrderRequiresPrice'] = false and supply the total cost value in the 'amount' argument or in the 'amount' extra parameter (the exchange-specific behaviour)"))))
							}
						}
					} else {
						cost = OpTrinary(OpEq2(cost, MkUndefined()), amount, cost)
					}
					*(request).At(MkString("amount")) = this.CostToPrecision(symbol, cost)
				} else {
					if IsTrue(OpEq2(side, MkString("sell"))) {
						*(request).At(MkString("amount")) = this.AmountToPrecision(symbol, amount)
					}
				}
			}
		}
	} else {
		if IsTrue(*(market).At(MkString("swap"))) {
			*(request).At(MkString("order_type")) = MkString("0")
			*(request).At(MkString("client_oid")) = OpCopy(clientOrderId)
			orderType := this.SafeString(params, MkString("type"))
			_ = orderType
			if IsTrue(OpEq2(orderType, MkUndefined())) {
				panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" createOrder() requires a type parameter, '1' = open long, '2' = open short, '3' = close long, '4' = close short for "), OpAdd(*(market).At(MkString("type")), MkString(" orders"))))))
			}
			*(request).At(MkString("size")) = this.AmountToPrecision(symbol, amount)
			*(request).At(MkString("type")) = OpCopy(orderType)
			if IsTrue(OpEq2(type_, MkString("limit"))) {
				*(request).At(MkString("match_price")) = MkString("0")
				*(request).At(MkString("price")) = this.PriceToPrecision(symbol, price)
			} else {
				if IsTrue(OpEq2(type_, MkString("market"))) {
					*(request).At(MkString("match_price")) = MkString("1")
				}
			}
			method = MkString("swapPostOrderPlaceOrder")
		}
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	return this.ParseOrder(response, market)
}

func (this *Bitget) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := OpCopy(MkUndefined())
	_ = market
	type_ := OpCopy(MkUndefined())
	_ = type_
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		defaultType := this.SafeString2(*this.At(MkString("options")), MkString("cancelOrder"), MkString("defaultType"))
		_ = defaultType
		type_ := this.SafeString(params, MkString("type"), defaultType)
		_ = type_
		if IsTrue(OpEq2(type_, MkString("spot"))) {
			if IsTrue(OpEq2(symbol, MkUndefined())) {
				panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" cancelOrder() requires a symbol argument for spot orders"))))
			}
		}
	} else {
		market = this.Market(symbol)
		type_ = *(market).At(MkString("type"))
	}
	query := this.Omit(params, MkString("type"))
	_ = query
	method := OpCopy(MkUndefined())
	_ = method
	request := MkMap(&VarMap{})
	_ = request
	if IsTrue(OpEq2(type_, MkString("spot"))) {
		method = MkString("apiPostOrderOrdersOrderIdSubmitcancel")
		*(request).At(MkString("order_id")) = OpCopy(id)
		*(request).At(MkString("method")) = MkString("submitcancel")
	} else {
		if IsTrue(OpEq2(type_, MkString("swap"))) {
			method = MkString("swapPostOrderCancelOrder")
			*(request).At(MkString("orderId")) = OpCopy(id)
			*(request).At(MkString("symbol")) = *(market).At(MkString("id"))
		}
	}
	response := this.Call(method, this.Extend(request, query))
	_ = response
	return this.ParseOrder(response, market)
}

func (this *Bitget) CancelOrders(goArgs ...*Variant) *Variant {
	ids := GoGetArg(goArgs, 0, MkUndefined())
	_ = ids
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" cancelOrders() requires a symbol argument"))))
	}
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	type_ := this.SafeString(params, MkString("type"), *(market).At(MkString("type")))
	_ = type_
	if IsTrue(OpEq2(type_, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" cancelOrders() requires a type parameter (one of 'spot', 'swap')."))))
	}
	params = this.Omit(params, MkString("type"))
	request := MkMap(&VarMap{})
	_ = request
	method := OpCopy(MkUndefined())
	_ = method
	if IsTrue(OpEq2(type_, MkString("spot"))) {
		method = MkString("apiPostOrderOrdersBatchcancel")
		*(request).At(MkString("method")) = MkString("batchcancel")
		jsonIds := this.Json(ids)
		_ = jsonIds
		parts := jsonIds.Split(MkString("\""))
		_ = parts
		*(request).At(MkString("order_ids")) = parts.Join(MkString(""))
	} else {
		if IsTrue(OpEq2(type_, MkString("swap"))) {
			method = MkString("swapPostOrderCancelBatchOrders")
			*(request).At(MkString("symbol")) = *(market).At(MkString("id"))
			*(request).At(MkString("ids")) = OpCopy(ids)
		}
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	return response
}

func (this *Bitget) FetchOrder(goArgs ...*Variant) *Variant {
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
	type_ := this.SafeString(params, MkString("type"), *(market).At(MkString("type")))
	_ = type_
	if IsTrue(OpEq2(type_, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchOrder() requires a type parameter (one of 'spot', 'swap')."))))
	}
	method := OpCopy(MkUndefined())
	_ = method
	request := MkMap(&VarMap{})
	_ = request
	if IsTrue(OpEq2(type_, MkString("spot"))) {
		clientOid := this.SafeString(params, MkString("client_oid"))
		_ = clientOid
		if IsTrue(OpNotEq2(clientOid, MkUndefined())) {
			method = MkString("apiPostOrderOrdersClientOid")
			*(request).At(MkString("client_oid")) = OpCopy(clientOid)
		} else {
			method = MkString("apiPostOrderOrdersOrderId")
			*(request).At(MkString("order_id")) = OpCopy(id)
		}
		*(request).At(MkString("method")) = MkString("getOrder")
	} else {
		if IsTrue(OpEq2(type_, MkString("swap"))) {
			method = MkString("swapGetOrderDetail")
			*(request).At(MkString("symbol")) = *(market).At(MkString("id"))
			*(request).At(MkString("orderId")) = OpCopy(id)
		}
	}
	query := this.Omit(params, MkString("type"))
	_ = query
	response := this.Call(method, this.Extend(request, query))
	_ = response
	data := this.SafeValue(response, MkString("data"), response)
	_ = data
	return this.ParseOrder(data, market)
}

func (this *Bitget) FetchOpenOrders(goArgs ...*Variant) *Variant {
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
	type_ := this.SafeString(params, MkString("type"), *(market).At(MkString("type")))
	_ = type_
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	method := OpCopy(MkUndefined())
	_ = method
	if IsTrue(OpEq2(type_, MkString("spot"))) {
		method = MkString("apiGetOrderOrdersOpenOrders")
		*(request).At(MkString("method")) = MkString("openOrders")
		if IsTrue(OpEq2(limit, MkUndefined())) {
			*(request).At(MkString("size")) = OpCopy(limit)
		}
	} else {
		if IsTrue(OpEq2(type_, MkString("swap"))) {
			method = MkString("swapGetOrderOrders")
			*(request).At(MkString("status")) = MkString("3")
			*(request).At(MkString("from")) = MkString("1")
			*(request).At(MkString("to")) = MkString("1")
			if IsTrue(OpEq2(limit, MkUndefined())) {
				*(request).At(MkString("limit")) = MkInteger(100)
			}
		}
	}
	query := this.Omit(params, MkString("type"))
	_ = query
	response := this.Call(method, this.Extend(request, query))
	_ = response
	data := OpCopy(response)
	_ = data
	if IsTrue(OpNot(GoIsArray(response))) {
		data = this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	}
	return this.ParseOrders(data, market, MkUndefined(), limit)
}

func (this *Bitget) FetchClosedOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchClosedOrders() requires a symbol argument"))))
	}
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	type_ := this.SafeString(params, MkString("type"), *(market).At(MkString("type")))
	_ = type_
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	method := OpCopy(MkUndefined())
	_ = method
	if IsTrue(OpEq2(type_, MkString("spot"))) {
		method = MkString("apiGetOrderOrdersHistory")
		if IsTrue(OpNotEq2(since, MkUndefined())) {
			*(request).At(MkString("start_time")) = OpCopy(since)
		}
		*(request).At(MkString("method")) = MkString("openOrders")
		if IsTrue(OpEq2(limit, MkUndefined())) {
			*(request).At(MkString("size")) = OpCopy(limit)
		}
	} else {
		if IsTrue(OpEq2(type_, MkString("swap"))) {
			method = MkString("swapGetOrderOrders")
			*(request).At(MkString("status")) = MkString("2")
			*(request).At(MkString("from")) = MkString("1")
			*(request).At(MkString("to")) = MkString("1")
			if IsTrue(OpEq2(limit, MkUndefined())) {
				*(request).At(MkString("limit")) = MkInteger(100)
			}
		}
	}
	query := this.Omit(params, MkString("type"))
	_ = query
	response := this.Call(method, this.Extend(request, query))
	_ = response
	data := OpCopy(response)
	_ = data
	if IsTrue(OpNot(GoIsArray(response))) {
		data = this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	}
	return this.ParseOrders(data, market, MkUndefined(), limit)
}

func (this *Bitget) FetchDeposits(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{
		"currency": *(currency).At(MkString("id")),
		"method":   MkString("deposit_withdraw"),
		"type":     MkString("deposit"),
		"size":     MkInteger(12),
	})
	_ = request
	response := this.Call(MkString("apiGetOrderDepositWithdraw"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseTransactions(data, currency, since, limit, params)
}

func (this *Bitget) FetchWithdrawals(goArgs ...*Variant) *Variant {
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
	request := MkMap(&VarMap{
		"currency": *(currency).At(MkString("id")),
		"method":   MkString("deposit_withdraw"),
		"type":     MkString("withdraw"),
		"size":     MkInteger(12),
	})
	_ = request
	response := this.Call(MkString("apiGetOrderDepositWithdraw"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseTransactions(data, currency, since, limit, params)
}

func (this *Bitget) ParseTransactionStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"WaitForOperation": MkString("pending"),
		"OperationLock":    MkString("pending"),
		"OperationSuccess": MkString("ok"),
		"Cancel":           MkString("canceled"),
		"Sure":             MkString("ok"),
		"Fail":             MkString("failed"),
		"WaitForChainSure": MkString("ok"),
		"WAIT_0":           MkString("pending"),
		"WAIT_1":           MkString("pending"),
		"DATA_CHANGE":      MkString("pending"),
		"SUCCESS":          MkString("ok"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Bitget) ParseTransaction(goArgs ...*Variant) *Variant {
	transaction := GoGetArg(goArgs, 0, MkUndefined())
	_ = transaction
	currency := GoGetArg(goArgs, 1, MkUndefined())
	_ = currency
	id := this.SafeString(transaction, MkString("id"))
	_ = id
	address := this.SafeString(transaction, MkString("address"))
	_ = address
	tag := this.SafeString(transaction, MkString("address_tag"))
	_ = tag
	tagFrom := OpCopy(MkUndefined())
	_ = tagFrom
	tagTo := OpCopy(tag)
	_ = tagTo
	addressFrom := OpCopy(MkUndefined())
	_ = addressFrom
	addressTo := OpCopy(address)
	_ = addressTo
	type_ := this.SafeString(transaction, MkString("type"))
	_ = type_
	if IsTrue(OpEq2(type_, MkString("withdraw"))) {
		type_ = MkString("withdrawal")
	} else {
		if IsTrue(OpEq2(type_, MkString("deposit"))) {
			type_ = MkString("deposit")
		}
	}
	currencyId := this.SafeString(transaction, MkString("currency"))
	_ = currencyId
	code := this.SafeCurrencyCode(currencyId)
	_ = code
	amount := this.SafeNumber(transaction, MkString("amount"))
	_ = amount
	status := this.ParseTransactionStatus(this.SafeString(transaction, MkString("state")))
	_ = status
	txid := this.SafeString(transaction, MkString("tx_hash"))
	_ = txid
	timestamp := this.SafeInteger(transaction, MkString("created_at"))
	_ = timestamp
	updated := this.SafeInteger(transaction, MkString("updated_at"))
	_ = updated
	feeCost := this.SafeNumber(transaction, MkString("fee"))
	_ = feeCost
	fee := OpCopy(MkUndefined())
	_ = fee
	if IsTrue(OpNotEq2(feeCost, MkUndefined())) {
		fee = MkMap(&VarMap{
			"currency": code,
			"cost":     feeCost,
		})
	}
	return MkMap(&VarMap{
		"info":        transaction,
		"id":          id,
		"currency":    code,
		"amount":      amount,
		"addressFrom": addressFrom,
		"addressTo":   addressTo,
		"address":     address,
		"tagFrom":     tagFrom,
		"tagTo":       tagTo,
		"tag":         tag,
		"status":      status,
		"type":        type_,
		"updated":     updated,
		"txid":        txid,
		"timestamp":   timestamp,
		"datetime":    this.Iso8601(timestamp),
		"fee":         fee,
	})
}

func (this *Bitget) FetchMyTrades(goArgs ...*Variant) *Variant {
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
	type_ := this.SafeString(params, MkString("type"), *(market).At(MkString("type")))
	_ = type_
	query := this.Omit(params, MkString("type"))
	_ = query
	if IsTrue(OpEq2(type_, MkString("swap"))) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" fetchMyTrades() is not supported for "), OpAdd(type_, MkString(" type"))))))
	}
	request := MkMap(&VarMap{
		"symbol": *(market).At(MkString("id")),
		"method": MkString("matchresults"),
	})
	_ = request
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("start_date")) = this.Ymd(since)
		end := this.Sum(since, OpMulti(MkInteger(2), OpMulti(MkInteger(24), OpMulti(MkInteger(60), OpMulti(MkInteger(60), MkInteger(1000))))))
		_ = end
		*(request).At(MkString("end_date")) = this.Ymd(end)
	}
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("size")) = OpCopy(limit)
	}
	response := this.Call(MkString("apiPostOrderMatchresults"), this.Extend(request, query))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	return this.ParseTrades(data, market, since, limit)
}

func (this *Bitget) FetchOrderTrades(goArgs ...*Variant) *Variant {
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
	type_ := this.SafeString(params, MkString("type"), *(market).At(MkString("type")))
	_ = type_
	params = this.Omit(params, MkString("type"))
	method := OpCopy(MkUndefined())
	_ = method
	request := MkMap(&VarMap{})
	_ = request
	if IsTrue(OpEq2(type_, MkString("spot"))) {
		*(request).At(MkString("order_id")) = OpCopy(id)
		*(request).At(MkString("method")) = MkString("matchresults")
		method = MkString("apiPostOrderOrdersOrderIdMatchresults")
	} else {
		if IsTrue(OpEq2(type_, MkString("swap"))) {
			*(request).At(MkString("orderId")) = OpCopy(id)
			*(request).At(MkString("symbol")) = *(market).At(MkString("id"))
			method = MkString("swapGetOrderFills")
		}
	}
	response := this.Call(method, this.Extend(request, params))
	_ = response
	data := OpCopy(response)
	_ = data
	if IsTrue(OpNot(GoIsArray(data))) {
		data = this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	}
	return this.ParseTrades(data, market, since, limit)
}

func (this *Bitget) FetchPosition(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"symbol": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("swapGetPositionSinglePosition"), this.Extend(request, params))
	_ = response
	return response
}

func (this *Bitget) FetchPositions(goArgs ...*Variant) *Variant {
	symbols := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbols
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	response := this.Call(MkString("swapGetPositionAllPosition"), params)
	_ = response
	return response
}

func (this *Bitget) Sign(goArgs ...*Variant) *Variant {
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
	request := OpAdd(MkString("/"), this.ImplodeParams(path, params))
	_ = request
	if IsTrue(OpOr(OpEq2(api, MkString("capi")), OpEq2(api, MkString("swap")))) {
		request = OpAdd(MkString("/api/swap/"), OpAdd(*this.At(MkString("version")), request))
	} else {
		request = OpAdd(MkString("/"), OpAdd(api, OpAdd(MkString("/v1"), request)))
	}
	query := this.Omit(params, this.ExtractParams(path))
	_ = query
	url := OpAdd(this.ImplodeHostname(*(*(*this.At(MkString("urls"))).At(MkString("api"))).At(api)), request)
	_ = url
	if IsTrue(OpOr(OpEq2(api, MkString("data")), OpEq2(api, MkString("capi")))) {
		if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
			OpAddAssign(&url, OpAdd(MkString("?"), this.Urlencode(query)))
		}
	} else {
		if IsTrue(OpEq2(api, MkString("swap"))) {
			this.CheckRequiredCredentials()
			timestamp := (this.Milliseconds()).Call(MkString("toString"))
			_ = timestamp
			auth := OpAdd(timestamp, OpAdd(method, request))
			_ = auth
			if IsTrue(OpEq2(method, MkString("POST"))) {
				body = this.Json(params)
				OpAddAssign(&auth, body)
			} else {
				if IsTrue(*(GoGetKeys(params)).At(MkString("length"))) {
					query := this.Urlencode(this.Keysort(params))
					_ = query
					OpAddAssign(&url, OpAdd(MkString("?"), query))
					OpAddAssign(&auth, OpAdd(MkString("?"), query))
				}
			}
			signature := this.Hmac(this.Encode(auth), this.Encode(*this.At(MkString("secret"))), MkString("sha256"), MkString("base64"))
			_ = signature
			headers = MkMap(&VarMap{
				"ACCESS-KEY":        *this.At(MkString("apiKey")),
				"ACCESS-SIGN":       signature,
				"ACCESS-TIMESTAMP":  timestamp,
				"ACCESS-PASSPHRASE": *this.At(MkString("password")),
			})
			if IsTrue(OpEq2(method, MkString("POST"))) {
				*(headers).At(MkString("Content-Type")) = MkString("application/json")
			}
		} else {
			if IsTrue(OpEq2(api, MkString("api"))) {
				timestamp := (this.Milliseconds()).Call(MkString("toString"))
				_ = timestamp
				auth := MkString("")
				_ = auth
				query = this.Keysort(query)
				auth = this.Rawencode(query)
				hash := this.Hash(this.Encode(*this.At(MkString("secret"))), MkString("sha1"))
				_ = hash
				signed := OpCopy(auth)
				_ = signed
				signature := this.Hmac(this.Encode(auth), this.Encode(hash), MkString("md5"))
				_ = signature
				if IsTrue(OpGt(auth.Length, MkInteger(0))) {
					OpAddAssign(&signed, MkString("&"))
				}
				OpAddAssign(&signed, OpAdd(MkString("sign="), OpAdd(signature, OpAdd(MkString("&req_time="), OpAdd(timestamp, OpAdd(MkString("&accesskey="), *this.At(MkString("apiKey"))))))))
				if IsTrue(OpEq2(method, MkString("GET"))) {
					if IsTrue(*(GoGetKeys(query)).At(MkString("length"))) {
						OpAddAssign(&url, OpAdd(MkString("?"), signed))
					}
				} else {
					if IsTrue(OpEq2(method, MkString("POST"))) {
						OpAddAssign(&url, OpAdd(MkString("?sign="), OpAdd(signature, OpAdd(MkString("&req_time="), OpAdd(timestamp, OpAdd(MkString("&accesskey="), *this.At(MkString("apiKey"))))))))
						body = OpCopy(auth)
						headers = MkMap(&VarMap{"Content-Type": MkString("application/x-www-form-urlencoded")})
					}
				}
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

func (this *Bitget) HandleErrors(goArgs ...*Variant) *Variant {
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
		return MkUndefined()
	}
	message := this.SafeString(response, MkString("err_msg"))
	_ = message
	errorCode := this.SafeString2(response, MkString("code"), MkString("err_code"))
	_ = errorCode
	feedback := OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))
	_ = feedback
	nonEmptyMessage := OpAnd(OpNotEq2(message, MkUndefined()), OpNotEq2(message, MkString("")))
	_ = nonEmptyMessage
	if IsTrue(nonEmptyMessage) {
		this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("exact")), message, feedback)
		this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("broad")), message, feedback)
	}
	nonZeroErrorCode := OpAnd(OpNotEq2(errorCode, MkUndefined()), OpNotEq2(errorCode, MkString("00000")))
	_ = nonZeroErrorCode
	if IsTrue(nonZeroErrorCode) {
		this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions"))).At(MkString("exact")), errorCode, feedback)
	}
	if IsTrue(OpOr(nonZeroErrorCode, nonEmptyMessage)) {
		panic(NewExchangeError(feedback))
	}
	return MkUndefined()
}
