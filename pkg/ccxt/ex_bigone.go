package ccxt

import (
)

type Bigone struct {
	*ExchangeBase
}

var _ Exchange = (*Bigone)(nil)

func init() {
	exchange := &Bigone{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Bigone) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("bigone") ,
            "name":MkString("BigONE") ,
            "countries":MkArray(&VarArray{
                MkString("CN") ,
                }),
            "version":MkString("v3") ,
            "rateLimit":MkInteger(1200) ,
            "has":MkMap(&VarMap{
                "cancelAllOrders":MkBool(true) ,
                "cancelOrder":MkBool(true) ,
                "createOrder":MkBool(true) ,
                "fetchBalance":MkBool(true) ,
                "fetchClosedOrders":MkBool(true) ,
                "fetchDepositAddress":MkBool(true) ,
                "fetchDeposits":MkBool(true) ,
                "fetchMarkets":MkBool(true) ,
                "fetchMyTrades":MkBool(true) ,
                "fetchOHLCV":MkBool(true) ,
                "fetchOpenOrders":MkBool(true) ,
                "fetchOrder":MkBool(true) ,
                "fetchOrders":MkBool(true) ,
                "fetchOrderBook":MkBool(true) ,
                "fetchTicker":MkBool(true) ,
                "fetchTickers":MkBool(true) ,
                "fetchTime":MkBool(true) ,
                "fetchTrades":MkBool(true) ,
                "fetchWithdrawals":MkBool(true) ,
                "withdraw":MkBool(true) ,
                }),
            "timeframes":MkMap(&VarMap{
                "1m":MkString("min1") ,
                "5m":MkString("min5") ,
                "15m":MkString("min15") ,
                "30m":MkString("min30") ,
                "1h":MkString("hour1") ,
                "3h":MkString("hour3") ,
                "4h":MkString("hour4") ,
                "6h":MkString("hour6") ,
                "12h":MkString("hour12") ,
                "1d":MkString("day1") ,
                "1w":MkString("week1") ,
                "1M":MkString("month1") ,
                }),
            "hostname":MkString("big.one") ,
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/1294454/69354403-1d532180-0c91-11ea-88ed-44c06cefdf87.jpg") ,
                "api":MkMap(&VarMap{
                    "public":MkString("https://{hostname}/api/v3") ,
                    "private":MkString("https://{hostname}/api/v3/viewer") ,
                    }),
                "www":MkString("https://big.one") ,
                "doc":MkString("https://open.big.one/docs/api.html") ,
                "fees":MkString("https://bigone.zendesk.com/hc/en-us/articles/115001933374-BigONE-Fee-Policy") ,
                "referral":MkString("https://b1.run/users/new?code=D3LLBVFT") ,
                }),
            "api":MkMap(&VarMap{
                "public":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("ping") ,
                        MkString("asset_pairs") ,
                        MkString("asset_pairs/{asset_pair_name}/depth") ,
                        MkString("asset_pairs/{asset_pair_name}/trades") ,
                        MkString("asset_pairs/{asset_pair_name}/ticker") ,
                        MkString("asset_pairs/{asset_pair_name}/candles") ,
                        MkString("asset_pairs/tickers") ,
                        })}),
                "private":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("accounts") ,
                        MkString("fund/accounts") ,
                        MkString("assets/{asset_symbol}/address") ,
                        MkString("orders") ,
                        MkString("orders/{id}") ,
                        MkString("orders/multi") ,
                        MkString("trades") ,
                        MkString("withdrawals") ,
                        MkString("deposits") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("orders") ,
                        MkString("orders/{id}/cancel") ,
                        MkString("orders/cancel") ,
                        MkString("withdrawals") ,
                        MkString("transfer") ,
                        }),
                    }),
                }),
            "fees":MkMap(&VarMap{
                "trading":MkMap(&VarMap{
                    "maker":this.ParseNumber(MkString("0.001") ),
                    "taker":this.ParseNumber(MkString("0.001") ),
                    }),
                "funding":MkMap(&VarMap{"withdraw":MkMap(&VarMap{})}),
                }),
            "exceptions":MkMap(&VarMap{
                "exact":MkMap(&VarMap{
                    "10001":BadRequest ,
                    "10005":ExchangeError ,
                    "Amount's scale must greater than AssetPair's base scale":InvalidOrder ,
                    "Price mulit with amount should larger than AssetPair's min_quote_value":InvalidOrder ,
                    "10007":BadRequest ,
                    "10011":ExchangeError ,
                    "10013":OrderNotFound ,
                    "10014":InsufficientFunds ,
                    "10403":PermissionDenied ,
                    "10429":RateLimitExceeded ,
                    "40004":AuthenticationError ,
                    "40103":AuthenticationError ,
                    "40104":AuthenticationError ,
                    "40301":PermissionDenied ,
                    "40302":ExchangeError ,
                    "40601":ExchangeError ,
                    "40602":ExchangeError ,
                    "40603":InsufficientFunds ,
                    "40605":InvalidOrder ,
                    "40120":InvalidOrder ,
                    "40121":InvalidOrder ,
                    "60100":BadSymbol ,
                    }),
                "broad":MkMap(&VarMap{}),
                }),
            "commonCurrencies":MkMap(&VarMap{
                "CRE":MkString("Cybereits") ,
                "FXT":MkString("FXTTOKEN") ,
                "MBN":MkString("Mobilian Coin") ,
                "ONE":MkString("BigONE Token") ,
                }),
            }));
}

