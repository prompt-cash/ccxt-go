package ccxt

import (
)

type Phemex struct {
	*ExchangeBase
}

var _ Exchange = (*Phemex)(nil)

func init() {
	exchange := &Phemex{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Phemex) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("phemex") ,
            "name":MkString("Phemex") ,
            "countries":MkArray(&VarArray{
                MkString("CN") ,
                }),
            "rateLimit":MkInteger(100) ,
            "version":MkString("v1") ,
            "certified":MkBool(false) ,
            "pro":MkBool(true) ,
            "hostname":MkString("api.phemex.com") ,
            "has":MkMap(&VarMap{
                "cancelAllOrders":MkBool(true) ,
                "cancelOrder":MkBool(true) ,
                "createOrder":MkBool(true) ,
                "fetchBalance":MkBool(true) ,
                "fetchClosedOrders":MkBool(true) ,
                "fetchCurrencies":MkBool(true) ,
                "fetchDepositAddress":MkBool(true) ,
                "fetchDeposits":MkBool(true) ,
                "fetchMarkets":MkBool(true) ,
                "fetchMyTrades":MkBool(true) ,
                "fetchOHLCV":MkBool(true) ,
                "fetchOpenOrders":MkBool(true) ,
                "fetchOrder":MkBool(true) ,
                "fetchOrderBook":MkBool(true) ,
                "fetchOrders":MkBool(true) ,
                "fetchTicker":MkBool(true) ,
                "fetchTrades":MkBool(true) ,
                "fetchWithdrawals":MkBool(true) ,
                }),
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/1294454/85225056-221eb600-b3d7-11ea-930d-564d2690e3f6.jpg") ,
                "test":MkMap(&VarMap{
                    "v1":MkString("https://testnet-api.phemex.com/v1") ,
                    "public":MkString("https://testnet-api.phemex.com/exchange/public") ,
                    "private":MkString("https://testnet-api.phemex.com") ,
                    }),
                "api":MkMap(&VarMap{
                    "v1":MkString("https://{hostname}/v1") ,
                    "public":MkString("https://{hostname}/exchange/public") ,
                    "private":MkString("https://{hostname}") ,
                    }),
                "www":MkString("https://phemex.com") ,
                "doc":MkString("https://github.com/phemex/phemex-api-docs") ,
                "fees":MkString("https://phemex.com/fees-conditions") ,
                "referral":MkMap(&VarMap{
                    "url":MkString("https://phemex.com/register?referralCode=EDNVJ") ,
                    "discount":MkNumber(0.1) ,
                    }),
                }),
            "timeframes":MkMap(&VarMap{
                "1m":MkString("60") ,
                "3m":MkString("180") ,
                "5m":MkString("300") ,
                "15m":MkString("900") ,
                "30m":MkString("1800") ,
                "1h":MkString("3600") ,
                "2h":MkString("7200") ,
                "3h":MkString("10800") ,
                "4h":MkString("14400") ,
                "6h":MkString("21600") ,
                "12h":MkString("43200") ,
                "1d":MkString("86400") ,
                "1w":MkString("604800") ,
                "1M":MkString("2592000") ,
                }),
            "api":MkMap(&VarMap{
                "public":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("cfg/v2/products") ,
                        MkString("products") ,
                        MkString("nomics/trades") ,
                        MkString("md/kline") ,
                        })}),
                "v1":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("md/orderbook") ,
                        MkString("md/trade") ,
                        MkString("md/ticker/24hr") ,
                        MkString("md/ticker/24hr/all") ,
                        MkString("md/spot/ticker/24hr") ,
                        MkString("md/spot/ticker/24hr/all") ,
                        MkString("exchange/public/products") ,
                        })}),
                "private":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("spot/orders/active") ,
                        MkString("spot/orders") ,
                        MkString("spot/wallets") ,
                        MkString("exchange/spot/order") ,
                        MkString("exchange/spot/order/trades") ,
                        MkString("accounts/accountPositions") ,
                        MkString("accounts/positions") ,
                        MkString("orders/activeList") ,
                        MkString("exchange/order/list") ,
                        MkString("exchange/order") ,
                        MkString("exchange/order/trade") ,
                        MkString("phemex-user/users/children") ,
                        MkString("phemex-user/wallets/v2/depositAddress") ,
                        MkString("exchange/margins/transfer") ,
                        MkString("exchange/wallets/confirm/withdraw") ,
                        MkString("exchange/wallets/withdrawList") ,
                        MkString("exchange/wallets/depositList") ,
                        MkString("exchange/wallets/v2/depositAddress") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("spot/orders") ,
                        MkString("orders") ,
                        MkString("positions/assign") ,
                        MkString("exchange/wallets/transferOut") ,
                        MkString("exchange/wallets/transferIn") ,
                        MkString("exchange/margins") ,
                        MkString("exchange/wallets/createWithdraw") ,
                        MkString("exchange/wallets/cancelWithdraw") ,
                        MkString("exchange/wallets/createWithdrawAddress") ,
                        }),
                    "put":MkArray(&VarArray{
                        MkString("spot/orders") ,
                        MkString("orders/replace") ,
                        MkString("positions/leverage") ,
                        MkString("positions/riskLimit") ,
                        }),
                    "delete":MkArray(&VarArray{
                        MkString("spot/orders") ,
                        MkString("orders/cancel") ,
                        MkString("orders") ,
                        MkString("orders/all") ,
                        }),
                    }),
                }),
            "precisionMode":TICK_SIZE ,
            "fees":MkMap(&VarMap{"trading":MkMap(&VarMap{
                    "tierBased":MkBool(false) ,
                    "percentage":MkBool(true) ,
                    "taker":this.ParseNumber(MkString("0.001") ),
                    "maker":this.ParseNumber(MkString("0.001") ),
                    })}),
            "requiredCredentials":MkMap(&VarMap{
                "apiKey":MkBool(true) ,
                "secret":MkBool(true) ,
                }),
            "exceptions":MkMap(&VarMap{
                "exact":MkMap(&VarMap{
                    "412":BadRequest ,
                    "6001":BadRequest ,
                    "19999":BadRequest ,
                    "10001":DuplicateOrderId ,
                    "10002":OrderNotFound ,
                    "10003":CancelPending ,
                    "10004":CancelPending ,
                    "10005":CancelPending ,
                    "11001":InsufficientFunds ,
                    "11002":InvalidOrder ,
                    "11003":InsufficientFunds ,
                    "11004":InvalidOrder ,
                    "11005":InsufficientFunds ,
                    "11006":ExchangeError ,
                    "11007":ExchangeError ,
                    "11008":ExchangeError ,
                    "11009":ExchangeError ,
                    "11010":InsufficientFunds ,
                    "11011":InvalidOrder ,
                    "11012":InvalidOrder ,
                    "11013":InvalidOrder ,
                    "11014":InvalidOrder ,
                    "11015":InvalidOrder ,
                    "11016":BadRequest ,
                    "11017":ExchangeError ,
                    "11018":ExchangeError ,
                    "11019":ExchangeError ,
                    "11020":ExchangeError ,
                    "11021":ExchangeError ,
                    "11022":AccountSuspended ,
                    "11023":ExchangeError ,
                    "11024":ExchangeError ,
                    "11025":BadRequest ,
                    "11026":ExchangeError ,
                    "11027":BadSymbol ,
                    "11028":BadSymbol ,
                    "11029":ExchangeError ,
                    "11030":ExchangeError ,
                    "11031":DDoSProtection ,
                    "11032":DDoSProtection ,
                    "11033":DuplicateOrderId ,
                    "11034":InvalidOrder ,
                    "11035":InvalidOrder ,
                    "11036":InvalidOrder ,
                    "11037":InvalidOrder ,
                    "11038":InvalidOrder ,
                    "11039":InvalidOrder ,
                    "11040":InvalidOrder ,
                    "11041":InvalidOrder ,
                    "11042":InvalidOrder ,
                    "11043":InvalidOrder ,
                    "11044":InvalidOrder ,
                    "11045":InvalidOrder ,
                    "11046":InvalidOrder ,
                    "11047":InvalidOrder ,
                    "11048":InvalidOrder ,
                    "11049":InvalidOrder ,
                    "11050":InvalidOrder ,
                    "11051":InvalidOrder ,
                    "11052":InvalidOrder ,
                    "11053":InvalidOrder ,
                    "11054":InvalidOrder ,
                    "11055":InvalidOrder ,
                    "11056":InvalidOrder ,
                    "11057":InvalidOrder ,
                    "11058":InvalidOrder ,
                    "11059":InvalidOrder ,
                    "11060":InvalidOrder ,
                    "11061":CancelPending ,
                    "11062":InvalidOrder ,
                    "11063":InvalidOrder ,
                    "11064":InvalidOrder ,
                    "11065":InvalidOrder ,
                    "11066":InvalidOrder ,
                    "11067":InvalidOrder ,
                    "11068":InvalidOrder ,
                    "11069":ExchangeError ,
                    "11070":BadSymbol ,
                    "11071":InvalidOrder ,
                    "11072":InvalidOrder ,
                    "11073":InvalidOrder ,
                    "11074":InvalidOrder ,
                    "11075":InvalidOrder ,
                    "11076":InvalidOrder ,
                    "11077":InvalidOrder ,
                    "11078":InvalidOrder ,
                    "11079":InvalidOrder ,
                    "11080":InvalidOrder ,
                    "11081":InvalidOrder ,
                    "11082":InsufficientFunds ,
                    "11083":InvalidOrder ,
                    "11084":InvalidOrder ,
                    "11085":DuplicateOrderId ,
                    "11086":InvalidOrder ,
                    "11087":InvalidOrder ,
                    "11088":InvalidOrder ,
                    "11089":InvalidOrder ,
                    "11090":InvalidOrder ,
                    "11091":InvalidOrder ,
                    "11092":InvalidOrder ,
                    "11093":InvalidOrder ,
                    "11094":InvalidOrder ,
                    "11095":InvalidOrder ,
                    "11096":InvalidOrder ,
                    "11097":BadRequest ,
                    "11098":BadRequest ,
                    "11099":ExchangeError ,
                    "11100":InsufficientFunds ,
                    "11101":InsufficientFunds ,
                    "11102":BadRequest ,
                    "11103":BadRequest ,
                    "11104":BadRequest ,
                    "11105":InsufficientFunds ,
                    "11106":InsufficientFunds ,
                    "11107":ExchangeError ,
                    "11108":InvalidOrder ,
                    "11109":InvalidOrder ,
                    "11110":InvalidOrder ,
                    "11111":InvalidOrder ,
                    "11112":InvalidOrder ,
                    "11113":BadRequest ,
                    "11114":InvalidOrder ,
                    "11115":InvalidOrder ,
                    "30018":BadRequest ,
                    "39996":PermissionDenied ,
                    }),
                "broad":MkMap(&VarMap{
                    "Failed to find api-key":AuthenticationError ,
                    "Missing required parameter":BadRequest ,
                    "API Signature verification failed":AuthenticationError ,
                    }),
                }),
            "options":MkMap(&VarMap{
                "x-phemex-request-expiry":MkInteger(60) ,
                "createOrderByQuoteRequiresPrice":MkBool(true) ,
                }),
            }));
}

