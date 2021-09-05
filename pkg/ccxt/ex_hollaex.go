package ccxt

import (
)

type Hollaex struct {
	*ExchangeBase
}

var _ Exchange = (*Hollaex)(nil)

func init() {
	exchange := &Hollaex{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Hollaex) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("hollaex") ,
            "name":MkString("HollaEx") ,
            "countries":MkArray(&VarArray{
                MkString("KR") ,
                }),
            "rateLimit":MkInteger(333) ,
            "version":MkString("v2") ,
            "has":MkMap(&VarMap{
                "CORS":MkBool(false) ,
                "cancelAllOrders":MkBool(true) ,
                "cancelOrder":MkBool(true) ,
                "createLimitBuyOrder":MkBool(true) ,
                "createLimitSellOrder":MkBool(true) ,
                "createMarketBuyOrder":MkBool(true) ,
                "createMarketSellOrder":MkBool(true) ,
                "createOrder":MkBool(true) ,
                "fetchBalance":MkBool(true) ,
                "fetchClosedOrders":MkBool(true) ,
                "fetchCurrencies":MkBool(true) ,
                "fetchDepositAddress":MkString("emulated") ,
                "fetchDeposits":MkBool(true) ,
                "fetchMarkets":MkBool(true) ,
                "fetchMyTrades":MkBool(true) ,
                "fetchOHLCV":MkBool(true) ,
                "fetchOpenOrder":MkBool(true) ,
                "fetchOpenOrders":MkBool(true) ,
                "fetchOrder":MkBool(true) ,
                "fetchOrderBook":MkBool(true) ,
                "fetchOrderBooks":MkBool(true) ,
                "fetchOrders":MkBool(true) ,
                "fetchTicker":MkBool(true) ,
                "fetchTickers":MkBool(true) ,
                "fetchTrades":MkBool(true) ,
                "fetchTransactions":MkBool(false) ,
                "fetchWithdrawals":MkBool(true) ,
                "withdraw":MkBool(true) ,
                "fetchDepositAddresses":MkBool(true) ,
                }),
            "timeframes":MkMap(&VarMap{
                "1h":MkString("1h") ,
                "1d":MkString("1d") ,
                }),
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/1294454/75841031-ca375180-5ddd-11ea-8417-b975674c23cb.jpg") ,
                "api":MkString("https://api.hollaex.com") ,
                "www":MkString("https://hollaex.com") ,
                "doc":MkString("https://apidocs.hollaex.com") ,
                "referral":MkString("https://pro.hollaex.com/signup?affiliation_code=QSWA6G") ,
                }),
            "precisionMode":TICK_SIZE ,
            "requiredCredentials":MkMap(&VarMap{
                "apiKey":MkBool(true) ,
                "secret":MkBool(true) ,
                }),
            "api":MkMap(&VarMap{
                "public":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("health") ,
                        MkString("constants") ,
                        MkString("kit") ,
                        MkString("tiers") ,
                        MkString("ticker") ,
                        MkString("tickers") ,
                        MkString("orderbook") ,
                        MkString("orderbooks") ,
                        MkString("trades") ,
                        MkString("chart") ,
                        MkString("charts") ,
                        MkString("udf/config") ,
                        MkString("udf/history") ,
                        MkString("udf/symbols") ,
                        })}),
                "private":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("user") ,
                        MkString("user/balance") ,
                        MkString("user/deposits") ,
                        MkString("user/withdrawals") ,
                        MkString("user/withdrawal/fee") ,
                        MkString("user/trades") ,
                        MkString("orders") ,
                        MkString("orders/{order_id}") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("user/request-withdrawal") ,
                        MkString("order") ,
                        }),
                    "delete":MkArray(&VarArray{
                        MkString("order/all") ,
                        MkString("order") ,
                        }),
                    }),
                }),
            "fees":MkMap(&VarMap{"trading":MkMap(&VarMap{
                    "tierBased":MkBool(true) ,
                    "percentage":MkBool(true) ,
                    "taker":MkNumber(0.001) ,
                    "maker":MkNumber(0.001) ,
                    })}),
            "exceptions":MkMap(&VarMap{
                "broad":MkMap(&VarMap{
                    "Invalid token":AuthenticationError ,
                    "Order not found":OrderNotFound ,
                    "Insufficient balance":InsufficientFunds ,
                    }),
                "exact":MkMap(&VarMap{
                    "400":BadRequest ,
                    "403":AuthenticationError ,
                    "404":BadRequest ,
                    "405":BadRequest ,
                    "410":BadRequest ,
                    "429":BadRequest ,
                    "500":NetworkError ,
                    "503":NetworkError ,
                    }),
                }),
            "options":MkMap(&VarMap{"api-expires":parseInt(OpDiv(*this.At(MkString("timeout")) , MkInteger(1000) ))}),
            }));
}

