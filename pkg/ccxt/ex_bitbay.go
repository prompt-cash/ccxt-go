package ccxt

import (
)

type Bitbay struct {
	*ExchangeBase
}

var _ Exchange = (*Bitbay)(nil)

func init() {
	exchange := &Bitbay{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Bitbay) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("bitbay") ,
            "name":MkString("BitBay") ,
            "countries":MkArray(&VarArray{
                MkString("MT") ,
                MkString("EU") ,
                }),
            "rateLimit":MkInteger(1000) ,
            "has":MkMap(&VarMap{
                "cancelOrder":MkBool(true) ,
                "CORS":MkBool(true) ,
                "createOrder":MkBool(true) ,
                "fetchBalance":MkBool(true) ,
                "fetchLedger":MkBool(true) ,
                "fetchMarkets":MkBool(true) ,
                "fetchMyTrades":MkBool(true) ,
                "fetchOHLCV":MkBool(true) ,
                "fetchOpenOrders":MkBool(true) ,
                "fetchOrderBook":MkBool(true) ,
                "fetchTicker":MkBool(true) ,
                "fetchTrades":MkBool(true) ,
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
                "3d":MkString("259200") ,
                "1w":MkString("604800") ,
                }),
            "hostname":MkString("bitbay.net") ,
            "urls":MkMap(&VarMap{
                "referral":MkString("https://auth.bitbay.net/ref/jHlbB4mIkdS1") ,
                "logo":MkString("https://user-images.githubusercontent.com/1294454/27766132-978a7bd8-5ece-11e7-9540-bc96d1e9bbb8.jpg") ,
                "www":MkString("https://bitbay.net") ,
                "api":MkMap(&VarMap{
                    "public":MkString("https://{hostname}/API/Public") ,
                    "private":MkString("https://{hostname}/API/Trading/tradingApi.php") ,
                    "v1_01Public":MkString("https://api.{hostname}/rest") ,
                    "v1_01Private":MkString("https://api.{hostname}/rest") ,
                    }),
                "doc":MkArray(&VarArray{
                    MkString("https://bitbay.net/public-api") ,
                    MkString("https://bitbay.net/en/private-api") ,
                    MkString("https://bitbay.net/account/tab-api") ,
                    MkString("https://github.com/BitBayNet/API") ,
                    MkString("https://docs.bitbay.net/v1.0.1-en/reference") ,
                    }),
                "support":MkString("https://support.bitbay.net") ,
                "fees":MkString("https://bitbay.net/en/fees") ,
                }),
            "api":MkMap(&VarMap{
                "public":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("{id}/all") ,
                        MkString("{id}/market") ,
                        MkString("{id}/orderbook") ,
                        MkString("{id}/ticker") ,
                        MkString("{id}/trades") ,
                        })}),
                "private":MkMap(&VarMap{"post":MkArray(&VarArray{
                        MkString("info") ,
                        MkString("trade") ,
                        MkString("cancel") ,
                        MkString("orderbook") ,
                        MkString("orders") ,
                        MkString("transfer") ,
                        MkString("withdraw") ,
                        MkString("history") ,
                        MkString("transactions") ,
                        })}),
                "v1_01Public":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("trading/ticker") ,
                        MkString("trading/ticker/{symbol}") ,
                        MkString("trading/stats") ,
                        MkString("trading/orderbook/{symbol}") ,
                        MkString("trading/transactions/{symbol}") ,
                        MkString("trading/candle/history/{symbol}/{resolution}") ,
                        })}),
                "v1_01Private":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("payments/withdrawal/{detailId}") ,
                        MkString("payments/deposit/{detailId}") ,
                        MkString("trading/offer") ,
                        MkString("trading/config/{symbol}") ,
                        MkString("trading/history/transactions") ,
                        MkString("balances/BITBAY/history") ,
                        MkString("balances/BITBAY/balance") ,
                        MkString("fiat_cantor/rate/{baseId}/{quoteId}") ,
                        MkString("fiat_cantor/history") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("trading/offer/{symbol}") ,
                        MkString("trading/config/{symbol}") ,
                        MkString("balances/BITBAY/balance") ,
                        MkString("balances/BITBAY/balance/transfer/{source}/{destination}") ,
                        MkString("fiat_cantor/exchange") ,
                        }),
                    "delete":MkArray(&VarArray{
                        MkString("trading/offer/{symbol}/{id}/{side}/{price}") ,
                        }),
                    "put":MkArray(&VarArray{
                        MkString("balances/BITBAY/balance/{id}") ,
                        }),
                    }),
                }),
            "fees":MkMap(&VarMap{
                "trading":MkMap(&VarMap{
                    "maker":this.ParseNumber(MkString("0.0") ),
                    "taker":this.ParseNumber(MkString("0.001") ),
                    "percentage":MkBool(true) ,
                    "tierBased":MkBool(false) ,
                    }),
                "fiat":MkMap(&VarMap{
                    "maker":OpDiv(MkNumber(0.30) , MkInteger(100) ),
                    "taker":OpDiv(MkNumber(0.43) , MkInteger(100) ),
                    "percentage":MkBool(true) ,
                    "tierBased":MkBool(true) ,
                    "tiers":MkMap(&VarMap{
                        "taker":MkArray(&VarArray{
                            MkArray(&VarArray{
                                MkNumber(0.0043) ,
                                MkInteger(0) ,
                                }),
                            MkArray(&VarArray{
                                MkNumber(0.0042) ,
                                MkInteger(1250) ,
                                }),
                            MkArray(&VarArray{
                                MkNumber(0.0041) ,
                                MkInteger(3750) ,
                                }),
                            MkArray(&VarArray{
                                MkNumber(0.0040) ,
                                MkInteger(7500) ,
                                }),
                            MkArray(&VarArray{
                                MkNumber(0.0039) ,
                                MkInteger(10000) ,
                                }),
                            MkArray(&VarArray{
                                MkNumber(0.0038) ,
                                MkInteger(15000) ,
                                }),
                            MkArray(&VarArray{
                                MkNumber(0.0037) ,
                                MkInteger(20000) ,
                                }),
                            MkArray(&VarArray{
                                MkNumber(0.0036) ,
                                MkInteger(25000) ,
                                }),
                            MkArray(&VarArray{
                                MkNumber(0.0035) ,
                                MkInteger(37500) ,
                                }),
                            MkArray(&VarArray{
                                MkNumber(0.0034) ,
                                MkInteger(50000) ,
                                }),
                            MkArray(&VarArray{
                                MkNumber(0.0033) ,
                                MkInteger(75000) ,
                                }),
                            MkArray(&VarArray{
                                MkNumber(0.0032) ,
                                MkInteger(100000) ,
                                }),
                            MkArray(&VarArray{
                                MkNumber(0.0031) ,
                                MkInteger(150000) ,
                                }),
                            MkArray(&VarArray{
                                MkNumber(0.0030) ,
                                MkInteger(200000) ,
                                }),
                            MkArray(&VarArray{
                                MkNumber(0.0029) ,
                                MkInteger(250000) ,
                                }),
                            MkArray(&VarArray{
                                MkNumber(0.0028) ,
                                MkInteger(375000) ,
                                }),
                            MkArray(&VarArray{
                                MkNumber(0.0027) ,
                                MkInteger(500000) ,
                                }),
                            MkArray(&VarArray{
                                MkNumber(0.0026) ,
                                MkInteger(625000) ,
                                }),
                            MkArray(&VarArray{
                                MkNumber(0.0025) ,
                                MkInteger(875000) ,
                                }),
                            }),
                        "maker":MkArray(&VarArray{
                            MkArray(&VarArray{
                                MkNumber(0.0030) ,
                                MkInteger(0) ,
                                }),
                            MkArray(&VarArray{
                                MkNumber(0.0029) ,
                                MkInteger(1250) ,
                                }),
                            MkArray(&VarArray{
                                MkNumber(0.0028) ,
                                MkInteger(3750) ,
                                }),
                            MkArray(&VarArray{
                                MkNumber(0.0028) ,
                                MkInteger(7500) ,
                                }),
                            MkArray(&VarArray{
                                MkNumber(0.0027) ,
                                MkInteger(10000) ,
                                }),
                            MkArray(&VarArray{
                                MkNumber(0.0026) ,
                                MkInteger(15000) ,
                                }),
                            MkArray(&VarArray{
                                MkNumber(0.0025) ,
                                MkInteger(20000) ,
                                }),
                            MkArray(&VarArray{
                                MkNumber(0.0025) ,
                                MkInteger(25000) ,
                                }),
                            MkArray(&VarArray{
                                MkNumber(0.0024) ,
                                MkInteger(37500) ,
                                }),
                            MkArray(&VarArray{
                                MkNumber(0.0023) ,
                                MkInteger(50000) ,
                                }),
                            MkArray(&VarArray{
                                MkNumber(0.0023) ,
                                MkInteger(75000) ,
                                }),
                            MkArray(&VarArray{
                                MkNumber(0.0022) ,
                                MkInteger(100000) ,
                                }),
                            MkArray(&VarArray{
                                MkNumber(0.0021) ,
                                MkInteger(150000) ,
                                }),
                            MkArray(&VarArray{
                                MkNumber(0.0021) ,
                                MkInteger(200000) ,
                                }),
                            MkArray(&VarArray{
                                MkNumber(0.0020) ,
                                MkInteger(250000) ,
                                }),
                            MkArray(&VarArray{
                                MkNumber(0.0019) ,
                                MkInteger(375000) ,
                                }),
                            MkArray(&VarArray{
                                MkNumber(0.0018) ,
                                MkInteger(500000) ,
                                }),
                            MkArray(&VarArray{
                                MkNumber(0.0018) ,
                                MkInteger(625000) ,
                                }),
                            MkArray(&VarArray{
                                MkNumber(0.0017) ,
                                MkInteger(875000) ,
                                }),
                            }),
                        }),
                    }),
                "funding":MkMap(&VarMap{"withdraw":MkMap(&VarMap{})}),
                }),
            "options":MkMap(&VarMap{"fiatCurrencies":MkArray(&VarArray{
                    MkString("EUR") ,
                    MkString("USD") ,
                    MkString("GBP") ,
                    MkString("PLN") ,
                    })}),
            "exceptions":MkMap(&VarMap{
                "400":ExchangeError ,
                "401":InvalidOrder ,
                "402":InvalidOrder ,
                "403":InvalidOrder ,
                "404":InvalidOrder ,
                "405":InvalidOrder ,
                "406":InsufficientFunds ,
                "408":InvalidOrder ,
                "501":AuthenticationError ,
                "502":AuthenticationError ,
                "503":InvalidNonce ,
                "504":ExchangeError ,
                "505":AuthenticationError ,
                "506":AccountSuspended ,
                "509":ExchangeError ,
                "510":BadSymbol ,
                "FUNDS_NOT_SUFFICIENT":InsufficientFunds ,
                "OFFER_FUNDS_NOT_EXCEEDING_MINIMUMS":InvalidOrder ,
                "OFFER_NOT_FOUND":OrderNotFound ,
                "OFFER_WOULD_HAVE_BEEN_PARTIALLY_FILLED":OrderImmediatelyFillable ,
                "ACTION_LIMIT_EXCEEDED":RateLimitExceeded ,
                "UNDER_MAINTENANCE":OnMaintenance ,
                "REQUEST_TIMESTAMP_TOO_OLD":InvalidNonce ,
                "PERMISSIONS_NOT_SUFFICIENT":PermissionDenied ,
                }),
            "commonCurrencies":MkMap(&VarMap{"GGC":MkString("Global Game Coin") }),
            }));
}

