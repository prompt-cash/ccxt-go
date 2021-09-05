package ccxt

import (
)

type Currencycom struct {
	*ExchangeBase
}

var _ Exchange = (*Currencycom)(nil)

func init() {
	exchange := &Currencycom{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Currencycom) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("currencycom") ,
            "name":MkString("Currency.com") ,
            "countries":MkArray(&VarArray{
                MkString("BY") ,
                }),
            "rateLimit":MkInteger(500) ,
            "certified":MkBool(true) ,
            "pro":MkBool(true) ,
            "version":MkString("v1") ,
            "has":MkMap(&VarMap{
                "cancelOrder":MkBool(true) ,
                "CORS":MkBool(false) ,
                "createOrder":MkBool(true) ,
                "fetchAccounts":MkBool(true) ,
                "fetchBalance":MkBool(true) ,
                "fetchMarkets":MkBool(true) ,
                "fetchMyTrades":MkBool(true) ,
                "fetchOHLCV":MkBool(true) ,
                "fetchOpenOrders":MkBool(true) ,
                "fetchOrderBook":MkBool(true) ,
                "fetchTicker":MkBool(true) ,
                "fetchTickers":MkBool(true) ,
                "fetchTime":MkBool(true) ,
                "fetchTradingFees":MkBool(true) ,
                "fetchTrades":MkBool(true) ,
                }),
            "timeframes":MkMap(&VarMap{
                "1m":MkString("1m") ,
                "3m":MkString("3m") ,
                "5m":MkString("5m") ,
                "15m":MkString("15m") ,
                "30m":MkString("30m") ,
                "1h":MkString("1h") ,
                "4h":MkString("4h") ,
                "1d":MkString("1d") ,
                "1w":MkString("1w") ,
                }),
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/1294454/83718672-36745c00-a63e-11ea-81a9-677b1f789a4d.jpg") ,
                "api":MkMap(&VarMap{
                    "public":MkString("https://api-adapter.backend.currency.com/api") ,
                    "private":MkString("https://api-adapter.backend.currency.com/api") ,
                    }),
                "test":MkMap(&VarMap{
                    "public":MkString("https://demo-api-adapter.backend.currency.com/api") ,
                    "private":MkString("https://demo-api-adapter.backend.currency.com/api") ,
                    }),
                "www":MkString("https://www.currency.com") ,
                "referral":MkString("https://currency.com/trading/signup?c=362jaimv&pid=referral") ,
                "doc":MkArray(&VarArray{
                    MkString("https://currency.com/api") ,
                    }),
                "fees":MkString("https://currency.com/fees-charges") ,
                }),
            "api":MkMap(&VarMap{
                "public":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("time") ,
                        MkString("exchangeInfo") ,
                        MkString("depth") ,
                        MkString("aggTrades") ,
                        MkString("klines") ,
                        MkString("ticker/24hr") ,
                        })}),
                "private":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("leverageSettings") ,
                        MkString("openOrders") ,
                        MkString("tradingPositions") ,
                        MkString("account") ,
                        MkString("myTrades") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("order") ,
                        MkString("updateTradingPosition") ,
                        MkString("updateTradingOrder") ,
                        MkString("closeTradingPosition") ,
                        }),
                    "delete":MkArray(&VarArray{
                        MkString("order") ,
                        }),
                    }),
                }),
            "fees":MkMap(&VarMap{"trading":MkMap(&VarMap{
                    "feeSide":MkString("get") ,
                    "tierBased":MkBool(false) ,
                    "percentage":MkBool(true) ,
                    "taker":this.ParseNumber(MkString("0.002") ),
                    "maker":this.ParseNumber(MkString("0.002") ),
                    })}),
            "precisionMode":TICK_SIZE ,
            "options":MkMap(&VarMap{
                "defaultTimeInForce":MkString("GTC") ,
                "warnOnFetchOpenOrdersWithoutSymbol":MkBool(true) ,
                "recvWindow":OpMulti(MkInteger(5) , MkInteger(1000) ),
                "timeDifference":MkInteger(0) ,
                "adjustForTimeDifference":MkBool(false) ,
                "parseOrderToPrecision":MkBool(false) ,
                "newOrderRespType":MkMap(&VarMap{
                    "market":MkString("FULL") ,
                    "limit":MkString("RESULT") ,
                    "stop":MkString("RESULT") ,
                    }),
                }),
            "exceptions":MkMap(&VarMap{
                "broad":MkMap(&VarMap{
                    "FIELD_VALIDATION_ERROR Cancel is available only for LIMIT order":InvalidOrder ,
                    "API key does not exist":AuthenticationError ,
                    "Order would trigger immediately.":InvalidOrder ,
                    "Account has insufficient balance for requested action.":InsufficientFunds ,
                    "Rest API trading is not enabled.":ExchangeNotAvailable ,
                    }),
                "exact":MkMap(&VarMap{
                    "-1000":ExchangeNotAvailable ,
                    "-1013":InvalidOrder ,
                    "-1021":InvalidNonce ,
                    "-1022":AuthenticationError ,
                    "-1100":InvalidOrder ,
                    "-1104":ExchangeError ,
                    "-1025":AuthenticationError ,
                    "-1128":BadRequest ,
                    "-2010":ExchangeError ,
                    "-2011":OrderNotFound ,
                    "-2013":OrderNotFound ,
                    "-2014":AuthenticationError ,
                    "-2015":AuthenticationError ,
                    }),
                }),
            "commonCurrencies":MkMap(&VarMap{
                "BNS":MkString("Bank of Nova Scotia") ,
                "EDU":MkString("New Oriental Education & Technology Group Inc") ,
                "ETN":MkString("Eaton") ,
                "IQ":MkString("iQIYI") ,
                "PLAY":MkString("Dave & Buster's Entertainment") ,
                }),
            }));
}

