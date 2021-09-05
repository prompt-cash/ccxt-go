package ccxt

import (
)

type Bitmex struct {
	*ExchangeBase
}

var _ Exchange = (*Bitmex)(nil)

func init() {
	exchange := &Bitmex{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Bitmex) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("bitmex") ,
            "name":MkString("BitMEX") ,
            "countries":MkArray(&VarArray{
                MkString("SC") ,
                }),
            "version":MkString("v1") ,
            "userAgent":MkUndefined() ,
            "rateLimit":MkInteger(2000) ,
            "pro":MkBool(true) ,
            "has":MkMap(&VarMap{
                "cancelAllOrders":MkBool(true) ,
                "cancelOrder":MkBool(true) ,
                "CORS":MkBool(false) ,
                "createOrder":MkBool(true) ,
                "editOrder":MkBool(true) ,
                "fetchBalance":MkBool(true) ,
                "fetchClosedOrders":MkBool(true) ,
                "fetchLedger":MkBool(true) ,
                "fetchMarkets":MkBool(true) ,
                "fetchMyTrades":MkBool(true) ,
                "fetchOHLCV":MkBool(true) ,
                "fetchOpenOrders":MkBool(true) ,
                "fetchOrder":MkBool(true) ,
                "fetchOrderBook":MkBool(true) ,
                "fetchOrders":MkBool(true) ,
                "fetchTicker":MkBool(true) ,
                "fetchTickers":MkBool(true) ,
                "fetchTrades":MkBool(true) ,
                "fetchTransactions":MkString("emulated") ,
                "withdraw":MkBool(true) ,
                }),
            "timeframes":MkMap(&VarMap{
                "1m":MkString("1m") ,
                "5m":MkString("5m") ,
                "1h":MkString("1h") ,
                "1d":MkString("1d") ,
                }),
            "urls":MkMap(&VarMap{
                "test":MkMap(&VarMap{
                    "public":MkString("https://testnet.bitmex.com") ,
                    "private":MkString("https://testnet.bitmex.com") ,
                    }),
                "logo":MkString("https://user-images.githubusercontent.com/1294454/27766319-f653c6e6-5ed4-11e7-933d-f0bc3699ae8f.jpg") ,
                "api":MkMap(&VarMap{
                    "public":MkString("https://www.bitmex.com") ,
                    "private":MkString("https://www.bitmex.com") ,
                    }),
                "www":MkString("https://www.bitmex.com") ,
                "doc":MkArray(&VarArray{
                    MkString("https://www.bitmex.com/app/apiOverview") ,
                    MkString("https://github.com/BitMEX/api-connectors/tree/master/official-http") ,
                    }),
                "fees":MkString("https://www.bitmex.com/app/fees") ,
                "referral":MkString("https://www.bitmex.com/register/upZpOX") ,
                }),
            "api":MkMap(&VarMap{
                "public":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("announcement") ,
                        MkString("announcement/urgent") ,
                        MkString("funding") ,
                        MkString("instrument") ,
                        MkString("instrument/active") ,
                        MkString("instrument/activeAndIndices") ,
                        MkString("instrument/activeIntervals") ,
                        MkString("instrument/compositeIndex") ,
                        MkString("instrument/indices") ,
                        MkString("insurance") ,
                        MkString("leaderboard") ,
                        MkString("liquidation") ,
                        MkString("orderBook") ,
                        MkString("orderBook/L2") ,
                        MkString("quote") ,
                        MkString("quote/bucketed") ,
                        MkString("schema") ,
                        MkString("schema/websocketHelp") ,
                        MkString("settlement") ,
                        MkString("stats") ,
                        MkString("stats/history") ,
                        MkString("trade") ,
                        MkString("trade/bucketed") ,
                        })}),
                "private":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("apiKey") ,
                        MkString("chat") ,
                        MkString("chat/channels") ,
                        MkString("chat/connected") ,
                        MkString("execution") ,
                        MkString("execution/tradeHistory") ,
                        MkString("notification") ,
                        MkString("order") ,
                        MkString("position") ,
                        MkString("user") ,
                        MkString("user/affiliateStatus") ,
                        MkString("user/checkReferralCode") ,
                        MkString("user/commission") ,
                        MkString("user/depositAddress") ,
                        MkString("user/executionHistory") ,
                        MkString("user/margin") ,
                        MkString("user/minWithdrawalFee") ,
                        MkString("user/wallet") ,
                        MkString("user/walletHistory") ,
                        MkString("user/walletSummary") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("apiKey") ,
                        MkString("apiKey/disable") ,
                        MkString("apiKey/enable") ,
                        MkString("chat") ,
                        MkString("order") ,
                        MkString("order/bulk") ,
                        MkString("order/cancelAllAfter") ,
                        MkString("order/closePosition") ,
                        MkString("position/isolate") ,
                        MkString("position/leverage") ,
                        MkString("position/riskLimit") ,
                        MkString("position/transferMargin") ,
                        MkString("user/cancelWithdrawal") ,
                        MkString("user/confirmEmail") ,
                        MkString("user/confirmEnableTFA") ,
                        MkString("user/confirmWithdrawal") ,
                        MkString("user/disableTFA") ,
                        MkString("user/logout") ,
                        MkString("user/logoutAll") ,
                        MkString("user/preferences") ,
                        MkString("user/requestEnableTFA") ,
                        MkString("user/requestWithdrawal") ,
                        }),
                    "put":MkArray(&VarArray{
                        MkString("order") ,
                        MkString("order/bulk") ,
                        MkString("user") ,
                        }),
                    "delete":MkArray(&VarArray{
                        MkString("apiKey") ,
                        MkString("order") ,
                        MkString("order/all") ,
                        }),
                    }),
                }),
            "exceptions":MkMap(&VarMap{
                "exact":MkMap(&VarMap{
                    "Invalid API Key.":AuthenticationError ,
                    "This key is disabled.":PermissionDenied ,
                    "Access Denied":PermissionDenied ,
                    "Duplicate clOrdID":InvalidOrder ,
                    "orderQty is invalid":InvalidOrder ,
                    "Invalid price":InvalidOrder ,
                    "Invalid stopPx for ordType":InvalidOrder ,
                    }),
                "broad":MkMap(&VarMap{
                    "Signature not valid":AuthenticationError ,
                    "overloaded":ExchangeNotAvailable ,
                    "Account has insufficient Available Balance":InsufficientFunds ,
                    "Service unavailable":ExchangeNotAvailable ,
                    "Server Error":ExchangeError ,
                    "Unable to cancel order due to existing state":InvalidOrder ,
                    }),
                }),
            "precisionMode":TICK_SIZE ,
            "options":MkMap(&VarMap{
                "api-expires":MkInteger(5) ,
                "fetchOHLCVOpenTimestamp":MkBool(true) ,
                }),
            }));
}

