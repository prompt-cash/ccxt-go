package ccxt

import (
)

type Whitebit struct {
	*ExchangeBase
}

var _ Exchange = (*Whitebit)(nil)

func init() {
	exchange := &Whitebit{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Whitebit) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("whitebit") ,
            "name":MkString("WhiteBit") ,
            "version":MkString("v2") ,
            "countries":MkArray(&VarArray{
                MkString("EE") ,
                }),
            "rateLimit":MkInteger(500) ,
            "has":MkMap(&VarMap{
                "cancelOrder":MkBool(false) ,
                "CORS":MkBool(false) ,
                "createDepositAddress":MkBool(false) ,
                "createLimitOrder":MkBool(false) ,
                "createMarketOrder":MkBool(false) ,
                "createOrder":MkBool(false) ,
                "deposit":MkBool(false) ,
                "editOrder":MkBool(false) ,
                "fetchBalance":MkBool(false) ,
                "fetchBidsAsks":MkBool(false) ,
                "fetchCurrencies":MkBool(true) ,
                "fetchMarkets":MkBool(true) ,
                "fetchOHLCV":MkBool(true) ,
                "fetchOrderBook":MkBool(true) ,
                "fetchStatus":MkBool(true) ,
                "fetchTicker":MkBool(true) ,
                "fetchTickers":MkBool(true) ,
                "fetchTrades":MkBool(true) ,
                "fetchTradingFees":MkBool(true) ,
                "privateAPI":MkBool(false) ,
                "publicAPI":MkBool(true) ,
                }),
            "timeframes":MkMap(&VarMap{
                "1m":MkString("1m") ,
                "3m":MkString("3m") ,
                "5m":MkString("5m") ,
                "15m":MkString("15m") ,
                "30m":MkString("30m") ,
                "1h":MkString("1h") ,
                "2h":MkString("2h") ,
                "4h":MkString("4h") ,
                "6h":MkString("6h") ,
                "8h":MkString("8h") ,
                "12h":MkString("12h") ,
                "1d":MkString("1d") ,
                "3d":MkString("3d") ,
                "1w":MkString("1w") ,
                "1M":MkString("1M") ,
                }),
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/1294454/66732963-8eb7dd00-ee66-11e9-849b-10d9282bb9e0.jpg") ,
                "api":MkMap(&VarMap{
                    "web":MkString("https://whitebit.com/") ,
                    "publicV2":MkString("https://whitebit.com/api/v2/public") ,
                    "publicV1":MkString("https://whitebit.com/api/v1/public") ,
                    }),
                "www":MkString("https://www.whitebit.com") ,
                "doc":MkString("https://documenter.getpostman.com/view/7473075/Szzj8dgv?version=latest") ,
                "fees":MkString("https://whitebit.com/fee-schedule") ,
                "referral":MkString("https://whitebit.com/referral/d9bdf40e-28f2-4b52-b2f9-cd1415d82963") ,
                }),
            "api":MkMap(&VarMap{
                "web":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("v1/healthcheck") ,
                        })}),
                "publicV1":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("markets") ,
                        MkString("tickers") ,
                        MkString("ticker") ,
                        MkString("symbols") ,
                        MkString("depth/result") ,
                        MkString("history") ,
                        MkString("kline") ,
                        })}),
                "publicV2":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("markets") ,
                        MkString("ticker") ,
                        MkString("assets") ,
                        MkString("fee") ,
                        MkString("depth/{market}") ,
                        MkString("trades/{market}") ,
                        })}),
                }),
            "fees":MkMap(&VarMap{"trading":MkMap(&VarMap{
                    "tierBased":MkBool(false) ,
                    "percentage":MkBool(true) ,
                    "taker":this.ParseNumber(MkString("0.001") ),
                    "maker":this.ParseNumber(MkString("0.001") ),
                    })}),
            "options":MkMap(&VarMap{"fetchTradesMethod":MkString("fetchTradesV1") }),
            "exceptions":MkMap(&VarMap{
                "exact":MkMap(&VarMap{"503":ExchangeNotAvailable }),
                "broad":MkMap(&VarMap{"Market is not available":BadSymbol }),
                }),
            }));
}

