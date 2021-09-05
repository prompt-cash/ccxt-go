package ccxt

import (
)

type Hitbtc struct {
	*ExchangeBase
}

var _ Exchange = (*Hitbtc)(nil)

func init() {
	exchange := &Hitbtc{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Hitbtc) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("hitbtc") ,
            "name":MkString("HitBTC") ,
            "countries":MkArray(&VarArray{
                MkString("HK") ,
                }),
            "rateLimit":MkInteger(1500) ,
            "version":MkString("2") ,
            "pro":MkBool(true) ,
            "has":MkMap(&VarMap{
                "cancelOrder":MkBool(true) ,
                "CORS":MkBool(false) ,
                "createDepositAddress":MkBool(true) ,
                "createOrder":MkBool(true) ,
                "editOrder":MkBool(true) ,
                "fetchBalance":MkBool(true) ,
                "fetchClosedOrders":MkBool(true) ,
                "fetchCurrencies":MkBool(true) ,
                "fetchDepositAddress":MkBool(true) ,
                "fetchDeposits":MkBool(false) ,
                "fetchMarkets":MkBool(true) ,
                "fetchMyTrades":MkBool(true) ,
                "fetchOHLCV":MkBool(true) ,
                "fetchOpenOrder":MkBool(true) ,
                "fetchOpenOrders":MkBool(true) ,
                "fetchOrder":MkBool(true) ,
                "fetchOrderBook":MkBool(true) ,
                "fetchOrders":MkBool(false) ,
                "fetchOrderTrades":MkBool(true) ,
                "fetchTicker":MkBool(true) ,
                "fetchTickers":MkBool(true) ,
                "fetchTrades":MkBool(true) ,
                "fetchTradingFee":MkBool(true) ,
                "fetchTransactions":MkBool(true) ,
                "fetchWithdrawals":MkBool(false) ,
                "withdraw":MkBool(true) ,
                "transfer":MkBool(true) ,
                }),
            "timeframes":MkMap(&VarMap{
                "1m":MkString("M1") ,
                "3m":MkString("M3") ,
                "5m":MkString("M5") ,
                "15m":MkString("M15") ,
                "30m":MkString("M30") ,
                "1h":MkString("H1") ,
                "4h":MkString("H4") ,
                "1d":MkString("D1") ,
                "1w":MkString("D7") ,
                "1M":MkString("1M") ,
                }),
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/1294454/27766555-8eaec20e-5edc-11e7-9c5b-6dc69fc42f5e.jpg") ,
                "test":MkMap(&VarMap{
                    "public":MkString("https://api.demo.hitbtc.com") ,
                    "private":MkString("https://api.demo.hitbtc.com") ,
                    }),
                "api":MkMap(&VarMap{
                    "public":MkString("https://api.hitbtc.com") ,
                    "private":MkString("https://api.hitbtc.com") ,
                    }),
                "www":MkString("https://hitbtc.com") ,
                "referral":MkString("https://hitbtc.com/?ref_id=5a5d39a65d466") ,
                "doc":MkArray(&VarArray{
                    MkString("https://api.hitbtc.com") ,
                    MkString("https://github.com/hitbtc-com/hitbtc-api/blob/master/APIv2.md") ,
                    }),
                "fees":MkArray(&VarArray{
                    MkString("https://hitbtc.com/fees-and-limits") ,
                    MkString("https://support.hitbtc.com/hc/en-us/articles/115005148605-Fees-and-limits") ,
                    }),
                }),
            "api":MkMap(&VarMap{
                "public":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("currency") ,
                        MkString("currency/{currency}") ,
                        MkString("symbol") ,
                        MkString("symbol/{symbol}") ,
                        MkString("ticker") ,
                        MkString("ticker/{symbol}") ,
                        MkString("trades") ,
                        MkString("trades/{symbol}") ,
                        MkString("orderbook") ,
                        MkString("orderbook/{symbol}") ,
                        MkString("candles") ,
                        MkString("candles/{symbol}") ,
                        })}),
                "private":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("trading/balance") ,
                        MkString("order") ,
                        MkString("order/{clientOrderId}") ,
                        MkString("trading/fee/all") ,
                        MkString("trading/fee/{symbol}") ,
                        MkString("margin/account") ,
                        MkString("margin/account/{symbol}") ,
                        MkString("margin/position") ,
                        MkString("margin/position/{symbol}") ,
                        MkString("margin/order") ,
                        MkString("margin/order/{clientOrderId}") ,
                        MkString("history/order") ,
                        MkString("history/trades") ,
                        MkString("history/order/{orderId}/trades") ,
                        MkString("account/balance") ,
                        MkString("account/crypto/address/{currency}") ,
                        MkString("account/crypto/addresses/{currency}") ,
                        MkString("account/crypto/used-addresses/{currency}") ,
                        MkString("account/crypto/estimate-withdraw") ,
                        MkString("account/crypto/is-mine/{address}") ,
                        MkString("account/transactions") ,
                        MkString("account/transactions/{id}") ,
                        MkString("sub-acc") ,
                        MkString("sub-acc/acl") ,
                        MkString("sub-acc/balance/{subAccountUserID}") ,
                        MkString("sub-acc/deposit-address/{subAccountUserId}/{currency}") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("order") ,
                        MkString("margin/order") ,
                        MkString("account/crypto/address/{currency}") ,
                        MkString("account/crypto/withdraw") ,
                        MkString("account/crypto/transfer-convert") ,
                        MkString("account/transfer") ,
                        MkString("account/transfer/internal") ,
                        MkString("sub-acc/freeze") ,
                        MkString("sub-acc/activate") ,
                        MkString("sub-acc/transfer") ,
                        }),
                    "put":MkArray(&VarArray{
                        MkString("order/{clientOrderId}") ,
                        MkString("margin/account/{symbol}") ,
                        MkString("margin/order/{clientOrderId}") ,
                        MkString("account/crypto/withdraw/{id}") ,
                        MkString("sub-acc/acl/{subAccountUserId}") ,
                        }),
                    "delete":MkArray(&VarArray{
                        MkString("order") ,
                        MkString("order/{clientOrderId}") ,
                        MkString("margin/account") ,
                        MkString("margin/account/{symbol}") ,
                        MkString("margin/position") ,
                        MkString("margin/position/{symbol}") ,
                        MkString("margin/order") ,
                        MkString("margin/order/{clientOrderId}") ,
                        MkString("account/crypto/withdraw/{id}") ,
                        }),
                    "patch":MkArray(&VarArray{
                        MkString("order/{clientOrderId}") ,
                        }),
                    }),
                }),
            "precisionMode":TICK_SIZE ,
            "fees":MkMap(&VarMap{"trading":MkMap(&VarMap{
                    "tierBased":MkBool(false) ,
                    "percentage":MkBool(true) ,
                    "maker":OpDiv(MkNumber(0.1) , MkInteger(100) ),
                    "taker":OpDiv(MkNumber(0.2) , MkInteger(100) ),
                    })}),
            "options":MkMap(&VarMap{
                "defaultTimeInForce":MkString("FOK") ,
                "accountsByType":MkMap(&VarMap{
                    "bank":MkString("bank") ,
                    "exchange":MkString("exchange") ,
                    "main":MkString("bank") ,
                    "funding":MkString("bank") ,
                    "spot":MkString("exchange") ,
                    "trade":MkString("exchange") ,
                    "trading":MkString("exchange") ,
                    }),
                "fetchBalanceMethod":MkMap(&VarMap{
                    "account":MkString("account") ,
                    "bank":MkString("account") ,
                    "main":MkString("account") ,
                    "funding":MkString("account") ,
                    "exchange":MkString("trading") ,
                    "spot":MkString("trading") ,
                    "trade":MkString("trading") ,
                    "trading":MkString("trading") ,
                    }),
                }),
            "commonCurrencies":MkMap(&VarMap{
                "AUTO":MkString("Cube") ,
                "BCC":MkString("BCC") ,
                "BDP":MkString("BidiPass") ,
                "BET":MkString("DAO.Casino") ,
                "BIT":MkString("BitRewards") ,
                "BOX":MkString("BOX Token") ,
                "CPT":MkString("Cryptaur") ,
                "GET":MkString("Themis") ,
                "HSR":MkString("HC") ,
                "IQ":MkString("IQ.Cash") ,
                "LNC":MkString("LinkerCoin") ,
                "PLA":MkString("PlayChip") ,
                "PNT":MkString("Penta") ,
                "SBTC":MkString("Super Bitcoin") ,
                "STX":MkString("Stox") ,
                "TV":MkString("Tokenville") ,
                "USD":MkString("USDT") ,
                "XPNT":MkString("PNT") ,
                }),
            "exceptions":MkMap(&VarMap{
                "504":RequestTimeout ,
                "1002":AuthenticationError ,
                "1003":PermissionDenied ,
                "2010":InvalidOrder ,
                "2001":BadSymbol ,
                "2011":InvalidOrder ,
                "2020":InvalidOrder ,
                "20002":OrderNotFound ,
                "20001":InsufficientFunds ,
                }),
            }));
}

