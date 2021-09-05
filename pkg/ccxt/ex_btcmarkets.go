package ccxt

import (
)

type Btcmarkets struct {
	*ExchangeBase
}

var _ Exchange = (*Btcmarkets)(nil)

func init() {
	exchange := &Btcmarkets{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Btcmarkets) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("btcmarkets") ,
            "name":MkString("BTC Markets") ,
            "countries":MkArray(&VarArray{
                MkString("AU") ,
                }),
            "rateLimit":MkInteger(1000) ,
            "version":MkString("v3") ,
            "has":MkMap(&VarMap{
                "cancelOrder":MkBool(true) ,
                "cancelOrders":MkBool(true) ,
                "CORS":MkBool(false) ,
                "createOrder":MkBool(true) ,
                "fetchBalance":MkBool(true) ,
                "fetchClosedOrders":MkString("emulated") ,
                "fetchDeposits":MkBool(true) ,
                "fetchMarkets":MkBool(true) ,
                "fetchMyTrades":MkBool(true) ,
                "fetchOHLCV":MkBool(true) ,
                "fetchOpenOrders":MkBool(true) ,
                "fetchOrder":MkBool(true) ,
                "fetchOrderBook":MkBool(true) ,
                "fetchOrders":MkBool(true) ,
                "fetchTicker":MkBool(true) ,
                "fetchTime":MkBool(true) ,
                "fetchTrades":MkBool(true) ,
                "fetchTransactions":MkBool(true) ,
                "fetchWithdrawals":MkBool(true) ,
                }),
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/51840849/89731817-b3fb8480-da52-11ea-817f-783b08aaf32b.jpg") ,
                "api":MkMap(&VarMap{
                    "public":MkString("https://api.btcmarkets.net") ,
                    "private":MkString("https://api.btcmarkets.net") ,
                    }),
                "www":MkString("https://btcmarkets.net") ,
                "doc":MkArray(&VarArray{
                    MkString("https://api.btcmarkets.net/doc/v3") ,
                    MkString("https://github.com/BTCMarkets/API") ,
                    }),
                }),
            "api":MkMap(&VarMap{
                "public":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("markets") ,
                        MkString("markets/{marketId}/ticker") ,
                        MkString("markets/{marketId}/trades") ,
                        MkString("markets/{marketId}/orderbook") ,
                        MkString("markets/{marketId}/candles") ,
                        MkString("markets/tickers") ,
                        MkString("markets/orderbooks") ,
                        MkString("time") ,
                        })}),
                "private":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("orders") ,
                        MkString("orders/{id}") ,
                        MkString("batchorders/{ids}") ,
                        MkString("trades") ,
                        MkString("trades/{id}") ,
                        MkString("withdrawals") ,
                        MkString("withdrawals/{id}") ,
                        MkString("deposits") ,
                        MkString("deposits/{id}") ,
                        MkString("transfers") ,
                        MkString("transfers/{id}") ,
                        MkString("addresses") ,
                        MkString("withdrawal-fees") ,
                        MkString("assets") ,
                        MkString("accounts/me/trading-fees") ,
                        MkString("accounts/me/withdrawal-limits") ,
                        MkString("accounts/me/balances") ,
                        MkString("accounts/me/transactions") ,
                        MkString("reports/{id}") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("orders") ,
                        MkString("batchorders") ,
                        MkString("withdrawals") ,
                        MkString("reports") ,
                        }),
                    "delete":MkArray(&VarArray{
                        MkString("orders") ,
                        MkString("orders/{id}") ,
                        MkString("batchorders/{ids}") ,
                        }),
                    "put":MkArray(&VarArray{
                        MkString("orders/{id}") ,
                        }),
                    }),
                }),
            "timeframes":MkMap(&VarMap{
                "1m":MkString("1m") ,
                "1h":MkString("1h") ,
                "1d":MkString("1d") ,
                }),
            "exceptions":MkMap(&VarMap{
                "3":InvalidOrder ,
                "6":DDoSProtection ,
                "InsufficientFund":InsufficientFunds ,
                "InvalidPrice":InvalidOrder ,
                "InvalidAmount":InvalidOrder ,
                "MissingArgument":InvalidOrder ,
                "OrderAlreadyCancelled":InvalidOrder ,
                "OrderNotFound":OrderNotFound ,
                "OrderStatusIsFinal":InvalidOrder ,
                "InvalidPaginationParameter":BadRequest ,
                }),
            "fees":MkMap(&VarMap{
                "percentage":MkBool(true) ,
                "tierBased":MkBool(true) ,
                "maker":this.ParseNumber(MkString("-0.0005") ),
                "taker":this.ParseNumber(MkString("0.0020") ),
                }),
            "options":MkMap(&VarMap{"fees":MkMap(&VarMap{"AUD":MkMap(&VarMap{
                        "maker":OpDiv(MkNumber(0.85) , MkInteger(100) ),
                        "taker":OpDiv(MkNumber(0.85) , MkInteger(100) ),
                        })})}),
            }));
}

