package ccxt

import (
)

type Tidex struct {
	*ExchangeBase
}

var _ Exchange = (*Tidex)(nil)

func init() {
	exchange := &Tidex{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Tidex) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("tidex") ,
            "name":MkString("Tidex") ,
            "countries":MkArray(&VarArray{
                MkString("UK") ,
                }),
            "rateLimit":MkInteger(2000) ,
            "version":MkString("3") ,
            "userAgent":*(*this.At(MkString("userAgents")) ).At(MkString("chrome") ),
            "has":MkMap(&VarMap{
                "cancelOrder":MkBool(true) ,
                "CORS":MkBool(false) ,
                "createMarketOrder":MkBool(false) ,
                "createOrder":MkBool(true) ,
                "fetchBalance":MkBool(true) ,
                "fetchCurrencies":MkBool(true) ,
                "fetchMarkets":MkBool(true) ,
                "fetchMyTrades":MkBool(true) ,
                "fetchOpenOrders":MkBool(true) ,
                "fetchOrder":MkBool(true) ,
                "fetchOrderBook":MkBool(true) ,
                "fetchOrderBooks":MkBool(true) ,
                "fetchTicker":MkBool(true) ,
                "fetchTickers":MkBool(true) ,
                "fetchTrades":MkBool(true) ,
                "withdraw":MkBool(true) ,
                }),
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/1294454/30781780-03149dc4-a12e-11e7-82bb-313b269d24d4.jpg") ,
                "api":MkMap(&VarMap{
                    "web":MkString("https://gate.tidex.com/api") ,
                    "public":MkString("https://api.tidex.com/api/3") ,
                    "private":MkString("https://api.tidex.com/tapi") ,
                    }),
                "www":MkString("https://tidex.com") ,
                "doc":MkString("https://tidex.com/exchange/public-api") ,
                "referral":MkString("https://tidex.com/exchange/?ref=57f5638d9cd7") ,
                "fees":MkArray(&VarArray{
                    MkString("https://tidex.com/exchange/assets-spec") ,
                    MkString("https://tidex.com/exchange/pairs-spec") ,
                    }),
                }),
            "api":MkMap(&VarMap{
                "web":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("currency") ,
                        MkString("pairs") ,
                        MkString("tickers") ,
                        MkString("orders") ,
                        MkString("ordershistory") ,
                        MkString("trade-data") ,
                        MkString("trade-data/{id}") ,
                        })}),
                "public":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("info") ,
                        MkString("ticker/{pair}") ,
                        MkString("depth/{pair}") ,
                        MkString("trades/{pair}") ,
                        })}),
                "private":MkMap(&VarMap{"post":MkArray(&VarArray{
                        MkString("getInfoExt") ,
                        MkString("getInfo") ,
                        MkString("Trade") ,
                        MkString("ActiveOrders") ,
                        MkString("OrderInfo") ,
                        MkString("CancelOrder") ,
                        MkString("TradeHistory") ,
                        MkString("CoinDepositAddress") ,
                        MkString("WithdrawCoin") ,
                        MkString("CreateCoupon") ,
                        MkString("RedeemCoupon") ,
                        })}),
                }),
            "fees":MkMap(&VarMap{"trading":MkMap(&VarMap{
                    "feeSide":MkString("get") ,
                    "tierBased":MkBool(false) ,
                    "percentage":MkBool(true) ,
                    "taker":this.ParseNumber(MkString("0.001") ),
                    "maker":this.ParseNumber(MkString("0.001") ),
                    })}),
            "commonCurrencies":MkMap(&VarMap{
                "DSH":MkString("DASH") ,
                "EMGO":MkString("MGO") ,
                "MGO":MkString("WMGO") ,
                }),
            "exceptions":MkMap(&VarMap{
                "exact":MkMap(&VarMap{
                    "803":InvalidOrder ,
                    "804":InvalidOrder ,
                    "805":InvalidOrder ,
                    "806":InvalidOrder ,
                    "807":InvalidOrder ,
                    "831":InsufficientFunds ,
                    "832":InsufficientFunds ,
                    "833":OrderNotFound ,
                    }),
                "broad":MkMap(&VarMap{
                    "Invalid pair name":ExchangeError ,
                    "invalid api key":AuthenticationError ,
                    "invalid sign":AuthenticationError ,
                    "api key dont have trade permission":AuthenticationError ,
                    "invalid parameter":InvalidOrder ,
                    "invalid order":InvalidOrder ,
                    "Requests too often":DDoSProtection ,
                    "not available":ExchangeNotAvailable ,
                    "data unavailable":ExchangeNotAvailable ,
                    "external service unavailable":ExchangeNotAvailable ,
                    }),
                }),
            "options":MkMap(&VarMap{"fetchTickersMaxLength":MkInteger(2048) }),
            "orders":MkMap(&VarMap{}),
            }));
}