func (this *Bitmex) FetchMarkets(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetInstrumentActiveAndIndices"), params ); _ = response;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , response.Length )); OpIncr(& i ){
    market := *(response ).At(i ); _ = market;
    active := OpNotEq2(*(market ).At(MkString("state") ), MkString("Unlisted") ); _ = active;
    id := *(market ).At(MkString("symbol") ); _ = id;
    baseId := *(market ).At(MkString("underlying") ); _ = baseId;
    quoteId := *(market ).At(MkString("quoteCurrency") ); _ = quoteId;
    basequote := OpAdd(baseId , quoteId ); _ = basequote;
    base := this.SafeCurrencyCode(baseId ); _ = base;
    quote := this.SafeCurrencyCode(quoteId ); _ = quote;
    swap := OpEq2(id , basequote ); _ = swap;
    positionId := this.SafeString2(market , MkString("positionCurrency") , MkString("quoteCurrency") ); _ = positionId;
    type_ := OpCopy(MkUndefined() ); _ = type_;
    future := OpCopy(MkBool(false) ); _ = future;
    prediction := OpCopy(MkBool(false) ); _ = prediction;
    position := this.SafeCurrencyCode(positionId ); _ = position;
    symbol := OpCopy(id ); _ = symbol;
    if IsTrue(swap ) {
      type_ = MkString("swap") ;
      symbol = OpAdd(base , OpAdd(MkString("/") , quote ));
    } else {
      if IsTrue(OpGtEq(id.IndexOf(MkString("B_") ), MkInteger(0) )) {
        prediction = OpCopy(MkBool(true) );
        type_ = MkString("prediction") ;
      } else {
        future = OpCopy(MkBool(true) );
        type_ = MkString("future") ;
      }
    }
    precision := MkMap(&VarMap{
          "amount":MkUndefined() ,
          "price":MkUndefined() ,
          }); _ = precision;
    lotSize := this.SafeNumber(market , MkString("lotSize") ); _ = lotSize;
    tickSize := this.SafeNumber(market , MkString("tickSize") ); _ = tickSize;
    if IsTrue(OpNotEq2(lotSize , MkUndefined() )) {
      *(precision ).At(MkString("amount") )= OpCopy(lotSize );
    }
    if IsTrue(OpNotEq2(tickSize , MkUndefined() )) {
      *(precision ).At(MkString("price") )= OpCopy(tickSize );
    }
    limits := MkMap(&VarMap{
          "amount":MkMap(&VarMap{
              "min":MkUndefined() ,
              "max":MkUndefined() ,
              }),
          "price":MkMap(&VarMap{
              "min":tickSize ,
              "max":this.SafeNumber(market , MkString("maxPrice") ),
              }),
          "cost":MkMap(&VarMap{
              "min":MkUndefined() ,
              "max":MkUndefined() ,
              }),
          }); _ = limits;
    limitField := OpTrinary(OpEq2(position , quote ), MkString("cost") , MkString("amount") ); _ = limitField;
    *(limits ).At(limitField )= MkMap(&VarMap{
        "min":lotSize ,
        "max":this.SafeNumber(market , MkString("maxOrderQty") ),
        });
    result.Push(MkMap(&VarMap{
            "id":id ,
            "symbol":symbol ,
            "base":base ,
            "quote":quote ,
            "baseId":baseId ,
            "quoteId":quoteId ,
            "active":active ,
            "precision":precision ,
            "limits":limits ,
            "taker":this.SafeNumber(market , MkString("takerFee") ),
            "maker":this.SafeNumber(market , MkString("makerFee") ),
            "type":type_ ,
            "spot":MkBool(false) ,
            "swap":swap ,
            "future":future ,
            "prediction":prediction ,
            "info":market ,
            }));
  }
  return result ;
}

