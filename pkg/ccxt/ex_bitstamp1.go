package ccxt

import (
)

type Bitstamp1 struct {
	*ExchangeBase
}

var _ Exchange = (*Bitstamp1)(nil)

func init() {
	exchange := &Bitstamp1{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Bitstamp1) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("bitstamp1") ,
            "name":MkString("Bitstamp") ,
            "countries":MkArray(&VarArray{
                MkString("GB") ,
                }),
            "rateLimit":MkInteger(1000) ,
            "version":MkString("v1") ,
            "has":MkMap(&VarMap{
                "cancelOrder":MkBool(true) ,
                "CORS":MkBool(true) ,
                "createOrder":MkBool(true) ,
                "fetchBalance":MkBool(true) ,
                "fetchMyTrades":MkBool(true) ,
                "fetchOrder":MkBool(false) ,
                "fetchOrderBook":MkBool(true) ,
                "fetchTicker":MkBool(true) ,
                "fetchTrades":MkBool(true) ,
                }),
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/1294454/27786377-8c8ab57e-5fe9-11e7-8ea4-2b05b6bcceec.jpg") ,
                "api":MkString("https://www.bitstamp.net/api") ,
                "www":MkString("https://www.bitstamp.net") ,
                "doc":MkString("https://www.bitstamp.net/api") ,
                }),
            "requiredCredentials":MkMap(&VarMap{
                "apiKey":MkBool(true) ,
                "secret":MkBool(true) ,
                "uid":MkBool(true) ,
                }),
            "api":MkMap(&VarMap{
                "public":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("ticker") ,
                        MkString("ticker_hour") ,
                        MkString("order_book") ,
                        MkString("transactions") ,
                        MkString("eur_usd") ,
                        })}),
                "private":MkMap(&VarMap{"post":MkArray(&VarArray{
                        MkString("balance") ,
                        MkString("user_transactions") ,
                        MkString("open_orders") ,
                        MkString("order_status") ,
                        MkString("cancel_order") ,
                        MkString("cancel_all_orders") ,
                        MkString("buy") ,
                        MkString("sell") ,
                        MkString("bitcoin_deposit_address") ,
                        MkString("unconfirmed_btc") ,
                        MkString("ripple_withdrawal") ,
                        MkString("ripple_address") ,
                        MkString("withdrawal_requests") ,
                        MkString("bitcoin_withdrawal") ,
                        })}),
                }),
            "markets":MkMap(&VarMap{
                "BTC/USD":MkMap(&VarMap{
                    "id":MkString("btcusd") ,
                    "symbol":MkString("BTC/USD") ,
                    "base":MkString("BTC") ,
                    "quote":MkString("USD") ,
                    "baseId":MkString("btc") ,
                    "quoteId":MkString("usd") ,
                    "maker":MkNumber(0.005) ,
                    "taker":MkNumber(0.005) ,
                    }),
                "BTC/EUR":MkMap(&VarMap{
                    "id":MkString("btceur") ,
                    "symbol":MkString("BTC/EUR") ,
                    "base":MkString("BTC") ,
                    "quote":MkString("EUR") ,
                    "baseId":MkString("btc") ,
                    "quoteId":MkString("eur") ,
                    "maker":MkNumber(0.005) ,
                    "taker":MkNumber(0.005) ,
                    }),
                "EUR/USD":MkMap(&VarMap{
                    "id":MkString("eurusd") ,
                    "symbol":MkString("EUR/USD") ,
                    "base":MkString("EUR") ,
                    "quote":MkString("USD") ,
                    "baseId":MkString("eur") ,
                    "quoteId":MkString("usd") ,
                    "maker":MkNumber(0.005) ,
                    "taker":MkNumber(0.005) ,
                    }),
                "XRP/USD":MkMap(&VarMap{
                    "id":MkString("xrpusd") ,
                    "symbol":MkString("XRP/USD") ,
                    "base":MkString("XRP") ,
                    "quote":MkString("USD") ,
                    "baseId":MkString("xrp") ,
                    "quoteId":MkString("usd") ,
                    "maker":MkNumber(0.005) ,
                    "taker":MkNumber(0.005) ,
                    }),
                "XRP/EUR":MkMap(&VarMap{
                    "id":MkString("xrpeur") ,
                    "symbol":MkString("XRP/EUR") ,
                    "base":MkString("XRP") ,
                    "quote":MkString("EUR") ,
                    "baseId":MkString("xrp") ,
                    "quoteId":MkString("eur") ,
                    "maker":MkNumber(0.005) ,
                    "taker":MkNumber(0.005) ,
                    }),
                "XRP/BTC":MkMap(&VarMap{
                    "id":MkString("xrpbtc") ,
                    "symbol":MkString("XRP/BTC") ,
                    "base":MkString("XRP") ,
                    "quote":MkString("BTC") ,
                    "baseId":MkString("xrp") ,
                    "quoteId":MkString("btc") ,
                    "maker":MkNumber(0.005) ,
                    "taker":MkNumber(0.005) ,
                    }),
                "LTC/USD":MkMap(&VarMap{
                    "id":MkString("ltcusd") ,
                    "symbol":MkString("LTC/USD") ,
                    "base":MkString("LTC") ,
                    "quote":MkString("USD") ,
                    "baseId":MkString("ltc") ,
                    "quoteId":MkString("usd") ,
                    "maker":MkNumber(0.005) ,
                    "taker":MkNumber(0.005) ,
                    }),
                "LTC/EUR":MkMap(&VarMap{
                    "id":MkString("ltceur") ,
                    "symbol":MkString("LTC/EUR") ,
                    "base":MkString("LTC") ,
                    "quote":MkString("EUR") ,
                    "baseId":MkString("ltc") ,
                    "quoteId":MkString("eur") ,
                    "maker":MkNumber(0.005) ,
                    "taker":MkNumber(0.005) ,
                    }),
                "LTC/BTC":MkMap(&VarMap{
                    "id":MkString("ltcbtc") ,
                    "symbol":MkString("LTC/BTC") ,
                    "base":MkString("LTC") ,
                    "quote":MkString("BTC") ,
                    "baseId":MkString("ltc") ,
                    "quoteId":MkString("btc") ,
                    "maker":MkNumber(0.005) ,
                    "taker":MkNumber(0.005) ,
                    }),
                "ETH/USD":MkMap(&VarMap{
                    "id":MkString("ethusd") ,
                    "symbol":MkString("ETH/USD") ,
                    "base":MkString("ETH") ,
                    "quote":MkString("USD") ,
                    "baseId":MkString("eth") ,
                    "quoteId":MkString("usd") ,
                    "maker":MkNumber(0.005) ,
                    "taker":MkNumber(0.005) ,
                    }),
                "ETH/EUR":MkMap(&VarMap{
                    "id":MkString("etheur") ,
                    "symbol":MkString("ETH/EUR") ,
                    "base":MkString("ETH") ,
                    "quote":MkString("EUR") ,
                    "baseId":MkString("eth") ,
                    "quoteId":MkString("eur") ,
                    "maker":MkNumber(0.005) ,
                    "taker":MkNumber(0.005) ,
                    }),
                "ETH/BTC":MkMap(&VarMap{
                    "id":MkString("ethbtc") ,
                    "symbol":MkString("ETH/BTC") ,
                    "base":MkString("ETH") ,
                    "quote":MkString("BTC") ,
                    "baseId":MkString("eth") ,
                    "quoteId":MkString("btc") ,
                    "maker":MkNumber(0.005) ,
                    "taker":MkNumber(0.005) ,
                    }),
                }),
            }));
}

