package ccxt

import (
)

type Digifinex struct {
	*ExchangeBase
}

var _ Exchange = (*Digifinex)(nil)

func init() {
	exchange := &Digifinex{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Digifinex) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("digifinex") ,
            "name":MkString("DigiFinex") ,
            "countries":MkArray(&VarArray{
                MkString("SG") ,
                }),
            "version":MkString("v3") ,
            "rateLimit":MkInteger(900) ,
            "has":MkMap(&VarMap{
                "cancelOrder":MkBool(true) ,
                "cancelOrders":MkBool(true) ,
                "createOrder":MkBool(true) ,
                "fetchBalance":MkBool(true) ,
                "fetchCurrencies":MkBool(true) ,
                "fetchDepositAddress":MkBool(true) ,
                "fetchDeposits":MkBool(true) ,
                "fetchLedger":MkBool(true) ,
                "fetchMarkets":MkBool(true) ,
                "fetchMyTrades":MkBool(true) ,
                "fetchOHLCV":MkBool(true) ,
                "fetchOpenOrders":MkBool(true) ,
                "fetchOrder":MkBool(true) ,
                "fetchOrderBook":MkBool(true) ,
                "fetchOrders":MkBool(true) ,
                "fetchStatus":MkBool(true) ,
                "fetchTicker":MkBool(true) ,
                "fetchTickers":MkBool(true) ,
                "fetchTime":MkBool(true) ,
                "fetchTrades":MkBool(true) ,
                "fetchWithdrawals":MkBool(true) ,
                "withdraw":MkBool(true) ,
                }),
            "timeframes":MkMap(&VarMap{
                "1m":MkString("1") ,
                "5m":MkString("5") ,
                "15m":MkString("15") ,
                "30m":MkString("30") ,
                "1h":MkString("60") ,
                "4h":MkString("240") ,
                "12h":MkString("720") ,
                "1d":MkString("1D") ,
                "1w":MkString("1W") ,
                }),
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/51840849/87443315-01283a00-c5fe-11ea-8628-c2a0feaf07ac.jpg") ,
                "api":MkString("https://openapi.digifinex.com") ,
                "www":MkString("https://www.digifinex.com") ,
                "doc":MkArray(&VarArray{
                    MkString("https://docs.digifinex.com") ,
                    }),
                "fees":MkString("https://digifinex.zendesk.com/hc/en-us/articles/360000328422-Fee-Structure-on-DigiFinex") ,
                "referral":MkString("https://www.digifinex.com/en-ww/from/DhOzBg?channelCode=ljaUPp") ,
                }),
            "api":MkMap(&VarMap{
                "public":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("{market}/symbols") ,
                        MkString("kline") ,
                        MkString("margin/currencies") ,
                        MkString("margin/symbols") ,
                        MkString("markets") ,
                        MkString("order_book") ,
                        MkString("ping") ,
                        MkString("spot/symbols") ,
                        MkString("time") ,
                        MkString("trades") ,
                        MkString("trades/symbols") ,
                        MkString("ticker") ,
                        MkString("currencies") ,
                        })}),
                "private":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("{market}/financelog") ,
                        MkString("{market}/mytrades") ,
                        MkString("{market}/order") ,
                        MkString("{market}â/orderâ/detail") ,
                        MkString("{market}/order/current") ,
                        MkString("{market}/order/history") ,
                        MkString("margin/assets") ,
                        MkString("margin/financelog") ,
                        MkString("margin/mytrades") ,
                        MkString("margin/order") ,
                        MkString("margin/order/current") ,
                        MkString("margin/order/history") ,
                        MkString("margin/positions") ,
                        MkString("otc/financelog") ,
                        MkString("spot/assets") ,
                        MkString("spot/financelog") ,
                        MkString("spot/mytrades") ,
                        MkString("spot/order") ,
                        MkString("spot/order/current") ,
                        MkString("spot/order/history") ,
                        MkString("deposit/address") ,
                        MkString("deposit/history") ,
                        MkString("withdraw/history") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("{market}/order/cancel") ,
                        MkString("{market}/order/new") ,
                        MkString("{market}â/orderâ/batch_new") ,
                        MkString("margin/order/cancel") ,
                        MkString("margin/order/new") ,
                        MkString("margin/position/close") ,
                        MkString("spot/order/cancel") ,
                        MkString("spot/order/new") ,
                        MkString("transfer") ,
                        MkString("withdraw/new") ,
                        MkString("withdraw/cancel") ,
                        }),
                    }),
                }),
            "fees":MkMap(&VarMap{"trading":MkMap(&VarMap{
                    "tierBased":MkBool(false) ,
                    "percentage":MkBool(true) ,
                    "maker":this.ParseNumber(MkString("0.002") ),
                    "taker":this.ParseNumber(MkString("0.002") ),
                    })}),
            "exceptions":MkMap(&VarMap{
                "exact":MkMap(&VarMap{
                    "10001":MkArray(&VarArray{
                        BadRequest ,
                        MkString("Wrong request method, please check it's a GET ot POST request") ,
                        }),
                    "10002":MkArray(&VarArray{
                        AuthenticationError ,
                        MkString("Invalid ApiKey") ,
                        }),
                    "10003":MkArray(&VarArray{
                        AuthenticationError ,
                        MkString("Sign doesn't match") ,
                        }),
                    "10004":MkArray(&VarArray{
                        BadRequest ,
                        MkString("Illegal request parameters") ,
                        }),
                    "10005":MkArray(&VarArray{
                        DDoSProtection ,
                        MkString("Request frequency exceeds the limit") ,
                        }),
                    "10006":MkArray(&VarArray{
                        PermissionDenied ,
                        MkString("Unauthorized to execute this request") ,
                        }),
                    "10007":MkArray(&VarArray{
                        PermissionDenied ,
                        MkString("IP address Unauthorized") ,
                        }),
                    "10008":MkArray(&VarArray{
                        InvalidNonce ,
                        MkString("Timestamp for this request is invalid, timestamp must within 1 minute") ,
                        }),
                    "10009":MkArray(&VarArray{
                        NetworkError ,
                        MkString("Unexist endpoint, please check endpoint URL") ,
                        }),
                    "10011":MkArray(&VarArray{
                        AccountSuspended ,
                        MkString("ApiKey expired. Please go to client side to re-create an ApiKey") ,
                        }),
                    "20001":MkArray(&VarArray{
                        PermissionDenied ,
                        MkString("Trade is not open for this trading pair") ,
                        }),
                    "20002":MkArray(&VarArray{
                        PermissionDenied ,
                        MkString("Trade of this trading pair is suspended") ,
                        }),
                    "20003":MkArray(&VarArray{
                        InvalidOrder ,
                        MkString("Invalid price or amount") ,
                        }),
                    "20007":MkArray(&VarArray{
                        InvalidOrder ,
                        MkString("Price precision error") ,
                        }),
                    "20008":MkArray(&VarArray{
                        InvalidOrder ,
                        MkString("Amount precision error") ,
                        }),
                    "20009":MkArray(&VarArray{
                        InvalidOrder ,
                        MkString("Amount is less than the minimum requirement") ,
                        }),
                    "20010":MkArray(&VarArray{
                        InvalidOrder ,
                        MkString("Cash Amount is less than the minimum requirement") ,
                        }),
                    "20011":MkArray(&VarArray{
                        InsufficientFunds ,
                        MkString("Insufficient balance") ,
                        }),
                    "20012":MkArray(&VarArray{
                        BadRequest ,
                        MkString("Invalid trade type, valid value: buy/sell)") ,
                        }),
                    "20013":MkArray(&VarArray{
                        InvalidOrder ,
                        MkString("No order info found") ,
                        }),
                    "20014":MkArray(&VarArray{
                        BadRequest ,
                        MkString("Invalid date, Valid format: 2018-07-25)") ,
                        }),
                    "20015":MkArray(&VarArray{
                        BadRequest ,
                        MkString("Date exceeds the limit") ,
                        }),
                    "20018":MkArray(&VarArray{
                        PermissionDenied ,
                        MkString("Your trading rights have been banned by the system") ,
                        }),
                    "20019":MkArray(&VarArray{
                        BadSymbol ,
                        MkString("Wrong trading pair symbol. Correct format:\"usdt_btc\". Quote asset is in the front") ,
                        }),
                    "20020":MkArray(&VarArray{
                        DDoSProtection ,
                        MkString("You have violated the API operation trading rules and temporarily forbid trading. At present, we have certain restrictions on the user's transaction rate and withdrawal rate.") ,
                        }),
                    "50000":MkArray(&VarArray{
                        ExchangeError ,
                        MkString("Exception error") ,
                        }),
                    "20021":MkArray(&VarArray{
                        BadRequest ,
                        MkString("Invalid currency") ,
                        }),
                    "20022":MkArray(&VarArray{
                        BadRequest ,
                        MkString("The ending timestamp must be larger than the starting timestamp") ,
                        }),
                    "20023":MkArray(&VarArray{
                        BadRequest ,
                        MkString("Invalid transfer type") ,
                        }),
                    "20024":MkArray(&VarArray{
                        BadRequest ,
                        MkString("Invalid amount") ,
                        }),
                    "20025":MkArray(&VarArray{
                        BadRequest ,
                        MkString("This currency is not transferable at the moment") ,
                        }),
                    "20026":MkArray(&VarArray{
                        InsufficientFunds ,
                        MkString("Transfer amount exceed your balance") ,
                        }),
                    "20027":MkArray(&VarArray{
                        PermissionDenied ,
                        MkString("Abnormal account status") ,
                        }),
                    "20028":MkArray(&VarArray{
                        PermissionDenied ,
                        MkString("Blacklist for transfer") ,
                        }),
                    "20029":MkArray(&VarArray{
                        PermissionDenied ,
                        MkString("Transfer amount exceed your daily limit") ,
                        }),
                    "20030":MkArray(&VarArray{
                        BadRequest ,
                        MkString("You have no position on this trading pair") ,
                        }),
                    "20032":MkArray(&VarArray{
                        PermissionDenied ,
                        MkString("Withdrawal limited") ,
                        }),
                    "20033":MkArray(&VarArray{
                        BadRequest ,
                        MkString("Wrong Withdrawal ID") ,
                        }),
                    "20034":MkArray(&VarArray{
                        PermissionDenied ,
                        MkString("Withdrawal service of this crypto has been closed") ,
                        }),
                    "20035":MkArray(&VarArray{
                        PermissionDenied ,
                        MkString("Withdrawal limit") ,
                        }),
                    "20036":MkArray(&VarArray{
                        ExchangeError ,
                        MkString("Withdrawal cancellation failed") ,
                        }),
                    "20037":MkArray(&VarArray{
                        InvalidAddress ,
                        MkString("The withdrawal address, Tag or chain type is not included in the withdrawal management list") ,
                        }),
                    "20038":MkArray(&VarArray{
                        InvalidAddress ,
                        MkString("The withdrawal address is not on the white list") ,
                        }),
                    "20039":MkArray(&VarArray{
                        ExchangeError ,
                        MkString("Can't be canceled in current status") ,
                        }),
                    "20040":MkArray(&VarArray{
                        RateLimitExceeded ,
                        MkString("Withdraw too frequently; limitation: 3 times a minute, 100 times a day") ,
                        }),
                    "20041":MkArray(&VarArray{
                        PermissionDenied ,
                        MkString("Beyond the daily withdrawal limit") ,
                        }),
                    "20042":MkArray(&VarArray{
                        BadSymbol ,
                        MkString("Current trading pair does not support API trading") ,
                        }),
                    }),
                "broad":MkMap(&VarMap{}),
                }),
            "options":MkMap(&VarMap{
                "defaultType":MkString("spot") ,
                "types":MkArray(&VarArray{
                    MkString("spot") ,
                    MkString("margin") ,
                    MkString("otc") ,
                    }),
                }),
            "commonCurrencies":MkMap(&VarMap{
                "BHT":MkString("Black House Test") ,
                "EPS":MkString("Epanus") ,
                "MBN":MkString("Mobilian Coin") ,
                "TEL":MkString("TEL666") ,
                }),
            }));
}

