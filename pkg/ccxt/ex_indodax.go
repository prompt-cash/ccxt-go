package ccxt

import (
)

type Indodax struct {
	*ExchangeBase
}

var _ Exchange = (*Indodax)(nil)

func init() {
	exchange := &Indodax{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Indodax) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("indodax") ,
            "name":MkString("INDODAX") ,
            "countries":MkArray(&VarArray{
                MkString("ID") ,
                }),
            "has":MkMap(&VarMap{
                "cancelOrder":MkBool(true) ,
                "CORS":MkBool(false) ,
                "createMarketOrder":MkBool(false) ,
                "createOrder":MkBool(true) ,
                "fetchBalance":MkBool(true) ,
                "fetchClosedOrders":MkBool(true) ,
                "fetchCurrencies":MkBool(false) ,
                "fetchMarkets":MkBool(true) ,
                "fetchMyTrades":MkBool(false) ,
                "fetchOpenOrders":MkBool(true) ,
                "fetchOrder":MkBool(true) ,
                "fetchOrderBook":MkBool(true) ,
                "fetchOrders":MkBool(false) ,
                "fetchTicker":MkBool(true) ,
                "fetchTickers":MkBool(false) ,
                "fetchTime":MkBool(true) ,
                "fetchTrades":MkBool(true) ,
                "withdraw":MkBool(true) ,
                }),
            "version":MkString("2.0") ,
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/51840849/87070508-9358c880-c221-11ea-8dc5-5391afbbb422.jpg") ,
                "api":MkMap(&VarMap{
                    "public":MkString("https://indodax.com/api") ,
                    "private":MkString("https://indodax.com/tapi") ,
                    }),
                "www":MkString("https://www.indodax.com") ,
                "doc":MkString("https://github.com/btcid/indodax-official-api-docs") ,
                "referral":MkString("https://indodax.com/ref/testbitcoincoid/1") ,
                }),
            "api":MkMap(&VarMap{
                "public":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("server_time") ,
                        MkString("pairs") ,
                        MkString("{pair}/ticker") ,
                        MkString("{pair}/trades") ,
                        MkString("{pair}/depth") ,
                        })}),
                "private":MkMap(&VarMap{"post":MkArray(&VarArray{
                        MkString("getInfo") ,
                        MkString("transHistory") ,
                        MkString("trade") ,
                        MkString("tradeHistory") ,
                        MkString("getOrder") ,
                        MkString("openOrders") ,
                        MkString("cancelOrder") ,
                        MkString("orderHistory") ,
                        MkString("withdrawCoin") ,
                        })}),
                }),
            "fees":MkMap(&VarMap{"trading":MkMap(&VarMap{
                    "tierBased":MkBool(false) ,
                    "percentage":MkBool(true) ,
                    "maker":MkInteger(0) ,
                    "taker":MkNumber(0.003) ,
                    })}),
            "exceptions":MkMap(&VarMap{
                "exact":MkMap(&VarMap{
                    "invalid_pair":BadSymbol ,
                    "Insufficient balance.":InsufficientFunds ,
                    "invalid order.":OrderNotFound ,
                    "Invalid credentials. API not found or session has expired.":AuthenticationError ,
                    "Invalid credentials. Bad sign.":AuthenticationError ,
                    }),
                "broad":MkMap(&VarMap{
                    "Minimum price":InvalidOrder ,
                    "Minimum order":InvalidOrder ,
                    }),
                }),
            "options":MkMap(&VarMap{
                "recvWindow":OpMulti(MkInteger(5) , MkInteger(1000) ),
                "timeDifference":MkInteger(0) ,
                "adjustForTimeDifference":MkBool(false) ,
                }),
            "commonCurrencies":MkMap(&VarMap{
                "STR":MkString("XLM") ,
                "BCHABC":MkString("BCH") ,
                "BCHSV":MkString("BSV") ,
                "DRK":MkString("DASH") ,
                "NEM":MkString("XEM") ,
                }),
            }));
}

func (this *Indodax) Nonce(goArgs ...*Variant) *Variant{
  return OpSub(this.Milliseconds(), *(*this.At(MkString("options")) ).At(MkString("timeDifference") ));
}

func (this *Indodax) FetchTime(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetServerTime"), params ); _ = response;
  return this.SafeInteger(response , MkString("server_time") );
}

