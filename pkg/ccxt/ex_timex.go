package ccxt

import (
)

type Timex struct {
	*ExchangeBase
}

var _ Exchange = (*Timex)(nil)

func init() {
	exchange := &Timex{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Timex) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("timex") ,
            "name":MkString("TimeX") ,
            "countries":MkArray(&VarArray{
                MkString("AU") ,
                }),
            "version":MkString("v1") ,
            "rateLimit":MkInteger(1500) ,
            "has":MkMap(&VarMap{
                "cancelOrder":MkBool(true) ,
                "cancelOrders":MkBool(true) ,
                "CORS":MkBool(false) ,
                "createOrder":MkBool(true) ,
                "editOrder":MkBool(true) ,
                "fetchBalance":MkBool(true) ,
                "fetchClosedOrders":MkBool(true) ,
                "fetchCurrencies":MkBool(true) ,
                "fetchMarkets":MkBool(true) ,
                "fetchMyTrades":MkBool(true) ,
                "fetchOHLCV":MkBool(true) ,
                "fetchOpenOrders":MkBool(true) ,
                "fetchOrder":MkBool(true) ,
                "fetchOrderBook":MkBool(true) ,
                "fetchTicker":MkBool(true) ,
                "fetchTickers":MkBool(true) ,
                "fetchTrades":MkBool(true) ,
                "fetchTradingFee":MkBool(true) ,
                }),
            "timeframes":MkMap(&VarMap{
                "1m":MkString("I1") ,
                "5m":MkString("I5") ,
                "15m":MkString("I15") ,
                "30m":MkString("I30") ,
                "1h":MkString("H1") ,
                "2h":MkString("H2") ,
                "4h":MkString("H4") ,
                "6h":MkString("H6") ,
                "12h":MkString("H12") ,
                "1d":MkString("D1") ,
                "1w":MkString("W1") ,
                }),
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/1294454/70423869-6839ab00-1a7f-11ea-8f94-13ae72c31115.jpg") ,
                "api":MkString("https://plasma-relay-backend.timex.io") ,
                "www":MkString("https://timex.io") ,
                "doc":MkString("https://docs.timex.io") ,
                "referral":MkString("https://timex.io/?refcode=1x27vNkTbP1uwkCck") ,
                }),
            "api":MkMap(&VarMap{
                "custody":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("credentials") ,
                        MkString("credentials/h/{hash}") ,
                        MkString("credentials/k/{key}") ,
                        MkString("credentials/me/address") ,
                        MkString("deposit-addresses") ,
                        MkString("deposit-addresses/h/{hash}") ,
                        })}),
                "history":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("orders") ,
                        MkString("orders/details") ,
                        MkString("orders/export/csv") ,
                        MkString("trades") ,
                        MkString("trades/export/csv") ,
                        })}),
                "currencies":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("a/{address}") ,
                        MkString("i/{id}") ,
                        MkString("s/{symbol}") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("perform") ,
                        MkString("prepare") ,
                        MkString("remove/perform") ,
                        MkString("s/{symbol}/remove/prepare") ,
                        MkString("s/{symbol}/update/perform") ,
                        MkString("s/{symbol}/update/prepare") ,
                        }),
                    }),
                "markets":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("i/{id}") ,
                        MkString("s/{symbol}") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("perform") ,
                        MkString("prepare") ,
                        MkString("remove/perform") ,
                        MkString("s/{symbol}/remove/prepare") ,
                        MkString("s/{symbol}/update/perform") ,
                        MkString("s/{symbol}/update/prepare") ,
                        }),
                    }),
                "public":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("candles") ,
                        MkString("currencies") ,
                        MkString("markets") ,
                        MkString("orderbook") ,
                        MkString("orderbook/raw") ,
                        MkString("orderbook/v2") ,
                        MkString("tickers") ,
                        MkString("trades") ,
                        })}),
                "statistics":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("address") ,
                        })}),
                "trading":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("balances") ,
                        MkString("fees") ,
                        MkString("orders") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("orders") ,
                        MkString("orders/json") ,
                        }),
                    "put":MkArray(&VarArray{
                        MkString("orders") ,
                        MkString("orders/json") ,
                        }),
                    "delete":MkArray(&VarArray{
                        MkString("orders") ,
                        MkString("orders/json") ,
                        }),
                    }),
                "tradingview":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("config") ,
                        MkString("history") ,
                        MkString("symbol_info") ,
                        MkString("time") ,
                        })}),
                }),
            "exceptions":MkMap(&VarMap{
                "exact":MkMap(&VarMap{
                    "0":ExchangeError ,
                    "1":NotSupported ,
                    "4000":BadRequest ,
                    "4001":BadRequest ,
                    "4002":InsufficientFunds ,
                    "4003":AuthenticationError ,
                    "4004":AuthenticationError ,
                    "4005":BadRequest ,
                    "4006":BadRequest ,
                    "4007":BadRequest ,
                    "4300":PermissionDenied ,
                    "4100":AuthenticationError ,
                    "4400":OrderNotFound ,
                    "5001":InvalidOrder ,
                    "5002":ExchangeError ,
                    "400":BadRequest ,
                    "401":AuthenticationError ,
                    "403":PermissionDenied ,
                    "404":OrderNotFound ,
                    "429":RateLimitExceeded ,
                    "500":ExchangeError ,
                    "503":ExchangeNotAvailable ,
                    }),
                "broad":MkMap(&VarMap{"Insufficient":InsufficientFunds }),
                }),
            "options":MkMap(&VarMap{
                "fetchTickers":MkMap(&VarMap{"period":MkString("1d") }),
                "fetchTrades":MkMap(&VarMap{"sort":MkString("timestamp,asc") }),
                "fetchMyTrades":MkMap(&VarMap{"sort":MkString("timestamp,asc") }),
                "fetchOpenOrders":MkMap(&VarMap{"sort":MkString("createdAt,asc") }),
                "fetchClosedOrders":MkMap(&VarMap{"sort":MkString("createdAt,asc") }),
                "defaultSort":MkString("timestamp,asc") ,
                "defaultSortOrders":MkString("createdAt,asc") ,
                }),
            }));
}

