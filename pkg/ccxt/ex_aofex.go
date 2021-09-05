package ccxt

import (
)

type Aofex struct {
	*ExchangeBase
}

var _ Exchange = (*Aofex)(nil)

func init() {
	exchange := &Aofex{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Aofex) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("aofex") ,
            "name":MkString("AOFEX") ,
            "countries":MkArray(&VarArray{
                MkString("GB") ,
                }),
            "rateLimit":MkInteger(1000) ,
            "hostname":MkString("openapi.aofex.com") ,
            "has":MkMap(&VarMap{
                "fetchMarkets":MkBool(true) ,
                "fetchCurrencies":MkBool(false) ,
                "fetchOrderBook":MkBool(true) ,
                "fetchTrades":MkBool(true) ,
                "fetchTicker":MkBool(true) ,
                "fetchTickers":MkBool(true) ,
                "fetchOHLCV":MkBool(true) ,
                "fetchBalance":MkBool(true) ,
                "createOrder":MkBool(true) ,
                "cancelOrder":MkBool(true) ,
                "cancelAllOrders":MkBool(true) ,
                "fetchOpenOrders":MkBool(true) ,
                "fetchClosedOrders":MkBool(true) ,
                "fetchClosedOrder":MkBool(true) ,
                "fetchOrderTrades":MkBool(true) ,
                "fetchTradingFee":MkBool(true) ,
                }),
            "timeframes":MkMap(&VarMap{
                "1m":MkString("1min") ,
                "5m":MkString("5min") ,
                "15m":MkString("15min") ,
                "30m":MkString("30min") ,
                "1h":MkString("1hour") ,
                "6h":MkString("6hour") ,
                "12h":MkString("12hour") ,
                "1d":MkString("1day") ,
                "1w":MkString("1week") ,
                }),
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/51840849/77670271-056d1080-6f97-11ea-9ac2-4268e9ed0c1f.jpg") ,
                "api":MkMap(&VarMap{
                    "public":MkString("https://{hostname}/openApi") ,
                    "private":MkString("https://{hostname}/openApi") ,
                    }),
                "www":MkString("https://aofex.com") ,
                "doc":MkString("https://aofex.zendesk.com/hc/en-us/sections/360005576574-API") ,
                "fees":MkString("https://aofex.zendesk.com/hc/en-us/articles/360025814934-Fees-on-AOFEX") ,
                "referral":MkString("https://aofex.com/#/register?key=9763840") ,
                }),
            "api":MkMap(&VarMap{
                "public":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("market/symbols") ,
                        MkString("market/trade") ,
                        MkString("market/depth") ,
                        MkString("market/kline") ,
                        MkString("market/precision") ,
                        MkString("market/24kline") ,
                        MkString("market/gears_depth") ,
                        MkString("market/detail") ,
                        })}),
                "private":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("entrust/currentList") ,
                        MkString("entrust/historyList") ,
                        MkString("entrust/rate") ,
                        MkString("wallet/list") ,
                        MkString("entrust/detail") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("entrust/add") ,
                        MkString("entrust/cancel") ,
                        }),
                    }),
                }),
            "fees":MkMap(&VarMap{"trading":MkMap(&VarMap{
                    "maker":MkNumber(0.0019) ,
                    "taker":MkNumber(0.002) ,
                    })}),
            "exceptions":MkMap(&VarMap{
                "exact":MkMap(&VarMap{
                    "20001":ExchangeError ,
                    "20401":PermissionDenied ,
                    "20500":ExchangeError ,
                    "20501":BadSymbol ,
                    "20502":ExchangeError ,
                    "20503":ExchangeError ,
                    "20504":InsufficientFunds ,
                    "20505":BadRequest ,
                    "20506":AuthenticationError ,
                    "20507":ExchangeError ,
                    "20508":InvalidAddress ,
                    "20509":InsufficientFunds ,
                    "20510":InvalidOrder ,
                    "20511":InvalidOrder ,
                    "20512":InvalidOrder ,
                    "20513":InvalidOrder ,
                    "20514":InvalidOrder ,
                    "20515":InvalidOrder ,
                    "20516":InvalidOrder ,
                    "20517":InvalidOrder ,
                    "50518":InvalidOrder ,
                    "20519":InvalidNonce ,
                    "20520":InvalidNonce ,
                    "20521":BadRequest ,
                    "20522":ExchangeError ,
                    "20523":AuthenticationError ,
                    "20524":PermissionDenied ,
                    "20525":AuthenticationError ,
                    "20526":PermissionDenied ,
                    "20527":PermissionDenied ,
                    "20528":PermissionDenied ,
                    "20529":AuthenticationError ,
                    "20530":PermissionDenied ,
                    }),
                "broad":MkMap(&VarMap{}),
                }),
            "options":MkMap(&VarMap{"fetchBalance":MkMap(&VarMap{"show_all":MkString("0") })}),
            "commonCurrencies":MkMap(&VarMap{"CPC":MkString("Consensus Planet Coin") }),
            }));
}