func (this *Bitmex) ParseBalanceResponse(goArgs ...*Variant) *Variant{
  response := GoGetArg(goArgs, 0, MkUndefined()); _ = response;
  result := MkMap(&VarMap{"info":response }); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , response.Length )); OpIncr(& i ){
    balance := *(response ).At(i ); _ = balance;
    currencyId := this.SafeString(balance , MkString("currency") ); _ = currencyId;
    code := this.SafeCurrencyCode(currencyId ); _ = code;
    account := this.Account(); _ = account;
    free := this.SafeString(balance , MkString("availableMargin") ); _ = free;
    total := this.SafeString(balance , MkString("marginBalance") ); _ = total;
    if IsTrue(OpEq2(code , MkString("BTC") )) {
      free = Precise.StringDiv(free , MkString("1e8") );
      total = Precise.StringDiv(total , MkString("1e8") );
    }
    *(account ).At(MkString("free") )= OpCopy(free );
    *(account ).At(MkString("total") )= OpCopy(total );
    *(result ).At(code )= OpCopy(account );
  }
  return this.ParseBalance(result );
}

func (this *Bitmex) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"currency":MkString("all") }); _ = request;
  response := this.Call(MkString("privateGetUserMargin"), this.Extend(request , params )); _ = response;
  return this.ParseBalanceResponse(response );
}

func (this *Bitmex) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"symbol":*(market ).At(MkString("id") )}); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("depth") )= OpCopy(limit );
  }
  response := this.Call(MkString("publicGetOrderBookL2"), this.Extend(request , params )); _ = response;
  result := MkMap(&VarMap{
        "symbol":symbol ,
        "bids":MkArray(&VarArray{}),
        "asks":MkArray(&VarArray{}),
        "timestamp":MkUndefined() ,
        "datetime":MkUndefined() ,
        "nonce":MkUndefined() ,
        }); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , response.Length )); OpIncr(& i ){
    order := *(response ).At(i ); _ = order;
    side := OpTrinary(OpEq2(*(order ).At(MkString("side") ), MkString("Sell") ), MkString("asks") , MkString("bids") ); _ = side;
    amount := this.SafeNumber(order , MkString("size") ); _ = amount;
    price := this.SafeNumber(order , MkString("price") ); _ = price;
    if IsTrue(OpNotEq2(price , MkUndefined() )) {
      (*(result ).At(side )).Call(MkString("push") , MkArray(&VarArray{
              price ,
              amount ,
              }));
    }
  }
  *(result ).At(MkString("bids") )= this.SortBy(*(result ).At(MkString("bids") ), MkInteger(0) , MkBool(true) );
  *(result ).At(MkString("asks") )= this.SortBy(*(result ).At(MkString("asks") ), MkInteger(0) );
  return result ;
}

func (this *Bitmex) FetchOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  filter := MkMap(&VarMap{"filter":MkMap(&VarMap{"orderID":id })}); _ = filter;
  response := this.FetchOrders(symbol , MkUndefined() , MkUndefined() , this.DeepExtend(filter , params )); _ = response;
  numResults := OpCopy(response.Length ); _ = numResults;
  if IsTrue(OpEq2(numResults , MkInteger(1) )) {
    return *(response ).At(MkInteger(0) );
  }
  panic(NewOrderNotFound(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(": The order ") , OpAdd(id , MkString(" not found.") )))));
  return MkUndefined()
}

func (this *Bitmex) FetchOrders(goArgs ...*Variant) *Variant{
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
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("startTime") )= this.Iso8601(since );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("count") )= OpCopy(limit );
  }
  request = this.DeepExtend(request , params );
  if IsTrue(OpHasMember(MkString("filter") , request )) {
    *(request ).At(MkString("filter") )= this.Json(*(request ).At(MkString("filter") ));
  }
  response := this.Call(MkString("privateGetOrder"), request ); _ = response;
  return this.ParseOrders(response , market , since , limit );
}

func (this *Bitmex) FetchOpenOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"filter":MkMap(&VarMap{"open":MkBool(true) })}); _ = request;
  return this.FetchOrders(symbol , since , limit , this.DeepExtend(request , params ));
}

func (this *Bitmex) FetchClosedOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  orders := this.FetchOrders(symbol , since , limit , params ); _ = orders;
  return this.FilterBy(orders , MkString("status") , MkString("closed") );
}

func (this *Bitmex) FetchMyTrades(goArgs ...*Variant) *Variant{
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
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("startTime") )= this.Iso8601(since );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("count") )= OpCopy(limit );
  }
  request = this.DeepExtend(request , params );
  if IsTrue(OpHasMember(MkString("filter") , request )) {
    *(request ).At(MkString("filter") )= this.Json(*(request ).At(MkString("filter") ));
  }
  response := this.Call(MkString("privateGetExecutionTradeHistory"), request ); _ = response;
  return this.ParseTrades(response , market , since , limit );
}

func (this *Bitmex) ParseLedgerEntryType(goArgs ...*Variant) *Variant{
  type_ := GoGetArg(goArgs, 0, MkUndefined()); _ = type_;
  types := MkMap(&VarMap{
        "Withdrawal":MkString("transaction") ,
        "RealisedPNL":MkString("margin") ,
        "UnrealisedPNL":MkString("margin") ,
        "Deposit":MkString("transaction") ,
        "Transfer":MkString("transfer") ,
        "AffiliatePayout":MkString("referral") ,
        }); _ = types;
  return this.SafeString(types , type_ , type_ );
}

