package ccxt

import (
)

type Tidebit struct {
	*ExchangeBase
}

var _ Exchange = (*Tidebit)(nil)

func init() {
	exchange := &Tidebit{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Tidebit) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("tidebit") ,
            "name":MkString("TideBit") ,
            "countries":MkArray(&VarArray{
                MkString("HK") ,
                }),
            "rateLimit":MkInteger(1000) ,
            "version":MkString("v2") ,
            "has":MkMap(&VarMap{
                "cancelOrder":MkBool(true) ,
                "CORS":MkBool(false) ,
                "createOrder":MkBool(true) ,
                "fetchBalance":MkBool(true) ,
                "fetchDepositAddress":MkBool(true) ,
                "fetchMarkets":MkBool(true) ,
                "fetchOHLCV":MkBool(true) ,
                "fetchOrderBook":MkBool(true) ,
                "fetchTicker":MkBool(true) ,
                "fetchTickers":MkBool(true) ,
                "fetchTrades":MkBool(true) ,
                "withdraw":MkBool(true) ,
                }),
            "timeframes":MkMap(&VarMap{
                "1m":MkString("1") ,
                "5m":MkString("5") ,
                "15m":MkString("15") ,
                "30m":MkString("30") ,
                "1h":MkString("60") ,
                "2h":MkString("120") ,
                "4h":MkString("240") ,
                "12h":MkString("720") ,
                "1d":MkString("1440") ,
                "3d":MkString("4320") ,
                "1w":MkString("10080") ,
                }),
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/51840849/87460811-1e690280-c616-11ea-8652-69f187305add.jpg") ,
                "api":MkString("https://www.tidebit.com") ,
                "www":MkString("https://www.tidebit.com") ,
                "doc":MkArray(&VarArray{
                    MkString("https://www.tidebit.com/documents/api/guide") ,
                    MkString("https://www.tidebit.com/swagger/#/default") ,
                    }),
                "referral":MkString("http://bit.ly/2IX0LrM") ,
                }),
            "api":MkMap(&VarMap{
                "public":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("markets") ,
                        MkString("tickers") ,
                        MkString("tickers/{market}") ,
                        MkString("timestamp") ,
                        MkString("trades") ,
                        MkString("trades/{market}") ,
                        MkString("order_book") ,
                        MkString("order") ,
                        MkString("k_with_pending_trades") ,
                        MkString("k") ,
                        MkString("depth") ,
                        }),
                    "post":MkArray(&VarArray{}),
                    }),
                "private":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("addresses/{address}") ,
                        MkString("deposits/history") ,
                        MkString("deposits/get_deposit") ,
                        MkString("deposits/deposit_address") ,
                        MkString("historys/orders") ,
                        MkString("historys/vouchers") ,
                        MkString("historys/accounts") ,
                        MkString("historys/snapshots") ,
                        MkString("linkage/get_status") ,
                        MkString("members/me") ,
                        MkString("order") ,
                        MkString("orders") ,
                        MkString("partners/orders/{id}/trades") ,
                        MkString("referral_commissions/get_undeposited") ,
                        MkString("referral_commissions/get_graph_data") ,
                        MkString("trades/my") ,
                        MkString("withdraws/bind_account_list") ,
                        MkString("withdraws/get_withdraw_account") ,
                        MkString("withdraws/fetch_bind_info") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("deposits/deposit_cash") ,
                        MkString("favorite_markets/update") ,
                        MkString("order/delete") ,
                        MkString("orders") ,
                        MkString("orders/multi") ,
                        MkString("orders/clear") ,
                        MkString("referral_commissions/deposit") ,
                        MkString("withdraws/apply") ,
                        MkString("withdraws/bind_bank") ,
                        MkString("withdraws/bind_address") ,
                        }),
                    }),
                }),
            "fees":MkMap(&VarMap{
                "trading":MkMap(&VarMap{
                    "tierBased":MkBool(false) ,
                    "percentage":MkBool(true) ,
                    "maker":this.ParseNumber(MkString("0.003") ),
                    "taker":this.ParseNumber(MkString("0.003") ),
                    }),
                "funding":MkMap(&VarMap{
                    "tierBased":MkBool(false) ,
                    "percentage":MkBool(true) ,
                    "withdraw":MkMap(&VarMap{}),
                    }),
                }),
            "exceptions":MkMap(&VarMap{
                "2002":InsufficientFunds ,
                "2003":OrderNotFound ,
                }),
            }));
}

