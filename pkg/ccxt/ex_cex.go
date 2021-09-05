package ccxt

import (
)

type Cex struct {
	*ExchangeBase
}

var _ Exchange = (*Cex)(nil)

func init() {
	exchange := &Cex{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Cex) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("cex") ,
            "name":MkString("CEX.IO") ,
            "countries":MkArray(&VarArray{
                MkString("GB") ,
                MkString("EU") ,
                MkString("CY") ,
                MkString("RU") ,
                }),
            "rateLimit":MkInteger(1500) ,
            "has":MkMap(&VarMap{
                "cancelOrder":MkBool(true) ,
                "CORS":MkBool(false) ,
                "createOrder":MkBool(true) ,
                "editOrder":MkBool(true) ,
                "fetchBalance":MkBool(true) ,
                "fetchClosedOrders":MkBool(true) ,
                "fetchCurrencies":MkBool(true) ,
                "fetchDepositAddress":MkBool(true) ,
                "fetchMarkets":MkBool(true) ,
                "fetchOHLCV":MkBool(true) ,
                "fetchOpenOrders":MkBool(true) ,
                "fetchOrder":MkBool(true) ,
                "fetchOrderBook":MkBool(true) ,
                "fetchOrders":MkBool(true) ,
                "fetchTicker":MkBool(true) ,
                "fetchTickers":MkBool(true) ,
                "fetchTrades":MkBool(true) ,
                }),
            "timeframes":MkMap(&VarMap{"1m":MkString("1m") }),
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/1294454/27766442-8ddc33b0-5ed8-11e7-8b98-f786aef0f3c9.jpg") ,
                "api":MkString("https://cex.io/api") ,
                "www":MkString("https://cex.io") ,
                "doc":MkString("https://cex.io/cex-api") ,
                "fees":MkArray(&VarArray{
                    MkString("https://cex.io/fee-schedule") ,
                    MkString("https://cex.io/limits-commissions") ,
                    }),
                "referral":MkString("https://cex.io/r/0/up105393824/0/") ,
                }),
            "requiredCredentials":MkMap(&VarMap{
                "apiKey":MkBool(true) ,
                "secret":MkBool(true) ,
                "uid":MkBool(true) ,
                }),
            "api":MkMap(&VarMap{
                "public":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("currency_profile") ,
                        MkString("currency_limits/") ,
                        MkString("last_price/{pair}/") ,
                        MkString("last_prices/{currencies}/") ,
                        MkString("ohlcv/hd/{yyyymmdd}/{pair}") ,
                        MkString("order_book/{pair}/") ,
                        MkString("ticker/{pair}/") ,
                        MkString("tickers/{currencies}/") ,
                        MkString("trade_history/{pair}/") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("convert/{pair}") ,
                        MkString("price_stats/{pair}") ,
                        }),
                    }),
                "private":MkMap(&VarMap{"post":MkArray(&VarArray{
                        MkString("active_orders_status/") ,
                        MkString("archived_orders/{pair}/") ,
                        MkString("balance/") ,
                        MkString("cancel_order/") ,
                        MkString("cancel_orders/{pair}/") ,
                        MkString("cancel_replace_order/{pair}/") ,
                        MkString("close_position/{pair}/") ,
                        MkString("get_address/") ,
                        MkString("get_myfee/") ,
                        MkString("get_order/") ,
                        MkString("get_order_tx/") ,
                        MkString("open_orders/{pair}/") ,
                        MkString("open_orders/") ,
                        MkString("open_position/{pair}/") ,
                        MkString("open_positions/{pair}/") ,
                        MkString("place_order/{pair}/") ,
                        MkString("raw_tx_history") ,
                        })}),
                }),
            "fees":MkMap(&VarMap{
                "trading":MkMap(&VarMap{
                    "maker":this.ParseNumber(MkString("0.0016") ),
                    "taker":this.ParseNumber(MkString("0.0025") ),
                    }),
                "funding":MkMap(&VarMap{
                    "withdraw":MkMap(&VarMap{}),
                    "deposit":MkMap(&VarMap{
                        "BTC":MkNumber(0.0) ,
                        "ETH":MkNumber(0.0) ,
                        "BCH":MkNumber(0.0) ,
                        "DASH":MkNumber(0.0) ,
                        "BTG":MkNumber(0.0) ,
                        "ZEC":MkNumber(0.0) ,
                        "XRP":MkNumber(0.0) ,
                        "XLM":MkNumber(0.0) ,
                        }),
                    }),
                }),
            "exceptions":MkMap(&VarMap{
                "exact":MkMap(&VarMap{}),
                "broad":MkMap(&VarMap{
                    "Insufficient funds":InsufficientFunds ,
                    "Nonce must be incremented":InvalidNonce ,
                    "Invalid Order":InvalidOrder ,
                    "Order not found":OrderNotFound ,
                    "limit exceeded":RateLimitExceeded ,
                    "Invalid API key":AuthenticationError ,
                    "There was an error while placing your order":InvalidOrder ,
                    "Sorry, too many clients already":DDoSProtection ,
                    }),
                }),
            "options":MkMap(&VarMap{
                "fetchOHLCVWarning":MkBool(true) ,
                "createMarketBuyOrderRequiresPrice":MkBool(true) ,
                "order":MkMap(&VarMap{"status":MkMap(&VarMap{
                        "c":MkString("canceled") ,
                        "d":MkString("closed") ,
                        "cd":MkString("canceled") ,
                        "a":MkString("open") ,
                        })}),
                }),
            }));
}

