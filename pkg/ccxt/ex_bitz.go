package ccxt

import (
)

type Bitz struct {
	*ExchangeBase
}

var _ Exchange = (*Bitz)(nil)

func init() {
	exchange := &Bitz{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Bitz) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("bitz") ,
            "name":MkString("Bit-Z") ,
            "countries":MkArray(&VarArray{
                MkString("HK") ,
                }),
            "rateLimit":MkInteger(2000) ,
            "version":MkString("v2") ,
            "userAgent":*(*this.At(MkString("userAgents")) ).At(MkString("chrome") ),
            "has":MkMap(&VarMap{
                "cancelOrder":MkBool(true) ,
                "cancelOrders":MkBool(true) ,
                "createOrder":MkBool(true) ,
                "createMarketOrder":MkBool(false) ,
                "fetchBalance":MkBool(true) ,
                "fetchDeposits":MkBool(true) ,
                "fetchClosedOrders":MkBool(true) ,
                "fetchMarkets":MkBool(true) ,
                "fetchOHLCV":MkBool(true) ,
                "fetchOpenOrders":MkBool(true) ,
                "fetchOrder":MkBool(true) ,
                "fetchOrderBook":MkBool(true) ,
                "fetchOrders":MkBool(true) ,
                "fetchTicker":MkBool(true) ,
                "fetchTickers":MkBool(true) ,
                "fetchTime":MkBool(true) ,
                "fetchTrades":MkBool(true) ,
                "fetchTransactions":MkBool(false) ,
                "fetchWithdrawals":MkBool(true) ,
                "withdraw":MkBool(true) ,
                }),
            "timeframes":MkMap(&VarMap{
                "1m":MkString("1min") ,
                "5m":MkString("5min") ,
                "15m":MkString("15min") ,
                "30m":MkString("30min") ,
                "1h":MkString("60min") ,
                "4h":MkString("4hour") ,
                "1d":MkString("1day") ,
                "5d":MkString("5day") ,
                "1w":MkString("1week") ,
                "1M":MkString("1mon") ,
                }),
            "hostname":MkString("apiv2.bitz.com") ,
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/51840849/87443304-fec5e000-c5fd-11ea-98f8-ba8e67f7eaff.jpg") ,
                "api":MkMap(&VarMap{
                    "market":MkString("https://{hostname}") ,
                    "trade":MkString("https://{hostname}") ,
                    "assets":MkString("https://{hostname}") ,
                    }),
                "www":MkString("https://www.bitz.com") ,
                "doc":MkString("https://apidocv2.bitz.plus/en/") ,
                "fees":MkString("https://www.bitz.com/fee?type=1") ,
                "referral":MkString("https://u.bitz.com/register?invite_code=1429193") ,
                }),
            "api":MkMap(&VarMap{
                "market":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("ticker") ,
                        MkString("depth") ,
                        MkString("order") ,
                        MkString("tickerall") ,
                        MkString("kline") ,
                        MkString("symbolList") ,
                        MkString("getServerTime") ,
                        MkString("currencyRate") ,
                        MkString("currencyCoinRate") ,
                        MkString("coinRate") ,
                        MkString("getContractCoin") ,
                        MkString("getContractKline") ,
                        MkString("getContractOrderBook") ,
                        MkString("getContractTradesHistory") ,
                        MkString("getContractTickers") ,
                        })}),
                "trade":MkMap(&VarMap{"post":MkArray(&VarArray{
                        MkString("addEntrustSheet") ,
                        MkString("cancelEntrustSheet") ,
                        MkString("cancelAllEntrustSheet") ,
                        MkString("coinOut") ,
                        MkString("getUserHistoryEntrustSheet") ,
                        MkString("getUserNowEntrustSheet") ,
                        MkString("getEntrustSheetInfo") ,
                        MkString("depositOrWithdraw") ,
                        MkString("getCoinAddress") ,
                        MkString("getCoinAddressList") ,
                        MkString("marketTrade") ,
                        MkString("addEntrustSheetBatch") ,
                        })}),
                "assets":MkMap(&VarMap{"post":MkArray(&VarArray{
                        MkString("getUserAssets") ,
                        })}),
                "contract":MkMap(&VarMap{"post":MkArray(&VarArray{
                        MkString("addContractTrade") ,
                        MkString("cancelContractTrade") ,
                        MkString("getContractActivePositions") ,
                        MkString("getContractAccountInfo") ,
                        MkString("getContractMyPositions") ,
                        MkString("getContractOrderResult") ,
                        MkString("getContractTradeResult") ,
                        MkString("getContractOrder") ,
                        MkString("getContractMyHistoryTrade") ,
                        MkString("getContractMyTrades") ,
                        })}),
                }),
            "fees":MkMap(&VarMap{
                "trading":MkMap(&VarMap{
                    "maker":this.ParseNumber(MkString("0.002") ),
                    "taker":this.ParseNumber(MkString("0.002") ),
                    }),
                "funding":MkMap(&VarMap{"withdraw":MkMap(&VarMap{})}),
                }),
            "precision":MkMap(&VarMap{
                "amount":MkInteger(8) ,
                "price":MkInteger(8) ,
                }),
            "options":MkMap(&VarMap{
                "fetchOHLCVVolume":MkBool(true) ,
                "fetchOHLCVWarning":MkBool(true) ,
                "lastNonceTimestamp":MkInteger(0) ,
                }),
            "commonCurrencies":MkMap(&VarMap{
                "BOX":MkString("BOX Token") ,
                "LEO":MkString("LeoCoin") ,
                "XRB":MkString("NANO") ,
                "PXC":MkString("Pixiecoin") ,
                "VTC":MkString("VoteCoin") ,
                "TTC":MkString("TimesChain") ,
                }),
            "exceptions":MkMap(&VarMap{
                "-102":ExchangeError ,
                "-103":AuthenticationError ,
                "-104":ExchangeNotAvailable ,
                "-105":AuthenticationError ,
                "-106":ExchangeNotAvailable ,
                "-109":AuthenticationError ,
                "-110":DDoSProtection ,
                "-111":PermissionDenied ,
                "-112":OnMaintenance ,
                "-114":RateLimitExceeded ,
                "-117":AuthenticationError ,
                "-100015":AuthenticationError ,
                "-100044":ExchangeError ,
                "-100101":ExchangeError ,
                "-100201":ExchangeError ,
                "-100301":ExchangeError ,
                "-100401":ExchangeError ,
                "-100302":ExchangeError ,
                "-100303":ExchangeError ,
                "-200003":AuthenticationError ,
                "-200005":PermissionDenied ,
                "-200025":ExchangeNotAvailable ,
                "-200027":InvalidOrder ,
                "-200028":InvalidOrder ,
                "-200029":InvalidOrder ,
                "-200030":InvalidOrder ,
                "-200031":InsufficientFunds ,
                "-200032":ExchangeError ,
                "-200033":ExchangeError ,
                "-200034":OrderNotFound ,
                "-200035":OrderNotFound ,
                "-200037":InvalidOrder ,
                "-200038":ExchangeError ,
                "-200055":OrderNotFound ,
                "-300069":AuthenticationError ,
                "-300101":ExchangeError ,
                "-300102":InvalidOrder ,
                "-300103":AuthenticationError ,
                "-301001":ExchangeNotAvailable ,
                }),
            }));
}