func (this *Btcmarkets) FetchTransactionsWithMethod(goArgs ...*Variant) *Variant{
  method := GoGetArg(goArgs, 0, MkUndefined()); _ = method;
  code := GoGetArg(goArgs, 1, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("after") )= OpCopy(since );
  }
  currency := OpCopy(MkUndefined() ); _ = currency;
  if IsTrue(OpNotEq2(code , MkUndefined() )) {
    currency = this.Currency(code );
  }
  response := this.Call(method , this.Extend(request , params )); _ = response;
  return this.ParseTransactions(response , currency , since , limit );
}

func (this *Btcmarkets) FetchTransactions(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  return this.FetchTransactionsWithMethod(MkString("privateGetTransfers") , code , since , limit , params );
}

func (this *Btcmarkets) FetchDeposits(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  return this.FetchTransactionsWithMethod(MkString("privateGetDeposits") , code , since , limit , params );
}

func (this *Btcmarkets) FetchWithdrawals(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  return this.FetchTransactionsWithMethod(MkString("privateGetWithdrawals") , code , since , limit , params );
}

func (this *Btcmarkets) ParseTransactionStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{"Complete":MkString("ok") }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Btcmarkets) ParseTransactionType(goArgs ...*Variant) *Variant{
  type_ := GoGetArg(goArgs, 0, MkUndefined()); _ = type_;
  statuses := MkMap(&VarMap{
        "Withdraw":MkString("withdrawal") ,
        "Deposit":MkString("deposit") ,
        }); _ = statuses;
  return this.SafeString(statuses , type_ , type_ );
}

func (this *Btcmarkets) ParseTransaction(goArgs ...*Variant) *Variant{
  transaction := GoGetArg(goArgs, 0, MkUndefined()); _ = transaction;
  currency := GoGetArg(goArgs, 1, MkUndefined() ); _ = currency;
  timestamp := this.Parse8601(this.SafeString(transaction , MkString("creationTime") )); _ = timestamp;
  lastUpdate := this.Parse8601(this.SafeString(transaction , MkString("lastUpdate") )); _ = lastUpdate;
  type_ := this.ParseTransactionType(this.SafeStringLower(transaction , MkString("type") )); _ = type_;
  if IsTrue(OpEq2(type_ , MkString("withdraw") )) {
    type_ = MkString("withdrawal") ;
  }
  cryptoPaymentDetail := this.SafeValue(transaction , MkString("paymentDetail") , MkMap(&VarMap{})); _ = cryptoPaymentDetail;
  txid := this.SafeString(cryptoPaymentDetail , MkString("txId") ); _ = txid;
  address := this.SafeString(cryptoPaymentDetail , MkString("address") ); _ = address;
  tag := OpCopy(MkUndefined() ); _ = tag;
  if IsTrue(OpNotEq2(address , MkUndefined() )) {
    addressParts := address.Split(MkString("?dt=") ); _ = addressParts;
    numParts := OpCopy(addressParts.Length ); _ = numParts;
    if IsTrue(OpGt(numParts , MkInteger(1) )) {
      address = *(addressParts ).At(MkInteger(0) );
      tag = *(addressParts ).At(MkInteger(1) );
    }
  }
  addressTo := OpCopy(address ); _ = addressTo;
  tagTo := OpCopy(tag ); _ = tagTo;
  addressFrom := OpCopy(MkUndefined() ); _ = addressFrom;
  tagFrom := OpCopy(MkUndefined() ); _ = tagFrom;
  fee := this.SafeNumber(transaction , MkString("fee") ); _ = fee;
  status := this.ParseTransactionStatus(this.SafeString(transaction , MkString("status") )); _ = status;
  currencyId := this.SafeString(transaction , MkString("assetName") ); _ = currencyId;
  code := this.SafeCurrencyCode(currencyId ); _ = code;
  amount := this.SafeNumber(transaction , MkString("amount") ); _ = amount;
  if IsTrue(fee ) {
    OpSubAssign(& amount , fee );
  }
  return MkMap(&VarMap{
        "id":this.SafeString(transaction , MkString("id") ),
        "txid":txid ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "address":address ,
        "addressTo":addressTo ,
        "addressFrom":addressFrom ,
        "tag":tag ,
        "tagTo":tagTo ,
        "tagFrom":tagFrom ,
        "type":type_ ,
        "amount":amount ,
        "currency":code ,
        "status":status ,
        "updated":lastUpdate ,
        "fee":MkMap(&VarMap{
            "currency":code ,
            "cost":fee ,
            }),
        "info":transaction ,
        });
}

