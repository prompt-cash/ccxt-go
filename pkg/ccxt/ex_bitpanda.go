package ccxt

import (
)

type Bitpanda struct {
	*ExchangeBase
}

var _ Exchange = (*Bitpanda)(nil)

func init() {
	exchange := &Bitpanda{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Bitpanda) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("bitpanda") ,
            "name":MkString("Bitpanda Pro") ,
            "countries":MkArray(&VarArray{
                MkString("AT") ,
                }),
            "rateLimit":MkInteger(300) ,
            "version":MkString("v1") ,
            "has":MkMap(&VarMap{
                "CORS":MkBool(false) ,
                "publicAPI":MkBool(true) ,
                "privateAPI":MkBool(true) ,
                "cancelAllOrders":MkBool(true) ,
                "cancelOrder":MkBool(true) ,
                "cancelOrders":MkBool(true) ,
                "createDepositAddress":MkBool(true) ,
                "createOrder":MkBool(true) ,
                "fetchBalance":MkBool(true) ,
                "fetchClosedOrders":MkBool(true) ,
                "fetchCurrencies":MkBool(true) ,
                "fetchDeposits":MkBool(true) ,
                "fetchDepositAddress":MkBool(true) ,
                "fetchMarkets":MkBool(true) ,
                "fetchMyTrades":MkBool(true) ,
                "fetchOHLCV":MkBool(true) ,
                "fetchOpenOrders":MkBool(true) ,
                "fetchOrder":MkBool(true) ,
                "fetchOrderBook":MkBool(true) ,
                "fetchOrderTrades":MkBool(true) ,
                "fetchTime":MkBool(true) ,
                "fetchTrades":MkBool(true) ,
                "fetchTradingFees":MkBool(true) ,
                "fetchTicker":MkBool(true) ,
                "fetchTickers":MkBool(true) ,
                "fetchWithdrawals":MkBool(true) ,
                "withdraw":MkBool(true) ,
                }),
            "timeframes":MkMap(&VarMap{
                "1m":MkString("1/MINUTES") ,
                "5m":MkString("5/MINUTES") ,
                "15m":MkString("15/MINUTES") ,
                "30m":MkString("30/MINUTES") ,
                "1h":MkString("1/HOURS") ,
                "4h":MkString("4/HOURS") ,
                "1d":MkString("1/DAYS") ,
                "1w":MkString("1/WEEKS") ,
                "1M":MkString("1/MONTHS") ,
                }),
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/51840849/87591171-9a377d80-c6f0-11ea-94ac-97a126eac3bc.jpg") ,
                "api":MkMap(&VarMap{
                    "public":MkString("https://api.exchange.bitpanda.com/public") ,
                    "private":MkString("https://api.exchange.bitpanda.com/public") ,
                    }),
                "www":MkString("https://www.bitpanda.com/en/pro") ,
                "doc":MkArray(&VarArray{
                    MkString("https://developers.bitpanda.com/exchange/") ,
                    }),
                "fees":MkString("https://www.bitpanda.com/en/pro/fees") ,
                }),
            "api":MkMap(&VarMap{
                "public":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("currencies") ,
                        MkString("candlesticks/{instrument_code}") ,
                        MkString("fees") ,
                        MkString("instruments") ,
                        MkString("order-book/{instrument_code}") ,
                        MkString("market-ticker") ,
                        MkString("market-ticker/{instrument_code}") ,
                        MkString("price-ticks/{instrument_code}") ,
                        MkString("time") ,
                        })}),
                "private":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("account/balances") ,
                        MkString("account/deposit/crypto/{currency_code}") ,
                        MkString("account/deposit/fiat/EUR") ,
                        MkString("account/deposits") ,
                        MkString("account/deposits/bitpanda") ,
                        MkString("account/withdrawals") ,
                        MkString("account/withdrawals/bitpanda") ,
                        MkString("account/fees") ,
                        MkString("account/orders") ,
                        MkString("account/orders/{order_id}") ,
                        MkString("account/orders/{order_id}/trades") ,
                        MkString("account/trades") ,
                        MkString("account/trades/{trade_id}") ,
                        MkString("account/trading-volume") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("account/deposit/crypto") ,
                        MkString("account/withdraw/crypto") ,
                        MkString("account/withdraw/fiat") ,
                        MkString("account/fees") ,
                        MkString("account/orders") ,
                        }),
                    "delete":MkArray(&VarArray{
                        MkString("account/orders") ,
                        MkString("account/orders/{order_id}") ,
                        MkString("account/orders/client/{client_id}") ,
                        }),
                    }),
                }),
            "fees":MkMap(&VarMap{"trading":MkMap(&VarMap{
                    "tierBased":MkBool(true) ,
                    "percentage":MkBool(true) ,
                    "taker":this.ParseNumber(MkString("0.0015") ),
                    "maker":this.ParseNumber(MkString("0.001") ),
                    "tiers":MkArray(&VarArray{
                        MkMap(&VarMap{
                            "taker":MkArray(&VarArray{
                                MkArray(&VarArray{
                                    this.ParseNumber(MkString("0") ),
                                    this.ParseNumber(MkString("0.0015") ),
                                    }),
                                MkArray(&VarArray{
                                    this.ParseNumber(MkString("100") ),
                                    this.ParseNumber(MkString("0.0013") ),
                                    }),
                                MkArray(&VarArray{
                                    this.ParseNumber(MkString("250") ),
                                    this.ParseNumber(MkString("0.0013") ),
                                    }),
                                MkArray(&VarArray{
                                    this.ParseNumber(MkString("1000") ),
                                    this.ParseNumber(MkString("0.001") ),
                                    }),
                                MkArray(&VarArray{
                                    this.ParseNumber(MkString("5000") ),
                                    this.ParseNumber(MkString("0.0009") ),
                                    }),
                                MkArray(&VarArray{
                                    this.ParseNumber(MkString("10000") ),
                                    this.ParseNumber(MkString("0.00075") ),
                                    }),
                                MkArray(&VarArray{
                                    this.ParseNumber(MkString("20000") ),
                                    this.ParseNumber(MkString("0.00065") ),
                                    }),
                                }),
                            "maker":MkArray(&VarArray{
                                MkArray(&VarArray{
                                    this.ParseNumber(MkString("0") ),
                                    this.ParseNumber(MkString("0.001") ),
                                    }),
                                MkArray(&VarArray{
                                    this.ParseNumber(MkString("100") ),
                                    this.ParseNumber(MkString("0.001") ),
                                    }),
                                MkArray(&VarArray{
                                    this.ParseNumber(MkString("250") ),
                                    this.ParseNumber(MkString("0.0009") ),
                                    }),
                                MkArray(&VarArray{
                                    this.ParseNumber(MkString("1000") ),
                                    this.ParseNumber(MkString("0.00075") ),
                                    }),
                                MkArray(&VarArray{
                                    this.ParseNumber(MkString("5000") ),
                                    this.ParseNumber(MkString("0.0006") ),
                                    }),
                                MkArray(&VarArray{
                                    this.ParseNumber(MkString("10000") ),
                                    this.ParseNumber(MkString("0.0005") ),
                                    }),
                                MkArray(&VarArray{
                                    this.ParseNumber(MkString("20000") ),
                                    this.ParseNumber(MkString("0.0005") ),
                                    }),
                                }),
                            }),
                        }),
                    })}),
            "requiredCredentials":MkMap(&VarMap{
                "apiKey":MkBool(true) ,
                "secret":MkBool(false) ,
                }),
            "exceptions":MkMap(&VarMap{
                "exact":MkMap(&VarMap{
                    "INVALID_CLIENT_UUID":InvalidOrder ,
                    "ORDER_NOT_FOUND":OrderNotFound ,
                    "ONLY_ONE_ERC20_ADDRESS_ALLOWED":InvalidAddress ,
                    "DEPOSIT_ADDRESS_NOT_USED":InvalidAddress ,
                    "INVALID_CREDENTIALS":AuthenticationError ,
                    "MISSING_CREDENTIALS":AuthenticationError ,
                    "INVALID_APIKEY":AuthenticationError ,
                    "INVALID_SCOPES":AuthenticationError ,
                    "INVALID_SUBJECT":AuthenticationError ,
                    "INVALID_ISSUER":AuthenticationError ,
                    "INVALID_AUDIENCE":AuthenticationError ,
                    "INVALID_DEVICE_ID":AuthenticationError ,
                    "INVALID_IP_RESTRICTION":AuthenticationError ,
                    "APIKEY_REVOKED":AuthenticationError ,
                    "APIKEY_EXPIRED":AuthenticationError ,
                    "SYNCHRONIZER_TOKEN_MISMATCH":AuthenticationError ,
                    "SESSION_EXPIRED":AuthenticationError ,
                    "INTERNAL_ERROR":AuthenticationError ,
                    "CLIENT_IP_BLOCKED":PermissionDenied ,
                    "MISSING_PERMISSION":PermissionDenied ,
                    "ILLEGAL_CHARS":BadRequest ,
                    "UNSUPPORTED_MEDIA_TYPE":BadRequest ,
                    "ACCOUNT_HISTORY_TIME_RANGE_TOO_BIG":BadRequest ,
                    "CANDLESTICKS_TIME_RANGE_TOO_BIG":BadRequest ,
                    "INVALID_INSTRUMENT_CODE":BadRequest ,
                    "INVALID_ORDER_TYPE":BadRequest ,
                    "INVALID_UNIT":BadRequest ,
                    "INVALID_PERIOD":BadRequest ,
                    "INVALID_TIME":BadRequest ,
                    "INVALID_DATE":BadRequest ,
                    "INVALID_CURRENCY":BadRequest ,
                    "INVALID_AMOUNT":BadRequest ,
                    "INVALID_PRICE":BadRequest ,
                    "INVALID_LIMIT":BadRequest ,
                    "INVALID_QUERY":BadRequest ,
                    "INVALID_CURSOR":BadRequest ,
                    "INVALID_ACCOUNT_ID":BadRequest ,
                    "INVALID_SIDE":InvalidOrder ,
                    "INVALID_ACCOUNT_HISTORY_FROM_TIME":BadRequest ,
                    "INVALID_ACCOUNT_HISTORY_MAX_PAGE_SIZE":BadRequest ,
                    "INVALID_ACCOUNT_HISTORY_TIME_PERIOD":BadRequest ,
                    "INVALID_ACCOUNT_HISTORY_TO_TIME":BadRequest ,
                    "INVALID_CANDLESTICKS_GRANULARITY":BadRequest ,
                    "INVALID_CANDLESTICKS_UNIT":BadRequest ,
                    "INVALID_ORDER_BOOK_DEPTH":BadRequest ,
                    "INVALID_ORDER_BOOK_LEVEL":BadRequest ,
                    "INVALID_PAGE_CURSOR":BadRequest ,
                    "INVALID_TIME_RANGE":BadRequest ,
                    "INVALID_TRADE_ID":BadRequest ,
                    "INVALID_UI_ACCOUNT_SETTINGS":BadRequest ,
                    "NEGATIVE_AMOUNT":InvalidOrder ,
                    "NEGATIVE_PRICE":InvalidOrder ,
                    "MIN_SIZE_NOT_SATISFIED":InvalidOrder ,
                    "BAD_AMOUNT_PRECISION":InvalidOrder ,
                    "BAD_PRICE_PRECISION":InvalidOrder ,
                    "BAD_TRIGGER_PRICE_PRECISION":InvalidOrder ,
                    "MAX_OPEN_ORDERS_EXCEEDED":BadRequest ,
                    "MISSING_PRICE":InvalidOrder ,
                    "MISSING_ORDER_TYPE":InvalidOrder ,
                    "MISSING_SIDE":InvalidOrder ,
                    "MISSING_CANDLESTICKS_PERIOD_PARAM":ArgumentsRequired ,
                    "MISSING_CANDLESTICKS_UNIT_PARAM":ArgumentsRequired ,
                    "MISSING_FROM_PARAM":ArgumentsRequired ,
                    "MISSING_INSTRUMENT_CODE":ArgumentsRequired ,
                    "MISSING_ORDER_ID":InvalidOrder ,
                    "MISSING_TO_PARAM":ArgumentsRequired ,
                    "MISSING_TRADE_ID":ArgumentsRequired ,
                    "INVALID_ORDER_ID":OrderNotFound ,
                    "NOT_FOUND":OrderNotFound ,
                    "INSUFFICIENT_LIQUIDITY":InsufficientFunds ,
                    "INSUFFICIENT_FUNDS":InsufficientFunds ,
                    "NO_TRADING":ExchangeNotAvailable ,
                    "SERVICE_UNAVAILABLE":ExchangeNotAvailable ,
                    "GATEWAY_TIMEOUT":ExchangeNotAvailable ,
                    "RATELIMIT":DDoSProtection ,
                    "CF_RATELIMIT":DDoSProtection ,
                    "INTERNAL_SERVER_ERROR":ExchangeError ,
                    }),
                "broad":MkMap(&VarMap{}),
                }),
            "commonCurrencies":MkMap(&VarMap{"MIOTA":MkString("IOTA") }),
            "options":MkMap(&VarMap{
                "fetchTradingFees":MkMap(&VarMap{"method":MkString("fetchPrivateTradingFees") }),
                "fiat":MkArray(&VarArray{
                    MkString("EUR") ,
                    MkString("CHF") ,
                    }),
                }),
            }));
}

