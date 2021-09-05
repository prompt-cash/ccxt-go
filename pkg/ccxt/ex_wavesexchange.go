package ccxt

import ()

type Wavesexchange struct {
	*ExchangeBase
}

var _ Exchange = (*Wavesexchange)(nil)

func init() {
	exchange := &Wavesexchange{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Wavesexchange) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("wavesexchange"),
		"name": MkString("Waves.Exchange"),
		"countries": MkArray(&VarArray{
			MkString("CH"),
		}),
		"rateLimit": MkInteger(500),
		"certified": MkBool(true),
		"pro":       MkBool(false),
		"has": MkMap(&VarMap{
			"cancelOrder":         MkBool(true),
			"createMarketOrder":   MkBool(false),
			"createOrder":         MkBool(true),
			"fetchBalance":        MkBool(true),
			"fetchClosedOrders":   MkBool(true),
			"fetchDepositAddress": MkBool(true),
			"fetchMarkets":        MkBool(true),
			"fetchMyTrades":       MkBool(true),
			"fetchOHLCV":          MkBool(true),
			"fetchOpenOrders":     MkBool(true),
			"fetchOrderBook":      MkBool(true),
			"fetchOrders":         MkBool(true),
			"fetchTicker":         MkBool(true),
			"fetchTrades":         MkBool(true),
			"withdraw":            MkBool(true),
		}),
		"timeframes": MkMap(&VarMap{
			"1m":  MkString("1m"),
			"5m":  MkString("5m"),
			"15m": MkString("15m"),
			"30m": MkString("30m"),
			"1h":  MkString("1h"),
			"2h":  MkString("2h"),
			"3h":  MkString("3h"),
			"4h":  MkString("4h"),
			"6h":  MkString("6h"),
			"12h": MkString("12h"),
			"1d":  MkString("1d"),
			"1w":  MkString("1w"),
			"1M":  MkString("1M"),
		}),
		"urls": MkMap(&VarMap{
			"logo": MkString("https://user-images.githubusercontent.com/1294454/84547058-5fb27d80-ad0b-11ea-8711-78ac8b3c7f31.jpg"),
			"api": MkMap(&VarMap{
				"matcher": MkString("http://matcher.waves.exchange"),
				"node":    MkString("https://nodes.waves.exchange"),
				"public":  MkString("https://api.wavesplatform.com/v0"),
				"private": MkString("https://api.waves.exchange/v1"),
				"forward": MkString("https://waves.exchange/api/v1/forward/matcher"),
				"market":  MkString("https://waves.exchange/api/v1/forward/marketdata/api/v1"),
			}),
			"doc": MkString("https://docs.waves.exchange"),
			"www": MkString("https://waves.exchange"),
		}),
		"api": MkMap(&VarMap{
			"matcher": MkMap(&VarMap{
				"get": MkArray(&VarArray{
					MkString("matcher"),
					MkString("matcher/settings"),
					MkString("matcher/settings/rates"),
					MkString("matcher/balance/reserved/{publicKey}"),
					MkString("matcher/debug/allSnashotOffsets"),
					MkString("matcher/debug/currentOffset"),
					MkString("matcher/debug/lastOffset"),
					MkString("matcher/debug/oldestSnapshotOffset"),
					MkString("matcher/orderbook"),
					MkString("matcher/orderbook/{amountAsset}/{priceAsset}"),
					MkString("matcher/orderbook/{baseId}/{quoteId}/publicKey/{publicKey}"),
					MkString("matcher/orderbook/{baseId}/{quoteId}/{orderId}"),
					MkString("matcher/orderbook/{baseId}/{quoteId}/info"),
					MkString("matcher/orderbook/{baseId}/{quoteId}/status"),
					MkString("matcher/orderbook/{baseId}/{quoteId}/tradeableBalance/{address}"),
					MkString("matcher/orderbook/{publicKey}"),
					MkString("matcher/orderbook/{publicKey}/{orderId}"),
					MkString("matcher/orders/{address}"),
					MkString("matcher/orders/{address}/{orderId}"),
					MkString("matcher/transactions/{orderId}"),
				}),
				"post": MkArray(&VarArray{
					MkString("matcher/orderbook"),
					MkString("matcher/orderbook/market"),
					MkString("matcher/orderbook/cancel"),
					MkString("matcher/orderbook/{baseId}/{quoteId}/cancel"),
					MkString("matcher/debug/saveSnapshots"),
					MkString("matcher/orders/{address}/cancel"),
					MkString("matcher/orders/cancel/{orderId}"),
				}),
				"delete": MkArray(&VarArray{
					MkString("matcher/orderbook/{baseId}/{quoteId}"),
					MkString("matcher/settings/rates/{assetId}"),
				}),
				"put": MkArray(&VarArray{
					MkString("matcher/settings/rates/{assetId}"),
				}),
			}),
			"node": MkMap(&VarMap{
				"get": MkArray(&VarArray{
					MkString("addresses"),
					MkString("addresses/balance/{address}"),
					MkString("addresses/balance/{address}/{confirmations}"),
					MkString("addresses/balance/details/{address}"),
					MkString("addresses/data/{address}"),
					MkString("addresses/data/{address}/{key}"),
					MkString("addresses/effectiveBalance/{address}"),
					MkString("addresses/effectiveBalance/{address}/{confirmations}"),
					MkString("addresses/publicKey/{publicKey}"),
					MkString("addresses/scriptInfo/{address}"),
					MkString("addresses/scriptInfo/{address}/meta"),
					MkString("addresses/seed/{address}"),
					MkString("addresses/seq/{from}/{to}"),
					MkString("addresses/validate/{address}"),
					MkString("alias/by-address/{address}"),
					MkString("alias/by-alias/{alias}"),
					MkString("assets/{assetId}/distribution/{height}/{limit}"),
					MkString("assets/balance/{address}"),
					MkString("assets/balance/{address}/{assetId}"),
					MkString("assets/details/{assetId}"),
					MkString("assets/nft/{address}/limit/{limit}"),
					MkString("blockchain/rewards"),
					MkString("blockchain/rewards/height"),
					MkString("blocks/address/{address}/{from}/{to}/"),
					MkString("blocks/at/{height}"),
					MkString("blocks/delay/{signature}/{blockNum}"),
					MkString("blocks/first"),
					MkString("blocks/headers/last"),
					MkString("blocks/headers/seq/{from}/{to}"),
					MkString("blocks/height"),
					MkString("blocks/height/{signature}"),
					MkString("blocks/last"),
					MkString("blocks/seq/{from}/{to}"),
					MkString("blocks/signature/{signature}"),
					MkString("consensus/algo"),
					MkString("consensus/basetarget"),
					MkString("consensus/basetarget/{blockId}"),
					MkString("consensus/{generatingbalance}/address"),
					MkString("consensus/generationsignature"),
					MkString("consensus/generationsignature/{blockId}"),
					MkString("debug/balances/history/{address}"),
					MkString("debug/blocks/{howMany}"),
					MkString("debug/configInfo"),
					MkString("debug/historyInfo"),
					MkString("debug/info"),
					MkString("debug/minerInfo"),
					MkString("debug/portfolios/{address}"),
					MkString("debug/state"),
					MkString("debug/stateChanges/address/{address}"),
					MkString("debug/stateChanges/info/{id}"),
					MkString("debug/stateWaves/{height}"),
					MkString("leasing/active/{address}"),
					MkString("node/state"),
					MkString("node/version"),
					MkString("peers/all"),
					MkString("peers/blacklisted"),
					MkString("peers/connected"),
					MkString("peers/suspended"),
					MkString("transactions/address/{address}/limit/{limit}"),
					MkString("transactions/info/{id}"),
					MkString("transactions/status"),
					MkString("transactions/unconfirmed"),
					MkString("transactions/unconfirmed/info/{id}"),
					MkString("transactions/unconfirmed/size"),
					MkString("utils/seed"),
					MkString("utils/seed/{length}"),
					MkString("utils/time"),
					MkString("wallet/seed"),
				}),
				"post": MkArray(&VarArray{
					MkString("addresses"),
					MkString("addresses/data/{address}"),
					MkString("addresses/sign/{address}"),
					MkString("addresses/signText/{address}"),
					MkString("addresses/verify/{address}"),
					MkString("addresses/verifyText/{address}"),
					MkString("debug/blacklist"),
					MkString("debug/print"),
					MkString("debug/rollback"),
					MkString("debug/validate"),
					MkString("node/stop"),
					MkString("peers/clearblacklist"),
					MkString("peers/connect"),
					MkString("transactions/broadcast"),
					MkString("transactions/calculateFee"),
					MkString("tranasctions/sign"),
					MkString("transactions/sign/{signerAddress}"),
					MkString("tranasctions/status"),
					MkString("utils/hash/fast"),
					MkString("utils/hash/secure"),
					MkString("utils/script/compileCode"),
					MkString("utils/script/compileWithImports"),
					MkString("utils/script/decompile"),
					MkString("utils/script/estimate"),
					MkString("utils/sign/{privateKey}"),
					MkString("utils/transactionsSerialize"),
				}),
				"delete": MkArray(&VarArray{
					MkString("addresses/{address}"),
					MkString("debug/rollback-to/{signature}"),
				}),
			}),
			"public": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("pairs"),
				MkString("candles/{baseId}/{quoteId}"),
				MkString("transactions/exchange"),
			})}),
			"private": MkMap(&VarMap{
				"get": MkArray(&VarArray{
					MkString("deposit/addresses/{code}"),
					MkString("deposit/currencies"),
					MkString("withdraw/currencies"),
					MkString("withdraw/addresses/{currency}/{address}"),
				}),
				"post": MkArray(&VarArray{
					MkString("oauth2/token"),
				}),
			}),
			"forward": MkMap(&VarMap{
				"get": MkArray(&VarArray{
					MkString("matcher/orders/{address}"),
					MkString("matcher/orders/{address}/{orderId}"),
				}),
				"post": MkArray(&VarArray{
					MkString("matcher/orders/{wavesAddress}/cancel"),
				}),
			}),
			"market": MkMap(&VarMap{"get": MkArray(&VarArray{
				MkString("tickers"),
			})}),
		}),
		"options": MkMap(&VarMap{
			"allowedCandles":           MkInteger(1440),
			"accessToken":              MkUndefined(),
			"matcherPublicKey":         MkUndefined(),
			"quotes":                   MkUndefined(),
			"createOrderDefaultExpiry": MkInteger(2419200000),
			"wavesAddress":             MkUndefined(),
			"withdrawFeeUSDN":          MkInteger(7420),
			"withdrawFeeWAVES":         MkInteger(100000),
			"wavesPrecision":           MkInteger(8),
		}),
		"requiresEddsa": MkBool(true),
		"exceptions": MkMap(&VarMap{
			"3147270":   InsufficientFunds,
			"112":       InsufficientFunds,
			"4":         ExchangeError,
			"13":        ExchangeNotAvailable,
			"14":        ExchangeNotAvailable,
			"3145733":   AccountSuspended,
			"3148040":   DuplicateOrderId,
			"3148801":   AuthenticationError,
			"9440512":   AuthenticationError,
			"9440771":   BadSymbol,
			"9441026":   InvalidOrder,
			"9441282":   InvalidOrder,
			"9441286":   InvalidOrder,
			"9441295":   InvalidOrder,
			"9441540":   InvalidOrder,
			"9441542":   InvalidOrder,
			"106954752": AuthenticationError,
			"106954769": AuthenticationError,
			"106957828": AuthenticationError,
			"106960131": AuthenticationError,
			"106981137": AuthenticationError,
			"9437193":   OrderNotFound,
			"1048577":   BadRequest,
			"1051904":   AuthenticationError,
		}),
	}))
}

