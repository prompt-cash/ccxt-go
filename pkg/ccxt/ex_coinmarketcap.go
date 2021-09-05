package ccxt

import (
)

type Coinmarketcap struct {
	*ExchangeBase
}

var _ Exchange = (*Coinmarketcap)(nil)

func init() {
	exchange := &Coinmarketcap{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Coinmarketcap) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("coinmarketcap") ,
            "name":MkString("CoinMarketCap") ,
            "rateLimit":MkInteger(10000) ,
            "version":MkString("v1") ,
            "countries":MkArray(&VarArray{
                MkString("US") ,
                }),
            "has":MkMap(&VarMap{
                "cancelOrder":MkBool(false) ,
                "CORS":MkBool(true) ,
                "createLimitOrder":MkBool(false) ,
                "createMarketOrder":MkBool(false) ,
                "createOrder":MkBool(false) ,
                "editOrder":MkBool(false) ,
                "privateAPI":MkBool(false) ,
                "fetchBalance":MkBool(false) ,
                "fetchCurrencies":MkBool(true) ,
                "fetchL2OrderBook":MkBool(false) ,
                "fetchMarkets":MkBool(true) ,
                "fetchOHLCV":MkBool(false) ,
                "fetchOrderBook":MkBool(false) ,
                "fetchTicker":MkBool(true) ,
                "fetchTickers":MkBool(true) ,
                "fetchTrades":MkBool(false) ,
                }),
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/51840849/87182086-1cd4cd00-c2ec-11ea-9ec4-d0cf2a2abf62.jpg") ,
                "api":MkMap(&VarMap{
                    "public":MkString("https://api.coinmarketcap.com") ,
                    "files":MkString("https://files.coinmarketcap.com") ,
                    "charts":MkString("https://graph.coinmarketcap.com") ,
                    }),
                "www":MkString("https://coinmarketcap.com") ,
                "doc":MkString("https://coinmarketcap.com/api") ,
                }),
            "requiredCredentials":MkMap(&VarMap{
                "apiKey":MkBool(false) ,
                "secret":MkBool(false) ,
                }),
            "api":MkMap(&VarMap{
                "files":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("generated/stats/global.json") ,
                        })}),
                "graphs":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("currencies/{name}/") ,
                        })}),
                "public":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("ticker/") ,
                        MkString("ticker/{id}/") ,
                        MkString("global/") ,
                        })}),
                }),
            "currencyCodes":MkArray(&VarArray{
                MkString("AUD") ,
                MkString("BRL") ,
                MkString("CAD") ,
                MkString("CHF") ,
                MkString("CNY") ,
                MkString("EUR") ,
                MkString("GBP") ,
                MkString("HKD") ,
                MkString("IDR") ,
                MkString("INR") ,
                MkString("JPY") ,
                MkString("KRW") ,
                MkString("MXN") ,
                MkString("RUB") ,
                MkString("USD") ,
                MkString("BTC") ,
                MkString("ETH") ,
                MkString("LTC") ,
                }),
            }));
}

func (this *Coinmarketcap) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  panic(NewExchangeError(OpAdd(MkString("Fetching order books is not supported by the API of ") , *this.At(MkString("id")) )));
  return MkUndefined()
}

