package ccxt

import (
)

type Qtrade struct {
	*ExchangeBase
}

var _ Exchange = (*Qtrade)(nil)

func init() {
	exchange := &Qtrade{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Qtrade) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("qtrade") ,
            "name":MkString("qTrade") ,
            "countries":MkArray(&VarArray{
                MkString("US") ,
                }),
            "rateLimit":MkInteger(1000) ,
            "version":MkString("v1") ,
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/51840849/80491487-74a99c00-896b-11ea-821e-d307e832f13e.jpg") ,
                "api":MkString("https://api.qtrade.io") ,
                "www":MkString("https://qtrade.io") ,
                "doc":MkString("https://qtrade-exchange.github.io/qtrade-docs") ,
                "referral":MkString("https://qtrade.io/?ref=BKOQWVFGRH2C") ,
                }),
            "has":MkMap(&VarMap{
                "CORS":MkBool(false) ,
                "fetchTrades":MkBool(true) ,
                "fetchTicker":MkBool(true) ,
                "fetchTickers":MkBool(true) ,
                "fetchMarkets":MkBool(true) ,
                "fetchCurrencies":MkBool(true) ,
                "fetchBalance":MkBool(true) ,
                "fetchOrderBook":MkBool(true) ,
                "fetchOrder":MkBool(true) ,
                "fetchOrders":MkBool(true) ,
                "fetchMyTrades":MkBool(true) ,
                "fetchClosedOrders":MkBool(true) ,
                "fetchOpenOrders":MkBool(true) ,
                "fetchOHLCV":MkBool(true) ,
                "createOrder":MkBool(true) ,
                "cancelOrder":MkBool(true) ,
                "createMarketOrder":MkBool(false) ,
                "withdraw":MkBool(true) ,
                "fetchDepositAddress":MkBool(true) ,
                "fetchTransactions":MkBool(false) ,
                "fetchDeposits":MkBool(true) ,
                "fetchWithdrawals":MkBool(true) ,
                "fetchDeposit":MkBool(true) ,
                "fetchWithdrawal":MkBool(true) ,
                }),
            "timeframes":MkMap(&VarMap{
                "5m":MkString("fivemin") ,
                "15m":MkString("fifteenmin") ,
                "30m":MkString("thirtymin") ,
                "1h":MkString("onehour") ,
                "2h":MkString("twohour") ,
                "4h":MkString("fourhour") ,
                "1d":MkString("oneday") ,
                }),
            "api":MkMap(&VarMap{
                "public":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("ticker/{market_string}") ,
                        MkString("tickers") ,
                        MkString("currency/{code}") ,
                        MkString("currencies") ,
                        MkString("common") ,
                        MkString("market/{market_string}") ,
                        MkString("markets") ,
                        MkString("market/{market_string}/trades") ,
                        MkString("orderbook/{market_string}") ,
                        MkString("market/{market_string}/ohlcv/{interval}") ,
                        })}),
                "private":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("me") ,
                        MkString("balances") ,
                        MkString("balances_all") ,
                        MkString("market/{market_string}") ,
                        MkString("orders") ,
                        MkString("order/{order_id}") ,
                        MkString("trades") ,
                        MkString("withdraw/{withdraw_id}") ,
                        MkString("withdraws") ,
                        MkString("deposit/{deposit_id}") ,
                        MkString("deposits") ,
                        MkString("transfers") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("cancel_order") ,
                        MkString("withdraw") ,
                        MkString("deposit_address/{currency}") ,
                        MkString("sell_limit") ,
                        MkString("buy_limit") ,
                        }),
                    }),
                }),
            "fees":MkMap(&VarMap{
                "trading":MkMap(&VarMap{
                    "tierBased":MkBool(true) ,
                    "percentage":MkBool(true) ,
                    "taker":MkNumber(0.0025) ,
                    "maker":MkNumber(0.0) ,
                    }),
                "funding":MkMap(&VarMap{"withdraw":MkMap(&VarMap{})}),
                }),
            "exceptions":MkMap(&VarMap{"exact":MkMap(&VarMap{
                    "invalid_auth":AuthenticationError ,
                    "insuff_funds":InsufficientFunds ,
                    "market_not_found":BadSymbol ,
                    "too_small":InvalidOrder ,
                    "limit_exceeded":RateLimitExceeded ,
                    })}),
            }));
}

