package ccxt

import (
)

type Liquid struct {
	*ExchangeBase
}

var _ Exchange = (*Liquid)(nil)

func init() {
	exchange := &Liquid{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Liquid) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("liquid") ,
            "name":MkString("Liquid") ,
            "countries":MkArray(&VarArray{
                MkString("JP") ,
                MkString("CN") ,
                MkString("TW") ,
                }),
            "version":MkString("2") ,
            "rateLimit":MkInteger(1000) ,
            "has":MkMap(&VarMap{
                "cancelOrder":MkBool(true) ,
                "CORS":MkBool(false) ,
                "createOrder":MkBool(true) ,
                "editOrder":MkBool(true) ,
                "fetchBalance":MkBool(true) ,
                "fetchClosedOrders":MkBool(true) ,
                "fetchCurrencies":MkBool(true) ,
                "fetchMarkets":MkBool(true) ,
                "fetchMyTrades":MkBool(true) ,
                "fetchOpenOrders":MkBool(true) ,
                "fetchOrder":MkBool(true) ,
                "fetchOrderBook":MkBool(true) ,
                "fetchOrders":MkBool(true) ,
                "fetchTicker":MkBool(true) ,
                "fetchTickers":MkBool(true) ,
                "fetchTrades":MkBool(true) ,
                "fetchWithdrawals":MkBool(true) ,
                "withdraw":MkBool(true) ,
                }),
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/1294454/45798859-1a872600-bcb4-11e8-8746-69291ce87b04.jpg") ,
                "api":MkString("https://api.liquid.com") ,
                "www":MkString("https://www.liquid.com") ,
                "doc":MkArray(&VarArray{
                    MkString("https://developers.liquid.com") ,
                    }),
                "fees":MkString("https://help.liquid.com/getting-started-with-liquid/the-platform/fee-structure") ,
                "referral":MkString("https://www.liquid.com/sign-up/?affiliate=SbzC62lt30976") ,
                }),
            "api":MkMap(&VarMap{
                "public":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("currencies") ,
                        MkString("products") ,
                        MkString("products/{id}") ,
                        MkString("products/{id}/price_levels") ,
                        MkString("executions") ,
                        MkString("ir_ladders/{currency}") ,
                        MkString("fees") ,
                        })}),
                "private":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("accounts") ,
                        MkString("accounts/balance") ,
                        MkString("accounts/main_asset") ,
                        MkString("accounts/{id}") ,
                        MkString("accounts/{currency}/reserved_balance_details") ,
                        MkString("crypto_accounts") ,
                        MkString("crypto_withdrawals") ,
                        MkString("executions/me") ,
                        MkString("fiat_accounts") ,
                        MkString("fund_infos") ,
                        MkString("loan_bids") ,
                        MkString("loans") ,
                        MkString("orders") ,
                        MkString("orders/{id}") ,
                        MkString("orders/{id}/trades") ,
                        MkString("trades") ,
                        MkString("trades/{id}/loans") ,
                        MkString("trading_accounts") ,
                        MkString("trading_accounts/{id}") ,
                        MkString("transactions") ,
                        MkString("withdrawals") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("crypto_withdrawals") ,
                        MkString("fund_infos") ,
                        MkString("fiat_accounts") ,
                        MkString("loan_bids") ,
                        MkString("orders") ,
                        MkString("withdrawals") ,
                        }),
                    "put":MkArray(&VarArray{
                        MkString("crypto_withdrawal/{id}/cancel") ,
                        MkString("loan_bids/{id}/close") ,
                        MkString("loans/{id}") ,
                        MkString("orders/{id}") ,
                        MkString("orders/{id}/cancel") ,
                        MkString("trades/{id}") ,
                        MkString("trades/{id}/adjust_margin") ,
                        MkString("trades/{id}/close") ,
                        MkString("trades/close_all") ,
                        MkString("trading_accounts/{id}") ,
                        MkString("withdrawals/{id}/cancel") ,
                        }),
                    }),
                }),
            "fees":MkMap(&VarMap{"trading":MkMap(&VarMap{
                    "tierBased":MkBool(true) ,
                    "percentage":MkBool(true) ,
                    "taker":MkNumber(0.0030) ,
                    "maker":MkNumber(0.0000) ,
                    "tiers":MkMap(&VarMap{
                        "perpetual":MkMap(&VarMap{
                            "maker":MkArray(&VarArray{
                                MkArray(&VarArray{
                                    MkInteger(0) ,
                                    MkNumber(0.0000) ,
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(25000) ,
                                    MkNumber(0.0000) ,
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(50000) ,
                                    OpNeg(MkNumber(0.00025) ),
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(100000) ,
                                    OpNeg(MkNumber(0.00025) ),
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(1000000) ,
                                    OpNeg(MkNumber(0.00025) ),
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(10000000) ,
                                    OpNeg(MkNumber(0.00025) ),
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(25000000) ,
                                    OpNeg(MkNumber(0.00025) ),
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(50000000) ,
                                    OpNeg(MkNumber(0.00025) ),
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(75000000) ,
                                    OpNeg(MkNumber(0.00025) ),
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(100000000) ,
                                    OpNeg(MkNumber(0.00025) ),
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(200000000) ,
                                    OpNeg(MkNumber(0.00025) ),
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(300000000) ,
                                    OpNeg(MkNumber(0.00025) ),
                                    }),
                                }),
                            "taker":MkArray(&VarArray{
                                MkArray(&VarArray{
                                    MkInteger(0) ,
                                    MkNumber(0.00120) ,
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(25000) ,
                                    MkNumber(0.00115) ,
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(50000) ,
                                    MkNumber(0.00110) ,
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(100000) ,
                                    MkNumber(0.00105) ,
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(1000000) ,
                                    MkNumber(0.00100) ,
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(10000000) ,
                                    MkNumber(0.00095) ,
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(25000000) ,
                                    MkNumber(0.00090) ,
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(50000000) ,
                                    MkNumber(0.00085) ,
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(75000000) ,
                                    MkNumber(0.00080) ,
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(100000000) ,
                                    MkNumber(0.00075) ,
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(200000000) ,
                                    MkNumber(0.00070) ,
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(300000000) ,
                                    MkNumber(0.00065) ,
                                    }),
                                }),
                            }),
                        "spot":MkMap(&VarMap{
                            "taker":MkArray(&VarArray{
                                MkArray(&VarArray{
                                    MkInteger(0) ,
                                    MkNumber(0.003) ,
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(10000) ,
                                    MkNumber(0.0029) ,
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(20000) ,
                                    MkNumber(0.0028) ,
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(50000) ,
                                    MkNumber(0.0026) ,
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(100000) ,
                                    MkNumber(0.0020) ,
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(1000000) ,
                                    MkNumber(0.0016) ,
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(5000000) ,
                                    MkNumber(0.0012) ,
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(10000000) ,
                                    MkNumber(0.0010) ,
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(25000000) ,
                                    MkNumber(0.0009) ,
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(50000000) ,
                                    MkNumber(0.0008) ,
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(100000000) ,
                                    MkNumber(0.0007) ,
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(200000000) ,
                                    MkNumber(0.0006) ,
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(500000000) ,
                                    MkNumber(0.0004) ,
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(1000000000) ,
                                    MkNumber(0.0003) ,
                                    }),
                                }),
                            "maker":MkArray(&VarArray{
                                MkArray(&VarArray{
                                    MkInteger(0) ,
                                    MkNumber(0.0000) ,
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(10000) ,
                                    MkNumber(0.0020) ,
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(20000) ,
                                    MkNumber(0.0019) ,
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(50000) ,
                                    MkNumber(0.0018) ,
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(100000) ,
                                    MkNumber(0.0016) ,
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(1000000) ,
                                    MkNumber(0.0008) ,
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(5000000) ,
                                    MkNumber(0.0007) ,
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(10000000) ,
                                    MkNumber(0.0005) ,
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(25000000) ,
                                    MkNumber(0.0000) ,
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(50000000) ,
                                    MkNumber(0.0000) ,
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(100000000) ,
                                    MkNumber(0.0000) ,
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(200000000) ,
                                    MkNumber(0.0000) ,
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(500000000) ,
                                    MkNumber(0.0000) ,
                                    }),
                                MkArray(&VarArray{
                                    MkInteger(1000000000) ,
                                    MkNumber(0.0000) ,
                                    }),
                                }),
                            }),
                        }),
                    })}),
            "precisionMode":TICK_SIZE ,
            "exceptions":MkMap(&VarMap{
                "API rate limit exceeded. Please retry after 300s":DDoSProtection ,
                "API Authentication failed":AuthenticationError ,
                "Nonce is too small":InvalidNonce ,
                "Order not found":OrderNotFound ,
                "Can not update partially filled order":InvalidOrder ,
                "Can not update non-live order":OrderNotFound ,
                "not_enough_free_balance":InsufficientFunds ,
                "must_be_positive":InvalidOrder ,
                "less_than_order_size":InvalidOrder ,
                "price_too_high":InvalidOrder ,
                "price_too_small":InvalidOrder ,
                "product_disabled":BadSymbol ,
                }),
            "commonCurrencies":MkMap(&VarMap{
                "WIN":MkString("WCOIN") ,
                "HOT":MkString("HOT Token") ,
                "MIOTA":MkString("IOTA") ,
                }),
            "options":MkMap(&VarMap{"cancelOrderException":MkBool(true) }),
            }));
}