func (this *Digifinex) FetchCurrencies(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetCurrencies"), params ); _ = response;
  data := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = data;
  result := MkMap(&VarMap{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , data.Length )); OpIncr(& i ){
    currency := *(data ).At(i ); _ = currency;
    id := this.SafeString(currency , MkString("currency") ); _ = id;
    code := this.SafeCurrencyCode(id ); _ = code;
    depositStatus := this.SafeValue(currency , MkString("deposit_status") , MkInteger(1) ); _ = depositStatus;
    withdrawStatus := this.SafeValue(currency , MkString("withdraw_status") , MkInteger(1) ); _ = withdrawStatus;
    active := OpAnd(depositStatus , withdrawStatus ); _ = active;
    fee := this.SafeNumber(currency , MkString("withdraw_fee_rate") ); _ = fee;
    if IsTrue(OpHasMember(code , result )) {
      if IsTrue(GoIsArray(*(*(result ).At(code )).At(MkString("info") ))) {
        (*(*(result ).At(code )).At(MkString("info") )).Call(MkString("push") , currency );
      } else {
        *(*(result ).At(code )).At(MkString("info") )= MkArray(&VarArray{
            *(*(result ).At(code )).At(MkString("info") ),
            currency ,
            });
      }
    } else {
      *(result ).At(code )= MkMap(&VarMap{
          "id":id ,
          "code":code ,
          "info":currency ,
          "type":MkUndefined() ,
          "name":MkUndefined() ,
          "active":active ,
          "fee":fee ,
          "precision":MkInteger(8) ,
          "limits":MkMap(&VarMap{
              "amount":MkMap(&VarMap{
                  "min":MkUndefined() ,
                  "max":MkUndefined() ,
                  }),
              "withdraw":MkMap(&VarMap{
                  "min":this.SafeNumber(currency , MkString("min_withdraw_amount") ),
                  "max":MkUndefined() ,
                  }),
              }),
          });
    }
  }
  return result ;
}

