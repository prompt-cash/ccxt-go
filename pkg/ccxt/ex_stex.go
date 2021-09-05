package ccxt

import (
)

type Stex struct {
	*ExchangeBase
}

var _ Exchange = (*Stex)(nil)

func init() {
	exchange := &Stex{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Stex) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("stex") ,
            "name":MkString("STEX") ,
            "countries":MkArray(&VarArray{
                MkString("EE") ,
                }),
            "rateLimit":MkInteger(500) ,
            "certified":MkBool(false) ,
            "has":MkMap(&VarMap{
                "cancelAllOrders":MkBool(true) ,
                "cancelOrder":MkBool(true) ,
                "CORS":MkBool(false) ,
                "createDepositAddress":MkBool(true) ,
                "createMarketOrder":MkBool(false) ,
                "createOrder":MkBool(true) ,
                "fetchBalance":MkBool(true) ,
                "fetchCurrencies":MkBool(true) ,
                "fetchDepositAddress":MkBool(true) ,
                "fetchDeposits":MkBool(true) ,
                "fetchFundingFees":MkBool(true) ,
                "fetchMarkets":MkBool(true) ,
                "fetchMyTrades":MkBool(true) ,
                "fetchOHLCV":MkBool(true) ,
                "fetchOpenOrders":MkBool(true) ,
                "fetchOrder":MkBool(true) ,
                "fetchOrderBook":MkBool(true) ,
                "fetchOrderTrades":MkBool(true) ,
                "fetchTicker":MkBool(true) ,
                "fetchTickers":MkBool(true) ,
                "fetchTime":MkBool(true) ,
                "fetchTrades":MkBool(true) ,
                "fetchWithdrawals":MkBool(true) ,
                "withdraw":MkBool(true) ,
                }),
            "version":MkString("v3") ,
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/1294454/69680782-03fd0b80-10bd-11ea-909e-7f603500e9cc.jpg") ,
                "api":MkString("https://api3.stex.com") ,
                "www":MkString("https://www.stex.com") ,
                "doc":MkArray(&VarArray{
                    MkString("https://help.stex.com/en/collections/1593608-api-v3-documentation") ,
                    }),
                "fees":MkString("https://app.stex.com/en/pairs-specification") ,
                "referral":MkString("https://app.stex.com?ref=36416021") ,
                }),
            "requiredCredentials":MkMap(&VarMap{
                "apiKey":MkBool(false) ,
                "secret":MkBool(false) ,
                "token":MkBool(true) ,
                }),
            "timeframes":MkMap(&VarMap{
                "1m":MkString("1") ,
                "5m":MkString("5") ,
                "30m":MkString("30") ,
                "1h":MkString("60") ,
                "4h":MkString("240") ,
                "12h":MkString("720") ,
                "1d":MkString("1D") ,
                }),
            "api":MkMap(&VarMap{
                "public":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("currencies") ,
                        MkString("currencies/{currencyId}") ,
                        MkString("markets") ,
                        MkString("pairs-groups") ,
                        MkString("currency_pairs/list/{code}") ,
                        MkString("currency_pairs/group/{currencyPairGroupId}") ,
                        MkString("currency_pairs/{currencyPairId}") ,
                        MkString("ticker") ,
                        MkString("ticker/{currencyPairId}") ,
                        MkString("trades/{currencyPairId}") ,
                        MkString("orderbook/{currencyPairId}") ,
                        MkString("chart/{currencyPairId}/{candlesType}") ,
                        MkString("deposit-statuses") ,
                        MkString("deposit-statuses/{statusId}") ,
                        MkString("withdrawal-statuses") ,
                        MkString("withdrawal-statuses/{statusId}") ,
                        MkString("ping") ,
                        MkString("mobile-versions") ,
                        })}),
                "trading":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("fees/{currencyPairId}") ,
                        MkString("orders") ,
                        MkString("orders/{currencyPairId}") ,
                        MkString("order/{orderId}") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("orders/{currencyPairId}") ,
                        }),
                    "delete":MkArray(&VarArray{
                        MkString("orders") ,
                        MkString("orders/{currencyPairId}") ,
                        MkString("order/{orderId}") ,
                        }),
                    }),
                "reports":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("orders") ,
                        MkString("orders/{orderId}") ,
                        MkString("trades/{currencyPairId}") ,
                        MkString("background/{listMode}") ,
                        MkString("background/{id}") ,
                        MkString("background/download/{id}") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("background/create") ,
                        }),
                    "delete":MkArray(&VarArray{
                        MkString("background/{id}") ,
                        }),
                    }),
                "profile":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("info") ,
                        MkString("wallets") ,
                        MkString("wallets/{walletId}") ,
                        MkString("wallets/address/{walletId}") ,
                        MkString("deposits") ,
                        MkString("deposits/{id}") ,
                        MkString("withdrawals") ,
                        MkString("withdrawals/{id}") ,
                        MkString("notifications") ,
                        MkString("favorite/currency_pairs") ,
                        MkString("token-scopes") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("wallets/burn/{walletId}") ,
                        MkString("wallets/{currencyId}") ,
                        MkString("wallets/address/{walletId}") ,
                        MkString("withdraw") ,
                        MkString("referral/program") ,
                        MkString("referral/insert/{code}") ,
                        MkString("referral/bonus_transfer/{currencyId}") ,
                        }),
                    "put":MkArray(&VarArray{
                        MkString("profile/favorite/currency_pairs/set") ,
                        }),
                    "delete":MkArray(&VarArray{
                        MkString("profile/withdraw/{withdrawalId}") ,
                        }),
                    }),
                "verification":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("verification/countries") ,
                        MkString("verification/stex") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("verification/stex") ,
                        }),
                    }),
                "settings":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("notifications/{event}") ,
                        MkString("notifications") ,
                        }),
                    "put":MkArray(&VarArray{
                        MkString("notifications") ,
                        MkString("notifications/set") ,
                        }),
                    }),
                }),
            "fees":MkMap(&VarMap{"trading":MkMap(&VarMap{
                    "tierBased":MkBool(false) ,
                    "percentage":MkBool(true) ,
                    "taker":this.ParseNumber(MkString("0.002") ),
                    "maker":this.ParseNumber(MkString("0.002") ),
                    })}),
            "commonCurrencies":MkMap(&VarMap{
                "BC":MkString("Bitcoin Confidential") ,
                "BITS":MkString("Bitcoinus") ,
                "BITSW":MkString("BITS") ,
                "BHD":MkString("Bithold") ,
                "BTH":MkString("Bithereum") ,
                "MPH":MkString("Chasyr Token") ,
                "SBTC":MkString("SBTCT") ,
                }),
            "options":MkMap(&VarMap{"parseOrderToPrecision":MkBool(false) }),
            "exceptions":MkMap(&VarMap{
                "exact":MkMap(&VarMap{
                    "Wrong parameters":BadRequest ,
                    "Unauthenticated.":AuthenticationError ,
                    "Server Error":ExchangeError ,
                    "This feature is only enabled for users verifies by Cryptonomica":PermissionDenied ,
                    "Too Many Attempts.":DDoSProtection ,
                    "Selected Pair is disabled":BadSymbol ,
                    }),
                "broad":MkMap(&VarMap{"Not enough":InsufficientFunds }),
                }),
            }));
}

