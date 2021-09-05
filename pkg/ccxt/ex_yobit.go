package ccxt

import (
)

type Yobit struct {
	*ExchangeBase
}

var _ Exchange = (*Yobit)(nil)

func init() {
	exchange := &Yobit{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Yobit) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("yobit") ,
            "name":MkString("YoBit") ,
            "countries":MkArray(&VarArray{
                MkString("RU") ,
                }),
            "rateLimit":MkInteger(3000) ,
            "version":MkString("3") ,
            "has":MkMap(&VarMap{
                "cancelOrder":MkBool(true) ,
                "CORS":MkBool(false) ,
                "createDepositAddress":MkBool(true) ,
                "createMarketOrder":MkBool(false) ,
                "createOrder":MkBool(true) ,
                "fetchBalance":MkBool(true) ,
                "fetchDepositAddress":MkBool(true) ,
                "fetchDeposits":MkBool(false) ,
                "fetchMarkets":MkBool(true) ,
                "fetchMyTrades":MkBool(true) ,
                "fetchOpenOrders":MkBool(true) ,
                "fetchOrder":MkBool(true) ,
                "fetchOrderBook":MkBool(true) ,
                "fetchOrderBooks":MkBool(true) ,
                "fetchTicker":MkBool(true) ,
                "fetchTickers":MkBool(true) ,
                "fetchTrades":MkBool(true) ,
                "fetchTransactions":MkBool(false) ,
                "fetchWithdrawals":MkBool(false) ,
                "withdraw":MkBool(true) ,
                }),
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/1294454/27766910-cdcbfdae-5eea-11e7-9859-03fea873272d.jpg") ,
                "api":MkMap(&VarMap{
                    "public":MkString("https://yobit.net/api") ,
                    "private":MkString("https://yobit.net/tapi") ,
                    }),
                "www":MkString("https://www.yobit.net") ,
                "doc":MkString("https://www.yobit.net/en/api/") ,
                "fees":MkString("https://www.yobit.net/en/fees/") ,
                }),
            "api":MkMap(&VarMap{
                "public":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("depth/{pair}") ,
                        MkString("info") ,
                        MkString("ticker/{pair}") ,
                        MkString("trades/{pair}") ,
                        })}),
                "private":MkMap(&VarMap{"post":MkArray(&VarArray{
                        MkString("ActiveOrders") ,
                        MkString("CancelOrder") ,
                        MkString("GetDepositAddress") ,
                        MkString("getInfo") ,
                        MkString("OrderInfo") ,
                        MkString("Trade") ,
                        MkString("TradeHistory") ,
                        MkString("WithdrawCoinsToAddress") ,
                        })}),
                }),
            "fees":MkMap(&VarMap{
                "trading":MkMap(&VarMap{
                    "maker":MkNumber(0.002) ,
                    "taker":MkNumber(0.002) ,
                    }),
                "funding":MkMap(&VarMap{"withdraw":MkMap(&VarMap{})}),
                }),
            "commonCurrencies":MkMap(&VarMap{
                "AIR":MkString("AirCoin") ,
                "ANI":MkString("ANICoin") ,
                "ANT":MkString("AntsCoin") ,
                "ATMCHA":MkString("ATM") ,
                "ASN":MkString("Ascension") ,
                "AST":MkString("Astral") ,
                "ATM":MkString("Autumncoin") ,
                "BCC":MkString("BCH") ,
                "BCS":MkString("BitcoinStake") ,
                "BITS":MkString("Bitstar") ,
                "BLN":MkString("Bulleon") ,
                "BNS":MkString("Benefit Bonus Coin") ,
                "BOT":MkString("BOTcoin") ,
                "BON":MkString("BONES") ,
                "BPC":MkString("BitcoinPremium") ,
                "BST":MkString("BitStone") ,
                "BTS":MkString("Bitshares2") ,
                "CAT":MkString("BitClave") ,
                "CBC":MkString("CryptoBossCoin") ,
                "CMT":MkString("CometCoin") ,
                "COIN":MkString("Coin.com") ,
                "COV":MkString("Coven Coin") ,
                "COVX":MkString("COV") ,
                "CPC":MkString("Capricoin") ,
                "CREDIT":MkString("Creditbit") ,
                "CS":MkString("CryptoSpots") ,
                "DCT":MkString("Discount") ,
                "DFT":MkString("DraftCoin") ,
                "DGD":MkString("DarkGoldCoin") ,
                "DIRT":MkString("DIRTY") ,
                "DROP":MkString("FaucetCoin") ,
                "DSH":MkString("DASH") ,
                "EKO":MkString("EkoCoin") ,
                "ENTER":MkString("ENTRC") ,
                "EPC":MkString("ExperienceCoin") ,
                "ESC":MkString("EdwardSnowden") ,
                "EUROPE":MkString("EUROP") ,
                "EXT":MkString("LifeExtension") ,
                "FUND":MkString("FUNDChains") ,
                "FUNK":MkString("FUNKCoin") ,
                "GCC":MkString("GlobalCryptocurrency") ,
                "GEN":MkString("Genstake") ,
                "GENE":MkString("Genesiscoin") ,
                "GOLD":MkString("GoldMint") ,
                "GOT":MkString("Giotto Coin") ,
                "GSX":MkString("GlowShares") ,
                "GT":MkString("GTcoin") ,
                "HTML5":MkString("HTML") ,
                "HYPERX":MkString("HYPER") ,
                "ICN":MkString("iCoin") ,
                "INSANE":MkString("INSN") ,
                "JNT":MkString("JointCoin") ,
                "JPC":MkString("JupiterCoin") ,
                "JWL":MkString("Jewels") ,
                "KNC":MkString("KingN Coin") ,
                "LBTCX":MkString("LiteBitcoin") ,
                "LIZI":MkString("LiZi") ,
                "LOC":MkString("LocoCoin") ,
                "LOCX":MkString("LOC") ,
                "LUNYR":MkString("LUN") ,
                "LUN":MkString("LunarCoin") ,
                "LUNA":MkString("Luna Coin") ,
                "MASK":MkString("Yobit MASK") ,
                "MDT":MkString("Midnight") ,
                "MIS":MkString("MIScoin") ,
                "NAV":MkString("NavajoCoin") ,
                "NBT":MkString("NiceBytes") ,
                "OMG":MkString("OMGame") ,
                "ONX":MkString("Onix") ,
                "PAC":MkString("$PAC") ,
                "PLAY":MkString("PlayCoin") ,
                "PIVX":MkString("Darknet") ,
                "PRS":MkString("PRE") ,
                "PUTIN":MkString("PutinCoin") ,
                "STK":MkString("StakeCoin") ,
                "SUB":MkString("Subscriptio") ,
                "PAY":MkString("EPAY") ,
                "PLC":MkString("Platin Coin") ,
                "RAI":MkString("RaiderCoin") ,
                "RCN":MkString("RCoin") ,
                "REP":MkString("Republicoin") ,
                "RUR":MkString("RUB") ,
                "SBTC":MkString("Super Bitcoin") ,
                "TTC":MkString("TittieCoin") ,
                "UNI":MkString("Universe") ,
                "UST":MkString("Uservice") ,
                "VOL":MkString("VolumeCoin") ,
                "XIN":MkString("XINCoin") ,
                "XRA":MkString("Ratecoin") ,
                }),
            "options":MkMap(&VarMap{
                "fetchOrdersRequiresSymbol":MkBool(true) ,
                "fetchTickersMaxLength":MkInteger(512) ,
                }),
            "exceptions":MkMap(&VarMap{
                "exact":MkMap(&VarMap{
                    "803":InvalidOrder ,
                    "804":InvalidOrder ,
                    "805":InvalidOrder ,
                    "806":InvalidOrder ,
                    "807":InvalidOrder ,
                    "831":InsufficientFunds ,
                    "832":InsufficientFunds ,
                    "833":OrderNotFound ,
                    }),
                "broad":MkMap(&VarMap{
                    "Invalid pair name":ExchangeError ,
                    "invalid api key":AuthenticationError ,
                    "invalid sign":AuthenticationError ,
                    "api key dont have trade permission":AuthenticationError ,
                    "invalid parameter":InvalidOrder ,
                    "invalid order":InvalidOrder ,
                    "The given order has already been cancelled":InvalidOrder ,
                    "Requests too often":DDoSProtection ,
                    "not available":ExchangeNotAvailable ,
                    "data unavailable":ExchangeNotAvailable ,
                    "external service unavailable":ExchangeNotAvailable ,
                    "Total transaction amount":InvalidOrder ,
                    "The given order has already been closed and cannot be cancelled":InvalidOrder ,
                    "Insufficient funds":InsufficientFunds ,
                    "invalid key":AuthenticationError ,
                    "invalid nonce":InvalidNonce ,
                    "Total order amount is less than minimal amount":InvalidOrder ,
                    "Rate Limited":RateLimitExceeded ,
                    }),
                }),
            "orders":MkMap(&VarMap{}),
            }));
}