func (this *Phemex) ParseSafeNumber(goArgs ...*Variant) *Variant{
  value := GoGetArg(goArgs, 0, MkUndefined() ); _ = value;
  if IsTrue(OpEq2(value , MkUndefined() )) {
    return value ;
  }
  value = value.Replace(MkString(",") , MkString("") );
  parts := value.Split(MkString(" ") ); _ = parts;
  return this.SafeNumber(parts , MkInteger(0) );
}

func (this *Phemex) ParseSwapMarket(goArgs ...*Variant) *Variant{
  market := GoGetArg(goArgs, 0, MkUndefined()); _ = market;
  id := this.SafeString(market , MkString("symbol") ); _ = id;
  baseId := this.SafeString2(market , MkString("baseCurrency") , MkString("contractUnderlyingAssets") ); _ = baseId;
  quoteId := this.SafeString(market , MkString("quoteCurrency") ); _ = quoteId;
  base := this.SafeCurrencyCode(baseId ); _ = base;
  quote := this.SafeCurrencyCode(quoteId ); _ = quote;
  symbol := OpAdd(base , OpAdd(MkString("/") , quote )); _ = symbol;
  type_ := this.SafeStringLower(market , MkString("type") ); _ = type_;
  inverse := OpCopy(MkBool(false) ); _ = inverse;
  spot := OpCopy(MkBool(false) ); _ = spot;
  swap := OpCopy(MkBool(true) ); _ = swap;
  settlementCurrencyId := this.SafeString(market , MkString("settlementCurrency") ); _ = settlementCurrencyId;
  if IsTrue(OpNotEq2(settlementCurrencyId , quoteId )) {
    inverse = OpCopy(MkBool(true) );
  }
  linear := OpNot(inverse ); _ = linear;
  precision := MkMap(&VarMap{
        "amount":this.SafeNumber(market , MkString("lotSize") ),
        "price":this.SafeNumber(market , MkString("tickSize") ),
        }); _ = precision;
  priceScale := this.SafeInteger(market , MkString("priceScale") ); _ = priceScale;
  ratioScale := this.SafeInteger(market , MkString("ratioScale") ); _ = ratioScale;
  valueScale := this.SafeInteger(market , MkString("valueScale") ); _ = valueScale;
  minPriceEp := this.SafeString(market , MkString("minPriceEp") ); _ = minPriceEp;
  maxPriceEp := this.SafeString(market , MkString("maxPriceEp") ); _ = maxPriceEp;
  makerFeeRateEr := this.SafeString(market , MkString("makerFeeRateEr") ); _ = makerFeeRateEr;
  takerFeeRateEr := this.SafeString(market , MkString("takerFeeRateEr") ); _ = takerFeeRateEr;
  maker := this.ParseNumber(this.FromEn(makerFeeRateEr , ratioScale )); _ = maker;
  taker := this.ParseNumber(this.FromEn(takerFeeRateEr , ratioScale )); _ = taker;
  limits := MkMap(&VarMap{
        "amount":MkMap(&VarMap{
            "min":*(precision ).At(MkString("amount") ),
            "max":MkUndefined() ,
            }),
        "price":MkMap(&VarMap{
            "min":this.ParseNumber(this.FromEn(minPriceEp , priceScale )),
            "max":this.ParseNumber(this.FromEn(maxPriceEp , priceScale )),
            }),
        "cost":MkMap(&VarMap{
            "min":MkUndefined() ,
            "max":this.ParseNumber(this.SafeString(market , MkString("maxOrderQty") )),
            }),
        }); _ = limits;
  status := this.SafeString(market , MkString("status") ); _ = status;
  active := OpEq2(status , MkString("Listed") ); _ = active;
  return MkMap(&VarMap{
        "id":id ,
        "symbol":symbol ,
        "base":base ,
        "quote":quote ,
        "baseId":baseId ,
        "quoteId":quoteId ,
        "info":market ,
        "type":type_ ,
        "spot":spot ,
        "swap":swap ,
        "linear":linear ,
        "inverse":inverse ,
        "active":active ,
        "taker":taker ,
        "maker":maker ,
        "priceScale":priceScale ,
        "valueScale":valueScale ,
        "ratioScale":ratioScale ,
        "precision":precision ,
        "limits":limits ,
        });
}

func (this *Phemex) ParseSpotMarket(goArgs ...*Variant) *Variant{
  market := GoGetArg(goArgs, 0, MkUndefined()); _ = market;
  type_ := this.SafeStringLower(market , MkString("type") ); _ = type_;
  id := this.SafeString(market , MkString("symbol") ); _ = id;
  quoteId := this.SafeString(market , MkString("quoteCurrency") ); _ = quoteId;
  baseId := this.SafeString(market , MkString("baseCurrency") ); _ = baseId;
  linear := OpCopy(MkUndefined() ); _ = linear;
  inverse := OpCopy(MkUndefined() ); _ = inverse;
  spot := OpCopy(MkBool(true) ); _ = spot;
  swap := OpCopy(MkBool(false) ); _ = swap;
  taker := this.SafeNumber(market , MkString("defaultTakerFee") ); _ = taker;
  maker := this.SafeNumber(market , MkString("defaultMakerFee") ); _ = maker;
  precision := MkMap(&VarMap{
        "amount":this.ParseSafeNumber(this.SafeString(market , MkString("baseTickSize") )),
        "price":this.ParseSafeNumber(this.SafeString(market , MkString("quoteTickSize") )),
        }); _ = precision;
  limits := MkMap(&VarMap{
        "amount":MkMap(&VarMap{
            "min":*(precision ).At(MkString("amount") ),
            "max":this.ParseSafeNumber(this.SafeString(market , MkString("maxBaseOrderSize") )),
            }),
        "price":MkMap(&VarMap{
            "min":*(precision ).At(MkString("price") ),
            "max":MkUndefined() ,
            }),
        "cost":MkMap(&VarMap{
            "min":this.ParseSafeNumber(this.SafeString(market , MkString("minOrderValue") )),
            "max":this.ParseSafeNumber(this.SafeString(market , MkString("maxOrderValue") )),
            }),
        }); _ = limits;
  base := this.SafeCurrencyCode(baseId ); _ = base;
  quote := this.SafeCurrencyCode(quoteId ); _ = quote;
  symbol := OpAdd(base , OpAdd(MkString("/") , quote )); _ = symbol;
  status := this.SafeString(market , MkString("status") ); _ = status;
  active := OpEq2(status , MkString("Listed") ); _ = active;
  return MkMap(&VarMap{
        "id":id ,
        "symbol":symbol ,
        "base":base ,
        "quote":quote ,
        "baseId":baseId ,
        "quoteId":quoteId ,
        "info":market ,
        "type":type_ ,
        "spot":spot ,
        "swap":swap ,
        "linear":linear ,
        "inverse":inverse ,
        "active":active ,
        "taker":taker ,
        "maker":maker ,
        "precision":precision ,
        "priceScale":MkInteger(8) ,
        "valueScale":MkInteger(8) ,
        "ratioScale":MkInteger(8) ,
        "limits":limits ,
        });
}