func (this *Wavesexchange) GetQuotes(goArgs ...*Variant) *Variant {
	quotes := this.SafeValue(*this.At(MkString("options")), MkString("quotes"))
	_ = quotes
	if IsTrue(quotes) {
		return quotes
	} else {
		response := this.Call(MkString("matcherGetMatcherSettings"))
		_ = response
		quotes = MkMap(&VarMap{})
		priceAssets := this.SafeValue(response, MkString("priceAssets"))
		_ = priceAssets
		for i := MkInteger(0); IsTrue(OpLw(i, priceAssets.Length)); OpIncr(&i) {
			*(quotes).At(*(priceAssets).At(i)) = OpCopy(MkBool(true))
		}
		*(*this.At(MkString("options"))).At(MkString("quotes")) = OpCopy(quotes)
		return quotes
	}
	return MkUndefined()
}

func (this *Wavesexchange) FetchMarkets(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	response := this.Call(MkString("marketGetTickers"))
	_ = response
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, response.Length)); OpIncr(&i) {
		entry := *(response).At(i)
		_ = entry
		baseId := this.SafeString(entry, MkString("amountAssetID"))
		_ = baseId
		quoteId := this.SafeString(entry, MkString("priceAssetID"))
		_ = quoteId
		id := OpAdd(baseId, OpAdd(MkString("/"), quoteId))
		_ = id
		marketId := this.SafeString(entry, MkString("symbol"))
		_ = marketId
		base := MkUndefined()
		_ = base
		quote := MkUndefined()
		_ = quote
		ArrayUnpack(marketId.Split(MkString("/")), &base, &quote)
		symbol := OpAdd(this.SafeCurrencyCode(base), OpAdd(MkString("/"), this.SafeCurrencyCode(quote)))
		_ = symbol
		precision := MkMap(&VarMap{
			"amount": this.SafeInteger(entry, MkString("amountAssetDecimals")),
			"price":  this.SafeInteger(entry, MkString("priceAssetDecimals")),
		})
		_ = precision
		result.Push(MkMap(&VarMap{
			"symbol":    symbol,
			"id":        id,
			"base":      base,
			"quote":     quote,
			"baseId":    baseId,
			"quoteId":   quoteId,
			"info":      entry,
			"precision": precision,
		}))
	}
	return result
}

func (this *Wavesexchange) FetchOrderBook(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	limit := GoGetArg(goArgs, 1, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := this.Extend(MkMap(&VarMap{
		"amountAsset": *(market).At(MkString("baseId")),
		"priceAsset":  *(market).At(MkString("quoteId")),
	}), params)
	_ = request
	response := this.Call(MkString("matcherGetMatcherOrderbookAmountAssetPriceAsset"), request)
	_ = response
	timestamp := this.SafeInteger(response, MkString("timestamp"))
	_ = timestamp
	bids := this.ParseOrderBookSide(this.SafeValue(response, MkString("bids")), market, limit)
	_ = bids
	asks := this.ParseOrderBookSide(this.SafeValue(response, MkString("asks")), market, limit)
	_ = asks
	return MkMap(&VarMap{
		"symbol":    symbol,
		"bids":      bids,
		"asks":      asks,
		"timestamp": timestamp,
		"datetime":  this.Iso8601(timestamp),
		"nonce":     MkUndefined(),
	})
}

func (this *Wavesexchange) ParseOrderBookSide(goArgs ...*Variant) *Variant {
	bookSide := GoGetArg(goArgs, 0, MkUndefined())
	_ = bookSide
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	precision := *(market).At(MkString("precision"))
	_ = precision
	wavesPrecision := this.SafeInteger(*this.At(MkString("options")), MkString("wavesPrecision"), MkInteger(8))
	_ = wavesPrecision
	amountPrecision := Math.Pow(MkInteger(10), *(precision).At(MkString("amount")))
	_ = amountPrecision
	difference := OpSub(*(precision).At(MkString("amount")), *(precision).At(MkString("price")))
	_ = difference
	pricePrecision := Math.Pow(MkInteger(10), OpSub(wavesPrecision, difference))
	_ = pricePrecision
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, bookSide.Length)); OpIncr(&i) {
		entry := *(bookSide).At(i)
		_ = entry
		price := OpDiv(this.SafeInteger(entry, MkString("price"), MkInteger(0)), pricePrecision)
		_ = price
		amount := OpDiv(this.SafeInteger(entry, MkString("amount"), MkInteger(0)), amountPrecision)
		_ = amount
		if IsTrue(OpAnd(OpNotEq2(limit, MkUndefined()), OpGt(i, limit))) {
			break
		}
		result.Push(MkArray(&VarArray{
			price,
			amount,
		}))
	}
	return result
}