func (this *Stex) FetchCurrencies(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetCurrencies"), params ); _ = response;
  result := MkMap(&VarMap{}); _ = result;
  currencies := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = currencies;
  for i := MkInteger(0) ; IsTrue(OpLw(i , currencies.Length )); OpIncr(& i ){
    currency := *(currencies ).At(i ); _ = currency;
    id := this.SafeString(currency , MkString("id") ); _ = id;
    numericId := this.SafeInteger(currency , MkString("id") ); _ = numericId;
    code := this.SafeCurrencyCode(this.SafeString(currency , MkString("code") )); _ = code;
    precision := this.SafeString(currency , MkString("precision") ); _ = precision;
    amountLimit := this.ParsePrecision(precision ); _ = amountLimit;
    fee := this.SafeNumber(currency , MkString("withdrawal_fee_const") ); _ = fee;
    active := this.SafeValue(currency , MkString("active") , MkBool(true) ); _ = active;
    *(result ).At(code )= MkMap(&VarMap{
        "id":id ,
        "numericId":numericId ,
        "code":code ,
        "info":currency ,
        "type":MkUndefined() ,
        "name":this.SafeString(currency , MkString("name") ),
        "active":active ,
        "fee":fee ,
        "precision":parseInt(precision ),
        "limits":MkMap(&VarMap{
            "amount":MkMap(&VarMap{
                "min":this.ParseNumber(amountLimit ),
                "max":MkUndefined() ,
                }),
            "deposit":MkMap(&VarMap{
                "min":this.SafeNumber(currency , MkString("minimum_deposit_amount") ),
                "max":MkUndefined() ,
                }),
            "withdraw":MkMap(&VarMap{
                "min":this.SafeNumber(currency , MkString("minimum_withdrawal_amount") ),
                "max":MkUndefined() ,
                }),
            }),
        });
  }
  return result ;
}

