package ccxt

import (
)

type Bitfinex2 struct {
	*ExchangeBase
}

var _ Exchange = (*Bitfinex2)(nil)

func init() {
	exchange := &Bitfinex2{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Bitfinex2) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("bitfinex2") ,
            "name":MkString("Bitfinex") ,
            "countries":MkArray(&VarArray{
                MkString("VG") ,
                }),
            "version":MkString("v2") ,
            "certified":MkBool(false) ,
            "pro":MkBool(false) ,
            "has":MkMap(&VarMap{
                "CORS":MkBool(false) ,
                "cancelAllOrders":MkBool(true) ,
                "cancelOrder":MkBool(true) ,
                "createDepositAddress":MkBool(true) ,
                "createLimitOrder":MkBool(true) ,
                "createMarketOrder":MkBool(true) ,
                "createOrder":MkBool(true) ,
                "deposit":MkBool(false) ,
                "editOrder":MkBool(false) ,
                "fetchBalance":MkBool(true) ,
                "fetchClosedOrder":MkBool(true) ,
                "fetchClosedOrders":MkBool(false) ,
                "fetchCurrencies":MkBool(true) ,
                "fetchDepositAddress":MkBool(true) ,
                "fetchFundingFees":MkBool(false) ,
                "fetchMyTrades":MkBool(true) ,
                "fetchOHLCV":MkBool(true) ,
                "fetchOpenOrder":MkBool(true) ,
                "fetchOpenOrders":MkBool(true) ,
                "fetchOrder":MkBool(false) ,
                "fetchOrderTrades":MkBool(true) ,
                "fetchStatus":MkBool(true) ,
                "fetchTickers":MkBool(true) ,
                "fetchTradingFee":MkBool(false) ,
                "fetchTradingFees":MkBool(false) ,
                "fetchTransactions":MkBool(true) ,
                "withdraw":MkBool(true) ,
                }),
            "timeframes":MkMap(&VarMap{
                "1m":MkString("1m") ,
                "5m":MkString("5m") ,
                "15m":MkString("15m") ,
                "30m":MkString("30m") ,
                "1h":MkString("1h") ,
                "3h":MkString("3h") ,
                "4h":MkString("4h") ,
                "6h":MkString("6h") ,
                "12h":MkString("12h") ,
                "1d":MkString("1D") ,
                "1w":MkString("7D") ,
                "2w":MkString("14D") ,
                "1M":MkString("1M") ,
                }),
            "rateLimit":MkInteger(1500) ,
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/1294454/27766244-e328a50c-5ed2-11e7-947b-041416579bb3.jpg") ,
                "api":MkMap(&VarMap{
                    "v1":MkString("https://api.bitfinex.com") ,
                    "public":MkString("https://api-pub.bitfinex.com") ,
                    "private":MkString("https://api.bitfinex.com") ,
                    }),
                "www":MkString("https://www.bitfinex.com") ,
                "doc":MkArray(&VarArray{
                    MkString("https://docs.bitfinex.com/v2/docs/") ,
                    MkString("https://github.com/bitfinexcom/bitfinex-api-node") ,
                    }),
                "fees":MkString("https://www.bitfinex.com/fees") ,
                }),
            "api":MkMap(&VarMap{
                "v1":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("symbols") ,
                        MkString("symbols_details") ,
                        })}),
                "public":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("conf/{config}") ,
                        MkString("conf/pub:{action}:{object}") ,
                        MkString("conf/pub:{action}:{object}:{detail}") ,
                        MkString("conf/pub:map:{object}") ,
                        MkString("conf/pub:map:{object}:{detail}") ,
                        MkString("conf/pub:map:currency:{detail}") ,
                        MkString("conf/pub:map:currency:sym") ,
                        MkString("conf/pub:map:currency:label") ,
                        MkString("conf/pub:map:currency:unit") ,
                        MkString("conf/pub:map:currency:undl") ,
                        MkString("conf/pub:map:currency:pool") ,
                        MkString("conf/pub:map:currency:explorer") ,
                        MkString("conf/pub:map:currency:tx:fee") ,
                        MkString("conf/pub:map:tx:method") ,
                        MkString("conf/pub:list:{object}") ,
                        MkString("conf/pub:list:{object}:{detail}") ,
                        MkString("conf/pub:list:currency") ,
                        MkString("conf/pub:list:pair:exchange") ,
                        MkString("conf/pub:list:pair:margin") ,
                        MkString("conf/pub:list:pair:futures") ,
                        MkString("conf/pub:list:competitions") ,
                        MkString("conf/pub:info:{object}") ,
                        MkString("conf/pub:info:{object}:{detail}") ,
                        MkString("conf/pub:info:pair") ,
                        MkString("conf/pub:info:tx:status") ,
                        MkString("conf/pub:fees") ,
                        MkString("platform/status") ,
                        MkString("tickers") ,
                        MkString("ticker/{symbol}") ,
                        MkString("trades/{symbol}/hist") ,
                        MkString("book/{symbol}/{precision}") ,
                        MkString("book/{symbol}/P0") ,
                        MkString("book/{symbol}/P1") ,
                        MkString("book/{symbol}/P2") ,
                        MkString("book/{symbol}/P3") ,
                        MkString("book/{symbol}/R0") ,
                        MkString("stats1/{key}:{size}:{symbol}:{side}/{section}") ,
                        MkString("stats1/{key}:{size}:{symbol}:{side}/last") ,
                        MkString("stats1/{key}:{size}:{symbol}:{side}/hist") ,
                        MkString("stats1/{key}:{size}:{symbol}/{section}") ,
                        MkString("stats1/{key}:{size}:{symbol}/last") ,
                        MkString("stats1/{key}:{size}:{symbol}/hist") ,
                        MkString("stats1/{key}:{size}:{symbol}:long/last") ,
                        MkString("stats1/{key}:{size}:{symbol}:long/hist") ,
                        MkString("stats1/{key}:{size}:{symbol}:short/last") ,
                        MkString("stats1/{key}:{size}:{symbol}:short/hist") ,
                        MkString("candles/trade:{timeframe}:{symbol}:{period}/{section}") ,
                        MkString("candles/trade:{timeframe}:{symbol}/{section}") ,
                        MkString("candles/trade:{timeframe}:{symbol}/last") ,
                        MkString("candles/trade:{timeframe}:{symbol}/hist") ,
                        MkString("status/{type}") ,
                        MkString("status/deriv") ,
                        MkString("liquidations/hist") ,
                        MkString("rankings/{key}:{timeframe}:{symbol}/{section}") ,
                        MkString("rankings/{key}:{timeframe}:{symbol}/hist") ,
                        MkString("pulse/hist") ,
                        MkString("pulse/profile/{nickname}") ,
                        MkString("funding/stats/{symbol}/hist") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("calc/trade/avg") ,
                        MkString("calc/fx") ,
                        }),
                    }),
                "private":MkMap(&VarMap{"post":MkArray(&VarArray{
                        MkString("auth/r/wallets") ,
                        MkString("auth/r/wallets/hist") ,
                        MkString("auth/r/orders") ,
                        MkString("auth/r/orders/{symbol}") ,
                        MkString("auth/w/order/submit") ,
                        MkString("auth/w/order/update") ,
                        MkString("auth/w/order/cancel") ,
                        MkString("auth/w/order/multi") ,
                        MkString("auth/w/order/cancel/multi") ,
                        MkString("auth/r/orders/{symbol}/hist") ,
                        MkString("auth/r/orders/hist") ,
                        MkString("auth/r/order/{symbol}:{id}/trades") ,
                        MkString("auth/r/trades/{symbol}/hist") ,
                        MkString("auth/r/trades/hist") ,
                        MkString("auth/r/ledgers/{currency}/hist") ,
                        MkString("auth/r/ledgers/hist") ,
                        MkString("auth/r/info/margin/{key}") ,
                        MkString("auth/r/info/margin/base") ,
                        MkString("auth/r/info/margin/sym_all") ,
                        MkString("auth/r/positions") ,
                        MkString("auth/w/position/claim") ,
                        MkString("auth/r/positions/hist") ,
                        MkString("auth/r/positions/audit") ,
                        MkString("auth/r/positions/snap") ,
                        MkString("auth/w/deriv/collateral/set") ,
                        MkString("auth/w/deriv/collateral/limits") ,
                        MkString("auth/r/funding/offers") ,
                        MkString("auth/r/funding/offers/{symbol}") ,
                        MkString("auth/w/funding/offer/submit") ,
                        MkString("auth/w/funding/offer/cancel") ,
                        MkString("auth/w/funding/offer/cancel/all") ,
                        MkString("auth/w/funding/close") ,
                        MkString("auth/w/funding/auto") ,
                        MkString("auth/w/funding/keep") ,
                        MkString("auth/r/funding/offers/{symbol}/hist") ,
                        MkString("auth/r/funding/offers/hist") ,
                        MkString("auth/r/funding/loans") ,
                        MkString("auth/r/funding/loans/hist") ,
                        MkString("auth/r/funding/loans/{symbol}") ,
                        MkString("auth/r/funding/loans/{symbol}/hist") ,
                        MkString("auth/r/funding/credits") ,
                        MkString("auth/r/funding/credits/hist") ,
                        MkString("auth/r/funding/credits/{symbol}") ,
                        MkString("auth/r/funding/credits/{symbol}/hist") ,
                        MkString("auth/r/funding/trades/{symbol}/hist") ,
                        MkString("auth/r/funding/trades/hist") ,
                        MkString("auth/r/info/funding/{key}") ,
                        MkString("auth/r/info/user") ,
                        MkString("auth/r/logins/hist") ,
                        MkString("auth/w/transfer") ,
                        MkString("auth/w/deposit/address") ,
                        MkString("auth/w/deposit/invoice") ,
                        MkString("auth/w/withdraw") ,
                        MkString("auth/r/movements/{currency}/hist") ,
                        MkString("auth/r/movements/hist") ,
                        MkString("auth/r/alerts") ,
                        MkString("auth/w/alert/set") ,
                        MkString("auth/w/alert/price:{symbol}:{price}/del") ,
                        MkString("auth/w/alert/{type}:{symbol}:{price}/del") ,
                        MkString("auth/calc/order/avail") ,
                        MkString("auth/w/settings/set") ,
                        MkString("auth/r/settings") ,
                        MkString("auth/w/settings/del") ,
                        })}),
                }),
            "fees":MkMap(&VarMap{
                "trading":MkMap(&VarMap{
                    "feeSide":MkString("get") ,
                    "percentage":MkBool(true) ,
                    "tierBased":MkBool(true) ,
                    "maker":this.ParseNumber(MkString("0.001") ),
                    "taker":this.ParseNumber(MkString("0.002") ),
                    "tiers":MkMap(&VarMap{
                        "taker":MkArray(&VarArray{
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("0") ),
                                this.ParseNumber(MkString("0.002") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("500000") ),
                                this.ParseNumber(MkString("0.002") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("1000000") ),
                                this.ParseNumber(MkString("0.002") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("2500000") ),
                                this.ParseNumber(MkString("0.002") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("5000000") ),
                                this.ParseNumber(MkString("0.002") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("7500000") ),
                                this.ParseNumber(MkString("0.002") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("10000000") ),
                                this.ParseNumber(MkString("0.0018") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("15000000") ),
                                this.ParseNumber(MkString("0.0016") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("20000000") ),
                                this.ParseNumber(MkString("0.0014") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("25000000") ),
                                this.ParseNumber(MkString("0.0012") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("30000000") ),
                                this.ParseNumber(MkString("0.001") ),
                                }),
                            }),
                        "maker":MkArray(&VarArray{
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("0") ),
                                this.ParseNumber(MkString("0.001") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("500000") ),
                                this.ParseNumber(MkString("0.0008") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("1000000") ),
                                this.ParseNumber(MkString("0.0006") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("2500000") ),
                                this.ParseNumber(MkString("0.0004") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("5000000") ),
                                this.ParseNumber(MkString("0.0002") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("7500000") ),
                                this.ParseNumber(MkString("0") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("10000000") ),
                                this.ParseNumber(MkString("0") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("15000000") ),
                                this.ParseNumber(MkString("0") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("20000000") ),
                                this.ParseNumber(MkString("0") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("25000000") ),
                                this.ParseNumber(MkString("0") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("30000000") ),
                                this.ParseNumber(MkString("0") ),
                                }),
                            }),
                        }),
                    }),
                "funding":MkMap(&VarMap{"withdraw":MkMap(&VarMap{})}),
                }),
            "options":MkMap(&VarMap{
                "precision":MkString("R0") ,
                "exchangeTypes":MkMap(&VarMap{
                    "EXCHANGE MARKET":MkString("market") ,
                    "EXCHANGE LIMIT":MkString("limit") ,
                    }),
                "orderTypes":MkMap(&VarMap{
                    "market":MkString("EXCHANGE MARKET") ,
                    "limit":MkString("EXCHANGE LIMIT") ,
                    }),
                "fiat":MkMap(&VarMap{
                    "USD":MkString("USD") ,
                    "EUR":MkString("EUR") ,
                    "JPY":MkString("JPY") ,
                    "GBP":MkString("GBP") ,
                    }),
                "v2AccountsByType":MkMap(&VarMap{
                    "spot":MkString("exchange") ,
                    "exchange":MkString("exchange") ,
                    "funding":MkString("funding") ,
                    "margin":MkString("margin") ,
                    "derivatives":MkString("margin") ,
                    }),
                }),
            "exceptions":MkMap(&VarMap{
                "exact":MkMap(&VarMap{
                    "10001":PermissionDenied ,
                    "10020":BadRequest ,
                    "10100":AuthenticationError ,
                    "10114":InvalidNonce ,
                    "20060":OnMaintenance ,
                    }),
                "broad":MkMap(&VarMap{
                    "address":InvalidAddress ,
                    "available balance is only":InsufficientFunds ,
                    "not enough exchange balance":InsufficientFunds ,
                    "Order not found":OrderNotFound ,
                    "symbol: invalid":BadSymbol ,
                    "Invalid order":InvalidOrder ,
                    }),
                }),
            }));
}