func (this *Wavesexchange) CheckRequiredKeys(goArgs ...*Variant) *Variant {
	if IsTrue(OpEq2(*this.At(MkString("apiKey")), MkUndefined())) {
		panic(NewAuthenticationError(OpAdd(*this.At(MkString("id")), MkString(" requires apiKey credential"))))
	}
	if IsTrue(OpEq2(*this.At(MkString("secret")), MkUndefined())) {
		panic(NewAuthenticationError(OpAdd(*this.At(MkString("id")), MkString(" requires secret credential"))))
	}
	apiKeyBytes := OpCopy(MkUndefined())
	_ = apiKeyBytes
	secretKeyBytes := OpCopy(MkUndefined())
	_ = secretKeyBytes
	{
		ret__ := func(this *Wavesexchange) (ret_ *Variant) {
			defer func() {
				if e := recover().(*Variant); e != nil {
					ret_ = func(this *Wavesexchange) *Variant {
						// catch block:
						panic(NewAuthenticationError(OpAdd(*this.At(MkString("id")), MkString(" apiKey must be a base58 encoded public key"))))
						return nil
					}(this)
				}
			}()
			// try block:
			apiKeyBytes = this.Base58ToBinary(*this.At(MkString("apiKey")))
			return nil
		}(this)
		if ret__ != nil {
			return ret__
		}
	}
	{
		ret__ := func(this *Wavesexchange) (ret_ *Variant) {
			defer func() {
				if e := recover().(*Variant); e != nil {
					ret_ = func(this *Wavesexchange) *Variant {
						// catch block:
						panic(NewAuthenticationError(OpAdd(*this.At(MkString("id")), MkString(" secret must be a base58 encoded private key"))))
						return nil
					}(this)
				}
			}()
			// try block:
			secretKeyBytes = this.Base58ToBinary(*this.At(MkString("secret")))
			return nil
		}(this)
		if ret__ != nil {
			return ret__
		}
	}
	hexApiKeyBytes := this.BinaryToBase16(apiKeyBytes)
	_ = hexApiKeyBytes
	hexSecretKeyBytes := this.BinaryToBase16(secretKeyBytes)
	_ = hexSecretKeyBytes
	if IsTrue(OpNotEq2(hexApiKeyBytes.Length, MkInteger(64))) {
		panic(NewAuthenticationError(OpAdd(*this.At(MkString("id")), MkString(" apiKey must be a base58 encoded public key"))))
	}
	if IsTrue(OpNotEq2(hexSecretKeyBytes.Length, MkInteger(64))) {
		panic(NewAuthenticationError(OpAdd(*this.At(MkString("id")), MkString(" secret must be a base58 encoded private key"))))
	}
	return MkUndefined()
}

func (this *Wavesexchange) Sign(goArgs ...*Variant) *Variant {
	path := GoGetArg(goArgs, 0, MkUndefined())
	_ = path
	api := GoGetArg(goArgs, 1, MkString("public"))
	_ = api
	method := GoGetArg(goArgs, 2, MkString("GET"))
	_ = method
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	headers := GoGetArg(goArgs, 4, MkUndefined())
	_ = headers
	body := GoGetArg(goArgs, 5, MkUndefined())
	_ = body
	query := this.Omit(params, this.ExtractParams(path))
	_ = query
	isCancelOrder := OpEq2(path, MkString("matcher/orders/{wavesAddress}/cancel"))
	_ = isCancelOrder
	path = this.ImplodeParams(path, params)
	url := OpAdd(*(*(*this.At(MkString("urls"))).At(MkString("api"))).At(api), OpAdd(MkString("/"), path))
	_ = url
	queryString := this.Urlencode(query)
	_ = queryString
	if IsTrue(OpOr(OpEq2(api, MkString("private")), OpEq2(api, MkString("forward")))) {
		headers = MkMap(&VarMap{"Accept": MkString("application/json")})
		accessToken := this.SafeString(*this.At(MkString("options")), MkString("accessToken"))
		_ = accessToken
		if IsTrue(accessToken) {
			*(headers).At(MkString("Authorization")) = OpAdd(MkString("Bearer "), accessToken)
		}
		if IsTrue(OpEq2(method, MkString("POST"))) {
			*(headers).At(MkString("content-type")) = MkString("application/json")
		} else {
			*(headers).At(MkString("content-type")) = MkString("application/x-www-form-urlencoded")
		}
		if IsTrue(isCancelOrder) {
			body = this.Json(MkArray(&VarArray{
				*(query).At(MkString("orderId")),
			}))
			queryString = MkString("")
		}
		if IsTrue(OpGt(queryString.Length, MkInteger(0))) {
			OpAddAssign(&url, OpAdd(MkString("?"), queryString))
		}
	} else {
		if IsTrue(OpEq2(api, MkString("matcher"))) {
			if IsTrue(OpEq2(method, MkString("POST"))) {
				headers = MkMap(&VarMap{"content-type": MkString("application/json")})
				body = this.Json(query)
			} else {
				headers = OpCopy(query)
			}
		} else {
			if IsTrue(OpEq2(method, MkString("POST"))) {
				headers = MkMap(&VarMap{"content-type": MkString("application/json")})
				body = this.Json(query)
			} else {
				headers = MkMap(&VarMap{"content-type": MkString("application/x-www-form-urlencoded")})
				if IsTrue(OpGt(queryString.Length, MkInteger(0))) {
					OpAddAssign(&url, OpAdd(MkString("?"), queryString))
				}
			}
		}
	}
	return MkMap(&VarMap{
		"url":     url,
		"method":  method,
		"body":    body,
		"headers": headers,
	})
}

func (this *Wavesexchange) GetAccessToken(goArgs ...*Variant) *Variant {
	if IsTrue(OpNot(this.SafeString(*this.At(MkString("options")), MkString("accessToken")))) {
		prefix := MkString("ffffff01")
		_ = prefix
		expiresDelta := OpMulti(MkInteger(60), OpMulti(MkInteger(60), OpMulti(MkInteger(24), MkInteger(7))))
		_ = expiresDelta
		seconds := this.Sum(this.Seconds(), expiresDelta)
		_ = seconds
		seconds = seconds.ToString()
		clientId := MkString("waves.exchange")
		_ = clientId
		message := OpAdd(MkString("W:"), OpAdd(clientId, OpAdd(MkString(":"), seconds)))
		_ = message
		messageHex := this.BinaryToBase16(this.StringToBinary(this.Encode(message)))
		_ = messageHex
		payload := OpAdd(prefix, messageHex)
		_ = payload
		hexKey := this.BinaryToBase16(this.Base58ToBinary(*this.At(MkString("secret"))))
		_ = hexKey
		signature := this.Eddsa(payload, hexKey, MkString("ed25519"))
		_ = signature
		request := MkMap(&VarMap{
			"grant_type": MkString("password"),
			"scope":      MkString("general"),
			"username":   *this.At(MkString("apiKey")),
			"password":   OpAdd(seconds, OpAdd(MkString(":"), signature)),
			"client_id":  clientId,
		})
		_ = request
		response := this.Call(MkString("privatePostOauth2Token"), request)
		_ = response
		*(*this.At(MkString("options"))).At(MkString("accessToken")) = this.SafeString(response, MkString("access_token"))
		return *(*this.At(MkString("options"))).At(MkString("accessToken"))
	}
	return MkUndefined()
}

func (this *Wavesexchange) ParseTicker(goArgs ...*Variant) *Variant {
	ticker := GoGetArg(goArgs, 0, MkUndefined())
	_ = ticker
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := OpCopy(MkUndefined())
	_ = timestamp
	baseId := this.SafeString(ticker, MkString("amountAsset"))
	_ = baseId
	quoteId := this.SafeString(ticker, MkString("priceAsset"))
	_ = quoteId
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(OpAnd(OpNotEq2(baseId, MkUndefined()), OpNotEq2(quoteId, MkUndefined()))) {
		marketId := OpAdd(baseId, OpAdd(MkString("/"), quoteId))
		_ = marketId
		if IsTrue(OpHasMember(marketId, *this.At(MkString("markets_by_id")))) {
			market = *(*this.At(MkString("markets_by_id"))).At(marketId)
		} else {
			base := this.SafeCurrencyCode(baseId)
			_ = base
			quote := this.SafeCurrencyCode(quoteId)
			_ = quote
			symbol = OpAdd(base, OpAdd(MkString("/"), quote))
		}
	}
	if IsTrue(OpAnd(OpEq2(symbol, MkUndefined()), OpNotEq2(market, MkUndefined()))) {
		symbol = *(market).At(MkString("symbol"))
	}
	data := this.SafeValue(ticker, MkString("data"), MkMap(&VarMap{}))
	_ = data
	last := this.SafeNumber(data, MkString("lastPrice"))
	_ = last
	low := this.SafeNumber(data, MkString("low"))
	_ = low
	high := this.SafeNumber(data, MkString("high"))
	_ = high
	vwap := this.SafeNumber(data, MkString("weightedAveragePrice"))
	_ = vwap
	baseVolume := this.SafeNumber(data, MkString("volume"))
	_ = baseVolume
	quoteVolume := this.SafeNumber(data, MkString("quoteVolume"))
	_ = quoteVolume
	open := this.SafeNumber(data, MkString("firstPrice"))
	_ = open
	change := OpCopy(MkUndefined())
	_ = change
	average := OpCopy(MkUndefined())
	_ = average
	percentage := OpCopy(MkUndefined())
	_ = percentage
	if IsTrue(OpAnd(OpNotEq2(last, MkUndefined()), OpNotEq2(open, MkUndefined()))) {
		change = OpSub(last, open)
		average = OpDiv(this.Sum(last, open), MkInteger(2))
		if IsTrue(OpGt(open, MkInteger(0))) {
			percentage = OpDiv(change, OpMulti(open, MkInteger(100)))
		}
	}
	return MkMap(&VarMap{
		"symbol":        symbol,
		"timestamp":     timestamp,
		"datetime":      this.Iso8601(timestamp),
		"high":          high,
		"low":           low,
		"bid":           MkUndefined(),
		"bidVolume":     MkUndefined(),
		"ask":           MkUndefined(),
		"askVolume":     MkUndefined(),
		"vwap":          vwap,
		"open":          open,
		"close":         last,
		"last":          last,
		"previousClose": MkUndefined(),
		"change":        change,
		"percentage":    percentage,
		"average":       average,
		"baseVolume":    baseVolume,
		"quoteVolume":   quoteVolume,
		"info":          ticker,
	})
}