func (this *Liquid) FetchCurrencies(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetCurrencies"), params ); _ = response;
  result := MkMap(&VarMap{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , response.Length )); OpIncr(& i ){
    currency := *(response ).At(i ); _ = currency;
    id := this.SafeString(currency , MkString("currency") ); _ = id;
    code := this.SafeCurrencyCode(id ); _ = code;
    name := this.SafeString(currency , MkString("name") ); _ = name;
    active := OpAnd(*(currency ).At(MkString("depositable") ), *(currency ).At(MkString("withdrawable") )); _ = active;
    amountPrecision := this.SafeInteger(currency , MkString("assets_precision") ); _ = amountPrecision;
    *(result ).At(code )= MkMap(&VarMap{
        "id":id ,
        "code":code ,
        "info":currency ,
        "name":name ,
        "active":active ,
        "fee":this.SafeNumber(currency , MkString("withdrawal_fee") ),
        "precision":amountPrecision ,
        "limits":MkMap(&VarMap{
            "amount":MkMap(&VarMap{
                "min":Math.Pow(MkInteger(10) , OpNeg(amountPrecision )),
                "max":Math.Pow(MkInteger(10) , amountPrecision ),
                }),
            "withdraw":MkMap(&VarMap{
                "min":this.SafeNumber(currency , MkString("minimum_withdrawal") ),
                "max":MkUndefined() ,
                }),
            }),
        });
  }
  return result ;
}

