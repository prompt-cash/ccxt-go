package ccxt

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

func (this *ExchangeBase) Account() *Variant {
	return MkMap(&VarMap{
		"free":  MkUndefined(),
		"used":  MkUndefined(),
		"total": MkUndefined(),
	})
}

func (this *ExchangeBase) FilterByArray(goArgs ...*Variant) *Variant {
	// transpiled function
	objects := GoGetArg(goArgs, 0, MkUndefined())
	key := GoGetArg(goArgs, 1, MkUndefined())
	values := GoGetArg(goArgs, 2, MkUndefined())
	indexed := GoGetArg(goArgs, 3, MkBool(true))
	objects = GoGetValues(objects)
	if IsTrue(OpOr(OpEq2(values, MkUndefined()), OpEq2(values, MkUndefined()))) {
		return OpTrinary(indexed, this.IndexBy(objects, key), objects)
	}
	result := MkArray(&VarArray{})
	_ = result
	for i := MkInteger(0); IsTrue(OpLw(i, objects.Length)); OpIncr(&i) {
		if IsTrue(OpNotEq(values.IndexOf(*(*(objects).At(i)).At(key)), MkInteger(-1))) {
			result.Push(*(objects).At(i))
		}
	}
	return OpTrinary(indexed, this.IndexBy(result, key), result)
}

func (this *ExchangeBase) Currency(code *Variant) *Variant {

	vCurrencies := *this.At(MkString("currencies"))
	if vCurrencies.Type == Undefined {
		panic(NewExchangeError(MkString(this.Id() + " currencies not loaded")))
	}

	if code.Type == String {
		if vCurrencies.Has(code.ToStr()) {
			return *vCurrencies.At(code)
		}
	}

	panic(NewExchangeError(MkString(this.Id() + " does not have currency code " + code.ToStr())))
}

func (this *ExchangeBase) SafeCurrency(goArgs ...*Variant) *Variant {
	// transpiled function
	currencyId := GoGetArg(goArgs, 0, MkUndefined())
	currency := GoGetArg(goArgs, 1, MkUndefined())
	if IsTrue(OpAnd(OpEq2(currencyId, MkUndefined()), OpNotEq2(currency, MkUndefined()))) {
		return currency
	}
	if IsTrue(OpAnd(OpNotEq2(*this.At(MkString("currencies_by_id")), MkUndefined()), OpHasMember(currencyId, *this.At(MkString("currencies_by_id"))))) {
		return *(*this.At(MkString("currencies_by_id"))).At(currencyId)
	}
	return MkMap(&VarMap{
		"id":   currencyId,
		"code": OpTrinary(OpEq2(currencyId, MkUndefined()), currencyId, this.CommonCurrencyCode(currencyId.ToUpperCase())),
	})
}

func (this *ExchangeBase) commonCurrencyCode(currency *Variant) *Variant {
	return this.SafeString(*this.At(MkString("commonCurrencies")), currency, currency)
}

func (this *ExchangeBase) SafeCurrencyCode(v ...*Variant) *Variant {
	m := this.SafeCurrency(v...)
	return *m.At(MkString("code"))
}

func (this *ExchangeBase) Market(symbol *Variant) *Variant {

	vMarkets := *this.At(MkString("markets"))
	if vMarkets.Type == Undefined {
		panic(NewExchangeError(MkString(this.Id() + " markets not loaded")))
	}

	if symbol.Type == String {
		if vMarkets.Has(symbol.ToStr()) {
			return *vMarkets.At(symbol)
		} else if (*this.At(MkString("markets_by_id"))).Has(symbol.ToStr()) {
			return *(*this.At(MkString("markets_by_id"))).At(symbol)
		}
	}

	panic(NewBadSymbol(MkString(this.Id() + " does not have market symbol " + symbol.ToStr())))
}

func (this *ExchangeBase) SafeMarket(goArgs ...*Variant) *Variant {
	// transpiled function
	marketId := GoGetArg(goArgs, 0, MkUndefined())
	market := GoGetArg(goArgs, 1, MkUndefined())
	delimiter := GoGetArg(goArgs, 2, MkUndefined())
	if IsTrue(OpNotEq2(marketId, MkUndefined())) {
		if IsTrue(OpAnd(OpNotEq2((*this.At(MkString("markets_by_id"))), MkUndefined()), OpHasMember(marketId, (*this.At(MkString("markets_by_id")))))) {
			market = *(*this.At(MkString("markets_by_id"))).At(marketId)
		} else {
			if IsTrue(OpNotEq2(delimiter, MkUndefined())) {
				baseId := MkUndefined()
				quoteId := MkUndefined()
				ArrayUnpack(marketId.Split(delimiter), &baseId, &quoteId)
				base := this.SafeCurrencyCode(baseId)
				quote := this.SafeCurrencyCode(quoteId)
				symbol := OpAdd(base, OpAdd(MkString("/"), quote))
				return MkMap(&VarMap{
					"id":      marketId,
					"symbol":  symbol,
					"base":    base,
					"quote":   quote,
					"baseId":  baseId,
					"quoteId": quoteId,
				})
			}
		}
	}
	if IsTrue(OpNotEq2(market, MkUndefined())) {
		return market
	}
	return MkMap(&VarMap{
		"id":      marketId,
		"symbol":  marketId,
		"base":    MkUndefined(),
		"quote":   MkUndefined(),
		"baseId":  MkUndefined(),
		"quoteId": MkUndefined(),
	})
}

func (this *ExchangeBase) MarketId(symbol *Variant) *Variant {
	vMarket := this.Market(symbol)
	if vMarket.Type == Undefined {
		return OpCopy(symbol)
	}
	return OpCopy(*vMarket.At(MkString("id")))
}

func (this *ExchangeBase) MarketIds(symbols *Variant) *Variant {
	if symbols.Type != Array {
		return MkUndefined()
	}

	vMap := VarMap{}
	for _, symbol := range *symbols.ToArray() {
		vMap[(*symbol).ToStr()] = this.MarketId(*symbol)
	}
	return MkMap(&vMap)
}

func (this *ExchangeBase) LoadMarkets() *Variant {
	markets := *this.At(MkString("markets"))
	if markets.Type != Undefined {
		if (*this.At(MkString("markets_by_id"))).Type == Undefined {
			return this.SetMarkets(markets, nil)
		}
		return markets
	}
	currencies := MkUndefined()
	if (*this.At(MkString("has"))).Has("fetchCurrencies") {
		currencies = this.VCall("FetchCurrencies") // call function on derived exctange
	}
	markets = this.VCall("FetchMarkets") // call function on derived exctange
	return this.SetMarkets(markets, currencies)
}

