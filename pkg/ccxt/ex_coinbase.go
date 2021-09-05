package ccxt

import (
)

type Coinbase struct {
	*ExchangeBase
}

var _ Exchange = (*Coinbase)(nil)

func init() {
	exchange := &Coinbase{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Coinbase) Describe(goArgs ...*Variant) *Variant{
  return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
            "id":MkString("coinbase") ,
            "name":MkString("Coinbase") ,
            "countries":MkArray(&VarArray{
                MkString("US") ,
                }),
            "rateLimit":MkInteger(400) ,
            "version":MkString("v2") ,
            "userAgent":*(*this.At(MkString("userAgents")) ).At(MkString("chrome") ),
            "headers":MkMap(&VarMap{"CB-VERSION":MkString("2018-05-30") }),
            "has":MkMap(&VarMap{
                "CORS":MkBool(true) ,
                "cancelOrder":MkBool(false) ,
                "createDepositAddress":MkBool(true) ,
                "createOrder":MkBool(false) ,
                "deposit":MkBool(false) ,
                "fetchBalance":MkBool(true) ,
                "fetchClosedOrders":MkBool(false) ,
                "fetchCurrencies":MkBool(true) ,
                "fetchDepositAddress":MkBool(false) ,
                "fetchMarkets":MkBool(true) ,
                "fetchMyTrades":MkBool(false) ,
                "fetchOHLCV":MkBool(false) ,
                "fetchOpenOrders":MkBool(false) ,
                "fetchOrder":MkBool(false) ,
                "fetchOrderBook":MkBool(false) ,
                "fetchL2OrderBook":MkBool(false) ,
                "fetchLedger":MkBool(true) ,
                "fetchOrders":MkBool(false) ,
                "fetchTicker":MkBool(true) ,
                "fetchTickers":MkBool(false) ,
                "fetchTime":MkBool(true) ,
                "fetchBidsAsks":MkBool(false) ,
                "fetchTrades":MkBool(false) ,
                "withdraw":MkBool(false) ,
                "fetchTransactions":MkBool(false) ,
                "fetchDeposits":MkBool(true) ,
                "fetchWithdrawals":MkBool(true) ,
                "fetchMySells":MkBool(true) ,
                "fetchMyBuys":MkBool(true) ,
                }),
            "urls":MkMap(&VarMap{
                "logo":MkString("https://user-images.githubusercontent.com/1294454/40811661-b6eceae2-653a-11e8-829e-10bfadb078cf.jpg") ,
                "api":MkString("https://api.coinbase.com") ,
                "www":MkString("https://www.coinbase.com") ,
                "doc":MkString("https://developers.coinbase.com/api/v2") ,
                "fees":MkString("https://support.coinbase.com/customer/portal/articles/2109597-buy-sell-bank-transfer-fees") ,
                "referral":MkString("https://www.coinbase.com/join/58cbe25a355148797479dbd2") ,
                }),
            "requiredCredentials":MkMap(&VarMap{
                "apiKey":MkBool(true) ,
                "secret":MkBool(true) ,
                }),
            "api":MkMap(&VarMap{
                "public":MkMap(&VarMap{"get":MkArray(&VarArray{
                        MkString("currencies") ,
                        MkString("time") ,
                        MkString("exchange-rates") ,
                        MkString("users/{user_id}") ,
                        MkString("prices/{symbol}/buy") ,
                        MkString("prices/{symbol}/sell") ,
                        MkString("prices/{symbol}/spot") ,
                        })}),
                "private":MkMap(&VarMap{
                    "get":MkArray(&VarArray{
                        MkString("accounts") ,
                        MkString("accounts/{account_id}") ,
                        MkString("accounts/{account_id}/addresses") ,
                        MkString("accounts/{account_id}/addresses/{address_id}") ,
                        MkString("accounts/{account_id}/addresses/{address_id}/transactions") ,
                        MkString("accounts/{account_id}/transactions") ,
                        MkString("accounts/{account_id}/transactions/{transaction_id}") ,
                        MkString("accounts/{account_id}/buys") ,
                        MkString("accounts/{account_id}/buys/{buy_id}") ,
                        MkString("accounts/{account_id}/sells") ,
                        MkString("accounts/{account_id}/sells/{sell_id}") ,
                        MkString("accounts/{account_id}/deposits") ,
                        MkString("accounts/{account_id}/deposits/{deposit_id}") ,
                        MkString("accounts/{account_id}/withdrawals") ,
                        MkString("accounts/{account_id}/withdrawals/{withdrawal_id}") ,
                        MkString("payment-methods") ,
                        MkString("payment-methods/{payment_method_id}") ,
                        MkString("user") ,
                        MkString("user/auth") ,
                        }),
                    "post":MkArray(&VarArray{
                        MkString("accounts") ,
                        MkString("accounts/{account_id}/primary") ,
                        MkString("accounts/{account_id}/addresses") ,
                        MkString("accounts/{account_id}/transactions") ,
                        MkString("accounts/{account_id}/transactions/{transaction_id}/complete") ,
                        MkString("accounts/{account_id}/transactions/{transaction_id}/resend") ,
                        MkString("accounts/{account_id}/buys") ,
                        MkString("accounts/{account_id}/buys/{buy_id}/commit") ,
                        MkString("accounts/{account_id}/sells") ,
                        MkString("accounts/{account_id}/sells/{sell_id}/commit") ,
                        MkString("accounts/{account_id}/deposists") ,
                        MkString("accounts/{account_id}/deposists/{deposit_id}/commit") ,
                        MkString("accounts/{account_id}/withdrawals") ,
                        MkString("accounts/{account_id}/withdrawals/{withdrawal_id}/commit") ,
                        }),
                    "put":MkArray(&VarArray{
                        MkString("accounts/{account_id}") ,
                        MkString("user") ,
                        }),
                    "delete":MkArray(&VarArray{
                        MkString("accounts/{id}") ,
                        MkString("accounts/{account_id}/transactions/{transaction_id}") ,
                        }),
                    }),
                }),
            "exceptions":MkMap(&VarMap{
                "exact":MkMap(&VarMap{
                    "two_factor_required":AuthenticationError ,
                    "param_required":ExchangeError ,
                    "validation_error":ExchangeError ,
                    "invalid_request":ExchangeError ,
                    "personal_details_required":AuthenticationError ,
                    "identity_verification_required":AuthenticationError ,
                    "jumio_verification_required":AuthenticationError ,
                    "jumio_face_match_verification_required":AuthenticationError ,
                    "unverified_email":AuthenticationError ,
                    "authentication_error":AuthenticationError ,
                    "invalid_token":AuthenticationError ,
                    "revoked_token":AuthenticationError ,
                    "expired_token":AuthenticationError ,
                    "invalid_scope":AuthenticationError ,
                    "not_found":ExchangeError ,
                    "rate_limit_exceeded":RateLimitExceeded ,
                    "internal_server_error":ExchangeError ,
                    }),
                "broad":MkMap(&VarMap{"request timestamp expired":InvalidNonce }),
                }),
            "commonCurrencies":MkMap(&VarMap{"CGLD":MkString("CELO") }),
            "options":MkMap(&VarMap{
                "fetchCurrencies":MkMap(&VarMap{"expires":MkInteger(5000) }),
                "accounts":MkArray(&VarArray{
                    MkString("wallet") ,
                    MkString("fiat") ,
                    }),
                }),
            }));
}