func (this *Bitfinex2) IsFiat(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  return OpHasMember(code , *(*this.At(MkString("options")) ).At(MkString("fiat") ));
}

func (this *Bitfinex2) GetCurrencyId(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  return OpAdd(MkString("f") , code );
}

func (this *Bitfinex2) FetchStatus(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetPlatformStatus"), params ); _ = response;
  status := this.SafeInteger(response , MkInteger(0) ); _ = status;
  formattedStatus := OpTrinary(OpEq2(status , MkInteger(1) ), MkString("ok") , MkString("maintenance") ); _ = formattedStatus;
  *this.At(MkString("status")) = this.Extend(*this.At(MkString("status")) , MkMap(&VarMap{
          "status":formattedStatus ,
          "updated":this.Milliseconds(),
          }));
  return *this.At(MkString("status")) ;
}

func (this *Bitfinex2) FetchMarkets(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  v2response := this.Call(MkString("publicGetConfPubListPairFutures"), params ); _ = v2response;
  v1response := this.Call(MkString("v1GetSymbolsDetails"), params ); _ = v1response;
  futuresMarketIds := this.SafeValue(v2response , MkInteger(0) , MkArray(&VarArray{})); _ = futuresMarketIds;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , v1response.Length )); OpIncr(& i ){
    market := *(v1response ).At(i ); _ = market;
    id := this.SafeStringUpper(market , MkString("pair") ); _ = id;
    spot := OpCopy(MkBool(true) ); _ = spot;
    if IsTrue(this.InArray(id , futuresMarketIds )) {
      spot = OpCopy(MkBool(false) );
    }
    futures := OpNot(spot ); _ = futures;
    type_ := OpTrinary(spot , MkString("spot") , MkString("futures") ); _ = type_;
    baseId := OpCopy(MkUndefined() ); _ = baseId;
    quoteId := OpCopy(MkUndefined() ); _ = quoteId;
    if IsTrue(OpGtEq(id.IndexOf(MkString(":") ), MkInteger(0) )) {
      parts := id.Split(MkString(":") ); _ = parts;
      baseId = *(parts ).At(MkInteger(0) );
      quoteId = *(parts ).At(MkInteger(1) );
    } else {
      baseId = id.Slice(MkInteger(0) , MkInteger(3) );
      quoteId = id.Slice(MkInteger(3) , MkInteger(6) );
    }
    base := this.SafeCurrencyCode(baseId ); _ = base;
    quote := this.SafeCurrencyCode(quoteId ); _ = quote;
    symbol := OpAdd(base , OpAdd(MkString("/") , quote )); _ = symbol;
    id = OpAdd(MkString("t") , id );
    baseId = this.GetCurrencyId(baseId );
    quoteId = this.GetCurrencyId(quoteId );
    precision := MkMap(&VarMap{
          "price":this.SafeInteger(market , MkString("price_precision") ),
          "amount":MkInteger(8) ,
          }); _ = precision;
    minOrderSizeString := this.SafeString(market , MkString("minimum_order_size") ); _ = minOrderSizeString;
    maxOrderSizeString := this.SafeString(market , MkString("maximum_order_size") ); _ = maxOrderSizeString;
    limits := MkMap(&VarMap{
          "amount":MkMap(&VarMap{
              "min":this.ParseNumber(minOrderSizeString ),
              "max":this.ParseNumber(maxOrderSizeString ),
              }),
          "price":MkMap(&VarMap{
              "min":this.ParseNumber(MkString("1e-8") ),
              "max":MkUndefined() ,
              }),
          }); _ = limits;
    *(limits ).At(MkString("cost") )= MkMap(&VarMap{
        "min":MkUndefined() ,
        "max":MkUndefined() ,
        });
    margin := this.SafeValue(market , MkString("margin") ); _ = margin;
    result.Push(MkMap(&VarMap{
            "id":id ,
            "symbol":symbol ,
            "base":base ,
            "quote":quote ,
            "baseId":baseId ,
            "quoteId":quoteId ,
            "active":MkBool(true) ,
            "precision":precision ,
            "limits":limits ,
            "info":market ,
            "type":type_ ,
            "swap":MkBool(false) ,
            "spot":spot ,
            "margin":margin ,
            "futures":futures ,
            }));
  }
  return result ;
}