func (this *Qtrade) FetchMarkets(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetMarkets"), params ); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  markets := this.SafeValue(data , MkString("markets") , MkArray(&VarArray{})); _ = markets;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , markets.Length )); OpIncr(& i ){
    market := *(markets ).At(i ); _ = market;
    marketId := this.SafeString(market , MkString("market_string") ); _ = marketId;
    numericId := this.SafeInteger(market , MkString("id") ); _ = numericId;
    baseId := this.SafeString(market , MkString("market_currency") ); _ = baseId;
    quoteId := this.SafeString(market , MkString("base_currency") ); _ = quoteId;
    base := this.SafeCurrencyCode(baseId ); _ = base;
    quote := this.SafeCurrencyCode(quoteId ); _ = quote;
    symbol := OpAdd(base , OpAdd(MkString("/") , quote )); _ = symbol;
    precision := MkMap(&VarMap{
          "amount":this.SafeInteger(market , MkString("market_precision") ),
          "price":this.SafeInteger(market , MkString("base_precision") ),
          }); _ = precision;
    canView := this.SafeValue(market , MkString("can_view") , MkBool(false) ); _ = canView;
    canTrade := this.SafeValue(market , MkString("can_trade") , MkBool(false) ); _ = canTrade;
    active := OpAnd(canTrade , canView ); _ = active;
    result.Push(MkMap(&VarMap{
            "symbol":symbol ,
            "id":marketId ,
            "numericId":numericId ,
            "baseId":baseId ,
            "quoteId":quoteId ,
            "base":base ,
            "quote":quote ,
            "active":active ,
            "precision":precision ,
            "taker":this.SafeNumber(market , MkString("taker_fee") ),
            "maker":this.SafeNumber(market , MkString("maker_fee") ),
            "limits":MkMap(&VarMap{
                "amount":MkMap(&VarMap{
                    "min":this.SafeNumber(market , MkString("minimum_sell_value") ),
                    "max":MkUndefined() ,
                    }),
                "price":MkMap(&VarMap{
                    "min":MkUndefined() ,
                    "max":MkUndefined() ,
                    }),
                "cost":MkMap(&VarMap{
                    "min":this.SafeNumber(market , MkString("minimum_buy_value") ),
                    "max":MkUndefined() ,
                    }),
                }),
            "info":market ,
            }));
  }
  return result ;
}

func (this *Qtrade) FetchCurrencies(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetCurrencies"), params ); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  currencies := this.SafeValue(data , MkString("currencies") , MkArray(&VarArray{})); _ = currencies;
  result := MkMap(&VarMap{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , currencies.Length )); OpIncr(& i ){
    currency := *(currencies ).At(i ); _ = currency;
    id := this.SafeString(currency , MkString("code") ); _ = id;
    code := this.SafeCurrencyCode(id ); _ = code;
    name := this.SafeString(currency , MkString("long_name") ); _ = name;
    type_ := this.SafeString(currency , MkString("type") ); _ = type_;
    canWithdraw := this.SafeValue(currency , MkString("can_withdraw") , MkBool(true) ); _ = canWithdraw;
    depositDisabled := this.SafeValue(currency , MkString("deposit_disabled") , MkBool(false) ); _ = depositDisabled;
    config := this.SafeValue(currency , MkString("config") , MkMap(&VarMap{})); _ = config;
    status := this.SafeString(currency , MkString("status") ); _ = status;
    active := OpAnd(canWithdraw , OpAnd(OpEq2(status , MkString("ok") ), OpNot(depositDisabled ))); _ = active;
    *(result ).At(code )= MkMap(&VarMap{
        "id":id ,
        "code":code ,
        "info":currency ,
        "type":type_ ,
        "name":name ,
        "fee":this.SafeNumber(config , MkString("withdraw_fee") ),
        "precision":this.SafeInteger(currency , MkString("precision") ),
        "active":active ,
        "limits":MkMap(&VarMap{
            "amount":MkMap(&VarMap{
                "min":this.SafeNumber(currency , MkString("minimum_order") ),
                "max":MkUndefined() ,
                }),
            "withdraw":MkMap(&VarMap{
                "min":MkUndefined() ,
                "max":MkUndefined() ,
                }),
            }),
        });
  }
  return result ;
}