func (this *Coinbase) FetchTime(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.Call(MkString("publicGetTime"), params ); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  return this.SafeTimestamp(data , MkString("epoch") );
}

func (this *Coinbase) FetchAccounts(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"limit":MkInteger(100) }); _ = request;
  response := this.Call(MkString("privateGetAccounts"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkArray(&VarArray{})); _ = data;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , data.Length )); OpIncr(& i ){
    account := *(data ).At(i ); _ = account;
    currency := this.SafeValue(account , MkString("currency") , MkMap(&VarMap{})); _ = currency;
    currencyId := this.SafeString(currency , MkString("code") ); _ = currencyId;
    code := this.SafeCurrencyCode(currencyId ); _ = code;
    result.Push(MkMap(&VarMap{
            "id":this.SafeString(account , MkString("id") ),
            "type":this.SafeString(account , MkString("type") ),
            "code":code ,
            "info":account ,
            }));
  }
  return result ;
}

func (this *Coinbase) CreateDepositAddress(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  accountId := this.SafeString(params , MkString("account_id") ); _ = accountId;
  params = this.Omit(params , MkString("account_id") );
  if IsTrue(OpEq2(accountId , MkUndefined() )) {
    this.LoadAccounts();
    for i := MkInteger(0) ; IsTrue(OpLw(i , (*((this).At(MkString("accounts")))).Length )); OpIncr(& i ){
      account := *(*this.At(MkString("accounts")) ).At(i ); _ = account;
      if IsTrue(OpAnd(OpEq2(*(account ).At(MkString("code") ), code ), OpEq2(*(account ).At(MkString("type") ), MkString("wallet") ))) {
        accountId = *(account ).At(MkString("id") );
        break ;
      }
    }
  }
  if IsTrue(OpEq2(accountId , MkUndefined() )) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , MkString(" createDepositAddress could not find the account with matching currency code, specify an `account_id` extra param") )));
  }
  request := MkMap(&VarMap{"account_id":accountId }); _ = request;
  response := this.Call(MkString("privatePostAccountsAccountIdAddresses"), this.Extend(request , params )); _ = response;
  data := this.SafeValue(response , MkString("data") , MkMap(&VarMap{})); _ = data;
  tag := this.SafeString(data , MkString("destination_tag") ); _ = tag;
  address := this.SafeString(data , MkString("address") ); _ = address;
  return MkMap(&VarMap{
        "currency":code ,
        "tag":tag ,
        "address":address ,
        "info":response ,
        });
}