func (this *Bitbay) FetchMarkets(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("v1_01PublicGetTradingTicker"), params ); _ = response;
  fiatCurrencies := this.SafeValue(*this.At(MkString("options")) , MkString("fiatCurrencies") , MkArray(&VarArray{})); _ = fiatCurrencies;
  result := MkArray(&VarArray{}); _ = result;
  items := this.SafeValue(response , MkString("items") ); _ = items;
  keys := GoGetKeys(items ); _ = keys;
  for i := MkInteger(0) ; IsTrue(OpLw(i , keys.Length )); OpIncr(& i ){
    key := *(keys ).At(i ); _ = key;
    item := *(items ).At(key ); _ = item;
    market := this.SafeValue(item , MkString("market") , MkMap(&VarMap{})); _ = market;
    first := this.SafeValue(market , MkString("first") , MkMap(&VarMap{})); _ = first;
    second := this.SafeValue(market , MkString("second") , MkMap(&VarMap{})); _ = second;
    baseId := this.SafeString(first , MkString("currency") ); _ = baseId;
    quoteId := this.SafeString(second , MkString("currency") ); _ = quoteId;
    id := OpAdd(baseId , quoteId ); _ = id;
    base := this.SafeCurrencyCode(baseId ); _ = base;
    quote := this.SafeCurrencyCode(quoteId ); _ = quote;
    symbol := OpAdd(base , OpAdd(MkString("/") , quote )); _ = symbol;
    precision := MkMap(&VarMap{
          "amount":this.SafeInteger(first , MkString("scale") ),
          "price":this.SafeInteger(second , MkString("scale") ),
          }); _ = precision;
    fees := this.SafeValue(*this.At(MkString("fees")) , MkString("trading") , MkMap(&VarMap{})); _ = fees;
    if IsTrue(OpOr(this.InArray(base , fiatCurrencies ), this.InArray(quote , fiatCurrencies ))) {
      fees = this.SafeValue(*this.At(MkString("fees")) , MkString("fiat") , MkMap(&VarMap{}));
    }
    maker := this.SafeNumber(fees , MkString("maker") ); _ = maker;
    taker := this.SafeNumber(fees , MkString("taker") ); _ = taker;
    result.Push(MkMap(&VarMap{
            "id":id ,
            "symbol":symbol ,
            "base":base ,
            "quote":quote ,
            "baseId":baseId ,
            "quoteId":quoteId ,
            "precision":precision ,
            "active":MkUndefined() ,
            "maker":maker ,
            "taker":taker ,
            "limits":MkMap(&VarMap{
                "amount":MkMap(&VarMap{
                    "min":this.SafeNumber(first , MkString("minOffer") ),
                    "max":MkUndefined() ,
                    }),
                "price":MkMap(&VarMap{
                    "min":MkUndefined() ,
                    "max":MkUndefined() ,
                    }),
                "cost":MkMap(&VarMap{
                    "min":this.SafeNumber(second , MkString("minOffer") ),
                    "max":MkUndefined() ,
                    }),
                }),
            "info":item ,
            }));
  }
  return result ;
}