func (this *ExchangeBase) SetMarkets(markets *Variant, currencies *Variant) *Variant {
	// partially transpiled function
	values := GoGetValues(markets).Map(func(market *Variant) *Variant {
		return this.DeepExtend(MkMap(&VarMap{"limits": *this.At(MkString("limits")), "precision": *this.At(MkString("precision"))}), *(*this.At(MkString("fees"))).At(MkString("trading")), market)
	})
	*this.At(MkString("markets")) = this.IndexBy(values, MkString("symbol"))
	*this.At(MkString("markets_by_id")) = this.IndexBy(markets, MkString("id"))
	*this.At(MkString("symbols")) = (GoGetKeys(*this.At(MkString("markets")))).Sort()
	*this.At(MkString("ids")) = (GoGetKeys((*this.At(MkString("markets_by_id"))))).Sort()
	if IsTrue(currencies) {
		*this.At(MkString("currencies")) = this.DeepExtend(currencies, *this.At(MkString("currencies")))
	} else {
		baseCurrencies := values.Filter(func(market *Variant) *Variant {
			return MkBool(market.Has("base"))
		}).Map(func(market *Variant) *Variant {
			id := *market.At(MkString("baseId"))
			if id.Type == Undefined {
				id = *market.At(MkString("base"))
			}
			precisionObj := *market.At(MkString("precision"))
			precision := *precisionObj.At(MkString("base"))
			if precision.Type == Undefined {
				precision = *precisionObj.At(MkString("amount"))
			}
			if precision.Type == Undefined {
				precision = MkInteger(8)
			}
			return MkMap(&VarMap{
				"id":        id,
				"numericId": *market.At(MkString("baseNumericId")),
				"code":      *market.At(MkString("base")),
				"precision": precision,
			})
		})
		quoteCurrencies := values.Filter(func(market *Variant) *Variant {
			return MkBool(market.Has("base"))
		}).Map(func(market *Variant) *Variant {
			id := *market.At(MkString("baseId"))
			if id.Type == Undefined {
				id = *market.At(MkString("base"))
			}
			precisionObj := *market.At(MkString("precision"))
			precision := *precisionObj.At(MkString("quote"))
			if precision.Type == Undefined {
				precision = *precisionObj.At(MkString("price"))
			}
			if precision.Type == Undefined {
				precision = MkInteger(8)
			}
			return MkMap(&VarMap{
				"id":        id,
				"numericId": *market.At(MkString("baseNumericId")),
				"code":      *market.At(MkString("base")),
				"precision": precision,
			})
		})
		baseCurrencies = this.SortBy(baseCurrencies, MkString("code"))
		quoteCurrencies = this.SortBy(quoteCurrencies, MkString("code"))
		*this.At(MkString("baseCurrencies")) = this.IndexBy(baseCurrencies, MkString("code"))
		*this.At(MkString("quoteCurrencies")) = this.IndexBy(quoteCurrencies, MkString("code"))
		currencies := baseCurrencies.Concat(quoteCurrencies)
		sortedCurrencies := this.SortBy(this.Flatten(currencies), MkString("code"))
		*this.At(MkString("currencies")) = this.DeepExtend(this.IndexBy(sortedCurrencies, MkString("code")), *this.At(MkString("currencies")))
	}
	*this.At(MkString("currencies_by_id")) = this.IndexBy(*this.At(MkString("currencies")), MkString("id"))
	return *this.At(MkString("markets"))
}

func (this *ExchangeBase) SafeSymbol(v ...*Variant) *Variant {
	m := this.SafeMarket(v...)
	return *m.At(MkString("symbol"))
}

func (this *ExchangeBase) CheckRequiredCredentials(goArgs ...*Variant) *Variant {
	// transpiled function
	error := GoGetArg(goArgs, 0, MkBool(true))
	keys := GoGetKeys(*this.At(MkString("requiredCredentials")))
	for i := MkInteger(0); IsTrue(OpLw(i, keys.Length)); OpIncr(&i) {
		key := *(keys).At(i)
		if IsTrue(OpAnd(*(*this.At(MkString("requiredCredentials"))).At(key), OpNot(*(this).At(key)))) {
			if IsTrue(error) {
				panic(NewAuthenticationError(OpAdd(*this.At(MkString("id")), OpAdd(MkString(" requires `"), OpAdd(key, MkString("` credential"))))))
			} else {
				return error
			}
		}
	}
	return MkBool(true)
}

func (this *ExchangeBase) CheckAddress(addr *Variant) *Variant {

	if addr.Type == Undefined {
		panic(NewInvalidAddress(MkString(this.Id() + " address is undefined")))
	}

	strAddr := addr.ToStr()
	// check the address is not the same letter like 'aaaaa' nor too short nor has a space
	if /*(unique (address).length === 1) ||*/ len(strAddr) < int((*this.At(MkString("minFundingAddressLength"))).ToInt()) || strings.Index(strAddr, " ") != -1 {
		panic(NewInvalidAddress(MkString(this.Id() + " address is invalid or has less than the minimum characters: " + addr.ToStr())))
	}

	return addr
}

func (this *ExchangeBase) ParseTrades(goArgs ...*Variant) *Variant {
	// partially transpiled function
	trades := GoGetArg(goArgs, 0, MkUndefined())
	market := GoGetArg(goArgs, 1, MkUndefined())
	since := GoGetArg(goArgs, 2, MkUndefined())
	limit := GoGetArg(goArgs, 3, MkUndefined())
	params := GoGetArg(goArgs, 4, MkMap(&VarMap{}))
	if !IsTrue(trades) {
		trades = MkArray(&VarArray{})
	}
	result := GoGetValues(trades).Map(func(trade *Variant) *Variant {
		return this.Extend(this.VCall("ParseTrade", trade, market), params)
	})
	result = this.SortBy(result, MkString("timestamp"))
	symbol := OpTrinary(OpNotEq2(market, MkUndefined()), *(market).At(MkString("symbol")), MkUndefined())
	tail := OpEq2(since, MkUndefined())
	return this.FilterBySymbolSinceLimit(result, symbol, since, limit, tail)
}

func (this *ExchangeBase) ParseTransactions(goArgs ...*Variant) *Variant {
	// partially transpiled function
	transactions := GoGetArg(goArgs, 0, MkUndefined())
	currency := GoGetArg(goArgs, 1, MkUndefined())
	since := GoGetArg(goArgs, 2, MkUndefined())
	limit := GoGetArg(goArgs, 3, MkUndefined())
	params := GoGetArg(goArgs, 4, MkMap(&VarMap{}))
	if !IsTrue(transactions) {
		transactions = MkArray(&VarArray{})
	}
	result := GoGetValues(transactions).Map(func(transaction *Variant) *Variant {
		return this.Extend(this.VCall("ParseTransaction", transaction, currency), params)
	})
	result = this.SortBy(result, MkString("timestamp"))
	code := OpTrinary(OpNotEq2(currency, MkUndefined()), *(currency).At(MkString("code")), MkUndefined())
	tail := OpEq2(since, MkUndefined())
	return this.FilterByCurrencySinceLimit(result, code, since, limit, tail)
}

