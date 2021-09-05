package ccxt

import (
)

type Bitstamp struct {
	*ExchangeBase
}

var _ Exchange = (*Bitstamp)(nil)

func init() {
	exchange := &Bitstamp{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Bitstamp) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("bitstamp") ,
            "name":MkString("Bitstamp") ,
            "countries":MkArray(&VarArray{
                MkString("GB") ,
                }),
            "rateLimit":MkInteger(1000) ,
            "version":MkString("v2") ,
            "userAgent":*(*this.At(MkString("userAgents")) ).At(MkString("chrome") ),
            "pro":MkBool(true) ,
            "has":MkMap(&VarMap{
                "CORS":MkBool(true) ,
                "cancelAllOrders":MkBool(true) ,
                "cancelOrder":MkBool(true) ,
                "createOrder":MkBool(true) ,
                "fetchBalance":MkBool(true) ,
                "fetchDepositAddress":MkBool(true) ,
                "fetchMarkets":MkBool(true) ,
                "fetchCurrencies":MkBool(true) ,
                "fetchMyTrades":MkBool(true) ,
                "fetchOHLCV":MkBool(true) ,
                "fetchOpenOrders":MkBool(true) ,
                "fetchOrder":MkBool(true) ,
                "fetchOrderBook":MkBool(true) ,
                "fetchTicker":MkBool(true) ,
                "fetchTrades":MkBool(true) ,
                "fetchTransactions":MkBool(true) ,
                "fetchWithdrawals":MkBool(true) ,
                "withdraw":MkBool(true) ,
                "fetchTradingFee":MkBool(true) ,
                "fetchTradingFees":MkBool(true) ,
                "fetchFundingFees":MkBool(true) ,
                "fetchFees":MkBool(true) ,
                "fetchLedger":MkBool(true) ,
                }),
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/1294454/27786377-8c8ab57e-5fe9-11e7-8ea4-2b05b6bcceec.jpg") ,
                "api":MkMap(&VarMap{
                    "public":MkString("https://www.bitstamp.net/api") ,
                    "private":MkString("https://www.bitstamp.net/api") ,
                    }),
                "www":MkString("https://www.bitstamp.net") ,
                "doc":MkString("https://www.bitstamp.net/api") ,
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
                "1w":MkString("259200") ,
                }),
            "requiredCredentials":MkMap(&VarMap{
                "apiKey":MkBool(true) ,
                "secret":MkBool(true) ,
                }),
            "api":MkMap(&VarMap{
                "public":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("ohlc/{pair}/") ,
                        MkString("order_book/{pair}/") ,
                        MkString("ticker_hour/{pair}/") ,
                        MkString("ticker/{pair}/") ,
                        MkString("transactions/{pair}/") ,
                        MkString("trading-pairs-info/") ,
                        })}),
                "private":MkMap(&VarMap{"post":MkArray(&VarArray{
                        MkString("balance/") ,
                        MkString("balance/{pair}/") ,
                        MkString("bch_withdrawal/") ,
                        MkString("bch_address/") ,
                        MkString("user_transactions/") ,
                        MkString("user_transactions/{pair}/") ,
                        MkString("open_orders/all/") ,
                        MkString("open_orders/{pair}/") ,
                        MkString("order_status/") ,
                        MkString("cancel_order/") ,
                        MkString("cancel_all_orders/") ,
                        MkString("cancel_all_orders/{pair}/") ,
                        MkString("buy/{pair}/") ,
                        MkString("buy/market/{pair}/") ,
                        MkString("buy/instant/{pair}/") ,
                        MkString("sell/{pair}/") ,
                        MkString("sell/market/{pair}/") ,
                        MkString("sell/instant/{pair}/") ,
                        MkString("btc_withdrawal/") ,
                        MkString("btc_address/") ,
                        MkString("ripple_withdrawal/") ,
                        MkString("ripple_address/") ,
                        MkString("ltc_withdrawal/") ,
                        MkString("ltc_address/") ,
                        MkString("eth_withdrawal/") ,
                        MkString("eth_address/") ,
                        MkString("xrp_withdrawal/") ,
                        MkString("xrp_address/") ,
                        MkString("xlm_withdrawal/") ,
                        MkString("xlm_address/") ,
                        MkString("pax_withdrawal/") ,
                        MkString("pax_address/") ,
                        MkString("link_withdrawal/") ,
                        MkString("link_address/") ,
                        MkString("usdc_withdrawal/") ,
                        MkString("usdc_address/") ,
                        MkString("omg_withdrawal/") ,
                        MkString("omg_address/") ,
                        MkString("dai_withdrawal/") ,
                        MkString("dai_address/") ,
                        MkString("knc_withdrawal/") ,
                        MkString("knc_address/") ,
                        MkString("mkr_withdrawal/") ,
                        MkString("mkr_address/") ,
                        MkString("zrx_withdrawal/") ,
                        MkString("zrx_address/") ,
                        MkString("gusd_withdrawal/") ,
                        MkString("gusd_address/") ,
                        MkString("aave_withdrawal/") ,
                        MkString("aave_address/") ,
                        MkString("bat_withdrawal/") ,
                        MkString("bat_address/") ,
                        MkString("uma_withdrawal/") ,
                        MkString("uma_address/") ,
                        MkString("snx_withdrawal/") ,
                        MkString("snx_address/") ,
                        MkString("uni_withdrawal/") ,
                        MkString("uni_address/") ,
                        MkString("yfi_withdrawal/") ,
                        MkString("yfi_address") ,
                        MkString("audio_withdrawal/") ,
                        MkString("audio_address/") ,
                        MkString("crv_withdrawal/") ,
                        MkString("crv_address/") ,
                        MkString("algo_withdrawal/") ,
                        MkString("algo_address/") ,
                        MkString("comp_withdrawal/") ,
                        MkString("comp_address/") ,
                        MkString("grt_withdrawal") ,
                        MkString("grt_address/") ,
                        MkString("usdt_withdrawal/") ,
                        MkString("usdt_address/") ,
                        MkString("eurt_withdrawal/") ,
                        MkString("eurt_address/") ,
                        MkString("matic_withdrawal/") ,
                        MkString("matic_address/") ,
                        MkString("sushi_withdrawal/") ,
                        MkString("sushi_address/") ,
                        MkString("chz_withdrawal/") ,
                        MkString("chz_address/") ,
                        MkString("enj_withdrawal/") ,
                        MkString("enj_address/") ,
                        MkString("transfer-to-main/") ,
                        MkString("transfer-from-main/") ,
                        MkString("withdrawal-requests/") ,
                        MkString("withdrawal/open/") ,
                        MkString("withdrawal/status/") ,
                        MkString("withdrawal/cancel/") ,
                        MkString("liquidation_address/new/") ,
                        MkString("liquidation_address/info/") ,
                        MkString("btc_unconfirmed/") ,
                        })}),
                }),
            "fees":MkMap(&VarMap{
                "trading":MkMap(&VarMap{
                    "tierBased":MkBool(true) ,
                    "percentage":MkBool(true) ,
                    "taker":this.ParseNumber(MkString("0.005") ),
                    "maker":this.ParseNumber(MkString("0.005") ),
                    "tiers":MkMap(&VarMap{
                        "taker":MkArray(&VarArray{
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("0") ),
                                this.ParseNumber(MkString("0.005") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("20000") ),
                                this.ParseNumber(MkString("0.0025") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("100000") ),
                                this.ParseNumber(MkString("0.0024") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("200000") ),
                                this.ParseNumber(MkString("0.0022") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("400000") ),
                                this.ParseNumber(MkString("0.0020") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("600000") ),
                                this.ParseNumber(MkString("0.0015") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("1000000") ),
                                this.ParseNumber(MkString("0.0014") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("2000000") ),
                                this.ParseNumber(MkString("0.0013") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("4000000") ),
                                this.ParseNumber(MkString("0.0012") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("20000000") ),
                                this.ParseNumber(MkString("0.0011") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("50000000") ),
                                this.ParseNumber(MkString("0.0010") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("100000000") ),
                                this.ParseNumber(MkString("0.0007") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("500000000") ),
                                this.ParseNumber(MkString("0.0005") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("2000000000") ),
                                this.ParseNumber(MkString("0.0003") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("6000000000") ),
                                this.ParseNumber(MkString("0.0001") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("20000000000") ),
                                this.ParseNumber(MkString("0.00005") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("20000000001") ),
                                this.ParseNumber(MkString("0") ),
                                }),
                            }),
                        "maker":MkArray(&VarArray{
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("0") ),
                                this.ParseNumber(MkString("0.005") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("20000") ),
                                this.ParseNumber(MkString("0.0025") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("100000") ),
                                this.ParseNumber(MkString("0.0024") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("200000") ),
                                this.ParseNumber(MkString("0.0022") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("400000") ),
                                this.ParseNumber(MkString("0.0020") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("600000") ),
                                this.ParseNumber(MkString("0.0015") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("1000000") ),
                                this.ParseNumber(MkString("0.0014") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("2000000") ),
                                this.ParseNumber(MkString("0.0013") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("4000000") ),
                                this.ParseNumber(MkString("0.0012") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("20000000") ),
                                this.ParseNumber(MkString("0.0011") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("50000000") ),
                                this.ParseNumber(MkString("0.0010") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("100000000") ),
                                this.ParseNumber(MkString("0.0007") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("500000000") ),
                                this.ParseNumber(MkString("0.0005") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("2000000000") ),
                                this.ParseNumber(MkString("0.0003") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("6000000000") ),
                                this.ParseNumber(MkString("0.0001") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("20000000000") ),
                                this.ParseNumber(MkString("0.00005") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("20000000001") ),
                                this.ParseNumber(MkString("0") ),
                                }),
                            }),
                        }),
                    }),
                "funding":MkMap(&VarMap{
                    "tierBased":MkBool(false) ,
                    "percentage":MkBool(false) ,
                    "withdraw":MkMap(&VarMap{}),
                    "deposit":MkMap(&VarMap{
                        "BTC":MkInteger(0) ,
                        "BCH":MkInteger(0) ,
                        "LTC":MkInteger(0) ,
                        "ETH":MkInteger(0) ,
                        "XRP":MkInteger(0) ,
                        "XLM":MkInteger(0) ,
                        "PAX":MkInteger(0) ,
                        "USD":MkNumber(7.5) ,
                        "EUR":MkInteger(0) ,
                        }),
                    }),
                }),
            "exceptions":MkMap(&VarMap{
                "exact":MkMap(&VarMap{
                    "No permission found":PermissionDenied ,
                    "API key not found":AuthenticationError ,
                    "IP address not allowed":PermissionDenied ,
                    "Invalid nonce":InvalidNonce ,
                    "Invalid signature":AuthenticationError ,
                    "Authentication failed":AuthenticationError ,
                    "Missing key, signature and nonce parameters":AuthenticationError ,
                    "Your account is frozen":PermissionDenied ,
                    "Please update your profile with your FATCA information, before using API.":PermissionDenied ,
                    "Order not found":OrderNotFound ,
                    "Price is more than 20% below market price.":InvalidOrder ,
                    "Bitstamp.net is under scheduled maintenance.":OnMaintenance ,
                    "Order could not be placed.":ExchangeNotAvailable ,
                    }),
                "broad":MkMap(&VarMap{
                    "Minimum order size is":InvalidOrder ,
                    "Check your account balance for details.":InsufficientFunds ,
                    "Ensure this value has at least":InvalidAddress ,
                    }),
                }),
            }));
}