func (this *Currencycom) Nonce(goArgs ...*Variant) *Variant{
  return OpSub(this.Milliseconds(), *(*this.At(MkString("options")) ).At(MkString("timeDifference") ));
}

func (this *Currencycom) FetchTime(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetTime"), params ); _ = response;
  return this.SafeInteger(response , MkString("serverTime") );
}

func (this *Currencycom) LoadTimeDifference(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetTime"), params ); _ = response;
  after := this.Milliseconds(); _ = after;
  *(*this.At(MkString("options")) ).At(MkString("timeDifference") )= parseInt(OpSub(after , *(response ).At(MkString("serverTime") )));
  return *(*this.At(MkString("options")) ).At(MkString("timeDifference") );
}

func (this *Currencycom) FetchMarkets(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetExchangeInfo"), params ); _ = response;
  if IsTrue(*(*this.At(MkString("options")) ).At(MkString("adjustForTimeDifference") )) {
    this.LoadTimeDifference();
  }
  markets := this.SafeValue(response , MkString("symbols") ); _ = markets;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , markets.Length )); OpIncr(& i ){
    market := *(markets ).At(i ); _ = market;
    id := this.SafeString(market , MkString("symbol") ); _ = id;
    baseId := this.SafeString(market , MkString("baseAsset") ); _ = baseId;
    quoteId := this.SafeString(market , MkString("quoteAsset") ); _ = quoteId;
    base := this.SafeCurrencyCode(baseId ); _ = base;
    quote := this.SafeCurrencyCode(quoteId ); _ = quote;
    symbol := OpAdd(base , OpAdd(MkString("/") , quote )); _ = symbol;
    if IsTrue(OpGtEq(id.IndexOf(MkString("/") ), MkInteger(0) )) {
      symbol = OpCopy(id );
    }
    filters := this.SafeValue(market , MkString("filters") , MkArray(&VarArray{})); _ = filters;
    filtersByType := this.IndexBy(filters , MkString("filterType") ); _ = filtersByType;
    precision := MkMap(&VarMap{
          "amount":OpDiv(MkInteger(1) , Math.Pow(MkInteger(1) , this.SafeInteger(market , MkString("baseAssetPrecision") ))),
          "price":this.SafeNumber(market , MkString("tickSize") ),
          }); _ = precision;
    status := this.SafeString(market , MkString("status") ); _ = status;
    active := OpEq2(status , MkString("TRADING") ); _ = active;
    type_ := this.SafeStringLower(market , MkString("marketType") ); _ = type_;
    if IsTrue(OpEq2(type_ , MkString("leverage") )) {
      type_ = MkString("margin") ;
    }
    spot := OpEq2(type_ , MkString("spot") ); _ = spot;
    margin := OpEq2(type_ , MkString("margin") ); _ = margin;
    entry := MkMap(&VarMap{
          "id":id ,
          "symbol":symbol ,
          "base":base ,
          "quote":quote ,
          "baseId":baseId ,
          "quoteId":quoteId ,
          "type":type_ ,
          "spot":spot ,
          "margin":margin ,
          "info":market ,
          "active":active ,
          "precision":precision ,
          "limits":MkMap(&VarMap{
              "amount":MkMap(&VarMap{
                  "min":Math.Pow(MkInteger(10) , OpNeg(*(precision ).At(MkString("amount") ))),
                  "max":MkUndefined() ,
                  }),
              "price":MkMap(&VarMap{
                  "min":MkUndefined() ,
                  "max":MkUndefined() ,
                  }),
              "cost":MkMap(&VarMap{
                  "min":OpNeg(Math.Log10(*(precision ).At(MkString("amount") ))),
                  "max":MkUndefined() ,
                  }),
              }),
          }); _ = entry;
    exchangeFee := this.SafeNumber2(market , MkString("exchangeFee") , MkString("tradingFee") ); _ = exchangeFee;
    makerFee := this.SafeNumber(market , MkString("makerFee") , exchangeFee ); _ = makerFee;
    takerFee := this.SafeNumber(market , MkString("takerFee") , exchangeFee ); _ = takerFee;
    if IsTrue(OpNotEq2(makerFee , MkUndefined() )) {
      *(entry ).At(MkString("maker") )= OpDiv(makerFee , MkInteger(100) );
    }
    if IsTrue(OpNotEq2(takerFee , MkUndefined() )) {
      *(entry ).At(MkString("taker") )= OpDiv(takerFee , MkInteger(100) );
    }
    if IsTrue(OpHasMember(MkString("PRICE_FILTER") , filtersByType )) {
      filter := this.SafeValue(filtersByType , MkString("PRICE_FILTER") , MkMap(&VarMap{})); _ = filter;
      *(*(entry ).At(MkString("precision") )).At(MkString("price") )= this.SafeNumber(filter , MkString("tickSize") );
      *(*(entry ).At(MkString("limits") )).At(MkString("price") )= MkMap(&VarMap{
          "min":this.SafeNumber(filter , MkString("minPrice") ),
          "max":MkUndefined() ,
          });
      maxPrice := this.SafeNumber(filter , MkString("maxPrice") ); _ = maxPrice;
      if IsTrue(OpAnd(OpNotEq2(maxPrice , MkUndefined() ), OpGt(maxPrice , MkInteger(0) ))) {
        *(*(*(entry ).At(MkString("limits") )).At(MkString("price") )).At(MkString("max") )= OpCopy(maxPrice );
      }
    }
    if IsTrue(OpHasMember(MkString("LOT_SIZE") , filtersByType )) {
      filter := this.SafeValue(filtersByType , MkString("LOT_SIZE") , MkMap(&VarMap{})); _ = filter;
      *(*(entry ).At(MkString("precision") )).At(MkString("amount") )= this.SafeNumber(filter , MkString("stepSize") );
      *(*(entry ).At(MkString("limits") )).At(MkString("amount") )= MkMap(&VarMap{
          "min":this.SafeNumber(filter , MkString("minQty") ),
          "max":this.SafeNumber(filter , MkString("maxQty") ),
          });
    }
    if IsTrue(OpHasMember(MkString("MARKET_LOT_SIZE") , filtersByType )) {
      filter := this.SafeValue(filtersByType , MkString("MARKET_LOT_SIZE") , MkMap(&VarMap{})); _ = filter;
      *(*(entry ).At(MkString("limits") )).At(MkString("market") )= MkMap(&VarMap{
          "min":this.SafeNumber(filter , MkString("minQty") ),
          "max":this.SafeNumber(filter , MkString("maxQty") ),
          });
    }
    if IsTrue(OpHasMember(MkString("MIN_NOTIONAL") , filtersByType )) {
      filter := this.SafeValue(filtersByType , MkString("MIN_NOTIONAL") , MkMap(&VarMap{})); _ = filter;
      *(*(*(entry ).At(MkString("limits") )).At(MkString("cost") )).At(MkString("min") )= this.SafeNumber(filter , MkString("minNotional") );
    }
    result.Push(entry );
  }
  return result ;
}

