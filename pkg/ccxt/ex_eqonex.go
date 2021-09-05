package ccxt

import (
)

type Eqonex struct {
	*ExchangeBase
}

var _ Exchange = (*Eqonex)(nil)

func init() {
	exchange := &Eqonex{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Eqonex) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("eqonex") ,
            "name":MkString("EQONEX") ,
            "countries":MkArray(&VarArray{
                MkString("US") ,
                MkString("SG") ,
                }),
            "rateLimit":MkInteger(10) ,
            "has":MkMap(&VarMap{
                "CORS":MkBool(false) ,
                "cancelOrder":MkBool(true) ,
                "createOrder":MkBool(true) ,
                "editOrder":MkBool(true) ,
                "fetchBalance":MkBool(true) ,
                "fetchCanceledOrders":MkBool(true) ,
                "fetchClosedOrders":MkBool(true) ,
                "fetchCurrencies":MkBool(true) ,
                "fetchDepositAddress":MkBool(true) ,
                "fetchDeposits":MkBool(true) ,
                "fetchMarkets":MkBool(true) ,
                "fetchMyTrades":MkBool(true) ,
                "fetchOHLCV":MkBool(true) ,
                "fetchOrder":MkBool(true) ,
                "fetchOrderBook":MkBool(true) ,
                "fetchOrders":MkBool(true) ,
                "fetchTicker":MkBool(false) ,
                "fetchTrades":MkBool(true) ,
                "fetchTradingFees":MkBool(true) ,
                "fetchWithdrawals":MkBool(true) ,
                "withdraw":MkBool(true) ,
                }),
            "timeframes":MkMap(&VarMap{
                "1m":MkInteger(1) ,
                "5m":MkInteger(2) ,
                "15m":MkInteger(3) ,
                "1h":MkInteger(4) ,
                "6h":MkInteger(5) ,
                "1d":MkInteger(6) ,
                "7d":MkInteger(7) ,
                }),
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/51840849/122649755-1a076c80-d138-11eb-8f2e-9a9166a03d79.jpg") ,
                "test":MkMap(&VarMap{
                    "public":MkString("https://testnet.eqonex.com/api") ,
                    "private":MkString("https://testnet.eqonex.com/api") ,
                    }),
                "api":MkMap(&VarMap{
                    "public":MkString("https://eqonex.com/api") ,
                    "private":MkString("https://eqonex.com/api") ,
                    }),
                "www":MkString("https://eqonex.com") ,
                "doc":MkArray(&VarArray{
                    MkString("https://developer.eqonex.com") ,
                    }),
                "referral":MkString("https://eqonex.com?referredByCode=zpa8kij4ouvBFup3") ,
                }),
            "api":MkMap(&VarMap{
                "public":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("health") ,
                        MkString("getInstruments") ,
                        MkString("getInstrumentPairs") ,
                        MkString("getOrderBook") ,
                        MkString("getRisk") ,
                        MkString("getTradeHistory") ,
                        MkString("getFundingRateHistory") ,
                        MkString("getChart") ,
                        MkString("getExchangeInfo") ,
                        })}),
                "private":MkMap(&VarMap{"post":MkArray(&VarArray{
                        MkString("logon") ,
                        MkString("order") ,
                        MkString("cancelOrder") ,
                        MkString("cancelReplaceOrder") ,
                        MkString("getOrder") ,
                        MkString("getOrders") ,
                        MkString("getOrderStatus") ,
                        MkString("getOrderHistory") ,
                        MkString("userTrades") ,
                        MkString("getPositions") ,
                        MkString("cancelAll") ,
                        MkString("getUserHistory") ,
                        MkString("getRisk") ,
                        MkString("getDepositAddresses") ,
                        MkString("getDepositHistory") ,
                        MkString("getWithdrawRequests") ,
                        MkString("sendWithdrawRequest") ,
                        MkString("getTransferHistory") ,
                        })}),
                }),
            "requiredCredentials":MkMap(&VarMap{
                "apiKey":MkBool(true) ,
                "secret":MkBool(true) ,
                "uid":MkBool(true) ,
                }),
            "exceptions":MkMap(&VarMap{"broad":MkMap(&VarMap{"symbol not found":BadSymbol })}),
            }));
}

func (this *Eqonex) FetchMarkets(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"verbose":MkBool(true) }); _ = request;
  response := this.Call(MkString("publicGetGetInstrumentPairs"), this.Extend(request , params )); _ = response;
  instrumentPairs := this.SafeValue(response , MkString("instrumentPairs") , MkArray(&VarArray{})); _ = instrumentPairs;
  markets := MkArray(&VarArray{}); _ = markets;
  for i := MkInteger(0) ; IsTrue(OpLw(i , instrumentPairs.Length )); OpIncr(& i ){
    market := this.ParseMarket(*(instrumentPairs ).At(i )); _ = market;
    markets.Push(market );
  }
  return markets ;
}

