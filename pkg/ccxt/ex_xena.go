package ccxt

import (
)

type Xena struct {
	*ExchangeBase
}

var _ Exchange = (*Xena)(nil)

func init() {
	exchange := &Xena{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Xena) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("xena") ,
            "name":MkString("Xena Exchange") ,
            "countries":MkArray(&VarArray{
                MkString("VC") ,
                MkString("UK") ,
                }),
            "rateLimit":MkInteger(100) ,
            "has":MkMap(&VarMap{
                "CORS":MkBool(false) ,
                "cancelAllOrders":MkBool(true) ,
                "cancelOrder":MkBool(true) ,
                "createDepositAddress":MkBool(true) ,
                "createOrder":MkBool(true) ,
                "editOrder":MkBool(true) ,
                "fetchBalance":MkBool(true) ,
                "fetchClosedOrders":MkBool(true) ,
                "fetchCurrencies":MkBool(true) ,
                "fetchDepositAddress":MkBool(true) ,
                "fetchDeposits":MkBool(true) ,
                "fetchLedger":MkBool(true) ,
                "fetchMarkets":MkBool(true) ,
                "fetchMyTrades":MkBool(true) ,
                "fetchOHLCV":MkBool(true) ,
                "fetchOpenOrders":MkBool(true) ,
                "fetchOrderBook":MkBool(true) ,
                "fetchTicker":MkBool(true) ,
                "fetchTickers":MkBool(true) ,
                "fetchTime":MkBool(true) ,
                "fetchTrades":MkBool(true) ,
                "fetchWithdrawals":MkBool(true) ,
                "withdraw":MkBool(true) ,
                }),
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/51840849/87489843-bb469280-c64c-11ea-91aa-69c6326506af.jpg") ,
                "test":MkMap(&VarMap{
                    "public":MkString("https://trading.demo.xena.io/api") ,
                    "private":MkString("https://api.demo.xena.io") ,
                    }),
                "api":MkMap(&VarMap{
                    "public":MkString("https://trading.xena.exchange/api") ,
                    "private":MkString("https://api.xena.exchange") ,
                    }),
                "www":MkString("https://xena.exchange") ,
                "doc":MkString("https://support.xena.exchange/support/solutions/44000808700") ,
                "fees":MkString("https://trading.xena.exchange/en/platform-specification/fee-schedule") ,
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
                        MkString("common/currencies") ,
                        MkString("common/instruments") ,
                        MkString("common/features") ,
                        MkString("common/commissions") ,
                        MkString("common/news") ,
                        MkString("market-data/candles/{marketId}/{timeframe}") ,
                        MkString("market-data/market-watch") ,
                        MkString("market-data/dom/{symbol}") ,
                        MkString("market-data/candles/{symbol}/{timeframe}") ,
                        MkString("market-data/trades/{symbol}") ,
                        MkString("market-data/server-time") ,
                        MkString("market-data/v2/candles/{symbol}/{timeframe}") ,
                        MkString("market-data/v2/trades/{symbol}") ,
                        MkString("market-data/v2/dom/{symbol}/") ,
                        MkString("market-data/v2/server-time") ,
                        })}),
                "private":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("trading/accounts/{accountId}/order") ,
                        MkString("trading/accounts/{accountId}/active-orders") ,
                        MkString("trading/accounts/{accountId}/last-order-statuses") ,
                        MkString("trading/accounts/{accountId}/positions") ,
                        MkString("trading/accounts/{accountId}/positions-history") ,
                        MkString("trading/accounts/{accountId}/margin-requirements") ,
                        MkString("trading/accounts") ,
                        MkString("trading/accounts/{accountId}/balance") ,
                        MkString("trading/accounts/{accountId}/trade-history") ,
                        MkString("transfers/accounts") ,
                        MkString("transfers/accounts/{accountId}") ,
                        MkString("transfers/accounts/{accountId}/deposit-address/{currency}") ,
                        MkString("transfers/accounts/{accountId}/deposits") ,
                        MkString("transfers/accounts/{accountId}/trusted-addresses") ,
                        MkString("transfers/accounts/{accountId}/withdrawals") ,
                        MkString("transfers/accounts/{accountId}/balance-history") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("trading/order/new") ,
                        MkString("trading/order/heartbeat") ,
                        MkString("trading/order/cancel") ,
                        MkString("trading/order/mass-cancel") ,
                        MkString("trading/order/replace") ,
                        MkString("trading/position/maintenance") ,
                        MkString("transfers/accounts/{accountId}/withdrawals") ,
                        MkString("transfers/accounts/{accountId}/deposit-address/{currency}") ,
                        }),
                    }),
                }),
            "fees":MkMap(&VarMap{
                "trading":MkMap(&VarMap{
                    "maker":MkNumber(0.0005) ,
                    "taker":MkNumber(0.001) ,
                    "tierBased":MkBool(true) ,
                    "percentage":MkBool(true) ,
                    }),
                "funding":MkMap(&VarMap{
                    "tierBased":MkBool(false) ,
                    "percentage":MkBool(false) ,
                    "withdraw":MkMap(&VarMap{}),
                    "deposit":MkMap(&VarMap{}),
                    }),
                }),
            "exceptions":MkMap(&VarMap{
                "exact":MkMap(&VarMap{
                    "Validation failed":BadRequest ,
                    "Unknown derivative symbol":BadSymbol ,
                    "Unknown account":BadRequest ,
                    "Wrong TransactTime":BadRequest ,
                    "ClOrdId is empty":BadRequest ,
                    }),
                "broad":MkMap(&VarMap{
                    "Invalid aggregation ratio or depth":BadRequest ,
                    "address":InvalidAddress ,
                    "Money not enough":InsufficientFunds ,
                    "parse error":BadRequest ,
                    "Not enough":InsufficientFunds ,
                    }),
                }),
            "options":MkMap(&VarMap{
                "defaultType":MkString("margin") ,
                "accountId":MkUndefined() ,
                }),
            }));
}