func (this *Bitstamp) FetchMarkets(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.FetchMarketsFromCache(params ); _ = response;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , response.Length )); OpIncr(& i ){
    market := *(response ).At(i ); _ = market;
    name := this.SafeString(market , MkString("name") ); _ = name;
    base := MkUndefined(); _ = base;
    quote := MkUndefined(); _ = quote;
    ArrayUnpack(name.Split(MkString("/") ), &base, &quote);
    baseId := base.ToLowerCase(); _ = baseId;
    quoteId := quote.ToLowerCase(); _ = quoteId;
    base = this.SafeCurrencyCode(base );
    quote = this.SafeCurrencyCode(quote );
    symbol := OpAdd(base , OpAdd(MkString("/") , quote )); _ = symbol;
    symbolId := OpAdd(baseId , OpAdd(MkString("_") , quoteId )); _ = symbolId;
    id := this.SafeString(market , MkString("url_symbol") ); _ = id;
    amountPrecisionString := this.SafeString(market , MkString("base_decimals") ); _ = amountPrecisionString;
    pricePrecisionString := this.SafeString(market , MkString("counter_decimals") ); _ = pricePrecisionString;
    amountLimit := this.ParsePrecision(amountPrecisionString ); _ = amountLimit;
    priceLimit := this.ParsePrecision(pricePrecisionString ); _ = priceLimit;
    precision := MkMap(&VarMap{
          "amount":parseInt(amountPrecisionString ),
          "price":parseInt(pricePrecisionString ),
          }); _ = precision;
    minimumOrder := this.SafeString(market , MkString("minimum_order") ); _ = minimumOrder;
    parts := minimumOrder.Split(MkString(" ") ); _ = parts;
    cost := *(parts ).At(MkInteger(0) ); _ = cost;
    trading := this.SafeString(market , MkString("trading") ); _ = trading;
    active := OpEq2(trading , MkString("Enabled") ); _ = active;
    result.Push(MkMap(&VarMap{
            "id":id ,
            "symbol":symbol ,
            "base":base ,
            "quote":quote ,
            "baseId":baseId ,
            "quoteId":quoteId ,
            "symbolId":symbolId ,
            "info":market ,
            "active":active ,
            "precision":precision ,
            "limits":MkMap(&VarMap{
                "amount":MkMap(&VarMap{
                    "min":this.ParseNumber(amountLimit ),
                    "max":MkUndefined() ,
                    }),
                "price":MkMap(&VarMap{
                    "min":this.ParseNumber(priceLimit ),
                    "max":MkUndefined() ,
                    }),
                "cost":MkMap(&VarMap{
                    "min":this.ParseNumber(cost ),
                    "max":MkUndefined() ,
                    }),
                }),
            }));
  }
  return result ;
}

