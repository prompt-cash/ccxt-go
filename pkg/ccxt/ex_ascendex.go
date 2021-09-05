package ccxt

import (
)

type Ascendex struct {
	*ExchangeBase
}

var _ Exchange = (*Ascendex)(nil)

func init() {
	exchange := &Ascendex{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Ascendex) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("ascendex") ,
            "name":MkString("AscendEX") ,
            "countries":MkArray(&VarArray{
                MkString("SG") ,
                }),
            "rateLimit":MkInteger(500) ,
            "certified":MkBool(true) ,
            "has":MkMap(&VarMap{
                "CORS":MkBool(false) ,
                "fetchMarkets":MkBool(true) ,
                "fetchCurrencies":MkBool(true) ,
                "fetchOrderBook":MkBool(true) ,
                "fetchTicker":MkBool(true) ,
                "fetchTickers":MkBool(true) ,
                "fetchOHLCV":MkBool(true) ,
                "fetchTrades":MkBool(true) ,
                "fetchAccounts":MkBool(true) ,
                "fetchBalance":MkBool(true) ,
                "createOrder":MkBool(true) ,
                "cancelOrder":MkBool(true) ,
                "cancelAllOrders":MkBool(true) ,
                "fetchDepositAddress":MkBool(true) ,
                "fetchTransactions":MkBool(true) ,
                "fetchDeposits":MkBool(true) ,
                "fetchWithdrawals":MkBool(true) ,
                "fetchOrder":MkBool(true) ,
                "fetchOrders":MkBool(true) ,
                "fetchOpenOrders":MkBool(true) ,
                "fetchClosedOrders":MkBool(true) ,
                }),
            "timeframes":MkMap(&VarMap{
                "1m":MkString("1") ,
                "5m":MkString("5") ,
                "15m":MkString("15") ,
                "30m":MkString("30") ,
                "1h":MkString("60") ,
                "2h":MkString("120") ,
                "4h":MkString("240") ,
                "6h":MkString("360") ,
                "12h":MkString("720") ,
                "1d":MkString("1d") ,
                "1w":MkString("1w") ,
                "1M":MkString("1m") ,
                }),
            "version":MkString("v1") ,
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/1294454/112027508-47984600-8b48-11eb-9e17-d26459cc36c6.jpg") ,
                "api":MkString("https://ascendex.com") ,
                "test":MkString("https://bitmax-test.io") ,
                "www":MkString("https://ascendex.com") ,
                "doc":MkArray(&VarArray{
                    MkString("https://bitmax-exchange.github.io/bitmax-pro-api/#bitmax-pro-api-documentation") ,
                    }),
                "fees":MkString("https://ascendex.com/en/feerate/transactionfee-traderate") ,
                "referral":MkMap(&VarMap{
                    "url":MkString("https://ascendex.com/en-us/register?inviteCode=EL6BXBQM") ,
                    "discount":MkNumber(0.25) ,
                    }),
                }),
            "api":MkMap(&VarMap{
                "public":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("assets") ,
                        MkString("products") ,
                        MkString("ticker") ,
                        MkString("barhist/info") ,
                        MkString("barhist") ,
                        MkString("depth") ,
                        MkString("trades") ,
                        MkString("cash/assets") ,
                        MkString("cash/products") ,
                        MkString("margin/assets") ,
                        MkString("margin/products") ,
                        MkString("futures/collateral") ,
                        MkString("futures/contracts") ,
                        MkString("futures/ref-px") ,
                        MkString("futures/market-data") ,
                        MkString("futures/funding-rates") ,
                        })}),
                "accountCategory":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("balance") ,
                        MkString("order/open") ,
                        MkString("order/status") ,
                        MkString("order/hist/current") ,
                        MkString("risk") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("order") ,
                        MkString("order/batch") ,
                        }),
                    "delete":MkArray(&VarArray{
                        MkString("order") ,
                        MkString("order/all") ,
                        MkString("order/batch") ,
                        }),
                    }),
                "accountGroup":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("cash/balance") ,
                        MkString("margin/balance") ,
                        MkString("margin/risk") ,
                        MkString("transfer") ,
                        MkString("futures/collateral-balance") ,
                        MkString("futures/position") ,
                        MkString("futures/risk") ,
                        MkString("futures/funding-payments") ,
                        MkString("order/hist") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("futures/transfer/deposit") ,
                        MkString("futures/transfer/withdraw") ,
                        }),
                    }),
                "private":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("info") ,
                        MkString("wallet/transactions") ,
                        MkString("wallet/deposit/address") ,
                        })}),
                }),
            "fees":MkMap(&VarMap{"trading":MkMap(&VarMap{
                    "feeSide":MkString("get") ,
                    "tierBased":MkBool(true) ,
                    "percentage":MkBool(true) ,
                    "taker":this.ParseNumber(MkString("0.002") ),
                    "maker":this.ParseNumber(MkString("0.002") ),
                    })}),
            "precisionMode":TICK_SIZE ,
            "options":MkMap(&VarMap{
                "account-category":MkString("cash") ,
                "account-group":MkUndefined() ,
                "fetchClosedOrders":MkMap(&VarMap{"method":MkString("accountGroupGetOrderHist") }),
                }),
            "exceptions":MkMap(&VarMap{
                "exact":MkMap(&VarMap{
                    "1900":BadRequest ,
                    "2100":AuthenticationError ,
                    "5002":BadSymbol ,
                    "6001":BadSymbol ,
                    "6010":InsufficientFunds ,
                    "60060":InvalidOrder ,
                    "600503":InvalidOrder ,
                    "100001":BadRequest ,
                    "100002":BadRequest ,
                    "100003":BadRequest ,
                    "100004":BadRequest ,
                    "100005":BadRequest ,
                    "100006":BadRequest ,
                    "100007":BadRequest ,
                    "100008":BadSymbol ,
                    "100009":AuthenticationError ,
                    "100010":BadRequest ,
                    "100011":BadRequest ,
                    "100012":BadRequest ,
                    "100013":BadRequest ,
                    "100101":ExchangeError ,
                    "150001":BadRequest ,
                    "200001":AuthenticationError ,
                    "200002":ExchangeError ,
                    "200003":ExchangeError ,
                    "200004":ExchangeError ,
                    "200005":ExchangeError ,
                    "200006":ExchangeError ,
                    "200007":ExchangeError ,
                    "200008":ExchangeError ,
                    "200009":ExchangeError ,
                    "200010":AuthenticationError ,
                    "200011":ExchangeError ,
                    "200012":ExchangeError ,
                    "200013":ExchangeError ,
                    "200014":PermissionDenied ,
                    "200015":PermissionDenied ,
                    "300001":InvalidOrder ,
                    "300002":InvalidOrder ,
                    "300003":InvalidOrder ,
                    "300004":InvalidOrder ,
                    "300005":InvalidOrder ,
                    "300006":InvalidOrder ,
                    "300007":InvalidOrder ,
                    "300008":InvalidOrder ,
                    "300009":InvalidOrder ,
                    "300011":InsufficientFunds ,
                    "300012":BadSymbol ,
                    "300013":InvalidOrder ,
                    "300020":InvalidOrder ,
                    "300021":InvalidOrder ,
                    "300031":InvalidOrder ,
                    "310001":InsufficientFunds ,
                    "310002":InvalidOrder ,
                    "310003":InvalidOrder ,
                    "310004":BadSymbol ,
                    "310005":InvalidOrder ,
                    "510001":ExchangeError ,
                    "900001":ExchangeError ,
                    }),
                "broad":MkMap(&VarMap{}),
                }),
            "commonCurrencies":MkMap(&VarMap{
                "BOND":MkString("BONDED") ,
                "BTCBEAR":MkString("BEAR") ,
                "BTCBULL":MkString("BULL") ,
                "BYN":MkString("Beyond Finance") ,
                }),
            }));
}

