package ccxt

import (
)

type Zb struct {
	*ExchangeBase
}

var _ Exchange = (*Zb)(nil)

func init() {
	exchange := &Zb{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Zb) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("zb") ,
            "name":MkString("ZB") ,
            "countries":MkArray(&VarArray{
                MkString("CN") ,
                }),
            "rateLimit":MkInteger(100) ,
            "version":MkString("v1") ,
            "certified":MkBool(true) ,
            "pro":MkBool(true) ,
            "has":MkMap(&VarMap{
                "cancelOrder":MkBool(true) ,
                "CORS":MkBool(false) ,
                "createMarketOrder":MkBool(false) ,
                "createOrder":MkBool(true) ,
                "fetchBalance":MkBool(true) ,
                "fetchCurrencies":MkBool(true) ,
                "fetchDepositAddress":MkBool(true) ,
                "fetchDepositAddresses":MkBool(true) ,
                "fetchDeposits":MkBool(true) ,
                "fetchMarkets":MkBool(true) ,
                "fetchOHLCV":MkBool(true) ,
                "fetchOpenOrders":MkBool(true) ,
                "fetchOrder":MkBool(true) ,
                "fetchOrderBook":MkBool(true) ,
                "fetchOrders":MkBool(true) ,
                "fetchClosedOrders":MkBool(true) ,
                "fetchTicker":MkBool(true) ,
                "fetchTickers":MkBool(true) ,
                "fetchTrades":MkBool(true) ,
                "fetchWithdrawals":MkBool(true) ,
                "withdraw":MkBool(true) ,
                }),
            "timeframes":MkMap(&VarMap{
                "1m":MkString("1min") ,
                "3m":MkString("3min") ,
                "5m":MkString("5min") ,
                "15m":MkString("15min") ,
                "30m":MkString("30min") ,
                "1h":MkString("1hour") ,
                "2h":MkString("2hour") ,
                "4h":MkString("4hour") ,
                "6h":MkString("6hour") ,
                "12h":MkString("12hour") ,
                "1d":MkString("1day") ,
                "3d":MkString("3day") ,
                "1w":MkString("1week") ,
                }),
            "exceptions":MkMap(&VarMap{
                "ws":MkMap(&VarMap{
                    "1001":ExchangeError ,
                    "1002":ExchangeError ,
                    "1003":AuthenticationError ,
                    "1004":AuthenticationError ,
                    "1005":AuthenticationError ,
                    "1006":PermissionDenied ,
                    "1007":ExchangeError ,
                    "1009":OnMaintenance ,
                    "1010":ExchangeNotAvailable ,
                    "1012":PermissionDenied ,
                    "1013":ExchangeError ,
                    "1014":ExchangeError ,
                    "2001":InsufficientFunds ,
                    "2002":InsufficientFunds ,
                    "2003":InsufficientFunds ,
                    "2005":InsufficientFunds ,
                    "2006":InsufficientFunds ,
                    "2007":InsufficientFunds ,
                    "2008":InsufficientFunds ,
                    "2009":InsufficientFunds ,
                    "3001":OrderNotFound ,
                    "3002":InvalidOrder ,
                    "3003":InvalidOrder ,
                    "3004":AuthenticationError ,
                    "3005":BadRequest ,
                    "3006":PermissionDenied ,
                    "3007":RequestTimeout ,
                    "3008":ExchangeError ,
                    "3009":InvalidOrder ,
                    "3010":PermissionDenied ,
                    "3011":InvalidOrder ,
                    "3012":InvalidOrder ,
                    "4001":AccountSuspended ,
                    "4002":RateLimitExceeded ,
                    }),
                "exact":MkMap(&VarMap{
                    "1001":ExchangeError ,
                    "1002":ExchangeError ,
                    "1003":AuthenticationError ,
                    "1004":AuthenticationError ,
                    "1005":AuthenticationError ,
                    "1006":AuthenticationError ,
                    "1009":ExchangeNotAvailable ,
                    "1010":ExchangeNotAvailable ,
                    "1012":PermissionDenied ,
                    "1013":ExchangeError ,
                    "1014":ExchangeError ,
                    "2001":InsufficientFunds ,
                    "2002":InsufficientFunds ,
                    "2003":InsufficientFunds ,
                    "2005":InsufficientFunds ,
                    "2006":InsufficientFunds ,
                    "2007":InsufficientFunds ,
                    "2008":InsufficientFunds ,
                    "2009":InsufficientFunds ,
                    "3001":OrderNotFound ,
                    "3002":InvalidOrder ,
                    "3003":InvalidOrder ,
                    "3004":AuthenticationError ,
                    "3005":BadRequest ,
                    "3006":AuthenticationError ,
                    "3007":AuthenticationError ,
                    "3008":OrderNotFound ,
                    "3009":InvalidOrder ,
                    "3010":PermissionDenied ,
                    "3011":InvalidOrder ,
                    "3012":InvalidOrder ,
                    "4001":ExchangeNotAvailable ,
                    "4002":RateLimitExceeded ,
                    }),
                "broad":MkMap(&VarMap{"æå¸å°åæè¯¯ï¼è¯·åæ·»å æå¸å°åã":InvalidAddress }),
                }),
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/1294454/32859187-cd5214f0-ca5e-11e7-967d-96568e2e2bd1.jpg") ,
                "api":MkMap(&VarMap{
                    "public":MkString("https://api.zb.today/data") ,
                    "private":MkString("https://trade.zb.today/api") ,
                    "trade":MkString("https://trade.zb.today/api") ,
                    }),
                "www":MkString("https://www.zb.com") ,
                "doc":MkString("https://www.zb.com/i/developer") ,
                "fees":MkString("https://www.zb.com/i/rate") ,
                "referral":MkMap(&VarMap{
                    "url":MkString("https://www.zbex.club/en/register?ref=4301lera") ,
                    "discount":MkNumber(0.16) ,
                    }),
                }),
            "api":MkMap(&VarMap{
                "trade":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("getFeeInfo") ,
                        })}),
                "public":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("markets") ,
                        MkString("ticker") ,
                        MkString("allTicker") ,
                        MkString("depth") ,
                        MkString("trades") ,
                        MkString("kline") ,
                        MkString("getGroupMarkets") ,
                        })}),
                "private":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("order") ,
                        MkString("orderMoreV2") ,
                        MkString("cancelOrder") ,
                        MkString("getOrder") ,
                        MkString("getOrders") ,
                        MkString("getOrdersNew") ,
                        MkString("getOrdersIgnoreTradeType") ,
                        MkString("getUnfinishedOrdersIgnoreTradeType") ,
                        MkString("getFinishedAndPartialOrders") ,
                        MkString("getAccountInfo") ,
                        MkString("getUserAddress") ,
                        MkString("getPayinAddress") ,
                        MkString("getWithdrawAddress") ,
                        MkString("getWithdrawRecord") ,
                        MkString("getChargeRecord") ,
                        MkString("getCnyWithdrawRecord") ,
                        MkString("getCnyChargeRecord") ,
                        MkString("withdraw") ,
                        MkString("addSubUser") ,
                        MkString("getSubUserList") ,
                        MkString("doTransferFunds") ,
                        MkString("createSubUserKey") ,
                        MkString("getLeverAssetsInfo") ,
                        MkString("getLeverBills") ,
                        MkString("transferInLever") ,
                        MkString("transferOutLever") ,
                        MkString("loan") ,
                        MkString("cancelLoan") ,
                        MkString("getLoans") ,
                        MkString("getLoanRecords") ,
                        MkString("borrow") ,
                        MkString("autoBorrow") ,
                        MkString("repay") ,
                        MkString("doAllRepay") ,
                        MkString("getRepayments") ,
                        MkString("getFinanceRecords") ,
                        MkString("changeInvestMark") ,
                        MkString("changeLoop") ,
                        MkString("getCrossAssets") ,
                        MkString("getCrossBills") ,
                        MkString("transferInCross") ,
                        MkString("transferOutCross") ,
                        MkString("doCrossLoan") ,
                        MkString("doCrossRepay") ,
                        MkString("getCrossRepayRecords") ,
                        })}),
                }),
            "fees":MkMap(&VarMap{
                "funding":MkMap(&VarMap{"withdraw":MkMap(&VarMap{})}),
                "trading":MkMap(&VarMap{
                    "maker":OpDiv(MkNumber(0.2) , MkInteger(100) ),
                    "taker":OpDiv(MkNumber(0.2) , MkInteger(100) ),
                    }),
                }),
            "commonCurrencies":MkMap(&VarMap{
                "ANG":MkString("Anagram") ,
                "ENT":MkString("ENTCash") ,
                "BCHABC":MkString("BCHABC") ,
                "BCHSV":MkString("BCHSV") ,
                }),
            }));
}