func (this *Coinbase) FetchMySells(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  request := this.PrepareAccountRequest(limit , params ); _ = request;
  this.LoadMarkets();
  query := this.Omit(params , MkArray(&VarArray{
            MkString("account_id") ,
            MkString("accountId") ,
            })); _ = query;
  sells := this.Call(MkString("privateGetAccountsAccountIdSells"), this.Extend(request , query )); _ = sells;
  return this.ParseTrades(*(sells ).At(MkString("data") ), MkUndefined() , since , limit );
}

func (this *Coinbase) FetchMyBuys(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined() ); _ = symbol;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  request := this.PrepareAccountRequest(limit , params ); _ = request;
  this.LoadMarkets();
  query := this.Omit(params , MkArray(&VarArray{
            MkString("account_id") ,
            MkString("accountId") ,
            })); _ = query;
  buys := this.Call(MkString("privateGetAccountsAccountIdBuys"), this.Extend(request , query )); _ = buys;
  return this.ParseTrades(*(buys ).At(MkString("data") ), MkUndefined() , since , limit );
}

func (this *Coinbase) FetchTransactionsWithMethod(goArgs ...*Variant) *Variant{
  method := GoGetArg(goArgs, 0, MkUndefined()); _ = method;
  code := GoGetArg(goArgs, 1, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 2, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 3, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 4, MkMap(&VarMap{})); _ = params;
  request := this.PrepareAccountRequestWithCurrencyCode(code , limit , params ); _ = request;
  this.LoadMarkets();
  query := this.Omit(params , MkArray(&VarArray{
            MkString("account_id") ,
            MkString("accountId") ,
            })); _ = query;
  response := this.Call(method , this.Extend(request , query )); _ = response;
  return this.ParseTransactions(*(response ).At(MkString("data") ), MkUndefined() , since , limit );
}

func (this *Coinbase) FetchWithdrawals(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  return this.FetchTransactionsWithMethod(MkString("privateGetAccountsAccountIdWithdrawals") , code , since , limit , params );
}

func (this *Coinbase) FetchDeposits(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  return this.FetchTransactionsWithMethod(MkString("privateGetAccountsAccountIdDeposits") , code , since , limit , params );
}

func (this *Coinbase) ParseTransactionStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  statuses := MkMap(&VarMap{
        "created":MkString("pending") ,
        "completed":MkString("ok") ,
        "canceled":MkString("canceled") ,
        }); _ = statuses;
  return this.SafeString(statuses , status , status );
}

func (this *Coinbase) ParseTransaction(goArgs ...*Variant) *Variant{
  transaction := GoGetArg(goArgs, 0, MkUndefined()); _ = transaction;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  subtotalObject := this.SafeValue(transaction , MkString("subtotal") , MkMap(&VarMap{})); _ = subtotalObject;
  feeObject := this.SafeValue(transaction , MkString("fee") , MkMap(&VarMap{})); _ = feeObject;
  id := this.SafeString(transaction , MkString("id") ); _ = id;
  timestamp := this.Parse8601(this.SafeValue(transaction , MkString("created_at") )); _ = timestamp;
  updated := this.Parse8601(this.SafeValue(transaction , MkString("updated_at") )); _ = updated;
  type_ := this.SafeString(transaction , MkString("resource") ); _ = type_;
  amount := this.SafeNumber(subtotalObject , MkString("amount") ); _ = amount;
  currencyId := this.SafeString(subtotalObject , MkString("currency") ); _ = currencyId;
  currency := this.SafeCurrencyCode(currencyId ); _ = currency;
  feeCost := this.SafeNumber(feeObject , MkString("amount") ); _ = feeCost;
  feeCurrencyId := this.SafeString(feeObject , MkString("currency") ); _ = feeCurrencyId;
  feeCurrency := this.SafeCurrencyCode(feeCurrencyId ); _ = feeCurrency;
  fee := MkMap(&VarMap{
        "cost":feeCost ,
        "currency":feeCurrency ,
        }); _ = fee;
  status := this.ParseTransactionStatus(this.SafeString(transaction , MkString("status") )); _ = status;
  if IsTrue(OpEq2(status , MkUndefined() )) {
    committed := this.SafeValue(transaction , MkString("committed") ); _ = committed;
    status = OpTrinary(committed , MkString("ok") , MkString("pending") );
  }
  return MkMap(&VarMap{
        "info":transaction ,
        "id":id ,
        "txid":id ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "address":MkUndefined() ,
        "tag":MkUndefined() ,
        "type":type_ ,
        "amount":amount ,
        "currency":currency ,
        "status":status ,
        "updated":updated ,
        "fee":fee ,
        });
}