func (this *Bitstamp) ConstructCurrencyObject(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  code := GoGetArg(goArgs, 1, MkUndefined()); _ = code;
  name := GoGetArg(goArgs, 2, MkUndefined()); _ = name;
  precision := GoGetArg(goArgs, 3, MkUndefined()); _ = precision;
  minCost := GoGetArg(goArgs, 4, MkUndefined()); _ = minCost;
  originalPayload := GoGetArg(goArgs, 5, MkUndefined()); _ = originalPayload;
  currencyType := MkString("crypto") ; _ = currencyType;
  description := this.Describe(); _ = description;
  if IsTrue(this.IsFiat(code )) {
    currencyType = MkString("fiat") ;
  }
  return MkMap(&VarMap{
        "id":id ,
        "code":code ,
        "info":originalPayload ,
        "type":currencyType ,
        "name":name ,
        "active":MkBool(true) ,
        "fee":this.SafeNumber(*(*(*(description ).At(MkString("fees") )).At(MkString("funding") )).At(MkString("withdraw") ), code ),
        "precision":precision ,
        "limits":MkMap(&VarMap{
            "amount":MkMap(&VarMap{
                "min":Math.Pow(MkInteger(10) , OpNeg(precision )),
                "max":MkUndefined() ,
                }),
            "price":MkMap(&VarMap{
                "min":Math.Pow(MkInteger(10) , OpNeg(precision )),
                "max":MkUndefined() ,
                }),
            "cost":MkMap(&VarMap{
                "min":minCost ,
                "max":MkUndefined() ,
                }),
            "withdraw":MkMap(&VarMap{
                "min":MkUndefined() ,
                "max":MkUndefined() ,
                }),
            }),
        });
}

func (this *Bitstamp) FetchMarketsFromCache(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  options := this.SafeValue(*this.At(MkString("options")) , MkString("fetchMarkets") , MkMap(&VarMap{})); _ = options;
  timestamp := this.SafeInteger(options , MkString("timestamp") ); _ = timestamp;
  expires := this.SafeInteger(options , MkString("expires") , MkInteger(1000) ); _ = expires;
  now := this.Milliseconds(); _ = now;
  if IsTrue(OpOr(OpEq2(timestamp , MkUndefined() ), OpGt(OpSub(now , timestamp ), expires ))) {
    response := this.Call(MkString("publicGetTradingPairsInfo"), params ); _ = response;
    *(*this.At(MkString("options")) ).At(MkString("fetchMarkets") )= this.Extend(options , MkMap(&VarMap{
            "response":response ,
            "timestamp":now ,
            }));
  }
  return this.SafeValue(*(*this.At(MkString("options")) ).At(MkString("fetchMarkets") ), MkString("response") );
}

func (this *Bitstamp) FetchCurrencies(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.FetchMarketsFromCache(params ); _ = response;
  result := MkMap(&VarMap{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , response.Length )); OpIncr(& i ){
    market := *(response ).At(i ); _ = market;
    name := this.SafeString(market , MkString("name") ); _ = name;
    base := MkUndefined(); _ = base;
    quote := MkUndefined(); _ = quote;
    ArrayUnpack(name.Split(MkString("/") ), &base, &quote);
    baseId := base.ToLowerCase(); _ = baseId;
    quoteId := quote.ToLowerCase(); _ = quoteId;
    base = this.SafeCurrencyCode(base );
    quote = this.SafeCurrencyCode(quote );
    description := this.SafeString(market , MkString("description") ); _ = description;
    baseDescription := MkUndefined(); _ = baseDescription;
    quoteDescription := MkUndefined(); _ = quoteDescription;
    ArrayUnpack(description.Split(MkString(" / ") ), &baseDescription, &quoteDescription);
    minimumOrder := this.SafeString(market , MkString("minimum_order") ); _ = minimumOrder;
    parts := minimumOrder.Split(MkString(" ") ); _ = parts;
    cost := *(parts ).At(MkInteger(0) ); _ = cost;
    if IsTrue(OpNot(OpHasMember(base , result ))) {
      baseDecimals := this.SafeInteger(market , MkString("base_decimals") ); _ = baseDecimals;
      *(result ).At(base )= this.ConstructCurrencyObject(baseId , base , baseDescription , baseDecimals , MkUndefined() , market );
    }
    if IsTrue(OpNot(OpHasMember(quote , result ))) {
      counterDecimals := this.SafeInteger(market , MkString("counter_decimals") ); _ = counterDecimals;
      *(result ).At(quote )= this.ConstructCurrencyObject(quoteId , quote , quoteDescription , counterDecimals , this.ParseNumber(cost ), market );
    }
  }
  return result ;
}

func (this *Bitstamp) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"pair":this.MarketId(symbol )}); _ = request;
  response := this.Call(MkString("publicGetOrderBookPair"), this.Extend(request , params )); _ = response;
  microtimestamp := this.SafeInteger(response , MkString("microtimestamp") ); _ = microtimestamp;
  timestamp := parseInt(OpDiv(microtimestamp , MkInteger(1000) )); _ = timestamp;
  orderbook := this.ParseOrderBook(response , symbol , timestamp ); _ = orderbook;
  *(orderbook ).At(MkString("nonce") )= OpCopy(microtimestamp );
  return orderbook ;
}

func (this *Bitstamp) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"pair":this.MarketId(symbol )}); _ = request;
  ticker := this.Call(MkString("publicGetTickerPair"), this.Extend(request , params )); _ = ticker;
  timestamp := this.SafeTimestamp(ticker , MkString("timestamp") ); _ = timestamp;
  vwap := this.SafeNumber(ticker , MkString("vwap") ); _ = vwap;
  baseVolume := this.SafeNumber(ticker , MkString("volume") ); _ = baseVolume;
  quoteVolume := OpCopy(MkUndefined() ); _ = quoteVolume;
  if IsTrue(OpAnd(OpNotEq2(baseVolume , MkUndefined() ), OpNotEq2(vwap , MkUndefined() ))) {
    quoteVolume = OpMulti(baseVolume , vwap );
  }
  last := this.SafeNumber(ticker , MkString("last") ); _ = last;
  return MkMap(&VarMap{
        "symbol":symbol ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "high":this.SafeNumber(ticker , MkString("high") ),
        "low":this.SafeNumber(ticker , MkString("low") ),
        "bid":this.SafeNumber(ticker , MkString("bid") ),
        "bidVolume":MkUndefined() ,
        "ask":this.SafeNumber(ticker , MkString("ask") ),
        "askVolume":MkUndefined() ,
        "vwap":vwap ,
        "open":this.SafeNumber(ticker , MkString("open") ),
        "close":last ,
        "last":last ,
        "previousClose":MkUndefined() ,
        "change":MkUndefined() ,
        "percentage":MkUndefined() ,
        "average":MkUndefined() ,
        "baseVolume":baseVolume ,
        "quoteVolume":quoteVolume ,
        "info":ticker ,
        });
}

func (this *Bitstamp) GetCurrencyIdFromTransaction(goArgs ...*Variant) *Variant{
  transaction := GoGetArg(goArgs, 0, MkUndefined()); _ = transaction;
  currencyId := this.SafeStringLower(transaction , MkString("currency") ); _ = currencyId;
  if IsTrue(OpNotEq2(currencyId , MkUndefined() )) {
    return currencyId ;
  }
  transaction = this.Omit(transaction , MkArray(&VarArray{
          MkString("fee") ,
          MkString("price") ,
          MkString("datetime") ,
          MkString("type") ,
          MkString("status") ,
          MkString("id") ,
          }));
  ids := GoGetKeys(transaction ); _ = ids;
  for i := MkInteger(0) ; IsTrue(OpLw(i , ids.Length )); OpIncr(& i ){
    id := *(ids ).At(i ); _ = id;
    if IsTrue(OpLw(id.IndexOf(MkString("_") ), MkInteger(0) )) {
      value := this.SafeNumber(transaction , id ); _ = value;
      if IsTrue(OpAnd(OpNotEq2(value , MkUndefined() ), OpNotEq2(value , MkInteger(0) ))) {
        return id ;
      }
    }
  }
  return MkUndefined() ;
}