func (this *Bitbay) FetchOpenOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  response := this.Call(MkString("v1_01PrivateGetTradingOffer"), this.Extend(request , params )); _ = response;
  items := this.SafeValue(response , MkString("items") , MkArray(&VarArray{})); _ = items;
  return this.ParseOrders(items , MkUndefined() , since , limit , MkMap(&VarMap{"status":MkString("open") }));
}

func (this *Bitbay) ParseOrder(goArgs ...*Variant) *Variant{
  order := GoGetArg(goArgs, 0, MkUndefined()); _ = order;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  marketId := this.SafeString(order , MkString("market") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market , MkString("-") ); _ = symbol;
  timestamp := this.SafeInteger(order , MkString("time") ); _ = timestamp;
  amount := this.SafeNumber(order , MkString("startAmount") ); _ = amount;
  remaining := this.SafeNumber(order , MkString("currentAmount") ); _ = remaining;
  postOnly := this.SafeValue(order , MkString("postOnly") ); _ = postOnly;
  return this.SafeOrder(MkMap(&VarMap{
            "id":this.SafeString(order , MkString("id") ),
            "clientOrderId":MkUndefined() ,
            "info":order ,
            "timestamp":timestamp ,
            "datetime":this.Iso8601(timestamp ),
            "lastTradeTimestamp":MkUndefined() ,
            "status":MkUndefined() ,
            "symbol":symbol ,
            "type":this.SafeString(order , MkString("mode") ),
            "timeInForce":MkUndefined() ,
            "postOnly":postOnly ,
            "side":this.SafeStringLower(order , MkString("offerType") ),
            "price":this.SafeNumber(order , MkString("rate") ),
            "stopPrice":MkUndefined() ,
            "amount":amount ,
            "cost":MkUndefined() ,
            "filled":MkUndefined() ,
            "remaining":remaining ,
            "average":MkUndefined() ,
            "fee":MkUndefined() ,
            "trades":MkUndefined() ,
            }));
}

