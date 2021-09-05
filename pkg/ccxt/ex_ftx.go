package ccxt

import (
)

type Ftx struct {
	*ExchangeBase
}

var _ Exchange = (*Ftx)(nil)

func init() {
	exchange := &Ftx{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Ftx) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("ftx") ,
            "name":MkString("FTX") ,
            "countries":MkArray(&VarArray{
                MkString("HK") ,
                }),
            "rateLimit":MkInteger(50) ,
            "certified":MkBool(true) ,
            "pro":MkBool(true) ,
            "hostname":MkString("ftx.com") ,
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/1294454/67149189-df896480-f2b0-11e9-8816-41593e17f9ec.jpg") ,
                "www":MkString("https://ftx.com") ,
                "api":MkMap(&VarMap{
                    "public":MkString("https://{hostname}") ,
                    "private":MkString("https://{hostname}") ,
                    }),
                "doc":MkString("https://github.com/ftexchange/ftx") ,
                "fees":MkString("https://ftexchange.zendesk.com/hc/en-us/articles/360024479432-Fees") ,
                "referral":MkMap(&VarMap{
                    "url":MkString("https://ftx.com/#a=ccxt") ,
                    "discount":MkNumber(0.05) ,
                    }),
                }),
            "has":MkMap(&VarMap{
                "cancelAllOrders":MkBool(true) ,
                "cancelOrder":MkBool(true) ,
                "createOrder":MkBool(true) ,
                "editOrder":MkBool(true) ,
                "fetchBalance":MkBool(true) ,
                "fetchClosedOrders":MkBool(false) ,
                "fetchCurrencies":MkBool(true) ,
                "fetchDepositAddress":MkBool(true) ,
                "fetchDeposits":MkBool(true) ,
                "fetchFundingFees":MkBool(false) ,
                "fetchMarkets":MkBool(true) ,
                "fetchMyTrades":MkBool(true) ,
                "fetchOHLCV":MkBool(true) ,
                "fetchOpenOrders":MkBool(true) ,
                "fetchOrder":MkBool(true) ,
                "fetchOrderBook":MkBool(true) ,
                "fetchOrders":MkBool(true) ,
                "fetchPositions":MkBool(true) ,
                "fetchTicker":MkBool(true) ,
                "fetchTickers":MkBool(true) ,
                "fetchTrades":MkBool(true) ,
                "fetchTradingFees":MkBool(true) ,
                "fetchWithdrawals":MkBool(true) ,
                "setLeverage":MkBool(true) ,
                "withdraw":MkBool(true) ,
                }),
            "timeframes":MkMap(&VarMap{
                "15s":MkString("15") ,
                "1m":MkString("60") ,
                "5m":MkString("300") ,
                "15m":MkString("900") ,
                "1h":MkString("3600") ,
                "4h":MkString("14400") ,
                "1d":MkString("86400") ,
                }),
            "api":MkMap(&VarMap{
                "public":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("coins") ,
                        MkString("markets") ,
                        MkString("markets/{market_name}") ,
                        MkString("markets/{market_name}/orderbook") ,
                        MkString("markets/{market_name}/trades") ,
                        MkString("markets/{market_name}/candles") ,
                        MkString("futures") ,
                        MkString("futures/{future_name}") ,
                        MkString("futures/{future_name}/stats") ,
                        MkString("funding_rates") ,
                        MkString("indexes/{index_name}/weights") ,
                        MkString("expired_futures") ,
                        MkString("indexes/{market_name}/candles") ,
                        MkString("wallet/coins") ,
                        MkString("lt/tokens") ,
                        MkString("lt/{token_name}") ,
                        MkString("etfs/rebalance_info") ,
                        MkString("options/requests") ,
                        MkString("options/trades") ,
                        MkString("options/historical_volumes/BTC") ,
                        MkString("stats/24h_options_volume") ,
                        MkString("options/open_interest/BTC") ,
                        MkString("options/historical_open_interest/BTC") ,
                        MkString("spot_margin/history") ,
                        MkString("spot_margin/borrow_summary") ,
                        MkString("nft/nfts") ,
                        MkString("nft/{nft_id}") ,
                        MkString("nft/{nft_id}/trades") ,
                        MkString("nft/all_trades") ,
                        MkString("nft/{nft_id}/account_info") ,
                        MkString("nft/collections") ,
                        MkString("ftxpay/apps/{user_specific_id}/details") ,
                        MkString("stats/latency_stats") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("ftxpay/apps/{user_specific_id}/orders") ,
                        }),
                    }),
                "private":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("subaccounts") ,
                        MkString("subaccounts/{nickname}/balances") ,
                        MkString("account") ,
                        MkString("positions") ,
                        MkString("wallet/balances") ,
                        MkString("wallet/all_balances") ,
                        MkString("wallet/deposit_address/{coin}") ,
                        MkString("wallet/deposits") ,
                        MkString("wallet/withdrawals") ,
                        MkString("wallet/airdrops") ,
                        MkString("wallet/withdrawal_fee") ,
                        MkString("wallet/saved_addresses") ,
                        MkString("orders") ,
                        MkString("orders/history") ,
                        MkString("orders/{order_id}") ,
                        MkString("orders/by_client_id/{client_order_id}") ,
                        MkString("conditional_orders") ,
                        MkString("conditional_orders/{conditional_order_id}/triggers") ,
                        MkString("conditional_orders/history") ,
                        MkString("fills") ,
                        MkString("funding_payments") ,
                        MkString("lt/balances") ,
                        MkString("lt/creations") ,
                        MkString("lt/redemptions") ,
                        MkString("options/my_requests") ,
                        MkString("options/requests/{request_id}/quotes") ,
                        MkString("options/my_quotes") ,
                        MkString("options/account_info") ,
                        MkString("options/positions") ,
                        MkString("options/fills") ,
                        MkString("staking/stakes") ,
                        MkString("staking/unstake_requests") ,
                        MkString("staking/balances") ,
                        MkString("staking/staking_rewards") ,
                        MkString("otc/quotes/{quoteId}") ,
                        MkString("spot_margin/borrow_rates") ,
                        MkString("spot_margin/lending_rates") ,
                        MkString("spot_margin/market_info") ,
                        MkString("spot_margin/borrow_history") ,
                        MkString("spot_margin/lending_history") ,
                        MkString("spot_margin/offers") ,
                        MkString("spot_margin/lending_info") ,
                        MkString("nft/balances") ,
                        MkString("nft/bids") ,
                        MkString("nft/deposits") ,
                        MkString("nft/withdrawals") ,
                        MkString("nft/fills") ,
                        MkString("nft/gallery/{gallery_id}") ,
                        MkString("nft/gallery_settings") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("subaccounts") ,
                        MkString("subaccounts/update_name") ,
                        MkString("subaccounts/transfer") ,
                        MkString("account/leverage") ,
                        MkString("wallet/withdrawals") ,
                        MkString("wallet/saved_addresses") ,
                        MkString("orders") ,
                        MkString("conditional_orders") ,
                        MkString("orders/{order_id}/modify") ,
                        MkString("orders/by_client_id/{client_order_id}/modify") ,
                        MkString("conditional_orders/{order_id}/modify") ,
                        MkString("lt/{token_name}/create") ,
                        MkString("lt/{token_name}/redeem") ,
                        MkString("options/requests") ,
                        MkString("options/requests/{request_id}/quotes") ,
                        MkString("options/quotes/{quote_id}/accept") ,
                        MkString("staking/unstake_requests") ,
                        MkString("srm_stakes/stakes") ,
                        MkString("otc/quotes/{quote_id}/accept") ,
                        MkString("otc/quotes") ,
                        MkString("spot_margin/offers") ,
                        MkString("nft/offer") ,
                        MkString("nft/buy") ,
                        MkString("nft/auction") ,
                        MkString("nft/edit_auction") ,
                        MkString("nft/cancel_auction") ,
                        MkString("nft/bids") ,
                        MkString("nft/redeem") ,
                        MkString("nft/gallery_settings") ,
                        MkString("ftxpay/apps/{user_specific_id}/orders") ,
                        }),
                    "delete":MkArray(&VarArray{
                        MkString("subaccounts") ,
                        MkString("wallet/saved_addresses/{saved_address_id}") ,
                        MkString("orders/{order_id}") ,
                        MkString("orders/by_client_id/{client_order_id}") ,
                        MkString("orders") ,
                        MkString("conditional_orders/{order_id}") ,
                        MkString("options/requests/{request_id}") ,
                        MkString("options/quotes/{quote_id}") ,
                        MkString("staking/unstake_requests/{request_id}") ,
                        }),
                    }),
                }),
            "fees":MkMap(&VarMap{
                "trading":MkMap(&VarMap{
                    "tierBased":MkBool(true) ,
                    "percentage":MkBool(true) ,
                    "maker":this.ParseNumber(MkString("0.0002") ),
                    "taker":this.ParseNumber(MkString("0.0007") ),
                    "tiers":MkMap(&VarMap{
                        "taker":MkArray(&VarArray{
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("0") ),
                                this.ParseNumber(MkString("0.0007") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("2000000") ),
                                this.ParseNumber(MkString("0.0006") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("5000000") ),
                                this.ParseNumber(MkString("0.00055") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("10000000") ),
                                this.ParseNumber(MkString("0.0005") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("25000000") ),
                                this.ParseNumber(MkString("0.045") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("50000000") ),
                                this.ParseNumber(MkString("0.0004") ),
                                }),
                            }),
                        "maker":MkArray(&VarArray{
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("0") ),
                                this.ParseNumber(MkString("0.0002") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("2000000") ),
                                this.ParseNumber(MkString("0.00015") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("5000000") ),
                                this.ParseNumber(MkString("0.0001") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("10000000") ),
                                this.ParseNumber(MkString("0.00005") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("25000000") ),
                                this.ParseNumber(MkString("0") ),
                                }),
                            MkArray(&VarArray{
                                this.ParseNumber(MkString("50000000") ),
                                this.ParseNumber(MkString("0") ),
                                }),
                            }),
                        }),
                    }),
                "funding":MkMap(&VarMap{"withdraw":MkMap(&VarMap{})}),
                }),
            "exceptions":MkMap(&VarMap{
                "exact":MkMap(&VarMap{
                    "Please slow down":RateLimitExceeded ,
                    "Size too small for provide":InvalidOrder ,
                    "Not logged in":AuthenticationError ,
                    "Not enough balances":InsufficientFunds ,
                    "InvalidPrice":InvalidOrder ,
                    "Size too small":InvalidOrder ,
                    "Size too large":InvalidOrder ,
                    "Missing parameter price":InvalidOrder ,
                    "Order not found":OrderNotFound ,
                    "Order already closed":InvalidOrder ,
                    "Trigger price too high":InvalidOrder ,
                    "Trigger price too low":InvalidOrder ,
                    "Order already queued for cancellation":CancelPending ,
                    "Duplicate client order ID":DuplicateOrderId ,
                    "Spot orders cannot be reduce-only":InvalidOrder ,
                    "Invalid reduce-only order":InvalidOrder ,
                    "Account does not have enough balances":InsufficientFunds ,
                    }),
                "broad":MkMap(&VarMap{
                    "Account does not have enough margin for order":InsufficientFunds ,
                    "Invalid parameter":BadRequest ,
                    "The requested URL was not found on the server":BadRequest ,
                    "No such coin":BadRequest ,
                    "No such subaccount":BadRequest ,
                    "No such future":BadSymbol ,
                    "No such market":BadSymbol ,
                    "Do not send more than":RateLimitExceeded ,
                    "An unexpected error occurred":ExchangeNotAvailable ,
                    "Please retry request":ExchangeNotAvailable ,
                    "Please try again":ExchangeNotAvailable ,
                    "Try again":ExchangeNotAvailable ,
                    "Only have permissions for subaccount":PermissionDenied ,
                    }),
                }),
            "precisionMode":TICK_SIZE ,
            "options":MkMap(&VarMap{
                "cancelOrder":MkMap(&VarMap{"method":MkString("privateDeleteOrdersOrderId") }),
                "fetchOpenOrders":MkMap(&VarMap{"method":MkString("privateGetOrders") }),
                "fetchOrders":MkMap(&VarMap{"method":MkString("privateGetOrdersHistory") }),
                "sign":MkMap(&VarMap{
                    "ftx.com":MkString("FTX") ,
                    "ftx.us":MkString("FTXUS") ,
                    }),
                }),
            }));
}

