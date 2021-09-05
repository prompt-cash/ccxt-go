package ccxt

import (
)

type Bl3p struct {
	*ExchangeBase
}

var _ Exchange = (*Bl3p)(nil)

func init() {
	exchange := &Bl3p{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Bl3p) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("bl3p") ,
            "name":MkString("BL3P") ,
            "countries":MkArray(&VarArray{
                MkString("NL") ,
                MkString("EU") ,
                }),
            "rateLimit":MkInteger(1000) ,
            "version":MkString("1") ,
            "comment":MkString("An exchange market by BitonicNL") ,
            "has":MkMap(&VarMap{
                "CORS":MkBool(false) ,
                "cancelOrder":MkBool(true) ,
                "createOrder":MkBool(true) ,
                "fetchBalance":MkBool(true) ,
                "fetchOrderBook":MkBool(true) ,
                "fetchTicker":MkBool(true) ,
                "fetchTrades":MkBool(true) ,
                }),
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/1294454/28501752-60c21b82-6feb-11e7-818b-055ee6d0e754.jpg") ,
                "api":MkString("https://api.bl3p.eu") ,
                "www":MkString("https://bl3p.eu") ,
                "doc":MkArray(&VarArray{
                    MkString("https://github.com/BitonicNL/bl3p-api/tree/master/docs") ,
                    MkString("https://bl3p.eu/api") ,
                    MkString("https://bitonic.nl/en/api") ,
                    }),
                }),
            "api":MkMap(&VarMap{
                "public":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("{market}/ticker") ,
                        MkString("{market}/orderbook") ,
                        MkString("{market}/trades") ,
                        })}),
                "private":MkMap(&VarMap{"post":MkArray(&VarArray{
                        MkString("{market}/money/depth/full") ,
                        MkString("{market}/money/order/add") ,
                        MkString("{market}/money/order/cancel") ,
                        MkString("{market}/money/order/result") ,
                        MkString("{market}/money/orders") ,
                        MkString("{market}/money/orders/history") ,
                        MkString("{market}/money/trades/fetch") ,
                        MkString("GENMKT/money/info") ,
                        MkString("GENMKT/money/deposit_address") ,
                        MkString("GENMKT/money/new_deposit_address") ,
                        MkString("GENMKT/money/wallet/history") ,
                        MkString("GENMKT/money/withdraw") ,
                        })}),
                }),
            "markets":MkMap(&VarMap{
                "BTC/EUR":MkMap(&VarMap{
                    "id":MkString("BTCEUR") ,
                    "symbol":MkString("BTC/EUR") ,
                    "base":MkString("BTC") ,
                    "quote":MkString("EUR") ,
                    "baseId":MkString("BTC") ,
                    "quoteId":MkString("EUR") ,
                    "maker":MkNumber(0.0025) ,
                    "taker":MkNumber(0.0025) ,
                    }),
                "LTC/EUR":MkMap(&VarMap{
                    "id":MkString("LTCEUR") ,
                    "symbol":MkString("LTC/EUR") ,
                    "base":MkString("LTC") ,
                    "quote":MkString("EUR") ,
                    "baseId":MkString("LTC") ,
                    "quoteId":MkString("EUR") ,
                    "maker":MkNumber(0.0025) ,
                    "taker":MkNumber(0.0025) ,
                    }),
                }),
            }));
}

func (this *Bl3p) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("privatePostGENMKTMoneyInfo"), params ); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  wallets := this.SafeValue(data , MkString("wallets") ); _ = wallets;
  result := MkMap(&VarMap{"info":data }); _ = result;
  codes := GoGetKeys(*this.At(MkString("currencies")) ); _ = codes;
  for i := MkInteger(0) ; IsTrue(OpLw(i , codes.Length )); OpIncr(& i ){
    code := *(codes ).At(i ); _ = code;
    currency := this.Currency(code ); _ = currency;
    currencyId := *(currency ).At(MkString("id") ); _ = currencyId;
    wallet := this.SafeValue(wallets , currencyId , MkMap(&VarMap{})); _ = wallet;
    available := this.SafeValue(wallet , MkString("available") , MkMap(&VarMap{})); _ = available;
    balance := this.SafeValue(wallet , MkString("balance") , MkMap(&VarMap{})); _ = balance;
    account := this.Account(); _ = account;
    *(account ).At(MkString("free") )= this.SafeString(available , MkString("value") );
    *(account ).At(MkString("total") )= this.SafeString(balance , MkString("value") );
    *(result ).At(code )= OpCopy(account );
  }
  return this.ParseBalance(result );
}

func (this *Bl3p) ParseBidAsk(goArgs ...*Variant) *Variant{
  bidask := GoGetArg(goArgs, 0, MkUndefined()); _ = bidask;
  priceKey := GoGetArg(goArgs, 1, MkInteger(0) ); _ = priceKey;
  amountKey := GoGetArg(goArgs, 2, MkInteger(1) ); _ = amountKey;
  price := this.SafeNumber(bidask , priceKey ); _ = price;
  size := this.SafeNumber(bidask , amountKey ); _ = size;
  return MkArray(&VarArray{
        OpDiv(price , MkNumber(100000.0) ),
        OpDiv(size , MkNumber(100000000.0) ),
        });
}