func (this *Zb) FetchMarkets(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  markets := this.Call(MkString("publicGetMarkets"), params ); _ = markets;
  keys := GoGetKeys(markets ); _ = keys;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , keys.Length )); OpIncr(& i ){
    id := *(keys ).At(i ); _ = id;
    market := *(markets ).At(id ); _ = market;
    baseId := MkUndefined(); _ = baseId;
    quoteId := MkUndefined(); _ = quoteId;
    ArrayUnpack(id.Split(MkString("_") ), &baseId, &quoteId);
    base := this.SafeCurrencyCode(baseId ); _ = base;
    quote := this.SafeCurrencyCode(quoteId ); _ = quote;
    symbol := OpAdd(base , OpAdd(MkString("/") , quote )); _ = symbol;
    amountPrecisionString := this.SafeString(market , MkString("amountScale") ); _ = amountPrecisionString;
    pricePrecisionString := this.SafeString(market , MkString("priceScale") ); _ = pricePrecisionString;
    amountLimit := this.ParsePrecision(amountPrecisionString ); _ = amountLimit;
    priceLimit := this.ParsePrecision(pricePrecisionString ); _ = priceLimit;
    precision := MkMap(&VarMap{
          "amount":parseInt(amountPrecisionString ),
          "price":parseInt(pricePrecisionString ),
          }); _ = precision;
    result.Push(MkMap(&VarMap{
            "id":id ,
            "symbol":symbol ,
            "baseId":baseId ,
            "quoteId":quoteId ,
            "base":base ,
            "quote":quote ,
            "active":MkBool(true) ,
            "precision":precision ,
            "limits":MkMap(&VarMap{
                "amount":MkMap(&VarMap{
                    "min":this.ParseNumber(amountLimit ),
                    "max":MkUndefined() ,
                    }),
                "price":MkMap(&VarMap{
                    "min":this.ParseNumber(priceLimit ),
                    "max":MkUndefined() ,
                    }),
                "cost":MkMap(&VarMap{
                    "min":MkInteger(0) ,
                    "max":MkUndefined() ,
                    }),
                }),
            "info":market ,
            }));
  }
  return result ;
}