func (this *Wavesexchange) FetchTicker(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{"pairs": *(market).At(MkString("id"))})
	_ = request
	response := this.Call(MkString("publicGetPairs"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	ticker := this.SafeValue(data, MkInteger(0), MkMap(&VarMap{}))
	_ = ticker
	return this.ParseTicker(ticker, market)
}

func (this *Wavesexchange) FetchOHLCV(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	timeframe := GoGetArg(goArgs, 1, MkString("1m"))
	_ = timeframe
	since := GoGetArg(goArgs, 2, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 3, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 4, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{
		"baseId":   *(market).At(MkString("baseId")),
		"quoteId":  *(market).At(MkString("quoteId")),
		"interval": *(*this.At(MkString("timeframes"))).At(timeframe),
	})
	_ = request
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("timeStart")) = since.ToString()
	} else {
		allowedCandles := this.SafeInteger(*this.At(MkString("options")), MkString("allowedCandles"), MkInteger(1440))
		_ = allowedCandles
		timeframeUnix := OpMulti(this.ParseTimeframe(timeframe), MkInteger(1000))
		_ = timeframeUnix
		currentTime := OpMulti(Math.Floor(OpDiv(this.Milliseconds(), timeframeUnix)), timeframeUnix)
		_ = currentTime
		delta := OpMulti(OpSub(allowedCandles, MkInteger(1)), timeframeUnix)
		_ = delta
		timeStart := OpSub(currentTime, delta)
		_ = timeStart
		*(request).At(MkString("timeStart")) = timeStart.ToString()
	}
	response := this.Call(MkString("publicGetCandlesBaseIdQuoteId"), this.Extend(request, params))
	_ = response
	data := this.SafeValue(response, MkString("data"), MkArray(&VarArray{}))
	_ = data
	result := this.ParseOHLCVs(data, market, timeframe, since, limit)
	_ = result
	lastClose := OpCopy(MkUndefined())
	_ = lastClose
	length := OpCopy(result.Length)
	_ = length
	for i := MkInteger(0); IsTrue(OpLw(i, result.Length)); OpIncr(&i) {
		j := OpSub(length, OpSub(i, MkInteger(1)))
		_ = j
		entry := *(result).At(j)
		_ = entry
		open := *(entry).At(MkInteger(1))
		_ = open
		if IsTrue(OpEq2(open, MkUndefined())) {
			*(entry).At(MkInteger(1)) = OpCopy(lastClose)
			*(entry).At(MkInteger(2)) = OpCopy(lastClose)
			*(entry).At(MkInteger(3)) = OpCopy(lastClose)
			*(entry).At(MkInteger(4)) = OpCopy(lastClose)
			*(result).At(j) = OpCopy(entry)
		}
		lastClose = *(entry).At(MkInteger(4))
	}
	return result
}

func (this *Wavesexchange) ParseOHLCV(goArgs ...*Variant) *Variant {
	ohlcv := GoGetArg(goArgs, 0, MkUndefined())
	_ = ohlcv
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	data := this.SafeValue(ohlcv, MkString("data"), MkMap(&VarMap{}))
	_ = data
	return MkArray(&VarArray{
		this.Parse8601(this.SafeString(data, MkString("time"))),
		this.SafeNumber(data, MkString("open")),
		this.SafeNumber(data, MkString("high")),
		this.SafeNumber(data, MkString("low")),
		this.SafeNumber(data, MkString("close")),
		this.SafeNumber(data, MkString("volume"), MkInteger(0)),
	})
}

func (this *Wavesexchange) FetchDepositAddress(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	_ = params
	this.GetAccessToken()
	supportedCurrencies := this.Call(MkString("privateGetDepositCurrencies"))
	_ = supportedCurrencies
	currencies := MkMap(&VarMap{})
	_ = currencies
	items := this.SafeValue(supportedCurrencies, MkString("items"), MkArray(&VarArray{}))
	_ = items
	for i := MkInteger(0); IsTrue(OpLw(i, items.Length)); OpIncr(&i) {
		entry := *(items).At(i)
		_ = entry
		currencyCode := this.SafeString(entry, MkString("id"))
		_ = currencyCode
		*(currencies).At(currencyCode) = OpCopy(MkBool(true))
	}
	if IsTrue(OpNot(OpHasMember(code, currencies))) {
		codes := GoGetKeys(currencies)
		_ = codes
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" fetch "), OpAdd(code, OpAdd(MkString(" deposit address not supported. Currency code must be one of "), codes.ToString()))))))
	}
	request := this.Extend(MkMap(&VarMap{"code": code}), params)
	_ = request
	response := this.Call(MkString("privateGetDepositAddressesCode"), request)
	_ = response
	addresses := this.SafeValue(response, MkString("deposit_addresses"))
	_ = addresses
	address := this.SafeString(addresses, MkInteger(0))
	_ = address
	return MkMap(&VarMap{
		"address": address,
		"code":    code,
		"tag":     MkUndefined(),
		"info":    response,
	})
}

func (this *Wavesexchange) GetMatcherPublicKey(goArgs ...*Variant) *Variant {
	matcherPublicKey := this.SafeString(*this.At(MkString("options")), MkString("matcherPublicKey"))
	_ = matcherPublicKey
	if IsTrue(matcherPublicKey) {
		return matcherPublicKey
	} else {
		response := this.Call(MkString("matcherGetMatcher"))
		_ = response
		*(*this.At(MkString("options"))).At(MkString("matcherPublicKey")) = response.Slice(MkInteger(1), OpSub(response.Length, MkInteger(1)))
		return *(*this.At(MkString("options"))).At(MkString("matcherPublicKey"))
	}
	return MkUndefined()
}

func (this *Wavesexchange) GetAssetBytes(goArgs ...*Variant) *Variant {
	currencyId := GoGetArg(goArgs, 0, MkUndefined())
	_ = currencyId
	if IsTrue(OpEq2(currencyId, MkString("WAVES"))) {
		return this.NumberToBE(MkInteger(0), MkInteger(1))
	} else {
		return this.BinaryConcat(this.NumberToBE(MkInteger(1), MkInteger(1)), this.Base58ToBinary(currencyId))
	}
	return MkUndefined()
}

func (this *Wavesexchange) GetAssetId(goArgs ...*Variant) *Variant {
	currencyId := GoGetArg(goArgs, 0, MkUndefined())
	_ = currencyId
	if IsTrue(OpEq2(currencyId, MkString("WAVES"))) {
		return MkString("")
	}
	return currencyId
}

func (this *Wavesexchange) PriceToPrecision(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	price := GoGetArg(goArgs, 1, MkUndefined())
	_ = price
	market := *(*this.At(MkString("markets"))).At(symbol)
	_ = market
	wavesPrecision := this.SafeInteger(*this.At(MkString("options")), MkString("wavesPrecision"), MkInteger(8))
	_ = wavesPrecision
	difference := OpSub(*(*(market).At(MkString("precision"))).At(MkString("amount")), *(*(market).At(MkString("precision"))).At(MkString("price")))
	_ = difference
	return parseInt(parseFloat(this.ToPrecision(price, OpSub(wavesPrecision, difference))))
}

func (this *Wavesexchange) AmountToPrecision(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	amount := GoGetArg(goArgs, 1, MkUndefined())
	_ = amount
	return parseInt(parseFloat(this.ToPrecision(amount, *(*(*(*this.At(MkString("markets"))).At(symbol)).At(MkString("precision"))).At(MkString("amount")))))
}

func (this *Wavesexchange) CurrencyToPrecision(goArgs ...*Variant) *Variant {
	currency := GoGetArg(goArgs, 0, MkUndefined())
	_ = currency
	amount := GoGetArg(goArgs, 1, MkUndefined())
	_ = amount
	return parseInt(parseFloat(this.ToPrecision(amount, *(*(*this.At(MkString("currencies"))).At(currency)).At(MkString("precision")))))
}