func (this *Liquid) FetchMarkets(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  spot := this.Call(MkString("publicGetProducts"), params ); _ = spot;
  perpetual := this.Call(MkString("publicGetProducts"), MkMap(&VarMap{"perpetual":MkString("1") })); _ = perpetual;
  currencies := this.FetchCurrencies(); _ = currencies;
  currenciesByCode := this.IndexBy(currencies , MkString("code") ); _ = currenciesByCode;
  result := MkArray(&VarArray{}); _ = result;
  markets := this.ArrayConcat(spot , perpetual ); _ = markets;
  for i := MkInteger(0) ; IsTrue(OpLw(i , markets.Length )); OpIncr(& i ){
    market := *(markets ).At(i ); _ = market;
    id := this.SafeString(market , MkString("id") ); _ = id;
    baseId := this.SafeString(market , MkString("base_currency") ); _ = baseId;
    quoteId := this.SafeString(market , MkString("quoted_currency") ); _ = quoteId;
    productType := this.SafeString(market , MkString("product_type") ); _ = productType;
    type_ := MkString("spot") ; _ = type_;
    spot := OpCopy(MkBool(true) ); _ = spot;
    swap := OpCopy(MkBool(false) ); _ = swap;
    if IsTrue(OpEq2(productType , MkString("Perpetual") )) {
      spot = OpCopy(MkBool(false) );
      swap = OpCopy(MkBool(true) );
      type_ = MkString("swap") ;
    }
    base := this.SafeCurrencyCode(baseId ); _ = base;
    quote := this.SafeCurrencyCode(quoteId ); _ = quote;
    symbol := OpCopy(MkUndefined() ); _ = symbol;
    if IsTrue(swap ) {
      symbol = this.SafeString(market , MkString("currency_pair_code") );
    } else {
      symbol = OpAdd(base , OpAdd(MkString("/") , quote ));
    }
    maker := *(*(*this.At(MkString("fees")) ).At(MkString("trading") )).At(MkString("maker") ); _ = maker;
    taker := *(*(*this.At(MkString("fees")) ).At(MkString("trading") )).At(MkString("taker") ); _ = taker;
    if IsTrue(OpEq2(type_ , MkString("swap") )) {
      maker = this.SafeNumber(market , MkString("maker_fee") , *(*(*this.At(MkString("fees")) ).At(MkString("trading") )).At(MkString("maker") ));
      taker = this.SafeNumber(market , MkString("taker_fee") , *(*(*this.At(MkString("fees")) ).At(MkString("trading") )).At(MkString("taker") ));
    }
    disabled := this.SafeValue(market , MkString("disabled") , MkBool(false) ); _ = disabled;
    active := OpNot(disabled ); _ = active;
    baseCurrency := this.SafeValue(currenciesByCode , base ); _ = baseCurrency;
    precision := MkMap(&VarMap{
          "amount":MkNumber(0.00000001) ,
          "price":this.SafeNumber(market , MkString("tick_size") ),
          }); _ = precision;
    minAmount := OpCopy(MkUndefined() ); _ = minAmount;
    if IsTrue(OpNotEq2(baseCurrency , MkUndefined() )) {
      minAmount = this.SafeNumber(*(baseCurrency ).At(MkString("info") ), MkString("minimum_order_quantity") );
    }
    lastPrice := this.SafeNumber(market , MkString("last_traded_price") ); _ = lastPrice;
    minPrice := OpCopy(MkUndefined() ); _ = minPrice;
    maxPrice := OpCopy(MkUndefined() ); _ = maxPrice;
    if IsTrue(lastPrice ) {
      multiplierDown := this.SafeNumber(market , MkString("multiplier_down") ); _ = multiplierDown;
      multiplierUp := this.SafeNumber(market , MkString("multiplier_up") ); _ = multiplierUp;
      if IsTrue(OpNotEq2(multiplierDown , MkUndefined() )) {
        minPrice = OpMulti(lastPrice , multiplierDown );
      }
      if IsTrue(OpNotEq2(multiplierUp , MkUndefined() )) {
        maxPrice = OpMulti(lastPrice , multiplierUp );
      }
    }
    limits := MkMap(&VarMap{
          "amount":MkMap(&VarMap{
              "min":minAmount ,
              "max":MkUndefined() ,
              }),
          "price":MkMap(&VarMap{
              "min":minPrice ,
              "max":maxPrice ,
              }),
          "cost":MkMap(&VarMap{
              "min":MkUndefined() ,
              "max":MkUndefined() ,
              }),
          }); _ = limits;
    result.Push(MkMap(&VarMap{
            "id":id ,
            "symbol":symbol ,
            "base":base ,
            "quote":quote ,
            "baseId":baseId ,
            "quoteId":quoteId ,
            "type":type_ ,
            "spot":spot ,
            "swap":swap ,
            "maker":maker ,
            "taker":taker ,
            "limits":limits ,
            "precision":precision ,
            "active":active ,
            "info":market ,
            }));
  }
  return result ;
}

