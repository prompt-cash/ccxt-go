package ccxt

import (
)

type Lykke struct {
	*ExchangeBase
}

var _ Exchange = (*Lykke)(nil)

func init() {
	exchange := &Lykke{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Lykke) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("lykke") ,
            "name":MkString("Lykke") ,
            "countries":MkArray(&VarArray{
                MkString("CH") ,
                }),
            "version":MkString("v1") ,
            "rateLimit":MkInteger(200) ,
            "has":MkMap(&VarMap{
                "CORS":MkBool(false) ,
                "fetchOHLCV":MkBool(false) ,
                "fetchOpenOrders":MkBool(true) ,
                "fetchClosedOrders":MkBool(true) ,
                "fetchOrder":MkBool(true) ,
                "fetchOrders":MkBool(true) ,
                "fetchTrades":MkBool(true) ,
                "fetchMyTrades":MkBool(true) ,
                "createOrder":MkBool(true) ,
                "cancelOrder":MkBool(true) ,
                "cancelAllOrders":MkBool(true) ,
                "fetchBalance":MkBool(true) ,
                "fetchMarkets":MkBool(true) ,
                "fetchOrderBook":MkBool(true) ,
                "fetchTicker":MkBool(true) ,
                }),
            "timeframes":MkMap(&VarMap{
                "1m":MkString("Minute") ,
                "5m":MkString("Min5") ,
                "15m":MkString("Min15") ,
                "30m":MkString("Min30") ,
                "1h":MkString("Hour") ,
                "4h":MkString("Hour4") ,
                "6h":MkString("Hour6") ,
                "12h":MkString("Hour12") ,
                "1d":MkString("Day") ,
                "1w":MkString("Week") ,
                "1M":MkString("Month") ,
                }),
            "requiredCredentials":MkMap(&VarMap{
                "apiKey":MkBool(true) ,
                "secret":MkBool(false) ,
                }),
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/1294454/34487620-3139a7b0-efe6-11e7-90f5-e520cef74451.jpg") ,
                "api":MkMap(&VarMap{
                    "mobile":MkString("https://public-api.lykke.com/api") ,
                    "public":MkString("https://hft-api.lykke.com/api") ,
                    "private":MkString("https://hft-api.lykke.com/api") ,
                    }),
                "test":MkMap(&VarMap{
                    "mobile":MkString("https://public-api.lykke.com/api") ,
                    "public":MkString("https://hft-service-dev.lykkex.net/api") ,
                    "private":MkString("https://hft-service-dev.lykkex.net/api") ,
                    }),
                "www":MkString("https://www.lykke.com") ,
                "doc":MkArray(&VarArray{
                    MkString("https://hft-api.lykke.com/swagger/ui/") ,
                    MkString("https://www.lykke.com/lykke_api") ,
                    }),
                "fees":MkString("https://www.lykke.com/trading-conditions") ,
                }),
            "api":MkMap(&VarMap{
                "mobile":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("AssetPairs/rate") ,
                        MkString("AssetPairs/rate/{assetPairId}") ,
                        MkString("AssetPairs/dictionary/{market}") ,
                        MkString("Assets/dictionary") ,
                        MkString("Candles/history/{market}/available") ,
                        MkString("Candles/history/{market}/{assetPair}/{period}/{type}/{from}/{to}") ,
                        MkString("Company/ownershipStructure") ,
                        MkString("Company/registrationsCount") ,
                        MkString("IsAlive") ,
                        MkString("Market") ,
                        MkString("Market/{market}") ,
                        MkString("Market/capitalization/{market}") ,
                        MkString("OrderBook") ,
                        MkString("OrderBook/{assetPairId}") ,
                        MkString("Trades/{AssetPairId}") ,
                        MkString("Trades/Last/{assetPair}/{n}") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("AssetPairs/rate/history") ,
                        MkString("AssetPairs/rate/history/{assetPairId}") ,
                        }),
                    }),
                "public":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("AssetPairs") ,
                        MkString("AssetPairs/{id}") ,
                        MkString("IsAlive") ,
                        MkString("OrderBooks") ,
                        MkString("OrderBooks/{AssetPairId}") ,
                        })}),
                "private":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("Orders") ,
                        MkString("Orders/{id}") ,
                        MkString("Wallets") ,
                        MkString("History/trades") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("Orders/limit") ,
                        MkString("Orders/market") ,
                        MkString("Orders/{id}/Cancel") ,
                        MkString("Orders/v2/market") ,
                        MkString("Orders/v2/limit") ,
                        MkString("Orders/stoplimit") ,
                        MkString("Orders/bulk") ,
                        }),
                    "delete":MkArray(&VarArray{
                        MkString("Orders") ,
                        MkString("Orders/{id}") ,
                        }),
                    }),
                }),
            "fees":MkMap(&VarMap{
                "trading":MkMap(&VarMap{
                    "tierBased":MkBool(false) ,
                    "percentage":MkBool(true) ,
                    "maker":MkNumber(0.0) ,
                    "taker":MkNumber(0.0) ,
                    }),
                "funding":MkMap(&VarMap{
                    "tierBased":MkBool(false) ,
                    "percentage":MkBool(false) ,
                    "withdraw":MkMap(&VarMap{"BTC":MkNumber(0.001) }),
                    "deposit":MkMap(&VarMap{"BTC":MkInteger(0) }),
                    }),
                }),
            "commonCurrencies":MkMap(&VarMap{
                "CAN":MkString("CanYaCoin") ,
                "XPD":MkString("Lykke XPD") ,
                }),
            }));
}