func (this *Bitfinex2) FetchCurrencies(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  labels := MkArray(&VarArray{
        MkString("pub:list:currency") ,
        MkString("pub:map:currency:sym") ,
        MkString("pub:map:currency:label") ,
        MkString("pub:map:currency:unit") ,
        MkString("pub:map:currency:undl") ,
        MkString("pub:map:currency:pool") ,
        MkString("pub:map:currency:explorer") ,
        MkString("pub:map:currency:tx:fee") ,
        }); _ = labels;
  config := labels.Join(MkString(",") ); _ = config;
  request := MkMap(&VarMap{"config":config }); _ = request;
  response := this.Call(MkString("publicGetConfConfig"), this.Extend(request , params )); _ = response;
  indexed := MkMap(&VarMap{
        "sym":this.IndexBy(this.SafeValue(response , MkInteger(1) , MkArray(&VarArray{})), MkInteger(0) ),
        "label":this.IndexBy(this.SafeValue(response , MkInteger(2) , MkArray(&VarArray{})), MkInteger(0) ),
        "unit":this.IndexBy(this.SafeValue(response , MkInteger(3) , MkArray(&VarArray{})), MkInteger(0) ),
        "undl":this.IndexBy(this.SafeValue(response , MkInteger(4) , MkArray(&VarArray{})), MkInteger(0) ),
        "pool":this.IndexBy(this.SafeValue(response , MkInteger(5) , MkArray(&VarArray{})), MkInteger(0) ),
        "explorer":this.IndexBy(this.SafeValue(response , MkInteger(6) , MkArray(&VarArray{})), MkInteger(0) ),
        "fees":this.IndexBy(this.SafeValue(response , MkInteger(7) , MkArray(&VarArray{})), MkInteger(0) ),
        }); _ = indexed;
  ids := this.SafeValue(response , MkInteger(0) , MkArray(&VarArray{})); _ = ids;
  result := MkMap(&VarMap{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , ids.Length )); OpIncr(& i ){
    id := *(ids ).At(i ); _ = id;
    code := this.SafeCurrencyCode(id ); _ = code;
    label := this.SafeValue(*(indexed ).At(MkString("label") ), id , MkArray(&VarArray{})); _ = label;
    name := this.SafeString(label , MkInteger(1) ); _ = name;
    pool := this.SafeValue(*(indexed ).At(MkString("pool") ), id , MkArray(&VarArray{})); _ = pool;
    type_ := this.SafeString(pool , MkInteger(1) ); _ = type_;
    feeValues := this.SafeValue(*(indexed ).At(MkString("fees") ), id , MkArray(&VarArray{})); _ = feeValues;
    fees := this.SafeValue(feeValues , MkInteger(1) , MkArray(&VarArray{})); _ = fees;
    fee := this.SafeNumber(fees , MkInteger(1) ); _ = fee;
    undl := this.SafeValue(*(indexed ).At(MkString("undl") ), id , MkArray(&VarArray{})); _ = undl;
    precision := MkInteger(8) ; _ = precision;
    fid := OpAdd(MkString("f") , id ); _ = fid;
    *(result ).At(code )= MkMap(&VarMap{
        "id":fid ,
        "code":code ,
        "info":MkArray(&VarArray{
            id ,
            label ,
            pool ,
            feeValues ,
            undl ,
            }),
        "type":type_ ,
        "name":name ,
        "active":MkBool(true) ,
        "fee":fee ,
        "precision":precision ,
        "limits":MkMap(&VarMap{
            "amount":MkMap(&VarMap{
                "min":OpDiv(MkInteger(1) , Math.Pow(MkInteger(10) , precision )),
                "max":MkUndefined() ,
                }),
            "withdraw":MkMap(&VarMap{
                "min":fee ,
                "max":MkUndefined() ,
                }),
            }),
        });
  }
  return result ;
}

func (this *Bitfinex2) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  accountsByType := this.SafeValue(*this.At(MkString("options")) , MkString("v2AccountsByType") , MkMap(&VarMap{})); _ = accountsByType;
  requestedType := this.SafeString(params , MkString("type") , MkString("exchange") ); _ = requestedType;
  accountType := this.SafeString(accountsByType , requestedType ); _ = accountType;
  if IsTrue(OpEq2(accountType , MkUndefined() )) {
    keys := GoGetKeys(accountsByType ); _ = keys;
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" fetchBalance type parameter must be one of ") , keys.Join(MkString(", ") )))));
  }
  isDerivative := OpEq2(requestedType , MkString("derivatives") ); _ = isDerivative;
  query := this.Omit(params , MkString("type") ); _ = query;
  response := this.Call(MkString("privatePostAuthRWallets"), query ); _ = response;
  result := MkMap(&VarMap{"info":response }); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , response.Length )); OpIncr(& i ){
    balance := *(response ).At(i ); _ = balance;
    type_ := this.SafeString(balance , MkInteger(0) ); _ = type_;
    currencyId := this.SafeStringLower(balance , MkInteger(1) , MkString("") ); _ = currencyId;
    start := OpSub(currencyId.Length , MkInteger(2) ); _ = start;
    isDerivativeCode := OpEq2(currencyId.Slice(start ), MkString("f0") ); _ = isDerivativeCode;
    derivativeCondition := OpOr(OpNot(isDerivative ), isDerivativeCode ); _ = derivativeCondition;
    if IsTrue(OpAnd(OpEq2(accountType , type_ ), derivativeCondition )) {
      code := this.SafeCurrencyCode(currencyId ); _ = code;
      account := this.Account(); _ = account;
      *(account ).At(MkString("total") )= this.SafeString(balance , MkInteger(2) );
      *(account ).At(MkString("free") )= this.SafeString(balance , MkInteger(4) );
      *(result ).At(code )= OpCopy(account );
    }
  }
  return this.ParseBalance(result );
}

