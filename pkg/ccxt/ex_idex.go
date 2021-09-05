package ccxt

import (
)

type Idex struct {
	*ExchangeBase
}

var _ Exchange = (*Idex)(nil)

func init() {
	exchange := &Idex{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Idex) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("idex") ,
            "name":MkString("IDEX") ,
            "countries":MkArray(&VarArray{
                MkString("US") ,
                }),
            "rateLimit":MkInteger(1500) ,
            "version":MkString("v2") ,
            "certified":MkBool(true) ,
            "pro":MkBool(true) ,
            "requiresWeb3":MkBool(true) ,
            "has":MkMap(&VarMap{
                "cancelOrder":MkBool(true) ,
                "createOrder":MkBool(true) ,
                "fetchBalance":MkBool(true) ,
                "fetchMarkets":MkBool(true) ,
                "fetchCurrencies":MkBool(true) ,
                "fetchMyTrades":MkBool(true) ,
                "fetchOHLCV":MkBool(true) ,
                "fetchOpenOrders":MkBool(true) ,
                "fetchClosedOrders":MkBool(true) ,
                "fetchOrders":MkBool(false) ,
                "fetchOrder":MkBool(true) ,
                "fetchOrderBook":MkBool(true) ,
                "fetchTicker":MkBool(true) ,
                "fetchTickers":MkBool(true) ,
                "fetchTrades":MkBool(true) ,
                "fetchTransactions":MkBool(false) ,
                "fetchDeposits":MkBool(true) ,
                "fetchWithdrawals":MkBool(true) ,
                "withdraw":MkBool(true) ,
                }),
            "timeframes":MkMap(&VarMap{
                "1m":MkString("1m") ,
                "5m":MkString("5m") ,
                "15m":MkString("15m") ,
                "30m":MkString("30m") ,
                "1h":MkString("1h") ,
                "6h":MkString("6h") ,
                "1d":MkString("1d") ,
                }),
            "urls":MkMap(&VarMap{
                "test":MkMap(&VarMap{
                    "public":MkString("https://api-sandbox.idex.io") ,
                    "private":MkString("https://api-sandbox.idex.io") ,
                    }),
                "logo":MkString("https://user-images.githubusercontent.com/51840849/94481303-2f222100-01e0-11eb-97dd-bc14c5943a86.jpg") ,
                "api":MkMap(&VarMap{
                    "ETH":MkString("https://api-eth.idex.io") ,
                    "BSC":MkString("https://api-bsc.idex.io") ,
                    }),
                "www":MkString("https://idex.io") ,
                "doc":MkArray(&VarArray{
                    MkString("https://docs.idex.io/") ,
                    }),
                }),
            "api":MkMap(&VarMap{
                "public":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("ping") ,
                        MkString("time") ,
                        MkString("exchange") ,
                        MkString("assets") ,
                        MkString("markets") ,
                        MkString("tickers") ,
                        MkString("candles") ,
                        MkString("trades") ,
                        MkString("orderbook") ,
                        MkString("wsToken") ,
                        })}),
                "private":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("user") ,
                        MkString("wallets") ,
                        MkString("balances") ,
                        MkString("orders") ,
                        MkString("fills") ,
                        MkString("deposits") ,
                        MkString("withdrawals") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("wallets") ,
                        MkString("orders") ,
                        MkString("orders/test") ,
                        MkString("withdrawals") ,
                        }),
                    "delete":MkArray(&VarArray{
                        MkString("orders") ,
                        }),
                    }),
                }),
            "options":MkMap(&VarMap{
                "defaultTimeInForce":MkString("gtc") ,
                "defaultSelfTradePrevention":MkString("cn") ,
                "network":MkString("ETH") ,
                }),
            "exceptions":MkMap(&VarMap{
                "INVALID_ORDER_QUANTITY":InvalidOrder ,
                "INSUFFICIENT_FUNDS":InsufficientFunds ,
                "SERVICE_UNAVAILABLE":ExchangeNotAvailable ,
                "EXCEEDED_RATE_LIMIT":DDoSProtection ,
                "INVALID_PARAMETER":BadRequest ,
                "WALLET_NOT_ASSOCIATED":InvalidAddress ,
                "INVALID_WALLET_SIGNATURE":AuthenticationError ,
                }),
            "requiredCredentials":MkMap(&VarMap{
                "walletAddress":MkBool(true) ,
                "privateKey":MkBool(true) ,
                "apiKey":MkBool(true) ,
                "secret":MkBool(true) ,
                }),
            "paddingMode":PAD_WITH_ZERO ,
            "commonCurrencies":MkMap(&VarMap{}),
            "requireCredentials":MkMap(&VarMap{
                "privateKey":MkBool(true) ,
                "apiKey":MkBool(true) ,
                "secret":MkBool(true) ,
                }),
            }));
}