func (this *Tidebit) FetchDepositAddress(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  currency := this.Currency(code ); _ = currency;
  request := MkMap(&VarMap{"currency":*(currency ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("privateGetDepositAddress"), this.Extend(request , params )); _ = response;
  if IsTrue(OpHasMember(MkString("success") , response )) {
    if IsTrue(*(response ).At(MkString("success") )) {
      address := this.SafeString(response , MkString("address") ); _ = address;
      tag := this.SafeString(response , MkString("addressTag") ); _ = tag;
      return MkMap(&VarMap{
            "currency":code ,
            "address":this.CheckAddress(address ),
            "tag":tag ,
            "info":response ,
            });
    }
  }
  return MkUndefined()
}

func (this *Tidebit) FetchMarkets(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetMarkets"), params ); _ = response;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , response.Length )); OpIncr(& i ){
    market := *(response ).At(i ); _ = market;
    id := this.SafeString(market , MkString("id") ); _ = id;
    symbol := this.SafeString(market , MkString("name") ); _ = symbol;
    baseId := MkUndefined(); _ = baseId;
    quoteId := MkUndefined(); _ = quoteId;
    ArrayUnpack(symbol.Split(MkString("/") ), &baseId, &quoteId);
    base := this.SafeCurrencyCode(baseId ); _ = base;
    quote := this.SafeCurrencyCode(quoteId ); _ = quote;
    result.Push(MkMap(&VarMap{
            "id":id ,
            "symbol":symbol ,
            "base":base ,
            "quote":quote ,
            "baseId":baseId ,
            "quoteId":quoteId ,
            "info":market ,
            "active":MkUndefined() ,
            "precision":*this.At(MkString("precision")) ,
            "limits":*this.At(MkString("limits")) ,
            }));
  }
  return result ;
}

func (this *Tidebit) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("privateGetMembersMe"), params ); _ = response;
  balances := this.SafeValue(response , MkString("accounts") ); _ = balances;
  result := MkMap(&VarMap{"info":balances }); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , balances.Length )); OpIncr(& i ){
    balance := *(balances ).At(i ); _ = balance;
    currencyId := this.SafeString(balance , MkString("currency") ); _ = currencyId;
    code := this.SafeCurrencyCode(currencyId ); _ = code;
    account := this.Account(); _ = account;
    *(account ).At(MkString("free") )= this.SafeString(balance , MkString("balance") );
    *(account ).At(MkString("used") )= this.SafeString(balance , MkString("locked") );
    *(result ).At(code )= OpCopy(account );
  }
  return this.ParseBalance(result );
}

func (this *Tidebit) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"market":*(market ).At(MkString("id") )}); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  *(request ).At(MkString("market") )= *(market ).At(MkString("id") );
  response := this.Call(MkString("publicGetDepth"), this.Extend(request , params )); _ = response;
  timestamp := this.SafeTimestamp(response , MkString("timestamp") ); _ = timestamp;
  return this.ParseOrderBook(response , symbol , timestamp );
}