func (this *Stex) FetchMarkets(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"code":MkString("ALL") }); _ = request;
  response := this.Call(MkString("publicGetCurrencyPairsListCode"), this.Extend(request , params )); _ = response;
  result := MkArray(&VarArray{}); _ = result;
  markets := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = markets;
  for i := MkInteger(0) ; IsTrue(OpLw(i , markets.Length )); OpIncr(& i ){
    market := *(markets ).At(i ); _ = market;
    id := this.SafeString(market , MkString("id") ); _ = id;
    numericId := this.SafeInteger(market , MkString("id") ); _ = numericId;
    baseId := this.SafeString(market , MkString("currency_id") ); _ = baseId;
    quoteId := this.SafeString(market , MkString("market_currency_id") ); _ = quoteId;
    baseNumericId := this.SafeInteger(market , MkString("currency_id") ); _ = baseNumericId;
    quoteNumericId := this.SafeInteger(market , MkString("market_currency_id") ); _ = quoteNumericId;
    base := this.SafeCurrencyCode(this.SafeString(market , MkString("currency_code") )); _ = base;
    quote := this.SafeCurrencyCode(this.SafeString(market , MkString("market_code") )); _ = quote;
    symbol := OpAdd(base , OpAdd(MkString("/") , quote )); _ = symbol;
    precision := MkMap(&VarMap{
          "amount":this.SafeInteger(market , MkString("currency_precision") ),
          "price":this.SafeInteger(market , MkString("market_precision") ),
          }); _ = precision;
    active := this.SafeValue(market , MkString("active") ); _ = active;
    minBuyPrice := this.SafeNumber(market , MkString("min_buy_price") ); _ = minBuyPrice;
    minSellPrice := this.SafeNumber(market , MkString("min_sell_price") ); _ = minSellPrice;
    minPrice := Math.Max(minBuyPrice , minSellPrice ); _ = minPrice;
    buyFee := OpDiv(this.SafeNumber(market , MkString("buy_fee_percent") ), MkInteger(100) ); _ = buyFee;
    sellFee := OpDiv(this.SafeNumber(market , MkString("sell_fee_percent") ), MkInteger(100) ); _ = sellFee;
    fee := Math.Max(buyFee , sellFee ); _ = fee;
    result.Push(MkMap(&VarMap{
            "id":id ,
            "numericId":numericId ,
            "symbol":symbol ,
            "base":base ,
            "quote":quote ,
            "baseId":baseId ,
            "quoteId":quoteId ,
            "baseNumericId":baseNumericId ,
            "quoteNumericId":quoteNumericId ,
            "info":market ,
            "active":active ,
            "maker":fee ,
            "taker":fee ,
            "precision":precision ,
            "limits":MkMap(&VarMap{
                "amount":MkMap(&VarMap{
                    "min":this.SafeNumber(market , MkString("min_order_amount") ),
                    "max":MkUndefined() ,
                    }),
                "price":MkMap(&VarMap{
                    "min":minPrice ,
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

func (this *Stex) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"currencyPairId":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("publicGetTickerCurrencyPairId"), this.Extend(request , params )); _ = response;
  ticker := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = ticker;
  return this.ParseTicker(ticker , market );
}

func (this *Stex) FetchTime(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetPing"), params ); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  serverDatetime := this.SafeValue(data , MkString("server_datetime") , MkMap(&VarMap{})); _ = serverDatetime;
  return this.Parse8601(this.SafeString(serverDatetime , MkString("date") ));
}

func (this *Stex) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"currencyPairId":*(market ).At(MkString("id") )}); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit_bids") )= OpCopy(limit );
    *(request ).At(MkString("limit_asks") )= OpCopy(limit );
  }
  response := this.Call(MkString("publicGetOrderbookCurrencyPairId"), this.Extend(request , params )); _ = response;
  orderbook := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = orderbook;
  return this.ParseOrderBook(orderbook , symbol , MkUndefined() , MkString("bid") , MkString("ask") , MkString("price") , MkString("amount") );
}

func (this *Stex) ParseTicker(goArgs ...*Variant) *Variant{
  ticker := GoGetArg(goArgs, 0, MkUndefined()); _ = ticker;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.SafeInteger(ticker , MkString("timestamp") ); _ = timestamp;
  marketId := this.SafeString2(ticker , MkString("id") , MkString("symbol") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market , MkString("_") ); _ = symbol;
  last := this.SafeNumber(ticker , MkString("last") ); _ = last;
  open := this.SafeNumber(ticker , MkString("open") ); _ = open;
  change := OpCopy(MkUndefined() ); _ = change;
  percentage := OpCopy(MkUndefined() ); _ = percentage;
  if IsTrue(OpNotEq2(last , MkUndefined() )) {
    if IsTrue(OpAnd(OpNotEq2(open , MkUndefined() ), OpGt(open , MkInteger(0) ))) {
      change = OpSub(last , open );
      percentage = OpSub(OpMulti(OpDiv(MkInteger(100) , open ), last ), MkInteger(100) );
    }
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
        "open":open ,
        "close":last ,
        "last":last ,
        "previousClose":MkUndefined() ,
        "change":change ,
        "percentage":percentage ,
        "average":MkUndefined() ,
        "baseVolume":this.SafeNumber(ticker , MkString("volumeQuote") ),
        "quoteVolume":this.SafeNumber(ticker , MkString("volume") ),
        "info":ticker ,
        });
}

func (this *Stex) FetchTickers(goArgs ...*Variant) *Variant{
  symbols := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbols;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("publicGetTicker"), params ); _ = response;
  tickers := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = tickers;
  return this.ParseTickers(tickers , symbols );
}