func (this *Coinmarketcap) CurrencyCode(goArgs ...*Variant) *Variant{
  base := GoGetArg(goArgs, 0, MkUndefined()); _ = base;
  name := GoGetArg(goArgs, 1, MkUndefined()); _ = name;
  currencies := MkMap(&VarMap{
        "ACChain":MkString("ACChain") ,
        "AdCoin":MkString("AdCoin") ,
        "BatCoin":MkString("BatCoin") ,
        "BigONE Token":MkString("BigONE Token") ,
        "Bitgem":MkString("Bitgem") ,
        "BlazeCoin":MkString("BlazeCoin") ,
        "BlockCAT":MkString("BlockCAT") ,
        "Blocktrade Token":MkString("Blocktrade Token") ,
        "BOX Token":MkString("BOX Token") ,
        "Catcoin":MkString("Catcoin") ,
        "CanYaCoin":MkString("CanYaCoin") ,
        "CryptoBossCoin":MkString("CryptoBossCoin") ,
        "Comet":MkString("Comet") ,
        "CPChain":MkString("CPChain") ,
        "CrowdCoin":MkString("CrowdCoin") ,
        "Cryptaur":MkString("Cryptaur") ,
        "Cubits":MkString("Cubits") ,
        "DAO.Casino":MkString("DAO.Casino") ,
        "DefiBox":MkString("DefiBox") ,
        "E-Dinar Coin":MkString("E-Dinar Coin") ,
        "EDRcoin":MkString("EDRcoin") ,
        "ENTCash":MkString("ENTCash") ,
        "FairCoin":MkString("FairCoin") ,
        "Fabric Token":MkString("Fabric Token") ,
        "GHOSTPRISM":MkString("GHOSTPRISM") ,
        "Global Tour Coin":MkString("Global Tour Coin") ,
        "GuccioneCoin":MkString("GuccioneCoin") ,
        "HarmonyCoin":MkString("HarmonyCoin") ,
        "Harvest Masternode Coin":MkString("Harvest Masternode Coin") ,
        "HOT Token":MkString("HOT Token") ,
        "Hydro Protocol":MkString("Hydro Protocol") ,
        "Huncoin":MkString("Huncoin") ,
        "iCoin":MkString("iCoin") ,
        "Infinity Economics":MkString("Infinity Economics") ,
        "IQ.cash":MkString("IQ.cash") ,
        "KingN Coin":MkString("KingN Coin") ,
        "LiteBitcoin":MkString("LiteBitcoin") ,
        "Maggie":MkString("Maggie") ,
        "Menlo One":MkString("Menlo One") ,
        "Mobilian Coin":MkString("Mobilian Coin") ,
        "Monarch":MkString("Monarch") ,
        "MTC Mesh Network":MkString("MTC Mesh Network") ,
        "IOTA":MkString("IOTA") ,
        "NetCoin":MkString("NetCoin") ,
        "PCHAIN":MkString("PCHAIN") ,
        "Penta":MkString("Penta") ,
        "Plair":MkString("Plair") ,
        "PlayChip":MkString("PlayChip") ,
        "Polcoin":MkString("Polcoin") ,
        "PutinCoin":MkString("PutinCoin") ,
        "Rcoin":MkString("Rcoin") ,
        "SBTCT":MkString("SiamBitcoin") ,
        "Super Bitcoin":MkString("Super Bitcoin") ,
        "TerraCredit":MkString("TerraCredit") ,
        "Themis":MkString("Themis") ,
        "UNI COIN":MkString("UNI COIN") ,
        "UNICORN Token":MkString("UNICORN Token") ,
        "Universe":MkString("Universe") ,
        }); _ = currencies;
  return this.SafeValue(currencies , name , base );
}

func (this *Coinmarketcap) FetchMarkets(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"limit":MkInteger(0) }); _ = request;
  response := this.Call(MkString("publicGetTicker"), this.Extend(request , params )); _ = response;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , response.Length )); OpIncr(& i ){
    market := *(response ).At(i ); _ = market;
    currencies := OpCopy(*this.At(MkString("currencyCodes")) ); _ = currencies;
    for j := MkInteger(0) ; IsTrue(OpLw(j , currencies.Length )); OpIncr(& j ){
      quote := *(currencies ).At(j ); _ = quote;
      quoteId := quote.ToLowerCase(); _ = quoteId;
      baseId := *(market ).At(MkString("id") ); _ = baseId;
      base := this.CurrencyCode(*(market ).At(MkString("symbol") ), *(market ).At(MkString("name") )); _ = base;
      symbol := OpAdd(base , OpAdd(MkString("/") , quote )); _ = symbol;
      id := OpAdd(baseId , OpAdd(MkString("/") , quoteId )); _ = id;
      result.Push(MkMap(&VarMap{
              "id":id ,
              "symbol":symbol ,
              "base":base ,
              "quote":quote ,
              "baseId":baseId ,
              "quoteId":quoteId ,
              "info":market ,
              "active":MkUndefined() ,
              "precision":*this.At(MkString("precision")) ,
              "limits":*this.At(MkString("limits")) ,
              }));
    }
  }
  return result ;
}

func (this *Coinmarketcap) FetchGlobal(goArgs ...*Variant) *Variant{
  currency := GoGetArg(goArgs, 0, MkString("USD") ); _ = currency;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  if IsTrue(currency ) {
    *(request ).At(MkString("convert") )= OpCopy(currency );
  }
  return this.Call(MkString("publicGetGlobal"), request );
}