func (this *ExchangeBase) ParseBalance(goArgs ...*Variant) *Variant {
	// transpiled function
	balance := GoGetArg(goArgs, 0, MkUndefined())
	legacy := GoGetArg(goArgs, 1, MkBool(false))
	codes := GoGetKeys(this.Omit(balance, MkArray(&VarArray{
		MkString("info"),
		MkString("timestamp"),
		MkString("datetime"),
		MkString("free"),
		MkString("used"),
		MkString("total"),
	})))
	*(balance).At(MkString("free")) = MkMap(&VarMap{})
	*(balance).At(MkString("used")) = MkMap(&VarMap{})
	*(balance).At(MkString("total")) = MkMap(&VarMap{})
	for i := MkInteger(0); IsTrue(OpLw(i, codes.Length)); OpIncr(&i) {
		code := *(codes).At(i)
		if IsTrue(OpEq2(*(*(balance).At(code)).At(MkString("total")), MkUndefined())) {
			if IsTrue(OpAnd(OpNotEq2(*(*(balance).At(code)).At(MkString("free")), MkUndefined()), OpNotEq2(*(*(balance).At(code)).At(MkString("used")), MkUndefined()))) {
				if IsTrue(legacy) {
					*(*(balance).At(code)).At(MkString("total")) = this.Sum(*(*(balance).At(code)).At(MkString("free")), *(*(balance).At(code)).At(MkString("used")))
				} else {
					*(*(balance).At(code)).At(MkString("total")) = Precise.StringAdd(*(*(balance).At(code)).At(MkString("free")), *(*(balance).At(code)).At(MkString("used")))
				}
			}
		}
		if IsTrue(OpEq2(*(*(balance).At(code)).At(MkString("free")), MkUndefined())) {
			if IsTrue(OpAnd(OpNotEq2(*(*(balance).At(code)).At(MkString("total")), MkUndefined()), OpNotEq2(*(*(balance).At(code)).At(MkString("used")), MkUndefined()))) {
				if IsTrue(legacy) {
					*(*(balance).At(code)).At(MkString("free")) = this.Sum(*(*(balance).At(code)).At(MkString("total")), OpNeg(*(*(balance).At(code)).At(MkString("used"))))
				} else {
					*(*(balance).At(code)).At(MkString("free")) = Precise.StringSub(*(*(balance).At(code)).At(MkString("total")), *(*(balance).At(code)).At(MkString("used")))
				}
			}
		}
		if IsTrue(OpEq2(*(*(balance).At(code)).At(MkString("used")), MkUndefined())) {
			if IsTrue(OpAnd(OpNotEq2(*(*(balance).At(code)).At(MkString("total")), MkUndefined()), OpNotEq2(*(*(balance).At(code)).At(MkString("free")), MkUndefined()))) {
				if IsTrue(legacy) {
					*(*(balance).At(code)).At(MkString("used")) = this.Sum(*(*(balance).At(code)).At(MkString("total")), OpNeg(*(*(balance).At(code)).At(MkString("free"))))
				} else {
					*(*(balance).At(code)).At(MkString("used")) = Precise.StringSub(*(*(balance).At(code)).At(MkString("total")), *(*(balance).At(code)).At(MkString("free")))
				}
			}
		}
		*(*(balance).At(code)).At(MkString("free")) = this.ParseNumber(*(*(balance).At(code)).At(MkString("free")))
		*(*(balance).At(code)).At(MkString("used")) = this.ParseNumber(*(*(balance).At(code)).At(MkString("used")))
		*(*(balance).At(code)).At(MkString("total")) = this.ParseNumber(*(*(balance).At(code)).At(MkString("total")))
		*(*balance.At(MkString("Free"))).At(code) = *(*(balance).At(code)).At(MkString("free"))
		*(*balance.At(MkString("Used"))).At(code) = *(*(balance).At(code)).At(MkString("used"))
		*(*balance.At(MkString("Total"))).At(code) = *(*(balance).At(code)).At(MkString("total"))
	}
	return balance
}

func (this *ExchangeBase) ParseOrders(goArgs ...*Variant) *Variant {
	// partially transpiled function
	orders := GoGetArg(goArgs, 0, MkUndefined())
	market := GoGetArg(goArgs, 1, MkUndefined())
	since := GoGetArg(goArgs, 2, MkUndefined())
	limit := GoGetArg(goArgs, 3, MkUndefined())
	params := GoGetArg(goArgs, 4, MkMap(&VarMap{}))
	var result *Variant
	if orders.Type == Array {
		result = GoGetValues(orders).Map(func(order *Variant) *Variant {
			return this.Extend(this.VCall("ParseOrder", order, market), params)
		})
	} else {
		vArray := VarArray{}
		i := 0
		for key, value := range *orders.ToMap() {
			vArray[i] = this.Extend(this.VCall("ParseOrder", this.Extend(MkMap(&VarMap{"id": MkString(key)}), *value, market), params))
			i++
		}
		result = MkArray(&vArray)
	}
	result = this.SortBy(result, MkString("timestamp"))
	symbol := OpTrinary(OpNotEq2(market, MkUndefined()), *(market).At(MkString("symbol")), MkUndefined())
	tail := OpEq2(since, MkUndefined())
	return this.FilterBySymbolSinceLimit(result, symbol, since, limit, tail)
}

func (this *ExchangeBase) ParseOrderBook(goArgs ...*Variant) *Variant {
	// transpiled function
	orderbook := GoGetArg(goArgs, 0, MkUndefined())
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	timestamp := GoGetArg(goArgs, 2, MkUndefined())
	bidsKey := GoGetArg(goArgs, 3, MkString("bids"))
	asksKey := GoGetArg(goArgs, 4, MkString("asks"))
	priceKey := GoGetArg(goArgs, 5, MkInteger(0))
	amountKey := GoGetArg(goArgs, 6, MkInteger(1))
	return MkMap(&VarMap{
		"symbol":    symbol,
		"bids":      this.SortBy(OpTrinary(OpHasMember(bidsKey, orderbook), this.ParseBidsAsks(*(orderbook).At(bidsKey), priceKey, amountKey), MkArray(&VarArray{})), MkInteger(0), MkBool(true)),
		"asks":      this.SortBy(OpTrinary(OpHasMember(asksKey, orderbook), this.ParseBidsAsks(*(orderbook).At(asksKey), priceKey, amountKey), MkArray(&VarArray{})), MkInteger(0)),
		"timestamp": timestamp,
		"datetime":  this.Iso8601(timestamp),
		"nonce":     MkUndefined(),
	})
}

func (this *ExchangeBase) FilterByValueSinceLimit(goArgs ...*Variant) *Variant {
	// transpiled function
	array := GoGetArg(goArgs, 0, MkUndefined())
	field := GoGetArg(goArgs, 1, MkUndefined())
	value := GoGetArg(goArgs, 2, MkUndefined())
	since := GoGetArg(goArgs, 3, MkUndefined())
	limit := GoGetArg(goArgs, 4, MkUndefined())
	key := GoGetArg(goArgs, 5, MkString("timestamp"))
	tail := GoGetArg(goArgs, 6, MkBool(false))
	valueIsDefined := OpAnd(OpNotEq2(value, MkUndefined()), OpNotEq2(value, MkUndefined()))
	_ = valueIsDefined
	sinceIsDefined := OpAnd(OpNotEq2(since, MkUndefined()), OpNotEq2(since, MkUndefined()))
	_ = sinceIsDefined
	if IsTrue(OpOr(valueIsDefined, sinceIsDefined)) {
		array = array.Filter(func(entry *Variant) *Variant {
			return MkBool((!IsTrue(valueIsDefined) || IsTrue(OpEq2(*entry.At(field), value))) && (!IsTrue(sinceIsDefined) || IsTrue(OpGtEq(*entry.At(key), since))))
		})
	}
	if IsTrue(OpAnd(OpNotEq2(limit, MkUndefined()), OpNotEq2(limit, MkUndefined()))) {
		array = OpTrinary(tail, array.Slice(OpNeg(limit)), array.Slice(MkInteger(0), limit))
	}
	return array
}

func (this *ExchangeBase) FilterBySymbolSinceLimit(goArgs ...*Variant) *Variant {
	// transpiled function
	array := GoGetArg(goArgs, 0, MkUndefined())
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	since := GoGetArg(goArgs, 2, MkUndefined())
	limit := GoGetArg(goArgs, 3, MkUndefined())
	tail := GoGetArg(goArgs, 4, MkBool(false))
	_ = tail
	return this.FilterByValueSinceLimit(array, MkString("symbol"), symbol, since, limit, MkString("timestamp"), tail)
}

func (this *ExchangeBase) FilterByCurrencySinceLimit(goArgs ...*Variant) *Variant {
	// transpiled function
	array := GoGetArg(goArgs, 0, MkUndefined())
	code := GoGetArg(goArgs, 1, MkUndefined())
	since := GoGetArg(goArgs, 2, MkUndefined())
	limit := GoGetArg(goArgs, 3, MkUndefined())
	tail := GoGetArg(goArgs, 4, MkBool(false))
	return this.FilterByValueSinceLimit(array, MkString("currency"), code, since, limit, MkString("timestamp"), tail)
}

