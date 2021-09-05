package ccxt

import (
)

type Btcalpha struct {
	*ExchangeBase
}

var _ Exchange = (*Btcalpha)(nil)

func init() {
	exchange := &Btcalpha{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Btcalpha) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("btcalpha") ,
            "name":MkString("BTC-Alpha") ,
            "countries":MkArray(&VarArray{
                MkString("US") ,
                }),
            "version":MkString("v1") ,
            "has":MkMap(&VarMap{
                "cancelOrder":MkBool(true) ,
                "createOrder":MkBool(true) ,
                "fetchBalance":MkBool(true) ,
                "fetchClosedOrders":MkBool(true) ,
                "fetchMarkets":MkBool(true) ,
                "fetchMyTrades":MkBool(true) ,
                "fetchOHLCV":MkBool(true) ,
                "fetchOpenOrders":MkBool(true) ,
                "fetchOrder":MkBool(true) ,
                "fetchOrderBook":MkBool(true) ,
                "fetchOrders":MkBool(true) ,
                "fetchTicker":MkBool(false) ,
                "fetchTrades":MkBool(true) ,
                }),
            "timeframes":MkMap(&VarMap{
                "1m":MkString("1") ,
                "5m":MkString("5") ,
                "15m":MkString("15") ,
                "30m":MkString("30") ,
                "1h":MkString("60") ,
                "4h":MkString("240") ,
                "1d":MkString("D") ,
                }),
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/1294454/42625213-dabaa5da-85cf-11e8-8f99-aa8f8f7699f0.jpg") ,
                "api":MkString("https://btc-alpha.com/api") ,
                "www":MkString("https://btc-alpha.com") ,
                "doc":MkString("https://btc-alpha.github.io/api-docs") ,
                "fees":MkString("https://btc-alpha.com/fees/") ,
                "referral":MkString("https://btc-alpha.com/?r=123788") ,
                }),
            "api":MkMap(&VarMap{
                "public":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("currencies/") ,
                        MkString("pairs/") ,
                        MkString("orderbook/{pair_name}/") ,
                        MkString("exchanges/") ,
                        MkString("charts/{pair}/{type}/chart/") ,
                        })}),
                "private":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("wallets/") ,
                        MkString("orders/own/") ,
                        MkString("order/{id}/") ,
                        MkString("exchanges/own/") ,
                        MkString("deposits/") ,
                        MkString("withdraws/") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("order/") ,
                        MkString("order-cancel/") ,
                        }),
                    }),
                }),
            "fees":MkMap(&VarMap{
                "trading":MkMap(&VarMap{
                    "maker":this.ParseNumber(MkString("0.002") ),
                    "taker":this.ParseNumber(MkString("0.002") ),
                    }),
                "funding":MkMap(&VarMap{"withdraw":MkMap(&VarMap{})}),
                }),
            "commonCurrencies":MkMap(&VarMap{"CBC":MkString("Cashbery") }),
            "exceptions":MkMap(&VarMap{
                "exact":MkMap(&VarMap{}),
                "broad":MkMap(&VarMap{"Out of balance":InsufficientFunds }),
                }),
            }));
}

func (this *Btcalpha) FetchMarkets(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetPairs"), params ); _ = response;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , response.Length )); OpIncr(& i ){
    market := *(response ).At(i ); _ = market;
    id := this.SafeString(market , MkString("name") ); _ = id;
    baseId := this.SafeString(market , MkString("currency1") ); _ = baseId;
    quoteId := this.SafeString(market , MkString("currency2") ); _ = quoteId;
    base := this.SafeCurrencyCode(baseId ); _ = base;
    quote := this.SafeCurrencyCode(quoteId ); _ = quote;
    symbol := OpAdd(base , OpAdd(MkString("/") , quote )); _ = symbol;
    pricePrecision := this.SafeString(market , MkString("price_precision") ); _ = pricePrecision;
    priceLimit := this.ParsePrecision(pricePrecision ); _ = priceLimit;
    precision := MkMap(&VarMap{
          "amount":MkInteger(8) ,
          "price":parseInt(pricePrecision ),
          }); _ = precision;
    amountLimit := this.SafeString(market , MkString("minimum_order_size") ); _ = amountLimit;
    result.Push(MkMap(&VarMap{
            "id":id ,
            "symbol":symbol ,
            "base":base ,
            "quote":quote ,
            "active":MkBool(true) ,
            "precision":precision ,
            "limits":MkMap(&VarMap{
                "amount":MkMap(&VarMap{
                    "min":this.ParseNumber(amountLimit ),
                    "max":this.SafeNumber(market , MkString("maximum_order_size") ),
                    }),
                "price":MkMap(&VarMap{
                    "min":this.ParseNumber(priceLimit ),
                    "max":MkUndefined() ,
                    }),
                "cost":MkMap(&VarMap{
                    "min":this.ParseNumber(Precise.StringMul(priceLimit , amountLimit )),
                    "max":MkUndefined() ,
                    }),
                }),
            "info":market ,
            "baseId":MkUndefined() ,
            "quoteId":MkUndefined() ,
            }));
  }
  return result ;
}

