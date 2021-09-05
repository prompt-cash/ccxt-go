package ccxt

import (
)

type Okex3 struct {
	*ExchangeBase
}

var _ Exchange = (*Okex3)(nil)

func init() {
	exchange := &Okex3{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Okex3) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("okex3") ,
            "name":MkString("OKEX") ,
            "countries":MkArray(&VarArray{
                MkString("CN") ,
                }),
            "version":MkString("v3") ,
            "rateLimit":MkInteger(1000) ,
            "pro":MkBool(true) ,
            "has":MkMap(&VarMap{
                "cancelOrder":MkBool(true) ,
                "CORS":MkBool(false) ,
                "createOrder":MkBool(true) ,
                "fetchBalance":MkBool(true) ,
                "fetchClosedOrders":MkBool(true) ,
                "fetchCurrencies":MkBool(false) ,
                "fetchDepositAddress":MkBool(true) ,
                "fetchDeposits":MkBool(true) ,
                "fetchLedger":MkBool(true) ,
                "fetchMarkets":MkBool(true) ,
                "fetchMyTrades":MkBool(true) ,
                "fetchOHLCV":MkBool(true) ,
                "fetchOpenOrders":MkBool(true) ,
                "fetchOrder":MkBool(true) ,
                "fetchOrderBook":MkBool(true) ,
                "fetchOrders":MkBool(false) ,
                "fetchOrderTrades":MkBool(true) ,
                "fetchTime":MkBool(true) ,
                "fetchTicker":MkBool(true) ,
                "fetchTickers":MkBool(true) ,
                "fetchTrades":MkBool(true) ,
                "fetchTransactions":MkBool(false) ,
                "fetchWithdrawals":MkBool(true) ,
                "futures":MkBool(true) ,
                "withdraw":MkBool(true) ,
                }),
            "timeframes":MkMap(&VarMap{
                "1m":MkString("60") ,
                "3m":MkString("180") ,
                "5m":MkString("300") ,
                "15m":MkString("900") ,
                "30m":MkString("1800") ,
                "1h":MkString("3600") ,
                "2h":MkString("7200") ,
                "4h":MkString("14400") ,
                "6h":MkString("21600") ,
                "12h":MkString("43200") ,
                "1d":MkString("86400") ,
                "1w":MkString("604800") ,
                "1M":MkString("2678400") ,
                "3M":MkString("8035200") ,
                "6M":MkString("16070400") ,
                "1y":MkString("31536000") ,
                }),
            "hostname":MkString("okex.com") ,
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/1294454/32552768-0d6dd3c6-c4a6-11e7-90f8-c043b64756a7.jpg") ,
                "api":MkMap(&VarMap{"rest":MkString("https://www.{hostname}") }),
                "www":MkString("https://www.okex.com") ,
                "doc":MkString("https://www.okex.com/docs/en/") ,
                "fees":MkString("https://www.okex.com/pages/products/fees.html") ,
                "referral":MkString("https://www.okex.com/join/1888677") ,
                "test":MkMap(&VarMap{"rest":MkString("https://testnet.okex.com") }),
                }),
            "api":MkMap(&VarMap{
                "general":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("time") ,
                        })}),
                "account":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("wallet") ,
                        MkString("sub-account") ,
                        MkString("asset-valuation") ,
                        MkString("wallet/{currency}") ,
                        MkString("withdrawal/history") ,
                        MkString("withdrawal/history/{currency}") ,
                        MkString("ledger") ,
                        MkString("deposit/address") ,
                        MkString("deposit/history") ,
                        MkString("deposit/history/{currency}") ,
                        MkString("currencies") ,
                        MkString("withdrawal/fee") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("transfer") ,
                        MkString("withdrawal") ,
                        }),
                    }),
                "spot":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("accounts") ,
                        MkString("accounts/{currency}") ,
                        MkString("accounts/{currency}/ledger") ,
                        MkString("orders") ,
                        MkString("amend_order/{instrument_id}") ,
                        MkString("orders_pending") ,
                        MkString("orders/{order_id}") ,
                        MkString("orders/{client_oid}") ,
                        MkString("trade_fee") ,
                        MkString("fills") ,
                        MkString("algo") ,
                        MkString("instruments") ,
                        MkString("instruments/{instrument_id}/book") ,
                        MkString("instruments/ticker") ,
                        MkString("instruments/{instrument_id}/ticker") ,
                        MkString("instruments/{instrument_id}/trades") ,
                        MkString("instruments/{instrument_id}/candles") ,
                        MkString("instruments/{instrument_id}/history/candles") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("order_algo") ,
                        MkString("orders") ,
                        MkString("batch_orders") ,
                        MkString("cancel_orders/{order_id}") ,
                        MkString("cancel_orders/{client_oid}") ,
                        MkString("cancel_batch_algos") ,
                        MkString("cancel_batch_orders") ,
                        }),
                    }),
                "margin":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("accounts") ,
                        MkString("accounts/{instrument_id}") ,
                        MkString("accounts/{instrument_id}/ledger") ,
                        MkString("accounts/availability") ,
                        MkString("accounts/{instrument_id}/availability") ,
                        MkString("accounts/borrowed") ,
                        MkString("accounts/{instrument_id}/borrowed") ,
                        MkString("orders") ,
                        MkString("accounts/{instrument_id}/leverage") ,
                        MkString("orders/{order_id}") ,
                        MkString("orders/{client_oid}") ,
                        MkString("orders_pending") ,
                        MkString("fills") ,
                        MkString("instruments/{instrument_id}/mark_price") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("accounts/borrow") ,
                        MkString("accounts/repayment") ,
                        MkString("orders") ,
                        MkString("batch_orders") ,
                        MkString("cancel_orders") ,
                        MkString("cancel_orders/{order_id}") ,
                        MkString("cancel_orders/{client_oid}") ,
                        MkString("cancel_batch_orders") ,
                        MkString("accounts/{instrument_id}/leverage") ,
                        }),
                    }),
                "futures":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("position") ,
                        MkString("{instrument_id}/position") ,
                        MkString("accounts") ,
                        MkString("accounts/{underlying}") ,
                        MkString("accounts/{underlying}/leverage") ,
                        MkString("accounts/{underlying}/ledger") ,
                        MkString("order_algo/{instrument_id}") ,
                        MkString("orders/{instrument_id}") ,
                        MkString("orders/{instrument_id}/{order_id}") ,
                        MkString("orders/{instrument_id}/{client_oid}") ,
                        MkString("fills") ,
                        MkString("trade_fee") ,
                        MkString("accounts/{instrument_id}/holds") ,
                        MkString("order_algo/{instrument_id}") ,
                        MkString("instruments") ,
                        MkString("instruments/{instrument_id}/book") ,
                        MkString("instruments/ticker") ,
                        MkString("instruments/{instrument_id}/ticker") ,
                        MkString("instruments/{instrument_id}/trades") ,
                        MkString("instruments/{instrument_id}/candles") ,
                        MkString("instruments/{instrument_id}/history/candles") ,
                        MkString("instruments/{instrument_id}/index") ,
                        MkString("rate") ,
                        MkString("instruments/{instrument_id}/estimated_price") ,
                        MkString("instruments/{instrument_id}/open_interest") ,
                        MkString("instruments/{instrument_id}/price_limit") ,
                        MkString("instruments/{instrument_id}/mark_price") ,
                        MkString("instruments/{instrument_id}/liquidation") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("accounts/{underlying}/leverage") ,
                        MkString("order") ,
                        MkString("amend_order/{instrument_id}") ,
                        MkString("orders") ,
                        MkString("cancel_order/{instrument_id}/{order_id}") ,
                        MkString("cancel_order/{instrument_id}/{client_oid}") ,
                        MkString("cancel_batch_orders/{instrument_id}") ,
                        MkString("accounts/margin_mode") ,
                        MkString("close_position") ,
                        MkString("cancel_all") ,
                        MkString("order_algo") ,
                        MkString("cancel_algos") ,
                        }),
                    }),
                "swap":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("position") ,
                        MkString("{instrument_id}/position") ,
                        MkString("accounts") ,
                        MkString("{instrument_id}/accounts") ,
                        MkString("accounts/{instrument_id}/settings") ,
                        MkString("accounts/{instrument_id}/ledger") ,
                        MkString("orders/{instrument_id}") ,
                        MkString("orders/{instrument_id}/{order_id}") ,
                        MkString("orders/{instrument_id}/{client_oid}") ,
                        MkString("fills") ,
                        MkString("accounts/{instrument_id}/holds") ,
                        MkString("trade_fee") ,
                        MkString("order_algo/{instrument_id}") ,
                        MkString("instruments") ,
                        MkString("instruments/{instrument_id}/depth") ,
                        MkString("instruments/ticker") ,
                        MkString("instruments/{instrument_id}/ticker") ,
                        MkString("instruments/{instrument_id}/trades") ,
                        MkString("instruments/{instrument_id}/candles") ,
                        MkString("instruments/{instrument_id}/history/candles") ,
                        MkString("instruments/{instrument_id}/index") ,
                        MkString("rate") ,
                        MkString("instruments/{instrument_id}/open_interest") ,
                        MkString("instruments/{instrument_id}/price_limit") ,
                        MkString("instruments/{instrument_id}/liquidation") ,
                        MkString("instruments/{instrument_id}/funding_time") ,
                        MkString("instruments/{instrument_id}/mark_price") ,
                        MkString("instruments/{instrument_id}/historical_funding_rate") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("accounts/{instrument_id}/leverage") ,
                        MkString("order") ,
                        MkString("amend_order/{instrument_id}") ,
                        MkString("orders") ,
                        MkString("cancel_order/{instrument_id}/{order_id}") ,
                        MkString("cancel_order/{instrument_id}/{client_oid}") ,
                        MkString("cancel_batch_orders/{instrument_id}") ,
                        MkString("order_algo") ,
                        MkString("cancel_algos") ,
                        MkString("close_position") ,
                        MkString("cancel_all") ,
                        MkString("order_algo") ,
                        MkString("cancel_algos") ,
                        }),
                    }),
                "option":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("accounts") ,
                        MkString("position") ,
                        MkString("{underlying}/position") ,
                        MkString("accounts/{underlying}") ,
                        MkString("orders/{underlying}") ,
                        MkString("fills/{underlying}") ,
                        MkString("accounts/{underlying}/ledger") ,
                        MkString("trade_fee") ,
                        MkString("orders/{underlying}/{order_id}") ,
                        MkString("orders/{underlying}/{client_oid}") ,
                        MkString("underlying") ,
                        MkString("instruments/{underlying}") ,
                        MkString("instruments/{underlying}/summary") ,
                        MkString("instruments/{underlying}/summary/{instrument_id}") ,
                        MkString("instruments/{instrument_id}/book") ,
                        MkString("instruments/{instrument_id}/trades") ,
                        MkString("instruments/{instrument_id}/ticker") ,
                        MkString("instruments/{instrument_id}/candles") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("order") ,
                        MkString("orders") ,
                        MkString("cancel_order/{underlying}/{order_id}") ,
                        MkString("cancel_order/{underlying}/{client_oid}") ,
                        MkString("cancel_batch_orders/{underlying}") ,
                        MkString("amend_order/{underlying}") ,
                        MkString("amend_batch_orders/{underlying}") ,
                        }),
                    }),
                "information":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("{currency}/long_short_ratio") ,
                        MkString("{currency}/volume") ,
                        MkString("{currency}/taker") ,
                        MkString("{currency}/sentiment") ,
                        MkString("{currency}/margin") ,
                        })}),
                "index":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("{instrument_id}/constituents") ,
                        })}),
                }),
            "fees":MkMap(&VarMap{
                "trading":MkMap(&VarMap{
                    "taker":MkNumber(0.0015) ,
                    "maker":MkNumber(0.0010) ,
                    }),
                "spot":MkMap(&VarMap{
                    "taker":MkNumber(0.0015) ,
                    "maker":MkNumber(0.0010) ,
                    }),
                "futures":MkMap(&VarMap{
                    "taker":MkNumber(0.0005) ,
                    "maker":MkNumber(0.0002) ,
                    }),
                "swap":MkMap(&VarMap{
                    "taker":MkNumber(0.00075) ,
                    "maker":MkNumber(0.00020) ,
                    }),
                }),
            "requiredCredentials":MkMap(&VarMap{
                "apiKey":MkBool(true) ,
                "secret":MkBool(true) ,
                "password":MkBool(true) ,
                }),
            "exceptions":MkMap(&VarMap{
                "exact":MkMap(&VarMap{
                    "1":ExchangeError ,
                    "failure to get a peer from the ring-balancer":ExchangeNotAvailable ,
                    "Server is busy, please try again.":ExchangeNotAvailable ,
                    "An unexpected error occurred":ExchangeError ,
                    "System error":ExchangeError ,
                    "4010":PermissionDenied ,
                    "4001":ExchangeError ,
                    "4002":ExchangeError ,
                    "30001":AuthenticationError ,
                    "30002":AuthenticationError ,
                    "30003":AuthenticationError ,
                    "30004":AuthenticationError ,
                    "30005":InvalidNonce ,
                    "30006":AuthenticationError ,
                    "30007":BadRequest ,
                    "30008":RequestTimeout ,
                    "30009":ExchangeError ,
                    "30010":AuthenticationError ,
                    "30011":PermissionDenied ,
                    "30012":AuthenticationError ,
                    "30013":AuthenticationError ,
                    "30014":DDoSProtection ,
                    "30015":AuthenticationError ,
                    "30016":ExchangeError ,
                    "30017":ExchangeError ,
                    "30018":ExchangeError ,
                    "30019":ExchangeNotAvailable ,
                    "30020":BadRequest ,
                    "30021":BadRequest ,
                    "30022":PermissionDenied ,
                    "30023":BadRequest ,
                    "30024":BadSymbol ,
                    "30025":BadRequest ,
                    "30026":DDoSProtection ,
                    "30027":AuthenticationError ,
                    "30028":PermissionDenied ,
                    "30029":AccountSuspended ,
                    "30030":ExchangeNotAvailable ,
                    "30031":BadRequest ,
                    "30032":BadSymbol ,
                    "30033":BadRequest ,
                    "30034":ExchangeError ,
                    "30035":ExchangeError ,
                    "30036":ExchangeError ,
                    "30037":ExchangeNotAvailable ,
                    "30038":OnMaintenance ,
                    "30044":RequestTimeout ,
                    "32001":AccountSuspended ,
                    "32002":PermissionDenied ,
                    "32003":CancelPending ,
                    "32004":ExchangeError ,
                    "32005":InvalidOrder ,
                    "32006":InvalidOrder ,
                    "32007":InvalidOrder ,
                    "32008":InvalidOrder ,
                    "32009":InvalidOrder ,
                    "32010":ExchangeError ,
                    "32011":ExchangeError ,
                    "32012":ExchangeError ,
                    "32013":ExchangeError ,
                    "32014":ExchangeError ,
                    "32015":ExchangeError ,
                    "32016":ExchangeError ,
                    "32017":ExchangeError ,
                    "32018":ExchangeError ,
                    "32019":ExchangeError ,
                    "32020":ExchangeError ,
                    "32021":ExchangeError ,
                    "32022":ExchangeError ,
                    "32023":ExchangeError ,
                    "32024":ExchangeError ,
                    "32025":ExchangeError ,
                    "32026":ExchangeError ,
                    "32027":ExchangeError ,
                    "32028":ExchangeError ,
                    "32029":ExchangeError ,
                    "32030":InvalidOrder ,
                    "32031":ArgumentsRequired ,
                    "32038":AuthenticationError ,
                    "32040":ExchangeError ,
                    "32044":ExchangeError ,
                    "32045":ExchangeError ,
                    "32046":ExchangeError ,
                    "32047":ExchangeError ,
                    "32048":InvalidOrder ,
                    "32049":ExchangeError ,
                    "32050":InvalidOrder ,
                    "32051":InvalidOrder ,
                    "32052":ExchangeError ,
                    "32053":ExchangeError ,
                    "32057":ExchangeError ,
                    "32054":ExchangeError ,
                    "32055":InvalidOrder ,
                    "32056":ExchangeError ,
                    "32058":ExchangeError ,
                    "32059":InvalidOrder ,
                    "32060":InvalidOrder ,
                    "32061":InvalidOrder ,
                    "32062":InvalidOrder ,
                    "32063":InvalidOrder ,
                    "32064":ExchangeError ,
                    "32065":ExchangeError ,
                    "32066":ExchangeError ,
                    "32067":ExchangeError ,
                    "32068":ExchangeError ,
                    "32069":ExchangeError ,
                    "32070":ExchangeError ,
                    "32071":ExchangeError ,
                    "32072":ExchangeError ,
                    "32073":ExchangeError ,
                    "32074":ExchangeError ,
                    "32075":ExchangeError ,
                    "32076":ExchangeError ,
                    "32077":ExchangeError ,
                    "32078":ExchangeError ,
                    "32079":ExchangeError ,
                    "32080":ExchangeError ,
                    "32083":ExchangeError ,
                    "33001":PermissionDenied ,
                    "33002":AccountSuspended ,
                    "33003":InsufficientFunds ,
                    "33004":ExchangeError ,
                    "33005":ExchangeError ,
                    "33006":ExchangeError ,
                    "33007":ExchangeError ,
                    "33008":InsufficientFunds ,
                    "33009":ExchangeError ,
                    "33010":ExchangeError ,
                    "33011":ExchangeError ,
                    "33012":ExchangeError ,
                    "33013":InvalidOrder ,
                    "33014":OrderNotFound ,
                    "33015":InvalidOrder ,
                    "33016":ExchangeError ,
                    "33017":InsufficientFunds ,
                    "33018":ExchangeError ,
                    "33020":ExchangeError ,
                    "33021":BadRequest ,
                    "33022":InvalidOrder ,
                    "33023":ExchangeError ,
                    "33024":InvalidOrder ,
                    "33025":InvalidOrder ,
                    "33026":ExchangeError ,
                    "33027":InvalidOrder ,
                    "33028":InvalidOrder ,
                    "33029":InvalidOrder ,
                    "33034":ExchangeError ,
                    "33035":ExchangeError ,
                    "33036":ExchangeError ,
                    "33037":ExchangeError ,
                    "33038":ExchangeError ,
                    "33039":ExchangeError ,
                    "33040":ExchangeError ,
                    "33041":ExchangeError ,
                    "33042":ExchangeError ,
                    "33043":ExchangeError ,
                    "33044":ExchangeError ,
                    "33045":ExchangeError ,
                    "33046":ExchangeError ,
                    "33047":ExchangeError ,
                    "33048":ExchangeError ,
                    "33049":ExchangeError ,
                    "33050":ExchangeError ,
                    "33051":ExchangeError ,
                    "33059":BadRequest ,
                    "33060":BadRequest ,
                    "33061":ExchangeError ,
                    "33062":ExchangeError ,
                    "33063":ExchangeError ,
                    "33064":ExchangeError ,
                    "33065":ExchangeError ,
                    "33085":InvalidOrder ,
                    "21009":ExchangeError ,
                    "34001":PermissionDenied ,
                    "34002":InvalidAddress ,
                    "34003":ExchangeError ,
                    "34004":ExchangeError ,
                    "34005":ExchangeError ,
                    "34006":ExchangeError ,
                    "34007":ExchangeError ,
                    "34008":InsufficientFunds ,
                    "34009":ExchangeError ,
                    "34010":ExchangeError ,
                    "34011":ExchangeError ,
                    "34012":ExchangeError ,
                    "34013":ExchangeError ,
                    "34014":ExchangeError ,
                    "34015":ExchangeError ,
                    "34016":PermissionDenied ,
                    "34017":AccountSuspended ,
                    "34018":AuthenticationError ,
                    "34019":PermissionDenied ,
                    "34020":PermissionDenied ,
                    "34021":InvalidAddress ,
                    "34022":ExchangeError ,
                    "34023":PermissionDenied ,
                    "34026":RateLimitExceeded ,
                    "34036":ExchangeError ,
                    "34037":ExchangeError ,
                    "34038":ExchangeError ,
                    "34039":ExchangeError ,
                    "35001":ExchangeError ,
                    "35002":ExchangeError ,
                    "35003":ExchangeError ,
                    "35004":ExchangeError ,
                    "35005":AuthenticationError ,
                    "35008":InvalidOrder ,
                    "35010":InvalidOrder ,
                    "35012":InvalidOrder ,
                    "35014":InvalidOrder ,
                    "35015":InvalidOrder ,
                    "35017":ExchangeError ,
                    "35019":InvalidOrder ,
                    "35020":InvalidOrder ,
                    "35021":InvalidOrder ,
                    "35022":BadRequest ,
                    "35024":BadRequest ,
                    "35025":InsufficientFunds ,
                    "35026":BadRequest ,
                    "35029":OrderNotFound ,
                    "35030":InvalidOrder ,
                    "35031":InvalidOrder ,
                    "35032":ExchangeError ,
                    "35037":ExchangeError ,
                    "35039":InsufficientFunds ,
                    "35040":InvalidOrder ,
                    "35044":ExchangeError ,
                    "35046":InsufficientFunds ,
                    "35047":InsufficientFunds ,
                    "35048":ExchangeError ,
                    "35049":InvalidOrder ,
                    "35050":InvalidOrder ,
                    "35052":InsufficientFunds ,
                    "35053":ExchangeError ,
                    "35055":InsufficientFunds ,
                    "35057":ExchangeError ,
                    "35058":ExchangeError ,
                    "35059":BadRequest ,
                    "35060":BadRequest ,
                    "35061":BadRequest ,
                    "35062":InvalidOrder ,
                    "35063":InvalidOrder ,
                    "35064":InvalidOrder ,
                    "35066":InvalidOrder ,
                    "35067":InvalidOrder ,
                    "35068":InvalidOrder ,
                    "35069":InvalidOrder ,
                    "35070":InvalidOrder ,
                    "35071":InvalidOrder ,
                    "35072":InvalidOrder ,
                    "35073":InvalidOrder ,
                    "35074":InvalidOrder ,
                    "35075":InvalidOrder ,
                    "35076":InvalidOrder ,
                    "35077":InvalidOrder ,
                    "35078":InvalidOrder ,
                    "35079":InvalidOrder ,
                    "35080":InvalidOrder ,
                    "35081":InvalidOrder ,
                    "35082":InvalidOrder ,
                    "35083":InvalidOrder ,
                    "35084":InvalidOrder ,
                    "35085":InvalidOrder ,
                    "35086":InvalidOrder ,
                    "35087":InvalidOrder ,
                    "35088":InvalidOrder ,
                    "35089":InvalidOrder ,
                    "35090":ExchangeError ,
                    "35091":ExchangeError ,
                    "35092":ExchangeError ,
                    "35093":ExchangeError ,
                    "35094":ExchangeError ,
                    "35095":BadRequest ,
                    "35096":ExchangeError ,
                    "35097":ExchangeError ,
                    "35098":ExchangeError ,
                    "35099":ExchangeError ,
                    "35102":RateLimitExceeded ,
                    "36001":BadRequest ,
                    "36002":BadRequest ,
                    "36005":ExchangeError ,
                    "36101":AuthenticationError ,
                    "36102":PermissionDenied ,
                    "36103":PermissionDenied ,
                    "36104":PermissionDenied ,
                    "36105":PermissionDenied ,
                    "36106":PermissionDenied ,
                    "36107":PermissionDenied ,
                    "36108":InsufficientFunds ,
                    "36109":PermissionDenied ,
                    "36201":PermissionDenied ,
                    "36202":PermissionDenied ,
                    "36203":InvalidOrder ,
                    "36204":ExchangeError ,
                    "36205":BadRequest ,
                    "36206":BadRequest ,
                    "36207":InvalidOrder ,
                    "36208":InvalidOrder ,
                    "36209":InvalidOrder ,
                    "36210":InvalidOrder ,
                    "36211":InvalidOrder ,
                    "36212":InvalidOrder ,
                    "36213":InvalidOrder ,
                    "36214":ExchangeError ,
                    "36216":OrderNotFound ,
                    "36217":InvalidOrder ,
                    "36218":InvalidOrder ,
                    "36219":InvalidOrder ,
                    "36220":InvalidOrder ,
                    "36221":InvalidOrder ,
                    "36222":InvalidOrder ,
                    "36223":InvalidOrder ,
                    "36224":InvalidOrder ,
                    "36225":InvalidOrder ,
                    "36226":InvalidOrder ,
                    "36227":InvalidOrder ,
                    "36228":InvalidOrder ,
                    "36229":InvalidOrder ,
                    "36230":InvalidOrder ,
                    }),
                "broad":MkMap(&VarMap{}),
                }),
            "precisionMode":TICK_SIZE ,
            "options":MkMap(&VarMap{
                "fetchOHLCV":MkMap(&VarMap{"type":MkString("Candles") }),
                "createMarketBuyOrderRequiresPrice":MkBool(true) ,
                "fetchMarkets":MkArray(&VarArray{
                    MkString("spot") ,
                    MkString("futures") ,
                    MkString("swap") ,
                    MkString("option") ,
                    }),
                "defaultType":MkString("spot") ,
                "auth":MkMap(&VarMap{
                    "time":MkString("public") ,
                    "currencies":MkString("private") ,
                    "instruments":MkString("public") ,
                    "rate":MkString("public") ,
                    "{instrument_id}/constituents":MkString("public") ,
                    }),
                }),
            "commonCurrencies":MkMap(&VarMap{
                "AE":MkString("AET") ,
                "BOX":MkString("DefiBox") ,
                "HOT":MkString("Hydro Protocol") ,
                "HSR":MkString("HC") ,
                "MAG":MkString("Maggie") ,
                "SBTC":MkString("Super Bitcoin") ,
                "YOYO":MkString("YOYOW") ,
                "WIN":MkString("WinToken") ,
                }),
            }));
}