func (this *Qtrade) ParseOHLCV(goArgs ...*Variant) *Variant{
  ohlcv := GoGetArg(goArgs, 0, MkUndefined()); _ = ohlcv;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  return MkArray(&VarArray{
        this.Parse8601(this.SafeString(ohlcv , MkString("time") )),
        this.SafeNumber(ohlcv , MkString("open") ),
        this.SafeNumber(ohlcv , MkString("high") ),
        this.SafeNumber(ohlcv , MkString("low") ),
        this.SafeNumber(ohlcv , MkString("close") ),
        this.SafeNumber(ohlcv , MkString("market_volume") ),
        });
}

func (this *Qtrade) FetchOHLCV(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  timeframe := GoGetArg(goArgs, 1, MkString("5m") ); _ = timeframe;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "market_string":*(market ).At(MkString("id") ),
        "interval":*(*this.At(MkString("timeframes")) ).At(timeframe ),
        }); _ = request;
  response := this.Call(MkString("publicGetMarketMarketStringOhlcvInterval"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  ohlcvs := this.SafeValue(data , MkString("slices") , MkArray(&VarArray{})); _ = ohlcvs;
  return this.ParseOHLCVs(ohlcvs , market , timeframe , since , limit );
}

func (this *Qtrade) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  marketId := this.MarketId(symbol ); _ = marketId;
  request := MkMap(&VarMap{"market_string":marketId }); _ = request;
  response := this.Call(MkString("publicGetOrderbookMarketString"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  orderbook := MkMap(&VarMap{}); _ = orderbook;
  sides := MkMap(&VarMap{
        "buy":MkString("bids") ,
        "sell":MkString("asks") ,
        }); _ = sides;
  keys := GoGetKeys(sides ); _ = keys;
  for i := MkInteger(0) ; IsTrue(OpLw(i , keys.Length )); OpIncr(& i ){
    key := *(keys ).At(i ); _ = key;
    side := *(sides ).At(key ); _ = side;
    bidasks := this.SafeValue(data , key , MkMap(&VarMap{})); _ = bidasks;
    prices := GoGetKeys(bidasks ); _ = prices;
    result := MkArray(&VarArray{}); _ = result;
    for j := MkInteger(0) ; IsTrue(OpLw(j , prices.Length )); OpIncr(& j ){
      priceAsString := *(prices ).At(j ); _ = priceAsString;
      price := this.SafeNumber(prices , j ); _ = price;
      amount := this.SafeNumber(bidasks , priceAsString ); _ = amount;
      result.Push(MkArray(&VarArray{
              price ,
              amount ,
              }));
    }
    *(orderbook ).At(side )= OpCopy(result );
  }
  timestamp := this.SafeIntegerProduct(data , MkString("last_change") , MkNumber(0.001) ); _ = timestamp;
  return this.ParseOrderBook(orderbook , symbol , timestamp );
}

func (this *Qtrade) ParseTicker(goArgs ...*Variant) *Variant{
  ticker := GoGetArg(goArgs, 0, MkUndefined()); _ = ticker;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  marketId := this.SafeString(ticker , MkString("id_hr") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market , MkString("_") ); _ = symbol;
  timestamp := this.SafeIntegerProduct(ticker , MkString("last_change") , MkNumber(0.001) ); _ = timestamp;
  previous := this.SafeNumber(ticker , MkString("day_open") ); _ = previous;
  last := this.SafeNumber(ticker , MkString("last") ); _ = last;
  day_change := this.SafeNumber(ticker , MkString("day_change") ); _ = day_change;
  percentage := OpCopy(MkUndefined() ); _ = percentage;
  change := OpCopy(MkUndefined() ); _ = change;
  average := this.SafeNumber(ticker , MkString("day_avg_price") ); _ = average;
  if IsTrue(OpNotEq2(day_change , MkUndefined() )) {
    percentage = OpMulti(day_change , MkInteger(100) );
    if IsTrue(OpNotEq2(previous , MkUndefined() )) {
      change = OpMulti(day_change , previous );
    }
  }
  if IsTrue(OpAnd(OpEq2(average , MkUndefined() ), OpAnd(OpNotEq2(last , MkUndefined() ), OpNotEq2(previous , MkUndefined() )))) {
    average = OpDiv(this.Sum(last , previous ), MkInteger(2) );
  }
  baseVolume := this.SafeNumber(ticker , MkString("day_volume_market") ); _ = baseVolume;
  quoteVolume := this.SafeNumber(ticker , MkString("day_volume_base") ); _ = quoteVolume;
  vwap := this.Vwap(baseVolume , quoteVolume ); _ = vwap;
  return MkMap(&VarMap{
        "symbol":symbol ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "high":this.SafeNumber(ticker , MkString("day_high") ),
        "low":this.SafeNumber(ticker , MkString("day_low") ),
        "bid":this.SafeNumber(ticker , MkString("bid") ),
        "bidVolume":MkUndefined() ,
        "ask":this.SafeNumber(ticker , MkString("ask") ),
        "askVolume":MkUndefined() ,
        "vwap":vwap ,
        "open":previous ,
        "close":last ,
        "last":last ,
        "previousClose":MkUndefined() ,
        "change":change ,
        "percentage":percentage ,
        "average":average ,
        "baseVolume":baseVolume ,
        "quoteVolume":quoteVolume ,
        "info":ticker ,
        });
}

func (this *Qtrade) FetchTickers(goArgs ...*Variant) *Variant{
  symbols := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbols;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("publicGetTickers"), params ); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  tickers := this.SafeValue(data , MkString("markets") , MkArray(&VarArray{})); _ = tickers;
  result := MkMap(&VarMap{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , tickers.Length )); OpIncr(& i ){
    ticker := this.ParseTicker(*(tickers ).At(i )); _ = ticker;
    symbol := *(ticker ).At(MkString("symbol") ); _ = symbol;
    *(result ).At(symbol )= OpCopy(ticker );
  }
  return this.FilterByArray(result , MkString("symbol") , symbols );
}

func (this *Qtrade) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"market_string":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("publicGetTickerMarketString"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  return this.ParseTicker(data , market );
}

func (this *Qtrade) FetchTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"market_string":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("publicGetMarketMarketStringTrades"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  trades := this.SafeValue(data , MkString("trades") , MkArray(&VarArray{})); _ = trades;
  return this.ParseTrades(trades , market , since , limit );
}

func (this *Qtrade) FetchMyTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"desc":MkBool(true) }); _ = request;
  market := OpCopy(MkUndefined() ); _ = market;
  numericId := this.SafeValue(params , MkString("market_id") ); _ = numericId;
  if IsTrue(OpNotEq2(numericId , MkUndefined() )) {
    *(request ).At(MkString("market_id") )= OpCopy(numericId );
  } else {
    if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
      market = this.Market(symbol );
      *(request ).At(MkString("market_string") )= *(market ).At(MkString("id") );
    }
  }
  response := this.Call(MkString("privateGetTrades"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  trades := this.SafeValue(data , MkString("trades") , MkArray(&VarArray{})); _ = trades;
  return this.ParseTrades(trades , market , since , limit );
}

func (this *Qtrade) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  id := this.SafeString(trade , MkString("id") ); _ = id;
  timestamp := this.SafeIntegerProduct(trade , MkString("created_at_ts") , MkNumber(0.001) ); _ = timestamp;
  if IsTrue(OpEq2(timestamp , MkUndefined() )) {
    timestamp = this.Parse8601(this.SafeString(trade , MkString("created_at") ));
  }
  side := this.SafeString(trade , MkString("side") ); _ = side;
  marketId := this.SafeString(trade , MkString("market_string") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market , MkString("_") ); _ = symbol;
  cost := this.SafeNumber2(trade , MkString("base_volume") , MkString("base_amount") ); _ = cost;
  priceString := this.SafeString(trade , MkString("price") ); _ = priceString;
  amountString := this.SafeString2(trade , MkString("market_amount") , MkString("amount") ); _ = amountString;
  price := this.ParseNumber(priceString ); _ = price;
  amount := this.ParseNumber(amountString ); _ = amount;
  if IsTrue(OpEq2(cost , MkUndefined() )) {
    cost = this.ParseNumber(Precise.StringMul(priceString , amountString ));
  }
  fee := OpCopy(MkUndefined() ); _ = fee;
  feeCost := this.SafeNumber(trade , MkString("base_fee") ); _ = feeCost;
  if IsTrue(OpNotEq2(feeCost , MkUndefined() )) {
    feeCurrencyCode := OpTrinary(OpEq2(market , MkUndefined() ), MkUndefined() , *(market ).At(MkString("quote") )); _ = feeCurrencyCode;
    fee = MkMap(&VarMap{
        "currency":feeCurrencyCode ,
        "cost":feeCost ,
        });
  }
  taker := this.SafeValue(trade , MkString("taker") , MkBool(true) ); _ = taker;
  takerOrMaker := OpTrinary(taker , MkString("taker") , MkString("maker") ); _ = takerOrMaker;
  orderId := this.SafeString(trade , MkString("order_id") ); _ = orderId;
  result := MkMap(&VarMap{
        "id":id ,
        "info":trade ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "symbol":symbol ,
        "order":orderId ,
        "type":MkUndefined() ,
        "side":side ,
        "takerOrMaker":takerOrMaker ,
        "price":price ,
        "amount":amount ,
        "cost":cost ,
        "fee":fee ,
        }); _ = result;
  return result ;
}