func (this *Tidex) FetchCurrencies(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("webGetCurrency"), params ); _ = response;
  result := MkMap(&VarMap{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , response.Length )); OpIncr(& i ){
    currency := *(response ).At(i ); _ = currency;
    id := this.SafeString(currency , MkString("symbol") ); _ = id;
    precision := this.SafeInteger(currency , MkString("amountPoint") ); _ = precision;
    code := this.SafeCurrencyCode(id ); _ = code;
    visible := this.SafeValue(currency , MkString("visible") ); _ = visible;
    active := OpEq2(visible , MkBool(true) ); _ = active;
    withdrawEnable := this.SafeValue(currency , MkString("withdrawEnable") ); _ = withdrawEnable;
    depositEnable := this.SafeValue(currency , MkString("depositEnable") ); _ = depositEnable;
    canWithdraw := OpEq2(withdrawEnable , MkBool(true) ); _ = canWithdraw;
    canDeposit := OpEq2(depositEnable , MkBool(true) ); _ = canDeposit;
    if IsTrue(OpOr(OpNot(canWithdraw ), OpNot(canDeposit ))) {
      active = OpCopy(MkBool(false) );
    }
    name := this.SafeString(currency , MkString("name") ); _ = name;
    fee := this.SafeNumber(currency , MkString("withdrawFee") ); _ = fee;
    *(result ).At(code )= MkMap(&VarMap{
        "id":id ,
        "code":code ,
        "name":name ,
        "active":active ,
        "precision":precision ,
        "funding":MkMap(&VarMap{
            "withdraw":MkMap(&VarMap{
                "active":canWithdraw ,
                "fee":fee ,
                }),
            "deposit":MkMap(&VarMap{
                "active":canDeposit ,
                "fee":this.ParseNumber(MkString("0") ),
                }),
            }),
        "limits":MkMap(&VarMap{
            "amount":MkMap(&VarMap{
                "min":MkUndefined() ,
                "max":MkUndefined() ,
                }),
            "withdraw":MkMap(&VarMap{
                "min":this.SafeNumber(currency , MkString("withdrawMinAmount") ),
                "max":MkUndefined() ,
                }),
            "deposit":MkMap(&VarMap{
                "min":this.SafeNumber(currency , MkString("depositMinAmount") ),
                "max":MkUndefined() ,
                }),
            }),
        "info":currency ,
        });
  }
  return result ;
}