func (this *Btcmarkets) FetchMarkets(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetMarkets"), params ); _ = response;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , response.Length )); OpIncr(& i ){
    market := *(response ).At(i ); _ = market;
    baseId := this.SafeString(market , MkString("baseAssetName") ); _ = baseId;
    quoteId := this.SafeString(market , MkString("quoteAssetName") ); _ = quoteId;
    id := this.SafeString(market , MkString("marketId") ); _ = id;
    base := this.SafeCurrencyCode(baseId ); _ = base;
    quote := this.SafeCurrencyCode(quoteId ); _ = quote;
    symbol := OpAdd(base , OpAdd(MkString("/") , quote )); _ = symbol;
    fees := this.SafeValue(this.SafeValue(*this.At(MkString("options")) , MkString("fees") , MkMap(&VarMap{})), quote , *this.At(MkString("fees")) ); _ = fees;
    pricePrecision := this.SafeInteger(market , MkString("priceDecimals") ); _ = pricePrecision;
    amountPrecision := this.SafeInteger(market , MkString("amountDecimals") ); _ = amountPrecision;
    minAmount := this.SafeNumber(market , MkString("minOrderAmount") ); _ = minAmount;
    maxAmount := this.SafeNumber(market , MkString("maxOrderAmount") ); _ = maxAmount;
    minPrice := OpCopy(MkUndefined() ); _ = minPrice;
    if IsTrue(OpEq2(quote , MkString("AUD") )) {
      minPrice = Math.Pow(MkInteger(10) , OpNeg(pricePrecision ));
    }
    precision := MkMap(&VarMap{
          "amount":amountPrecision ,
          "price":pricePrecision ,
          }); _ = precision;
    limits := MkMap(&VarMap{
          "amount":MkMap(&VarMap{
              "min":minAmount ,
              "max":maxAmount ,
              }),
          "price":MkMap(&VarMap{
              "min":minPrice ,
              "max":MkUndefined() ,
              }),
          "cost":MkMap(&VarMap{
              "min":MkUndefined() ,
              "max":MkUndefined() ,
              }),
          }); _ = limits;
    result.Push(MkMap(&VarMap{
            "info":market ,
            "id":id ,
            "symbol":symbol ,
            "base":base ,
            "quote":quote ,
            "baseId":baseId ,
            "quoteId":quoteId ,
            "active":MkUndefined() ,
            "maker":*(fees ).At(MkString("maker") ),
            "taker":*(fees ).At(MkString("taker") ),
            "limits":limits ,
            "precision":precision ,
            }));
  }
  return result ;
}

func (this *Btcmarkets) FetchTime(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetTime"), params ); _ = response;
  return this.Parse8601(this.SafeString(response , MkString("timestamp") ));
}

func (this *Btcmarkets) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("privateGetAccountsMeBalances"), params ); _ = response;
  result := MkMap(&VarMap{"info":response }); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , response.Length )); OpIncr(& i ){
    balance := *(response ).At(i ); _ = balance;
    currencyId := this.SafeString(balance , MkString("assetName") ); _ = currencyId;
    code := this.SafeCurrencyCode(currencyId ); _ = code;
    account := this.Account(); _ = account;
    *(account ).At(MkString("used") )= this.SafeString(balance , MkString("locked") );
    *(account ).At(MkString("total") )= this.SafeString(balance , MkString("balance") );
    *(result ).At(code )= OpCopy(account );
  }
  return this.ParseBalance(result );
}