func (this *Qtrade) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("privateGetBalancesAll"), params ); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  balances := this.SafeValue(data , MkString("balances") , MkArray(&VarArray{})); _ = balances;
  result := MkMap(&VarMap{
        "info":response ,
        "timestamp":MkUndefined() ,
        "datetime":MkUndefined() ,
        }); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , balances.Length )); OpIncr(& i ){
    balance := *(balances ).At(i ); _ = balance;
    currencyId := this.SafeString(balance , MkString("currency") ); _ = currencyId;
    code := this.SafeCurrencyCode(currencyId ); _ = code;
    account := OpTrinary(OpHasMember(code , result ), *(result ).At(code ), this.Account()); _ = account;
    *(account ).At(MkString("free") )= this.SafeString(balance , MkString("balance") );
    *(account ).At(MkString("used") )= MkString("0") ;
    *(result ).At(code )= OpCopy(account );
  }
  balances = this.SafeValue(data , MkString("order_balances") , MkArray(&VarArray{}));
  for i := MkInteger(0) ; IsTrue(OpLw(i , balances.Length )); OpIncr(& i ){
    balance := *(balances ).At(i ); _ = balance;
    currencyId := this.SafeString(balance , MkString("currency") ); _ = currencyId;
    code := this.SafeCurrencyCode(currencyId ); _ = code;
    account := OpTrinary(OpHasMember(code , result ), *(result ).At(code ), this.Account()); _ = account;
    *(account ).At(MkString("used") )= this.SafeString(balance , MkString("balance") );
    *(result ).At(code )= OpCopy(account );
  }
  return this.ParseBalance(result );
}

