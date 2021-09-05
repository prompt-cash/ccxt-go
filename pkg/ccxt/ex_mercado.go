package ccxt

import (
)

type Mercado struct {
	*ExchangeBase
}

var _ Exchange = (*Mercado)(nil)

func init() {
	exchange := &Mercado{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Mercado) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("mercado") ,
            "name":MkString("Mercado Bitcoin") ,
            "countries":MkArray(&VarArray{
                MkString("BR") ,
                }),
            "rateLimit":MkInteger(1000) ,
            "version":MkString("v3") ,
            "has":MkMap(&VarMap{
                "cancelOrder":MkBool(true) ,
                "CORS":MkBool(true) ,
                "createMarketOrder":MkBool(true) ,
                "createOrder":MkBool(true) ,
                "fetchBalance":MkBool(true) ,
                "fetchMarkets":MkBool(true) ,
                "fetchMyTrades":MkString("emulated") ,
                "fetchOHLCV":MkBool(true) ,
                "fetchOpenOrders":MkBool(true) ,
                "fetchOrder":MkBool(true) ,
                "fetchOrderBook":MkBool(true) ,
                "fetchOrders":MkBool(true) ,
                "fetchTicker":MkBool(true) ,
                "fetchTickers":MkBool(false) ,
                "fetchTrades":MkBool(true) ,
                "withdraw":MkBool(true) ,
                }),
            "timeframes":MkMap(&VarMap{
                "1m":MkString("1m") ,
                "5m":MkString("5m") ,
                "15m":MkString("15m") ,
                "30m":MkString("30m") ,
                "1h":MkString("1h") ,
                "6h":MkString("6h") ,
                "12h":MkString("12h") ,
                "1d":MkString("1d") ,
                "3d":MkString("3d") ,
                "1w":MkString("1w") ,
                "2w":MkString("2w") ,
                }),
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/1294454/27837060-e7c58714-60ea-11e7-9192-f05e86adb83f.jpg") ,
                "api":MkMap(&VarMap{
                    "public":MkString("https://www.mercadobitcoin.net/api") ,
                    "private":MkString("https://www.mercadobitcoin.net/tapi") ,
                    "v4Public":MkString("https://www.mercadobitcoin.com.br/v4") ,
                    }),
                "www":MkString("https://www.mercadobitcoin.com.br") ,
                "doc":MkArray(&VarArray{
                    MkString("https://www.mercadobitcoin.com.br/api-doc") ,
                    MkString("https://www.mercadobitcoin.com.br/trade-api") ,
                    }),
                }),
            "api":MkMap(&VarMap{
                "public":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("coins") ,
                        MkString("{coin}/orderbook/") ,
                        MkString("{coin}/ticker/") ,
                        MkString("{coin}/trades/") ,
                        MkString("{coin}/trades/{from}/") ,
                        MkString("{coin}/trades/{from}/{to}") ,
                        MkString("{coin}/day-summary/{year}/{month}/{day}/") ,
                        })}),
                "private":MkMap(&VarMap{"post":MkArray(&VarArray{
                        MkString("cancel_order") ,
                        MkString("get_account_info") ,
                        MkString("get_order") ,
                        MkString("get_withdrawal") ,
                        MkString("list_system_messages") ,
                        MkString("list_orders") ,
                        MkString("list_orderbook") ,
                        MkString("place_buy_order") ,
                        MkString("place_sell_order") ,
                        MkString("place_market_buy_order") ,
                        MkString("place_market_sell_order") ,
                        MkString("withdraw_coin") ,
                        })}),
                "v4Public":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("{coin}/candle/") ,
                        })}),
                }),
            "fees":MkMap(&VarMap{"trading":MkMap(&VarMap{
                    "maker":MkNumber(0.003) ,
                    "taker":MkNumber(0.007) ,
                    })}),
            "options":MkMap(&VarMap{"limits":MkMap(&VarMap{
                    "BTC":MkNumber(0.001) ,
                    "BCH":MkNumber(0.001) ,
                    "ETH":MkNumber(0.01) ,
                    "LTC":MkNumber(0.01) ,
                    "XRP":MkNumber(0.1) ,
                    })}),
            }));
}