func (this *Hollaex) FetchMarkets(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetConstants"), params ); _ = response;
  pairs := this.SafeValue(response , MkString("pairs") , MkMap(&VarMap{})); _ = pairs;
  keys := GoGetKeys(pairs ); _ = keys;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , keys.Length )); OpIncr(& i ){
    key := *(keys ).At(i ); _ = key;
    market := *(pairs ).At(key ); _ = market;
    id := this.SafeString(market , MkString("name") ); _ = id;
    baseId := this.SafeString(market , MkString("pair_base") ); _ = baseId;
    quoteId := this.SafeString(market , MkString("pair_2") ); _ = quoteId;
    base := this.CommonCurrencyCode(baseId.ToUpperCase()); _ = base;
    quote := this.CommonCurrencyCode(quoteId.ToUpperCase()); _ = quote;
    symbol := OpAdd(base , OpAdd(MkString("/") , quote )); _ = symbol;
    active := this.SafeValue(market , MkString("active") ); _ = active;
    maker := *(*(*this.At(MkString("fees")) ).At(MkString("trading") )).At(MkString("maker") ); _ = maker;
    taker := *(*(*this.At(MkString("fees")) ).At(MkString("trading") )).At(MkString("taker") ); _ = taker;
    result.Push(MkMap(&VarMap{
            "id":id ,
            "symbol":symbol ,
            "base":base ,
            "quote":quote ,
            "baseId":baseId ,
            "quoteId":quoteId ,
            "active":active ,
            "precision":MkMap(&VarMap{
                "price":this.SafeNumber(market , MkString("increment_price") ),
                "amount":this.SafeNumber(market , MkString("increment_size") ),
                }),
            "limits":MkMap(&VarMap{
                "amount":MkMap(&VarMap{
                    "min":this.SafeNumber(market , MkString("min_size") ),
                    "max":this.SafeNumber(market , MkString("max_size") ),
                    }),
                "price":MkMap(&VarMap{
                    "min":this.SafeNumber(market , MkString("min_price") ),
                    "max":this.SafeNumber(market , MkString("max_price") ),
                    }),
                "cost":MkMap(&VarMap{
                    "min":MkUndefined() ,
                    "max":MkUndefined() ,
                    }),
                }),
            "taker":taker ,
            "maker":maker ,
            "info":market ,
            }));
  }
  return result ;
}

func (this *Hollaex) FetchCurrencies(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetConstants"), params ); _ = response;
  coins := this.SafeValue(response , MkString("coins") , MkMap(&VarMap{})); _ = coins;
  keys := GoGetKeys(coins ); _ = keys;
  result := MkMap(&VarMap{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , keys.Length )); OpIncr(& i ){
    key := *(keys ).At(i ); _ = key;
    currency := *(coins ).At(key ); _ = currency;
    id := this.SafeString(currency , MkString("symbol") ); _ = id;
    numericId := this.SafeInteger(currency , MkString("id") ); _ = numericId;
    code := this.SafeCurrencyCode(id ); _ = code;
    name := this.SafeString(currency , MkString("fullname") ); _ = name;
    active := this.SafeValue(currency , MkString("active") ); _ = active;
    fee := this.SafeNumber(currency , MkString("withdrawal_fee") ); _ = fee;
    precision := this.SafeNumber(currency , MkString("increment_unit") ); _ = precision;
    withdrawalLimits := this.SafeValue(currency , MkString("withdrawal_limits") , MkArray(&VarArray{})); _ = withdrawalLimits;
    *(result ).At(code )= MkMap(&VarMap{
        "id":id ,
        "numericId":numericId ,
        "code":code ,
        "info":currency ,
        "name":name ,
        "active":active ,
        "fee":fee ,
        "precision":precision ,
        "limits":MkMap(&VarMap{
            "amount":MkMap(&VarMap{
                "min":this.SafeNumber(currency , MkString("min") ),
                "max":this.SafeNumber(currency , MkString("max") ),
                }),
            "withdraw":MkMap(&VarMap{
                "min":MkUndefined() ,
                "max":this.SafeValue(withdrawalLimits , MkInteger(0) ),
                }),
            }),
        });
  }
  return result ;
}