func (this *Btcalpha) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"pair_name":this.MarketId(symbol )}); _ = request;
  if IsTrue(limit ) {
    *(request ).At(MkString("limit_sell") )= OpCopy(limit );
    *(request ).At(MkString("limit_buy") )= OpCopy(limit );
  }
  response := this.Call(MkString("publicGetOrderbookPairName"), this.Extend(request , params )); _ = response;
  return this.ParseOrderBook(response , symbol , MkUndefined() , MkString("buy") , MkString("sell") , MkString("price") , MkString("amount") );
}

func (this *Btcalpha) ParseBidsAsks(goArgs ...*Variant) *Variant{
  bidasks := GoGetArg(goArgs, 0, MkUndefined()); _ = bidasks;
  priceKey := GoGetArg(goArgs, 1, MkInteger(0) ); _ = priceKey;
  amountKey := GoGetArg(goArgs, 2, MkInteger(1) ); _ = amountKey;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , bidasks.Length )); OpIncr(& i ){
    bidask := *(bidasks ).At(i ); _ = bidask;
    if IsTrue(bidask ) {
      result.Push(this.ParseBidAsk(bidask , priceKey , amountKey ));
    }
  }
  return result ;
}

func (this *Btcalpha) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  if IsTrue(OpEq2(market , MkUndefined() )) {
    market = this.SafeValue(*this.At(MkString("markets_by_id")) , *(trade ).At(MkString("pair") ));
  }
  if IsTrue(OpNotEq2(market , MkUndefined() )) {
    symbol = *(market ).At(MkString("symbol") );
  }
  timestamp := this.SafeTimestamp(trade , MkString("timestamp") ); _ = timestamp;
  priceString := this.SafeString(trade , MkString("price") ); _ = priceString;
  amountString := this.SafeString(trade , MkString("amount") ); _ = amountString;
  price := this.ParseNumber(priceString ); _ = price;
  amount := this.ParseNumber(amountString ); _ = amount;
  cost := this.ParseNumber(Precise.StringMul(priceString , amountString )); _ = cost;
  id := this.SafeString2(trade , MkString("id") , MkString("tid") ); _ = id;
  side := this.SafeString2(trade , MkString("my_side") , MkString("side") ); _ = side;
  orderId := this.SafeString(trade , MkString("o_id") ); _ = orderId;
  return MkMap(&VarMap{
        "id":id ,
        "info":trade ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "symbol":symbol ,
        "order":orderId ,
        "type":MkString("limit") ,
        "side":side ,
        "takerOrMaker":MkUndefined() ,
        "price":price ,
        "amount":amount ,
        "cost":cost ,
        "fee":MkUndefined() ,
        });
}

func (this *Btcalpha) FetchTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := OpCopy(MkUndefined() ); _ = market;
  request := MkMap(&VarMap{}); _ = request;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("pair") )= *(market ).At(MkString("id") );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  trades := this.Call(MkString("publicGetExchanges"), this.Extend(request , params )); _ = trades;
  return this.ParseTrades(trades , market , since , limit );
}

func (this *Btcalpha) ParseOHLCV(goArgs ...*Variant) *Variant{
  ohlcv := GoGetArg(goArgs, 0, MkUndefined()); _ = ohlcv;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  return MkArray(&VarArray{
        this.SafeTimestamp(ohlcv , MkString("time") ),
        this.SafeNumber(ohlcv , MkString("open") ),
        this.SafeNumber(ohlcv , MkString("high") ),
        this.SafeNumber(ohlcv , MkString("low") ),
        this.SafeNumber(ohlcv , MkString("close") ),
        this.SafeNumber(ohlcv , MkString("volume") ),
        });
}