func (this *Bigone) FetchMarkets(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetAssetPairs"), params ); _ = response;
  markets := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = markets;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , markets.Length )); OpIncr(& i ){
    market := *(markets ).At(i ); _ = market;
    id := this.SafeString(market , MkString("name") ); _ = id;
    uuid := this.SafeString(market , MkString("id") ); _ = uuid;
    baseAsset := this.SafeValue(market , MkString("base_asset") , MkMap(&VarMap{})); _ = baseAsset;
    quoteAsset := this.SafeValue(market , MkString("quote_asset") , MkMap(&VarMap{})); _ = quoteAsset;
    baseId := this.SafeString(baseAsset , MkString("symbol") ); _ = baseId;
    quoteId := this.SafeString(quoteAsset , MkString("symbol") ); _ = quoteId;
    base := this.SafeCurrencyCode(baseId ); _ = base;
    quote := this.SafeCurrencyCode(quoteId ); _ = quote;
    symbol := OpAdd(base , OpAdd(MkString("/") , quote )); _ = symbol;
    amountPrecisionString := this.SafeString(market , MkString("base_scale") ); _ = amountPrecisionString;
    pricePrecisionString := this.SafeString(market , MkString("quote_scale") ); _ = pricePrecisionString;
    amountLimit := this.ParsePrecision(amountPrecisionString ); _ = amountLimit;
    priceLimit := this.ParsePrecision(pricePrecisionString ); _ = priceLimit;
    precision := MkMap(&VarMap{
          "amount":parseInt(amountPrecisionString ),
          "price":parseInt(pricePrecisionString ),
          }); _ = precision;
    minCost := this.SafeNumber(market , MkString("min_quote_value") ); _ = minCost;
    entry := MkMap(&VarMap{
          "id":id ,
          "uuid":uuid ,
          "symbol":symbol ,
          "base":base ,
          "quote":quote ,
          "baseId":baseId ,
          "quoteId":quoteId ,
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
                  "min":minCost ,
                  "max":MkUndefined() ,
                  }),
              }),
          "info":market ,
          }); _ = entry;
    result.Push(entry );
  }
  return result ;
}