func (this *Idex) FetchMarkets(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetMarkets"), params ); _ = response;
  response2 := this.Call(MkString("publicGetExchange")); _ = response2;
  maker := this.SafeNumber(response2 , MkString("makerFeeRate") ); _ = maker;
  taker := this.SafeNumber(response2 , MkString("takerFeeRate") ); _ = taker;
  makerMin := this.SafeNumber(response2 , MkString("makerTradeMinimum") ); _ = makerMin;
  takerMin := this.SafeNumber(response2 , MkString("takerTradeMinimum") ); _ = takerMin;
  minCostETH := Math.Min(makerMin , takerMin ); _ = minCostETH;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , response.Length )); OpIncr(& i ){
    entry := *(response ).At(i ); _ = entry;
    marketId := this.SafeString(entry , MkString("market") ); _ = marketId;
    baseId := this.SafeString(entry , MkString("baseAsset") ); _ = baseId;
    quoteId := this.SafeString(entry , MkString("quoteAsset") ); _ = quoteId;
    base := this.SafeCurrencyCode(baseId ); _ = base;
    quote := this.SafeCurrencyCode(quoteId ); _ = quote;
    symbol := OpAdd(base , OpAdd(MkString("/") , quote )); _ = symbol;
    basePrecisionString := this.SafeString(entry , MkString("baseAssetPrecision") ); _ = basePrecisionString;
    quotePrecisionString := this.SafeString(entry , MkString("quoteAssetPrecision") ); _ = quotePrecisionString;
    basePrecision := this.ParsePrecision(basePrecisionString ); _ = basePrecision;
    quotePrecision := this.ParsePrecision(quotePrecisionString ); _ = quotePrecision;
    status := this.SafeString(entry , MkString("status") ); _ = status;
    active := OpEq2(status , MkString("active") ); _ = active;
    minCost := OpCopy(MkUndefined() ); _ = minCost;
    if IsTrue(OpEq2(quote , MkString("ETH") )) {
      minCost = OpCopy(minCostETH );
    }
    precision := MkMap(&VarMap{
          "amount":parseInt(basePrecisionString ),
          "price":parseInt(quotePrecisionString ),
          }); _ = precision;
    result.Push(MkMap(&VarMap{
            "symbol":symbol ,
            "id":marketId ,
            "base":base ,
            "quote":quote ,
            "baseId":baseId ,
            "quoteId":quoteId ,
            "active":active ,
            "info":entry ,
            "precision":precision ,
            "taker":taker ,
            "maker":maker ,
            "limits":MkMap(&VarMap{
                "amount":MkMap(&VarMap{
                    "min":this.ParseNumber(basePrecision ),
                    "max":MkUndefined() ,
                    }),
                "price":MkMap(&VarMap{
                    "min":this.ParseNumber(quotePrecision ),
                    "max":MkUndefined() ,
                    }),
                "cost":MkMap(&VarMap{
                    "min":minCost ,
                    "max":MkUndefined() ,
                    }),
                }),
            }));
  }
  return result ;
}

func (this *Idex) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"market":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("publicGetTickers"), this.Extend(request , params )); _ = response;
  ticker := this.SafeValue(response , MkInteger(0) ); _ = ticker;
  return this.ParseTicker(ticker , market );
}

func (this *Idex) FetchTickers(goArgs ...*Variant) *Variant{
  symbols := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbols;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("publicGetTickers"), params ); _ = response;
  return this.ParseTickers(response , symbols );
}

func (this *Idex) ParseTicker(goArgs ...*Variant) *Variant{
  ticker := GoGetArg(goArgs, 0, MkUndefined()); _ = ticker;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  marketId := this.SafeString(ticker , MkString("market") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market , MkString("-") ); _ = symbol;
  baseVolume := this.SafeNumber(ticker , MkString("baseVolume") ); _ = baseVolume;
  quoteVolume := this.SafeNumber(ticker , MkString("quoteVolume") ); _ = quoteVolume;
  timestamp := this.SafeInteger(ticker , MkString("time") ); _ = timestamp;
  open := this.SafeNumber(ticker , MkString("open") ); _ = open;
  high := this.SafeNumber(ticker , MkString("high") ); _ = high;
  low := this.SafeNumber(ticker , MkString("low") ); _ = low;
  close := this.SafeNumber(ticker , MkString("close") ); _ = close;
  ask := this.SafeNumber(ticker , MkString("ask") ); _ = ask;
  bid := this.SafeNumber(ticker , MkString("bid") ); _ = bid;
  percentage := this.SafeNumber(ticker , MkString("percentChange") ); _ = percentage;
  if IsTrue(OpNotEq2(percentage , MkUndefined() )) {
    percentage = OpAdd(MkInteger(1) , OpDiv(percentage , MkInteger(100) ));
  }
  change := OpCopy(MkUndefined() ); _ = change;
  if IsTrue(OpAnd(OpNotEq2(close , MkUndefined() ), OpNotEq2(open , MkUndefined() ))) {
    change = OpSub(close , open );
  }
  return MkMap(&VarMap{
        "symbol":symbol ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "high":high ,
        "low":low ,
        "bid":bid ,
        "bidVolume":MkUndefined() ,
        "ask":ask ,
        "askVolume":MkUndefined() ,
        "vwap":MkUndefined() ,
        "open":open ,
        "close":close ,
        "last":close ,
        "previousClose":MkUndefined() ,
        "change":change ,
        "percentage":percentage ,
        "average":MkUndefined() ,
        "baseVolume":baseVolume ,
        "quoteVolume":quoteVolume ,
        "info":ticker ,
        });
}

func (this *Idex) FetchOHLCV(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  timeframe := GoGetArg(goArgs, 1, MkString("1m") ); _ = timeframe;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "market":*(market ).At(MkString("id") ),
        "interval":timeframe ,
        }); _ = request;
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("start") )= OpCopy(since );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("publicGetCandles"), this.Extend(request , params )); _ = response;
  if IsTrue(GoIsArray(response )) {
    return this.ParseOHLCVs(response , market , timeframe , since , limit );
  } else {
    return MkArray(&VarArray{});
  }
  return MkUndefined()
}