func (this *Bitbay) FetchMyTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  if IsTrue(symbol ) {
    markets := MkArray(&VarArray{
          this.MarketId(symbol ),
          }); _ = markets;
    *(request ).At(MkString("markets") )= OpCopy(markets );
  }
  query := MkMap(&VarMap{"query":this.Json(this.Extend(request , params ))}); _ = query;
  response := this.Call(MkString("v1_01PrivateGetTradingHistoryTransactions"), query ); _ = response;
  items := this.SafeValue(response , MkString("items") ); _ = items;
  result := this.ParseTrades(items , MkUndefined() , since , limit ); _ = result;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    return result ;
  }
  return this.FilterBySymbol(result , symbol );
}

func (this *Bitbay) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("v1_01PrivateGetBalancesBITBAYBalance"), params ); _ = response;
  balances := this.SafeValue(response , MkString("balances") ); _ = balances;
  if IsTrue(OpEq2(balances , MkUndefined() )) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" empty balance response ") , this.Json(response )))));
  }
  result := MkMap(&VarMap{"info":response }); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , balances.Length )); OpIncr(& i ){
    balance := *(balances ).At(i ); _ = balance;
    currencyId := this.SafeString(balance , MkString("currency") ); _ = currencyId;
    code := this.SafeCurrencyCode(currencyId ); _ = code;
    account := this.Account(); _ = account;
    *(account ).At(MkString("used") )= this.SafeString(balance , MkString("lockedFunds") );
    *(account ).At(MkString("free") )= this.SafeString(balance , MkString("availableFunds") );
    *(result ).At(code )= OpCopy(account );
  }
  return this.ParseBalance(result );
}

func (this *Bitbay) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"id":this.MarketId(symbol )}); _ = request;
  orderbook := this.Call(MkString("publicGetIdOrderbook"), this.Extend(request , params )); _ = orderbook;
  return this.ParseOrderBook(orderbook , symbol );
}

func (this *Bitbay) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"id":this.MarketId(symbol )}); _ = request;
  ticker := this.Call(MkString("publicGetIdTicker"), this.Extend(request , params )); _ = ticker;
  timestamp := this.Milliseconds(); _ = timestamp;
  baseVolume := this.SafeNumber(ticker , MkString("volume") ); _ = baseVolume;
  vwap := this.SafeNumber(ticker , MkString("vwap") ); _ = vwap;
  quoteVolume := OpCopy(MkUndefined() ); _ = quoteVolume;
  if IsTrue(OpAnd(OpNotEq2(baseVolume , MkUndefined() ), OpNotEq2(vwap , MkUndefined() ))) {
    quoteVolume = OpMulti(baseVolume , vwap );
  }
  last := this.SafeNumber(ticker , MkString("last") ); _ = last;
  return MkMap(&VarMap{
        "symbol":symbol ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "high":this.SafeNumber(ticker , MkString("max") ),
        "low":this.SafeNumber(ticker , MkString("min") ),
        "bid":this.SafeNumber(ticker , MkString("bid") ),
        "bidVolume":MkUndefined() ,
        "ask":this.SafeNumber(ticker , MkString("ask") ),
        "askVolume":MkUndefined() ,
        "vwap":vwap ,
        "open":MkUndefined() ,
        "close":last ,
        "last":last ,
        "previousClose":MkUndefined() ,
        "change":MkUndefined() ,
        "percentage":MkUndefined() ,
        "average":this.SafeNumber(ticker , MkString("average") ),
        "baseVolume":baseVolume ,
        "quoteVolume":quoteVolume ,
        "info":ticker ,
        });
}