func (this *Bitfinex2) Transfer(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  amount := GoGetArg(goArgs, 1, MkUndefined()); _ = amount;
  fromAccount := GoGetArg(goArgs, 2, MkUndefined()); _ = fromAccount;
  toAccount := GoGetArg(goArgs, 3, MkUndefined()); _ = toAccount;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  accountsByType := this.SafeValue(*this.At(MkString("options")) , MkString("v2AccountsByType") , MkMap(&VarMap{})); _ = accountsByType;
  fromId := this.SafeString(accountsByType , fromAccount ); _ = fromId;
  if IsTrue(OpEq2(fromId , MkUndefined() )) {
    keys := GoGetKeys(accountsByType ); _ = keys;
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" transfer fromAccount must be one of ") , keys.Join(MkString(", ") )))));
  }
  toId := this.SafeString(accountsByType , toAccount ); _ = toId;
  if IsTrue(OpEq2(toId , MkUndefined() )) {
    keys := GoGetKeys(accountsByType ); _ = keys;
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" transfer toAccount must be one of ") , keys.Join(MkString(", ") )))));
  }
  currency := this.Currency(code ); _ = currency;
  fromCurrencyId := this.ConvertDerivativesId(currency , fromAccount ); _ = fromCurrencyId;
  toCurrencyId := this.ConvertDerivativesId(currency , toAccount ); _ = toCurrencyId;
  requestedAmount := this.CurrencyToPrecision(code , amount ); _ = requestedAmount;
  request := MkMap(&VarMap{
        "amount":requestedAmount ,
        "currency":fromCurrencyId ,
        "currency_to":toCurrencyId ,
        "from":fromId ,
        "to":toId ,
        }); _ = request;
  response := this.Call(MkString("privatePostAuthWTransfer"), this.Extend(request , params )); _ = response;
  timestamp := this.SafeInteger(response , MkInteger(0) ); _ = timestamp;
  error := this.SafeString(response , MkInteger(0) ); _ = error;
  if IsTrue(OpEq2(error , MkString("error") )) {
    message := this.SafeString(response , MkInteger(2) , MkString("") ); _ = message;
    this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("exact") ), message , OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , message )));
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , message ))));
  }
  info := this.SafeValue(response , MkInteger(4) ); _ = info;
  fromResponse := this.SafeString(info , MkInteger(1) ); _ = fromResponse;
  toResponse := this.SafeString(info , MkInteger(2) ); _ = toResponse;
  toCode := this.SafeCurrencyCode(this.SafeString(info , MkInteger(5) )); _ = toCode;
  success := this.SafeString(response , MkInteger(6) ); _ = success;
  status := OpTrinary(OpEq2(success , MkString("SUCCESS") ), MkString("ok") , MkUndefined() ); _ = status;
  return MkMap(&VarMap{
        "info":response ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "status":status ,
        "amount":requestedAmount ,
        "code":toCode ,
        "fromAccount":fromResponse ,
        "toAccount":toResponse ,
        });
}

func (this *Bitfinex2) ConvertDerivativesId(goArgs ...*Variant) *Variant{
  currency := GoGetArg(goArgs, 0, MkUndefined()); _ = currency;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  info := this.SafeValue(currency , MkString("info") ); _ = info;
  transferId := this.SafeString(info , MkInteger(0) ); _ = transferId;
  underlying := this.SafeValue(info , MkInteger(4) , MkArray(&VarArray{})); _ = underlying;
  currencyId := OpCopy(MkUndefined() ); _ = currencyId;
  if IsTrue(OpEq2(type_ , MkString("derivatives") )) {
    currencyId = this.SafeString(underlying , MkInteger(0) , transferId );
    start := OpSub(currencyId.Length , MkInteger(2) ); _ = start;
    isDerivativeCode := OpEq2(currencyId.Slice(start ), MkString("F0") ); _ = isDerivativeCode;
    if IsTrue(OpNot(isDerivativeCode )) {
      currencyId = OpAdd(currencyId , MkString("F0") );
    }
  } else {
    if IsTrue(OpNotEq2(type_ , MkString("margin") )) {
      currencyId = this.SafeString(underlying , MkInteger(1) , transferId );
    } else {
      currencyId = OpCopy(transferId );
    }
  }
  return currencyId ;
}

func (this *Bitfinex2) FetchOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  panic(NewNotSupported(OpAdd(*this.At(MkString("id")) , MkString(" fetchOrder is not implemented yet") )));
  return MkUndefined()
}

func (this *Bitfinex2) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  precision := this.SafeValue(*this.At(MkString("options")) , MkString("precision") , MkString("R0") ); _ = precision;
  request := MkMap(&VarMap{
        "symbol":this.MarketId(symbol ),
        "precision":precision ,
        }); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("len") )= OpCopy(limit );
  }
  fullRequest := this.Extend(request , params ); _ = fullRequest;
  orderbook := this.Call(MkString("publicGetBookSymbolPrecision"), fullRequest ); _ = orderbook;
  timestamp := this.Milliseconds(); _ = timestamp;
  result := MkMap(&VarMap{
        "symbol":symbol ,
        "bids":MkArray(&VarArray{}),
        "asks":MkArray(&VarArray{}),
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "nonce":MkUndefined() ,
        }); _ = result;
  priceIndex := OpTrinary(OpEq2(*(fullRequest ).At(MkString("precision") ), MkString("R0") ), MkInteger(1) , MkInteger(0) ); _ = priceIndex;
  for i := MkInteger(0) ; IsTrue(OpLw(i , orderbook.Length )); OpIncr(& i ){
    order := *(orderbook ).At(i ); _ = order;
    price := this.SafeNumber(order , priceIndex ); _ = price;
    signedAmount := this.SafeNumber(order , MkInteger(2) ); _ = signedAmount;
    amount := Math.Abs(signedAmount ); _ = amount;
    side := OpTrinary(OpGt(signedAmount , MkInteger(0) ), MkString("bids") , MkString("asks") ); _ = side;
    (*(result ).At(side )).Call(MkString("push") , MkArray(&VarArray{
            price ,
            amount ,
            }));
  }
  *(result ).At(MkString("bids") )= this.SortBy(*(result ).At(MkString("bids") ), MkInteger(0) , MkBool(true) );
  *(result ).At(MkString("asks") )= this.SortBy(*(result ).At(MkString("asks") ), MkInteger(0) );
  return result ;
}

func (this *Bitfinex2) ParseTicker(goArgs ...*Variant) *Variant{
  ticker := GoGetArg(goArgs, 0, MkUndefined()); _ = ticker;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.Milliseconds(); _ = timestamp;
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  if IsTrue(OpNotEq2(market , MkUndefined() )) {
    symbol = *(market ).At(MkString("symbol") );
  }
  length := OpCopy(ticker.Length ); _ = length;
  last := this.SafeNumber(ticker , OpSub(length , MkInteger(4) )); _ = last;
  return MkMap(&VarMap{
        "symbol":symbol ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "high":this.SafeNumber(ticker , OpSub(length , MkInteger(2) )),
        "low":this.SafeNumber(ticker , OpSub(length , MkInteger(1) )),
        "bid":this.SafeNumber(ticker , OpSub(length , MkInteger(10) )),
        "bidVolume":MkUndefined() ,
        "ask":this.SafeNumber(ticker , OpSub(length , MkInteger(8) )),
        "askVolume":MkUndefined() ,
        "vwap":MkUndefined() ,
        "open":MkUndefined() ,
        "close":last ,
        "last":last ,
        "previousClose":MkUndefined() ,
        "change":this.SafeNumber(ticker , OpSub(length , MkInteger(6) )),
        "percentage":OpMulti(this.SafeNumber(ticker , OpSub(length , MkInteger(5) )), MkInteger(100) ),
        "average":MkUndefined() ,
        "baseVolume":this.SafeNumber(ticker , OpSub(length , MkInteger(3) )),
        "quoteVolume":MkUndefined() ,
        "info":ticker ,
        });
}

