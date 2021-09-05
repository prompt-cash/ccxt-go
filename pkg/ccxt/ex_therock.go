package ccxt

import (
)

type Therock struct {
	*ExchangeBase
}

var _ Exchange = (*Therock)(nil)

func init() {
	exchange := &Therock{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Therock) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("therock") ,
            "name":MkString("TheRockTrading") ,
            "countries":MkArray(&VarArray{
                MkString("MT") ,
                }),
            "rateLimit":MkInteger(1000) ,
            "version":MkString("v1") ,
            "has":MkMap(&VarMap{
                "cancelOrder":MkBool(true) ,
                "CORS":MkBool(false) ,
                "createOrder":MkBool(true) ,
                "fetchBalance":MkBool(true) ,
                "fetchClosedOrders":MkBool(true) ,
                "fetchDeposits":MkBool(true) ,
                "fetchLedger":MkBool(true) ,
                "fetchMarkets":MkBool(true) ,
                "fetchMyTrades":MkBool(true) ,
                "fetchOpenOrders":MkBool(true) ,
                "fetchOrder":MkBool(true) ,
                "fetchOrderBook":MkBool(true) ,
                "fetchOrders":MkBool(true) ,
                "fetchTicker":MkBool(true) ,
                "fetchTickers":MkBool(true) ,
                "fetchTrades":MkBool(true) ,
                "fetchTransactions":MkString("emulated") ,
                "fetchWithdrawals":MkBool(true) ,
                }),
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/1294454/27766869-75057fa2-5ee9-11e7-9a6f-13e641fa4707.jpg") ,
                "api":MkString("https://api.therocktrading.com") ,
                "www":MkString("https://therocktrading.com") ,
                "doc":MkArray(&VarArray{
                    MkString("https://api.therocktrading.com/doc/v1/index.html") ,
                    MkString("https://api.therocktrading.com/doc/") ,
                    }),
                }),
            "api":MkMap(&VarMap{
                "public":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("funds") ,
                        MkString("funds/{id}/orderbook") ,
                        MkString("funds/{id}/ticker") ,
                        MkString("funds/{id}/trades") ,
                        MkString("funds/tickers") ,
                        })}),
                "private":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("balances") ,
                        MkString("balances/{id}") ,
                        MkString("discounts") ,
                        MkString("discounts/{id}") ,
                        MkString("funds") ,
                        MkString("funds/{id}") ,
                        MkString("funds/{id}/trades") ,
                        MkString("funds/{fund_id}/orders") ,
                        MkString("funds/{fund_id}/orders/{id}") ,
                        MkString("funds/{fund_id}/position_balances") ,
                        MkString("funds/{fund_id}/positions") ,
                        MkString("funds/{fund_id}/positions/{id}") ,
                        MkString("transactions") ,
                        MkString("transactions/{id}") ,
                        MkString("withdraw_limits/{id}") ,
                        MkString("withdraw_limits") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("atms/withdraw") ,
                        MkString("funds/{fund_id}/orders") ,
                        }),
                    "delete":MkArray(&VarArray{
                        MkString("funds/{fund_id}/orders/{id}") ,
                        MkString("funds/{fund_id}/orders/remove_all") ,
                        }),
                    }),
                }),
            "fees":MkMap(&VarMap{
                "trading":MkMap(&VarMap{
                    "maker":this.ParseNumber(MkString("0.002") ),
                    "taker":this.ParseNumber(MkString("0.002") ),
                    }),
                "funding":MkMap(&VarMap{
                    "tierBased":MkBool(false) ,
                    "percentage":MkBool(false) ,
                    "withdraw":MkMap(&VarMap{}),
                    "deposit":MkMap(&VarMap{
                        "BTC":MkInteger(0) ,
                        "BCH":MkInteger(0) ,
                        "PPC":MkInteger(0) ,
                        "ETH":MkInteger(0) ,
                        "ZEC":MkInteger(0) ,
                        "LTC":MkInteger(0) ,
                        "EUR":MkInteger(0) ,
                        }),
                    }),
                }),
            "exceptions":MkMap(&VarMap{
                "exact":MkMap(&VarMap{
                    "Request already running":BadRequest ,
                    "cannot specify multiple address types":BadRequest ,
                    "Currency is not included in the list":BadRequest ,
                    "Record not found":OrderNotFound ,
                    }),
                "broad":MkMap(&VarMap{
                    "before must be greater than after param":BadRequest ,
                    "must be shorter than 60 days":BadRequest ,
                    "must be a multiple of (period param) in minutes":BadRequest ,
                    "Address allocation limit reached for currency":InvalidAddress ,
                    "is not a valid value for param currency":BadRequest ,
                    " is invalid":InvalidAddress ,
                    }),
                }),
            }));
}

