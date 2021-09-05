package ccxt

import (
)

type Independentreserve struct {
	*ExchangeBase
}

var _ Exchange = (*Independentreserve)(nil)

func init() {
	exchange := &Independentreserve{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Independentreserve) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("independentreserve") ,
            "name":MkString("Independent Reserve") ,
            "countries":MkArray(&VarArray{
                MkString("AU") ,
                MkString("NZ") ,
                }),
            "rateLimit":MkInteger(1000) ,
            "has":MkMap(&VarMap{
                "cancelOrder":MkBool(true) ,
                "CORS":MkBool(false) ,
                "createOrder":MkBool(true) ,
                "fetchBalance":MkBool(true) ,
                "fetchClosedOrders":MkBool(true) ,
                "fetchMarkets":MkBool(true) ,
                "fetchMyTrades":MkBool(true) ,
                "fetchOpenOrders":MkBool(true) ,
                "fetchOrder":MkBool(true) ,
                "fetchOrderBook":MkBool(true) ,
                "fetchTicker":MkBool(true) ,
                "fetchTrades":MkBool(true) ,
                }),
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/51840849/87182090-1e9e9080-c2ec-11ea-8e49-563db9a38f37.jpg") ,
                "api":MkMap(&VarMap{
                    "public":MkString("https://api.independentreserve.com/Public") ,
                    "private":MkString("https://api.independentreserve.com/Private") ,
                    }),
                "www":MkString("https://www.independentreserve.com") ,
                "doc":MkString("https://www.independentreserve.com/API") ,
                }),
            "api":MkMap(&VarMap{
                "public":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("GetValidPrimaryCurrencyCodes") ,
                        MkString("GetValidSecondaryCurrencyCodes") ,
                        MkString("GetValidLimitOrderTypes") ,
                        MkString("GetValidMarketOrderTypes") ,
                        MkString("GetValidOrderTypes") ,
                        MkString("GetValidTransactionTypes") ,
                        MkString("GetMarketSummary") ,
                        MkString("GetOrderBook") ,
                        MkString("GetAllOrders") ,
                        MkString("GetTradeHistorySummary") ,
                        MkString("GetRecentTrades") ,
                        MkString("GetFxRates") ,
                        MkString("GetOrderMinimumVolumes") ,
                        MkString("GetCryptoWithdrawalFees") ,
                        })}),
                "private":MkMap(&VarMap{"post":MkArray(&VarArray{
                        MkString("GetOpenOrders") ,
                        MkString("GetClosedOrders") ,
                        MkString("GetClosedFilledOrders") ,
                        MkString("GetOrderDetails") ,
                        MkString("GetAccounts") ,
                        MkString("GetTransactions") ,
                        MkString("GetFiatBankAccounts") ,
                        MkString("GetDigitalCurrencyDepositAddress") ,
                        MkString("GetDigitalCurrencyDepositAddresses") ,
                        MkString("GetTrades") ,
                        MkString("GetBrokerageFees") ,
                        MkString("GetDigitalCurrencyWithdrawal") ,
                        MkString("PlaceLimitOrder") ,
                        MkString("PlaceMarketOrder") ,
                        MkString("CancelOrder") ,
                        MkString("SynchDigitalCurrencyDepositAddressWithBlockchain") ,
                        MkString("RequestFiatWithdrawal") ,
                        MkString("WithdrawFiatCurrency") ,
                        MkString("WithdrawDigitalCurrency") ,
                        })}),
                }),
            "fees":MkMap(&VarMap{"trading":MkMap(&VarMap{
                    "taker":OpDiv(MkNumber(0.5) , MkInteger(100) ),
                    "maker":OpDiv(MkNumber(0.5) , MkInteger(100) ),
                    "percentage":MkBool(true) ,
                    "tierBased":MkBool(false) ,
                    })}),
            "commonCurrencies":MkMap(&VarMap{"PLA":MkString("PlayChip") }),
            }));
}

