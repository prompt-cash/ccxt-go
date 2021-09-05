package ccxt

import (
)

type Coinmate struct {
	*ExchangeBase
}

var _ Exchange = (*Coinmate)(nil)

func init() {
	exchange := &Coinmate{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Coinmate) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("coinmate") ,
            "name":MkString("CoinMate") ,
            "countries":MkArray(&VarArray{
                MkString("GB") ,
                MkString("CZ") ,
                MkString("EU") ,
                }),
            "rateLimit":MkInteger(1000) ,
            "has":MkMap(&VarMap{
                "cancelOrder":MkBool(true) ,
                "CORS":MkBool(true) ,
                "createOrder":MkBool(true) ,
                "fetchBalance":MkBool(true) ,
                "fetchMarkets":MkBool(true) ,
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
                "logo":MkString("https://user-images.githubusercontent.com/51840849/87460806-1c9f3f00-c616-11ea-8c46-a77018a8f3f4.jpg") ,
                "api":MkString("https://coinmate.io/api") ,
                "www":MkString("https://coinmate.io") ,
                "fees":MkString("https://coinmate.io/fees") ,
                "doc":MkArray(&VarArray{
                    MkString("https://coinmate.docs.apiary.io") ,
                    MkString("https://coinmate.io/developers") ,
                    }),
                "referral":MkString("https://coinmate.io?referral=YTFkM1RsOWFObVpmY1ZjMGREQmpTRnBsWjJJNVp3PT0") ,
                }),
            "requiredCredentials":MkMap(&VarMap{
                "apiKey":MkBool(true) ,
                "secret":MkBool(true) ,
                "uid":MkBool(true) ,
                }),
            "api":MkMap(&VarMap{
                "public":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("orderBook") ,
                        MkString("ticker") ,
                        MkString("transactions") ,
                        MkString("tradingPairs") ,
                        })}),
                "private":MkMap(&VarMap{"post":MkArray(&VarArray{
                        MkString("balances") ,
                        MkString("bitcoinCashWithdrawal") ,
                        MkString("bitcoinCashDepositAddresses") ,
                        MkString("bitcoinDepositAddresses") ,
                        MkString("bitcoinWithdrawal") ,
                        MkString("bitcoinWithdrawalFees") ,
                        MkString("buyInstant") ,
                        MkString("buyLimit") ,
                        MkString("cancelOrder") ,
                        MkString("cancelOrderWithInfo") ,
                        MkString("createVoucher") ,
                        MkString("dashDepositAddresses") ,
                        MkString("dashWithdrawal") ,
                        MkString("ethereumWithdrawal") ,
                        MkString("ethereumDepositAddresses") ,
                        MkString("litecoinWithdrawal") ,
                        MkString("litecoinDepositAddresses") ,
                        MkString("openOrders") ,
                        MkString("order") ,
                        MkString("orderHistory") ,
                        MkString("orderById") ,
                        MkString("pusherAuth") ,
                        MkString("redeemVoucher") ,
                        MkString("replaceByBuyLimit") ,
                        MkString("replaceByBuyInstant") ,
                        MkString("replaceBySellLimit") ,
                        MkString("replaceBySellInstant") ,
                        MkString("rippleDepositAddresses") ,
                        MkString("rippleWithdrawal") ,
                        MkString("sellInstant") ,
                        MkString("sellLimit") ,
                        MkString("transactionHistory") ,
                        MkString("traderFees") ,
                        MkString("tradeHistory") ,
                        MkString("transfer") ,
                        MkString("transferHistory") ,
                        MkString("unconfirmedBitcoinDeposits") ,
                        MkString("unconfirmedBitcoinCashDeposits") ,
                        MkString("unconfirmedDashDeposits") ,
                        MkString("unconfirmedEthereumDeposits") ,
                        MkString("unconfirmedLitecoinDeposits") ,
                        MkString("unconfirmedRippleDeposits") ,
                        })}),
                }),
            "fees":MkMap(&VarMap{
                "trading":MkMap(&VarMap{
                    "tierBased":MkBool(true) ,
                    "percentage":MkBool(true) ,
                    "maker":OpDiv(MkNumber(0.12) , MkInteger(100) ),
                    "taker":OpDiv(MkNumber(0.25) , MkInteger(100) ),
                    "tiers":MkMap(&VarMap{
                        "taker":MkArray(&VarArray{
                            MkArray(&VarArray{
                                MkInteger(0) ,
                                OpDiv(MkNumber(0.25) , MkInteger(100) ),
                                }),
                            MkArray(&VarArray{
                                MkInteger(10000) ,
                                OpDiv(MkNumber(0.23) , MkInteger(100) ),
                                }),
                            MkArray(&VarArray{
                                MkInteger(100000) ,
                                OpDiv(MkNumber(0.21) , MkInteger(100) ),
                                }),
                            MkArray(&VarArray{
                                MkInteger(250000) ,
                                OpDiv(MkNumber(0.20) , MkInteger(100) ),
                                }),
                            MkArray(&VarArray{
                                MkInteger(500000) ,
                                OpDiv(MkNumber(0.15) , MkInteger(100) ),
                                }),
                            MkArray(&VarArray{
                                MkInteger(1000000) ,
                                OpDiv(MkNumber(0.13) , MkInteger(100) ),
                                }),
                            MkArray(&VarArray{
                                MkInteger(3000000) ,
                                OpDiv(MkNumber(0.10) , MkInteger(100) ),
                                }),
                            MkArray(&VarArray{
                                MkInteger(15000000) ,
                                OpDiv(MkNumber(0.05) , MkInteger(100) ),
                                }),
                            }),
                        "maker":MkArray(&VarArray{
                            MkArray(&VarArray{
                                MkInteger(0) ,
                                OpDiv(MkNumber(0.12) , MkInteger(100) ),
                                }),
                            MkArray(&VarArray{
                                MkInteger(10000) ,
                                OpDiv(MkNumber(0.11) , MkInteger(100) ),
                                }),
                            MkArray(&VarArray{
                                MkInteger(1000000) ,
                                OpDiv(MkNumber(0.10) , MkInteger(100) ),
                                }),
                            MkArray(&VarArray{
                                MkInteger(250000) ,
                                OpDiv(MkNumber(0.08) , MkInteger(100) ),
                                }),
                            MkArray(&VarArray{
                                MkInteger(500000) ,
                                OpDiv(MkNumber(0.05) , MkInteger(100) ),
                                }),
                            MkArray(&VarArray{
                                MkInteger(1000000) ,
                                OpDiv(MkNumber(0.03) , MkInteger(100) ),
                                }),
                            MkArray(&VarArray{
                                MkInteger(3000000) ,
                                OpDiv(MkNumber(0.02) , MkInteger(100) ),
                                }),
                            MkArray(&VarArray{
                                MkInteger(15000000) ,
                                MkInteger(0) ,
                                }),
                            }),
                        }),
                    }),
                "promotional":MkMap(&VarMap{"trading":MkMap(&VarMap{
                        "maker":OpDiv(MkNumber(0.05) , MkInteger(100) ),
                        "taker":OpDiv(MkNumber(0.15) , MkInteger(100) ),
                        "tiers":MkMap(&VarMap{
                            "taker":MkArray(&VarArray{
                                MkArray(&VarArray{
                                    MkInteger(0) ,
                                    OpDiv(MkNumber(0.15) , MkInteger(100) ),
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(10000) ,
                                    OpDiv(MkNumber(0.14) , MkInteger(100) ),
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(100000) ,
                                    OpDiv(MkNumber(0.13) , MkInteger(100) ),
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(250000) ,
                                    OpDiv(MkNumber(0.12) , MkInteger(100) ),
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(500000) ,
                                    OpDiv(MkNumber(0.11) , MkInteger(100) ),
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(1000000) ,
                                    OpDiv(MkNumber(0.1) , MkInteger(100) ),
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(3000000) ,
                                    OpDiv(MkNumber(0.08) , MkInteger(100) ),
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(15000000) ,
                                    OpDiv(MkNumber(0.05) , MkInteger(100) ),
                                    }),
                                }),
                            "maker":MkArray(&VarArray{
                                MkArray(&VarArray{
                                    MkInteger(0) ,
                                    OpDiv(MkNumber(0.05) , MkInteger(100) ),
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(10000) ,
                                    OpDiv(MkNumber(0.04) , MkInteger(100) ),
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(1000000) ,
                                    OpDiv(MkNumber(0.03) , MkInteger(100) ),
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(250000) ,
                                    OpDiv(MkNumber(0.02) , MkInteger(100) ),
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(500000) ,
                                    MkInteger(0) ,
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(1000000) ,
                                    MkInteger(0) ,
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(3000000) ,
                                    MkInteger(0) ,
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(15000000) ,
                                    MkInteger(0) ,
                                    }),
                                }),
                            }),
                        })}),
                }),
            "options":MkMap(&VarMap{"promotionalMarkets":MkArray(&VarArray{
                    MkString("ETH/EUR") ,
                    MkString("ETH/CZK") ,
                    MkString("ETH/BTC") ,
                    MkString("XRP/EUR") ,
                    MkString("XRP/CZK") ,
                    MkString("XRP/BTC") ,
                    MkString("DASH/EUR") ,
                    MkString("DASH/CZK") ,
                    MkString("DASH/BTC") ,
                    MkString("BCH/EUR") ,
                    MkString("BCH/CZK") ,
                    MkString("BCH/BTC") ,
                    })}),
            "exceptions":MkMap(&VarMap{
                "exact":MkMap(&VarMap{"No order with given ID":OrderNotFound }),
                "broad":MkMap(&VarMap{
                    "Not enough account balance available":InsufficientFunds ,
                    "Incorrect order ID":InvalidOrder ,
                    "Minimum Order Size ":InvalidOrder ,
                    "TOO MANY REQUESTS":RateLimitExceeded ,
                    }),
                }),
            }));
}