func (this *Eqonex) ParseMarket(goArgs ...*Variant) *Variant{
  market := GoGetArg(goArgs, 0, MkUndefined()); _ = market;
  id := this.SafeString(market , MkString("instrumentId") ); _ = id;
  uppercaseId := this.SafeString(market , MkString("symbol") ); _ = uppercaseId;
  assetType := this.SafeString(market , MkString("assetType") ); _ = assetType;
  spot := OpEq2(assetType , MkString("PAIR") ); _ = spot;
  swap := OpEq2(assetType , MkString("PERPETUAL_SWAP") ); _ = swap;
  type_ := OpTrinary(swap , MkString("swap") , MkString("spot") ); _ = type_;
  baseId := this.SafeString(market , MkString("currency") ); _ = baseId;
  quoteId := this.SafeString(market , MkString("contAmtCurr") ); _ = quoteId;
  base := this.SafeCurrencyCode(baseId ); _ = base;
  quote := this.SafeCurrencyCode(quoteId ); _ = quote;
  symbol := OpTrinary(swap , uppercaseId , OpAdd(base , OpAdd(MkString("/") , quote ))); _ = symbol;
  status := this.SafeInteger(market , MkString("securityStatus") ); _ = status;
  active := OpEq2(status , MkInteger(1) ); _ = active;
  precision := MkMap(&VarMap{
        "amount":this.SafeInteger(market , MkString("quantity_scale") ),
        "price":this.SafeInteger(market , MkString("price_scale") ),
        }); _ = precision;
  return MkMap(&VarMap{
        "id":id ,
        "uppercaseId":uppercaseId ,
        "symbol":symbol ,
        "base":base ,
        "quote":quote ,
        "baseId":baseId ,
        "quoteId":quoteId ,
        "type":type_ ,
        "spot":spot ,
        "swap":swap ,
        "active":active ,
        "precision":precision ,
        "limits":MkMap(&VarMap{
            "amount":MkMap(&VarMap{
                "min":this.SafeNumber(market , MkString("minTradeVol") ),
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
        "info":market ,
        });
}

func (this *Eqonex) FetchCurrencies(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetGetInstruments"), params ); _ = response;
  currencies := MkMap(&VarMap{}); _ = currencies;
  instruments := this.SafeValue(response , MkString("instruments") , MkArray(&VarArray{})); _ = instruments;
  for i := MkInteger(0) ; IsTrue(OpLw(i , instruments.Length )); OpIncr(& i ){
    currency := this.ParseCurrency(*(instruments ).At(i )); _ = currency;
    code := *(currency ).At(MkString("code") ); _ = code;
    *(currencies ).At(code )= OpCopy(currency );
  }
  return currencies ;
}

func (this *Eqonex) ParseCurrency(goArgs ...*Variant) *Variant{
  currency := GoGetArg(goArgs, 0, MkUndefined()); _ = currency;
  id := this.SafeString(currency , MkInteger(0) ); _ = id;
  uppercaseId := this.SafeString(currency , MkInteger(1) ); _ = uppercaseId;
  code := this.SafeCurrencyCode(uppercaseId ); _ = code;
  priceScale := this.SafeInteger(currency , MkInteger(2) ); _ = priceScale;
  amountScale := this.SafeInteger(currency , MkInteger(3) ); _ = amountScale;
  precision := Math.Max(priceScale , amountScale ); _ = precision;
  name := this.SafeString(currency , MkInteger(6) ); _ = name;
  status := this.SafeInteger(currency , MkInteger(4) ); _ = status;
  active := OpEq2(status , MkInteger(1) ); _ = active;
  fee := this.SafeNumber(currency , MkInteger(5) ); _ = fee;
  return MkMap(&VarMap{
        "id":id ,
        "info":currency ,
        "uppercaseId":uppercaseId ,
        "code":code ,
        "name":name ,
        "precision":precision ,
        "fee":fee ,
        "active":active ,
        "limits":MkMap(&VarMap{
            "amount":MkMap(&VarMap{
                "min":MkUndefined() ,
                "max":MkUndefined() ,
                }),
            "withdraw":MkMap(&VarMap{
                "min":MkUndefined() ,
                "max":MkUndefined() ,
                }),
            }),
        });
}

func (this *Eqonex) FetchOHLCV(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  timeframe := GoGetArg(goArgs, 1, MkString("1m") ); _ = timeframe;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "pairId":parseInt(*(market ).At(MkString("id") )),
        "timespan":*(*this.At(MkString("timeframes")) ).At(timeframe ),
        }); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("publicGetGetChart"), this.Extend(request , params )); _ = response;
  chart := this.SafeValue(response , MkString("chart") , MkArray(&VarArray{})); _ = chart;
  return this.ParseOHLCVs(chart , market , timeframe , since , limit );
}

func (this *Eqonex) ParseOHLCV(goArgs ...*Variant) *Variant{
  ohlcv := GoGetArg(goArgs, 0, MkUndefined()); _ = ohlcv;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.SafeInteger(ohlcv , MkInteger(0) ); _ = timestamp;
  open := this.ParseNumber(this.ConvertFromScale(this.SafeString(ohlcv , MkInteger(1) ), *(*(market ).At(MkString("precision") )).At(MkString("price") ))); _ = open;
  high := this.ParseNumber(this.ConvertFromScale(this.SafeString(ohlcv , MkInteger(2) ), *(*(market ).At(MkString("precision") )).At(MkString("price") ))); _ = high;
  low := this.ParseNumber(this.ConvertFromScale(this.SafeString(ohlcv , MkInteger(3) ), *(*(market ).At(MkString("precision") )).At(MkString("price") ))); _ = low;
  close := this.ParseNumber(this.ConvertFromScale(this.SafeString(ohlcv , MkInteger(4) ), *(*(market ).At(MkString("precision") )).At(MkString("price") ))); _ = close;
  volume := this.ParseNumber(this.ConvertFromScale(this.SafeString(ohlcv , MkInteger(5) ), *(*(market ).At(MkString("precision") )).At(MkString("amount") ))); _ = volume;
  return MkArray(&VarArray{
        timestamp ,
        open ,
        high ,
        low ,
        close ,
        volume ,
        });
}

func (this *Eqonex) ParseBidAsk(goArgs ...*Variant) *Variant{
  bidask := GoGetArg(goArgs, 0, MkUndefined()); _ = bidask;
  priceKey := GoGetArg(goArgs, 1, MkInteger(0) ); _ = priceKey;
  amountKey := GoGetArg(goArgs, 2, MkInteger(1) ); _ = amountKey;
  market := GoGetArg(goArgs, 3, MkUndefined() ); _ = market;
  if IsTrue(OpEq2(market , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" parseBidAsk() requires a market argument") )));
  }
  priceString := this.SafeString(bidask , priceKey ); _ = priceString;
  amountString := this.SafeString(bidask , amountKey ); _ = amountString;
  return MkArray(&VarArray{
        this.ParseNumber(this.ConvertFromScale(priceString , *(*(market ).At(MkString("precision") )).At(MkString("price") ))),
        this.ParseNumber(this.ConvertFromScale(amountString , *(*(market ).At(MkString("precision") )).At(MkString("amount") ))),
        });
}