func (this *Ftx) FetchCurrencies(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetCoins"), params ); _ = response;
  currencies := this.SafeValue(response , MkString("result") , MkArray(&VarArray{})); _ = currencies;
  result := MkMap(&VarMap{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , currencies.Length )); OpIncr(& i ){
    currency := *(currencies ).At(i ); _ = currency;
    id := this.SafeString(currency , MkString("id") ); _ = id;
    code := this.SafeCurrencyCode(id ); _ = code;
    name := this.SafeString(currency , MkString("name") ); _ = name;
    *(result ).At(code )= MkMap(&VarMap{
        "id":id ,
        "code":code ,
        "info":currency ,
        "type":MkUndefined() ,
        "name":name ,
        "active":MkUndefined() ,
        "fee":MkUndefined() ,
        "precision":MkUndefined() ,
        "limits":MkMap(&VarMap{
            "withdraw":MkMap(&VarMap{
                "min":MkUndefined() ,
                "max":MkUndefined() ,
                }),
            "amount":MkMap(&VarMap{
                "min":MkUndefined() ,
                "max":MkUndefined() ,
                }),
            }),
        });
  }
  return result ;
}

func (this *Ftx) FetchMarkets(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetMarkets"), params ); _ = response;
  result := MkArray(&VarArray{}); _ = result;
  markets := this.SafeValue(response , MkString("result") , MkArray(&VarArray{})); _ = markets;
  for i := MkInteger(0) ; IsTrue(OpLw(i , markets.Length )); OpIncr(& i ){
    market := *(markets ).At(i ); _ = market;
    id := this.SafeString(market , MkString("name") ); _ = id;
    baseId := this.SafeString2(market , MkString("baseCurrency") , MkString("underlying") ); _ = baseId;
    quoteId := this.SafeString(market , MkString("quoteCurrency") , MkString("USD") ); _ = quoteId;
    type_ := this.SafeString(market , MkString("type") ); _ = type_;
    base := this.SafeCurrencyCode(baseId ); _ = base;
    quote := this.SafeCurrencyCode(quoteId ); _ = quote;
    symbol := OpTrinary(OpEq2(type_ , MkString("future") ), this.SafeString(market , MkString("name") ), OpAdd(base , OpAdd(MkString("/") , quote ))); _ = symbol;
    active := this.SafeValue(market , MkString("enabled") ); _ = active;
    sizeIncrement := this.SafeNumber(market , MkString("sizeIncrement") ); _ = sizeIncrement;
    priceIncrement := this.SafeNumber(market , MkString("priceIncrement") ); _ = priceIncrement;
    precision := MkMap(&VarMap{
          "amount":sizeIncrement ,
          "price":priceIncrement ,
          }); _ = precision;
    result.Push(MkMap(&VarMap{
            "id":id ,
            "symbol":symbol ,
            "base":base ,
            "quote":quote ,
            "baseId":baseId ,
            "quoteId":quoteId ,
            "type":type_ ,
            "future":OpEq2(type_ , MkString("future") ),
            "spot":OpEq2(type_ , MkString("spot") ),
            "active":active ,
            "precision":precision ,
            "limits":MkMap(&VarMap{
                "amount":MkMap(&VarMap{
                    "min":sizeIncrement ,
                    "max":MkUndefined() ,
                    }),
                "price":MkMap(&VarMap{
                    "min":priceIncrement ,
                    "max":MkUndefined() ,
                    }),
                "cost":MkMap(&VarMap{
                    "min":MkUndefined() ,
                    "max":MkUndefined() ,
                    }),
                }),
            "info":market ,
            }));
  }
  return result ;
}