func (this *Bitz) FetchMarkets(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("marketGetSymbolList"), params ); _ = response;
  markets := this.SafeValue(response , MkString("data") ); _ = markets;
  ids := GoGetKeys(markets ); _ = ids;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , ids.Length )); OpIncr(& i ){
    id := *(ids ).At(i ); _ = id;
    market := *(markets ).At(id ); _ = market;
    numericId := this.SafeString(market , MkString("id") ); _ = numericId;
    baseId := this.SafeString(market , MkString("coinFrom") ); _ = baseId;
    quoteId := this.SafeString(market , MkString("coinTo") ); _ = quoteId;
    base := baseId.ToUpperCase(); _ = base;
    quote := quoteId.ToUpperCase(); _ = quote;
    base = this.SafeCurrencyCode(base );
    quote = this.SafeCurrencyCode(quote );
    symbol := OpAdd(base , OpAdd(MkString("/") , quote )); _ = symbol;
    pricePrecisionString := this.SafeString(market , MkString("priceFloat") ); _ = pricePrecisionString;
    minPrice := this.ParsePrecision(pricePrecisionString ); _ = minPrice;
    precision := MkMap(&VarMap{
          "amount":this.SafeInteger(market , MkString("numberFloat") ),
          "price":parseInt(pricePrecisionString ),
          }); _ = precision;
    minAmount := this.SafeString(market , MkString("minTrade") ); _ = minAmount;
    result.Push(MkMap(&VarMap{
            "info":market ,
            "id":id ,
            "numericId":numericId ,
            "symbol":symbol ,
            "base":base ,
            "quote":quote ,
            "baseId":baseId ,
            "quoteId":quoteId ,
            "active":MkBool(true) ,
            "precision":precision ,
            "limits":MkMap(&VarMap{
                "amount":MkMap(&VarMap{
                    "min":this.ParseNumber(minAmount ),
                    "max":this.SafeNumber(market , MkString("maxTrade") ),
                    }),
                "price":MkMap(&VarMap{
                    "min":this.ParseNumber(minPrice ),
                    "max":MkUndefined() ,
                    }),
                "cost":MkMap(&VarMap{
                    "min":this.ParseNumber(Precise.StringMul(minPrice , minAmount )),
                    "max":MkUndefined() ,
                    }),
                }),
            }));
  }
  return result ;
}