func (this *Bitstamp) GetMarketFromTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  trade = this.Omit(trade , MkArray(&VarArray{
          MkString("fee") ,
          MkString("price") ,
          MkString("datetime") ,
          MkString("tid") ,
          MkString("type") ,
          MkString("order_id") ,
          MkString("side") ,
          }));
  currencyIds := GoGetKeys(trade ); _ = currencyIds;
  numCurrencyIds := OpCopy(currencyIds.Length ); _ = numCurrencyIds;
  if IsTrue(OpGt(numCurrencyIds , MkInteger(2) )) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" getMarketFromTrade too many keys: ") , OpAdd(this.Json(currencyIds ), OpAdd(MkString(" in the trade: ") , this.Json(trade )))))));
  }
  if IsTrue(OpEq2(numCurrencyIds , MkInteger(2) )) {
    marketId := OpAdd(*(currencyIds ).At(MkInteger(0) ), *(currencyIds ).At(MkInteger(1) )); _ = marketId;
    if IsTrue(OpHasMember(marketId , *this.At(MkString("markets_by_id")) )) {
      return *(*this.At(MkString("markets_by_id")) ).At(marketId );
    }
    marketId = OpAdd(*(currencyIds ).At(MkInteger(1) ), *(currencyIds ).At(MkInteger(0) ));
    if IsTrue(OpHasMember(marketId , *this.At(MkString("markets_by_id")) )) {
      return *(*this.At(MkString("markets_by_id")) ).At(marketId );
    }
  }
  return MkUndefined() ;
}

func (this *Bitstamp) GetMarketFromTrades(goArgs ...*Variant) *Variant{
  trades := GoGetArg(goArgs, 0, MkUndefined()); _ = trades;
  tradesBySymbol := this.IndexBy(trades , MkString("symbol") ); _ = tradesBySymbol;
  symbols := GoGetKeys(tradesBySymbol ); _ = symbols;
  numSymbols := OpCopy(symbols.Length ); _ = numSymbols;
  if IsTrue(OpEq2(numSymbols , MkInteger(1) )) {
    return *(*this.At(MkString("markets")) ).At(*(symbols ).At(MkInteger(0) ));
  }
  return MkUndefined() ;
}

func (this *Bitstamp) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  id := this.SafeString2(trade , MkString("id") , MkString("tid") ); _ = id;
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  side := OpCopy(MkUndefined() ); _ = side;
  price := this.SafeNumber(trade , MkString("price") ); _ = price;
  amount := this.SafeNumber(trade , MkString("amount") ); _ = amount;
  orderId := this.SafeString(trade , MkString("order_id") ); _ = orderId;
  type_ := OpCopy(MkUndefined() ); _ = type_;
  cost := this.SafeNumber(trade , MkString("cost") ); _ = cost;
  if IsTrue(OpEq2(market , MkUndefined() )) {
    keys := GoGetKeys(trade ); _ = keys;
    for i := MkInteger(0) ; IsTrue(OpLw(i , keys.Length )); OpIncr(& i ){
      if IsTrue(OpGtEq((*(keys ).At(i )).Call(MkString("indexOf") , MkString("_") ), MkInteger(0) )) {
        marketId := (*(keys ).At(i )).Call(MkString("replace") , MkString("_") , MkString("") ); _ = marketId;
        if IsTrue(OpHasMember(marketId , *this.At(MkString("markets_by_id")) )) {
          market = *(*this.At(MkString("markets_by_id")) ).At(marketId );
        }
      }
    }
    if IsTrue(OpEq2(market , MkUndefined() )) {
      market = this.GetMarketFromTrade(trade );
    }
  }
  feeCost := this.SafeNumber(trade , MkString("fee") ); _ = feeCost;
  feeCurrency := OpCopy(MkUndefined() ); _ = feeCurrency;
  if IsTrue(OpNotEq2(market , MkUndefined() )) {
    price = this.SafeNumber(trade , *(market ).At(MkString("symbolId") ), price );
    amount = this.SafeNumber(trade , *(market ).At(MkString("baseId") ), amount );
    cost = this.SafeNumber(trade , *(market ).At(MkString("quoteId") ), cost );
    feeCurrency = *(market ).At(MkString("quote") );
    symbol = *(market ).At(MkString("symbol") );
  }
  timestamp := this.SafeString2(trade , MkString("date") , MkString("datetime") ); _ = timestamp;
  if IsTrue(OpNotEq2(timestamp , MkUndefined() )) {
    if IsTrue(OpGtEq(timestamp.IndexOf(MkString(" ") ), MkInteger(0) )) {
      timestamp = this.Parse8601(timestamp );
    } else {
      timestamp = parseInt(timestamp );
      timestamp = OpMulti(timestamp , MkInteger(1000) );
    }
  }
  if IsTrue(OpHasMember(MkString("id") , trade )) {
    if IsTrue(OpNotEq2(amount , MkUndefined() )) {
      if IsTrue(OpLw(amount , MkInteger(0) )) {
        side = MkString("sell") ;
        amount = OpNeg(amount );
      } else {
        side = MkString("buy") ;
      }
    }
  } else {
    side = this.SafeString(trade , MkString("type") );
    if IsTrue(OpEq2(side , MkString("1") )) {
      side = MkString("sell") ;
    } else {
      if IsTrue(OpEq2(side , MkString("0") )) {
        side = MkString("buy") ;
      }
    }
  }
  if IsTrue(OpEq2(cost , MkUndefined() )) {
    if IsTrue(OpNotEq2(price , MkUndefined() )) {
      if IsTrue(OpNotEq2(amount , MkUndefined() )) {
        cost = OpMulti(price , amount );
      }
    }
  }
  if IsTrue(OpNotEq2(cost , MkUndefined() )) {
    cost = Math.Abs(cost );
  }
  fee := OpCopy(MkUndefined() ); _ = fee;
  if IsTrue(OpNotEq2(feeCost , MkUndefined() )) {
    fee = MkMap(&VarMap{
        "cost":feeCost ,
        "currency":feeCurrency ,
        });
  }
  return MkMap(&VarMap{
        "id":id ,
        "info":trade ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "symbol":symbol ,
        "order":orderId ,
        "type":type_ ,
        "side":side ,
        "takerOrMaker":MkUndefined() ,
        "price":price ,
        "amount":amount ,
        "cost":cost ,
        "fee":fee ,
        });
}

func (this *Bitstamp) ParseTradingFee(goArgs ...*Variant) *Variant{
  balances := GoGetArg(goArgs, 0, MkUndefined()); _ = balances;
  symbol := GoGetArg(goArgs, 1, MkUndefined()); _ = symbol;
  market := this.Market(symbol ); _ = market;
  tradeFee := this.SafeNumber(balances , OpAdd(*(market ).At(MkString("id") ), MkString("_fee") )); _ = tradeFee;
  return MkMap(&VarMap{
        "symbol":symbol ,
        "maker":tradeFee ,
        "taker":tradeFee ,
        });
}

func (this *Bitstamp) FetchTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "pair":*(market ).At(MkString("id") ),
        "time":MkString("hour") ,
        }); _ = request;
  response := this.Call(MkString("publicGetTransactionsPair"), this.Extend(request , params )); _ = response;
  return this.ParseTrades(response , market , since , limit );
}

func (this *Bitstamp) ParseOHLCV(goArgs ...*Variant) *Variant{
  ohlcv := GoGetArg(goArgs, 0, MkUndefined()); _ = ohlcv;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  return MkArray(&VarArray{
        this.SafeTimestamp(ohlcv , MkString("timestamp") ),
        this.SafeNumber(ohlcv , MkString("open") ),
        this.SafeNumber(ohlcv , MkString("high") ),
        this.SafeNumber(ohlcv , MkString("low") ),
        this.SafeNumber(ohlcv , MkString("close") ),
        this.SafeNumber(ohlcv , MkString("volume") ),
        });
}