func (this *Lykke) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined()); _ = market;
  marketId := this.SafeString(trade , MkString("AssetPairId") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market ); _ = symbol;
  id := this.SafeString2(trade , MkString("id") , MkString("Id") ); _ = id;
  orderId := this.SafeString(trade , MkString("OrderId") ); _ = orderId;
  timestamp := this.Parse8601(this.SafeString2(trade , MkString("dateTime") , MkString("DateTime") )); _ = timestamp;
  priceString := this.SafeString2(trade , MkString("price") , MkString("Price") ); _ = priceString;
  amountString := this.SafeString2(trade , MkString("volume") , MkString("Amount") ); _ = amountString;
  side := this.SafeStringLower(trade , MkString("action") ); _ = side;
  if IsTrue(OpEq2(side , MkUndefined() )) {
    side = OpTrinary(OpEq2(*(amountString ).At(MkInteger(0) ), MkString("-") ), MkString("sell") , MkString("buy") );
  }
  amountString = Precise.StringAbs(amountString );
  price := this.ParseNumber(priceString ); _ = price;
  amount := this.ParseNumber(amountString ); _ = amount;
  cost := this.ParseNumber(Precise.StringMul(priceString , amountString )); _ = cost;
  fee := MkMap(&VarMap{
        "cost":MkInteger(0) ,
        "currency":*(market ).At(MkString("quote") ),
        }); _ = fee;
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

func (this *Lykke) FetchTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  if IsTrue(OpEq2(limit , MkUndefined() )) {
    limit = MkInteger(100) ;
  }
  request := MkMap(&VarMap{
        "AssetPairId":*(market ).At(MkString("id") ),
        "skip":MkInteger(0) ,
        "take":limit ,
        }); _ = request;
  response := this.Call(MkString("mobileGetTradesAssetPairId"), this.Extend(request , params )); _ = response;
  return this.ParseTrades(response , market , since , limit );
}

func (this *Lykke) FetchMyTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("take") )= OpCopy(limit );
  }
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("assetPairId") )= *(market ).At(MkString("id") );
  }
  response := this.Call(MkString("privateGetHistoryTrades"), this.Extend(request , params )); _ = response;
  return this.ParseTrades(response , market , since , limit );
}

func (this *Lykke) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("privateGetWallets"), params ); _ = response;
  result := MkMap(&VarMap{"info":response }); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , response.Length )); OpIncr(& i ){
    balance := *(response ).At(i ); _ = balance;
    currencyId := this.SafeString(balance , MkString("AssetId") ); _ = currencyId;
    code := this.SafeCurrencyCode(currencyId ); _ = code;
    account := this.Account(); _ = account;
    *(account ).At(MkString("total") )= this.SafeString(balance , MkString("Balance") );
    *(account ).At(MkString("used") )= this.SafeString(balance , MkString("Reserved") );
    *(result ).At(code )= OpCopy(account );
  }
  return this.ParseBalance(result );
}