func (this *Indodax) LoadTimeDifference(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  serverTime := this.FetchTime(params ); _ = serverTime;
  after := this.Milliseconds(); _ = after;
  *(*this.At(MkString("options")) ).At(MkString("timeDifference") )= OpSub(after , serverTime );
  return *(*this.At(MkString("options")) ).At(MkString("timeDifference") );
}

func (this *Indodax) FetchMarkets(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetPairs"), params ); _ = response;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , response.Length )); OpIncr(& i ){
    market := *(response ).At(i ); _ = market;
    id := this.SafeString(market , MkString("ticker_id") ); _ = id;
    baseId := this.SafeString(market , MkString("traded_currency") ); _ = baseId;
    quoteId := this.SafeString(market , MkString("base_currency") ); _ = quoteId;
    base := this.SafeCurrencyCode(baseId ); _ = base;
    quote := this.SafeCurrencyCode(quoteId ); _ = quote;
    symbol := OpAdd(base , OpAdd(MkString("/") , quote )); _ = symbol;
    taker := this.SafeNumber(market , MkString("trade_fee_percent") ); _ = taker;
    isMaintenance := this.SafeInteger(market , MkString("is_maintenance") ); _ = isMaintenance;
    active := OpTrinary(isMaintenance , MkBool(false) , MkBool(true) ); _ = active;
    pricePrecision := this.SafeInteger(market , MkString("price_round") ); _ = pricePrecision;
    precision := MkMap(&VarMap{
          "amount":MkInteger(8) ,
          "price":pricePrecision ,
          }); _ = precision;
    limits := MkMap(&VarMap{
          "amount":MkMap(&VarMap{
              "min":this.SafeNumber(market , MkString("trade_min_traded_currency") ),
              "max":MkUndefined() ,
              }),
          "price":MkMap(&VarMap{
              "min":this.SafeNumber(market , MkString("trade_min_base_currency") ),
              "max":MkUndefined() ,
              }),
          "cost":MkMap(&VarMap{
              "min":MkUndefined() ,
              "max":MkUndefined() ,
              }),
          }); _ = limits;
    result.Push(MkMap(&VarMap{
            "id":id ,
            "symbol":symbol ,
            "base":base ,
            "quote":quote ,
            "baseId":baseId ,
            "quoteId":quoteId ,
            "taker":taker ,
            "percentage":MkBool(true) ,
            "precision":precision ,
            "limits":limits ,
            "info":market ,
            "active":active ,
            }));
  }
  return result ;
}

func (this *Indodax) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("privatePostGetInfo"), params ); _ = response;
  balances := this.SafeValue(response , MkString("return") , MkMap(&VarMap{})); _ = balances;
  free := this.SafeValue(balances , MkString("balance") , MkMap(&VarMap{})); _ = free;
  used := this.SafeValue(balances , MkString("balance_hold") , MkMap(&VarMap{})); _ = used;
  timestamp := this.SafeTimestamp(balances , MkString("server_time") ); _ = timestamp;
  result := MkMap(&VarMap{
        "info":response ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        }); _ = result;
  currencyIds := GoGetKeys(free ); _ = currencyIds;
  for i := MkInteger(0) ; IsTrue(OpLw(i , currencyIds.Length )); OpIncr(& i ){
    currencyId := *(currencyIds ).At(i ); _ = currencyId;
    code := this.SafeCurrencyCode(currencyId ); _ = code;
    account := this.Account(); _ = account;
    *(account ).At(MkString("free") )= this.SafeString(free , currencyId );
    *(account ).At(MkString("used") )= this.SafeString(used , currencyId );
    *(result ).At(code )= OpCopy(account );
  }
  return this.ParseBalance(result );
}

func (this *Indodax) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"pair":this.MarketId(symbol )}); _ = request;
  orderbook := this.Call(MkString("publicGetPairDepth"), this.Extend(request , params )); _ = orderbook;
  return this.ParseOrderBook(orderbook , symbol , MkUndefined() , MkString("buy") , MkString("sell") );
}

