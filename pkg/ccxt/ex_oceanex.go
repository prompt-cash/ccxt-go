package ccxt

import (
)

type Oceanex struct {
	*ExchangeBase
}

var _ Exchange = (*Oceanex)(nil)

func init() {
	exchange := &Oceanex{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Oceanex) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("oceanex") ,
            "name":MkString("OceanEx") ,
            "countries":MkArray(&VarArray{
                MkString("LU") ,
                MkString("CN") ,
                MkString("SG") ,
                }),
            "version":MkString("v1") ,
            "rateLimit":MkInteger(3000) ,
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/1294454/58385970-794e2d80-8001-11e9-889c-0567cd79b78e.jpg") ,
                "api":MkString("https://api.oceanex.pro") ,
                "www":MkString("https://www.oceanex.pro.com") ,
                "doc":MkString("https://api.oceanex.pro/doc/v1") ,
                "referral":MkString("https://oceanex.pro/signup?referral=VE24QX") ,
                }),
            "has":MkMap(&VarMap{
                "fetchMarkets":MkBool(true) ,
                "fetchCurrencies":MkBool(false) ,
                "fetchTicker":MkBool(true) ,
                "fetchTickers":MkBool(true) ,
                "fetchOrderBook":MkBool(true) ,
                "fetchOrderBooks":MkBool(true) ,
                "fetchTrades":MkBool(true) ,
                "fetchTradingLimits":MkBool(false) ,
                "fetchTradingFees":MkBool(false) ,
                "fetchAllTradingFees":MkBool(true) ,
                "fetchFundingFees":MkBool(false) ,
                "fetchTime":MkBool(true) ,
                "fetchOrder":MkBool(true) ,
                "fetchOrders":MkBool(true) ,
                "fetchOpenOrders":MkBool(true) ,
                "fetchClosedOrders":MkBool(true) ,
                "fetchBalance":MkBool(true) ,
                "createMarketOrder":MkBool(true) ,
                "createOrder":MkBool(true) ,
                "cancelOrder":MkBool(true) ,
                "cancelOrders":MkBool(true) ,
                "cancelAllOrders":MkBool(true) ,
                }),
            "timeframes":MkMap(&VarMap{
                "1m":MkString("1m") ,
                "5m":MkString("5m") ,
                "15m":MkString("15m") ,
                "30m":MkString("30m") ,
                "1h":MkString("1h") ,
                "4h":MkString("4h") ,
                "12h":MkString("12h") ,
                "1d":MkString("1d") ,
                "1w":MkString("1w") ,
                }),
            "api":MkMap(&VarMap{
                "public":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("markets") ,
                        MkString("tickers/{pair}") ,
                        MkString("tickers_multi") ,
                        MkString("order_book") ,
                        MkString("order_book/multi") ,
                        MkString("fees/trading") ,
                        MkString("trades") ,
                        MkString("timestamp") ,
                        })}),
                "private":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("key") ,
                        MkString("members/me") ,
                        MkString("orders") ,
                        MkString("orders/filter") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("orders") ,
                        MkString("orders/multi") ,
                        MkString("order/delete") ,
                        MkString("order/delete/multi") ,
                        MkString("orders/clear") ,
                        }),
                    }),
                }),
            "fees":MkMap(&VarMap{"trading":MkMap(&VarMap{
                    "tierBased":MkBool(false) ,
                    "percentage":MkBool(true) ,
                    "maker":OpDiv(MkNumber(0.1) , MkInteger(100) ),
                    "taker":OpDiv(MkNumber(0.1) , MkInteger(100) ),
                    })}),
            "commonCurrencies":MkMap(&VarMap{"PLA":MkString("Plair") }),
            "exceptions":MkMap(&VarMap{
                "codes":MkMap(&VarMap{
                    "-1":BadRequest ,
                    "-2":BadRequest ,
                    "1001":BadRequest ,
                    "1004":ArgumentsRequired ,
                    "1006":AuthenticationError ,
                    "1008":AuthenticationError ,
                    "1010":AuthenticationError ,
                    "1011":PermissionDenied ,
                    "2001":AuthenticationError ,
                    "2002":InvalidOrder ,
                    "2004":OrderNotFound ,
                    "9003":PermissionDenied ,
                    }),
                "exact":MkMap(&VarMap{
                    "market does not have a valid value":BadRequest ,
                    "side does not have a valid value":BadRequest ,
                    "Account::AccountError: Cannot lock funds":InsufficientFunds ,
                    "The account does not exist":AuthenticationError ,
                    }),
                }),
            }));
}