func (this *Ftx) ParseTicker(goArgs ...*Variant) *Variant{
  ticker := GoGetArg(goArgs, 0, MkUndefined()); _ = ticker;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  marketId := this.SafeString(ticker , MkString("name") ); _ = marketId;
  if IsTrue(OpHasMember(marketId , *this.At(MkString("markets_by_id")) )) {
    market = *(*this.At(MkString("markets_by_id")) ).At(marketId );
  } else {
    type_ := this.SafeString(ticker , MkString("type") ); _ = type_;
    if IsTrue(OpEq2(type_ , MkString("future") )) {
      symbol = OpCopy(marketId );
    } else {
      base := this.SafeCurrencyCode(this.SafeString(ticker , MkString("baseCurrency") )); _ = base;
      quote := this.SafeCurrencyCode(this.SafeString(ticker , MkString("quoteCurrency") )); _ = quote;
      if IsTrue(OpAnd(OpNotEq2(base , MkUndefined() ), OpNotEq2(quote , MkUndefined() ))) {
        symbol = OpAdd(base , OpAdd(MkString("/") , quote ));
      }
    }
  }
  if IsTrue(OpAnd(OpEq2(symbol , MkUndefined() ), OpNotEq2(market , MkUndefined() ))) {
    symbol = *(market ).At(MkString("symbol") );
  }
  last := this.SafeNumber(ticker , MkString("last") ); _ = last;
  timestamp := this.SafeTimestamp(ticker , MkString("time") , this.Milliseconds()); _ = timestamp;
  return MkMap(&VarMap{
        "symbol":symbol ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "high":this.SafeNumber(ticker , MkString("high") ),
        "low":this.SafeNumber(ticker , MkString("low") ),
        "bid":this.SafeNumber(ticker , MkString("bid") ),
        "bidVolume":this.SafeNumber(ticker , MkString("bidSize") ),
        "ask":this.SafeNumber(ticker , MkString("ask") ),
        "askVolume":this.SafeNumber(ticker , MkString("askSize") ),
        "vwap":MkUndefined() ,
        "open":MkUndefined() ,
        "close":last ,
        "last":last ,
        "previousClose":MkUndefined() ,
        "change":MkUndefined() ,
        "percentage":this.SafeNumber(ticker , MkString("change24h") ),
        "average":MkUndefined() ,
        "baseVolume":MkUndefined() ,
        "quoteVolume":this.SafeNumber(ticker , MkString("quoteVolume24h") ),
        "info":ticker ,
        });
}

func (this *Ftx) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"market_name":*(market ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("publicGetMarketsMarketName"), this.Extend(request , params )); _ = response;
  result := this.SafeValue(response , MkString("result") , MkMap(&VarMap{})); _ = result;
  return this.ParseTicker(result , market );
}

func (this *Ftx) FetchTickers(goArgs ...*Variant) *Variant{
  symbols := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbols;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("publicGetMarkets"), params ); _ = response;
  tickers := this.SafeValue(response , MkString("result") , MkArray(&VarArray{})); _ = tickers;
  return this.ParseTickers(tickers , symbols );
}

func (this *Ftx) FetchOrderBook(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{"market_name":*(market ).At(MkString("id") )}); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("depth") )= OpCopy(limit );
  }
  response := this.Call(MkString("publicGetMarketsMarketNameOrderbook"), this.Extend(request , params )); _ = response;
  result := this.SafeValue(response , MkString("result") , MkMap(&VarMap{})); _ = result;
  return this.ParseOrderBook(result , symbol );
}

func (this *Ftx) ParseOHLCV(goArgs ...*Variant) *Variant{
  ohlcv := GoGetArg(goArgs, 0, MkUndefined()); _ = ohlcv;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  return MkArray(&VarArray{
        this.SafeInteger(ohlcv , MkString("time") ),
        this.SafeNumber(ohlcv , MkString("open") ),
        this.SafeNumber(ohlcv , MkString("high") ),
        this.SafeNumber(ohlcv , MkString("low") ),
        this.SafeNumber(ohlcv , MkString("close") ),
        this.SafeNumber(ohlcv , MkString("volume") ),
        });
}

func (this *Ftx) GetMarketId(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  key := GoGetArg(goArgs, 1, MkUndefined()); _ = key;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  parts := this.GetMarketParams(symbol , key , params ); _ = parts;
  return this.SafeString(parts , MkInteger(1) , symbol );
}

func (this *Ftx) GetMarketParams(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  key := GoGetArg(goArgs, 1, MkUndefined()); _ = key;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  market := OpCopy(MkUndefined() ); _ = market;
  marketId := OpCopy(MkUndefined() ); _ = marketId;
  if IsTrue(OpHasMember(symbol , *this.At(MkString("markets")) )) {
    market = this.Market(symbol );
    marketId = *(market ).At(MkString("id") );
  } else {
    marketId = this.SafeString(params , key , symbol );
  }
  return MkArray(&VarArray{
        market ,
        marketId ,
        });
}