func (this *Hitbtc) FeeToPrecision(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  fee := GoGetArg(goArgs, 1, MkUndefined()); _ = fee;
  return this.DecimalToPrecision(fee , TRUNCATE , MkInteger(8) , DECIMAL_PLACES );
}

func (this *Hitbtc) FetchMarkets(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetSymbol"), params ); _ = response;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , response.Length )); OpIncr(& i ){
    market := *(response ).At(i ); _ = market;
    id := this.SafeString(market , MkString("id") ); _ = id;
    baseId := this.SafeString(market , MkString("baseCurrency") ); _ = baseId;
    quoteId := this.SafeString(market , MkString("quoteCurrency") ); _ = quoteId;
    base := this.SafeCurrencyCode(baseId ); _ = base;
    quote := this.SafeCurrencyCode(quoteId ); _ = quote;
    symbol := OpAdd(base , OpAdd(MkString("/") , quote )); _ = symbol;
    if IsTrue(OpGtEq(id.IndexOf(MkString("_") ), MkInteger(0) )) {
      symbol = OpCopy(id );
    }
    lotString := this.SafeString(market , MkString("quantityIncrement") ); _ = lotString;
    stepString := this.SafeString(market , MkString("tickSize") ); _ = stepString;
    lot := this.ParseNumber(lotString ); _ = lot;
    step := this.ParseNumber(stepString ); _ = step;
    precision := MkMap(&VarMap{
          "price":step ,
          "amount":lot ,
          }); _ = precision;
    taker := this.SafeNumber(market , MkString("takeLiquidityRate") ); _ = taker;
    maker := this.SafeNumber(market , MkString("provideLiquidityRate") ); _ = maker;
    feeCurrencyId := this.SafeString(market , MkString("feeCurrency") ); _ = feeCurrencyId;
    feeCurrencyCode := this.SafeCurrencyCode(feeCurrencyId ); _ = feeCurrencyCode;
    result.Push(this.Extend(*(*this.At(MkString("fees")) ).At(MkString("trading") ), MkMap(&VarMap{
                "info":market ,
                "id":id ,
                "symbol":symbol ,
                "base":base ,
                "quote":quote ,
                "baseId":baseId ,
                "quoteId":quoteId ,
                "active":MkBool(true) ,
                "taker":taker ,
                "maker":maker ,
                "precision":precision ,
                "feeCurrency":feeCurrencyCode ,
                "limits":MkMap(&VarMap{
                    "amount":MkMap(&VarMap{
                        "min":lot ,
                        "max":MkUndefined() ,
                        }),
                    "price":MkMap(&VarMap{
                        "min":step ,
                        "max":MkUndefined() ,
                        }),
                    "cost":MkMap(&VarMap{
                        "min":this.ParseNumber(Precise.StringMul(lotString , stepString )),
                        "max":MkUndefined() ,
                        }),
                    }),
                })));
  }
  return result ;
}