func (this *Yobit) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("privatePostGetInfo"), params ); _ = response;
  balances := this.SafeValue(response , MkString("return") , MkMap(&VarMap{})); _ = balances;
  timestamp := this.SafeInteger(balances , MkString("server_time") ); _ = timestamp;
  result := MkMap(&VarMap{
        "info":response ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        }); _ = result;
  free := this.SafeValue(balances , MkString("funds") , MkMap(&VarMap{})); _ = free;
  total := this.SafeValue(balances , MkString("funds_incl_orders") , MkMap(&VarMap{})); _ = total;
  currencyIds := GoGetKeys(this.Extend(free , total )); _ = currencyIds;
  for i := MkInteger(0) ; IsTrue(OpLw(i , currencyIds.Length )); OpIncr(& i ){
    currencyId := *(currencyIds ).At(i ); _ = currencyId;
    code := this.SafeCurrencyCode(currencyId ); _ = code;
    account := this.Account(); _ = account;
    *(account ).At(MkString("free") )= this.SafeString(free , currencyId );
    *(account ).At(MkString("total") )= this.SafeString(total , currencyId );
    *(result ).At(code )= OpCopy(account );
  }
  return this.ParseBalance(result );
}

func (this *Yobit) FetchMarkets(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetInfo"), params ); _ = response;
  markets := this.SafeValue(response , MkString("pairs") ); _ = markets;
  keys := GoGetKeys(markets ); _ = keys;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , keys.Length )); OpIncr(& i ){
    id := *(keys ).At(i ); _ = id;
    market := *(markets ).At(id ); _ = market;
    baseId := MkUndefined(); _ = baseId;
    quoteId := MkUndefined(); _ = quoteId;
    ArrayUnpack(id.Split(MkString("_") ), &baseId, &quoteId);
    base := baseId.ToUpperCase(); _ = base;
    quote := quoteId.ToUpperCase(); _ = quote;
    base = this.SafeCurrencyCode(base );
    quote = this.SafeCurrencyCode(quote );
    symbol := OpAdd(base , OpAdd(MkString("/") , quote )); _ = symbol;
    precision := MkMap(&VarMap{
          "amount":this.SafeInteger(market , MkString("decimal_places") ),
          "price":this.SafeInteger(market , MkString("decimal_places") ),
          }); _ = precision;
    amountLimits := MkMap(&VarMap{
          "min":this.SafeNumber(market , MkString("min_amount") ),
          "max":this.SafeNumber(market , MkString("max_amount") ),
          }); _ = amountLimits;
    priceLimits := MkMap(&VarMap{
          "min":this.SafeNumber(market , MkString("min_price") ),
          "max":this.SafeNumber(market , MkString("max_price") ),
          }); _ = priceLimits;
    costLimits := MkMap(&VarMap{"min":this.SafeNumber(market , MkString("min_total") )}); _ = costLimits;
    limits := MkMap(&VarMap{
          "amount":amountLimits ,
          "price":priceLimits ,
          "cost":costLimits ,
          }); _ = limits;
    hidden := this.SafeInteger(market , MkString("hidden") ); _ = hidden;
    active := OpEq2(hidden , MkInteger(0) ); _ = active;
    feeString := this.SafeString(market , MkString("fee") ); _ = feeString;
    feeString = Precise.StringDiv(feeString , MkString("100") );
    takerFee := this.ParseNumber(feeString ); _ = takerFee;
    makerFee := this.ParseNumber(feeString ); _ = makerFee;
    result.Push(MkMap(&VarMap{
            "id":id ,
            "symbol":symbol ,
            "base":base ,
            "quote":quote ,
            "baseId":baseId ,
            "quoteId":quoteId ,
            "active":active ,
            "taker":takerFee ,
            "maker":makerFee ,
            "precision":precision ,
            "limits":limits ,
            "info":market ,
            }));
  }
  return result ;
}