func (this *Bitstamp1) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpNotEq2(symbol , MkString("BTC/USD") )) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , OpAdd(*this.At(MkString("version")) , OpAdd(MkString(" fetchOrderBook doesn't support ") , OpAdd(symbol , MkString(", use it for BTC/USD only") )))))));
  }
  this.LoadMarkets();
  orderbook := this.Call(MkString("publicGetOrderBook"), params ); _ = orderbook;
  timestamp := this.SafeTimestamp(orderbook , MkString("timestamp") ); _ = timestamp;
  return this.ParseOrderBook(orderbook , symbol , timestamp );
}

func (this *Bitstamp1) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpNotEq2(symbol , MkString("BTC/USD") )) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , OpAdd(*this.At(MkString("version")) , OpAdd(MkString(" fetchTicker doesn't support ") , OpAdd(symbol , MkString(", use it for BTC/USD only") )))))));
  }
  this.LoadMarkets();
  ticker := this.Call(MkString("publicGetTicker"), params ); _ = ticker;
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

func (this *Bitstamp1) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.SafeTimestamp2(trade , MkString("date") , MkString("datetime") ); _ = timestamp;
  side := OpTrinary(OpEq2(*(trade ).At(MkString("type") ), MkInteger(0) ), MkString("buy") , MkString("sell") ); _ = side;
  orderId := this.SafeString(trade , MkString("order_id") ); _ = orderId;
  if IsTrue(OpHasMember(MkString("currency_pair") , trade )) {
    if IsTrue(OpHasMember(*(trade ).At(MkString("currency_pair") ), *this.At(MkString("markets_by_id")) )) {
      market = *(*this.At(MkString("markets_by_id")) ).At(*(trade ).At(MkString("currency_pair") ));
    }
  }
  id := this.SafeString(trade , MkString("tid") ); _ = id;
  priceString := this.SafeString(trade , MkString("price") ); _ = priceString;
  amountString := this.SafeString(trade , MkString("amount") ); _ = amountString;
  price := this.ParseNumber(priceString ); _ = price;
  amount := this.ParseNumber(amountString ); _ = amount;
  cost := this.ParseNumber(Precise.StringMul(priceString , amountString )); _ = cost;
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  if IsTrue(OpNotEq2(market , MkUndefined() )) {
    symbol = *(market ).At(MkString("symbol") );
  }
  return MkMap(&VarMap{
        "id":id ,
        "info":trade ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "symbol":symbol ,
        "order":orderId ,
        "type":MkUndefined() ,
        "side":side ,
        "takerOrMaker":MkUndefined() ,
        "price":price ,
        "amount":amount ,
        "cost":cost ,
        "fee":MkUndefined() ,
        });
}