func (this *Ascendex) GetAccount(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  account := this.SafeValue(params , MkString("account") , *(*this.At(MkString("options")) ).At(MkString("account") )); _ = account;
  return (account.ToLowerCase()).Call(MkString("capitalize") );
}

func (this *Ascendex) FetchCurrencies(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  assets := this.Call(MkString("publicGetAssets"), params ); _ = assets;
  margin := this.Call(MkString("publicGetMarginAssets"), params ); _ = margin;
  cash := this.Call(MkString("publicGetCashAssets"), params ); _ = cash;
  assetsData := this.SafeValue(assets , MkString("data") , MkArray(&VarArray{})); _ = assetsData;
  marginData := this.SafeValue(margin , MkString("data") , MkArray(&VarArray{})); _ = marginData;
  cashData := this.SafeValue(cash , MkString("data") , MkArray(&VarArray{})); _ = cashData;
  assetsById := this.IndexBy(assetsData , MkString("assetCode") ); _ = assetsById;
  marginById := this.IndexBy(marginData , MkString("assetCode") ); _ = marginById;
  cashById := this.IndexBy(cashData , MkString("assetCode") ); _ = cashById;
  dataById := this.DeepExtend(assetsById , marginById , cashById ); _ = dataById;
  ids := GoGetKeys(dataById ); _ = ids;
  result := MkMap(&VarMap{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , ids.Length )); OpIncr(& i ){
    id := *(ids ).At(i ); _ = id;
    currency := *(dataById ).At(id ); _ = currency;
    code := this.SafeCurrencyCode(id ); _ = code;
    precision := this.SafeString2(currency , MkString("precisionScale") , MkString("nativeScale") ); _ = precision;
    minAmount := this.ParsePrecision(precision ); _ = minAmount;
    fee := this.SafeNumber2(currency , MkString("withdrawFee") , MkString("withdrawalFee") ); _ = fee;
    status := this.SafeString2(currency , MkString("status") , MkString("statusCode") ); _ = status;
    active := OpEq2(status , MkString("Normal") ); _ = active;
    margin := OpHasMember(MkString("borrowAssetCode") , currency ); _ = margin;
    *(result ).At(code )= MkMap(&VarMap{
        "id":id ,
        "code":code ,
        "info":currency ,
        "type":MkUndefined() ,
        "margin":margin ,
        "name":this.SafeString(currency , MkString("assetName") ),
        "active":active ,
        "fee":fee ,
        "precision":parseInt(precision ),
        "limits":MkMap(&VarMap{
            "amount":MkMap(&VarMap{
                "min":this.ParseNumber(minAmount ),
                "max":MkUndefined() ,
                }),
            "withdraw":MkMap(&VarMap{
                "min":this.SafeNumber(currency , MkString("minWithdrawalAmt") ),
                "max":MkUndefined() ,
                }),
            }),
        });
  }
  return result ;
}