func (this *Hollaex) FetchOrderBooks(goArgs ...*Variant) *Variant{
  symbols := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbols;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("publicGetOrderbooks"), params ); _ = response;
  result := MkMap(&VarMap{}); _ = result;
  marketIds := GoGetKeys(response ); _ = marketIds;
  for i := MkInteger(0) ; IsTrue(OpLw(i , marketIds.Length )); OpIncr(& i ){
    marketId := *(marketIds ).At(i ); _ = marketId;
    orderbook := *(response ).At(marketId ); _ = orderbook;
    symbol := this.SafeSymbol(marketId , MkUndefined() , MkString("-") ); _ = symbol;
    timestamp := this.Parse8601(this.SafeString(orderbook , MkString("timestamp") )); _ = timestamp;
    *(result ).At(symbol )= this.ParseOrderBook(*(response ).At(marketId ), timestamp );
  }
  return result ;
}

func (this *Hollaex) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  marketId := this.MarketId(symbol ); _ = marketId;
  request := MkMap(&VarMap{"symbol":marketId }); _ = request;
  response := this.Call(MkString("publicGetOrderbooks"), this.Extend(request , params )); _ = response;
  orderbook := this.SafeValue(response , marketId ); _ = orderbook;
  timestamp := this.Parse8601(this.SafeString(orderbook , MkString("timestamp") )); _ = timestamp;
  return this.ParseOrderBook(orderbook , symbol , timestamp );
}

func (this *Hollaex) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"symbol":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("publicGetTicker"), this.Extend(request , params )); _ = response;
  return this.ParseTicker(response , market );
}

func (this *Hollaex) FetchTickers(goArgs ...*Variant) *Variant{
  symbols := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbols;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("publicGetTickers"), this.Extend(params )); _ = response;
  return this.ParseTickers(response , symbols );
}

func (this *Hollaex) ParseTickers(goArgs ...*Variant) *Variant{
  response := GoGetArg(goArgs, 0, MkUndefined()); _ = response;
  symbols := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbols;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  result := MkMap(&VarMap{}); _ = result;
  keys := GoGetKeys(response ); _ = keys;
  for i := MkInteger(0) ; IsTrue(OpLw(i , keys.Length )); OpIncr(& i ){
    key := *(keys ).At(i ); _ = key;
    ticker := *(response ).At(key ); _ = ticker;
    marketId := this.SafeString(ticker , MkString("symbol") , key ); _ = marketId;
    market := this.SafeMarket(marketId , MkUndefined() , MkString("-") ); _ = market;
    symbol := *(market ).At(MkString("symbol") ); _ = symbol;
    *(result ).At(symbol )= this.Extend(this.ParseTicker(ticker , market ), params );
  }
  return this.FilterByArray(result , MkString("symbol") , symbols );
}

