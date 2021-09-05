package ccxt

import ()

type Binancecoinm struct {
	*ExchangeBase
}

var _ Exchange = (*Binancecoinm)(nil)

func init() {
	exchange := &Binancecoinm{}
	Exchanges = append(Exchanges, exchange)
}

func (this *Binancecoinm) Describe(goArgs ...*Variant) *Variant {
	return this.DeepExtend(this.BaseDescribe(), MkMap(&VarMap{
		"id":   MkString("binancecoinm"),
		"name": MkString("Binance COIN-M"),
		"urls": MkMap(&VarMap{
			"logo": MkString("https://user-images.githubusercontent.com/1294454/117738721-668c8d80-b205-11eb-8c49-3fad84c4a07f.jpg"),
			"doc": MkArray(&VarArray{
				MkString("https://binance-docs.github.io/apidocs/delivery/en/"),
				MkString("https://binance-docs.github.io/apidocs/spot/en"),
			}),
		}),
		"options": MkMap(&VarMap{
			"defaultType":      MkString("delivery"),
			"leverageBrackets": MkUndefined(),
		}),
	}))
}

func (this *Binancecoinm) TransferIn(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	amount := GoGetArg(goArgs, 1, MkUndefined())
	_ = amount
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	return this.FuturesTransfer(code, amount, MkInteger(3), params)
}

func (this *Binancecoinm) TransferOut(goArgs ...*Variant) *Variant {
	code := GoGetArg(goArgs, 0, MkUndefined())
	_ = code
	amount := GoGetArg(goArgs, 1, MkUndefined())
	_ = amount
	params := GoGetArg(goArgs, 2, MkMap(&VarMap{}))
	_ = params
	return this.FuturesTransfer(code, amount, MkInteger(4), params)
}