func (this *Therock) FetchMarkets(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetFunds"), params ); _ = response;
  markets := this.SafeValue(response , MkString("funds") ); _ = markets;
  result := MkArray(&VarArray{}); _ = result;
  if IsTrue(OpEq2(markets , MkUndefined() )) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , MkString(" fetchMarkets got an unexpected response") )));
  } else {
    for i := MkInteger(0) ; IsTrue(OpLw(i , markets.Length )); OpIncr(& i ){
      market := *(markets ).At(i ); _ = market;
      id := this.SafeString(market , MkString("id") ); _ = id;
      baseId := this.SafeString(market , MkString("trade_currency") ); _ = baseId;
      quoteId := this.SafeString(market , MkString("base_currency") ); _ = quoteId;
      base := this.SafeCurrencyCode(baseId ); _ = base;
      quote := this.SafeCurrencyCode(quoteId ); _ = quote;
      symbol := OpAdd(base , OpAdd(MkString("/") , quote )); _ = symbol;
      buy_fee := this.SafeNumber(market , MkString("buy_fee") ); _ = buy_fee;
      sell_fee := this.SafeNumber(market , MkString("sell_fee") ); _ = sell_fee;
      taker := Math.Max(buy_fee , sell_fee ); _ = taker;
      taker = OpDiv(taker , MkInteger(100) );
      maker := OpCopy(taker ); _ = maker;
      result.Push(MkMap(&VarMap{
              "id":id ,
              "symbol":symbol ,
              "base":base ,
              "quote":quote ,
              "baseId":baseId ,
              "quoteId":quoteId ,
              "info":market ,
              "active":MkBool(true) ,
              "maker":maker ,
              "taker":taker ,
              "precision":MkMap(&VarMap{
                  "amount":this.SafeInteger(market , MkString("trade_currency_decimals") ),
                  "price":this.SafeInteger(market , MkString("base_currency_decimals") ),
                  }),
              "limits":MkMap(&VarMap{
                  "amount":MkMap(&VarMap{
                      "min":this.SafeNumber(market , MkString("minimum_quantity_offer") ),
                      "max":MkUndefined() ,
                      }),
                  "price":MkMap(&VarMap{
                      "min":this.SafeNumber(market , MkString("minimum_price_offer") ),
                      "max":MkUndefined() ,
                      }),
                  "cost":MkMap(&VarMap{
                      "min":MkUndefined() ,
                      "max":MkUndefined() ,
                      }),
                  }),
              }));
    }
  }
  return result ;
}

func (this *Therock) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("privateGetBalances"), params ); _ = response;
  balances := this.SafeValue(response , MkString("balances") , MkArray(&VarArray{})); _ = balances;
  result := MkMap(&VarMap{"info":response }); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , balances.Length )); OpIncr(& i ){
    balance := *(balances ).At(i ); _ = balance;
    currencyId := this.SafeString(balance , MkString("currency") ); _ = currencyId;
    code := this.SafeCurrencyCode(currencyId ); _ = code;
    account := this.Account(); _ = account;
    *(account ).At(MkString("free") )= this.SafeString(balance , MkString("trading_balance") );
    *(account ).At(MkString("total") )= this.SafeString(balance , MkString("balance") );
    *(result ).At(code )= OpCopy(account );
  }
  return this.ParseBalance(result );
}