func (this *Btcmarkets) ParseOHLCV(goArgs ...*Variant) *Variant{
  ohlcv := GoGetArg(goArgs, 0, MkUndefined()); _ = ohlcv;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  return MkArray(&VarArray{
        this.Parse8601(this.SafeString(ohlcv , MkInteger(0) )),
        this.SafeNumber(ohlcv , MkInteger(1) ),
        this.SafeNumber(ohlcv , MkInteger(2) ),
        this.SafeNumber(ohlcv , MkInteger(3) ),
        this.SafeNumber(ohlcv , MkInteger(4) ),
        this.SafeNumber(ohlcv , MkInteger(5) ),
        });
}

func (this *Btcmarkets) FetchOHLCV(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  timeframe := GoGetArg(goArgs, 1, MkString("1m") ); _ = timeframe;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "marketId":*(market ).At(MkString("id") ),
        "timeWindow":*(*this.At(MkString("timeframes")) ).At(timeframe ),
        }); _ = request;
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("from") )= this.Iso8601(since );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("publicGetMarketsMarketIdCandles"), this.Extend(request , params )); _ = response;
  return this.ParseOHLCVs(response , market , timeframe , since , limit );
}

func (this *Btcmarkets) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"marketId":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("publicGetMarketsMarketIdOrderbook"), this.Extend(request , params )); _ = response;
  timestamp := this.SafeIntegerProduct(response , MkString("snapshotId") , MkNumber(0.001) ); _ = timestamp;
  orderbook := this.ParseOrderBook(response , symbol , timestamp ); _ = orderbook;
  *(orderbook ).At(MkString("nonce") )= this.SafeInteger(response , MkString("snapshotId") );
  return orderbook ;
}

func (this *Btcmarkets) ParseTicker(goArgs ...*Variant) *Variant{
  ticker := GoGetArg(goArgs, 0, MkUndefined()); _ = ticker;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  marketId := this.SafeString(ticker , MkString("marketId") ); _ = marketId;
  if IsTrue(OpNotEq2(marketId , MkUndefined() )) {
    if IsTrue(OpHasMember(marketId , *this.At(MkString("markets_by_id")) )) {
      market = *(*this.At(MkString("markets_by_id")) ).At(marketId );
    } else {
      baseId := MkUndefined(); _ = baseId;
      quoteId := MkUndefined(); _ = quoteId;
      ArrayUnpack(marketId.Split(MkString("-") ), &baseId, &quoteId);
      base := this.SafeCurrencyCode(baseId ); _ = base;
      quote := this.SafeCurrencyCode(quoteId ); _ = quote;
      symbol = OpAdd(base , OpAdd(MkString("/") , quote ));
    }
  }
  if IsTrue(OpAnd(OpEq2(symbol , MkUndefined() ), OpNotEq2(market , MkUndefined() ))) {
    symbol = *(market ).At(MkString("symbol") );
  }
  timestamp := this.Parse8601(this.SafeString(ticker , MkString("timestamp") )); _ = timestamp;
  last := this.SafeNumber(ticker , MkString("lastPrice") ); _ = last;
  baseVolume := this.SafeNumber(ticker , MkString("volume24h") ); _ = baseVolume;
  quoteVolume := this.SafeNumber(ticker , MkString("volumeQte24h") ); _ = quoteVolume;
  vwap := this.Vwap(baseVolume , quoteVolume ); _ = vwap;
  change := this.SafeNumber(ticker , MkString("price24h") ); _ = change;
  percentage := this.SafeNumber(ticker , MkString("pricePct24h") ); _ = percentage;
  return MkMap(&VarMap{
        "symbol":symbol ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "high":this.SafeNumber(ticker , MkString("high24h") ),
        "low":this.SafeNumber(ticker , MkString("low") ),
        "bid":this.SafeNumber(ticker , MkString("bestBid") ),
        "bidVolume":MkUndefined() ,
        "ask":this.SafeNumber(ticker , MkString("bestAsk") ),
        "askVolume":MkUndefined() ,
        "vwap":vwap ,
        "open":MkUndefined() ,
        "close":last ,
        "last":last ,
        "previousClose":MkUndefined() ,
        "change":change ,
        "percentage":percentage ,
        "average":MkUndefined() ,
        "baseVolume":baseVolume ,
        "quoteVolume":quoteVolume ,
        "info":ticker ,
        });
}

func (this *Btcmarkets) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"marketId":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("publicGetMarketsMarketIdTicker"), this.Extend(request , params )); _ = response;
  return this.ParseTicker(response , market );
}

func (this *Btcmarkets) FetchTicker2(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"id":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("publicGetMarketIdTick"), this.Extend(request , params )); _ = response;
  return this.ParseTicker(response , market );
}