func (this *Timex) FetchMarkets(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetMarkets"), params ); _ = response;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , response.Length )); OpIncr(& i ){
    result.Push(this.ParseMarket(*(response ).At(i )));
  }
  return result ;
}

func (this *Timex) FetchCurrencies(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetCurrencies"), params ); _ = response;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , response.Length )); OpIncr(& i ){
    currency := *(response ).At(i ); _ = currency;
    result.Push(this.ParseCurrency(currency ));
  }
  return this.IndexBy(result , MkString("code") );
}

func (this *Timex) FetchTickers(goArgs ...*Variant) *Variant{
  symbols := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbols;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  period := this.SafeString(*(*this.At(MkString("options")) ).At(MkString("fetchTickers") ), MkString("period") , MkString("1d") ); _ = period;
  request := MkMap(&VarMap{"period":*(*this.At(MkString("timeframes")) ).At(period )}); _ = request;
  response := this.Call(MkString("publicGetTickers"), this.Extend(request , params )); _ = response;
  return this.ParseTickers(response , symbols );
}

func (this *Timex) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  period := this.SafeString(*(*this.At(MkString("options")) ).At(MkString("fetchTickers") ), MkString("period") , MkString("1d") ); _ = period;
  request := MkMap(&VarMap{
        "market":*(market ).At(MkString("id") ),
        "period":*(*this.At(MkString("timeframes")) ).At(period ),
        }); _ = request;
  response := this.Call(MkString("publicGetTickers"), this.Extend(request , params )); _ = response;
  ticker := this.SafeValue(response , MkInteger(0) ); _ = ticker;
  return this.ParseTicker(ticker , market );
}

func (this *Timex) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"market":*(market ).At(MkString("id") )}); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("publicGetOrderbookV2"), this.Extend(request , params )); _ = response;
  timestamp := this.Parse8601(this.SafeString(response , MkString("timestamp") )); _ = timestamp;
  return this.ParseOrderBook(response , symbol , timestamp , MkString("bid") , MkString("ask") , MkString("price") , MkString("baseTokenAmount") );
}