func (this *Phemex) FetchMarkets(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  v2Products := this.Call(MkString("publicGetCfgV2Products"), params ); _ = v2Products;
  v1Products := this.Call(MkString("v1GetExchangePublicProducts"), params ); _ = v1Products;
  v1ProductsData := this.SafeValue(v1Products , MkString("data") , MkArray(&VarArray{})); _ = v1ProductsData;
  v2ProductsData := this.SafeValue(v2Products , MkString("data") , MkMap(&VarMap{})); _ = v2ProductsData;
  products := this.SafeValue(v2ProductsData , MkString("products") , MkArray(&VarArray{})); _ = products;
  riskLimits := this.SafeValue(v2ProductsData , MkString("riskLimits") , MkArray(&VarArray{})); _ = riskLimits;
  riskLimitsById := this.IndexBy(riskLimits , MkString("symbol") ); _ = riskLimitsById;
  v1ProductsById := this.IndexBy(v1ProductsData , MkString("symbol") ); _ = v1ProductsById;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , products.Length )); OpIncr(& i ){
    market := *(products ).At(i ); _ = market;
    type_ := this.SafeStringLower(market , MkString("type") ); _ = type_;
    if IsTrue(OpEq2(type_ , MkString("perpetual") )) {
      id := this.SafeString(market , MkString("symbol") ); _ = id;
      riskLimitValues := this.SafeValue(riskLimitsById , id , MkMap(&VarMap{})); _ = riskLimitValues;
      market = this.Extend(market , riskLimitValues );
      v1ProductsValues := this.SafeValue(v1ProductsById , id , MkMap(&VarMap{})); _ = v1ProductsValues;
      market = this.Extend(market , v1ProductsValues );
      market = this.ParseSwapMarket(market );
    } else {
      market = this.ParseSpotMarket(market );
    }
    result.Push(market );
  }
  return result ;
}

func (this *Phemex) FetchCurrencies(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetCfgV2Products"), params ); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  currencies := this.SafeValue(data , MkString("currencies") , MkArray(&VarArray{})); _ = currencies;
  result := MkMap(&VarMap{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , currencies.Length )); OpIncr(& i ){
    currency := *(currencies ).At(i ); _ = currency;
    id := this.SafeString(currency , MkString("currency") ); _ = id;
    name := this.SafeString(currency , MkString("name") ); _ = name;
    code := this.SafeCurrencyCode(id ); _ = code;
    valueScaleString := this.SafeString(currency , MkString("valueScale") ); _ = valueScaleString;
    valueScale := parseInt(valueScaleString ); _ = valueScale;
    minValueEv := this.SafeString(currency , MkString("minValueEv") ); _ = minValueEv;
    maxValueEv := this.SafeString(currency , MkString("maxValueEv") ); _ = maxValueEv;
    minAmount := OpCopy(MkUndefined() ); _ = minAmount;
    maxAmount := OpCopy(MkUndefined() ); _ = maxAmount;
    precision := OpCopy(MkUndefined() ); _ = precision;
    if IsTrue(OpNotEq2(valueScale , MkUndefined() )) {
      precisionString := this.ParsePrecision(valueScaleString ); _ = precisionString;
      precision = this.ParseNumber(precisionString );
      minAmount = this.ParseNumber(Precise.StringMul(minValueEv , precisionString ));
      maxAmount = this.ParseNumber(Precise.StringMul(maxValueEv , precisionString ));
    }
    *(result ).At(code )= MkMap(&VarMap{
        "id":id ,
        "info":currency ,
        "code":code ,
        "name":name ,
        "active":MkUndefined() ,
        "fee":MkUndefined() ,
        "precision":precision ,
        "limits":MkMap(&VarMap{
            "amount":MkMap(&VarMap{
                "min":minAmount ,
                "max":maxAmount ,
                }),
            "withdraw":MkMap(&VarMap{
                "min":MkUndefined() ,
                "max":MkUndefined() ,
                }),
            }),
        "valueScale":valueScale ,
        });
  }
  return result ;
}

func (this *Phemex) ParseBidAsk(goArgs ...*Variant) *Variant{
  bidask := GoGetArg(goArgs, 0, MkUndefined()); _ = bidask;
  priceKey := GoGetArg(goArgs, 1, MkInteger(0) ); _ = priceKey;
  amountKey := GoGetArg(goArgs, 2, MkInteger(1) ); _ = amountKey;
  market := GoGetArg(goArgs, 3, MkUndefined() ); _ = market;
  if IsTrue(OpEq2(market , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" parseBidAsk() requires a market argument") )));
  }
  amount := this.SafeString(bidask , amountKey ); _ = amount;
  if IsTrue(*(market ).At(MkString("spot") )) {
    amount = this.FromEv(amount , market );
  }
  return MkArray(&VarArray{
        this.ParseNumber(this.FromEp(this.SafeString(bidask , priceKey ), market )),
        this.ParseNumber(amount ),
        });
}