func (this *Zb) FetchCurrencies(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("tradeGetGetFeeInfo"), params ); _ = response;
  currencies := this.SafeValue(response , MkString("result") , MkMap(&VarMap{})); _ = currencies;
  ids := GoGetKeys(currencies ); _ = ids;
  result := MkMap(&VarMap{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , ids.Length )); OpIncr(& i ){
    id := *(ids ).At(i ); _ = id;
    currency := *(currencies ).At(id ); _ = currency;
    code := this.SafeCurrencyCode(id ); _ = code;
    precision := OpCopy(MkUndefined() ); _ = precision;
    isWithdrawEnabled := OpCopy(MkBool(true) ); _ = isWithdrawEnabled;
    isDepositEnabled := OpCopy(MkBool(true) ); _ = isDepositEnabled;
    fees := MkMap(&VarMap{}); _ = fees;
    for j := MkInteger(0) ; IsTrue(OpLw(j , currency.Length )); OpIncr(& j ){
      networkItem := *(currency ).At(j ); _ = networkItem;
      network := this.SafeString(networkItem , MkString("chainName") ); _ = network;
      withdrawFee := this.SafeNumber(networkItem , MkString("fee") ); _ = withdrawFee;
      depositEnable := this.SafeValue(networkItem , MkString("canDeposit") ); _ = depositEnable;
      withdrawEnable := this.SafeValue(networkItem , MkString("canWithdraw") ); _ = withdrawEnable;
      isDepositEnabled = OpOr(isDepositEnabled , depositEnable );
      isWithdrawEnabled = OpOr(isWithdrawEnabled , withdrawEnable );
      *(fees ).At(network )= OpCopy(withdrawFee );
    }
    active := OpAnd(isWithdrawEnabled , isDepositEnabled ); _ = active;
    *(result ).At(code )= MkMap(&VarMap{
        "id":id ,
        "name":MkUndefined() ,
        "code":code ,
        "precision":precision ,
        "info":currency ,
        "active":active ,
        "fee":MkUndefined() ,
        "fees":fees ,
        "limits":*this.At(MkString("limits")) ,
        });
  }
  return result ;
}

func (this *Zb) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("privateGetGetAccountInfo"), params ); _ = response;
  balances := this.SafeValue(*(response ).At(MkString("result") ), MkString("coins") ); _ = balances;
  result := MkMap(&VarMap{
        "info":response ,
        "timestamp":MkUndefined() ,
        "datetime":MkUndefined() ,
        }); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , balances.Length )); OpIncr(& i ){
    balance := *(balances ).At(i ); _ = balance;
    account := this.Account(); _ = account;
    currencyId := this.SafeString(balance , MkString("key") ); _ = currencyId;
    code := this.SafeCurrencyCode(currencyId ); _ = code;
    *(account ).At(MkString("free") )= this.SafeString(balance , MkString("available") );
    *(account ).At(MkString("used") )= this.SafeString(balance , MkString("freez") );
    *(result ).At(code )= OpCopy(account );
  }
  return this.ParseBalance(result );
}