func (this *Timex) FetchTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  options := this.SafeValue(*this.At(MkString("options")) , MkString("fetchTrades") , MkMap(&VarMap{})); _ = options;
  defaultSort := this.SafeValue(options , MkString("sort") , MkString("timestamp,asc") ); _ = defaultSort;
  sort := this.SafeString(params , MkString("sort") , defaultSort ); _ = sort;
  query := this.Omit(params , MkString("sort") ); _ = query;
  request := MkMap(&VarMap{
        "market":*(market ).At(MkString("id") ),
        "sort":sort ,
        }); _ = request;
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("from") )= this.Iso8601(since );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("size") )= OpCopy(limit );
  }
  response := this.Call(MkString("publicGetTrades"), this.Extend(request , query )); _ = response;
  return this.ParseTrades(response , market , since , limit );
}

func (this *Timex) FetchOHLCV(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  timeframe := GoGetArg(goArgs, 1, MkString("1m") ); _ = timeframe;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "market":*(market ).At(MkString("id") ),
        "period":*(*this.At(MkString("timeframes")) ).At(timeframe ),
        }); _ = request;
  duration := this.ParseTimeframe(timeframe ); _ = duration;
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("from") )= this.Iso8601(since );
    if IsTrue(OpNotEq2(limit , MkUndefined() )) {
      *(request ).At(MkString("till") )= this.Iso8601(this.Sum(since , OpMulti(this.Sum(limit , MkInteger(1) ), OpMulti(duration , MkInteger(1000) ))));
    }
  } else {
    if IsTrue(OpNotEq2(limit , MkUndefined() )) {
      now := this.Milliseconds(); _ = now;
      *(request ).At(MkString("till") )= this.Iso8601(now );
      *(request ).At(MkString("from") )= this.Iso8601(OpSub(now , OpSub(OpMulti(limit , OpMulti(duration , MkInteger(1000) )), MkInteger(1) )));
    } else {
      *(request ).At(MkString("till") )= this.Iso8601(this.Milliseconds());
    }
  }
  response := this.Call(MkString("publicGetCandles"), this.Extend(request , params )); _ = response;
  return this.ParseOHLCVs(response , market , timeframe , since , limit );
}

func (this *Timex) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("tradingGetBalances"), params ); _ = response;
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
    *(account ).At(MkString("total") )= this.SafeString(balance , MkString("totalBalance") );
    *(account ).At(MkString("used") )= this.SafeString(balance , MkString("lockedBalance") );
    *(result ).At(code )= OpCopy(account );
  }
  return this.ParseBalance(result );
}

func (this *Timex) CreateOrder(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 2, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 3, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 4, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 5, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  uppercaseSide := side.ToUpperCase(); _ = uppercaseSide;
  uppercaseType := type_.ToUpperCase(); _ = uppercaseType;
  request := MkMap(&VarMap{
        "symbol":*(market ).At(MkString("id") ),
        "quantity":this.AmountToPrecision(symbol , amount ),
        "side":uppercaseSide ,
        "orderTypes":uppercaseType ,
        }); _ = request;
  query := OpCopy(params ); _ = query;
  if IsTrue(OpEq2(uppercaseType , MkString("LIMIT") )) {
    *(request ).At(MkString("price") )= this.PriceToPrecision(symbol , price );
    defaultExpireIn := this.SafeInteger(*this.At(MkString("options")) , MkString("expireIn") ); _ = defaultExpireIn;
    expireTime := this.SafeValue(params , MkString("expireTime") ); _ = expireTime;
    expireIn := this.SafeValue(params , MkString("expireIn") , defaultExpireIn ); _ = expireIn;
    if IsTrue(OpNotEq2(expireTime , MkUndefined() )) {
      *(request ).At(MkString("expireTime") )= OpCopy(expireTime );
    } else {
      if IsTrue(OpNotEq2(expireIn , MkUndefined() )) {
        *(request ).At(MkString("expireIn") )= OpCopy(expireIn );
      } else {
        panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" createOrder() method requires a expireTime or expireIn param for a ") , OpAdd(type_ , MkString(" order, you can also set the expireIn exchange-wide option") )))));
      }
    }
    query = this.Omit(params , MkArray(&VarArray{
            MkString("expireTime") ,
            MkString("expireIn") ,
            }));
  } else {
    *(request ).At(MkString("price") )= MkInteger(0) ;
  }
  response := this.Call(MkString("tradingPostOrders"), this.Extend(request , query )); _ = response;
  orders := this.SafeValue(response , MkString("orders") , MkArray(&VarArray{})); _ = orders;
  order := this.SafeValue(orders , MkInteger(0) , MkMap(&VarMap{})); _ = order;
  return this.ParseOrder(order , market );
}