func (this *Bitstamp) FetchOHLCV(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  timeframe := GoGetArg(goArgs, 1, MkString("1m") ); _ = timeframe;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "pair":*(market ).At(MkString("id") ),
        "step":*(*this.At(MkString("timeframes")) ).At(timeframe ),
        }); _ = request;
  duration := this.ParseTimeframe(timeframe ); _ = duration;
  if IsTrue(OpEq2(limit , MkUndefined() )) {
    if IsTrue(OpEq2(since , MkUndefined() )) {
      panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchOHLCV() requires a since argument or a limit argument") )));
    } else {
      limit = MkInteger(1000) ;
      start := parseInt(OpDiv(since , MkInteger(1000) )); _ = start;
      *(request ).At(MkString("start") )= OpCopy(start );
      *(request ).At(MkString("end") )= this.Sum(start , OpMulti(limit , duration ));
      *(request ).At(MkString("limit") )= OpCopy(limit );
    }
  } else {
    if IsTrue(OpNotEq2(since , MkUndefined() )) {
      start := parseInt(OpDiv(since , MkInteger(1000) )); _ = start;
      *(request ).At(MkString("start") )= OpCopy(start );
      *(request ).At(MkString("end") )= this.Sum(start , OpMulti(limit , duration ));
    }
    *(request ).At(MkString("limit") )= Math.Min(limit , MkInteger(1000) );
  }
  response := this.Call(MkString("publicGetOhlcPair"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  ohlc := this.SafeValue(data , MkString("ohlc") , MkArray(&VarArray{})); _ = ohlc;
  return this.ParseOHLCVs(ohlc , market , timeframe , since , limit );
}

func (this *Bitstamp) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  balance := this.Call(MkString("privatePostBalance"), params ); _ = balance;
  result := MkMap(&VarMap{
        "info":balance ,
        "timestamp":MkUndefined() ,
        "datetime":MkUndefined() ,
        }); _ = result;
  codes := GoGetKeys(*this.At(MkString("currencies")) ); _ = codes;
  for i := MkInteger(0) ; IsTrue(OpLw(i , codes.Length )); OpIncr(& i ){
    code := *(codes ).At(i ); _ = code;
    currency := this.Currency(code ); _ = currency;
    currencyId := *(currency ).At(MkString("id") ); _ = currencyId;
    account := this.Account(); _ = account;
    *(account ).At(MkString("free") )= this.SafeString(balance , OpAdd(currencyId , MkString("_available") ));
    *(account ).At(MkString("used") )= this.SafeString(balance , OpAdd(currencyId , MkString("_reserved") ));
    *(account ).At(MkString("total") )= this.SafeString(balance , OpAdd(currencyId , MkString("_balance") ));
    *(result ).At(code )= OpCopy(account );
  }
  return this.ParseBalance(result );
}

func (this *Bitstamp) FetchTradingFee(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  method := MkString("privatePostBalance") ; _ = method;
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("pair") )= *(market ).At(MkString("id") );
    OpAddAssign(& method , MkString("Pair") );
  }
  balance := this.Call(method , this.Extend(request , params )); _ = balance;
  return MkMap(&VarMap{
        "info":balance ,
        "symbol":symbol ,
        "maker":*(balance ).At(MkString("fee") ),
        "taker":*(balance ).At(MkString("fee") ),
        });
}

func (this *Bitstamp) PraseTradingFees(goArgs ...*Variant) *Variant{
  balance := GoGetArg(goArgs, 0, MkUndefined()); _ = balance;
  result := MkMap(&VarMap{"info":balance }); _ = result;
  markets := GoGetKeys(*this.At(MkString("markets")) ); _ = markets;
  for i := MkInteger(0) ; IsTrue(OpLw(i , markets.Length )); OpIncr(& i ){
    symbol := *(markets ).At(i ); _ = symbol;
    fee := this.ParseTradingFee(balance , symbol ); _ = fee;
    *(result ).At(symbol )= OpCopy(fee );
  }
  return result ;
}

func (this *Bitstamp) FetchTradingFees(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  balance := this.Call(MkString("privatePostBalance"), params ); _ = balance;
  return this.PraseTradingFees(balance );
}

func (this *Bitstamp) ParseFundingFees(goArgs ...*Variant) *Variant{
  balance := GoGetArg(goArgs, 0, MkUndefined()); _ = balance;
  withdraw := MkMap(&VarMap{}); _ = withdraw;
  ids := GoGetKeys(balance ); _ = ids;
  for i := MkInteger(0) ; IsTrue(OpLw(i , ids.Length )); OpIncr(& i ){
    id := *(ids ).At(i ); _ = id;
    if IsTrue(OpGtEq(id.IndexOf(MkString("_withdrawal_fee") ), MkInteger(0) )) {
      currencyId := *(id.Split(MkString("_") )).At(MkInteger(0) ); _ = currencyId;
      code := this.SafeCurrencyCode(currencyId ); _ = code;
      *(withdraw ).At(code )= this.SafeNumber(balance , id );
    }
  }
  return MkMap(&VarMap{
        "info":balance ,
        "withdraw":withdraw ,
        "deposit":MkMap(&VarMap{}),
        });
}

func (this *Bitstamp) FetchFundingFees(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  balance := this.Call(MkString("privatePostBalance"), params ); _ = balance;
  return this.ParseFundingFees(balance );
}

func (this *Bitstamp) FetchFees(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  balance := this.Call(MkString("privatePostBalance"), params ); _ = balance;
  tradingFees := this.PraseTradingFees(balance ); _ = tradingFees;
  *(tradingFees ).At(MkString("info") ) = MkUndefined();
  fundingFees := this.ParseFundingFees(balance ); _ = fundingFees;
  *(fundingFees ).At(MkString("info") ) = MkUndefined();
  return MkMap(&VarMap{
        "info":balance ,
        "trading":tradingFees ,
        "funding":fundingFees ,
        });
}

func (this *Bitstamp) CreateOrder(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 2, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 3, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 4, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 5, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  method := OpAdd(MkString("privatePost") , this.Capitalize(side )); _ = method;
  request := MkMap(&VarMap{
        "pair":*(market ).At(MkString("id") ),
        "amount":this.AmountToPrecision(symbol , amount ),
        }); _ = request;
  if IsTrue(OpEq2(type_ , MkString("market") )) {
    OpAddAssign(& method , MkString("Market") );
  } else {
    if IsTrue(OpEq2(type_ , MkString("instant") )) {
      OpAddAssign(& method , MkString("Instant") );
    } else {
      *(request ).At(MkString("price") )= this.PriceToPrecision(symbol , price );
    }
  }
  OpAddAssign(& method , MkString("Pair") );
  response := this.Call(method , this.Extend(request , params )); _ = response;
  order := this.ParseOrder(response , market ); _ = order;
  return this.Extend(order , MkMap(&VarMap{"type":type_ }));
}

func (this *Bitstamp) CancelOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"id":id }); _ = request;
  return this.Call(MkString("privatePostCancelOrder"), this.Extend(request , params ));
}

func (this *Bitstamp) CancelAllOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := OpCopy(MkUndefined() ); _ = market;
  request := MkMap(&VarMap{}); _ = request;
  method := MkString("privatePostCancelAllOrders") ; _ = method;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("pair") )= *(market ).At(MkString("id") );
    method = MkString("privatePostCancelAllOrdersPair") ;
  }
  return this.Call(method , this.Extend(request , params ));
}