func (this *Coinmate) FetchMarkets(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetTradingPairs"), params ); _ = response;
  data := this.SafeValue(response , MkString("data") ); _ = data;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , data.Length )); OpIncr(& i ){
    market := *(data ).At(i ); _ = market;
    id := this.SafeString(market , MkString("name") ); _ = id;
    baseId := this.SafeString(market , MkString("firstCurrency") ); _ = baseId;
    quoteId := this.SafeString(market , MkString("secondCurrency") ); _ = quoteId;
    base := this.SafeCurrencyCode(baseId ); _ = base;
    quote := this.SafeCurrencyCode(quoteId ); _ = quote;
    symbol := OpAdd(base , OpAdd(MkString("/") , quote )); _ = symbol;
    promotionalMarkets := this.SafeValue(*this.At(MkString("options")) , MkString("promotionalMarkets") , MkArray(&VarArray{})); _ = promotionalMarkets;
    fees := this.SafeValue(*this.At(MkString("fees")) , MkString("trading") ); _ = fees;
    if IsTrue(this.InArray(symbol , promotionalMarkets )) {
      promotionalFees := this.SafeValue(*this.At(MkString("fees")) , MkString("promotional") , MkMap(&VarMap{})); _ = promotionalFees;
      fees = this.SafeValue(promotionalFees , MkString("trading") , fees );
    }
    result.Push(MkMap(&VarMap{
            "id":id ,
            "symbol":symbol ,
            "base":base ,
            "quote":quote ,
            "baseId":baseId ,
            "quoteId":quoteId ,
            "active":MkUndefined() ,
            "maker":*(fees ).At(MkString("maker") ),
            "taker":*(fees ).At(MkString("taker") ),
            "info":market ,
            "precision":MkMap(&VarMap{
                "price":this.SafeInteger(market , MkString("priceDecimals") ),
                "amount":this.SafeInteger(market , MkString("lotDecimals") ),
                }),
            "limits":MkMap(&VarMap{
                "amount":MkMap(&VarMap{
                    "min":this.SafeNumber(market , MkString("minAmount") ),
                    "max":MkUndefined() ,
                    }),
                "price":MkMap(&VarMap{
                    "min":MkUndefined() ,
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

func (this *Coinmate) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("privatePostBalances"), params ); _ = response;
  balances := this.SafeValue(response , MkString("data") ); _ = balances;
  result := MkMap(&VarMap{"info":response }); _ = result;
  currencyIds := GoGetKeys(balances ); _ = currencyIds;
  for i := MkInteger(0) ; IsTrue(OpLw(i , currencyIds.Length )); OpIncr(& i ){
    currencyId := *(currencyIds ).At(i ); _ = currencyId;
    code := this.SafeCurrencyCode(currencyId ); _ = code;
    balance := this.SafeValue(balances , currencyId ); _ = balance;
    account := this.Account(); _ = account;
    *(account ).At(MkString("free") )= this.SafeString(balance , MkString("available") );
    *(account ).At(MkString("used") )= this.SafeString(balance , MkString("reserved") );
    *(account ).At(MkString("total") )= this.SafeString(balance , MkString("balance") );
    *(result ).At(code )= OpCopy(account );
  }
  return this.ParseBalance(result );
}

func (this *Coinmate) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{
        "currencyPair":this.MarketId(symbol ),
        "groupByPriceLimit":MkString("False") ,
        }); _ = request;
  response := this.Call(MkString("publicGetOrderBook"), this.Extend(request , params )); _ = response;
  orderbook := *(response ).At(MkString("data") ); _ = orderbook;
  timestamp := this.SafeTimestamp(orderbook , MkString("timestamp") ); _ = timestamp;
  return this.ParseOrderBook(orderbook , symbol , timestamp , MkString("bids") , MkString("asks") , MkString("price") , MkString("amount") );
}

func (this *Coinmate) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"currencyPair":this.MarketId(symbol )}); _ = request;
  response := this.Call(MkString("publicGetTicker"), this.Extend(request , params )); _ = response;
  ticker := this.SafeValue(response , MkString("data") ); _ = ticker;
  timestamp := this.SafeTimestamp(ticker , MkString("timestamp") ); _ = timestamp;
  last := this.SafeNumber(ticker , MkString("last") ); _ = last;
  return MkMap(&VarMap{
        "symbol":symbol ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "high":this.SafeNumber(ticker , MkString("high") ),
        "low":this.SafeNumber(ticker , MkString("low") ),
        "bid":this.SafeNumber(ticker , MkString("bid") ),
        "bidVolume":MkUndefined() ,
        "ask":this.SafeNumber(ticker , MkString("ask") ),
        "vwap":MkUndefined() ,
        "askVolume":MkUndefined() ,
        "open":MkUndefined() ,
        "close":last ,
        "last":last ,
        "previousClose":MkUndefined() ,
        "change":MkUndefined() ,
        "percentage":MkUndefined() ,
        "average":MkUndefined() ,
        "baseVolume":this.SafeNumber(ticker , MkString("amount") ),
        "quoteVolume":MkUndefined() ,
        "info":ticker ,
        });
}

func (this *Coinmate) FetchTransactions(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"limit":MkInteger(1000) }); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("timestampFrom") )= OpCopy(since );
  }
  if IsTrue(OpNotEq2(code , MkUndefined() )) {
    currency := this.Currency(code ); _ = currency;
    *(request ).At(MkString("currency") )= *(currency ).At(MkString("id") );
  }
  response := this.Call(MkString("privatePostTransferHistory"), this.Extend(request , params )); _ = response;
  items := *(response ).At(MkString("data") ); _ = items;
  return this.ParseTransactions(items , MkUndefined() , since , limit );
}