func (this *Bitfinex2) FetchTickers(goArgs ...*Variant) *Variant{
  symbols := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbols;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  if IsTrue(OpNotEq2(symbols , MkUndefined() )) {
    ids := this.MarketIds(symbols ); _ = ids;
    *(request ).At(MkString("symbols") )= ids.Join(MkString(",") );
  } else {
    *(request ).At(MkString("symbols") )= MkString("ALL") ;
  }
  tickers := this.Call(MkString("publicGetTickers"), this.Extend(request , params )); _ = tickers;
  result := MkMap(&VarMap{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , tickers.Length )); OpIncr(& i ){
    ticker := *(tickers ).At(i ); _ = ticker;
    id := *(ticker ).At(MkInteger(0) ); _ = id;
    if IsTrue(OpHasMember(id , *this.At(MkString("markets_by_id")) )) {
      market := *(*this.At(MkString("markets_by_id")) ).At(id ); _ = market;
      symbol := *(market ).At(MkString("symbol") ); _ = symbol;
      *(result ).At(symbol )= this.ParseTicker(ticker , market );
    }
  }
  return this.FilterByArray(result , MkString("symbol") , symbols );
}

func (this *Bitfinex2) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"symbol":*(market ).At(MkString("id") )}); _ = request;
  ticker := this.Call(MkString("publicGetTickerSymbol"), this.Extend(request , params )); _ = ticker;
  return this.ParseTicker(ticker , market );
}

func (this *Bitfinex2) ParseSymbol(goArgs ...*Variant) *Variant{
  marketId := GoGetArg(goArgs, 0, MkUndefined()); _ = marketId;
  if IsTrue(OpEq2(marketId , MkUndefined() )) {
    return marketId ;
  }
  marketId = marketId.Replace(MkString("t") , MkString("") );
  baseId := OpCopy(MkUndefined() ); _ = baseId;
  quoteId := OpCopy(MkUndefined() ); _ = quoteId;
  if IsTrue(OpGtEq(marketId.IndexOf(MkString(":") ), MkInteger(0) )) {
    parts := marketId.Split(MkString(":") ); _ = parts;
    baseId = *(parts ).At(MkInteger(0) );
    quoteId = *(parts ).At(MkInteger(1) );
  } else {
    baseId = marketId.Slice(MkInteger(0) , MkInteger(3) );
    quoteId = marketId.Slice(MkInteger(3) , MkInteger(6) );
  }
  base := this.SafeCurrencyCode(baseId ); _ = base;
  quote := this.SafeCurrencyCode(quoteId ); _ = quote;
  return OpAdd(base , OpAdd(MkString("/") , quote ));
}

func (this *Bitfinex2) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  tradeLength := OpCopy(trade.Length ); _ = tradeLength;
  isPrivate := OpGt(tradeLength , MkInteger(5) ); _ = isPrivate;
  id := this.SafeString(trade , MkInteger(0) ); _ = id;
  amountIndex := OpTrinary(isPrivate , MkInteger(4) , MkInteger(2) ); _ = amountIndex;
  side := OpCopy(MkUndefined() ); _ = side;
  amountString := this.SafeString(trade , amountIndex ); _ = amountString;
  priceIndex := OpTrinary(isPrivate , MkInteger(5) , MkInteger(3) ); _ = priceIndex;
  priceString := this.SafeString(trade , priceIndex ); _ = priceString;
  if IsTrue(OpEq2(*(amountString ).At(MkInteger(0) ), MkString("-") )) {
    side = MkString("sell") ;
    amountString = amountString.Slice(MkInteger(1) );
  } else {
    side = MkString("buy") ;
  }
  amount := this.ParseNumber(amountString ); _ = amount;
  price := this.ParseNumber(priceString ); _ = price;
  cost := this.ParseNumber(Precise.StringMul(priceString , amountString )); _ = cost;
  orderId := OpCopy(MkUndefined() ); _ = orderId;
  takerOrMaker := OpCopy(MkUndefined() ); _ = takerOrMaker;
  type_ := OpCopy(MkUndefined() ); _ = type_;
  fee := OpCopy(MkUndefined() ); _ = fee;
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  timestampIndex := OpTrinary(isPrivate , MkInteger(2) , MkInteger(1) ); _ = timestampIndex;
  timestamp := this.SafeInteger(trade , timestampIndex ); _ = timestamp;
  if IsTrue(isPrivate ) {
    marketId := *(trade ).At(MkInteger(1) ); _ = marketId;
    if IsTrue(OpHasMember(marketId , *this.At(MkString("markets_by_id")) )) {
      market = *(*this.At(MkString("markets_by_id")) ).At(marketId );
      symbol = *(market ).At(MkString("symbol") );
    } else {
      symbol = this.ParseSymbol(marketId );
    }
    orderId = this.SafeString(trade , MkInteger(3) );
    maker := this.SafeInteger(trade , MkInteger(8) ); _ = maker;
    takerOrMaker = OpTrinary(OpEq2(maker , MkInteger(1) ), MkString("maker") , MkString("taker") );
    feeCostString := this.SafeString(trade , MkInteger(9) ); _ = feeCostString;
    feeCostString = Precise.StringNeg(feeCostString );
    feeCost := this.ParseNumber(feeCostString ); _ = feeCost;
    feeCurrencyId := this.SafeString(trade , MkInteger(10) ); _ = feeCurrencyId;
    feeCurrency := this.SafeCurrencyCode(feeCurrencyId ); _ = feeCurrency;
    fee = MkMap(&VarMap{
        "cost":feeCost ,
        "currency":feeCurrency ,
        });
    orderType := *(trade ).At(MkInteger(6) ); _ = orderType;
    type_ = this.SafeString(*(*this.At(MkString("options")) ).At(MkString("exchangeTypes") ), orderType );
  }
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    if IsTrue(OpNotEq2(market , MkUndefined() )) {
      symbol = *(market ).At(MkString("symbol") );
    }
  }
  return MkMap(&VarMap{
        "id":id ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "symbol":symbol ,
        "order":orderId ,
        "side":side ,
        "type":type_ ,
        "takerOrMaker":takerOrMaker ,
        "price":price ,
        "amount":amount ,
        "cost":cost ,
        "fee":fee ,
        "info":trade ,
        });
}

func (this *Bitfinex2) FetchTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  sort := MkString("-1") ; _ = sort;
  request := MkMap(&VarMap{"symbol":*(market ).At(MkString("id") )}); _ = request;
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("start") )= OpCopy(since );
    sort = MkString("1") ;
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  *(request ).At(MkString("sort") )= OpCopy(sort );
  response := this.Call(MkString("publicGetTradesSymbolHist"), this.Extend(request , params )); _ = response;
  trades := this.SortBy(response , MkInteger(1) ); _ = trades;
  return this.ParseTrades(trades , market , MkUndefined() , limit );
}

func (this *Bitfinex2) FetchOHLCV(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  timeframe := GoGetArg(goArgs, 1, MkString("1m") ); _ = timeframe;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkInteger(100) ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  if IsTrue(OpEq2(limit , MkUndefined() )) {
    limit = MkInteger(100) ;
  }
  if IsTrue(OpEq2(since , MkUndefined() )) {
    duration := this.ParseTimeframe(timeframe ); _ = duration;
    since = OpSub(this.Milliseconds(), OpMulti(duration , OpMulti(limit , MkInteger(1000) )));
  }
  request := MkMap(&VarMap{
        "symbol":*(market ).At(MkString("id") ),
        "timeframe":*(*this.At(MkString("timeframes")) ).At(timeframe ),
        "sort":MkInteger(1) ,
        "start":since ,
        "limit":limit ,
        }); _ = request;
  response := this.Call(MkString("publicGetCandlesTradeTimeframeSymbolHist"), this.Extend(request , params )); _ = response;
  return this.ParseOHLCVs(response , market , timeframe , since , limit );
}