func (this *Therock) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"id":this.MarketId(symbol )}); _ = request;
  orderbook := this.Call(MkString("publicGetFundsIdOrderbook"), this.Extend(request , params )); _ = orderbook;
  timestamp := this.Parse8601(this.SafeString(orderbook , MkString("date") )); _ = timestamp;
  return this.ParseOrderBook(orderbook , symbol , timestamp , MkString("bids") , MkString("asks") , MkString("price") , MkString("amount") );
}

func (this *Therock) ParseTicker(goArgs ...*Variant) *Variant{
  ticker := GoGetArg(goArgs, 0, MkUndefined()); _ = ticker;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.Parse8601(*(ticker ).At(MkString("date") )); _ = timestamp;
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
        "bid":this.SafeNumber(ticker , MkString("bid") ),
        "bidVolume":MkUndefined() ,
        "ask":this.SafeNumber(ticker , MkString("ask") ),
        "askVolume":MkUndefined() ,
        "vwap":MkUndefined() ,
        "open":this.SafeNumber(ticker , MkString("open") ),
        "close":last ,
        "last":last ,
        "previousClose":this.SafeNumber(ticker , MkString("close") ),
        "change":MkUndefined() ,
        "percentage":MkUndefined() ,
        "average":MkUndefined() ,
        "baseVolume":this.SafeNumber(ticker , MkString("volume_traded") ),
        "quoteVolume":this.SafeNumber(ticker , MkString("volume") ),
        "info":ticker ,
        });
}

func (this *Therock) FetchTickers(goArgs ...*Variant) *Variant{
  symbols := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbols;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("publicGetFundsTickers"), params ); _ = response;
  tickers := this.IndexBy(*(response ).At(MkString("tickers") ), MkString("fund_id") ); _ = tickers;
  ids := GoGetKeys(tickers ); _ = ids;
  result := MkMap(&VarMap{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , ids.Length )); OpIncr(& i ){
    id := *(ids ).At(i ); _ = id;
    market := this.SafeMarket(id ); _ = market;
    symbol := *(market ).At(MkString("symbol") ); _ = symbol;
    ticker := *(tickers ).At(id ); _ = ticker;
    *(result ).At(symbol )= this.ParseTicker(ticker , market );
  }
  return this.FilterByArray(result , MkString("symbol") , symbols );
}

func (this *Therock) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  ticker := this.Call(MkString("publicGetFundsIdTicker"), this.Extend(MkMap(&VarMap{"id":*(market ).At(MkString("id") )}), params )); _ = ticker;
  return this.ParseTicker(ticker , market );
}

func (this *Therock) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  marketId := this.SafeString(trade , MkString("fund_id") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market ); _ = symbol;
  timestamp := this.Parse8601(this.SafeString(trade , MkString("date") )); _ = timestamp;
  id := this.SafeString(trade , MkString("id") ); _ = id;
  orderId := this.SafeString(trade , MkString("order_id") ); _ = orderId;
  side := this.SafeString(trade , MkString("side") ); _ = side;
  priceString := this.SafeString(trade , MkString("price") ); _ = priceString;
  amountString := this.SafeString(trade , MkString("amount") ); _ = amountString;
  price := this.ParseNumber(priceString ); _ = price;
  amount := this.ParseNumber(amountString ); _ = amount;
  cost := this.ParseNumber(Precise.StringMul(priceString , amountString )); _ = cost;
  fee := OpCopy(MkUndefined() ); _ = fee;
  feeCost := OpCopy(MkUndefined() ); _ = feeCost;
  transactions := this.SafeValue(trade , MkString("transactions") , MkArray(&VarArray{})); _ = transactions;
  transactionsByType := this.GroupBy(transactions , MkString("type") ); _ = transactionsByType;
  feeTransactions := this.SafeValue(transactionsByType , MkString("paid_commission") , MkArray(&VarArray{})); _ = feeTransactions;
  for i := MkInteger(0) ; IsTrue(OpLw(i , feeTransactions.Length )); OpIncr(& i ){
    if IsTrue(OpEq2(feeCost , MkUndefined() )) {
      feeCost = MkInteger(0) ;
    }
    feeCost = this.Sum(feeCost , this.SafeNumber(*(feeTransactions ).At(i ), MkString("price") ));
  }
  if IsTrue(OpNotEq2(feeCost , MkUndefined() )) {
    fee = MkMap(&VarMap{
        "cost":feeCost ,
        "currency":*(market ).At(MkString("quote") ),
        });
  }
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