func (this *Eqonex) ParseOrderBook(goArgs ...*Variant) *Variant{
  orderbook := GoGetArg(goArgs, 0, MkUndefined()); _ = orderbook;
  symbol := GoGetArg(goArgs, 1, MkUndefined()); _ = symbol;
  timestamp := GoGetArg(goArgs, 2, MkUndefined() ); _ = timestamp;
  bidsKey := GoGetArg(goArgs, 3, MkString("bids") ); _ = bidsKey;
  asksKey := GoGetArg(goArgs, 4, MkString("asks") ); _ = asksKey;
  priceKey := GoGetArg(goArgs, 5, MkInteger(0) ); _ = priceKey;
  amountKey := GoGetArg(goArgs, 6, MkInteger(1) ); _ = amountKey;
  market := GoGetArg(goArgs, 7, MkUndefined() ); _ = market;
  result := MkMap(&VarMap{
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "nonce":MkUndefined() ,
        }); _ = result;
  sides := MkArray(&VarArray{
        bidsKey ,
        asksKey ,
        }); _ = sides;
  for i := MkInteger(0) ; IsTrue(OpLw(i , sides.Length )); OpIncr(& i ){
    side := *(sides ).At(i ); _ = side;
    orders := MkArray(&VarArray{}); _ = orders;
    bidasks := this.SafeValue(orderbook , side ); _ = bidasks;
    for k := MkInteger(0) ; IsTrue(OpLw(k , bidasks.Length )); OpIncr(& k ){
      orders.Push(this.ParseBidAsk(*(bidasks ).At(k ), priceKey , amountKey , market ));
    }
    *(result ).At(side )= OpCopy(orders );
  }
  *(result ).At(bidsKey )= this.SortBy(*(result ).At(bidsKey ), MkInteger(0) , MkBool(true) );
  *(result ).At(asksKey )= this.SortBy(*(result ).At(asksKey ), MkInteger(0) );
  return result ;
}

func (this *Eqonex) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"pairId":parseInt(*(market ).At(MkString("id") ))}); _ = request;
  response := this.Call(MkString("publicGetGetOrderBook"), this.Extend(request , params )); _ = response;
  return this.ParseOrderBook(response , symbol , MkUndefined() , MkString("bids") , MkString("asks") , MkInteger(0) , MkInteger(1) , market );
}

func (this *Eqonex) FetchTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"pairId":parseInt(*(market ).At(MkString("id") ))}); _ = request;
  response := this.Call(MkString("publicGetGetTradeHistory"), this.Extend(request , params )); _ = response;
  trades := this.SafeValue(response , MkString("trades") , MkArray(&VarArray{})); _ = trades;
  return this.ParseTrades(trades , market , since , limit , params );
}

func (this *Eqonex) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  id := OpCopy(MkUndefined() ); _ = id;
  timestamp := OpCopy(MkUndefined() ); _ = timestamp;
  orderId := OpCopy(MkUndefined() ); _ = orderId;
  type_ := OpCopy(MkUndefined() ); _ = type_;
  side := OpCopy(MkUndefined() ); _ = side;
  priceString := OpCopy(MkUndefined() ); _ = priceString;
  amountString := OpCopy(MkUndefined() ); _ = amountString;
  fee := OpCopy(MkUndefined() ); _ = fee;
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  if IsTrue(GoIsArray(trade )) {
    id = this.SafeString(trade , MkInteger(3) );
    priceString = this.ConvertFromScale(this.SafeString(trade , MkInteger(0) ), *(*(market ).At(MkString("precision") )).At(MkString("price") ));
    amountString = this.ConvertFromScale(this.SafeString(trade , MkInteger(1) ), *(*(market ).At(MkString("precision") )).At(MkString("amount") ));
    timestamp = this.ToMilliseconds(this.SafeString(trade , MkInteger(2) ));
    takerSide := this.SafeInteger(trade , MkInteger(4) ); _ = takerSide;
    if IsTrue(OpEq2(takerSide , MkInteger(1) )) {
      side = MkString("buy") ;
    } else {
      if IsTrue(OpEq2(takerSide , MkInteger(2) )) {
        side = MkString("sell") ;
      }
    }
  } else {
    id = this.SafeString(trade , MkString("execId") );
    timestamp = this.SafeInteger(trade , MkString("time") );
    marketId := this.SafeString(trade , MkString("symbol") ); _ = marketId;
    symbol = this.SafeSymbol(marketId , market );
    orderId = this.SafeString(trade , MkString("orderId") );
    side = this.SafeStringLower(trade , MkString("side") );
    type_ = this.ParseOrderType(this.SafeString(trade , MkString("ordType") ));
    priceString = this.SafeString(trade , MkString("lastPx") );
    amountString = this.SafeString(trade , MkString("qty") );
    feeCost := this.SafeNumber(trade , MkString("commission") ); _ = feeCost;
    if IsTrue(OpNotEq2(feeCost , MkUndefined() )) {
      feeCost = OpNeg(feeCost );
      feeCurrencyId := this.SafeString(trade , MkString("commCurrency") ); _ = feeCurrencyId;
      feeCurrencyCode := this.SafeCurrencyCode(feeCurrencyId ); _ = feeCurrencyCode;
      fee = MkMap(&VarMap{
          "cost":feeCost ,
          "currency":feeCurrencyCode ,
          });
    }
  }
  if IsTrue(OpAnd(OpEq2(symbol , MkUndefined() ), OpNotEq2(market , MkUndefined() ))) {
    symbol = *(market ).At(MkString("symbol") );
  }
  cost := this.ParseNumber(Precise.StringMul(amountString , priceString )); _ = cost;
  price := this.ParseNumber(priceString ); _ = price;
  amount := this.ParseNumber(amountString ); _ = amount;
  return MkMap(&VarMap{
        "info":trade ,
        "id":id ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "symbol":symbol ,
        "order":orderId ,
        "type":type_ ,
        "side":side ,
        "takerOrMaker":MkUndefined() ,
        "price":price ,
        "amount":amount ,
        "cost":cost ,
        "fee":fee ,
        });
}