func (this *Digifinex) FetchMarkets(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  options := this.SafeValue(*this.At(MkString("options")) , MkString("fetchMarkets") , MkMap(&VarMap{})); _ = options;
  method := this.SafeString(options , MkString("method") , MkString("fetch_markets_v2") ); _ = method;
  return this.Call(method , params );
}

func (this *Digifinex) FetchMarketsV2(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetTradesSymbols"), params ); _ = response;
  markets := this.SafeValue(response , MkString("symbol_list") , MkArray(&VarArray{})); _ = markets;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , markets.Length )); OpIncr(& i ){
    market := *(markets ).At(i ); _ = market;
    id := this.SafeString(market , MkString("symbol") ); _ = id;
    baseId := this.SafeString(market , MkString("base_asset") ); _ = baseId;
    quoteId := this.SafeString(market , MkString("quote_asset") ); _ = quoteId;
    base := this.SafeCurrencyCode(baseId ); _ = base;
    quote := this.SafeCurrencyCode(quoteId ); _ = quote;
    symbol := OpAdd(base , OpAdd(MkString("/") , quote )); _ = symbol;
    precision := MkMap(&VarMap{
          "amount":this.SafeInteger(market , MkString("amount_precision") ),
          "price":this.SafeInteger(market , MkString("price_precision") ),
          }); _ = precision;
    limits := MkMap(&VarMap{
          "amount":MkMap(&VarMap{
              "min":this.SafeNumber(market , MkString("minimum_amount") ),
              "max":MkUndefined() ,
              }),
          "price":MkMap(&VarMap{
              "min":MkUndefined() ,
              "max":MkUndefined() ,
              }),
          "cost":MkMap(&VarMap{
              "min":this.SafeNumber(market , MkString("minimum_value") ),
              "max":MkUndefined() ,
              }),
          }); _ = limits;
    isAllowed := this.SafeInteger(market , MkString("is_allow") , MkInteger(1) ); _ = isAllowed;
    active := OpTrinary(isAllowed , MkBool(true) , MkBool(false) ); _ = active;
    type_ := MkString("spot") ; _ = type_;
    spot := OpEq2(type_ , MkString("spot") ); _ = spot;
    margin := OpEq2(type_ , MkString("margin") ); _ = margin;
    result.Push(MkMap(&VarMap{
            "id":id ,
            "symbol":symbol ,
            "base":base ,
            "quote":quote ,
            "baseId":baseId ,
            "quoteId":quoteId ,
            "active":active ,
            "type":type_ ,
            "spot":spot ,
            "margin":margin ,
            "precision":precision ,
            "limits":limits ,
            "info":market ,
            }));
  }
  return result ;
}

func (this *Digifinex) FetchMarketsV1(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetMarkets"), params ); _ = response;
  markets := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = markets;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , markets.Length )); OpIncr(& i ){
    market := *(markets ).At(i ); _ = market;
    id := this.SafeString(market , MkString("market") ); _ = id;
    baseId := MkUndefined(); _ = baseId;
    quoteId := MkUndefined(); _ = quoteId;
    ArrayUnpack(id.Split(MkString("_") ), &baseId, &quoteId);
    base := this.SafeCurrencyCode(baseId ); _ = base;
    quote := this.SafeCurrencyCode(quoteId ); _ = quote;
    symbol := OpAdd(base , OpAdd(MkString("/") , quote )); _ = symbol;
    precision := MkMap(&VarMap{
          "amount":this.SafeInteger(market , MkString("volume_precision") ),
          "price":this.SafeInteger(market , MkString("price_precision") ),
          }); _ = precision;
    limits := MkMap(&VarMap{
          "amount":MkMap(&VarMap{
              "min":this.SafeNumber(market , MkString("min_volume") ),
              "max":MkUndefined() ,
              }),
          "price":MkMap(&VarMap{
              "min":MkUndefined() ,
              "max":MkUndefined() ,
              }),
          "cost":MkMap(&VarMap{
              "min":this.SafeNumber(market , MkString("min_amount") ),
              "max":MkUndefined() ,
              }),
          }); _ = limits;
    active := OpCopy(MkUndefined() ); _ = active;
    result.Push(MkMap(&VarMap{
            "id":id ,
            "symbol":symbol ,
            "base":base ,
            "quote":quote ,
            "baseId":baseId ,
            "quoteId":quoteId ,
            "active":active ,
            "precision":precision ,
            "limits":limits ,
            "info":market ,
            }));
  }
  return result ;
}