func (this *Yobit) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"pair":*(market ).At(MkString("id") )}); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("publicGetDepthPair"), this.Extend(request , params )); _ = response;
  market_id_in_reponse := OpHasMember(*(market ).At(MkString("id") ), response ); _ = market_id_in_reponse;
  if IsTrue(OpNot(market_id_in_reponse )) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , OpAdd(*(market ).At(MkString("symbol") ), MkString(" order book is empty or not available") )))));
  }
  orderbook := *(response ).At(*(market ).At(MkString("id") )); _ = orderbook;
  return this.ParseOrderBook(orderbook , symbol );
}

func (this *Yobit) FetchOrderBooks(goArgs ...*Variant) *Variant{
  symbols := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbols;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  ids := OpCopy(MkUndefined() ); _ = ids;
  if IsTrue(OpEq2(symbols , MkUndefined() )) {
    ids = (*((this).At(MkString("ids")))).Call(MkString("join"), MkString("-") );
    if IsTrue(OpGt(ids.Length , MkInteger(2048) )) {
      numIds := OpCopy((*((this).At(MkString("ids")))).Length ); _ = numIds;
      panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" has ") , OpAdd(numIds.ToString(), MkString(" symbols exceeding max URL length, you are required to specify a list of symbols in the first argument to fetchOrderBooks") )))));
    }
  } else {
    ids = this.MarketIds(symbols );
    ids = ids.Join(MkString("-") );
  }
  request := MkMap(&VarMap{"pair":ids }); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("publicGetDepthPair"), this.Extend(request , params )); _ = response;
  result := MkMap(&VarMap{}); _ = result;
  ids = GoGetKeys(response );
  for i := MkInteger(0) ; IsTrue(OpLw(i , ids.Length )); OpIncr(& i ){
    id := *(ids ).At(i ); _ = id;
    symbol := this.SafeSymbol(id ); _ = symbol;
    *(result ).At(symbol )= this.ParseOrderBook(*(response ).At(id ));
  }
  return result ;
}