func (this *Bitpanda) FetchTime(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetTime"), params ); _ = response;
  return this.SafeInteger(response , MkString("epoch_millis") );
}

func (this *Bitpanda) FetchCurrencies(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetCurrencies"), params ); _ = response;
  result := MkMap(&VarMap{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , response.Length )); OpIncr(& i ){
    currency := *(response ).At(i ); _ = currency;
    id := this.SafeString(currency , MkString("code") ); _ = id;
    code := this.SafeCurrencyCode(id ); _ = code;
    *(result ).At(code )= MkMap(&VarMap{
        "id":id ,
        "code":code ,
        "name":MkUndefined() ,
        "info":currency ,
        "active":MkUndefined() ,
        "fee":MkUndefined() ,
        "precision":this.SafeInteger(currency , MkString("precision") ),
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
  return result ;
}

func (this *Bitpanda) FetchMarkets(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetInstruments"), params ); _ = response;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , response.Length )); OpIncr(& i ){
    market := *(response ).At(i ); _ = market;
    baseAsset := this.SafeValue(market , MkString("base") , MkMap(&VarMap{})); _ = baseAsset;
    quoteAsset := this.SafeValue(market , MkString("quote") , MkMap(&VarMap{})); _ = quoteAsset;
    baseId := this.SafeString(baseAsset , MkString("code") ); _ = baseId;
    quoteId := this.SafeString(quoteAsset , MkString("code") ); _ = quoteId;
    id := OpAdd(baseId , OpAdd(MkString("_") , quoteId )); _ = id;
    base := this.SafeCurrencyCode(baseId ); _ = base;
    quote := this.SafeCurrencyCode(quoteId ); _ = quote;
    symbol := OpAdd(base , OpAdd(MkString("/") , quote )); _ = symbol;
    precision := MkMap(&VarMap{
          "amount":this.SafeInteger(market , MkString("amount_precision") ),
          "price":this.SafeInteger(market , MkString("market_precision") ),
          }); _ = precision;
    limits := MkMap(&VarMap{
          "amount":MkMap(&VarMap{
              "min":MkUndefined() ,
              "max":MkUndefined() ,
              }),
          "price":MkMap(&VarMap{
              "min":MkUndefined() ,
              "max":MkUndefined() ,
              }),
          "cost":MkMap(&VarMap{
              "min":this.SafeNumber(market , MkString("min_size") ),
              "max":MkUndefined() ,
              }),
          }); _ = limits;
    state := this.SafeString(market , MkString("state") ); _ = state;
    active := OpEq2(state , MkString("ACTIVE") ); _ = active;
    result.Push(MkMap(&VarMap{
            "id":id ,
            "symbol":symbol ,
            "base":base ,
            "quote":quote ,
            "baseId":baseId ,
            "quoteId":quoteId ,
            "precision":precision ,
            "limits":limits ,
            "info":market ,
            "active":active ,
            }));
  }
  return result ;
}

func (this *Bitpanda) FetchTradingFees(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  method := this.SafeString(params , MkString("method") ); _ = method;
  params = this.Omit(params , MkString("method") );
  if IsTrue(OpEq2(method , MkUndefined() )) {
    options := this.SafeValue(*this.At(MkString("options")) , MkString("fetchTradingFees") , MkMap(&VarMap{})); _ = options;
    method = this.SafeString(options , MkString("method") , MkString("fetchPrivateTradingFees") );
  }
  return this.Call(method , params );
}

func (this *Bitpanda) FetchPublicTradingFees(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("publicGetFees"), params ); _ = response;
  feeGroupsById := this.IndexBy(response , MkString("fee_group_id") ); _ = feeGroupsById;
  feeGroupId := this.SafeValue(*this.At(MkString("options")) , MkString("fee_group_id") , MkString("default") ); _ = feeGroupId;
  feeGroup := this.SafeValue(feeGroupsById , feeGroupId , MkMap(&VarMap{})); _ = feeGroup;
  feeTiers := this.SafeValue(feeGroup , MkString("fee_tiers") ); _ = feeTiers;
  result := MkMap(&VarMap{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , (*((this).At(MkString("symbols")))).Length )); OpIncr(& i ){
    symbol := *(*this.At(MkString("symbols")) ).At(i ); _ = symbol;
    fee := MkMap(&VarMap{
          "info":feeGroup ,
          "symbol":symbol ,
          "maker":MkUndefined() ,
          "taker":MkUndefined() ,
          "percentage":MkBool(true) ,
          "tierBased":MkBool(true) ,
          }); _ = fee;
    takerFees := MkArray(&VarArray{}); _ = takerFees;
    makerFees := MkArray(&VarArray{}); _ = makerFees;
    for i := MkInteger(0) ; IsTrue(OpLw(i , feeTiers.Length )); OpIncr(& i ){
      tier := *(feeTiers ).At(i ); _ = tier;
      volume := this.SafeNumber(tier , MkString("volume") ); _ = volume;
      taker := this.SafeNumber(tier , MkString("taker_fee") ); _ = taker;
      maker := this.SafeNumber(tier , MkString("maker_fee") ); _ = maker;
      OpDivAssign(& taker , MkInteger(100) );
      OpDivAssign(& maker , MkInteger(100) );
      takerFees.Push(MkArray(&VarArray{
              volume ,
              taker ,
              }));
      makerFees.Push(MkArray(&VarArray{
              volume ,
              maker ,
              }));
      if IsTrue(OpEq2(i , MkInteger(0) )) {
        *(fee ).At(MkString("taker") )= OpCopy(taker );
        *(fee ).At(MkString("maker") )= OpCopy(maker );
      }
    }
    tiers := MkMap(&VarMap{
          "taker":takerFees ,
          "maker":makerFees ,
          }); _ = tiers;
    *(fee ).At(MkString("tiers") )= OpCopy(tiers );
    *(result ).At(symbol )= OpCopy(fee );
  }
  return result ;
}

func (this *Bitpanda) FetchPrivateTradingFees(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("privateGetAccountFees"), params ); _ = response;
  activeFeeTier := this.SafeValue(response , MkString("active_fee_tier") , MkMap(&VarMap{})); _ = activeFeeTier;
  result := MkMap(&VarMap{
        "info":response ,
        "maker":this.SafeNumber(activeFeeTier , MkString("maker_fee") ),
        "taker":this.SafeNumber(activeFeeTier , MkString("taker_fee") ),
        "percentage":MkBool(true) ,
        "tierBased":MkBool(true) ,
        }); _ = result;
  feeTiers := this.SafeValue(response , MkString("fee_tiers") ); _ = feeTiers;
  takerFees := MkArray(&VarArray{}); _ = takerFees;
  makerFees := MkArray(&VarArray{}); _ = makerFees;
  for i := MkInteger(0) ; IsTrue(OpLw(i , feeTiers.Length )); OpIncr(& i ){
    tier := *(feeTiers ).At(i ); _ = tier;
    volume := this.SafeNumber(tier , MkString("volume") ); _ = volume;
    taker := this.SafeNumber(tier , MkString("taker_fee") ); _ = taker;
    maker := this.SafeNumber(tier , MkString("maker_fee") ); _ = maker;
    OpDivAssign(& taker , MkInteger(100) );
    OpDivAssign(& maker , MkInteger(100) );
    takerFees.Push(MkArray(&VarArray{
            volume ,
            taker ,
            }));
    makerFees.Push(MkArray(&VarArray{
            volume ,
            maker ,
            }));
  }
  tiers := MkMap(&VarMap{
        "taker":takerFees ,
        "maker":makerFees ,
        }); _ = tiers;
  *(result ).At(MkString("tiers") )= OpCopy(tiers );
  return result ;
}