func (this *Bitstamp) ParseOrderStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "In Queue":MkString("open") ,
        "Open":MkString("open") ,
        "Finished":MkString("closed") ,
        "Canceled":MkString("canceled") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Bitstamp) FetchOrderStatus(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"id":id }); _ = request;
  response := this.Call(MkString("privatePostOrderStatus"), this.Extend(request , params )); _ = response;
  return this.ParseOrderStatus(this.SafeString(response , MkString("status") ));
}

func (this *Bitstamp) FetchOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
  }
  request := MkMap(&VarMap{"id":id }); _ = request;
  response := this.Call(MkString("privatePostOrderStatus"), this.Extend(request , params )); _ = response;
  return this.ParseOrder(response , market );
}

func (this *Bitstamp) FetchMyTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  method := MkString("privatePostUserTransactions") ; _ = method;
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("pair") )= *(market ).At(MkString("id") );
    OpAddAssign(& method , MkString("Pair") );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(method , this.Extend(request , params )); _ = response;
  result := this.FilterBy(response , MkString("type") , MkString("2") ); _ = result;
  return this.ParseTrades(result , market , since , limit );
}

func (this *Bitstamp) FetchTransactions(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("privatePostUserTransactions"), this.Extend(request , params )); _ = response;
  currency := OpCopy(MkUndefined() ); _ = currency;
  if IsTrue(OpNotEq2(code , MkUndefined() )) {
    currency = this.Currency(code );
  }
  transactions := this.FilterByArray(response , MkString("type") , MkArray(&VarArray{
            MkString("0") ,
            MkString("1") ,
            }), MkBool(false) ); _ = transactions;
  return this.ParseTransactions(transactions , currency , since , limit );
}

func (this *Bitstamp) FetchWithdrawals(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("timedelta") )= OpSub(this.Milliseconds(), since );
  } else {
    *(request ).At(MkString("timedelta") )= MkInteger(50000000) ;
  }
  response := this.Call(MkString("privatePostWithdrawalRequests"), this.Extend(request , params )); _ = response;
  return this.ParseTransactions(response , MkUndefined() , since , limit );
}

func (this *Bitstamp) ParseTransaction(goArgs ...*Variant) *Variant{
  transaction := GoGetArg(goArgs, 0, MkUndefined()); _ = transaction;
  currency := GoGetArg(goArgs, 1, MkUndefined() ); _ = currency;
  timestamp := this.Parse8601(this.SafeString(transaction , MkString("datetime") )); _ = timestamp;
  id := this.SafeString(transaction , MkString("id") ); _ = id;
  currencyId := this.GetCurrencyIdFromTransaction(transaction ); _ = currencyId;
  code := this.SafeCurrencyCode(currencyId , currency ); _ = code;
  feeCost := this.SafeNumber(transaction , MkString("fee") ); _ = feeCost;
  feeCurrency := OpCopy(MkUndefined() ); _ = feeCurrency;
  amount := OpCopy(MkUndefined() ); _ = amount;
  if IsTrue(OpHasMember(MkString("amount") , transaction )) {
    amount = this.SafeNumber(transaction , MkString("amount") );
  } else {
    if IsTrue(OpNotEq2(currency , MkUndefined() )) {
      amount = this.SafeNumber(transaction , *(currency ).At(MkString("id") ), amount );
      feeCurrency = *(currency ).At(MkString("code") );
    } else {
      if IsTrue(OpAnd(OpNotEq2(code , MkUndefined() ), OpNotEq2(currencyId , MkUndefined() ))) {
        amount = this.SafeNumber(transaction , currencyId , amount );
        feeCurrency = OpCopy(code );
      }
    }
  }
  if IsTrue(OpNotEq2(amount , MkUndefined() )) {
    amount = Math.Abs(amount );
  }
  status := MkString("ok") ; _ = status;
  if IsTrue(OpHasMember(MkString("status") , transaction )) {
    status = this.ParseTransactionStatus(this.SafeString(transaction , MkString("status") ));
  }
  type_ := OpCopy(MkUndefined() ); _ = type_;
  if IsTrue(OpHasMember(MkString("type") , transaction )) {
    rawType := this.SafeString(transaction , MkString("type") ); _ = rawType;
    if IsTrue(OpEq2(rawType , MkString("0") )) {
      type_ = MkString("deposit") ;
    } else {
      if IsTrue(OpEq2(rawType , MkString("1") )) {
        type_ = MkString("withdrawal") ;
      }
    }
  } else {
    type_ = MkString("withdrawal") ;
  }
  txid := this.SafeString(transaction , MkString("transaction_id") ); _ = txid;
  tag := OpCopy(MkUndefined() ); _ = tag;
  address := this.SafeString(transaction , MkString("address") ); _ = address;
  if IsTrue(OpNotEq2(address , MkUndefined() )) {
    addressParts := address.Split(MkString("?dt=") ); _ = addressParts;
    numParts := OpCopy(addressParts.Length ); _ = numParts;
    if IsTrue(OpGt(numParts , MkInteger(1) )) {
      address = *(addressParts ).At(MkInteger(0) );
      tag = *(addressParts ).At(MkInteger(1) );
    }
  }
  addressFrom := OpCopy(MkUndefined() ); _ = addressFrom;
  addressTo := OpCopy(address ); _ = addressTo;
  tagFrom := OpCopy(MkUndefined() ); _ = tagFrom;
  tagTo := OpCopy(tag ); _ = tagTo;
  fee := OpCopy(MkUndefined() ); _ = fee;
  if IsTrue(OpNotEq2(feeCost , MkUndefined() )) {
    fee = MkMap(&VarMap{
        "currency":feeCurrency ,
        "cost":feeCost ,
        "rate":MkUndefined() ,
        });
  }
  return MkMap(&VarMap{
        "info":transaction ,
        "id":id ,
        "txid":txid ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "addressFrom":addressFrom ,
        "addressTo":addressTo ,
        "address":address ,
        "tagFrom":tagFrom ,
        "tagTo":tagTo ,
        "tag":tag ,
        "type":type_ ,
        "amount":amount ,
        "currency":code ,
        "status":status ,
        "updated":MkUndefined() ,
        "fee":fee ,
        });
}