func (this *Liquid) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("privateGetAccounts"), params ); _ = response;
  result := MkMap(&VarMap{
        "info":response ,
        "timestamp":MkUndefined() ,
        "datetime":MkUndefined() ,
        }); _ = result;
  crypto := this.SafeValue(response , MkString("crypto_accounts") , MkArray(&VarArray{})); _ = crypto;
  fiat := this.SafeValue(response , MkString("fiat_accounts") , MkArray(&VarArray{})); _ = fiat;
  for i := MkInteger(0) ; IsTrue(OpLw(i , crypto.Length )); OpIncr(& i ){
    balance := *(crypto ).At(i ); _ = balance;
    currencyId := this.SafeString(balance , MkString("currency") ); _ = currencyId;
    code := this.SafeCurrencyCode(currencyId ); _ = code;
    account := this.Account(); _ = account;
    *(account ).At(MkString("total") )= this.SafeString(balance , MkString("balance") );
    *(account ).At(MkString("used") )= this.SafeString(balance , MkString("reserved_balance") );
    *(result ).At(code )= OpCopy(account );
  }
  for i := MkInteger(0) ; IsTrue(OpLw(i , fiat.Length )); OpIncr(& i ){
    balance := *(fiat ).At(i ); _ = balance;
    currencyId := this.SafeString(balance , MkString("currency") ); _ = currencyId;
    code := this.SafeCurrencyCode(currencyId ); _ = code;
    account := this.Account(); _ = account;
    *(account ).At(MkString("total") )= this.SafeString(balance , MkString("balance") );
    *(account ).At(MkString("used") )= this.SafeString(balance , MkString("reserved_balance") );
    *(result ).At(code )= OpCopy(account );
  }
  return this.ParseBalance(result );
}

func (this *Liquid) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"id":this.MarketId(symbol )}); _ = request;
  response := this.Call(MkString("publicGetProductsIdPriceLevels"), this.Extend(request , params )); _ = response;
  return this.ParseOrderBook(response , symbol , MkUndefined() , MkString("buy_price_levels") , MkString("sell_price_levels") );
}

func (this *Liquid) ParseTicker(goArgs ...*Variant) *Variant{
  ticker := GoGetArg(goArgs, 0, MkUndefined()); _ = ticker;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.Milliseconds(); _ = timestamp;
  last := OpCopy(MkUndefined() ); _ = last;
  if IsTrue(OpHasMember(MkString("last_traded_price") , ticker )) {
    if IsTrue(*(ticker ).At(MkString("last_traded_price") )) {
      length := *(*(ticker ).At(MkString("last_traded_price") )).At(MkString("length") ); _ = length;
      if IsTrue(OpGt(length , MkInteger(0) )) {
        last = this.SafeNumber(ticker , MkString("last_traded_price") );
      }
    }
  }
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  if IsTrue(OpEq2(market , MkUndefined() )) {
    marketId := this.SafeString(ticker , MkString("id") ); _ = marketId;
    if IsTrue(OpHasMember(marketId , *this.At(MkString("markets_by_id")) )) {
      market = *(*this.At(MkString("markets_by_id")) ).At(marketId );
    } else {
      baseId := this.SafeString(ticker , MkString("base_currency") ); _ = baseId;
      quoteId := this.SafeString(ticker , MkString("quoted_currency") ); _ = quoteId;
      if IsTrue(OpHasMember(symbol , *this.At(MkString("markets")) )) {
        market = *(*this.At(MkString("markets")) ).At(symbol );
      } else {
        symbol = OpAdd(this.SafeCurrencyCode(baseId ), OpAdd(MkString("/") , this.SafeCurrencyCode(quoteId )));
      }
    }
  }
  if IsTrue(OpNotEq2(market , MkUndefined() )) {
    symbol = *(market ).At(MkString("symbol") );
  }
  change := OpCopy(MkUndefined() ); _ = change;
  percentage := OpCopy(MkUndefined() ); _ = percentage;
  average := OpCopy(MkUndefined() ); _ = average;
  open := this.SafeNumber(ticker , MkString("last_price_24h") ); _ = open;
  if IsTrue(OpAnd(OpNotEq2(open , MkUndefined() ), OpNotEq2(last , MkUndefined() ))) {
    change = OpSub(last , open );
    average = OpDiv(this.Sum(last , open ), MkInteger(2) );
    if IsTrue(OpGt(open , MkInteger(0) )) {
      percentage = OpDiv(change , OpMulti(open , MkInteger(100) ));
    }
  }
  return MkMap(&VarMap{
        "symbol":symbol ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "high":this.SafeNumber(ticker , MkString("high_market_ask") ),
        "low":this.SafeNumber(ticker , MkString("low_market_bid") ),
        "bid":this.SafeNumber(ticker , MkString("market_bid") ),
        "bidVolume":MkUndefined() ,
        "ask":this.SafeNumber(ticker , MkString("market_ask") ),
        "askVolume":MkUndefined() ,
        "vwap":MkUndefined() ,
        "open":open ,
        "close":last ,
        "last":last ,
        "previousClose":MkUndefined() ,
        "change":change ,
        "percentage":percentage ,
        "average":average ,
        "baseVolume":this.SafeNumber(ticker , MkString("volume_24h") ),
        "quoteVolume":MkUndefined() ,
        "info":ticker ,
        });
}