func (this *Idex) ParseOHLCV(goArgs ...*Variant) *Variant{
  ohlcv := GoGetArg(goArgs, 0, MkUndefined()); _ = ohlcv;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.SafeInteger(ohlcv , MkString("start") ); _ = timestamp;
  open := this.SafeNumber(ohlcv , MkString("open") ); _ = open;
  high := this.SafeNumber(ohlcv , MkString("high") ); _ = high;
  low := this.SafeNumber(ohlcv , MkString("low") ); _ = low;
  close := this.SafeNumber(ohlcv , MkString("close") ); _ = close;
  volume := this.SafeNumber(ohlcv , MkString("volume") ); _ = volume;
  return MkArray(&VarArray{
        timestamp ,
        open ,
        high ,
        low ,
        close ,
        volume ,
        });
}

func (this *Idex) FetchTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"market":*(market ).At(MkString("id") )}); _ = request;
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("start") )= OpCopy(since );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("publicGetTrades"), this.Extend(request , params )); _ = response;
  return this.ParseTrades(response , market , since , limit );
}

func (this *Idex) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  id := this.SafeString(trade , MkString("fillId") ); _ = id;
  priceString := this.SafeString(trade , MkString("price") ); _ = priceString;
  amountString := this.SafeString(trade , MkString("quantity") ); _ = amountString;
  price := this.ParseNumber(priceString ); _ = price;
  amount := this.ParseNumber(amountString ); _ = amount;
  cost := this.SafeNumber(trade , MkString("quoteQuantity") ); _ = cost;
  if IsTrue(OpEq2(cost , MkUndefined() )) {
    cost = this.ParseNumber(Precise.StringMul(priceString , amountString ));
  }
  timestamp := this.SafeInteger(trade , MkString("time") ); _ = timestamp;
  marketId := this.SafeString(trade , MkString("market") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market , MkString("-") ); _ = symbol;
  makerSide := this.SafeString(trade , MkString("makerSide") ); _ = makerSide;
  oppositeSide := OpTrinary(OpEq2(makerSide , MkString("buy") ), MkString("sell") , MkString("buy") ); _ = oppositeSide;
  side := this.SafeString(trade , MkString("side") , oppositeSide ); _ = side;
  takerOrMaker := this.SafeString(trade , MkString("liquidity") , MkString("taker") ); _ = takerOrMaker;
  feeCost := this.SafeNumber(trade , MkString("fee") ); _ = feeCost;
  fee := OpCopy(MkUndefined() ); _ = fee;
  if IsTrue(OpNotEq2(feeCost , MkUndefined() )) {
    feeCurrencyId := this.SafeString(trade , MkString("feeAsset") ); _ = feeCurrencyId;
    fee = MkMap(&VarMap{
        "cost":feeCost ,
        "currency":this.SafeCurrencyCode(feeCurrencyId ),
        });
  }
  orderId := this.SafeString(trade , MkString("orderId") ); _ = orderId;
  return MkMap(&VarMap{
        "info":trade ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "symbol":symbol ,
        "id":id ,
        "order":orderId ,
        "type":MkString("limit") ,
        "side":side ,
        "takerOrMaker":takerOrMaker ,
        "price":price ,
        "amount":amount ,
        "cost":cost ,
        "fee":fee ,
        });
}

func (this *Idex) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "market":*(market ).At(MkString("id") ),
        "level":MkInteger(2) ,
        }); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("publicGetOrderbook"), this.Extend(request , params )); _ = response;
  nonce := this.SafeInteger(response , MkString("sequence") ); _ = nonce;
  return MkMap(&VarMap{
        "symbol":symbol ,
        "timestamp":MkUndefined() ,
        "datetime":MkUndefined() ,
        "nonce":nonce ,
        "bids":this.ParseSide(response , MkString("bids") ),
        "asks":this.ParseSide(response , MkString("asks") ),
        });
}

func (this *Idex) ParseSide(goArgs ...*Variant) *Variant{
  book := GoGetArg(goArgs, 0, MkUndefined()); _ = book;
  side := GoGetArg(goArgs, 1, MkUndefined()); _ = side;
  bookSide := this.SafeValue(book , side , MkArray(&VarArray{})); _ = bookSide;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , bookSide.Length )); OpIncr(& i ){
    order := *(bookSide ).At(i ); _ = order;
    price := this.SafeNumber(order , MkInteger(0) ); _ = price;
    amount := this.SafeNumber(order , MkInteger(1) ); _ = amount;
    orderCount := this.SafeInteger(order , MkInteger(2) ); _ = orderCount;
    result.Push(MkArray(&VarArray{
            price ,
            amount ,
            orderCount ,
            }));
  }
  descending := OpEq2(side , MkString("bids") ); _ = descending;
  return this.SortBy(result , MkInteger(0) , descending );
}

func (this *Idex) FetchCurrencies(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetAssets"), params ); _ = response;
  result := MkMap(&VarMap{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , response.Length )); OpIncr(& i ){
    entry := *(response ).At(i ); _ = entry;
    name := this.SafeString(entry , MkString("name") ); _ = name;
    currencyId := this.SafeString(entry , MkString("symbol") ); _ = currencyId;
    precisionString := this.SafeString(entry , MkString("exchangeDecimals") ); _ = precisionString;
    code := this.SafeCurrencyCode(currencyId ); _ = code;
    precision := this.ParsePrecision(precisionString ); _ = precision;
    lot := this.ParseNumber(precision ); _ = lot;
    *(result ).At(code )= MkMap(&VarMap{
        "id":currencyId ,
        "code":code ,
        "info":entry ,
        "type":MkUndefined() ,
        "name":name ,
        "active":MkUndefined() ,
        "fee":MkUndefined() ,
        "precision":parseInt(precisionString ),
        "limits":MkMap(&VarMap{
            "amount":MkMap(&VarMap{
                "min":lot ,
                "max":MkUndefined() ,
                }),
            "withdraw":MkMap(&VarMap{
                "min":lot ,
                "max":MkUndefined() ,
                }),
            }),
        });
  }
  return result ;
}