func (this *Xena) FetchTime(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetMarketDataV2ServerTime"), params ); _ = response;
  transactTime := this.SafeInteger(response , MkString("transactTime") ); _ = transactTime;
  return parseInt(OpDiv(transactTime , MkInteger(1000000) ));
}

func (this *Xena) FetchMarkets(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetCommonInstruments"), params ); _ = response;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , response.Length )); OpIncr(& i ){
    market := *(response ).At(i ); _ = market;
    type_ := this.SafeStringLower(market , MkString("type") ); _ = type_;
    id := this.SafeString(market , MkString("symbol") ); _ = id;
    numericId := this.SafeString(market , MkString("id") ); _ = numericId;
    marginType := this.SafeString(market , MkString("marginType") ); _ = marginType;
    baseId := this.SafeString(market , MkString("baseCurrency") ); _ = baseId;
    quoteId := this.SafeString(market , MkString("quoteCurrency") ); _ = quoteId;
    base := this.SafeCurrencyCode(baseId ); _ = base;
    quote := this.SafeCurrencyCode(quoteId ); _ = quote;
    symbol := OpCopy(id ); _ = symbol;
    if IsTrue(OpEq2(type_ , MkString("margin") )) {
      if IsTrue(OpEq2(marginType , MkString("XenaFuture") )) {
        type_ = MkString("future") ;
      } else {
        if IsTrue(OpEq2(marginType , MkString("XenaListedPerpetual") )) {
          type_ = MkString("swap") ;
          symbol = OpAdd(base , OpAdd(MkString("/") , quote ));
        }
      }
    }
    future := OpEq2(type_ , MkString("future") ); _ = future;
    swap := OpEq2(type_ , MkString("swap") ); _ = swap;
    pricePrecision := this.SafeInteger2(market , MkString("tickSize") , MkString("pricePrecision") ); _ = pricePrecision;
    precision := MkMap(&VarMap{
          "price":pricePrecision ,
          "amount":MkInteger(0) ,
          }); _ = precision;
    maxCost := this.SafeNumber(market , MkString("maxOrderQty") ); _ = maxCost;
    minCost := this.SafeNumber(market , MkString("minOrderQuantity") ); _ = minCost;
    limits := MkMap(&VarMap{
          "amount":MkMap(&VarMap{
              "min":MkUndefined() ,
              "max":MkUndefined() ,
              }),
          "price":MkMap(&VarMap{
              "min":MkUndefined() ,
              "max":MkUndefined() ,
              }),
          "cost":MkMap(&VarMap{
              "min":minCost ,
              "max":maxCost ,
              }),
          }); _ = limits;
    active := this.SafeValue(market , MkString("enabled") , MkBool(false) ); _ = active;
    inverse := this.SafeValue(market , MkString("inverse") , MkBool(false) ); _ = inverse;
    result.Push(MkMap(&VarMap{
            "id":id ,
            "symbol":symbol ,
            "base":base ,
            "quote":quote ,
            "baseId":baseId ,
            "quoteId":quoteId ,
            "numericId":numericId ,
            "active":active ,
            "type":type_ ,
            "spot":MkBool(false) ,
            "future":future ,
            "swap":swap ,
            "inverse":inverse ,
            "precision":precision ,
            "limits":limits ,
            "info":market ,
            }));
  }
  return result ;
}

func (this *Xena) FetchCurrencies(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetCommonCurrencies"), params ); _ = response;
  ids := GoGetKeys(response ); _ = ids;
  result := MkMap(&VarMap{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , ids.Length )); OpIncr(& i ){
    id := *(ids ).At(i ); _ = id;
    currency := *(response ).At(id ); _ = currency;
    code := this.SafeCurrencyCode(id ); _ = code;
    name := this.SafeString(currency , MkString("title") ); _ = name;
    precision := this.SafeInteger(currency , MkString("precision") ); _ = precision;
    enabled := this.SafeValue(currency , MkString("enabled") ); _ = enabled;
    active := OpEq2(enabled , MkBool(true) ); _ = active;
    withdraw := this.SafeValue(currency , MkString("withdraw") , MkMap(&VarMap{})); _ = withdraw;
    *(result ).At(code )= MkMap(&VarMap{
        "id":id ,
        "code":code ,
        "info":currency ,
        "name":name ,
        "active":active ,
        "fee":this.SafeNumber(withdraw , MkString("commission") ),
        "precision":precision ,
        "limits":MkMap(&VarMap{
            "amount":MkMap(&VarMap{
                "min":MkUndefined() ,
                "max":MkUndefined() ,
                }),
            "withdraw":MkMap(&VarMap{
                "min":this.SafeNumber(withdraw , MkString("minAmount") ),
                "max":MkUndefined() ,
                }),
            }),
        });
  }
  return result ;
}

func (this *Xena) ParseTicker(goArgs ...*Variant) *Variant{
  ticker := GoGetArg(goArgs, 0, MkUndefined()); _ = ticker;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.Milliseconds(); _ = timestamp;
  marketId := this.SafeString(ticker , MkString("symbol") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market ); _ = symbol;
  last := this.SafeNumber(ticker , MkString("lastPx") ); _ = last;
  open := this.SafeNumber(ticker , MkString("firstPx") ); _ = open;
  percentage := OpCopy(MkUndefined() ); _ = percentage;
  change := OpCopy(MkUndefined() ); _ = change;
  average := OpCopy(MkUndefined() ); _ = average;
  if IsTrue(OpAnd(OpNotEq2(last , MkUndefined() ), OpNotEq2(open , MkUndefined() ))) {
    change = OpSub(last , open );
    average = OpDiv(this.Sum(last , open ), MkInteger(2) );
    if IsTrue(OpGt(open , MkInteger(0) )) {
      percentage = OpDiv(change , OpMulti(open , MkInteger(100) ));
    }
  }
  buyVolume := this.SafeNumber(ticker , MkString("buyVolume") ); _ = buyVolume;
  sellVolume := this.SafeNumber(ticker , MkString("sellVolume") ); _ = sellVolume;
  baseVolume := this.Sum(buyVolume , sellVolume ); _ = baseVolume;
  return MkMap(&VarMap{
        "symbol":symbol ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "high":this.SafeNumber(ticker , MkString("highPx") ),
        "low":this.SafeNumber(ticker , MkString("lowPx") ),
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
        "average":average ,
        "baseVolume":baseVolume ,
        "quoteVolume":MkUndefined() ,
        "info":ticker ,
        });
}

