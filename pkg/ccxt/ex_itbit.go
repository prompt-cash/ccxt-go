package ccxt

import (
)

type Itbit struct {
	*ExchangeBase
}

var _ Exchange = (*Itbit)(nil)

func init() {
	exchange := &Itbit{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Itbit) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("itbit") ,
            "name":MkString("itBit") ,
            "countries":MkArray(&VarArray{
                MkString("US") ,
                }),
            "rateLimit":MkInteger(2000) ,
            "version":MkString("v1") ,
            "has":MkMap(&VarMap{
                "cancelOrder":MkBool(true) ,
                "CORS":MkBool(true) ,
                "createMarketOrder":MkBool(false) ,
                "createOrder":MkBool(true) ,
                "fetchBalance":MkBool(true) ,
                "fetchClosedOrders":MkBool(true) ,
                "fetchMyTrades":MkBool(true) ,
                "fetchOpenOrders":MkBool(true) ,
                "fetchOrder":MkBool(true) ,
                "fetchOrderBook":MkBool(true) ,
                "fetchOrders":MkBool(true) ,
                "fetchTicker":MkBool(true) ,
                "fetchTrades":MkBool(true) ,
                "fetchTransactions":MkBool(true) ,
                }),
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/1294454/27822159-66153620-60ad-11e7-89e7-005f6d7f3de0.jpg") ,
                "api":MkString("https://api.itbit.com") ,
                "www":MkString("https://www.itbit.com") ,
                "doc":MkArray(&VarArray{
                    MkString("https://api.itbit.com/docs") ,
                    MkString("https://www.itbit.com/api") ,
                    }),
                }),
            "api":MkMap(&VarMap{
                "public":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("markets/{symbol}/ticker") ,
                        MkString("markets/{symbol}/order_book") ,
                        MkString("markets/{symbol}/trades") ,
                        })}),
                "private":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("wallets") ,
                        MkString("wallets/{walletId}") ,
                        MkString("wallets/{walletId}/balances/{currencyCode}") ,
                        MkString("wallets/{walletId}/funding_history") ,
                        MkString("wallets/{walletId}/trades") ,
                        MkString("wallets/{walletId}/orders") ,
                        MkString("wallets/{walletId}/orders/{id}") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("wallet_transfers") ,
                        MkString("wallets") ,
                        MkString("wallets/{walletId}/cryptocurrency_deposits") ,
                        MkString("wallets/{walletId}/cryptocurrency_withdrawals") ,
                        MkString("wallets/{walletId}/orders") ,
                        MkString("wire_withdrawal") ,
                        }),
                    "delete":MkArray(&VarArray{
                        MkString("wallets/{walletId}/orders/{id}") ,
                        }),
                    }),
                }),
            "markets":MkMap(&VarMap{
                "BTC/USD":MkMap(&VarMap{
                    "id":MkString("XBTUSD") ,
                    "symbol":MkString("BTC/USD") ,
                    "base":MkString("BTC") ,
                    "quote":MkString("USD") ,
                    "baseId":MkString("XBT") ,
                    "quoteId":MkString("USD") ,
                    }),
                "BTC/SGD":MkMap(&VarMap{
                    "id":MkString("XBTSGD") ,
                    "symbol":MkString("BTC/SGD") ,
                    "base":MkString("BTC") ,
                    "quote":MkString("SGD") ,
                    "baseId":MkString("XBT") ,
                    "quoteId":MkString("SGD") ,
                    }),
                "BTC/EUR":MkMap(&VarMap{
                    "id":MkString("XBTEUR") ,
                    "symbol":MkString("BTC/EUR") ,
                    "base":MkString("BTC") ,
                    "quote":MkString("EUR") ,
                    "baseId":MkString("XBT") ,
                    "quoteId":MkString("EUR") ,
                    }),
                "ETH/USD":MkMap(&VarMap{
                    "id":MkString("ETHUSD") ,
                    "symbol":MkString("ETH/USD") ,
                    "base":MkString("ETH") ,
                    "quote":MkString("USD") ,
                    "baseId":MkString("ETH") ,
                    "quoteId":MkString("USD") ,
                    }),
                "ETH/EUR":MkMap(&VarMap{
                    "id":MkString("ETHEUR") ,
                    "symbol":MkString("ETH/EUR") ,
                    "base":MkString("ETH") ,
                    "quote":MkString("EUR") ,
                    "baseId":MkString("ETH") ,
                    "quoteId":MkString("EUR") ,
                    }),
                "ETH/SGD":MkMap(&VarMap{
                    "id":MkString("ETHSGD") ,
                    "symbol":MkString("ETH/SGD") ,
                    "base":MkString("ETH") ,
                    "quote":MkString("SGD") ,
                    "baseId":MkString("ETH") ,
                    "quoteId":MkString("SGD") ,
                    }),
                "PAXGUSD":MkMap(&VarMap{
                    "id":MkString("PAXGUSD") ,
                    "symbol":MkString("PAXG/USD") ,
                    "base":MkString("PAXG") ,
                    "quote":MkString("USD") ,
                    "baseId":MkString("PAXG") ,
                    "quoteId":MkString("USD") ,
                    }),
                "BCHUSD":MkMap(&VarMap{
                    "id":MkString("BCHUSD") ,
                    "symbol":MkString("BCH/USD") ,
                    "base":MkString("BCH") ,
                    "quote":MkString("USD") ,
                    "baseId":MkString("BCH") ,
                    "quoteId":MkString("USD") ,
                    }),
                "LTCUSD":MkMap(&VarMap{
                    "id":MkString("LTCUSD") ,
                    "symbol":MkString("LTC/USD") ,
                    "base":MkString("LTC") ,
                    "quote":MkString("USD") ,
                    "baseId":MkString("LTC") ,
                    "quoteId":MkString("USD") ,
                    }),
                }),
            "fees":MkMap(&VarMap{"trading":MkMap(&VarMap{
                    "maker":OpDiv(OpNeg(MkNumber(0.03) ), MkInteger(100) ),
                    "taker":OpDiv(MkNumber(0.35) , MkInteger(100) ),
                    })}),
            "commonCurrencies":MkMap(&VarMap{"XBT":MkString("BTC") }),
            }));
}

