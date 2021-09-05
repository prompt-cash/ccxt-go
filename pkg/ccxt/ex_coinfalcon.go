package ccxt

import (
)

type Coinfalcon struct {
	*ExchangeBase
}

var _ Exchange = (*Coinfalcon)(nil)

func init() {
	exchange := &Coinfalcon{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Coinfalcon) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("coinfalcon") ,
            "name":MkString("CoinFalcon") ,
            "countries":MkArray(&VarArray{
                MkString("GB") ,
                }),
            "rateLimit":MkInteger(1000) ,
            "version":MkString("v1") ,
            "has":MkMap(&VarMap{
                "cancelOrder":MkBool(true) ,
                "createOrder":MkBool(true) ,
                "fetchBalance":MkBool(true) ,
                "fetchMarkets":MkBool(true) ,
                "fetchMyTrades":MkBool(true) ,
                "fetchOpenOrders":MkBool(true) ,
                "fetchOrder":MkBool(true) ,
                "fetchOrderBook":MkBool(true) ,
                "fetchTicker":MkBool(true) ,
                "fetchTickers":MkBool(true) ,
                "fetchTrades":MkBool(true) ,
                }),
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/1294454/41822275-ed982188-77f5-11e8-92bb-496bcd14ca52.jpg") ,
                "api":MkString("https://coinfalcon.com") ,
                "www":MkString("https://coinfalcon.com") ,
                "doc":MkString("https://docs.coinfalcon.com") ,
                "fees":MkString("https://coinfalcon.com/fees") ,
                "referral":MkString("https://coinfalcon.com/?ref=CFJSVGTUPASB") ,
                }),
            "api":MkMap(&VarMap{
                "public":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("markets") ,
                        MkString("markets/{market}") ,
                        MkString("markets/{market}/orders") ,
                        MkString("markets/{market}/trades") ,
                        })}),
                "private":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("user/accounts") ,
                        MkString("user/orders") ,
                        MkString("user/orders/{id}") ,
                        MkString("user/orders/{id}/trades") ,
                        MkString("user/trades") ,
                        MkString("user/fees") ,
                        MkString("account/withdrawals/{id}") ,
                        MkString("account/withdrawals") ,
                        MkString("account/deposit/{id}") ,
                        MkString("account/deposits") ,
                        MkString("account/deposit_address") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("user/orders") ,
                        MkString("account/withdraw") ,
                        }),
                    "delete":MkArray(&VarArray{
                        MkString("user/orders/{id}") ,
                        MkString("account/withdrawals/{id}") ,
                        }),
                    }),
                }),
            "fees":MkMap(&VarMap{"trading":MkMap(&VarMap{
                    "tierBased":MkBool(true) ,
                    "maker":MkNumber(0.0) ,
                    "taker":MkNumber(0.002) ,
                    })}),
            "precision":MkMap(&VarMap{
                "amount":MkInteger(8) ,
                "price":MkInteger(8) ,
                }),
            }));
}

func (this *Coinfalcon) FetchMarkets(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetMarkets"), params ); _ = response;
  markets := this.SafeValue(response , MkString("data") ); _ = markets;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , markets.Length )); OpIncr(& i ){
    market := *(markets ).At(i ); _ = market;
    baseId := MkUndefined(); _ = baseId;
    quoteId := MkUndefined(); _ = quoteId;
    ArrayUnpack((*(market ).At(MkString("name") )).Call(MkString("split") , MkString("-") ), &baseId, &quoteId);
    base := this.SafeCurrencyCode(baseId ); _ = base;
    quote := this.SafeCurrencyCode(quoteId ); _ = quote;
    symbol := OpAdd(base , OpAdd(MkString("/") , quote )); _ = symbol;
    precision := MkMap(&VarMap{
          "amount":this.SafeInteger(market , MkString("size_precision") ),
          "price":this.SafeInteger(market , MkString("price_precision") ),
          }); _ = precision;
    result.Push(MkMap(&VarMap{
            "id":*(market ).At(MkString("name") ),
            "symbol":symbol ,
            "base":base ,
            "quote":quote ,
            "baseId":baseId ,
            "quoteId":quoteId ,
            "active":MkBool(true) ,
            "precision":precision ,
            "limits":MkMap(&VarMap{
                "amount":MkMap(&VarMap{
                    "min":Math.Pow(MkInteger(10) , OpNeg(*(precision ).At(MkString("amount") ))),
                    "max":MkUndefined() ,
                    }),
                "price":MkMap(&VarMap{
                    "min":Math.Pow(MkInteger(10) , OpNeg(*(precision ).At(MkString("price") ))),
                    "max":MkUndefined() ,
                    }),
                "cost":MkMap(&VarMap{
                    "min":MkUndefined() ,
                    "max":MkUndefined() ,
                    }),
                }),
            "info":market ,
            }));
  }
  return result ;
}