func (this *Bitfinex2) ParseOrderStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  if IsTrue(OpEq2(status , MkUndefined() )) {
    return status ;
  }
  parts := status.Split(MkString(" ") ); _ = parts;
  state := this.SafeString(parts , MkInteger(0) ); _ = state;
  statuses := MkMap(&VarMap{
        "ACTIVE":MkString("open") ,
        "PARTIALLY":MkString("open") ,
        "EXECUTED":MkString("closed") ,
        "CANCELED":MkString("canceled") ,
        "INSUFFICIENT":MkString("canceled") ,
        "RSN_DUST":MkString("rejected") ,
        "RSN_PAUSE":MkString("rejected") ,
        }); _ = statuses;
  return this.SafeString(statuses , state , status );
}

func (this *Bitfinex2) ParseOrder(goArgs ...*Variant) *Variant{
  order := GoGetArg(goArgs, 0, MkUndefined()); _ = order;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  id := this.SafeString(order , MkInteger(0) ); _ = id;
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  marketId := this.SafeString(order , MkInteger(3) ); _ = marketId;
  if IsTrue(OpHasMember(marketId , *this.At(MkString("markets_by_id")) )) {
    market = *(*this.At(MkString("markets_by_id")) ).At(marketId );
  } else {
    symbol = this.ParseSymbol(marketId );
  }
  if IsTrue(OpAnd(OpEq2(symbol , MkUndefined() ), OpNotEq2(market , MkUndefined() ))) {
    symbol = *(market ).At(MkString("symbol") );
  }
  timestamp := this.SafeInteger(order , MkInteger(5) ); _ = timestamp;
  remaining := Math.Abs(this.SafeNumber(order , MkInteger(6) )); _ = remaining;
  signedAmount := this.SafeNumber(order , MkInteger(7) ); _ = signedAmount;
  amount := Math.Abs(signedAmount ); _ = amount;
  side := OpTrinary(OpLw(signedAmount , MkInteger(0) ), MkString("sell") , MkString("buy") ); _ = side;
  orderType := this.SafeString(order , MkInteger(8) ); _ = orderType;
  type_ := this.SafeString(this.SafeValue(*this.At(MkString("options")) , MkString("exchangeTypes") ), orderType ); _ = type_;
  status := OpCopy(MkUndefined() ); _ = status;
  statusString := this.SafeString(order , MkInteger(13) ); _ = statusString;
  if IsTrue(OpNotEq2(statusString , MkUndefined() )) {
    parts := statusString.Split(MkString(" @ ") ); _ = parts;
    status = this.ParseOrderStatus(this.SafeString(parts , MkInteger(0) ));
  }
  price := this.SafeNumber(order , MkInteger(16) ); _ = price;
  average := this.SafeNumber(order , MkInteger(17) ); _ = average;
  clientOrderId := this.SafeString(order , MkInteger(2) ); _ = clientOrderId;
  return this.SafeOrder(MkMap(&VarMap{
            "info":order ,
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
            "stopPrice":MkUndefined() ,
            "amount":amount ,
            "cost":MkUndefined() ,
            "average":average ,
            "filled":MkUndefined() ,
            "remaining":remaining ,
            "status":status ,
            "fee":MkUndefined() ,
            "trades":MkUndefined() ,
            }));
}

func (this *Bitfinex2) CreateOrder(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 2, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 3, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 4, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 5, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  orderTypes := this.SafeValue(*this.At(MkString("options")) , MkString("orderTypes") , MkMap(&VarMap{})); _ = orderTypes;
  orderType := this.SafeStringUpper(orderTypes , type_ , type_ ); _ = orderType;
  amount = OpTrinary(OpEq2(side , MkString("sell") ), OpNeg(amount ), amount );
  request := MkMap(&VarMap{
        "symbol":*(market ).At(MkString("id") ),
        "type":orderType ,
        "amount":this.NumberToString(amount ),
        }); _ = request;
  if IsTrue(OpOr(OpEq2(orderType , MkString("LIMIT") ), OpEq2(orderType , MkString("EXCHANGE LIMIT") ))) {
    *(request ).At(MkString("price") )= this.NumberToString(price );
  } else {
    if IsTrue(OpOr(OpEq2(orderType , MkString("STOP") ), OpEq2(orderType , MkString("EXCHANGE STOP") ))) {
      stopPrice := this.SafeNumber(params , MkString("stopPrice") , price ); _ = stopPrice;
      *(request ).At(MkString("price") )= this.NumberToString(stopPrice );
    } else {
      if IsTrue(OpOr(OpEq2(orderType , MkString("STOP LIMIT") ), OpEq2(orderType , MkString("EXCHANGE STOP LIMIT") ))) {
        priceAuxLimit := this.SafeNumber(params , MkString("price_aux_limit") ); _ = priceAuxLimit;
        stopPrice := this.SafeNumber(params , MkString("stopPrice") ); _ = stopPrice;
        if IsTrue(OpEq2(priceAuxLimit , MkUndefined() )) {
          if IsTrue(OpEq2(stopPrice , MkUndefined() )) {
            panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" createOrder() requires a stopPrice parameter or a price_aux_limit parameter for a ") , OpAdd(orderType , MkString(" order") )))));
          } else {
            *(request ).At(MkString("price_aux_limit") )= this.NumberToString(price );
          }
        } else {
          *(request ).At(MkString("price_aux_limit") )= this.NumberToString(priceAuxLimit );
          if IsTrue(OpEq2(stopPrice , MkUndefined() )) {
            stopPrice = OpCopy(price );
          }
        }
        *(request ).At(MkString("price") )= this.NumberToString(stopPrice );
      } else {
        if IsTrue(OpOr(OpEq2(orderType , MkString("TRAILING STOP") ), OpEq2(orderType , MkString("EXCHANGE TRAILING STOP") ))) {
          priceTrailing := this.SafeNumber(params , MkString("price_trailing") ); _ = priceTrailing;
          *(request ).At(MkString("price_trailing") )= this.NumberToString(priceTrailing );
          stopPrice := this.SafeNumber(params , MkString("stopPrice") , price ); _ = stopPrice;
          *(request ).At(MkString("price") )= this.NumberToString(stopPrice );
        } else {
          if IsTrue(OpOr(OpEq2(orderType , MkString("FOK") ), OpOr(OpEq2(orderType , MkString("EXCHANGE FOK") ), OpOr(OpEq2(orderType , MkString("IOC") ), OpEq2(orderType , MkString("EXCHANGE IOC") ))))) {
            *(request ).At(MkString("price") )= this.NumberToString(price );
          }
        }
      }
    }
  }
  params = this.Omit(params , MkArray(&VarArray{
          MkString("stopPrice") ,
          MkString("price_aux_limit") ,
          MkString("price_trailing") ,
          }));
  clientOrderId := this.SafeValue2(params , MkString("cid") , MkString("clientOrderId") ); _ = clientOrderId;
  if IsTrue(OpNotEq2(clientOrderId , MkUndefined() )) {
    *(request ).At(MkString("cid") )= OpCopy(clientOrderId );
    params = this.Omit(params , MkArray(&VarArray{
            MkString("cid") ,
            MkString("clientOrderId") ,
            }));
  }
  response := this.Call(MkString("privatePostAuthWOrderSubmit"), this.Extend(request , params )); _ = response;
  status := this.SafeString(response , MkInteger(6) ); _ = status;
  if IsTrue(OpNotEq2(status , MkString("SUCCESS") )) {
    errorCode := *(response ).At(MkInteger(5) ); _ = errorCode;
    errorText := *(response ).At(MkInteger(7) ); _ = errorText;
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , OpAdd(*(response ).At(MkInteger(6) ), OpAdd(MkString(": ") , OpAdd(errorText , OpAdd(MkString(" (#") , OpAdd(errorCode , MkString(")") )))))))));
  }
  orders := this.SafeValue(response , MkInteger(4) , MkArray(&VarArray{})); _ = orders;
  order := this.SafeValue(orders , MkInteger(0) ); _ = order;
  return this.ParseOrder(order , market );
}