func (this *Okex3) FetchTime(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("generalGetTime"), params ); _ = response;
  return this.Parse8601(this.SafeString(response , MkString("iso") ));
}

func (this *Okex3) FetchMarkets(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  types := this.SafeValue(*this.At(MkString("options")) , MkString("fetchMarkets") ); _ = types;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , types.Length )); OpIncr(& i ){
    markets := this.FetchMarketsByType(*(types ).At(i ), params ); _ = markets;
    result = this.ArrayConcat(result , markets );
  }
  return result ;
}

func (this *Okex3) ParseMarkets(goArgs ...*Variant) *Variant{
  markets := GoGetArg(goArgs, 0, MkUndefined()); _ = markets;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , markets.Length )); OpIncr(& i ){
    result.Push(this.ParseMarket(*(markets ).At(i )));
  }
  return result ;
}

func (this *Okex3) ParseMarket(goArgs ...*Variant) *Variant{
  market := GoGetArg(goArgs, 0, MkUndefined()); _ = market;
  id := this.SafeString(market , MkString("instrument_id") ); _ = id;
  marketType := MkString("spot") ; _ = marketType;
  spot := OpCopy(MkBool(true) ); _ = spot;
  future := OpCopy(MkBool(false) ); _ = future;
  swap := OpCopy(MkBool(false) ); _ = swap;
  option := OpCopy(MkBool(false) ); _ = option;
  baseId := this.SafeString(market , MkString("base_currency") ); _ = baseId;
  quoteId := this.SafeString(market , MkString("quote_currency") ); _ = quoteId;
  contractVal := this.SafeNumber(market , MkString("contract_val") ); _ = contractVal;
  if IsTrue(OpNotEq2(contractVal , MkUndefined() )) {
    if IsTrue(OpHasMember(MkString("option_type") , market )) {
      marketType = MkString("option") ;
      spot = OpCopy(MkBool(false) );
      option = OpCopy(MkBool(true) );
      underlying := this.SafeString(market , MkString("underlying") ); _ = underlying;
      parts := underlying.Split(MkString("-") ); _ = parts;
      baseId = this.SafeString(parts , MkInteger(0) );
      quoteId = this.SafeString(parts , MkInteger(1) );
    } else {
      marketType = MkString("swap") ;
      spot = OpCopy(MkBool(false) );
      swap = OpCopy(MkBool(true) );
      futuresAlias := this.SafeString(market , MkString("alias") ); _ = futuresAlias;
      if IsTrue(OpNotEq2(futuresAlias , MkUndefined() )) {
        swap = OpCopy(MkBool(false) );
        future = OpCopy(MkBool(true) );
        marketType = MkString("futures") ;
        baseId = this.SafeString(market , MkString("underlying_index") );
      }
    }
  }
  base := this.SafeCurrencyCode(baseId ); _ = base;
  quote := this.SafeCurrencyCode(quoteId ); _ = quote;
  symbol := OpTrinary(spot , OpAdd(base , OpAdd(MkString("/") , quote )), id ); _ = symbol;
  lotSize := this.SafeNumber2(market , MkString("lot_size") , MkString("trade_increment") ); _ = lotSize;
  minPrice := this.SafeString(market , MkString("tick_size") ); _ = minPrice;
  precision := MkMap(&VarMap{
        "amount":this.SafeNumber(market , MkString("size_increment") , lotSize ),
        "price":this.ParseNumber(minPrice ),
        }); _ = precision;
  minAmountString := this.SafeString2(market , MkString("min_size") , MkString("base_min_size") ); _ = minAmountString;
  minAmount := this.ParseNumber(minAmountString ); _ = minAmount;
  minCost := OpCopy(MkUndefined() ); _ = minCost;
  if IsTrue(OpAnd(OpNotEq2(minAmount , MkUndefined() ), OpNotEq2(minPrice , MkUndefined() ))) {
    minCost = this.ParseNumber(Precise.StringMul(minPrice , minAmountString ));
  }
  active := OpCopy(MkBool(true) ); _ = active;
  fees := this.SafeValue2(*this.At(MkString("fees")) , marketType , MkString("trading") , MkMap(&VarMap{})); _ = fees;
  return this.Extend(fees , MkMap(&VarMap{
            "id":id ,
            "symbol":symbol ,
            "base":base ,
            "quote":quote ,
            "baseId":baseId ,
            "quoteId":quoteId ,
            "info":market ,
            "type":marketType ,
            "spot":spot ,
            "futures":future ,
            "swap":swap ,
            "option":option ,
            "active":active ,
            "precision":precision ,
            "limits":MkMap(&VarMap{
                "amount":MkMap(&VarMap{
                    "min":minAmount ,
                    "max":MkUndefined() ,
                    }),
                "price":MkMap(&VarMap{
                    "min":*(precision ).At(MkString("price") ),
                    "max":MkUndefined() ,
                    }),
                "cost":MkMap(&VarMap{
                    "min":minCost ,
                    "max":MkUndefined() ,
                    }),
                }),
            }));
}