func (this *Therock) ParseLedgerEntryDirection(goArgs ...*Variant) *Variant{
  direction := GoGetArg(goArgs, 0, MkUndefined()); _ = direction;
  directions := MkMap(&VarMap{
        "affiliate_earnings":MkString("in") ,
        "atm_payment":MkString("in") ,
        "bought_currency_from_fund":MkString("out") ,
        "bought_shares":MkString("out") ,
        "paid_commission":MkString("out") ,
        "paypal_payment":MkString("in") ,
        "pos_payment":MkString("in") ,
        "released_currency_to_fund":MkString("out") ,
        "rollover_commission":MkString("out") ,
        "sold_currency_to_fund":MkString("in") ,
        "sold_shares":MkString("in") ,
        "transfer_received":MkString("in") ,
        "transfer_sent":MkString("out") ,
        "withdraw":MkString("out") ,
        }); _ = directions;
  return this.SafeString(directions , direction , direction );
}

func (this *Therock) ParseLedgerEntryType(goArgs ...*Variant) *Variant{
  type_ := GoGetArg(goArgs, 0, MkUndefined()); _ = type_;
  types := MkMap(&VarMap{
        "affiliate_earnings":MkString("referral") ,
        "atm_payment":MkString("transaction") ,
        "bought_currency_from_fund":MkString("trade") ,
        "bought_shares":MkString("trade") ,
        "paid_commission":MkString("fee") ,
        "paypal_payment":MkString("transaction") ,
        "pos_payment":MkString("transaction") ,
        "released_currency_to_fund":MkString("trade") ,
        "rollover_commission":MkString("fee") ,
        "sold_currency_to_fund":MkString("trade") ,
        "sold_shares":MkString("trade") ,
        "transfer_received":MkString("transfer") ,
        "transfer_sent":MkString("transfer") ,
        "withdraw":MkString("transaction") ,
        }); _ = types;
  return this.SafeString(types , type_ , type_ );
}

func (this *Therock) ParseLedgerEntry(goArgs ...*Variant) *Variant{
  item := GoGetArg(goArgs, 0, MkUndefined()); _ = item;
  currency := GoGetArg(goArgs, 1, MkUndefined() ); _ = currency;
  id := this.SafeString(item , MkString("id") ); _ = id;
  referenceId := OpCopy(MkUndefined() ); _ = referenceId;
  type_ := this.SafeString(item , MkString("type") ); _ = type_;
  direction := this.ParseLedgerEntryDirection(type_ ); _ = direction;
  type_ = this.ParseLedgerEntryType(type_ );
  if IsTrue(OpOr(OpEq2(type_ , MkString("trade") ), OpEq2(type_ , MkString("fee") ))) {
    referenceId = this.SafeString(item , MkString("trade_id") );
  }
  currencyId := this.SafeString(item , MkString("currency") ); _ = currencyId;
  code := this.SafeCurrencyCode(currencyId ); _ = code;
  amount := this.SafeNumber(item , MkString("price") ); _ = amount;
  timestamp := this.Parse8601(this.SafeString(item , MkString("date") )); _ = timestamp;
  status := MkString("ok") ; _ = status;
  return MkMap(&VarMap{
        "info":item ,
        "id":id ,
        "direction":direction ,
        "account":MkUndefined() ,
        "referenceId":referenceId ,
        "referenceAccount":MkUndefined() ,
        "type":type_ ,
        "currency":code ,
        "amount":amount ,
        "before":MkUndefined() ,
        "after":MkUndefined() ,
        "status":status ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "fee":MkUndefined() ,
        });
}

func (this *Therock) FetchLedger(goArgs ...*Variant) *Variant{
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
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("after") )= this.Iso8601(since );
  }
  response := this.Call(MkString("privateGetTransactions"), this.Extend(request , params )); _ = response;
  transactions := this.SafeValue(response , MkString("transactions") , MkArray(&VarArray{})); _ = transactions;
  return this.ParseLedger(transactions , currency , since , limit );
}