func (this *Digifinex) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  defaultType := this.SafeString(*this.At(MkString("options")) , MkString("defaultType") , MkString("spot") ); _ = defaultType;
  type_ := this.SafeString(params , MkString("type") , defaultType ); _ = type_;
  params = this.Omit(params , MkString("type") );
  method := OpAdd(MkString("privateGet") , OpAdd(this.Capitalize(type_ ), MkString("Assets") )); _ = method;
  response := this.Call(method , params ); _ = response;
  balances := this.SafeValue(response , MkString("list") , MkArray(&VarArray{})); _ = balances;
  result := MkMap(&VarMap{"info":response }); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , balances.Length )); OpIncr(& i ){
    balance := *(balances ).At(i ); _ = balance;
    currencyId := this.SafeString(balance , MkString("currency") ); _ = currencyId;
    code := this.SafeCurrencyCode(currencyId ); _ = code;
    account := this.Account(); _ = account;
    *(account ).At(MkString("used") )= this.SafeString(balance , MkString("frozen") );
    *(account ).At(MkString("free") )= this.SafeString(balance , MkString("free") );
    *(account ).At(MkString("total") )= this.SafeString(balance , MkString("total") );
    *(result ).At(code )= OpCopy(account );
  }
  return this.ParseBalance(result );
}

func (this *Digifinex) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"symbol":*(market ).At(MkString("id") )}); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("publicGetOrderBook"), this.Extend(request , params )); _ = response;
  timestamp := this.SafeTimestamp(response , MkString("date") ); _ = timestamp;
  return this.ParseOrderBook(response , symbol , timestamp );
}

func (this *Digifinex) FetchTickers(goArgs ...*Variant) *Variant{
  symbols := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbols;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("publicGetTicker"), params ); _ = response;
  result := MkMap(&VarMap{}); _ = result;
  tickers := this.SafeValue(response , MkString("ticker") , MkArray(&VarArray{})); _ = tickers;
  date := this.SafeInteger(response , MkString("date") ); _ = date;
  for i := MkInteger(0) ; IsTrue(OpLw(i , tickers.Length )); OpIncr(& i ){
    rawTicker := this.Extend(MkMap(&VarMap{"date":date }), *(tickers ).At(i )); _ = rawTicker;
    ticker := this.ParseTicker(rawTicker ); _ = ticker;
    symbol := *(ticker ).At(MkString("symbol") ); _ = symbol;
    *(result ).At(symbol )= OpCopy(ticker );
  }
  return this.FilterByArray(result , MkString("symbol") , symbols );
}

func (this *Digifinex) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"symbol":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("publicGetTicker"), this.Extend(request , params )); _ = response;
  date := this.SafeInteger(response , MkString("date") ); _ = date;
  tickers := this.SafeValue(response , MkString("ticker") , MkArray(&VarArray{})); _ = tickers;
  firstTicker := this.SafeValue(tickers , MkInteger(0) , MkMap(&VarMap{})); _ = firstTicker;
  result := this.Extend(MkMap(&VarMap{"date":date }), firstTicker ); _ = result;
  return this.ParseTicker(result , market );
}

func (this *Digifinex) ParseTicker(goArgs ...*Variant) *Variant{
  ticker := GoGetArg(goArgs, 0, MkUndefined()); _ = ticker;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  marketId := this.SafeStringUpper(ticker , MkString("symbol") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market , MkString("_") ); _ = symbol;
  timestamp := this.SafeTimestamp(ticker , MkString("date") ); _ = timestamp;
  last := this.SafeNumber(ticker , MkString("last") ); _ = last;
  percentage := this.SafeNumber(ticker , MkString("change") ); _ = percentage;
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
        "percentage":percentage ,
        "average":MkUndefined() ,
        "baseVolume":this.SafeNumber(ticker , MkString("vol") ),
        "quoteVolume":this.SafeNumber(ticker , MkString("base_vol") ),
        "info":ticker ,
        });
}

func (this *Digifinex) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  id := this.SafeString(trade , MkString("id") ); _ = id;
  orderId := this.SafeString(trade , MkString("order_id") ); _ = orderId;
  timestamp := this.SafeTimestamp2(trade , MkString("date") , MkString("timestamp") ); _ = timestamp;
  side := this.SafeString2(trade , MkString("type") , MkString("side") ); _ = side;
  priceString := this.SafeString(trade , MkString("price") ); _ = priceString;
  amountString := this.SafeString(trade , MkString("amount") ); _ = amountString;
  price := this.ParseNumber(priceString ); _ = price;
  amount := this.ParseNumber(amountString ); _ = amount;
  cost := this.ParseNumber(Precise.StringMul(priceString , amountString )); _ = cost;
  marketId := this.SafeString(trade , MkString("symbol") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market , MkString("_") ); _ = symbol;
  takerOrMaker := this.SafeValue(trade , MkString("is_maker") ); _ = takerOrMaker;
  feeCost := this.SafeNumber(trade , MkString("fee") ); _ = feeCost;
  fee := OpCopy(MkUndefined() ); _ = fee;
  if IsTrue(OpNotEq2(feeCost , MkUndefined() )) {
    feeCurrencyId := this.SafeString(trade , MkString("fee_currency") ); _ = feeCurrencyId;
    feeCurrencyCode := this.SafeCurrencyCode(feeCurrencyId ); _ = feeCurrencyCode;
    fee = MkMap(&VarMap{
        "cost":feeCost ,
        "currency":feeCurrencyCode ,
        });
  }
  return MkMap(&VarMap{
        "id":id ,
        "info":trade ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "symbol":symbol ,
        "type":MkUndefined() ,
        "order":orderId ,
        "side":side ,
        "price":price ,
        "amount":amount ,
        "cost":cost ,
        "takerOrMaker":takerOrMaker ,
        "fee":fee ,
        });
}