func (this *Aofex) FetchMarkets(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  markets := this.Call(MkString("publicGetMarketSymbols"), params ); _ = markets;
  precisions := this.Call(MkString("publicGetMarketPrecision")); _ = precisions;
  precisions = this.SafeValue(precisions , MkString("result") , MkMap(&VarMap{}));
  markets = this.SafeValue(markets , MkString("result") , MkArray(&VarArray{}));
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , markets.Length )); OpIncr(& i ){
    market := *(markets ).At(i ); _ = market;
    id := this.SafeString(market , MkString("symbol") ); _ = id;
    baseId := this.SafeString(market , MkString("base_currency") ); _ = baseId;
    quoteId := this.SafeString(market , MkString("quote_currency") ); _ = quoteId;
    base := this.SafeCurrencyCode(baseId ); _ = base;
    quote := this.SafeCurrencyCode(quoteId ); _ = quote;
    symbol := OpAdd(base , OpAdd(MkString("/") , quote )); _ = symbol;
    numericId := this.SafeInteger(market , MkString("id") ); _ = numericId;
    precision := this.SafeValue(precisions , id , MkMap(&VarMap{})); _ = precision;
    makerFeeString := this.SafeString(market , MkString("maker_fee") ); _ = makerFeeString;
    takerFeeString := this.SafeString(market , MkString("taker_fee") ); _ = takerFeeString;
    makerFee := this.ParseNumber(Precise.StringDiv(makerFeeString , MkString("1000") )); _ = makerFee;
    takerFee := this.ParseNumber(Precise.StringDiv(takerFeeString , MkString("1000") )); _ = takerFee;
    result.Push(MkMap(&VarMap{
            "id":id ,
            "numericId":numericId ,
            "symbol":symbol ,
            "baseId":baseId ,
            "quoteId":quoteId ,
            "base":base ,
            "quote":quote ,
            "active":MkUndefined() ,
            "maker":makerFee ,
            "taker":takerFee ,
            "precision":MkMap(&VarMap{
                "amount":this.SafeInteger(precision , MkString("amount") ),
                "price":this.SafeInteger(precision , MkString("price") ),
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
            "info":market ,
            }));
  }
  return result ;
}

func (this *Aofex) ParseOHLCV(goArgs ...*Variant) *Variant{
  ohlcv := GoGetArg(goArgs, 0, MkUndefined()); _ = ohlcv;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  return MkArray(&VarArray{
        this.SafeTimestamp(ohlcv , MkString("id") ),
        this.SafeNumber(ohlcv , MkString("open") ),
        this.SafeNumber(ohlcv , MkString("high") ),
        this.SafeNumber(ohlcv , MkString("low") ),
        this.SafeNumber(ohlcv , MkString("close") ),
        this.SafeNumber(ohlcv , MkString("amount") ),
        });
}

func (this *Aofex) FetchOHLCV(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  timeframe := GoGetArg(goArgs, 1, MkString("1m") ); _ = timeframe;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  if IsTrue(OpEq2(limit , MkUndefined() )) {
    limit = MkInteger(150) ;
  }
  request := MkMap(&VarMap{
        "symbol":*(market ).At(MkString("id") ),
        "period":*(*this.At(MkString("timeframes")) ).At(timeframe ),
        "size":limit ,
        }); _ = request;
  response := this.Call(MkString("publicGetMarketKline"), this.Extend(request , params )); _ = response;
  result := this.SafeValue(response , MkString("result") , MkMap(&VarMap{})); _ = result;
  data := this.SafeValue(result , MkString("data") , MkArray(&VarArray{})); _ = data;
  return this.ParseOHLCVs(data , market , since , limit );
}

func (this *Aofex) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  options := this.SafeValue(*this.At(MkString("options")) , MkString("fetchBalance") , MkMap(&VarMap{})); _ = options;
  showAll := this.SafeValue(options , MkString("show_all") , MkString("0") ); _ = showAll;
  request := MkMap(&VarMap{"show_all":showAll }); _ = request;
  response := this.Call(MkString("privateGetWalletList"), this.Extend(request , params )); _ = response;
  result := MkMap(&VarMap{
        "info":response ,
        "timestamp":MkUndefined() ,
        "datetime":MkUndefined() ,
        }); _ = result;
  balances := this.SafeValue(response , MkString("result") , MkArray(&VarArray{})); _ = balances;
  for i := MkInteger(0) ; IsTrue(OpLw(i , balances.Length )); OpIncr(& i ){
    balance := *(balances ).At(i ); _ = balance;
    currencyId := this.SafeString(balance , MkString("currency") ); _ = currencyId;
    code := this.SafeCurrencyCode(currencyId ); _ = code;
    account := this.Account(); _ = account;
    *(account ).At(MkString("free") )= this.SafeString(balance , MkString("available") );
    *(account ).At(MkString("used") )= this.SafeString(balance , MkString("frozen") );
    *(result ).At(code )= OpCopy(account );
  }
  return this.ParseBalance(result );
}