func (this *Phemex) ParseOrderBook(goArgs ...*Variant) *Variant{
  orderbook := GoGetArg(goArgs, 0, MkUndefined()); _ = orderbook;
  symbol := GoGetArg(goArgs, 1, MkUndefined()); _ = symbol;
  timestamp := GoGetArg(goArgs, 2, MkUndefined() ); _ = timestamp;
  bidsKey := GoGetArg(goArgs, 3, MkString("bids") ); _ = bidsKey;
  asksKey := GoGetArg(goArgs, 4, MkString("asks") ); _ = asksKey;
  priceKey := GoGetArg(goArgs, 5, MkInteger(0) ); _ = priceKey;
  amountKey := GoGetArg(goArgs, 6, MkInteger(1) ); _ = amountKey;
  market := GoGetArg(goArgs, 7, MkUndefined() ); _ = market;
  result := MkMap(&VarMap{
        "symbol":symbol ,
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

func (this *Phemex) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"symbol":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("v1GetMdOrderbook"), this.Extend(request , params )); _ = response;
  result := this.SafeValue(response , MkString("result") , MkMap(&VarMap{})); _ = result;
  book := this.SafeValue(result , MkString("book") , MkMap(&VarMap{})); _ = book;
  timestamp := this.SafeIntegerProduct(result , MkString("timestamp") , MkNumber(0.000001) ); _ = timestamp;
  orderbook := this.ParseOrderBook(book , symbol , timestamp , MkString("bids") , MkString("asks") , MkInteger(0) , MkInteger(1) , market ); _ = orderbook;
  *(orderbook ).At(MkString("nonce") )= this.SafeInteger(result , MkString("sequence") );
  return orderbook ;
}

func (this *Phemex) ToEn(goArgs ...*Variant) *Variant{
  n := GoGetArg(goArgs, 0, MkUndefined()); _ = n;
  scale := GoGetArg(goArgs, 1, MkUndefined()); _ = scale;
  stringN := n.ToString(); _ = stringN;
  precise := NewPrecise(stringN ); _ = precise;
  precise.Decimals = OpSub(precise.Decimals , scale );
  precise.Reduce();
  stringValue := precise.ToString(); _ = stringValue;
  return parseInt(parseFloat(stringValue ));
}

func (this *Phemex) ToEv(goArgs ...*Variant) *Variant{
  amount := GoGetArg(goArgs, 0, MkUndefined()); _ = amount;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  if IsTrue(OpOr(OpEq2(amount , MkUndefined() ), OpEq2(market , MkUndefined() ))) {
    return amount ;
  }
  return this.ToEn(amount , *(market ).At(MkString("valueScale") ));
}

func (this *Phemex) ToEp(goArgs ...*Variant) *Variant{
  price := GoGetArg(goArgs, 0, MkUndefined()); _ = price;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  if IsTrue(OpOr(OpEq2(price , MkUndefined() ), OpEq2(market , MkUndefined() ))) {
    return price ;
  }
  return this.ToEn(price , *(market ).At(MkString("priceScale") ));
}

func (this *Phemex) FromEn(goArgs ...*Variant) *Variant{
  en := GoGetArg(goArgs, 0, MkUndefined()); _ = en;
  scale := GoGetArg(goArgs, 1, MkUndefined()); _ = scale;
  if IsTrue(OpEq2(en , MkUndefined() )) {
    return MkUndefined() ;
  }
  precise := NewPrecise(en ); _ = precise;
  precise.Decimals = this.Sum(precise.Decimals , scale );
  precise.Reduce();
  return precise.ToString();
}

func (this *Phemex) FromEp(goArgs ...*Variant) *Variant{
  ep := GoGetArg(goArgs, 0, MkUndefined()); _ = ep;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  if IsTrue(OpOr(OpEq2(ep , MkUndefined() ), OpEq2(market , MkUndefined() ))) {
    return ep ;
  }
  return this.FromEn(ep , this.SafeInteger(market , MkString("priceScale") ));
}

func (this *Phemex) FromEv(goArgs ...*Variant) *Variant{
  ev := GoGetArg(goArgs, 0, MkUndefined()); _ = ev;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  if IsTrue(OpOr(OpEq2(ev , MkUndefined() ), OpEq2(market , MkUndefined() ))) {
    return ev ;
  }
  return this.FromEn(ev , this.SafeInteger(market , MkString("valueScale") ));
}

func (this *Phemex) FromEr(goArgs ...*Variant) *Variant{
  er := GoGetArg(goArgs, 0, MkUndefined()); _ = er;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  if IsTrue(OpOr(OpEq2(er , MkUndefined() ), OpEq2(market , MkUndefined() ))) {
    return er ;
  }
  return this.FromEn(er , this.SafeInteger(market , MkString("ratioScale") ));
}

func (this *Phemex) ParseOHLCV(goArgs ...*Variant) *Variant{
  ohlcv := GoGetArg(goArgs, 0, MkUndefined()); _ = ohlcv;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  baseVolume := OpCopy(MkUndefined() ); _ = baseVolume;
  if IsTrue(OpAnd(OpNotEq2(market , MkUndefined() ), *(market ).At(MkString("spot") ))) {
    baseVolume = this.ParseNumber(this.FromEv(this.SafeString(ohlcv , MkInteger(7) ), market ));
  } else {
    baseVolume = this.SafeNumber(ohlcv , MkInteger(7) );
  }
  return MkArray(&VarArray{
        this.SafeTimestamp(ohlcv , MkInteger(0) ),
        this.ParseNumber(this.FromEp(this.SafeString(ohlcv , MkInteger(3) ), market )),
        this.ParseNumber(this.FromEp(this.SafeString(ohlcv , MkInteger(4) ), market )),
        this.ParseNumber(this.FromEp(this.SafeString(ohlcv , MkInteger(5) ), market )),
        this.ParseNumber(this.FromEp(this.SafeString(ohlcv , MkInteger(6) ), market )),
        baseVolume ,
        });
}

func (this *Phemex) FetchOHLCV(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  timeframe := GoGetArg(goArgs, 1, MkString("1m") ); _ = timeframe;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"resolution":*(*this.At(MkString("timeframes")) ).At(timeframe )}); _ = request;
  duration := this.ParseTimeframe(timeframe ); _ = duration;
  now := this.Seconds(); _ = now;
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    if IsTrue(OpEq2(limit , MkUndefined() )) {
      limit = MkInteger(2000) ;
    }
    since = parseInt(OpDiv(since , MkInteger(1000) ));
    *(request ).At(MkString("from") )= OpCopy(since );
    *(request ).At(MkString("to") )= Math.Min(now , this.Sum(since , OpMulti(duration , limit )));
  } else {
    if IsTrue(OpNotEq2(limit , MkUndefined() )) {
      limit = Math.Min(limit , MkInteger(2000) );
      *(request ).At(MkString("from") )= OpSub(now , OpMulti(duration , this.Sum(limit , MkInteger(1) )));
      *(request ).At(MkString("to") )= OpCopy(now );
    } else {
      panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchOHLCV() requires a since argument, or a limit argument, or both") )));
    }
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  *(request ).At(MkString("symbol") )= *(market ).At(MkString("id") );
  response := this.Call(MkString("publicGetMdKline"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  rows := this.SafeValue(data , MkString("rows") , MkArray(&VarArray{})); _ = rows;
  return this.ParseOHLCVs(rows , market , timeframe , since , limit );
}

func (this *Phemex) ParseTicker(goArgs ...*Variant) *Variant{
  ticker := GoGetArg(goArgs, 0, MkUndefined()); _ = ticker;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  marketId := this.SafeString(ticker , MkString("symbol") ); _ = marketId;
  market = this.SafeMarket(marketId , market );
  symbol := *(market ).At(MkString("symbol") ); _ = symbol;
  timestamp := this.SafeIntegerProduct(ticker , MkString("timestamp") , MkNumber(0.000001) ); _ = timestamp;
  lastString := this.FromEp(this.SafeString(ticker , MkString("lastEp") ), market ); _ = lastString;
  last := this.ParseNumber(lastString ); _ = last;
  quoteVolume := this.ParseNumber(this.FromEv(this.SafeString(ticker , MkString("turnoverEv") ), market )); _ = quoteVolume;
  baseVolume := this.SafeNumber(ticker , MkString("volume") ); _ = baseVolume;
  if IsTrue(OpEq2(baseVolume , MkUndefined() )) {
    baseVolume = this.ParseNumber(this.FromEv(this.SafeString(ticker , MkString("volumeEv") ), market ));
  }
  vwap := OpCopy(MkUndefined() ); _ = vwap;
  if IsTrue(OpAnd(OpNotEq2(market , MkUndefined() ), *(market ).At(MkString("spot") ))) {
    vwap = this.Vwap(baseVolume , quoteVolume );
  }
  change := OpCopy(MkUndefined() ); _ = change;
  percentage := OpCopy(MkUndefined() ); _ = percentage;
  average := OpCopy(MkUndefined() ); _ = average;
  openString := this.FromEp(this.SafeString(ticker , MkString("openEp") ), market ); _ = openString;
  open := this.ParseNumber(openString ); _ = open;
  if IsTrue(OpAnd(OpNotEq2(openString , MkUndefined() ), OpNotEq2(lastString , MkUndefined() ))) {
    change = this.ParseNumber(Precise.StringSub(lastString , openString ));
    average = this.ParseNumber(Precise.StringDiv(Precise.StringAdd(lastString , openString ), MkString("2") ));
    percentage = this.ParseNumber(Precise.StringMul(Precise.StringSub(Precise.StringDiv(lastString , openString ), MkString("1") ), MkString("100") ));
  }
  result := MkMap(&VarMap{
        "symbol":symbol ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "high":this.ParseNumber(this.FromEp(this.SafeString(ticker , MkString("highEp") ), market )),
        "low":this.ParseNumber(this.FromEp(this.SafeString(ticker , MkString("lowEp") ), market )),
        "bid":this.ParseNumber(this.FromEp(this.SafeString(ticker , MkString("bidEp") ), market )),
        "bidVolume":MkUndefined() ,
        "ask":this.ParseNumber(this.FromEp(this.SafeString(ticker , MkString("askEp") ), market )),
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
        }); _ = result;
  return result ;
}

func (this *Phemex) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"symbol":*(market ).At(MkString("id") )}); _ = request;
  method := OpTrinary(*(market ).At(MkString("spot") ), MkString("v1GetMdSpotTicker24hr") , MkString("v1GetMdTicker24hr") ); _ = method;
  response := this.Call(method , this.Extend(request , params )); _ = response;
  result := this.SafeValue(response , MkString("result") , MkMap(&VarMap{})); _ = result;
  return this.ParseTicker(result , market );
}

func (this *Phemex) FetchTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"symbol":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("v1GetMdTrade"), this.Extend(request , params )); _ = response;
  result := this.SafeValue(response , MkString("result") , MkMap(&VarMap{})); _ = result;
  trades := this.SafeValue(result , MkString("trades") , MkArray(&VarArray{})); _ = trades;
  return this.ParseTrades(trades , market , since , limit );
}