func (this *Independentreserve) FetchMarkets(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  baseCurrencies := this.Call(MkString("publicGetGetValidPrimaryCurrencyCodes"), params ); _ = baseCurrencies;
  quoteCurrencies := this.Call(MkString("publicGetGetValidSecondaryCurrencyCodes"), params ); _ = quoteCurrencies;
  limits := this.Call(MkString("publicGetGetOrderMinimumVolumes"), params ); _ = limits;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , baseCurrencies.Length )); OpIncr(& i ){
    baseId := *(baseCurrencies ).At(i ); _ = baseId;
    base := this.SafeCurrencyCode(baseId ); _ = base;
    minAmount := this.SafeNumber(limits , baseId ); _ = minAmount;
    for j := MkInteger(0) ; IsTrue(OpLw(j , quoteCurrencies.Length )); OpIncr(& j ){
      quoteId := *(quoteCurrencies ).At(j ); _ = quoteId;
      quote := this.SafeCurrencyCode(quoteId ); _ = quote;
      id := OpAdd(baseId , OpAdd(MkString("/") , quoteId )); _ = id;
      symbol := OpAdd(base , OpAdd(MkString("/") , quote )); _ = symbol;
      result.Push(MkMap(&VarMap{
              "id":id ,
              "symbol":symbol ,
              "base":base ,
              "quote":quote ,
              "baseId":baseId ,
              "quoteId":quoteId ,
              "info":id ,
              "active":MkUndefined() ,
              "precision":*this.At(MkString("precision")) ,
              "limits":MkMap(&VarMap{"amount":MkMap(&VarMap{
                      "min":minAmount ,
                      "max":MkUndefined() ,
                      })}),
              }));
    }
  }
  return result ;
}

func (this *Independentreserve) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  balances := this.Call(MkString("privatePostGetAccounts"), params ); _ = balances;
  result := MkMap(&VarMap{"info":balances }); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , balances.Length )); OpIncr(& i ){
    balance := *(balances ).At(i ); _ = balance;
    currencyId := this.SafeString(balance , MkString("CurrencyCode") ); _ = currencyId;
    code := this.SafeCurrencyCode(currencyId ); _ = code;
    account := this.Account(); _ = account;
    *(account ).At(MkString("free") )= this.SafeString(balance , MkString("AvailableBalance") );
    *(account ).At(MkString("total") )= this.SafeString(balance , MkString("TotalBalance") );
    *(result ).At(code )= OpCopy(account );
  }
  return this.ParseBalance(result );
}

func (this *Independentreserve) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "primaryCurrencyCode":*(market ).At(MkString("baseId") ),
        "secondaryCurrencyCode":*(market ).At(MkString("quoteId") ),
        }); _ = request;
  response := this.Call(MkString("publicGetGetOrderBook"), this.Extend(request , params )); _ = response;
  timestamp := this.Parse8601(this.SafeString(response , MkString("CreatedTimestampUtc") )); _ = timestamp;
  return this.ParseOrderBook(response , symbol , timestamp , MkString("BuyOrders") , MkString("SellOrders") , MkString("Price") , MkString("Volume") );
}

func (this *Independentreserve) ParseTicker(goArgs ...*Variant) *Variant{
  ticker := GoGetArg(goArgs, 0, MkUndefined()); _ = ticker;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.Parse8601(this.SafeString(ticker , MkString("CreatedTimestampUtc") )); _ = timestamp;
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  if IsTrue(market ) {
    symbol = *(market ).At(MkString("symbol") );
  }
  last := this.SafeNumber(ticker , MkString("LastPrice") ); _ = last;
  return MkMap(&VarMap{
        "symbol":symbol ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "high":this.SafeNumber(ticker , MkString("DayHighestPrice") ),
        "low":this.SafeNumber(ticker , MkString("DayLowestPrice") ),
        "bid":this.SafeNumber(ticker , MkString("CurrentHighestBidPrice") ),
        "bidVolume":MkUndefined() ,
        "ask":this.SafeNumber(ticker , MkString("CurrentLowestOfferPrice") ),
        "askVolume":MkUndefined() ,
        "vwap":MkUndefined() ,
        "open":MkUndefined() ,
        "close":last ,
        "last":last ,
        "previousClose":MkUndefined() ,
        "change":MkUndefined() ,
        "percentage":MkUndefined() ,
        "average":this.SafeNumber(ticker , MkString("DayAvgPrice") ),
        "baseVolume":this.SafeNumber(ticker , MkString("DayVolumeXbtInSecondaryCurrrency") ),
        "quoteVolume":MkUndefined() ,
        "info":ticker ,
        });
}

func (this *Independentreserve) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "primaryCurrencyCode":*(market ).At(MkString("baseId") ),
        "secondaryCurrencyCode":*(market ).At(MkString("quoteId") ),
        }); _ = request;
  response := this.Call(MkString("publicGetGetMarketSummary"), this.Extend(request , params )); _ = response;
  return this.ParseTicker(response , market );
}