func (this *Oceanex) FetchMarkets(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"show_details":MkBool(true) }); _ = request;
  response := this.Call(MkString("publicGetMarkets"), this.Extend(request , params )); _ = response;
  result := MkArray(&VarArray{}); _ = result;
  markets := this.SafeValue(response , MkString("data") ); _ = markets;
  for i := MkInteger(0) ; IsTrue(OpLw(i , markets.Length )); OpIncr(& i ){
    market := *(markets ).At(i ); _ = market;
    id := this.SafeValue(market , MkString("id") ); _ = id;
    name := this.SafeValue(market , MkString("name") ); _ = name;
    baseId := MkUndefined(); _ = baseId;
    quoteId := MkUndefined(); _ = quoteId;
    ArrayUnpack(name.Split(MkString("/") ), &baseId, &quoteId);
    base := this.SafeCurrencyCode(baseId ); _ = base;
    quote := this.SafeCurrencyCode(quoteId ); _ = quote;
    baseId = baseId.ToLowerCase();
    quoteId = quoteId.ToLowerCase();
    symbol := OpAdd(base , OpAdd(MkString("/") , quote )); _ = symbol;
    result.Push(MkMap(&VarMap{
            "id":id ,
            "symbol":symbol ,
            "base":base ,
            "quote":quote ,
            "baseId":baseId ,
            "quoteId":quoteId ,
            "active":MkBool(true) ,
            "info":market ,
            "precision":MkMap(&VarMap{
                "amount":this.SafeInteger(market , MkString("amount_precision") ),
                "price":this.SafeInteger(market , MkString("price_precision") ),
                "base":this.SafeInteger(market , MkString("ask_precision") ),
                "quote":this.SafeInteger(market , MkString("bid_precision") ),
                }),
            "limits":MkMap(&VarMap{
                "amount":MkMap(&VarMap{
                    "min":MkUndefined() ,
                    "max":MkUndefined() ,
                    }),
                "price":MkMap(&VarMap{
                    "min":MkUndefined() ,
                    "max":MkUndefined() ,
                    }),
                "cost":MkMap(&VarMap{
                    "min":this.SafeNumber(market , MkString("minimum_trading_amount") ),
                    "max":MkUndefined() ,
                    }),
                }),
            }));
  }
  return result ;
}

func (this *Oceanex) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"pair":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("publicGetTickersPair"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  return this.ParseTicker(data , market );
}

func (this *Oceanex) FetchTickers(goArgs ...*Variant) *Variant{
  symbols := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbols;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  if IsTrue(OpEq2(symbols , MkUndefined() )) {
    symbols = OpCopy(*this.At(MkString("symbols")) );
  }
  marketIds := this.MarketIds(symbols ); _ = marketIds;
  request := MkMap(&VarMap{"markets":marketIds }); _ = request;
  response := this.Call(MkString("publicGetTickersMulti"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") ); _ = data;
  result := MkMap(&VarMap{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , data.Length )); OpIncr(& i ){
    ticker := *(data ).At(i ); _ = ticker;
    marketId := this.SafeString(ticker , MkString("market") ); _ = marketId;
    market := this.SafeMarket(marketId ); _ = market;
    symbol := *(market ).At(MkString("symbol") ); _ = symbol;
    *(result ).At(symbol )= this.ParseTicker(ticker , market );
  }
  return this.FilterByArray(result , MkString("symbol") , symbols );
}

func (this *Oceanex) ParseTicker(goArgs ...*Variant) *Variant{
  data := GoGetArg(goArgs, 0, MkUndefined()); _ = data;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  ticker := this.SafeValue(data , MkString("ticker") , MkMap(&VarMap{})); _ = ticker;
  timestamp := this.SafeTimestamp(data , MkString("at") ); _ = timestamp;
  return MkMap(&VarMap{
        "symbol":*(market ).At(MkString("symbol") ),
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
        "close":this.SafeNumber(ticker , MkString("last") ),
        "last":this.SafeNumber(ticker , MkString("last") ),
        "previousClose":MkUndefined() ,
        "change":MkUndefined() ,
        "percentage":MkUndefined() ,
        "average":MkUndefined() ,
        "baseVolume":this.SafeNumber(ticker , MkString("volume") ),
        "quoteVolume":MkUndefined() ,
        "info":ticker ,
        });
}

func (this *Oceanex) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"market":*(market ).At(MkString("id") )}); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("publicGetOrderBook"), this.Extend(request , params )); _ = response;
  orderbook := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = orderbook;
  timestamp := this.SafeTimestamp(orderbook , MkString("timestamp") ); _ = timestamp;
  return this.ParseOrderBook(orderbook , symbol , timestamp );
}