func (this *Itbit) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"symbol":this.MarketId(symbol )}); _ = request;
  orderbook := this.Call(MkString("publicGetMarketsSymbolOrderBook"), this.Extend(request , params )); _ = orderbook;
  return this.ParseOrderBook(orderbook , symbol );
}

func (this *Itbit) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"symbol":this.MarketId(symbol )}); _ = request;
  ticker := this.Call(MkString("publicGetMarketsSymbolTicker"), this.Extend(request , params )); _ = ticker;
  serverTimeUTC := this.SafeString(ticker , MkString("serverTimeUTC") ); _ = serverTimeUTC;
  if IsTrue(OpNot(serverTimeUTC )) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" fetchTicker returned a bad response: ") , this.Json(ticker )))));
  }
  timestamp := this.Parse8601(serverTimeUTC ); _ = timestamp;
  vwap := this.SafeNumber(ticker , MkString("vwap24h") ); _ = vwap;
  baseVolume := this.SafeNumber(ticker , MkString("volume24h") ); _ = baseVolume;
  quoteVolume := OpCopy(MkUndefined() ); _ = quoteVolume;
  if IsTrue(OpAnd(OpNotEq2(baseVolume , MkUndefined() ), OpNotEq2(vwap , MkUndefined() ))) {
    quoteVolume = OpMulti(baseVolume , vwap );
  }
  last := this.SafeNumber(ticker , MkString("lastPrice") ); _ = last;
  return MkMap(&VarMap{
        "symbol":symbol ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "high":this.SafeNumber(ticker , MkString("high24h") ),
        "low":this.SafeNumber(ticker , MkString("low24h") ),
        "bid":this.SafeNumber(ticker , MkString("bid") ),
        "bidVolume":MkUndefined() ,
        "ask":this.SafeNumber(ticker , MkString("ask") ),
        "askVolume":MkUndefined() ,
        "vwap":vwap ,
        "open":this.SafeNumber(ticker , MkString("openToday") ),
        "close":last ,
        "last":last ,
        "previousClose":MkUndefined() ,
        "change":MkUndefined() ,
        "percentage":MkUndefined() ,
        "average":MkUndefined() ,
        "baseVolume":baseVolume ,
        "quoteVolume":quoteVolume ,
        "info":ticker ,
        });
}