func (this *Aofex) FetchTradingFee(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"symbol":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("privateGetEntrustRate"), this.Extend(request , params )); _ = response;
  result := this.SafeValue(response , MkString("result") , MkMap(&VarMap{})); _ = result;
  return MkMap(&VarMap{
        "info":response ,
        "symbol":symbol ,
        "maker":this.SafeNumber(result , MkString("fromFee") ),
        "taker":this.SafeNumber(result , MkString("toFee") ),
        });
}

func (this *Aofex) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"symbol":this.MarketId(symbol )}); _ = request;
  response := this.Call(MkString("publicGetMarketDepth"), this.Extend(request , params )); _ = response;
  result := this.SafeValue(response , MkString("result") , MkMap(&VarMap{})); _ = result;
  timestamp := this.SafeInteger(result , MkString("ts") ); _ = timestamp;
  return this.ParseOrderBook(result , symbol , timestamp );
}

func (this *Aofex) ParseTicker(goArgs ...*Variant) *Variant{
  ticker := GoGetArg(goArgs, 0, MkUndefined()); _ = ticker;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.SafeTimestamp(ticker , MkString("id") ); _ = timestamp;
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  if IsTrue(market ) {
    symbol = *(market ).At(MkString("symbol") );
  }
  open := this.SafeNumber(ticker , MkString("open") ); _ = open;
  last := this.SafeNumber(ticker , MkString("close") ); _ = last;
  change := OpCopy(MkUndefined() ); _ = change;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    change = this.ParseNumber(this.PriceToPrecision(symbol , OpSub(last , open )));
  } else {
    change = OpSub(last , open );
  }
  average := OpDiv(this.Sum(last , open ), MkInteger(2) ); _ = average;
  percentage := OpDiv(change , OpMulti(open , MkInteger(100) )); _ = percentage;
  baseVolume := this.SafeNumber(ticker , MkString("amount") ); _ = baseVolume;
  quoteVolume := this.SafeNumber(ticker , MkString("vol") ); _ = quoteVolume;
  vwap := this.Vwap(baseVolume , quoteVolume ); _ = vwap;
  if IsTrue(OpNotEq2(vwap , MkUndefined() )) {
    vwap = this.ParseNumber(this.PriceToPrecision(symbol , vwap ));
  }
  return MkMap(&VarMap{
        "symbol":symbol ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "high":this.SafeNumber(ticker , MkString("high") ),
        "low":this.SafeNumber(ticker , MkString("low") ),
        "bid":MkUndefined() ,
        "bidVolume":MkUndefined() ,
        "ask":MkUndefined() ,
        "askVolume":MkUndefined() ,
        "vwap":vwap ,
        "open":open ,
        "close":last ,
        "last":last ,
        "previousClose":MkUndefined() ,
        "change":change ,
        "percentage":percentage ,
        "average":average ,
        "baseVolume":baseVolume ,
        "quoteVolume":quoteVolume ,
        "info":ticker ,
        });
}