func (this *Oceanex) FetchOrderBooks(goArgs ...*Variant) *Variant{
  symbols := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbols;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  if IsTrue(OpEq2(symbols , MkUndefined() )) {
    symbols = OpCopy(*this.At(MkString("symbols")) );
  }
  marketIds := this.MarketIds(symbols ); _ = marketIds;
  request := MkMap(&VarMap{"markets":marketIds }); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("publicGetOrderBookMulti"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = data;
  result := MkMap(&VarMap{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , data.Length )); OpIncr(& i ){
    orderbook := *(data ).At(i ); _ = orderbook;
    marketId := this.SafeString(orderbook , MkString("market") ); _ = marketId;
    symbol := this.SafeSymbol(marketId ); _ = symbol;
    timestamp := this.SafeTimestamp(orderbook , MkString("timestamp") ); _ = timestamp;
    *(result ).At(symbol )= this.ParseOrderBook(orderbook , symbol , timestamp );
  }
  return result ;
}

func (this *Oceanex) FetchTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"market":*(market ).At(MkString("id") )}); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("publicGetTrades"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") ); _ = data;
  return this.ParseTrades(data , market , since , limit );
}

func (this *Oceanex) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  side := this.SafeValue(trade , MkString("side") ); _ = side;
  if IsTrue(OpEq2(side , MkString("bid") )) {
    side = MkString("buy") ;
  } else {
    if IsTrue(OpEq2(side , MkString("ask") )) {
      side = MkString("sell") ;
    }
  }
  marketId := this.SafeValue(trade , MkString("market") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market ); _ = symbol;
  timestamp := this.SafeTimestamp(trade , MkString("created_on") ); _ = timestamp;
  if IsTrue(OpEq2(timestamp , MkUndefined() )) {
    timestamp = this.Parse8601(this.SafeString(trade , MkString("created_at") ));
  }
  return MkMap(&VarMap{
        "info":trade ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "symbol":symbol ,
        "id":this.SafeString(trade , MkString("id") ),
        "order":MkUndefined() ,
        "type":MkString("limit") ,
        "takerOrMaker":MkUndefined() ,
        "side":side ,
        "price":this.SafeNumber(trade , MkString("price") ),
        "amount":this.SafeNumber(trade , MkString("volume") ),
        "cost":MkUndefined() ,
        "fee":MkUndefined() ,
        });
}

func (this *Oceanex) FetchTime(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetTimestamp"), params ); _ = response;
  return this.SafeTimestamp(response , MkString("data") );
}

func (this *Oceanex) FetchAllTradingFees(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetFeesTrading"), params ); _ = response;
  data := this.SafeValue(response , MkString("data") ); _ = data;
  result := MkMap(&VarMap{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , data.Length )); OpIncr(& i ){
    group := *(data ).At(i ); _ = group;
    maker := this.SafeValue(group , MkString("ask_fee") , MkMap(&VarMap{})); _ = maker;
    taker := this.SafeValue(group , MkString("bid_fee") , MkMap(&VarMap{})); _ = taker;
    marketId := this.SafeString(group , MkString("market") ); _ = marketId;
    symbol := this.SafeSymbol(marketId ); _ = symbol;
    *(result ).At(symbol )= MkMap(&VarMap{
        "info":group ,
        "symbol":symbol ,
        "maker":this.SafeNumber(maker , MkString("value") ),
        "taker":this.SafeNumber(taker , MkString("value") ),
        });
  }
  return result ;
}

func (this *Oceanex) FetchKey(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("privateGetKey"), params ); _ = response;
  return this.SafeValue(response , MkString("data") );
}

func (this *Oceanex) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("privateGetMembersMe"), params ); _ = response;
  data := this.SafeValue(response , MkString("data") ); _ = data;
  balances := this.SafeValue(data , MkString("accounts") ); _ = balances;
  result := MkMap(&VarMap{"info":response }); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , balances.Length )); OpIncr(& i ){
    balance := *(balances ).At(i ); _ = balance;
    currencyId := this.SafeValue(balance , MkString("currency") ); _ = currencyId;
    code := this.SafeCurrencyCode(currencyId ); _ = code;
    account := this.Account(); _ = account;
    *(account ).At(MkString("free") )= this.SafeString(balance , MkString("balance") );
    *(account ).At(MkString("used") )= this.SafeString(balance , MkString("locked") );
    *(result ).At(code )= OpCopy(account );
  }
  return this.ParseBalance(result );
}