func (this *Bigone) LoadMarkets(goArgs ...*Variant) *Variant{
  reload := GoGetArg(goArgs, 0, MkBool(false) ); _ = reload;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  markets := this.LoadMarkets(reload , params ); _ = markets;
  marketsByUuid := this.SafeValue(*this.At(MkString("options")) , MkString("marketsByUuid") ); _ = marketsByUuid;
  if IsTrue(OpOr(OpEq2(marketsByUuid , MkUndefined() ), reload )) {
    marketsByUuid = MkMap(&VarMap{});
    for i := MkInteger(0) ; IsTrue(OpLw(i , (*((this).At(MkString("symbols")))).Length )); OpIncr(& i ){
      symbol := *(*this.At(MkString("symbols")) ).At(i ); _ = symbol;
      market := *(*this.At(MkString("markets")) ).At(symbol ); _ = market;
      uuid := this.SafeString(market , MkString("uuid") ); _ = uuid;
      *(marketsByUuid ).At(uuid )= OpCopy(market );
    }
    *(*this.At(MkString("options")) ).At(MkString("marketsByUuid") )= OpCopy(marketsByUuid );
  }
  return markets ;
}

func (this *Bigone) ParseTicker(goArgs ...*Variant) *Variant{
  ticker := GoGetArg(goArgs, 0, MkUndefined()); _ = ticker;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  marketId := this.SafeString(ticker , MkString("asset_pair_name") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market , MkString("-") ); _ = symbol;
  timestamp := OpCopy(MkUndefined() ); _ = timestamp;
  close := this.SafeNumber(ticker , MkString("close") ); _ = close;
  bid := this.SafeValue(ticker , MkString("bid") , MkMap(&VarMap{})); _ = bid;
  ask := this.SafeValue(ticker , MkString("ask") , MkMap(&VarMap{})); _ = ask;
  return MkMap(&VarMap{
        "symbol":symbol ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "high":this.SafeNumber(ticker , MkString("high") ),
        "low":this.SafeNumber(ticker , MkString("low") ),
        "bid":this.SafeNumber(bid , MkString("price") ),
        "bidVolume":this.SafeNumber(bid , MkString("quantity") ),
        "ask":this.SafeNumber(ask , MkString("price") ),
        "askVolume":this.SafeNumber(ask , MkString("quantity") ),
        "vwap":MkUndefined() ,
        "open":this.SafeNumber(ticker , MkString("open") ),
        "close":close ,
        "last":close ,
        "previousClose":MkUndefined() ,
        "change":this.SafeNumber(ticker , MkString("daily_change") ),
        "percentage":MkUndefined() ,
        "average":MkUndefined() ,
        "baseVolume":this.SafeNumber(ticker , MkString("volume") ),
        "quoteVolume":MkUndefined() ,
        "info":ticker ,
        });
}

func (this *Bigone) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"asset_pair_name":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("publicGetAssetPairsAssetPairNameTicker"), this.Extend(request , params )); _ = response;
  ticker := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = ticker;
  return this.ParseTicker(ticker , market );
}

func (this *Bigone) FetchTickers(goArgs ...*Variant) *Variant{
  symbols := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbols;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  if IsTrue(OpNotEq2(symbols , MkUndefined() )) {
    ids := this.MarketIds(symbols ); _ = ids;
    *(request ).At(MkString("pair_names") )= ids.Join(MkString(",") );
  }
  response := this.Call(MkString("publicGetAssetPairsTickers"), this.Extend(request , params )); _ = response;
  tickers := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = tickers;
  result := MkMap(&VarMap{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , tickers.Length )); OpIncr(& i ){
    ticker := this.ParseTicker(*(tickers ).At(i )); _ = ticker;
    symbol := *(ticker ).At(MkString("symbol") ); _ = symbol;
    *(result ).At(symbol )= OpCopy(ticker );
  }
  return this.FilterByArray(result , MkString("symbol") , symbols );
}

func (this *Bigone) FetchTime(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetPing"), params ); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  timestamp := this.SafeInteger(data , MkString("timestamp") ); _ = timestamp;
  return parseInt(OpDiv(timestamp , MkInteger(1000000) ));
}

func (this *Bigone) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"asset_pair_name":*(market ).At(MkString("id") )}); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("publicGetAssetPairsAssetPairNameDepth"), this.Extend(request , params )); _ = response;
  orderbook := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = orderbook;
  return this.ParseOrderBook(orderbook , symbol , MkUndefined() , MkString("bids") , MkString("asks") , MkString("price") , MkString("quantity") );
}