func (this *Eqonex) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("privatePostGetPositions"), params ); _ = response;
  positions := this.SafeValue(response , MkString("positions") , MkArray(&VarArray{})); _ = positions;
  result := MkMap(&VarMap{"info":response }); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , positions.Length )); OpIncr(& i ){
    position := *(positions ).At(i ); _ = position;
    assetType := this.SafeString(position , MkString("assetType") ); _ = assetType;
    if IsTrue(OpEq2(assetType , MkString("ASSET") )) {
      currencyId := this.SafeString(position , MkString("symbol") ); _ = currencyId;
      code := this.SafeCurrencyCode(currencyId ); _ = code;
      quantityString := this.SafeString(position , MkString("quantity") ); _ = quantityString;
      availableQuantityString := this.SafeString(position , MkString("availableQuantity") ); _ = availableQuantityString;
      scale := this.SafeInteger(position , MkString("quantity_scale") ); _ = scale;
      account := this.Account(); _ = account;
      *(account ).At(MkString("free") )= this.ConvertFromScale(availableQuantityString , scale );
      *(account ).At(MkString("total") )= this.ConvertFromScale(quantityString , scale );
      *(result ).At(code )= OpCopy(account );
    }
  }
  return this.ParseBalance(result );
}

func (this *Eqonex) CreateOrder(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 2, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 3, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 4, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 5, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  orderSide := OpTrinary(OpEq2(side , MkString("buy") ), MkInteger(1) , MkInteger(2) ); _ = orderSide;
  quantityScale := this.GetScale(amount ); _ = quantityScale;
  request := MkMap(&VarMap{
        "instrumentId":parseInt(*(market ).At(MkString("id") )),
        "symbol":*(market ).At(MkString("uppercaseId") ),
        "side":orderSide ,
        "quantity":this.ConvertToScale(this.NumberToString(amount ), quantityScale ),
        "quantity_scale":quantityScale ,
        }); _ = request;
  if IsTrue(OpEq2(type_ , MkString("market") )) {
    *(request ).At(MkString("ordType") )= MkInteger(1) ;
  } else {
    if IsTrue(OpEq2(type_ , MkString("limit") )) {
      *(request ).At(MkString("ordType") )= MkInteger(2) ;
      priceScale := this.GetScale(price ); _ = priceScale;
      *(request ).At(MkString("price") )= this.ConvertToScale(this.NumberToString(price ), priceScale );
      *(request ).At(MkString("priceScale") )= OpCopy(priceScale );
    } else {
      stopPrice := this.SafeNumber2(params , MkString("stopPrice") , MkString("stopPx") ); _ = stopPrice;
      params = this.Omit(params , MkArray(&VarArray{
              MkString("stopPrice") ,
              MkString("stopPx") ,
              }));
      if IsTrue(OpEq2(stopPrice , MkUndefined() )) {
        if IsTrue(OpEq2(type_ , MkString("stop") )) {
          if IsTrue(OpEq2(price , MkUndefined() )) {
            panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" createOrder() requires a price argument or a stopPrice parameter or a stopPx parameter for ") , OpAdd(type_ , MkString(" orders") )))));
          }
          *(request ).At(MkString("ordType") )= MkInteger(3) ;
          *(request ).At(MkString("stopPx") )= this.ConvertToScale(this.NumberToString(price ), this.GetScale(price ));
        } else {
          if IsTrue(OpEq2(type_ , MkString("stop limit") )) {
            panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" createOrder() requires a stopPrice parameter or a stopPx parameter for ") , OpAdd(type_ , MkString(" orders") )))));
          }
        }
      } else {
        if IsTrue(OpEq2(type_ , MkString("stop") )) {
          *(request ).At(MkString("ordType") )= MkInteger(3) ;
          *(request ).At(MkString("stopPx") )= this.ConvertToScale(this.NumberToString(stopPrice ), this.GetScale(stopPrice ));
        } else {
          if IsTrue(OpEq2(type_ , MkString("stop limit") )) {
            *(request ).At(MkString("ordType") )= MkInteger(4) ;
            priceScale := this.GetScale(price ); _ = priceScale;
            stopPriceScale := this.GetScale(stopPrice ); _ = stopPriceScale;
            *(request ).At(MkString("price_scale") )= OpCopy(priceScale );
            *(request ).At(MkString("stopPx_scale") )= OpCopy(stopPriceScale );
            *(request ).At(MkString("stopPx") )= this.ConvertToScale(this.NumberToString(stopPrice ), stopPriceScale );
            *(request ).At(MkString("price") )= this.ConvertToScale(this.NumberToString(price ), priceScale );
          }
        }
      }
    }
  }
  response := this.Call(MkString("privatePostOrder"), this.Extend(request , params )); _ = response;
  return this.ParseOrder(response , market );
}

func (this *Eqonex) CancelOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" cancelOrder() requires a symbol argument") )));
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "origOrderId":parseInt(id ),
        "instrumentId":parseInt(*(market ).At(MkString("id") )),
        }); _ = request;
  response := this.Call(MkString("privatePostCancelOrder"), this.Extend(request , params )); _ = response;
  return this.ParseOrder(response , market );
}