func (this *Coinbase) ParseTrade(goArgs ...*Variant) *Variant{
  trade := GoGetArg(goArgs, 0, MkUndefined()); _ = trade;
  market := GoGetArg(goArgs, 1, MkUndefined() ); _ = market;
  symbol := OpCopy(MkUndefined() ); _ = symbol;
  totalObject := this.SafeValue(trade , MkString("total") , MkMap(&VarMap{})); _ = totalObject;
  amountObject := this.SafeValue(trade , MkString("amount") , MkMap(&VarMap{})); _ = amountObject;
  subtotalObject := this.SafeValue(trade , MkString("subtotal") , MkMap(&VarMap{})); _ = subtotalObject;
  feeObject := this.SafeValue(trade , MkString("fee") , MkMap(&VarMap{})); _ = feeObject;
  id := this.SafeString(trade , MkString("id") ); _ = id;
  timestamp := this.Parse8601(this.SafeValue(trade , MkString("created_at") )); _ = timestamp;
  if IsTrue(OpEq2(market , MkUndefined() )) {
    baseId := this.SafeString(amountObject , MkString("currency") ); _ = baseId;
    quoteId := this.SafeString(totalObject , MkString("currency") ); _ = quoteId;
    if IsTrue(OpAnd(OpNotEq2(baseId , MkUndefined() ), OpNotEq2(quoteId , MkUndefined() ))) {
      base := this.SafeCurrencyCode(baseId ); _ = base;
      quote := this.SafeCurrencyCode(quoteId ); _ = quote;
      symbol = OpAdd(base , OpAdd(MkString("/") , quote ));
    }
  }
  orderId := OpCopy(MkUndefined() ); _ = orderId;
  side := this.SafeString(trade , MkString("resource") ); _ = side;
  type_ := OpCopy(MkUndefined() ); _ = type_;
  costString := this.SafeString(subtotalObject , MkString("amount") ); _ = costString;
  amountString := this.SafeString(amountObject , MkString("amount") ); _ = amountString;
  cost := this.ParseNumber(costString ); _ = cost;
  amount := this.ParseNumber(amountString ); _ = amount;
  price := this.ParseNumber(Precise.StringDiv(costString , amountString )); _ = price;
  feeCost := this.SafeNumber(feeObject , MkString("amount") ); _ = feeCost;
  feeCurrencyId := this.SafeString(feeObject , MkString("currency") ); _ = feeCurrencyId;
  feeCurrency := this.SafeCurrencyCode(feeCurrencyId ); _ = feeCurrency;
  fee := MkMap(&VarMap{
        "cost":feeCost ,
        "currency":feeCurrency ,
        }); _ = fee;
  return MkMap(&VarMap{
        "info":trade ,
        "id":id ,
        "order":orderId ,
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
        });
}

func (this *Coinbase) FetchMarkets(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.FetchCurrenciesFromCache(params ); _ = response;
  currencies := this.SafeValue(response , MkString("currencies") , MkMap(&VarMap{})); _ = currencies;
  exchangeRates := this.SafeValue(response , MkString("exchangeRates") , MkMap(&VarMap{})); _ = exchangeRates;
  data := this.SafeValue(currencies , MkString("data") , MkArray(&VarArray{})); _ = data;
  dataById := this.IndexBy(data , MkString("id") ); _ = dataById;
  rates := this.SafeValue(this.SafeValue(exchangeRates , MkString("data") , MkMap(&VarMap{})), MkString("rates") , MkMap(&VarMap{})); _ = rates;
  baseIds := GoGetKeys(rates ); _ = baseIds;
  result := MkArray(&VarArray{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , baseIds.Length )); OpIncr(& i ){
    baseId := *(baseIds ).At(i ); _ = baseId;
    base := this.SafeCurrencyCode(baseId ); _ = base;
    type_ := OpTrinary(OpHasMember(baseId , dataById ), MkString("fiat") , MkString("crypto") ); _ = type_;
    if IsTrue(OpEq2(type_ , MkString("crypto") )) {
      for j := MkInteger(0) ; IsTrue(OpLw(j , data.Length )); OpIncr(& j ){
        quoteCurrency := *(data ).At(j ); _ = quoteCurrency;
        quoteId := this.SafeString(quoteCurrency , MkString("id") ); _ = quoteId;
        quote := this.SafeCurrencyCode(quoteId ); _ = quote;
        symbol := OpAdd(base , OpAdd(MkString("/") , quote )); _ = symbol;
        id := OpAdd(baseId , OpAdd(MkString("-") , quoteId )); _ = id;
        result.Push(MkMap(&VarMap{
                "id":id ,
                "symbol":symbol ,
                "base":base ,
                "quote":quote ,
                "baseId":baseId ,
                "quoteId":quoteId ,
                "active":MkUndefined() ,
                "info":quoteCurrency ,
                "precision":MkMap(&VarMap{
                    "amount":MkUndefined() ,
                    "price":MkUndefined() ,
                    }),
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
                        "min":this.SafeNumber(quoteCurrency , MkString("min_size") ),
                        "max":MkUndefined() ,
                        }),
                    }),
                }));
      }
    }
  }
  return result ;
}