func (this *Bigone) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.Parse8601(this.SafeString2(trade , MkString("created_at") , MkString("inserted_at") )); _ = timestamp;
  priceString := this.SafeString(trade , MkString("price") ); _ = priceString;
  amountString := this.SafeString(trade , MkString("amount") ); _ = amountString;
  price := this.ParseNumber(priceString ); _ = price;
  amount := this.ParseNumber(amountString ); _ = amount;
  cost := this.ParseNumber(Precise.StringMul(priceString , amountString )); _ = cost;
  marketId := this.SafeString(trade , MkString("asset_pair_name") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market , MkString("-") ); _ = symbol;
  side := this.SafeString(trade , MkString("side") ); _ = side;
  takerSide := this.SafeString(trade , MkString("taker_side") ); _ = takerSide;
  takerOrMaker := OpCopy(MkUndefined() ); _ = takerOrMaker;
  if IsTrue(OpAnd(OpNotEq2(takerSide , MkUndefined() ), OpAnd(OpNotEq2(side , MkUndefined() ), OpNotEq2(side , MkString("SELF_TRADING") )))) {
    takerOrMaker = OpTrinary(OpEq2(takerSide , side ), MkString("taker") , MkString("maker") );
  }
  if IsTrue(OpEq2(side , MkUndefined() )) {
    side = OpTrinary(OpEq2(takerSide , MkString("ASK") ), MkString("sell") , MkString("buy") );
  } else {
    if IsTrue(OpEq2(side , MkString("BID") )) {
      side = MkString("buy") ;
    } else {
      if IsTrue(OpEq2(side , MkString("ASK") )) {
        side = MkString("sell") ;
      }
    }
  }
  makerOrderId := this.SafeString(trade , MkString("maker_order_id") ); _ = makerOrderId;
  takerOrderId := this.SafeString(trade , MkString("taker_order_id") ); _ = takerOrderId;
  orderId := OpCopy(MkUndefined() ); _ = orderId;
  if IsTrue(OpNotEq2(makerOrderId , MkUndefined() )) {
    if IsTrue(OpNotEq2(takerOrderId , MkUndefined() )) {
      orderId = MkArray(&VarArray{
          makerOrderId ,
          takerOrderId ,
          });
    } else {
      orderId = OpCopy(makerOrderId );
    }
  } else {
    if IsTrue(OpNotEq2(takerOrderId , MkUndefined() )) {
      orderId = OpCopy(takerOrderId );
    }
  }
  id := this.SafeString(trade , MkString("id") ); _ = id;
  result := MkMap(&VarMap{
        "id":id ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "symbol":symbol ,
        "order":orderId ,
        "type":MkString("limit") ,
        "side":side ,
        "takerOrMaker":takerOrMaker ,
        "price":price ,
        "amount":amount ,
        "cost":this.ParseNumber(cost ),
        "info":trade ,
        }); _ = result;
  makerCurrencyCode := OpCopy(MkUndefined() ); _ = makerCurrencyCode;
  takerCurrencyCode := OpCopy(MkUndefined() ); _ = takerCurrencyCode;
  if IsTrue(OpAnd(OpNotEq2(market , MkUndefined() ), OpNotEq2(takerOrMaker , MkUndefined() ))) {
    if IsTrue(OpEq2(side , MkString("buy") )) {
      if IsTrue(OpEq2(takerOrMaker , MkString("maker") )) {
        makerCurrencyCode = *(market ).At(MkString("base") );
        takerCurrencyCode = *(market ).At(MkString("quote") );
      } else {
        makerCurrencyCode = *(market ).At(MkString("quote") );
        takerCurrencyCode = *(market ).At(MkString("base") );
      }
    } else {
      if IsTrue(OpEq2(takerOrMaker , MkString("maker") )) {
        makerCurrencyCode = *(market ).At(MkString("quote") );
        takerCurrencyCode = *(market ).At(MkString("base") );
      } else {
        makerCurrencyCode = *(market ).At(MkString("base") );
        takerCurrencyCode = *(market ).At(MkString("quote") );
      }
    }
  } else {
    if IsTrue(OpEq2(side , MkString("SELF_TRADING") )) {
      if IsTrue(OpEq2(takerSide , MkString("BID") )) {
        makerCurrencyCode = *(market ).At(MkString("quote") );
        takerCurrencyCode = *(market ).At(MkString("base") );
      } else {
        if IsTrue(OpEq2(takerSide , MkString("ASK") )) {
          makerCurrencyCode = *(market ).At(MkString("base") );
          takerCurrencyCode = *(market ).At(MkString("quote") );
        }
      }
    }
  }
  makerFeeCost := this.SafeNumber(trade , MkString("maker_fee") ); _ = makerFeeCost;
  takerFeeCost := this.SafeNumber(trade , MkString("taker_fee") ); _ = takerFeeCost;
  if IsTrue(OpNotEq2(makerFeeCost , MkUndefined() )) {
    if IsTrue(OpNotEq2(takerFeeCost , MkUndefined() )) {
      *(result ).At(MkString("fees") )= MkArray(&VarArray{
          MkMap(&VarMap{
              "cost":makerFeeCost ,
              "currency":makerCurrencyCode ,
              }),
          MkMap(&VarMap{
              "cost":takerFeeCost ,
              "currency":takerCurrencyCode ,
              }),
          });
    } else {
      *(result ).At(MkString("fee") )= MkMap(&VarMap{
          "cost":makerFeeCost ,
          "currency":makerCurrencyCode ,
          });
    }
  } else {
    if IsTrue(OpNotEq2(takerFeeCost , MkUndefined() )) {
      *(result ).At(MkString("fee") )= MkMap(&VarMap{
          "cost":takerFeeCost ,
          "currency":takerCurrencyCode ,
          });
    } else {
      *(result ).At(MkString("fee") )= OpCopy(MkUndefined() );
    }
  }
  return result ;
}

