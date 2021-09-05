package ccxt

import (
)

type Bittrex struct {
	*ExchangeBase
}

var _ Exchange = (*Bittrex)(nil)

func init() {
	exchange := &Bittrex{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Bittrex) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("bittrex") ,
            "name":MkString("Bittrex") ,
            "countries":MkArray(&VarArray{
                MkString("US") ,
                }),
            "version":MkString("v3") ,
            "rateLimit":MkInteger(1500) ,
            "certified":MkBool(false) ,
            "pro":MkBool(true) ,
            "has":MkMap(&VarMap{
                "CORS":MkBool(false) ,
                "cancelAllOrders":MkBool(true) ,
                "cancelOrder":MkBool(true) ,
                "createDepositAddress":MkBool(true) ,
                "createMarketOrder":MkBool(true) ,
                "createOrder":MkBool(true) ,
                "fetchBalance":MkBool(true) ,
                "fetchDeposits":MkBool(true) ,
                "fetchDepositAddress":MkBool(true) ,
                "fetchClosedOrders":MkBool(true) ,
                "fetchCurrencies":MkBool(true) ,
                "fetchMarkets":MkBool(true) ,
                "fetchMyTrades":MkString("emulated") ,
                "fetchOHLCV":MkBool(true) ,
                "fetchOrder":MkBool(true) ,
                "fetchOrderTrades":MkBool(true) ,
                "fetchOrderBook":MkBool(true) ,
                "fetchOpenOrders":MkBool(true) ,
                "fetchTicker":MkBool(true) ,
                "fetchTickers":MkBool(true) ,
                "fetchTime":MkBool(true) ,
                "fetchTrades":MkBool(true) ,
                "fetchTransactions":MkBool(false) ,
                "fetchWithdrawals":MkBool(true) ,
                "withdraw":MkBool(true) ,
                }),
            "timeframes":MkMap(&VarMap{
                "1m":MkString("MINUTE_1") ,
                "5m":MkString("MINUTE_5") ,
                "1h":MkString("HOUR_1") ,
                "1d":MkString("DAY_1") ,
                }),
            "hostname":MkString("bittrex.com") ,
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/51840849/87153921-edf53180-c2c0-11ea-96b9-f2a9a95a455b.jpg") ,
                "api":MkMap(&VarMap{
                    "public":MkString("https://api.bittrex.com") ,
                    "private":MkString("https://api.bittrex.com") ,
                    }),
                "www":MkString("https://bittrex.com") ,
                "doc":MkArray(&VarArray{
                    MkString("https://bittrex.github.io/api/v3") ,
                    }),
                "fees":MkArray(&VarArray{
                    MkString("https://bittrex.zendesk.com/hc/en-us/articles/115003684371-BITTREX-SERVICE-FEES-AND-WITHDRAWAL-LIMITATIONS") ,
                    MkString("https://bittrex.zendesk.com/hc/en-us/articles/115000199651-What-fees-does-Bittrex-charge-") ,
                    }),
                "referral":MkString("https://bittrex.com/Account/Register?referralCode=1ZE-G0G-M3B") ,
                }),
            "api":MkMap(&VarMap{
                "public":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("ping") ,
                        MkString("currencies") ,
                        MkString("currencies/{symbol}") ,
                        MkString("markets") ,
                        MkString("markets/tickers") ,
                        MkString("markets/summaries") ,
                        MkString("markets/{marketSymbol}") ,
                        MkString("markets/{marketSymbol}/summary") ,
                        MkString("markets/{marketSymbol}/orderbook") ,
                        MkString("markets/{marketSymbol}/trades") ,
                        MkString("markets/{marketSymbol}/ticker") ,
                        MkString("markets/{marketSymbol}/candles/{candleInterval}/recent") ,
                        MkString("markets/{marketSymbol}/candles/{candleInterval}/historical/{year}/{month}/{day}") ,
                        MkString("markets/{marketSymbol}/candles/{candleInterval}/historical/{year}/{month}") ,
                        MkString("markets/{marketSymbol}/candles/{candleInterval}/historical/{year}") ,
                        })}),
                "private":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("account") ,
                        MkString("account/volume") ,
                        MkString("addresses") ,
                        MkString("addresses/{currencySymbol}") ,
                        MkString("balances") ,
                        MkString("balances/{currencySymbol}") ,
                        MkString("deposits/open") ,
                        MkString("deposits/closed") ,
                        MkString("deposits/ByTxId/{txId}") ,
                        MkString("deposits/{depositId}") ,
                        MkString("orders/closed") ,
                        MkString("orders/open") ,
                        MkString("orders/{orderId}") ,
                        MkString("orders/{orderId}/executions") ,
                        MkString("ping") ,
                        MkString("subaccounts/{subaccountId}") ,
                        MkString("subaccounts") ,
                        MkString("withdrawals/open") ,
                        MkString("withdrawals/closed") ,
                        MkString("withdrawals/ByTxId/{txId}") ,
                        MkString("withdrawals/{withdrawalId}") ,
                        MkString("withdrawals/whitelistAddresses") ,
                        MkString("conditional-orders/{conditionalOrderId}") ,
                        MkString("conditional-orders/closed") ,
                        MkString("conditional-orders/open") ,
                        MkString("transfers/sent") ,
                        MkString("transfers/received") ,
                        MkString("transfers/{transferId}") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("addresses") ,
                        MkString("orders") ,
                        MkString("subaccounts") ,
                        MkString("withdrawals") ,
                        MkString("conditional-orders") ,
                        MkString("transfers") ,
                        }),
                    "delete":MkArray(&VarArray{
                        MkString("orders/open") ,
                        MkString("orders/{orderId}") ,
                        MkString("withdrawals/{withdrawalId}") ,
                        MkString("conditional-orders/{conditionalOrderId}") ,
                        }),
                    }),
                }),
            "fees":MkMap(&VarMap{
                "trading":MkMap(&VarMap{
                    "tierBased":MkBool(true) ,
                    "percentage":MkBool(true) ,
                    "maker":this.ParseNumber(MkString("0.0075") ),
                    "taker":this.ParseNumber(MkString("0.0075") ),
                    }),
                "funding":MkMap(&VarMap{
                    "tierBased":MkBool(false) ,
                    "percentage":MkBool(false) ,
                    }),
                }),
            "exceptions":MkMap(&VarMap{
                "exact":MkMap(&VarMap{
                    "BAD_REQUEST":BadRequest ,
                    "STARTDATE_OUT_OF_RANGE":BadRequest ,
                    "APISIGN_NOT_PROVIDED":AuthenticationError ,
                    "INVALID_SIGNATURE":AuthenticationError ,
                    "INVALID_CURRENCY":ExchangeError ,
                    "INVALID_PERMISSION":AuthenticationError ,
                    "INSUFFICIENT_FUNDS":InsufficientFunds ,
                    "INVALID_CEILING_MARKET_BUY":InvalidOrder ,
                    "INVALID_FIAT_ACCOUNT":InvalidOrder ,
                    "INVALID_ORDER_TYPE":InvalidOrder ,
                    "QUANTITY_NOT_PROVIDED":InvalidOrder ,
                    "MIN_TRADE_REQUIREMENT_NOT_MET":InvalidOrder ,
                    "NOT_FOUND":OrderNotFound ,
                    "ORDER_NOT_OPEN":OrderNotFound ,
                    "INVALID_ORDER":InvalidOrder ,
                    "UUID_INVALID":OrderNotFound ,
                    "RATE_NOT_PROVIDED":InvalidOrder ,
                    "INVALID_MARKET":BadSymbol ,
                    "WHITELIST_VIOLATION_IP":PermissionDenied ,
                    "DUST_TRADE_DISALLOWED_MIN_VALUE":InvalidOrder ,
                    "RESTRICTED_MARKET":BadSymbol ,
                    "We are down for scheduled maintenance, but we\u2019ll be back up shortly.":OnMaintenance ,
                    }),
                "broad":MkMap(&VarMap{
                    "throttled":DDoSProtection ,
                    "problem":ExchangeNotAvailable ,
                    }),
                }),
            "options":MkMap(&VarMap{
                "fetchTicker":MkMap(&VarMap{"method":MkString("publicGetMarketsMarketSymbolTicker") }),
                "fetchTickers":MkMap(&VarMap{"method":MkString("publicGetMarketsTickers") }),
                "parseOrderStatus":MkBool(false) ,
                "hasAlreadyAuthenticatedSuccessfully":MkBool(false) ,
                "tag":MkMap(&VarMap{
                    "NXT":MkBool(true) ,
                    "CRYPTO_NOTE_PAYMENTID":MkBool(true) ,
                    "BITSHAREX":MkBool(true) ,
                    "RIPPLE":MkBool(true) ,
                    "NEM":MkBool(true) ,
                    "STELLAR":MkBool(true) ,
                    "STEEM":MkBool(true) ,
                    }),
                "subaccountId":MkUndefined() ,
                "fetchClosedOrdersFilterBySince":MkBool(true) ,
                }),
            "commonCurrencies":MkMap(&VarMap{"REPV2":MkString("REP") }),
            }));
}