func (this *Bitpanda) ParseTicker(goArgs ...*Variant) *Variant{
  ticker := GoGetArg(goArgs, 0, MkUndefined()); _ = ticker;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.Parse8601(this.SafeString(ticker , MkString("time") )); _ = timestamp;
  marketId := this.SafeString(ticker , MkString("instrument_code") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market , MkString("_") ); _ = symbol;
  last := this.SafeNumber(ticker , MkString("last_price") ); _ = last;
  percentage := this.SafeNumber(ticker , MkString("price_change_percentage") ); _ = percentage;
  change := this.SafeNumber(ticker , MkString("price_change") ); _ = change;
  open := OpCopy(MkUndefined() ); _ = open;
  average := OpCopy(MkUndefined() ); _ = average;
  if IsTrue(OpAnd(OpNotEq2(last , MkUndefined() ), OpNotEq2(change , MkUndefined() ))) {
    open = OpSub(last , change );
    average = OpDiv(this.Sum(last , open ), MkInteger(2) );
  }
  baseVolume := this.SafeNumber(ticker , MkString("base_volume") ); _ = baseVolume;
  quoteVolume := this.SafeNumber(ticker , MkString("quote_volume") ); _ = quoteVolume;
  vwap := this.Vwap(baseVolume , quoteVolume ); _ = vwap;
  return MkMap(&VarMap{
        "symbol":symbol ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "high":this.SafeNumber(ticker , MkString("high") ),
        "low":this.SafeNumber(ticker , MkString("low") ),
        "bid":this.SafeNumber(ticker , MkString("best_bid") ),
        "bidVolume":MkUndefined() ,
        "ask":this.SafeNumber(ticker , MkString("best_ask") ),
        "askVolume":MkUndefined() ,
        "vwap":vwap ,
        "open":open ,
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

func (this *Bitpanda) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"instrument_code":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("publicGetMarketTickerInstrumentCode"), this.Extend(request , params )); _ = response;
  return this.ParseTicker(response , market );
}

func (this *Bitpanda) FetchTickers(goArgs ...*Variant) *Variant{
  symbols := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbols;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("publicGetMarketTicker"), params ); _ = response;
  result := MkMap(&VarMap{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , response.Length )); OpIncr(& i ){
    ticker := this.ParseTicker(*(response ).At(i )); _ = ticker;
    symbol := *(ticker ).At(MkString("symbol") ); _ = symbol;
    *(result ).At(symbol )= OpCopy(ticker );
  }
  return this.FilterByArray(result , MkString("symbol") , symbols );
}