func (this *Wavesexchange) FromPrecision(goArgs ...*Variant) *Variant {
	amount := GoGetArg(goArgs, 0, MkUndefined())
	_ = amount
	scale := GoGetArg(goArgs, 1, MkUndefined())
	_ = scale
	if IsTrue(OpEq2(amount, MkUndefined())) {
		return MkUndefined()
	}
	precise := NewPrecise(amount)
	_ = precise
	precise.Decimals = OpAdd(precise.Decimals, scale)
	precise.Reduce()
	return precise.ToString()
}

func (this *Wavesexchange) ToPrecision(goArgs ...*Variant) *Variant {
	amount := GoGetArg(goArgs, 0, MkUndefined())
	_ = amount
	scale := GoGetArg(goArgs, 1, MkUndefined())
	_ = scale
	amountString := amount.ToString()
	_ = amountString
	precise := NewPrecise(amountString)
	_ = precise
	precise.Decimals = OpSub(precise.Decimals, scale)
	precise.Reduce()
	return precise.ToString()
}

func (this *Wavesexchange) CurrencyFromPrecision(goArgs ...*Variant) *Variant {
	currency := GoGetArg(goArgs, 0, MkUndefined())
	_ = currency
	amount := GoGetArg(goArgs, 1, MkUndefined())
	_ = amount
	scale := *(*(*this.At(MkString("currencies"))).At(currency)).At(MkString("precision"))
	_ = scale
	return this.FromPrecision(amount, scale)
}

func (this *Wavesexchange) PriceFromPrecision(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	price := GoGetArg(goArgs, 1, MkUndefined())
	_ = price
	market := *(*this.At(MkString("markets"))).At(symbol)
	_ = market
	wavesPrecision := this.SafeInteger(*this.At(MkString("options")), MkString("wavesPrecision"), MkInteger(8))
	_ = wavesPrecision
	scale := OpSub(wavesPrecision, OpAdd(*(*(market).At(MkString("precision"))).At(MkString("amount")), *(*(market).At(MkString("precision"))).At(MkString("price"))))
	_ = scale
	return this.FromPrecision(price, scale)
}

func (this *Wavesexchange) CreateOrder(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	type_ := GoGetArg(goArgs, 1, MkUndefined())
	_ = type_
	side := GoGetArg(goArgs, 2, MkUndefined())
	_ = side
	amount := GoGetArg(goArgs, 3, MkUndefined())
	_ = amount
	price := GoGetArg(goArgs, 4, MkUndefined())
	_ = price
	params := GoGetArg(goArgs, 5, MkMap(&VarMap{}))
	_ = params
	this.CheckRequiredDependencies()
	this.CheckRequiredKeys()
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	matcherPublicKey := this.GetMatcherPublicKey()
	_ = matcherPublicKey
	amountAsset := this.GetAssetId(*(market).At(MkString("baseId")))
	_ = amountAsset
	priceAsset := this.GetAssetId(*(market).At(MkString("quoteId")))
	_ = priceAsset
	amount = this.AmountToPrecision(symbol, amount)
	price = this.PriceToPrecision(symbol, price)
	orderType := OpTrinary(OpEq2(side, MkString("buy")), MkInteger(0), MkInteger(1))
	_ = orderType
	timestamp := this.Milliseconds()
	_ = timestamp
	defaultExpiryDelta := this.SafeInteger(*this.At(MkString("options")), MkString("createOrderDefaultExpiry"), MkInteger(2419200000))
	_ = defaultExpiryDelta
	expiration := this.Sum(timestamp, defaultExpiryDelta)
	_ = expiration
	settings := this.Call(MkString("matcherGetMatcherSettings"))
	_ = settings
	orderFee := this.SafeValue(settings, MkString("orderFee"))
	_ = orderFee
	dynamic := this.SafeValue(orderFee, MkString("dynamic"))
	_ = dynamic
	baseMatcherFee := this.SafeString(dynamic, MkString("baseFee"))
	_ = baseMatcherFee
	wavesMatcherFee := this.CurrencyFromPrecision(MkString("WAVES"), baseMatcherFee)
	_ = wavesMatcherFee
	rates := this.SafeValue(dynamic, MkString("rates"))
	_ = rates
	priceAssets := GoGetKeys(rates)
	_ = priceAssets
	matcherFeeAssetId := OpCopy(MkUndefined())
	_ = matcherFeeAssetId
	matcherFee := OpCopy(MkUndefined())
	_ = matcherFee
	if IsTrue(OpHasMember(MkString("feeAssetId"), params)) {
		matcherFeeAssetId = *(params).At(MkString("feeAssetId"))
	} else {
		if IsTrue(OpHasMember(MkString("feeAssetId"), *this.At(MkString("options")))) {
			matcherFeeAssetId = *(*this.At(MkString("options"))).At(MkString("feeAssetId"))
		} else {
			balances := this.FetchBalance()
			_ = balances
			if IsTrue(OpGt(*(*(balances).At(MkString("WAVES"))).At(MkString("free")), wavesMatcherFee)) {
				matcherFeeAssetId = MkString("WAVES")
				matcherFee = OpCopy(baseMatcherFee)
			} else {
				for i := MkInteger(0); IsTrue(OpLw(i, priceAssets.Length)); OpIncr(&i) {
					assetId := *(priceAssets).At(i)
					_ = assetId
					code := this.SafeCurrencyCode(assetId)
					_ = code
					balance := this.SafeValue(this.SafeValue(balances, code, MkMap(&VarMap{})), MkString("free"))
					_ = balance
					assetFee := Precise.StringMul(*(rates).At(assetId), wavesMatcherFee)
					_ = assetFee
					if IsTrue(OpAnd(OpNotEq2(balance, MkUndefined()), OpGt(balance, assetFee))) {
						matcherFeeAssetId = OpCopy(assetId)
						break
					}
				}
			}
		}
	}
	if IsTrue(OpEq2(matcherFeeAssetId, MkUndefined())) {
		panic(NewInsufficientFunds(OpAdd(*this.At(MkString("id")), MkString(" not enough funds to cover the fee, specify feeAssetId in params or options, or buy some WAVES"))))
	}
	if IsTrue(OpEq2(matcherFee, MkUndefined())) {
		wavesPrecision := this.SafeInteger(*this.At(MkString("options")), MkString("wavesPrecision"), MkInteger(8))
		_ = wavesPrecision
		rate := this.SafeString(rates, matcherFeeAssetId)
		_ = rate
		code := this.SafeCurrencyCode(matcherFeeAssetId)
		_ = code
		currency := this.Currency(code)
		_ = currency
		newPrecison := OpSub(wavesPrecision, *(currency).At(MkString("precision")))
		_ = newPrecison
		matcherFee = this.FromPrecision(Precise.StringMul(rate, baseMatcherFee), newPrecison)
		matcherFee = Precise.StringDiv(Precise.StringAdd(matcherFee, MkString("1")), MkString("1"), MkInteger(0))
	}
	byteArray := MkArray(&VarArray{
		this.NumberToBE(MkInteger(3), MkInteger(1)),
		this.Base58ToBinary(*this.At(MkString("apiKey"))),
		this.Base58ToBinary(matcherPublicKey),
		this.GetAssetBytes(*(market).At(MkString("baseId"))),
		this.GetAssetBytes(*(market).At(MkString("quoteId"))),
		this.NumberToBE(orderType, MkInteger(1)),
		this.NumberToBE(price, MkInteger(8)),
		this.NumberToBE(amount, MkInteger(8)),
		this.NumberToBE(timestamp, MkInteger(8)),
		this.NumberToBE(expiration, MkInteger(8)),
		this.NumberToBE(matcherFee, MkInteger(8)),
		this.GetAssetBytes(matcherFeeAssetId),
	})
	_ = byteArray
	binary := this.BinaryConcatArray(byteArray)
	_ = binary
	signature := this.Eddsa(this.BinaryToBase16(binary), this.BinaryToBase16(this.Base58ToBinary(*this.At(MkString("secret")))), MkString("ed25519"))
	_ = signature
	assetPair := MkMap(&VarMap{
		"amountAsset": amountAsset,
		"priceAsset":  priceAsset,
	})
	_ = assetPair
	body := MkMap(&VarMap{
		"senderPublicKey":  *this.At(MkString("apiKey")),
		"matcherPublicKey": matcherPublicKey,
		"assetPair":        assetPair,
		"orderType":        side,
		"price":            price,
		"amount":           amount,
		"timestamp":        timestamp,
		"expiration":       expiration,
		"matcherFee":       parseInt(matcherFee),
		"signature":        signature,
		"version":          MkInteger(3),
	})
	_ = body
	if IsTrue(OpNotEq2(matcherFeeAssetId, MkString("WAVES"))) {
		*(body).At(MkString("matcherFeeAssetId")) = OpCopy(matcherFeeAssetId)
	}
	response := this.Call(MkString("matcherPostMatcherOrderbook"), body)
	_ = response
	value := this.SafeValue(response, MkString("message"))
	_ = value
	return this.ParseOrder(value, market)
}