func (this *Bitz) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("assetsPostGetUserAssets"), params ); _ = response;
  balances := this.SafeValue(*(response ).At(MkString("data") ), MkString("info") ); _ = balances;
  timestamp := this.ParseMicrotime(this.SafeString(response , MkString("microtime") )); _ = timestamp;
  result := MkMap(&VarMap{
        "info":response ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        }); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , balances.Length )); OpIncr(& i ){
    balance := *(balances ).At(i ); _ = balance;
    currencyId := this.SafeString(balance , MkString("name") ); _ = currencyId;
    code := this.SafeCurrencyCode(currencyId ); _ = code;
    account := this.Account(); _ = account;
    *(account ).At(MkString("used") )= this.SafeString(balance , MkString("lock") );
    *(account ).At(MkString("total") )= this.SafeString(balance , MkString("num") );
    *(account ).At(MkString("free") )= this.SafeString(balance , MkString("over") );
    *(result ).At(code )= OpCopy(account );
  }
  return this.ParseBalance(result );
}

func (this *Bitz) ParseTicker(goArgs ...*Variant) *Variant{
  ticker := GoGetArg(goArgs, 0, MkUndefined()); _ = ticker;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := OpCopy(MkUndefined() ); _ = timestamp;
  marketId := this.SafeString(ticker , MkString("symbol") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market , MkString("_") ); _ = symbol;
  last := this.SafeNumber(ticker , MkString("now") ); _ = last;
  open := this.SafeNumber(ticker , MkString("open") ); _ = open;
  change := OpCopy(MkUndefined() ); _ = change;
  average := OpCopy(MkUndefined() ); _ = average;
  if IsTrue(OpAnd(OpNotEq2(last , MkUndefined() ), OpNotEq2(open , MkUndefined() ))) {
    change = OpSub(last , open );
    average = OpDiv(this.Sum(last , open ), MkInteger(2) );
  }
  return MkMap(&VarMap{
        "symbol":symbol ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "high":this.SafeNumber(ticker , MkString("high") ),
        "low":this.SafeNumber(ticker , MkString("low") ),
        "bid":this.SafeNumber(ticker , MkString("bidPrice") ),
        "bidVolume":this.SafeNumber(ticker , MkString("bidQty") ),
        "ask":this.SafeNumber(ticker , MkString("askPrice") ),
        "askVolume":this.SafeNumber(ticker , MkString("askQty") ),
        "vwap":MkUndefined() ,
        "open":open ,
        "close":last ,
        "last":last ,
        "previousClose":MkUndefined() ,
        "change":change ,
        "percentage":this.SafeNumber(ticker , MkString("priceChange24h") ),
        "average":average ,
        "baseVolume":this.SafeNumber(ticker , MkString("volume") ),
        "quoteVolume":this.SafeNumber(ticker , MkString("quoteVolume") ),
        "info":ticker ,
        });
}

func (this *Bitz) ParseMicrotime(goArgs ...*Variant) *Variant{
  microtime := GoGetArg(goArgs, 0, MkUndefined()); _ = microtime;
  if IsTrue(OpEq2(microtime , MkUndefined() )) {
    return microtime ;
  }
  parts := microtime.Split(MkString(" ") ); _ = parts;
  milliseconds := parseFloat(*(parts ).At(MkInteger(0) )); _ = milliseconds;
  seconds := parseInt(*(parts ).At(MkInteger(1) )); _ = seconds;
  total := this.Sum(seconds , milliseconds ); _ = total;
  return parseInt(OpMulti(total , MkInteger(1000) ));
}