func (this *Zb) ParseDepositAddress(goArgs ...*Variant) *Variant{
  depositAddress := GoGetArg(goArgs, 0, MkUndefined()); _ = depositAddress;
  currency := GoGetArg(goArgs, 1, MkUndefined() ); _ = currency;
  address := this.SafeString2(depositAddress , MkString("key") , MkString("address") ); _ = address;
  tag := OpCopy(MkUndefined() ); _ = tag;
  memo := this.SafeString(depositAddress , MkString("memo") ); _ = memo;
  if IsTrue(OpNotEq2(memo , MkUndefined() )) {
    tag = OpCopy(memo );
  } else {
    if IsTrue(OpGtEq(address.IndexOf(MkString("_") ), MkInteger(0) )) {
      parts := address.Split(MkString("_") ); _ = parts;
      address = *(parts ).At(MkInteger(0) );
      tag = *(parts ).At(MkInteger(1) );
    }
  }
  currencyId := this.SafeString(depositAddress , MkString("blockChain") ); _ = currencyId;
  code := this.SafeCurrencyCode(currencyId , currency ); _ = code;
  return MkMap(&VarMap{
        "currency":code ,
        "address":address ,
        "tag":tag ,
        "info":depositAddress ,
        });
}

func (this *Zb) FetchDepositAddresses(goArgs ...*Variant) *Variant{
  codes := GoGetArg(goArgs, 0, MkUndefined() ); _ = codes;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("privateGetGetPayinAddress"), params ); _ = response;
  message := this.SafeValue(response , MkString("message") , MkMap(&VarMap{})); _ = message;
  datas := this.SafeValue(message , MkString("datas") , MkArray(&VarArray{})); _ = datas;
  return this.ParseDepositAddresses(datas , codes );
}

func (this *Zb) FetchDepositAddress(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  currency := this.Currency(code ); _ = currency;
  request := MkMap(&VarMap{"currency":*(currency ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("privateGetGetUserAddress"), this.Extend(request , params )); _ = response;
  message := this.SafeValue(response , MkString("message") , MkMap(&VarMap{})); _ = message;
  datas := this.SafeValue(message , MkString("datas") , MkMap(&VarMap{})); _ = datas;
  return this.ParseDepositAddress(datas , currency );
}

func (this *Zb) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"market":*(market ).At(MkString("id") )}); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("size") )= OpCopy(limit );
  }
  response := this.Call(MkString("publicGetDepth"), this.Extend(request , params )); _ = response;
  return this.ParseOrderBook(response , symbol );
}

func (this *Zb) FetchTickers(goArgs ...*Variant) *Variant{
  symbols := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbols;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("publicGetAllTicker"), params ); _ = response;
  result := MkMap(&VarMap{}); _ = result;
  marketsByIdWithoutUnderscore := MkMap(&VarMap{}); _ = marketsByIdWithoutUnderscore;
  marketIds := GoGetKeys(*this.At(MkString("markets_by_id")) ); _ = marketIds;
  for i := MkInteger(0) ; IsTrue(OpLw(i , marketIds.Length )); OpIncr(& i ){
    tickerId := (*(marketIds ).At(i )).Call(MkString("replace") , MkString("_") , MkString("") ); _ = tickerId;
    *(marketsByIdWithoutUnderscore ).At(tickerId )= *(*this.At(MkString("markets_by_id")) ).At(*(marketIds ).At(i ));
  }
  ids := GoGetKeys(response ); _ = ids;
  for i := MkInteger(0) ; IsTrue(OpLw(i , ids.Length )); OpIncr(& i ){
    market := *(marketsByIdWithoutUnderscore ).At(*(ids ).At(i )); _ = market;
    *(result ).At(*(market ).At(MkString("symbol") ))= this.ParseTicker(*(response ).At(*(ids ).At(i )), market );
  }
  return this.FilterByArray(result , MkString("symbol") , symbols );
}

func (this *Zb) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"market":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("publicGetTicker"), this.Extend(request , params )); _ = response;
  ticker := this.SafeValue(response , MkString("ticker") , MkMap(&VarMap{})); _ = ticker;
  *(ticker ).At(MkString("date") )= this.SafeValue(response , MkString("date") );
  return this.ParseTicker(ticker , market );
}