func (this *Wavesexchange) CancelOrder(goArgs ...*Variant) *Variant {
	id := GoGetArg(goArgs, 0, MkUndefined())
	_ = id
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	_ = symbol
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	this.CheckRequiredDependencies()
	this.CheckRequiredKeys()
	this.GetAccessToken()
	wavesAddress := this.GetWavesAddress()
	_ = wavesAddress
	response := this.Call(MkString("forwardPostMatcherOrdersWavesAddressCancel"), MkMap(&VarMap{
		"wavesAddress": wavesAddress,
		"orderId":      id,
	}))
	_ = response
	message := this.SafeValue(response, MkString("message"))
	_ = message
	firstMessage := this.SafeValue(message, MkInteger(0))
	_ = firstMessage
	firstOrder := this.SafeValue(firstMessage, MkInteger(0))
	_ = firstOrder
	returnedId := this.SafeString(firstOrder, MkString("orderId"))
	_ = returnedId
	return MkMap(&VarMap{
		"info":               response,
		"id":                 returnedId,
		"clientOrderId":      MkUndefined(),
		"timestamp":          MkUndefined(),
		"datetime":           MkUndefined(),
		"lastTradeTimestamp": MkUndefined(),
		"symbol":             symbol,
		"type":               MkUndefined(),
		"side":               MkUndefined(),
		"price":              MkUndefined(),
		"amount":             MkUndefined(),
		"cost":               MkUndefined(),
		"average":            MkUndefined(),
		"filled":             MkUndefined(),
		"remaining":          MkUndefined(),
		"status":             MkUndefined(),
		"fee":                MkUndefined(),
		"trades":             MkUndefined(),
	})
}

func (this *Wavesexchange) FetchOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.CheckRequiredDependencies()
	this.CheckRequiredKeys()
	if IsTrue(OpEq2(symbol, MkUndefined())) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" fetchOrders() requires symbol argument"))))
	}
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	timestamp := this.Milliseconds()
	_ = timestamp
	byteArray := MkArray(&VarArray{
		this.Base58ToBinary(*this.At(MkString("apiKey"))),
		this.NumberToBE(timestamp, MkInteger(8)),
	})
	_ = byteArray
	binary := this.BinaryConcatArray(byteArray)
	_ = binary
	hexSecret := this.BinaryToBase16(this.Base58ToBinary(*this.At(MkString("secret"))))
	_ = hexSecret
	signature := this.Eddsa(this.BinaryToBase16(binary), hexSecret, MkString("ed25519"))
	_ = signature
	request := MkMap(&VarMap{
		"Accept":    MkString("application/json"),
		"Timestamp": timestamp.ToString(),
		"Signature": signature,
		"publicKey": *this.At(MkString("apiKey")),
		"baseId":    *(market).At(MkString("baseId")),
		"quoteId":   *(market).At(MkString("quoteId")),
	})
	_ = request
	response := this.Call(MkString("matcherGetMatcherOrderbookBaseIdQuoteIdPublicKeyPublicKey"), this.Extend(request, params))
	_ = response
	return this.ParseOrders(response, market, since, limit)
}

func (this *Wavesexchange) FetchOpenOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	this.GetAccessToken()
	market := OpCopy(MkUndefined())
	_ = market
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
	}
	address := this.GetWavesAddress()
	_ = address
	request := MkMap(&VarMap{
		"address":    address,
		"activeOnly": MkBool(true),
	})
	_ = request
	response := this.Call(MkString("forwardGetMatcherOrdersAddress"), request)
	_ = response
	return this.ParseOrders(response, market, since, limit)
}

func (this *Wavesexchange) FetchClosedOrders(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	this.GetAccessToken()
	market := OpCopy(MkUndefined())
	_ = market
	if IsTrue(OpNotEq2(symbol, MkUndefined())) {
		market = this.Market(symbol)
	}
	address := this.GetWavesAddress()
	_ = address
	request := MkMap(&VarMap{
		"address":    address,
		"closedOnly": MkBool(true),
	})
	_ = request
	response := this.Call(MkString("forwardGetMatcherOrdersAddress"), request)
	_ = response
	return this.ParseOrders(response, market, since, limit)
}

func (this *Wavesexchange) ParseOrderStatus(goArgs ...*Variant) *Variant {
	status := GoGetArg(goArgs, 0, MkUndefined())
	_ = status
	statuses := MkMap(&VarMap{
		"Cancelled":       MkString("canceled"),
		"Accepted":        MkString("open"),
		"Filled":          MkString("closed"),
		"PartiallyFilled": MkString("open"),
	})
	_ = statuses
	return this.SafeString(statuses, status, status)
}

func (this *Wavesexchange) GetSymbolFromAssetPair(goArgs ...*Variant) *Variant {
	assetPair := GoGetArg(goArgs, 0, MkUndefined())
	_ = assetPair
	baseId := this.SafeString(assetPair, MkString("amountAsset"), MkString("WAVES"))
	_ = baseId
	quoteId := this.SafeString(assetPair, MkString("priceAsset"), MkString("WAVES"))
	_ = quoteId
	return OpAdd(this.SafeCurrencyCode(baseId), OpAdd(MkString("/"), this.SafeCurrencyCode(quoteId)))
}

func (this *Wavesexchange) ParseOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	_ = order
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	timestamp := this.SafeInteger(order, MkString("timestamp"))
	_ = timestamp
	side := this.SafeString2(order, MkString("type"), MkString("orderType"))
	_ = side
	type_ := MkString("limit")
	_ = type_
	if IsTrue(OpHasMember(MkString("type"), order)) {
		type_ = this.SafeString(order, MkString("orderType"), type_)
	}
	id := this.SafeString(order, MkString("id"))
	_ = id
	filledString := this.SafeString(order, MkString("filled"))
	_ = filledString
	priceString := this.SafeString(order, MkString("price"))
	_ = priceString
	amountString := this.SafeString(order, MkString("amount"))
	_ = amountString
	assetPair := this.SafeValue(order, MkString("assetPair"))
	_ = assetPair
	symbol := OpCopy(MkUndefined())
	_ = symbol
	if IsTrue(OpNotEq2(assetPair, MkUndefined())) {
		symbol = this.GetSymbolFromAssetPair(assetPair)
	} else {
		if IsTrue(OpNotEq2(market, MkUndefined())) {
			symbol = *(market).At(MkString("symbol"))
		}
	}
	amountCurrency := this.SafeCurrencyCode(this.SafeString(assetPair, MkString("amountAsset"), MkString("WAVES")))
	_ = amountCurrency
	price := this.ParseNumber(this.PriceFromPrecision(symbol, priceString))
	_ = price
	amount := this.ParseNumber(this.CurrencyFromPrecision(amountCurrency, amountString))
	_ = amount
	filled := this.ParseNumber(this.CurrencyFromPrecision(amountCurrency, filledString))
	_ = filled
	average := this.ParseNumber(this.PriceFromPrecision(symbol, this.SafeString(order, MkString("avgWeighedPrice"))))
	_ = average
	status := this.ParseOrderStatus(this.SafeString(order, MkString("status")))
	_ = status
	fee := OpCopy(MkUndefined())
	_ = fee
	if IsTrue(OpHasMember(MkString("type"), order)) {
		currency := this.SafeCurrencyCode(this.SafeString(order, MkString("feeAsset")))
		_ = currency
		fee = MkMap(&VarMap{
			"currency": currency,
			"fee":      this.ParseNumber(this.CurrencyFromPrecision(currency, this.SafeString(order, MkString("filledFee")))),
		})
	} else {
		currency := this.SafeCurrencyCode(this.SafeString(order, MkString("matcherFeeAssetId"), MkString("WAVES")))
		_ = currency
		fee = MkMap(&VarMap{
			"currency": currency,
			"fee":      this.ParseNumber(this.CurrencyFromPrecision(currency, this.SafeString(order, MkString("matcherFee")))),
		})
	}
	return this.SafeOrder(MkMap(&VarMap{
		"info":               order,
		"id":                 id,
		"clientOrderId":      MkUndefined(),
		"timestamp":          timestamp,
		"datetime":           this.Iso8601(timestamp),
		"lastTradeTimestamp": MkUndefined(),
		"symbol":             symbol,
		"type":               type_,
		"timeInForce":        MkUndefined(),
		"postOnly":           MkUndefined(),
		"side":               side,
		"price":              price,
		"stopPrice":          MkUndefined(),
		"amount":             amount,
		"cost":               MkUndefined(),
		"average":            average,
		"filled":             filled,
		"remaining":          MkUndefined(),
		"status":             status,
		"fee":                fee,
		"trades":             MkUndefined(),
	}))
}