func (this *Phemex) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  priceString := OpCopy(MkUndefined() ); _ = priceString;
  amountString := OpCopy(MkUndefined() ); _ = amountString;
  timestamp := OpCopy(MkUndefined() ); _ = timestamp;
  id := OpCopy(MkUndefined() ); _ = id;
  side := OpCopy(MkUndefined() ); _ = side;
  costString := OpCopy(MkUndefined() ); _ = costString;
  type_ := OpCopy(MkUndefined() ); _ = type_;
  fee := OpCopy(MkUndefined() ); _ = fee;
  marketId := this.SafeString(trade , MkString("symbol") ); _ = marketId;
  market = this.SafeMarket(marketId , market );
  symbol := *(market ).At(MkString("symbol") ); _ = symbol;
  orderId := OpCopy(MkUndefined() ); _ = orderId;
  takerOrMaker := OpCopy(MkUndefined() ); _ = takerOrMaker;
  if IsTrue(GoIsArray(trade )) {
    tradeLength := OpCopy(trade.Length ); _ = tradeLength;
    timestamp = this.SafeIntegerProduct(trade , MkInteger(0) , MkNumber(0.000001) );
    if IsTrue(OpGt(tradeLength , MkInteger(4) )) {
      id = this.SafeString(trade , OpSub(tradeLength , MkInteger(4) ));
    }
    side = this.SafeStringLower(trade , OpSub(tradeLength , MkInteger(3) ));
    priceString = this.FromEp(this.SafeString(trade , OpSub(tradeLength , MkInteger(2) )), market );
    amountString = this.FromEv(this.SafeString(trade , OpSub(tradeLength , MkInteger(1) )), market );
  } else {
    timestamp = this.SafeIntegerProduct(trade , MkString("transactTimeNs") , MkNumber(0.000001) );
    id = this.SafeString2(trade , MkString("execId") , MkString("execID") );
    orderId = this.SafeString(trade , MkString("orderID") );
    side = this.SafeStringLower(trade , MkString("side") );
    type_ = this.ParseOrderType(this.SafeString(trade , MkString("ordType") ));
    execStatus := this.SafeString(trade , MkString("execStatus") ); _ = execStatus;
    if IsTrue(OpEq2(execStatus , MkString("MakerFill") )) {
      takerOrMaker = MkString("maker") ;
    }
    priceString = this.FromEp(this.SafeString(trade , MkString("execPriceEp") ), market );
    amountString = this.FromEv(this.SafeString(trade , MkString("execBaseQtyEv") ), market );
    amountString = this.SafeString(trade , MkString("execQty") , amountString );
    costString = this.FromEv(this.SafeString2(trade , MkString("execQuoteQtyEv") , MkString("execValueEv") ), market );
    feeCostString := this.FromEv(this.SafeString(trade , MkString("execFeeEv") ), market ); _ = feeCostString;
    if IsTrue(OpNotEq2(feeCostString , MkUndefined() )) {
      feeRateString := this.FromEr(this.SafeString(trade , MkString("feeRateEr") ), market ); _ = feeRateString;
      feeCurrencyCode := OpCopy(MkUndefined() ); _ = feeCurrencyCode;
      if IsTrue(*(market ).At(MkString("spot") )) {
        feeCurrencyCode = OpTrinary(OpEq2(side , MkString("buy") ), *(market ).At(MkString("base") ), *(market ).At(MkString("quote") ));
      } else {
        info := this.SafeValue(market , MkString("info") ); _ = info;
        if IsTrue(OpNotEq2(info , MkUndefined() )) {
          settlementCurrencyId := this.SafeString(info , MkString("settlementCurrency") ); _ = settlementCurrencyId;
          feeCurrencyCode = this.SafeCurrencyCode(settlementCurrencyId );
        }
      }
      fee = MkMap(&VarMap{
          "cost":this.ParseNumber(feeCostString ),
          "rate":this.ParseNumber(feeRateString ),
          "currency":feeCurrencyCode ,
          });
    }
  }
  price := this.ParseNumber(priceString ); _ = price;
  amount := this.ParseNumber(amountString ); _ = amount;
  if IsTrue(OpEq2(costString , MkUndefined() )) {
    costString = Precise.StringMul(priceString , amountString );
  }
  cost := this.ParseNumber(costString ); _ = cost;
  return MkMap(&VarMap{
        "info":trade ,
        "id":id ,
        "symbol":symbol ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "order":orderId ,
        "type":type_ ,
        "side":side ,
        "takerOrMaker":takerOrMaker ,
        "price":price ,
        "amount":amount ,
        "cost":cost ,
        "fee":fee ,
        });
}

func (this *Phemex) ParseSpotBalance(goArgs ...*Variant) *Variant{
  response := GoGetArg(goArgs, 0, MkUndefined()); _ = response;
  timestamp := OpCopy(MkUndefined() ); _ = timestamp;
  result := MkMap(&VarMap{"info":response }); _ = result;
  data := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = data;
  for i := MkInteger(0) ; IsTrue(OpLw(i , data.Length )); OpIncr(& i ){
    balance := *(data ).At(i ); _ = balance;
    currencyId := this.SafeString(balance , MkString("currency") ); _ = currencyId;
    code := this.SafeCurrencyCode(currencyId ); _ = code;
    currency := this.SafeValue(*this.At(MkString("currencies")) , code , MkMap(&VarMap{})); _ = currency;
    scale := this.SafeInteger(currency , MkString("valueScale") , MkInteger(8) ); _ = scale;
    account := this.Account(); _ = account;
    balanceEv := this.SafeString(balance , MkString("balanceEv") ); _ = balanceEv;
    lockedTradingBalanceEv := this.SafeString(balance , MkString("lockedTradingBalanceEv") ); _ = lockedTradingBalanceEv;
    lockedWithdrawEv := this.SafeString(balance , MkString("lockedWithdrawEv") ); _ = lockedWithdrawEv;
    total := this.FromEn(balanceEv , scale ); _ = total;
    lockedTradingBalance := this.FromEn(lockedTradingBalanceEv , scale ); _ = lockedTradingBalance;
    lockedWithdraw := this.FromEn(lockedWithdrawEv , scale ); _ = lockedWithdraw;
    used := Precise.StringAdd(lockedTradingBalance , lockedWithdraw ); _ = used;
    lastUpdateTimeNs := this.SafeIntegerProduct(balance , MkString("lastUpdateTimeNs") , MkNumber(0.000001) ); _ = lastUpdateTimeNs;
    timestamp = OpTrinary(OpEq2(timestamp , MkUndefined() ), lastUpdateTimeNs , Math.Max(timestamp , lastUpdateTimeNs ));
    *(account ).At(MkString("total") )= OpCopy(total );
    *(account ).At(MkString("used") )= OpCopy(used );
    *(result ).At(code )= OpCopy(account );
  }
  *(result ).At(MkString("timestamp") )= OpCopy(timestamp );
  *(result ).At(MkString("datetime") )= this.Iso8601(timestamp );
  return this.ParseBalance(result );
}

func (this *Phemex) ParseSwapBalance(goArgs ...*Variant) *Variant{
  response := GoGetArg(goArgs, 0, MkUndefined()); _ = response;
  result := MkMap(&VarMap{"info":response }); _ = result;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  balance := this.SafeValue(data , MkString("account") , MkMap(&VarMap{})); _ = balance;
  currencyId := this.SafeString(balance , MkString("currency") ); _ = currencyId;
  code := this.SafeCurrencyCode(currencyId ); _ = code;
  currency := this.Currency(code ); _ = currency;
  account := this.Account(); _ = account;
  accountBalanceEv := this.SafeString(balance , MkString("accountBalanceEv") ); _ = accountBalanceEv;
  totalUsedBalanceEv := this.SafeString(balance , MkString("totalUsedBalanceEv") ); _ = totalUsedBalanceEv;
  valueScale := this.SafeInteger(currency , MkString("valueScale") , MkInteger(8) ); _ = valueScale;
  *(account ).At(MkString("total") )= this.FromEn(accountBalanceEv , valueScale );
  *(account ).At(MkString("used") )= this.FromEn(totalUsedBalanceEv , valueScale );
  *(result ).At(code )= OpCopy(account );
  return this.ParseBalance(result );
}