func (this *Liquid) FetchTickers(goArgs ...*Variant) *Variant{
  symbols := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbols;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("publicGetProducts"), params ); _ = response;
  result := MkMap(&VarMap{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , response.Length )); OpIncr(& i ){
    ticker := this.ParseTicker(*(response ).At(i )); _ = ticker;
    symbol := *(ticker ).At(MkString("symbol") ); _ = symbol;
    *(result ).At(symbol )= OpCopy(ticker );
  }
  return this.FilterByArray(result , MkString("symbol") , symbols );
}

func (this *Liquid) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"id":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("publicGetProductsId"), this.Extend(request , params )); _ = response;
  return this.ParseTicker(response , market );
}

func (this *Liquid) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.SafeTimestamp(trade , MkString("created_at") ); _ = timestamp;
  orderId := this.SafeString(trade , MkString("order_id") ); _ = orderId;
  takerSide := this.SafeString(trade , MkString("taker_side") ); _ = takerSide;
  mySide := this.SafeString(trade , MkString("my_side") ); _ = mySide;
  side := OpTrinary(OpNotEq2(mySide , MkUndefined() ), mySide , takerSide ); _ = side;
  takerOrMaker := OpCopy(MkUndefined() ); _ = takerOrMaker;
  if IsTrue(OpNotEq2(mySide , MkUndefined() )) {
    takerOrMaker = OpTrinary(OpEq2(takerSide , mySide ), MkString("taker") , MkString("maker") );
  }
  priceString := this.SafeString(trade , MkString("price") ); _ = priceString;
  amountString := this.SafeString(trade , MkString("quantity") ); _ = amountString;
  price := this.ParseNumber(priceString ); _ = price;
  amount := this.ParseNumber(amountString ); _ = amount;
  cost := this.ParseNumber(Precise.StringMul(priceString , amountString )); _ = cost;
  id := this.SafeString(trade , MkString("id") ); _ = id;
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  if IsTrue(OpNotEq2(market , MkUndefined() )) {
    symbol = *(market ).At(MkString("symbol") );
  }
  return MkMap(&VarMap{
        "info":trade ,
        "id":id ,
        "order":orderId ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "symbol":symbol ,
        "type":MkUndefined() ,
        "side":side ,
        "takerOrMaker":takerOrMaker ,
        "price":price ,
        "amount":amount ,
        "cost":cost ,
        "fee":MkUndefined() ,
        });
}

func (this *Liquid) FetchTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"product_id":*(market ).At(MkString("id") )}); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("timestamp") )= parseInt(OpDiv(since , MkInteger(1000) ));
  }
  response := this.Call(MkString("publicGetExecutions"), this.Extend(request , params )); _ = response;
  result := OpTrinary(OpNotEq2(since , MkUndefined() ), response , *(response ).At(MkString("models") )); _ = result;
  return this.ParseTrades(result , market , since , limit );
}

func (this *Liquid) FetchMyTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "product_id":*(market ).At(MkString("id") ),
        "with_details":MkBool(true) ,
        }); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("privateGetExecutionsMe"), this.Extend(request , params )); _ = response;
  return this.ParseTrades(*(response ).At(MkString("models") ), market , since , limit );
}

func (this *Liquid) CreateOrder(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 2, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 3, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 4, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 5, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  clientOrderId := this.SafeString2(params , MkString("clientOrderId") , MkString("client_order_id") ); _ = clientOrderId;
  params = this.Omit(params , MkArray(&VarArray{
          MkString("clientOrderId") ,
          MkString("client_order_id") ,
          }));
  request := MkMap(&VarMap{
        "order_type":type_ ,
        "product_id":this.MarketId(symbol ),
        "side":side ,
        "quantity":this.AmountToPrecision(symbol , amount ),
        }); _ = request;
  if IsTrue(OpNotEq2(clientOrderId , MkUndefined() )) {
    *(request ).At(MkString("client_order_id") )= OpCopy(clientOrderId );
  }
  if IsTrue(OpOr(OpEq2(type_ , MkString("limit") ), OpOr(OpEq2(type_ , MkString("limit_post_only") ), OpOr(OpEq2(type_ , MkString("market_with_range") ), OpEq2(type_ , MkString("stop") ))))) {
    *(request ).At(MkString("price") )= this.PriceToPrecision(symbol , price );
  }
  response := this.Call(MkString("privatePostOrders"), this.Extend(request , params )); _ = response;
  return this.ParseOrder(response );
}