func (this *Whitebit) FetchMarkets(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicV2GetMarkets"), params ); _ = response;
  markets := this.SafeValue(response , MkString("result") ); _ = markets;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , markets.Length )); OpIncr(& i ){
    market := *(markets ).At(i ); _ = market;
    id := this.SafeString(market , MkString("name") ); _ = id;
    baseId := this.SafeString(market , MkString("stock") ); _ = baseId;
    quoteId := this.SafeString(market , MkString("money") ); _ = quoteId;
    base := this.SafeCurrencyCode(baseId ); _ = base;
    quote := this.SafeCurrencyCode(quoteId ); _ = quote;
    symbol := OpAdd(base , OpAdd(MkString("/") , quote )); _ = symbol;
    active := this.SafeValue(market , MkString("tradesEnabled") ); _ = active;
    entry := MkMap(&VarMap{
          "id":id ,
          "symbol":symbol ,
          "base":base ,
          "quote":quote ,
          "baseId":baseId ,
          "quoteId":quoteId ,
          "info":market ,
          "active":active ,
          "precision":MkMap(&VarMap{
              "amount":this.SafeInteger(market , MkString("stockPrec") ),
              "price":this.SafeInteger(market , MkString("moneyPrec") ),
              }),
          "limits":MkMap(&VarMap{
              "amount":MkMap(&VarMap{
                  "min":this.SafeNumber(market , MkString("minAmount") ),
                  "max":MkUndefined() ,
                  }),
              "price":MkMap(&VarMap{
                  "min":MkUndefined() ,
                  "max":MkUndefined() ,
                  }),
              "cost":MkMap(&VarMap{
                  "min":this.SafeNumber(market , MkString("minTotal") ),
                  "max":MkUndefined() ,
                  }),
              }),
          }); _ = entry;
    result.Push(entry );
  }
  return result ;
}

func (this *Whitebit) FetchCurrencies(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicV2GetAssets"), params ); _ = response;
  currencies := this.SafeValue(response , MkString("result") ); _ = currencies;
  ids := GoGetKeys(currencies ); _ = ids;
  result := MkMap(&VarMap{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , ids.Length )); OpIncr(& i ){
    id := *(ids ).At(i ); _ = id;
    currency := *(currencies ).At(id ); _ = currency;
    canDeposit := this.SafeValue(currency , MkString("canDeposit") , MkBool(true) ); _ = canDeposit;
    canWithdraw := this.SafeValue(currency , MkString("canWithdraw") , MkBool(true) ); _ = canWithdraw;
    active := OpAnd(canDeposit , canWithdraw ); _ = active;
    code := this.SafeCurrencyCode(id ); _ = code;
    *(result ).At(code )= MkMap(&VarMap{
        "id":id ,
        "code":code ,
        "info":currency ,
        "name":MkUndefined() ,
        "active":active ,
        "fee":MkUndefined() ,
        "precision":MkUndefined() ,
        "limits":MkMap(&VarMap{
            "amount":MkMap(&VarMap{
                "min":MkUndefined() ,
                "max":MkUndefined() ,
                }),
            "withdraw":MkMap(&VarMap{
                "min":this.SafeNumber(currency , MkString("minWithdrawal") ),
                "max":this.SafeNumber(currency , MkString("maxWithdrawal") ),
                }),
            }),
        });
  }
  return result ;
}

func (this *Whitebit) FetchTradingFees(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicV2GetFee"), params ); _ = response;
  fees := this.SafeValue(response , MkString("result") ); _ = fees;
  return MkMap(&VarMap{
        "maker":this.SafeNumber(fees , MkString("makerFee") ),
        "taker":this.SafeNumber(fees , MkString("takerFee") ),
        });
}

func (this *Whitebit) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"market":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("publicV1GetTicker"), this.Extend(request , params )); _ = response;
  ticker := this.SafeValue(response , MkString("result") , MkMap(&VarMap{})); _ = ticker;
  return this.ParseTicker(ticker , market );
}

func (this *Whitebit) ParseTicker(goArgs ...*Variant) *Variant{
  ticker := GoGetArg(goArgs, 0, MkUndefined()); _ = ticker;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.SafeTimestamp(ticker , MkString("at") , this.Milliseconds()); _ = timestamp;
  ticker = this.SafeValue(ticker , MkString("ticker") , ticker );
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  if IsTrue(OpNotEq2(market , MkUndefined() )) {
    symbol = *(market ).At(MkString("symbol") );
  }
  last := this.SafeNumber(ticker , MkString("last") ); _ = last;
  percentage := this.SafeNumber(ticker , MkString("change") ); _ = percentage;
  change := OpCopy(MkUndefined() ); _ = change;
  if IsTrue(OpNotEq2(percentage , MkUndefined() )) {
    change = this.NumberToString(OpMulti(percentage , MkNumber(0.01) ));
  }
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
        "open":this.SafeNumber(ticker , MkString("open") ),
        "close":last ,
        "last":last ,
        "previousClose":MkUndefined() ,
        "change":change ,
        "percentage":percentage ,
        "average":MkUndefined() ,
        "baseVolume":this.SafeNumber(ticker , MkString("volume") ),
        "quoteVolume":this.SafeNumber(ticker , MkString("deal") ),
        "info":ticker ,
        });
}