func (this *Bitpanda) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"instrument_code":this.MarketId(symbol )}); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("depth") )= OpCopy(limit );
  }
  response := this.Call(MkString("publicGetOrderBookInstrumentCode"), this.Extend(request , params )); _ = response;
  timestamp := this.Parse8601(this.SafeString(response , MkString("time") )); _ = timestamp;
  return this.ParseOrderBook(response , symbol , timestamp , MkString("bids") , MkString("asks") , MkString("price") , MkString("amount") );
}

func (this *Bitpanda) ParseOHLCV(goArgs ...*Variant) *Variant{
  ohlcv := GoGetArg(goArgs, 0, MkUndefined()); _ = ohlcv;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  granularity := this.SafeValue(ohlcv , MkString("granularity") ); _ = granularity;
  unit := this.SafeString(granularity , MkString("unit") ); _ = unit;
  period := this.SafeString(granularity , MkString("period") ); _ = period;
  units := MkMap(&VarMap{
        "MINUTES":MkString("m") ,
        "HOURS":MkString("h") ,
        "DAYS":MkString("d") ,
        "WEEKS":MkString("w") ,
        "MONTHS":MkString("M") ,
        }); _ = units;
  lowercaseUnit := this.SafeString(units , unit ); _ = lowercaseUnit;
  timeframe := OpAdd(period , lowercaseUnit ); _ = timeframe;
  durationInSeconds := this.ParseTimeframe(timeframe ); _ = durationInSeconds;
  duration := OpMulti(durationInSeconds , MkInteger(1000) ); _ = duration;
  timestamp := this.Parse8601(this.SafeString(ohlcv , MkString("time") )); _ = timestamp;
  alignedTimestamp := OpMulti(duration , parseInt(OpDiv(timestamp , duration ))); _ = alignedTimestamp;
  options := this.SafeValue(*this.At(MkString("options")) , MkString("fetchOHLCV") , MkMap(&VarMap{})); _ = options;
  volumeField := this.SafeString(options , MkString("volume") , MkString("total_amount") ); _ = volumeField;
  return MkArray(&VarArray{
        alignedTimestamp ,
        this.SafeNumber(ohlcv , MkString("open") ),
        this.SafeNumber(ohlcv , MkString("high") ),
        this.SafeNumber(ohlcv , MkString("low") ),
        this.SafeNumber(ohlcv , MkString("close") ),
        this.SafeNumber(ohlcv , volumeField ),
        });
}