func (this *Tidex) FetchMarkets(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetInfo"), params ); _ = response;
  markets := *(response ).At(MkString("pairs") ); _ = markets;
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
    precision := MkMap(&VarMap{
          "amount":this.SafeInteger(market , MkString("decimal_places") ),
          "price":this.SafeInteger(market , MkString("decimal_places") ),
          }); _ = precision;
    limits := MkMap(&VarMap{
          "amount":MkMap(&VarMap{
              "min":this.SafeNumber(market , MkString("min_amount") ),
              "max":this.SafeNumber(market , MkString("max_amount") ),
              }),
          "price":MkMap(&VarMap{
              "min":this.SafeNumber(market , MkString("min_price") ),
              "max":this.SafeNumber(market , MkString("max_price") ),
              }),
          "cost":MkMap(&VarMap{"min":this.SafeNumber(market , MkString("min_total") )}),
          }); _ = limits;
    hidden := this.SafeInteger(market , MkString("hidden") ); _ = hidden;
    active := OpEq2(hidden , MkInteger(0) ); _ = active;
    takerFeeString := this.SafeString(market , MkString("fee") ); _ = takerFeeString;
    takerFeeString = Precise.StringDiv(takerFeeString , MkString("100") );
    takerFee := this.ParseNumber(takerFeeString ); _ = takerFee;
    result.Push(MkMap(&VarMap{
            "id":id ,
            "symbol":symbol ,
            "base":base ,
            "quote":quote ,
            "baseId":baseId ,
            "quoteId":quoteId ,
            "active":active ,
            "taker":takerFee ,
            "precision":precision ,
            "limits":limits ,
            "info":market ,
            }));
  }
  return result ;
}

func (this *Tidex) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("privatePostGetInfoExt"), params ); _ = response;
  balances := this.SafeValue(response , MkString("return") ); _ = balances;
  timestamp := this.SafeTimestamp(balances , MkString("server_time") ); _ = timestamp;
  result := MkMap(&VarMap{
        "info":response ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        }); _ = result;
  funds := this.SafeValue(balances , MkString("funds") , MkMap(&VarMap{})); _ = funds;
  currencyIds := GoGetKeys(funds ); _ = currencyIds;
  for i := MkInteger(0) ; IsTrue(OpLw(i , currencyIds.Length )); OpIncr(& i ){
    currencyId := *(currencyIds ).At(i ); _ = currencyId;
    code := this.SafeCurrencyCode(currencyId ); _ = code;
    balance := this.SafeValue(funds , currencyId , MkMap(&VarMap{})); _ = balance;
    account := this.Account(); _ = account;
    *(account ).At(MkString("free") )= this.SafeString(balance , MkString("value") );
    *(account ).At(MkString("used") )= this.SafeString(balance , MkString("inOrders") );
    *(result ).At(code )= OpCopy(account );
  }
  return this.ParseBalance(result );
}

func (this *Tidex) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"pair":*(market ).At(MkString("id") )}); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("publicGetDepthPair"), this.Extend(request , params )); _ = response;
  market_id_in_reponse := OpHasMember(*(market ).At(MkString("id") ), response ); _ = market_id_in_reponse;
  if IsTrue(OpNot(market_id_in_reponse )) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , OpAdd(*(market ).At(MkString("symbol") ), MkString(" order book is empty or not available") )))));
  }
  orderbook := *(response ).At(*(market ).At(MkString("id") )); _ = orderbook;
  return this.ParseOrderBook(orderbook , symbol );
}

func (this *Tidex) FetchOrderBooks(goArgs ...*Variant) *Variant{
  symbols := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbols;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  ids := OpCopy(MkUndefined() ); _ = ids;
  if IsTrue(OpEq2(symbols , MkUndefined() )) {
    ids = (*((this).At(MkString("ids")))).Call(MkString("join"), MkString("-") );
    if IsTrue(OpGt(ids.Length , MkInteger(2048) )) {
      numIds := OpCopy((*((this).At(MkString("ids")))).Length ); _ = numIds;
      panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" has ") , OpAdd(numIds.ToString(), MkString(" symbols exceeding max URL length, you are required to specify a list of symbols in the first argument to fetchOrderBooks") )))));
    }
  } else {
    ids = this.MarketIds(symbols );
    ids = ids.Join(MkString("-") );
  }
  request := MkMap(&VarMap{"pair":ids }); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("publicGetDepthPair"), this.Extend(request , params )); _ = response;
  result := MkMap(&VarMap{}); _ = result;
  ids = GoGetKeys(response );
  for i := MkInteger(0) ; IsTrue(OpLw(i , ids.Length )); OpIncr(& i ){
    id := *(ids ).At(i ); _ = id;
    symbol := this.SafeSymbol(id ); _ = symbol;
    *(result ).At(symbol )= this.ParseOrderBook(*(response ).At(id ));
  }
  return result ;
}