func (this *Bigone) FetchTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"asset_pair_name":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("publicGetAssetPairsAssetPairNameTrades"), this.Extend(request , params )); _ = response;
  trades := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = trades;
  return this.ParseTrades(trades , market , since , limit );
}

func (this *Bigone) ParseOHLCV(goArgs ...*Variant) *Variant{
  ohlcv := GoGetArg(goArgs, 0, MkUndefined()); _ = ohlcv;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  return MkArray(&VarArray{
        this.Parse8601(this.SafeString(ohlcv , MkString("time") )),
        this.SafeNumber(ohlcv , MkString("open") ),
        this.SafeNumber(ohlcv , MkString("high") ),
        this.SafeNumber(ohlcv , MkString("low") ),
        this.SafeNumber(ohlcv , MkString("close") ),
        this.SafeNumber(ohlcv , MkString("volume") ),
        });
}

func (this *Bigone) FetchOHLCV(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  timeframe := GoGetArg(goArgs, 1, MkString("1m") ); _ = timeframe;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  if IsTrue(OpEq2(limit , MkUndefined() )) {
    limit = MkInteger(100) ;
  }
  request := MkMap(&VarMap{
        "asset_pair_name":*(market ).At(MkString("id") ),
        "period":*(*this.At(MkString("timeframes")) ).At(timeframe ),
        "limit":limit ,
        }); _ = request;
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    duration := this.ParseTimeframe(timeframe ); _ = duration;
    end := this.Sum(since , OpMulti(limit , OpMulti(duration , MkInteger(1000) ))); _ = end;
    *(request ).At(MkString("time") )= this.Iso8601(end );
  }
  response := this.Call(MkString("publicGetAssetPairsAssetPairNameCandles"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = data;
  return this.ParseOHLCVs(data , market , timeframe , since , limit );
}

func (this *Bigone) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  type_ := this.SafeString(params , MkString("type") , MkString("") ); _ = type_;
  params = this.Omit(params , MkString("type") );
  method := OpAdd(MkString("privateGet") , OpAdd(this.Capitalize(type_ ), MkString("Accounts") )); _ = method;
  response := this.Call(method , params ); _ = response;
  result := MkMap(&VarMap{
        "info":response ,
        "timestamp":MkUndefined() ,
        "datetime":MkUndefined() ,
        }); _ = result;
  balances := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = balances;
  for i := MkInteger(0) ; IsTrue(OpLw(i , balances.Length )); OpIncr(& i ){
    balance := *(balances ).At(i ); _ = balance;
    symbol := this.SafeString(balance , MkString("asset_symbol") ); _ = symbol;
    code := this.SafeCurrencyCode(symbol ); _ = code;
    account := this.Account(); _ = account;
    *(account ).At(MkString("total") )= this.SafeString(balance , MkString("balance") );
    *(account ).At(MkString("used") )= this.SafeString(balance , MkString("locked_balance") );
    *(result ).At(code )= OpCopy(account );
  }
  return this.ParseBalance(result );
}