func (this *Independentreserve) ParseOrder(goArgs ...*Variant) *Variant{
  order := GoGetArg(goArgs, 0, MkUndefined()); _ = order;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  baseId := this.SafeString(order , MkString("PrimaryCurrencyCode") ); _ = baseId;
  quoteId := this.SafeString(order , MkString("SecondaryCurrencyCode") ); _ = quoteId;
  base := OpCopy(MkUndefined() ); _ = base;
  quote := OpCopy(MkUndefined() ); _ = quote;
  if IsTrue(OpAnd(OpNotEq2(baseId , MkUndefined() ), OpNotEq2(quoteId , MkUndefined() ))) {
    base = this.SafeCurrencyCode(baseId );
    quote = this.SafeCurrencyCode(quoteId );
    symbol = OpAdd(base , OpAdd(MkString("/") , quote ));
  } else {
    if IsTrue(OpNotEq2(market , MkUndefined() )) {
      symbol = *(market ).At(MkString("symbol") );
      base = *(market ).At(MkString("base") );
      quote = *(market ).At(MkString("quote") );
    }
  }
  orderType := this.SafeString2(order , MkString("Type") , MkString("OrderType") ); _ = orderType;
  side := OpCopy(MkUndefined() ); _ = side;
  if IsTrue(OpGtEq(orderType.IndexOf(MkString("Bid") ), MkInteger(0) )) {
    side = MkString("buy") ;
  } else {
    if IsTrue(OpGtEq(orderType.IndexOf(MkString("Offer") ), MkInteger(0) )) {
      side = MkString("sell") ;
    }
  }
  if IsTrue(OpGtEq(orderType.IndexOf(MkString("Market") ), MkInteger(0) )) {
    orderType = MkString("market") ;
  } else {
    if IsTrue(OpGtEq(orderType.IndexOf(MkString("Limit") ), MkInteger(0) )) {
      orderType = MkString("limit") ;
    }
  }
  timestamp := this.Parse8601(this.SafeString(order , MkString("CreatedTimestampUtc") )); _ = timestamp;
  amount := this.SafeNumber2(order , MkString("VolumeOrdered") , MkString("Volume") ); _ = amount;
  filled := this.SafeNumber(order , MkString("VolumeFilled") ); _ = filled;
  remaining := this.SafeNumber(order , MkString("Outstanding") ); _ = remaining;
  feeRate := this.SafeNumber(order , MkString("FeePercent") ); _ = feeRate;
  feeCost := OpCopy(MkUndefined() ); _ = feeCost;
  if IsTrue(OpNotEq2(feeRate , MkUndefined() )) {
    feeCost = OpMulti(feeRate , filled );
  }
  fee := MkMap(&VarMap{
        "rate":feeRate ,
        "cost":feeCost ,
        "currency":base ,
        }); _ = fee;
  id := this.SafeString(order , MkString("OrderGuid") ); _ = id;
  status := this.ParseOrderStatus(this.SafeString(order , MkString("Status") )); _ = status;
  cost := this.SafeNumber(order , MkString("Value") ); _ = cost;
  average := this.SafeNumber(order , MkString("AvgPrice") ); _ = average;
  price := this.SafeNumber(order , MkString("Price") ); _ = price;
  return this.SafeOrder(MkMap(&VarMap{
            "info":order ,
            "id":id ,
            "clientOrderId":MkUndefined() ,
            "timestamp":timestamp ,
            "datetime":this.Iso8601(timestamp ),
            "lastTradeTimestamp":MkUndefined() ,
            "symbol":symbol ,
            "type":orderType ,
            "timeInForce":MkUndefined() ,
            "postOnly":MkUndefined() ,
            "side":side ,
            "price":price ,
            "stopPrice":MkUndefined() ,
            "cost":cost ,
            "average":average ,
            "amount":amount ,
            "filled":filled ,
            "remaining":remaining ,
            "status":status ,
            "fee":fee ,
            "trades":MkUndefined() ,
            }));
}