func (this *Ascendex) FetchMarkets(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  products := this.Call(MkString("publicGetProducts"), params ); _ = products;
  cash := this.Call(MkString("publicGetCashProducts"), params ); _ = cash;
  futures := this.Call(MkString("publicGetFuturesContracts"), params ); _ = futures;
  productsData := this.SafeValue(products , MkString("data") , MkArray(&VarArray{})); _ = productsData;
  productsById := this.IndexBy(productsData , MkString("symbol") ); _ = productsById;
  cashData := this.SafeValue(cash , MkString("data") , MkArray(&VarArray{})); _ = cashData;
  futuresData := this.SafeValue(futures , MkString("data") , MkArray(&VarArray{})); _ = futuresData;
  cashAndFuturesData := this.ArrayConcat(cashData , futuresData ); _ = cashAndFuturesData;
  cashAndFuturesById := this.IndexBy(cashAndFuturesData , MkString("symbol") ); _ = cashAndFuturesById;
  dataById := this.DeepExtend(productsById , cashAndFuturesById ); _ = dataById;
  ids := GoGetKeys(dataById ); _ = ids;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , ids.Length )); OpIncr(& i ){
    id := *(ids ).At(i ); _ = id;
    market := *(dataById ).At(id ); _ = market;
    baseId := this.SafeString(market , MkString("baseAsset") ); _ = baseId;
    quoteId := this.SafeString(market , MkString("quoteAsset") ); _ = quoteId;
    base := this.SafeCurrencyCode(baseId ); _ = base;
    quote := this.SafeCurrencyCode(quoteId ); _ = quote;
    precision := MkMap(&VarMap{
          "amount":this.SafeNumber(market , MkString("lotSize") ),
          "price":this.SafeNumber(market , MkString("tickSize") ),
          }); _ = precision;
    status := this.SafeString(market , MkString("status") ); _ = status;
    active := OpEq2(status , MkString("Normal") ); _ = active;
    type_ := OpTrinary(OpHasMember(MkString("useLot") , market ), MkString("spot") , MkString("future") ); _ = type_;
    spot := OpEq2(type_ , MkString("spot") ); _ = spot;
    future := OpEq2(type_ , MkString("future") ); _ = future;
    margin := this.SafeValue(market , MkString("marginTradable") , MkBool(false) ); _ = margin;
    symbol := OpCopy(id ); _ = symbol;
    if IsTrue(OpNot(future )) {
      symbol = OpAdd(base , OpAdd(MkString("/") , quote ));
    }
    fee := this.SafeNumber(market , MkString("commissionReserveRate") ); _ = fee;
    result.Push(MkMap(&VarMap{
            "id":id ,
            "symbol":symbol ,
            "base":base ,
            "quote":quote ,
            "baseId":baseId ,
            "quoteId":quoteId ,
            "info":market ,
            "type":type_ ,
            "spot":spot ,
            "margin":margin ,
            "future":future ,
            "active":active ,
            "precision":precision ,
            "taker":fee ,
            "maker":fee ,
            "limits":MkMap(&VarMap{
                "amount":MkMap(&VarMap{
                    "min":this.SafeNumber(market , MkString("minQty") ),
                    "max":this.SafeNumber(market , MkString("maxQty") ),
                    }),
                "price":MkMap(&VarMap{
                    "min":this.SafeNumber(market , MkString("tickSize") ),
                    "max":MkUndefined() ,
                    }),
                "cost":MkMap(&VarMap{
                    "min":this.SafeNumber(market , MkString("minNotional") ),
                    "max":this.SafeNumber(market , MkString("maxNotional") ),
                    }),
                }),
            }));
  }
  return result ;
}

func (this *Ascendex) FetchAccounts(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  accountGroup := this.SafeString(*this.At(MkString("options")) , MkString("account-group") ); _ = accountGroup;
  response := OpCopy(MkUndefined() ); _ = response;
  if IsTrue(OpEq2(accountGroup , MkUndefined() )) {
    response = this.Call(MkString("privateGetInfo"), params );
    data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
    accountGroup = this.SafeString(data , MkString("accountGroup") );
    *(*this.At(MkString("options")) ).At(MkString("account-group") )= OpCopy(accountGroup );
  }
  return MkArray(&VarArray{
        MkMap(&VarMap{
            "id":accountGroup ,
            "type":MkUndefined() ,
            "currency":MkUndefined() ,
            "info":response ,
            }),
        });
}

func (this *Ascendex) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  this.LoadAccounts();
  defaultAccountCategory := this.SafeString(*this.At(MkString("options")) , MkString("account-category") , MkString("cash") ); _ = defaultAccountCategory;
  options := this.SafeValue(*this.At(MkString("options")) , MkString("fetchBalance") , MkMap(&VarMap{})); _ = options;
  accountCategory := this.SafeString(options , MkString("account-category") , defaultAccountCategory ); _ = accountCategory;
  accountCategory = this.SafeString(params , MkString("account-category") , accountCategory );
  params = this.Omit(params , MkString("account-category") );
  account := this.SafeValue(*this.At(MkString("accounts")) , MkInteger(0) , MkMap(&VarMap{})); _ = account;
  accountGroup := this.SafeString(account , MkString("id") ); _ = accountGroup;
  request := MkMap(&VarMap{"account-group":accountGroup }); _ = request;
  method := MkString("accountCategoryGetBalance") ; _ = method;
  if IsTrue(OpEq2(accountCategory , MkString("futures") )) {
    method = MkString("accountGroupGetFuturesCollateralBalance") ;
  } else {
    *(request ).At(MkString("account-category") )= OpCopy(accountCategory );
  }
  response := this.Call(method , this.Extend(request , params )); _ = response;
  result := MkMap(&VarMap{
        "info":response ,
        "timestamp":MkUndefined() ,
        "datetime":MkUndefined() ,
        }); _ = result;
  balances := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = balances;
  for i := MkInteger(0) ; IsTrue(OpLw(i , balances.Length )); OpIncr(& i ){
    balance := *(balances ).At(i ); _ = balance;
    code := this.SafeCurrencyCode(this.SafeString(balance , MkString("asset") )); _ = code;
    account := this.Account(); _ = account;
    *(account ).At(MkString("free") )= this.SafeString(balance , MkString("availableBalance") );
    *(account ).At(MkString("total") )= this.SafeString(balance , MkString("totalBalance") );
    *(result ).At(code )= OpCopy(account );
  }
  return this.ParseBalance(result );
}