func (this *Tidebit) ParseTicker(goArgs ...*Variant) *Variant{
  ticker := GoGetArg(goArgs, 0, MkUndefined()); _ = ticker;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.SafeTimestamp(ticker , MkString("at") ); _ = timestamp;
  ticker = this.SafeValue(ticker , MkString("ticker") , MkMap(&VarMap{}));
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  if IsTrue(OpNotEq2(market , MkUndefined() )) {
    symbol = *(market ).At(MkString("symbol") );
  }
  last := this.SafeNumber(ticker , MkString("last") ); _ = last;
  return MkMap(&VarMap{
        "symbol":symbol ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "high":this.SafeNumber(ticker , MkString("high") ),
        "low":this.SafeNumber(ticker , MkString("low") ),
        "bid":this.SafeNumber(ticker , MkString("buy") ),
        "ask":this.SafeNumber(ticker , MkString("sell") ),
        "bidVolume":MkUndefined() ,
        "askVolume":MkUndefined() ,
        "vwap":MkUndefined() ,
        "open":MkUndefined() ,
        "close":last ,
        "last":last ,
        "change":MkUndefined() ,
        "percentage":MkUndefined() ,
        "previousClose":MkUndefined() ,
        "average":MkUndefined() ,
        "baseVolume":this.SafeNumber(ticker , MkString("vol") ),
        "quoteVolume":MkUndefined() ,
        "info":ticker ,
        });
}

func (this *Tidebit) FetchTickers(goArgs ...*Variant) *Variant{
  symbols := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbols;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  tickers := this.Call(MkString("publicGetTickers"), params ); _ = tickers;
  ids := GoGetKeys(tickers ); _ = ids;
  result := MkMap(&VarMap{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , ids.Length )); OpIncr(& i ){
    id := *(ids ).At(i ); _ = id;
    market := this.SafeMarket(id ); _ = market;
    symbol := *(market ).At(MkString("symbol") ); _ = symbol;
    ticker := *(tickers ).At(id ); _ = ticker;
    *(result ).At(symbol )= this.ParseTicker(ticker , market );
  }
  return this.FilterByArray(result , MkString("symbol") , symbols );
}

func (this *Tidebit) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"market":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("publicGetTickersMarket"), this.Extend(request , params )); _ = response;
  return this.ParseTicker(response , market );
}

func (this *Tidebit) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.Parse8601(this.SafeString(trade , MkString("created_at") )); _ = timestamp;
  id := this.SafeString(trade , MkString("id") ); _ = id;
  priceString := this.SafeString(trade , MkString("price") ); _ = priceString;
  amountString := this.SafeString(trade , MkString("volume") ); _ = amountString;
  price := this.ParseNumber(priceString ); _ = price;
  amount := this.ParseNumber(amountString ); _ = amount;
  cost := this.SafeNumber(trade , MkString("funds") ); _ = cost;
  if IsTrue(OpEq2(cost , MkUndefined() )) {
    cost = this.ParseNumber(Precise.StringMul(priceString , amountString ));
  }
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
        "type":MkUndefined() ,
        "side":MkUndefined() ,
        "order":MkUndefined() ,
        "takerOrMaker":MkUndefined() ,
        "price":price ,
        "amount":amount ,
        "cost":cost ,
        "fee":MkUndefined() ,
        });
}

func (this *Tidebit) FetchTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"market":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("publicGetTrades"), this.Extend(request , params )); _ = response;
  return this.ParseTrades(response , market , since , limit );
}

func (this *Tidebit) ParseOHLCV(goArgs ...*Variant) *Variant{
  ohlcv := GoGetArg(goArgs, 0, MkUndefined()); _ = ohlcv;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  return MkArray(&VarArray{
        this.SafeTimestamp(ohlcv , MkInteger(0) ),
        this.SafeNumber(ohlcv , MkInteger(1) ),
        this.SafeNumber(ohlcv , MkInteger(2) ),
        this.SafeNumber(ohlcv , MkInteger(3) ),
        this.SafeNumber(ohlcv , MkInteger(4) ),
        this.SafeNumber(ohlcv , MkInteger(5) ),
        });
}

func (this *Tidebit) FetchOHLCV(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  timeframe := GoGetArg(goArgs, 1, MkString("1m") ); _ = timeframe;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  if IsTrue(OpEq2(limit , MkUndefined() )) {
    limit = MkInteger(30) ;
  }
  request := MkMap(&VarMap{
        "market":*(market ).At(MkString("id") ),
        "period":*(*this.At(MkString("timeframes")) ).At(timeframe ),
        "limit":limit ,
        }); _ = request;
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("timestamp") )= parseInt(OpDiv(since , MkInteger(1000) ));
  } else {
    *(request ).At(MkString("timestamp") )= MkInteger(1800000) ;
  }
  response := this.Call(MkString("publicGetK"), this.Extend(request , params )); _ = response;
  if IsTrue(OpEq2(response , MkString("null") )) {
    return MkArray(&VarArray{});
  }
  return this.ParseOHLCVs(response , market , timeframe , since , limit );
}