func (this *Hitbtc) Transfer(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  amount := GoGetArg(goArgs, 1, MkUndefined()); _ = amount;
  fromAccount := GoGetArg(goArgs, 2, MkUndefined()); _ = fromAccount;
  toAccount := GoGetArg(goArgs, 3, MkUndefined()); _ = toAccount;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  currency := this.Currency(code ); _ = currency;
  requestAmount := this.CurrencyToPrecision(code , amount ); _ = requestAmount;
  request := MkMap(&VarMap{
        "currency":*(currency ).At(MkString("id") ),
        "amount":requestAmount ,
        }); _ = request;
  type_ := this.SafeString(params , MkString("type") ); _ = type_;
  if IsTrue(OpEq2(type_ , MkUndefined() )) {
    accountsByType := this.SafeValue(*this.At(MkString("options")) , MkString("accountsByType") , MkMap(&VarMap{})); _ = accountsByType;
    fromId := this.SafeString(accountsByType , fromAccount ); _ = fromId;
    toId := this.SafeString(accountsByType , toAccount ); _ = toId;
    keys := GoGetKeys(accountsByType ); _ = keys;
    if IsTrue(OpEq2(fromId , MkUndefined() )) {
      panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" fromAccount must be one of ") , OpAdd(keys.Join(MkString(", ") ), OpAdd(MkString(" instead of ") , fromId ))))));
    }
    if IsTrue(OpEq2(toId , MkUndefined() )) {
      panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" toAccount must be one of ") , OpAdd(keys.Join(MkString(", ") ), OpAdd(MkString(" instead of ") , toId ))))));
    }
    if IsTrue(OpEq2(fromId , toId )) {
      panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , MkString(" from and to cannot be the same account") )));
    }
    type_ = OpAdd(fromId , OpAdd(MkString("To") , this.Capitalize(toId )));
  }
  *(request ).At(MkString("type") )= OpCopy(type_ );
  response := this.Call(MkString("privatePostAccountTransfer"), this.Extend(request , params )); _ = response;
  id := this.SafeString(response , MkString("id") ); _ = id;
  return MkMap(&VarMap{
        "info":response ,
        "id":id ,
        "timestamp":MkUndefined() ,
        "datetime":MkUndefined() ,
        "amount":requestAmount ,
        "currency":code ,
        "fromAccount":fromAccount ,
        "toAccount":toAccount ,
        "status":MkUndefined() ,
        });
}