func (this *Stex) ParseOHLCV(goArgs ...*Variant) *Variant{
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

func (this *Stex) FetchOHLCV(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  timeframe := GoGetArg(goArgs, 1, MkString("1d") ); _ = timeframe;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "currencyPairId":*(market ).At(MkString("id") ),
        "candlesType":*(*this.At(MkString("timeframes")) ).At(timeframe ),
        }); _ = request;
  if IsTrue(OpEq2(limit , MkUndefined() )) {
    limit = MkInteger(100) ;
  } else {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  duration := this.ParseTimeframe(timeframe ); _ = duration;
  timerange := OpMulti(limit , duration ); _ = timerange;
  if IsTrue(OpEq2(since , MkUndefined() )) {
    *(request ).At(MkString("timeEnd") )= this.Seconds();
    *(request ).At(MkString("timeStart") )= OpSub(*(request ).At(MkString("timeEnd") ), timerange );
  } else {
    *(request ).At(MkString("timeStart") )= parseInt(OpDiv(since , MkInteger(1000) ));
    *(request ).At(MkString("timeEnd") )= this.Sum(*(request ).At(MkString("timeStart") ), timerange );
  }
  response := this.Call(MkString("publicGetChartCurrencyPairIdCandlesType"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = data;
  return this.ParseOHLCVs(data , market , timeframe , since , limit );
}

func (this *Stex) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  id := this.SafeString(trade , MkString("id") ); _ = id;
  timestamp := this.SafeTimestamp(trade , MkString("timestamp") ); _ = timestamp;
  priceString := this.SafeString(trade , MkString("price") ); _ = priceString;
  amountString := this.SafeString(trade , MkString("amount") ); _ = amountString;
  price := this.ParseNumber(priceString ); _ = price;
  amount := this.ParseNumber(amountString ); _ = amount;
  cost := this.ParseNumber(Precise.StringMul(priceString , amountString )); _ = cost;
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  if IsTrue(OpAnd(OpEq2(symbol , MkUndefined() ), OpNotEq2(market , MkUndefined() ))) {
    symbol = *(market ).At(MkString("symbol") );
  }
  side := this.SafeStringLower2(trade , MkString("type") , MkString("trade_type") ); _ = side;
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

func (this *Stex) FetchTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"currencyPairId":*(market ).At(MkString("id") )}); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("sort") )= MkString("ASC") ;
    *(request ).At(MkString("from") )= parseInt(OpDiv(since , MkInteger(1000) ));
  }
  response := this.Call(MkString("publicGetTradesCurrencyPairId"), this.Extend(request , params )); _ = response;
  trades := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = trades;
  return this.ParseTrades(trades , market , since , limit );
}

func (this *Stex) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("profileGetWallets"), params ); _ = response;
  result := MkMap(&VarMap{
        "info":response ,
        "timestamp":MkUndefined() ,
        "datetime":MkUndefined() ,
        }); _ = result;
  balances := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = balances;
  for i := MkInteger(0) ; IsTrue(OpLw(i , balances.Length )); OpIncr(& i ){
    balance := *(balances ).At(i ); _ = balance;
    code := this.SafeCurrencyCode(this.SafeString(balance , MkString("currency_id") )); _ = code;
    account := this.Account(); _ = account;
    *(account ).At(MkString("free") )= this.SafeString(balance , MkString("balance") );
    *(account ).At(MkString("used") )= this.SafeString(balance , MkString("frozen_balance") );
    *(result ).At(code )= OpCopy(account );
  }
  return this.ParseBalance(result );
}