func (this *Therock) ParseTransactionType(goArgs ...*Variant) *Variant{
  type_ := GoGetArg(goArgs, 0, MkUndefined()); _ = type_;
  types := MkMap(&VarMap{
        "withdraw":MkString("withdrawal") ,
        "atm_payment":MkString("deposit") ,
        }); _ = types;
  return this.SafeString(types , type_ , type_ );
}

func (this *Therock) ParseTransaction(goArgs ...*Variant) *Variant{
  transaction := GoGetArg(goArgs, 0, MkUndefined()); _ = transaction;
  currency := GoGetArg(goArgs, 1, MkUndefined() ); _ = currency;
  id := this.SafeString(transaction , MkString("id") ); _ = id;
  type_ := this.ParseTransactionType(this.SafeString(transaction , MkString("type") )); _ = type_;
  detail := this.SafeValue(transaction , MkString("transfer_detail") , MkMap(&VarMap{})); _ = detail;
  method := this.SafeString(detail , MkString("method") ); _ = method;
  txid := OpCopy(MkUndefined() ); _ = txid;
  address := OpCopy(MkUndefined() ); _ = address;
  if IsTrue(OpNotEq2(method , MkUndefined() )) {
    if IsTrue(OpNotEq2(method , MkString("wire_transfer") )) {
      txid = this.SafeString(detail , MkString("id") );
      address = this.SafeString(detail , MkString("recipient") );
    }
  }
  currencyId := this.SafeString(transaction , MkString("currency") ); _ = currencyId;
  code := this.SafeCurrencyCode(currencyId ); _ = code;
  amount := this.SafeNumber(transaction , MkString("price") ); _ = amount;
  timestamp := this.Parse8601(this.SafeString(transaction , MkString("date") )); _ = timestamp;
  status := MkString("ok") ; _ = status;
  return MkMap(&VarMap{
        "info":transaction ,
        "id":id ,
        "currency":code ,
        "amount":amount ,
        "addressFrom":MkUndefined() ,
        "addressTo":address ,
        "address":address ,
        "tagFrom":MkUndefined() ,
        "tagTo":MkUndefined() ,
        "tag":MkUndefined() ,
        "status":status ,
        "type":type_ ,
        "updated":MkUndefined() ,
        "txid":txid ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "fee":MkUndefined() ,
        });
}

func (this *Therock) FetchWithdrawals(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"type":MkString("withdraw") }); _ = request;
  return this.FetchTransactions(code , since , limit , this.Extend(request , params ));
}

func (this *Therock) FetchDeposits(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"type":MkString("atm_payment") }); _ = request;
  return this.FetchTransactions(code , since , limit , this.Extend(request , params ));
}

func (this *Therock) FetchTransactions(goArgs ...*Variant) *Variant{
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
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("after") )= this.Iso8601(since );
  }
  params = this.Extend(request , params );
  response := this.Call(MkString("privateGetTransactions"), params ); _ = response;
  transactions := this.SafeValue(response , MkString("transactions") , MkArray(&VarArray{})); _ = transactions;
  transactionTypes := MkArray(&VarArray{
        MkString("withdraw") ,
        MkString("atm_payment") ,
        }); _ = transactionTypes;
  depositsAndWithdrawals := this.FilterByArray(transactions , MkString("type") , transactionTypes , MkBool(false) ); _ = depositsAndWithdrawals;
  return this.ParseTransactions(depositsAndWithdrawals , currency , since , limit );
}