func (this *Bittrex) CostToPrecision(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  cost := GoGetArg(goArgs, 1, MkUndefined()); _ = cost;
  return this.DecimalToPrecision(cost , TRUNCATE , *(*(*(*this.At(MkString("markets")) ).At(symbol )).At(MkString("precision") )).At(MkString("price") ), DECIMAL_PLACES );
}

func (this *Bittrex) FeeToPrecision(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  fee := GoGetArg(goArgs, 1, MkUndefined()); _ = fee;
  return this.DecimalToPrecision(fee , TRUNCATE , *(*(*(*this.At(MkString("markets")) ).At(symbol )).At(MkString("precision") )).At(MkString("price") ), DECIMAL_PLACES );
}

func (this *Bittrex) FetchMarkets(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetMarkets"), params ); _ = response;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , response.Length )); OpIncr(& i ){
    market := *(response ).At(i ); _ = market;
    baseId := this.SafeString(market , MkString("baseCurrencySymbol") ); _ = baseId;
    quoteId := this.SafeString(market , MkString("quoteCurrencySymbol") ); _ = quoteId;
    id := this.SafeString(market , MkString("symbol") ); _ = id;
    base := this.SafeCurrencyCode(baseId ); _ = base;
    quote := this.SafeCurrencyCode(quoteId ); _ = quote;
    symbol := OpAdd(base , OpAdd(MkString("/") , quote )); _ = symbol;
    pricePrecision := this.SafeInteger(market , MkString("precision") , MkInteger(8) ); _ = pricePrecision;
    precision := MkMap(&VarMap{
          "amount":MkInteger(8) ,
          "price":pricePrecision ,
          }); _ = precision;
    status := this.SafeString(market , MkString("status") ); _ = status;
    active := OpEq2(status , MkString("ONLINE") ); _ = active;
    result.Push(MkMap(&VarMap{
            "id":id ,
            "symbol":symbol ,
            "base":base ,
            "quote":quote ,
            "baseId":baseId ,
            "quoteId":quoteId ,
            "active":active ,
            "info":market ,
            "precision":precision ,
            "limits":MkMap(&VarMap{
                "amount":MkMap(&VarMap{
                    "min":this.SafeNumber(market , MkString("minTradeSize") ),
                    "max":MkUndefined() ,
                    }),
                "price":MkMap(&VarMap{
                    "min":OpDiv(MkInteger(1) , Math.Pow(MkInteger(10) , *(precision ).At(MkString("price") ))),
                    "max":MkUndefined() ,
                    }),
                "cost":MkMap(&VarMap{
                    "min":MkUndefined() ,
                    "max":MkUndefined() ,
                    }),
                }),
            }));
  }
  return result ;
}