func (this *Currencycom) FetchAccounts(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("privateGetAccount"), params ); _ = response;
  accounts := this.SafeValue(response , MkString("balances") , MkArray(&VarArray{})); _ = accounts;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , accounts.Length )); OpIncr(& i ){
    account := *(accounts ).At(i ); _ = account;
    accountId := this.SafeInteger(account , MkString("accountId") ); _ = accountId;
    currencyId := this.SafeString(account , MkString("asset") ); _ = currencyId;
    currencyCode := this.SafeCurrencyCode(currencyId ); _ = currencyCode;
    result.Push(MkMap(&VarMap{
            "id":accountId ,
            "type":MkUndefined() ,
            "currency":currencyCode ,
            "info":response ,
            }));
  }
  return result ;
}

func (this *Currencycom) FetchTradingFees(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("privateGetAccount"), params ); _ = response;
  return MkMap(&VarMap{
        "info":response ,
        "maker":this.SafeNumber(response , MkString("makerCommission") ),
        "taker":this.SafeNumber(response , MkString("takerCommission") ),
        });
}

func (this *Currencycom) ParseBalanceResponse(goArgs ...*Variant) *Variant{
  response := GoGetArg(goArgs, 0, MkUndefined()); _ = response;
  result := MkMap(&VarMap{"info":response }); _ = result;
  balances := this.SafeValue(response , MkString("balances") , MkArray(&VarArray{})); _ = balances;
  for i := MkInteger(0) ; IsTrue(OpLw(i , balances.Length )); OpIncr(& i ){
    balance := *(balances ).At(i ); _ = balance;
    currencyId := this.SafeString(balance , MkString("asset") ); _ = currencyId;
    code := this.SafeCurrencyCode(currencyId ); _ = code;
    account := this.Account(); _ = account;
    *(account ).At(MkString("free") )= this.SafeString(balance , MkString("free") );
    *(account ).At(MkString("used") )= this.SafeString(balance , MkString("locked") );
    *(result ).At(code )= OpCopy(account );
  }
  return this.ParseBalance(result );
}