func (this *Hollaex) ParseTicker(goArgs ...*Variant) *Variant{
  ticker := GoGetArg(goArgs, 0, MkUndefined()); _ = ticker;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  marketId := this.SafeString(ticker , MkString("symbol") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market , MkString("-") ); _ = symbol;
  timestamp := this.Parse8601(this.SafeString2(ticker , MkString("time") , MkString("timestamp") )); _ = timestamp;
  close := this.SafeNumber(ticker , MkString("close") ); _ = close;
  result := MkMap(&VarMap{
        "symbol":symbol ,
        "info":ticker ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "high":this.SafeNumber(ticker , MkString("high") ),
        "low":this.SafeNumber(ticker , MkString("low") ),
        "bid":MkUndefined() ,
        "bidVolume":MkUndefined() ,
        "ask":MkUndefined() ,
        "askVolume":MkUndefined() ,
        "vwap":MkUndefined() ,
        "open":this.SafeNumber(ticker , MkString("open") ),
        "close":close ,
        "last":this.SafeNumber(ticker , MkString("last") , close ),
        "previousClose":MkUndefined() ,
        "change":MkUndefined() ,
        "percentage":MkUndefined() ,
        "average":MkUndefined() ,
        "baseVolume":this.SafeNumber(ticker , MkString("volume") ),
        "quoteVolume":MkUndefined() ,
        }); _ = result;
  return result ;
}

func (this *Hollaex) FetchTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"symbol":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("publicGetTrades"), this.Extend(request , params )); _ = response;
  trades := this.SafeValue(response , *(market ).At(MkString("id") ), MkArray(&VarArray{})); _ = trades;
  return this.ParseTrades(trades , market , since , limit );
}

func (this *Hollaex) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  marketId := this.SafeString(trade , MkString("symbol") ); _ = marketId;
  market = this.SafeMarket(marketId , market , MkString("-") );
  symbol := *(market ).At(MkString("symbol") ); _ = symbol;
  datetime := this.SafeString(trade , MkString("timestamp") ); _ = datetime;
  timestamp := this.Parse8601(datetime ); _ = timestamp;
  side := this.SafeString(trade , MkString("side") ); _ = side;
  priceString := this.SafeString(trade , MkString("price") ); _ = priceString;
  amountString := this.SafeString(trade , MkString("size") ); _ = amountString;
  price := this.ParseNumber(priceString ); _ = price;
  amount := this.ParseNumber(amountString ); _ = amount;
  cost := this.ParseNumber(Precise.StringMul(priceString , amountString )); _ = cost;
  feeCost := this.SafeNumber(trade , MkString("fee") ); _ = feeCost;
  fee := OpCopy(MkUndefined() ); _ = fee;
  if IsTrue(OpNotEq2(feeCost , MkUndefined() )) {
    quote := *(market ).At(MkString("quote") ); _ = quote;
    feeCurrencyCode := OpTrinary(OpNotEq2(market , MkUndefined() ), *(market ).At(MkString("quote") ), quote ); _ = feeCurrencyCode;
    fee = MkMap(&VarMap{
        "cost":feeCost ,
        "currency":feeCurrencyCode ,
        });
  }
  return MkMap(&VarMap{
        "info":trade ,
        "id":MkUndefined() ,
        "timestamp":timestamp ,
        "datetime":datetime ,
        "symbol":symbol ,
        "order":MkUndefined() ,
        "type":MkUndefined() ,
        "side":side ,
        "takerOrMaker":MkUndefined() ,
        "price":price ,
        "amount":amount ,
        "cost":cost ,
        "fee":fee ,
        });
}

func (this *Hollaex) FetchOHLCV(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  timeframe := GoGetArg(goArgs, 1, MkString("1h") ); _ = timeframe;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "symbol":*(market ).At(MkString("id") ),
        "resolution":*(*this.At(MkString("timeframes")) ).At(timeframe ),
        }); _ = request;
  duration := this.ParseTimeframe(timeframe ); _ = duration;
  if IsTrue(OpEq2(since , MkUndefined() )) {
    if IsTrue(OpEq2(limit , MkUndefined() )) {
      panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchOHLCV() requires a 'since' or a 'limit' argument") )));
    } else {
      end := this.Seconds(); _ = end;
      start := OpSub(end , OpMulti(duration , limit )); _ = start;
      *(request ).At(MkString("to") )= OpCopy(end );
      *(request ).At(MkString("from") )= OpCopy(start );
    }
  } else {
    if IsTrue(OpEq2(limit , MkUndefined() )) {
      *(request ).At(MkString("from") )= parseInt(OpDiv(since , MkInteger(1000) ));
      *(request ).At(MkString("to") )= this.Seconds();
    } else {
      start := parseInt(OpDiv(since , MkInteger(1000) )); _ = start;
      *(request ).At(MkString("from") )= OpCopy(start );
      *(request ).At(MkString("to") )= this.Sum(start , OpMulti(duration , limit ));
    }
  }
  response := this.Call(MkString("publicGetChart"), this.Extend(request , params )); _ = response;
  return this.ParseOHLCVs(response , market , timeframe , since , limit );
}