func (this *Eqonex) EditOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 2, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 3, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 4, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 5, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 6, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  orderSide := OpTrinary(OpEq2(side , MkString("buy") ), MkInteger(1) , MkInteger(2) ); _ = orderSide;
  quantityScale := this.GetScale(amount ); _ = quantityScale;
  request := MkMap(&VarMap{
        "origOrderId":id ,
        "instrumentId":parseInt(*(market ).At(MkString("id") )),
        "symbol":*(market ).At(MkString("uppercaseId") ),
        "side":orderSide ,
        "quantity":this.ConvertToScale(this.NumberToString(amount ), quantityScale ),
        "quantity_scale":quantityScale ,
        }); _ = request;
  if IsTrue(OpEq2(type_ , MkString("market") )) {
    *(request ).At(MkString("ordType") )= MkInteger(1) ;
  } else {
    if IsTrue(OpEq2(type_ , MkString("limit") )) {
      *(request ).At(MkString("ordType") )= MkInteger(2) ;
      *(request ).At(MkString("price") )= this.ConvertToScale(this.NumberToString(price ), this.GetScale(price ));
    } else {
      stopPrice := this.SafeNumber2(params , MkString("stopPrice") , MkString("stopPx") ); _ = stopPrice;
      params = this.Omit(params , MkArray(&VarArray{
              MkString("stopPrice") ,
              MkString("stopPx") ,
              }));
      if IsTrue(OpEq2(stopPrice , MkUndefined() )) {
        if IsTrue(OpEq2(type_ , MkString("stop") )) {
          if IsTrue(OpEq2(price , MkUndefined() )) {
            panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" editOrder() requires a price argument or a stopPrice parameter or a stopPx parameter for ") , OpAdd(type_ , MkString(" orders") )))));
          }
          *(request ).At(MkString("ordType") )= MkInteger(3) ;
          *(request ).At(MkString("stopPx") )= this.ConvertToScale(this.NumberToString(price ), this.GetScale(price ));
        } else {
          if IsTrue(OpEq2(type_ , MkString("stop limit") )) {
            panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" editOrder() requires a stopPrice parameter or a stopPx parameter for ") , OpAdd(type_ , MkString(" orders") )))));
          }
        }
      } else {
        if IsTrue(OpEq2(type_ , MkString("stop") )) {
          *(request ).At(MkString("ordType") )= MkInteger(3) ;
          *(request ).At(MkString("stopPx") )= this.ConvertToScale(this.NumberToString(stopPrice ), this.GetScale(stopPrice ));
        } else {
          if IsTrue(OpEq2(type_ , MkString("stop limit") )) {
            *(request ).At(MkString("ordType") )= MkInteger(4) ;
            priceScale := this.GetScale(price ); _ = priceScale;
            stopPriceScale := this.GetScale(stopPrice ); _ = stopPriceScale;
            *(request ).At(MkString("price_scale") )= OpCopy(priceScale );
            *(request ).At(MkString("stopPx_scale") )= OpCopy(stopPriceScale );
            *(request ).At(MkString("stopPx") )= this.ConvertToScale(this.NumberToString(stopPrice ), stopPriceScale );
            *(request ).At(MkString("price") )= this.ConvertToScale(this.NumberToString(price ), priceScale );
          }
        }
      }
    }
  }
  response := this.Call(MkString("privatePostOrder"), this.Extend(request , params )); _ = response;
  return this.ParseOrder(response , market );
}

func (this *Eqonex) FetchOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"orderId":parseInt(id )}); _ = request;
  response := this.Call(MkString("privatePostGetOrderStatus"), this.Extend(request , params )); _ = response;
  return this.ParseOrder(response );
}

func (this *Eqonex) FetchClosedOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"ordStatus":MkString("2") }); _ = request;
  return this.FetchOrders(symbol , since , limit , this.Extend(request , params ));
}

func (this *Eqonex) FetchCanceledOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"ordStatus":MkString("4") }); _ = request;
  return this.FetchOrders(symbol , since , limit , this.Extend(request , params ));
}

func (this *Eqonex) FetchOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := OpCopy(MkUndefined() ); _ = market;
  request := MkMap(&VarMap{}); _ = request;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("instrumentId") )= parseInt(*(market ).At(MkString("id") ));
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("privatePostGetOrders"), this.Extend(request , params )); _ = response;
  orders := this.SafeValue(response , MkString("orders") , MkArray(&VarArray{})); _ = orders;
  return this.ParseOrders(orders , market , since , limit , params );
}

func (this *Eqonex) FetchMyTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("instrumentId") )= *(market ).At(MkString("id") );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("startTime") )= OpCopy(since );
  }
  response := this.Call(MkString("privatePostUserTrades"), this.Extend(request , params )); _ = response;
  trades := this.SafeValue(response , MkString("trades") , MkArray(&VarArray{})); _ = trades;
  return this.ParseTrades(trades , market , since , limit , params );
}

func (this *Eqonex) FetchDepositAddress(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  currency := this.Currency(code ); _ = currency;
  request := MkMap(&VarMap{"instrumentId":parseInt(*(currency ).At(MkString("id") ))}); _ = request;
  response := this.Call(MkString("privatePostGetDepositAddresses"), this.Extend(request , params )); _ = response;
  addresses := this.SafeValue(response , MkString("addresses") , MkArray(&VarArray{})); _ = addresses;
  address := this.SafeValue(addresses , MkInteger(0) ); _ = address;
  return this.ParseDepositAddress(address , currency );
}

func (this *Eqonex) ParseDepositAddress(goArgs ...*Variant) *Variant{
  depositAddress := GoGetArg(goArgs, 0, MkUndefined()); _ = depositAddress;
  currency := GoGetArg(goArgs, 1, MkUndefined() ); _ = currency;
  currencyId := this.SafeString(depositAddress , MkString("symbol") ); _ = currencyId;
  code := this.SafeCurrencyCode(currencyId , currency ); _ = code;
  address := this.SafeString(depositAddress , MkString("address") ); _ = address;
  this.CheckAddress(address );
  return MkMap(&VarMap{
        "currency":code ,
        "address":address ,
        "tag":MkUndefined() ,
        "info":depositAddress ,
        });
}