func (this *Zb) ParseTicker(goArgs ...*Variant) *Variant{
  ticker := GoGetArg(goArgs, 0, MkUndefined()); _ = ticker;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.SafeInteger(ticker , MkString("date") , this.Milliseconds()); _ = timestamp;
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  if IsTrue(OpNotEq2(market , MkUndefined() )) {
    symbol = *(market ).At(MkString("symbol") );
  }
  last := this.SafeNumber(ticker , MkString("last") ); _ = last;
  return MkMap(&VarMap{
        "symbol":symbol ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "high":this.SafeNumber(ticker , MkString("high") ),
        "low":this.SafeNumber(ticker , MkString("low") ),
        "bid":this.SafeNumber(ticker , MkString("buy") ),
        "bidVolume":MkUndefined() ,
        "ask":this.SafeNumber(ticker , MkString("sell") ),
        "askVolume":MkUndefined() ,
        "vwap":MkUndefined() ,
        "open":MkUndefined() ,
        "close":last ,
        "last":last ,
        "previousClose":MkUndefined() ,
        "change":MkUndefined() ,
        "percentage":MkUndefined() ,
        "average":MkUndefined() ,
        "baseVolume":this.SafeNumber(ticker , MkString("vol") ),
        "quoteVolume":MkUndefined() ,
        "info":ticker ,
        });
}

func (this *Zb) ParseOHLCV(goArgs ...*Variant) *Variant{
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

func (this *Zb) FetchOHLCV(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  timeframe := GoGetArg(goArgs, 1, MkString("1m") ); _ = timeframe;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  if IsTrue(OpEq2(limit , MkUndefined() )) {
    limit = MkInteger(1000) ;
  }
  request := MkMap(&VarMap{
        "market":*(market ).At(MkString("id") ),
        "type":*(*this.At(MkString("timeframes")) ).At(timeframe ),
        "limit":limit ,
        }); _ = request;
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("since") )= OpCopy(since );
  }
  response := this.Call(MkString("publicGetKline"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = data;
  return this.ParseOHLCVs(data , market , timeframe , since , limit );
}

func (this *Zb) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.SafeTimestamp(trade , MkString("date") ); _ = timestamp;
  side := this.SafeString(trade , MkString("trade_type") ); _ = side;
  side = OpTrinary(OpEq2(side , MkString("bid") ), MkString("buy") , MkString("sell") );
  id := this.SafeString(trade , MkString("tid") ); _ = id;
  priceString := this.SafeString(trade , MkString("price") ); _ = priceString;
  amountString := this.SafeString(trade , MkString("amount") ); _ = amountString;
  costString := Precise.StringMul(priceString , amountString ); _ = costString;
  price := this.ParseNumber(priceString ); _ = price;
  amount := this.ParseNumber(amountString ); _ = amount;
  cost := this.ParseNumber(costString ); _ = cost;
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  if IsTrue(OpNotEq2(market , MkUndefined() )) {
    symbol = *(market ).At(MkString("symbol") );
  }
  return MkMap(&VarMap{
        "info":trade ,
        "id":id ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "symbol":symbol ,
        "type":MkUndefined() ,
        "side":side ,
        "order":MkUndefined() ,
        "takerOrMaker":MkUndefined() ,
        "price":price ,
        "amount":amount ,
        "cost":cost ,
        "fee":MkUndefined() ,
        });
}

func (this *Zb) FetchTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"market":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("publicGetTrades"), this.Extend(request , params )); _ = response;
  return this.ParseTrades(response , market , since , limit );
}

func (this *Zb) CreateOrder(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 2, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 3, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 4, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 5, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpNotEq2(type_ , MkString("limit") )) {
    panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")) , MkString(" allows limit orders only") )));
  }
  this.LoadMarkets();
  request := MkMap(&VarMap{
        "price":this.PriceToPrecision(symbol , price ),
        "amount":this.AmountToPrecision(symbol , amount ),
        "tradeType":OpTrinary(OpEq2(side , MkString("buy") ), MkString("1") , MkString("0") ),
        "currency":this.MarketId(symbol ),
        }); _ = request;
  response := this.Call(MkString("privateGetOrder"), this.Extend(request , params )); _ = response;
  return MkMap(&VarMap{
        "info":response ,
        "id":*(response ).At(MkString("id") ),
        });
}

func (this *Zb) CancelOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{
        "id":id.ToString(),
        "currency":this.MarketId(symbol ),
        }); _ = request;
  return this.Call(MkString("privateGetCancelOrder"), this.Extend(request , params ));
}

func (this *Zb) FetchOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchOrder() requires a symbol argument") )));
  }
  this.LoadMarkets();
  request := MkMap(&VarMap{
        "id":id.ToString(),
        "currency":this.MarketId(symbol ),
        }); _ = request;
  response := this.Call(MkString("privateGetGetOrder"), this.Extend(request , params )); _ = response;
  return this.ParseOrder(response , MkUndefined() );
}