func (this *Yobit) ParseTicker(goArgs ...*Variant) *Variant{
  ticker := GoGetArg(goArgs, 0, MkUndefined()); _ = ticker;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.SafeTimestamp(ticker , MkString("updated") ); _ = timestamp;
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  if IsTrue(OpNotEq2(market , MkUndefined() )) {
    symbol = *(market ).At(MkString("symbol") );
  }
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
        "average":this.SafeNumber(ticker , MkString("avg") ),
        "baseVolume":this.SafeNumber(ticker , MkString("vol_cur") ),
        "quoteVolume":this.SafeNumber(ticker , MkString("vol") ),
        "info":ticker ,
        });
}

func (this *Yobit) FetchTickers(goArgs ...*Variant) *Variant{
  symbols := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbols;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  ids := OpCopy(*this.At(MkString("ids")) ); _ = ids;
  if IsTrue(OpEq2(symbols , MkUndefined() )) {
    numIds := OpCopy(ids.Length ); _ = numIds;
    ids = ids.Join(MkString("-") );
    maxLength := this.SafeInteger(*this.At(MkString("options")) , MkString("fetchTickersMaxLength") , MkInteger(2048) ); _ = maxLength;
    if IsTrue(OpGt(ids.Length , *(*this.At(MkString("options")) ).At(MkString("fetchTickersMaxLength") ))) {
      panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" has ") , OpAdd(numIds.ToString(), OpAdd(MkString(" markets exceeding max URL length for this endpoint (") , OpAdd(maxLength.ToString(), MkString(" characters), please, specify a list of symbols of interest in the first argument to fetchTickers") )))))));
    }
  } else {
    ids = this.MarketIds(symbols );
    ids = ids.Join(MkString("-") );
  }
  request := MkMap(&VarMap{"pair":ids }); _ = request;
  tickers := this.Call(MkString("publicGetTickerPair"), this.Extend(request , params )); _ = tickers;
  result := MkMap(&VarMap{}); _ = result;
  keys := GoGetKeys(tickers ); _ = keys;
  for k := MkInteger(0) ; IsTrue(OpLw(k , keys.Length )); OpIncr(& k ){
    id := *(keys ).At(k ); _ = id;
    ticker := *(tickers ).At(id ); _ = ticker;
    market := this.SafeMarket(id ); _ = market;
    symbol := *(market ).At(MkString("symbol") ); _ = symbol;
    *(result ).At(symbol )= this.ParseTicker(ticker , market );
  }
  return this.FilterByArray(result , MkString("symbol") , symbols );
}