func (this *Digifinex) FetchTime(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetTime"), params ); _ = response;
  return this.SafeTimestamp(response , MkString("server_time") );
}

func (this *Digifinex) FetchStatus(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.Call(MkString("publicGetPing"), params );
  *this.At(MkString("status")) = this.Extend(*this.At(MkString("status")) , MkMap(&VarMap{
          "status":MkString("ok") ,
          "updated":this.Milliseconds(),
          }));
  return *this.At(MkString("status")) ;
}

func (this *Digifinex) FetchTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"symbol":*(market ).At(MkString("id") )}); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("publicGetTrades"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = data;
  return this.ParseTrades(data , market , since , limit );
}

func (this *Digifinex) ParseOHLCV(goArgs ...*Variant) *Variant{
  ohlcv := GoGetArg(goArgs, 0, MkUndefined()); _ = ohlcv;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  return MkArray(&VarArray{
        this.SafeTimestamp(ohlcv , MkInteger(0) ),
        this.SafeNumber(ohlcv , MkInteger(5) ),
        this.SafeNumber(ohlcv , MkInteger(3) ),
        this.SafeNumber(ohlcv , MkInteger(4) ),
        this.SafeNumber(ohlcv , MkInteger(2) ),
        this.SafeNumber(ohlcv , MkInteger(1) ),
        });
}

func (this *Digifinex) FetchOHLCV(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  timeframe := GoGetArg(goArgs, 1, MkString("1m") ); _ = timeframe;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "symbol":*(market ).At(MkString("id") ),
        "period":*(*this.At(MkString("timeframes")) ).At(timeframe ),
        }); _ = request;
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    startTime := parseInt(OpDiv(since , MkInteger(1000) )); _ = startTime;
    *(request ).At(MkString("start_time") )= OpCopy(startTime );
    if IsTrue(OpNotEq2(limit , MkUndefined() )) {
      duration := this.ParseTimeframe(timeframe ); _ = duration;
      *(request ).At(MkString("end_time") )= this.Sum(startTime , OpMulti(limit , duration ));
    }
  } else {
    if IsTrue(OpNotEq2(limit , MkUndefined() )) {
      endTime := this.Seconds(); _ = endTime;
      duration := this.ParseTimeframe(timeframe ); _ = duration;
      *(request ).At(MkString("startTime") )= this.Sum(endTime , OpMulti(OpNeg(limit ), duration ));
    }
  }
  response := this.Call(MkString("publicGetKline"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = data;
  return this.ParseOHLCVs(data , market , timeframe , since , limit );
}

func (this *Digifinex) CreateOrder(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 2, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 3, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 4, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 5, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  defaultType := this.SafeString(*this.At(MkString("options")) , MkString("defaultType") , MkString("spot") ); _ = defaultType;
  orderType := this.SafeString(params , MkString("type") , defaultType ); _ = orderType;
  params = this.Omit(params , MkString("type") );
  request := MkMap(&VarMap{
        "market":orderType ,
        "symbol":*(market ).At(MkString("id") ),
        "amount":this.AmountToPrecision(symbol , amount ),
        }); _ = request;
  suffix := MkString("") ; _ = suffix;
  if IsTrue(OpEq2(type_ , MkString("market") )) {
    suffix = MkString("_market") ;
  } else {
    *(request ).At(MkString("price") )= this.PriceToPrecision(symbol , price );
  }
  *(request ).At(MkString("type") )= OpAdd(side , suffix );
  response := this.Call(MkString("privatePostMarketOrderNew"), this.Extend(request , params )); _ = response;
  result := this.ParseOrder(response , market ); _ = result;
  return this.Extend(result , MkMap(&VarMap{
            "symbol":symbol ,
            "side":side ,
            "type":type_ ,
            "amount":amount ,
            "price":price ,
            }));
}

func (this *Digifinex) CancelOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  defaultType := this.SafeString(*this.At(MkString("options")) , MkString("defaultType") , MkString("spot") ); _ = defaultType;
  orderType := this.SafeString(params , MkString("type") , defaultType ); _ = orderType;
  params = this.Omit(params , MkString("type") );
  request := MkMap(&VarMap{
        "market":orderType ,
        "order_id":id ,
        }); _ = request;
  response := this.Call(MkString("privatePostMarketOrderCancel"), this.Extend(request , params )); _ = response;
  canceledOrders := this.SafeValue(response , MkString("success") , MkArray(&VarArray{})); _ = canceledOrders;
  numCanceledOrders := OpCopy(canceledOrders.Length ); _ = numCanceledOrders;
  if IsTrue(OpNotEq2(numCanceledOrders , MkInteger(1) )) {
    panic(NewOrderNotFound(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" cancelOrder ") , OpAdd(id , MkString(" not found") )))));
  }
  return response ;
}

func (this *Digifinex) CancelOrders(goArgs ...*Variant) *Variant{
  ids := GoGetArg(goArgs, 0, MkUndefined()); _ = ids;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  defaultType := this.SafeString(*this.At(MkString("options")) , MkString("defaultType") , MkString("spot") ); _ = defaultType;
  orderType := this.SafeString(params , MkString("type") , defaultType ); _ = orderType;
  params = this.Omit(params , MkString("type") );
  request := MkMap(&VarMap{
        "market":orderType ,
        "order_id":ids.Join(MkString(",") ),
        }); _ = request;
  response := this.Call(MkString("privatePostCancelOrder"), this.Extend(request , params )); _ = response;
  canceledOrders := this.SafeValue(response , MkString("success") , MkArray(&VarArray{})); _ = canceledOrders;
  numCanceledOrders := OpCopy(canceledOrders.Length ); _ = numCanceledOrders;
  if IsTrue(OpLw(numCanceledOrders , MkInteger(1) )) {
    panic(NewOrderNotFound(OpAdd(*this.At(MkString("id")) , MkString(" cancelOrders error") )));
  }
  return response ;
}