func (this *Tidex) ParseTicker(goArgs ...*Variant) *Variant{
  ticker := GoGetArg(goArgs, 0, MkUndefined()); _ = ticker;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.SafeTimestamp(ticker , MkString("updated") ); _ = timestamp;
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  if IsTrue(OpNotEq2(market , MkUndefined() )) {
    symbol = *(market ).At(MkString("symbol") );
    if IsTrue(OpNot(*(market ).At(MkString("active") ))) {
      timestamp = OpCopy(MkUndefined() );
    }
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
        "average":this.SafeNumber(ticker , MkString("avg") ),
        "baseVolume":this.SafeNumber(ticker , MkString("vol_cur") ),
        "quoteVolume":this.SafeNumber(ticker , MkString("vol") ),
        "info":ticker ,
        });
}

func (this *Tidex) FetchTickers(goArgs ...*Variant) *Variant{
  symbols := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbols;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  ids := OpCopy(*this.At(MkString("ids")) ); _ = ids;
  if IsTrue(OpEq2(symbols , MkUndefined() )) {
    numIds := OpCopy(ids.Length ); _ = numIds;
    ids = ids.Join(MkString("-") );
    if IsTrue(OpGt(ids.Length , *(*this.At(MkString("options")) ).At(MkString("fetchTickersMaxLength") ))) {
      maxLength := this.SafeInteger(*this.At(MkString("options")) , MkString("fetchTickersMaxLength") , MkInteger(2048) ); _ = maxLength;
      panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" has ") , OpAdd(numIds.ToString(), OpAdd(MkString(" markets exceeding max URL length for this endpoint (") , OpAdd(maxLength.ToString(), MkString(" characters), please, specify a list of symbols of interest in the first argument to fetchTickers") )))))));
    }
  } else {
    ids = this.MarketIds(symbols );
    ids = ids.Join(MkString("-") );
  }
  request := MkMap(&VarMap{"pair":ids }); _ = request;
  response := this.Call(MkString("publicGetTickerPair"), this.Extend(request , params )); _ = response;
  result := MkMap(&VarMap{}); _ = result;
  keys := GoGetKeys(response ); _ = keys;
  for i := MkInteger(0) ; IsTrue(OpLw(i , keys.Length )); OpIncr(& i ){
    id := *(keys ).At(i ); _ = id;
    market := this.SafeMarket(id ); _ = market;
    symbol := *(market ).At(MkString("symbol") ); _ = symbol;
    *(result ).At(symbol )= this.ParseTicker(*(response ).At(id ), market );
  }
  return this.FilterByArray(result , MkString("symbol") , symbols );
}

func (this *Tidex) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  tickers := this.FetchTickers(MkArray(&VarArray{
            symbol ,
            }), params ); _ = tickers;
  return *(tickers ).At(symbol );
}