func (this *Coinfalcon) ParseTicker(goArgs ...*Variant) *Variant{
  ticker := GoGetArg(goArgs, 0, MkUndefined()); _ = ticker;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  marketId := this.SafeString(ticker , MkString("name") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market , MkString("-") ); _ = symbol;
  timestamp := this.Milliseconds(); _ = timestamp;
  last := this.SafeNumber(ticker , MkString("last_price") ); _ = last;
  return MkMap(&VarMap{
        "symbol":symbol ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "high":MkUndefined() ,
        "low":MkUndefined() ,
        "bid":MkUndefined() ,
        "bidVolume":MkUndefined() ,
        "ask":MkUndefined() ,
        "askVolume":MkUndefined() ,
        "vwap":MkUndefined() ,
        "open":MkUndefined() ,
        "close":last ,
        "last":last ,
        "previousClose":MkUndefined() ,
        "change":this.SafeNumber(ticker , MkString("change_in_24h") ),
        "percentage":MkUndefined() ,
        "average":MkUndefined() ,
        "baseVolume":MkUndefined() ,
        "quoteVolume":this.SafeNumber(ticker , MkString("volume") ),
        "info":ticker ,
        });
}

func (this *Coinfalcon) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  tickers := this.FetchTickers(params ); _ = tickers;
  return *(tickers ).At(symbol );
}

func (this *Coinfalcon) FetchTickers(goArgs ...*Variant) *Variant{
  symbols := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbols;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("publicGetMarkets"), params ); _ = response;
  tickers := this.SafeValue(response , MkString("data") ); _ = tickers;
  result := MkMap(&VarMap{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , tickers.Length )); OpIncr(& i ){
    ticker := this.ParseTicker(*(tickers ).At(i )); _ = ticker;
    symbol := *(ticker ).At(MkString("symbol") ); _ = symbol;
    *(result ).At(symbol )= OpCopy(ticker );
  }
  return this.FilterByArray(result , MkString("symbol") , symbols );
}

func (this *Coinfalcon) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{
        "market":this.MarketId(symbol ),
        "level":MkString("3") ,
        }); _ = request;
  response := this.Call(MkString("publicGetMarketsMarketOrders"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  return this.ParseOrderBook(data , symbol , MkUndefined() , MkString("bids") , MkString("asks") , MkString("price") , MkString("size") );
}

func (this *Coinfalcon) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.Parse8601(this.SafeString(trade , MkString("created_at") )); _ = timestamp;
  priceString := this.SafeString(trade , MkString("price") ); _ = priceString;
  amountString := this.SafeString(trade , MkString("size") ); _ = amountString;
  price := this.ParseNumber(priceString ); _ = price;
  amount := this.ParseNumber(amountString ); _ = amount;
  cost := this.ParseNumber(Precise.StringMul(priceString , amountString )); _ = cost;
  symbol := *(market ).At(MkString("symbol") ); _ = symbol;
  tradeId := this.SafeString(trade , MkString("id") ); _ = tradeId;
  side := this.SafeString(trade , MkString("side") ); _ = side;
  orderId := this.SafeString(trade , MkString("order_id") ); _ = orderId;
  fee := OpCopy(MkUndefined() ); _ = fee;
  feeCost := this.SafeNumber(trade , MkString("fee") ); _ = feeCost;
  if IsTrue(OpNotEq2(feeCost , MkUndefined() )) {
    feeCurrencyCode := this.SafeString(trade , MkString("fee_currency_code") ); _ = feeCurrencyCode;
    fee = MkMap(&VarMap{
        "cost":feeCost ,
        "currency":this.SafeCurrencyCode(feeCurrencyCode ),
        });
  }
  return MkMap(&VarMap{
        "info":trade ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "symbol":symbol ,
        "id":tradeId ,
        "order":orderId ,
        "type":MkUndefined() ,
        "side":side ,
        "takerOrMaker":MkUndefined() ,
        "price":price ,
        "amount":amount ,
        "cost":cost ,
        "fee":fee ,
        });
}