func (this *Okex3) FetchMarketsByType(goArgs ...*Variant) *Variant{
  type_ := GoGetArg(goArgs, 0, MkUndefined()); _ = type_;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(type_ , MkString("option") )) {
    underlying := this.Call(MkString("optionGetUnderlying"), params ); _ = underlying;
    result := MkArray(&VarArray{}); _ = result;
    for i := MkInteger(0) ; IsTrue(OpLw(i , underlying.Length )); OpIncr(& i ){
      response := this.Call(MkString("optionGetInstrumentsUnderlying"), MkMap(&VarMap{"underlying":*(underlying ).At(i )})); _ = response;
      result = this.ArrayConcat(result , response );
    }
    return this.ParseMarkets(result );
  } else {
    if IsTrue(OpOr(OpEq2(type_ , MkString("spot") ), OpOr(OpEq2(type_ , MkString("futures") ), OpEq2(type_ , MkString("swap") )))) {
      method := OpAdd(type_ , MkString("GetInstruments") ); _ = method;
      response := this.Call(method , params ); _ = response;
      return this.ParseMarkets(response );
    } else {
      panic(NewNotSupported(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" fetchMarketsByType does not support market type ") , type_ ))));
    }
  }
  return MkUndefined()
}

func (this *Okex3) FetchCurrencies(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("accountGetCurrencies"), params ); _ = response;
  result := MkMap(&VarMap{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , response.Length )); OpIncr(& i ){
    currency := *(response ).At(i ); _ = currency;
    id := this.SafeString(currency , MkString("currency") ); _ = id;
    code := this.SafeCurrencyCode(id ); _ = code;
    precision := MkNumber(0.00000001) ; _ = precision;
    name := this.SafeString(currency , MkString("name") ); _ = name;
    canDeposit := this.SafeInteger(currency , MkString("can_deposit") ); _ = canDeposit;
    canWithdraw := this.SafeInteger(currency , MkString("can_withdraw") ); _ = canWithdraw;
    active := OpTrinary(OpAnd(canDeposit , canWithdraw ), MkBool(true) , MkBool(false) ); _ = active;
    *(result ).At(code )= MkMap(&VarMap{
        "id":id ,
        "code":code ,
        "info":currency ,
        "type":MkUndefined() ,
        "name":name ,
        "active":active ,
        "fee":MkUndefined() ,
        "precision":precision ,
        "limits":MkMap(&VarMap{
            "amount":MkMap(&VarMap{
                "min":MkUndefined() ,
                "max":MkUndefined() ,
                }),
            "withdraw":MkMap(&VarMap{
                "min":this.SafeNumber(currency , MkString("min_withdrawal") ),
                "max":MkUndefined() ,
                }),
            }),
        });
  }
  return result ;
}

func (this *Okex3) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  method := OpAdd(*(market ).At(MkString("type") ), MkString("GetInstrumentsInstrumentId") ); _ = method;
  OpTrinary(OpAddAssign(& method , OpEq2(*(market ).At(MkString("type") ), MkString("swap") )), MkString("Depth") , MkString("Book") );
  request := MkMap(&VarMap{"instrument_id":*(market ).At(MkString("id") )}); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("size") )= OpCopy(limit );
  }
  response := this.Call(method , this.Extend(request , params )); _ = response;
  timestamp := this.Parse8601(this.SafeString2(response , MkString("timestamp") , MkString("time") )); _ = timestamp;
  return this.ParseOrderBook(response , symbol , timestamp );
}

func (this *Okex3) ParseTicker(goArgs ...*Variant) *Variant{
  ticker := GoGetArg(goArgs, 0, MkUndefined()); _ = ticker;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.Parse8601(this.SafeString(ticker , MkString("timestamp") )); _ = timestamp;
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  marketId := this.SafeString(ticker , MkString("instrument_id") ); _ = marketId;
  if IsTrue(OpHasMember(marketId , *this.At(MkString("markets_by_id")) )) {
    market = *(*this.At(MkString("markets_by_id")) ).At(marketId );
    symbol = *(market ).At(MkString("symbol") );
  } else {
    if IsTrue(OpNotEq2(marketId , MkUndefined() )) {
      parts := marketId.Split(MkString("-") ); _ = parts;
      numParts := OpCopy(parts.Length ); _ = numParts;
      if IsTrue(OpEq2(numParts , MkInteger(2) )) {
        baseId := MkUndefined(); _ = baseId;
        quoteId := MkUndefined(); _ = quoteId;
        ArrayUnpack(parts , &baseId, &quoteId);
        base := this.SafeCurrencyCode(baseId ); _ = base;
        quote := this.SafeCurrencyCode(quoteId ); _ = quote;
        symbol = OpAdd(base , OpAdd(MkString("/") , quote ));
      } else {
        symbol = OpCopy(marketId );
      }
    }
  }
  if IsTrue(OpAnd(OpEq2(symbol , MkUndefined() ), OpNotEq2(market , MkUndefined() ))) {
    symbol = *(market ).At(MkString("symbol") );
  }
  last := this.SafeNumber(ticker , MkString("last") ); _ = last;
  open := this.SafeNumber(ticker , MkString("open_24h") ); _ = open;
  return MkMap(&VarMap{
        "symbol":symbol ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "high":this.SafeNumber(ticker , MkString("high_24h") ),
        "low":this.SafeNumber(ticker , MkString("low_24h") ),
        "bid":this.SafeNumber(ticker , MkString("best_bid") ),
        "bidVolume":this.SafeNumber(ticker , MkString("best_bid_size") ),
        "ask":this.SafeNumber(ticker , MkString("best_ask") ),
        "askVolume":this.SafeNumber(ticker , MkString("best_ask_size") ),
        "vwap":MkUndefined() ,
        "open":open ,
        "close":last ,
        "last":last ,
        "previousClose":MkUndefined() ,
        "change":MkUndefined() ,
        "percentage":MkUndefined() ,
        "average":MkUndefined() ,
        "baseVolume":this.SafeNumber(ticker , MkString("base_volume_24h") ),
        "quoteVolume":this.SafeNumber(ticker , MkString("quote_volume_24h") ),
        "info":ticker ,
        });
}