func (this *Tidex) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.SafeTimestamp(trade , MkString("timestamp") ); _ = timestamp;
  side := this.SafeString(trade , MkString("type") ); _ = side;
  if IsTrue(OpEq2(side , MkString("ask") )) {
    side = MkString("sell") ;
  } else {
    if IsTrue(OpEq2(side , MkString("bid") )) {
      side = MkString("buy") ;
    }
  }
  priceString := this.SafeString2(trade , MkString("rate") , MkString("price") ); _ = priceString;
  id := this.SafeString2(trade , MkString("trade_id") , MkString("tid") ); _ = id;
  orderId := this.SafeString(trade , MkString("order_id") ); _ = orderId;
  marketId := this.SafeString(trade , MkString("pair") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market ); _ = symbol;
  amountString := this.SafeString(trade , MkString("amount") ); _ = amountString;
  price := this.ParseNumber(priceString ); _ = price;
  amount := this.ParseNumber(amountString ); _ = amount;
  cost := this.ParseNumber(Precise.StringMul(priceString , amountString )); _ = cost;
  type_ := MkString("limit") ; _ = type_;
  takerOrMaker := OpCopy(MkUndefined() ); _ = takerOrMaker;
  fee := OpCopy(MkUndefined() ); _ = fee;
  feeCost := this.SafeNumber(trade , MkString("commission") ); _ = feeCost;
  if IsTrue(OpNotEq2(feeCost , MkUndefined() )) {
    feeCurrencyId := this.SafeString(trade , MkString("commissionCurrency") ); _ = feeCurrencyId;
    feeCurrencyCode := this.SafeCurrencyCode(feeCurrencyId ); _ = feeCurrencyCode;
    fee = MkMap(&VarMap{
        "cost":feeCost ,
        "currency":feeCurrencyCode ,
        });
  }
  isYourOrder := this.SafeValue(trade , MkString("is_your_order") ); _ = isYourOrder;
  if IsTrue(OpNotEq2(isYourOrder , MkUndefined() )) {
    takerOrMaker = MkString("taker") ;
    if IsTrue(isYourOrder ) {
      takerOrMaker = MkString("maker") ;
    }
    if IsTrue(OpEq2(fee , MkUndefined() )) {
      fee = this.CalculateFee(symbol , type_ , side , amount , price , takerOrMaker );
    }
  }
  return MkMap(&VarMap{
        "id":id ,
        "order":orderId ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "symbol":symbol ,
        "type":type_ ,
        "side":side ,
        "takerOrMaker":takerOrMaker ,
        "price":price ,
        "amount":amount ,
        "cost":cost ,
        "fee":fee ,
        "info":trade ,
        });
}

func (this *Tidex) FetchTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"pair":*(market ).At(MkString("id") )}); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("publicGetTradesPair"), this.Extend(request , params )); _ = response;
  if IsTrue(GoIsArray(response )) {
    numElements := OpCopy(response.Length ); _ = numElements;
    if IsTrue(OpEq2(numElements , MkInteger(0) )) {
      return MkArray(&VarArray{});
    }
  }
  return this.ParseTrades(*(response ).At(*(market ).At(MkString("id") )), market , since , limit );
}

func (this *Tidex) CreateOrder(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 2, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 3, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 4, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 5, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(type_ , MkString("market") )) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , MkString(" allows limit orders only") )));
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "pair":*(market ).At(MkString("id") ),
        "type":side ,
        "amount":this.AmountToPrecision(symbol , amount ),
        "rate":this.PriceToPrecision(symbol , price ),
        }); _ = request;
  response := this.Call(MkString("privatePostTrade"), this.Extend(request , params )); _ = response;
  id := OpCopy(MkUndefined() ); _ = id;
  status := MkString("open") ; _ = status;
  filled := MkNumber(0.0) ; _ = filled;
  remaining := OpCopy(amount ); _ = remaining;
  returnResult := this.SafeValue(response , MkString("return") ); _ = returnResult;
  if IsTrue(OpNotEq2(returnResult , MkUndefined() )) {
    id = this.SafeString(returnResult , MkString("order_id") );
    if IsTrue(OpEq2(id , MkString("0") )) {
      id = this.SafeString(returnResult , MkString("init_order_id") );
      status = MkString("closed") ;
    }
    filled = this.SafeNumber(returnResult , MkString("received") , filled );
    remaining = this.SafeNumber(returnResult , MkString("remains") , amount );
  }
  timestamp := this.Milliseconds(); _ = timestamp;
  return this.SafeOrder(MkMap(&VarMap{
            "id":id ,
            "timestamp":timestamp ,
            "datetime":this.Iso8601(timestamp ),
            "lastTradeTimestamp":MkUndefined() ,
            "status":status ,
            "symbol":symbol ,
            "type":type_ ,
            "side":side ,
            "price":price ,
            "cost":MkUndefined() ,
            "amount":amount ,
            "remaining":remaining ,
            "filled":filled ,
            "fee":MkUndefined() ,
            "info":response ,
            "clientOrderId":MkUndefined() ,
            "average":MkUndefined() ,
            "trades":MkUndefined() ,
            }));
}