func (this *Liquid) CancelOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"id":id }); _ = request;
  response := this.Call(MkString("privatePutOrdersIdCancel"), this.Extend(request , params )); _ = response;
  order := this.ParseOrder(response ); _ = order;
  if IsTrue(OpEq2(*(order ).At(MkString("status") ), MkString("closed") )) {
    if IsTrue(*(*this.At(MkString("options")) ).At(MkString("cancelOrderException") )) {
      panic(NewOrderNotFound(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" order closed already: ") , this.Json(response )))));
    }
  }
  return order ;
}

func (this *Liquid) EditOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 2, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 3, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 4, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 5, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 6, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  if IsTrue(OpEq2(price , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" editOrder() requires the price argument") )));
  }
  request := MkMap(&VarMap{
        "order":MkMap(&VarMap{
            "quantity":this.AmountToPrecision(symbol , amount ),
            "price":this.PriceToPrecision(symbol , price ),
            }),
        "id":id ,
        }); _ = request;
  response := this.Call(MkString("privatePutOrdersId"), this.Extend(request , params )); _ = response;
  return this.ParseOrder(response );
}

func (this *Liquid) ParseOrderStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "live":MkString("open") ,
        "filled":MkString("closed") ,
        "cancelled":MkString("canceled") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Liquid) ParseOrder(goArgs ...*Variant) *Variant{
  order := GoGetArg(goArgs, 0, MkUndefined()); _ = order;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  orderId := this.SafeString(order , MkString("id") ); _ = orderId;
  timestamp := this.SafeTimestamp(order , MkString("created_at") ); _ = timestamp;
  marketId := this.SafeString(order , MkString("product_id") ); _ = marketId;
  market = this.SafeValue(*this.At(MkString("markets_by_id")) , marketId );
  status := this.ParseOrderStatus(this.SafeString(order , MkString("status") )); _ = status;
  amount := this.SafeNumber(order , MkString("quantity") ); _ = amount;
  filled := this.SafeNumber(order , MkString("filled_quantity") ); _ = filled;
  price := this.SafeNumber(order , MkString("price") ); _ = price;
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  feeCurrency := OpCopy(MkUndefined() ); _ = feeCurrency;
  if IsTrue(OpNotEq2(market , MkUndefined() )) {
    symbol = *(market ).At(MkString("symbol") );
    feeCurrency = *(market ).At(MkString("quote") );
  }
  type_ := this.SafeString(order , MkString("order_type") ); _ = type_;
  tradeCost := MkInteger(0) ; _ = tradeCost;
  tradeFilled := MkInteger(0) ; _ = tradeFilled;
  average := this.SafeNumber(order , MkString("average_price") ); _ = average;
  trades := this.ParseTrades(this.SafeValue(order , MkString("executions") , MkArray(&VarArray{})), market , MkUndefined() , MkUndefined() , MkMap(&VarMap{
            "order":orderId ,
            "type":type_ ,
            })); _ = trades;
  numTrades := OpCopy(trades.Length ); _ = numTrades;
  for i := MkInteger(0) ; IsTrue(OpLw(i , numTrades )); OpIncr(& i ){
    trade := *(trades ).At(i ); _ = trade;
    *(trade ).At(MkString("order") )= OpCopy(orderId );
    *(trade ).At(MkString("type") )= OpCopy(type_ );
    tradeFilled = this.Sum(tradeFilled , *(trade ).At(MkString("amount") ));
    tradeCost = this.Sum(tradeCost , *(trade ).At(MkString("cost") ));
  }
  cost := OpCopy(MkUndefined() ); _ = cost;
  lastTradeTimestamp := OpCopy(MkUndefined() ); _ = lastTradeTimestamp;
  if IsTrue(OpGt(numTrades , MkInteger(0) )) {
    lastTradeTimestamp = *(*(trades ).At(OpSub(numTrades , MkInteger(1) ))).At(MkString("timestamp") );
    if IsTrue(OpAnd(OpNot(average ), OpGt(tradeFilled , MkInteger(0) ))) {
      average = OpDiv(tradeCost , tradeFilled );
    }
    if IsTrue(OpEq2(cost , MkUndefined() )) {
      cost = OpCopy(tradeCost );
    }
    if IsTrue(OpEq2(filled , MkUndefined() )) {
      filled = OpCopy(tradeFilled );
    }
  }
  remaining := OpCopy(MkUndefined() ); _ = remaining;
  if IsTrue(OpAnd(OpNotEq2(amount , MkUndefined() ), OpNotEq2(filled , MkUndefined() ))) {
    remaining = OpSub(amount , filled );
  }
  side := this.SafeString(order , MkString("side") ); _ = side;
  clientOrderId := this.SafeString(order , MkString("client_order_id") ); _ = clientOrderId;
  return MkMap(&VarMap{
        "id":orderId ,
        "clientOrderId":clientOrderId ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "lastTradeTimestamp":lastTradeTimestamp ,
        "type":type_ ,
        "timeInForce":MkUndefined() ,
        "postOnly":MkUndefined() ,
        "status":status ,
        "symbol":symbol ,
        "side":side ,
        "price":price ,
        "stopPrice":MkUndefined() ,
        "amount":amount ,
        "filled":filled ,
        "cost":cost ,
        "remaining":remaining ,
        "average":average ,
        "trades":trades ,
        "fee":MkMap(&VarMap{
            "currency":feeCurrency ,
            "cost":this.SafeNumber(order , MkString("order_fee") ),
            }),
        "info":order ,
        });
}