func (this *Cex) FetchCurrenciesFromCache(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  options := this.SafeValue(*this.At(MkString("options")) , MkString("fetchCurrencies") , MkMap(&VarMap{})); _ = options;
  timestamp := this.SafeInteger(options , MkString("timestamp") ); _ = timestamp;
  expires := this.SafeInteger(options , MkString("expires") , MkInteger(1000) ); _ = expires;
  now := this.Milliseconds(); _ = now;
  if IsTrue(OpOr(OpEq2(timestamp , MkUndefined() ), OpGt(OpSub(now , timestamp ), expires ))) {
    response := this.Call(MkString("publicGetCurrencyProfile"), params ); _ = response;
    *(*this.At(MkString("options")) ).At(MkString("fetchCurrencies") )= this.Extend(options , MkMap(&VarMap{
            "response":response ,
            "timestamp":now ,
            }));
  }
  return this.SafeValue(*(*this.At(MkString("options")) ).At(MkString("fetchCurrencies") ), MkString("response") );
}

func (this *Cex) FetchCurrencies(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.FetchCurrenciesFromCache(params ); _ = response;
  *(*this.At(MkString("options")) ).At(MkString("currencies") )= MkMap(&VarMap{
      "timestamp":this.Milliseconds(),
      "response":response ,
      });
  data := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = data;
  currencies := this.SafeValue(data , MkString("symbols") , MkArray(&VarArray{})); _ = currencies;
  result := MkMap(&VarMap{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , currencies.Length )); OpIncr(& i ){
    currency := *(currencies ).At(i ); _ = currency;
    id := this.SafeString(currency , MkString("code") ); _ = id;
    code := this.SafeCurrencyCode(id ); _ = code;
    precision := this.SafeInteger(currency , MkString("precision") ); _ = precision;
    active := OpCopy(MkBool(true) ); _ = active;
    *(result ).At(code )= MkMap(&VarMap{
        "id":id ,
        "code":code ,
        "name":id ,
        "active":active ,
        "precision":precision ,
        "fee":MkUndefined() ,
        "limits":MkMap(&VarMap{
            "amount":MkMap(&VarMap{
                "min":this.SafeNumber(currency , MkString("minimumCurrencyAmount") ),
                "max":MkUndefined() ,
                }),
            "withdraw":MkMap(&VarMap{
                "min":this.SafeNumber(currency , MkString("minimalWithdrawalAmount") ),
                "max":MkUndefined() ,
                }),
            }),
        "info":currency ,
        });
  }
  return result ;
}

func (this *Cex) FetchMarkets(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  currenciesResponse := this.FetchCurrenciesFromCache(params ); _ = currenciesResponse;
  currenciesData := this.SafeValue(currenciesResponse , MkString("data") , MkMap(&VarMap{})); _ = currenciesData;
  currencies := this.SafeValue(currenciesData , MkString("symbols") , MkArray(&VarArray{})); _ = currencies;
  currenciesById := this.IndexBy(currencies , MkString("code") ); _ = currenciesById;
  pairs := this.SafeValue(currenciesData , MkString("pairs") , MkArray(&VarArray{})); _ = pairs;
  response := this.Call(MkString("publicGetCurrencyLimits"), params ); _ = response;
  result := MkArray(&VarArray{}); _ = result;
  markets := this.SafeValue(*(response ).At(MkString("data") ), MkString("pairs") ); _ = markets;
  for i := MkInteger(0) ; IsTrue(OpLw(i , markets.Length )); OpIncr(& i ){
    market := *(markets ).At(i ); _ = market;
    baseId := this.SafeString(market , MkString("symbol1") ); _ = baseId;
    quoteId := this.SafeString(market , MkString("symbol2") ); _ = quoteId;
    id := OpAdd(baseId , OpAdd(MkString("/") , quoteId )); _ = id;
    base := this.SafeCurrencyCode(baseId ); _ = base;
    quote := this.SafeCurrencyCode(quoteId ); _ = quote;
    symbol := OpAdd(base , OpAdd(MkString("/") , quote )); _ = symbol;
    baseCurrency := this.SafeValue(currenciesById , baseId , MkMap(&VarMap{})); _ = baseCurrency;
    quoteCurrency := this.SafeValue(currenciesById , quoteId , MkMap(&VarMap{})); _ = quoteCurrency;
    pricePrecision := this.SafeInteger(quoteCurrency , MkString("precision") , MkInteger(8) ); _ = pricePrecision;
    for j := MkInteger(0) ; IsTrue(OpLw(j , pairs.Length )); OpIncr(& j ){
      pair := *(pairs ).At(j ); _ = pair;
      if IsTrue(OpAnd(OpEq2(*(pair ).At(MkString("symbol1") ), baseId ), OpEq2(*(pair ).At(MkString("symbol2") ), quoteId ))) {
        pricePrecision = this.SafeInteger(pair , MkString("pricePrecision") , pricePrecision );
      }
    }
    baseCcyPrecision := this.SafeInteger(baseCurrency , MkString("precision") , MkInteger(8) ); _ = baseCcyPrecision;
    baseCcyScale := this.SafeInteger(baseCurrency , MkString("scale") , MkInteger(0) ); _ = baseCcyScale;
    amountPrecision := OpSub(baseCcyPrecision , baseCcyScale ); _ = amountPrecision;
    precision := MkMap(&VarMap{
          "amount":amountPrecision ,
          "price":pricePrecision ,
          }); _ = precision;
    result.Push(MkMap(&VarMap{
            "id":id ,
            "info":market ,
            "symbol":symbol ,
            "base":base ,
            "quote":quote ,
            "baseId":baseId ,
            "quoteId":quoteId ,
            "precision":precision ,
            "limits":MkMap(&VarMap{
                "amount":MkMap(&VarMap{
                    "min":this.SafeNumber(market , MkString("minLotSize") ),
                    "max":this.SafeNumber(market , MkString("maxLotSize") ),
                    }),
                "price":MkMap(&VarMap{
                    "min":this.SafeNumber(market , MkString("minPrice") ),
                    "max":this.SafeNumber(market , MkString("maxPrice") ),
                    }),
                "cost":MkMap(&VarMap{
                    "min":this.SafeNumber(market , MkString("minLotSizeS2") ),
                    "max":MkUndefined() ,
                    }),
                }),
            "active":MkUndefined() ,
            }));
  }
  return result ;
}