func (this *Bitmex) ParseLedgerEntry(goArgs ...*Variant) *Variant{
  item := GoGetArg(goArgs, 0, MkUndefined()); _ = item;
  currency := GoGetArg(goArgs, 1, MkUndefined() ); _ = currency;
  id := this.SafeString(item , MkString("transactID") ); _ = id;
  account := this.SafeString(item , MkString("account") ); _ = account;
  referenceId := this.SafeString(item , MkString("tx") ); _ = referenceId;
  referenceAccount := OpCopy(MkUndefined() ); _ = referenceAccount;
  type_ := this.ParseLedgerEntryType(this.SafeString(item , MkString("transactType") )); _ = type_;
  currencyId := this.SafeString(item , MkString("currency") ); _ = currencyId;
  code := this.SafeCurrencyCode(currencyId , currency ); _ = code;
  amount := this.SafeNumber(item , MkString("amount") ); _ = amount;
  if IsTrue(OpNotEq2(amount , MkUndefined() )) {
    amount = OpDiv(amount , MkInteger(100000000) );
  }
  timestamp := this.Parse8601(this.SafeString(item , MkString("transactTime") )); _ = timestamp;
  if IsTrue(OpEq2(timestamp , MkUndefined() )) {
    timestamp = MkInteger(0) ;
  }
  feeCost := this.SafeNumber(item , MkString("fee") , MkInteger(0) ); _ = feeCost;
  if IsTrue(OpNotEq2(feeCost , MkUndefined() )) {
    feeCost = OpDiv(feeCost , MkInteger(100000000) );
  }
  fee := MkMap(&VarMap{
        "cost":feeCost ,
        "currency":code ,
        }); _ = fee;
  after := this.SafeNumber(item , MkString("walletBalance") ); _ = after;
  if IsTrue(OpNotEq2(after , MkUndefined() )) {
    after = OpDiv(after , MkInteger(100000000) );
  }
  before := this.Sum(after , OpNeg(amount )); _ = before;
  direction := OpCopy(MkUndefined() ); _ = direction;
  if IsTrue(OpLw(amount , MkInteger(0) )) {
    direction = MkString("out") ;
    amount = Math.Abs(amount );
  } else {
    direction = MkString("in") ;
  }
  status := this.ParseTransactionStatus(this.SafeString(item , MkString("transactStatus") )); _ = status;
  return MkMap(&VarMap{
        "id":id ,
        "info":item ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
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
        "fee":fee ,
        });
}

func (this *Bitmex) FetchLedger(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  currency := OpCopy(MkUndefined() ); _ = currency;
  if IsTrue(OpNotEq2(code , MkUndefined() )) {
    currency = this.Currency(code );
  }
  request := MkMap(&VarMap{}); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("count") )= OpCopy(limit );
  }
  response := this.Call(MkString("privateGetUserWalletHistory"), this.Extend(request , params )); _ = response;
  return this.ParseLedger(response , currency , since , limit );
}

func (this *Bitmex) FetchTransactions(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("count") )= OpCopy(limit );
  }
  response := this.Call(MkString("privateGetUserWalletHistory"), this.Extend(request , params )); _ = response;
  transactions := this.FilterByArray(response , MkString("transactType") , MkArray(&VarArray{
            MkString("Withdrawal") ,
            MkString("Deposit") ,
            }), MkBool(false) ); _ = transactions;
  currency := OpCopy(MkUndefined() ); _ = currency;
  if IsTrue(OpNotEq2(code , MkUndefined() )) {
    currency = this.Currency(code );
  }
  return this.ParseTransactions(transactions , currency , since , limit );
}

func (this *Bitmex) ParseTransactionStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "Canceled":MkString("canceled") ,
        "Completed":MkString("ok") ,
        "Pending":MkString("pending") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Bitmex) ParseTransaction(goArgs ...*Variant) *Variant{
  transaction := GoGetArg(goArgs, 0, MkUndefined()); _ = transaction;
  currency := GoGetArg(goArgs, 1, MkUndefined() ); _ = currency;
  id := this.SafeString(transaction , MkString("transactID") ); _ = id;
  transactTime := this.Parse8601(this.SafeString(transaction , MkString("transactTime") )); _ = transactTime;
  timestamp := this.Parse8601(this.SafeString(transaction , MkString("timestamp") )); _ = timestamp;
  type_ := this.SafeStringLower(transaction , MkString("transactType") ); _ = type_;
  address := OpCopy(MkUndefined() ); _ = address;
  addressFrom := OpCopy(MkUndefined() ); _ = addressFrom;
  addressTo := OpCopy(MkUndefined() ); _ = addressTo;
  if IsTrue(OpEq2(type_ , MkString("withdrawal") )) {
    address = this.SafeString(transaction , MkString("address") );
    addressFrom = this.SafeString(transaction , MkString("tx") );
    addressTo = OpCopy(address );
  }
  amountString := this.SafeString(transaction , MkString("amount") ); _ = amountString;
  amountString = Precise.StringDiv(Precise.StringAbs(amountString ), MkString("1e8") );
  feeCostString := this.SafeString(transaction , MkString("fee") ); _ = feeCostString;
  feeCostString = Precise.StringDiv(feeCostString , MkString("1e8") );
  fee := MkMap(&VarMap{
        "cost":this.ParseNumber(feeCostString ),
        "currency":MkString("BTC") ,
        }); _ = fee;
  status := this.SafeString(transaction , MkString("transactStatus") ); _ = status;
  if IsTrue(OpNotEq2(status , MkUndefined() )) {
    status = this.ParseTransactionStatus(status );
  }
  return MkMap(&VarMap{
        "info":transaction ,
        "id":id ,
        "txid":MkUndefined() ,
        "timestamp":transactTime ,
        "datetime":this.Iso8601(transactTime ),
        "addressFrom":addressFrom ,
        "address":address ,
        "addressTo":addressTo ,
        "tagFrom":MkUndefined() ,
        "tag":MkUndefined() ,
        "tagTo":MkUndefined() ,
        "type":type_ ,
        "amount":this.ParseNumber(amountString ),
        "currency":MkString("BTC") ,
        "status":status ,
        "updated":timestamp ,
        "comment":MkUndefined() ,
        "fee":fee ,
        });
}