func (this *Stex) ParseOrderStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "PROCESSING":MkString("open") ,
        "PENDING":MkString("open") ,
        "PARTIAL":MkString("open") ,
        "FINISHED":MkString("closed") ,
        "CANCELLED":MkString("canceled") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Stex) ParseOrder(goArgs ...*Variant) *Variant{
  order := GoGetArg(goArgs, 0, MkUndefined()); _ = order;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  id := this.SafeString(order , MkString("id") ); _ = id;
  status := this.ParseOrderStatus(this.SafeString(order , MkString("status") )); _ = status;
  marketId := this.SafeString2(order , MkString("currency_pair_id") , MkString("currency_pair_name") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market , MkString("_") ); _ = symbol;
  timestamp := this.SafeTimestamp(order , MkString("timestamp") ); _ = timestamp;
  price := this.SafeNumber(order , MkString("price") ); _ = price;
  amount := this.SafeNumber(order , MkString("initial_amount") ); _ = amount;
  filled := this.SafeNumber(order , MkString("processed_amount") ); _ = filled;
  remaining := OpCopy(MkUndefined() ); _ = remaining;
  cost := OpCopy(MkUndefined() ); _ = cost;
  if IsTrue(OpNotEq2(filled , MkUndefined() )) {
    if IsTrue(OpNotEq2(amount , MkUndefined() )) {
      remaining = OpSub(amount , filled );
      if IsTrue(*(*this.At(MkString("options")) ).At(MkString("parseOrderToPrecision") )) {
        remaining = parseFloat(this.AmountToPrecision(symbol , remaining ));
      }
      remaining = Math.Max(remaining , MkNumber(0.0) );
    }
    if IsTrue(OpNotEq2(price , MkUndefined() )) {
      if IsTrue(OpEq2(cost , MkUndefined() )) {
        cost = OpMulti(price , filled );
      }
    }
  }
  type_ := this.SafeString(order , MkString("original_type") ); _ = type_;
  if IsTrue(OpOr(OpEq2(type_ , MkString("BUY") ), OpEq2(type_ , MkString("SELL") ))) {
    type_ = OpCopy(MkUndefined() );
  }
  side := this.SafeStringLower(order , MkString("type") ); _ = side;
  rawTrades := this.SafeValue(order , MkString("trades") ); _ = rawTrades;
  trades := OpCopy(MkUndefined() ); _ = trades;
  if IsTrue(OpNotEq2(rawTrades , MkUndefined() )) {
    trades = this.ParseTrades(rawTrades , market , MkUndefined() , MkUndefined() , MkMap(&VarMap{
            "symbol":symbol ,
            "order":id ,
            }));
  }
  stopPrice := this.SafeNumber(order , MkString("trigger_price") ); _ = stopPrice;
  result := MkMap(&VarMap{
        "info":order ,
        "id":id ,
        "clientOrderId":MkUndefined() ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "lastTradeTimestamp":MkUndefined() ,
        "symbol":symbol ,
        "type":type_ ,
        "timeInForce":MkUndefined() ,
        "postOnly":MkUndefined() ,
        "side":side ,
        "price":price ,
        "stopPrice":stopPrice ,
        "amount":amount ,
        "cost":cost ,
        "average":MkUndefined() ,
        "filled":filled ,
        "remaining":remaining ,
        "status":status ,
        "trades":trades ,
        }); _ = result;
  fees := this.SafeValue(order , MkString("fees") ); _ = fees;
  if IsTrue(OpEq2(fees , MkUndefined() )) {
    *(result ).At(MkString("fee") )= OpCopy(MkUndefined() );
  } else {
    numFees := OpCopy(fees.Length ); _ = numFees;
    if IsTrue(OpGt(numFees , MkInteger(0) )) {
      *(result ).At(MkString("fees") )= MkArray(&VarArray{});
      for i := MkInteger(0) ; IsTrue(OpLw(i , fees.Length )); OpIncr(& i ){
        feeCost := this.SafeNumber(*(fees ).At(i ), MkString("amount") ); _ = feeCost;
        if IsTrue(OpNotEq2(feeCost , MkUndefined() )) {
          feeCurrencyId := this.SafeString(*(fees ).At(i ), MkString("currency_id") ); _ = feeCurrencyId;
          feeCurrencyCode := this.SafeCurrencyCode(feeCurrencyId ); _ = feeCurrencyCode;
          (*(result ).At(MkString("fees") )).Call(MkString("push") , MkMap(&VarMap{
                  "cost":feeCost ,
                  "currency":feeCurrencyCode ,
                  }));
        }
      }
    } else {
      *(result ).At(MkString("fee") )= OpCopy(MkUndefined() );
    }
  }
  return result ;
}