func (this *Okex3) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  method := OpAdd(*(market ).At(MkString("type") ), MkString("GetInstrumentsInstrumentIdTicker") ); _ = method;
  request := MkMap(&VarMap{"instrument_id":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(method , this.Extend(request , params )); _ = response;
  return this.ParseTicker(response );
}

func (this *Okex3) FetchTickersByType(goArgs ...*Variant) *Variant{
  type_ := GoGetArg(goArgs, 0, MkUndefined()); _ = type_;
  symbols := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbols;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  method := OpAdd(type_ , MkString("GetInstrumentsTicker") ); _ = method;
  response := this.Call(method , params ); _ = response;
  result := MkMap(&VarMap{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , response.Length )); OpIncr(& i ){
    ticker := this.ParseTicker(*(response ).At(i )); _ = ticker;
    symbol := *(ticker ).At(MkString("symbol") ); _ = symbol;
    *(result ).At(symbol )= OpCopy(ticker );
  }
  return this.FilterByArray(result , MkString("symbol") , symbols );
}

func (this *Okex3) FetchTickers(goArgs ...*Variant) *Variant{
  symbols := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbols;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  defaultType := this.SafeString2(*this.At(MkString("options")) , MkString("fetchTickers") , MkString("defaultType") ); _ = defaultType;
  type_ := this.SafeString(params , MkString("type") , defaultType ); _ = type_;
  return this.FetchTickersByType(type_ , symbols , this.Omit(params , MkString("type") ));
}

func (this *Okex3) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  marketId := this.SafeString(trade , MkString("instrument_id") ); _ = marketId;
  base := OpCopy(MkUndefined() ); _ = base;
  quote := OpCopy(MkUndefined() ); _ = quote;
  if IsTrue(OpHasMember(marketId , *this.At(MkString("markets_by_id")) )) {
    market = *(*this.At(MkString("markets_by_id")) ).At(marketId );
    symbol = *(market ).At(MkString("symbol") );
    base = *(market ).At(MkString("base") );
    quote = *(market ).At(MkString("quote") );
  } else {
    if IsTrue(OpNotEq2(marketId , MkUndefined() )) {
      parts := marketId.Split(MkString("-") ); _ = parts;
      numParts := OpCopy(parts.Length ); _ = numParts;
      if IsTrue(OpEq2(numParts , MkInteger(2) )) {
        baseId := MkUndefined(); _ = baseId;
        quoteId := MkUndefined(); _ = quoteId;
        ArrayUnpack(parts , &baseId, &quoteId);
        base = this.SafeCurrencyCode(baseId );
        quote = this.SafeCurrencyCode(quoteId );
        symbol = OpAdd(base , OpAdd(MkString("/") , quote ));
      } else {
        symbol = OpCopy(marketId );
      }
    }
  }
  if IsTrue(OpAnd(OpEq2(symbol , MkUndefined() ), OpNotEq2(market , MkUndefined() ))) {
    symbol = *(market ).At(MkString("symbol") );
    base = *(market ).At(MkString("base") );
    quote = *(market ).At(MkString("quote") );
  }
  timestamp := this.Parse8601(this.SafeString2(trade , MkString("timestamp") , MkString("created_at") )); _ = timestamp;
  priceString := this.SafeString(trade , MkString("price") ); _ = priceString;
  amountString := this.SafeString2(trade , MkString("size") , MkString("qty") ); _ = amountString;
  amountString = this.SafeString(trade , MkString("order_qty") , amountString );
  price := this.ParseNumber(priceString ); _ = price;
  amount := this.ParseNumber(amountString ); _ = amount;
  cost := this.ParseNumber(Precise.StringMul(priceString , amountString )); _ = cost;
  takerOrMaker := this.SafeString2(trade , MkString("exec_type") , MkString("liquidity") ); _ = takerOrMaker;
  if IsTrue(OpEq2(takerOrMaker , MkString("M") )) {
    takerOrMaker = MkString("maker") ;
  } else {
    if IsTrue(OpEq2(takerOrMaker , MkString("T") )) {
      takerOrMaker = MkString("taker") ;
    }
  }
  side := this.SafeString(trade , MkString("side") ); _ = side;
  feeCost := this.SafeNumber(trade , MkString("fee") ); _ = feeCost;
  fee := OpCopy(MkUndefined() ); _ = fee;
  if IsTrue(OpNotEq2(feeCost , MkUndefined() )) {
    feeCurrency := OpTrinary(OpEq2(side , MkString("buy") ), base , quote ); _ = feeCurrency;
    fee = MkMap(&VarMap{
        "cost":OpNeg(feeCost ),
        "currency":feeCurrency ,
        });
  }
  orderId := this.SafeString(trade , MkString("order_id") ); _ = orderId;
  return MkMap(&VarMap{
        "info":trade ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "symbol":symbol ,
        "id":this.SafeString2(trade , MkString("trade_id") , MkString("ledger_id") ),
        "order":orderId ,
        "type":MkUndefined() ,
        "takerOrMaker":takerOrMaker ,
        "side":side ,
        "price":price ,
        "amount":amount ,
        "cost":cost ,
        "fee":fee ,
        });
}

func (this *Okex3) FetchTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  method := OpAdd(*(market ).At(MkString("type") ), MkString("GetInstrumentsInstrumentIdTrades") ); _ = method;
  if IsTrue(OpOr(OpEq2(limit , MkUndefined() ), OpGt(limit , MkInteger(100) ))) {
    limit = MkInteger(100) ;
  }
  request := MkMap(&VarMap{
        "instrument_id":*(market ).At(MkString("id") ),
        "limit":limit ,
        }); _ = request;
  response := this.Call(method , this.Extend(request , params )); _ = response;
  return this.ParseTrades(response , market , since , limit );
}

func (this *Okex3) ParseOHLCV(goArgs ...*Variant) *Variant{
  ohlcv := GoGetArg(goArgs, 0, MkUndefined()); _ = ohlcv;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  if IsTrue(GoIsArray(ohlcv )) {
    numElements := OpCopy(ohlcv.Length ); _ = numElements;
    volumeIndex := OpTrinary(OpGt(numElements , MkInteger(6) ), MkInteger(6) , MkInteger(5) ); _ = volumeIndex;
    timestamp := this.SafeValue(ohlcv , MkInteger(0) ); _ = timestamp;
    if IsTrue(OpEq2(OpType(timestamp ), MkString("string") )) {
      timestamp = this.Parse8601(timestamp );
    }
    return MkArray(&VarArray{
          timestamp ,
          this.SafeNumber(ohlcv , MkInteger(1) ),
          this.SafeNumber(ohlcv , MkInteger(2) ),
          this.SafeNumber(ohlcv , MkInteger(3) ),
          this.SafeNumber(ohlcv , MkInteger(4) ),
          this.SafeNumber(ohlcv , volumeIndex ),
          });
  } else {
    return MkArray(&VarArray{
          this.Parse8601(this.SafeString(ohlcv , MkString("time") )),
          this.SafeNumber(ohlcv , MkString("open") ),
          this.SafeNumber(ohlcv , MkString("high") ),
          this.SafeNumber(ohlcv , MkString("low") ),
          this.SafeNumber(ohlcv , MkString("close") ),
          this.SafeNumber(ohlcv , MkString("volume") ),
          });
  }
  return MkUndefined()
}

func (this *Okex3) FetchOHLCV(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  timeframe := GoGetArg(goArgs, 1, MkString("1m") ); _ = timeframe;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  duration := this.ParseTimeframe(timeframe ); _ = duration;
  request := MkMap(&VarMap{
        "instrument_id":*(market ).At(MkString("id") ),
        "granularity":*(*this.At(MkString("timeframes")) ).At(timeframe ),
        }); _ = request;
  options := this.SafeValue(*this.At(MkString("options")) , MkString("fetchOHLCV") , MkMap(&VarMap{})); _ = options;
  defaultType := this.SafeString(options , MkString("type") , MkString("Candles") ); _ = defaultType;
  type_ := this.SafeString(params , MkString("type") , defaultType ); _ = type_;
  params = this.Omit(params , MkString("type") );
  method := OpAdd(*(market ).At(MkString("type") ), OpAdd(MkString("GetInstrumentsInstrumentId") , type_ )); _ = method;
  if IsTrue(OpEq2(type_ , MkString("Candles") )) {
    if IsTrue(OpNotEq2(since , MkUndefined() )) {
      if IsTrue(OpNotEq2(limit , MkUndefined() )) {
        *(request ).At(MkString("end") )= this.Iso8601(this.Sum(since , OpMulti(limit , OpMulti(duration , MkInteger(1000) ))));
      }
      *(request ).At(MkString("start") )= this.Iso8601(since );
    } else {
      if IsTrue(OpNotEq2(limit , MkUndefined() )) {
        now := this.Milliseconds(); _ = now;
        *(request ).At(MkString("start") )= this.Iso8601(OpSub(now , OpMulti(limit , OpMulti(duration , MkInteger(1000) ))));
        *(request ).At(MkString("end") )= this.Iso8601(now );
      }
    }
  } else {
    if IsTrue(OpEq2(type_ , MkString("HistoryCandles") )) {
      if IsTrue(*(market ).At(MkString("option") )) {
        panic(NewNotSupported(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" fetchOHLCV does not have ") , OpAdd(type_ , OpAdd(MkString(" for ") , OpAdd(*(market ).At(MkString("type") ), MkString(" markets") )))))));
      }
      if IsTrue(OpNotEq2(since , MkUndefined() )) {
        if IsTrue(OpEq2(limit , MkUndefined() )) {
          limit = MkInteger(300) ;
        }
        *(request ).At(MkString("start") )= this.Iso8601(this.Sum(since , OpMulti(limit , OpMulti(duration , MkInteger(1000) ))));
        *(request ).At(MkString("end") )= this.Iso8601(since );
      } else {
        if IsTrue(OpNotEq2(limit , MkUndefined() )) {
          now := this.Milliseconds(); _ = now;
          *(request ).At(MkString("end") )= this.Iso8601(OpSub(now , OpMulti(limit , OpMulti(duration , MkInteger(1000) ))));
          *(request ).At(MkString("start") )= this.Iso8601(now );
        }
      }
    }
  }
  response := this.Call(method , this.Extend(request , params )); _ = response;
  return this.ParseOHLCVs(response , market , timeframe , since , limit );
}

func (this *Okex3) ParseAccountBalance(goArgs ...*Variant) *Variant{
  response := GoGetArg(goArgs, 0, MkUndefined()); _ = response;
  result := MkMap(&VarMap{
        "info":response ,
        "timestamp":MkUndefined() ,
        "datetime":MkUndefined() ,
        }); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , response.Length )); OpIncr(& i ){
    balance := *(response ).At(i ); _ = balance;
    currencyId := this.SafeString(balance , MkString("currency") ); _ = currencyId;
    code := this.SafeCurrencyCode(currencyId ); _ = code;
    account := this.Account(); _ = account;
    *(account ).At(MkString("total") )= this.SafeString(balance , MkString("balance") );
    *(account ).At(MkString("used") )= this.SafeString(balance , MkString("hold") );
    *(account ).At(MkString("free") )= this.SafeString(balance , MkString("available") );
    *(result ).At(code )= OpCopy(account );
  }
  return this.ParseBalance(result );
}

func (this *Okex3) ParseMarginBalance(goArgs ...*Variant) *Variant{
  response := GoGetArg(goArgs, 0, MkUndefined()); _ = response;
  result := MkMap(&VarMap{
        "info":response ,
        "timestamp":MkUndefined() ,
        "datetime":MkUndefined() ,
        }); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , response.Length )); OpIncr(& i ){
    balance := *(response ).At(i ); _ = balance;
    marketId := this.SafeString(balance , MkString("instrument_id") ); _ = marketId;
    market := this.SafeValue(*this.At(MkString("markets_by_id")) , marketId ); _ = market;
    symbol := OpCopy(MkUndefined() ); _ = symbol;
    if IsTrue(OpEq2(market , MkUndefined() )) {
      baseId := MkUndefined(); _ = baseId;
      quoteId := MkUndefined(); _ = quoteId;
      ArrayUnpack(marketId.Split(MkString("-") ), &baseId, &quoteId);
      base := this.SafeCurrencyCode(baseId ); _ = base;
      quote := this.SafeCurrencyCode(quoteId ); _ = quote;
      symbol = OpAdd(base , OpAdd(MkString("/") , quote ));
    } else {
      symbol = *(market ).At(MkString("symbol") );
    }
    omittedBalance := this.Omit(balance , MkArray(&VarArray{
              MkString("instrument_id") ,
              MkString("liquidation_price") ,
              MkString("product_id") ,
              MkString("risk_rate") ,
              MkString("margin_ratio") ,
              MkString("maint_margin_ratio") ,
              MkString("tiers") ,
              })); _ = omittedBalance;
    keys := GoGetKeys(omittedBalance ); _ = keys;
    accounts := MkMap(&VarMap{}); _ = accounts;
    for k := MkInteger(0) ; IsTrue(OpLw(k , keys.Length )); OpIncr(& k ){
      key := *(keys ).At(k ); _ = key;
      marketBalance := *(balance ).At(key ); _ = marketBalance;
      if IsTrue(OpGtEq(key.IndexOf(MkString(":") ), MkInteger(0) )) {
        parts := key.Split(MkString(":") ); _ = parts;
        currencyId := *(parts ).At(MkInteger(1) ); _ = currencyId;
        code := this.SafeCurrencyCode(currencyId ); _ = code;
        account := this.Account(); _ = account;
        *(account ).At(MkString("total") )= this.SafeString(marketBalance , MkString("balance") );
        *(account ).At(MkString("used") )= this.SafeString(marketBalance , MkString("hold") );
        *(account ).At(MkString("free") )= this.SafeString(marketBalance , MkString("available") );
        *(accounts ).At(code )= OpCopy(account );
      } else {
        panic(NewNotSupported(OpAdd(*this.At(MkString("id")) , MkString(" margin balance response format has changed!") )));
      }
    }
    *(result ).At(symbol )= this.ParseBalance(accounts );
  }
  return result ;
}