func (this *Bitmex) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  if IsTrue(OpNot(*(market ).At(MkString("active") ))) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(": symbol ") , OpAdd(symbol , MkString(" is delisted") )))));
  }
  tickers := this.FetchTickers(MkArray(&VarArray{
            symbol ,
            }), params ); _ = tickers;
  ticker := this.SafeValue(tickers , symbol ); _ = ticker;
  if IsTrue(OpEq2(ticker , MkUndefined() )) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ticker symbol ") , OpAdd(symbol , MkString(" not found") )))));
  }
  return ticker ;
}

func (this *Bitmex) FetchTickers(goArgs ...*Variant) *Variant{
  symbols := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbols;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("publicGetInstrumentActiveAndIndices"), params ); _ = response;
  result := MkMap(&VarMap{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , response.Length )); OpIncr(& i ){
    ticker := this.ParseTicker(*(response ).At(i )); _ = ticker;
    symbol := this.SafeString(ticker , MkString("symbol") ); _ = symbol;
    if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
      *(result ).At(symbol )= OpCopy(ticker );
    }
  }
  return this.FilterByArray(result , MkString("symbol") , symbols );
}

func (this *Bitmex) ParseTicker(goArgs ...*Variant) *Variant{
  ticker := GoGetArg(goArgs, 0, MkUndefined()); _ = ticker;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  marketId := this.SafeString(ticker , MkString("symbol") ); _ = marketId;
  market = this.SafeValue(*this.At(MkString("markets_by_id")) , marketId , market );
  if IsTrue(OpNotEq2(market , MkUndefined() )) {
    symbol = *(market ).At(MkString("symbol") );
  }
  timestamp := this.Parse8601(this.SafeString(ticker , MkString("timestamp") )); _ = timestamp;
  open := this.SafeNumber(ticker , MkString("prevPrice24h") ); _ = open;
  last := this.SafeNumber(ticker , MkString("lastPrice") ); _ = last;
  change := OpCopy(MkUndefined() ); _ = change;
  percentage := OpCopy(MkUndefined() ); _ = percentage;
  if IsTrue(OpAnd(OpNotEq2(last , MkUndefined() ), OpNotEq2(open , MkUndefined() ))) {
    change = OpSub(last , open );
    if IsTrue(OpGt(open , MkInteger(0) )) {
      percentage = OpDiv(change , OpMulti(open , MkInteger(100) ));
    }
  }
  return MkMap(&VarMap{
        "symbol":symbol ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "high":this.SafeNumber(ticker , MkString("highPrice") ),
        "low":this.SafeNumber(ticker , MkString("lowPrice") ),
        "bid":this.SafeNumber(ticker , MkString("bidPrice") ),
        "bidVolume":MkUndefined() ,
        "ask":this.SafeNumber(ticker , MkString("askPrice") ),
        "askVolume":MkUndefined() ,
        "vwap":this.SafeNumber(ticker , MkString("vwap") ),
        "open":open ,
        "close":last ,
        "last":last ,
        "previousClose":MkUndefined() ,
        "change":change ,
        "percentage":percentage ,
        "average":OpDiv(this.Sum(open , last ), MkInteger(2) ),
        "baseVolume":this.SafeNumber(ticker , MkString("homeNotional24h") ),
        "quoteVolume":this.SafeNumber(ticker , MkString("foreignNotional24h") ),
        "info":ticker ,
        });
}

func (this *Bitmex) ParseOHLCV(goArgs ...*Variant) *Variant{
  ohlcv := GoGetArg(goArgs, 0, MkUndefined()); _ = ohlcv;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  return MkArray(&VarArray{
        this.Parse8601(this.SafeString(ohlcv , MkString("timestamp") )),
        this.SafeNumber(ohlcv , MkString("open") ),
        this.SafeNumber(ohlcv , MkString("high") ),
        this.SafeNumber(ohlcv , MkString("low") ),
        this.SafeNumber(ohlcv , MkString("close") ),
        this.SafeNumber(ohlcv , MkString("volume") ),
        });
}

func (this *Bitmex) FetchOHLCV(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  timeframe := GoGetArg(goArgs, 1, MkString("1m") ); _ = timeframe;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "symbol":*(market ).At(MkString("id") ),
        "binSize":*(*this.At(MkString("timeframes")) ).At(timeframe ),
        "partial":MkBool(true) ,
        }); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("count") )= OpCopy(limit );
  }
  duration := OpMulti(this.ParseTimeframe(timeframe ), MkInteger(1000) ); _ = duration;
  fetchOHLCVOpenTimestamp := this.SafeValue(*this.At(MkString("options")) , MkString("fetchOHLCVOpenTimestamp") , MkBool(true) ); _ = fetchOHLCVOpenTimestamp;
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    timestamp := OpCopy(since ); _ = timestamp;
    if IsTrue(fetchOHLCVOpenTimestamp ) {
      timestamp = this.Sum(timestamp , duration );
    }
    ymdhms := this.Ymdhms(timestamp ); _ = ymdhms;
    *(request ).At(MkString("startTime") )= OpCopy(ymdhms );
  } else {
    *(request ).At(MkString("reverse") )= OpCopy(MkBool(true) );
  }
  response := this.Call(MkString("publicGetTradeBucketed"), this.Extend(request , params )); _ = response;
  result := this.ParseOHLCVs(response , market , timeframe , since , limit ); _ = result;
  if IsTrue(fetchOHLCVOpenTimestamp ) {
    for i := MkInteger(0) ; IsTrue(OpLw(i , result.Length )); OpIncr(& i ){
      *(*(result ).At(i )).At(MkInteger(0) )= OpSub(*(*(result ).At(i )).At(MkInteger(0) ), duration );
    }
  }
  return result ;
}