func (this *Timex) EditOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 2, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 3, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 4, MkUndefined() ); _ = amount;
  price := GoGetArg(goArgs, 5, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 6, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"id":id }); _ = request;
  if IsTrue(OpNotEq2(amount , MkUndefined() )) {
    *(request ).At(MkString("quantity") )= this.AmountToPrecision(symbol , amount );
  }
  if IsTrue(OpNotEq2(price , MkUndefined() )) {
    *(request ).At(MkString("price") )= this.PriceToPrecision(symbol , price );
  }
  response := this.Call(MkString("tradingPutOrders"), this.Extend(request , params )); _ = response;
  if IsTrue(OpHasMember(MkString("unchangedOrders") , response )) {
    orderIds := this.SafeValue(response , MkString("unchangedOrders") , MkArray(&VarArray{})); _ = orderIds;
    orderId := this.SafeString(orderIds , MkInteger(0) ); _ = orderId;
    return MkMap(&VarMap{
          "id":orderId ,
          "info":response ,
          });
  }
  orders := this.SafeValue(response , MkString("changedOrders") , MkArray(&VarArray{})); _ = orders;
  firstOrder := this.SafeValue(orders , MkInteger(0) , MkMap(&VarMap{})); _ = firstOrder;
  order := this.SafeValue(firstOrder , MkString("newOrder") , MkMap(&VarMap{})); _ = order;
  return this.ParseOrder(order , market );
}

func (this *Timex) CancelOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  return this.CancelOrders(MkArray(&VarArray{
            id ,
            }), symbol , params );
}

func (this *Timex) CancelOrders(goArgs ...*Variant) *Variant{
  ids := GoGetArg(goArgs, 0, MkUndefined()); _ = ids;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"id":ids }); _ = request;
  response := this.Call(MkString("tradingDeleteOrders"), this.Extend(request , params )); _ = response;
  return response ;
}

func (this *Timex) FetchOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"orderHash":id }); _ = request;
  response := this.Call(MkString("historyGetOrdersDetails"), request ); _ = response;
  order := this.SafeValue(response , MkString("order") , MkMap(&VarMap{})); _ = order;
  trades := this.SafeValue(response , MkString("trades") , MkArray(&VarArray{})); _ = trades;
  return this.ParseOrder(this.Extend(order , MkMap(&VarMap{"trades":trades })));
}

func (this *Timex) FetchOpenOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  options := this.SafeValue(*this.At(MkString("options")) , MkString("fetchOpenOrders") , MkMap(&VarMap{})); _ = options;
  defaultSort := this.SafeValue(options , MkString("sort") , MkString("createdAt,asc") ); _ = defaultSort;
  sort := this.SafeString(params , MkString("sort") , defaultSort ); _ = sort;
  query := this.Omit(params , MkString("sort") ); _ = query;
  request := MkMap(&VarMap{"sort":sort }); _ = request;
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("symbol") )= *(market ).At(MkString("id") );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("size") )= OpCopy(limit );
  }
  response := this.Call(MkString("tradingGetOrders"), this.Extend(request , query )); _ = response;
  orders := this.SafeValue(response , MkString("orders") , MkArray(&VarArray{})); _ = orders;
  return this.ParseOrders(orders , market , since , limit );
}