func (this *Bitz) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"symbol":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("marketGetTicker"), this.Extend(request , params )); _ = response;
  ticker := this.ParseTicker(*(response ).At(MkString("data") ), market ); _ = ticker;
  timestamp := this.ParseMicrotime(this.SafeString(response , MkString("microtime") )); _ = timestamp;
  return this.Extend(ticker , MkMap(&VarMap{
            "timestamp":timestamp ,
            "datetime":this.Iso8601(timestamp ),
            }));
}

func (this *Bitz) FetchTickers(goArgs ...*Variant) *Variant{
  symbols := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbols;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  if IsTrue(OpNotEq2(symbols , MkUndefined() )) {
    ids := this.MarketIds(symbols ); _ = ids;
    *(request ).At(MkString("symbols") )= ids.Join(MkString(",") );
  }
  response := this.Call(MkString("marketGetTickerall"), this.Extend(request , params )); _ = response;
  tickers := this.SafeValue(response , MkString("data") ); _ = tickers;
  timestamp := this.ParseMicrotime(this.SafeString(response , MkString("microtime") )); _ = timestamp;
  result := MkMap(&VarMap{}); _ = result;
  ids := GoGetKeys(tickers ); _ = ids;
  for i := MkInteger(0) ; IsTrue(OpLw(i , ids.Length )); OpIncr(& i ){
    id := *(ids ).At(i ); _ = id;
    ticker := *(tickers ).At(id ); _ = ticker;
    market := OpCopy(MkUndefined() ); _ = market;
    if IsTrue(OpHasMember(id , *this.At(MkString("markets_by_id")) )) {
      market = *(*this.At(MkString("markets_by_id")) ).At(id );
    }
    ticker = this.ParseTicker(*(tickers ).At(id ), market );
    symbol := *(ticker ).At(MkString("symbol") ); _ = symbol;
    if IsTrue(OpEq2(symbol , MkUndefined() )) {
      if IsTrue(OpNotEq2(market , MkUndefined() )) {
        symbol = *(market ).At(MkString("symbol") );
      } else {
        baseId := MkUndefined(); _ = baseId;
        quoteId := MkUndefined(); _ = quoteId;
        ArrayUnpack(id.Split(MkString("_") ), &baseId, &quoteId);
        base := this.SafeCurrencyCode(baseId ); _ = base;
        quote := this.SafeCurrencyCode(quoteId ); _ = quote;
        symbol = OpAdd(base , OpAdd(MkString("/") , quote ));
      }
    }
    if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
      *(result ).At(symbol )= this.Extend(ticker , MkMap(&VarMap{
              "timestamp":timestamp ,
              "datetime":this.Iso8601(timestamp ),
              }));
    }
  }
  return this.FilterByArray(result , MkString("symbol") , symbols );
}

func (this *Bitz) FetchTime(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("marketGetGetServerTime"), params ); _ = response;
  return this.SafeTimestamp(response , MkString("time") );
}

func (this *Bitz) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"symbol":this.MarketId(symbol )}); _ = request;
  response := this.Call(MkString("marketGetDepth"), this.Extend(request , params )); _ = response;
  orderbook := this.SafeValue(response , MkString("data") ); _ = orderbook;
  timestamp := this.ParseMicrotime(this.SafeString(response , MkString("microtime") )); _ = timestamp;
  return this.ParseOrderBook(orderbook , symbol , timestamp );
}

func (this *Bitz) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  id := this.SafeString(trade , MkString("id") ); _ = id;
  timestamp := this.SafeTimestamp(trade , MkString("T") ); _ = timestamp;
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  if IsTrue(OpNotEq2(market , MkUndefined() )) {
    symbol = *(market ).At(MkString("symbol") );
  }
  priceString := this.SafeString(trade , MkString("p") ); _ = priceString;
  amountString := this.SafeString(trade , MkString("n") ); _ = amountString;
  price := this.ParseNumber(priceString ); _ = price;
  amount := this.ParseNumber(amountString ); _ = amount;
  cost := this.ParseNumber(Precise.StringMul(priceString , amountString )); _ = cost;
  side := this.SafeString(trade , MkString("s") ); _ = side;
  return MkMap(&VarMap{
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "symbol":symbol ,
        "id":id ,
        "order":MkUndefined() ,
        "type":MkString("limit") ,
        "side":side ,
        "takerOrMaker":MkUndefined() ,
        "price":price ,
        "amount":amount ,
        "cost":cost ,
        "fee":MkUndefined() ,
        "info":trade ,
        });
}