func (this *Bittrex) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  balances := this.Call(MkString("privateGetBalances"), params ); _ = balances;
  result := MkMap(&VarMap{"info":balances }); _ = result;
  indexed := this.IndexBy(balances , MkString("currencySymbol") ); _ = indexed;
  currencyIds := GoGetKeys(indexed ); _ = currencyIds;
  for i := MkInteger(0) ; IsTrue(OpLw(i , currencyIds.Length )); OpIncr(& i ){
    currencyId := *(currencyIds ).At(i ); _ = currencyId;
    code := this.SafeCurrencyCode(currencyId ); _ = code;
    account := this.Account(); _ = account;
    balance := *(indexed ).At(currencyId ); _ = balance;
    *(account ).At(MkString("free") )= this.SafeString(balance , MkString("available") );
    *(account ).At(MkString("total") )= this.SafeString(balance , MkString("total") );
    *(result ).At(code )= OpCopy(account );
  }
  return this.ParseBalance(result );
}

func (this *Bittrex) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"marketSymbol":this.MarketId(symbol )}); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    if IsTrue(OpAnd(OpNotEq2(limit , MkInteger(1) ), OpAnd(OpNotEq2(limit , MkInteger(25) ), OpNotEq2(limit , MkInteger(500) )))) {
      panic(NewBadRequest(OpAdd(*this.At(MkString("id")) , MkString(" fetchOrderBook() limit argument must be undefined, 1, 25 or 500, default is 25") )));
    }
    *(request ).At(MkString("depth") )= OpCopy(limit );
  }
  response := this.Call(MkString("publicGetMarketsMarketSymbolOrderbook"), this.Extend(request , params )); _ = response;
  sequence := this.SafeInteger(*this.At(MkString("last_response_headers")) , MkString("Sequence") ); _ = sequence;
  orderbook := this.ParseOrderBook(response , symbol , MkUndefined() , MkString("bid") , MkString("ask") , MkString("rate") , MkString("quantity") ); _ = orderbook;
  *(orderbook ).At(MkString("nonce") )= OpCopy(sequence );
  return orderbook ;
}

func (this *Bittrex) FetchCurrencies(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetCurrencies"), params ); _ = response;
  result := MkMap(&VarMap{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , response.Length )); OpIncr(& i ){
    currency := *(response ).At(i ); _ = currency;
    id := this.SafeString(currency , MkString("symbol") ); _ = id;
    code := this.SafeCurrencyCode(id ); _ = code;
    precision := MkInteger(8) ; _ = precision;
    fee := this.SafeNumber(currency , MkString("txFee") ); _ = fee;
    isActive := this.SafeString(currency , MkString("status") ); _ = isActive;
    *(result ).At(code )= MkMap(&VarMap{
        "id":id ,
        "code":code ,
        "address":this.SafeString(currency , MkString("baseAddress") ),
        "info":currency ,
        "type":this.SafeString(currency , MkString("coinType") ),
        "name":this.SafeString(currency , MkString("name") ),
        "active":OpEq2(isActive , MkString("ONLINE") ),
        "fee":fee ,
        "precision":precision ,
        "limits":MkMap(&VarMap{
            "amount":MkMap(&VarMap{
                "min":OpDiv(MkInteger(1) , Math.Pow(MkInteger(10) , precision )),
                "max":MkUndefined() ,
                }),
            "withdraw":MkMap(&VarMap{
                "min":fee ,
                "max":MkUndefined() ,
                }),
            }),
        });
  }
  return result ;
}

func (this *Bittrex) ParseTicker(goArgs ...*Variant) *Variant{
  ticker := GoGetArg(goArgs, 0, MkUndefined()); _ = ticker;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.Parse8601(this.SafeString(ticker , MkString("updatedAt") )); _ = timestamp;
  marketId := this.SafeString(ticker , MkString("symbol") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market , MkString("-") ); _ = symbol;
  percentage := this.SafeNumber(ticker , MkString("percentChange") ); _ = percentage;
  last := this.SafeNumber(ticker , MkString("lastTradeRate") ); _ = last;
  return MkMap(&VarMap{
        "symbol":symbol ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "high":this.SafeNumber(ticker , MkString("high") ),
        "low":this.SafeNumber(ticker , MkString("low") ),
        "bid":this.SafeNumber(ticker , MkString("bidRate") ),
        "bidVolume":MkUndefined() ,
        "ask":this.SafeNumber(ticker , MkString("askRate") ),
        "askVolume":MkUndefined() ,
        "vwap":MkUndefined() ,
        "open":MkUndefined() ,
        "close":last ,
        "last":last ,
        "previousClose":MkUndefined() ,
        "change":MkUndefined() ,
        "percentage":percentage ,
        "average":MkUndefined() ,
        "baseVolume":this.SafeNumber(ticker , MkString("volume") ),
        "quoteVolume":this.SafeNumber(ticker , MkString("quoteVolume") ),
        "info":ticker ,
        });
}

func (this *Bittrex) FetchTickers(goArgs ...*Variant) *Variant{
  symbols := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbols;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  options := this.SafeValue(*this.At(MkString("options")) , MkString("fetchTickers") , MkMap(&VarMap{})); _ = options;
  defaultMethod := this.SafeString(options , MkString("method") , MkString("publicGetMarketsTickers") ); _ = defaultMethod;
  method := this.SafeString(params , MkString("method") , defaultMethod ); _ = method;
  params = this.Omit(params , MkString("method") );
  response := this.Call(method , params ); _ = response;
  tickers := MkArray(&VarArray{}); _ = tickers;
  for i := MkInteger(0) ; IsTrue(OpLw(i , response.Length )); OpIncr(& i ){
    ticker := this.ParseTicker(*(response ).At(i )); _ = ticker;
    tickers.Push(ticker );
  }
  return this.FilterByArray(tickers , MkString("symbol") , symbols );
}

func (this *Bittrex) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"marketSymbol":*(market ).At(MkString("id") )}); _ = request;
  options := this.SafeValue(*this.At(MkString("options")) , MkString("fetchTicker") , MkMap(&VarMap{})); _ = options;
  defaultMethod := this.SafeString(options , MkString("method") , MkString("publicGetMarketsMarketSymbolTicker") ); _ = defaultMethod;
  method := this.SafeString(params , MkString("method") , defaultMethod ); _ = method;
  params = this.Omit(params , MkString("method") );
  response := this.Call(method , this.Extend(request , params )); _ = response;
  return this.ParseTicker(response , market );
}