func (this *Ftx) FetchOHLCV(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  timeframe := GoGetArg(goArgs, 1, MkString("1m") ); _ = timeframe;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := MkUndefined(); _ = market;
  marketId := MkUndefined(); _ = marketId;
  ArrayUnpack(this.GetMarketParams(symbol , MkString("market_name") , params ), &market, &marketId);
  request := MkMap(&VarMap{
        "resolution":*(*this.At(MkString("timeframes")) ).At(timeframe ),
        "market_name":marketId ,
        }); _ = request;
  limit = OpTrinary(OpEq2(limit , MkUndefined() ), MkInteger(1501) , limit );
  if IsTrue(OpEq2(since , MkUndefined() )) {
    *(request ).At(MkString("end_time") )= this.Seconds();
    *(request ).At(MkString("limit") )= OpCopy(limit );
    *(request ).At(MkString("start_time") )= OpSub(*(request ).At(MkString("end_time") ), OpMulti(limit , this.ParseTimeframe(timeframe )));
  } else {
    *(request ).At(MkString("start_time") )= parseInt(OpDiv(since , MkInteger(1000) ));
    *(request ).At(MkString("limit") )= OpCopy(limit );
    *(request ).At(MkString("end_time") )= this.Sum(*(request ).At(MkString("start_time") ), OpMulti(limit , this.ParseTimeframe(timeframe )));
  }
  response := this.Call(MkString("publicGetMarketsMarketNameCandles"), this.Extend(request , params )); _ = response;
  result := this.SafeValue(response , MkString("result") , MkArray(&VarArray{})); _ = result;
  return this.ParseOHLCVs(result , market , timeframe , since , limit );
}

func (this *Ftx) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  id := this.SafeString(trade , MkString("id") ); _ = id;
  takerOrMaker := this.SafeString(trade , MkString("liquidity") ); _ = takerOrMaker;
  marketId := this.SafeString(trade , MkString("market") ); _ = marketId;
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  if IsTrue(OpHasMember(marketId , *this.At(MkString("markets_by_id")) )) {
    market = *(*this.At(MkString("markets_by_id")) ).At(marketId );
    symbol = *(market ).At(MkString("symbol") );
  } else {
    base := this.SafeCurrencyCode(this.SafeString(trade , MkString("baseCurrency") )); _ = base;
    quote := this.SafeCurrencyCode(this.SafeString(trade , MkString("quoteCurrency") )); _ = quote;
    if IsTrue(OpAnd(OpNotEq2(base , MkUndefined() ), OpNotEq2(quote , MkUndefined() ))) {
      symbol = OpAdd(base , OpAdd(MkString("/") , quote ));
    } else {
      symbol = OpCopy(marketId );
    }
  }
  timestamp := this.Parse8601(this.SafeString(trade , MkString("time") )); _ = timestamp;
  priceString := this.SafeString(trade , MkString("price") ); _ = priceString;
  amountString := this.SafeString(trade , MkString("size") ); _ = amountString;
  price := this.ParseNumber(priceString ); _ = price;
  amount := this.ParseNumber(amountString ); _ = amount;
  cost := this.ParseNumber(Precise.StringMul(priceString , amountString )); _ = cost;
  if IsTrue(OpAnd(OpEq2(symbol , MkUndefined() ), OpNotEq2(market , MkUndefined() ))) {
    symbol = *(market ).At(MkString("symbol") );
  }
  side := this.SafeString(trade , MkString("side") ); _ = side;
  fee := OpCopy(MkUndefined() ); _ = fee;
  feeCost := this.SafeNumber(trade , MkString("fee") ); _ = feeCost;
  if IsTrue(OpNotEq2(feeCost , MkUndefined() )) {
    feeCurrencyId := this.SafeString(trade , MkString("feeCurrency") ); _ = feeCurrencyId;
    feeCurrencyCode := this.SafeCurrencyCode(feeCurrencyId ); _ = feeCurrencyCode;
    fee = MkMap(&VarMap{
        "cost":feeCost ,
        "currency":feeCurrencyCode ,
        "rate":this.SafeNumber(trade , MkString("feeRate") ),
        });
  }
  orderId := this.SafeString(trade , MkString("orderId") ); _ = orderId;
  return MkMap(&VarMap{
        "info":trade ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "symbol":symbol ,
        "id":id ,
        "order":orderId ,
        "type":MkUndefined() ,
        "takerOrMaker":takerOrMaker ,
        "side":side ,
        "price":price ,
        "amount":amount ,
        "cost":cost ,
        "fee":fee ,
        });
}

func (this *Ftx) FetchTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := MkUndefined(); _ = market;
  marketId := MkUndefined(); _ = marketId;
  ArrayUnpack(this.GetMarketParams(symbol , MkString("market_name") , params ), &market, &marketId);
  request := MkMap(&VarMap{"market_name":marketId }); _ = request;
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("start_time") )= parseInt(OpDiv(since , MkInteger(1000) ));
    *(request ).At(MkString("end_time") )= this.Seconds();
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  response := this.Call(MkString("publicGetMarketsMarketNameTrades"), this.Extend(request , params )); _ = response;
  result := this.SafeValue(response , MkString("result") , MkArray(&VarArray{})); _ = result;
  return this.ParseTrades(result , market , since , limit );
}

func (this *Ftx) FetchTradingFees(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("privateGetAccount"), params ); _ = response;
  result := this.SafeValue(response , MkString("result") , MkMap(&VarMap{})); _ = result;
  return MkMap(&VarMap{
        "info":response ,
        "maker":this.SafeNumber(result , MkString("makerFee") ),
        "taker":this.SafeNumber(result , MkString("takerFee") ),
        });
}

func (this *Ftx) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("privateGetWalletBalances"), params ); _ = response;
  result := MkMap(&VarMap{"info":response }); _ = result;
  balances := this.SafeValue(response , MkString("result") , MkArray(&VarArray{})); _ = balances;
  for i := MkInteger(0) ; IsTrue(OpLw(i , balances.Length )); OpIncr(& i ){
    balance := *(balances ).At(i ); _ = balance;
    code := this.SafeCurrencyCode(this.SafeString(balance , MkString("coin") )); _ = code;
    account := this.Account(); _ = account;
    *(account ).At(MkString("free") )= this.SafeString2(balance , MkString("availableWithoutBorrow") , MkString("free") );
    *(account ).At(MkString("total") )= this.SafeString(balance , MkString("total") );
    *(result ).At(code )= OpCopy(account );
  }
  return this.ParseBalance(result );
}