func (this *Coinmate) ParseTransactionStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{"COMPLETED":MkString("ok") }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Coinmate) ParseTransaction(goArgs ...*Variant) *Variant{
  item := GoGetArg(goArgs, 0, MkUndefined()); _ = item;
  currency := GoGetArg(goArgs, 1, MkUndefined() ); _ = currency;
  timestamp := this.SafeInteger(item , MkString("timestamp") ); _ = timestamp;
  amount := this.SafeNumber(item , MkString("amount") ); _ = amount;
  fee := this.SafeNumber(item , MkString("fee") ); _ = fee;
  txid := this.SafeString(item , MkString("txid") ); _ = txid;
  address := this.SafeString(item , MkString("destination") ); _ = address;
  tag := this.SafeString(item , MkString("destinationTag") ); _ = tag;
  currencyId := this.SafeString(item , MkString("amountCurrency") ); _ = currencyId;
  code := this.SafeCurrencyCode(currencyId , currency ); _ = code;
  type_ := this.SafeStringLower(item , MkString("transferType") ); _ = type_;
  status := this.ParseTransactionStatus(this.SafeString(item , MkString("transferStatus") )); _ = status;
  id := this.SafeString(item , MkString("transactionId") ); _ = id;
  return MkMap(&VarMap{
        "id":id ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "currency":code ,
        "amount":amount ,
        "type":type_ ,
        "txid":txid ,
        "address":address ,
        "tag":tag ,
        "status":status ,
        "fee":MkMap(&VarMap{
            "cost":fee ,
            "currency":code ,
            }),
        "info":item ,
        });
}