func (this *Currencycom) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("privateGetAccount"), params ); _ = response;
  return this.ParseBalanceResponse(response );
}

func (this *Currencycom) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"symbol":*(market ).At(MkString("id") )}); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("publicGetDepth"), this.Extend(request , params )); _ = response;
  orderbook := this.ParseOrderBook(response , symbol ); _ = orderbook;
  *(orderbook ).At(MkString("nonce") )= this.SafeInteger(response , MkString("lastUpdateId") );
  return orderbook ;
}

func (this *Currencycom) ParseTicker(goArgs ...*Variant) *Variant{
  ticker := GoGetArg(goArgs, 0, MkUndefined()); _ = ticker;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.SafeInteger(ticker , MkString("closeTime") ); _ = timestamp;
  marketId := this.SafeString(ticker , MkString("symbol") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market ); _ = symbol;
  last := this.SafeNumber(ticker , MkString("lastPrice") ); _ = last;
  open := this.SafeNumber(ticker , MkString("openPrice") ); _ = open;
  average := OpCopy(MkUndefined() ); _ = average;
  if IsTrue(OpAnd(OpNotEq2(open , MkUndefined() ), OpNotEq2(last , MkUndefined() ))) {
    average = OpDiv(this.Sum(open , last ), MkInteger(2) );
  }
  return MkMap(&VarMap{
        "symbol":symbol ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "high":this.SafeNumber(ticker , MkString("highPrice") ),
        "low":this.SafeNumber(ticker , MkString("lowPrice") ),
        "bid":this.SafeNumber(ticker , MkString("bidPrice") ),
        "bidVolume":this.SafeNumber(ticker , MkString("bidQty") ),
        "ask":this.SafeNumber(ticker , MkString("askPrice") ),
        "askVolume":this.SafeNumber(ticker , MkString("askQty") ),
        "vwap":this.SafeNumber(ticker , MkString("weightedAvgPrice") ),
        "open":open ,
        "close":last ,
        "last":last ,
        "previousClose":this.SafeNumber(ticker , MkString("prevClosePrice") ),
        "change":this.SafeNumber(ticker , MkString("priceChange") ),
        "percentage":this.SafeNumber(ticker , MkString("priceChangePercent") ),
        "average":average ,
        "baseVolume":this.SafeNumber(ticker , MkString("volume") ),
        "quoteVolume":this.SafeNumber(ticker , MkString("quoteVolume") ),
        "info":ticker ,
        });
}