func (this *Mercado) FetchMarkets(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetCoins"), params ); _ = response;
  result := MkArray(&VarArray{}); _ = result;
  amountLimits := this.SafeValue(*this.At(MkString("options")) , MkString("limits") , MkMap(&VarMap{})); _ = amountLimits;
  for i := MkInteger(0) ; IsTrue(OpLw(i , response.Length )); OpIncr(& i ){
    coin := *(response ).At(i ); _ = coin;
    baseId := OpCopy(coin ); _ = baseId;
    quoteId := MkString("BRL") ; _ = quoteId;
    base := this.SafeCurrencyCode(baseId ); _ = base;
    quote := this.SafeCurrencyCode(quoteId ); _ = quote;
    symbol := OpAdd(base , OpAdd(MkString("/") , quote )); _ = symbol;
    id := OpAdd(quote , base ); _ = id;
    precision := MkMap(&VarMap{
          "amount":MkInteger(8) ,
          "price":MkInteger(5) ,
          }); _ = precision;
    priceLimit := MkString("1e-5") ; _ = priceLimit;
    result.Push(MkMap(&VarMap{
            "id":id ,
            "symbol":symbol ,
            "base":base ,
            "quote":quote ,
            "baseId":baseId ,
            "quoteId":quoteId ,
            "active":MkUndefined() ,
            "info":coin ,
            "precision":precision ,
            "limits":MkMap(&VarMap{
                "amount":MkMap(&VarMap{
                    "min":this.SafeNumber(amountLimits , baseId ),
                    "max":MkUndefined() ,
                    }),
                "price":MkMap(&VarMap{
                    "min":this.ParseNumber(priceLimit ),
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

func (this *Mercado) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"coin":*(market ).At(MkString("base") )}); _ = request;
  response := this.Call(MkString("publicGetCoinOrderbook"), this.Extend(request , params )); _ = response;
  return this.ParseOrderBook(response , symbol );
}

func (this *Mercado) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"coin":*(market ).At(MkString("base") )}); _ = request;
  response := this.Call(MkString("publicGetCoinTicker"), this.Extend(request , params )); _ = response;
  ticker := this.SafeValue(response , MkString("ticker") , MkMap(&VarMap{})); _ = ticker;
  timestamp := this.SafeTimestamp(ticker , MkString("date") ); _ = timestamp;
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
        "average":MkUndefined() ,
        "baseVolume":this.SafeNumber(ticker , MkString("vol") ),
        "quoteVolume":MkUndefined() ,
        "info":ticker ,
        });
}

func (this *Mercado) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.SafeTimestamp2(trade , MkString("date") , MkString("executed_timestamp") ); _ = timestamp;
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  if IsTrue(OpNotEq2(market , MkUndefined() )) {
    symbol = *(market ).At(MkString("symbol") );
  }
  id := this.SafeString2(trade , MkString("tid") , MkString("operation_id") ); _ = id;
  type_ := OpCopy(MkUndefined() ); _ = type_;
  side := this.SafeString(trade , MkString("type") ); _ = side;
  priceString := this.SafeString(trade , MkString("price") ); _ = priceString;
  amountString := this.SafeString2(trade , MkString("amount") , MkString("quantity") ); _ = amountString;
  price := this.ParseNumber(priceString ); _ = price;
  amount := this.ParseNumber(amountString ); _ = amount;
  cost := this.ParseNumber(Precise.StringMul(priceString , amountString )); _ = cost;
  feeCost := this.SafeNumber(trade , MkString("fee_rate") ); _ = feeCost;
  fee := OpCopy(MkUndefined() ); _ = fee;
  if IsTrue(OpNotEq2(feeCost , MkUndefined() )) {
    fee = MkMap(&VarMap{
        "cost":feeCost ,
        "currency":MkUndefined() ,
        });
  }
  return MkMap(&VarMap{
        "id":id ,
        "info":trade ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "symbol":symbol ,
        "order":MkUndefined() ,
        "type":type_ ,
        "side":side ,
        "takerOrMaker":MkUndefined() ,
        "price":price ,
        "amount":amount ,
        "cost":cost ,
        "fee":fee ,
        });
}