func (this *Yobit) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  tickers := this.FetchTickers(MkArray(&VarArray{
            symbol ,
            }), params ); _ = tickers;
  return *(tickers ).At(symbol );
}

func (this *Yobit) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  timestamp := this.SafeTimestamp(trade , MkString("timestamp") ); _ = timestamp;
  side := this.SafeString(trade , MkString("type") ); _ = side;
  if IsTrue(OpEq2(side , MkString("ask") )) {
    side = MkString("sell") ;
  } else {
    if IsTrue(OpEq2(side , MkString("bid") )) {
      side = MkString("buy") ;
    }
  }
  priceString := this.SafeString2(trade , MkString("rate") , MkString("price") ); _ = priceString;
  id := this.SafeString2(trade , MkString("trade_id") , MkString("tid") ); _ = id;
  order := this.SafeString(trade , MkString("order_id") ); _ = order;
  marketId := this.SafeString(trade , MkString("pair") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market ); _ = symbol;
  amountString := this.SafeString(trade , MkString("amount") ); _ = amountString;
  cost := this.ParseNumber(Precise.StringMul(priceString , amountString )); _ = cost;
  price := this.ParseNumber(priceString ); _ = price;
  amount := this.ParseNumber(amountString ); _ = amount;
  type_ := MkString("limit") ; _ = type_;
  fee := OpCopy(MkUndefined() ); _ = fee;
  feeCost := this.SafeNumber(trade , MkString("commission") ); _ = feeCost;
  if IsTrue(OpNotEq2(feeCost , MkUndefined() )) {
    feeCurrencyId := this.SafeString(trade , MkString("commissionCurrency") ); _ = feeCurrencyId;
    feeCurrencyCode := this.SafeCurrencyCode(feeCurrencyId ); _ = feeCurrencyCode;
    fee = MkMap(&VarMap{
        "cost":feeCost ,
        "currency":feeCurrencyCode ,
        });
  }
  isYourOrder := this.SafeValue(trade , MkString("is_your_order") ); _ = isYourOrder;
  if IsTrue(OpNotEq2(isYourOrder , MkUndefined() )) {
    if IsTrue(OpEq2(fee , MkUndefined() )) {
      fee = this.CalculateFee(symbol , type_ , side , amount , price , MkString("taker") );
    }
  }
  return MkMap(&VarMap{
        "id":id ,
        "order":order ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "symbol":symbol ,
        "type":type_ ,
        "side":side ,
        "takerOrMaker":MkUndefined() ,
        "price":price ,
        "amount":amount ,
        "cost":cost ,
        "fee":fee ,
        "info":trade ,
        });
}