func (this *Zb) FetchOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkInteger(50) ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString("fetchOrders() requires a symbol argument") )));
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "currency":*(market ).At(MkString("id") ),
        "pageIndex":MkInteger(1) ,
        "pageSize":limit ,
        }); _ = request;
  method := MkString("privateGetGetOrdersIgnoreTradeType") ; _ = method;
  if IsTrue(OpHasMember(MkString("tradeType") , params )) {
    method = MkString("privateGetGetOrdersNew") ;
  }
  response := OpCopy(MkUndefined() ); _ = response;
  {
  ret__ := func(this *Zb) (ret_ *Variant) {
    defer func() {
      if e := recover().(*Variant); e != nil {
        ret_ = func(this *Zb) *Variant {
          // catch block:
          if IsTrue(OpIsType(e , OrderNotFound )) {
            return MkArray(&VarArray{});
          }
          panic(e );
          return nil
        }(this)
      }
    }()
    // try block:
    response = this.Call(method , this.Extend(request , params ));
    return nil
  }(this)
  if ret__ != nil {
    return ret__;
   }
  }
  return this.ParseOrders(response , market , since , limit );
}

func (this *Zb) FetchClosedOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString("fetchClosedOrders() requires a symbol argument") )));
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "currency":*(market ).At(MkString("id") ),
        "pageIndex":MkInteger(1) ,
        "pageSize":MkInteger(10) ,
        }); _ = request;
  response := this.Call(MkString("privateGetGetFinishedAndPartialOrders"), this.Extend(request , params )); _ = response;
  return this.ParseOrders(response , market , since , limit );
}

func (this *Zb) FetchOpenOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkInteger(10) ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString("fetchOpenOrders() requires a symbol argument") )));
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "currency":*(market ).At(MkString("id") ),
        "pageIndex":MkInteger(1) ,
        "pageSize":limit ,
        }); _ = request;
  method := MkString("privateGetGetUnfinishedOrdersIgnoreTradeType") ; _ = method;
  if IsTrue(OpHasMember(MkString("tradeType") , params )) {
    method = MkString("privateGetGetOrdersNew") ;
  }
  response := OpCopy(MkUndefined() ); _ = response;
  {
  ret__ := func(this *Zb) (ret_ *Variant) {
    defer func() {
      if e := recover().(*Variant); e != nil {
        ret_ = func(this *Zb) *Variant {
          // catch block:
          if IsTrue(OpIsType(e , OrderNotFound )) {
            return MkArray(&VarArray{});
          }
          panic(e );
          return nil
        }(this)
      }
    }()
    // try block:
    response = this.Call(method , this.Extend(request , params ));
    return nil
  }(this)
  if ret__ != nil {
    return ret__;
   }
  }
  return this.ParseOrders(response , market , since , limit );
}

func (this *Zb) ParseOrder(goArgs ...*Variant) *Variant{
  order := GoGetArg(goArgs, 0, MkUndefined()); _ = order;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  side := this.SafeInteger(order , MkString("type") ); _ = side;
  side = OpTrinary(OpEq2(side , MkInteger(1) ), MkString("buy") , MkString("sell") );
  type_ := MkString("limit") ; _ = type_;
  timestamp := this.SafeInteger(order , MkString("trade_date") ); _ = timestamp;
  marketId := this.SafeString(order , MkString("currency") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market , MkString("_") ); _ = symbol;
  price := this.SafeNumber(order , MkString("price") ); _ = price;
  filled := this.SafeNumber(order , MkString("trade_amount") ); _ = filled;
  amount := this.SafeNumber(order , MkString("total_amount") ); _ = amount;
  cost := this.SafeNumber(order , MkString("trade_money") ); _ = cost;
  status := this.ParseOrderStatus(this.SafeString(order , MkString("status") )); _ = status;
  id := this.SafeString(order , MkString("id") ); _ = id;
  feeCost := this.SafeNumber(order , MkString("fees") ); _ = feeCost;
  fee := OpCopy(MkUndefined() ); _ = fee;
  if IsTrue(OpNotEq2(feeCost , MkUndefined() )) {
    feeCurrency := OpCopy(MkUndefined() ); _ = feeCurrency;
    zbFees := this.SafeValue(order , MkString("useZbFee") ); _ = zbFees;
    if IsTrue(OpEq2(zbFees , MkBool(true) )) {
      feeCurrency = MkString("ZB") ;
    } else {
      if IsTrue(OpNotEq2(market , MkUndefined() )) {
        feeCurrency = OpTrinary(OpEq2(side , MkString("sell") ), *(market ).At(MkString("quote") ), *(market ).At(MkString("base") ));
      }
    }
    fee = MkMap(&VarMap{
        "cost":feeCost ,
        "currency":feeCurrency ,
        });
  }
  return this.SafeOrder(MkMap(&VarMap{
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
            "stopPrice":MkUndefined() ,
            "average":MkUndefined() ,
            "cost":cost ,
            "amount":amount ,
            "filled":filled ,
            "remaining":MkUndefined() ,
            "status":status ,
            "fee":fee ,
            "trades":MkUndefined() ,
            }));
}