func (this *Bitstamp) ParseTransactionStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "0":MkString("pending") ,
        "1":MkString("pending") ,
        "2":MkString("ok") ,
        "3":MkString("canceled") ,
        "4":MkString("failed") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Bitstamp) ParseOrder(goArgs ...*Variant) *Variant{
  order := GoGetArg(goArgs, 0, MkUndefined()); _ = order;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  id := this.SafeString(order , MkString("id") ); _ = id;
  side := this.SafeString(order , MkString("type") ); _ = side;
  if IsTrue(OpNotEq2(side , MkUndefined() )) {
    side = OpTrinary(OpEq2(side , MkString("1") ), MkString("sell") , MkString("buy") );
  }
  timestamp := this.Parse8601(this.SafeString(order , MkString("datetime") )); _ = timestamp;
  marketId := this.SafeStringLower(order , MkString("currency_pair") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market , MkString("/") ); _ = symbol;
  status := this.ParseOrderStatus(this.SafeString(order , MkString("status") )); _ = status;
  amount := this.SafeNumber(order , MkString("amount") ); _ = amount;
  transactions := this.SafeValue(order , MkString("transactions") , MkArray(&VarArray{})); _ = transactions;
  trades := this.ParseTrades(transactions , market ); _ = trades;
  length := OpCopy(trades.Length ); _ = length;
  if IsTrue(length ) {
    symbol = *(*(trades ).At(MkInteger(0) )).At(MkString("symbol") );
  }
  price := this.SafeNumber(order , MkString("price") ); _ = price;
  return this.SafeOrder(MkMap(&VarMap{
            "id":id ,
            "clientOrderId":MkUndefined() ,
            "datetime":this.Iso8601(timestamp ),
            "timestamp":timestamp ,
            "lastTradeTimestamp":MkUndefined() ,
            "status":status ,
            "symbol":symbol ,
            "type":MkUndefined() ,
            "timeInForce":MkUndefined() ,
            "postOnly":MkUndefined() ,
            "side":side ,
            "price":price ,
            "stopPrice":MkUndefined() ,
            "cost":MkUndefined() ,
            "amount":amount ,
            "filled":MkUndefined() ,
            "remaining":MkUndefined() ,
            "trades":trades ,
            "fee":MkUndefined() ,
            "info":order ,
            "average":MkUndefined() ,
            }));
}

func (this *Bitstamp) ParseLedgerEntryType(goArgs ...*Variant) *Variant{
  type_ := GoGetArg(goArgs, 0, MkUndefined()); _ = type_;
  types := MkMap(&VarMap{
        "0":MkString("transaction") ,
        "1":MkString("transaction") ,
        "2":MkString("trade") ,
        "14":MkString("transfer") ,
        }); _ = types;
  return this.SafeString(types , type_ , type_ );
}

func (this *Bitstamp) ParseLedgerEntry(goArgs ...*Variant) *Variant{
  item := GoGetArg(goArgs, 0, MkUndefined()); _ = item;
  currency := GoGetArg(goArgs, 1, MkUndefined() ); _ = currency;
  type_ := this.ParseLedgerEntryType(this.SafeString(item , MkString("type") )); _ = type_;
  if IsTrue(OpEq2(type_ , MkString("trade") )) {
    parsedTrade := this.ParseTrade(item ); _ = parsedTrade;
    market := OpCopy(MkUndefined() ); _ = market;
    keys := GoGetKeys(item ); _ = keys;
    for i := MkInteger(0) ; IsTrue(OpLw(i , keys.Length )); OpIncr(& i ){
      if IsTrue(OpGtEq((*(keys ).At(i )).Call(MkString("indexOf") , MkString("_") ), MkInteger(0) )) {
        marketId := (*(keys ).At(i )).Call(MkString("replace") , MkString("_") , MkString("") ); _ = marketId;
        if IsTrue(OpHasMember(marketId , *this.At(MkString("markets_by_id")) )) {
          market = *(*this.At(MkString("markets_by_id")) ).At(marketId );
        }
      }
    }
    if IsTrue(OpEq2(market , MkUndefined() )) {
      market = this.GetMarketFromTrade(item );
    }
    direction := OpTrinary(OpEq2(*(parsedTrade ).At(MkString("side") ), MkString("buy") ), MkString("in") , MkString("out") ); _ = direction;
    return MkMap(&VarMap{
          "id":*(parsedTrade ).At(MkString("id") ),
          "info":item ,
          "timestamp":*(parsedTrade ).At(MkString("timestamp") ),
          "datetime":*(parsedTrade ).At(MkString("datetime") ),
          "direction":direction ,
          "account":MkUndefined() ,
          "referenceId":*(parsedTrade ).At(MkString("order") ),
          "referenceAccount":MkUndefined() ,
          "type":type_ ,
          "currency":*(market ).At(MkString("base") ),
          "amount":*(parsedTrade ).At(MkString("amount") ),
          "before":MkUndefined() ,
          "after":MkUndefined() ,
          "status":MkString("ok") ,
          "fee":*(parsedTrade ).At(MkString("fee") ),
          });
  } else {
    parsedTransaction := this.ParseTransaction(item , currency ); _ = parsedTransaction;
    direction := OpCopy(MkUndefined() ); _ = direction;
    if IsTrue(OpHasMember(MkString("amount") , item )) {
      amount := this.SafeNumber(item , MkString("amount") ); _ = amount;
      direction = OpTrinary(OpGt(amount , MkInteger(0) ), MkString("in") , MkString("out") );
    } else {
      if IsTrue(OpAnd(OpHasMember(MkString("currency") , parsedTransaction ), OpNotEq2(*(parsedTransaction ).At(MkString("currency") ), MkUndefined() ))) {
        code := *(parsedTransaction ).At(MkString("currency") ); _ = code;
        currencyId := this.SafeString(*this.At(MkString("currencies_by_id")) , code , code ); _ = currencyId;
        amount := this.SafeNumber(item , currencyId ); _ = amount;
        direction = OpTrinary(OpGt(amount , MkInteger(0) ), MkString("in") , MkString("out") );
      }
    }
    return MkMap(&VarMap{
          "id":*(parsedTransaction ).At(MkString("id") ),
          "info":item ,
          "timestamp":*(parsedTransaction ).At(MkString("timestamp") ),
          "datetime":*(parsedTransaction ).At(MkString("datetime") ),
          "direction":direction ,
          "account":MkUndefined() ,
          "referenceId":*(parsedTransaction ).At(MkString("txid") ),
          "referenceAccount":MkUndefined() ,
          "type":type_ ,
          "currency":*(parsedTransaction ).At(MkString("currency") ),
          "amount":*(parsedTransaction ).At(MkString("amount") ),
          "before":MkUndefined() ,
          "after":MkUndefined() ,
          "status":*(parsedTransaction ).At(MkString("status") ),
          "fee":*(parsedTransaction ).At(MkString("fee") ),
          });
  }
  return MkUndefined()
}

func (this *Bitstamp) FetchLedger(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("privatePostUserTransactions"), this.Extend(request , params )); _ = response;
  currency := OpCopy(MkUndefined() ); _ = currency;
  if IsTrue(OpNotEq2(code , MkUndefined() )) {
    currency = this.Currency(code );
  }
  return this.ParseLedger(response , currency , since , limit );
}

func (this *Bitstamp) FetchOpenOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  market := OpCopy(MkUndefined() ); _ = market;
  this.LoadMarkets();
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
  }
  response := this.Call(MkString("privatePostOpenOrdersAll"), params ); _ = response;
  return this.ParseOrders(response , market , since , limit , MkMap(&VarMap{
            "status":MkString("open") ,
            "type":MkString("limit") ,
            }));
}

func (this *Bitstamp) GetCurrencyName(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  return code.ToLowerCase();
}

func (this *Bitstamp) IsFiat(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  return OpOr(OpEq2(code , MkString("USD") ), OpOr(OpEq2(code , MkString("EUR") ), OpEq2(code , MkString("GBP") )));
}

func (this *Bitstamp) FetchDepositAddress(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  if IsTrue(this.IsFiat(code )) {
    panic(NewNotSupported(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" fiat fetchDepositAddress() for ") , OpAdd(code , MkString(" is not supported!") )))));
  }
  name := this.GetCurrencyName(code ); _ = name;
  method := OpAdd(MkString("privatePost") , OpAdd(this.Capitalize(name ), MkString("Address") )); _ = method;
  response := this.Call(method , params ); _ = response;
  address := this.SafeString(response , MkString("address") ); _ = address;
  tag := this.SafeString2(response , MkString("memo_id") , MkString("destination_tag") ); _ = tag;
  this.CheckAddress(address );
  return MkMap(&VarMap{
        "currency":code ,
        "address":address ,
        "tag":tag ,
        "info":response ,
        });
}