func (this *Digifinex) ParseOrderStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "0":MkString("open") ,
        "1":MkString("open") ,
        "2":MkString("closed") ,
        "3":MkString("canceled") ,
        "4":MkString("canceled") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Digifinex) ParseOrder(goArgs ...*Variant) *Variant{
  order := GoGetArg(goArgs, 0, MkUndefined()); _ = order;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  id := this.SafeString(order , MkString("order_id") ); _ = id;
  timestamp := this.SafeTimestamp(order , MkString("created_date") ); _ = timestamp;
  lastTradeTimestamp := this.SafeTimestamp(order , MkString("finished_date") ); _ = lastTradeTimestamp;
  side := this.SafeString(order , MkString("type") ); _ = side;
  type_ := OpCopy(MkUndefined() ); _ = type_;
  if IsTrue(OpNotEq2(side , MkUndefined() )) {
    parts := side.Split(MkString("_") ); _ = parts;
    numParts := OpCopy(parts.Length ); _ = numParts;
    if IsTrue(OpGt(numParts , MkInteger(1) )) {
      side = *(parts ).At(MkInteger(0) );
      type_ = *(parts ).At(MkInteger(1) );
    } else {
      type_ = MkString("limit") ;
    }
  }
  status := this.ParseOrderStatus(this.SafeString(order , MkString("status") )); _ = status;
  marketId := this.SafeString(order , MkString("symbol") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market , MkString("_") ); _ = symbol;
  amount := this.SafeNumber(order , MkString("amount") ); _ = amount;
  filled := this.SafeNumber(order , MkString("executed_amount") ); _ = filled;
  price := this.SafeNumber(order , MkString("price") ); _ = price;
  average := this.SafeNumber(order , MkString("avg_price") ); _ = average;
  return this.SafeOrder(MkMap(&VarMap{
            "info":order ,
            "id":id ,
            "clientOrderId":MkUndefined() ,
            "timestamp":timestamp ,
            "datetime":this.Iso8601(timestamp ),
            "lastTradeTimestamp":lastTradeTimestamp ,
            "symbol":symbol ,
            "type":type_ ,
            "timeInForce":MkUndefined() ,
            "postOnly":MkUndefined() ,
            "side":side ,
            "price":price ,
            "stopPrice":MkUndefined() ,
            "amount":amount ,
            "filled":filled ,
            "remaining":MkUndefined() ,
            "cost":MkUndefined() ,
            "average":average ,
            "status":status ,
            "fee":MkUndefined() ,
            "trades":MkUndefined() ,
            }));
}

func (this *Digifinex) FetchOpenOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  defaultType := this.SafeString(*this.At(MkString("options")) , MkString("defaultType") , MkString("spot") ); _ = defaultType;
  orderType := this.SafeString(params , MkString("type") , defaultType ); _ = orderType;
  params = this.Omit(params , MkString("type") );
  this.LoadMarkets();
  market := OpCopy(MkUndefined() ); _ = market;
  request := MkMap(&VarMap{"market":orderType }); _ = request;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("symbol") )= *(market ).At(MkString("id") );
  }
  response := this.Call(MkString("privateGetMarketOrderCurrent"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = data;
  return this.ParseOrders(data , market , since , limit );
}

func (this *Digifinex) FetchOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  defaultType := this.SafeString(*this.At(MkString("options")) , MkString("defaultType") , MkString("spot") ); _ = defaultType;
  orderType := this.SafeString(params , MkString("type") , defaultType ); _ = orderType;
  params = this.Omit(params , MkString("type") );
  this.LoadMarkets();
  market := OpCopy(MkUndefined() ); _ = market;
  request := MkMap(&VarMap{"market":orderType }); _ = request;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("symbol") )= *(market ).At(MkString("id") );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("start_time") )= parseInt(OpDiv(since , MkInteger(1000) ));
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("privateGetMarketOrderHistory"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = data;
  return this.ParseOrders(data , market , since , limit );
}

func (this *Digifinex) FetchOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  defaultType := this.SafeString(*this.At(MkString("options")) , MkString("defaultType") , MkString("spot") ); _ = defaultType;
  orderType := this.SafeString(params , MkString("type") , defaultType ); _ = orderType;
  params = this.Omit(params , MkString("type") );
  this.LoadMarkets();
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
  }
  request := MkMap(&VarMap{
        "market":orderType ,
        "order_id":id ,
        }); _ = request;
  response := this.Call(MkString("privateGetMarketOrder"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = data;
  order := this.SafeValue(data , MkInteger(0) ); _ = order;
  if IsTrue(OpEq2(order , MkUndefined() )) {
    panic(NewOrderNotFound(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" fetchOrder() order ") , OpAdd(id , MkString(" not found") )))));
  }
  return this.ParseOrder(order , market );
}

func (this *Digifinex) FetchMyTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  defaultType := this.SafeString(*this.At(MkString("options")) , MkString("defaultType") , MkString("spot") ); _ = defaultType;
  orderType := this.SafeString(params , MkString("type") , defaultType ); _ = orderType;
  params = this.Omit(params , MkString("type") );
  this.LoadMarkets();
  market := OpCopy(MkUndefined() ); _ = market;
  request := MkMap(&VarMap{"market":orderType }); _ = request;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("symbol") )= *(market ).At(MkString("id") );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("start_time") )= parseInt(OpDiv(since , MkInteger(1000) ));
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("privateGetMarketMytrades"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("list") , MkArray(&VarArray{})); _ = data;
  return this.ParseTrades(data , market , since , limit );
}