func (this *Currencycom) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"symbol":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("publicGetTicker24hr"), this.Extend(request , params )); _ = response;
  return this.ParseTicker(response , market );
}

func (this *Currencycom) FetchTickers(goArgs ...*Variant) *Variant{
  symbols := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbols;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("publicGetTicker24hr"), params ); _ = response;
  return this.ParseTickers(response , symbols );
}

func (this *Currencycom) ParseOHLCV(goArgs ...*Variant) *Variant{
  ohlcv := GoGetArg(goArgs, 0, MkUndefined()); _ = ohlcv;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  return MkArray(&VarArray{
        this.SafeInteger(ohlcv , MkInteger(0) ),
        this.SafeNumber(ohlcv , MkInteger(1) ),
        this.SafeNumber(ohlcv , MkInteger(2) ),
        this.SafeNumber(ohlcv , MkInteger(3) ),
        this.SafeNumber(ohlcv , MkInteger(4) ),
        this.SafeNumber(ohlcv , MkInteger(5) ),
        });
}

func (this *Currencycom) FetchOHLCV(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  timeframe := GoGetArg(goArgs, 1, MkString("1m") ); _ = timeframe;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "symbol":*(market ).At(MkString("id") ),
        "interval":*(*this.At(MkString("timeframes")) ).At(timeframe ),
        }); _ = request;
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("startTime") )= OpCopy(since );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("publicGetKlines"), this.Extend(request , params )); _ = response;
  return this.ParseOHLCVs(response , market , timeframe , since , limit );
}

func (this *Currencycom) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.SafeInteger2(trade , MkString("T") , MkString("time") ); _ = timestamp;
  priceString := this.SafeString2(trade , MkString("p") , MkString("price") ); _ = priceString;
  amountString := this.SafeString2(trade , MkString("q") , MkString("qty") ); _ = amountString;
  price := this.ParseNumber(priceString ); _ = price;
  amount := this.ParseNumber(amountString ); _ = amount;
  cost := this.ParseNumber(Precise.StringMul(priceString , amountString )); _ = cost;
  id := this.SafeString2(trade , MkString("a") , MkString("id") ); _ = id;
  side := OpCopy(MkUndefined() ); _ = side;
  orderId := this.SafeString(trade , MkString("orderId") ); _ = orderId;
  if IsTrue(OpHasMember(MkString("m") , trade )) {
    side = OpTrinary(*(trade ).At(MkString("m") ), MkString("sell") , MkString("buy") );
  } else {
    if IsTrue(OpHasMember(MkString("isBuyerMaker") , trade )) {
      side = OpTrinary(*(trade ).At(MkString("isBuyerMaker") ), MkString("sell") , MkString("buy") );
    } else {
      if IsTrue(OpHasMember(MkString("isBuyer") , trade )) {
        side = OpTrinary(*(trade ).At(MkString("isBuyer") ), MkString("buy") , MkString("sell") );
      }
    }
  }
  fee := OpCopy(MkUndefined() ); _ = fee;
  if IsTrue(OpHasMember(MkString("commission") , trade )) {
    fee = MkMap(&VarMap{
        "cost":this.SafeNumber(trade , MkString("commission") ),
        "currency":this.SafeCurrencyCode(this.SafeString(trade , MkString("commissionAsset") )),
        });
  }
  takerOrMaker := OpCopy(MkUndefined() ); _ = takerOrMaker;
  if IsTrue(OpHasMember(MkString("isMaker") , trade )) {
    takerOrMaker = OpTrinary(*(trade ).At(MkString("isMaker") ), MkString("maker") , MkString("taker") );
  }
  marketId := this.SafeString(trade , MkString("symbol") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market ); _ = symbol;
  return MkMap(&VarMap{
        "info":trade ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "symbol":symbol ,
        "id":id ,
        "order":orderId ,
        "type":MkUndefined() ,
        "takerOrMaker":takerOrMaker ,
        "side":side ,
        "price":price ,
        "amount":amount ,
        "cost":cost ,
        "fee":fee ,
        });
}