func (this *Tidex) CancelOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"order_id":parseInt(id )}); _ = request;
  return this.Call(MkString("privatePostCancelOrder"), this.Extend(request , params ));
}

func (this *Tidex) ParseOrderStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "0":MkString("open") ,
        "1":MkString("closed") ,
        "2":MkString("canceled") ,
        "3":MkString("canceled") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Tidex) ParseOrder(goArgs ...*Variant) *Variant{
  order := GoGetArg(goArgs, 0, MkUndefined()); _ = order;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  id := this.SafeString(order , MkString("id") ); _ = id;
  status := this.ParseOrderStatus(this.SafeString(order , MkString("status") )); _ = status;
  timestamp := this.SafeTimestamp(order , MkString("timestamp_created") ); _ = timestamp;
  marketId := this.SafeString(order , MkString("pair") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market ); _ = symbol;
  remaining := OpCopy(MkUndefined() ); _ = remaining;
  amount := OpCopy(MkUndefined() ); _ = amount;
  price := this.SafeNumber(order , MkString("rate") ); _ = price;
  if IsTrue(OpHasMember(MkString("start_amount") , order )) {
    amount = this.SafeNumber(order , MkString("start_amount") );
    remaining = this.SafeNumber(order , MkString("amount") );
  } else {
    remaining = this.SafeNumber(order , MkString("amount") );
  }
  fee := OpCopy(MkUndefined() ); _ = fee;
  return this.SafeOrder(MkMap(&VarMap{
            "info":order ,
            "id":id ,
            "clientOrderId":MkUndefined() ,
            "symbol":symbol ,
            "timestamp":timestamp ,
            "datetime":this.Iso8601(timestamp ),
            "lastTradeTimestamp":MkUndefined() ,
            "type":MkString("limit") ,
            "timeInForce":MkUndefined() ,
            "postOnly":MkUndefined() ,
            "side":this.SafeString(order , MkString("type") ),
            "price":price ,
            "stopPrice":MkUndefined() ,
            "cost":MkUndefined() ,
            "amount":amount ,
            "remaining":remaining ,
            "filled":MkUndefined() ,
            "status":status ,
            "fee":fee ,
            "average":MkUndefined() ,
            "trades":MkUndefined() ,
            }));
}

func (this *Tidex) FetchOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"order_id":parseInt(id )}); _ = request;
  response := this.Call(MkString("privatePostOrderInfo"), this.Extend(request , params )); _ = response;
  id = id.ToString();
  result := this.SafeValue(response , MkString("return") , MkMap(&VarMap{})); _ = result;
  order := this.SafeValue(result , id ); _ = order;
  return this.ParseOrder(this.Extend(MkMap(&VarMap{"id":id }), order ));
}

func (this *Tidex) FetchOpenOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("pair") )= *(market ).At(MkString("id") );
  }
  response := this.Call(MkString("privatePostActiveOrders"), this.Extend(request , params )); _ = response;
  orders := this.SafeValue(response , MkString("return") , MkArray(&VarArray{})); _ = orders;
  return this.ParseOrders(orders , market , since , limit );
}