func (this *Bitstamp) Withdraw(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  amount := GoGetArg(goArgs, 1, MkUndefined()); _ = amount;
  address := GoGetArg(goArgs, 2, MkUndefined()); _ = address;
  tag := GoGetArg(goArgs, 3, MkUndefined() ); _ = tag;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  this.CheckAddress(address );
  request := MkMap(&VarMap{"amount":amount }); _ = request;
  method := OpCopy(MkUndefined() ); _ = method;
  if IsTrue(OpNot(this.IsFiat(code ))) {
    name := this.GetCurrencyName(code ); _ = name;
    method = OpAdd(MkString("privatePost") , OpAdd(this.Capitalize(name ), MkString("Withdrawal") ));
    if IsTrue(OpEq2(code , MkString("XRP") )) {
      if IsTrue(OpNotEq2(tag , MkUndefined() )) {
        *(request ).At(MkString("destination_tag") )= OpCopy(tag );
      }
    } else {
      if IsTrue(OpEq2(code , MkString("XLM") )) {
        if IsTrue(OpNotEq2(tag , MkUndefined() )) {
          *(request ).At(MkString("memo_id") )= OpCopy(tag );
        }
      }
    }
    *(request ).At(MkString("address") )= OpCopy(address );
  } else {
    method = MkString("privatePostWithdrawalOpen") ;
    currency := this.Currency(code ); _ = currency;
    *(request ).At(MkString("iban") )= OpCopy(address );
    *(request ).At(MkString("account_currency") )= *(currency ).At(MkString("id") );
  }
  response := this.Call(method , this.Extend(request , params )); _ = response;
  return MkMap(&VarMap{
        "info":response ,
        "id":*(response ).At(MkString("id") ),
        });
}

func (this *Bitstamp) Nonce(goArgs ...*Variant) *Variant{
  return this.Milliseconds();
}

func (this *Bitstamp) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("public") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  url := OpAdd(*(*(*this.At(MkString("urls")) ).At(MkString("api") )).At(api ), MkString("/") ); _ = url;
  OpAddAssign(& url , OpAdd(*this.At(MkString("version")) , MkString("/") ));
  OpAddAssign(& url , this.ImplodeParams(path , params ));
  query := this.Omit(params , this.ExtractParams(path )); _ = query;
  if IsTrue(OpEq2(api , MkString("public") )) {
    if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
      OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(query )));
    }
  } else {
    this.CheckRequiredCredentials();
    xAuth := OpAdd(MkString("BITSTAMP ") , *this.At(MkString("apiKey")) ); _ = xAuth;
    xAuthNonce := this.Uuid(); _ = xAuthNonce;
    xAuthTimestamp := (this.Milliseconds()).Call(MkString("toString") ); _ = xAuthTimestamp;
    xAuthVersion := MkString("v2") ; _ = xAuthVersion;
    contentType := MkString("") ; _ = contentType;
    headers = MkMap(&VarMap{
        "X-Auth":xAuth ,
        "X-Auth-Nonce":xAuthNonce ,
        "X-Auth-Timestamp":xAuthTimestamp ,
        "X-Auth-Version":xAuthVersion ,
        });
    if IsTrue(OpEq2(method , MkString("POST") )) {
      if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
        body = this.Urlencode(query );
        contentType = MkString("application/x-www-form-urlencoded") ;
        *(headers ).At(MkString("Content-Type") )= OpCopy(contentType );
      } else {
        body = this.Urlencode(MkMap(&VarMap{"foo":MkString("bar") }));
        contentType = MkString("application/x-www-form-urlencoded") ;
        *(headers ).At(MkString("Content-Type") )= OpCopy(contentType );
      }
    }
    authBody := OpTrinary(body , body , MkString("") ); _ = authBody;
    auth := OpAdd(xAuth , OpAdd(method , OpAdd(url.Replace(MkString("https://") , MkString("") ), OpAdd(contentType , OpAdd(xAuthNonce , OpAdd(xAuthTimestamp , OpAdd(xAuthVersion , authBody ))))))); _ = auth;
    signature := this.Hmac(this.Encode(auth ), this.Encode(*this.At(MkString("secret")) )); _ = signature;
    *(headers ).At(MkString("X-Auth-Signature") )= OpCopy(signature );
  }
  return MkMap(&VarMap{
        "url":url ,
        "method":method ,
        "body":body ,
        "headers":headers ,
        });
}

func (this *Bitstamp) HandleErrors(goArgs ...*Variant) *Variant{
  httpCode := GoGetArg(goArgs, 0, MkUndefined()); _ = httpCode;
  reason := GoGetArg(goArgs, 1, MkUndefined()); _ = reason;
  url := GoGetArg(goArgs, 2, MkUndefined()); _ = url;
  method := GoGetArg(goArgs, 3, MkUndefined()); _ = method;
  headers := GoGetArg(goArgs, 4, MkUndefined()); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined()); _ = body;
  response := GoGetArg(goArgs, 6, MkUndefined()); _ = response;
  requestHeaders := GoGetArg(goArgs, 7, MkUndefined()); _ = requestHeaders;
  requestBody := GoGetArg(goArgs, 8, MkUndefined()); _ = requestBody;
  if IsTrue(OpEq2(response , MkUndefined() )) {
    return MkUndefined();
  }
  status := this.SafeString(response , MkString("status") ); _ = status;
  error := this.SafeValue(response , MkString("error") ); _ = error;
  if IsTrue(OpOr(OpEq2(status , MkString("error") ), OpNotEq2(error , MkUndefined() ))) {
    errors := MkArray(&VarArray{}); _ = errors;
    if IsTrue(OpEq2(OpType(error ), MkString("string") )) {
      errors.Push(error );
    } else {
      if IsTrue(OpNotEq2(error , MkUndefined() )) {
        keys := GoGetKeys(error ); _ = keys;
        for i := MkInteger(0) ; IsTrue(OpLw(i , keys.Length )); OpIncr(& i ){
          key := *(keys ).At(i ); _ = key;
          value := this.SafeValue(error , key ); _ = value;
          if IsTrue(GoIsArray(value )) {
            errors = this.ArrayConcat(errors , value );
          } else {
            errors.Push(value );
          }
        }
      }
    }
    reason := this.SafeValue(response , MkString("reason") , MkMap(&VarMap{})); _ = reason;
    if IsTrue(OpEq2(OpType(reason ), MkString("string") )) {
      errors.Push(reason );
    } else {
      all := this.SafeValue(reason , MkString("__all__") , MkArray(&VarArray{})); _ = all;
      for i := MkInteger(0) ; IsTrue(OpLw(i , all.Length )); OpIncr(& i ){
        errors.Push(*(all ).At(i ));
      }
    }
    code := this.SafeString(response , MkString("code") ); _ = code;
    if IsTrue(OpEq2(code , MkString("API0005") )) {
      panic(NewAuthenticationError(OpAdd(*this.At(MkString("id")) , MkString(" invalid signature, use the uid for the main account if you have subaccounts") )));
    }
    feedback := OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , body )); _ = feedback;
    for i := MkInteger(0) ; IsTrue(OpLw(i , errors.Length )); OpIncr(& i ){
      value := *(errors ).At(i ); _ = value;
      this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("exact") ), value , feedback );
      this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("broad") ), value , feedback );
    }
    panic(NewExchangeError(feedback ));
  }
  return MkUndefined()
}