func (this *Ftx) ParseOrderStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "new":MkString("open") ,
        "open":MkString("open") ,
        "closed":MkString("closed") ,
        "triggered":MkString("closed") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Ftx) ParseOrder(goArgs ...*Variant) *Variant{
  order := GoGetArg(goArgs, 0, MkUndefined()); _ = order;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  id := this.SafeString(order , MkString("id") ); _ = id;
  timestamp := this.Parse8601(this.SafeString(order , MkString("createdAt") )); _ = timestamp;
  status := this.ParseOrderStatus(this.SafeString(order , MkString("status") )); _ = status;
  amount := this.SafeNumber(order , MkString("size") ); _ = amount;
  filled := this.SafeNumber(order , MkString("filledSize") ); _ = filled;
  remaining := this.SafeNumber(order , MkString("remainingSize") ); _ = remaining;
  if IsTrue(OpAnd(OpEq2(remaining , MkNumber(0.0) ), OpAnd(OpNotEq2(amount , MkUndefined() ), OpNotEq2(filled , MkUndefined() )))) {
    remaining = Math.Max(OpSub(amount , filled ), MkInteger(0) );
    if IsTrue(OpGt(remaining , MkInteger(0) )) {
      status = MkString("canceled") ;
    }
  }
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  marketId := this.SafeString(order , MkString("market") ); _ = marketId;
  if IsTrue(OpNotEq2(marketId , MkUndefined() )) {
    if IsTrue(OpHasMember(marketId , *this.At(MkString("markets_by_id")) )) {
      market = *(*this.At(MkString("markets_by_id")) ).At(marketId );
      symbol = *(market ).At(MkString("symbol") );
    } else {
      symbol = OpCopy(marketId );
    }
  }
  if IsTrue(OpAnd(OpEq2(symbol , MkUndefined() ), OpNotEq2(market , MkUndefined() ))) {
    symbol = *(market ).At(MkString("symbol") );
  }
  side := this.SafeString(order , MkString("side") ); _ = side;
  type_ := this.SafeString(order , MkString("type") ); _ = type_;
  average := this.SafeNumber(order , MkString("avgFillPrice") ); _ = average;
  price := this.SafeNumber2(order , MkString("price") , MkString("triggerPrice") , average ); _ = price;
  cost := OpCopy(MkUndefined() ); _ = cost;
  if IsTrue(OpAnd(OpNotEq2(filled , MkUndefined() ), OpNotEq2(price , MkUndefined() ))) {
    cost = OpMulti(filled , price );
  }
  lastTradeTimestamp := this.Parse8601(this.SafeString(order , MkString("triggeredAt") )); _ = lastTradeTimestamp;
  clientOrderId := this.SafeString(order , MkString("clientId") ); _ = clientOrderId;
  stopPrice := this.SafeNumber(order , MkString("triggerPrice") ); _ = stopPrice;
  postOnly := this.SafeValue(order , MkString("postOnly") ); _ = postOnly;
  return MkMap(&VarMap{
        "info":order ,
        "id":id ,
        "clientOrderId":clientOrderId ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "lastTradeTimestamp":lastTradeTimestamp ,
        "symbol":symbol ,
        "type":type_ ,
        "timeInForce":MkUndefined() ,
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
        "fee":MkUndefined() ,
        "trades":MkUndefined() ,
        });
}

func (this *Ftx) CreateOrder(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 1, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 2, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 3, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 4, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 5, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{
        "market":*(market ).At(MkString("id") ),
        "side":side ,
        "type":type_ ,
        "size":parseFloat(this.AmountToPrecision(symbol , amount )),
        }); _ = request;
  clientOrderId := this.SafeString2(params , MkString("clientId") , MkString("clientOrderId") ); _ = clientOrderId;
  if IsTrue(OpNotEq2(clientOrderId , MkUndefined() )) {
    *(request ).At(MkString("clientId") )= OpCopy(clientOrderId );
    params = this.Omit(params , MkArray(&VarArray{
            MkString("clientId") ,
            MkString("clientOrderId") ,
            }));
  }
  method := OpCopy(MkUndefined() ); _ = method;
  if IsTrue(OpEq2(type_ , MkString("limit") )) {
    method = MkString("privatePostOrders") ;
    *(request ).At(MkString("price") )= parseFloat(this.PriceToPrecision(symbol , price ));
  } else {
    if IsTrue(OpEq2(type_ , MkString("market") )) {
      method = MkString("privatePostOrders") ;
      *(request ).At(MkString("price") )= OpCopy(MkUndefined() );
    } else {
      if IsTrue(OpOr(OpEq2(type_ , MkString("stop") ), OpEq2(type_ , MkString("takeProfit") ))) {
        method = MkString("privatePostConditionalOrders") ;
        stopPrice := this.SafeNumber2(params , MkString("stopPrice") , MkString("triggerPrice") ); _ = stopPrice;
        if IsTrue(OpEq2(stopPrice , MkUndefined() )) {
          panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" createOrder () requires a stopPrice parameter or a triggerPrice parameter for ") , OpAdd(type_ , MkString(" orders") )))));
        } else {
          params = this.Omit(params , MkArray(&VarArray{
                  MkString("stopPrice") ,
                  MkString("triggerPrice") ,
                  }));
          *(request ).At(MkString("triggerPrice") )= parseFloat(this.PriceToPrecision(symbol , stopPrice ));
        }
        if IsTrue(OpNotEq2(price , MkUndefined() )) {
          *(request ).At(MkString("orderPrice") )= parseFloat(this.PriceToPrecision(symbol , price ));
        }
      } else {
        if IsTrue(OpEq2(type_ , MkString("trailingStop") )) {
          method = MkString("privatePostConditionalOrders") ;
          *(request ).At(MkString("trailValue") )= parseFloat(this.PriceToPrecision(symbol , price ));
        } else {
          panic(NewInvalidOrder(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" createOrder () does not support order type ") , OpAdd(type_ , MkString(", only limit, market, stop, trailingStop, or takeProfit orders are supported") )))));
        }
      }
    }
  }
  response := this.Call(method , this.Extend(request , params )); _ = response;
  result := this.SafeValue(response , MkString("result") , MkArray(&VarArray{})); _ = result;
  return this.ParseOrder(result , market );
}