func (this *Okex3) ParseFuturesBalance(goArgs ...*Variant) *Variant{
  response := GoGetArg(goArgs, 0, MkUndefined()); _ = response;
  result := MkMap(&VarMap{
        "info":response ,
        "timestamp":MkUndefined() ,
        "datetime":MkUndefined() ,
        }); _ = result;
  info := this.SafeValue(response , MkString("info") , MkMap(&VarMap{})); _ = info;
  ids := GoGetKeys(info ); _ = ids;
  for i := MkInteger(0) ; IsTrue(OpLw(i , ids.Length )); OpIncr(& i ){
    id := *(ids ).At(i ); _ = id;
    code := this.SafeCurrencyCode(id ); _ = code;
    balance := this.SafeValue(info , id , MkMap(&VarMap{})); _ = balance;
    account := this.Account(); _ = account;
    totalAvailBalance := this.SafeString(balance , MkString("total_avail_balance") ); _ = totalAvailBalance;
    if IsTrue(OpEq2(this.SafeString(balance , MkString("margin_mode") ), MkString("fixed") )) {
      contracts := this.SafeValue(balance , MkString("contracts") , MkArray(&VarArray{})); _ = contracts;
      free := OpCopy(totalAvailBalance ); _ = free;
      for i := MkInteger(0) ; IsTrue(OpLw(i , contracts.Length )); OpIncr(& i ){
        contract := *(contracts ).At(i ); _ = contract;
        fixedBalance := this.SafeString(contract , MkString("fixed_balance") ); _ = fixedBalance;
        realizedPnl := this.SafeString(contract , MkString("realized_pnl") ); _ = realizedPnl;
        marginFrozen := this.SafeString(contract , MkString("margin_frozen") ); _ = marginFrozen;
        marginForUnfilled := this.SafeString(contract , MkString("margin_for_unfilled") ); _ = marginForUnfilled;
        margin := Precise.StringSub(Precise.StringSub(Precise.StringAdd(fixedBalance , realizedPnl ), marginFrozen ), marginForUnfilled ); _ = margin;
        free = Precise.StringAdd(free , margin );
      }
      *(account ).At(MkString("free") )= OpCopy(free );
    } else {
      realizedPnl := this.SafeString(balance , MkString("realized_pnl") ); _ = realizedPnl;
      unrealizedPnl := this.SafeString(balance , MkString("unrealized_pnl") ); _ = unrealizedPnl;
      marginFrozen := this.SafeString(balance , MkString("margin_frozen") ); _ = marginFrozen;
      marginForUnfilled := this.SafeString(balance , MkString("margin_for_unfilled") ); _ = marginForUnfilled;
      positive := Precise.StringAdd(Precise.StringAdd(totalAvailBalance , realizedPnl ), unrealizedPnl ); _ = positive;
      *(account ).At(MkString("free") )= Precise.StringSub(Precise.StringSub(positive , marginFrozen ), marginForUnfilled );
    }
    *(account ).At(MkString("total") )= this.SafeString(balance , MkString("equity") );
    *(result ).At(code )= OpCopy(account );
  }
  return this.ParseBalance(result );
}

func (this *Okex3) ParseSwapBalance(goArgs ...*Variant) *Variant{
  response := GoGetArg(goArgs, 0, MkUndefined()); _ = response;
  result := MkMap(&VarMap{"info":response }); _ = result;
  timestamp := OpCopy(MkUndefined() ); _ = timestamp;
  info := this.SafeValue(response , MkString("info") , MkArray(&VarArray{})); _ = info;
  for i := MkInteger(0) ; IsTrue(OpLw(i , info.Length )); OpIncr(& i ){
    balance := *(info ).At(i ); _ = balance;
    marketId := this.SafeString(balance , MkString("instrument_id") ); _ = marketId;
    symbol := OpCopy(marketId ); _ = symbol;
    if IsTrue(OpHasMember(marketId , *this.At(MkString("markets_by_id")) )) {
      symbol = *(*(*this.At(MkString("markets_by_id")) ).At(marketId )).At(MkString("symbol") );
    }
    balanceTimestamp := this.Parse8601(this.SafeString(balance , MkString("timestamp") )); _ = balanceTimestamp;
    timestamp = OpTrinary(OpEq2(timestamp , MkUndefined() ), balanceTimestamp , Math.Max(timestamp , balanceTimestamp ));
    account := this.Account(); _ = account;
    *(account ).At(MkString("total") )= this.SafeString(balance , MkString("equity") );
    *(account ).At(MkString("free") )= this.SafeString(balance , MkString("total_avail_balance") );
    *(result ).At(symbol )= OpCopy(account );
  }
  *(result ).At(MkString("timestamp") )= OpCopy(timestamp );
  *(result ).At(MkString("datetime") )= this.Iso8601(timestamp );
  return this.ParseBalance(result );
}

func (this *Okex3) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  defaultType := this.SafeString2(*this.At(MkString("options")) , MkString("fetchBalance") , MkString("defaultType") ); _ = defaultType;
  type_ := this.SafeString(params , MkString("type") , defaultType ); _ = type_;
  if IsTrue(OpEq2(type_ , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchBalance() requires a type parameter (one of 'account', 'spot', 'margin', 'futures', 'swap')") )));
  }
  this.LoadMarkets();
  suffix := OpTrinary(OpEq2(type_ , MkString("account") ), MkString("Wallet") , MkString("Accounts") ); _ = suffix;
  method := OpAdd(type_ , OpAdd(MkString("Get") , suffix )); _ = method;
  query := this.Omit(params , MkString("type") ); _ = query;
  response := this.Call(method , query ); _ = response;
  return this.ParseBalanceByType(type_ , response );
}

func (this *Okex3) ParseBalanceByType(goArgs ...*Variant) *Variant{
  type_ := GoGetArg(goArgs, 0, MkUndefined()); _ = type_;
  response := GoGetArg(goArgs, 1, MkUndefined()); _ = response;
  if IsTrue(OpOr(OpEq2(type_ , MkString("account") ), OpEq2(type_ , MkString("spot") ))) {
    return this.ParseAccountBalance(response );
  } else {
    if IsTrue(OpEq2(type_ , MkString("margin") )) {
      return this.ParseMarginBalance(response );
    } else {
      if IsTrue(OpEq2(type_ , MkString("futures") )) {
        return this.ParseFuturesBalance(response );
      } else {
        if IsTrue(OpEq2(type_ , MkString("swap") )) {
          return this.ParseSwapBalance(response );
        }
      }
    }
  }
  panic(NewNotSupported(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" fetchBalance does not support the '") , OpAdd(type_ , MkString("' type (the type must be one of 'account', 'spot', 'margin', 'futures', 'swap')") )))));
  return MkUndefined()
}

func (this *Okex3) CreateOrder(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 2, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 3, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 4, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 5, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"instrument_id":*(market ).At(MkString("id") )}); _ = request;
  clientOrderId := this.SafeString2(params , MkString("client_oid") , MkString("clientOrderId") ); _ = clientOrderId;
  if IsTrue(OpNotEq2(clientOrderId , MkUndefined() )) {
    *(request ).At(MkString("client_oid") )= OpCopy(clientOrderId );
    params = this.Omit(params , MkArray(&VarArray{
            MkString("client_oid") ,
            MkString("clientOrderId") ,
            }));
  }
  method := OpCopy(MkUndefined() ); _ = method;
  if IsTrue(OpOr(*(market ).At(MkString("futures") ), *(market ).At(MkString("swap") ))) {
    size := OpTrinary(*(market ).At(MkString("futures") ), this.NumberToString(amount ), this.AmountToPrecision(symbol , amount )); _ = size;
    request = this.Extend(request , MkMap(&VarMap{
            "type":type_ ,
            "size":size ,
            }));
    orderType := this.SafeString(params , MkString("order_type") ); _ = orderType;
    isMarketOrder := OpOr(OpEq2(type_ , MkString("market") ), OpEq2(orderType , MkString("4") )); _ = isMarketOrder;
    if IsTrue(isMarketOrder ) {
      *(request ).At(MkString("order_type") )= MkString("4") ;
    } else {
      *(request ).At(MkString("price") )= this.PriceToPrecision(symbol , price );
    }
    if IsTrue(*(market ).At(MkString("futures") )) {
      *(request ).At(MkString("leverage") )= MkString("10") ;
    }
    method = OpAdd(*(market ).At(MkString("type") ), MkString("PostOrder") );
  } else {
    marginTrading := this.SafeString(params , MkString("margin_trading") , MkString("1") ); _ = marginTrading;
    request = this.Extend(request , MkMap(&VarMap{
            "side":side ,
            "type":type_ ,
            "margin_trading":marginTrading ,
            }));
    if IsTrue(OpEq2(type_ , MkString("limit") )) {
      *(request ).At(MkString("price") )= this.PriceToPrecision(symbol , price );
      *(request ).At(MkString("size") )= this.AmountToPrecision(symbol , amount );
    } else {
      if IsTrue(OpEq2(type_ , MkString("market") )) {
        if IsTrue(OpEq2(side , MkString("buy") )) {
          notional := this.SafeNumber(params , MkString("notional") ); _ = notional;
          createMarketBuyOrderRequiresPrice := this.SafeValue(*this.At(MkString("options")) , MkString("createMarketBuyOrderRequiresPrice") , MkBool(true) ); _ = createMarketBuyOrderRequiresPrice;
          if IsTrue(createMarketBuyOrderRequiresPrice ) {
            if IsTrue(OpNotEq2(price , MkUndefined() )) {
              if IsTrue(OpEq2(notional , MkUndefined() )) {
                notional = OpMulti(amount , price );
              }
            } else {
              if IsTrue(OpEq2(notional , MkUndefined() )) {
                panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")) , MkString(" createOrder() requires the price argument with market buy orders to calculate total order cost (amount to spend), where cost = amount * price. Supply a price argument to createOrder() call if you want the cost to be calculated for you from price and amount, or, alternatively, add .options['createMarketBuyOrderRequiresPrice'] = false and supply the total cost value in the 'amount' argument or in the 'notional' extra parameter (the exchange-specific behaviour)") )));
              }
            }
          } else {
            notional = OpTrinary(OpEq2(notional , MkUndefined() ), amount , notional );
          }
          precision := *(*(market ).At(MkString("precision") )).At(MkString("price") ); _ = precision;
          *(request ).At(MkString("notional") )= this.DecimalToPrecision(notional , TRUNCATE , precision , *this.At(MkString("precisionMode")) );
        } else {
          *(request ).At(MkString("size") )= this.AmountToPrecision(symbol , amount );
        }
      }
    }
    method = OpTrinary(OpEq2(marginTrading , MkString("2") ), MkString("marginPostOrders") , MkString("spotPostOrders") );
  }
  response := this.Call(method , this.Extend(request , params )); _ = response;
  order := this.ParseOrder(response , market ); _ = order;
  return this.Extend(order , MkMap(&VarMap{
            "type":type_ ,
            "side":side ,
            }));
}

func (this *Okex3) CancelOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" cancelOrder() requires a symbol argument") )));
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  type_ := OpCopy(MkUndefined() ); _ = type_;
  if IsTrue(OpOr(*(market ).At(MkString("futures") ), *(market ).At(MkString("swap") ))) {
    type_ = *(market ).At(MkString("type") );
  } else {
    defaultType := this.SafeString2(*this.At(MkString("options")) , MkString("cancelOrder") , MkString("defaultType") , *(market ).At(MkString("type") )); _ = defaultType;
    type_ = this.SafeString(params , MkString("type") , defaultType );
  }
  if IsTrue(OpEq2(type_ , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" cancelOrder() requires a type parameter (one of 'spot', 'margin', 'futures', 'swap').") )));
  }
  method := OpAdd(type_ , MkString("PostCancelOrder") ); _ = method;
  request := MkMap(&VarMap{"instrument_id":*(market ).At(MkString("id") )}); _ = request;
  if IsTrue(OpOr(*(market ).At(MkString("futures") ), *(market ).At(MkString("swap") ))) {
    OpAddAssign(& method , MkString("InstrumentId") );
  } else {
    OpAddAssign(& method , MkString("s") );
  }
  clientOrderId := this.SafeString2(params , MkString("client_oid") , MkString("clientOrderId") ); _ = clientOrderId;
  if IsTrue(OpNotEq2(clientOrderId , MkUndefined() )) {
    OpAddAssign(& method , MkString("ClientOid") );
    *(request ).At(MkString("client_oid") )= OpCopy(clientOrderId );
  } else {
    OpAddAssign(& method , MkString("OrderId") );
    *(request ).At(MkString("order_id") )= OpCopy(id );
  }
  query := this.Omit(params , MkArray(&VarArray{
            MkString("type") ,
            MkString("client_oid") ,
            MkString("clientOrderId") ,
            })); _ = query;
  response := this.Call(method , this.Extend(request , query )); _ = response;
  result := OpTrinary(OpHasMember(MkString("result") , response ), response , this.SafeValue(response , *(market ).At(MkString("id") ), MkMap(&VarMap{}))); _ = result;
  return this.ParseOrder(result , market );
}

func (this *Okex3) ParseOrderStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "-2":MkString("failed") ,
        "-1":MkString("canceled") ,
        "0":MkString("open") ,
        "1":MkString("open") ,
        "2":MkString("closed") ,
        "3":MkString("open") ,
        "4":MkString("canceled") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Okex3) ParseOrderSide(goArgs ...*Variant) *Variant{
  side := GoGetArg(goArgs, 0, MkUndefined()); _ = side;
  sides := MkMap(&VarMap{
        "1":MkString("buy") ,
        "2":MkString("sell") ,
        "3":MkString("sell") ,
        "4":MkString("buy") ,
        }); _ = sides;
  return this.SafeString(sides , side , side );
}