func (this *Itbit) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  id := this.SafeString2(trade , MkString("executionId") , MkString("matchNumber") ); _ = id;
  timestamp := this.Parse8601(this.SafeString(trade , MkString("timestamp") )); _ = timestamp;
  side := this.SafeString(trade , MkString("direction") ); _ = side;
  orderId := this.SafeString(trade , MkString("orderId") ); _ = orderId;
  feeCost := this.SafeNumber(trade , MkString("commissionPaid") ); _ = feeCost;
  feeCurrencyId := this.SafeString(trade , MkString("commissionCurrency") ); _ = feeCurrencyId;
  feeCurrency := this.SafeCurrencyCode(feeCurrencyId ); _ = feeCurrency;
  rebatesApplied := this.SafeNumber(trade , MkString("rebatesApplied") ); _ = rebatesApplied;
  if IsTrue(OpNotEq2(rebatesApplied , MkUndefined() )) {
    rebatesApplied = OpNeg(rebatesApplied );
  }
  rebateCurrencyId := this.SafeString(trade , MkString("rebateCurrency") ); _ = rebateCurrencyId;
  rebateCurrency := this.SafeCurrencyCode(rebateCurrencyId ); _ = rebateCurrency;
  priceString := this.SafeString2(trade , MkString("price") , MkString("rate") ); _ = priceString;
  amountString := this.SafeString2(trade , MkString("currency1Amount") , MkString("amount") ); _ = amountString;
  price := this.ParseNumber(priceString ); _ = price;
  amount := this.ParseNumber(amountString ); _ = amount;
  cost := this.ParseNumber(Precise.StringMul(priceString , amountString )); _ = cost;
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  marketId := this.SafeString(trade , MkString("instrument") ); _ = marketId;
  if IsTrue(OpNotEq2(marketId , MkUndefined() )) {
    if IsTrue(OpHasMember(marketId , *this.At(MkString("markets_by_id")) )) {
      market = *(*this.At(MkString("markets_by_id")) ).At(marketId );
    } else {
      baseId := this.SafeString(trade , MkString("currency1") ); _ = baseId;
      quoteId := this.SafeString(trade , MkString("currency2") ); _ = quoteId;
      base := this.SafeCurrencyCode(baseId ); _ = base;
      quote := this.SafeCurrencyCode(quoteId ); _ = quote;
      symbol = OpAdd(base , OpAdd(MkString("/") , quote ));
    }
  }
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    if IsTrue(OpNotEq2(market , MkUndefined() )) {
      symbol = *(market ).At(MkString("symbol") );
    }
  }
  result := MkMap(&VarMap{
        "info":trade ,
        "id":id ,
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
        }); _ = result;
  if IsTrue(OpNotEq2(feeCost , MkUndefined() )) {
    if IsTrue(OpNotEq2(rebatesApplied , MkUndefined() )) {
      if IsTrue(OpEq2(feeCurrency , rebateCurrency )) {
        feeCost = this.Sum(feeCost , rebatesApplied );
        *(result ).At(MkString("fee") )= MkMap(&VarMap{
            "cost":feeCost ,
            "currency":feeCurrency ,
            });
      } else {
        *(result ).At(MkString("fees") )= MkArray(&VarArray{
            MkMap(&VarMap{
                "cost":feeCost ,
                "currency":feeCurrency ,
                }),
            MkMap(&VarMap{
                "cost":rebatesApplied ,
                "currency":rebateCurrency ,
                }),
            });
      }
    } else {
      *(result ).At(MkString("fee") )= MkMap(&VarMap{
          "cost":feeCost ,
          "currency":feeCurrency ,
          });
    }
  }
  if IsTrue(OpNot(OpHasMember(MkString("fee") , result ))) {
    if IsTrue(OpNot(OpHasMember(MkString("fees") , result ))) {
      *(result ).At(MkString("fee") )= OpCopy(MkUndefined() );
    }
  }
  return result ;
}