func (this *Bitz) FetchTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"symbol":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("marketGetOrder"), this.Extend(request , params )); _ = response;
  return this.ParseTrades(*(response ).At(MkString("data") ), market , since , limit );
}

func (this *Bitz) ParseOHLCV(goArgs ...*Variant) *Variant{
  ohlcv := GoGetArg(goArgs, 0, MkUndefined()); _ = ohlcv;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  return MkArray(&VarArray{
        this.SafeInteger(ohlcv , MkString("time") ),
        this.SafeNumber(ohlcv , MkString("open") ),
        this.SafeNumber(ohlcv , MkString("high") ),
        this.SafeNumber(ohlcv , MkString("low") ),
        this.SafeNumber(ohlcv , MkString("close") ),
        this.SafeNumber(ohlcv , MkString("volume") ),
        });
}

func (this *Bitz) FetchOHLCV(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  timeframe := GoGetArg(goArgs, 1, MkString("1m") ); _ = timeframe;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  duration := this.ParseTimeframe(timeframe ); _ = duration;
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "symbol":*(market ).At(MkString("id") ),
        "resolution":*(*this.At(MkString("timeframes")) ).At(timeframe ),
        }); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("size") )= Math.Min(limit , MkInteger(300) );
    if IsTrue(OpNotEq2(since , MkUndefined() )) {
      *(request ).At(MkString("to") )= this.Sum(since , OpMulti(limit , OpMulti(duration , MkInteger(1000) )));
    }
  } else {
    if IsTrue(OpNotEq2(since , MkUndefined() )) {
      panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchOHLCV() requires a limit argument if the since argument is specified") )));
    }
  }
  response := this.Call(MkString("marketGetKline"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  bars := this.SafeValue(data , MkString("bars") , MkArray(&VarArray{})); _ = bars;
  return this.ParseOHLCVs(bars , market , timeframe , since , limit );
}

func (this *Bitz) ParseOrderStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "0":MkString("open") ,
        "1":MkString("open") ,
        "2":MkString("closed") ,
        "3":MkString("canceled") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Bitz) ParseOrder(goArgs ...*Variant) *Variant{
  order := GoGetArg(goArgs, 0, MkUndefined()); _ = order;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  id := this.SafeString(order , MkString("id") ); _ = id;
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  if IsTrue(OpEq2(market , MkUndefined() )) {
    baseId := this.SafeString(order , MkString("coinFrom") ); _ = baseId;
    quoteId := this.SafeString(order , MkString("coinTo") ); _ = quoteId;
    if IsTrue(OpAnd(OpNotEq2(baseId , MkUndefined() ), OpNotEq2(quoteId , MkUndefined() ))) {
      marketId := OpAdd(baseId , OpAdd(MkString("_") , quoteId )); _ = marketId;
      if IsTrue(OpHasMember(marketId , *this.At(MkString("markets_by_id")) )) {
        market = this.SafeValue(*this.At(MkString("markets_by_id")) , marketId );
      } else {
        base := this.SafeCurrencyCode(baseId ); _ = base;
        quote := this.SafeCurrencyCode(quoteId ); _ = quote;
        symbol = OpAdd(base , OpAdd(MkString("/") , quote ));
      }
    }
  }
  if IsTrue(OpNotEq2(market , MkUndefined() )) {
    symbol = *(market ).At(MkString("symbol") );
  }
  side := this.SafeString(order , MkString("flag") ); _ = side;
  if IsTrue(OpNotEq2(side , MkUndefined() )) {
    side = OpTrinary(OpEq2(side , MkString("sale") ), MkString("sell") , MkString("buy") );
  }
  price := this.SafeNumber(order , MkString("price") ); _ = price;
  amount := this.SafeNumber(order , MkString("number") ); _ = amount;
  remaining := this.SafeNumber(order , MkString("numberOver") ); _ = remaining;
  filled := this.SafeNumber(order , MkString("numberDeal") ); _ = filled;
  timestamp := this.SafeInteger(order , MkString("timestamp") ); _ = timestamp;
  if IsTrue(OpEq2(timestamp , MkUndefined() )) {
    timestamp = this.SafeTimestamp(order , MkString("created") );
  }
  cost := this.SafeNumber(order , MkString("orderTotalPrice") ); _ = cost;
  status := this.ParseOrderStatus(this.SafeString(order , MkString("status") )); _ = status;
  return this.SafeOrder(MkMap(&VarMap{
            "id":id ,
            "clientOrderId":MkUndefined() ,
            "datetime":this.Iso8601(timestamp ),
            "timestamp":timestamp ,
            "lastTradeTimestamp":MkUndefined() ,
            "status":status ,
            "symbol":symbol ,
            "type":MkString("limit") ,
            "timeInForce":MkUndefined() ,
            "postOnly":MkUndefined() ,
            "side":side ,
            "price":price ,
            "stopPrice":MkUndefined() ,
            "cost":cost ,
            "amount":amount ,
            "filled":filled ,
            "remaining":remaining ,
            "trades":MkUndefined() ,
            "fee":MkUndefined() ,
            "info":order ,
            "average":MkUndefined() ,
            }));
}