func (this *Eqonex) FetchDeposits(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  currency := OpCopy(MkUndefined() ); _ = currency;
  if IsTrue(OpNotEq2(code , MkUndefined() )) {
    currency = this.Currency(code );
    *(request ).At(MkString("instrumentId") )= parseInt(*(currency ).At(MkString("id") ));
  }
  response := this.Call(MkString("privatePostGetDepositHistory"), this.Extend(request , params )); _ = response;
  deposits := this.SafeValue(response , MkString("deposits") , MkArray(&VarArray{})); _ = deposits;
  return this.ParseTransactions(deposits , currency , since , limit , MkMap(&VarMap{"type":MkString("deposit") }));
}

func (this *Eqonex) FetchWithdrawals(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  currency := OpCopy(MkUndefined() ); _ = currency;
  if IsTrue(OpNotEq2(code , MkUndefined() )) {
    currency = this.Currency(code );
    *(request ).At(MkString("instrumentId") )= parseInt(*(currency ).At(MkString("id") ));
  }
  response := this.Call(MkString("privatePostGetWithdrawRequests"), this.Extend(request , params )); _ = response;
  withdrawals := this.SafeValue(response , MkString("addresses") , MkArray(&VarArray{})); _ = withdrawals;
  return this.ParseTransactions(withdrawals , currency , since , limit , MkMap(&VarMap{"type":MkString("withdrawal") }));
}

func (this *Eqonex) ParseTransaction(goArgs ...*Variant) *Variant{
  transaction := GoGetArg(goArgs, 0, MkUndefined()); _ = transaction;
  currency := GoGetArg(goArgs, 1, MkUndefined() ); _ = currency;
  id := this.SafeString(transaction , MkString("id") , MkString("transactionId") ); _ = id;
  txid := this.SafeString(transaction , MkString("transactionUuid") ); _ = txid;
  timestamp := this.SafeInteger(transaction , MkString("timestamp") ); _ = timestamp;
  address := this.SafeString(transaction , MkString("address") ); _ = address;
  if IsTrue(OpEq2(address , MkString("null") )) {
    address = OpCopy(MkUndefined() );
  }
  type_ := this.SafeString(transaction , MkString("type") ); _ = type_;
  amount := this.SafeNumber(transaction , MkString("balance_change") ); _ = amount;
  if IsTrue(OpEq2(amount , MkUndefined() )) {
    amount = this.SafeString(transaction , MkString("quantity") );
    amountScale := this.SafeInteger(transaction , MkString("quantity_scale") ); _ = amountScale;
    amount = this.ParseNumber(this.ConvertFromScale(amount , amountScale ));
  }
  currencyId := this.SafeString(transaction , MkString("symbol") ); _ = currencyId;
  code := this.SafeCurrencyCode(currencyId , currency ); _ = code;
  status := this.ParseTransactionStatus(this.SafeString(transaction , MkString("status") )); _ = status;
  return MkMap(&VarMap{
        "info":transaction ,
        "id":id ,
        "txid":txid ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "addressFrom":MkUndefined() ,
        "address":address ,
        "addressTo":MkUndefined() ,
        "tagFrom":MkUndefined() ,
        "tag":MkUndefined() ,
        "tagTo":MkUndefined() ,
        "type":type_ ,
        "amount":amount ,
        "currency":code ,
        "status":status ,
        "updated":MkUndefined() ,
        "comment":MkUndefined() ,
        "fee":MkUndefined() ,
        });
}

func (this *Eqonex) ParseTransactionStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "0":MkString("pending") ,
        "1":MkString("ok") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Eqonex) Withdraw(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  amount := GoGetArg(goArgs, 1, MkUndefined()); _ = amount;
  address := GoGetArg(goArgs, 2, MkUndefined()); _ = address;
  tag := GoGetArg(goArgs, 3, MkUndefined() ); _ = tag;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.CheckAddress(address );
  this.LoadMarkets();
  currency := this.Currency(code ); _ = currency;
  scale := this.GetScale(amount ); _ = scale;
  quantity := this.ConvertToScale(amount , scale ); _ = quantity;
  request := MkMap(&VarMap{
        "instrumentId":parseInt(*(currency ).At(MkString("id") )),
        "symbol":*(currency ).At(MkString("uppercaseId") ),
        "quantity":quantity ,
        "quantity_scale":scale ,
        "address":address ,
        }); _ = request;
  response := this.Call(MkString("privatePostSendWithdrawRequest"), this.Extend(request , params )); _ = response;
  return this.ParseTransaction(response , currency );
}

func (this *Eqonex) FetchTradingFees(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetGetExchangeInfo"), params ); _ = response;
  tradingFees := this.SafeValue(response , MkString("spotFees") , MkArray(&VarArray{})); _ = tradingFees;
  taker := MkMap(&VarMap{}); _ = taker;
  maker := MkMap(&VarMap{}); _ = maker;
  for i := MkInteger(0) ; IsTrue(OpLw(i , tradingFees.Length )); OpIncr(& i ){
    tradingFee := *(tradingFees ).At(i ); _ = tradingFee;
    if IsTrue(OpNotEq2(this.SafeString(tradingFee , MkString("tier") ), MkUndefined() )) {
      *(taker ).At(*(tradingFee ).At(MkString("tier") ))= this.SafeNumber(tradingFee , MkString("taker") );
      *(maker ).At(*(tradingFee ).At(MkString("tier") ))= this.SafeNumber(tradingFee , MkString("maker") );
    }
  }
  return MkMap(&VarMap{
        "info":tradingFees ,
        "tierBased":MkBool(true) ,
        "maker":maker ,
        "taker":taker ,
        });
}