func (this *Independentreserve) ParseOrderStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "Open":MkString("open") ,
        "PartiallyFilled":MkString("open") ,
        "Filled":MkString("closed") ,
        "PartiallyFilledAndCancelled":MkString("canceled") ,
        "Cancelled":MkString("canceled") ,
        "PartiallyFilledAndExpired":MkString("canceled") ,
        "Expired":MkString("canceled") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Independentreserve) FetchOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("privatePostGetOrderDetails"), this.Extend(MkMap(&VarMap{"orderGuid":id }), params )); _ = response;
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
  }
  return this.ParseOrder(response , market );
}

func (this *Independentreserve) FetchOpenOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := this.Ordered(MkMap(&VarMap{})); _ = request;
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("primaryCurrencyCode") )= *(market ).At(MkString("baseId") );
    *(request ).At(MkString("secondaryCurrencyCode") )= *(market ).At(MkString("quoteId") );
  }
  if IsTrue(OpEq2(limit , MkUndefined() )) {
    limit = MkInteger(50) ;
  }
  *(request ).At(MkString("pageIndex") )= MkInteger(1) ;
  *(request ).At(MkString("pageSize") )= OpCopy(limit );
  response := this.Call(MkString("privatePostGetOpenOrders"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("Data") , MkArray(&VarArray{})); _ = data;
  return this.ParseOrders(data , market , since , limit );
}

func (this *Independentreserve) FetchClosedOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := this.Ordered(MkMap(&VarMap{})); _ = request;
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("primaryCurrencyCode") )= *(market ).At(MkString("baseId") );
    *(request ).At(MkString("secondaryCurrencyCode") )= *(market ).At(MkString("quoteId") );
  }
  if IsTrue(OpEq2(limit , MkUndefined() )) {
    limit = MkInteger(50) ;
  }
  *(request ).At(MkString("pageIndex") )= MkInteger(1) ;
  *(request ).At(MkString("pageSize") )= OpCopy(limit );
  response := this.Call(MkString("privatePostGetClosedOrders"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("Data") , MkArray(&VarArray{})); _ = data;
  return this.ParseOrders(data , market , since , limit );
}

func (this *Independentreserve) FetchMyTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkInteger(50) ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  pageIndex := this.SafeInteger(params , MkString("pageIndex") , MkInteger(1) ); _ = pageIndex;
  if IsTrue(OpEq2(limit , MkUndefined() )) {
    limit = MkInteger(50) ;
  }
  request := this.Ordered(MkMap(&VarMap{
            "pageIndex":pageIndex ,
            "pageSize":limit ,
            })); _ = request;
  response := this.Call(MkString("privatePostGetTrades"), this.Extend(request , params )); _ = response;
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
  }
  return this.ParseTrades(*(response ).At(MkString("Data") ), market , since , limit );
}

func (this *Independentreserve) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.Parse8601(*(trade ).At(MkString("TradeTimestampUtc") )); _ = timestamp;
  id := this.SafeString(trade , MkString("TradeGuid") ); _ = id;
  orderId := this.SafeString(trade , MkString("OrderGuid") ); _ = orderId;
  priceString := this.SafeString2(trade , MkString("Price") , MkString("SecondaryCurrencyTradePrice") ); _ = priceString;
  amountString := this.SafeString2(trade , MkString("VolumeTraded") , MkString("PrimaryCurrencyAmount") ); _ = amountString;
  price := this.ParseNumber(priceString ); _ = price;
  amount := this.ParseNumber(amountString ); _ = amount;
  cost := this.ParseNumber(Precise.StringMul(priceString , amountString )); _ = cost;
  baseId := this.SafeString(trade , MkString("PrimaryCurrencyCode") ); _ = baseId;
  quoteId := this.SafeString(trade , MkString("SecondaryCurrencyCode") ); _ = quoteId;
  marketId := OpCopy(MkUndefined() ); _ = marketId;
  if IsTrue(OpAnd(OpNotEq2(baseId , MkUndefined() ), OpNotEq2(quoteId , MkUndefined() ))) {
    marketId = OpAdd(baseId , OpAdd(MkString("/") , quoteId ));
  }
  symbol := this.SafeSymbol(marketId , market , MkString("/") ); _ = symbol;
  side := this.SafeString(trade , MkString("OrderType") ); _ = side;
  if IsTrue(OpNotEq2(side , MkUndefined() )) {
    if IsTrue(OpGtEq(side.IndexOf(MkString("Bid") ), MkInteger(0) )) {
      side = MkString("buy") ;
    } else {
      if IsTrue(OpGtEq(side.IndexOf(MkString("Offer") ), MkInteger(0) )) {
        side = MkString("sell") ;
      }
    }
  }
  return MkMap(&VarMap{
        "id":id ,
        "info":trade ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "symbol":symbol ,
        "order":orderId ,
        "type":MkUndefined() ,
        "side":side ,
        "takerOrMaker":MkUndefined() ,
        "price":price ,
        "amount":amount ,
        "cost":cost ,
        "fee":MkUndefined() ,
        });
}