func (this *Coinmarketcap) ParseTicker(goArgs ...*Variant) *Variant{
  ticker := GoGetArg(goArgs, 0, MkUndefined()); _ = ticker;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.SafeTimestamp(ticker , MkString("last_updated") ); _ = timestamp;
  if IsTrue(OpEq2(timestamp , MkUndefined() )) {
    timestamp = this.Milliseconds();
  }
  change := this.SafeNumber(ticker , MkString("percent_change_24h") ); _ = change;
  last := OpCopy(MkUndefined() ); _ = last;
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  volume := OpCopy(MkUndefined() ); _ = volume;
  if IsTrue(OpNotEq2(market , MkUndefined() )) {
    symbol = *(market ).At(MkString("symbol") );
    priceKey := OpAdd(MkString("price_") , *(market ).At(MkString("quoteId") )); _ = priceKey;
    last = this.SafeNumber(ticker , priceKey );
    volumeKey := OpAdd(MkString("24h_volume_") , *(market ).At(MkString("quoteId") )); _ = volumeKey;
    volume = this.SafeNumber(ticker , volumeKey );
  }
  return MkMap(&VarMap{
        "symbol":symbol ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "high":MkUndefined() ,
        "low":MkUndefined() ,
        "bid":MkUndefined() ,
        "bidVolume":MkUndefined() ,
        "ask":MkUndefined() ,
        "askVolume":MkUndefined() ,
        "vwap":MkUndefined() ,
        "open":MkUndefined() ,
        "close":last ,
        "last":last ,
        "previousClose":MkUndefined() ,
        "change":MkUndefined() ,
        "percentage":change ,
        "average":MkUndefined() ,
        "baseVolume":MkUndefined() ,
        "quoteVolume":volume ,
        "info":ticker ,
        });
}

func (this *Coinmarketcap) FetchTickers(goArgs ...*Variant) *Variant{
  currency := GoGetArg(goArgs, 0, MkString("USD") ); _ = currency;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"limit":MkInteger(10000) }); _ = request;
  if IsTrue(currency ) {
    *(request ).At(MkString("convert") )= OpCopy(currency );
  }
  response := this.Call(MkString("publicGetTicker"), this.Extend(request , params )); _ = response;
  result := MkMap(&VarMap{}); _ = result;
  for t := MkInteger(0) ; IsTrue(OpLw(t , response.Length )); OpIncr(& t ){
    ticker := *(response ).At(t ); _ = ticker;
    currencyId := currency.ToLowerCase(); _ = currencyId;
    id := OpAdd(*(ticker ).At(MkString("id") ), OpAdd(MkString("/") , currencyId )); _ = id;
    symbol := OpCopy(id ); _ = symbol;
    market := OpCopy(MkUndefined() ); _ = market;
    if IsTrue(OpHasMember(id , *this.At(MkString("markets_by_id")) )) {
      market = *(*this.At(MkString("markets_by_id")) ).At(id );
      symbol = *(market ).At(MkString("symbol") );
    }
    *(result ).At(symbol )= this.ParseTicker(ticker , market );
  }
  return result ;
}

func (this *Coinmarketcap) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "convert":*(market ).At(MkString("quote") ),
        "id":*(market ).At(MkString("baseId") ),
        }); _ = request;
  response := this.Call(MkString("publicGetTickerId"), this.Extend(request , params )); _ = response;
  ticker := *(response ).At(MkInteger(0) ); _ = ticker;
  return this.ParseTicker(ticker , market );
}

func (this *Coinmarketcap) FetchCurrencies(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"limit":MkInteger(0) }); _ = request;
  response := this.Call(MkString("publicGetTicker"), this.Extend(request , params )); _ = response;
  result := MkMap(&VarMap{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , response.Length )); OpIncr(& i ){
    currency := *(response ).At(i ); _ = currency;
    id := this.SafeString(currency , MkString("symbol") ); _ = id;
    name := this.SafeString(currency , MkString("name") ); _ = name;
    code := this.CurrencyCode(id , name ); _ = code;
    *(result ).At(code )= MkMap(&VarMap{
        "id":id ,
        "code":code ,
        "info":currency ,
        "name":name ,
        "active":MkBool(true) ,
        "fee":MkUndefined() ,
        "precision":MkUndefined() ,
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
            "withdraw":MkMap(&VarMap{
                "min":MkUndefined() ,
                "max":MkUndefined() ,
                }),
            }),
        });
  }
  return result ;
}

func (this *Coinmarketcap) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("public") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  url := OpAdd(*(*(*this.At(MkString("urls")) ).At(MkString("api") )).At(api ), OpAdd(MkString("/") , OpAdd(*this.At(MkString("version")) , OpAdd(MkString("/") , this.ImplodeParams(path , params ))))); _ = url;
  query := this.Omit(params , this.ExtractParams(path )); _ = query;
  if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
    OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(query )));
  }
  return MkMap(&VarMap{
        "url":url ,
        "method":method ,
        "body":body ,
        "headers":headers ,
        });
}

func (this *Coinmarketcap) HandleErrors(goArgs ...*Variant) *Variant{
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
  error := this.SafeValue(response , MkString("error") , MkBool(false) ); _ = error;
  if IsTrue(error ) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , this.Json(response )))));
  }
  return MkUndefined()
}