func (this *Lykke) CancelOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"id":id }); _ = request;
  return this.Call(MkString("privateDeleteOrdersId"), this.Extend(request , params ));
}

func (this *Lykke) CancelAllOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("assetPairId") )= *(market ).At(MkString("id") );
  }
  return this.Call(MkString("privateDeleteOrders"), this.Extend(request , params ));
}

func (this *Lykke) CreateOrder(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 2, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 3, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 4, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 5, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  query := MkMap(&VarMap{
        "AssetPairId":*(market ).At(MkString("id") ),
        "OrderAction":this.Capitalize(side ),
        "Volume":amount ,
        "Asset":*(market ).At(MkString("baseId") ),
        }); _ = query;
  if IsTrue(OpEq2(type_ , MkString("limit") )) {
    *(query ).At(MkString("Price") )= OpCopy(price );
  }
  method := OpAdd(MkString("privatePostOrdersV2") , this.Capitalize(type_ )); _ = method;
  result := this.Call(method , this.Extend(query , params )); _ = result;
  id := this.SafeString(result , MkString("Id") ); _ = id;
  price = this.SafeNumber(result , MkString("Price") );
  return MkMap(&VarMap{
        "id":id ,
        "info":result ,
        "clientOrderId":MkUndefined() ,
        "timestamp":MkUndefined() ,
        "datetime":MkUndefined() ,
        "lastTradeTimestamp":MkUndefined() ,
        "symbol":symbol ,
        "type":type_ ,
        "side":side ,
        "price":price ,
        "amount":amount ,
        "cost":MkUndefined() ,
        "average":MkUndefined() ,
        "filled":MkUndefined() ,
        "remaining":MkUndefined() ,
        "status":MkUndefined() ,
        "fee":MkUndefined() ,
        "trades":MkUndefined() ,
        });
}

func (this *Lykke) FetchMarkets(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  markets := this.Call(MkString("publicGetAssetPairs")); _ = markets;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , markets.Length )); OpIncr(& i ){
    market := *(markets ).At(i ); _ = market;
    id := this.SafeString(market , MkString("Id") ); _ = id;
    name := this.SafeString(market , MkString("Name") ); _ = name;
    baseId := MkUndefined(); _ = baseId;
    quoteId := MkUndefined(); _ = quoteId;
    ArrayUnpack(name.Split(MkString("/") ), &baseId, &quoteId);
    base := this.SafeCurrencyCode(baseId ); _ = base;
    quote := this.SafeCurrencyCode(quoteId ); _ = quote;
    symbol := OpAdd(base , OpAdd(MkString("/") , quote )); _ = symbol;
    pricePrecision := this.SafeString(market , MkString("Accuracy") ); _ = pricePrecision;
    priceLimit := this.ParsePrecision(pricePrecision ); _ = priceLimit;
    precision := MkMap(&VarMap{
          "price":parseInt(pricePrecision ),
          "amount":this.SafeInteger(market , MkString("InvertedAccuracy") ),
          }); _ = precision;
    result.Push(MkMap(&VarMap{
            "id":id ,
            "symbol":symbol ,
            "base":base ,
            "quote":quote ,
            "active":MkBool(true) ,
            "info":market ,
            "precision":precision ,
            "limits":MkMap(&VarMap{
                "amount":MkMap(&VarMap{
                    "min":this.SafeNumber(market , MkString("MinVolume") ),
                    "max":MkUndefined() ,
                    }),
                "price":MkMap(&VarMap{
                    "min":this.ParseNumber(priceLimit ),
                    "max":MkUndefined() ,
                    }),
                "cost":MkMap(&VarMap{
                    "min":this.SafeNumber(market , MkString("MinInvertedVolume") ),
                    "max":MkUndefined() ,
                    }),
                }),
            "baseId":MkUndefined() ,
            "quoteId":MkUndefined() ,
            }));
  }
  return result ;
}