func (this *Btcmarkets) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.Parse8601(this.SafeString(trade , MkString("timestamp") )); _ = timestamp;
  marketId := this.SafeString(trade , MkString("marketId") ); _ = marketId;
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  base := OpCopy(MkUndefined() ); _ = base;
  quote := OpCopy(MkUndefined() ); _ = quote;
  if IsTrue(OpNotEq2(marketId , MkUndefined() )) {
    if IsTrue(OpHasMember(marketId , *this.At(MkString("markets_by_id")) )) {
      market = *(*this.At(MkString("markets_by_id")) ).At(marketId );
    } else {
      baseId := MkUndefined(); _ = baseId;
      quoteId := MkUndefined(); _ = quoteId;
      ArrayUnpack(marketId.Split(MkString("-") ), &baseId, &quoteId);
      base = this.SafeCurrencyCode(baseId );
      quote = this.SafeCurrencyCode(quoteId );
      symbol = OpAdd(base , OpAdd(MkString("/") , quote ));
    }
  }
  if IsTrue(OpAnd(OpEq2(symbol , MkUndefined() ), OpNotEq2(market , MkUndefined() ))) {
    symbol = *(market ).At(MkString("symbol") );
    base = *(market ).At(MkString("base") );
    quote = *(market ).At(MkString("quote") );
  }
  feeCurrencyCode := OpCopy(MkUndefined() ); _ = feeCurrencyCode;
  if IsTrue(OpEq2(quote , MkString("AUD") )) {
    feeCurrencyCode = OpCopy(quote );
  } else {
    feeCurrencyCode = OpCopy(base );
  }
  side := this.SafeString(trade , MkString("side") ); _ = side;
  if IsTrue(OpEq2(side , MkString("Bid") )) {
    side = MkString("buy") ;
  } else {
    if IsTrue(OpEq2(side , MkString("Ask") )) {
      side = MkString("sell") ;
    }
  }
  id := this.SafeString(trade , MkString("id") ); _ = id;
  priceString := this.SafeString(trade , MkString("price") ); _ = priceString;
  amountString := this.SafeString(trade , MkString("amount") ); _ = amountString;
  price := this.ParseNumber(priceString ); _ = price;
  amount := this.ParseNumber(amountString ); _ = amount;
  cost := this.ParseNumber(Precise.StringMul(priceString , amountString )); _ = cost;
  orderId := this.SafeString(trade , MkString("orderId") ); _ = orderId;
  fee := OpCopy(MkUndefined() ); _ = fee;
  feeCost := this.SafeNumber(trade , MkString("fee") ); _ = feeCost;
  if IsTrue(OpNotEq2(feeCost , MkUndefined() )) {
    fee = MkMap(&VarMap{
        "cost":feeCost ,
        "currency":feeCurrencyCode ,
        });
  }
  takerOrMaker := this.SafeStringLower(trade , MkString("liquidityType") ); _ = takerOrMaker;
  return MkMap(&VarMap{
        "info":trade ,
        "id":id ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "order":orderId ,
        "symbol":symbol ,
        "type":MkUndefined() ,
        "side":side ,
        "price":price ,
        "amount":amount ,
        "cost":cost ,
        "takerOrMaker":takerOrMaker ,
        "fee":fee ,
        });
}

func (this *Btcmarkets) FetchTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"marketId":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("publicGetMarketsMarketIdTrades"), this.Extend(request , params )); _ = response;
  return this.ParseTrades(response , market , since , limit );
}