func (this *Xena) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  tickers := this.FetchTickers(MkUndefined() , params ); _ = tickers;
  if IsTrue(OpHasMember(symbol , tickers )) {
    return *(tickers ).At(symbol );
  }
  panic(NewBadSymbol(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" fetchTicker could not find a ticker with symbol ") , symbol ))));
  return MkUndefined()
}

func (this *Xena) FetchTickers(goArgs ...*Variant) *Variant{
  symbols := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbols;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  tickers := this.Call(MkString("publicGetMarketDataMarketWatch"), params ); _ = tickers;
  result := MkMap(&VarMap{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , tickers.Length )); OpIncr(& i ){
    ticker := this.ParseTicker(*(tickers ).At(i )); _ = ticker;
    symbol := *(ticker ).At(MkString("symbol") ); _ = symbol;
    *(result ).At(symbol )= OpCopy(ticker );
  }
  return this.FilterByArray(result , MkString("symbol") , symbols );
}

func (this *Xena) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"symbol":this.MarketId(symbol )}); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("depth") )= OpCopy(limit );
  }
  response := this.Call(MkString("publicGetMarketDataV2DomSymbol"), this.Extend(request , params )); _ = response;
  mdEntry := this.SafeValue(response , MkString("mdEntry") , MkArray(&VarArray{})); _ = mdEntry;
  mdEntriesByType := this.GroupBy(mdEntry , MkString("mdEntryType") ); _ = mdEntriesByType;
  lastUpdateTime := this.SafeInteger(response , MkString("lastUpdateTime") ); _ = lastUpdateTime;
  timestamp := parseInt(OpDiv(lastUpdateTime , MkInteger(1000000) )); _ = timestamp;
  return this.ParseOrderBook(mdEntriesByType , symbol , timestamp , MkString("0") , MkString("1") , MkString("mdEntryPx") , MkString("mdEntrySize") );
}

func (this *Xena) FetchAccounts(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("privateGetTradingAccounts"), params ); _ = response;
  accounts := this.SafeValue(response , MkString("accounts") ); _ = accounts;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , accounts.Length )); OpIncr(& i ){
    account := *(accounts ).At(i ); _ = account;
    accountId := this.SafeString(account , MkString("id") ); _ = accountId;
    currencyId := this.SafeString(account , MkString("currency") ); _ = currencyId;
    code := this.SafeCurrencyCode(currencyId ); _ = code;
    type_ := this.SafeStringLower(account , MkString("kind") ); _ = type_;
    result.Push(MkMap(&VarMap{
            "id":accountId ,
            "type":type_ ,
            "currency":code ,
            "info":account ,
            }));
  }
  return result ;
}

func (this *Xena) FindAccountByType(goArgs ...*Variant) *Variant{
  type_ := GoGetArg(goArgs, 0, MkUndefined()); _ = type_;
  this.LoadMarkets();
  this.LoadAccounts();
  accountsByType := this.GroupBy(*this.At(MkString("accounts")) , MkString("type") ); _ = accountsByType;
  accounts := this.SafeValue(accountsByType , type_ ); _ = accounts;
  if IsTrue(OpEq2(accounts , MkUndefined() )) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" findAccountByType() could not find an accountId with type '") , OpAdd(type_ , MkString("', specify the 'accountId' parameter instead") )))));
  }
  numAccounts := OpCopy(accounts.Length ); _ = numAccounts;
  if IsTrue(OpGt(numAccounts , MkInteger(1) )) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" findAccountByType() found more than one accountId with type '") , OpAdd(type_ , MkString("', specify the 'accountId' parameter instead") )))));
  }
  return *(accounts ).At(MkInteger(0) );
}

func (this *Xena) GetAccountId(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkUndefined()); _ = params;
  this.LoadMarkets();
  this.LoadAccounts();
  defaultAccountId := this.SafeString(*this.At(MkString("options")) , MkString("accountId") ); _ = defaultAccountId;
  accountId := this.SafeString(params , MkString("accountId") , defaultAccountId ); _ = accountId;
  if IsTrue(OpNotEq2(accountId , MkUndefined() )) {
    return accountId ;
  }
  defaultType := this.SafeString(*this.At(MkString("options")) , MkString("defaultType") , MkString("margin") ); _ = defaultType;
  type_ := this.SafeString(params , MkString("type") , defaultType ); _ = type_;
  params = this.Omit(params , MkString("type") );
  if IsTrue(OpEq2(type_ , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" requires an 'accountId' parameter or a 'type' parameter ('spot' or 'margin')") )));
  }
  account := this.FindAccountByType(type_ ); _ = account;
  return *(account ).At(MkString("id") );
}