func (this *Indodax) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"pair":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("publicGetPairTicker"), this.Extend(request , params )); _ = response;
  ticker := *(response ).At(MkString("ticker") ); _ = ticker;
  timestamp := this.SafeTimestamp(ticker , MkString("server_time") ); _ = timestamp;
  baseVolume := OpAdd(MkString("vol_") , (*(market ).At(MkString("baseId") )).Call(MkString("toLowerCase") )); _ = baseVolume;
  quoteVolume := OpAdd(MkString("vol_") , (*(market ).At(MkString("quoteId") )).Call(MkString("toLowerCase") )); _ = quoteVolume;
  last := this.SafeNumber(ticker , MkString("last") ); _ = last;
  return MkMap(&VarMap{
        "symbol":symbol ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "high":this.SafeNumber(ticker , MkString("high") ),
        "low":this.SafeNumber(ticker , MkString("low") ),
        "bid":this.SafeNumber(ticker , MkString("buy") ),
        "bidVolume":MkUndefined() ,
        "ask":this.SafeNumber(ticker , MkString("sell") ),
        "askVolume":MkUndefined() ,
        "vwap":MkUndefined() ,
        "open":MkUndefined() ,
        "close":last ,
        "last":last ,
        "previousClose":MkUndefined() ,
        "change":MkUndefined() ,
        "percentage":MkUndefined() ,
        "average":MkUndefined() ,
        "baseVolume":this.SafeNumber(ticker , baseVolume ),
        "quoteVolume":this.SafeNumber(ticker , quoteVolume ),
        "info":ticker ,
        });
}

func (this *Indodax) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.SafeTimestamp(trade , MkString("date") ); _ = timestamp;
  id := this.SafeString(trade , MkString("tid") ); _ = id;
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  if IsTrue(OpNotEq2(market , MkUndefined() )) {
    symbol = *(market ).At(MkString("symbol") );
  }
  type_ := OpCopy(MkUndefined() ); _ = type_;
  side := this.SafeString(trade , MkString("type") ); _ = side;
  priceString := this.SafeString(trade , MkString("price") ); _ = priceString;
  amountString := this.SafeString(trade , MkString("amount") ); _ = amountString;
  price := this.ParseNumber(priceString ); _ = price;
  amount := this.ParseNumber(amountString ); _ = amount;
  cost := this.ParseNumber(Precise.StringMul(priceString , amountString )); _ = cost;
  return MkMap(&VarMap{
        "id":id ,
        "info":trade ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "symbol":symbol ,
        "type":type_ ,
        "side":side ,
        "order":MkUndefined() ,
        "takerOrMaker":MkUndefined() ,
        "price":price ,
        "amount":amount ,
        "cost":cost ,
        "fee":MkUndefined() ,
        });
}

func (this *Indodax) FetchTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"pair":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("publicGetPairTrades"), this.Extend(request , params )); _ = response;
  return this.ParseTrades(response , market , since , limit );
}

func (this *Indodax) ParseOrderStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "open":MkString("open") ,
        "filled":MkString("closed") ,
        "cancelled":MkString("canceled") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Indodax) ParseOrder(goArgs ...*Variant) *Variant{
  order := GoGetArg(goArgs, 0, MkUndefined()); _ = order;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  side := OpCopy(MkUndefined() ); _ = side;
  if IsTrue(OpHasMember(MkString("type") , order )) {
    side = *(order ).At(MkString("type") );
  }
  status := this.ParseOrderStatus(this.SafeString(order , MkString("status") , MkString("open") )); _ = status;
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  cost := OpCopy(MkUndefined() ); _ = cost;
  price := this.SafeNumber(order , MkString("price") ); _ = price;
  amount := OpCopy(MkUndefined() ); _ = amount;
  remaining := OpCopy(MkUndefined() ); _ = remaining;
  if IsTrue(OpNotEq2(market , MkUndefined() )) {
    symbol = *(market ).At(MkString("symbol") );
    quoteId := *(market ).At(MkString("quoteId") ); _ = quoteId;
    baseId := *(market ).At(MkString("baseId") ); _ = baseId;
    if IsTrue(OpAnd(OpEq2(*(market ).At(MkString("quoteId") ), MkString("idr") ), OpHasMember(MkString("order_rp") , order ))) {
      quoteId = MkString("rp") ;
    }
    if IsTrue(OpAnd(OpEq2(*(market ).At(MkString("baseId") ), MkString("idr") ), OpHasMember(MkString("remain_rp") , order ))) {
      baseId = MkString("rp") ;
    }
    cost = this.SafeNumber(order , OpAdd(MkString("order_") , quoteId ));
    if IsTrue(OpNot(cost )) {
      amount = this.SafeNumber(order , OpAdd(MkString("order_") , baseId ));
      remaining = this.SafeNumber(order , OpAdd(MkString("remain_") , baseId ));
    }
  }
  timestamp := this.SafeInteger(order , MkString("submit_time") ); _ = timestamp;
  fee := OpCopy(MkUndefined() ); _ = fee;
  id := this.SafeString(order , MkString("order_id") ); _ = id;
  return this.SafeOrder(MkMap(&VarMap{
            "info":order ,
            "id":id ,
            "clientOrderId":MkUndefined() ,
            "timestamp":timestamp ,
            "datetime":this.Iso8601(timestamp ),
            "lastTradeTimestamp":MkUndefined() ,
            "symbol":symbol ,
            "type":MkString("limit") ,
            "timeInForce":MkUndefined() ,
            "postOnly":MkUndefined() ,
            "side":side ,
            "price":price ,
            "stopPrice":MkUndefined() ,
            "cost":cost ,
            "average":MkUndefined() ,
            "amount":amount ,
            "filled":MkUndefined() ,
            "remaining":remaining ,
            "status":status ,
            "fee":fee ,
            "trades":MkUndefined() ,
            }));
}