func (this *Tidebit) ParseOrderStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "done":MkString("closed") ,
        "wait":MkString("open") ,
        "cancel":MkString("canceled") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Tidebit) ParseOrder(goArgs ...*Variant) *Variant{
  order := GoGetArg(goArgs, 0, MkUndefined()); _ = order;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  marketId := this.SafeString(order , MkString("market") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market ); _ = symbol;
  timestamp := this.Parse8601(this.SafeString(order , MkString("created_at") )); _ = timestamp;
  status := this.ParseOrderStatus(this.SafeString(order , MkString("state") )); _ = status;
  id := this.SafeString(order , MkString("id") ); _ = id;
  type_ := this.SafeString(order , MkString("ord_type") ); _ = type_;
  side := this.SafeString(order , MkString("side") ); _ = side;
  price := this.SafeNumber(order , MkString("price") ); _ = price;
  amount := this.SafeNumber(order , MkString("volume") ); _ = amount;
  filled := this.SafeNumber(order , MkString("executed_volume") ); _ = filled;
  remaining := this.SafeNumber(order , MkString("remaining_volume") ); _ = remaining;
  average := this.SafeNumber(order , MkString("avg_price") ); _ = average;
  return this.SafeOrder(MkMap(&VarMap{
            "id":id ,
            "clientOrderId":MkUndefined() ,
            "timestamp":timestamp ,
            "datetime":this.Iso8601(timestamp ),
            "lastTradeTimestamp":MkUndefined() ,
            "status":status ,
            "symbol":symbol ,
            "type":type_ ,
            "timeInForce":MkUndefined() ,
            "postOnly":MkUndefined() ,
            "side":side ,
            "price":price ,
            "stopPrice":MkUndefined() ,
            "amount":amount ,
            "filled":filled ,
            "remaining":remaining ,
            "cost":MkUndefined() ,
            "trades":MkUndefined() ,
            "fee":MkUndefined() ,
            "info":order ,
            "average":average ,
            }));
}

func (this *Tidebit) CreateOrder(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 2, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 3, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 4, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 5, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{
        "market":this.MarketId(symbol ),
        "side":side ,
        "volume":amount.ToString(),
        "ord_type":type_ ,
        }); _ = request;
  if IsTrue(OpEq2(type_ , MkString("limit") )) {
    *(request ).At(MkString("price") )= price.ToString();
  }
  response := this.Call(MkString("privatePostOrders"), this.Extend(request , params )); _ = response;
  return this.ParseOrder(response );
}

func (this *Tidebit) CancelOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"id":id }); _ = request;
  result := this.Call(MkString("privatePostOrderDelete"), this.Extend(request , params )); _ = result;
  order := this.ParseOrder(result ); _ = order;
  status := this.SafeString(order , MkString("status") ); _ = status;
  if IsTrue(OpOr(OpEq2(status , MkString("closed") ), OpEq2(status , MkString("canceled") ))) {
    panic(NewOrderNotFound(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , this.Json(order )))));
  }
  return order ;
}