func (this *Idex) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.CheckRequiredCredentials();
  this.LoadMarkets();
  nonce1 := this.Uuidv1(); _ = nonce1;
  request := MkMap(&VarMap{
        "nonce":nonce1 ,
        "wallet":*this.At(MkString("walletAddress")) ,
        }); _ = request;
  extendedRequest := this.Extend(request , params ); _ = extendedRequest;
  if IsTrue(OpEq2(*(extendedRequest ).At(MkString("wallet") ), MkUndefined() )) {
    panic(NewBadRequest(OpAdd(*this.At(MkString("id")) , MkString(" wallet is undefined, set this.walletAddress or \"address\" in params") )));
  }
  response := OpCopy(MkUndefined() ); _ = response;
  {
  ret__ := func(this *Idex) (ret_ *Variant) {
    defer func() {
      if e := recover().(*Variant); e != nil {
        ret_ = func(this *Idex) *Variant {
          // catch block:
          if IsTrue(OpIsType(e , InvalidAddress )) {
            walletAddress := *(extendedRequest ).At(MkString("wallet") ); _ = walletAddress;
            this.AssociateWallet(walletAddress );
            response = this.Call(MkString("privateGetBalances"), extendedRequest );
          } else {
            panic(e );
          }
          return nil
        }(this)
      }
    }()
    // try block:
    response = this.Call(MkString("privateGetBalances"), extendedRequest );
    return nil
  }(this)
  if ret__ != nil {
    return ret__;
   }
  }
  result := MkMap(&VarMap{
        "info":response ,
        "timestamp":MkUndefined() ,
        "datetime":MkUndefined() ,
        }); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , response.Length )); OpIncr(& i ){
    entry := *(response ).At(i ); _ = entry;
    currencyId := this.SafeString(entry , MkString("asset") ); _ = currencyId;
    code := this.SafeCurrencyCode(currencyId ); _ = code;
    account := this.Account(); _ = account;
    *(account ).At(MkString("total") )= this.SafeString(entry , MkString("quantity") );
    *(account ).At(MkString("free") )= this.SafeString(entry , MkString("availableForTrade") );
    *(account ).At(MkString("used") )= this.SafeString(entry , MkString("locked") );
    *(result ).At(code )= OpCopy(account );
  }
  return this.ParseBalance(result );
}

func (this *Idex) FetchMyTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.CheckRequiredCredentials();
  this.LoadMarkets();
  market := OpCopy(MkUndefined() ); _ = market;
  request := MkMap(&VarMap{
        "nonce":this.Uuidv1(),
        "wallet":*this.At(MkString("walletAddress")) ,
        }); _ = request;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("market") )= *(market ).At(MkString("id") );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("start") )= OpCopy(since );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  extendedRequest := this.Extend(request , params ); _ = extendedRequest;
  if IsTrue(OpEq2(*(extendedRequest ).At(MkString("wallet") ), MkUndefined() )) {
    panic(NewBadRequest(OpAdd(*this.At(MkString("id")) , MkString(" walletAddress is undefined, set this.walletAddress or \"address\" in params") )));
  }
  response := OpCopy(MkUndefined() ); _ = response;
  {
  ret__ := func(this *Idex) (ret_ *Variant) {
    defer func() {
      if e := recover().(*Variant); e != nil {
        ret_ = func(this *Idex) *Variant {
          // catch block:
          if IsTrue(OpIsType(e , InvalidAddress )) {
            walletAddress := *(extendedRequest ).At(MkString("wallet") ); _ = walletAddress;
            this.AssociateWallet(walletAddress );
            response = this.Call(MkString("privateGetFills"), extendedRequest );
          } else {
            panic(e );
          }
          return nil
        }(this)
      }
    }()
    // try block:
    response = this.Call(MkString("privateGetFills"), extendedRequest );
    return nil
  }(this)
  if ret__ != nil {
    return ret__;
   }
  }
  return this.ParseTrades(response , market , since , limit );
}

func (this *Idex) FetchOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"orderId":id }); _ = request;
  return this.FetchOrdersHelper(symbol , MkUndefined() , MkUndefined() , this.Extend(request , params ));
}

func (this *Idex) FetchOpenOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"closed":MkBool(false) }); _ = request;
  return this.FetchOrdersHelper(symbol , since , limit , this.Extend(request , params ));
}

func (this *Idex) FetchClosedOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"closed":MkBool(true) }); _ = request;
  return this.FetchOrdersHelper(symbol , since , limit , this.Extend(request , params ));
}

func (this *Idex) FetchOrdersHelper(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{
        "nonce":this.Uuidv1(),
        "wallet":*this.At(MkString("walletAddress")) ,
        }); _ = request;
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("market") )= *(market ).At(MkString("id") );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("start") )= OpCopy(since );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("privateGetOrders"), this.Extend(request , params )); _ = response;
  if IsTrue(GoIsArray(response )) {
    return this.ParseOrders(response , market , since , limit );
  } else {
    return this.ParseOrder(response , market );
  }
  return MkUndefined()
}