func (this *Itbit) FetchTransactions(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  walletId := this.SafeString(params , MkString("walletId") ); _ = walletId;
  if IsTrue(OpEq2(walletId , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchMyTrades() requires a walletId parameter") )));
  }
  request := MkMap(&VarMap{"walletId":walletId }); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("perPage") )= OpCopy(limit );
  }
  response := this.Call(MkString("privateGetWalletsWalletIdFundingHistory"), this.Extend(request , params )); _ = response;
  items := *(response ).At(MkString("fundingHistory") ); _ = items;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , items.Length )); OpIncr(& i ){
    item := *(items ).At(i ); _ = item;
    time := this.SafeString(item , MkString("time") ); _ = time;
    timestamp := this.Parse8601(time ); _ = timestamp;
    currency := this.SafeString(item , MkString("currency") ); _ = currency;
    destinationAddress := this.SafeString(item , MkString("destinationAddress") ); _ = destinationAddress;
    txnHash := this.SafeString(item , MkString("txnHash") ); _ = txnHash;
    transactionType := this.SafeStringLower(item , MkString("transactionType") ); _ = transactionType;
    transactionStatus := this.SafeString(item , MkString("status") ); _ = transactionStatus;
    status := this.ParseTransferStatus(transactionStatus ); _ = status;
    result.Push(MkMap(&VarMap{
            "id":this.SafeString(item , MkString("withdrawalId") ),
            "timestamp":timestamp ,
            "datetime":this.Iso8601(timestamp ),
            "currency":this.SafeCurrencyCode(currency ),
            "address":destinationAddress ,
            "tag":MkUndefined() ,
            "txid":txnHash ,
            "type":transactionType ,
            "status":status ,
            "amount":this.SafeNumber(item , MkString("amount") ),
            "fee":MkUndefined() ,
            "info":item ,
            }));
  }
  return result ;
}

func (this *Itbit) ParseTransferStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  options := MkMap(&VarMap{
        "cancelled":MkString("canceled") ,
        "completed":MkString("ok") ,
        }); _ = options;
  return this.SafeString(options , status , MkString("pending") );
}

func (this *Itbit) FetchMyTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  walletId := this.SafeString(params , MkString("walletId") ); _ = walletId;
  if IsTrue(OpEq2(walletId , MkUndefined() )) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , MkString(" fetchMyTrades() requires a walletId parameter") )));
  }
  request := MkMap(&VarMap{"walletId":walletId }); _ = request;
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("rangeStart") )= this.Ymdhms(since , MkString("T") );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("perPage") )= OpCopy(limit );
  }
  response := this.Call(MkString("privateGetWalletsWalletIdTrades"), this.Extend(request , params )); _ = response;
  trades := this.SafeValue(response , MkString("tradingHistory") , MkArray(&VarArray{})); _ = trades;
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
  }
  return this.ParseTrades(trades , market , since , limit );
}

func (this *Itbit) FetchTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"symbol":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("publicGetMarketsSymbolTrades"), this.Extend(request , params )); _ = response;
  trades := this.SafeValue(response , MkString("recentTrades") , MkArray(&VarArray{})); _ = trades;
  return this.ParseTrades(trades , market , since , limit );
}

func (this *Itbit) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.FetchWallets(params ); _ = response;
  balances := *(*(response ).At(MkInteger(0) )).At(MkString("balances") ); _ = balances;
  result := MkMap(&VarMap{"info":response }); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , balances.Length )); OpIncr(& i ){
    balance := *(balances ).At(i ); _ = balance;
    currencyId := this.SafeString(balance , MkString("currency") ); _ = currencyId;
    code := this.SafeCurrencyCode(currencyId ); _ = code;
    account := this.Account(); _ = account;
    *(account ).At(MkString("free") )= this.SafeString(balance , MkString("availableBalance") );
    *(account ).At(MkString("total") )= this.SafeString(balance , MkString("totalBalance") );
    *(result ).At(code )= OpCopy(account );
  }
  return this.ParseBalance(result );
}

func (this *Itbit) FetchWallets(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  if IsTrue(OpNot(*this.At(MkString("uid")) )) {
    panic(NewAuthenticationError(OpAdd(*this.At(MkString("id")) , MkString(" fetchWallets() requires uid API credential") )));
  }
  request := MkMap(&VarMap{"userId":*this.At(MkString("uid")) }); _ = request;
  return this.Call(MkString("privateGetWallets"), this.Extend(request , params ));
}