func (this *Cex) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("privatePostBalance"), params ); _ = response;
  result := MkMap(&VarMap{"info":response }); _ = result;
  ommited := MkArray(&VarArray{
        MkString("username") ,
        MkString("timestamp") ,
        }); _ = ommited;
  balances := this.Omit(response , ommited ); _ = balances;
  currencyIds := GoGetKeys(balances ); _ = currencyIds;
  for i := MkInteger(0) ; IsTrue(OpLw(i , currencyIds.Length )); OpIncr(& i ){
    currencyId := *(currencyIds ).At(i ); _ = currencyId;
    balance := this.SafeValue(balances , currencyId , MkMap(&VarMap{})); _ = balance;
    account := this.Account(); _ = account;
    *(account ).At(MkString("free") )= this.SafeString(balance , MkString("available") );
    *(account ).At(MkString("used") )= this.SafeString(balance , MkString("orders") , MkString("0") );
    code := this.SafeCurrencyCode(currencyId ); _ = code;
    *(result ).At(code )= OpCopy(account );
  }
  return this.ParseBalance(result );
}

func (this *Cex) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"pair":this.MarketId(symbol )}); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("depth") )= OpCopy(limit );
  }
  response := this.Call(MkString("publicGetOrderBookPair"), this.Extend(request , params )); _ = response;
  timestamp := this.SafeTimestamp(response , MkString("timestamp") ); _ = timestamp;
  return this.ParseOrderBook(response , symbol , timestamp );
}

func (this *Cex) ParseOHLCV(goArgs ...*Variant) *Variant{
  ohlcv := GoGetArg(goArgs, 0, MkUndefined()); _ = ohlcv;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  return MkArray(&VarArray{
        this.SafeTimestamp(ohlcv , MkInteger(0) ),
        this.SafeNumber(ohlcv , MkInteger(1) ),
        this.SafeNumber(ohlcv , MkInteger(2) ),
        this.SafeNumber(ohlcv , MkInteger(3) ),
        this.SafeNumber(ohlcv , MkInteger(4) ),
        this.SafeNumber(ohlcv , MkInteger(5) ),
        });
}

func (this *Cex) FetchOHLCV(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  timeframe := GoGetArg(goArgs, 1, MkString("1m") ); _ = timeframe;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  if IsTrue(OpEq2(since , MkUndefined() )) {
    since = OpSub(this.Milliseconds(), MkInteger(86400000) );
  } else {
    if IsTrue(*(*this.At(MkString("options")) ).At(MkString("fetchOHLCVWarning") )) {
      panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , MkString(" fetchOHLCV warning: CEX can return historical candles for a certain date only, this might produce an empty or null reply. Set exchange.options['fetchOHLCVWarning'] = false or add ({ 'options': { 'fetchOHLCVWarning': false }}) to constructor params to suppress this warning message.") )));
    }
  }
  ymd := this.Ymd(since ); _ = ymd;
  ymd = ymd.Split(MkString("-") );
  ymd = ymd.Join(MkString("") );
  request := MkMap(&VarMap{
        "pair":*(market ).At(MkString("id") ),
        "yyyymmdd":ymd ,
        }); _ = request;
  {
  ret__ := func(this *Cex) (ret_ *Variant) {
    defer func() {
      if e := recover().(*Variant); e != nil {
        ret_ = func(this *Cex) *Variant {
          // catch block:
          if IsTrue(OpIsType(e , NullResponse )) {
            return MkArray(&VarArray{});
          }
          return nil
        }(this)
      }
    }()
    // try block:
    response := this.Call(MkString("publicGetOhlcvHdYyyymmddPair"), this.Extend(request , params )); _ = response;
    key := OpAdd(MkString("data") , *(*this.At(MkString("timeframes")) ).At(timeframe )); _ = key;
    data := this.SafeString(response , key ); _ = data;
    ohlcvs := JSON.Parse(data ); _ = ohlcvs;
    return this.ParseOHLCVs(ohlcvs , market , timeframe , since , limit );
    return nil
  }(this)
  if ret__ != nil {
    return ret__;
   }
  }
  return MkUndefined()
}