func (this *Bitpanda) FetchOHLCV(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  timeframe := GoGetArg(goArgs, 1, MkString("1m") ); _ = timeframe;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  periodUnit := this.SafeString(*this.At(MkString("timeframes")) , timeframe ); _ = periodUnit;
  period := MkUndefined(); _ = period;
  unit := MkUndefined(); _ = unit;
  ArrayUnpack(periodUnit.Split(MkString("/") ), &period, &unit);
  durationInSeconds := this.ParseTimeframe(timeframe ); _ = durationInSeconds;
  duration := OpMulti(durationInSeconds , MkInteger(1000) ); _ = duration;
  if IsTrue(OpEq2(limit , MkUndefined() )) {
    limit = MkInteger(1500) ;
  }
  request := MkMap(&VarMap{
        "instrument_code":*(market ).At(MkString("id") ),
        "period":period ,
        "unit":unit ,
        }); _ = request;
  if IsTrue(OpEq2(since , MkUndefined() )) {
    now := this.Milliseconds(); _ = now;
    *(request ).At(MkString("to") )= this.Iso8601(now );
    *(request ).At(MkString("from") )= this.Iso8601(OpSub(now , OpMulti(limit , duration )));
  } else {
    *(request ).At(MkString("from") )= this.Iso8601(since );
    *(request ).At(MkString("to") )= this.Iso8601(this.Sum(since , OpMulti(limit , duration )));
  }
  response := this.Call(MkString("publicGetCandlesticksInstrumentCode"), this.Extend(request , params )); _ = response;
  return this.ParseOHLCVs(response , market , timeframe , since , limit );
}

func (this *Bitpanda) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  feeInfo := this.SafeValue(trade , MkString("fee") , MkMap(&VarMap{})); _ = feeInfo;
  trade = this.SafeValue(trade , MkString("trade") , trade );
  timestamp := this.SafeInteger(trade , MkString("trade_timestamp") ); _ = timestamp;
  if IsTrue(OpEq2(timestamp , MkUndefined() )) {
    timestamp = this.Parse8601(this.SafeString(trade , MkString("time") ));
  }
  side := this.SafeStringLower2(trade , MkString("side") , MkString("taker_side") ); _ = side;
  price := this.SafeNumber(trade , MkString("price") ); _ = price;
  amount := this.SafeNumber(trade , MkString("amount") ); _ = amount;
  cost := this.SafeNumber(trade , MkString("volume") ); _ = cost;
  if IsTrue(OpAnd(OpEq2(cost , MkUndefined() ), OpAnd(OpNotEq2(amount , MkUndefined() ), OpNotEq2(price , MkUndefined() )))) {
    cost = OpMulti(amount , price );
  }
  marketId := this.SafeString(trade , MkString("instrument_code") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market , MkString("_") ); _ = symbol;
  feeCost := this.SafeNumber(feeInfo , MkString("fee_amount") ); _ = feeCost;
  takerOrMaker := OpCopy(MkUndefined() ); _ = takerOrMaker;
  fee := OpCopy(MkUndefined() ); _ = fee;
  if IsTrue(OpNotEq2(feeCost , MkUndefined() )) {
    feeCurrencyId := this.SafeString(feeInfo , MkString("fee_currency") ); _ = feeCurrencyId;
    feeCurrencyCode := this.SafeCurrencyCode(feeCurrencyId ); _ = feeCurrencyCode;
    feeRate := this.SafeNumber(feeInfo , MkString("fee_percentage") ); _ = feeRate;
    fee = MkMap(&VarMap{
        "cost":feeCost ,
        "currency":feeCurrencyCode ,
        "rate":feeRate ,
        });
    takerOrMaker = this.SafeStringLower(feeInfo , MkString("fee_type") );
  }
  return MkMap(&VarMap{
        "id":this.SafeString2(trade , MkString("trade_id") , MkString("sequence") ),
        "order":this.SafeString(trade , MkString("order_id") ),
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "symbol":symbol ,
        "type":MkUndefined() ,
        "side":side ,
        "price":price ,
        "amount":amount ,
        "cost":cost ,
        "takerOrMaker":takerOrMaker ,
        "fee":fee ,
        "info":trade ,
        });
}

func (this *Bitpanda) FetchTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"instrument_code":*(market ).At(MkString("id") )}); _ = request;
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("from") )= this.Iso8601(since );
    *(request ).At(MkString("to") )= this.Iso8601(this.Sum(since , MkInteger(14400000) ));
  }
  response := this.Call(MkString("publicGetPriceTicksInstrumentCode"), this.Extend(request , params )); _ = response;
  return this.ParseTrades(response , market , since , limit );
}

func (this *Bitpanda) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("privateGetAccountBalances"), params ); _ = response;
  balances := this.SafeValue(response , MkString("balances") , MkArray(&VarArray{})); _ = balances;
  result := MkMap(&VarMap{"info":response }); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , balances.Length )); OpIncr(& i ){
    balance := *(balances ).At(i ); _ = balance;
    currencyId := this.SafeString(balance , MkString("currency_code") ); _ = currencyId;
    code := this.SafeCurrencyCode(currencyId ); _ = code;
    account := this.Account(); _ = account;
    *(account ).At(MkString("free") )= this.SafeString(balance , MkString("available") );
    *(account ).At(MkString("used") )= this.SafeString(balance , MkString("locked") );
    *(result ).At(code )= OpCopy(account );
  }
  return this.ParseBalance(result );
}