func (this *Aofex) FetchTickers(goArgs ...*Variant) *Variant{
  symbols := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbols;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  if IsTrue(OpNotEq2(symbols , MkUndefined() )) {
    ids := this.MarketIds(symbols ); _ = ids;
    *(request ).At(MkString("symbol") )= ids.Join(MkString(",") );
  }
  response := this.Call(MkString("publicGetMarket24kline"), this.Extend(request , params )); _ = response;
  tickers := this.SafeValue(response , MkString("result") , MkArray(&VarArray{})); _ = tickers;
  result := MkMap(&VarMap{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , tickers.Length )); OpIncr(& i ){
    marketId := this.SafeString(*(tickers ).At(i ), MkString("symbol") ); _ = marketId;
    market := this.SafeMarket(marketId , MkUndefined() , MkString("-") ); _ = market;
    symbol := *(market ).At(MkString("symbol") ); _ = symbol;
    data := this.SafeValue(*(tickers ).At(i ), MkString("data") , MkMap(&VarMap{})); _ = data;
    *(result ).At(symbol )= this.ParseTicker(data , market );
  }
  return this.FilterByArray(result , MkString("symbol") , symbols );
}

func (this *Aofex) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"symbol":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("publicGetMarketDetail"), this.Extend(request , params )); _ = response;
  result := this.SafeValue(response , MkString("result") , MkMap(&VarMap{})); _ = result;
  return this.ParseTicker(result , market );
}

func (this *Aofex) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  id := this.SafeString(trade , MkString("id") ); _ = id;
  ctime := this.Parse8601(this.SafeString(trade , MkString("ctime") )); _ = ctime;
  timestamp := OpSub(this.SafeTimestamp(trade , MkString("ts") , ctime ), MkInteger(28800000) ); _ = timestamp;
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  if IsTrue(OpAnd(OpEq2(symbol , MkUndefined() ), OpNotEq2(market , MkUndefined() ))) {
    symbol = *(market ).At(MkString("symbol") );
  }
  side := this.SafeString(trade , MkString("direction") ); _ = side;
  priceString := this.SafeString(trade , MkString("price") ); _ = priceString;
  amountString := this.SafeString2(trade , MkString("amount") , MkString("number") ); _ = amountString;
  price := this.ParseNumber(priceString ); _ = price;
  amount := this.ParseNumber(amountString ); _ = amount;
  cost := this.SafeNumber(trade , MkString("total_price") ); _ = cost;
  if IsTrue(OpEq2(cost , MkUndefined() )) {
    cost = this.ParseNumber(Precise.StringMul(priceString , amountString ));
  }
  feeCost := this.SafeNumber(trade , MkString("fee") ); _ = feeCost;
  fee := OpCopy(MkUndefined() ); _ = fee;
  if IsTrue(OpNotEq2(feeCost , MkUndefined() )) {
    feeCurrencyCode := OpCopy(MkUndefined() ); _ = feeCurrencyCode;
    if IsTrue(OpNotEq2(market , MkUndefined() )) {
      if IsTrue(OpEq2(side , MkString("buy") )) {
        feeCurrencyCode = *(market ).At(MkString("base") );
      } else {
        if IsTrue(OpEq2(side , MkString("sell") )) {
          feeCurrencyCode = *(market ).At(MkString("quote") );
        }
      }
    }
    fee = MkMap(&VarMap{
        "cost":feeCost ,
        "currency":feeCurrencyCode ,
        });
  }
  return MkMap(&VarMap{
        "id":id ,
        "info":trade ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
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

func (this *Aofex) FetchTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"symbol":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("publicGetMarketTrade"), this.Extend(request , params )); _ = response;
  result := this.SafeValue(response , MkString("result") , MkMap(&VarMap{})); _ = result;
  data := this.SafeValue(result , MkString("data") , MkArray(&VarArray{})); _ = data;
  return this.ParseTrades(data , market , since , limit );
}