func (this *Qtrade) CreateOrder(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 2, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 3, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 4, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 5, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpNotEq2(type_ , MkString("limit") )) {
    panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")) , MkString(" createOrder() allows limit orders only") )));
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "amount":this.AmountToPrecision(symbol , amount ),
        "market_id":*(market ).At(MkString("numericId") ),
        "price":this.PriceToPrecision(symbol , price ),
        }); _ = request;
  method := OpTrinary(OpEq2(side , MkString("sell") ), MkString("privatePostSellLimit") , MkString("privatePostBuyLimit") ); _ = method;
  response := this.Call(method , this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  order := this.SafeValue(data , MkString("order") , MkMap(&VarMap{})); _ = order;
  return this.ParseOrder(order , market );
}

func (this *Qtrade) ParseOrder(goArgs ...*Variant) *Variant{
  order := GoGetArg(goArgs, 0, MkUndefined()); _ = order;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  id := this.SafeString(order , MkString("id") ); _ = id;
  timestamp := this.Parse8601(this.SafeString(order , MkString("created_at") )); _ = timestamp;
  sideType := this.SafeString(order , MkString("order_type") ); _ = sideType;
  orderType := OpCopy(MkUndefined() ); _ = orderType;
  side := OpCopy(MkUndefined() ); _ = side;
  if IsTrue(OpNotEq2(sideType , MkUndefined() )) {
    parts := sideType.Split(MkString("_") ); _ = parts;
    side = this.SafeString(parts , MkInteger(0) );
    orderType = this.SafeString(parts , MkInteger(1) );
  }
  price := this.SafeNumber(order , MkString("price") ); _ = price;
  amount := this.SafeNumber(order , MkString("market_amount") ); _ = amount;
  remaining := this.SafeNumber(order , MkString("market_amount_remaining") ); _ = remaining;
  open := this.SafeValue(order , MkString("open") , MkBool(false) ); _ = open;
  closeReason := this.SafeString(order , MkString("close_reason") ); _ = closeReason;
  status := OpCopy(MkUndefined() ); _ = status;
  if IsTrue(open ) {
    status = MkString("open") ;
  } else {
    if IsTrue(OpEq2(closeReason , MkString("canceled") )) {
      status = MkString("canceled") ;
    } else {
      status = MkString("closed") ;
    }
  }
  marketId := this.SafeString(order , MkString("market_string") ); _ = marketId;
  market = this.SafeMarket(marketId , market , MkString("_") );
  symbol := *(market ).At(MkString("symbol") ); _ = symbol;
  rawTrades := this.SafeValue(order , MkString("trades") , MkArray(&VarArray{})); _ = rawTrades;
  parsedTrades := this.ParseTrades(rawTrades , market , MkUndefined() , MkUndefined() , MkMap(&VarMap{
            "order":id ,
            "side":side ,
            "type":orderType ,
            })); _ = parsedTrades;
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
            "average":MkUndefined() ,
            "amount":amount ,
            "remaining":remaining ,
            "filled":MkUndefined() ,
            "status":status ,
            "fee":MkUndefined() ,
            "fees":MkUndefined() ,
            "cost":MkUndefined() ,
            "trades":parsedTrades ,
            }));
}