func (this *Cex) ParseTicker(goArgs ...*Variant) *Variant{
  ticker := GoGetArg(goArgs, 0, MkUndefined()); _ = ticker;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.SafeTimestamp(ticker , MkString("timestamp") ); _ = timestamp;
  volume := this.SafeNumber(ticker , MkString("volume") ); _ = volume;
  high := this.SafeNumber(ticker , MkString("high") ); _ = high;
  low := this.SafeNumber(ticker , MkString("low") ); _ = low;
  bid := this.SafeNumber(ticker , MkString("bid") ); _ = bid;
  ask := this.SafeNumber(ticker , MkString("ask") ); _ = ask;
  last := this.SafeNumber(ticker , MkString("last") ); _ = last;
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  if IsTrue(market ) {
    symbol = *(market ).At(MkString("symbol") );
  }
  return MkMap(&VarMap{
        "symbol":symbol ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "high":high ,
        "low":low ,
        "bid":bid ,
        "bidVolume":MkUndefined() ,
        "ask":ask ,
        "askVolume":MkUndefined() ,
        "vwap":MkUndefined() ,
        "open":MkUndefined() ,
        "close":last ,
        "last":last ,
        "previousClose":MkUndefined() ,
        "change":MkUndefined() ,
        "percentage":MkUndefined() ,
        "average":MkUndefined() ,
        "baseVolume":volume ,
        "quoteVolume":MkUndefined() ,
        "info":ticker ,
        });
}

func (this *Cex) FetchTickers(goArgs ...*Variant) *Variant{
  symbols := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbols;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  currencies := GoGetKeys(*this.At(MkString("currencies")) ); _ = currencies;
  request := MkMap(&VarMap{"currencies":currencies.Join(MkString("/") )}); _ = request;
  response := this.Call(MkString("publicGetTickersCurrencies"), this.Extend(request , params )); _ = response;
  tickers := *(response ).At(MkString("data") ); _ = tickers;
  result := MkMap(&VarMap{}); _ = result;
  for t := MkInteger(0) ; IsTrue(OpLw(t , tickers.Length )); OpIncr(& t ){
    ticker := *(tickers ).At(t ); _ = ticker;
    symbol := (*(ticker ).At(MkString("pair") )).Call(MkString("replace") , MkString(":") , MkString("/") ); _ = symbol;
    market := *(*this.At(MkString("markets")) ).At(symbol ); _ = market;
    *(result ).At(symbol )= this.ParseTicker(ticker , market );
  }
  return this.FilterByArray(result , MkString("symbol") , symbols );
}

func (this *Cex) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"pair":*(market ).At(MkString("id") )}); _ = request;
  ticker := this.Call(MkString("publicGetTickerPair"), this.Extend(request , params )); _ = ticker;
  return this.ParseTicker(ticker , market );
}

func (this *Cex) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.SafeTimestamp(trade , MkString("date") ); _ = timestamp;
  id := this.SafeString(trade , MkString("tid") ); _ = id;
  type_ := OpCopy(MkUndefined() ); _ = type_;
  side := this.SafeString(trade , MkString("type") ); _ = side;
  priceString := this.SafeString(trade , MkString("price") ); _ = priceString;
  amountString := this.SafeString(trade , MkString("amount") ); _ = amountString;
  price := this.ParseNumber(priceString ); _ = price;
  amount := this.ParseNumber(amountString ); _ = amount;
  cost := this.ParseNumber(Precise.StringMul(priceString , amountString )); _ = cost;
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
        "type":type_ ,
        "side":side ,
        "order":MkUndefined() ,
        "takerOrMaker":MkUndefined() ,
        "price":price ,
        "amount":amount ,
        "cost":cost ,
        "fee":MkUndefined() ,
        });
}

func (this *Cex) FetchTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"pair":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("publicGetTradeHistoryPair"), this.Extend(request , params )); _ = response;
  return this.ParseTrades(response , market , since , limit );
}