func (this *Bitpanda) ParseDepositAddress(goArgs ...*Variant) *Variant{
  depositAddress := GoGetArg(goArgs, 0, MkUndefined()); _ = depositAddress;
  currency := GoGetArg(goArgs, 1, MkUndefined() ); _ = currency;
  code := OpCopy(MkUndefined() ); _ = code;
  if IsTrue(OpNotEq2(currency , MkUndefined() )) {
    code = *(currency ).At(MkString("code") );
  }
  address := this.SafeString(depositAddress , MkString("address") ); _ = address;
  tag := this.SafeString(depositAddress , MkString("destination_tag") ); _ = tag;
  this.CheckAddress(address );
  return MkMap(&VarMap{
        "currency":code ,
        "address":address ,
        "tag":tag ,
        "info":depositAddress ,
        });
}

func (this *Bitpanda) CreateDepositAddress(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  currency := this.Currency(code ); _ = currency;
  request := MkMap(&VarMap{"currency":*(currency ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("privatePostAccountDepositCrypto"), this.Extend(request , params )); _ = response;
  return this.ParseDepositAddress(response , currency );
}

func (this *Bitpanda) FetchDepositAddress(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  currency := this.Currency(code ); _ = currency;
  request := MkMap(&VarMap{"currency_code":*(currency ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("privateGetAccountDepositCryptoCurrencyCode"), this.Extend(request , params )); _ = response;
  return this.ParseDepositAddress(response , currency );
}

func (this *Bitpanda) FetchDeposits(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  currency := OpCopy(MkUndefined() ); _ = currency;
  if IsTrue(OpNotEq2(code , MkUndefined() )) {
    currency = this.Currency(code );
    *(request ).At(MkString("currency_code") )= *(currency ).At(MkString("id") );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("max_page_size") )= OpCopy(limit );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    to := this.SafeString(params , MkString("to") ); _ = to;
    if IsTrue(OpEq2(to , MkUndefined() )) {
      panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchDeposits() requires a \"to\" iso8601 string param with the since argument is specified") )));
    }
    *(request ).At(MkString("from") )= this.Iso8601(since );
  }
  response := this.Call(MkString("privateGetAccountDeposits"), this.Extend(request , params )); _ = response;
  depositHistory := this.SafeValue(response , MkString("deposit_history") , MkArray(&VarArray{})); _ = depositHistory;
  return this.ParseTransactions(depositHistory , currency , since , limit , MkMap(&VarMap{"type":MkString("deposit") }));
}

func (this *Bitpanda) FetchWithdrawals(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  currency := OpCopy(MkUndefined() ); _ = currency;
  if IsTrue(OpNotEq2(code , MkUndefined() )) {
    currency = this.Currency(code );
    *(request ).At(MkString("currency_code") )= *(currency ).At(MkString("id") );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("max_page_size") )= OpCopy(limit );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    to := this.SafeString(params , MkString("to") ); _ = to;
    if IsTrue(OpEq2(to , MkUndefined() )) {
      panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchWithdrawals() requires a \"to\" iso8601 string param with the since argument is specified") )));
    }
    *(request ).At(MkString("from") )= this.Iso8601(since );
  }
  response := this.Call(MkString("privateGetAccountWithdrawals"), this.Extend(request , params )); _ = response;
  withdrawalHistory := this.SafeValue(response , MkString("withdrawal_history") , MkArray(&VarArray{})); _ = withdrawalHistory;
  return this.ParseTransactions(withdrawalHistory , currency , since , limit , MkMap(&VarMap{"type":MkString("withdrawal") }));
}

func (this *Bitpanda) Withdraw(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  amount := GoGetArg(goArgs, 1, MkUndefined()); _ = amount;
  address := GoGetArg(goArgs, 2, MkUndefined()); _ = address;
  tag := GoGetArg(goArgs, 3, MkUndefined() ); _ = tag;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.CheckAddress(address );
  this.LoadMarkets();
  currency := this.Currency(code ); _ = currency;
  request := MkMap(&VarMap{
        "currency":code ,
        "amount":this.CurrencyToPrecision(code , amount ),
        }); _ = request;
  options := this.SafeValue(*this.At(MkString("options")) , MkString("fiat") , MkArray(&VarArray{})); _ = options;
  isFiat := this.InArray(code , options ); _ = isFiat;
  method := OpTrinary(isFiat , MkString("privatePostAccountWithdrawFiat") , MkString("privatePostAccountWithdrawCrypto") ); _ = method;
  if IsTrue(isFiat ) {
    payoutAccountId := this.SafeString(params , MkString("payout_account_id") ); _ = payoutAccountId;
    if IsTrue(OpEq2(payoutAccountId , MkUndefined() )) {
      panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" withdraw() requires a payout_account_id param for fiat ") , OpAdd(code , MkString(" withdrawals") )))));
    }
  } else {
    recipient := MkMap(&VarMap{"address":address }); _ = recipient;
    if IsTrue(OpNotEq2(tag , MkUndefined() )) {
      *(recipient ).At(MkString("destination_tag") )= OpCopy(tag );
    }
    *(request ).At(MkString("recipient") )= OpCopy(recipient );
  }
  response := this.Call(method , this.Extend(request , params )); _ = response;
  return this.ParseTransaction(response , currency );
}