func (this *Bitbay) FetchLedger(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  balanceCurrencies := MkArray(&VarArray{}); _ = balanceCurrencies;
  if IsTrue(OpNotEq2(code , MkUndefined() )) {
    currency := this.Currency(code ); _ = currency;
    balanceCurrencies.Push(*(currency ).At(MkString("id") ));
  }
  request := MkMap(&VarMap{"balanceCurrencies":balanceCurrencies }); _ = request;
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("fromTime") )= OpCopy(since );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  request = this.Extend(request , params );
  response := this.Call(MkString("v1_01PrivateGetBalancesBITBAYHistory"), MkMap(&VarMap{"query":this.Json(request )})); _ = response;
  items := *(response ).At(MkString("items") ); _ = items;
  return this.ParseLedger(items , MkUndefined() , since , limit );
}

func (this *Bitbay) ParseLedgerEntry(goArgs ...*Variant) *Variant{
  item := GoGetArg(goArgs, 0, MkUndefined()); _ = item;
  currency := GoGetArg(goArgs, 1, MkUndefined() ); _ = currency;
  timestamp := this.SafeInteger(item , MkString("time") ); _ = timestamp;
  balance := this.SafeValue(item , MkString("balance") , MkMap(&VarMap{})); _ = balance;
  currencyId := this.SafeString(balance , MkString("currency") ); _ = currencyId;
  code := this.SafeCurrencyCode(currencyId ); _ = code;
  change := this.SafeValue(item , MkString("change") , MkMap(&VarMap{})); _ = change;
  amount := this.SafeNumber(change , MkString("total") ); _ = amount;
  direction := MkString("in") ; _ = direction;
  if IsTrue(OpLw(amount , MkInteger(0) )) {
    direction = MkString("out") ;
    amount = OpNeg(amount );
  }
  id := this.SafeString(item , MkString("historyId") ); _ = id;
  referenceId := this.SafeString(item , MkString("detailId") ); _ = referenceId;
  type_ := this.ParseLedgerEntryType(this.SafeString(item , MkString("type") )); _ = type_;
  fundsBefore := this.SafeValue(item , MkString("fundsBefore") , MkMap(&VarMap{})); _ = fundsBefore;
  before := this.SafeNumber(fundsBefore , MkString("total") ); _ = before;
  fundsAfter := this.SafeValue(item , MkString("fundsAfter") , MkMap(&VarMap{})); _ = fundsAfter;
  after := this.SafeNumber(fundsAfter , MkString("total") ); _ = after;
  return MkMap(&VarMap{
        "info":item ,
        "id":id ,
        "direction":direction ,
        "account":MkUndefined() ,
        "referenceId":referenceId ,
        "referenceAccount":MkUndefined() ,
        "type":type_ ,
        "currency":code ,
        "amount":amount ,
        "before":before ,
        "after":after ,
        "status":MkString("ok") ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "fee":MkUndefined() ,
        });
}

func (this *Bitbay) ParseLedgerEntryType(goArgs ...*Variant) *Variant{
  type_ := GoGetArg(goArgs, 0, MkUndefined()); _ = type_;
  types := MkMap(&VarMap{
        "ADD_FUNDS":MkString("transaction") ,
        "BITCOIN_GOLD_FORK":MkString("transaction") ,
        "CREATE_BALANCE":MkString("transaction") ,
        "FUNDS_MIGRATION":MkString("transaction") ,
        "WITHDRAWAL_LOCK_FUNDS":MkString("transaction") ,
        "WITHDRAWAL_SUBTRACT_FUNDS":MkString("transaction") ,
        "WITHDRAWAL_UNLOCK_FUNDS":MkString("transaction") ,
        "TRANSACTION_COMMISSION_OUTCOME":MkString("fee") ,
        "TRANSACTION_COMMISSION_RETURN":MkString("fee") ,
        "TRANSACTION_OFFER_ABORTED_RETURN":MkString("trade") ,
        "TRANSACTION_OFFER_COMPLETED_RETURN":MkString("trade") ,
        "TRANSACTION_POST_INCOME":MkString("trade") ,
        "TRANSACTION_POST_OUTCOME":MkString("trade") ,
        "TRANSACTION_PRE_LOCKING":MkString("trade") ,
        }); _ = types;
  return this.SafeString(types , type_ , type_ );
}