func (this *Itbit) FetchWallet(goArgs ...*Variant) *Variant{
  walletId := GoGetArg(goArgs, 0, MkUndefined()); _ = walletId;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"walletId":walletId }); _ = request;
  return this.Call(MkString("privateGetWalletsWalletId"), this.Extend(request , params ));
}

func (this *Itbit) FetchOpenOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"status":MkString("open") }); _ = request;
  return this.FetchOrders(symbol , since , limit , this.Extend(request , params ));
}

func (this *Itbit) FetchClosedOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"status":MkString("filled") }); _ = request;
  return this.FetchOrders(symbol , since , limit , this.Extend(request , params ));
}

func (this *Itbit) FetchOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
  }
  walletIdInParams := OpHasMember(MkString("walletId") , params ); _ = walletIdInParams;
  if IsTrue(OpNot(walletIdInParams )) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , MkString(" fetchOrders() requires a walletId parameter") )));
  }
  walletId := *(params ).At(MkString("walletId") ); _ = walletId;
  request := MkMap(&VarMap{"walletId":walletId }); _ = request;
  response := this.Call(MkString("privateGetWalletsWalletIdOrders"), this.Extend(request , params )); _ = response;
  return this.ParseOrders(response , market , since , limit );
}

func (this *Itbit) ParseOrderStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "submitted":MkString("open") ,
        "open":MkString("open") ,
        "filled":MkString("closed") ,
        "cancelled":MkString("canceled") ,
        "rejected":MkString("canceled") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Itbit) ParseOrder(goArgs ...*Variant) *Variant{
  order := GoGetArg(goArgs, 0, MkUndefined()); _ = order;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  side := this.SafeString(order , MkString("side") ); _ = side;
  type_ := this.SafeString(order , MkString("type") ); _ = type_;
  symbol := *(*(*this.At(MkString("markets_by_id")) ).At(*(order ).At(MkString("instrument") ))).At(MkString("symbol") ); _ = symbol;
  timestamp := this.Parse8601(*(order ).At(MkString("createdTime") )); _ = timestamp;
  amount := this.SafeNumber(order , MkString("amount") ); _ = amount;
  filled := this.SafeNumber(order , MkString("amountFilled") ); _ = filled;
  fee := OpCopy(MkUndefined() ); _ = fee;
  price := this.SafeNumber(order , MkString("price") ); _ = price;
  average := this.SafeNumber(order , MkString("volumeWeightedAveragePrice") ); _ = average;
  clientOrderId := this.SafeString(order , MkString("clientOrderIdentifier") ); _ = clientOrderId;
  id := this.SafeString(order , MkString("id") ); _ = id;
  postOnlyString := this.SafeString(order , MkString("postOnly") ); _ = postOnlyString;
  postOnly := OpEq2(postOnlyString , MkString("True") ); _ = postOnly;
  return this.SafeOrder(MkMap(&VarMap{
            "id":id ,
            "clientOrderId":clientOrderId ,
            "info":order ,
            "timestamp":timestamp ,
            "datetime":this.Iso8601(timestamp ),
            "lastTradeTimestamp":MkUndefined() ,
            "status":this.ParseOrderStatus(this.SafeString(order , MkString("status") )),
            "symbol":symbol ,
            "type":type_ ,
            "timeInForce":MkUndefined() ,
            "postOnly":postOnly ,
            "side":side ,
            "price":price ,
            "stopPrice":MkUndefined() ,
            "cost":MkUndefined() ,
            "average":average ,
            "amount":amount ,
            "filled":filled ,
            "remaining":MkUndefined() ,
            "fee":fee ,
            "trades":MkUndefined() ,
            }));
}

func (this *Itbit) Nonce(goArgs ...*Variant) *Variant{
  return this.Milliseconds();
}

