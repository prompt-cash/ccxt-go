package ccxt

import (
)

type Flowbtc struct {
	*ExchangeBase
}

var _ Exchange = (*Flowbtc)(nil)

func init() {
	exchange := &Flowbtc{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Flowbtc) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("flowbtc") ,
            "name":MkString("flowBTC") ,
            "countries":MkArray(&VarArray{
                MkString("BR") ,
                }),
            "version":MkString("v1") ,
            "rateLimit":MkInteger(1000) ,
            "has":MkMap(&VarMap{
                "cancelOrder":MkBool(true) ,
                "CORS":MkBool(false) ,
                "createOrder":MkBool(true) ,
                "fetchBalance":MkBool(true) ,
                "fetchMarkets":MkBool(true) ,
                "fetchOrderBook":MkBool(true) ,
                "fetchTicker":MkBool(true) ,
                "fetchTrades":MkBool(true) ,
                }),
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/51840849/87443317-01c0d080-c5fe-11ea-95c2-9ebe1a8fafd9.jpg") ,
                "api":MkString("https://publicapi.flowbtc.com.br") ,
                "www":MkString("https://www.flowbtc.com.br") ,
                "doc":MkString("https://www.flowbtc.com.br/api.html") ,
                }),
            "requiredCredentials":MkMap(&VarMap{
                "apiKey":MkBool(true) ,
                "secret":MkBool(true) ,
                "uid":MkBool(true) ,
                }),
            "api":MkMap(&VarMap{
                "public":MkMap(&VarMap{"post":MkArray(&VarArray{
                        MkString("GetTicker") ,
                        MkString("GetTrades") ,
                        MkString("GetTradesByDate") ,
                        MkString("GetOrderBook") ,
                        MkString("GetProductPairs") ,
                        MkString("GetProducts") ,
                        })}),
                "private":MkMap(&VarMap{"post":MkArray(&VarArray{
                        MkString("CreateAccount") ,
                        MkString("GetUserInfo") ,
                        MkString("SetUserInfo") ,
                        MkString("GetAccountInfo") ,
                        MkString("GetAccountTrades") ,
                        MkString("GetDepositAddresses") ,
                        MkString("Withdraw") ,
                        MkString("CreateOrder") ,
                        MkString("ModifyOrder") ,
                        MkString("CancelOrder") ,
                        MkString("CancelAllOrders") ,
                        MkString("GetAccountOpenOrders") ,
                        MkString("GetOrderFee") ,
                        })}),
                }),
            "fees":MkMap(&VarMap{"trading":MkMap(&VarMap{
                    "tierBased":MkBool(false) ,
                    "percentage":MkBool(true) ,
                    "maker":MkNumber(0.0025) ,
                    "taker":MkNumber(0.005) ,
                    })}),
            }));
}

func (this *Flowbtc) FetchMarkets(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicPostGetProductPairs"), params ); _ = response;
  markets := this.SafeValue(response , MkString("productPairs") ); _ = markets;
  result := MkMap(&VarMap{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , markets.Length )); OpIncr(& i ){
    market := *(markets ).At(i ); _ = market;
    id := this.SafeString(market , MkString("name") ); _ = id;
    baseId := this.SafeString(market , MkString("product1Label") ); _ = baseId;
    quoteId := this.SafeString(market , MkString("product2Label") ); _ = quoteId;
    base := this.SafeCurrencyCode(baseId ); _ = base;
    quote := this.SafeCurrencyCode(quoteId ); _ = quote;
    precision := MkMap(&VarMap{
          "amount":this.SafeInteger(market , MkString("product1DecimalPlaces") ),
          "price":this.SafeInteger(market , MkString("product2DecimalPlaces") ),
          }); _ = precision;
    symbol := OpAdd(base , OpAdd(MkString("/") , quote )); _ = symbol;
    *(result ).At(symbol )= MkMap(&VarMap{
        "id":id ,
        "symbol":symbol ,
        "base":base ,
        "quote":quote ,
        "baseId":baseId ,
        "quoteId":quoteId ,
        "precision":precision ,
        "limits":MkMap(&VarMap{
            "amount":MkMap(&VarMap{
                "min":MkUndefined() ,
                "max":MkUndefined() ,
                }),
            "price":MkMap(&VarMap{
                "min":MkUndefined() ,
                "max":MkUndefined() ,
                }),
            "cost":MkMap(&VarMap{
                "min":MkUndefined() ,
                "max":MkUndefined() ,
                }),
            }),
        "info":market ,
        "active":MkUndefined() ,
        });
  }
  return result ;
}

func (this *Flowbtc) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("privatePostGetAccountInfo"), params ); _ = response;
  balances := this.SafeValue(response , MkString("currencies") ); _ = balances;
  result := MkMap(&VarMap{"info":response }); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , balances.Length )); OpIncr(& i ){
    balance := *(balances ).At(i ); _ = balance;
    currencyId := *(balance ).At(MkString("name") ); _ = currencyId;
    code := this.SafeCurrencyCode(currencyId ); _ = code;
    account := this.Account(); _ = account;
    *(account ).At(MkString("free") )= this.SafeString(balance , MkString("balance") );
    *(account ).At(MkString("total") )= this.SafeString(balance , MkString("hold") );
    *(result ).At(code )= OpCopy(account );
  }
  return this.ParseBalance(result );
}

func (this *Flowbtc) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"productPair":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("publicPostGetOrderBook"), this.Extend(request , params )); _ = response;
  return this.ParseOrderBook(response , symbol , MkUndefined() , MkString("bids") , MkString("asks") , MkString("px") , MkString("qty") );
}