func (this *Liquid) FetchOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"id":id }); _ = request;
  response := this.Call(MkString("privateGetOrdersId"), this.Extend(request , params )); _ = response;
  return this.ParseOrder(response );
}

func (this *Liquid) FetchOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := OpCopy(MkUndefined() ); _ = market;
  request := MkMap(&VarMap{"with_details":MkInteger(1) }); _ = request;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("product_id") )= *(market ).At(MkString("id") );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("privateGetOrders"), this.Extend(request , params )); _ = response;
  orders := this.SafeValue(response , MkString("models") , MkArray(&VarArray{})); _ = orders;
  return this.ParseOrders(orders , market , since , limit );
}

func (this *Liquid) FetchOpenOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"status":MkString("live") }); _ = request;
  return this.FetchOrders(symbol , since , limit , this.Extend(request , params ));
}

func (this *Liquid) FetchClosedOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"status":MkString("filled") }); _ = request;
  return this.FetchOrders(symbol , since , limit , this.Extend(request , params ));
}

func (this *Liquid) Withdraw(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  amount := GoGetArg(goArgs, 1, MkUndefined()); _ = amount;
  address := GoGetArg(goArgs, 2, MkUndefined()); _ = address;
  tag := GoGetArg(goArgs, 3, MkUndefined() ); _ = tag;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.CheckAddress(address );
  this.LoadMarkets();
  currency := this.Currency(code ); _ = currency;
  request := MkMap(&VarMap{
        "currency":*(currency ).At(MkString("id") ),
        "address":address ,
        "amount":amount ,
        }); _ = request;
  if IsTrue(OpNotEq2(tag , MkUndefined() )) {
    if IsTrue(OpEq2(code , MkString("XRP") )) {
      *(request ).At(MkString("payment_id") )= OpCopy(tag );
    } else {
      if IsTrue(OpEq2(code , MkString("XLM") )) {
        *(request ).At(MkString("memo_type") )= MkString("text") ;
        *(request ).At(MkString("memo_value") )= OpCopy(tag );
      } else {
        panic(NewNotSupported(OpAdd(*this.At(MkString("id")) , MkString(" withdraw() only supports a tag along the address for XRP or XLM") )));
      }
    }
  }
  response := this.Call(MkString("privatePostCryptoWithdrawals"), this.Extend(request , params )); _ = response;
  return this.ParseTransaction(response , currency );
}

func (this *Liquid) FetchWithdrawals(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  currency := OpCopy(MkUndefined() ); _ = currency;
  if IsTrue(OpNotEq2(code , MkUndefined() )) {
    currency = this.Currency(code );
  }
  response := this.Call(MkString("privateGetCryptoWithdrawals"), this.Extend(request , params )); _ = response;
  transactions := this.SafeValue(response , MkString("models") , MkArray(&VarArray{})); _ = transactions;
  return this.ParseTransactions(transactions , currency , since , limit );
}