func (this *Bittrex) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.Parse8601(this.SafeString(trade , MkString("executedAt") )); _ = timestamp;
  id := this.SafeString(trade , MkString("id") ); _ = id;
  order := this.SafeString(trade , MkString("orderId") ); _ = order;
  marketId := this.SafeString(trade , MkString("marketSymbol") ); _ = marketId;
  market = this.SafeMarket(marketId , market , MkString("-") );
  priceString := this.SafeString(trade , MkString("rate") ); _ = priceString;
  amountString := this.SafeString(trade , MkString("quantity") ); _ = amountString;
  price := this.ParseNumber(priceString ); _ = price;
  amount := this.ParseNumber(amountString ); _ = amount;
  cost := this.ParseNumber(Precise.StringMul(priceString , amountString )); _ = cost;
  takerOrMaker := OpCopy(MkUndefined() ); _ = takerOrMaker;
  isTaker := this.SafeValue(trade , MkString("isTaker") ); _ = isTaker;
  if IsTrue(OpNotEq2(isTaker , MkUndefined() )) {
    takerOrMaker = OpTrinary(isTaker , MkString("taker") , MkString("maker") );
  }
  fee := OpCopy(MkUndefined() ); _ = fee;
  feeCost := this.SafeNumber(trade , MkString("commission") ); _ = feeCost;
  if IsTrue(OpNotEq2(feeCost , MkUndefined() )) {
    fee = MkMap(&VarMap{
        "cost":feeCost ,
        "currency":*(market ).At(MkString("quote") ),
        });
  }
  side := this.SafeStringLower(trade , MkString("takerSide") ); _ = side;
  return MkMap(&VarMap{
        "info":trade ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "symbol":*(market ).At(MkString("symbol") ),
        "id":id ,
        "order":order ,
        "takerOrMaker":takerOrMaker ,
        "type":MkUndefined() ,
        "side":side ,
        "price":price ,
        "amount":amount ,
        "cost":cost ,
        "fee":fee ,
        });
}

func (this *Bittrex) FetchTime(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetPing"), params ); _ = response;
  return this.SafeInteger(response , MkString("serverTime") );
}

func (this *Bittrex) FetchTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"marketSymbol":this.MarketId(symbol )}); _ = request;
  response := this.Call(MkString("publicGetMarketsMarketSymbolTrades"), this.Extend(request , params )); _ = response;
  return this.ParseTrades(response , market , since , limit );
}

func (this *Bittrex) ParseOHLCV(goArgs ...*Variant) *Variant{
  ohlcv := GoGetArg(goArgs, 0, MkUndefined()); _ = ohlcv;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  return MkArray(&VarArray{
        this.Parse8601(this.SafeString(ohlcv , MkString("startsAt") )),
        this.SafeNumber(ohlcv , MkString("open") ),
        this.SafeNumber(ohlcv , MkString("high") ),
        this.SafeNumber(ohlcv , MkString("low") ),
        this.SafeNumber(ohlcv , MkString("close") ),
        this.SafeNumber(ohlcv , MkString("volume") ),
        });
}

func (this *Bittrex) FetchOHLCV(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  timeframe := GoGetArg(goArgs, 1, MkString("1m") ); _ = timeframe;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  reverseId := OpAdd(*(market ).At(MkString("baseId") ), OpAdd(MkString("-") , *(market ).At(MkString("quoteId") ))); _ = reverseId;
  request := MkMap(&VarMap{
        "candleInterval":*(*this.At(MkString("timeframes")) ).At(timeframe ),
        "marketSymbol":reverseId ,
        }); _ = request;
  method := MkString("publicGetMarketsMarketSymbolCandlesCandleIntervalRecent") ; _ = method;
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    now := this.Milliseconds(); _ = now;
    difference := Math.Abs(OpSub(now , since )); _ = difference;
    sinceDate := this.Ymd(since ); _ = sinceDate;
    parts := sinceDate.Split(MkString("-") ); _ = parts;
    sinceYear := this.SafeInteger(parts , MkInteger(0) ); _ = sinceYear;
    sinceMonth := this.SafeInteger(parts , MkInteger(1) ); _ = sinceMonth;
    sinceDay := this.SafeInteger(parts , MkInteger(2) ); _ = sinceDay;
    if IsTrue(OpEq2(timeframe , MkString("1d") )) {
      if IsTrue(OpGt(difference , MkInteger(31622400000) )) {
        method = MkString("publicGetMarketsMarketSymbolCandlesCandleIntervalHistoricalYear") ;
        *(request ).At(MkString("year") )= OpCopy(sinceYear );
      }
    } else {
      if IsTrue(OpEq2(timeframe , MkString("1h") )) {
        if IsTrue(OpGt(difference , MkInteger(2678400000) )) {
          method = MkString("publicGetMarketsMarketSymbolCandlesCandleIntervalHistoricalYearMonth") ;
          *(request ).At(MkString("year") )= OpCopy(sinceYear );
          *(request ).At(MkString("month") )= OpCopy(sinceMonth );
        }
      } else {
        if IsTrue(OpGt(difference , MkInteger(86400000) )) {
          method = MkString("publicGetMarketsMarketSymbolCandlesCandleIntervalHistoricalYearMonthDay") ;
          *(request ).At(MkString("year") )= OpCopy(sinceYear );
          *(request ).At(MkString("month") )= OpCopy(sinceMonth );
          *(request ).At(MkString("day") )= OpCopy(sinceDay );
        }
      }
    }
  }
  response := this.Call(method , this.Extend(request , params )); _ = response;
  return this.ParseOHLCVs(response , market , timeframe , since , limit );
}

func (this *Bittrex) FetchOpenOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("marketSymbol") )= *(market ).At(MkString("id") );
  }
  response := this.Call(MkString("privateGetOrdersOpen"), this.Extend(request , params )); _ = response;
  return this.ParseOrders(response , market , since , limit );
}

func (this *Bittrex) FetchOrderTrades(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"orderId":id }); _ = request;
  response := this.Call(MkString("privateGetOrdersOrderIdExecutions"), this.Extend(request , params )); _ = response;
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
  }
  return this.ParseTrades(response , market , since , limit );
}