func (this *Okex3) ParseOrder(goArgs ...*Variant) *Variant{
  order := GoGetArg(goArgs, 0, MkUndefined()); _ = order;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  id := this.SafeString(order , MkString("order_id") ); _ = id;
  timestamp := this.Parse8601(this.SafeString(order , MkString("timestamp") )); _ = timestamp;
  side := this.SafeString(order , MkString("side") ); _ = side;
  type_ := this.SafeString(order , MkString("type") ); _ = type_;
  if IsTrue(OpAnd(OpNotEq2(side , MkString("buy") ), OpNotEq2(side , MkString("sell") ))) {
    side = this.ParseOrderSide(type_ );
  }
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  marketId := this.SafeString(order , MkString("instrument_id") ); _ = marketId;
  if IsTrue(OpHasMember(marketId , *this.At(MkString("markets_by_id")) )) {
    market = *(*this.At(MkString("markets_by_id")) ).At(marketId );
    symbol = *(market ).At(MkString("symbol") );
  } else {
    symbol = OpCopy(marketId );
  }
  if IsTrue(OpNotEq2(market , MkUndefined() )) {
    if IsTrue(OpEq2(symbol , MkUndefined() )) {
      symbol = *(market ).At(MkString("symbol") );
    }
  }
  amount := this.SafeNumber(order , MkString("size") ); _ = amount;
  filled := this.SafeNumber2(order , MkString("filled_size") , MkString("filled_qty") ); _ = filled;
  remaining := OpCopy(MkUndefined() ); _ = remaining;
  if IsTrue(OpNotEq2(amount , MkUndefined() )) {
    if IsTrue(OpNotEq2(filled , MkUndefined() )) {
      amount = Math.Max(amount , filled );
      remaining = Math.Max(MkInteger(0) , OpSub(amount , filled ));
    }
  }
  if IsTrue(OpEq2(type_ , MkString("market") )) {
    remaining = MkInteger(0) ;
  }
  cost := this.SafeNumber2(order , MkString("filled_notional") , MkString("funds") ); _ = cost;
  price := this.SafeNumber(order , MkString("price") ); _ = price;
  average := this.SafeNumber(order , MkString("price_avg") ); _ = average;
  if IsTrue(OpEq2(cost , MkUndefined() )) {
    if IsTrue(OpAnd(OpNotEq2(filled , MkUndefined() ), OpNotEq2(average , MkUndefined() ))) {
      cost = OpMulti(average , filled );
    }
  } else {
    if IsTrue(OpAnd(OpEq2(average , MkUndefined() ), OpAnd(OpNotEq2(filled , MkUndefined() ), OpGt(filled , MkInteger(0) )))) {
      average = OpDiv(cost , filled );
    }
  }
  status := this.ParseOrderStatus(this.SafeString(order , MkString("state") )); _ = status;
  feeCost := this.SafeNumber(order , MkString("fee") ); _ = feeCost;
  fee := OpCopy(MkUndefined() ); _ = fee;
  if IsTrue(OpNotEq2(feeCost , MkUndefined() )) {
    feeCurrency := OpCopy(MkUndefined() ); _ = feeCurrency;
    fee = MkMap(&VarMap{
        "cost":feeCost ,
        "currency":feeCurrency ,
        });
  }
  clientOrderId := this.SafeString(order , MkString("client_oid") ); _ = clientOrderId;
  if IsTrue(OpAnd(OpNotEq2(clientOrderId , MkUndefined() ), OpLw(clientOrderId.Length , MkInteger(1) ))) {
    clientOrderId = OpCopy(MkUndefined() );
  }
  stopPrice := this.SafeNumber(order , MkString("trigger_price") ); _ = stopPrice;
  return MkMap(&VarMap{
        "info":order ,
        "id":id ,
        "clientOrderId":clientOrderId ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "lastTradeTimestamp":MkUndefined() ,
        "symbol":symbol ,
        "type":type_ ,
        "timeInForce":MkUndefined() ,
        "postOnly":MkUndefined() ,
        "side":side ,
        "price":price ,
        "stopPrice":stopPrice ,
        "average":average ,
        "cost":cost ,
        "amount":amount ,
        "filled":filled ,
        "remaining":remaining ,
        "status":status ,
        "fee":fee ,
        "trades":MkUndefined() ,
        });
}

func (this *Okex3) FetchOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchOrder() requires a symbol argument") )));
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  defaultType := this.SafeString2(*this.At(MkString("options")) , MkString("fetchOrder") , MkString("defaultType") , *(market ).At(MkString("type") )); _ = defaultType;
  type_ := this.SafeString(params , MkString("type") , defaultType ); _ = type_;
  if IsTrue(OpEq2(type_ , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchOrder() requires a type parameter (one of 'spot', 'margin', 'futures', 'swap').") )));
  }
  instrumentId := OpTrinary(OpOr(*(market ).At(MkString("futures") ), *(market ).At(MkString("swap") )), MkString("InstrumentId") , MkString("") ); _ = instrumentId;
  method := OpAdd(type_ , OpAdd(MkString("GetOrders") , instrumentId )); _ = method;
  request := MkMap(&VarMap{"instrument_id":*(market ).At(MkString("id") )}); _ = request;
  clientOid := this.SafeString(params , MkString("client_oid") ); _ = clientOid;
  if IsTrue(OpNotEq2(clientOid , MkUndefined() )) {
    OpAddAssign(& method , MkString("ClientOid") );
    *(request ).At(MkString("client_oid") )= OpCopy(clientOid );
  } else {
    OpAddAssign(& method , MkString("OrderId") );
    *(request ).At(MkString("order_id") )= OpCopy(id );
  }
  query := this.Omit(params , MkString("type") ); _ = query;
  response := this.Call(method , this.Extend(request , query )); _ = response;
  return this.ParseOrder(response );
}

func (this *Okex3) FetchOrdersByState(goArgs ...*Variant) *Variant{
  state := GoGetArg(goArgs, 0, MkUndefined()); _ = state;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchOrdersByState() requires a symbol argument") )));
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  type_ := OpCopy(MkUndefined() ); _ = type_;
  if IsTrue(OpOr(*(market ).At(MkString("futures") ), *(market ).At(MkString("swap") ))) {
    type_ = *(market ).At(MkString("type") );
  } else {
    defaultType := this.SafeString2(*this.At(MkString("options")) , MkString("fetchOrder") , MkString("defaultType") , *(market ).At(MkString("type") )); _ = defaultType;
    type_ = this.SafeString(params , MkString("type") , defaultType );
  }
  if IsTrue(OpEq2(type_ , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchOrdersByState() requires a type parameter (one of 'spot', 'margin', 'futures', 'swap').") )));
  }
  request := MkMap(&VarMap{
        "instrument_id":*(market ).At(MkString("id") ),
        "state":state ,
        }); _ = request;
  method := OpAdd(type_ , MkString("GetOrders") ); _ = method;
  if IsTrue(OpOr(*(market ).At(MkString("futures") ), *(market ).At(MkString("swap") ))) {
    OpAddAssign(& method , MkString("InstrumentId") );
  }
  query := this.Omit(params , MkString("type") ); _ = query;
  response := this.Call(method , this.Extend(request , query )); _ = response;
  orders := OpCopy(MkUndefined() ); _ = orders;
  if IsTrue(OpOr(*(market ).At(MkString("swap") ), *(market ).At(MkString("futures") ))) {
    orders = this.SafeValue(response , MkString("order_info") , MkArray(&VarArray{}));
  } else {
    orders = OpCopy(response );
    responseLength := OpCopy(response.Length ); _ = responseLength;
    if IsTrue(OpLw(responseLength , MkInteger(1) )) {
      return MkArray(&VarArray{});
    }
    if IsTrue(OpGt(responseLength , MkInteger(1) )) {
      before := this.SafeValue(*(response ).At(MkInteger(1) ), MkString("before") ); _ = before;
      if IsTrue(OpNotEq2(before , MkUndefined() )) {
        orders = *(response ).At(MkInteger(0) );
      }
    }
  }
  return this.ParseOrders(orders , market , since , limit );
}

func (this *Okex3) FetchOpenOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  return this.FetchOrdersByState(MkString("6") , symbol , since , limit , params );
}

func (this *Okex3) FetchClosedOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  return this.FetchOrdersByState(MkString("7") , symbol , since , limit , params );
}

func (this *Okex3) ParseDepositAddress(goArgs ...*Variant) *Variant{
  depositAddress := GoGetArg(goArgs, 0, MkUndefined()); _ = depositAddress;
  currency := GoGetArg(goArgs, 1, MkUndefined() ); _ = currency;
  address := this.SafeString(depositAddress , MkString("address") ); _ = address;
  tag := this.SafeString2(depositAddress , MkString("tag") , MkString("payment_id") ); _ = tag;
  tag = this.SafeString2(depositAddress , MkString("memo") , MkString("Memo") , tag );
  currencyId := this.SafeString(depositAddress , MkString("currency") ); _ = currencyId;
  code := this.SafeCurrencyCode(currencyId ); _ = code;
  this.CheckAddress(address );
  return MkMap(&VarMap{
        "currency":code ,
        "address":address ,
        "tag":tag ,
        "info":depositAddress ,
        });
}

func (this *Okex3) FetchDepositAddress(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  parts := code.Split(MkString("-") ); _ = parts;
  currency := this.Currency(*(parts ).At(MkInteger(0) )); _ = currency;
  request := MkMap(&VarMap{"currency":*(currency ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("accountGetDepositAddress"), this.Extend(request , params )); _ = response;
  addressesByCode := this.ParseDepositAddresses(response ); _ = addressesByCode;
  address := this.SafeValue(addressesByCode , code ); _ = address;
  if IsTrue(OpEq2(address , MkUndefined() )) {
    panic(NewInvalidAddress(OpAdd(*this.At(MkString("id")) , MkString(" fetchDepositAddress cannot return nonexistent addresses, you should create withdrawal addresses with the exchange website first") )));
  }
  return address ;
}

func (this *Okex3) Withdraw(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  amount := GoGetArg(goArgs, 1, MkUndefined()); _ = amount;
  address := GoGetArg(goArgs, 2, MkUndefined()); _ = address;
  tag := GoGetArg(goArgs, 3, MkUndefined() ); _ = tag;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.CheckAddress(address );
  this.LoadMarkets();
  currency := this.Currency(code ); _ = currency;
  if IsTrue(tag ) {
    address = OpAdd(address , OpAdd(MkString(":") , tag ));
  }
  fee := this.SafeString(params , MkString("fee") ); _ = fee;
  if IsTrue(OpEq2(fee , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" withdraw() requires a `fee` string parameter, network transaction fee must be â¥ 0. Withdrawals to OKCoin or OKEx are fee-free, please set '0'. Withdrawing to external digital asset address requires network transaction fee.") )));
  }
  request := MkMap(&VarMap{
        "currency":*(currency ).At(MkString("id") ),
        "to_address":address ,
        "destination":MkString("4") ,
        "amount":this.NumberToString(amount ),
        "fee":fee ,
        }); _ = request;
  if IsTrue(OpHasMember(MkString("password") , params )) {
    *(request ).At(MkString("trade_pwd") )= *(params ).At(MkString("password") );
  } else {
    if IsTrue(OpHasMember(MkString("trade_pwd") , params )) {
      *(request ).At(MkString("trade_pwd") )= *(params ).At(MkString("trade_pwd") );
    } else {
      if IsTrue(*this.At(MkString("password")) ) {
        *(request ).At(MkString("trade_pwd") )= OpCopy(*this.At(MkString("password")) );
      }
    }
  }
  query := this.Omit(params , MkArray(&VarArray{
            MkString("fee") ,
            MkString("password") ,
            MkString("trade_pwd") ,
            })); _ = query;
  if IsTrue(OpNot(OpHasMember(MkString("trade_pwd") , request ))) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , MkString(" withdraw() requires this.password set on the exchange instance or a password / trade_pwd parameter") )));
  }
  response := this.Call(MkString("accountPostWithdrawal"), this.Extend(request , query )); _ = response;
  return MkMap(&VarMap{
        "info":response ,
        "id":this.SafeString(response , MkString("withdrawal_id") ),
        });
}

func (this *Okex3) FetchDeposits(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  method := MkString("accountGetDepositHistory") ; _ = method;
  currency := OpCopy(MkUndefined() ); _ = currency;
  if IsTrue(OpNotEq2(code , MkUndefined() )) {
    currency = this.Currency(code );
    *(request ).At(MkString("currency") )= *(currency ).At(MkString("id") );
    OpAddAssign(& method , MkString("Currency") );
  }
  response := this.Call(method , this.Extend(request , params )); _ = response;
  return this.ParseTransactions(response , currency , since , limit , params );
}