func (this *Timex) FetchClosedOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  options := this.SafeValue(*this.At(MkString("options")) , MkString("fetchClosedOrders") , MkMap(&VarMap{})); _ = options;
  defaultSort := this.SafeValue(options , MkString("sort") , MkString("createdAt,asc") ); _ = defaultSort;
  sort := this.SafeString(params , MkString("sort") , defaultSort ); _ = sort;
  query := this.Omit(params , MkString("sort") ); _ = query;
  request := MkMap(&VarMap{
        "sort":sort ,
        "side":MkString("BUY") ,
        }); _ = request;
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("symbol") )= *(market ).At(MkString("id") );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("from") )= this.Iso8601(since );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("size") )= OpCopy(limit );
  }
  response := this.Call(MkString("historyGetOrders"), this.Extend(request , query )); _ = response;
  orders := this.SafeValue(response , MkString("orders") , MkArray(&VarArray{})); _ = orders;
  return this.ParseOrders(orders , market , since , limit );
}

func (this *Timex) FetchMyTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  options := this.SafeValue(*this.At(MkString("options")) , MkString("fetchMyTrades") , MkMap(&VarMap{})); _ = options;
  defaultSort := this.SafeValue(options , MkString("sort") , MkString("timestamp,asc") ); _ = defaultSort;
  sort := this.SafeString(params , MkString("sort") , defaultSort ); _ = sort;
  query := this.Omit(params , MkString("sort") ); _ = query;
  request := MkMap(&VarMap{"sort":sort }); _ = request;
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("symbol") )= *(market ).At(MkString("id") );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("from") )= this.Iso8601(since );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("size") )= OpCopy(limit );
  }
  response := this.Call(MkString("historyGetTrades"), this.Extend(request , query )); _ = response;
  trades := this.SafeValue(response , MkString("trades") , MkArray(&VarArray{})); _ = trades;
  return this.ParseTrades(trades , market , since , limit );
}

func (this *Timex) FetchTradingFee(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"markets":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("tradingGetFees"), this.Extend(request , params )); _ = response;
  result := this.SafeValue(response , MkInteger(0) , MkMap(&VarMap{})); _ = result;
  return MkMap(&VarMap{
        "info":response ,
        "maker":this.SafeNumber(result , MkString("fee") ),
        "taker":MkUndefined() ,
        });
}

func (this *Timex) ParseMarket(goArgs ...*Variant) *Variant{
  market := GoGetArg(goArgs, 0, MkUndefined()); _ = market;
  locked := this.SafeValue(market , MkString("locked") ); _ = locked;
  active := OpNot(locked ); _ = active;
  id := this.SafeString(market , MkString("symbol") ); _ = id;
  baseId := this.SafeString(market , MkString("baseCurrency") ); _ = baseId;
  quoteId := this.SafeString(market , MkString("quoteCurrency") ); _ = quoteId;
  base := this.SafeCurrencyCode(baseId ); _ = base;
  quote := this.SafeCurrencyCode(quoteId ); _ = quote;
  symbol := OpAdd(base , OpAdd(MkString("/") , quote )); _ = symbol;
  precision := MkMap(&VarMap{
        "amount":this.PrecisionFromString(this.SafeString(market , MkString("quantityIncrement") )),
        "price":this.PrecisionFromString(this.SafeString(market , MkString("tickSize") )),
        }); _ = precision;
  amountIncrement := this.SafeNumber(market , MkString("quantityIncrement") ); _ = amountIncrement;
  minBase := this.SafeNumber(market , MkString("baseMinSize") ); _ = minBase;
  minAmount := Math.Max(amountIncrement , minBase ); _ = minAmount;
  priceIncrement := this.SafeNumber(market , MkString("tickSize") ); _ = priceIncrement;
  minCost := this.SafeNumber(market , MkString("quoteMinSize") ); _ = minCost;
  limits := MkMap(&VarMap{
        "amount":MkMap(&VarMap{
            "min":minAmount ,
            "max":MkUndefined() ,
            }),
        "price":MkMap(&VarMap{
            "min":priceIncrement ,
            "max":MkUndefined() ,
            }),
        "cost":MkMap(&VarMap{
            "min":Math.Max(minCost , OpMulti(minAmount , priceIncrement )),
            "max":MkUndefined() ,
            }),
        }); _ = limits;
  taker := this.SafeNumber(market , MkString("takerFee") ); _ = taker;
  maker := this.SafeNumber(market , MkString("makerFee") ); _ = maker;
  return MkMap(&VarMap{
        "id":id ,
        "symbol":symbol ,
        "base":base ,
        "quote":quote ,
        "baseId":baseId ,
        "quoteId":quoteId ,
        "type":MkString("spot") ,
        "active":active ,
        "precision":precision ,
        "limits":limits ,
        "taker":taker ,
        "maker":maker ,
        "info":market ,
        });
}