func (this *Mercado) FetchTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  method := MkString("publicGetCoinTrades") ; _ = method;
  request := MkMap(&VarMap{"coin":*(market ).At(MkString("base") )}); _ = request;
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    OpAddAssign(& method , MkString("From") );
    *(request ).At(MkString("from") )= parseInt(OpDiv(since , MkInteger(1000) ));
  }
  to := this.SafeInteger(params , MkString("to") ); _ = to;
  if IsTrue(OpNotEq2(to , MkUndefined() )) {
    OpAddAssign(& method , MkString("To") );
  }
  response := this.Call(method , this.Extend(request , params )); _ = response;
  return this.ParseTrades(response , market , since , limit );
}

func (this *Mercado) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("privatePostGetAccountInfo"), params ); _ = response;
  data := this.SafeValue(response , MkString("response_data") , MkMap(&VarMap{})); _ = data;
  balances := this.SafeValue(data , MkString("balance") , MkMap(&VarMap{})); _ = balances;
  result := MkMap(&VarMap{"info":response }); _ = result;
  currencyIds := GoGetKeys(balances ); _ = currencyIds;
  for i := MkInteger(0) ; IsTrue(OpLw(i , currencyIds.Length )); OpIncr(& i ){
    currencyId := *(currencyIds ).At(i ); _ = currencyId;
    code := this.SafeCurrencyCode(currencyId ); _ = code;
    if IsTrue(OpHasMember(currencyId , balances )) {
      balance := this.SafeValue(balances , currencyId , MkMap(&VarMap{})); _ = balance;
      account := this.Account(); _ = account;
      *(account ).At(MkString("free") )= this.SafeString(balance , MkString("available") );
      *(account ).At(MkString("total") )= this.SafeString(balance , MkString("total") );
      *(result ).At(code )= OpCopy(account );
    }
  }
  return this.ParseBalance(result );
}

func (this *Mercado) CreateOrder(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 2, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 3, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 4, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 5, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"coin_pair":this.MarketId(symbol )}); _ = request;
  method := OpAdd(this.Capitalize(side ), MkString("Order") ); _ = method;
  if IsTrue(OpEq2(type_ , MkString("limit") )) {
    method = OpAdd(MkString("privatePostPlace") , method );
    *(request ).At(MkString("limit_price") )= this.PriceToPrecision(symbol , price );
    *(request ).At(MkString("quantity") )= this.AmountToPrecision(symbol , amount );
  } else {
    method = OpAdd(MkString("privatePostPlaceMarket") , method );
    if IsTrue(OpEq2(side , MkString("buy") )) {
      if IsTrue(OpEq2(price , MkUndefined() )) {
        panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")) , MkString(" createOrder() requires the price argument with market buy orders to calculate total order cost (amount to spend), where cost = amount * price. Supply a price argument to createOrder() call if you want the cost to be calculated for you from price and amount") )));
      }
      *(request ).At(MkString("cost") )= this.PriceToPrecision(symbol , OpMulti(amount , price ));
    } else {
      *(request ).At(MkString("quantity") )= this.AmountToPrecision(symbol , amount );
    }
  }
  response := this.Call(method , this.Extend(request , params )); _ = response;
  return MkMap(&VarMap{
        "info":response ,
        "id":(*(*(*(response ).At(MkString("response_data") )).At(MkString("order") )).At(MkString("order_id") )).Call(MkString("toString") ),
        });
}

func (this *Mercado) CancelOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" cancelOrder () requires a symbol argument") )));
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "coin_pair":*(market ).At(MkString("id") ),
        "order_id":id ,
        }); _ = request;
  response := this.Call(MkString("privatePostCancelOrder"), this.Extend(request , params )); _ = response;
  responseData := this.SafeValue(response , MkString("response_data") , MkMap(&VarMap{})); _ = responseData;
  order := this.SafeValue(responseData , MkString("order") , MkMap(&VarMap{})); _ = order;
  return this.ParseOrder(order , market );
}