func (this *Tidex) FetchMyTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := OpCopy(MkUndefined() ); _ = market;
  request := MkMap(&VarMap{}); _ = request;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("pair") )= *(market ).At(MkString("id") );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("count") )= parseInt(limit );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("since") )= parseInt(OpDiv(since , MkInteger(1000) ));
  }
  response := this.Call(MkString("privatePostTradeHistory"), this.Extend(request , params )); _ = response;
  trades := this.SafeValue(response , MkString("return") , MkArray(&VarArray{})); _ = trades;
  return this.ParseTrades(trades , market , since , limit );
}

func (this *Tidex) Withdraw(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  amount := GoGetArg(goArgs, 1, MkUndefined()); _ = amount;
  address := GoGetArg(goArgs, 2, MkUndefined()); _ = address;
  tag := GoGetArg(goArgs, 3, MkUndefined() ); _ = tag;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.CheckAddress(address );
  this.LoadMarkets();
  currency := this.Currency(code ); _ = currency;
  request := MkMap(&VarMap{
        "coinName":*(currency ).At(MkString("id") ),
        "amount":parseFloat(amount ),
        "address":address ,
        }); _ = request;
  if IsTrue(OpNotEq2(tag , MkUndefined() )) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , MkString(" withdraw() does not support the tag argument yet due to a lack of docs on withdrawing with tag/memo on behalf of the exchange.") )));
  }
  response := this.Call(MkString("privatePostWithdrawCoin"), this.Extend(request , params )); _ = response;
  return MkMap(&VarMap{
        "info":response ,
        "id":*(*(response ).At(MkString("return") )).At(MkString("tId") ),
        });
}

func (this *Tidex) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("public") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  url := *(*(*this.At(MkString("urls")) ).At(MkString("api") )).At(api ); _ = url;
  query := this.Omit(params , this.ExtractParams(path )); _ = query;
  if IsTrue(OpEq2(api , MkString("private") )) {
    this.CheckRequiredCredentials();
    nonce := this.Nonce(); _ = nonce;
    body = this.Urlencode(this.Extend(MkMap(&VarMap{
                "nonce":nonce ,
                "method":path ,
                }), query ));
    signature := this.Hmac(this.Encode(body ), this.Encode(*this.At(MkString("secret")) ), MkString("sha512") ); _ = signature;
    headers = MkMap(&VarMap{
        "Content-Type":MkString("application/x-www-form-urlencoded") ,
        "Key":*this.At(MkString("apiKey")) ,
        "Sign":signature ,
        });
  } else {
    if IsTrue(OpEq2(api , MkString("public") )) {
      OpAddAssign(& url , OpAdd(MkString("/") , this.ImplodeParams(path , params )));
      if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
        OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(query )));
      }
    } else {
      OpAddAssign(& url , OpAdd(MkString("/") , this.ImplodeParams(path , params )));
      if IsTrue(OpEq2(method , MkString("GET") )) {
        if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
          OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(query )));
        }
      } else {
        if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
          body = this.Json(query );
          headers = MkMap(&VarMap{"Content-Type":MkString("application/json") });
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

func (this *Tidex) HandleErrors(goArgs ...*Variant) *Variant{
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
  if IsTrue(OpHasMember(MkString("success") , response )) {
    success := this.SafeValue(response , MkString("success") , MkBool(false) ); _ = success;
    if IsTrue(OpEq2(OpType(success ), MkString("string") )) {
      if IsTrue(OpOr(OpEq2(success , MkString("true") ), OpEq2(success , MkString("1") ))) {
        success = OpCopy(MkBool(true) );
      } else {
        success = OpCopy(MkBool(false) );
      }
    }
    if IsTrue(OpNot(success )) {
      code := this.SafeString(response , MkString("code") ); _ = code;
      message := this.SafeString(response , MkString("error") ); _ = message;
      feedback := OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , body )); _ = feedback;
      this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("exact") ), code , feedback );
      this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("exact") ), message , feedback );
      this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("broad") ), message , feedback );
      panic(NewExchangeError(feedback ));
    }
  }
  return MkUndefined()
}