func (this *Aofex) ParseOrderStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "1":MkString("open") ,
        "2":MkString("open") ,
        "3":MkString("closed") ,
        "4":MkString("canceled") ,
        "5":MkString("canceled") ,
        "6":MkString("canceled") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Aofex) ParseOrder(goArgs ...*Variant) *Variant{
  order := GoGetArg(goArgs, 0, MkUndefined()); _ = order;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  id := this.SafeString(order , MkString("order_sn") ); _ = id;
  orderStatus := this.SafeString(order , MkString("status") ); _ = orderStatus;
  status := this.ParseOrderStatus(orderStatus ); _ = status;
  marketId := this.SafeString(order , MkString("symbol") ); _ = marketId;
  market = this.SafeMarket(marketId , market , MkString("-") );
  timestamp := this.Parse8601(this.SafeString(order , MkString("ctime") )); _ = timestamp;
  if IsTrue(OpNotEq2(timestamp , MkUndefined() )) {
    OpSubAssign(& timestamp , MkInteger(28800000) );
  }
  orderType := this.SafeString(order , MkString("type") ); _ = orderType;
  type_ := OpTrinary(OpEq2(orderType , MkString("2") ), MkString("limit") , MkString("market") ); _ = type_;
  side := this.SafeString(order , MkString("side") ); _ = side;
  cost := OpCopy(MkUndefined() ); _ = cost;
  price := OpCopy(MkUndefined() ); _ = price;
  amount := OpCopy(MkUndefined() ); _ = amount;
  average := OpCopy(MkUndefined() ); _ = average;
  number := this.SafeNumber(order , MkString("number") ); _ = number;
  totalPrice := this.SafeNumber(order , MkString("total_price") ); _ = totalPrice;
  if IsTrue(OpEq2(type_ , MkString("limit") )) {
    amount = OpCopy(number );
    price = this.SafeNumber(order , MkString("price") );
  } else {
    average = this.SafeNumber(order , MkString("deal_price") );
    if IsTrue(OpEq2(side , MkString("buy") )) {
      amount = this.SafeNumber(order , MkString("deal_number") );
    } else {
      amount = OpCopy(number );
    }
  }
  rawTrades := this.SafeValue(order , MkString("trades") , MkArray(&VarArray{})); _ = rawTrades;
  for i := MkInteger(0) ; IsTrue(OpLw(i , rawTrades.Length )); OpIncr(& i ){
    *(*(rawTrades ).At(i )).At(MkString("direction") )= OpCopy(side );
  }
  trades := this.ParseTrades(rawTrades , market , MkUndefined() , MkUndefined() , MkMap(&VarMap{
            "symbol":*(market ).At(MkString("symbol") ),
            "order":id ,
            "type":type_ ,
            })); _ = trades;
  if IsTrue(OpEq2(type_ , MkString("limit") )) {
    cost = OpCopy(totalPrice );
  } else {
    if IsTrue(OpEq2(side , MkString("buy") )) {
      cost = OpCopy(number );
    }
  }
  filled := OpCopy(MkUndefined() ); _ = filled;
  if IsTrue(OpAnd(OpEq2(type_ , MkString("limit") ), OpEq2(orderStatus , MkString("3") ))) {
    filled = OpCopy(amount );
  }
  return this.SafeOrder(MkMap(&VarMap{
            "info":order ,
            "id":id ,
            "clientOrderId":MkUndefined() ,
            "timestamp":timestamp ,
            "datetime":this.Iso8601(timestamp ),
            "lastTradeTimestamp":MkUndefined() ,
            "status":status ,
            "symbol":*(market ).At(MkString("symbol") ),
            "type":type_ ,
            "timeInForce":MkUndefined() ,
            "postOnly":MkUndefined() ,
            "side":side ,
            "price":price ,
            "stopPrice":MkUndefined() ,
            "cost":cost ,
            "average":average ,
            "amount":amount ,
            "filled":filled ,
            "remaining":MkUndefined() ,
            "trades":trades ,
            "fee":MkUndefined() ,
            }));
}

func (this *Aofex) FetchClosedOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"order_sn":id }); _ = request;
  response := this.Call(MkString("privateGetEntrustDetail"), this.Extend(request , params )); _ = response;
  result := this.SafeValue(response , MkString("result") , MkMap(&VarMap{})); _ = result;
  trades := this.SafeValue(result , MkString("trades") , MkArray(&VarArray{})); _ = trades;
  order := this.SafeValue(result , MkString("entrust") , MkMap(&VarMap{})); _ = order;
  *(order ).At(MkString("trades") )= OpCopy(trades );
  return this.ParseOrder(order );
}