func (this *Ftx) EditOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined()); _ = symbol;
  type_ := GoGetArg(goArgs, 2, MkUndefined()); _ = type_;
  side := GoGetArg(goArgs, 3, MkUndefined()); _ = side;
  amount := GoGetArg(goArgs, 4, MkUndefined()); _ = amount;
  price := GoGetArg(goArgs, 5, MkUndefined() ); _ = price;
  params := GoGetArg(goArgs, 6, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := this.Market(symbol ); _ = market;
  request := MkMap(&VarMap{}); _ = request;
  method := OpCopy(MkUndefined() ); _ = method;
  clientOrderId := this.SafeString2(params , MkString("client_order_id") , MkString("clientOrderId") ); _ = clientOrderId;
  triggerPrice := this.SafeNumber(params , MkString("triggerPrice") ); _ = triggerPrice;
  orderPrice := this.SafeNumber(params , MkString("orderPrice") ); _ = orderPrice;
  trailValue := this.SafeNumber(params , MkString("trailValue") ); _ = trailValue;
  params = this.Omit(params , MkArray(&VarArray{
          MkString("client_order_id") ,
          MkString("clientOrderId") ,
          MkString("triggerPrice") ,
          MkString("orderPrice") ,
          MkString("trailValue") ,
          }));
  triggerPriceIsDefined := OpNotEq2(triggerPrice , MkUndefined() ); _ = triggerPriceIsDefined;
  orderPriceIsDefined := OpNotEq2(orderPrice , MkUndefined() ); _ = orderPriceIsDefined;
  trailValueIsDefined := OpNotEq2(trailValue , MkUndefined() ); _ = trailValueIsDefined;
  if IsTrue(OpOr(triggerPriceIsDefined , OpOr(orderPriceIsDefined , trailValueIsDefined ))) {
    method = MkString("privatePostConditionalOrdersOrderIdModify") ;
    *(request ).At(MkString("order_id") )= OpCopy(id );
    if IsTrue(triggerPriceIsDefined ) {
      *(request ).At(MkString("triggerPrice") )= parseFloat(this.PriceToPrecision(symbol , triggerPrice ));
    }
    if IsTrue(orderPriceIsDefined ) {
      *(request ).At(MkString("orderPrice") )= parseFloat(this.PriceToPrecision(symbol , orderPrice ));
    }
    if IsTrue(trailValueIsDefined ) {
      *(request ).At(MkString("trailValue") )= parseFloat(this.PriceToPrecision(symbol , trailValue ));
    }
  } else {
    if IsTrue(OpEq2(clientOrderId , MkUndefined() )) {
      method = MkString("privatePostOrdersOrderIdModify") ;
      *(request ).At(MkString("order_id") )= OpCopy(id );
    } else {
      method = MkString("privatePostOrdersByClientIdClientOrderIdModify") ;
      *(request ).At(MkString("client_order_id") )= OpCopy(clientOrderId );
    }
    if IsTrue(OpNotEq2(price , MkUndefined() )) {
      *(request ).At(MkString("price") )= parseFloat(this.PriceToPrecision(symbol , price ));
    }
  }
  if IsTrue(OpNotEq2(amount , MkUndefined() )) {
    *(request ).At(MkString("size") )= parseFloat(this.AmountToPrecision(symbol , amount ));
  }
  response := this.Call(method , this.Extend(request , params )); _ = response;
  result := this.SafeValue(response , MkString("result") , MkMap(&VarMap{})); _ = result;
  return this.ParseOrder(result , market );
}

func (this *Ftx) CancelOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  options := this.SafeValue(*this.At(MkString("options")) , MkString("cancelOrder") , MkMap(&VarMap{})); _ = options;
  defaultMethod := this.SafeString(options , MkString("method") , MkString("privateDeleteOrdersOrderId") ); _ = defaultMethod;
  method := this.SafeString(params , MkString("method") , defaultMethod ); _ = method;
  type_ := this.SafeValue(params , MkString("type") ); _ = type_;
  clientOrderId := this.SafeValue2(params , MkString("client_order_id") , MkString("clientOrderId") ); _ = clientOrderId;
  if IsTrue(OpEq2(clientOrderId , MkUndefined() )) {
    *(request ).At(MkString("order_id") )= parseInt(id );
    if IsTrue(OpOr(OpEq2(type_ , MkString("stop") ), OpOr(OpEq2(type_ , MkString("trailingStop") ), OpEq2(type_ , MkString("takeProfit") )))) {
      method = MkString("privateDeleteConditionalOrdersOrderId") ;
    }
  } else {
    *(request ).At(MkString("client_order_id") )= OpCopy(clientOrderId );
    method = MkString("privateDeleteOrdersByClientIdClientOrderId") ;
  }
  query := this.Omit(params , MkArray(&VarArray{
            MkString("method") ,
            MkString("type") ,
            MkString("client_order_id") ,
            MkString("clientOrderId") ,
            })); _ = query;
  response := this.Call(method , this.Extend(request , query )); _ = response;
  result := this.SafeValue(response , MkString("result") , MkMap(&VarMap{})); _ = result;
  return result ;
}

func (this *Ftx) CancelAllOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  marketId := this.GetMarketId(symbol , MkString("market") , params ); _ = marketId;
  if IsTrue(OpNotEq2(marketId , MkUndefined() )) {
    *(request ).At(MkString("market") )= OpCopy(marketId );
  }
  response := this.Call(MkString("privateDeleteOrders"), this.Extend(request , params )); _ = response;
  result := this.SafeValue(response , MkString("result") , MkMap(&VarMap{})); _ = result;
  return result ;
}

func (this *Ftx) FetchOrder(goArgs ...*Variant) *Variant{
  id := GoGetArg(goArgs, 0, MkUndefined()); _ = id;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  clientOrderId := this.SafeValue2(params , MkString("client_order_id") , MkString("clientOrderId") ); _ = clientOrderId;
  method := MkString("privateGetOrdersOrderId") ; _ = method;
  if IsTrue(OpEq2(clientOrderId , MkUndefined() )) {
    *(request ).At(MkString("order_id") )= OpCopy(id );
  } else {
    *(request ).At(MkString("client_order_id") )= OpCopy(clientOrderId );
    params = this.Omit(params , MkArray(&VarArray{
            MkString("client_order_id") ,
            MkString("clientOrderId") ,
            }));
    method = MkString("privateGetOrdersByClientIdClientOrderId") ;
  }
  response := this.Call(method , this.Extend(request , params )); _ = response;
  result := this.SafeValue(response , MkString("result") , MkMap(&VarMap{})); _ = result;
  return this.ParseOrder(result );
}

func (this *Ftx) FetchOpenOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  market := MkUndefined(); _ = market;
  marketId := MkUndefined(); _ = marketId;
  ArrayUnpack(this.GetMarketParams(symbol , MkString("market") , params ), &market, &marketId);
  if IsTrue(OpNotEq2(marketId , MkUndefined() )) {
    *(request ).At(MkString("market") )= OpCopy(marketId );
  }
  options := this.SafeValue(*this.At(MkString("options")) , MkString("fetchOpenOrders") , MkMap(&VarMap{})); _ = options;
  defaultMethod := this.SafeString(options , MkString("method") , MkString("privateGetOrders") ); _ = defaultMethod;
  method := this.SafeString(params , MkString("method") , defaultMethod ); _ = method;
  type_ := this.SafeValue(params , MkString("type") ); _ = type_;
  if IsTrue(OpOr(OpEq2(type_ , MkString("stop") ), OpOr(OpEq2(type_ , MkString("trailingStop") ), OpEq2(type_ , MkString("takeProfit") )))) {
    method = MkString("privateGetConditionalOrders") ;
  }
  query := this.Omit(params , MkArray(&VarArray{
            MkString("method") ,
            MkString("type") ,
            })); _ = query;
  response := this.Call(method , this.Extend(request , query )); _ = response;
  result := this.SafeValue(response , MkString("result") , MkArray(&VarArray{})); _ = result;
  return this.ParseOrders(result , market , since , limit );
}