func (this *Bittrex) CreateOrder(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 2, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 3, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 4, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 5, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  uppercaseType := type_.ToUpperCase(); _ = uppercaseType;
  reverseId := OpAdd(*(market ).At(MkString("baseId") ), OpAdd(MkString("-") , *(market ).At(MkString("quoteId") ))); _ = reverseId;
  request := MkMap(&VarMap{
        "marketSymbol":reverseId ,
        "direction":side.ToUpperCase(),
        "type":uppercaseType ,
        }); _ = request;
  isCeilingLimit := OpEq2(uppercaseType , MkString("CEILING_LIMIT") ); _ = isCeilingLimit;
  isCeilingMarket := OpEq2(uppercaseType , MkString("CEILING_MARKET") ); _ = isCeilingMarket;
  isCeilingOrder := OpOr(isCeilingLimit , isCeilingMarket ); _ = isCeilingOrder;
  if IsTrue(isCeilingOrder ) {
    cost := OpCopy(MkUndefined() ); _ = cost;
    if IsTrue(isCeilingLimit ) {
      *(request ).At(MkString("limit") )= this.PriceToPrecision(symbol , price );
      cost = this.SafeNumber2(params , MkString("ceiling") , MkString("cost") , amount );
    } else {
      if IsTrue(isCeilingMarket ) {
        cost = this.SafeNumber2(params , MkString("ceiling") , MkString("cost") );
        if IsTrue(OpEq2(cost , MkUndefined() )) {
          if IsTrue(OpEq2(price , MkUndefined() )) {
            cost = OpCopy(amount );
          } else {
            cost = OpMulti(amount , price );
          }
        }
      }
    }
    params = this.Omit(params , MkArray(&VarArray{
            MkString("ceiling") ,
            MkString("cost") ,
            }));
    *(request ).At(MkString("ceiling") )= this.CostToPrecision(symbol , cost );
    *(request ).At(MkString("timeInForce") )= MkString("IMMEDIATE_OR_CANCEL") ;
  } else {
    *(request ).At(MkString("quantity") )= this.AmountToPrecision(symbol , amount );
    if IsTrue(OpEq2(uppercaseType , MkString("LIMIT") )) {
      *(request ).At(MkString("limit") )= this.PriceToPrecision(symbol , price );
      *(request ).At(MkString("timeInForce") )= MkString("GOOD_TIL_CANCELLED") ;
    } else {
      *(request ).At(MkString("timeInForce") )= MkString("IMMEDIATE_OR_CANCEL") ;
    }
  }
  response := this.Call(MkString("privatePostOrders"), this.Extend(request , params )); _ = response;
  return this.ParseOrder(response , market );
}

func (this *Bittrex) CancelOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"orderId":id }); _ = request;
  response := this.Call(MkString("privateDeleteOrdersOrderId"), this.Extend(request , params )); _ = response;
  return this.Extend(this.ParseOrder(response ), MkMap(&VarMap{
            "id":id ,
            "info":response ,
            "status":MkString("canceled") ,
            }));
}

func (this *Bittrex) CancelAllOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("marketSymbol") )= *(market ).At(MkString("id") );
  }
  response := this.Call(MkString("privateDeleteOrdersOpen"), this.Extend(request , params )); _ = response;
  orders := MkArray(&VarArray{}); _ = orders;
  for i := MkInteger(0) ; IsTrue(OpLw(i , response.Length )); OpIncr(& i ){
    result := this.SafeValue(*(response ).At(i ), MkString("result") , MkMap(&VarMap{})); _ = result;
    orders.Push(result );
  }
  return this.ParseOrders(orders , market );
}

func (this *Bittrex) FetchDeposits(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  currency := OpCopy(MkUndefined() ); _ = currency;
  if IsTrue(OpNotEq2(code , MkUndefined() )) {
    currency = this.Currency(code );
    *(request ).At(MkString("currencySymbol") )= *(currency ).At(MkString("id") );
  }
  response := this.Call(MkString("privateGetDepositsClosed"), this.Extend(request , params )); _ = response;
  return this.ParseTransactions(response , currency , MkUndefined() , limit );
}

func (this *Bittrex) FetchWithdrawals(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  currency := OpCopy(MkUndefined() ); _ = currency;
  if IsTrue(OpNotEq2(code , MkUndefined() )) {
    currency = this.Currency(code );
    *(request ).At(MkString("currencySymbol") )= *(currency ).At(MkString("id") );
  }
  response := this.Call(MkString("privateGetWithdrawalsClosed"), this.Extend(request , params )); _ = response;
  return this.ParseTransactions(response , currency , since , limit );
}

func (this *Bittrex) ParseTransaction(goArgs ...*Variant) *Variant{
  transaction := GoGetArg(goArgs, 0, MkUndefined()); _ = transaction;
  currency := GoGetArg(goArgs, 1, MkUndefined() ); _ = currency;
  id := this.SafeString(transaction , MkString("id") ); _ = id;
  amount := this.SafeNumber(transaction , MkString("quantity") ); _ = amount;
  address := this.SafeString(transaction , MkString("cryptoAddress") ); _ = address;
  txid := this.SafeString(transaction , MkString("txId") ); _ = txid;
  updated := this.Parse8601(this.SafeString(transaction , MkString("updatedAt") )); _ = updated;
  opened := this.Parse8601(this.SafeString(transaction , MkString("createdAt") )); _ = opened;
  timestamp := OpTrinary(opened , opened , updated ); _ = timestamp;
  type_ := OpTrinary(OpEq2(opened , MkUndefined() ), MkString("deposit") , MkString("withdrawal") ); _ = type_;
  currencyId := this.SafeString(transaction , MkString("currencySymbol") ); _ = currencyId;
  code := this.SafeCurrencyCode(currencyId , currency ); _ = code;
  status := MkString("pending") ; _ = status;
  if IsTrue(OpEq2(type_ , MkString("deposit") )) {
    status = MkString("ok") ;
  } else {
    responseStatus := this.SafeString(transaction , MkString("status") ); _ = responseStatus;
    if IsTrue(OpEq2(responseStatus , MkString("ERROR_INVALID_ADDRESS") )) {
      status = MkString("failed") ;
    } else {
      if IsTrue(OpEq2(responseStatus , MkString("CANCELLED") )) {
        status = MkString("canceled") ;
      } else {
        if IsTrue(OpEq2(responseStatus , MkString("PENDING") )) {
          status = MkString("pending") ;
        } else {
          if IsTrue(OpEq2(responseStatus , MkString("COMPLETED") )) {
            status = MkString("ok") ;
          } else {
            if IsTrue(OpAnd(OpEq2(responseStatus , MkString("AUTHORIZED") ), OpNotEq2(txid , MkUndefined() ))) {
              status = MkString("ok") ;
            }
          }
        }
      }
    }
  }
  feeCost := this.SafeNumber(transaction , MkString("txCost") ); _ = feeCost;
  if IsTrue(OpEq2(feeCost , MkUndefined() )) {
    if IsTrue(OpEq2(type_ , MkString("deposit") )) {
      feeCost = MkInteger(0) ;
    }
  }
  return MkMap(&VarMap{
        "info":transaction ,
        "id":id ,
        "currency":code ,
        "amount":amount ,
        "address":address ,
        "tag":MkUndefined() ,
        "status":status ,
        "type":type_ ,
        "updated":updated ,
        "txid":txid ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "fee":MkMap(&VarMap{
            "currency":code ,
            "cost":feeCost ,
            }),
        });
}