func (this *Bigone) ParseOrder(goArgs ...*Variant) *Variant{
  order := GoGetArg(goArgs, 0, MkUndefined()); _ = order;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  id := this.SafeString(order , MkString("id") ); _ = id;
  marketId := this.SafeString(order , MkString("asset_pair_name") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market , MkString("-") ); _ = symbol;
  timestamp := this.Parse8601(this.SafeString(order , MkString("created_at") )); _ = timestamp;
  price := this.SafeNumber(order , MkString("price") ); _ = price;
  amount := this.SafeNumber(order , MkString("amount") ); _ = amount;
  filled := this.SafeNumber(order , MkString("filled_amount") ); _ = filled;
  status := this.ParseOrderStatus(this.SafeString(order , MkString("state") )); _ = status;
  side := this.SafeString(order , MkString("side") ); _ = side;
  if IsTrue(OpEq2(side , MkString("BID") )) {
    side = MkString("buy") ;
  } else {
    side = MkString("sell") ;
  }
  lastTradeTimestamp := this.Parse8601(this.SafeString(order , MkString("updated_at") )); _ = lastTradeTimestamp;
  average := this.SafeNumber(order , MkString("avg_deal_price") ); _ = average;
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

func (this *Bigone) CreateOrder(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 2, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 3, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 4, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 5, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  side = OpTrinary(OpEq2(side , MkString("buy") ), MkString("BID") , MkString("ASK") );
  uppercaseType := type_.ToUpperCase(); _ = uppercaseType;
  request := MkMap(&VarMap{
        "asset_pair_name":*(market ).At(MkString("id") ),
        "side":side ,
        "amount":this.AmountToPrecision(symbol , amount ),
        "type":uppercaseType ,
        }); _ = request;
  if IsTrue(OpEq2(uppercaseType , MkString("LIMIT") )) {
    *(request ).At(MkString("price") )= this.PriceToPrecision(symbol , price );
  } else {
    isStopLimit := OpEq2(uppercaseType , MkString("STOP_LIMIT") ); _ = isStopLimit;
    isStopMarket := OpEq2(uppercaseType , MkString("STOP_MARKET") ); _ = isStopMarket;
    if IsTrue(OpOr(isStopLimit , isStopMarket )) {
      stopPrice := this.SafeNumber2(params , MkString("stop_price") , MkString("stopPrice") ); _ = stopPrice;
      if IsTrue(OpEq2(stopPrice , MkUndefined() )) {
        panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" createOrder() requires a stop_price parameter") )));
      }
      *(request ).At(MkString("stop_price") )= this.PriceToPrecision(symbol , stopPrice );
      params = this.Omit(params , MkArray(&VarArray{
              MkString("stop_price") ,
              MkString("stopPrice") ,
              }));
    }
    if IsTrue(isStopLimit ) {
      *(request ).At(MkString("price") )= this.PriceToPrecision(symbol , price );
    }
  }
  response := this.Call(MkString("privatePostOrders"), this.Extend(request , params )); _ = response;
  order := this.SafeValue(response , MkString("data") ); _ = order;
  return this.ParseOrder(order , market );
}

func (this *Bigone) CancelOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"id":id }); _ = request;
  response := this.Call(MkString("privatePostOrdersIdCancel"), this.Extend(request , params )); _ = response;
  order := this.SafeValue(response , MkString("data") ); _ = order;
  return this.ParseOrder(order );
}

func (this *Bigone) CancelAllOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"asset_pair_name":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("privatePostOrdersCancel"), this.Extend(request , params )); _ = response;
  return response ;
}