func (this *Bitmex) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.Parse8601(this.SafeString(trade , MkString("timestamp") )); _ = timestamp;
  priceString := this.SafeString2(trade , MkString("avgPx") , MkString("price") ); _ = priceString;
  amountString := this.SafeString2(trade , MkString("size") , MkString("lastQty") ); _ = amountString;
  id := this.SafeString(trade , MkString("trdMatchID") ); _ = id;
  order := this.SafeString(trade , MkString("orderID") ); _ = order;
  side := this.SafeStringLower(trade , MkString("side") ); _ = side;
  costString := this.SafeString(trade , MkString("execCost") ); _ = costString;
  costString = Precise.StringDiv(Precise.StringAbs(costString ), MkString("1e8") );
  fee := OpCopy(MkUndefined() ); _ = fee;
  feeCostString := Precise.StringDiv(this.SafeString(trade , MkString("execComm") ), MkString("1e8") ); _ = feeCostString;
  if IsTrue(OpNotEq2(feeCostString , MkUndefined() )) {
    currencyId := this.SafeString(trade , MkString("settlCurrency") ); _ = currencyId;
    feeCurrencyCode := this.SafeCurrencyCode(currencyId ); _ = feeCurrencyCode;
    feeRateString := this.SafeString(trade , MkString("commission") ); _ = feeRateString;
    fee = MkMap(&VarMap{
        "cost":this.ParseNumber(feeCostString ),
        "currency":feeCurrencyCode ,
        "rate":this.ParseNumber(feeRateString ),
        });
  }
  execType := this.SafeString(trade , MkString("execType") ); _ = execType;
  takerOrMaker := OpCopy(MkUndefined() ); _ = takerOrMaker;
  if IsTrue(OpAnd(OpNotEq2(feeCostString , MkUndefined() ), OpEq2(execType , MkString("Trade") ))) {
    takerOrMaker = OpTrinary(Precise.StringLt(feeCostString , MkString("0") ), MkString("maker") , MkString("taker") );
  }
  marketId := this.SafeString(trade , MkString("symbol") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market ); _ = symbol;
  type_ := this.SafeStringLower(trade , MkString("ordType") ); _ = type_;
  return MkMap(&VarMap{
        "info":trade ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "symbol":symbol ,
        "id":id ,
        "order":order ,
        "type":type_ ,
        "takerOrMaker":takerOrMaker ,
        "side":side ,
        "price":this.ParseNumber(priceString ),
        "cost":this.ParseNumber(costString ),
        "amount":this.ParseNumber(amountString ),
        "fee":fee ,
        });
}

func (this *Bitmex) ParseOrderStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "New":MkString("open") ,
        "PartiallyFilled":MkString("open") ,
        "Filled":MkString("closed") ,
        "DoneForDay":MkString("open") ,
        "Canceled":MkString("canceled") ,
        "PendingCancel":MkString("open") ,
        "PendingNew":MkString("open") ,
        "Rejected":MkString("rejected") ,
        "Expired":MkString("expired") ,
        "Stopped":MkString("open") ,
        "Untriggered":MkString("open") ,
        "Triggered":MkString("open") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Bitmex) ParseTimeInForce(goArgs ...*Variant) *Variant{
  timeInForce := GoGetArg(goArgs, 0, MkUndefined()); _ = timeInForce;
  timeInForces := MkMap(&VarMap{
        "Day":MkString("Day") ,
        "GoodTillCancel":MkString("GTC") ,
        "ImmediateOrCancel":MkString("IOC") ,
        "FillOrKill":MkString("FOK") ,
        }); _ = timeInForces;
  return this.SafeString(timeInForces , timeInForce , timeInForce );
}

func (this *Bitmex) ParseOrder(goArgs ...*Variant) *Variant{
  order := GoGetArg(goArgs, 0, MkUndefined()); _ = order;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  status := this.ParseOrderStatus(this.SafeString(order , MkString("ordStatus") )); _ = status;
  marketId := this.SafeString(order , MkString("symbol") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market ); _ = symbol;
  timestamp := this.Parse8601(this.SafeString(order , MkString("timestamp") )); _ = timestamp;
  lastTradeTimestamp := this.Parse8601(this.SafeString(order , MkString("transactTime") )); _ = lastTradeTimestamp;
  price := this.SafeNumber(order , MkString("price") ); _ = price;
  amount := this.SafeNumber(order , MkString("orderQty") ); _ = amount;
  filled := this.SafeNumber(order , MkString("cumQty") , MkNumber(0.0) ); _ = filled;
  average := this.SafeNumber(order , MkString("avgPx") ); _ = average;
  id := this.SafeString(order , MkString("orderID") ); _ = id;
  type_ := this.SafeStringLower(order , MkString("ordType") ); _ = type_;
  side := this.SafeStringLower(order , MkString("side") ); _ = side;
  clientOrderId := this.SafeString(order , MkString("clOrdID") ); _ = clientOrderId;
  timeInForce := this.ParseTimeInForce(this.SafeString(order , MkString("timeInForce") )); _ = timeInForce;
  stopPrice := this.SafeNumber(order , MkString("stopPx") ); _ = stopPrice;
  execInst := this.SafeString(order , MkString("execInst") ); _ = execInst;
  postOnly := OpEq2(execInst , MkString("ParticipateDoNotInitiate") ); _ = postOnly;
  return this.SafeOrder(MkMap(&VarMap{
            "info":order ,
            "id":id ,
            "clientOrderId":clientOrderId ,
            "timestamp":timestamp ,
            "datetime":this.Iso8601(timestamp ),
            "lastTradeTimestamp":lastTradeTimestamp ,
            "symbol":symbol ,
            "type":type_ ,
            "timeInForce":timeInForce ,
            "postOnly":postOnly ,
            "side":side ,
            "price":price ,
            "stopPrice":stopPrice ,
            "amount":amount ,
            "cost":MkUndefined() ,
            "average":average ,
            "filled":filled ,
            "remaining":MkUndefined() ,
            "status":status ,
            "fee":MkUndefined() ,
            "trades":MkUndefined() ,
            }));
}