func (this *Mercado) ParseOrderStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "2":MkString("open") ,
        "3":MkString("canceled") ,
        "4":MkString("closed") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Mercado) ParseOrder(goArgs ...*Variant) *Variant{
  order := GoGetArg(goArgs, 0, MkUndefined()); _ = order;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  id := this.SafeString(order , MkString("order_id") ); _ = id;
  order_type := this.SafeString(order , MkString("order_type") ); _ = order_type;
  side := OpCopy(MkUndefined() ); _ = side;
  if IsTrue(OpHasMember(MkString("order_type") , order )) {
    side = OpTrinary(OpEq2(order_type , MkString("1") ), MkString("buy") , MkString("sell") );
  }
  status := this.ParseOrderStatus(this.SafeString(order , MkString("status") )); _ = status;
  marketId := this.SafeString(order , MkString("coin_pair") ); _ = marketId;
  market = this.SafeMarket(marketId , market );
  timestamp := this.SafeTimestamp(order , MkString("created_timestamp") ); _ = timestamp;
  fee := MkMap(&VarMap{
        "cost":this.SafeNumber(order , MkString("fee") ),
        "currency":*(market ).At(MkString("quote") ),
        }); _ = fee;
  price := this.SafeNumber(order , MkString("limit_price") ); _ = price;
  average := this.SafeNumber(order , MkString("executed_price_avg") ); _ = average;
  amount := this.SafeNumber(order , MkString("quantity") ); _ = amount;
  filled := this.SafeNumber(order , MkString("executed_quantity") ); _ = filled;
  lastTradeTimestamp := this.SafeTimestamp(order , MkString("updated_timestamp") ); _ = lastTradeTimestamp;
  rawTrades := this.SafeValue(order , MkString("operations") , MkArray(&VarArray{})); _ = rawTrades;
  trades := this.ParseTrades(rawTrades , market , MkUndefined() , MkUndefined() , MkMap(&VarMap{
            "side":side ,
            "order":id ,
            })); _ = trades;
  return this.SafeOrder(MkMap(&VarMap{
            "info":order ,
            "id":id ,
            "clientOrderId":MkUndefined() ,
            "timestamp":timestamp ,
            "datetime":this.Iso8601(timestamp ),
            "lastTradeTimestamp":lastTradeTimestamp ,
            "symbol":*(market ).At(MkString("symbol") ),
            "type":MkString("limit") ,
            "timeInForce":MkUndefined() ,
            "postOnly":MkUndefined() ,
            "side":side ,
            "price":price ,
            "stopPrice":MkUndefined() ,
            "cost":MkUndefined() ,
            "average":average ,
            "amount":amount ,
            "filled":filled ,
            "remaining":MkUndefined() ,
            "status":status ,
            "fee":fee ,
            "trades":trades ,
            }));
}

func (this *Mercado) FetchOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchOrder () requires a symbol argument") )));
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "coin_pair":*(market ).At(MkString("id") ),
        "order_id":parseInt(id ),
        }); _ = request;
  response := this.Call(MkString("privatePostGetOrder"), this.Extend(request , params )); _ = response;
  responseData := this.SafeValue(response , MkString("response_data") , MkMap(&VarMap{})); _ = responseData;
  order := this.SafeValue(responseData , MkString("order") ); _ = order;
  return this.ParseOrder(order , market );
}