func (this *Oceanex) CreateOrder(goArgs ...*Variant) *Variant{
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
        "side":side ,
        "ord_type":type_ ,
        "volume":this.AmountToPrecision(symbol , amount ),
        }); _ = request;
  if IsTrue(OpEq2(type_ , MkString("limit") )) {
    *(request ).At(MkString("price") )= this.PriceToPrecision(symbol , price );
  }
  response := this.Call(MkString("privatePostOrders"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") ); _ = data;
  return this.ParseOrder(data , market );
}

func (this *Oceanex) FetchOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  ids := OpCopy(id ); _ = ids;
  if IsTrue(OpNot(GoIsArray(id ))) {
    ids = MkArray(&VarArray{
        id ,
        });
  }
  this.LoadMarkets();
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
  }
  request := MkMap(&VarMap{"ids":ids }); _ = request;
  response := this.Call(MkString("privateGetOrders"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") ); _ = data;
  dataLength := OpCopy(data.Length ); _ = dataLength;
  if IsTrue(OpEq2(data , MkUndefined() )) {
    panic(NewOrderNotFound(OpAdd(*this.At(MkString("id")) , MkString(" could not found matching order") )));
  }
  if IsTrue(GoIsArray(id )) {
    return this.ParseOrders(data , market );
  }
  if IsTrue(OpEq2(dataLength , MkInteger(0) )) {
    panic(NewOrderNotFound(OpAdd(*this.At(MkString("id")) , MkString(" could not found matching order") )));
  }
  return this.ParseOrder(*(data ).At(MkInteger(0) ), market );
}

func (this *Oceanex) FetchOpenOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"states":MkArray(&VarArray{
            MkString("wait") ,
            })}); _ = request;
  return this.FetchOrders(symbol , since , limit , this.Extend(request , params ));
}

func (this *Oceanex) FetchClosedOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"states":MkArray(&VarArray{
            MkString("done") ,
            MkString("cancel") ,
            })}); _ = request;
  return this.FetchOrders(symbol , since , limit , this.Extend(request , params ));
}

func (this *Oceanex) FetchOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchOrders() requires a `symbol` argument") )));
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  states := this.SafeValue(params , MkString("states") , MkArray(&VarArray{
            MkString("wait") ,
            MkString("done") ,
            MkString("cancel") ,
            })); _ = states;
  query := this.Omit(params , MkString("states") ); _ = query;
  request := MkMap(&VarMap{
        "market":*(market ).At(MkString("id") ),
        "states":states ,
        "need_price":MkString("True") ,
        }); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("privateGetOrdersFilter"), this.Extend(request , query )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = data;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , data.Length )); OpIncr(& i ){
    orders := this.SafeValue(*(data ).At(i ), MkString("orders") , MkArray(&VarArray{})); _ = orders;
    status := this.ParseOrderStatus(this.SafeValue(*(data ).At(i ), MkString("state") )); _ = status;
    parsedOrders := this.ParseOrders(orders , market , since , limit , MkMap(&VarMap{"status":status })); _ = parsedOrders;
    result = this.ArrayConcat(result , parsedOrders );
  }
  return result ;
}

func (this *Oceanex) ParseOrder(goArgs ...*Variant) *Variant{
  order := GoGetArg(goArgs, 0, MkUndefined()); _ = order;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  status := this.ParseOrderStatus(this.SafeValue(order , MkString("state") )); _ = status;
  marketId := this.SafeString2(order , MkString("market") , MkString("market_id") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market ); _ = symbol;
  timestamp := this.SafeTimestamp(order , MkString("created_on") ); _ = timestamp;
  if IsTrue(OpEq2(timestamp , MkUndefined() )) {
    timestamp = this.Parse8601(this.SafeString(order , MkString("created_at") ));
  }
  return this.SafeOrder(MkMap(&VarMap{
            "info":order ,
            "id":this.SafeString(order , MkString("id") ),
            "clientOrderId":MkUndefined() ,
            "timestamp":timestamp ,
            "datetime":this.Iso8601(timestamp ),
            "lastTradeTimestamp":MkUndefined() ,
            "symbol":symbol ,
            "type":this.SafeValue(order , MkString("ord_type") ),
            "timeInForce":MkUndefined() ,
            "postOnly":MkUndefined() ,
            "side":this.SafeValue(order , MkString("side") ),
            "price":this.SafeNumber(order , MkString("price") ),
            "stopPrice":MkUndefined() ,
            "average":this.SafeNumber(order , MkString("avg_price") ),
            "amount":this.SafeNumber(order , MkString("volume") ),
            "remaining":this.SafeNumber(order , MkString("remaining_volume") ),
            "filled":this.SafeNumber(order , MkString("executed_volume") ),
            "status":status ,
            "cost":MkUndefined() ,
            "trades":MkUndefined() ,
            "fee":MkUndefined() ,
            }));
}