func (this *Hitbtc) FetchCurrencies(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetCurrency"), params ); _ = response;
  result := MkMap(&VarMap{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , response.Length )); OpIncr(& i ){
    currency := *(response ).At(i ); _ = currency;
    id := this.SafeString(currency , MkString("id") ); _ = id;
    decimals := this.SafeInteger(currency , MkString("precisionTransfer") , MkInteger(8) ); _ = decimals;
    precision := OpDiv(MkInteger(1) , Math.Pow(MkInteger(10) , decimals )); _ = precision;
    code := this.SafeCurrencyCode(id ); _ = code;
    payin := this.SafeValue(currency , MkString("payinEnabled") ); _ = payin;
    payout := this.SafeValue(currency , MkString("payoutEnabled") ); _ = payout;
    transfer := this.SafeValue(currency , MkString("transferEnabled") ); _ = transfer;
    active := OpAnd(payin , OpAnd(payout , transfer )); _ = active;
    if IsTrue(OpHasMember(MkString("disabled") , currency )) {
      if IsTrue(*(currency ).At(MkString("disabled") )) {
        active = OpCopy(MkBool(false) );
      }
    }
    type_ := MkString("fiat") ; _ = type_;
    if IsTrue(OpAnd(OpHasMember(MkString("crypto") , currency ), *(currency ).At(MkString("crypto") ))) {
      type_ = MkString("crypto") ;
    }
    name := this.SafeString(currency , MkString("fullName") ); _ = name;
    *(result ).At(code )= MkMap(&VarMap{
        "id":id ,
        "code":code ,
        "type":type_ ,
        "payin":payin ,
        "payout":payout ,
        "transfer":transfer ,
        "info":currency ,
        "name":name ,
        "active":active ,
        "fee":this.SafeNumber(currency , MkString("payoutFee") ),
        "precision":precision ,
        "limits":MkMap(&VarMap{
            "amount":MkMap(&VarMap{
                "min":OpDiv(MkInteger(1) , Math.Pow(MkInteger(10) , decimals )),
                "max":Math.Pow(MkInteger(10) , decimals ),
                }),
            "withdraw":MkMap(&VarMap{
                "min":MkUndefined() ,
                "max":Math.Pow(MkInteger(10) , precision ),
                }),
            }),
        });
  }
  return result ;
}

func (this *Hitbtc) FetchTradingFee(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := this.Extend(MkMap(&VarMap{"symbol":*(market ).At(MkString("id") )}), this.Omit(params , MkString("symbol") )); _ = request;
  response := this.Call(MkString("privateGetTradingFeeSymbol"), request ); _ = response;
  return MkMap(&VarMap{
        "info":response ,
        "maker":this.SafeNumber(response , MkString("provideLiquidityRate") ),
        "taker":this.SafeNumber(response , MkString("takeLiquidityRate") ),
        });
}

func (this *Hitbtc) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  type_ := this.SafeString(params , MkString("type") , MkString("trading") ); _ = type_;
  fetchBalanceAccounts := this.SafeValue(*this.At(MkString("options")) , MkString("fetchBalanceMethod") , MkMap(&VarMap{})); _ = fetchBalanceAccounts;
  typeId := this.SafeString(fetchBalanceAccounts , type_ ); _ = typeId;
  if IsTrue(OpEq2(typeId , MkUndefined() )) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , MkString(" fetchBalance account type must be either main or trading") )));
  }
  method := OpAdd(MkString("privateGet") , OpAdd(this.Capitalize(typeId ), MkString("Balance") )); _ = method;
  query := this.Omit(params , MkString("type") ); _ = query;
  response := this.Call(method , query ); _ = response;
  result := MkMap(&VarMap{
        "info":response ,
        "timestamp":MkUndefined() ,
        "datetime":MkUndefined() ,
        }); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , response.Length )); OpIncr(& i ){
    balance := *(response ).At(i ); _ = balance;
    currencyId := this.SafeString(balance , MkString("currency") ); _ = currencyId;
    code := this.SafeCurrencyCode(currencyId ); _ = code;
    account := this.Account(); _ = account;
    *(account ).At(MkString("free") )= this.SafeString(balance , MkString("available") );
    *(account ).At(MkString("used") )= this.SafeString(balance , MkString("reserved") );
    *(result ).At(code )= OpCopy(account );
  }
  return this.ParseBalance(result );
}

func (this *Hitbtc) ParseOHLCV(goArgs ...*Variant) *Variant{
  ohlcv := GoGetArg(goArgs, 0, MkUndefined()); _ = ohlcv;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  return MkArray(&VarArray{
        this.Parse8601(this.SafeString(ohlcv , MkString("timestamp") )),
        this.SafeNumber(ohlcv , MkString("open") ),
        this.SafeNumber(ohlcv , MkString("max") ),
        this.SafeNumber(ohlcv , MkString("min") ),
        this.SafeNumber(ohlcv , MkString("close") ),
        this.SafeNumber(ohlcv , MkString("volume") ),
        });
}

func (this *Hitbtc) FetchOHLCV(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  timeframe := GoGetArg(goArgs, 1, MkString("1m") ); _ = timeframe;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "symbol":*(market ).At(MkString("id") ),
        "period":*(*this.At(MkString("timeframes")) ).At(timeframe ),
        }); _ = request;
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("from") )= this.Iso8601(since );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("publicGetCandlesSymbol"), this.Extend(request , params )); _ = response;
  return this.ParseOHLCVs(response , market , timeframe , since , limit );
}

func (this *Hitbtc) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"symbol":this.MarketId(symbol )}); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("publicGetOrderbookSymbol"), this.Extend(request , params )); _ = response;
  return this.ParseOrderBook(response , symbol , MkUndefined() , MkString("bid") , MkString("ask") , MkString("price") , MkString("size") );
}