func (this *Btcmarkets) CreateOrder(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 2, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 3, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 4, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 5, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "marketId":*(market ).At(MkString("id") ),
        "amount":this.AmountToPrecision(symbol , amount ),
        "side":OpTrinary(OpEq2(side , MkString("buy") ), MkString("Bid") , MkString("Ask") ),
        }); _ = request;
  lowercaseType := type_.ToLowerCase(); _ = lowercaseType;
  orderTypes := this.SafeValue(*this.At(MkString("options")) , MkString("orderTypes") , MkMap(&VarMap{
            "limit":MkString("Limit") ,
            "market":MkString("Market") ,
            "stop":MkString("Stop") ,
            "stop limit":MkString("Stop Limit") ,
            "take profit":MkString("Take Profit") ,
            })); _ = orderTypes;
  *(request ).At(MkString("type") )= this.SafeString(orderTypes , lowercaseType , type_ );
  priceIsRequired := OpCopy(MkBool(false) ); _ = priceIsRequired;
  triggerPriceIsRequired := OpCopy(MkBool(false) ); _ = triggerPriceIsRequired;
  if IsTrue(OpEq2(lowercaseType , MkString("limit") )) {
    priceIsRequired = OpCopy(MkBool(true) );
  } else {
    if IsTrue(OpEq2(lowercaseType , MkString("stop limit") )) {
      triggerPriceIsRequired = OpCopy(MkBool(true) );
      priceIsRequired = OpCopy(MkBool(true) );
    } else {
      if IsTrue(OpEq2(lowercaseType , MkString("take profit") )) {
        triggerPriceIsRequired = OpCopy(MkBool(true) );
      } else {
        if IsTrue(OpEq2(lowercaseType , MkString("stop") )) {
          triggerPriceIsRequired = OpCopy(MkBool(true) );
        }
      }
    }
  }
  if IsTrue(priceIsRequired ) {
    if IsTrue(OpEq2(price , MkUndefined() )) {
      panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" createOrder() requires a price argument for a ") , OpAdd(type_ , MkString("order") )))));
    } else {
      *(request ).At(MkString("price") )= this.PriceToPrecision(symbol , price );
    }
  }
  if IsTrue(triggerPriceIsRequired ) {
    triggerPrice := this.SafeNumber(params , MkString("triggerPrice") ); _ = triggerPrice;
    params = this.Omit(params , MkString("triggerPrice") );
    if IsTrue(OpEq2(triggerPrice , MkUndefined() )) {
      panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" createOrder() requires a triggerPrice parameter for a ") , OpAdd(type_ , MkString("order") )))));
    } else {
      *(request ).At(MkString("triggerPrice") )= this.PriceToPrecision(symbol , triggerPrice );
    }
  }
  clientOrderId := this.SafeString(params , MkString("clientOrderId") ); _ = clientOrderId;
  if IsTrue(OpNotEq2(clientOrderId , MkUndefined() )) {
    *(request ).At(MkString("clientOrderId") )= OpCopy(clientOrderId );
  }
  params = this.Omit(params , MkString("clientOrderId") );
  response := this.Call(MkString("privatePostOrders"), this.Extend(request , params )); _ = response;
  return this.ParseOrder(response , market );
}

func (this *Btcmarkets) CancelOrders(goArgs ...*Variant) *Variant{
  ids := GoGetArg(goArgs, 0, MkUndefined()); _ = ids;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  for i := MkInteger(0) ; IsTrue(OpLw(i , ids.Length )); OpIncr(& i ){
    *(ids ).At(i )= parseInt(*(ids ).At(i ));
  }
  request := MkMap(&VarMap{"ids":ids }); _ = request;
  return this.Call(MkString("privateDeleteBatchordersIds"), this.Extend(request , params ));
}

func (this *Btcmarkets) CancelOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"id":id }); _ = request;
  return this.Call(MkString("privateDeleteOrdersId"), this.Extend(request , params ));
}

func (this *Btcmarkets) CalculateFee(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 2, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 3, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 4, MkUndefined()); _ = price;
  takerOrMaker := GoGetArg(goArgs, 5, MkString("taker") ); _ = takerOrMaker;
  params := GoGetArg(goArgs, 6, MkMap(&VarMap{})); _ = params;
  market := *(*this.At(MkString("markets")) ).At(symbol ); _ = market;
  rate := *(market ).At(takerOrMaker ); _ = rate;
  currency := OpCopy(MkUndefined() ); _ = currency;
  cost := OpCopy(MkUndefined() ); _ = cost;
  if IsTrue(OpEq2(*(market ).At(MkString("quote") ), MkString("AUD") )) {
    currency = *(market ).At(MkString("quote") );
    cost = parseFloat(this.CostToPrecision(symbol , OpMulti(amount , price )));
  } else {
    currency = *(market ).At(MkString("base") );
    cost = parseFloat(this.AmountToPrecision(symbol , amount ));
  }
  return MkMap(&VarMap{
        "type":takerOrMaker ,
        "currency":currency ,
        "rate":rate ,
        "cost":parseFloat(this.FeeToPrecision(symbol , OpMulti(rate , cost ))),
        });
}