func (this *Mercado) Withdraw(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  amount := GoGetArg(goArgs, 1, MkUndefined()); _ = amount;
  address := GoGetArg(goArgs, 2, MkUndefined()); _ = address;
  tag := GoGetArg(goArgs, 3, MkUndefined() ); _ = tag;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.CheckAddress(address );
  this.LoadMarkets();
  currency := this.Currency(code ); _ = currency;
  request := MkMap(&VarMap{
        "coin":*(currency ).At(MkString("id") ),
        "quantity":amount.ToFixed(MkInteger(10) ),
        "address":address ,
        }); _ = request;
  if IsTrue(OpEq2(code , MkString("BRL") )) {
    account_ref := OpHasMember(MkString("account_ref") , params ); _ = account_ref;
    if IsTrue(OpNot(account_ref )) {
      panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" withdraw() requires account_ref parameter to withdraw ") , code ))));
    }
  } else {
    if IsTrue(OpNotEq2(code , MkString("LTC") )) {
      tx_fee := OpHasMember(MkString("tx_fee") , params ); _ = tx_fee;
      if IsTrue(OpNot(tx_fee )) {
        panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" withdraw() requires tx_fee parameter to withdraw ") , code ))));
      }
      if IsTrue(OpEq2(code , MkString("XRP") )) {
        if IsTrue(OpEq2(tag , MkUndefined() )) {
          if IsTrue(OpNot(OpHasMember(MkString("destination_tag") , params ))) {
            panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" withdraw() requires a tag argument or destination_tag parameter to withdraw ") , code ))));
          }
        } else {
          *(request ).At(MkString("destination_tag") )= OpCopy(tag );
        }
      }
    }
  }
  response := this.Call(MkString("privatePostWithdrawCoin"), this.Extend(request , params )); _ = response;
  return MkMap(&VarMap{
        "info":response ,
        "id":*(*(*(response ).At(MkString("response_data") )).At(MkString("withdrawal") )).At(MkString("id") ),
        });
}

func (this *Mercado) ParseOHLCV(goArgs ...*Variant) *Variant{
  ohlcv := GoGetArg(goArgs, 0, MkUndefined()); _ = ohlcv;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  return MkArray(&VarArray{
        this.SafeTimestamp(ohlcv , MkString("timestamp") ),
        this.SafeNumber(ohlcv , MkString("open") ),
        this.SafeNumber(ohlcv , MkString("high") ),
        this.SafeNumber(ohlcv , MkString("low") ),
        this.SafeNumber(ohlcv , MkString("close") ),
        this.SafeNumber(ohlcv , MkString("volume") ),
        });
}

func (this *Mercado) FetchOHLCV(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  timeframe := GoGetArg(goArgs, 1, MkString("5m") ); _ = timeframe;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "precision":*(*this.At(MkString("timeframes")) ).At(timeframe ),
        "coin":(*(market ).At(MkString("id") )).Call(MkString("toLowerCase") ),
        }); _ = request;
  if IsTrue(OpAnd(OpNotEq2(limit , MkUndefined() ), OpNotEq2(since , MkUndefined() ))) {
    *(request ).At(MkString("from") )= parseInt(OpDiv(since , MkInteger(1000) ));
    *(request ).At(MkString("to") )= this.Sum(*(request ).At(MkString("from") ), OpMulti(limit , this.ParseTimeframe(timeframe )));
  } else {
    if IsTrue(OpNotEq2(since , MkUndefined() )) {
      *(request ).At(MkString("from") )= parseInt(OpDiv(since , MkInteger(1000) ));
      *(request ).At(MkString("to") )= this.Sum(this.Seconds(), MkInteger(1) );
    } else {
      if IsTrue(OpNotEq2(limit , MkUndefined() )) {
        *(request ).At(MkString("to") )= this.Seconds();
        *(request ).At(MkString("from") )= OpSub(*(request ).At(MkString("to") ), OpMulti(limit , this.ParseTimeframe(timeframe )));
      }
    }
  }
  response := this.Call(MkString("v4PublicGetCoinCandle"), this.Extend(request , params )); _ = response;
  candles := this.SafeValue(response , MkString("candles") , MkArray(&VarArray{})); _ = candles;
  return this.ParseOHLCVs(candles , market , timeframe , since , limit );
}

func (this *Mercado) FetchOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchOrders () requires a symbol argument") )));
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"coin_pair":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("privatePostListOrders"), this.Extend(request , params )); _ = response;
  responseData := this.SafeValue(response , MkString("response_data") , MkMap(&VarMap{})); _ = responseData;
  orders := this.SafeValue(responseData , MkString("orders") , MkArray(&VarArray{})); _ = orders;
  return this.ParseOrders(orders , market , since , limit );
}