func (this *Bitpanda) ParseTransaction(goArgs ...*Variant) *Variant{
  transaction := GoGetArg(goArgs, 0, MkUndefined()); _ = transaction;
  currency := GoGetArg(goArgs, 1, MkUndefined() ); _ = currency;
  id := this.SafeString(transaction , MkString("transaction_id") ); _ = id;
  amount := this.SafeNumber(transaction , MkString("amount") ); _ = amount;
  timestamp := this.Parse8601(this.SafeString(transaction , MkString("time") )); _ = timestamp;
  currencyId := this.SafeString(transaction , MkString("currency") ); _ = currencyId;
  currency = this.SafeCurrency(currencyId , currency );
  status := MkString("ok") ; _ = status;
  feeCost := this.SafeNumber2(transaction , MkString("fee_amount") , MkString("fee") ); _ = feeCost;
  fee := OpCopy(MkUndefined() ); _ = fee;
  addressTo := this.SafeString(transaction , MkString("recipient") ); _ = addressTo;
  tagTo := this.SafeString(transaction , MkString("destination_tag") ); _ = tagTo;
  if IsTrue(OpNotEq2(feeCost , MkUndefined() )) {
    feeCurrencyId := this.SafeString(transaction , MkString("fee_currency") , currencyId ); _ = feeCurrencyId;
    feeCurrencyCode := this.SafeCurrencyCode(feeCurrencyId ); _ = feeCurrencyCode;
    fee = MkMap(&VarMap{
        "cost":feeCost ,
        "currency":feeCurrencyCode ,
        });
  }
  return MkMap(&VarMap{
        "info":transaction ,
        "id":id ,
        "currency":*(currency ).At(MkString("code") ),
        "amount":amount ,
        "address":addressTo ,
        "addressFrom":MkUndefined() ,
        "addressTo":addressTo ,
        "tag":tagTo ,
        "tagFrom":MkUndefined() ,
        "tagTo":tagTo ,
        "status":status ,
        "type":MkUndefined() ,
        "updated":MkUndefined() ,
        "txid":this.SafeString(transaction , MkString("blockchain_transaction_id") ),
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "fee":fee ,
        });
}

func (this *Bitpanda) ParseOrderStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "FILLED":MkString("open") ,
        "FILLED_FULLY":MkString("closed") ,
        "FILLED_CLOSED":MkString("canceled") ,
        "FILLED_REJECTED":MkString("rejected") ,
        "OPEN":MkString("open") ,
        "REJECTED":MkString("rejected") ,
        "CLOSED":MkString("canceled") ,
        "FAILED":MkString("failed") ,
        "STOP_TRIGGERED":MkString("triggered") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Bitpanda) ParseOrder(goArgs ...*Variant) *Variant{
  order := GoGetArg(goArgs, 0, MkUndefined()); _ = order;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  rawOrder := this.SafeValue(order , MkString("order") , order ); _ = rawOrder;
  id := this.SafeString(rawOrder , MkString("order_id") ); _ = id;
  clientOrderId := this.SafeString(rawOrder , MkString("client_id") ); _ = clientOrderId;
  timestamp := this.Parse8601(this.SafeString(rawOrder , MkString("time") )); _ = timestamp;
  rawStatus := this.ParseOrderStatus(this.SafeString(rawOrder , MkString("status") )); _ = rawStatus;
  status := this.ParseOrderStatus(rawStatus ); _ = status;
  marketId := this.SafeString(rawOrder , MkString("instrument_code") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market , MkString("_") ); _ = symbol;
  price := this.SafeNumber(rawOrder , MkString("price") ); _ = price;
  amount := this.SafeNumber(rawOrder , MkString("amount") ); _ = amount;
  filledString := this.SafeString(rawOrder , MkString("filled_amount") ); _ = filledString;
  filled := this.ParseNumber(filledString ); _ = filled;
  side := this.SafeStringLower(rawOrder , MkString("side") ); _ = side;
  type_ := this.SafeStringLower(rawOrder , MkString("type") ); _ = type_;
  timeInForce := this.ParseTimeInForce(this.SafeString(rawOrder , MkString("time_in_force") )); _ = timeInForce;
  stopPrice := this.SafeNumber(rawOrder , MkString("trigger_price") ); _ = stopPrice;
  postOnly := this.SafeValue(rawOrder , MkString("is_post_only") ); _ = postOnly;
  rawTrades := this.SafeValue(order , MkString("trades") , MkArray(&VarArray{})); _ = rawTrades;
  trades := this.ParseTrades(rawTrades , market , MkUndefined() , MkUndefined() , MkMap(&VarMap{"type":type_ })); _ = trades;
  return this.SafeOrder(MkMap(&VarMap{
            "id":id ,
            "clientOrderId":clientOrderId ,
            "info":order ,
            "timestamp":timestamp ,
            "datetime":this.Iso8601(timestamp ),
            "lastTradeTimestamp":MkUndefined() ,
            "symbol":symbol ,
            "type":type_ ,
            "timeInForce":timeInForce ,
            "postOnly":postOnly ,
            "side":side ,
            "price":price ,
            "stopPrice":stopPrice ,
            "amount":amount ,
            "cost":MkUndefined() ,
            "average":MkUndefined() ,
            "filled":filled ,
            "remaining":MkUndefined() ,
            "status":status ,
            "trades":trades ,
            }));
}

func (this *Bitpanda) ParseTimeInForce(goArgs ...*Variant) *Variant{
  timeInForce := GoGetArg(goArgs, 0, MkUndefined()); _ = timeInForce;
  timeInForces := MkMap(&VarMap{
        "GOOD_TILL_CANCELLED":MkString("GTC") ,
        "GOOD_TILL_TIME":MkString("GTT") ,
        "IMMEDIATE_OR_CANCELLED":MkString("IOC") ,
        "FILL_OR_KILL":MkString("FOK") ,
        }); _ = timeInForces;
  return this.SafeString(timeInForces , timeInForce , timeInForce );
}

func (this *Bitpanda) CreateOrder(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 2, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 3, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 4, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 5, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  uppercaseType := type_.ToUpperCase(); _ = uppercaseType;
  request := MkMap(&VarMap{
        "instrument_code":*(market ).At(MkString("id") ),
        "type":uppercaseType ,
        "side":side.ToUpperCase(),
        "amount":this.AmountToPrecision(symbol , amount ),
        }); _ = request;
  priceIsRequired := OpCopy(MkBool(false) ); _ = priceIsRequired;
  if IsTrue(OpOr(OpEq2(uppercaseType , MkString("LIMIT") ), OpEq2(uppercaseType , MkString("STOP") ))) {
    priceIsRequired = OpCopy(MkBool(true) );
  }
  if IsTrue(OpEq2(uppercaseType , MkString("STOP") )) {
    triggerPrice := this.SafeNumber(params , MkString("trigger_price") ); _ = triggerPrice;
    if IsTrue(OpEq2(triggerPrice , MkUndefined() )) {
      panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" createOrder() requires a trigger_price param for ") , OpAdd(type_ , MkString(" orders") )))));
    }
    *(request ).At(MkString("trigger_price") )= this.PriceToPrecision(symbol , triggerPrice );
    params = this.Omit(params , MkString("trigger_price") );
  }
  if IsTrue(priceIsRequired ) {
    *(request ).At(MkString("price") )= this.PriceToPrecision(symbol , price );
  }
  clientOrderId := this.SafeString2(params , MkString("clientOrderId") , MkString("client_id") ); _ = clientOrderId;
  if IsTrue(OpNotEq2(clientOrderId , MkUndefined() )) {
    *(request ).At(MkString("client_id") )= OpCopy(clientOrderId );
    params = this.Omit(params , MkArray(&VarArray{
            MkString("clientOrderId") ,
            MkString("client_id") ,
            }));
  }
  response := this.Call(MkString("privatePostAccountOrders"), this.Extend(request , params )); _ = response;
  return this.ParseOrder(response , market );
}