func (this *Hollaex) ParseOHLCV(goArgs ...*Variant) *Variant{
  response := GoGetArg(goArgs, 0, MkUndefined()); _ = response;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timeframe := GoGetArg(goArgs, 2, MkString("1h") ); _ = timeframe;
  since := GoGetArg(goArgs, 3, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 4, MkUndefined() ); _ = limit;
  return MkArray(&VarArray{
        this.Parse8601(this.SafeString(response , MkString("time") )),
        this.SafeNumber(response , MkString("open") ),
        this.SafeNumber(response , MkString("high") ),
        this.SafeNumber(response , MkString("low") ),
        this.SafeNumber(response , MkString("close") ),
        this.SafeNumber(response , MkString("volume") ),
        });
}

func (this *Hollaex) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("privateGetUserBalance"), params ); _ = response;
  timestamp := this.Parse8601(this.SafeString(response , MkString("updated_at") )); _ = timestamp;
  result := MkMap(&VarMap{
        "info":response ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        }); _ = result;
  currencyIds := GoGetKeys(*this.At(MkString("currencies_by_id")) ); _ = currencyIds;
  for i := MkInteger(0) ; IsTrue(OpLw(i , currencyIds.Length )); OpIncr(& i ){
    currencyId := *(currencyIds ).At(i ); _ = currencyId;
    code := this.SafeCurrencyCode(currencyId ); _ = code;
    account := this.Account(); _ = account;
    *(account ).At(MkString("free") )= this.SafeString(response , OpAdd(currencyId , MkString("_available") ));
    *(account ).At(MkString("total") )= this.SafeString(response , OpAdd(currencyId , MkString("_balance") ));
    *(result ).At(code )= OpCopy(account );
  }
  return this.ParseBalance(result );
}

func (this *Hollaex) FetchOpenOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"order_id":id }); _ = request;
  response := this.Call(MkString("privateGetOrdersOrderId"), this.Extend(request , params )); _ = response;
  return this.ParseOrder(response );
}

func (this *Hollaex) FetchOpenOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"open":MkBool(true) }); _ = request;
  return this.FetchOrders(symbol , since , limit , this.Extend(request , params ));
}

func (this *Hollaex) FetchClosedOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"status":MkString("filled") }); _ = request;
  return this.FetchOrders(symbol , since , limit , this.Extend(request , params ));
}

func (this *Hollaex) FetchOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"order_id":id }); _ = request;
  response := this.Call(MkString("privateGetOrders"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = data;
  order := this.SafeValue(data , MkInteger(0) ); _ = order;
  if IsTrue(OpEq2(order , MkUndefined() )) {
    panic(NewOrderNotFound(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" fetchOrder() could not find order id ") , id ))));
  }
  return this.ParseOrder(order );
}

func (this *Hollaex) FetchOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := OpCopy(MkUndefined() ); _ = market;
  request := MkMap(&VarMap{}); _ = request;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("symbol") )= *(market ).At(MkString("id") );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("start_date") )= this.Iso8601(since );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("privateGetOrders"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = data;
  return this.ParseOrders(data , market , since , limit );
}