func (this *ExchangeBase) ReduceFeesByCurrency(goArgs ...*Variant) *Variant {
	// transpiled function
	fees := GoGetArg(goArgs, 0, MkUndefined())
	reduced := MkMap(&VarMap{})
	for i := MkInteger(0); IsTrue(OpLw(i, fees.Length)); OpIncr(&i) {
		fee := *(fees).At(i)
		feeCurrencyCode := this.SafeValue(fee, MkString("currency"))
		if IsTrue(OpNotEq2(feeCurrencyCode, MkUndefined())) {
			if IsTrue(OpHasMember(feeCurrencyCode, reduced)) {
				*(*(reduced).At(feeCurrencyCode)).At(MkString("cost")) = this.Sum(*(*(reduced).At(feeCurrencyCode)).At(MkString("cost")), *(fee).At(MkString("cost")))
			} else {
				*(reduced).At(feeCurrencyCode) = MkMap(&VarMap{
					"cost":     *(fee).At(MkString("cost")),
					"currency": feeCurrencyCode,
				})
			}
		}
	}
	return GoGetValues(reduced)
}

func (this *ExchangeBase) SafeOrder(goArgs ...*Variant) *Variant {
	// transpiled function
	order := GoGetArg(goArgs, 0, MkUndefined())
	amount := this.SafeValue(order, MkString("amount"))
	remaining := this.SafeValue(order, MkString("remaining"))
	filled := this.SafeValue(order, MkString("filled"))
	cost := this.SafeValue(order, MkString("cost"))
	average := this.SafeValue(order, MkString("average"))
	price := this.SafeValue(order, MkString("price"))
	lastTradeTimeTimestamp := this.SafeInteger(order, MkString("lastTradeTimestamp"))
	parseFilled := OpEq2(filled, MkUndefined())
	parseCost := OpEq2(cost, MkUndefined())
	parseLastTradeTimeTimestamp := OpEq2(lastTradeTimeTimestamp, MkUndefined())
	parseFee := OpEq2(this.SafeValue(order, MkString("fee")), MkUndefined())
	parseFees := OpEq2(this.SafeValue(order, MkString("fees")), MkUndefined())
	shouldParseFees := OpOr(parseFee, parseFees)
	fees := this.SafeValue(order, MkString("fees"), MkArray(&VarArray{}))
	if IsTrue(OpOr(parseFilled, OpOr(parseCost, shouldParseFees))) {
		trades := this.SafeValue(order, MkString("trades"))
		if IsTrue(GoIsArray(trades)) {
			if IsTrue(parseFilled) {
				filled = MkInteger(0)
			}
			if IsTrue(parseCost) {
				cost = MkInteger(0)
			}
			for i := MkInteger(0); IsTrue(OpLw(i, trades.Length)); OpIncr(&i) {
				trade := *(trades).At(i)
				tradeAmount := this.SafeValue(trade, MkString("amount"))
				if IsTrue(OpAnd(parseFilled, OpNotEq2(tradeAmount, MkUndefined()))) {
					filled = this.Sum(filled, tradeAmount)
				}
				tradeCost := this.SafeValue(trade, MkString("cost"))
				if IsTrue(OpAnd(parseCost, OpNotEq2(tradeCost, MkUndefined()))) {
					cost = this.Sum(cost, tradeCost)
				}
				tradeTimestamp := this.SafeValue(trade, MkString("timestamp"))
				if IsTrue(OpAnd(parseLastTradeTimeTimestamp, OpNotEq2(tradeTimestamp, MkUndefined()))) {
					if IsTrue(OpEq2(lastTradeTimeTimestamp, MkUndefined())) {
						lastTradeTimeTimestamp = OpCopy(tradeTimestamp)
					} else {
						lastTradeTimeTimestamp = Math.Max(lastTradeTimeTimestamp, tradeTimestamp)
					}
				}
				if IsTrue(shouldParseFees) {
					tradeFees := this.SafeValue(trade, MkString("fees"))
					if IsTrue(OpNotEq2(tradeFees, MkUndefined())) {
						for j := MkInteger(0); IsTrue(OpLw(j, tradeFees.Length)); OpIncr(&j) {
							tradeFee := *(tradeFees).At(j)
							fees.Push(this.Extend(MkMap(&VarMap{}), tradeFee))
						}
					} else {
						tradeFee := this.SafeValue(trade, MkString("fee"))
						if IsTrue(OpNotEq2(tradeFee, MkUndefined())) {
							fees.Push(this.Extend(MkMap(&VarMap{}), tradeFee))
						}
					}
				}
			}
		}
	}
	if IsTrue(shouldParseFees) {
		reducedFees := OpTrinary(*this.At(MkString("reduceFees")), this.ReduceFeesByCurrency(fees), fees)
		reducedLength := OpCopy(reducedFees.Length)
		if IsTrue(OpAnd(OpNot(parseFee), OpEq2(reducedLength, MkInteger(0)))) {
			reducedFees.Push(*(order).At(MkString("fee")))
		}
		if IsTrue(parseFees) {
			*(order).At(MkString("fees")) = OpCopy(reducedFees)
		}
		if IsTrue(OpAnd(parseFee, OpEq2(reducedLength, MkInteger(1)))) {
			*(order).At(MkString("fee")) = *(reducedFees).At(MkInteger(0))
		}
	}
	if IsTrue(OpEq2(amount, MkUndefined())) {
		if IsTrue(OpAnd(OpNotEq2(filled, MkUndefined()), OpNotEq2(remaining, MkUndefined()))) {
			amount = this.Sum(filled, remaining)
		} else {
			if IsTrue(OpEq2(this.SafeString(order, MkString("status")), MkString("closed"))) {
				amount = OpCopy(filled)
			}
		}
	}
	if IsTrue(OpEq2(filled, MkUndefined())) {
		if IsTrue(OpAnd(OpNotEq2(amount, MkUndefined()), OpNotEq2(remaining, MkUndefined()))) {
			filled = Math.Max(this.Sum(amount, OpNeg(remaining)), MkInteger(0))
		}
	}
	if IsTrue(OpEq2(remaining, MkUndefined())) {
		if IsTrue(OpAnd(OpNotEq2(amount, MkUndefined()), OpNotEq2(filled, MkUndefined()))) {
			remaining = Math.Max(this.Sum(amount, OpNeg(filled)), MkInteger(0))
		}
	}
	if IsTrue(OpEq2(average, MkUndefined())) {
		if IsTrue(OpAnd(OpNotEq2(filled, MkUndefined()), OpAnd(OpNotEq2(cost, MkUndefined()), OpGt(filled, MkInteger(0))))) {
			average = OpDiv(cost, filled)
		}
	}
	costPriceExists := OpOr(OpNotEq2(average, MkUndefined()), OpNotEq2(price, MkUndefined()))
	if IsTrue(OpAnd(parseCost, OpAnd(OpNotEq2(filled, MkUndefined()), costPriceExists))) {
		cost = OpTrinary(OpEq2(average, MkUndefined()), OpMulti(price, filled), OpMulti(average, filled))
	}
	orderType := this.SafeValue(order, MkString("type"))
	emptyPrice := OpOr(OpEq2(price, MkUndefined()), OpEq2(price, MkNumber(0.0)))
	if IsTrue(OpAnd(emptyPrice, OpEq2(orderType, MkString("market")))) {
		price = OpCopy(average)
	}
	return this.Extend(order, MkMap(&VarMap{
		"lastTradeTimestamp": lastTradeTimeTimestamp,
		"price":              price,
		"amount":             amount,
		"cost":               cost,
		"average":            average,
		"filled":             filled,
		"remaining":          remaining,
	}))
}