func (this *Flowbtc) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"productPair":*(market ).At(MkString("id") )}); _ = request;
  ticker := this.Call(MkString("publicPostGetTicker"), this.Extend(request , params )); _ = ticker;
  timestamp := this.Milliseconds(); _ = timestamp;
  last := this.SafeNumber(ticker , MkString("last") ); _ = last;
  return MkMap(&VarMap{
        "symbol":symbol ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "high":this.SafeNumber(ticker , MkString("high") ),
        "low":this.SafeNumber(ticker , MkString("low") ),
        "bid":this.SafeNumber(ticker , MkString("bid") ),
        "bidVolume":MkUndefined() ,
        "ask":this.SafeNumber(ticker , MkString("ask") ),
        "askVolume":MkUndefined() ,
        "vwap":MkUndefined() ,
        "open":MkUndefined() ,
        "close":last ,
        "last":last ,
        "previousClose":MkUndefined() ,
        "change":MkUndefined() ,
        "percentage":MkUndefined() ,
        "average":MkUndefined() ,
        "baseVolume":this.SafeNumber(ticker , MkString("volume24hr") ),
        "quoteVolume":this.SafeNumber(ticker , MkString("volume24hrProduct2") ),
        "info":ticker ,
        });
}

func (this *Flowbtc) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined()); _ = market;
  timestamp := this.SafeTimestamp(trade , MkString("unixtime") ); _ = timestamp;
  side := OpTrinary(OpEq2(*(trade ).At(MkString("incomingOrderSide") ), MkInteger(0) ), MkString("buy") , MkString("sell") ); _ = side;
  id := this.SafeString(trade , MkString("tid") ); _ = id;
  priceString := this.SafeString(trade , MkString("px") ); _ = priceString;
  amountString := this.SafeString(trade , MkString("qty") ); _ = amountString;
  price := this.ParseNumber(priceString ); _ = price;
  amount := this.ParseNumber(amountString ); _ = amount;
  cost := this.ParseNumber(Precise.StringMul(priceString , amountString )); _ = cost;
  return MkMap(&VarMap{
        "info":trade ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "symbol":*(market ).At(MkString("symbol") ),
        "id":id ,
        "order":MkUndefined() ,
        "type":MkUndefined() ,
        "side":side ,
        "price":price ,
        "amount":amount ,
        "cost":cost ,
        "takerOrMaker":MkUndefined() ,
        "fee":MkUndefined() ,
        });
}

func (this *Flowbtc) FetchTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "ins":*(market ).At(MkString("id") ),
        "startIndex":OpNeg(MkInteger(1) ),
        }); _ = request;
  response := this.Call(MkString("publicPostGetTrades"), this.Extend(request , params )); _ = response;
  return this.ParseTrades(*(response ).At(MkString("trades") ), market , since , limit );
}

func (this *Flowbtc) CreateOrder(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 2, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 3, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 4, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 5, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  orderType := OpTrinary(OpEq2(type_ , MkString("market") ), MkInteger(1) , MkInteger(0) ); _ = orderType;
  request := MkMap(&VarMap{
        "ins":this.MarketId(symbol ),
        "side":side ,
        "orderType":orderType ,
        "qty":amount ,
        "px":this.PriceToPrecision(symbol , price ),
        }); _ = request;
  response := this.Call(MkString("privatePostCreateOrder"), this.Extend(request , params )); _ = response;
  return MkMap(&VarMap{
        "info":response ,
        "id":*(response ).At(MkString("serverOrderId") ),
        });
}

func (this *Flowbtc) CancelOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  if IsTrue(OpHasMember(MkString("ins") , params )) {
    request := MkMap(&VarMap{"serverOrderId":id }); _ = request;
    return this.Call(MkString("privatePostCancelOrder"), this.Extend(request , params ));
  }
  panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , MkString(" cancelOrder() requires an `ins` symbol parameter for cancelling an order") )));
  return MkUndefined()
}

func (this *Flowbtc) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("public") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  url := OpAdd(*(*this.At(MkString("urls")) ).At(MkString("api") ), OpAdd(MkString("/") , OpAdd(*this.At(MkString("version")) , OpAdd(MkString("/") , path )))); _ = url;
  if IsTrue(OpEq2(api , MkString("public") )) {
    if IsTrue(*(GoGetKeys(params )).At(MkString("length") )) {
      body = this.Json(params );
    }
  } else {
    this.CheckRequiredCredentials();
    nonce := this.Nonce(); _ = nonce;
    auth := OpAdd(nonce.ToString(), OpAdd(*this.At(MkString("uid")) , *this.At(MkString("apiKey")) )); _ = auth;
    signature := this.Hmac(this.Encode(auth ), this.Encode(*this.At(MkString("secret")) )); _ = signature;
    body = this.Json(this.Extend(MkMap(&VarMap{
                "apiKey":*this.At(MkString("apiKey")) ,
                "apiNonce":nonce ,
                "apiSig":signature.ToUpperCase(),
                }), params ));
    headers = MkMap(&VarMap{"Content-Type":MkString("application/json") });
  }
  return MkMap(&VarMap{
        "url":url ,
        "method":method ,
        "body":body ,
        "headers":headers ,
        });
}

func (this *Flowbtc) HandleErrors(goArgs ...*Variant) *Variant{
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
  isAccepted := this.SafeValue(response , MkString("isAccepted") , MkBool(true) ); _ = isAccepted;
  if IsTrue(OpNot(isAccepted )) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , this.Json(response )))));
  }
  return MkUndefined()
}