func (this *Zb) ParseOrderStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "0":MkString("open") ,
        "1":MkString("canceled") ,
        "2":MkString("closed") ,
        "3":MkString("open") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Zb) ParseTransactionStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "0":MkString("pending") ,
        "1":MkString("failed") ,
        "2":MkString("ok") ,
        "3":MkString("canceled") ,
        "5":MkString("ok") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Zb) ParseTransaction(goArgs ...*Variant) *Variant{
  transaction := GoGetArg(goArgs, 0, MkUndefined()); _ = transaction;
  currency := GoGetArg(goArgs, 1, MkUndefined() ); _ = currency;
  id := this.SafeString(transaction , MkString("id") ); _ = id;
  txid := this.SafeString(transaction , MkString("hash") ); _ = txid;
  amount := this.SafeNumber(transaction , MkString("amount") ); _ = amount;
  timestamp := this.Parse8601(this.SafeString(transaction , MkString("submit_time") )); _ = timestamp;
  timestamp = this.SafeInteger(transaction , MkString("submitTime") , timestamp );
  address := this.SafeString2(transaction , MkString("toAddress") , MkString("address") ); _ = address;
  tag := OpCopy(MkUndefined() ); _ = tag;
  if IsTrue(OpNotEq2(address , MkUndefined() )) {
    parts := address.Split(MkString("_") ); _ = parts;
    address = this.SafeString(parts , MkInteger(0) );
    tag = this.SafeString(parts , MkInteger(1) );
  }
  confirmTimes := this.SafeInteger(transaction , MkString("confirmTimes") ); _ = confirmTimes;
  updated := this.SafeInteger(transaction , MkString("manageTime") ); _ = updated;
  type_ := OpCopy(MkUndefined() ); _ = type_;
  currencyId := this.SafeString(transaction , MkString("currency") ); _ = currencyId;
  code := this.SafeCurrencyCode(currencyId , currency ); _ = code;
  if IsTrue(OpNotEq2(address , MkUndefined() )) {
    type_ = OpTrinary(OpEq2(confirmTimes , MkUndefined() ), MkString("withdrawal") , MkString("deposit") );
  }
  status := this.ParseTransactionStatus(this.SafeString(transaction , MkString("status") )); _ = status;
  fee := OpCopy(MkUndefined() ); _ = fee;
  feeCost := this.SafeNumber(transaction , MkString("fees") ); _ = feeCost;
  if IsTrue(OpNotEq2(feeCost , MkUndefined() )) {
    fee = MkMap(&VarMap{
        "cost":feeCost ,
        "currency":code ,
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

func (this *Zb) Withdraw(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  amount := GoGetArg(goArgs, 1, MkUndefined()); _ = amount;
  address := GoGetArg(goArgs, 2, MkUndefined()); _ = address;
  tag := GoGetArg(goArgs, 3, MkUndefined() ); _ = tag;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  password := this.SafeString(params , MkString("safePwd") , *this.At(MkString("password")) ); _ = password;
  if IsTrue(OpEq2(password , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" withdraw() requires exchange.password or a safePwd parameter") )));
  }
  fees := this.SafeNumber(params , MkString("fees") ); _ = fees;
  if IsTrue(OpEq2(fees , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" withdraw() requires a fees parameter") )));
  }
  this.CheckAddress(address );
  this.LoadMarkets();
  currency := this.Currency(code ); _ = currency;
  if IsTrue(OpNotEq2(tag , MkUndefined() )) {
    OpAddAssign(& address , OpAdd(MkString("_") , tag ));
  }
  request := MkMap(&VarMap{
        "amount":this.CurrencyToPrecision(code , amount ),
        "currency":*(currency ).At(MkString("id") ),
        "fees":this.CurrencyToPrecision(code , fees ),
        "method":MkString("withdraw") ,
        "receiveAddr":address ,
        "safePwd":password ,
        }); _ = request;
  response := this.Call(MkString("privateGetWithdraw"), this.Extend(request , params )); _ = response;
  transaction := this.ParseTransaction(response , currency ); _ = transaction;
  return this.Extend(transaction , MkMap(&VarMap{
            "type":MkString("withdrawal") ,
            "address":address ,
            "addressTo":address ,
            "amount":amount ,
            }));
}