func (this *Bigone) FetchOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"id":id }); _ = request;
  response := this.Call(MkString("privateGetOrdersId"), this.Extend(request , params )); _ = response;
  order := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = order;
  return this.ParseOrder(order );
}

func (this *Bigone) FetchOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchOrders() requires a symbol argument") )));
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"asset_pair_name":*(market ).At(MkString("id") )}); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("privateGetOrders"), this.Extend(request , params )); _ = response;
  orders := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = orders;
  return this.ParseOrders(orders , market , since , limit );
}

func (this *Bigone) FetchMyTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchMyTrades() requires a symbol argument") )));
  }
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"asset_pair_name":*(market ).At(MkString("id") )}); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("privateGetTrades"), this.Extend(request , params )); _ = response;
  trades := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = trades;
  return this.ParseTrades(trades , market , since , limit );
}

func (this *Bigone) ParseOrderStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "PENDING":MkString("open") ,
        "FILLED":MkString("closed") ,
        "CANCELLED":MkString("canceled") ,
        }); _ = statuses;
  return this.SafeString(statuses , status );
}

func (this *Bigone) FetchOpenOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"state":MkString("PENDING") }); _ = request;
  return this.FetchOrders(symbol , since , limit , this.Extend(request , params ));
}

func (this *Bigone) FetchClosedOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"state":MkString("FILLED") }); _ = request;
  return this.FetchOrders(symbol , since , limit , this.Extend(request , params ));
}

func (this *Bigone) Nonce(goArgs ...*Variant) *Variant{
  return OpMulti(this.Microseconds(), MkInteger(1000) );
}