func (this *Phemex) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  defaultType := this.SafeString2(*this.At(MkString("options")) , MkString("defaultType") , MkString("fetchBalance") , MkString("spot") ); _ = defaultType;
  type_ := this.SafeString(params , MkString("type") , defaultType ); _ = type_;
  method := MkString("privateGetSpotWallets") ; _ = method;
  request := MkMap(&VarMap{}); _ = request;
  if IsTrue(OpEq2(type_ , MkString("swap") )) {
    code := this.SafeString(params , MkString("code") ); _ = code;
    if IsTrue(OpNotEq2(code , MkUndefined() )) {
      currency := this.Currency(code ); _ = currency;
      *(request ).At(MkString("currency") )= *(currency ).At(MkString("id") );
      params = this.Omit(params , MkString("code") );
    } else {
      currency := this.SafeString(params , MkString("currency") ); _ = currency;
      if IsTrue(OpEq2(currency , MkUndefined() )) {
        panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" fetchBalance() requires a code parameter or a currency parameter for ") , OpAdd(type_ , MkString(" type") )))));
      }
    }
    method = MkString("privateGetAccountsAccountPositions") ;
  }
  params = this.Omit(params , MkString("type") );
  response := this.Call(method , this.Extend(request , params )); _ = response;
  result := OpTrinary(OpEq2(type_ , MkString("swap") ), this.ParseSwapBalance(response ), this.ParseSpotBalance(response )); _ = result;
  return result ;
}

func (this *Phemex) ParseOrderStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "Created":MkString("open") ,
        "Untriggered":MkString("open") ,
        "Deactivated":MkString("closed") ,
        "Triggered":MkString("open") ,
        "Rejected":MkString("rejected") ,
        "New":MkString("open") ,
        "PartiallyFilled":MkString("open") ,
        "Filled":MkString("closed") ,
        "Canceled":MkString("canceled") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Phemex) ParseOrderType(goArgs ...*Variant) *Variant{
  type_ := GoGetArg(goArgs, 0, MkUndefined()); _ = type_;
  types := MkMap(&VarMap{
        "Limit":MkString("limit") ,
        "Market":MkString("market") ,
        }); _ = types;
  return this.SafeString(types , type_ , type_ );
}

func (this *Phemex) ParseTimeInForce(goArgs ...*Variant) *Variant{
  timeInForce := GoGetArg(goArgs, 0, MkUndefined()); _ = timeInForce;
  timeInForces := MkMap(&VarMap{
        "GoodTillCancel":MkString("GTC") ,
        "PostOnly":MkString("PO") ,
        "ImmediateOrCancel":MkString("IOC") ,
        "FillOrKill":MkString("FOK") ,
        }); _ = timeInForces;
  return this.SafeString(timeInForces , timeInForce , timeInForce );
}

func (this *Phemex) ParseSpotOrder(goArgs ...*Variant) *Variant{
  order := GoGetArg(goArgs, 0, MkUndefined()); _ = order;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  id := this.SafeString(order , MkString("orderID") ); _ = id;
  clientOrderId := this.SafeString(order , MkString("clOrdID") ); _ = clientOrderId;
  if IsTrue(OpAnd(OpNotEq2(clientOrderId , MkUndefined() ), OpLw(clientOrderId.Length , MkInteger(1) ))) {
    clientOrderId = OpCopy(MkUndefined() );
  }
  marketId := this.SafeString(order , MkString("symbol") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market ); _ = symbol;
  price := this.ParseNumber(this.OmitZero(this.FromEp(this.SafeString(order , MkString("priceEp") ), market ))); _ = price;
  amount := this.ParseNumber(this.OmitZero(this.FromEv(this.SafeString(order , MkString("baseQtyEv") ), market ))); _ = amount;
  remaining := this.ParseNumber(this.OmitZero(this.FromEv(this.SafeString(order , MkString("leavesBaseQtyEv") ), market ))); _ = remaining;
  filled := this.ParseNumber(this.OmitZero(this.FromEv(this.SafeString(order , MkString("cumBaseQtyEv") ), market ))); _ = filled;
  cost := this.ParseNumber(this.OmitZero(this.FromEv(this.SafeString(order , MkString("quoteQtyEv") ), market ))); _ = cost;
  average := this.ParseNumber(this.OmitZero(this.FromEp(this.SafeString(order , MkString("avgPriceEp") ), market ))); _ = average;
  status := this.ParseOrderStatus(this.SafeString(order , MkString("ordStatus") )); _ = status;
  side := this.SafeStringLower(order , MkString("side") ); _ = side;
  type_ := this.ParseOrderType(this.SafeString(order , MkString("ordType") )); _ = type_;
  timestamp := this.SafeIntegerProduct2(order , MkString("actionTimeNs") , MkString("createTimeNs") , MkNumber(0.000001) ); _ = timestamp;
  fee := OpCopy(MkUndefined() ); _ = fee;
  feeCost := this.ParseNumber(this.FromEv(this.SafeString(order , MkString("cumFeeEv") ), market )); _ = feeCost;
  if IsTrue(OpNotEq2(feeCost , MkUndefined() )) {
    fee = MkMap(&VarMap{
        "cost":feeCost ,
        "currency":MkUndefined() ,
        });
  }
  timeInForce := this.ParseTimeInForce(this.SafeString(order , MkString("timeInForce") )); _ = timeInForce;
  stopPrice := this.ParseNumber(this.OmitZero(this.FromEp(this.SafeString(order , MkString("stopPxEp") , market )))); _ = stopPrice;
  postOnly := OpEq2(timeInForce , MkString("PO") ); _ = postOnly;
  return this.SafeOrder(MkMap(&VarMap{
            "info":order ,
            "id":id ,
            "clientOrderId":clientOrderId ,
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
            "cost":cost ,
            "average":average ,
            "filled":filled ,
            "remaining":remaining ,
            "status":status ,
            "fee":fee ,
            "trades":MkUndefined() ,
            }));
}

func (this *Phemex) ParseSwapOrder(goArgs ...*Variant) *Variant{
  order := GoGetArg(goArgs, 0, MkUndefined()); _ = order;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  id := this.SafeString(order , MkString("orderID") ); _ = id;
  clientOrderId := this.SafeString(order , MkString("clOrdID") ); _ = clientOrderId;
  if IsTrue(OpAnd(OpNotEq2(clientOrderId , MkUndefined() ), OpLw(clientOrderId.Length , MkInteger(1) ))) {
    clientOrderId = OpCopy(MkUndefined() );
  }
  marketId := this.SafeString(order , MkString("symbol") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market ); _ = symbol;
  status := this.ParseOrderStatus(this.SafeString(order , MkString("ordStatus") )); _ = status;
  side := this.SafeStringLower(order , MkString("side") ); _ = side;
  type_ := this.ParseOrderType(this.SafeString(order , MkString("orderType") )); _ = type_;
  price := this.ParseNumber(this.FromEp(this.SafeString(order , MkString("priceEp") ), market )); _ = price;
  amount := this.SafeNumber(order , MkString("orderQty") ); _ = amount;
  filled := this.SafeNumber(order , MkString("cumQty") ); _ = filled;
  remaining := this.SafeNumber(order , MkString("leavesQty") ); _ = remaining;
  timestamp := this.SafeIntegerProduct(order , MkString("actionTimeNs") , MkNumber(0.000001) ); _ = timestamp;
  cost := this.SafeNumber(order , MkString("cumValue") ); _ = cost;
  lastTradeTimestamp := this.SafeIntegerProduct(order , MkString("transactTimeNs") , MkNumber(0.000001) ); _ = lastTradeTimestamp;
  if IsTrue(OpEq2(lastTradeTimestamp , MkInteger(0) )) {
    lastTradeTimestamp = OpCopy(MkUndefined() );
  }
  timeInForce := this.ParseTimeInForce(this.SafeString(order , MkString("timeInForce") )); _ = timeInForce;
  stopPrice := this.SafeNumber(order , MkString("stopPx") ); _ = stopPrice;
  postOnly := OpEq2(timeInForce , MkString("PO") ); _ = postOnly;
  return MkMap(&VarMap{
        "info":order ,
        "id":id ,
        "clientOrderId":clientOrderId ,
        "datetime":this.Iso8601(timestamp ),
        "timestamp":timestamp ,
        "lastTradeTimestamp":lastTradeTimestamp ,
        "symbol":symbol ,
        "type":type_ ,
        "timeInForce":timeInForce ,
        "postOnly":postOnly ,
        "side":side ,
        "price":price ,
        "stopPrice":stopPrice ,
        "amount":amount ,
        "filled":filled ,
        "remaining":remaining ,
        "cost":cost ,
        "average":MkUndefined() ,
        "status":status ,
        "fee":MkUndefined() ,
        "trades":MkUndefined() ,
        });
}