func (this *Ascendex) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"symbol":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("publicGetDepth"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  orderbook := this.SafeValue(data , MkString("data") , MkMap(&VarMap{})); _ = orderbook;
  timestamp := this.SafeInteger(orderbook , MkString("ts") ); _ = timestamp;
  result := this.ParseOrderBook(orderbook , symbol , timestamp ); _ = result;
  *(result ).At(MkString("nonce") )= this.SafeInteger(orderbook , MkString("seqnum") );
  return result ;
}

func (this *Ascendex) ParseTicker(goArgs ...*Variant) *Variant{
  ticker := GoGetArg(goArgs, 0, MkUndefined()); _ = ticker;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := OpCopy(MkUndefined() ); _ = timestamp;
  marketId := this.SafeString(ticker , MkString("symbol") ); _ = marketId;
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  if IsTrue(OpHasMember(marketId , *this.At(MkString("markets_by_id")) )) {
    market = *(*this.At(MkString("markets_by_id")) ).At(marketId );
  } else {
    if IsTrue(OpNotEq2(marketId , MkUndefined() )) {
      type_ := this.SafeString(ticker , MkString("type") ); _ = type_;
      if IsTrue(OpEq2(type_ , MkString("spot") )) {
        baseId := MkUndefined(); _ = baseId;
        quoteId := MkUndefined(); _ = quoteId;
        ArrayUnpack(marketId.Split(MkString("/") ), &baseId, &quoteId);
        base := this.SafeCurrencyCode(baseId ); _ = base;
        quote := this.SafeCurrencyCode(quoteId ); _ = quote;
        symbol = OpAdd(base , OpAdd(MkString("/") , quote ));
      }
    }
  }
  if IsTrue(OpAnd(OpEq2(symbol , MkUndefined() ), OpNotEq2(market , MkUndefined() ))) {
    symbol = *(market ).At(MkString("symbol") );
  }
  close := this.SafeNumber(ticker , MkString("close") ); _ = close;
  bid := this.SafeValue(ticker , MkString("bid") , MkArray(&VarArray{})); _ = bid;
  ask := this.SafeValue(ticker , MkString("ask") , MkArray(&VarArray{})); _ = ask;
  open := this.SafeNumber(ticker , MkString("open") ); _ = open;
  change := OpCopy(MkUndefined() ); _ = change;
  percentage := OpCopy(MkUndefined() ); _ = percentage;
  average := OpCopy(MkUndefined() ); _ = average;
  if IsTrue(OpAnd(OpNotEq2(open , MkUndefined() ), OpNotEq2(close , MkUndefined() ))) {
    change = OpSub(close , open );
    if IsTrue(OpGt(open , MkInteger(0) )) {
      percentage = OpDiv(change , OpMulti(open , MkInteger(100) ));
    }
    average = OpDiv(this.Sum(open , close ), MkInteger(2) );
  }
  return MkMap(&VarMap{
        "symbol":symbol ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "high":this.SafeNumber(ticker , MkString("high") ),
        "low":this.SafeNumber(ticker , MkString("low") ),
        "bid":this.SafeNumber(bid , MkInteger(0) ),
        "bidVolume":this.SafeNumber(bid , MkInteger(1) ),
        "ask":this.SafeNumber(ask , MkInteger(0) ),
        "askVolume":this.SafeNumber(ask , MkInteger(1) ),
        "vwap":MkUndefined() ,
        "open":open ,
        "close":close ,
        "last":close ,
        "previousClose":MkUndefined() ,
        "change":change ,
        "percentage":percentage ,
        "average":average ,
        "baseVolume":this.SafeNumber(ticker , MkString("volume") ),
        "quoteVolume":MkUndefined() ,
        "info":ticker ,
        });
}

func (this *Ascendex) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"symbol":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("publicGetTicker"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  return this.ParseTicker(data , market );
}

func (this *Ascendex) FetchTickers(goArgs ...*Variant) *Variant{
  symbols := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbols;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  if IsTrue(OpNotEq2(symbols , MkUndefined() )) {
    marketIds := this.MarketIds(symbols ); _ = marketIds;
    *(request ).At(MkString("symbol") )= marketIds.Join(MkString(",") );
  }
  response := this.Call(MkString("publicGetTicker"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = data;
  return this.ParseTickers(data , symbols );
}

func (this *Ascendex) ParseOHLCV(goArgs ...*Variant) *Variant{
  ohlcv := GoGetArg(goArgs, 0, MkUndefined()); _ = ohlcv;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  data := this.SafeValue(ohlcv , MkString("data") , MkMap(&VarMap{})); _ = data;
  return MkArray(&VarArray{
        this.SafeInteger(data , MkString("ts") ),
        this.SafeNumber(data , MkString("o") ),
        this.SafeNumber(data , MkString("h") ),
        this.SafeNumber(data , MkString("l") ),
        this.SafeNumber(data , MkString("c") ),
        this.SafeNumber(data , MkString("v") ),
        });
}

func (this *Ascendex) FetchOHLCV(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  timeframe := GoGetArg(goArgs, 1, MkString("1m") ); _ = timeframe;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "symbol":*(market ).At(MkString("id") ),
        "interval":*(*this.At(MkString("timeframes")) ).At(timeframe ),
        }); _ = request;
  duration := this.ParseTimeframe(timeframe ); _ = duration;
  options := this.SafeValue(*this.At(MkString("options")) , MkString("fetchOHLCV") , MkMap(&VarMap{})); _ = options;
  defaultLimit := this.SafeInteger(options , MkString("limit") , MkInteger(500) ); _ = defaultLimit;
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("from") )= OpCopy(since );
    if IsTrue(OpEq2(limit , MkUndefined() )) {
      limit = OpCopy(defaultLimit );
    } else {
      limit = Math.Min(limit , defaultLimit );
    }
    *(request ).At(MkString("to") )= this.Sum(since , OpMulti(limit , OpMulti(duration , MkInteger(1000) )), MkInteger(1) );
  } else {
    if IsTrue(OpNotEq2(limit , MkUndefined() )) {
      *(request ).At(MkString("n") )= OpCopy(limit );
    }
  }
  response := this.Call(MkString("publicGetBarhist"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = data;
  return this.ParseOHLCVs(data , market , timeframe , since , limit );
}

func (this *Ascendex) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.SafeInteger(trade , MkString("ts") ); _ = timestamp;
  priceString := this.SafeString2(trade , MkString("price") , MkString("p") ); _ = priceString;
  amountString := this.SafeString(trade , MkString("q") ); _ = amountString;
  price := this.ParseNumber(priceString ); _ = price;
  amount := this.ParseNumber(amountString ); _ = amount;
  cost := this.ParseNumber(Precise.StringMul(priceString , amountString )); _ = cost;
  buyerIsMaker := this.SafeValue(trade , MkString("bm") , MkBool(false) ); _ = buyerIsMaker;
  makerOrTaker := OpTrinary(buyerIsMaker , MkString("maker") , MkString("taker") ); _ = makerOrTaker;
  side := OpTrinary(buyerIsMaker , MkString("buy") , MkString("sell") ); _ = side;
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  if IsTrue(OpAnd(OpEq2(symbol , MkUndefined() ), OpNotEq2(market , MkUndefined() ))) {
    symbol = *(market ).At(MkString("symbol") );
  }
  return MkMap(&VarMap{
        "info":trade ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "symbol":symbol ,
        "id":MkUndefined() ,
        "order":MkUndefined() ,
        "type":MkUndefined() ,
        "takerOrMaker":makerOrTaker ,
        "side":side ,
        "price":price ,
        "amount":amount ,
        "cost":cost ,
        "fee":MkUndefined() ,
        });
}