func (this *Cex) CreateOrder(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 2, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 3, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 4, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 5, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpAnd(OpEq2(type_ , MkString("market") ), OpEq2(side , MkString("buy") ))) {
    if IsTrue(*(*this.At(MkString("options")) ).At(MkString("createMarketBuyOrderRequiresPrice") )) {
      if IsTrue(OpEq2(price , MkUndefined() )) {
        panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")) , MkString(" createOrder() requires the price argument with market buy orders to calculate total order cost (amount to spend), where cost = amount * price. Supply a price argument to createOrder() call if you want the cost to be calculated for you from price and amount, or, alternatively, add .options['createMarketBuyOrderRequiresPrice'] = false to supply the cost in the amount argument (the exchange-specific behaviour)") )));
      } else {
        amount = OpMulti(amount , price );
      }
    }
  }
  this.LoadMarkets();
  request := MkMap(&VarMap{
        "pair":this.MarketId(symbol ),
        "type":side ,
        "amount":amount ,
        }); _ = request;
  if IsTrue(OpEq2(type_ , MkString("limit") )) {
    *(request ).At(MkString("price") )= OpCopy(price );
  } else {
    *(request ).At(MkString("order_type") )= OpCopy(type_ );
  }
  response := this.Call(MkString("privatePostPlaceOrderPair"), this.Extend(request , params )); _ = response;
  placedAmount := this.SafeNumber(response , MkString("amount") ); _ = placedAmount;
  remaining := this.SafeNumber(response , MkString("pending") ); _ = remaining;
  timestamp := this.SafeValue(response , MkString("time") ); _ = timestamp;
  complete := this.SafeValue(response , MkString("complete") ); _ = complete;
  status := OpTrinary(complete , MkString("closed") , MkString("open") ); _ = status;
  filled := OpCopy(MkUndefined() ); _ = filled;
  if IsTrue(OpAnd(OpNotEq2(placedAmount , MkUndefined() ), OpNotEq2(remaining , MkUndefined() ))) {
    filled = Math.Max(OpSub(placedAmount , remaining ), MkInteger(0) );
  }
  return MkMap(&VarMap{
        "id":this.SafeString(response , MkString("id") ),
        "info":response ,
        "clientOrderId":MkUndefined() ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "lastTradeTimestamp":MkUndefined() ,
        "type":type_ ,
        "side":this.SafeString(response , MkString("type") ),
        "symbol":symbol ,
        "status":status ,
        "price":this.SafeNumber(response , MkString("price") ),
        "amount":placedAmount ,
        "cost":MkUndefined() ,
        "average":MkUndefined() ,
        "remaining":remaining ,
        "filled":filled ,
        "fee":MkUndefined() ,
        "trades":MkUndefined() ,
        });
}

func (this *Cex) CancelOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"id":id }); _ = request;
  return this.Call(MkString("privatePostCancelOrder"), this.Extend(request , params ));
}