func (this *Eqonex) FetchTradingLimits(goArgs ...*Variant) *Variant{
  symbols := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbols;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("publicGetGetExchangeInfo"), params ); _ = response;
  tradingLimits := this.SafeValue(response , MkString("tradingLimits") , MkArray(&VarArray{})); _ = tradingLimits;
  return MkMap(&VarMap{
        "info":tradingLimits ,
        "limits":MkMap(&VarMap{
            "amount":MkMap(&VarMap{
                "min":MkUndefined() ,
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
        });
}

func (this *Eqonex) FetchFundingLimits(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetGetExchangeInfo"), params ); _ = response;
  withdrawLimits := this.SafeValue(response , MkString("withdrawLimits") , MkArray(&VarArray{})); _ = withdrawLimits;
  return MkMap(&VarMap{
        "info":withdrawLimits ,
        "withdraw":MkUndefined() ,
        });
}

func (this *Eqonex) ParseOrder(goArgs ...*Variant) *Variant{
  order := GoGetArg(goArgs, 0, MkUndefined()); _ = order;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  id := this.SafeString2(order , MkString("orderId") , MkString("id") ); _ = id;
  id = this.SafeString(order , MkString("origOrderId") , id );
  clientOrderId := this.SafeString(order , MkString("clOrdId") ); _ = clientOrderId;
  type_ := this.ParseOrderType(this.SafeString(order , MkString("ordType") )); _ = type_;
  side := this.ParseOrderSide(this.SafeString(order , MkString("side") )); _ = side;
  status := this.ParseOrderStatus(this.SafeString(order , MkString("ordStatus") )); _ = status;
  marketId := this.SafeString(order , MkString("instrumentId") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market ); _ = symbol;
  timestamp := this.ToMilliseconds(this.SafeString(order , MkString("timeStamp") )); _ = timestamp;
  lastTradeTimestamp := OpCopy(MkUndefined() ); _ = lastTradeTimestamp;
  priceString := this.SafeString(order , MkString("price") ); _ = priceString;
  priceScale := this.SafeInteger(order , MkString("price_scale") ); _ = priceScale;
  price := this.ParseNumber(this.ConvertFromScale(priceString , priceScale )); _ = price;
  amountString := this.SafeString(order , MkString("quantity") ); _ = amountString;
  amountScale := this.SafeInteger(order , MkString("quantity_scale") ); _ = amountScale;
  amount := this.ParseNumber(this.ConvertFromScale(amountString , amountScale )); _ = amount;
  filledString := this.SafeString(order , MkString("cumQty") ); _ = filledString;
  filledScale := this.SafeInteger(order , MkString("cumQty_scale") ); _ = filledScale;
  filled := this.ParseNumber(this.ConvertFromScale(filledString , filledScale )); _ = filled;
  remainingString := this.SafeString(order , MkString("leavesQty") ); _ = remainingString;
  remainingScale := this.SafeInteger(order , MkString("leavesQty_scale") ); _ = remainingScale;
  remaining := this.ParseNumber(this.ConvertFromScale(remainingString , remainingScale )); _ = remaining;
  fee := OpCopy(MkUndefined() ); _ = fee;
  currencyId := this.SafeInteger(order , MkString("feeInstrumentId") ); _ = currencyId;
  feeCurrencyCode := this.SafeCurrencyCode(currencyId ); _ = feeCurrencyCode;
  feeCost := this.SafeString(order , MkString("feeTotal") ); _ = feeCost;
  feeScale := this.SafeInteger(order , MkString("fee_scale") ); _ = feeScale;
  if IsTrue(OpNotEq2(feeCost , MkUndefined() )) {
    feeCost = Precise.StringNeg(feeCost );
    feeCost = this.ParseNumber(this.ConvertFromScale(feeCost , feeScale ));
  }
  if IsTrue(OpNotEq2(feeCost , MkUndefined() )) {
    fee = MkMap(&VarMap{
        "currency":feeCurrencyCode ,
        "cost":feeCost ,
        "rate":MkUndefined() ,
        });
  }
  timeInForce := this.ParseTimeInForce(this.SafeString(order , MkString("timeInForce") )); _ = timeInForce;
  if IsTrue(OpEq2(timeInForce , MkString("0") )) {
    timeInForce = OpCopy(MkUndefined() );
  }
  stopPriceScale := this.SafeInteger(order , MkString("stopPx_scale") , MkInteger(0) ); _ = stopPriceScale;
  stopPrice := this.ParseNumber(this.ConvertFromScale(this.SafeString(order , MkString("stopPx") ), stopPriceScale )); _ = stopPrice;
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
            "postOnly":MkUndefined() ,
            "side":side ,
            "price":price ,
            "stopPrice":stopPrice ,
            "amount":amount ,
            "cost":MkUndefined() ,
            "average":MkUndefined() ,
            "filled":filled ,
            "remaining":remaining ,
            "status":status ,
            "fee":fee ,
            "trades":MkUndefined() ,
            }));
}

func (this *Eqonex) ParseOrderStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "0":MkString("open") ,
        "1":MkString("open") ,
        "2":MkString("closed") ,
        "3":MkString("open") ,
        "4":MkString("canceled") ,
        "5":MkString("canceled") ,
        "6":MkString("canceling") ,
        "7":MkString("canceled") ,
        "8":MkString("rejected") ,
        "9":MkString("canceled") ,
        "A":MkString("open") ,
        "B":MkString("open") ,
        "C":MkString("expired") ,
        "D":MkString("open") ,
        "E":MkString("canceling") ,
        "F":MkString("open") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Eqonex) ParseOrderSide(goArgs ...*Variant) *Variant{
  side := GoGetArg(goArgs, 0, MkUndefined()); _ = side;
  sides := MkMap(&VarMap{
        "1":MkString("buy") ,
        "2":MkString("sell") ,
        }); _ = sides;
  return this.SafeString(sides , side , side );
}

func (this *Eqonex) ParseOrderType(goArgs ...*Variant) *Variant{
  type_ := GoGetArg(goArgs, 0, MkUndefined()); _ = type_;
  types := MkMap(&VarMap{
        "1":MkString("market") ,
        "2":MkString("limit") ,
        "3":MkString("stop") ,
        "4":MkString("stop limit") ,
        }); _ = types;
  return this.SafeString(types , type_ , type_ );
}