func (this *Aofex) FetchOrderTrades(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  response := this.FetchClosedOrder(id , symbol , params ); _ = response;
  return this.SafeValue(response , MkString("trades") , MkArray(&VarArray{}));
}

func (this *Aofex) FetchOrdersWithMethod(goArgs ...*Variant) *Variant{
  method := GoGetArg(goArgs, 0, MkUndefined()); _ = method;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"direct":MkString("prev") }); _ = request;
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("symbol") )= *(market ).At(MkString("id") );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(method , this.Extend(request , params )); _ = response;
  result := this.SafeValue(response , MkString("result") , MkArray(&VarArray{})); _ = result;
  return this.ParseOrders(result , market , since , limit );
}

func (this *Aofex) FetchOpenOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  return this.FetchOrdersWithMethod(MkString("privateGetEntrustCurrentList") , symbol , since , limit , params );
}

func (this *Aofex) FetchClosedOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  return this.FetchOrdersWithMethod(MkString("privateGetEntrustHistoryList") , symbol , since , limit , params );
}

func (this *Aofex) CreateOrder(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 2, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 3, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 4, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 5, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  orderType := OpAdd(side , OpAdd(MkString("-") , type_ )); _ = orderType;
  request := MkMap(&VarMap{
        "symbol":*(market ).At(MkString("id") ),
        "type":orderType ,
        }); _ = request;
  if IsTrue(OpEq2(type_ , MkString("limit") )) {
    *(request ).At(MkString("amount") )= this.AmountToPrecision(symbol , amount );
    *(request ).At(MkString("price") )= this.PriceToPrecision(symbol , price );
  } else {
    if IsTrue(OpEq2(type_ , MkString("market") )) {
      if IsTrue(OpEq2(side , MkString("buy") )) {
        createMarketBuyOrderRequiresPrice := this.SafeValue(*this.At(MkString("options")) , MkString("createMarketBuyOrderRequiresPrice") , MkBool(true) ); _ = createMarketBuyOrderRequiresPrice;
        cost := OpCopy(amount ); _ = cost;
        if IsTrue(createMarketBuyOrderRequiresPrice ) {
          if IsTrue(OpNotEq2(price , MkUndefined() )) {
            cost = OpMulti(amount , price );
          } else {
            panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")) , MkString(" createOrder() requires the price argument with market buy orders to calculate total order cost (amount to spend), where cost = amount * price. Supply a price argument to createOrder() call if you want the cost to be calculated for you from price and amount, or, alternatively, add .options['createMarketBuyOrderRequiresPrice'] = false and supply the total cost value in the 'amount' argument") )));
          }
        }
        precision := *(*(market ).At(MkString("precision") )).At(MkString("price") ); _ = precision;
        *(request ).At(MkString("amount") )= this.DecimalToPrecision(cost , TRUNCATE , precision , *this.At(MkString("precisionMode")) );
      } else {
        *(request ).At(MkString("amount") )= this.AmountToPrecision(symbol , amount );
      }
    }
  }
  response := this.Call(MkString("privatePostEntrustAdd"), this.Extend(request , params )); _ = response;
  result := this.SafeValue(response , MkString("result") , MkMap(&VarMap{})); _ = result;
  order := this.ParseOrder(result , market ); _ = order;
  timestamp := this.Milliseconds(); _ = timestamp;
  return this.Extend(order , MkMap(&VarMap{
            "timestamp":timestamp ,
            "datetime":this.Iso8601(timestamp ),
            "amount":amount ,
            "price":price ,
            "type":type_ ,
            "side":side ,
            }));
}