func (this *Timex) ParseCurrency(goArgs ...*Variant) *Variant{
  currency := GoGetArg(goArgs, 0, MkUndefined()); _ = currency;
  id := this.SafeString(currency , MkString("symbol") ); _ = id;
  code := this.SafeCurrencyCode(id ); _ = code;
  name := this.SafeString(currency , MkString("name") ); _ = name;
  precision := this.SafeInteger(currency , MkString("decimals") ); _ = precision;
  active := this.SafeValue(currency , MkString("active") ); _ = active;
  feeString := this.SafeString(currency , MkString("withdrawalFee") ); _ = feeString;
  tradeDecimals := this.SafeInteger(currency , MkString("tradeDecimals") ); _ = tradeDecimals;
  fee := OpCopy(MkUndefined() ); _ = fee;
  if IsTrue(OpAnd(OpNotEq2(feeString , MkUndefined() ), OpNotEq2(tradeDecimals , MkUndefined() ))) {
    feeStringLen := OpCopy(feeString.Length ); _ = feeStringLen;
    dotIndex := OpSub(feeStringLen , tradeDecimals ); _ = dotIndex;
    if IsTrue(OpGt(dotIndex , MkInteger(0) )) {
      whole := feeString.Slice(MkInteger(0) , dotIndex ); _ = whole;
      fraction := feeString.Slice(OpNeg(dotIndex )); _ = fraction;
      fee = this.ParseNumber(OpAdd(whole , OpAdd(MkString(".") , fraction )));
    } else {
      fraction := MkString(".") ; _ = fraction;
      for i := MkInteger(0) ; IsTrue(OpLw(i , OpNeg(dotIndex ))); OpIncr(& i ){
        OpAddAssign(& fraction , MkString("0") );
      }
      fee = this.ParseNumber(OpAdd(fraction , feeString ));
    }
  }
  return MkMap(&VarMap{
        "id":code ,
        "code":code ,
        "info":currency ,
        "type":MkUndefined() ,
        "name":name ,
        "active":active ,
        "fee":fee ,
        "precision":precision ,
        "limits":MkMap(&VarMap{
            "withdraw":MkMap(&VarMap{
                "min":fee ,
                "max":MkUndefined() ,
                }),
            "amount":MkMap(&VarMap{
                "min":MkUndefined() ,
                "max":MkUndefined() ,
                }),
            }),
        });
}

func (this *Timex) ParseTicker(goArgs ...*Variant) *Variant{
  ticker := GoGetArg(goArgs, 0, MkUndefined()); _ = ticker;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  marketId := this.SafeString(ticker , MkString("market") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market , MkString("/") ); _ = symbol;
  timestamp := this.Parse8601(this.SafeString(ticker , MkString("timestamp") )); _ = timestamp;
  last := this.SafeNumber(ticker , MkString("last") ); _ = last;
  open := this.SafeNumber(ticker , MkString("open") ); _ = open;
  change := OpCopy(MkUndefined() ); _ = change;
  average := OpCopy(MkUndefined() ); _ = average;
  if IsTrue(OpAnd(OpNotEq2(last , MkUndefined() ), OpNotEq2(open , MkUndefined() ))) {
    change = OpSub(last , open );
    average = OpDiv(this.Sum(last , open ), MkInteger(2) );
  }
  percentage := OpCopy(MkUndefined() ); _ = percentage;
  if IsTrue(OpAnd(OpNotEq2(change , MkUndefined() ), open )) {
    percentage = OpMulti(OpDiv(change , open ), MkInteger(100) );
  }
  return MkMap(&VarMap{
        "symbol":symbol ,
        "info":ticker ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "high":this.SafeNumber(ticker , MkString("high") ),
        "low":this.SafeNumber(ticker , MkString("low") ),
        "bid":this.SafeNumber(ticker , MkString("bid") ),
        "bidVolume":MkUndefined() ,
        "ask":this.SafeNumber(ticker , MkString("ask") ),
        "askVolume":MkUndefined() ,
        "vwap":MkUndefined() ,
        "open":open ,
        "close":last ,
        "last":last ,
        "previousClose":MkUndefined() ,
        "change":change ,
        "percentage":percentage ,
        "average":average ,
        "baseVolume":this.SafeNumber(ticker , MkString("volume") ),
        "quoteVolume":this.SafeNumber(ticker , MkString("volumeQuote") ),
        });
}