func (this *Bitfinex2) CancelAllOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"all":MkInteger(1) }); _ = request;
  response := this.Call(MkString("privatePostAuthWOrderCancelMulti"), this.Extend(request , params )); _ = response;
  orders := this.SafeValue(response , MkInteger(4) , MkArray(&VarArray{})); _ = orders;
  return this.ParseOrders(orders );
}

func (this *Bitfinex2) CancelOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  cid := this.SafeValue2(params , MkString("cid") , MkString("clientOrderId") ); _ = cid;
  request := OpCopy(MkUndefined() ); _ = request;
  if IsTrue(OpNotEq2(cid , MkUndefined() )) {
    cidDate := this.SafeValue(params , MkString("cidDate") ); _ = cidDate;
    if IsTrue(OpEq2(cidDate , MkUndefined() )) {
      panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")) , MkString(" canceling an order by clientOrderId ('cid') requires both 'cid' and 'cid_date' ('YYYY-MM-DD')") )));
    }
    request = MkMap(&VarMap{
        "cid":cid ,
        "cid_date":cidDate ,
        });
    params = this.Omit(params , MkArray(&VarArray{
            MkString("cid") ,
            MkString("clientOrderId") ,
            }));
  } else {
    request = MkMap(&VarMap{"id":parseInt(id )});
  }
  response := this.Call(MkString("privatePostAuthWOrderCancel"), this.Extend(request , params )); _ = response;
  order := this.SafeValue(response , MkInteger(4) ); _ = order;
  return this.ParseOrder(order );
}

func (this *Bitfinex2) FetchOpenOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"id":MkArray(&VarArray{
            parseInt(id ),
            })}); _ = request;
  orders := this.FetchOpenOrders(symbol , MkUndefined() , MkUndefined() , this.Extend(request , params )); _ = orders;
  order := this.SafeValue(orders , MkInteger(0) ); _ = order;
  if IsTrue(OpEq2(order , MkUndefined() )) {
    panic(NewOrderNotFound(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" order ") , OpAdd(id , MkString(" not found") )))));
  }
  return order ;
}

func (this *Bitfinex2) FetchClosedOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"id":MkArray(&VarArray{
            parseInt(id ),
            })}); _ = request;
  orders := this.FetchClosedOrders(symbol , MkUndefined() , MkUndefined() , this.Extend(request , params )); _ = orders;
  order := this.SafeValue(orders , MkInteger(0) ); _ = order;
  if IsTrue(OpEq2(order , MkUndefined() )) {
    panic(NewOrderNotFound(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" order ") , OpAdd(id , MkString(" not found") )))));
  }
  return order ;
}

func (this *Bitfinex2) FetchOpenOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  market := OpCopy(MkUndefined() ); _ = market;
  response := OpCopy(MkUndefined() ); _ = response;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    response = this.Call(MkString("privatePostAuthROrders"), this.Extend(request , params ));
  } else {
    market = this.Market(symbol );
    *(request ).At(MkString("symbol") )= *(market ).At(MkString("id") );
    response = this.Call(MkString("privatePostAuthROrdersSymbol"), this.Extend(request , params ));
  }
  return this.ParseOrders(response , market , since , limit );
}

func (this *Bitfinex2) FetchClosedOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  market := OpCopy(MkUndefined() ); _ = market;
  response := OpCopy(MkUndefined() ); _ = response;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    response = this.Call(MkString("privatePostAuthROrdersHist"), this.Extend(request , params ));
  } else {
    market = this.Market(symbol );
    *(request ).At(MkString("symbol") )= *(market ).At(MkString("id") );
    response = this.Call(MkString("privatePostAuthROrdersSymbolHist"), this.Extend(request , params ));
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("start") )= OpCopy(since );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  return this.ParseOrders(response , market , since , limit );
}

func (this *Bitfinex2) FetchOrderTrades(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchOrderTrades() requires a symbol argument") )));
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  orderId := parseInt(id ); _ = orderId;
  request := MkMap(&VarMap{
        "id":orderId ,
        "symbol":*(market ).At(MkString("id") ),
        }); _ = request;
  response := this.Call(MkString("privatePostAuthROrderSymbolIdTrades"), this.Extend(request , params )); _ = response;
  return this.ParseTrades(response , market , since , limit );
}

func (this *Bitfinex2) FetchMyTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := OpCopy(MkUndefined() ); _ = market;
  request := MkMap(&VarMap{"end":this.Milliseconds()}); _ = request;
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("start") )= OpCopy(since );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  method := MkString("privatePostAuthRTradesHist") ; _ = method;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("symbol") )= *(market ).At(MkString("id") );
    method = MkString("privatePostAuthRTradesSymbolHist") ;
  }
  response := this.Call(method , this.Extend(request , params )); _ = response;
  return this.ParseTrades(response , market , since , limit );
}

func (this *Bitfinex2) CreateDepositAddress(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"op_renew":MkInteger(1) }); _ = request;
  response := this.FetchDepositAddress(code , this.Extend(request , params )); _ = response;
  return response ;
}

func (this *Bitfinex2) FetchDepositAddress(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  name := this.Call(MkString("getCurrencyName"), code ); _ = name;
  request := MkMap(&VarMap{
        "method":name ,
        "wallet":MkString("exchange") ,
        "op_renew":MkInteger(0) ,
        }); _ = request;
  response := this.Call(MkString("privatePostAuthWDepositAddress"), this.Extend(request , params )); _ = response;
  result := this.SafeValue(response , MkInteger(4) , MkArray(&VarArray{})); _ = result;
  poolAddress := this.SafeString(result , MkInteger(5) ); _ = poolAddress;
  address := OpTrinary(OpEq2(poolAddress , MkUndefined() ), this.SafeString(result , MkInteger(4) ), poolAddress ); _ = address;
  tag := OpTrinary(OpEq2(poolAddress , MkUndefined() ), MkUndefined() , this.SafeString(result , MkInteger(4) )); _ = tag;
  this.CheckAddress(address );
  return MkMap(&VarMap{
        "currency":code ,
        "address":address ,
        "tag":tag ,
        "info":response ,
        });
}