func (this *Oceanex) ParseOrderStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "wait":MkString("open") ,
        "done":MkString("closed") ,
        "cancel":MkString("canceled") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Oceanex) CreateOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  orders := GoGetArg(goArgs, 1, MkUndefined()); _ = orders;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "market":*(market ).At(MkString("id") ),
        "orders":orders ,
        }); _ = request;
  response := this.Call(MkString("privatePostOrdersMulti"), this.Extend(request , params )); _ = response;
  data := *(response ).At(MkString("data") ); _ = data;
  return this.ParseOrders(data );
}

func (this *Oceanex) CancelOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("privatePostOrderDelete"), this.Extend(MkMap(&VarMap{"id":id }), params )); _ = response;
  data := this.SafeValue(response , MkString("data") ); _ = data;
  return this.ParseOrder(data );
}

func (this *Oceanex) CancelOrders(goArgs ...*Variant) *Variant{
  ids := GoGetArg(goArgs, 0, MkUndefined()); _ = ids;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("privatePostOrderDeleteMulti"), this.Extend(MkMap(&VarMap{"ids":ids }), params )); _ = response;
  data := this.SafeValue(response , MkString("data") ); _ = data;
  return this.ParseOrders(data );
}

func (this *Oceanex) CancelAllOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("privatePostOrdersClear"), params ); _ = response;
  data := this.SafeValue(response , MkString("data") ); _ = data;
  return this.ParseOrders(data );
}

func (this *Oceanex) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("public") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  url := OpAdd(*(*this.At(MkString("urls")) ).At(MkString("api") ), OpAdd(MkString("/") , OpAdd(*this.At(MkString("version")) , OpAdd(MkString("/") , this.ImplodeParams(path , params ))))); _ = url;
  query := this.Omit(params , this.ExtractParams(path )); _ = query;
  if IsTrue(OpEq2(api , MkString("public") )) {
    if IsTrue(OpOr(OpEq2(path , MkString("tickers_multi") ), OpEq2(path , MkString("order_book/multi") ))) {
      request := MkString("?") ; _ = request;
      markets := this.SafeValue(params , MkString("markets") ); _ = markets;
      for i := MkInteger(0) ; IsTrue(OpLw(i , markets.Length )); OpIncr(& i ){
        OpAddAssign(& request , OpAdd(MkString("markets[]=") , OpAdd(*(markets ).At(i ), MkString("&") )));
      }
      limit := this.SafeValue(params , MkString("limit") ); _ = limit;
      if IsTrue(OpNotEq2(limit , MkUndefined() )) {
        OpAddAssign(& request , OpAdd(MkString("limit=") , limit ));
      }
      OpAddAssign(& url , request );
    } else {
      if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
        OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(query )));
      }
    }
  } else {
    if IsTrue(OpEq2(api , MkString("private") )) {
      this.CheckRequiredCredentials();
      request := MkMap(&VarMap{
            "uid":*this.At(MkString("apiKey")) ,
            "data":query ,
            }); _ = request;
      jwt_token := this.Jwt(request , this.Encode(*this.At(MkString("secret")) ), MkString("RS256") ); _ = jwt_token;
      OpAddAssign(& url , OpAdd(MkString("?user_jwt=") , jwt_token ));
    }
  }
  headers = MkMap(&VarMap{"Content-Type":MkString("application/json") });
  return MkMap(&VarMap{
        "url":url ,
        "method":method ,
        "body":body ,
        "headers":headers ,
        });
}

func (this *Oceanex) HandleErrors(goArgs ...*Variant) *Variant{
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
  errorCode := this.SafeString(response , MkString("code") ); _ = errorCode;
  message := this.SafeString(response , MkString("message") ); _ = message;
  if IsTrue(OpAnd(OpNotEq2(errorCode , MkUndefined() ), OpNotEq2(errorCode , MkString("0") ))) {
    feedback := OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , body )); _ = feedback;
    this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("codes") ), errorCode , feedback );
    this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("exact") ), message , feedback );
    panic(NewExchangeError(feedback ));
  }
  return MkUndefined()
}