func (this *Itbit) CreateOrder(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 2, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 3, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 4, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 5, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  if IsTrue(OpEq2(type_ , MkString("market") )) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , MkString(" allows limit orders only") )));
  }
  walletIdInParams := OpHasMember(MkString("walletId") , params ); _ = walletIdInParams;
  if IsTrue(OpNot(walletIdInParams )) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , MkString(" createOrder() requires a walletId parameter") )));
  }
  amount = amount.ToString();
  price = price.ToString();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "side":side ,
        "type":type_ ,
        "currency":(*(market ).At(MkString("id") )).Call(MkString("replace") , *(market ).At(MkString("quote") ), MkString("") ),
        "amount":amount ,
        "display":amount ,
        "price":price ,
        "instrument":*(market ).At(MkString("id") ),
        }); _ = request;
  response := this.Call(MkString("privatePostWalletsWalletIdOrders"), this.Extend(request , params )); _ = response;
  return MkMap(&VarMap{
        "info":response ,
        "id":*(response ).At(MkString("id") ),
        });
}

func (this *Itbit) FetchOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  walletIdInParams := OpHasMember(MkString("walletId") , params ); _ = walletIdInParams;
  if IsTrue(OpNot(walletIdInParams )) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , MkString(" fetchOrder() requires a walletId parameter") )));
  }
  request := MkMap(&VarMap{"id":id }); _ = request;
  response := this.Call(MkString("privateGetWalletsWalletIdOrdersId"), this.Extend(request , params )); _ = response;
  return this.ParseOrder(response );
}

func (this *Itbit) CancelOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  walletIdInParams := OpHasMember(MkString("walletId") , params ); _ = walletIdInParams;
  if IsTrue(OpNot(walletIdInParams )) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , MkString(" cancelOrder() requires a walletId parameter") )));
  }
  request := MkMap(&VarMap{"id":id }); _ = request;
  return this.Call(MkString("privateDeleteWalletsWalletIdOrdersId"), this.Extend(request , params ));
}

func (this *Itbit) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("public") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  url := OpAdd(*(*this.At(MkString("urls")) ).At(MkString("api") ), OpAdd(MkString("/") , OpAdd(*this.At(MkString("version")) , OpAdd(MkString("/") , this.ImplodeParams(path , params ))))); _ = url;
  query := this.Omit(params , this.ExtractParams(path )); _ = query;
  if IsTrue(OpAnd(OpEq2(method , MkString("GET") ), *(GoGetKeys(query )).At(MkString("length") ))) {
    OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(query )));
  }
  if IsTrue(OpAnd(OpEq2(method , MkString("POST") ), *(GoGetKeys(query )).At(MkString("length") ))) {
    body = this.Json(query );
  }
  if IsTrue(OpEq2(api , MkString("private") )) {
    this.CheckRequiredCredentials();
    nonce := (this.Nonce()).Call(MkString("toString") ); _ = nonce;
    timestamp := OpCopy(nonce ); _ = timestamp;
    authBody := OpTrinary(OpEq2(method , MkString("POST") ), body , MkString("") ); _ = authBody;
    auth := MkArray(&VarArray{
          method ,
          url ,
          authBody ,
          nonce ,
          timestamp ,
          }); _ = auth;
    message := OpAdd(nonce , (this.Json(auth )).Call(MkString("replace") , MkString("\\/") , MkString("/") )); _ = message;
    hash := this.Hash(this.Encode(message ), MkString("sha256") , MkString("binary") ); _ = hash;
    binaryUrl := this.StringToBinary(this.Encode(url )); _ = binaryUrl;
    binhash := this.BinaryConcat(binaryUrl , hash ); _ = binhash;
    signature := this.Hmac(binhash , this.Encode(*this.At(MkString("secret")) ), MkString("sha512") , MkString("base64") ); _ = signature;
    headers = MkMap(&VarMap{
        "Authorization":OpAdd(*this.At(MkString("apiKey")) , OpAdd(MkString(":") , signature )),
        "Content-Type":MkString("application/json") ,
        "X-Auth-Timestamp":timestamp ,
        "X-Auth-Nonce":nonce ,
        });
  }
  return MkMap(&VarMap{
        "url":url ,
        "method":method ,
        "body":body ,
        "headers":headers ,
        });
}

func (this *Itbit) HandleErrors(goArgs ...*Variant) *Variant{
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
  code := this.SafeString(response , MkString("code") ); _ = code;
  if IsTrue(OpNotEq2(code , MkUndefined() )) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , this.Json(response )))));
  }
  return MkUndefined()
}