func (this *Xena) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  this.LoadAccounts();
  accountId := this.GetAccountId(params ); _ = accountId;
  request := MkMap(&VarMap{"accountId":accountId }); _ = request;
  response := this.Call(MkString("privateGetTradingAccountsAccountIdBalance"), this.Extend(request , params )); _ = response;
  result := MkMap(&VarMap{"info":response }); _ = result;
  timestamp := OpCopy(MkUndefined() ); _ = timestamp;
  balances := this.SafeValue(response , MkString("balances") , MkArray(&VarArray{})); _ = balances;
  for i := MkInteger(0) ; IsTrue(OpLw(i , balances.Length )); OpIncr(& i ){
    balance := *(balances ).At(i ); _ = balance;
    lastUpdateTime := this.SafeString(balance , MkString("lastUpdateTime") ); _ = lastUpdateTime;
    lastUpdated := lastUpdateTime.Slice(MkInteger(0) , MkInteger(13) ); _ = lastUpdated;
    currentTimestamp := parseInt(lastUpdated ); _ = currentTimestamp;
    timestamp = OpTrinary(OpEq2(timestamp , MkUndefined() ), currentTimestamp , Math.Max(timestamp , currentTimestamp ));
    currencyId := this.SafeString(balance , MkString("currency") ); _ = currencyId;
    code := this.SafeCurrencyCode(currencyId ); _ = code;
    account := this.Account(); _ = account;
    *(account ).At(MkString("free") )= this.SafeString(balance , MkString("available") );
    *(account ).At(MkString("used") )= this.SafeString(balance , MkString("onHold") );
    *(result ).At(code )= OpCopy(account );
  }
  *(result ).At(MkString("timestamp") )= OpCopy(timestamp );
  *(result ).At(MkString("datetime") )= this.Iso8601(timestamp );
  return this.ParseBalance(result );
}

func (this *Xena) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  id := this.SafeString(trade , MkString("tradeId") ); _ = id;
  timestamp := this.SafeInteger(trade , MkString("transactTime") ); _ = timestamp;
  if IsTrue(OpNotEq2(timestamp , MkUndefined() )) {
    timestamp = parseInt(OpDiv(timestamp , MkInteger(1000000) ));
  }
  side := this.SafeStringLower2(trade , MkString("side") , MkString("aggressorSide") ); _ = side;
  if IsTrue(OpEq2(side , MkString("1") )) {
    side = MkString("buy") ;
  } else {
    if IsTrue(OpEq2(side , MkString("2") )) {
      side = MkString("sell") ;
    }
  }
  orderId := this.SafeString(trade , MkString("orderId") ); _ = orderId;
  marketId := this.SafeString(trade , MkString("symbol") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market ); _ = symbol;
  priceString := this.SafeString2(trade , MkString("lastPx") , MkString("mdEntryPx") ); _ = priceString;
  amountString := this.SafeString2(trade , MkString("lastQty") , MkString("mdEntrySize") ); _ = amountString;
  price := this.ParseNumber(priceString ); _ = price;
  amount := this.ParseNumber(amountString ); _ = amount;
  cost := this.ParseNumber(Precise.StringMul(priceString , amountString )); _ = cost;
  fee := OpCopy(MkUndefined() ); _ = fee;
  feeCost := this.SafeNumber(trade , MkString("commission") ); _ = feeCost;
  if IsTrue(OpNotEq2(feeCost , MkUndefined() )) {
    feeCurrencyId := this.SafeString(trade , MkString("commCurrency") ); _ = feeCurrencyId;
    feeCurrencyCode := this.SafeCurrencyCode(feeCurrencyId ); _ = feeCurrencyCode;
    feeRate := this.SafeNumber(trade , MkString("commRate") ); _ = feeRate;
    fee = MkMap(&VarMap{
        "cost":feeCost ,
        "rate":feeRate ,
        "currency":feeCurrencyCode ,
        });
  }
  return MkMap(&VarMap{
        "id":id ,
        "info":trade ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "symbol":symbol ,
        "type":MkUndefined() ,
        "order":orderId ,
        "side":side ,
        "takerOrMaker":MkUndefined() ,
        "price":price ,
        "amount":amount ,
        "cost":cost ,
        "fee":fee ,
        });
}

func (this *Xena) FetchMyTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  this.LoadAccounts();
  accountId := this.GetAccountId(params ); _ = accountId;
  request := MkMap(&VarMap{"accountId":accountId }); _ = request;
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("symbol") )= *(market ).At(MkString("id") );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("from") )= OpMulti(since , MkInteger(1000000) );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("privateGetTradingAccountsAccountIdTradeHistory"), this.Extend(request , params )); _ = response;
  return this.ParseTrades(response , market , since , limit );
}

func (this *Xena) ParseOHLCV(goArgs ...*Variant) *Variant{
  ohlcv := GoGetArg(goArgs, 0, MkUndefined()); _ = ohlcv;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  transactTime := this.SafeInteger(ohlcv , MkString("transactTime") ); _ = transactTime;
  timestamp := parseInt(OpDiv(transactTime , MkInteger(1000000) )); _ = timestamp;
  buyVolume := this.SafeNumber(ohlcv , MkString("buyVolume") ); _ = buyVolume;
  sellVolume := this.SafeNumber(ohlcv , MkString("sellVolume") ); _ = sellVolume;
  volume := this.Sum(buyVolume , sellVolume ); _ = volume;
  return MkArray(&VarArray{
        timestamp ,
        this.SafeNumber(ohlcv , MkString("firstPx") ),
        this.SafeNumber(ohlcv , MkString("highPx") ),
        this.SafeNumber(ohlcv , MkString("lowPx") ),
        this.SafeNumber(ohlcv , MkString("lastPx") ),
        volume ,
        });
}

func (this *Xena) FetchOHLCV(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  timeframe := GoGetArg(goArgs, 1, MkString("1m") ); _ = timeframe;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "symbol":*(market ).At(MkString("id") ),
        "timeframe":*(*this.At(MkString("timeframes")) ).At(timeframe ),
        }); _ = request;
  durationInSeconds := this.ParseTimeframe(timeframe ); _ = durationInSeconds;
  duration := OpMulti(durationInSeconds , MkInteger(1000) ); _ = duration;
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("from") )= OpMulti(since , MkInteger(1000000) );
    if IsTrue(OpNotEq2(limit , MkUndefined() )) {
      *(request ).At(MkString("to") )= OpMulti(this.Sum(since , OpMulti(limit , duration )), MkInteger(1000000) );
    }
  } else {
    now := this.Milliseconds(); _ = now;
    if IsTrue(OpNotEq2(limit , MkUndefined() )) {
      *(request ).At(MkString("from") )= OpMulti(OpSub(now , OpMulti(limit , duration )), MkInteger(1000000) );
    }
  }
  response := this.Call(MkString("publicGetMarketDataV2CandlesSymbolTimeframe"), this.Extend(request , params )); _ = response;
  mdEntry := this.SafeValue(response , MkString("mdEntry") , MkArray(&VarArray{})); _ = mdEntry;
  return this.ParseOHLCVs(mdEntry , market , timeframe , since , limit );
}