func (this *Yobit) FetchTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"pair":*(market ).At(MkString("id") )}); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("publicGetTradesPair"), this.Extend(request , params )); _ = response;
  if IsTrue(GoIsArray(response )) {
    numElements := OpCopy(response.Length ); _ = numElements;
    if IsTrue(OpEq2(numElements , MkInteger(0) )) {
      return MkArray(&VarArray{});
    }
  }
  return this.ParseTrades(*(response ).At(*(market ).At(MkString("id") )), market , since , limit );
}

func (this *Yobit) CreateOrder(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 2, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 3, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 4, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 5, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(type_ , MkString("market") )) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , MkString(" allows limit orders only") )));
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "pair":*(market ).At(MkString("id") ),
        "type":side ,
        "amount":this.AmountToPrecision(symbol , amount ),
        "rate":this.PriceToPrecision(symbol , price ),
        }); _ = request;
  response := this.Call(MkString("privatePostTrade"), this.Extend(request , params )); _ = response;
  id := OpCopy(MkUndefined() ); _ = id;
  status := MkString("open") ; _ = status;
  filled := MkNumber(0.0) ; _ = filled;
  remaining := OpCopy(amount ); _ = remaining;
  if IsTrue(OpHasMember(MkString("return") , response )) {
    id = this.SafeString(*(response ).At(MkString("return") ), MkString("order_id") );
    if IsTrue(OpEq2(id , MkString("0") )) {
      id = this.SafeString(*(response ).At(MkString("return") ), MkString("init_order_id") );
      status = MkString("closed") ;
    }
    filled = this.SafeNumber(*(response ).At(MkString("return") ), MkString("received") , MkNumber(0.0) );
    remaining = this.SafeNumber(*(response ).At(MkString("return") ), MkString("remains") , amount );
  }
  timestamp := this.Milliseconds(); _ = timestamp;
  return MkMap(&VarMap{
        "id":id ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "lastTradeTimestamp":MkUndefined() ,
        "status":status ,
        "symbol":symbol ,
        "type":type_ ,
        "side":side ,
        "price":price ,
        "cost":OpMulti(price , filled ),
        "amount":amount ,
        "remaining":remaining ,
        "filled":filled ,
        "fee":MkUndefined() ,
        "info":response ,
        "clientOrderId":MkUndefined() ,
        "average":MkUndefined() ,
        "trades":MkUndefined() ,
        });
}

func (this *Yobit) CancelOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"order_id":parseInt(id )}); _ = request;
  return this.Call(MkString("privatePostCancelOrder"), this.Extend(request , params ));
}