func (this *Okex3) FetchWithdrawals(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  method := MkString("accountGetWithdrawalHistory") ; _ = method;
  currency := OpCopy(MkUndefined() ); _ = currency;
  if IsTrue(OpNotEq2(code , MkUndefined() )) {
    currency = this.Currency(code );
    *(request ).At(MkString("currency") )= *(currency ).At(MkString("id") );
    OpAddAssign(& method , MkString("Currency") );
  }
  response := this.Call(method , this.Extend(request , params )); _ = response;
  return this.ParseTransactions(response , currency , since , limit , params );
}

func (this *Okex3) ParseTransactionStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "-3":MkString("pending") ,
        "-2":MkString("canceled") ,
        "-1":MkString("failed") ,
        "0":MkString("pending") ,
        "1":MkString("pending") ,
        "2":MkString("ok") ,
        "3":MkString("pending") ,
        "4":MkString("pending") ,
        "5":MkString("pending") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Okex3) ParseTransaction(goArgs ...*Variant) *Variant{
  transaction := GoGetArg(goArgs, 0, MkUndefined()); _ = transaction;
  currency := GoGetArg(goArgs, 1, MkUndefined() ); _ = currency;
  type_ := OpCopy(MkUndefined() ); _ = type_;
  id := OpCopy(MkUndefined() ); _ = id;
  address := OpCopy(MkUndefined() ); _ = address;
  withdrawalId := this.SafeString(transaction , MkString("withdrawal_id") ); _ = withdrawalId;
  addressFrom := this.SafeString(transaction , MkString("from") ); _ = addressFrom;
  addressTo := this.SafeString(transaction , MkString("to") ); _ = addressTo;
  tagTo := this.SafeString(transaction , MkString("tag") ); _ = tagTo;
  if IsTrue(OpNotEq2(withdrawalId , MkUndefined() )) {
    type_ = MkString("withdrawal") ;
    id = OpCopy(withdrawalId );
    address = OpCopy(addressTo );
  } else {
    id = this.SafeString2(transaction , MkString("payment_id") , MkString("deposit_id") );
    type_ = MkString("deposit") ;
    address = OpCopy(addressTo );
  }
  currencyId := this.SafeString(transaction , MkString("currency") ); _ = currencyId;
  code := this.SafeCurrencyCode(currencyId ); _ = code;
  amount := this.SafeNumber(transaction , MkString("amount") ); _ = amount;
  status := this.ParseTransactionStatus(this.SafeString(transaction , MkString("status") )); _ = status;
  txid := this.SafeString(transaction , MkString("txid") ); _ = txid;
  timestamp := this.Parse8601(this.SafeString(transaction , MkString("timestamp") )); _ = timestamp;
  feeCost := OpCopy(MkUndefined() ); _ = feeCost;
  if IsTrue(OpEq2(type_ , MkString("deposit") )) {
    feeCost = MkInteger(0) ;
  } else {
    if IsTrue(OpNotEq2(currencyId , MkUndefined() )) {
      feeWithCurrencyId := this.SafeString(transaction , MkString("fee") ); _ = feeWithCurrencyId;
      if IsTrue(OpNotEq2(feeWithCurrencyId , MkUndefined() )) {
        lowercaseCurrencyId := currencyId.ToLowerCase(); _ = lowercaseCurrencyId;
        feeWithoutCurrencyId := feeWithCurrencyId.Replace(lowercaseCurrencyId , MkString("") ); _ = feeWithoutCurrencyId;
        feeCost = parseFloat(feeWithoutCurrencyId );
      }
    }
  }
  return MkMap(&VarMap{
        "info":transaction ,
        "id":id ,
        "currency":code ,
        "amount":amount ,
        "addressFrom":addressFrom ,
        "addressTo":addressTo ,
        "address":address ,
        "tagFrom":MkUndefined() ,
        "tagTo":tagTo ,
        "tag":tagTo ,
        "status":status ,
        "type":type_ ,
        "updated":MkUndefined() ,
        "txid":txid ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "fee":MkMap(&VarMap{
            "currency":code ,
            "cost":feeCost ,
            }),
        });
}

func (this *Okex3) ParseMyTrade(goArgs ...*Variant) *Variant{
  pair := GoGetArg(goArgs, 0, MkUndefined()); _ = pair;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  userTrade := this.SafeValue(pair , MkInteger(1) ); _ = userTrade;
  otherTrade := this.SafeValue(pair , MkInteger(0) ); _ = otherTrade;
  firstMarketId := this.SafeString(otherTrade , MkString("instrument_id") ); _ = firstMarketId;
  secondMarketId := this.SafeString(userTrade , MkString("instrument_id") ); _ = secondMarketId;
  if IsTrue(OpNotEq2(firstMarketId , secondMarketId )) {
    panic(NewNotSupported(OpAdd(*this.At(MkString("id")) , MkString(" parseMyTrade() received unrecognized response format, differing instrument_ids in one fill, the exchange API might have changed, paste your verbose output: https://github.com/ccxt/ccxt/wiki/FAQ#what-is-required-to-get-help") )));
  }
  marketId := OpCopy(firstMarketId ); _ = marketId;
  market = this.SafeMarket(marketId , market );
  symbol := *(market ).At(MkString("symbol") ); _ = symbol;
  quoteId := *(market ).At(MkString("quoteId") ); _ = quoteId;
  side := OpCopy(MkUndefined() ); _ = side;
  amount := OpCopy(MkUndefined() ); _ = amount;
  cost := OpCopy(MkUndefined() ); _ = cost;
  receivedCurrencyId := this.SafeString(userTrade , MkString("currency") ); _ = receivedCurrencyId;
  feeCurrencyId := OpCopy(MkUndefined() ); _ = feeCurrencyId;
  if IsTrue(OpEq2(receivedCurrencyId , quoteId )) {
    side = this.SafeString(otherTrade , MkString("side") );
    amount = this.SafeNumber(otherTrade , MkString("size") );
    cost = this.SafeNumber(userTrade , MkString("size") );
    feeCurrencyId = this.SafeString(otherTrade , MkString("currency") );
  } else {
    side = this.SafeString(userTrade , MkString("side") );
    amount = this.SafeNumber(userTrade , MkString("size") );
    cost = this.SafeNumber(otherTrade , MkString("size") );
    feeCurrencyId = this.SafeString(userTrade , MkString("currency") );
  }
  id := this.SafeString(userTrade , MkString("trade_id") ); _ = id;
  price := this.SafeNumber(userTrade , MkString("price") ); _ = price;
  feeCostFirst := this.SafeNumber(otherTrade , MkString("fee") ); _ = feeCostFirst;
  feeCostSecond := this.SafeNumber(userTrade , MkString("fee") ); _ = feeCostSecond;
  feeCurrencyCodeFirst := this.SafeCurrencyCode(this.SafeString(otherTrade , MkString("currency") )); _ = feeCurrencyCodeFirst;
  feeCurrencyCodeSecond := this.SafeCurrencyCode(this.SafeString(userTrade , MkString("currency") )); _ = feeCurrencyCodeSecond;
  fee := OpCopy(MkUndefined() ); _ = fee;
  fees := OpCopy(MkUndefined() ); _ = fees;
  if IsTrue(OpAnd(OpNotEq2(feeCostFirst , MkUndefined() ), OpNotEq2(feeCostFirst , MkInteger(0) ))) {
    if IsTrue(OpAnd(OpNotEq2(feeCostSecond , MkUndefined() ), OpNotEq2(feeCostSecond , MkInteger(0) ))) {
      fees = MkArray(&VarArray{
          MkMap(&VarMap{
              "cost":OpNeg(feeCostFirst ),
              "currency":feeCurrencyCodeFirst ,
              }),
          MkMap(&VarMap{
              "cost":OpNeg(feeCostSecond ),
              "currency":feeCurrencyCodeSecond ,
              }),
          });
    } else {
      fee = MkMap(&VarMap{
          "cost":OpNeg(feeCostFirst ),
          "currency":feeCurrencyCodeFirst ,
          });
    }
  } else {
    if IsTrue(OpAnd(OpNotEq2(feeCostSecond , MkUndefined() ), OpNotEq2(feeCostSecond , MkInteger(0) ))) {
      fee = MkMap(&VarMap{
          "cost":OpNeg(feeCostSecond ),
          "currency":feeCurrencyCodeSecond ,
          });
    } else {
      fee = MkMap(&VarMap{
          "cost":MkInteger(0) ,
          "currency":this.SafeCurrencyCode(feeCurrencyId ),
          });
    }
  }
  timestamp := this.Parse8601(this.SafeString2(userTrade , MkString("timestamp") , MkString("created_at") )); _ = timestamp;
  takerOrMaker := this.SafeString2(userTrade , MkString("exec_type") , MkString("liquidity") ); _ = takerOrMaker;
  if IsTrue(OpEq2(takerOrMaker , MkString("M") )) {
    takerOrMaker = MkString("maker") ;
  } else {
    if IsTrue(OpEq2(takerOrMaker , MkString("T") )) {
      takerOrMaker = MkString("taker") ;
    }
  }
  orderId := this.SafeString(userTrade , MkString("order_id") ); _ = orderId;
  result := MkMap(&VarMap{
        "info":pair ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "symbol":symbol ,
        "id":id ,
        "order":orderId ,
        "type":MkUndefined() ,
        "takerOrMaker":takerOrMaker ,
        "side":side ,
        "price":price ,
        "amount":amount ,
        "cost":cost ,
        "fee":fee ,
        }); _ = result;
  if IsTrue(OpNotEq2(fees , MkUndefined() )) {
    *(result ).At(MkString("fees") )= OpCopy(fees );
  }
  return result ;
}

func (this *Okex3) ParseMyTrades(goArgs ...*Variant) *Variant{
  trades := GoGetArg(goArgs, 0, MkUndefined()); _ = trades;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  grouped := this.GroupBy(trades , MkString("trade_id") ); _ = grouped;
  tradeIds := GoGetKeys(grouped ); _ = tradeIds;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , tradeIds.Length )); OpIncr(& i ){
    tradeId := *(tradeIds ).At(i ); _ = tradeId;
    pair := *(grouped ).At(tradeId ); _ = pair;
    numTradesInPair := OpCopy(pair.Length ); _ = numTradesInPair;
    if IsTrue(OpEq2(numTradesInPair , MkInteger(2) )) {
      trade := this.ParseMyTrade(pair ); _ = trade;
      result.Push(trade );
    }
  }
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  if IsTrue(OpNotEq2(market , MkUndefined() )) {
    symbol = *(market ).At(MkString("symbol") );
  }
  return this.FilterBySymbolSinceLimit(result , symbol , since , limit );
}

func (this *Okex3) FetchMyTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchMyTrades() requires a symbol argument") )));
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  if IsTrue(OpAnd(OpNotEq2(limit , MkUndefined() ), OpGt(limit , MkInteger(100) ))) {
    limit = MkInteger(100) ;
  }
  request := MkMap(&VarMap{"instrument_id":*(market ).At(MkString("id") )}); _ = request;
  defaultType := this.SafeString2(*this.At(MkString("options")) , MkString("fetchMyTrades") , MkString("defaultType") ); _ = defaultType;
  type_ := this.SafeString(params , MkString("type") , defaultType ); _ = type_;
  query := this.Omit(params , MkString("type") ); _ = query;
  method := OpAdd(type_ , MkString("GetFills") ); _ = method;
  response := this.Call(method , this.Extend(request , query )); _ = response;
  return this.ParseMyTrades(response , market , since , limit , params );
}

func (this *Okex3) FetchOrderTrades(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"order_id":id }); _ = request;
  return this.FetchMyTrades(symbol , since , limit , this.Extend(request , params ));
}

func (this *Okex3) FetchPosition(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  method := OpCopy(MkUndefined() ); _ = method;
  request := MkMap(&VarMap{"instrument_id":*(market ).At(MkString("id") )}); _ = request;
  type_ := *(market ).At(MkString("type") ); _ = type_;
  if IsTrue(OpOr(OpEq2(type_ , MkString("futures") ), OpEq2(type_ , MkString("swap") ))) {
    method = OpAdd(type_ , MkString("GetInstrumentIdPosition") );
  } else {
    if IsTrue(OpEq2(type_ , MkString("option") )) {
      underlying := this.SafeString(params , MkString("underlying") ); _ = underlying;
      if IsTrue(OpEq2(underlying , MkUndefined() )) {
        panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" fetchPosition() requires an underlying parameter for ") , OpAdd(type_ , OpAdd(MkString(" market ") , symbol ))))));
      }
      method = OpAdd(type_ , MkString("GetUnderlyingPosition") );
    } else {
      panic(NewNotSupported(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" fetchPosition() does not support ") , OpAdd(type_ , OpAdd(MkString(" market ") , OpAdd(symbol , MkString(", supported market types are futures, swap or option") )))))));
    }
  }
  response := this.Call(method , this.Extend(request , params )); _ = response;
  return response ;
}