func (this *Qtrade) CancelOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"id":parseInt(id )}); _ = request;
  return this.Call(MkString("privatePostCancelOrder"), this.Extend(request , params ));
}

func (this *Qtrade) FetchOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"order_id":id }); _ = request;
  response := this.Call(MkString("privateGetOrderOrderId"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  order := this.SafeValue(data , MkString("order") , MkMap(&VarMap{})); _ = order;
  return this.ParseOrder(order );
}

func (this *Qtrade) FetchOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  market := OpCopy(MkUndefined() ); _ = market;
  numericId := this.SafeValue(params , MkString("market_id") ); _ = numericId;
  if IsTrue(OpNotEq2(numericId , MkUndefined() )) {
    *(request ).At(MkString("market_id") )= OpCopy(numericId );
  } else {
    if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
      market = this.Market(symbol );
      *(request ).At(MkString("market_string") )= *(market ).At(MkString("id") );
    }
  }
  response := this.Call(MkString("privateGetOrders"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  orders := this.SafeValue(data , MkString("orders") , MkArray(&VarArray{})); _ = orders;
  return this.ParseOrders(orders , market , since , limit );
}

func (this *Qtrade) FetchOpenOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"open":MkBool(true) }); _ = request;
  return this.FetchOrders(symbol , since , limit , this.Extend(request , params ));
}

func (this *Qtrade) FetchClosedOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"open":MkBool(false) }); _ = request;
  return this.FetchOrders(symbol , since , limit , this.Extend(request , params ));
}

func (this *Qtrade) ParseDepositAddress(goArgs ...*Variant) *Variant{
  depositAddress := GoGetArg(goArgs, 0, MkUndefined()); _ = depositAddress;
  currency := GoGetArg(goArgs, 1, MkUndefined() ); _ = currency;
  code := OpTrinary(OpEq2(currency , MkUndefined() ), MkUndefined() , *(currency ).At(MkString("code") )); _ = code;
  address := this.SafeString(depositAddress , MkString("address") ); _ = address;
  tag := OpCopy(MkUndefined() ); _ = tag;
  if IsTrue(OpNotEq2(address , MkUndefined() )) {
    parts := address.Split(MkString(":") ); _ = parts;
    address = this.SafeString(parts , MkInteger(0) );
    tag = this.SafeString(parts , MkInteger(1) );
  }
  this.CheckAddress(address );
  return MkMap(&VarMap{
        "currency":code ,
        "address":address ,
        "tag":tag ,
        "info":depositAddress ,
        });
}

