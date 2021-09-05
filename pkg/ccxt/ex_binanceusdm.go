package ccxt

import ()

type Binanceusdm struct {
	*ExchangeBase
}

var _ Exchange = (*Binanceusdm)(nil)

func init() {
	exchange := &Binanceusdm{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Binanceusdm) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("binanceusdm"),
		"name": MkString("Binance USDâ-M"),
		"urls": MkMap(&VarMap{
			"logo": MkString("https://user-images.githubusercontent.com/1294454/117738721-668c8d80-b205-11eb-8c49-3fad84c4a07f.jpg"),
			"doc": MkArray(&VarArray{
				MkString("https://binance-docs.github.io/apidocs/futures/en/"),
				MkString("https://binance-docs.github.io/apidocs/spot/en"),
			}),
		}),
		"options": MkMap(&VarMap{
			"defaultType":      MkString("future"),
			"leverageBrackets": MkUndefined(),
			"marginTypes":      MkMap(&VarMap{}),
		}),
	}))
}

func (this *Binanceusdm) TransferIn(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	amount := GoGetArg(goArgs, 1, MkUndefined())
	_ = amount
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	return this.FuturesTransfer(code, amount, MkInteger(1), params)
}

func (this *Binanceusdm) TransferOut(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	amount := GoGetArg(goArgs, 1, MkUndefined())
	_ = amount
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	return this.FuturesTransfer(code, amount, MkInteger(2), params)
}