func (this *Zb) FetchWithdrawals(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  currency := OpCopy(MkUndefined() ); _ = currency;
  if IsTrue(OpNotEq2(code , MkUndefined() )) {
    currency = this.Currency(code );
    *(request ).At(MkString("currency") )= *(currency ).At(MkString("id") );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("pageSize") )= OpCopy(limit );
  }
  response := this.Call(MkString("privateGetGetWithdrawRecord"), this.Extend(request , params )); _ = response;
  message := this.SafeValue(response , MkString("message") , MkMap(&VarMap{})); _ = message;
  datas := this.SafeValue(message , MkString("datas") , MkMap(&VarMap{})); _ = datas;
  withdrawals := this.SafeValue(datas , MkString("list") , MkArray(&VarArray{})); _ = withdrawals;
  return this.ParseTransactions(withdrawals , currency , since , limit );
}

func (this *Zb) FetchDeposits(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  currency := OpCopy(MkUndefined() ); _ = currency;
  if IsTrue(OpNotEq2(code , MkUndefined() )) {
    currency = this.Currency(code );
    *(request ).At(MkString("currency") )= *(currency ).At(MkString("id") );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("pageSize") )= OpCopy(limit );
  }
  response := this.Call(MkString("privateGetGetChargeRecord"), this.Extend(request , params )); _ = response;
  message := this.SafeValue(response , MkString("message") , MkMap(&VarMap{})); _ = message;
  datas := this.SafeValue(message , MkString("datas") , MkMap(&VarMap{})); _ = datas;
  deposits := this.SafeValue(datas , MkString("list") , MkArray(&VarArray{})); _ = deposits;
  return this.ParseTransactions(deposits , currency , since , limit );
}

func (this *Zb) Nonce(goArgs ...*Variant) *Variant{
  return this.Milliseconds();
}

func (this *Zb) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("public") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  url := *(*(*this.At(MkString("urls")) ).At(MkString("api") )).At(api ); _ = url;
  if IsTrue(OpEq2(api , MkString("public") )) {
    OpAddAssign(& url , OpAdd(MkString("/") , OpAdd(*this.At(MkString("version")) , OpAdd(MkString("/") , path ))));
    if IsTrue(*(GoGetKeys(params )).At(MkString("length") )) {
      OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(params )));
    }
  } else {
    if IsTrue(OpEq2(api , MkString("trade") )) {
      OpAddAssign(& url , OpAdd(MkString("/") , path ));
      if IsTrue(*(GoGetKeys(params )).At(MkString("length") )) {
        OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(params )));
      }
    } else {
      query := this.Keysort(this.Extend(MkMap(&VarMap{
                    "method":path ,
                    "accesskey":*this.At(MkString("apiKey")) ,
                    }), params )); _ = query;
      nonce := this.Nonce(); _ = nonce;
      query = this.Keysort(query );
      auth := this.Rawencode(query ); _ = auth;
      secret := this.Hash(this.Encode(*this.At(MkString("secret")) ), MkString("sha1") ); _ = secret;
      signature := this.Hmac(this.Encode(auth ), this.Encode(secret ), MkString("md5") ); _ = signature;
      suffix := OpAdd(MkString("sign=") , OpAdd(signature , OpAdd(MkString("&reqTime=") , nonce.ToString()))); _ = suffix;
      OpAddAssign(& url , OpAdd(MkString("/") , OpAdd(path , OpAdd(MkString("?") , OpAdd(auth , OpAdd(MkString("&") , suffix ))))));
    }
  }
  return MkMap(&VarMap{
        "url":url ,
        "method":method ,
        "body":body ,
        "headers":headers ,
        });
}

func (this *Zb) HandleErrors(goArgs ...*Variant) *Variant{
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
  if IsTrue(OpEq2(*(body ).At(MkInteger(0) ), MkString("{") )) {
    feedback := OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , body )); _ = feedback;
    this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("broad") ), body , feedback );
    if IsTrue(OpHasMember(MkString("code") , response )) {
      code := this.SafeString(response , MkString("code") ); _ = code;
      this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("exact") ), code , feedback );
      if IsTrue(OpNotEq2(code , MkString("1000") )) {
        panic(NewExchangeError(feedback ));
      }
    }
    result := this.SafeValue(response , MkString("result") ); _ = result;
    if IsTrue(OpNotEq2(result , MkUndefined() )) {
      if IsTrue(OpNot(result )) {
        message := this.SafeString(response , MkString("message") ); _ = message;
        if IsTrue(OpEq2(message , MkString("æå¡ç«¯å¿ç¢") )) {
          panic(NewExchangeNotAvailable(feedback ));
        } else {
          panic(NewExchangeError(feedback ));
        }
      }
    }
  }
  return MkUndefined()
}