func (this *ExchangeBase) ParseTimeframe(timeframe *Variant) *Variant {
	amount := timeframe.Slice(MkInteger(0), OpNeg(MkInteger(1)))
	unit := timeframe.Slice(OpNeg(MkInteger(1)))
	scale := OpCopy(MkUndefined())
	if IsTrue(OpEq2(unit, MkString("y"))) {
		scale = OpMulti(MkInteger(60), OpMulti(MkInteger(60), OpMulti(MkInteger(24), MkInteger(365))))
	} else if IsTrue(OpEq2(unit, MkString("M"))) {
		scale = OpMulti(MkInteger(60), OpMulti(MkInteger(60), OpMulti(MkInteger(24), MkInteger(30))))
	} else if IsTrue(OpEq2(unit, MkString("w"))) {
		scale = OpMulti(MkInteger(60), OpMulti(MkInteger(60), OpMulti(MkInteger(24), MkInteger(7))))
	} else if IsTrue(OpEq2(unit, MkString("d"))) {
		scale = OpMulti(MkInteger(60), OpMulti(MkInteger(60), MkInteger(24)))
	} else if IsTrue(OpEq2(unit, MkString("h"))) {
		scale = OpMulti(MkInteger(60), MkInteger(60))
	} else if IsTrue(OpEq2(unit, MkString("m"))) {
		scale = MkInteger(60)
	} else if IsTrue(OpEq2(unit, MkString("s"))) {
		scale = MkInteger(1)
	} else {
		panic(NewNotSupported(MkString("timeframe unit " + unit.ToStr() + " is not supported")))
	}

	return OpMulti(amount, scale)
}

func (this *ExchangeBase) ParseTickers(goArgs ...*Variant) *Variant {
	// transpiled function
	tickers := GoGetArg(goArgs, 0, MkUndefined())
	symbols := GoGetArg(goArgs, 1, MkUndefined())
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	result := MkArray(&VarArray{})
	values := GoGetValues(OpOr(tickers, MkArray(&VarArray{})))
	for i := MkInteger(0); IsTrue(OpLw(i, values.Length)); OpIncr(&i) {
		result.Push(this.Extend(this.VCall("ParseTicker", *(values).At(i)), params)) // defined on exchange
	}
	return this.FilterByArray(result, MkString("symbol"), symbols)
}

func (this *ExchangeBase) ParseTransfers(goArgs ...*Variant) *Variant {
	// transpiled function
	transfers := GoGetArg(goArgs, 0, MkUndefined())
	currency := GoGetArg(goArgs, 1, MkUndefined())
	since := GoGetArg(goArgs, 2, MkUndefined())
	limit := GoGetArg(goArgs, 3, MkUndefined())
	params := GoGetArg(goArgs, 4, MkMap(&VarMap{}))
	if !IsTrue(transfers) {
		transfers = MkArray(&VarArray{})
	}
	result := GoGetValues(transfers).Map(func(transfer *Variant) *Variant {
		return this.Extend(this.VCall("ParseTransfer", transfer, currency), params) // defined on exchange
	})
	result = this.SortBy(result, MkString("timestamp"))
	code := OpTrinary(OpNotEq2(currency, MkUndefined()), *(currency).At(MkString("code")), MkUndefined())
	tail := OpEq2(since, MkUndefined())
	return this.FilterByCurrencySinceLimit(result, code, since, limit, tail)
}

func (this *ExchangeBase) Vwap(goArgs ...*Variant) *Variant {
	// transpiled function
	baseVolume := GoGetArg(goArgs, 0, MkUndefined())
	quoteVolume := GoGetArg(goArgs, 1, MkUndefined())
	return OpTrinary(OpAnd(OpNotEq2(baseVolume, MkUndefined()), OpAnd(OpNotEq2(quoteVolume, MkUndefined()), OpGt(baseVolume, MkInteger(0)))), OpDiv(quoteVolume, baseVolume), MkUndefined())
}

func (this *ExchangeBase) ParseLedger(goArgs ...*Variant) *Variant {
	// transpiled function
	data := GoGetArg(goArgs, 0, MkUndefined())
	currency := GoGetArg(goArgs, 1, MkUndefined())
	since := GoGetArg(goArgs, 2, MkUndefined())
	limit := GoGetArg(goArgs, 3, MkUndefined())
	params := GoGetArg(goArgs, 4, MkMap(&VarMap{}))
	result := MkArray(&VarArray{})
	array := GoGetValues(OpOr(data, MkArray(&VarArray{})))
	for i := MkInteger(0); IsTrue(OpLw(i, array.Length)); OpIncr(&i) {
		itemOrItems := this.VCall("ParseLedgerEntry", *(array).At(i), currency)
		if IsTrue(GoIsArray(itemOrItems)) {
			for j := MkInteger(0); IsTrue(OpLw(j, itemOrItems.Length)); OpIncr(&j) {
				result.Push(this.Extend(*(itemOrItems).At(j), params))
			}
		} else {
			result.Push(this.Extend(itemOrItems, params))
		}
	}
	result = this.SortBy(result, MkString("timestamp"))
	code := OpTrinary(OpNotEq2(currency, MkUndefined()), *(currency).At(MkString("code")), MkUndefined())
	tail := OpEq2(since, MkUndefined())
	return this.FilterByCurrencySinceLimit(result, code, since, limit, tail)
}

func (this *ExchangeBase) PrecisionFromString(str *Variant) *Variant {
	p := len(strings.Trim(str.ToStr(), "0")) - 1
	if p < 0 {
		p = 0
	}
	return MkInteger(int64(p))
}

func (this *ExchangeBase) ParsePrecision(goArgs ...*Variant) *Variant {
	// transpiled function
	precision := GoGetArg(goArgs, 0, MkUndefined())
	if IsTrue(OpEq2(precision, MkUndefined())) {
		return MkUndefined()
	}
	return OpAdd(MkString("1e"), Precise.StringNeg(precision))
}

func (this *ExchangeBase) CostToPrecision(goArgs ...*Variant) *Variant {
	// transpiled function
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	cost := GoGetArg(goArgs, 1, MkUndefined())
	return this.DecimalToPrecision(cost, TRUNCATE, *(*(*(*this.At(MkString("markets"))).At(symbol)).At(MkString("precision"))).At(MkString("price")), *this.At(MkString("precisionMode")), *this.At(MkString("paddingMode")))
}

func (this *ExchangeBase) PriceToPrecision(goArgs ...*Variant) *Variant {
	// transpiled function
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	price := GoGetArg(goArgs, 1, MkUndefined())
	return this.DecimalToPrecision(price, ROUND, *(*(*(*this.At(MkString("markets"))).At(symbol)).At(MkString("precision"))).At(MkString("price")), *this.At(MkString("precisionMode")), *this.At(MkString("paddingMode")))
}

func (this *ExchangeBase) AmountToPrecision(goArgs ...*Variant) *Variant {
	// transpiled function
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	amount := GoGetArg(goArgs, 1, MkUndefined())
	return this.DecimalToPrecision(amount, TRUNCATE, *(*(*(*this.At(MkString("markets"))).At(symbol)).At(MkString("precision"))).At(MkString("amount")), *this.At(MkString("precisionMode")), *this.At(MkString("paddingMode")))
}

func (this *ExchangeBase) FeeToPrecision(goArgs ...*Variant) *Variant {
	// transpiled function
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	fee := GoGetArg(goArgs, 1, MkUndefined())
	return this.DecimalToPrecision(fee, ROUND, *(*(*(*this.At(MkString("markets"))).At(symbol)).At(MkString("precision"))).At(MkString("price")), *this.At(MkString("precisionMode")), *this.At(MkString("paddingMode")))
}

func (this *ExchangeBase) CurrencyToPrecision(goArgs ...*Variant) *Variant {
	// transpiled function
	currency := GoGetArg(goArgs, 0, MkUndefined())
	fee := GoGetArg(goArgs, 1, MkUndefined())
	return this.DecimalToPrecision(fee, ROUND, *(*(*this.At(MkString("currencies"))).At(currency)).At(MkString("precision")), *this.At(MkString("precisionMode")), *this.At(MkString("paddingMode")))
}