func (this *Wavesexchange) GetWavesAddress(goArgs ...*Variant) *Variant {
	cachedAddreess := this.SafeString(*this.At(MkString("options")), MkString("wavesAddress"))
	_ = cachedAddreess
	if IsTrue(OpEq2(cachedAddreess, MkUndefined())) {
		request := MkMap(&VarMap{"publicKey": *this.At(MkString("apiKey"))})
		_ = request
		response := this.Call(MkString("nodeGetAddressesPublicKeyPublicKey"), request)
		_ = response
		*(*this.At(MkString("options"))).At(MkString("wavesAddress")) = this.SafeString(response, MkString("address"))
		return *(*this.At(MkString("options"))).At(MkString("wavesAddress"))
	} else {
		return cachedAddreess
	}
	return MkUndefined()
}

func (this *Wavesexchange) FetchBalance(goArgs ...*Variant) *Variant {
	params := GoGetArg(goArgs, 0, MkMap(&VarMap{}))
	_ = params
	this.CheckRequiredDependencies()
	this.CheckRequiredKeys()
	this.LoadMarkets()
	wavesAddress := this.GetWavesAddress()
	_ = wavesAddress
	request := MkMap(&VarMap{"address": wavesAddress})
	_ = request
	totalBalance := this.Call(MkString("nodeGetAssetsBalanceAddress"), request)
	_ = totalBalance
	balances := this.SafeValue(totalBalance, MkString("balances"))
	_ = balances
	result := MkMap(&VarMap{})
	_ = result
	timestamp := OpCopy(MkUndefined())
	_ = timestamp
	for i := MkInteger(0); IsTrue(OpLw(i, balances.Length)); OpIncr(&i) {
		entry := *(balances).At(i)
		_ = entry
		entryTimestamp := this.SafeInteger(entry, MkString("timestamp"))
		_ = entryTimestamp
		timestamp = OpTrinary(OpEq2(timestamp, MkUndefined()), entryTimestamp, Math.Max(timestamp, entryTimestamp))
		issueTransaction := this.SafeValue(entry, MkString("issueTransaction"))
		_ = issueTransaction
		decimals := this.SafeInteger(issueTransaction, MkString("decimals"))
		_ = decimals
		currencyId := this.SafeString(entry, MkString("assetId"))
		_ = currencyId
		balance := this.SafeString(entry, MkString("balance"))
		_ = balance
		code := OpCopy(MkUndefined())
		_ = code
		if IsTrue(OpHasMember(currencyId, *this.At(MkString("currencies_by_id")))) {
			code = this.SafeCurrencyCode(currencyId)
			*(result).At(code) = this.Account()
			*(*(result).At(code)).At(MkString("total")) = this.FromPrecision(balance, decimals)
		}
	}
	currentTimestamp := this.Milliseconds()
	_ = currentTimestamp
	byteArray := MkArray(&VarArray{
		this.Base58ToBinary(*this.At(MkString("apiKey"))),
		this.NumberToBE(currentTimestamp, MkInteger(8)),
	})
	_ = byteArray
	binary := this.BinaryConcatArray(byteArray)
	_ = binary
	hexSecret := this.BinaryToBase16(this.Base58ToBinary(*this.At(MkString("secret"))))
	_ = hexSecret
	signature := this.Eddsa(this.BinaryToBase16(binary), hexSecret, MkString("ed25519"))
	_ = signature
	matcherRequest := MkMap(&VarMap{
		"publicKey": *this.At(MkString("apiKey")),
		"signature": signature,
		"timestamp": currentTimestamp.ToString(),
	})
	_ = matcherRequest
	reservedBalance := this.Call(MkString("matcherGetMatcherBalanceReservedPublicKey"), matcherRequest)
	_ = reservedBalance
	reservedKeys := GoGetKeys(reservedBalance)
	_ = reservedKeys
	for i := MkInteger(0); IsTrue(OpLw(i, reservedKeys.Length)); OpIncr(&i) {
		currencyId := *(reservedKeys).At(i)
		_ = currencyId
		code := this.SafeCurrencyCode(currencyId)
		_ = code
		if IsTrue(OpNot(OpHasMember(code, result))) {
			*(result).At(code) = this.Account()
		}
		amount := this.SafeString(reservedBalance, currencyId)
		_ = amount
		*(*(result).At(code)).At(MkString("used")) = this.CurrencyFromPrecision(code, amount)
	}
	wavesRequest := MkMap(&VarMap{"address": wavesAddress})
	_ = wavesRequest
	wavesTotal := this.Call(MkString("nodeGetAddressesBalanceAddress"), wavesRequest)
	_ = wavesTotal
	*(result).At(MkString("WAVES")) = this.SafeValue(result, MkString("WAVES"), MkMap(&VarMap{}))
	*(*(result).At(MkString("WAVES"))).At(MkString("total")) = this.CurrencyFromPrecision(MkString("WAVES"), this.SafeString(wavesTotal, MkString("balance")))
	codes := GoGetKeys(result)
	_ = codes
	for i := MkInteger(0); IsTrue(OpLw(i, codes.Length)); OpIncr(&i) {
		code := *(codes).At(i)
		_ = code
		if IsTrue(OpEq2(this.SafeValue(*(result).At(code), MkString("used")), MkUndefined())) {
			*(*(result).At(code)).At(MkString("used")) = MkString("0")
		}
	}
	*(result).At(MkString("timestamp")) = OpCopy(timestamp)
	*(result).At(MkString("datetime")) = this.Iso8601(timestamp)
	return this.ParseBalance(result)
}

func (this *Wavesexchange) FetchMyTrades(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	address := this.GetWavesAddress()
	_ = address
	request := MkMap(&VarMap{
		"sender":      address,
		"amountAsset": *(market).At(MkString("baseId")),
		"priceAsset":  *(market).At(MkString("quoteId")),
	})
	_ = request
	response := this.Call(MkString("publicGetTransactionsExchange"), request)
	_ = response
	data := this.SafeValue(response, MkString("data"))
	_ = data
	return this.ParseTrades(data, market, since, limit)
}

func (this *Wavesexchange) FetchTrades(goArgs ...*Variant) *Variant {
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	_ = symbol
	since := GoGetArg(goArgs, 1, MkUndefined())
	_ = since
	limit := GoGetArg(goArgs, 2, MkUndefined())
	_ = limit
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	_ = params
	this.LoadMarkets()
	market := this.Market(symbol)
	_ = market
	request := MkMap(&VarMap{
		"amountAsset": *(market).At(MkString("baseId")),
		"priceAsset":  *(market).At(MkString("quoteId")),
	})
	_ = request
	if IsTrue(OpNotEq2(limit, MkUndefined())) {
		*(request).At(MkString("limit")) = OpCopy(limit)
	}
	if IsTrue(OpNotEq2(since, MkUndefined())) {
		*(request).At(MkString("timeStart")) = OpCopy(since)
	}
	response := this.Call(MkString("publicGetTransactionsExchange"), request)
	_ = response
	data := this.SafeValue(response, MkString("data"))
	_ = data
	return this.ParseTrades(data, market, since, limit)
}

func (this *Wavesexchange) ParseTrade(goArgs ...*Variant) *Variant {
	trade := GoGetArg(goArgs, 0, MkUndefined())
	_ = trade
	market := GoGetArg(goArgs, 1, MkUndefined())
	_ = market
	data := this.SafeValue(trade, MkString("data"))
	_ = data
	datetime := this.SafeString(data, MkString("timestamp"))
	_ = datetime
	timestamp := this.Parse8601(datetime)
	_ = timestamp
	id := this.SafeString(data, MkString("id"))
	_ = id
	priceString := this.SafeString(data, MkString("price"))
	_ = priceString
	amountString := this.SafeString(data, MkString("amount"))
	_ = amountString
	price := this.ParseNumber(priceString)
	_ = price
	amount := this.ParseNumber(amountString)
	_ = amount
	cost := this.ParseNumber(Precise.StringMul(priceString, amountString))
	_ = cost
	order1 := this.SafeValue(data, MkString("order1"))
	_ = order1
	order2 := this.SafeValue(data, MkString("order2"))
	_ = order2
	order := OpCopy(MkUndefined())
	_ = order
	if IsTrue(OpEq2(this.SafeString(order1, MkString("senderPublicKey")), *this.At(MkString("apiKey")))) {
		order = OpCopy(order1)
	} else {
		order = OpCopy(order2)
	}
	symbol := OpCopy(MkUndefined())
	_ = symbol
	assetPair := this.SafeValue(order, MkString("assetPair"))
	_ = assetPair
	if IsTrue(OpNotEq2(assetPair, MkUndefined())) {
		symbol = this.GetSymbolFromAssetPair(assetPair)
	} else {
		if IsTrue(OpNotEq2(market, MkUndefined())) {
			symbol = *(market).At(MkString("symbol"))
		}
	}
	side := this.SafeString(order, MkString("orderType"))
	_ = side
	orderId := this.SafeString(order, MkString("id"))
	_ = orderId
	fee := MkMap(&VarMap{
		"cost":     this.SafeNumber(order, MkString("matcherFee")),
		"currency": this.SafeCurrencyCode(this.SafeString(order, MkString("matcherFeeAssetId"), MkString("WAVES"))),
	})
	_ = fee
	return MkMap(&VarMap{
		"info":         trade,
		"timestamp":    timestamp,
		"datetime":     datetime,
		"symbol":       symbol,
		"id":           id,
		"order":        orderId,
		"type":         MkUndefined(),
		"side":         side,
		"takerOrMaker": MkUndefined(),
		"price":        price,
		"amount":       amount,
		"cost":         cost,
		"fee":          fee,
	})
}