func (this *Bitmex) FetchTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"symbol":*(market ).At(MkString("id") )}); _ = request;
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("startTime") )= this.Iso8601(since );
  } else {
    *(request ).At(MkString("reverse") )= OpCopy(MkBool(true) );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("count") )= OpCopy(limit );
  }
  response := this.Call(MkString("publicGetTrade"), this.Extend(request , params )); _ = response;
  return this.ParseTrades(response , market , since , limit );
}

func (this *Bitmex) CreateOrder(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 2, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 3, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 4, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 5, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  orderType := this.Capitalize(type_ ); _ = orderType;
  request := MkMap(&VarMap{
        "symbol":*(market ).At(MkString("id") ),
        "side":this.Capitalize(side ),
        "orderQty":parseFloat(this.AmountToPrecision(symbol , amount )),
        "ordType":orderType ,
        }); _ = request;
  if IsTrue(OpOr(OpEq2(orderType , MkString("Stop") ), OpOr(OpEq2(orderType , MkString("StopLimit") ), OpOr(OpEq2(orderType , MkString("MarketIfTouched") ), OpEq2(orderType , MkString("LimitIfTouched") ))))) {
    stopPrice := this.SafeNumber2(params , MkString("stopPx") , MkString("stopPrice") ); _ = stopPrice;
    if IsTrue(OpEq2(stopPrice , MkUndefined() )) {
      panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" createOrder() requires a stopPx or stopPrice parameter for the ") , OpAdd(orderType , MkString(" order type") )))));
    } else {
      *(request ).At(MkString("stopPx") )= parseFloat(this.PriceToPrecision(symbol , stopPrice ));
      params = this.Omit(params , MkArray(&VarArray{
              MkString("stopPx") ,
              MkString("stopPrice") ,
              }));
    }
  }
  if IsTrue(OpOr(OpEq2(orderType , MkString("Limit") ), OpOr(OpEq2(orderType , MkString("StopLimit") ), OpEq2(orderType , MkString("LimitIfTouched") )))) {
    *(request ).At(MkString("price") )= parseFloat(this.PriceToPrecision(symbol , price ));
  }
  clientOrderId := this.SafeString2(params , MkString("clOrdID") , MkString("clientOrderId") ); _ = clientOrderId;
  if IsTrue(OpNotEq2(clientOrderId , MkUndefined() )) {
    *(request ).At(MkString("clOrdID") )= OpCopy(clientOrderId );
    params = this.Omit(params , MkArray(&VarArray{
            MkString("clOrdID") ,
            MkString("clientOrderId") ,
            }));
  }
  response := this.Call(MkString("privatePostOrder"), this.Extend(request , params )); _ = response;
  return this.ParseOrder(response , market );
}

func (this *Bitmex) EditOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 2, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 3, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 4, MkUndefined() ); _ = amount;
  price := GoGetArg(goArgs, 5, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 6, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  origClOrdID := this.SafeString2(params , MkString("origClOrdID") , MkString("clientOrderId") ); _ = origClOrdID;
  if IsTrue(OpNotEq2(origClOrdID , MkUndefined() )) {
    *(request ).At(MkString("origClOrdID") )= OpCopy(origClOrdID );
    clientOrderId := this.SafeString(params , MkString("clOrdID") , MkString("clientOrderId") ); _ = clientOrderId;
    if IsTrue(OpNotEq2(clientOrderId , MkUndefined() )) {
      *(request ).At(MkString("clOrdID") )= OpCopy(clientOrderId );
    }
    params = this.Omit(params , MkArray(&VarArray{
            MkString("origClOrdID") ,
            MkString("clOrdID") ,
            MkString("clientOrderId") ,
            }));
  } else {
    *(request ).At(MkString("orderID") )= OpCopy(id );
  }
  if IsTrue(OpNotEq2(amount , MkUndefined() )) {
    *(request ).At(MkString("orderQty") )= OpCopy(amount );
  }
  if IsTrue(OpNotEq2(price , MkUndefined() )) {
    *(request ).At(MkString("price") )= OpCopy(price );
  }
  response := this.Call(MkString("privatePutOrder"), this.Extend(request , params )); _ = response;
  return this.ParseOrder(response );
}