func (this *Eqonex) ParseTimeInForce(goArgs ...*Variant) *Variant{
  timeInForce := GoGetArg(goArgs, 0, MkUndefined()); _ = timeInForce;
  timeInForces := MkMap(&VarMap{
        "1":MkString("GTC") ,
        "3":MkString("IOC") ,
        "4":MkString("FOK") ,
        "5":MkString("GTX") ,
        "6":MkString("GTD") ,
        }); _ = timeInForces;
  return this.SafeString(timeInForces , timeInForce , timeInForce );
}

func (this *Eqonex) ToMilliseconds(goArgs ...*Variant) *Variant{
  dateString := GoGetArg(goArgs, 0, MkUndefined()); _ = dateString;
  if IsTrue(OpEq2(dateString , MkUndefined() )) {
    return dateString ;
  }
  splits := dateString.Split(MkString("-") ); _ = splits;
  partOne := this.SafeString(splits , MkInteger(0) ); _ = partOne;
  partTwo := this.SafeString(splits , MkInteger(1) ); _ = partTwo;
  if IsTrue(OpOr(OpEq2(partOne , MkUndefined() ), OpEq2(partTwo , MkUndefined() ))) {
    return MkUndefined() ;
  }
  if IsTrue(OpNotEq2(partOne.Length , MkInteger(8) )) {
    return MkUndefined() ;
  }
  date := OpAdd(partOne.Slice(MkInteger(0) , MkInteger(4) ), OpAdd(MkString("-") , OpAdd(partOne.Slice(MkInteger(4) , MkInteger(6) ), OpAdd(MkString("-") , partOne.Slice(MkInteger(6) , MkInteger(8) ))))); _ = date;
  return this.Parse8601(OpAdd(date , OpAdd(MkString(" ") , partTwo )));
}

func (this *Eqonex) ConvertFromScale(goArgs ...*Variant) *Variant{
  number := GoGetArg(goArgs, 0, MkUndefined()); _ = number;
  scale := GoGetArg(goArgs, 1, MkUndefined()); _ = scale;
  if IsTrue(OpOr(OpEq2(number , MkUndefined() ), OpEq2(scale , MkUndefined() ))) {
    return MkUndefined() ;
  }
  precise := NewPrecise(number ); _ = precise;
  precise.Decimals = OpAdd(precise.Decimals , scale );
  precise.Reduce();
  return precise.ToString();
}

func (this *Eqonex) GetScale(goArgs ...*Variant) *Variant{
  num := GoGetArg(goArgs, 0, MkUndefined()); _ = num;
  s := this.NumberToString(num ); _ = s;
  return this.PrecisionFromString(s );
}

func (this *Eqonex) ConvertToScale(goArgs ...*Variant) *Variant{
  number := GoGetArg(goArgs, 0, MkUndefined()); _ = number;
  scale := GoGetArg(goArgs, 1, MkUndefined()); _ = scale;
  if IsTrue(OpOr(OpEq2(number , MkUndefined() ), OpEq2(scale , MkUndefined() ))) {
    return MkUndefined() ;
  }
  precise := NewPrecise(number ); _ = precise;
  precise.Decimals = OpSub(precise.Decimals , scale );
  precise.Reduce();
  preciseString := precise.ToString(); _ = preciseString;
  return parseInt(preciseString );
}

func (this *Eqonex) Nonce(goArgs ...*Variant) *Variant{
  return this.Milliseconds();
}

func (this *Eqonex) HandleErrors(goArgs ...*Variant) *Variant{
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
  error := this.SafeString(response , MkString("error") ); _ = error;
  if IsTrue(OpNotEq2(error , MkUndefined() )) {
    feedback := OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , body )); _ = feedback;
    this.ThrowExactlyMatchedException(*this.At(MkString("exceptions")) , error , feedback );
    this.ThrowBroadlyMatchedException(*this.At(MkString("exceptions")) , body , feedback );
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , body ))));
  }
  return MkUndefined()
}

func (this *Eqonex) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("public") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  url := this.ImplodeParams(path , params ); _ = url;
  query := this.Omit(params , this.ExtractParams(path )); _ = query;
  if IsTrue(OpEq2(api , MkString("public") )) {
    if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
      OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(query )));
    }
  } else {
    if IsTrue(OpEq2(api , MkString("private") )) {
      format := this.SafeValue(params , MkString("format") ); _ = format;
      type_ := this.SafeValue(params , MkString("type") ); _ = type_;
      extension := MkMap(&VarMap{}); _ = extension;
      if IsTrue(OpNotEq2(format , MkUndefined() )) {
        *(extension ).At(MkString("format") )= OpCopy(format );
      }
      if IsTrue(OpNotEq2(type_ , MkUndefined() )) {
        *(extension ).At(MkString("type") )= OpCopy(type_ );
      }
      if IsTrue(*(GoGetKeys(extension )).At(MkString("length") )) {
        OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(extension )));
      }
      params = this.Omit(params , MkArray(&VarArray{
              MkString("format") ,
              MkString("type") ,
              }));
      this.CheckRequiredCredentials();
      nonce := this.Nonce(); _ = nonce;
      query = this.Extend(query , MkMap(&VarMap{
              "userId":*this.At(MkString("uid")) ,
              "nonce":nonce ,
              }));
      *(params ).At(MkString("nonce") )= this.Nonce();
      body = this.Json(query );
      signature := this.Hmac(this.Encode(body ), this.Encode(*this.At(MkString("secret")) ), MkString("sha384") ); _ = signature;
      headers = MkMap(&VarMap{
          "Content-Type":MkString("application/json") ,
          "requestToken":*this.At(MkString("apiKey")) ,
          "signature":signature ,
          });
    }
  }
  url = OpAdd(*(*(*this.At(MkString("urls")) ).At(MkString("api") )).At(api ), OpAdd(MkString("/") , url ));
  return MkMap(&VarMap{
        "url":url ,
        "method":method ,
        "body":body ,
        "headers":headers ,
        });
}