func (this *Phemex) ParseOrder(goArgs ...*Variant) *Variant{
  order := GoGetArg(goArgs, 0, MkUndefined()); _ = order;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  if IsTrue(OpHasMember(MkString("closedPnl") , order )) {
    return this.ParseSwapOrder(order , market );
  }
  return this.ParseSpotOrder(order , market );
}

func (this *Phemex) CreateOrder(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 2, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 3, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 4, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 5, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  side = this.Capitalize(side );
  type_ = this.Capitalize(type_ );
  request := MkMap(&VarMap{
        "symbol":*(market ).At(MkString("id") ),
        "side":side ,
        "ordType":type_ ,
        }); _ = request;
  if IsTrue(*(market ).At(MkString("spot") )) {
    qtyType := this.SafeValue(params , MkString("qtyType") , MkString("ByBase") ); _ = qtyType;
    if IsTrue(OpOr(OpEq2(type_ , MkString("Market") ), OpOr(OpEq2(type_ , MkString("Stop") ), OpEq2(type_ , MkString("MarketIfTouched") )))) {
      if IsTrue(OpNotEq2(price , MkUndefined() )) {
        qtyType = MkString("ByQuote") ;
      }
    }
    *(request ).At(MkString("qtyType") )= OpCopy(qtyType );
    if IsTrue(OpEq2(qtyType , MkString("ByQuote") )) {
      cost := this.SafeNumber(params , MkString("cost") ); _ = cost;
      params = this.Omit(params , MkString("cost") );
      if IsTrue(*(*this.At(MkString("options")) ).At(MkString("createOrderByQuoteRequiresPrice") )) {
        if IsTrue(OpNotEq2(price , MkUndefined() )) {
          cost = OpMulti(amount , price );
        } else {
          if IsTrue(OpEq2(cost , MkUndefined() )) {
            panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" createOrder() ") , OpAdd(qtyType , MkString(" requires a price argument or a cost parameter") )))));
          }
        }
      }
      cost = OpTrinary(OpEq2(cost , MkUndefined() ), amount , cost );
      costString := cost.ToString(); _ = costString;
      *(request ).At(MkString("quoteQtyEv") )= this.ToEv(costString , market );
    } else {
      amountString := amount.ToString(); _ = amountString;
      *(request ).At(MkString("baseQtyEv") )= this.ToEv(amountString , market );
    }
  } else {
    if IsTrue(*(market ).At(MkString("swap") )) {
      *(request ).At(MkString("orderQty") )= parseInt(amount );
    }
  }
  if IsTrue(OpEq2(type_ , MkString("Limit") )) {
    priceString := price.ToString(); _ = priceString;
    *(request ).At(MkString("priceEp") )= this.ToEp(priceString , market );
  }
  stopPrice := this.SafeString2(params , MkString("stopPx") , MkString("stopPrice") ); _ = stopPrice;
  if IsTrue(OpNotEq2(stopPrice , MkUndefined() )) {
    *(request ).At(MkString("stopPxEp") )= this.ToEp(stopPrice , market );
  }
  params = this.Omit(params , MkArray(&VarArray{
          MkString("stopPx") ,
          MkString("stopPrice") ,
          }));
  method := OpTrinary(*(market ).At(MkString("spot") ), MkString("privatePostSpotOrders") , MkString("privatePostOrders") ); _ = method;
  response := this.Call(method , this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  return this.ParseOrder(data , market );
}

func (this *Phemex) CancelOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" cancelOrder() requires a symbol argument") )));
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"symbol":*(market ).At(MkString("id") )}); _ = request;
  clientOrderId := this.SafeString2(params , MkString("clientOrderId") , MkString("clOrdID") ); _ = clientOrderId;
  params = this.Omit(params , MkArray(&VarArray{
          MkString("clientOrderId") ,
          MkString("clOrdID") ,
          }));
  if IsTrue(OpNotEq2(clientOrderId , MkUndefined() )) {
    *(request ).At(MkString("clOrdID") )= OpCopy(clientOrderId );
  } else {
    *(request ).At(MkString("orderID") )= OpCopy(id );
  }
  method := OpTrinary(*(market ).At(MkString("spot") ), MkString("privateDeleteSpotOrders") , MkString("privateDeleteOrdersCancel") ); _ = method;
  response := this.Call(method , this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  return this.ParseOrder(data , market );
}

func (this *Phemex) CancelAllOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    if IsTrue(OpNot(*(market ).At(MkString("swap") ))) {
      panic(NewNotSupported(OpAdd(*this.At(MkString("id")) , MkString(" cancelAllOrders() supports swap market type orders only") )));
    }
    *(request ).At(MkString("symbol") )= *(market ).At(MkString("id") );
  }
  return this.Call(MkString("privateDeleteOrdersAll"), this.Extend(request , params ));
}

func (this *Phemex) FetchOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchOrder() requires a symbol argument") )));
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  method := OpTrinary(*(market ).At(MkString("spot") ), MkString("privateGetSpotOrdersActive") , MkString("privateGetExchangeOrder") ); _ = method;
  request := MkMap(&VarMap{"symbol":*(market ).At(MkString("id") )}); _ = request;
  clientOrderId := this.SafeString2(params , MkString("clientOrderId") , MkString("clOrdID") ); _ = clientOrderId;
  params = this.Omit(params , MkArray(&VarArray{
          MkString("clientOrderId") ,
          MkString("clOrdID") ,
          }));
  if IsTrue(OpNotEq2(clientOrderId , MkUndefined() )) {
    *(request ).At(MkString("clOrdID") )= OpCopy(clientOrderId );
  } else {
    *(request ).At(MkString("orderID") )= OpCopy(id );
  }
  response := this.Call(method , this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  order := OpCopy(data ); _ = order;
  if IsTrue(GoIsArray(data )) {
    numOrders := OpCopy(data.Length ); _ = numOrders;
    if IsTrue(OpLw(numOrders , MkInteger(1) )) {
      if IsTrue(OpNotEq2(clientOrderId , MkUndefined() )) {
        panic(NewOrderNotFound(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" fetchOrder ") , OpAdd(symbol , OpAdd(MkString(" order with clientOrderId ") , OpAdd(clientOrderId , MkString(" not found") )))))));
      } else {
        panic(NewOrderNotFound(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" fetchOrder ") , OpAdd(symbol , OpAdd(MkString(" order with id ") , OpAdd(id , MkString(" not found") )))))));
      }
    }
    order = this.SafeValue(data , MkInteger(0) , MkMap(&VarMap{}));
  }
  return this.ParseOrder(order , market );
}

func (this *Phemex) FetchOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchOrders() requires a symbol argument") )));
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  method := OpTrinary(*(market ).At(MkString("spot") ), MkString("privateGetSpotOrders") , MkString("privateGetExchangeOrderList") ); _ = method;
  request := MkMap(&VarMap{"symbol":*(market ).At(MkString("id") )}); _ = request;
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("start") )= OpCopy(since );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(method , this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  rows := this.SafeValue(data , MkString("rows") , MkArray(&VarArray{})); _ = rows;
  return this.ParseOrders(rows , market , since , limit );
}

func (this *Phemex) FetchOpenOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchOpenOrders() requires a symbol argument") )));
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  method := OpTrinary(*(market ).At(MkString("spot") ), MkString("privateGetSpotOrders") , MkString("privateGetOrdersActiveList") ); _ = method;
  request := MkMap(&VarMap{"symbol":*(market ).At(MkString("id") )}); _ = request;
  response := OpCopy(MkUndefined() ); _ = response;
  {
  ret__ := func(this *Phemex) (ret_ *Variant) {
    defer func() {
      if e := recover().(*Variant); e != nil {
        ret_ = func(this *Phemex) *Variant {
          // catch block:
          if IsTrue(OpIsType(e , OrderNotFound )) {
            return MkArray(&VarArray{});
          }
          return nil
        }(this)
      }
    }()
    // try block:
    response = this.Call(method , this.Extend(request , params ));
    return nil
  }(this)
  if ret__ != nil {
    return ret__;
   }
  }
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  if IsTrue(GoIsArray(data )) {
    return this.ParseOrders(data , market , since , limit );
  } else {
    rows := this.SafeValue(data , MkString("rows") , MkArray(&VarArray{})); _ = rows;
    return this.ParseOrders(rows , market , since , limit );
  }
  return MkUndefined()
}