func (this *Currencycom) FetchTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"symbol":*(market ).At(MkString("id") )}); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("publicGetAggTrades"), this.Extend(request , params )); _ = response;
  return this.ParseTrades(response , market , since , limit );
}

func (this *Currencycom) ParseOrderStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "NEW":MkString("open") ,
        "PARTIALLY_FILLED":MkString("open") ,
        "FILLED":MkString("closed") ,
        "CANCELED":MkString("canceled") ,
        "PENDING_CANCEL":MkString("canceling") ,
        "REJECTED":MkString("rejected") ,
        "EXPIRED":MkString("expired") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Currencycom) ParseOrder(goArgs ...*Variant) *Variant{
  order := GoGetArg(goArgs, 0, MkUndefined()); _ = order;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  status := this.ParseOrderStatus(this.SafeString(order , MkString("status") )); _ = status;
  marketId := this.SafeString(order , MkString("symbol") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market , MkString("/") ); _ = symbol;
  timestamp := OpCopy(MkUndefined() ); _ = timestamp;
  if IsTrue(OpHasMember(MkString("time") , order )) {
    timestamp = this.SafeInteger(order , MkString("time") );
  } else {
    if IsTrue(OpHasMember(MkString("transactTime") , order )) {
      timestamp = this.SafeInteger(order , MkString("transactTime") );
    }
  }
  price := this.SafeNumber(order , MkString("price") ); _ = price;
  amount := this.SafeNumber(order , MkString("origQty") ); _ = amount;
  filled := this.SafeNumber(order , MkString("executedQty") ); _ = filled;
  remaining := OpCopy(MkUndefined() ); _ = remaining;
  cost := this.SafeNumber(order , MkString("cummulativeQuoteQty") ); _ = cost;
  id := this.SafeString(order , MkString("orderId") ); _ = id;
  type_ := this.SafeStringLower(order , MkString("type") ); _ = type_;
  side := this.SafeStringLower(order , MkString("side") ); _ = side;
  trades := OpCopy(MkUndefined() ); _ = trades;
  fills := this.SafeValue(order , MkString("fills") ); _ = fills;
  if IsTrue(OpNotEq2(fills , MkUndefined() )) {
    trades = this.ParseTrades(fills , market );
  }
  timeInForce := this.SafeString(order , MkString("timeInForce") ); _ = timeInForce;
  return this.SafeOrder(MkMap(&VarMap{
            "info":order ,
            "id":id ,
            "timestamp":timestamp ,
            "datetime":this.Iso8601(timestamp ),
            "lastTradeTimestamp":MkUndefined() ,
            "symbol":symbol ,
            "type":type_ ,
            "timeInForce":timeInForce ,
            "side":side ,
            "price":price ,
            "stopPrice":MkUndefined() ,
            "amount":amount ,
            "cost":cost ,
            "average":MkUndefined() ,
            "filled":filled ,
            "remaining":remaining ,
            "status":status ,
            "fee":MkUndefined() ,
            "trades":trades ,
            }));
}