func (this *Ftx) FetchOrders(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  market := MkUndefined(); _ = market;
  marketId := MkUndefined(); _ = marketId;
  ArrayUnpack(this.GetMarketParams(symbol , MkString("market") , params ), &market, &marketId);
  if IsTrue(OpNotEq2(marketId , MkUndefined() )) {
    *(request ).At(MkString("market") )= OpCopy(marketId );
  }
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("start_time") )= parseInt(OpDiv(since , MkInteger(1000) ));
  }
  options := this.SafeValue(*this.At(MkString("options")) , MkString("fetchOrders") , MkMap(&VarMap{})); _ = options;
  defaultMethod := this.SafeString(options , MkString("method") , MkString("privateGetOrdersHistory") ); _ = defaultMethod;
  method := this.SafeString(params , MkString("method") , defaultMethod ); _ = method;
  type_ := this.SafeValue(params , MkString("type") ); _ = type_;
  if IsTrue(OpOr(OpEq2(type_ , MkString("stop") ), OpOr(OpEq2(type_ , MkString("trailingStop") ), OpEq2(type_ , MkString("takeProfit") )))) {
    method = MkString("privateGetConditionalOrdersHistory") ;
  }
  query := this.Omit(params , MkArray(&VarArray{
            MkString("method") ,
            MkString("type") ,
            })); _ = query;
  response := this.Call(method , this.Extend(request , query )); _ = response;
  result := this.SafeValue(response , MkString("result") , MkArray(&VarArray{})); _ = result;
  return this.ParseOrders(result , market , since , limit );
}

func (this *Ftx) FetchMyTrades(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  market := MkUndefined(); _ = market;
  marketId := MkUndefined(); _ = marketId;
  ArrayUnpack(this.GetMarketParams(symbol , MkString("market") , params ), &market, &marketId);
  request := MkMap(&VarMap{}); _ = request;
  if IsTrue(OpNotEq2(marketId , MkUndefined() )) {
    *(request ).At(MkString("market") )= OpCopy(marketId );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("start_time") )= parseInt(OpDiv(since , MkInteger(1000) ));
    *(request ).At(MkString("end_time") )= this.Seconds();
  }
  response := this.Call(MkString("privateGetFills"), this.Extend(request , params )); _ = response;
  trades := this.SafeValue(response , MkString("result") , MkArray(&VarArray{})); _ = trades;
  return this.ParseTrades(trades , market , since , limit );
}

func (this *Ftx) Withdraw(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  amount := GoGetArg(goArgs, 1, MkUndefined()); _ = amount;
  address := GoGetArg(goArgs, 2, MkUndefined()); _ = address;
  tag := GoGetArg(goArgs, 3, MkUndefined() ); _ = tag;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  this.CheckAddress(address );
  currency := this.Currency(code ); _ = currency;
  request := MkMap(&VarMap{
        "coin":*(currency ).At(MkString("id") ),
        "size":amount ,
        "address":address ,
        }); _ = request;
  if IsTrue(OpNotEq2(*this.At(MkString("password")) , MkUndefined() )) {
    *(request ).At(MkString("password") )= OpCopy(*this.At(MkString("password")) );
  }
  if IsTrue(OpNotEq2(tag , MkUndefined() )) {
    *(request ).At(MkString("tag") )= OpCopy(tag );
  }
  response := this.Call(MkString("privatePostWalletWithdrawals"), this.Extend(request , params )); _ = response;
  result := this.SafeValue(response , MkString("result") , MkMap(&VarMap{})); _ = result;
  return this.ParseTransaction(result , currency );
}

func (this *Ftx) FetchPositions(goArgs ...*Variant) *Variant{
  symbols := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbols;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{}); _ = request;
  response := this.Call(MkString("privateGetPositions"), this.Extend(request , params )); _ = response;
  return this.SafeValue(response , MkString("result") , MkArray(&VarArray{}));
}

func (this *Ftx) FetchAccountPositions(goArgs ...*Variant) *Variant{
  symbols := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbols;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("privateGetAccount"), params ); _ = response;
  result := this.SafeValue(response , MkString("result") , MkMap(&VarMap{})); _ = result;
  return this.SafeValue(result , MkString("positions") , MkArray(&VarArray{}));
}

func (this *Ftx) FetchDepositAddress(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  currency := this.Currency(code ); _ = currency;
  request := MkMap(&VarMap{"coin":*(currency ).At(MkString("id") )}); _ = request;
  response := this.Call(MkString("privateGetWalletDepositAddressCoin"), this.Extend(request , params )); _ = response;
  result := this.SafeValue(response , MkString("result") , MkMap(&VarMap{})); _ = result;
  address := this.SafeString(result , MkString("address") ); _ = address;
  tag := this.SafeString(result , MkString("tag") ); _ = tag;
  this.CheckAddress(address );
  return MkMap(&VarMap{
        "currency":code ,
        "address":address ,
        "tag":tag ,
        "info":response ,
        });
}

func (this *Ftx) ParseTransactionStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "confirmed":MkString("ok") ,
        "complete":MkString("ok") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Ftx) ParseTransaction(goArgs ...*Variant) *Variant{
  transaction := GoGetArg(goArgs, 0, MkUndefined()); _ = transaction;
  currency := GoGetArg(goArgs, 1, MkUndefined() ); _ = currency;
  code := this.SafeCurrencyCode(this.SafeString(transaction , MkString("coin") )); _ = code;
  id := this.SafeString(transaction , MkString("id") ); _ = id;
  amount := this.SafeNumber(transaction , MkString("size") ); _ = amount;
  status := this.ParseTransactionStatus(this.SafeString(transaction , MkString("status") )); _ = status;
  timestamp := this.Parse8601(this.SafeString(transaction , MkString("time") )); _ = timestamp;
  txid := this.SafeString(transaction , MkString("txid") ); _ = txid;
  tag := OpCopy(MkUndefined() ); _ = tag;
  address := this.SafeValue(transaction , MkString("address") ); _ = address;
  if IsTrue(OpNotEq2(OpType(address ), MkString("string") )) {
    tag = this.SafeString(address , MkString("tag") );
    address = this.SafeString(address , MkString("address") );
  }
  if IsTrue(OpEq2(address , MkUndefined() )) {
    notes := this.SafeString(transaction , MkString("notes") ); _ = notes;
    if IsTrue(OpAnd(OpNotEq2(notes , MkUndefined() ), OpGtEq(notes.IndexOf(MkString("Transfer to") ), MkInteger(0) ))) {
      address = notes.Slice(MkInteger(12) );
    }
  }
  fee := this.SafeNumber(transaction , MkString("fee") ); _ = fee;
  return MkMap(&VarMap{
        "info":transaction ,
        "id":id ,
        "txid":txid ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "addressFrom":MkUndefined() ,
        "address":address ,
        "addressTo":address ,
        "tagFrom":MkUndefined() ,
        "tag":tag ,
        "tagTo":tag ,
        "type":MkUndefined() ,
        "amount":amount ,
        "currency":code ,
        "status":status ,
        "updated":MkUndefined() ,
        "fee":MkMap(&VarMap{
            "currency":code ,
            "cost":fee ,
            "rate":MkUndefined() ,
            }),
        });
}

