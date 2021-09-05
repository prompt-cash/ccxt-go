package ccxt

import (
)

type Coincheck struct {
	*ExchangeBase
}

var _ Exchange = (*Coincheck)(nil)

func init() {
	exchange := &Coincheck{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Coincheck) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("coincheck") ,
            "name":MkString("coincheck") ,
            "countries":MkArray(&VarArray{
                MkString("JP") ,
                MkString("ID") ,
                }),
            "rateLimit":MkInteger(1500) ,
            "has":MkMap(&VarMap{
                "cancelOrder":MkBool(true) ,
                "CORS":MkBool(false) ,
                "createOrder":MkBool(true) ,
                "fetchBalance":MkBool(true) ,
                "fetchMyTrades":MkBool(true) ,
                "fetchOrderBook":MkBool(true) ,
                "fetchOpenOrders":MkBool(true) ,
                "fetchTicker":MkBool(true) ,
                "fetchTrades":MkBool(true) ,
                }),
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/51840849/87182088-1d6d6380-c2ec-11ea-9c64-8ab9f9b289f5.jpg") ,
                "api":MkString("https://coincheck.com/api") ,
                "www":MkString("https://coincheck.com") ,
                "doc":MkString("https://coincheck.com/documents/exchange/api") ,
                "fees":MkArray(&VarArray{
                    MkString("https://coincheck.com/exchange/fee") ,
                    MkString("https://coincheck.com/info/fee") ,
                    }),
                }),
            "api":MkMap(&VarMap{
                "public":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("exchange/orders/rate") ,
                        MkString("order_books") ,
                        MkString("rate/{pair}") ,
                        MkString("ticker") ,
                        MkString("trades") ,
                        })}),
                "private":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("accounts") ,
                        MkString("accounts/balance") ,
                        MkString("accounts/leverage_balance") ,
                        MkString("bank_accounts") ,
                        MkString("deposit_money") ,
                        MkString("exchange/orders/opens") ,
                        MkString("exchange/orders/transactions") ,
                        MkString("exchange/orders/transactions_pagination") ,
                        MkString("exchange/leverage/positions") ,
                        MkString("lending/borrows/matches") ,
                        MkString("send_money") ,
                        MkString("withdraws") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("bank_accounts") ,
                        MkString("deposit_money/{id}/fast") ,
                        MkString("exchange/orders") ,
                        MkString("exchange/transfers/to_leverage") ,
                        MkString("exchange/transfers/from_leverage") ,
                        MkString("lending/borrows") ,
                        MkString("lending/borrows/{id}/repay") ,
                        MkString("send_money") ,
                        MkString("withdraws") ,
                        }),
                    "delete":MkArray(&VarArray{
                        MkString("bank_accounts/{id}") ,
                        MkString("exchange/orders/{id}") ,
                        MkString("withdraws/{id}") ,
                        }),
                    }),
                }),
            "markets":MkMap(&VarMap{
                "BTC/JPY":MkMap(&VarMap{
                    "id":MkString("btc_jpy") ,
                    "symbol":MkString("BTC/JPY") ,
                    "base":MkString("BTC") ,
                    "quote":MkString("JPY") ,
                    "baseId":MkString("btc") ,
                    "quoteId":MkString("jpy") ,
                    }),
                "ETC/JPY":MkMap(&VarMap{
                    "id":MkString("etc_jpy") ,
                    "symbol":MkString("ETC/JPY") ,
                    "base":MkString("ETC") ,
                    "quote":MkString("JPY") ,
                    "baseId":MkString("etc") ,
                    "quoteId":MkString("jpy") ,
                    }),
                "FCT/JPY":MkMap(&VarMap{
                    "id":MkString("fct_jpy") ,
                    "symbol":MkString("FCT/JPY") ,
                    "base":MkString("FCT") ,
                    "quote":MkString("JPY") ,
                    "baseId":MkString("fct") ,
                    "quoteId":MkString("jpy") ,
                    }),
                "MONA/JPY":MkMap(&VarMap{
                    "id":MkString("mona_jpy") ,
                    "symbol":MkString("MONA/JPY") ,
                    "base":MkString("MONA") ,
                    "quote":MkString("JPY") ,
                    "baseId":MkString("mona") ,
                    "quoteId":MkString("jpy") ,
                    }),
                "ETC/BTC":MkMap(&VarMap{
                    "id":MkString("etc_btc") ,
                    "symbol":MkString("ETC/BTC") ,
                    "base":MkString("ETC") ,
                    "quote":MkString("BTC") ,
                    "baseId":MkString("etc") ,
                    "quoteId":MkString("btc") ,
                    }),
                }),
            "fees":MkMap(&VarMap{"trading":MkMap(&VarMap{
                    "tierBased":MkBool(false) ,
                    "percentage":MkBool(true) ,
                    "maker":this.ParseNumber(MkString("0") ),
                    "taker":this.ParseNumber(MkString("0") ),
                    })}),
            }));
}