func (this *Therock) ParseOrderStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "active":MkString("open") ,
        "executed":MkString("closed") ,
        "deleted":MkString("canceled") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Therock) ParseOrder(goArgs ...*Variant) *Variant{
  order := GoGetArg(goArgs, 0, MkUndefined()); _ = order;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  id := this.SafeString(order , MkString("id") ); _ = id;
  marketId := this.SafeString(order , MkString("fund_id") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market ); _ = symbol;
  status := this.ParseOrderStatus(this.SafeString(order , MkString("status") )); _ = status;
  timestamp := this.Parse8601(this.SafeString(order , MkString("date") )); _ = timestamp;
  type_ := this.SafeString(order , MkString("type") ); _ = type_;
  side := this.SafeString(order , MkString("side") ); _ = side;
  amount := this.SafeNumber(order , MkString("amount") ); _ = amount;
  remaining := this.SafeNumber(order , MkString("amount_unfilled") ); _ = remaining;
  filled := OpCopy(MkUndefined() ); _ = filled;
  if IsTrue(OpNotEq2(amount , MkUndefined() )) {
    if IsTrue(OpNotEq2(remaining , MkUndefined() )) {
      filled = OpSub(amount , remaining );
    }
  }
  price := this.SafeNumber(order , MkString("price") ); _ = price;
  trades := this.SafeValue(order , MkString("trades") ); _ = trades;
  cost := OpCopy(MkUndefined() ); _ = cost;
  average := OpCopy(MkUndefined() ); _ = average;
  lastTradeTimestamp := OpCopy(MkUndefined() ); _ = lastTradeTimestamp;
  if IsTrue(OpNotEq2(trades , MkUndefined() )) {
    numTrades := OpCopy(trades.Length ); _ = numTrades;
    if IsTrue(OpGt(numTrades , MkInteger(0) )) {
      trades = this.ParseTrades(trades , market , MkUndefined() , MkUndefined() , MkMap(&VarMap{"orderId":id }));
      cost = MkInteger(0) ;
      filled = MkInteger(0) ;
      for i := MkInteger(0) ; IsTrue(OpLw(i , numTrades )); OpIncr(& i ){
        trade := *(trades ).At(i ); _ = trade;
        cost = this.Sum(cost , *(trade ).At(MkString("cost") ));
        filled = this.Sum(filled , *(trade ).At(MkString("amount") ));
      }
      if IsTrue(OpGt(filled , MkInteger(0) )) {
        average = OpDiv(cost , filled );
      }
      lastTradeTimestamp = *(*(trades ).At(OpSub(numTrades , MkInteger(1) ))).At(MkString("timestamp") );
    } else {
      cost = MkInteger(0) ;
    }
  }
  stopPrice := this.SafeNumber(order , MkString("conditional_price") ); _ = stopPrice;
  return MkMap(&VarMap{
        "id":id ,
        "clientOrderId":MkUndefined() ,
        "info":order ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "lastTradeTimestamp":lastTradeTimestamp ,
        "status":status ,
        "symbol":symbol ,
        "type":type_ ,
        "timeInForce":MkUndefined() ,
        "postOnly":MkUndefined() ,
        "side":side ,
        "price":price ,
        "stopPrice":stopPrice ,
        "cost":cost ,
        "amount":amount ,
        "filled":filled ,
        "average":average ,
        "remaining":remaining ,
        "fee":MkUndefined() ,
        "trades":trades ,
        });
}

func (this *Therock) FetchOpenOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"status":MkString("active") }); _ = request;
  return this.FetchOrders(symbol , since , limit , this.Extend(request , params ));
}

func (this *Therock) FetchClosedOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"status":MkString("executed") }); _ = request;
  return this.FetchOrders(symbol , since , limit , this.Extend(request , params ));
}

func (this *Therock) FetchOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchOrders() requires a symbol argument") )));
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"fund_id":*(market ).At(MkString("id") )}); _ = request;
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("after") )= this.Iso8601(since );
  }
  response := this.Call(MkString("privateGetFundsFundIdOrders"), this.Extend(request , params )); _ = response;
  orders := this.SafeValue(response , MkString("orders") , MkArray(&VarArray{})); _ = orders;
  return this.ParseOrders(orders , market , since , limit );
}

func (this *Therock) FetchOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchOrder() requires a symbol argument") )));
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "id":id ,
        "fund_id":*(market ).At(MkString("id") ),
        }); _ = request;
  response := this.Call(MkString("privatePostFundsFundIdOrdersId"), this.Extend(request , params )); _ = response;
  return this.ParseOrder(response );
}