func (this *Bitbay) ParseOHLCV(goArgs ...*Variant) *Variant{
  ohlcv := GoGetArg(goArgs, 0, MkUndefined()); _ = ohlcv;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  first := this.SafeValue(ohlcv , MkInteger(1) , MkMap(&VarMap{})); _ = first;
  return MkArray(&VarArray{
        this.SafeInteger(ohlcv , MkInteger(0) ),
        this.SafeNumber(first , MkString("o") ),
        this.SafeNumber(first , MkString("h") ),
        this.SafeNumber(first , MkString("l") ),
        this.SafeNumber(first , MkString("c") ),
        this.SafeNumber(first , MkString("v") ),
        });
}

func (this *Bitbay) FetchOHLCV(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  timeframe := GoGetArg(goArgs, 1, MkString("1m") ); _ = timeframe;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  tradingSymbol := OpAdd(*(market ).At(MkString("baseId") ), OpAdd(MkString("-") , *(market ).At(MkString("quoteId") ))); _ = tradingSymbol;
  request := MkMap(&VarMap{
        "symbol":tradingSymbol ,
        "resolution":*(*this.At(MkString("timeframes")) ).At(timeframe ),
        }); _ = request;
  if IsTrue(OpEq2(limit , MkUndefined() )) {
    limit = MkInteger(100) ;
  }
  duration := this.ParseTimeframe(timeframe ); _ = duration;
  timerange := OpMulti(limit , OpMulti(duration , MkInteger(1000) )); _ = timerange;
  if IsTrue(OpEq2(since , MkUndefined() )) {
    *(request ).At(MkString("to") )= this.Milliseconds();
    *(request ).At(MkString("from") )= OpSub(*(request ).At(MkString("to") ), timerange );
  } else {
    *(request ).At(MkString("from") )= parseInt(since );
    *(request ).At(MkString("to") )= this.Sum(*(request ).At(MkString("from") ), timerange );
  }
  response := this.Call(MkString("v1_01PublicGetTradingCandleHistorySymbolResolution"), this.Extend(request , params )); _ = response;
  items := this.SafeValue(response , MkString("items") , MkArray(&VarArray{})); _ = items;
  return this.ParseOHLCVs(items , market , timeframe , since , limit );
}

func (this *Bitbay) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.SafeInteger2(trade , MkString("time") , MkString("t") ); _ = timestamp;
  side := this.SafeStringLower2(trade , MkString("userAction") , MkString("ty") ); _ = side;
  wasTaker := this.SafeValue(trade , MkString("wasTaker") ); _ = wasTaker;
  takerOrMaker := OpCopy(MkUndefined() ); _ = takerOrMaker;
  if IsTrue(OpNotEq2(wasTaker , MkUndefined() )) {
    takerOrMaker = OpTrinary(wasTaker , MkString("taker") , MkString("maker") );
  }
  priceString := this.SafeString2(trade , MkString("rate") , MkString("r") ); _ = priceString;
  amountString := this.SafeString2(trade , MkString("amount") , MkString("a") ); _ = amountString;
  price := this.ParseNumber(priceString ); _ = price;
  amount := this.ParseNumber(amountString ); _ = amount;
  cost := this.ParseNumber(Precise.StringMul(priceString , amountString )); _ = cost;
  feeCost := this.SafeNumber(trade , MkString("commissionValue") ); _ = feeCost;
  marketId := this.SafeString(trade , MkString("market") ); _ = marketId;
  market = this.SafeMarket(marketId , market , MkString("-") );
  symbol := *(market ).At(MkString("symbol") ); _ = symbol;
  fee := OpCopy(MkUndefined() ); _ = fee;
  if IsTrue(OpNotEq2(feeCost , MkUndefined() )) {
    feeCcy := OpTrinary(OpEq2(side , MkString("buy") ), *(market ).At(MkString("base") ), *(market ).At(MkString("quote") )); _ = feeCcy;
    fee = MkMap(&VarMap{
        "currency":feeCcy ,
        "cost":feeCost ,
        });
  }
  order := this.SafeString(trade , MkString("offerId") ); _ = order;
  type_ := OpCopy(MkUndefined() ); _ = type_;
  if IsTrue(OpNotEq2(order , MkUndefined() )) {
    type_ = OpTrinary(order , MkString("limit") , MkString("market") );
  }
  return MkMap(&VarMap{
        "id":this.SafeString(trade , MkString("id") ),
        "order":order ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "symbol":symbol ,
        "type":type_ ,
        "side":side ,
        "price":price ,
        "amount":amount ,
        "cost":cost ,
        "takerOrMaker":takerOrMaker ,
        "fee":fee ,
        "info":trade ,
        });
}