func (this *Coincheck) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  balances := this.Call(MkString("privateGetAccountsBalance"), params ); _ = balances;
  result := MkMap(&VarMap{"info":balances }); _ = result;
  codes := GoGetKeys(*this.At(MkString("currencies")) ); _ = codes;
  for i := MkInteger(0) ; IsTrue(OpLw(i , codes.Length )); OpIncr(& i ){
    code := *(codes ).At(i ); _ = code;
    currency := this.Currency(code ); _ = currency;
    currencyId := *(currency ).At(MkString("id") ); _ = currencyId;
    if IsTrue(OpHasMember(currencyId , balances )) {
      account := this.Account(); _ = account;
      reserved := OpAdd(currencyId , MkString("_reserved") ); _ = reserved;
      *(account ).At(MkString("free") )= this.SafeString(balances , currencyId );
      *(account ).At(MkString("used") )= this.SafeString(balances , reserved );
      *(result ).At(code )= OpCopy(account );
    }
  }
  return this.ParseBalance(result );
}

func (this *Coincheck) FetchOpenOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
  }
  response := this.Call(MkString("privateGetExchangeOrdersOpens"), params ); _ = response;
  rawOrders := this.SafeValue(response , MkString("orders") , MkArray(&VarArray{})); _ = rawOrders;
  parsedOrders := this.ParseOrders(rawOrders , market , since , limit ); _ = parsedOrders;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , parsedOrders.Length )); OpIncr(& i ){
    result.Push(this.Extend(*(parsedOrders ).At(i ), MkMap(&VarMap{"status":MkString("open") })));
  }
  return result ;
}

func (this *Coincheck) ParseOrder(goArgs ...*Variant) *Variant{
  order := GoGetArg(goArgs, 0, MkUndefined()); _ = order;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  id := this.SafeString(order , MkString("id") ); _ = id;
  side := this.SafeString(order , MkString("order_type") ); _ = side;
  timestamp := this.Parse8601(this.SafeString(order , MkString("created_at") )); _ = timestamp;
  amount := this.SafeNumber(order , MkString("pending_amount") ); _ = amount;
  remaining := this.SafeNumber(order , MkString("pending_amount") ); _ = remaining;
  price := this.SafeNumber(order , MkString("rate") ); _ = price;
  status := OpCopy(MkUndefined() ); _ = status;
  marketId := this.SafeString(order , MkString("pair") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market , MkString("_") ); _ = symbol;
  return this.SafeOrder(MkMap(&VarMap{
            "id":id ,
            "clientOrderId":MkUndefined() ,
            "timestamp":timestamp ,
            "datetime":this.Iso8601(timestamp ),
            "lastTradeTimestamp":MkUndefined() ,
            "amount":amount ,
            "remaining":remaining ,
            "filled":MkUndefined() ,
            "side":side ,
            "type":MkUndefined() ,
            "timeInForce":MkUndefined() ,
            "postOnly":MkUndefined() ,
            "status":status ,
            "symbol":symbol ,
            "price":price ,
            "stopPrice":MkUndefined() ,
            "cost":MkUndefined() ,
            "fee":MkUndefined() ,
            "info":order ,
            "average":MkUndefined() ,
            "trades":MkUndefined() ,
            }));
}