func (this *Xena) FetchTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"symbol":*(market ).At(MkString("id") )}); _ = request;
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("from") )= this.Iso8601(since );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("publicGetMarketDataV2TradesSymbol"), this.Extend(request , params )); _ = response;
  mdEntry := this.SafeValue(response , MkString("mdEntry") , MkArray(&VarArray{})); _ = mdEntry;
  return this.ParseTrades(mdEntry , market , since , limit );
}

func (this *Xena) ParseOrderStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "A":MkString("open") ,
        "0":MkString("open") ,
        "1":MkString("open") ,
        "2":MkString("closed") ,
        "6":MkString("canceled") ,
        "4":MkString("canceled") ,
        "E":MkString("open") ,
        "8":MkString("rejected") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Xena) ParseOrder(goArgs ...*Variant) *Variant{
  order := GoGetArg(goArgs, 0, MkUndefined()); _ = order;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  id := this.SafeString(order , MkString("orderId") ); _ = id;
  clientOrderId := this.SafeString(order , MkString("clOrdId") ); _ = clientOrderId;
  transactTime := this.SafeInteger(order , MkString("transactTime") ); _ = transactTime;
  timestamp := parseInt(OpDiv(transactTime , MkInteger(1000000) )); _ = timestamp;
  status := this.ParseOrderStatus(this.SafeString(order , MkString("ordStatus") )); _ = status;
  marketId := this.SafeString(order , MkString("symbol") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market ); _ = symbol;
  price := this.SafeNumber(order , MkString("price") ); _ = price;
  amount := this.SafeNumber(order , MkString("orderQty") ); _ = amount;
  filled := this.SafeNumber(order , MkString("cumQty") ); _ = filled;
  remaining := this.SafeNumber(order , MkString("leavesQty") ); _ = remaining;
  side := this.SafeStringLower(order , MkString("side") ); _ = side;
  if IsTrue(OpEq2(side , MkString("1") )) {
    side = MkString("buy") ;
  } else {
    if IsTrue(OpEq2(side , MkString("1") )) {
      side = MkString("sell") ;
    }
  }
  type_ := this.SafeStringLower(order , MkString("ordType") ); _ = type_;
  if IsTrue(OpEq2(type_ , MkString("1") )) {
    type_ = MkString("market") ;
  } else {
    if IsTrue(OpEq2(type_ , MkString("2") )) {
      type_ = MkString("limit") ;
    } else {
      if IsTrue(OpEq2(type_ , MkString("3") )) {
        type_ = MkString("stop") ;
      } else {
        if IsTrue(OpEq2(type_ , MkString("4") )) {
          type_ = MkString("stop-limit") ;
        }
      }
    }
  }
  return this.SafeOrder(MkMap(&VarMap{
            "id":id ,
            "clientOrderId":clientOrderId ,
            "info":order ,
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
            "average":MkUndefined() ,
            "filled":filled ,
            "remaining":remaining ,
            "status":status ,
            "fee":MkUndefined() ,
            "trades":MkUndefined() ,
            }));
}

func (this *Xena) CreateOrder(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 2, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 3, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 4, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 5, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  this.LoadAccounts();
  accountId := this.GetAccountId(params ); _ = accountId;
  orderTypes := MkMap(&VarMap{
        "market":MkString("1") ,
        "limit":MkString("2") ,
        "stop":MkString("3") ,
        "stop-limit":MkString("4") ,
        }); _ = orderTypes;
  orderType := this.SafeString(orderTypes , type_ ); _ = orderType;
  if IsTrue(OpEq2(orderType , MkUndefined() )) {
    panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" createOrder does not support order type ") , OpAdd(type_ , MkString(", supported order types are market, limit, stop, stop-limit") )))));
  }
  orderSides := MkMap(&VarMap{
        "buy":MkString("1") ,
        "sell":MkString("2") ,
        }); _ = orderSides;
  orderSide := this.SafeString(orderSides , side ); _ = orderSide;
  if IsTrue(OpEq2(orderSide , MkUndefined() )) {
    panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" createOrder does not support order side ") , OpAdd(side , MkString(", supported order sides are buy, sell") )))));
  }
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "account":parseInt(accountId ),
        "symbol":*(market ).At(MkString("id") ),
        "ordType":orderType ,
        "side":orderSide ,
        "orderQty":this.AmountToPrecision(symbol , amount ),
        "transactTime":OpMulti(this.Milliseconds(), MkInteger(1000000) ),
        }); _ = request;
  if IsTrue(OpOr(OpEq2(type_ , MkString("limit") ), OpEq2(type_ , MkString("stop-limit") ))) {
    if IsTrue(OpEq2(price , MkUndefined() )) {
      panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" createOrder() requires a price argument for order type ") , type_ ))));
    }
    *(request ).At(MkString("price") )= this.PriceToPrecision(symbol , price );
  }
  if IsTrue(OpOr(OpEq2(type_ , MkString("stop") ), OpEq2(type_ , MkString("stop-limit") ))) {
    stopPx := this.SafeNumber(params , MkString("stopPx") ); _ = stopPx;
    if IsTrue(OpEq2(stopPx , MkUndefined() )) {
      panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" createOrder() requires a stopPx param for order type ") , type_ ))));
    }
    *(request ).At(MkString("stopPx") )= this.PriceToPrecision(symbol , stopPx );
    params = this.Omit(params , MkString("stopPx") );
  }
  clientOrderId := this.SafeString2(params , MkString("clientOrderId") , MkString("clOrdId") , this.Uuid()); _ = clientOrderId;
  if IsTrue(OpNotEq2(clientOrderId , MkUndefined() )) {
    *(request ).At(MkString("clOrdId") )= OpCopy(clientOrderId );
    params = this.Omit(params , MkArray(&VarArray{
            MkString("clientOrderId") ,
            MkString("clOrdId") ,
            }));
  }
  response := this.Call(MkString("privatePostTradingOrderNew"), this.Extend(request , params )); _ = response;
  return this.ParseOrder(response , market );
}