func (this *Cex) ParseOrder(goArgs ...*Variant) *Variant{
  order := GoGetArg(goArgs, 0, MkUndefined()); _ = order;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.SafeValue(order , MkString("time") ); _ = timestamp;
  if IsTrue(OpAnd(OpEq2(OpType(timestamp ), MkString("string") ), OpGtEq(timestamp.IndexOf(MkString("T") ), MkInteger(0) ))) {
    timestamp = this.Parse8601(timestamp );
  } else {
    timestamp = parseInt(timestamp );
  }
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  if IsTrue(OpEq2(market , MkUndefined() )) {
    baseId := this.SafeString(order , MkString("symbol1") ); _ = baseId;
    quoteId := this.SafeString(order , MkString("symbol2") ); _ = quoteId;
    base := this.SafeCurrencyCode(baseId ); _ = base;
    quote := this.SafeCurrencyCode(quoteId ); _ = quote;
    symbol = OpAdd(base , OpAdd(MkString("/") , quote ));
    if IsTrue(OpHasMember(symbol , *this.At(MkString("markets")) )) {
      market = this.Market(symbol );
    }
  }
  status := this.ParseOrderStatus(this.SafeString(order , MkString("status") )); _ = status;
  price := this.SafeNumber(order , MkString("price") ); _ = price;
  amount := this.SafeNumber(order , MkString("amount") ); _ = amount;
  if IsTrue(OpNotEq2(amount , MkUndefined() )) {
    amount = Math.Abs(amount );
  }
  remaining := this.SafeNumber2(order , MkString("pending") , MkString("remains") ); _ = remaining;
  filled := OpSub(amount , remaining ); _ = filled;
  fee := OpCopy(MkUndefined() ); _ = fee;
  cost := OpCopy(MkUndefined() ); _ = cost;
  if IsTrue(OpNotEq2(market , MkUndefined() )) {
    symbol = *(market ).At(MkString("symbol") );
    taCost := this.SafeNumber(order , OpAdd(MkString("ta:") , *(market ).At(MkString("quote") ))); _ = taCost;
    ttaCost := this.SafeNumber(order , OpAdd(MkString("tta:") , *(market ).At(MkString("quote") ))); _ = ttaCost;
    cost = this.Sum(taCost , ttaCost );
    baseFee := OpAdd(MkString("fa:") , *(market ).At(MkString("base") )); _ = baseFee;
    baseTakerFee := OpAdd(MkString("tfa:") , *(market ).At(MkString("base") )); _ = baseTakerFee;
    quoteFee := OpAdd(MkString("fa:") , *(market ).At(MkString("quote") )); _ = quoteFee;
    quoteTakerFee := OpAdd(MkString("tfa:") , *(market ).At(MkString("quote") )); _ = quoteTakerFee;
    feeRate := this.SafeNumber(order , MkString("tradingFeeMaker") ); _ = feeRate;
    if IsTrue(OpNot(feeRate )) {
      feeRate = this.SafeNumber(order , MkString("tradingFeeTaker") , feeRate );
    }
    if IsTrue(feeRate ) {
      OpDivAssign(& feeRate , MkNumber(100.0) );
    }
    if IsTrue(OpOr(OpHasMember(baseFee , order ), OpHasMember(baseTakerFee , order ))) {
      baseFeeCost := this.SafeNumber2(order , baseFee , baseTakerFee ); _ = baseFeeCost;
      fee = MkMap(&VarMap{
          "currency":*(market ).At(MkString("base") ),
          "rate":feeRate ,
          "cost":baseFeeCost ,
          });
    } else {
      if IsTrue(OpOr(OpHasMember(quoteFee , order ), OpHasMember(quoteTakerFee , order ))) {
        quoteFeeCost := this.SafeNumber2(order , quoteFee , quoteTakerFee ); _ = quoteFeeCost;
        fee = MkMap(&VarMap{
            "currency":*(market ).At(MkString("quote") ),
            "rate":feeRate ,
            "cost":quoteFeeCost ,
            });
      }
    }
  }
  if IsTrue(OpNot(cost )) {
    cost = OpMulti(price , filled );
  }
  side := *(order ).At(MkString("type") ); _ = side;
  trades := OpCopy(MkUndefined() ); _ = trades;
  orderId := *(order ).At(MkString("id") ); _ = orderId;
  if IsTrue(OpHasMember(MkString("vtx") , order )) {
    trades = MkArray(&VarArray{});
    for i := MkInteger(0) ; IsTrue(OpLw(i , *(*(order ).At(MkString("vtx") )).At(MkString("length") ))); OpIncr(& i ){
      item := *(*(order ).At(MkString("vtx") )).At(i ); _ = item;
      tradeSide := this.SafeString(item , MkString("type") ); _ = tradeSide;
      if IsTrue(OpEq2(tradeSide , MkString("cancel") )) {
        continue ;
      }
      tradePrice := this.SafeNumber(item , MkString("price") ); _ = tradePrice;
      if IsTrue(OpEq2(tradePrice , MkUndefined() )) {
        continue ;
      }
      if IsTrue(OpEq2(tradeSide , MkString("costsNothing") )) {
        continue ;
      }
      tradeTimestamp := this.Parse8601(this.SafeString(item , MkString("time") )); _ = tradeTimestamp;
      tradeAmount := this.SafeNumber(item , MkString("amount") ); _ = tradeAmount;
      feeCost := this.SafeNumber(item , MkString("fee_amount") ); _ = feeCost;
      absTradeAmount := OpTrinary(OpLw(tradeAmount , MkInteger(0) ), OpNeg(tradeAmount ), tradeAmount ); _ = absTradeAmount;
      tradeCost := OpCopy(MkUndefined() ); _ = tradeCost;
      if IsTrue(OpEq2(tradeSide , MkString("sell") )) {
        tradeCost = OpCopy(absTradeAmount );
        absTradeAmount = OpDiv(this.Sum(feeCost , tradeCost ), tradePrice );
      } else {
        tradeCost = OpMulti(absTradeAmount , tradePrice );
      }
      trades.Push(MkMap(&VarMap{
              "id":this.SafeString(item , MkString("id") ),
              "timestamp":tradeTimestamp ,
              "datetime":this.Iso8601(tradeTimestamp ),
              "order":orderId ,
              "symbol":symbol ,
              "price":tradePrice ,
              "amount":absTradeAmount ,
              "cost":tradeCost ,
              "side":tradeSide ,
              "fee":MkMap(&VarMap{
                  "cost":feeCost ,
                  "currency":*(market ).At(MkString("quote") ),
                  }),
              "info":item ,
              "type":MkUndefined() ,
              "takerOrMaker":MkUndefined() ,
              }));
    }
  }
  return MkMap(&VarMap{
        "id":orderId ,
        "clientOrderId":MkUndefined() ,
        "datetime":this.Iso8601(timestamp ),
        "timestamp":timestamp ,
        "lastTradeTimestamp":MkUndefined() ,
        "status":status ,
        "symbol":symbol ,
        "type":OpTrinary(OpEq2(price , MkUndefined() ), MkString("market") , MkString("limit") ),
        "timeInForce":MkUndefined() ,
        "postOnly":MkUndefined() ,
        "side":side ,
        "price":price ,
        "stopPrice":MkUndefined() ,
        "cost":cost ,
        "amount":amount ,
        "filled":filled ,
        "remaining":remaining ,
        "trades":trades ,
        "fee":fee ,
        "info":order ,
        "average":MkUndefined() ,
        });
}