func (this *Bitz) CreateOrder(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 2, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 3, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 4, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 5, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  if IsTrue(OpNotEq2(type_ , MkString("limit") )) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , MkString(" createOrder allows limit orders only") )));
  }
  market := this.Market(symbol ); _ = market;
  orderType := OpTrinary(OpEq2(side , MkString("buy") ), MkString("1") , MkString("2") ); _ = orderType;
  if IsTrue(OpNot(*this.At(MkString("password")) )) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , MkString(" createOrder() requires you to set exchange.password = \"YOUR_TRADING_PASSWORD\" (a trade password is NOT THE SAME as your login password)") )));
  }
  request := MkMap(&VarMap{
        "symbol":*(market ).At(MkString("id") ),
        "type":orderType ,
        "price":this.PriceToPrecision(symbol , price ),
        "number":this.AmountToPrecision(symbol , amount ),
        "tradePwd":*this.At(MkString("password")) ,
        }); _ = request;
  response := this.Call(MkString("tradePostAddEntrustSheet"), this.Extend(request , params )); _ = response;
  timestamp := this.ParseMicrotime(this.SafeString(response , MkString("microtime") )); _ = timestamp;
  order := this.Extend(MkMap(&VarMap{"timestamp":timestamp }), *(response ).At(MkString("data") )); _ = order;
  return this.ParseOrder(order , market );
}

func (this *Bitz) CancelOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"entrustSheetId":id }); _ = request;
  response := this.Call(MkString("tradePostCancelEntrustSheet"), this.Extend(request , params )); _ = response;
  return response ;
}

func (this *Bitz) CancelOrders(goArgs ...*Variant) *Variant{
  ids := GoGetArg(goArgs, 0, MkUndefined()); _ = ids;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"ids":ids.Join(MkString(",") )}); _ = request;
  response := this.Call(MkString("tradePostCancelEntrustSheet"), this.Extend(request , params )); _ = response;
  return response ;
}

func (this *Bitz) FetchOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"entrustSheetId":id }); _ = request;
  response := this.Call(MkString("tradePostGetEntrustSheetInfo"), this.Extend(request , params )); _ = response;
  return this.ParseOrder(*(response ).At(MkString("data") ));
}

func (this *Bitz) FetchOrdersWithMethod(goArgs ...*Variant) *Variant{
  method := GoGetArg(goArgs, 0, MkUndefined()); _ = method;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchOpenOrders() requires a symbol argument") )));
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "coinFrom":*(market ).At(MkString("baseId") ),
        "coinTo":*(market ).At(MkString("quoteId") ),
        }); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("page") )= MkInteger(1) ;
    *(request ).At(MkString("pageSize") )= OpCopy(limit );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("startTime") )= parseInt(OpDiv(since , MkInteger(1000) ));
  }
  response := this.Call(method , this.Extend(request , params )); _ = response;
  orders := this.SafeValue(*(response ).At(MkString("data") ), MkString("data") , MkArray(&VarArray{})); _ = orders;
  return this.ParseOrders(orders , MkUndefined() , since , limit );
}

func (this *Bitz) FetchOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  return this.FetchOrdersWithMethod(MkString("tradePostGetUserHistoryEntrustSheet") , symbol , since , limit , params );
}