func (this *Ascendex) FetchTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"symbol":*(market ).At(MkString("id") )}); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("n") )= OpCopy(limit );
  }
  response := this.Call(MkString("publicGetTrades"), this.Extend(request , params )); _ = response;
  records := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = records;
  trades := this.SafeValue(records , MkString("data") , MkArray(&VarArray{})); _ = trades;
  return this.ParseTrades(trades , market , since , limit );
}

func (this *Ascendex) ParseOrderStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "PendingNew":MkString("open") ,
        "New":MkString("open") ,
        "PartiallyFilled":MkString("open") ,
        "Filled":MkString("closed") ,
        "Canceled":MkString("canceled") ,
        "Rejected":MkString("rejected") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Ascendex) ParseOrder(goArgs ...*Variant) *Variant{
  order := GoGetArg(goArgs, 0, MkUndefined()); _ = order;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  status := this.ParseOrderStatus(this.SafeString(order , MkString("status") )); _ = status;
  marketId := this.SafeString(order , MkString("symbol") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market , MkString("/") ); _ = symbol;
  timestamp := this.SafeInteger2(order , MkString("timestamp") , MkString("sendingTime") ); _ = timestamp;
  lastTradeTimestamp := this.SafeInteger(order , MkString("lastExecTime") ); _ = lastTradeTimestamp;
  price := this.SafeNumber(order , MkString("price") ); _ = price;
  amount := this.SafeNumber(order , MkString("orderQty") ); _ = amount;
  average := this.SafeNumber(order , MkString("avgPx") ); _ = average;
  filled := this.SafeNumber2(order , MkString("cumFilledQty") , MkString("cumQty") ); _ = filled;
  id := this.SafeString(order , MkString("orderId") ); _ = id;
  clientOrderId := this.SafeString(order , MkString("id") ); _ = clientOrderId;
  if IsTrue(OpNotEq2(clientOrderId , MkUndefined() )) {
    if IsTrue(OpLw(clientOrderId.Length , MkInteger(1) )) {
      clientOrderId = OpCopy(MkUndefined() );
    }
  }
  type_ := this.SafeStringLower(order , MkString("orderType") ); _ = type_;
  side := this.SafeStringLower(order , MkString("side") ); _ = side;
  feeCost := this.SafeNumber(order , MkString("cumFee") ); _ = feeCost;
  fee := OpCopy(MkUndefined() ); _ = fee;
  if IsTrue(OpNotEq2(feeCost , MkUndefined() )) {
    feeCurrencyId := this.SafeString(order , MkString("feeAsset") ); _ = feeCurrencyId;
    feeCurrencyCode := this.SafeCurrencyCode(feeCurrencyId ); _ = feeCurrencyCode;
    fee = MkMap(&VarMap{
        "cost":feeCost ,
        "currency":feeCurrencyCode ,
        });
  }
  stopPrice := this.SafeNumber(order , MkString("stopPrice") ); _ = stopPrice;
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
            "stopPrice":stopPrice ,
            "amount":amount ,
            "cost":MkUndefined() ,
            "average":average ,
            "filled":filled ,
            "remaining":MkUndefined() ,
            "status":status ,
            "fee":fee ,
            "trades":MkUndefined() ,
            }));
}