func (this *Lykke) ParseTicker(goArgs ...*Variant) *Variant{
  ticker := GoGetArg(goArgs, 0, MkUndefined()); _ = ticker;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.Milliseconds(); _ = timestamp;
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  if IsTrue(market ) {
    symbol = *(market ).At(MkString("symbol") );
  }
  close := this.SafeNumber(ticker , MkString("lastPrice") ); _ = close;
  return MkMap(&VarMap{
        "symbol":symbol ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "high":MkUndefined() ,
        "low":MkUndefined() ,
        "bid":this.SafeNumber(ticker , MkString("bid") ),
        "bidVolume":MkUndefined() ,
        "ask":this.SafeNumber(ticker , MkString("ask") ),
        "askVolume":MkUndefined() ,
        "vwap":MkUndefined() ,
        "open":MkUndefined() ,
        "close":close ,
        "last":close ,
        "previousClose":MkUndefined() ,
        "change":MkUndefined() ,
        "percentage":MkUndefined() ,
        "average":MkUndefined() ,
        "baseVolume":MkUndefined() ,
        "quoteVolume":this.SafeNumber(ticker , MkString("volume24H") ),
        "info":ticker ,
        });
}

func (this *Lykke) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"market":*(market ).At(MkString("id") )}); _ = request;
  ticker := this.Call(MkString("mobileGetMarketMarket"), this.Extend(request , params )); _ = ticker;
  return this.ParseTicker(ticker , market );
}

func (this *Lykke) ParseOrderStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "Open":MkString("open") ,
        "Pending":MkString("open") ,
        "InOrderBook":MkString("open") ,
        "Processing":MkString("open") ,
        "Matched":MkString("closed") ,
        "Cancelled":MkString("canceled") ,
        "Rejected":MkString("rejected") ,
        "Replaced":MkString("canceled") ,
        "Placed":MkString("open") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Lykke) ParseOrder(goArgs ...*Variant) *Variant{
  order := GoGetArg(goArgs, 0, MkUndefined()); _ = order;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  status := this.ParseOrderStatus(this.SafeString(order , MkString("Status") )); _ = status;
  marketId := this.SafeString(order , MkString("AssetPairId") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market ); _ = symbol;
  lastTradeTimestamp := this.Parse8601(this.SafeString(order , MkString("LastMatchTime") )); _ = lastTradeTimestamp;
  timestamp := OpCopy(MkUndefined() ); _ = timestamp;
  if IsTrue(OpAnd(OpHasMember(MkString("Registered") , order ), *(order ).At(MkString("Registered") ))) {
    timestamp = this.Parse8601(*(order ).At(MkString("Registered") ));
  } else {
    if IsTrue(OpAnd(OpHasMember(MkString("CreatedAt") , order ), *(order ).At(MkString("CreatedAt") ))) {
      timestamp = this.Parse8601(*(order ).At(MkString("CreatedAt") ));
    }
  }
  price := this.SafeNumber(order , MkString("Price") ); _ = price;
  side := OpCopy(MkUndefined() ); _ = side;
  amount := this.SafeNumber(order , MkString("Volume") ); _ = amount;
  if IsTrue(OpLw(amount , MkInteger(0) )) {
    side = MkString("sell") ;
    amount = Math.Abs(amount );
  } else {
    side = MkString("buy") ;
  }
  remaining := Math.Abs(this.SafeNumber(order , MkString("RemainingVolume") )); _ = remaining;
  id := this.SafeString(order , MkString("Id") ); _ = id;
  return this.SafeOrder(MkMap(&VarMap{
            "info":order ,
            "id":id ,
            "clientOrderId":MkUndefined() ,
            "timestamp":timestamp ,
            "datetime":this.Iso8601(timestamp ),
            "lastTradeTimestamp":lastTradeTimestamp ,
            "symbol":symbol ,
            "type":MkUndefined() ,
            "timeInForce":MkUndefined() ,
            "postOnly":MkUndefined() ,
            "side":side ,
            "price":price ,
            "stopPrice":MkUndefined() ,
            "cost":MkUndefined() ,
            "average":MkUndefined() ,
            "amount":amount ,
            "filled":MkUndefined() ,
            "remaining":remaining ,
            "status":status ,
            "fee":MkUndefined() ,
            "trades":MkUndefined() ,
            }));
}