func (this *Hitbtc) ParseTicker(goArgs ...*Variant) *Variant{
  ticker := GoGetArg(goArgs, 0, MkUndefined()); _ = ticker;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.Parse8601(*(ticker ).At(MkString("timestamp") )); _ = timestamp;
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  if IsTrue(OpNotEq2(market , MkUndefined() )) {
    symbol = *(market ).At(MkString("symbol") );
  }
  baseVolume := this.SafeNumber(ticker , MkString("volume") ); _ = baseVolume;
  quoteVolume := this.SafeNumber(ticker , MkString("volumeQuote") ); _ = quoteVolume;
  open := this.SafeNumber(ticker , MkString("open") ); _ = open;
  last := this.SafeNumber(ticker , MkString("last") ); _ = last;
  change := OpCopy(MkUndefined() ); _ = change;
  percentage := OpCopy(MkUndefined() ); _ = percentage;
  average := OpCopy(MkUndefined() ); _ = average;
  if IsTrue(OpAnd(OpNotEq2(last , MkUndefined() ), OpNotEq2(open , MkUndefined() ))) {
    change = OpSub(last , open );
    average = OpDiv(this.Sum(last , open ), MkInteger(2) );
    if IsTrue(OpGt(open , MkInteger(0) )) {
      percentage = OpDiv(change , OpMulti(open , MkInteger(100) ));
    }
  }
  vwap := this.Vwap(baseVolume , quoteVolume ); _ = vwap;
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

func (this *Hitbtc) FetchTickers(goArgs ...*Variant) *Variant{
  symbols := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbols;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("publicGetTicker"), params ); _ = response;
  result := MkMap(&VarMap{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , response.Length )); OpIncr(& i ){
    ticker := *(response ).At(i ); _ = ticker;
    marketId := this.SafeString(ticker , MkString("symbol") ); _ = marketId;
    market := this.SafeMarket(marketId ); _ = market;
    symbol := *(market ).At(MkString("symbol") ); _ = symbol;
    *(result ).At(symbol )= this.ParseTicker(ticker , market );
  }
  return this.FilterByArray(result , MkString("symbol") , symbols );
}

func (this *Hitbtc) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"symbol":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("publicGetTickerSymbol"), this.Extend(request , params )); _ = response;
  if IsTrue(OpHasMember(MkString("message") , response )) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , *(response ).At(MkString("message") )))));
  }
  return this.ParseTicker(response , market );
}

func (this *Hitbtc) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.Parse8601(*(trade ).At(MkString("timestamp") )); _ = timestamp;
  marketId := this.SafeString(trade , MkString("symbol") ); _ = marketId;
  market = this.SafeMarket(marketId , market );
  symbol := *(market ).At(MkString("symbol") ); _ = symbol;
  fee := OpCopy(MkUndefined() ); _ = fee;
  feeCost := this.SafeNumber(trade , MkString("fee") ); _ = feeCost;
  if IsTrue(OpNotEq2(feeCost , MkUndefined() )) {
    feeCurrencyCode := OpTrinary(market , *(market ).At(MkString("feeCurrency") ), MkUndefined() ); _ = feeCurrencyCode;
    fee = MkMap(&VarMap{
        "cost":feeCost ,
        "currency":feeCurrencyCode ,
        });
  }
  orderId := this.SafeString(trade , MkString("clientOrderId") ); _ = orderId;
  priceString := this.SafeString(trade , MkString("price") ); _ = priceString;
  amountString := this.SafeString(trade , MkString("quantity") ); _ = amountString;
  price := this.ParseNumber(priceString ); _ = price;
  amount := this.ParseNumber(amountString ); _ = amount;
  cost := this.ParseNumber(Precise.StringMul(priceString , amountString )); _ = cost;
  side := this.SafeString(trade , MkString("side") ); _ = side;
  id := this.SafeString(trade , MkString("id") ); _ = id;
  return MkMap(&VarMap{
        "info":trade ,
        "id":id ,
        "order":orderId ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "symbol":symbol ,
        "type":MkUndefined() ,
        "side":side ,
        "takerOrMaker":MkUndefined() ,
        "price":price ,
        "amount":amount ,
        "cost":cost ,
        "fee":fee ,
        });
}

func (this *Hitbtc) FetchTransactions(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  currency := OpCopy(MkUndefined() ); _ = currency;
  request := MkMap(&VarMap{}); _ = request;
  if IsTrue(OpNotEq2(code , MkUndefined() )) {
    currency = this.Currency(code );
    *(request ).At(MkString("asset") )= *(currency ).At(MkString("id") );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("startTime") )= OpCopy(since );
  }
  response := this.Call(MkString("privateGetAccountTransactions"), this.Extend(request , params )); _ = response;
  return this.ParseTransactions(response , currency , since , limit );
}