func (this *Bitbay) FetchTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  tradingSymbol := OpAdd(*(market ).At(MkString("baseId") ), OpAdd(MkString("-") , *(market ).At(MkString("quoteId") ))); _ = tradingSymbol;
  request := MkMap(&VarMap{"symbol":tradingSymbol }); _ = request;
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("fromTime") )= OpSub(since , MkInteger(1) );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("v1_01PublicGetTradingTransactionsSymbol"), this.Extend(request , params )); _ = response;
  items := this.SafeValue(response , MkString("items") ); _ = items;
  return this.ParseTrades(items , market , since , limit );
}

func (this *Bitbay) CreateOrder(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 2, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 3, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 4, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 5, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  tradingSymbol := OpAdd(*(market ).At(MkString("baseId") ), OpAdd(MkString("-") , *(market ).At(MkString("quoteId") ))); _ = tradingSymbol;
  request := MkMap(&VarMap{
        "symbol":tradingSymbol ,
        "offerType":side ,
        "amount":amount ,
        "mode":type_ ,
        }); _ = request;
  if IsTrue(OpEq2(type_ , MkString("limit") )) {
    *(request ).At(MkString("rate") )= OpCopy(price );
    price = parseFloat(price );
  }
  amount = parseFloat(amount );
  response := this.Call(MkString("v1_01PrivatePostTradingOfferSymbol"), this.Extend(request , params )); _ = response;
  timestamp := this.Milliseconds(); _ = timestamp;
  id := this.SafeString(response , MkString("offerId") ); _ = id;
  completed := this.SafeValue(response , MkString("completed") , MkBool(false) ); _ = completed;
  status := OpTrinary(completed , MkString("closed") , MkString("open") ); _ = status;
  filled := MkInteger(0) ; _ = filled;
  cost := OpCopy(MkUndefined() ); _ = cost;
  transactions := this.SafeValue(response , MkString("transactions") ); _ = transactions;
  trades := OpCopy(MkUndefined() ); _ = trades;
  if IsTrue(OpNotEq2(transactions , MkUndefined() )) {
    trades = this.ParseTrades(transactions , market , MkUndefined() , MkUndefined() , MkMap(&VarMap{
            "timestamp":timestamp ,
            "datetime":this.Iso8601(timestamp ),
            "symbol":symbol ,
            "side":side ,
            "type":type_ ,
            "orderId":id ,
            }));
    cost = MkInteger(0) ;
    for i := MkInteger(0) ; IsTrue(OpLw(i , trades.Length )); OpIncr(& i ){
      filled = this.Sum(filled , *(*(trades ).At(i )).At(MkString("amount") ));
      cost = this.Sum(cost , *(*(trades ).At(i )).At(MkString("cost") ));
    }
  }
  remaining := OpSub(amount , filled ); _ = remaining;
  return MkMap(&VarMap{
        "id":id ,
        "info":response ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "lastTradeTimestamp":MkUndefined() ,
        "status":status ,
        "symbol":symbol ,
        "type":type_ ,
        "side":side ,
        "price":price ,
        "amount":amount ,
        "cost":cost ,
        "filled":filled ,
        "remaining":remaining ,
        "average":MkUndefined() ,
        "fee":MkUndefined() ,
        "trades":trades ,
        "clientOrderId":MkUndefined() ,
        });
}

func (this *Bitbay) CancelOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  side := this.SafeString(params , MkString("side") ); _ = side;
  if IsTrue(OpEq2(side , MkUndefined() )) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , MkString(" cancelOrder() requires a `side` parameter (\"buy\" or \"sell\")") )));
  }
  price := this.SafeValue(params , MkString("price") ); _ = price;
  if IsTrue(OpEq2(price , MkUndefined() )) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , MkString(" cancelOrder() requires a `price` parameter (float or string)") )));
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  tradingSymbol := OpAdd(*(market ).At(MkString("baseId") ), OpAdd(MkString("-") , *(market ).At(MkString("quoteId") ))); _ = tradingSymbol;
  request := MkMap(&VarMap{
        "symbol":tradingSymbol ,
        "id":id ,
        "side":side ,
        "price":price ,
        }); _ = request;
  return this.Call(MkString("v1_01PrivateDeleteTradingOfferSymbolIdSidePrice"), this.Extend(request , params ));
}

func (this *Bitbay) IsFiat(goArgs ...*Variant) *Variant{
  currency := GoGetArg(goArgs, 0, MkUndefined()); _ = currency;
  fiatCurrencies := MkMap(&VarMap{
        "USD":MkBool(true) ,
        "EUR":MkBool(true) ,
        "PLN":MkBool(true) ,
        }); _ = fiatCurrencies;
  return this.SafeValue(fiatCurrencies , currency , MkBool(false) );
}