func (this *Indodax) FetchOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchOrder() requires a symbol") )));
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "pair":*(market ).At(MkString("id") ),
        "order_id":id ,
        }); _ = request;
  response := this.Call(MkString("privatePostGetOrder"), this.Extend(request , params )); _ = response;
  orders := *(response ).At(MkString("return") ); _ = orders;
  order := this.ParseOrder(this.Extend(MkMap(&VarMap{"id":id }), *(orders ).At(MkString("order") )), market ); _ = order;
  return this.Extend(MkMap(&VarMap{"info":response }), order );
}

func (this *Indodax) FetchOpenOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
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
  response := this.Call(MkString("privatePostOpenOrders"), this.Extend(request , params )); _ = response;
  rawOrders := *(*(response ).At(MkString("return") )).At(MkString("orders") ); _ = rawOrders;
  if IsTrue(OpNot(rawOrders )) {
    return MkArray(&VarArray{});
  }
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    return this.ParseOrders(rawOrders , market , since , limit );
  }
  marketIds := GoGetKeys(rawOrders ); _ = marketIds;
  exchangeOrders := MkArray(&VarArray{}); _ = exchangeOrders;
  for i := MkInteger(0) ; IsTrue(OpLw(i , marketIds.Length )); OpIncr(& i ){
    marketId := *(marketIds ).At(i ); _ = marketId;
    marketOrders := *(rawOrders ).At(marketId ); _ = marketOrders;
    market = *(*this.At(MkString("markets_by_id")) ).At(marketId );
    parsedOrders := this.ParseOrders(marketOrders , market , since , limit ); _ = parsedOrders;
    exchangeOrders = this.ArrayConcat(exchangeOrders , parsedOrders );
  }
  return exchangeOrders ;
}

func (this *Indodax) FetchClosedOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchOrders() requires a symbol argument") )));
  }
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("pair") )= *(market ).At(MkString("id") );
  }
  response := this.Call(MkString("privatePostOrderHistory"), this.Extend(request , params )); _ = response;
  orders := this.ParseOrders(*(*(response ).At(MkString("return") )).At(MkString("orders") ), market ); _ = orders;
  orders = this.FilterBy(orders , MkString("status") , MkString("closed") );
  return this.FilterBySymbolSinceLimit(orders , symbol , since , limit );
}

func (this *Indodax) CreateOrder(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 2, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 3, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 4, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 5, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpNotEq2(type_ , MkString("limit") )) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , MkString(" allows limit orders only") )));
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "pair":*(market ).At(MkString("id") ),
        "type":side ,
        "price":price ,
        }); _ = request;
  currency := *(market ).At(MkString("baseId") ); _ = currency;
  if IsTrue(OpEq2(side , MkString("buy") )) {
    *(request ).At(*(market ).At(MkString("quoteId") ))= OpMulti(amount , price );
  } else {
    *(request ).At(*(market ).At(MkString("baseId") ))= OpCopy(amount );
  }
  *(request ).At(currency )= OpCopy(amount );
  result := this.Call(MkString("privatePostTrade"), this.Extend(request , params )); _ = result;
  data := this.SafeValue(result , MkString("return") , MkMap(&VarMap{})); _ = data;
  id := this.SafeString(data , MkString("order_id") ); _ = id;
  return MkMap(&VarMap{
        "info":result ,
        "id":id ,
        });
}