func (this *Cex) FetchOpenOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  method := MkString("privatePostOpenOrders") ; _ = method;
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("pair") )= *(market ).At(MkString("id") );
    OpAddAssign(& method , MkString("Pair") );
  }
  orders := this.Call(method , this.Extend(request , params )); _ = orders;
  for i := MkInteger(0) ; IsTrue(OpLw(i , orders.Length )); OpIncr(& i ){
    *(orders ).At(i )= this.Extend(*(orders ).At(i ), MkMap(&VarMap{"status":MkString("open") }));
  }
  return this.ParseOrders(orders , market , since , limit );
}

func (this *Cex) FetchClosedOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  method := MkString("privatePostArchivedOrdersPair") ; _ = method;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchClosedOrders() requires a symbol argument") )));
  }
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"pair":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(method , this.Extend(request , params )); _ = response;
  return this.ParseOrders(response , market , since , limit );
}

func (this *Cex) FetchOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"id":id.ToString()}); _ = request;
  response := this.Call(MkString("privatePostGetOrderTx"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  return this.ParseOrder(data );
}

func (this *Cex) FetchOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "limit":limit ,
        "pair":*(market ).At(MkString("id") ),
        "dateFrom":since ,
        }); _ = request;
  response := this.Call(MkString("privatePostArchivedOrdersPair"), this.Extend(request , params )); _ = response;
  results := MkArray(&VarArray{}); _ = results;
  for i := MkInteger(0) ; IsTrue(OpLw(i , response.Length )); OpIncr(& i ){
    order := *(response ).At(i ); _ = order;
    status := this.ParseOrderStatus(this.SafeString(order , MkString("status") )); _ = status;
    baseId := this.SafeString(order , MkString("symbol1") ); _ = baseId;
    quoteId := this.SafeString(order , MkString("symbol2") ); _ = quoteId;
    base := this.SafeCurrencyCode(baseId ); _ = base;
    quote := this.SafeCurrencyCode(quoteId ); _ = quote;
    symbol := OpAdd(base , OpAdd(MkString("/") , quote )); _ = symbol;
    side := this.SafeString(order , MkString("type") ); _ = side;
    baseAmount := this.SafeNumber(order , OpAdd(MkString("a:") , OpAdd(baseId , MkString(":cds") ))); _ = baseAmount;
    quoteAmount := this.SafeNumber(order , OpAdd(MkString("a:") , OpAdd(quoteId , MkString(":cds") ))); _ = quoteAmount;
    fee := this.SafeNumber(order , OpAdd(MkString("f:") , OpAdd(quoteId , MkString(":cds") ))); _ = fee;
    amount := this.SafeNumber(order , MkString("amount") ); _ = amount;
    price := this.SafeNumber(order , MkString("price") ); _ = price;
    remaining := this.SafeNumber(order , MkString("remains") ); _ = remaining;
    filled := OpSub(amount , remaining ); _ = filled;
    orderAmount := OpCopy(MkUndefined() ); _ = orderAmount;
    cost := OpCopy(MkUndefined() ); _ = cost;
    average := OpCopy(MkUndefined() ); _ = average;
    type_ := OpCopy(MkUndefined() ); _ = type_;
    if IsTrue(OpNot(price )) {
      type_ = MkString("market") ;
      orderAmount = OpCopy(baseAmount );
      cost = OpCopy(quoteAmount );
      average = OpDiv(orderAmount , cost );
    } else {
      ta := this.SafeNumber(order , OpAdd(MkString("ta:") , quoteId ), MkInteger(0) ); _ = ta;
      tta := this.SafeNumber(order , OpAdd(MkString("tta:") , quoteId ), MkInteger(0) ); _ = tta;
      fa := this.SafeNumber(order , OpAdd(MkString("fa:") , quoteId ), MkInteger(0) ); _ = fa;
      tfa := this.SafeNumber(order , OpAdd(MkString("tfa:") , quoteId ), MkInteger(0) ); _ = tfa;
      if IsTrue(OpEq2(side , MkString("sell") )) {
        cost = this.Sum(this.Sum(ta , tta ), this.Sum(fa , tfa ));
      } else {
        cost = OpSub(this.Sum(ta , tta ), this.Sum(fa , tfa ));
      }
      type_ = MkString("limit") ;
      orderAmount = OpCopy(amount );
      average = OpDiv(cost , filled );
    }
    time := this.SafeString(order , MkString("time") ); _ = time;
    lastTxTime := this.SafeString(order , MkString("lastTxTime") ); _ = lastTxTime;
    timestamp := this.Parse8601(time ); _ = timestamp;
    results.Push(MkMap(&VarMap{
            "id":this.SafeString(order , MkString("id") ),
            "timestamp":timestamp ,
            "datetime":this.Iso8601(timestamp ),
            "lastUpdated":this.Parse8601(lastTxTime ),
            "status":status ,
            "symbol":symbol ,
            "side":side ,
            "price":price ,
            "amount":orderAmount ,
            "average":average ,
            "type":type_ ,
            "filled":filled ,
            "cost":cost ,
            "remaining":remaining ,
            "fee":MkMap(&VarMap{
                "cost":fee ,
                "currency":quote ,
                }),
            "info":order ,
            }));
  }
  return results ;
}