func (this *Mercado) FetchOpenOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchOpenOrders () requires a symbol argument") )));
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "coin_pair":*(market ).At(MkString("id") ),
        "status_list":MkString("[2]") ,
        }); _ = request;
  response := this.Call(MkString("privatePostListOrders"), this.Extend(request , params )); _ = response;
  responseData := this.SafeValue(response , MkString("response_data") , MkMap(&VarMap{})); _ = responseData;
  orders := this.SafeValue(responseData , MkString("orders") , MkArray(&VarArray{})); _ = orders;
  return this.ParseOrders(orders , market , since , limit );
}

func (this *Mercado) FetchMyTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchMyTrades () requires a symbol argument") )));
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "coin_pair":*(market ).At(MkString("id") ),
        "has_fills":MkBool(true) ,
        }); _ = request;
  response := this.Call(MkString("privatePostListOrders"), this.Extend(request , params )); _ = response;
  responseData := this.SafeValue(response , MkString("response_data") , MkMap(&VarMap{})); _ = responseData;
  ordersRaw := this.SafeValue(responseData , MkString("orders") , MkArray(&VarArray{})); _ = ordersRaw;
  orders := this.ParseOrders(ordersRaw , market , since , limit ); _ = orders;
  trades := this.OrdersToTrades(orders ); _ = trades;
  return this.FilterBySymbolSinceLimit(trades , symbol , since , limit );
}

func (this *Mercado) OrdersToTrades(goArgs ...*Variant) *Variant{
  orders := GoGetArg(goArgs, 0, MkUndefined()); _ = orders;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , orders.Length )); OpIncr(& i ){
    trades := this.SafeValue(*(orders ).At(i ), MkString("trades") , MkArray(&VarArray{})); _ = trades;
    for y := MkInteger(0) ; IsTrue(OpLw(y , trades.Length )); OpIncr(& y ){
      result.Push(*(trades ).At(y ));
    }
  }
  return result ;
}

func (this *Mercado) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("public") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  url := OpAdd(*(*(*this.At(MkString("urls")) ).At(MkString("api") )).At(api ), MkString("/") ); _ = url;
  query := this.Omit(params , this.ExtractParams(path )); _ = query;
  if IsTrue(OpOr(OpEq2(api , MkString("public") ), OpEq2(api , MkString("v4Public") ))) {
    OpAddAssign(& url , this.ImplodeParams(path , params ));
    if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
      OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(query )));
    }
  } else {
    this.CheckRequiredCredentials();
    OpAddAssign(& url , OpAdd(*this.At(MkString("version")) , MkString("/") ));
    nonce := this.Nonce(); _ = nonce;
    body = this.Urlencode(this.Extend(MkMap(&VarMap{
                "tapi_method":path ,
                "tapi_nonce":nonce ,
                }), params ));
    auth := OpAdd(MkString("/tapi/") , OpAdd(*this.At(MkString("version")) , OpAdd(MkString("/") , OpAdd(MkString("?") , body )))); _ = auth;
    headers = MkMap(&VarMap{
        "Content-Type":MkString("application/x-www-form-urlencoded") ,
        "TAPI-ID":*this.At(MkString("apiKey")) ,
        "TAPI-MAC":this.Hmac(this.Encode(auth ), this.Encode(*this.At(MkString("secret")) ), MkString("sha512") ),
        });
  }
  return MkMap(&VarMap{
        "url":url ,
        "method":method ,
        "body":body ,
        "headers":headers ,
        });
}

func (this *Mercado) HandleErrors(goArgs ...*Variant) *Variant{
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
  errorMessage := this.SafeValue(response , MkString("error_message") ); _ = errorMessage;
  if IsTrue(OpNotEq2(errorMessage , MkUndefined() )) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , this.Json(response )))));
  }
  return MkUndefined()
}