func (this *Bittrex) ParseTimeInForce(goArgs ...*Variant) *Variant{
  timeInForce := GoGetArg(goArgs, 0, MkUndefined()); _ = timeInForce;
  timeInForces := MkMap(&VarMap{
        "GOOD_TIL_CANCELLED":MkString("GTC") ,
        "IMMEDIATE_OR_CANCEL":MkString("IOC") ,
        "FILL_OR_KILL":MkString("FOK") ,
        "POST_ONLY_GOOD_TIL_CANCELLED":MkString("PO") ,
        }); _ = timeInForces;
  return this.SafeString(timeInForces , timeInForce , timeInForce );
}

func (this *Bittrex) ParseOrder(goArgs ...*Variant) *Variant{
  order := GoGetArg(goArgs, 0, MkUndefined()); _ = order;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  marketSymbol := this.SafeString(order , MkString("marketSymbol") ); _ = marketSymbol;
  market = this.SafeMarket(marketSymbol , market , MkString("-") );
  symbol := *(market ).At(MkString("symbol") ); _ = symbol;
  feeCurrency := *(market ).At(MkString("quote") ); _ = feeCurrency;
  direction := this.SafeStringLower(order , MkString("direction") ); _ = direction;
  createdAt := this.SafeString(order , MkString("createdAt") ); _ = createdAt;
  updatedAt := this.SafeString(order , MkString("updatedAt") ); _ = updatedAt;
  closedAt := this.SafeString(order , MkString("closedAt") ); _ = closedAt;
  lastTradeTimestamp := OpCopy(MkUndefined() ); _ = lastTradeTimestamp;
  if IsTrue(OpNotEq2(closedAt , MkUndefined() )) {
    lastTradeTimestamp = this.Parse8601(closedAt );
  } else {
    if IsTrue(updatedAt ) {
      lastTradeTimestamp = this.Parse8601(updatedAt );
    }
  }
  timestamp := this.Parse8601(createdAt ); _ = timestamp;
  type_ := this.SafeStringLower(order , MkString("type") ); _ = type_;
  quantity := this.SafeNumber(order , MkString("quantity") ); _ = quantity;
  limit := this.SafeNumber(order , MkString("limit") ); _ = limit;
  fillQuantity := this.SafeNumber(order , MkString("fillQuantity") ); _ = fillQuantity;
  commission := this.SafeNumber(order , MkString("commission") ); _ = commission;
  proceeds := this.SafeNumber(order , MkString("proceeds") ); _ = proceeds;
  status := this.SafeStringLower(order , MkString("status") ); _ = status;
  timeInForce := this.ParseTimeInForce(this.SafeString(order , MkString("timeInForce") )); _ = timeInForce;
  postOnly := OpEq2(timeInForce , MkString("PO") ); _ = postOnly;
  return this.SafeOrder(MkMap(&VarMap{
            "id":this.SafeString(order , MkString("id") ),
            "clientOrderId":MkUndefined() ,
            "timestamp":timestamp ,
            "datetime":this.Iso8601(timestamp ),
            "lastTradeTimestamp":lastTradeTimestamp ,
            "symbol":symbol ,
            "type":type_ ,
            "timeInForce":timeInForce ,
            "postOnly":postOnly ,
            "side":direction ,
            "price":limit ,
            "stopPrice":MkUndefined() ,
            "cost":proceeds ,
            "average":MkUndefined() ,
            "amount":quantity ,
            "filled":fillQuantity ,
            "remaining":MkUndefined() ,
            "status":status ,
            "fee":MkMap(&VarMap{
                "cost":commission ,
                "currency":feeCurrency ,
                }),
            "info":order ,
            "trades":MkUndefined() ,
            }));
}

func (this *Bittrex) ParseOrders(goArgs ...*Variant) *Variant{
  orders := GoGetArg(goArgs, 0, MkUndefined()); _ = orders;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  if IsTrue(*(*this.At(MkString("options")) ).At(MkString("fetchClosedOrdersFilterBySince") )) {
    return this.ParseOrders(orders , market , since , limit , params );
  } else {
    return this.ParseOrders(orders , market , MkUndefined() , limit , params );
  }
  return MkUndefined()
}

func (this *Bittrex) ParseOrderStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "CLOSED":MkString("closed") ,
        "OPEN":MkString("open") ,
        "CANCELLED":MkString("canceled") ,
        "CANCELED":MkString("canceled") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Bittrex) FetchOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := OpCopy(MkUndefined() ); _ = response;
  {
  ret__ := func(this *Bittrex) (ret_ *Variant) {
    defer func() {
      if e := recover().(*Variant); e != nil {
        ret_ = func(this *Bittrex) *Variant {
          // catch block:
          if IsTrue(*this.At(MkString("last_json_response")) ) {
            message := this.SafeString(*this.At(MkString("last_json_response")) , MkString("message") ); _ = message;
            if IsTrue(OpEq2(message , MkString("UUID_INVALID") )) {
              panic(NewOrderNotFound(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" fetchOrder() error: ") , *this.At(MkString("last_http_response")) ))));
            }
          }
          panic(e );
          return nil
        }(this)
      }
    }()
    // try block:
    request := MkMap(&VarMap{"orderId":id }); _ = request;
    response = this.Call(MkString("privateGetOrdersOrderId"), this.Extend(request , params ));
    return nil
  }(this)
  if ret__ != nil {
    return ret__;
   }
  }
  return this.ParseOrder(response );
}