func (this *Ascendex) CreateOrder(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 2, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 3, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 4, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 5, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  this.LoadAccounts();
  market := this.Market(symbol ); _ = market;
  defaultAccountCategory := this.SafeString(*this.At(MkString("options")) , MkString("account-category") , MkString("cash") ); _ = defaultAccountCategory;
  options := this.SafeValue(*this.At(MkString("options")) , MkString("createOrder") , MkMap(&VarMap{})); _ = options;
  accountCategory := this.SafeString(options , MkString("account-category") , defaultAccountCategory ); _ = accountCategory;
  accountCategory = this.SafeString(params , MkString("account-category") , accountCategory );
  params = this.Omit(params , MkString("account-category") );
  account := this.SafeValue(*this.At(MkString("accounts")) , MkInteger(0) , MkMap(&VarMap{})); _ = account;
  accountGroup := this.SafeValue(account , MkString("id") ); _ = accountGroup;
  clientOrderId := this.SafeString2(params , MkString("clientOrderId") , MkString("id") ); _ = clientOrderId;
  request := MkMap(&VarMap{
        "account-group":accountGroup ,
        "account-category":accountCategory ,
        "symbol":*(market ).At(MkString("id") ),
        "time":this.Milliseconds(),
        "orderQty":this.AmountToPrecision(symbol , amount ),
        "orderType":type_ ,
        "side":side ,
        }); _ = request;
  if IsTrue(OpNotEq2(clientOrderId , MkUndefined() )) {
    *(request ).At(MkString("id") )= OpCopy(clientOrderId );
    params = this.Omit(params , MkArray(&VarArray{
            MkString("clientOrderId") ,
            MkString("id") ,
            }));
  }
  if IsTrue(OpOr(OpEq2(type_ , MkString("limit") ), OpEq2(type_ , MkString("stop_limit") ))) {
    *(request ).At(MkString("orderPrice") )= this.PriceToPrecision(symbol , price );
  }
  if IsTrue(OpOr(OpEq2(type_ , MkString("stop_limit") ), OpEq2(type_ , MkString("stop_market") ))) {
    stopPrice := this.SafeNumber(params , MkString("stopPrice") ); _ = stopPrice;
    if IsTrue(OpEq2(stopPrice , MkUndefined() )) {
      panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" createOrder() requires a stopPrice parameter for ") , OpAdd(type_ , MkString(" orders") )))));
    } else {
      *(request ).At(MkString("stopPrice") )= this.PriceToPrecision(symbol , stopPrice );
      params = this.Omit(params , MkString("stopPrice") );
    }
  }
  response := this.Call(MkString("accountCategoryPostOrder"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  info := this.SafeValue(data , MkString("info") , MkMap(&VarMap{})); _ = info;
  return this.ParseOrder(info , market );
}

func (this *Ascendex) FetchOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  this.LoadAccounts();
  defaultAccountCategory := this.SafeString(*this.At(MkString("options")) , MkString("account-category") , MkString("cash") ); _ = defaultAccountCategory;
  options := this.SafeValue(*this.At(MkString("options")) , MkString("fetchOrder") , MkMap(&VarMap{})); _ = options;
  accountCategory := this.SafeString(options , MkString("account-category") , defaultAccountCategory ); _ = accountCategory;
  accountCategory = this.SafeString(params , MkString("account-category") , accountCategory );
  params = this.Omit(params , MkString("account-category") );
  account := this.SafeValue(*this.At(MkString("accounts")) , MkInteger(0) , MkMap(&VarMap{})); _ = account;
  accountGroup := this.SafeValue(account , MkString("id") ); _ = accountGroup;
  request := MkMap(&VarMap{
        "account-group":accountGroup ,
        "account-category":accountCategory ,
        "orderId":id ,
        }); _ = request;
  response := this.Call(MkString("accountCategoryGetOrderStatus"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  return this.ParseOrder(data );
}

func (this *Ascendex) FetchOpenOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  this.LoadAccounts();
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
  }
  defaultAccountCategory := this.SafeString(*this.At(MkString("options")) , MkString("account-category") , MkString("cash") ); _ = defaultAccountCategory;
  options := this.SafeValue(*this.At(MkString("options")) , MkString("fetchOpenOrders") , MkMap(&VarMap{})); _ = options;
  accountCategory := this.SafeString(options , MkString("account-category") , defaultAccountCategory ); _ = accountCategory;
  accountCategory = this.SafeString(params , MkString("account-category") , accountCategory );
  params = this.Omit(params , MkString("account-category") );
  account := this.SafeValue(*this.At(MkString("accounts")) , MkInteger(0) , MkMap(&VarMap{})); _ = account;
  accountGroup := this.SafeValue(account , MkString("id") ); _ = accountGroup;
  request := MkMap(&VarMap{
        "account-group":accountGroup ,
        "account-category":accountCategory ,
        }); _ = request;
  response := this.Call(MkString("accountCategoryGetOrderOpen"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = data;
  if IsTrue(OpEq2(accountCategory , MkString("futures") )) {
    return this.ParseOrders(data , market , since , limit );
  }
  orders := MkArray(&VarArray{}); _ = orders;
  for i := MkInteger(0) ; IsTrue(OpLw(i , data.Length )); OpIncr(& i ){
    order := this.ParseOrder(*(data ).At(i ), market ); _ = order;
    orders.Push(order );
  }
  return this.FilterBySymbolSinceLimit(orders , symbol , since , limit );
}

func (this *Ascendex) FetchClosedOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  this.LoadAccounts();
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
  }
  defaultAccountCategory := this.SafeString(*this.At(MkString("options")) , MkString("account-category") ); _ = defaultAccountCategory;
  options := this.SafeValue(*this.At(MkString("options")) , MkString("fetchClosedOrders") , MkMap(&VarMap{})); _ = options;
  accountCategory := this.SafeString(options , MkString("account-category") , defaultAccountCategory ); _ = accountCategory;
  accountCategory = this.SafeString(params , MkString("account-category") , accountCategory );
  params = this.Omit(params , MkString("account-category") );
  account := this.SafeValue(*this.At(MkString("accounts")) , MkInteger(0) , MkMap(&VarMap{})); _ = account;
  accountGroup := this.SafeValue(account , MkString("id") ); _ = accountGroup;
  request := MkMap(&VarMap{"account-group":accountGroup }); _ = request;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("symbol") )= *(market ).At(MkString("id") );
  }
  method := this.SafeValue(options , MkString("method") , MkString("accountGroupGetOrderHist") ); _ = method;
  if IsTrue(OpEq2(method , MkString("accountGroupGetOrderHist") )) {
    if IsTrue(OpNotEq2(accountCategory , MkUndefined() )) {
      *(request ).At(MkString("category") )= OpCopy(accountCategory );
    }
  } else {
    *(request ).At(MkString("account-category") )= OpCopy(accountCategory );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("startTime") )= OpCopy(since );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("pageSize") )= OpCopy(limit );
  }
  response := this.Call(method , this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") ); _ = data;
  isArray := GoIsArray(data ); _ = isArray;
  if IsTrue(OpNot(isArray )) {
    data = this.SafeValue(data , MkString("data") , MkArray(&VarArray{}));
  }
  return this.ParseOrders(data , market , since , limit );
}