func (this *Hitbtc) ParseTransaction(goArgs ...*Variant) *Variant{
  transaction := GoGetArg(goArgs, 0, MkUndefined()); _ = transaction;
  currency := GoGetArg(goArgs, 1, MkUndefined() ); _ = currency;
  id := this.SafeString(transaction , MkString("id") ); _ = id;
  timestamp := this.Parse8601(this.SafeString(transaction , MkString("createdAt") )); _ = timestamp;
  updated := this.Parse8601(this.SafeString(transaction , MkString("updatedAt") )); _ = updated;
  currencyId := this.SafeString(transaction , MkString("currency") ); _ = currencyId;
  code := this.SafeCurrencyCode(currencyId , currency ); _ = code;
  status := this.ParseTransactionStatus(this.SafeString(transaction , MkString("status") )); _ = status;
  amount := this.SafeNumber(transaction , MkString("amount") ); _ = amount;
  address := this.SafeString(transaction , MkString("address") ); _ = address;
  txid := this.SafeString(transaction , MkString("hash") ); _ = txid;
  fee := OpCopy(MkUndefined() ); _ = fee;
  feeCost := this.SafeNumber(transaction , MkString("fee") ); _ = feeCost;
  if IsTrue(OpNotEq2(feeCost , MkUndefined() )) {
    fee = MkMap(&VarMap{
        "cost":feeCost ,
        "currency":code ,
        });
  }
  type_ := this.ParseTransactionType(this.SafeString(transaction , MkString("type") )); _ = type_;
  return MkMap(&VarMap{
        "info":transaction ,
        "id":id ,
        "txid":txid ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "address":address ,
        "tag":MkUndefined() ,
        "type":type_ ,
        "amount":amount ,
        "currency":code ,
        "status":status ,
        "updated":updated ,
        "fee":fee ,
        });
}

func (this *Hitbtc) ParseTransactionStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "pending":MkString("pending") ,
        "failed":MkString("failed") ,
        "success":MkString("ok") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Hitbtc) ParseTransactionType(goArgs ...*Variant) *Variant{
  type_ := GoGetArg(goArgs, 0, MkUndefined()); _ = type_;
  types := MkMap(&VarMap{
        "payin":MkString("deposit") ,
        "payout":MkString("withdrawal") ,
        "withdraw":MkString("withdrawal") ,
        }); _ = types;
  return this.SafeString(types , type_ , type_ );
}

func (this *Hitbtc) FetchTrades(goArgs ...*Variant) *Variant{
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
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("sort") )= MkString("ASC") ;
    *(request ).At(MkString("from") )= this.Iso8601(since );
  }
  response := this.Call(MkString("publicGetTradesSymbol"), this.Extend(request , params )); _ = response;
  return this.ParseTrades(response , market , since , limit );
}

func (this *Hitbtc) CreateOrder(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 2, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 3, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 4, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 5, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  uuid := this.Uuid(); _ = uuid;
  parts := uuid.Split(MkString("-") ); _ = parts;
  clientOrderId := parts.Join(MkString("") ); _ = clientOrderId;
  clientOrderId = clientOrderId.Slice(MkInteger(0) , MkInteger(32) );
  amount = parseFloat(amount );
  request := MkMap(&VarMap{
        "clientOrderId":clientOrderId ,
        "symbol":*(market ).At(MkString("id") ),
        "side":side ,
        "quantity":this.AmountToPrecision(symbol , amount ),
        "type":type_ ,
        }); _ = request;
  if IsTrue(OpEq2(type_ , MkString("limit") )) {
    *(request ).At(MkString("price") )= this.PriceToPrecision(symbol , price );
  } else {
    *(request ).At(MkString("timeInForce") )= *(*this.At(MkString("options")) ).At(MkString("defaultTimeInForce") );
  }
  response := this.Call(MkString("privatePostOrder"), this.Extend(request , params )); _ = response;
  order := this.ParseOrder(response ); _ = order;
  if IsTrue(OpEq2(*(order ).At(MkString("status") ), MkString("rejected") )) {
    panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" order was rejected by the exchange ") , this.Json(order )))));
  }
  return order ;
}

func (this *Hitbtc) EditOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 2, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 3, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 4, MkUndefined() ); _ = amount;
  price := GoGetArg(goArgs, 5, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 6, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  uuid := this.Uuid(); _ = uuid;
  parts := uuid.Split(MkString("-") ); _ = parts;
  requestClientId := parts.Join(MkString("") ); _ = requestClientId;
  requestClientId = requestClientId.Slice(MkInteger(0) , MkInteger(32) );
  request := MkMap(&VarMap{
        "clientOrderId":id ,
        "requestClientId":requestClientId ,
        }); _ = request;
  if IsTrue(OpNotEq2(amount , MkUndefined() )) {
    *(request ).At(MkString("quantity") )= this.AmountToPrecision(symbol , amount );
  }
  if IsTrue(OpNotEq2(price , MkUndefined() )) {
    *(request ).At(MkString("price") )= this.PriceToPrecision(symbol , price );
  }
  response := this.Call(MkString("privatePatchOrderClientOrderId"), this.Extend(request , params )); _ = response;
  return this.ParseOrder(response );
}

func (this *Hitbtc) CancelOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"clientOrderId":id }); _ = request;
  response := this.Call(MkString("privateDeleteOrderClientOrderId"), this.Extend(request , params )); _ = response;
  return this.ParseOrder(response );
}