func (this *Bitz) FetchOpenOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  return this.FetchOrdersWithMethod(MkString("tradePostGetUserNowEntrustSheet") , symbol , since , limit , params );
}

func (this *Bitz) FetchClosedOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  return this.FetchOrdersWithMethod(MkString("tradePostGetUserHistoryEntrustSheet") , symbol , since , limit , params );
}

func (this *Bitz) ParseTransactionStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "1":MkString("pending") ,
        "2":MkString("pending") ,
        "3":MkString("pending") ,
        "4":MkString("ok") ,
        "5":MkString("canceled") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Bitz) ParseTransaction(goArgs ...*Variant) *Variant{
  transaction := GoGetArg(goArgs, 0, MkUndefined()); _ = transaction;
  currency := GoGetArg(goArgs, 1, MkUndefined() ); _ = currency;
  timestamp := this.SafeInteger(transaction , MkString("updated") ); _ = timestamp;
  if IsTrue(OpEq2(timestamp , MkInteger(0) )) {
    timestamp = OpCopy(MkUndefined() );
  }
  currencyId := this.SafeString(transaction , MkString("coin") ); _ = currencyId;
  code := this.SafeCurrencyCode(currencyId , currency ); _ = code;
  type_ := this.SafeStringLower(transaction , MkString("type") ); _ = type_;
  status := this.ParseTransactionStatus(this.SafeString(transaction , MkString("status") )); _ = status;
  fee := OpCopy(MkUndefined() ); _ = fee;
  feeCost := this.SafeNumber(transaction , MkString("network_fee") ); _ = feeCost;
  if IsTrue(OpNotEq2(feeCost , MkUndefined() )) {
    fee = MkMap(&VarMap{
        "cost":feeCost ,
        "code":code ,
        });
  }
  return MkMap(&VarMap{
        "id":this.SafeString(transaction , MkString("id") ),
        "txid":this.SafeString(transaction , MkString("txid") ),
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "address":this.SafeString(transaction , MkString("wallet") ),
        "tag":this.SafeString(transaction , MkString("memo") ),
        "type":type_ ,
        "amount":this.SafeNumber(transaction , MkString("number") ),
        "currency":code ,
        "status":status ,
        "updated":timestamp ,
        "fee":fee ,
        "info":transaction ,
        });
}

func (this *Bitz) ParseTransactionsByType(goArgs ...*Variant) *Variant{
  type_ := GoGetArg(goArgs, 0, MkUndefined()); _ = type_;
  transactions := GoGetArg(goArgs, 1, MkUndefined()); _ = transactions;
  code := GoGetArg(goArgs, 2, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 3, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 4, MkUndefined() ); _ = limit;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , transactions.Length )); OpIncr(& i ){
    transaction := this.ParseTransaction(this.Extend(MkMap(&VarMap{"type":type_ }), *(transactions ).At(i ))); _ = transaction;
    result.Push(transaction );
  }
  return this.FilterByCurrencySinceLimit(result , code , since , limit );
}

func (this *Bitz) ParseTransactionType(goArgs ...*Variant) *Variant{
  type_ := GoGetArg(goArgs, 0, MkUndefined()); _ = type_;
  types := MkMap(&VarMap{
        "deposit":MkInteger(1) ,
        "withdrawal":MkInteger(2) ,
        }); _ = types;
  return this.SafeInteger(types , type_ , type_ );
}

func (this *Bitz) FetchDeposits(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  return this.FetchTransactionsForType(MkString("deposit") , code , since , limit , params );
}

func (this *Bitz) FetchWithdrawals(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  return this.FetchTransactionsForType(MkString("withdrawal") , code , since , limit , params );
}

func (this *Bitz) FetchTransactionsForType(goArgs ...*Variant) *Variant{
  type_ := GoGetArg(goArgs, 0, MkUndefined()); _ = type_;
  code := GoGetArg(goArgs, 1, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(code , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchTransactions() requires a currency `code` argument") )));
  }
  this.LoadMarkets();
  currency := this.Currency(code ); _ = currency;
  request := MkMap(&VarMap{
        "coin":*(currency ).At(MkString("id") ),
        "type":this.ParseTransactionType(type_ ),
        }); _ = request;
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("startTime") )= (parseInt(OpDiv(since , MkInteger(1000) ))).Call(MkString("toString") );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("page") )= MkInteger(1) ;
    *(request ).At(MkString("pageSize") )= OpCopy(limit );
  }
  response := this.Call(MkString("tradePostDepositOrWithdraw"), this.Extend(request , params )); _ = response;
  transactions := this.SafeValue(*(response ).At(MkString("data") ), MkString("data") , MkArray(&VarArray{})); _ = transactions;
  return this.ParseTransactionsByType(type_ , transactions , code , since , limit );
}