func (this *Bittrex) OrderToTrade(goArgs ...*Variant) *Variant{
  order := GoGetArg(goArgs, 0, MkUndefined()); _ = order;
  timestamp := this.SafeInteger2(order , MkString("lastTradeTimestamp") , MkString("timestamp") ); _ = timestamp;
  return MkMap(&VarMap{
        "id":this.SafeString(order , MkString("id") ),
        "side":this.SafeString(order , MkString("side") ),
        "order":this.SafeString(order , MkString("id") ),
        "type":this.SafeString(order , MkString("type") ),
        "price":this.SafeNumber(order , MkString("average") ),
        "amount":this.SafeNumber(order , MkString("filled") ),
        "cost":this.SafeNumber(order , MkString("cost") ),
        "symbol":this.SafeString(order , MkString("symbol") ),
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "fee":this.SafeValue(order , MkString("fee") ),
        "info":order ,
        "takerOrMaker":MkUndefined() ,
        });
}

func (this *Bittrex) OrdersToTrades(goArgs ...*Variant) *Variant{
  orders := GoGetArg(goArgs, 0, MkUndefined()); _ = orders;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , orders.Length )); OpIncr(& i ){
    result.Push(this.OrderToTrade(*(orders ).At(i )));
  }
  return result ;
}

func (this *Bittrex) FetchMyTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("pageSize") )= OpCopy(limit );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("startDate") )= OpAdd(this.Ymdhms(since , MkString("T") ), MkString("Z") );
  }
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("marketSymbol") )= OpAdd(*(market ).At(MkString("base") ), OpAdd(MkString("-") , *(market ).At(MkString("quote") )));
  }
  response := this.Call(MkString("privateGetOrdersClosed"), this.Extend(request , params )); _ = response;
  orders := this.ParseOrders(response , market ); _ = orders;
  trades := this.OrdersToTrades(orders ); _ = trades;
  return this.FilterBySymbolSinceLimit(trades , symbol , since , limit );
}

func (this *Bittrex) FetchClosedOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("pageSize") )= OpCopy(limit );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("startDate") )= OpAdd(this.Ymdhms(since , MkString("T") ), MkString("Z") );
  }
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("marketSymbol") )= OpAdd(*(market ).At(MkString("base") ), OpAdd(MkString("-") , *(market ).At(MkString("quote") )));
  }
  response := this.Call(MkString("privateGetOrdersClosed"), this.Extend(request , params )); _ = response;
  return this.ParseOrders(response , market , since , limit );
}

func (this *Bittrex) CreateDepositAddress(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  currency := this.Currency(code ); _ = currency;
  request := MkMap(&VarMap{"currencySymbol":*(currency ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("privatePostAddresses"), this.Extend(request , params )); _ = response;
  address := this.SafeString(response , MkString("cryptoAddress") ); _ = address;
  message := this.SafeString(response , MkString("status") ); _ = message;
  if IsTrue(OpOr(OpNot(address ), OpEq2(message , MkString("REQUESTED") ))) {
    panic(NewAddressPending(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" the address for ") , OpAdd(code , MkString(" is being generated (pending, not ready yet, retry again later)") )))));
  }
  tag := this.SafeString(response , MkString("cryptoAddressTag") ); _ = tag;
  if IsTrue(OpAnd(OpEq2(tag , MkUndefined() ), OpHasMember(*(currency ).At(MkString("type") ), *(*this.At(MkString("options")) ).At(MkString("tag") )))) {
    tag = OpCopy(address );
    address = *(currency ).At(MkString("address") );
  }
  this.CheckAddress(address );
  return MkMap(&VarMap{
        "currency":code ,
        "address":address ,
        "tag":tag ,
        "info":response ,
        });
}

func (this *Bittrex) FetchDepositAddress(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  currency := this.Currency(code ); _ = currency;
  request := MkMap(&VarMap{"currencySymbol":*(currency ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("privateGetAddressesCurrencySymbol"), this.Extend(request , params )); _ = response;
  address := this.SafeString(response , MkString("cryptoAddress") ); _ = address;
  message := this.SafeString(response , MkString("status") ); _ = message;
  if IsTrue(OpOr(OpNot(address ), OpEq2(message , MkString("REQUESTED") ))) {
    panic(NewAddressPending(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" the address for ") , OpAdd(code , MkString(" is being generated (pending, not ready yet, retry again later)") )))));
  }
  tag := this.SafeString(response , MkString("cryptoAddressTag") ); _ = tag;
  if IsTrue(OpAnd(OpEq2(tag , MkUndefined() ), OpHasMember(*(currency ).At(MkString("type") ), *(*this.At(MkString("options")) ).At(MkString("tag") )))) {
    tag = OpCopy(address );
    address = *(currency ).At(MkString("address") );
  }
  this.CheckAddress(address );
  return MkMap(&VarMap{
        "currency":code ,
        "address":address ,
        "tag":tag ,
        "info":response ,
        });
}

func (this *Bittrex) Withdraw(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  amount := GoGetArg(goArgs, 1, MkUndefined()); _ = amount;
  address := GoGetArg(goArgs, 2, MkUndefined()); _ = address;
  tag := GoGetArg(goArgs, 3, MkUndefined() ); _ = tag;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.CheckAddress(address );
  this.LoadMarkets();
  currency := this.Currency(code ); _ = currency;
  request := MkMap(&VarMap{
        "currencySymbol":*(currency ).At(MkString("id") ),
        "quantity":amount ,
        "cryptoAddress":address ,
        }); _ = request;
  if IsTrue(OpNotEq2(tag , MkUndefined() )) {
    *(request ).At(MkString("cryptoAddressTag") )= OpCopy(tag );
  }
  response := this.Call(MkString("privatePostWithdrawals"), this.Extend(request , params )); _ = response;
  id := this.SafeString(response , MkString("id") ); _ = id;
  return MkMap(&VarMap{
        "info":response ,
        "id":id ,
        });
}