func (this *Bitbay) Withdraw(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  amount := GoGetArg(goArgs, 1, MkUndefined()); _ = amount;
  address := GoGetArg(goArgs, 2, MkUndefined()); _ = address;
  tag := GoGetArg(goArgs, 3, MkUndefined() ); _ = tag;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.CheckAddress(address );
  this.LoadMarkets();
  method := OpCopy(MkUndefined() ); _ = method;
  currency := this.Currency(code ); _ = currency;
  request := MkMap(&VarMap{
        "currency":*(currency ).At(MkString("id") ),
        "quantity":amount ,
        }); _ = request;
  if IsTrue(this.IsFiat(code )) {
    method = MkString("privatePostWithdraw") ;
  } else {
    method = MkString("privatePostTransfer") ;
    if IsTrue(OpNotEq2(tag , MkUndefined() )) {
      OpAddAssign(& address , OpAdd(MkString("?dt=") , tag.ToString()));
    }
    *(request ).At(MkString("address") )= OpCopy(address );
  }
  response := this.Call(method , this.Extend(request , params )); _ = response;
  return MkMap(&VarMap{
        "info":response ,
        "id":MkUndefined() ,
        });
}

func (this *Bitbay) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("public") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  url := this.ImplodeHostname(*(*(*this.At(MkString("urls")) ).At(MkString("api") )).At(api )); _ = url;
  if IsTrue(OpEq2(api , MkString("public") )) {
    query := this.Omit(params , this.ExtractParams(path )); _ = query;
    OpAddAssign(& url , OpAdd(MkString("/") , OpAdd(this.ImplodeParams(path , params ), MkString(".json") )));
    if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
      OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(query )));
    }
  } else {
    if IsTrue(OpEq2(api , MkString("v1_01Public") )) {
      query := this.Omit(params , this.ExtractParams(path )); _ = query;
      OpAddAssign(& url , OpAdd(MkString("/") , this.ImplodeParams(path , params )));
      if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
        OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(query )));
      }
    } else {
      if IsTrue(OpEq2(api , MkString("v1_01Private") )) {
        this.CheckRequiredCredentials();
        query := this.Omit(params , this.ExtractParams(path )); _ = query;
        OpAddAssign(& url , OpAdd(MkString("/") , this.ImplodeParams(path , params )));
        nonce := (this.Milliseconds()).Call(MkString("toString") ); _ = nonce;
        payload := OpCopy(MkUndefined() ); _ = payload;
        if IsTrue(OpNotEq2(method , MkString("POST") )) {
          if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
            OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(query )));
          }
          payload = OpAdd(*this.At(MkString("apiKey")) , nonce );
        } else {
          if IsTrue(OpEq2(body , MkUndefined() )) {
            body = this.Json(query );
            payload = OpAdd(*this.At(MkString("apiKey")) , OpAdd(nonce , body ));
          }
        }
        headers = MkMap(&VarMap{
            "Request-Timestamp":nonce ,
            "Operation-Id":this.Uuid(),
            "API-Key":*this.At(MkString("apiKey")) ,
            "API-Hash":this.Hmac(this.Encode(payload ), this.Encode(*this.At(MkString("secret")) ), MkString("sha512") ),
            "Content-Type":MkString("application/json") ,
            });
      } else {
        this.CheckRequiredCredentials();
        body = this.Urlencode(this.Extend(MkMap(&VarMap{
                    "method":path ,
                    "moment":this.Nonce(),
                    }), params ));
        headers = MkMap(&VarMap{
            "Content-Type":MkString("application/x-www-form-urlencoded") ,
            "API-Key":*this.At(MkString("apiKey")) ,
            "API-Hash":this.Hmac(this.Encode(body ), this.Encode(*this.At(MkString("secret")) ), MkString("sha512") ),
            });
      }
    }
  }
  return MkMap(&VarMap{
        "url":url ,
        "method":method ,
        "body":body ,
        "headers":headers ,
        });
}

func (this *Bitbay) HandleErrors(goArgs ...*Variant) *Variant{
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
  if IsTrue(OpHasMember(MkString("code") , response )) {
    code := this.SafeString(response , MkString("code") ); _ = code;
    feedback := OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , body )); _ = feedback;
    this.ThrowExactlyMatchedException(*this.At(MkString("exceptions")) , code , feedback );
    panic(NewExchangeError(feedback ));
  } else {
    if IsTrue(OpHasMember(MkString("status") , response )) {
      status := this.SafeString(response , MkString("status") ); _ = status;
      if IsTrue(OpEq2(status , MkString("Fail") )) {
        errors := this.SafeValue(response , MkString("errors") ); _ = errors;
        feedback := OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , body )); _ = feedback;
        for i := MkInteger(0) ; IsTrue(OpLw(i , errors.Length )); OpIncr(& i ){
          error := *(errors ).At(i ); _ = error;
          this.ThrowExactlyMatchedException(*this.At(MkString("exceptions")) , error , feedback );
        }
        panic(NewExchangeError(feedback ));
      }
    }
  }
  return MkUndefined()
}