func (this *Bitz) Withdraw(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  amount := GoGetArg(goArgs, 1, MkUndefined()); _ = amount;
  address := GoGetArg(goArgs, 2, MkUndefined()); _ = address;
  tag := GoGetArg(goArgs, 3, MkUndefined() ); _ = tag;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.CheckAddress(address );
  this.LoadMarkets();
  currency := this.Currency(code ); _ = currency;
  request := MkMap(&VarMap{
        "coin":*(currency ).At(MkString("id") ),
        "number":this.CurrencyToPrecision(code , amount ),
        "address":address ,
        }); _ = request;
  if IsTrue(OpNotEq2(tag , MkUndefined() )) {
    *(request ).At(MkString("memo") )= OpCopy(tag );
  }
  response := this.Call(MkString("tradePostCoinOut"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  return this.ParseTransaction(data , currency );
}

func (this *Bitz) Nonce(goArgs ...*Variant) *Variant{
  currentTimestamp := this.Seconds(); _ = currentTimestamp;
  if IsTrue(OpGt(currentTimestamp , *(*this.At(MkString("options")) ).At(MkString("lastNonceTimestamp") ))) {
    *(*this.At(MkString("options")) ).At(MkString("lastNonceTimestamp") )= OpCopy(currentTimestamp );
    *(*this.At(MkString("options")) ).At(MkString("lastNonce") )= MkInteger(100000) ;
  }
  *(*this.At(MkString("options")) ).At(MkString("lastNonce") )= this.Sum(*(*this.At(MkString("options")) ).At(MkString("lastNonce") ), MkInteger(1) );
  return *(*this.At(MkString("options")) ).At(MkString("lastNonce") );
}

func (this *Bitz) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("market") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  baseUrl := this.ImplodeHostname(*(*(*this.At(MkString("urls")) ).At(MkString("api") )).At(api )); _ = baseUrl;
  url := OpAdd(baseUrl , OpAdd(MkString("/") , OpAdd(this.Capitalize(api ), OpAdd(MkString("/") , path )))); _ = url;
  query := OpCopy(MkUndefined() ); _ = query;
  if IsTrue(OpEq2(api , MkString("market") )) {
    query = this.Urlencode(params );
    if IsTrue(query.Length ) {
      OpAddAssign(& url , OpAdd(MkString("?") , query ));
    }
  } else {
    this.CheckRequiredCredentials();
    body = this.Rawencode(this.Keysort(this.Extend(MkMap(&VarMap{
                    "apiKey":*this.At(MkString("apiKey")) ,
                    "timeStamp":this.Seconds(),
                    "nonce":this.Nonce(),
                    }), params )));
    OpAddAssign(& body , OpAdd(MkString("&sign=") , this.Hash(this.Encode(OpAdd(body , *this.At(MkString("secret")) )))));
    headers = MkMap(&VarMap{"Content-type":MkString("application/x-www-form-urlencoded") });
  }
  return MkMap(&VarMap{
        "url":url ,
        "method":method ,
        "body":body ,
        "headers":headers ,
        });
}

func (this *Bitz) HandleErrors(goArgs ...*Variant) *Variant{
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
  status := this.SafeString(response , MkString("status") ); _ = status;
  if IsTrue(OpNotEq2(status , MkUndefined() )) {
    feedback := OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , body )); _ = feedback;
    if IsTrue(OpEq2(status , MkString("200") )) {
      code := this.SafeInteger(response , MkString("data") ); _ = code;
      if IsTrue(OpNotEq2(code , MkUndefined() )) {
        this.ThrowExactlyMatchedException(*this.At(MkString("exceptions")) , code , feedback );
        panic(NewExchangeError(feedback ));
      } else {
        return MkUndefined();
      }
    }
    this.ThrowExactlyMatchedException(*this.At(MkString("exceptions")) , status , feedback );
    panic(NewExchangeError(feedback ));
  }
  return MkUndefined()
}