func (this *Coinfalcon) FetchMyTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchMyTrades() requires a symbol argument") )));
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"market":*(market ).At(MkString("id") )}); _ = request;
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("start_time") )= this.Iso8601(since );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("privateGetUserTrades"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = data;
  return this.ParseTrades(data , market , since , limit );
}

func (this *Coinfalcon) FetchTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"market":*(market ).At(MkString("id") )}); _ = request;
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("since") )= this.Iso8601(since );
  }
  response := this.Call(MkString("publicGetMarketsMarketTrades"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = data;
  return this.ParseTrades(data , market , since , limit );
}

func (this *Coinfalcon) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("privateGetUserAccounts"), params ); _ = response;
  result := MkMap(&VarMap{"info":response }); _ = result;
  balances := this.SafeValue(response , MkString("data") ); _ = balances;
  for i := MkInteger(0) ; IsTrue(OpLw(i , balances.Length )); OpIncr(& i ){
    balance := *(balances ).At(i ); _ = balance;
    currencyId := this.SafeString(balance , MkString("currency_code") ); _ = currencyId;
    code := this.SafeCurrencyCode(currencyId ); _ = code;
    account := this.Account(); _ = account;
    *(account ).At(MkString("free") )= this.SafeString(balance , MkString("available_balance") );
    *(account ).At(MkString("used") )= this.SafeString(balance , MkString("hold_balance") );
    *(account ).At(MkString("total") )= this.SafeString(balance , MkString("balance") );
    *(result ).At(code )= OpCopy(account );
  }
  return this.ParseBalance(result );
}

func (this *Coinfalcon) ParseOrderStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "fulfilled":MkString("closed") ,
        "canceled":MkString("canceled") ,
        "pending":MkString("open") ,
        "open":MkString("open") ,
        "partially_filled":MkString("open") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Coinfalcon) ParseOrder(goArgs ...*Variant) *Variant{
  order := GoGetArg(goArgs, 0, MkUndefined()); _ = order;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  marketId := this.SafeString(order , MkString("market") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market , MkString("-") ); _ = symbol;
  timestamp := this.Parse8601(this.SafeString(order , MkString("created_at") )); _ = timestamp;
  price := this.SafeNumber(order , MkString("price") ); _ = price;
  amount := this.SafeNumber(order , MkString("size") ); _ = amount;
  filled := this.SafeNumber(order , MkString("size_filled") ); _ = filled;
  status := this.ParseOrderStatus(this.SafeString(order , MkString("status") )); _ = status;
  type_ := this.SafeString(order , MkString("operation_type") ); _ = type_;
  if IsTrue(OpNotEq2(type_ , MkUndefined() )) {
    type_ = type_.Split(MkString("_") );
    type_ = *(type_ ).At(MkInteger(0) );
  }
  side := this.SafeString(order , MkString("order_type") ); _ = side;
  postOnly := this.SafeValue(order , MkString("post_only") ); _ = postOnly;
  return this.SafeOrder(MkMap(&VarMap{
            "id":this.SafeString(order , MkString("id") ),
            "clientOrderId":MkUndefined() ,
            "datetime":this.Iso8601(timestamp ),
            "timestamp":timestamp ,
            "status":status ,
            "symbol":symbol ,
            "type":type_ ,
            "timeInForce":MkUndefined() ,
            "postOnly":postOnly ,
            "side":side ,
            "price":price ,
            "stopPrice":MkUndefined() ,
            "cost":MkUndefined() ,
            "amount":amount ,
            "filled":filled ,
            "remaining":MkUndefined() ,
            "trades":MkUndefined() ,
            "fee":MkUndefined() ,
            "info":order ,
            "lastTradeTimestamp":MkUndefined() ,
            "average":MkUndefined() ,
            }));
}