func (this *Bitmex) CancelOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  clientOrderId := this.SafeString2(params , MkString("clOrdID") , MkString("clientOrderId") ); _ = clientOrderId;
  request := MkMap(&VarMap{}); _ = request;
  if IsTrue(OpEq2(clientOrderId , MkUndefined() )) {
    *(request ).At(MkString("orderID") )= OpCopy(id );
  } else {
    *(request ).At(MkString("clOrdID") )= OpCopy(clientOrderId );
    params = this.Omit(params , MkArray(&VarArray{
            MkString("clOrdID") ,
            MkString("clientOrderId") ,
            }));
  }
  response := this.Call(MkString("privateDeleteOrder"), this.Extend(request , params )); _ = response;
  order := this.SafeValue(response , MkInteger(0) , MkMap(&VarMap{})); _ = order;
  error := this.SafeString(order , MkString("error") ); _ = error;
  if IsTrue(OpNotEq2(error , MkUndefined() )) {
    if IsTrue(OpGtEq(error.IndexOf(MkString("Unable to cancel order due to existing state") ), MkInteger(0) )) {
      panic(NewOrderNotFound(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" cancelOrder() failed: ") , error ))));
    }
  }
  return this.ParseOrder(order );
}

func (this *Bitmex) CancelAllOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("symbol") )= *(market ).At(MkString("id") );
  }
  response := this.Call(MkString("privateDeleteOrderAll"), this.Extend(request , params )); _ = response;
  return this.ParseOrders(response , market );
}

func (this *Bitmex) FetchPositions(goArgs ...*Variant) *Variant{
  symbols := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbols;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("privateGetPosition"), params ); _ = response;
  return response ;
}

func (this *Bitmex) IsFiat(goArgs ...*Variant) *Variant{
  currency := GoGetArg(goArgs, 0, MkUndefined()); _ = currency;
  if IsTrue(OpEq2(currency , MkString("EUR") )) {
    return MkBool(true) ;
  }
  if IsTrue(OpEq2(currency , MkString("PLN") )) {
    return MkBool(true) ;
  }
  return MkBool(false) ;
}

func (this *Bitmex) Withdraw(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  amount := GoGetArg(goArgs, 1, MkUndefined()); _ = amount;
  address := GoGetArg(goArgs, 2, MkUndefined()); _ = address;
  tag := GoGetArg(goArgs, 3, MkUndefined() ); _ = tag;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.CheckAddress(address );
  this.LoadMarkets();
  if IsTrue(OpNotEq2(code , MkString("BTC") )) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , MkString(" supoprts BTC withdrawals only, other currencies coming soon...") )));
  }
  request := MkMap(&VarMap{
        "currency":MkString("XBt") ,
        "amount":amount ,
        "address":address ,
        }); _ = request;
  response := this.Call(MkString("privatePostUserRequestWithdrawal"), this.Extend(request , params )); _ = response;
  return MkMap(&VarMap{
        "info":response ,
        "id":this.SafeString(response , MkString("transactID") ),
        });
}

func (this *Bitmex) HandleErrors(goArgs ...*Variant) *Variant{
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
  if IsTrue(OpEq2(code , MkInteger(429) )) {
    panic(NewDDoSProtection(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , body ))));
  }
  if IsTrue(OpGtEq(code , MkInteger(400) )) {
    error := this.SafeValue(response , MkString("error") , MkMap(&VarMap{})); _ = error;
    message := this.SafeString(error , MkString("message") ); _ = message;
    feedback := OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , body )); _ = feedback;
    this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("exact") ), message , feedback );
    this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("broad") ), message , feedback );
    if IsTrue(OpEq2(code , MkInteger(400) )) {
      panic(NewBadRequest(feedback ));
    }
    panic(NewExchangeError(feedback ));
  }
  return MkUndefined()
}

func (this *Bitmex) Nonce(goArgs ...*Variant) *Variant{
  return this.Milliseconds();
}

func (this *Bitmex) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("public") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  query := OpAdd(MkString("/api/") , OpAdd(*this.At(MkString("version")) , OpAdd(MkString("/") , path ))); _ = query;
  if IsTrue(OpEq2(method , MkString("GET") )) {
    if IsTrue(*(GoGetKeys(params )).At(MkString("length") )) {
      OpAddAssign(& query , OpAdd(MkString("?") , this.Urlencode(params )));
    }
  } else {
    format := this.SafeString(params , MkString("_format") ); _ = format;
    if IsTrue(OpNotEq2(format , MkUndefined() )) {
      OpAddAssign(& query , OpAdd(MkString("?") , this.Urlencode(MkMap(&VarMap{"_format":format }))));
      params = this.Omit(params , MkString("_format") );
    }
  }
  url := OpAdd(*(*(*this.At(MkString("urls")) ).At(MkString("api") )).At(api ), query ); _ = url;
  if IsTrue(OpAnd(*this.At(MkString("apiKey")) , *this.At(MkString("secret")) )) {
    auth := OpAdd(method , query ); _ = auth;
    expires := this.SafeInteger(*this.At(MkString("options")) , MkString("api-expires") ); _ = expires;
    headers = MkMap(&VarMap{
        "Content-Type":MkString("application/json") ,
        "api-key":*this.At(MkString("apiKey")) ,
        });
    expires = this.Sum(this.Seconds(), expires );
    expires = expires.ToString();
    OpAddAssign(& auth , expires );
    *(headers ).At(MkString("api-expires") )= OpCopy(expires );
    if IsTrue(OpOr(OpEq2(method , MkString("POST") ), OpOr(OpEq2(method , MkString("PUT") ), OpEq2(method , MkString("DELETE") )))) {
      if IsTrue(*(GoGetKeys(params )).At(MkString("length") )) {
        body = this.Json(params );
        OpAddAssign(& auth , body );
      }
    }
    *(headers ).At(MkString("api-signature") )= this.Hmac(this.Encode(auth ), this.Encode(*this.At(MkString("secret")) ));
  }
  return MkMap(&VarMap{
        "url":url ,
        "method":method ,
        "body":body ,
        "headers":headers ,
        });
}