func (this *Idex) ParseOrderStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "active":MkString("open") ,
        "partiallyFilled":MkString("open") ,
        "rejected":MkString("canceled") ,
        "filled":MkString("closed") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Idex) ParseOrder(goArgs ...*Variant) *Variant{
  order := GoGetArg(goArgs, 0, MkUndefined()); _ = order;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.SafeInteger(order , MkString("time") ); _ = timestamp;
  fills := this.SafeValue(order , MkString("fills") , MkArray(&VarArray{})); _ = fills;
  id := this.SafeString(order , MkString("orderId") ); _ = id;
  clientOrderId := this.SafeString(order , MkString("clientOrderId") ); _ = clientOrderId;
  marketId := this.SafeString(order , MkString("market") ); _ = marketId;
  side := this.SafeString(order , MkString("side") ); _ = side;
  symbol := this.SafeSymbol(marketId , market , MkString("-") ); _ = symbol;
  trades := this.ParseTrades(fills , market ); _ = trades;
  type_ := this.SafeString(order , MkString("type") ); _ = type_;
  amount := this.SafeNumber(order , MkString("originalQuantity") ); _ = amount;
  filled := this.SafeNumber(order , MkString("executedQuantity") ); _ = filled;
  average := this.SafeNumber(order , MkString("avgExecutionPrice") ); _ = average;
  price := this.SafeNumber(order , MkString("price") ); _ = price;
  rawStatus := this.SafeString(order , MkString("status") ); _ = rawStatus;
  status := this.ParseOrderStatus(rawStatus ); _ = status;
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
            "average":average ,
            "filled":filled ,
            "remaining":MkUndefined() ,
            "status":status ,
            "fee":MkUndefined() ,
            "trades":trades ,
            }));
}

func (this *Idex) AssociateWallet(goArgs ...*Variant) *Variant{
  walletAddress := GoGetArg(goArgs, 0, MkUndefined()); _ = walletAddress;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  nonce := this.Uuidv1(); _ = nonce;
  noPrefix := this.Remove0xPrefix(walletAddress ); _ = noPrefix;
  byteArray := MkArray(&VarArray{
        this.Base16ToBinary(nonce ),
        this.Base16ToBinary(noPrefix ),
        }); _ = byteArray;
  binary := this.BinaryConcatArray(byteArray ); _ = binary;
  hash := this.Hash(binary , MkString("keccak") , MkString("hex") ); _ = hash;
  signature := this.SignMessageString(hash , *this.At(MkString("privateKey")) ); _ = signature;
  request := MkMap(&VarMap{
        "parameters":MkMap(&VarMap{
            "nonce":nonce ,
            "wallet":walletAddress ,
            }),
        "signature":signature ,
        }); _ = request;
  result := this.Call(MkString("privatePostWallets"), request ); _ = result;
  return result ;
}