func (this *Digifinex) ParseLedgerEntryType(goArgs ...*Variant) *Variant{
  type_ := GoGetArg(goArgs, 0, MkUndefined()); _ = type_;
  types := MkMap(&VarMap{}); _ = types;
  return this.SafeString(types , type_ , type_ );
}

func (this *Digifinex) ParseLedgerEntry(goArgs ...*Variant) *Variant{
  item := GoGetArg(goArgs, 0, MkUndefined()); _ = item;
  currency := GoGetArg(goArgs, 1, MkUndefined() ); _ = currency;
  id := this.SafeString(item , MkString("num") ); _ = id;
  account := OpCopy(MkUndefined() ); _ = account;
  type_ := this.ParseLedgerEntryType(this.SafeString(item , MkString("type") )); _ = type_;
  code := this.SafeCurrencyCode(this.SafeString(item , MkString("currency_mark") ), currency ); _ = code;
  timestamp := this.SafeTimestamp(item , MkString("time") ); _ = timestamp;
  before := OpCopy(MkUndefined() ); _ = before;
  after := this.SafeNumber(item , MkString("balance") ); _ = after;
  status := MkString("ok") ; _ = status;
  return MkMap(&VarMap{
        "info":item ,
        "id":id ,
        "direction":MkUndefined() ,
        "account":account ,
        "referenceId":MkUndefined() ,
        "referenceAccount":MkUndefined() ,
        "type":type_ ,
        "currency":code ,
        "amount":MkUndefined() ,
        "before":before ,
        "after":after ,
        "status":status ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "fee":MkUndefined() ,
        });
}

func (this *Digifinex) FetchLedger(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  defaultType := this.SafeString(*this.At(MkString("options")) , MkString("defaultType") , MkString("spot") ); _ = defaultType;
  orderType := this.SafeString(params , MkString("type") , defaultType ); _ = orderType;
  params = this.Omit(params , MkString("type") );
  this.LoadMarkets();
  request := MkMap(&VarMap{"market":orderType }); _ = request;
  currency := OpCopy(MkUndefined() ); _ = currency;
  if IsTrue(OpNotEq2(code , MkUndefined() )) {
    currency = this.Currency(code );
    *(request ).At(MkString("currency_mark") )= *(currency ).At(MkString("id") );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("start_time") )= parseInt(OpDiv(since , MkInteger(1000) ));
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("privateGetMarketFinancelog"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  items := this.SafeValue(data , MkString("finance") , MkArray(&VarArray{})); _ = items;
  return this.ParseLedger(items , currency , since , limit );
}

func (this *Digifinex) ParseDepositAddress(goArgs ...*Variant) *Variant{
  depositAddress := GoGetArg(goArgs, 0, MkUndefined()); _ = depositAddress;
  currency := GoGetArg(goArgs, 1, MkUndefined() ); _ = currency;
  address := this.SafeString(depositAddress , MkString("address") ); _ = address;
  tag := this.SafeString(depositAddress , MkString("addressTag") ); _ = tag;
  currencyId := this.SafeStringUpper(depositAddress , MkString("currency") ); _ = currencyId;
  code := this.SafeCurrencyCode(currencyId ); _ = code;
  return MkMap(&VarMap{
        "info":depositAddress ,
        "code":code ,
        "address":address ,
        "tag":tag ,
        });
}

func (this *Digifinex) FetchDepositAddress(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  currency := this.Currency(code ); _ = currency;
  request := MkMap(&VarMap{"currency":*(currency ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("privateGetDepositAddress"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = data;
  addresses := this.ParseDepositAddresses(data ); _ = addresses;
  address := this.SafeValue(addresses , code ); _ = address;
  if IsTrue(OpEq2(address , MkUndefined() )) {
    panic(NewInvalidAddress(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" fetchDepositAddress did not return an address for ") , OpAdd(code , MkString(" - create the deposit address in the user settings on the exchange website first.") )))));
  }
  return address ;
}

func (this *Digifinex) FetchTransactionsByType(goArgs ...*Variant) *Variant{
  type_ := GoGetArg(goArgs, 0, MkUndefined()); _ = type_;
  code := GoGetArg(goArgs, 1, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  currency := OpCopy(MkUndefined() ); _ = currency;
  request := MkMap(&VarMap{}); _ = request;
  if IsTrue(OpNotEq2(code , MkUndefined() )) {
    currency = this.Currency(code );
    *(request ).At(MkString("currency") )= *(currency ).At(MkString("id") );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("size") )= Math.Min(MkInteger(500) , limit );
  }
  method := OpTrinary(OpEq2(type_ , MkString("deposit") ), MkString("privateGetDepositHistory") , MkString("privateGetWithdrawHistory") ); _ = method;
  response := this.Call(method , this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = data;
  return this.ParseTransactions(data , currency , since , limit , MkMap(&VarMap{"type":type_ }));
}

func (this *Digifinex) FetchDeposits(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  return this.FetchTransactionsByType(MkString("deposit") , code , since , limit , params );
}

func (this *Digifinex) FetchWithdrawals(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  return this.FetchTransactionsByType(MkString("withdrawal") , code , since , limit , params );
}

func (this *Digifinex) ParseTransactionStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "0":MkString("pending") ,
        "1":MkString("canceled") ,
        "2":MkString("pending") ,
        "3":MkString("failed") ,
        "4":MkString("pending") ,
        "5":MkString("failed") ,
        "6":MkString("ok") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Digifinex) ParseTransaction(goArgs ...*Variant) *Variant{
  transaction := GoGetArg(goArgs, 0, MkUndefined()); _ = transaction;
  currency := GoGetArg(goArgs, 1, MkUndefined() ); _ = currency;
  id := this.SafeString2(transaction , MkString("id") , MkString("withdraw_id") ); _ = id;
  address := this.SafeString(transaction , MkString("address") ); _ = address;
  tag := this.SafeString(transaction , MkString("memo") ); _ = tag;
  if IsTrue(OpNotEq2(tag , MkUndefined() )) {
    if IsTrue(OpLw(tag.Length , MkInteger(1) )) {
      tag = OpCopy(MkUndefined() );
    }
  }
  txid := this.SafeString(transaction , MkString("hash") ); _ = txid;
  currencyId := this.SafeStringUpper(transaction , MkString("currency") ); _ = currencyId;
  code := this.SafeCurrencyCode(currencyId , currency ); _ = code;
  timestamp := this.Parse8601(this.SafeString(transaction , MkString("created_date") )); _ = timestamp;
  updated := this.Parse8601(this.SafeString(transaction , MkString("finished_date") )); _ = updated;
  status := this.ParseTransactionStatus(this.SafeString(transaction , MkString("state") )); _ = status;
  amount := this.SafeNumber(transaction , MkString("amount") ); _ = amount;
  feeCost := this.SafeNumber(transaction , MkString("fee") ); _ = feeCost;
  fee := OpCopy(MkUndefined() ); _ = fee;
  if IsTrue(OpNotEq2(feeCost , MkUndefined() )) {
    fee = MkMap(&VarMap{
        "currency":code ,
        "cost":feeCost ,
        });
  }
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
        "type":MkUndefined() ,
        "amount":amount ,
        "currency":code ,
        "status":status ,
        "updated":updated ,
        "fee":fee ,
        });
}