func (this *Hollaex) ParseOrderStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "new":MkString("open") ,
        "pfilled":MkString("open") ,
        "filled":MkString("closed") ,
        "canceled":MkString("canceled") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Hollaex) ParseOrder(goArgs ...*Variant) *Variant{
  order := GoGetArg(goArgs, 0, MkUndefined()); _ = order;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  marketId := this.SafeString(order , MkString("symbol") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market , MkString("-") ); _ = symbol;
  id := this.SafeString(order , MkString("id") ); _ = id;
  timestamp := this.Parse8601(this.SafeString(order , MkString("created_at") )); _ = timestamp;
  type_ := this.SafeString(order , MkString("type") ); _ = type_;
  side := this.SafeString(order , MkString("side") ); _ = side;
  price := this.SafeNumber(order , MkString("price") ); _ = price;
  amount := this.SafeNumber(order , MkString("size") ); _ = amount;
  filled := this.SafeNumber(order , MkString("filled") ); _ = filled;
  status := this.ParseOrderStatus(this.SafeString(order , MkString("status") )); _ = status;
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
            "remaining":MkUndefined() ,
            "cost":MkUndefined() ,
            "trades":MkUndefined() ,
            "fee":MkUndefined() ,
            "info":order ,
            "average":MkUndefined() ,
            }));
}

func (this *Hollaex) CreateOrder(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 2, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 3, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 4, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 5, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "symbol":*(market ).At(MkString("id") ),
        "side":side ,
        "size":amount ,
        "type":type_ ,
        }); _ = request;
  if IsTrue(OpNotEq2(type_ , MkString("market") )) {
    *(request ).At(MkString("price") )= OpCopy(price );
  }
  stopPrice := this.SafeFloat2(params , MkString("stopPrice") , MkString("stop") ); _ = stopPrice;
  if IsTrue(OpNotEq2(stopPrice , MkUndefined() )) {
    *(request ).At(MkString("stop") )= parseFloat(this.PriceToPrecision(symbol , stopPrice ));
    params = this.Omit(params , MkArray(&VarArray{
            MkString("stopPrice") ,
            MkString("stop") ,
            }));
  }
  response := this.Call(MkString("privatePostOrder"), this.Extend(request , params )); _ = response;
  return this.ParseOrder(response , market );
}

func (this *Hollaex) CancelOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"order_id":id }); _ = request;
  response := this.Call(MkString("privateDeleteOrder"), this.Extend(request , params )); _ = response;
  return this.ParseOrder(response );
}

func (this *Hollaex) CancelAllOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("symbol") )= *(market ).At(MkString("id") );
  }
  response := this.Call(MkString("privateDeleteOrderAll"), this.Extend(request , params )); _ = response;
  return this.ParseOrders(response , market );
}