func (this *Currencycom) CreateOrder(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 2, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 3, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 4, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 5, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  accountId := OpCopy(MkUndefined() ); _ = accountId;
  if IsTrue(*(market ).At(MkString("margin") )) {
    accountId = this.SafeInteger(params , MkString("accountId") );
    if IsTrue(OpEq2(accountId , MkUndefined() )) {
      panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" createOrder() requires an accountId parameter for ") , OpAdd(*(market ).At(MkString("type") ), OpAdd(MkString(" market ") , symbol ))))));
    }
  }
  uppercaseType := type_.ToUpperCase(); _ = uppercaseType;
  newOrderRespType := this.SafeValue(*(*this.At(MkString("options")) ).At(MkString("newOrderRespType") ), type_ , MkString("RESULT") ); _ = newOrderRespType;
  request := MkMap(&VarMap{
        "symbol":*(market ).At(MkString("id") ),
        "quantity":this.AmountToPrecision(symbol , amount ),
        "type":uppercaseType ,
        "side":side.ToUpperCase(),
        "newOrderRespType":newOrderRespType ,
        }); _ = request;
  if IsTrue(OpEq2(uppercaseType , MkString("LIMIT") )) {
    *(request ).At(MkString("price") )= this.PriceToPrecision(symbol , price );
    *(request ).At(MkString("timeInForce") )= *(*this.At(MkString("options")) ).At(MkString("defaultTimeInForce") );
  } else {
    if IsTrue(OpEq2(uppercaseType , MkString("STOP") )) {
      *(request ).At(MkString("price") )= this.PriceToPrecision(symbol , price );
    }
  }
  response := this.Call(MkString("privatePostOrder"), this.Extend(request , params )); _ = response;
  return this.ParseOrder(response , market );
}

func (this *Currencycom) FetchOpenOrders(goArgs ...*Variant) *Variant{
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
  } else {
    if IsTrue(*(*this.At(MkString("options")) ).At(MkString("warnOnFetchOpenOrdersWithoutSymbol") )) {
      symbols := OpCopy(*this.At(MkString("symbols")) ); _ = symbols;
      numSymbols := OpCopy(symbols.Length ); _ = numSymbols;
      fetchOpenOrdersRateLimit := parseInt(OpDiv(numSymbols , MkInteger(2) )); _ = fetchOpenOrdersRateLimit;
      panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" fetchOpenOrders() WARNING: fetching open orders without specifying a symbol is rate-limited to one call per ") , OpAdd(fetchOpenOrdersRateLimit.ToString(), OpAdd(MkString(" seconds. Do not call this method frequently to avoid ban. Set ") , OpAdd(*this.At(MkString("id")) , MkString(".options[\"warnOnFetchOpenOrdersWithoutSymbol\"] = false to suppress this warning message.") )))))));
    }
  }
  response := this.Call(MkString("privateGetOpenOrders"), this.Extend(request , params )); _ = response;
  return this.ParseOrders(response , market , since , limit );
}

func (this *Currencycom) CancelOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" cancelOrder() requires a symbol argument") )));
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  origClientOrderId := this.SafeValue(params , MkString("origClientOrderId") ); _ = origClientOrderId;
  request := MkMap(&VarMap{"symbol":*(market ).At(MkString("id") )}); _ = request;
  if IsTrue(OpEq2(origClientOrderId , MkUndefined() )) {
    *(request ).At(MkString("orderId") )= OpCopy(id );
  } else {
    *(request ).At(MkString("origClientOrderId") )= OpCopy(origClientOrderId );
  }
  response := this.Call(MkString("privateDeleteOrder"), this.Extend(request , params )); _ = response;
  return this.ParseOrder(response , market );
}

func (this *Currencycom) FetchMyTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchMyTrades() requires a symbol argument") )));
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"symbol":*(market ).At(MkString("id") )}); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("privateGetMyTrades"), this.Extend(request , params )); _ = response;
  return this.ParseTrades(response , market , since , limit );
}