func (this *ExchangeBase) CurrencyFromPrecision(goArgs ...*Variant) *Variant {
	// transpiled function
	currency := GoGetArg(goArgs, 0, MkUndefined())
	amount := GoGetArg(goArgs, 1, MkUndefined())
	scale := *(*(*this.At(MkString("currencies"))).At(currency)).At(MkString("precision"))
	return this.FromPrecision(amount, scale)
}

func (this *ExchangeBase) PriceFromPrecision(goArgs ...*Variant) *Variant {
	// transpiled function
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	price := GoGetArg(goArgs, 1, MkUndefined())
	market := *(*this.At(MkString("markets"))).At(symbol)
	wavesPrecision := this.SafeInteger(*this.At(MkString("options")), MkString("wavesPrecision"), MkInteger(8))
	scale := OpSub(wavesPrecision, OpAdd(*(*(market).At(MkString("precision"))).At(MkString("amount")), *(*(market).At(MkString("precision"))).At(MkString("price"))))
	return this.FromPrecision(price, scale)
}

func (this *ExchangeBase) FromPrecision(goArgs ...*Variant) *Variant {
	// transpiled function
	amount := GoGetArg(goArgs, 0, MkUndefined())
	scale := GoGetArg(goArgs, 1, MkUndefined())
	if IsTrue(OpEq2(amount, MkUndefined())) {
		return MkUndefined()
	}
	precise := NewPrecise(amount)
	precise.Decimals = OpAdd(precise.Decimals, scale)
	precise.Reduce()
	return precise.ToString()
}

func (this *ExchangeBase) ToPrecision(goArgs ...*Variant) *Variant {
	// transpiled function
	amount := GoGetArg(goArgs, 0, MkUndefined())
	scale := GoGetArg(goArgs, 1, MkUndefined())
	amountString := amount.ToString()
	_ = amountString
	precise := NewPrecise(amountString)
	_ = precise
	precise.Decimals = OpSub(precise.Decimals, scale)
	precise.Reduce()
	return precise.ToString()
}

func (this *ExchangeBase) ParseOHLCV(goArgs ...*Variant) *Variant {
	// transpiled function
	ohlcv := GoGetArg(goArgs, 0, MkUndefined())
	//market := GoGetArg(goArgs, 1, MkUndefined());
	return OpTrinary(GoIsArray(ohlcv), ohlcv.Slice(MkInteger(0), MkInteger(6)), ohlcv)
}

func (this *ExchangeBase) ParseOHLCVs(goArgs ...*Variant) *Variant {
	// transpiled function
	ohlcvs := GoGetArg(goArgs, 0, MkUndefined())
	market := GoGetArg(goArgs, 1, MkUndefined())
	//timeframe := GoGetArg(goArgs, 2, MkString("1m"));
	since := GoGetArg(goArgs, 3, MkUndefined())
	limit := GoGetArg(goArgs, 4, MkUndefined())
	parsed := ohlcvs.Map(func(ohlcv *Variant) *Variant {
		return this.ParseOHLCV(ohlcv, market)
	})
	sorted := this.SortBy(parsed, MkInteger(0))
	tail := OpEq2(since, MkUndefined())
	return this.FilterBySinceLimit(sorted, since, limit, MkInteger(0), tail)
}

func (this *ExchangeBase) FilterBySinceLimit(goArgs ...*Variant) *Variant {
	// transpiled function
	array := GoGetArg(goArgs, 0, MkUndefined())
	since := GoGetArg(goArgs, 1, MkUndefined())
	limit := GoGetArg(goArgs, 2, MkUndefined())
	key := GoGetArg(goArgs, 3, MkString("timestamp"))
	tail := GoGetArg(goArgs, 4, MkBool(false))
	sinceIsDefined := OpAnd(OpNotEq2(since, MkUndefined()), OpNotEq2(since, MkUndefined()))
	if IsTrue(sinceIsDefined) {
		array = array.Filter(func(entry *Variant) *Variant {
			return OpGtEq(*(entry).At(key), since)
		})
	}
	if IsTrue(OpAnd(OpNotEq2(limit, MkUndefined()), OpNotEq2(limit, MkUndefined()))) {
		array = OpTrinary(tail, array.Slice(OpNeg(limit)), array.Slice(MkInteger(0), limit))
	}
	return array
}

func (this *ExchangeBase) ParseTradingViewOHLCV(goArgs ...*Variant) *Variant {
	// transpiled function
	ohlcvs := GoGetArg(goArgs, 0, MkUndefined())
	market := GoGetArg(goArgs, 1, MkUndefined())
	timeframe := GoGetArg(goArgs, 2, MkString("1m"))
	since := GoGetArg(goArgs, 3, MkUndefined())
	limit := GoGetArg(goArgs, 4, MkUndefined())
	result := this.ConvertTradingViewToOHLCV(ohlcvs)
	return this.ParseOHLCVs(result, market, timeframe, since, limit)
}

func (this *ExchangeBase) ConvertTradingViewToOHLCV(goArgs ...*Variant) *Variant {
	// transpiled function
	ohlcvs := GoGetArg(goArgs, 0, MkUndefined())
	t := GoGetArg(goArgs, 1, MkString("t"))
	o := GoGetArg(goArgs, 2, MkString("o"))
	h := GoGetArg(goArgs, 3, MkString("h"))
	l := GoGetArg(goArgs, 4, MkString("l"))
	c := GoGetArg(goArgs, 5, MkString("c"))
	v := GoGetArg(goArgs, 6, MkString("v"))
	ms := GoGetArg(goArgs, 7, MkBool(false))
	result := MkArray(&VarArray{})
	for i := MkInteger(0); IsTrue(OpLw(i, *(*(ohlcvs).At(t)).At(MkString("length")))); OpIncr(&i) {
		result.Push(MkArray(&VarArray{
			OpTrinary(ms, *(*(ohlcvs).At(t)).At(i), OpMulti(*(*(ohlcvs).At(t)).At(i), MkInteger(1000))),
			*(*(ohlcvs).At(o)).At(i),
			*(*(ohlcvs).At(h)).At(i),
			*(*(ohlcvs).At(l)).At(i),
			*(*(ohlcvs).At(c)).At(i),
			*(*(ohlcvs).At(v)).At(i),
		}))
	}
	return result
}

func (this *ExchangeBase) LoadAccounts(goArgs ...*Variant) *Variant {
	// transpiled function
	reload := GoGetArg(goArgs, 0, MkBool(false))
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	if IsTrue(reload) {
		*this.At(MkString("accounts")) = this.VCall("FetchAccounts", params) // defined on exchange
	} else {
		if IsTrue(*this.At(MkString("accounts"))) {
			return *this.At(MkString("accounts"))
		} else {
			*this.At(MkString("accounts")) = this.VCall("FetchAccounts", params) // defined on exchange
		}
	}
	*this.At(MkString("accountsById")) = this.IndexBy(*this.At(MkString("accounts")), MkString("id"))
	return *this.At(MkString("accounts"))
}