func (this *Stex) CreateOrder(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 2, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 3, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 4, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 5, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(type_ , MkString("market") )) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , MkString(" createOrder allows limit orders only") )));
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  if IsTrue(OpEq2(type_ , MkString("limit") )) {
    type_ = OpCopy(side );
  }
  request := MkMap(&VarMap{
        "currencyPairId":*(market ).At(MkString("id") ),
        "type":type_.ToUpperCase(),
        "amount":parseFloat(this.AmountToPrecision(symbol , amount )),
        "price":parseFloat(this.PriceToPrecision(symbol , price )),
        }); _ = request;
  response := this.Call(MkString("tradingPostOrdersCurrencyPairId"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  return this.ParseOrder(data , market );
}

func (this *Stex) FetchOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"orderId":id }); _ = request;
  response := this.Call(MkString("tradingGetOrderOrderId"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
  }
  return this.ParseOrder(data , market );
}

func (this *Stex) FetchClosedOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"orderId":id }); _ = request;
  response := this.Call(MkString("reportsGetOrdersOrderId"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
  }
  return this.ParseOrder(data , market );
}

func (this *Stex) FetchOrderTrades(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  order := this.FetchClosedOrder(id , symbol , params ); _ = order;
  return *(order ).At(MkString("trades") );
}

func (this *Stex) FetchOpenOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := OpCopy(MkUndefined() ); _ = market;
  method := MkString("tradingGetOrders") ; _ = method;
  request := MkMap(&VarMap{}); _ = request;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    method = MkString("tradingGetOrdersCurrencyPairId") ;
    market = this.Market(symbol );
    *(request ).At(MkString("currencyPairId") )= *(market ).At(MkString("id") );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(method , this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = data;
  return this.ParseOrders(data , market , since , limit );
}

func (this *Stex) CancelOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"orderId":id }); _ = request;
  response := this.Call(MkString("tradingDeleteOrderOrderId"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  acceptedOrders := this.SafeValue(data , MkString("put_into_processing_queue") , MkArray(&VarArray{})); _ = acceptedOrders;
  rejectedOrders := this.SafeValue(data , MkString("not_put_into_processing_queue") , MkArray(&VarArray{})); _ = rejectedOrders;
  numAcceptedOrders := OpCopy(acceptedOrders.Length ); _ = numAcceptedOrders;
  numRejectedOrders := OpCopy(rejectedOrders.Length ); _ = numRejectedOrders;
  if IsTrue(OpLw(numAcceptedOrders , MkInteger(1) )) {
    if IsTrue(OpLw(numRejectedOrders , MkInteger(1) )) {
      panic(NewOrderNotFound(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" cancelOrder received an empty response: ") , this.Json(response )))));
    } else {
      return this.ParseOrder(*(rejectedOrders ).At(MkInteger(0) ));
    }
  } else {
    if IsTrue(OpLw(numRejectedOrders , MkInteger(1) )) {
      return this.ParseOrder(*(acceptedOrders ).At(MkInteger(0) ));
    } else {
      panic(NewOrderNotFound(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" cancelOrder received an empty response: ") , this.Json(response )))));
    }
  }
  return MkUndefined()
}

func (this *Stex) CancelAllOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  method := MkString("tradingDeleteOrders") ; _ = method;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market := this.Market(symbol ); _ = market;
    *(request ).At(MkString("currencyPairId") )= *(market ).At(MkString("id") );
    method = MkString("tradingDeleteOrdersCurrencyPairId") ;
  }
  response := this.Call(method , this.Extend(request , params )); _ = response;
  return response ;
}

func (this *Stex) FetchMyTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchMyTrades() requires a symbol argument") )));
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"currencyPairId":*(market ).At(MkString("id") )}); _ = request;
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("timeStart") )= this.Iso8601(since );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("reportsGetTradesCurrencyPairId"), this.Extend(request , params )); _ = response;
  trades := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = trades;
  return this.ParseTrades(trades , market , since , limit );
}