func (this *Currencycom) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("public") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  url := OpAdd(*(*(*this.At(MkString("urls")) ).At(MkString("api") )).At(api ), OpAdd(MkString("/") , OpAdd(*this.At(MkString("version")) , OpAdd(MkString("/") , path )))); _ = url;
  if IsTrue(OpEq2(path , MkString("historicalTrades") )) {
    headers = MkMap(&VarMap{"X-MBX-APIKEY":*this.At(MkString("apiKey")) });
  }
  if IsTrue(OpEq2(api , MkString("private") )) {
    this.CheckRequiredCredentials();
    query := this.Urlencode(this.Extend(MkMap(&VarMap{
                  "timestamp":this.Nonce(),
                  "recvWindow":*(*this.At(MkString("options")) ).At(MkString("recvWindow") ),
                  }), params )); _ = query;
    signature := this.Hmac(this.Encode(query ), this.Encode(*this.At(MkString("secret")) )); _ = signature;
    OpAddAssign(& query , OpAdd(MkString("&") , OpAdd(MkString("signature=") , signature )));
    headers = MkMap(&VarMap{"X-MBX-APIKEY":*this.At(MkString("apiKey")) });
    if IsTrue(OpOr(OpEq2(method , MkString("GET") ), OpEq2(method , MkString("DELETE") ))) {
      OpAddAssign(& url , OpAdd(MkString("?") , query ));
    } else {
      body = OpCopy(query );
      *(headers ).At(MkString("Content-Type") )= MkString("application/x-www-form-urlencoded") ;
    }
  } else {
    if IsTrue(*(GoGetKeys(params )).At(MkString("length") )) {
      OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(params )));
    }
  }
  return MkMap(&VarMap{
        "url":url ,
        "method":method ,
        "body":body ,
        "headers":headers ,
        });
}

func (this *Currencycom) HandleErrors(goArgs ...*Variant) *Variant{
  httpCode := GoGetArg(goArgs, 0, MkUndefined()); _ = httpCode;
  reason := GoGetArg(goArgs, 1, MkUndefined()); _ = reason;
  url := GoGetArg(goArgs, 2, MkUndefined()); _ = url;
  method := GoGetArg(goArgs, 3, MkUndefined()); _ = method;
  headers := GoGetArg(goArgs, 4, MkUndefined()); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined()); _ = body;
  response := GoGetArg(goArgs, 6, MkUndefined()); _ = response;
  requestHeaders := GoGetArg(goArgs, 7, MkUndefined()); _ = requestHeaders;
  requestBody := GoGetArg(goArgs, 8, MkUndefined()); _ = requestBody;
  if IsTrue(OpOr(OpEq2(httpCode , MkInteger(418) ), OpEq2(httpCode , MkInteger(429) ))) {
    panic(NewDDoSProtection(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , OpAdd(httpCode.ToString(), OpAdd(MkString(" ") , OpAdd(reason , OpAdd(MkString(" ") , body ))))))));
  }
  if IsTrue(OpGtEq(httpCode , MkInteger(400) )) {
    if IsTrue(OpGtEq(body.IndexOf(MkString("Price * QTY is zero or less") ), MkInteger(0) )) {
      panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" order cost = amount * price is zero or less ") , body ))));
    }
    if IsTrue(OpGtEq(body.IndexOf(MkString("LOT_SIZE") ), MkInteger(0) )) {
      panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" order amount should be evenly divisible by lot size ") , body ))));
    }
    if IsTrue(OpGtEq(body.IndexOf(MkString("PRICE_FILTER") ), MkInteger(0) )) {
      panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" order price is invalid, i.e. exceeds allowed price precision, exceeds min price or max price limits or is invalid float value in general, use this.priceToPrecision (symbol, amount) ") , body ))));
    }
  }
  if IsTrue(OpEq2(response , MkUndefined() )) {
    return MkUndefined();
  }
  errorCode := this.SafeString(response , MkString("code") ); _ = errorCode;
  if IsTrue(OpAnd(OpNotEq2(errorCode , MkUndefined() ), OpNotEq2(errorCode , MkString("0") ))) {
    feedback := OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , this.Json(response ))); _ = feedback;
    this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("exact") ), errorCode , feedback );
    message := this.SafeString(response , MkString("msg") ); _ = message;
    this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("broad") ), message , feedback );
    panic(NewExchangeError(feedback ));
  }
  return MkUndefined()
}