func (this *Wavesexchange) HandleErrors(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	reason := GoGetArg(goArgs, 1, MkUndefined())
	_ = reason
	url := GoGetArg(goArgs, 2, MkUndefined())
	_ = url
	method := GoGetArg(goArgs, 3, MkUndefined())
	_ = method
	headers := GoGetArg(goArgs, 4, MkUndefined())
	_ = headers
	body := GoGetArg(goArgs, 5, MkUndefined())
	_ = body
	response := GoGetArg(goArgs, 6, MkUndefined())
	_ = response
	requestHeaders := GoGetArg(goArgs, 7, MkUndefined())
	_ = requestHeaders
	requestBody := GoGetArg(goArgs, 8, MkUndefined())
	_ = requestBody
	errorCode := this.SafeString(response, MkString("error"))
	_ = errorCode
	success := this.SafeValue(response, MkString("success"), MkBool(true))
	_ = success
	Exception := this.SafeValue(*this.At(MkString("exceptions")), errorCode)
	_ = Exception
	if IsTrue(OpNotEq2(Exception, MkUndefined())) {
		message := this.SafeString(response, MkString("message"))
		_ = message
		panic(NewException(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), message))))
	}
	message := this.SafeString(response, MkString("message"))
	_ = message
	if IsTrue(OpEq2(message, MkString("Validation Error"))) {
		panic(NewBadRequest(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))))
	}
	if IsTrue(OpNot(success)) {
		panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), body))))
	}
	return MkUndefined()
}

func (this *Wavesexchange) Withdraw(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	amount := GoGetArg(goArgs, 1, MkUndefined())
	_ = amount
	address := GoGetArg(goArgs, 2, MkUndefined())
	_ = address
	tag := GoGetArg(goArgs, 3, MkUndefined())
	_ = tag
	params := GoGetArg(goArgs, 4, MkMap(&VarMap{}))
	_ = params
	if IsTrue(OpNotEq2(code, MkString("WAVES"))) {
		supportedCurrencies := this.Call(MkString("privateGetWithdrawCurrencies"))
		_ = supportedCurrencies
		currencies := MkMap(&VarMap{})
		_ = currencies
		items := this.SafeValue(supportedCurrencies, MkString("items"), MkArray(&VarArray{}))
		_ = items
		for i := MkInteger(0); IsTrue(OpLw(i, items.Length)); OpIncr(&i) {
			entry := *(items).At(i)
			_ = entry
			currencyCode := this.SafeString(entry, MkString("id"))
			_ = currencyCode
			*(currencies).At(currencyCode) = OpCopy(MkBool(true))
		}
		if IsTrue(OpNot(OpHasMember(code, currencies))) {
			codes := GoGetKeys(currencies)
			_ = codes
			panic(NewExchangeError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" fetch "), OpAdd(code, OpAdd(MkString(" withdrawals are not supported. Currency code must be one of "), codes.ToString()))))))
		}
	}
	this.LoadMarkets()
	hexChars := MkArray(&VarArray{
		MkString("0"),
		MkString("1"),
		MkString("2"),
		MkString("3"),
		MkString("4"),
		MkString("5"),
		MkString("6"),
		MkString("7"),
		MkString("8"),
		MkString("9"),
		MkString("a"),
		MkString("b"),
		MkString("c"),
		MkString("d"),
		MkString("e"),
		MkString("f"),
	})
	_ = hexChars
	set := MkMap(&VarMap{})
	_ = set
	for i := MkInteger(0); IsTrue(OpLw(i, hexChars.Length)); OpIncr(&i) {
		key := *(hexChars).At(i)
		_ = key
		*(set).At(key) = OpCopy(MkBool(true))
	}
	isErc20 := OpCopy(MkBool(true))
	_ = isErc20
	noPrefix := this.Remove0xPrefix(address)
	_ = noPrefix
	lower := noPrefix.ToLowerCase()
	_ = lower
	for i := MkInteger(0); IsTrue(OpLw(i, lower.Length)); OpIncr(&i) {
		character := *(lower).At(i)
		_ = character
		if IsTrue(OpNot(OpHasMember(character, set))) {
			isErc20 = OpCopy(MkBool(false))
			break
		}
	}
	this.GetAccessToken()
	proxyAddress := OpCopy(MkUndefined())
	_ = proxyAddress
	if IsTrue(OpAnd(OpEq2(code, MkString("WAVES")), OpNot(isErc20))) {
		proxyAddress = OpCopy(address)
	} else {
		withdrawAddressRequest := MkMap(&VarMap{
			"address":  address,
			"currency": code,
		})
		_ = withdrawAddressRequest
		withdrawAddress := this.Call(MkString("privateGetWithdrawAddressesCurrencyAddress"), withdrawAddressRequest)
		_ = withdrawAddress
		currency := this.SafeValue(withdrawAddress, MkString("currency"))
		_ = currency
		allowedAmount := this.SafeValue(currency, MkString("allowed_amount"))
		_ = allowedAmount
		minimum := this.SafeNumber(allowedAmount, MkString("min"))
		_ = minimum
		if IsTrue(OpLwEq(amount, minimum)) {
			panic(NewBadRequest(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" "), OpAdd(code, OpAdd(MkString(" withdraw failed, amount "), OpAdd(amount.ToString(), OpAdd(MkString(" must be greater than the minimum allowed amount of "), minimum.ToString()))))))))
		}
		proxyAddresses := this.SafeValue(withdrawAddress, MkString("proxy_addresses"), MkArray(&VarArray{}))
		_ = proxyAddresses
		proxyAddress = this.SafeString(proxyAddresses, MkInteger(0))
	}
	fee := this.SafeInteger(*this.At(MkString("options")), MkString("withdrawFeeWAVES"), MkInteger(100000))
	_ = fee
	feeAssetId := MkString("WAVES")
	_ = feeAssetId
	type_ := MkInteger(4)
	_ = type_
	version := MkInteger(2)
	_ = version
	amountInteger := this.CurrencyToPrecision(code, amount)
	_ = amountInteger
	currency := this.Currency(code)
	_ = currency
	timestamp := this.Milliseconds()
	_ = timestamp
	byteArray := MkArray(&VarArray{
		this.NumberToBE(MkInteger(4), MkInteger(1)),
		this.NumberToBE(MkInteger(2), MkInteger(1)),
		this.Base58ToBinary(*this.At(MkString("apiKey"))),
		this.GetAssetBytes(*(currency).At(MkString("id"))),
		this.GetAssetBytes(feeAssetId),
		this.NumberToBE(timestamp, MkInteger(8)),
		this.NumberToBE(amountInteger, MkInteger(8)),
		this.NumberToBE(fee, MkInteger(8)),
		this.Base58ToBinary(proxyAddress),
		this.NumberToBE(MkInteger(0), MkInteger(2)),
	})
	_ = byteArray
	binary := this.BinaryConcatArray(byteArray)
	_ = binary
	hexSecret := this.BinaryToBase16(this.Base58ToBinary(*this.At(MkString("secret"))))
	_ = hexSecret
	signature := this.Eddsa(this.BinaryToBase16(binary), hexSecret, MkString("ed25519"))
	_ = signature
	request := MkMap(&VarMap{
		"senderPublicKey": *this.At(MkString("apiKey")),
		"amount":          amountInteger,
		"fee":             fee,
		"type":            type_,
		"version":         version,
		"attachment":      MkString(""),
		"feeAssetId":      this.GetAssetId(feeAssetId),
		"proofs": MkArray(&VarArray{
			signature,
		}),
		"assetId":   this.GetAssetId(*(currency).At(MkString("id"))),
		"recipient": proxyAddress,
		"timestamp": timestamp,
		"signature": signature,
	})
	_ = request
	return this.Call(MkString("nodePostTransactionsBroadcast"), request)
}