func (this *Qtrade) FetchDepositAddress(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  currency := this.Currency(code ); _ = currency;
  request := MkMap(&VarMap{"currency":*(currency ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("privatePostDepositAddressCurrency"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  return this.ParseDepositAddress(data , currency );
}

func (this *Qtrade) FetchDeposit(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  code := GoGetArg(goArgs, 1, MkUndefined() ); _ = code;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"deposit_id":id }); _ = request;
  response := this.Call(MkString("privateGetDepositDepositId"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  deposit := this.SafeValue(data , MkString("deposit") , MkMap(&VarMap{})); _ = deposit;
  return this.ParseTransaction(deposit );
}

func (this *Qtrade) FetchDeposits(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  currency := OpCopy(MkUndefined() ); _ = currency;
  if IsTrue(OpNotEq2(code , MkUndefined() )) {
    currency = this.Currency(code );
  }
  response := this.Call(MkString("privateGetDeposits"), params ); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  deposits := this.SafeValue(data , MkString("deposits") , MkArray(&VarArray{})); _ = deposits;
  return this.ParseTransactions(deposits , currency , since , limit );
}

func (this *Qtrade) FetchWithdrawal(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  code := GoGetArg(goArgs, 1, MkUndefined() ); _ = code;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"withdraw_id":id }); _ = request;
  response := this.Call(MkString("privateGetWithdrawWithdrawId"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  withdrawal := this.SafeValue(data , MkString("withdraw") , MkMap(&VarMap{})); _ = withdrawal;
  return this.ParseTransaction(withdrawal );
}

func (this *Qtrade) FetchWithdrawals(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  currency := OpCopy(MkUndefined() ); _ = currency;
  if IsTrue(OpNotEq2(code , MkUndefined() )) {
    currency = this.Currency(code );
  }
  response := this.Call(MkString("privateGetWithdraws"), params ); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  withdrawals := this.SafeValue(data , MkString("withdraws") , MkArray(&VarArray{})); _ = withdrawals;
  return this.ParseTransactions(withdrawals , currency , since , limit );
}

func (this *Qtrade) ParseTransaction(goArgs ...*Variant) *Variant{
  transaction := GoGetArg(goArgs, 0, MkUndefined()); _ = transaction;
  currency := GoGetArg(goArgs, 1, MkUndefined() ); _ = currency;
  timestamp := this.Parse8601(this.SafeString(transaction , MkString("created_at") )); _ = timestamp;
  id := this.SafeString(transaction , MkString("id") ); _ = id;
  networkData := this.SafeValue(transaction , MkString("network_data") , MkMap(&VarMap{})); _ = networkData;
  unsignedTx := this.SafeValue(networkData , MkString("unsigned_tx") , MkMap(&VarMap{})); _ = unsignedTx;
  addressFrom := this.SafeString(unsignedTx , MkString("from") ); _ = addressFrom;
  txid := this.SafeString(networkData , MkString("txid") ); _ = txid;
  address := this.SafeString(transaction , MkString("address") ); _ = address;
  tag := OpCopy(MkUndefined() ); _ = tag;
  if IsTrue(OpNotEq2(address , MkUndefined() )) {
    parts := address.Split(MkString(":") ); _ = parts;
    numParts := OpCopy(parts.Length ); _ = numParts;
    if IsTrue(OpGt(numParts , MkInteger(1) )) {
      address = this.SafeString(parts , MkInteger(0) );
      tag = this.SafeString(parts , MkInteger(1) );
    }
  }
  addressTo := OpCopy(address ); _ = addressTo;
  tagFrom := OpCopy(MkUndefined() ); _ = tagFrom;
  tagTo := OpCopy(tag ); _ = tagTo;
  cancelRequested := this.SafeValue(transaction , MkString("cancel_requested") ); _ = cancelRequested;
  type_ := OpTrinary(OpEq2(cancelRequested , MkUndefined() ), MkString("deposit") , MkString("withdrawal") ); _ = type_;
  amount := this.SafeNumber(transaction , MkString("amount") ); _ = amount;
  currencyId := this.SafeString(transaction , MkString("currency") ); _ = currencyId;
  code := this.SafeCurrencyCode(currencyId ); _ = code;
  status := this.ParseTransactionStatus(this.SafeString(transaction , MkString("status") )); _ = status;
  statusCode := this.SafeString(transaction , MkString("code") ); _ = statusCode;
  if IsTrue(cancelRequested ) {
    status = MkString("canceled") ;
  } else {
    if IsTrue(OpEq2(status , MkUndefined() )) {
      status = this.ParseTransactionStatus(statusCode );
    }
  }
  fee := OpCopy(MkUndefined() ); _ = fee;
  return MkMap(&VarMap{
        "info":transaction ,
        "id":id ,
        "txid":txid ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "addressFrom":addressFrom ,
        "addressTo":addressTo ,
        "address":address ,
        "tagFrom":tagFrom ,
        "tagTo":tagTo ,
        "tag":tag ,
        "type":type_ ,
        "amount":amount ,
        "currency":code ,
        "status":status ,
        "updated":MkUndefined() ,
        "fee":fee ,
        });
}

func (this *Qtrade) ParseTransactionStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "initiated":MkString("pending") ,
        "needs_create":MkString("pending") ,
        "credited":MkString("ok") ,
        "confirmed":MkString("ok") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Qtrade) Withdraw(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  amount := GoGetArg(goArgs, 1, MkUndefined()); _ = amount;
  address := GoGetArg(goArgs, 2, MkUndefined()); _ = address;
  tag := GoGetArg(goArgs, 3, MkUndefined() ); _ = tag;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  currency := this.Currency(code ); _ = currency;
  request := MkMap(&VarMap{
        "address":address ,
        "amount":amount ,
        "currency":*(currency ).At(MkString("id") ),
        }); _ = request;
  if IsTrue(OpNotEq2(tag , MkUndefined() )) {
    OpAddAssign(& *(request ).At(MkString("address") ), OpAdd(MkString(":") , tag ));
  }
  response := this.Call(MkString("privatePostWithdraw"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  result := this.ParseTransaction(data ); _ = result;
  return this.Extend(result , MkMap(&VarMap{
            "currency":code ,
            "address":address ,
            "addressTo":address ,
            "tag":tag ,
            "tagTo":tag ,
            "amount":amount ,
            }));
}

func (this *Qtrade) Nonce(goArgs ...*Variant) *Variant{
  return this.Milliseconds();
}

func (this *Qtrade) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("public") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  url := OpAdd(MkString("/") , OpAdd(*this.At(MkString("version")) , MkString("/") )); _ = url;
  if IsTrue(OpEq2(api , MkString("private") )) {
    OpAddAssign(& url , MkString("user/") );
  }
  OpAddAssign(& url , this.ImplodeParams(path , params ));
  request := this.Omit(params , this.ExtractParams(path )); _ = request;
  if IsTrue(OpEq2(method , MkString("POST") )) {
    body = this.Json(request );
  } else {
    if IsTrue(*(GoGetKeys(request )).At(MkString("length") )) {
      OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(request )));
    }
  }
  if IsTrue(OpEq2(api , MkString("private") )) {
    timestamp := (this.Milliseconds()).Call(MkString("toString") ); _ = timestamp;
    bodyAsString := OpTrinary(OpEq2(method , MkString("POST") ), body , MkString("") ); _ = bodyAsString;
    auth := (MkArray(&VarArray{
            method ,
            url ,
            timestamp ,
            bodyAsString ,
            *this.At(MkString("secret")) ,
            })).Call(MkString("join") , MkString("\n") ); _ = auth;
    hash := this.Hash(this.Encode(auth ), MkString("sha256") , MkString("base64") ); _ = hash;
    key := OpCopy(*this.At(MkString("apiKey")) ); _ = key;
    if IsTrue(OpNotEq2(OpType(key ), MkString("string") )) {
      key = key.ToString();
    }
    signature := OpAdd(MkString("HMAC-SHA256 ") , OpAdd(key , OpAdd(MkString(":") , hash ))); _ = signature;
    headers = MkMap(&VarMap{
        "Authorization":signature ,
        "HMAC-Timestamp":timestamp ,
        });
    if IsTrue(OpEq2(method , MkString("POST") )) {
      *(headers ).At(MkString("Content-Type") )= MkString("application/json") ;
    }
  }
  url = OpAdd(*(*this.At(MkString("urls")) ).At(MkString("api") ), url );
  return MkMap(&VarMap{
        "url":url ,
        "method":method ,
        "body":body ,
        "headers":headers ,
        });
}

func (this *Qtrade) HandleErrors(goArgs ...*Variant) *Variant{
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
  errors := this.SafeValue(response , MkString("errors") , MkArray(&VarArray{})); _ = errors;
  numErrors := OpCopy(errors.Length ); _ = numErrors;
  if IsTrue(OpLw(numErrors , MkInteger(1) )) {
    return MkUndefined();
  }
  feedback := OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , body )); _ = feedback;
  for i := MkInteger(0) ; IsTrue(OpLw(i , errors.Length )); OpIncr(& i ){
    error := *(errors ).At(i ); _ = error;
    errorCode := this.SafeString(error , MkString("code") ); _ = errorCode;
    this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("exact") ), errorCode , feedback );
  }
  panic(NewExchangeError(feedback ));
  return MkUndefined()
}