func (this *Hitbtc) ParseOrderStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "new":MkString("open") ,
        "suspended":MkString("open") ,
        "partiallyFilled":MkString("open") ,
        "filled":MkString("closed") ,
        "canceled":MkString("canceled") ,
        "expired":MkString("failed") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Hitbtc) ParseOrder(goArgs ...*Variant) *Variant{
  order := GoGetArg(goArgs, 0, MkUndefined()); _ = order;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  created := this.Parse8601(this.SafeString(order , MkString("createdAt") )); _ = created;
  updated := this.Parse8601(this.SafeString(order , MkString("updatedAt") )); _ = updated;
  marketId := this.SafeString(order , MkString("symbol") ); _ = marketId;
  market = this.SafeMarket(marketId , market );
  symbol := *(market ).At(MkString("symbol") ); _ = symbol;
  amount := this.SafeNumber(order , MkString("quantity") ); _ = amount;
  filled := this.SafeNumber(order , MkString("cumQuantity") ); _ = filled;
  status := this.ParseOrderStatus(this.SafeString(order , MkString("status") )); _ = status;
  id := this.SafeString(order , MkString("clientOrderId") ); _ = id;
  clientOrderId := OpCopy(id ); _ = clientOrderId;
  price := this.SafeNumber(order , MkString("price") ); _ = price;
  type_ := this.SafeString(order , MkString("type") ); _ = type_;
  side := this.SafeString(order , MkString("side") ); _ = side;
  trades := this.SafeValue(order , MkString("tradesReport") ); _ = trades;
  fee := OpCopy(MkUndefined() ); _ = fee;
  average := this.SafeNumber(order , MkString("avgPrice") ); _ = average;
  if IsTrue(OpNotEq2(trades , MkUndefined() )) {
    trades = this.ParseTrades(trades , market );
  }
  timeInForce := this.SafeString(order , MkString("timeInForce") ); _ = timeInForce;
  return this.SafeOrder(MkMap(&VarMap{
            "id":id ,
            "clientOrderId":clientOrderId ,
            "timestamp":created ,
            "datetime":this.Iso8601(created ),
            "lastTradeTimestamp":updated ,
            "status":status ,
            "symbol":symbol ,
            "type":type_ ,
            "timeInForce":timeInForce ,
            "side":side ,
            "price":price ,
            "stopPrice":MkUndefined() ,
            "average":average ,
            "amount":amount ,
            "cost":MkUndefined() ,
            "filled":filled ,
            "remaining":MkUndefined() ,
            "fee":fee ,
            "trades":trades ,
            "info":order ,
            }));
}

func (this *Hitbtc) FetchOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"clientOrderId":id }); _ = request;
  response := this.Call(MkString("privateGetHistoryOrder"), this.Extend(request , params )); _ = response;
  numOrders := OpCopy(response.Length ); _ = numOrders;
  if IsTrue(OpGt(numOrders , MkInteger(0) )) {
    return this.ParseOrder(*(response ).At(MkInteger(0) ));
  }
  panic(NewOrderNotFound(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" order ") , OpAdd(id , MkString(" not found") )))));
  return MkUndefined()
}

func (this *Hitbtc) FetchOpenOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"clientOrderId":id }); _ = request;
  response := this.Call(MkString("privateGetOrderClientOrderId"), this.Extend(request , params )); _ = response;
  return this.ParseOrder(response );
}

func (this *Hitbtc) FetchOpenOrders(goArgs ...*Variant) *Variant{
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
  }
  response := this.Call(MkString("privateGetOrder"), this.Extend(request , params )); _ = response;
  return this.ParseOrders(response , market , since , limit );
}

func (this *Hitbtc) FetchClosedOrders(goArgs ...*Variant) *Variant{
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
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("from") )= this.Iso8601(since );
  }
  response := this.Call(MkString("privateGetHistoryOrder"), this.Extend(request , params )); _ = response;
  parsedOrders := this.ParseOrders(response , market ); _ = parsedOrders;
  orders := MkArray(&VarArray{}); _ = orders;
  for i := MkInteger(0) ; IsTrue(OpLw(i , parsedOrders.Length )); OpIncr(& i ){
    order := *(parsedOrders ).At(i ); _ = order;
    status := *(order ).At(MkString("status") ); _ = status;
    if IsTrue(OpOr(OpEq2(status , MkString("closed") ), OpEq2(status , MkString("canceled") ))) {
      orders.Push(order );
    }
  }
  return this.FilterBySinceLimit(orders , since , limit );
}

func (this *Hitbtc) FetchMyTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("symbol") )= *(market ).At(MkString("id") );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("from") )= this.Iso8601(since );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("privateGetHistoryTrades"), this.Extend(request , params )); _ = response;
  return this.ParseTrades(response , market , since , limit );
}

func (this *Hitbtc) FetchOrderTrades(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
  }
  request := MkMap(&VarMap{"orderId":id }); _ = request;
  response := this.Call(MkString("privateGetHistoryOrderOrderIdTrades"), this.Extend(request , params )); _ = response;
  numOrders := OpCopy(response.Length ); _ = numOrders;
  if IsTrue(OpGt(numOrders , MkInteger(0) )) {
    return this.ParseTrades(response , market , since , limit );
  }
  panic(NewOrderNotFound(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" order ") , OpAdd(id , OpAdd(MkString(" not found, ") , OpAdd(*this.At(MkString("id")) , MkString(".fetchOrderTrades() requires an exchange-specific order id, you need to grab it from order[\"info\"][\"id\"]") )))))));
  return MkUndefined()
}