func (this *Bigone) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("public") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  query := this.Omit(params , this.ExtractParams(path )); _ = query;
  baseUrl := this.ImplodeHostname(*(*(*this.At(MkString("urls")) ).At(MkString("api") )).At(api )); _ = baseUrl;
  url := OpAdd(baseUrl , OpAdd(MkString("/") , this.ImplodeParams(path , params ))); _ = url;
  if IsTrue(OpEq2(api , MkString("public") )) {
    if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
      OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(query )));
    }
  } else {
    this.CheckRequiredCredentials();
    nonce := (this.Nonce()).Call(MkString("toString") ); _ = nonce;
    request := MkMap(&VarMap{
          "type":MkString("OpenAPIV2") ,
          "sub":*this.At(MkString("apiKey")) ,
          "nonce":nonce ,
          }); _ = request;
    jwt := this.Jwt(request , this.Encode(*this.At(MkString("secret")) )); _ = jwt;
    headers = MkMap(&VarMap{"Authorization":OpAdd(MkString("Bearer ") , jwt )});
    if IsTrue(OpEq2(method , MkString("GET") )) {
      if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
        OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(query )));
      }
    } else {
      if IsTrue(OpEq2(method , MkString("POST") )) {
        *(headers ).At(MkString("Content-Type") )= MkString("application/json") ;
        body = this.Json(query );
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

func (this *Bigone) FetchDepositAddress(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  currency := this.Currency(code ); _ = currency;
  request := MkMap(&VarMap{"asset_symbol":*(currency ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("privateGetAssetsAssetSymbolAddress"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = data;
  dataLength := OpCopy(data.Length ); _ = dataLength;
  if IsTrue(OpLw(dataLength , MkInteger(1) )) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , MkString("fetchDepositAddress() returned empty address response") )));
  }
  firstElement := *(data ).At(MkInteger(0) ); _ = firstElement;
  address := this.SafeString(firstElement , MkString("value") ); _ = address;
  tag := this.SafeString(firstElement , MkString("memo") ); _ = tag;
  this.CheckAddress(address );
  return MkMap(&VarMap{
        "currency":code ,
        "address":address ,
        "tag":tag ,
        "info":response ,
        });
}

func (this *Bigone) ParseTransactionStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "WITHHOLD":MkString("ok") ,
        "UNCONFIRMED":MkString("pending") ,
        "CONFIRMED":MkString("ok") ,
        "COMPLETED":MkString("ok") ,
        "PENDING":MkString("pending") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Bigone) ParseTransaction(goArgs ...*Variant) *Variant{
  transaction := GoGetArg(goArgs, 0, MkUndefined()); _ = transaction;
  currency := GoGetArg(goArgs, 1, MkUndefined() ); _ = currency;
  currencyId := this.SafeString(transaction , MkString("asset_symbol") ); _ = currencyId;
  code := this.SafeCurrencyCode(currencyId ); _ = code;
  id := this.SafeInteger(transaction , MkString("id") ); _ = id;
  amount := this.SafeNumber(transaction , MkString("amount") ); _ = amount;
  status := this.ParseTransactionStatus(this.SafeString(transaction , MkString("state") )); _ = status;
  timestamp := this.Parse8601(this.SafeString(transaction , MkString("inserted_at") )); _ = timestamp;
  updated := this.Parse8601(this.SafeString2(transaction , MkString("updated_at") , MkString("completed_at") )); _ = updated;
  txid := this.SafeString(transaction , MkString("txid") ); _ = txid;
  address := this.SafeString(transaction , MkString("target_address") ); _ = address;
  tag := this.SafeString(transaction , MkString("memo") ); _ = tag;
  type_ := OpTrinary(OpHasMember(MkString("customer_id") , transaction ), MkString("deposit") , MkString("withdrawal") ); _ = type_;
  return MkMap(&VarMap{
        "info":transaction ,
        "id":id ,
        "txid":txid ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "addressFrom":MkUndefined() ,
        "address":MkUndefined() ,
        "addressTo":address ,
        "tagFrom":MkUndefined() ,
        "tag":tag ,
        "tagTo":MkUndefined() ,
        "type":type_ ,
        "amount":amount ,
        "currency":code ,
        "status":status ,
        "updated":updated ,
        "fee":MkUndefined() ,
        });
}

func (this *Bigone) FetchDeposits(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  currency := OpCopy(MkUndefined() ); _ = currency;
  if IsTrue(OpNotEq2(code , MkUndefined() )) {
    currency = this.Currency(code );
    *(request ).At(MkString("asset_symbol") )= *(currency ).At(MkString("id") );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("privateGetDeposits"), this.Extend(request , params )); _ = response;
  deposits := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = deposits;
  return this.ParseTransactions(deposits , code , since , limit );
}

func (this *Bigone) FetchWithdrawals(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  currency := OpCopy(MkUndefined() ); _ = currency;
  if IsTrue(OpNotEq2(code , MkUndefined() )) {
    currency = this.Currency(code );
    *(request ).At(MkString("asset_symbol") )= *(currency ).At(MkString("id") );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("privateGetWithdrawals"), this.Extend(request , params )); _ = response;
  withdrawals := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = withdrawals;
  return this.ParseTransactions(withdrawals , code , since , limit );
}

func (this *Bigone) Withdraw(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  amount := GoGetArg(goArgs, 1, MkUndefined()); _ = amount;
  address := GoGetArg(goArgs, 2, MkUndefined()); _ = address;
  tag := GoGetArg(goArgs, 3, MkUndefined() ); _ = tag;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  currency := this.Currency(code ); _ = currency;
  request := MkMap(&VarMap{
        "symbol":*(currency ).At(MkString("id") ),
        "target_address":address ,
        "amount":this.CurrencyToPrecision(code , amount ),
        }); _ = request;
  if IsTrue(OpNotEq2(tag , MkUndefined() )) {
    *(request ).At(MkString("memo") )= OpCopy(tag );
  }
  response := this.Call(MkString("privatePostWithdrawals"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  return this.ParseTransaction(data , currency );
}

func (this *Bigone) HandleErrors(goArgs ...*Variant) *Variant{
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
  message := this.SafeString(response , MkString("message") ); _ = message;
  if IsTrue(OpNotEq2(code , MkString("0") )) {
    feedback := OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , body )); _ = feedback;
    this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("exact") ), message , feedback );
    this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("exact") ), code , feedback );
    this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("broad") ), message , feedback );
    panic(NewExchangeError(feedback ));
  }
  return MkUndefined()
}