func (this *Timex) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  marketId := this.SafeString(trade , MkString("symbol") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market ); _ = symbol;
  timestamp := this.Parse8601(this.SafeString(trade , MkString("timestamp") )); _ = timestamp;
  priceString := this.SafeString(trade , MkString("price") ); _ = priceString;
  amountString := this.SafeString(trade , MkString("quantity") ); _ = amountString;
  price := this.ParseNumber(priceString ); _ = price;
  amount := this.ParseNumber(amountString ); _ = amount;
  cost := this.ParseNumber(Precise.StringMul(priceString , amountString )); _ = cost;
  id := this.SafeString(trade , MkString("id") ); _ = id;
  side := this.SafeStringLower2(trade , MkString("direction") , MkString("side") ); _ = side;
  takerOrMaker := this.SafeStringLower(trade , MkString("makerOrTaker") ); _ = takerOrMaker;
  orderId := OpCopy(MkUndefined() ); _ = orderId;
  if IsTrue(OpNotEq2(takerOrMaker , MkUndefined() )) {
    orderId = this.SafeString(trade , OpAdd(takerOrMaker , MkString("OrderId") ));
  }
  fee := OpCopy(MkUndefined() ); _ = fee;
  feeCost := this.SafeNumber(trade , MkString("fee") ); _ = feeCost;
  feeCurrency := this.SafeCurrencyCode(this.SafeString(trade , MkString("feeToken") )); _ = feeCurrency;
  if IsTrue(OpNotEq2(feeCost , MkUndefined() )) {
    fee = MkMap(&VarMap{
        "cost":feeCost ,
        "currency":feeCurrency ,
        });
  }
  return MkMap(&VarMap{
        "info":trade ,
        "id":id ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "symbol":symbol ,
        "order":orderId ,
        "type":MkUndefined() ,
        "side":side ,
        "price":price ,
        "amount":amount ,
        "cost":cost ,
        "takerOrMaker":takerOrMaker ,
        "fee":fee ,
        });
}

func (this *Timex) ParseOHLCV(goArgs ...*Variant) *Variant{
  ohlcv := GoGetArg(goArgs, 0, MkUndefined()); _ = ohlcv;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  return MkArray(&VarArray{
        this.Parse8601(this.SafeString(ohlcv , MkString("timestamp") )),
        this.SafeNumber(ohlcv , MkString("open") ),
        this.SafeNumber(ohlcv , MkString("high") ),
        this.SafeNumber(ohlcv , MkString("low") ),
        this.SafeNumber(ohlcv , MkString("close") ),
        this.SafeNumber(ohlcv , MkString("volume") ),
        });
}