func (this *Btcalpha) FetchOHLCV(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  timeframe := GoGetArg(goArgs, 1, MkString("5m") ); _ = timeframe;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "pair":*(market ).At(MkString("id") ),
        "type":*(*this.At(MkString("timeframes")) ).At(timeframe ),
        }); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("since") )= parseInt(OpDiv(since , MkInteger(1000) ));
  }
  response := this.Call(MkString("publicGetChartsPairTypeChart"), this.Extend(request , params )); _ = response;
  return this.ParseOHLCVs(response , market , timeframe , since , limit );
}

func (this *Btcalpha) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("privateGetWallets"), params ); _ = response;
  result := MkMap(&VarMap{"info":response }); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , response.Length )); OpIncr(& i ){
    balance := *(response ).At(i ); _ = balance;
    currencyId := this.SafeString(balance , MkString("currency") ); _ = currencyId;
    code := this.SafeCurrencyCode(currencyId ); _ = code;
    account := this.Account(); _ = account;
    *(account ).At(MkString("used") )= this.SafeString(balance , MkString("reserve") );
    *(account ).At(MkString("total") )= this.SafeString(balance , MkString("balance") );
    *(result ).At(code )= OpCopy(account );
  }
  return this.ParseBalance(result );
}

func (this *Btcalpha) ParseOrderStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "1":MkString("open") ,
        "2":MkString("canceled") ,
        "3":MkString("closed") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Btcalpha) ParseOrder(goArgs ...*Variant) *Variant{
  order := GoGetArg(goArgs, 0, MkUndefined()); _ = order;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  if IsTrue(OpEq2(market , MkUndefined() )) {
    market = this.SafeValue(*this.At(MkString("markets_by_id")) , *(order ).At(MkString("pair") ));
  }
  if IsTrue(OpNotEq2(market , MkUndefined() )) {
    symbol = *(market ).At(MkString("symbol") );
  }
  timestamp := this.SafeTimestamp(order , MkString("date") ); _ = timestamp;
  price := this.SafeNumber(order , MkString("price") ); _ = price;
  amount := this.SafeNumber(order , MkString("amount") ); _ = amount;
  status := this.ParseOrderStatus(this.SafeString(order , MkString("status") )); _ = status;
  id := this.SafeString2(order , MkString("oid") , MkString("id") ); _ = id;
  trades := this.SafeValue(order , MkString("trades") , MkArray(&VarArray{})); _ = trades;
  trades = this.ParseTrades(trades , market );
  side := this.SafeString2(order , MkString("my_side") , MkString("type") ); _ = side;
  filled := OpCopy(MkUndefined() ); _ = filled;
  numTrades := OpCopy(trades.Length ); _ = numTrades;
  if IsTrue(OpGt(numTrades , MkInteger(0) )) {
    filled = MkNumber(0.0) ;
    for i := MkInteger(0) ; IsTrue(OpLw(i , numTrades )); OpIncr(& i ){
      filled = this.Sum(filled , *(*(trades ).At(i )).At(MkString("amount") ));
    }
  }
  remaining := OpCopy(MkUndefined() ); _ = remaining;
  if IsTrue(OpAnd(OpNotEq2(amount , MkUndefined() ), OpAnd(OpGt(amount , MkInteger(0) ), OpNotEq2(filled , MkUndefined() )))) {
    remaining = Math.Max(MkInteger(0) , OpSub(amount , filled ));
  }
  return MkMap(&VarMap{
        "id":id ,
        "clientOrderId":MkUndefined() ,
        "datetime":this.Iso8601(timestamp ),
        "timestamp":timestamp ,
        "status":status ,
        "symbol":symbol ,
        "type":MkString("limit") ,
        "timeInForce":MkUndefined() ,
        "postOnly":MkUndefined() ,
        "side":side ,
        "price":price ,
        "stopPrice":MkUndefined() ,
        "cost":MkUndefined() ,
        "amount":amount ,
        "filled":filled ,
        "remaining":remaining ,
        "trades":trades ,
        "fee":MkUndefined() ,
        "info":order ,
        "lastTradeTimestamp":MkUndefined() ,
        "average":MkUndefined() ,
        });
}

func (this *Btcalpha) CreateOrder(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 2, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 3, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 4, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 5, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "pair":*(market ).At(MkString("id") ),
        "type":side ,
        "amount":amount ,
        "price":this.PriceToPrecision(symbol , price ),
        }); _ = request;
  response := this.Call(MkString("privatePostOrder"), this.Extend(request , params )); _ = response;
  if IsTrue(OpNot(*(response ).At(MkString("success") ))) {
    panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , this.Json(response )))));
  }
  order := this.ParseOrder(response , market ); _ = order;
  amount = OpTrinary(OpGt(*(order ).At(MkString("amount") ), MkInteger(0) ), *(order ).At(MkString("amount") ), amount );
  return this.Extend(order , MkMap(&VarMap{"amount":amount }));
}