func (this *Indodax) CancelOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" cancelOrder() requires a symbol argument") )));
  }
  side := this.SafeValue(params , MkString("side") ); _ = side;
  if IsTrue(OpEq2(side , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" cancelOrder() requires an extra \"side\" param") )));
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "order_id":id ,
        "pair":*(market ).At(MkString("id") ),
        "type":side ,
        }); _ = request;
  return this.Call(MkString("privatePostCancelOrder"), this.Extend(request , params ));
}

func (this *Indodax) Withdraw(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  amount := GoGetArg(goArgs, 1, MkUndefined()); _ = amount;
  address := GoGetArg(goArgs, 2, MkUndefined()); _ = address;
  tag := GoGetArg(goArgs, 3, MkUndefined() ); _ = tag;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.CheckAddress(address );
  this.LoadMarkets();
  currency := this.Currency(code ); _ = currency;
  requestId := this.Milliseconds(); _ = requestId;
  request := MkMap(&VarMap{
        "currency":*(currency ).At(MkString("id") ),
        "withdraw_amount":amount ,
        "withdraw_address":address ,
        "request_id":requestId.ToString(),
        }); _ = request;
  if IsTrue(tag ) {
    *(request ).At(MkString("withdraw_memo") )= OpCopy(tag );
  }
  response := this.Call(MkString("privatePostWithdrawCoin"), this.Extend(request , params )); _ = response;
  id := OpCopy(MkUndefined() ); _ = id;
  if IsTrue(OpAnd(OpHasMember(MkString("txid") , response ), OpGt(*(*(response ).At(MkString("txid") )).At(MkString("length") ), MkInteger(0) ))) {
    id = *(response ).At(MkString("txid") );
  }
  return MkMap(&VarMap{
        "info":response ,
        "id":id ,
        });
}

func (this *Indodax) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("public") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  url := *(*(*this.At(MkString("urls")) ).At(MkString("api") )).At(api ); _ = url;
  if IsTrue(OpEq2(api , MkString("public") )) {
    OpAddAssign(& url , OpAdd(MkString("/") , this.ImplodeParams(path , params )));
  } else {
    this.CheckRequiredCredentials();
    body = this.Urlencode(this.Extend(MkMap(&VarMap{
                "method":path ,
                "timestamp":this.Nonce(),
                "recvWindow":*(*this.At(MkString("options")) ).At(MkString("recvWindow") ),
                }), params ));
    headers = MkMap(&VarMap{
        "Content-Type":MkString("application/x-www-form-urlencoded") ,
        "Key":*this.At(MkString("apiKey")) ,
        "Sign":this.Hmac(this.Encode(body ), this.Encode(*this.At(MkString("secret")) ), MkString("sha512") ),
        });
  }
  return MkMap(&VarMap{
        "url":url ,
        "method":method ,
        "body":body ,
        "headers":headers ,
        });
}

func (this *Indodax) HandleErrors(goArgs ...*Variant) *Variant{
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
  if IsTrue(GoIsArray(response )) {
    return MkUndefined();
  }
  error := this.SafeValue(response , MkString("error") , MkString("") ); _ = error;
  if IsTrue(OpAnd(OpNot(OpHasMember(MkString("success") , response )), OpEq2(error , MkString("") ))) {
    return MkUndefined();
  }
  if IsTrue(OpEq2(this.SafeInteger(response , MkString("success") , MkInteger(0) ), MkInteger(1) )) {
    if IsTrue(OpNot(OpHasMember(MkString("return") , response ))) {
      panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(": malformed response: ") , this.Json(response )))));
    } else {
      return MkUndefined();
    }
  }
  feedback := OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , body )); _ = feedback;
  this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("exact") ), error , feedback );
  this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("broad") ), error , feedback );
  panic(NewExchangeError(feedback ));
  return MkUndefined()
}