func (this *Coincheck) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"pair":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("publicGetOrderBooks"), this.Extend(request , params )); _ = response;
  return this.ParseOrderBook(response , symbol );
}

func (this *Coincheck) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpNotEq2(symbol , MkString("BTC/JPY") )) {
    panic(NewBadSymbol(OpAdd(*this.At(MkString("id")) , MkString(" fetchTicker () supports BTC/JPY only") )));
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"pair":*(market ).At(MkString("id") )}); _ = request;
  ticker := this.Call(MkString("publicGetTicker"), this.Extend(request , params )); _ = ticker;
  timestamp := this.SafeTimestamp(ticker , MkString("timestamp") ); _ = timestamp;
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
        "vwap":MkUndefined() ,
        "open":MkUndefined() ,
        "close":last ,
        "last":last ,
        "previousClose":MkUndefined() ,
        "change":MkUndefined() ,
        "percentage":MkUndefined() ,
        "average":MkUndefined() ,
        "baseVolume":this.SafeNumber(ticker , MkString("volume") ),
        "quoteVolume":MkUndefined() ,
        "info":ticker ,
        });
}

func (this *Coincheck) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.Parse8601(this.SafeString(trade , MkString("created_at") )); _ = timestamp;
  id := this.SafeString(trade , MkString("id") ); _ = id;
  priceString := this.SafeString(trade , MkString("rate") ); _ = priceString;
  marketId := this.SafeString(trade , MkString("pair") ); _ = marketId;
  market = this.SafeValue(*this.At(MkString("markets_by_id")) , marketId , market );
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  baseId := OpCopy(MkUndefined() ); _ = baseId;
  quoteId := OpCopy(MkUndefined() ); _ = quoteId;
  if IsTrue(OpNotEq2(marketId , MkUndefined() )) {
    if IsTrue(OpHasMember(marketId , *this.At(MkString("markets_by_id")) )) {
      market = *(*this.At(MkString("markets_by_id")) ).At(marketId );
      baseId = *(market ).At(MkString("baseId") );
      quoteId = *(market ).At(MkString("quoteId") );
      symbol = *(market ).At(MkString("symbol") );
    } else {
      ids := marketId.Split(MkString("_") ); _ = ids;
      baseId = *(ids ).At(MkInteger(0) );
      quoteId = *(ids ).At(MkInteger(1) );
      base := this.SafeCurrencyCode(baseId ); _ = base;
      quote := this.SafeCurrencyCode(quoteId ); _ = quote;
      symbol = OpAdd(base , OpAdd(MkString("/") , quote ));
    }
  }
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    if IsTrue(OpNotEq2(market , MkUndefined() )) {
      symbol = *(market ).At(MkString("symbol") );
    }
  }
  takerOrMaker := OpCopy(MkUndefined() ); _ = takerOrMaker;
  amountString := OpCopy(MkUndefined() ); _ = amountString;
  cost := OpCopy(MkUndefined() ); _ = cost;
  side := OpCopy(MkUndefined() ); _ = side;
  fee := OpCopy(MkUndefined() ); _ = fee;
  orderId := OpCopy(MkUndefined() ); _ = orderId;
  if IsTrue(OpHasMember(MkString("liquidity") , trade )) {
    if IsTrue(OpEq2(this.SafeString(trade , MkString("liquidity") ), MkString("T") )) {
      takerOrMaker = MkString("taker") ;
    } else {
      if IsTrue(OpEq2(this.SafeString(trade , MkString("liquidity") ), MkString("M") )) {
        takerOrMaker = MkString("maker") ;
      }
    }
    funds := this.SafeValue(trade , MkString("funds") , MkMap(&VarMap{})); _ = funds;
    amountString = this.SafeString(funds , baseId );
    cost = this.SafeNumber(funds , quoteId );
    fee = MkMap(&VarMap{
        "currency":this.SafeString(trade , MkString("fee_currency") ),
        "cost":this.SafeNumber(trade , MkString("fee") ),
        });
    side = this.SafeString(trade , MkString("side") );
    orderId = this.SafeString(trade , MkString("order_id") );
  } else {
    amountString = this.SafeString(trade , MkString("amount") );
    side = this.SafeString(trade , MkString("order_type") );
  }
  price := this.ParseNumber(priceString ); _ = price;
  amount := this.ParseNumber(amountString ); _ = amount;
  if IsTrue(OpEq2(cost , MkUndefined() )) {
    cost = this.ParseNumber(Precise.StringMul(priceString , amountString ));
  }
  return MkMap(&VarMap{
        "id":id ,
        "info":trade ,
        "datetime":this.Iso8601(timestamp ),
        "timestamp":timestamp ,
        "symbol":symbol ,
        "type":MkUndefined() ,
        "side":side ,
        "order":orderId ,
        "takerOrMaker":takerOrMaker ,
        "price":price ,
        "amount":amount ,
        "cost":cost ,
        "fee":fee ,
        });
}