func (this *Coinfalcon) CreateOrder(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 2, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 3, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 4, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 5, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "market":*(market ).At(MkString("id") ),
        "size":this.AmountToPrecision(symbol , amount ),
        "order_type":side ,
        }); _ = request;
  if IsTrue(OpEq2(type_ , MkString("limit") )) {
    price = this.PriceToPrecision(symbol , price );
    *(request ).At(MkString("price") )= price.ToString();
  }
  *(request ).At(MkString("operation_type") )= OpAdd(type_ , MkString("_order") );
  response := this.Call(MkString("privatePostUserOrders"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  return this.ParseOrder(data , market );
}

func (this *Coinfalcon) CancelOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"id":id }); _ = request;
  response := this.Call(MkString("privateDeleteUserOrdersId"), this.Extend(request , params )); _ = response;
  market := this.Market(symbol ); _ = market;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  return this.ParseOrder(data , market );
}

func (this *Coinfalcon) FetchOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"id":id }); _ = request;
  response := this.Call(MkString("privateGetUserOrdersId"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  return this.ParseOrder(data );
}

func (this *Coinfalcon) FetchOpenOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("market") )= *(market ).At(MkString("id") );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("since_time") )= this.Iso8601(since );
  }
  response := this.Call(MkString("privateGetUserOrders"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = data;
  orders := this.FilterByArray(data , MkString("status") , MkArray(&VarArray{
            MkString("pending") ,
            MkString("open") ,
            MkString("partially_filled") ,
            }), MkBool(false) ); _ = orders;
  return this.ParseOrders(orders , market , since , limit );
}

func (this *Coinfalcon) Nonce(goArgs ...*Variant) *Variant{
  return this.Milliseconds();
}

func (this *Coinfalcon) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("public") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  request := OpAdd(MkString("/api/") , OpAdd(*this.At(MkString("version")) , OpAdd(MkString("/") , this.ImplodeParams(path , params )))); _ = request;
  query := this.Omit(params , this.ExtractParams(path )); _ = query;
  if IsTrue(OpEq2(api , MkString("public") )) {
    if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
      OpAddAssign(& request , OpAdd(MkString("?") , this.Urlencode(query )));
    }
  } else {
    this.CheckRequiredCredentials();
    if IsTrue(OpEq2(method , MkString("GET") )) {
      if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
        OpAddAssign(& request , OpAdd(MkString("?") , this.Urlencode(query )));
      }
    } else {
      body = this.Json(query );
    }
    seconds := (this.Seconds()).Call(MkString("toString") ); _ = seconds;
    payload := (MkArray(&VarArray{
            seconds ,
            method ,
            request ,
            })).Call(MkString("join") , MkString("|") ); _ = payload;
    if IsTrue(body ) {
      OpAddAssign(& payload , OpAdd(MkString("|") , body ));
    }
    signature := this.Hmac(this.Encode(payload ), this.Encode(*this.At(MkString("secret")) )); _ = signature;
    headers = MkMap(&VarMap{
        "CF-API-KEY":*this.At(MkString("apiKey")) ,
        "CF-API-TIMESTAMP":seconds ,
        "CF-API-SIGNATURE":signature ,
        "Content-Type":MkString("application/json") ,
        });
  }
  url := OpAdd(*(*this.At(MkString("urls")) ).At(MkString("api") ), request ); _ = url;
  return MkMap(&VarMap{
        "url":url ,
        "method":method ,
        "body":body ,
        "headers":headers ,
        });
}

func (this *Coinfalcon) HandleErrors(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  reason := GoGetArg(goArgs, 1, MkUndefined()); _ = reason;
  url := GoGetArg(goArgs, 2, MkUndefined()); _ = url;
  method := GoGetArg(goArgs, 3, MkUndefined()); _ = method;
  headers := GoGetArg(goArgs, 4, MkUndefined()); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined()); _ = body;
  response := GoGetArg(goArgs, 6, MkUndefined()); _ = response;
  requestHeaders := GoGetArg(goArgs, 7, MkUndefined()); _ = requestHeaders;
  requestBody := GoGetArg(goArgs, 8, MkUndefined()); _ = requestBody;
  if IsTrue(OpLw(code , MkInteger(400) )) {
    return MkUndefined();
  }
  ErrorClass := this.SafeValue(MkMap(&VarMap{
            "401":AuthenticationError ,
            "429":RateLimitExceeded ,
            }), code , ExchangeError ); _ = ErrorClass;
  panic(NewErrorClass(body ));
  return MkUndefined()
}