func (this *Coinmate) FetchMyTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  if IsTrue(OpEq2(limit , MkUndefined() )) {
    limit = MkInteger(1000) ;
  }
  request := MkMap(&VarMap{"limit":limit }); _ = request;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market := this.Market(symbol ); _ = market;
    *(request ).At(MkString("currencyPair") )= *(market ).At(MkString("id") );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("timestampFrom") )= OpCopy(since );
  }
  response := this.Call(MkString("privatePostTradeHistory"), this.Extend(request , params )); _ = response;
  items := *(response ).At(MkString("data") ); _ = items;
  return this.ParseTrades(items , MkUndefined() , since , limit );
}

func (this *Coinmate) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  marketId := this.SafeString(trade , MkString("currencyPair") ); _ = marketId;
  market = this.SafeMarket(marketId , market , MkString("_") );
  priceString := this.SafeString(trade , MkString("price") ); _ = priceString;
  amountString := this.SafeString(trade , MkString("amount") ); _ = amountString;
  price := this.ParseNumber(priceString ); _ = price;
  amount := this.ParseNumber(amountString ); _ = amount;
  cost := this.ParseNumber(Precise.StringMul(priceString , amountString )); _ = cost;
  side := this.SafeStringLower2(trade , MkString("type") , MkString("tradeType") ); _ = side;
  type_ := this.SafeStringLower(trade , MkString("orderType") ); _ = type_;
  orderId := this.SafeString(trade , MkString("orderId") ); _ = orderId;
  id := this.SafeString(trade , MkString("transactionId") ); _ = id;
  timestamp := this.SafeInteger2(trade , MkString("timestamp") , MkString("createdTimestamp") ); _ = timestamp;
  fee := OpCopy(MkUndefined() ); _ = fee;
  feeCost := this.SafeNumber(trade , MkString("fee") ); _ = feeCost;
  if IsTrue(OpNotEq2(feeCost , MkUndefined() )) {
    fee = MkMap(&VarMap{
        "cost":feeCost ,
        "currency":*(market ).At(MkString("quote") ),
        });
  }
  takerOrMaker := this.SafeString(trade , MkString("feeType") ); _ = takerOrMaker;
  takerOrMaker = OpTrinary(OpEq2(takerOrMaker , MkString("MAKER") ), MkString("maker") , MkString("taker") );
  return MkMap(&VarMap{
        "id":id ,
        "info":trade ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "symbol":*(market ).At(MkString("symbol") ),
        "type":type_ ,
        "side":side ,
        "order":orderId ,
        "takerOrMaker":takerOrMaker ,
        "price":price ,
        "amount":amount ,
        "cost":cost ,
        "fee":fee ,
        });
}