func (this *Xena) EditOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 2, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 3, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 4, MkUndefined() ); _ = amount;
  price := GoGetArg(goArgs, 5, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 6, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" cancelOrder() requires a symbol argument") )));
  }
  this.LoadMarkets();
  this.LoadAccounts();
  accountId := this.GetAccountId(params ); _ = accountId;
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "account":parseInt(accountId ),
        "clOrdId":this.Uuid(),
        "symbol":*(market ).At(MkString("id") ),
        "transactTime":OpMulti(this.Milliseconds(), MkInteger(1000000) ),
        }); _ = request;
  clientOrderId := this.SafeString2(params , MkString("clientOrderId") , MkString("origClOrdId") ); _ = clientOrderId;
  if IsTrue(OpNotEq2(clientOrderId , MkUndefined() )) {
    *(request ).At(MkString("origClOrdId") )= OpCopy(clientOrderId );
    params = this.Omit(params , MkArray(&VarArray{
            MkString("clientOrderId") ,
            MkString("origClOrdId") ,
            }));
  } else {
    *(request ).At(MkString("orderId") )= OpCopy(id );
  }
  if IsTrue(OpNotEq2(amount , MkUndefined() )) {
    *(request ).At(MkString("orderQty") )= this.AmountToPrecision(symbol , amount );
  }
  if IsTrue(OpNotEq2(price , MkUndefined() )) {
    *(request ).At(MkString("price") )= this.PriceToPrecision(symbol , price );
  }
  stopPx := this.SafeNumber(params , MkString("stopPx") ); _ = stopPx;
  if IsTrue(OpNotEq2(stopPx , MkUndefined() )) {
    *(request ).At(MkString("stopPx") )= this.PriceToPrecision(symbol , stopPx );
    params = this.Omit(params , MkString("stopPx") );
  }
  capPrice := this.SafeNumber(params , MkString("capPrice") ); _ = capPrice;
  if IsTrue(OpNotEq2(capPrice , MkUndefined() )) {
    *(request ).At(MkString("capPrice") )= this.PriceToPrecision(symbol , capPrice );
    params = this.Omit(params , MkString("capPrice") );
  }
  response := this.Call(MkString("privatePostTradingOrderReplace"), this.Extend(request , params )); _ = response;
  return this.ParseOrder(response , market );
}

func (this *Xena) CancelOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" cancelOrder() requires a symbol argument") )));
  }
  this.LoadMarkets();
  this.LoadAccounts();
  accountId := this.GetAccountId(params ); _ = accountId;
  clientOrderId := this.SafeString2(params , MkString("clientOrderId") , MkString("origClOrdId") ); _ = clientOrderId;
  params = this.Omit(params , MkArray(&VarArray{
          MkString("clientOrderId") ,
          MkString("origClOrdId") ,
          }));
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "account":parseInt(accountId ),
        "symbol":*(market ).At(MkString("id") ),
        "clOrdId":this.Uuid(),
        "transactTime":OpMulti(this.Milliseconds(), MkInteger(1000000) ),
        }); _ = request;
  if IsTrue(OpNotEq2(clientOrderId , MkUndefined() )) {
    *(request ).At(MkString("origClOrdId") )= OpCopy(clientOrderId );
  } else {
    *(request ).At(MkString("orderId") )= OpCopy(id );
  }
  response := this.Call(MkString("privatePostTradingOrderCancel"), this.Extend(request , params )); _ = response;
  return this.ParseOrder(response , market );
}

func (this *Xena) CancelAllOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  this.LoadAccounts();
  accountId := this.GetAccountId(params ); _ = accountId;
  request := MkMap(&VarMap{
        "account":parseInt(accountId ),
        "clOrdId":this.Uuid(),
        }); _ = request;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market := this.Market(symbol ); _ = market;
    *(request ).At(MkString("symbol") )= *(market ).At(MkString("id") );
    *(request ).At(MkString("massCancelRequestType") )= MkString("1") ;
  } else {
    *(request ).At(MkString("massCancelRequestType") )= MkString("7") ;
  }
  response := this.Call(MkString("privatePostTradingOrderMassCancel"), this.Extend(request , params )); _ = response;
  return response ;
}

func (this *Xena) FetchOpenOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  this.LoadAccounts();
  accountId := this.GetAccountId(params ); _ = accountId;
  request := MkMap(&VarMap{"accountId":accountId }); _ = request;
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("symbol") )= *(market ).At(MkString("id") );
  }
  response := this.Call(MkString("privateGetTradingAccountsAccountIdActiveOrders"), this.Extend(request , params )); _ = response;
  return this.ParseOrders(response , market , since , limit );
}

func (this *Xena) FetchClosedOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  this.LoadAccounts();
  accountId := this.GetAccountId(params ); _ = accountId;
  request := MkMap(&VarMap{"accountId":accountId }); _ = request;
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("symbol") )= *(market ).At(MkString("id") );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("from") )= OpMulti(this.Iso8601(since ), MkInteger(1000000) );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("privateGetTradingAccountsAccountIdLastOrderStatuses"), this.Extend(request , params )); _ = response;
  return this.ParseOrders(response , market , since , limit );
}

func (this *Xena) CreateDepositAddress(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  this.LoadAccounts();
  accountId := this.GetAccountId(params ); _ = accountId;
  currency := this.Currency(code ); _ = currency;
  request := MkMap(&VarMap{
        "accountId":accountId ,
        "currency":*(currency ).At(MkString("id") ),
        }); _ = request;
  response := this.Call(MkString("privatePostTransfersAccountsAccountIdDepositAddressCurrency"), this.Extend(request , params )); _ = response;
  address := this.SafeValue(response , MkString("address") ); _ = address;
  tag := OpCopy(MkUndefined() ); _ = tag;
  this.CheckAddress(address );
  return MkMap(&VarMap{
        "currency":code ,
        "address":address ,
        "tag":tag ,
        "info":response ,
        });
}