func (this *ExchangeBase) BuildOHLCVC(goArgs ...*Variant) *Variant {
	// transpiled function
	trades := GoGetArg(goArgs, 0, MkUndefined())
	timeframe := GoGetArg(goArgs, 1, MkString("1m"))
	since := GoGetArg(goArgs, 2, MkNumber(math.Inf(-1)))
	limit := GoGetArg(goArgs, 3, MkNumber(math.Inf(1)))
	ms := OpMulti(this.ParseTimeframe(timeframe), MkInteger(1000))
	ohlcvs := MkArray(&VarArray{})
	timestamp := MkUndefined()
	open := MkUndefined()
	high := MkUndefined()
	low := MkUndefined()
	close := MkUndefined()
	volume := MkUndefined()
	count := MkUndefined()
	ArrayUnpack(MkArray(&VarArray{
		MkInteger(0),
		MkInteger(1),
		MkInteger(2),
		MkInteger(3),
		MkInteger(4),
		MkInteger(5),
		MkInteger(6),
	}), &timestamp, &open, &high, &low, &close, &volume, &count)
	oldest := Math.Min(OpSub(trades.Length, MkInteger(1)), limit)
	for i := MkInteger(0); IsTrue(OpLwEq(i, oldest)); OpIncr(&i) {
		trade := *(trades).At(i)
		if IsTrue(OpLw(*trade.At(MkString("timestamp")), since)) {
			continue
		}
		openingTime := OpMulti(Math.Floor(OpDiv(*trade.At(MkString("timestamp")), ms)), ms)
		candle := OpSub(ohlcvs.Length, MkInteger(1))
		if IsTrue(OpOr(OpEq2(candle, OpNeg(MkInteger(1))), OpGtEq(openingTime, OpAdd(*(*(ohlcvs).At(candle)).At(timestamp), ms)))) {
			ohlcvs.Push(MkArray(&VarArray{
				openingTime,
				*trade.At(MkString("price")),
				*trade.At(MkString("price")),
				*trade.At(MkString("price")),
				*trade.At(MkString("price")),
				*trade.At(MkString("Amount")),
				MkInteger(1),
			}))
		} else {
			*(*(ohlcvs).At(candle)).At(high) = Math.Max(*(*(ohlcvs).At(candle)).At(high), *trade.At(MkString("price")))
			*(*(ohlcvs).At(candle)).At(low) = Math.Min(*(*(ohlcvs).At(candle)).At(low), *trade.At(MkString("price")))
			*(*(ohlcvs).At(candle)).At(close) = OpCopy(*trade.At(MkString("price")))
			OpAddAssign(&*(*(ohlcvs).At(candle)).At(volume), *trade.At(MkString("Amount")))
			OpIncr(&*(*(ohlcvs).At(candle)).At(count))
		}
	}
	return ohlcvs
}

var mt_rand_init = false

func mt_rand(min int, max int) int {
	if !mt_rand_init {
		rand.Seed(time.Now().UnixNano())
		mt_rand_init = true
	}
	return rand.Intn(max-min) + min
}

func (this *ExchangeBase) Uuid() *Variant {
	return MkString(fmt.Sprintf("%04x%04x-%04x-%04x-%04x-%04x%04x%04x",
		// 32 bits for "time_low"
		mt_rand(0, 0xffff), mt_rand(0, 0xffff),

		// 16 bits for "time_mid"
		mt_rand(0, 0xffff),

		// 16 bits for "time_hi_and_version",
		// four most significant bits holds version number 4
		mt_rand(0, 0x0fff)|0x4000,

		// 16 bits, 8 bits for "clk_seq_hi_res", 8 bits for "clk_seq_low",
		// two most significant bits holds zero and one for variant DCE1.1
		mt_rand(0, 0x3fff)|0x8000,

		// 48 bits for "node"
		mt_rand(0, 0xffff), mt_rand(0, 0xffff), mt_rand(0, 0xffff)))
}

func (this *ExchangeBase) Uuid16() *Variant {
	return MkString(fmt.Sprintf("%08x%08x", mt_rand(0, 0xffffffff), mt_rand(0, 0xffffffff)))
}

func (this *ExchangeBase) Uuid22() *Variant {
	return MkString(fmt.Sprintf("%08x%08x%06x", mt_rand(0, 0xffffffff), mt_rand(0, 0xffffffff), mt_rand(0, 0xffffff)))
}

func (this *ExchangeBase) ParseDepositAddresses(goArgs ...*Variant) *Variant {
	addresses := GoGetArg(goArgs, 0, MkUndefined())
	codes := GoGetArg(goArgs, 1, MkUndefined())
	result := MkArray(&VarArray{})
	for i := MkInteger(0); IsTrue(OpLw(i, addresses.Length)); OpIncr(&i) {
		address := this.VCall("ParseDepositAddress", *(addresses).At(i)) // defined on exchange
		result.Push(address)
	}
	if IsTrue(codes) {
		result = this.FilterByArray(result, MkString("currency"), codes)
	}
	return this.IndexBy(result, MkString("currency"))
}

func (this *ExchangeBase) ParseBidAsk(goArgs ...*Variant) *Variant {
	// transpiled function
	bidask := GoGetArg(goArgs, 0, MkUndefined())
	priceKey := GoGetArg(goArgs, 1, MkInteger(0))
	amountKey := GoGetArg(goArgs, 2, MkInteger(1))
	price := this.SafeNumber(bidask, priceKey)
	amount := this.SafeNumber(bidask, amountKey)
	return MkArray(&VarArray{
		price,
		amount,
	})
}

func (this *ExchangeBase) ParseBidsAsks(goArgs ...*Variant) *Variant {
	bidasks := GoGetArg(goArgs, 0, MkUndefined())
	priceKey := GoGetArg(goArgs, 1, MkInteger(0))
	amountKey := GoGetArg(goArgs, 2, MkInteger(1))
	if !IsTrue(bidasks) {
		bidasks = MkArray(&VarArray{})
	}
	return bidasks.Map(func(bidask *Variant) *Variant {
		return this.ParseBidAsk(bidask, priceKey, amountKey)
	})
}

func (this *ExchangeBase) CommonCurrencyCode(goArgs ...*Variant) *Variant {
	// transpiled function
	currency := GoGetArg(goArgs, 0, MkUndefined())
	if IsTrue(OpNot(*this.At(MkString("substituteCommonCurrencyCodes")))) {
		return currency
	}
	return this.SafeString(*this.At(MkString("commonCurrencies")), currency, currency)
}

func (this *ExchangeBase) CalculateFee(goArgs ...*Variant) *Variant {
	// transpiled function
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	//type_ := GoGetArg(goArgs, 1, MkUndefined());
	side := GoGetArg(goArgs, 2, MkUndefined())
	amount := GoGetArg(goArgs, 3, MkUndefined())
	price := GoGetArg(goArgs, 4, MkUndefined())
	takerOrMaker := GoGetArg(goArgs, 5, MkString("taker"))
	//params := GoGetArg(goArgs, 6, MkMap(&VarMap{}));
	market := *(*this.At(MkString("markets"))).At(symbol)
	feeSide := this.SafeString(market, MkString("feeSide"), MkString("quote"))
	key := MkString("quote")
	cost := OpCopy(MkUndefined())
	if IsTrue(OpEq2(feeSide, MkString("quote"))) {
		cost = OpMulti(amount, price)
	} else {
		if IsTrue(OpEq2(feeSide, MkString("base"))) {
			cost = OpCopy(amount)
		} else {
			if IsTrue(OpEq2(feeSide, MkString("get"))) {
				cost = OpCopy(amount)
				if IsTrue(OpEq2(side, MkString("sell"))) {
					OpMultiAssign(&cost, price)
				} else {
					key = MkString("base")
				}
			} else {
				if IsTrue(OpEq2(feeSide, MkString("give"))) {
					cost = OpCopy(amount)
					if IsTrue(OpEq2(side, MkString("buy"))) {
						OpMultiAssign(&cost, price)
					} else {
						key = MkString("base")
					}
				}
			}
		}
	}
	rate := *(market).At(takerOrMaker)
	if IsTrue(OpNotEq2(cost, MkUndefined())) {
		OpMultiAssign(&cost, rate)
	}
	return MkMap(&VarMap{
		"type":     takerOrMaker,
		"currency": *(market).At(key),
		"rate":     rate,
		"cost":     cost,
	})
}