func (this *Idex) CreateOrder(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 2, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 3, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 4, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 5, MkMap(&VarMap{})); _ = params;
  this.CheckRequiredCredentials();
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  nonce := this.Uuidv1(); _ = nonce;
  typeEnum := OpCopy(MkUndefined() ); _ = typeEnum;
  stopLossTypeEnums := MkMap(&VarMap{
        "stopLoss":MkInteger(3) ,
        "stopLossLimit":MkInteger(4) ,
        "takeProfit":MkInteger(5) ,
        "takeProfitLimit":MkInteger(6) ,
        }); _ = stopLossTypeEnums;
  stopPriceString := OpCopy(MkUndefined() ); _ = stopPriceString;
  if IsTrue(OpOr(OpEq2(type_ , MkString("stopLossLimit") ), OpOr(OpEq2(type_ , MkString("takeProfitLimit") ), OpHasMember(MkString("stopPrice") , params )))) {
    if IsTrue(OpNot(OpHasMember(MkString("stopPrice") , params ))) {
      panic(NewBadRequest(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" stopPrice is a required parameter for ") , OpAdd(type_ , MkString("orders") )))));
    }
    stopPriceString = this.PriceToPrecision(symbol , *(params ).At(MkString("stopPrice") ));
  }
  limitTypeEnums := MkMap(&VarMap{
        "limit":MkInteger(1) ,
        "limitMaker":MkInteger(2) ,
        }); _ = limitTypeEnums;
  priceString := OpCopy(MkUndefined() ); _ = priceString;
  typeLower := type_.ToLowerCase(); _ = typeLower;
  limitOrder := OpGt(typeLower.IndexOf(MkString("limit") ), OpNeg(MkInteger(1) )); _ = limitOrder;
  if IsTrue(OpHasMember(type_ , limitTypeEnums )) {
    typeEnum = *(limitTypeEnums ).At(type_ );
    priceString = this.PriceToPrecision(symbol , price );
  } else {
    if IsTrue(OpHasMember(type_ , stopLossTypeEnums )) {
      typeEnum = *(stopLossTypeEnums ).At(type_ );
      priceString = this.PriceToPrecision(symbol , price );
    } else {
      if IsTrue(OpEq2(type_ , MkString("market") )) {
        typeEnum = MkInteger(0) ;
      } else {
        panic(NewBadRequest(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , OpAdd(type_ , MkString(" is not a valid order type") )))));
      }
    }
  }
  amountEnum := MkInteger(0) ; _ = amountEnum;
  if IsTrue(OpHasMember(MkString("quoteOrderQuantity") , params )) {
    if IsTrue(OpNotEq2(type_ , MkString("market") )) {
      panic(NewNotSupported(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" quoteOrderQuantity is not supported for ") , OpAdd(type_ , MkString(" orders, only supported for market orders") )))));
    }
    amountEnum = MkInteger(1) ;
    amount = this.SafeNumber(params , MkString("quoteOrderQuantity") );
  }
  sideEnum := OpTrinary(OpEq2(side , MkString("buy") ), MkInteger(0) , MkInteger(1) ); _ = sideEnum;
  walletBytes := this.Remove0xPrefix(*this.At(MkString("walletAddress")) ); _ = walletBytes;
  network := this.SafeString(*this.At(MkString("options")) , MkString("network") , MkString("ETH") ); _ = network;
  orderVersion := OpTrinary(OpEq2(network , MkString("ETH") ), MkInteger(1) , MkInteger(2) ); _ = orderVersion;
  amountString := this.AmountToPrecision(symbol , amount ); _ = amountString;
  timeInForceEnums := MkMap(&VarMap{
        "gtc":MkInteger(0) ,
        "ioc":MkInteger(2) ,
        "fok":MkInteger(3) ,
        }); _ = timeInForceEnums;
  defaultTimeInForce := this.SafeString(*this.At(MkString("options")) , MkString("defaultTimeInForce") , MkString("gtc") ); _ = defaultTimeInForce;
  timeInForce := this.SafeString(params , MkString("timeInForce") , defaultTimeInForce ); _ = timeInForce;
  timeInForceEnum := OpCopy(MkUndefined() ); _ = timeInForceEnum;
  if IsTrue(OpHasMember(timeInForce , timeInForceEnums )) {
    timeInForceEnum = *(timeInForceEnums ).At(timeInForce );
  } else {
    allOptions := GoGetKeys(timeInForceEnums ); _ = allOptions;
    asString := allOptions.Join(MkString(", ") ); _ = asString;
    panic(NewBadRequest(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , OpAdd(timeInForce , OpAdd(MkString(" is not a valid timeInForce, please choose one of ") , asString ))))));
  }
  selfTradePreventionEnums := MkMap(&VarMap{
        "dc":MkInteger(0) ,
        "co":MkInteger(1) ,
        "cn":MkInteger(2) ,
        "cb":MkInteger(3) ,
        }); _ = selfTradePreventionEnums;
  defaultSelfTradePrevention := this.SafeString(*this.At(MkString("options")) , MkString("defaultSelfTradePrevention") , MkString("cn") ); _ = defaultSelfTradePrevention;
  selfTradePrevention := this.SafeString(params , MkString("selfTradePrevention") , defaultSelfTradePrevention ); _ = selfTradePrevention;
  selfTradePreventionEnum := OpCopy(MkUndefined() ); _ = selfTradePreventionEnum;
  if IsTrue(OpHasMember(selfTradePrevention , selfTradePreventionEnums )) {
    selfTradePreventionEnum = *(selfTradePreventionEnums ).At(selfTradePrevention );
  } else {
    allOptions := GoGetKeys(selfTradePreventionEnums ); _ = allOptions;
    asString := allOptions.Join(MkString(", ") ); _ = asString;
    panic(NewBadRequest(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , OpAdd(selfTradePrevention , OpAdd(MkString(" is not a valid selfTradePrevention, please choose one of ") , asString ))))));
  }
  byteArray := MkArray(&VarArray{
        this.NumberToBE(orderVersion , MkInteger(1) ),
        this.Base16ToBinary(nonce ),
        this.Base16ToBinary(walletBytes ),
        this.StringToBinary(this.Encode(*(market ).At(MkString("id") ))),
        this.NumberToBE(typeEnum , MkInteger(1) ),
        this.NumberToBE(sideEnum , MkInteger(1) ),
        this.StringToBinary(this.Encode(amountString )),
        this.NumberToBE(amountEnum , MkInteger(1) ),
        }); _ = byteArray;
  if IsTrue(limitOrder ) {
    encodedPrice := this.StringToBinary(this.Encode(priceString )); _ = encodedPrice;
    byteArray.Push(encodedPrice );
  }
  if IsTrue(OpHasMember(type_ , stopLossTypeEnums )) {
    encodedPrice := this.StringToBinary(this.Encode(OpOr(stopPriceString , priceString ))); _ = encodedPrice;
    byteArray.Push(encodedPrice );
  }
  clientOrderId := this.SafeString(params , MkString("clientOrderId") ); _ = clientOrderId;
  if IsTrue(OpNotEq2(clientOrderId , MkUndefined() )) {
    byteArray.Push(this.StringToBinary(this.Encode(clientOrderId )));
  }
  after := MkArray(&VarArray{
        this.NumberToBE(timeInForceEnum , MkInteger(1) ),
        this.NumberToBE(selfTradePreventionEnum , MkInteger(1) ),
        this.NumberToBE(MkInteger(0) , MkInteger(8) ),
        }); _ = after;
  allBytes := this.ArrayConcat(byteArray , after ); _ = allBytes;
  binary := this.BinaryConcatArray(allBytes ); _ = binary;
  hash := this.Hash(binary , MkString("keccak") , MkString("hex") ); _ = hash;
  signature := this.SignMessageString(hash , *this.At(MkString("privateKey")) ); _ = signature;
  request := MkMap(&VarMap{
        "parameters":MkMap(&VarMap{
            "nonce":nonce ,
            "market":*(market ).At(MkString("id") ),
            "side":side ,
            "type":type_ ,
            "wallet":*this.At(MkString("walletAddress")) ,
            "timeInForce":timeInForce ,
            "selfTradePrevention":selfTradePrevention ,
            }),
        "signature":signature ,
        }); _ = request;
  if IsTrue(limitOrder ) {
    *(*(request ).At(MkString("parameters") )).At(MkString("price") )= OpCopy(priceString );
  }
  if IsTrue(OpHasMember(type_ , stopLossTypeEnums )) {
    *(*(request ).At(MkString("parameters") )).At(MkString("stopPrice") )= OpOr(stopPriceString , priceString );
  }
  if IsTrue(OpEq2(amountEnum , MkInteger(0) )) {
    *(*(request ).At(MkString("parameters") )).At(MkString("quantity") )= OpCopy(amountString );
  } else {
    *(*(request ).At(MkString("parameters") )).At(MkString("quoteOrderQuantity") )= OpCopy(amountString );
  }
  if IsTrue(OpNotEq2(clientOrderId , MkUndefined() )) {
    *(*(request ).At(MkString("parameters") )).At(MkString("clientOrderId") )= OpCopy(clientOrderId );
  }
  response := this.Call(MkString("privatePostOrders"), request ); _ = response;
  return this.ParseOrder(response , market );
}