func (this *Yobit) ParseOrderStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "0":MkString("open") ,
        "1":MkString("closed") ,
        "2":MkString("canceled") ,
        "3":MkString("open") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Yobit) ParseOrder(goArgs ...*Variant) *Variant{
  order := GoGetArg(goArgs, 0, MkUndefined()); _ = order;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  id := this.SafeString(order , MkString("id") ); _ = id;
  status := this.ParseOrderStatus(this.SafeString(order , MkString("status") )); _ = status;
  timestamp := this.SafeTimestamp(order , MkString("timestamp_created") ); _ = timestamp;
  marketId := this.SafeString(order , MkString("pair") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market ); _ = symbol;
  remaining := this.SafeNumber(order , MkString("amount") ); _ = remaining;
  amount := this.SafeNumber(order , MkString("start_amount") ); _ = amount;
  price := this.SafeNumber(order , MkString("rate") ); _ = price;
  fee := OpCopy(MkUndefined() ); _ = fee;
  type_ := MkString("limit") ; _ = type_;
  side := this.SafeString(order , MkString("type") ); _ = side;
  return this.SafeOrder(MkMap(&VarMap{
            "info":order ,
            "id":id ,
            "clientOrderId":MkUndefined() ,
            "symbol":symbol ,
            "timestamp":timestamp ,
            "datetime":this.Iso8601(timestamp ),
            "lastTradeTimestamp":MkUndefined() ,
            "type":type_ ,
            "timeInForce":MkUndefined() ,
            "postOnly":MkUndefined() ,
            "side":side ,
            "price":price ,
            "stopPrice":MkUndefined() ,
            "cost":MkUndefined() ,
            "amount":amount ,
            "remaining":remaining ,
            "filled":MkUndefined() ,
            "status":status ,
            "fee":fee ,
            "average":MkUndefined() ,
            "trades":MkUndefined() ,
            }));
}

func (this *Yobit) FetchOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"order_id":parseInt(id )}); _ = request;
  response := this.Call(MkString("privatePostOrderInfo"), this.Extend(request , params )); _ = response;
  id = id.ToString();
  orders := this.SafeValue(response , MkString("return") , MkMap(&VarMap{})); _ = orders;
  return this.ParseOrder(this.Extend(MkMap(&VarMap{"id":id }), *(orders ).At(id )));
}

func (this *Yobit) FetchOpenOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchOpenOrders() requires a symbol argument") )));
  }
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market := this.Market(symbol ); _ = market;
    *(request ).At(MkString("pair") )= *(market ).At(MkString("id") );
  }
  response := this.Call(MkString("privatePostActiveOrders"), this.Extend(request , params )); _ = response;
  orders := this.SafeValue(response , MkString("return") , MkArray(&VarArray{})); _ = orders;
  return this.ParseOrders(orders , market , since , limit );
}

func (this *Yobit) FetchMyTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpEq2(symbol , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" fetchMyTrades() requires a `symbol` argument") )));
  }
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"pair":*(market ).At(MkString("id") )}); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("count") )= parseInt(limit );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("since") )= parseInt(OpDiv(since , MkInteger(1000) ));
  }
  response := this.Call(MkString("privatePostTradeHistory"), this.Extend(request , params )); _ = response;
  trades := this.SafeValue(response , MkString("return") , MkMap(&VarMap{})); _ = trades;
  ids := GoGetKeys(trades ); _ = ids;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , ids.Length )); OpIncr(& i ){
    id := *(ids ).At(i ); _ = id;
    trade := this.ParseTrade(this.Extend(*(trades ).At(id ), MkMap(&VarMap{"trade_id":id })), market ); _ = trade;
    result.Push(trade );
  }
  return this.FilterBySymbolSinceLimit(result , symbol , since , limit );
}

func (this *Yobit) CreateDepositAddress(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  request := MkMap(&VarMap{"need_new":MkInteger(1) }); _ = request;
  response := this.FetchDepositAddress(code , this.Extend(request , params )); _ = response;
  address := this.SafeString(response , MkString("address") ); _ = address;
  this.CheckAddress(address );
  return MkMap(&VarMap{
        "currency":code ,
        "address":address ,
        "tag":MkUndefined() ,
        "info":*(response ).At(MkString("info") ),
        });
}

func (this *Yobit) FetchDepositAddress(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  currency := this.Currency(code ); _ = currency;
  request := MkMap(&VarMap{
        "coinName":*(currency ).At(MkString("id") ),
        "need_new":MkInteger(0) ,
        }); _ = request;
  response := this.Call(MkString("privatePostGetDepositAddress"), this.Extend(request , params )); _ = response;
  address := this.SafeString(*(response ).At(MkString("return") ), MkString("address") ); _ = address;
  this.CheckAddress(address );
  return MkMap(&VarMap{
        "currency":code ,
        "address":address ,
        "tag":MkUndefined() ,
        "info":response ,
        });
}