func (this *Btcalpha) CancelOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"order":id }); _ = request;
  response := this.Call(MkString("privatePostOrderCancel"), this.Extend(request , params )); _ = response;
  return response ;
}

func (this *Btcalpha) FetchOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"id":id }); _ = request;
  order := this.Call(MkString("privateGetOrderId"), this.Extend(request , params )); _ = order;
  return this.ParseOrder(order );
}

func (this *Btcalpha) FetchOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("pair") )= *(market ).At(MkString("id") );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  orders := this.Call(MkString("privateGetOrdersOwn"), this.Extend(request , params )); _ = orders;
  return this.ParseOrders(orders , market , since , limit );
}

func (this *Btcalpha) FetchOpenOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"status":MkString("1") }); _ = request;
  return this.FetchOrders(symbol , since , limit , this.Extend(request , params ));
}

func (this *Btcalpha) FetchClosedOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"status":MkString("3") }); _ = request;
  return this.FetchOrders(symbol , since , limit , this.Extend(request , params ));
}

func (this *Btcalpha) FetchMyTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market := this.Market(symbol ); _ = market;
    *(request ).At(MkString("pair") )= *(market ).At(MkString("id") );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  trades := this.Call(MkString("privateGetExchangesOwn"), this.Extend(request , params )); _ = trades;
  return this.ParseTrades(trades , MkUndefined() , since , limit );
}

func (this *Btcalpha) Nonce(goArgs ...*Variant) *Variant{
  return this.Milliseconds();
}

func (this *Btcalpha) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("public") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  query := this.Urlencode(this.Keysort(this.Omit(params , this.ExtractParams(path )))); _ = query;
  url := OpAdd(*(*this.At(MkString("urls")) ).At(MkString("api") ), MkString("/") ); _ = url;
  if IsTrue(OpNotEq2(path , MkString("charts/{pair}/{type}/chart/") )) {
    OpAddAssign(& url , MkString("v1/") );
  }
  OpAddAssign(& url , this.ImplodeParams(path , params ));
  headers = MkMap(&VarMap{"Accept":MkString("application/json") });
  if IsTrue(OpEq2(api , MkString("public") )) {
    if IsTrue(query.Length ) {
      OpAddAssign(& url , OpAdd(MkString("?") , query ));
    }
  } else {
    this.CheckRequiredCredentials();
    payload := OpCopy(*this.At(MkString("apiKey")) ); _ = payload;
    if IsTrue(OpEq2(method , MkString("POST") )) {
      *(headers ).At(MkString("Content-Type") )= MkString("application/x-www-form-urlencoded") ;
      body = OpCopy(query );
      OpAddAssign(& payload , body );
    } else {
      if IsTrue(query.Length ) {
        OpAddAssign(& url , OpAdd(MkString("?") , query ));
      }
    }
    *(headers ).At(MkString("X-KEY") )= OpCopy(*this.At(MkString("apiKey")) );
    *(headers ).At(MkString("X-SIGN") )= this.Hmac(this.Encode(payload ), this.Encode(*this.At(MkString("secret")) ));
    *(headers ).At(MkString("X-NONCE") )= (this.Nonce()).Call(MkString("toString") );
  }
  return MkMap(&VarMap{
        "url":url ,
        "method":method ,
        "body":body ,
        "headers":headers ,
        });
}

func (this *Btcalpha) HandleErrors(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
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
  error := this.SafeString(response , MkString("error") ); _ = error;
  feedback := OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , body )); _ = feedback;
  if IsTrue(OpNotEq2(error , MkUndefined() )) {
    this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("exact") ), error , feedback );
    this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("broad") ), error , feedback );
  }
  if IsTrue(OpOr(OpEq2(code , MkInteger(401) ), OpEq2(code , MkInteger(403) ))) {
    panic(NewAuthenticationError(feedback ));
  } else {
    if IsTrue(OpEq2(code , MkInteger(429) )) {
      panic(NewDDoSProtection(feedback ));
    }
  }
  if IsTrue(OpLw(code , MkInteger(400) )) {
    return MkUndefined();
  }
  panic(NewExchangeError(feedback ));
  return MkUndefined()
}