func (this *Okex3) FetchPositions(goArgs ...*Variant) *Variant{
  symbols := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbols;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  method := OpCopy(MkUndefined() ); _ = method;
  defaultType := this.SafeString2(*this.At(MkString("options")) , MkString("fetchPositions") , MkString("defaultType") ); _ = defaultType;
  type_ := this.SafeString(params , MkString("type") , defaultType ); _ = type_;
  if IsTrue(OpOr(OpEq2(type_ , MkString("futures") ), OpEq2(type_ , MkString("swap") ))) {
    method = OpAdd(type_ , MkString("GetPosition") );
  } else {
    if IsTrue(OpEq2(type_ , MkString("option") )) {
      underlying := this.SafeString(params , MkString("underlying") ); _ = underlying;
      if IsTrue(OpEq2(underlying , MkUndefined() )) {
        panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" fetchPositions() requires an underlying parameter for ") , OpAdd(type_ , MkString(" markets") )))));
      }
      method = OpAdd(type_ , MkString("GetUnderlyingPosition") );
    } else {
      panic(NewNotSupported(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" fetchPositions() does not support ") , OpAdd(type_ , MkString(" markets, supported market types are futures, swap or option") )))));
    }
  }
  params = this.Omit(params , MkString("type") );
  response := this.Call(method , params ); _ = response;
  return response ;
}

func (this *Okex3) FetchLedger(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  defaultType := this.SafeString2(*this.At(MkString("options")) , MkString("fetchLedger") , MkString("defaultType") ); _ = defaultType;
  type_ := this.SafeString(params , MkString("type") , defaultType ); _ = type_;
  query := this.Omit(params , MkString("type") ); _ = query;
  suffix := OpTrinary(OpEq2(type_ , MkString("account") ), MkString("") , MkString("Accounts") ); _ = suffix;
  argument := MkString("") ; _ = argument;
  request := MkMap(&VarMap{}); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  currency := OpCopy(MkUndefined() ); _ = currency;
  if IsTrue(OpEq2(type_ , MkString("spot") )) {
    if IsTrue(OpEq2(code , MkUndefined() )) {
      panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" fetchLedger() requires a currency code argument for '") , OpAdd(type_ , MkString("' markets") )))));
    }
    argument = MkString("Currency") ;
    currency = this.Currency(code );
    *(request ).At(MkString("currency") )= *(currency ).At(MkString("id") );
  } else {
    if IsTrue(OpEq2(type_ , MkString("futures") )) {
      if IsTrue(OpEq2(code , MkUndefined() )) {
        panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" fetchLedger() requires an underlying symbol for '") , OpAdd(type_ , MkString("' markets") )))));
      }
      argument = MkString("Underlying") ;
      market := this.Market(code ); _ = market;
      marketInfo := this.SafeValue(market , MkString("info") , MkMap(&VarMap{})); _ = marketInfo;
      settlementCurrencyId := this.SafeString(marketInfo , MkString("settlement_currency") ); _ = settlementCurrencyId;
      settlementCurrencyCode := this.SafeCurrencyCode(settlementCurrencyId ); _ = settlementCurrencyCode;
      currency = this.Currency(settlementCurrencyCode );
      underlyingId := this.SafeString(marketInfo , MkString("underlying") ); _ = underlyingId;
      *(request ).At(MkString("underlying") )= OpCopy(underlyingId );
    } else {
      if IsTrue(OpOr(OpEq2(type_ , MkString("margin") ), OpEq2(type_ , MkString("swap") ))) {
        if IsTrue(OpEq2(code , MkUndefined() )) {
          panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" fetchLedger() requires a code argument (a market symbol) for '") , OpAdd(type_ , MkString("' markets") )))));
        }
        argument = MkString("InstrumentId") ;
        market := this.Market(code ); _ = market;
        currency = this.Currency(*(market ).At(MkString("base") ));
        *(request ).At(MkString("instrument_id") )= *(market ).At(MkString("id") );
      } else {
        if IsTrue(OpEq2(type_ , MkString("account") )) {
          if IsTrue(OpNotEq2(code , MkUndefined() )) {
            currency = this.Currency(code );
            *(request ).At(MkString("currency") )= *(currency ).At(MkString("id") );
          }
        } else {
          panic(NewNotSupported(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" fetchLedger does not support the '") , OpAdd(type_ , MkString("' type (the type must be one of 'account', 'spot', 'margin', 'futures', 'swap')") )))));
        }
      }
    }
  }
  method := OpAdd(type_ , OpAdd(MkString("Get") , OpAdd(suffix , OpAdd(argument , MkString("Ledger") )))); _ = method;
  response := this.Call(method , this.Extend(request , query )); _ = response;
  responseLength := OpCopy(response.Length ); _ = responseLength;
  if IsTrue(OpLw(responseLength , MkInteger(1) )) {
    return MkArray(&VarArray{});
  }
  isArray := GoIsArray(*(response ).At(MkInteger(0) )); _ = isArray;
  isMargin := OpEq2(type_ , MkString("margin") ); _ = isMargin;
  entries := OpTrinary(OpAnd(isMargin , isArray ), *(response ).At(MkInteger(0) ), response ); _ = entries;
  if IsTrue(OpEq2(type_ , MkString("swap") )) {
    ledgerEntries := this.ParseLedger(entries ); _ = ledgerEntries;
    return this.FilterBySymbolSinceLimit(ledgerEntries , code , since , limit );
  }
  return this.ParseLedger(entries , currency , since , limit );
}

func (this *Okex3) ParseLedgerEntryType(goArgs ...*Variant) *Variant{
  type_ := GoGetArg(goArgs, 0, MkUndefined()); _ = type_;
  types := MkMap(&VarMap{
        "transfer":MkString("transfer") ,
        "trade":MkString("trade") ,
        "rebate":MkString("rebate") ,
        "match":MkString("trade") ,
        "fee":MkString("fee") ,
        "settlement":MkString("trade") ,
        "liquidation":MkString("trade") ,
        "funding":MkString("fee") ,
        "margin":MkString("margin") ,
        }); _ = types;
  return this.SafeString(types , type_ , type_ );
}

func (this *Okex3) ParseLedgerEntry(goArgs ...*Variant) *Variant{
  item := GoGetArg(goArgs, 0, MkUndefined()); _ = item;
  currency := GoGetArg(goArgs, 1, MkUndefined() ); _ = currency;
  id := this.SafeString(item , MkString("ledger_id") ); _ = id;
  account := OpCopy(MkUndefined() ); _ = account;
  details := this.SafeValue(item , MkString("details") , MkMap(&VarMap{})); _ = details;
  referenceId := this.SafeString(details , MkString("order_id") ); _ = referenceId;
  referenceAccount := OpCopy(MkUndefined() ); _ = referenceAccount;
  type_ := this.ParseLedgerEntryType(this.SafeString(item , MkString("type") )); _ = type_;
  code := this.SafeCurrencyCode(this.SafeString(item , MkString("currency") ), currency ); _ = code;
  amount := this.SafeNumber(item , MkString("amount") ); _ = amount;
  timestamp := this.Parse8601(this.SafeString(item , MkString("timestamp") )); _ = timestamp;
  fee := MkMap(&VarMap{
        "cost":this.SafeNumber(item , MkString("fee") ),
        "currency":code ,
        }); _ = fee;
  before := OpCopy(MkUndefined() ); _ = before;
  after := this.SafeNumber(item , MkString("balance") ); _ = after;
  status := MkString("ok") ; _ = status;
  marketId := this.SafeString(item , MkString("instrument_id") ); _ = marketId;
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  if IsTrue(OpHasMember(marketId , *this.At(MkString("markets_by_id")) )) {
    market := *(*this.At(MkString("markets_by_id")) ).At(marketId ); _ = market;
    symbol = *(market ).At(MkString("symbol") );
  }
  return MkMap(&VarMap{
        "info":item ,
        "id":id ,
        "account":account ,
        "referenceId":referenceId ,
        "referenceAccount":referenceAccount ,
        "type":type_ ,
        "currency":code ,
        "symbol":symbol ,
        "amount":amount ,
        "before":before ,
        "after":after ,
        "status":status ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "fee":fee ,
        });
}

func (this *Okex3) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("public") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  isArray := GoIsArray(params ); _ = isArray;
  request := OpAdd(MkString("/api/") , OpAdd(api , OpAdd(MkString("/") , OpAdd(*this.At(MkString("version")) , MkString("/") )))); _ = request;
  OpTrinary(OpAddAssign(& request , isArray ), path , this.ImplodeParams(path , params ));
  query := OpTrinary(isArray , params , this.Omit(params , this.ExtractParams(path ))); _ = query;
  url := OpAdd(this.ImplodeHostname(*(*(*this.At(MkString("urls")) ).At(MkString("api") )).At(MkString("rest") )), request ); _ = url;
  type_ := this.GetPathAuthenticationType(path ); _ = type_;
  if IsTrue(OpOr(OpEq2(type_ , MkString("public") ), OpEq2(type_ , MkString("information") ))) {
    if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
      OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(query )));
    }
  } else {
    if IsTrue(OpEq2(type_ , MkString("private") )) {
      this.CheckRequiredCredentials();
      timestamp := this.Iso8601(this.Milliseconds()); _ = timestamp;
      headers = MkMap(&VarMap{
          "OK-ACCESS-KEY":*this.At(MkString("apiKey")) ,
          "OK-ACCESS-PASSPHRASE":*this.At(MkString("password")) ,
          "OK-ACCESS-TIMESTAMP":timestamp ,
          });
      auth := OpAdd(timestamp , OpAdd(method , request )); _ = auth;
      if IsTrue(OpEq2(method , MkString("GET") )) {
        if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
          urlencodedQuery := OpAdd(MkString("?") , this.Urlencode(query )); _ = urlencodedQuery;
          OpAddAssign(& url , urlencodedQuery );
          OpAddAssign(& auth , urlencodedQuery );
        }
      } else {
        if IsTrue(OpOr(isArray , *(GoGetKeys(query )).At(MkString("length") ))) {
          body = this.Json(query );
          OpAddAssign(& auth , body );
        }
        *(headers ).At(MkString("Content-Type") )= MkString("application/json") ;
      }
      signature := this.Hmac(this.Encode(auth ), this.Encode(*this.At(MkString("secret")) ), MkString("sha256") , MkString("base64") ); _ = signature;
      *(headers ).At(MkString("OK-ACCESS-SIGN") )= OpCopy(signature );
    }
  }
  return MkMap(&VarMap{
        "url":url ,
        "method":method ,
        "body":body ,
        "headers":headers ,
        });
}

func (this *Okex3) GetPathAuthenticationType(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  if IsTrue(OpEq2(path , MkString("underlying") )) {
    return MkString("public") ;
  }
  auth := this.SafeValue(*this.At(MkString("options")) , MkString("auth") , MkMap(&VarMap{})); _ = auth;
  key := this.FindBroadlyMatchedKey(auth , path ); _ = key;
  return this.SafeString(auth , key , MkString("private") );
}

func (this *Okex3) HandleErrors(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  reason := GoGetArg(goArgs, 1, MkUndefined()); _ = reason;
  url := GoGetArg(goArgs, 2, MkUndefined()); _ = url;
  method := GoGetArg(goArgs, 3, MkUndefined()); _ = method;
  headers := GoGetArg(goArgs, 4, MkUndefined()); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined()); _ = body;
  response := GoGetArg(goArgs, 6, MkUndefined()); _ = response;
  requestHeaders := GoGetArg(goArgs, 7, MkUndefined()); _ = requestHeaders;
  requestBody := GoGetArg(goArgs, 8, MkUndefined()); _ = requestBody;
  if IsTrue(OpNot(response )) {
    return MkUndefined();
  }
  feedback := OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , body )); _ = feedback;
  if IsTrue(OpEq2(code , MkInteger(503) )) {
    panic(NewExchangeNotAvailable(feedback ));
  }
  message := this.SafeString(response , MkString("message") ); _ = message;
  errorCode := this.SafeString2(response , MkString("code") , MkString("error_code") ); _ = errorCode;
  nonEmptyMessage := OpAnd(OpNotEq2(message , MkUndefined() ), OpNotEq2(message , MkString("") )); _ = nonEmptyMessage;
  nonZeroErrorCode := OpAnd(OpNotEq2(errorCode , MkUndefined() ), OpNotEq2(errorCode , MkString("0") )); _ = nonZeroErrorCode;
  if IsTrue(nonEmptyMessage ) {
    this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("exact") ), message , feedback );
    this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("broad") ), message , feedback );
  }
  if IsTrue(nonZeroErrorCode ) {
    this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("exact") ), errorCode , feedback );
  }
  if IsTrue(OpOr(nonZeroErrorCode , nonEmptyMessage )) {
    panic(NewExchangeError(feedback ));
  }
  return MkUndefined()
}