func (this *Btcmarkets) ParseOrderStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "Accepted":MkString("open") ,
        "Placed":MkString("open") ,
        "Partially Matched":MkString("open") ,
        "Fully Matched":MkString("closed") ,
        "Cancelled":MkString("canceled") ,
        "Partially Cancelled":MkString("canceled") ,
        "Failed":MkString("rejected") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Btcmarkets) ParseOrder(goArgs ...*Variant) *Variant{
  order := GoGetArg(goArgs, 0, MkUndefined()); _ = order;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.Parse8601(this.SafeString(order , MkString("creationTime") )); _ = timestamp;
  marketId := this.SafeString(order , MkString("marketId") ); _ = marketId;
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  if IsTrue(OpNotEq2(marketId , MkUndefined() )) {
    if IsTrue(OpHasMember(marketId , *this.At(MkString("markets_by_id")) )) {
      market = *(*this.At(MkString("markets_by_id")) ).At(marketId );
    } else {
      baseId := MkUndefined(); _ = baseId;
      quoteId := MkUndefined(); _ = quoteId;
      ArrayUnpack(marketId.Split(MkString("-") ), &baseId, &quoteId);
      base := this.SafeCurrencyCode(baseId ); _ = base;
      quote := this.SafeCurrencyCode(quoteId ); _ = quote;
      symbol = OpAdd(base , OpAdd(MkString("/") , quote ));
    }
  }
  if IsTrue(OpAnd(OpEq2(symbol , MkUndefined() ), OpNotEq2(market , MkUndefined() ))) {
    symbol = *(market ).At(MkString("symbol") );
  }
  side := this.SafeString(order , MkString("side") ); _ = side;
  if IsTrue(OpEq2(side , MkString("Bid") )) {
    side = MkString("buy") ;
  } else {
    if IsTrue(OpEq2(side , MkString("Ask") )) {
      side = MkString("sell") ;
    }
  }
  type_ := this.SafeStringLower(order , MkString("type") ); _ = type_;
  price := this.SafeNumber(order , MkString("price") ); _ = price;
  amount := this.SafeNumber(order , MkString("amount") ); _ = amount;
  remaining := this.SafeNumber(order , MkString("openAmount") ); _ = remaining;
  status := this.ParseOrderStatus(this.SafeString(order , MkString("status") )); _ = status;
  id := this.SafeString(order , MkString("orderId") ); _ = id;
  clientOrderId := this.SafeString(order , MkString("clientOrderId") ); _ = clientOrderId;
  timeInForce := this.SafeString(order , MkString("timeInForce") ); _ = timeInForce;
  stopPrice := this.SafeNumber(order , MkString("triggerPrice") ); _ = stopPrice;
  postOnly := this.SafeValue(order , MkString("postOnly") ); _ = postOnly;
  return this.SafeOrder(MkMap(&VarMap{
            "info":order ,
            "id":id ,
            "clientOrderId":clientOrderId ,
            "timestamp":timestamp ,
            "datetime":this.Iso8601(timestamp ),
            "lastTradeTimestamp":MkUndefined() ,
            "symbol":symbol ,
            "type":type_ ,
            "timeInForce":timeInForce ,
            "postOnly":postOnly ,
            "side":side ,
            "price":price ,
            "stopPrice":stopPrice ,
            "cost":MkUndefined() ,
            "amount":amount ,
            "filled":MkUndefined() ,
            "remaining":remaining ,
            "average":MkUndefined() ,
            "status":status ,
            "trades":MkUndefined() ,
            "fee":MkUndefined() ,
            }));
}

func (this *Btcmarkets) FetchOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"id":id }); _ = request;
  response := this.Call(MkString("privateGetOrdersId"), this.Extend(request , params )); _ = response;
  return this.ParseOrder(response );
}

func (this *Btcmarkets) FetchOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"status":MkString("all") }); _ = request;
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("marketId") )= *(market ).At(MkString("id") );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("after") )= OpCopy(since );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("privateGetOrders"), this.Extend(request , params )); _ = response;
  return this.ParseOrders(response , market , since , limit );
}

func (this *Btcmarkets) FetchOpenOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"status":MkString("open") }); _ = request;
  return this.FetchOrders(symbol , since , limit , this.Extend(request , params ));
}

func (this *Btcmarkets) FetchClosedOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  orders := this.FetchOrders(symbol , since , limit , params ); _ = orders;
  return this.FilterBy(orders , MkString("status") , MkString("closed") );
}

func (this *Btcmarkets) FetchMyTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("marketId") )= *(market ).At(MkString("id") );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("after") )= OpCopy(since );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("privateGetTrades"), this.Extend(request , params )); _ = response;
  return this.ParseTrades(response , market , since , limit );
}