func (this *Bitstamp1) FetchTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpNotEq2(symbol , MkString("BTC/USD") )) {
    panic(NewBadSymbol(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , OpAdd(*this.At(MkString("version")) , OpAdd(MkString(" fetchTrades doesn't support ") , OpAdd(symbol , MkString(", use it for BTC/USD only") )))))));
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"time":MkString("minute") }); _ = request;
  response := this.Call(MkString("publicGetTransactions"), this.Extend(request , params )); _ = response;
  return this.ParseTrades(response , market , since , limit );
}

func (this *Bitstamp1) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  balance := this.Call(MkString("privatePostBalance"), params ); _ = balance;
  result := MkMap(&VarMap{"info":balance }); _ = result;
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

func (this *Bitstamp1) CreateOrder(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 2, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 3, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 4, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 5, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpNotEq2(type_ , MkString("limit") )) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , OpAdd(*this.At(MkString("version")) , MkString(" accepts limit orders only") )))));
  }
  if IsTrue(OpNotEq2(symbol , MkString("BTC/USD") )) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , MkString(" v1 supports BTC/USD orders only") )));
  }
  this.LoadMarkets();
  method := OpAdd(MkString("privatePost") , this.Capitalize(side )); _ = method;
  request := MkMap(&VarMap{
        "amount":amount ,
        "price":price ,
        }); _ = request;
  response := this.Call(method , this.Extend(request , params )); _ = response;
  id := this.SafeString(response , MkString("id") ); _ = id;
  return MkMap(&VarMap{
        "info":response ,
        "id":id ,
        });
}

func (this *Bitstamp1) CancelOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  return this.Call(MkString("privatePostCancelOrder"), MkMap(&VarMap{"id":id }));
}

func (this *Bitstamp1) ParseOrderStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "In Queue":MkString("open") ,
        "Open":MkString("open") ,
        "Finished":MkString("closed") ,
        "Canceled":MkString("canceled") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Bitstamp1) FetchOrderStatus(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"id":id }); _ = request;
  response := this.Call(MkString("privatePostOrderStatus"), this.Extend(request , params )); _ = response;
  return this.ParseOrderStatus(response );
}

func (this *Bitstamp1) FetchMyTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
  }
  pair := OpTrinary(market , *(market ).At(MkString("id") ), MkString("all") ); _ = pair;
  request := MkMap(&VarMap{"id":pair }); _ = request;
  response := this.Call(MkString("privatePostOpenOrdersId"), this.Extend(request , params )); _ = response;
  return this.ParseTrades(response , market , since , limit );
}

func (this *Bitstamp1) FetchOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  panic(NewNotSupported(OpAdd(*this.At(MkString("id")) , MkString(" fetchOrder is not implemented yet") )));
  return MkUndefined()
}

func (this *Bitstamp1) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("public") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  url := OpAdd(*(*this.At(MkString("urls")) ).At(MkString("api") ), OpAdd(MkString("/") , this.ImplodeParams(path , params ))); _ = url;
  query := this.Omit(params , this.ExtractParams(path )); _ = query;
  if IsTrue(OpEq2(api , MkString("public") )) {
    if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
      OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(query )));
    }
  } else {
    this.CheckRequiredCredentials();
    nonce := (this.Nonce()).Call(MkString("toString") ); _ = nonce;
    auth := OpAdd(nonce , OpAdd(*this.At(MkString("uid")) , *this.At(MkString("apiKey")) )); _ = auth;
    signature := this.Encode(this.Hmac(this.Encode(auth ), this.Encode(*this.At(MkString("secret")) ))); _ = signature;
    query = this.Extend(MkMap(&VarMap{
            "key":*this.At(MkString("apiKey")) ,
            "signature":signature.ToUpperCase(),
            "nonce":nonce ,
            }), query );
    body = this.Urlencode(query );
    headers = MkMap(&VarMap{"Content-Type":MkString("application/x-www-form-urlencoded") });
  }
  return MkMap(&VarMap{
        "url":url ,
        "method":method ,
        "body":body ,
        "headers":headers ,
        });
}

func (this *Bitstamp1) HandleErrors(goArgs ...*Variant) *Variant{
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
  if IsTrue(OpEq2(status , MkString("error") )) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , this.Json(response )))));
  }
  return MkUndefined()
}