func (this *Bitfinex2) ParseTransactionStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "SUCCESS":MkString("ok") ,
        "ERROR":MkString("failed") ,
        "FAILURE":MkString("failed") ,
        "CANCELED":MkString("canceled") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Bitfinex2) ParseTransaction(goArgs ...*Variant) *Variant{
  transaction := GoGetArg(goArgs, 0, MkUndefined()); _ = transaction;
  currency := GoGetArg(goArgs, 1, MkUndefined() ); _ = currency;
  transactionLength := OpCopy(transaction.Length ); _ = transactionLength;
  timestamp := OpCopy(MkUndefined() ); _ = timestamp;
  updated := OpCopy(MkUndefined() ); _ = updated;
  code := OpCopy(MkUndefined() ); _ = code;
  amount := OpCopy(MkUndefined() ); _ = amount;
  id := OpCopy(MkUndefined() ); _ = id;
  status := OpCopy(MkUndefined() ); _ = status;
  tag := OpCopy(MkUndefined() ); _ = tag;
  type_ := OpCopy(MkUndefined() ); _ = type_;
  feeCost := OpCopy(MkUndefined() ); _ = feeCost;
  txid := OpCopy(MkUndefined() ); _ = txid;
  addressTo := OpCopy(MkUndefined() ); _ = addressTo;
  if IsTrue(OpLw(transactionLength , MkInteger(9) )) {
    data := this.SafeValue(transaction , MkInteger(4) , MkArray(&VarArray{})); _ = data;
    timestamp = this.SafeInteger(transaction , MkInteger(0) );
    if IsTrue(OpNotEq2(currency , MkUndefined() )) {
      code = *(currency ).At(MkString("code") );
    }
    feeCost = this.SafeNumber(data , MkInteger(8) );
    if IsTrue(OpNotEq2(feeCost , MkUndefined() )) {
      feeCost = OpNeg(feeCost );
    }
    amount = this.SafeNumber(data , MkInteger(5) );
    id = this.SafeValue(data , MkInteger(0) );
    status = MkString("ok") ;
    if IsTrue(OpEq2(id , MkInteger(0) )) {
      id = OpCopy(MkUndefined() );
      status = MkString("failed") ;
    }
    tag = this.SafeString(data , MkInteger(3) );
    type_ = MkString("withdrawal") ;
  } else {
    id = this.SafeString(transaction , MkInteger(0) );
    timestamp = this.SafeInteger(transaction , MkInteger(5) );
    updated = this.SafeInteger(transaction , MkInteger(6) );
    status = this.ParseTransactionStatus(this.SafeString(transaction , MkInteger(9) ));
    amount = this.SafeNumber(transaction , MkInteger(12) );
    if IsTrue(OpNotEq2(amount , MkUndefined() )) {
      if IsTrue(OpLw(amount , MkInteger(0) )) {
        type_ = MkString("withdrawal") ;
      } else {
        type_ = MkString("deposit") ;
      }
    }
    feeCost = this.SafeNumber(transaction , MkInteger(13) );
    if IsTrue(OpNotEq2(feeCost , MkUndefined() )) {
      feeCost = OpNeg(feeCost );
    }
    addressTo = this.SafeString(transaction , MkInteger(16) );
    txid = this.SafeString(transaction , MkInteger(20) );
  }
  return MkMap(&VarMap{
        "info":transaction ,
        "id":id ,
        "txid":txid ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "addressFrom":MkUndefined() ,
        "address":addressTo ,
        "addressTo":addressTo ,
        "tagFrom":MkUndefined() ,
        "tag":tag ,
        "tagTo":tag ,
        "type":type_ ,
        "amount":amount ,
        "currency":code ,
        "status":status ,
        "updated":updated ,
        "fee":MkMap(&VarMap{
            "currency":code ,
            "cost":feeCost ,
            "rate":MkUndefined() ,
            }),
        });
}

func (this *Bitfinex2) FetchTransactions(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  currency := OpCopy(MkUndefined() ); _ = currency;
  request := MkMap(&VarMap{}); _ = request;
  method := MkString("privatePostAuthRMovementsHist") ; _ = method;
  if IsTrue(OpNotEq2(code , MkUndefined() )) {
    currency = this.Currency(code );
    *(request ).At(MkString("currency") )= *(currency ).At(MkString("id") );
    method = MkString("privatePostAuthRMovementsCurrencyHist") ;
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("start") )= OpCopy(since );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(method , this.Extend(request , params )); _ = response;
  return this.ParseTransactions(response , currency , since , limit );
}

func (this *Bitfinex2) Withdraw(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  amount := GoGetArg(goArgs, 1, MkUndefined()); _ = amount;
  address := GoGetArg(goArgs, 2, MkUndefined()); _ = address;
  tag := GoGetArg(goArgs, 3, MkUndefined() ); _ = tag;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.CheckAddress(address );
  this.LoadMarkets();
  currency := this.Currency(code ); _ = currency;
  name := this.Call(MkString("getCurrencyName"), code ); _ = name;
  request := MkMap(&VarMap{
        "method":name ,
        "wallet":MkString("exchange") ,
        "amount":this.NumberToString(amount ),
        "address":address ,
        }); _ = request;
  if IsTrue(OpNotEq2(tag , MkUndefined() )) {
    *(request ).At(MkString("payment_id") )= OpCopy(tag );
  }
  response := this.Call(MkString("privatePostAuthWWithdraw"), this.Extend(request , params )); _ = response;
  text := this.SafeString(response , MkInteger(7) ); _ = text;
  if IsTrue(OpNotEq2(text , MkString("success") )) {
    this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("broad") ), text , text );
  }
  transaction := this.ParseTransaction(response , currency ); _ = transaction;
  return this.Extend(transaction , MkMap(&VarMap{"address":address }));
}

func (this *Bitfinex2) FetchPositions(goArgs ...*Variant) *Variant{
  symbols := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbols;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("privatePostPositions"), params ); _ = response;
  return response ;
}

func (this *Bitfinex2) Nonce(goArgs ...*Variant) *Variant{
  return this.Milliseconds();
}

func (this *Bitfinex2) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("public") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  request := OpAdd(MkString("/") , this.ImplodeParams(path , params )); _ = request;
  query := this.Omit(params , this.ExtractParams(path )); _ = query;
  if IsTrue(OpEq2(api , MkString("v1") )) {
    request = OpAdd(api , request );
  } else {
    request = OpAdd(*this.At(MkString("version")) , request );
  }
  url := OpAdd(*(*(*this.At(MkString("urls")) ).At(MkString("api") )).At(api ), OpAdd(MkString("/") , request )); _ = url;
  if IsTrue(OpEq2(api , MkString("public") )) {
    if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
      OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(query )));
    }
  }
  if IsTrue(OpEq2(api , MkString("private") )) {
    this.CheckRequiredCredentials();
    nonce := (this.Nonce()).Call(MkString("toString") ); _ = nonce;
    body = this.Json(query );
    auth := OpAdd(MkString("/api/") , OpAdd(request , OpAdd(nonce , body ))); _ = auth;
    signature := this.Hmac(this.Encode(auth ), this.Encode(*this.At(MkString("secret")) ), MkString("sha384") ); _ = signature;
    headers = MkMap(&VarMap{
        "bfx-nonce":nonce ,
        "bfx-apikey":*this.At(MkString("apiKey")) ,
        "bfx-signature":signature ,
        "Content-Type":MkString("application/json") ,
        });
  }
  return MkMap(&VarMap{
        "url":url ,
        "method":method ,
        "body":body ,
        "headers":headers ,
        });
}

func (this *Bitfinex2) HandleErrors(goArgs ...*Variant) *Variant{
  statusCode := GoGetArg(goArgs, 0, MkUndefined()); _ = statusCode;
  statusText := GoGetArg(goArgs, 1, MkUndefined()); _ = statusText;
  url := GoGetArg(goArgs, 2, MkUndefined()); _ = url;
  method := GoGetArg(goArgs, 3, MkUndefined()); _ = method;
  responseHeaders := GoGetArg(goArgs, 4, MkUndefined()); _ = responseHeaders;
  responseBody := GoGetArg(goArgs, 5, MkUndefined()); _ = responseBody;
  response := GoGetArg(goArgs, 6, MkUndefined()); _ = response;
  requestHeaders := GoGetArg(goArgs, 7, MkUndefined()); _ = requestHeaders;
  requestBody := GoGetArg(goArgs, 8, MkUndefined()); _ = requestBody;
  if IsTrue(OpNotEq2(response , MkUndefined() )) {
    if IsTrue(OpNot(GoIsArray(response ))) {
      message := this.SafeString(response , MkString("message") ); _ = message;
      if IsTrue(OpAnd(OpNotEq2(message , MkUndefined() ), OpGtEq(message.IndexOf(MkString("not enough exchange balance") ), MkInteger(0) ))) {
        panic(NewInsufficientFunds(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , this.Json(response )))));
      }
      panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , this.Json(response )))));
    }
  } else {
    if IsTrue(OpEq2(response , MkString("") )) {
      panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , MkString(" returned empty response") )));
    }
  }
  if IsTrue(OpEq2(statusCode , MkInteger(500) )) {
    errorCode := this.NumberToString(*(response ).At(MkInteger(1) )); _ = errorCode;
    errorText := *(response ).At(MkInteger(2) ); _ = errorText;
    feedback := OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , errorText )); _ = feedback;
    this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("exact") ), errorCode , feedback );
    this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("exact") ), errorText , feedback );
    this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("broad") ), errorText , feedback );
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , OpAdd(errorText , OpAdd(MkString(" (#") , OpAdd(errorCode , MkString(")") )))))));
  }
  return response ;
}