func (this *Hitbtc) CreateDepositAddress(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  currency := this.Currency(code ); _ = currency;
  request := MkMap(&VarMap{"currency":*(currency ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("privatePostAccountCryptoAddressCurrency"), this.Extend(request , params )); _ = response;
  address := this.SafeString(response , MkString("address") ); _ = address;
  this.CheckAddress(address );
  tag := this.SafeString(response , MkString("paymentId") ); _ = tag;
  return MkMap(&VarMap{
        "currency":currency ,
        "address":address ,
        "tag":tag ,
        "info":response ,
        });
}

func (this *Hitbtc) FetchDepositAddress(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  currency := this.Currency(code ); _ = currency;
  request := MkMap(&VarMap{"currency":*(currency ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("privateGetAccountCryptoAddressCurrency"), this.Extend(request , params )); _ = response;
  address := this.SafeString(response , MkString("address") ); _ = address;
  this.CheckAddress(address );
  tag := this.SafeString(response , MkString("paymentId") ); _ = tag;
  return MkMap(&VarMap{
        "currency":*(currency ).At(MkString("code") ),
        "address":address ,
        "tag":tag ,
        "info":response ,
        });
}

func (this *Hitbtc) Withdraw(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  amount := GoGetArg(goArgs, 1, MkUndefined()); _ = amount;
  address := GoGetArg(goArgs, 2, MkUndefined()); _ = address;
  tag := GoGetArg(goArgs, 3, MkUndefined() ); _ = tag;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  this.CheckAddress(address );
  currency := this.Currency(code ); _ = currency;
  request := MkMap(&VarMap{
        "currency":*(currency ).At(MkString("id") ),
        "amount":parseFloat(amount ),
        "address":address ,
        }); _ = request;
  if IsTrue(tag ) {
    *(request ).At(MkString("paymentId") )= OpCopy(tag );
  }
  response := this.Call(MkString("privatePostAccountCryptoWithdraw"), this.Extend(request , params )); _ = response;
  return MkMap(&VarMap{
        "info":response ,
        "id":*(response ).At(MkString("id") ),
        });
}

func (this *Hitbtc) Nonce(goArgs ...*Variant) *Variant{
  return this.Milliseconds();
}

func (this *Hitbtc) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("public") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  url := OpAdd(MkString("/api/") , OpAdd(*this.At(MkString("version")) , MkString("/") )); _ = url;
  query := this.Omit(params , this.ExtractParams(path )); _ = query;
  if IsTrue(OpEq2(api , MkString("public") )) {
    OpAddAssign(& url , OpAdd(api , OpAdd(MkString("/") , this.ImplodeParams(path , params ))));
    if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
      OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(query )));
    }
  } else {
    this.CheckRequiredCredentials();
    OpAddAssign(& url , this.ImplodeParams(path , params ));
    if IsTrue(OpEq2(method , MkString("GET") )) {
      if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
        OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(query )));
      }
    } else {
      if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
        body = this.Json(query );
      }
    }
    payload := this.Encode(OpAdd(*this.At(MkString("apiKey")) , OpAdd(MkString(":") , *this.At(MkString("secret")) ))); _ = payload;
    auth := this.StringToBase64(payload ); _ = auth;
    headers = MkMap(&VarMap{
        "Authorization":OpAdd(MkString("Basic ") , this.Decode(auth )),
        "Content-Type":MkString("application/json") ,
        });
  }
  url = OpAdd(*(*(*this.At(MkString("urls")) ).At(MkString("api") )).At(api ), url );
  return MkMap(&VarMap{
        "url":url ,
        "method":method ,
        "body":body ,
        "headers":headers ,
        });
}

func (this *Hitbtc) HandleErrors(goArgs ...*Variant) *Variant{
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
    feedback := OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , body )); _ = feedback;
    if IsTrue(OpOr(OpEq2(code , MkInteger(503) ), OpEq2(code , MkInteger(504) ))) {
      panic(NewExchangeNotAvailable(feedback ));
    }
    if IsTrue(OpEq2(code , MkInteger(429) )) {
      return MkUndefined();
    }
    if IsTrue(OpEq2(*(body ).At(MkInteger(0) ), MkString("{") )) {
      if IsTrue(OpHasMember(MkString("error") , response )) {
        errorCode := this.SafeString(*(response ).At(MkString("error") ), MkString("code") ); _ = errorCode;
        this.ThrowExactlyMatchedException(*this.At(MkString("exceptions")) , errorCode , feedback );
        message := this.SafeString(*(response ).At(MkString("error") ), MkString("message") ); _ = message;
        if IsTrue(OpEq2(message , MkString("Duplicate clientOrderId") )) {
          panic(NewInvalidOrder(feedback ));
        }
      }
    }
    panic(NewExchangeError(feedback ));
  }
  return MkUndefined()
}