func (this *Tidebit) Withdraw(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  amount := GoGetArg(goArgs, 1, MkUndefined()); _ = amount;
  address := GoGetArg(goArgs, 2, MkUndefined()); _ = address;
  tag := GoGetArg(goArgs, 3, MkUndefined() ); _ = tag;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.CheckAddress(address );
  this.LoadMarkets();
  currency := this.Currency(code ); _ = currency;
  id := this.SafeString(params , MkString("id") ); _ = id;
  if IsTrue(OpEq2(id , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" withdraw() requires an extra `id` param (withdraw account id according to withdraws/bind_account_list endpoint") )));
  }
  request := MkMap(&VarMap{
        "id":id ,
        "currency_type":MkString("coin") ,
        "currency":*(currency ).At(MkString("id") ),
        "body":amount ,
        }); _ = request;
  if IsTrue(OpNotEq2(tag , MkUndefined() )) {
    *(request ).At(MkString("memo") )= OpCopy(tag );
  }
  result := this.Call(MkString("privatePostWithdrawsApply"), this.Extend(request , params )); _ = result;
  return MkMap(&VarMap{
        "info":result ,
        "id":MkUndefined() ,
        });
}

func (this *Tidebit) Nonce(goArgs ...*Variant) *Variant{
  return this.Milliseconds();
}

func (this *Tidebit) EncodeParams(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkUndefined()); _ = params;
  return this.Urlencode(this.Keysort(params ));
}

func (this *Tidebit) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("public") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  request := OpAdd(MkString("/") , OpAdd(MkString("api/") , OpAdd(*this.At(MkString("version")) , OpAdd(MkString("/") , OpAdd(this.ImplodeParams(path , params ), MkString(".json") ))))); _ = request;
  query := this.Omit(params , this.ExtractParams(path )); _ = query;
  url := OpAdd(*(*this.At(MkString("urls")) ).At(MkString("api") ), request ); _ = url;
  if IsTrue(OpEq2(api , MkString("public") )) {
    if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
      OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(query )));
    }
  } else {
    this.CheckRequiredCredentials();
    nonce := (this.Nonce()).Call(MkString("toString") ); _ = nonce;
    sortedByKey := this.Keysort(this.Extend(MkMap(&VarMap{
                  "access_key":*this.At(MkString("apiKey")) ,
                  "tonce":nonce ,
                  }), params )); _ = sortedByKey;
    query := this.Urlencode(sortedByKey ); _ = query;
    payload := OpAdd(method , OpAdd(MkString("|") , OpAdd(request , OpAdd(MkString("|") , query )))); _ = payload;
    signature := this.Hmac(this.Encode(payload ), this.Encode(*this.At(MkString("secret")) )); _ = signature;
    suffix := OpAdd(query , OpAdd(MkString("&signature=") , signature )); _ = suffix;
    if IsTrue(OpEq2(method , MkString("GET") )) {
      OpAddAssign(& url , OpAdd(MkString("?") , suffix ));
    } else {
      body = OpCopy(suffix );
      headers = MkMap(&VarMap{"Content-Type":MkString("application/x-www-form-urlencoded") });
    }
  }
  return MkMap(&VarMap{
        "url":url ,
        "method":method ,
        "body":body ,
        "headers":headers ,
        });
}

func (this *Tidebit) HandleErrors(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  reason := GoGetArg(goArgs, 1, MkUndefined()); _ = reason;
  url := GoGetArg(goArgs, 2, MkUndefined()); _ = url;
  method := GoGetArg(goArgs, 3, MkUndefined()); _ = method;
  headers := GoGetArg(goArgs, 4, MkUndefined()); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined()); _ = body;
  response := GoGetArg(goArgs, 6, MkUndefined()); _ = response;
  requestHeaders := GoGetArg(goArgs, 7, MkUndefined()); _ = requestHeaders;
  requestBody := GoGetArg(goArgs, 8, MkUndefined()); _ = requestBody;
  if IsTrue(OpOr(OpEq2(code , MkInteger(400) ), OpEq2(response , MkUndefined() ))) {
    feedback := OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , body )); _ = feedback;
    if IsTrue(OpEq2(response , MkUndefined() )) {
      panic(NewExchangeError(feedback ));
    }
    error := this.SafeValue(response , MkString("error") , MkMap(&VarMap{})); _ = error;
    errorCode := this.SafeString(error , MkString("code") ); _ = errorCode;
    this.ThrowExactlyMatchedException(*this.At(MkString("exceptions")) , errorCode , feedback );
  }
  return MkUndefined()
}