func (this *Coinmate) FetchTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "currencyPair":*(market ).At(MkString("id") ),
        "minutesIntoHistory":MkInteger(10) ,
        }); _ = request;
  response := this.Call(MkString("publicGetTransactions"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = data;
  return this.ParseTrades(data , market , since , limit );
}

func (this *Coinmate) FetchOpenOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("privatePostOpenOrders"), this.Extend(MkMap(&VarMap{}), params )); _ = response;
  extension := MkMap(&VarMap{"status":MkString("open") }); _ = extension;
  return this.ParseOrders(*(response ).At(MkString("data") ), MkUndefined() , since , limit , extension );
}

func (this *Coinmate) FetchOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchOrders() requires a symbol argument") )));
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"currencyPair":*(market ).At(MkString("id") )}); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("privatePostOrderHistory"), this.Extend(request , params )); _ = response;
  return this.ParseOrders(*(response ).At(MkString("data") ), market , since , limit );
}

func (this *Coinmate) ParseOrderStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "FILLED":MkString("closed") ,
        "CANCELLED":MkString("canceled") ,
        "PARTIALLY_FILLED":MkString("open") ,
        "OPEN":MkString("open") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Coinmate) ParseOrderType(goArgs ...*Variant) *Variant{
  type_ := GoGetArg(goArgs, 0, MkUndefined()); _ = type_;
  types := MkMap(&VarMap{
        "LIMIT":MkString("limit") ,
        "MARKET":MkString("market") ,
        }); _ = types;
  return this.SafeString(types , type_ , type_ );
}