func (this *Coinbase) FetchCurrenciesFromCache(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  options := this.SafeValue(*this.At(MkString("options")) , MkString("fetchCurrencies") , MkMap(&VarMap{})); _ = options;
  timestamp := this.SafeInteger(options , MkString("timestamp") ); _ = timestamp;
  expires := this.SafeInteger(options , MkString("expires") , MkInteger(1000) ); _ = expires;
  now := this.Milliseconds(); _ = now;
  if IsTrue(OpOr(OpEq2(timestamp , MkUndefined() ), OpGt(OpSub(now , timestamp ), expires ))) {
    currencies := this.Call(MkString("publicGetCurrencies"), params ); _ = currencies;
    exchangeRates := this.Call(MkString("publicGetExchangeRates"), params ); _ = exchangeRates;
    *(*this.At(MkString("options")) ).At(MkString("fetchCurrencies") )= this.Extend(options , MkMap(&VarMap{
            "currencies":currencies ,
            "exchangeRates":exchangeRates ,
            "timestamp":now ,
            }));
  }
  return this.SafeValue(*this.At(MkString("options")) , MkString("fetchCurrencies") , MkMap(&VarMap{}));
}

func (this *Coinbase) FetchCurrencies(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  response := this.FetchCurrenciesFromCache(params ); _ = response;
  currencies := this.SafeValue(response , MkString("currencies") , MkMap(&VarMap{})); _ = currencies;
  exchangeRates := this.SafeValue(response , MkString("exchangeRates") , MkMap(&VarMap{})); _ = exchangeRates;
  data := this.SafeValue(currencies , MkString("data") , MkArray(&VarArray{})); _ = data;
  dataById := this.IndexBy(data , MkString("id") ); _ = dataById;
  rates := this.SafeValue(this.SafeValue(exchangeRates , MkString("data") , MkMap(&VarMap{})), MkString("rates") , MkMap(&VarMap{})); _ = rates;
  keys := GoGetKeys(rates ); _ = keys;
  result := MkMap(&VarMap{}); _ = result;
  for i := MkInteger(0) ; IsTrue(OpLw(i , keys.Length )); OpIncr(& i ){
    key := *(keys ).At(i ); _ = key;
    type_ := OpTrinary(OpHasMember(key , dataById ), MkString("fiat") , MkString("crypto") ); _ = type_;
    currency := this.SafeValue(dataById , key , MkMap(&VarMap{})); _ = currency;
    id := this.SafeString(currency , MkString("id") , key ); _ = id;
    name := this.SafeString(currency , MkString("name") ); _ = name;
    code := this.SafeCurrencyCode(id ); _ = code;
    *(result ).At(code )= MkMap(&VarMap{
        "id":id ,
        "code":code ,
        "info":currency ,
        "type":type_ ,
        "name":name ,
        "active":MkBool(true) ,
        "fee":MkUndefined() ,
        "precision":MkUndefined() ,
        "limits":MkMap(&VarMap{
            "amount":MkMap(&VarMap{
                "min":this.SafeNumber(currency , MkString("min_size") ),
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

func (this *Coinbase) FetchTicker(goArgs ...*Variant) *Variant{
  symbol := GoGetArg(goArgs, 0, MkUndefined()); _ = symbol;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  timestamp := this.Seconds(); _ = timestamp;
  market := this.Market(symbol ); _ = market;
  request := this.Extend(MkMap(&VarMap{"symbol":*(market ).At(MkString("id") )}), params ); _ = request;
  buy := this.Call(MkString("publicGetPricesSymbolBuy"), request ); _ = buy;
  sell := this.Call(MkString("publicGetPricesSymbolSell"), request ); _ = sell;
  spot := this.Call(MkString("publicGetPricesSymbolSpot"), request ); _ = spot;
  ask := this.SafeNumber(*(buy ).At(MkString("data") ), MkString("amount") ); _ = ask;
  bid := this.SafeNumber(*(sell ).At(MkString("data") ), MkString("amount") ); _ = bid;
  last := this.SafeNumber(*(spot ).At(MkString("data") ), MkString("amount") ); _ = last;
  return MkMap(&VarMap{
        "symbol":symbol ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "bid":bid ,
        "ask":ask ,
        "last":last ,
        "high":MkUndefined() ,
        "low":MkUndefined() ,
        "bidVolume":MkUndefined() ,
        "askVolume":MkUndefined() ,
        "vwap":MkUndefined() ,
        "open":MkUndefined() ,
        "close":last ,
        "previousClose":MkUndefined() ,
        "change":MkUndefined() ,
        "percentage":MkUndefined() ,
        "average":MkUndefined() ,
        "baseVolume":MkUndefined() ,
        "quoteVolume":MkUndefined() ,
        "info":MkMap(&VarMap{
            "buy":buy ,
            "sell":sell ,
            "spot":spot ,
            }),
        });
}

func (this *Coinbase) FetchBalance(goArgs ...*Variant) *Variant{
  params := GoGetArg(goArgs, 0, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  request := MkMap(&VarMap{"limit":MkInteger(100) }); _ = request;
  response := this.Call(MkString("privateGetAccounts"), this.Extend(request , params )); _ = response;
  balances := this.SafeValue(response , MkString("data") ); _ = balances;
  accounts := this.SafeValue(params , MkString("type") , *(*this.At(MkString("options")) ).At(MkString("accounts") )); _ = accounts;
  result := MkMap(&VarMap{"info":response }); _ = result;
  for b := MkInteger(0) ; IsTrue(OpLw(b , balances.Length )); OpIncr(& b ){
    balance := *(balances ).At(b ); _ = balance;
    type_ := this.SafeString(balance , MkString("type") ); _ = type_;
    if IsTrue(this.InArray(type_ , accounts )) {
      value := this.SafeValue(balance , MkString("balance") ); _ = value;
      if IsTrue(OpNotEq2(value , MkUndefined() )) {
        currencyId := this.SafeString(value , MkString("currency") ); _ = currencyId;
        code := this.SafeCurrencyCode(currencyId ); _ = code;
        total := this.SafeString(value , MkString("amount") ); _ = total;
        free := OpCopy(total ); _ = free;
        account := this.SafeValue(result , code ); _ = account;
        if IsTrue(OpEq2(account , MkUndefined() )) {
          account = this.Account();
          *(account ).At(MkString("free") )= OpCopy(free );
          *(account ).At(MkString("total") )= OpCopy(total );
        } else {
          *(account ).At(MkString("free") )= Precise.StringAdd(*(account ).At(MkString("free") ), total );
          *(account ).At(MkString("total") )= Precise.StringAdd(*(account ).At(MkString("total") ), total );
        }
        *(result ).At(code )= OpCopy(account );
      }
    }
  }
  return this.ParseBalance(result );
}

func (this *Coinbase) FetchLedger(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  since := GoGetArg(goArgs, 1, MkUndefined() ); _ = since;
  limit := GoGetArg(goArgs, 2, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  this.LoadMarkets();
  currency := OpCopy(MkUndefined() ); _ = currency;
  if IsTrue(OpNotEq2(code , MkUndefined() )) {
    currency = this.Currency(code );
  }
  request := this.PrepareAccountRequestWithCurrencyCode(code , limit , params ); _ = request;
  query := this.Omit(params , MkArray(&VarArray{
            MkString("account_id") ,
            MkString("accountId") ,
            })); _ = query;
  response := this.Call(MkString("privateGetAccountsAccountIdTransactions"), this.Extend(request , query )); _ = response;
  return this.ParseLedger(*(response ).At(MkString("data") ), currency , since , limit );
}

func (this *Coinbase) ParseLedgerEntryStatus(goArgs ...*Variant) *Variant{
  status := GoGetArg(goArgs, 0, MkUndefined()); _ = status;
  types := MkMap(&VarMap{"completed":MkString("ok") }); _ = types;
  return this.SafeString(types , status , status );
}

func (this *Coinbase) ParseLedgerEntryType(goArgs ...*Variant) *Variant{
  type_ := GoGetArg(goArgs, 0, MkUndefined()); _ = type_;
  types := MkMap(&VarMap{
        "buy":MkString("trade") ,
        "sell":MkString("trade") ,
        "fiat_deposit":MkString("transaction") ,
        "fiat_withdrawal":MkString("transaction") ,
        "exchange_deposit":MkString("transaction") ,
        "exchange_withdrawal":MkString("transaction") ,
        "send":MkString("transaction") ,
        "pro_deposit":MkString("transaction") ,
        "pro_withdrawal":MkString("transaction") ,
        }); _ = types;
  return this.SafeString(types , type_ , type_ );
}

func (this *Coinbase) ParseLedgerEntry(goArgs ...*Variant) *Variant{
  item := GoGetArg(goArgs, 0, MkUndefined()); _ = item;
  currency := GoGetArg(goArgs, 1, MkUndefined() ); _ = currency;
  amountInfo := this.SafeValue(item , MkString("amount") , MkMap(&VarMap{})); _ = amountInfo;
  amount := this.SafeNumber(amountInfo , MkString("amount") ); _ = amount;
  direction := OpCopy(MkUndefined() ); _ = direction;
  if IsTrue(OpLw(amount , MkInteger(0) )) {
    direction = MkString("out") ;
    amount = OpNeg(amount );
  } else {
    direction = MkString("in") ;
  }
  currencyId := this.SafeString(amountInfo , MkString("currency") ); _ = currencyId;
  code := this.SafeCurrencyCode(currencyId , currency ); _ = code;
  fee := OpCopy(MkUndefined() ); _ = fee;
  networkInfo := this.SafeValue(item , MkString("network") , MkMap(&VarMap{})); _ = networkInfo;
  feeInfo := this.SafeValue(networkInfo , MkString("transaction_fee") ); _ = feeInfo;
  if IsTrue(OpNotEq2(feeInfo , MkUndefined() )) {
    feeCurrencyId := this.SafeString(feeInfo , MkString("currency") ); _ = feeCurrencyId;
    feeCurrencyCode := this.SafeCurrencyCode(feeCurrencyId , currency ); _ = feeCurrencyCode;
    feeAmount := this.SafeNumber(feeInfo , MkString("amount") ); _ = feeAmount;
    fee = MkMap(&VarMap{
        "cost":feeAmount ,
        "currency":feeCurrencyCode ,
        });
  }
  timestamp := this.Parse8601(this.SafeValue(item , MkString("created_at") )); _ = timestamp;
  id := this.SafeString(item , MkString("id") ); _ = id;
  type_ := this.ParseLedgerEntryType(this.SafeString(item , MkString("type") )); _ = type_;
  status := this.ParseLedgerEntryStatus(this.SafeString(item , MkString("status") )); _ = status;
  path := this.SafeString(item , MkString("resource_path") ); _ = path;
  accountId := OpCopy(MkUndefined() ); _ = accountId;
  if IsTrue(OpNotEq2(path , MkUndefined() )) {
    parts := path.Split(MkString("/") ); _ = parts;
    numParts := OpCopy(parts.Length ); _ = numParts;
    if IsTrue(OpGt(numParts , MkInteger(3) )) {
      accountId = *(parts ).At(MkInteger(3) );
    }
  }
  return MkMap(&VarMap{
        "info":item ,
        "id":id ,
        "timestamp":timestamp ,
        "datetime":this.Iso8601(timestamp ),
        "direction":direction ,
        "account":accountId ,
        "referenceId":MkUndefined() ,
        "referenceAccount":MkUndefined() ,
        "type":type_ ,
        "currency":code ,
        "amount":amount ,
        "before":MkUndefined() ,
        "after":MkUndefined() ,
        "status":status ,
        "fee":fee ,
        });
}

func (this *Coinbase) FindAccountId(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined()); _ = code;
  this.LoadMarkets();
  this.LoadAccounts();
  for i := MkInteger(0) ; IsTrue(OpLw(i , (*((this).At(MkString("accounts")))).Length )); OpIncr(& i ){
    account := *(*this.At(MkString("accounts")) ).At(i ); _ = account;
    if IsTrue(OpEq2(*(account ).At(MkString("code") ), code )) {
      return *(account ).At(MkString("id") );
    }
  }
  return MkUndefined() ;
}

func (this *Coinbase) PrepareAccountRequest(goArgs ...*Variant) *Variant{
  limit := GoGetArg(goArgs, 0, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 1, MkMap(&VarMap{})); _ = params;
  accountId := this.SafeString2(params , MkString("account_id") , MkString("accountId") ); _ = accountId;
  if IsTrue(OpEq2(accountId , MkUndefined() )) {
    panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" method requires an account_id (or accountId) parameter") )));
  }
  request := MkMap(&VarMap{"account_id":accountId }); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  return request ;
}

func (this *Coinbase) PrepareAccountRequestWithCurrencyCode(goArgs ...*Variant) *Variant{
  code := GoGetArg(goArgs, 0, MkUndefined() ); _ = code;
  limit := GoGetArg(goArgs, 1, MkUndefined() ); _ = limit;
  params := GoGetArg(goArgs, 2, MkMap(&VarMap{})); _ = params;
  accountId := this.SafeString2(params , MkString("account_id") , MkString("accountId") ); _ = accountId;
  if IsTrue(OpEq2(accountId , MkUndefined() )) {
    if IsTrue(OpEq2(code , MkUndefined() )) {
      panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")) , MkString(" method requires an account_id (or accountId) parameter OR a currency code argument") )));
    }
    accountId = this.FindAccountId(code );
    if IsTrue(OpEq2(accountId , MkUndefined() )) {
      panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" could not find account id for ") , code ))));
    }
  }
  request := MkMap(&VarMap{"account_id":accountId }); _ = request;
  if IsTrue(OpNotEq2(limit , MkUndefined() )) {
    *(request ).At(MkString("limit") )= OpCopy(limit );
  }
  return request ;
}