func (this *Idex) Withdraw(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  amount := GoGetArg(goArgs, 1, MkUndefined()); _ = amount;
  address := GoGetArg(goArgs, 2, MkUndefined()); _ = address;
  tag := GoGetArg(goArgs, 3, MkUndefined() ); _ = tag;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.CheckRequiredCredentials();
  this.LoadMarkets();
  nonce := this.Uuidv1(); _ = nonce;
  amountString := this.CurrencyToPrecision(code , amount ); _ = amountString;
  currency := this.Currency(code ); _ = currency;
  walletBytes := this.Remove0xPrefix(*this.At(MkString("walletAddress")) ); _ = walletBytes;
  byteArray := MkArray(&VarArray{
        this.Base16ToBinary(nonce ),
        this.Base16ToBinary(walletBytes ),
        this.StringToBinary(this.Encode(*(currency ).At(MkString("id") ))),
        this.StringToBinary(this.Encode(amountString )),
        this.NumberToBE(MkInteger(1) , MkInteger(1) ),
        }); _ = byteArray;
  binary := this.BinaryConcatArray(byteArray ); _ = binary;
  hash := this.Hash(binary , MkString("keccak") , MkString("hex") ); _ = hash;
  signature := this.SignMessageString(hash , *this.At(MkString("privateKey")) ); _ = signature;
  request := MkMap(&VarMap{
        "parameters":MkMap(&VarMap{
            "nonce":nonce ,
            "wallet":address ,
            "asset":*(currency ).At(MkString("id") ),
            "quantity":amountString ,
            }),
        "signature":signature ,
        }); _ = request;
  response := this.Call(MkString("privatePostWithdrawals"), request ); _ = response;
  id := this.SafeString(response , MkString("withdrawalId") ); _ = id;
  return MkMap(&VarMap{
        "info":response ,
        "id":id ,
        });
}

func (this *Idex) CancelOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.CheckRequiredCredentials();
  this.LoadMarkets();
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
  }
  nonce := this.Uuidv1(); _ = nonce;
  walletBytes := this.Remove0xPrefix(*this.At(MkString("walletAddress")) ); _ = walletBytes;
  byteArray := MkArray(&VarArray{
        this.Base16ToBinary(nonce ),
        this.Base16ToBinary(walletBytes ),
        this.StringToBinary(this.Encode(id )),
        }); _ = byteArray;
  binary := this.BinaryConcatArray(byteArray ); _ = binary;
  hash := this.Hash(binary , MkString("keccak") , MkString("hex") ); _ = hash;
  signature := this.SignMessageString(hash , *this.At(MkString("privateKey")) ); _ = signature;
  request := MkMap(&VarMap{
        "parameters":MkMap(&VarMap{
            "nonce":nonce ,
            "wallet":*this.At(MkString("walletAddress")) ,
            "orderId":id ,
            }),
        "signature":signature ,
        }); _ = request;
  response := this.Call(MkString("privateDeleteOrders"), this.Extend(request , params )); _ = response;
  canceledOrder := this.SafeValue(response , MkInteger(0) ); _ = canceledOrder;
  return this.ParseOrder(canceledOrder , market );
}

func (this *Idex) HandleErrors(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  reason := GoGetArg(goArgs, 1, MkUndefined()); _ = reason;
  url := GoGetArg(goArgs, 2, MkUndefined()); _ = url;
  method := GoGetArg(goArgs, 3, MkUndefined()); _ = method;
  headers := GoGetArg(goArgs, 4, MkUndefined()); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined()); _ = body;
  response := GoGetArg(goArgs, 6, MkUndefined()); _ = response;
  requestHeaders := GoGetArg(goArgs, 7, MkUndefined()); _ = requestHeaders;
  requestBody := GoGetArg(goArgs, 8, MkUndefined()); _ = requestBody;
  errorCode := this.SafeString(response , MkString("code") ); _ = errorCode;
  message := this.SafeString(response , MkString("message") ); _ = message;
  if IsTrue(OpHasMember(errorCode , *this.At(MkString("exceptions")) )) {
    Exception := *(*this.At(MkString("exceptions")) ).At(errorCode ); _ = Exception;
    panic(NewException(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , message ))));
  }
  if IsTrue(OpNotEq2(errorCode , MkUndefined() )) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , message ))));
  }
  return MkUndefined()
}