func (this *Cex) ParseOrderStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  return this.SafeString(*(*(*this.At(MkString("options")) ).At(MkString("order") )).At(MkString("status") ), status , status );
}

func (this *Cex) EditOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 2, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 3, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 4, MkUndefined() ); _ = amount;
  price := GoGetArg(goArgs, 5, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 6, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(amount , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" editOrder() requires a amount argument") )));
  }
  if IsTrue(OpEq2(price , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" editOrder() requires a price argument") )));
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "pair":*(market ).At(MkString("id") ),
        "type":side ,
        "amount":amount ,
        "price":price ,
        "order_id":id ,
        }); _ = request;
  response := this.Call(MkString("privatePostCancelReplaceOrderPair"), this.Extend(request , params )); _ = response;
  return this.ParseOrder(response , market );
}

func (this *Cex) FetchDepositAddress(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpOr(OpEq2(code , MkString("XRP") ), OpEq2(code , MkString("XLM") ))) {
    panic(NewNotSupported(OpAdd(*this.At(MkString("id")) , MkString(" fetchDepositAddress does not support XRP and XLM addresses yet (awaiting docs from CEX.io)") )));
  }
  this.LoadMarkets();
  currency := this.Currency(code ); _ = currency;
  request := MkMap(&VarMap{"currency":*(currency ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("privatePostGetAddress"), this.Extend(request , params )); _ = response;
  address := this.SafeString(response , MkString("data") ); _ = address;
  this.CheckAddress(address );
  return MkMap(&VarMap{
        "currency":code ,
        "address":address ,
        "tag":MkUndefined() ,
        "info":response ,
        });
}

func (this *Cex) Nonce(goArgs ...*Variant) *Variant{
  return this.Milliseconds();
}

func (this *Cex) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("public") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  url := OpAdd(*(*this.At(MkString("urls")) ).At(MkString("api") ), OpAdd(MkString("/") , this.ImplodeParams(path , params ))); _ = url;
  query := this.Omit(params , this.ExtractParams(path )); _ = query;
  if IsTrue(OpEq2(api , MkString("public") )) {
    if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
      OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(query )));
    }
  } else {
    this.CheckRequiredCredentials();
    nonce := (this.Nonce()).Call(MkString("toString") ); _ = nonce;
    auth := OpAdd(nonce , OpAdd(*this.At(MkString("uid")) , *this.At(MkString("apiKey")) )); _ = auth;
    signature := this.Hmac(this.Encode(auth ), this.Encode(*this.At(MkString("secret")) )); _ = signature;
    body = this.Json(this.Extend(MkMap(&VarMap{
                "key":*this.At(MkString("apiKey")) ,
                "signature":signature.ToUpperCase(),
                "nonce":nonce ,
                }), query ));
    headers = MkMap(&VarMap{"Content-Type":MkString("application/json") });
  }
  return MkMap(&VarMap{
        "url":url ,
        "method":method ,
        "body":body ,
        "headers":headers ,
        });
}

func (this *Cex) HandleErrors(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  reason := GoGetArg(goArgs, 1, MkUndefined()); _ = reason;
  url := GoGetArg(goArgs, 2, MkUndefined()); _ = url;
  method := GoGetArg(goArgs, 3, MkUndefined()); _ = method;
  headers := GoGetArg(goArgs, 4, MkUndefined()); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined()); _ = body;
  response := GoGetArg(goArgs, 6, MkUndefined()); _ = response;
  requestHeaders := GoGetArg(goArgs, 7, MkUndefined()); _ = requestHeaders;
  requestBody := GoGetArg(goArgs, 8, MkUndefined()); _ = requestBody;
  if IsTrue(GoIsArray(response )) {
    return response ;
  }
  if IsTrue(OpEq2(body , MkString("true") )) {
    return MkUndefined();
  }
  if IsTrue(OpEq2(response , MkUndefined() )) {
    panic(NewNullResponse(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" returned ") , this.Json(response )))));
  }
  if IsTrue(OpHasMember(MkString("e") , response )) {
    if IsTrue(OpHasMember(MkString("ok") , response )) {
      if IsTrue(OpEq2(*(response ).At(MkString("ok") ), MkString("ok") )) {
        return MkUndefined();
      }
    }
  }
  if IsTrue(OpHasMember(MkString("error") , response )) {
    message := this.SafeString(response , MkString("error") ); _ = message;
    feedback := OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , body )); _ = feedback;
    this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("exact") ), message , feedback );
    this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("broad") ), message , feedback );
    panic(NewExchangeError(feedback ));
  }
  return MkUndefined()
}