func (this *Ascendex) CancelOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" cancelOrder() requires a symbol argument") )));
  }
  this.LoadMarkets();
  this.LoadAccounts();
  market := this.Market(symbol ); _ = market;
  defaultAccountCategory := this.SafeString(*this.At(MkString("options")) , MkString("account-category") , MkString("cash") ); _ = defaultAccountCategory;
  options := this.SafeValue(*this.At(MkString("options")) , MkString("cancelOrder") , MkMap(&VarMap{})); _ = options;
  accountCategory := this.SafeString(options , MkString("account-category") , defaultAccountCategory ); _ = accountCategory;
  accountCategory = this.SafeString(params , MkString("account-category") , accountCategory );
  params = this.Omit(params , MkString("account-category") );
  account := this.SafeValue(*this.At(MkString("accounts")) , MkInteger(0) , MkMap(&VarMap{})); _ = account;
  accountGroup := this.SafeValue(account , MkString("id") ); _ = accountGroup;
  clientOrderId := this.SafeString2(params , MkString("clientOrderId") , MkString("id") ); _ = clientOrderId;
  request := MkMap(&VarMap{
        "account-group":accountGroup ,
        "account-category":accountCategory ,
        "symbol":*(market ).At(MkString("id") ),
        "time":this.Milliseconds(),
        "id":MkString("foobar") ,
        }); _ = request;
  if IsTrue(OpEq2(clientOrderId , MkUndefined() )) {
    *(request ).At(MkString("orderId") )= OpCopy(id );
  } else {
    *(request ).At(MkString("id") )= OpCopy(clientOrderId );
    params = this.Omit(params , MkArray(&VarArray{
            MkString("clientOrderId") ,
            MkString("id") ,
            }));
  }
  response := this.Call(MkString("accountCategoryDeleteOrder"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  info := this.SafeValue(data , MkString("info") , MkMap(&VarMap{})); _ = info;
  return this.ParseOrder(info , market );
}

func (this *Ascendex) CancelAllOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  this.LoadAccounts();
  defaultAccountCategory := this.SafeString(*this.At(MkString("options")) , MkString("account-category") , MkString("cash") ); _ = defaultAccountCategory;
  options := this.SafeValue(*this.At(MkString("options")) , MkString("cancelAllOrders") , MkMap(&VarMap{})); _ = options;
  accountCategory := this.SafeString(options , MkString("account-category") , defaultAccountCategory ); _ = accountCategory;
  accountCategory = this.SafeString(params , MkString("account-category") , accountCategory );
  params = this.Omit(params , MkString("account-category") );
  account := this.SafeValue(*this.At(MkString("accounts")) , MkInteger(0) , MkMap(&VarMap{})); _ = account;
  accountGroup := this.SafeValue(account , MkString("id") ); _ = accountGroup;
  request := MkMap(&VarMap{
        "account-group":accountGroup ,
        "account-category":accountCategory ,
        "time":this.Milliseconds(),
        }); _ = request;
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("symbol") )= *(market ).At(MkString("id") );
  }
  response := this.Call(MkString("accountCategoryDeleteOrderAll"), this.Extend(request , params )); _ = response;
  return response ;
}

func (this *Ascendex) ParseDepositAddress(goArgs ...*Variant) *Variant{
  depositAddress := GoGetArg(goArgs, 0, MkUndefined()); _ = depositAddress;
  currency := GoGetArg(goArgs, 1, MkUndefined() ); _ = currency;
  address := this.SafeString(depositAddress , MkString("address") ); _ = address;
  tagId := this.SafeString(depositAddress , MkString("tagId") ); _ = tagId;
  tag := this.SafeString(depositAddress , tagId ); _ = tag;
  this.CheckAddress(address );
  code := OpTrinary(OpEq2(currency , MkUndefined() ), MkUndefined() , *(currency ).At(MkString("code") )); _ = code;
  return MkMap(&VarMap{
        "currency":code ,
        "address":address ,
        "tag":tag ,
        "info":depositAddress ,
        });
}

func (this *Ascendex) FetchDepositAddress(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  currency := this.Currency(code ); _ = currency;
  chainName := this.SafeString(params , MkString("chainName") ); _ = chainName;
  params = this.Omit(params , MkString("chainName") );
  request := MkMap(&VarMap{"asset":*(currency ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("privateGetWalletDepositAddress"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  addresses := this.SafeValue(data , MkString("address") , MkArray(&VarArray{})); _ = addresses;
  numAddresses := OpCopy(addresses.Length ); _ = numAddresses;
  address := OpCopy(MkUndefined() ); _ = address;
  if IsTrue(OpGt(numAddresses , MkInteger(1) )) {
    addressesByChainName := this.IndexBy(addresses , MkString("chainName") ); _ = addressesByChainName;
    if IsTrue(OpEq2(chainName , MkUndefined() )) {
      chainNames := GoGetKeys(addressesByChainName ); _ = chainNames;
      chains := chainNames.Join(MkString(", ") ); _ = chains;
      panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" fetchDepositAddress returned more than one address, a chainName parameter is required, one of ") , chains ))));
    }
    address = this.SafeValue(addressesByChainName , chainName , MkMap(&VarMap{}));
  } else {
    address = this.SafeValue(addresses , MkInteger(0) , MkMap(&VarMap{}));
  }
  result := this.ParseDepositAddress(address , currency ); _ = result;
  return this.Extend(result , MkMap(&VarMap{"info":response }));
}