func (this *Xena) FetchDepositAddress(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  this.LoadAccounts();
  accountId := this.GetAccountId(params ); _ = accountId;
  currency := this.Currency(code ); _ = currency;
  request := MkMap(&VarMap{
        "accountId":accountId ,
        "currency":*(currency ).At(MkString("id") ),
        }); _ = request;
  response := this.Call(MkString("privateGetTransfersAccountsAccountIdDepositAddressCurrency"), this.Extend(request , params )); _ = response;
  address := this.SafeValue(response , MkString("address") ); _ = address;
  tag := OpCopy(MkUndefined() ); _ = tag;
  this.CheckAddress(address );
  return MkMap(&VarMap{
        "currency":code ,
        "address":address ,
        "tag":tag ,
        "info":response ,
        });
}

func (this *Xena) FetchTransactionsByType(goArgs ...*Variant) *Variant{
  type_ := GoGetArg(goArgs, 0, MkUndefined()); _ = type_;
  code := GoGetArg(goArgs, 1, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(code , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchTransactions() requires a currency `code` argument") )));
  }
  this.LoadMarkets();
  this.LoadAccounts();
  accountId := this.GetAccountId(params ); _ = accountId;
  currency := this.Currency(code ); _ = currency;
  request := MkMap(&VarMap{
        "currency":*(currency ).At(MkString("id") ),
        "accountId":accountId ,
        }); _ = request;
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("since") )= parseInt(OpDiv(since , MkInteger(1000) ));
  }
  method := OpAdd(MkString("privateGetTransfersAccountsAccountId") , this.Capitalize(type_ )); _ = method;
  response := this.Call(method , this.Extend(request , params )); _ = response;
  transactions := this.SafeValue(response , type_ , MkArray(&VarArray{})); _ = transactions;
  return this.ParseTransactions(transactions , currency , since , limit );
}

func (this *Xena) FetchWithdrawals(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  return this.FetchTransactionsByType(MkString("withdrawals") , code , since , limit , params );
}

func (this *Xena) FetchDeposits(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  return this.FetchTransactionsByType(MkString("deposits") , code , since , limit , params );
}

func (this *Xena) ParseTransaction(goArgs ...*Variant) *Variant{
  transaction := GoGetArg(goArgs, 0, MkUndefined()); _ = transaction;
  currency := GoGetArg(goArgs, 1, MkUndefined() ); _ = currency;
  id := this.SafeString(transaction , MkString("withdrawalRequestId") ); _ = id;
  type_ := OpTrinary(OpEq2(id , MkUndefined() ), MkString("deposit") , MkString("withdrawal") ); _ = type_;
  updated := this.SafeInteger(transaction , MkString("lastUpdated") ); _ = updated;
  if IsTrue(OpNotEq2(updated , MkUndefined() )) {
    updated = parseInt(OpDiv(updated , MkInteger(1000000) ));
  }
  timestamp := OpCopy(MkUndefined() ); _ = timestamp;
  txid := this.SafeString(transaction , MkString("txId") ); _ = txid;
  currencyId := this.SafeString(transaction , MkString("currency") ); _ = currencyId;
  code := this.SafeCurrencyCode(currencyId , currency ); _ = code;
  address := this.SafeString(transaction , MkString("address") ); _ = address;
  addressFrom := OpCopy(MkUndefined() ); _ = addressFrom;
  addressTo := OpCopy(address ); _ = addressTo;
  amount := this.SafeNumber(transaction , MkString("amount") ); _ = amount;
  status := this.ParseTransactionStatus(this.SafeString(transaction , MkString("status") )); _ = status;
  fee := OpCopy(MkUndefined() ); _ = fee;
  return MkMap(&VarMap{
        "info":transaction ,
        "id":id ,
        "txid":txid ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "addressFrom":addressFrom ,
        "addressTo":addressTo ,
        "address":address ,
        "tagFrom":MkUndefined() ,
        "tagTo":MkUndefined() ,
        "tag":MkUndefined() ,
        "type":type_ ,
        "amount":amount ,
        "currency":code ,
        "status":status ,
        "updated":updated ,
        "fee":fee ,
        });
}

func (this *Xena) ParseTransactionStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "1":MkString("pending") ,
        "2":MkString("ok") ,
        "3":MkString("failed") ,
        "4":MkString("failed") ,
        "5":MkString("pending") ,
        "100":MkString("pending") ,
        "101":MkString("pending") ,
        "102":MkString("pending") ,
        "103":MkString("pending") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Xena) Withdraw(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  amount := GoGetArg(goArgs, 1, MkUndefined()); _ = amount;
  address := GoGetArg(goArgs, 2, MkUndefined()); _ = address;
  tag := GoGetArg(goArgs, 3, MkUndefined() ); _ = tag;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.CheckAddress(address );
  this.LoadMarkets();
  this.LoadAccounts();
  accountId := this.GetAccountId(params ); _ = accountId;
  currency := this.Currency(code ); _ = currency;
  uuid := this.Uuid(); _ = uuid;
  uuid = uuid.Split(MkString("-") );
  uuid = uuid.Join(MkString("") );
  request := MkMap(&VarMap{
        "currency":*(currency ).At(MkString("id") ),
        "accountId":accountId ,
        "amount":this.CurrencyToPrecision(code , amount ),
        "address":address ,
        "id":uuid ,
        }); _ = request;
  response := this.Call(MkString("privatePostTransfersAccountsAccountIdWithdrawals"), this.Extend(request , params )); _ = response;
  return this.ParseTransaction(response , currency );
}