func (this *Timex) ParseOrder(goArgs ...*Variant) *Variant{
  order := GoGetArg(goArgs, 0, MkUndefined()); _ = order;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  id := this.SafeString(order , MkString("id") ); _ = id;
  type_ := this.SafeStringLower(order , MkString("type") ); _ = type_;
  side := this.SafeStringLower(order , MkString("side") ); _ = side;
  marketId := this.SafeString(order , MkString("symbol") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market ); _ = symbol;
  timestamp := this.Parse8601(this.SafeString(order , MkString("createdAt") )); _ = timestamp;
  price := this.SafeNumber(order , MkString("price") ); _ = price;
  amount := this.SafeNumber(order , MkString("quantity") ); _ = amount;
  filled := this.SafeNumber(order , MkString("filledQuantity") ); _ = filled;
  canceledQuantity := this.SafeNumber(order , MkString("cancelledQuantity") ); _ = canceledQuantity;
  status := OpCopy(MkUndefined() ); _ = status;
  if IsTrue(OpAnd(OpNotEq2(amount , MkUndefined() ), OpNotEq2(filled , MkUndefined() ))) {
    if IsTrue(OpGtEq(filled , amount )) {
      status = MkString("closed") ;
    } else {
      if IsTrue(OpAnd(OpNotEq2(canceledQuantity , MkUndefined() ), OpGt(canceledQuantity , MkInteger(0) ))) {
        status = MkString("canceled") ;
      } else {
        status = MkString("open") ;
      }
    }
  }
  rawTrades := this.SafeValue(order , MkString("trades") , MkArray(&VarArray{})); _ = rawTrades;
  trades := this.ParseTrades(rawTrades , market , MkUndefined() , MkUndefined() , MkMap(&VarMap{
            "order":id ,
            "type":type_ ,
            })); _ = trades;
  clientOrderId := this.SafeString(order , MkString("clientOrderId") ); _ = clientOrderId;
  return this.SafeOrder(MkMap(&VarMap{
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
            "stopPrice":MkUndefined() ,
            "amount":amount ,
            "cost":MkUndefined() ,
            "average":MkUndefined() ,
            "filled":filled ,
            "remaining":MkUndefined() ,
            "status":status ,
            "fee":MkUndefined() ,
            "trades":trades ,
            }));
}

func (this *Timex) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("public") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  url := OpAdd(*(*this.At(MkString("urls")) ).At(MkString("api") ), OpAdd(MkString("/") , OpAdd(api , OpAdd(MkString("/") , path )))); _ = url;
  if IsTrue(*(GoGetKeys(params )).At(MkString("length") )) {
    OpAddAssign(& url , OpAdd(MkString("?") , this.UrlencodeWithArrayRepeat(params )));
  }
  if IsTrue(OpNotEq2(api , MkString("public") )) {
    this.CheckRequiredCredentials();
    auth := this.StringToBase64(OpAdd(*this.At(MkString("apiKey")) , OpAdd(MkString(":") , *this.At(MkString("secret")) ))); _ = auth;
    secret := OpAdd(MkString("Basic ") , this.Decode(auth )); _ = secret;
    headers = MkMap(&VarMap{"authorization":secret });
  }
  return MkMap(&VarMap{
        "url":url ,
        "method":method ,
        "body":body ,
        "headers":headers ,
        });
}

func (this *Timex) HandleErrors(goArgs ...*Variant) *Variant{
  statusCode := GoGetArg(goArgs, 0, MkUndefined()); _ = statusCode;
  statusText := GoGetArg(goArgs, 1, MkUndefined()); _ = statusText;
  url := GoGetArg(goArgs, 2, MkUndefined()); _ = url;
  method := GoGetArg(goArgs, 3, MkUndefined()); _ = method;
  responseHeaders := GoGetArg(goArgs, 4, MkUndefined()); _ = responseHeaders;
  responseBody := GoGetArg(goArgs, 5, MkUndefined()); _ = responseBody;
  response := GoGetArg(goArgs, 6, MkUndefined()); _ = response;
  requestHeaders := GoGetArg(goArgs, 7, MkUndefined()); _ = requestHeaders;
  requestBody := GoGetArg(goArgs, 8, MkUndefined()); _ = requestBody;
  if IsTrue(OpEq2(response , MkUndefined() )) {
    return MkUndefined();
  }
  if IsTrue(OpGtEq(statusCode , MkInteger(400) )) {
    feedback := OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , responseBody )); _ = feedback;
    error := this.SafeValue(response , MkString("error") ); _ = error;
    if IsTrue(OpEq2(error , MkUndefined() )) {
      error = OpCopy(response );
    }
    code := this.SafeString2(error , MkString("code") , MkString("status") ); _ = code;
    message := this.SafeString2(error , MkString("message") , MkString("debugMessage") ); _ = message;
    this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("broad") ), message , feedback );
    this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("exact") ), code , feedback );
    this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("exact") ), message , feedback );
    panic(NewExchangeError(feedback ));
  }
  return MkUndefined()
}