func (this *Coinmate) ParseOrder(goArgs ...*Variant) *Variant{
  order := GoGetArg(goArgs, 0, MkUndefined()); _ = order;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  id := this.SafeString(order , MkString("id") ); _ = id;
  timestamp := this.SafeInteger(order , MkString("timestamp") ); _ = timestamp;
  side := this.SafeStringLower(order , MkString("type") ); _ = side;
  price := this.SafeNumber(order , MkString("price") ); _ = price;
  amount := this.SafeNumber(order , MkString("originalAmount") ); _ = amount;
  remaining := this.SafeNumber(order , MkString("remainingAmount") ); _ = remaining;
  if IsTrue(OpEq2(remaining , MkUndefined() )) {
    remaining = this.SafeNumber(order , MkString("amount") );
  }
  status := this.ParseOrderStatus(this.SafeString(order , MkString("status") )); _ = status;
  type_ := this.ParseOrderType(this.SafeString(order , MkString("orderTradeType") )); _ = type_;
  average := this.SafeNumber(order , MkString("avgPrice") ); _ = average;
  marketId := this.SafeString(order , MkString("currencyPair") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market , MkString("_") ); _ = symbol;
  clientOrderId := this.SafeString(order , MkString("clientOrderId") ); _ = clientOrderId;
  stopPrice := this.SafeNumber(order , MkString("stopPrice") ); _ = stopPrice;
  return this.SafeOrder(MkMap(&VarMap{
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
            "stopPrice":stopPrice ,
            "amount":amount ,
            "cost":MkUndefined() ,
            "average":average ,
            "filled":MkUndefined() ,
            "remaining":remaining ,
            "status":status ,
            "trades":MkUndefined() ,
            "info":order ,
            "fee":MkUndefined() ,
            }));
}