func (this *Stex) CreateDepositAddress(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  currency := this.Currency(code ); _ = currency;
  request := MkMap(&VarMap{"currencyId":*(currency ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("profilePostWalletsCurrencyId"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  depositAddress := this.SafeValue(data , MkString("deposit_address") , MkMap(&VarMap{})); _ = depositAddress;
  address := this.SafeString(depositAddress , MkString("address") ); _ = address;
  tag := this.SafeString(depositAddress , MkString("additional_address_parameter") ); _ = tag;
  this.CheckAddress(address );
  return MkMap(&VarMap{
        "currency":code ,
        "address":address ,
        "tag":tag ,
        "info":response ,
        });
}

func (this *Stex) FetchDepositAddress(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  balance := this.FetchBalance(); _ = balance;
  wallets := this.SafeValue(*(balance ).At(MkString("info") ), MkString("data") , MkArray(&VarArray{})); _ = wallets;
  walletsByCurrencyId := this.IndexBy(wallets , MkString("currency_id") ); _ = walletsByCurrencyId;
  currency := this.Currency(code ); _ = currency;
  wallet := this.SafeValue(walletsByCurrencyId , *(currency ).At(MkString("id") )); _ = wallet;
  if IsTrue(OpEq2(wallet , MkUndefined() )) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" fetchDepositAddress() could not find the wallet id for currency code ") , OpAdd(code , MkString(", try to call createDepositAddress() first") )))));
  }
  walletId := this.SafeInteger(wallet , MkString("id") ); _ = walletId;
  if IsTrue(OpEq2(walletId , MkUndefined() )) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" fetchDepositAddress() could not find the wallet id for currency code ") , OpAdd(code , MkString(", try to call createDepositAddress() first") )))));
  }
  request := MkMap(&VarMap{"walletId":walletId }); _ = request;
  response := this.Call(MkString("profileGetWalletsWalletId"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = data;
  depositAddress := this.SafeValue(data , MkString("deposit_address") , MkMap(&VarMap{})); _ = depositAddress;
  address := this.SafeString(depositAddress , MkString("address") ); _ = address;
  tag := this.SafeString(depositAddress , MkString("additional_address_parameter") ); _ = tag;
  this.CheckAddress(address );
  return MkMap(&VarMap{
        "currency":code ,
        "address":address ,
        "tag":tag ,
        "info":response ,
        });
}

func (this *Stex) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("public") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  url := OpAdd(*(*this.At(MkString("urls")) ).At(MkString("api") ), OpAdd(MkString("/") , OpAdd(api , OpAdd(MkString("/") , this.ImplodeParams(path , params ))))); _ = url;
  query := this.Omit(params , this.ExtractParams(path )); _ = query;
  if IsTrue(OpEq2(api , MkString("public") )) {
    if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
      OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(query )));
    }
  } else {
    this.CheckRequiredCredentials();
    headers = MkMap(&VarMap{"Authorization":OpAdd(MkString("Bearer ") , *this.At(MkString("token")) )});
    if IsTrue(OpOr(OpEq2(method , MkString("GET") ), OpEq2(method , MkString("DELETE") ))) {
      if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
        OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(query )));
      }
    } else {
      body = this.Json(query );
      if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
        *(headers ).At(MkString("Content-Type") )= MkString("application/json") ;
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

func (this *Stex) ParseTransactionStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "processing":MkString("pending") ,
        "checking by system":MkString("pending") ,
        "hodl":MkString("pending") ,
        "amount too low":MkString("failed") ,
        "not confirmed":MkString("pending") ,
        "cancelled by user":MkString("canceled") ,
        "approved":MkString("pending") ,
        "finished":MkString("ok") ,
        "withdrawal error":MkString("failed") ,
        "deposit error":MkString("failed") ,
        "cancelled by admin":MkString("canceled") ,
        "awaiting":MkString("pending") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Stex) ParseTransaction(goArgs ...*Variant) *Variant{
  transaction := GoGetArg(goArgs, 0, MkUndefined()); _ = transaction;
  currency := GoGetArg(goArgs, 1, MkUndefined() ); _ = currency;
  id := this.SafeString(transaction , MkString("id") ); _ = id;
  withdrawalAddress := this.SafeValue(transaction , MkString("withdrawal_address") , MkMap(&VarMap{})); _ = withdrawalAddress;
  address := this.SafeString(withdrawalAddress , MkString("address") ); _ = address;
  tag := this.SafeString(withdrawalAddress , MkString("additional_address_parameter") ); _ = tag;
  currencyId := this.SafeString(transaction , MkString("currency_id") ); _ = currencyId;
  code := OpCopy(MkUndefined() ); _ = code;
  if IsTrue(OpHasMember(currencyId , *this.At(MkString("currencies_by_id")) )) {
    currency = *(*this.At(MkString("currencies_by_id")) ).At(currencyId );
  } else {
    code = this.CommonCurrencyCode(this.SafeString(transaction , MkString("currency_code") ));
  }
  if IsTrue(OpAnd(OpEq2(code , MkUndefined() ), OpNotEq2(currency , MkUndefined() ))) {
    code = *(currency ).At(MkString("code") );
  }
  type_ := OpTrinary(OpHasMember(MkString("deposit_status_id") , transaction ), MkString("deposit") , MkString("withdrawal") ); _ = type_;
  amount := this.SafeNumber(transaction , MkString("amount") ); _ = amount;
  status := this.ParseTransactionStatus(this.SafeStringLower(transaction , MkString("status") )); _ = status;
  timestamp := this.SafeTimestamp2(transaction , MkString("timestamp") , MkString("created_ts") ); _ = timestamp;
  updated := this.SafeTimestamp(transaction , MkString("updated_ts") ); _ = updated;
  txid := this.SafeString(transaction , MkString("txid") ); _ = txid;
  fee := OpCopy(MkUndefined() ); _ = fee;
  feeCost := this.SafeNumber(transaction , MkString("fee") ); _ = feeCost;
  if IsTrue(OpNotEq2(feeCost , MkUndefined() )) {
    feeCurrencyId := this.SafeString(transaction , MkString("fee_currency_id") , MkString("deposit_fee_currency_id") ); _ = feeCurrencyId;
    feeCurrencyCode := this.SafeCurrencyCode(feeCurrencyId ); _ = feeCurrencyCode;
    fee = MkMap(&VarMap{
        "cost":feeCost ,
        "currency":feeCurrencyCode ,
        });
  }
  return MkMap(&VarMap{
        "info":transaction ,
        "id":id ,
        "txid":txid ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "addressFrom":MkUndefined() ,
        "address":address ,
        "addressTo":address ,
        "tagFrom":MkUndefined() ,
        "tag":tag ,
        "tagTo":tag ,
        "type":type_ ,
        "amount":amount ,
        "currency":code ,
        "status":status ,
        "updated":updated ,
        "fee":fee ,
        });
}