func (this *Digifinex) Withdraw(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  amount := GoGetArg(goArgs, 1, MkUndefined()); _ = amount;
  address := GoGetArg(goArgs, 2, MkUndefined()); _ = address;
  tag := GoGetArg(goArgs, 3, MkUndefined() ); _ = tag;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.CheckAddress(address );
  this.LoadMarkets();
  currency := this.Currency(code ); _ = currency;
  request := MkMap(&VarMap{
        "address":address ,
        "amount":parseFloat(amount ),
        "currency":*(currency ).At(MkString("id") ),
        }); _ = request;
  if IsTrue(OpNotEq2(tag , MkUndefined() )) {
    *(request ).At(MkString("memo") )= OpCopy(tag );
  }
  response := this.Call(MkString("privatePostWithdrawNew"), this.Extend(request , params )); _ = response;
  return this.ParseTransaction(response , currency );
}

func (this *Digifinex) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("public") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  version := OpCopy(*this.At(MkString("version")) ); _ = version;
  url := OpAdd(*(*this.At(MkString("urls")) ).At(MkString("api") ), OpAdd(MkString("/") , OpAdd(version , OpAdd(MkString("/") , this.ImplodeParams(path , params ))))); _ = url;
  query := this.Omit(params , this.ExtractParams(path )); _ = query;
  urlencoded := this.Urlencode(this.Keysort(query )); _ = urlencoded;
  if IsTrue(OpEq2(api , MkString("private") )) {
    nonce := (this.Nonce()).Call(MkString("toString") ); _ = nonce;
    auth := OpCopy(urlencoded ); _ = auth;
    signature := this.Hmac(this.Encode(auth ), this.Encode(*this.At(MkString("secret")) )); _ = signature;
    if IsTrue(OpEq2(method , MkString("GET") )) {
      if IsTrue(urlencoded ) {
        OpAddAssign(& url , OpAdd(MkString("?") , urlencoded ));
      }
    } else {
      if IsTrue(OpEq2(method , MkString("POST") )) {
        headers = MkMap(&VarMap{"Content-Type":MkString("application/x-www-form-urlencoded") });
        if IsTrue(urlencoded ) {
          body = OpCopy(urlencoded );
        }
      }
    }
    headers = MkMap(&VarMap{
        "ACCESS-KEY":*this.At(MkString("apiKey")) ,
        "ACCESS-SIGN":signature ,
        "ACCESS-TIMESTAMP":nonce ,
        });
  } else {
    if IsTrue(urlencoded ) {
      OpAddAssign(& url , OpAdd(MkString("?") , urlencoded ));
    }
  }
  return MkMap(&VarMap{
        "url":url ,
        "method":method ,
        "body":body ,
        "headers":headers ,
        });
}

func (this *Digifinex) HandleErrors(goArgs ...*Variant) *Variant{
  statusCode := GoGetArg(goArgs, 0, MkUndefined()); _ = statusCode;
  statusText := GoGetArg(goArgs, 1, MkUndefined()); _ = statusText;
  url := GoGetArg(goArgs, 2, MkUndefined()); _ = url;
  method := GoGetArg(goArgs, 3, MkUndefined()); _ = method;
  responseHeaders := GoGetArg(goArgs, 4, MkUndefined()); _ = responseHeaders;
  responseBody := GoGetArg(goArgs, 5, MkUndefined()); _ = responseBody;
  response := GoGetArg(goArgs, 6, MkUndefined()); _ = response;
  requestHeaders := GoGetArg(goArgs, 7, MkUndefined()); _ = requestHeaders;
  requestBody := GoGetArg(goArgs, 8, MkUndefined()); _ = requestBody;
  if IsTrue(OpNot(response )) {
    return MkUndefined();
  }
  code := this.SafeString(response , MkString("code") ); _ = code;
  if IsTrue(OpOr(OpEq2(code , MkString("0") ), OpEq2(code , MkString("200") ))) {
    return MkUndefined();
  }
  feedback := OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , responseBody )); _ = feedback;
  if IsTrue(OpEq2(code , MkUndefined() )) {
    panic(NewBadResponse(feedback ));
  }
  unknownError := MkArray(&VarArray{
        ExchangeError ,
        feedback ,
        }); _ = unknownError;
  ExceptionClass := MkUndefined(); _ = ExceptionClass;
  message := MkUndefined(); _ = message;
  ArrayUnpack(this.SafeValue(*(*this.At(MkString("exceptions")) ).At(MkString("exact") ), code , unknownError ), &ExceptionClass, &message);
  panic(NewExceptionClass(message ));
  return MkUndefined()
}