func (this *ExchangeBase) FilterBySymbol(goArgs ...*Variant) *Variant {
	array := GoGetArg(goArgs, 0, MkUndefined())
	symbol := GoGetArg(goArgs, 1, MkUndefined())
	return OpTrinary(OpNotEq2(symbol, MkUndefined()), array.Filter(func(entry *Variant) *Variant {
		return OpEq2(*entry.At(MkString("symbol")), symbol)
	}), array)
}

func (this *ExchangeBase) FuturesTransfer(goArgs ...*Variant) *Variant {
	// transpiled function
	code := GoGetArg(goArgs, 0, MkUndefined())
	amount := GoGetArg(goArgs, 1, MkUndefined())
	type_ := GoGetArg(goArgs, 2, MkUndefined())
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	if IsTrue(OpOr(OpLw(type_, MkInteger(1)), OpGt(type_, MkInteger(4)))) {
		panic(NewArgumentsRequired(OpAdd(*this.At(MkString("id")), MkString(" type must be between 1 and 4"))))
	}
	this.LoadMarkets()
	currency := this.Currency(code)
	request := MkMap(&VarMap{
		"asset":  *(currency).At(MkString("id")),
		"amount": amount,
		"type":   type_,
	})
	// this function is only defined on binance
	response := this.VCall("SapiPostFuturesTransfer", this.Extend(request, params))
	return this.VCall("ParseTransfer", response, currency)
}

func (this *ExchangeBase) Remove0xPrefix(goArgs ...*Variant) *Variant {
	// transpiled function
	hexData := GoGetArg(goArgs, 0, MkUndefined())
	if IsTrue(OpEq2(hexData.Slice(MkInteger(0), MkInteger(2)), MkString("0x"))) {
		return hexData.Slice(MkInteger(2))
	} else {
		return hexData
	}
	return MkUndefined()
}

func (this *ExchangeBase) CheckRequiredDependencies(goArgs ...*Variant) *Variant {
	return MkUndefined() // default implementation doe snto do anything
}

///////////////////////////////////////////////////////////////////////////////////////////////
// convinence functions for the caller, not used by the exchanges directly

func (this *ExchangeBase) CreateLimitOrder(goArgs ...*Variant) *Variant {
	// transpiled function
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	side := GoGetArg(goArgs, 1, MkUndefined())
	amount := GoGetArg(goArgs, 2, MkUndefined())
	price := GoGetArg(goArgs, 3, MkUndefined())
	params := GoGetArg(goArgs, 4, MkMap(&VarMap{}))
	return this.VCall("CreateOrder", symbol, MkString("limit"), side, amount, price, params)
}

func (this *ExchangeBase) CreateMarketOrder(goArgs ...*Variant) *Variant {
	// transpiled function
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	side := GoGetArg(goArgs, 1, MkUndefined())
	amount := GoGetArg(goArgs, 2, MkUndefined())
	price := GoGetArg(goArgs, 3, MkUndefined())
	params := GoGetArg(goArgs, 4, MkMap(&VarMap{}))
	return this.VCall("CreateOrder", symbol, MkString("market"), side, amount, price, params)
}

func (this *ExchangeBase) CreateLimitBuyOrder(goArgs ...*Variant) *Variant {
	// transpiled function
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	amount := GoGetArg(goArgs, 1, MkUndefined())
	price := GoGetArg(goArgs, 2, MkUndefined())
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	return this.VCall("CreateOrder", symbol, MkString("limit"), MkString("buy"), amount, price, params)
}

func (this *ExchangeBase) CreateLimitSellOrder(goArgs ...*Variant) *Variant {
	// transpiled function
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	amount := GoGetArg(goArgs, 1, MkUndefined())
	price := GoGetArg(goArgs, 2, MkUndefined())
	params := GoGetArg(goArgs, 3, MkMap(&VarMap{}))
	return this.VCall("CreateOrder", symbol, MkString("limit"), MkString("sell"), amount, price, params)
}

func (this *ExchangeBase) CreateMarketBuyOrder(goArgs ...*Variant) *Variant {
	// transpiled function
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	amount := GoGetArg(goArgs, 1, MkUndefined())
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	return this.VCall("CreateOrder", symbol, MkString("market"), MkString("buy"), amount, MkUndefined(), params)
}

func (this *ExchangeBase) CreateMarketSellOrder(goArgs ...*Variant) *Variant {
	// transpiled function
	symbol := GoGetArg(goArgs, 0, MkUndefined())
	amount := GoGetArg(goArgs, 1, MkUndefined())
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	return this.VCall("CreateOrder", symbol, MkString("market"), MkString("sell"), amount, MkUndefined(), params)
}

func (this *ExchangeBase) FetchTickers(goArgs ...*Variant) *Variant {
	panic(NewNotSupported(OpAdd(*this.At(MkString("id")), MkString(" fetchTickers not supported yet"))))
	return MkUndefined()
}

func (this *ExchangeBase) FetchOrder(goArgs ...*Variant) *Variant {
	panic(NewNotSupported(OpAdd(*this.At(MkString("id")), MkString(" fetchOrder not supported yet"))))
	return MkUndefined()
}

func (this *ExchangeBase) FetchUnifiedOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	return this.VCall("FetchOrder", this.SafeValue(order, MkString("id")), this.SafeValue(order, MkString("symbol")), params)
}

func (this *ExchangeBase) CreateOrder(goArgs ...*Variant) *Variant {
	panic(NewNotSupported(OpAdd(*this.At(MkString("id")), MkString(" createOrder not supported yet"))))
	return MkUndefined()
}

func (this *ExchangeBase) CancelOrder(goArgs ...*Variant) *Variant {
	panic(NewNotSupported(OpAdd(*this.At(MkString("id")), MkString(" cancelOrder not supported yet"))))
	return MkUndefined()
}

func (this *ExchangeBase) CancelUnifiedOrder(goArgs ...*Variant) *Variant {
	order := GoGetArg(goArgs, 0, MkUndefined())
	params := GoGetArg(goArgs, 1, MkMap(&VarMap{}))
	return this.VCall("CancelOrder", this.SafeValue(order, MkString("id")), this.SafeValue(order, MkString("symbol")), params)
}

func (this *ExchangeBase) FetchOrders(goArgs ...*Variant) *Variant {
	panic(NewNotSupported(OpAdd(*this.At(MkString("id")), MkString(" fetchOrders not supported yet"))))
	return MkUndefined()
}

func (this *ExchangeBase) FetchOpenOrders(goArgs ...*Variant) *Variant {
	panic(NewNotSupported(OpAdd(*this.At(MkString("id")), MkString(" fetchOpenOrders not supported yet"))))
	return MkUndefined()
}

func (this *ExchangeBase) FetchClosedOrders(goArgs ...*Variant) *Variant {
	panic(NewNotSupported(OpAdd(*this.At(MkString("id")), MkString(" fetchClosedOrders not supported yet"))))
	return MkUndefined()
}

func (this *ExchangeBase) FetchMyTrades(goArgs ...*Variant) *Variant {
	panic(NewNotSupported(OpAdd(*this.At(MkString("id")), MkString(" fetchMyTrades not supported yet"))))
	return MkUndefined()
}

func (this *ExchangeBase) FetchTransactions(goArgs ...*Variant) *Variant {
	panic(NewNotSupported(OpAdd(*this.At(MkString("id")), MkString(" fetchTransactions not supported yet"))))
	return MkUndefined()
}

func (this *ExchangeBase) FetchDeposits(goArgs ...*Variant) *Variant {
	panic(NewNotSupported(OpAdd(*this.At(MkString("id")), MkString(" fetchDeposits not supported yet"))))
	return MkUndefined()
}

func (this *ExchangeBase) FetchWithdrawals(goArgs ...*Variant) *Variant {
	panic(NewNotSupported(OpAdd(*this.At(MkString("id")), MkString(" fetchWithdrawals not supported yet"))))
	return MkUndefined()
}