func (this *Whitebit) FetchTickers(goArgs ...*Variant) *Variant{
  symbols := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbols;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("publicV1GetTickers"), params ); _ = response;
  data := this.SafeValue(response , MkString("result") ); _ = data;
  marketIds := GoGetKeys(data ); _ = marketIds;
  result := MkMap(&VarMap{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , marketIds.Length )); OpIncr(& i ){
    marketId := *(marketIds ).At(i ); _ = marketId;
    market := this.SafeMarket(marketId ); _ = market;
    ticker := this.ParseTicker(*(data ).At(marketId ), market ); _ = ticker;
    symbol := *(ticker ).At(MkString("symbol") ); _ = symbol;
    *(result ).At(symbol )= OpCopy(ticker );
  }
  return this.FilterByArray(result , MkString("symbol") , symbols );
}

func (this *Whitebit) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"market":*(market ).At(MkString("id") )}); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("publicV2GetDepthMarket"), this.Extend(request , params )); _ = response;
  result := this.SafeValue(response , MkString("result") , MkMap(&VarMap{})); _ = result;
  timestamp := this.Parse8601(this.SafeString(result , MkString("lastUpdateTimestamp") )); _ = timestamp;
  return this.ParseOrderBook(result , symbol , timestamp );
}

func (this *Whitebit) FetchTradesV1(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "market":*(market ).At(MkString("id") ),
        "lastId":MkInteger(1) ,
        }); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("publicV1GetHistory"), this.Extend(request , params )); _ = response;
  result := this.SafeValue(response , MkString("result") , MkArray(&VarArray{})); _ = result;
  return this.ParseTrades(result , market , since , limit );
}

func (this *Whitebit) FetchTradesV2(goArgs ...*Variant) *Variant{
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
  response := this.Call(MkString("publicV2GetTradesMarket"), this.Extend(request , params )); _ = response;
  result := this.SafeValue(response , MkString("result") , MkArray(&VarArray{})); _ = result;
  return this.ParseTrades(result , market , since , limit );
}

func (this *Whitebit) FetchTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  method := this.SafeString(*this.At(MkString("options")) , MkString("fetchTradesMethod") , MkString("fetchTradesV2") ); _ = method;
  return this.Call(method , symbol , since , limit , params );
}

func (this *Whitebit) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.SafeValue(trade , MkString("time") ); _ = timestamp;
  if IsTrue(OpEq2(OpType(timestamp ), MkString("string") )) {
    timestamp = this.Parse8601(timestamp );
  } else {
    timestamp = parseInt(OpMulti(timestamp , MkInteger(1000) ));
  }
  priceString := this.SafeString(trade , MkString("price") ); _ = priceString;
  amountString := this.SafeString2(trade , MkString("amount") , MkString("volume") ); _ = amountString;
  cost := this.ParseNumber(Precise.StringMul(priceString , amountString )); _ = cost;
  price := this.ParseNumber(priceString ); _ = price;
  amount := this.ParseNumber(amountString ); _ = amount;
  id := this.SafeString2(trade , MkString("id") , MkString("tradeId") ); _ = id;
  side := this.SafeString(trade , MkString("type") ); _ = side;
  if IsTrue(OpEq2(side , MkUndefined() )) {
    isBuyerMaker := this.SafeValue(trade , MkString("isBuyerMaker") ); _ = isBuyerMaker;
    side = OpTrinary(isBuyerMaker , MkString("buy") , MkString("sell") );
  }
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  if IsTrue(OpNotEq2(market , MkUndefined() )) {
    symbol = *(market ).At(MkString("symbol") );
  }
  return MkMap(&VarMap{
        "info":trade ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "symbol":symbol ,
        "id":id ,
        "order":MkUndefined() ,
        "type":MkUndefined() ,
        "takerOrMaker":MkUndefined() ,
        "side":side ,
        "price":price ,
        "amount":amount ,
        "cost":cost ,
        "fee":MkUndefined() ,
        });
}