func (this *Hollaex) FetchMyTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("symbol") )= *(market ).At(MkString("id") );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("start_date") )= this.Iso8601(since );
  }
  response := this.Call(MkString("privateGetUserTrades"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = data;
  return this.ParseTrades(data , market , since , limit );
}

func (this *Hollaex) ParseDepositAddress(goArgs ...*Variant) *Variant{
  depositAddress := GoGetArg(goArgs, 0, MkUndefined()); _ = depositAddress;
  currency := GoGetArg(goArgs, 1, MkUndefined() ); _ = currency;
  address := this.SafeString(depositAddress , MkString("address") ); _ = address;
  tag := OpCopy(MkUndefined() ); _ = tag;
  if IsTrue(OpNotEq2(address , MkUndefined() )) {
    parts := address.Split(MkString(":") ); _ = parts;
    address = this.SafeString(parts , MkInteger(0) );
    tag = this.SafeString(parts , MkInteger(1) );
  }
  this.CheckAddress(address );
  currencyId := this.SafeString(depositAddress , MkString("currency") ); _ = currencyId;
  currency = this.SafeCurrency(currencyId , currency );
  network := this.SafeString(depositAddress , MkString("network") ); _ = network;
  return MkMap(&VarMap{
        "currency":*(currency ).At(MkString("code") ),
        "address":address ,
        "tag":tag ,
        "network":network ,
        "info":depositAddress ,
        });
}

func (this *Hollaex) FetchDepositAddresses(goArgs ...*Variant) *Variant{
  codes := GoGetArg(goArgs, 0, MkUndefined() ); _ = codes;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  network := this.SafeString(params , MkString("network") ); _ = network;
  params = this.Omit(params , MkString("network") );
  response := this.Call(MkString("privateGetUser"), params ); _ = response;
  wallet := this.SafeValue(response , MkString("wallet") , MkArray(&VarArray{})); _ = wallet;
  addresses := OpTrinary(OpEq2(network , MkUndefined() ), wallet , this.FilterBy(wallet , MkString("network") , network )); _ = addresses;
  return this.ParseDepositAddresses(addresses , codes );
}

func (this *Hollaex) FetchDeposits(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  currency := OpCopy(MkUndefined() ); _ = currency;
  if IsTrue(OpNotEq2(code , MkUndefined() )) {
    currency = this.Currency(code );
    *(request ).At(MkString("currency") )= *(currency ).At(MkString("id") );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("start_date") )= this.Iso8601(since );
  }
  response := this.Call(MkString("privateGetUserDeposits"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = data;
  return this.ParseTransactions(data , currency , since , limit );
}

func (this *Hollaex) FetchWithdrawals(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  currency := OpCopy(MkUndefined() ); _ = currency;
  if IsTrue(OpNotEq2(code , MkUndefined() )) {
    currency = this.Currency(code );
    *(request ).At(MkString("currency") )= *(currency ).At(MkString("id") );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("start_date") )= this.Iso8601(since );
  }
  response := this.Call(MkString("privateGetUserWithdrawals"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = data;
  return this.ParseTransactions(data , currency , since , limit );
}

func (this *Hollaex) ParseTransaction(goArgs ...*Variant) *Variant{
  transaction := GoGetArg(goArgs, 0, MkUndefined()); _ = transaction;
  currency := GoGetArg(goArgs, 1, MkUndefined() ); _ = currency;
  id := this.SafeString(transaction , MkString("id") ); _ = id;
  txid := this.SafeString(transaction , MkString("transaction_id") ); _ = txid;
  timestamp := this.Parse8601(this.SafeString(transaction , MkString("created_at") )); _ = timestamp;
  updated := this.Parse8601(this.SafeString(transaction , MkString("updated_at") )); _ = updated;
  type_ := this.SafeString(transaction , MkString("type") ); _ = type_;
  amount := this.SafeNumber(transaction , MkString("amount") ); _ = amount;
  address := this.SafeString(transaction , MkString("address") ); _ = address;
  addressTo := OpCopy(MkUndefined() ); _ = addressTo;
  addressFrom := OpCopy(MkUndefined() ); _ = addressFrom;
  tag := OpCopy(MkUndefined() ); _ = tag;
  tagTo := OpCopy(MkUndefined() ); _ = tagTo;
  tagFrom := OpCopy(MkUndefined() ); _ = tagFrom;
  if IsTrue(OpNotEq2(address , MkUndefined() )) {
    parts := address.Split(MkString(":") ); _ = parts;
    address = this.SafeString(parts , MkInteger(0) );
    tag = this.SafeString(parts , MkInteger(1) );
    addressTo = OpCopy(address );
    tagTo = OpCopy(tag );
  }
  currencyId := this.SafeString(transaction , MkString("currency") ); _ = currencyId;
  code := this.SafeCurrencyCode(currencyId ); _ = code;
  status := this.SafeValue(transaction , MkString("status") ); _ = status;
  dismissed := this.SafeValue(transaction , MkString("dismissed") ); _ = dismissed;
  rejected := this.SafeValue(transaction , MkString("rejected") ); _ = rejected;
  if IsTrue(status ) {
    status = MkString("ok") ;
  } else {
    if IsTrue(dismissed ) {
      status = MkString("canceled") ;
    } else {
      if IsTrue(rejected ) {
        status = MkString("failed") ;
      } else {
        status = MkString("pending") ;
      }
    }
  }
  fee := MkMap(&VarMap{
        "currency":code ,
        "cost":this.SafeNumber(transaction , MkString("fee") ),
        }); _ = fee;
  return MkMap(&VarMap{
        "info":transaction ,
        "id":id ,
        "txid":txid ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "addressFrom":addressFrom ,
        "address":address ,
        "addressTo":addressTo ,
        "tagFrom":tagFrom ,
        "tag":tag ,
        "tagTo":tagTo ,
        "type":type_ ,
        "amount":amount ,
        "currency":code ,
        "status":status ,
        "updated":updated ,
        "fee":fee ,
        });
}

func (this *Hollaex) Withdraw(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  amount := GoGetArg(goArgs, 1, MkUndefined()); _ = amount;
  address := GoGetArg(goArgs, 2, MkUndefined()); _ = address;
  tag := GoGetArg(goArgs, 3, MkUndefined() ); _ = tag;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.CheckAddress(address );
  this.LoadMarkets();
  currency := this.Currency(code ); _ = currency;
  if IsTrue(OpNotEq2(tag , MkUndefined() )) {
    OpAddAssign(& address , OpAdd(MkString(":") , tag ));
  }
  request := MkMap(&VarMap{
        "currency":*(currency ).At(MkString("id") ),
        "amount":amount ,
        "address":address ,
        }); _ = request;
  otp := this.SafeString(params , MkString("otp_code") ); _ = otp;
  if IsTrue(OpOr(OpNotEq2(otp , MkUndefined() ), OpNotEq2(*this.At(MkString("twofa")) , MkUndefined() ))) {
    if IsTrue(OpEq2(otp , MkUndefined() )) {
      otp = this.Oath();
    }
    *(request ).At(MkString("otp_code") )= OpCopy(otp );
  }
  response := this.Call(MkString("privatePostUserRequestWithdrawal"), this.Extend(request , params )); _ = response;
  return MkMap(&VarMap{
        "info":response ,
        "id":MkUndefined() ,
        });
}

func (this *Hollaex) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("public") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  query := this.Omit(params , this.ExtractParams(path )); _ = query;
  path = OpAdd(MkString("/") , OpAdd(*this.At(MkString("version")) , OpAdd(MkString("/") , this.ImplodeParams(path , params ))));
  if IsTrue(OpOr(OpEq2(method , MkString("GET") ), OpEq2(method , MkString("DELETE") ))) {
    if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
      OpAddAssign(& path , OpAdd(MkString("?") , this.Urlencode(query )));
    }
  }
  url := OpAdd(*(*this.At(MkString("urls")) ).At(MkString("api") ), path ); _ = url;
  if IsTrue(OpEq2(api , MkString("private") )) {
    this.CheckRequiredCredentials();
    defaultExpires := this.SafeInteger2(*this.At(MkString("options")) , MkString("api-expires") , MkString("expires") , parseInt(OpDiv(*this.At(MkString("timeout")) , MkInteger(1000) ))); _ = defaultExpires;
    expires := this.Sum(this.Seconds(), defaultExpires ); _ = expires;
    expiresString := expires.ToString(); _ = expiresString;
    auth := OpAdd(method , OpAdd(path , expiresString )); _ = auth;
    headers = MkMap(&VarMap{
        "api-key":this.Encode(*this.At(MkString("apiKey")) ),
        "api-expires":expiresString ,
        });
    if IsTrue(OpEq2(method , MkString("POST") )) {
      *(headers ).At(MkString("Content-type") )= MkString("application/json") ;
      if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
        body = this.Json(query );
        OpAddAssign(& auth , body );
      }
    }
    signature := this.Hmac(this.Encode(auth ), this.Encode(*this.At(MkString("secret")) )); _ = signature;
    *(headers ).At(MkString("api-signature") )= OpCopy(signature );
  }
  return MkMap(&VarMap{
        "url":url ,
        "method":method ,
        "body":body ,
        "headers":headers ,
        });
}

func (this *Hollaex) HandleErrors(goArgs ...*Variant) *Variant{
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
  if IsTrue(OpAnd(OpGtEq(code , MkInteger(400) ), OpLwEq(code , MkInteger(503) ))) {
    feedback := OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , body )); _ = feedback;
    message := this.SafeString(response , MkString("message") ); _ = message;
    this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("broad") ), message , feedback );
    status := code.ToString(); _ = status;
    this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("exact") ), status , feedback );
  }
  return MkUndefined()
}