func (this *Lykke) FetchOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"id":id }); _ = request;
  response := this.Call(MkString("privateGetOrdersId"), this.Extend(request , params )); _ = response;
  return this.ParseOrder(response );
}

func (this *Lykke) FetchOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("privateGetOrders"), params ); _ = response;
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
  }
  return this.ParseOrders(response , market , since , limit );
}

func (this *Lykke) FetchOpenOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"status":MkString("InOrderBook") }); _ = request;
  return this.FetchOrders(symbol , since , limit , this.Extend(request , params ));
}

func (this *Lykke) FetchClosedOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"status":MkString("Matched") }); _ = request;
  return this.FetchOrders(symbol , since , limit , this.Extend(request , params ));
}

func (this *Lykke) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("publicGetOrderBooksAssetPairId"), this.Extend(MkMap(&VarMap{"AssetPairId":this.MarketId(symbol )}), params )); _ = response;
  orderbook := MkMap(&VarMap{
        "timestamp":MkUndefined() ,
        "bids":MkArray(&VarArray{}),
        "asks":MkArray(&VarArray{}),
        }); _ = orderbook;
  timestamp := OpCopy(MkUndefined() ); _ = timestamp;
  for i := MkInteger(0) ; IsTrue(OpLw(i , response.Length )); OpIncr(& i ){
    side := *(response ).At(i ); _ = side;
    if IsTrue(*(side ).At(MkString("IsBuy") )) {
      *(orderbook ).At(MkString("bids") )= this.ArrayConcat(*(orderbook ).At(MkString("bids") ), *(side ).At(MkString("Prices") ));
    } else {
      *(orderbook ).At(MkString("asks") )= this.ArrayConcat(*(orderbook ).At(MkString("asks") ), *(side ).At(MkString("Prices") ));
    }
    sideTimestamp := this.Parse8601(*(side ).At(MkString("Timestamp") )); _ = sideTimestamp;
    timestamp = OpTrinary(OpEq2(timestamp , MkUndefined() ), sideTimestamp , Math.Max(timestamp , sideTimestamp ));
  }
  return this.ParseOrderBook(orderbook , symbol , timestamp , MkString("bids") , MkString("asks") , MkString("Price") , MkString("Volume") );
}

func (this *Lykke) ParseBidAsk(goArgs ...*Variant) *Variant{
  bidask := GoGetArg(goArgs, 0, MkUndefined()); _ = bidask;
  priceKey := GoGetArg(goArgs, 1, MkInteger(0) ); _ = priceKey;
  amountKey := GoGetArg(goArgs, 2, MkInteger(1) ); _ = amountKey;
  price := this.SafeNumber(bidask , priceKey ); _ = price;
  amount := this.SafeNumber(bidask , amountKey ); _ = amount;
  if IsTrue(OpLw(amount , MkInteger(0) )) {
    amount = OpNeg(amount );
  }
  return MkArray(&VarArray{
        price ,
        amount ,
        });
}

func (this *Lykke) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("public") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  url := OpAdd(*(*(*this.At(MkString("urls")) ).At(MkString("api") )).At(api ), OpAdd(MkString("/") , this.ImplodeParams(path , params ))); _ = url;
  query := this.Omit(params , this.ExtractParams(path )); _ = query;
  if IsTrue(OpEq2(api , MkString("mobile") )) {
    if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
      OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(query )));
    }
  } else {
    if IsTrue(OpEq2(api , MkString("public") )) {
      if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
        OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(query )));
      }
    } else {
      if IsTrue(OpEq2(api , MkString("private") )) {
        if IsTrue(OpOr(OpEq2(method , MkString("GET") ), OpEq2(method , MkString("DELETE") ))) {
          if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
            OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(query )));
          }
        }
        this.CheckRequiredCredentials();
        headers = MkMap(&VarMap{
            "api-key":*this.At(MkString("apiKey")) ,
            "Accept":MkString("application/json") ,
            "Content-Type":MkString("application/json") ,
            });
        if IsTrue(OpEq2(method , MkString("POST") )) {
          if IsTrue(*(GoGetKeys(params )).At(MkString("length") )) {
            body = this.Json(params );
          }
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