func (this *Whitebit) FetchOHLCV(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  timeframe := GoGetArg(goArgs, 1, MkString("1m") ); _ = timeframe;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "market":*(market ).At(MkString("id") ),
        "interval":*(*this.At(MkString("timeframes")) ).At(timeframe ),
        }); _ = request;
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    maxLimit := MkInteger(1440) ; _ = maxLimit;
    if IsTrue(OpEq2(limit , MkUndefined() )) {
      limit = OpCopy(maxLimit );
    }
    limit = Math.Min(limit , maxLimit );
    start := parseInt(OpDiv(since , MkInteger(1000) )); _ = start;
    duration := this.ParseTimeframe(timeframe ); _ = duration;
    end := this.Sum(start , OpMulti(duration , limit )); _ = end;
    *(request ).At(MkString("start") )= OpCopy(start );
    *(request ).At(MkString("end") )= OpCopy(end );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("publicV1GetKline"), this.Extend(request , params )); _ = response;
  result := this.SafeValue(response , MkString("result") , MkArray(&VarArray{})); _ = result;
  return this.ParseOHLCVs(result , market , timeframe , since , limit );
}

func (this *Whitebit) ParseOHLCV(goArgs ...*Variant) *Variant{
  ohlcv := GoGetArg(goArgs, 0, MkUndefined()); _ = ohlcv;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  return MkArray(&VarArray{
        this.SafeTimestamp(ohlcv , MkInteger(0) ),
        this.SafeNumber(ohlcv , MkInteger(1) ),
        this.SafeNumber(ohlcv , MkInteger(3) ),
        this.SafeNumber(ohlcv , MkInteger(4) ),
        this.SafeNumber(ohlcv , MkInteger(2) ),
        this.SafeNumber(ohlcv , MkInteger(5) ),
        });
}

func (this *Whitebit) FetchStatus(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("webGetV1Healthcheck"), params ); _ = response;
  status := this.SafeInteger(response , MkString("status") ); _ = status;
  formattedStatus := MkString("ok") ; _ = formattedStatus;
  if IsTrue(OpEq2(status , MkInteger(503) )) {
    formattedStatus = MkString("maintenance") ;
  }
  *this.At(MkString("status")) = this.Extend(*this.At(MkString("status")) , MkMap(&VarMap{
          "status":formattedStatus ,
          "updated":this.Milliseconds(),
          }));
  return *this.At(MkString("status")) ;
}

func (this *Whitebit) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("publicV1") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  query := this.Omit(params , this.ExtractParams(path )); _ = query;
  url := OpAdd(*(*(*this.At(MkString("urls")) ).At(MkString("api") )).At(api ), OpAdd(MkString("/") , this.ImplodeParams(path , params ))); _ = url;
  if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
    OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(query )));
  }
  return MkMap(&VarMap{
        "url":url ,
        "method":method ,
        "body":body ,
        "headers":headers ,
        });
}

func (this *Whitebit) HandleErrors(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  reason := GoGetArg(goArgs, 1, MkUndefined()); _ = reason;
  url := GoGetArg(goArgs, 2, MkUndefined()); _ = url;
  method := GoGetArg(goArgs, 3, MkUndefined()); _ = method;
  headers := GoGetArg(goArgs, 4, MkUndefined()); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined()); _ = body;
  response := GoGetArg(goArgs, 6, MkUndefined()); _ = response;
  requestHeaders := GoGetArg(goArgs, 7, MkUndefined()); _ = requestHeaders;
  requestBody := GoGetArg(goArgs, 8, MkUndefined()); _ = requestBody;
  if IsTrue(OpOr(OpEq2(code , MkInteger(418) ), OpEq2(code , MkInteger(429) ))) {
    panic(NewDDoSProtection(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , OpAdd(code.ToString(), OpAdd(MkString(" ") , OpAdd(reason , OpAdd(MkString(" ") , body ))))))));
  }
  if IsTrue(OpEq2(code , MkInteger(404) )) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , OpAdd(code.ToString(), MkString(" endpoint not found") )))));
  }
  if IsTrue(OpNotEq2(response , MkUndefined() )) {
    success := this.SafeValue(response , MkString("success") ); _ = success;
    if IsTrue(OpNot(success )) {
      feedback := OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , body )); _ = feedback;
      status := this.SafeString(response , MkString("status") ); _ = status;
      if IsTrue(OpEq2(OpType(status ), MkString("string") )) {
        this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("exact") ), status , feedback );
      }
      this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("broad") ), body , feedback );
      panic(NewExchangeError(feedback ));
    }
  }
  return MkUndefined()
}