func (this *Ascendex) FetchDeposits(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"txType":MkString("deposit") }); _ = request;
  return this.FetchTransactions(code , since , limit , this.Extend(request , params ));
}

func (this *Ascendex) FetchWithdrawals(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"txType":MkString("withdrawal") }); _ = request;
  return this.FetchTransactions(code , since , limit , this.Extend(request , params ));
}

func (this *Ascendex) FetchTransactions(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  currency := OpCopy(MkUndefined() ); _ = currency;
  if IsTrue(OpNotEq2(code , MkUndefined() )) {
    currency = this.Currency(code );
    *(request ).At(MkString("asset") )= *(currency ).At(MkString("id") );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("startTs") )= OpCopy(since );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("pageSize") )= OpCopy(limit );
  }
  response := this.Call(MkString("privateGetWalletTransactions"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  transactions := this.SafeValue(data , MkString("data") , MkArray(&VarArray{})); _ = transactions;
  return this.ParseTransactions(transactions , currency , since , limit );
}

func (this *Ascendex) ParseTransactionStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "reviewing":MkString("pending") ,
        "pending":MkString("pending") ,
        "confirmed":MkString("ok") ,
        "rejected":MkString("rejected") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Ascendex) ParseTransaction(goArgs ...*Variant) *Variant{
  transaction := GoGetArg(goArgs, 0, MkUndefined()); _ = transaction;
  currency := GoGetArg(goArgs, 1, MkUndefined() ); _ = currency;
  id := this.SafeString(transaction , MkString("requestId") ); _ = id;
  amount := this.SafeNumber(transaction , MkString("amount") ); _ = amount;
  destAddress := this.SafeValue(transaction , MkString("destAddress") , MkMap(&VarMap{})); _ = destAddress;
  address := this.SafeString(destAddress , MkString("address") ); _ = address;
  tag := this.SafeString(destAddress , MkString("destTag") ); _ = tag;
  txid := this.SafeString(transaction , MkString("networkTransactionId") ); _ = txid;
  type_ := this.SafeString(transaction , MkString("transactionType") ); _ = type_;
  timestamp := this.SafeInteger(transaction , MkString("time") ); _ = timestamp;
  currencyId := this.SafeString(transaction , MkString("asset") ); _ = currencyId;
  code := this.SafeCurrencyCode(currencyId , currency ); _ = code;
  status := this.ParseTransactionStatus(this.SafeString(transaction , MkString("status") )); _ = status;
  feeCost := this.SafeNumber(transaction , MkString("commission") ); _ = feeCost;
  return MkMap(&VarMap{
        "info":transaction ,
        "id":id ,
        "currency":code ,
        "amount":amount ,
        "address":address ,
        "addressTo":address ,
        "addressFrom":MkUndefined() ,
        "tag":tag ,
        "tagTo":tag ,
        "tagFrom":MkUndefined() ,
        "status":status ,
        "type":type_ ,
        "updated":MkUndefined() ,
        "txid":txid ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "fee":MkMap(&VarMap{
            "currency":code ,
            "cost":feeCost ,
            }),
        });
}

func (this *Ascendex) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("public") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  url := MkString("") ; _ = url;
  query := OpCopy(params ); _ = query;
  accountCategory := OpEq2(api , MkString("accountCategory") ); _ = accountCategory;
  if IsTrue(OpOr(accountCategory , OpEq2(api , MkString("accountGroup") ))) {
    OpAddAssign(& url , this.ImplodeParams(MkString("/{account-group}") , params ));
    query = this.Omit(params , MkString("account-group") );
  }
  request := this.ImplodeParams(path , query ); _ = request;
  OpAddAssign(& url , OpAdd(MkString("/api/pro/") , *this.At(MkString("version")) ));
  if IsTrue(accountCategory ) {
    OpAddAssign(& url , this.ImplodeParams(MkString("/{account-category}") , query ));
    query = this.Omit(query , MkString("account-category") );
  }
  OpAddAssign(& url , OpAdd(MkString("/") , request ));
  query = this.Omit(query , this.ExtractParams(path ));
  if IsTrue(OpEq2(api , MkString("public") )) {
    if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
      OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(query )));
    }
  } else {
    this.CheckRequiredCredentials();
    timestamp := (this.Milliseconds()).Call(MkString("toString") ); _ = timestamp;
    payload := OpAdd(timestamp , OpAdd(MkString("+") , request )); _ = payload;
    hmac := this.Hmac(this.Encode(payload ), this.Encode(*this.At(MkString("secret")) ), MkString("sha256") , MkString("base64") ); _ = hmac;
    headers = MkMap(&VarMap{
        "x-auth-key":*this.At(MkString("apiKey")) ,
        "x-auth-timestamp":timestamp ,
        "x-auth-signature":hmac ,
        });
    if IsTrue(OpEq2(method , MkString("GET") )) {
      if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
        OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(query )));
      }
    } else {
      *(headers ).At(MkString("Content-Type") )= MkString("application/json") ;
      body = this.Json(query );
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

func (this *Ascendex) HandleErrors(goArgs ...*Variant) *Variant{
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
  code := this.SafeString(response , MkString("code") ); _ = code;
  message := this.SafeString(response , MkString("message") ); _ = message;
  error := OpAnd(OpNotEq2(code , MkUndefined() ), OpNotEq2(code , MkString("0") )); _ = error;
  if IsTrue(OpOr(error , OpNotEq2(message , MkUndefined() ))) {
    feedback := OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , body )); _ = feedback;
    this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("exact") ), code , feedback );
    this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("exact") ), message , feedback );
    this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("broad") ), message , feedback );
    panic(NewExchangeError(feedback ));
  }
  return MkUndefined()
}