func (this *Ftx) FetchDeposits(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("privateGetWalletDeposits"), params ); _ = response;
  result := this.SafeValue(response , MkString("result") , MkArray(&VarArray{})); _ = result;
  currency := OpCopy(MkUndefined() ); _ = currency;
  if IsTrue(OpNotEq2(code , MkUndefined() )) {
    currency = this.Currency(code );
  }
  return this.ParseTransactions(result , currency , since , limit , MkMap(&VarMap{"type":MkString("deposit") }));
}

func (this *Ftx) FetchWithdrawals(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  response := this.Call(MkString("privateGetWalletWithdrawals"), params ); _ = response;
  result := this.SafeValue(response , MkString("result") , MkArray(&VarArray{})); _ = result;
  currency := OpCopy(MkUndefined() ); _ = currency;
  if IsTrue(OpNotEq2(code , MkUndefined() )) {
    currency = this.Currency(code );
  }
  return this.ParseTransactions(result , currency , since , limit , MkMap(&VarMap{"type":MkString("withdrawal") }));
}

func (this *Ftx) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("public") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  request := OpAdd(MkString("/api/") , this.ImplodeParams(path , params )); _ = request;
  query := this.Omit(params , this.ExtractParams(path )); _ = query;
  baseUrl := this.ImplodeHostname(*(*(*this.At(MkString("urls")) ).At(MkString("api") )).At(api )); _ = baseUrl;
  url := OpAdd(baseUrl , request ); _ = url;
  if IsTrue(OpNotEq2(method , MkString("POST") )) {
    if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
      suffix := OpAdd(MkString("?") , this.Urlencode(query )); _ = suffix;
      OpAddAssign(& url , suffix );
      OpAddAssign(& request , suffix );
    }
  }
  if IsTrue(OpEq2(api , MkString("private") )) {
    this.CheckRequiredCredentials();
    timestamp := (this.Milliseconds()).Call(MkString("toString") ); _ = timestamp;
    auth := OpAdd(timestamp , OpAdd(method , request )); _ = auth;
    headers = MkMap(&VarMap{});
    if IsTrue(OpOr(OpEq2(method , MkString("POST") ), OpEq2(method , MkString("DELETE") ))) {
      body = this.Json(query );
      OpAddAssign(& auth , body );
      *(headers ).At(MkString("Content-Type") )= MkString("application/json") ;
    }
    signature := this.Hmac(this.Encode(auth ), this.Encode(*this.At(MkString("secret")) ), MkString("sha256") ); _ = signature;
    options := this.SafeValue(*this.At(MkString("options")) , MkString("sign") , MkMap(&VarMap{})); _ = options;
    headerPrefix := this.SafeString(options , *this.At(MkString("hostname")) , MkString("FTX") ); _ = headerPrefix;
    keyField := OpAdd(headerPrefix , MkString("-KEY") ); _ = keyField;
    tsField := OpAdd(headerPrefix , MkString("-TS") ); _ = tsField;
    signField := OpAdd(headerPrefix , MkString("-SIGN") ); _ = signField;
    *(headers ).At(keyField )= OpCopy(*this.At(MkString("apiKey")) );
    *(headers ).At(tsField )= OpCopy(timestamp );
    *(headers ).At(signField )= OpCopy(signature );
  }
  return MkMap(&VarMap{
        "url":url ,
        "method":method ,
        "body":body ,
        "headers":headers ,
        });
}

func (this *Ftx) HandleErrors(goArgs ...*Variant) *Variant{
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
  success := this.SafeValue(response , MkString("success") ); _ = success;
  if IsTrue(OpNot(success )) {
    feedback := OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , body )); _ = feedback;
    error := this.SafeString(response , MkString("error") ); _ = error;
    this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("exact") ), error , feedback );
    this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("broad") ), error , feedback );
    panic(NewExchangeError(feedback ));
  }
  return MkUndefined()
}

func (this *Ftx) SetLeverage(goArgs ...*Variant) *Variant{
  leverage := GoGetArg(goArgs, 0, MkUndefined()); _ = leverage;
  symbol := GoGetArg(goArgs, 1, MkUndefined() ); _ = symbol;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  if IsTrue(OpOr(OpLw(leverage , MkInteger(1) ), OpGt(leverage , MkInteger(20) ))) {
    panic(NewBadRequest(OpAdd(*this.At(MkString("id")) , MkString(" leverage should be between 1 and 20") )));
  }
  request := MkMap(&VarMap{"leverage":leverage }); _ = request;
  return this.Call(MkString("privatePostAccountLeverage"), this.Extend(request , params ));
}

func (this *Ftx) ParseIncome(goArgs ...*Variant) *Variant{
  income := GoGetArg(goArgs, 0, MkUndefined()); _ = income;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  marketId := this.SafeString(income , MkString("future") ); _ = marketId;
  symbol := this.SafeSymbol(marketId , market ); _ = symbol;
  amount := this.SafeNumber(income , MkString("payment") ); _ = amount;
  code := this.SafeCurrencyCode(MkString("USD") ); _ = code;
  id := this.SafeString(income , MkString("id") ); _ = id;
  timestamp := this.SafeInteger(income , MkString("time") ); _ = timestamp;
  rate := this.Call(MkString("safe_number"), income , MkString("rate") ); _ = rate;
  return MkMap(&VarMap{
        "info":income ,
        "symbol":symbol ,
        "code":code ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "id":id ,
        "amount":amount ,
        "rate":rate ,
        });
}

func (this *Ftx) ParseIncomes(goArgs ...*Variant) *Variant{
  incomes := GoGetArg(goArgs, 0, MkUndefined()); _ = incomes;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , incomes.Length )); OpIncr(& i ){
    entry := *(incomes ).At(i ); _ = entry;
    parsed := this.ParseIncome(entry , market ); _ = parsed;
    result.Push(parsed );
  }
  return this.FilterBySinceLimit(result , since , limit , MkString("timestamp") );
}

func (this *Ftx) FetchFundingHistory(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  method := MkString("private_get_funding_payments") ; _ = method;
  request := MkMap(&VarMap{}); _ = request;
  market := OpCopy(MkUndefined() ); _ = market;
  if IsTrue(OpNotEq2(symbol , MkUndefined() )) {
    market = this.Market(symbol );
    *(request ).At(MkString("future") )= *(market ).At(MkString("id") );
  }
  if IsTrue(OpNotEq2(since , MkUndefined() )) {
    *(request ).At(MkString("startTime") )= OpCopy(since );
  }
  response := this.Call(method , this.Extend(request , params )); _ = response;
  return this.ParseIncomes(response , market , since , limit );
}