func (this *Aofex) CancelOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"order_ids":id }); _ = request;
  response := this.Call(MkString("privatePostEntrustCancel"), this.Extend(request , params )); _ = response;
  result := this.SafeValue(response , MkString("result") , MkMap(&VarMap{})); _ = result;
  success := this.SafeValue(result , MkString("success") , MkArray(&VarArray{})); _ = success;
  if IsTrue(OpNot(this.InArray(id , success ))) {
    panic(NewOrderNotFound(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" order id ") , OpAdd(id , OpAdd(MkString(" not found in successfully canceled orders: ") , this.Json(response )))))));
  }
  timestamp := OpCopy(MkUndefined() ); _ = timestamp;
  return MkMap(&VarMap{
        "info":response ,
        "id":id ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "lastTradeTimestamp":MkUndefined() ,
        "status":MkString("canceled") ,
        "symbol":symbol ,
        "type":MkUndefined() ,
        "side":MkUndefined() ,
        "price":MkUndefined() ,
        "cost":MkUndefined() ,
        "average":MkUndefined() ,
        "amount":MkUndefined() ,
        "filled":MkUndefined() ,
        "remaining":MkUndefined() ,
        "trades":MkUndefined() ,
        "fee":MkUndefined() ,
        "clientOrderId":MkUndefined() ,
        });
}

func (this *Aofex) CancelAllOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" cancelAllOrders() requires a symbol argument") )));
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"symbol":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("privatePostEntrustCancel"), this.Extend(request , params )); _ = response;
  return response ;
}

func (this *Aofex) Nonce(goArgs ...*Variant) *Variant{
  return this.Milliseconds();
}

func (this *Aofex) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("public") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  url := OpAdd(this.ImplodeHostname(*(*(*this.At(MkString("urls")) ).At(MkString("api") )).At(api )), OpAdd(MkString("/") , path )); _ = url;
  keys := GoGetKeys(params ); _ = keys;
  keysLength := OpCopy(keys.Length ); _ = keysLength;
  if IsTrue(OpEq2(api , MkString("public") )) {
    if IsTrue(OpGt(keysLength , MkInteger(0) )) {
      OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(params )));
    }
  } else {
    nonce := (this.Nonce()).Call(MkString("toString") ); _ = nonce;
    uuid := this.Uuid(); _ = uuid;
    randomString := uuid.Slice(MkInteger(0) , MkInteger(5) ); _ = randomString;
    nonceString := OpAdd(nonce , OpAdd(MkString("_") , randomString )); _ = nonceString;
    auth := MkMap(&VarMap{}); _ = auth;
    *(auth ).At(*this.At(MkString("apiKey")) )= OpCopy(*this.At(MkString("apiKey")) );
    *(auth ).At(*this.At(MkString("secret")) )= OpCopy(*this.At(MkString("secret")) );
    *(auth ).At(nonceString )= OpCopy(nonceString );
    for i := MkInteger(0) ; IsTrue(OpLw(i , keysLength )); OpIncr(& i ){
      key := *(keys ).At(i ); _ = key;
      *(auth ).At(key )= OpAdd(key , OpAdd(MkString("=") , *(params ).At(key )));
    }
    keysorted := this.Keysort(auth ); _ = keysorted;
    stringToSign := MkString("") ; _ = stringToSign;
    keys = GoGetKeys(keysorted );
    for i := MkInteger(0) ; IsTrue(OpLw(i , keys.Length )); OpIncr(& i ){
      key := *(keys ).At(i ); _ = key;
      OpAddAssign(& stringToSign , *(keysorted ).At(key ));
    }
    signature := this.Hash(this.Encode(stringToSign ), MkString("sha1") ); _ = signature;
    headers = MkMap(&VarMap{
        "Nonce":nonceString ,
        "Token":*this.At(MkString("apiKey")) ,
        "Signature":signature ,
        });
    if IsTrue(OpEq2(method , MkString("POST") )) {
      *(headers ).At(MkString("Content-Type") )= MkString("application/x-www-form-urlencoded") ;
      if IsTrue(OpGt(keysLength , MkInteger(0) )) {
        body = this.Urlencode(params );
      }
    } else {
      if IsTrue(OpGt(keysLength , MkInteger(0) )) {
        OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(params )));
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

func (this *Aofex) HandleErrors(goArgs ...*Variant) *Variant{
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
  error := this.SafeString(response , MkString("errno") ); _ = error;
  if IsTrue(OpAnd(OpNotEq2(error , MkUndefined() ), OpNotEq2(error , MkString("0") ))) {
    message := this.SafeString(response , MkString("errmsg") ); _ = message;
    feedback := OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , body )); _ = feedback;
    this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("exact") ), error , feedback );
    this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("broad") ), message , feedback );
    panic(NewExchangeError(feedback ));
  }
  return MkUndefined()
}