func (this *Coinmate) CreateOrder(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 2, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 3, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 4, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 5, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  method := OpAdd(MkString("privatePost") , this.Capitalize(side )); _ = method;
  request := MkMap(&VarMap{"currencyPair":this.MarketId(symbol )}); _ = request;
  if IsTrue(OpEq2(type_ , MkString("market") )) {
    if IsTrue(OpEq2(side , MkString("buy") )) {
      *(request ).At(MkString("total") )= this.AmountToPrecision(symbol , amount );
    } else {
      *(request ).At(MkString("amount") )= this.AmountToPrecision(symbol , amount );
    }
    OpAddAssign(& method , MkString("Instant") );
  } else {
    *(request ).At(MkString("amount") )= this.AmountToPrecision(symbol , amount );
    *(request ).At(MkString("price") )= this.PriceToPrecision(symbol , price );
    OpAddAssign(& method , this.Capitalize(type_ ));
  }
  response := this.Call(method , this.Extend(request , params )); _ = response;
  id := this.SafeString(response , MkString("data") ); _ = id;
  return MkMap(&VarMap{
        "info":response ,
        "id":id ,
        });
}

func (this *Coinmate) FetchOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"orderId":id }); _ = request;
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(symbol ) {
    market = this.Market(symbol );
  }
  response := this.Call(MkString("privatePostOrderById"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") ); _ = data;
  return this.ParseOrder(data , market );
}

func (this *Coinmate) CancelOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"orderId":id }); _ = request;
  response := this.Call(MkString("privatePostCancelOrderWithInfo"), this.Extend(request , params )); _ = response;
  return MkMap(&VarMap{"info":response });
}

func (this *Coinmate) Nonce(goArgs ...*Variant) *Variant{
  return this.Milliseconds();
}

func (this *Coinmate) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("public") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  url := OpAdd(*(*this.At(MkString("urls")) ).At(MkString("api") ), OpAdd(MkString("/") , path )); _ = url;
  if IsTrue(OpEq2(api , MkString("public") )) {
    if IsTrue(*(GoGetKeys(params )).At(MkString("length") )) {
      OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(params )));
    }
  } else {
    this.CheckRequiredCredentials();
    nonce := (this.Nonce()).Call(MkString("toString") ); _ = nonce;
    auth := OpAdd(nonce , OpAdd(*this.At(MkString("uid")) , *this.At(MkString("apiKey")) )); _ = auth;
    signature := this.Hmac(this.Encode(auth ), this.Encode(*this.At(MkString("secret")) )); _ = signature;
    body = this.Urlencode(this.Extend(MkMap(&VarMap{
                "clientId":*this.At(MkString("uid")) ,
                "nonce":nonce ,
                "publicKey":*this.At(MkString("apiKey")) ,
                "signature":signature.ToUpperCase(),
                }), params ));
    headers = MkMap(&VarMap{"Content-Type":MkString("application/x-www-form-urlencoded") });
  }
  return MkMap(&VarMap{
        "url":url ,
        "method":method ,
        "body":body ,
        "headers":headers ,
        });
}

func (this *Coinmate) HandleErrors(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  reason := GoGetArg(goArgs, 1, MkUndefined()); _ = reason;
  url := GoGetArg(goArgs, 2, MkUndefined()); _ = url;
  method := GoGetArg(goArgs, 3, MkUndefined()); _ = method;
  headers := GoGetArg(goArgs, 4, MkUndefined()); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined()); _ = body;
  response := GoGetArg(goArgs, 6, MkUndefined()); _ = response;
  requestHeaders := GoGetArg(goArgs, 7, MkUndefined()); _ = requestHeaders;
  requestBody := GoGetArg(goArgs, 8, MkUndefined()); _ = requestBody;
  if IsTrue(OpNotEq2(response , MkUndefined() )) {
    if IsTrue(OpHasMember(MkString("error") , response )) {
      if IsTrue(*(response ).At(MkString("error") )) {
        message := this.SafeString(response , MkString("errorMessage") ); _ = message;
        feedback := OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , message )); _ = feedback;
        this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("exact") ), message , feedback );
        this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("broad") ), message , feedback );
        panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , this.Json(response )))));
      }
    }
  }
  if IsTrue(OpGt(code , MkInteger(400) )) {
    if IsTrue(body ) {
      feedback := OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , body )); _ = feedback;
      this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("exact") ), body , feedback );
      this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("broad") ), body , feedback );
      panic(NewExchangeError(feedback ));
    }
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , body ))));
  }
  return MkUndefined()
}