func (this *Coincheck) FetchMyTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  response := this.Call(MkString("privateGetExchangeOrdersTransactions"), this.Extend(MkMap(&VarMap{}), params )); _ = response;
  transactions := this.SafeValue(response , MkString("transactions") , MkArray(&VarArray{})); _ = transactions;
  return this.ParseTrades(transactions , market , since , limit );
}

func (this *Coincheck) FetchTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"pair":*(market ).At(MkString("id") )}); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("publicGetTrades"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = data;
  return this.ParseTrades(data , market , since , limit );
}

func (this *Coincheck) CreateOrder(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 2, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 3, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 4, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 5, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"pair":this.MarketId(symbol )}); _ = request;
  if IsTrue(OpEq2(type_ , MkString("market") )) {
    order_type := OpAdd(type_ , OpAdd(MkString("_") , side )); _ = order_type;
    *(request ).At(MkString("order_type") )= OpCopy(order_type );
    prefix := OpTrinary(OpEq2(side , MkString("buy") ), OpAdd(order_type , MkString("_") ), MkString("") ); _ = prefix;
    *(request ).At(OpAdd(prefix , MkString("amount") ))= OpCopy(amount );
  } else {
    *(request ).At(MkString("order_type") )= OpCopy(side );
    *(request ).At(MkString("rate") )= OpCopy(price );
    *(request ).At(MkString("amount") )= OpCopy(amount );
  }
  response := this.Call(MkString("privatePostExchangeOrders"), this.Extend(request , params )); _ = response;
  id := this.SafeString(response , MkString("id") ); _ = id;
  return MkMap(&VarMap{
        "info":response ,
        "id":id ,
        });
}

func (this *Coincheck) CancelOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"id":id }); _ = request;
  return this.Call(MkString("privateDeleteExchangeOrdersId"), this.Extend(request , params ));
}

func (this *Coincheck) Nonce(goArgs ...*Variant) *Variant{
  return this.Milliseconds();
}

func (this *Coincheck) Sign(goArgs ...*Variant) *Variant{
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
    queryString := MkString("") ; _ = queryString;
    if IsTrue(OpEq2(method , MkString("GET") )) {
      if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
        OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(this.Keysort(query ))));
      }
    } else {
      if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
        body = this.Urlencode(this.Keysort(query ));
        queryString = OpCopy(body );
      }
    }
    auth := OpAdd(nonce , OpAdd(url , queryString )); _ = auth;
    headers = MkMap(&VarMap{
        "Content-Type":MkString("application/x-www-form-urlencoded") ,
        "ACCESS-KEY":*this.At(MkString("apiKey")) ,
        "ACCESS-NONCE":nonce ,
        "ACCESS-SIGNATURE":this.Hmac(this.Encode(auth ), this.Encode(*this.At(MkString("secret")) )),
        });
  }
  return MkMap(&VarMap{
        "url":url ,
        "method":method ,
        "body":body ,
        "headers":headers ,
        });
}

func (this *Coincheck) HandleErrors(goArgs ...*Variant) *Variant{
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
  success := this.SafeValue(response , MkString("success") , MkBool(true) ); _ = success;
  if IsTrue(OpNot(success )) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , this.Json(response )))));
  }
  return MkUndefined()
}