func (this *Xena) ParseLedgerEntryType(goArgs ...*Variant) *Variant{
  type_ := GoGetArg(goArgs, 0, MkUndefined()); _ = type_;
  types := MkMap(&VarMap{
        "deposit":MkString("transaction") ,
        "withdrawal":MkString("transaction") ,
        "internal deposit":MkString("transfer") ,
        "internal withdrawal":MkString("transfer") ,
        "rebate":MkString("rebate") ,
        "reward":MkString("reward") ,
        }); _ = types;
  return this.SafeString(types , type_ , type_ );
}

func (this *Xena) ParseLedgerEntry(goArgs ...*Variant) *Variant{
  item := GoGetArg(goArgs, 0, MkUndefined()); _ = item;
  currency := GoGetArg(goArgs, 1, MkUndefined() ); _ = currency;
  id := this.SafeString(item , MkString("id") ); _ = id;
  direction := OpCopy(MkUndefined() ); _ = direction;
  account := this.SafeString(item , MkString("accountId") ); _ = account;
  referenceId := OpCopy(MkUndefined() ); _ = referenceId;
  referenceAccount := OpCopy(MkUndefined() ); _ = referenceAccount;
  type_ := this.ParseLedgerEntryType(this.SafeString(item , MkString("kind") )); _ = type_;
  code := this.SafeCurrencyCode(this.SafeString(item , MkString("currency") ), currency ); _ = code;
  amount := this.SafeNumber(item , MkString("amount") ); _ = amount;
  if IsTrue(OpLw(amount , MkInteger(0) )) {
    direction = MkString("out") ;
    amount = Math.Abs(amount );
  } else {
    direction = MkString("in") ;
  }
  timestamp := this.SafeInteger(item , MkString("ts") ); _ = timestamp;
  if IsTrue(OpNotEq2(timestamp , MkUndefined() )) {
    timestamp = parseInt(OpDiv(timestamp , MkInteger(1000000) ));
  }
  fee := MkMap(&VarMap{
        "cost":this.SafeNumber(item , MkString("commission") ),
        "currency":code ,
        }); _ = fee;
  before := OpCopy(MkUndefined() ); _ = before;
  after := this.SafeNumber(item , MkString("balance") ); _ = after;
  status := MkString("ok") ; _ = status;
  return MkMap(&VarMap{
        "info":item ,
        "id":id ,
        "direction":direction ,
        "account":account ,
        "referenceId":referenceId ,
        "referenceAccount":referenceAccount ,
        "type":type_ ,
        "currency":code ,
        "amount":amount ,
        "before":before ,
        "after":after ,
        "status":status ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "fee":fee ,
        });
}

func (this *Xena) FetchLedger(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  this.LoadAccounts();
  accountId := this.GetAccountId(params ); _ = accountId;
  request := MkMap(&VarMap{"accountId":accountId }); _ = request;
  currency := OpCopy(MkUndefined() ); _ = currency;
  if IsTrue(OpNotEq2(code , MkUndefined() )) {
    currency = this.Currency(code );
    *(request ).At(MkString("symbol") )= *(currency ).At(MkString("id") );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("from") )= OpMulti(since , MkInteger(1000000) );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("privateGetTransfersAccountsAccountIdBalanceHistory"), this.Extend(request , params )); _ = response;
  return this.ParseLedger(response , currency , since , limit );
}

func (this *Xena) Nonce(goArgs ...*Variant) *Variant{
  return this.Milliseconds();
}

func (this *Xena) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("public") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  url := OpAdd(*(*(*this.At(MkString("urls")) ).At(MkString("api") )).At(api ), OpAdd(MkString("/") , this.ImplodeParams(path , params ))); _ = url;
  query := this.Omit(params , this.ExtractParams(path )); _ = query;
  if IsTrue(OpEq2(api , MkString("public") )) {
    if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
      OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(query )));
    }
  } else {
    if IsTrue(OpEq2(api , MkString("private") )) {
      this.CheckRequiredCredentials();
      nonce := this.Nonce(); _ = nonce;
      nonce = nonce.ToString();
      nonce = OpAdd(nonce , MkString("000000") );
      payload := OpAdd(MkString("AUTH") , nonce ); _ = payload;
      secret := (*((this).At(MkString("secret")))).Call(MkString("slice"), MkInteger(14) , MkInteger(78) ); _ = secret;
      ecdsa := this.Ecdsa(payload , secret , MkString("p256") , MkString("sha256") ); _ = ecdsa;
      signature := OpAdd(*(ecdsa ).At(MkString("r") ), *(ecdsa ).At(MkString("s") )); _ = signature;
      headers = MkMap(&VarMap{
          "X-AUTH-API-KEY":*this.At(MkString("apiKey")) ,
          "X-AUTH-API-PAYLOAD":payload ,
          "X-AUTH-API-SIGNATURE":signature ,
          "X-AUTH-API-NONCE":nonce ,
          });
      if IsTrue(OpEq2(method , MkString("GET") )) {
        if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
          OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(query )));
        }
      } else {
        if IsTrue(OpEq2(method , MkString("POST") )) {
          body = this.Json(query );
          *(headers ).At(MkString("Content-Type") )= MkString("application/json") ;
        }
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

func (this *Xena) HandleErrors(goArgs ...*Variant) *Variant{
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
  if IsTrue(OpGtEq(code , MkInteger(400) )) {
    feedback := OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , this.Json(response ))); _ = feedback;
    message := this.SafeString(response , MkString("error") ); _ = message;
    exact := *(*this.At(MkString("exceptions")) ).At(MkString("exact") ); _ = exact;
    if IsTrue(OpHasMember(message , exact )) {
      panic(exact.Call(message , feedback ));
    }
    broad := *(*this.At(MkString("exceptions")) ).At(MkString("broad") ); _ = broad;
    broadKey := this.FindBroadlyMatchedKey(broad , body ); _ = broadKey;
    if IsTrue(OpNotEq2(broadKey , MkUndefined() )) {
      panic(broad.Call(broadKey , feedback ));
    }
    panic(NewExchangeError(feedback ));
  }
  return MkUndefined()
}