func (this *Btcmarkets) LookupSymbolFromMarketId(goArgs ...*Variant) *Variant{
  marketId := GoGetArg(goArgs, 0, MkUndefined()); _ = marketId;
  market := OpCopy(MkUndefined() ); _ = market;
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  if IsTrue(OpNotEq2(marketId , MkUndefined() )) {
    if IsTrue(OpHasMember(marketId , *this.At(MkString("markets_by_id")) )) {
      market = *(*this.At(MkString("markets_by_id")) ).At(marketId );
    } else {
      baseId := MkUndefined(); _ = baseId;
      quoteId := MkUndefined(); _ = quoteId;
      ArrayUnpack(marketId.Split(MkString("-") ), &baseId, &quoteId);
      base := this.SafeCurrencyCode(baseId ); _ = base;
      quote := this.SafeCurrencyCode(quoteId ); _ = quote;
      symbol = OpAdd(base , OpAdd(MkString("/") , quote ));
    }
  }
  if IsTrue(OpAnd(OpEq2(symbol , MkUndefined() ), OpNotEq2(market , MkUndefined() ))) {
    symbol = *(market ).At(MkString("symbol") );
  }
  return symbol ;
}

func (this *Btcmarkets) Nonce(goArgs ...*Variant) *Variant{
  return this.Milliseconds();
}

func (this *Btcmarkets) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("public") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  request := OpAdd(MkString("/") , OpAdd(*this.At(MkString("version")) , OpAdd(MkString("/") , this.ImplodeParams(path , params )))); _ = request;
  query := this.Keysort(this.Omit(params , this.ExtractParams(path ))); _ = query;
  if IsTrue(OpEq2(api , MkString("private") )) {
    this.CheckRequiredCredentials();
    nonce := (this.Nonce()).Call(MkString("toString") ); _ = nonce;
    secret := this.Base64ToBinary(this.Encode(*this.At(MkString("secret")) )); _ = secret;
    auth := OpAdd(method , OpAdd(request , nonce )); _ = auth;
    if IsTrue(OpOr(OpEq2(method , MkString("GET") ), OpEq2(method , MkString("DELETE") ))) {
      if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
        OpAddAssign(& request , OpAdd(MkString("?") , this.Urlencode(query )));
      }
    } else {
      body = this.Json(query );
      OpAddAssign(& auth , body );
    }
    signature := this.Hmac(this.Encode(auth ), secret , MkString("sha512") , MkString("base64") ); _ = signature;
    headers = MkMap(&VarMap{
        "Accept":MkString("application/json") ,
        "Accept-Charset":MkString("UTF-8") ,
        "Content-Type":MkString("application/json") ,
        "BM-AUTH-APIKEY":*this.At(MkString("apiKey")) ,
        "BM-AUTH-TIMESTAMP":nonce ,
        "BM-AUTH-SIGNATURE":signature ,
        });
  } else {
    if IsTrue(OpEq2(api , MkString("public") )) {
      if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
        OpAddAssign(& request , OpAdd(MkString("?") , this.Urlencode(query )));
      }
    }
  }
  url := OpAdd(*(*(*this.At(MkString("urls")) ).At(MkString("api") )).At(api ), request ); _ = url;
  return MkMap(&VarMap{
        "url":url ,
        "method":method ,
        "body":body ,
        "headers":headers ,
        });
}

func (this *Btcmarkets) HandleErrors(goArgs ...*Variant) *Variant{
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
  if IsTrue(OpHasMember(MkString("success") , response )) {
    if IsTrue(OpNot(*(response ).At(MkString("success") ))) {
      error := this.SafeString(response , MkString("errorCode") ); _ = error;
      feedback := OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , body )); _ = feedback;
      this.ThrowExactlyMatchedException(*this.At(MkString("exceptions")) , error , feedback );
      panic(NewExchangeError(feedback ));
    }
  }
  if IsTrue(OpGtEq(code , MkInteger(400) )) {
    errorCode := this.SafeString(response , MkString("code") ); _ = errorCode;
    message := this.SafeString(response , MkString("message") ); _ = message;
    feedback := OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , body )); _ = feedback;
    this.ThrowExactlyMatchedException(*this.At(MkString("exceptions")) , errorCode , feedback );
    this.ThrowExactlyMatchedException(*this.At(MkString("exceptions")) , message , feedback );
    panic(NewExchangeError(feedback ));
  }
  return MkUndefined()
}