func (this *Bittrex) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("v3") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  url := OpAdd(this.ImplodeParams(*(*(*this.At(MkString("urls")) ).At(MkString("api") )).At(api ), MkMap(&VarMap{"hostname":*this.At(MkString("hostname")) })), MkString("/") ); _ = url;
  if IsTrue(OpEq2(api , MkString("private") )) {
    OpAddAssign(& url , OpAdd(*this.At(MkString("version")) , MkString("/") ));
    this.CheckRequiredCredentials();
    OpAddAssign(& url , this.ImplodeParams(path , params ));
    params = this.Omit(params , this.ExtractParams(path ));
    hashString := MkString("") ; _ = hashString;
    if IsTrue(OpEq2(method , MkString("POST") )) {
      body = this.Json(params );
      hashString = OpCopy(body );
    } else {
      if IsTrue(*(GoGetKeys(params )).At(MkString("length") )) {
        OpAddAssign(& url , OpAdd(MkString("?") , this.Rawencode(params )));
      }
    }
    contentHash := this.Hash(this.Encode(hashString ), MkString("sha512") , MkString("hex") ); _ = contentHash;
    timestamp := (this.Milliseconds()).Call(MkString("toString") ); _ = timestamp;
    auth := OpAdd(timestamp , OpAdd(url , OpAdd(method , contentHash ))); _ = auth;
    subaccountId := this.SafeValue(*this.At(MkString("options")) , MkString("subaccountId") ); _ = subaccountId;
    if IsTrue(OpNotEq2(subaccountId , MkUndefined() )) {
      OpAddAssign(& auth , subaccountId );
    }
    signature := this.Hmac(this.Encode(auth ), this.Encode(*this.At(MkString("secret")) ), MkString("sha512") ); _ = signature;
    headers = MkMap(&VarMap{
        "Api-Key":*this.At(MkString("apiKey")) ,
        "Api-Timestamp":timestamp ,
        "Api-Content-Hash":contentHash ,
        "Api-Signature":signature ,
        });
    if IsTrue(OpNotEq2(subaccountId , MkUndefined() )) {
      *(headers ).At(MkString("Api-Subaccount-Id") )= OpCopy(subaccountId );
    }
    if IsTrue(OpEq2(method , MkString("POST") )) {
      *(headers ).At(MkString("Content-Type") )= MkString("application/json") ;
    }
  } else {
    if IsTrue(OpEq2(api , MkString("public") )) {
      OpAddAssign(& url , OpAdd(*this.At(MkString("version")) , MkString("/") ));
    }
    OpAddAssign(& url , this.ImplodeParams(path , params ));
    params = this.Omit(params , this.ExtractParams(path ));
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

func (this *Bittrex) HandleErrors(goArgs ...*Variant) *Variant{
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
  if IsTrue(OpEq2(*(body ).At(MkInteger(0) ), MkString("{") )) {
    feedback := OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , body )); _ = feedback;
    success := this.SafeValue(response , MkString("success") ); _ = success;
    if IsTrue(OpEq2(success , MkUndefined() )) {
      code := this.SafeString(response , MkString("code") ); _ = code;
      if IsTrue(OpAnd(OpEq2(code , MkString("NOT_FOUND") ), OpGtEq(url.IndexOf(MkString("addresses") ), MkInteger(0) ))) {
        panic(NewInvalidAddress(feedback ));
      }
      if IsTrue(OpNotEq2(code , MkUndefined() )) {
        this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("exact") ), code , feedback );
        this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("broad") ), code , feedback );
      }
      return MkUndefined();
    }
    if IsTrue(OpEq2(OpType(success ), MkString("string") )) {
      success = OpEq2(success , MkString("true") );
    }
    if IsTrue(OpNot(success )) {
      message := this.SafeString(response , MkString("message") ); _ = message;
      if IsTrue(OpEq2(message , MkString("APIKEY_INVALID") )) {
        if IsTrue(*(*this.At(MkString("options")) ).At(MkString("hasAlreadyAuthenticatedSuccessfully") )) {
          panic(NewDDoSProtection(feedback ));
        } else {
          panic(NewAuthenticationError(feedback ));
        }
      }
      if IsTrue(OpEq2(message , MkString("INVALID_ORDER") )) {
        cancel := MkString("cancel") ; _ = cancel;
        indexOfCancel := url.IndexOf(cancel ); _ = indexOfCancel;
        if IsTrue(OpGtEq(indexOfCancel , MkInteger(0) )) {
          urlParts := url.Split(MkString("?") ); _ = urlParts;
          numParts := OpCopy(urlParts.Length ); _ = numParts;
          if IsTrue(OpGt(numParts , MkInteger(1) )) {
            query := *(urlParts ).At(MkInteger(1) ); _ = query;
            params := query.Split(MkString("&") ); _ = params;
            numParams := OpCopy(params.Length ); _ = numParams;
            orderId := OpCopy(MkUndefined() ); _ = orderId;
            for i := MkInteger(0) ; IsTrue(OpLw(i , numParams )); OpIncr(& i ){
              param := *(params ).At(i ); _ = param;
              keyValue := param.Split(MkString("=") ); _ = keyValue;
              if IsTrue(OpEq2(*(keyValue ).At(MkInteger(0) ), MkString("uuid") )) {
                orderId = *(keyValue ).At(MkInteger(1) );
                break ;
              }
            }
            if IsTrue(OpNotEq2(orderId , MkUndefined() )) {
              panic(NewOrderNotFound(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" cancelOrder ") , OpAdd(orderId , OpAdd(MkString(" ") , this.Json(response )))))));
            } else {
              panic(NewOrderNotFound(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" cancelOrder ") , this.Json(response )))));
            }
          }
        }
      }
      this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("exact") ), message , feedback );
      if IsTrue(OpNotEq2(message , MkUndefined() )) {
        this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("broad") ), message , feedback );
      }
      panic(NewExchangeError(feedback ));
    }
  }
  return MkUndefined()
}