func (this *Yobit) Withdraw(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  amount := GoGetArg(goArgs, 1, MkUndefined()); _ = amount;
  address := GoGetArg(goArgs, 2, MkUndefined()); _ = address;
  tag := GoGetArg(goArgs, 3, MkUndefined() ); _ = tag;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.CheckAddress(address );
  this.LoadMarkets();
  currency := this.Currency(code ); _ = currency;
  request := MkMap(&VarMap{
        "coinName":*(currency ).At(MkString("id") ),
        "amount":amount ,
        "address":address ,
        }); _ = request;
  if IsTrue(OpNotEq2(tag , MkUndefined() )) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , MkString(" withdraw() does not support the tag argument yet due to a lack of docs on withdrawing with tag/memo on behalf of the exchange.") )));
  }
  response := this.Call(MkString("privatePostWithdrawCoinsToAddress"), this.Extend(request , params )); _ = response;
  return MkMap(&VarMap{
        "info":response ,
        "id":MkUndefined() ,
        });
}

func (this *Yobit) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("public") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  url := *(*(*this.At(MkString("urls")) ).At(MkString("api") )).At(api ); _ = url;
  query := this.Omit(params , this.ExtractParams(path )); _ = query;
  if IsTrue(OpEq2(api , MkString("private") )) {
    this.CheckRequiredCredentials();
    nonce := this.Nonce(); _ = nonce;
    body = this.Urlencode(this.Extend(MkMap(&VarMap{
                "nonce":nonce ,
                "method":path ,
                }), query ));
    signature := this.Hmac(this.Encode(body ), this.Encode(*this.At(MkString("secret")) ), MkString("sha512") ); _ = signature;
    headers = MkMap(&VarMap{
        "Content-Type":MkString("application/x-www-form-urlencoded") ,
        "Key":*this.At(MkString("apiKey")) ,
        "Sign":signature ,
        });
  } else {
    if IsTrue(OpEq2(api , MkString("public") )) {
      OpAddAssign(& url , OpAdd(MkString("/") , OpAdd(*this.At(MkString("version")) , OpAdd(MkString("/") , this.ImplodeParams(path , params )))));
      if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
        OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(query )));
      }
    } else {
      OpAddAssign(& url , OpAdd(MkString("/") , this.ImplodeParams(path , params )));
      if IsTrue(OpEq2(method , MkString("GET") )) {
        if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
          OpAddAssign(& url , OpAdd(MkString("?") , this.Urlencode(query )));
        }
      } else {
        if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
          body = this.Json(query );
          headers = MkMap(&VarMap{"Content-Type":MkString("application/json") });
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

func (this *Yobit) HandleErrors(goArgs ...*Variant) *Variant{
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
  if IsTrue(OpHasMember(MkString("success") , response )) {
    success := this.SafeValue(response , MkString("success") , MkBool(false) ); _ = success;
    if IsTrue(OpEq2(OpType(success ), MkString("string") )) {
      if IsTrue(OpOr(OpEq2(success , MkString("true") ), OpEq2(success , MkString("1") ))) {
        success = OpCopy(MkBool(true) );
      } else {
        success = OpCopy(MkBool(false) );
      }
    }
    if IsTrue(OpNot(success )) {
      code := this.SafeString(response , MkString("code") ); _ = code;
      message := this.SafeString(response , MkString("error") ); _ = message;
      feedback := OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , body )); _ = feedback;
      this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("exact") ), code , feedback );
      this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("exact") ), message , feedback );
      this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("broad") ), message , feedback );
      panic(NewExchangeError(feedback ));
    }
  }
  return MkUndefined()
}