func (this *Phemex) FetchClosedOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchClosedOrders() requires a symbol argument") )));
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  method := OpTrinary(*(market ).At(MkString("spot") ), MkString("privateGetExchangeSpotOrder") , MkString("privateGetExchangeOrderList") ); _ = method;
  request := MkMap(&VarMap{"symbol":*(market ).At(MkString("id") )}); _ = request;
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("start") )= OpCopy(since );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(method , this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  if IsTrue(GoIsArray(data )) {
    return this.ParseOrders(data , market , since , limit );
  } else {
    rows := this.SafeValue(data , MkString("rows") , MkArray(&VarArray{})); _ = rows;
    return this.ParseOrders(rows , market , since , limit );
  }
  return MkUndefined()
}

func (this *Phemex) FetchMyTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchClosedOrders() requires a symbol argument") )));
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  method := OpTrinary(*(market ).At(MkString("spot") ), MkString("privateGetExchangeSpotOrderTrades") , MkString("privateGetExchangeOrderTrade") ); _ = method;
  request := MkMap(&VarMap{"symbol":*(market ).At(MkString("id") )}); _ = request;
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("start") )= OpCopy(since );
  }
  if IsTrue(OpAnd(*(market ).At(MkString("swap") ), OpNotEq2(limit , MkUndefined() ))) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(method , this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  rows := this.SafeValue(data , MkString("rows") , MkArray(&VarArray{})); _ = rows;
  return this.ParseTrades(rows , market , since , limit );
}

func (this *Phemex) FetchDepositAddress(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  currency := this.Currency(code ); _ = currency;
  request := MkMap(&VarMap{"currency":*(currency ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("privateGetPhemexUserWalletsV2DepositAddress"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  address := this.SafeString(data , MkString("address") ); _ = address;
  tag := this.SafeString(data , MkString("tag") ); _ = tag;
  this.CheckAddress(address );
  return MkMap(&VarMap{
        "currency":code ,
        "address":address ,
        "tag":tag ,
        "info":response ,
        });
}

func (this *Phemex) FetchDeposits(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  currency := OpCopy(MkUndefined() ); _ = currency;
  if IsTrue(OpNotEq2(code , MkUndefined() )) {
    currency = this.Currency(code );
  }
  response := this.Call(MkString("privateGetExchangeWalletsDepositList"), params ); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  return this.ParseTransactions(data , currency , since , limit );
}

func (this *Phemex) FetchWithdrawals(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  currency := OpCopy(MkUndefined() ); _ = currency;
  if IsTrue(OpNotEq2(code , MkUndefined() )) {
    currency = this.Currency(code );
  }
  response := this.Call(MkString("privateGetExchangeWalletsWithdrawList"), params ); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  return this.ParseTransactions(data , currency , since , limit );
}

func (this *Phemex) ParseTransactionStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "Success":MkString("ok") ,
        "Succeed":MkString("ok") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Phemex) ParseTransaction(goArgs ...*Variant) *Variant{
  transaction := GoGetArg(goArgs, 0, MkUndefined()); _ = transaction;
  currency := GoGetArg(goArgs, 1, MkUndefined() ); _ = currency;
  id := this.SafeString(transaction , MkString("id") ); _ = id;
  address := this.SafeString(transaction , MkString("address") ); _ = address;
  tag := OpCopy(MkUndefined() ); _ = tag;
  txid := this.SafeString(transaction , MkString("txHash") ); _ = txid;
  currencyId := this.SafeString(transaction , MkString("currency") ); _ = currencyId;
  currency = this.SafeCurrency(currencyId , currency );
  code := *(currency ).At(MkString("code") ); _ = code;
  timestamp := this.SafeInteger2(transaction , MkString("createdAt") , MkString("submitedAt") ); _ = timestamp;
  type_ := this.SafeStringLower(transaction , MkString("type") ); _ = type_;
  feeCost := this.ParseNumber(this.FromEn(this.SafeString(transaction , MkString("feeEv") ), *(currency ).At(MkString("valueScale") ))); _ = feeCost;
  fee := OpCopy(MkUndefined() ); _ = fee;
  if IsTrue(OpNotEq2(feeCost , MkUndefined() )) {
    type_ = MkString("withdrawal") ;
    fee = MkMap(&VarMap{
        "cost":feeCost ,
        "currency":code ,
        });
  }
  status := this.ParseTransactionStatus(this.SafeString(transaction , MkString("status") )); _ = status;
  amount := this.ParseNumber(this.FromEn(this.SafeString(transaction , MkString("amountEv") ), *(currency ).At(MkString("valueScale") ))); _ = amount;
  return MkMap(&VarMap{
        "info":transaction ,
        "id":id ,
        "txid":txid ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "address":address ,
        "addressTo":address ,
        "addressFrom":MkUndefined() ,
        "tag":tag ,
        "tagTo":tag ,
        "tagFrom":MkUndefined() ,
        "type":type_ ,
        "amount":amount ,
        "currency":code ,
        "status":status ,
        "updated":MkUndefined() ,
        "fee":fee ,
        });
}

func (this *Phemex) FetchPositions(goArgs ...*Variant) *Variant{
  symbols := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbols;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  code := this.SafeString(params , MkString("code") ); _ = code;
  request := MkMap(&VarMap{}); _ = request;
  if IsTrue(OpEq2(code , MkUndefined() )) {
    currencyId := this.SafeString(params , MkString("currency") ); _ = currencyId;
    if IsTrue(OpEq2(currencyId , MkUndefined() )) {
      panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchPositions() requires a currency parameter or a code parameter") )));
    }
  } else {
    currency := this.Currency(code ); _ = currency;
    params = this.Omit(params , MkString("code") );
    *(request ).At(MkString("currency") )= *(currency ).At(MkString("id") );
  }
  response := this.Call(MkString("privateGetAccountsAccountPositions"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  positions := this.SafeValue(data , MkString("positions") , MkArray(&VarArray{})); _ = positions;
  return positions ;
}

func (this *Phemex) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("public") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  query := this.Omit(params , this.ExtractParams(path )); _ = query;
  requestPath := OpAdd(MkString("/") , this.ImplodeParams(path , params )); _ = requestPath;
  url := OpCopy(requestPath ); _ = url;
  queryString := MkString("") ; _ = queryString;
  if IsTrue(OpOr(OpEq2(method , MkString("GET") ), OpOr(OpEq2(method , MkString("DELETE") ), OpEq2(method , MkString("PUT") )))) {
    if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
      queryString = this.UrlencodeWithArrayRepeat(query );
      OpAddAssign(& url , OpAdd(MkString("?") , queryString ));
    }
  }
  if IsTrue(OpEq2(api , MkString("private") )) {
    this.CheckRequiredCredentials();
    timestamp := this.Seconds(); _ = timestamp;
    xPhemexRequestExpiry := this.SafeInteger(*this.At(MkString("options")) , MkString("x-phemex-request-expiry") , MkInteger(60) ); _ = xPhemexRequestExpiry;
    expiry := this.Sum(timestamp , xPhemexRequestExpiry ); _ = expiry;
    expiryString := expiry.ToString(); _ = expiryString;
    headers = MkMap(&VarMap{
        "x-phemex-access-token":*this.At(MkString("apiKey")) ,
        "x-phemex-request-expiry":expiryString ,
        });
    payload := MkString("") ; _ = payload;
    if IsTrue(OpEq2(method , MkString("POST") )) {
      payload = this.Json(params );
      body = OpCopy(payload );
      *(headers ).At(MkString("Content-Type") )= MkString("application/json") ;
    }
    auth := OpAdd(requestPath , OpAdd(queryString , OpAdd(expiryString , payload ))); _ = auth;
    *(headers ).At(MkString("x-phemex-request-signature") )= this.Hmac(this.Encode(auth ), this.Encode(*this.At(MkString("secret")) ));
  }
  url = OpAdd(this.ImplodeHostname(*(*(*this.At(MkString("urls")) ).At(MkString("api") )).At(api )), url );
  return MkMap(&VarMap{
        "url":url ,
        "method":method ,
        "body":body ,
        "headers":headers ,
        });
}

func (this *Phemex) HandleErrors(goArgs ...*Variant) *Variant{
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
  error := this.SafeValue(response , MkString("error") , response ); _ = error;
  errorCode := this.SafeString(error , MkString("code") ); _ = errorCode;
  message := this.SafeString(error , MkString("msg") ); _ = message;
  if IsTrue(OpAnd(OpNotEq2(errorCode , MkUndefined() ), OpNotEq2(errorCode , MkString("0") ))) {
    feedback := OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , body )); _ = feedback;
    this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("exact") ), errorCode , feedback );
    this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("broad") ), message , feedback );
    panic(NewExchangeError(feedback ));
  }
  return MkUndefined()
}