func (this *Stex) FetchDeposits(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  currency := OpCopy(MkUndefined() ); _ = currency;
  request := MkMap(&VarMap{}); _ = request;
  if IsTrue(OpNotEq2(code , MkUndefined() )) {
    currency = this.Currency(code );
    *(request ).At(MkString("currencyId") )= *(currency ).At(MkString("id") );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("timeStart") )= OpCopy(since );
  }
  response := this.Call(MkString("profileGetDeposits"), this.Extend(request , params )); _ = response;
  deposits := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = deposits;
  return this.ParseTransactions(deposits , code , since , limit );
}

func (this *Stex) FetchWithdrawals(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  currency := OpCopy(MkUndefined() ); _ = currency;
  request := MkMap(&VarMap{}); _ = request;
  if IsTrue(OpNotEq2(code , MkUndefined() )) {
    currency = this.Currency(code );
    *(request ).At(MkString("currencyId") )= *(currency ).At(MkString("id") );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("timeStart") )= OpCopy(since );
  }
  response := this.Call(MkString("profileGetWithdrawals"), this.Extend(request , params )); _ = response;
  withdrawals := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = withdrawals;
  return this.ParseTransactions(withdrawals , code , since , limit );
}

func (this *Stex) Withdraw(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  amount := GoGetArg(goArgs, 1, MkUndefined()); _ = amount;
  address := GoGetArg(goArgs, 2, MkUndefined()); _ = address;
  tag := GoGetArg(goArgs, 3, MkUndefined() ); _ = tag;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.CheckAddress(address );
  this.LoadMarkets();
  currency := this.Currency(code ); _ = currency;
  request := MkMap(&VarMap{
        "currency_id":*(currency ).At(MkString("id") ),
        "amount":parseFloat(this.CurrencyToPrecision(code , amount )),
        "address":address ,
        }); _ = request;
  if IsTrue(OpNotEq2(tag , MkUndefined() )) {
    *(request ).At(MkString("additional_address_parameter") )= OpCopy(tag );
  }
  response := this.Call(MkString("profilePostWithdraw"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  return this.ParseTransaction(data , currency );
}

func (this *Stex) FetchFundingFees(goArgs ...*Variant) *Variant{
  codes := GoGetArg(goArgs, 0, MkUndefined() ); _ = codes;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("publicGetCurrencies"), params ); _ = response;
  data := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = data;
  withdrawFees := MkMap(&VarMap{}); _ = withdrawFees;
  depositFees := MkMap(&VarMap{}); _ = depositFees;
  for i := MkInteger(0) ; IsTrue(OpLw(i , data.Length )); OpIncr(& i ){
    id := this.SafeString(*(data ).At(i ), MkString("id") ); _ = id;
    code := this.SafeCurrencyCode(id ); _ = code;
    *(withdrawFees ).At(code )= this.SafeNumber(*(data ).At(i ), MkString("withdrawal_fee_const") );
    *(depositFees ).At(code )= this.SafeNumber(*(data ).At(i ), MkString("deposit_fee_const") );
  }
  return MkMap(&VarMap{
        "withdraw":withdrawFees ,
        "deposit":depositFees ,
        "info":response ,
        });
}

func (this *Stex) HandleErrors(goArgs ...*Variant) *Variant{
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
  success := this.SafeValue(response , MkString("success") , MkBool(false) ); _ = success;
  if IsTrue(OpNot(success )) {
    message := this.SafeString(response , MkString("message") ); _ = message;
    feedback := OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , body )); _ = feedback;
    this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("exact") ), message , feedback );
    this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("broad") ), message , feedback );
    panic(NewExchangeError(feedback ));
  }
  return MkUndefined()
}