func (this *Bl3p) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"market":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("publicGetMarketOrderbook"), this.Extend(request , params )); _ = response;
  orderbook := this.SafeValue(response , MkString("data") ); _ = orderbook;
  return this.ParseOrderBook(orderbook , symbol , MkUndefined() , MkString("bids") , MkString("asks") , MkString("price_int") , MkString("amount_int") );
}

func (this *Bl3p) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"market":this.MarketId(symbol )}); _ = request;
  ticker := this.Call(MkString("publicGetMarketTicker"), this.Extend(request , params )); _ = ticker;
  timestamp := this.SafeTimestamp(ticker , MkString("timestamp") ); _ = timestamp;
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
        "baseVolume":this.SafeNumber(*(ticker ).At(MkString("volume") ), MkString("24h") ),
        "quoteVolume":MkUndefined() ,
        "info":ticker ,
        });
}

func (this *Bl3p) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  id := this.SafeString(trade , MkString("trade_id") ); _ = id;
  timestamp := this.SafeInteger(trade , MkString("date") ); _ = timestamp;
  priceString := this.SafeString(trade , MkString("price_int") ); _ = priceString;
  priceString = Precise.StringDiv(priceString , MkString("100000") );
  amountString := this.SafeString(trade , MkString("amount_int") ); _ = amountString;
  amountString = Precise.StringDiv(amountString , MkString("100000000") );
  price := this.ParseNumber(priceString ); _ = price;
  amount := this.ParseNumber(amountString ); _ = amount;
  cost := this.ParseNumber(Precise.StringMul(priceString , amountString )); _ = cost;
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  if IsTrue(OpNotEq2(market , MkUndefined() )) {
    symbol = *(market ).At(MkString("symbol") );
  }
  return MkMap(&VarMap{
        "id":id ,
        "info":trade ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "symbol":symbol ,
        "type":MkUndefined() ,
        "side":MkUndefined() ,
        "order":MkUndefined() ,
        "takerOrMaker":MkUndefined() ,
        "price":price ,
        "amount":amount ,
        "cost":cost ,
        "fee":MkUndefined() ,
        });
}

func (this *Bl3p) FetchTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  market := this.Market(symbol ); _ = market;
  response := this.Call(MkString("publicGetMarketTrades"), this.Extend(MkMap(&VarMap{"market":*(market ).At(MkString("id") )}), params )); _ = response;
  result := this.ParseTrades(*(*(response ).At(MkString("data") )).At(MkString("trades") ), market , since , limit ); _ = result;
  return result ;
}

func (this *Bl3p) CreateOrder(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 2, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 3, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 4, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 5, MkMap(&VarMap{})); _ = params;
  market := this.Market(symbol ); _ = market;
  order := MkMap(&VarMap{
        "market":*(market ).At(MkString("id") ),
        "amount_int":parseInt(OpMulti(amount , MkInteger(100000000) )),
        "fee_currency":*(market ).At(MkString("quote") ),
        "type":OpTrinary(OpEq2(side , MkString("buy") ), MkString("bid") , MkString("ask") ),
        }); _ = order;
  if IsTrue(OpEq2(type_ , MkString("limit") )) {
    *(order ).At(MkString("price_int") )= parseInt(OpMulti(price , MkNumber(100000.0) ));
  }
  response := this.Call(MkString("privatePostMarketMoneyOrderAdd"), this.Extend(order , params )); _ = response;
  orderId := this.SafeString(*(response ).At(MkString("data") ), MkString("order_id") ); _ = orderId;
  return MkMap(&VarMap{
        "info":response ,
        "id":orderId ,
        });
}

func (this *Bl3p) CancelOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"order_id":id }); _ = request;
  return this.Call(MkString("privatePostMarketMoneyOrderCancel"), this.Extend(request , params ));
}

func (this *Bl3p) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("public") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  request := this.ImplodeParams(path , params ); _ = request;
  url := OpAdd(*(*this.At(MkString("urls")) ).At(MkString("api") ), OpAdd(MkString("/") , OpAdd(*this.At(MkString("version")) , OpAdd(MkString("/") , request )))); _ = url;
  query := this.Omit(params , this.ExtractParams(path )); _ = query;
  if IsTrue(OpEq2(api , MkString("public") )) {
    if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
      OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(query )));
    }
  } else {
    this.CheckRequiredCredentials();
    nonce := this.Nonce(); _ = nonce;
    body = this.Urlencode(this.Extend(MkMap(&VarMap{"nonce":nonce }), query ));
    secret := this.Base64ToBinary(*this.At(MkString("secret")) ); _ = secret;
    auth := OpAdd(request , OpAdd(MkString("\000") , body )); _ = auth;
    signature := this.Hmac(this.Encode(auth ), secret , MkString("sha512") , MkString("base64") ); _ = signature;
    headers = MkMap(&VarMap{
        "Content-Type":MkString("application/x-www-form-urlencoded") ,
        "Rest-Key":*this.At(MkString("apiKey")) ,
        "Rest-Sign":signature ,
        });
  }
  return MkMap(&VarMap{
        "url":url ,
        "method":method ,
        "body":body ,
        "headers":headers ,
        });
}