func (this *Bitpanda) CancelOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  clientOrderId := this.SafeString2(params , MkString("clientOrderId") , MkString("client_id") ); _ = clientOrderId;
  params = this.Omit(params , MkArray(&VarArray{
          MkString("clientOrderId") ,
          MkString("client_id") ,
          }));
  method := MkString("privateDeleteAccountOrdersOrderId") ; _ = method;
  request := MkMap(&VarMap{}); _ = request;
  if IsTrue(OpNotEq2(clientOrderId , MkUndefined() )) {
    method = MkString("privateDeleteAccountOrdersClientClientId") ;
    *(request ).At(MkString("client_id") )= OpCopy(clientOrderId );
  } else {
    *(request ).At(MkString("order_id") )= OpCopy(id );
  }
  response := this.Call(method , this.Extend(request , params )); _ = response;
  return response ;
}

func (this *Bitpanda) CancelAllOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market := this.Market(symbol ); _ = market;
    *(request ).At(MkString("instrument_code") )= *(market ).At(MkString("id") );
  }
  response := this.Call(MkString("privateDeleteAccountOrders"), this.Extend(request , params )); _ = response;
  return response ;
}

func (this *Bitpanda) CancelOrders(goArgs ...*Variant) *Variant{
  ids := GoGetArg(goArgs, 0, MkUndefined()); _ = ids;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"ids":ids.Join(MkString(",") )}); _ = request;
  response := this.Call(MkString("privateDeleteAccountOrders"), this.Extend(request , params )); _ = response;
  return response ;
}

func (this *Bitpanda) FetchOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"order_id":id }); _ = request;
  response := this.Call(MkString("privateGetAccountOrdersOrderId"), this.Extend(request , params )); _ = response;
  return this.ParseOrder(response );
}

func (this *Bitpanda) FetchOpenOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("instrument_code") )= *(market ).At(MkString("id") );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    to := this.SafeString(params , MkString("to") ); _ = to;
    if IsTrue(OpEq2(to , MkUndefined() )) {
      panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchOrders() requires a \"to\" iso8601 string param with the since argument is specified, max range is 100 days") )));
    }
    *(request ).At(MkString("from") )= this.Iso8601(since );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("max_page_size") )= OpCopy(limit );
  }
  response := this.Call(MkString("privateGetAccountOrders"), this.Extend(request , params )); _ = response;
  orderHistory := this.SafeValue(response , MkString("order_history") , MkArray(&VarArray{})); _ = orderHistory;
  return this.ParseOrders(orderHistory , market , since , limit );
}

func (this *Bitpanda) FetchClosedOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"with_cancelled_and_rejected":MkBool(true) }); _ = request;
  return this.FetchOpenOrders(symbol , since , limit , this.Extend(request , params ));
}

func (this *Bitpanda) FetchOrderTrades(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"order_id":id }); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("max_page_size") )= OpCopy(limit );
  }
  response := this.Call(MkString("privateGetAccountOrdersOrderIdTrades"), this.Extend(request , params )); _ = response;
  tradeHistory := this.SafeValue(response , MkString("trade_history") , MkArray(&VarArray{})); _ = tradeHistory;
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
  }
  return this.ParseTrades(tradeHistory , market , since , limit );
}

func (this *Bitpanda) FetchMyTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("instrument_code") )= *(market ).At(MkString("id") );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    to := this.SafeString(params , MkString("to") ); _ = to;
    if IsTrue(OpEq2(to , MkUndefined() )) {
      panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchMyTrades() requires a \"to\" iso8601 string param with the since argument is specified, max range is 100 days") )));
    }
    *(request ).At(MkString("from") )= this.Iso8601(since );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("max_page_size") )= OpCopy(limit );
  }
  response := this.Call(MkString("privateGetAccountTrades"), this.Extend(request , params )); _ = response;
  tradeHistory := this.SafeValue(response , MkString("trade_history") , MkArray(&VarArray{})); _ = tradeHistory;
  return this.ParseTrades(tradeHistory , market , since , limit );
}

func (this *Bitpanda) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("public") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  url := OpAdd(*(*(*this.At(MkString("urls")) ).At(MkString("api") )).At(api ), OpAdd(MkString("/") , OpAdd(*this.At(MkString("version")) , OpAdd(MkString("/") , this.ImplodeParams(path , params ))))); _ = url;
  query := this.Omit(params , this.ExtractParams(path )); _ = query;
  if IsTrue(OpEq2(api , MkString("public") )) {
    if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
      OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(query )));
    }
  } else {
    if IsTrue(OpEq2(api , MkString("private") )) {
      this.CheckRequiredCredentials();
      headers = MkMap(&VarMap{
          "Accept":MkString("application/json") ,
          "Authorization":OpAdd(MkString("Bearer ") , *this.At(MkString("apiKey")) ),
          });
      if IsTrue(OpEq2(method , MkString("POST") )) {
        body = this.Json(query );
        *(headers ).At(MkString("Content-Type") )= MkString("application/json") ;
      } else {
        if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
          OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(query )));
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

func (this *Bitpanda) HandleErrors(goArgs ...*Variant) *Variant{
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
  message := this.SafeString(response , MkString("error") ); _ = message;
  if IsTrue(OpNotEq2(message , MkUndefined() )) {
    feedback := OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , body )); _ = feedback;
    this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("exact") ), message , feedback );
    this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("broad") ), message , feedback );
    panic(NewExchangeError(feedback ));
  }
  return MkUndefined()
}