func (this *Liquid) ParseTransactionStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "pending":MkString("pending") ,
        "cancelled":MkString("canceled") ,
        "approved":MkString("ok") ,
        "processing":MkString("pending") ,
        "processed":MkString("ok") ,
        "reverted":MkString("failed") ,
        "to_be_reviewed":MkString("pending") ,
        "declined":MkString("failed") ,
        "broadcasted":MkString("ok") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Liquid) ParseTransaction(goArgs ...*Variant) *Variant{
  transaction := GoGetArg(goArgs, 0, MkUndefined()); _ = transaction;
  currency := GoGetArg(goArgs, 1, MkUndefined() ); _ = currency;
  id := this.SafeString(transaction , MkString("id") ); _ = id;
  address := this.SafeString(transaction , MkString("address") ); _ = address;
  tag := this.SafeString2(transaction , MkString("payment_id") , MkString("memo_value") ); _ = tag;
  txid := this.SafeString(transaction , MkString("transaction_hash") ); _ = txid;
  currencyId := this.SafeString2(transaction , MkString("currency") , MkString("asset") ); _ = currencyId;
  code := this.SafeCurrencyCode(currencyId , currency ); _ = code;
  timestamp := this.SafeTimestamp(transaction , MkString("created_at") ); _ = timestamp;
  updated := this.SafeTimestamp(transaction , MkString("updated_at") ); _ = updated;
  type_ := MkString("withdrawal") ; _ = type_;
  status := this.ParseTransactionStatus(this.SafeString(transaction , MkString("state") )); _ = status;
  amountString := this.SafeString(transaction , MkString("amount") ); _ = amountString;
  feeCostString := this.SafeString(transaction , MkString("withdrawal_fee") ); _ = feeCostString;
  amount := this.ParseNumber(Precise.StringSub(amountString , feeCostString )); _ = amount;
  return MkMap(&VarMap{
        "info":transaction ,
        "id":id ,
        "txid":txid ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "address":address ,
        "tag":tag ,
        "type":type_ ,
        "amount":amount ,
        "currency":code ,
        "status":status ,
        "updated":updated ,
        "fee":MkMap(&VarMap{
            "currency":code ,
            "cost":this.ParseNumber(feeCostString ),
            }),
        });
}

func (this *Liquid) Nonce(goArgs ...*Variant) *Variant{
  return this.Milliseconds();
}

func (this *Liquid) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("public") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  url := OpAdd(MkString("/") , this.ImplodeParams(path , params )); _ = url;
  query := this.Omit(params , this.ExtractParams(path )); _ = query;
  headers = MkMap(&VarMap{
      "X-Quoine-API-Version":*this.At(MkString("version")) ,
      "Content-Type":MkString("application/json") ,
      });
  if IsTrue(OpEq2(api , MkString("private") )) {
    this.CheckRequiredCredentials();
    if IsTrue(OpEq2(method , MkString("GET") )) {
      if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
        OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(query )));
      }
    } else {
      if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
        body = this.Json(query );
      }
    }
    nonce := this.Nonce(); _ = nonce;
    request := MkMap(&VarMap{
          "path":url ,
          "token_id":*this.At(MkString("apiKey")) ,
          "iat":Math.Floor(OpDiv(nonce , MkInteger(1000) )),
          }); _ = request;
    if IsTrue(OpNot(OpHasMember(MkString("client_order_id") , query ))) {
      *(request ).At(MkString("nonce") )= OpCopy(nonce );
    }
    *(headers ).At(MkString("X-Quoine-Auth") )= this.Jwt(request , this.Encode(*this.At(MkString("secret")) ));
  } else {
    if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
      OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(query )));
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

func (this *Liquid) HandleErrors(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  reason := GoGetArg(goArgs, 1, MkUndefined()); _ = reason;
  url := GoGetArg(goArgs, 2, MkUndefined()); _ = url;
  method := GoGetArg(goArgs, 3, MkUndefined()); _ = method;
  headers := GoGetArg(goArgs, 4, MkUndefined()); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined()); _ = body;
  response := GoGetArg(goArgs, 6, MkUndefined()); _ = response;
  requestHeaders := GoGetArg(goArgs, 7, MkUndefined()); _ = requestHeaders;
  requestBody := GoGetArg(goArgs, 8, MkUndefined()); _ = requestBody;
  if IsTrue(OpAnd(OpGtEq(code , MkInteger(200) ), OpLw(code , MkInteger(300) ))) {
    return MkUndefined();
  }
  if IsTrue(OpEq2(code , MkInteger(401) )) {
    this.ThrowExactlyMatchedException(*this.At(MkString("exceptions")) , body , body );
    return MkUndefined();
  }
  if IsTrue(OpEq2(code , MkInteger(429) )) {
    panic(NewDDoSProtection(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , body ))));
  }
  if IsTrue(OpEq2(response , MkUndefined() )) {
    return MkUndefined();
  }
  feedback := OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , body )); _ = feedback;
  message := this.SafeString(response , MkString("message") ); _ = message;
  errors := this.SafeValue(response , MkString("errors") ); _ = errors;
  if IsTrue(OpNotEq2(message , MkUndefined() )) {
    this.ThrowExactlyMatchedException(*this.At(MkString("exceptions")) , message , feedback );
  } else {
    if IsTrue(OpNotEq2(errors , MkUndefined() )) {
      types := GoGetKeys(errors ); _ = types;
      for i := MkInteger(0) ; IsTrue(OpLw(i , types.Length )); OpIncr(& i ){
        type_ := *(types ).At(i ); _ = type_;
        errorMessages := *(errors ).At(type_ ); _ = errorMessages;
        for j := MkInteger(0) ; IsTrue(OpLw(j , errorMessages.Length )); OpIncr(& j ){
          message := *(errorMessages ).At(j ); _ = message;
          this.ThrowExactlyMatchedException(*this.At(MkString("exceptions")) , message , feedback );
        }
      }
    } else {
      panic(NewExchangeError(feedback ));
    }
  }
  return MkUndefined()
}