func (this *Coinbase) Sign(goArgs ...*Variant) *Variant{
  path := GoGetArg(goArgs, 0, MkUndefined()); _ = path;
  api := GoGetArg(goArgs, 1, MkString("public") ); _ = api;
  method := GoGetArg(goArgs, 2, MkString("GET") ); _ = method;
  params := GoGetArg(goArgs, 3, MkMap(&VarMap{})); _ = params;
  headers := GoGetArg(goArgs, 4, MkUndefined() ); _ = headers;
  body := GoGetArg(goArgs, 5, MkUndefined() ); _ = body;
  fullPath := OpAdd(MkString("/") , OpAdd(*this.At(MkString("version")) , OpAdd(MkString("/") , this.ImplodeParams(path , params )))); _ = fullPath;
  query := this.Omit(params , this.ExtractParams(path )); _ = query;
  if IsTrue(OpEq2(method , MkString("GET") )) {
    if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
      OpAddAssign(& fullPath , OpAdd(MkString("?") , this.Urlencode(query )));
    }
  }
  url := OpAdd(*(*this.At(MkString("urls")) ).At(MkString("api") ), fullPath ); _ = url;
  if IsTrue(OpEq2(api , MkString("private") )) {
    authorization := this.SafeString(*this.At(MkString("headers")) , MkString("Authorization") ); _ = authorization;
    if IsTrue(OpNotEq2(authorization , MkUndefined() )) {
      headers = MkMap(&VarMap{
          "Authorization":authorization ,
          "Content-Type":MkString("application/json") ,
          });
    } else {
      if IsTrue(*this.At(MkString("token")) ) {
        headers = MkMap(&VarMap{
            "Authorization":OpAdd(MkString("Bearer ") , *this.At(MkString("token")) ),
            "Content-Type":MkString("application/json") ,
            });
      } else {
        this.CheckRequiredCredentials();
        nonce := (this.Nonce()).Call(MkString("toString") ); _ = nonce;
        payload := MkString("") ; _ = payload;
        if IsTrue(OpNotEq2(method , MkString("GET") )) {
          if IsTrue(*(GoGetKeys(query )).At(MkString("length") )) {
            body = this.Json(query );
            payload = OpCopy(body );
          }
        }
        auth := OpAdd(nonce , OpAdd(method , OpAdd(fullPath , payload ))); _ = auth;
        signature := this.Hmac(this.Encode(auth ), this.Encode(*this.At(MkString("secret")) )); _ = signature;
        headers = MkMap(&VarMap{
            "CB-ACCESS-KEY":*this.At(MkString("apiKey")) ,
            "CB-ACCESS-SIGN":signature ,
            "CB-ACCESS-TIMESTAMP":nonce ,
            "Content-Type":MkString("application/json") ,
            });
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

func (this *Coinbase) HandleErrors(goArgs ...*Variant) *Variant{
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
  feedback := OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" ") , body )); _ = feedback;
  errorCode := this.SafeString(response , MkString("error") ); _ = errorCode;
  if IsTrue(OpNotEq2(errorCode , MkUndefined() )) {
    errorMessage := this.SafeString(response , MkString("error_description") ); _ = errorMessage;
    this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("exact") ), errorCode , feedback );
    this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("broad") ), errorMessage , feedback );
    panic(NewExchangeError(feedback ));
  }
  errors := this.SafeValue(response , MkString("errors") ); _ = errors;
  if IsTrue(OpNotEq2(errors , MkUndefined() )) {
    if IsTrue(GoIsArray(errors )) {
      numErrors := OpCopy(errors.Length ); _ = numErrors;
      if IsTrue(OpGt(numErrors , MkInteger(0) )) {
        errorCode = this.SafeString(*(errors ).At(MkInteger(0) ), MkString("id") );
        errorMessage := this.SafeString(*(errors ).At(MkInteger(0) ), MkString("message") ); _ = errorMessage;
        if IsTrue(OpNotEq2(errorCode , MkUndefined() )) {
          this.ThrowExactlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("exact") ), errorCode , feedback );
          this.ThrowBroadlyMatchedException(*(*this.At(MkString("exceptions")) ).At(MkString("broad") ), errorMessage , feedback );
          panic(NewExchangeError(feedback ));
        }
      }
    }
  }
  data := this.SafeValue(response , MkString("data") ); _ = data;
  if IsTrue(OpEq2(data , MkUndefined() )) {
    panic(NewExchangeError(OpAdd(*this.At(MkString("id")) , OpAdd(MkString(" failed due to a malformed response ") , this.Json(response )))));
  }
  return MkUndefined()
}