func (this *Therock) CreateOrder(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 2, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 3, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 4, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 5, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  if IsTrue(OpEq2(type_ , MkString("market") )) {
    price = MkInteger(0) ;
  }
  request := MkMap(&VarMap{
        "fund_id":this.MarketId(symbol ),
        "side":side ,
        "amount":amount ,
        "price":price ,
        }); _ = request;
  response := this.Call(MkString("privatePostFundsFundIdOrders"), this.Extend(request , params )); _ = response;
  return this.ParseOrder(response );
}

func (this *Therock) CancelOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{
        "id":id ,
        "fund_id":this.MarketId(symbol ),
        }); _ = request;
  response := this.Call(MkString("privateDeleteFundsFundIdOrdersId"), this.Extend(request , params )); _ = response;
  return this.ParseOrder(response );
}

func (this *Therock) FetchMyTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchMyTrades() requires a symbol argument") )));
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"id":*(market ).At(MkString("id") )}); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("per_page") )= OpCopy(limit );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("after") )= this.Iso8601(since );
  }
  response := this.Call(MkString("privateGetFundsIdTrades"), this.Extend(request , params )); _ = response;
  return this.ParseTrades(*(response ).At(MkString("trades") ), market , since , limit );
}

func (this *Therock) FetchTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"id":*(market ).At(MkString("id") )}); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("per_page") )= OpCopy(limit );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("after") )= this.Iso8601(since );
  }
  response := this.Call(MkString("publicGetFundsIdTrades"), this.Extend(request , params )); _ = response;
  return this.ParseTrades(*(response ).At(MkString("trades") ), market , since , limit );
}

func (this *Therock) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("public") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  url := OpAdd(*(*this.At(MkString("urls")) ).At(MkString("api") ), OpAdd(MkString("/") , OpAdd(*this.At(MkString("version")) , OpAdd(MkString("/") , this.ImplodeParams(path , params ))))); _ = url;
  query := this.Omit(params , this.ExtractParams(path )); _ = query;
  headers = OpTrinary(OpEq2(headers , MkUndefined() ), MkMap(&VarMap{}), headers );
  if IsTrue(OpEq2(api , MkString("private") )) {
    this.CheckRequiredCredentials();
    if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
      if IsTrue(OpEq2(method , MkString("POST") )) {
        body = this.Json(query );
        *(headers ).At(MkString("Content-Type") )= MkString("application/json") ;
      } else {
        queryString := this.Rawencode(query ); _ = queryString;
        if IsTrue(queryString.Length ) {
          OpAddAssign(& url , OpAdd(MkString("?") , queryString ));
        }
      }
    }
    nonce := (this.Nonce()).Call(MkString("toString") ); _ = nonce;
    auth := OpAdd(nonce , url ); _ = auth;
    *(headers ).At(MkString("X-TRT-KEY") )= OpCopy(*this.At(MkString("apiKey")) );
    *(headers ).At(MkString("X-TRT-NONCE") )= OpCopy(nonce );
    *(headers ).At(MkString("X-TRT-SIGN") )= this.Hmac(this.Encode(auth ), this.Encode(*this.At(MkString("secret")) ), MkString("sha512") );
  } else {
    if IsTrue(OpEq2(api , MkString("public") )) {
      if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
        OpAddAssign(& url , OpAdd(MkString("?") , this.Rawencode(query )));
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

func (this *Therock) HandleErrors(goArgs ...*Variant) *Variant{
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
  errors := this.SafeValue(response , MkString("errors") , MkArray(&VarArray{})); _ = errors;
  numErrors := OpCopy(errors.Length ); _ = numErrors;
  if IsTrue(OpGt(numErrors , MkInteger(0) )) {
    feedback := OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , body )); _ = feedback;
    for i := MkInteger(0) ; IsTrue(OpLw(i , numErrors )); OpIncr(& i ){
      error := *(errors ).At(i ); _ = error;
      message := this.SafeString(error , MkString("message") ); _ = message;
      this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("exact") ), message , feedback );
      this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("broad") ), message , feedback );
    }
    panic(NewExchangeError(feedback ));
  }
  return MkUndefined()
}