func (this *Idex) FetchDeposits(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  params = this.Extend(MkMap(&VarMap{"method":MkString("privateGetDeposits") }), params );
  return this.FetchTransactionsHelper(code , since , limit , params );
}

func (this *Idex) FetchWithdrawals(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  params = this.Extend(MkMap(&VarMap{"method":MkString("privateGetWithdrawals") }), params );
  return this.FetchTransactionsHelper(code , since , limit , params );
}

func (this *Idex) FetchTransactionsHelper(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  nonce := this.Uuidv1(); _ = nonce;
  request := MkMap(&VarMap{
        "nonce":nonce ,
        "wallet":*this.At(MkString("walletAddress")) ,
        }); _ = request;
  currency := OpCopy(MkUndefined() ); _ = currency;
  if IsTrue(OpNotEq2(code , MkUndefined() )) {
    currency = this.Currency(code );
    *(request ).At(MkString("asset") )= *(currency ).At(MkString("id") );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("start") )= OpCopy(since );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  method := *(params ).At(MkString("method") ); _ = method;
  params = this.Omit(params , MkString("method") );
  response := this.Call(method , this.Extend(request , params )); _ = response;
  return this.ParseTransactions(response , currency , since , limit );
}

func (this *Idex) ParseTransactionStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{"mined":MkString("ok") }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Idex) ParseTransaction(goArgs ...*Variant) *Variant{
  transaction := GoGetArg(goArgs, 0, MkUndefined()); _ = transaction;
  currency := GoGetArg(goArgs, 1, MkUndefined() ); _ = currency;
  type_ := OpCopy(MkUndefined() ); _ = type_;
  if IsTrue(OpHasMember(MkString("depositId") , transaction )) {
    type_ = MkString("deposit") ;
  } else {
    if IsTrue(OpHasMember(MkString("withdrawalId") , transaction )) {
      type_ = MkString("withdrawal") ;
    }
  }
  id := this.SafeString2(transaction , MkString("depositId") , MkString("withdrawId") ); _ = id;
  code := this.SafeCurrencyCode(this.SafeString(transaction , MkString("asset") ), currency ); _ = code;
  amount := this.SafeNumber(transaction , MkString("quantity") ); _ = amount;
  txid := this.SafeString(transaction , MkString("txId") ); _ = txid;
  timestamp := this.SafeInteger(transaction , MkString("txTime") ); _ = timestamp;
  fee := OpCopy(MkUndefined() ); _ = fee;
  if IsTrue(OpHasMember(MkString("fee") , transaction )) {
    fee = MkMap(&VarMap{
        "cost":this.SafeNumber(transaction , MkString("fee") ),
        "currency":MkString("ETH") ,
        });
  }
  rawStatus := this.SafeString(transaction , MkString("txStatus") ); _ = rawStatus;
  status := this.ParseTransactionStatus(rawStatus ); _ = status;
  updated := this.SafeInteger(transaction , MkString("confirmationTime") ); _ = updated;
  return MkMap(&VarMap{
        "info":transaction ,
        "id":id ,
        "txid":txid ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "address":MkUndefined() ,
        "tag":MkUndefined() ,
        "type":type_ ,
        "amount":amount ,
        "currency":code ,
        "status":status ,
        "updated":updated ,
        "fee":fee ,
        });
}

func (this *Idex) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("public") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  network := this.SafeString(*this.At(MkString("options")) , MkString("network") , MkString("ETH") ); _ = network;
  version := this.SafeString(*this.At(MkString("options")) , MkString("version") , MkString("v1") ); _ = version;
  url := OpAdd(*(*(*this.At(MkString("urls")) ).At(MkString("api") )).At(network ), OpAdd(MkString("/") , OpAdd(version , OpAdd(MkString("/") , path )))); _ = url;
  keys := GoGetKeys(params ); _ = keys;
  length := OpCopy(keys.Length ); _ = length;
  query := OpCopy(MkUndefined() ); _ = query;
  if IsTrue(OpGt(length , MkInteger(0) )) {
    if IsTrue(OpEq2(method , MkString("GET") )) {
      query = this.Urlencode(params );
      url = OpAdd(url , OpAdd(MkString("?") , query ));
    } else {
      body = this.Json(params );
    }
  }
  headers = MkMap(&VarMap{"Content-Type":MkString("application/json") });
  if IsTrue(OpNotEq2(*this.At(MkString("apiKey")) , MkUndefined() )) {
    *(headers ).At(MkString("IDEX-API-Key") )= OpCopy(*this.At(MkString("apiKey")) );
  }
  if IsTrue(OpEq2(api , MkString("private") )) {
    payload := OpCopy(MkUndefined() ); _ = payload;
    if IsTrue(OpEq2(method , MkString("GET") )) {
      payload = OpCopy(query );
    } else {
      payload = OpCopy(body );
    }
    *(headers ).At(MkString("IDEX-HMAC-Signature") )= this.Hmac(this.Encode(payload ), this.Encode(*this.At(MkString("secret")) ), MkString("sha256") , MkString("hex") );
  }
  return MkMap(&VarMap{
        "url":url ,
        "method":method ,
        "body":body ,
        "headers":headers ,
        });
}