func (this *Independentreserve) FetchTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "primaryCurrencyCode":*(market ).At(MkString("baseId") ),
        "secondaryCurrencyCode":*(market ).At(MkString("quoteId") ),
        "numberOfRecentTradesToRetrieve":MkInteger(50) ,
        }); _ = request;
  response := this.Call(MkString("publicGetGetRecentTrades"), this.Extend(request , params )); _ = response;
  return this.ParseTrades(*(response ).At(MkString("Trades") ), market , since , limit );
}

func (this *Independentreserve) CreateOrder(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 2, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 3, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 4, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 5, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  capitalizedOrderType := this.Capitalize(type_ ); _ = capitalizedOrderType;
  method := OpAdd(MkString("privatePostPlace") , OpAdd(capitalizedOrderType , MkString("Order") )); _ = method;
  orderType := OpCopy(capitalizedOrderType ); _ = orderType;
  OpTrinary(OpAddAssign(& orderType , OpEq2(side , MkString("sell") )), MkString("Offer") , MkString("Bid") );
  request := this.Ordered(MkMap(&VarMap{
            "primaryCurrencyCode":*(market ).At(MkString("baseId") ),
            "secondaryCurrencyCode":*(market ).At(MkString("quoteId") ),
            "orderType":orderType ,
            })); _ = request;
  if IsTrue(OpEq2(type_ , MkString("limit") )) {
    *(request ).At(MkString("price") )= OpCopy(price );
  }
  *(request ).At(MkString("volume") )= OpCopy(amount );
  response := this.Call(method , this.Extend(request , params )); _ = response;
  return MkMap(&VarMap{
        "info":response ,
        "id":*(response ).At(MkString("OrderGuid") ),
        });
}

func (this *Independentreserve) CancelOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"orderGuid":id }); _ = request;
  return this.Call(MkString("privatePostCancelOrder"), this.Extend(request , params ));
}

func (this *Independentreserve) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("public") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  url := OpAdd(*(*(*this.At(MkString("urls")) ).At(MkString("api") )).At(api ), OpAdd(MkString("/") , path )); _ = url;
  if IsTrue(OpEq2(api , MkString("public") )) {
    if IsTrue(*(GoGetKeys(params )).At(MkString("length") )) {
      OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(params )));
    }
  } else {
    this.CheckRequiredCredentials();
    nonce := this.Nonce(); _ = nonce;
    auth := MkArray(&VarArray{
          url ,
          OpAdd(MkString("apiKey=") , *this.At(MkString("apiKey")) ),
          OpAdd(MkString("nonce=") , nonce.ToString()),
          }); _ = auth;
    keys := GoGetKeys(params ); _ = keys;
    for i := MkInteger(0) ; IsTrue(OpLw(i , keys.Length )); OpIncr(& i ){
      key := *(keys ).At(i ); _ = key;
      value := (*(params ).At(key )).Call(MkString("toString") ); _ = value;
      auth.Push(OpAdd(key , OpAdd(MkString("=") , value )));
    }
    message := auth.Join(MkString(",") ); _ = message;
    signature := this.Hmac(this.Encode(message ), this.Encode(*this.At(MkString("secret")) )); _ = signature;
    query := this.Ordered(MkMap(&VarMap{})); _ = query;
    *(query ).At(MkString("apiKey") )= OpCopy(*this.At(MkString("apiKey")) );
    *(query ).At(MkString("nonce") )= OpCopy(nonce );
    *(query ).At(MkString("signature") )= signature.ToUpperCase();
    for i := MkInteger(0) ; IsTrue(OpLw(i , keys.Length )); OpIncr(& i ){
      key := *(keys ).At(i ); _ = key;
      *(query ).At(key )= *(params ).At(key );
    }
    body = this.Json(query );
    headers = MkMap(&VarMap{"Content-Type":MkString("application/json") });
  }
  return MkMap(&VarMap{
        "url":url ,
        "method":method ,
        "body":body ,
        "headers":headers ,
        });
}

